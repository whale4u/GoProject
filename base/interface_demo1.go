package main

import (
	"fmt"
)

//interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

//我们给接受者类型（Receiver Type） MyString 添加了方法 FindVowels() []rune。
// 现在，我们称 MyString 实现了 VowelsFinder 接口。
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := MyString("Niko Buck")
	var v VowelsFinder
	// v 的类型为 VowelsFinder，name 的类型为 MyString，我们把 name 赋值给了 v。
	// 由于 MyString 实现了 VowelFinder，因此这是合法的。
	v = name
	fmt.Printf("Vowels are %c", v.FindVowels())
}
