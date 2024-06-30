package main

import "fmt"

func main() {
	xs := []string{"Alabama", "Alaska", "American Samoa", "Alabama", "Alaska", "American Samoa",
		"Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Colorado", "Connecticut",
		"Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Colorado", "Connecticut",
		"Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Colorado", "Connecticut",
		"Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Colorado", "Connecticut",
		"Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Colorado", "Connecticut", "Colorado", "Connecticut",
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
		"West Virginia", "Wisconsin", "Wyoming"}

	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}

	for i, v := range m {
		fmt.Println(v, i)
	}

}
