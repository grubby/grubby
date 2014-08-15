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
const DO = 57352
const DEF = 57353
const END = 57354
const CLASS = 57355
const MODULE = 57356
const TRUE = 57357
const FALSE = 57358
const LESSTHAN = 57359
const GREATERTHAN = 57360
const EQUALTO = 57361
const BANG = 57362
const COMPLEMENT = 57363
const POSITIVE = 57364
const NEGATIVE = 57365
const STAR = 57366
const WHITESPACE = 57367
const NEWLINE = 57368
const SEMICOLON = 57369
const COLON = 57370
const DOT = 57371
const PIPE = 57372
const LBRACKET = 57373
const RBRACKET = 57374
const LBRACE = 57375
const RBRACE = 57376
const DOLLARSIGN = 57377
const ATSIGN = 57378
const FILE_CONST_REF = 57379
const EOF = 57380

var RubyToknames = []string{
	"NODE",
	"REF",
	"CAPITAL_REF",
	"LPAREN",
	"RPAREN",
	"COMMA",
	"DO",
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
	"PIPE",
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

//line parser.y:409

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 11,
	-1, 10,
	19, 11,
	-2, 16,
	-1, 28,
	1, 34,
	9, 34,
	12, 34,
	24, 34,
	26, 34,
	34, 34,
	38, 34,
	-2, 11,
	-1, 71,
	9, 11,
	-2, 41,
	-1, 74,
	9, 11,
	-2, 43,
	-1, 75,
	22, 11,
	23, 11,
	24, 11,
	-2, 44,
	-1, 77,
	1, 34,
	9, 34,
	12, 34,
	24, 34,
	26, 34,
	34, 34,
	38, 34,
	-2, 11,
	-1, 85,
	22, 11,
	23, 11,
	24, 11,
	-2, 73,
	-1, 86,
	22, 11,
	23, 11,
	24, 11,
	-2, 74,
	-1, 87,
	22, 11,
	23, 11,
	24, 11,
	-2, 75,
	-1, 88,
	22, 11,
	23, 11,
	24, 11,
	-2, 76,
	-1, 109,
	28, 69,
	-2, 11,
	-1, 115,
	22, 11,
	23, 11,
	24, 11,
	-2, 77,
	-1, 116,
	22, 11,
	23, 11,
	24, 11,
	-2, 78,
	-1, 117,
	22, 11,
	23, 11,
	24, 11,
	-2, 79,
	-1, 118,
	22, 11,
	23, 11,
	24, 11,
	-2, 72,
	-1, 120,
	22, 11,
	23, 11,
	24, 11,
	-2, 71,
	-1, 125,
	24, 11,
	-2, 10,
	-1, 142,
	24, 34,
	-2, 11,
	-1, 156,
	22, 11,
	23, 11,
	24, 11,
	-2, 46,
	-1, 166,
	28, 70,
	-2, 11,
	-1, 183,
	22, 11,
	23, 11,
	24, 11,
	-2, 89,
	-1, 192,
	22, 11,
	23, 11,
	24, 11,
	-2, 86,
	-1, 198,
	26, 90,
	34, 90,
	-2, 11,
	-1, 201,
	26, 87,
	34, 87,
	-2, 11,
}

const RubyNprod = 102
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 506

