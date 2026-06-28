func hasDuplicate(nums []int) bool {
    dup :=  make(map[int]int)
    for i := 0; i < len(nums); i++ {
        _, isHere := dup[nums[i]]
        if isHere {
            return true
        }else {
        dup[nums[i]] = i
        }
    }
    return false
}
