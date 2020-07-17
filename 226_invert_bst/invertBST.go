package invertbst

// Time O(N)
// Space O(N)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	invert(root)

	return root
}

func invert(t *TreeNode) {
	if t == nil {
		return
	}

	invert(t.Left)
	invert(t.Right)

	t.Left, t.Right = t.Right, t.Left
}
