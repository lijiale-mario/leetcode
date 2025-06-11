package p0001_two_sum // 与源文件包名一致

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "example 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "example 2",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "example 3",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := TwoSum(tc.nums, tc.target) // 调用导出的 TwoSum
			// 简单比较，实际中可能需要排序后比较或更复杂的逻辑
			if !reflect.DeepEqual(got, tc.want) && !reflect.DeepEqual(got, []int{tc.want[1], tc.want[0]}) {
				t.Errorf("TwoSum(%v, %d) = %v; want %v or %v", tc.nums, tc.target, got, tc.want, []int{tc.want[1], tc.want[0]})
			}
		})
	}
}