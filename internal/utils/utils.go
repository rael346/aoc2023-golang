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
	return strings.Split(strings.Trim(literal, "\n"), "\n")
}

func ExpectInt(t *testing.T, actual int, expected int) {
	if expected != actual {
		t.Fatalf("Expected %d but gotten %d", expected, actual)
	}
}

// Euclidean algorithm for finding GCD
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func All[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if !pred(t) {
			return false
		}
	}
	return true
}
