// if c < 'a' { c += 32 } 
func isVowel(c rune) bool {
    c = unicode.ToLower(c)
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func reverseVowels(s string) string {
    ss := []rune(s)
    n := len(ss)
    l, r := 0, n - 1
    for l < r {
        // find the index of the first vowel in the given range
        for l < r && !isVowel(ss[l]) {
            l += 1
        }
        // find the index of last vowel in the given range
        for r > l && !isVowel(ss[r]) {
            r -= 1
        }
        // swap ss[l] and ss[r]
        ss[l], ss[r] = ss[r], ss[l]
        // since we've processed the character ss[l],
        // we move the left pointer to the right
        l += 1
        // since we've processed the character ss[r],
        // we move the right pointer to the left
        r -= 1
    }
    return string(ss)
}