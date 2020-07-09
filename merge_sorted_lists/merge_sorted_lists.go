package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var empty = false

	var lCopy = lists
	var result *ListNode = nil
	current := result
	for !empty {
		var smallest int
		var smallestIndex int
		first := true
		for i, l := range lCopy {
			if l != nil && (first || l.Val < smallest) {
				smallestIndex, smallest = i, l.Val
				first = false
			}
		}
		if first == false {
			node := &ListNode{smallest, nil}
			if current == nil {
				result = node
				current = node
			} else {
				current.Next = node
				current = current.Next
			}
			lists[smallestIndex] = lists[smallestIndex].Next
		} else {
			empty = true
		}
	}
	return result
}

func main() {
	list1 := &ListNode{1, &ListNode{3, &ListNode{6, &ListNode{12, nil}}}}
	list2 := &ListNode{3, &ListNode{9, &ListNode{13, &ListNode{15, nil}}}}
	assert("1,3,3,6,9,12,13,15", mergeKLists([]*ListNode{list1, list2}))

	list1 = &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, nil}}}}
	list2 = &ListNode{2, &ListNode{2, &ListNode{2, &ListNode{2, nil}}}}
	assert("1,1,1,1,2,2,2,2", mergeKLists([]*ListNode{list1, list2}))

	list1 = &ListNode{1, &ListNode{1, nil}}
	list2 = &ListNode{2, &ListNode{2, &ListNode{2, &ListNode{2, nil}}}}
	assert("1,1,2,2,2,2", mergeKLists([]*ListNode{list1, list2}))

	list1 = nil
	list2 = &ListNode{2, &ListNode{2, &ListNode{2, &ListNode{2, nil}}}}
	assert("2,2,2,2", mergeKLists([]*ListNode{list1, list2}))

	list1 = nil
	list2 = nil
	assert("nil", mergeKLists([]*ListNode{list1, list2}))

}

func assert(str string, nodes *ListNode) {
	var resStr string
	if nodes == nil {
		resStr = "nil"
	}
	for nodes != nil {
		if nodes.Next == nil {
			resStr += fmt.Sprintf("%d", nodes.Val)
		} else {
			resStr += fmt.Sprintf("%d,", nodes.Val)
		}
		nodes = nodes.Next
	}
	if resStr != str {
		panic("Expected: " + str + " Actual: " + resStr)
	}
}
