# What is SQLite? 🗃️

SQLite is a **serverless relational database**.

That single sentence explains almost everything important.

Unlike databases like PostgreSQL or MySQL, SQLite does **not** run as a separate database server process.

Instead:

* The entire database is stored in a **single file**
* Our Go application talks directly to that file
* No database server setup is needed
* No username/password configuration is required for local development
* Perfect for beginners, small apps, desktop apps, CLI tools, prototypes, local tools, and learning SQL

Example:

```txt
myapp.db
```

That file itself is the database.

---

# Why SQLite is Amazing for Beginners

SQLite is one of the best databases to start with because:

| Feature                   | SQLite           |
| ------------------------- | ---------------- |
| Easy setup                | ✅ Extremely easy |
| Separate DB server needed | ❌ No             |
| Database is a file        | ✅ Yes            |
| Fast for local apps       | ✅ Very           |
| SQL support               | ✅ Full SQL       |
| Good for learning         | ✅ Excellent      |
| Used in real apps         | ✅ Yes            |

SQLite is used in:

* Android
* iOS
* Browsers
* Desktop apps
* Embedded systems
* Local tools
* Small SaaS projects
* Testing environments

---

# How Databases Work Conceptually

Before Go code, we must understand the mental model.

A relational database stores data in **tables**.

Example:

## Users Table

| id | name | age |
| -- | ---- | --- |
| 1  | Skyy | 29  |
| 2  | Alex | 25  |

Each row = a record.

Each column = a field.

---

# SQL — The Language of Databases

Databases use SQL.

SQL = Structured Query Language.

Common SQL commands:

| Command      | Purpose      |
| ------------ | ------------ |
| CREATE TABLE | create table |
| INSERT INTO  | add data     |
| SELECT       | read data    |
| UPDATE       | modify data  |
| DELETE       | remove data  |

Example:

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT,
    age INTEGER
);
```

---

# How Go Works with Databases

Go uses a package called:

```go
database/sql
```

Important:

`database/sql` does NOT talk to databases directly.

Instead:

```txt
Go app
   ↓
database/sql
   ↓
Database Driver
   ↓
SQLite / PostgreSQL / MySQL
```

We need:

1. `database/sql`
2. A database driver

For SQLite, a popular driver is:

```txt
github.com/mattn/go-sqlite3
```

---

# Installing SQLite Driver in Go

Inside our Go project:

```bash
go mod init myapp
```

Then install driver:

```bash
go get github.com/mattn/go-sqlite3
```

---

# First SQLite Program in Go

---

# Step 1 — Import Packages

```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)
```

---

# VERY IMPORTANT — Blank Identifier Import

This line:

```go
_ "github.com/mattn/go-sqlite3"
```

looks strange.

Why `_`?

Because:

* We don't directly use the package
* The package registers itself internally with `database/sql`

This is called a **side-effect import**.

Without it:

```go
sql.Open("sqlite3", ...)
```

would fail because Go wouldn't know the driver.

---

# Step 2 — Open Database

```go
db, err := sql.Open("sqlite3", "./test.db")
if err != nil {
	log.Fatal(err)
}
defer db.Close()
```

Explanation:

```go
sql.Open(driverName, databasePath)
```

* `"sqlite3"` = driver
* `"./test.db"` = database file

If file doesn't exist:

✅ SQLite creates it automatically.

---

# Step 3 — Create Table

```go
query := `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	age INTEGER
);
`
```

Execute it:

```go
_, err = db.Exec(query)
if err != nil {
	log.Fatal(err)
}
```

---

# What is Exec()?

Used for queries that do NOT return rows.

Examples:

* CREATE TABLE
* INSERT
* UPDATE
* DELETE

---

# Full Example

```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER
	);
	`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table created successfully")
}
```

---

# What Happens After Running?

SQLite creates:

```txt
test.db
```

inside our project folder.

That file contains the entire database.

---

# Inserting Data

SQL:

```sql
INSERT INTO users(name, age)
VALUES('Skyy', 29);
```

