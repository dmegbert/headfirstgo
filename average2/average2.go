package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	var sum float64 = 0
	for _, arg := range arguments {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += num
	}
	sampleCount := float64(len(arguments))
	fmt.Printf("Average is: %0.2f\n", sum / sampleCount)
	maxNum, err := max(arguments...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Max is %f\n", maxNum)
}

func max(numbers ...string) (float64, error) {
	maxNum := 0.0
	for _, num := range numbers {
		num, err := strconv.ParseFloat(num, 65)
		if err != nil {
			return 0, err
		}
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum, nil
}
