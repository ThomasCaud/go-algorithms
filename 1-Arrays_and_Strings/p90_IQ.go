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

