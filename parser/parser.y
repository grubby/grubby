%{

package parser

import(
  "strings"
  "github.com/grubby/grubby/ast"
)

var Statements []ast.Node

%}

// fields inside this union end up as the fields in a structure known
// as RubySymType, of which a reference is passed to the lexer.
%union{
  genericValue ast.Node
  genericSlice ast.Nodes
  stringSlice []string
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%token <genericValue> NODE
%token <genericValue> REF
%token <genericValue> CAPITAL_REF
%token <genericValue> LPAREN
%token <genericValue> RPAREN
%token <genericValue> COMMA

// keywords
%token <genericValue> DEF
%token <genericValue> END
%token <genericValue> CLASS
%token <genericValue> MODULE

// booleans
%token <genericValue> TRUE
%token <genericValue> FALSE

// operators
%token <genericValue> LESSTHAN
%token <genericValue> GREATERTHAN
%token <genericValue> EQUALTO
%token <genericValue> BANG
%token <genericValue> COMPLEMENT
%token <genericValue> POSITIVE
%token <genericValue> NEGATIVE
%token <genericValue> STAR
%token <genericValue> HASHROCKET

// misc
%token <genericValue> COLON
%token <genericValue> DOT
%token <genericValue> LBRACKET // "["
%token <genericValue> RBRACKET // "]"
%token <genericValue> LBRACE   // "{"
%token <genericValue> RBRACE   // "}"

/*
  eg: if you want to be able to assign to something in the RubySymType
      struct, or if you want a terminating node below, you will want to
      declare a type (or possibly just a token)
*/

%type <genericValue> whitespace
%type <genericValue> optional_newline

// single nodes
%type <genericValue> expr
%type <genericValue> true
%type <genericValue> hash
%type <genericValue> false
%type <genericValue> array
%type <genericValue> symbol
%type <genericValue> callexpr
%type <genericValue> statement
%type <genericValue> assignment
%type <genericValue> func_declaration
%type <genericValue> class_declaration
%type <genericValue> module_declaration
%type <genericValue> class_name_with_modules

// unary operator nodes
%type <genericValue> negation   // !
%type <genericValue> complement // ~
%type <genericValue> positive   // +
%type <genericValue> negative   // -

// binary operator nodes
%type <genericValue> binary_addition       // 2 + 3
%type <genericValue> binary_subtraction    // 2 - 3
%type <genericValue> binary_multiplication // 2 * 3

// slice nodes
%type <genericSlice> list
%type <genericSlice> callargs
%type <genericSlice> capture_list
%type <genericSlice> key_value_pairs
%type <genericSlice> symbol_key_value_pairs
%type <genericSlice> nodes_with_commas
%type <stringSlice> namespaced_modules

%%

capture_list : /* empty */
  { Statements = []ast.Node{} }
| capture_list statement
  {
		if $2 != nil {
			Statements = append(Statements, $2)
		}
	};

list : /* empty */
  { $$ = []ast.Node{} }
| list statement
  {
		if $2 != nil {
			$$ = append($$, $2)
		}
	};

symbol : COLON REF
  {
    $$ = ast.Symbol{Name: $2.(ast.BareReference).Name}
  };

statement : callexpr whitespace
  { $$ = $1 }
| expr whitespace
  { $$ = $1 }
| '\n'
  { $$ = nil; /* ignores new lines around statements */ }
| whitespace
  { $$ = nil; /* ignores whitespace around statements */ };

callexpr : REF whitespace callargs
  {
    $$ = ast.CallExpression{
      Func: $1,
      Args: $3,
    }
  };

expr : NODE | REF | CAPITAL_REF | func_declaration | class_declaration | symbol | module_declaration | assignment | true | false | negation | complement | positive | negative | binary_addition | binary_subtraction | binary_multiplication | array | hash;

callargs : /* empty */ { $$ = ast.Nodes{} }
| NODE
  { $$ = append($$, $1) }
| REF
  { $$ = append($$, $1) }
| LPAREN whitespace nodes_with_commas whitespace RPAREN
  {
		$$ = $3
	}
| nodes_with_commas
  {
    $$ = $1
  };

nodes_with_commas : /* empty */ { $$ = ast.Nodes{} }
| REF
  { $$ = append($$, $1); }
| NODE
  { $$ = append($$, $1); }
| nodes_with_commas whitespace COMMA whitespace NODE
  { $$ = append($$, $5); };
| nodes_with_commas whitespace COMMA whitespace REF
  { $$ = append($$, $5); };

// FIXME: this should use a different type than callargs
// call args can be a list of expressions. This is just a list of REFs or NODEs
func_declaration : DEF whitespace REF whitespace callargs whitespace "\n" list END
  {
		$$ = ast.FuncDecl{
			Name: $3.(ast.BareReference),
      Args: $5,
			Body: $8,
    }
  };

class_declaration: CLASS whitespace class_name_with_modules "\n" list END
  {
    $$ = ast.ClassDecl{
       Name: $3.(ast.Class).Name,
       Body: $5,
    }
  }
| CLASS whitespace class_name_with_modules LESSTHAN class_name_with_modules "\n" list END
  {
    $$ = ast.ClassDecl{
       Name: $3.(ast.Class).Name,
       SuperClass: $5.(ast.Class),
       Namespace: $3.(ast.Class).Namespace,
       Body: $7,
    }
  };

module_declaration: MODULE whitespace class_name_with_modules "\n" list END
  {
    $$ = ast.ModuleDecl{
      Name: $3.(ast.Class).Name,
      Namespace: $3.(ast.Class).Namespace,
      Body: $5,
    }
  };

class_name_with_modules: whitespace CAPITAL_REF whitespace
  {
    $$ = ast.Class{
      Name: $2.(ast.BareReference).Name,
    }
  }
| whitespace namespaced_modules COLON COLON CAPITAL_REF whitespace
  {
    $$ = ast.Class{
       Name: $5.(ast.BareReference).Name,
       Namespace: strings.Join($2, "::"),
    }
  };

namespaced_modules : CAPITAL_REF
  {
    $$ = append($$, $1.(ast.BareReference).Name)
  }
|  namespaced_modules COLON COLON CAPITAL_REF
  {
    $$ = append($$, $4.(ast.BareReference).Name)
  };

assignment: REF whitespace EQUALTO whitespace expr
  {
    $$ = ast.Assignment{
      LHS: $1,
      RHS: $5,
    }
  };

negation: BANG whitespace expr { $$ = ast.Negation{Target: $3} };
complement: COMPLEMENT whitespace expr { $$ = ast.Complement{Target: $3} };
positive: POSITIVE whitespace expr { $$ = ast.Positive{Target: $3} };
negative: NEGATIVE whitespace expr { $$ = ast.Negative{Target: $3} };

binary_addition: expr whitespace POSITIVE whitespace expr
  {
    $$ = ast.Addition{
      LHS: $1,
      RHS: $5,
    }
  };

binary_subtraction: expr whitespace NEGATIVE whitespace expr
  {
    $$ = ast.Subtraction{
      LHS: $1,
      RHS: $5,
    }
  };

binary_multiplication: expr whitespace STAR whitespace expr
  {
    $$ = ast.Multiplication{
      LHS: $1,
      RHS: $5,
    }
  };

whitespace : /* zero or more */
  { $$ = nil }
| " " whitespace
  { $$ = nil }
| "\t" whitespace
  { $$ = nil };

true: TRUE { $$ = ast.Boolean{Value: true} }
false: FALSE { $$ = ast.Boolean{Value: false} }
array : LBRACKET whitespace nodes_with_commas whitespace RBRACKET
  {
    $$ = ast.Array{Nodes: $3};
  };
hash: LBRACE whitespace optional_newline whitespace key_value_pairs whitespace optional_newline whitespace RBRACE
  {
    pairs := []ast.HashKeyValuePair{}
    for _, node := range $5 {
      pairs = append(pairs, node.(ast.HashKeyValuePair))
    }
    $$ = ast.Hash{Pairs: pairs}
  }
| LBRACE whitespace optional_newline whitespace symbol_key_value_pairs whitespace optional_newline whitespace RBRACE
  {
    pairs := []ast.HashKeyValuePair{}
    for _, node := range $5 {
      pairs = append(pairs, node.(ast.HashKeyValuePair))
    }
    $$ = ast.Hash{Pairs: pairs}
  };

optional_newline : /* possibly nothing */
  { $$ = nil }
| "\n"
  { $$ = nil }
| optional_newline whitespace "\n" whitespace
  { $$ = nil };

key_value_pairs: /* empty */ { $$ = ast.Nodes{} }
| expr whitespace HASHROCKET whitespace expr
  { $$ = append($$, ast.HashKeyValuePair{Key: $1, Value: $5}) }
| key_value_pairs whitespace COMMA optional_newline whitespace expr whitespace HASHROCKET whitespace expr
  { $$ = append($$, ast.HashKeyValuePair{Key: $5, Value: $9}) }
| key_value_pairs whitespace COMMA optional_newline whitespace expr whitespace HASHROCKET whitespace expr COMMA
  { $$ = append($$, ast.HashKeyValuePair{Key: $5, Value: $9}) };

symbol_key_value_pairs: /* empty */ { $$ = ast.Nodes{} }
| REF COLON whitespace expr
  {
    $$ = append($$, ast.HashKeyValuePair{
      Key: ast.Symbol{Name: $1.(ast.BareReference).Name},
      Value: $4,
    })
  }
| symbol_key_value_pairs whitespace COMMA optional_newline whitespace REF COLON whitespace expr
  {
    $$ = append($$, ast.HashKeyValuePair{
      Key: ast.Symbol{Name: $6.(ast.BareReference).Name},
      Value: $9,
    })
  }
| symbol_key_value_pairs whitespace COMMA optional_newline whitespace REF COLON whitespace expr COMMA
  {
    $$ = append($$, ast.HashKeyValuePair{
      Key: ast.Symbol{Name: $6.(ast.BareReference).Name},
      Value: $9,
    })
  };


optional_comma: /* nothing */ | COMMA;

%%
