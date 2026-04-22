**structs + receivers + methods** are *core Go*, and once we get them, a lot of Go suddenly “clicks”.

---

# 🧱 1. What is a Struct in Go?

A **struct** is Go’s way of grouping related data together.

Think of it like:

> “We want one object that holds multiple fields.”

```go
type Person struct {
    Name string
    Age  int
}
```

Now we can create values:

```go
p1 := Person{
    Name: "Skyy",
    Age:  29,
}
```

### Key idea:

Struct = **custom data type with fields**

---

# 🧠 2. Why Structs Matter in Go

Go does NOT have traditional classes like Java/C++.

Instead:

* Struct = data
* Methods = behavior
* Together → behave like a class

So Go’s “OOP” is:

> **Struct + Methods = Object-like behavior**

---

# ⚙️ 3. Methods in Go (Core Concept)

In Go, functions can be attached to types → these are called **methods**.

### Syntax:

```go
func (receiver Type) methodName() {
    // logic
}
```

Example:

```go
func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.Name)
}
```

Usage:

```go
p1.Greet()
```

---

# 🎯 4. What is a Receiver?

The receiver is the **thing the method operates on**.

```go
func (p Person) Greet()
```

Here:

* `p` → receiver variable
* `Person` → receiver type

👉 This means:

> “This method belongs to Person”

---

# 🧪 5. Value Receiver vs Pointer Receiver (VERY IMPORTANT)

This is where most beginners struggle. Let’s get it crystal clear.

---

## ✅ Value Receiver

```go
func (p Person) UpdateName(newName string) {
    p.Name = newName
}
```

### Problem:

```go
p1 := Person{Name: "Skyy"}
p1.UpdateName("John")

fmt.Println(p1.Name) // STILL "Skyy"
```

### Why?

Because:

> Value receiver = **copy of struct is passed**

So we modify the copy, not original.

---

## ✅ Pointer Receiver

```go
func (p *Person) UpdateName(newName string) {
    p.Name = newName
}
```

Now:

```go
p1 := Person{Name: "Skyy"}
p1.UpdateName("John")

fmt.Println(p1.Name) // "John"
```

### Why this works:

> Pointer receiver = method gets reference to original struct

---

# 🧠 Rule of Thumb (Remember This)

Use pointer receivers when:

* We want to **modify data**
* Struct is **large** (avoid copying)
* We want **consistent behavior**

---

# 🔄 6. Go’s Smart Behavior (Important Detail)

Even if we use pointer receiver:

```go
p1.UpdateName("John")
```

Go automatically converts:

```go
(&p1).UpdateName("John")
```

So we don’t need to manually use `&`.

---

# 🧩 7. Multiple Methods on Struct

```go
type Rectangle struct {
    Width  float64
    Height float64
}
```

### Method 1:

```go
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

### Method 2:

```go
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}
```

Usage:

```go
rect := Rectangle{Width: 10, Height: 5}

fmt.Println(rect.Area())      // 50
fmt.Println(rect.Perimeter()) // 30
```

---

# 🔥 8. Method Sets (Advanced but Important)

This is subtle but powerful.

### If method has:

#### Value receiver:

* Can be called on:

  * value
  * pointer

#### Pointer receiver:

* Can be called on:

  * pointer
  * value (Go auto converts)

---

### Example:

```go
func (p Person) A() {}
func (p *Person) B() {}
```

Then:

```go
p := Person{}
ptr := &p

p.A()    // ✅
p.B()    // ✅ (auto converts)

ptr.A()  // ✅
ptr.B()  // ✅
```

---

# 🧱 9. Struct Embedding (Mini Inheritance)

Go doesn’t have inheritance, but we simulate it.

```go
type Animal struct {
    Name string
}

type Dog struct {
    Animal
    Breed string
}
```

Now:

```go
d := Dog{
    Animal: Animal{Name: "Buddy"},
    Breed:  "Labrador",
}

