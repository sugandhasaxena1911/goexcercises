package main

import (
	"fmt"
)

// assumption: binary tree

type Node struct {
	val   int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(val int) {
	node := Node{val, nil, nil}
	q := []*Node{}
	if t.root == nil {
		t.root = &node
		return
	}
	current := t.root
	for current != nil {
		if current.left == nil {
			current.left = &node
			return
		} else {
			q = append(q, current.left)
		}
		if current.right == nil {
			current.right = &node
			return
		} else {
			q = append(q, current.right)
		}

		current = q[0]
		q = q[1:]
	}

}

// Left root Right
func (t *Tree) InorderTraverse() {
	inorder(t.root)
}

func inorder(current *Node) {
	if current == nil {
		return
	}
	inorder(current.left)
	fmt.Println(current.val)
	inorder(current.right)
}

// root left right
func (t *Tree) PreorderTraverse() {
	preorder(t.root)
}
func preorder(current *Node) {
	if current == nil {
		return
	}
	fmt.Println(current.val)
	preorder(current.left)
	preorder(current.right)
}

// left  right root
func (t *Tree) PostTraverse() {
	postorder(t.root)
}
func postorder(current *Node) {
	if current == nil {
		return
	}
	postorder(current.left)
	postorder(current.right)
	fmt.Println(current.val)

}

func (t *Tree) Height() int {
	return height(t.root)

}
func height(current *Node) int {
	if current == nil {
		return 0
	}
	l := height(current.left)
	r := height(current.right)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}

}

func main() {
	t := Tree{}
	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(4)
	t.Insert(5)
	t.Insert(6)
	t.Insert(7)
	t.InorderTraverse()
	t.PreorderTraverse()
	fmt.Println(t.Height())
}
