package main

import (
    "testing"
)

func TestSimulation(t *testing.T) {
    input := LoadInput("a")
    output := Simulation(input)
    toStdOut(output)
}

