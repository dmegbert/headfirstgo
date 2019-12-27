package main

import (
	"fmt"
	datafile2 "github.com/dmegbert/headfirstgo/datafile"
	"log"
	"sort"
)

func main() {
	counts := make(map[string]int)
	lines, err := datafile2.GetStrings("/Users/degbert/go/bin/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, candidate := range lines {
		counts[candidate]++
	}

	for key, value := range counts {
		fmt.Printf("The candidate %s has %d votes\n", key, value)
	}
	var names []string
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("The candidate %s has %d votes\n", name, counts[name])
	}
}
