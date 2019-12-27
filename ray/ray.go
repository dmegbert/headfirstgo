package main

import (
	"fmt"
	"time"
)

func main() {
	//Basic array
	var nums [3] int
	nums[0] = 3
	nums[1] = 2
	nums[2] = 1
	for i := 0; i < len(nums); i++ {
		println("array is", nums[i])
	}
	//array using time
	var datetimes [2] time.Time
	datetimes[0] = time.Now()
	datetimes[1] = time.Now()
	for d := 0; d < len(datetimes); d++ {
		p := fmt.Println
		p(datetimes[d].Format(time.RFC3339))
		p(datetimes[d].Format("Mon Jan _2 15:04:05 2006"))
	}
	//array literal
	five :=  [5] int {1, 2, 3, 4, 5}
	for y := 0; y < 5; y++ {
		fmt.Println("thing is", five[y])
	}
	fmt.Println(five)
	fmt.Printf("%#v\n", five)

	for _, num := range five {
		println("is", num)
	}
	thing := [] int {3, 5 , 7}
	for _, x := range thing {
		println(x)
	}
	notes := make([]string, 7)
	for idx, _ := range notes {
		notes[idx] = fmt.Sprint("my idx is", idx)
	}
	for x := 1; x < 5; x++ {
		notes = append(notes, fmt.Sprint("my idx is", x + 6))
	}

	fmt.Printf("my notes %#v", notes)

}
