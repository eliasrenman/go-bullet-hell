package assets

import (
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Font struct {
	font.Face
}

func LoadFont(path string, op opentype.FaceOptions) *Font {
	data, err := Assets.ReadFile("data/" + path)
	if err != nil {
		log.Fatal(err)
	}

	tt, err := opentype.Parse(data)
	if err != nil {
		log.Fatal(err)
	}

	if op.Size == 0 {
		op = opentype.FaceOptions{
			Size:    12,
			DPI:     72,
			Hinting: font.HintingFull,
		}
	}

	fnt, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	return &Font{fnt}
}
