package imagedraw

import (
	"bytes"
	"image"
	"image/color"
	"image/png"

	"log"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
)

type Field struct {
	Key  string
	Text FieldText
}

type FieldText struct {
	PositionX int
	PositionY int
	FontSize  float64
	FontBytes []byte
	Value     string
}

func Draw(imgBytes []byte, fields []Field) image.Image {
	img, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		log.Fatal("ERROR: decode image bytes failed ", err)
		panic(err)
	}

	dc := gg.NewContextForImage(img)
	dc.DrawImage(img, 0, 0)

	for _, field := range fields {

		f, err := truetype.Parse(field.Text.FontBytes)
		if err != nil {
			log.Fatal("ERROR: parse font failed ", err)
			panic(err)
		}

		dc.SetFontFace(truetype.NewFace(f, &truetype.Options{Size: field.Text.FontSize}))
		dc.SetColor(color.Black)
		dc.DrawString(field.Text.Value, float64(field.Text.PositionX), float64(field.Text.PositionY))
	}

	return dc.Image()
}

func DrawAndEconde(imgBytes []byte, fields []Field) []byte {
	imageDraw := Draw(imgBytes, fields)
	b := new(bytes.Buffer)
	if err := png.Encode(b, imageDraw); err != nil {
		log.Fatal("ERROR: unable to encode image ", err)
		panic(err)
	}
	return b.Bytes()
}
