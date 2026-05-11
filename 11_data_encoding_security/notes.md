# Marshalling and Unmarshalling in Golang 💻

Marshalling and unmarshalling are core concepts in Go when we work with:

* APIs
* JSON
* files
* databases
* network communication
* configuration systems

If we build backend services in Go, we will use these concepts constantly.

---

# What is Marshalling?

**Marshalling** means:

> Converting a Go data structure into another format.

Usually:

* Go struct → JSON
* Go struct → XML
* Go struct → bytes

The most common case is:

```go
struct -> JSON string
```

---

# What is Unmarshalling?

**Unmarshalling** means:

> Converting external data into Go data structures.

Usually:

* JSON → struct
* XML → struct
* bytes → struct

The most common case is:

```go
JSON -> struct
```

---

# Real World Analogy

Imagine we have a Go struct:

```go
type User struct {
	Name string
	Age  int
}
```

Inside our Go application, this is a normal Go object in memory.

But:

* browsers
* mobile apps
* databases
* APIs

cannot understand Go structs directly.

So we must convert it into a portable format like JSON.

That conversion is:

* Marshalling

Then when another system sends JSON back:

* we convert JSON into Go structs
* that's Unmarshalling

---

# The `encoding/json` Package

Go provides:

```go
import "encoding/json"
```

This package handles:

* JSON encoding
* JSON decoding

Main functions:

| Function           | Purpose   |
| ------------------ | --------- |
| `json.Marshal()`   | Go → JSON |
| `json.Unmarshal()` | JSON → Go |

---

# Basic Marshalling Example

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{
		Name: "Skyy",
		Age:  29,
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}
```

Output:

```json
{"Name":"Skyy","Age":29}
```

---

# What Happened Internally?

This line:

```go
json.Marshal(user)
```

did several things internally:

1. Reflected over the struct
2. Found exported fields
3. Created JSON keys
4. Converted values
5. Returned bytes

Result type:

```go
[]byte
```

NOT string.

That is why we do:

```go
string(jsonData)
```

for printing.

---

# Why Does Marshal Return `[]byte`?

Because JSON is usually:

* transmitted over networks
* written to files
* sent in HTTP responses

All of those work with bytes.

So Go returns:

```go
[]byte
```

instead of string.

---

# Basic Unmarshalling Example

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	jsonData := []byte(`{
		"Name":"Skyy",
		"Age":29
	}`)

	var user User

	err := json.Unmarshal(jsonData, &user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
```

Output:

```go
{Skyy 29}
```

---

# VERY IMPORTANT:

# Why Do We Pass Pointer in Unmarshal?

This:

```go
&user
```

is critical.

Because:

```go
json.Unmarshal()
```

needs to MODIFY the variable.

Without pointer:

```go
json.Unmarshal(jsonData, user)
```

Go would pass a copy.

So the original struct would not change.

---

# Marshal vs Unmarshal

| Operation | Meaning   |
| --------- | --------- |
| Marshal   | Go → JSON |
| Unmarshal | JSON → Go |

---

# Exported Fields Matter

This is EXTREMELY important.

---

## Works

```go
type User struct {
	Name string
	Age  int
}
```

---

## Does NOT Work

```go
type User struct {
	name string
	age  int
}
```

Output:

```json
{}
```

Why?

Because:

* lowercase fields are unexported
* `encoding/json` uses reflection
* reflection cannot access unexported fields safely

So:

* ONLY exported fields participate in marshalling/unmarshalling.

---

# JSON Struct Tags

Usually we do NOT want:

```json
{
  "Name":"Skyy"
}
```

We want:

```json
{
  "name":"Skyy"
}
```

We use struct tags.

---

# Struct Tags

```go
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

Now output becomes:

```json
{
  "name":"Skyy",
  "age":29
}
```

---

# Anatomy of a Struct Tag

```go
`json:"name"`
```

Means:

| Part     | Meaning          |
| -------- | ---------------- |
| `json`   | package/tag name |
| `"name"` | JSON key         |

---

# Full Example

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user := User{
		Name: "Skyy",
		Age:  29,
	}

	data, _ := json.Marshal(user)

	fmt.Println(string(data))
}
```

