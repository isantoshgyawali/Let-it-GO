package main
import (
	"fmt"
)

/**
* They are like objects in js to some extent
* keys could be of multiple different types : string, int ,float etc
* but all the keys in the single maps should have the same type and so do the values
**/
func usingMap() {
	// [string] = key types ; []float64 = value types
	mapMenu := map[string]float64{
		"one":   10.203,
		"two":   12.487,
		"three": 23.858,
	}

	fmt.Println("\n", mapMenu)
	fmt.Println("value with key one:", mapMenu["one"])

	for key, value := range mapMenu {
		fmt.Println(key, "-", value)
	}
}
