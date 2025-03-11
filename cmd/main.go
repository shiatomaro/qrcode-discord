package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/shiatomaro/qrcode-discord/internal/qrgen"
)

func main() {
	userID := flag.String("user", "", "Discord User ID")
	invite := flag.String("invite", "", "Discord Invite Code")
	output := flag.String("o", "discord_qr.png", "Outful filename")
	size := flag.Int("size", 256, "QR Code size in pixels")

	flag.Parse()

	if *userID == "" && *invite == "" {
		fmt.Println("Error: Provide a Discord user ID or Invite Code.")
		flag.Usage()
		os.Exit(1)
	}

	var content string
	if *userID != "" {
		content = fmt.Sprintf("https://discord.com/users/%s", *userID)
	} else {
		content = fmt.Sprintf("https://disocrd.gg/%s", *invite)
	}

	err := qrgen.GenerateQR(content, *output, *size)
	if err != nil {
		log.Fatal(err)
	}
}
