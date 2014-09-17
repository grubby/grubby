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
%token <genericValue> UNLESS
%token <genericValue> CLASS
%token <genericValue> MODULE
%token <genericValue> FOR
%token <genericValue> WHILE
%token <genericValue> UNTIL
%token <genericValue> BEGIN
%token <genericValue> RESCUE
%token <genericValue> ENSURE
%token <genericValue> BREAK
%token <genericValue> REDO
%token <genericValue> RETRY
%token <genericValue> RETURN

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

// single nodes
%type <genericValue> expr
%type <genericValue> line
%type <genericValue> true
%type <genericValue> hash
%type <genericValue> block
%type <genericValue> false
%type <genericValue> array
%type <genericValue> group
%type <genericValue> global
%type <genericValue> if_block
%type <genericValue> assignment
%type <genericValue> begin_block
%type <genericValue> single_node
%type <genericValue> class_variable
%type <genericValue> call_expression
%type <genericValue> func_declaration
%type <genericValue> binary_expression
%type <genericValue> class_declaration
%type <genericValue> instance_variable
%type <genericValue> module_declaration
%type <genericValue> class_name_with_modules
%type <genericValue> default_value_arg

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
%type <genericSlice> lines
%type <genericSlice> call_args
%type <genericSlice> block_args
%type <genericSlice> elsif_block
%type <genericSlice> capture_list
%type <genericSlice> function_args
%type <genericSlice> key_value_pairs
%type <genericSlice> optional_rescues
%type <genericSlice> nodes_with_commas
%type <genericSlice> comma_delimited_refs
%type <genericSlice> comma_delimited_nodes
%type <genericSlice> symbol_key_value_pairs
%type <genericSlice> nonempty_nodes_with_commas
%type <genericSlice> nodes_with_commas_and_optional_block
%type <genericSlice> comma_delimited_args_with_default_values
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
| optional_whitespace single_node optional_whitespace { $$ = $2 };
| optional_whitespace expr optional_whitespace { $$ = $2 }

list : /* empty */
  { $$ = []ast.Node{} }
| list NEWLINE
  { }
| list WHITESPACE
  { }
| list expr
  {
		if $2 != nil {
			$$ = append($$, $2)
		}
	};

optional_newline : /* empty */ | optional_newline NEWLINE;
one_or_more_newlines : NEWLINE | one_or_more_newlines NEWLINE;
optional_whitespace : /* zero or more */ | WHITESPACE optional_whitespace;

// e.g.: not a complex set of tokens (e.g.: call expression)
single_node : NODE | REF | CAPITAL_REF | instance_variable | class_variable | global | true | false | array | hash | class_name_with_modules;

binary_expression : binary_addition | binary_subtraction | binary_multiplication | binary_division | bitwise_and | bitwise_or;

expr : single_node | call_expression | func_declaration | class_declaration | module_declaration | assignment | negation | complement | positive | negative | if_block | group | begin_block | binary_expression;

call_expression : REF LPAREN optional_whitespace RPAREN
  {
    $$ = ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: []ast.Node{},
    }
  }
| REF LPAREN single_node RPAREN
  {
    $$ = ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: []ast.Node{$3},
    }
  }
| REF LPAREN nodes_with_commas RPAREN
  {
    $$ = ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: $3,
    }
  }
| REF LPAREN optional_whitespace nodes_with_commas optional_whitespace RPAREN
  {
    $$ = ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: $4,
    }
  }
| single_node optional_whitespace call_args
  {
    $$ = ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: $3,
    }
  }
| single_node DOT REF
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: $3.(ast.BareReference),
    };
  }
| single_node DOT REF optional_whitespace call_args
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: $3.(ast.BareReference),
      Args: $5,
    };
  }

// e.g.: `puts 'whatever' do ; end;` or with_a_block { puts 'foo' }
| REF optional_whitespace nodes_with_commas_and_optional_block
  {
    $$ = ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: $3,
    };
  }
| single_node optional_whitespace OPERATOR optional_whitespace single_node
  {
    $$ = ast.CallExpression{
      Func: ast.BareReference{Name: $3},
      Target: $1,
      Args: []ast.Node{$5},
    }
  }

// hash / array retrieval at index
| REF optional_whitespace LBRACKET optional_whitespace single_node optional_whitespace RBRACKET
  {
    $$ = ast.CallExpression{
      Func: ast.BareReference{Name: "[]"},
      Target: $1,
      Args: []ast.Node{$5},
    }
  }

