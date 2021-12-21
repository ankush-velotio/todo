package utils

func Contains(slice []string, element string) bool {
	for _, val := range slice {
		if element == val {
			return true
		}
	}
	return false
}
