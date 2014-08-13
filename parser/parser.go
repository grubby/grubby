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

//line parser.y:369

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 10,
	-1, 10,
	18, 10,
	-2, 15,
	-1, 27,
	4, 10,
	5, 10,
	7, 10,
	18, 10,
	-2, 32,
	-1, 70,
	9, 10,
	-2, 39,
	-1, 77,
	21, 10,
	22, 10,
	23, 10,
	-2, 61,
	-1, 78,
	21, 10,
	22, 10,
	23, 10,
	-2, 62,
	-1, 79,
	21, 10,
	22, 10,
	23, 10,
	-2, 63,
	-1, 80,
	21, 10,
	22, 10,
	23, 10,
	-2, 64,
	-1, 97,
	27, 57,
	-2, 10,
	-1, 104,
	21, 10,
	22, 10,
	23, 10,
	-2, 65,
	-1, 105,
	21, 10,
	22, 10,
	23, 10,
	-2, 66,
	-1, 106,
	21, 10,
	22, 10,
	23, 10,
	-2, 67,
	-1, 107,
	21, 10,
	22, 10,
	23, 10,
	-2, 60,
	-1, 108,
	21, 10,
	22, 10,
	23, 10,
	-2, 59,
	-1, 123,
	21, 32,
	22, 32,
	23, 32,
	-2, 10,
	-1, 142,
	23, 10,
	-2, 9,
	-1, 144,
	27, 58,
	-2, 10,
	-1, 161,
	21, 10,
	22, 10,
	23, 10,
	-2, 77,
	-1, 170,
	21, 10,
	22, 10,
	23, 10,
	-2, 74,
	-1, 176,
	25, 78,
	32, 78,
	-2, 10,
	-1, 179,
	25, 75,
	32, 75,
	-2, 10,
}

const RubyNprod = 88
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 244

