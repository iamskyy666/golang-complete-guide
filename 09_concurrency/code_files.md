```go
package main

// 📂 09_concurrency
// starting a new groutine with the "go" keyword

import (
	"fmt"
	"time"
)

func sayHello(msg string, delay time.Duration){
	time.Sleep(delay)
	fmt.Println("sayHello f(x):",msg)
}

func main() {
	fmt.Println("First message from main() goroutine")
	sayHello("HELLO CONCURRENCY ⏳",time.Second)
    // ⬇️ blocks this for time.Second
	fmt.Println("Last message from main() goroutine")
}

//O/P
// $ go run main.go
// First message from main() goroutine
// sayHello f(x): HELLO CONCURRENCY ⏳
// Last message from main() goroutine

// soln. - go sayHello("HELLO CONCURRENCY ⏳",time.Second)
// STILL PROBLEMATIC ❌

```
```go 
package main

// 📂 09_concurrency
// starting a new groutine with the "go" keyword

import (
	"fmt"
	"time"
)

func sayHello(msg string, delay time.Duration){
	time.Sleep(delay)
	fmt.Println("sayHello f(x):",msg)
}

func main() {
	fmt.Println("1.First message from main() goroutine")
	fmt.Println("2. Second message from main() goroutine")
	go sayHello("HELLO CONCURRENCY 1 ⏳",time.Second)
	go sayHello("HELLO CONCURRENCY 2 ⏳",time.Second)
	go sayHello("HELLO CONCURRENCY 2 seconds ⏳",2*time.Second)
	go sayHello("HELLO CONCURRENCY 3 seconds ⏳",3*time.Second)
	fmt.Println("Last message from main() goroutine")
	time.Sleep(2 * time.Second)
} 

//O/P // ❌ PROBLEMATIC (solution? -  waitgroup)
// $ go run main.go
// 1.First message from main() goroutine
// 2. Second message from main() goroutine
// Last message from main() goroutine
// sayHello f(x): HELLO CONCURRENCY 2 ⏳
// sayHello f(x): HELLO CONCURRENCY 1 ⏳
```

```go
package main

// 📂 09_concurrency
// understanding wait_groups
// golang memory-model: https://go.dev/ref/mem 🧠

// 💡RULES OF WGs -
// 1. Add?increment outside of goroutines.
// 2. We must decrease/decrement the counter by calling wg.Done() inside the goroutine, not outside.
// 3. We mustn't forget to call wg.Wait() -  if we forget, then nothing will happen/ No effect at all.
// 4. BONUS - Always pass a *reference of a wait-group variable, i nstead of a copy.

import (
	"fmt"
	"sync"
	"time"
)

// 🔵 BASIC WAY

func sayHello(msg string, delay time.Duration, wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(delay)
	fmt.Println("sayHello f(x):",msg)
}

func main() {

	var wg sync.WaitGroup // ready for use
	wg.Add(5) // no. of goroutines

	fmt.Println("1. First message from main() goroutine")
	fmt.Println("2. Second message from main() goroutine")
	go sayHello("HELLO CONCURRENCY 1 - 1s ⏳",time.Second,&wg)
	go sayHello("HELLO CONCURRENCY 2 - 2s ⏳",time.Second,&wg)
	go sayHello("HELLO CONCURRENCY 3 - 2s ⏳",2*time.Second,&wg)
	go sayHello("HELLO CONCURRENCY 4 - 5s ⏳",5*time.Second,&wg)
	go sayHello("HELLO CONCURRENCY 5 - 1s ⏳",time.Second,&wg)
	fmt.Println("Last message from main() goroutine")
	time.Sleep(2 * time.Second)

	wg.Wait()
}

// $ go run main.go
// 1. First message from main() goroutine
// 2. Second message from main() goroutine
// Last message from main() goroutine
// sayHello f(x): HELLO CONCURRENCY 2 - 2s ⏳
// sayHello f(x): HELLO CONCURRENCY 5 - 1s ⏳
// sayHello f(x): HELLO CONCURRENCY 1 - 1s ⏳
// sayHello f(x): HELLO CONCURRENCY 3 - 2s ⏳
// sayHello f(x): HELLO CONCURRENCY 4 - 5s ⏳

// ---------------------------------------------------------------------
// ---------------------------------------------------------------------

// 🟢 ALTERNATE WAY (MORE COMMON)
func sayHello(msg string, delay time.Duration, wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(delay)
	fmt.Println("sayHello f(x):",msg)
}

func main() {

	var wg sync.WaitGroup // ready for use
	totalJobs:=5

	for i := range totalJobs {
	wg.Add(1) // no. of goroutines
	go sayHello(fmt.Sprintf("JOB: %d,",i),time.Second,&wg)
	}
	fmt.Println("Last message from main() goroutine")
	wg.Wait()
} 

// $ go run main.go
// Last message from main() goroutine
// sayHello f(x): JOB: 0,
// sayHello f(x): JOB: 2,
// sayHello f(x): JOB: 1,
// sayHello f(x): JOB: 4,
// sayHello f(x): JOB: 3,

```

