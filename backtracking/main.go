package backtracking

func backtrack(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}

	if len(nums) == 2 {
		// tmp := make([]int, len(nums))
		// copy(tmp, nums)
		return [][]int{nums, {nums[1], nums[0]}}
	}

	ans := [][]int{}

	for i := 0; i < len(nums); i++ {

		cp := make([]int, len(nums))
		copy(cp, nums)
		cp[i], cp[0] = cp[0], cp[i]


		res := backtrack(cp[1:])

		for _, val := range res {
			val = append([]int{nums[i]}, val...)
			ans = append(ans, val)
		}

	}

	return ans
}
