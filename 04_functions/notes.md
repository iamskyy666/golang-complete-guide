If we understand functions properly in Go, our code becomes **clean, reusable, and predictable**.

---

# 🧠 1. What is a Function?

A **function** is:

> A reusable block of code that performs a task

---

## 🔹 Basic Syntax

```go
func functionName(parameters) returnType {
  // logic
}
```

---

## ✅ Example

```go
func greet(name string) string {
  return "Hello " + name
}
```

---

## Usage:

```go
msg := greet("Skyy")
fmt.Println(msg)
```

---

# 🔥 2. Function Parameters

---

## 🔹 Typed parameters

```go
func add(a int, b int) int {
  return a + b
}
```

---

## 🔹 Short form (same type)

```go
func add(a, b int) int {
  return a + b
}
```

---

## 🧠 Important

Go is:

> **Strictly typed**

No implicit conversion

---

# 🔁 3. Multiple Return Values (VERY IMPORTANT)

This is a **signature feature of Go**.

---

## Example:

```go
func divide(a, b int) (int, error) {
  if b == 0 {
    return 0, fmt.Errorf("cannot divide by zero")
  }
  return a / b, nil
}
```

---

## Usage:

```go
result, err := divide(10, 2)
```

---

## Ignore value:

```go
result, _ := divide(10, 2)
```

---

## 🧠 Why this matters

Go avoids exceptions → uses **explicit error handling**

---

# 🔥 4. Named Return Values

---

## Example:

```go
func add(a, b int) (result int) {
  result = a + b
  return
}
```

---

## ⚠️ Reality

Avoid overusing this—it can reduce clarity.

---

# 🔁 5. Variadic Functions

---

## 👉 Accept multiple arguments

```go
func sum(nums ...int) int {
  total := 0
  for _, v := range nums {
    total += v
  }
  return total
}
```

---

## Usage:

```go
sum(1,2,3,4)
```

---

## Pass slice:

```go
nums := []int{1,2,3}
sum(nums...)
```

---

# 🧠 6. Functions are First-Class Citizens

👉 This is powerful.

---

## 🔹 Assign to variable

```go
add := func(a, b int) int {
  return a + b
}
```

---

## 🔹 Pass as argument

```go
func operate(a, b int, fn func(int, int) int) int {
  return fn(a, b)
}
```

---

## Usage:

```go
result := operate(2, 3, add)
```

---

# 🔥 7. Anonymous Functions

---

## Example:

```go
func() {
  fmt.Println("Hello")
}()
```

👉 Immediately invoked

---

## With variables:

```go
x := func(a int) int {
  return a * 2
}
```

---

# 🔄 8. Closures (IMPORTANT)

---

## Example:

```go
func counter() func() int {
  count := 0

  return func() int {
    count++
    return count
  }
}
```

---

## Usage:

```go
c := counter()
fmt.Println(c()) // 1
fmt.Println(c()) // 2
```

---

## 🧠 Why important?

* Maintains state
* Used in middleware, handlers

---

# ⚔️ 9. Pass by Value vs Pointer (Critical)

---

## Value

```go
func update(x int) {
  x = 100
}
```

👉 No change outside

---

## Pointer

```go
func update(x *int) {
  *x = 100
}
```

👉 Changes original

---

---

# 🔥 10. Methods (Functions on Structs)

---

## Example:

```go
type User struct {
  Name string
}
```

---

## Method:

```go
func (u User) greet() {
  fmt.Println("Hello", u.Name)
}
```

---

## Usage:

```go
u := User{Name: "Skyy"}
u.greet()
```

---

## Pointer receiver

```go
func (u *User) updateName(name string) {
  u.Name = name
}
```

---

👉 Used for modifying struct

---

# 🧠 11. Function Scope

---

## Local scope

```go
func test() {
  x := 10
}
```

👉 `x` only inside function

---

## Package-level

```go
var x = 10
```

---

---

# 🔥 12. init() Function

---

## Example:

```go
func init() {
  fmt.Println("Runs before main")
}
```

---

👉 Used for:

* setup
* config
* initialization

---

---

# ⚠️ 13. Common Mistakes

---

## ❌ Ignoring error returns

```go
result, _ := divide(10, 0) // bad practice
```

---

## ❌ Overusing pointers

Not every function needs pointer params

---

## ❌ Too many responsibilities

Bad:

```go
func doEverything() {}
```

---

👉 Functions should be small and focused

---

# 🧠 Mental Model

Think like this:

