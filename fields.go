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
		return minuteRange(input)
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

func minuteRange(input string) (string, error) {
	span := strings.Split(input, `-`)

	from, err := strconv.Atoi(span[0])
	if err != nil {
		return "", fmt.Errorf("unable to parse range '%s' : %v", input, err)
	}

	to, err := strconv.Atoi(span[1])
	if err != nil {
		return "", fmt.Errorf("unable to parse range '%s' : %v", input, err)
	}

	var sb strings.Builder
	for i := from; i <= to; i++ {
		sb.WriteString(fmt.Sprintf("%d ", i))
	}
	return strings.Trim(sb.String(), " "), nil

}