```go 
package main

import (
	"fmt"
	"time"
)

// 📂 09_concurrency
// channels

type User struct{
	Name string
	Age int
}

func main() {
	messages:=make(chan string)
	users:=make(chan User) // any type of data

	go func ()  {
		fmt.Println("<- Sending a message to messages chan.")
		// variable <- channel data
		messages<-"Hello from messages channel!"
	}()

	go func ()  {
		fmt.Println("<- Sending data to users chan.")
		// variable <- channel data
		users<-User{
			Name: "Skyy",
			Age: 30,
		}
	}()

	time.Sleep(1*time.Second)
	fmt.Println("About to get data from the channels. =>")

	msg:=<-messages
	fmt.Println("🟢 Messages data:",msg)

	skyy:=<-users
	fmt.Println("🔵 User data:",skyy)
}

// $ go run main.go
// <- Sending data to users chan.
// <- Sending a message to messages chan.
// About to get data from the channels. =>
// 🟢 Messages data: Hello from messages channel!
// 🔵 User data: {Skyy 30}

```
```go 

package main

import (
	"fmt"
)

// 📂 09_concurrency
// Buffered Channels


// Buffer - temporary location, keep something first, process later.
// Buffered Chans.- Keep receiving data, without using/processing immediately.
// Even when the channel is full, we can keep sending data to it (unlike unbuffered channels).
// Channels must be closed.
func main() {
	messages:=make(chan string,3)

	fmt.Println("Sending messages to buffered chan.")
	messages<-"Message 1"
	messages<-"Message 2"
	messages<-"Message 3"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}

// $ go run main.go
// Sending messages to buffered chan.
// Message 1
// Message 2
// Message 3

```

```go 

package main

import "fmt"

// 📂 09_concurrency
// Closing Channels

func main() {
	// double channel
	jobs:=make(chan int,5)
	done:=make(chan bool)

	// start a goroutine
	go func() {
		for {
			r,ok:=<-jobs
			if ok {
				fmt.Println("Got this message!",r)
			}else{
				done<-false
				return
			}
		}
	}()
		for i := 1; i <= 3; i++ {
			jobs<-i
			fmt.Println("Sending..",i)
		}

		close(jobs)
	<-done
}

// $ go run main.go
// Sending.. 1
// Sending.. 2
// Sending.. 3
// Got this message! 1
// Got this message! 2
// Got this message! 3



// 💡 NOW THE SAME, WITH WAIT-GROUPS

// Closing Channels - with WaitGroups

func main() {
	jobs:=make(chan int,5)
	//done:=make(chan bool)

	var wg sync.WaitGroup

	wg.Add(1)

	// start a goroutine
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			r,ok:=<-jobs
			if ok {
				fmt.Println("Got this message!",r)
			}else{
				//done<-false
				fmt.Println("Channel Closed..!")
				return
			}
		}
	}(&wg)

		for i := 1; i <= 3; i++ {
			jobs<-i
			fmt.Println("Sending..",i)
		}

		close(jobs)
	// <-done
	wg.Wait()
}

// $ go run main.go
// Sending.. 1
// Sending.. 2
// Sending.. 3
// Got this message! 1
// Got this message! 2
// Got this message! 3
// Channel Closed..!
```

Let's break this down extremely carefully, because this code demonstrates several foundational Go concurrency concepts simultaneously:

* buffered channels
* goroutines
* channel closing
* detecting closed channels
* synchronization
* wait groups
* producer-consumer pattern

This is actually a classic concurrent systems design pattern.

---

# High-Level Idea

This program has:

| Component          | Role                |
| ------------------ | ------------------- |
| Main goroutine     | Producer            |
| Worker goroutine   | Consumer            |
| `jobs` channel     | Communication queue |
| `done` / WaitGroup | Synchronization     |

---

# PART 1 — VERSION USING `done` CHANNEL

---

# Full Code

```go
package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			r, ok := <-jobs

			if ok {
				fmt.Println("Got this message!", r)
			} else {
				done <- false
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("Sending..", i)
	}

	close(jobs)

	<-done
}
```

---

# Step 1 — Channel Creation

```go
jobs := make(chan int, 5)
```

This creates:

# buffered channel

with:

```text id="0g4g0q"
capacity = 5
```

---

