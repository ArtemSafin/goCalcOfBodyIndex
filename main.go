package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Калькулятор индекса массы тела")
	weight, height := getUserInput()
	index := calculateIndex(weight, height)
	fmt.Println(outputResult(index))

}

func getUserInput() (float64, float64) {
	var weight, height float64
	fmt.Print("Введите ваш рост:")
	fmt.Scan(&height)
	fmt.Print("Введите ваш вес:")
	fmt.Scan(&weight)
	return weight, height
}

func calculateIndex(weight float64, height float64) float64 {
	index := weight / math.Pow(height/100, 2)
	return index
}

func getCategory(index float64) string {
	var category string

	switch {
	case index < 16:
		category = "Сильный дефицит"
	case index >= 16 && index < 19:
		category = "Дефицит"
	case index >= 19 && index <= 25:
		category = "Норма"
	case index > 25 && index <= 30:
		category = "Избыток массы"
	case index > 30 && index <= 35:
		category = "Ожирение"
	default:
		category = "Сильное ожирение (индекс > 35)"
	}

	return category
}

func outputResult(index float64) string {
	category := getCategory(index)
	return fmt.Sprintf("РЕЗУЛЬТАТ: Индекс массы тела: %.1f\nКатегория: %s", index, category)
}
