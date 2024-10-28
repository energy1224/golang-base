package main

import "fmt"

func main() {
	const USDeur = 0.92
	const USDrub = 96
	const EURrub = USDrub / USDeur

	inputSumm := getUserInput()

	fmt.Print(inputSumm)

}
func getUserInput() float64 {
	var inputSumm float64
	fmt.Print("_Введите сумму_")
	fmt.Scan(&inputSumm)
	return inputSumm
	

}

func calculate(baseCurrency float64, targetCurrency float64, requiredSumm float64 )float64 {
        var result float64
	return result
}