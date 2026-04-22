package main
import ("fmt")
func main(){
	var myArray = [4]string{"Rajan", "Prabin", "Raaj"}
	cars := [...]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(myArray)
	fmt.Println(cars)
	
	var prices = [4]int{120, 344, 122, 123}
	fmt.Println(prices[2])
	
	prices[2] = 50
	fmt.Println(prices)
}