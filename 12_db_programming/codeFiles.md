```go 
// REVISED VERSION, BECAUSE OF C-COMPILER ISSUE ⚠️
// connecting to a DATABASE and TABLE Creation
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	hashed_password BLOB NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
)
`

func main() {

	dbName:="data.db"
	_=os.Remove(dbName) // remove each time with a fresh start (don't so this in prod.)

	db,err:= sql.Open("sqlite",dbName)

	if err!=nil{
		log.Fatal("ERROR: ",err.Error())
	}

	defer func ()  {
		fmt.Println("Closing DB-Connection!")
		if err:=db.Close();err!=nil{
			log.Printf("ERROR closing DB connection: %v",err.Error())
		}
	}() // like every io ops.

	err=db.Ping() // verification
	if err!=nil{
		log.Fatal("DB-CONNECTION ERROR: ",err.Error())
	}

	fmt.Println("Connected to DB ✅")

	_,err=db.Exec(schema)
	if err!=nil{
		log.Fatal("SQL-ERROR: ",err.Error())
	}

	fmt.Println("Table was created ☑️")
}

// $ go run main.go
// Connected to DB ✅
// Table was created ☑️
// Closing DB-Connection!

```
```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
)

// Inserting RECORDS into the DATABASE

var schema = `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	hashed_password TEXT NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
)
`

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

	fmt.Println("Connected to DB ✅")

	// CreateTable(db)

	// fmt.Println("Table was created ☑️")

	lastId,err:= CreateUser(db,"Skyy Banerjee","skyy_banerjee@test.com","abcd@1234")
	if err!=nil{
		log.Fatal("SQL-ERROR: ",err.Error())
	}

	fmt.Println("LAST USER_ID:",lastId)

	lastId,err= CreateUser(db,"Soumadip Benerjee","soumadip@test.com","abcd@####")
	if err!=nil{
		log.Fatal("SQL-ERROR: ",err.Error())
	}

	fmt.Println("LAST USER_ID:",lastId)

	lastId,err= CreateUser(db,"Bruce Wayne","bruce@test.com","bruce@1234")
	if err!=nil{
		log.Fatal("SQL-ERROR: ",err.Error())
	}

	fmt.Println("LAST USER_ID:",lastId)
}

func CreateTable(db *sql.DB){
	_,err:=db.Exec(schema)
	if err!=nil{
		log.Fatal("SQL-ERROR: ",err.Error())
	}
}

func CreateUser(db *sql.DB, name, email, hashed_password string)(int64,error){
	//query
	stmt:=`INSERT INTO users (name,email,hashed_password)  VALUES (?, ?, ?)`

	//! hash the password first
	hp,err:=bcrypt.GenerateFromPassword([]byte(hashed_password),bcrypt.DefaultCost)
	if err!=nil{
		return 0,err
	}

	res,err:= db.Exec(stmt,name,email,string(hp))
	if err!=nil{
		// log.Fatal("SQL-ERROR: ",err.Error())
		return 0,err
	}

	return res.LastInsertId()
}

// Connected to DB ✅
// LAST USER_ID: 1
// LAST USER_ID: 2
// LAST USER_ID: 3
// LAST USER_ID: 4

```
```go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
)

// Fetching RECORDS from the DB

var schema = `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	hashed_password TEXT NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
)
`

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	HashedPassword string `json:"-"`
	CreatedAt time.Time `json:"created_at"`
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

	// CreateTable(db)

	skyy,err:=GetUserByEmail(db,"skyy@test.com")
	if err!=nil{
		log.Fatal("SQL-ERROR: ",err.Error())
	}

	// byte-slice
	skyyBS,err:=json.MarshalIndent(skyy,""," ")
	if err!=nil{
		log.Fatal("ERROR: ",err.Error())
	}

	fmt.Printf("Skyy Details: \n%v\n",string(skyyBS))

	allUsers,err:=GetAllUsers(db)
	if err!=nil{
		log.Fatal("ERROR: ",err.Error())
	}

	// byte-slice
	allUsersBS,err:=json.MarshalIndent(allUsers,""," ")

	fmt.Printf("All Users: \n%v\n",string(allUsersBS))

}

