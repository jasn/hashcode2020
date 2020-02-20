package main

func LibraryPicker(input *Input, daysLeft int, usedLibs map[LibraryID]bool) (*Library, bool) {
	bestScore := 0
	bestIdx := 0
	foundBest := false
	for i, lib := range input.Libraries {
		libID := LibraryID(i)
		if _, ok := usedLibs[libID]; ok {
			continue
		}

		foundBest = true
		v := Value(daysLeft, lib)
		if v >= bestScore {
			bestScore = v
			bestIdx = i
		}
	}
	return input.Libraries[bestIdx], foundBest
}

func Value(daysLeft int, lib *Library) int {
	score := 0
	for i := 0; i < daysLeft; i++ {
		for j := 0; j < lib.BooksShippedPerDay; j++ {
			if i+j >= daysLeft || i+j >= len(lib.BestBooks) {
				return score
			}
			score += lib.BestBooks[i+j].Score
		}
	}
	return score
}
