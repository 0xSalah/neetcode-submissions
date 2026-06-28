func isAnagram(s string, t string) bool {
	arg := make(map[byte]int)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		arg[s[i]] += 1
	}
	for  i := 0; i < len(s); i++{
		arg[t[i]]--
		if arg[t[i]] == -1 {
			return false
		}
	}
	return true
}
