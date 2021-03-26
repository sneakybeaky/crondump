package crondump

import (
	"errors"
	"strconv"
)

// Minute expands the input in cron syntax to show the minutes included
func Minute(input string) (string, error) {

	i, err := strconv.Atoi(input)

	if err != nil {
		return "", err
	}

	if i > 59 {
		return "", errors.New("minute can't be larger than 59")
	}

	return input, nil
}
