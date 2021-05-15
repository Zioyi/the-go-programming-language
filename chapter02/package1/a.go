package package1

import "fmt"

var A int = B

// var B int = C   // 变异报错 因为B已经在b.go中声明了

func F() {
	fmt.Println(B)
}

func increasingBST(root *TreeNode) *TreeNode {
	dummuyNode := &TreeNode{}
	ptrNode := dummuyNode
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)

		ptrNode.Right = node
		node.Left = nil
		ptrNode = node

		inorder(node.Right)
	}
	inorder(root)
	return dummuyNode.Right
}

func increasingBST2(root *TreeNode) *TreeNode {
	dummyNode := &TreeNode{}
	resNode := dummyNode
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)

		// 在中序遍历的过程中修改节点指向
		resNode.Right = node
		node.Left = nil
		resNode = node

		inorder(node.Right)
	}
	inorder(root)
	return dummyNode.Right
}
