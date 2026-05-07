/*
In Go language, a channel is a medium through which a goroutine communicates with another goroutine and this communication is lock-free. Or in other words, a channel is a technique which allows to let one goroutine to send data to another goroutine.
By default channel is bidirectional, means the goroutines can send or receive data through the same channel.
*/

/*
In Go language, a channel is created using chan keyword and it can only transfer data of the same type, different types of data are not allowed to transport from the same channel. Syntax:

	var Channel_name chan Type

You can also create a channel using make() function using a shorthand declaration. Syntax:

	channel_name:= make(chan Type)
*/
package main

import (
	"fmt"
)

func main() {
	// Go program to illustrate
	// how to create a channel

	// Creating a channel
	// Using var keyword
	var a chan int
	fmt.Println("Value of the channel: ", a)
	fmt.Printf("Type of the channel: %T ", a)

	fmt.Println()
	// Creating a channel using make() function
	mychannel := make(chan int)
	fmt.Println("Value of the channel: ", mychannel)
	fmt.Printf("Type of the channel: %T ", mychannel)

	fmt.Println()
	ch := make(chan int)
	go myfunc(ch)
	ch <- 25
	/*
	   Send and Receive Data From a Channel
	   In Go language, channel work with two principal operations one is sending and another one is receiving, both the operations collectively known as communication. And the direction of <- operator indicates whether the data is received or send. In the channel, the send and receive operation block until another side is not ready by default. It allows goroutine to synchronize with each other without explicit locks or condition variables.
	   Send operation: The send operation is used to send data from one goroutine to another goroutine with the help of a channel. Values like int, float64, and bool can safe and easy to send through a channel because they are copied so there is no risk of accidental concurrent access of the same value. Similarly, strings are also safe to transfer because they are immutable. But for sending pointers or reference like a slice, map, etc. through a channel are not safe because the value of pointers or reference may change by sending goroutine or by the receiving goroutine at the same time and the result is unpredicted. So, when you use pointers or references in the channel you must make sure that they can only access by the one goroutine at a time.
	   Mychannel <- element
	   The above statement indicates that the data(element) send to the channel(Mychannel) with the help of a <- operator.
	   Receive operation: The receive operation is used to receive the data sent by the send operator.
	   element := <-Mychannel
	   The above statement indicates that the element receives data from the channel(Mychannel). If the result of the received statement is not going to use is also a valid statement. You can also write a receive statement as:
	   <-Mychannel

	*/

	// Creating a channel
	c := make(chan string)

	// calling Goroutine
	go myFun(c)

	// When the value of ok is
	// set to true means the
	// channel is open and it
	// can send or receive data
	// When the value of ok is set to
	// false means the channel is closed
	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}

	//Length of the Channel: In channel, you can find the length of the channel using len() function. 
	// Here, the length indicates the number of value queued in the channel buffer. Example:
	mychnl := make(chan string, 4)
    mychnl <- "GFG"
    mychnl <- "gfg"
    mychnl <- "Geeks"
    mychnl <- "GeeksforGeeks"
	fmt.Println("Length of the channel is: ", len(mychnl))

	//Capacity of the Channel: In channel, you can find the capacity of the channel using cap() function. 
	// Here, the capacity indicates the size of the buffer. Example:
	fmt.Println("Capacity of the channel is: ", cap(mychnl))
}
func myfunc(ch chan int) {

	fmt.Println(234 + <-ch)
}

// Function
func myFun(mychnl chan string) {

	for v := 0; v < 4; v++ {
		mychnl <- "Rajan Poudel"
	}
	close(mychnl)
}
