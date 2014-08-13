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
	-1, 8,
	23, 10,
	-2, 6,
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
	-1, 103,
	21, 10,
	22, 10,
	23, 10,
	-2, 65,
	-1, 104,
	21, 10,
	22, 10,
	23, 10,
	-2, 66,
	-1, 105,
	21, 10,
	22, 10,
	23, 10,
	-2, 67,
	-1, 106,
	21, 10,
	22, 10,
	23, 10,
	-2, 60,
	-1, 107,
	21, 10,
	22, 10,
	23, 10,
	-2, 59,
	-1, 122,
	21, 32,
	22, 32,
	23, 32,
	-2, 10,
	-1, 141,
	23, 10,
	-2, 9,
	-1, 143,
	27, 58,
	-2, 10,
	-1, 160,
	21, 10,
	22, 10,
	23, 10,
	-2, 77,
	-1, 169,
	21, 10,
	22, 10,
	23, 10,
	-2, 74,
	-1, 175,
	25, 78,
	32, 78,
	-2, 10,
	-1, 178,
	25, 75,
	32, 75,
	-2, 10,
}

const RubyNprod = 88
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 243

var RubyAct = []int{

	46, 67, 5, 8, 84, 112, 81, 45, 58, 59,
	74, 47, 7, 2, 164, 162, 6, 171, 161, 136,
	130, 115, 172, 117, 140, 63, 64, 65, 48, 6,
	101, 6, 126, 49, 50, 51, 153, 60, 52, 53,
	54, 55, 56, 57, 118, 96, 129, 131, 99, 6,
	126, 75, 75, 179, 95, 6, 77, 78, 79, 80,
	6, 126, 76, 128, 87, 88, 89, 90, 6, 91,
	92, 93, 6, 4, 94, 150, 6, 126, 63, 64,
	65, 66, 100, 149, 174, 102, 177, 63, 64, 65,
	159, 103, 104, 105, 106, 107, 111, 75, 114, 108,
	147, 6, 137, 117, 109, 116, 121, 113, 143, 123,
	124, 125, 97, 127, 85, 86, 168, 127, 132, 73,
	133, 134, 135, 48, 61, 62, 72, 71, 98, 69,
	70, 141, 120, 72, 71, 142, 69, 151, 146, 148,
	68, 144, 145, 127, 154, 119, 152, 155, 1, 157,
	138, 139, 156, 127, 158, 160, 110, 163, 24, 165,
	166, 83, 82, 23, 22, 21, 20, 167, 170, 19,
	169, 18, 173, 14, 29, 176, 13, 175, 12, 30,
	178, 9, 27, 28, 15, 10, 11, 32, 31, 33,
	34, 35, 36, 25, 17, 26, 37, 38, 39, 40,
	16, 3, 0, 0, 0, 0, 41, 0, 42, 0,
	44, 43, 9, 122, 28, 0, 0, 0, 32, 0,
	33, 34, 35, 36, 0, 0, 0, 37, 38, 39,
	40, 0, 0, 0, 0, 0, 0, 41, 0, 42,
	0, 44, 43,
}
var RubyPact = []int{

	-22, 48, -1000, -23, -1000, 177, 31, -1000, 31, -1000,
	31, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 31, -1000, -1000,
	-1000, -1000, 31, 31, 31, -1000, -1000, 31, 31, 31,
	31, 31, 31, 3, 119, -1000, 66, 63, 122, 114,
	31, 31, 177, 177, 177, 177, 157, -1000, -1000, -1000,
	109, -1000, -1000, 31, 31, 31, 31, -1000, 31, 31,
	31, -1000, -1000, 31, 29, 106, 23, 31, 31, 31,
	31, 31, -1000, -1000, 5, -1000, -1000, 177, 177, 177,
	177, 177, 157, 95, 129, -1000, 31, 31, -6, -1000,
	14, -1000, 208, 31, 31, 31, 31, 31, 31, 31,
	31, -1000, 52, 21, -1000, -7, 36, 31, -1000, 31,
	31, 31, -8, 94, 146, -1, -1000, 177, -1000, -1000,
	102, -1000, 137, 91, 74, 57, 31, -1000, -1000, -1000,
	-1000, 31, 25, 31, -1000, -1000, 5, -1000, 5, -1000,
	73, 177, 7, -1000, -1000, -17, 5, -18, 5, 31,
	31, -1000, -1000, 177, -1000, 111, 177, 31, -10, 31,
	4, 31, 67, 177, 31, 77, 177, -1000, 44, -1000,
}
var RubyPgo = []int{

	0, 202, 3, 201, 200, 195, 194, 193, 188, 186,
	185, 184, 179, 178, 176, 174, 173, 10, 171, 169,
	166, 165, 164, 163, 158, 5, 1, 156, 148, 145,
	132, 6, 130, 128, 0, 4,
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

	0, 0, 1, 2, 3, 1, 2, 0, 2, 3,
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
	-34, 25, -34, -2, -2, -2, -2, -2, -31, 9,
	-27, -26, -25, -17, -34, 27, -25, 9, 30, -29,
	-30, -2, 5, -34, -34, -34, 25, -34, 11, 25,
	27, 11, -34, -34, -34, -34, 27, 8, 4, 5,
	25, -2, -25, 6, 4, 5, -35, 9, -35, 9,
	18, -34, -25, 11, -34, -34, -35, -34, -35, 17,
	-2, 11, 32, -34, 32, -34, -34, -2, 5, -2,
	-34, 27, 18, -34, 17, -2, -34, 9, -2, 9,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 0, 10, 4, -2, 14,
	-2, 16, 17, 18, 19, 20, 21, 22, 23, 24,
	25, 26, 27, 28, 29, 30, 31, -2, 33, 34,
	35, 36, 10, 10, 10, 68, 69, 10, 10, 10,
	10, 10, 10, 0, 0, 11, 0, 0, 0, 0,
	10, 10, 0, 0, 0, 0, 44, 12, 84, 85,
	0, 82, 83, 10, 10, 10, 10, 37, 10, 10,
	-2, 40, 41, 10, 0, 0, 0, -2, -2, -2,
	-2, 10, 45, 46, 10, 86, 87, 0, 0, 0,
	0, 0, 44, 0, 50, 7, 10, -2, 0, 7,
	0, 13, 73, -2, -2, -2, -2, -2, 10, 10,
	10, 51, 10, 0, 55, 0, 10, 10, 70, 10,
	10, 10, -2, 0, 0, 0, 8, 0, 52, 7,
	0, 54, 0, 12, 12, 0, 10, 38, 42, 43,
	7, -2, 10, -2, 47, 48, 10, 12, 10, 12,
	0, 0, 10, 53, 56, 0, 10, 0, 10, 10,
	-2, 49, 71, 0, 72, 0, 0, 10, 0, -2,
	0, 10, 0, 0, 10, -2, 0, 79, -2, 76,
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
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
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
