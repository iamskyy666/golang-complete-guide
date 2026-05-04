package main

import (
	"fmt"
	"sync"
	"time"
)

// 📂 09_concurrency
// Understaning Mutex

type BankAcc struct {
	Balance int
	Mu sync.Mutex
}

func (b *BankAcc) Deposit(amt int){
	b.Mu.Lock()
	defer b.Mu.Unlock()

	b.Balance+=amt

	fmt.Println("Deposit:",amt)
}

func (b *BankAcc) Withdraw(amt int){
	b.Mu.Lock()
	defer b.Mu.Unlock()

	if b.Balance < amt {
		fmt.Println("Cannot withdraw that much amt:",amt)
		return
	}

		b.Balance-=amt

	fmt.Println("Withdraw:",amt)
}

func (b *BankAcc) BalanceInfo()int{
	b.Mu.Lock()
	defer b.Mu.Unlock()

	return b.Balance
}


func main() {
	var wg sync.WaitGroup
	var acc = &BankAcc{Balance: 200}

	for i:=range 10 {
		wg.Add(1)

		go func(amt int) {
			defer wg.Done()
			time.Sleep(time.Duration(amt) * time.Millisecond)
			acc.Deposit(amt)

		}(i+1)
	}



	wg.Wait()

	fmt.Println("Account:",acc.Balance)
}

// $ go run main.go
// Deposit: 1
// Deposit: 2
// Deposit: 3
// Deposit: 4
// Deposit: 5
// Deposit: 6
// Deposit: 8
// Deposit: 7
// Deposit: 9
// Deposit: 10
// Account: 255