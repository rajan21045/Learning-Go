package main
import ("fmt")
const name = "Rajan"
const address = "Gaindakot - 04"
const (
	collegeName = "Lumbini I.C.T Campus"
	semester = "Fifth"
	rollNo = 20
	isStudent = true
)
func main(){
	fmt.Print(name)
	fmt.Print(address)

	fmt.Println()

	fmt.Println(collegeName)
	fmt.Println(semester)
	fmt.Println(rollNo)
	fmt.Println(isStudent)
}