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
const HASHROCKET = 57366
const COLON = 57367
const DOT = 57368
const LBRACKET = 57369
const RBRACKET = 57370
const LBRACE = 57371
const RBRACE = 57372

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
	"HASHROCKET",
	"COLON",
	"DOT",
	"LBRACKET",
	"RBRACKET",
	"LBRACE",
	"RBRACE",
}
var RubyStatenames = []string{}

const RubyEofCode = 1
const RubyErrCode = 2
const RubyMaxDepth = 200

//line parser.y:340

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 7,
	7, 56,
	9, 56,
	18, 56,
	-2, 12,
	-1, 42,
	9, 35,
	-2, 30,
	-1, 60,
	9, 37,
	-2, 31,
	-1, 61,
	9, 36,
	-2, 32,
	-1, 63,
	9, 56,
	-2, 34,
	-1, 69,
	18, 56,
	-2, 12,
	-1, 84,
	9, 35,
	-2, 30,
	-1, 87,
	25, 46,
	-2, 56,
	-1, 111,
	18, 56,
	-2, 12,
	-1, 130,
	25, 47,
	-2, 56,
	-1, 145,
	21, 56,
	22, 56,
	23, 56,
	-2, 72,
	-1, 151,
	21, 56,
	22, 56,
	23, 56,
	-2, 68,
	-1, 159,
	30, 73,
	31, 73,
	-2, 56,
	-1, 160,
	30, 69,
	31, 69,
	-2, 56,
}

const RubyNprod = 77
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 397

