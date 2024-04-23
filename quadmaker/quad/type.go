package quad

type Quad struct {
	leftCorner          string
	rightCorner         string
	leftBottomCorner    string
	rightBottomCorner   string
	topBottomRowFill    string
	leftRightColumnFill string
	middleOfQuadFill    string
}

func NewQuad(leftCorner, rightCorner, leftBottomCorner, rightBottomCorner, topBottomRowFill, leftRightColumnFill, middleOfQuadFill string) *Quad {
	return &Quad{
		leftCorner,
		rightCorner,
		leftBottomCorner,
		rightBottomCorner,
		topBottomRowFill,
		leftRightColumnFill,
		middleOfQuadFill,
	}
}
