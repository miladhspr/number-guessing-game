package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	ID      int
	Title   string
	Chances int
}

var userAttemptCount = 1

func main() {
	for {
		playGame()
		fmt.Print("\nDo you want to play again? (y/n): ")
		playAgain := getUserInput()
		if strings.ToLower(playAgain) != "y" {
			fmt.Println("Thanks for playing! Goodbye!")
			break
		}
	}
}

func playGame() {
	gamesLevel := startGame()
	difficultyLevel, _ := chooseLevel(gamesLevel)

	randomNumber := generateRandomNumber()
	fmt.Println(randomNumber)
	for i := 1; i <= difficultyLevel.Chances; i++ {
		fmt.Printf("Attempt %d of %d. Enter your guess: ", i, difficultyLevel.Chances)
		guessNumber := getGuess()
		var helper string
		if guessNumber == randomNumber {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.", userAttemptCount)
			return
		} else {
			if guessNumber > randomNumber {
				helper = "less"
			}
			if guessNumber < randomNumber {
				helper = "bigger"
			}
			fmt.Printf("Incorrect! The number is %s than %d.\n", helper, guessNumber)
			userAttemptCount++
		}
	}
	fmt.Printf("Sorry! You've used all %d attempts. The correct number was %d.", difficultyLevel.Chances, randomNumber)
}

func startGame() []Game {
	gamesLevel := []Game{
		{ID: 1, Title: "Easy", Chances: 10},
		{ID: 2, Title: "Medium", Chances: 5},
		{ID: 3, Title: "Hard", Chances: 3},
	}
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have multiple chances to guess the correct number.")
	fmt.Println()
	return gamesLevel
}

func chooseLevel(gamesLevel []Game) (Game, error) {
	fmt.Println("Please select the difficulty level:")
	for _, game := range gamesLevel {
		fmt.Printf("%d. %s (%d chances) \n", game.ID, game.Title, game.Chances)
	}
	fmt.Println("Enter your choice:")
	levelOfGameInput := getUserInput()
	number, err := strconv.Atoi(levelOfGameInput)
	if err != nil || number < 1 || number > len(gamesLevel) {
		fmt.Println("Invalid choice. Please try again.")
		return chooseLevel(gamesLevel)
	}
	difficultyLevel := gamesLevel[number-1]
	fmt.Printf("Great! You have selected the %s difficulty level.\n", difficultyLevel.Title)
	return difficultyLevel, nil
}

func generateRandomNumber() int {
	return rand.Intn(100) + 1
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
func getGuess() int {
	guess := getUserInput()
	guessNumber, _ := strconv.Atoi(guess)
	return guessNumber
}