// Controllers

func CreateTable(db *sql.DB){
	_,err:=db.Exec(schema)
	if err!=nil{
		log.Fatal("SQL-ERROR: ",err.Error())
	}
}

func CreateUser(db *sql.DB, name, email, hashed_password string)(int64,error){
	//query
	stmt:=`INSERT INTO users (name,email,hashed_password)  VALUES (?, ?, ?)`

	//! hash the password first
	hp,err:=bcrypt.GenerateFromPassword([]byte(hashed_password),bcrypt.DefaultCost)
		if err!=nil{
		return 0,err
	}

	res,err:= db.Exec(stmt,name,email,string(hp))
	if err!=nil{
		// log.Fatal("SQL-ERROR: ",err.Error())
		return 0,err
	}
	return res.LastInsertId()
}

func GetUserByEmail(db *sql.DB, email string)(*User,error){
	stmt:=`SELECT id,name,email,hashed_password,created_at FROM users WHERE email = ?`

	row:=db.QueryRow(stmt,email)

	// scan the data back into the struct
	var user User

	err:=row.Scan(&user.ID,&user.Name,&user.Email,&user.HashedPassword, &user.CreatedAt)

	if err!=nil{
		if err==sql.ErrNoRows{
			log.Fatal("No records found:",err.Error())
		}
		return nil,err
	}

	return &user, nil
}

func GetAllUsers(db *sql.DB)([]User,error){
	stmt:=`SELECT id,name,email,hashed_password,created_at FROM users`

	rows,err:=db.Query(stmt)

	if err!=nil{
		return nil,err
	}

	defer rows.Close()

	var users []User
		for rows.Next(){
			var user User
			if err:=rows.Scan(&user.ID,&user.Name,&user.Email,&user.HashedPassword, &user.CreatedAt); err!=nil{
			return nil,err
			}
			users = append(users, user)
		
		}
		if err:=rows.Err();err!=nil{
			return nil,err
		}
		return users,nil

}

// $ go run main.go
//
//--------------------------------------------------------------
//
// Skyy Details: 
// {
//  "id": 1,
//  "name": "Skyy",
//  "email": "skyy@test.com",
//  "created_at": "2026-05-12T05:48:57Z"
// }
//
// --------------------------------------------------------------
//
// All Users: 
// [
//  {
//   "id": 1,
//   "name": "Skyy",
//   "email": "skyy@test.com",
//   "created_at": "2026-05-12T05:48:57Z"
//  },
//  {
//   "id": 2,
//   "name": "Skyy Banerjee",
//   "email": "skyy_banerjee@test.com",
//   "created_at": "2026-05-12T05:52:51Z"
//  },
//  {
//   "id": 3,
//   "name": "Soumadip Benerjee",
//   "email": "soumadip@test.com",
//   "created_at": "2026-05-12T05:52:51Z"
//  },
//  {
//   "id": 4,
//   "name": "Bruce Wayne",
//   "email": "bruce@test.com",
//   "created_at": "2026-05-12T05:52:52Z"
//  }
// ]
```

```go
func main(){
//! Create With Prepare( )
	lastId,err:= CreateUserWithPrepare(db,"Alfred Pennywise","butler@wayne-enterprise.com","batman@****")
	if err!=nil{
		log.Fatal("SQL-ERROR: ",err.Error())
	}
}

// PREPARE stmt.
func CreateUserWithPrepare(db *sql.DB, name, email, hashed_password string)(int64,error){
	//query
	stmt,err:=db.Prepare(`INSERT INTO users (name,email,hashed_password)  VALUES (?, ?, ?)`)
	if err!=nil{
		return 0,err
		
	}

	defer stmt.Close()

	hp,err:=bcrypt.GenerateFromPassword([]byte(hashed_password),bcrypt.DefaultCost)
		if err!=nil{
		return 0,err
	}

	res,err:= stmt.Exec(name,email,string(hp))
	if err!=nil{
		return 0,err
	}
	return res.LastInsertId()
}

// $ go run main.go
// Last_ID: 5

```
```go
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
```