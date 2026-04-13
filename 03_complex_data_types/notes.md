If we truly understand **arrays vs slices**, we’ll avoid a ton of bugs (especially in APIs, DB work, and concurrency).

Let’s go deep—but in a way that actually builds intuition.

---

# 🧠 1. Arrays in Go

## 👉 What is an array?

An **array** is:

> A fixed-size collection of elements of the same type

---

## 🔹 Declaration

```go
var nums [3]int = [3]int{1, 2, 3}
```

or:

```go
nums := [3]int{1, 2, 3}
```

---

## 🔹 Key Properties

### ✅ Fixed size (VERY IMPORTANT)

```go
[3]int
```

👉 Size is part of the type

---

### ⚠️ Different sizes = different types

```go
var a [3]int
var b [4]int

a = b // ❌ error
```

---

## 🔹 Accessing elements

```go
nums[0] = 10
fmt.Println(nums[1])
```

---

## 🔹 Length

```go
len(nums) // 3
```

---

## 🔹 Default values

```go
var nums [3]int
```

👉 Output:

```go
[0 0 0]
```

---

## 🔹 Arrays are VALUE types

This is critical.

```go
a := [3]int{1,2,3}
b := a

b[0] = 100

fmt.Println(a) // [1 2 3]
```

👉 Copy is made → no shared memory

---

## 🔥 Why arrays are rarely used directly

Because:

* Fixed size
* Hard to pass around
* Not flexible

👉 That’s why Go gives us **slices**

---

# 🚀 2. Slices (THE REAL DEAL)

## 👉 What is a slice?

A **slice** is:

> A dynamic, flexible view over an underlying array

---

## 🔹 Declaration

```go
nums := []int{1, 2, 3}
```

---

## 🧠 Internal Structure (VERY IMPORTANT)

A slice is NOT just data.

It contains:

1. Pointer → to underlying array
2. Length → current size
3. Capacity → max size before reallocation

---

### Visual:

```
Slice:
[ pointer ] → [1 2 3 0 0]
[ length = 3 ]
[ capacity = 5 ]
```

---

## 🔹 Length vs Capacity

```go
nums := []int{1,2,3}

len(nums) // 3
cap(nums) // 3
```

---

## 🔹 Using make()

```go
nums := make([]int, 3, 5)
```

👉 Creates:

* length = 3
* capacity = 5

---

# 🔥 3. Append (Core Operation)

```go
nums := []int{1,2,3}
nums = append(nums, 4)
```

---

## 🧠 What happens internally?

If capacity is enough:

* Same array used

If capacity exceeded:

* New array created
* Data copied

---

👉 This is where bugs can happen (sharing vs copying)

---

# ⚠️ 4. Slice is a REFERENCE type

```go
a := []int{1,2,3}
b := a

b[0] = 100

fmt.Println(a) // [100 2 3]
```

👉 Both point to same underlying array

---

# 🔪 5. Slicing (Sub-slices)

```go
nums := []int{1,2,3,4,5}

sub := nums[1:4]
```

👉 Output:

```go
[2 3 4]
```

---

## 🧠 Important

* Start index inclusive
* End index exclusive

---

## ⚠️ Shared memory again

```go
sub[0] = 100
fmt.Println(nums) // affected!
```

---

👉 Slices share underlying array unless copied

---

# 🔁 6. Copying Slices (IMPORTANT)

To avoid shared memory issues:

```go
a := []int{1,2,3}
b := make([]int, len(a))

copy(b, a)
```

---

👉 Now independent

---

# 🔥 7. Nil vs Empty Slice

---

## Nil slice

```go
var nums []int
```

👉 nums == nil → true

---

## Empty slice

```go
nums := []int{}
```

👉 nums != nil

---

### Why this matters

In APIs / JSON:

* `nil` → null
* `[]` → empty array

---

# 🔄 8. Iteration with slices

```go
nums := []int{10,20,30}

for i, v := range nums {
  fmt.Println(i, v)
}
```

---

# ⚠️ Common Mistakes (Very Important)

---

## ❌ Thinking slice is fully independent

