package main

func Simulation(input *Input) Output {

    usedLibs := map[LibraryID]bool{}

    libraries := []*LibraryAnswer{}

    for daysLeft := input.Days; daysLeft > 0; daysLeft++ {
        lib := LibraryPicker(input, daysLeft, usedLibs)
        usedLibs[lib.ID] = true
        a := &LibraryAnswer{
            ID:    lib.ID,
            Books: convertToBookIDs(lib.BestBooks),
        }
        libraries = append(libraries, a)
    }

    return Output{}
}

func convertToBookIDs(sortedBooks []SortedBook) []BookID {
    books := []BookID{}
    for _, b := range sortedBooks {
        books = append(books, b.Book)
    }
    return books
}