# Meaning

The channel can store:

```text id="c3g82v"
5 integers
```

before blocking sender.

---

# Internally

Conceptually:

```text id="79t6ww"
jobs:
[ ][ ][ ][ ][ ]
```

Initially empty.

---

# Why Buffered Here?

Because producer can send several jobs without waiting for consumer immediately.

This creates:

```text id="tds2o9"
producer-consumer decoupling
```

---

# Next

```go
done := make(chan bool)
```

Unbuffered channel.

Used purely for:

# synchronization signal

NOT data transfer.

---

# Step 2 — Start Worker Goroutine

```go
go func() {
```

Creates concurrent worker.

---

# Important

Now we have:

| Goroutine | Purpose       |
| --------- | ------------- |
| Main      | Produces jobs |
| Worker    | Consumes jobs |

---

# Step 3 — Infinite Loop

```go
for {
```

Worker continuously waits for jobs.

---

# Step 4 — Receiving From Channel

```go
r, ok := <-jobs
```

This is extremely important.

---

# Two-Value Receive

When receiving from channel:

```go
value, ok := <-channel
```

---

# Meaning of `ok`

| Situation                | ok    |
| ------------------------ | ----- |
| Channel open             | true  |
| Channel closed and empty | false |

---

# Very Important

When channel closes:

receivers do NOT instantly fail.

They continue receiving buffered values first.

Only when:

```text id="jlwm11"
channel closed
AND
buffer empty
```

does:

```text id="jlwm12"
ok == false
```

---

# Step 5 — If Channel Open

```go
if ok {
	fmt.Println("Got this message!", r)
}
```

Worker processes received job.

---

# Step 6 — If Channel Closed

```go
else {
	done <- false
	return
}
```

This means:

```text id="jlwm13"
"No more jobs exist."
```

---

# Important Synchronization Event

Worker sends signal:

```go
done <- false
```

to notify main goroutine:

```text id="jlwm14"
"I'm finished."
```

---

# Then

```go
return
```

Worker goroutine exits.

---

# Step 7 — Producer Loop

```go
for i := 1; i <= 3; i++ {
	jobs <- i
	fmt.Println("Sending..", i)
}
```

Producer sends:

```text id="jlwm15"
1
2
3
```

into buffered channel.

---

# Internal Buffer Evolution

Initially:

```text id="jlwm16"
[]
```

After sends:

```text id="jlwm17"
[1]
[1 2]
[1 2 3]
```

---

# Since Buffer Capacity = 5

Producer never blocks.

---

# Important Concurrent Behavior

Meanwhile worker goroutine may already consume items concurrently.

Scheduling determines exact interleaving.

---

# Step 8 — Closing Channel

```go
close(jobs)
```

This is one of the most important concurrency operations in Go.

---

# What `close()` Actually Means

It means:

```text id="jlwm18"
"No more values will ever be sent."
```

NOT:

```text id="jlwm19"
"Destroy channel immediately"
```

---

# Critical Detail

Buffered values remain available.

---

# After Close

Channel state becomes:

| Property                    | State |
| --------------------------- | ----- |
| New sends allowed?          | ❌ No  |
| Remaining receives allowed? | ✅ Yes |
| Buffered data accessible?   | ✅ Yes |

---

# Why Close Channel?

Because otherwise worker loop:

```go
for {
	r, ok := <-jobs
}
```

would block forever waiting for more data.

---

# Step 9 — Main Waits

```go
<-done
```

Main blocks until worker signals completion.

---

# Why Needed?

Without this:

main goroutine might exit early.

Remember:

```text id="jlwm20"
When main exits,
entire program exits.
```

---

# Timeline of Execution

---

# Main

Creates channels.

---

# Main

Launches worker.

---

# Worker

Blocks on:

```go
<-jobs
```

waiting for data.

---

# Main

Sends:

```text id="jlwm21"
1
2
3
```

---

# Worker

Receives and prints them.

---

# Main

Closes channel.

---

# Worker

Eventually receives:

```text id="jlwm22"
ok == false
```

meaning:

```text id="jlwm23"
closed + empty
```

---

# Worker

Signals done.

---

# Main

Unblocks and exits.

---

# Extremely Important Channel Concept

This line:

```go
r, ok := <-jobs
```

is the canonical Go way to detect:

# channel exhaustion

This is fundamental in pipelines and worker systems.

---

# PART 2 — WAITGROUP VERSION

Now the same program rewritten more idiomatically.

---

# Main Difference

Instead of:

```go
done := make(chan bool)
```

we use:

```go
var wg sync.WaitGroup
```

---

# Why Better?

Because WaitGroup is specifically designed for:

