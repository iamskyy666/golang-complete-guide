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

```go 
package main

import (
	"context"
	"fmt"
	"time"
)

// 📂 09_concurrency
// 💻 Create a ping-ponger small project

func ping(ctx context.Context, ch chan string){
	for {
		select{
		case <-ctx.Done():
			return
		case ch<-fmt.Sprintf("🟢 ping %v",time.Now()):
			time.Sleep(1 * time.Second)	
		}
	}
}

func pong(ctx context.Context, ch chan string){
	for {
		select{
		case <-ctx.Done():
			return
		case ch <-fmt.Sprintf("🔵 pong %v",time.Now()):
			time.Sleep(1 * time.Second)	
		}
	}
}

func main() {
	ctx,cancel:=context.WithCancel(context.Background())
	defer cancel()

	pingCh:= make(chan string)
	done:= make(chan struct{})

	go ping(ctx,pingCh)
	go pong(ctx,pingCh)

	go func() {
		timeout:=time.After(5 * time.Second)
		for{
			select{
			case <-timeout:
				fmt.Println("Ops. completed ✅")
				close(pingCh)
				done <-struct{}{}
				return
			case msg:= <-pingCh:
				fmt.Println(msg)	
			}
		}
	}()

	<-done
	fmt.Println("Done ☑️")
}

// $ go run main.go
// 🟢 ping 2026-05-02 14:41:53.6709971 +0530 IST m=+0.001432601
// 🔵 pong 2026-05-02 14:41:53.6709971 +0530 IST m=+0.001432601
// 🟢 ping 2026-05-02 14:41:54.6724315 +0530 IST m=+1.002867001
// 🔵 pong 2026-05-02 14:41:54.6724315 +0530 IST m=+1.002867001
// 🔵 pong 2026-05-02 14:41:55.6729275 +0530 IST m=+2.003363001
// 🟢 ping 2026-05-02 14:41:55.6729275 +0530 IST m=+2.003363001
// 🟢 ping 2026-05-02 14:41:56.6733093 +0530 IST m=+3.003744801
// 🔵 pong 2026-05-02 14:41:56.6733093 +0530 IST m=+3.003744801
// 🟢 ping 2026-05-02 14:41:57.6737194 +0530 IST m=+4.004154901
// 🔵 pong 2026-05-02 14:41:57.674064 +0530 IST m=+4.004499501
// Ops. completed ✅
// Done ☑️
```

This is actually a very rich concurrency example.

It demonstrates:

* goroutines
* channels
* `select`
* context cancellation
* cooperative shutdown
* timeout handling
* synchronization
* producer-consumer concurrency
* coordination between multiple goroutines

This is very close to real-world concurrent systems design.

Let's dissect it deeply.

---

# High-Level Architecture

This program creates:

| Component          | Role                        |
| ------------------ | --------------------------- |
| `ping()` goroutine | Producer                    |
| `pong()` goroutine | Producer                    |
| `pingCh` channel   | Shared communication stream |
| monitor goroutine  | Consumer/controller         |
| `context.Context`  | Cancellation system         |
| `done` channel     | Completion synchronization  |

---

# Conceptual Flow

```text id="jlwm40"
ping goroutine ─┐
                ├──> pingCh ───> controller goroutine
pong goroutine ─┘
```

---

# Main Goal of Program

The program:

1. Starts two concurrent producers
2. Both continuously send messages
3. Main controller consumes messages
4. After 5 seconds:

   * stop everything
   * cleanup
   * exit gracefully

---

# Step 1 — Imports

```go id="jlwm41"
import (
	"context"
	"fmt"
	"time"
)
```

---

# `context`

Provides:

* cancellation
* deadlines
* lifecycle propagation

Very important in production Go systems.

---

# `time`

Used for:

* delays
* timestamps
* timeout handling

---

# Step 2 — `ping()` Function

```go id="jlwm42"
func ping(ctx context.Context, ch chan string)
```

---

# Parameters

| Parameter | Purpose              |
| --------- | -------------------- |
| `ctx`     | cancellation control |
| `ch`      | output channel       |

---

# Infinite Loop

```go id="jlwm43"
for {
```

This goroutine runs forever until cancelled.

---

# `select`

```go id="jlwm44"
select {
```

Very important concurrency primitive.

`select` waits for:

```text id="jlwm45"
multiple concurrent communication events
```

---

# Case 1 — Context Cancellation

```go id="jlwm46"
case <-ctx.Done():
	return
```

---

# What Is `ctx.Done()`?

`Done()` returns a channel.

When context cancelled:

```text id="jlwm47"
channel closes
```

---

# Meaning

This line means:

```text id="jlwm48"
"If cancellation signal arrives,
exit goroutine."
```

---

# Extremely Important

This is:

# cooperative cancellation

The goroutine voluntarily stops itself.

---

# Case 2 — Send Message

```go id="jlwm49"
case ch <- fmt.Sprintf("🟢 ping %v", time.Now()):
```

This attempts:

```text id="jlwm50"
send message into channel
```

---

# Important Detail

