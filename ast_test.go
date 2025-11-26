package ast

// main_test.go (pour montrer la construction manuelle de l'AST)
// import (
// 	"fmt"
// 	"testing"
// )

// func TestASTConstruction(t *testing.T) {
// 	// Exemple 1: Une simple action avec une affectation et un appel de fonction
// 	program := &Action{
// 		Name: "action d'affection",
// 		Statements: []Statement{
// 			&ExpressionStatement{
// 				Expression: &InfixExpression{
// 					Left:     &Identifier{Value: "x"},
// 					Operator: "=",
// 					Right: &InfixExpression{
// 						Left:     &NumberLiteral{Value: "10"},
// 						Operator: "+",
// 						Right:    &NumberLiteral{Value: "5"},
// 					},
// 				},
// 			},
// 			&ExpressionStatement{
// 				Expression: &CallExpression{
// 					Function:  &Identifier{Value: "print"},
// 					Arguments: []Expression{&Identifier{Value: "x"}},
// 				},
// 			},
// 		},
// 	}

// 	// Ceci est la représentation textuelle de l'AST que nous avons construit manuellement.
// 	// Dans un vrai scénario, le parser construirait cela à partir du code source.
// 	expected := "\t\taction d'affection\r\n(x = (10 + 5))\nprint(x)\n"
// 	if program.String() != expected {
// 		t.Errorf("action.String() wrong. expected=\n%q\ngot=\n%q", expected, program.String())
// 	}

// 	fmt.Println("--- Exemple d'action Simple ---")
// 	fmt.Println(program.String())

// 	// Exemple 2: Déclaration de fonction, if, while
// 	programWithComplexities := &Action{
// 		Name: "Fonction",
// 		Statements: []Statement{
// 			&FunctionDeclaration{
// 				Name: &Identifier{Value: "fibonacci", Type: &Type_name{Value: "Number",
// 					Intval: &NumberLiteral{Value: ""}, Decval: &NumberLiteral{Value: ""}, MinVal: &NumberLiteral{Value: ""},
// 					MaxVal: &NumberLiteral{Value: ""}}},
// 				Parameters: []*Identifier{{Value: "n", Type: &Type_name{Value: "Number",
// 					Intval: &NumberLiteral{Value: "10"}, Decval: &NumberLiteral{Value: ""}, MinVal: &NumberLiteral{Value: ""},
// 					MaxVal: &NumberLiteral{Value: ""}}}},
// 				Body: &BlockStatement{
// 					Statements: []Statement{
// 						&IfStatement{
// 							Condition: &InfixExpression{
// 								Left:     &Identifier{Value: "n"},
// 								Operator: "<",
// 								Right:    &NumberLiteral{Value: "2"},
// 							},
// 							Consequence: &BlockStatement{
// 								Statements: []Statement{
// 									&ReturnStatement{ReturnValue: &Identifier{Value: "n"}},
// 								},
// 							},
// 						},
// 						&ReturnStatement{
// 							ReturnValue: &InfixExpression{
// 								Left: &CallExpression{
// 									Function: &Identifier{Value: "fibonacci"},
// 									Arguments: []Expression{&InfixExpression{
// 										Left:     &Identifier{Value: "n"},
// 										Operator: "-",
// 										Right:    &NumberLiteral{Value: "1"},
// 									}},
// 								},
// 								Operator: "+",
// 								Right: &CallExpression{
// 									Function: &Identifier{Value: "fibonacci"},
// 									Arguments: []Expression{&InfixExpression{
// 										Left:     &Identifier{Value: "n"},
// 										Operator: "-",
// 										Right:    &NumberLiteral{Value: "2"},
// 									}},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			&WhileStatement{
// 				Condition: &InfixExpression{
// 					Left:     &Identifier{Value: "i"},
// 					Operator: "<",
// 					Right:    &NumberLiteral{Value: "10"},
// 				},
// 				Body: &BlockStatement{
// 					Statements: []Statement{
// 						&ExpressionStatement{
// 							Expression: &CallExpression{
// 								Function:  &Identifier{Value: "print"},
// 								Arguments: []Expression{&Identifier{Value: "i"}},
// 							},
// 						},
// 						&ExpressionStatement{
// 							Expression: &InfixExpression{
// 								Left:     &Identifier{Value: "i"},
// 								Operator: "=",
// 								Right: &InfixExpression{
// 									Left:     &Identifier{Value: "i"},
// 									Operator: "+",
// 									Right:    &NumberLiteral{Value: "1"},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	fmt.Println("\n--- Exemple de Programme avec Fonctions, If, While ---")
// 	fmt.Println(programWithComplexities.String())

// 	// Un vrai test vérifierait la structure de l'AST, pas juste la chaîne.
// 	// Mais pour l'affichage, c'est utile.
// 	// Vous pouvez ajouter d'autres assertions ici pour vérifier la structure exacte.
// 	// Par exemple:
// 	// if len(programWithComplexities.Statements) != 2 { t.Errorf("Expected 2 statements") }
// 	// if _, ok := programWithComplexities.Statements[0].(*FunctionDeclaration); !ok { t.Errorf("Expected FunctionDeclaration") }
// }