```go
a := []int{1,2,3}
b := a
```

👉 Not a copy!

---

## ❌ Modifying slice inside range incorrectly

```go
for _, v := range nums {
  v *= 2 // ❌ doesn't update original
}
```

✔ Fix:

```go
for i := range nums {
  nums[i] *= 2
}
```

---

## ❌ Forgetting append returns new slice

```go
append(nums, 4) // ❌ lost result
```

✔ Always:

```go
nums = append(nums, 4)
```

---

## ❌ Unexpected mutation via subslices

This causes real bugs in production.

---

# ⚔️ Arrays vs Slices (Clear Comparison)

| Feature            | Array      | Slice          |
| ------------------ | ---------- | -------------- |
| Size               | Fixed      | Dynamic        |
| Type includes size | ✅ Yes      | ❌ No           |
| Memory             | Value type | Reference type |
| Flexibility        | Low        | High           |
| Usage              | Rare       | Everywhere     |

---

# 🧠 Mental Model (Lock this in)

Think like this:

* **Array** → actual storage (rarely used directly)
* **Slice** → handle/view on top of array (what we use daily)

---

# 🚀 Real Backend Perspective

In real apps, we use slices for:

* API responses:

```go
[]User
```

* DB queries:

```go
[]Order
```

* JSON arrays:

```json
[]
```

---

# 🔥 Final Insight (This separates beginners from pros)

> A slice is NOT the data—it’s a **window into data**

That’s why:

* Changes can affect others
* Append can create new memory
* Copying matters

---

Arrays and slices are only useful if we can **manipulate and iterate over them cleanly**.

Let’s focus on:

* The **important built-in functions (“methods”)**
* The **right way to loop**
* The **gotchas that break real apps**

---

# 🧠 First: Important Reality

Go does NOT have “methods” on slices like JavaScript (`map`, `filter`, etc.)

Instead, Go gives us:

> **Built-in functions + patterns**

So we write logic explicitly. That’s by design.

---

# 🔥 1. Core Built-in Functions (Must Know)

---

## ✅ `len()` → length

```go
nums := []int{1,2,3}
fmt.Println(len(nums)) // 3
```

👉 Works for:

* arrays
* slices
* maps
* strings

---

## ✅ `cap()` → capacity (slice only)

```go
nums := make([]int, 3, 5)
fmt.Println(cap(nums)) // 5
```

👉 Important for performance understanding

---

## ✅ `append()` → add elements

```go
nums := []int{1,2}
nums = append(nums, 3)
```

---

### ⚠️ Key Rule

Always reassign:

```go
nums = append(nums, 4)
```

---

### Append multiple

```go
nums = append(nums, 5, 6, 7)
```

---

### Append another slice

```go
a := []int{1,2}
b := []int{3,4}

a = append(a, b...)
```

👉 `...` is required

---

---

## ✅ `copy()` → copy slices

```go
a := []int{1,2,3}
b := make([]int, len(a))

copy(b, a)
```

---

### Why this matters

Without copy:
👉 slices share memory → bugs

---

---

## ✅ `make()` → create slice

```go
nums := make([]int, 3, 5)
```

* length = 3
* capacity = 5

---

---

## ✅ `new()` (rare for slices)

```go
p := new([]int)
```

👉 Not commonly used for slices

---

---

# 🔁 2. Looping Over Arrays & Slices

This is where most beginners either:

* Write unidiomatic code
* Or introduce subtle bugs

Let’s fix that.

---

## 🔹 1. Classic for loop

```go
nums := []int{10,20,30}

for i := 0; i < len(nums); i++ {
  fmt.Println(nums[i])
}
```

---

### When to use this?

* When we need index control
* When modifying elements

---

---

## 🔹 2. `range` loop (MOST IMPORTANT)

```go
nums := []int{10,20,30}

for i, v := range nums {
  fmt.Println(i, v)
}
```

---

### 🔥 Key behavior

* `i` → index
* `v` → **copy of value**

---

## ⚠️ VERY IMPORTANT GOTCHA

```go
for _, v := range nums {
  v = v * 2 // ❌ doesn't change original
}
```

