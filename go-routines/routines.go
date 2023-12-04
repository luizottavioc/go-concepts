package main

import (
	"fmt"
	"time"
)

func loopFunc(name string, qtd int) {
	for i := 1; i <= qtd; i++ {
		fmt.Printf("Loop '%s' count: %d\n", name, i)
		time.Sleep(time.Millisecond * 350)
	}
}

func main() {
	go loopFunc("a", 5)
	go loopFunc("b", 6)
	go loopFunc("c", 7)
	go loopFunc("d", 8)

	time.Sleep(time.Second * 10)
}
