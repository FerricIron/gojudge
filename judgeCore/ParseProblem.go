package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)
import "../def"

func ParseProblemFile(filename string, problem *def.Problem) (err error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(file, problem)
	if err != nil {
		return
	}
	if len(problem.JudgeList) == 0 {
		return fmt.Errorf("empty problem file")
	}
	return nil
}
