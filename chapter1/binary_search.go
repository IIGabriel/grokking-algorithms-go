package chapter1

// BinarySearch performs a binary search on a sorted list of integers to find a target value.
// - Only works when the list is sorted in ascending order.
// - Returns the index of the target value if found, otherwise it returns -1.
// - The time complexity of this function is O(log n).
func BinarySearch(data []int, target int) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := data[mid]
		if guess == target {
			return mid
		}
		if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
