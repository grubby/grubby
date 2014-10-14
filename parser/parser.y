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
%token <genericValue> SPECIAL_CHAR_REF
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
%token <genericValue> NEXT
%token <genericValue> REDO
%token <genericValue> RETRY
%token <genericValue> RETURN
%token <genericValue> YIELD

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
%token <genericValue> DOUBLECOLON
%token <genericValue> DOT
%token <genericValue> PIPE          // "|"
%token <genericValue> SLASH         // "/"
%token <genericValue> AMPERSAND     // "&"
%token <genericValue> QUESTIONMARK  // "?"
%token <genericValue> CARET         // "^"
%token <genericValue> LBRACKET      // "["
%token <genericValue> RBRACKET      // "]"
%token <genericValue> LBRACE        // "{"
%token <genericValue> RBRACE        // "}"
%token <genericValue> DOLLARSIGN    // "$"
%token <genericValue> ATSIGN        // "@"
%token <genericValue> FILE_CONST_REF // __FILE__
%token <genericValue> EOF

/*
  eg: if you want to be able to assign to something in the RubySymType
      struct, or if you want a terminating node below, you will want to
      declare a type (or possibly just a token)
*/

// single nodes
%type <genericValue> expr
%type <genericValue> true
%type <genericValue> hash
%type <genericValue> block
%type <genericValue> false
%type <genericValue> array
%type <genericValue> group
%type <genericValue> global
%type <genericValue> ternary
%type <genericValue> if_block
%type <genericValue> assignment
%type <genericValue> while_loop
%type <genericValue> begin_block
%type <genericValue> single_node
%type <genericValue> class_variable
%type <genericValue> loop_statement
%type <genericValue> call_expression
%type <genericValue> func_declaration
%type <genericValue> yield_expression
%type <genericValue> return_expression
%type <genericValue> binary_expression
%type <genericValue> class_declaration
%type <genericValue> default_value_arg
%type <genericValue> instance_variable
%type <genericValue> module_declaration
%type <genericValue> class_name_with_modules

%type <genericValue> assignable_variables;

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

// misc
%type <genericValue> optional_newlines

%%

capture_list : /* empty */
  { Statements = []ast.Node{} }
| NEWLINE
  { }
| SEMICOLON
  { }
| EOF
  { }
| capture_list expr SEMICOLON
  { Statements = append(Statements, $2) }
| capture_list expr NEWLINE
  { Statements = append(Statements, $2) }
| capture_list expr EOF
  {
    Statements = append(Statements, $2)
	}
| capture_list NEWLINE
| capture_list SEMICOLON
| capture_list EOF
  { };

optional_newlines : /* empty */ { }
| optional_newlines NEWLINE { }

list : /* empty */
  { $$ = ast.Nodes{} }
| list NEWLINE
  {  }
| list SEMICOLON
  {  }
| list single_node
  {  $$ = append($$, $2) };
| list expr
  {  $$ = append($$, $2) };

// e.g.: not a complex set of tokens (e.g.: call expression)
single_node : NODE | REF | CAPITAL_REF | instance_variable | class_variable | global | true | false | array | hash | class_name_with_modules | ternary | call_expression | group;

binary_expression : binary_addition | binary_subtraction | binary_multiplication | binary_division | bitwise_and | bitwise_or;

expr : single_node | func_declaration | class_declaration | module_declaration | assignment | negation | complement | positive | negative | if_block | begin_block | binary_expression | yield_expression | return_expression | while_loop | loop_statement;

loop_statement: BREAK
  { $$ = ast.Break{} }
| NEXT
  { $$ = ast.Next{} };

call_expression : REF LPAREN nodes_with_commas RPAREN
  {
    $$ = ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: $3,
    }
  }
| SPECIAL_CHAR_REF
  {
    $$ = ast.CallExpression{Func: $1.(ast.BareReference)}
  }
| SPECIAL_CHAR_REF LPAREN nodes_with_commas RPAREN
  {
    $$ = ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: $3,
    }
  }
