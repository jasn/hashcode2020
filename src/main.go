package main

import "os"

func main() {
    name := os.Args[1]

    input := LoadInput(name)
    output := naive(input)
    toStdOut(output)
}
