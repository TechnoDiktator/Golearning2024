package main
import(
	"fmt"
)


func main(){
	
	//while loop in go looks like a for loop
	// x:=0
	// for x < 5{
	// 	fmt.Println("value of x is: ", x)
	// 	x++
	// }
	
	//for loop in go looks like this 
	// for i:=0 ; i<5 ; i++ {
	// 	fmt.Println("Value of i is:" , i)
	// }

	names := []string{"mario" , "luigi" , "yoshi" ,  "peach"}

	for i:= 0 ; i<len(names) ; i++{
		fmt.Println(" " , names[i])
	}

	//basically a for in 
	// for index , value := range names {
	// 	fmt.Printf("The positiion at index %v is %v \n" , index , value)
	// }


	//in go you have to use every variable that you declare
	for _,value := range names{
		fmt.Printf(" The value is %v \n" , value)
	}
}





















