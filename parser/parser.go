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
const FILE_CONST_REF = 57377
const EOF = 57378

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
	"FILE_CONST_REF",
	"EOF",
}
var RubyStatenames = []string{}

const RubyEofCode = 1
const RubyErrCode = 2
const RubyMaxDepth = 200

//line parser.y:388

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 11,
	-1, 10,
	18, 11,
	-2, 16,
	-1, 28,
	4, 11,
	5, 11,
	7, 11,
	18, 11,
	-2, 34,
	-1, 73,
	9, 11,
	-2, 42,
	-1, 81,
	21, 11,
	22, 11,
	23, 11,
	-2, 67,
	-1, 82,
	21, 11,
	22, 11,
	23, 11,
	-2, 68,
	-1, 83,
	21, 11,
	22, 11,
	23, 11,
	-2, 69,
	-1, 84,
	21, 11,
	22, 11,
	23, 11,
	-2, 70,
	-1, 102,
	27, 63,
	-2, 11,
	-1, 109,
	21, 11,
	22, 11,
	23, 11,
	-2, 71,
	-1, 110,
	21, 11,
	22, 11,
	23, 11,
	-2, 72,
	-1, 111,
	21, 11,
	22, 11,
	23, 11,
	-2, 73,
	-1, 112,
	21, 11,
	22, 11,
	23, 11,
	-2, 66,
	-1, 113,
	21, 11,
	22, 11,
	23, 11,
	-2, 65,
	-1, 129,
	21, 34,
	22, 34,
	23, 34,
	-2, 11,
	-1, 135,
	23, 11,
	-2, 10,
	-1, 152,
	27, 64,
	-2, 11,
	-1, 169,
	21, 11,
	22, 11,
	23, 11,
	-2, 83,
	-1, 178,
	21, 11,
	22, 11,
	23, 11,
	-2, 80,
	-1, 184,
	25, 84,
	32, 84,
	-2, 11,
	-1, 187,
	25, 81,
	32, 81,
	-2, 11,
}

const RubyNprod = 95
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 404

