package src

import (
	"fmt"
	"testing"
)

func TestNaive(t *testing.T) {

	input := Input{
		BooksTotal:     6,
		LibrariesTotal: 2,
		Days:           7,
		BooksScore:     []int{1, 2, 3, 6, 5, 4},
		Libraries: []Library{
			{
				BooksShippedPerDay: 2,
				DaysForSignUp:      2,
				Books: map[BookID]struct{}{
					0: struct{}{},
					1: struct{}{},
					2: struct{}{},
					3: struct{}{},
					4: struct{}{},
				},
			},
			{
				BooksShippedPerDay: 1,
				DaysForSignUp:      3,
				Books: map[BookID]struct{}{
					3: struct{}{},
					2: struct{}{},
					5: struct{}{},
					0: struct{}{},
				},
			},
		},
	}

	output := naive(input)

	toStdOut(output)
	val, _ := Score(input, output)
	fmt.Println(val)
}
