package day5

import (
	"bufio"
	"fmt"
	"github.com/collinzh/AdventOfCode2019/util"
	"os"
	"strings"
)

type D5Solution struct {
}

type IntCodeComputer struct {
	Memory         []int
	ProgramCounter int
}

func NewIntCodeComputer(memory []int) *IntCodeComputer {

	return &IntCodeComputer{Memory: memory, ProgramCounter: 0}
}

func (icc *IntCodeComputer) readNext() int {
	val := icc.Memory[icc.ProgramCounter]
	icc.ProgramCounter++
	return val
}

func (icc *IntCodeComputer) HasNext() bool {
	return icc.ProgramCounter < len(icc.Memory)
}

func (icc *IntCodeComputer) RunNext() bool {
	if !icc.HasNext() {
		return false
	}

	instruction := icc.readNext()
	opcode := instruction % 100
	p1Mode := instruction / 100 % 10
	p2Mode := instruction / 1000 % 10
	//p3Mode := instruction / 10000 % 10

	var v1, v2, res int

	switch opcode {
	case 1:
		v1 = icc.resolve(p1Mode, icc.readNext())
		v2 = icc.resolve(p2Mode, icc.readNext())
		res = icc.readNext()
		icc.Memory[res] = v1 + v2
	case 2:
		v1 = icc.resolve(p1Mode, icc.readNext())
		v2 = icc.resolve(p2Mode, icc.readNext())
		res = icc.readNext()
		icc.Memory[res] = v1 * v2
	case 3:
		v1 := icc.readNext()
		fmt.Print("Input: ")
		iScanner := bufio.NewScanner(os.Stdin)
		if iScanner.Scan() {
			input := util.MustAtoi(iScanner.Text())
			icc.Memory[v1] = input
		}
	case 4:
		v1 := icc.resolve(p1Mode, icc.readNext())
		fmt.Printf("Output: %d\n", v1)
	case 5:
		v1 := icc.resolve(p1Mode, icc.readNext())
		v2 := icc.resolve(p2Mode, icc.readNext())
		if v1 != 0 {
			icc.ProgramCounter = v2
		}
	case 6:
		v1 := icc.resolve(p1Mode, icc.readNext())
		v2 := icc.resolve(p2Mode, icc.readNext())
		if v1 == 0 {
			icc.ProgramCounter = v2
		}
	case 7:
		v1 := icc.resolve(p1Mode, icc.readNext())
		v2 := icc.resolve(p2Mode, icc.readNext())
		v3 := icc.readNext()
		if v1 < v2 {
			icc.Memory[v3] = 1
		} else {
			icc.Memory[v3] = 0
		}
	case 8:
		v1 := icc.resolve(p1Mode, icc.readNext())
		v2 := icc.resolve(p2Mode, icc.readNext())
		v3 := icc.readNext()
		if v1 == v2 {
			icc.Memory[v3] = 1
		} else {
			icc.Memory[v3] = 0
		}
	case 99:
		return false
	}

	return true
}

func (icc *IntCodeComputer) resolve(mode, param int) int {
	if mode == 1 {
		return param
	}
	return icc.Memory[param]
}

func (icc *IntCodeComputer) PrintMemoryMap() {
	for idx, data := range icc.Memory {
		fmt.Printf("%d\t%d\n", idx, data)
	}
}

func (d D5Solution) RunSolutionP1() {
	d.runInputFile("day5.txt", false)
}

func (d D5Solution) RunSolutionP2() {
	d.runInputFile("day5.txt", false)
}

func (d D5Solution) runInputFile(input string, printMemory bool) {
	f := util.MustOpenFile(input)
	lines := util.ScanToStringSlices(f)
	splitted := strings.Split(lines[0], ",")
	mem := make([]int, len(splitted))
	for idx, str := range splitted {
		mem[idx] = util.MustAtoi(str)
	}

	icc := NewIntCodeComputer(mem)
	for icc.RunNext() {

	}
	if printMemory {
		icc.PrintMemoryMap()
	}
}
