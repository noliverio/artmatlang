package artmatlang

import (
	"fmt"
	"reflect"
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

func TestHasSubstrings(t *testing.T) {

	initialString := []token{
		token{token: byte('\x00'), lexeme: []byte{byte('5')}},
		token{token: byte('+')},
		token{token: byte('(')},
		token{token: byte('\x00'), lexeme: []byte{byte('4')}},
		token{token: byte('+')},
		token{token: byte('\x00'), lexeme: []byte{byte('2')}},
		token{token: byte(')')},
	}

	if !hasSubstrings(initialString) {
		t.Fail()
	}
}

func TestFindSubstrings(t *testing.T) {
	// pass a string of tokens for 5 + (4 + 2 ) to findSubstring
	initialString := []token{
		token{token: byte('\x00'), lexeme: []byte{byte('5')}},
		token{token: byte('+')},
		token{token: byte('(')},
		token{token: byte('\x00'), lexeme: []byte{byte('4')}},
		token{token: byte('+')},
		token{token: byte('\x00'), lexeme: []byte{byte('2')}},
		token{token: byte(')')},
	}
	expectedSubstring := []token{
		token{token: byte('\x00'), lexeme: []byte{byte('4')}},
		token{token: byte('+')},
		token{token: byte('\x00'), lexeme: []byte{byte('2')}},
	}
	substring := findSubstrings(initialString)
	if !reflect.DeepEqual(substring, expectedSubstring) {
		t.Fail()
	}
}

func TestCreateValueNode(t *testing.T) {
	initialToken := token{token: byte('\x00'), lexeme: []byte{byte('5')}}
	returnNode, err := createValueNode(initialToken)
	panicOnError(err)
	expectedNode := syntaxTreeNode{oper: byte('V'), val: 8.3}
	if !reflect.DeepEqual(returnNode, expectedNode) {
		t.Fail()
	}
}