👉 Because `v` is a copy

---

✔ Correct:

```go
for i := range nums {
  nums[i] *= 2
}
```

---

---

## 🔹 3. Ignore index or value

```go
for _, v := range nums {
  fmt.Println(v)
}
```

---

```go
for i := range nums {
  fmt.Println(i)
}
```

---

---

## 🔹 4. Reverse loop

```go
for i := len(nums)-1; i >= 0; i-- {
  fmt.Println(nums[i])
}
```

---

---

## 🔹 5. Infinite + manual break

```go
i := 0

for {
  if i >= len(nums) {
    break
  }
  fmt.Println(nums[i])
  i++
}
```

---

👉 Rare, but useful in advanced flows

---

# 🔥 3. Common Patterns (VERY IMPORTANT)

Go doesn’t give `map/filter/reduce`—we build them.

---

## 🔹 Filter pattern

```go
nums := []int{1,2,3,4}

var result []int

for _, v := range nums {
  if v % 2 == 0 {
    result = append(result, v)
  }
}
```

---

## 🔹 Map (transform) pattern

```go
nums := []int{1,2,3}

result := make([]int, len(nums))

for i, v := range nums {
  result[i] = v * 2
}
```

---

## 🔹 Reduce (sum)

```go
sum := 0

for _, v := range nums {
  sum += v
}
```

---

---

## 🔹 Remove element (IMPORTANT)

```go
nums := []int{1,2,3,4}

// remove index 1
nums = append(nums[:1], nums[2:]...)
```

---

👉 This is a classic Go trick

---

---

## 🔹 Insert element

```go
nums = append(nums[:i], append([]int{val}, nums[i:]...)...)
```

---

👉 Looks ugly, but very common

---

# ⚠️ 4. Real Gotchas (Pay attention)

---

## ❌ Forgetting append reassignment

```go
append(nums, 4) // ❌ lost
```

---

## ❌ Shared slice memory

```go
a := []int{1,2,3}
b := a
```

👉 Changes affect both

---

## ❌ Range variable pointer bug

```go
for _, v := range nums {
  go func() {
    fmt.Println(v) // ❌ tricky bug in concurrency
  }()
}
```

👉 Fix:

```go
for _, v := range nums {
  v := v
  go func() {
    fmt.Println(v)
  }()
}
```

---

## ❌ Modifying slice while iterating

Dangerous:

```go
for i := range nums {
  nums = append(nums, i) // ❌ unpredictable
}
```

---

# 🧠 Mental Model

Think like this:

* `len` → how many elements
* `cap` → how much room
* `append` → grow slice
* `copy` → avoid shared memory
* `range` → iterate safely (but returns copy)

---

# 🚀 What actually matters for us (backend focus)

We’ll use these constantly for:

* API responses (`[]User`)
* Filtering data
* Transforming DB results
* Building JSON arrays

---

# 🔥 Final Advice (No fluff)

If we don’t master slices:
👉 We’ll keep hitting weird bugs

If we do:
👉 Go becomes extremely predictable and powerful

---

Maps are one of the **most used and most misunderstood** parts of Go. If we really understand them, a lot of backend work (APIs, caching, lookups, grouping) becomes much easier.

Let’s go deep, but keep it practical.

---

# 🧠 1. What is a Map in Go?

A **map** is:

> A collection of key-value pairs (like a hash table / dictionary)

---

## 🔹 Basic syntax

```go
m := map[string]int{
  "a": 1,
  "b": 2,
}
```

---

## 👉 Structure

```
key   → value
"a"   → 1
"b"   → 2
```

---

## 🔥 Important

* Keys must be **unique**
* Values can repeat

---

# 🚀 2. Creating Maps

---

## 🔹 Literal

```go
m := map[string]int{
  "x": 10,
  "y": 20,
}
```

---

## 🔹 Using `make()` (VERY COMMON)

```go
m := make(map[string]int)
```

---

## ⚠️ Critical Rule

```go
var m map[string]int
m["a"] = 1 // ❌ panic
```

👉 Because map is **nil**

---

