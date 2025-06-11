package tree

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode creates a new TreeNode.
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

// SliceToTree converts a slice of integers (level order, -1 for nil) to a binary tree.
func SliceToTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(nums) {
		node := queue[0]
		queue = queue[1:]

		if i < len(nums) && nums[i] != -1 { // -1 represents nil node
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(nums) && nums[i] != -1 { // -1 represents nil node
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// TreeToSliceLevelOrder converts a binary tree to a slice of integers (level order, -1 for nil).
func TreeToSliceLevelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			result = append(result, node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		} else {
			// We don't append -1 for trailing nil nodes in the conventional sense,
			// but to reconstruct, we might need them if there are non-nil nodes later at the same level or deeper.
			// For simplicity in this example, we'll only add -1 if it's not a trailing nil sequence.
			// A more robust serialization would handle this better.
			// However, to keep it simple and often sufficient for LeetCode, we can do this:
			// If we add -1, we need a way to know when to stop adding them if all subsequent nodes are nil.
			// This basic version might produce slices with trailing -1s that are not strictly necessary.
			result = append(result, -1) // Represent nil node with -1
		}
	}

	// Trim trailing -1s
	i := len(result) - 1
	for i >= 0 && result[i] == -1 {
		i--
	}
	return result[:i+1]
}