Because this send is inside `select`:

the send only occurs if:

```text id="jlwm51"
channel operation is ready
```

---

# What Message Looks Like

```text id="jlwm52"
🟢 ping 2026-05-02 ...
```

Includes current timestamp.

---

# Then

```go id="jlwm53"
time.Sleep(1 * time.Second)
```

Pauses producer for 1 second.

---

# Why?

Prevents infinite high-speed message spam.

Without sleep:

```text id="jlwm54"
CPU usage skyrockets
millions of messages generated
```

---

# `pong()` Function

Almost identical.

Only difference:

```go id="jlwm55"
"🔵 pong"
```

instead of:

```go id="jlwm56"
"🟢 ping"
```

---

# Important Design Detail

Both goroutines send into SAME channel:

```go id="jlwm57"
pingCh
```

This creates:

# fan-in pattern

---

# Fan-In Pattern

Multiple producers:

```text id="jlwm58"
merge messages into one stream
```

---

# Step 3 — Main Function

---

# Create Context

```go id="jlwm59"
ctx, cancel := context.WithCancel(context.Background())
```

---

# What Is Happening?

Creates:

| Object   | Purpose                          |
| -------- | -------------------------------- |
| `ctx`    | shared cancellation context      |
| `cancel` | function to trigger cancellation |

---

# `context.Background()`

Root context.

Base parent context.

---

# `WithCancel`

Creates child context with cancellation capability.

---

# Why Important?

Now all goroutines can listen for:

```text id="jlwm60"
shutdown signal
```

---

# defer cancel()

```go id="jlwm61"
defer cancel()
```

Ensures cleanup when main exits.

---

# Step 4 — Create Channels

---

# Main Communication Channel

```go id="jlwm62"
pingCh := make(chan string)
```

---

# Important

This is:

# unbuffered channel

---

# Meaning

Each send requires active receiver.

Strong synchronization.

---

# Completion Signal Channel

```go id="jlwm63"
done := make(chan struct{})
```

---

# Why `struct{}`?

Very important Go idiom.

---

# Empty Struct

```go id="jlwm64"
struct{}
```

occupies:

```text id="jlwm65"
zero bytes
```

---

# Meaning

We do NOT care about data.

Only signal matters.

---

# Common Idiom

```go id="jlwm66"
chan struct{}
```

means:

```text id="jlwm67"
signal-only channel
```

---

# Step 5 — Start Producers

```go id="jlwm68"
go ping(ctx, pingCh)
go pong(ctx, pingCh)
```

Now we have:

| Goroutine | Purpose    |
| --------- | ---------- |
| ping      | producer   |
| pong      | producer   |
| main      | controller |

---

# Step 6 — Start Controller Goroutine

```go id="jlwm69"
go func() {
```

This goroutine:

* receives messages
* handles timeout
* coordinates shutdown

---

# Create Timeout

```go id="jlwm70"
timeout := time.After(5 * time.Second)
```

---

# What Does `time.After()` Do?

Returns channel.

After 5 seconds:

```text id="jlwm71"
channel receives value
```

---

# Important

This creates:

# timeout event source

---

# Infinite Loop

```go id="jlwm72"
for {
```

Controller continuously monitors:

* timeout
* incoming messages

---

# select

```go id="jlwm73"
select {
```

This is the concurrency heart of the program.

---

# Case 1 — Timeout Occurs

```go id="jlwm74"
case <-timeout:
```

After 5 seconds:

timeout channel becomes ready.

---

# Then

```go id="jlwm75"
fmt.Println("Ops. completed ✅")
```

---

# CRITICAL PROBLEMATIC LINE

```go id="jlwm76"
close(pingCh)
```

This is actually dangerous.

---

# Why Dangerous?

Because:

```text id="jlwm77"
multiple goroutines are still sending
```

Specifically:

* `ping()`
* `pong()`

both still executing:

```go id="jlwm78"
ch <- ...
```

---

# Sending On Closed Channel

Causes:

```text id="jlwm79"
panic: send on closed channel
```

---

# Why Didn't Program Crash?

Race timing.

The goroutines likely stopped before next send.

But this code is NOT safe.

---

# Proper Shutdown Should Be

```go id="jlwm80"
cancel()
```

NOT:

```go id="jlwm81"
close(pingCh)
```

because producers own sending side.

---

# Important Channel Rule

# Sender closes channel.

Receivers should NOT close channels they do not own.

---

# Then

```go id="jlwm82"
done <- struct{}{}
```

Signals completion.

---

# Then

```go id="jlwm83"
return
```

Controller goroutine exits.

---

# Case 2 — Receive Message

```go id="jlwm84"
case msg := <-pingCh:
```

Receives message from:

* ping OR pong

---

# Important

Because channel shared:

messages interleave nondeterministically.

---

# Example

Could be:

```text id="jlwm85"
ping
pong
pong
ping
```

Scheduling decides order.

---

# Then

```go id="jlwm86"
fmt.Println(msg)
```

prints message.

---

# Step 7 — Main Waits

```go id="jlwm87"
<-done
```

