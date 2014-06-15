%{

package main

import (
  "bufio"
  "fmt"
  "os"
  "unicode"
)

var regs = make([]int, 26)
var base int

%}

%union{
  val int
}

// any token that returns a value needs a type
// these will become field names in the above struct
%type <val> expr number

%token <val> DIGIT LETTER

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

%%   /* start of legit parser? */

type RubyLex struct {
  s string
  pos int
}

func (l *RubyLex) Lex(lval *RubySymType) int {
  var c rune = ' '
  for c == ' ' {
    if l.pos == len(l.s) {
      return 0
    }

    c = rune(l.s[l.pos])
    l.pos++
  }

  if unicode.IsDigit(c) {
    fmt.Printf("storing an int value in a lex val %#v\n", lval)
    lval.val = int(c - '0')
    return DIGIT
  }

  return int(c)
}

func (l *RubyLex) Error(s string) {
  panic(fmt.Sprintf("syntax error: %s\n", s))
}

func main() {
  stdin := bufio.NewReader(os.NewFile(0, "stdin"))

  var line string
  var err error
  for {
    fmt.Printf("give me your best rubby: ")

    if line, err = stdin.ReadString('\n'); err == nil{
      ret := RubyParse(&RubyLex{s: line})
      fmt.Printf("got a %T :: %#v\n", ret, ret)
    } else {
      break
    }
  }
}
