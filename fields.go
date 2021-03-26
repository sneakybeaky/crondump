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

type expander interface {
	expand() (string, error)
}

// ExpandMinute expands the input in cron syntax to show the minutes included
func ExpandMinute(input string) (string, error) {

	switch {
	case isList(input):
		ml, err := newMinuteList(input)
		if err != nil {
			return "", err
		}
		return ml.expand()

	case isRange(input):
		mr, err := newMinuteRange(input)
		if err != nil {
			return "", err
		}
		return mr.expand()

	default:
		m, err := newMinute(input)

		if err != nil {
			return "", err
		}

		return m.expand()

	}
}

func isRange(input string) bool {
	reg := regexp.MustCompile(`[\d]+-[\d]+`)
	return reg.MatchString(input)
}

func isList(input string) bool {
	reg := regexp.MustCompile(`[\d]+,[\d]+`)
	return reg.MatchString(input)
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
	minutes []expander
}

func newMinuteList(cronexp string) (*minuteList, error) {

	ml := &minuteList{}

	terms := strings.Split(cronexp, `,`)

	for _, term := range terms {
		m, err := newMinute(term)

		if err != nil {
			return nil, fmt.Errorf("unable to parse list '%s' : %v", cronexp, err)
		}

		ml.minutes = append(ml.minutes, m)
	}

	return ml, nil

}

func (ml *minuteList) expand() (string, error) {
	var sb strings.Builder

	for _, m := range ml.minutes {

		s, err := m.expand()
		if err != nil {
			return "", fmt.Errorf("unable to expand list : %v", err)
		}

		sb.WriteString(fmt.Sprintf("%s ", s))
	}

	return strings.TrimSpace(sb.String()), nil
}
