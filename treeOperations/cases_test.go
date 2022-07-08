package treeOperations

var insertTestCases = []struct {
	t           *Tree
	v           int
	want        *Tree
	expectError bool
}{
	{
		&Tree{},
		5,
		&Tree{
			Value: 5,
		},
		false,
	},
	{
		&Tree{
			Value: 5,
		},
		10,
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
			},
		},
		false,
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
			},
		},
		3,
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
			},
			Left: &Tree{
				Value: 3,
			},
		},
		false,
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
			},
			Left: &Tree{
				Value: 3,
			},
		},
		7,
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
			},
			Left: &Tree{
				Value: 3,
			},
		},
		false,
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
			},
			Left: &Tree{
				Value: 3,
			},
		},
		12,
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
			},
		},
		false,
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
			},
		},
		2,
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
				},
			},
		},
		false,
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
				},
			},
		},
		4,
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
				},
				Right: &Tree{
					Value: 4,
				},
			},
		},
		false,
	},
}

var compareTestCases = []struct {
	t1   *Tree
	t2   *Tree
	want bool
}{
	{
		&Tree{
			Value: 5,
		},
		&Tree{
			Value: 5,
		},
		true,
	},
	{
		&Tree{
			Value: 5,
			Left: &Tree{
				Value: 3,
			},
		},
		&Tree{
			Value: 5,
			Left: &Tree{
				Value: 3,
			},
		},
		true,
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 3,
			},
		},
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 3,
			},
		},
		true,
	},
	{
		&Tree{
			Value: 6,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
				},
				Right: &Tree{
					Value: 4,
				},
			},
		},
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
				},
				Right: &Tree{
					Value: 4,
				},
			},
		},
		false,
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
				},
				Right: &Tree{
					Value: 4,
				},
			},
		},
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
				},
				Right: &Tree{
					Value: 1,
				},
			},
		},
		false,
	},
	{&Tree{
		Value: 5,
		Right: &Tree{
			Value: 10,
			Left: &Tree{
				Value: 7,
			},
			Right: &Tree{
				Value: 12,
			},
		},
		Left: &Tree{
			Value: 3,
			Left: &Tree{
				Value: 2,
			},
			Right: &Tree{
				Value: 4,
			},
		},
	},
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
				},
				Right: &Tree{
					Value: 4,
				},
			},
		},
		true,
	},
}

var searchTestCases = []struct {
	t    *Tree
	v    int
	want *Tree
}{
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
				},
				Right: &Tree{
					Value: 4,
				},
			},
		},
		5,
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
				},
				Right: &Tree{
					Value: 4,
				},
			},
		},
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
				},
				Right: &Tree{
					Value: 4,
				},
			},
		},
		10,
		&Tree{
			Value: 10,
			Left: &Tree{
				Value: 7,
			},
			Right: &Tree{
				Value: 12,
			},
		},
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
				},
				Right: &Tree{
					Value: 4,
				},
			},
		},
		4,
		&Tree{
			Value: 4,
		},
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
				},
				Right: &Tree{
					Value: 4,
				},
			},
		},
		8,
		nil,
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
				},
				Right: &Tree{
					Value: 4,
				},
			},
		},
		13,
		nil,
	},
}

var minTestCases = []struct {
	t    *Tree
	want *Tree
}{
	{
		&Tree{
			Value: 5,
		},
		&Tree{
			Value: 5,
		},
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
		},
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
		},
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
				},
				Right: &Tree{
					Value: 4,
				},
			},
		},
		&Tree{
			Value: 2,
		},
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
					Left: &Tree{
						Value: 1,
					},
				},
				Right: &Tree{
					Value: 4,
				},
			},
		},
		&Tree{
			Value: 1,
		},
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Right: &Tree{
					Value: 4,
				},
			},
		},
		&Tree{
			Value: 3,
			Right: &Tree{
				Value: 4,
			},
		},
	},
}

var traverseTestCases = []struct {
	t    *Tree
	want []int
}{
	{
		&Tree{
			Value: 5,
		},
		[]int{5},
	},
	{
		&Tree{
			Value: 5,
			Left: &Tree{
				Value: 2,
			},
		},
		[]int{2, 5},
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 6,
			},
		},
		[]int{5, 6},
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
					Left: &Tree{
						Value: 1,
					},
				},
				Right: &Tree{
					Value: 4,
				},
			},
		},
		[]int{1, 2, 3, 4, 5, 7, 10, 12},
	},
}

var deleteTestCases = []struct {
	t           *Tree
	v           int
	want        *Tree
	expectError bool
}{
	{
		&Tree{
			Value: 5,
		},
		5,
		nil,
		false,
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
			},
		},
		10,
		&Tree{
			Value: 5,
		},
		false,
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
			},
			Left: &Tree{
				Value: 3,
			},
		},
		3,
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
			},
		},
		false,
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
			},
			Left: &Tree{
				Value: 3,
			},
		},
		10,
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 7,
			},
			Left: &Tree{
				Value: 3,
			},
		},
		false,
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
					Left: &Tree{
						Value: 11,
					},
				},
			},
			Left: &Tree{
				Value: 3,
			},
		},
		10,
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 11,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
			},
		},
		false,
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
				},
			},
		},
		2,
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
			},
		},
		false,
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
				},
				Right: &Tree{
					Value: 4,
				},
			},
		},
		3,
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 4,
				Left: &Tree{
					Value: 2,
				},
			},
		},
		false,
	},
}

var findParentTestCases = []struct {
	p    *Tree
	c    *Tree
	want *Tree
}{
	{
		&Tree{
			Value: 5,
		},
		&Tree{
			Value: 5,
		},
		nil,
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 2,
				Right: &Tree{
					Value: 4,
				},
			},
		},
		&Tree{
			Value: 10,
			Left: &Tree{
				Value: 7,
			},
			Right: &Tree{
				Value: 12,
			},
		},
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 2,
				Right: &Tree{
					Value: 4,
				},
			},
		},
	},
	{
		&Tree{
			Value: 5,
			Right: &Tree{
				Value: 10,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 12,
				},
			},
			Left: &Tree{
				Value: 2,
				Right: &Tree{
					Value: 4,
				},
			},
		},
		&Tree{
			Value: 4,
		},
		&Tree{
			Value: 2,
			Right: &Tree{
				Value: 4,
			},
		},
	},
}
