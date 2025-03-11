package qrgen

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func GenerateQRWithTemplate(input, urlTemplate, filename string, size int) error {
	if input == "" {
		return fmt.Errorf("no input provided")
	}

	url := fmt.Sprintf(urlTemplate, input)

	if err := qrcode.WriteFile(url, qrcode.Medium, size, filename); err != nil {
		return err
	}

	fmt.Println("QR Code saved as", filename)
	return nil
}
