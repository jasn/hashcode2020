package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"sync"
)

func LoadBestOutput(name string) *Output{
	f, err := ioutil.ReadFile(filepath.Join(dataFolder, name+".best"))
	if err != nil {
		return nil
	}

	var output Output
	err = json.Unmarshal(f, &output)
	if err != nil {
		panic("invalid output")
	}
	return &output
}

func ImproveAllSolutions(names []string) {
	rand.Seed(42)

	var inputs []*Input
	var outputs []*Output
	var scores []int
	for _, name := range names {
		input := LoadInput(name)
		output := LoadBestOutput(name)
		if output == nil {
			res := Simulation(input)
			println(fmt.Sprintf("Computed simulation for %s", name))
			output = &res
		}
		score, err := Score(input, *output)
		if err != nil {
			panic(err.Error())
		}

		inputs = append(inputs, input)
		outputs = append(outputs, output)
		scores = append(scores, score)
	}


	exec := NewThrottledExecutor(context.Background(), 10)
	var lock sync.RWMutex
	for {
		exec.Go(func() error {
			i := rand.Intn(len(outputs))

			lock.RLock()
			curScore := scores[i]
			curOutput := copyOutput(outputs[i])
			input := inputs[i]
			lock.RUnlock()

			newOutput, newScore := tryImprove(input, curOutput, curScore)
			if newOutput != nil {
				lock.Lock()
				println(fmt.Sprintf("Improved %s from %d to %d, total: %d", names[i], scores[i], newScore, total(scores)))

				bytes, err := json.MarshalIndent(newOutput, "", "  ")
				if err != nil {
					panic("failed marshalling output")
				}
				err = ioutil.WriteFile(filepath.Join(dataFolder, names[i]+".best"), bytes, 0775)
				if err != nil {
					panic("failed writing output")
				}

				outputs[i] = newOutput
				scores[i] = newScore
				lock.Unlock()
			}
			return nil
		})
	}
}

func total(scores []int) int {
	cur := 0
	for _, score := range scores {
		cur += score
	}
	return cur
}

func tryImprove(input *Input, output *Output, curScore int) (*Output, int) {
	operation := rand.Intn(1)
	switch operation {
	case 0:
		swapLibrary(output)
	}
	newScore := score(input, output)
	if newScore > curScore {
		return output, newScore
	}
	return nil, -1
}

func swapLibrary(output *Output) {
	x := rand.Intn(len(output.Libraries))
	y := rand.Intn(len(output.Libraries))
	xLibrary := output.Libraries[x]
	output.Libraries[x] = output.Libraries[y]
	output.Libraries[y] = xLibrary
}

func copyOutput(out *Output) *Output {
	bytes, err := json.Marshal(out)
	if err != nil {
		panic(err.Error())
	}
	var copy Output
	err = json.Unmarshal(bytes, &copy)
	if err != nil {
		panic(err.Error())
	}
	return &copy
}

func score(input *Input, output *Output) int {
	score, err := Score(input, *output)
	if err != nil {
		panic(err.Error())
	}
	return score
}