Go:

```go
insertQuery := `
INSERT INTO users(name, age)
VALUES(?, ?)
`

_, err = db.Exec(insertQuery, "Skyy", 29)
if err != nil {
	log.Fatal(err)
}
```

---

# What are `?` Placeholders?

These are parameter placeholders.

Very important for:

* security
* preventing SQL injection
* safer queries

Bad:

```go
query := "INSERT INTO users VALUES('" + name + "')"
```

Good:

```go
VALUES(?, ?)
```

---

# Reading Data (SELECT)

Now the interesting part.

---

# Query()

Used when multiple rows are returned.

```go
rows, err := db.Query("SELECT id, name, age FROM users")
if err != nil {
	log.Fatal(err)
}
defer rows.Close()
```

---

# rows.Next()

Moves to next row.

```go
for rows.Next() {

}
```

---

# Scan()

Copies column values into Go variables.

```go
var id int
var name string
var age int

err := rows.Scan(&id, &name, &age)
```

IMPORTANT:

We pass pointers because Scan modifies variables.

---

# Full Read Example

```go
rows, err := db.Query("SELECT id, name, age FROM users")
if err != nil {
	log.Fatal(err)
}
defer rows.Close()

for rows.Next() {

	var id int
	var name string
	var age int

	err := rows.Scan(&id, &name, &age)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id, name, age)
}
```

---

# Output

```txt
1 Skyy 29
2 Alex 25
```

---

# QueryRow()

Used when expecting ONE row.

Example:

```go
var name string

err := db.QueryRow(
	"SELECT name FROM users WHERE id = ?",
	1,
).Scan(&name)

if err != nil {
	log.Fatal(err)
}

fmt.Println(name)
```

---

# Updating Data

SQL:

```sql
UPDATE users
SET age = 30
WHERE id = 1;
```

Go:

```go
_, err = db.Exec(
	"UPDATE users SET age = ? WHERE id = ?",
	30,
	1,
)
```

---

# Deleting Data

```go
_, err = db.Exec(
	"DELETE FROM users WHERE id = ?",
	1,
)
```

---

# Structs + Database

Usually we map rows into structs.

Example:

```go
type User struct {
	ID   int
	Name string
	Age  int
}
```

Then:

```go
var users []User

for rows.Next() {

	var user User

	err := rows.Scan(
		&user.ID,
		&user.Name,
		&user.Age,
	)

	if err != nil {
		log.Fatal(err)
	}

	users = append(users, user)
}
```

This is how real applications work.

---

# Understanding the Flow Deeply

The full flow:

```txt
Go Struct
   ↓
SQL Query
   ↓
SQLite Database File
   ↓
Returned Rows
   ↓
Scan into Struct
```

---

# Common SQLite Data Types

SQLite is flexible with types.

| SQLite Type | Go Type |
| ----------- | ------- |
| INTEGER     | int     |
| TEXT        | string  |
| REAL        | float64 |
| BLOB        | []byte  |
| BOOLEAN     | bool    |

---

# Primary Keys

```sql
id INTEGER PRIMARY KEY AUTOINCREMENT
```

Meaning:

* unique ID
* auto-generated
* increments automatically

Example:

```txt
1
2
3
4
```

---

# SQL Injection — Critical Concept

Never build raw queries from user input.

BAD:

```go
query := "SELECT * FROM users WHERE name = '" + name + "'"
```

Attackers can inject SQL.

ALWAYS use placeholders:

```go
WHERE name = ?
```

---

# Error Handling in Database Code

Database operations fail often:

* DB locked
* syntax error
* missing table
* invalid query
* closed connection

So Go code must always check errors.

---

# defer db.Close()

Important.

Releases database resources properly.

---

# What is a Driver?

The driver translates Go database calls into SQLite operations.

Different databases need different drivers.

| Database   | Driver       |
| ---------- | ------------ |
| SQLite     | go-sqlite3   |
| PostgreSQL | pgx          |
| MySQL      | mysql driver |

