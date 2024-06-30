package main
/*** 
Follow these steps :
starting with the slice 
 x:=[]int{42,43,44,45,46,47,48,49,50,51}
 Append  to that slices this values 
 52
 in one STATEMENT append  to that slice  these valeus 
 53,54,55,

 print out the slices 
 append to the slices 
 y:= []int{56,57,58,59,60}
 printout the slices 
*/
import "fmt"


func main(){
	x:=[]int{42,43,44,45,46,47,48,49,50,51}
	fmt.Println(x)

	x= append(x,52)
    fmt.Println(x)

}