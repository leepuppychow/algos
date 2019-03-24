package linked_list

// Add two numbers. Each num is represented in REVERSE order as a linked list
// Return linked list as the solution

// WORKS, but should refactor (iterative solution is probably more elegant)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := ListNode{
		Val:  0,
		Next: nil,
	}
	addTwo(l1, l2, &result)
	return &result
}

func addTwo(l1, l2, current *ListNode) (*ListNode, *ListNode, *ListNode) {
	if l1 == nil && l2 == nil {
		return nil, nil, nil
	}
	sum := current.Val

	if l1 != nil && l2 == nil {
		sum += l1.Val
		prepNextDigit(sum, l1.Next, nil, current)
		return addTwo(l1.Next, nil, current.Next)
	} else if l2 != nil && l1 == nil {
		sum += l2.Val
		prepNextDigit(sum, nil, l2.Next, current)
		return addTwo(nil, l2.Next, current.Next)
	} else {
		sum += (l1.Val + l2.Val)
		prepNextDigit(sum, l1.Next, l2.Next, current)
		return addTwo(l1.Next, l2.Next, current.Next)
	}
}

func prepNextDigit(sum int, l1Next, l2Next, current *ListNode) {
	current.Val = sum % 10
	if l1Next == nil && l2Next == nil && sum/10 == 0 {
		return
	}
	current.Next = &ListNode{
		Val:  sum / 10,
		Next: nil,
	}
}