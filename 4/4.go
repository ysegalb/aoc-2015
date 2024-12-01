package main

import (
	"crypto/md5"
	"strconv"
)

/*
--- Day 4: The Ideal Stocking Stuffer ---
Santa needs help mining some AdventCoins (very similar to bitcoins) to use as gifts for all the economically forward-thinking little girls and boys.

To do this, he needs to find MD5 hashes which, in hexadecimal, start with at least five zeroes. The input to the MD5 hash is some secret key (your puzzle input, given below) followed by a number in decimal. To mine AdventCoins, you must find Santa the lowest positive number (no leading zeroes: 1, 2, 3, ...) that produces such a hash.

For example:

If your secret key is abcdef, the answer is 609043, because the MD5 hash of abcdef609043 starts with five zeroes (000001dbbfa...), and it is the lowest such number to do so.
If your secret key is pqrstuv, the lowest number it combines with to make an MD5 hash starting with five zeroes is 1048970; that is, the MD5 hash of pqrstuv1048970 looks like 000006136ef....

*/

func MineCoins(input string) (int, int) {
	hashFiveZeroes := 0
	hashSixZeroes := 0
	i := 0
	for {
		data := []byte(input + strconv.Itoa(i))
		result := md5.Sum(data)
		if hashFiveZeroes == 0 && FiveZeroes(result) {
			hashFiveZeroes = i
		}
		if hashSixZeroes == 0 && SixZeroes(result) {
			hashSixZeroes = i
		}
		if hashFiveZeroes != 0 && hashSixZeroes != 0 {
			return hashFiveZeroes, hashSixZeroes
		}
		i++
	}
}

func FiveZeroes(r [16]byte) bool {
	// Checks first two bites for 0 bitwise. And the third byte using bitwise with 0xF0 implies only the most significant nibble is compared.
	// Returns true only if all of them are 0.
	return r[0]|r[1]|(r[2]&0xF0) == 0 // If below 0xF0, it is 0
}

func SixZeroes(r [16]byte) bool {
	// Checks first three bites for 0 bitwise. That will be true if all of them are 0.
	return r[0]|r[1]|r[2] == 0
}
