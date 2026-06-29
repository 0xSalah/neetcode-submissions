func min(a int, b int) int {
	if a > b {
		return b
	}else {
		return a
	}
}
func container (a int, b int, index_a int, index_b int) int{
	width := index_b - index_a
	return width * min(a,b)
}
func maxArea(heights []int) int {
	f := 0
	l := len(heights) -1
	max := 0
	for f < l {
		result := container(heights[f],heights[l],f,l)
		if result > max {
			max = result
		}
		if heights[f] < heights[l] {
			f++
		}else {
			l--
		}
	}
	return max
}
