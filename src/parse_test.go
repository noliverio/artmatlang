package artmatlang

import (
	"fmt"
	"testing"
)

func TestPrintParseStart(t *testing.T) {
	fmt.Println("Starting Parse Tests")
}

func testCreateSubnode(t *testing.T) {
	// pass in the lexed subtring for 5 + 6
	substring := []token{token{token: byte('\x00'), lexeme: []byte("5")}, token{token: byte('+')}, token{token: byte('\x00'), lexeme: []byte("6")}}
	node, err := createSubnode(substring)
	fmt.Println(err)
	fmt.Println(node)

}

func TestFindSubstrings(t *testing.T) {
	// pass a string of tokens for 5 + (4 + 2 ) to findSubstring
	initialString := []token{token{token: byte('\x00')}, token{token: byte('+')}, token{token: byte('(')}, token{token: byte('\x00')}, token{token: byte('+')}, token{token: byte('\x00')}, token{token: byte(')')}}
	_ = findSubstrings(initialString)
}
