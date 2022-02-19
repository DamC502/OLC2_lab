// ExprNonTerminal
package Interpreter

import (
	"fmt"
	"log"
	"strconv"
)

type Types int8

const (
	Type_int Types = iota
	Type_double
)

type Result struct {
	int_value     int64
	decimal_value float64
	type_value    Types
}

func (r *Result) String() string {
	switch r.type_value {
	case Type_double:
		return fmt.Sprintf("%f", r.decimal_value)
	case Type_int:
		return fmt.Sprintf("%d", r.int_value)
	default:
		return "Error, valor no definido"
	}
}

func (r *Result) Cast(type_to Types) {
	if type_to == r.type_value {
		return
	}

	switch r.type_value {
	case Type_double:
		switch type_to {
		case Type_int:
			return
		default:
			fmt.Println("Advertencia conversión de tipos inválida")
			return
		}
		return
	case Type_int:
		switch type_to {
		case Type_int:
			return
		case Type_double:
			r.decimal_value = float64(r.int_value)
			r.type_value = Type_double
			return
		default:
			fmt.Println("Advertencia conversión de tipos inválida")
			return
		}
	default:
		fmt.Println("Advertencia conversión de tipos inválida")
		return
	}
}

type ExprInterpreter interface {
	Interpret(r *Result) Result
}

type ExprNode struct {
	left  ExprInterpreter
	right ExprInterpreter
	op    string
}

func NewExprNode(left ExprInterpreter, right ExprInterpreter, op string) ExprNode {
	return ExprNode{left, right, op}
}

type DoubleNode struct {
	number string
}

func NewDoubleNode(s string) DoubleNode {
	return DoubleNode{s}
}

func (d DoubleNode) Interpret(r *Result) Result {
	f, _ := strconv.ParseFloat(d.number, 64)
	return Result{0, f, Type_double}
}

type IntNode struct {
	number string
}

func NewIntNode(s string) IntNode {
	return IntNode{s}
}

func (i IntNode) Interpret(r *Result) Result {
	in, _ := strconv.ParseInt(i.number, 0, 64)
	return Result{in, 0.0, Type_int}
}

func (e ExprNode) Interpret(r *Result) Result {
	if e.left == nil || e.right == nil {
		log.Fatal("nodo Expr Inconsistente")
		return Result{0, 0, Type_int}
	} else if e.right == nil {
		return e.left.Interpret(r)
	} else if e.left == nil {
		return e.right.Interpret(r)
	} else {
		left := e.left.Interpret(r)
		right := e.right.Interpret(r)
		switch e.op {
		case "+":
			return add_op(&left, &right)
		case "*":
			return mul_op(&left, &right)
		default:
			fmt.Println("Operación no soportada")
			return Result{0, 0, Type_int}
		}
	}
}

func mul_op(left *Result, right *Result) Result {
	right.Cast(left.type_value)
	left.Cast(right.type_value)

	if right.type_value == left.type_value {
		switch right.type_value {
		case Type_int:
			right.int_value = left.int_value * right.int_value
		case Type_double:
			right.decimal_value = left.decimal_value * right.decimal_value
		default:
			fmt.Println("Operacion invalida con el tipo: ", right.type_value)
		}
		return *right
	} else {
		fmt.Print("Error tipos inconpatibles")
		return Result{0, 0.0, Type_int}
	}
}

func add_op(left *Result, right *Result) Result {
	right.Cast(left.type_value)
	left.Cast(right.type_value)

	if right.type_value == left.type_value {
		switch right.type_value {
		case Type_int:
			right.int_value = left.int_value + right.int_value
		case Type_double:
			right.decimal_value = left.decimal_value + right.decimal_value
		default:
			fmt.Println("Operacion invalida con el tipo: ", right.type_value)
		}
		return *right
	} else {
		fmt.Print("Error tipos inconpatibles")
		return Result{0, 0.0, Type_int}
	}
}
