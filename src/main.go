package main

import (
    "fmt"
    "os"
)

func main() {
    algo := os.Args[1]
    name := os.Args[2]

    input := LoadInput(name)
    var output Output
    if algo == "naive" {
        output = naive(input)
    } else if algo == "sim" {
        output = Simulation(input)
    } else {
        fmt.Println("unknown argument 1")
    }
    toStdOut(output)
}
