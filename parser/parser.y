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
  genericBlock    ast.Block
  genericValue    ast.Node
  genericSlice    ast.Nodes
  genericString   string
  stringSlice     []string
  switchCaseSlice []ast.SwitchCase
}

%token <genericValue> OPERATOR

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%token <genericValue> NODE
%token <genericValue> REF
%token <genericValue> SYMBOL
%token <genericValue> SPECIAL_CHAR_REF
%token <genericValue> CAPITAL_REF
%token <genericValue> NAMESPACED_CAPITAL_REF
%token <genericValue> LPAREN
%token <genericValue> RPAREN
%token <genericValue> COMMA

%token <genericString> NamespacedModule
%token <genericValue> ProcArg

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
%token <genericValue> AND
%token <genericValue> OR
%token <genericValue> LAMBDA
%token <genericValue> CASE
%token <genericValue> WHEN
%token <genericValue> ALIAS
%token <genericValue> SELF
%token <genericValue> NIL

// operators
%token <genericValue> LESSTHAN
%token <genericValue> GREATERTHAN
%token <genericValue> EQUALTO
%token <genericValue> BANG
%token <genericValue> COMPLEMENT

%token <genericValue> BINARY_PLUS
%token <genericValue> UNARY_PLUS

%token <genericValue> BINARY_MINUS
%token <genericValue> UNARY_MINUS

%token <genericValue> STAR
%token <genericValue> RANGE

%token <genericValue> OR_EQUALS

// misc
%token <genericValue> WHITESPACE
%token <genericValue> NEWLINE
%token <genericValue> SEMICOLON
%token <genericValue> COLON
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
%token <genericValue> LINE_CONST_REF // __LINE__
%token <genericValue> EOF

/*
  eg: if you want to be able to assign to something in the RubySymType
      struct, or if you want a terminating node below, you will want to
      declare a type (or possibly just a token)
*/

 // single nodes
%type <genericValue> nil
%type <genericValue> expr
%type <genericValue> self
%type <genericValue> hash
%type <genericValue> range
%type <genericBlock> block
%type <genericValue> alias
%type <genericValue> array
%type <genericValue> group
%type <genericValue> global
%type <genericValue> lambda
%type <genericValue> rescue
%type <genericValue> ternary
%type <genericValue> if_block
%type <genericValue> proc_arg
%type <genericValue> splat_arg
%type <genericValue> assignment
%type <genericValue> multiple_assignment
%type <genericValue> begin_block
%type <genericValue> single_node
%type <genericValue> simple_node
%type <genericValue> class_variable
%type <genericValue> call_expression
%type <genericValue> operator_expression
%type <genericValue> method_declaration
%type <genericValue> yield_expression
%type <genericValue> retry_expression
%type <genericValue> return_expression
%type <genericValue> break_expression
%type <genericValue> next_expression
%type <genericValue> binary_expression
%type <genericValue> class_declaration
%type <genericValue> eigenclass_declaration
%type <genericValue> default_value_arg
%type <genericValue> instance_variable
%type <genericValue> module_declaration
%type <genericValue> conditional_assignment
%type <genericValue> class_name_with_modules

%type <switchCaseSlice> switch_cases;
%type <genericValue> switch_statement;

%type <genericValue> logical_or;
%type <genericValue> logical_and;

%type <genericValue> rescue_modifier;

// loops and expressions that can be inside a loop
%type <genericValue> while_loop
%type <genericValue> loop_if_block;
%type <genericSlice> loop_elsif_block;

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
%type <genericSlice> rescues
%type <genericSlice> call_args
%type <genericSlice> block_args
%type <genericSlice> elsif_block
%type <genericSlice> capture_list
%type <genericSlice> method_args
%type <genericSlice> key_value_pairs
%type <genericSlice> loop_expressions
%type <genericSlice> optional_rescues
%type <genericSlice> nodes_with_commas
%type <genericSlice> comma_delimited_refs
%type <genericSlice> comma_delimited_nodes
%type <genericSlice> symbol_key_value_pairs
%type <genericSlice> nonempty_nodes_with_commas
%type <genericSlice> two_or_more_call_expressions
%type <genericSlice> comma_delimited_args_with_default_values
%type <genericSlice> comma_delimited_class_names
%type <genericSlice> nodes_with_commas_and_optional_newlines

// misc
%type <genericValue> optional_comma
%type <genericValue> optional_newlines

%left DOT
%left QUESTIONMARK

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

optional_comma : /* empty */ { }
| COMMA { };

optional_newlines : /* empty */ { }
| optional_newlines NEWLINE { }

list : /* empty */
  { $$ = ast.Nodes{} }
| list NEWLINE
  {  }
| list SEMICOLON
  {  }
| list expr
{  $$ = append($$, $2) };

simple_node : SYMBOL | NODE | REF | CAPITAL_REF | instance_variable | class_variable | global | LINE_CONST_REF | FILE_CONST_REF | self | nil;

single_node : simple_node | array | hash | class_name_with_modules | call_expression | operator_expression | group | lambda | negation | complement | positive | negative | splat_arg | logical_and | logical_or | binary_expression;

binary_expression : binary_addition | binary_subtraction | binary_multiplication | binary_division | bitwise_and | bitwise_or;

expr : single_node | method_declaration | class_declaration | module_declaration | eigenclass_declaration | assignment | multiple_assignment | conditional_assignment | if_block | begin_block | yield_expression | while_loop | switch_statement | return_expression | break_expression | next_expression | rescue_modifier | range | retry_expression | ternary | alias;

rescue_modifier : single_node RESCUE single_node
  { $$ = ast.RescueModifier{Statement: $1, Rescue: $3} };

splat_arg : STAR single_node
  { $$ = ast.StarSplat{Value: $2} };