fmt.Println(d.Name) // Buddy
```

---

# 🧠 10. Methods with Embedded Structs

```go
func (a Animal) Speak() {
    fmt.Println("Animal speaks")
}
```

Now:

```go
d.Speak() // works!
```

👉 This is called **method promotion**

---

# ⚠️ 11. Common Mistakes (Be Careful)

### ❌ Mixing receivers

```go
func (p Person) A() {}
func (p *Person) B() {}
```

This works, but:

> It creates confusion → avoid mixing unless necessary

---

### ❌ Forgetting pointer when modifying

If something isn’t updating:

> 90% chance we forgot pointer receiver

---

# 🚀 12. Real-World Example (Closer to Backend Work)

```go
type User struct {
    Name  string
    Email string
}
```

### Method:

```go
func (u *User) UpdateEmail(newEmail string) {
    u.Email = newEmail
}
```

### Usage:

```go
user := User{Name: "Skyy", Email: "old@mail.com"}
user.UpdateEmail("new@mail.com")
```

---

# 🧩 Final Mental Model

Think like this:

| Concept          | Meaning                 |
| ---------------- | ----------------------- |
| Struct           | Data container          |
| Method           | Function tied to struct |
| Receiver         | The struct instance     |
| Pointer receiver | Modify original         |
| Value receiver   | Work on copy            |

---

Interface - This is the piece that makes Go feel *very different* from languages like Java or C++. If we don’t internalize interfaces properly, Go will always feel a bit “loose” and confusing.

Let’s build this from the ground up.

---

# 🧠 1. What is an Interface in Go?

An **interface** is a type that defines a set of method signatures.

```go
type Shape interface {
    Area() float64
}
```

👉 This does **NOT** contain implementation.
It just says:

> “Any type that has an `Area()` method is a Shape.”

---

# 🔥 2. The Big Idea (This is where Go is unique)

In Go:

> **Types implement interfaces implicitly**

There is NO:

* `implements` keyword ❌
* explicit declaration ❌

If a type has the required methods → it automatically satisfies the interface.

---

# 🧱 3. Example: Interface in Action

```go
type Shape interface {
    Area() float64
}
```

### Struct 1:

```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

### Struct 2:

```go
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}
```

👉 Both `Rectangle` and `Circle` now satisfy `Shape`.

---

# ⚙️ 4. Using Interfaces

```go
func printArea(s Shape) {
    fmt.Println(s.Area())
}
```

Usage:

```go
r := Rectangle{10, 5}
c := Circle{7}

printArea(r)
printArea(c)
```

👉 Same function, different types → **polymorphism**

---

# 🧠 5. Mental Model

Think like this:

| Thing     | Meaning                                            |
| --------- | -------------------------------------------------- |
| Struct    | "I am data + methods"                              |
| Interface | "I don’t care what you are, just behave like this" |

---

# ⚠️ 6. Pointer vs Value with Interfaces (CRITICAL)

This is where people mess up.

---

## Case 1: Value receiver

```go
func (r Rectangle) Area() float64
```

Then BOTH work:

```go
var s Shape

r := Rectangle{10, 5}
s = r      // ✅
s = &r     // ✅
```

---

## Case 2: Pointer receiver

```go
func (r *Rectangle) Area() float64
```

Now:

```go
s = r      // ❌ ERROR
s = &r     // ✅
```

👉 Why?

Because:

> Method belongs to `*Rectangle`, not `Rectangle`

---

### 🔥 Rule:

> Interface satisfaction depends on **method set**

---

# 🧩 7. Interface as a Type (Very Important)

Interfaces are actual types.

```go
var s Shape
```

This variable can hold:

* Rectangle
* Circle
* ANY type that implements `Area()`

---

# 📦 8. Empty Interface (`interface{}` / `any`)

```go
var x interface{}
```

or (modern Go):

```go
var x any
```

👉 This can hold **anything**

```go
x = 10
x = "hello"
x = true
```

---

### ⚠️ But here’s the catch:

We lose type safety.

---

# 🔍 9. Type Assertion

Used to extract actual value from interface.

```go
value := x.(int)
```

Safe version:

```go
value, ok := x.(int)

if ok {
    fmt.Println("It's an int:", value)
}
```

---

# 🔀 10. Type Switch

Cleaner way to handle multiple types:

```go
switch v := x.(type) {
case int:
    fmt.Println("int:", v)
case string:
    fmt.Println("string:", v)
default:
    fmt.Println("unknown type")
}
```

---

# 🧱 11. Multiple Methods in Interface

```go
type Animal interface {
    Speak()
    Move()
}
```

To satisfy this:

```go
type Dog struct{}

func (d Dog) Speak() {}
func (d Dog) Move() {}
```

👉 Must implement ALL methods.

---