Output:

```json
{"name":"Skyy","age":29}
```

---

# Ignoring Fields

We can exclude fields from JSON.

```go
type User struct {
	Name     string `json:"name"`
	Password string `json:"-"`
}
```

Now password is ignored.

Very common in APIs.

---

# `omitempty`

This is heavily used in backend development.

---

## Example

```go
type User struct {
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}
```

If email is empty:

```go
Email == ""
```

then it is omitted entirely.

---

## Without omitempty

```json
{
  "name":"Skyy",
  "email":""
}
```

---

## With omitempty

```json
{
  "name":"Skyy"
}
```

---

# What Counts as Empty?

| Type    | Empty Value |
| ------- | ----------- |
| string  | `""`        |
| int     | `0`         |
| bool    | `false`     |
| slice   | `nil`       |
| map     | `nil`       |
| pointer | `nil`       |

---

# Nested Struct Marshalling

```go
type Address struct {
	City string `json:"city"`
}

type User struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
}
```

Output:

```json
{
  "name":"Skyy",
  "address":{
    "city":"Kolkata"
  }
}
```

---

# Slices and Arrays

```go
type User struct {
	Name    string   `json:"name"`
	Hobbies []string `json:"hobbies"`
}
```

Output:

```json
{
  "name":"Skyy",
  "hobbies":["coding","gaming"]
}
```

---

# Maps

```go
m := map[string]int{
	"apple": 5,
	"mango": 10,
}
```

Can be marshalled too.

Output:

```json
{
  "apple":5,
  "mango":10
}
```

---

# Pretty JSON with `MarshalIndent`

Normal JSON:

```json
{"name":"Skyy","age":29}
```

Pretty JSON:

```go
data, _ := json.MarshalIndent(user, "", "  ")
```

Output:

```json
{
  "name": "Skyy",
  "age": 29
}
```

---

# Parameters of MarshalIndent

```go
json.MarshalIndent(v, prefix, indent)
```

| Parameter | Meaning                |
| --------- | ---------------------- |
| `v`       | data                   |
| `prefix`  | added before each line |
| `indent`  | indentation spaces     |

---

# Unmarshal Into Maps

Sometimes JSON structure is unknown.

Then we use:

```go
map[string]interface{}
```

Example:

```go
var result map[string]interface{}

json.Unmarshal(data, &result)
```

---

# Why `interface{}`?

Because JSON values can be:

* string
* bool
* array
* object
* number
* null

Go needs a flexible container.

---

# Type Assertions Become Necessary

```go
name := result["name"].(string)
```

Because:

* interface{} stores unknown types
* we must extract actual type

---

# Important JSON Number Behavior

JSON numbers become:

```go
float64
```

by default when unmarshalling into `interface{}`.

Example:

```go
age := result["age"].(float64)
```

NOT int.

This surprises many beginners.

---

# Working with Raw JSON

Sometimes we store raw JSON temporarily.

Go provides:

```go
json.RawMessage
```

Example:

```go
type Event struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}
```

Very useful in:

* event systems
* webhooks
* dynamic APIs

---

# Streaming JSON with Decoder

Instead of:

```go
json.Unmarshal()
```

we can use:

```go
json.NewDecoder()
```

Example:

```go
file, _ := os.Open("data.json")

decoder := json.NewDecoder(file)

var user User

decoder.Decode(&user)
```

Useful for:

* files
* HTTP requests
* large JSON streams

---

# Encoding JSON to Writer

Instead of:

```go
json.Marshal()
```

we can use:

```go
json.NewEncoder()
```

Example:

```go
json.NewEncoder(os.Stdout).Encode(user)
```

Useful for:

* HTTP responses
* files
* streaming

---

# Real HTTP API Example

---

## Sending JSON Response

```go
func handler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name: "Skyy",
		Age:  29,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}
```

---

# Receiving JSON Request

```go
func handler(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Println(user)
}
```

This is the backbone of REST APIs in Go.

---

# Common Beginner Mistakes

---

## 1. Using lowercase fields

```go
type User struct {
	name string
}
```

Will not work.

---

## 2. Forgetting pointer in Unmarshal

