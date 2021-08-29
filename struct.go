package main

import (
	"fmt"
)

func main() {
	type userDetails struct {
		name        string
		age         int16
		phoneNumber []string
	}
	userList := []userDetails{
		{
			name:        "sanat",
			age:         23,
			phoneNumber: []string{"9770003301"},
		},
	}
	fmt.Print("Array of Object => ", userList)
}
