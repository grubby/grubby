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

//line parser.y:375

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 11,
	-1, 10,
	18, 11,
	-2, 16,
	-1, 27,
	4, 11,
	5, 11,
	7, 11,
	18, 11,
	-2, 33,
	-1, 70,
	9, 11,
	-2, 40,
	-1, 77,
	21, 11,
	22, 11,
	23, 11,
	-2, 65,
	-1, 78,
	21, 11,
	22, 11,
	23, 11,
	-2, 66,
	-1, 79,
	21, 11,
	22, 11,
	23, 11,
	-2, 67,
	-1, 80,
	21, 11,
	22, 11,
	23, 11,
	-2, 68,
	-1, 97,
	27, 61,
	-2, 11,
	-1, 104,
	21, 11,
	22, 11,
	23, 11,
	-2, 69,
	-1, 105,
	21, 11,
	22, 11,
	23, 11,
	-2, 70,
	-1, 106,
	21, 11,
	22, 11,
	23, 11,
	-2, 71,
	-1, 107,
	21, 11,
	22, 11,
	23, 11,
	-2, 64,
	-1, 108,
	21, 11,
	22, 11,
	23, 11,
	-2, 63,
	-1, 123,
	21, 33,
	22, 33,
	23, 33,
	-2, 11,
	-1, 129,
	23, 11,
	-2, 10,
	-1, 146,
	27, 62,
	-2, 11,
	-1, 163,
	21, 11,
	22, 11,
	23, 11,
	-2, 81,
	-1, 172,
	21, 11,
	22, 11,
	23, 11,
	-2, 78,
	-1, 178,
	25, 82,
	32, 82,
	-2, 11,
	-1, 181,
	25, 79,
	32, 79,
	-2, 11,
}

const RubyNprod = 92
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 391

