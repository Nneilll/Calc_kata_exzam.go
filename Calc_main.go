package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Println("Введите выражение (например, 5 + 2):")
	fmt.Scanln(&input)

	// Разделяем строку по пробелам
	parts := strings.Split(input, " ")

	if len(parts) != 3 {
		fmt.Println("Некорректное выражение. Попробуйте еще раз.")
		return
	}

	// Парсим числа
	num1, err1 := strconv.Atoi(parts[0])
	num2, err2 := strconv.Atoi(parts[2])
	operator := parts[1]

	if err1 != nil || err2 != nil {
		fmt.Println("Ошибка ввода чисел. Попробуйте еще раз.")
		return
	}

	// Проверка на 0
	if num1 == 0 || num2 == 0 {
		fmt.Println("Число 0 не допускается. Попробуйте еще раз.")
		return
	}

	// Выполнение операции
	switch operator {
	case "+":
		fmt.Printf("Результат: %d\n", num1+num2)
	case "-":
		fmt.Printf("Результат: %d\n", num1-num2)
	case "*":
		fmt.Printf("Результат: %d\n", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка: деление на 0!")
		} else {
			fmt.Printf("Результат: %d\n", num1/num2)
		}
	default:
		fmt.Println("Некорректный оператор. Попробуйте еще раз.")
	}
}
