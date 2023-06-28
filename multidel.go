package multidel

import (
	"sort"
)

// Delete removes all elements at specified indices
func Delete[T any](slice *[]T, indices ...int) {
	sort.Slice(indices, func(i, j int) bool { return indices[i] < indices[j] })

	for i, idx := range indices {
		indices[i] = idx - i
	}

	for _, idx := range indices {
		*slice = append((*slice)[:idx], (*slice)[idx+1:]...)
	}
}

// DeleteFunc removes all elements for which the shouldDelete function returned true.
func DeleteFunc[T any](slice *[]T, shouldDelete func(T) bool) {
	for i := 0; i < len(*slice); {
		if !shouldDelete((*slice)[i]) {
			i += 1
			continue
		}

		*slice = append((*slice)[:i], (*slice)[i+1:]...)
		i = 0
	}
}
