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

//line parser.y:361

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	23, 6,
	-2, 2,
	-1, 4,
	18, 6,
	-2, 11,
	-1, 21,
	4, 6,
	5, 6,
	7, 6,
	9, 6,
	18, 6,
	-2, 28,
	-1, 74,
	21, 6,
	22, 6,
	23, 6,
	-2, 57,
	-1, 75,
	21, 6,
	22, 6,
	23, 6,
	-2, 58,
	-1, 76,
	21, 6,
	22, 6,
	23, 6,
	-2, 59,
	-1, 77,
	21, 6,
	22, 6,
	23, 6,
	-2, 60,
	-1, 94,
	27, 53,
	-2, 6,
	-1, 100,
	21, 6,
	22, 6,
	23, 6,
	-2, 61,
	-1, 101,
	21, 6,
	22, 6,
	23, 6,
	-2, 62,
	-1, 102,
	21, 6,
	22, 6,
	23, 6,
	-2, 63,
	-1, 103,
	21, 6,
	22, 6,
	23, 6,
	-2, 56,
	-1, 104,
	21, 6,
	22, 6,
	23, 6,
	-2, 55,
	-1, 118,
	21, 28,
	22, 28,
	23, 28,
	-2, 6,
	-1, 122,
	23, 6,
	-2, 5,
	-1, 137,
	27, 54,
	-2, 6,
	-1, 154,
	21, 6,
	22, 6,
	23, 6,
	-2, 73,
	-1, 163,
	21, 6,
	22, 6,
	23, 6,
	-2, 70,
	-1, 169,
	25, 74,
	32, 74,
	-2, 6,
	-1, 172,
	25, 71,
	32, 71,
	-2, 6,
}

const RubyNprod = 84
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 360

