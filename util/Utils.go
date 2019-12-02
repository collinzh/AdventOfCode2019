package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Position struct {
	X, Y int
}

type Direction int

const (
	Up    = Direction(1)
	Down  = Direction(2)
	Left  = Direction(3)
	Right = Direction(4)
)

func (d Direction) TurnLeft() Direction {
	switch d {
	case Up:
		return Left
	case Down:
		return Right
	case Left:
		return Down
	case Right:
		return Up
	}
	panic(errors.New(fmt.Sprintf("unrecognized direction: %d", int(d))))
}

func (d Direction) TurnRight() Direction {
	switch d {
	case Up:
		return Right
	case Down:
		return Left
	case Left:
		return Up
	case Right:
		return Down
	}
	panic(errors.New(fmt.Sprintf("unrecognized direction: %d", int(d))))
}

func ScanToIntSlice(f *os.File) []int {
	numbers := make([]int, 0)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func ScanToStringSlices(file *os.File) []string {
	str := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	return str
}

func ScanToRuneSlices(file *os.File) []rune {
	runes := make([]rune, 0)
	reader := bufio.NewReader(file)
	for true {
		r, s, e := reader.ReadRune()
		if e != nil || s == 0 {
			break
		}
		runes = append(runes, r)
	}
	return runes
}

func Distance(a, b *Position) int {
	return IntegerAbs(a.X-b.X) + IntegerAbs(a.Y-b.Y)
}

func IntegerAbs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

const Letters string = "abcdefghijklmnopqrstuvwxzy"
