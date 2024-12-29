package rotatebraket

// func solution(s string) int {
// 	cnt := 0
// 	brakets := strings.Split(s, "")

// 	var verify func(strArr []string) bool

// 	verify = func(strArr []string) bool {
// 		verifier := make(map[string][]int)

// 		for idx, sa := range strArr {
// 			switch sa {
// 			case "}", "]", ")":
// 				r := reverseBraket(sa)
// 				if len(verifier[r]) == 0 {
// 					return false
// 				}
// 				v := verifier[r][len(verifier[r])-1]
// 				if !verify(strArr[v+1 : idx]) {
// 					return false
// 				}
// 				verifier[r] = verifier[r][:len(verifier[r])-1]
// 			case "{", "[", "(":
// 				verifier[sa] = append(verifier[sa], idx)
// 			}
// 		}

// 		flag := true
// 		for _, v := range verifier {
// 			if len(v) != 0 {
// 				flag = false
// 			}
// 		}

// 		return flag
// 	}

// 	for i := 0; i < len(brakets); i++ {
// 		if verify(brakets) {
// 			cnt++
// 		}
// 		brakets = append(brakets, brakets[0])
// 		brakets = brakets[1:]
// 	}
// 	return cnt
// }

// func reverseBraket(b string) string {
// 	switch b {
// 	case "}":
// 		return "{"
// 	case "]":
// 		return "["
// 	case ")":
// 		return "("
// 	}
// 	panic("invalid braket")
// }

func solution(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if isValid(s[i:] + s[:i]) {
			result++
		}
	}
	return result
}

func isValid(s string) bool {
	var stack []rune
	for _, c := range s {
		switch c {
		case '[', '{', '(':
			stack = append(stack, c)
			continue
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
