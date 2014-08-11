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
}
var RubyStatenames = []string{}

const RubyEofCode = 1
const RubyErrCode = 2
const RubyMaxDepth = 200

//line parser.y:367

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 7,
	7, 61,
	9, 61,
	-2, 29,
	-1, 48,
	9, 39,
	-2, 34,
	-1, 71,
	9, 41,
	-2, 35,
	-1, 72,
	9, 40,
	-2, 36,
	-1, 74,
	9, 61,
	-2, 38,
	-1, 98,
	9, 39,
	-2, 34,
	-1, 101,
	24, 50,
	-2, 61,
	-1, 145,
	24, 51,
	-2, 61,
	-1, 160,
	21, 61,
	22, 61,
	23, 61,
	-2, 77,
	-1, 169,
	21, 61,
	22, 61,
	23, 61,
	-2, 73,
	-1, 175,
	29, 78,
	32, 78,
	-2, 61,
	-1, 178,
	29, 74,
	32, 74,
	-2, 61,
}

const RubyNprod = 86
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 432

var RubyAct = []int{

	104, 69, 6, 77, 46, 47, 88, 168, 48, 74,
	149, 116, 139, 26, 27, 164, 89, 179, 124, 177,
	162, 26, 27, 124, 147, 143, 49, 50, 51, 52,
	53, 132, 55, 89, 124, 56, 57, 58, 59, 60,
	61, 26, 27, 26, 27, 100, 103, 89, 113, 66,
	67, 68, 62, 70, 78, 171, 78, 133, 119, 79,
	172, 99, 75, 66, 67, 68, 121, 91, 92, 93,
	85, 94, 4, 174, 95, 96, 97, 98, 63, 113,
	150, 159, 105, 66, 67, 68, 106, 140, 113, 107,
	64, 65, 71, 72, 90, 73, 71, 72, 76, 73,
	115, 78, 118, 145, 117, 112, 70, 141, 142, 87,
	86, 102, 101, 127, 128, 120, 129, 6, 54, 123,
	130, 6, 2, 135, 136, 137, 138, 105, 122, 80,
	82, 83, 84, 1, 22, 21, 20, 19, 18, 17,
	151, 16, 146, 148, 144, 6, 154, 155, 12, 157,
	41, 10, 9, 6, 156, 152, 158, 163, 42, 165,
	166, 13, 25, 3, 108, 109, 110, 111, 170, 11,
	114, 43, 173, 23, 15, 176, 8, 7, 40, 24,
	125, 14, 28, 161, 29, 31, 32, 33, 0, 0,
	0, 34, 35, 36, 37, 0, 30, 0, 38, 0,
	39, 0, 45, 44, 5, 26, 27, 0, 0, 0,
	0, 8, 81, 40, 0, 0, 0, 28, 0, 29,
	31, 32, 33, 0, 160, 0, 34, 35, 36, 37,
	0, 30, 0, 38, 0, 39, 167, 45, 44, 169,
	0, 0, 0, 0, 0, 0, 175, 0, 0, 178,
	8, 7, 40, 0, 0, 0, 28, 153, 29, 31,
	32, 33, 0, 0, 0, 34, 35, 36, 37, 0,
	30, 0, 38, 0, 39, 0, 45, 44, 5, 26,
	27, 8, 7, 40, 0, 0, 0, 28, 134, 29,
	31, 32, 33, 0, 0, 0, 34, 35, 36, 37,
	0, 30, 0, 38, 0, 39, 0, 45, 44, 5,
	26, 27, 8, 7, 40, 0, 0, 0, 28, 131,
	29, 31, 32, 33, 0, 0, 0, 34, 35, 36,
	37, 0, 30, 0, 38, 0, 39, 0, 45, 44,
	5, 26, 27, 8, 7, 40, 0, 0, 0, 28,
	0, 29, 31, 32, 33, 0, 0, 0, 34, 35,
	36, 37, 0, 30, 0, 38, 0, 39, 0, 45,
	44, 5, 26, 27, 8, 81, 40, 0, 0, 0,
	28, 0, 29, 31, 32, 33, 0, 0, 0, 34,
	35, 36, 37, 0, 30, 0, 38, 0, 39, 0,
	45, 44, 124, 8, 126, 40, 0, 0, 0, 28,
	0, 29, 31, 32, 33, 0, 0, 0, 34, 35,
	36, 37, 0, 30, 0, 38, 0, 39, 0, 45,
	44, 124,
}
var RubyPact = []int{

	-1000, 339, -1000, -20, -20, -1000, -1000, -20, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -20, -20, -20, -20, -20,
	113, -20, -1000, -1000, -20, -20, -20, -20, -20, -20,
	-1000, -1000, -1000, -1000, 47, 85, -1000, 28, 88, 44,
	-1000, -1000, 93, -20, -1000, -20, 207, 207, 207, 207,
	105, -16, -1000, 89, -1000, -1000, -20, -20, -20, -1000,
	-20, -1000, -1000, -20, -20, -20, -20, 29, 106, 14,
	-20, -20, -20, -20, -20, -20, -1000, -1000, -20, -1000,
	-1000, 207, 207, 207, 207, 105, 70, 207, 92, -1000,
	-20, -20, 34, -1000, 28, 35, 39, 399, -20, -20,
	-20, -20, -20, -20, -20, -20, 308, -1, -1000, 33,
	277, -1000, -20, -20, -20, -20, -12, 79, 103, -7,
	-1000, -1000, -1000, 97, -1000, 15, 1, -1000, 62, -20,
	-1000, -1000, -1000, -1000, 246, -20, -20, -16, -20, -16,
	64, 207, 172, -1000, -1000, -9, -20, -14, -20, -20,
	-20, -1000, -1000, 370, -1000, 2, 207, -20, 31, -20,
	42, -20, 56, 207, -20, 10, 207, -1000, 8, -1000,
}
var RubyPgo = []int{

	0, 0, 188, 6, 72, 181, 179, 174, 173, 171,
	169, 163, 162, 120, 161, 158, 152, 151, 150, 148,
	3, 141, 139, 138, 137, 136, 135, 134, 11, 1,
	133, 128, 119, 9, 111,
}
var RubyR1 = []int{

	0, 30, 30, 28, 28, 10, 13, 13, 13, 13,
	11, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 12,
	12, 12, 12, 12, 29, 29, 29, 29, 29, 33,
	33, 33, 33, 33, 16, 17, 17, 19, 20, 20,
	34, 34, 14, 14, 21, 22, 23, 24, 25, 26,
	27, 1, 1, 1, 5, 7, 8, 6, 6, 3,
	3, 3, 31, 31, 31, 31, 32, 32, 32, 32,
	2, 2, 9, 9, 18, 15,
}
var RubyR2 = []int{

	0, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 0, 1, 1, 5, 1, 0,
	1, 1, 5, 5, 9, 6, 8, 6, 3, 6,
	1, 4, 5, 5, 3, 3, 3, 3, 5, 5,
	5, 0, 2, 2, 1, 1, 5, 9, 9, 0,
	1, 4, 0, 6, 11, 12, 0, 4, 9, 10,
	0, 1, 2, 2, 2, 3,
}
var RubyChk = []int{

	-1000, -30, -13, -11, -4, 32, -1, 5, 4, -16,
	-17, -10, -19, -14, -5, -7, -21, -22, -23, -24,
	-25, -26, -27, -8, -6, -12, 33, 34, 10, 12,
	24, 13, 14, 15, 19, 20, 21, 22, 26, 28,
	6, -18, -15, -9, 31, 30, -1, -1, -1, -1,
	-1, -1, -1, -1, 5, -1, -1, -1, -1, -1,
	-1, -1, 5, 31, 5, 6, 21, 22, 23, -29,
	18, 4, 5, 7, -33, 18, 5, -20, -1, -20,
	-4, 5, -4, -4, -4, -33, 5, 4, -3, 32,
	5, -1, -1, -1, -1, -1, -1, -1, -1, 32,
	16, 6, -34, 32, -1, -1, -1, -1, -4, -4,
	-4, -4, -33, 9, -4, -29, -28, -20, -1, 24,
	-28, 27, -31, -32, 32, -4, 5, -1, -1, -1,
	-13, 11, 32, 24, 11, -1, -1, -1, -1, 24,
	8, 4, 5, 32, -28, 6, -3, 9, -3, 9,
	18, -1, -28, 11, -1, -1, -3, -1, -3, 17,
	-4, 11, 29, -1, 29, -1, -1, -4, 5, -4,
	-1, 24, 18, -1, 17, -4, -1, 9, -4, 9,
}
var RubyDef = []int{

	1, -2, 2, 61, 61, 8, 9, -2, 11, 12,
	13, 14, 15, 16, 17, 18, 19, 20, 21, 22,
	23, 24, 25, 26, 27, 28, 61, 61, 61, 61,
	0, 61, 64, 65, 61, 61, 61, 61, 61, 61,
	30, 31, 32, 33, 0, 0, 6, 7, -2, 0,
	62, 63, 0, 61, 5, 61, 0, 0, 0, 0,
	39, 69, 84, 0, 82, 83, 61, 61, 61, 10,
	61, -2, -2, 61, -2, 61, 61, 0, 0, 0,
	54, 29, 55, 56, 57, 61, 40, 41, 61, 70,
	85, 0, 0, 0, 0, 39, 0, 0, -2, 3,
	61, -2, 0, 3, 0, 0, 0, 72, 58, 59,
	60, 52, 61, 61, 53, 61, 0, 0, 48, 0,
	0, 66, 61, 61, 61, 61, 29, 0, 0, 0,
	4, 45, 3, 0, 47, 69, 69, 71, 0, 61,
	37, 42, 43, 3, 0, -2, 61, 69, 61, 69,
	0, 0, 0, 46, 49, 0, 61, 0, 61, 61,
	-2, 44, 67, 0, 68, 0, 0, 61, 0, -2,
	0, 61, 0, 0, 61, -2, 0, 79, -2, 75,
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
		//line parser.y:122
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 4:
		//line parser.y:124
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
			}
		}
	case 5:
		//line parser.y:131
		{
			RubyVAL.genericValue = ast.Symbol{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 6:
		//line parser.y:136
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 7:
		//line parser.y:138
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 8:
		//line parser.y:140
		{
			RubyVAL.genericValue = nil /* ignores new lines around statements */
		}
	case 9:
		//line parser.y:142
		{
			RubyVAL.genericValue = nil /* ignores whitespace around statements */
		}
	case 10:
		//line parser.y:145
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 32:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 33:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 34:
		//line parser.y:156
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 35:
		//line parser.y:158
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 36:
		//line parser.y:160
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 37:
		//line parser.y:162
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 38:
		//line parser.y:166
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 39:
		//line parser.y:170
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 40:
		//line parser.y:172
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 41:
		//line parser.y:174
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 42:
		//line parser.y:176
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 43:
		//line parser.y:178
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 44:
		//line parser.y:183
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 45:
		//line parser.y:192
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 46:
		//line parser.y:199
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 47:
		//line parser.y:209
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 48:
		//line parser.y:218
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 49:
		//line parser.y:224
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 50:
		//line parser.y:232
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 51:
		//line parser.y:236
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 52:
		//line parser.y:241
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 53:
		//line parser.y:248
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 54:
		//line parser.y:255
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 55:
		//line parser.y:256
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 56:
		//line parser.y:257
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 57:
		//line parser.y:258
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 58:
		//line parser.y:261
		{
			RubyVAL.genericValue = ast.Addition{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 59:
		//line parser.y:269
		{
			RubyVAL.genericValue = ast.Subtraction{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 60:
		//line parser.y:277
		{
			RubyVAL.genericValue = ast.Multiplication{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 61:
		//line parser.y:285
		{
			RubyVAL.genericValue = nil
		}
	case 62:
		//line parser.y:287
		{
			RubyVAL.genericValue = nil
		}
	case 63:
		//line parser.y:289
		{
			RubyVAL.genericValue = nil
		}
	case 64:
		//line parser.y:291
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 65:
		//line parser.y:292
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 66:
		//line parser.y:294
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 67:
		//line parser.y:298
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 68:
		//line parser.y:306
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 69:
		//line parser.y:315
		{
			RubyVAL.genericValue = nil
		}
	case 70:
		//line parser.y:317
		{
			RubyVAL.genericValue = nil
		}
	case 71:
		//line parser.y:319
		{
			RubyVAL.genericValue = nil
		}
	case 72:
		//line parser.y:321
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 73:
		//line parser.y:323
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 74:
		//line parser.y:325
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 75:
		//line parser.y:327
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-7].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 76:
		//line parser.y:329
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 77:
		//line parser.y:331
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 78:
		//line parser.y:338
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 79:
		//line parser.y:345
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 80:
		//line parser.y:353
		{
			RubyVAL.genericValue = nil
		}
	case 81:
		//line parser.y:354
		{
			RubyVAL.genericValue = nil
		}
	case 82:
		//line parser.y:357
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 83:
		//line parser.y:359
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 84:
		//line parser.y:362
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 85:
		//line parser.y:365
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	}
	goto Rubystack /* stack new state and value */
}
