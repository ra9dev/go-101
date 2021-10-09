package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func likesGenerator(user *User) chan<- string

const numWrites = 5

func boringGenerator(msg string) <-chan string {
	ch := make(chan string)

	go func() {
		for i := 0; i < numWrites; i++ {
			ch <- fmt.Sprintf("%s - %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(ch)
	}()

	return ch
}

func main() {
	genCh := boringGenerator("boring")

	for msg := range genCh { // мы блочимся на этой строчке до закрытия канала
		fmt.Println(msg)
	}

	fmt.Println("Too boring...")
}