Wrong:

```go
json.Unmarshal(data, user)
```

Correct:

```go
json.Unmarshal(data, &user)
```

---

## 3. Forgetting error handling

Wrong:

```go
data, _ := json.Marshal(user)
```

Always handle errors properly.

---

## 4. Invalid JSON

JSON keys MUST use double quotes.

Wrong:

```json
{
  'name':'Skyy'
}
```

Correct:

```json
{
  "name":"Skyy"
}
```

---

# Custom Marshalling

We can define custom behavior.

Go uses interfaces:

```go
json.Marshaler
json.Unmarshaler
```

---

# Custom Marshal Example

```go
type User struct {
	Name string
}

func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"username": u.Name,
	})
}
```

Now JSON becomes:

```json
{
  "username":"Skyy"
}
```

---

# Custom Unmarshal Example

```go
func (u *User) UnmarshalJSON(data []byte) error {
	var temp struct {
		Username string `json:"username"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	u.Name = temp.Username

	return nil
}
```

---

# Internal Mechanism: Reflection

The `encoding/json` package heavily uses:

Reflection in Computer Science

Reflection allows Go to:

* inspect struct fields
* inspect tags
* determine types dynamically
* assign values at runtime

Without reflection:

* generic JSON handling would be impossible.

---

# Performance Considerations

Go's standard JSON package is:

* reliable
* readable
* easy to use

But not the fastest.

Large-scale systems sometimes use:

* `jsoniter`
* `easyjson`
* code generation approaches

for higher performance.

---

# Mental Model

Think of it like this:

```text
GO STRUCT
    ↓ Marshal
JSON/BYTES
    ↓ Unmarshal
GO STRUCT
```

---

# Most Important Things to Remember

1. Marshal = Go → JSON
2. Unmarshal = JSON → Go
3. Marshal returns `[]byte`
4. Unmarshal requires POINTERS
5. Only EXPORTED fields work
6. Struct tags control JSON keys
7. `omitempty` removes empty fields
8. `Encoder/Decoder` are great for streams/files/HTTP
9. Reflection powers the entire system
10. This is fundamental for APIs in Go

# Encode and Decode in Golang

In Go, `Encode()` and `Decode()` are closely related to:

* `Marshal()`
* `Unmarshal()`

But they are designed for:

* streams
* files
* HTTP requests/responses
* large data handling

instead of directly working with raw byte slices.

---

# Big Picture

| Function                     | Works With      |
| ---------------------------- | --------------- |
| `json.Marshal()`             | bytes           |
| `json.Unmarshal()`           | bytes           |
| `json.NewEncoder().Encode()` | writers/streams |
| `json.NewDecoder().Decode()` | readers/streams |

---

# Mental Model

```text id="h10ukg"
Marshal   -> convert to []byte
Unmarshal -> convert from []byte

Encode    -> write JSON directly somewhere
Decode    -> read JSON directly from somewhere
```

---

# Core Difference

---

## Marshal / Unmarshal

We manually deal with bytes.

```go id="8mab4q"
data, _ := json.Marshal(user)

json.Unmarshal(data, &user)
```

---

## Encode / Decode

Go handles streaming automatically.

```go id="s7z3k5"
json.NewEncoder(w).Encode(user)

json.NewDecoder(r.Body).Decode(&user)
```

Much cleaner for:

* APIs
* files
* sockets
* streaming systems

---

# What is Encoding?

Encoding means:

> Convert Go data into JSON and directly WRITE it somewhere.

Usually:

* HTTP response
* file
* stdout
* network connection

---

# What is Decoding?

Decoding means:

> READ JSON from somewhere and convert it into Go data.

Usually:

* HTTP request body
* file
* socket
* stream

---

# The Encoder

Created using:

```go id="56h8hc"
json.NewEncoder(writer)
```

It requires an:

```go id="i6k9dn"
io.Writer
```

---

# The Decoder

Created using:

```go id="5svt2w"
json.NewDecoder(reader)
```

It requires an:

```go id="rbdjlwm"
io.Reader
```

---

# VERY IMPORTANT:

# io.Reader and io.Writer

These are foundational Go interfaces.

---

## io.Reader

Something we can READ data from.

Examples:

* files
* HTTP request body
* strings
* network sockets

---

## io.Writer

Something we can WRITE data to.

Examples:

* files
* HTTP response
* terminal
* buffers

---

# Encode Example

```go id="0n2v11"
package main

