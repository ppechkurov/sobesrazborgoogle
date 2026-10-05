package sums

func getIndices(nums []int, target int) (first, second int, success bool) {
	seen := make(map[int]int, len(nums))
	for i, v := range nums {
		j, ok := seen[target-v]
		if ok {
			return j, i, true
		}

		seen[v] = i
	}

	return -1, -1, false
}
