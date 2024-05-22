package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

func ReadInputFromFile(fileName string) []string {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := []string{}
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err.Error())
	}

	return input
}

func ToInput(literal string) []string {
	return strings.Split(literal, "\n")
}

func ExpectInt(t *testing.T, actual int, expected int) {
	if expected != actual {
		t.Fatalf("Expected %d but gotten %d", expected, actual)
	}
}
