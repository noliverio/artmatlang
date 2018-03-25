package artmatlang

import (
	"fmt"
	"testing"
)

func TestPrintLexStart(t *testing.T) {
	fmt.Println("Starting Lex Tests")
}

func TestLex(t *testing.T) {
	testEq1 := []byte("( 52 + ( 3 * 2 ) )")

	testEq1Lexed := lex(testEq1)

	fmt.Println(testEq1Lexed)
}
