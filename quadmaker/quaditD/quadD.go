package main

import (
	"os"
	q "quad"
)

var quadD = q.NewQuad("A", "C", "A", "C", "B", "B", " ")

func QuadD(x, y int) {
	os.Stdout.WriteString(q.GenerateQuad(x, y, *quadD))
}

func main() {
	x, y := q.ProcessArgs()
	QuadD(x, y)
}
