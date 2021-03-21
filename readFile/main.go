package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//program to open a file, scan the words, count them and print the word count.
	wordCount()
	charCount()
	countLetters()
	howManyOfEach()
}

func wordCount() {

	file, err := os.Open("text.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		scanner.Text()
		count++
	}
	if err != nil {
		panic(err)
	}

	//char := len()
	fmt.Printf("There are %d words in the file\n", count)

}

func charCount() {

	file, err := ioutil.ReadFile("text.txt")
	if err != nil {
		panic(err)
	}
	contents := string(file)
	words := bufio.NewScanner(strings.NewReader(contents))
	words.Split(bufio.ScanRunes)

	for words.Scan() {
		fmt.Print(words.Text())
	}
}

func countLetters() {

	file, err := ioutil.ReadFile("text.txt")
	if err != nil {
		panic(err)
	}

	wordsInFile := string(file)
	characters := len([]rune(wordsInFile))
	fmt.Printf("\nThere are %d characters in this file.\n", characters)
}

func howManyOfEach() {

	file, err := ioutil.ReadFile("text.txt")
	if err != nil {
		panic(err)
	}

	words := string(file)
	//Fields counts the number of spaces between fields in a string
	numWords := len(strings.Fields(words))
	countCharacters := strings.Count(words, "o")
	fmt.Printf("In my file, I have %d words, and %d occurences of %s", numWords, countCharacters, "\"o\". ")
}
