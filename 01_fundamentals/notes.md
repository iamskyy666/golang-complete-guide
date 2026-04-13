#data-types
#variables
#constants

# 🧠 Big Picture: What are Data Types in Go?

In Go, a **data type defines**:

* What kind of data we store
* How much memory it uses
* What operations we can perform on it

Go is:

> **Statically typed + strongly typed**

Meaning:

* Type is known at compile time
* No implicit conversions (strict behavior)

---

# 🔥 1. Basic (Primitive) Data Types

These are the building blocks.

---

## 🔢 Integers

### Types:

```go
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
```

### Example:

```go
var age int = 25
```

---

### 👉 Key Concepts

### 1. Signed vs Unsigned

* `int` → can be negative
* `uint` → only positive

---

### 2. Size matters

| Type  | Size | Range       |
| ----- | ---- | ----------- |
| int8  | 1B   | -128 to 127 |
| int32 | 4B   | large       |
| int64 | 8B   | very large  |

---

### ⚠️ Important

`int` is:

* **platform dependent**

  * 32-bit system → int32
  * 64-bit system → int64

👉 That’s why in production:

* Use `int64` for DB, APIs

---

## 🔢 Floating Point

```go
float32, float64
```

### Example:

```go
var price float64 = 99.99
```

---

### ⚠️ Precision issue

```go
fmt.Println(0.1 + 0.2) // not exactly 0.3
```

👉 Because floating-point math is approximate.

---

## 🔤 Strings

```go
var name string = "Skyy"
```

---

### 👉 Important Properties

* Immutable (cannot change directly)
* Stored as **byte slices internally**

---

### Example:

```go
name := "Go"
fmt.Println(len(name)) // 2
```

---

### 🧠 UTF-8 Gotcha

```go
s := "😊"
fmt.Println(len(s)) // 4 (bytes, not characters)
```

👉 Use `rune` for proper character handling.

---

## 🔣 Boolean

```go
var isLoggedIn bool = true
```

Only:

```go
true / false
```

---

# ⚡ 2. Derived (Composite) Types

Now we move into **real power of Go**

---

## 📦 Arrays

```go
var nums [3]int = [3]int{1, 2, 3}
```

---

### Key Properties:

* Fixed size
* Same type elements

---

### ⚠️ Important

```go
[3]int != [4]int
```

👉 Different types entirely

---

---

## 🔗 Slices (VERY IMPORTANT)

This is what we actually use instead of arrays.

```go
nums := []int{1, 2, 3}
```

---

### 🧠 Internals

A slice has:

* Pointer to array
* Length
* Capacity

---

### Example:

```go
nums := []int{1,2,3}
nums = append(nums, 4)
```

---

### Why slices matter

* Dynamic size
* Flexible
* Used everywhere (APIs, DB, JSON)

---

---

## 🗺️ Maps (Hashmaps)

```go
m := map[string]int{
  "a": 1,
  "b": 2,
}
```

---

### Key Features:

* Key-value store
* Fast lookup

---

### Example:

```go
value, ok := m["a"]
```

👉 `ok` tells if key exists

---

---

## 🧩 Structs (Custom Types)

This is where real backend modeling starts.

```go
type User struct {
  Name string
  Age  int
}
```

---

### Example:

```go
u := User{Name: "Skyy", Age: 29}
```

---

### Why structs matter

* Represent real-world data
* Used in:

  * APIs
  * DB models
  * JSON

---

---

# 🔁 3. Type System Concepts (VERY IMPORTANT)

---

## 🧠 Zero Values

Go gives default values automatically:

| Type   | Zero Value |
| ------ | ---------- |
| int    | 0          |
| string | ""         |
| bool   | false      |
| slice  | nil        |
| map    | nil        |

---

### Example:

```go
var x int
fmt.Println(x) // 0
```

---

## 🔄 Type Conversion (Explicit only)

```go
var a int = 10
var b float64 = float64(a)
```