✔ Fix:

```go
m := make(map[string]int)
m["a"] = 1
```

---

# 🔑 3. Accessing Values

```go
value := m["a"]
```

---

## ⚠️ Important behavior

If key doesn’t exist:

```go
fmt.Println(m["z"]) // 0 (zero value)
```

---

## ✅ Safe access (VERY IMPORTANT)

```go
value, ok := m["a"]
```

---

### Meaning:

* `value` → actual value
* `ok` → true if key exists

---

### Example:

```go
if val, ok := m["a"]; ok {
  fmt.Println("Found:", val)
} else {
  fmt.Println("Not found")
}
```

---

# 🔁 4. Updating & Adding

Same syntax:

```go
m["a"] = 100 // update
m["c"] = 300 // add
```

---

# ❌ 5. Deleting

```go
delete(m, "a")
```

---

👉 No error if key doesn’t exist

---

# 🔁 6. Looping Over Maps

---

## 🔹 Using `range`

```go
for key, value := range m {
  fmt.Println(key, value)
}
```

---

## ⚠️ Important

👉 Order is **NOT guaranteed**

Every iteration may be different.

---

## 🔹 Only keys

```go
for key := range m {
  fmt.Println(key)
}
```

---

---

# 🧠 7. Map is a Reference Type

This is where many bugs happen.

---

```go
a := map[string]int{"x": 1}
b := a

b["x"] = 100

fmt.Println(a) // affected!
```

---

👉 Both point to same underlying data

---

# 🔥 8. Copying Maps (Manual)

No built-in deep copy.

---

## Correct way:

```go
a := map[string]int{"x": 1}

b := make(map[string]int)

for k, v := range a {
  b[k] = v
}
```

---

👉 Now independent

---

# 🔑 9. Key Types (VERY IMPORTANT)

---

## ✅ Allowed

Keys must be **comparable types**:

* int
* string
* bool
* structs (if fields are comparable)

---

## ❌ Not allowed

```go
map[[]int]int // ❌ slices not allowed
```

👉 Because slices are not comparable

---

---

# 🔥 10. Nested Maps

```go
m := map[string]map[string]int{
  "user1": {"score": 10},
}
```

---

## Access:

```go
m["user1"]["score"]
```

---

---

# 🧠 11. Zero Value of Map

```go
var m map[string]int
```

👉 `m == nil`

---

### Behavior:

| Operation | Result  |
| --------- | ------- |
| Read      | ✅ works |
| Write     | ❌ panic |

---

---

# 🔥 12. Common Patterns (Real Usage)

---

## 🔹 Counting frequency

```go
nums := []int{1,2,2,3}

freq := make(map[int]int)

for _, v := range nums {
  freq[v]++
}
```

---

---

## 🔹 Grouping data

```go
words := []string{"go", "java", "go"}

groups := make(map[string][]string)

for _, w := range words {
  groups[w] = append(groups[w], w)
}
```

---

---

## 🔹 Set implementation

```go
set := make(map[string]bool)

set["a"] = true
```

---

👉 Go doesn’t have built-in set → we use maps

---

---

# ⚠️ 13. Common Mistakes (Fix these early)

---

## ❌ Writing to nil map

Biggest beginner error.

---

## ❌ Assuming order

```go
for k := range m {
  // order is random
}
```

---

## ❌ Forgetting `ok` check

```go
val := m["missing"] // might mislead
```

---

## ❌ Expecting deep copy

```go
b := a // ❌ not copy
```

---

---

# 🔒 14. Maps and Concurrency (IMPORTANT)

Maps are:

> ❌ NOT safe for concurrent writes

---

## ❌ Dangerous:

```go
go func() {
  m["a"] = 1
}()
```

---

## ✔ Solutions:

* `sync.Mutex`
* `sync.Map` (special case)

---

---

# 🧠 Mental Model (Lock this in)

Think like this:

* Map = fast lookup table
* Key → index
* Value → stored data

---

And most important:

> A map is NOT the data—it’s a **reference to data**

---

# ⚔️ Map vs Slice (Quick Insight)

