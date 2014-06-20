%{

package parser

var regs = make([]int, 26)
var base int

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
  | list stat '\n'
  ;

stat  :    expr
    {
      return $1
    }
  ;

expr  :    number
  ;

number  :   DIGIT
    {
      $$ = $1
      if $1 == 0 {
        base = 8
      } else {
        base = 10
      }
    }
  |    number DIGIT
    { $$ = base * $1 + $2 }
  ;

%%
