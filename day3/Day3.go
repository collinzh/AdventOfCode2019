package day3

import (
	"errors"
	"fmt"
	"github.com/collinzh/AdventOfCode2019/util"
	"strings"
)

type D3Solution struct {
}

func parseMoves() [][]util.Move {
	f := util.MustOpenFile("day3.txt")
	lines := util.ScanToStringSlices(f)
	ret := make([][]util.Move, len(lines))
	for idx, line := range lines {
		moves := strings.Split(line, ",")
		parsed := make([]util.Move, len(moves))
		for mIdx, move := range moves {
			parsed[mIdx] = util.Move{
				Direction: util.ParseDirection(string(move[0])),
				Step:      util.MustAtoi(move[1:]),
			}
		}
		ret[idx] = parsed
	}

	return ret
}

func makeCanvas(allMoves [][]util.Move) map[string]int {
	canvas := make(map[string]int)

	x, y := 0, 0
	counter := 0
	moves := allMoves[0]
	for _, move := range moves {
		stepX, stepY := steps(move.Direction)

		for i := 0; i < move.Step; i++ {
			x += stepX
			y += stepY
			counter++
			canvas[posName(x, y)] = counter
		}
	}

	moves = allMoves[1]
	x, y = 0, 0
	counter = 0
	minDistance := util.MaxInteger
	minSteps := util.MaxInteger
	for _, move := range moves {
		stepX, stepY := steps(move.Direction)

		for i := 0; i < move.Step; i++ {
			x += stepX
			y += stepY
			counter++
			steps := canvas[posName(x, y)]
			if steps > 0 {
				dist := util.IntegerAbs(x) + util.IntegerAbs(y)
				if dist < minDistance {
					minDistance = dist
				}
				if steps+counter < minSteps {
					minSteps = steps + counter
				}
			}
		}
	}

	fmt.Printf("Min distance: %d\n", minDistance)
	fmt.Printf("Min steps: %d\n", minSteps)

	return canvas
}

func posName(x, y int) string {
	return fmt.Sprintf("%d_%d", x, y)
}

func steps(dir util.Direction) (stepX, stepY int) {
	switch dir {
	case util.Up:
		stepX = 0
		stepY = 1
	case util.Down:
		stepX = 0
		stepY = -1
	case util.Right:
		stepX = 1
		stepY = 0
	case util.Left:
		stepX = -1
		stepY = 0
	case util.None:
		panic(errors.New("unrecognized move"))
	}
	return
}

func (d D3Solution) RunSolutionP1() {
	makeCanvas(parseMoves())
}

func (d D3Solution) RunSolutionP2() {
}
