package helper

func RowEchelonForm(augmentedMatrix []AugmentedEquation) []AugmentedEquation {

	n := len(augmentedMatrix)
	augmentedSize := len(augmentedMatrix[0].Expression)
	expressionSize := augmentedSize - 1

	for col := 0; col < expressionSize; col++ {
		pivot := col
		for row := col; row < n; row++ {
			if augmentedMatrix[row].Expression[col] > augmentedMatrix[pivot].Expression[col] {
				pivot = row
			}
		}

		if pivot != col {
			for j := col; j < expressionSize; j++ {
				augmentedMatrix[pivot].Expression[j], augmentedMatrix[col].Expression[j] =
					augmentedMatrix[col].Expression[j], augmentedMatrix[pivot].Expression[j]
			}
		}

		for row := col + 1; row < n; row++ {
			numerator := float64(augmentedMatrix[row].Expression[col])
			denominator := float64(augmentedMatrix[col].Expression[col])
			factor := numerator / denominator

			for j := col; j < expressionSize; j++ {
				currElement := float64(augmentedMatrix[row].Expression[j])

				currElement -= factor * denominator

				augmentedMatrix[row].Expression[j] = currElement
			}

		}
	}

	return augmentedMatrix
}
