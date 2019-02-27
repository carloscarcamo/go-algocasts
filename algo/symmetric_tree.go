package algo

import . "github.com/cxjwin/go-algocasts/datastructure"

import (
	"github.com/golang-collections/collections/stack"
)

func isSymmetric(l *TreeNode, r *TreeNode) bool {
	if l != nil && r != nil {
		return l.Val == r.Val &&
			isSymmetric(l.Left, r.Right) &&
			isSymmetric(l.Right, r.Left)
	}

	return l == nil && r == nil
}

func isSymmetricTreeRecursive(t *TreeNode) bool {
	if t == nil {
		return true
	}

	return isSymmetric(t.Left, t.Right)
}

func isSymmetricTreeIterative(t *TreeNode) bool {
	if t == nil {
		return true
	}

	s := new(stack.Stack)
	s.Push(t.Left)
	s.Push(t.Right)

	for s.Len() != 0 {
		l := s.Pop()
		r := s.Pop()
		if l == nil && r == nil {
			continue
		}

		if l == nil || r == nil {
			return false
		}

		tl := l.(*TreeNode)
		tr := r.(*TreeNode)
		if tl.Val != tr.Val {
			return false
		}
		if tl.Left != nil {
			s.Push(tl.Left)
		}
		if tr.Right != nil {
			s.Push(tr.Right)
		}
		if tl.Right != nil {
			s.Push(tl.Right)
		}
		if tr.Left != nil {
			s.Push(tr.Left)
		}
	}

	return true
}
