package main
import ("fmt")
const PI = 3.14
const A int = 1
const (
	name = "Rajan"
	age = 21
)
func main(){
	/*
	Go Constants
	If a variable should have a fixed value that cannot be changed, you can use the const keyword.

	The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.

	Syntax
	const CONSTNAME type = value
	Note: The value of a constant must be assigned when you declare it.
	*/
	fmt.Println(PI)
	fmt.Println(A)
	fmt.Println(name)
	fmt.Println(age)
}