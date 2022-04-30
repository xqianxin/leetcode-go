package pslice

// We have two special characters:

// The first character can be represented by one bit 0.
// The second character can be represented by two bits (10 or 11).
// Given a binary array bits that ends with 0, return true if the last character must be a one-bit character.

// Example 1:

// Input: bits = [1,0,0]
// Output: true
// Explanation: The only way to decode it is two-bit character and one-bit character.
// So the last character is one-bit character.
// Example 2:

// Input: bits = [1,1,1,0]
// Output: false
// Explanation: The only way to decode it is two-bit character and two-bit character.
// So the last character is not one-bit character.

// Constraints:

// 1 <= bits.length <= 1000
// bits[i] is either 0 or 1.

func isOneBitCharacter(bits []int) bool {
	last := len(bits) - 1
	if bits[last] == 1 {
		return false
	}
	return isValid(bits[:last])
}

func isValid(bits []int) bool {
	if len(bits) == 0 {
		return true
	}
	if len(bits) == 1 {
		return bits[0] == 0
	}
	if len(bits) == 2 {
		return bits[1] == 0 || bits[0] == 1
	}
	end := len(bits) - 1
	if bits[end] == 1 && bits[end-1] == 0 {
		return false
	}
	if bits[end] == 1 && bits[end-1] == 1 {
		return isValid(bits[:end-1])
	}
	if bits[end] == 0 {
		return isValid(bits[:end]) || isValid(bits[:end-1])
	}
	return false
}
