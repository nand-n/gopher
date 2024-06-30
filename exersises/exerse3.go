package main

/***
To delete from slices we use append  along  with sling. for this hands on exercise follow  these steps
Start with slices

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	use append and slicing  to get the values here  which you should  assign  to the variables x  and then print
	[42 43 44 48 49 50 51]
*/
import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}
