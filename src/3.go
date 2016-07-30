package main

import "fmt"

func main() {

	sum := 0
	x := 1
	y := 2

	for true {
		z := y + x
		x = y
		y = z

		if x%2 == 0 {
			sum += x
		}

		if y > 4000000 {
			fmt.Println(sum)
			return
		}
	}

	fmt.Println(sum)

}
