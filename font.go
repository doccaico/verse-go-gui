package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func LoadFont(size float64) (text.Face, error) {
	fontdata, err := os.ReadFile(filepath.Join("resources", "fonts", "NotoSansJP-Regular.ttf"))
	if err != nil {
		t.Fatal(err)
	}

	s, err := text.NewGoTextFaceSource(fontdata)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &text.GoTextFace{
		Source: s,
		Size:   size,
	}, nil
}
