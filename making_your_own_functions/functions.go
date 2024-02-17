package main

import (	
	"fmt"
	"math"
	
)

func sayGreeting(n string) {
	fmt.Printf("Good mi=orning %v \n" , n)
}

func sayBye(n string){
	fmt.Printf("Goodbye $v \n" , n)
}

func circleArea(r float64){
	return math.Pi*r*r
}

//maa ki aankh callbacks
func cycleNames(n []string , f func(string)){
	for  _ , v := range n {
		f(v)
	}
}






func main(){
	sayGreeting("TABANG")
	sayGreeting("LUIGI")
	sayBye("Mario")

	cycleNames([]string{"cloud" , "tifa" , "barret"} , sayGreeting)

	a1:= circleArea(1.2)
	a2:= circleArea(2.24)
	fmt.Println(a1 , a2)
}















