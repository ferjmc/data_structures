package btree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) insert(v int) {
	if n == nil {
		return
	}
	if n.Left == nil {
		n.Left = &TreeNode{Val: v}
		return
	}
	if n.Right == nil {
		n.Right = &TreeNode{Val: v}
		return
	}
	n.Left.insert(v)
	return

	//	else if v <= n.Val {
	//		if n.Left == nil {
	//			n.Left = &TreeNode{Val: v, Left: nil, Right: nil}
	//		} else {
	//			n.Left.insert(v)
	//		}
	//	} else {
	//		if n.Right == nil {
	//			n.Right = &TreeNode{Val: v, Left: nil, Right: nil}
	//		} else {
	//			n.Right.insert(v)
	//		}
	//	}
}

func (n *TreeNode) String() string {
	var res string
	if n == nil {
		return res
	}
	res += fmt.Sprintf("%d,", n.Val)
	res += fmt.Sprint(n.Left)
	res += fmt.Sprint(n.Right)
	return res
}

func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	getDiameter(root, &res)
	return res
}

func getDiameter(tree *TreeNode, diameter *int) int {
	if tree == nil {
		return 0
	}
	var leftDepth, rightDepth int
	if tree.Left != nil {
		leftDepth = 1 + getDiameter(tree.Left, diameter)
	}
	if tree.Right != nil {
		rightDepth = 1 + getDiameter(tree.Right, diameter)
	}
	currDiameter := leftDepth + rightDepth
	if *diameter < currDiameter {
		*diameter = currDiameter
	}
	if leftDepth > rightDepth {
		return leftDepth
	}
	return rightDepth
}
