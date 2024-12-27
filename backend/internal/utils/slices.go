package utils

func AreSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	seen := make(map[string]bool)
	for _, v := range a {
		seen[v] = true
	}

	for _, v := range b {
		if !seen[v] {
			return false
		}
	}

	return true
}
