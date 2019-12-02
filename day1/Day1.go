package day1

import (
	"fmt"
	"github.com/collinzh/AdventOfCode2019/util"
)

type D1Solution struct {
}

func (s D1Solution) RunSolutionP1() {
	f := util.MustOpenFile("day1.txt")
	defer f.Close()
	numbers := util.ScanToIntSlice(f)
	sum := 0
	for _, num := range numbers {
		sum += num/3 - 2
	}

	fmt.Printf("Sum is %d\n", sum)
}

func (s D1Solution) RunSolutionP2() {
	f := util.MustOpenFile("day1.txt")
	defer f.Close()
	numbers := util.ScanToIntSlice(f)
	sum := 0
	for _, num := range numbers {
		fuel := num/3 - 2
		sum += calculateFuel(fuel) + fuel
	}

	fmt.Printf("Sum is %d\n", sum)
}

func calculateFuel(input int) int {
	output := input/3 - 2
	if output <= 0 {
		return 0
	}

	return output + calculateFuel(output)
}
