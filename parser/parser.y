%{

package parser

var Statements []interface{}

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	intval int
  genericNumber interface{}
  floatval float64
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <genericNumber> expr number ffloat

// same for terminals
%token <intval> DIGIT
%token <floatval> FLOAT

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
      Statements = append(Statements, $1 );
		};

expr	:    '(' expr ')'
		{ $$  =  $2 }
	/* |    expr '+' expr */
	/* 	{ $$  =  $1 + $3 } */
	/* |    expr '-' expr */
	/* 	{ $$  =  $1 - $3 } */
	/* |    expr '*' expr */
	/* 	{ $$  =  $1 * $3 } */
	/* |    expr '/' expr */
	/* 	{ $$  =  $1 / $3 } */
	/* |    expr '%' expr */
	/* 	{ $$  =  $1 % $3 } */
	/* |    expr '&' expr */
	/* 	{ $$  =  $1 & $3 } */
	/* |    expr '|' expr */
	/* 	{ $$  =  $1 | $3 } */
	/* |    '-'  expr        %prec  UMINUS */
	/* 	{ $$  = -$2  } */
  |    DIGIT
    { $$ = $1 }
  |    FLOAT
    { $$ = $1; }
	;

number	:    DIGIT
		{ $$ = $1; }
  ;

ffloat   :    FLOAT
    { $$ = $1; }
  ;

%%
