package qrgen

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func GenerateQR(input string, filename string, size int) error {
	if len(input) == 0 {
		return fmt.Errorf("no input provided")
	}

	var content string
	if len(input) > 15 {
		content = fmt.Sprintf("https://discord.com/users/%s", input)
	} else {
		content = fmt.Sprintf("https://discord.gg/%s", input)
	}

	err := qrcode.WriteFile(content, qrcode.Medium, size, filename)
	if err != nil {
		return err
	}

	fmt.Println("QR Code Saved As ", filename)
	return nil
}
