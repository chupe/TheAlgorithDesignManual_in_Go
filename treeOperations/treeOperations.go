package treeOperations

func Insert(t *Tree, v int) *Tree {
	if t.Value == 0 {
		t.Value = v
	}
	if v > t.Value {
		if t.Right == nil {
			t.Right = &Tree{
				Value: v,
			}
		} else {
			Insert(t.Right, v)
		}
	}
	if v < t.Value {
		if t.Left == nil {
			t.Left = &Tree{
				Value: v,
			}
		} else {
			Insert(t.Left, v)
		}
	}

	return t
}

func Search(t *Tree, v int) *Tree {
	if t == nil {
		return nil
	}
	if t.Value == v {
		return t
	}
	if v < t.Value {
		return Search(t.Left, v)
	}
	if v > t.Value {
		return Search(t.Right, v)
	}

	return nil
}

func Min(t *Tree) *Tree {
	if t == nil {
		return nil
	}
	min := t
	if t.Left != nil {
		min = Min(t.Left)
	}

	return min
}

func TraverseTree(t *Tree) []int {
	if t == nil {
		return nil
	}
	var result []int
	result = append(result, TraverseTree(t.Left)...)
	result = append(result, t.Value)
	result = append(result, TraverseTree(t.Right)...)

	return result
}

func Delete(t *Tree, v int) *Tree {
	td := Search(t, v)
	if t == nil || td == nil {
		return nil
	}
	if td.Left == nil && td.Right == nil {
		parent := td.Parent(t)
		if parent == nil {
			return nil
		}
		if td.Value > parent.Value {
			parent.Right = nil
		} else {
			parent.Left = nil
		}
		return t
	}

	if td.Left == nil && td.Right != nil {
		parent := td.Parent(t)
		if td.Value > parent.Value {
			parent.Right = td.Right
		} else {
			parent.Left = td.Right
		}
	}

	if td.Left != nil && td.Right == nil {
		parent := td.Parent(t)
		if td.Value > parent.Value {
			parent.Right = td.Left
		} else {
			parent.Left = td.Left
		}
	}

	if td.Left != nil && td.Right != nil {
		parent := td.Parent(t)
		successor := Min(td.Right)
		Delete(td, successor.Value)
		successor.Left = td.Left
		successor.Right = td.Right

		if td.Value > parent.Value {
			parent.Right = successor
		} else {
			parent.Left = successor
		}
	}

	return t
}
