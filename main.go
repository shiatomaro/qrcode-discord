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

	userID := flag.String("user", "", "Discord User ID")
	invite := flag.String("invite", "", "Discord Invite Code")
	output := flag.String("o", "discord_qr.png", "Output filename")
	size := flag.Int("size", 256, "QR Code size in pixels")
	flag.Parse()

	var input string
	if *userID == "" && *invite == "" {
		fmt.Print("Enter Discord User ID or Invite Code: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
	} else {
		if *userID != "" {
			input = *userID
		} else {
			input = *invite
		}
	}
	if input == "" {
		fmt.Println("Error: No Input Provided")
		os.Exit(1)
	}
	err := qrgen.GenerateQR(input, *output, *size)
	if err != nil {
		log.Fatal(err)
	}
}
