```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// 📂 10_file_io_sys_prog

// working with files

func main() {
	data:="Working with files in Golang!"
	// write
	err:=os.WriteFile("text.txt",[]byte(data),0644)
	if err != nil {
		log.Fatal("ERROR:",err.Error())
	}

	fmt.Println("Done writing.. ✅")

	// read
	filePath:="text.txt"
	content,err:=os.ReadFile(filePath)
	if err != nil {
		log.Fatal("ERROR:",err.Error())
	}

	fmt.Println("FILE-CONTENT:",string(content))

	data2,err:=os.Create("file-via-create.txt")
	if err != nil {
		log.Fatal("ERROR:",err.Error())
	}

	defer data2.Close()

	_,err = data2.WriteString("WriteString() is like Write, but writes the contents of string s rather than a slice of bytes. ☑️")
	if err != nil {
		log.Fatal("ERROR:",err.Error())
	}

	// read line-by-line

	newfile,err:= os.Open("line-by-line.txt")
	if err != nil {
		log.Fatal("ERROR:",err.Error())
	}

	defer newfile.Close()

	scanner:=bufio.NewScanner(newfile)

	lineNum:=0

	fmt.Println("*** READING/SCANNING line-by-line ***")

	for scanner.Scan(){
		lineNum++
		fmt.Println(lineNum,scanner.Text())
	}

	if err:=scanner.Err(); err != nil {
		if err!=io.EOF{
			log.Fatal("ERROR:",err.Error())
		}
	}

	// append content (instead of overriding)
	newfile2,err:=os.OpenFile("line-by-line.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil {
		log.Fatal("ERROR:",err.Error())
	}
	defer newfile2.Close()

	_, err = newfile2.WriteString("\n- Gengar 👻\n")
if err != nil {
	log.Fatal("ERROR:", err.Error())
}

_, err = newfile2.WriteString("- Dragonite 🐲\n")
if err != nil {
	log.Fatal("ERROR:", err.Error())
}

_, err = newfile2.WriteString("- Shiftry 🌿/🖤\n")
if err != nil {
	log.Fatal("ERROR:", err.Error())
}

_, err = newfile2.WriteString("- Lapras 💦\n")
if err != nil {
	log.Fatal("ERROR:", err.Error())
}

_, err = newfile2.WriteString("- Poliwhirl 💦\n")
if err != nil {
	log.Fatal("ERROR:", err.Error())
}

fmt.Println("Done appending.. ✅")

}

// OUTPUT:
// $ go run  main.go
// Done writing.. ✅
// FILE-CONTENT: Working with files in Golang!
// *** READING/SCANNING line-by-line ***
// 1 - Umbreon 🖤
// 2 - Noctowl 🔮
// 3 - Absol 🖤
// 4 - Articuno ❄️
// 5 - Gengar 👻
// 6 - Dragonite 🐲
// 7 - Shiftry 🌿/🖤
// 8 - Lapras 💦
// 9 - Poliwhirl 💦
// Done appending.. ✅
```
```go
package main

import (
	"fmt"
	"path/filepath"
)

// 📂 10_file_io_sys_prog

// handling file-paths

func main() {

	// "path/filepath" - platform independent
	path1:=filepath.Join("C:","Users","Documents","rand-folder")
	fmt.Println(path1) // C:Users\Documents\rand-folder

	path2:=filepath.Join("config","app.yaml")
	fmt.Println(path2) // config\app.yaml

	// base
	fmt.Println(filepath.Base(path1)) // rand-folder

	// extension
	fmt.Println(filepath.Ext(path2)) // .yaml


	maliciousDir:="./users/./dir../other"
	fmt.Println(filepath.Clean(maliciousDir)) // users\dir..\other

}
```

```go
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// 📂 10_file_io_sys_prog

// working with directories

func main() {
	err:=os.Mkdir("exampleDir",0755)
	if err != nil {
		log.Fatal("ERROR:",err.Error())
	}

	// err=os.MkdirAll("Downloads/static/images",0755)
	// Better way:
	dir:="Downloads/static/images"
	err=os.MkdirAll(filepath.Clean(dir),0755)
	if err != nil {
		log.Fatal("ERROR:",err.Error())
	}

	err=os.RemoveAll("toBDeleted")
	if err != nil {
		log.Fatal("ERROR:",err.Error())
	}

	fmt.Println("All ops. done ✅")
}

// All ops. done ✅
```

```go
package main

import (
	"fmt"
	"log"
	"os"
)

// 📂 10_file_io_sys_prog
// working with temp. files and directories
// mainly used for TESTING

func main() {

	tempFile, err := os.CreateTemp("", "logs.txt")
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}

	// cleanup
	defer func() {

		// CLOSE FIRST
		err := tempFile.Close()
		if err != nil {
			log.Fatal("ERROR:", err.Error())
		}

		// THEN REMOVE
		fmt.Println("✅ Removing temp. file..", tempFile.Name())

		err = os.Remove(tempFile.Name())
		if err != nil {
			log.Fatal("ERROR:", err.Error())
		}
	}()

	// write to temp file
	_, err = tempFile.Write(
		[]byte("Writing gibberish to temporary-file 📂\n"),
	)

	if err != nil {
		tempFile.Close()
		log.Fatal("ERROR:", err.Error())
		return
	}

	fmt.Println("✅ Temp file created:", tempFile.Name())

	// temp-dir
	tempDir,err:=os.MkdirTemp("","my_app_logs")
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}

	// cleanup
	defer func() {

		fmt.Println("tempDir:",tempDir)

		//REMOVE
		fmt.Println("✅ Removing temp. dir..", tempDir)

		err = os.RemoveAll(tempDir)
		if err != nil {
			log.Fatal("ERROR:", err.Error())
		}
	}()
}


// $ go run main.go
// ✅ Temp file created: C:\Users\ASUS\AppData\Local\Temp\logs.txt547468223
// tempDir: C:\Users\ASUS\AppData\Local\Temp\my_app_logs2924086784
// ✅ Removing temp. dir.. C:\Users\ASUS\AppData\Local\Temp\my_app_logs2924086784
// ✅ Removing temp. file.. C:\Users\ASUS\AppData\Local\Temp\logs.txt547468223
```