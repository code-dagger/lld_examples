package utils

func Contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

func PullSingleFromSlice[T comparable](slice []T, target T) []T {
	result := []T{}
	for idx, val := range slice {
		if val == target {
			result = append(slice[:idx], slice[idx+1:]...)
			break
		}
	}
	return result
}
