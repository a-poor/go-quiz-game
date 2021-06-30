package main

import (
	"fmt"
	//"os"
	"quiz-game/color"
	"strings"
)


//const PROBLEMS_PATH = "./problems.csv"


//func read_questions(path string) {
//	f, err := os.Open(path)
//	if err != nil {
//		color.Redln("Oh no! An error occured")
//		panic(err)
//	}
//	defer f.Close()
//
//	var data [][]string
//	data = make([][]string, )
//
//	return data
//}

func ask(question string) string {
	var ans string
	color.Blue(question, " ")
	fmt.Scanln(&ans)
	return strings.TrimSpace(ans)
}

func correct(message string) {
	color.Greenln(message)
}

func wrong(message string) {
	color.Redln(message)
}

func main() {
	// Load the questions


	// Print the intro
	color.Blueln("Welcome to my Quiz Game!")
	color.Grayln("========================\n")

	// Get the player's name
	var name string
	for {
		name = ask("What is your name?")
		if len(name) == 0 {
			wrong("That's not a name!")
		} else {
			correct("Great to meet you, "+name+"!\n")
			break
		}
	}

	// Ask the questions
}
