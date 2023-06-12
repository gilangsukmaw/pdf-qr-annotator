package pdigenerator

import (
	"bytes"
	"fmt"
	pdfcpuapi "github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"io"
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
