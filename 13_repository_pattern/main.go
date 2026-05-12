package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/iamskyy666/db_programming/13_repository_pattern/repository"
	_ "modernc.org/sqlite"
)


var schema = `
CREATE TABLE IF NOT EXISTS profile (
	user_id INTEGER PRIMARY KEY REFERENCES users(user_id) ON DELETE CASCADE,
	avatar TEXT NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
)
`

func main() {
	dbName:="users_db.db"
	db,err:=ConnectDb(dbName)
	CheckErr(err)
	 defer db.Close() // don't forget
	fmt.Println("Connected to DB ✅")

	repo:=repository.NewSQLUserRepository(db)

	user,err:=repo.CreateUser("User from Repo","repouser@test.com","123456@abcdef","avatar12")
	CheckErr(err)

	fmt.Println("user:",user)
	
}

func CheckErr(err error){
	if err != nil {
		log.Fatal("ERROR: ",err)
	}
}

func ConnectDb(name string)(*sql.DB,error){
db,err:= sql.Open("sqlite",name)

	if err!=nil{
		log.Fatal("ERROR: ",err.Error())
	}

	err=db.Ping() // verification
	if err!=nil{
		log.Fatal("DB-CONNECTION ERROR: ",err.Error())
	}
	return db,nil
}



func CreateTable(db *sql.DB){
	_,err:=db.Exec(schema)
	if err!=nil{
		log.Fatal("SQL-ERROR: ",err.Error())
	}
}

