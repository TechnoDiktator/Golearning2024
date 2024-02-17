package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	greetings := "hello three friends"
	fmt.Println(strings.Contains(greetings , "hell"))
	fmt.Println(strings.ReplaceAll(greetings , "hello" ,  "Teri maa ka"))

	fmt.Println(strings.ToUpper(greetings))

	// the original value is unchanged
	fmt.Println(strings.Index(greetings , "ll"))
	fmt.Println(strings.Split(greetings , " "))

	ages := []int{45,20,35,30,75,50}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages , 30)
	fmt.Println(index)

	names := []string{"yoshi" , "mario" , "peach" , "bowser" , "luigi"}
	sort.Strings(names)

	fmt.Println(sort.SearchStrings(names , "bowser"))
	


}

