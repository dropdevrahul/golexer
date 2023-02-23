package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dropdevrahul/golexer/golexer"
)

func main() {
	usage := `Usage: golexer FILE`
	//save := flag.Bool("c", false, "Save results to a file instead of printing")

	if len(os.Args) < 2 {
		fmt.Println(usage)
		return
	}

	tz := golexer.NewTokenizer()

	tokens, err := tz.LexFile(os.Args[1])
	if err != nil {
		log.Panic(err)
	}
	for _, t := range tokens {
		fmt.Printf("%s\n", t.Value)
	}
	fmt.Println()
}
