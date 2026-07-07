func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)- 1
	pivotIn := rand.Int() % len(a)

	a[pivotIn], a[right] = a[right], a[pivotIn]

	for i := 0; i < len(a)- 1; i++ {
		if a[i] < a[right] { 
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	a[right], a[left] = a[left], a[right]
	quicksort(a[:left])
	quicksort(a[left+1:])
	return a
}



func sortArray(nums []int) []int {
    return quicksort(nums)
}
