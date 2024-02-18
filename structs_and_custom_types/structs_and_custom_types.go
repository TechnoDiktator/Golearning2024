package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string , r *bufio.Reader)(string , error){
	fmt.Print(prompt)
	input,err := r.ReadString('\n')

	return strings.TrimSpace(input) , err
}

func createBill() bill{
	reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Create a new bill name:")
	// name , _ := reader.ReadString('\n')
	
	//the readstring returns the readvalue from the std input as awell 
	//as the error
	//name = strings.TrimSpace(name)
	name , _ := getInput("Create a new bill name:" , reader)
	b := newBill(name)
	fmt.Println("Creted the bill - " , b.name)
	return b
}

func promptOptions(b bill){
	reader := bufio.NewReader(os.Stdin)
	opt,_ := getInput("Choose an option (a - add item , s - save the bill , t - add tip) " , reader) 
	//fmt.Println(opt)
	switch opt{
	case "a" : 
		fmt.Println("you chose a ")
		name , _ := getInput("Item name " , reader)
		price , _ := getInput("Item price:" , reader)
		
		//typecasting in go 
		p,err:=strconv.ParseFloat(price , 64)
		if err != nil {
			fmt.Println("The price must be a number !!")
			promptOptions(b)
		}
		b.addItem(name , p)

		fmt.Println(name,price)

	case "t":
		fmt.Println("you chose t")

		tip , _ := getInput("Enter the tip in amount of $" , reader)
		fmt.Println(tip)

		//typecasting in go 
		p,err:=strconv.ParseFloat(tip , 64)
		if err != nil {
			fmt.Println("The tip must be a number !!")
			promptOptions(b)
		}
		b.addItem(tip , p)

	case "s":
		fmt.Println("you chose s , we will save the bill now"  , b)
		b.save()
		fmt.Println("you saved the file! ," , b.name)
	default:
		fmt.Print("plase only select the valid option !!!!")
		promptOptions(b)
	}
}


func main(){
	// mybill := newBill("mario")
	// fmt.Println(mybill.format())

	// mybill.updateTip(10)
	// fmt.Println(mybill.format())
	mybill := createBill() 
	promptOptions(mybill)
	fmt.Println(mybill)
}

















