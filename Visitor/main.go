package main

import (
	"fmt"
	"strings"
)

func main() {
	//Expression:  1 + (2+3)
	stringsBuilder := strings.Builder{}

	// //Intrusive
	// fmt.Println("############### INTRUSIVE ##################")
	// expression := &AdditionExpression{
	// 	&DoubleExpression{1},
	// 	&AdditionExpression{
	// 		&DoubleExpression{2},
	// 		&DoubleExpression{3},
	// 	},
	// }

	// expression.Print(&stringsBuilder)
	// fmt.Println(stringsBuilder.String())

	// //Reflective
	// fmt.Println("############### REFLECTIVE ##################")
	// expressionTwo := &AdditionExpressionTwo{
	// 	&DoubleExpressionTwo{1},
	// 	&AdditionExpressionTwo{
	// 		&DoubleExpressionTwo{2},
	// 		&DoubleExpressionTwo{3},
	// 	},
	// }

	// Print(expressionTwo, &stringsBuilder)
	// fmt.Println(stringsBuilder.String())

	//Reflective
	fmt.Println("############### CLASSIC ##################")
	expressionThree := &AdditionExpressionThree{
		&DoubleExpressionThree{1},
		&AdditionExpressionThree{
			&DoubleExpressionThree{2},
			&DoubleExpressionThree{3},
		},
	}

	visitor := &ExpressionPrinter{sb: &stringsBuilder}

	expressionThree.Accept(visitor)
	fmt.Println(stringsBuilder.String())
}

//------------------------------INTRUSIVE VISITOR---------------------------------------------
// Inicialmente a Interface não possui o método Print
// Adicionamos o método Print devido a uma necessidade (de negócio) de todos os tipos que implementam esta
// interface precisarem ter uma maneira particular de impressão
// A modificação na Interface quebra o OCP e caracteriza o Visitante Intrusivo
type Expression interface {
	Print(sb *strings.Builder)
}

type DoubleExpression struct {
	value float64
}

func (doubleExpression *DoubleExpression) Print(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%g", doubleExpression.value))
}

type AdditionExpression struct {
	left, right Expression
}

func (additionExpression *AdditionExpression) Print(sb *strings.Builder) {
	sb.WriteRune('(')
	additionExpression.left.Print(sb)
	sb.WriteRune('+')
	additionExpression.right.Print(sb)
	sb.WriteRune(')')
}

//-----------------------------REFLECTIVE VISITOR---------------------------
//Diferentemente do Visitante Intrusivo, que é caracterizado por um quebra no OCP, aqui o Visitante
//é caracterizado por ser uma simples FUNÇÃO
//Não quebra o OCP e também respeita o SRP, visto que a responsabilidade da nova funcionalidade é
//isolada nesta nova função.
type ExpressionTwo interface{}

type DoubleExpressionTwo struct {
	value float64
}

type AdditionExpressionTwo struct {
	left, right ExpressionTwo
}

func Print(expression ExpressionTwo, sb *strings.Builder) {
	if doubleExpression, ok := expression.(*DoubleExpressionTwo); ok {
		sb.WriteString(fmt.Sprintf("%g", doubleExpression.value))
	} else if additionExpression, ok := expression.(*AdditionExpressionTwo); ok {
		sb.WriteRune('(')
		Print(additionExpression.left, sb)
		sb.WriteRune('+')
		Print(additionExpression.right, sb)
		sb.WriteRune(')')
	}
}

//-------------------------------CLASSIC VISITOR---------------------------------
//Interfaces e Objetos que aceitam visitantes
//Os objetos e interfaces passam o controle para o Visitante, o qual, por sua vez, realiza a funcionalidade
//desejada sobre os mesmos
type ExpressionVisitor interface {
	VisitDoubleExpression(doubleExpression *DoubleExpressionThree)
	VisitAdditionExpression(additionExpression *AdditionExpressionThree)
}

type ExpressionThree interface {
	Accept(visitor ExpressionVisitor)
}

type DoubleExpressionThree struct {
	value float64
}

func (doubleExpression *DoubleExpressionThree) Accept(visitor ExpressionVisitor) {
	//passa o controle para o visitante
	visitor.VisitDoubleExpression(doubleExpression)
}

type AdditionExpressionThree struct {
	left, right ExpressionThree
}

func (additionExpression *AdditionExpressionThree) Accept(visitor ExpressionVisitor) {
	visitor.VisitAdditionExpression(additionExpression)
}

type ExpressionPrinter struct {
	sb *strings.Builder
}

func (expressionPrinter *ExpressionPrinter) VisitDoubleExpression(doubleExpression *DoubleExpressionThree) {
	expressionPrinter.sb.WriteString(fmt.Sprintf("%g", doubleExpression.value))
}

func (expressionPrinter *ExpressionPrinter) VisitAdditionExpression(additionExpression *AdditionExpressionThree) {
	expressionPrinter.sb.WriteRune('(')
	additionExpression.left.Accept(expressionPrinter)
	expressionPrinter.sb.WriteRune('+')
	additionExpression.right.Accept(expressionPrinter)
	expressionPrinter.sb.WriteRune(')')

}
