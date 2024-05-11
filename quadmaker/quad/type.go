package quad

type Quad struct {
	leftTopCorner       string
	rightTopCorner      string
	leftBottomCorner    string
	rightBottomCorner   string
	topBottomRowFill    string
	leftRightColumnFill string
	middleOfQuadFill    string
}

func NewQuad(leftTopCorner, rightTopCorner, leftBottomCorner, rightBottomCorner, topBottomRowFill, leftRightColumnFill, middleOfQuadFill string) *Quad {
	return &Quad{
		leftTopCorner,
		rightTopCorner,
		leftBottomCorner,
		rightBottomCorner,
		topBottomRowFill,
		leftRightColumnFill,
		middleOfQuadFill,
	}
}
