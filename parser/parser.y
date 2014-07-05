%{

package parser

import (
  "fmt"
)

var regs = make([]int, 26)
var Statements []interface{}

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	intval int
  floatval float64
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <intval> expr number
// %type <floatval> ffloat

// same for terminals
%token <intval> DIGIT LETTER

%left '|'
%left '&'
%left '+'  '-'
%left '*'  '/'  '%'
%left UMINUS      /*  supplies  precedence  for  unary  minus  */

%%

list	: /* empty */
	| list stat '\n'
	;

stat	:    expr
		{
			fmt.Printf( "you typed '%d'\n", $1 );
      Statements = append(Statements, $1 );
		}
	|    LETTER '=' expr
		{
			regs[$1]  =  $3
		}
	;

expr	:    '(' expr ')'
		{ $$  =  $2 }
	|    expr '+' expr
		{ $$  =  $1 + $3 }
	|    expr '-' expr
		{ $$  =  $1 - $3 }
	|    expr '*' expr
		{ $$  =  $1 * $3 }
	|    expr '/' expr
		{ $$  =  $1 / $3 }
	|    expr '%' expr
		{ $$  =  $1 % $3 }
	|    expr '&' expr
		{ $$  =  $1 & $3 }
	|    expr '|' expr
		{ $$  =  $1 | $3 }
	|    '-'  expr        %prec  UMINUS
		{ $$  = -$2  }
	|    LETTER
		{ $$  = regs[$1] }
	|    number
	;

number	:    DIGIT
		{ $$ = $1; }
	|    number DIGIT
    { $$ = 10 * $1 + $2 }
  ;

%%      /*  start  of  programs  */
