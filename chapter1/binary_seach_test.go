package chapter1

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type testModel struct {
		data   []int
		target int
		want   int
	}

	testCase := func(t *testing.T, testParams testModel) {
		got := BinarySearch(testParams.data, testParams.target)
		if got != testParams.want {
			t.Errorf("BinarySearch(%v, %v) = %v; want %v", testParams.data, testParams.target, got, testParams.want)
		}
	}

	t.Run("happy path", func(t *testing.T) {
		testParam := testModel{
			[]int{1, 3, 5, 7, 9}, 3, 1,
		}
		testCase(t, testParam)
	})

	t.Run("target not found", func(t *testing.T) {
		testParam := testModel{
			[]int{4, 6, 10, 14, 20}, -4, -1,
		}
		testCase(t, testParam)
	})

	t.Run("empty list", func(t *testing.T) {
		testParam := testModel{
			[]int{}, 3, -1,
		}
		testCase(t, testParam)
	})

	t.Run("single element - found", func(t *testing.T) {
		testParam := testModel{
			[]int{5}, 5, 0,
		}
		testCase(t, testParam)
	})

	t.Run("single element - not found", func(t *testing.T) {
		testParam := testModel{
			[]int{5}, 3, -1,
		}
		testCase(t, testParam)
	})

	t.Run("target is first element", func(t *testing.T) {
		testParam := testModel{
			[]int{1, 3, 5, 7, 9}, 1, 0,
		}
		testCase(t, testParam)
	})

	t.Run("target is last element", func(t *testing.T) {
		testParam := testModel{
			[]int{1, 3, 5, 7, 9}, 9, 4,
		}
		testCase(t, testParam)
	})

}
