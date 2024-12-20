package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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
// Опции утилиты
type GrepOptions struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
}

// Функция для чтения файла
func readLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file.", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// Фукнция для проверки строки на вхождение
func matchLine(pattern, line string, opt *GrepOptions) bool {
	if opt.IgnoreCase {
		pattern = strings.ToLower(pattern)
		line = strings.ToLower(line)
	}

	if opt.Fixed {
		return strings.Contains(line, pattern)
	}

	matched, err := regexp.MatchString(pattern, line)
	if err != nil {
		log.Fatal("Failed to match line.", err)
	}
	return matched
}

func parseFlags() (*GrepOptions, string, string) {
	opt := &GrepOptions{}

	flag.IntVar(&opt.After, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&opt.Before, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&opt.Context, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&opt.Count, "c", false, "количество строк")
	flag.BoolVar(&opt.IgnoreCase, "i", false, "игнорировать регистр")
	flag.BoolVar(&opt.Invert, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&opt.Fixed, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&opt.LineNum, "n", false, "напечатать номер строки")

	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: go-grep [...OPTIONS] [PATTERN] [FILENAME]")
		os.Exit(1)
	}
	pattern, filename := args[0], args[1]
	return opt, pattern, filename
}

func findLinesForPrint(result []int, opt *GrepOptions, lines []string) []int {
	var printResult []int
	for _, idx := range result {
		start := idx - opt.Before
		end := idx + opt.After + 1

		if start < 0 {
			start = 0
		} else if len(printResult) > 0 && printResult[len(printResult)-1] > start {
			start = printResult[len(printResult)-1] + 1
		}

		if end >= len(lines) {
			end = len(lines)
		}

		for i := start; i < end; i++ {
			printResult = append(printResult, i)
		}
	}
	return printResult
}

func findCorrectLines(lines []string, pattern string, opt *GrepOptions) ([]int, int) {
	var result []int
	var count int
	for i, line := range lines {
		match := matchLine(pattern, line, opt)

		if opt.Invert {
			match = !match
		}

		if match {
			count++
			result = append(result, i)
		}
	}
	return result, count
}

func main() {
	opt, pattern, filename := parseFlags()
	lines := readLines(filename)

	// Слайс индексов подходящих строк
	result, count := findCorrectLines(lines, pattern, opt)
	if opt.Count {
		fmt.Println(count)
		return
	}

	if opt.Context > 0 {
		opt.After = opt.Context
		opt.Before = opt.Context
	}

	// Слайс индексов строк, которые необходимо напечатать
	printLines := findLinesForPrint(result, opt, lines)
	for _, idx := range printLines {
		if opt.LineNum {
			fmt.Print(idx+1, ": ")
		}
		fmt.Println(lines[idx])
	}
}
