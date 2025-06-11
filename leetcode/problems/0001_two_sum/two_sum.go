package p0001_two_sum // 包名更改，方便导入

// TwoSum 接受一个整数数组 nums 和一个目标值 target，
// 返回数组中和为 target 的两个数的索引。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 函数名首字母大写以导出
func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // 值 -> 索引
	for i, num := range nums {
		complement := target - num
		if index, ok := numMap[complement]; ok {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return nil // 题目保证有解，这里理论上不会执行到
}