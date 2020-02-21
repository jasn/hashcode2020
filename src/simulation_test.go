package main

import (
    "fmt"
    "sort"
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

func TestSimulationA(t *testing.T) {

    score := 0
    for _, name := range []string{"a"} {
        input := LoadInput(name)
        output := Simulation(input)
        sc, err := Score(input, output)
        if err != nil {
            fmt.Println(err)
        }
        score += sc
        fmt.Println(fmt.Sprintf("score (%s): %d", name, sc))
        toStdOut(output)
    }
    fmt.Println(fmt.Sprintf("score (total): %d", score))
}


func TestSimulationFo(t *testing.T) {

    score := 0
    for _, name := range []string{"f"} {
        input := LoadInput(name)
        output := Simulation(input)
        //for _, l := range output.Libraries {
            //fmt.Println(input.Libraries[l.ID].DaysForSignUp)
        //}

        sc, err := Score(input, output)
        if err != nil {
            fmt.Println(err)
        }
        score += sc
        fmt.Println(fmt.Sprintf("score (%s): %d", name, sc))
        fmt.Println("---- ")
        total := 0
        min := 10000000
        //libs := []Library{}
        for _, l := range input.Libraries {
            //fmt.Println(l.DaysForSignUp)
            total += l.DaysForSignUp
            if l.DaysForSignUp < min {
                min = l.DaysForSignUp
            }
        }
        sort.Slice(input.Libraries, func(i, j int) bool {
           return input.Libraries[i].DaysForSignUp < input.Libraries[j].DaysForSignUp
        })

        fmt.Println("---- ")
        s := 0
        libs := []*Library{}
        for inx, l := range input.Libraries {
            if inx > 24 {
                break
            }
            s += l.DaysForSignUp
            //fmt.Println(l.DaysForSignUp)
            libs = append(libs, l)
            l.ID = LibraryID(inx)
        }

        fmt.Println("---- ")
        fmt.Println(s)
        fmt.Println(total/len(input.Libraries))
        fmt.Println(min)

        input.Libraries = libs
        input.LibrariesTotal = len(libs)
        output = Simulation(input)
        sc, err = Score(input, output)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(fmt.Sprintf("score (%s): %d", name, sc))


        //toStdOut(output)
    }
    fmt.Println(fmt.Sprintf("score (total): %d", score))
}

