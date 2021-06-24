package main

func isValid(s string) bool {
	var stk []byte
	for _, c := range s {
		if c == ')' {
			if len(stk) == 0 || stk[len(stk)-1] != '(' {
				return false
			} else {
				stk = stk[:len(stk)-1]
			}
		} else if c == ']' {
			if len(stk) == 0 || stk[len(stk)-1] != '[' {
				return false
			} else {
				stk = stk[:len(stk)-1]
			}
		} else if c == '}' {
			if len(stk) == 0 || stk[len(stk)-1] != '{' {
				return false
			} else {
				stk = stk[:len(stk)-1]
			}
		} else {
			stk = append(stk, byte(c))
		}
	}
	return len(stk) == 0
}
