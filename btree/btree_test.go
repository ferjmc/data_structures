package btree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree := &TreeNode{Val: 1}
	tree.insert(2)
	tree.insert(3)
	tree.insert(4)
	tree.insert(5)
	fmt.Println(tree)
	fmt.Println(diameterOfBinaryTree(tree))
}
