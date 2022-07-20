package substr

import "strings"

func Sub(s string, trail string, head ...string) string {
	for _, h := range head {
		i := strings.Index(s, h)
		if i < 0 {
			return ""
		}
		s = s[i+len(h):]
	}
	if trail != "" {
		i := strings.Index(s, trail)
		if i < 0 {
			return ""
		}
		s = s[:i]
	}
	return s + ""
}

func RSub(s string, head string, trail ...string) string {
	for _, t := range trail {
		i := strings.LastIndex(s, t)
		if i < 0 {
			return ""
		}
		s = s[:i]
	}
	if head != "" {
		i := strings.LastIndex(s, head)
		if i < 0 {
			return ""
		}
		s = s[i+len(head):]
	}
	return s + ""
}
