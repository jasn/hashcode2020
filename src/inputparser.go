package src


type BookID int

type Input struct {
	BooksTotal int
	LibrariesTotal int
	Days int

	BooksScore BooksScore
	Libraries []Library
}

type BooksScore []int

type Library struct {
	

}