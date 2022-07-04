package findMatch

var testCases = []struct {
	s1          string
	s2          string
	want        int
	expectError bool
}{
	{
		"abaababbaba",
		"abba",
		5,
		false,
	},
	{
		"abaababababa",
		"abba",
		-1,
		false,
	},
	{
		"find a string match in this string",
		"this",
		23,
		false,
	},
	{
		"find a string match in this string",
		"abba",
		-1,
		false,
	},
	{
		"find a string match in this string",
		"find a",
		0,
		false,
	},
	{
		"find a string match in this string!",
		"!",
		34,
		false,
	},
}