import (
	"encoding/json"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user := User{
		Name: "Skyy",
		Age:  29,
	}

	json.NewEncoder(os.Stdout).Encode(user)
}
```

Output:

```json id="1r8o9v"
{"name":"Skyy","age":29}
```

---

# What Happened Internally?

This:

```go id="c6j6m6"
json.NewEncoder(os.Stdout)
```

creates an encoder that writes directly to terminal output.

Then:

```go id="zbqpj3"
Encode(user)
```

1. Converts struct → JSON
2. Writes JSON directly to stdout

No manual byte handling needed.

---

# Equivalent Marshal Version

This:

```go id="i98zqc"
json.NewEncoder(os.Stdout).Encode(user)
```

is roughly equivalent to:

```go id="fc7mzg"
data, _ := json.Marshal(user)

os.Stdout.Write(data)
```

But encoder is:

* cleaner
* more memory efficient
* stream friendly

---

# Decode Example

```go id="s7jj5p"
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonData := `{
		"name":"Skyy",
		"age":29
	}`

	reader := strings.NewReader(jsonData)

	var user User

	err := json.NewDecoder(reader).Decode(&user)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", user)
}
```

Output:

```go id="k0vk3h"
{Name:Skyy Age:29}
```

---

# What is `strings.NewReader()`?

This:

```go id="mghkrc"
strings.NewReader(jsonData)
```

creates an:

Input Stream

that implements:

```go id="w4kz4j"
io.Reader
```

So the decoder can read from it.

---

# Why Use Decoder Instead of Unmarshal?

---

# 1. Better for Files

---

## Using Unmarshal

```go id="3fy3r6"
data, _ := os.ReadFile("user.json")

json.Unmarshal(data, &user)
```

Entire file loads into memory.

---

## Using Decoder

```go id="2as8ix"
file, _ := os.Open("user.json")

json.NewDecoder(file).Decode(&user)
```

Streams directly from file.

Better for large files.

---

# 2. Better for HTTP APIs

This is HUGE in backend development.

---

# Reading Request Body

```go id="1apm1u"
func handler(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
```

Why?

Because:

```go id="7h6q5m"
r.Body
```

is already an:

```go id="u8tjlwm"
io.Reader
```

Perfect for decoder.

---

# Sending Response Body

```go id="1mhh9w"
func handler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name: "Skyy",
		Age:  29,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}
```

Because:

```go id="r5xjsh"
w
```

implements:

```go id="jlwmz2"
io.Writer
```

Perfect for encoder.

---

# This is the Standard Pattern in Go APIs

---

## Incoming JSON

```go id="jlwm8h"
json.NewDecoder(r.Body).Decode(&data)
```

---

## Outgoing JSON

```go id="jlwm0j"
json.NewEncoder(w).Encode(response)
```

We will see this everywhere in Go backend development.

---

# Encode Automatically Adds Newline

This surprises many beginners.

---

# Example

```go id="jlwmga"
json.NewEncoder(os.Stdout).Encode(user)
```

Output:

```json id="jlwmj0"
{"name":"Skyy","age":29}
```

with a trailing newline.

---

# Marshal Does NOT Add Newline

```go id="jlwmrv"
json.Marshal(user)
```

returns only raw JSON bytes.

---

# Why Encoder/Decoder Are Better for Streams

Imagine a 5GB JSON file.

---

## Unmarshal Approach

```go id="jlwm6r"
data, _ := os.ReadFile("big.json")
```

BAD:

* entire file enters memory

---

## Decoder Approach

```go id="jlwmzh"
json.NewDecoder(file).Decode(&data)
```

GOOD:

* reads progressively

Much more scalable.

---

# Multiple JSON Objects in Stream

Decoder can process continuous JSON streams.

Example:

```json id="jlwm44"
{"name":"A"}
{"name":"B"}
{"name":"C"}
```

We can repeatedly decode:

```go id="jlwm1g"
decoder := json.NewDecoder(file)

