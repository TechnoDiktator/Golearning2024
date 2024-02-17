package main
import "fmt"

func main() {
	//strings
	var nameOne string = "Heello"
	var nameTwo = "luigi"
	nameThree:= "Rastogi" //you cannot use this shorthand notation outside of a function
	fmt.Println(nameOne , nameTwo , nameThree)

	nameOne = "peach"
	nameThree = "bowser"
	nameFour:="yoshi"
	fmt.Println(nameOne , nameTwo , nameThree)
	fmt.Println(nameFour)

	
	//ints
	var ageOne int  = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne , ageTwo , ageThree)

	//bits and memory
	// var numOne int8 = 215   aukaat ke bahar ka number
	var numTwo  int8= -128
	var numThree uint = 25
	//var numThree uint8 = 256  unsigned int can only go till 255
	var numFour uint16 = 256

	fmt.Println(numTwo , numThree , numFour)
  

	//floats
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 234.11
	scoreThree := 1.5   //will be inferred as float 64

	fmt.Println(scoreOne , scoreTwo , scoreThree)

}
















