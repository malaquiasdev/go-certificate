package imagedraw

import (
	"bytes"
	"log"

	"github.com/signintech/gopdf"
)

func ImageToPdf(imageDraw []byte, imageDraw2 []byte) *bytes.Buffer {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4Landscape})

	addPageWithImage(&pdf, imageDraw, gopdf.PageSizeA4Landscape.W, gopdf.PageSizeA4Landscape.H)
	addPageWithImage(&pdf, imageDraw2, gopdf.PageSizeA4Landscape.W, gopdf.PageSizeA4Landscape.H)

	pdfBuffer := bytes.NewBuffer([]byte{})
	if _, err := pdf.WriteTo(pdfBuffer); err != nil {
		log.Fatal("ERROR: unable to convert PDF to buffer ", err)
		panic(err)
	}
	return pdfBuffer
}

func addPageWithImage(pdf *gopdf.GoPdf, imageData []byte, width float64, height float64) {
	img, err := gopdf.ImageHolderByBytes(imageData)
	if err != nil {
		log.Fatal("ERROR: unable to create image holder ", err)
	}
	pdf.AddPage()
	pdf.ImageByHolderWithOptions(img,
		gopdf.ImageOptions{
			Rect: &gopdf.Rect{
				W: width,
				H: height,
			},
		})
}
