package artmatlang

import (
	//	"errors"
	"fmt"
	"math"
)

type syntaxTreeNode struct {
	oper  byte
	val   float64
	nodes []syntaxTreeNode
}

// validate verifies that the node is makes sense.
// It requires that the node have an acceptable number of sub nodes for it's operation
// eg can not have three nodes and try to use exp
func (astnode *syntaxTreeNode) validate() error {
	switch astnode.oper {
	case byte('+'):
		return validNodeCount(astnode.nodes, 2, false)
	case byte('-'):
		return validNodeCount(astnode.nodes, 2, false)
	case byte('*'):
		return validNodeCount(astnode.nodes, 2, false)
	case byte('/'):
		return validNodeCount(astnode.nodes, 2, true)
	case byte('^'):
		return validNodeCount(astnode.nodes, 2, true)
	default:
		return nil
	}
}

func validNodeCount(nodes []syntaxTreeNode, count int, exact bool) error {
	if exact {
		if len(nodes) != count {
			return fmt.Errorf("Invalid number of nodes: found %v, expected %v", len(nodes), count)
		}
	} else {
		if len(nodes) < count {
			return fmt.Errorf("Invalid number of nodes: found %v, expected at least %v", len(nodes), count)
		}
	}
	return nil

}

func (astnode *syntaxTreeNode) evaluate() {
	switch astnode.oper {
	case byte('+'):
		astnode.add()
	case byte('-'):
		astnode.subtract()
	case byte('*'):
		astnode.multiply()
	case byte('/'):
		astnode.divide()
	case byte('^'):
		astnode.exp()
	}
}

func (astnode *syntaxTreeNode) traverse() {
	for node := range astnode.nodes {
		if astnode.nodes[node].oper != byte('V') {
			astnode.nodes[node].traverse()
		}
	}
	astnode.evaluate()
}

func (astnode *syntaxTreeNode) add() {
	sum := 0.0
	for _, v := range astnode.nodes {
		sum += v.val
	}
	astnode.val = sum
}

func (astnode *syntaxTreeNode) subtract() {
	diff := astnode.nodes[0].val
	for i, v := range astnode.nodes {
		if i != 0 {
			diff -= v.val
		}
	}
	astnode.val = diff
}

func (astnode *syntaxTreeNode) multiply() {
	total := astnode.nodes[0].val
	for i, v := range astnode.nodes {
		if i != 0 {
			total *= v.val
		}
	}
	astnode.val = total
}

func (astnode *syntaxTreeNode) divide() {
	result := astnode.nodes[0].val / astnode.nodes[1].val
	astnode.val = result
}

func (astnode *syntaxTreeNode) exp() {
	result := math.Pow(astnode.nodes[0].val, astnode.nodes[1].val)
	astnode.val = result
}
