// @leet start
func search(nums []int, target int) int {
	left , right := 0 , len(nums) -1

	for left <= right {
		mid := left + (right-left) / 2// avoid overflow

		switch  {
		case nums[mid] == target:
				return mid
		case nums[mid] < target:
				left = mid + 1
		default:
				right = mid -1
			
		}
	}
	return -1
}
// @leet end
