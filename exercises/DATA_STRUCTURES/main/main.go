package main

import (
	"fmt"
	"github.com/sugandhasaxena1911/goexcercises/exercises/DATA_STRUCTURES"
)

func main() {

	/*
		var ll1 DATA_STRUCTURES.LinkedList

		fmt.Println("INSERT 10 ELEMENTS")
		for x := 1; x <= 10; x++ {
			node := DATA_STRUCTURES.ListNode{x, nil}
			ll1.InsertNodeAtEnd(&node)

		}
		fmt.Println("TRAVERSE")
		ll1.TraverseNode()

		fmt.Println("DELETE FIRST NODE ")
		node := DATA_STRUCTURES.ListNode{1, nil}
		ll1.DeleteNodeByValue(&node)
		fmt.Println("TRAVERSE")
		ll1.TraverseNode()

		fmt.Println("DELETE LAST NODE ")
		node = DATA_STRUCTURES.ListNode{10, nil}
		ll1.DeleteNodeByValue(&node)
		fmt.Println("TRAVERSE")
		ll1.TraverseNode()

		fmt.Println("DELETE middle NODE ")
		node = DATA_STRUCTURES.ListNode{6, nil}
		ll1.DeleteNodeByValue(&node)
		fmt.Println("TRAVERSE")
		ll1.TraverseNode()

		fmt.Println("DeleteNthFromEnd , n==2")
		ll1.DeleteNthFromEnd(2)
		ll1.TraverseNode()
		fmt.Println("DeleteNthFromEnd , n==1")
		ll1.DeleteNthFromEnd(1)
		ll1.TraverseNode()

		fmt.Println("DeleteNthFromEnd , n==4")
		ll1.DeleteNthFromEnd(4)
		ll1.TraverseNode()

		var ll2 DATA_STRUCTURES.LinkedList

		fmt.Println("INSERT 10 ELEMENTS")
		for x := 1; x <= 10; x++ {
			node := DATA_STRUCTURES.ListNode{x, nil}
			ll2.InsertNodeAtEnd(&node)

		}
		fmt.Println("TRAVERSE")
		ll2.TraverseNode()
		head := ll2.ReverseLinkedList()
		fmt.Println("reversed linked list")
		ll3 := DATA_STRUCTURES.LinkedList{head}
		ll3.TraverseNode()

		fmt.Println("Original linked list ")
		ll2.TraverseNode()
		fmt.Println("reverse linked list in place ")
		ll2.ReverseLinkedListV2()
		ll2.TraverseNode()

		fmt.Println("MERGE LINKED LISTS")
		fmt.Println("INSERT 5 ELEMENTS to list1")
		var llA DATA_STRUCTURES.LinkedList

		for x := 1; x <= 5; x++ {
			node := DATA_STRUCTURES.ListNode{x, nil}
			llA.InsertNodeAtEnd(&node)

		}
		fmt.Println("TRAVERSE 1")
		llA.TraverseNode()

		fmt.Println("INSERT 5 ELEMENTS to list2")
		var llB DATA_STRUCTURES.LinkedList

		for x := 4; x <= 8; x++ {
			node := DATA_STRUCTURES.ListNode{x, nil}
			llB.InsertNodeAtEnd(&node)

		}
		fmt.Println("TRAVERSE 2")
		llB.TraverseNode()
		fmt.Println("calling MERGE")
		ll3head := DATA_STRUCTURES.MergeLinkedLists(llA.Head, llB.Head)
		ll4 := DATA_STRUCTURES.LinkedList{ll3head}
		fmt.Println("calling MERGED LIST")

		ll4.TraverseNode()

		fmt.Println("PALLINDROME")
		node1 := DATA_STRUCTURES.ListNode{1, nil}
		node2 := DATA_STRUCTURES.ListNode{2, nil}
		node3 := DATA_STRUCTURES.ListNode{2, nil}
		node4 := DATA_STRUCTURES.ListNode{1, nil}
		ll5 := DATA_STRUCTURES.LinkedList{nil}
		ll5.InsertNodeAtEnd(&node1)
		ll5.InsertNodeAtEnd(&node2)
		ll5.InsertNodeAtEnd(&node3)
		ll5.InsertNodeAtEnd(&node4)
		//ll5.InsertNodeAtEnd(&node5)

		ll5.TraverseNode()
		fmt.Println("Is pallindorme ??", ll5.IsPallindromeV3())

		fmt.Println("HAS CYCLE")
		node1 = DATA_STRUCTURES.ListNode{1, nil}
		node2 = DATA_STRUCTURES.ListNode{2, nil}
		node3 = DATA_STRUCTURES.ListNode{3, nil}
		node4 = DATA_STRUCTURES.ListNode{4, nil}
		node5 := DATA_STRUCTURES.ListNode{5, nil}

		ll6 := DATA_STRUCTURES.LinkedList{nil}
		ll6.InsertNodeAtEnd(&node1)
		ll6.InsertNodeAtEnd(&node2)
		ll6.InsertNodeAtEnd(&node3)
		ll6.InsertNodeAtEnd(&node4)
		ll6.InsertNodeAtEnd(&node5)

		//ll6.TraverseNode()   : keep on printing
		fmt.Println("has cycle ??  ", ll6.HasCycle())


	*/

	fmt.Println("TREE")
	tree := DATA_STRUCTURES.Tree{}
	tn1 := DATA_STRUCTURES.TreeNode{11, nil, nil}
	tn2 := DATA_STRUCTURES.TreeNode{12, nil, nil}
	tn3 := DATA_STRUCTURES.TreeNode{13, nil, nil}
	tn4 := DATA_STRUCTURES.TreeNode{14, nil, nil}
	tn5 := DATA_STRUCTURES.TreeNode{15, nil, nil}
	tn6 := DATA_STRUCTURES.TreeNode{16, nil, nil}
	tn7 := DATA_STRUCTURES.TreeNode{17, nil, nil}

	tree.InsertTreeNodes(&tn1)
	tree.InsertTreeNodes(&tn2)
	tree.InsertTreeNodes(&tn3)
	tree.InsertTreeNodes(&tn4)
	tree.InsertTreeNodes(&tn5)
	tree.InsertTreeNodes(&tn6)
	tree.InsertTreeNodes(&tn7)

	fmt.Println("INORDER")
	tree.InorderTraverseTree()
	fmt.Println("POSTORDER")
	tree.PostorderTraverseTree()
	fmt.Println("PREORDER")
	tree.PreorderTraverseTree()

	//start deleting Node
	fmt.Println("DELETE NODE 12")
	key := 12
	tree.DeleteNode(key)
	fmt.Println("INORDER")
	tree.InorderTraverseTree()

	fmt.Println("DELETE NODE 14")
	key = 14
	tree.DeleteNode(key)
	fmt.Println("INORDER")
	tree.InorderTraverseTree()

	fmt.Println("DELETE NODE 16")
	key = 16
	tree.DeleteNode(key)
	fmt.Println("INORDER")
	tree.InorderTraverseTree()

	fmt.Println("TREE-2")
	tree = DATA_STRUCTURES.Tree{}
	tn1 = DATA_STRUCTURES.TreeNode{11, nil, nil}
	tn2 = DATA_STRUCTURES.TreeNode{12, nil, nil}
	tn3 = DATA_STRUCTURES.TreeNode{13, nil, nil}
	tn4 = DATA_STRUCTURES.TreeNode{14, nil, nil}
	tn5 = DATA_STRUCTURES.TreeNode{15, nil, nil}
	tn6 = DATA_STRUCTURES.TreeNode{16, nil, nil}
	tn7 = DATA_STRUCTURES.TreeNode{17, nil, nil}

	tree.InsertTreeNodes(&tn1)
	tree.InsertTreeNodes(&tn2)
	tree.InsertTreeNodes(&tn3)
	tree.InsertTreeNodes(&tn4)
	tree.InsertTreeNodes(&tn5)
	tree.InsertTreeNodes(&tn6)
	tree.InsertTreeNodes(&tn7)
	fmt.Println("Traverse")
	tree.InorderTraverseTree()

	fmt.Println("HEIGHT", tree.HeightOfTree())
	fmt.Println("HEIGHT2", tree.HeightOfTreeV2())

	tn8 := DATA_STRUCTURES.TreeNode{18, nil, nil}
	tree.InsertTreeNodes(&tn8)
	fmt.Println("HEIGHT", tree.HeightOfTree())
	fmt.Println("HEIGHT2", tree.HeightOfTreeV2())
	fmt.Println("Total nodes : ", tree.TotalNumberOfNodes())

	fmt.Println("Insert into binary tree")
	tree = DATA_STRUCTURES.Tree{}
	tn1 = DATA_STRUCTURES.TreeNode{11, nil, nil}
	tn2 = DATA_STRUCTURES.TreeNode{13, nil, nil}
	tn3 = DATA_STRUCTURES.TreeNode{12, nil, nil}
	tn4 = DATA_STRUCTURES.TreeNode{14, nil, nil}
	tn5 = DATA_STRUCTURES.TreeNode{10, nil, nil}
	tn6 = DATA_STRUCTURES.TreeNode{2, nil, nil}
	tn7 = DATA_STRUCTURES.TreeNode{7, nil, nil}
	tree.InsertIntoBinaryTree(&tn1)
	tree.InsertIntoBinaryTree(&tn2)
	tree.InsertIntoBinaryTree(&tn3)
	tree.InsertIntoBinaryTree(&tn4)
	tree.InsertIntoBinaryTree(&tn5)
	tree.InsertIntoBinaryTree(&tn6)
	tree.InsertIntoBinaryTree(&tn7)
	fmt.Println("Traverse Tree")
	tree.InorderTraverseTree()

	fmt.Println("Another tree ")
	tree2 := DATA_STRUCTURES.Tree{}
	tn1 = DATA_STRUCTURES.TreeNode{11, nil, nil}
	tn2 = DATA_STRUCTURES.TreeNode{13, nil, nil}
	tn3 = DATA_STRUCTURES.TreeNode{12, nil, nil}
	tn4 = DATA_STRUCTURES.TreeNode{14, nil, nil}
	tn5 = DATA_STRUCTURES.TreeNode{10, nil, nil}
	tn6 = DATA_STRUCTURES.TreeNode{2, nil, nil}
	tn7 = DATA_STRUCTURES.TreeNode{7, nil, nil}
	tree2.InsertIntoBinaryTreeV2(&tn1)
	tree2.InsertIntoBinaryTreeV2(&tn2)
	tree2.InsertIntoBinaryTreeV2(&tn3)
	tree2.InsertIntoBinaryTreeV2(&tn4)
	tree2.InsertIntoBinaryTreeV2(&tn5)
	tree2.InsertIntoBinaryTreeV2(&tn6)
	tree2.InsertIntoBinaryTreeV2(&tn7)
	fmt.Println("Traverse Tree")
	tree2.InorderTraverseTree()
}
