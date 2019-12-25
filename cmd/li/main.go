package main

import (
	"fmt"
	"log"

	"github.com/reinerRubin/li"
)

func main() {
	if err := app(); err != nil {
		log.Fatal(err)
	}
}

func app() error {
	lines, err := li.ReadStdin()
	if err != nil {
		return err
	}

	frequencies, err := li.NewFrequencies(lines)
	if err != nil {
		return err
	}

	barChatSectionWidth, err := li.BarChatSectionWidth()
	if err != nil {
		return err
	}

	barChat := li.FrequenciesToBarChat(frequencies, barChatSectionWidth)
	fmt.Print(barChat)

	return nil
}
