package insertionSort

var testCases = []struct {
	s1          []string
	want        []string
	expectError bool
}{
	{
		[]string{"a"},
		[]string{"a"},
		false,
	},
	{
		[]string{"i", "n", "s", "e", "r", "t", "i", "o", "n", "s", "o", "r", "t"},
		[]string{"e", "i", "i", "n", "n", "o", "o", "r", "r", "s", "s", "t", "t"},
		false,
	},
	{
		[]string{"bca", "bfe", "cda", "cdc"},
		[]string{"bca", "bfe", "cda", "cdc"},
		false,
	},
}
