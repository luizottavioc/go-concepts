package main

import (
	"fmt"
	"time"
)

func printNum(c chan <- int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Number added: %d\n", i)
		c <- i
		
		time.Sleep(time.Millisecond * 150)
	}

	close(c)
}

func printLetter(c chan <- string) {
	for i := 'a'; i < 'j'; i++ {
		fmt.Printf("Letter added: %c\n", i)
		c <- string(i)
		
		time.Sleep(time.Millisecond * 100)
	}

	close(c)
}

func main() {
	cn := make(chan int, 10)
	cl := make(chan string)

	cnEnd := make(chan bool)
	clEnd := make(chan bool)

	go printLetter(cl)
	go printNum(cn)

	go func() {
		for n := range cn {
			fmt.Printf("Number read: %d\n", n)
			time.Sleep(time.Millisecond * 350)
		}

		cnEnd <- true
	}()

	go func() {
		for l := range cl {
			fmt.Printf("Letter read: %s\n", string(l))
		}

		clEnd <- true
	}()

	<- cnEnd
	<- clEnd

	fmt.Println()
}