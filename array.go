package main

import (
	"fmt"
)

//  to take pointer refernce in array we need to add => &
func main() {
	studentsRollNo := []int{1, 2, 3, 4, 5}
	studentId := studentsRollNo
	studentId[1] = 100
	// studentId = append(studentId, 6)
	fmt.Println("studentRoll No => ", studentsRollNo, studentId)
}
