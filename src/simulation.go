package main

func Simulation(input *Input) Output {

    usedLibs := map[LibraryID]bool{}
    usedBooks := map[BookID]bool{}

    libraries := []LibraryAnswer{}
    daysLeft := input.Days
    for daysLeft > 0 {
        lib, ok := LibraryPicker(input, daysLeft, usedLibs, usedBooks)
        if !ok {
            break
        }
        usedLibs[lib.ID] = true
        for _, b := range lib.BestUniqueBooks {
            usedBooks[b.Book] = true
        }
        a := LibraryAnswer{
            ID:    lib.ID,
            Books: convertToBookIDs(lib.BestUniqueBooks),
        }
        libraries = append(libraries, a)
        daysLeft -= lib.DaysForSignUp
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
