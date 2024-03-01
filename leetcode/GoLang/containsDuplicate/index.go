package containsDuplicate

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	hasTwoOrMore := false
	for i := 0; i < len(nums); i++ {
		current := m[nums[i]]
		if current != 0 {
			hasTwoOrMore = true
			break
		} else {
			m[nums[i]] = 1
		}
	}
	return hasTwoOrMore
}
