package main

func LibraryPicker(input *Input, daysLeft int, usedLibs map[LibraryID]bool, usedBooks map[BookID]bool) (*Library, bool) {
	//bestScore := 0
	bestIdx := 0
	bestRatio := 0.0
	foundBest := false
	for i, lib := range input.Libraries {
		libID := LibraryID(i)
		if _, ok := usedLibs[libID]; ok {
			continue
		}

		foundBest = true
		v := Value(daysLeft, lib, usedBooks)
		ratio := float64(v)/float64(lib.DaysForSignUp)
		if ratio >= bestRatio {
			//bestScore = v
			bestRatio = ratio
			bestIdx = i
		}
		//if v >= bestScore {
		//	bestScore = v
		//	bestIdx = i
		//}
	}
	return input.Libraries[bestIdx], foundBest
}

func Value(daysLeft int, lib *Library, usedBooks map[BookID]bool) int {
	score := 0
	for i := 0; i < daysLeft; i++ {
		if i >= daysLeft {
			return score
		}
		for j := 0; j < lib.BooksShippedPerDay; j++ {
			if i+j >= len(lib.BestBooks) {
				return score
			}
			//if _, ok := usedBooks[BookID(i+j)]; ok {
			//	continue
			//}
			score += lib.BestBooks[i+j].Score
		}
	}
	return score
}
