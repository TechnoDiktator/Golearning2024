package main
import "fmt"

func main(){
	
	age:=35
	name:="tarang"

	fmt.Print("Hello \n")
	fmt.Print("World \n")

	//Println -> for next line
	fmt.Println("Hello")
	fmt.Println("Ninhjas")

	fmt.Println("My age is :" , age )
	fmt.Println("My name is :"  , name)

	fmt.Println("my age is"  ,age , "and my naem is " , name)
	
	//(formatted string) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n" , age , name)
	fmt.Printf("my age is %q and my naem is %q \n",age , name)

	fmt.Printf("age is of type %T and name is of type %T" ,age , name)
	fmt.Printf("you scored %2f points! \n" , 255.5555)

	// Sprintf (save formatted strings in a avariable )
	var str = fmt.Sprintf("my age is %v and my name is %v \n" , age , name)
	fmt.Println("The saved string is :"  , str)

	


}