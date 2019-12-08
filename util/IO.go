package util

import (
	"os"
	"strconv"
)

func MustOpenFile(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return f
}

func MustAtoi(str string) int {
	if num, err := strconv.Atoi(str); err != nil {
		panic(err)
	} else {
		return num
	}
}
