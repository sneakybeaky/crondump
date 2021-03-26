package crondump

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Minute expands the input in cron syntax to show the minutes included
func Minute(input string) (string, error) {

	r, err := isRange(input)

	if err != nil {
		return "", err
	}

	if r {
		mr, err := newMinuteRange(input)
		if err != nil {
			return "", err
		}
		return mr.expand()
	}

	i, err := strconv.Atoi(input)

	if err != nil {
		return "", err
	}

	if i < 0 || i > 59 {
		return "", errors.New("minute must be between 0 and 59")
	}

	return input, nil
}

func isRange(input string) (bool, error) {
	return regexp.MatchString(`[\d]+-[\d]+`, input)
}

type minuteRange struct {
	from int
	to   int
}

func newMinuteRange(cronexp string) (*minuteRange, error) {
	span := strings.Split(cronexp, `-`)

	from, err := strconv.Atoi(span[0])
	if err != nil {
		return nil, fmt.Errorf("unable to parse range '%s' : %v", cronexp, err)
	}

	to, err := strconv.Atoi(span[1])
	if err != nil {
		return nil, fmt.Errorf("unable to parse range '%s' : %v", cronexp, err)
	}

	if to > 59 {
		return nil, fmt.Errorf("range %s ends too high at %d", cronexp, to)
	}

	return &minuteRange{
		from: from,
		to:   to,
	}, nil
}

func (mr *minuteRange) expand() (string, error) {
	var sb strings.Builder
	for i := mr.from; i <= mr.to; i++ {
		sb.WriteString(fmt.Sprintf("%d ", i))
	}
	return strings.TrimSpace(sb.String()), nil
}
