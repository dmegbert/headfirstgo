package main

import (
	"fmt"
	util2 "github.com/dmegbert/headfirstgo/util"
	"github.com/headfirstgo/greeting/dansk"
	"log"
)

func main() {
	dansk.Hej()
	width, err := util2.GetFloatInput("What is the width of the wall: ")
	if err != nil {
		log.Fatal(err)
	}
	height, err := util2.GetFloatInput("What is the height of the wall: ")
	if err != nil {
		log.Fatal(err)
	}

	instructions, err := getPaint(width, height)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(instructions)
}

func getPaint(width float64, height float64) (string, error) {
	if width < 0 {
		return "", fmt.Errorf("a width of %0.2f is invalid", width)
	}
	area := width * height
	paint := area / 10.0
	return fmt.Sprintf("You need %0.2f liters of paint", paint), nil
}
