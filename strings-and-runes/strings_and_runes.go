package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"

	fmt.Println("len:", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("rune len:", utf8.RuneCountInString(s))

	for idx, r := range s {
		fmt.Printf("%d => %#U\n", idx, r)
	}

	for i, w := 0, 0; i < len(s); i += w {
		runeVal, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeVal, i)
		w = width

		examineRune(runeVal)
	}

}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("Got a tee")
	} else if r == 'ส' {
		fmt.Println("Got a so sua")
	}
}