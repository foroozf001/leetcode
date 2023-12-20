package main

func main() {}

func IsMatch(s string, p string) bool {
	if len(p) == 0 { return len(s) == 0 }
	match := len(s) > 0 && (s[0] == p[0] || p[0] == '.')
	if len(p) > 1 && p[1] == '*' {
		return IsMatch(s, p[2:]) || (match && IsMatch(s[1:], p))
	}
	return match && IsMatch(s[1:], p[1:])
}
