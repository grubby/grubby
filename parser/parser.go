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

//line parser.y:365

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 4,
	23, 8,
	-2, 4,
	-1, 6,
	18, 8,
	-2, 13,
	-1, 23,
	4, 8,
	5, 8,
	7, 8,
	9, 8,
	18, 8,
	-2, 30,
	-1, 76,
	21, 8,
	22, 8,
	23, 8,
	-2, 59,
	-1, 77,
	21, 8,
	22, 8,
	23, 8,
	-2, 60,
	-1, 78,
	21, 8,
	22, 8,
	23, 8,
	-2, 61,
	-1, 79,
	21, 8,
	22, 8,
	23, 8,
	-2, 62,
	-1, 96,
	27, 55,
	-2, 8,
	-1, 102,
	21, 8,
	22, 8,
	23, 8,
	-2, 63,
	-1, 103,
	21, 8,
	22, 8,
	23, 8,
	-2, 64,
	-1, 104,
	21, 8,
	22, 8,
	23, 8,
	-2, 65,
	-1, 105,
	21, 8,
	22, 8,
	23, 8,
	-2, 58,
	-1, 106,
	21, 8,
	22, 8,
	23, 8,
	-2, 57,
	-1, 120,
	21, 30,
	22, 30,
	23, 30,
	-2, 8,
	-1, 124,
	23, 8,
	-2, 7,
	-1, 139,
	27, 56,
	-2, 8,
	-1, 156,
	21, 8,
	22, 8,
	23, 8,
	-2, 75,
	-1, 165,
	21, 8,
	22, 8,
	23, 8,
	-2, 72,
	-1, 171,
	25, 76,
	32, 76,
	-2, 8,
	-1, 174,
	25, 73,
	32, 73,
	-2, 8,
}

const RubyNprod = 86
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 393

