package Arrays_and_Slices

func Sum(n []int) int {
	final := 0

	// Range like enumerate in Python where
	// i is the index position and v is the value
	// at the corresponding position. Means you can
	// stop indexing in an array and get at the value
	// more easily.

	for _, v := range n {

		final += v
	}
	return final
}

func SumAll(nums ...[]int) []int {
	//go over reach arg, call SUM and add that result to a final slice.
	var final []int
	for _, v := range nums {
		final = append(final, Sum(v))
	}

	return final
}

func SumAllTails(nums ...[]int) []int {
	var final []int
	for _, v := range nums {
		if len(v) == 0 {
			final = append(final, 0)
		} else {
			final = append(final, Sum(v[1:]))
		}
	}
	return final
}
