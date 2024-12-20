package main

import "testing"

func TestEmptyString(t *testing.T) {
	testCorrectInput(t, "", "")
}

func TestCorrectStringWithoutDigits(t *testing.T) {
	testCorrectInput(t, "abcd", "abcd")
}

func TestCorrectStringWithDigits(t *testing.T) {
	testCorrectInput(t, "a3bc4d", "aaabccccd")
}

func TestCorrectStringWithEscapedRunes(t *testing.T) {
	testCorrectInput(t, "\\\\3bc\\52c2", "\\\\\\bc55cc")
}

func TestCorrectStringWithNonAscii(t *testing.T) {
	testCorrectInput(t, "в2б3", "ввббб")
}

func TestTwoUnescapedDigits(t *testing.T) {
	input := "abc23"
	result, err := Unpack(input)

	if result != "" {
		t.Errorf("wrong result. want=%s, got=%s", "", result)
	}

	errorMessage := "invalid input: two unescaped digits in a row"
	if err.Error() != errorMessage {
		t.Errorf("wrong error. want=%s, got=%s", errorMessage, err.Error())
	}
}

func TestSlashInEnd(t *testing.T) {
	input := "abc2\\"
	result, err := Unpack(input)

	if result != "" {
		t.Errorf("wrong result. want=%s, got=%s", "", result)
	}

	errorMessage := "invalid input: string not terminated"
	if err.Error() != errorMessage {
		t.Errorf("wrong error. want=%s, got=%s", errorMessage, err.Error())
	}
}

func testCorrectInput(t *testing.T, input, expected string) {
	result, err := Unpack(input)

	if err != nil {
		t.Fatalf("unexpected error: %v\n", err)
	}

	if result != expected {
		t.Errorf("wrong result. got=%s, want=%s", result, expected)
	}
}