But `database/sql` API remains mostly similar.

This is one of Go's biggest strengths.

---

# Beginner Project Ideas

SQLite + Go is fantastic for:

1. Notes app
2. Todo app
3. Expense tracker
4. CLI password manager
5. Blog backend
6. Local inventory system
7. Medicine tracker
8. Camp booking system

---

# Recommended Learning Order

For someone learning Go databases:

## Phase 1

Learn:

* SQL basics
* CREATE
* INSERT
* SELECT
* UPDATE
* DELETE

## Phase 2

Learn:

* database/sql
* Query()
* QueryRow()
* Exec()
* Scan()

## Phase 3

Learn:

* Struct mapping
* Relationships
* Transactions

## Phase 4

Move to:

* PostgreSQL
* REST APIs
* ORMs
* Authentication systems

---

# What is an ORM?

ORM = Object Relational Mapper.

Popular Go ORM:

GORM

It lets us write:

```go
db.Create(&user)
```

instead of raw SQL.

But beginners should first learn raw SQL + `database/sql`.

That foundation matters a lot.

---

# Important Beginner Reality

Many new developers jump straight into ORMs.

That creates weak fundamentals.

If we understand:

* SQL
* tables
* queries
* relationships
* scanning rows

then ORMs become easy later.

Without SQL knowledge, ORMs become confusing magic.

---

# SQLite vs PostgreSQL

| Feature       | SQLite              | PostgreSQL            |
| ------------- | ------------------- | --------------------- |
| Setup         | Easy                | Harder                |
| Server needed | No                  | Yes                   |
| Scalability   | Smaller apps        | Large apps            |
| Performance   | Excellent local     | Excellent distributed |
| Best for      | Learning/small apps | Production backend    |

SQLite is the perfect stepping stone before PostgreSQL.

---

# One Important Limitation of SQLite

SQLite allows many readers.

But write concurrency is limited compared to PostgreSQL.

So massive-scale apps usually use PostgreSQL/MySQL.

But for learning?

SQLite is outstanding.

---

# Final Mental Model

Think of SQLite like this:

```txt
Excel file + SQL powers + reliability
```

except much more powerful.

And Go talks to it through:

```txt
database/sql + sqlite driver
```

That’s the core architecture.

# What is a Prepared Statement?

A prepared statement is a **precompiled SQL query template**.

Instead of sending a raw SQL query to the database every time:

```txt id="kj7q0x"
Parse SQL
→ Validate SQL
→ Compile SQL
→ Execute SQL
```

we prepare it once:

```txt id="9xjq9v"
Prepare once
→ Reuse many times
```

This improves:

* performance
* organization
* security
* readability

---

# Mental Model

Think of it like a reusable function.

Without prepared statement:

```txt id="8m6zrq"
Write full SQL every time
```

With prepared statement:

```txt id="jlwm01"
Create reusable SQL template
```

---

# Basic Example

Without prepare:

```go id="jlwm02"
_, err := db.Exec(
	"INSERT INTO users(name,email) VALUES(?, ?)",
	name,
	email,
)
```

This is perfectly valid.

But every time we call it:

* SQL gets parsed again
* query gets processed again

---

# With Prepare

```go id="jlwm03"
stmt, err := db.Prepare(
	"INSERT INTO users(name,email) VALUES(?, ?)",
)
```

Now the database creates a reusable prepared query.

Then:

```go id="jlwm04"
_, err = stmt.Exec(name, email)
```

can be called repeatedly.

---

# Full Example

```go id="jlwm05"
stmt, err := db.Prepare(`
	INSERT INTO users(name,email)
	VALUES(?, ?)
`)
if err != nil {
	log.Fatal(err)
}

defer stmt.Close()

_, err = stmt.Exec("Skyy", "skyy@test.com")
if err != nil {
	log.Fatal(err)
}

_, err = stmt.Exec("Bruce", "bruce@test.com")
if err != nil {
	log.Fatal(err)
}
```

