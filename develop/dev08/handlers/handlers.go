package handlers

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/mitchellh/go-ps"
)

type HandlerFunc func(args []string) (output string, err error)

func HandleExit(args []string) (string, error) {
	os.Exit(0)
	return "", nil
}

func HandlePwd(args []string) (string, error) {
	return GetCurrentDirectory()
}

func HandleCd(args []string) (string, error) {
	if len(args) != 2 {
		return "", errors.New("Usage: cd [DIRECTORY]")
	}
	err := os.Chdir(args[1])
	if err != nil {
		return "", errors.New(fmt.Sprintf("Can't find directory: %s\n", args[1]))
	}
	return "", nil
}

func HandlePs(args []string) (string, error) {
	processes, err := ps.Processes()
	if err != nil {
		return "", err
	}
	result := make([]string, 0, len(processes))
	for _, process := range processes {
		p := fmt.Sprintf("%d: %s", process.Pid(), process.Executable())
		result = append(result, p)
	}
	return strings.Join(result, "\n"), nil
}

func HandleKill(args []string) (string, error) {
	if len(args) != 2 {
		return "", errors.New("Usage: kill [PID]")
	}
	pid, err := strconv.Atoi(args[1])
	if err != nil {
		return "", errors.New("Invalid PID: " + args[1])
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return "", err
	}
	process.Kill()
	return "", nil
}

func HandleNetcat(args []string) (string, error) {
	if len(args) != 3 {
		fmt.Println("Usage: netcat [HOST] [PORT]")
	}

	host, port := args[1], args[2]
	address := net.JoinHostPort(host, port)

	timeout := 10 * time.Second
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open connection: %v\n", err)
	}

	fmt.Printf("Connected to %s\n", address)

	// Writer to server
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			message := scanner.Text() + "\n"
			if _, err := conn.Write([]byte(message)); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to send message: %v\n", err)
				conn.Close()
				return
			}
		}
	}()

	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Connection closed")
				return
			}
			fmt.Println(message)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	conn.Close()
	return "", nil
}

func GetCurrentDirectory() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}
