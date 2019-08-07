package utilits

import "strings"

func Replacer(s string) string {

	r := strings.NewReplacer(" ", "",
		"-", "",
		".", "")
	s = r.Replace(s)

	if strings.HasPrefix(s, "0") {
		s = strings.TrimLeft(s, "0")
		s = "0" + s
	}

	return s
}
