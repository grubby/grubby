%{

package parser

import(
  "github.com/grubby/grubby/ast"
)

var Statements []ast.Node

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
  genericValue ast.Node
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%token <genericValue> NODE

%%

list	: /* empty */
	| list statement '\n'
	;

statement : NODE
  { Statements = append(Statements, $1); };


%%
