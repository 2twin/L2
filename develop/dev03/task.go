package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Конфиг нашего кастомного сорта
type customSort struct {
	lines         []string
	key           int
	numeric       bool
	reverse       bool
	unique        bool
	byMonth       bool
	ignoreBlanks  bool
	checkSort     bool
	numericSuffix bool
}

// Методы кастомного сорта для реализации интерфейса sort.Interface
func (cs *customSort) Len() int {
	return len(cs.lines)
}

func (cs *customSort) Swap(i, j int) {
	cs.lines[i], cs.lines[j] = cs.lines[j], cs.lines[i]
}

// Основной метод сорировки
func (cs *customSort) Less(i, j int) bool {
	// Забираем текущую и предыдущую строку по определенному столбцу
	line1 := getColumnValue(cs.lines[i], cs.key)
	line2 := getColumnValue(cs.lines[j], cs.key)

	// Если стоит флаг -b, убираем хвостовые пробелы у строк
	if cs.ignoreBlanks {
		line1 = strings.TrimSpace(line1)
		line2 = strings.TrimSpace(line2)
	}

	// Если стоит флаг -M, конвертируем строки в имена месяцев
	if cs.byMonth {
		line1 = convertToMonthName(line1)
		line2 = convertToMonthName(line2)
	}

	// Если стоит флаг -h, конвертируем строки в числа без суффиксов
	if cs.numericSuffix {
		line1, line2 = convertToNumeric(convertToNumericSuffix(line1), convertToNumericSuffix(line2))
		// line2 = convertToNumericSuffix(line2)
	}

	// Если стоит флаг -n, конвертируем строки в числа, и если
	// не возникло ошибок, конвертируем обратно в строку
	if cs.numeric {
		num1, err1 := strconv.Atoi(line1)
		num2, err2 := strconv.Atoi(line2)

		if err1 == nil && err2 == nil {
			line1 = fmt.Sprintf("%064d", num1)
			line2 = fmt.Sprintf("%064d", num2)
		}
	}

	// Если стоит флаг -u, проверяем равны ли 2 соседние строки
	if cs.unique {
		return line1 != line2
	}

	// Если стоит флаг -r, сортируем в обратном порядке
	if cs.reverse {

		// Если стоит флаг -c, проверям сортировку (в обратном порядке), и если находим
		// неотсортированный участок, печатаем текущую строку и завершаем программу
		if cs.checkSort && line1 < line2 {
			log.Fatal("disorder: ", line1)
		}

		return line1 > line2
	}

	// Если стоит флаг -c, проверям сортировку, и если находим
	// неотсортированный участок, печатаем текущую строку и завершаем программу
	if cs.checkSort && line1 < line2 {
		log.Fatal("disorder: ", line1)
	}

	return line1 < line2
}

// getColumnValue возвращает строку по определенному столбцу
func getColumnValue(line string, key int) string {
	cols := strings.Fields(line)
	if key > 0 && key <= len(cols) {
		return cols[key-1]
	}

	return line
}

// convertToMonthName конвертирует строку в порядковый номер месяца
func convertToMonthName(line string) string {
	t, err := time.Parse("Jan", line)
	if err != nil {
		return line
	}

	return fmt.Sprintf("%02d", t.Month())
}

// convertToNumericSuffix конвертирует строку в числовой вид без суффиксов
func convertToNumericSuffix(line string) string {
	if strings.HasSuffix(line, "K") {
		num, _ := strconv.Atoi(strings.TrimSuffix(line, "K"))
		return fmt.Sprintf("%064d", num * 1_000)
	}

	if strings.HasSuffix(line, "M") {
		num, _ := strconv.Atoi(strings.TrimSuffix(line, "M"))
		return fmt.Sprintf("%064d", num * 1_000_000)
	}

	if strings.HasSuffix(line, "G") {
		num, _ := strconv.Atoi(strings.TrimSuffix(line, "G"))
		return fmt.Sprintf("%064d", num * 1_000_000_000)
	}

	return line
}

// convertToNumeric конвертирует строки в числовые значения
func convertToNumeric(line1, line2 string) (string, string) {
	num1, err1 := strconv.Atoi(line1)
	num2, err2 := strconv.Atoi(line2)

	if err1 == nil && err2 == nil {
		return fmt.Sprintf("%064d", num1), fmt.Sprintf("%064d", num2)
	}

	return line1, line2
}

// sortLine сортирует строки по заданным параметрам конфига
func sortLines(lines []string, key int, numeric, reverse, byMonth, ignoreBlanks, checkSort, numericSuffix, unique bool) []string {
	cs := customSort{
		lines:         lines,
		key:           key,
		numeric:       numeric,
		reverse:       reverse,
		byMonth:       byMonth,
		ignoreBlanks:  ignoreBlanks,
		checkSort:     checkSort,
		numericSuffix: numericSuffix,
		unique:        unique,
	}

	sort.Sort(&cs)
	return cs.lines
}

func main() {
	// Задаем все возможные флаги
	key := flag.Int("k", -1, "Column sort by (1-indexed)")
	numeric := flag.Bool("n", false, "Sort numerically")
	reverse := flag.Bool("r", false, "Sort in reversed order")
	byMonth := flag.Bool("M", false, "Sort by month")
	ignoreBlanks := flag.Bool("b", false, "Ignore trailing spaces")
	checkSort := flag.Bool("c", false, "Check if the input is already sorted")
	numericSuffix := flag.Bool("h", false, "Compare human-readable numbers (e.g. 2K, 10M)")
	unique := flag.Bool("u", false, "Suppress lines that appear more than once")

	flag.Parse()

	// Забираем название файла (последний аргумент командной строки)
	inputFile := os.Args[len(os.Args)-1]

	inputFileHandle, err := os.Open(inputFile)
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}
	defer inputFileHandle.Close()

	var lines []string
	scanner := bufio.NewScanner(inputFileHandle)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading input file:", err)
	}

	sortedLines := sortLines(lines, *key, *numeric, *reverse, *byMonth, *ignoreBlanks, *checkSort, *numericSuffix, *unique)

	outputFileHandle, err := os.Create("out.txt")
	if err != nil {
		log.Fatal("Error creating output file:", err)
	}
	defer outputFileHandle.Close()

	writer := bufio.NewWriter(outputFileHandle)

	for _, line := range sortedLines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			log.Fatal("Error writing to output file:", err)
		}
	}

	writer.Flush()

	log.Println("Sort completed")
}
