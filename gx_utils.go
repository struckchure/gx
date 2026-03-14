package gx

import (
	"fmt"
	"strings"
)

func formatToOpenApi(path string) string {
	if path == "/" {
		return path
	}

	split := strings.Split(path, "/")
	result := make([]string, len(split))

	for i, s := range split {
		if after, ok := strings.CutPrefix(s, ":"); ok {
			result[i] = fmt.Sprintf("{%s}", after)
		} else {
			result[i] = s
		}
	}

	return strings.Join(result, "/")
}

func last[T any](slice []T, n int) []T {
	if n <= 0 {
		return nil
	}
	if len(slice) < n {
		return slice
	}
	return slice[len(slice)-n:]
}
