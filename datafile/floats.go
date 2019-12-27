// Package datafile allows reading data samples from files.
package datafile

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("Opening", fileName)
	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing File")
	file.Close()
}

// GetFloats reads a float64 from each line of a file
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	CloseFile(file)
	if scanner.Err() != nil {
		return nil, err
	}
	return numbers, err
}

func GetStrings(fileName string) ([]string, error) {
	var myStrings []string
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		myStr := scanner.Text()
		myStrings = append(myStrings, myStr)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, err
	}
	return myStrings, err
}