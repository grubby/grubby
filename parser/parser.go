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
	operator     string
	genericValue ast.Node
	genericSlice ast.Nodes
	stringSlice  []string
}

const OPERATOR = 57346
const NODE = 57347
const REF = 57348
const CAPITAL_REF = 57349
const LPAREN = 57350
const RPAREN = 57351
const COMMA = 57352
const DO = 57353
const DEF = 57354
const END = 57355
const IF = 57356
const ELSE = 57357
const ELSIF = 57358
const CLASS = 57359
const MODULE = 57360
const TRUE = 57361
const FALSE = 57362
const LESSTHAN = 57363
const GREATERTHAN = 57364
const EQUALTO = 57365
const BANG = 57366
const COMPLEMENT = 57367
const POSITIVE = 57368
const NEGATIVE = 57369
const STAR = 57370
const WHITESPACE = 57371
const NEWLINE = 57372
const SEMICOLON = 57373
const COLON = 57374
const DOT = 57375
const PIPE = 57376
const SLASH = 57377
const AMPERSAND = 57378
const MODULO = 57379
const CARET = 57380
const LBRACKET = 57381
const RBRACKET = 57382
const LBRACE = 57383
const RBRACE = 57384
const DOLLARSIGN = 57385
const ATSIGN = 57386
const FILE_CONST_REF = 57387
const EOF = 57388

var RubyToknames = []string{
	"OPERATOR",
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
	"SLASH",
	"AMPERSAND",
	"MODULO",
	"CARET",
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

//line parser.y:543

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 11,
	-1, 10,
	23, 11,
	-2, 16,
	-1, 32,
	1, 38,
	4, 38,
	10, 38,
	13, 38,
	15, 38,
	16, 38,
	28, 38,
	30, 38,
	34, 38,
	35, 38,
	36, 38,
	42, 38,
	46, 38,
	-2, 11,
	-1, 80,
	5, 11,
	6, 11,
	8, 11,
	-2, 45,
	-1, 83,
	10, 11,
	-2, 47,
	-1, 86,
	10, 11,
	-2, 50,
	-1, 87,
	4, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	-2, 51,
	-1, 89,
	1, 38,
	4, 38,
	10, 38,
	13, 38,
	15, 38,
	16, 38,
	28, 38,
	30, 38,
	34, 38,
	35, 38,
	36, 38,
	42, 38,
	46, 38,
	-2, 11,
	-1, 91,
	34, 11,
	-2, 7,
	-1, 97,
	4, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	-2, 80,
	-1, 98,
	4, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	-2, 81,
	-1, 99,
	4, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	-2, 82,
	-1, 100,
	4, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	-2, 83,
	-1, 105,
	4, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	-2, 7,
	-1, 128,
	32, 76,
	-2, 11,
	-1, 135,
	4, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	-2, 48,
	-1, 136,
	4, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	-2, 84,
	-1, 137,
	4, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	-2, 85,
	-1, 138,
	4, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	-2, 86,
	-1, 139,
	4, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	-2, 87,
	-1, 140,
	4, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	-2, 88,
	-1, 141,
	4, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	-2, 89,
	-1, 145,
	4, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	-2, 79,
	-1, 147,
	4, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	-2, 78,
	-1, 152,
	4, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	-2, 10,
	-1, 155,
	34, 114,
	-2, 60,
	-1, 169,
	4, 38,
	28, 38,
	34, 38,
	35, 38,
	36, 38,
	-2, 11,
	-1, 179,
	34, 115,
	-2, 61,
	-1, 196,
	4, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	-2, 53,
	-1, 210,
	32, 77,
	-2, 11,
	-1, 221,
	4, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	-2, 7,
	-1, 233,
	10, 99,
	30, 99,
	42, 99,
	-2, 11,
	-1, 234,
	4, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	-2, 7,
	-1, 245,
	34, 116,
	-2, 64,
	-1, 248,
	10, 96,
	30, 96,
	42, 96,
	-2, 11,
	-1, 254,
	30, 100,
	42, 100,
	-2, 11,
	-1, 257,
	30, 97,
	42, 97,
	-2, 11,
}

const RubyNprod = 126
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 749

