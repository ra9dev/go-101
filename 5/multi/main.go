package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func boringGenerator(msg string) <-chan string {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%s - %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(ch)
	}()

	return ch
}

func fanIn(cs ...<-chan string) <-chan string { // 0
	ch := make(chan string)
	wg := new(sync.WaitGroup) // { count: 0 }

	// for _, c := range cs в области замыкания 0
	for _, c := range cs { // 1
		wg.Add(1) // сount++
		// Когда мы запускаем горутины в цикле, они настолько конкурентно быстро запускаются, что фиксируют для себя последнее значение
		// слайса в цикле

		localC := c // можем зафиксироать в замыкании{} итерации цикла
		go func() { // 2
			defer wg.Done() // count--

			for in := range localC {
				ch <- in
			}
		}()
	}

	go func() { // 1
		wg.Wait() // count > 0, block
		close(ch)
	}()

	return ch
}

func main() {
	fannedInCh := fanIn(boringGenerator("Joe"), boringGenerator("Ann"), boringGenerator("Еркебулан"), boringGenerator("Кирилл"))

	for v := range fannedInCh {
		fmt.Println(v)
	}

	fmt.Println("Still boriing...")
}
