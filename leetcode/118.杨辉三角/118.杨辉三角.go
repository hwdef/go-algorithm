

// @lc code=start
func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	triangle := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		rows := []int{1}
		for j := 1; j <= i-1; j++ {
			num := triangle[i-1][j-1] + triangle[i-1][j]
			rows = append(rows, num)
		}
		rows = append(rows, 1)
		triangle = append(triangle, rows)
	}
	return triangle
}

// @lc code=end

