package pdigenerator

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

type Output struct {
	file []byte
}

func (o *Output) ExportToPDF() error {
	permissions := 0644 // or whatever you need
	err := ioutil.WriteFile("./output/output.pdf", o.file, fs.FileMode(permissions))
	if err != nil {
		fmt.Println("Couldn't export output: %w", err)
		return err
	}

	return nil
}
