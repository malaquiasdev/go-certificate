package certificate

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

type Field struct {
	Key  string
	Text FieldText
}

type FieldText struct {
	PositionX int
	PositionY int
	FontSize  float64
	FontPath  string
	Value     string
}

func Generate(imgPath string, field Field) (image.Image, error) {
	bgImage, err := gg.LoadImage(imgPath)
	if err != nil {
		return nil, err
	}

	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)

	if err := dc.LoadFontFace(field.Text.FontPath, field.Text.FontSize); err != nil {
		return nil, err
	}

	// x := float64(imgWidth / 2)
	// y := float64((imgHeight / 2) - 80)
	// maxWidth := float64(imgWidth) - 60.0
	dc.SetColor(color.Black)
	dc.DrawString(field.Text.Value, float64(field.Text.PositionX), float64(field.Text.PositionY))
	// dc.DrawStringWrapped(field.Text.Value, x, y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)

	return dc.Image(), nil
}
