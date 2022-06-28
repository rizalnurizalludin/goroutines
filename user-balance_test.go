package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Rizal",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Ebon",
		Balance: 1000000,
	}
	go Transfer(&user1, &user2, 100000)

	time.Sleep(10 * time.Second)

	fmt.Println("User ", user1.Name, "Balence ", user1.Balance)
	fmt.Println("User ", user2.Name, "Balence ", user2.Balance)
}
