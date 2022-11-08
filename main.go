package main

import (
	"fmt"
	"github.com/ohoherror/datafile"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("D:\\xc-work\\study\\go\\test\\src\\readfile\\data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Println("Average:%0.2f\n", sum/sampleCount)
}
