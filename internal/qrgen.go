package qrgen

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func GenerateQR(content string, filename string, size int) error {
	err := qrcode.WriteFile(content, qrcode.Medium, size, filename)
	if err != nil {
		return err
	}
	fmt.Println("QR Code saved as", filename)
	return nil
}