var RubyAct = []int{

	40, 108, 122, 69, 2, 42, 93, 71, 63, 135,
	124, 96, 53, 54, 158, 81, 156, 113, 41, 41,
	165, 131, 43, 125, 111, 41, 92, 44, 45, 46,
	39, 62, 47, 48, 49, 50, 51, 52, 114, 143,
	166, 55, 61, 58, 59, 60, 72, 72, 41, 98,
	74, 75, 76, 77, 73, 78, 58, 59, 60, 84,
	85, 86, 173, 87, 168, 88, 171, 153, 89, 144,
	90, 91, 58, 59, 60, 132, 113, 41, 141, 97,
	106, 41, 99, 65, 66, 162, 67, 100, 101, 102,
	103, 104, 137, 105, 72, 110, 94, 64, 112, 70,
	107, 109, 117, 82, 83, 95, 119, 120, 121, 65,
	66, 68, 67, 116, 127, 115, 128, 129, 130, 43,
	56, 57, 138, 139, 133, 134, 136, 80, 79, 1,
	18, 17, 145, 16, 15, 14, 13, 146, 148, 12,
	8, 149, 23, 151, 140, 142, 7, 6, 154, 24,
	9, 157, 4, 159, 160, 5, 25, 150, 19, 152,
	161, 11, 164, 163, 20, 10, 167, 0, 0, 170,
	169, 0, 0, 172, 3, 21, 22, 0, 0, 0,
	26, 155, 27, 28, 29, 30, 0, 0, 0, 31,
	32, 33, 34, 0, 0, 0, 0, 0, 0, 35,
	0, 36, 0, 38, 37, 3, 21, 22, 0, 0,
	0, 26, 147, 27, 28, 29, 30, 0, 0, 0,
	31, 32, 33, 34, 0, 0, 0, 0, 0, 0,
	35, 0, 36, 0, 38, 37, 3, 21, 22, 0,
	0, 0, 26, 126, 27, 28, 29, 30, 0, 0,
	0, 31, 32, 33, 34, 0, 0, 0, 0, 0,
	0, 35, 0, 36, 0, 38, 37, 3, 21, 22,
	0, 0, 0, 26, 123, 27, 28, 29, 30, 0,
	0, 0, 31, 32, 33, 34, 0, 0, 0, 0,
	0, 0, 35, 0, 36, 0, 38, 37, 3, 21,
	22, 0, 0, 0, 26, 0, 27, 28, 29, 30,
	0, 0, 0, 31, 32, 33, 34, 0, 0, 0,
	0, 0, 0, 35, 0, 36, 0, 38, 37, 3,
	118, 22, 0, 0, 0, 26, 0, 27, 28, 29,
	30, 0, 0, 0, 31, 32, 33, 34, 0, 0,
	0, 0, 0, 0, 35, 0, 36, 0, 38, 37,
}
var RubyPact = []int{

	-1000, 294, -5, -1000, 1, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 1, -1000, -1000, -1000, -1000, 1, 1, 1, -1000,
	-1000, 1, 1, 1, 1, 1, 1, 7, 115, -1000,
	35, 1, 13, 79, 94, 1, 1, 294, 294, 294,
	294, 123, -1000, -1000, -1000, 98, -1000, -1000, 1, 1,
	1, -1000, 1, -1000, 1, -1000, -1000, 1, -1000, 1,
	1, -10, 90, -25, 1, 1, 1, 1, 1, -1000,
	-1000, 24, -1000, -1000, 294, 294, 294, 294, 294, 123,
	71, 105, -1000, 1, 1, -3, -1000, 8, -1000, 325,
	1, 1, 1, 1, 1, 1, 1, 1, 263, -26,
	-1000, -4, 232, 1, -1000, 1, 1, 1, -6, 67,
	120, -27, 1, -1000, -1000, 86, -1000, 118, 69, 30,
	51, 1, -1000, -1000, -1000, -1000, 201, 1, -1000, -1000,
	24, -1000, 24, -1000, 50, 294, 170, -1000, -1000, -16,
	24, -18, 24, 1, 1, -1000, -1000, 294, -1000, 80,
	294, 1, -7, 1, 22, 1, 47, 294, 1, 57,
	294, -1000, 53, -1000,
}
var RubyPgo = []int{

	0, 167, 2, 165, 164, 161, 158, 156, 155, 152,
	150, 149, 147, 146, 142, 140, 7, 139, 136, 135,
	134, 133, 131, 130, 1, 8, 129, 115, 113, 3,
	111, 105, 0, 15,
}
var RubyR1 = []int{

	0, 26, 26, 26, 24, 24, 32, 32, 33, 33,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 9, 9,
	9, 9, 9, 8, 25, 25, 25, 25, 30, 30,
	30, 30, 29, 29, 29, 29, 29, 12, 13, 13,
	15, 16, 16, 31, 31, 10, 10, 17, 18, 19,
	20, 21, 22, 23, 3, 5, 6, 4, 4, 27,
	27, 27, 27, 28, 28, 28, 1, 1, 7, 7,
	14, 14, 11, 11,
}
var RubyR2 = []int{

	0, 0, 2, 3, 0, 2, 0, 2, 0, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 1, 1, 5, 1, 1, 1,
	5, 5, 0, 1, 1, 5, 5, 9, 6, 8,
	6, 3, 6, 1, 4, 5, 5, 3, 3, 3,
	3, 5, 5, 5, 1, 1, 5, 9, 9, 0,
	6, 11, 12, 4, 9, 10, 0, 1, 2, 2,
	2, 2, 3, 3,
}
var RubyChk = []int{

	-1000, -26, -2, 4, -9, -8, -12, -13, -15, -10,
	-3, -5, -17, -18, -19, -20, -21, -22, -23, -6,
	-4, 5, 6, -14, -11, -7, 10, 12, 13, 14,
	15, 19, 20, 21, 22, 29, 31, 34, 33, 35,
	-32, 24, -32, -32, -32, -32, -32, -32, -32, -32,
	-32, -32, -32, 5, 6, 34, 5, 6, 21, 22,
	23, -32, 18, -25, 18, 4, 5, 7, -30, -29,
	5, -16, -32, -16, -2, -2, -2, -2, -29, 5,
	4, -33, 5, 6, -32, -32, -32, -32, -32, -32,
	-32, -32, 36, 16, 6, -31, 36, -32, 25, -32,
	-2, -2, -2, -2, -2, -29, 9, -25, -24, -16,
	-32, 27, -24, 9, 30, -27, -28, -2, 5, -32,
	-32, -32, -2, 11, 36, 27, 11, -32, -32, -32,
	-32, 27, 8, 4, 5, 36, -24, 6, 4, 5,
	-33, 9, -33, 9, 18, -32, -24, 11, -32, -32,
	-33, -32, -33, 17, -2, 11, 32, -32, 32, -32,
	-32, -2, 5, -2, -32, 27, 18, -32, 17, -2,
	-32, 9, -2, 9,
}
var RubyDef = []int{

	1, -2, -2, 10, -2, 12, 13, 14, 15, 16,
	17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
	27, -2, 29, 30, 31, 32, 6, 6, 6, 64,
	65, 6, 6, 6, 6, 6, 6, 0, 0, 3,
	0, 6, 0, 42, 0, 6, 6, 0, 0, 0,
	0, 42, 8, 80, 81, 0, 78, 79, 6, 6,
	6, 7, 6, 33, 6, 34, 35, 6, 37, 6,
	6, 0, 0, 0, -2, -2, -2, -2, 6, 43,
	44, 6, 82, 83, 0, 0, 0, 0, 0, 42,
	0, 42, 4, 6, -2, 0, 4, 0, 9, 69,
	-2, -2, -2, -2, -2, 6, 6, 6, 0, 0,
	51, 0, 0, 6, 66, 6, 6, 6, -2, 0,
	0, 0, -2, 48, 4, 0, 50, 0, 8, 8,
	0, 6, 36, 40, 41, 4, 0, -2, 45, 46,
	6, 8, 6, 8, 0, 0, 0, 49, 52, 0,
	6, 0, 6, 6, -2, 47, 67, 0, 68, 0,
	0, 6, 0, -2, 0, 6, 0, 0, 6, -2,
	0, 75, -2, 72,
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
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 3:
		//line parser.y:122
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 4:
		//line parser.y:129
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 5:
		//line parser.y:131
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
		//line parser.y:145
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 34:
		//line parser.y:153
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 35:
		//line parser.y:155
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 36:
		//line parser.y:157
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 37:
		//line parser.y:161
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 38:
		//line parser.y:166
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 39:
		//line parser.y:168
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
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
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 43:
		//line parser.y:176
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 44:
		//line parser.y:178
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
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
		//line parser.y:188
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 48:
		//line parser.y:197
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 49:
		//line parser.y:204
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 50:
		//line parser.y:214
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 51:
		//line parser.y:223
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 52:
		//line parser.y:229
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 53:
		//line parser.y:237
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 54:
		//line parser.y:241
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 55:
		//line parser.y:246
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 56:
		//line parser.y:253
		{ // FIXME: this is a hack. If this rule is not present, "foo = 5" will not parse
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 57:
		//line parser.y:260
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 58:
		//line parser.y:261
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 59:
		//line parser.y:262
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 60:
		//line parser.y:263
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 61:
		//line parser.y:266
		{
			RubyVAL.genericValue = ast.Addition{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 62:
		//line parser.y:274
		{
			RubyVAL.genericValue = ast.Subtraction{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 63:
		//line parser.y:282
		{
			RubyVAL.genericValue = ast.Multiplication{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 64:
		//line parser.y:289
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 65:
		//line parser.y:290
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 66:
		//line parser.y:292
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 67:
		//line parser.y:296
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 68:
		//line parser.y:304
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 69:
		//line parser.y:312
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 70:
		//line parser.y:314
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 71:
		//line parser.y:316
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 72:
		//line parser.y:318
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 73:
		//line parser.y:321
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 74:
		//line parser.y:328
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 75:
		//line parser.y:335
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 76:
		//line parser.y:343
		{
			RubyVAL.genericValue = nil
		}
	case 77:
		//line parser.y:344
		{
			RubyVAL.genericValue = nil
		}
	case 78:
		//line parser.y:347
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 79:
		//line parser.y:349
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 80:
		//line parser.y:352
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 81:
		//line parser.y:354
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 82:
		//line parser.y:357
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 83:
		//line parser.y:359
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	}
	goto Rubystack /* stack new state and value */
}