| CAPITAL_REF LPAREN nodes_with_commas RPAREN
  {
    $$ = ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: $3,
    }
  }
| REF call_args
  {
    $$ = ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: $2,
    }
  }
| single_node DOT REF
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: $3.(ast.BareReference),
    };
  }
| call_expression DOT REF
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: $3.(ast.BareReference),
    };
  }
| single_node DOT REF nodes_with_commas_and_optional_block
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: $3.(ast.BareReference),
      Args: $4,
    }
  }
| single_node DOT REF call_args
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: $3.(ast.BareReference),
      Args: $4,
    };
  }
| group DOT REF
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: $3.(ast.BareReference),
      Args: []ast.Node{},
    }
  }

// e.g.: `puts 'whatever' do ; end;` or with_a_block { puts 'foo' }
| REF nodes_with_commas_and_optional_block
  {
    $$ = ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: $2,
    };
  }
| single_node LESSTHAN single_node
  {
    $$ = ast.CallExpression{
      Func: ast.BareReference{Name: "<"},
      Target: $1,
      Args: []ast.Node{$3},
    }
  }
| call_expression LESSTHAN single_node
  {
    $$ = ast.CallExpression{
      Func: ast.BareReference{Name: "<"},
      Target: $1,
      Args: []ast.Node{$3},
    }
  }
| single_node GREATERTHAN single_node
  {
    $$ = ast.CallExpression{
      Func: ast.BareReference{Name: ">"},
      Target: $1,
      Args: []ast.Node{$3},
    }
  }
| call_expression GREATERTHAN single_node
  {
    $$ = ast.CallExpression{
      Func: ast.BareReference{Name: ">"},
      Target: $1,
      Args: []ast.Node{$3},
    }
  }
| single_node OPERATOR single_node
  {
    $$ = ast.CallExpression{
      Func: ast.BareReference{Name: $2},
      Target: $1,
      Args: []ast.Node{$3},
    }
  }

// hash / array retrieval at index
| REF LBRACKET single_node RBRACKET
  {
    $$ = ast.CallExpression{
      Func: ast.BareReference{Name: "[]"},
      Target: $1,
      Args: []ast.Node{$3},
    }
  }

// hash assignment
| REF LBRACKET single_node RBRACKET EQUALTO expr
  {
    $$ = ast.CallExpression{
      Func: ast.BareReference{Name: "[]="},
      Target: $1,
      Args: []ast.Node{$3, $6},
    }
  };

call_args : LPAREN nodes_with_commas RPAREN
  { $$ = $2; }
| nonempty_nodes_with_commas
  { $$ = $1; };

comma_delimited_nodes : single_node
  { $$ = append($$, $1); }
| comma_delimited_nodes COMMA single_node
  { $$ = append($$, $3); };

nodes_with_commas : /* empty */ { $$ = ast.Nodes{} }
| single_node
  { $$ = append($$, $1); }
| binary_expression
  { $$ = append($$, $1); }
| nodes_with_commas COMMA single_node
  { $$ = append($$, $3); };

// FIXME: this should ONLY have a block at the end (not in the middle)
nodes_with_commas_and_optional_block : single_node
  { $$ = append($$, $1); }
| block
  { $$ = append($$, $1); }
| nodes_with_commas_and_optional_block COMMA single_node
  { $$ = append($$, $3); }
| nodes_with_commas_and_optional_block COMMA block
  { $$ = append($$, $3); };

nonempty_nodes_with_commas : single_node
  { $$ = append($$, $1); }
| nonempty_nodes_with_commas COMMA single_node
  { $$ = append($$, $3); }
| nonempty_nodes_with_commas COMMA block
  { $$ = append($$, $3); };


// FIXME: this should use a different type than call_args
// call args can be a list of expressions. This is just a list of REFs or NODEs
func_declaration : DEF REF function_args optional_newlines list optional_newlines END
  {
		$$ = ast.FuncDecl{
			Name: $2.(ast.BareReference),
      Args: $3,
			Body: $5,
    }
  }
