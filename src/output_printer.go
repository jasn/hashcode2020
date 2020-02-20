package src

import "fmt"

func toStdOut(output Output) {
    fmt.Println(len(output.Libraries))
    for _, l := range output.Libraries {
        libraryHeader(l.ID, len(l.Books))
        libraryBooks(l.Books)
    }
}

func libraryBooks(books []BookID) {
    for inx, b := range books {
        fmt.Print(b)
        if inx < len(books) {
            fmt.Print(" ")
        }
    }
    fmt.Println("")
}

func libraryHeader(id LibraryID, books int) {
    fmt.Print(id)
    fmt.Print(" ")
    fmt.Println(books)
}
