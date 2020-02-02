// 给你一个整数数组 arr。你可以从中选出一个整数集合，并删除这些整数在数组中的每次出现。

// 返回 至少 能删除数组中的一半整数的整数集合的最小大小。

// 示例 1：

// 输入：arr = [3,3,3,3,5,5,5,2,2,7]
// 输出：2
// 解释：选择 {3,7} 使得结果数组为 [5,5,5,2,2]、长度为 5（原数组长度的一半）。
// 大小为 2 的可行集合有 {3,5},{3,2},{5,2}。
// 选择 {2,7} 是不可行的，它的结果数组为 [3,3,3,3,5,5,5]，新数组长度大于原数组的二分之一。
// 示例 2：

// 输入：arr = [7,7,7,7,7,7]
// 输出：1
// 解释：我们只能选择集合 {7}，结果数组为空。
// 示例 3：

// 输入：arr = [1,9]
// 输出：1
// 示例 4：

// 输入：arr = [1000,1000,3,7]
// 输出：1
// 示例 5：

// 输入：arr = [1,2,3,4,5,6,7,8,9,10]
// 输出：5

// 提示：

// 1 <= arr.length <= 10^5
// arr.length 为偶数
// 1 <= arr[i] <= 10^5

package main

import (
	"fmt"
	"sort"
)

type sortableArr []int

func (arr sortableArr) Len() int {
	return len(arr)
}

func (arr sortableArr) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr sortableArr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

type groups struct {
	num []int
}

func (g *groups) Len() int {
	return len(g.num)
}

func (g *groups) Swap(i, j int) {
	g.num[i], g.num[j] = g.num[j], g.num[i]
}

func (g *groups) Less(i, j int) bool {
	return g.num[i] > g.num[j]
}

func generateGroups(arr []int) groups {
	res := groups{[]int{}}
	current := arr[0]
	n := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == current {
			n++
		} else {
			res.num = append(res.num, n)
			current = arr[i]
			n = 1
		}
	}
	res.num = append(res.num, n)
	sort.Sort(&res)
	return res
}

func minSetSize(arr []int) int {
	rest := len(arr) / 2
	sort.Sort(sortableArr(arr))
	//fmt.Println(arr)
	nums := generateGroups(arr).num
	//fmt.Println(nums)
	min := 0
	for rest > 0 {
		rest -= nums[min]
		min++
	}
	return min
}

func main() {
	fmt.Println(minSetSize([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}))

	fmt.Println(minSetSize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))

	fmt.Println(minSetSize([]int{7, 7, 7, 7, 7, 7}))
}