package main
/**** 
Using composite literal : 
create an Array which holds 5 values of type int assign values  to each index  position 
Range over  the array and print  the values  out  print out the value and the type 
*/
import "fmt"



func main(){
 x:=[5]int{}
 for i:=0; i<5; i++ {
	x[i] = i
 }

 for i,v := range x {
	fmt.Printf("%v - %T - %v \n" , v , v , i );
 } 
}