# 🔗 12. Interface Composition

We can combine interfaces:

```go
type Reader interface {
    Read()
}

type Writer interface {
    Write()
}

type ReadWriter interface {
    Reader
    Writer
}
```

👉 This is very common in Go.

---

# 🧠 13. Real Go Philosophy (Important Insight)

In Go:

> Interfaces are usually **small and focused**

Example from standard library:

* `io.Reader` → only `Read()`
* `io.Writer` → only `Write()`

---

### Why?

Because:

> Small interfaces = flexible code

---

# ⚡ 14. Dependency Inversion (Real Backend Use)

Instead of:

```go
type MySQL struct{}
```

We do:

```go
type Database interface {
    Save(data string)
}
```

Then:

```go
func process(db Database) {
    db.Save("data")
}
```

👉 Now we can swap:

* MySQL
* MongoDB
* Mock (for testing)

---

# 🧪 15. Interface + Struct Together

```go
type PaymentProcessor interface {
    Pay(amount float64)
}
```

Implementations:

```go
type Stripe struct{}
func (s Stripe) Pay(a float64) {}

type Razorpay struct{}
func (r Razorpay) Pay(a float64) {}
```

Usage:

```go
func checkout(p PaymentProcessor) {
    p.Pay(100)
}
```

---

# 🚨 16. Nil Interface Trap (Very Important)

```go
var s Shape
fmt.Println(s == nil) // true
```

BUT:

```go
var r *Rectangle = nil
s = r

fmt.Println(s == nil) // ❌ false
```

👉 Why?

Because interface = (type + value)

Here:

* type = *Rectangle
* value = nil

So interface ≠ nil

---

# 🧩 17. When SHOULD We Use Interfaces?

Use interfaces when:

* We want **flexibility**
* We want **decoupling**
* We want **testability (mocking)**

---

# ❌ When NOT to Use

Don’t create interfaces:

* “just in case”
* without real need

👉 This is a common beginner mistake.

---

# 🧠 Final Mental Model (Lock This In)

> **Interfaces describe behavior, not data**

And the most important line:

> “If it walks like a duck and quacks like a duck, Go treats it as a duck.”

---

 **`Stringer` is small, but it teaches how Go *really* uses interfaces in practice**.

---

# 🧠 1. What is the Stringer Interface?

In Go, there’s a built-in interface:

```go
type Stringer interface {
    String() string
}
```

👉 It lives inside the fmt package.

---

# 🔥 Core Idea

If a type implements:

```go
String() string
```

Then Go will automatically use that method when printing.

---

# ⚙️ 2. Default Behavior Without Stringer

```go
type Person struct {
    Name string
    Age  int
}

p := Person{"Skyy", 29}
fmt.Println(p)
```

### Output:

```
{Skyy 29}
```

👉 This is Go’s default struct printing.

---

# ✨ 3. Implementing Stringer

Now we define our own representation:

```go
func (p Person) String() string {
    return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}
```

Now:

```go
fmt.Println(p)
```

### Output:

```
Name: Skyy, Age: 29
```

---

# 🧠 What just happened?

Because `Person` implements:

```go
String() string
```

It automatically satisfies `Stringer`.

👉 No explicit declaration needed.

---

# 🔍 4. Where Stringer is Used

Anywhere in Go that prints values:

* `fmt.Println()`
* `fmt.Printf("%v", x)`
* `fmt.Sprint()`

They internally check:

> “Does this value implement Stringer?”

If yes → call `String()`

---

# ⚙️ 5. Under the Hood (Important)

When we do:

```go
fmt.Println(p)
```

Internally:

1. Check: does `p` implement `Stringer`?
2. If yes → call `p.String()`
3. If no → use default formatting

---

# 🧩 6. Example with Multiple Types

```go
type Car struct {
    Brand string
}

func (c Car) String() string {
    return "Car brand is " + c.Brand
}
```

```go
type Book struct {
    Title string
}

func (b Book) String() string {
    return "Book: " + b.Title
}
```

Now:

```go
fmt.Println(Car{"BMW"})
fmt.Println(Book{"Go in Action"})
```

👉 Each prints its own custom format.

---

# ⚠️ 7. Pointer vs Value Receiver (Important Again)

### Case 1: Value receiver

```go
func (p Person) String() string
```

Works for:

```go
p := Person{}
fmt.Println(p)   // ✅
fmt.Println(&p)  // ✅
```

