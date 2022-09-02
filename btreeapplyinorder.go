package challenge_go

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {

	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	_, _ = f(root.Data)
	BTreeApplyInorder(root.Right, f)

}
