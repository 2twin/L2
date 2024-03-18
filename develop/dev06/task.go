package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/


// getFields возвращает массив чисел из строкового флага fields
func getFields(f string) ([]int, error) {
	var res []int
	split := strings.Split(f, ",")

	for _, str := range split {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}

		res = append(res, num)
	}

	return res, nil
}


// cut реализует функционал утилиты cut
func cut(fields []int, delimiter string, separated bool) {
	// создаем сканнер stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// под каждую новую строку в stdin создаем переменную и разбиваем ее по delimiter
		line := scanner.Text()
		parts := strings.Split(line, delimiter)

		// если стоит флаг "-s" и очередная строка не содержит наш delimiter, пропускаем эту строку
        if separated && !strings.Contains(line, delimiter) {
            continue
        }

		// создаем массив для хранения найденных строк
        var selectedFields []string
		// для каждого значения fields проверяем если оно меньше 0,
		// значит флаг не стоит, следовательно сохраняем всю строку в наш массив
        for _, field := range fields {
			if field < 0 {
				selectedFields = append(selectedFields, line)
				continue
			}

			// если field больше нуля и меньше длины строки, добавляем в массив нужную часть строки
            if field <= len(parts) {
                selectedFields = append(selectedFields, parts[field-1])
            }
        }
		// выводим результат в stdout, объединяя массив с результатами в одну строку
        fmt.Println(strings.Join(selectedFields, delimiter))
	}

	if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error reading standard input:", err)
    }
}

func main() {
	fields := flag.String("f", "-1", "Select fields (columns)")
	delimiter := flag.String("d", "\t", "Use different delimiter")
	separated := flag.Bool("s", false, "Only separated strings")

	flag.Parse()

	f, err := getFields(*fields)
	if err != nil {
		log.Fatal(err)
	}

	/*
		пример использования:
		```
			echo -e "1, 2, 3\n4, 5, 6\n" | go run task.go -f 1,3 -d ","
		```
	*/
	cut(f, *delimiter, *separated)
}
