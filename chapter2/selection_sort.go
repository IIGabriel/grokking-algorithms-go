package chapter2

// SelectionSort sorts the given slice of integers in ascending order using the selection sort algorithm.
// - The time complexity of this function is O(n^2).
// - The space complexity of this function is O(1).
func SelectionSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		lowerIndex := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[lowerIndex] {
				lowerIndex = j
			}
		}
		data[i], data[lowerIndex] = data[lowerIndex], data[i]
	}
}
