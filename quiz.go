package main

import (
	"fmt"
	//"os"
	"github.com/a-poor/golors"
	"strings"
)


//const PROBLEMS_PATH = "./problems.csv"


//func read_questions(path string) {
//	f, err := os.Open(path)
//	if err != nil {
//		golors.Redln("Oh no! An error occured")
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
	golors.Blue(question, " ")
	fmt.Scanln(&ans)
	return strings.TrimSpace(ans)
}

func correct(message string) {
	golors.Greenln(message)
}

func wrong(message string) {
	golors.Redln(message)
}

func main() {
	// Load the questions


	// Print the intro
	golors.Blueln("Welcome to my Quiz Game!")
	golors.Grayln("========================\n")

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