var RubyAct = []int{

	90, 58, 6, 65, 40, 41, 76, 124, 42, 134,
	162, 86, 63, 155, 26, 27, 26, 27, 153, 161,
	149, 109, 147, 109, 132, 101, 85, 43, 44, 45,
	46, 77, 48, 26, 27, 49, 50, 51, 52, 53,
	54, 4, 26, 27, 109, 77, 77, 66, 128, 66,
	117, 89, 67, 99, 118, 104, 78, 79, 80, 59,
	81, 125, 99, 82, 83, 84, 73, 55, 56, 57,
	91, 99, 106, 130, 92, 126, 127, 93, 55, 56,
	57, 156, 55, 56, 57, 135, 100, 66, 103, 87,
	102, 68, 70, 71, 72, 98, 75, 74, 64, 112,
	113, 114, 6, 47, 60, 61, 6, 62, 120, 121,
	122, 123, 91, 60, 61, 105, 62, 115, 59, 2,
	94, 95, 96, 97, 88, 136, 108, 131, 133, 107,
	6, 139, 140, 1, 142, 110, 144, 23, 6, 141,
	22, 143, 148, 129, 150, 21, 20, 19, 18, 17,
	13, 11, 10, 154, 137, 14, 157, 158, 3, 12,
	24, 16, 25, 15, 8, 7, 9, 0, 0, 0,
	28, 146, 29, 31, 32, 33, 0, 0, 145, 34,
	35, 36, 37, 0, 0, 30, 151, 38, 0, 39,
	152, 5, 26, 27, 0, 8, 7, 9, 0, 159,
	160, 28, 138, 29, 31, 32, 33, 0, 0, 0,
	34, 35, 36, 37, 0, 0, 30, 0, 38, 0,
	39, 0, 5, 26, 27, 8, 7, 9, 0, 0,
	0, 28, 119, 29, 31, 32, 33, 0, 0, 0,
	34, 35, 36, 37, 0, 0, 30, 0, 38, 0,
	39, 0, 5, 26, 27, 8, 7, 9, 0, 0,
	0, 28, 116, 29, 31, 32, 33, 0, 0, 0,
	34, 35, 36, 37, 0, 0, 30, 0, 38, 0,
	39, 0, 5, 26, 27, 8, 7, 9, 0, 0,
	0, 28, 0, 29, 31, 32, 33, 0, 0, 0,
	34, 35, 36, 37, 0, 0, 30, 0, 38, 0,
	39, 0, 5, 26, 27, 8, 69, 9, 0, 0,
	0, 28, 0, 29, 31, 32, 33, 0, 0, 0,
	34, 35, 36, 37, 0, 0, 30, 0, 38, 0,
	39, 0, 109, 8, 111, 9, 0, 0, 0, 28,
	0, 29, 31, 32, 33, 0, 0, 0, 34, 35,
	36, 37, 0, 0, 30, 0, 38, 0, 39, 0,
	109, 8, 69, 9, 0, 0, 0, 28, 0, 29,
	31, 32, 33, 0, 0, 0, 34, 35, 36, 37,
	0, 0, 30, 0, 38, 0, 39,
}
var RubyPact = []int{

	-1000, 281, -1000, -16, -16, -1000, -1000, -16, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -16, -16, -16, -16,
	98, -16, -1000, -1000, -16, -16, -16, -16, -16, -16,
	-1000, 46, 100, -1000, -1000, 93, -16, -1000, -16, 367,
	367, 367, 367, 92, 14, -16, -16, -16, -1000, -16,
	-1000, -1000, -16, -16, -16, -5, 83, 20, -16, -16,
	-16, -16, -16, -16, -1000, -1000, -16, -1000, 367, 367,
	367, 367, 92, 62, 109, -1000, -16, -16, 30, -1000,
	46, 41, 44, 339, -16, -16, -16, -16, -16, -16,
	-16, 251, 19, -1000, 29, 221, -1000, -16, -16, -16,
	-16, -18, 53, 71, 17, -1000, -1000, -1000, 67, -1000,
	15, 0, -1000, 61, -16, -1000, -1000, -1000, -1000, 191,
	-16, -16, 14, -16, 14, -16, 367, 160, -1000, -1000,
	-8, -16, -10, -16, 367, -16, -1000, -1000, 311, -1000,
	13, -16, -16, -12, 57, -16, -16, 367, 367, 10,
	1, -1000, -1000,
}
var RubyPgo = []int{

	0, 0, 6, 41, 163, 162, 161, 160, 159, 158,
	117, 155, 152, 151, 150, 3, 149, 148, 147, 146,
	145, 140, 137, 25, 1, 133, 129, 126, 12, 124,
	124,
}
var RubyR1 = []int{

	0, 25, 25, 23, 23, 8, 10, 10, 10, 10,
	9, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	24, 24, 24, 24, 24, 28, 28, 28, 28, 28,
	12, 13, 13, 14, 15, 15, 29, 29, 11, 16,
	17, 18, 19, 20, 21, 22, 1, 1, 1, 4,
	6, 7, 5, 5, 2, 2, 2, 26, 26, 26,
	26, 27, 27, 27, 27, 30, 30,
}
var RubyR2 = []int{

	0, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	0, 1, 1, 5, 1, 0, 1, 1, 5, 5,
	9, 6, 8, 6, 3, 6, 1, 4, 5, 3,
	3, 3, 3, 5, 5, 5, 0, 2, 2, 1,
	1, 5, 9, 9, 0, 1, 4, 0, 5, 10,
	11, 0, 4, 9, 10, 0, 1,
}
var RubyChk = []int{

	-1000, -25, -10, -9, -3, 31, -1, 5, 4, 6,
	-12, -13, -8, -14, -11, -4, -6, -16, -17, -18,
	-19, -20, -21, -22, -7, -5, 32, 33, 10, 12,
	25, 13, 14, 15, 19, 20, 21, 22, 27, 29,
	-1, -1, -1, -1, -1, -1, -1, 5, -1, -1,
	-1, -1, -1, -1, -1, 21, 22, 23, -24, 18,
	4, 5, 7, -28, 5, -15, -1, -15, -3, 5,
	-3, -3, -3, -28, 5, 4, -2, 31, -1, -1,
	-1, -1, -1, -1, -1, 31, 16, 6, -29, 31,
	-1, -1, -1, -1, -3, -3, -3, -3, -28, 9,
	-24, -23, -15, -1, 25, -23, 28, -26, -27, 31,
	-3, 5, -1, -1, -1, -10, 11, 31, 25, 11,
	-1, -1, -1, -1, 25, 8, 4, 5, 31, -23,
	6, -2, 9, -2, 9, 24, -1, -23, 11, -1,
	-1, -2, -1, -2, -1, -3, 11, 30, -1, 30,
	-1, -3, -3, 5, -1, 25, 24, -1, -1, -3,
	-3, 9, 9,
}
var RubyDef = []int{

	1, -2, 2, 56, 56, 8, 9, -2, 11, 13,
	14, 15, 16, 17, 18, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 28, 29, 56, 56, 56, 56,
	0, 56, 59, 60, 56, 56, 56, 56, 56, 56,
	6, 7, -2, 57, 58, 0, 56, 5, 56, 0,
	0, 0, 0, 35, 64, 56, 56, 56, 10, 56,
	-2, -2, 56, -2, 56, 0, 0, 0, 49, -2,
	50, 51, 52, 56, 36, 37, 56, 65, 0, 0,
	0, 0, 35, 0, -2, 3, 56, -2, 0, 3,
	0, 0, 0, 67, 53, 54, 55, 48, 56, 56,
	56, 0, 0, 44, 0, 0, 61, 56, 56, 56,
	56, -2, 0, 0, 0, 4, 41, 3, 0, 43,
	64, 64, 66, 0, 56, 33, 38, 39, 3, 0,
	-2, 56, 64, 56, 64, 56, 0, 0, 42, 45,
	0, 56, 0, 56, 0, -2, 40, 62, 0, 63,
	0, -2, 56, 0, 0, 56, 56, 0, 0, -2,
	-2, 74, 70,
}
var RubyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 33,
	31, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 32,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30,
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
		//line parser.y:107
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:109
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 3:
		//line parser.y:116
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 4:
		//line parser.y:118
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
			}
		}
	case 5:
		//line parser.y:125
		{
			RubyVAL.genericValue = ast.Symbol{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 6:
		//line parser.y:130
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 7:
		//line parser.y:132
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 8:
		//line parser.y:134
		{
			RubyVAL.genericValue = nil /* ignores new lines around statements */
		}
	case 9:
		//line parser.y:136
		{
			RubyVAL.genericValue = nil /* ignores whitespace around statements */
		}
	case 10:
		//line parser.y:139
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
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
		//line parser.y:148
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 31:
		//line parser.y:150
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 32:
		//line parser.y:152
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 33:
		//line parser.y:154
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 34:
		//line parser.y:158
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 35:
		//line parser.y:162
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 36:
		//line parser.y:164
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 37:
		//line parser.y:166
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 38:
		//line parser.y:168
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 39:
		//line parser.y:170
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 40:
		//line parser.y:175
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 41:
		//line parser.y:184
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 42:
		//line parser.y:191
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 43:
		//line parser.y:201
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 44:
		//line parser.y:210
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 45:
		//line parser.y:216
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 46:
		//line parser.y:224
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 47:
		//line parser.y:228
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 48:
		//line parser.y:233
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 49:
		//line parser.y:240
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 50:
		//line parser.y:241
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 51:
		//line parser.y:242
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 52:
		//line parser.y:243
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 53:
		//line parser.y:246
		{
			RubyVAL.genericValue = ast.Addition{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 54:
		//line parser.y:254
		{
			RubyVAL.genericValue = ast.Subtraction{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 55:
		//line parser.y:262
		{
			RubyVAL.genericValue = ast.Multiplication{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 56:
		//line parser.y:270
		{
			RubyVAL.genericValue = nil
		}
	case 57:
		//line parser.y:272
		{
			RubyVAL.genericValue = nil
		}
	case 58:
		//line parser.y:274
		{
			RubyVAL.genericValue = nil
		}
	case 59:
		//line parser.y:276
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 60:
		//line parser.y:277
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 61:
		//line parser.y:279
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 62:
		//line parser.y:283
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 63:
		//line parser.y:291
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 64:
		//line parser.y:300
		{
			RubyVAL.genericValue = nil
		}
	case 65:
		//line parser.y:302
		{
			RubyVAL.genericValue = nil
		}
	case 66:
		//line parser.y:304
		{
			RubyVAL.genericValue = nil
		}
	case 67:
		//line parser.y:306
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 68:
		//line parser.y:308
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 69:
		//line parser.y:310
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 70:
		//line parser.y:312
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 71:
		//line parser.y:314
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 72:
		//line parser.y:316
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 73:
		//line parser.y:323
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 74:
		//line parser.y:330
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	}
	goto Rubystack /* stack new state and value */
}
