package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "modernc.org/sqlite"
)

// DEPENDANCY INJECTION
type application struct{
	errorLog *log.Logger
	infoLog  *log.Logger
	userRepo UserRepo
	mux *http.ServeMux
}

func main() {

	 mux:=http.NewServeMux()

	 db,err:= ConnectDb("users_database.db")
	 if err!=nil{
		log.Fatal("ERROR: ",err.Error())
	}
	defer db.Close()
	 
	 app:= &application{
		errorLog: log.New(os.Stderr, "ERROR\t",log.Ltime | log.LstdFlags | log.Lmicroseconds | log.Lshortfile),
		infoLog: log.New(os.Stderr, "ERROR\t",log.Ltime | log.LstdFlags),
		userRepo: NewSQLUserRepository(db),
		mux: mux,

	 }

	 log.Println("Server listening on PORT: 8080 ✅")

	 app.mount(mux)

	if err:= app.Serve(); err!=nil{
		log.Fatal("ERROR: ",err.Error())
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