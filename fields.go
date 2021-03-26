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

	if i < 0 || i > 59 {
		return "", errors.New("minute must be between 0 and 59")
	}

	return input, nil
}
