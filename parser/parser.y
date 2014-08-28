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
  operator     string
  genericValue ast.Node
  genericSlice ast.Nodes
  stringSlice []string
}

%token <operator> OPERATOR

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%token <genericValue> NODE
%token <genericValue> REF
%token <genericValue> CAPITAL_REF
%token <genericValue> LPAREN
%token <genericValue> RPAREN
%token <genericValue> COMMA

// keywords
%token <genericValue> DO
%token <genericValue> DEF
%token <genericValue> END
%token <genericValue> IF
%token <genericValue> ELSE
%token <genericValue> ELSIF
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

// misc
%token <genericValue> WHITESPACE
%token <genericValue> NEWLINE
%token <genericValue> SEMICOLON
%token <genericValue> COLON
%token <genericValue> DOT
%token <genericValue> PIPE       // "|"
%token <genericValue> SLASH      // "/"
%token <genericValue> AMPERSAND  // "&"
%token <genericValue> MODULO     // "%"
%token <genericValue> CARET      // "^"
%token <genericValue> LBRACKET   // "["
%token <genericValue> RBRACKET   // "]"
%token <genericValue> LBRACE     // "{"
%token <genericValue> RBRACE     // "}"
%token <genericValue> DOLLARSIGN // "$"
%token <genericValue> ATSIGN     // "@"
%token <genericValue> FILE_CONST_REF // __FILE__
%token <genericValue> EOF

/*
  eg: if you want to be able to assign to something in the RubySymType
      struct, or if you want a terminating node below, you will want to
      declare a type (or possibly just a token)
*/

%type <genericValue> optional_comma

// single nodes
%type <genericValue> expr
%type <genericValue> line
%type <genericValue> true
%type <genericValue> hash
%type <genericValue> block
%type <genericValue> false
%type <genericValue> array
%type <genericValue> global
%type <genericValue> callexpr
%type <genericValue> if_block
%type <genericValue> assignment
%type <genericValue> class_variable
%type <genericValue> func_declaration
%type <genericValue> class_declaration
%type <genericValue> instance_variable
%type <genericValue> module_declaration
%type <genericValue> class_name_with_modules
%type <genericValue> filename_const_reference

// unary operator nodes
%type <genericValue> negation   // !
%type <genericValue> complement // ~
%type <genericValue> positive   // +
%type <genericValue> negative   // -

// binary operator nodes
%type <genericValue> binary_addition       // 2 + 3
%type <genericValue> binary_subtraction    // 2 - 3
%type <genericValue> binary_multiplication // 2 * 3
%type <genericValue> binary_division       // 2 / 3
%type <genericValue> bitwise_and           // 2 & 5
%type <genericValue> bitwise_or            // 2 | 5

// slice nodes
%type <genericSlice> list
%type <genericSlice> call_args
%type <genericSlice> block_args
%type <genericSlice> elsif_block
%type <genericSlice> function_args
%type <genericSlice> capture_list
%type <genericSlice> key_value_pairs
%type <genericSlice> nodes_with_commas
%type <genericSlice> comma_delimited_refs
%type <genericSlice> symbol_key_value_pairs
%type <genericSlice> nonempty_nodes_with_commas
%type <genericSlice> nodes_with_commas_and_optional_block
%type <stringSlice> namespaced_modules

%%

capture_list : /* empty */
  { Statements = []ast.Node{} }
| EOF
  { Statements = []ast.Node{} }
| capture_list line
  {
		if $2 != nil {
			Statements = append(Statements, $2)
		}
	}
| capture_list line EOF
  {
		if $2 != nil {
			Statements = append(Statements, $2)
		}
	}

line : NEWLINE { $$ = nil }
| SEMICOLON { $$ = nil }
| whitespace expr whitespace { $$ = $2 };

list : /* empty */
  { $$ = []ast.Node{} }
| list NEWLINE
  { /* do nothing */ }
| list WHITESPACE
  { /* do nothing */ }
| list expr
  {
		if $2 != nil {
			$$ = append($$, $2)
		}
	};

whitespace : /* zero or more */ | WHITESPACE whitespace
optional_newline : /* empty */ | optional_newline NEWLINE;

expr : NODE | REF | CAPITAL_REF | instance_variable | class_variable | global | callexpr | func_declaration | class_declaration | module_declaration | assignment | true | false | negation | complement | positive | negative | binary_addition | binary_subtraction | binary_multiplication | binary_division | bitwise_and | bitwise_or | array | hash | filename_const_reference | if_block;

callexpr : REF whitespace call_args
  {
    $$ = ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: $3,
    }
  }
