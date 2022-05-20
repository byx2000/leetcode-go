package my_calendar_iii

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/my-calendar-iii/

type MyCalendarThree struct {
	diff map[int]int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{map[int]int{}}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	this.diff[start]++
	this.diff[end]--

	// 对diff中的所有key进行排序
	var keys []int
	for k := range this.diff {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	result := 0
	sum := 0
	for _, k := range keys {
		sum += this.diff[k]
		result = int(math.Max(float64(result), float64(sum)))
	}

	return result
}