// hash assignment
| REF optional_whitespace LBRACKET optional_whitespace single_node optional_whitespace RBRACKET optional_whitespace EQUALTO optional_whitespace expr
  {
    $$ = ast.CallExpression{
      Func: ast.BareReference{Name: "[]="},
      Target: $1,
      Args: []ast.Node{$5, $11},
    }
  };

call_args : LPAREN optional_whitespace nodes_with_commas optional_whitespace RPAREN
  { $$ = $3; }
| nonempty_nodes_with_commas
  { $$ = $1; };

comma_delimited_nodes : NODE
  { $$ = append($$, $1); }
| comma_delimited_nodes optional_whitespace COMMA optional_whitespace NODE
  { $$ = append($$, $5); };

nodes_with_commas : /* empty */ { $$ = ast.Nodes{} }
| single_node
  { $$ = append($$, $1); }
| binary_expression
  { $$ = append($$, $1); }
| call_expression
  { $$ = append($$, $1); }
| nodes_with_commas optional_whitespace COMMA optional_whitespace single_node
  { $$ = append($$, $5); };

// FIXME: this should ONLY have a block at the end (not in the middle)
nodes_with_commas_and_optional_block : expr
  { $$ = append($$, $1); }
| block
  { $$ = append($$, $1); }
| nodes_with_commas_and_optional_block optional_whitespace COMMA optional_whitespace expr
  { $$ = append($$, $5); }
| nodes_with_commas_and_optional_block optional_whitespace COMMA optional_whitespace block
  { $$ = append($$, $5); };

nonempty_nodes_with_commas : single_node
  { $$ = append($$, $1); }
| call_expression
  { $$ = append($$, $1); }
| nonempty_nodes_with_commas optional_whitespace COMMA optional_whitespace single_node
  { $$ = append($$, $5); }
| nonempty_nodes_with_commas optional_whitespace COMMA optional_whitespace call_expression
  { $$ = append($$, $5); }
| nonempty_nodes_with_commas optional_whitespace COMMA optional_whitespace block
  { $$ = append($$, $5); };

whitespace_and_newlines : /* empty */
  | whitespace_and_newlines WHITESPACE;
  | whitespace_and_newlines NEWLINE;

// FIXME: this should use a different type than call_args
// call args can be a list of expressions. This is just a list of REFs or NODEs
func_declaration :
  DEF optional_whitespace REF optional_whitespace function_args whitespace_and_newlines list whitespace_and_newlines END
  {
		$$ = ast.FuncDecl{
			Name: $3.(ast.BareReference),
      Args: $5,
			Body: $7,
    }
  };

function_args : comma_delimited_args_with_default_values
  { $$ = $1 }
| LPAREN optional_whitespace comma_delimited_args_with_default_values optional_whitespace RPAREN
  { $$ = $3 };

default_value_arg : REF optional_whitespace
  { $$ = ast.MethodParam{Name: $1.(ast.BareReference)} }
| REF optional_whitespace EQUALTO optional_whitespace expr
  { $$ = ast.MethodParam{Name: $1.(ast.BareReference), DefaultValue: $5} };

comma_delimited_args_with_default_values : /* empty */ { $$ = ast.Nodes{} }
| default_value_arg
  {
    $$ = append($$, $1)
  }
| comma_delimited_args_with_default_values optional_whitespace COMMA optional_whitespace default_value_arg
  {
    $$ = append($$, $5)
  };

class_declaration : CLASS optional_whitespace class_name_with_modules optional_whitespace NEWLINE list END
  {
    $$ = ast.ClassDecl{
       Name: $3.(ast.Class).Name,
       Body: $6,
    }
  }
| CLASS optional_whitespace class_name_with_modules optional_whitespace LESSTHAN optional_whitespace class_name_with_modules optional_whitespace NEWLINE list END
  {
    $$ = ast.ClassDecl{
       Name: $3.(ast.Class).Name,
       SuperClass: $7.(ast.Class),
       Namespace: $3.(ast.Class).Namespace,
       Body: $10,
    }
  };

module_declaration : MODULE optional_whitespace class_name_with_modules optional_whitespace NEWLINE list END
  {
    $$ = ast.ModuleDecl{
      Name: $3.(ast.Class).Name,
      Namespace: $3.(ast.Class).Namespace,
      Body: $6,
    }
  };

class_name_with_modules : CAPITAL_REF
  {
    $$ = ast.Class{
      Name: $1.(ast.BareReference).Name,
    }
  }
