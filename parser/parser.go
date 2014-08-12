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
const TRUE = 57356
const FALSE = 57357
const LESSTHAN = 57358
const GREATERTHAN = 57359
const EQUALTO = 57360
const BANG = 57361
const COMPLEMENT = 57362
const POSITIVE = 57363
const NEGATIVE = 57364
const STAR = 57365
const WHITESPACE = 57366
const NEWLINE = 57367
const SEMICOLON = 57368
const COLON = 57369
const DOT = 57370
const LBRACKET = 57371
const RBRACKET = 57372
const LBRACE = 57373
const RBRACE = 57374
const DOLLARSIGN = 57375
const ATSIGN = 57376
const EOF = 57377

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
	"TRUE",
	"FALSE",
	"LESSTHAN",
	"GREATERTHAN",
	"EQUALTO",
	"BANG",
	"COMPLEMENT",
	"POSITIVE",
	"NEGATIVE",
	"STAR",
	"WHITESPACE",
	"NEWLINE",
	"SEMICOLON",
	"COLON",
	"DOT",
	"LBRACKET",
	"RBRACKET",
	"LBRACE",
	"RBRACE",
	"DOLLARSIGN",
	"ATSIGN",
	"EOF",
}
var RubyStatenames = []string{}

const RubyEofCode = 1
const RubyErrCode = 2
const RubyMaxDepth = 200

//line parser.y:351

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	23, 6,
	-2, 2,
	-1, 20,
	18, 6,
	-2, 27,
	-1, 68,
	9, 6,
	-2, 38,
	-1, 73,
	21, 6,
	22, 6,
	23, 6,
	-2, 54,
	-1, 74,
	21, 6,
	22, 6,
	23, 6,
	-2, 55,
	-1, 75,
	21, 6,
	22, 6,
	23, 6,
	-2, 56,
	-1, 76,
	21, 6,
	22, 6,
	23, 6,
	-2, 57,
	-1, 90,
	9, 39,
	-2, 34,
	-1, 93,
	27, 50,
	-2, 6,
	-1, 99,
	21, 6,
	22, 6,
	23, 6,
	-2, 58,
	-1, 100,
	21, 6,
	22, 6,
	23, 6,
	-2, 59,
	-1, 101,
	21, 6,
	22, 6,
	23, 6,
	-2, 60,
	-1, 102,
	21, 6,
	22, 6,
	23, 6,
	-2, 53,
	-1, 103,
	21, 6,
	22, 6,
	23, 6,
	-2, 52,
	-1, 120,
	23, 6,
	-2, 5,
	-1, 134,
	27, 51,
	-2, 6,
	-1, 149,
	21, 6,
	22, 6,
	23, 6,
	-2, 70,
	-1, 158,
	21, 6,
	22, 6,
	23, 6,
	-2, 67,
	-1, 164,
	25, 71,
	32, 71,
	-2, 6,
	-1, 167,
	25, 68,
	32, 68,
	-2, 6,
}

const RubyNprod = 81
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 355

