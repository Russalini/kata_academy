package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func plus_stroka(a, b string) string {
	cnt := 0
	if len(a) > 10 || len(b) > 10 {
		panic("Ошибка: строки не может быть больше 10 символов")
	}
	_, err := strconv.Atoi(a)
	if err == nil {
		cnt++
	}
	_, err2 := strconv.Atoi(b)
	if err2 != nil && cnt > 0 {
		panic("Первая строка не может быть числом!")
	} else {
		return strconv.Quote(a + b)
	}
}

func minus_stroka(a, b string) string {
	cnt := 0
	if len(a) > 10 || len(b) > 10 {
		panic("Ошибка: строки не может быть больше 10 символов")
	}
	_, err := strconv.Atoi(a)
	if err == nil {
		cnt++
	}
	_, err2 := strconv.Atoi(b)
	if err2 != nil && cnt > 0 {
		panic("Первая строка не может быть числом!")
	} else {
		return strconv.Quote(strings.Replace(a, b, "", -1))
	}
}

func umnoj_stroka(a string, n int) string {
	if len(a) > 10 {
		panic("Ошибка: строки не может быть больше 10 символов")
	}
	if n > 0 && n < 11 {
		return strconv.Quote(strings.Repeat(a, n))
	} else {
		panic("Ошибка: число может только принадлежать отрезку от 0 до 10 включительно.")
	}
}

func delenie_stroka(a string, n int) string {
	if len(a) > 10 {
		panic("Ошибка: строки не может быть больше 10 символов")
	}
	if n > 0 && n < 11 {
		return strconv.Quote(a[:len(a)/n])
	} else {
		panic("Ошибка: число может только принадлежать отрезку от 0 до 10 включительно.")
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		panic("Ошибка ввода. В случае умножения - используйте (x) вместо (*)")
	}

	str := args[0]
	op := args[1]
	num, _ := strconv.Atoi(args[2])

	var result string
	switch op {
	case "+":
		result = plus_stroka(str, args[2])
	case "-":
		result = minus_stroka(str, args[2])
	case "x":
		result = umnoj_stroka(str, num)
	case "/":
		result = delenie_stroka(str, num)
	default:
		panic("Операция не распознана. Используйте что-то одно из этого: +, -, x, /")
	}

	if len(result) > 40 {
		result = result[:40] + "..."
	}

	fmt.Println(result)
}
