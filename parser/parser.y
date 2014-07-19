%{

package parser

import(
  "github.com/grubby/grubby/ast"
)

var Statements []ast.Node

%}

// fields inside this union end up as the fields in a structure known
// as RubySymType, of which a reference is passed to the lexer.
%union{
  genericValue ast.Node
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%token <genericValue> NODE
%token <genericValue> REF
%token <genericValue> LPAREN
%token <genericValue> RPAREN
%token <genericValue> COMMA

/*
  eg: if you want to be able to assign to something in the RubySymType
      struct, or if you want a terminating node below, you'll want to
      declare a type (or possibly just a token)
*/

%type <genericValue> expr
%type <genericValue> callexpr
%type <genericValue> callargs

%%

list	: /* empty */ | list statement;

statement : callexpr
  { Statements = append(Statements, $1); }
| expr
  { Statements = append(Statements, $1); }
| expr '\n'
  { Statements = append(Statements, $1); };

callexpr : REF callargs
  {
    $$ = ast.CallExpression{
      Func: $1.StringValue(),
      Args: []ast.Node{$2},
    }
  }
| REF " " callargs
  {
    $$ = ast.CallExpression{
      Func: $1.StringValue(),
      Args: []ast.Node{$3},
    }
  };

callargs : NODE | LPAREN nodes_with_commas RPAREN;
nodes_with_commas: NODE | nodes_with_commas COMMA NODE;
expr : NODE | REF;


%%
