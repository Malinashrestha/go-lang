package main
import ("fmt"
 "time"
 )
func main(){
// WE USED SWTICH FOR REPLACING IF/ELSE CONDITION BASICALLY WHEN THERE IS MULTIPLE  CASES

/*
// SIMPLE SWITCH
i:=5

switch i {

case 1:
fmt.Println("one")

case 2:
fmt.Println("two")

case 3:
fmt.Println("Three")	

default:
fmt.Println("other")
}
*/

//Multiple condition  switch case
switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
	fmt.Println("it's weekend")
default:
	fmt.Println("it's workday")
}	
}
