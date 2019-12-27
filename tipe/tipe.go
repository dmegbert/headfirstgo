package main

import "fmt"

func main() {
	tank := Liters(10)
	otherTank := tank.LitersToGallons()
	tank.DoubleMe()
	fmt.Printf("other tank is %#v\n", otherTank)
	fmt.Printf("Doubled is %0.2f", tank)
}

type Liters float64
type Gallons float64

func (l *Liters) LitersToGallons() Gallons {
	return Gallons(*l * 0.264)
}

func (l *Liters) DoubleMe() {
	*l *= 2
}
