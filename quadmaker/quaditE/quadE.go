package main

import (
	"os"
	q "quad"
)

var quadE = q.NewQuad("A", "C", "C", "A", "B", "B", " ")

func QuadE(x, y int) {
	os.Stdout.WriteString(q.GenerateQuad(x, y, *quadE))
}

func main() {
	x, y := q.ProcessArgs(os.Args)
	QuadE(x, y)
}
