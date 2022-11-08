package main

import "fmt"

func main() {
	numbers := [3]float64{71.8, 56.2, 99.5}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Println("Average: %0.2f\n", sum/sampleCount)
	fmt.Println("I change the", sum)
	//fmt.Println(sum)
}
