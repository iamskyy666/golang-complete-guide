Composition in Go is one of those things that feels simple at first—but it’s actually a *core philosophy* of the language. If we really get this, we start writing idiomatic Go instead of trying to force OOP patterns from Java/C++.

---

# 🔥 What is Composition in Go?

**Composition = building complex types by combining simpler types**

Instead of inheritance (like in Java), Go says:

> “Don’t extend types. Just *compose* them.”

---

# 🧠 Why Go avoids inheritance

Languages like Java use:

```java
class Dog extends Animal {}
```

Go **does NOT have inheritance**.

Instead, Go prefers:

* smaller structs
* reusable components
* combining them together

This leads to:

* less tight coupling
* more flexibility
* easier testing

---

# 🧱 Basic Composition (Struct inside Struct)

Let’s start simple:

```go
type Engine struct {
    HorsePower int
}

type Car struct {
    Brand  string
    Engine Engine
}
```

### Usage:

```go
car := Car{
    Brand: "BMW",
    Engine: Engine{
        HorsePower: 300,
    },
}

fmt.Println(car.Engine.HorsePower)
```

👉 Here, `Car` **has an** `Engine`
This is called **HAS-A relationship**

---

# ⚡ Struct Embedding (The Real Power)

Now things get interesting.

```go
type Engine struct {
    HorsePower int
}

type Car struct {
    Brand string
    Engine   // embedded struct
}
```

### Usage:

```go
car := Car{
    Brand: "BMW",
    Engine: Engine{
        HorsePower: 300,
    },
}

fmt.Println(car.HorsePower) // 🔥 No need for car.Engine.HorsePower
```

👉 This is called **embedding**

---

# 🧩 What embedding actually does

When we embed:

```go
Engine
```

Go:

* “promotes” fields and methods
* makes them directly accessible

So internally it's still:

```go
car.Engine.HorsePower
```

But we can write:

```go
car.HorsePower
```

---

# 🧠 Composition vs Inheritance (Important mindset shift)

Let’s compare:

### ❌ Inheritance thinking:

> “Car IS-A Engine” (wrong)

### ✅ Composition thinking:

> “Car HAS-A Engine” (correct)

---

# 🔁 Method Promotion in Composition

Methods also get promoted.

```go
type Engine struct {}

func (e Engine) Start() {
    fmt.Println("Engine started")
}

type Car struct {
    Engine
}
```

### Usage:

```go
car := Car{}
car.Start() // 🔥 works directly
```

👉 Even though `Start()` belongs to `Engine`, `Car` can use it.

---

# ⚠️ Method Overriding (Sort of…)

Go doesn’t support true overriding, but we can **shadow methods**:

```go
func (c Car) Start() {
    fmt.Println("Car starting...")
}
```

Now:

```go
car.Start() // calls Car's version
```

But we can still access original:

```go
car.Engine.Start()
```

---

# 🧱 Composition with Multiple Embeddings

```go
type Engine struct {
    Power int
}

type Wheels struct {
    Count int
}

type Car struct {
    Engine
    Wheels
}
```

### Usage:

```go
car := Car{
    Engine: Engine{Power: 200},
    Wheels: Wheels{Count: 4},
}

fmt.Println(car.Power)  // from Engine
fmt.Println(car.Count)  // from Wheels
```

---

# ⚠️ Name Conflicts

If two embedded structs have same field:

```go
type A struct {
    Value int
}

type B struct {
    Value int
}

type C struct {
    A
    B
}
```

Then:

```go
c.Value ❌ // ambiguous
```

We must use:

```go
c.A.Value
c.B.Value
```

---

# 🧠 Composition + Interfaces (VERY IMPORTANT)

This is where Go becomes powerful.

We can compose behavior using interfaces:

```go
type Speaker interface {
    Speak()
}

type Dog struct {}

func (d Dog) Speak() {
    fmt.Println("Bark")
}
```

Now compose:

```go
type Robot struct {
    Speaker
}
```

👉 Robot can use ANY type that implements `Speaker`

---

# 💡 Real-world Pattern (Used everywhere)

```go
type Logger struct {}

func (l Logger) Log(msg string) {
    fmt.Println(msg)
}

type Service struct {
    Logger
}
```

Now:

```go
s := Service{}
s.Log("Hello") // 🔥 reuse without inheritance
```

---

# 🚀 Why Composition is Better (Practical View)

Let’s be real—this is why Go enforces it:

### 1. Flexible

We can mix and match behaviors easily

### 2. No deep inheritance trees

(No “God classes” like Java nightmares)

### 3. Easier testing

We can swap components

### 4. Cleaner design

Small reusable building blocks

---

# ⚠️ Common Beginner Mistake (Important)

Trying to replicate inheritance:

```go
type Animal struct {}
type Dog struct {
    Animal
}
```

Then expecting polymorphism like Java.

👉 That’s not how Go is designed.

Instead, use:

* interfaces
* composition

---

# 🧠 Mental Model to Keep

