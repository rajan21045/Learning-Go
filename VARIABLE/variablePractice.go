package main

import ("fmt")

func main() {
    /*
    Go Variable Types
    In Go, there are different types of variables, for example:

    int- stores integers (whole numbers), such as 123 or -123
    float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
    string - stores text, such as "Hello World". String values are surrounded by double quotes
    bool- stores values with two states: true or false
    */
    var name string = "Rajan Poudel"
    var age int = 20
    var marks float32 = 20.12
    var isPass = false
    fmt.Printf("Hi, %v. Your Age Is %v, Marks Is %v And You are %v.", name, age, marks, isPass)
    fmt.Println()


    //Go Multiple Variable Declaration
    var a, b, c int = 1, 2, 3
    fmt.Printf("%v, %v, %v", a, b, c)
}