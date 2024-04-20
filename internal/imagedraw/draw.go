package imagedraw

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"
	"log"
	"strings"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

func GenerateImage(textContent string, fgColorHex string, bgColorHex string, fontSize float64, fontBytes []byte, imgBytes []byte) ([]byte, error) {

	/*
		fgColor := color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff} // Default font color
		if len(fgColorHex) == 7 {
			_, err := fmt.Sscanf(fgColorHex, "#%02x%02x%02x", &fgColor.R, &fgColor.G, &fgColor.B)
			if err != nil {
				log.Println(err)
				fgColor = color.RGBA{R: 0x2e, G: 0x34, B: 0x36, A: 0xff}
			}
		}

		bgColor := color.RGBA{R: 0x30, G: 0x0a, B: 0x24, A: 0xff} // Default background color
		if len(bgColorHex) == 7 {
			_, err := fmt.Sscanf(bgColorHex, "#%02x%02x%02x", &bgColor.R, &bgColor.G, &bgColor.B)
			if err != nil {
				log.Println(err)
				bgColor = color.RGBA{R: 0x30, G: 0x0a, B: 0x24, A: 0xff}
			}
		}
	*/

	loadedFont, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Fatal("ERROR: parse font failed ", err)
		return nil, err
	}

	code := strings.Replace(textContent, "\t", "    ", -1) // convert tabs into spaces
	text := strings.Split(code, "\n")                      // split newlines into arrays

	/*
		fg := image.NewUniform(fgColor)
		bg := image.NewUniform(bgColor)
		rgba := image.NewRGBA(image.Rect(0, 0, 1200, 630))
		draw.Draw(rgba, rgba.Bounds(), bg, image.Pt(0, 0), draw.Src)
	*/

	// Decode the provided image bytes into an image.RGBA type
	img, err := png.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		log.Fatal("ERROR: decode image bytes failed ", err)
		return nil, err
	}

	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, rgba.Bounds(), img, image.Pt(0, 0), draw.Src)

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(loadedFont)
	c.SetFontSize(fontSize)
	// c.SetClip(rgba.Bounds())
	c.SetClip(bounds)
	// c.SetDst(rgba)
	c.SetDst(rgba)
	// c.SetSrc(fg)
	c.SetHinting(font.HintingNone)

	textXOffset := 50
	textYOffset := 10 + int(c.PointToFixed(fontSize)>>6) // Note shift/truncate 6 bits first

	pt := freetype.Pt(textXOffset, textYOffset)
	for _, s := range text {
		_, err = c.DrawString(strings.Replace(s, "\r", "", -1), pt)
		if err != nil {
			return nil, err
		}
		pt.Y += c.PointToFixed(fontSize * 1.5)
	}

	b := new(bytes.Buffer)
	if err := png.Encode(b, rgba); err != nil {
		log.Fatal("ERROR: unable to encode image ", err)
		return nil, err
	}
	return b.Bytes(), nil
}