call_expression : REF LPAREN nodes_with_commas RPAREN
  {
    callExpr := ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: $3,
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| REF LPAREN nodes_with_commas RPAREN block
  {
    callExpr := ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: $3,
      OptionalBlock: $5,
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| SPECIAL_CHAR_REF
  {
    callExpr := ast.CallExpression{Func: $1.(ast.BareReference)}
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| SPECIAL_CHAR_REF LPAREN nodes_with_commas RPAREN
  {
    callExpr := ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: $3,
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| CAPITAL_REF LPAREN nodes_with_commas RPAREN
  {
    callExpr := ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: $3,
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| REF call_args
  {
    callExpr := ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: $2,
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| REF call_args block
  {
    callExpr := ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: $2,
      OptionalBlock: $3,
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| REF block
  {
    callExpr := ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: []ast.Node{},
      OptionalBlock: $2,
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| single_node DOT REF
  {
    callExpr := ast.CallExpression{
      Target: $1,
      Func: $3.(ast.BareReference),
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| single_node DOT REF block
  {
    callExpr := ast.CallExpression{
      Target: $1,
      Func: $3.(ast.BareReference),
      Args: []ast.Node{},
      OptionalBlock: $4,
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| single_node DOT REF call_args block
  {
    callExpr := ast.CallExpression{
      Target: $1,
      Func: $3.(ast.BareReference),
      Args: $4,
      OptionalBlock: $5,
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| single_node DOT REF call_args
  {
    callExpr := ast.CallExpression{
      Target: $1,
      Func: $3.(ast.BareReference),
      Args: $4,
    };
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| group DOT REF
  {
    callExpr := ast.CallExpression{
      Target: $1,
      Func: $3.(ast.BareReference),
      Args: []ast.Node{},
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| group DOT REF block
  {
    callExpr := ast.CallExpression{
      Target: $1,
      Func: $3.(ast.BareReference),
      Args: []ast.Node{},
      OptionalBlock: $4,
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| single_node DOT REF EQUALTO expr
  {
    methodName := $3.(ast.BareReference).Name + "="
    callExpr := ast.CallExpression{
      Func: ast.BareReference{Name: methodName},
      Target: $1,
      Args: []ast.Node{$5},
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }

// e.g.: `puts 'whatever' do ; end;` or with_a_block { puts 'foo' }
| REF nodes_with_commas
  {
    callExpr := ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: $2,
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| REF nodes_with_commas block
  {
    callExpr := ast.CallExpression{
      Func: $1.(ast.BareReference),
      Args: $2,
      OptionalBlock: $3,
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| single_node LESSTHAN single_node
  {
    callExpr := ast.CallExpression{
      Func: ast.BareReference{Line: $2.LineNumber(), Name: "<"},
      Target: $1,
      Args: []ast.Node{$3},
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| call_expression LESSTHAN single_node
  {
    callExpr := ast.CallExpression{
      Func: ast.BareReference{Line: $2.LineNumber(), Name: "<"},
      Target: $1,
      Args: []ast.Node{$3},
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| single_node GREATERTHAN single_node
  {
    callExpr := ast.CallExpression{
      Func: ast.BareReference{Line: $2.LineNumber(), Name: ">"},
      Target: $1,
      Args: []ast.Node{$3},
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }

// hash / array retrieval at index
| REF LBRACKET single_node RBRACKET
  {
    callExpr := ast.CallExpression{
      Func: ast.BareReference{Line: $2.LineNumber(), Name: "[]"},
      Target: $1,
      Args: []ast.Node{$3},
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| CAPITAL_REF LBRACKET single_node RBRACKET
  {
    callExpr := ast.CallExpression{
      Func: ast.BareReference{Line: $2.LineNumber(), Name: "[]"},
      Target: $1,
      Args: []ast.Node{$3},
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| instance_variable LBRACKET single_node RBRACKET
  {
    callExpr := ast.CallExpression{
      Func: ast.BareReference{Line: $2.LineNumber(), Name: "[]"},
      Target: $1,
      Args: []ast.Node{$3},
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| REF LBRACKET range RBRACKET
  {
    callExpr := ast.CallExpression{
      Func: ast.BareReference{Line: $2.LineNumber(), Name: "[]"},
      Target: $1,
      Args: []ast.Node{$3},
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| CAPITAL_REF LBRACKET range RBRACKET
  {
    callExpr := ast.CallExpression{
      Func: ast.BareReference{Line: $2.LineNumber(), Name: "[]"},
      Target: $1,
      Args: []ast.Node{$3},
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| REF LBRACKET nonempty_nodes_with_commas RBRACKET
  {
    callExpr := ast.CallExpression{
      Func: ast.BareReference{Line: $2.LineNumber(), Name: "[]"},
      Target: $1,
      Args: $3,
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| CAPITAL_REF LBRACKET nonempty_nodes_with_commas RBRACKET
  {
    callExpr := ast.CallExpression{
      Func: ast.BareReference{Line: $2.LineNumber(), Name: "[]"},
      Target: $1,
      Args: $3,
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| call_expression LBRACKET nonempty_nodes_with_commas RBRACKET
  {
    callExpr := ast.CallExpression{
      Func: ast.BareReference{Line: $2.LineNumber(), Name: "[]"},
      Target: $1,
      Args: $3,
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| call_expression LBRACKET single_node RBRACKET
  {
    callExpr := ast.CallExpression{
      Func: ast.BareReference{Line: $3.LineNumber(), Name: "[]"},
      Target: $1,
      Args: []ast.Node{$3},
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }

// hash assignment
| REF LBRACKET single_node RBRACKET EQUALTO expr
  {
    callExpr := ast.CallExpression{
      Func: ast.BareReference{Line: $3.LineNumber(), Name: "[]="},
      Target: $1,
      Args: []ast.Node{$3, $6},
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| CAPITAL_REF LBRACKET single_node RBRACKET EQUALTO optional_newlines expr
{
  if $7 == nil {
      panic("WHAT THE EVER COMPILING FUCK")
    }
    callExpr := ast.CallExpression{
      Func: ast.BareReference{Line: $3.LineNumber(), Name: "[]="},
      Target: $1,
      Args: []ast.Node{$3, $7},
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| instance_variable LBRACKET single_node RBRACKET EQUALTO expr
  {
    callExpr := ast.CallExpression{
      Func: ast.BareReference{Line: $3.LineNumber(), Name: "[]="},
      Target: $1,
      Args: []ast.Node{$3, $6},
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  }
| call_expression LBRACKET single_node RBRACKET EQUALTO expr
  {
    callExpr := ast.CallExpression{
      Func: ast.BareReference{Line: $3.LineNumber(), Name: "[]="},
      Target: $1,
      Args: []ast.Node{$3, $6},
    }
    callExpr.Line = $1.LineNumber()
    $$ = callExpr
  };


operator_expression : single_node OPERATOR optional_newlines single_node
  {
    callExpr := ast.CallExpression{
      Line: $1.LineNumber(),
      Func: $2.(ast.BareReference),
      Target: $1,
      Args: []ast.Node{$4},
    }
    $$ = callExpr
  };


call_args : LPAREN nodes_with_commas RPAREN
  { $$ = $2 }
| LPAREN nodes_with_commas COMMA optional_newlines proc_arg RPAREN
  { $$ = append($2, $5) }
| nonempty_nodes_with_commas
  { $$ = $1 }
| nonempty_nodes_with_commas COMMA optional_newlines proc_arg
  { $$ = append($1, $4) };

comma_delimited_nodes : single_node
  { $$ = append($$, $1) }
| comma_delimited_nodes COMMA single_node
  { $$ = append($$, $3) };

nodes_with_commas : /* empty */ { $$ = ast.Nodes{} }
| single_node
  { $$ = append($$, $1) }
| assignment
  { $$ = append($$, $1) }
| proc_arg
  { $$ = append($$, $1) }
| nodes_with_commas COMMA optional_newlines single_node
  { $$ = append($$, $4) }
| nodes_with_commas COMMA optional_newlines assignment
  { $$ = append($$, $4) }
| nodes_with_commas COMMA optional_newlines proc_arg
  { $$ = append($$, $4) };

proc_arg : ProcArg single_node
  {
    callExpr := ast.CallExpression{
      Line: $2.LineNumber(),
      Func: ast.BareReference{Line: $2.LineNumber(), Name: "to_proc"},
      Target: $2,
    }
    $$ = callExpr
  };

nonempty_nodes_with_commas : single_node
  { $$ = append($$, $1); }
| nonempty_nodes_with_commas COMMA single_node
  { $$ = append($$, $3); }


method_declaration : DEF REF method_args list END
  {
		method := ast.FuncDecl{
			Name: $2.(ast.BareReference),
      Args: $3,
			Body: $4,
    }
    method.Line = $2.LineNumber()
    $$ = method
  }
| DEF REF method_args list rescues END
  {
		method := ast.FuncDecl{
			Name: $2.(ast.BareReference),
      Args: $3,
			Body: $4,
      Rescues: $5,
    }
    method.Line = $2.LineNumber()
    $$ = method
  }
| DEF REF DOT REF method_args list END
  {
		method := ast.FuncDecl{
      Target: $2,
			Name: $4.(ast.BareReference),
      Args: $5,
			Body: $6,
    }
    method.Line = $2.LineNumber()
    $$ = method
  }
| DEF self DOT REF method_args list END
  {
		method := ast.FuncDecl{
      Target: $2,
			Name: $4.(ast.BareReference),
      Args: $5,
			Body: $6,
    }
    method.Line = $2.LineNumber()
    $$ = method
  }
| DEF REF DOT REF method_args list rescues END
  {
		method := ast.FuncDecl{
      Target: $2,
			Name: $4.(ast.BareReference),
      Args: $5,
			Body: $6,
      Rescues: $7,
    }
    method.Line = $2.LineNumber()
    $$ = method
  }
| DEF self DOT REF method_args list rescues END
  {
		method := ast.FuncDecl{
      Target: $2,
			Name: $4.(ast.BareReference),
      Args: $5,
			Body: $6,
      Rescues: $7,
    }
    method.Line = $2.LineNumber()
    $$ = method
  }
| DEF OPERATOR method_args list END
  {
		method := ast.FuncDecl{
			Name: $2.(ast.BareReference),
      Args: $3,
      Body: $4,
    }
    method.Line = $2.LineNumber()
    $$ = method
  }
| DEF OPERATOR method_args list rescues END
  {
		method := ast.FuncDecl{
			Name: $2.(ast.BareReference),
      Args: $3,
      Body: $4,
      Rescues: $5,
    }
    method.Line = $2.LineNumber()
    $$ = method
  };


method_args : comma_delimited_args_with_default_values
  { $$ = $1 }
| LPAREN comma_delimited_args_with_default_values RPAREN
  { $$ = $2 };

comma_delimited_args_with_default_values : /* empty */ { $$ = ast.Nodes{} }
| default_value_arg
  { $$ = append($$, $1) }
| comma_delimited_args_with_default_values COMMA default_value_arg
  { $$ = append($$, $3) };

default_value_arg : REF
  { $$ = ast.MethodParam{Name: $1.(ast.BareReference)} }
| STAR REF
  { $$ = ast.MethodParam{Name: $2.(ast.BareReference), IsSplat: true} }
| REF EQUALTO single_node
  { $$ = ast.MethodParam{Name: $1.(ast.BareReference), DefaultValue: $3} }
| ProcArg REF
  { $$ = ast.MethodParam{Name: $2.(ast.BareReference), IsProc: true} };


class_declaration : CLASS class_name_with_modules list END
  {
    class := ast.ClassDecl{
       Name: $2.(ast.Class).Name,
       Namespace: $2.(ast.Class).Namespace,
       Body: $3,
    }
    class.Line = $2.LineNumber()
    $$ = class
  }
| CLASS class_name_with_modules LESSTHAN class_name_with_modules list END
  {
    class := ast.ClassDecl{
       Name: $2.(ast.Class).Name,
       SuperClass: $4.(ast.Class),
       Namespace: $2.(ast.Class).Namespace,
       Body: $5,
    }
    class.Line = $2.LineNumber()
    $$ = class
  };

eigenclass_declaration : CLASS OPERATOR single_node list END
  {
    name := $2.(ast.BareReference).Name
    if name != "<<" {
      panic("FREAKOUT :: impossible operator after 'class' keyword (" + name + ")")
    }

    class := ast.EigenClass{
      Target: $3,
      Body: $4,
    }
    class.Line = $2.LineNumber()
    $$ = class
  };

module_declaration : MODULE class_name_with_modules list END
  {
    module := ast.ModuleDecl{
      Name: $2.(ast.Class).Name,
      Namespace: $2.(ast.Class).Namespace,
      Body: $3,
    }
    module.Line = $2.LineNumber()
    $$ = module
  };

class_name_with_modules : CAPITAL_REF
  {
    class := ast.Class{
      Name: $1.(ast.BareReference).Name,
      IsGlobalNamespace: false,
    }
    class.Line = $1.LineNumber()
    $$ = class
  }
| CAPITAL_REF NAMESPACED_CAPITAL_REF
  {
    firstPart := $1.(ast.BareReference).Name
    fullName := strings.Join([]string{firstPart, $2.(ast.BareReference).Name}, "")
    pieces := strings.Split(fullName, "::")
    name := pieces[len(pieces)-1]
    var namespace []string
    if len(pieces) > 1 {
      namespace = pieces[0:len(pieces)-1]
    }

    class := ast.Class{
       Name: name,
       Namespace: strings.Join(namespace, "::"),
       IsGlobalNamespace: false,
    }
    class.Line = $1.LineNumber()
    $$ = class
  }
| NAMESPACED_CAPITAL_REF
  {
    pieces := strings.Split($1.(ast.BareReference).Name, "::")
    name := pieces[len(pieces)-1]
    var namespace []string
    if len(pieces) > 1 {
      namespace = pieces[0:len(pieces)-1]
    }

    $$ = ast.Class{
      Line: $1.LineNumber(),
      Name: strings.TrimPrefix(name, "::"),
      Namespace: strings.TrimPrefix(strings.Join(namespace, "::"), "::"),
      IsGlobalNamespace: true,
    }
  };


assignment : REF EQUALTO single_node
  {
    eql := ast.Assignment{
      LHS: $1,
      RHS: $3,
    }
    eql.Line = $1.LineNumber()
    $$ = eql
  }
| REF EQUALTO rescue_modifier
  {
    eql := ast.Assignment{LHS: $1, RHS: $3}
    eql.Line = $1.LineNumber()
    $$ = eql
  }
| REF EQUALTO ternary
  {
    eql := ast.Assignment{LHS: $1, RHS: $3}
    eql.Line = $1.LineNumber()
    $$ = eql
  }
| CAPITAL_REF EQUALTO expr
  {
    eql := ast.Assignment{
      LHS: $1,
      RHS: $3,
    }
    eql.Line = $1.LineNumber()
    $$ = eql
  }
| instance_variable EQUALTO expr
  {
    eql := ast.Assignment{
      LHS: $1,
      RHS: $3,
    }
    eql.Line = $1.LineNumber()
    $$ = eql
  }
| class_variable EQUALTO expr
  {
    eql := ast.Assignment{
      LHS: $1,
      RHS: $3,
    }
    eql.Line = $1.LineNumber()
    $$ = eql
  }
| global EQUALTO expr
  {
    eql := ast.Assignment{
      LHS: $1,
      RHS: $3,
    }
    eql.Line = $1.LineNumber()
    $$ = eql
  };

multiple_assignment : assignable_variables EQUALTO expr
  {
    eql := ast.Assignment{
      LHS: $1,
      RHS: $3,
    }
    eql.Line = $1.LineNumber()
    $$ = eql
  }
| two_or_more_call_expressions EQUALTO two_or_more_call_expressions
  {
    eql := ast.Assignment{
      LHS: ast.Array{Nodes: $1},
      RHS: ast.Array{Nodes: $3},
    }
    eql.Line = $1[0].(ast.CallExpression).Target.LineNumber()
    $$ = eql
  };

two_or_more_call_expressions : REF LBRACKET single_node RBRACKET COMMA REF LBRACKET single_node RBRACKET
  {
    $$ = []ast.Node{
      ast.CallExpression{
        Target: $1,
        Func: ast.BareReference{Line: $3.LineNumber(), Name: "[]="},
        Args: []ast.Node{$3},
      },
      ast.CallExpression{
        Target: $6,
        Func: ast.BareReference{Line: $3.LineNumber(), Name: "[]="},
        Args: []ast.Node{$8},
      },
    }
  }
| two_or_more_call_expressions COMMA REF LBRACKET single_node RBRACKET
  {
    tail := ast.CallExpression{Line: $3.LineNumber(), Target: $3, Func: ast.BareReference{Name: "[]="}, Args: []ast.Node{$5}}
    $$ = append($1, tail)
  };

conditional_assignment : REF OR_EQUALS single_node
  {
    eql := ast.ConditionalAssignment{
      LHS: $1,
      RHS: $3,
    }
    eql.Line = $1.LineNumber()
    $$ = eql
  }
| REF OR_EQUALS ternary
  {
    eql := ast.ConditionalAssignment{LHS: $1, RHS: $3}
    eql.Line = $1.LineNumber()
    $$ = eql
  }
| CAPITAL_REF OR_EQUALS expr
  {
    eql := ast.ConditionalAssignment{
      LHS: $1,
      RHS: $3,
    }
    eql.Line = $1.LineNumber()
    $$ = eql
  }
| instance_variable OR_EQUALS expr
  {
    eql := ast.ConditionalAssignment{
      LHS: $1,
      RHS: $3,
    }
    eql.Line = $1.LineNumber()
    $$ = eql
  }
| class_variable OR_EQUALS expr
  {
    eql := ast.ConditionalAssignment{
      LHS: $1,
      RHS: $3,
    }
    eql.Line = $1.LineNumber()
    $$ = eql
  }
| global OR_EQUALS expr
  {
    eql := ast.ConditionalAssignment{
      LHS: $1,
      RHS: $3,
    }
    eql.Line = $1.LineNumber()
    $$ = eql
  }
| call_expression OR_EQUALS expr
  {
    eql := ast.ConditionalAssignment{LHS: $1, RHS: $3}
    eql.Line = $1.LineNumber()
    $$ = eql
  };

global : DOLLARSIGN REF
  { $$ = ast.GlobalVariable{Name: $2.(ast.BareReference).Name} }
| DOLLARSIGN CAPITAL_REF
  { $$ = ast.GlobalVariable{Name: $2.(ast.BareReference).Name} };

instance_variable : ATSIGN REF
  {
    ivar := ast.InstanceVariable{Name: $2.(ast.BareReference).Name}
    ivar.Line = $2.LineNumber()
    $$ = ivar
  }
| ATSIGN CAPITAL_REF
  {
    ivar := ast.InstanceVariable{Name: $2.(ast.BareReference).Name}
    ivar.Line = $2.LineNumber()
    $$ = ivar
  };

class_variable : ATSIGN ATSIGN REF
  {
    classvar := ast.ClassVariable{Name: $3.(ast.BareReference).Name}
    classvar.Line = $3.LineNumber()
    $$ = classvar
  }
| ATSIGN ATSIGN CAPITAL_REF
  {
    classvar := ast.ClassVariable{Name: $3.(ast.BareReference).Name}
    classvar.Line = $3.LineNumber()
    $$ = classvar
  };

assignable_variables : REF COMMA REF
  { vars := ast.Array{Nodes: []ast.Node{$1, $3}}; vars.Line = $1.LineNumber(); $$ = vars }
| REF COMMA instance_variable
  { vars := ast.Array{Nodes: []ast.Node{$1, $3}}; vars.Line = $1.LineNumber(); $$ = vars }
| REF COMMA class_variable
  { vars := ast.Array{Nodes: []ast.Node{$1, $3}}; vars.Line = $1.LineNumber(); $$ = vars }
| REF COMMA STAR REF
  { vars := ast.Array{Nodes: []ast.Node{$1, ast.StarSplat{Value: $4}}}; vars.Line = $1.LineNumber(); $$ = vars }

| instance_variable COMMA REF
  { vars := ast.Array{Nodes: []ast.Node{$1, $3}}; vars.Line = $1.LineNumber(); $$ = vars }
| instance_variable COMMA instance_variable
  { vars := ast.Array{Nodes: []ast.Node{$1, $3}}; vars.Line = $1.LineNumber(); $$ = vars }
| instance_variable COMMA class_variable
  { vars := ast.Array{Nodes: []ast.Node{$1, $3}}; vars.Line = $1.LineNumber(); $$ = vars }
| instance_variable COMMA STAR REF
  { vars := ast.Array{Nodes: []ast.Node{$1, ast.StarSplat{Value: $4}}}; vars.Line = $1.LineNumber(); $$ = vars }

| class_variable COMMA REF
  { vars := ast.Array{Nodes: []ast.Node{$1, $3}}; vars.Line = $1.LineNumber(); $$ = vars }
| class_variable COMMA instance_variable
  { vars := ast.Array{Nodes: []ast.Node{$1, $3}}; vars.Line = $1.LineNumber(); $$ = vars }
| class_variable COMMA class_variable
  { vars := ast.Array{Nodes: []ast.Node{$1, $3}}; vars.Line = $1.LineNumber(); $$ = vars }
| class_variable COMMA STAR REF
  { vars := ast.Array{Nodes: []ast.Node{$1, ast.StarSplat{Value: $4}}}; vars.Line = $1.LineNumber(); $$ = vars }

| assignable_variables COMMA REF
  { vars := ast.Array{Nodes: append($$.(ast.Array).Nodes, $3)}; vars.Line = $1.LineNumber(); $$ = vars }
| assignable_variables COMMA instance_variable
  { vars := ast.Array{Nodes: append($$.(ast.Array).Nodes, $3)}; vars.Line = $1.LineNumber(); $$ = vars }
| assignable_variables COMMA class_variable
  { vars := ast.Array{Nodes: append($$.(ast.Array).Nodes, $3)}; vars.Line = $1.LineNumber(); $$ = vars }
| assignable_variables COMMA STAR REF
  { vars := ast.Array{Nodes: []ast.Node{$1, ast.StarSplat{Value: $4}}} ; vars.Line = $1.LineNumber(); $$ = vars };


negation : BANG expr
  { bang := ast.Negation{Target: $2}; bang.Line = $2.LineNumber(); $$ = bang };
complement : COMPLEMENT expr
  { comp := ast.Complement{Target: $2}; comp.Line = $2.LineNumber(); $$ = comp };
positive : UNARY_PLUS single_node
  { plus := ast.Positive{Target: $2}; plus.Line = $2.LineNumber(); $$ = plus };
negative : UNARY_MINUS single_node
  { minus := ast.Negative{Target: $2}; minus.Line = $2.LineNumber(); $$ = minus };

binary_addition : single_node BINARY_PLUS single_node
  {
    add := ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Line: $2.LineNumber(), Name: "+"},
      Args: []ast.Node{$3},
    }
    add.Line = $1.LineNumber()
    $$ = add
  };

binary_subtraction : single_node BINARY_MINUS expr
  {
    sub := ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Line: $2.LineNumber(), Name: "-"},
      Args: []ast.Node{$3},
    }
    sub.Line = $1.LineNumber()
    $$ = sub
  };

binary_multiplication : single_node STAR single_node
  {
    mult := ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Line: $2.LineNumber(), Name: "*"},
      Args: []ast.Node{$3},
    }
    mult.Line = $1.LineNumber()
    $$ = mult
  };

binary_division : single_node SLASH single_node
  {
    divis := ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Line: $2.LineNumber(), Name: "/"},
      Args: []ast.Node{$3},
    }
    divis.Line = $1.LineNumber()
    $$ = divis
  };

bitwise_and: single_node AMPERSAND single_node
  {
    and := ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Line: $2.LineNumber(), Name: "&"},
      Args: []ast.Node{$3},
    }
    and.Line = $1.LineNumber()
    $$ = and
  };

bitwise_or: single_node PIPE single_node
  {
    or := ast.CallExpression{
      Target: $1,
      Func: ast.BareReference{Line: $2.LineNumber(), Name: "|"},
      Args: []ast.Node{$3},
    }
    or.Line = $1.LineNumber()
    $$ = or
  };

array : LBRACKET optional_newlines nodes_with_commas_and_optional_newlines optional_newlines RBRACKET
  { $$ = ast.Array{Line: $1.LineNumber(), Nodes: $3} };

self : SELF { };
nil : NIL { };

nodes_with_commas_and_optional_newlines : /* empty */ { $$ = ast.Nodes{} }
| single_node
  { $$ = append($$, $1) }
| assignment
  { $$ = append($$, $1) }
| proc_arg
  { $$ = append($$, $1) }
| nodes_with_commas_and_optional_newlines COMMA optional_newlines single_node
  { $$ = append($$, $4) }
| nodes_with_commas_and_optional_newlines COMMA optional_newlines assignment
  { $$ = append($$, $4) }
| nodes_with_commas_and_optional_newlines COMMA optional_newlines proc_arg
  { $$ = append($$, $4) };

hash : LBRACE optional_newlines RBRACE
  { $$ = ast.Hash{Line: $1.LineNumber()} }
| LBRACE optional_newlines key_value_pairs optional_newlines RBRACE
  {
    pairs := []ast.HashKeyValuePair{}
    for _, node := range $3 {
      pairs = append(pairs, node.(ast.HashKeyValuePair))
    }
    $$ = ast.Hash{Line: $1.LineNumber(), Pairs: pairs}
  }
| LBRACE optional_newlines symbol_key_value_pairs optional_newlines RBRACE
  {
    pairs := []ast.HashKeyValuePair{}
    for _, node := range $3 {
      pairs = append(pairs, node.(ast.HashKeyValuePair))
    }
    $$ = ast.Hash{Line: $1.LineNumber(), Pairs: pairs}
  };

key_value_pairs : single_node OPERATOR expr
  {
    name := $2.(ast.BareReference).Name
    if name != "=>" {
      panic("FREAKOUT")
        }
    $$ = append($$, ast.HashKeyValuePair{Key: $1, Value: $3})
  }
| key_value_pairs COMMA optional_newlines single_node OPERATOR expr optional_comma
  {
    name := $5.(ast.BareReference).Name
    if name != "=>" {
      panic("FREAKOUT")
    }
    $$ = append($$, ast.HashKeyValuePair{Key: $4, Value: $6})
  };

symbol_key_value_pairs : REF COLON single_node
  {
    $$ = append($$, ast.HashKeyValuePair{
      Key: ast.Symbol{Line: $1.LineNumber(), Name: $1.(ast.BareReference).Name},
      Value: $3,
    })
  }
| symbol_key_value_pairs COMMA optional_newlines REF COLON single_node optional_newlines
  {
    $$ = append($$, ast.HashKeyValuePair{
      Key: ast.Symbol{Line: $4.LineNumber(), Name: $4.(ast.BareReference).Name},
      Value: $6,
    })
  }
| symbol_key_value_pairs COMMA optional_newlines REF COLON single_node COMMA optional_newlines
  {
    $$ = append($$, ast.HashKeyValuePair{
      Key: ast.Symbol{Line: $4.LineNumber(), Name: $4.(ast.BareReference).Name},
      Value: $6,
    })
  };

block : DO list END
  {
    block := ast.Block{Line: $1.LineNumber(), Body: $2}
    $$ = block
  }
| DO block_args list END
  {
    block := ast.Block{Line: $1.LineNumber(), Args: $2, Body: $3}
    $$ = block
  }
| LBRACE optional_newlines list optional_newlines RBRACE
  {
    block := ast.Block{Line: $1.LineNumber(), Body: $3}
    $$ = block
  }
| LBRACE optional_newlines block_args list RBRACE
  {
    block := ast.Block{Line: $1.LineNumber(), Args: $3, Body: $4}
    $$ = block
  }
| LBRACE optional_newlines call_expression optional_newlines RBRACE
  {
    block := ast.Block{Line: $1.LineNumber(), Body: []ast.Node{$3}}
    $$ = block
  }
| LBRACE optional_newlines call_expression list optional_newlines RBRACE
  {
      head := []ast.Node{$3}
      tail := $4
      body := append(head, tail...)
      block := ast.Block{Line: $1.LineNumber(), Body: body}
      $$ = block
  };

block_args : PIPE comma_delimited_refs PIPE
  { $$ = $2 };

comma_delimited_refs : /* empty */ { $$ = ast.Nodes{} }
| REF
  { $$ = append($$, $1); }
| comma_delimited_refs COMMA REF
  { $$ = append($$, $3); };

if_block : IF expr list END
  {
    cond := ast.IfBlock{
      Condition: $2,
      Body: $3,
    }
    cond.Line = $2.LineNumber()
    $$ = cond
  }
| IF expr list elsif_block END
  {
    cond := ast.IfBlock{
      Condition: $2,
      Body: $3,
      Else: $4,
    }
    cond.Line = $2.LineNumber()
    $$ = cond
  }
| expr IF expr
  {
    cond := ast.IfBlock{
      Condition: $3,
      Body: []ast.Node{$1},
    }
    cond.Line = $1.LineNumber()
    $$ = cond
  }
| call_expression IF expr
  {
    cond := ast.IfBlock{
      Condition: $3,
      Body: []ast.Node{$1},
    }
    cond.Line = $1.LineNumber()
    $$ = cond
  }
| single_node UNLESS expr
  {
    cond := ast.IfBlock{
      Condition: ast.Negation{Line: $1.LineNumber(), Target: $3},
      Body: []ast.Node{$1},
    }
    cond.Line = $1.LineNumber()
    $$ = cond
  }
| call_expression UNLESS expr
  {
    cond := ast.IfBlock{
      Condition: ast.Negation{Line: $1.LineNumber(), Target: $3},
      Body: ast.Nodes{$1},
    }
    cond.Line = $1.LineNumber()
    $$ = cond
  }
| assignment UNLESS expr
  {
    cond := ast.IfBlock{
      Condition: ast.Negation{Line: $1.LineNumber(), Target: $3},
      Body: ast.Nodes{$1},
    }
    cond.Line = $1.LineNumber()
    $$ = cond
  }
| UNLESS expr NEWLINE list END
  {
    cond := ast.IfBlock{
      Condition: ast.Negation{Line: $2.LineNumber(), Target: $2},
      Body: $4,
    }
    cond.Line = $2.LineNumber()
    $$ = cond
  }
| UNLESS expr NEWLINE list elsif_block END
  {
    cond := ast.IfBlock{
      Condition: ast.Negation{Line: $2.LineNumber(), Target: $2},
      Body: $4,
      Else: $5,
    }
    cond.Line = $2.LineNumber()
    $$ = cond
  }
| UNLESS expr SEMICOLON list END
  {
    cond := ast.IfBlock{
      Condition: ast.Negation{Line: $2.LineNumber(), Target: $2},
      Body: $4,
    }
    cond.Line = $2.LineNumber()
    $$ = cond
  }
| expr UNLESS expr
  {
    cond := ast.IfBlock{
      Condition: ast.Negation{Line: $1.LineNumber(), Target: $3},
      Body: []ast.Node{$1},
    }
    cond.Line = $1.LineNumber()
    $$ = cond
  };


elsif_block : elsif_block ELSIF expr list
  {
    ifblock := ast.IfBlock{
      Line: $3.LineNumber(),
      Condition: $3,
      Body: $4,
    }
    $$ = append($$, ifblock)
  }
| elsif_block ELSE list
  {
    $$ = append($$, ast.IfBlock{
      Line: $2.LineNumber(),
      Condition: ast.Boolean{Line: $2.LineNumber(), Value: true},
      Body: $3,
    })
      }
| ELSIF expr list
  {
    $$ = append($$, ast.IfBlock{
      Line: $2.LineNumber(),
      Condition: $2,
      Body: $3,
    })
  }
| ELSE list
  {
    $$ = append($$, ast.IfBlock{
      Line: $1.LineNumber(),
      Condition: ast.Boolean{Line: $1.LineNumber(), Value: true},
      Body: $2,
    })
      };

lines : /* empty */ { }
| lines expr { $$ = append($$, $2) }
| lines SEMICOLON { };

group : LPAREN lines RPAREN
  { group := ast.Group{Body: $2}; group.Line = $1.(ast.Nil).Line; $$ = group };

begin_block : BEGIN list optional_rescues END
  {
    begin := ast.Begin{
      Body: $2,
      Rescue: $3,
    }
    begin.Line = $1.LineNumber()
    $$ = begin
  }
| BEGIN list optional_rescues ELSE list END
  {
    begin := ast.Begin{
      Body: $2,
      Rescue: $3,
      Else: $5,
    }
    begin.Line = $1.LineNumber()
    $$ = begin
  };

rescue : RESCUE list
  { $$ = ast.Rescue{Line: $1.LineNumber(), Body: $2} }
| RESCUE comma_delimited_class_names list
  {
    classes := []ast.Class{}
    for _, class := range $2 {
      classes = append(classes, class.(ast.Class))
    }
    $$ = ast.Rescue{
      Line: $1.LineNumber(),
      Body: $3,
      Exception: ast.RescueException{
        Classes: classes,
      },
    }
  }
| RESCUE comma_delimited_class_names OPERATOR REF list
  {
    name := $3.(ast.BareReference).Name
    if name != "=>" {
      panic("FREAKOUT")
    }

    classes := []ast.Class{}
    for _, class := range $2 {
      classes = append(classes, class.(ast.Class))
    }

    $$ = ast.Rescue{
      Line: $1.LineNumber(),
      Body: $5,
      Exception: ast.RescueException{
        Var: $4.(ast.BareReference),
        Classes: classes,
      },
    }
  }
| RESCUE OPERATOR REF list
  {
    name := $2.(ast.BareReference).Name
    if name != "=>" {
      panic("FREAKOUT")
    }

    $$ = ast.Rescue{
      Line: $1.LineNumber(),
      Body: $4,
      Exception: ast.RescueException{
        Var: $3.(ast.BareReference),
      },
    }
  };

comma_delimited_class_names : class_name_with_modules
  { $$ = append($$, $1) }
| comma_delimited_class_names COMMA class_name_with_modules
  { $$ = append($$, $3) };

optional_rescues : /* empty */
  { $$ = []ast.Node{} }
| optional_rescues rescue
  { $$ = append($$, $2) };

rescues : rescue
  { $$ = append($$, $1) }
| rescues rescue
  { $$ = append($$, $2) };

yield_expression : YIELD comma_delimited_nodes
  {
    if len($2) == 1 {
      $$ = ast.Yield{Line: $1.LineNumber(), Value: $2[0]}
    } else {
      $$ = ast.Yield{Line: $1.LineNumber(), Value: $2}
    }
  }
| YIELD { $$ = ast.Yield{Line: $1.LineNumber(), } };

retry_expression : RETRY { $$ = ast.Retry{Line: $1.LineNumber(), } };

return_expression : RETURN comma_delimited_nodes
  {
    if len($2) == 1 {
      $$ = ast.Return{Line: $1.LineNumber(), Value: $2[0]}
    } else {
      $$ = ast.Return{Line: $1.LineNumber(), Value: $2}
    }
  }
| RETURN
  { $$ = ast.Return{Line: $1.LineNumber(), } };


next_expression : NEXT
  { $$ = ast.Next{} }
| NEXT IF expr
  { $$ = ast.IfBlock{Line: $3.LineNumber(), Condition: $3, Body: []ast.Node{ast.Next{}}} }
| NEXT UNLESS expr
  { $$ = ast.IfBlock{Line: $3.LineNumber(), Condition: ast.Negation{Line: $3.LineNumber(), Target: $3}, Body: []ast.Node{ast.Next{}}} };


break_expression: BREAK
  { $$ = ast.Break{} }
| BREAK IF expr
  { $$ = ast.IfBlock{Line: $3.LineNumber(), Condition: $3, Body: []ast.Node{ast.Break{}}} }
| BREAK UNLESS expr
  { $$ = ast.IfBlock{Line: $3.LineNumber(), Condition: ast.Negation{Line: $3.LineNumber(), Target: $3}, Body: []ast.Node{ast.Break{}}} };


ternary : single_node QUESTIONMARK single_node COLON single_node
  {
    ternary := ast.Ternary{
      Condition: $1,
      True: $3,
      False: $5,
    }
    ternary.Line = $1.LineNumber()
    $$ = ternary
  }

while_loop : WHILE expr NEWLINE loop_expressions END
  {
    loop := ast.Loop{Condition: $2, Body: $4}
    loop.Line = $2.LineNumber()
    $$ = loop
  }
| UNTIL expr NEWLINE loop_expressions END
  {
    condition := ast.Negation{Line: $2.LineNumber(), Target:$2}
    loop := ast.Loop{Condition: condition, Body: $4}
    loop.Line = $2.LineNumber()
    $$ = loop
  }
| expr WHILE expr
  {
    loop := ast.Loop{Condition: $3, Body: []ast.Node{$1}}
    loop.Line = $3.LineNumber()
    $$ = loop
  };

loop_expressions : /* empty */
  { $$ = ast.Nodes{} }
| loop_expressions NEWLINE
  {  }
| loop_expressions SEMICOLON
  {  }
| loop_expressions expr
  {  $$ = append($$, $2) }
| loop_expressions loop_if_block
  {  $$ = append($$, $2) };

loop_if_block : IF expr NEWLINE loop_expressions END
  {
    $$ = ast.IfBlock{
      Line: $2.LineNumber(),
      Condition: $2,
      Body: $4,
    }
  }
| IF expr NEWLINE loop_expressions loop_elsif_block END
  {
    $$ = ast.IfBlock{
      Line: $2.LineNumber(),
      Condition: $2,
      Body: $4,
      Else: $5,
    }
  }
| UNLESS expr NEWLINE loop_expressions END
  {
    $$ = ast.IfBlock{
      Line: $2.LineNumber(),
      Condition: ast.Negation{Line: $2.LineNumber(), Target: $2},
      Body: $4,
    }
  }
| UNLESS expr NEWLINE loop_expressions loop_elsif_block END
  {
    $$ = ast.IfBlock{
      Line: $2.LineNumber(),
      Condition: ast.Negation{Line: $2.LineNumber(), Target: $2},
      Body: $4,
      Else: $5,
    }
  }
| UNLESS expr SEMICOLON loop_expressions END
  {
    $$ = ast.IfBlock{
      Line: $2.LineNumber(),
      Condition: ast.Negation{Line: $2.LineNumber(), Target: $2},
      Body: $4,
    }
  };

loop_elsif_block : loop_elsif_block ELSIF expr loop_expressions
  {
    $$ = append($$, ast.IfBlock{
      Line: $3.LineNumber(),
      Condition: $3,
      Body: $4,
    })
  }
| loop_elsif_block ELSE loop_expressions
  {
    $$ = append($$, ast.IfBlock{
      Line: $1.LineNumber(),
      Condition: ast.Boolean{Line: $1.LineNumber(), Value: true},
      Body: $3,
    })
      }
| ELSIF expr loop_expressions
  {
    $$ = append($$, ast.IfBlock{
      Line: $2.LineNumber(),
      Condition: $2,
      Body: $3,
    })
  }
| ELSE loop_expressions
  {
    $$ = append($$, ast.IfBlock{
      Line: $1.LineNumber(),
      Condition: ast.Boolean{Line: $1.LineNumber(), Value: true},
      Body: $2,
    })
   };

logical_and : single_node AND optional_newlines single_node
  { $$ = ast.WeakLogicalAnd{Line: $1.LineNumber(), LHS: $1, RHS: $4} };

logical_or : single_node OR optional_newlines single_node
  { $$ = ast.WeakLogicalOr{Line: $1.LineNumber(), LHS: $1, RHS: $4} };

lambda : LAMBDA block
  {
    lambda := ast.Lambda{Body: $2}
    lambda.Line = $2.LineNumber()
    $$ = lambda
  };

switch_statement : CASE single_node optional_newlines switch_cases END
  {
    switchstmt := ast.SwitchStatement{Condition: $2, Cases: $4}
    switchstmt.Line = $1.LineNumber()
    $$ = switchstmt
  }
| CASE single_node optional_newlines switch_cases ELSE list END
  {
    switchstmt := ast.SwitchStatement{Condition: $2, Cases: $4, Else: $6}
    switchstmt.Line = $1.LineNumber()
    $$ = switchstmt
  }
| CASE optional_newlines switch_cases END
  {
    switchstmt := ast.SwitchStatement{Cases: $3}
    switchstmt.Line = $1.LineNumber()
    $$ = switchstmt
  }
| CASE optional_newlines switch_cases ELSE list END
  {
    switchstmt := ast.SwitchStatement{Cases: $3, Else: $5}
    switchstmt.Line = $1.LineNumber()
    $$ = switchstmt
  };

switch_cases : WHEN comma_delimited_nodes list optional_newlines
  { $$ = append($$, ast.SwitchCase{Conditions: $2, Body: $3}) }
| switch_cases WHEN comma_delimited_nodes list optional_newlines
  { $$ = append($$, ast.SwitchCase{Conditions: $3, Body: $4}) };

range : single_node RANGE single_node
  {
    rng := ast.Range{Start: $1, End: $3}
    rng.Line = $1.LineNumber();
    $$ = rng
  };

alias : ALIAS SYMBOL SYMBOL
  {
    alias := ast.Alias{To: $2.(ast.Symbol), From: $3.(ast.Symbol)}
    alias.Line = $1.LineNumber()
    $$ = alias
  };

%%
