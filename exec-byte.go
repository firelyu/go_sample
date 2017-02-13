package main

import (
	"fmt"
	"unicode/utf8"
)

func sample1() {
	const sampleString = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	//var sampleString []byte = []byte("\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")

	fmt.Println("Println:")
	fmt.Println(sampleString)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sampleString); i++ {
		fmt.Printf("%x ", sampleString[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sampleString)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sampleString)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sampleString)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sampleString)

	fmt.Println("%q loop:")
	for i := 0; i < len(sampleString); i++ {
		fmt.Printf("%q ", sampleString[i])
	}
	fmt.Printf("\n")

	fmt.Println("%+q loop:")
	for i := 0; i < len(sampleString); i++ {
		fmt.Printf("%+q ", sampleString[i])
	}
	fmt.Printf("\n")

	fmt.Println("Println []byte(string)")
	fmt.Println([]byte(sampleString))
}

func sample2() {
	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")
}

func sample3()  {
	const nihongo = "日本語"

	fmt.Println("The len is", len(nihongo))

	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}

	character := rune(0x00e0)
	fmt.Printf("%#U", character)
	fmt.Println()

	sentence := []rune{0x5ed6, 0x715c}
	fmt.Printf("%q", sentence)
}


// https://blog.golang.org/strings
func main() {
	sample2()
}
