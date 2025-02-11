package utils

import (
	"strings"
	"regexp"
)

func HtmlTrim(source string) []string {
	// 1. Regex patterns
	// 2. Replace
	// 3. Return slice

	// 1
	pattern_old := []string{
		"<", "\t", "  *", "^ ", " $",
	}
	pattern_new := []string{
		"\n<", " ", " ", "", "",
	}

	// 2
	res := func(text *string, in, out *[]string, res []string, pos int) []string {
		var fn func(*string, *[]string, *[]string, []string, int) []string

		fn = func(text *string, in, out *[]string, res []string, pos int) []string {
			if pos > len(*in) - 1 {
				res = strings.Split(*text, "\n")
				return res
			}

			re := regexp.MustCompile((*in)[pos])
			*text = re.ReplaceAllString(*text, (*out)[pos])

			return fn(text, in, out, res, pos + 1)
		}

		return fn(text, in, out, res, pos)
	}(&source, &pattern_old, &pattern_new, []string{}, 0)

	// 3
	return res
}
