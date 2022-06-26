package robotTourOptimization

type cases struct {
	points      []Point
	want        float64
	expectError bool
}

var testCasesUnoptimized = []cases{
	{
		[]Point{
			{0, 0},
			{0, -1},
			{0, 1},
			{0, 3},
			{0, -5},
			{0, 11},
			{0, -21},
		},
		82,
		false,
	},
	{
		[]Point{
			{-2, -2},
			{0, 2},
			{0, -2},
			{2, 0},
			{-2, 0},
			{-2, 2},
			{2, -2},
			{2, 2},
		},
		16,
		false,
	},
}
