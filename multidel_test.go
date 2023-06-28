package multidel_test

import (
	"testing"

	"github.com/H3Cki/go-multidel"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		indices []int
		result  []int
		err     error
	}{
		{
			"no indices",
			[]int{1, 2, 3, 4, 5},
			[]int{},
			[]int{1, 2, 3, 4, 5},
			nil,
		},
		{
			"remove 1 index",
			[]int{1, 2, 3, 4, 5},
			[]int{0},
			[]int{2, 3, 4, 5},
			nil,
		},
		{
			"remove 2 indices",
			[]int{1, 2, 3, 4, 5},
			[]int{0, 2},
			[]int{2, 4, 5},
			nil,
		},
		{
			"remove all indices",
			[]int{1, 2, 3, 4, 5},
			[]int{0, 1, 2, 3, 4},
			[]int{},
			nil,
		},
		{
			"remove all indices unsorted",
			[]int{1, 2, 3, 4, 5},
			[]int{1, 4, 2, 3, 0},
			[]int{},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multidel.Delete(&tt.slice, tt.indices...)
			assert.EqualValues(t, tt.result, tt.slice)
		})
	}
}

func TestDeleteFunc(t *testing.T) {
	tests := []struct {
		name   string
		slice  []int
		should func(int) bool
		result []int
	}{
		{
			"no removals",
			[]int{1, 2, 3, 4, 5},
			func(int) bool { return false },
			[]int{1, 2, 3, 4, 5},
		},
		{
			"remove first and last",
			[]int{1, 2, 3, 4, 5},
			func(value int) bool { return value == 1 || value == 5 },
			[]int{2, 3, 4},
		},
		{
			"remove all",
			[]int{1, 2, 3, 4, 5},
			func(int) bool { return true },
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multidel.DeleteFunc(&tt.slice, tt.should)
			assert.EqualValues(t, tt.result, tt.slice)
		})
	}
}
