package linkedlist

// ListNode defines a node in a singly linked list.
// LeetCode 经常使用这个定义。
type ListNode struct {
    Val  int
    Next *ListNode
}

// 你可以在这里添加链表相关的辅助函数，例如：
// - 从数组创建链表
// - 打印链表
// - 反转链表等

// NewListNode creates a new list node
func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

// CreateLinkedListFromArray creates a linked list from an integer array
func CreateLinkedListFromArray(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

// ToArray converts a linked list to an array of integers
func (head *ListNode) ToArray() []int {
	var res []int
	current := head
	for current != nil {
		res = append(res, current.Val)
		current = current.Next
	}
	return res
}