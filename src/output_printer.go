package main

import (
    "fmt"
    "strconv"
    "strings"
)

func toStdOut(output Output) {
    print(toString(output))
}

func toString(output Output) string {
    var sb strings.Builder
    sb.WriteString(fmt.Sprintf("%d\n", len(output.Libraries)))
    for _, l := range output.Libraries {
        libraryHeader(&sb, l.ID, len(l.Books))
        libraryBooks(&sb, l.Books)
    }
    return sb.String()
}

func libraryBooks(sb *strings.Builder, books []BookID) {
    for inx, b := range books {
        sb.WriteString(strconv.Itoa(int(b)))
        if inx < len(books) {
            sb.WriteString(" ")
        }
    }
    sb.WriteString("\n")
}

func libraryHeader(sb *strings.Builder, id LibraryID, books int) {
    sb.WriteString(strconv.Itoa(int(id)))
    sb.WriteString(" ")
    sb.WriteString(strconv.Itoa(books))
}
