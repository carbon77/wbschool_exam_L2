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

// CutOptions содержит параметры для обработки файла
type CutOptions struct {
	Fields    []int  // Номера полей для вывода (нумерация с 0)
	Delimiter string // Разделитель полей
	Separated bool   // Выводить только строки с разделителем
}

// readLines читает все строки из файла и возвращает их в виде слайса
func readLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

// parseFlags разбирает аргументы командной строки и возвращает настройки и имя файла
func parseFlags() (*CutOptions, string) {
	fields := flag.String("f", "", "выберите поля (колонки)")
	delimiter := flag.String("d", "\t", "использовать РАЗДЕЛИТЕЛЬ вместо TAB для разделения полей")
	separated := flag.Bool("s", false, "выводить только строки с разделителем")
	flag.Parse()

	if *fields == "" {
		log.Fatal("флаг fields (-f) обязателен")
	}

	var fieldsList []int
	for _, f := range strings.Split(*fields, ",") {
		field, err := strconv.Atoi(f)
		if err != nil {
			log.Fatal(err)
		}
		fieldsList = append(fieldsList, field-1)
	}

	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("необходимо указать имя файла")
	}
	filename := args[0]

	return &CutOptions{
		Fields:    fieldsList,
		Delimiter: *delimiter,
		Separated: *separated,
	}, filename
}

// cutLines обрабатывает строки согласно заданным параметрам
func cutLines(lines []string, opt *CutOptions) []string {
	var result []string
	for _, line := range lines {
		// Пропускаем строки без разделителя, если установлен флаг separated
		if opt.Separated && !strings.Contains(line, opt.Delimiter) {
			continue
		}

		var newLine []string
		fields := strings.Split(line, opt.Delimiter)

		// Выбираем нужные поля
		for _, field := range opt.Fields {
			if len(fields) < field {
				newLine = append(newLine, "")
			} else {
				newLine = append(newLine, fields[field])
			}
		}

		result = append(result, strings.Join(newLine, opt.Delimiter))
	}
	return result
}

func main() {
	opt, filename := parseFlags()
	lines := readLines(filename)
	cutLines(lines, opt)
	for _, line := range cutLines(lines, opt) {
		fmt.Println(line)
	}
}
