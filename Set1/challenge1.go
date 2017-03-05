package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

/*
Convert hex to base64
The string:

49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
Should produce:

SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
So go ahead and make that happen. You'll need to use this code for the rest of the exercises.

Cryptopals Rule
Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.
*/

func challenge1() {
	var hexString = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	var b64String = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	byteString, _ := hex.DecodeString(hexString) //hex to bytes
	fmt.Println(string(byteString))
	byteToBase64 := base64.StdEncoding.EncodeToString(byteString) //bytes to b64
	fmt.Println(byteToBase64)
	if b64String == byteToBase64 {
		fmt.Println("Challenge 1 Pass")
	}

}