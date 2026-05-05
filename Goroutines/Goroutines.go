package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 13; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func smessage(s string) {
	for i := 0; i < 13; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}


func sum(x, y, z int){
	for i:=0; i<13; i++{
		time.Sleep(100* time.Millisecond)
		fmt.Println(x, y, z)
		fmt.Println(x + y + z)
	}
}
func main() {
	/*
	   A goroutine is a lightweight thread managed by the Go runtime.
	   The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.
	   Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides useful primitives, although you won't need them much in Go as there are other primitives.
	*/
	go say("Hello")
	smessage("How Are You")
	say("Rajan")
	sum(12, 100,12)
}