Whenever we design in Go, think:

> ❌ “What should this inherit from?”
> ✅ “What should this be composed of?”

---

If we really want to *think in Go*, struct embedding is one of those concepts we can’t treat as just syntax—it’s how Go replaces a big chunk of classical OOP.

Let’s break it down properly.

---

# 🔥 What is Struct Embedding?

Struct embedding is a special form of composition where we include a struct **without giving it a field name**.

```go
type Address struct {
	City string
}

type User struct {
	Name string
	Address // 👈 embedded
}
```

👉 This is **not just nesting**. It unlocks something extra.

---

# ⚡ What makes embedding special?

Normally (with composition):

```go
type User struct {
	Name string
	Addr Address
}
```

We access like:

```go
user.Addr.City
```

---

But with embedding:

```go
type User struct {
	Name string
	Address
}
```

We can do:

```go
user.City // 🔥 directly accessible
```

👉 This is called **field promotion**

---

# 🧠 Mental Model (Very Important)

Embedding means:

> “This struct *contains* another struct, and I want to expose its fields/methods as if they are part of me.”

---

# 🧱 Under the Hood (Reality Check)

Even though we write:

```go
user.City
```

Internally it is still:

```go
user.Address.City
```

👉 Go is just making our life easier with promotion.

---

# 🔁 Method Promotion (Even more powerful)

Embedding doesn’t just promote fields—it promotes methods too.

```go
type Logger struct {}

func (l Logger) Log(msg string) {
	fmt.Println(msg)
}

type Service struct {
	Logger
}
```

Now:

```go
s := Service{}
s.Log("Hello") // 🔥 works
```

👉 `Service` didn’t define `Log`, but still has it.

---

# ⚠️ Pointer vs Value Embedding (Important nuance)

### Value embedding:

```go
type Service struct {
	Logger
}
```

### Pointer embedding:

```go
type Service struct {
	*Logger
}
```

---

## Difference:

### Value embedding:

* copies the struct
* safe, simple

### Pointer embedding:

* shares the same instance
* allows mutation

```go
s := Service{
	Logger: &Logger{},
}
```

👉 Use pointer embedding when:

* struct is large
* we need shared state
* methods use pointer receivers

---

# ⚠️ Name Conflicts (Ambiguity)

If two embedded structs have same field:

```go
type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
	A
	B
}
```

Then:

```go
c.Name ❌ // ambiguous
```

We must do:

```go
c.A.Name
c.B.Name
```

👉 Go forces clarity instead of guessing.

---

# ⚠️ Shadowing (Overriding-like behavior)

```go
type Contact struct {
	Email string
}

type User struct {
	Contact
	Email string
}
```

Now:

```go
user.Email              // own field
user.Contact.Email      // embedded field
```

👉 This is **shadowing**, not true overriding.

---

# 🧠 Embedding + Interfaces (This is the real power)

This is where Go becomes elegant.

```go
type Reader interface {
	Read()
}

type File struct {}

func (f File) Read() {
	fmt.Println("Reading file")
}
```

Now embed:

```go
type Document struct {
	File
}
```

👉 `Document` automatically satisfies `Reader`
because `File` does.

---

# 🚀 Why Go uses embedding instead of inheritance

Let’s be blunt—inheritance creates:

* deep hierarchies
* tight coupling
* hard-to-maintain code

Embedding gives:

* flat structures
* reusable components
* flexible design

---

# 🧠 Real-world patterns

### 1. Logging

```go
type Service struct {
	Logger
}
```

### 2. HTTP handlers

```go
type MyHandler struct {
	http.Handler
}
```

### 3. Database models

```go
type BaseModel struct {
	ID int
}

type User struct {
	BaseModel
	Name string
}
```

---

# ⚠️ Common Beginner Mistakes

### ❌ Thinking embedding = inheritance

No. There is:

* no superclass
* no polymorphism like Java

---

### ❌ Overusing embedding

Not everything should be embedded.

Bad:

```go
type User struct {
	Logger
	Config
	DB
	Cache
}
```

👉 This becomes messy quickly.

---

### ❌ Ignoring explicit access when needed

Sometimes this is better:

```go
Addr Address
```

Instead of embedding.

---

# 🧠 When should we use embedding?

Use embedding when:

* we want to **promote behavior**
* we want to reuse methods cleanly
* the relationship feels like:

  > “this type is built using this capability”

Avoid embedding when:

* fields may conflict
* we need clear boundaries
* we have multiple instances (like billing/shipping address)

---

# 🎯 Final Mental Model

If we remember just one thing:

> Embedding = Composition + Promotion

---

# 🚀 Quick Summary

* Embedding = anonymous struct field
* Promotes:

  * fields
  * methods
* Enables:

  * reuse
  * cleaner APIs
* Supports:

  * interface satisfaction
* Does NOT provide:

  * inheritance
  * true overriding

---

