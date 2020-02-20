package main

func LibraryPicker(input Input, daysLeft int, usedLibs map[LibraryID]bool) Library {
	bestScore := 0
	bestIdx := 0
	for i, lib := range input.Libraries {
		libID := LibraryID(i)
		if _, ok := usedLibs[libID]; ok {
			continue
		}

		v := Value(daysLeft, lib, input.Books)
		if v >= bestScore {
			bestScore = v
			bestIdx = i
		}
	}
	return bestIdx
}

func Value(daysLeft int, lib Library, books BooksScore) int {
	score := 0
	for i := 0; i < daysLeft; i++ {
		for j := 0; j < lib.BooksShippedPerDay; j++ {
			if i+j >= daysLeft || i+j >= len(lib.BestBooks) {
				return score
			}
			score += lib.BestBooks[i+j]
		}
	}
	return score
}
