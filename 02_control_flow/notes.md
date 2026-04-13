# 🧠 1. Conditionals in Go (`if`, `else`)

## 👉 Basic syntax

```go
if condition {
  // run this
} else {
  // run this
}
```

---

## ✅ Example

```go
age := 18

if age >= 18 {
  fmt.Println("Adult")
} else {
  fmt.Println("Minor")
}
```

---

## ⚠️ Important Rules

### ❌ No parentheses

```go
if (age > 18) { } // ❌ not idiomatic
```

✔ Correct:

```go
if age > 18 { }
```

---

### ❌ Condition must be boolean

```go
if age { } // ❌ invalid
```

👉 Go is strict—no truthy/falsy like JavaScript

---

## 🔥 If with initialization (VERY IMPORTANT)

This is heavily used in real code.

```go
if x := 10; x > 5 {
  fmt.Println(x)
}
```

---

### 🧠 Scope behavior

* `x` exists **only inside the if block**

```go
fmt.Println(x) // ❌ undefined
```

---

### Real-world usage

```go
if err := doSomething(); err != nil {
  fmt.Println("Error:", err)
}
```

👉 This is idiomatic Go

---

# 🔀 2. Switch Statement

Cleaner alternative to multiple `if-else`.

---

## 👉 Basic example

```go
day := 2

switch day {
case 1:
  fmt.Println("Monday")
case 2:
  fmt.Println("Tuesday")
default:
  fmt.Println("Unknown")
}
```

---

## 🔥 Important Features

---

### ✅ No break needed

Unlike C/JS:

* Go automatically breaks after each case

---

### 🔥 Multiple values in one case

```go
switch day {
case 1, 2, 3:
  fmt.Println("Start of week")
}
```

---

### 🔥 Switch without expression

```go
x := 10

switch {
case x > 5:
  fmt.Println("Greater than 5")
case x < 5:
  fmt.Println("Less than 5")
}
```

👉 Works like `if-else`

---

### 🔥 Fallthrough (rare but important)

```go
switch x := 1; x {
case 1:
  fmt.Println("One")
  fallthrough
case 2:
  fmt.Println("Two")
}
```

👉 Forces next case to execute

---

⚠️ Use carefully—can confuse logic

---

# 🔁 3. Loops in Go (ONLY ONE LOOP: `for`)

Go has:

> ❗ No `while`, no `do-while`
> Only `for`

But `for` can behave like all of them.

---

## 👉 Basic for loop

```go
for i := 0; i < 5; i++ {
  fmt.Println(i)
}
```

---

## 👉 While-like loop

```go
i := 0

for i < 5 {
  fmt.Println(i)
  i++
}
```

---

## 👉 Infinite loop

```go
for {
  fmt.Println("Running forever")
}
```

---

👉 Used in:

* Servers
* Workers
* Background jobs

---

# 🔥 4. Range Loop (VERY IMPORTANT)

Used for iterating collections.

---

## 👉 Slice example

```go
nums := []int{10, 20, 30}

for index, value := range nums {
  fmt.Println(index, value)
}
```

---

## 👉 Ignoring values

```go
for _, value := range nums {
  fmt.Println(value)
}
```

---

## 👉 Map example

```go
m := map[string]int{"a": 1, "b": 2}

for key, value := range m {
  fmt.Println(key, value)
}
```

---

## 👉 String iteration (UTF-8 aware)

```go
for i, ch := range "Go" {
  fmt.Println(i, ch)
}
```

👉 `ch` is a **rune**, not byte

---

# ⚠️ Range Gotcha (Important)

```go
nums := []int{1,2,3}

for _, v := range nums {
  v = v * 2
}
```

👉 ❌ Does NOT modify original slice

---

✔ Correct:

```go
for i := range nums {
  nums[i] *= 2
}
```

---

# 🛑 5. Loop Control Keywords

---

## 🔹 break

Stops loop

```go
for i := 0; i < 10; i++ {
  if i == 5 {
    break
  }
}
```

---

## 🔹 continue

Skip current iteration

```go
for i := 0; i < 5; i++ {
  if i == 2 {
    continue
  }
  fmt.Println(i)
}
```

---

# 🧠 6. Nested Loops

```go
for i := 0; i < 3; i++ {
  for j := 0; j < 3; j++ {
    fmt.Println(i, j)
  }
}
```

---

👉 Used in:

* Matrices
* Grid problems
* Algorithms

---

# ⚠️ Common Mistakes (Fix these early)

---

## ❌ Forgetting `range` behavior