* Function = behavior
* Struct = data
* Pointer = control over mutation

---

# 🚀 Backend Perspective

We’ll use functions for:

* API handlers
* DB operations
* Business logic
* Validation

---

# 🔥 Final Insight

Go functions are designed to be:

* Simple
* Explicit
* Predictable
---

Closures are one of those concepts that feel “meh” at first, then suddenly become **very powerful** once we see where they actually matter.

---

# 🧠 1. What is a Closure?

A **closure** is:

> A function that **remembers and uses variables from its outer scope**, even after that outer function has finished executing

---

## 🔥 Simple Example

```go
func outer() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}
```

---

## Usage:

```go
counter := outer()

fmt.Println(counter()) // 1
fmt.Println(counter()) // 2
fmt.Println(counter()) // 3
```

---

## 🧠 What’s happening?

* `count` is defined inside `outer`
* Normally it should disappear after `outer` finishes
* But the inner function **captures it**

👉 That’s the closure

---

# 🔍 2. Why does this work?

Because Go stores:

* The function
* AND the variables it uses

👉 Together as a **closure**

---

## Think of it like:

```text
function + its environment = closure
```

---

# 🔥 3. Real Mental Model

---

## Without closure:

```go
func increment(count int) int {
	count++
	return count
}
```

👉 No memory → resets every time

---

## With closure:

```go
counter := outer()
```

👉 Now:

* State is preserved
* `count` keeps increasing

---

# ⚔️ 4. Closures vs Normal Functions

| Feature | Normal Function | Closure        |
| ------- | --------------- | -------------- |
| Memory  | ❌ No            | ✅ Yes          |
| State   | ❌ Reset         | ✅ Persist      |
| Scope   | Local only      | Captures outer |

---

# 🔁 5. Multiple Closures = Separate State

```go
c1 := outer()
c2 := outer()

fmt.Println(c1()) // 1
fmt.Println(c1()) // 2

fmt.Println(c2()) // 1 (separate state!)
```

---

👉 Each closure has its **own copy of variables**

---

# 🔥 6. Closures with Parameters

```go
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
```

---

## Usage:

```go
double := multiplier(2)
triple := multiplier(3)

fmt.Println(double(5)) // 10
fmt.Println(triple(5)) // 15
```

---

👉 `factor` is captured

---

# 🔥 7. Closures in Loops (VERY IMPORTANT BUG)

This is where many developers mess up.

---

## ❌ Wrong

```go
nums := []int{1,2,3}

for _, v := range nums {
	go func() {
		fmt.Println(v)
	}()
}
```

---

👉 Output might be:

```text
3 3 3
```

---

## 🧠 Why?

Because closure captures **same variable `v`**, not value

---

## ✅ Fix

```go
for _, v := range nums {
	v := v // create new variable

	go func() {
		fmt.Println(v)
	}()
}
```

---

👉 Now correct output

---

# 🔥 8. Closures as Function Factories

We can generate functions dynamically.

---

## Example:

```go
func adder(base int) func(int) int {
	return func(x int) int {
		return base + x
	}
}
```

---

## Usage:

```go
add10 := adder(10)

fmt.Println(add10(5)) // 15
```

---

👉 This is used in:

* Middleware
* Config-based logic

---

# 🔥 9. Closures vs Pointers

Closures:

* Capture variables automatically

Pointers:

* Manually pass references

---

## Example:

```go
func outer() func() {
	x := 10

	return func() {
		fmt.Println(x)
	}
}
```

👉 No pointer needed—closure handles it

---

# 🚀 10. Real Backend Use Cases

This is where closures actually shine.

---

## 🔹 Middleware (very common)

```go
func authMiddleware(role string) func() {
	return func() {
		fmt.Println("Checking role:", role)
	}
}
```

---

---

## 🔹 Lazy initialization

```go
func lazyLoader() func() int {
	data := 0

	return func() int {
		if data == 0 {
			data = 100 // load once
		}
		return data
	}
}
```

---

---

## 🔹 State tracking

* Counters
* Sessions
* Rate limiting

---

# ⚠️ 11. Common Mistakes

---

## ❌ Loop variable capture (biggest one)

---

## ❌ Overusing closures

👉 Makes code harder to read

---

## ❌ Hidden state bugs

Closures hide state → debugging harder

---

# 🧠 Final Mental Model

Think like this:

> A closure is a function that **carries its own memory**

---

# 🔥 One-line intuition

> “Even after the outer function is gone, the inner function still remembers everything it needs.”

