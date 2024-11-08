package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	//b2 - 4ac

	var a, b, c int
	fmt.Scan(&a, &b, &c)

	d := b*b - 4*a*c
	var sliceRoots []float64

	if d > 0 {
		x1 := (-float64(b) + math.Sqrt(float64(d))) / (2 * float64(a))
		x2 := (-float64(b) - math.Sqrt(float64(d))) / (2 * float64(a))
		sliceRoots = append(sliceRoots, x1, x2)
		sort.Float64s(sliceRoots)
		fmt.Println(2)
		fmt.Printf("%.6f %.6f\n", sliceRoots[0], sliceRoots[1])
	} else if d == 0 {
		x := -float64(b) / (2 * float64(a))
		sliceRoots = append(sliceRoots, x)
		fmt.Println(1)
		fmt.Printf("%.6f\n", sliceRoots[0])
	} else {
		fmt.Println(0)
	}

}
