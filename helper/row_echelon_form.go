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
				a := &augmentedMatrix[pivot].Expression[j]
				b := &augmentedMatrix[col].Expression[j]

				tmp := a
				a = b
				b = tmp
			}
		}

		for row := col + 1; row < n; row++ {
			numerator := float64(augmentedMatrix[row].Expression[col])
			denominator := float64(augmentedMatrix[col].Expression[col])
			factor := numerator / denominator

			for j := col; j < augmentedSize; j++ {
				augmentedMatrix[row].Expression[j] -= factor * augmentedMatrix[col].Expression[j]
			}
		}
	}

	return augmentedMatrix
}
