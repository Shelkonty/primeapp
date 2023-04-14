package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	intro()

	doneChan := make(chan bool)

	go readUserInput(os.Stdin, doneChan)

	<-doneChan

	close(doneChan)
	fmt.Println("Bye...")
}
func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)
	for {
		res, done := checkNumbers(scanner)
		if done {
			doneChan <- true
			return
		}
		fmt.Println(res)
	}
}
func checkNumbers(scanner *bufio.Scanner) (string, bool) {

	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}
	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Enter whole number", false
	}
	_, msg := isPrime(numToCheck)
	return msg, false
}
func intro() {
	fmt.Println("Prime Number")
	fmt.Println("------------")
	fmt.Println("Enter whole number (q for quit).")
	myPrompt()
}
func myPrompt() {
	fmt.Print("-> ")
}
func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime", n)
	}
	if n < 0 {
		return false, "negative is not prime"
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not prime because is divisible by %d", n, i)
		}
	}
	return true, fmt.Sprintf("%d is prime number", n)
}
