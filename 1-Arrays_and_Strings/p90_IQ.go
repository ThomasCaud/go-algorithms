package main

func isUnique(s string) bool {
	char := make(map[rune]bool);
	for _, c := range s {
		_, ok := char[c];
		if ok {
			return false
		}
		char[c] = true
	}
	return true;
}

func checkPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	runes := []rune(s1);
	for i := 0 ; i < len(runes) ; i++ {
		c2 := len(s2) - 1 - i;
		if s1[i] != s2[c2] {
			return false
		}
	}
	return true
}
