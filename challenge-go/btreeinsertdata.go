package challenge_go

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if root.Data == data {
		return nil
	}
	if root.Data > data {
		if root.Left == nil {
			root.Left = &TreeNode{Data: data}
			root.Left.Parent = root
		}
		return BTreeInsertData(root.Left, data)
	}
	if root.Data < data {
		if root.Right == nil {
			root.Right = &TreeNode{Data: data}
			root.Right.Parent = root
		}
		return BTreeInsertData(root.Right, data)
	}
	return root
}