```text id="jlwm24"
goroutine lifecycle synchronization
```

---

# Step 1 — Create WaitGroup

```go
var wg sync.WaitGroup
```

Internally contains:

* atomic counter
* semaphore
* synchronization state

---

# Step 2 — Add Worker Count

```go
wg.Add(1)
```

Meaning:

```text id="jlwm25"
"One goroutine must finish."
```

Counter:

```text id="jlwm26"
0 → 1
```

---

# Step 3 — Start Worker

```go
go func(wg *sync.WaitGroup)
```

Pass pointer because WaitGroup must NOT be copied.

---

# Step 4 — defer wg.Done()

```go
defer wg.Done()
```

Critical concurrency pattern.

---

# Meaning

When goroutine exits:

```go
wg.Done()
```

automatically runs.

Equivalent to:

```go
wg.Add(-1)
```

---

# Why defer?

Guarantees cleanup even if:

* return happens
* panic occurs
* error occurs

---

# Step 5 — Same Consumer Loop

```go
r, ok := <-jobs
```

Same logic.

---

# Difference

Instead of:

```go
done <- false
```

we simply:

```go
return
```

because:

```text id="jlwm27"
WaitGroup already tracks completion
```

---

# Step 6 — Main Waits

```go
wg.Wait()
```

Main blocks until:

```text id="jlwm28"
WaitGroup counter becomes zero
```

---

# Internally

When worker exits:

```go
wg.Done()
```

Counter:

```text id="jlwm29"
1 → 0
```

Runtime wakes goroutine blocked on:

```go
wg.Wait()
```

---

# Why WaitGroup Version Is Better

---

# `done` Channel Version

Works, but:

* manual signaling
* extra channel
* more boilerplate
* less scalable

---

# WaitGroup Version

Cleaner:

* purpose-built
* scalable
* idiomatic Go
* clearer intent

---

# Important Concurrency Concepts Demonstrated

This example demonstrates MANY important ideas.

---

# 1. Producer-Consumer Pattern

Main goroutine:

```text id="jlwm30"
producer
```

Worker goroutine:

```text id="jlwm31"
consumer
```

---

# 2. Buffered Queue

```go
make(chan int, 5)
```

acts like bounded concurrent queue.

---

# 3. Coordination Through Channels

Communication occurs safely without shared memory.

---

# 4. Channel Closing

Producer signals:

```text id="jlwm32"
"No more work."
```

---

# 5. Graceful Worker Shutdown

Worker exits cleanly after processing remaining jobs.

---

# 6. Synchronization

Main waits safely for worker completion.

---

# 7. Happens-Before Guarantees

Channel operations synchronize memory visibility.

---

# IMPORTANT Subtle Detail

This worker loop:

```go
for {
	r, ok := <-jobs
}
```

is usually written more idiomatically as:

```go
for r := range jobs {
	fmt.Println(r)
}
```

---

# Why?

Because:

```go
range channel
```

automatically:

* receives values
* stops when channel closes

---

# Equivalent Cleaner Version

```go
go func() {
	defer wg.Done()

	for r := range jobs {
		fmt.Println("Got this message!", r)
	}

	fmt.Println("Channel Closed..!")
}()
```

This is the idiomatic Go approach.

---

# Why Channel Closing Is So Important

Without closing:

worker blocks forever:

```go
<-jobs
```

This creates:

# goroutine leak

---

# Closing Channels Is A Coordination Mechanism

It is NOT merely cleanup.

It semantically communicates:

```text id="jlwm33"
"The stream has ended."
```

---

# Very Important Rule

# Sender closes the channel.

NOT receiver.

---

# Why?

Because sender knows:

```text id="jlwm34"
whether more values may arrive
```

Receiver usually does not.

---

# Final Mental Model

This program is essentially:

```text id="jlwm35"
Producer
   ↓
Buffered Queue (channel)
   ↓
Consumer Worker
   ↓
Completion Synchronization
```

---

# Final Summary

This code demonstrates:

| Concept                 | Purpose                   |
| ----------------------- | ------------------------- |
| Goroutine               | Concurrent worker         |
| Buffered channel        | Job queue                 |
| Channel receive         | Consume work              |
| `ok` value              | Detect channel close      |
| `close(jobs)`           | Signal no more work       |
| `done` channel          | Manual synchronization    |
| WaitGroup               | Idiomatic synchronization |
| Blocking                | Safe coordination         |
| Producer-consumer model | Concurrent architecture   |

---

# Most Important Takeaway

The deepest idea in this example is:

> Channels are not just for passing data.

They also communicate:

* lifecycle
* completion
* coordination
* synchronization
* stream termination

That is why channels are such a powerful abstraction in Go concurrency.
