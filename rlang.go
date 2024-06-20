package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
)

func ClassicLangs() []string {
	return []string{
		"C",
		"C++",
		"C#",
		"Javascript",
		"Typescript",
		"Python",
		"Java",
		"Kotlin",
		"Go",
		"Rust",
	}
}

func FunctionalLangs() []string {
	return []string{
		"Clojure",
		"Elixir",
		"Elm",
		"F#",
		"Haskell",
		"Idris",
		"OCaml",
		"PureScript",
		"Racket",
		"Scheme",
		"Scala",
	}
}

func MiscLangs() []string {
	return []string{
		"Basic",
		"Bash",
		"Crystal",
		"Dart",
		"Fortran",
		"Lua",
		"Pascal",
		"Perl/Raku",
	}
}

func ArrayLangs() []string {
	return []string{
		"APL",
		"BQN",
		"UIUA",
		"J",
	}
}

func SelectLang(langList []string) string {
	var randomN = rand.IntN(len(langList))
	return langList[randomN]
}

func main() {
	var buffer []string

	classic := flag.Bool("classic", true, "Allows classic programming languages like C, C++, Python, Java, ...")
	functional := flag.Bool("functional", false, "Allows functional programming languages like Clojure, Elixir, Haskell, OCaml, ...")
	misc := flag.Bool("misc", false, "Allows miscellaneous programming languages like Crystal, Fortran, Bash, Lua, ...")
	array := flag.Bool("array", false, "Allows array programming languages like APL, BQN, UIUA and J")

	flag.Parse()

	if *classic {
		for _, v := range ClassicLangs() {
			buffer = append(buffer, v)
		}
	}
	if *functional {
		for _, v := range FunctionalLangs() {
			buffer = append(buffer, v)
		}
	}
	if *misc {
		for _, v := range MiscLangs() {
			buffer = append(buffer, v)
		}
	}
	if *array {
		for _, v := range ArrayLangs() {
			buffer = append(buffer, v)
		}
	}
	fmt.Println(SelectLang(buffer))
}
