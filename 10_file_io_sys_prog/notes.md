Working with files is one of the most important parts of backend and systems programming in Go. Go gives us a very clean and powerful standard library for file handling through packages like:

* `os`
* `io`
* `bufio`
* `io/ioutil` (older, mostly deprecated)
* `path/filepath`

The design philosophy is:

* Everything is treated as streams of bytes
* Files implement interfaces like `io.Reader` and `io.Writer`
* Small composable interfaces are preferred

---

# 1. The `os` Package

The main package for file operations is:

```go
import "os"
```

The core type is:

```go
type File struct
```

A file in Go is represented by:

```go
*os.File
```

It represents:

* regular files
* directories
* pipes
* terminals
* sockets

---

# 2. Opening Files

---

# `os.Open()`

Opens a file in **read-only mode**.

```go
file, err := os.Open("notes.txt")
if err != nil {
    panic(err)
}

defer file.Close()
```

---

## What does it return?

```go
(*os.File, error)
```

---

# Why `defer file.Close()`?

VERY important.

Files consume OS resources.

If we don't close them:

* memory leaks happen
* file descriptor leaks happen
* program may crash eventually

So we usually do:

```go
defer file.Close()
```

immediately after successful open.

---

# 3. Creating Files

---

# `os.Create()`

Creates a new file.

```go
file, err := os.Create("hello.txt")
if err != nil {
    panic(err)
}

defer file.Close()
```

---

## Important behavior

If file exists:

* it gets TRUNCATED (emptied)

Equivalent behavior:

```text
create if not exists
overwrite if exists
```

---

# 4. Writing to Files

There are MANY ways.

---

# Method 1 — `WriteString`

```go
file.WriteString("Hello Go\n")
```

Example:

```go
package main

import (
    "os"
)

func main() {
    file, err := os.Create("test.txt")
    if err != nil {
        panic(err)
    }

    defer file.Close()

    file.WriteString("Hello World\n")
}
```

---

# Method 2 — `Write`

`Write()` accepts bytes.

```go
data := []byte("Hello")
file.Write(data)
```

---

# Why bytes?

Files fundamentally store bytes.

Strings are converted into bytes.

---

# 5. Reading Files

---

# Method 1 — `os.ReadFile()`

Simplest method.

```go
data, err := os.ReadFile("test.txt")
if err != nil {
    panic(err)
}

fmt.Println(string(data))
```

---

## Return Type

```go
([]byte, error)
```

Again:

files are bytes.

---

# Full Example

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    data, err := os.ReadFile("test.txt")
    if err != nil {
        panic(err)
    }

    fmt.Println(string(data))
}
```

---

# 6. Appending to Files

By default:

* `os.Create()` overwrites

To append:

```go
os.OpenFile()
```

---

# `os.OpenFile()`

This is the advanced version.

```go
file, err := os.OpenFile(
    "test.txt",
    os.O_APPEND|os.O_WRONLY,
    os.ModeAppend,
)
```

---

# Important Flags

| Flag          | Meaning           |
| ------------- | ----------------- |
| `os.O_RDONLY` | read only         |
| `os.O_WRONLY` | write only        |
| `os.O_RDWR`   | read + write      |
| `os.O_APPEND` | append            |
| `os.O_CREATE` | create if missing |
| `os.O_TRUNC`  | truncate          |
| `os.O_EXCL`   | exclusive create  |

---

# Append Example

```go
package main

import (
    "os"
)

func main() {

    file, err := os.OpenFile(
        "test.txt",
        os.O_APPEND|os.O_WRONLY,
        0644,
    )

    if err != nil {
        panic(err)
    }

    defer file.Close()

    file.WriteString("New Line\n")
}
```

---

# What is `0644`?

Linux/Unix permissions.

---

# File Permission System

```text
Owner Group Others
```

---

# `0644`

```text
6 = read + write
4 = read
4 = read
```

Binary:

```text
110 100 100
```

---

# Common Permissions

| Permission | Meaning     |
| ---------- | ----------- |
| `0644`     | normal file |
| `0755`     | executable  |
| `0700`     | private     |

---

# 7. Reading Large Files Efficiently

`os.ReadFile()` loads entire file into memory.

BAD for huge files.

Use buffered reading.

---

# `bufio.NewScanner()`

Reads line by line.

```go
scanner := bufio.NewScanner(file)
```

---

# Example

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {

    file, err := os.Open("big.txt")
    if err != nil {
        panic(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
```

