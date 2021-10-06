package main

import "fmt"

func Average(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return int(total / len(s))
}

func main() {
	fmt.Println(Average([]int{1, 2, 3, 4, 5, 6, 7}))
}