* It gives **copy**, not reference

---

## ❌ Infinite loop by mistake

```go
for i < 5 {
  // forgot i++
}
```

---

## ❌ Overusing fallthrough

👉 Makes code harder to read

---

## ❌ Using index when not needed

```go
for i := 0; i < len(nums); i++ { } // less idiomatic
```

✔ Better:

```go
for _, v := range nums { }
```

---

# 🧠 Mental Model

Think like this:

* `if` → decision making
* `switch` → cleaner multi-branch logic
* `for` → everything looping
* `range` → iterate collections

---

# 🚀 What actually matters for us (backend focus)

We’ll use this constantly for:

* ✅ Validating requests (`if err != nil`)
* ✅ Iterating DB results (`range`)
* ✅ Handling statuses (`switch`)
* ✅ Running workers (`for {}` infinite loops)

---
# CODE FILES 💻
## 1️⃣ LOOPS
```go
package main

import "fmt"

func main() {
	fmt.Println("----------C-style FOR loop--------")
	for i:=1;i<=10;i++{
    fmt.Println(i)
	}

	fmt.Println("----------WHILE loop--------")
	k:=10
	for k>0{
		fmt.Println(k)
		k--
	}

	fmt.Println("----------INFINITE loop--------")
	counter:=0
	for{
		fmt.Println("Counter:",counter)
		counter++
		if counter>=5{
			break
		}
	}

	fmt.Println("----------SKIPPING --------")
	for i:=1;i<=10;i++{
		if i%2==0{
			continue
		}
    fmt.Println(i)
	}

	fmt.Println("----------ARRAYS[] --------")
	skillset:=[3]string{"ReactJs","Golang","NodeJs"}
	for _,v:= range skillset{
		fmt.Println(v)
	}
	
}

// O/P:

// $ go run main.go
// ----------C-style FOR loop--------
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 10
// ----------WHILE loop--------
// 10
// 9
// 8
// 7
// 6
// 5
// 4
// 3
// 2
// 1
// ----------INFINITE loop------
// Counter: 0
// Counter: 1
// Counter: 2
// Counter: 3
// Counter: 4
// ----------SKIPPING --------
// 1
// 3
// 5
// 7
// 9
// ----------ARRAYS[] --------
// ReactJs
// Golang
// NodeJs
```

## 2️⃣ IF-ELSE 
```go
package main

import "fmt"

func main() {
	temp:=25
	if temp>30{
		fmt.Println("Greater than 30 ☀️")
	}else{
		fmt.Println("Less than 30 🌿")
	}

	score:=85
	if score>=80 {
		fmt.Println("Grade: A")
	}else if score>=60{
		fmt.Println("Grade: B")
	}else{
		fmt.Println("Could improve!")
	}

	userAccess:=map[string]bool{
		"Skyy":true,
		"Soumadip":false,
		"Banerjee":true,
	}

	if hasAccess,ok:=userAccess["Skyy"];ok && hasAccess{
		fmt.Println("Skyy can access this system ✅")
	}else{
		fmt.Println("Access denied! ⚠️")
	}
	
}

// O/P:

// $ go run main.go
// Less than 30 🌿
// Grade: A
// Skyy can access this system ✅
```

## 3️⃣. SWITCH-statements
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	country:="Brazil"
	fmt.Println("I wanna travel to",country)

	switch country{
	case "Egypt":
		fmt.Println("Pyramids in Cairo! 🐪")
	case "Switzerland":
		fmt.Println("Skiing in the Alps! ⛷️")	
	case "Brazil":
		fmt.Println("Amazon-adventure! 🐍")	
	default:
		fmt.Println("UNKNOWN Country 🤔")	
	}

	switch hour:=time.Now().Hour();{
	case hour<12:
		fmt.Println("Guten Morgen!")
	case hour<17:
		fmt.Println("Guten Abend!")	
	default:
		fmt.Println("Servus!")	
	}	

	checkType:=func(i any){
		switch v:=i.(type){
		case int:
			fmt.Printf("Int value: %d\n",v)
		case string:
			fmt.Printf("String value: %s\n",v)	
		case bool:
			fmt.Printf("Boolean value: %t\n",v)	
		default:
			fmt.Printf("Unknown type: %T\n",v)	
		}
	}
	checkType(21)
	checkType("String test..")
	checkType(true)
}

// O/P:

// $ go run main.go
// I wanna travel to Brazil
// Amazon-adventure! 🐍
// Servus!
// Int value: 21
// String value: String test..
// Boolean value: true
```