| NODE DOT REF whitespace call_args
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: $3.(ast.BareReference),
      Args: $5,
    };
  }
| NODE DOT REF
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: $3.(ast.BareReference),
    };
  }
| REF DOT REF whitespace call_args
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: $3.(ast.BareReference),
      Args: $5,
    };
  }
| REF whitespace nodes_with_commas_and_optional_block
  {
    $$ = ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: $3,
    };
  }
| expr whitespace OPERATOR whitespace expr
  {
    $$ = ast.CallExpression{
      Func: ast.BareReference{Name: $3},
      Target: $1,
      Args: []ast.Node{$5},
    }
  };

call_args : LPAREN whitespace nodes_with_commas whitespace RPAREN
  { $$ = $3; }
| nonempty_nodes_with_commas
  { $$ = $1; };

nodes_with_commas_and_optional_block : expr
  { $$ = append($$, $1); }
| block
  { $$ = append($$, $1); }
| nodes_with_commas_and_optional_block whitespace COMMA whitespace expr
  { $$ = append($$, $5); }
| nodes_with_commas_and_optional_block whitespace COMMA whitespace block
  { $$ = append($$, $5); };

nonempty_nodes_with_commas : REF
  { $$ = append($$, $1); }
| NODE
  { $$ = append($$, $1); }
| nonempty_nodes_with_commas whitespace COMMA whitespace NODE
  { $$ = append($$, $5); };
| nonempty_nodes_with_commas whitespace COMMA whitespace REF
  { $$ = append($$, $5); }
| nonempty_nodes_with_commas whitespace COMMA whitespace block
  { $$ = append($$, $5); };

nodes_with_commas : /* empty */ { $$ = ast.Nodes{} }
| REF
  { $$ = append($$, $1); }
| NODE
  { $$ = append($$, $1); }
| nodes_with_commas whitespace COMMA whitespace NODE
  { $$ = append($$, $5); };
| nodes_with_commas whitespace COMMA whitespace REF
  { $$ = append($$, $5); };

whitespace_and_newlines : /* empty */
  | whitespace_and_newlines WHITESPACE;
  | whitespace_and_newlines NEWLINE;

// FIXME: this should use a different type than call_args
// call args can be a list of expressions. This is just a list of REFs or NODEs
func_declaration :
  DEF whitespace REF whitespace function_args whitespace_and_newlines list whitespace_and_newlines END
  {
		$$ = ast.FuncDecl{
			Name: $3.(ast.BareReference),
      Args: $5,
			Body: $7,
    }
  };

function_args : /* possibly nothing */ { $$ = []ast.Node{} };
| call_args { $$ = $1 };

class_declaration : CLASS whitespace class_name_with_modules NEWLINE list END
  {
    $$ = ast.ClassDecl{
       Name: $3.(ast.Class).Name,
       Body: $5,
    }
  }
| CLASS whitespace class_name_with_modules LESSTHAN class_name_with_modules NEWLINE list END
  {
    $$ = ast.ClassDecl{
       Name: $3.(ast.Class).Name,
       SuperClass: $5.(ast.Class),
       Namespace: $3.(ast.Class).Namespace,
       Body: $7,
    }
  };

module_declaration : MODULE whitespace class_name_with_modules NEWLINE list END
  {
    $$ = ast.ModuleDecl{
      Name: $3.(ast.Class).Name,
      Namespace: $3.(ast.Class).Namespace,
      Body: $5,
    }
  };

class_name_with_modules : whitespace CAPITAL_REF whitespace
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

assignment : REF whitespace EQUALTO whitespace expr
  {
    $$ = ast.Assignment{
      LHS: $1,
      RHS: $5,
    }
  }
|  CAPITAL_REF whitespace EQUALTO whitespace expr
  {
    $$ = ast.Assignment{
      LHS: $1,
      RHS: $5,
    }
  }
|  instance_variable whitespace EQUALTO whitespace expr
  {
    $$ = ast.Assignment{
      LHS: $1,
      RHS: $5,
    }
  }
|  class_variable whitespace EQUALTO whitespace expr
  {
    $$ = ast.Assignment{
      LHS: $1,
      RHS: $5,
    }
  }
|  global whitespace EQUALTO whitespace expr
  {
    $$ = ast.Assignment{
      LHS: $1,
      RHS: $5,
    }
  };

negation : BANG whitespace expr { $$ = ast.Negation{Target: $3} };
complement : COMPLEMENT whitespace expr { $$ = ast.Complement{Target: $3} };
positive : POSITIVE whitespace expr { $$ = ast.Positive{Target: $3} };
negative : NEGATIVE whitespace expr { $$ = ast.Negative{Target: $3} };

