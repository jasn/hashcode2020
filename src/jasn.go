package src

<<<<<<< Updated upstream
import (
	"errors"
)

=======
>>>>>>> Stashed changes
type Output struct {
	Libraries []LibraryAnswer
}

type LibraryAnswer struct {
	ID    LibraryID
	Books []BookID
}

<<<<<<< Updated upstream
func Score(input Input, output Output) (int, error) {
	score = 0
	librariesSeen := map[LibraryID]bool{}
	booksSeen := map[LibraryID]bool{}
	day := 0
	var err error
	for _, l := range output.Libraries {
		if librariesSeen[l.ID] {
			return 0, errors.New("Library %v is duplicated in output", l.ID)
		}
		librariesSeen[l.ID] = true
		scoreAdd, err := processLibrary(day, input, l, booksSeen, librariesSeen)

		day += getDaysForSignUp(input.Libraries, l.ID)
	}

	return score, nil
}

func processLibrary(day int, input Input, lib LibraryAnswer, booksSeen map[LibraryID]bool, libsSeen map[LibraryID]bool) int {
	signUpDays := input.Libraries.DaysForSignUp
	day += signUpDays

	scoreAdd := 0
	for i := 0; i < len(lib.Books); i++ {
		if input.Days <= day {
			break
		}
		for j := 0; j < input.Library.BooksShippedPerDay && i+j < len(lib.Books); j++ {
			bookID := lib.Books[i+j]
			if _, ok := booksSeen[bookID]; !ok {
				scoreAdd += input.BooksScore[bookID]
				booksSeen[bookID] = true
			}
			i++
		}
	}

	return scoreAdd
}
=======
//func Score(input Input, output Output) (int, error) {
//	score = 0
//	librariesSeen := map[LibraryID]bool{}
//	booksSeen := map[LibraryID]bool{}
//	for _, l := range output.Libraries {
//		if librariesSeen[l.ID] {
//			return 0, errors.New("Library %v is duplicated in output", l.ID)
//		}
//		librariesSeen[l.ID] = true
//
//		for _, bookID := range l.Books {
//			if _, ok := booksSeen[bookID]; !ok {
//				score += input.BooksScore[bookID]
//				booksSeen[bookID] = true
//			}
//		}
//	}
//
//	if days, ok := assertOnlySingleLibraryAtATime(input, output); !ok {
//		return errors.New("Cannot sign up that many libraries used days: %v", days)
//	}
//
//	return score
//}
//
//func assertOnlySingleLibraryAtATime(input Input, output Output) (bool, int) {
//	days := input.Days
//	day := 0
//	for _, library := range output.Libraries {
//		d := getDaysForSignUp(input.Libraries, library.ID)
//		day += d
//	}
//	if day > days {
//		return false, day
//	}
//	return true, day
//}
//
//func assertLibraryCanProduceBooks()
//
//func getDaysForSignUp(libraries []Library, id LibraryID) int {
//	return libraries[id].DaysForSignUp
//}
>>>>>>> Stashed changes
