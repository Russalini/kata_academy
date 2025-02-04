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

	stroka = strings.ReplaceAll(stroka, " ", "")

	if stroka[0] != '"' {
		panic("Первым аргументом должна быть строка.")
	}

	var parts []string
	var buffer strings.Builder
	Kavychki := false

	for _, char := range stroka {
		if char == '"' {
			Kavychki = !Kavychki
		}
		if !Kavychki && (char == '+' || char == '-' || char == '*' || char == '/') {
			parts = append(parts, buffer.String())
			parts = append(parts, string(char))
			buffer.Reset()
		} else {
			buffer.WriteRune(char)
		}
	}
	parts = append(parts, buffer.String())

	if len(parts) != 3 {
		panic("Неподдерживаемое выражение.")
	}

	left, op, right := parts[0], parts[1], parts[2]

	if left[0] != '"' || left[len(left)-1] != '"' {
		panic("Левый аргумент должен быть строкой.")
	}
	left = left[1 : len(left)-1]

	switch op {
	case "+":
		if right[0] != '"' || right[len(right)-1] != '"' {
			panic("Правый аргумент должен быть строкой.")
		}
		right = right[1 : len(right)-1]
		return plusStroka(left, right), nil
	case "-":
		if right[0] != '"' || right[len(right)-1] != '"' {
			panic("Правый аргумент должен быть строкой.")
		}
		right = right[1 : len(right)-1]
		return minusStroka(left, right), nil
	case "*":
		if right[0] == '"' && right[len(right)-1] == '"' {
			panic("Правый аргумент должен быть числом, а не строкой.")
		}
		n, err := strconv.Atoi(right)
		if err != nil || n < 1 || n > 10 {
			panic("Правый аргумент должен быть числом от 1 до 10.")
		}
		return umnojStroka(left, n), nil
	case "/":
		if right[0] == '"' && right[len(right)-1] == '"' {
			panic("Правый аргумент должен быть числом, а не строкой.")
		}
		n, err := strconv.Atoi(right)
		if err != nil || n < 1 || n > 10 {
			panic("Правый аргумент должен быть числом от 1 до 10.")
		}
		return delenieStroka(left, n), nil
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