---

# 🚀 What matters for us

We should focus on:

* Capturing variables safely
* Using closures for **stateful logic**
* Avoiding loop bugs

---

If we misunderstand `panic` and `recover`, we either:

* crash apps unnecessarily ❌
* or hide bugs we shouldn’t ❌

Let’s get this **clean and practical**.

---

# 🧠 1. What is `panic` in Go?

A **panic** is:

> A runtime error that immediately stops normal execution of the program

---

## 🔥 Example

```go
panic("something went wrong")
```

---

## 👉 What happens?

1. Current function stops
2. Go starts **unwinding the stack**
3. Runs all `defer` statements
4. If not recovered → program crashes

---

## Output looks like:

```
panic: something went wrong

goroutine 1 [running]:
main.main()
...
```

---

# ⚔️ 2. panic vs error (VERY IMPORTANT)

| Feature  | error                      | panic               |
| -------- | -------------------------- | ------------------- |
| Type     | return value               | runtime crash       |
| Handling | explicit (`if err != nil`) | implicit            |
| Use case | expected failures          | unexpected failures |

---

## 🧠 Rule

> ✅ Use `error` for expected problems
> ❌ Use `panic` only for truly broken states

---

# 🔥 3. When should we use `panic`?

---

## ✅ Valid cases

* Programmer mistakes
* Invalid assumptions
* Corrupted state

---

### Example:

```go
func divide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}
```

---

## ❌ Bad usage

```go
if userNotFound {
	panic("user not found") // ❌ should be error
}
```

---

👉 This should be handled, not crash the app

---

# 🔁 4. What is `recover`?

`recover` is:

> A built-in function that catches a panic and prevents the crash

---

## ⚠️ Important

`recover()` only works:

> ✅ inside a `defer` function

---

# 🔥 5. Basic Example

```go
func safe() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()

	panic("boom")
}
```

---

## Output:

```
Recovered from: boom
```

👉 Program continues

---

# 🧠 6. Execution Flow

---

## Without recover:

```text
panic → stack unwind → crash
```

---

## With recover:

```text
panic → defer runs → recover catches → continue
```

---

# 🔥 7. defer + panic interaction

---

## Example:

```go
func test() {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	panic("error")
}
```

---

## Output:

```
Deferred 2
Deferred 1
panic: error
```

---

👉 LIFO order (last defer runs first)

---

# 🔥 8. Real Example (Safe Execution Wrapper)

---

```go
func runSafe(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fn()
}
```

---

## Usage:

```go
runSafe(func() {
	panic("something broke")
})
```

---

👉 This is used in:

* servers
* job workers
* middleware

---

# 🔥 9. Panic in Goroutines (IMPORTANT)

Each goroutine is isolated.

---

## ❌ This will crash program:

```go
go func() {
	panic("oops")
}()
```

---

## ✔ Safe version:

```go
go func() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in goroutine:", r)
		}
	}()
	panic("oops")
}()
```

---

# ⚠️ 10. Common Mistakes

---

## ❌ Using panic for normal errors

Bad practice.

---

## ❌ Recover outside defer

```go
r := recover() // ❌ useless
```

---

## ❌ Swallowing errors silently

```go
recover() // no logging ❌
```

---

👉 Always log or handle properly

---

# 🔥 11. Real Backend Usage

---

## 🔹 HTTP server middleware

In web servers, panic should NOT crash the server.

---

### Example concept:

```go
func middleware(next func()) func() {
	return func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Internal Server Error")
			}
		}()
		next()
	}
}
```

---

👉 This prevents:

* server crash
* returns 500 instead

---

# 🧠 12. Mental Model (Lock this in)

Think like this:

* `panic` → **something is broken**
* `recover` → **last chance to save program**

---

# 🔥 13. Golden Rules (Don’t ignore)

---

## ✅ Use panic for:

* programmer errors
* impossible states

---

## ❌ Do NOT use panic for:

* user input errors
* validation failures
* DB “not found”

---

## ✅ Always recover at boundaries:

* HTTP handlers
* goroutines
* worker systems

---

# 🚀 Final Insight

> Panic is not for control flow—it’s for **fail-fast situations**

If we use it wrongly:
👉 unstable apps

If we use it correctly:
👉 robust systems with controlled recovery

---


# CODE-FILES 💻

