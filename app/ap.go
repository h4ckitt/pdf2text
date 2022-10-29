package app

import (
	"bytes"
	"github.com/gen2brain/go-fitz"
	"github.com/otiai10/gosseract/v2"
	"image/png"
	"log"
	"strings"
)

//var ocrClient *gosseract.Client

func ProcessDoc(file []byte) (string, error) {
	var text strings.Builder

	client := gosseract.NewClient()

	doc, err := fitz.NewFromMemory(file)

	if err != nil {
		log.Println(err)
		return "", nil
	}

	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)

		if err != nil {
			log.Println(err)
			return "", err
		}

		buf := new(bytes.Buffer)

		//err = jpeg.Encode(buf, img, nil)

		err = png.Encode(buf, img)

		if err != nil {
			log.Println(err)
			return "", err
		}

		client.SetImageFromBytes(buf.Bytes())

		res, err := client.Text()

		if err != nil {
			return "", err
		}

		text.WriteString(res)
	}
	return text.String(), nil
}
