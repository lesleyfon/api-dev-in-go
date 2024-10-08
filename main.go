package main

import (
	"fmt"
	"reflect"
)

func addNumbers(a, b int) int {
	allDataTypes()
	return a + b
}

func defineMultipleVar() {
	var cents, books, movie = 7.9, 4, "Star Wars"

	fmt.Println("Tom's change is ", cents, "cents.")
	fmt.Println("cents is of type", reflect.TypeOf(cents))
	fmt.Println("We bought a total of ", books, "books.")
	fmt.Println("books is of type", reflect.TypeOf(books))
	fmt.Println("My favorite movie is", movie, ".")
	fmt.Println("movie is of type", reflect.TypeOf(movie))
}

func allDataTypes() {
	var favorite bool = true
	var number int = 10
	var pi float32 = 4.14
	var firstName string = "Lesley"
	var specialNum = complex(9, 16)

	fmt.Println("A dog is an animal", favorite)
	fmt.Println("My favorite number is the number: ", number)
	fmt.Println("The value of Pi in two decimal places is: ", pi)
	fmt.Println("My name is: ", firstName)
	fmt.Println("An example of a complex number is: ", specialNum)
	fmt.Println("The data type f crew is = ", reflect.TypeOf(specialNum))
}
func constants() {

	const STATE string = "Lagos"
	const WEATHER = 23.7
	const (
		ONE   = 1
		THREE = 3
		SEVEN = 7
	)
}

func operators() {
	var crackers = 9
	var chips = 3

	// Addition
	fmt.Println("We have", crackers+chips, " crackers and chips all together.")
	//Subtraction
	fmt.Println("We have", crackers-chips, " crackers and chips all together.")
	//Multiplies
	fmt.Println("We have", crackers*chips, " crackers and chips all together.")
	//Division
	fmt.Println("We have", crackers/chips, " crackers and chips all together.")
	//Modulus
	fmt.Println("We have", crackers%chips, " crackers and chips all together.")
	//Increment
	crackers++
	fmt.Println("We have", crackers, "crackers")
	//Decrement
	chips--
	fmt.Println("We have", chips, "chips")
}

func comparisons() {
	noodles := 9
	cookies := 3

	//Equal to
	fmt.Println(noodles == cookies)
	//Not Equal to
	fmt.Println(noodles != cookies)
	//Greater than
	fmt.Println(noodles > cookies)
	//Less than
	fmt.Println(noodles < cookies)
	//Greater than or equal to
	fmt.Println(noodles >= cookies)
	//Less than or equal to
	fmt.Println(noodles <= cookies)
}

func assignments() {
	var favInt = 9
	fmt.Println("My favorite int is ", favInt)

	var sisInt = 10
	sisInt = sisInt + favInt
	fmt.Println("Sister int plus favInt", sisInt)

	//Subtraction assignment operator
	var candy = 9
	candy -= 3
	fmt.Println(candy, "candies is too much")
	//Multiply and assign
	var dogs = 9
	dogs *= 3
	fmt.Println("We have", dogs, "dogs")
	//Divide and assign
	var eggs = 9
	eggs /= 3
	fmt.Println("I have ", eggs, "eggs")
}

func logicalOperators() {

	var buses = 3
	//Logical and - returns true if both statements are true, and false otherwise.
	fmt.Println(buses < 3 && buses < 9) // False
	//Logical or - returns false if both statements are false, and true otherwise.
	fmt.Println(buses < 3 || buses < 9) // False
	//Logical not - returns true if the operand is false, and false otherwise.
	fmt.Println(!(buses < 3 && buses < 9))
}

func bitwiseOperation() {
	var buses uint = 9
	var cars uint = 3
	var trains uint

	trains = buses & cars
	fmt.Println("buses & cars  =", trains)

	trains = buses | cars
	fmt.Println("buses | cars  =", trains)

	trains = buses ^ cars
	fmt.Println("buses ^ cars  =", trains)

	trains = buses << 1
	fmt.Println("buses << 1 =", trains)

	trains = buses >> 1
	fmt.Println("buses >> 1 =", trains)
}

/*
This is a multiline comments
Another line here
*/
func main() {
	x := 5
	y := 7
	total := addNumbers(x, y)
	defineMultipleVar()
	constants()
	operators()

	// This line is a single comment
	fmt.Println(total)
}