var RubyAct = []int{

	103, 146, 5, 76, 92, 104, 125, 47, 70, 48,
	89, 49, 8, 7, 61, 62, 82, 2, 187, 185,
	194, 6, 6, 137, 155, 51, 51, 149, 135, 50,
	6, 113, 162, 163, 52, 53, 54, 184, 108, 55,
	56, 57, 58, 59, 60, 63, 138, 107, 6, 4,
	162, 163, 148, 111, 83, 83, 195, 75, 6, 66,
	67, 68, 85, 86, 87, 88, 69, 95, 96, 97,
	98, 84, 99, 100, 101, 102, 202, 197, 50, 200,
	173, 105, 106, 66, 67, 68, 66, 67, 68, 182,
	112, 172, 6, 114, 170, 6, 158, 137, 122, 119,
	159, 160, 115, 116, 117, 118, 79, 120, 166, 83,
	134, 109, 121, 132, 127, 131, 191, 136, 93, 94,
	143, 141, 144, 145, 81, 133, 129, 128, 80, 73,
	64, 65, 167, 168, 91, 90, 110, 71, 151, 74,
	152, 153, 154, 50, 140, 139, 1, 157, 130, 161,
	156, 24, 164, 23, 165, 22, 174, 169, 171, 21,
	20, 19, 18, 27, 14, 30, 175, 177, 13, 12,
	178, 31, 180, 15, 10, 179, 11, 181, 32, 25,
	186, 183, 188, 189, 17, 26, 16, 3, 0, 0,
	0, 193, 0, 190, 0, 196, 192, 0, 199, 0,
	9, 28, 29, 198, 0, 0, 201, 33, 176, 34,
	35, 36, 37, 0, 0, 0, 38, 39, 40, 41,
	0, 124, 123, 0, 0, 0, 0, 42, 0, 43,
	0, 46, 45, 44, 9, 28, 29, 0, 0, 0,
	0, 33, 150, 34, 35, 36, 37, 0, 0, 0,
	38, 39, 40, 41, 0, 124, 123, 0, 0, 0,
	0, 42, 0, 43, 0, 46, 45, 44, 9, 28,
	29, 0, 0, 0, 0, 33, 147, 34, 35, 36,
	37, 0, 0, 0, 38, 39, 40, 41, 0, 124,
	123, 0, 0, 0, 0, 42, 0, 43, 0, 46,
	45, 44, 9, 28, 29, 0, 0, 0, 0, 33,
	126, 34, 35, 36, 37, 0, 0, 0, 38, 39,
	40, 41, 0, 124, 123, 0, 0, 0, 0, 42,
	0, 43, 0, 46, 45, 44, 78, 77, 29, 73,
	0, 0, 79, 33, 0, 34, 35, 36, 37, 0,
	0, 72, 38, 39, 40, 41, 0, 0, 0, 0,
	0, 0, 0, 42, 0, 43, 0, 46, 45, 44,
	9, 28, 29, 0, 0, 0, 0, 33, 0, 34,
	35, 36, 37, 0, 0, 0, 38, 39, 40, 41,
	0, 124, 123, 0, 0, 0, 0, 42, 0, 43,
	0, 46, 45, 44, 9, 28, 29, 0, 0, 0,
	79, 33, 0, 34, 35, 36, 37, 0, 0, 0,
	38, 39, 40, 41, 0, 0, 0, 0, 0, 0,
	0, 42, 0, 43, 0, 46, 45, 44, 9, 28,
	29, 0, 0, 0, 0, 33, 0, 34, 35, 36,
	37, 0, 0, 0, 38, 39, 40, 41, 0, 0,
	0, 0, 0, 0, 0, 42, 0, 43, 0, 46,
	45, 44, 9, 142, 29, 0, 0, 0, 0, 33,
	0, 34, 35, 36, 37, 0, 0, 0, 38, 39,
	40, 41, 0, 0, 0, 0, 0, 0, 0, 42,
	0, 43, 0, 46, 45, 44,
}
var RubyPact = []int{

	-21, 23, -1000, -25, -1000, 434, 33, -1000, 33, -1000,
	33, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -3, -1000,
	-1000, -1000, -1000, 33, 33, 33, -1000, -1000, 33, 33,
	33, 33, 33, 33, -1000, 9, 125, -1000, 64, 47,
	332, 123, 119, 33, 33, 434, 434, 434, 434, 130,
	-1000, -1000, -1000, 113, -1000, -1000, 33, 33, 33, 33,
	-1000, 33, 33, 33, 33, 33, -1000, -3, -1000, -1000,
	33, 33, 21, 105, 27, 33, 33, 33, 33, 33,
	-1000, -1000, 5, -1000, -1000, 434, 434, 434, 434, 90,
	434, 130, 89, 64, 298, 122, 122, -1000, 33, 33,
	0, -1000, 14, -1000, 468, 33, 33, 33, 33, 33,
	33, 33, 33, -1000, -1000, 33, -1000, -1000, -1000, -1000,
	-1000, -1000, 264, 26, -1000, -1, 230, 33, -1000, 33,
	33, 33, -4, 400, 88, 96, 7, -1000, -1000, 102,
	-1000, 128, 85, 82, 61, 33, 33, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 366, 196, 33, -1000, -1000, 5,
	-1000, 5, -1000, 71, 434, 25, -1000, -1000, -15, 5,
	-16, 5, 33, 33, -1000, -1000, 434, -1000, 111, 434,
	33, -8, 33, 37, 33, 59, 434, 33, 70, 434,
	-1000, 67, -1000,
}
var RubyPgo = []int{

	0, 188, 6, 187, 186, 185, 3, 184, 179, 178,
	176, 174, 173, 171, 169, 168, 165, 164, 16, 163,
	162, 161, 160, 159, 155, 153, 151, 5, 8, 148,
	146, 145, 144, 10, 139, 137, 136, 0, 4, 1,
}
var RubyR1 = []int{

	0, 30, 30, 30, 30, 3, 3, 27, 27, 27,
	27, 37, 37, 38, 38, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 11, 11, 11, 11, 11, 10,
	10, 10, 28, 28, 35, 35, 35, 35, 34, 34,
	34, 34, 34, 33, 33, 33, 33, 33, 39, 39,
	39, 14, 29, 29, 15, 15, 17, 18, 18, 36,
	36, 12, 12, 20, 21, 22, 23, 24, 25, 26,
	4, 7, 8, 5, 5, 31, 31, 31, 31, 32,
	32, 32, 1, 1, 9, 9, 16, 16, 13, 13,
	19, 6,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 3, 0, 2, 2,
	2, 0, 2, 0, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
	5, 3, 5, 1, 1, 1, 5, 5, 1, 1,
	5, 5, 5, 0, 1, 1, 5, 5, 0, 2,
	2, 9, 0, 1, 6, 8, 6, 3, 6, 1,
	4, 5, 5, 3, 3, 3, 3, 5, 5, 5,
	1, 1, 5, 9, 9, 0, 6, 11, 12, 4,
	9, 10, 0, 1, 2, 2, 2, 2, 3, 3,
	1, 3,
}
var RubyChk = []int{

	-1000, -30, 38, -3, 26, -37, 25, 38, -2, 4,
	-11, -10, -14, -15, -17, -12, -4, -7, -20, -21,
	-22, -23, -24, -25, -26, -8, -5, -19, 5, 6,
	-16, -13, -9, 11, 13, 14, 15, 16, 20, 21,
	22, 23, 31, 33, 37, 36, 35, -37, -37, -37,
	-37, 29, -37, -37, -37, -37, -37, -37, -37, -37,
	-37, 5, 6, 36, 5, 6, 22, 23, 24, 19,
	-28, -35, 19, 7, -34, -2, -6, 5, 4, 10,
	5, 5, -18, -37, -18, -2, -2, -2, -2, -33,
	5, 4, -38, 5, 6, -37, -37, -37, -37, -37,
	-37, -37, -37, -37, -27, -37, -37, 26, 17, 6,
	-36, 26, -37, 26, -37, -2, -2, -2, -2, 9,
	-2, -33, 9, 26, 25, -2, 12, -28, 5, 4,
	-29, -28, -27, -18, -37, 28, -27, 9, 32, -31,
	-32, -2, 5, -37, -37, -37, -39, 12, 26, 28,
	12, -37, -37, -37, -37, 28, -2, -6, 8, 4,
	5, -6, 25, 26, -27, -27, 6, 4, 5, -38,
	9, -38, 9, 19, -37, -39, 12, -37, -37, -38,
	-37, -38, 18, -2, 12, 34, -37, 34, -37, -37,
	-2, 5, -2, -37, 28, 19, -37, 18, -2, -37,
	9, -2, 9,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 0, 11, 4, 11, 15,
	-2, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, -2, 35,
	36, 37, 38, 11, 11, 11, 80, 81, 11, 11,
	11, 11, 11, 11, 100, 0, 0, 12, 6, 0,
	0, 0, 0, 11, 11, 0, 0, 0, 0, 53,
	13, 96, 97, 0, 94, 95, 11, 11, 11, 11,
	39, -2, 11, 11, -2, -2, 45, -2, 15, 7,
	11, 11, 0, 0, 0, -2, -2, -2, -2, 11,
	54, 55, 11, 98, 99, 0, 0, 0, 0, 0,
	0, 53, 0, 0, 0, 0, 62, 7, 11, -2,
	0, 7, 0, 14, 85, -2, -2, -2, -2, 11,
	-2, 11, 11, 8, 9, -2, 101, 40, 48, 49,
	58, 63, 0, 0, 67, 0, 0, 11, 82, 11,
	11, 11, -2, 0, 0, 0, 7, 64, 7, 0,
	66, 0, 13, 13, 0, 11, -2, 47, 42, 50,
	51, 52, 59, 60, 58, 0, -2, 56, 57, 11,
	13, 11, 13, 0, 0, 0, 65, 68, 0, 11,
	0, 11, 11, -2, 61, 83, 0, 84, 0, 0,
	11, 0, -2, 0, 11, 0, 0, 11, -2, 0,
	91, -2, 88,
}
var RubyTok1 = []int{

	1,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38,
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
		//line parser.y:122
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:124
		{
			Statements = []ast.Node{}
		}
	case 3:
		//line parser.y:126
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 4:
		//line parser.y:132
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 5:
		//line parser.y:138
		{
			RubyVAL.genericValue = nil
		}
	case 6:
		//line parser.y:139
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 7:
		//line parser.y:142
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 8:
		//line parser.y:144
		{ /* do nothing */
		}
	case 9:
		//line parser.y:146
		{ /* do nothing */
		}
	case 10:
		//line parser.y:148
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
		//line parser.y:162
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 40:
		//line parser.y:169
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 41:
		//line parser.y:177
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 42:
		//line parser.y:185
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 43:
		//line parser.y:187
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 44:
		//line parser.y:190
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 45:
		//line parser.y:192
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 46:
		//line parser.y:194
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 47:
		//line parser.y:196
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 48:
		//line parser.y:199
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 49:
		//line parser.y:201
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 50:
		//line parser.y:203
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 51:
		//line parser.y:205
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 52:
		//line parser.y:207
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 53:
		//line parser.y:209
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 54:
		//line parser.y:211
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 55:
		//line parser.y:213
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 56:
		//line parser.y:215
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 57:
		//line parser.y:217
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 61:
		//line parser.y:227
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 62:
		//line parser.y:235
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 63:
		//line parser.y:236
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 64:
		//line parser.y:239
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 65:
		//line parser.y:246
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 66:
		//line parser.y:256
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 67:
		//line parser.y:265
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 68:
		//line parser.y:271
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 69:
		//line parser.y:279
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 70:
		//line parser.y:283
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 71:
		//line parser.y:288
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 72:
		//line parser.y:295
		{ // FIXME: this is a hack. If this rule is not present, "foo = 5" will not parse
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 73:
		//line parser.y:302
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 74:
		//line parser.y:303
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 75:
		//line parser.y:304
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 76:
		//line parser.y:305
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 77:
		//line parser.y:308
		{
			RubyVAL.genericValue = ast.Addition{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 78:
		//line parser.y:316
		{
			RubyVAL.genericValue = ast.Subtraction{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 79:
		//line parser.y:324
		{
			RubyVAL.genericValue = ast.Multiplication{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 80:
		//line parser.y:331
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 81:
		//line parser.y:332
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 82:
		//line parser.y:334
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 83:
		//line parser.y:338
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 84:
		//line parser.y:346
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 85:
		//line parser.y:354
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 86:
		//line parser.y:356
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 87:
		//line parser.y:358
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 88:
		//line parser.y:360
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 89:
		//line parser.y:363
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 90:
		//line parser.y:370
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 91:
		//line parser.y:377
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 92:
		//line parser.y:385
		{
			RubyVAL.genericValue = nil
		}
	case 93:
		//line parser.y:386
		{
			RubyVAL.genericValue = nil
		}
	case 94:
		//line parser.y:389
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 95:
		//line parser.y:391
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 96:
		//line parser.y:394
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 97:
		//line parser.y:396
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 98:
		//line parser.y:399
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 99:
		//line parser.y:401
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 100:
		//line parser.y:404
		{
			RubyVAL.genericValue = ast.FileNameConstReference{}
		}
	case 101:
		//line parser.y:407
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	}
	goto Rubystack /* stack new state and value */
}
