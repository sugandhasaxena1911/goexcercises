package DATA_STRUCTURES

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Tree struct {
	Root *TreeNode
}

func (tree *Tree) InsertTreeNodes(tn *TreeNode) {
	if tree.Root == nil {
		tree.Root = tn
		return
	}
	//First In First Out
	q := []*TreeNode{}
	current := tree.Root
	for current != nil {

		if current.Left == nil {
			//insert here
			current.Left = tn
			break
		} else {
			q = append(q, current.Left)
		}
		if current.Right == nil {
			current.Right = tn
			break
		} else {
			q = append(q, current.Right)
		}

		current = q[0]
		q = q[1:]

	}
}

func (tree *Tree) InorderTraverseTree() {
	inorderTraverseTree(tree.Root)
	return
}
func inorderTraverseTree(tn *TreeNode) {
	if tn == nil {
		return
	}
	current := tn
	inorderTraverseTree(current.Left)
	fmt.Println(current.Val)
	inorderTraverseTree(current.Right)

}
func (tree *Tree) PostorderTraverseTree() {
	postorderTraverseTree(tree.Root)
	return
}
func postorderTraverseTree(tn *TreeNode) {
	if tn == nil {
		return
	}
	current := tn
	postorderTraverseTree(current.Left)
	postorderTraverseTree(current.Right)
	fmt.Println(current.Val)

}
func (tree *Tree) PreorderTraverseTree() {
	preorderTraverseTree(tree.Root)
	return
}
func preorderTraverseTree(tn *TreeNode) {
	if tn == nil {
		return
	}
	current := tn
	fmt.Println(current.Val)
	preorderTraverseTree(current.Left)
	preorderTraverseTree(current.Right)

}

func (tree *Tree) DeleteNode(key int) {

	if tree.Root == nil {
		fmt.Println("Empty tree ")
		return
	}
	current := tree.Root
	q := []*TreeNode{}
	for current != nil {
		fmt.Println("Current node value ", current.Val)

		if current.Val == key {
			temp := current
			//lets find the rightmost value
			var prev *TreeNode
			current = tree.Root
			for current.Right != nil {
				prev = current
				current = current.Right

			}
			temp.Val = current.Val // replaced value
			//delete the righmostnode
			if current.Left == nil { // its leaf node
				prev.Right = nil
			} else {
				prev.Right = current.Left
			}

			return
		}

		if current.Left != nil {
			q = append(q, current.Left)
		}
		if current.Right != nil {
			q = append(q, current.Right)
		}

		current = q[0]
		q = q[1:]

	}

}

func (tree *Tree) TotalNumberOfNodes() int {
	current := tree.Root

	if current == nil {
		return 0
	}
	q := []*TreeNode{}
	q = append(q, current)
	n := 0

	for len(q) > 0 {
		current = q[0]
		q = q[1:]
		if current.Left != nil {
			q = append(q, current.Left)

		}
		if current.Right != nil {
			q = append(q, current.Right)

		}
		n++
	}
	return n
}

func (tree *Tree) HeightOfTree() int {
	return heightOfTree(tree.Root)

}

func heightOfTree(node *TreeNode) int {
	if node == nil {
		return 0
	}
	l := heightOfTree(node.Left)
	r := heightOfTree(node.Left)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}

}
func (tree *Tree) HeightOfTreeV2() int {
	return heightOfTreeV2(tree.Root)

}
func heightOfTreeV2(node *TreeNode) int {
	if node == nil {
		return 0
	}
	q := []*TreeNode{}
	q = append(q, node)
	height := 0
	for len(q) > 0 {
		n := len(q) // number of nodes at current level
		fmt.Println("Number of nodes at level ", height, "is: ", n)
		for n > 0 {
			current := q[0]
			q = q[1:]

			if current.Left != nil {
				q = append(q, current.Left)

			}
			if current.Right != nil {
				q = append(q, current.Right)

			}
			n--
		}

		height++
	}

	return height

}
