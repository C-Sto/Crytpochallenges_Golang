package main

import "strings"

/*
Single-byte XOR cipher
The hex encoded string:

1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
... has been XOR'd against a single character. Find the key, decrypt the message.

You can do this by hand. But don't: write code to do it for you.

How? Devise some method for "scoring" a piece of English plaintext. Character frequency is a good metric. Evaluate each output and choose the one with the best score.

Achievement Unlocked
You now have our permission to make "ETAOIN SHRDLU" jokes on Twitter.
*/

func scorePlaintext(candidate string) float32 {

	//freqs:
	/*
	  commits := map[string]int{
	    "rsc": 3711,
	    "r":   2138,
	    "gri": 1908,
	    "adg": 912,
	  }
	*/
	freq := map[string]float32{
		"a": 8.167, "b": 1.492, "c": 2.782, "d": 4.253, "e": 12.702,
		"f": 2.228, "g": 2.015, "h": 6.094, "i": 6.966, "j": 0.153,
		"k": 0.772, "l": 4.025, "m": 2.406, "n": 6.749, "o": 7.507,
		"p": 1.929, "q": 0.095, "r": 5.987, "s": 6.327, "t": 9.056,
		"u": 2.758, "v": 0.978, "w": 2.360, "x": 0.150, "y": 1.974,
		"z": 0.074, " ": 55.00, "\n": 0.01, ".": 6.53, ",": 6.16,
		";": 0.32, ":": 0.34, "!": 0.33, "?": 0.56, "'": 2.43,
		"-": 1.53,
	}

	//takes a string input, and gives a score based on how english it looks.
	//This function may be updated for later challenges depending on how well it works

	//for each letter of the english alphabet, do a chi-square test
	alphabet := "abcdefghijklmnopqrstuvwxyz \n.,;:!?'-'"
	score := float32(0)
	for _, letter := range alphabet {
		count := float32(strings.Count(candidate, string(letter)))
		expected := freq[string(letter)]
		score += ((count - expected) * (count - expected)) / expected
	}

	return score
}