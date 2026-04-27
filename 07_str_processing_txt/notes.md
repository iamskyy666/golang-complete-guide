If we’re serious about writing clean Go, the `strings` package is something we’ll use *every single day*. It’s not just utilities—it’s how we safely and efficiently manipulate text without reinventing the wheel.

Let’s break it down properly so it actually sticks.

---

# 🔥 What is the `strings` package?

The **`strings` package** provides functions to:

* search
* compare
* modify
* split
* build

👉 all kinds of **string manipulation**

Import:

```go
import "strings"
```

---

# 🧠 Important foundation (don’t skip this)

In Go:

* strings are **immutable**
* every modification → creates a **new string**

```go
s := "hello"
s2 := strings.ToUpper(s)

fmt.Println(s)  // hello
fmt.Println(s2) // HELLO
```

👉 This is why performance matters (we’ll cover `Builder` later)

---

# 🔍 1. Searching Functions

## 📌 `Contains`

```go
strings.Contains("golang", "go") // true
```

---

## 📌 `HasPrefix` / `HasSuffix`

```go
strings.HasPrefix("golang", "go") // true
strings.HasSuffix("file.txt", ".txt") // true
```

---

## 📌 `Index`

```go
strings.Index("golang", "lang") // 2
strings.Index("golang", "xyz")  // -1
```

👉 returns first occurrence

---

## 📌 `LastIndex`

```go
strings.LastIndex("go go go", "go") // 6
```

---

# ✂️ 2. Splitting Strings

## 📌 `Split`

```go
parts := strings.Split("a,b,c", ",")
// ["a", "b", "c"]
```

---

## 📌 `SplitN`

```go
strings.SplitN("a,b,c,d", ",", 2)
// ["a", "b,c,d"]
```

---

## 📌 `Fields`

```go
strings.Fields("  hello   world  ")
// ["hello", "world"]
```

👉 splits by whitespace (very useful)

---

# 🔗 3. Joining Strings

## 📌 `Join`

```go
arr := []string{"go", "is", "awesome"}
result := strings.Join(arr, " ")
// "go is awesome"
```

---

# ✏️ 4. Modifying Strings

## 📌 `Replace`

```go
strings.Replace("go go go", "go", "GO", 2)
// "GO GO go"
```

---

## 📌 `ReplaceAll`

```go
strings.ReplaceAll("go go go", "go", "GO")
// "GO GO GO"
```

---

## 📌 `ToUpper` / `ToLower`

```go
strings.ToUpper("go") // "GO"
strings.ToLower("GO") // "go"
```

---

## 📌 `Trim` family

### `Trim`

```go
strings.Trim("!!hello!!", "!") // "hello"
```

### `TrimSpace`

```go
strings.TrimSpace("   hello   ") // "hello"
```

### `TrimPrefix` / `TrimSuffix`

```go
strings.TrimPrefix("go-lang", "go-") // "lang"
```

---

# 🔁 5. Comparing Strings

## 📌 `EqualFold` (case-insensitive)

```go
strings.EqualFold("Go", "go") // true
```

---

## 📌 `Compare`

```go
strings.Compare("a", "b") // -1
```

Returns:

* 0 → equal
* -1 → less
* 1 → greater

👉 rarely used in practice

---

# 🔄 6. Repeating Strings

```go
strings.Repeat("go", 3)
// "gogogo"
```

---

# 🧱 7. strings.Builder (VERY IMPORTANT)

This is where most beginners mess up.

---

## ❌ Bad approach (inefficient)

```go
result := ""
for i := 0; i < 5; i++ {
	result += "go"
}
```

👉 creates new string every time → slow

---

## ✅ Good approach (efficient)

```go
var b strings.Builder

for i := 0; i < 5; i++ {
	b.WriteString("go")
}

result := b.String()
```

👉 uses internal buffer → **fast and memory-efficient**

---

## 📌 Builder methods

```go
b.WriteString("hello")
b.WriteByte('!')
b.WriteRune('✓')

str := b.String()
```

---

# 🧠 When to use Builder?

Use it when:

* building strings in loops
* concatenating many strings
* performance matters

---

# 🔧 8. Other Useful Functions

## 📌 `Count`

```go
strings.Count("go go go", "go") // 3
```

---

## 📌 `ContainsAny`

```go
strings.ContainsAny("hello", "xyz") // false
```

---

## 📌 `ContainsRune`

```go
strings.ContainsRune("hello", 'e') // true
```

---

## 📌 `Map` (advanced)

```go
result := strings.Map(func(r rune) rune {
	if r == 'a' {
		return 'A'
	}
	return r
}, "banana")
// "bAnAnA"
```

---

# ⚠️ Common Mistakes

## ❌ Thinking strings are mutable

They are NOT.

---

## ❌ Overusing `+` in loops

Kills performance.

---

## ❌ Ignoring `TrimSpace`

This causes real bugs (especially in APIs)

---

# 🧠 Real-world usage patterns

### ✅ Parsing CSV-like data

```go
strings.Split(data, ",")
```

---

### ✅ Cleaning user input

```go
strings.TrimSpace(input)
```

---

### ✅ Case-insensitive comparison

```go
strings.EqualFold(email1, email2)
```

---

### ✅ Building JSON / logs manually

```go
strings.Builder
```

---

# 🎯 Final Mental Model

Think of `strings` as:

> 🔧 A toolbox for safe, efficient text manipulation

---

# 🚀 Quick Summary

* Strings are **immutable**
* Use:

  * `Split`, `Join` → structure
  * `Trim`, `Replace` → clean
  * `Contains`, `Index` → search
  * `Builder` → performance
* `Builder` is critical for real-world apps

---

If we strip away the surface, `strings.Builder` is really about one thing:

