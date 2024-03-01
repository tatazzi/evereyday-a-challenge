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

// Runtime: 85ms
// Memory: 15.44mb

//func containsDuplicate(nums []int) bool {
//	sort.Slice(nums, func(i, j int) bool {
//		return nums[i] > nums[j]
//	})
//	n := len(nums)
//	for i := 0; i < n-1; i++ {
//		if nums[i] == nums[i+1] {
//			return true
//		}
//	}
//	return false
//}
//Runtime: 90ms
//memory: 8.12mb
