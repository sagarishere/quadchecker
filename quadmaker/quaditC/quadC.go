package main

import (
	"os"
	q "quad"
)

var quadC = q.NewQuad("A", "A", "C", "C", "B", "B", " ")

func QuadC(x, y int) {
	os.Stdout.WriteString(q.GenerateQuad(x, y, *quadC))
}

func main() {
	x, y := q.ProcessArgs()
	QuadC(x, y)
}
