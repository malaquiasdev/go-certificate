package image_draw

import (
	"bytes"
	"image"
	"image/color"
	"image/png"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
)

func Png(imgBytes []byte, params []DrawParams) ([]byte, error) {
	imageDraw, err := draw(imgBytes, params)
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)
	if err := png.Encode(b, imageDraw); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func draw(imgBytes []byte, params []DrawParams) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return nil, err
	}

	dc := gg.NewContextForImage(img)
	dc.DrawImage(img, 0, 0)

	for _, field := range params {

		f, err := truetype.Parse(field.Text.Font.File)
		if err != nil {
			return nil, err
		}

		dc.SetFontFace(truetype.NewFace(f, &truetype.Options{Size: field.Text.Font.Size}))
		dc.SetColor(color.Black)
		dc.DrawString(field.Text.Value, float64(field.Text.Position.X), float64(field.Text.Position.Y))
	}

	return dc.Image(), nil
}