---

# Why use Scanner?

Efficient memory usage.

Instead of:

```text
load ENTIRE file
```

we do:

```text
read small chunks incrementally
```

---

# 8. Buffered Writing

---

# `bufio.NewWriter()`

Writes efficiently.

Instead of hitting disk every write:

* data is accumulated in memory buffer
* written in batches

---

# Example

```go
package main

import (
    "bufio"
    "os"
)

func main() {

    file, _ := os.Create("fast.txt")
    defer file.Close()

    writer := bufio.NewWriter(file)

    writer.WriteString("Hello\n")
    writer.WriteString("World\n")

    writer.Flush()
}
```

---

# IMPORTANT: `Flush()`

Without flush:

buffer may never be written.

Always flush buffered writers.

---

# 9. File Metadata

---

# `os.Stat()`

Gets information about file.

```go
info, err := os.Stat("test.txt")
```

---

# Useful Methods

| Method           | Meaning       |
| ---------------- | ------------- |
| `info.Name()`    | filename      |
| `info.Size()`    | bytes         |
| `info.Mode()`    | permissions   |
| `info.ModTime()` | modified time |
| `info.IsDir()`   | is directory  |

---

# Example

```go
info, _ := os.Stat("test.txt")

fmt.Println(info.Name())
fmt.Println(info.Size())
fmt.Println(info.IsDir())
```

---

# 10. Checking File Existence

Common pattern:

```go
_, err := os.Stat("test.txt")

if os.IsNotExist(err) {
    fmt.Println("File does not exist")
}
```

---

# 11. Renaming Files

```go
err := os.Rename("old.txt", "new.txt")
```

---

# 12. Deleting Files

```go
err := os.Remove("test.txt")
```

---

# Remove Directory

```go
os.RemoveAll("temp")
```

VERY dangerous:

Deletes recursively.

---

# 13. Working with Directories

---

# Create Directory

```go
os.Mkdir("data", 0755)
```

---

# Create Nested Directories

```go
os.MkdirAll("a/b/c", 0755)
```

---

# Reading Directory Contents

```go
entries, err := os.ReadDir(".")
```

---

# Example

```go
for _, entry := range entries {
    fmt.Println(entry.Name())
}
```

---

# 14. Seeking in Files

Files have an internal cursor.

---

# `Seek()`

```go
file.Seek(offset, whence)
```

---

# Whence Values

| Value | Meaning        |
| ----- | -------------- |
| `0`   | from beginning |
| `1`   | from current   |
| `2`   | from end       |

---

# Example

```go
file.Seek(0, 0)
```

Move cursor to beginning.

---

# 15. Copying Files

---

# `io.Copy()`

```go
io.Copy(dstFile, srcFile)
```

---

# Full Example

```go
package main

import (
    "io"
    "os"
)

func main() {

    src, _ := os.Open("source.txt")
    defer src.Close()

    dst, _ := os.Create("dest.txt")
    defer dst.Close()

    io.Copy(dst, src)
}
```

---

# Why `io.Copy()` is Powerful

Works with ANY:

* file
* network socket
* HTTP response
* buffer
* pipe

because Go uses interfaces.

---

# 16. Temporary Files

---

# `os.CreateTemp()`

```go
file, err := os.CreateTemp("", "temp-*.txt")
```

---

# Why Useful?

* testing
* caching
* uploads
* intermediate processing

---

# 17. Interfaces Behind File Handling

This is VERY important.

Go file handling is built around interfaces.

---

# `io.Reader`

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

Anything that can be read from.

Examples:

* files
* HTTP bodies
* strings
* buffers

---

# `io.Writer`

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

Anything writable.

---

# HUGE Go Philosophy

Go focuses on behavior, not inheritance.

Instead of:

```text
File extends Stream
```

Go says:

```text
File implements Reader and Writer
```

Very composable.

---

# 18. Reading JSON Files

Very common backend task.

---

# Example

```go
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

    data, _ := os.ReadFile("user.json")

    var user User

    json.Unmarshal(data, &user)
}
```

---

# 19. Writing JSON Files

```go
data, _ := json.Marshal(user)

os.WriteFile("user.json", data, 0644)
```

---

# 20. `os.WriteFile()`

Shortcut helper.

```go
os.WriteFile(
    "notes.txt",
    []byte("hello"),
    0644,
)
```

Equivalent to:

* create/open
* write
* close

---

