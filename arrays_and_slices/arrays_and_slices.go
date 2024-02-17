package main 
import "fmt"


func main()  {

	//arrays are static in go 
	//just like java
	var ages [3]int = [3]int{1,2,3}

	//shorthand
	var ages2 = [3]int{1,2,3}
	ages2[2] = 234

	fmt.Println(ages , len(ages))
	fmt.Println(ages2 , len(ages2))

	names:= [4]string{"yoshi" , "mario" , "peach" ,  "bowser"}
	fmt.Println(ages , len(ages))
	fmt.Println(names , len(names))


	//slices (use arrays under the hood)
	var scores = []int{100 , 50 , 60}
	scores[2] = 25
	scores = append(scores , 456)
	fmt.Println(scores , len(scores))


	//slice ranges
	rangeOne:= names[1:3]
	fmt.Println(rangeOne)

	rangeTwo:= names[2:]
	fmt.Println(rangeTwo)

	range3 := names[:3]
	fmt.Println(range3)

	rangeOne = append(rangeOne , "koopa")
	names[2] = "ABasa"
	fmt.Println(rangeOne , "   " , names)

	











}


