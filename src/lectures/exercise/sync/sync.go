//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type Count struct {
	count int
	sync.Mutex
}

func countLetters(word string) int {
	counter := 0
	for i := 0; i < len(word); i++ {
		counter += 1
	}
	return counter
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var wg sync.WaitGroup
	count := Count{}

	words := strings.Split(scanner.Text(), " ")
	for _, word := range words {
		used_word := word
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Lock()
			defer count.Unlock()

			count.count += countLetters(used_word)
		}()
	}
	wg.Wait()
	fmt.Println("length of word = ", count.count)
}
