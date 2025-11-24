package parsers

import (
	"fmt"
)

// --- 1. Interfaces de Base de l'AST ---

// Node est l'interface racine pour tous les nœuds de l'AST.
type Node interface {
	String() string // Pour l'affichage debug de l'AST
}

// Statement est une instruction (par exemple, un `if`, une boucle, une affectation).
type Statement interface {
	Node
	statementNode() // Méthode de "marquage" pour distinguer les Statements
}

// Expression est une expression (par exemple, un nombre, une variable, un appel de fonction).
type Expression interface {
	Node
	expressionNode() // Méthode de "marquage" pour distinguer les Expressions
}

// --- 2. Implémentations concrètes des nœuds de l'AST ---

// Programme représente le nœud racine de tout le code source.
type Action struct {
	Name       string
	Statements []Statement
}

func (p *Action) String() string {
	var s string
	for _, stmt := range p.Statements {
		s += stmt.String() + "\n"
	}
	return "\t\t" + p.Name + "\r\n" + s
}

// --- 2.1 Nœuds d'Instructions (Statements) ---

// ExpressionStatement est une instruction qui est juste une expression (ex: "5 + x;")
type ExpressionStatement struct {
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) String() string { return es.Expression.String() }

// BlockStatement représente un bloc d'instructions (souvent entre accolades {}).
type BlockStatement struct {
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}
func (bs *BlockStatement) String() string {
	var s string
	for _, stmt := range bs.Statements {
		s += stmt.String() + "; " // Simplifié pour l'affichage
	}
	// return "{ " + s + " }"
	return s
}

// IfStatement représente une structure conditionnelle 'if'.
type IfStatement struct {
	Condition   Expression
	Consequence *BlockStatement // Bloc d'instructions si la condition est vraie
	Alternative *BlockStatement // Bloc d'instructions si la condition est fausse (optionnel)
}

func (is *IfStatement) statementNode() {}
func (is *IfStatement) String() string {
	str := fmt.Sprintf("IF (%s) THEN %s", is.Condition.String(), is.Consequence.String())
	if is.Alternative != nil {
		str += fmt.Sprintf(" ELSE %s", is.Alternative.String())
	}
	return str + "END IF"
}

// WhileStatement représente une boucle 'while'.
type WhileStatement struct {
	Condition Expression
	Body      *BlockStatement
}

func (ws *WhileStatement) statementNode() {}
func (ws *WhileStatement) String() string {
	return fmt.Sprintf("WHILE (%s) %s END WHILE", ws.Condition.String(), ws.Body.String())
}

// FunctionDeclaration représente la déclaration d'une fonction.
type FunctionDeclaration struct {
	Name       *Identifier
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fd *FunctionDeclaration) statementNode() {}
func (fd *FunctionDeclaration) String() string {
	params := make([]string, len(fd.Parameters))
	for i, p := range fd.Parameters {
		params[i] = p.String()
	}
	return fmt.Sprintf("FUNCTION %s(%s) %s END FUNCTION", fd.Name.String(), join(params, ", "), fd.Body.String())
}

// ReturnStatement représente l'instruction 'return'.
type ReturnStatement struct {
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) String() string {
	return fmt.Sprintf("RETURN %s", rs.ReturnValue.String())
}

// --- 2.2 Nœuds d'Expressions (Expressions) ---

// Identifier représente un nom (variable, nom de fonction).
type Identifier struct {
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) String() string  { return i.Value }

// IntegerLiteral représente un nombre entier.
type NumberLiteral struct {
	Value string
}

func (il *NumberLiteral) expressionNode() {}
func (il *NumberLiteral) String() string {
	// return strconv.FormatInt(il.Value, 10)
	return il.Value
}

// BooleanLiteral représente un booléen (true/false).
type BooleanLiteral struct {
	Value string
}

func (bl *BooleanLiteral) expressionNode() {}
func (bl *BooleanLiteral) String() string  { return bl.Value /*strconv.FormatBool(bl.Value)*/ }

// InfixExpression représente une opération binaire (ex: 5 + 3, x > y).
type InfixExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}
func (ie *InfixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", ie.Left.String(), ie.Operator, ie.Right.String())
}

// CallExpression représente un appel de fonction.
type CallExpression struct {
	Function  Expression // Peut être un Identifier ou une autre Expression (ex: une fonction anonyme)
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) String() string {
	args := make([]string, len(ce.Arguments))
	for i, a := range ce.Arguments {
		args[i] = a.String()
	}
	return fmt.Sprintf("%s(%s)", ce.Function.String(), join(args, ", "))
}

// --- Fonctions utilitaires ---
func join(s []string, sep string) string {
	if len(s) == 0 {
		return ""
	}
	res := s[0]
	for i := 1; i < len(s); i++ {
		res += sep + s[i]
	}
	return res
}
