package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gosuri/uilive"
)

var Green = "\033[32m"
var Red = "\033[31m"

func startSpinner() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[32], 100*time.Millisecond)
	s.Start()
	return s
}

func Commonword() {
	s := startSpinner()
	word := getWord()
	s.Stop()
	fmt.Println("Score:", readScore())
	fmt.Println("The word has", len(word), "letters")

	for i := 6; i > 0; i-- {
		fmt.Println()
		fmt.Print("Enter your guess ", Red, "(", i, " tries left): ")
		fmt.Printf("\x1b[0m")
		var guess string
		fmt.Scanln(&guess)
		guess = strings.ToLower(guess)

		if len(guess) != len(word) {
			fmt.Println("Your guess must be", len(word), "letters long")
			i++
			continue
		}

		if guess == word {
			fmt.Println()
			fmt.Println(Green + "HOORAY You Won! 🥳🎉")
			score(i)
			playAgain()
		}

		for i := 0; i < len(word); i++ {
			if guess[i] == word[i] {
				fmt.Printf("\x1b[32;1m%c\x1b[0m", guess[i])
			} else if strings.Contains(word, string(guess[i])) {
				fmt.Printf("\x1b[33;1m%c\x1b[0m", guess[i])
			} else {
				fmt.Printf("%c", guess[i])
			}
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println()

	fmt.Println("You lost! The word was " + Green + word)
	writeScore(0)

	playAgain()

}

func Rareword() {
	s := startSpinner()
	word, definition := getRareWord()
	s.Stop()

	fmt.Println("Score:", readScore())
	fmt.Println("The word has", len(word), "letters")
	for i := 6; i > 0; i-- {
		fmt.Println()
		fmt.Print("Enter your guess ", Red, "(", i, " tries left): ")
		fmt.Printf("\x1b[0m")
		var guess string
		fmt.Scanln(&guess)
		guess = strings.ToLower(guess)

		if len(guess) != len(word) {
			fmt.Println("Your guess must be", len(word), "letters long")
			i++
			continue
		}

		if guess == word {
			fmt.Println()
			fmt.Println(Green + "HOORAY You Won! 🥳🎉")
			fmt.Println("\x1b[0m" + "Definition📖: " + Green + definition)
			score(i)
			playAgain()
		}

		for i := 0; i < len(word); i++ {
			if guess[i] == word[i] {
				fmt.Printf("\x1b[32;1m%c\x1b[0m", guess[i])
			} else if strings.Contains(word, string(guess[i])) {
				fmt.Printf("\x1b[33;1m%c\x1b[0m", guess[i])
			} else {
				fmt.Printf("%c", guess[i])
			}
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println()

	fmt.Println("You lost! The word was " + Green + word)
	writeScore(0)
	fmt.Println("\x1b[0m" + "Definition📖: " + Green + definition)

	playAgain()
}

func playAgain() {
	fmt.Println()
	fmt.Print("Play again? (y/n): ")
	var choice string
	fmt.Scan(&choice)
	choice = strings.ToLower(choice)
	if choice == "y" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		if os.Args[1] == "--common" || os.Args[1] == "-c" {
			Commonword()
		}
		if os.Args[1] == "--rare" || os.Args[1] == "-r" {
			Rareword()
		}

	} else if choice == "n" {

		writer := uilive.New()
		writer.Start()

		fmt.Fprintln(writer, "Thanks For Playing! 👋🏼")
		time.Sleep(500 * time.Millisecond)
		fmt.Fprintln(writer, "Thanks For Playing! 👋🏼👋🏽")
		time.Sleep(500 * time.Millisecond)
		fmt.Fprintln(writer, "Thanks For Playing! 👋🏼👋🏽👋🏿")
		time.Sleep(500 * time.Millisecond)

		writer.Stop()
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		os.Exit(0)
	} else {
		fmt.Println("Invalid input")
		playAgain()
	}
}

func Help() {
	fmt.Println("Usage: wordy [--common(-c)/--rare(-r)/--help(-h)/--rules(-R)]")
	fmt.Println()
	fmt.Println("common: Guess a common word")
	fmt.Println("rare: Guess a rare word")
	fmt.Println("help: Display this help message")
	fmt.Print("rules: Display the rules of the game")
}

func Rules() {
	fmt.Println("Rules:")
	fmt.Println("1. You have 6 tries to guess the word")
	fmt.Println("2. Your guess must be the same length as the word")
	fmt.Println("3. If you guess a letter that is in the word, it will be highlighted in yellow")
	fmt.Println("4. If you guess a letter that is in the correct position, it will be highlighted in green")
	fmt.Println("5. If you guess a letter that is not in the word, it will be printed normally")
	fmt.Println("6. If you guess the word correctly, you win!")
	fmt.Println("7. If you run out of tries, you lose!")
	fmt.Print("8. Have fun!")
}

func score(i int) {
	switch i {
	case 6:
		fmt.Println("You got a perfect score! 🏆")
		writeScore(100)
	case 5:
		fmt.Println("You got a score of 83% 🥇")
		writeScore(83)
	case 4:
		fmt.Println("You got a score of 66% 🥈")
		writeScore(66)
	case 3:
		fmt.Println("You got a score of 50% 🥉")
		writeScore(50)
	case 2:
		fmt.Println("You got a score of 33% 🏅")
		writeScore(33)
	case 1:
		fmt.Println("You got a score of 16% 🎖")
		writeScore(16)
	}
}
