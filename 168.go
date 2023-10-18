// program which illustrates the recover in a goroutine

package main

import (
	"fmt"
	"time"
)

// for the recovery

func handlepanic() {
	if x := recover(); x != nil {
		fmt.Println("RECOVER", x)
	}
}

// here, this panic is not handled by recover funcion because of recover function is not called in the same goroutine in which panic occurs

func myfun1() {
	defer handlepanic()
	fmt.Println("Welcome to the Function1")
	go myfun2()
	time.Sleep(10 * time.Second)
}

func myfun2() {
	fmt.Println("Welcome to Function2")
	panic("Paniced!!")
}

func main() {
	myfun1()
	fmt.Println("The Return successfully from main function")
}
