package main

import (
	"fmt"
	"pdf/pdigenerator"
)

func main() {
	config := pdigenerator.FileConfig{
		Src:               "./example.png",
		TargetPath:        "./output",
		QrPage:            1,
		ResultName:        "woow",
		EncryptedPdf:      false,
		UpdateMetadata:    false,
		GenerateThumbnail: false,
		//WatermarkPosition: "br", // bottom right
		Pos: 8,
		Dx:  -10,
		Dy:  10,
	}
	file, err := pdigenerator.PDIgenerate(&config)

	if err != nil {
		fmt.Println("error generating pdf", err)
	}

	err = file.ExportToPDF()
	if err != nil {
		fmt.Println("error exporting pdf", err)
	}
}
