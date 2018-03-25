package artmatlang

import (
	"fmt"
)

// Take the lexed string as an input and break into ast
func Parse(lexedString []token) syntaxTreeNode {
	headNode := syntaxTreeNode{val: 0, oper: byte('H')}
	// create the subnodes
	return headNode
}

func createSubnode(substring []token) (syntaxTreeNode, error) {
	// if the substring is one token then that token should be a number,
	// verify that it is, and create the appropriate value node for it
	// else return an error
	if len(substring) == 1 {
		if substring[0].token == byte('\x00') {
			//			value := substring[0].lexeme
			value := 8.3
			return syntaxTreeNode{val: value, oper: byte('V'), nodes: nil}, nil
		} else {
			return syntaxTreeNode{}, fmt.Errorf("expected number token, received %s", substring[0].token)
		}
	}
	if hasSubstrings(substring) {
		findSubstrings(substring)
	}

	a := syntaxTreeNode{}
	b := syntaxTreeNode{}
	c := syntaxTreeNode{nodes: []syntaxTreeNode{a, b}}
	return c, nil
}

// parseNum takes a byte slice of nums and converts them to a number
func parseNum() {
	//	var value float64
	fmt.Println(int(byte(8)))
}

// initialise hasSubstrings as true,
// determine if the substring has further substrings
// if there are no paren tokens then this has further substrings
// if there are parens, set hasSubstrings to false
func hasSubstrings(substring []token) bool {

	hasSubstrings := false
	for _, token := range substring {
		if token.token == byte('(') {
			hasSubstrings = true
		}
		if token.token == byte(')') {
			hasSubstrings = true
		}
	}
	return hasSubstrings
}

// find first open paren, and coresponding close paren
// return the substring striped of it's open and close parens
func findSubstrings(substring []token) []token {
	openCount := 0
	openPos := 0
	closePos := 0
	for pos, token := range substring {
		if token.token == byte('(') {
			if openCount == 0 {
				openPos = pos + 1
			}
			openCount++
		} else if token.token == byte(')') {
			openCount--
			if openCount == 0 {
				closePos = pos
			}
		}
	}
	return substring[openPos:closePos]
}
