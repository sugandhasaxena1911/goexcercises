package DATA_STRUCTURES

import "fmt"

func (tree *Tree) InsertIntoBinaryTree(node *TreeNode) {
	if tree.Root == nil {

		tree.Root = node
		return
	}
	fmt.Println("Inserting ", node.Val)
	current := tree.Root
	key := node.Val
	for current != nil {
		fmt.Println(current.Val)
		if key < current.Val {
			if current.Left == nil {
				current.Left = node
				return
			} else {
				current = current.Left
			}
			continue
		}

		if key > current.Val {
			if current.Right == nil {
				current.Right = node
				return
			} else {

				current = current.Right
			}
			continue
		}

	}

}

func (tree *Tree) InsertIntoBinaryTreeV2(node *TreeNode) {
	if tree.Root == nil {
		tree.Root = node
		return
	}
	node = insertBinarytree(tree.Root, node.Val)
}

func insertBinarytree(nd *TreeNode, key int) *TreeNode {
	if nd == nil {
		return &TreeNode{key, nil, nil}
	}

	if key > nd.Val {
		nd.Right = insertBinarytree(nd.Right, key)
	} else {
		if key < nd.Val {
			nd.Left = insertBinarytree(nd.Left, key)
		}
	}
	return nd

}
