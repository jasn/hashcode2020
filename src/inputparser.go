package main

import (
	"io/ioutil"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const dataFolder = "../data"

type Input struct {
	BooksTotal     int
	LibrariesTotal int
	Days           int

	BooksScore BooksScore
	Libraries  []*Library
}

type BooksScore []int

type Library struct {
	DaysForSignUp      int
	BooksShippedPerDay int
	Books              []BookID
	BestBooks          []SortedBook
	BestUniqueBooks    []SortedBook
	ID                 LibraryID
}

type SortedBook struct {
	Book  BookID
	Score int
}

func LoadInput(name string) *Input {
	dataFiles, err := ioutil.ReadDir(dataFolder)
	if err != nil {
		panic(err.Error())
	}
	fileName := ""
	for _, file := range dataFiles {
		if strings.HasPrefix(file.Name(), name+"_") && strings.HasSuffix(file.Name(), ".txt") {
			fileName = file.Name()
			break
		}
	}
	if len(fileName) == 0 {
		panic("wrong filename")
	}
	fileContent, err := ioutil.ReadFile(filepath.Join(dataFolder, fileName))
	if err != nil {
		panic(err.Error())
	}

	return PostProcess(Parse(string(fileContent)))
}

func PostProcess(i *Input) *Input {
	for _, library := range i.Libraries {
		var sortedBooks []SortedBook
		for _, book := range library.Books {
			sortedBooks = append(sortedBooks, SortedBook{
				Book:  book,
				Score: i.BooksScore[book],
			})
		}
		library.BestBooks = sortedBooks
		sort.Slice(library.BestBooks, func(i, j int) bool {
			return library.BestBooks[i].Score > library.BestBooks[j].Score
		})
	}
	return i
}

func Parse(s string) *Input {
	lines := strings.Split(s, "\n")
	firstLine := strings.Split(lines[0], " ")

	input := &Input{
		BooksTotal:     parseInt(firstLine[0]),
		LibrariesTotal: parseInt(firstLine[1]),
		Days:           parseInt(firstLine[2]),
	}

	secondLine := strings.Split(lines[1], " ")
	for _, bookScoreString := range secondLine {
		input.BooksScore = append(input.BooksScore, parseInt(bookScoreString))
	}

	for i := 2; i+1 < len(lines); i += 2 {
		firstLibraryLine := strings.Split(lines[i], " ")
		secondLibraryLine := strings.Split(lines[i+1], " ")

		if len(firstLibraryLine) < 3 {
			continue
		}

		//books := make(map[BookID]struct{}, parseInt(firstLibraryLine[0]))
		var books []BookID
		for _, book := range secondLibraryLine {
			books = append(books, BookID(parseInt(book)))
		}

		input.Libraries = append(input.Libraries, &Library{
			DaysForSignUp:      parseInt(firstLibraryLine[1]),
			BooksShippedPerDay: parseInt(firstLibraryLine[2]),
			Books:              books,
			ID:                 LibraryID(len(input.Libraries)),
		})
	}
	return input
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return i
}
