package main

/***
Using Composite literal :
Create  a SLICE of  Type INT 
Assign these  10 values 
42-51
Range over the slice and print the values out  
*/
import "fmt"


func main(){
 x:=[]int{42,43,44,45,46,47,48,49,50,51}
 for i, v := range x {
	fmt.Printf("%v \t %T \t  %v\n" , v , v , i)
 }

}