package main

import "fmt"

type Output struct {
	Libraries []LibraryAnswer
}

type LibraryAnswer struct {
	ID    LibraryID
	Books []BookID
}

func Score(input Input, output Output) (int, error) {
	score := 0
	librariesSeen := map[LibraryID]bool{}
	booksSeen := map[BookID]bool{}
	day := 0
	for _, l := range output.Libraries {
		if librariesSeen[l.ID] {
			return 0, fmt.Errorf("Library %v is duplicated in output", l.ID)
		}
		librariesSeen[l.ID] = true
		scoreAdd := processLibrary(day, input, l, booksSeen, librariesSeen)

		day += getDaysForSignUp(input.Libraries, l.ID)
		score += scoreAdd
	}

	return score, nil
}

func getDaysForSignUp(libraries []Library, id LibraryID) int {
	return libraries[id].DaysForSignUp
}

func processLibrary(day int, input Input, lib LibraryAnswer, booksSeen map[BookID]bool, libsSeen map[LibraryID]bool) int {
	signUpDays := input.Libraries[lib.ID].DaysForSignUp
	day += signUpDays

	scoreAdd := 0
	for i := 0; i < len(lib.Books); i++ {
		if input.Days <= day {
			break
		}
		for j := 0; j < input.Libraries[lib.ID].BooksShippedPerDay && i+j < len(lib.Books); j++ {
			bookID := lib.Books[i+j]
			if _, ok := booksSeen[bookID]; !ok {
				scoreAdd += input.BooksScore[bookID]
				booksSeen[bookID] = true
			}
			i++
		}
		day++
	}

	return scoreAdd
}
