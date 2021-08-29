package main

import (
	"fmt"
)

func main() {
	var arrayList = [...]int{1, 2, 3, 4, 5}
	for v, key := range arrayList {
		fmt.Println(key, v)
	}
}