var RubyAct = []int{

	121, 181, 5, 88, 101, 94, 82, 52, 7, 53,
	2, 55, 241, 68, 69, 239, 204, 6, 155, 164,
	190, 57, 6, 54, 104, 238, 57, 250, 184, 162,
	6, 132, 73, 56, 223, 122, 203, 6, 58, 59,
	60, 206, 207, 61, 62, 63, 64, 65, 66, 165,
	67, 70, 206, 207, 74, 75, 76, 183, 127, 130,
	95, 95, 79, 77, 78, 81, 96, 126, 6, 4,
	253, 232, 258, 224, 108, 109, 110, 111, 112, 113,
	114, 115, 116, 216, 117, 118, 119, 120, 73, 152,
	56, 6, 123, 124, 125, 8, 256, 214, 191, 149,
	193, 192, 131, 199, 200, 133, 146, 251, 210, 91,
	74, 75, 76, 198, 164, 6, 106, 107, 79, 77,
	78, 128, 142, 247, 148, 71, 72, 93, 95, 161,
	92, 156, 158, 160, 144, 143, 80, 85, 211, 245,
	202, 134, 211, 212, 103, 179, 87, 174, 129, 175,
	176, 97, 98, 99, 100, 177, 83, 105, 103, 102,
	180, 86, 159, 167, 73, 186, 163, 187, 188, 189,
	56, 178, 166, 194, 1, 157, 171, 154, 197, 27,
	201, 205, 26, 217, 25, 24, 74, 75, 76, 23,
	22, 218, 21, 219, 79, 77, 78, 20, 135, 136,
	137, 138, 139, 140, 141, 19, 145, 18, 147, 195,
	225, 227, 213, 215, 228, 30, 230, 208, 14, 209,
	34, 13, 12, 168, 35, 237, 15, 10, 31, 220,
	240, 11, 242, 243, 36, 28, 17, 29, 222, 229,
	16, 231, 3, 0, 0, 0, 0, 249, 0, 0,
	0, 252, 0, 0, 255, 0, 0, 235, 0, 0,
	0, 0, 0, 0, 196, 0, 9, 32, 33, 0,
	244, 0, 0, 37, 170, 49, 173, 172, 38, 39,
	40, 41, 0, 0, 221, 42, 43, 44, 45, 0,
	151, 150, 0, 0, 0, 0, 0, 0, 0, 0,
	46, 0, 47, 0, 51, 50, 48, 0, 233, 234,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	246, 0, 0, 248, 0, 0, 0, 0, 0, 9,
	32, 33, 254, 0, 0, 257, 37, 236, 49, 0,
	0, 38, 39, 40, 41, 0, 0, 0, 42, 43,
	44, 45, 0, 151, 150, 0, 0, 0, 0, 0,
	0, 0, 0, 46, 0, 47, 0, 51, 50, 48,
	9, 32, 33, 0, 0, 0, 0, 37, 226, 49,
	0, 0, 38, 39, 40, 41, 0, 0, 0, 42,
	43, 44, 45, 0, 151, 150, 0, 0, 0, 0,
	0, 0, 0, 0, 46, 0, 47, 0, 51, 50,
	48, 9, 32, 33, 0, 0, 0, 0, 37, 185,
	49, 0, 0, 38, 39, 40, 41, 0, 0, 0,
	42, 43, 44, 45, 0, 151, 150, 0, 0, 0,
	0, 0, 0, 0, 0, 46, 0, 47, 0, 51,
	50, 48, 9, 32, 33, 0, 0, 0, 0, 37,
	182, 49, 0, 0, 38, 39, 40, 41, 0, 0,
	0, 42, 43, 44, 45, 0, 151, 150, 0, 0,
	0, 0, 0, 0, 0, 0, 46, 0, 47, 0,
	51, 50, 48, 9, 32, 33, 0, 0, 0, 0,
	37, 153, 49, 0, 0, 38, 39, 40, 41, 0,
	0, 0, 42, 43, 44, 45, 0, 151, 150, 0,
	0, 0, 0, 0, 0, 0, 0, 46, 0, 47,
	0, 51, 50, 48, 90, 89, 33, 85, 0, 0,
	91, 37, 0, 49, 0, 0, 38, 39, 40, 41,
	0, 0, 84, 42, 43, 44, 45, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 46, 0,
	47, 0, 51, 50, 48, 9, 32, 33, 0, 0,
	0, 0, 37, 0, 49, 0, 0, 38, 39, 40,
	41, 0, 0, 0, 42, 43, 44, 45, 0, 151,
	150, 0, 0, 0, 0, 0, 0, 0, 0, 46,
	0, 47, 0, 51, 50, 48, 9, 32, 33, 0,
	0, 0, 91, 37, 0, 49, 0, 0, 38, 39,
	40, 41, 0, 0, 0, 42, 43, 44, 45, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	46, 0, 47, 0, 51, 50, 48, 9, 32, 33,
	0, 0, 0, 0, 37, 0, 49, 0, 0, 38,
	39, 40, 41, 0, 0, 0, 42, 43, 44, 45,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 46, 0, 47, 0, 51, 50, 48, 9, 169,
	33, 0, 0, 0, 0, 37, 0, 49, 0, 0,
	38, 39, 40, 41, 0, 0, 0, 42, 43, 44,
	45, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 46, 0, 47, 0, 51, 50, 48,
}
var RubyPact = []int{

	-36, 39, -1000, -38, -1000, 662, 8, -1000, 8, -10,
	8, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -7, -1000, -1000, -1000, -1000, 8, 8, 8,
	-1000, -1000, 8, 8, 8, 8, 8, 8, -1000, 8,
	7, 119, -1000, 28, 130, 42, 539, 124, 121, 8,
	8, 662, 662, 662, 662, 153, -1000, 662, -1000, -1000,
	110, -1000, -1000, 8, 8, 8, 8, 8, 8, 8,
	8, 8, -1000, 8, 8, 8, 8, 8, -1000, -7,
	-10, 8, 8, 8, 37, 114, 29, 8, 8, 8,
	8, 8, -1000, -1000, 1, 8, -1000, -1000, 662, 662,
	662, 662, 662, 662, 662, 129, 662, 96, 662, 153,
	89, 28, 498, -16, 129, 129, -1000, 8, 8, -3,
	-1000, 9, -1000, 703, 261, 8, 8, 8, 8, 8,
	8, 8, -1000, -1000, -1000, 8, 8, 8, 8, 8,
	-1000, -1000, 8, -1000, 8, 139, -1000, -1000, -1000, 457,
	27, -1000, -4, 416, 8, -1000, 8, 8, 8, -12,
	-1000, 85, 8, -1000, 621, 104, 98, 6, -18, -1000,
	8, 23, -1000, -1000, 101, -1000, 137, 87, 73, 160,
	8, -1000, 8, -1000, 662, 580, 8, -1000, -1000, -1000,
	-1000, -1000, 4, -1000, -1000, 63, -1000, -1000, 580, 375,
	8, -1000, -1000, 1, -1000, 1, -1000, 49, 662, 662,
	580, 8, 334, -1000, 8, 12, -1000, -1000, -27, 1,
	-30, 1, 8, 8, 8, 580, -1000, 133, -1000, -1000,
	662, -1000, 117, 662, 580, -1000, 8, -5, 8, 84,
	8, 48, 662, 8, 86, 662, -1000, 62, -1000,
}
var RubyPgo = []int{

	0, 243, 89, 242, 240, 237, 3, 236, 235, 234,
	231, 228, 227, 226, 224, 222, 221, 220, 218, 5,
	215, 207, 205, 197, 192, 190, 189, 185, 184, 182,
	179, 35, 6, 177, 176, 175, 174, 172, 4, 171,
	163, 161, 156, 148, 0, 24, 1, 140,
}
var RubyR1 = []int{

	0, 36, 36, 36, 36, 3, 3, 31, 31, 31,
	31, 44, 44, 45, 45, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 12, 12,
	12, 12, 12, 10, 10, 10, 10, 10, 10, 32,
	32, 42, 42, 42, 42, 41, 41, 41, 41, 41,
	38, 38, 38, 38, 38, 46, 46, 46, 15, 35,
	35, 16, 16, 18, 19, 19, 43, 43, 13, 13,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	4, 7, 8, 5, 5, 37, 37, 37, 37, 40,
	40, 40, 1, 1, 9, 9, 17, 17, 14, 14,
	20, 6, 6, 33, 39, 39, 39, 47, 47, 11,
	11, 34, 34, 34, 34, 34,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 3, 0, 2, 2,
	2, 0, 2, 0, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 5, 3, 5, 3, 5, 5,
	1, 1, 1, 5, 5, 1, 1, 5, 5, 5,
	0, 1, 1, 5, 5, 0, 2, 2, 9, 0,
	1, 6, 8, 6, 3, 6, 1, 4, 5, 5,
	3, 3, 3, 3, 5, 5, 5, 5, 5, 5,
	1, 1, 5, 9, 9, 0, 6, 11, 12, 4,
	9, 10, 0, 1, 2, 2, 2, 2, 3, 3,
	1, 3, 7, 3, 0, 1, 5, 1, 2, 5,
	6, 0, 5, 3, 4, 2,
}
var RubyChk = []int{

	-1000, -36, 46, -3, 30, -44, 29, 46, -2, 5,
	-12, -10, -15, -16, -18, -13, -4, -7, -21, -22,
	-23, -24, -25, -26, -27, -28, -29, -30, -8, -5,
	-20, -11, 6, 7, -17, -14, -9, 12, 17, 18,
	19, 20, 24, 25, 26, 27, 39, 41, 45, 14,
	44, 43, -44, -44, 33, -44, -44, 33, -44, -44,
	-44, -44, -44, -44, -44, -44, -44, -44, 6, 7,
	44, 6, 7, 4, 26, 27, 28, 35, 36, 34,
	6, 23, -32, -42, 23, 8, -41, -2, -6, 6,
	5, 11, 6, 6, -19, -44, -19, -2, -2, -2,
	-2, -38, 6, 5, -45, -2, 6, 7, -44, -44,
	-44, -44, -44, -44, -44, -44, -44, -44, -44, -44,
	-44, -44, -31, -44, -44, -44, 30, 21, 7, -43,
	30, -44, 30, -44, -31, -2, -2, -2, -2, -2,
	-2, -2, -32, 6, 5, -2, 10, -2, -38, 10,
	30, 29, -2, 13, -33, 34, -32, -35, -32, -31,
	-19, -44, 32, -31, 10, 40, -37, -40, -2, 6,
	13, -34, 16, 15, -44, -44, -44, -44, -39, 6,
	-38, -46, 13, 30, 32, 13, -44, -44, -44, -44,
	32, 13, 16, 15, -44, -31, -2, -6, 9, 5,
	6, -6, -47, 30, 34, -44, 29, 30, -31, -31,
	7, 5, 6, -45, 10, -45, 10, 23, -44, -44,
	-31, -2, -31, 30, 10, -46, 13, -44, -44, -45,
	-44, -45, 22, -2, -2, -31, 13, -44, 13, 42,
	-44, 42, -44, -44, -31, 6, -2, 6, -2, -44,
	32, 23, -44, 22, -2, -44, 10, -2, 10,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 0, 11, 4, 11, 15,
	-2, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, -2, 39, 40, 41, 42, 11, 11, 11,
	90, 91, 11, 11, 11, 11, 11, 11, 110, 11,
	0, 0, 12, 6, 0, 0, 0, 0, 0, 11,
	11, 0, 0, 0, 0, 60, 13, 0, 106, 107,
	0, 104, 105, 11, 11, 11, 11, 11, 11, 11,
	-2, 11, 43, -2, 11, 11, -2, -2, 52, -2,
	15, -2, 11, 11, 0, 0, 0, -2, -2, -2,
	-2, 11, 61, 62, 11, -2, 108, 109, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 60,
	0, 0, 0, 0, 0, 69, 7, 11, -2, 0,
	7, 0, 14, 95, 0, -2, -2, -2, -2, -2,
	-2, -2, 44, 55, 56, -2, 11, -2, 11, 11,
	8, 9, -2, 111, 11, -2, 46, 65, 70, 0,
	0, 74, 0, 0, 11, 92, 11, 11, 11, -2,
	119, 0, 11, 7, 0, 0, 0, 0, 0, -2,
	11, 7, 71, 7, 0, 73, 0, 13, 13, 0,
	11, 120, 11, 7, 0, 125, -2, 54, 49, 57,
	58, 59, 7, 117, 113, 0, 66, 67, 65, 0,
	-2, 63, 64, 11, 13, 11, 13, 0, 0, 0,
	123, -2, 0, 118, 11, 0, 72, 75, 0, 11,
	0, 11, 11, -2, -2, 124, 112, 0, 68, 93,
	0, 94, 0, 0, 122, -2, 11, 0, -2, 0,
	11, 0, 0, 11, -2, 0, 101, -2, 98,
}
var RubyTok1 = []int{

	1,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46,
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
		//line parser.y:139
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:141
		{
			Statements = []ast.Node{}
		}
	case 3:
		//line parser.y:143
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 4:
		//line parser.y:149
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 5:
		//line parser.y:155
		{
			RubyVAL.genericValue = nil
		}
	case 6:
		//line parser.y:156
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 7:
		//line parser.y:159
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 8:
		//line parser.y:161
		{ /* do nothing */
		}
	case 9:
		//line parser.y:163
		{ /* do nothing */
		}
	case 10:
		//line parser.y:165
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 41:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 42:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 43:
		//line parser.y:179
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 44:
		//line parser.y:186
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 45:
		//line parser.y:194
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 46:
		//line parser.y:201
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 47:
		//line parser.y:209
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 48:
		//line parser.y:216
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 49:
		//line parser.y:225
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 50:
		//line parser.y:227
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 51:
		//line parser.y:230
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 52:
		//line parser.y:232
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 53:
		//line parser.y:234
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 54:
		//line parser.y:236
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 55:
		//line parser.y:239
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 56:
		//line parser.y:241
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 57:
		//line parser.y:243
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 58:
		//line parser.y:245
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 59:
		//line parser.y:247
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 60:
		//line parser.y:249
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 61:
		//line parser.y:251
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 62:
		//line parser.y:253
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 63:
		//line parser.y:255
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 64:
		//line parser.y:257
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:267
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 69:
		//line parser.y:275
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 70:
		//line parser.y:276
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 71:
		//line parser.y:279
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 72:
		//line parser.y:286
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 73:
		//line parser.y:296
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 74:
		//line parser.y:305
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 75:
		//line parser.y:311
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 76:
		//line parser.y:319
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 77:
		//line parser.y:323
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 78:
		//line parser.y:328
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 79:
		//line parser.y:335
		{ // FIXME: this is a hack. If this rule is not present, "foo = 5" will not parse
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 80:
		//line parser.y:342
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 81:
		//line parser.y:343
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 82:
		//line parser.y:344
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 83:
		//line parser.y:345
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 84:
		//line parser.y:348
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 85:
		//line parser.y:357
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 86:
		//line parser.y:366
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 87:
		//line parser.y:375
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 88:
		//line parser.y:384
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 89:
		//line parser.y:393
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 90:
		//line parser.y:401
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 91:
		//line parser.y:402
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 92:
		//line parser.y:404
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 93:
		//line parser.y:408
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 94:
		//line parser.y:416
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 95:
		//line parser.y:424
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 96:
		//line parser.y:426
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 97:
		//line parser.y:428
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 98:
		//line parser.y:430
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 99:
		//line parser.y:433
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 100:
		//line parser.y:440
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 101:
		//line parser.y:447
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 102:
		//line parser.y:455
		{
			RubyVAL.genericValue = nil
		}
	case 103:
		//line parser.y:456
		{
			RubyVAL.genericValue = nil
		}
	case 104:
		//line parser.y:459
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 105:
		//line parser.y:461
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 106:
		//line parser.y:464
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 107:
		//line parser.y:466
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 108:
		//line parser.y:469
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 109:
		//line parser.y:471
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 110:
		//line parser.y:474
		{
			RubyVAL.genericValue = ast.FileNameConstReference{}
		}
	case 111:
		//line parser.y:477
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 112:
		//line parser.y:479
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 113:
		//line parser.y:487
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 114:
		//line parser.y:489
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 115:
		//line parser.y:491
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 116:
		//line parser.y:493
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 119:
		//line parser.y:498
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 120:
		//line parser.y:505
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 121:
		//line parser.y:513
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 122:
		//line parser.y:515
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 123:
		//line parser.y:522
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 124:
		//line parser.y:529
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 125:
		//line parser.y:536
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	}
	goto Rubystack /* stack new state and value */
}