var RubyAct = []int{

	105, 132, 5, 119, 88, 78, 70, 47, 7, 48,
	2, 49, 85, 173, 171, 6, 61, 62, 144, 51,
	6, 180, 138, 188, 51, 135, 124, 122, 137, 50,
	104, 8, 6, 107, 52, 53, 54, 170, 6, 55,
	56, 57, 58, 59, 60, 63, 6, 125, 148, 149,
	148, 149, 101, 181, 79, 79, 66, 67, 68, 159,
	80, 100, 66, 67, 68, 186, 69, 91, 92, 93,
	94, 183, 95, 96, 97, 6, 4, 98, 99, 168,
	6, 81, 82, 83, 84, 158, 106, 145, 124, 108,
	66, 67, 68, 156, 115, 152, 75, 74, 177, 72,
	89, 90, 79, 121, 102, 116, 118, 120, 123, 114,
	71, 75, 74, 77, 72, 130, 131, 109, 110, 111,
	112, 113, 64, 65, 76, 140, 103, 141, 142, 143,
	50, 153, 154, 73, 128, 127, 150, 146, 147, 87,
	86, 151, 126, 1, 117, 160, 155, 157, 24, 23,
	22, 21, 161, 163, 20, 19, 164, 18, 166, 27,
	14, 165, 30, 167, 13, 12, 172, 31, 174, 175,
	15, 10, 11, 32, 25, 17, 26, 179, 16, 3,
	0, 182, 0, 0, 185, 0, 169, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 176, 0,
	0, 178, 0, 0, 0, 0, 0, 0, 184, 0,
	0, 187, 9, 28, 29, 0, 0, 0, 33, 162,
	34, 35, 36, 37, 0, 0, 0, 38, 39, 40,
	41, 0, 134, 133, 0, 0, 0, 42, 0, 43,
	0, 46, 45, 44, 9, 28, 29, 0, 0, 0,
	33, 139, 34, 35, 36, 37, 0, 0, 0, 38,
	39, 40, 41, 0, 134, 133, 0, 0, 0, 42,
	0, 43, 0, 46, 45, 44, 9, 28, 29, 0,
	0, 0, 33, 136, 34, 35, 36, 37, 0, 0,
	0, 38, 39, 40, 41, 0, 134, 133, 0, 0,
	0, 42, 0, 43, 0, 46, 45, 44, 9, 28,
	29, 0, 0, 0, 33, 0, 34, 35, 36, 37,
	0, 0, 0, 38, 39, 40, 41, 0, 134, 133,
	0, 0, 0, 42, 0, 43, 0, 46, 45, 44,
	9, 28, 29, 0, 0, 0, 33, 0, 34, 35,
	36, 37, 0, 0, 0, 38, 39, 40, 41, 0,
	0, 0, 0, 0, 0, 42, 0, 43, 0, 46,
	45, 44, 9, 129, 29, 0, 0, 0, 33, 0,
	34, 35, 36, 37, 0, 0, 0, 38, 39, 40,
	41, 0, 0, 0, 0, 0, 0, 42, 0, 43,
	0, 46, 45, 44,
}
var RubyPact = []int{

	-26, 51, -1000, -28, -1000, 336, 22, -1000, 22, -1000,
	22, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -4, -1000,
	-1000, -1000, -1000, 22, 22, 22, -1000, -1000, 22, 22,
	22, 22, 22, 22, -1000, 11, 117, -1000, 69, 48,
	92, 119, 108, 22, 22, 336, 336, 336, 336, 135,
	-1000, -1000, -1000, 95, -1000, -1000, 22, 22, 22, 22,
	-1000, 22, 22, 22, -1000, -1000, 22, 22, 36, 98,
	5, 22, 22, 22, 22, 22, -1000, -1000, 8, -1000,
	-1000, 336, 336, 336, 336, 336, 135, 85, 107, 107,
	-1000, 22, 22, 0, -1000, 69, 17, -1000, 368, 22,
	22, 22, 22, 22, 22, 22, -1000, -1000, -1000, 272,
	3, -1000, -5, 240, 22, -1000, 22, 22, 22, -9,
	79, 133, 24, -1000, -1000, 22, -1000, -1000, 89, -1000,
	127, 84, 76, 41, 22, -1000, -1000, -1000, -1000, -1000,
	304, 208, 22, -1000, -1000, 8, -1000, 8, -1000, 62,
	336, 26, -1000, -1000, -18, 8, -19, 8, 22, 22,
	-1000, -1000, 336, -1000, 93, 336, 22, -6, 22, 35,
	22, 54, 336, 22, 56, 336, -1000, 14, -1000,
}
var RubyPgo = []int{

	0, 180, 25, 179, 178, 176, 175, 174, 173, 172,
	171, 170, 167, 165, 164, 162, 160, 5, 159, 157,
	155, 154, 151, 150, 149, 148, 3, 6, 144, 143,
	142, 135, 12, 133, 126, 0, 4, 1,
}
var RubyR1 = []int{

	0, 29, 29, 29, 29, 3, 3, 26, 26, 26,
	26, 35, 35, 36, 36, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 10, 10, 10, 10, 10, 9,
	9, 27, 27, 33, 33, 33, 33, 32, 32, 32,
	32, 32, 37, 37, 37, 13, 28, 28, 14, 14,
	16, 17, 17, 34, 34, 11, 11, 19, 20, 21,
	22, 23, 24, 25, 4, 6, 7, 5, 5, 30,
	30, 30, 30, 31, 31, 31, 1, 1, 8, 8,
	15, 15, 12, 12, 18,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 3, 0, 2, 2,
	2, 0, 2, 0, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
	5, 5, 1, 1, 1, 5, 5, 0, 1, 1,
	5, 5, 0, 2, 2, 9, 0, 1, 6, 8,
	6, 3, 6, 1, 4, 5, 5, 3, 3, 3,
	3, 5, 5, 5, 1, 1, 5, 9, 9, 0,
	6, 11, 12, 4, 9, 10, 0, 1, 2, 2,
	2, 2, 3, 3, 1,
}
var RubyChk = []int{

	-1000, -29, 36, -3, 25, -35, 24, 36, -2, 4,
	-10, -9, -13, -14, -16, -11, -4, -6, -19, -20,
	-21, -22, -23, -24, -25, -7, -5, -18, 5, 6,
	-15, -12, -8, 10, 12, 13, 14, 15, 19, 20,
	21, 22, 29, 31, 35, 34, 33, -35, -35, -35,
	-35, 28, -35, -35, -35, -35, -35, -35, -35, -35,
	-35, 5, 6, 34, 5, 6, 21, 22, 23, 18,
	-27, 18, 7, -33, 5, 4, 5, 5, -17, -35,
	-17, -2, -2, -2, -2, -32, 5, 4, -36, 5,
	6, -35, -35, -35, -35, -35, -35, -35, -35, -35,
	25, 16, 6, -34, 25, -35, -35, 25, -35, -2,
	-2, -2, -2, -2, -32, 9, -27, -28, -27, -26,
	-17, -35, 27, -26, 9, 30, -30, -31, -2, 5,
	-35, -35, -37, 25, 24, -2, 11, 25, 27, 11,
	-35, -35, -35, -35, 27, 8, 4, 5, 24, 25,
	-26, -26, 6, 4, 5, -36, 9, -36, 9, 18,
	-35, -37, 11, -35, -35, -36, -35, -36, 17, -2,
	11, 32, -35, 32, -35, -35, -2, 5, -2, -35,
	27, 18, -35, 17, -2, -35, 9, -2, 9,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 0, 11, 4, 11, 15,
	-2, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, -2, 35,
	36, 37, 38, 11, 11, 11, 74, 75, 11, 11,
	11, 11, 11, 11, 94, 0, 0, 12, 6, 0,
	0, 0, 0, 11, 11, 0, 0, 0, 0, 47,
	13, 90, 91, 0, 88, 89, 11, 11, 11, 11,
	39, 11, 11, -2, 43, 44, 11, 11, 0, 0,
	0, -2, -2, -2, -2, 11, 48, 49, 11, 92,
	93, 0, 0, 0, 0, 0, 47, 0, 0, 56,
	7, 11, -2, 0, 7, 0, 0, 14, 79, -2,
	-2, -2, -2, -2, 11, 11, 40, 52, 57, 0,
	0, 61, 0, 0, 11, 76, 11, 11, 11, -2,
	0, 0, 7, 8, 9, -2, 58, 7, 0, 60,
	0, 13, 13, 0, 11, 41, 45, 46, 53, 54,
	52, 0, -2, 50, 51, 11, 13, 11, 13, 0,
	0, 0, 59, 62, 0, 11, 0, 11, 11, -2,
	55, 77, 0, 78, 0, 0, 11, 0, -2, 0,
	11, 0, 0, 11, -2, 0, 85, -2, 82,
}
var RubyTok1 = []int{

	1,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36,
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
		//line parser.y:118
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:120
		{
			Statements = []ast.Node{}
		}
	case 3:
		//line parser.y:122
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 4:
		//line parser.y:128
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 5:
		//line parser.y:134
		{
			RubyVAL.genericValue = nil
		}
	case 6:
		//line parser.y:135
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 7:
		//line parser.y:138
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 8:
		//line parser.y:140
		{ /* do nothing */
		}
	case 9:
		//line parser.y:142
		{ /* do nothing */
		}
	case 10:
		//line parser.y:144
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 39:
		//line parser.y:158
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 40:
		//line parser.y:165
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 41:
		//line parser.y:174
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 42:
		//line parser.y:178
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 43:
		//line parser.y:183
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 44:
		//line parser.y:185
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 45:
		//line parser.y:187
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 46:
		//line parser.y:189
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 47:
		//line parser.y:191
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 48:
		//line parser.y:193
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 49:
		//line parser.y:195
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 50:
		//line parser.y:197
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 51:
		//line parser.y:199
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 55:
		//line parser.y:209
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 56:
		//line parser.y:217
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 57:
		//line parser.y:218
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 58:
		//line parser.y:221
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 59:
		//line parser.y:228
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 60:
		//line parser.y:238
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 61:
		//line parser.y:247
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 62:
		//line parser.y:253
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 63:
		//line parser.y:261
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 64:
		//line parser.y:265
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 65:
		//line parser.y:270
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 66:
		//line parser.y:277
		{ // FIXME: this is a hack. If this rule is not present, "foo = 5" will not parse
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 67:
		//line parser.y:284
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 68:
		//line parser.y:285
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 69:
		//line parser.y:286
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 70:
		//line parser.y:287
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 71:
		//line parser.y:290
		{
			RubyVAL.genericValue = ast.Addition{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 72:
		//line parser.y:298
		{
			RubyVAL.genericValue = ast.Subtraction{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 73:
		//line parser.y:306
		{
			RubyVAL.genericValue = ast.Multiplication{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 74:
		//line parser.y:313
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 75:
		//line parser.y:314
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 76:
		//line parser.y:316
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 77:
		//line parser.y:320
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 78:
		//line parser.y:328
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 79:
		//line parser.y:336
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 80:
		//line parser.y:338
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 81:
		//line parser.y:340
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 82:
		//line parser.y:342
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 83:
		//line parser.y:345
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 84:
		//line parser.y:352
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 85:
		//line parser.y:359
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 86:
		//line parser.y:367
		{
			RubyVAL.genericValue = nil
		}
	case 87:
		//line parser.y:368
		{
			RubyVAL.genericValue = nil
		}
	case 88:
		//line parser.y:371
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 89:
		//line parser.y:373
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 90:
		//line parser.y:376
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 91:
		//line parser.y:378
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 92:
		//line parser.y:381
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 93:
		//line parser.y:383
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 94:
		//line parser.y:386
		{
			RubyVAL.genericValue = ast.FileNameConstReference{}
		}
	}
	goto Rubystack /* stack new state and value */
}
