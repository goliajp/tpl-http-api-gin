package strx

import "strings"

func SnakeCase(str string) string {
	return toSnake(str)
}

func toSnake(s string) string {
	return ToDelimited(s, '_')
}

func ToSnakeWithIgnore(s string, ignore string) string {
	return ToScreamingDelimited(s, '_', ignore, false)
}

// ToScreamingSnake converts a string to SCREAMING_SNAKE_CASE
func ToScreamingSnake(s string) string {
	return ToScreamingDelimited(s, '_', "", true)
}

// ToKebab converts a string to kebab-case
func ToKebab(s string) string {
	return ToDelimited(s, '-')
}

// ToScreamingKebab converts a string to SCREAMING-KEBAB-CASE
func ToScreamingKebab(s string) string {
	return ToScreamingDelimited(s, '-', "", true)
}

// ToDelimited converts a string to delimited.snake.case
// (in this case `delimiter = '.'`)
func ToDelimited(s string, delimiter uint8) string {
	return ToScreamingDelimited(s, delimiter, "", false)
}

// ToScreamingDelimited converts a string to SCREAMING.DELIMITED.SNAKE.CASE
// (in this case `delimiter = '.'; screaming = true`)
// or delimited.snake.case
// (in this case `delimiter = '.'; screaming = false`)
func ToScreamingDelimited(s string, delimiter uint8, ignore string, screaming bool) string {
	s = strings.TrimSpace(s)
	n := strings.Builder{}
	n.Grow(len(s) + 2) // nominal 2 bytes of extra space for inserted delimiters
	for i, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		vIsNum := v >= '0' && v <= '9'
		if vIsLow && screaming {
			v += 'A'
			v -= 'a'
		} else if vIsCap && !screaming {
			v += 'a'
			v -= 'A'
		}
		if i+1 < len(s) {
			next := s[i+1]
			nextIsCap := next >= 'A' && next <= 'Z'
			nextIsLow := next >= 'a' && next <= 'z'
			nextIsNum := next >= '0' && next <= '9'
			if (vIsCap && (nextIsLow || nextIsNum)) || (vIsLow && nextIsCap) || (vIsNum && (nextIsCap || nextIsLow)) {
				prevIgnore := ignore != "" && i > 0 && strings.ContainsAny(string(s[i-1]), ignore)
				if !prevIgnore {
					if vIsCap && nextIsLow {
						if prevIsCap := i > 0 && s[i-1] >= 'A' && s[i-1] <= 'Z'; prevIsCap {
							n.WriteByte(delimiter)
						}
					}
					n.WriteByte(v)
					if vIsLow || vIsNum {
						n.WriteByte(delimiter)
					}
					continue
				}
			}
		}
		if (v == ' ' || v == '_' || v == '-' || v == '.') && !strings.ContainsAny(string(v), ignore) {
			n.WriteByte(delimiter)
		} else {
			n.WriteByte(v)
		}
	}

	return n.String()
}