> **building strings efficiently without wasting memory**

Most beginners underestimate this—but in real systems (logs, JSON building, templating, APIs), this becomes *critical*.

Let’s go deep.

---

# 🔥 What is `strings.Builder`?

`strings.Builder` is a type from the `strings` package used to **efficiently construct strings step-by-step**.

```go
var b strings.Builder
```

---

# 🧠 Why do we even need it?

Because of this:

### ❌ Strings are immutable in Go

```go
result := ""
result += "go"
result += "lang"
```

👉 Every `+=`:

* creates a **new string**
* copies old data
* wastes memory

---

## ⚠️ What actually happens internally

```go
result = result + "go"
```

becomes:

1. allocate new memory
2. copy old string
3. append new string

👉 Do this in a loop → **O(n²) behavior**

---

# ✅ Builder solves this

`strings.Builder`:

* uses an internal **byte buffer**
* grows it dynamically
* avoids repeated allocations

👉 Result → **O(n)** instead of O(n²)

---

# 🧱 Basic Usage

```go
var b strings.Builder

b.WriteString("Go")
b.WriteString(" ")
b.WriteString("Lang")

result := b.String()

fmt.Println(result) // Go Lang
```

---

# ⚡ Core Methods (we should know these well)

## 1. `WriteString`

```go
b.WriteString("hello")
```

* most commonly used
* appends string

---

## 2. `WriteByte`

```go
b.WriteByte('!')
```

* adds a single byte
* faster than string for single chars

---

## 3. `WriteRune`

```go
b.WriteRune('✓')
```

* for Unicode characters
* safe for multi-byte characters

---

## 4. `String()`

```go
result := b.String()
```

* returns final string
* does **NOT copy unnecessarily**

---

# 🧠 Example: Loop (real-world case)

### ❌ Bad

```go
result := ""
for i := 0; i < 5; i++ {
	result += "go"
}
```

---

### ✅ Good

```go
var b strings.Builder

for i := 0; i < 5; i++ {
	b.WriteString("go")
}

result := b.String()
```

---

# ⚡ Performance Insight (this matters)

| Approach          | Time Complexity | Memory      |
| ----------------- | --------------- | ----------- |
| `+=`              | ❌ O(n²)         | ❌ High      |
| `strings.Builder` | ✅ O(n)          | ✅ Efficient |

---

# 🔧 Advanced: `Grow()` (hidden gem)

We can pre-allocate memory:

```go
var b strings.Builder
b.Grow(100) // allocate 100 bytes upfront
```

👉 Use when we roughly know final size

This avoids:

* repeated resizing
* internal reallocations

---

# 🧠 Example with Grow

```go
var b strings.Builder
b.Grow(50)

for i := 0; i < 5; i++ {
	b.WriteString("hello")
}
```

---

# ⚠️ Important Rules (people mess this up)

## ❌ Don’t copy a Builder

```go
b2 := b // BAD
```

👉 This can cause bugs because Builder has internal state

---

## ❌ Don’t use after copying

Always pass pointer if needed:

```go
func build(b *strings.Builder) {
	b.WriteString("data")
}
```

---

## ❌ Don’t mix with manual string building

Stick to Builder once we start using it.

---

# 🧠 When should we use Builder?

Use it when:

* building strings in loops
* concatenating many pieces
* generating logs / JSON / HTML
* performance matters

---

## ❌ When NOT needed

Don’t over-engineer:

```go
name := "Go" + "Lang"
```

👉 This is fine

---

# 🔁 Builder vs bytes.Buffer (important comparison)

| Feature            | strings.Builder | bytes.Buffer    |
| ------------------ | --------------- | --------------- |
| Works with strings | ✅ best          | ok              |
| Works with bytes   | ❌               | ✅               |
| Performance        | ✅ optimized     | slightly slower |
| Simplicity         | ✅               | medium          |

👉 Rule:

* **strings → Builder**
* **bytes → Buffer**

---

# 🧠 Internal Design (high-level)

`strings.Builder`:

* wraps a `[]byte` buffer
* grows capacity automatically
* avoids copying until `String()` is called

---

# 🚀 Real-world Example

### Building a CSV line

```go
var b strings.Builder

fields := []string{"John", "Doe", "30"}

for i, f := range fields {
	b.WriteString(f)
	if i != len(fields)-1 {
		b.WriteString(",")
	}
}

csv := b.String()
fmt.Println(csv) // John,Doe,30
```

---

# 🎯 Mental Model

If we remember just one thing:

> `strings.Builder` = efficient string construction using a growing buffer

---

# 🚀 Final Summary

* Strings are immutable → concatenation is expensive
* `strings.Builder` avoids repeated allocations
* Key methods:

  * `WriteString`
  * `WriteByte`
  * `WriteRune`
  * `String`
  * `Grow`
* Use in loops and heavy concatenation
* Don’t copy it

---

*FORMATTING VERBS* - This is one of those areas where most devs memorize a few verbs (`%s`, `%d`) and move on—but if we actually understand the full system, we get *precision control* over logs, debugging, formatting APIs, everything.

Let’s break it down properly.

---

# 🔥 Where formatting verbs are used

All of this belongs to the `fmt` package:

```go
fmt.Printf()
fmt.Sprintf()
fmt.Fprintf()
```

👉 Formatting verbs = placeholders like `%v`, `%d`, `%s`

---

# 🧠 Core idea

```go
fmt.Printf("Value: %d", 10)
```

* `%d` → tells Go how to format the value
* `10` → actual value

---

# 🧱 1. General-purpose verbs (MOST IMPORTANT)

## 📌 `%v` — default format

```go
fmt.Printf("%v", 10)        // 10
fmt.Printf("%v", "hello")   // hello
```