var RubyAct = []int{

	100, 126, 5, 113, 84, 74, 67, 45, 7, 46,
	2, 47, 81, 167, 165, 58, 59, 6, 174, 132,
	138, 116, 131, 164, 129, 118, 6, 102, 48, 99,
	8, 142, 143, 49, 50, 51, 142, 143, 52, 53,
	54, 55, 56, 57, 60, 175, 119, 96, 63, 64,
	65, 75, 75, 6, 4, 153, 95, 76, 63, 64,
	65, 182, 6, 66, 87, 88, 89, 90, 180, 91,
	92, 93, 177, 162, 94, 152, 6, 77, 78, 79,
	80, 150, 101, 6, 110, 103, 63, 64, 65, 146,
	72, 71, 98, 69, 97, 139, 118, 75, 115, 85,
	86, 112, 114, 117, 68, 109, 72, 71, 171, 69,
	124, 125, 104, 105, 106, 107, 108, 61, 62, 134,
	73, 135, 136, 137, 48, 147, 148, 70, 122, 121,
	144, 140, 141, 83, 82, 145, 120, 1, 111, 154,
	149, 151, 24, 23, 22, 21, 155, 157, 20, 19,
	158, 18, 160, 14, 29, 159, 13, 161, 12, 30,
	166, 15, 168, 169, 10, 11, 31, 25, 17, 26,
	16, 173, 3, 0, 0, 176, 0, 0, 179, 163,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 170, 0, 0, 172, 0, 0, 0, 0, 0,
	0, 178, 0, 0, 181, 9, 27, 28, 0, 0,
	0, 32, 156, 33, 34, 35, 36, 0, 0, 0,
	37, 38, 39, 40, 0, 128, 127, 0, 0, 0,
	41, 0, 42, 0, 44, 43, 9, 27, 28, 0,
	0, 0, 32, 133, 33, 34, 35, 36, 0, 0,
	0, 37, 38, 39, 40, 0, 128, 127, 0, 0,
	0, 41, 0, 42, 0, 44, 43, 9, 27, 28,
	0, 0, 0, 32, 130, 33, 34, 35, 36, 0,
	0, 0, 37, 38, 39, 40, 0, 128, 127, 0,
	0, 0, 41, 0, 42, 0, 44, 43, 9, 27,
	28, 0, 0, 0, 32, 0, 33, 34, 35, 36,
	0, 0, 0, 37, 38, 39, 40, 0, 128, 127,
	0, 0, 0, 41, 0, 42, 0, 44, 43, 9,
	27, 28, 0, 0, 0, 32, 0, 33, 34, 35,
	36, 0, 0, 0, 37, 38, 39, 40, 0, 0,
	0, 0, 0, 0, 41, 0, 42, 0, 44, 43,
	9, 123, 28, 0, 0, 0, 32, 0, 33, 34,
	35, 36, 0, 0, 0, 37, 38, 39, 40, 0,
	0, 0, 0, 0, 0, 41, 0, 42, 0, 44,
	43,
}
var RubyPact = []int{

	-25, 29, -1000, -27, -1000, 325, 38, -1000, 38, -1000,
	38, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 38, -1000, -1000,
	-1000, -1000, 38, 38, 38, -1000, -1000, 38, 38, 38,
	38, 38, 38, 10, 112, -1000, 65, 45, 86, 115,
	38, 38, 325, 325, 325, 325, 129, -1000, -1000, -1000,
	94, -1000, -1000, 38, 38, 38, 38, -1000, 38, 38,
	38, -1000, -1000, 38, 31, 88, 4, 38, 38, 38,
	38, 38, -1000, -1000, 2, -1000, -1000, 325, 325, 325,
	325, 325, 129, 75, 102, -1000, 38, 38, -6, -1000,
	65, 16, -1000, 356, 38, 38, 38, 38, 38, 38,
	38, -1000, -1000, 263, -3, -1000, -8, 232, 38, -1000,
	38, 38, 38, -7, 87, 127, 7, -1000, -1000, 38,
	-1000, -1000, 83, -1000, 121, 72, 66, 37, 38, -1000,
	-1000, -1000, -1000, -1000, 294, 201, 38, -1000, -1000, 2,
	-1000, 2, -1000, 56, 325, 12, -1000, -1000, -18, 2,
	-19, 2, 38, 38, -1000, -1000, 325, -1000, 103, 325,
	38, -9, 38, 27, 38, 55, 325, 38, 59, 325,
	-1000, 52, -1000,
}
var RubyPgo = []int{

	0, 173, 24, 172, 170, 169, 168, 167, 166, 165,
	164, 161, 159, 158, 156, 154, 153, 5, 151, 149,
	148, 145, 144, 143, 142, 3, 6, 138, 137, 136,
	129, 12, 127, 92, 0, 4, 1,
}
var RubyR1 = []int{

	0, 28, 28, 28, 28, 3, 3, 25, 25, 25,
	25, 34, 34, 35, 35, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 10, 10, 10, 10, 10, 9, 26,
	26, 32, 32, 32, 32, 31, 31, 31, 31, 31,
	36, 36, 36, 13, 27, 27, 14, 14, 16, 17,
	17, 33, 33, 11, 11, 18, 19, 20, 21, 22,
	23, 24, 4, 6, 7, 5, 5, 29, 29, 29,
	29, 30, 30, 30, 1, 1, 8, 8, 15, 15,
	12, 12,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 3, 0, 2, 2,
	2, 0, 2, 0, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 3, 5,
	1, 1, 1, 5, 5, 0, 1, 1, 5, 5,
	0, 2, 2, 9, 0, 1, 6, 8, 6, 3,
	6, 1, 4, 5, 5, 3, 3, 3, 3, 5,
	5, 5, 1, 1, 5, 9, 9, 0, 6, 11,
	12, 4, 9, 10, 0, 1, 2, 2, 2, 2,
	3, 3,
}
var RubyChk = []int{

	-1000, -28, 35, -3, 25, -34, 24, 35, -2, 4,
	-10, -9, -13, -14, -16, -11, -4, -6, -18, -19,
	-20, -21, -22, -23, -24, -7, -5, 5, 6, -15,
	-12, -8, 10, 12, 13, 14, 15, 19, 20, 21,
	22, 29, 31, 34, 33, -34, -34, -34, -34, -34,
	-34, -34, -34, -34, -34, -34, -34, -34, 5, 6,
	34, 5, 6, 21, 22, 23, 18, -26, 18, 7,
	-32, 5, 4, 5, -17, -34, -17, -2, -2, -2,
	-2, -31, 5, 4, -35, 5, 6, -34, -34, -34,
	-34, -34, -34, -34, -34, 25, 16, 6, -33, 25,
	-34, -34, 25, -34, -2, -2, -2, -2, -2, -31,
	9, -27, -26, -25, -17, -34, 27, -25, 9, 30,
	-29, -30, -2, 5, -34, -34, -36, 25, 24, -2,
	11, 25, 27, 11, -34, -34, -34, -34, 27, 8,
	4, 5, 24, 25, -25, -25, 6, 4, 5, -35,
	9, -35, 9, 18, -34, -36, 11, -34, -34, -35,
	-34, -35, 17, -2, 11, 32, -34, 32, -34, -34,
	-2, 5, -2, -34, 27, 18, -34, 17, -2, -34,
	9, -2, 9,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 0, 11, 4, 11, 15,
	-2, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, -2, 34, 35,
	36, 37, 11, 11, 11, 72, 73, 11, 11, 11,
	11, 11, 11, 0, 0, 12, 6, 0, 0, 0,
	11, 11, 0, 0, 0, 0, 45, 13, 88, 89,
	0, 86, 87, 11, 11, 11, 11, 38, 11, 11,
	-2, 41, 42, 11, 0, 0, 0, -2, -2, -2,
	-2, 11, 46, 47, 11, 90, 91, 0, 0, 0,
	0, 0, 45, 0, 54, 7, 11, -2, 0, 7,
	0, 0, 14, 77, -2, -2, -2, -2, -2, 11,
	11, 50, 55, 0, 0, 59, 0, 0, 11, 74,
	11, 11, 11, -2, 0, 0, 7, 8, 9, -2,
	56, 7, 0, 58, 0, 13, 13, 0, 11, 39,
	43, 44, 51, 52, 50, 0, -2, 48, 49, 11,
	13, 11, 13, 0, 0, 0, 57, 60, 0, 11,
	0, 11, 11, -2, 53, 75, 0, 76, 0, 0,
	11, 0, -2, 0, 11, 0, 0, 11, -2, 0,
	83, -2, 80,
}
var RubyTok1 = []int{

	1,
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
		//line parser.y:116
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:118
		{
			Statements = []ast.Node{}
		}
	case 3:
		//line parser.y:120
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 4:
		//line parser.y:126
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 5:
		//line parser.y:132
		{
			RubyVAL.genericValue = nil
		}
	case 6:
		//line parser.y:133
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 7:
		//line parser.y:136
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 8:
		//line parser.y:138
		{ /* do nothing */
		}
	case 9:
		//line parser.y:140
		{ /* do nothing */
		}
	case 10:
		//line parser.y:142
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
			}
		}
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 36:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 37:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 38:
		//line parser.y:156
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 39:
		//line parser.y:164
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 40:
		//line parser.y:168
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 41:
		//line parser.y:173
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 42:
		//line parser.y:175
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 43:
		//line parser.y:177
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 44:
		//line parser.y:179
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 45:
		//line parser.y:181
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 46:
		//line parser.y:183
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 47:
		//line parser.y:185
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 48:
		//line parser.y:187
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 49:
		//line parser.y:189
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 53:
		//line parser.y:199
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 54:
		//line parser.y:207
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 55:
		//line parser.y:208
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 56:
		//line parser.y:211
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 57:
		//line parser.y:218
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 58:
		//line parser.y:228
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 59:
		//line parser.y:237
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 60:
		//line parser.y:243
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 61:
		//line parser.y:251
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 62:
		//line parser.y:255
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 63:
		//line parser.y:260
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 64:
		//line parser.y:267
		{ // FIXME: this is a hack. If this rule is not present, "foo = 5" will not parse
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 65:
		//line parser.y:274
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 66:
		//line parser.y:275
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 67:
		//line parser.y:276
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 68:
		//line parser.y:277
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 69:
		//line parser.y:280
		{
			RubyVAL.genericValue = ast.Addition{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 70:
		//line parser.y:288
		{
			RubyVAL.genericValue = ast.Subtraction{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 71:
		//line parser.y:296
		{
			RubyVAL.genericValue = ast.Multiplication{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 72:
		//line parser.y:303
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 73:
		//line parser.y:304
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 74:
		//line parser.y:306
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 75:
		//line parser.y:310
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 76:
		//line parser.y:318
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 77:
		//line parser.y:326
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 78:
		//line parser.y:328
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 79:
		//line parser.y:330
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 80:
		//line parser.y:332
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 81:
		//line parser.y:335
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 82:
		//line parser.y:342
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 83:
		//line parser.y:349
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 84:
		//line parser.y:357
		{
			RubyVAL.genericValue = nil
		}
	case 85:
		//line parser.y:358
		{
			RubyVAL.genericValue = nil
		}
	case 86:
		//line parser.y:361
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 87:
		//line parser.y:363
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 88:
		//line parser.y:366
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 89:
		//line parser.y:368
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 90:
		//line parser.y:371
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 91:
		//line parser.y:373
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	}
	goto Rubystack /* stack new state and value */
}
