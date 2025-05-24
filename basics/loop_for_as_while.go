package main

import (
	"fmt"
	"math/rand"
	"time"
)

func guessingNumber() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// random number betwen 1 & 100
	target := random.Intn(100) + 1

	fmt.Println("Welcome to guessing game")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess what it is?")

	var guess int
	var exit string
	for {
		fmt.Println("Enter your guess:")
		// fmt.Scanln(guess)
		fmt.Scanln(&guess)
		fmt.Scanln(&exit)

		if guess == target {
			fmt.Println("You guess the correct number!!!")
			break
		}

		if guess > target {
			fmt.Println("Too high!!! Try guessing a lower number!!!")
		}

		if guess < target {
			fmt.Println("Too low!!! Try guessing a higher number!!!")
		}

		if exit == "q" {
			fmt.Println("process finished...")
			break
		}
	}
}

func main() {

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	sum := 0
	for {
		sum += 5
		if sum == 10 {
			continue
		}

		fmt.Println("Sum:", sum)

		// if sum >= 20 { // error this condition skips the next condition
		// 	continue
		// }

		if sum >= 25 {
			break
		}
	}

	guessingNumber()

}
