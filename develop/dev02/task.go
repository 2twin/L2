package main

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

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrIncorrectString = errors.New("incorrect string")
)

func Unzip(s string) (string, error) {
	if len(s) == 0 {
		return "", ErrIncorrectString
	}

	if _, err := strconv.Atoi(s); err == nil {
		return "", ErrIncorrectString
	}

	arr := strings.Split(s, "")
	res := []string{}

	for _, r := range arr {
		n, err := strconv.Atoi(r)
		if err != nil {
			res = append(res, r)
			continue
		}

		if len(res) == 0 {
			break
		}

		i := len(res) - 1
		current := res[i]

		if current == `\` {
			if res[i-1] != `\` {
				res[i] = r
				continue
			}

			i -= 1
			n -= 1
		}

		s := strings.Repeat(current, n)
		res[i] = s
	}

	return strings.Join(res, ""), nil
}

func main() {
	str := []string{"a4bc2d5e", "abcd", "45", "", `qwe\4\5`, `qwe\45`, `qwe\\5`}
	for _, s := range str {
		fmt.Println(Unzip(s))
	}
}
