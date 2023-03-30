package golexer_test

import (
	"os"
	"testing"

	"github.com/dropdevrahul/golexer"
	"github.com/stretchr/testify/assert"
)

func TestLexFile(t *testing.T) {
	tz := golexer.NewTokenizer(golexer.DefaultSeperators)

	input, err := os.Open("test-data/sample.txt")
	if err != nil {
		panic(err)
	}

	tokens, err := tz.Lex(input)
	assert.Equal(t, err, nil)
	assert.Equal(t, 11, len(tokens))
}
