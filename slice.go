package main

import (
	"fmt"
	"sort"
)

func main() {
	fruitSlice := []string{
		"apple",
		"orange",
		"peach",
		"blueberry",
		"graps",
		"lemon",
		"tomato",
		"mango",
	}
	sort.Strings(fruitSlice)
	newFruit := fruitSlice[2:]
	fmt.Println("fruit Slice", fruitSlice, newFruit)
}
