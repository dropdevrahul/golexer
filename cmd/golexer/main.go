package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dropdevrahul/golexer"
)

func main() {
	//usage := `Usage: golexer FILE`
	//save := flag.Bool("c", false, "Save results to a file instead of printing")
	var err error

	input := os.Stdin

	if len(os.Args) == 2 {
		input, err = os.Open(os.Args[0])
		if err != nil {
			panic(err)
		}
	}

	tz := golexer.NewTokenizer(golexer.DefaultSeperators)

	tokens, err := tz.Lex(input)
	if err != nil {
		log.Panic(err)
	}

	for _, t := range tokens {
		fmt.Printf("%s\n", t.Value)
	}

	fmt.Println()
}
