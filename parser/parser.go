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
const IF = 57355
const ELSE = 57356
const ELSIF = 57357
const CLASS = 57358
const MODULE = 57359
const TRUE = 57360
const FALSE = 57361
const LESSTHAN = 57362
const GREATERTHAN = 57363
const EQUALTO = 57364
const BANG = 57365
const COMPLEMENT = 57366
const POSITIVE = 57367
const NEGATIVE = 57368
const STAR = 57369
const WHITESPACE = 57370
const NEWLINE = 57371
const SEMICOLON = 57372
const COLON = 57373
const DOT = 57374
const PIPE = 57375
const LBRACKET = 57376
const RBRACKET = 57377
const LBRACE = 57378
const RBRACE = 57379
const DOLLARSIGN = 57380
const ATSIGN = 57381
const FILE_CONST_REF = 57382
const EOF = 57383

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
	"IF",
	"ELSE",
	"ELSIF",
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

//line parser.y:452

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 11,
	-1, 10,
	22, 11,
	-2, 16,
	-1, 29,
	1, 35,
	9, 35,
	12, 35,
	14, 35,
	27, 35,
	29, 35,
	37, 35,
	41, 35,
	-2, 11,
	-1, 74,
	9, 11,
	-2, 42,
	-1, 77,
	9, 11,
	-2, 44,
	-1, 78,
	25, 11,
	26, 11,
	27, 11,
	-2, 45,
	-1, 80,
	1, 35,
	9, 35,
	12, 35,
	14, 35,
	27, 35,
	29, 35,
	37, 35,
	41, 35,
	-2, 11,
	-1, 82,
	33, 11,
	-2, 7,
	-1, 88,
	25, 11,
	26, 11,
	27, 11,
	-2, 74,
	-1, 89,
	25, 11,
	26, 11,
	27, 11,
	-2, 75,
	-1, 90,
	25, 11,
	26, 11,
	27, 11,
	-2, 76,
	-1, 91,
	25, 11,
	26, 11,
	27, 11,
	-2, 77,
	-1, 96,
	27, 11,
	-2, 7,
	-1, 114,
	31, 70,
	-2, 11,
	-1, 121,
	25, 11,
	26, 11,
	27, 11,
	-2, 78,
	-1, 122,
	25, 11,
	26, 11,
	27, 11,
	-2, 79,
	-1, 123,
	25, 11,
	26, 11,
	27, 11,
	-2, 80,
	-1, 124,
	25, 11,
	26, 11,
	27, 11,
	-2, 73,
	-1, 126,
	25, 11,
	26, 11,
	27, 11,
	-2, 72,
	-1, 131,
	27, 11,
	-2, 10,
	-1, 134,
	33, 105,
	-2, 54,
	-1, 150,
	27, 35,
	-2, 11,
	-1, 158,
	33, 106,
	-2, 55,
	-1, 171,
	25, 11,
	26, 11,
	27, 11,
	-2, 47,
	-1, 185,
	31, 71,
	-2, 11,
	-1, 206,
	25, 11,
	26, 11,
	27, 11,
	-2, 90,
	-1, 215,
	33, 107,
	-2, 58,
	-1, 218,
	25, 11,
	26, 11,
	27, 11,
	-2, 87,
	-1, 224,
	29, 91,
	37, 91,
	-2, 11,
	-1, 227,
	29, 88,
	37, 88,
	-2, 11,
}

const RubyNprod = 112
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 677

