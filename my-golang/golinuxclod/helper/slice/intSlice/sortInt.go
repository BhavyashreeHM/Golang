package intslice

import "sort"

func SortInt(a []int) []int{
	sort.Ints(a)
	return a
}