👉 “just print it”

---

## 📌 `%+v` — include struct field names

```go
type User struct {
	Name string
	Age  int
}

u := User{"Skyy", 25}
fmt.Printf("%+v", u)
// {Name:Skyy Age:25}
```

---

## 📌 `%#v` — Go syntax representation

```go
fmt.Printf("%#v", u)
// main.User{Name:"Skyy", Age:25}
```

👉 very useful for debugging

---

## 📌 `%T` — type

```go
fmt.Printf("%T", 10) // int
```

---

## 📌 `%%` — literal percent

```go
fmt.Printf("100%%")
// 100%
```

---

# 🔢 2. Integer verbs

## 📌 `%d` — decimal

```go
fmt.Printf("%d", 42) // 42
```

---

## 📌 `%b` — binary

```go
fmt.Printf("%b", 5) // 101
```

---

## 📌 `%o` — octal

```go
fmt.Printf("%o", 8) // 10
```

---

## 📌 `%x` / `%X` — hexadecimal

```go
fmt.Printf("%x", 255) // ff
fmt.Printf("%X", 255) // FF
```

---

## 📌 `%c` — character

```go
fmt.Printf("%c", 65) // A
```

---

# 🔤 3. String & byte verbs

## 📌 `%s` — string

```go
fmt.Printf("%s", "go") // go
```

---

## 📌 `%q` — quoted string

```go
fmt.Printf("%q", "go")
// "go"
```

---

## 📌 `%x` — hex encoding of string

```go
fmt.Printf("%x", "go")
// 676f
```

---

# 🔡 4. Rune (Unicode) verbs

## 📌 `%c`

```go
fmt.Printf("%c", 'A') // A
```

---

## 📌 `%U` — Unicode format

```go
fmt.Printf("%U", 'A')
// U+0041
```

---

# 🔢 5. Floating-point verbs

## 📌 `%f` — decimal

```go
fmt.Printf("%f", 3.14) // 3.140000
```

---

## 📌 `%.2f` — precision

```go
fmt.Printf("%.2f", 3.14159) // 3.14
```

---

## 📌 `%e` / `%E` — scientific notation

```go
fmt.Printf("%e", 1000.0)
// 1.000000e+03
```

---

## 📌 `%g` — compact format

```go
fmt.Printf("%g", 1000.0) // 1000
```

👉 chooses best format automatically

---

# 🔗 6. Pointer verb

## 📌 `%p`

```go
x := 10
fmt.Printf("%p", &x)
```

👉 prints memory address

---

# 🧩 7. Boolean verb

## 📌 `%t`

```go
fmt.Printf("%t", true) // true
```

---

# 🧠 8. Width & Precision (VERY IMPORTANT)

This is where formatting becomes powerful.

---

## 📌 Width

```go
fmt.Printf("%5d", 42)
```

Output:

```
   42
```

---

## 📌 Left align

```go
fmt.Printf("%-5d", 42)
```

Output:

```
42   
```

---

## 📌 Zero padding

```go
fmt.Printf("%05d", 42)
```

Output:

```
00042
```

---

## 📌 Precision (floats)

```go
fmt.Printf("%.2f", 3.14159)
// 3.14
```

---

## 📌 Precision (strings)

```go
fmt.Printf("%.3s", "golang")
// gol
```

---

# 🧠 9. Argument indexing

```go
fmt.Printf("%[2]d %[1]d", 10, 20)
```

Output:

```
20 10
```

👉 reorder arguments

---

# ⚡ 10. Formatting with Sprintf

Instead of printing:

```go
str := fmt.Sprintf("Age: %d", 25)
```

👉 returns string

---

# 🧠 11. Interaction with Stringer

If a type implements:

```go
String() string
```

Then:

```go
fmt.Printf("%v", obj)
```

👉 calls `String()` automatically

---

# ⚠️ Common Mistakes

## ❌ Wrong verb type

```go
fmt.Printf("%d", "hello") // ERROR
```

---

## ❌ Forgetting precision

```go
fmt.Printf("%f", 3.1) // 3.100000
```

---

## ❌ Overusing `%v`

Good for debugging, but not for precise output

---

# 🧠 Real-world usage patterns

### Logging

```go
fmt.Printf("User %s logged in at %v", name, time.Now())
```

---

### Financial formatting

```go
fmt.Printf("Balance: %.2f", balance)
```

---

### Debugging structs

```go
fmt.Printf("%+v", obj)
```

---

# 🎯 Mental Model

Think of formatting verbs as:

> “Instructions to Go on how to represent data”

---

# 🚀 Quick Cheat Sheet

| Verb  | Meaning            |
| ----- | ------------------ |
| `%v`  | default            |
| `%+v` | struct with fields |
| `%#v` | Go syntax          |
| `%T`  | type               |
| `%d`  | integer            |
| `%f`  | float              |
| `%s`  | string             |
| `%q`  | quoted             |
| `%x`  | hex                |
| `%t`  | boolean            |
| `%p`  | pointer            |

---

# 🚀 Final Insight

If we really want to level up:

* `%v` → debugging
* `%+v` → structs
* `%.2f` → money
* `%#v` → deep debugging

Master these → we cover 90% of real-world use.

---

If we don’t understand runes properly, everything with strings, Unicode, even validation logic starts breaking in subtle ways.

---

# 🔥 What is a Rune in Go?

A **rune** is:

> an alias for `int32` that represents a **Unicode code point**

```go
type rune = int32
```

So when we write:

```go
var r rune = 'A'
```

👉 we’re storing the Unicode value of `'A'` (which is 65)

---

# 🧠 Rune vs Byte (critical distinction)

