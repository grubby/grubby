//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:3
import (
	"github.com/grubby/grubby/ast"
	"strings"
)

var Statements []ast.Node

//line parser.y:16
type RubySymType struct {
	yys          int
	genericValue ast.Node
	genericSlice ast.Nodes
	stringSlice  []string
}

const NODE = 57346
const REF = 57347
const CAPITAL_REF = 57348
const LPAREN = 57349
const RPAREN = 57350
const COMMA = 57351
const DEF = 57352
const END = 57353
const CLASS = 57354
const MODULE = 57355
const LESSTHAN = 57356
const GREATERTHAN = 57357
const EQUALTO = 57358
const COLON = 57359
const DOT = 57360

var RubyToknames = []string{
	"NODE",
	"REF",
	"CAPITAL_REF",
	"LPAREN",
	"RPAREN",
	"COMMA",
	"DEF",
	"END",
	"CLASS",
	"MODULE",
	"LESSTHAN",
	"GREATERTHAN",
	"EQUALTO",
	"COLON",
	"DOT",
}
var RubyStatenames = []string{}

const RubyEofCode = 1
const RubyErrCode = 2
const RubyMaxDepth = 200

//line parser.y:188

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 30,
	17, 32,
	-2, 30,
	-1, 49,
	17, 33,
	-2, 31,
}

const RubyNprod = 34
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 81

var RubyAct = []int{

	40, 29, 16, 8, 7, 9, 35, 36, 48, 13,
	54, 14, 38, 22, 18, 19, 15, 20, 5, 6,
	24, 8, 7, 9, 21, 52, 39, 13, 50, 14,
	17, 34, 42, 37, 15, 41, 5, 6, 32, 33,
	45, 8, 7, 9, 30, 49, 28, 13, 47, 14,
	51, 43, 44, 53, 15, 23, 5, 6, 8, 7,
	9, 31, 18, 19, 13, 20, 14, 27, 26, 25,
	46, 15, 2, 5, 6, 1, 11, 10, 3, 12,
	4,
}
var RubyPact = []int{

	-1000, 54, -1000, -1000, -1000, -1000, -1000, 10, -1000, -1000,
	-1000, -1000, -1000, 4, -7, 50, -1000, 58, -1000, -1000,
	63, 41, 38, -1000, -1000, 30, -1000, -1000, 58, -13,
	-1000, 16, -1000, -8, 7, -1000, 21, 15, 47, -1000,
	37, -12, 39, -1000, -1000, 17, -1000, -1000, 38, -1000,
	-1000, 6, -1000, -1, -1000,
}
var RubyPgo = []int{

	0, 80, 79, 78, 70, 77, 76, 1, 0, 2,
	75, 69, 61,
}
var RubyR1 = []int{

	0, 10, 10, 8, 8, 2, 4, 4, 4, 4,
	3, 3, 1, 1, 1, 1, 1, 1, 9, 9,
	9, 9, 11, 11, 11, 11, 11, 5, 6, 6,
	7, 7, 12, 12,
}
var RubyR2 = []int{

	0, 0, 2, 0, 2, 2, 1, 1, 1, 1,
	2, 3, 1, 1, 1, 1, 1, 1, 0, 1,
	1, 3, 0, 1, 1, 4, 4, 7, 6, 10,
	1, 4, 1, 4,
}
var RubyChk = []int{

	-1000, -10, -4, -3, -1, 19, 20, 5, 4, 6,
	-5, -6, -2, 10, 12, 17, -9, 20, 4, 5,
	7, 20, 20, 5, -9, -11, 5, 4, 5, -7,
	6, -12, 8, 9, -9, 19, 20, 17, 20, 19,
	-8, 14, 17, 4, 5, -8, -4, 11, 20, 6,
	11, -7, 19, -8, 11,
}
var RubyDef = []int{

	1, -2, 2, 6, 7, 8, 9, 13, 12, 14,
	15, 16, 17, 0, 0, 0, 10, 18, 19, 20,
	22, 0, 0, 5, 11, 0, 23, 24, 18, 0,
	-2, 0, 21, 0, 0, 3, 0, 0, 0, 3,
	0, 0, 0, 25, 26, 0, 4, 28, 0, -2,
	27, 0, 3, 0, 29,
}
var RubyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	19, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 20,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18,
}
var RubyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var RubyDebug = 0

type RubyLexer interface {
	Lex(lval *RubySymType) int
	Error(s string)
}

const RubyFlag = -1000

func RubyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(RubyToknames) {
		if RubyToknames[c-4] != "" {
			return RubyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func RubyStatname(s int) string {
	if s >= 0 && s < len(RubyStatenames) {
		if RubyStatenames[s] != "" {
			return RubyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func Rubylex1(lex RubyLexer, lval *RubySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = RubyTok1[0]
		goto out
	}
	if char < len(RubyTok1) {
		c = RubyTok1[char]
		goto out
	}
	if char >= RubyPrivate {
		if char < RubyPrivate+len(RubyTok2) {
			c = RubyTok2[char-RubyPrivate]
			goto out
		}
	}
	for i := 0; i < len(RubyTok3); i += 2 {
		c = RubyTok3[i+0]
		if c == char {
			c = RubyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = RubyTok2[1] /* unknown char */
	}
	if RubyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", RubyTokname(c), uint(char))
	}
	return c
}

func RubyParse(Rubylex RubyLexer) int {
	var Rubyn int
	var Rubylval RubySymType
	var RubyVAL RubySymType
	RubyS := make([]RubySymType, RubyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Rubystate := 0
	Rubychar := -1
	Rubyp := -1
	goto Rubystack

ret0:
	return 0

ret1:
	return 1

Rubystack:
	/* put a state and value onto the stack */
	if RubyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", RubyTokname(Rubychar), RubyStatname(Rubystate))
	}

	Rubyp++
	if Rubyp >= len(RubyS) {
		nyys := make([]RubySymType, len(RubyS)*2)
		copy(nyys, RubyS)
		RubyS = nyys
	}
	RubyS[Rubyp] = RubyVAL
	RubyS[Rubyp].yys = Rubystate

Rubynewstate:
	Rubyn = RubyPact[Rubystate]
	if Rubyn <= RubyFlag {
		goto Rubydefault /* simple state */
	}
	if Rubychar < 0 {
		Rubychar = Rubylex1(Rubylex, &Rubylval)
	}
	Rubyn += Rubychar
	if Rubyn < 0 || Rubyn >= RubyLast {
		goto Rubydefault
	}
	Rubyn = RubyAct[Rubyn]
	if RubyChk[Rubyn] == Rubychar { /* valid shift */
		Rubychar = -1
		RubyVAL = Rubylval
		Rubystate = Rubyn
		if Errflag > 0 {
			Errflag--
		}
		goto Rubystack
	}

Rubydefault:
	/* default state action */
	Rubyn = RubyDef[Rubystate]
	if Rubyn == -2 {
		if Rubychar < 0 {
			Rubychar = Rubylex1(Rubylex, &Rubylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if RubyExca[xi+0] == -1 && RubyExca[xi+1] == Rubystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Rubyn = RubyExca[xi+0]
			if Rubyn < 0 || Rubyn == Rubychar {
				break
			}
		}
		Rubyn = RubyExca[xi+1]
		if Rubyn < 0 {
			goto ret0
		}
	}
	if Rubyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Rubylex.Error("syntax error")
			Nerrs++
			if RubyDebug >= 1 {
				__yyfmt__.Printf("%s", RubyStatname(Rubystate))
				__yyfmt__.Printf(" saw %s\n", RubyTokname(Rubychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Rubyp >= 0 {
				Rubyn = RubyPact[RubyS[Rubyp].yys] + RubyErrCode
				if Rubyn >= 0 && Rubyn < RubyLast {
					Rubystate = RubyAct[Rubyn] /* simulate a shift of "error" */
					if RubyChk[Rubystate] == RubyErrCode {
						goto Rubystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if RubyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", RubyS[Rubyp].yys)
				}
				Rubyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if RubyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", RubyTokname(Rubychar))
			}
			if Rubychar == RubyEofCode {
				goto ret1
			}
			Rubychar = -1
			goto Rubynewstate /* try again in the same state */
		}
	}

	/* reduction by production Rubyn */
	if RubyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Rubyn, RubyStatname(Rubystate))
	}

	Rubynt := Rubyn
	Rubypt := Rubyp
	_ = Rubypt // guard against "declared and not used"

	Rubyp -= RubyR2[Rubyn]
	RubyVAL = RubyS[Rubyp+1]

	/* consult goto table to find next state */
	Rubyn = RubyR1[Rubyn]
	Rubyg := RubyPgo[Rubyn]
	Rubyj := Rubyg + RubyS[Rubyp].yys + 1

	if Rubyj >= RubyLast {
		Rubystate = RubyAct[Rubyg]
	} else {
		Rubystate = RubyAct[Rubyj]
		if RubyChk[Rubystate] != -Rubyn {
			Rubystate = RubyAct[Rubyg]
		}
	}
	// dummy call; replaced with literal code
	switch Rubynt {

	case 1:
		//line parser.y:71
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:73
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 3:
		//line parser.y:80
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 4:
		//line parser.y:82
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
			}
		}
	case 5:
		//line parser.y:89
		{
			RubyVAL.genericValue = ast.Symbol{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 6:
		//line parser.y:94
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 7:
		//line parser.y:96
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 8:
		//line parser.y:98
		{
			RubyVAL.genericValue = nil /* ignores new lines around statements */
		}
	case 9:
		//line parser.y:100
		{
			RubyVAL.genericValue = nil /* ignores whitespace around statements */
		}
	case 10:
		//line parser.y:103
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue,
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 11:
		//line parser.y:110
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue,
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 12:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 13:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 14:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 15:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 16:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 17:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 18:
		//line parser.y:119
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 19:
		//line parser.y:121
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 20:
		//line parser.y:123
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 21:
		//line parser.y:125
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 22:
		//line parser.y:129
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 23:
		//line parser.y:131
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 24:
		//line parser.y:133
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 25:
		//line parser.y:135
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 26:
		//line parser.y:137
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 27:
		//line parser.y:140
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-3].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 28:
		//line parser.y:149
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 29:
		//line parser.y:156
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-7].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-7].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 30:
		//line parser.y:166
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 31:
		//line parser.y:172
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 32:
		//line parser.y:180
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 33:
		//line parser.y:184
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	}
	goto Rubystack /* stack new state and value */
}
