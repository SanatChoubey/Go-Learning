package main

import (
	"fmt"
)
//  pointer reference take in map 
func main() {
	var mapData = map[string]string{
		"name": "Sanat",
		"age":  "23",
		"Job":  "Software Engineer",
	}
	mapData2 := mapData
	mapData2["Job"] = "coder!"
	fmt.Println("Map Data", mapData["Job"], mapData2)
}
