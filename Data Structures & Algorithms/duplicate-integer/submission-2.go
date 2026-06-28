func hasDuplicate(nums []int) bool {
    dup :=  make(map[int]bool)
    for i := 0; i < len(nums); i++ {
        if isHere := dup[nums[i]]; isHere {
            return true
        }else {
        dup[nums[i]] = true
        }
    }
    return false
}