# 21. Common Real-World Patterns

---

# Logging

```go
logFile, _ := os.OpenFile(
    "app.log",
    os.O_APPEND|os.O_CREATE|os.O_WRONLY,
    0644,
)
```

---

# CSV Processing

Usually:

* open file
* use `encoding/csv`
* scanner/writer

---

# Config Files

Usually:

* read file
* parse JSON/YAML/TOML

---

# Upload Systems

Usually:

* stream file
* avoid loading entire file into RAM

---

# 22. Common Mistakes

---

# Forgetting `Close()`

BAD.

---

# Ignoring Errors

BAD.

```go
data, _ := os.ReadFile("x")
```

acceptable only in tiny demos.

---

# Loading Huge Files into Memory

BAD:

```go
os.ReadFile("100GB.txt")
```

Use streaming.

---

# Forgetting `Flush()`

Buffered writers need flushing.

---

# Using Wrong File Flags

Very common beginner issue.

Example:

```go
os.O_APPEND
```

without:

```go
os.O_WRONLY
```

can fail.

---

# 23. Modern vs Old APIs

Old:

```go
ioutil.ReadFile()
ioutil.WriteFile()
```

Deprecated.

Modern:

```go
os.ReadFile()
os.WriteFile()
```

---

# 24. The Most Important Concept

The REAL core of Go file handling is:

# Streams + Interfaces

Everything revolves around:

```text
Reader
Writer
Closer
Seeker
```

That is why the same code can work with:

* files
* HTTP
* TCP sockets
* buffers
* compressed streams

without changing logic.

That is one of Go's greatest strengths in backend engineering.

---

File permissions in Go come from Unix/Linux permission systems.

In Go, we usually see them in functions like:

```go
os.OpenFile()
os.WriteFile()
os.Mkdir()
os.MkdirAll()
```

Example:

```go
os.WriteFile("test.txt", data, 0644)
```

That `0644` is a permission mode.

---

# The Good News

We do NOT need to memorize hundreds of permission combinations.

In real-world backend development, we mostly use around:

| Permission | Common Usage                   |
| ---------- | ------------------------------ |
| `0644`     | Normal files                   |
| `0755`     | Executable files/directories   |
| `0700`     | Private files/directories      |
| `0600`     | Sensitive private files        |
| `0777`     | Full permissions (rarely safe) |

If we deeply understand THESE, we understand almost everything important.

---

# 1. Understanding Permission Structure 🖥️

Permissions are:

```text
OWNER | GROUP | OTHERS
```

Example:

```text
0644
```

Break it:

```text
0 6 4 4
  | | |
  | | └── Others
  | └──── Group
  └────── Owner
```

---

# Why the Leading `0`?

This indicates:

# OCTAL (base-8)

Unix permissions are octal numbers.

Digits allowed:

```text
0-7
```

---

# 2. Permission Values

Each digit is made from:

| Value | Meaning |
| ----- | ------- |
| `4`   | Read    |
| `2`   | Write   |
| `1`   | Execute |

Permissions are combinations.

---

# 3. The Complete Essential Table

| Number | Binary | Meaning                |
| ------ | ------ | ---------------------- |
| `0`    | `000`  | No permission          |
| `1`    | `001`  | Execute                |
| `2`    | `010`  | Write                  |
| `3`    | `011`  | Write + Execute        |
| `4`    | `100`  | Read                   |
| `5`    | `101`  | Read + Execute         |
| `6`    | `110`  | Read + Write           |
| `7`    | `111`  | Read + Write + Execute |

This table is VERY important.

Once this clicks, permissions become easy.

---

# 4. Understanding `0644`

Break it:

```text
Owner = 6 = read + write
Group = 4 = read
Others = 4 = read
```

Meaning:

| User Type | Permission |
| --------- | ---------- |
| Owner     | read/write |
| Group     | read       |
| Others    | read       |

---

# Visual Form

```text
rw-r--r--
```

Explanation:

```text
rw-  r--  r--
│    │    │
│    │    └── others
│    └────── group
└────────── owner
```

---

# Why is `0644` Common?

Perfect for normal files.

Examples:

* config files
* text files
* JSON files
* logs
* source code

Owner can edit.

Others can only read.

---

# 5. Understanding `0755`

Break it:

```text
7 = read + write + execute
5 = read + execute
5 = read + execute
```

Visual:

```text
rwxr-xr-x
```

---

# Why is `0755` Common?

Used for:

