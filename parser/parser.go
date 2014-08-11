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
const COLON = 57366
const DOT = 57367
const LBRACKET = 57368
const RBRACKET = 57369
const LBRACE = 57370
const RBRACE = 57371
const DOLLARSIGN = 57372
const ATSIGN = 57373
const EOF = 57374

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

//line parser.y:374

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 62,
	-1, 8,
	7, 62,
	9, 62,
	-2, 30,
	-1, 50,
	9, 40,
	-2, 35,
	-1, 74,
	9, 42,
	-2, 36,
	-1, 75,
	9, 41,
	-2, 37,
	-1, 77,
	9, 62,
	-2, 39,
	-1, 102,
	9, 40,
	-2, 35,
	-1, 105,
	24, 51,
	-2, 62,
	-1, 149,
	24, 52,
	-2, 62,
	-1, 164,
	21, 62,
	22, 62,
	23, 62,
	-2, 78,
	-1, 173,
	21, 62,
	22, 62,
	23, 62,
	-2, 74,
	-1, 179,
	29, 79,
	33, 79,
	-2, 62,
	-1, 182,
	29, 75,
	33, 75,
	-2, 62,
}

const RubyNprod = 89
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 452

var RubyAct = []int{

	108, 72, 7, 80, 120, 48, 49, 27, 28, 50,
	77, 143, 168, 92, 91, 104, 128, 183, 172, 181,
	166, 27, 28, 153, 128, 151, 147, 51, 52, 53,
	54, 55, 103, 57, 136, 107, 58, 59, 60, 61,
	62, 63, 27, 28, 27, 28, 128, 92, 117, 92,
	64, 65, 47, 2, 175, 137, 81, 123, 81, 176,
	73, 82, 69, 70, 71, 78, 125, 69, 70, 71,
	95, 96, 97, 88, 98, 178, 66, 99, 100, 101,
	102, 163, 144, 117, 117, 109, 5, 93, 94, 110,
	154, 149, 111, 69, 70, 71, 105, 74, 75, 79,
	76, 67, 68, 56, 119, 81, 122, 134, 121, 3,
	116, 73, 124, 74, 75, 106, 76, 131, 132, 127,
	133, 7, 145, 146, 126, 7, 1, 139, 140, 141,
	142, 109, 90, 89, 23, 22, 21, 20, 19, 18,
	17, 148, 13, 42, 155, 83, 85, 86, 87, 7,
	158, 159, 156, 161, 150, 152, 11, 7, 10, 43,
	14, 167, 26, 169, 170, 4, 160, 12, 162, 44,
	24, 16, 174, 25, 15, 0, 177, 0, 0, 180,
	0, 0, 112, 113, 114, 115, 0, 0, 118, 0,
	0, 0, 0, 9, 8, 41, 0, 0, 129, 29,
	165, 30, 32, 33, 34, 0, 0, 0, 35, 36,
	37, 38, 0, 31, 0, 39, 0, 40, 0, 46,
	45, 0, 6, 27, 28, 0, 0, 0, 0, 9,
	8, 41, 0, 0, 0, 29, 157, 30, 32, 33,
	34, 0, 164, 0, 35, 36, 37, 38, 0, 31,
	0, 39, 0, 40, 171, 46, 45, 173, 6, 27,
	28, 0, 0, 0, 179, 0, 0, 182, 9, 8,
	41, 0, 0, 0, 29, 138, 30, 32, 33, 34,
	0, 0, 0, 35, 36, 37, 38, 0, 31, 0,
	39, 0, 40, 0, 46, 45, 0, 6, 27, 28,
	9, 8, 41, 0, 0, 0, 29, 135, 30, 32,
	33, 34, 0, 0, 0, 35, 36, 37, 38, 0,
	31, 0, 39, 0, 40, 0, 46, 45, 0, 6,
	27, 28, 9, 8, 41, 0, 0, 0, 29, 0,
	30, 32, 33, 34, 0, 0, 0, 35, 36, 37,
	38, 0, 31, 0, 39, 0, 40, 0, 46, 45,
	0, 6, 27, 28, 9, 84, 41, 0, 0, 0,
	29, 0, 30, 32, 33, 34, 0, 0, 0, 35,
	36, 37, 38, 0, 31, 0, 39, 0, 40, 0,
	46, 45, 0, 128, 9, 130, 41, 0, 0, 0,
	29, 0, 30, 32, 33, 34, 0, 0, 0, 35,
	36, 37, 38, 0, 31, 0, 39, 0, 40, 0,
	46, 45, 0, 128, 9, 84, 41, 0, 0, 0,
	29, 0, 30, 32, 33, 34, 0, 0, 0, 35,
	36, 37, 38, 0, 31, 0, 39, 0, 40, 0,
	46, 45,
}
var RubyPact = []int{

	21, 328, -1000, 20, -27, -27, -1000, -1000, -27, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -27, -27, -27, -27,
	-27, 98, -27, -1000, -1000, -27, -27, -27, -27, -27,
	-27, -1000, -1000, -1000, -1000, 45, 96, -1000, -1000, 46,
	93, 47, -1000, -1000, 94, -27, -1000, -27, 420, 420,
	420, 420, 128, -20, -1000, -1000, 82, -1000, -1000, -27,
	-27, -27, -1000, -27, -1000, -1000, -27, -27, -27, -27,
	-1, 90, 2, -27, -27, -27, -27, -27, -27, -1000,
	-1000, -27, -1000, -1000, -1000, 420, 420, 420, 420, 128,
	75, 420, 109, -1000, -27, -27, 33, -1000, 46, 42,
	39, 390, -27, -27, -27, -27, -27, -27, -27, -27,
	296, 1, -1000, 31, 264, -1000, -27, -27, -27, -27,
	-13, 74, 118, -7, -1000, -1000, -1000, 85, -1000, 16,
	14, -1000, 72, -27, -1000, -1000, -1000, -1000, 225, -27,
	-27, -20, -27, -20, 64, 420, 189, -1000, -1000, -9,
	-27, -17, -27, -27, -27, -1000, -1000, 360, -1000, 13,
	420, -27, 30, -27, 41, -27, 58, 420, -27, 10,
	420, -1000, 8, -1000,
}
var RubyPgo = []int{

	0, 0, 175, 14, 86, 174, 173, 171, 170, 169,
	167, 165, 162, 107, 160, 159, 158, 156, 143, 142,
	3, 140, 139, 138, 137, 136, 135, 134, 4, 1,
	126, 124, 119, 10, 115,
}
var RubyR1 = []int{

	0, 30, 30, 30, 28, 28, 10, 13, 13, 13,
	13, 11, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	12, 12, 12, 12, 12, 29, 29, 29, 29, 29,
	33, 33, 33, 33, 33, 16, 17, 17, 19, 20,
	20, 34, 34, 14, 14, 21, 22, 23, 24, 25,
	26, 27, 1, 1, 1, 5, 7, 8, 6, 6,
	3, 3, 3, 31, 31, 31, 31, 32, 32, 32,
	32, 2, 2, 9, 9, 18, 18, 15, 15,
}
var RubyR2 = []int{

	0, 0, 1, 3, 0, 2, 2, 2, 2, 1,
	1, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 0, 1, 1, 5, 1,
	0, 1, 1, 5, 5, 9, 6, 8, 6, 3,
	6, 1, 4, 5, 5, 3, 3, 3, 3, 5,
	5, 5, 0, 2, 2, 1, 1, 5, 9, 9,
	0, 1, 4, 0, 6, 11, 12, 0, 4, 9,
	10, 0, 1, 2, 2, 2, 2, 3, 3,
}
var RubyChk = []int{

	-1000, -30, 32, -13, -11, -4, 33, -1, 5, 4,
	-16, -17, -10, -19, -14, -5, -7, -21, -22, -23,
	-24, -25, -26, -27, -8, -6, -12, 34, 35, 10,
	12, 24, 13, 14, 15, 19, 20, 21, 22, 26,
	28, 6, -18, -15, -9, 31, 30, 32, -1, -1,
	-1, -1, -1, -1, -1, -1, 5, -1, -1, -1,
	-1, -1, -1, -1, 5, 6, 31, 5, 6, 21,
	22, 23, -29, 18, 4, 5, 7, -33, 18, 5,
	-20, -1, -20, -4, 5, -4, -4, -4, -33, 5,
	4, -3, 33, 5, 6, -1, -1, -1, -1, -1,
	-1, -1, -1, 33, 16, 6, -34, 33, -1, -1,
	-1, -1, -4, -4, -4, -4, -33, 9, -4, -29,
	-28, -20, -1, 24, -28, 27, -31, -32, 33, -4,
	5, -1, -1, -1, -13, 11, 33, 24, 11, -1,
	-1, -1, -1, 24, 8, 4, 5, 33, -28, 6,
	-3, 9, -3, 9, 18, -1, -28, 11, -1, -1,
	-3, -1, -3, 17, -4, 11, 29, -1, 29, -1,
	-1, -4, 5, -4, -1, 24, 18, -1, 17, -4,
	-1, 9, -4, 9,
}
var RubyDef = []int{

	1, -2, 2, 0, 62, 62, 9, 10, -2, 12,
	13, 14, 15, 16, 17, 18, 19, 20, 21, 22,
	23, 24, 25, 26, 27, 28, 29, 62, 62, 62,
	62, 0, 62, 65, 66, 62, 62, 62, 62, 62,
	62, 31, 32, 33, 34, 0, 0, 3, 7, 8,
	-2, 0, 63, 64, 0, 62, 6, 62, 0, 0,
	0, 0, 40, 70, 85, 86, 0, 83, 84, 62,
	62, 62, 11, 62, -2, -2, 62, -2, 62, 62,
	0, 0, 0, 55, 30, 56, 57, 58, 62, 41,
	42, 62, 71, 87, 88, 0, 0, 0, 0, 40,
	0, 0, -2, 4, 62, -2, 0, 4, 0, 0,
	0, 73, 59, 60, 61, 53, 62, 62, 54, 62,
	0, 0, 49, 0, 0, 67, 62, 62, 62, 62,
	30, 0, 0, 0, 5, 46, 4, 0, 48, 70,
	70, 72, 0, 62, 38, 43, 44, 4, 0, -2,
	62, 70, 62, 70, 0, 0, 0, 47, 50, 0,
	62, 0, 62, 62, -2, 45, 68, 0, 69, 0,
	0, 62, 0, -2, 0, 62, 0, 0, 62, -2,
	0, 80, -2, 76,
}
var RubyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 35,
	33, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 34,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32,
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
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 4:
		//line parser.y:125
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 5:
		//line parser.y:127
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
			}
		}
	case 6:
		//line parser.y:134
		{
			RubyVAL.genericValue = ast.Symbol{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 7:
		//line parser.y:139
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 8:
		//line parser.y:141
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 9:
		//line parser.y:143
		{
			RubyVAL.genericValue = nil /* ignores new lines around statements */
		}
	case 10:
		//line parser.y:145
		{
			RubyVAL.genericValue = nil /* ignores whitespace around statements */
		}
	case 11:
		//line parser.y:148
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
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
		//line parser.y:159
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 36:
		//line parser.y:161
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 37:
		//line parser.y:163
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 38:
		//line parser.y:165
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 39:
		//line parser.y:169
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 40:
		//line parser.y:173
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 41:
		//line parser.y:175
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 42:
		//line parser.y:177
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 43:
		//line parser.y:179
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 44:
		//line parser.y:181
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 45:
		//line parser.y:186
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 46:
		//line parser.y:195
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 47:
		//line parser.y:202
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 48:
		//line parser.y:212
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 49:
		//line parser.y:221
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 50:
		//line parser.y:227
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 51:
		//line parser.y:235
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 52:
		//line parser.y:239
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 53:
		//line parser.y:244
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 54:
		//line parser.y:251
		{ // FIXME: this is a hack. If this rule is not present, "foo = 5" will not parse
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 55:
		//line parser.y:258
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 56:
		//line parser.y:259
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 57:
		//line parser.y:260
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 58:
		//line parser.y:261
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 59:
		//line parser.y:264
		{
			RubyVAL.genericValue = ast.Addition{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 60:
		//line parser.y:272
		{
			RubyVAL.genericValue = ast.Subtraction{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 61:
		//line parser.y:280
		{
			RubyVAL.genericValue = ast.Multiplication{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 62:
		//line parser.y:288
		{
			RubyVAL.genericValue = nil
		}
	case 63:
		//line parser.y:290
		{
			RubyVAL.genericValue = nil
		}
	case 64:
		//line parser.y:292
		{
			RubyVAL.genericValue = nil
		}
	case 65:
		//line parser.y:294
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 66:
		//line parser.y:295
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 67:
		//line parser.y:297
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 68:
		//line parser.y:301
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 69:
		//line parser.y:309
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 70:
		//line parser.y:318
		{
			RubyVAL.genericValue = nil
		}
	case 71:
		//line parser.y:320
		{
			RubyVAL.genericValue = nil
		}
	case 72:
		//line parser.y:322
		{
			RubyVAL.genericValue = nil
		}
	case 73:
		//line parser.y:324
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 74:
		//line parser.y:326
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 75:
		//line parser.y:328
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 76:
		//line parser.y:330
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-7].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 77:
		//line parser.y:332
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 78:
		//line parser.y:334
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 79:
		//line parser.y:341
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 80:
		//line parser.y:348
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 81:
		//line parser.y:356
		{
			RubyVAL.genericValue = nil
		}
	case 82:
		//line parser.y:357
		{
			RubyVAL.genericValue = nil
		}
	case 83:
		//line parser.y:360
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 84:
		//line parser.y:362
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 85:
		//line parser.y:365
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 86:
		//line parser.y:367
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 87:
		//line parser.y:370
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 88:
		//line parser.y:372
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	}
	goto Rubystack /* stack new state and value */
}
