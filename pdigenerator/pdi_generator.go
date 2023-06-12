package pdigenerator

import (
	"fmt"
	"github.com/google/uuid"
	pdfcpuapi "github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"os"
	"strings"
)

func PDIgenerate(config *FileConfig) (Output, error) {
	var (
		conf          = model.NewDefaultConfiguration()
		defaultImport = pdfcpu.DefaultImportConfig() //for processing image png or jpg
	)

	file, err := os.Open(config.Src)
	if err != nil {
		return Output{file: nil}, err
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return Output{file: nil}, err
	}

	filename := stat.Name()

	//set the qrcode
	id := uuid.New()
	qr, err := GenerateQR(id.String())
	if err != nil {
		fmt.Println("Couldn't pdigenerator the file: %w", err)
		return Output{file: nil}, err
	}

	//set document watermark
	wm := SetWaterMark(config, qr)

	//if file is pdf
	if strings.Contains(filename[len(filename)-5:], ".pdf") {
		output, err := AddWatermark(file, wm)

		if err != nil {
			fmt.Println("Couldn't add watermark to the file: %w", err)
			return Output{file: nil}, err
		}

		//err = ExportFile(output, id.String())
		//if err != nil {
		//	fmt.Println("Couldn't export the file the file: %w", err)
		//	return nil, err
		//}

		return Output{file: output}, nil
	}

	//if file is png or jpg
	if strings.Contains(filename[len(filename)-5:], ".jpg") || strings.Contains(filename[len(filename)-5:], ".png") {
		err := pdfcpuapi.ImportImagesFile([]string{config.Src}, fmt.Sprintf("./temp/%s-%s.pdf", config.ResultName, id.String()), defaultImport, conf)
		if err != nil {
			fmt.Println("Couldn't import image to pdf: %w", err)
			return Output{file: nil}, err
		}

		imageFilePdf, err := os.Open(fmt.Sprintf(`./temp/%s-%s.pdf`, config.ResultName, id.String()))
		if err != nil {
			return Output{file: nil}, err
		}

		defer file.Close()

		output, err := AddWatermark(imageFilePdf, wm)

		if err != nil {
			fmt.Println("Couldn't add watermark to the file: %w", err)
			return Output{file: nil}, err
		}

		err = os.Remove(fmt.Sprintf(`./temp/%s-%s.pdf`, config.ResultName, id.String()))
		if err != nil {
			fmt.Println("Couldn't remove the file: %w", err)
			return Output{file: nil}, err
		}

		//err = ExportFile(output, id.String())
		//if err != nil {
		//	fmt.Println("Couldn't export the file: %w", err)
		//	return nil, err
		//}

		return Output{file: output}, nil
	}

	return Output{file: nil}, err
}
