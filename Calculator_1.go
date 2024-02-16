package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

/*func intToRoman(num int) string {
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
} */

func intToRoman(num int) string {
	arabicNumbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanNumerals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""

	for i := 0; i < len(arabicNumbers); i++ {
		for num >= arabicNumbers[i] {
			roman += romanNumerals[i]
			num -= arabicNumbers[i]
		}
	}

	return roman
}

func main() {

	var stroka, chis1, arf, chis2 string

	fmt.Println("Введите ваш пример")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stroka = scanner.Text()
	fmt.Println()
	fmt.Println(stroka)

	number := strings.Split(stroka, " ")

	if len(number) != 3 {
		panic("Формат введенных данных не соответствует изначально заявленным")
	}

	chis1 = number[0]
	arf = number[1]
	chis2 = number[2]

	isRoman := checkRoman(chis1)
	isArabic := checkArabic(chis1)

	if isRoman && !isArabic {
		fmt.Println("Введено римское число")
	} else if !isRoman && isArabic {
		fmt.Println("Введено арабское число")
	} else {
		panic("Некорректное число")
	}

	if isArabic == true {

		a1, _ := strconv.Atoi(chis1)
		b1, _ := strconv.Atoi(chis2)

		if !(1 <= a1 && a1 <= 10) || !(1 <= b1 && b1 <= 10) {
			panic("Значения вводимых чисел не соответствуют заявленным границам")
		}

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

		if !(1 <= a1 && a1 <= 10) || !(1 <= b1 && b1 <= 10) {
			panic("Значения вводимых чисел не соответствуют заявленным границам")
		}

		switch arf {
		case "+":
			fmt.Println(intToRoman(a1 + b1))
			//fmt.Println(a1 + b1)
		case "-":
			if a1 < b1 {
				panic("Первое число меньше второго")
			} else {
				fmt.Println(intToRoman(a1 - b1))
				//fmt.Println(a1 - b1)
			}
		case "*":
			fmt.Println(intToRoman(a1 * b1))
			//fmt.Println(a1 * b1)
		case "/":
			if a1 < b1 || a1%b1 != 0 {
				panic("Первое число меньше второго или вы хотите поделить с остатком")
			} else {
				fmt.Println(intToRoman(a1 / b1))
				//fmt.Println(a1 - b1)
			}

		}
	}

}
