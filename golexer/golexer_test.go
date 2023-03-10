package golexer_test

import (
	"testing"

	"github.com/dropdevrahul/lexer/golexer"
	"github.com/stretchr/testify/assert"
)

func TestLexFile(t *testing.T) {
	tz := golexer.NewTokenizer()

	tokens, err := tz.LexFile("sample.txt")
	assert.Equal(t, err, nil)
	assert.Equal(t, 11, len(tokens))
}
