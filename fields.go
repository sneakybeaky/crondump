package crondump

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	maxMinute = 59
)

// Minute expands the input in cron syntax to show the minutes included
func Minute(input string) (string, error) {

	l, err := isList(input)

	if err != nil {
		return "", err
	}

	if l {
		ml, err := newMinuteList(input)
		if err != nil {
			return "", err
		}
		return ml.expand()
	}

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

	m, err := newMinute(input)

	if err != nil {
		return "", err
	}

	return m.expand()
}

func isRange(input string) (bool, error) {
	return regexp.MatchString(`[\d]+-[\d]+`, input)
}

func isList(input string) (bool, error) {
	return regexp.MatchString(`[\d]+,[\d]+`, input)
}

type minute struct {
	value string
}

func newMinute(cronexp string) (minute, error) {
	m, err := strconv.Atoi(cronexp)

	if err != nil {
		return minute{}, err
	}

	if m < 0 || m > maxMinute {
		return minute{}, errors.New("minute must be between 0 and 59")
	}

	return minute{value: cronexp}, nil
}

func (m minute) expand() (string, error) {
	return m.value, nil
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

	if to > maxMinute {
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

type minuteList struct {
	minutes []string
}

func newMinuteList(cronexp string) (*minuteList, error) {

	ml := &minuteList{}

	terms := strings.Split(cronexp, `,`)

	for _, term := range terms {
		ml.minutes = append(ml.minutes, term)
	}

	return ml, nil

}

func (ml *minuteList) expand() (string, error) {
	var sb strings.Builder

	for _, m := range ml.minutes {
		sb.WriteString(fmt.Sprintf("%s ", m))
	}

	return strings.TrimSpace(sb.String()), nil
}
