package quad

import (
	"os"
	"strconv"
	"strings"
)

func GenerateQuad(x, y int, q Quad) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	rows := make([]string, y) // Create a slice of strings to hold each row

	for row := 1; row <= y; row++ {
		var currentRow strings.Builder // Use strings.Builder for efficient string concatenation

		for column := 1; column <= x; column++ {
			switch {
			case row == 1 && column == 1: // Top-left corner
				currentRow.WriteString(q.leftCorner)
			case row == 1 && column == x: // Top-right corner
				currentRow.WriteString(q.rightCorner)
			case row == y && column == 1: // Bottom-left corner
				currentRow.WriteString(q.leftBottomCorner)
			case row == y && column == x: // Bottom-right corner
				currentRow.WriteString(q.rightBottomCorner)
			case row == 1 || row == y: // Top or bottom rows
				currentRow.WriteString(q.topBottomRowFill)
			case column == 1 || column == x: // Left or right columns
				currentRow.WriteString(q.leftRightColumnFill)
			default: // Middle of the quad
				currentRow.WriteString(q.middleOfQuadFill)
			}
		}
		rows[row-1] = currentRow.String() // Store the complete row
	}

	return strings.Join(rows, "\n") + "\n"
}

func ProcessArgs() (x, y int) {
	if len(os.Args) != 3 {
		printUsage()
		os.Exit(1)
	}
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		println(err)
		printUsage()
		os.Exit(1)
	}
	y, err = strconv.Atoi(os.Args[2])
	if err != nil {
		println(err)
		printUsage()
		os.Exit(1)
	}
	return
}

func printUsage() {
	println("Usage: quadE <width: int> <height: int>")
}
