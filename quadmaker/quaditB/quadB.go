package main

import (
	"os"
	q "quad"
)

var quadB = q.NewQuad("/", "\\", "\\", "/", "*", "*", " ")

func QuadB(x, y int) {
	os.Stdout.WriteString(q.GenerateQuad(x, y, *quadB))
}

func main() {
	x, y := q.ProcessArgs()
	QuadB(x, y)
}