var RubyAct = []int{

	107, 79, 5, 160, 95, 7, 108, 49, 85, 50,
	92, 51, 2, 73, 64, 65, 212, 210, 6, 179,
	145, 169, 53, 6, 134, 220, 72, 53, 163, 143,
	52, 6, 118, 209, 196, 54, 55, 56, 181, 182,
	57, 58, 59, 60, 61, 62, 146, 63, 66, 181,
	182, 178, 113, 6, 4, 221, 86, 86, 69, 70,
	71, 112, 6, 192, 162, 87, 69, 70, 71, 228,
	99, 100, 101, 102, 116, 103, 104, 105, 106, 226,
	131, 52, 223, 109, 110, 111, 8, 205, 6, 69,
	70, 71, 197, 117, 174, 175, 119, 191, 6, 217,
	82, 173, 145, 120, 189, 137, 136, 128, 76, 125,
	97, 98, 67, 68, 86, 142, 127, 186, 215, 140,
	186, 187, 141, 144, 135, 139, 153, 185, 154, 155,
	94, 158, 84, 78, 156, 94, 93, 83, 88, 89,
	90, 91, 114, 177, 96, 159, 165, 115, 166, 167,
	168, 52, 74, 77, 148, 172, 157, 176, 147, 170,
	180, 1, 138, 133, 24, 23, 22, 183, 21, 184,
	193, 188, 190, 20, 19, 18, 27, 14, 31, 13,
	121, 122, 123, 124, 195, 126, 200, 198, 12, 201,
	32, 203, 15, 10, 202, 28, 204, 11, 208, 33,
	149, 25, 17, 211, 26, 213, 214, 16, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 219, 9, 29,
	30, 222, 0, 0, 225, 34, 151, 46, 152, 0,
	35, 36, 37, 38, 171, 0, 0, 39, 40, 41,
	42, 0, 130, 129, 0, 0, 0, 0, 43, 0,
	44, 0, 48, 47, 45, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 29, 30, 0, 0, 0,
	0, 34, 207, 46, 206, 0, 35, 36, 37, 38,
	0, 0, 0, 39, 40, 41, 42, 0, 130, 129,
	0, 0, 216, 0, 43, 218, 44, 0, 48, 47,
	45, 0, 0, 224, 0, 0, 227, 9, 29, 30,
	0, 0, 0, 0, 34, 199, 46, 0, 0, 35,
	36, 37, 38, 0, 0, 0, 39, 40, 41, 42,
	0, 130, 129, 0, 0, 0, 0, 43, 0, 44,
	0, 48, 47, 45, 9, 29, 30, 0, 0, 0,
	0, 34, 194, 46, 0, 0, 35, 36, 37, 38,
	0, 0, 0, 39, 40, 41, 42, 0, 130, 129,
	0, 0, 0, 0, 43, 0, 44, 0, 48, 47,
	45, 9, 29, 30, 0, 0, 0, 0, 34, 164,
	46, 0, 0, 35, 36, 37, 38, 0, 0, 0,
	39, 40, 41, 42, 0, 130, 129, 0, 0, 0,
	0, 43, 0, 44, 0, 48, 47, 45, 9, 29,
	30, 0, 0, 0, 0, 34, 161, 46, 0, 0,
	35, 36, 37, 38, 0, 0, 0, 39, 40, 41,
	42, 0, 130, 129, 0, 0, 0, 0, 43, 0,
	44, 0, 48, 47, 45, 9, 29, 30, 0, 0,
	0, 0, 34, 132, 46, 0, 0, 35, 36, 37,
	38, 0, 0, 0, 39, 40, 41, 42, 0, 130,
	129, 0, 0, 0, 0, 43, 0, 44, 0, 48,
	47, 45, 81, 80, 30, 76, 0, 0, 82, 34,
	0, 46, 0, 0, 35, 36, 37, 38, 0, 0,
	75, 39, 40, 41, 42, 0, 0, 0, 0, 0,
	0, 0, 43, 0, 44, 0, 48, 47, 45, 9,
	29, 30, 0, 0, 0, 0, 34, 0, 46, 0,
	0, 35, 36, 37, 38, 0, 0, 0, 39, 40,
	41, 42, 0, 130, 129, 0, 0, 0, 0, 43,
	0, 44, 0, 48, 47, 45, 9, 29, 30, 0,
	0, 0, 82, 34, 0, 46, 0, 0, 35, 36,
	37, 38, 0, 0, 0, 39, 40, 41, 42, 0,
	0, 0, 0, 0, 0, 0, 43, 0, 44, 0,
	48, 47, 45, 9, 29, 30, 0, 0, 0, 0,
	34, 0, 46, 0, 0, 35, 36, 37, 38, 0,
	0, 0, 39, 40, 41, 42, 0, 0, 0, 0,
	0, 0, 0, 43, 0, 44, 0, 48, 47, 45,
	9, 150, 30, 0, 0, 0, 0, 34, 0, 46,
	0, 0, 35, 36, 37, 38, 0, 0, 0, 39,
	40, 41, 42, 0, 0, 0, 0, 0, 0, 0,
	43, 0, 44, 0, 48, 47, 45,
}
var RubyPact = []int{

	-29, 25, -1000, -36, -1000, 599, 34, -1000, 34, -1000,
	34, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -5,
	-1000, -1000, -1000, -1000, 34, 34, 34, -1000, -1000, 34,
	34, 34, 34, 34, 34, -1000, 34, 9, 107, -1000,
	64, 4, 488, 132, 127, 34, 34, 599, 599, 599,
	599, 131, -1000, 599, -1000, -1000, 105, -1000, -1000, 34,
	34, 34, 34, -1000, 34, 34, 34, 34, 34, -1000,
	-5, -1000, 34, 34, 34, 32, 136, 45, 34, 34,
	34, 34, 34, -1000, -1000, 3, 34, -1000, -1000, 599,
	599, 599, 599, 100, 599, 131, 98, 64, 451, -9,
	101, 101, -1000, 34, 34, -2, -1000, 11, -1000, 636,
	214, 34, 34, 34, 34, 34, 34, 34, 34, -1000,
	-1000, 34, -1000, 34, 126, -1000, -1000, -1000, -1000, -1000,
	414, 35, -1000, -3, 377, 34, -1000, 34, 34, 34,
	-10, -1000, -1000, 562, 93, 90, 22, -14, -1000, 34,
	10, -1000, -1000, 121, -1000, 116, 95, 88, 41, 34,
	340, 34, -1000, -1000, -1000, -1000, -1000, 5, -1000, -1000,
	83, -1000, -1000, 525, 303, 34, -1000, -1000, 3, -1000,
	3, -1000, 66, 599, -1000, 260, -1000, 34, 21, -1000,
	-1000, -20, 3, -21, 3, 34, 34, -1000, 113, -1000,
	-1000, 599, -1000, 94, 599, -1000, 34, -6, 34, 33,
	34, 61, 599, 34, 70, 599, -1000, 60, -1000,
}
var RubyPgo = []int{

	0, 209, 80, 208, 207, 204, 1, 202, 201, 199,
	197, 195, 193, 192, 190, 188, 179, 178, 177, 8,
	176, 175, 174, 173, 168, 166, 165, 164, 6, 13,
	163, 162, 161, 158, 10, 156, 154, 153, 152, 147,
	0, 4, 3, 143,
}
var RubyR1 = []int{

	0, 32, 32, 32, 32, 3, 3, 28, 28, 28,
	28, 40, 40, 41, 41, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 12, 12, 12, 12, 12,
	10, 10, 10, 29, 29, 38, 38, 38, 38, 37,
	37, 37, 37, 37, 34, 34, 34, 34, 34, 42,
	42, 42, 15, 31, 31, 16, 16, 18, 19, 19,
	39, 39, 13, 13, 21, 22, 23, 24, 25, 26,
	27, 4, 7, 8, 5, 5, 33, 33, 33, 33,
	36, 36, 36, 1, 1, 9, 9, 17, 17, 14,
	14, 20, 6, 6, 30, 35, 35, 35, 43, 43,
	11, 11,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 3, 0, 2, 2,
	2, 0, 2, 0, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 5, 3, 5, 1, 1, 1, 5, 5, 1,
	1, 5, 5, 5, 0, 1, 1, 5, 5, 0,
	2, 2, 9, 0, 1, 6, 8, 6, 3, 6,
	1, 4, 5, 5, 3, 3, 3, 3, 5, 5,
	5, 1, 1, 5, 9, 9, 0, 6, 11, 12,
	4, 9, 10, 0, 1, 2, 2, 2, 2, 3,
	3, 1, 3, 7, 3, 0, 1, 5, 1, 2,
	5, 7,
}
var RubyChk = []int{

	-1000, -32, 41, -3, 29, -40, 28, 41, -2, 4,
	-12, -10, -15, -16, -18, -13, -4, -7, -21, -22,
	-23, -24, -25, -26, -27, -8, -5, -20, -11, 5,
	6, -17, -14, -9, 11, 16, 17, 18, 19, 23,
	24, 25, 26, 34, 36, 40, 13, 39, 38, -40,
	-40, -40, -40, 32, -40, -40, -40, -40, -40, -40,
	-40, -40, -40, -40, 5, 6, 39, 5, 6, 25,
	26, 27, 22, -29, -38, 22, 7, -37, -2, -6,
	5, 4, 10, 5, 5, -19, -40, -19, -2, -2,
	-2, -2, -34, 5, 4, -41, -2, 5, 6, -40,
	-40, -40, -40, -40, -40, -40, -40, -40, -28, -40,
	-40, -40, 29, 20, 6, -39, 29, -40, 29, -40,
	-28, -2, -2, -2, -2, 9, -2, -34, 9, 29,
	28, -2, 12, -30, 33, -29, 5, 4, -31, -29,
	-28, -19, -40, 31, -28, 9, 35, -33, -36, -2,
	5, 12, 14, -40, -40, -40, -40, -35, 5, -34,
	-42, 12, 29, 31, 12, -40, -40, -40, -40, 31,
	-28, -2, -6, 8, 4, 5, -6, -43, 29, 33,
	-40, 28, 29, -28, -28, 6, 4, 5, -41, 9,
	-41, 9, 22, -40, 12, -28, 29, 9, -42, 12,
	-40, -40, -41, -40, -41, 21, -2, 12, -40, 12,
	37, -40, 37, -40, -40, 5, -2, 5, -2, -40,
	31, 22, -40, 21, -2, -40, 9, -2, 9,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 0, 11, 4, 11, 15,
	-2, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, -2,
	36, 37, 38, 39, 11, 11, 11, 81, 82, 11,
	11, 11, 11, 11, 11, 101, 11, 0, 0, 12,
	6, 0, 0, 0, 0, 11, 11, 0, 0, 0,
	0, 54, 13, 0, 97, 98, 0, 95, 96, 11,
	11, 11, 11, 40, -2, 11, 11, -2, -2, 46,
	-2, 15, -2, 11, 11, 0, 0, 0, -2, -2,
	-2, -2, 11, 55, 56, 11, -2, 99, 100, 0,
	0, 0, 0, 0, 0, 54, 0, 0, 0, 0,
	0, 63, 7, 11, -2, 0, 7, 0, 14, 86,
	0, -2, -2, -2, -2, 11, -2, 11, 11, 8,
	9, -2, 102, 11, -2, 41, 49, 50, 59, 64,
	0, 0, 68, 0, 0, 11, 83, 11, 11, 11,
	-2, 110, 7, 0, 0, 0, 0, 0, -2, 11,
	7, 65, 7, 0, 67, 0, 13, 13, 0, 11,
	0, -2, 48, 43, 51, 52, 53, 7, 108, 104,
	0, 60, 61, 59, 0, -2, 57, 58, 11, 13,
	11, 13, 0, 0, 111, 0, 109, 11, 0, 66,
	69, 0, 11, 0, 11, 11, -2, 103, 0, 62,
	84, 0, 85, 0, 0, -2, 11, 0, -2, 0,
	11, 0, 0, 11, -2, 0, 92, -2, 89,
}
var RubyTok1 = []int{

	1,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
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
		//line parser.y:128
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:130
		{
			Statements = []ast.Node{}
		}
	case 3:
		//line parser.y:132
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 4:
		//line parser.y:138
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 5:
		//line parser.y:144
		{
			RubyVAL.genericValue = nil
		}
	case 6:
		//line parser.y:145
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 7:
		//line parser.y:148
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 8:
		//line parser.y:150
		{ /* do nothing */
		}
	case 9:
		//line parser.y:152
		{ /* do nothing */
		}
	case 10:
		//line parser.y:154
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 40:
		//line parser.y:168
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 41:
		//line parser.y:175
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 42:
		//line parser.y:183
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 43:
		//line parser.y:191
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 44:
		//line parser.y:193
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 45:
		//line parser.y:196
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 46:
		//line parser.y:198
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 47:
		//line parser.y:200
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 48:
		//line parser.y:202
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 49:
		//line parser.y:205
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 50:
		//line parser.y:207
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 51:
		//line parser.y:209
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 52:
		//line parser.y:211
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 53:
		//line parser.y:213
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 54:
		//line parser.y:215
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 55:
		//line parser.y:217
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 56:
		//line parser.y:219
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 57:
		//line parser.y:221
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 58:
		//line parser.y:223
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 62:
		//line parser.y:233
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 63:
		//line parser.y:241
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 64:
		//line parser.y:242
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 65:
		//line parser.y:245
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 66:
		//line parser.y:252
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 67:
		//line parser.y:262
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 68:
		//line parser.y:271
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 69:
		//line parser.y:277
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 70:
		//line parser.y:285
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 71:
		//line parser.y:289
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 72:
		//line parser.y:294
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 73:
		//line parser.y:301
		{ // FIXME: this is a hack. If this rule is not present, "foo = 5" will not parse
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 74:
		//line parser.y:308
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 75:
		//line parser.y:309
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 76:
		//line parser.y:310
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 77:
		//line parser.y:311
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 78:
		//line parser.y:314
		{
			RubyVAL.genericValue = ast.Addition{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 79:
		//line parser.y:322
		{
			RubyVAL.genericValue = ast.Subtraction{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 80:
		//line parser.y:330
		{
			RubyVAL.genericValue = ast.Multiplication{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 81:
		//line parser.y:337
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 82:
		//line parser.y:338
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 83:
		//line parser.y:340
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 84:
		//line parser.y:344
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 85:
		//line parser.y:352
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 86:
		//line parser.y:360
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 87:
		//line parser.y:362
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 88:
		//line parser.y:364
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 89:
		//line parser.y:366
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 90:
		//line parser.y:369
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 91:
		//line parser.y:376
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 92:
		//line parser.y:383
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 93:
		//line parser.y:391
		{
			RubyVAL.genericValue = nil
		}
	case 94:
		//line parser.y:392
		{
			RubyVAL.genericValue = nil
		}
	case 95:
		//line parser.y:395
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 96:
		//line parser.y:397
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 97:
		//line parser.y:400
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 98:
		//line parser.y:402
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 99:
		//line parser.y:405
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 100:
		//line parser.y:407
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 101:
		//line parser.y:410
		{
			RubyVAL.genericValue = ast.FileNameConstReference{}
		}
	case 102:
		//line parser.y:413
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 103:
		//line parser.y:415
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 104:
		//line parser.y:423
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 105:
		//line parser.y:425
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 106:
		//line parser.y:427
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 107:
		//line parser.y:429
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 110:
		//line parser.y:434
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 111:
		//line parser.y:441
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-3].genericSlice,
				Else: ast.IfBlock{
					Condition: ast.Boolean{Value: true},
					Body:      RubyS[Rubypt-1].genericSlice,
				},
			}
		}
	}
	goto Rubystack /* stack new state and value */
}
