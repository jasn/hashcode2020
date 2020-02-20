package main

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

var toImprove = []string{"b", "c", "d", "e", "f"}

func Test__improvthemall(t *testing.T) {
	ImproveAllSolutions(toImprove)
}

func Test__convertthemall(t *testing.T) {
	for _, name := range toImprove {
		out := LoadBestOutput(name)
		s := toString(*out)
		err := ioutil.WriteFile(filepath.Join(dataFolder, name+".best.out"), []byte(s), 0775)
		if err != nil {
			panic(err.Error())
		}
	}
}
