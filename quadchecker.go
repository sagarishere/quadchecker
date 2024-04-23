package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type QuadError string

func (e QuadError) Error() string {
	return string(e)
}

func getPreviousOutput() string {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		println("Failed to read input")
	}
	return string(input)
}

func execCommand(quad, x, y string) string {
	cmd := exec.Command(quad, x, y)
	output, err := cmd.Output()
	if err != nil {
		println("Failed to execute command")
	}
	return string(output)
}

func getDimensions(str string) (x, y int, err error) {
	lines := strings.Split(str, "\n")
	lines = lines[:len(lines)-1]
	if len(lines) == 0 {
		return 0, 0, QuadError("Empty string")
	}
	y, x = len(lines), len(lines[0])
	if strings.Contains(lines[0], " ") {
		return 0, 0, QuadError("Invalid character")
	}
	for _, line := range lines {
		if len(lines[0]) != len(line) {
			return 0, 0, QuadError("Different lengths")
		}
	}
	return
}

func findSameResults(x int, y int, resultOfGiven string) (sameQuads []string) {
	files := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"}
	strX, strY := strconv.Itoa(x), strconv.Itoa(y)
	for _, elem := range files {
		if resultOfGiven == execCommand(elem, strX, strY) {
			sameQuads = append(sameQuads, elem[2:])
		}
	}
	return
}

func quadChecker(previousOutput string) {
	x, y, err := getDimensions(previousOutput)
	if err != nil {
		println("Not a quad function")
		return
	}
	sameQuads := findSameResults(x, y, previousOutput)
	for i, elem := range sameQuads {
		fmt.Printf("[%s] [%d] [%d]", elem, x, y)
		if i != len(sameQuads)-1 {
			print(" || ")
		}
	}
	println()
}

func main() {
	previousOutput := getPreviousOutput()
	if len(previousOutput) > 0 {
		quadChecker(previousOutput)
	} else {
		println("Not a quad function")
	}
}
