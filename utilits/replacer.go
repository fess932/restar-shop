package utilits

import "strings"

func Replacer(s string) string {
	r := strings.NewReplacer(" ", "",
		"-", "",
		".", "")
	s = r.Replace(s)
	return s
}
