package golexer

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

type Position struct {
	Line     int
	ColStart int
	ColEnd   int
}

type Tokenizer struct {
	Current int
}

type Token struct {
	Value string
	Pos   Position
}

var seperators = map[rune]bool{
	'(': true,
	'}': true,
	')': true,
	'{': true,
	'[': true,
	']': true,
	',': true,
}

func NewTokenizer() Tokenizer {
	return Tokenizer{}
}

func (p *Tokenizer) GetNextToken(line []rune) (string, error) {
	res := []rune{}
	stringToken := false
	// strip spaces
	for line[p.Current] == ' ' {
		p.Current++
	}

	s := line[p.Current]
	if s == '"' {
		res = append(res, s)
		p.Current++
		stringToken = true
	}
	found := false

	for i := p.Current; i < len(line); i++ {
		if !stringToken {
			if _, ok := seperators[line[i]]; ok {
				if len(res) > 0 {
					return string(res), nil
				}
				p.Current++
				res = append(res, line[i])
				found = true
				break
			}

			if line[i] == ' ' {
				p.Current++
				found = true
				break
			}
		}

		res = append(res, line[i])
		p.Current++

		if stringToken && line[i] == s {
			found = true
			break
		}
	}

	// if currently parsing string token
	if !found && stringToken {
		return string(res), errors.New("Invalid string token")
	}

	return string(res), nil
}

func (p *Tokenizer) LexLine(line []rune, lineNo int) []Token {
	program := []Token{}
	if len(line) <= 0 {
		return program
	}
	word := ""
	var err error
	p.Current = 0
	for p.Current < len(line) {
		iCol := p.Current
		// we have a special case for character starting with " and are considered strings
		word, err = p.GetNextToken(line)
		if err != nil {
			log.Panicf("Error at parsing line %d col %d: %s", lineNo, iCol+1, err)
		}

		token := Token{
			Pos: Position{
				Line:     lineNo,
				ColStart: iCol + 1,
				ColEnd:   p.Current - iCol,
			},
			Value: word,
		}
		program = append(program, token)
	}

	return program
}

func (p *Tokenizer) LexFile(fpath string) ([]Token, error) {
	program := []Token{}
	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	l := 1

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		p := p.LexLine([]rune(line), l)

		program = append(program, p...)

		l++
	}

	return program, nil
}
