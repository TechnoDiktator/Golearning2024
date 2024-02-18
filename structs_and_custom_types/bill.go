package main

import (
	"fmt"
	"os"
)
type bill struct {
	name string
	items map[string]float64
	tip float64

}
// make new builds 
func newBill(name string)bill{
	b := bill{
		name: name,
		items: map[string]float64{},
		tip:0,
	}
	return b
}

//format the bill
func (b bill) format()string {
	fs := "Bill breakdown \n"
	var total float64 = 0
	//list items	
	for k,v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n"  , k + ":" , v)
		total +=v
	}
	//total
	fs += fmt.Sprintf("%-25v...$%0.2f" , "total:" , total)
	return fs
}



// update the tip
func (b *bill)updateTip(tip float64){
	//(*b).tip = tip
	//we dont need to do it like as above
	
	//we are passing b 
	b.updateTip(tip)
}

//add an item to the bill
func (b *bill)addItem(name string , price float64){
	 b.items[name] = price
}

func(b *bill)save()string{

	//we converted the object into bytes
	data := []byte(b.format())
	err := os.WriteFile("saved_bills/"+b.name+ ".txt" ,data , 0644)

	if(err !=nil){
		panic(err)
	}	
	fmt.Println("We have saved the bill successfully" , b)

	return ""
}




























