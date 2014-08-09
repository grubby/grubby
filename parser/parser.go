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
const DOLLARSIGN = 57373

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
	"DOLLARSIGN",
}
var RubyStatenames = []string{}

const RubyEofCode = 1
const RubyErrCode = 2
const RubyMaxDepth = 200

//line parser.y:347

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 7,
	7, 57,
	9, 57,
	18, 57,
	-2, 12,
	-1, 44,
	9, 36,
	-2, 31,
	-1, 64,
	9, 38,
	-2, 32,
	-1, 65,
	9, 37,
	-2, 33,
	-1, 67,
	9, 57,
	-2, 35,
	-1, 73,
	18, 57,
	-2, 12,
	-1, 88,
	9, 36,
	-2, 31,
	-1, 91,
	25, 47,
	-2, 57,
	-1, 115,
	18, 57,
	-2, 12,
	-1, 134,
	25, 48,
	-2, 57,
	-1, 149,
	21, 57,
	22, 57,
	23, 57,
	-2, 73,
	-1, 155,
	21, 57,
	22, 57,
	23, 57,
	-2, 69,
	-1, 163,
	30, 74,
	32, 74,
	-2, 57,
	-1, 164,
	30, 70,
	32, 70,
	-2, 57,
}

const RubyNprod = 80
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 403

