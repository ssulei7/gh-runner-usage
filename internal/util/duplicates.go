package util

// removeDuplicates returns a new slice with duplicates removed from the input slice.
func RemoveDuplicates(labels []string) []string {
	unique := make(map[string]bool)
	for _, label := range labels {
		unique[label] = true
	}

	result := make([]string, 0, len(unique))
	for label := range unique {
		result = append(result, label)
	}
	return result
}
