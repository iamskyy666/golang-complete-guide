```go
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
```

We’ve basically built a **mini `.env` parser** from scratch. Let’s walk through it like we would in a real code review so we *actually understand what’s happening under the hood*.

---

# 🔥 What this project does

At a high level:

> It takes `.env`-style text → parses it → converts it into a `map[string]string`

So this:

```
PORT=5000
ENV=development
```

Becomes:

```go
map[string]string{
  "PORT": "5000",
  "ENV":  "development",
}
```

---

# 🧠 Core Function: `parseConfig`

```go
func parseConfig(content string)(map[string]string,error)
```

### Input:

* raw string (your `.env` content)

### Output:

* `map[string]string` → parsed key-value pairs
* `error` → (currently unused, but good design choice)

---

# 🧱 Step-by-step breakdown

---

## 1. Create result map

```go
config := make(map[string]string)
```

👉 This will store:

```
KEY → VALUE
```

---

## 2. Regex (the heart of the parser)

```go
re := regexp.MustCompile(`^\s*([\w.-]+)\s*=\s*(?:'([^']*)'|"([^"]*)"|([^#\s]*))?(?:\s*#.*)?$`)
```

Let’s break this down properly.

---

### 🔍 Pattern explained

```
^\s*
```

* allow leading spaces

---

```
([\w.-]+)
```

👉 **Group 1 = KEY**

* letters, numbers, `_`, `.`, `-`

---

```
\s*=\s*
```

* equal sign with optional spaces

---

```
(?: ... )?
```

👉 optional value group

Inside it:

### Case 1: single quotes

```
'([^']*)'
```

👉 Group 2

---

### Case 2: double quotes

```
"([^"]*)"
```

👉 Group 3

---

### Case 3: unquoted value

```
([^#\s]*)
```

👉 Group 4

---

```
(?:\s*#.*)?
```

👉 optional comment at end

---

### ✅ So we support:

| Format                | Works? |
| --------------------- | ------ |
| `KEY=value`           | ✅      |
| `KEY="value"`         | ✅      |
| `KEY='value'`         | ✅      |
| `KEY=value # comment` | ✅      |
| `KEY=`                | ✅      |

---

# 🔁 3. Reading line by line

```go
scanner := bufio.NewScanner(strings.NewReader(content))
```

👉 Instead of reading file, we simulate it using a string

---

## Loop:

```go
for scanner.Scan()
```

Each iteration = one line

---

# ✂️ 4. Cleaning lines

```go
trimmedLine := strings.TrimSpace(line)
```

---

## Skip:

```go
if trimmedLine == "" || strings.HasPrefix(trimmedLine, "#") {
	continue
}
```

👉 ignores:

* empty lines
* comments

---

# 🔍 5. Apply regex

```go
matches := re.FindStringSubmatch(trimmedLine)
```

---

## If invalid:

```go
if matches == nil {
	fmt.Printf("Line %d: '%s' - Is invalid!\n", lineNo, line)
	continue
}
```

👉 graceful handling instead of crashing

---

# 🧠 6. Extract key

```go
key := matches[1]
```

👉 Group 1 = key

---

# 🧠 7. Extract value (smart logic)

```go
if matches[2] != "" {
	value = matches[2]
} else if matches[3] != "" {
	value = matches[3]
} else {
	value = matches[4]
}
```

👉 Priority:

1. single-quoted
2. double-quoted
3. unquoted

---

## Example:

```
API_ENDPOINT="https://api.example.com/v1"
```

👉 matches[3] used

---

# ⚠️ Special case: empty values

```
EMPTY_KEY=
```

👉 all groups empty → value = ""

Handled correctly ✔️

---

# 🧾 8. Store in map

```go
config[key] = value
```

---

# 🖨️ 9. Output

```go
fmt.Printf("%s=%q\n", k, v)
```

👉 `%q` prints quoted values:

```
PORT="5000"
```

---

# 🔥 What makes this project strong

### ✅ 1. Real-world parsing

You handled:

* comments
* spaces
* quoted values
* empty values

---

### ✅ 2. Regex design is solid

Covers most `.env` use cases

---

### ✅ 3. Safe iteration

Using `bufio.Scanner` → efficient

---

### ✅ 4. Graceful error handling

Invalid lines don’t crash program

---

# ⚠️ Limitations (important if we want to improve)

Let’s be honest—this is good, but not production-ready yet.

---

## ❌ 1. No multiline values

```
KEY="line1
line2"
```

👉 not supported

---

## ❌ 2. No escape handling

```
KEY="hello \"world\""
```

👉 breaks

---

## ❌ 3. No environment expansion

```
HOME=/home/user
PATH=$HOME/bin
```

👉 `$HOME` not resolved

---

## ❌ 4. Scanner limit (~64K per line)

Large lines may fail

---

# 🚀 How we’d improve it

If we wanted to level this up:

### 1. Add error return instead of print

### 2. Support escaped quotes

### 3. Add variable expansion

### 4. Switch to file input instead of string

### 5. Add validation layer

---

# 🧠 Mental Model

This project is basically:

> **Lexer (regex) + line parser + map builder**

---
