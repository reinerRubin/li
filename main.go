package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	if err := app(); err != nil {
		log.Fatal(err)
	}
}

func app() error {
	lines, err := in()
	if err != nil {
		return err
	}

	frequencies, err := newFrequencies(lines)
	if err != nil {
		return err
	}

	width, err := terminalWidth()
	if err != nil {
		return err
	}

	histogram := frequenciesToHistogram(frequencies, width)
	fmt.Print(histogram)

	return nil
}

func frequenciesToHistogram(fr frequencies, width int) string {
	var (
		maxCount = len(fmt.Sprintf("%d", fr.maxCount()))
		acc      = ""
	)

	if len(fr) == 0 {
		return acc
	}

	biggest := fr[0].count
	count := func(quantity int) int {
		w := quantity * width / biggest
		if w == 0 {
			w = 1 // for aesthetic
		}
		return w
	}

	for _, f := range fr {
		format := "%" + fmt.Sprintf("%d", maxCount) + "d"
		acc += fmt.Sprintf(" "+format+" ", f.count)

		rightOffset := count(f.count)
		for i := 0; i < rightOffset; i++ {
			acc += "■"
		}
		for i := 0; i < width-rightOffset; i++ {
			acc += " "
		}
		acc += fmt.Sprintf("  %s\n", f.item)
	}

	return acc
}

type frequency struct {
	item  string
	count int
}

type frequencies []frequency

func (v frequencies) maxString() string {
	r := ""
	for _, frequency := range v {
		if len(frequency.item) > len(r) {
			r = frequency.item
		}
	}

	return r
}

func (v frequencies) maxCount() int {
	c := 0
	for _, frequency := range v {
		if frequency.count > c {
			c = frequency.count
		}
	}

	return c
}

func (v frequencies) Len() int {
	return len(v)
}

func (v frequencies) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v frequencies) Less(i, j int) bool {
	if v[i].count == v[j].count {
		return v[i].item > v[j].item
	}

	return v[i].count > v[j].count
}

var uniqCRegexp = regexp.MustCompile(`^\s*(\d+)\s+(.*)$`)

func newFrequencies(lines []string) (frequencies, error) {
	fr := make(frequencies, 0, len(lines))

	for _, line := range lines {
		f, err := newFrequency(line)
		if err != nil {
			return nil, err
		}

		fr = append(fr, f)
	}

	sort.Sort(fr)

	return fr, nil
}

func newFrequency(line string) (frequency, error) {
	m := uniqCRegexp.FindStringSubmatch(line)
	if len(m) != 3 {
		return frequency{}, fmt.Errorf("failed to parse %s: %+v", line, m)
	}

	count, err := strconv.Atoi(m[1])
	if err != nil {
		return frequency{}, err
	}

	return frequency{
		item:  m[2],
		count: count,
	}, nil
}

func in() ([]string, error) {
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

func terminalWidth() (int, error) {
	return 30, nil
}

func times(str string, n int) (out string) {
	for i := 0; i < n; i++ {
		out += str
	}
	return
}

func rightPad(str string, pad string, length int) string {
	return str + times(pad, length-len(str))
}
