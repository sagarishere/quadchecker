package main

import (
	"os"
	q "quad"
)

var quadA = q.NewQuad("o", "o", "o", "o", "-", "|", " ")

func QuadA(x, y int) {
	os.Stdout.WriteString(q.GenerateQuad(x, y, *quadA))
}

func main() {
	x, y := q.ProcessArgs()
	QuadA(x, y)
}
