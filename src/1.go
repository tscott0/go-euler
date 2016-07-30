package main1

import "fmt"

func main() {

	var multiples []int
	var sum int = 0

	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			multiples = append(multiples, i)
			sum += i
		}
	}

	fmt.Println(multiples)
	fmt.Println(sum)

}
