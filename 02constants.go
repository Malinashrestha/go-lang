package main
import "fmt"

//constant  can also be declared here 

const age  = 3.1415

func  main(){

fmt.Println(age)
const name string = "GOLANG"
//name = "javascrpit" = cannot used this  because of constant
fmt.Println(name)

var same  =  "sexy lambo"
 same = "ROLLS"
fmt.Println(same)

//MULTIPLE CONSTANT CAN BE USED LIKE THIS 
const(
sugar = 110
oil = 320
)

fmt.Println(sugar)
fmt.Println(oil)

}
