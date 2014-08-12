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

//line parser.y:363

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 3,
	23, 7,
	-2, 3,
	-1, 5,
	18, 7,
	-2, 12,
	-1, 22,
	4, 7,
	5, 7,
	7, 7,
	9, 7,
	18, 7,
	-2, 29,
	-1, 75,
	21, 7,
	22, 7,
	23, 7,
	-2, 58,
	-1, 76,
	21, 7,
	22, 7,
	23, 7,
	-2, 59,
	-1, 77,
	21, 7,
	22, 7,
	23, 7,
	-2, 60,
	-1, 78,
	21, 7,
	22, 7,
	23, 7,
	-2, 61,
	-1, 95,
	27, 54,
	-2, 7,
	-1, 101,
	21, 7,
	22, 7,
	23, 7,
	-2, 62,
	-1, 102,
	21, 7,
	22, 7,
	23, 7,
	-2, 63,
	-1, 103,
	21, 7,
	22, 7,
	23, 7,
	-2, 64,
	-1, 104,
	21, 7,
	22, 7,
	23, 7,
	-2, 57,
	-1, 105,
	21, 7,
	22, 7,
	23, 7,
	-2, 56,
	-1, 119,
	21, 29,
	22, 29,
	23, 29,
	-2, 7,
	-1, 123,
	23, 7,
	-2, 6,
	-1, 138,
	27, 55,
	-2, 7,
	-1, 155,
	21, 7,
	22, 7,
	23, 7,
	-2, 74,
	-1, 164,
	21, 7,
	22, 7,
	23, 7,
	-2, 71,
	-1, 170,
	25, 75,
	32, 75,
	-2, 7,
	-1, 173,
	25, 72,
	32, 72,
	-2, 7,
}

const RubyNprod = 85
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 361