```go
package main

import "fmt"

// 📂 04_functions
// more on functions..

func factorial(n int)int{
	if n<=1{
		return 1
	}
	return n * factorial(n-1)
}

func intSeq() func() int{
	i:=0
	return func () int {
		i++
		return i
	}
}


func main() {
	fmt.Println(factorial(5))
	nextInt:=intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	logger:=func (msg string)  {
		fmt.Println(msg)
	}

	logger("Hiii")
	logger("Hello")
}
```

```go
package main

import "fmt"

// 📂 04_functions
// variadic functions..

func sum(nums ...int)int{
	total:=0
	for _,num:=range nums{
		total+=num
	}
	return total
}

func config(nums ...int){
	if len(nums)>0{
		first:=nums[0]
		fmt.Println("First Number:",first)
	}else{
		fmt.Println("Default Number")
	}
}


func main() {
fmt.Println(sum(1,2,3))
fmt.Println(sum(2,3))
fmt.Println(sum(2,3,1,1,1))

config(4)
config(21,22,33)
config()
}

// $ go run main.go
// 6
// 5
// 8
// First Number: 4
// First Number: 21
// Default Number
```
```go
package main

import (
	"fmt"
	"strings"
)

// 📂 04_functions
// returning multiple values from a function

func divide(a,b int)(int,error){
	//err is usually the last thing that's returned - convention
	return a/b,nil
}

func splitName(fullName string)(firstName, lastName string){
	parts:= strings.Split(fullName," ")
	firstName=parts[0]
	lastName=parts[1]

	return
}


func main() {
	val,err:=divide(10,5)
	if err!=nil{
		fmt.Println("ERROR:",err)
	}else{
	fmt.Println(val)
	}

	firstName,lastName:=splitName("Soumadip Banerjee")
	fmt.Println(firstName,lastName)
}

// $ go run main.go
// 2
// Soumadip Banerjee
```
```go
package main

import (
	"errors"
	"fmt"
	"time"
)

// 📂 04_functions
// custom error handling

var ErrDivideByZero = errors.New("division by zero!")
var ErrNumTooLarge = errors.New("number too large!")

func divide(a,b int)(int,error){
	if b==0{
		return 0,ErrDivideByZero
	}
	if a>1000{
		// silly example
		return 0,ErrNumTooLarge
	}
	return a/b,nil
}

type OpError struct{
	Op string
	Code int
	Message string
	Time time.Time
}

func (op *OpError)Error()string{
	return  op.Message
}

func NewOpError(op string,
	code int,
	message string,
	t time.Time)*OpError{
		return &OpError{
			Op:op,
			Code: code,
			Message: message,
			Time: t,
		}
}

func DoSomething()error{
	return NewOpError("do-something",100,"do something plz",time.Now())
}
func main() {
	val,err:=divide(1000000,2)
	if err != nil {
		// fmt.Println(err.Error())
		// fmt.Println(err)
		// return

		if errors.Is(err, ErrDivideByZero){
			fmt.Println("Divide by zero!")
		}else if errors.Is(err,ErrNumTooLarge){
			fmt.Println("Num too large!")
		}
	}
	fmt.Println("Value:",val)
}

// $ go run main.go
// Num too large!
// Value: 0
```
```go
package main

import "fmt"

// 📂 04_functions
// defer statement - LIFO

func simpleDefer(){
	fmt.Println("Function simpleDefer Start...")
	defer fmt.Println("Simple defer - deferred f(x)...")
	fmt.Println("Function simpleDefer Middle...")
}

func main() {
	fmt.Println("Function mainDefer: Start...")
	defer fmt.Println("First main: deferred...")
	defer fmt.Println("Second main: deferred...")
	defer func ()  {
		fmt.Println("Before the return of main")
	}()
	simpleDefer()
}

// $ go run main.go
// Function mainDefer: Start...
// Function simpleDefer Start...
// Function simpleDefer Middle...
// Simple defer - deferred f(x)...
// Before the return of main
// Second main: deferred...
// First main: deferred...
```
```go
package main

import "fmt"

// 📂 04_functions
// Panic and recovery...

func mightPanic(shouldPanic bool){
	if shouldPanic{
		panic("Something went wrong in mightPanic..!")
	}
	fmt.Println("This f(x) executed without panic!")
}

func recoverable(){
	defer func(){
		if r:=recover();r!=nil{
			fmt.Println("Recovered from PANIC!")
		}
	}()
	mightPanic(false)
}

func main() {
// panic("Something bad just happened!")
 mightPanic(false)
 recoverable()

}

// $ go run main.go
// This f(x) executed without panic!
// This f(x) executed without panic!
```
```go

```