package main

import (
	"github.com/dmegbert/headfirstgo/magazine"
)

func main() {
	sub := magazine.Subscriber{
		Name:   "Joe",
		Rate:   22,
		Active: true,
	}
	sub.City = "Euclid"
	sub.State = "OH"
	sub.PostalCode = "44119"
	sub.Street = "19350 Monterey Ave."
	defSub := magazine.DefaultSubscriber("Bob")
	magazine.ApplyDiscount(&sub)
	magazine.PrintInfo(&sub)
	magazine.PrintInfo(defSub)
}
