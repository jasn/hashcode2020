package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"path/filepath"
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

	for _, name := range names {
		input := LoadInput(name)
		output := LoadBestOutput(name)
		if output == nil {
			res := naive(input)
			output = &res
		}
	}

	for  {
		i := rand.Intn(len(outputs))
		newOutput := tryImprove(inputs[i], outputs[i])
		if newOutput != nil {

		}
	}

}
