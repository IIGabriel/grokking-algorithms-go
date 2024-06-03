package chapter2

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	type testModel struct {
		data     []int
		expected []int
	}

	testCase := func(t *testing.T, testParams testModel) {
		SelectionSort(testParams.data)
		if !reflect.DeepEqual(testParams.data, testParams.expected) {
			t.Errorf("SelectionSort(%v) = %v; want %v", testParams.data, testParams.data, testParams.expected)
		}
	}

	t.Run("already sorted list", func(t *testing.T) {
		testParam := testModel{
			[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5},
		}
		testCase(t, testParam)
	})

	t.Run("reverse order list", func(t *testing.T) {
		testParam := testModel{
			[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5},
		}
		testCase(t, testParam)
	})

	t.Run("list with duplicates", func(t *testing.T) {
		testParam := testModel{
			[]int{3, 1, 2, 1, 3}, []int{1, 1, 2, 3, 3},
		}
		testCase(t, testParam)
	})

	t.Run("empty list", func(t *testing.T) {
		testParam := testModel{
			[]int{}, []int{},
		}
		testCase(t, testParam)
	})

	t.Run("single element list", func(t *testing.T) {
		testParam := testModel{
			[]int{1}, []int{1},
		}
		testCase(t, testParam)
	})

	t.Run("two element list", func(t *testing.T) {
		testParam := testModel{
			[]int{2, 1}, []int{1, 2},
		}
		testCase(t, testParam)
	})

	t.Run("all elements same", func(t *testing.T) {
		testParam := testModel{
			[]int{5, 5, 5, 5, 5}, []int{5, 5, 5, 5, 5},
		}
		testCase(t, testParam)
	})
}
