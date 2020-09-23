package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var philosophers = []string{"Thomas", "Maxime", "Florent", "Romain", "Quentin"}

func pickUp(phName string, forkLeft, forkRight *sync.Mutex) {
	// fmt.Println(phName, "is picking up...")
	forkLeft.Lock()
	forkRight.Lock()
}

func pickDown(phName string, forkLeft, forkRight *sync.Mutex) {
	// fmt.Println(phName, "is picking down...")
	forkLeft.Unlock()
	forkRight.Unlock()
}

func eat(phName string) {
	// fmt.Println(phName, "is eating...")
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
}

func think(phName string) {
	// fmt.Println(phName, "is thinking...")
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func eatWhileHungry(phName string, forkLeft *sync.Mutex, forkRight *sync.Mutex) {
	// fmt.Println(phName, " seated")
	think(phName)
	fmt.Println(phName, "takes fork ", &forkLeft, " and ", &forkRight)
	pickUp(phName, forkLeft, forkRight)
	eat(phName)
	pickDown(phName, forkLeft, forkRight)
	fmt.Println(phName, "put down fork ", &forkLeft, " and ", &forkRight)
	fmt.Println(phName, "left")
	dining.Done()
}

var dining sync.WaitGroup

func main () {
	dining.Add(5)

	forkLeft := &sync.Mutex{}

	for _, n := range philosophers {
		forkRight := &sync.Mutex{}
		go eatWhileHungry(n, forkLeft, forkRight)
		forkLeft = forkRight
	}

	dining.Wait()
	fmt.Println("Everyone ate")
}