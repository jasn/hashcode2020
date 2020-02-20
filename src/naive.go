package main

func naive(input *Input) Output {
	libAnswers := []LibraryAnswer{}
	for idx, l := range input.Libraries {
		books := getBookFromLibrary(l)
		libAnswers = append(libAnswers, LibraryAnswer{
			ID:    LibraryID(idx),
			Books: books,
		})
	}
	return Output{
		Libraries: libAnswers,
	}
}

func getBookFromLibrary(l *Library) []BookID {
	bookIDs := []BookID{}
	for _, bookID := range l.Books {
		bookIDs = append(bookIDs, bookID)
	}
	return bookIDs
}
