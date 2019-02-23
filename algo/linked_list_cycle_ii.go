package algo

import "github.com/cxjwin/go-algocasts/datastructure"

// https://leetcode.com/problems/linked-list-cycle-ii/

func firstNodeOfCycleMap(head *ds.ListNode) *ds.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	m := make(map[*ds.ListNode]int)
	p := head

	for p.Next != nil {
		_, ok := m[p]
		if ok {
			return p
		}

		m[p] = p.Value.(int)
		p = p.Next
	}

	return nil
}

func firstNodeOfCycle2Pointer(head *ds.ListNode) *ds.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	s, f := head, head

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			p := head
			for {
				if p == s {
					return p
				}
				p = p.Next
				s = s.Next
			}
		}
	}

	return nil
}