| DEF OPERATOR function_args optional_newlines list optional_newlines END
  {
		$$ = ast.FuncDecl{
			Name: ast.BareReference{Name: $2},
      Args: $3,
			Body: $5,
    }
  };

function_args : comma_delimited_args_with_default_values
  { $$ = $1 }
| LPAREN comma_delimited_args_with_default_values RPAREN
  { $$ = $2 };

default_value_arg : REF
  { $$ = ast.MethodParam{Name: $1.(ast.BareReference)} }
| STAR REF
  { $$ = ast.MethodParam{Name: $2.(ast.BareReference), IsSplat: true} }
| REF EQUALTO expr
  { $$ = ast.MethodParam{Name: $1.(ast.BareReference), DefaultValue: $3} }
| AMPERSAND REF
  { $$ = ast.MethodParam{Name: $2.(ast.BareReference), IsProc: true} };

comma_delimited_args_with_default_values : /* empty */ { $$ = ast.Nodes{} }
| default_value_arg
  {
    $$ = append($$, $1)
  }
| comma_delimited_args_with_default_values COMMA default_value_arg
  {
    $$ = append($$, $3)
  };

class_declaration : CLASS class_name_with_modules list END
  {
    $$ = ast.ClassDecl{
       Name: $2.(ast.Class).Name,
       Body: $3,
    }
  }
| CLASS class_name_with_modules LESSTHAN class_name_with_modules list END
  {
    $$ = ast.ClassDecl{
       Name: $2.(ast.Class).Name,
       SuperClass: $4.(ast.Class),
       Namespace: $2.(ast.Class).Namespace,
       Body: $5,
    }
  };

module_declaration : MODULE class_name_with_modules list END
  {
    $$ = ast.ModuleDecl{
      Name: $2.(ast.Class).Name,
      Namespace: $2.(ast.Class).Namespace,
      Body: $3,
    }
  };

class_name_with_modules : CAPITAL_REF
  {
    $$ = ast.Class{
      Name: $1.(ast.BareReference).Name,
    }
  }
| namespaced_modules DOUBLECOLON CAPITAL_REF
  {
    $$ = ast.Class{
       Name: $3.(ast.BareReference).Name,
       Namespace: strings.Join($1, "::"),
    }
  };

namespaced_modules : CAPITAL_REF
  {
    $$ = append($$, $1.(ast.BareReference).Name)
  }
|  namespaced_modules DOUBLECOLON CAPITAL_REF
  {
    $$ = append($$, $3.(ast.BareReference).Name)
  };

assignment : REF EQUALTO single_node
  {
    $$ = ast.Assignment{
      LHS: $1,
      RHS: $3,
    }
  }
| CAPITAL_REF EQUALTO expr
  {
    $$ = ast.Assignment{
      LHS: $1,
      RHS: $3,
    }
  }
| instance_variable EQUALTO expr
  {
    $$ = ast.Assignment{
      LHS: $1,
      RHS: $3,
    }
  }
| class_variable EQUALTO expr
  {
    $$ = ast.Assignment{
      LHS: $1,
      RHS: $3,
    }
  }
| global EQUALTO expr
  {
    $$ = ast.Assignment{
      LHS: $1,
      RHS: $3,
    }
  }
| assignable_variables EQUALTO expr
  {
    $$ = ast.Assignment{
      LHS: $1,
      RHS: $3,
    }
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

assignable_variables : REF COMMA REF
  { $$ = ast.Array{Nodes: []ast.Node{$1, $3}} }
| assignable_variables COMMA REF
  { $$ = ast.Array{Nodes: append($$.(ast.Array).Nodes, $3)} }

negation : BANG expr { $$ = ast.Negation{Target: $2} };
complement : COMPLEMENT expr { $$ = ast.Complement{Target: $2} };
positive : POSITIVE expr { $$ = ast.Positive{Target: $2} };
negative : NEGATIVE expr { $$ = ast.Negative{Target: $2} };

binary_addition : single_node POSITIVE single_node
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "+"},
      Args: []ast.Node{$3},
    }
  };

