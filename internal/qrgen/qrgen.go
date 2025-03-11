package qrgen

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func GenerateQRFromInput(userInput string, filename string, size int) error {
	var content string
	if len(userInput) == 0 {
		return fmt.Errorf("no input provided")
	}
	if len(userInput) > 15 {
		content = fmt.Sprintf("https://discord.com/users/%s", userInput)
	} else {
		content = fmt.Sprintf("https://discord.gg/%s", userInput)
	}

	err := qrcode.WriteFile(content, qrcode.Medium, size, filename)
	if err != nil {
		return err
	}
	fmt.Println("QR Code saved as", filename)
	return nil
}
