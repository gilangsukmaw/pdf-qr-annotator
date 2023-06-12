package pdigenerator

import (
	"bytes"
	"fmt"
	pdfcpuapi "github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"io"
	"io/fs"
	"io/ioutil"
)

func AddWatermark(file io.ReadSeeker, wm *model.Watermark) ([]byte, error) {
	output := new(bytes.Buffer)
	err := pdfcpuapi.AddWatermarks(file, output, nil, wm, nil)
	if err != nil {
		fmt.Println("Couldn't add watermark to doc: %w", err)
		return nil, err
	}

	return output.Bytes(), nil
}

func ExportFile(outputBytes []byte) error {
	permissions := 0644 // or whatever you need
	err := ioutil.WriteFile("./output/output.pdf", outputBytes, fs.FileMode(permissions))
	if err != nil {
		fmt.Println("Couldn't export output: %w", err)
		return err
	}

	return nil
}
