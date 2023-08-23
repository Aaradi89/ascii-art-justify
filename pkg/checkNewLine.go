package pkg

func IsNewLine(input string) bool {
	for i, x := range input {
		if i > 0 && x == 'n' && input[i-1] == '\\' {
			return true
		}
	}

	return false
}