binary_subtraction : single_node NEGATIVE expr
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "-"},
      Args: []ast.Node{$3},
    }
  };

binary_multiplication : single_node STAR single_node
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "*"},
      Args: []ast.Node{$3},
    }
  };

binary_division : single_node SLASH single_node
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "/"},
      Args: []ast.Node{$3},
    }
  };

bitwise_and: single_node AMPERSAND single_node
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "&"},
      Args: []ast.Node{$3},
    }
  };

bitwise_or: single_node PIPE single_node
  {
    $$ = ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Name: "|"},
      Args: []ast.Node{$3},
    }
  };

true : TRUE { $$ = ast.Boolean{Value: true} }
false : FALSE { $$ = ast.Boolean{Value: false} }
array : LBRACKET comma_delimited_nodes RBRACKET
  { $$ = ast.Array{Nodes: $2} };
| LBRACKET nodes_with_commas RBRACKET
  { $$ = ast.Array{Nodes: $2} };

hash : LBRACE optional_newlines key_value_pairs optional_newlines RBRACE
  {
    pairs := []ast.HashKeyValuePair{}
    for _, node := range $3 {
      pairs = append(pairs, node.(ast.HashKeyValuePair))
    }
    $$ = ast.Hash{Pairs: pairs}
  }
| LBRACE optional_newlines symbol_key_value_pairs optional_newlines RBRACE
  {
    pairs := []ast.HashKeyValuePair{}
    for _, node := range $3 {
      pairs = append(pairs, node.(ast.HashKeyValuePair))
    }
    $$ = ast.Hash{Pairs: pairs}
  };

key_value_pairs : /* empty */ { $$ = ast.Nodes{} }
| single_node OPERATOR expr optional_newlines
  {
    if $2 != "=>" {
      panic("FREAKOUT")
    }
    $$ = append($$, ast.HashKeyValuePair{Key: $1, Value: $3})
  }
| key_value_pairs optional_newlines COMMA expr OPERATOR expr optional_newlines
  {
  if $5 != "=>" {
      panic("FREAKOUT")
    }
    $$ = append($$, ast.HashKeyValuePair{Key: $4, Value: $6})
  }
| key_value_pairs optional_newlines COMMA expr OPERATOR expr COMMA optional_newlines
  {
    if $5 != "=>" {
      panic("FREAKOUT")
    }
    $$ = append($$, ast.HashKeyValuePair{Key: $4, Value: $6})
  };

symbol_key_value_pairs : REF COLON single_node
  {
    $$ = append($$, ast.HashKeyValuePair{
      Key: ast.Symbol{Name: $1.(ast.BareReference).Name},
      Value: $3,
    })
  }
| symbol_key_value_pairs COMMA optional_newlines REF COLON single_node optional_newlines
  {
    $$ = append($$, ast.HashKeyValuePair{
      Key: ast.Symbol{Name: $4.(ast.BareReference).Name},
      Value: $6,
    })
  }
| symbol_key_value_pairs COMMA optional_newlines REF COLON single_node COMMA optional_newlines
  {
    $$ = append($$, ast.HashKeyValuePair{
      Key: ast.Symbol{Name: $4.(ast.BareReference).Name},
      Value: $6,
    })
  };

block : DO list END
  { $$ = ast.Block{Body: $2} }
| DO block_args list END
  {
    $$ = ast.Block{
    Body: $3,
    Args: $2,
    }
  }
| LBRACE block_args list RBRACE
  {
    $$ = ast.Block{
      Body: $3,
      Args: $2,
    }
  }
| LBRACE list RBRACE
  { $$ = ast.Block{Body: $2} };

block_args : PIPE comma_delimited_refs PIPE
  { $$ = $2 };

comma_delimited_refs : /* empty */ { $$ = ast.Nodes{} }
| REF
  { $$ = append($$, $1); }
| comma_delimited_refs COMMA REF
  { $$ = append($$, $3); };

