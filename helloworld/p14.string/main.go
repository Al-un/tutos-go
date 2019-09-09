package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	name := "Hello world"
	fmt.Println(name)
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}
	fmt.Printf("\n")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c ", name[i])
	}

	name = "Señor"
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}
	fmt.Printf("\n")

	nameRunes := []rune(name)
	for i := 0; i < len(nameRunes); i++ {
		fmt.Printf("%c ", nameRunes[i])
	}
	fmt.Printf("\n")
	for index, rune := range name {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str = string(runeSlice)
	fmt.Println(str)
	fmt.Printf("length of %s is %d\n", str, utf8.RuneCountInString(str))

	// mutation
	toto := "toto"
	fmt.Println(mutateStr([]rune(toto)))
}

func mutateStr(s []rune) string {
	s[0] = 'a'
	return string(s)
}
