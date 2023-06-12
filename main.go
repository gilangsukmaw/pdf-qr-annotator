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
		ResultName:        "asu",
		EncryptedPdf:      false,
		UpdateMetadata:    false,
		GenerateThumbnail: false,
		//WatermarkPosition: "br", // bottom right
		Dx: -10,
		Dy: 10,
	}
	err := pdigenerator.PDIgenerate(&config)

	if err != nil {
		fmt.Println("error generating pdf")
	}
}