Same prepared query reused multiple times.

---

# Why Prepared Statements Exist

There are several major reasons.

---

# 1. Performance

This is the classic reason.

---

## Without Prepared Statement

Every execution may involve:

```txt id="3tjlwm"
SQL string
   ↓
Parse
   ↓
Validate
   ↓
Build execution plan
   ↓
Execute
```

Repeated over and over.

---

## With Prepared Statement

Database does:

```txt id="rjlwm06"
Prepare once
   ↓
Reuse execution plan
   ↓
Execute faster repeatedly
```

Especially useful in:

* loops
* bulk inserts
* high-traffic apps

---

# Example — Bulk Insert

Suppose we insert 10,000 users.

BAD approach:

```go id="jlwm07"
for _, user := range users {
	db.Exec(...)
}
```

Database repeatedly parses SQL.

---

BETTER:

```go id="jlwm08"
stmt, _ := db.Prepare(...)

for _, user := range users {
	stmt.Exec(...)
}
```

Now SQL structure is reused.

Huge difference at scale.

---

# 2. Security

Prepared statements help prevent SQL injection.

This is CRITICAL.

---

# SQL Injection Problem

BAD:

```go id="jlwm09"
query := `
	SELECT * FROM users
	WHERE email='` + email + `'
`
```

Attacker enters:

```txt id="jlwm10"
' OR 1=1 --
```

Now query becomes dangerous.

---

# Prepared Statements Separate:

## SQL Structure

from:

## User Data

This is the key idea.

---

Prepared query:

```sql id="jlwm11"
SELECT * FROM users WHERE email = ?
```

User input is treated as:

* data
* NOT executable SQL

This is one of the biggest security concepts in backend development.

---

# 3. Cleaner Code

Prepared statements make repeated operations cleaner.

Example:

```go id="jlwm12"
stmt, err := db.Prepare(`
	UPDATE users
	SET name = ?
	WHERE id = ?
`)
```

Then reuse:

```go id="jlwm13"
stmt.Exec("Bruce", 1)
stmt.Exec("Clark", 2)
stmt.Exec("Diana", 3)
```

Cleaner than rewriting SQL repeatedly.

---

# What Does Prepare() Return?

It returns:

```go id="jlwm14"
*sql.Stmt
```

Meaning:

* prepared SQL statement object

Think of it as:

```txt id="jlwm15"
Reusable database query object
```

---

# Important Lifecycle

---

# Step 1 — Prepare

```go id="jlwm16"
stmt, err := db.Prepare(...)
```

---

# Step 2 — Execute Multiple Times

```go id="jlwm17"
stmt.Exec(...)
stmt.Exec(...)
stmt.Exec(...)
```

---

# Step 3 — Close Statement

```go id="jlwm18"
defer stmt.Close()
```

Very important.

Prepared statements use database resources.

---

# stmt.Exec()

Same idea as:

```go id="jlwm19"
db.Exec()
```

except:

* SQL already prepared
* only values are supplied

---

# stmt.Query()

Prepared SELECT:

```go id="jlwm20"
stmt, err := db.Prepare(`
	SELECT id,name
	FROM users
	WHERE age > ?
`)
```

Then:

```go id="jlwm21"
rows, err := stmt.Query(18)
```

---

# stmt.QueryRow()

Single row:

```go id="jlwm22"
stmt, err := db.Prepare(`
	SELECT name
	FROM users
	WHERE id = ?
`)
```

Then:

```go id="jlwm23"
err = stmt.QueryRow(1).Scan(&name)
```

---

# Important Beginner Understanding

This:

```go id="jlwm24"
db.Exec(...)
```

actually often uses prepared statements internally behind the scenes in many drivers/databases.

But explicit prepare gives:

* more control
* better reuse
* clearer optimization

---

# When Should We Use Prepare?

---

# Good Use Cases

✅ repeated queries
✅ loops
✅ bulk inserts
✅ repeated updates
✅ high-performance systems
✅ reusable repository methods