Main blocks until controller signals completion.

---

# Then

```go id="jlwm88"
fmt.Println("Done ☑️")
```

Program exits.

---

# VERY IMPORTANT CONCURRENCY CONCEPTS

---

# 1. Fan-In Concurrency

Two producers:

```text id="jlwm89"
merge into one communication stream
```

---

# 2. Select-Based Event Loop

Controller goroutine behaves like:

# event loop

similar to:

* Node.js event loop
* OS multiplexers
* network pollers

---

# 3. Cooperative Cancellation

```go id="jlwm90"
ctx.Done()
```

allows graceful goroutine shutdown.

---

# 4. Synchronization Through Channels

Unbuffered channel synchronizes sender and receiver.

---

# 5. Timeout Coordination

```go id="’wini91"
time.After()
```

creates asynchronous timeout event.

---

# 6. Concurrent Nondeterminism

Output order varies.

Concurrency inherently nondeterministic.

---

# IMPORTANT BUG IN CODE

This line is unsafe:

```go id="’wini92"
close(pingCh)
```

because producers still sending.

---

# Correct Version

Should be:

```go id="’wini93"
case <-timeout:
	fmt.Println("Ops. completed ✅")

	cancel()

	done <- struct{}{}
	return
```

---

# Why Better?

Because:

```text id="’wini94"
producers stop themselves gracefully
```

via:

```go id="
```


```go id="jlwm95"
case <-ctx.Done():
	return
```

Then channel does NOT need closing.

---

# Better Production Design

Even more correct version:

```go id="jlwm96"
var wg sync.WaitGroup
```

Wait for producers to fully exit before shutdown.

---

# Why?

Without waiting:

program may terminate while goroutines still running.

---

# Another Important Subtlety

This line:

```go id="’wini97"
case msg := <-pingCh:
```

does NOT check:

```go id="’wini98"
ok
```

If channel closes:

receives return:

```text id="’wini99"
zero value
```

For strings:

```text id="જ્મ્વિ100"
""
```

Could create strange behavior.

---

# Proper Receive

Should be:

```go id="’wini101"
case msg, ok := <-pingCh:
	if !ok {
		return
	}
	fmt.Println(msg)
```

---

# Why This Example Is Valuable

This small program demonstrates many real-world concurrency patterns:

| Concept              | Real-world Usage       |
| -------------------- | ---------------------- |
| Producers            | APIs/services/workers  |
| Shared channel       | event stream           |
| Select               | event multiplexing     |
| Context cancellation | graceful shutdown      |
| Timeout              | request deadlines      |
| Done channel         | lifecycle coordination |
| Fan-in               | aggregation systems    |

---

# Internal Runtime Perspective

At runtime:

---

# Goroutines

Runtime scheduler manages:

* ping goroutine
* pong goroutine
* controller goroutine
* main goroutine

---

# When Channel Blocks

Runtime:

```text id="’wini102"
parks goroutine
runs another
```

Efficient concurrency.

---

# When `select` Executes

Runtime checks readiness of:

* timeout channel
* pingCh receive

and chooses ready case.

---

# Mental Model of Whole System

```text id="’wini103"
           ┌─────────┐
           │ ping() │
           └────┬────┘
                │
                │
                ▼
           ┌─────────┐
           │ pingCh  │
           └─────────┘
                ▲
                │
                │
           ┌────┴────┐
           │ pong()  │
           └─────────┘

                │
                ▼

        Controller Goroutine
        ├── Receives messages
        ├── Handles timeout
        └── Coordinates shutdown
```

---

# Most Important Design Insight

This program is essentially implementing:

# concurrent event-driven architecture

using:

* goroutines
* channels
* select
* context cancellation

This is exactly the kind of architecture used in:

* servers
* streaming systems
* websocket systems
* distributed services
* background workers
* message processors

---

# Final Summary

This program demonstrates:

| Feature            | Purpose                  |
| ------------------ | ------------------------ |
| Goroutines         | Concurrent execution     |
| Shared channel     | Message communication    |
| Fan-in             | Merge producers          |
| Unbuffered channel | Synchronization          |
| Select             | Event multiplexing       |
| Context            | Cancellation propagation |
| Timeout            | Lifecycle control        |
| Done channel       | Shutdown synchronization |
| Sleep              | Rate limiting            |
| Infinite loops     | Long-running workers     |

---

# MOST IMPORTANT TAKEAWAY

The deepest idea in this code is:

> Go concurrency is fundamentally about coordinating independent concurrent activities through communication and cancellation signals.

Not merely:

```text id="’wini104"
"running things in parallel"
```

This example demonstrates how Go builds concurrent systems as:

```text id="’wini105"
communicating event-driven components
```

which is the core philosophy behind Go's concurrency model.

```go
package main

import (
	"fmt"
	"sync"
)

// 📂 09_concurrency
// Understaning Mutex
// race conditions..
// A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.

func main() {

	counter:=0 // critical-section

	var wg sync.WaitGroup
	var mu sync.Mutex

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counter)

	// $ go run main.go
	// 10
}
```
