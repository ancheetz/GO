package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type of 'deck, made of a slice of strings

type deck []string

//create a new function (method) that is called on type deck
func newDeck() deck {
	//create and lists a deck of cards
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}

	}
	return cards
}

//needs to list the receiver/ variable to print out, d --> function print is only callable on receiver "deck"
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//create a deal function on type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//convert a slice of deck into a string. this function only applies to d - the specific instance of type deck
func (d deck) toString() string {
	//strings is a package with function Join, which takes 2 arguments; we are joining the multiple strings w a comma to build one string
	return strings.Join([]string(d), ",")
}

//create a func to save a list of cards to a file on my local machine
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//create new deck from a file on harddrive

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//if an error occurs, we can either log the error and return a call to newDeck()
		//or log the error and completely exit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//converting the bs into a string
	s := strings.Split(string(bs), ",")
	return deck(s)
}

//shuffle all cards within the deck, by looping through the index, generating a random number between 0 and length of cards-1
//swap the current card and the card at cards[randomNum]

func (d deck) shuffle() {
	//create new random source, seed value
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	//calling upon r of type rand we can replace rand.Intn() with r.Intn() to generate a random source seed for the Intn()
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
