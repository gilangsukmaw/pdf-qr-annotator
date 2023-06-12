package pdigenerator

import (
	"bytes"
	"fmt"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"io"
)

func GenerateQR(uuid string) ([]byte, error) {

	qrc, err := qrcode.New(fmt.Sprintf(`https://privy.id/verify/%s`, uuid))
	// qrc, err := qrcode.New("with-custom-shape", qrcode.WithCircleShape())
	if err != nil {
		return nil, err
	}

	// get bytes
	buf := bytes.NewBuffer(nil)
	wr := nopCloser{Writer: buf}
	w2 := standard.NewWithWriter(
		wr,
		standard.WithQRWidth(5),
		standard.WithBorderWidth(1),
		standard.WithLogoImageFilePNG("./logo/logo_25_pdi.png"),
		standard.WithCircleShape(),
	)
	if err = qrc.Save(w2); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type nopCloser struct {
	io.Writer
}

func (nopCloser) Close() error { return nil }