| Feature  | Slice     | Map            |
| -------- | --------- | -------------- |
| Access   | index     | key            |
| Order    | preserved | not guaranteed |
| Use case | list      | lookup         |

---

# 🚀 Backend Reality (Where we use maps)

We’ll use maps constantly for:

* JSON decoding
* Caching
* Lookup tables
* Counting/grouping
* Query results

---

# 🔥 Final Advice

If we misuse maps:
👉 Bugs will be subtle and painful

If we understand them:
👉 We get **O(1) lookups + clean logic**

---

PONTERS🧠 — this is one of the most important concepts in Go. If we truly understand pointers, we stop guessing about memory and start writing **predictable, efficient code**.

Let’s go deep, but keep it grounded.

---

# 🧠 1. What is a Pointer?

A **pointer** is:

> A variable that stores the **memory address** of another variable

---

## 🔹 Basic Example

```go
x := 10
p := &x
```

---

## 👉 What’s happening?

* `x` → stores value `10`
* `&x` → gives address of `x`
* `p` → stores that address

---

## 🔍 Visual

```
x = 10        (stored at address 0x123)

p = 0x123     (points to x)
```

---

# 🔥 2. Dereferencing (`*`)

To access the value via pointer:

```go
fmt.Println(*p) // 10
```

---

## 👉 Meaning

* `*p` → “go to address and get value”

---

## 🔁 Modify via pointer

```go
*p = 20
fmt.Println(x) // 20
```

---

👉 This is the real power: **modify original data**

---

# ⚔️ 3. Why Do We Need Pointers?

Let’s be honest—this is the real question.

---

## ❌ Without pointer (copy happens)

```go
func update(x int) {
  x = 100
}

a := 10
update(a)

fmt.Println(a) // still 10
```

---

## ✅ With pointer

```go
func update(x *int) {
  *x = 100
}

a := 10
update(&a)

fmt.Println(a) // 100
```

---

👉 Now we modify original memory

---

# 🧠 4. Pass by Value vs Reference (Critical)

Go is:

> **Always pass-by-value**

Even pointers are passed by value (address is copied)

---

## Example:

```go
func change(p *int) {
  *p = 50
}
```

👉 We pass a copy of the address, but it still points to same data

---

# 🔥 5. Pointer Declaration

---

## 🔹 Basic

```go
var p *int
```

👉 `p` can point to an int

---

## 🔹 Assigning

```go
x := 10
p := &x
```

---

---

# ⚠️ 6. Nil Pointer (IMPORTANT)

```go
var p *int
```

👉 Default value = `nil`

---

## ❌ Dangerous:

```go
*p = 10 // panic
```

---

## ✔ Safe check:

```go
if p != nil {
  *p = 10
}
```

---

---

# 🔥 7. Pointers with Structs (VERY IMPORTANT)

This is where pointers are heavily used.

---

## Without pointer

```go
type User struct {
  Name string
}

func update(u User) {
  u.Name = "Updated"
}
```

👉 Copy is modified, original unchanged

---

## With pointer

```go
func update(u *User) {
  u.Name = "Updated"
}
```

---

### Usage:

```go
user := User{Name: "Skyy"}
update(&user)
```

---

👉 Now original is updated

---

## 🔥 Go Shortcut (Important)

```go
u := &User{Name: "Skyy"}
u.Name = "New"
```

👉 Go automatically dereferences

(No need for `(*u).Name`)

---

---

# 🔁 8. Pointers with Slices & Maps

This is subtle and important.

---

## ❗ Slices & Maps are already reference types

```go
func modify(nums []int) {
  nums[0] = 100
}
```

👉 Works without pointer

---

## When do we use pointer then?

* When modifying the **slice itself** (not just elements)

```go
func add(nums *[]int) {
  *nums = append(*nums, 4)
}
```

---

---

# 🔥 9. Pointer vs Value (When to use what?)

---

## ✅ Use pointers when:

* We want to modify original data
* Data is large (avoid copying)
* Working with structs in backend

---

## ❌ Avoid pointers when:

* Data is small (int, bool)
* No mutation needed

---