---

👉 No automatic conversion like JS

---

## 🧠 Type Inference

```go
x := 10
```

Go infers:

```go
int
```

---

# ⚔️ 4. Value Types vs Reference Types

This is where many beginners struggle—so pay attention.

---

## 📦 Value Types

* int, float, bool, array, struct

👉 Copy is made

```go
a := 10
b := a
b = 20

fmt.Println(a) // still 10
```

---

## 🔗 Reference Types

* slice, map, pointer, channel

👉 Shared memory

```go
a := []int{1,2}
b := a
b[0] = 100

fmt.Println(a) // [100,2]
```

---

# 🧠 5. Special Types

---

## 🔹 Rune (character)

```go
var ch rune = 'A'
```

👉 Alias for `int32`

---

## 🔹 Byte

```go
var b byte = 255
```

👉 Alias for `uint8`

---

## 🔹 Pointer

```go
var x int = 10
var p *int = &x
```

---

### Why pointers?

* Avoid copying large data
* Modify original value

---

---

# 🧱 Mental Model (Important)

Think like this:

* **Basic types** → raw data
* **Composite types** → structured data
* **Reference types** → shared data
* **Structs** → real-world models

---

# ⚠️ Common Beginner Mistakes

Let me correct these early:

### ❌ Treating slices like arrays

→ They are NOT the same

---

### ❌ Forgetting map initialization

```go
var m map[string]int
m["a"] = 1 // ❌ panic
```

✔ Fix:

```go
m := make(map[string]int)
```

---

### ❌ Ignoring zero values

→ Can cause bugs silently

---

# 🚀 What actually matters for us (practical focus)

If we’re building backend apps:

Focus heavily on:

* ✅ Structs
* ✅ Slices
* ✅ Maps
* ✅ Pointers (basic understanding)

Don’t over-focus on:

* int8, int16 unless needed

---

If data types are **what kind of data**, variables and constants are **how we store and manage that data in memory**.

Let’s actually *understand how Go behaves*

---

# 🧠 1. Variables in Go

## 👉 What is a variable?

A **variable** is:

> A named storage location in memory whose value can change

---

## 🔹 Basic Declaration

### Full form (explicit)

```go
var age int = 25
```

---

### Type inference (very common)

```go
var age = 25
```

Go figures out:

```go
int
```

---

### Short declaration (MOST used)

```go
age := 25
```

⚠️ Only allowed **inside functions**

---

## 🧠 What’s actually happening?

When we write:

```go
age := 25
```

Go:

1. Allocates memory
2. Stores value `25`
3. Labels it as `age`
4. Infers type `int`

---

## 🔁 Updating Variables

```go
age := 25
age = 30
```

👉 Variables are mutable (can change)

---

## 📦 Multiple Variables

```go
var a, b, c int = 1, 2, 3
```

or:

```go
x, y := 10, 20
```

---

## 🔹 Zero Values (VERY IMPORTANT)

If we don’t initialize:

```go
var x int
var s string
var b bool
```

We get:

| Type   | Default |
| ------ | ------- |
| int    | 0       |
| string | ""      |
| bool   | false   |

👉 No “undefined” like JavaScript—Go is predictable

---

## 🔹 Block Declaration

```go
var (
  name = "Skyy"
  age  = 29
)
```

Clean and used in real projects.

---

# ⚠️ Important Rules for Variables

---

## ❌ No unused variables

```go
x := 10 // ❌ error if unused
```

Go forces clean code.

---

## ❌ No redeclaration in same scope

```go
x := 10
x := 20 // ❌ error
```

✔ Correct:

```go
x = 20
```

---

## ⚠️ Shadowing (tricky)

```go
x := 10

if true {
  x := 20 // new variable, not same!
}
```

👉 This creates a **new variable**, not update

This causes real bugs if we’re careless.

---

# 🔥 2. Constants in Go

## 👉 What is a constant?

A **constant** is:

