package main
import (
	"fmt"
)

func main(){
	age:=45
	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 45)

	if age < 30 {
		fmt.Println("age ius less than 30")
	}else if age<40{
		fmt.Println("age is less than  forty")
	}else{
		fmt.Println("age is not less than 45")
	}
	//continue and brak
	names := []string{"mario" , "luigi" , "yoshi" , "peach" , "bowser"}

	for index , value := range names {
		if index == 1{
			fmt.Println("continuing at  pos " , index , value)
			continue
		}
		if index > 2 {
			fmt.Println("breakink at the pos %v" , index)
		}


		fmt.Printf("the value at pos %v is %v" , index , value)
	}
}


































