| Concept | Type    | Meaning           |
| ------- | ------- | ----------------- |
| byte    | `uint8` | raw data (8 bits) |
| rune    | `int32` | Unicode character |

---

## Example:

```go
s := "A"
```

* `'A'` → 1 byte → 1 rune

```go
s := "你"
```

* `'你'` → 3 bytes → 1 rune

👉 This is the whole point of runes.

---

# 🔢 Rune values (Unicode code points)

```go
r := 'A'
fmt.Println(r) // 65
```

```go
r := '你'
fmt.Println(r) // 20320
```

👉 These numbers are **Unicode code points**

---

# 🔤 Rune literal syntax

```go
r1 := 'A'
r2 := '你'
r3 := '😊'
```

👉 Single quotes = rune
👉 Double quotes = string

---

# ⚠️ Important rule

```go
'a'   // rune
" a " // string
```

---

# 🔁 Iterating with runes (this is where they shine)

```go
s := "Go你好"

for _, r := range s {
	fmt.Printf("%c ", r)
}
```

👉 Output:

```
G o 你 好
```

👉 `range` automatically converts bytes → runes

---

# ⚠️ Why not use indexing?

```go
s := "你"
fmt.Println(s[0])
```

👉 This prints a **byte**, not a rune → wrong

---

## ✅ Correct way

```go
runes := []rune(s)
fmt.Println(runes[0])        // Unicode value
fmt.Println(string(runes[0])) // 你
```

---

# 🔄 Converting between string and rune

## String → rune slice

```go
s := "hello"
r := []rune(s)
```

---

## Rune slice → string

```go
str := string(r)
```

---

# 🧠 Why rune slice matters

Because it gives us **character-level control**:

```go
s := "你好"
r := []rune(s)

r[0] = '我'

fmt.Println(string(r)) // 我好
```

👉 This is impossible with raw string indexing.

---

# 🔧 Common rune operations

## Print rune as character

```go
fmt.Printf("%c", 'A')
```

---

## Print Unicode code point

```go
fmt.Printf("%U", 'A') // U+0041
```

---

# 🔍 Using runes with `unicode` package

```go
import "unicode"

unicode.IsLetter('A') // true
unicode.IsDigit('5')  // true
unicode.IsSpace(' ')  // true
```

---

## Case conversion

```go
unicode.ToUpper('a') // 'A'
unicode.ToLower('A') // 'a'
```

---

# ⚠️ Subtle but important: Rune ≠ Visual Character

This is where most devs get surprised.

Example:

```go
"é"
```

Can be:

* 1 rune (`é`)
* OR 2 runes (`e` + accent)

👉 Both look identical

---

## So:

> rune = Unicode code point
> NOT always what the user *visually sees*

---

# 🔥 Runes in string building

```go
var b strings.Builder

b.WriteRune('你')
b.WriteRune('好')

fmt.Println(b.String()) // 你好
```

👉 Always use `WriteRune` for Unicode safety

---

# ⚠️ Common mistakes

## ❌ Treating string index as character

```go
s[0] // ❌ byte, not rune
```

---

## ❌ Using len() for characters

```go
len("你好") // 6 bytes, not 2 characters
```

---

## ❌ Ignoring rune conversion

Leads to:

* broken slicing
* corrupted Unicode

---

# 🧠 When should we use runes?

Use runes when:

* iterating characters
* modifying text
* handling Unicode input
* validation (letters, digits, etc.)

---

# 🚀 Practical example

## Reverse a string (Unicode-safe)

```go
func reverse(s string) string {
	r := []rune(s)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}
```

👉 Works for:

* ASCII
* Chinese
* Emojis

---

# 🎯 Mental Model

If we remember one thing:

> rune = one Unicode code point
> string = sequence of bytes (UTF-8)

---

# 🚀 Quick Summary

* rune = `int32`
* represents Unicode character
* needed because:

  * strings are bytes
* use:

  * `range` → iteration
  * `[]rune` → indexing/modification
* essential for real-world text handling

---

If we don’t understand Unicode properly in Go, we *will* write bugs—especially when dealing with real user input (names, emojis, international text). So let’s build this from the ground up and make it solid.

---

# 🔥 The core problem

In Go:

> A **string is a sequence of bytes**, not characters.

```go
s := "Hello"
```

👉 looks simple, but internally:

* `"Hello"` = 5 bytes → fine (ASCII)

Now:

```go
s := "你好"
```

👉 This is where things change:

* `"你"` = 3 bytes
* `"好"` = 3 bytes
* total = **6 bytes**

---

# 🧠 Key Concepts (must be clear)

## 1. Byte (`uint8`)

* raw data
* what strings are made of

---

## 2. Rune (`int32`)

* represents a **Unicode code point**
* basically = “character” (not always visually exact, but close enough)

```go
var r rune = '你'
```

---

## 3. UTF-8 encoding

Go strings are encoded in:

👉 **UTF-8**

* variable-length encoding
* 1 to 4 bytes per character

---

# ⚠️ The classic beginner mistake

```go
s := "你好"

fmt.Println(len(s)) // ❗ 6, not 2
```

👉 `len()` returns **bytes**, not characters

---

# ✅ Correct way to count characters

## Option 1: convert to rune slice

```go
runes := []rune(s)
fmt.Println(len(runes)) // 2
```

---

## Option 2: use utf8 package

```go
import "unicode/utf8"

utf8.RuneCountInString(s) // 2
```

---

# 🔁 Iterating over strings (VERY IMPORTANT)

## ❌ Wrong way (byte-wise)

```go
for i := 0; i < len(s); i++ {
	fmt.Println(s[i])
}
```

👉 breaks for Unicode

---

## ✅ Correct way (rune-wise)

