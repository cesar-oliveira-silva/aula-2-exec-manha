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
	min := Operation("minimum")
	r := min(1, 2, 3, 4, 5)
	fmt.Printf("Exec 4 o valor minimo é: %v \n", r)

	med := Operation("average")
	r = med(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Printf("Exec 4 o valor media é: %v \n", r)

	max := Operation("maximum")
	r = max(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Printf("Exec 4 o valor Maximo é: %v \n", r)

	//exercicio 5
	cachorros := Animal("Caes")
	f := cachorros(33.0)
	fmt.Printf("Exec 5 Cachorro precisa de: %v  kilos \n", f)

	gatos := Animal("Gatos")
	f = gatos(22.0)
	fmt.Printf("Exec 5 Gatos precisa de: %v  kilos \n", f)

	hamster := Animal("Hamster")
	f = hamster(5.0)
	fmt.Printf("Exec 5 Hamster precisa de: %v  kilos \n", f)

	aranha := Animal("Tarantula")
	f = aranha(5.0)
	fmt.Printf("Exec 5 Tarantula precisa de: %v  kilos \n", f)

	tartaruga := Animal("Tartaruga")
	f = tartaruga(6.0)
	fmt.Printf("Exec 5 Tartaruga precisa de: %v  kilos \n", f)
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

// exercicio 4
func Operation(operacao string) func(valores ...int) int {
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

	menor := 999999999999
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

// exercicio 5

func Animal(tipoAnimal string) func(quantidade float64) float64 {
	switch tipoAnimal {
	case "Caes":
		return alimentoCao
	case "Gatos":
		return alimentoGato
	case "Hamster":
		return alimentoHamster
	case "Tarantula":
		return alimentoTarantula
	default:
		return animalInexistente
	}

}

func alimentoCao(quantidade float64) float64 {
	return quantidade * 10
}

func alimentoGato(quantidade float64) float64 {
	return quantidade * 5
}

func alimentoHamster(quantidade float64) float64 {
	return quantidade * 0.250
}

func alimentoTarantula(quantidade float64) float64 {
	return quantidade * 0.150
}

func animalInexistente(quantidade float64) float64 {
	fmt.Println("Este animal nao esta cadastrado")
	return 0

}
