package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type grepConfig struct {
	lines      []string
	after      int
	before     int
	ctx        int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	result     map[int]string
	counter    int
}

func grep(cfg *grepConfig, target string) {
	for i, line := range cfg.lines {
		if cfg.ignoreCase {
			line = strings.ToLower(line)
			target = strings.ToLower(target)
		}

		if cfg.fixed {
			if line == target {
				process(cfg, i)
			}
			continue
		}

		if strings.Contains(line, target) {
			process(cfg, i)
		}
	}
}

func process(cfg *grepConfig, idx int) {
	start := idx
	end := idx

	if cfg.ctx > 0 {
		start -= cfg.ctx
		end += cfg.ctx
	}

	if cfg.before > 0 {
		start -= cfg.before
	}

	if cfg.after > 0 {
		end += cfg.after
	}

	if start < 0 {
		start = 0
	}

	if end >= len(cfg.lines) {
		end = len(cfg.lines) - 1
	}

	for i := start; i <= end; i++ {
		if cfg.lineNum {
			if i == idx {
				cfg.result[i] = fmt.Sprintf("%d:%s", i+1, cfg.lines[i])
				cfg.counter++
			} else {
				cfg.result[i] = fmt.Sprintf("%d-%s", i+1, cfg.lines[i])
			}
			continue
		}

		cfg.result[i] = cfg.lines[i]
		cfg.counter++
	}
}

func (cfg *grepConfig) printResult() {
	if cfg.count {
		fmt.Println(cfg.counter)
		return
	}

	if cfg.invert {
		for i, s := range cfg.lines {
			if _, ok := cfg.result[i]; !ok {
				fmt.Println(s)
			}
		}
		return
	}

	for _, line := range cfg.result {
		fmt.Println(line)
	}
}

func main() {
	after := flag.Int("A", 0, "Print +N lines after match")
	before := flag.Int("B", 0, "Print +N lines before match")
	ctx := flag.Int("C", 0, "Print ±N lines before and after match")
	count := flag.Bool("c", false, "Print number of matches")
	ignoreCase := flag.Bool("i", false, "Ignore case")
	invert := flag.Bool("v", false, "Invert matches")
	fixed := flag.Bool("F", false, "Exact match to line, not pattern")
	lineNum := flag.Bool("n", false, "Print number of line")

	flag.Parse()

	input := os.Args[len(os.Args)-1]
	inputFile, err := os.Open(input)
	if err != nil {
		log.Fatal("Failed to open input file:", err.Error())
	}

	defer inputFile.Close()

	var lines []string
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading input file:", err)
	}

	target := os.Args[len(os.Args)-2]

	cfg := &grepConfig{
		lines:      lines,
		after:      *after,
		before:     *before,
		ctx:        *ctx,
		count:      *count,
		ignoreCase: *ignoreCase,
		invert:     *invert,
		fixed:      *fixed,
		lineNum:    *lineNum,
		result:     make(map[int]string),
	}

	grep(cfg, target)
	cfg.printResult()
}