---

# ⚠️ 10. Common Mistakes (Fix these early)

---

## ❌ Dereferencing nil pointer

```go
var p *int
fmt.Println(*p) // ❌ panic
```

---

## ❌ Confusion with `*` and `&`

* `&` → get address
* `*` → get value

---

## ❌ Overusing pointers

```go
func add(a *int, b *int) *int // ❌ unnecessary
```

👉 Makes code messy

---

---

# 🔥 11. Pointer to Pointer (Advanced)

```go
x := 10
p := &x
pp := &p
```

---

👉 Rarely used, but exists

---

---

# 🧠 12. Memory Insight (Important)

When we use pointer:

```go
p := &x
```

We are:

* Avoiding copying
* Sharing memory

---

👉 This is critical in:

* APIs
* DB models
* performance-sensitive code

---

---

# 🚀 Real Backend Usage

We’ll see pointers everywhere in Go backend:

---

## 🔹 JSON decoding

```go
json.Unmarshal(data, &user)
```

---

## 🔹 DB operations

```go
func GetUser(id int) (*User, error)
```

---

## 🔹 Optional fields

```go
type User struct {
  Name *string
}
```

👉 nil means “not provided”

---

---

# 🧠 Mental Model (Lock this in)

Think like this:

* Variable → holds value
* Pointer → holds address
* `&` → where is it?
* `*` → what is inside?

---

# ⚔️ Final Insight (Very Important)

> Pointers are about **control over memory and mutation**

If we avoid them:
👉 We lose efficiency and flexibility

If we overuse them:
👉 Code becomes confusing

---

# CODE-FILES 💻
```go
package main

import "fmt"

// 📂 1️⃣. ARRAYS[]
func main() {
	var nums [2]int
	fmt.Printf("%+v\n",nums)

	nums[0]=1
	nums[1]=2

	fmt.Printf("%+v\n",nums)

	primes:=[4]int{2,3,4,5}
	fmt.Printf("%+v\n",primes)
	primes[3]=11
	fmt.Printf("%+v\n",primes)

	for prime:=range primes{
		fmt.Println(prime)
	}

	var matrix [2][3]int
	fmt.Printf("%+v\n",matrix)
	
}

// $ go run main.go
// [0 0]
// [1 2]
// [2 3 4 5]
// [2 3 4 11]
// 0
// 1
// 2
// 3
// [[0 0 0] [0 0 0]]
```
```go
package main

import "fmt"

// 📂 03_complex_data_types
// 📂 2️⃣. SLICES (DYNAMIC-ARRAYS[])

func main() {
	pokemons:=[]string{"Pikachu⚡","Charizard🔥","Umbreon🌛","Noctowl🔮"}

	// ALT. WAY
	nums:=make([]int,3,5)
	fmt.Printf("%+v\n",pokemons)
	fmt.Printf("%+v\n",nums)
	fmt.Printf("Items: %+v, Len: %d, Cap:%d\n",nums, len(nums),cap(nums))

	nums=append(nums, 1)
	nums=append(nums, 2)
	nums=append(nums, 3)
	nums=append(nums, 4)
	nums=append(nums, 5)

	fmt.Printf("Items: %+v, Len: %d, Cap:%d\n",nums, len(nums),cap(nums))
	fmt.Printf("%+v",nums[3:6]) // 3- included, 6- excluded


}

// $ go run main.go
// [Pikachu⚡ Charizard🔥 Umbreon🌛 Noctowl🔮]
// [0 0 0]
// Items: [0 0 0], Len: 3, Cap:5
// Items: [0 0 0 1 2 3 4 5], Len: 8, Cap:10
// [1 2 3]
```