for {
	var user User

	err := decoder.Decode(&user)

	if err == io.EOF {
		break
	}

	fmt.Println(user)
}
```

This is powerful for:

* logs
* streaming APIs
* sockets
* message queues

---

# Decoder Tokenization

Decoder can even parse token-by-token.

Example:

```go id="jlwm1d"
decoder.Token()
```

Used in:

* advanced parsers
* streaming parsers
* huge JSON processing

Less common for beginners.

---

# Setting Indentation with Encoder

```go id="jlwmok"
encoder := json.NewEncoder(os.Stdout)

encoder.SetIndent("", "  ")

encoder.Encode(user)
```

Output:

```json id="jlwm4j"
{
  "name": "Skyy",
  "age": 29
}
```

---

# Preventing HTML Escaping

By default:

```go id="jlwm1r"
< > &
```

get escaped.

Example:

```json id="jlwm6m"
"\u003cdiv\u003e"
```

Disable using:

```go id="jlwmfm"
encoder.SetEscapeHTML(false)
```

---

# Common Beginner Mistakes

---

# 1. Forgetting Pointer in Decode

Wrong:

```go id="jlwmlo"
Decode(user)
```

Correct:

```go id="jlwmvt"
Decode(&user)
```

Because decoder must modify the variable.

---

# 2. Ignoring Errors

Wrong:

```go id="jlwm6c"
json.NewDecoder(r.Body).Decode(&user)
```

Always check errors.

---

# 3. Using Marshal for Huge Files

This can waste memory badly.

Decoder is better for streams.

---

# 4. Forgetting to Close Files

```go id="jlwm2o"
file, _ := os.Open("data.json")
defer file.Close()
```

Important.

---

# Internal Flow of Encoder

```text id="jlwm3n"
Go Struct
    ↓
Encoder
    ↓
JSON
    ↓
io.Writer
```

---

# Internal Flow of Decoder

```text id="jlwm6v"
io.Reader
    ↓
JSON
    ↓
Decoder
    ↓