---

# Probably Unnecessary For

❌ one-time queries
❌ tiny scripts
❌ very small apps

---

# Real Backend Example

Suppose:

* 10,000 login requests
* same login query repeatedly

Prepared statement:

```sql id="jlwm25"
SELECT id, hashed_password
FROM users
WHERE email = ?
```

Prepared once.

Executed thousands of times efficiently.

---

# Important Distinction

Prepared statements do NOT mean:

* results are cached
* rows are stored

Only:

* SQL structure/execution plan is prepared

Data is still fetched fresh.

---

# One Extremely Important Thing

This is safe:

```go id="jlwm26"
WHERE email = ?
```

because placeholders are parameterized.

But THIS is dangerous:

```go id="jlwm27"
"WHERE email = '" + email + "'"
```

Never concatenate raw user input into SQL.

Ever.

---

# Real Architecture Pattern

In professional Go backend systems:

```txt id="jlwm28"
Repository layer
   ↓
Prepared statements
   ↓
Reusable database operations
```

Very common.

---

# Example Repository Pattern

```go id="jlwm29"
type UserRepo struct {
	createStmt *sql.Stmt
	getStmt    *sql.Stmt
}
```

Prepared once during app startup.

Reused throughout application lifetime.

Very efficient.

---

# Tradeoff

Prepared statements:

* improve repeated query performance
* use extra DB resources

So preparing thousands of unique one-time queries is wasteful.

---

# Final Mental Model

Think of prepared statements like this:

Without prepare:

```txt id="jlwm30"
"Build the SQL machine every time"
```

With prepare:

```txt id="jlwm31"
"Build the machine once, feed it new values repeatedly"
```

That’s the core idea.

# What is a Database Transaction?

A transaction is a group of database operations treated as **one single unit of work**.

Meaning:

```txt id="8x0jlwm"
Either ALL operations succeed
OR
NONE of them happen
```

This is one of the most important concepts in backend engineering.

---

# Why Transactions Exist

Imagine a banking system.

Suppose we transfer money:

```txt id="jlwm01"
Alice → Bob
₹500
```

What actually happens internally?

---

# Step 1

Subtract ₹500 from Alice.

---

# Step 2

Add ₹500 to Bob.

---

# Dangerous Situation

What if:

```txt id="jlwm02"
Step 1 succeeds
BUT
Step 2 fails
```

Now:

* Alice lost money
* Bob never received it

Data becomes corrupted.

This is EXACTLY the kind of problem transactions solve.

---

# Transaction Guarantee

With a transaction:

```txt id="jlwm03"
Both operations succeed together
OR
everything gets undone
```

This is called:

```txt id="jlwm04"
ROLLBACK
```

---

# Core Transaction Principle

Transactions guarantee database consistency.

---

# Real-World Examples

Transactions are critical for:

* banking
* payments
* ecommerce orders
* booking systems
* inventory systems
* auth/account creation
* wallet systems
* ticket reservations

---

# The 4 ACID Properties

Transactions are based on something famous called:

# ACID

---

# A — Atomicity

A \Rightarrow \text{All or Nothing}

A transaction is indivisible.

Either:

* all operations succeed
* or none happen

---

# C — Consistency

Database remains valid before and after transaction.

Rules stay preserved.

Example:

* balances never become negative accidentally
* foreign keys remain valid

---

# I — Isolation

Multiple transactions do not interfere incorrectly with each other.

Important for concurrency.

---

# D — Durability

Once committed:

```txt id="jlwm05"
data survives crashes
```

even if:

* app crashes
* server restarts
* power goes out

---

# How Transactions Work in Go

Go uses:

```go id="jlwm06"
db.Begin()
```

to start a transaction.

This returns:

```go id="jlwm07"
*sql.Tx
```

Transaction object.

---

# Important Architecture Change

Normally:

```go id="jlwm08"
db.Exec()
db.Query()
```

But inside transaction:

```go id="jlwm09"
tx.Exec()
tx.Query()
```