* executable programs
* scripts
* directories

---

# VERY IMPORTANT: Directories Behave Differently

For files:

```text
execute = can run file
```

For directories:

```text
execute = can ENTER/traverse directory
```

Without execute permission:

we cannot access contents properly.

---

# Typical Directory Permission

```go
os.Mkdir("data", 0755)
```

Meaning:

* owner fully controls
* others can enter/read

---

# 6. Understanding `0700`

```text
rwx------
```

Only owner can access.

Nobody else can even read.

---

# Used For

* SSH keys
* secrets
* credentials
* private configs

---

# Example

```go
os.WriteFile("secret.txt", data, 0700)
```

---

# 7. Understanding `0600`

```text
rw-------
```

Owner can:

* read
* write

But NOT execute.

Nobody else can access.

---

# Very Common For

* password files
* API keys
* token storage

---

# 8. Understanding `0777`

```text
rwxrwxrwx
```

EVERYONE can:

* read
* write
* execute

---

# Usually BAD

Security nightmare.

Avoid unless absolutely necessary.

---

# Why Dangerous?

Anyone can:

* modify files
* delete files
* inject malicious content

---

# 9. Directory Permissions vs File Permissions

This confuses many beginners.

---

# File Permissions

| Permission | Meaning         |
| ---------- | --------------- |
| Read       | read contents   |
| Write      | modify contents |
| Execute    | run file        |

---

# Directory Permissions

| Permission | Meaning                |
| ---------- | ---------------------- |
| Read       | list files             |
| Write      | create/delete files    |
| Execute    | enter/access directory |

---

# Important Example

A directory may allow listing:

```text
r--
```

But without execute:

we cannot actually access files inside properly.

---

# 10. Essential Real-World Defaults

These are the ones most Go developers actually use.

---

# Normal File

```go
0644
```

---

# Private Sensitive File

```go
0600
```

---

# Directory

```go
0755
```

---

# Private Directory

```go
0700
```

---

# Executable Script/Binary

```go
0755
```

---

# 11. FileMode Type in Go

Permissions use:

```go
os.FileMode
```

Example:

```go
var mode os.FileMode = 0644
```

---

# 12. Symbolic Permission Style

Linux also shows permissions like:

```text
-rw-r--r--
```

Let's decode:

```text
- rw- r-- r--
```

First character:

| Symbol | Meaning      |
| ------ | ------------ |
| `-`    | regular file |
| `d`    | directory    |
| `l`    | symlink      |

Then:

```text
rw- = owner
r-- = group
r-- = others
```

---

# 13. Converting Permissions Mentally

This becomes easy with practice.

---

# Example

```text
rwxr-x---
```

Break:

```text
rwx = 7
r-x = 5
--- = 0
```

Result:

```text
0750
```

---

# Example 2

```text
rw-r-----
```

Break:

```text
rw- = 6
r-- = 4
--- = 0
```

Result:

```text
0640
```

---

# 14. Go Examples

---

# Creating Secure File

```go
os.WriteFile("secret.txt", data, 0600)
```

---

# Creating Public Readable File

```go
os.WriteFile("notes.txt", data, 0644)
```

---

# Creating Directory

```go
os.Mkdir("uploads", 0755)
```

---

# Creating Private Directory

```go
os.Mkdir("private", 0700)
```

---

# 15. Important Security Advice

Backend developers SHOULD care about permissions.

Bad permissions can expose:

* passwords
* API keys
* JWT secrets
* databases
* uploads

A huge amount of real-world breaches happen because permissions were too open.

---

# 16. What We Actually Need to Remember

Honestly, these are enough for most professional Go backend work:

| Mode   | Meaning              |
| ------ | -------------------- |
| `0644` | standard file        |
| `0755` | executable/directory |
| `0600` | private file         |
| `0700` | private directory    |
| `0777` | unsafe/open          |

And remember:

| Value | Meaning |
| ----- | ------- |
| `4`   | read    |
| `2`   | write   |
| `1`   | execute |

Everything else is just combinations.

