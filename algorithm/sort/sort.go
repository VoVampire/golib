package sort

import "sort"

// https://visualgo.net/en/sorting

// 优化思路：
// 1. 提前结束
// 2. 边界哨兵
// 3. 小数组(5~15)切换到插入排序

type Sorter interface {
	Name() string
	Sort([]int)
}

type GoSort struct{}

func (g *GoSort) Name() string { return "go-sort" }
func (g *GoSort) Sort(a []int) { sort.Ints(a) }

type sorter struct{}

func (s *sorter) less(i, j int) bool { return i < j }
func (s *sorter) more(i, j int) bool { return i > j }
func (s *sorter) swap(i, j *int)     { *i, *j = *j, *i }
