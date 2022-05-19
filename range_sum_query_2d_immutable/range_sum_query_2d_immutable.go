package range_sum_query_2d_immutable

// https://leetcode.cn/problems/range-sum-query-2d-immutable/

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	// sum[i]表示第i行的前缀和数组
	sum := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		sum[i] = make([]int, len(matrix[i])+1)
		for j := 0; j < len(matrix[i]); j++ {
			sum[i][j+1] = sum[i][j] + matrix[i][j]
		}
	}

	return NumMatrix{sum}
}

func (this *NumMatrix) SumRegion(r1 int, c1 int, r2 int, c2 int) int {
	s := 0
	for i := r1; i <= r2; i++ {
		s += this.sum[i][c2+1] - this.sum[i][c1]
	}
	return s
}
