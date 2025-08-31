package helper

func BackSubstitution(REFMatrix []AugmentedEquation) map[int]float64 {

	variableMap := make(map[int]float64)

	rows := len(REFMatrix)
	cols := len(REFMatrix[0].Expression)

	for row := rows - 1; row >= 0; row-- {
		coeff := REFMatrix[row].Expression[row]
		if coeff == 0 {
			continue
		}

		sumAhead := 0.0
		for col := cols - 2; col > row; col-- {
			coeff := REFMatrix[row].Expression[col]
			sumAhead += coeff * variableMap[col]
		}

		value := REFMatrix[row].Expression[cols-1]
		variableValue := (value - sumAhead) / coeff

		variableMap[row] = variableValue
	}

	return variableMap
}
