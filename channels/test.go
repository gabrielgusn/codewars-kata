package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 3)

	go func() {
		// time.Sleep(2 * time.Second)
		ch <- "Hello from goroutine"
	}()

	fmt.Println("Waiting for message...")
	go func() {
		// msg := <-ch
		// fmt.Println("Received:", msg)
		fmt.Println("Received:", <-ch)
	}()

	go func() {
		// time.Sleep(time.Second)
		teste(ch)
	}()
	go func() {
		// time.Sleep(time.Second)
		teste2(ch)
	}()

	go func() {
		// time.Sleep(2 * time.Second)
		fmt.Println("Received 2:", <-ch)
	}()
	go func() {
		// time.Sleep(2 * time.Second)
		fmt.Println("Received 3:", <-ch)
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Main goroutine done")
}

func teste(ch chan string) {
	ch <- "isso eh um teste"
}

func teste2(ch chan string) {
	ch <- "isso eh um segundo teste"
}
