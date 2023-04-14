package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is prime number"},
		{"not prime", 4, false, "4 is not prime because is divisible by 2"},
		{"zero", 0, false, "0 is not prime"},
		{"one", 1, false, "1 is not prime"},
		{"negatif", -1, false, "negative is not prime"},
	}
	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}
		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}
		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}
func Test_prompt(t *testing.T) {

	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w
	myPrompt()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)
	//perform test
	if string(out) != "-> " {
		t.Errorf("Incorrect prompt: expected -> but got %s", string(out))
	}
}
func Test_intro(t *testing.T) {
	oldOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w
	intro()
	_ = w.Close()
	os.Stdout = oldOut

	out, _ := io.ReadAll(r)
	if !strings.Contains(string(out), "Enter whole number") {
		t.Errorf("Incorrect intro text, got %s", string(out))
	}
}
func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", "", "Enter whole number"},
		{"zero", "0", "0 is not prime"},
		{"one", "1", "1 is not prime"},
		{"two", "2", "2 is prime number"},
		{"four", "4", "4 is not prime because is divisible by 2"},
		{"negative", "-2", "negative is not prime"},
		{"decimal", "4.0", "Enter whole number"},
		{"quit", "q", ""},
	}
	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)
		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected, but got %s", e.expected, res)
		}
	}
}
func Test_readUserInput(t *testing.T) {

	doneChan := make(chan bool)

	var stdin bytes.Buffer
	stdin.Write([]byte("1\nq\n"))
	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
