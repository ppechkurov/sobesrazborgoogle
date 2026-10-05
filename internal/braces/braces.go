package braces

func compare(str []rune) bool {
	stack := make([]rune, 0, len(str))
	for _, r := range str {
		switch r {
		case '(', '{', '[':
			stack = append(stack, r)
		case ')', '}', ']':
			if last := len(stack) - 1; last >= 0 {
				stack = stack[:last]
			}
		}
	}

	return len(stack) == 0
}