VERY important.

Once transaction begins, all operations should go through `tx`.

---

# Basic Transaction Flow

---

# Step 1 — Begin Transaction

```go id="jlwm10"
tx, err := db.Begin()
```

---

# Step 2 — Execute Queries

```go id="jlwm11"
tx.Exec(...)
tx.Query(...)
```

---

# Step 3A — Commit if Successful

```go id="jlwm12"
tx.Commit()
```

Database permanently saves changes.

---

# Step 3B — Rollback if Error

```go id="jlwm13"
tx.Rollback()
```

Undo everything.

---

# Simple Example

Suppose:

* create user
* create profile

Both should succeed together.

---

# Without Transaction

Dangerous:

```txt id="jlwm14"
User created
Profile creation failed
```

Now database is partially broken.

---

# With Transaction

Safe.

---

# Full Example

```go id="jlwm15"
func CreateUserAndProfile(db *sql.DB) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO users(name)
		VALUES(?)
	`, "Skyy")

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO profiles(user_id,bio)
		VALUES(?,?)
	`, 1, "Developer")

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
```

---

# What Happens Internally

---

# Scenario 1 — Everything Succeeds

```txt id="jlwm16"
Begin
  ↓
Insert user
  ↓
Insert profile
  ↓
Commit
```

Database saves both.

---

# Scenario 2 — Error Happens

```txt id="jlwm17"
Begin
  ↓
Insert user
  ↓
Insert profile fails
  ↓
Rollback
```

Database undoes everything.

User insert disappears too.

---

# This Is Extremely Powerful

Transactions prevent:

* partial writes
* corrupted state
* inconsistent data

---

# Very Important Beginner Rule

Inside a transaction:

❌ DON'T use:

```go id="jlwm18"
db.Exec()
```

Use:

```go id="jlwm19"
tx.Exec()
```

Because transaction operations must stay inside the transaction context.

---

# Common Beginner Mistake

WRONG:

```go id="jlwm20"
tx, _ := db.Begin()

db.Exec(...)
```

This bypasses transaction entirely.

---

# Correct

```go id="jlwm21"
tx.Exec(...)
```

---

# Better Pattern Using defer

Very common production pattern:

```go id="jlwm22"
func Example(db *sql.DB) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec(...)
	if err != nil {
		return err
	}

	_, err = tx.Exec(...)
	if err != nil {
		return err
	}

	return tx.Commit()
}
```

---

# Why defer Rollback() Works

Important subtlety.

If:

* commit succeeds
* transaction already closed

then rollback becomes harmless.

But if:

* any error occurs
* function exits early

rollback automatically runs.

Very elegant Go pattern.

---

# Real Ecommerce Example

Suppose customer buys product.

We must:

1. create order
2. reduce inventory
3. create payment record
4. create shipment record

If step 3 fails:

```txt id="jlwm23"
Everything must undo
```

Otherwise inventory becomes wrong.

Transactions solve this.

---

# Transactions + Concurrency

Transactions also help prevent race conditions.

Example:

```txt id="jlwm24"
Two people buying last item simultaneously
```

Without transaction:

* both may succeed incorrectly

With proper transaction/isolation:

* only one succeeds safely

---

# SQLite and Transactions

SQLite fully supports transactions.

In fact:

* SQLite internally wraps many operations in transactions automatically

But explicit transactions are still extremely important.

---

# Performance Benefits

Transactions are not only about safety.

They also improve performance.

---

# Example

BAD:

```go id="jlwm25"
for _, user := range users {
	db.Exec(...)
}
```

Each insert may commit separately.

Very slow.

---

BETTER:

```go id="jlwm26"
tx, _ := db.Begin()

for _, user := range users {
	tx.Exec(...)
}

tx.Commit()
```

Huge performance improvement.

Because:

* fewer disk sync operations
* fewer commits

---

# Transactions + Prepared Statements

Very common combination:

```go id="jlwm27"
tx, _ := db.Begin()

stmt, _ := tx.Prepare(...)

for ... {
	stmt.Exec(...)
}

tx.Commit()
```

This is extremely common in high-performance systems.

---

# Important Limitation

Transactions are temporary.

Until commit:

```txt id="jlwm28"
changes are not permanently saved
```

Think of transaction like:

```txt id="jlwm29"
temporary private workspace
```

Commit = publish changes.

Rollback = discard changes.

---

# One Very Important Backend Insight

Many critical production bugs come from:

❌ forgetting transactions

Especially when:

* multiple tables involved
* financial operations
* inventory systems
* auth/session systems

Learning transactions early is a major backend milestone.

---

# Final Mental Model

Think of a transaction like this:

```txt id="jlwm30"
Start protected mode
   ↓
Perform multiple DB operations
   ↓
If everything succeeds:
    SAVE ALL
Else:
    UNDO ALL
```

That’s the heart of database transactions.

# What is the Repository Pattern?

The Repository Pattern is an architectural pattern used to separate:

```txt id="jlwm01"
Database logic
FROM
Business/application logic
```

It creates a clean boundary between:

* our app
* and the database

---

# Core Idea

Instead of scattering SQL everywhere:

❌ BAD:

```txt id="jlwm02"
main.go
handlers.go
services.go
auth.go
```

all directly running SQL queries...

we centralize database operations into repositories.

---

# Mental Model

Think of a repository as:

```txt id="jlwm03"
A database access layer
```

or:

```txt id="jlwm04"
A collection of database-related methods
```

---

# Why Repository Pattern Exists

Without it:

* SQL becomes scattered
* code becomes messy
* harder to test
* harder to maintain
* harder to switch databases
* duplicated queries everywhere

---

# Real Problem Without Repository Pattern

Suppose we repeatedly do:

```go id="jlwm05"
SELECT * FROM users WHERE email = ?
```

in:

* login
* signup
* password reset
* profile page

Now query logic becomes duplicated everywhere.

Very common beginner problem.

---

# Repository Pattern Solution

Create:

```txt id="jlwm06"
UserRepository
```

with methods like:

```go id="jlwm07"
GetByEmail()
Create()
Delete()
Update()
```

Now all SQL lives in one place.

Cleaner architecture.

---

# Basic Architecture

```txt id="jlwm08"
Handlers / Controllers
        ↓
Services / Business Logic
        ↓
Repositories
        ↓
Database
```

---

# Example Without Repository Pattern

BAD:

```go id="jlwm09"
func Login(db *sql.DB, email string) {

	row := db.QueryRow(`
		SELECT id,email,password
		FROM users
		WHERE email = ?
	`, email)

	...
}
```

Then later:

```go id="jlwm10"
func ResetPassword(db *sql.DB, email string) {

	row := db.QueryRow(`
		SELECT id,email,password
		FROM users
		WHERE email = ?
	`, email)

	...
}
```

Repeated logic everywhere.

---

# With Repository Pattern

Cleaner:

```go id="jlwm11"
user, err := repo.GetByEmail(email)
```

The SQL details become hidden inside repository.

---

# First Repository Example

---

# Step 1 — Define Repository Struct

```go id="jlwm12"
type UserRepository struct {
	db *sql.DB
}
```

This repository owns database access.

---

# Step 2 — Constructor

```go id="jlwm13"
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
```

---

# Step 3 — Repository Methods

```go id="jlwm14"
func (r *UserRepository) Create(
	name,
	email,
	password string,
) (int64, error) {

	stmt := `
	INSERT INTO users(
		name,
		email,
		hashed_password
	)
	VALUES(?, ?, ?)
	`

	res, err := r.db.Exec(
		stmt,
		name,
		email,
		password,
	)

	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}
```

---

# Step 4 — Usage

```go id="jlwm15"
repo := NewUserRepository(db)

id, err := repo.Create(
	"Skyy",
	"skyy@test.com",
	"12345",
)
```

Notice:

```txt id="jlwm16"
main() no longer contains SQL
```

That’s the key idea.

---

# Real Benefit

Now if database changes:

```txt id="jlwm17"
SQLite → PostgreSQL
```

we mainly modify repository layer.

The rest of app stays mostly unchanged.

This is huge in production systems.

---

# Repository Encapsulation

Repository hides:

* SQL syntax
* query details
* scanning logic
* table structure

Application only sees:

```go id="jlwm18"
repo.CreateUser(...)
repo.GetByEmail(...)
```

Much cleaner abstraction.

---

# Common Repository Methods

Usually repositories contain CRUD operations.

---

# Create

```go id="jlwm19"
Create()
```

---

# Read

```go id="jlwm20"
GetByID()
GetByEmail()
GetAll()
```

---

# Update

```go id="jlwm21"
Update()
```

---

# Delete

```go id="jlwm22"
Delete()
```

---

# Example — GetByEmail

```go id="jlwm23"
func (r *UserRepository) GetByEmail(
	email string,
) (*User, error) {

	stmt := `
	SELECT id,name,email
	FROM users
	WHERE email = ?
	`

	row := r.db.QueryRow(stmt, email)

	var user User

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
```

---

# Why This Matters So Much

Because backend apps grow FAST.

Today:

* 2 queries

Tomorrow:

* 200 queries

Without structure:

* chaos happens quickly

---

# Repository Pattern Helps With Testing

VERY important.

Without repository pattern:

```go id="jlwm24"
func Login(db *sql.DB)
```

hard to test.

---

With repository interface:

```go id="jlwm25"
type UserRepository interface {
	GetByEmail(email string) (*User, error)
}
```

Now we can create:

* mock repositories
* fake repositories
* test repositories

Huge advantage.

---

# Example Mock Repository

```go id="jlwm26"
type MockUserRepo struct{}

func (m *MockUserRepo) GetByEmail(
	email string,
) (*User, error) {

	return &User{
		ID: 1,
		Email: email,
	}, nil
}
```

Now business logic can be tested without real DB.

Very powerful.

---

# Repository Pattern + Services

Large systems usually separate:

---

# Repository

Handles:

* SQL
* persistence
* database operations

---

# Service Layer

Handles:

* business rules
* validation
* workflows
* permissions

---

# Example

Repository:

```go id="jlwm27"
repo.CreateUser()
```

Service:

```go id="jlwm28"
hash password
validate email
check uniqueness
send verification email
```

This separation becomes extremely valuable in larger systems.

---

# Transactions + Repository Pattern

Repositories often accept transactions too.

Example:

```go id="jlwm29"
func (r *UserRepository) CreateTx(
	tx *sql.Tx,
	user User,
) error
```

Now repositories can participate in transactions cleanly.

Very common architecture.

---

# Common Folder Structure

Typical Go backend:

```txt id="jlwm30"
internal/
├── handlers/
├── services/
├── repositories/
├── models/
├── database/
```

---

# Important Beginner Insight

Repository Pattern is NOT about:

* reducing SQL
* making things shorter

It’s about:

```txt id="jlwm31"
maintainability
organization
separation of concerns
scalability
testability
```

Those are major backend engineering principles.

---

# One Important Warning

Some developers OVER-engineer repositories.

Tiny apps don't always need:

* interfaces everywhere
* 12 abstraction layers
* complicated architecture

For small apps:

* simple repositories are enough

Architecture should match project size.

---

# In Go Specifically

Go developers often prefer:

* simpler architectures
* fewer abstractions
* explicit code

So Go repository patterns are usually lighter than Java/C# enterprise patterns.

That’s an important ecosystem difference.

---

# Final Mental Model

Think of repository pattern like this:

```txt id="jlwm32"
Application says:
   "Give me user"

Repository says:
   "I'll handle the SQL"
```

The app focuses on:

* behavior
* workflows
* business rules

The repository focuses on:

* persistence
* queries
* database interaction

That separation is the heart of the Repository Pattern.


