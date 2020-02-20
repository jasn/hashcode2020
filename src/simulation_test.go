package main

import (
    "fmt"
    "testing"
)

func TestSimulation(t *testing.T) {

    score := 0
    for _, name := range []string{"a", "b", "c", "d", "e", "f"} {
        input := LoadInput(name)
        output := Simulation(input)
        sc, err := Score(input, output)
        if err != nil {
            fmt.Println(err)
        }
        score += sc
        fmt.Println(fmt.Sprintf("score (%s): %d", name, sc))
        //toStdOut(output)
    }
    fmt.Println(fmt.Sprintf("score (total): %d", score))
}