```go
Embedding static files into a Go binary is one of the most useful modern Go features.

It allows us to package files INSIDE the compiled executable itself.

This became officially available in:

# Go 1.16

through:

```go id="d8ocmu"
embed
```

package.

---

# Why Embedding Exists

Normally, applications depend on external files:

* HTML templates
* CSS
* JS
* images
* config files
* SQL migrations
* text assets

Without embedding:

```text id="rk8dga"
binary + separate files/folders
```

This creates deployment problems.

---

# Example Problem Without Embedding

Suppose we build:

```text id="e0r7oc"
app.exe
```

but app needs:

```text id="6f6ngt"
templates/
static/
config/
```

If someone deletes/moves them:

application breaks.

---

# Embedding Solves This

We can compile everything INTO the executable:

```text id="o8p7pi"
ONE self-contained binary
```

Very powerful.

---

# Real-World Uses

Embedding is commonly used for:

| Use Case        | Example        |
| --------------- | -------------- |
| HTML templates  | web apps       |
| Static frontend | CSS/JS         |
| Images/icons    | desktop apps   |
| SQL migrations  | databases      |
| Default configs | CLIs           |
| Documentation   | help systems   |
| Certificates    | internal tools |

---

# 1. Importing `embed`

```go id="d9r5cs"
import "embed"
```

---

# VERY IMPORTANT

Even if we don't directly use `embed` package functions,

we STILL need:

```go id="6l8f0z"
import _ "embed"
```

or:

```go id="odjqlx"
import "embed"
```

because the compiler processes directives from it.

---

# 2. The `//go:embed` Directive

This is special compiler syntax.

Example:

```go id="n7mn3d"
//go:embed hello.txt
var content string
```

---

# IMPORTANT

This is NOT a normal comment.

Compiler reads it specially.

Must appear:

* immediately above variable declaration
* no blank line

---

# Simplest Example

---

# File Structure

```text id="ctyu8w"
project/
├── main.go
└── hello.txt
```

---

# hello.txt

```text id="jlwm2q"
Hello from embedded file!
```

---

# main.go

```go id="zjlwm8"
package main

import (
	_ "embed"
	"fmt"
)

//go:embed hello.txt
var data string

func main() {
	fmt.Println(data)
}
```

---

# Output

```text id="0rq3u6"
Hello from embedded file!
```

---

# What Happened?

At compile time:

Go reads:

```text id="e53g4k"
hello.txt
```

and stores contents INSIDE executable.

---

# 3. Embedding as `string`

Useful for:

* text files
* HTML
* JSON
* SQL

Example:

```go id="wz3h0k"
//go:embed query.sql
var sqlQuery string
```

---

# 4. Embedding as `[]byte`

Useful for:

* images
* binaries
* arbitrary data

Example:

```go id="mxrj6x"
//go:embed logo.png
var logo []byte
```

---

# Difference

| Type     | Best For   |
| -------- | ---------- |
| `string` | text       |
| `[]byte` | binary/raw |

---

# 5. Embedding Multiple Files

Example:

```go id="sv9nca"
//go:embed a.txt b.txt
var files embed.FS
```

---

# What is `embed.FS`?

This is a virtual file system.

Very important concept.

---

# 6. `embed.FS`

Type:

```go id="bhuz6z"
embed.FS
```

acts like:

# read-only filesystem

inside memory/binary.

---

# Example Structure

```text id="cb8j6q"
project/
├── main.go
├── a.txt
└── b.txt
```

---

# Example

```go id="v0g06n"
package main

import (
	"embed"
	"fmt"
)

//go:embed *.txt
var files embed.FS

func main() {

	data, err := files.ReadFile("a.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
```

---

# 7. Wildcards

Supported.

Example:

```go id="o7wn36"
//go:embed static/*
var staticFiles embed.FS
```

---

# This Embeds Entire Folder

```text id="hhm9ei"
static/
├── css/
├── js/
└── images/
```

---

# VERY Common in Web Apps

Especially for:

* frontend assets
* templates

---

# 8. Serving Embedded Static Files in Web Servers

SUPER common.

---

# Example

```go id="0x6z0x"
package main

import (
	"embed"
	"net/http"
)

//go:embed static/*
var staticFiles embed.FS

func main() {

	fs := http.FS(staticFiles)

	http.Handle("/", http.FileServer(fs))

	http.ListenAndServe(":8080", nil)
}
```

---

# What Happens?

Frontend files are served directly from executable.

No separate static folder needed.

---

# 9. Embedding HTML Templates

Very common.

---

# Example Structure

```text id="3w4t0e"
templates/
├── home.html
└── about.html
```

---

# Example

```go id="0yr7ry"
package main

import (
	"embed"
	"html/template"
)

//go:embed templates/*
var templateFS embed.FS

func main() {

	tmpl := template.Must(
		template.ParseFS(templateFS, "templates/*.html"),
	)

	_ = tmpl
}
```

