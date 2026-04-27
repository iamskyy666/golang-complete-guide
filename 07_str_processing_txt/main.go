package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// 📂 07_str_processing_txt
// Proj: Config-Parser

func parseConfig(content string)(map[string]string,error){
	config:=make(map[string]string)
	re:=regexp.MustCompile(`^\s*([\w.-]+)\s*=\s*(?:'([^']*)'|"([^"]*)"|([^#\s]*))?(?:\s*#.*)?$`)

	scanner:=bufio.NewScanner(strings.NewReader(content)) // scan from string
	lineNo:=0

	for scanner.Scan(){
		lineNo++
		line:=scanner.Text()

		trimmedLine:=strings.TrimSpace(line)

		if trimmedLine=="" || strings.HasPrefix(trimmedLine,"#"){
			continue
		}

		matches:=re.FindStringSubmatch(trimmedLine)
		if matches == nil{
			fmt.Printf("Line %d: '%s' - Is invalid!\n",lineNo,line)
			continue
		}

		key:=matches[1]

		var value string

		if matches[2] != ""{
		value=matches[2]
		}else if matches[3]!=""{
			value=matches[3]
		}else{
			value=matches[4]
		}

		config[key] = value

	}

	return config,nil
}


func main() {
	envfileContent:=`
# Server Configuration

PORT=5000
ENV=development

# Database Configuration

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword1234
DB_NAME=myapp_db
API_ENDPOINT="https://api.example.com/v1"

# Authentication / JWT

JWT_SECRET=supersecretkey123
JWT_EXPIRY=24h

# Feature Flags / Debug

LOG_LEVEL=debug
ENABLE_CACHE=true

# Empty values
EMPTY_KEY=
ANOTHER_KEY_NO_VALUE=`


config,err:=parseConfig(envfileContent)

if err != nil {
	fmt.Println("ERROR:",err)
	os.Exit(1)
}

for k,v:=range config{
		fmt.Printf("%s=%q\n",k,v)
}
	
}

// OUTPUT:
// $ go run main.go
// DB_NAME="myapp_db"
// JWT_EXPIRY="24h"
// LOG_LEVEL="debug"
// ENABLE_CACHE="true"
// ENV="development"
// DB_HOST="localhost"
// DB_PASSWORD="yourpassword1234"
// API_ENDPOINT="https://api.example.com/v1"
// JWT_SECRET="supersecretkey123"
// EMPTY_KEY=""
// ANOTHER_KEY_NO_VALUE=""
// PORT="5000"
// DB_PORT="5432"
// DB_USER="postgres"
