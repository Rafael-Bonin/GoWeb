package utils

import "errors"

func Divide(n1, n2 float64) (float64, error) {
	if n1 == 0 || n2 == 0 {
		return 0, errors.New("nao é possivel realizar divisao por 0")
	}

	return (n1 / n2), nil
}