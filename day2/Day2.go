package day2

import (
	"fmt"
	"github.com/collinzh/AdventOfCode2019/util"
	"strconv"
	"strings"
)

type D2Solution struct {
}

func (d D2Solution) RunSolutionP1() {
	ins := parseInstructions("day2.txt")
	pc := 0

LOOP:
	for {
		code := ins[pc]
		switch code {
		case 1:
			ins[ins[pc+3]] = ins[ins[pc+1]] + ins[ins[pc+2]]
			pc += 4
		case 2:
			ins[ins[pc+3]] = ins[ins[pc+1]] * ins[ins[pc+2]]
			pc += 4
		case 99:
			break LOOP
		}
	}

	fmt.Printf("Result: %d\n", ins[0])
}

func (d D2Solution) RunSolutionP2() {
	template := parseInstructions("day2.txt")
	for noun := 1; noun <= 99; noun++ {
		for verb := 1; verb <= 99; verb++ {
			ins := make([]int, len(template))
			copy(ins, template)
			pc := 0
			ins[1] = noun
			ins[2] = verb

		LOOP:
			for {
				code := ins[pc]
				switch code {
				case 1:
					ins[ins[pc+3]] = ins[ins[pc+1]] + ins[ins[pc+2]]
					pc += 4
				case 2:
					ins[ins[pc+3]] = ins[ins[pc+1]] * ins[ins[pc+2]]
					pc += 4
				case 99:
					break LOOP
				}

				if ins[0] == 19690720 {
					fmt.Printf("Result: %d\n", 100*noun+verb)
					return
				}
			}
		}
	}
}

func parseInstructions(path string) []int {
	f := util.MustOpenFile(path)
	defer f.Close()
	line := util.ScanToStringSlices(f)[0]
	splitted := strings.Split(line, ",")
	instructions := make([]int, len(splitted))
	for idx, str := range splitted {
		instructions[idx], _ = strconv.Atoi(str)
	}
	return instructions
}
