package main

import (
	"fmt"
)

func main() {
	/*
		Go Slices
		Slices are similar to arrays, but are more powerful and flexible.

		Like arrays, slices are also used to store multiple values of the same type in a single variable.

		However, unlike arrays, the length of a slice can grow and shrink as you see fit.

		In Go, there are several ways to create a slice:

		Using the []datatype{values} format
		Create a slice from an array
		Using the make() function
		Create a Slice With []datatype{values}
		Syntax
		slice_name := []datatype{values}
		A common way of declaring a slice is like this:

		myslice := []int{}
		The code above declares an empty slice of 0 length and 0 capacity.

		To initialize the slice during declaration, use this:
	*/

	myslice := []int{1, 2, 3}
	fmt.Println(cap(myslice))
	fmt.Println(len(myslice))

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
	/*
		In the example above, we see that in the first slice (myslice1), the actual elements are not specified, so both the length and capacity of the slice will be zero.
		In the second slice (myslice2), the elements are specified, and both length and capacity is equal to the number of actual elements specified.
	*/

	/*
		Create a Slice From an Array
		You can create a slice by slicing an array:

		Syntax
		var myarray = [length]datatype{values} // An array
		myslice := myarray[start:end] // A slice made from the array
		Example
		This example shows how to create a slice from an array:
	*/
	myarr1 := [5]int{1, 2, 12, 13, 55}
	myslice3 := myarr1[2:4]

	fmt.Printf("My Slice: %v\n", myslice3)
	fmt.Printf("My Slice Length: %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))
	//In the example above myslice is a slice with length 2. It is made from arr1 which is an array with length 5.

	arr2 := [6]int{10, 11, 12, 13, 14, 15}
	myslice4 := arr2[2:4]

	fmt.Printf("myslice = %v\n", myslice4)
	fmt.Printf("length = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n", cap(myslice4))
	/*
		In the example above myslice is a slice with length 2. It is made from arr1 which is an array with length 6.
		The slice starts from the third element of the array which has value 12 (remember that array indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.). The slice can grow to the end of the array. This means that the capacity of the slice is 4.
		If myslice started from element 0, the slice capacity would be 6.
	*/

	/*
		Create a Slice With The make() Function
		The make() function can also be used to create a slice.

		Syntax
		slice_name := make([]type, length, capacity)
		Note: If the capacity parameter is not defined, it will be equal to length.

		Example
		This example shows how to create slices using the make() function:
	*/

	myslices := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslices)
	fmt.Printf("length = %d\n", len(myslices))
	fmt.Printf("capacity = %d\n", cap(myslices))

	// with omitted capacity
	myslices1 := make([]int, 5)
	fmt.Printf("myslices1 = %v\n", myslices1)
	fmt.Printf("length = %d\n", len(myslices1))
	fmt.Printf("capacity = %d\n", cap(myslices1))

}
