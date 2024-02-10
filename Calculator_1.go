package main

import (
	"fmt"
	"strconv"
)

func checkRoman(input string) bool {
	romanSymbols := "IVXLCDM"
	for _, c := range input {
		if !contains(romanSymbols, string(c)) {
			return false
		}
	}
	return true
}

func checkArabic(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

func contains(s string, char string) bool {
	for _, c := range s {
		if string(c) == char {
			return true
		}
	}
	return false
}

func romanToInt(roman string) int {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prev := 0
	for i := len(roman) - 1; i >= 0; i-- {
		current := romanNumerals[rune(roman[i])]
		if prev > current {
			result -= current
		} else {
			result += current
		}
		prev = current
	}

	return result
}

func intToRoman(num int) string {
	// Создаем словарь соответствия арабских чисел и римских цифр
	romanNumerals := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	// Создаем переменную для хранения результата
	roman := ""

	for arabic, numeral := range romanNumerals {
		// Пока арабское число больше или равно текущему значению
		for num >= arabic {
			// Добавляем соответствующую римскую цифру в результат
			roman += numeral
			// Вычитаем значение арабского числа из исходного числа
			num -= arabic
		}
	}

	return roman
}

func main() {

	var chis1, arf, chis2 string
	fmt.Println("Введите ваш пример")
	fmt.Scanln(&chis1, &arf, &chis2)
	fmt.Println()
	fmt.Println(chis1, arf, chis2)

	isRoman := checkRoman(chis1)
	isArabic := checkArabic(chis1)

	if isRoman && !isArabic {
		fmt.Println("Введено римское число")
	} else if !isRoman && isArabic {
		fmt.Println("Введено арабское число")
	} else {
		fmt.Println("Некорректное число")
	}

	if isArabic == true {

		a1, _ := strconv.Atoi(chis1)
		b1, _ := strconv.Atoi(chis2)

		switch arf {
		case "+":
			fmt.Println(a1 + b1)
		case "-":
			fmt.Println(a1 - b1)
		case "*":
			fmt.Println(a1 * b1)
		case "/":
			fmt.Println(a1 / b1)

		}
	} else if isRoman == true {

		a1 := romanToInt(chis1)
		b1 := romanToInt(chis2)

		switch arf {
		case "+":
			fmt.Println(intToRoman(a1 + b1))
		case "-":
			if a1 < b1 {
				fmt.Println("Ошибка")
			} else {
				fmt.Println(intToRoman(a1 - b1))
			}
		case "*":
			fmt.Println(intToRoman(a1 * b1))
		case "/":
			if a1 < b1 || a1%b1 != 0 {
				fmt.Println("Ошибка")
			} else {
				fmt.Println(intToRoman(a1 / b1))
			}

		}
	}

}