> A value that cannot change after being assigned

---

## 🔹 Declaration

```go
const pi = 3.14
```

---

## 🔹 Typed vs Untyped Constants

### Untyped (flexible)

```go
const x = 10
```

👉 Can behave as int, float, etc.

---

### Typed

```go
const x int = 10
```

👉 Strict type

---

## 🧠 Why untyped constants are powerful

```go
const x = 10

var y int64 = x // ✅ works
```

👉 Go adapts automatically

---

## 🔒 Constants cannot change

```go
const x = 10
x = 20 // ❌ error
```

---

## ❌ No runtime values

```go
const x = time.Now() // ❌ not allowed
```

👉 Constants must be known at **compile time**

---

# 🔥 3. iota (VERY IMPORTANT for enums)

This is where Go gets interesting.

---

## 👉 What is iota?

`iota` is:

> A counter that increments automatically

---

## Example:

```go
const (
  A = iota // 0
  B        // 1
  C        // 2
)
```

---

## Real-world usage (Enums)

```go
const (
  Pending = iota
  Approved
  Rejected
)
```

👉 Output:

```
Pending = 0
Approved = 1
Rejected = 2
```

---

## Advanced iota

```go
const (
  _ = iota
  KB = 1 << (10 * iota)
  MB
  GB
)
```

👉 Used for:

* File sizes
* Bit flags

---

# ⚔️ Variables vs Constants (Clear Comparison)

| Feature         | Variable     | Constant     |
| --------------- | ------------ | ------------ |
| Change value    | ✅ Yes        | ❌ No         |
| Runtime allowed | ✅ Yes        | ❌ No         |
| Memory          | Allocated    | Optimized    |
| Use case        | Dynamic data | Fixed values |

---

# 🧠 Mental Model (This will help long-term)

Think like this:

* **Variables** → things that change
  (user input, API data, DB values)

* **Constants** → things that never change
  (PI, status codes, config values)

---

# ⚠️ Common Beginner Mistakes

---

## ❌ Using := everywhere blindly

Remember:

* Only inside functions

---

## ❌ Not understanding shadowing

This one silently breaks logic.

---

## ❌ Using variables instead of constants

Bad:

```go
status := "ACTIVE"
```

Better:

```go
const StatusActive = "ACTIVE"
```

---

# 🚀 What actually matters for us (practical focus)

If we’re building backend apps:

Focus on:

* `:=` (short declaration)
* Proper variable scoping
* Using constants for:

  * Status
  * Roles
  * Config values

---

Go looks simple on the surface but has **important design trade-offs underneath**.

Let’s be clear upfront:

> ❗ Go does **NOT have real enums** like TypeScript, Java, or C#

Instead, Go uses:

> **Constants + custom types + iota** to simulate enums

If we understand this properly, we’ll write **clean, safe, production-level code**.

---

# 🧠 1. What is an Enum (Conceptually)?

An enum is:

> A type that has a fixed set of possible values

Example (conceptually):

* OrderStatus → Pending, Shipped, Delivered

---

# 🔥 2. How Go Implements Enums

We build enums using **3 pieces together**:

1. `type` → defines a custom type
2. `const` → defines fixed values
3. `iota` → auto-increments values

---

## ✅ Basic Example

```go
type Status int

const (
  Pending Status = iota
  Approved
  Rejected
)
```

---

## 👉 What’s happening?

* `Status` → new type (not just int)
* `Pending = 0`
* `Approved = 1`
* `Rejected = 2`

---

## Usage:

```go
var s Status = Pending
```

---

# ⚠️ Why not just use int?

Bad:

```go
var status int = 1
```

👉 Problem:

* What is `1`? No meaning

---

Good:

```go
var status Status = Approved
```

👉 Now it's readable and type-safe

---

# 🧠 3. Strong Typing Advantage

Go enums are **type-safe**

```go
type Status int
type Role int

var s Status = 1
var r Role = 1

s = r // ❌ error
```

