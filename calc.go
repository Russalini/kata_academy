package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите пример:")

	if scanner.Scan() {
		expression := scanner.Text()
		result, err := calculator(expression)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Результат: %q\n", result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при чтении ввода:", err)
	}
}

func calculator(stroka string) (string, error) {

	if stroka[0] != '"' {
		panic("Первым аргументом должна быть строка.")
	}

	var parts []string
	var buffer strings.Builder
	Kavychki := false

	for _, znak := range stroka {
		if znak == '"' {
			Kavychki = !Kavychki
		}
		if !Kavychki && (znak == '+' || znak == '-' || znak == '*' || znak == '/') {
			parts = append(parts, buffer.String())
			parts = append(parts, string(znak))
			buffer.Reset()
		} else {
			buffer.WriteRune(znak)
		}
	}
	parts = append(parts, buffer.String())

	if len(parts) != 3 {
		panic("Неподдерживаемое выражение.")
	}

	left, op, right := parts[0], parts[1], parts[2]
	var leftCH, rightCH string
	dleft := len(left)
	for i := 0; i < dleft-1; i++ {
		leftCH += string(left[i])
	}
	dright := len(right)
	for i := 1; i < dright; i++ {
		rightCH += string(right[i])
	}

	if leftCH[0] != '"' || leftCH[len(leftCH)-1] != '"' {
		panic("Левый аргумент должен быть строкой.")
	}
	leftCH = leftCH[1 : len(leftCH)-1]

	switch op {
	case "+":
		if rightCH[0] != '"' || rightCH[len(rightCH)-1] != '"' {
			panic("Правый аргумент должен быть строкой.")
		}
		rightCH = rightCH[1 : len(rightCH)-1]
		return plusStroka(leftCH, rightCH), nil
	case "-":
		if rightCH[0] != '"' || rightCH[len(rightCH)-1] != '"' {
			panic("Правый аргумент должен быть строкой.")
		}
		rightCH = rightCH[1 : len(rightCH)-1]
		return minusStroka(leftCH, rightCH), nil
	case "*":
		if rightCH[0] == '"' && rightCH[len(rightCH)-1] == '"' {
			panic("Правый аргумент должен быть числом, а не строкой.")
		}
		n, err := strconv.Atoi(rightCH)
		if err != nil || n < 1 || n > 10 {
			panic("Правый аргумент должен быть числом от 1 до 10.")
		}
		return umnojStroka(leftCH, n), nil
	case "/":
		if rightCH[0] == '"' && rightCH[len(rightCH)-1] == '"' {
			panic("Правый аргумент должен быть числом, а не строкой.")
		}
		n, err := strconv.Atoi(rightCH)
		if err != nil || n < 1 || n > 10 {
			panic("Правый аргумент должен быть числом от 1 до 10.")
		}
		return delenieStroka(leftCH, n), nil
	default:
		panic("Неподдерживаемая операция.")
	}
}

func plusStroka(a, b string) string {
	result := a + b
	return sorok(result)
}

func minusStroka(a, b string) string {
	result := strings.Replace(a, b, "", 1)
	return sorok(result)
}

func umnojStroka(a string, n int) string {
	result := strings.Repeat(a, n)
	return sorok(result)
}

func delenieStroka(a string, n int) string {
	length := len(a) / n
	if length == 0 {
		return ""
	}
	result := a[:length]
	return sorok(result)
}

func sorok(s string) string {
	if len(s) > 40 {
		return s[:40] + "..."
	}
	return s
}
