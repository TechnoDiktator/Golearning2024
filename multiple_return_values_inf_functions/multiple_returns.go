package main
import (
	"fmt"
	"strings"
	
)


func getInitials(n string) (string , string){
	s := strings.ToUpper(n)
	names := strings.Split(s , "")
	var initials []string

	for _ , v :=range names {
		initials = append(initials , v[:1])
	}
	fmt.Println(initials)
	return  initials[0] , initials[1]

}	


func main(){
	firstname , secondname :=getInitials("Tifa Lockhart")
	fmt.Println(firstname , secondname)
	
}