```go
package main

import "fmt"

// 📂 03_complex_data_types
// 📂 3️⃣. MAPS

func main() {
	pokemons:=map[string]int{
		"Pikachu⚡":25,
		"Charizard🔥":6,
		"Squirtle💦":7,
		"Bulbasaur🍃":1,

	}
	fmt.Printf("%+v\n",pokemons)
	pokemons["Charizard🔥"]=06
	pokemons["Squirtle💦"]=7+7
	fmt.Printf("%+v\n",pokemons)

	// check if property exists or not - ok check (bool)
	pikachu,ok:=pokemons["Pikachu⚡"]
	fmt.Println(pikachu,ok)

	if ok{
		fmt.Printf("PIKACHU: %+v\n",pikachu)
	}

	if _,ok=pokemons["Pdgey🐦"];ok{
		fmt.Println("Exists")
	}else{
		fmt.Println("Does not exist")
	}

	//! DELETION
	delete(pokemons, "Pikachu⚡")
	fmt.Printf("%+v\n",pokemons)
	
	// ALT WAY.
	configs:=make(map[string]int)
	fmt.Printf("%+v %T\n",configs,configs)

	// if configs == nil{
	// fmt.Println("Config is nil!")
	// }
}

// $ go run main.go
// map[Bulbasaur🍃:1 Charizard🔥:6 Pikachu⚡:25 Squirtle💦:7]
// map[Bulbasaur🍃:1 Charizard🔥:6 Pikachu⚡:25 Squirtle💦:14]
// 25 true
// PIKACHU: 25
// Does not exist
// map[Bulbasaur🍃:1 Charizard🔥:6 Squirtle💦:14]
// map[] map[string]int
```
```go
package main

import "fmt"

// 📂 03_complex_data_types
// 📂 4️⃣. Ptrs*

func modifyVal(val int){
	val=val*10
	fmt.Printf("modified value: %+v\n",val)
}

func modifyPtr(val *int){
	if val==nil{
		fmt.Println("val is nil!")
		return
	}else{
		*val=*val * 10 // dereferencing
		fmt.Printf("modified ptr: %+v\n",val)
	}
}

func main() {
	age:=30
	agePtr:=&age
	fmt.Printf("age: %d\n",age)
	fmt.Printf("agePtr: %d\n",agePtr)
	fmt.Printf("agePtr-2: %d\n",&age)
	fmt.Printf("agePtr-3: %d\n",&agePtr)

	num:=10
	modifyVal(num)
	fmt.Println("Original Num:",num)
	modifyPtr(&num)
	fmt.Println("Original Num:",num)
}

// $ go run main.go
// age: 30
// agePtr: 824634294440
// agePtr-2: 824634294440
// agePtr-3: 824634114144
// modified value: 100
// Original Num: 10
// modified ptr: 0xc00008c0c8
// Original Num: 100
```
```go
package main

import (
	"fmt"
	"slices"
)

// 📂 03_complex_data_types
// 📂 5️⃣. More on slices


func main() {
	fmt.Println("------- .ADVANCED SLICING OPERATIONS. -------")
	originalSlice:=[]int{0,1,2,3,4,5,6,7,8,9}
	fmt.Printf("Original[]: %+v, Len: %d, Cap:%d\n",originalSlice, len(originalSlice),cap(originalSlice))

	s1:=originalSlice[2:5]
	fmt.Printf("s1: %+v, Len: %d, Cap:%d\n",s1, len(s1),cap(s1))
	s2:=originalSlice[:4]
	fmt.Printf("s2: %+v, Len: %d, Cap:%d\n",s2, len(s2),cap(s2))
	s3:=originalSlice[2:8]
	fmt.Printf("s3: %+v, Len: %d, Cap:%d\n",s3, len(s3),cap(s3))

	fmt.Println(slices.Contains(s3,6))
	//slices.Insert(s2, 1000)
	s2=append(s2, 1000)
	fmt.Printf("modified-s2: %+v, Len: %d, Cap:%d\n",s2, len(s2),cap(s2))
		
}

// $ go run main.go
// ------- .ADVANCED SLICING OPERATIONS. -------
// Original[]: [0 1 2 3 4 5 6 7 8 9], Len: 10, Cap:10
// s1: [2 3 4], Len: 3, Cap:8
// s2: [0 1 2 3], Len: 4, Cap:10
// s3: [2 3 4 5 6 7], Len: 6, Cap:8
// true
// modified-s2: [0 1 2 3 1000], Len: 5, Cap:10
```