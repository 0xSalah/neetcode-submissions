func sortColors(nums []int) {
    left := 0
	mid := 0
	right := len(nums)- 1

	for mid <= right {
		if nums[mid] == 0 {
			nums[mid],nums[left] = nums[left], nums[mid]
			left++
			mid++
		}else if nums[mid] == 1 {
			mid++
		}else{
			nums[mid],nums[right] = nums[right], nums[mid]
			right--
		}
	}
}
