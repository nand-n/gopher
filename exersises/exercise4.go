package main

import "fmt"

/**
create a slice  to store the names  of all the states in the united states  of america
use make  and append to do this
Goal : do not  have an array that underlise the slice created  more than  once
Print out
	 the len
	 the cap
	 the values along with their index position with out  using the range clause

	{"Alabama", "Alaska", "American Samoa", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "District Of Columbia", "Federated States Of Micronesia", "Florida", "Georgia", "Guam", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Marshall Islands", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Northern Mariana Islands", "Ohio", "Oklahoma", "Oregon", "Palau", "Pennsylvania", "Puerto Rico", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virgin Islands", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming"}
*/

func main() {
	x := make([]int, 50)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x2 := make([]int, 0, 50)
	fmt.Println(x2)
	fmt.Println(len(x2))
	fmt.Println(cap(x2))

	fmt.Println("-----------------------")
	x = append(x, 98)
	x2 = append(x2, 99)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	fmt.Println(x2)
	fmt.Println(len(x2))
	fmt.Println(cap(x2))

	xs := make([]string, 0, 50)
	xs = append(xs, "Alabama", "Alaska", "American Samoa",
		"Arizona", "Arkansas", "California", "Colorado", "Connecticut",
		"Delaware", "District Of Columbia", "Federated States Of Micronesia",
		"Florida", "Georgia", "Guam", "Hawaii", "Idaho", "Illinois",
		"Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine",
		"Marshall Islands", "Maryland", "Massachusetts", "Michigan",
		"Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska",
		"Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York",
		"North Carolina", "North Dakota", "Northern Mariana Islands",
		"Ohio", "Oklahoma", "Oregon", "Palau", "Pennsylvania", "Puerto Rico",
		"Rhode Island", "South Carolina", "South Dakota", "Tennessee",
		"Texas", "Utah", "Vermont", "Virgin Islands", "Virginia", "Washington",
		"West Virginia", "Wisconsin", "Wyoming",
	)

	fmt.Println(len(xs))
	fmt.Println(cap(xs))
	for i := 0; i < len(xs); i++ {
		println(xs[i])
	}
}
