package main

import (
	"errors"
	"fmt"
)

const USDeur = 0.92
const USDrub = 96
const EURrub = USDrub / USDeur
const eur = "eur"
const usd = "usd"
const rub = "rub"

func main() {

	var baseCurrency string
	var targetCurrency string
	var inputSum float64
	var result float64
	fmt.Println("_Конвертер валют_")
	for {
		userCurrency, err := getUserCurrency()
		if err != nil {
			fmt.Println("Ошибка.Введите правильное значение вашей валюты")
			continue
		}
		baseCurrency = userCurrency
		break
	}
	for {
		sum, err := getUserSum(baseCurrency)
		if err != nil {
			fmt.Println("Ошибка. Введите корректную сумму.")
			continue
		}
		inputSum = sum
		break

	}
	for {
		target, err := getTargetCurrency(baseCurrency)
		if err != nil {
			fmt.Println("Ошибка.Введите правильное значение нужной валюты")
			continue
		}
		targetCurrency = target
		break
	}
	result = calculate(baseCurrency, targetCurrency, inputSum)
	fmt.Printf("Вы получили %.2f %s ", result, targetCurrency)

}
func getUserCurrency() (string, error) {
	var baseCurrency string
	fmt.Println("Введите исходную валюту: 'eur' ,'usd' или 'rub'.")
	fmt.Scan(&baseCurrency)
	if baseCurrency == eur || baseCurrency == usd || baseCurrency == rub {
		return baseCurrency, nil
	}
	return "", errors.New("NO_PARAMS_ERROR")
}

func getTargetCurrency(base string) (string, error) {
	var message string
	var target string
	if base == eur {
		message = "Введите нужную валюту:'usd' или 'rub'."
	} else if base == usd {
		message = "Введите нужную валюту:'eur' или 'rub'."
	} else {
		message = "Введите нужную валюту:'eur' или 'usd'."
	}
	fmt.Println(message)
	fmt.Scan(&target)
	if target == eur || target == usd || target == rub && target != base {
		return target, nil
	}
	return "", errors.New("NO_PARAMS_ERROR")
}

func getUserSum(base string) (float64, error) {
	var sum float64
	fmt.Printf("Введите какое колличество %s вы хотите поменять.\n", base)
	fmt.Scan(&sum)
	if sum <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	return sum, nil
}

func calculate(baseCurrency string, targetCurrency string, sum float64) float64 {
	var result float64
	if baseCurrency == usd && targetCurrency == eur {
		result = sum * USDeur
	} else if baseCurrency == usd && targetCurrency == rub {
		result = sum * USDrub
	} else if baseCurrency == eur && targetCurrency == usd {
		result = sum / USDeur
	} else if baseCurrency == eur && targetCurrency == rub {
		result = sum * EURrub
	} else if baseCurrency == rub && targetCurrency == usd {
		result = sum / USDrub
	} else if baseCurrency == rub && targetCurrency == eur {
		result = sum / EURrub
	}
	return result
}
