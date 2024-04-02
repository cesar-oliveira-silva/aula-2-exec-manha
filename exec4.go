package main

// const (
// 	minimum = "minimum"
// 	average = "average"
// 	maximum = "maximum"
// )

// minFunc, err := operation(minimum)
// averageFunc, err := operation(average)
// maxFunc, err := operation(maximum)

func operation(operacao string) func(valores ...int) int {
	switch operacao {
	case "minimum":
		return oppMinimum
	case "average":
		return oppAverage
	case "maximum":
		return oppMaximum

	}

	return nil

}

func oppMinimum(valores ...int) int {

	menor := 0
	for _, valor := range valores {
		if valor < menor {
			menor = valor
		}
	}
	return menor
}
func oppMaximum(valores ...int) int {

	maior := 0
	for _, valor := range valores {
		if valor > maior {
			maior = valor
		}
	}
	return maior
}

func oppAverage(valores ...int) int {
	soma := 0
	count := 0
	for _, valor := range valores {
		soma += valor
		count++
	}
	media := soma / count
	return media
}