---

### Case 2: Pointer receiver

```go
func (p *Person) String() string
```

Now:

```go
p := Person{}

fmt.Println(p)   // ❌ might NOT use String()
fmt.Println(&p)  // ✅
```

👉 Why?

Because:

> Only `*Person` implements Stringer, not `Person`

---

### 🔥 Rule:

If unsure → use **value receiver for String()**

---

# 🧪 8. Real Use Case

Let’s say we’re building logs:

```go
type Order struct {
    ID     int
    Amount float64
}
```

Without Stringer:

```
{101 250.5}
```

With Stringer:

```go
func (o Order) String() string {
    return fmt.Sprintf("Order #%d -> ₹%.2f", o.ID, o.Amount)
}
```

Now logs become readable:

```
Order #101 -> ₹250.50
```

👉 This is where it shines.

---

# 🧠 9. Stringer + Interfaces Together

We can even use it explicitly:

```go
func printString(s fmt.Stringer) {
    fmt.Println(s.String())
}
```

Now:

```go
printString(Person{"Skyy", 29})
```

👉 Any type implementing `String()` works here.

---

# ⚠️ 10. Common Mistakes

### ❌ Forgetting exact method signature

```go
func (p Person) ToString() string // ❌ WRONG
```

Must be:

```go
String() string // exact match
```

---

### ❌ Infinite recursion (VERY COMMON BUG)

```go
func (p Person) String() string {
    return fmt.Sprintf("%v", p) // ❌ infinite loop
}
```

👉 Why?

Because:

* `%v` → calls `String()`
* which calls `%v` again → loop forever 💥

---

### ✅ Fix:

```go
func (p Person) String() string {
    return fmt.Sprintf("Name: %s", p.Name)
}
```

---

# 🧩 11. Stringer vs JSON

Important distinction:

* `String()` → for **printing / debugging**
* JSON → for **API responses**

```go
json.Marshal(p)
```

👉 Does NOT use `String()`

---

# 🧠 12. Why Stringer Exists (Philosophy)

Go prefers:

> “Let types describe themselves”

Instead of external formatting logic.

---

# 🚀 Final Mental Model

| Concept  | Meaning                      |
| -------- | ---------------------------- |
| Stringer | Interface with `String()`    |
| Purpose  | Custom string representation |
| Used by  | fmt package                  |
| Benefit  | Clean logs, debugging        |

---

# 💡 Honest Advice

This looks simple—but don’t ignore it.

If we:

* build APIs
* log data
* debug complex structs

👉 `Stringer` becomes *extremely useful*

---

GENERICS - This is one of the biggest upgrades Go has had. If we get **generics** properly, we stop writing repetitive code and start writing *reusable, type-safe abstractions*.

Honestly: generics in Go are **simpler than in C++/Java**, and that’s intentional. Don’t try to overcomplicate them.

---

# 🧠 1. What are Generics?

Generics let us write code that works with **multiple types**, while still keeping **type safety**.

👉 Before generics, we had two bad options:

* Duplicate code ❌
* Use `interface{}` (lose type safety) ❌

Generics fix this.

---

# 🔥 2. Basic Syntax

```go
func functionName[T any](param T) T {
    return param
}
```

### Breakdown:

* `T` → type parameter
* `any` → constraint (means any type allowed)

---

# ⚙️ 3. Simple Example

```go
func Identity[T any](value T) T {
    return value
}
```

Usage:

```go
fmt.Println(Identity)
fmt.Println(Identity[string]("hello"))
```

👉 Same function, different types.

---

# 🧠 Key Insight

> `T` is just a placeholder for a real type.

---

# 📦 4. Why Not interface{}?

Old way:

```go
func Print(value interface{}) {
    fmt.Println(value)
}
```

Problem:

* No type safety
* Need type assertions

Generics:

```go
func Print[T any](value T) {
    fmt.Println(value)
}
```

👉 Safer + cleaner

---

# 🧩 5. Generics with Slices

```go
func PrintSlice[T any](items []T) {
    for _, item := range items {
        fmt.Println(item)
    }
}
```

Usage:

```go
PrintSlice([]int{1, 2, 3})
PrintSlice([]string{"a", "b"})
```

---

# ⚠️ 6. Constraints (VERY IMPORTANT)

We often don’t want “any type”—we want **restricted types**.

