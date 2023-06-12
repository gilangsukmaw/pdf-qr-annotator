package pdigenerator

import (
	"bytes"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type FileConfig struct {
	Src               string
	TargetPath        string
	ResultName        string
	Password          string
	EncryptedPdf      bool
	UpdateMetadata    bool
	GenerateThumbnail bool
	QrPosition        string
	QrPage            int
	Pos               int
	Dx                int
	Dy                int
}

func SetWaterMark(config *FileConfig, qr []byte) *model.Watermark {
	wm := model.DefaultWatermarkConfig()
	wm.PdfRes = map[int]model.PdfResources{}
	wm.Mode = 1
	wm.Page = config.QrPage
	wm.Image = bytes.NewReader(qr)
	wm.Update = false
	wm.OnTop = true
	wm.Rotation = 0
	wm.Diagonal = 0
	wm.ScaleAbs = true
	wm.Pos = types.Anchor(config.Pos)
	wm.Dx = float64(config.Dx)
	wm.Dy = float64(config.Dy)

	return wm
}
