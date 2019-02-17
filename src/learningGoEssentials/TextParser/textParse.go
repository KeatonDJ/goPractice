//Count how many times a word is inside of a phrase/textblock. I used the lyrics from "Misty Mountains" from The Hobbit written by Richard Billy Setford, Misty Mountains lyrics © Words & Music A Div Of Big Deal Music LLC

package main

import (
	"fmt"
	"strings"
)

func main() {

	phrase := `
Far over the misty mountains cold
To dungeons deep and caverns old
We must away ere break of day
To find our long-forgotten gold
The pines were roaring on the height
The winds were moaning in the night
The fire was red, it flaming spread
The trees like torches blazed with light`

	phrase = strings.ToLower(phrase)
	wordMap := map[string]int{}
	words := strings.Fields(phrase)

	var topWord string
	var count int

	for _, word := range words {

		wordMap[word]++

		if wordMap[word] > wordMap[topWord] {
			topWord, count = word, wordMap[word]
		}
	}

	fmt.Printf("Most used word is: %q and was used %d times\n", topWord, count)

	fmt.Println("Word usage:")

	for _, cur := range words {

		fmt.Printf("%q was used %d times.\n", cur, wordMap[cur])

	}

}
