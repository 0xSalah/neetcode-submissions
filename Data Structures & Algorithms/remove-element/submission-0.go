func removeElement(nums []int, val int) int {
	last := len(nums) - 1
	k := 0
	for k <= last {
		if nums[k] == val{
			nums[k] = nums[last]
			last--
		}else {
			k++
		}
	}
	return k
}
