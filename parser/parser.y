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
%type <genericNumber> expr

// same for terminals
%token <intval> DIGIT
%token <floatval> FLOAT

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
  |    DIGIT
    { $$ = $1 }
  |    FLOAT
    { $$ = $1; }
	;

%%