---

## Example: Add two numbers

```go
func Add[T int | float64](a, b T) T {
    return a + b
}
```

👉 Now only `int` or `float64` allowed.

---

# 🧠 7. Constraint Interfaces (Powerful Concept)

We can define reusable constraints:

```go
type Number interface {
    int | float64
}
```

Then:

```go
func Add[T Number](a, b T) T {
    return a + b
}
```

---

# 🔥 8. Predefined Constraint: comparable

Go gives us built-in constraints.

```go
func Contains[T comparable](slice []T, val T) bool {
    for _, v := range slice {
        if v == val {
            return true
        }
    }
    return false
}
```

👉 `comparable` means:

* `==` and `!=` must work

---

# 🧱 9. Generic Structs

Generics aren’t just for functions.

```go
type Box[T any] struct {
    Value T
}
```

Usage:

```go
b1 := Box[int]{Value: 10}
b2 := Box[string]{Value: "hello"}
```

---

# ⚙️ 10. Methods on Generic Types

```go
func (b Box[T]) GetValue() T {
    return b.Value
}
```

👉 Works just like normal methods.

---

# 🧠 11. Type Inference (Important Convenience)

We don’t always need to specify type:

```go
Identity(10)        // Go infers T = int
Identity("hello")   // T = string
```

👉 Cleaner code.

---

# 🔍 12. Real Example (Useful One)

### Find max value:

```go
func Max[T int | float64](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

---

# ⚠️ 13. Limitations of Generics in Go

Go keeps generics **simple and limited** on purpose.

---

### ❌ No operator overloading

We must explicitly define constraints:

```go
int | float64
```

---

### ❌ No "any method" calls

We can’t do:

```go
func DoSomething[T any](v T) {
    v.SomeMethod() // ❌ not allowed
}
```

👉 Unless constraint defines that method.

---

# 🧩 14. Combining Generics + Interfaces

```go
type Speaker interface {
    Speak()
}
```

```go
func MakeSpeak[T Speaker](obj T) {
    obj.Speak()
}
```

👉 Now only types with `Speak()` allowed.

---

# 🧠 15. When SHOULD We Use Generics?

Use generics when:

* Same logic works for multiple types
* We want **type safety**
* We want **reuse without duplication**

---

# ❌ When NOT to Use

Don’t force generics:

* If code is simple
* If only one type is used
* If interface is cleaner

👉 This is where many beginners go wrong.

---

# ⚡ 16. Generics vs Interfaces (Important Comparison)

| Feature     | Generics                    | Interfaces           |
| ----------- | --------------------------- | -------------------- |
| Purpose     | Code reuse                  | Behavior abstraction |
| Type safety | Compile-time                | Runtime              |
| Best for    | Algorithms, data structures | Polymorphism         |

---

# 🧠 Real Insight

* Generics = “same logic, different types”
* Interfaces = “different types, same behavior”

---

# 🚀 17. Real Backend Use Cases

Generics are useful for:

* Utility functions (map, filter, reduce)
* Data structures (stack, queue)
* Reusable libraries

---

# 🧩 Example: Generic Stack

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
    last := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return last
}
```

---

# ⚠️ 18. Common Mistakes

### ❌ Overusing generics

Not everything needs generics.

---

### ❌ Ignoring constraints

Leads to compile errors.

---

### ❌ Trying to write Java-style generics

Go is simpler—don’t fight it.

---

# 🧠 Final Mental Model

> Generics = **write once, work with many types safely**

---

# 💡 Straight Talk

If we’re coming from JavaScript/TypeScript:

* Generics feel familiar ✅
* But Go’s version is stricter and simpler ⚠️

Don’t try to make it “too clever”.

---


