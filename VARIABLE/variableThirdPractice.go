package main

import ("fmt")

func main() {
	accountAge := 2.6
	accountAgeInt := int(accountAge)
	fmt.Println("Account age in years:", accountAge)
	fmt.Println("Account age in days:", accountAgeInt)
	temperature := 98
	temperatureFloat := float64(temperature)
	fmt.Println("Temperature in Fahrenheit:", temperature)
	fmt.Println("Temperature in Celsius:", temperatureFloat)
}
