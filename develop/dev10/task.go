package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

type TelnetOptions struct {
	Timeout time.Duration
}

func parseFlags() (*TelnetOptions, string) {
	opts := &TelnetOptions{}
	flag.DurationVar(&opts.Timeout, "timeout", 10*time.Second, "Connetion timeout")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: go-telnet [--timeout=10s] [HOST] [PORT]")
		os.Exit(1)
	}

	host := args[0]
	port := args[1]
	address := net.JoinHostPort(host, port)

	return opts, address
}

func openConnection(opts *TelnetOptions, address string) net.Conn {
	conn, err := net.DialTimeout("tcp", address, opts.Timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open connection: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func writeServer(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text() + "\n"
		if _, err := conn.Write([]byte(text)); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write to server: %v\n", err)
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read from stdin: %v\n", err)
	}

	conn.Close()
	os.Exit(0)
}

func readServer(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed by server")
			os.Exit(0)
		}
		fmt.Println(message)
	}
}

func main() {
	opts, address := parseFlags()
	conn := openConnection(opts, address)
	defer conn.Close()

	fmt.Printf("Connected to %s\n", address)

	go writeServer(conn)
	readServer(conn)
}