---

# Why `ParseFS()`?

Because templates are inside embedded filesystem.

---

# 10. Directory Embedding

---

# Embed Entire Directory

```go id="pq5v3d"
//go:embed templates
var templates embed.FS
```

---

# Recursive?

YES.

Subdirectories included.

---

# 11. Build-Time Embedding

VERY important concept.

Embedding happens:

# during compilation

NOT runtime.

Meaning:

```text id="v2b53z"
changing original file later does NOT affect executable
```

We must rebuild app.

---

# Example

```text id="9lyux2"
go build
```

creates snapshot of files at build time.

---

# 12. Embedded Files are Read-Only

We cannot modify:

```go id="5nl9o1"
embed.FS
```

It's immutable.

---

# Why?

Because data is compiled into binary memory sections.

---

# 13. Common Mistake

Trying:

```go id="6h4f8w"
files.WriteFile(...)
```

Impossible.

`embed.FS` only supports reading.

---

# 14. Using `fs` Package

Very important.

`embed.FS` implements interfaces from:

```go id="jxv4ef"
io/fs
```

---

# Why This Matters

Same code can work with:

* real filesystem
* embedded filesystem
* zip filesystem
* network filesystem

because Go uses interfaces.

---

# 15. `fs.Sub()`

Very useful.

Suppose:

```text id="0c0tb8"
static/
    css/
    js/
```

---

# Example

```go id="c2bw9h"
subFS, err := fs.Sub(staticFiles, "static")
```

Now root becomes:

```text id="8wm0w3"
css/
js/
```

instead of:

```text id="2p7pw0"
static/css
```

---

# Commonly Used With HTTP Servers

```go id="s8vdnq"
http.FileServer(http.FS(subFS))
```

---

# 16. Limitations of Embedding

---

# Cannot Embed Outside Module

BAD:

```go id="lgdtx7"
//go:embed ../secret.txt
```

Not allowed.

---

# Cannot Use Runtime Variables

BAD:

```go id="mxg9mb"
//go:embed dynamicPath
```

Must be compile-time constant patterns.

---

# File Size Matters

Huge embedded assets:

* increase binary size
* increase memory usage

---

# 17. Binary Size Tradeoff

Embedding:

```text id="14j0y4"
simplicity ↑
binary size ↑
```

Tradeoff worth it for many apps.

---

# 18. Common Real-World Architectures

---

# Single Binary Web App

Very popular in Go.

Binary contains:

* backend
* frontend
* templates
* configs

Deploy ONE file.

---

# CLI Tools

Embedding:

* default config
* help docs
* shell scripts

---

# Desktop Apps

Embedding:

* icons
* fonts
* UI assets

---

# 19. Security Considerations

IMPORTANT.

Embedded assets are NOT encrypted.

People can often inspect binaries and extract assets.

Do NOT embed:

* secrets
* passwords
* API keys

---

# 20. Modern Go Philosophy Here

Embedding fits Go's philosophy perfectly:

```text id="7q0k48"
simple deployment
minimal dependencies
single executable
```

This is one reason Go became huge in:

* cloud
* DevOps
* backend
* CLI tooling

---

# 21. Full Real Example

---

# Structure

```text id="9vsvic"
project/
├── main.go
├── static/
│   └── style.css
└── templates/
    └── home.html
```

---

# main.go

```go id="jlwm4g"
package main

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates/*
var templateFS embed.FS

//go:embed static/*
var staticFS embed.FS

func main() {

	tmpl := template.Must(
		template.ParseFS(templateFS, "templates/*.html"),
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "home.html", nil)
	})

	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.FS(staticFS)),
		),
	)

	http.ListenAndServe(":8080", nil)
}
```

---

# 22. Most Important Concepts to Remember

---

# Core Concepts

| Concept      | Meaning                      |
| ------------ | ---------------------------- |
| `//go:embed` | compile-time embedding       |
| `embed.FS`   | virtual read-only filesystem |
| `string`     | text embedding               |
| `[]byte`     | binary embedding             |
| `ParseFS()`  | templates from embedded FS   |
| `http.FS()`  | serve embedded files         |

---

# MOST Important Mental Model

Think of embedding as:

```text id="wlcmmi"
copying files INTO executable during build
```

After compilation:

```text id="4cq25y"
binary already contains the files
```

No external dependency needed anymore.
```

