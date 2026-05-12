package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
)

// 🟡 Database Transactions -> Units of works
// ------------------------------------------

// EXAMPLE:
// 1. User creates a bank-account
// 2. Create a wallet for the user
// 3. Want to top-up wallet for the user
// 4. Want to write a transaction-log

var schema = `
CREATE TABLE IF NOT EXISTS profile (
	user_id INTEGER PRIMARY KEY REFERENCES users(user_id) ON DELETE CASCADE,
	avatar TEXT NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
)
`

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	HashedPassword string `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	Profile Profile `json:"profile"`
}

type Profile struct {
	UserId  int `json:"user_id"`
	Avatar string `json:"avatar"`
	Created time.Time `json:"created"`
}

func main() {
	dbName:="users_db.db"

	db,err:= sql.Open("sqlite",dbName)

	if err!=nil{
		log.Fatal("ERROR: ",err.Error())
	}

	defer db.Close() 

	err=db.Ping() // verification
	if err!=nil{
		log.Fatal("DB-CONNECTION ERROR: ",err.Error())
	}

	user,err:=GetUserByEmail(db,"tx2Localhost")
	if err!=nil{
		log.Fatal("ERROR: ",err.Error())
	}

	userBs,err:=json.MarshalIndent(user,""," ")
	if err!=nil{
		log.Fatal("ERROR: ",err.Error())
	}

	fmt.Printf("user:\n %+v\n",string(userBs))

	// CreateTable(db)
	// userId,err:=CreateUser(db, "User2 from Tx2","tx2Localhost","**345jajaja","http://avatar2.com/user2")
	// if err!=nil{
	// 	log.Fatal("ERROR: ",err.Error())
	// }

	// fmt.Println("✅ User created. User_id:",userId)

}

func CreateTable(db *sql.DB){
	_,err:=db.Exec(schema)
	if err!=nil{
		log.Fatal("SQL-ERROR: ",err.Error())
	}
}

// 3 phases of transactions -> Begin, RollBack or Commit
func CreateUser(db *sql.DB, name, email, hashed_password, avatar string)(int64,error){

	ctx:=context.Background()

	// begin the tx
	tx,err:=db.BeginTx(ctx,nil)
	if err!=nil{
		return 0,err
	}

	defer tx.Rollback()


	//query/stmt
	stmt,err:=tx.PrepareContext(ctx,`INSERT INTO users (name,email,hashed_password)  VALUES (?, ?, ?)`)
	if err!=nil{
		return 0,err
	}

	defer stmt.Close()

	//! hash the password first
	hp,err:=bcrypt.GenerateFromPassword([]byte(hashed_password),bcrypt.DefaultCost)
	if err!=nil{
		return 0,err
	}

	res,err:= stmt.Exec(name,email,string(hp))
	if err!=nil{
		return 0,err
	}

	userId,err:= res.LastInsertId()
	if err!=nil{
		return 0,err
	}

	// Now, create the profile

	profileStmt,err:=tx.PrepareContext(ctx, `INSERT INTO profile (user_id,avatar) VALUES(?, ?)`)

	if err!=nil{
		return 0,err
	}

	defer profileStmt.Close()
	_,err=profileStmt.Exec(userId,avatar)
	if err!=nil{
		return 0,err
	}

	// At the end, commit the tx.
	err = tx.Commit()
	if err!=nil{
		return 0,err
	}

	return userId,nil
}

func GetUserByEmail(db *sql.DB, email string)(*User,error){
	stmt:=`SELECT u.id,u.name,u.email,u.hashed_password,u.created_at, p.avatar FROM users u INNER JOIN profile p ON u.id = p.user_id WHERE u.email = ?`

	row:=db.QueryRow(stmt,email)

	// scan the data back into the struct
	var user User

	err:=row.Scan(&user.ID,&user.Name,&user.Email,&user.HashedPassword, &user.CreatedAt, &user.Profile.Avatar)

	if err!=nil{
		if err==sql.ErrNoRows{
			log.Fatal("No records found:",err.Error())
		}
		return nil,err
	}

	user.Profile.UserId=user.ID
	return &user, nil
}

// $ go run main.go
// user:
//  {
//  "id": 7,
//  "name": "User2 from Tx2",
//  "email": "tx2Localhost",
//  "created_at": "2026-05-12T09:58:16Z",
//  "profile": {
//   "user_id": 7,
//   "avatar": "http://avatar2.com/user2",
//   "created": "0001-01-01T00:00:00Z"
//  }
// }