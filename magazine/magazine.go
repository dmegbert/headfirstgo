package magazine

import "fmt"

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

type Employee struct {
	Name   string
	Salary float64
	Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

func PrintInfo(s *Subscriber) {
	fmt.Println("Name is", s.Name)
	fmt.Println("Rate is", s.Rate)
	fmt.Println("Active is ", s.Active)
	fmt.Printf("Address is %#v\n", s.Address)
}

func DefaultSubscriber(Name string) *Subscriber {
	var s Subscriber
	s.Name = Name
	s.Rate = 5.99
	s.Active = true
	return &s
}

func ApplyDiscount(s *Subscriber) {
	s.Rate = 4.99
}
