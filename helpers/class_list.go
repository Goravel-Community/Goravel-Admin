package helpers

import "strings"

func ClassList(classes map[string]bool) (res string) {
	res = ""
	for class, condition := range classes {
		if condition {
			res += " " + class
		}
	}
	return strings.TrimSpace(res)
}
