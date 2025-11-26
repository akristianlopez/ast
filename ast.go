package ast

type Node interface {
	Pos() int
	End() int
	// String()
}

type Program struct {
	Knowledge    StringLiteral
	Declarations *BlockDclStatement //Statement
	Start        *StartStatement
}

type StartStatement struct {
	ObjectName StringLiteral
	Statements []Statement
}

type LetFuncTypeDcl interface {
	Node
	statementNode()
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Déclarations
type FunctionDecl struct {
	StartPos   int
	Name       *Identifier
	Parameters []*Parameter
	ReturnType Type
	Body       *BlockStatement
}

type Parameter struct {
	Name *Identifier
	Type Type
}

// Records
type RecordDecl struct {
	StartPos   int
	Name       *Identifier
	Parameters []*Parameter
}

// Types
type Type interface {
	Node
	typeNode()
}

type BasicType struct {
	StartPos int
	Name     string // "integer", "float", "string", "boolean","time", "date"
}

type ArrayType struct {
	StartPos    int
	ElementType Type
}

// Statements
type BlockStatement struct {
	StartPos   int
	Statements []Statement
}

// Statements
type BlockDclStatement struct {
	StartPos   int
	Statements []VariableDecl
}

type IfStatement struct {
	StartPos  int
	Condition Expression
	Then      Statement
	Else      Statement
}

type SelectStatement struct {
	StartPos   int
	Condition  Expression
	Statements []CaseStatement
}

type CaseStatement struct {
	StartPos   int
	Expr       Expression
	Statements []Statement
}

type WhileStatement struct {
	StartPos  int
	Condition Expression
	Body      Statement
}

type ForStatement struct {
	StartPos  int
	Init      Statement
	Condition Expression
	Update    Statement
	Body      Statement
}

type ForEachStatement struct {
	StartPos int
	Variable *Identifier
	Iterator Expression
	Body     Statement
}

type TypeArgs struct {
	Name  string
	Value IntegerLiteral
}

type VariableDecl struct {
	StartPos  int
	Name      *Identifier
	Type      Type
	Arguments []TypeArgs
	Value     Expression
}

type ReturnStatement struct {
	StartPos int
	Value    Expression
}

type ExpressionStatement struct {
	StartPos   int
	Expression Expression
}

// Expressions
type Identifier struct {
	StartPos int
	Value    string
}

type IntegerLiteral struct {
	StartPos int
	Value    int64
}

type FloatLiteral struct {
	StartPos int
	Value    float64
}

type StringLiteral struct {
	StartPos int
	Value    string
}

type BooleanLiteral struct {
	StartPos int
	Value    bool
}

type BinaryExpression struct {
	StartPos int
	Left     Expression
	Operator string
	Right    Expression
}

type CallExpression struct {
	StartPos  int
	Function  Expression
	Arguments []Expression
}

type ArrayLiteral struct {
	StartPos int
	Elements []Expression
}

// Implémentations des méthodes Node
func (p *Program) Pos() int { return 0 }
func (p *Program) End() int { return 0 }

// func (b BasicType) typeNode() {}
// func (b BasicType) End() int  { return b.StartPos }
// func (b BasicType) Pos() int  { return b.StartPos }

func (f *FunctionDecl) Pos() int       { return f.StartPos }
func (f *FunctionDecl) End() int       { return f.Body.End() }
func (f *FunctionDecl) statementNode() {}

func (b *BlockStatement) Pos() int { return b.StartPos }
func (b *BlockStatement) End() int {
	if len(b.Statements) > 0 {
		return b.Statements[len(b.Statements)-1].End()
	}
	return b.StartPos
}
func (b *BlockStatement) statementNode() {}

func (i *IfStatement) Pos() int { return i.StartPos }
func (i *IfStatement) End() int {
	end := i.Then.End()
	if i.Else != nil {
		end = i.Else.End()
	}
	return end
}
func (i *IfStatement) statementNode() {}

func (w *WhileStatement) Pos() int       { return w.StartPos }
func (w *WhileStatement) End() int       { return w.Body.End() }
func (w *WhileStatement) statementNode() {}

func (f *ForStatement) Pos() int       { return f.StartPos }
func (f *ForStatement) End() int       { return f.Body.End() }
func (f *ForStatement) statementNode() {}

func (fe *ForEachStatement) Pos() int       { return fe.StartPos }
func (fe *ForEachStatement) End() int       { return fe.Body.End() }
func (fe *ForEachStatement) statementNode() {}

func (v *VariableDecl) Pos() int       { return v.StartPos }
func (v *VariableDecl) End() int       { return v.Value.End() }
func (v *VariableDecl) statementNode() {}

func (r *ReturnStatement) Pos() int       { return r.StartPos }
func (r *ReturnStatement) End() int       { return r.Value.End() }
func (r *ReturnStatement) statementNode() {}

func (e *ExpressionStatement) Pos() int       { return e.StartPos }
func (e *ExpressionStatement) End() int       { return e.Expression.End() }
func (e *ExpressionStatement) statementNode() {}

func (i *Identifier) Pos() int        { return i.StartPos }
func (i *Identifier) End() int        { return i.StartPos + len(i.Value) }
func (i *Identifier) expressionNode() {}

func (il *IntegerLiteral) Pos() int        { return il.StartPos }
func (il *IntegerLiteral) End() int        { return il.StartPos }
func (il *IntegerLiteral) expressionNode() {}

func (fl *FloatLiteral) Pos() int        { return fl.StartPos }
func (fl *FloatLiteral) End() int        { return fl.StartPos }
func (fl *FloatLiteral) expressionNode() {}

func (sl *StringLiteral) Pos() int        { return sl.StartPos }
func (sl *StringLiteral) End() int        { return sl.StartPos + len(sl.Value) + 2 }
func (sl *StringLiteral) expressionNode() {}

func (bl *BooleanLiteral) Pos() int        { return bl.StartPos }
func (bl *BooleanLiteral) End() int        { return bl.StartPos }
func (bl *BooleanLiteral) expressionNode() {}

func (b *BinaryExpression) Pos() int        { return b.Left.Pos() }
func (b *BinaryExpression) End() int        { return b.Right.End() }
func (b *BinaryExpression) expressionNode() {}

func (c *CallExpression) Pos() int        { return c.Function.Pos() }
func (c *CallExpression) End() int        { return c.Function.End() }
func (c *CallExpression) expressionNode() {}

func (a *ArrayLiteral) Pos() int        { return a.StartPos }
func (a *ArrayLiteral) End() int        { return a.StartPos }
func (a *ArrayLiteral) expressionNode() {}