if_block : IF expr list END
  {
    $$ = ast.IfBlock{
      Condition: $2,
      Body: $3,
    }
  }
| IF expr list elsif_block END
  {
    $$ = ast.IfBlock{
      Condition: $2,
      Body: $3,
      Else: $4,
    }
  }
| expr IF expr
  {
    $$ = ast.IfBlock{
      Condition: $3,
      Body: []ast.Node{$1},
    }
  }
| call_expression IF expr
  {
    $$ = ast.IfBlock{
      Condition: $3,
      Body: []ast.Node{$1},
    }
  }
| single_node UNLESS expr
  {
    $$ = ast.IfBlock{
      Condition: ast.Negation{Target: $3},
      Body: []ast.Node{$1},
    }
  }
| call_expression UNLESS expr
  {
    $$ = ast.IfBlock{
      Condition: ast.Negation{Target: $3},
      Body: ast.Nodes{$1},
    }
  }
| assignment UNLESS expr
  {
    $$ = ast.IfBlock{
      Condition: ast.Negation{Target: $3},
      Body: ast.Nodes{$1},
    }
  }
| UNLESS expr NEWLINE list END
  {
    $$ = ast.IfBlock{
      Condition: ast.Negation{Target: $2},
      Body: $4,
    }
  }
| UNLESS expr NEWLINE list elsif_block END
  {
    $$ = ast.IfBlock{
      Condition: ast.Negation{Target: $2},
      Body: $4,
      Else: $5,
    }
  }
| UNLESS expr SEMICOLON list END
  {
    $$ = ast.IfBlock{
      Condition: ast.Negation{Target: $2},
      Body: $4,
    }
  };


elsif_block : elsif_block ELSIF expr list
  {
    $$ = append($$, ast.IfBlock{
      Condition: $3,
      Body: $4,
    })
  }
| elsif_block ELSE list
  {
    $$ = append($$, ast.IfBlock{
      Condition: ast.Boolean{Value: true},
      Body: $3,
    })
      }
| ELSIF expr list
  {
    $$ = append($$, ast.IfBlock{
      Condition: $2,
      Body: $3,
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
| lines expr { $$ = append($$, $2) }
| lines SEMICOLON { };

group : LPAREN lines RPAREN
  { $$ = ast.Group{Body: $2} };

begin_block : BEGIN list optional_rescues END
  {
    $$ = ast.Begin{
      Body: $2,
      Rescue: $3,
    }
  };

optional_rescues : /* empty */
  { $$ = []ast.Node{} }
| optional_rescues RESCUE list
  { $$ = append($$, ast.Rescue{Body: $3}) }
| optional_rescues RESCUE CAPITAL_REF list
  {
    $$ = append($$, ast.Rescue{
      Body: $4,
      Exception: ast.RescueException{
        Class: $3.(ast.BareReference),
      },
    })
  }
| optional_rescues RESCUE CAPITAL_REF OPERATOR REF list
  {
    if $4 != "=>" {
      panic("FREAKOUT")
    }

    $$ = append($$, ast.Rescue{
      Body: $6,
      Exception: ast.RescueException{
        Var: $5.(ast.BareReference),
        Class: $3.(ast.BareReference),
      },
    })
  };

yield_expression : YIELD comma_delimited_nodes
  {
    if len($2) == 1 {
      $$ = ast.Yield{Value: $2[0]}
    } else {
      $$ = ast.Yield{Value: $2}
    }
  };
return_expression : RETURN comma_delimited_nodes
  {
    if len($2) == 1 {
      $$ = ast.Return{Value: $2[0]}
    } else {
      $$ = ast.Return{Value: $2}
    }
  };

ternary : single_node QUESTIONMARK single_node COLON single_node
  {
    $$ = ast.Ternary{
      Condition: $1,
      True: $3,
      False: $5,
    }
  }

while_loop : WHILE expr NEWLINE list END
  { $$ = ast.Loop{Condition: $2, Body: $4} };

%%
