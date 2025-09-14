package main
import "fmt"

// for --> only construct in  go for  looping
func main(){

//while loop using  for loop
a := 1
for a <=3{
fmt.Println(a)
a = a+1
}

//INFINTE LOOP
/*
for{
fmt.Println("9")

}
*/

// NOW CLASSIC FOR LOOP

for i:=0;i<5;i++{
//break
if i ==2{
continue
}
fmt.Println(i)

}
// NEW  FEATURE OF RANGE IN GO 1.22

for j:= range 5{
fmt.Println(j)
}



}
