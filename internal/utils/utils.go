package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func CodeString(v any) string {
	switch a := v.(type) {
	case int:
		return strconv.Itoa(a)
	case string:
		return WithQuotes(a)
	default:
		return fmt.Sprint(a)
	}
}

func WithQuotes(s string) string {
	return `"` + s + `"`
}

func TrimQuotes(s string) string {
	return strings.Trim(s, `"`)
}
