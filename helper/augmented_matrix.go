package helper

type AugmentedEquation struct {
	Expression []int
}

func CreateAugmentedMatrix(equations []Equation) []AugmentedEquation {
	var augmentedMatrix []AugmentedEquation
	for _, equation := range equations {
		var req AugmentedEquation
		req.Expression = append(req.Expression, equation.expression...)
		req.Expression = append(req.Expression, equation.value)

		augmentedMatrix = append(augmentedMatrix, req)
	}

	return augmentedMatrix
}
