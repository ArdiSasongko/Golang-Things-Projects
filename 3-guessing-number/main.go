package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	random := rand.Intn(100)
	diff := bufio.NewReader(os.Stdin)

	guest := 0
	fmt.Println("Guessing the random number!")
	fmt.Println("Pick difficulty :")
	fmt.Println("1. Easy (10 Guesses)")
	fmt.Println("2. Medium (7 Guesses)")
	fmt.Println("3. Hard (5 Guesses)")
	fmt.Println("4. Impossible (3 Guesses)")
	fmt.Printf("Choose your difficulty : ")
	inputDif, _ := diff.ReadString('\n')
	inputDif = strings.TrimSpace(inputDif)

	switch inputDif {
	case "1":
		guest = 10
		fmt.Println("You chose Easy difficulty.")
	case "2":
		guest = 7
		fmt.Println("You chose Medium difficulty.")
	case "3":
		guest = 5
		fmt.Println("You chose Hard difficulty.")
	case "4":
		guest = 3
		fmt.Println("You chose Hard difficulty.")
	default:
		fmt.Println("Invalid choice. Exiting the game.")
		return
	}

	fmt.Println("Start guessing ....")
	guestNumber := bufio.NewReader(os.Stdin)

	for i := 0; i < guest; i++ {
		fmt.Printf("Guess number (%d guesses left): ", guest-i)
		inputNums, _ := guestNumber.ReadString('\n')
		gM, err := strconv.Atoi(strings.TrimSpace(inputNums))

		if err != nil {
			fmt.Println("Invalid input, please enter a number")
			i--
			continue
		}

		if gM == random {
			fmt.Println("Correct!!")
			break
		} else if gM > random {
			fmt.Println("Incorrect, number too high")
		} else {
			fmt.Println("Incorrect, number too low")
		}
	}

	fmt.Println("")
	fmt.Println("The correct number is", random)
	fmt.Println("Thanks for playing!")
}
