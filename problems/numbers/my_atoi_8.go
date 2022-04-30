package numbers

import "math"

// Read in and ignore any leading whitespace.
// Check if the next character (if not already at the end of the string) is '-' or '+'. Read this character in if it is either. This determines if the final result is negative or positive respectively. Assume the result is positive if neither is present.
// Read in next the characters until the next non-digit character or the end of the input is reached. The rest of the string is ignored.
// Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no digits were read, then the integer is 0. Change the sign as necessary (from step 2).
// If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then clamp the integer so that it remains in the range. Specifically, integers less than -231 should be clamped to -231, and integers greater than 231 - 1 should be clamped to 231 - 1.
// Return the integer as the final result.
func myAtoi(s string) int {
	res := 0
	flag := 1
	start := false
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			if !start {
				start = true
			}
			if res*flag > math.MaxInt32/10 {
				return math.MaxInt32
			}
			if res*flag < math.MinInt32/10 {
				return math.MinInt32
			}
			cur := (int(s[i]) - int('0'))
			if flag == 1 && math.MaxInt32-res*10 < cur {
				return math.MaxInt32
			}
			if flag == -1 && flag*res*10-math.MinInt32 < cur {
				return math.MinInt32
			}
			res = res*10 + cur
		} else if !start && s[i] == ' ' {
			continue
		} else if !start && s[i] == '-' {
			flag = -1
			start = true
		} else if !start && s[i] == '+' {
			start = true
		} else {
			break
		}
	}
	return flag * res
}
