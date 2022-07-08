package treeOperations

import "fmt"

type Tree struct {
	Value int
	Right *Tree
	Left  *Tree
}

func (t1 *Tree) Parent(p *Tree) *Tree {
	if p == nil {
		return nil
	}
	if p.Left.Compare(t1) || p.Right.Compare(t1) {
		return p
	}
	if t1.Value > p.Value {
		return t1.Parent(p.Right)
	} else {
		return t1.Parent(p.Left)
	}
}

func (t1 *Tree) String() string {
	value := fmt.Sprintf("%d\n", t1.Value)
	if t1.Left != nil {
		value += fmt.Sprintf("L: %s", t1.Left.String())
	}
	if t1.Right != nil {
		value += fmt.Sprintf("R: %s", t1.Right.String())
	}
	return value
}

func (t1 *Tree) Compare(t2 *Tree) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}

	leftCmp := t1.Left == t2.Left
	rightCmp := t1.Right == t2.Right

	if t1.Value != t2.Value {
		return false
	}

	if t1.Left != nil && t2.Left != nil {
		leftCmp = t1.Left.Compare(t2.Left)
		if !leftCmp {
			return false
		}
	}

	if t1.Right != nil && t2.Right != nil {
		rightCmp = t1.Right.Compare(t2.Right)
		if !rightCmp {
			return false
		}
	}

	return leftCmp && rightCmp
}
