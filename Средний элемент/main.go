package main

import (
	"fmt"
	"sort"
)

func main() {

	var a, b, c int
	fmt.Scan(&a, &b, &c)

	number := []int{a, b, c}

	sort.Ints(number)
	fmt.Println(number[1])
}