Go Struct
```

---

# Marshal vs Encode

| Feature         | Marshal | Encode |
| --------------- | ------- | ------ |
| Returns bytes   | ✅       | ❌      |
| Writes directly | ❌       | ✅      |
| Stream friendly | ❌       | ✅      |
| Great for APIs  | ⚠️      | ✅      |
| Great for files | ⚠️      | ✅      |

---

# Unmarshal vs Decode

| Feature          | Unmarshal | Decode |
| ---------------- | --------- | ------ |
| Uses []byte      | ✅         | ❌      |
| Uses streams     | ❌         | ✅      |
| Good for files   | ⚠️        | ✅      |
| Good for HTTP    | ⚠️        | ✅      |
| Memory efficient | ❌         | ✅      |

---

# Real Production Mentality

---

## Small JSON already in memory?

Use:

```go id="jlwmdd"
Marshal / Unmarshal
```

---

## Working with:

* HTTP
* files
* sockets
* streams
* large JSON

Use:

```go id="jlwmkl"
Encoder / Decoder
```

---

# The Most Important Thing to Understand

`Marshal/Unmarshal` work with:

```go id="jlwmqy"
[]byte
```

while

`Encode/Decode` work with:

```go id="jlwm4o"
io.Writer / io.Reader
```

That is the fundamental distinction.

# Base64 Encoding & Decoding in Golang

Base64 is one of the most important encoding techniques used in:

* APIs
* authentication
* JWTs
* file transfer
* email systems
* image embedding
* cryptography workflows
* binary transmission

And it is VERY commonly misunderstood.

---

# First:

# Base64 is NOT Encryption

This is the biggest beginner misconception.

Base64:

* does NOT secure data
* does NOT hide secrets
* does NOT require a key

It is only:

> A binary-to-text encoding scheme.

Meaning:

* convert binary data into safe text format

Anyone can decode Base64 instantly.

---

# Why Base64 Exists

Computers handle:

* binary data
* raw bytes

But many systems historically handled only:

* plain text
* ASCII-safe characters

Examples:

* old email systems
* URLs
* JSON
* HTTP headers

Binary data could break these systems.

So Base64 converts:

```text id="5vhxq4"
binary data -> safe text
```

using only safe characters.

---

# What Characters Does Base64 Use?

Standard Base64 uses:

```text id="1cmy5x"
A-Z
a-z
0-9
+
/
```

Total:

```text id="a4wlw8"
64 characters
```

Hence the name:

```text id="v7ttjlwm"
Base64
```

---

# Why "Base64"?

Because each Base64 digit represents:

Binary Number System

6 bits of data.

Why?

```text id="jlwmn8"
2^6 = 64
```

So:

* every Base64 character encodes 6 bits

---

# The Core Problem Base64 Solves

Suppose we have raw binary:

```text id="jlwmv0"
010101110101010101
```

This is unsafe for:

* URLs
* JSON
* emails
* text protocols

Base64 converts it into safe printable text:

```text id="jlwm50"
VGVzdA==
```

Now it can travel safely through text systems.

---

# Important:

# Base64 Increases Size

Base64 makes data about:

```text id="jlwm8p"
33% larger
```

Why?

Because:

* 3 bytes (24 bits)
  become:
* 4 Base64 characters

---

# Golang Package

Go provides:

```go id="jlwmhe"
import "encoding/base64"
```

---

# Basic Encoding Example

```go id="jlwm55"
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "Hello World"

	encoded := base64.StdEncoding.EncodeToString([]byte(data))

	fmt.Println(encoded)
}
```

Output:

```text id="jlwmzt"
SGVsbG8gV29ybGQ=
```

---

# What Happened Here?

---

## Step 1

```go id="jlwm1c"
[]byte(data)
```

Convert string → bytes.

Because Base64 works on raw bytes.

---

## Step 2

```go id="jlwm55"
base64.StdEncoding.EncodeToString()
```

Converts bytes → Base64 string.

---

# Decoding Example

```go id="jlwm94"
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	encoded := "SGVsbG8gV29ybGQ="

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(decoded))
}
```

Output:

```text id="jlwm4g"
Hello World
```

---

# What DecodeString Returns

```go id="jlwmfs"
[]byte
```

Because original data may not be text.

Could be:

* image
* PDF
* ZIP
* executable
* audio

So Go returns raw bytes.

---

# Why We Convert Back to String

```go id="jlwm1m"
string(decoded)
```

because our original data was text.

---

# Internal Mechanism

Base64 processes:

* 3 bytes at a time
* 24 bits total

Then splits into:

* 4 groups of 6 bits

Each 6-bit value maps to:

* one Base64 character

---

# Visual Example

Suppose:

```text id="jlwm3y"
Cat
```

ASCII bytes:

```text id="jlwmvx"
C = 01000011
a = 01100001
t = 01110100
```

Combined:

```text id="jlwm8s"
010000110110000101110100
```

Split into 6-bit chunks:

```text id="jlwmz0"
010000
110110
000101
110100
```

Mapped to Base64 table:

```text id="jlwmq4"
Q2F0
```

---

# Padding in Base64

Notice outputs often end with:

```text id="jlwm0z"
=
```

or:

```text id="jlwm4p"
==
```

These are padding characters.

---

# Why Padding Exists

Base64 works in:

* groups of 3 bytes

If data length isn't divisible by 3:

* padding is added

---

# Example

```text id="jlwmc7"
Hi
```

Only 2 bytes.

Base64 adds padding:

```text id="jlwm3m"
SGk=
```

---

# Standard Encoding

Most common:

```go id="jlwm8m"
base64.StdEncoding
```

Uses:

* `+`
* `/`

---

# URL Encoding Variant

URLs treat:

* `+`
* `/`

specially.

So Go provides:

```go id="jlwm3p"
base64.URLEncoding
```

Uses:

* `-`
* `_`

instead.

---

# Example

```go id="jlwmgi"
encoded := base64.URLEncoding.EncodeToString([]byte("hello"))
```

Very common in:

* JWT tokens
* URLs
* OAuth systems

---

# JWT Uses Base64URL

JSON Web Tokens heavily use:

JSON Web Token

specifically:

* Base64 URL encoding

A JWT looks like:

```text id="jlwmgm"
xxxxx.yyyyy.zzzzz
```

Each section is Base64URL encoded.

---

# Encoding Binary Files

Base64 is heavily used for transmitting files.

---

# Example: Image → Base64

```go id="jlwmqf"
package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("image.png")
	if err != nil {
		panic(err)
	}

	encoded := base64.StdEncoding.EncodeToString(data)

	fmt.Println(encoded)
}
```

Now image becomes safe text.

---

# Decoding Back Into File

```go id="jlwm91"
decoded, _ := base64.StdEncoding.DecodeString(encoded)