```go
for _, r := range s {
	fmt.Printf("%c\n", r)
}
```

👉 Go automatically decodes UTF-8

---

# 🔍 Example

```go
s := "Go你好"

for i, r := range s {
	fmt.Printf("Index: %d, Rune: %c\n", i, r)
}
```

👉 Output:

* index jumps (because multi-byte chars)

---

# ⚠️ Indexing problem

```go
s := "你好"
fmt.Println(s[0]) // ❗ NOT '你'
```

👉 gives first **byte**, not character

---

## ✅ Correct way

```go
r := []rune(s)
fmt.Println(string(r[0])) // 你
```

---

# 🔧 Unicode helpers (`unicode` package)

Go gives us tools to work with Unicode properly.

---

## 📌 Check character type

```go
import "unicode"

unicode.IsLetter('A') // true
unicode.IsDigit('9')  // true
unicode.IsSpace(' ')  // true
```

---

## 📌 Case conversion

```go
unicode.ToUpper('a') // 'A'
unicode.ToLower('A') // 'a'
```

---

# 🔄 Transforming strings safely

## ❌ Wrong way

```go
strings.ToUpper(s)
```

👉 works for many cases, but not always Unicode-safe for complex scripts

---

## ✅ Advanced way (using runes)

```go
result := strings.Map(func(r rune) rune {
	if unicode.IsLetter(r) {
		return unicode.ToUpper(r)
	}
	return r
}, s)
```

---

# 🧠 Combining characters (advanced insight)

Some characters are not single runes.

Example:

```go
"é"
```

Can be:

* single rune (é)
* OR:

  * 'e' + accent

👉 visually same, internally different

---

## ⚠️ Why this matters

```go
len([]rune("é")) // may be 1 or 2 depending on encoding
```

👉 This is where normalization comes in (advanced topic)

---

# 🔥 Strings.Builder + Unicode

When building Unicode strings:

```go
var b strings.Builder

b.WriteRune('你')
b.WriteRune('好')

fmt.Println(b.String()) // 你好
```

👉 Always use `WriteRune` for safety

---

# 🧠 Real-world scenarios

## ✅ User input (names, emails)

Always assume Unicode:

* "José"
* "李雷"

---

## ✅ Emoji handling

```go
s := "😊"
len(s) // 4 bytes
```

---

## ✅ Text processing

* search
* replace
* validation

---

# ⚠️ Common mistakes

## ❌ Using `len()` for character count

Wrong for Unicode

---

## ❌ Indexing strings directly

Breaks multi-byte chars

---

## ❌ Treating string as ASCII

Real-world apps are NOT ASCII

---

# 🎯 Mental Model

If we remember one thing:

> string = bytes
> rune = character (Unicode code point)

---

# 🚀 Practical Rules

* Use `range` → iterate strings
* Use `rune` → when working with characters
* Use `utf8` → count runes
* Use `strings.Builder + WriteRune` → build safely

---

# 🚀 Quick Summary

* Go strings = UTF-8 encoded bytes
* Characters = runes (`int32`)
* `len()` ≠ number of characters
* Use:

  * `range` → iteration
  * `[]rune` → indexing
  * `utf8.RuneCountInString` → counting
* Unicode is everywhere—code must handle it

---

Regex in Go is powerful—but also easy to misuse if we treat it like a magic hammer. The goal is to know **when and how** to use it efficiently.

Let’s go step by step so it actually sticks.

---

# 🔥 What is Regex in Go?

Regex (regular expressions) lets us:

* match patterns
* validate strings
* extract data
* replace parts of text

In Go, this is handled by:

```go
import "regexp"
```

---

# 🧠 Important foundation

Go uses the **RE2 engine** (not PCRE like many languages)

👉 That means:

* ✅ safe (no catastrophic backtracking)
* ❌ no advanced features like backtracking, lookbehind

This is a *design choice for performance and safety*

---

# 🧱 1. Creating a Regex

## 📌 `Compile` (safe way)

```go
re, err := regexp.Compile(`go`)
if err != nil {
	panic(err)
}
```

---

## 📌 `MustCompile` (common in practice)

```go
re := regexp.MustCompile(`go`)
```

👉 panics if invalid
👉 used when pattern is known beforehand

---

# 🔍 2. Matching

## 📌 `MatchString`

```go
re := regexp.MustCompile(`go`)
fmt.Println(re.MatchString("golang")) // true
```

---

## 📌 Direct function

```go
matched, _ := regexp.MatchString(`go`, "golang")
```

👉 quick but less flexible

---

# 🔎 3. Finding matches

## 📌 `FindString`

```go
re := regexp.MustCompile(`go`)
fmt.Println(re.FindString("golang go")) // "go"
```

---

## 📌 `FindAllString`

```go
re := regexp.MustCompile(`go`)
fmt.Println(re.FindAllString("go go go", -1))
// ["go", "go", "go"]
```

---

## 📌 `FindStringIndex`

```go
re := regexp.MustCompile(`go`)
fmt.Println(re.FindStringIndex("golang"))
// [0 2]
```

👉 start & end index

---

# 🧩 4. Capturing groups

```go
re := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
```

---

## 📌 `FindStringSubmatch`

```go
match := re.FindStringSubmatch("test@gmail.com")

fmt.Println(match)
```

Output:

```go
[
  "test@gmail.com", // full match
  "test",
  "gmail",
  "com",
]
```

---

# 🔁 5. Replace operations

## 📌 `ReplaceAllString`

```go
re := regexp.MustCompile(`go`)
result := re.ReplaceAllString("go go go", "GO")

fmt.Println(result) // GO GO GO
```

---

## 📌 Using groups in replacement