binary_addition : expr whitespace POSITIVE whitespace expr
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "+"},
      Args: []ast.Node{$5},
    }
  };

binary_subtraction : expr whitespace NEGATIVE whitespace expr
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "-"},
      Args: []ast.Node{$5},
    }
  };

binary_multiplication : expr whitespace STAR whitespace expr
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "*"},
      Args: []ast.Node{$5},
    }
  };

binary_division : expr whitespace SLASH whitespace expr
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "/"},
      Args: []ast.Node{$5},
    }
  };

bitwise_and: expr whitespace AMPERSAND whitespace expr
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "&"},
      Args: []ast.Node{$5},
    }
  };

bitwise_or: expr whitespace PIPE whitespace expr
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "|"},
      Args: []ast.Node{$5},
    }
  };

true : TRUE { $$ = ast.Boolean{Value: true} }
false : FALSE { $$ = ast.Boolean{Value: false} }
array : LBRACKET whitespace nodes_with_commas whitespace RBRACKET
  {
    $$ = ast.Array{Nodes: $3};
  };
hash : LBRACE whitespace optional_newline whitespace key_value_pairs whitespace optional_newline whitespace RBRACE
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

key_value_pairs : /* empty */ { $$ = ast.Nodes{} }
| expr whitespace EQUALTO GREATERTHAN whitespace expr
  { $$ = append($$, ast.HashKeyValuePair{Key: $1, Value: $6}) }
| key_value_pairs whitespace COMMA optional_newline whitespace expr whitespace EQUALTO GREATERTHAN whitespace expr
  { $$ = append($$, ast.HashKeyValuePair{Key: $6, Value: $11}) }
| key_value_pairs whitespace COMMA optional_newline whitespace expr whitespace EQUALTO GREATERTHAN whitespace expr COMMA
  { $$ = append($$, ast.HashKeyValuePair{Key: $6, Value: $11}) };

symbol_key_value_pairs : REF COLON whitespace expr
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


optional_comma : /* nothing */ { $$ = nil; }
| COMMA  { $$ = nil; };

global : DOLLARSIGN REF
  { $$ = ast.GlobalVariable{Name: $2.(ast.BareReference).Name} }
| DOLLARSIGN CAPITAL_REF
  { $$ = ast.GlobalVariable{Name: $2.(ast.BareReference).Name} };

instance_variable : ATSIGN REF
  { $$ = ast.InstanceVariable{Name: $2.(ast.BareReference).Name} }
| ATSIGN CAPITAL_REF
  { $$ = ast.InstanceVariable{Name: $2.(ast.BareReference).Name} };

class_variable : ATSIGN ATSIGN REF
  { $$ = ast.ClassVariable{Name: $3.(ast.BareReference).Name} }
| ATSIGN ATSIGN CAPITAL_REF
  { $$ = ast.ClassVariable{Name: $3.(ast.BareReference).Name} };

filename_const_reference : FILE_CONST_REF
  { $$ = ast.FileNameConstReference{} };

block : DO list END
  { $$ = ast.Block{Body: $2} }
| DO whitespace block_args whitespace one_or_more_newlines list END
  {
    $$ = ast.Block{
    Body: $6,
    Args: $3,
    }
  };

block_args : PIPE comma_delimited_refs PIPE
  { $$ = $2 };

comma_delimited_refs : /* empty */ { $$ = ast.Nodes{} }
| REF
  { $$ = append($$, $1); }
| nodes_with_commas whitespace COMMA whitespace REF
  { $$ = append($$, $5); };

one_or_more_newlines : NEWLINE | one_or_more_newlines NEWLINE

if_block : IF whitespace expr list END
  {
    $$ = ast.IfBlock{
      Condition: $3,
      Body: $4,
    }
  }
| IF whitespace expr list elsif_block END
  {
    $$ = ast.IfBlock{
      Condition: $3,
      Body: $4,
      Else: $5,
    }
  };

elsif_block : /* nothing */ { $$ = []ast.Node{} };
| elsif_block ELSIF whitespace expr list
  {
    $$ = append($$, ast.IfBlock{
      Condition: $4,
      Body: $5,
    })
  }
| elsif_block ELSE list
  {
    $$ = append($$, ast.IfBlock{
      Condition: ast.Boolean{Value: true},
      Body: $3,
    })
      }
| ELSIF whitespace expr list
  {
    $$ = append($$, ast.IfBlock{
      Condition: $3,
      Body: $4,
    })
  }
| ELSE list
  {
    $$ = append($$, ast.IfBlock{
      Condition: ast.Boolean{Value: true},
      Body: $2,
    })
  };

%%