👉 Even though both are int underneath, Go prevents mixing

---

# 🔥 4. String-based Enums (VERY COMMON in APIs)

Sometimes numbers aren’t ideal.

---

## Example:

```go
type Status string

const (
  Pending  Status = "PENDING"
  Approved Status = "APPROVED"
  Rejected Status = "REJECTED"
)
```

---

## Why use string enums?

* Better for JSON APIs
* Readable in logs
* No confusion

---

## Example JSON:

```json
{
  "status": "APPROVED"
}
```

---

# ⚔️ Int vs String Enums (Important Decision)

| Type        | Pros                   | Cons             |
| ----------- | ---------------------- | ---------------- |
| int enum    | fast, memory efficient | not readable     |
| string enum | readable, API-friendly | slightly heavier |

---

👉 Practical advice:

* Internal logic → int enums
* APIs / DB → string enums

---

# 🔁 5. iota Deep Dive (Core of Enums)

---

## Basic:

```go
const (
  A = iota // 0
  B        // 1
  C        // 2
)
```

---

## Reset behavior

`iota` resets in each `const` block:

```go
const (
  A = iota // 0
)

const (
  B = iota // 0 again
)
```

---

## Skipping values

```go
const (
  A = iota
  _
  C
)
```

👉 A=0, C=2

---

## Custom increments

```go
const (
  A = iota * 10 // 0
  B             // 10
  C             // 20
)
```

---

## Bitmask enums (ADVANCED but important)

```go
const (
  Read = 1 << iota  // 1
  Write             // 2
  Execute           // 4
)
```

---

👉 Used for:

* Permissions
* Flags

---

# 🧠 6. Adding Methods to Enums (POWERFUL)

This is where Go becomes elegant.

---

## Example:

```go
type Status int

const (
  Pending Status = iota
  Approved
  Rejected
)
```

---

## Add String method:

```go
func (s Status) String() string {
  switch s {
  case Pending:
    return "PENDING"
  case Approved:
    return "APPROVED"
  case Rejected:
    return "REJECTED"
  default:
    return "UNKNOWN"
  }
}
```

---

## Usage:

```go
fmt.Println(Pending.String()) // PENDING
```

---

👉 This is how Go replaces built-in enum features

---

# 🔄 7. Enum Validation (IMPORTANT in real apps)

Go doesn’t enforce valid enum values automatically.

---

## Problem:

```go
var s Status = 100 // ❌ allowed!
```

---

## Solution:

```go
func IsValidStatus(s Status) bool {
  switch s {
  case Pending, Approved, Rejected:
    return true
  }
  return false
}
```

---

👉 You must validate manually

---

# 🧱 8. Real Backend Example

Let’s make it real for your MERN/backend mindset:

```go
type OrderStatus string

const (
  OrderPending   OrderStatus = "PENDING"
  OrderShipped   OrderStatus = "SHIPPED"
  OrderDelivered OrderStatus = "DELIVERED"
)
```

---

## Struct:

```go
type Order struct {
  ID     string
  Status OrderStatus
}
```

---

👉 Clean, readable, API-friendly

---

# ⚠️ Common Mistakes (I’ll stop you early)

---

## ❌ Using plain strings everywhere

```go
status := "PENDING"
```

👉 Leads to:

* Typos
* Bugs
* No validation

---

## ❌ Not creating custom type

```go
const Pending = "PENDING"
```

👉 Weak typing → bad practice

---

## ❌ Ignoring validation

Go won’t protect you automatically.

---

# 🧠 Final Mental Model

Think like this:

* Go enums = **typed constants**
* `iota` = auto-number generator
* Custom type = safety layer
* Methods = behavior layer

---

# 🚀 What actually matters for us

For backend development, focus on:

* ✅ String enums for APIs
* ✅ Custom types for safety
* ✅ Validation functions
* ✅ Adding methods (like `.String()`)

---


