package cmd



func longestPalindrome(s string) string {
	var longestStr string
	findLongestPandStr(s[:len(s)-1], &longestStr) 
	findLongestPandStr(s[1:], &longestStr)
	return longestStr
}

func findLongestPandStr(s string, start int, end int, longestStr *string) string {
	if IsPanlindrome(s) {
		if len (s) < len(*longestStr) {
			return *longestStr
		} else {
			longestStr = &s
		}
	}
    if len(s)==1 {
		return s
	}

	return *longestStr

}


func IsPanlindrome(s string) bool {
	leng := len(s)
	isPan := false
	for i, _  := range s {
		if i > (leng / 2) {
			return true
		}
		if s[i] != s[leng-1-i] {
			return false
		}
	}
	return isPan
}


// CHAT GPT
func longestPalindrome1(s string) string {
	if len(s) == 0 {
		return ""
	}
	var longestStr string

	// Iterate through all possible centers of the palindrome
	for i := 0; i < len(s); i++ {
		// Odd-length palindromes
		checkPalindrome(s, i, i, &longestStr)
		// Even-length palindromes
		checkPalindrome(s, i, i+1, &longestStr)
	}
	return longestStr
}

func checkPalindrome(s string, left, right int, longestStr *string) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		// Expand around the center
		left--
		right++
	}
	// Compare and update the longest palindrome found
	if len(s[left+1:right]) > len(*longestStr) {
		*longestStr = s[left+1 : right]
	}
}