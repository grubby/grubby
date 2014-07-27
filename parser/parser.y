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
%token <genericValue> LPAREN
%token <genericValue> RPAREN
%token <genericValue> COMMA

// keywords
%token <genericValue> DEF
%token <genericValue> END
%token <genericValue> CLASS
%token <genericValue> MODULE

// operators
%token <genericValue> LESSTHAN
%token <genericValue> GREATERTHAN
%token <genericValue> EQUALTO

// misc
%token <genericValue> COLON
%token <genericValue> DOT

/*
  eg: if you want to be able to assign to something in the RubySymType
      struct, or if you want a terminating node below, you will want to
      declare a type (or possibly just a token)
*/

// single nodes
%type <genericValue> expr
%type <genericValue> symbol
%type <genericValue> callexpr
%type <genericValue> statement
%type <genericValue> func_declaration
%type <genericValue> class_declaration
%type <genericValue> class_name_with_modules

// slice nodes
%type <genericSlice> list
%type <genericSlice> callargs
%type <genericSlice> capture_list
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

statement : callexpr
  { $$ = $1 }
| expr
  { $$ = $1 }
| '\n'
  { $$ = nil; /* ignores new lines around statements */ }
| ' '
  { $$ = nil; /* ignores whitespace around statements */ };

callexpr : REF callargs
  {
    $$ = ast.CallExpression{
      Func: $1,
      Args: $2,
    }
  }
| REF " " callargs
  {
    $$ = ast.CallExpression{
      Func: $1,
      Args: $3,
    }
  };

expr : NODE | REF | func_declaration | class_declaration | symbol;

callargs : /* empty */ { $$ = ast.Nodes{} }
| NODE
  { $$ = append($$, $1) }
| REF
  { $$ = append($$, $1) }
| LPAREN nodes_with_commas RPAREN
  {
		$$ = $2
	};

nodes_with_commas : /* empty */ { $$ = ast.Nodes{} }
| REF
  { $$ = append($$, $1); }
| NODE
  { $$ = append($$, $1); }
| nodes_with_commas COMMA " " NODE
  { $$ = append($$, $4); };
| nodes_with_commas COMMA " " REF
  { $$ = append($$, $4); };

func_declaration : DEF " " REF callargs "\n" list END
  {
		$$ = ast.FuncDecl{
			Name: $3.(ast.BareReference),
      Args: $4,
			Body: $6,
    }
  };

class_declaration: CLASS " " class_name_with_modules "\n" list END
  {
    $$ = ast.ClassDecl{
       Name: $3.(ast.Class).Name,
       Body: $5,
    }
  }
| CLASS " " class_name_with_modules " " LESSTHAN " " class_name_with_modules "\n" list END
  {
    $$ = ast.ClassDecl{
       Name: $3.(ast.Class).Name,
       SuperClass: $7.(ast.Class),
       Namespace: $3.(ast.Class).Namespace,
       Body: $9,
    }
  }

class_name_with_modules: REF
  {
    $$ = ast.Class{
      Name: $1.(ast.BareReference).Name,
    }
  }
| namespaced_modules COLON COLON REF
  {
    $$ = ast.Class{
       Name: $4.(ast.BareReference).Name,
       Namespace: strings.Join($1, "::"),
    }
  };

namespaced_modules : REF
  {
    $$ = append($$, $1.(ast.BareReference).Name)
  }
|  namespaced_modules COLON COLON REF
  {
    $$ = append($$, $4.(ast.BareReference).Name)
  };

%%
