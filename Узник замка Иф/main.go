package main

import (
	"fmt"
)

func main() {
	var A, B, C, D, E int
	fmt.Scan(&A, &B, &C, &D, &E)

	// Проверяем все комбинации размеров кирпича
	if (A <= D && B <= E) || (A <= E && B <= D) ||
		(A <= D && C <= E) || (A <= E && C <= D) ||
		(B <= D && C <= E) || (B <= E && C <= D) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
