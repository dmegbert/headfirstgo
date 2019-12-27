package main

import (
	"fmt"
	datafile2 "github.com/dmegbert/headfirstgo/datafile"
	encaps2 "github.com/dmegbert/headfirstgo/encaps"
	"log"
)

func main() {
	numbers, err := datafile2.GetFloats("/Users/degbert/go/bin/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, num := range numbers {
		sum += num
	}
	avg := sum / float64(len(numbers))
	fmt.Printf("avg is %0.2f\n", avg)

	date := encaps2.Date{}
	date.SetDay(20)
	fmt.Println("Today", date)

}