var RubyAct = []int{

	100, 67, 5, 8, 84, 113, 81, 45, 7, 46,
	2, 47, 165, 163, 172, 118, 74, 58, 59, 6,
	131, 116, 137, 6, 102, 6, 4, 141, 48, 130,
	99, 162, 6, 49, 50, 51, 119, 66, 52, 53,
	54, 55, 56, 57, 6, 127, 60, 154, 175, 132,
	173, 75, 75, 63, 64, 65, 77, 78, 79, 80,
	6, 127, 6, 127, 87, 88, 89, 90, 76, 91,
	92, 93, 151, 129, 94, 63, 64, 65, 160, 96,
	138, 118, 101, 150, 148, 103, 6, 127, 95, 180,
	110, 104, 105, 106, 107, 108, 112, 75, 115, 109,
	178, 85, 86, 144, 6, 117, 97, 122, 61, 62,
	124, 125, 126, 114, 128, 6, 145, 146, 128, 133,
	169, 134, 135, 136, 48, 63, 64, 65, 72, 71,
	73, 69, 142, 98, 72, 71, 143, 69, 152, 147,
	149, 70, 68, 121, 128, 155, 120, 153, 156, 1,
	158, 139, 140, 157, 128, 159, 161, 111, 164, 24,
	166, 167, 83, 82, 23, 22, 21, 20, 168, 171,
	19, 170, 18, 174, 14, 29, 177, 13, 176, 12,
	30, 179, 9, 27, 28, 15, 10, 11, 32, 31,
	33, 34, 35, 36, 25, 17, 26, 37, 38, 39,
	40, 16, 3, 0, 0, 0, 0, 41, 0, 42,
	0, 44, 43, 9, 123, 28, 0, 0, 0, 32,
	0, 33, 34, 35, 36, 0, 0, 0, 37, 38,
	39, 40, 0, 0, 0, 0, 0, 0, 41, 0,
	42, 0, 44, 43,
}
var RubyPact = []int{

	-25, 1, -1000, -27, -1000, 178, 8, -1000, 8, -1000,
	8, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 8, -1000, -1000,
	-1000, -1000, 8, 8, 8, -1000, -1000, 8, 8, 8,
	8, 8, 8, 12, 103, -1000, 104, 19, 124, 125,
	8, 8, 178, 178, 178, 178, 158, -1000, -1000, -1000,
	96, -1000, -1000, 8, 8, 8, 8, -1000, 8, 8,
	8, -1000, -1000, 8, 63, 100, 5, 8, 8, 8,
	8, 8, -1000, -1000, -1, -1000, -1000, 178, 178, 178,
	178, 178, 158, 81, 130, -1000, 8, 8, -6, -1000,
	104, 6, -1000, 209, 8, 8, 8, 8, 8, 8,
	8, 8, -1000, 62, 4, -1000, -7, 38, 8, -1000,
	8, 8, 8, -5, 72, 147, 2, -1000, 178, -1000,
	-1000, 97, -1000, 112, 75, 74, 54, 8, -1000, -1000,
	-1000, -1000, 8, 36, 8, -1000, -1000, -1, -1000, -1,
	-1000, 61, 178, 20, -1000, -1000, -19, -1, -20, -1,
	8, 8, -1000, -1000, 178, -1000, 115, 178, 8, -13,
	8, 32, 8, 31, 178, 8, 91, 178, -1000, 80,
	-1000,
}
var RubyPgo = []int{

	0, 203, 3, 202, 201, 196, 195, 194, 189, 187,
	186, 185, 180, 179, 177, 175, 174, 16, 172, 170,
	167, 166, 165, 164, 159, 5, 1, 157, 149, 146,
	143, 6, 141, 133, 0, 4,
}
var RubyR1 = []int{

	0, 28, 28, 28, 28, 3, 3, 25, 25, 25,
	34, 34, 35, 35, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 10, 10, 10, 10, 10, 9, 26, 26,
	32, 32, 32, 32, 31, 31, 31, 31, 31, 13,
	27, 27, 14, 14, 16, 17, 17, 33, 33, 11,
	11, 18, 19, 20, 21, 22, 23, 24, 4, 6,
	7, 5, 5, 29, 29, 29, 29, 30, 30, 30,
	1, 1, 8, 8, 15, 15, 12, 12,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 3, 0, 2, 3,
	0, 2, 0, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 5, 1,
	1, 1, 5, 5, 0, 1, 1, 5, 5, 9,
	0, 1, 6, 8, 6, 3, 6, 1, 4, 5,
	5, 3, 3, 3, 3, 5, 5, 5, 1, 1,
	5, 9, 9, 0, 6, 11, 12, 4, 9, 10,
	0, 1, 2, 2, 2, 2, 3, 3,
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
	-29, -30, -2, 5, -34, -34, -34, 25, -34, 11,
	25, 27, 11, -34, -34, -34, -34, 27, 8, 4,
	5, 25, -2, -25, 6, 4, 5, -35, 9, -35,
	9, 18, -34, -25, 11, -34, -34, -35, -34, -35,
	17, -2, 11, 32, -34, 32, -34, -34, -2, 5,
	-2, -34, 27, 18, -34, 17, -2, -34, 9, -2,
	9,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 0, 10, 4, 10, 14,
	-2, 16, 17, 18, 19, 20, 21, 22, 23, 24,
	25, 26, 27, 28, 29, 30, 31, -2, 33, 34,
	35, 36, 10, 10, 10, 68, 69, 10, 10, 10,
	10, 10, 10, 0, 0, 11, 6, 0, 0, 0,
	10, 10, 0, 0, 0, 0, 44, 12, 84, 85,
	0, 82, 83, 10, 10, 10, 10, 37, 10, 10,
	-2, 40, 41, 10, 0, 0, 0, -2, -2, -2,
	-2, 10, 45, 46, 10, 86, 87, 0, 0, 0,
	0, 0, 44, 0, 50, 7, 10, -2, 0, 7,
	0, 0, 13, 73, -2, -2, -2, -2, -2, 10,
	10, 10, 51, 10, 0, 55, 0, 10, 10, 70,
	10, 10, 10, -2, 0, 0, 0, 8, 0, 52,
	7, 0, 54, 0, 12, 12, 0, 10, 38, 42,
	43, 7, -2, 10, -2, 47, 48, 10, 12, 10,
	12, 0, 0, 10, 53, 56, 0, 10, 0, 10,
	10, -2, 49, 71, 0, 72, 0, 0, 10, 0,
	-2, 0, 10, 0, 0, 10, -2, 0, 79, -2,
	76,
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
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
			}
		}
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 36:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 37:
		//line parser.y:154
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 38:
		//line parser.y:162
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 39:
		//line parser.y:166
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 40:
		//line parser.y:171
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
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
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 45:
		//line parser.y:181
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
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
		//line parser.y:193
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
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 51:
		//line parser.y:202
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 52:
		//line parser.y:205
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 53:
		//line parser.y:212
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 54:
		//line parser.y:222
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 55:
		//line parser.y:231
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 56:
		//line parser.y:237
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 57:
		//line parser.y:245
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 58:
		//line parser.y:249
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 59:
		//line parser.y:254
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 60:
		//line parser.y:261
		{ // FIXME: this is a hack. If this rule is not present, "foo = 5" will not parse
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 61:
		//line parser.y:268
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 62:
		//line parser.y:269
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 63:
		//line parser.y:270
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 64:
		//line parser.y:271
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 65:
		//line parser.y:274
		{
			RubyVAL.genericValue = ast.Addition{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 66:
		//line parser.y:282
		{
			RubyVAL.genericValue = ast.Subtraction{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 67:
		//line parser.y:290
		{
			RubyVAL.genericValue = ast.Multiplication{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 68:
		//line parser.y:297
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 69:
		//line parser.y:298
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 70:
		//line parser.y:300
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 71:
		//line parser.y:304
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 72:
		//line parser.y:312
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 73:
		//line parser.y:320
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 74:
		//line parser.y:322
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 75:
		//line parser.y:324
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 76:
		//line parser.y:326
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 77:
		//line parser.y:329
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 78:
		//line parser.y:336
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 79:
		//line parser.y:343
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 80:
		//line parser.y:351
		{
			RubyVAL.genericValue = nil
		}
	case 81:
		//line parser.y:352
		{
			RubyVAL.genericValue = nil
		}
	case 82:
		//line parser.y:355
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 83:
		//line parser.y:357
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 84:
		//line parser.y:360
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 85:
		//line parser.y:362
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 86:
		//line parser.y:365
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 87:
		//line parser.y:367
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	}
	goto Rubystack /* stack new state and value */
}
