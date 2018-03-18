package artmatlang

import (
	"fmt"
	"testing"
)

func TestValidateAddPass(t *testing.T) {

	a := syntaxTreeNode{oper: 0, val: 5, nodes: nil}
	b := syntaxTreeNode{oper: 0, val: 2, nodes: nil}
	c := syntaxTreeNode{oper: byte('+'), val: 0, nodes: []syntaxTreeNode{a, b}}

	err := c.validate()
	if err != nil {
		t.Fail()
	}
}

func TestValidateAddFail(t *testing.T) {

	a := syntaxTreeNode{oper: 0, val: 5, nodes: nil}
	b := syntaxTreeNode{oper: byte('+'), val: 0, nodes: []syntaxTreeNode{a}}

	err := b.validate()
	if err == nil {
		t.Fail()
	}
}

func TestValidateSubtractPass(t *testing.T) {

	a := syntaxTreeNode{oper: 0, val: 5, nodes: nil}
	b := syntaxTreeNode{oper: 0, val: 2, nodes: nil}
	c := syntaxTreeNode{oper: byte('-'), val: 0, nodes: []syntaxTreeNode{a, b}}

	err := c.validate()
	if err != nil {
		t.Fail()
	}
}

func TestValidateSubtactFail(t *testing.T) {

	a := syntaxTreeNode{oper: 0, val: 5, nodes: nil}
	b := syntaxTreeNode{oper: byte('-'), val: 0, nodes: []syntaxTreeNode{a}}

	err := b.validate()
	if err == nil {
		t.Fail()
	}
}

func TestValidateMultiplyPass(t *testing.T) {

	a := syntaxTreeNode{oper: 0, val: 5, nodes: nil}
	b := syntaxTreeNode{oper: 0, val: 2, nodes: nil}
	c := syntaxTreeNode{oper: byte('*'), val: 0, nodes: []syntaxTreeNode{a, b}}

	err := c.validate()
	if err != nil {
		t.Fail()
	}
}

func TestValidateMultiplyFail(t *testing.T) {

	a := syntaxTreeNode{oper: 0, val: 5, nodes: nil}
	b := syntaxTreeNode{oper: byte('*'), val: 0, nodes: []syntaxTreeNode{a}}

	err := b.validate()
	if err == nil {
		t.Fail()
	}
}

func TestValidateDividePass(t *testing.T) {

	a := syntaxTreeNode{oper: 0, val: 5, nodes: nil}
	b := syntaxTreeNode{oper: 0, val: 2, nodes: nil}
	c := syntaxTreeNode{oper: byte('/'), val: 0, nodes: []syntaxTreeNode{a, b}}

	err := c.validate()
	if err != nil {
		t.Fail()
	}
}

func TestValidateDivideFail(t *testing.T) {

	a := syntaxTreeNode{oper: 0, val: 5, nodes: nil}
	b := syntaxTreeNode{oper: 0, val: 2, nodes: nil}
	c := syntaxTreeNode{oper: 0, val: 2, nodes: nil}

	d := syntaxTreeNode{oper: byte('/'), val: 0, nodes: []syntaxTreeNode{a, b, c}}

	err := d.validate()
	if err == nil {
		t.Fail()
	}
}

func TestValidateExpPass(t *testing.T) {

	a := syntaxTreeNode{oper: 0, val: 5, nodes: nil}
	b := syntaxTreeNode{oper: 0, val: 2, nodes: nil}
	c := syntaxTreeNode{oper: byte('^'), val: 0, nodes: []syntaxTreeNode{a, b}}

	err := c.validate()
	if err != nil {
		t.Fail()
	}
}

func TestValidateExpFail(t *testing.T) {

	a := syntaxTreeNode{oper: 0, val: 5, nodes: nil}
	b := syntaxTreeNode{oper: 0, val: 2, nodes: nil}
	c := syntaxTreeNode{oper: 0, val: 2, nodes: nil}

	d := syntaxTreeNode{oper: byte('^'), val: 0, nodes: []syntaxTreeNode{a, b, c}}

	err := d.validate()
	if err == nil {
		t.Fail()
	}
}

func TestTraverse(t *testing.T) {

	a := syntaxTreeNode{oper: 0, val: 5, nodes: nil}
	b := syntaxTreeNode{oper: 0, val: 2, nodes: nil}
	c := syntaxTreeNode{oper: 0, val: 0, nodes: []syntaxTreeNode{a, b}}

	c.traverse()
	for _, node := range c.nodes {
		fmt.Println(node)
	}
}

func TestAdd(t *testing.T) {

	a := syntaxTreeNode{oper: 0, val: 5, nodes: nil}
	b := syntaxTreeNode{oper: 0, val: 2, nodes: nil}
	c := syntaxTreeNode{oper: 0, val: 0, nodes: []syntaxTreeNode{a, b}}
	c.add()
	if c.val != 7 {
		t.Fail()
		fmt.Println(c)
	}

}

func TestSubtract(t *testing.T) {

	a := syntaxTreeNode{oper: 0, val: 5, nodes: nil}
	b := syntaxTreeNode{oper: 0, val: 2, nodes: nil}
	c := syntaxTreeNode{oper: 0, val: 0, nodes: []syntaxTreeNode{a, b}}
	c.subtract()
	if c.val != 3 {
		t.Fail()
		fmt.Println(c)
	}

}

func TestMultiply(t *testing.T) {

	a := syntaxTreeNode{oper: 0, val: 5, nodes: nil}
	b := syntaxTreeNode{oper: 0, val: 2, nodes: nil}
	c := syntaxTreeNode{oper: 0, val: 0, nodes: []syntaxTreeNode{a, b}}
	c.multiply()
	if c.val != 10 {
		t.Fail()
		fmt.Println(c)
	}

}

func TestDivide(t *testing.T) {

	a := syntaxTreeNode{oper: 0, val: 5, nodes: nil}
	b := syntaxTreeNode{oper: 0, val: 2, nodes: nil}
	c := syntaxTreeNode{oper: 0, val: 0, nodes: []syntaxTreeNode{a, b}}
	c.divide()
	if c.val != 2.5 {
		t.Fail()
		fmt.Println(c)
	}

}

func TestExp(t *testing.T) {

	a := syntaxTreeNode{oper: 0, val: 5, nodes: nil}
	b := syntaxTreeNode{oper: 0, val: 2, nodes: nil}
	c := syntaxTreeNode{oper: 0, val: 0, nodes: []syntaxTreeNode{a, b}}
	c.exp()
	if c.val != 25 {
		t.Fail()
		fmt.Println(c)
	}

}

func TestEvalutate(t *testing.T) {

	a := syntaxTreeNode{oper: 0, val: 5, nodes: nil}
	b := syntaxTreeNode{oper: 0, val: 2, nodes: nil}
	c := syntaxTreeNode{oper: byte('+'), val: 0, nodes: []syntaxTreeNode{a, b}}
	c.evaluate()
	if c.val != 7 {
		t.Fail()
		fmt.Println(c)
	}
}
