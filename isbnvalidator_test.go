package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestValidateIsbn10(t *testing.T) {
	ISBNs := readFile("isbn10-examples.txt")
	for _, isbn := range ISBNs {
		result := validateIsbn10(isbn)
		if !result {
			t.Fatal("Invalid ISBN: " + isbn)
		}
	}
}

func TestValidateIsbn13(t *testing.T) {
	ISBNs := readFile("isbn13-examples.txt")
	for _, isbn := range ISBNs {
		result := validateIsbn13(isbn)
		if !result {
			t.Fatal("Invalid ISBN: " + isbn)
		}
	}
}

func readFile(filename string) []string {
	file, _ := ioutil.ReadFile(filename)
	contents := string(file)
	ISBNs := strings.Split(string(contents), "\n")[1:]

	return ISBNs
}
