# Number Guessing Game

## Overview
The Number Guessing Game is a simple CLI-based game where the computer randomly selects a number, and the user must guess it. The user has a limited number of chances based on the selected difficulty level. The game provides feedback on whether the guess was too high or too low and ends when the user guesses correctly or runs out of attempts. 

The game also tracks high scores and the time taken to guess the number.

## Features
- CLI-based interaction build with pure Go language.
- Selectable difficulty levels (Easy, Medium, Hard).
- Feedback on whether the guess is too high or too low.
- Multiple rounds of play.
- Timer to track the duration of each game.
- High score tracking for each difficulty level.

## How to Play

1. **Start the Game**:
   - When the game starts, it displays a welcome message and explains the rules.

2. **Select Difficulty Level**:
   - You can choose from three difficulty levels:
     - **Easy**: 10 chances to guess.
     - **Medium**: 5 chances to guess.
     - **Hard**: 3 chances to guess.

3. **Make a Guess**:
   - After selecting the difficulty level, you will be prompted to enter a guess.
   - If your guess is correct, the game ends with a congratulatory message and displays the number of attempts it took.
   - If your guess is incorrect, the game will tell you whether the correct number is higher or lower than your guess.

4. **End of the Game**:
   - The game ends either when you guess the correct number or run out of chances.
   - After the game ends, you can choose to play again.

5. **High Scores**:
   - If you beat the current high score for a difficulty level, the game will update and display the new high score.

## Example Output
```
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have multiple chances to guess the correct number.

Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)
   Enter your choice: 2

Great! You have selected the Medium difficulty level.
Let's start the game!

Attempt 1 of 5. Enter your guess: 50
Incorrect! The number is less than your guess.

Attempt 2 of 5. Enter your guess: 25
Incorrect! The number is greater than your guess.

Attempt 3 of 5. Enter your guess: 35
Congratulations! You guessed the correct number in 3 attempts.
It took you 12.45 seconds.

New high score for Medium level! 3 attempts.

Do you want to play again? (y/n):
```

## Installation and Setup

### Prerequisites
- [Go](https://golang.org/doc/install) installed on your machine.

### Clone the Repository
```sh
git clone https://github.com/miladhspr/number-guessing-game.git
cd number-guessing-game
```

### Build the Application
```sh
go build -o number-guessing-game
```

### Run the Game
```sh
./number-guessing-game
```

## Contributing
If you would like to contribute to this project, please fork the repository and submit a pull request. Any contributions, such as bug fixes, feature additions, or documentation improvements, are welcome.

## Roadmap.sh
https://roadmap.sh/projects/number-guessing-game