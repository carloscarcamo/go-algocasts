package algo

import "github.com/cxjwin/go-algocasts/datastructure"

// https://leetcode.com/problems/linked-list-cycle/

func linkedListHasCycle(head *ds.ListNode) bool {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}
