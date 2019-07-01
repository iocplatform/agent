package api

import "strings"

// RemoveDuplicates entries from a string array
func RemoveDuplicates(a []string) []string {
	result := []string{}
	seen := map[string]string{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}

// NormalizeHash hash list
func NormalizeHash(a []string) []string {
	result := []string{}
	for _, val := range a {
		if strings.HasPrefix(val, "0x") {
			val = val[2:]
		}
		result = append(result, strings.ToLower(val))
	}
	return result
}
