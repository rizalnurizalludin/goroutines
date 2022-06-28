package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWord() {
	fmt.Println("Hello Word")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWord()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}