var RubyAct = []int{

	94, 62, 6, 4, 42, 43, 80, 67, 44, 138,
	166, 128, 27, 28, 153, 81, 113, 105, 165, 27,
	28, 157, 151, 136, 113, 132, 69, 121, 45, 46,
	47, 48, 81, 50, 27, 28, 51, 52, 53, 54,
	55, 56, 27, 28, 93, 159, 81, 122, 113, 70,
	90, 70, 103, 108, 63, 72, 74, 75, 76, 103,
	82, 83, 84, 77, 85, 134, 89, 86, 87, 88,
	68, 110, 129, 103, 95, 57, 58, 71, 96, 130,
	131, 97, 59, 60, 61, 160, 98, 99, 100, 101,
	104, 70, 107, 91, 102, 59, 60, 61, 139, 64,
	65, 114, 66, 116, 117, 118, 6, 59, 60, 61,
	6, 109, 124, 125, 126, 127, 95, 106, 49, 64,
	65, 92, 66, 79, 78, 119, 112, 2, 111, 140,
	1, 135, 137, 63, 6, 143, 144, 23, 146, 133,
	148, 22, 6, 145, 149, 147, 152, 21, 154, 20,
	141, 19, 155, 18, 17, 13, 156, 158, 11, 10,
	161, 162, 8, 7, 9, 163, 164, 14, 29, 150,
	30, 32, 33, 34, 3, 12, 26, 35, 36, 37,
	38, 24, 16, 31, 25, 39, 15, 40, 0, 41,
	5, 27, 28, 8, 7, 9, 0, 0, 0, 29,
	142, 30, 32, 33, 34, 0, 0, 0, 35, 36,
	37, 38, 0, 0, 31, 0, 39, 0, 40, 0,
	41, 5, 27, 28, 8, 7, 9, 0, 0, 0,
	29, 123, 30, 32, 33, 34, 0, 0, 0, 35,
	36, 37, 38, 0, 0, 31, 0, 39, 0, 40,
	0, 41, 5, 27, 28, 8, 7, 9, 0, 0,
	0, 29, 120, 30, 32, 33, 34, 0, 0, 0,
	35, 36, 37, 38, 0, 0, 31, 0, 39, 0,
	40, 0, 41, 5, 27, 28, 8, 7, 9, 0,
	0, 0, 29, 0, 30, 32, 33, 34, 0, 0,
	0, 35, 36, 37, 38, 0, 0, 31, 0, 39,
	0, 40, 0, 41, 5, 27, 28, 8, 73, 9,
	0, 0, 0, 29, 0, 30, 32, 33, 34, 0,
	0, 0, 35, 36, 37, 38, 0, 0, 31, 0,
	39, 0, 40, 0, 41, 113, 8, 115, 9, 0,
	0, 0, 29, 0, 30, 32, 33, 34, 0, 0,
	0, 35, 36, 37, 38, 0, 0, 31, 0, 39,
	0, 40, 0, 41, 113, 8, 73, 9, 0, 0,
	0, 29, 0, 30, 32, 33, 34, 0, 0, 0,
	35, 36, 37, 38, 0, 0, 31, 0, 39, 0,
	40, 0, 41,
}
var RubyPact = []int{

	-1000, 282, -1000, -21, -21, -1000, -1000, -21, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -21, -21, -21,
	-21, 113, -21, -1000, -1000, -21, -21, -21, -21, -21,
	-21, 70, -1000, 86, 115, -1000, -1000, 65, -21, -1000,
	-21, 371, 371, 371, 371, 119, -17, -1000, -1000, -21,
	-21, -21, -1000, -21, -1000, -1000, -21, -21, -21, 34,
	87, 12, -21, -21, -21, -21, -21, -21, -1000, -1000,
	-21, -1000, 371, 371, 371, 371, 119, 50, 95, -1000,
	-21, -21, 28, -1000, 86, 36, 43, 342, -21, -21,
	-21, -21, -21, -21, -21, 251, -5, -1000, 22, 220,
	-1000, -21, -21, -21, -21, -14, 64, 75, -7, -1000,
	-1000, -1000, 59, -1000, 14, 0, -1000, 74, -21, -1000,
	-1000, -1000, -1000, 189, -21, -21, -17, -21, -17, -21,
	371, 158, -1000, -1000, -8, -21, -16, -21, 371, -21,
	-1000, -1000, 313, -1000, 16, -21, -21, 20, 61, -21,
	-21, 371, 371, 9, 1, -1000, -1000,
}
var RubyPgo = []int{

	0, 0, 6, 3, 186, 184, 182, 181, 176, 175,
	174, 125, 167, 159, 158, 155, 26, 154, 153, 151,
	149, 147, 141, 137, 17, 1, 130, 128, 126, 7,
	121, 121,
}
var RubyR1 = []int{

	0, 26, 26, 24, 24, 9, 11, 11, 11, 11,
	10, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 25, 25, 25, 25, 25, 29, 29, 29, 29,
	29, 13, 14, 14, 15, 16, 16, 30, 30, 12,
	17, 18, 19, 20, 21, 22, 23, 1, 1, 1,
	4, 6, 7, 5, 5, 2, 2, 2, 27, 27,
	27, 27, 28, 28, 28, 28, 31, 31, 8, 8,
}
var RubyR2 = []int{

	0, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 0, 1, 1, 5, 1, 0, 1, 1, 5,
	5, 9, 6, 8, 6, 3, 6, 1, 4, 5,
	3, 3, 3, 3, 5, 5, 5, 0, 2, 2,
	1, 1, 5, 9, 9, 0, 1, 4, 0, 5,
	10, 11, 0, 4, 9, 10, 0, 1, 2, 2,
}
var RubyChk = []int{

	-1000, -26, -11, -10, -3, 32, -1, 5, 4, 6,
	-13, -14, -9, -15, -12, -4, -6, -17, -18, -19,
	-20, -21, -22, -23, -7, -5, -8, 33, 34, 10,
	12, 25, 13, 14, 15, 19, 20, 21, 22, 27,
	29, 31, -1, -1, -1, -1, -1, -1, -1, 5,
	-1, -1, -1, -1, -1, -1, -1, 5, 6, 21,
	22, 23, -25, 18, 4, 5, 7, -29, 5, -16,
	-1, -16, -3, 5, -3, -3, -3, -29, 5, 4,
	-2, 32, -1, -1, -1, -1, -1, -1, -1, 32,
	16, 6, -30, 32, -1, -1, -1, -1, -3, -3,
	-3, -3, -29, 9, -25, -24, -16, -1, 25, -24,
	28, -27, -28, 32, -3, 5, -1, -1, -1, -11,
	11, 32, 25, 11, -1, -1, -1, -1, 25, 8,
	4, 5, 32, -24, 6, -2, 9, -2, 9, 24,
	-1, -24, 11, -1, -1, -2, -1, -2, -1, -3,
	11, 30, -1, 30, -1, -3, -3, 5, -1, 25,
	24, -1, -1, -3, -3, 9, 9,
}
var RubyDef = []int{

	1, -2, 2, 57, 57, 8, 9, -2, 11, 13,
	14, 15, 16, 17, 18, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 28, 29, 30, 57, 57, 57,
	57, 0, 57, 60, 61, 57, 57, 57, 57, 57,
	57, 0, 6, 7, -2, 58, 59, 0, 57, 5,
	57, 0, 0, 0, 0, 36, 65, 78, 79, 57,
	57, 57, 10, 57, -2, -2, 57, -2, 57, 0,
	0, 0, 50, -2, 51, 52, 53, 57, 37, 38,
	57, 66, 0, 0, 0, 0, 36, 0, -2, 3,
	57, -2, 0, 3, 0, 0, 0, 68, 54, 55,
	56, 49, 57, 57, 57, 0, 0, 45, 0, 0,
	62, 57, 57, 57, 57, -2, 0, 0, 0, 4,
	42, 3, 0, 44, 65, 65, 67, 0, 57, 34,
	39, 40, 3, 0, -2, 57, 65, 57, 65, 57,
	0, 0, 43, 46, 0, 57, 0, 57, 0, -2,
	41, 63, 0, 64, 0, -2, 57, 0, 0, 57,
	57, 0, 0, -2, -2, 75, 71,
}
var RubyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 34,
	32, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 33,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
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
		//line parser.y:109
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:111
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 3:
		//line parser.y:118
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 4:
		//line parser.y:120
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
			}
		}
	case 5:
		//line parser.y:127
		{
			RubyVAL.genericValue = ast.Symbol{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 6:
		//line parser.y:132
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 7:
		//line parser.y:134
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 8:
		//line parser.y:136
		{
			RubyVAL.genericValue = nil /* ignores new lines around statements */
		}
	case 9:
		//line parser.y:138
		{
			RubyVAL.genericValue = nil /* ignores whitespace around statements */
		}
	case 10:
		//line parser.y:141
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 31:
		//line parser.y:150
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 32:
		//line parser.y:152
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 33:
		//line parser.y:154
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 34:
		//line parser.y:156
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 35:
		//line parser.y:160
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 36:
		//line parser.y:164
		{
			RubyVAL.genericSlice = ast.Nodes{}
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
		//line parser.y:172
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 41:
		//line parser.y:177
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 42:
		//line parser.y:186
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 43:
		//line parser.y:193
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 44:
		//line parser.y:203
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 45:
		//line parser.y:212
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 46:
		//line parser.y:218
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 47:
		//line parser.y:226
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 48:
		//line parser.y:230
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 49:
		//line parser.y:235
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 50:
		//line parser.y:242
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 51:
		//line parser.y:243
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 52:
		//line parser.y:244
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 53:
		//line parser.y:245
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 54:
		//line parser.y:248
		{
			RubyVAL.genericValue = ast.Addition{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 55:
		//line parser.y:256
		{
			RubyVAL.genericValue = ast.Subtraction{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 56:
		//line parser.y:264
		{
			RubyVAL.genericValue = ast.Multiplication{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
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
			RubyVAL.genericValue = nil
		}
	case 60:
		//line parser.y:278
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 61:
		//line parser.y:279
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 62:
		//line parser.y:281
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 63:
		//line parser.y:285
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 64:
		//line parser.y:293
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
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
			RubyVAL.genericValue = nil
		}
	case 68:
		//line parser.y:308
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 69:
		//line parser.y:310
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 70:
		//line parser.y:312
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 71:
		//line parser.y:314
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 72:
		//line parser.y:316
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 73:
		//line parser.y:318
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 74:
		//line parser.y:325
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 75:
		//line parser.y:332
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 78:
		//line parser.y:343
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 79:
		//line parser.y:345
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	}
	goto Rubystack /* stack new state and value */
}
