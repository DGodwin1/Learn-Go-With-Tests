package Iteration

// Takes a string and does nice things
func Repeat(c string, n int) string{
	final := ""
	for i := 0; i < n; i++ {
		final += c
	}
	return final
}

