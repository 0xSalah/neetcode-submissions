func majorityElement(nums []int) int {
   count :=0
   var element int
   r := make(map[int]int) 

   for _, v := range nums {
		r[v] += 1
		if count < r[v] {
			element = v
			count = r[v]
		}
   }
	return element
}
