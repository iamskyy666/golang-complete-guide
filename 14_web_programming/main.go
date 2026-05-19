package main

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

// Template Layouts And CSS
type application struct{
	errorLog *log.Logger
	infoLog  *log.Logger
	userRepo UserRepo
	templateDir string
	tp *TmpltRenderer
	publicPath string
}

func main() {

	 db,err:= ConnectDb("users_database.db")
	 if err!=nil{
		log.Fatal("ERROR: ",err.Error())
	}
	defer db.Close()
	 
	 app:= &application{
		errorLog: log.New(os.Stderr, "ERROR\t",log.Ltime | log.LstdFlags | log.Lmicroseconds | log.Lshortfile),
		infoLog: log.New(os.Stderr, "ERROR\t",log.Ltime | log.LstdFlags),
		userRepo: NewSQLUserRepository(db),
		templateDir: "./14_web_programming/templates",
		publicPath: "14_web_programming/public",
	 }

	 app.tp = NewTemplateRenderer(app.templateDir,true) // template rendering with cache

	 log.Println("Server listening on PORT: 8080 ✅")
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