package main

import (
	"strings"

	"github.com/brianvoe/gofakeit/v7"
)

func main() {

	name := gofakeit.Name()
	println(name)
	//range
	//String to []rune
	for _, value := range name {
		println(string(value))
	}
	//length
	println("Name Length", len(name))
	//uppercase
	println("Uppercase:", strings.ToUpper(name))
	//lowercase
	println("Lowercase:", strings.ToLower(name))
	//contains
	println("Contains 'a':", strings.Contains(name, "a"))
	//replace
	println("Replace a with o:", strings.ReplaceAll(name, "a", "o"))
	//split
	parts := strings.Split(name, " ")
	for i, part := range parts {
		println("Part", i, ":", part)
	}
	//join
	joined := strings.Join(parts, "-")
	println("Joined with - :", joined)
	//trim
	trimmed := strings.TrimSpace("   " + name + "   ")
	println("Trimmed:", trimmed)
	//repeat
	println("Repeated 3 times:", strings.Repeat(name+" ", 3))
	//index
	println("Index of 'a':", strings.Index(name, "a"))
	//last index
	println("Last Index of 'a':", strings.LastIndex(name, "a"))
	//substring
	if len(name) > 5 {
		println("Substring (0,5):", name[0:5])
	}

}
