// @leet start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // stores number -> index

	for i , num := range nums {
		otherNumVal := target - num
		
		// check if other index is in the map
		if index , ok := m[otherNumVal]; ok {
			return []int{index,i}
		}
		// if not found, store the other index
		m[num] = i
	}
	return nil
}
// @leet end
