// Package golexer implements a very basic lexer to lex any text file
// or text source into a list of tokens.

// The lexer can recognize text within string qoutes "" as special case and generates a single token
// for it.
package golexer

import (
	"bufio"
	"errors"
	"io"
	"log"
	"strings"
)

type Position struct {
	Line     int
	ColStart int
	ColEnd   int
}

type Tokenizer struct {
	seperators map[rune]bool
	current    int
}

type Token struct {
	Value string
	Pos   Position
}

var DefaultSeperators = map[rune]bool{
	'(': true,
	'}': true,
	')': true,
	'{': true,
	'[': true,
	']': true,
	',': true,
}

func NewTokenizer(seperators map[rune]bool) Tokenizer {
	return Tokenizer{
		seperators: seperators,
	}
}

// GetNextToken returns the next token from the passed line it works on rune list
// instead of string. Uses the state Tokenizer.current to track/update current index
// being parsed

// a map of seperators can be passed to parse these runes as single token if they are
// not part of a string, the return type is string and not package specific Token type
// hence this function can be used anywhere.
func (p *Tokenizer) GetNextToken(line []rune, seperators map[rune]bool) (string, error) {
	res := []rune{}
	stringToken := false
	found := false

	// strip spaces
	for line[p.current] == ' ' {
		p.current++
	}

	s := line[p.current]
	// if current character starts with " we have to handle it seperately since
	// the next for loop can't handle this case
	if s == '"' {
		res = append(res, s)
		p.current++
		stringToken = true
	}

	for i := p.current; i < len(line); i++ {
		if !stringToken {
			// seperators in this case refers to special characters that
			// have to be parsed as seperate token
			// TODO make it configurable
			if _, ok := seperators[line[i]]; ok {
				if len(res) > 0 {
					return string(res), nil
				}
				p.current++
				res = append(res, line[i])
				found = true
				break
			}

			if line[i] == ' ' {
				p.current++
				found = true
				break
			}
		}

		res = append(res, line[i])
		p.current++

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

// LexLine handles converting a string into array of tokens
func (p *Tokenizer) LexLine(line string, lc int) []Token {
	var err error

	program := []Token{}
	word := ""
	p.current = 0

	line = strings.TrimSpace(line)
	if len(line) <= 0 {
		return program
	}

	for p.current < len(line) {
		iCol := p.current
		// we have a special case for character starting with " and are considered strings
		word, err = p.GetNextToken([]rune(line), p.seperators)
		if err != nil {
			log.Panicf("Error at parsing line %d col %d: %s", lc, iCol+1, err)
		}

		token := Token{
			Pos: Position{
				Line:     lc,
				ColStart: iCol + 1,
				ColEnd:   p.current - iCol,
			},
			Value: word,
		}
		program = append(program, token)
	}

	return program
}

func (p *Tokenizer) Lex(f io.Reader) ([]Token, error) {
	var program []Token
	scanner := bufio.NewScanner(f)

	l := 1
	for scanner.Scan() {
		line := scanner.Text()
		tokens := p.LexLine(line, l)
		program = append(program, tokens...)
		l++
	}

	return program, nil
}
