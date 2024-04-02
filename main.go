package main

import "fmt"

// exercicio3
type funcionario struct {
	categoria rune
	horas     float64
}

func main() {
	//exercicio 1
	fmt.Printf("imposto do salario de 50000: %f \n", imposto(50000))
	fmt.Printf("imposto do salario de 50001: %f \n", imposto(50001))
	fmt.Printf("imposto do salario de 75000: %f \n", imposto(75000))
	fmt.Printf("imposto do salario de 150000: %f \n", imposto(150000))
	fmt.Printf("imposto do salario de 150001: %f \n", imposto(150001))

	//exercicio 2
	fmt.Printf("exemplo de media 1: %v \n", calculaMedia(3, 4, 5, 7, 8, 9, 4, 6))
	fmt.Printf("exemplo de media com nota negativa: %v \n", calculaMedia(3, 4, 5, 7, -8, 9, 4, 6))

	// exercicio 3
	func1 := funcionario{'C', 162}
	func2 := funcionario{'B', 176}
	func3 := funcionario{'A', 172}

	fmt.Printf("O salario do funcionario 1: %.2f \n", calculaSalario(func1))
	fmt.Printf("O salario do funcionario 2: %.2f \n", calculaSalario(func2))
	fmt.Printf("O salario do funcionario 3: %.2f \n", calculaSalario(func3))

	//exercicio 4
	oper := operation("minimum")
	r := oper(1, 2, 3, 4, 5)
	fmt.Printf("Exec 4 o valor minimo é: %v", r)
}

// exercicio 1
func imposto(salario float64) float64 {
	switch {
	case salario <= 50000:
		{
			return 0
		}
	case salario > 50000 && salario <= 150000:
		{
			return salario * 0.17
		}
	case salario > 150000:
		{
			return salario * 0.27
		}
	default:
		{
			return 0
		}
	}
}

//exercicio 2

// Precisamos criar uma função na qual possamos passar N quantidade de números inteiros e devolva a média.
func calculaMedia(notas ...int) int {

	numero_notas := 0
	notas_somadas := 0

	for _, nota := range notas {
		if nota > 0 {
			notas_somadas += nota
			numero_notas++
		} else {
			// Obs: Caso um dos números digitados seja negativo, a aplicação deve retornar um erro
			fmt.Printf("Numero de nota invalido: %v \n", nota)
			break
		}
	}

	return notas_somadas / numero_notas

}

// exercicio 3

func calculaSalario(pessoa funcionario) float64 {

	switch {
	case pessoa.categoria == 'C':
		return 1000 * pessoa.horas

	case pessoa.categoria == 'B':
		if pessoa.horas > 160 {
			adicionao := pessoa.horas - 160
			pagto := 1500.0 * 160
			pagto = pagto + (adicionao * 1500.0) + (adicionao * (1500.0 / 0.2))
			return pagto
		} else {
			return pessoa.horas * 1500.0
		}
	case pessoa.categoria == 'A':
		if pessoa.horas > 160 {
			adicionao := pessoa.horas - 160
			pagto := 3000.0 * 160
			pagto = pagto + (adicionao * 3000.0) + (adicionao * (3000.0 / 0.5))
			return pagto
		} else {
			return pessoa.horas * 3000.0
		}
	}

	return 0.0
}
