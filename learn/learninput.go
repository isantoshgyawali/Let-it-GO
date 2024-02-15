package main

import (
	"fmt"
)

func insertMap() {
	var input int
	m := make(map[string]int)

	fmt.Printf("\nEnter the no. of items to insert: ")
	fmt.Scanf("%d", &input)

	var keys string
	var values int
	for i := 0; i < input; i++ {
		fmt.Printf("\nEnter the %v.key :", i+1)
		fmt.Scanf("%s", &keys)
		fmt.Printf("Enter the value for %v : ", keys)
		fmt.Scanf("%d", &values)

		m[keys] = values
	}
	fmt.Printf("%v\n", m)
}
