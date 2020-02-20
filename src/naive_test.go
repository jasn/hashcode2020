package main

import (
	"fmt"
	"testing"
)

func TestNaive(t *testing.T) {

	input := &Input{
		BooksTotal:     6,
		LibrariesTotal: 2,
		Days:           7,
		BooksScore:     []int{1, 2, 3, 6, 5, 4},
		Libraries: []*Library{
			{
				BooksShippedPerDay: 2,
				DaysForSignUp:      2,
				Books:              []BookID{0, 1, 2, 3, 4},
			},
			{
				BooksShippedPerDay: 1,
				DaysForSignUp:      3,
				Books:              []BookID{3, 2, 5, 0},
			},
		},
	}

	output := naive(input)

	toStdOut(output)
	val, _ := Score(input, output)
	fmt.Println("score:", val)
}
