package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"ru/zakat/shell/handlers"
	"runtime"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nHello, %s!\n", user.Username)

	in := bufio.NewReader(os.Stdin)
	handlerFuncs := map[string]handlers.HandlerFunc{
		"exit":   handlers.HandleExit,
		"pwd":    handlers.HandlePwd,
		"cd":     handlers.HandleCd,
		"ps":     handlers.HandlePs,
		"kill":   handlers.HandleKill,
		"netcat": handlers.HandleNetcat,
		"echo": func(args []string) (string, error) {
			if len(args) != 2 {
				return "", errors.New("Usage: echo [MESSAGE]")
			}
			return args[1], nil
		},
	}

	for {
		fmt.Printf("\nmyshell>")

		input, err := in.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read string: %v\n", err)
		}
		input = strings.TrimSpace(input)
		// Разбиваем команду по | если это пайп
		commands := strings.Split(input, "|")
		var output string
		for _, command := range commands {
			args := parseCommand(command)
			if output != "" {
				args = append(args, fmt.Sprintf("\"%s\"", output))
			}

			handler, ok := handlerFuncs[args[0]]
			if !ok {
				output, err = executeCommand(args)
			} else {
				output, err = handler(args)
			}

			if err != nil {
				fmt.Println(err)
				break
			}
		}
		fmt.Println(output)
	}
}

// Функция для парсинга команды с аргументами в кавычках
func parseCommand(text string) []string {
	runes := []rune(strings.TrimSpace(text))
	command := make([][]rune, 1)
	curIdx := 0

	for curIdx < len(runes) {
		switch runes[curIdx] {
		case '"':
			curIdx++
			for runes[curIdx] != '"' {
				addToLastArg(command, runes[curIdx])
				curIdx++
			}
		case ' ':
			command = append(command, make([]rune, 0))
		default:
			addToLastArg(command, runes[curIdx])
		}
		curIdx++
	}

	result := make([]string, 0, len(command))
	for _, arg := range command {
		result = append(result, string(arg))
	}
	return result
}

// Добавить символ в конец последнего аргумента команды
func addToLastArg(command [][]rune, char rune) {
	command[len(command)-1] = append(command[len(command)-1], char)
}

// Фукнция для выполнения fork/exec команд
func executeCommand(args []string) (string, error) {
	var cmd *exec.Cmd
	// Если программа запущена в Windows необходимо вызвать программу через cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", append([]string{"/C"}, args...)...)
	} else {
		cmd = exec.Command(args[0], args[1:]...)
	}
	result := bytes.NewBufferString("")
	errResult := bytes.NewBufferString("")

	// Записываем результат в строку и возвращаем
	cmd.Stdout = result
	cmd.Stdin = os.Stdin
	cmd.Stderr = errResult
	if err := cmd.Run(); err != nil {
		fmt.Printf("%s: %v\n", args[0], err)
	}
	return result.String(), errors.New(errResult.String())
}
