package image_draw

import (
	"bytes"

	"github.com/signintech/gopdf"
)

func ToPdf(coverImg []byte, backCoverImg []byte) (*bytes.Buffer, error) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4Landscape})

	err := addPageWithImage(&pdf, coverImg, gopdf.PageSizeA4Landscape.W, gopdf.PageSizeA4Landscape.H)
	if err != nil {
		return nil, err
	}

	err = addPageWithImage(&pdf, backCoverImg, gopdf.PageSizeA4Landscape.W, gopdf.PageSizeA4Landscape.H)
	if err != nil {
		return nil, err
	}

	pdfBuffer := bytes.NewBuffer([]byte{})
	if _, err := pdf.WriteTo(pdfBuffer); err != nil {
		return nil, err
	}

	return pdfBuffer, nil
}

func addPageWithImage(pdf *gopdf.GoPdf, imageData []byte, width float64, height float64) error {
	img, err := gopdf.ImageHolderByBytes(imageData)
	if err != nil {
		return err
	}

	pdf.AddPage()
	pdf.ImageByHolderWithOptions(img,
		gopdf.ImageOptions{
			Rect: &gopdf.Rect{
				W: width,
				H: height,
			},
		})

	return nil
}