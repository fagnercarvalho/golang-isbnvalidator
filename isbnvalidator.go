package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: isbnvalidator.exe <isbn> <ISBN-10 or ISBN-13>")
		os.Exit(-1)
	} else {
		isbn := os.Args[1]
		check := os.Args[2]

		result, err := validateIsbn(isbn, check)
		if err != nil {
			panic(err)
		}

		if result {
			fmt.Println("The ISBN is valid!")
		} else {
			fmt.Println("I'm sorry, the ISBN isn't valid!")
		}
	}
}

func validateIsbn(isbn string, check string) (bool, error) {
	switch check {
	case "ISBN-10":
		return validateIsbn10(isbn), nil
	case "ISBN-13":
		return validateIsbn13(isbn), nil
	default:
		return false, errors.New("The check format for the ISBN is invalid. Select either ISBN-10 or ISBN-13.")
	}
}

func validateIsbn10(isbn string) bool {

	if len(isbn) != 10 {
		return false
	}

	var sum int
	var multiply int = 10
	for i, v := range isbn {
		digitString := string(v)

		if i == 9 && digitString == "X" {
			digitString = "10"
		}

		digit, err := strconv.Atoi(digitString)
		if err != nil {
			panic(err)
		} else {
			sum = sum + (multiply * digit)
			multiply--
		}
	}

	return sum%11 == 0
}

func validateIsbn13(isbn string) bool {

	if len(isbn) != 13 {
		return false
	}

	var sum int
	for i, v := range isbn {
		var multiply int
		if i%2 == 0 {
			multiply = 1
		} else {
			multiply = 3
		}

		digit, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		} else {
			sum = sum + (multiply * digit)
		}
	}

	return sum%10 == 0
}