var RubyAct = []int{

	41, 109, 123, 70, 3, 82, 43, 72, 64, 136,
	125, 94, 97, 54, 55, 42, 2, 159, 157, 114,
	166, 42, 63, 44, 132, 126, 40, 112, 45, 46,
	47, 93, 138, 48, 49, 50, 51, 52, 53, 42,
	115, 167, 56, 62, 59, 60, 61, 73, 73, 42,
	99, 75, 76, 77, 78, 74, 79, 59, 60, 61,
	85, 86, 87, 174, 88, 169, 89, 172, 154, 90,
	145, 91, 92, 59, 60, 61, 133, 114, 42, 144,
	98, 142, 42, 100, 66, 67, 107, 68, 101, 102,
	103, 104, 105, 95, 106, 73, 111, 163, 65, 113,
	71, 108, 110, 118, 83, 84, 96, 120, 121, 122,
	66, 67, 69, 68, 117, 128, 116, 129, 130, 131,
	44, 57, 58, 139, 140, 134, 135, 137, 81, 80,
	1, 19, 18, 146, 17, 141, 143, 16, 147, 149,
	15, 14, 150, 13, 152, 9, 24, 8, 151, 155,
	153, 7, 158, 25, 160, 161, 10, 5, 6, 26,
	20, 162, 12, 165, 164, 21, 11, 168, 0, 0,
	171, 170, 0, 0, 173, 4, 22, 23, 0, 0,
	0, 27, 156, 28, 29, 30, 31, 0, 0, 0,
	32, 33, 34, 35, 0, 0, 0, 0, 0, 0,
	36, 0, 37, 0, 39, 38, 4, 22, 23, 0,
	0, 0, 27, 148, 28, 29, 30, 31, 0, 0,
	0, 32, 33, 34, 35, 0, 0, 0, 0, 0,
	0, 36, 0, 37, 0, 39, 38, 4, 22, 23,
	0, 0, 0, 27, 127, 28, 29, 30, 31, 0,
	0, 0, 32, 33, 34, 35, 0, 0, 0, 0,
	0, 0, 36, 0, 37, 0, 39, 38, 4, 22,
	23, 0, 0, 0, 27, 124, 28, 29, 30, 31,
	0, 0, 0, 32, 33, 34, 35, 0, 0, 0,
	0, 0, 0, 36, 0, 37, 0, 39, 38, 4,
	22, 23, 0, 0, 0, 27, 0, 28, 29, 30,
	31, 0, 0, 0, 32, 33, 34, 35, 0, 0,
	0, 0, 0, 0, 36, 0, 37, 0, 39, 38,
	4, 119, 23, 0, 0, 0, 27, 0, 28, 29,
	30, 31, 0, 0, 0, 32, 33, 34, 35, 0,
	0, 0, 0, 0, 0, 36, 0, 37, 0, 39,
	38,
}
var RubyPact = []int{

	-19, 295, -1000, -9, -1000, 15, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 15, -1000, -1000, -1000, -1000, 15, 15, 15,
	-1000, -1000, 15, 15, 15, 15, 15, 15, 8, 116,
	-1000, 36, 15, 4, 80, 95, 15, 15, 295, 295,
	295, 295, 124, -1000, -1000, -1000, 99, -1000, -1000, 15,
	15, 15, -1000, 15, -1000, 15, -1000, -1000, 15, -1000,
	15, 15, -5, 87, -24, 15, 15, 15, 15, 15,
	-1000, -1000, 25, -1000, -1000, 295, 295, 295, 295, 295,
	124, 77, 106, -1000, 15, 15, 0, -1000, 10, -1000,
	326, 15, 15, 15, 15, 15, 15, 15, 15, 264,
	-26, -1000, -2, 233, 15, -1000, 15, 15, 15, -3,
	68, 121, -27, 15, -1000, -1000, 26, -1000, 119, 72,
	70, 52, 15, -1000, -1000, -1000, -1000, 202, 15, -1000,
	-1000, 25, -1000, 25, -1000, 51, 295, 171, -1000, -1000,
	-14, 25, -15, 25, 15, 15, -1000, -1000, 295, -1000,
	92, 295, 15, -7, 15, 23, 15, 48, 295, 15,
	58, 295, -1000, 54, -1000,
}
var RubyPgo = []int{

	0, 168, 2, 166, 165, 162, 160, 159, 158, 157,
	156, 153, 151, 147, 146, 145, 7, 143, 141, 140,
	137, 134, 132, 131, 1, 8, 130, 116, 114, 3,
	112, 106, 0, 5,
}
var RubyR1 = []int{

	0, 26, 26, 26, 26, 24, 24, 32, 32, 33,
	33, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 9,
	9, 9, 9, 9, 8, 25, 25, 25, 25, 30,
	30, 30, 30, 29, 29, 29, 29, 29, 12, 13,
	13, 15, 16, 16, 31, 31, 10, 10, 17, 18,
	19, 20, 21, 22, 23, 3, 5, 6, 4, 4,
	27, 27, 27, 27, 28, 28, 28, 1, 1, 7,
	7, 14, 14, 11, 11,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 0, 2, 0, 2, 0,
	2, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 3, 1, 1, 5, 1, 1,
	1, 5, 5, 0, 1, 1, 5, 5, 9, 6,
	8, 6, 3, 6, 1, 4, 5, 5, 3, 3,
	3, 3, 5, 5, 5, 1, 1, 5, 9, 9,
	0, 6, 11, 12, 4, 9, 10, 0, 1, 2,
	2, 2, 2, 3, 3,
}
var RubyChk = []int{

	-1000, -26, 35, -2, 4, -9, -8, -12, -13, -15,
	-10, -3, -5, -17, -18, -19, -20, -21, -22, -23,
	-6, -4, 5, 6, -14, -11, -7, 10, 12, 13,
	14, 15, 19, 20, 21, 22, 29, 31, 34, 33,
	35, -32, 24, -32, -32, -32, -32, -32, -32, -32,
	-32, -32, -32, -32, 5, 6, 34, 5, 6, 21,
	22, 23, -32, 18, -25, 18, 4, 5, 7, -30,
	-29, 5, -16, -32, -16, -2, -2, -2, -2, -29,
	5, 4, -33, 5, 6, -32, -32, -32, -32, -32,
	-32, -32, -32, 36, 16, 6, -31, 36, -32, 25,
	-32, -2, -2, -2, -2, -2, -29, 9, -25, -24,
	-16, -32, 27, -24, 9, 30, -27, -28, -2, 5,
	-32, -32, -32, -2, 11, 36, 27, 11, -32, -32,
	-32, -32, 27, 8, 4, 5, 36, -24, 6, 4,
	5, -33, 9, -33, 9, 18, -32, -24, 11, -32,
	-32, -33, -32, -33, 17, -2, 11, 32, -32, 32,
	-32, -32, -2, 5, -2, -32, 27, 18, -32, 17,
	-2, -32, 9, -2, 9,
}
var RubyDef = []int{

	1, -2, 2, -2, 11, -2, 13, 14, 15, 16,
	17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
	27, 28, -2, 30, 31, 32, 33, 7, 7, 7,
	65, 66, 7, 7, 7, 7, 7, 7, 0, 0,
	4, 0, 7, 0, 43, 0, 7, 7, 0, 0,
	0, 0, 43, 9, 81, 82, 0, 79, 80, 7,
	7, 7, 8, 7, 34, 7, 35, 36, 7, 38,
	7, 7, 0, 0, 0, -2, -2, -2, -2, 7,
	44, 45, 7, 83, 84, 0, 0, 0, 0, 0,
	43, 0, 43, 5, 7, -2, 0, 5, 0, 10,
	70, -2, -2, -2, -2, -2, 7, 7, 7, 0,
	0, 52, 0, 0, 7, 67, 7, 7, 7, -2,
	0, 0, 0, -2, 49, 5, 0, 51, 0, 9,
	9, 0, 7, 37, 41, 42, 5, 0, -2, 46,
	47, 7, 9, 7, 9, 0, 0, 0, 50, 53,
	0, 7, 0, 7, 7, -2, 48, 68, 0, 69,
	0, 0, 7, 0, -2, 0, 7, 0, 0, 7,
	-2, 0, 76, -2, 73,
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
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 4:
		//line parser.y:124
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 5:
		//line parser.y:131
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 6:
		//line parser.y:133
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
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
		//line parser.y:147
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 35:
		//line parser.y:155
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 36:
		//line parser.y:157
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 37:
		//line parser.y:159
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 38:
		//line parser.y:163
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
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
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 43:
		//line parser.y:176
		{
			RubyVAL.genericSlice = ast.Nodes{}
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
		//line parser.y:184
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 48:
		//line parser.y:190
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 49:
		//line parser.y:199
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 50:
		//line parser.y:206
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 51:
		//line parser.y:216
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 52:
		//line parser.y:225
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 53:
		//line parser.y:231
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 54:
		//line parser.y:239
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 55:
		//line parser.y:243
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 56:
		//line parser.y:248
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 57:
		//line parser.y:255
		{ // FIXME: this is a hack. If this rule is not present, "foo = 5" will not parse
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 58:
		//line parser.y:262
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 59:
		//line parser.y:263
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 60:
		//line parser.y:264
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 61:
		//line parser.y:265
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 62:
		//line parser.y:268
		{
			RubyVAL.genericValue = ast.Addition{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 63:
		//line parser.y:276
		{
			RubyVAL.genericValue = ast.Subtraction{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 64:
		//line parser.y:284
		{
			RubyVAL.genericValue = ast.Multiplication{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 65:
		//line parser.y:291
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 66:
		//line parser.y:292
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 67:
		//line parser.y:294
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 68:
		//line parser.y:298
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 69:
		//line parser.y:306
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 70:
		//line parser.y:314
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 71:
		//line parser.y:316
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 72:
		//line parser.y:318
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 73:
		//line parser.y:320
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 74:
		//line parser.y:323
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 75:
		//line parser.y:330
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 76:
		//line parser.y:337
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 77:
		//line parser.y:345
		{
			RubyVAL.genericValue = nil
		}
	case 78:
		//line parser.y:346
		{
			RubyVAL.genericValue = nil
		}
	case 79:
		//line parser.y:349
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 80:
		//line parser.y:351
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 81:
		//line parser.y:354
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 82:
		//line parser.y:356
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 83:
		//line parser.y:359
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 84:
		//line parser.y:361
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	}
	goto Rubystack /* stack new state and value */
}
