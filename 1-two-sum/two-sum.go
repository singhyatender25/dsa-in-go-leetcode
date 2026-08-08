func twoSum(nums []int, target int) []int {

	mp := make(map[int]int)

	for i := range nums {
		complement := target - nums[i]

		idx, exists := mp[complement]
		if exists {
			return []int{idx, i}
		}
		mp[nums[i]] = i
	}
	return nil
}