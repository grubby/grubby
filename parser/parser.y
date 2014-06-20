%{

package parser

var regs = make([]int, 26)

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.

// effectively you can use this to declare what values the lexer
// will want to save for each token it handles
%union{
  Val int
  FloatVal float64
}

// any token that returns a value needs a type
// these will become field names in the above struct
%type <Val> expr number

%token <Val> DIGIT FLOAT

%%

list  : /* empty */
  | list statement '\n'
  ;

statement  :    expr;

expr  :    number;

number  :   DIGIT
    { $$ = $1 }
  |    FLOAT
    { $$ = $1 }
  ;

%%
