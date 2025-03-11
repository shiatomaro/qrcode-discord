package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/shiatomaro/qrcode-discord/qrgen"
)

func main() {
	userInput := flag.String("input", "", "Discord username/ID or invite code")
	urlType := flag.String("type", "invite", "Type of URL: 'invite' or 'user'")
	output := flag.String("o", "discord_qr.png", "Output filename")
	size := flag.Int("size", 256, "QR Code size in pixels")
	flag.Parse()

	var input string
	if *userInput == "" {
		fmt.Print("Enter Discord username/ID or invite code: ")
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		input = strings.TrimSpace(line)
	} else {
		input = *userInput
	}

	if input == "" {
		fmt.Println("Error: No Input Provided")
		os.Exit(1)
	}

	var urlTemplate string
	switch strings.ToLower(*urlType) {
	case "user":
		urlTemplate = "https://discord.com/users/%s"
	case "invite":
		urlTemplate = "https://discord.gg/%s"
	default:
		fmt.Println("Invalid type. Use 'user' or 'invite'.")
		os.Exit(1)
	}

	err := qrgen.GenerateQRWithTemplate(input, urlTemplate, *output, *size)
	if err != nil {
		log.Fatal(err)
	}
}
