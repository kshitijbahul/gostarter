package main

import (
	"fmt"
)

func main() {
	var a [11]int
	for i := 0; i <= 10; i++ {
		a[i] = i
		fmt.Println(checkEvenOrOdd(a[i]))
	}

}
func checkEvenOrOdd(ele int) string {
	if ele%2 != 0 {
		return "odd"
	}
	return "even"
}
