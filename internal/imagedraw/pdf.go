package imagedraw

import (
	"bytes"
	"image"
	"image/png"
	"log"

	"github.com/signintech/gopdf"
)

func ImageToPdf(imageDraw image.Image, imageDraw2 image.Image) *bytes.Buffer {
	b := new(bytes.Buffer)
	if err := png.Encode(b, imageDraw); err != nil {
		log.Fatal("ERROR: unable to encode image ", err)
		panic(err)
	}

	img, _ := gopdf.ImageHolderByBytes(b.Bytes())

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4Landscape})
	pdf.AddPage()

	pdf.ImageByHolderWithOptions(img,
		gopdf.ImageOptions{
			Rect: &gopdf.Rect{
				W: 842,
				H: 595,
			},
		})

	b2 := new(bytes.Buffer)
	if err := png.Encode(b2, imageDraw2); err != nil {
		log.Fatal("ERROR: unable to encode image ", err)
		panic(err)
	}

	img2, _ := gopdf.ImageHolderByBytes(b2.Bytes())

	pdf.AddPage()
	// Draw the image onto the PDF page (replace with resized image if used)
	pdf.ImageByHolderWithOptions(img2,
		gopdf.ImageOptions{
			Rect: &gopdf.Rect{
				W: 842,
				H: 595,
			},
		})

	pdfBuffer := bytes.NewBuffer([]byte{})
	if _, err := pdf.WriteTo(pdfBuffer); err != nil {
		log.Fatal("ERROR: unable to convert PDF to buffer ", err)
		panic(err)
	}
	return pdfBuffer
}
