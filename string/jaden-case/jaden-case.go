package jadencase

import "strings"

func solution(s string) string {
	words := strings.Split(s, "")
	convert := make([]string, 0)

	temp := ""
	for _, word := range words {
		if word == " " {
			if temp != "" {
				convert = append(convert, temp)
			}
			convert = append(convert, " ")
			temp = ""
		} else {
			temp += word
		}
	}

	if temp != "" {
		convert = append(convert, temp)
	}

	for i, c := range convert {
		if c != " " {
			word := strings.ToLower(c)
			if !detectNumber(word[0]) {
				word = strings.ToUpper(string(word[0])) + word[1:]
			}
			convert[i] = word
		}
	}

	return strings.Join(convert, "")
}

func detectNumber(r byte) bool {
	return r >= '0' && r <= '9'
}
