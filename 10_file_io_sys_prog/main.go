package main

import (
	"embed"
	"fmt"
	"log"
)

// 📂 10_file_io_sys_prog
// Embedding static assets into our Go program

//! Below is not a comment -> It's the embedding syntax.

//go:embed hello.txt

var data string

//go:embed public
var public embed.FS

// ----------------------------------------
// entire enterprise application in Go 🔥 |
// ----------------------------------------

func main() {
fmt.Println("data:",data)

folderData,err:=public.ReadFile("public/data.txt")
if err != nil {
		log.Fatal("ERROR:", err.Error())
	}
fmt.Println("folderData:\n",string(folderData))



// $ go run main.go
// data: Static Asset/File embedding in Golang
// folderData:
//  1 - Golang
// 2 - Javascript
// 3 - Typescript

}

//⭐ *** Example of Cleaner Architecture ***
//⭐ *** This is how professional Go web servers commonly structure embedded assets. ***
// package main

// import (
// 	"embed"
// 	"fmt"
// 	"io/fs"
// 	"log"
// )

// //go:embed public/*
// var public embed.FS

// func main() {

// 	publicFS, err := fs.Sub(public, "public")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	data, err := publicFS.ReadFile("data.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(string(data))
// }



