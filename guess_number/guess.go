// guess a random number via cli
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	//playGame := true
	x := 5
	for x != 0{
		fmt.Print("Guess a number: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if target == guess {
			fmt.Println("You Win!")
			break
		} else if target < guess{
			fmt.Println("Your guess was too high")
		} else if target > guess {
			fmt.Println("Your guess was too low")
		}
		x--
		fmt.Println("You have", x, "turns left.")
	}
	if x == 0 {
		fmt.Println("You are a loser. The number was", target)
	}
}
