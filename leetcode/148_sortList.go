/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//merge merging two Lists as ascending
func merge(left *ListNode, right *ListNode) *ListNode {
	//when left merged after right
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	if left.Val <= right.Val {
		left.Next = merge(left.Next, right)
		return left
	} else {
		right.Next = merge(left, right.Next)
		return right
	}
}

//sortList divides List when List have only one element
//then merge List
func sortList(head *ListNode) *ListNode {
	//when cannot divide List
	if head == nil || head.Next == nil {
		return head
	}

	//divide List as left and right.
	//right List can have one more element than left List
	right := head
	tail := head
	var cutPoint *ListNode

	for tail != nil && tail.Next != nil {
		cutPoint = right
		right = right.Next
		tail = tail.Next.Next
	}
	cutPoint.Next = nil

	return merge(sortList(head), sortList(right))
}