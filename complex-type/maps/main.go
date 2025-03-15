package main

import (
	"fmt"
	"sort"
)

func main() {
	districts := make(map[string]string)
	districts["DH"] = "Dhaka"
	districts["CH"] = "Chittagaon"
	districts["RJ"] = "Rajshahi"
	districts["TH"] = "Thakurgaon"
	fmt.Println(districts)

	districts["DIN"] = "Dinajpur"
	fmt.Println(districts)

	delete(districts, "TH")
	fmt.Println(districts)

	keys := make([]string, len(districts))
	i := 0
	for k := range districts {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	fmt.Println("\nSorted")

	for i := range keys {
		fmt.Println(districts[keys[i]])
	}
}