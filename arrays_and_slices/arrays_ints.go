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
