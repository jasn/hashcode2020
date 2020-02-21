package main

func LibraryPicker(input *Input, daysLeft int, usedLibs map[LibraryID]bool, usedBooks map[BookID]bool) (*Library, bool) {
	bestIdx := 0
	bestUniqueBooks := []SortedBook{}
	bestRatio := 0.0
	foundBest := false
	for i, lib := range input.Libraries {
		libID := LibraryID(i)
		if _, ok := usedLibs[libID]; ok {
			continue
		}

		foundBest = true
		v, books := Value(daysLeft, lib, usedBooks)
		ratio := float64(v)/float64(lib.DaysForSignUp)
		if ratio >= bestRatio {
			bestRatio = ratio
			bestIdx = i
			bestUniqueBooks = books
		}
	}
	library := input.Libraries[bestIdx]
	library.BestUniqueBooks = bestUniqueBooks
	return library, foundBest
}

func Value(daysLeft int, lib *Library, usedBooks map[BookID]bool) (int, []SortedBook) {
	score := 0
	bestUniqueBooks := []SortedBook{}
	extra := 0
	for i := 0; i < daysLeft-lib.DaysForSignUp; i++ {
		if i >= daysLeft {
			return score, bestUniqueBooks
		}
		j := 0
		for j < lib.BooksShippedPerDay {
			inx := i * lib.BooksShippedPerDay
			i2 := inx + j + extra
			if i2 >= len(lib.BestBooks) {
				return score, bestUniqueBooks
			}
			bestBook := lib.BestBooks[i2]
			if _, ok := usedBooks[bestBook.Book]; ok {
				extra += 1
				continue
			}
			score += bestBook.Score
			bestUniqueBooks = append(bestUniqueBooks, bestBook)
			j++
		}
	}

	return score, bestUniqueBooks
}
