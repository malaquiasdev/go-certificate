package utils

import (
	"time"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func FormatDateTimeToDateOnly(fullDateTime *string) (string, error) {
	dateTime, err := time.Parse(time.RFC3339Nano, *fullDateTime)
	if err != nil {
		return "", err
	}
	formatted := dateTime.Format("02/01/2006")
	return formatted, nil
}

func GetDateTimeNowFormatted() string {
	return time.Now().Format(time.RFC3339Nano)
}

func NormalizeString(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}
