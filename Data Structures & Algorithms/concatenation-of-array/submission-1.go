func getConcatenation(nums []int) []int {
    n := len(nums)
    arr := make([]int, n*2)
    for i := 0; i < n; i++ {
        arr[i] = nums[i]
		arr[i+n] = nums[i]
    }
    return arr
}
