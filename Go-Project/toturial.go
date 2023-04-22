package main 

import (
	"fmt"
)

func main () {
	fmt.Println("Welcome to my Quiz Game")
	var name string
	var age uint	

	fmt.Printf("Enter your name: ")
	fmt.Scan( &name)
	fmt.Printf("Enter your age: ")
	fmt.Scan( &age)
	if age >= 10 {
		fmt.Println("You can play the game!")
	} else {
		fmt.Println("You must be at least 10 years")
		fmt.Println("You cannot play the game!")
		return
	}
	
	fmt.Printf("Welcome %v and your are %v years old.\n", name, age)

	score := 0
	numofQuestions := 2


	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string 
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer +" "+ answer2 == "RTX 3090" {
		fmt.Println("Correct!")
		score += 1
	} else if answer +" "+ answer2 == "rtx 3090" {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900 have? ")
	var cores uint
	fmt.Scan(&cores)
	if cores == 12 {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("You scored %v out of %v.\n", score, numofQuestions)
	percent := (float64(score) / float64(numofQuestions)) * 100
	fmt.Printf("Your percentage is %v%%.", percent)
}