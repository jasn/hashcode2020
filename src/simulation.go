package main

func Simulation(input *Input) Output {

    usedLibs := map[LibraryID]bool{}
    usedBooks := map[BookID]bool{}

    libraries := []LibraryAnswer{}

    for daysLeft := input.Days; daysLeft > 0; daysLeft-- {
        lib, ok := LibraryPicker(input, daysLeft, usedLibs, usedBooks)
        if !ok {
            break
        }
        usedLibs[lib.ID] = true
        a := LibraryAnswer{
            ID:    lib.ID,
            Books: convertToBookIDs(lib.BestBooks),
        }
        libraries = append(libraries, a)
    }

    return Output{
        Libraries: libraries,
    }
}

func convertToBookIDs(sortedBooks []SortedBook) []BookID {
    books := []BookID{}
    for _, b := range sortedBooks {
        books = append(books, b.Book)
    }
    return books
}
