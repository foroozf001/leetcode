package main

func main() {}

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	match := len(s) > 0 && (s[0] == p[0] || p[0] == '?')
	if match {
		return isMatch(s[1:], p[1:])
	} else if p[0] == '*' {
		return isMatch(s, p[1:]) || (len(s) > 0 && isMatch(s[1:], p))
	}
	return false
}