```go
re := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
result := re.ReplaceAllString("test@gmail.com", "$1 at $2 dot $3")

fmt.Println(result)
// test at gmail dot com
```

---

# ✂️ 6. Splitting with regex

```go
re := regexp.MustCompile(`[,\s]+`)
parts := re.Split("go,java python", -1)

// ["go", "java", "python"]
```

---

# 🔤 7. Common regex patterns

## 📌 Digits

```go
\d
```

---

## 📌 Word characters

```go
\w
```

---

## 📌 Whitespace

```go
\s
```

---

## 📌 Any character

```go
.
```

---

## 📌 Repetition

```go
a+   // one or more
a*   // zero or more
a?   // optional
```

---

## 📌 Anchors

```go
^start
end$
```

---

# 🧠 Example: Email validation

```go
re := regexp.MustCompile(`^[\w._%+-]+@[\w.-]+\.[a-zA-Z]{2,}$`)

fmt.Println(re.MatchString("test@gmail.com")) // true
```

---

# ⚠️ Important Go-specific limitations

Because of RE2:

## ❌ No lookbehind

```regex
(?<=abc)def
```

👉 not supported

---

## ❌ Limited backtracking

👉 avoids performance issues but reduces flexibility

---

# ⚡ Performance Tip (VERY IMPORTANT)

## ❌ Bad

```go
for i := 0; i < 1000; i++ {
	regexp.MatchString(`go`, "golang")
}
```

👉 compiles regex every time

---

## ✅ Good

```go
re := regexp.MustCompile(`go`)

for i := 0; i < 1000; i++ {
	re.MatchString("golang")
}
```

👉 compile once, reuse

---

# 🧠 Raw strings (backticks) are your friend

Always prefer:

```go
re := regexp.MustCompile(`\d+`)
```

Instead of:

```go
"\\d+"
```

👉 cleaner, less error-prone

---

# 🔥 Real-world use cases

## ✅ Validation

* email
* phone
* password rules

---

## ✅ Parsing logs

```go
ERROR: 2026-01-01 message
```

---

## ✅ Extracting data

* URLs
* IDs
* tokens

---

## ⚠️ When NOT to use regex

Don’t force regex when:

* simple string functions work
* readability suffers

👉 Example:

```go
strings.Contains(s, "@") // better than regex
```

---

# 🎯 Mental Model

Think of regex as:

> a pattern engine for matching and extracting structured text

---

# 🚀 Quick Summary

* Use `regexp` package
* Prefer `MustCompile`
* Key methods:

  * `MatchString`
  * `FindString`
  * `FindAllString`
  * `ReplaceAllString`