# CODE-FILES 💻
```go
package main

import "fmt"

// 📂 06_composition_design_patterns

// Composition --> Has a relationship

// Inheritance --> Relationship (Different concept)
// Car -> composed of many parts

type Address struct {
	Street string
	City string
	State string
	ZipCode string
}

func(a Address) FullAddress()string{
	if a.Street=="" && a.City==""{
		return "No address provided!"
	}
	return fmt.Sprintf("%s, %s, %s, %s", a.Street, a.City, a.State, a.ZipCode)
}

type Customer struct{
	CustomerID int
	Name string
	Email string
	BillingAddr Address // embedded
	ShippingAddr Address // embedded
}

func (c Customer) PrintDetails(){
	fmt.Printf("Customer_ID: %d\n",c.CustomerID)
	fmt.Printf("Name: %s\n",c.Name)
	fmt.Printf("Email: %s\n",c.Email)
	fmt.Println("Billing Address:",c.BillingAddr.FullAddress())
	fmt.Println("Shiipping Address:",c.ShippingAddr.FullAddress())
	fmt.Println("----------------------------------")
}

func main() {
	fmt.Println("----------- COMPOSITION ------------")
	customer1 := Customer{
	CustomerID: 1,
	Name:       "Rahul Sharma",
	Email:      "rahul.sharma@example.com",
	BillingAddr: Address{
		Street:  "12 MG Road",
		City:    "Kolkata",
		State:   "West Bengal",
		ZipCode: "700001",
	},
	ShippingAddr: Address{
		Street:  "45 Park Street",
		City:    "Kolkata",
		State:   "West Bengal",
		ZipCode: "700016",
	},
}


// CUSTOMER WITH THE SAME SHIPPING AND BILLING-ADDR.

mainAddr:= Address{
		Street:  "22 Ballygunge Place",
		City:    "Kolkata",
		State:   "West Bengal",
		ZipCode: "700019",
	}

customer2 := Customer{
	CustomerID: 2,
	Name:       "Ananya Sen",
	Email:      "ananya.sen@example.com",
	BillingAddr: mainAddr,
	ShippingAddr: mainAddr,
}

customer1.PrintDetails()
customer2.PrintDetails()

// $ go run main.go
// ----------- COMPOSITION ------------
// Customer_ID: 1
// Name: Rahul Sharma
// Email: rahul.sharma@example.com
// Billing Address: 12 MG Road, Kolkata, West Bengal, 700001
// Shiipping Address: 45 Park Street, Kolkata, West Bengal, 700016
// ----------------------------------
// Customer_ID: 2
// Name: Ananya Sen
// Email: ananya.sen@example.com
// Billing Address: 22 Ballygunge Place, Kolkata, West Bengal, 700019
// Shiipping Address: 22 Ballygunge Place, Kolkata, West Bengal, 700019
// ----------------------------------
}
```
```go
package main

import "fmt"

// 📂 06_composition_design_patterns

// Struct-Embedding - An alt. to inheritance

type Address struct {
	Street string
	City string
	State string
	ZipCode string
}

type ContactInfo struct {
	Email string
	Phone string
}

// Embedding
type Company struct {
	Name string
	Address
	ContactInfo
	BusinessType string

}

type CompanyWithOwnEmail struct {
	Name string
	Address
	ContactInfo
	Email string

}

// methods()

func(a Address) FullAddress()string{
	if a.Street=="" && a.City==""{
		return "No address provided!"
	}
	return fmt.Sprintf("%s, %s, %s, %s", a.Street, a.City, a.State, a.ZipCode)
}


func (ci ContactInfo) DisplayContact()string{
	return fmt.Sprintf("Email: %s, Phone: %s",ci.Email,ci.Phone)
}

func (c Company) GetProfile(){
	fmt.Printf("Company Name: %s\n",c.Name)
	fmt.Printf("Location: %s\n",c.FullAddress())
	fmt.Printf("Street (promoted fx): %s\n",c.Street)

	fmt.Printf("Email (promoted fx): %s\n",c.Email)
	fmt.Printf("Business Type: %s\n",c.BusinessType)

}

func main() {
	fmt.Println("----------- Struct-EMBEDDING ------------")
	company1 := Company{
	Name: "TechNova Solutions",
	Address: Address{
		Street:  "101 Sector V",
		City:    "Kolkata",
		State:   "West Bengal",
		ZipCode: "700091",
	},
	ContactInfo: ContactInfo{
		Email: "contact@technova.com",
		Phone: "+91-9876543210",
	},
	BusinessType: "Software Development",
}

company1.GetProfile()
fmt.Println("Comp. City: ",company1.City)
fmt.Println("Comp. State: ",company1.State)
	
	


// $ go run main.go
// ----------- Struct-EMBEDDING ------------
// Company Name: TechNova Solutions
// Location: 101 Sector V, Kolkata, West Bengal, 700091
// Street (promoted fx): 101 Sector V
// Email (promoted fx): contact@technova.com
// Business Type: Software Development
// Comp. City:  Kolkata
// Comp. State:  West Bengal
}
```