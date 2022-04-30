package dynamicpro

// 5. Longest Palindromic Substring
// Medium
// Given a string s, return the longest palindromic substring in s.

// Example 1:

// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.
// Example 2:

// Input: s = "cbbd"
// Output: "bb"

// Constraints:

// 1 <= s.length <= 1000
// s consist of only digits and English letters.
func longestPalindrome(s string) string {
	dp := make(map[int]map[int]bool)
	n := len(s)
	for i := 0; i < n; i++ {
		dp[i] = make(map[int]bool)
		dp[i][i] = true
	}
	max, start, end := 1, 0, 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				dp[i][j] = true
			} else if s[i] != s[j] {
				dp[i][j] = false
			} else if i+1 >= n || j-1 < i+1 {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i+1][j-1]
			}
			if dp[i][j] && j-i+1 > max {
				max = j - i + 1
				start, end = i, j
			}
		}
	}
	return s[start : end+1]
}
