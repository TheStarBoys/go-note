package simpleScript

import (
	"testing"
	"fmt"
)

func TestSimpleLexer(t *testing.T) {
	reader := Tokenize("int age= ")
	outputTokens(reader)
}

func outputTokens(reader TokenReader) {
	for reader.IsEnd() == false {
		token := reader.Read()
		fmt.Printf("type: %d, text: %s\n", token.GetType(), string(token.GetText()))
	}
}