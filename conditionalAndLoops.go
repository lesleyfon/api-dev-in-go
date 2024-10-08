package main

import "fmt"

func ifStatements() {
	bicycle := 2
	tricycle := 3
	if bicycle > tricycle {
		fmt.Println("We have more bicycles than tricycles")
	}

	if tricycle > bicycle {
		fmt.Println("We have more tricycles than Bicycles")
	}

}

func ifElseStatements() {
	Username := "Greghu"
	if Username == "Lizzy" {
		fmt.Println("Elizabeth's username is Lizzy")
	} else {
		fmt.Println("Greg's username is", Username)
	}
}

func ifElseIfStatements() {
	Username := "Greghu"
	if Username == "Lizzy" {
		fmt.Println("Elizabeth's username is Lizzy")
	} else if Username == "Greghu" {
		fmt.Println("Greg's username is", Username)
	} else {
		fmt.Println("Ooops, please sign up!")
	}
}

func switchStatements() {
	age := 31
	switch age {
	case 1:
		fmt.Println("toddler")
	case 13:
		fmt.Println("Teenager")
	case 18:
		fmt.Println("Adult")
	}

	animal := "dog"
	fmt.Print("A ", animal, " is a ")
	switch animal {
	case "dog", "cat", "frog":
		fmt.Println("pet animal")
	case "lion", "tiger", "elephant":
		fmt.Println("zoo animal")
	default:
		fmt.Println("Ooops, we don't know this animal")
	}

}

func loops() {
	//Single condition loop
	fmt.Println("Single condition loop")
	i := 0
	for i <= 8 {
		i = i + 1
	}
	//Classic loop structure
	fmt.Println("Classic loop structure with 3 statements")

	for i := 0; i < 7; i++ {
		fmt.Println(i)
	}
	// /continue statement
	for i := 0; i < 7; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
