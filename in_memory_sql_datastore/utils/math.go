package utils

type T any

func PullFirstValueFromArray[T comparable](sourceArr []T, target T) []T {
	result := []T{}
	for idx, v := range sourceArr {
		if v != target {
			continue
		}
		result = append(result[:idx], result[idx+1:]...)
	}
	return result
}