var RubyAct = []int{

	42, 110, 124, 71, 4, 83, 95, 44, 65, 73,
	137, 126, 98, 55, 56, 2, 43, 160, 158, 115,
	43, 167, 127, 133, 45, 113, 94, 41, 175, 46,
	47, 48, 173, 43, 49, 50, 51, 52, 53, 54,
	116, 170, 57, 43, 63, 43, 100, 43, 74, 74,
	64, 155, 76, 77, 78, 79, 145, 80, 75, 134,
	115, 86, 87, 88, 143, 89, 108, 90, 84, 85,
	91, 168, 92, 93, 60, 61, 62, 60, 61, 62,
	139, 99, 146, 96, 101, 60, 61, 62, 164, 102,
	103, 104, 105, 106, 72, 107, 74, 112, 67, 68,
	114, 69, 109, 97, 119, 111, 58, 59, 121, 122,
	123, 70, 66, 118, 67, 68, 129, 69, 130, 131,
	132, 45, 140, 141, 135, 136, 82, 81, 138, 117,
	1, 20, 19, 18, 147, 17, 142, 144, 16, 148,
	150, 15, 14, 151, 10, 153, 25, 9, 8, 152,
	156, 154, 26, 159, 11, 161, 162, 6, 7, 27,
	21, 13, 163, 22, 166, 165, 12, 0, 169, 0,
	0, 172, 171, 0, 0, 174, 5, 23, 24, 0,
	0, 0, 28, 157, 29, 30, 31, 32, 0, 0,
	0, 33, 34, 35, 36, 0, 0, 0, 0, 0,
	0, 37, 0, 38, 0, 40, 39, 5, 23, 24,
	0, 0, 0, 28, 149, 29, 30, 31, 32, 0,
	0, 0, 33, 34, 35, 36, 0, 0, 0, 0,
	0, 0, 37, 0, 38, 0, 40, 39, 5, 23,
	24, 0, 0, 0, 28, 128, 29, 30, 31, 32,
	0, 0, 0, 33, 34, 35, 36, 0, 0, 0,
	0, 0, 0, 37, 0, 38, 0, 40, 39, 5,
	23, 24, 0, 0, 0, 28, 125, 29, 30, 31,
	32, 0, 0, 0, 33, 34, 35, 36, 0, 0,
	0, 0, 0, 0, 37, 0, 38, 0, 40, 39,
	5, 23, 24, 0, 0, 0, 28, 0, 29, 30,
	31, 32, 0, 0, 0, 33, 34, 35, 36, 0,
	0, 3, 0, 0, 0, 37, 0, 38, 0, 40,
	39, 5, 23, 24, 0, 0, 0, 28, 0, 29,
	30, 31, 32, 0, 0, 0, 33, 34, 35, 36,
	0, 0, 0, 0, 0, 0, 37, 0, 38, 0,
	40, 39, 5, 120, 24, 0, 0, 0, 28, 0,
	29, 30, 31, 32, 0, 0, 0, 33, 34, 35,
	36, 0, 0, 0, 0, 0, 0, 37, 0, 38,
	0, 40, 39,
}
var RubyPact = []int{

	-20, 296, -1000, -1000, -8, -1000, 9, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 9, -1000, -1000, -1000, -1000, 9, 9,
	9, -1000, -1000, 9, 9, 9, 9, 9, 9, 8,
	101, -1000, 56, 9, 32, 94, 89, 9, 9, 327,
	327, 327, 327, 122, -1000, -1000, -1000, 63, -1000, -1000,
	9, 9, 9, -1000, 9, -1000, 9, -1000, -1000, 9,
	-1000, 9, 9, -10, 77, -24, 9, 9, 9, 9,
	9, -1000, -1000, 21, -1000, -1000, 327, 327, 327, 327,
	327, 122, 57, 110, -1000, 9, 9, -2, -1000, 10,
	-1000, 358, 9, 9, 9, 9, 9, 9, 9, 9,
	265, -25, -1000, -5, 234, 9, -1000, 9, 9, 9,
	-4, 51, 120, -26, 9, -1000, -1000, 74, -1000, 118,
	55, 47, 64, 9, -1000, -1000, -1000, -1000, 203, 9,
	-1000, -1000, 21, -1000, 21, -1000, 34, 327, 172, -1000,
	-1000, -14, 21, -15, 21, 9, 9, -1000, -1000, 327,
	-1000, 83, 327, 9, -6, 9, 53, 9, 24, 327,
	9, 23, 327, -1000, 19, -1000,
}
var RubyPgo = []int{

	0, 167, 2, 166, 163, 161, 160, 159, 158, 157,
	154, 152, 148, 147, 146, 144, 9, 142, 141, 138,
	135, 133, 132, 131, 1, 8, 130, 129, 113, 3,
	111, 103, 0, 5,
}
var RubyR1 = []int{

	0, 26, 26, 26, 26, 26, 24, 24, 32, 32,
	33, 33, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	9, 9, 9, 9, 9, 8, 25, 25, 25, 25,
	30, 30, 30, 30, 29, 29, 29, 29, 29, 12,
	13, 13, 15, 16, 16, 31, 31, 10, 10, 17,
	18, 19, 20, 21, 22, 23, 3, 5, 6, 4,
	4, 27, 27, 27, 27, 28, 28, 28, 1, 1,
	7, 7, 14, 14, 11, 11,
}
var RubyR2 = []int{

	0, 0, 1, 2, 2, 3, 0, 2, 0, 2,
	0, 2, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 3, 1, 1, 5, 1,
	1, 1, 5, 5, 0, 1, 1, 5, 5, 9,
	6, 8, 6, 3, 6, 1, 4, 5, 5, 3,
	3, 3, 3, 5, 5, 5, 1, 1, 5, 9,
	9, 0, 6, 11, 12, 4, 9, 10, 0, 1,
	2, 2, 2, 2, 3, 3,
}
var RubyChk = []int{

	-1000, -26, 35, 25, -2, 4, -9, -8, -12, -13,
	-15, -10, -3, -5, -17, -18, -19, -20, -21, -22,
	-23, -6, -4, 5, 6, -14, -11, -7, 10, 12,
	13, 14, 15, 19, 20, 21, 22, 29, 31, 34,
	33, 35, -32, 24, -32, -32, -32, -32, -32, -32,
	-32, -32, -32, -32, -32, 5, 6, 34, 5, 6,
	21, 22, 23, -32, 18, -25, 18, 4, 5, 7,
	-30, -29, 5, -16, -32, -16, -2, -2, -2, -2,
	-29, 5, 4, -33, 5, 6, -32, -32, -32, -32,
	-32, -32, -32, -32, 36, 16, 6, -31, 36, -32,
	25, -32, -2, -2, -2, -2, -2, -29, 9, -25,
	-24, -16, -32, 27, -24, 9, 30, -27, -28, -2,
	5, -32, -32, -32, -2, 11, 36, 27, 11, -32,
	-32, -32, -32, 27, 8, 4, 5, 36, -24, 6,
	4, 5, -33, 9, -33, 9, 18, -32, -24, 11,
	-32, -32, -33, -32, -33, 17, -2, 11, 32, -32,
	32, -32, -32, -2, 5, -2, -32, 27, 18, -32,
	17, -2, -32, 9, -2, 9,
}
var RubyDef = []int{

	1, -2, 2, 3, -2, 12, -2, 14, 15, 16,
	17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
	27, 28, 29, -2, 31, 32, 33, 34, 8, 8,
	8, 66, 67, 8, 8, 8, 8, 8, 8, 0,
	0, 5, 0, 8, 0, 44, 0, 8, 8, 0,
	0, 0, 0, 44, 10, 82, 83, 0, 80, 81,
	8, 8, 8, 9, 8, 35, 8, 36, 37, 8,
	39, 8, 8, 0, 0, 0, -2, -2, -2, -2,
	8, 45, 46, 8, 84, 85, 0, 0, 0, 0,
	0, 44, 0, 44, 6, 8, -2, 0, 6, 0,
	11, 71, -2, -2, -2, -2, -2, 8, 8, 8,
	0, 0, 53, 0, 0, 8, 68, 8, 8, 8,
	-2, 0, 0, 0, -2, 50, 6, 0, 52, 0,
	10, 10, 0, 8, 38, 42, 43, 6, 0, -2,
	47, 48, 8, 10, 8, 10, 0, 0, 0, 51,
	54, 0, 8, 0, 8, 8, -2, 49, 69, 0,
	70, 0, 0, 8, 0, -2, 0, 8, 0, 0,
	8, -2, 0, 77, -2, 74,
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
		//line parser.y:114
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:116
		{
			Statements = []ast.Node{}
		}
	case 3:
		//line parser.y:118
		{ /* do nothing */
		}
	case 4:
		//line parser.y:120
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 5:
		//line parser.y:126
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 6:
		//line parser.y:133
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 7:
		//line parser.y:135
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 34:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 35:
		//line parser.y:149
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 36:
		//line parser.y:157
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 37:
		//line parser.y:159
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 38:
		//line parser.y:161
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 39:
		//line parser.y:165
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 40:
		//line parser.y:170
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 41:
		//line parser.y:172
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 42:
		//line parser.y:174
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 43:
		//line parser.y:176
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 44:
		//line parser.y:178
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 45:
		//line parser.y:180
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 46:
		//line parser.y:182
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 47:
		//line parser.y:184
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 48:
		//line parser.y:186
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 49:
		//line parser.y:192
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 50:
		//line parser.y:201
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 51:
		//line parser.y:208
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 52:
		//line parser.y:218
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 53:
		//line parser.y:227
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 54:
		//line parser.y:233
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 55:
		//line parser.y:241
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 56:
		//line parser.y:245
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 57:
		//line parser.y:250
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 58:
		//line parser.y:257
		{ // FIXME: this is a hack. If this rule is not present, "foo = 5" will not parse
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 59:
		//line parser.y:264
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 60:
		//line parser.y:265
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 61:
		//line parser.y:266
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 62:
		//line parser.y:267
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 63:
		//line parser.y:270
		{
			RubyVAL.genericValue = ast.Addition{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 64:
		//line parser.y:278
		{
			RubyVAL.genericValue = ast.Subtraction{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 65:
		//line parser.y:286
		{
			RubyVAL.genericValue = ast.Multiplication{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 66:
		//line parser.y:293
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 67:
		//line parser.y:294
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 68:
		//line parser.y:296
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 69:
		//line parser.y:300
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 70:
		//line parser.y:308
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 71:
		//line parser.y:316
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 72:
		//line parser.y:318
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 73:
		//line parser.y:320
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 74:
		//line parser.y:322
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 75:
		//line parser.y:325
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 76:
		//line parser.y:332
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 77:
		//line parser.y:339
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 78:
		//line parser.y:347
		{
			RubyVAL.genericValue = nil
		}
	case 79:
		//line parser.y:348
		{
			RubyVAL.genericValue = nil
		}
	case 80:
		//line parser.y:351
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 81:
		//line parser.y:353
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 82:
		//line parser.y:356
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 83:
		//line parser.y:358
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 84:
		//line parser.y:361
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 85:
		//line parser.y:363
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	}
	goto Rubystack /* stack new state and value */
}