* Use raw strings (`` ` ``)
* Compile once for performance
* RE2 = safe but limited

---

# 🚀 Honest Advice

Regex is powerful—but dangerous when overused.

> If a problem can be solved with `strings` package, do that first.
> Use regex when pattern matching becomes complex.

---

If we’re building real applications—emails, configs, HTML pages, logs—**text templating** becomes unavoidable. Go’s approach is clean, safe, and surprisingly powerful once we understand the flow.

Let’s break it down in a way we can actually *use*.

---

# 🔥 What is `text/template`?

The `text/template` package lets us:

> **generate dynamic text output using templates + data**

It’s logic-less *enough* to stay clean, but powerful enough to be practical.

---

# 🧠 Core idea

We define a template:

```go id="x5x6cz"
"Hello, {{.Name}}!"
```

Then pass data:

```go id="3a8lgs"
{Name: "Skyy"}
```

👉 Output:

```
Hello, Skyy!
```

---

# 🧱 Basic Example

```go id="tn9h5f"
package main

import (
	"os"
	"text/template"
)

func main() {
	tmpl := template.Must(template.New("example").Parse("Hello, {{.Name}}!"))

	data := struct {
		Name string
	}{
		Name: "Skyy",
	}

	tmpl.Execute(os.Stdout, data)
}
```

---

# 🧠 Key Concept: `{{.}}` (dot)

The `.` (dot) means:

> “current data context”

---

## Example:

```go id="wm7y1x"
{{.Name}}
```

👉 Access `Name` field from struct

---

# 🔁 Working with Structs

```go id="1vlhpd"
type User struct {
	Name string
	Age  int
}
```

Template:

```go id="dzzpnp"
"Name: {{.Name}}, Age: {{.Age}}"
```

---

# 🧩 Working with Maps

```go id="y8ssfp"
data := map[string]string{
	"name": "Skyy",
}
```

Template:

```go id="ntunhf"
{{.name}}
```

---

# 🔁 Loops (`range`)

```go id="6fwls8"
{{range .}}
- {{.}}
{{end}}
```

---

## Example:

```go id="r61kqg"
data := []string{"Go", "Java", "Python"}
```

Output:

```
- Go
- Java
- Python
```

---

# 🔀 Conditionals (`if`)

```go id="b0n9rp"
{{if .IsAdmin}}
Admin User
{{else}}
Regular User
{{end}}
```

---

# 🧠 Nested data

```go id="ucwuvx"
type User struct {
	Name string
	Address struct {
		City string
	}
}
```

Template:

```go id="6j6qvf"
{{.Address.City}}
```

---

# 🔧 Functions in templates

Go lets us use built-in functions.

## Example:

```go id="9gqjqn"
{{len .}}
```

---

## Custom functions (VERY IMPORTANT)

```go id="0pxr5m"
func toUpper(s string) string {
	return strings.ToUpper(s)
}

tmpl := template.Must(
	template.New("t").
	Funcs(template.FuncMap{
		"upper": toUpper,
	}).
	Parse("{{upper .Name}}"),
)
```

---

# 📂 Parsing templates from files

```go id="48mm8y"
tmpl := template.Must(template.ParseFiles("template.txt"))
```

---

# 🔗 Multiple templates

```go id="txh2r3"
tmpl := template.Must(template.ParseFiles("header.txt", "body.txt"))
```

---

# 🔥 Template composition (important)

```go id="3rfmvy"
{{define "header"}}Header Content{{end}}
{{template "header"}}
```

👉 reuse templates

---

# ⚠️ Error handling

```go id="9dy7oy"
tmpl, err := template.New("t").Parse("...")
if err != nil {
	panic(err)
}
```

---

# ⚡ `template.Must`

```go id="nxq9x2"
tmpl := template.Must(template.New("t").Parse("..."))
```

👉 panics if error
👉 used when template is static

---

# 🧠 `Execute` vs `ExecuteTemplate`

## `Execute`

```go id="zslz2k"
tmpl.Execute(os.Stdout, data)
```

---

## `ExecuteTemplate`

```go id="gr5d4b"
tmpl.ExecuteTemplate(os.Stdout, "header", data)
```

---

# 🔒 Security note

`text/template`:

* safe for text

For HTML:

👉 use `html/template` (auto-escapes)

---

# ⚠️ Common mistakes

## ❌ Forgetting dot context

```go id="yr3n1o"
{{Name}} // ❌ wrong
{{.Name}} // ✅ correct
```

---

## ❌ Overloading logic

Templates should NOT contain heavy logic.

Bad:

```go id="n2lxqj"
{{if and (gt .Age 18) (lt .Age 60)}}
```

👉 move logic to Go code

---

## ❌ Ignoring errors in Execute

```go id="4s2r9l"
tmpl.Execute(...)
```

👉 always check error in real apps

---

# 🧠 Real-world use cases

## ✅ Email templates

* OTP emails
* welcome messages

---

## ✅ Config generation

* YAML / JSON

---

## ✅ CLI output formatting

---

## ✅ Code generation

---

# 🧠 Mental Model

Think of templates as:

> **data + structure = output**

---

# 🚀 Full Example (practical)

```go id="5q6c2n"
type Product struct {
	Name  string
	Price float64
}

tmpl := template.Must(template.New("p").Parse(`
Product List:
{{range .}}
- {{.Name}}: ${{.Price}}
{{end}}
`))

products := []Product{
	{"Laptop", 999.99},
	{"Phone", 499.99},
}

tmpl.Execute(os.Stdout, products)
```

---

# 🚀 Quick Summary

* `text/template` → generate dynamic text
* `{{.}}` → current data
* `range`, `if` → control flow
* `Funcs()` → custom logic
* `ParseFiles()` → real templates
* `Execute()` → render output

---

# 🚀 Honest Advice

Don’t treat templates as mini programming languages.

> Keep logic in Go, keep templates for structure.

---


# CODE-FILES 💻
```go
package main

import (
	"fmt"
	"strings"
)

// 📂 07_str_processing_txt
// strings and documentation

func main() {
	s1:=" I am Batman"
	s2:=strings.Clone(s1)

	fmt.Println(s1,s2)

	b:= strings.Builder{}
	b.Write([]byte("Builder.Write() example.. with []byte"))
	fmt.Println(b.String())

	bStr:=strings.Builder{}
	bStr.WriteString("Builder.WriteString() example.. with 'str'.")
	fmt.Println(bStr.String())

	// More..
	fmt.Println(strings.ToLower(s1))
	fmt.Println(strings.ToUpper(s1))
	fmt.Println("Length of s1:",len(s1))
	fmt.Println(strings.TrimSpace(s1))
	fmt.Println("Length of s1 after TrimSpace():",len(s1))

	fmt.Println(strings.HasSuffix("test@hotmail.com","hotmail.com"))
	fmt.Println(strings.HasPrefix("test@hotmail.com","test"))
	fmt.Println(strings.Replace("test@hotmail.com","hot","cold",1))

	parts:=strings.Split("skyy@gmail.com","@")
	username,domain:=parts[0],parts[1]

	fmt.Println("username:",username,",","domain:",domain)
	
	fields:=strings.Fields("Hello there")
	field1,field2:=fields[0],fields[1]
	fmt.Println("field1:",field1,",","field2:",field2)

	fmt.Println(strings.Join(fields,","))

}

// $ go run main.go
//  I am Batman  I am Batman
// Builder.Write() example.. with []byte
// Builder.WriteString() example.. with 'str'.
//  i am batman
//  I AM BATMAN
// Length of s1: 12
// I am Batman
// Length of s1 after TrimSpace(): 12
// true
// true
// test@coldmail.com
// username: skyy , domain: gmail.com
// field1: Hello , field2: there
// Hello,there
```
```go
package main

import (
	"errors"
	"fmt"
)

// 📂 07_str_processing_txt
// formatting strs. using formatting-verbs

type ConfigItem struct{
	Key string
	Value any
	IsSet bool
}


func (ci ConfigItem)String()string{
	return fmt.Sprintf("Key: %s, Value: %s, IsSet: %t",ci.Key, ci.Value, ci.IsSet)
}

func main() {
	
	appName:="EnvParser"
	version:= 1.2
	port:=8080
	isEnabled:=true

	status:=fmt.Sprintf("Application: %s (Version %.1f) running on PORT %d. Enabled: %t",appName, version, port, isEnabled)

	configItem1:=ConfigItem{Key: "API_URL",Value: "http://localhost:3000/api",IsSet: true}
	configItem2:=ConfigItem{Key: "TIMEOUT_MS",Value: 5000,IsSet: true}
	configItem3:=ConfigItem{Key: "DEBUG_MODE",Value: false,IsSet: false}

	fmt.Println("STATUS:",status)

	fmt.Printf("Config-Item 1 (%%v) => %v\n",configItem1)
	fmt.Printf("Config-Item 2 (%%+v) => %+v\n",configItem2)
	fmt.Printf("Config-Item 3 (%%#v)=> %#v\n",configItem3)

	err:=errors.New("testing err.")
	errInfo:=fmt.Errorf("here is the error on port %d: %w",port,err)
	fmt.Println(errInfo) // print it/use it


}

// $ go run main.go
// STATUS: Application: EnvParser (Version 1.2) running on PORT 8080. Enabled: true
// Config-Item 1 (%v) => Key: API_URL, Value: http://localhost:3000/api, IsSet: true
// Config-Item 2 (%+v) => Key: TIMEOUT_MS, Value: %!s(int=5000), IsSet: true
// Config-Item 3 (%#v)=> main.ConfigItem{Key:"DEBUG_MODE", Value:false, IsSet:false}
// here is the error on port 8080: testing err.

```

```go
package main

import (
	"fmt"
	"unicode"
)

// 📂 07_str_processing_txt
// Working with unicode chars. in Golang

func main() {
	username:="test"
	fmt.Println(len(username)) // 4 -> Think in terms of chars (4 bytes)

	test1:="বাংলা বর্ণমালা"
	fmt.Println(len(test1)) // 40 -> different encoding, UTF-8 (variable-length)
	fmt.Println(test1[0]) // 224 -> ASCII code ❌ PROBLEM

	for _,i:=range test1{
		fmt.Println(string(i))
	}

	// DIRECT LOOPING -> best way with UNICODE -> runes & range
	// UNICODE PKG.
	zh:= []rune{'汉','字','漢'}
	for _,i:=range zh{
		fmt.Println(string(i),unicode.IsLetter(i),unicode.IsLower(i))
	}

	/*
	$ go run main.go
4
40
224
ব
া
ং
ল
া
 
ব
র

ণ
ম
া
ল
া
汉 true false
字 true false
漢 true false
	*/
}
 ```
 ```go
 package main

import (
	"fmt"
	"os"
	"regexp"
)

// 📂 07_str_processing_txt
// using regex with Golang
// website to play: https://www.regexpal.com/
func main() {
	txt1:="Hello World! Welcome to Go!"
	regex1, err:=regexp.Compile(`Go`)

	if err != nil {
		fmt.Println("⚠️ ERROR:",err)
		os.Exit(1)
		//OR
		// log.Fatal(err)
	}

	fmt.Printf("Text1 '%s', matches 'Go':%t\n",txt1, regex1.MatchString(txt1)) // true

	txt2:="Product codes: P1234, X567, P789"
	rProductP:=regexp.MustCompile(`P\d+`)
	firstProduct:=rProductP.FindString(txt2)
	allProducts:=rProductP.FindAllString(txt2,-1)
	fmt.Println(firstProduct) // P1234
	fmt.Println(allProducts) // [P1234 P789]
}
 ```
 ```go
 package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

// 📂 07_str_processing_txt
// working with text/templates

type EmailData struct{
	RecipentName string
	SenderName string
	Subject string
	Body string
	Items []string // loop
	UnreadCount int
}

func main() {
	fmt.Println("------------ TEXT/TEMPLATE -----------")
	emailTemplate:=`
	Subject: {{ .Subject }}
	{{.Body}}


	{{if .Items}}
	Related Items: 
	{{range .Items}}
	-{{.}}
	{{end}}
	{{end}}

	{{if gt .UnreadCount 0}}

	You have {{.UnreadCount}} unread-messages.
	{{else}}
	You have no messages!
	{{end}}


	- Thanks
	{{.SenderName}} 😊
	`
	tmpltObj,err:= template.New("email-message").Parse(emailTemplate)
	if err != nil {
		fmt.Println("ERROR:",err.Error())
		os.Exit(1)
	}

	// sample-data
	emailTest := EmailData{
	RecipentName: "Rahul Sharma",
	SenderName:   "Skyy",
	Subject:      "Your Weekly Updates",
	Body:         "Hi Rahul,\nHere are your latest updates:",
	Items: []string{
		"New course available: Go Concurrency",
		"50% discount on premium plan",
		"Reminder: Complete your profile",
	},
	UnreadCount: 3,
    }

	var output strings.Builder

	// Strings output on the BUFFER
	err=tmpltObj.Execute(&output,emailTest)
	if err != nil {
		fmt.Println("ERROR:",err.Error())
		os.Exit(1)
	}

	fmt.Println(strings.ToUpper(output.String()))
	fmt.Println("------------------------------------------------")

	// Output on the screen (stdout)
	err=tmpltObj.Execute(os.Stdout,emailTest)
	if err != nil {
		fmt.Println("ERROR:",err.Error())
		os.Exit(1)
	}

}

// $ go run main.go
// ------------ TEXT/TEMPLATE -----------

//         SUBJECT: YOUR WEEKLY UPDATES
//         HI RAHUL,
// HERE ARE YOUR LATEST UPDATES:



//         RELATED ITEMS: 

//         -NEW COURSE AVAILABLE: GO CONCURRENCY

//         -50% DISCOUNT ON PREMIUM PLAN

//         -REMINDER: COMPLETE YOUR PROFILE





//         YOU HAVE 3 UNREAD-MESSAGES.



//         - THANKS
//         SKYY 😊

// ------------------------------------------------

//         Subject: Your Weekly Updates
//         Hi Rahul,
// Here are your latest updates:



//         Related Items: 

//         -New course available: Go Concurrency

//         -50% discount on premium plan

//         -Reminder: Complete your profile





//         You have 3 unread-messages.



//         - Thanks
//         Skyy 😊
 ```