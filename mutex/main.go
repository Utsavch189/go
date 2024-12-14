package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	balance int
	mu      sync.Mutex
}

func (a *Account) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
	time.Sleep(time.Second * 2)
	fmt.Println("Deposit completed")
}

func (a *Account) Withdraw(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance >= amount {
		a.balance -= amount
		time.Sleep(time.Second * 2)
		fmt.Println("Withdrawal completed")
	} else {
		fmt.Println("Insufficient balance")
	}
}

func (a *Account) GetBalance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	fmt.Println("Reading balance...")
	return a.balance
}

func main() {
	account := &Account{}
	var wg sync.WaitGroup

	wg.Add(2) // Ensure we add for both goroutines

	go func() {
		defer wg.Done()
		account.Deposit(100)
	}()

	go func() {
		defer wg.Done()
		account.Withdraw(50)
	}()

	wg.Wait() // Wait for both Deposit and Withdraw to finish

	fmt.Printf("Final Balance: %d\n", account.GetBalance())
}

/*
Methods on the Account
	Deposit(amount int):
		Locks the mutex to prevent concurrent modification of balance.
		Increments the balance by amount.
		Unlocks the mutex when done (using defer ensures unlocking happens even if a panic occurs).

	Withdraw(amount int):
		Locks the mutex.
		Checks if balance is sufficient for the withdrawal. If yes, it deducts the amount from balance.
		Unlocks the mutex.

	GetBalance():
		Locks the mutex to safely read the balance.
		Unlocks the mutex after retrieving the value.
		These methods ensure that operations on balance are atomic, meaning they cannot be interrupted by other goroutines.

Main Function
	Step-by-step Execution:

	Account Initialization:
		account := &Account{} creates a new account with an initial balance of 0.

	WaitGroup Setup:
		A sync.WaitGroup is used to wait for two goroutines (Deposit and Withdraw)
		to complete before printing the final balance.

	Deposit Operation in a Goroutine:
		A goroutine calls account.Deposit(100).
		This locks the mutex, increments the balance by 100, and then unlocks the mutex.

	Withdraw Operation in a Goroutine:
		Another goroutine calls account.Withdraw(50).
		This locks the mutex, checks if the balance is at least 50 (which it is after the deposit),
		and deducts 50. The mutex is then unlocked.

	Waiting for Goroutines:
		The wg.Wait() ensures the main goroutine waits for the Deposit and Withdraw operations
		to complete before proceeding.

	Getting Final Balance:
		After the goroutines finish, account.GetBalance()
		is called to safely read the balance and print it.
*/
