package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// This problem is a bruteforce search probem

func FindAdventCoin(secretKey string) int {
	count := 1 // we'll start with one as the lowest number.

	for {
		// concatenate the input string and the count.
		data := fmt.Sprintf("%s%d", secretKey, count)

		// generate the hash for it.

		hash := md5.Sum([]byte(data))

		// encode it to a string.

		hashStr := hex.EncodeToString(hash[:])

		// if the condition is met then return the count otherwise keep increasing it.

		if hashStr[:5] == "00000" && hashStr[:6] == "000000" {
			return count
		}

		count++
	}
}

func main() {

	// word := "bgvyzdsv"

	// word2 := "vsdzyvgb"

	word3 := "dgthioke"

	// fmt.Println(FindAdventCoin(word))
	fmt.Println(FindAdventCoin(word3))
	// fmt.Println(FindAdventCoin(word))
}
