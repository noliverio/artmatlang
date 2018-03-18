package artmatlang

// Take the lexed string as an input and break into ast
func Parse(lexedString []token) syntaxTreeNode {
	headNode := syntaxTreeNode{val: 0, oper: byte('H')}
	headNode.nodes, _ = createSubnode(lexedString)
	return headNode
}

func createSubnode(substring []token) ([]syntaxTreeNode, error) {
	// find first paren going from r to l
	// find lastparen going from l to r
	a := syntaxTreeNode{}
	b := syntaxTreeNode{}
	c := []syntaxTreeNode{a, b}
	return c, nil
}
