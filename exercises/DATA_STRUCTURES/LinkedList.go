package DATA_STRUCTURES

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (ll *LinkedList) InsertNodeAtEnd(node *ListNode) {
	if ll.Head == nil {
		ll.Head = node
		return
	} else {
		current := ll.Head
		//traverse uptil last element
		for current.Next != nil {
			current = current.Next
		}

		current.Next = node

	}

}
func (ll *LinkedList) DeleteNodeByValue(node *ListNode) {
	value := node.Val
	//only one element
	if ll.Head.Next == nil && ll.Head.Val == value {
		ll.Head = nil
		return
	} else {
		current := ll.Head
		var previous *ListNode
		for current != nil {
			if current.Val == value {

				if current == ll.Head {
					ll.Head = ll.Head.Next
					return
				}

				previous.Next = current.Next
			}
			previous = current
			current = current.Next

		}
	}

}

func (ll *LinkedList) TraverseNode() {
	current := ll.Head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next

	}

}

func (ll *LinkedList) DeleteNthFromEnd(n int) {
	current := ll.Head
	pointer := ll.Head
	counter := 0
	for current != nil {
		if counter > n {
			pointer = pointer.Next

		}
		counter++
		current = current.Next

	}
	//edge case , ie first elemnt
	if n == counter {
		ll.Head = ll.Head.Next
		return
	}
	// nothing to be deleted
	if n > counter {
		return
	}
	if n < counter {
		pointer.Next = pointer.Next.Next

	}

}

// ReverseLinkedList reverse and return the head
func (ll *LinkedList) ReverseLinkedList() *ListNode {
	current := ll.Head
	sum := 0
	sl := []int{}
	for current != nil {
		sl = append(sl, current.Val)
		sum++
		current = current.Next

	}
	high := len(sl) - 1
	var ll2 LinkedList
	for high >= 0 {
		node := ListNode{sl[high], nil}
		ll2.InsertNodeAtEnd(&node)
		high--
	}
	return ll2.Head
}

func (ll *LinkedList) ReverseLinkedListV2() {
	current := ll.Head
	var previous *ListNode
	for current != nil {
		forward := current.Next // to use in later
		current.Next = previous
		previous = current
		current = forward
	}
	ll.Head = previous
}

func MergeLinkedLists(head1 *ListNode, head2 *ListNode) *ListNode {
	// assumption : they are sorted
	current1 := head1
	current2 := head2
	ll := LinkedList{nil}
	check := 0
	for current1 != nil && current2 != nil {
		fmt.Println("INSIDE", current1.Val, current2.Val)
		if current1.Val < current2.Val {
			fmt.Println("INSIDE")
			node := ListNode{current1.Val, nil}
			ll.InsertNodeAtEnd(&node)
			check = 1
		}
		if current2.Val < current1.Val {
			node := ListNode{current2.Val, nil}
			ll.InsertNodeAtEnd(&node)
			check = 2
		}
		if current2.Val == current1.Val {
			fmt.Println("EQUALS")
			node1 := ListNode{current1.Val, nil}
			node2 := ListNode{current2.Val, nil}

			fmt.Println("PRINTING")
			ll.TraverseNode()
			ll.InsertNodeAtEnd(&node1)
			ll.InsertNodeAtEnd(&node2)
			current2 = current2.Next
			current1 = current1.Next
			check = 0

		}
		fmt.Println("check value", check)
		if check == 1 {
			current1 = current1.Next
		} else if check == 2 {
			current2 = current2.Next

		}

	}
	fmt.Println("LOOP ENDS")
	ll.TraverseNode()
	for current1 != nil {
		node := ListNode{current1.Val, nil}
		ll.InsertNodeAtEnd(&node)
		current1 = current1.Next
	}
	for current2 != nil {
		node := ListNode{current2.Val, nil}
		ll.InsertNodeAtEnd(&node)
		current2 = current2.Next
	}
	return ll.Head

}

func (ll *LinkedList) IsPallindrome() bool {
	current1 := ll.Head
	current2 := ll.ReverseLinkedList() // Note that thsi func passes another created Linked List
	for current1 != nil && current2 != nil {
		fmt.Println("values ", current1.Val, current2.Val)
		if current1.Val != current2.Val {
			return false
		}
		current1 = current1.Next
		current2 = current2.Next

	}
	return true
}

// This method is bit bad , 1. if yo have already craeted slice for LL , then its better to parse slice to check pallindrome
// therefore no need to back trace the linked list & reversing the pointers
// better approach : when just put half elements in slice
func (ll *LinkedList) IsPallindromeV2() bool {
	current := ll.Head
	var forward *ListNode
	var previous *ListNode

	forward = nil
	previous = nil
	sl := []int{}
	for current != nil {
		sl = append(sl, current.Val)
		forward = current.Next
		current.Next = previous
		previous = current
		current = forward
		if current == nil {
			fmt.Println("Reached last elemt ")
			current = previous
			n := 0

			for current != nil && n <= len(sl)-1 {
				fmt.Println("Reached last elemt ", current.Val, sl[n])

				if current.Val != sl[n] {
					return false
				}
				current = current.Next
				n++
			}

		}

	}
	return true

}

func (ll *LinkedList) IsPallindromeV3() bool {
	current := ll.Head
	pointer := ll.Head
	sl := []int{}
	for current != nil && current.Next != nil {
		sl = append(sl, pointer.Val)
		current = current.Next.Next
		pointer = pointer.Next

	}
	if current != nil && current.Next == nil {
		pointer = pointer.Next
	}
	n := len(sl) - 1
	for pointer != nil && n >= 0 {
		fmt.Println("Hello ", pointer.Val, sl[n])
		if pointer.Val != sl[n] {
			return false
		}
		pointer = pointer.Next
		n--
	}

	return true

}

func (ll *LinkedList) HasCycle() bool {

	slow := ll.Head
	fast := ll.Head.Next

	for fast != nil && fast.Next != nil && fast != slow {
		fmt.Println(slow.Val, fast.Val)

		slow = slow.Next
		fast = fast.Next.Next

	}
	if fast == slow {
		return true
	} else {
		return false
	}
}