# CODE-FILES 💻
```go
package main

// 📂 05_oop

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

func NewEmployee(id int, firstName,lastName,position string, salary int, isActive bool) Employee{
	return Employee{
		ID: id,
		FirstName: firstName,
		LastName: lastName,
		Position: position,
		Salary: salary,
		JoinedAt: time.Now(),
	}
}

func main() {

	skyy := Employee{
		ID: 1,
		FirstName: "Skyy",
		LastName: "Banerjee",
		Position: "Frontend Dev.",
		Salary: 55000,
		IsActive: true,
		JoinedAt: time.Now(),
	}

	fmt.Println("skyy:",skyy)
	fmt.Println("salary:",skyy.Salary)


	soumadip := NewEmployee(2,"Soumadip","Banerjee","golang-dev",59000,true)

	fmt.Println("salary:",soumadip.Salary)
	soumadiPtr:=&soumadip

	soumadiPtr.Salary=50000000
	fmt.Println("updated_salary:",soumadip.Salary)
	fmt.Println("soumadip:",soumadip)


// $ go run main.go
// skyy: {1 Skyy Banerjee Frontend Dev. 55000 true 2026-04-22 15:14:12.2144169 +0530 IST m=+0.000692001}
// salary: 55000
// salary: 59000
// updated_salary: 50000000
// soumadip: {2 Soumadip Banerjee golang-dev 50000000 false 2026-04-22 15:14:12.2155256 +0530 IST m=+0.001800701}

}
```
```go
package main

// 📂 05_oop
// Methods and receivers

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

// Attaching the methods() to the struct
// VALUE RECEIVER
func (e Employee) FullName() string{
	return e.FirstName+" "+e.LastName
}

// func (e Employee) Deactivate(){
// 	e.IsActive=false
// }

// 💡 PTR. RECEIVER -> Whenever we manipulate/change the original struct{}
func (e *Employee) DeactivateWPtr(){
	e.IsActive=false
}




func main() {
	skyy := Employee{
		ID: 1,
		FirstName: "Skyy",
		LastName: "Banerjee",
		Position: "Frontend Dev.",
		Salary: 55000,
		IsActive: true,
		JoinedAt: time.Now(),
	}

	fmt.Println("skyy:",skyy)
	fmt.Println("full-name:",skyy.FullName())
	fmt.Printf("%+v\n",skyy.IsActive)
	skyy.DeactivateWPtr() // works now!
	fmt.Printf("%+v\n",skyy.IsActive)
	fmt.Println("skyy:",skyy)

	// joe:=&Employee{}
	// joe=nil //⚠️Bad Idea!
	// joe.FullName() // Panic!


// $ go run main.go
// skyy: {1 Skyy Banerjee Frontend Dev. 55000 true 2026-04-22 16:20:00.3762711 +0530 IST m=+0.000652701}
// full-name: Skyy Banerjee
// true
// false
// skyy: {1 Skyy Banerjee Frontend Dev. 55000 false 2026-04-22 16:20:00.3762711 +0530 IST m=+0.000652701}
}
```
```go
package main

import (
	"fmt"
)

// 📂 05_oop
// Interface - A contract

type Pokemon interface{
	GetName() string
}

type ElectricPokemon struct {
	ID        int
	Name	string

	
}

type FirePokemon struct {
	ID int
	Name string
}

func (p ElectricPokemon) GetName()string{
	return p.Name
}

func (p FirePokemon) GetName()string{
	return p.Name
}


// Interface implementation
func DisplayPokemon(p Pokemon){
	fmt.Println("Name:",p.GetName())
}



func main() {

	p:=ElectricPokemon{
		ID:25,
		Name: "Pikachu",
	}

	c:=FirePokemon{
		ID:6,
		Name: "Charizard",
	}

	DisplayPokemon(p)
	DisplayPokemon(c)

	// Go native interfaces
	// io.WriteCloser


// $ go run main.go
// Name: Pikachu
// Name: Charizard
}
```

```go 
package main

import "fmt"

// 📂 05_oop
// Generics

func Sum[T int|float32|float64](nums ...T)T{
	var tot T

	for _,i:=range nums{
	 tot+=i
	}
	return tot
}

type Number interface{
	int | float32 | string | float64
}

func Sum2[T Number](nums ...T)T{
	var tot T

	for _,i:=range nums{
	 tot+=i
	}
	return tot
}

func main() {
	grades:=[]int{90,85}
	people:=[]string{"John","Jane","Peter"}

	fmt.Println(len(grades), len(people)) // len - gnereic f(x)

	f:=Sum[float32](1.5,1.2,1.1)

	fmt.Println(Sum(1,2,3,4))
	fmt.Println("Sum:",f)
	fmt.Println(Sum(7.7,88.65,3,4))

	fmt.Println(Sum2("Ha ","ha ", " ha"))


// $ go run main.go
// 2 3
// 10
// Sum: 3.8000002
// 103.35000000000001
// Ha ha  ha
}
```