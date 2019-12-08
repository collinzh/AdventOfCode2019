package day4

import (
	"fmt"
	"strconv"
)

type D4Solution struct {
}

func isValidPassword1(pwd int) bool {
	pwdStr := strconv.Itoa(pwd)
	valid := false
	for i := 0; i < len(pwdStr)-1; i++ {
		if pwdStr[i] == pwdStr[i+1] {
			valid = true
		} else if pwdStr[i] > pwdStr[i+1] {
			return false
		}
	}

	return valid
}

func (d D4Solution) RunSolutionP1() {
	counter := 0
	for num := 307237; num <= 769058; num++ {
		if isValidPassword1(num) {
			counter++
		}
	}

	fmt.Printf("Found %d valid passwords\n", counter)
}

func isValidPassword2(pwd int) bool {
	pwdStr := strconv.Itoa(pwd)
	size := len(pwdStr)
	for i := 0; i < size-1; i++ {
		if pwdStr[i] > pwdStr[i+1] {
			return false
		}
	}

	i, j := 0, 0
	for {
		j = i
		if j >= size {
			break
		}

		for j < size-1 && pwdStr[j] == pwdStr[j+1] {
			j++
		}

		if j-i == 1 {
			return true
		} else {
			i = j + 1
		}
	}

	return false
}

func (d D4Solution) RunSolutionP2() {
	counter := 0
	for num := 307237; num <= 769058; num++ {
		if isValidPassword2(num) {
			counter++
		}
	}

	fmt.Printf("Found %d valid passwords\n", counter)
}