| namespaced_modules COLON COLON CAPITAL_REF
  {
    $$ = ast.Class{
       Name: $4.(ast.BareReference).Name,
       Namespace: strings.Join($1, "::"),
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

assignment : REF optional_whitespace EQUALTO optional_whitespace expr
  {
    $$ = ast.Assignment{
      LHS: $1,
      RHS: $5,
    }
  }
|  CAPITAL_REF optional_whitespace EQUALTO optional_whitespace expr
  {
    $$ = ast.Assignment{
      LHS: $1,
      RHS: $5,
    }
  }
|  instance_variable optional_whitespace EQUALTO optional_whitespace expr
  {
    $$ = ast.Assignment{
      LHS: $1,
      RHS: $5,
    }
  }
|  class_variable optional_whitespace EQUALTO optional_whitespace expr
  {
    $$ = ast.Assignment{
      LHS: $1,
      RHS: $5,
    }
  }
|  global optional_whitespace EQUALTO optional_whitespace expr
  {
    $$ = ast.Assignment{
      LHS: $1,
      RHS: $5,
    }
  };

negation : BANG optional_whitespace expr { $$ = ast.Negation{Target: $3} };
complement : COMPLEMENT optional_whitespace expr { $$ = ast.Complement{Target: $3} };
positive : POSITIVE optional_whitespace expr { $$ = ast.Positive{Target: $3} };
negative : NEGATIVE optional_whitespace expr { $$ = ast.Negative{Target: $3} };

binary_addition : single_node optional_whitespace POSITIVE optional_whitespace single_node
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "+"},
      Args: []ast.Node{$5},
    }
  }
| call_expression optional_whitespace POSITIVE optional_whitespace single_node
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "+"},
      Args: []ast.Node{$5},
    }
  }
| single_node optional_whitespace POSITIVE optional_whitespace call_expression
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "+"},
      Args: []ast.Node{$5},
    }
  }
| call_expression optional_whitespace POSITIVE optional_whitespace call_expression
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "+"},
      Args: []ast.Node{$5},
    }
  };

binary_subtraction : single_node optional_whitespace NEGATIVE optional_whitespace expr
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "-"},
      Args: []ast.Node{$5},
    }
  };

binary_multiplication : single_node optional_whitespace STAR optional_whitespace single_node
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "*"},
      Args: []ast.Node{$5},
    }
  };

binary_division : single_node optional_whitespace SLASH optional_whitespace single_node
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "/"},
      Args: []ast.Node{$5},
    }
  };

bitwise_and: single_node optional_whitespace AMPERSAND optional_whitespace single_node
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "&"},
      Args: []ast.Node{$5},
    }
  };

bitwise_or: single_node optional_whitespace PIPE optional_whitespace single_node
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "|"},
      Args: []ast.Node{$5},
    }
  };

true : TRUE { $$ = ast.Boolean{Value: true} }
false : FALSE { $$ = ast.Boolean{Value: false} }
array : LBRACKET optional_whitespace comma_delimited_nodes optional_whitespace RBRACKET
  { $$ = ast.Array{Nodes: $3} };
| LBRACKET optional_whitespace nodes_with_commas optional_whitespace RBRACKET
  { $$ = ast.Array{Nodes: $3} };

hash : LBRACE optional_whitespace optional_newline optional_whitespace key_value_pairs optional_whitespace optional_newline optional_whitespace RBRACE
  {
    pairs := []ast.HashKeyValuePair{}
    for _, node := range $5 {
      pairs = append(pairs, node.(ast.HashKeyValuePair))
    }
    $$ = ast.Hash{Pairs: pairs}
  }
| LBRACE optional_whitespace optional_newline optional_whitespace symbol_key_value_pairs optional_whitespace optional_newline optional_whitespace RBRACE
  {
    pairs := []ast.HashKeyValuePair{}
    for _, node := range $5 {
      pairs = append(pairs, node.(ast.HashKeyValuePair))
    }
    $$ = ast.Hash{Pairs: pairs}
  };

key_value_pairs : /* empty */ { $$ = ast.Nodes{} }
| single_node optional_whitespace OPERATOR optional_whitespace expr
  {
    if $3 != "=>" {
      panic("FREAKOUT")
    }
    $$ = append($$, ast.HashKeyValuePair{Key: $1, Value: $5})
  }
| key_value_pairs optional_whitespace COMMA optional_newline optional_whitespace expr optional_whitespace OPERATOR optional_whitespace expr
  {
  if $8 != "=>" {
      panic("FREAKOUT")
    }
    $$ = append($$, ast.HashKeyValuePair{Key: $6, Value: $10})
  }
