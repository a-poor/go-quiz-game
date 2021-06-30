package main

import (
	"bufio"
	"fmt"
	"github.com/a-poor/golors"
	"os"
	"strings"
)


const PROBLEMS_PATH = "./problems.csv"
const PASSING_SCORE = 70.0


type problem struct {
	question string
	answer string
}

func read_questions(path string) []problem {
	// Open the Questions file
	file, err := os.Open(path)
	if err != nil {
		golors.Redln("Oh no! An error occured")
		panic(err)
	}
	defer file.Close()

	// Create a scanner to read file
	scanner := bufio.NewScanner(file)

	// Create a slice to hold the results
	var data []problem
	data = make([]problem, 0)

	// Read the file line by line
	for scanner.Scan() {
		row := scanner.Text()

		split := strings.Split(row, ",")
		q, a := split[0], split[1]
		new_prob := problem{ q, a }
		data = append(data, new_prob)
	}

	// Check for error from the scanner
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Return the result
	return data
}

func ask(question string) string {
	var ans string
	golors.Blue(question, "? ")
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
	questions := read_questions(PROBLEMS_PATH)

	// Print the intro
	golors.Blueln("Welcome to my Quiz Game!")
	golors.Grayln("========================\n")

	// Get the player's name
	var name string
	for {
		name = ask("What is your name")
		if len(name) == 0 {
			wrong("That's not a name!")
		} else {
			correct("Great to meet you, "+name+"!\n")
			break
		}
	}

	// Ask the questions
	n_correct := 0
	out_of := 0

	for _, q := range questions {
		response := ask(q.question)
		fmt.Println("You said:       ", response)
		fmt.Println("The answer was: ", q.answer)
		golors.Gray("You said: ", response, ". That is...")
		if response == q.answer {
			golors.Greenln("CORRECT!")
			n_correct++
		} else {
			golors.Redln("WRONG! Sorry...")
		}
		out_of++
	}

	score := float32(n_correct) / float32(out_of) * 100.0

	// Print the results
	golors.Grayln("\nWay to go, ", name, "! You finished the quiz!")
	golors.Whiteln(fmt.Sprintf("You got %d correct out of %d total.", n_correct, out_of))
	golors.Whiteln(fmt.Sprintf("That's a score of: %.2f %%", score))
	if score > PASSING_SCORE {
		golors.Greenln(fmt.Sprintf("Since you scored at least %.2f %%, you passed!", PASSING_SCORE))
	} else {
		golors.Redln(fmt.Sprintf("Since you scored under %.1f %%, you...unfortunately...failed!", PASSING_SCORE))
	}
}
