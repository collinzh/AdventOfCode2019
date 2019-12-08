package main

import (
	"github.com/collinzh/AdventOfCode2019/day1"
	"github.com/collinzh/AdventOfCode2019/day2"
	"github.com/collinzh/AdventOfCode2019/day3"
	"github.com/collinzh/AdventOfCode2019/day4"
)

type Solution interface {
	RunSolutionP1()

	RunSolutionP2()
}

func main() {
	RunDay4()
}

func RunDay1() {
	s := &day1.D1Solution{}
	s.RunSolutionP1()
	s.RunSolutionP2()
}

func RunDay2() {
	s := &day2.D2Solution{}
	s.RunSolutionP1()
	s.RunSolutionP2()
}

func RunDay3() {
	s := &day3.D3Solution{}
	s.RunSolutionP1()
}

func RunDay4() {
	s := &day4.D4Solution{}
	s.RunSolutionP1()
	s.RunSolutionP2()
}
