package utils

import "time"

func FormatDateTimeToDateOnly(fullDateTime *string) (string, error) {
	dateTime, err := time.Parse(time.RFC3339Nano, *fullDateTime)
	if err != nil {
		return "", err
	}
	formatted := dateTime.Format("02/01/2006")
	return formatted, nil
}
