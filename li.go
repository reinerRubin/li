package li

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FrequenciesToBarChat(fr Frequencies, barChatWidth int) string {
	var (
		maxCount = len(fmt.Sprintf("%d", fr.maxCount()))
		acc      = strings.Builder{}
	)

	if len(fr) == 0 {
		return ""
	}

	biggest := fr[0].Count
	count := func(quantity int) int {
		w := quantity * barChatWidth / biggest
		if w == 0 {
			w = 1 // for aesthetic
		}
		return w
	}

	format := "%" + fmt.Sprintf("%d", maxCount) + "d"
	for _, f := range fr {
		acc.WriteString(fmt.Sprintf(" "+format+" ", f.Count))

		rightOffset := count(f.Count)
		for i := 0; i < rightOffset; i++ {
			acc.WriteString("■")
		}
		for i := 0; i < barChatWidth-rightOffset; i++ {
			acc.WriteString(" ")
		}
		acc.WriteString(fmt.Sprintf("  %s\n", f.Item))
	}

	return acc.String()
}

func ReadStdin() ([]string, error) {
	var (
		lines   = make([]string, 0)
		scanner = bufio.NewScanner(os.Stdin)
	)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func BarChatSectionWidth() (int, error) {
	return 30, nil
}
