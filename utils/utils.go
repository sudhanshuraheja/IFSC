package utils

func MergeMaps(map1 map[string]int, map2 map[string]int) map[string]int {
	result := make(map[string]int)

	for key, value := range map1 {
		result[key] = value
	}

	for key := range map2 {
		_, keyExists := map2[key]
		if keyExists {
			result[key] = result[key] + map2[key]
		} else {
			result[key] = map2[key]
		}
	}

	return result
}
