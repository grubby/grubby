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
  genericSlice ast.Nodes
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
%type <genericSlice> callargs
%type <genericSlice> nodes_with_commas

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
      Func: $1,
      Args: $2,
    }
  }
| REF " " callargs
  {
    $$ = ast.CallExpression{
      Func: $1,
      Args: $3,
    }
  };

callargs : NODE
  {  $$ = append($$, $1) }
| LPAREN nodes_with_commas RPAREN
  { $$ = $2 };

nodes_with_commas: /* empty */ { $$ = ast.Nodes{} }
| NODE
  { $$ = append($$, $1); }
| nodes_with_commas COMMA " " NODE
  { $$ = append($$, $4); };

expr : NODE | REF;


%%
