package main

import "testing"

func TestParseSingleCommand(t *testing.T) {
	command := "ps"
	args := parseCommand(command)

	if len(args) != 1 {
		t.Errorf("wrong len(args). want=1, got=%d\n", len(args))
	}

	if args[0] != "ps" {
		t.Errorf("wrong args[0]. want=ps, got=%s\n", args[0])
	}
}

func TestParseCommandWithArg(t *testing.T) {
	command := "cd handlers"
	args := parseCommand(command)

	if len(args) != 2 {
		t.Errorf("wrong len(args). want=2, got=%d\n", len(args))
	}

	if args[0] != "cd" {
		t.Errorf("wrong args[0]. want=cd, got=%s\n", args[0])
	}

	if args[1] != "handlers" {
		t.Errorf("wrong args[1]. want=handlers, got=%s\n", args[1])
	}
}

func TestParseCommandWithString(t *testing.T) {
	command := "echo \"Hello world\""
	args := parseCommand(command)

	if len(args) != 2 {
		t.Errorf("wrong len(args). want=2, got=%d\n", len(args))
	}

	if args[0] != "echo" {
		t.Errorf("wrong args[0]. want=echo, got=%s\n", args[0])
	}

	if args[1] != "Hello world" {
		t.Errorf("wrong args[1]. want=Hello world, got=%s\n", args[1])
	}
}