os.WriteFile("new.png", decoded, 0644)
```

---

# Very Common Real-World Uses

---

# 1. Basic Authentication

HTTP Basic Auth:

```text id="jlwmx4"
username:password
```

gets Base64 encoded.

Example:

```text id="jlwm6j"
Authorization: Basic c2t5eTpzZWNyZXQ=
```

---

# 2. JWT Tokens

JWT payloads are Base64URL encoded.

---

# 3. Embedding Images in HTML/CSS

Example:

```html id="jlwm9l"
<img src="data:image/png;base64,iVBORw0KGgoAAA..." />
```

---

# 4. Sending Binary Through JSON

JSON supports text safely.

Not raw binary.

So:

* images
* PDFs
* files

are often Base64 encoded first.

---

# Streaming Encoder

Go also supports stream encoding.

---

# Encoder Example

```go id="jlwmmb"
package main

import (
	"encoding/base64"
	"os"
)

func main() {
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)

	encoder.Write([]byte("Hello"))

	encoder.Close()
}
```

---

# VERY IMPORTANT:

# Must Close Encoder

```go id="jlwm7g"
encoder.Close()
```

is critical.

Because encoder buffers internally.

Without closing:

* final bytes may not flush

---

# Streaming Decoder

```go id="jlwmvv"
decoder := base64.NewDecoder(base64.StdEncoding, reader)
```

Useful for:

* large files
* streams
* network data

---

# Raw Base64 (Without Padding)

Sometimes systems omit:

```text id="jlwm0v"
=
```

padding.

Go provides:

```go id="jlwmnq"
base64.RawStdEncoding
```

and:

```go id="jlwm0m"
base64.RawURLEncoding
```

Common in:

* JWTs
* URLs
* compact tokens

---

# Example

```go id="jlwm8n"
base64.RawStdEncoding.EncodeToString([]byte("Hi"))
```

Output:

```text id="jlwmij"
SGk
```

No `=` padding.

---

# Common Beginner Mistakes

---

# 1. Thinking Base64 is Encryption

WRONG.

Anyone can decode it instantly.

---

# 2. Forgetting Error Handling

Decode can fail.

Example:

```go id="jlwm9x"
base64.StdEncoding.DecodeString("%%%")
```

Invalid Base64.

---

# 3. Forgetting Conversion to []byte

Encoding functions need bytes.

Wrong:

```go id="jlwm0w"
EncodeToString(data)
```

Correct:

```go id="jlwmj5"
EncodeToString([]byte(data))
```

---

# 4. Forgetting Encoder Close

Important for stream encoders.

---

# Performance Notes

Base64:

* increases size
* adds CPU overhead

So:

* avoid unnecessary encoding
* especially for huge files

---

# Mental Model

```text id="jlwmcv"
Original Data
      ↓
Bytes
      ↓
Base64 Encoding
      ↓
Safe Text
      ↓
Transport/Storage
      ↓
Base64 Decoding
      ↓
Original Bytes
```

---

# Most Important Things to Remember

1. Base64 is encoding, NOT encryption
2. Converts binary → safe text
3. Uses 64 ASCII-safe characters
4. Adds ~33% size overhead
5. Standard package:

   ```go
   encoding/base64
   ```
6. Main functions:

   ```go
   EncodeToString()
   DecodeString()
   ```
7. URL-safe variant exists
8. JWT heavily uses Base64URL
9. Base64 works on bytes
10. Common for APIs, auth, file transfer, and web systems

