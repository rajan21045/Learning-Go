package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	/*
		Q. Create and Print Slice
			- Create a slice of integers [1, 2, 3, 4, 5]
			- Print length and capacity
	*/
	mySlice := []int{1, 2, 3, 4, 5}
	var size int = len(mySlice)
	var capacity int = cap(mySlice)
	fmt.Printf("The Size/Length of Slice Name 'mySlice' Is %d", size)
	fmt.Println()
	fmt.Printf("The Capacity of Slice Name 'mySlice' Is %d", capacity)

	/*
		Q. Slice from Array
			- Given an array [10, 20, 30, 40, 50], create a slice containing only middle elements
	*/
	fmt.Println()
	myArr := [5]int{10, 20, 30, 40, 50}
	mySlice1 := myArr[1:4]
	fmt.Println()
	fmt.Println("The middle values of array name myArr are:")
	fmt.Println(mySlice1)

	/*
		Q. Append Elements
			- Start with an empty slice, append 5 numbers, print after each append
	*/
	fmt.Println()
	num := []int{}
	num = append(num, 1)
	fmt.Println(num)
	num = append(num, 3)
	fmt.Println(num)
	num = append(num, 2)
	fmt.Println(num)
	num = append(num, 4)
	fmt.Println(num)
	num = append(num, 5)
	fmt.Println(num)
	//Sorting The Slices
	sort.Ints(num)
	fmt.Println(num)

	/*
		Q. Slice Copy
			- Copy one slice into another and verify both are independent
	*/
	num2 := num
	fmt.Println(num2)

	/*
		Q. Remove Element
			Remove element at index i from a slice
			(Hint: use append with slicing)
	*/
	fmt.Println()
	myElement := []int{1, 2, 3, 4, 5}
	i := 2
	fmt.Println(myElement)
	myElement = append(myElement[:i], myElement[i+1:]...)
	fmt.Println(myElement)

	/*
		Q. Reverse a Slice
			- Reverse a slice without using extra space
	*/
	fmt.Println()
	nums3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("The Original Slice Named 'nums3' Is %d.", nums3)
	fmt.Println()
	slices.Reverse(nums3)
	fmt.Printf("The Reverse Slice Named 'nums3' Is %d.", nums3)

	/*
		Q. Check if Element Exists
			Write a function that returns true if a number exists in a slice
	*/
	fmt.Println()
	myNumber := []int{1, 8, 12, 5, 67, 44}
	number := 12
	var isExit bool = slices.Contains(myNumber, number)
	fmt.Println()
	fmt.Printf("Do The Number(%d) Exits In The Original Slice?: %b", number, isExit)

	/*
		Q. Find Max and Min
			Return max and min values from a slice
	*/
	fmt.Println()
	minNumber := slices.Min(myNumber)
	maxNumber := slices.Max(myNumber)
	fmt.Println()
	fmt.Printf("The Max Number In Slice Named 'myNumber' Is %d And Min Number In Slice Named 'myNumber' Is %d.", maxNumber, minNumber)

	/*
		Q. Merge Two Slices
			- Combine two slices into one
	*/
	fmt.Println()
	newSlice := slices.Concat(myNumber, nums3)
	fmt.Println()
	fmt.Println(newSlice)
	slices.Sort(newSlice)
	fmt.Println(newSlice)

	/*
		Slice Capacity Growth
			- Track how capacity changes when appending elements in a loop
	*/
		var s []int
		prevCap := cap(s)
		for i := 1; i <= 20; i++ {
			s = append(s, i)
			// Print only when capacity changes
			if cap(s) != prevCap {
				fmt.Printf("After adding %d → len=%d cap=%d\n", i, len(s), cap(s))
				prevCap = cap(s)
			}
		}
}
