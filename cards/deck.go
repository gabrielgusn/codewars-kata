package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) newDeck() deck {

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

func (d deck) printDeck() {
	for pos, card := range d {
		fmt.Println(pos, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveDeckToFile(filename string) error {
	err := os.WriteFile(filename, []byte(d.toString()), 0777)
	return err
}

func newDeckFromFile(filepath string) deck {
	newDeck := deck{}

	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(500)
	}

	newDeck = strings.Split(string(fileContent), ",")

	return newDeck

}

// func (d deck) shuffle() deck {
// 	var randPosArr []int
// 	var shuffledDeck []string = make([]string, len(d))
// 	for i := 0; i <= len(d); i++ {
// 		// randPosArr = append(randPosArr, rand.Intn((len(d))))
// 		rand_num = rand.Intn(len(d))

// 		if shuffledDeck[rand_num] != "" {
// 			positioned := false
// 			j := i + 1
// 			for !positioned {
// 				if j >= len(d)-1 {
// 					j = 0
// 				}
// 				if shuffledDeck[j] == "" {
// 					shuffledDeck[j] =
// 				}

// 				j++
// 			}
// 		} else {
// 			shuffledDeck[rand_num] =
// 		}
// 	}

// 	return shuffledDeck
// }

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	var auxCard string
	var randomNumber int
	for pos, card := range d {
		randomNumber = r.Intn(len(d))
		auxCard = d[randomNumber]
		d[randomNumber] = card
		d[pos] = auxCard
	}
}
