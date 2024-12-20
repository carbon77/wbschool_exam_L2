package main

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
  - "a4bc2d5e" => "aaaabccddddde"
  - "abcd" => "abcd"
  - "45" => "" (некорректная строка)
  - "" => ""

Дополнительное задание: поддержка escape - последовательностей
  - qwe\4\5 => qwe45 (*)
  - qwe\45 => qwe44444 (*)
  - qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
var (
	digits = []rune("0123456789")
)

func Unpack(input string) (string, error) {
	runes := []rune(input)
	var result string
	var escape bool

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if r == '\\' && !escape {
			if i == len(runes)-1 {
				return "", errors.New("invalid input: string not terminated")
			}
			escape = true
			continue
		}

		if slices.Contains(digits, r) && !escape {
			return "", errors.New("invalid input: two unescaped digits in a row")
		}

		count := 1
		if i < len(runes)-1 && slices.Contains(digits, runes[i+1]) {
			count, _ = strconv.Atoi(string(runes[i+1]))
			i++
		}

		result += strings.Repeat(string(r), count)
		escape = false
	}
	return result, nil
}

func main() {
	input := "abcd\\"
	str, err := Unpack(input)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %s\n", str)
}
