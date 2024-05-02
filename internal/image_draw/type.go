package image_draw

type Font struct {
	Size float64
	File []byte
}

type Position struct {
	X int
	Y int
}

type FieldText struct {
	Position Position
	Font     Font
	Value    string
}

type DrawParams struct {
	Key  string
	Text FieldText
}
