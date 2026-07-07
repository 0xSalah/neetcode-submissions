func twoSum(numbers []int, target int) []int {
	index1 := 0
	index2 := len(numbers)- 1

	for i:=  0; i < len(numbers)- 1; i++ {
		if numbers[index1] + numbers[index2] == target {
			return []int{index1+1,index2+1}
		}else if numbers[index1] + numbers[index2] > target {
			index2--
		}else {
			index1++
		}
	}
	return []int{}
}
