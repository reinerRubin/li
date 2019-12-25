package li

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

var uniqCRegexp = regexp.MustCompile(`^\s*(\d+)\s+(.*)$`)

type Frequency struct {
	Item  string
	Count int
}

// Frequencies is a slice of item/count pairs, must be sorted by count
type Frequencies []Frequency

func (v Frequencies) maxCount() (c int) {
	for _, frequency := range v {
		if frequency.Count > c {
			c = frequency.Count
		}
	}

	return
}

func (v Frequencies) Len() int {
	return len(v)
}

func (v Frequencies) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v Frequencies) Less(i, j int) bool {
	if v[i].Count == v[j].Count {
		return v[i].Item > v[j].Item
	}

	return v[i].Count > v[j].Count
}

func NewFrequencies(lines []string) (Frequencies, error) {
	fr := make(Frequencies, 0, len(lines))

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

func newFrequency(line string) (Frequency, error) {
	m := uniqCRegexp.FindStringSubmatch(line)
	if len(m) != 3 {
		return Frequency{}, fmt.Errorf("failed to parse %s: %+v", line, m)
	}

	count, err := strconv.Atoi(m[1])
	if err != nil {
		return Frequency{}, err
	}

	return Frequency{
		Item:  m[2],
		Count: count,
	}, nil
}