var RubyAct = []int{

	40, 107, 120, 68, 2, 80, 92, 70, 63, 132,
	122, 95, 53, 54, 153, 151, 41, 105, 160, 41,
	123, 42, 43, 44, 45, 46, 91, 39, 47, 48,
	49, 50, 51, 52, 41, 62, 110, 128, 112, 138,
	161, 55, 61, 58, 59, 60, 71, 71, 41, 97,
	73, 74, 75, 76, 72, 77, 163, 168, 166, 83,
	84, 85, 148, 86, 136, 87, 129, 105, 88, 89,
	90, 139, 41, 41, 58, 59, 60, 105, 96, 81,
	82, 98, 58, 59, 60, 134, 99, 100, 101, 102,
	103, 93, 104, 71, 109, 65, 66, 111, 67, 106,
	108, 115, 56, 57, 157, 117, 118, 119, 69, 64,
	65, 66, 94, 67, 125, 126, 127, 43, 130, 131,
	79, 78, 114, 113, 133, 1, 17, 16, 15, 140,
	14, 135, 137, 13, 141, 143, 144, 12, 146, 11,
	7, 34, 145, 149, 147, 6, 152, 5, 154, 155,
	35, 8, 20, 4, 36, 156, 18, 159, 158, 10,
	19, 162, 9, 0, 165, 164, 0, 0, 167, 3,
	21, 33, 0, 0, 0, 22, 150, 23, 24, 25,
	26, 0, 0, 0, 27, 28, 29, 30, 0, 0,
	0, 0, 0, 0, 31, 0, 32, 0, 38, 37,
	3, 21, 33, 0, 0, 0, 22, 142, 23, 24,
	25, 26, 0, 0, 0, 27, 28, 29, 30, 0,
	0, 0, 0, 0, 0, 31, 0, 32, 0, 38,
	37, 3, 21, 33, 0, 0, 0, 22, 124, 23,
	24, 25, 26, 0, 0, 0, 27, 28, 29, 30,
	0, 0, 0, 0, 0, 0, 31, 0, 32, 0,
	38, 37, 3, 21, 33, 0, 0, 0, 22, 121,
	23, 24, 25, 26, 0, 0, 0, 27, 28, 29,
	30, 0, 0, 0, 0, 0, 0, 31, 0, 32,
	0, 38, 37, 3, 21, 33, 0, 0, 0, 22,
	0, 23, 24, 25, 26, 0, 0, 0, 27, 28,
	29, 30, 0, 0, 0, 0, 0, 0, 31, 0,
	32, 0, 38, 37, 3, 116, 33, 0, 0, 0,
	22, 0, 23, 24, 25, 26, 0, 0, 0, 27,
	28, 29, 30, 0, 0, 0, 0, 0, 0, 31,
	0, 32, 0, 38, 37,
}
var RubyPact = []int{

	-1000, 289, -8, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-5, -5, -5, -5, -5, -1000, -1000, -5, -5, -5,
	-5, -5, -5, -1000, -1000, -1000, -1000, 7, 97, -1000,
	61, -5, 17, 91, 103, -5, -5, 289, 289, 289,
	289, 116, -1000, -1000, -1000, 74, -1000, -1000, -5, -5,
	-5, -1000, -5, -1000, -5, -1000, -1000, -5, -5, -5,
	-10, 85, -25, -5, -5, -5, -5, -5, -1000, -1000,
	24, -1000, -1000, 289, 289, 289, 289, 289, 116, 68,
	106, -1000, -5, -5, 9, -1000, 8, -1000, 320, -5,
	-5, -5, -5, -5, -5, -5, -5, 258, -26, -1000,
	-7, 227, -1000, -5, -5, -5, 10, 58, 114, -27,
	-5, -1000, -1000, 79, -1000, 55, 30, 53, -5, -1000,
	-1000, -1000, -1000, 196, -5, 24, -1000, 24, -1000, 45,
	289, 165, -1000, -1000, -17, 24, -18, 24, -5, -5,
	-1000, -1000, 289, -1000, 99, 289, -5, -9, -5, 22,
	-5, 39, 289, -5, 49, 289, -1000, 48, -1000,
}
var RubyPgo = []int{

	0, 163, 2, 162, 160, 159, 156, 154, 153, 152,
	151, 150, 147, 145, 141, 140, 7, 139, 137, 133,
	130, 128, 127, 126, 1, 8, 125, 123, 122, 3,
	112, 0, 5,
}
var RubyR1 = []int{

	0, 26, 26, 26, 24, 24, 31, 31, 32, 32,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 9, 9,
	9, 9, 9, 8, 25, 25, 25, 25, 25, 29,
	29, 29, 29, 29, 12, 13, 13, 15, 16, 16,
	30, 30, 10, 10, 17, 18, 19, 20, 21, 22,
	23, 3, 5, 6, 4, 4, 27, 27, 27, 27,
	28, 28, 28, 1, 1, 7, 7, 14, 14, 11,
	11,
}
var RubyR2 = []int{

	0, 0, 2, 3, 0, 2, 0, 2, 0, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 0, 1, 1, 5, 1, 0,
	1, 1, 5, 5, 9, 6, 8, 6, 3, 6,
	1, 4, 5, 5, 3, 3, 3, 3, 5, 5,
	5, 1, 1, 5, 9, 9, 0, 6, 11, 12,
	4, 9, 10, 0, 1, 2, 2, 2, 2, 3,
	3,
}
var RubyChk = []int{

	-1000, -26, -2, 4, -8, -12, -13, -15, -10, -3,
	-5, -17, -18, -19, -20, -21, -22, -23, -6, -4,
	-9, 5, 10, 12, 13, 14, 15, 19, 20, 21,
	22, 29, 31, 6, -14, -11, -7, 34, 33, 35,
	-31, 24, -31, -31, -31, -31, -31, -31, -31, -31,
	-31, -31, -31, 5, 6, 34, 5, 6, 21, 22,
	23, -31, 18, -25, 18, 4, 5, 7, -29, 5,
	-16, -31, -16, -2, -2, -2, -2, -29, 5, 4,
	-32, 5, 6, -31, -31, -31, -31, -31, -31, -31,
	-31, 36, 16, 6, -30, 36, -31, 25, -31, -2,
	-2, -2, -2, -2, -29, 9, -25, -24, -16, -31,
	27, -24, 30, -27, -28, -2, 5, -31, -31, -31,
	-2, 11, 36, 27, 11, -31, -31, -31, 27, 8,
	4, 5, 36, -24, 6, -32, 9, -32, 9, 18,
	-31, -24, 11, -31, -31, -32, -31, -32, 17, -2,
	11, 32, -31, 32, -31, -31, -2, 5, -2, -31,
	27, 18, -31, 17, -2, -31, 9, -2, 9,
}
var RubyDef = []int{

	1, -2, -2, 10, 11, 12, 13, 14, 15, 16,
	17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
	-2, 6, 6, 6, 6, 61, 62, 6, 6, 6,
	6, 6, 6, 29, 30, 31, 32, 0, 0, 3,
	0, 6, 0, 34, 0, 6, 6, 0, 0, 0,
	0, 39, 8, 77, 78, 0, 75, 76, 6, 6,
	6, 7, 6, 33, 6, 35, 36, 6, -2, 6,
	0, 0, 0, -2, -2, -2, -2, 6, 40, 41,
	6, 79, 80, 0, 0, 0, 0, 0, 39, 0,
	-2, 4, 6, -2, 0, 4, 0, 9, 66, -2,
	-2, -2, -2, -2, 6, 6, 6, 0, 0, 48,
	0, 0, 63, 6, 6, 6, 6, 0, 0, 0,
	-2, 45, 4, 0, 47, 8, 8, 0, 6, 37,
	42, 43, 4, 0, -2, 6, 8, 6, 8, 0,
	0, 0, 46, 49, 0, 6, 0, 6, 6, -2,
	44, 64, 0, 65, 0, 0, 6, 0, -2, 0,
	6, 0, 0, 6, -2, 0, 72, -2, 69,
}
var RubyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	36,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35,
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
		//line parser.y:113
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:115
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 3:
		//line parser.y:121
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 4:
		//line parser.y:128
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 5:
		//line parser.y:130
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
			}
		}
	case 10:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 11:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 19:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 20:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 21:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 22:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 23:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 24:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 25:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 26:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 27:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 28:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 29:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 30:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 31:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 32:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 33:
		//line parser.y:144
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 34:
		//line parser.y:151
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 35:
		//line parser.y:153
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 36:
		//line parser.y:155
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 37:
		//line parser.y:157
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 38:
		//line parser.y:161
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 39:
		//line parser.y:165
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 40:
		//line parser.y:167
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 41:
		//line parser.y:169
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 42:
		//line parser.y:171
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 43:
		//line parser.y:173
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 44:
		//line parser.y:178
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 45:
		//line parser.y:187
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 46:
		//line parser.y:194
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 47:
		//line parser.y:204
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 48:
		//line parser.y:213
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 49:
		//line parser.y:219
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 50:
		//line parser.y:227
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 51:
		//line parser.y:231
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 52:
		//line parser.y:236
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 53:
		//line parser.y:243
		{ // FIXME: this is a hack. If this rule is not present, "foo = 5" will not parse
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 54:
		//line parser.y:250
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 55:
		//line parser.y:251
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 56:
		//line parser.y:252
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 57:
		//line parser.y:253
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 58:
		//line parser.y:256
		{
			RubyVAL.genericValue = ast.Addition{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 59:
		//line parser.y:264
		{
			RubyVAL.genericValue = ast.Subtraction{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 60:
		//line parser.y:272
		{
			RubyVAL.genericValue = ast.Multiplication{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 61:
		//line parser.y:279
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 62:
		//line parser.y:280
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 63:
		//line parser.y:282
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 64:
		//line parser.y:286
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 65:
		//line parser.y:294
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 66:
		//line parser.y:302
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 67:
		//line parser.y:304
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 68:
		//line parser.y:306
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 69:
		//line parser.y:308
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 70:
		//line parser.y:311
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 71:
		//line parser.y:318
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 72:
		//line parser.y:325
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 73:
		//line parser.y:333
		{
			RubyVAL.genericValue = nil
		}
	case 74:
		//line parser.y:334
		{
			RubyVAL.genericValue = nil
		}
	case 75:
		//line parser.y:337
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 76:
		//line parser.y:339
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 77:
		//line parser.y:342
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 78:
		//line parser.y:344
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 79:
		//line parser.y:347
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 80:
		//line parser.y:349
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	}
	goto Rubystack /* stack new state and value */
}
