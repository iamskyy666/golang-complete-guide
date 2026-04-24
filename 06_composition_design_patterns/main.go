package main

import (
	"errors"
	"fmt"
)

// 📂 06_composition_design_patterns
// Proj: Bank Account Management

type Account struct {
	AccNumber string
	Balance float64
	AccOwner string
}

type SavingsAcc struct {
	Account // Embed Acc.
	InterestRate float64 // e.g., 0.2 for 2%
}

type OverdraftAccount struct {
	Account // Embedded
	OverdraftLimit float64
}

//! methods( )

func (acc *Account) Deposit(amt float64)error{
	if amt<=0{
		return errors.New("Deposit amount must be positive!")
	}
	acc.Balance+=amt
	fmt.Printf("✅ Deposited $%.2f to %s. Curr. Balance: $%.2f\n",amt,acc.AccNumber,acc.Balance)
	return nil
}

func (acc *Account) Withdraw(amt float64)error{
	if amt<=0{
		return errors.New("Withdrawal amount must be positive!")
	}
	if amt > acc.Balance {
	return errors.New("insufficient balance")
 }
	acc.Balance-=amt
	fmt.Printf("☑️ Withdrew $%.2f to %s. Curr Balance: $%.2f\n",amt,acc.AccNumber,acc.Balance)
	return nil
}

func (acc *Account) GetBlance()float64{
	return acc.Balance
}

func (sa *SavingsAcc) AddInterest(){
	interest := sa.Balance * sa.InterestRate
	fmt.Printf("Adding interest $%.2f to savings-acc. %s\n",interest, sa.AccNumber)
	err:=sa.Deposit(interest) // embedded f(x) --> promoted
	if err != nil {
		fmt.Printf("ERROR depositing $%.2f to savings-acc. %v\n",interest,err.Error())
	}
}

func (oa *OverdraftAccount) Withdraw(amt float64)error{
	if amt<=0{
		return errors.New("Withdrawal amount must be positive!")
	}
	if (oa.Balance + oa.OverdraftLimit)<amt{
		return fmt.Errorf("Withdrawal of $%.2f exceeds overdraft limit for %s. Available including overdraft: $%.2f\n",amt,oa.AccNumber,oa.Balance+oa.OverdraftLimit)
	}
	oa.Balance-=amt // Balance can go negative
	fmt.Printf("☑️ Withdrew $%.2f from overdraft acc. -  %s. Curr Balance: $%.2f\n",amt,oa.AccNumber,oa.Balance)
	return nil
}

//! Stringer-Interface implementation
func (acc *Account) String()string{
	return fmt.Sprintf("Account [%s], Acc. Holder: %s, Balance:$%.2f", acc.AccNumber, acc.AccOwner, acc.Balance)
}

func main() {
	saveAcc := &SavingsAcc{
	Account: Account{
		AccNumber: "SAV2001",
		AccOwner:  "Ananya Sen",
		Balance:   10000,
	},
	InterestRate: 0.02, // 2%
}

fmt.Println("----------------- Savings Acc. Ops. -------------------")
fmt.Println(saveAcc.Account.String())

err:=saveAcc.Deposit(200.00)
if err != nil {
	fmt.Printf("ERROR depositing $%.2f to savings account. %v\n",200.00,err.Error())
}

saveAcc.AddInterest()

err=saveAcc.Withdraw(50.00)
if err != nil {
	fmt.Println("ERROR:",err.Error())
}
fmt.Println("💰 Final Savings Details:",saveAcc.Account.String())

ovdAcc := &OverdraftAccount{
	Account: Account{
		AccNumber: "OD3001",
		AccOwner:  "Arjun Roy",
		Balance:   2000,
	},
	OverdraftLimit: 3000,
}

fmt.Println("----------------- Overdraft Acc. Ops. --------------------")
fmt.Println(ovdAcc.Account.String())
err=ovdAcc.Deposit(50)
if err != nil {
	fmt.Println("ERROR:",err.Error())
}

err=ovdAcc.Withdraw(100)
if err != nil {
	fmt.Println("ERROR:",err.Error())
}

fmt.Println("💰 Final Overdraft Details:",ovdAcc.Account.String())
}

// $ go run main.go
// ----------------- Savings Acc. Ops. -------------------
// Account [SAV2001], Acc. Holder: Ananya Sen, Balance:$10000.00
// ✅ Deposited $200.00 to SAV2001. Curr. Balance: $10200.00
// Adding interest $204.00 to savings-acc. SAV2001
// ✅ Deposited $204.00 to SAV2001. Curr. Balance: $10404.00
// ☑️ Withdrew $50.00 to SAV2001. Curr Balance: $10354.00
// 💰 Final Savings Details: Account [SAV2001], Acc. Holder: Ananya Sen, Balance:$10354.00
// ----------------- Overdraft Acc. Ops. --------------------
// Account [OD3001], Acc. Holder: Arjun Roy, Balance:$2000.00
// ✅ Deposited $50.00 to OD3001. Curr. Balance: $2050.00
// ☑️ Withdrew $100.00 from overdraft acc. -  OD3001. Curr Balance: $1950.00
// 💰 Final Overdraft Details: Account [OD3001], Acc. Holder: Arjun Roy, Balance:$1950.00