| key_value_pairs optional_whitespace COMMA optional_newline optional_whitespace expr optional_whitespace OPERATOR optional_whitespace expr COMMA
  {
    if $8 != "=>" {
      panic("FREAKOUT")
    }
    $$ = append($$, ast.HashKeyValuePair{Key: $6, Value: $10})
  };

symbol_key_value_pairs : REF COLON optional_whitespace expr
  {
    $$ = append($$, ast.HashKeyValuePair{
      Key: ast.Symbol{Name: $1.(ast.BareReference).Name},
      Value: $4,
    })
  }
| symbol_key_value_pairs optional_whitespace COMMA optional_newline optional_whitespace REF COLON optional_whitespace expr
  {
    $$ = append($$, ast.HashKeyValuePair{
      Key: ast.Symbol{Name: $6.(ast.BareReference).Name},
      Value: $9,
    })
  }
| symbol_key_value_pairs optional_whitespace COMMA optional_newline optional_whitespace REF COLON optional_whitespace expr COMMA
  {
    $$ = append($$, ast.HashKeyValuePair{
      Key: ast.Symbol{Name: $6.(ast.BareReference).Name},
      Value: $9,
    })
  };

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

block : DO list END
  { $$ = ast.Block{Body: $2} }
| DO optional_whitespace block_args optional_whitespace list END
  {
    $$ = ast.Block{
    Body: $5,
    Args: $3,
    }
  }
| LBRACE optional_whitespace block_args list RBRACE
  {
    $$ = ast.Block{
      Body: $4,
      Args: $3,
    }
  }
| LBRACE list RBRACE
  { $$ = ast.Block{Body: $2} };

block_args : PIPE optional_whitespace comma_delimited_refs optional_whitespace PIPE
  { $$ = $3 };

comma_delimited_refs : /* empty */ { $$ = ast.Nodes{} }
| REF
  { $$ = append($$, $1); }
| comma_delimited_refs optional_whitespace COMMA optional_whitespace REF
  { $$ = append($$, $5); };

if_block : IF optional_whitespace expr list END
  {
    $$ = ast.IfBlock{
      Condition: $3,
      Body: $4,
    }
  }
| IF optional_whitespace expr list elsif_block END
  {
    $$ = ast.IfBlock{
      Condition: $3,
      Body: $4,
      Else: $5,
    }
  }
| single_node optional_whitespace IF optional_whitespace single_node
  {
    $$ = ast.IfBlock{
      Condition: $5,
      Body: []ast.Node{$1},
    }
  }
| call_expression optional_whitespace IF optional_whitespace single_node
  {
    $$ = ast.IfBlock{
      Condition: $5,
      Body: []ast.Node{$1},
    }
  }
| single_node optional_whitespace UNLESS optional_whitespace expr
  {
    $$ = ast.IfBlock{
      Condition: ast.Negation{Target: $5},
      Body: []ast.Node{$1},
    }
  };

elsif_block : /* nothing */ { $$ = []ast.Node{} };
| elsif_block ELSIF optional_whitespace expr list
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
| ELSIF optional_whitespace expr list
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

lines : /* empty */ { }
| SEMICOLON { };
| WHITESPACE { };
| expr { $$ = []ast.Node{$1} }
| lines expr { $$ = append($$, $2) }
| lines SEMICOLON { }
| lines optional_whitespace { };

group : LPAREN lines RPAREN
  { $$ = ast.Group{Body: $2} };

begin_block : BEGIN NEWLINE list optional_rescues END
  {
    $$ = ast.Begin{
      Body: $3,
      Rescue: $4,
    }
  };

optional_rescues : /* empty */
  { $$ = []ast.Node{} }
| optional_rescues optional_newline RESCUE list
  { $$ = append($$, ast.Rescue{Body: $4}) }
| optional_rescues optional_newline RESCUE optional_whitespace CAPITAL_REF optional_whitespace list
  {
    $$ = append($$, ast.Rescue{
      Body: $7,
      Exception: ast.RescueException{
        Class: $5.(ast.BareReference),
      },
    })
  }
| optional_rescues optional_newline RESCUE optional_whitespace CAPITAL_REF optional_whitespace OPERATOR optional_whitespace REF list
  {
    if $7 != "=>" {
      panic("FREAKOUT")
    }

    $$ = append($$, ast.Rescue{
      Body: $10,
      Exception: ast.RescueException{
        Var: $9.(ast.BareReference),
        Class: $5.(ast.BareReference),
      },
    })
  };

%%
