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
const SLASH = 57376
const AMPERSAND = 57377
const MODULO = 57378
const CARET = 57379
const LBRACKET = 57380
const RBRACKET = 57381
const LBRACE = 57382
const RBRACE = 57383
const DOLLARSIGN = 57384
const ATSIGN = 57385
const FILE_CONST_REF = 57386
const EOF = 57387

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

//line parser.y:556

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 11,
	-1, 10,
	22, 11,
	-2, 16,
	-1, 36,
	1, 42,
	9, 42,
	12, 42,
	14, 42,
	15, 42,
	20, 42,
	21, 42,
	27, 42,
	29, 42,
	33, 42,
	34, 42,
	35, 42,
	36, 42,
	37, 42,
	41, 42,
	45, 42,
	-2, 11,
	-1, 88,
	9, 11,
	-2, 49,
	-1, 91,
	9, 11,
	-2, 51,
	-1, 92,
	20, 11,
	21, 11,
	25, 11,
	26, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 52,
	-1, 94,
	1, 42,
	9, 42,
	12, 42,
	14, 42,
	15, 42,
	20, 42,
	21, 42,
	27, 42,
	29, 42,
	33, 42,
	34, 42,
	35, 42,
	36, 42,
	37, 42,
	41, 42,
	45, 42,
	-2, 11,
	-1, 96,
	33, 11,
	-2, 7,
	-1, 102,
	20, 11,
	21, 11,
	25, 11,
	26, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 81,
	-1, 103,
	20, 11,
	21, 11,
	25, 11,
	26, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 82,
	-1, 104,
	20, 11,
	21, 11,
	25, 11,
	26, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 83,
	-1, 105,
	20, 11,
	21, 11,
	25, 11,
	26, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 84,
	-1, 110,
	20, 11,
	21, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 7,
	-1, 135,
	31, 77,
	-2, 11,
	-1, 142,
	20, 11,
	21, 11,
	25, 11,
	26, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 85,
	-1, 143,
	20, 11,
	21, 11,
	25, 11,
	26, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 86,
	-1, 144,
	20, 11,
	21, 11,
	25, 11,
	26, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 87,
	-1, 145,
	20, 11,
	21, 11,
	25, 11,
	26, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 88,
	-1, 146,
	20, 11,
	21, 11,
	25, 11,
	26, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 89,
	-1, 149,
	20, 11,
	21, 11,
	25, 11,
	26, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 92,
	-1, 150,
	20, 11,
	21, 11,
	25, 11,
	26, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 93,
	-1, 151,
	20, 11,
	21, 11,
	25, 11,
	26, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 94,
	-1, 152,
	20, 11,
	21, 11,
	25, 11,
	26, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 80,
	-1, 154,
	20, 11,
	21, 11,
	25, 11,
	26, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 79,
	-1, 159,
	20, 11,
	21, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 10,
	-1, 162,
	33, 119,
	-2, 61,
	-1, 178,
	20, 42,
	21, 42,
	27, 42,
	33, 42,
	34, 42,
	35, 42,
	36, 42,
	37, 42,
	-2, 11,
	-1, 183,
	20, 11,
	21, 11,
	25, 11,
	26, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 90,
	-1, 184,
	20, 11,
	21, 11,
	25, 11,
	26, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 91,
	-1, 190,
	33, 120,
	-2, 62,
	-1, 207,
	20, 11,
	21, 11,
	25, 11,
	26, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 54,
	-1, 221,
	31, 78,
	-2, 11,
	-1, 232,
	20, 11,
	21, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 7,
	-1, 244,
	9, 104,
	29, 104,
	41, 104,
	-2, 11,
	-1, 245,
	20, 11,
	21, 11,
	27, 11,
	33, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	-2, 7,
	-1, 256,
	33, 121,
	-2, 65,
	-1, 259,
	9, 101,
	29, 101,
	41, 101,
	-2, 11,
	-1, 265,
	29, 105,
	41, 105,
	-2, 11,
	-1, 268,
	29, 102,
	41, 102,
	-2, 11,
}

const RubyNprod = 131
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 741

var RubyAct = []int{

	128, 192, 5, 93, 106, 99, 87, 56, 7, 57,
	2, 58, 252, 250, 109, 173, 215, 162, 261, 9,
	36, 37, 71, 72, 195, 96, 41, 171, 53, 6,
	139, 42, 43, 44, 45, 217, 218, 59, 46, 47,
	48, 49, 61, 62, 63, 174, 6, 64, 65, 66,
	67, 68, 69, 50, 70, 51, 249, 55, 54, 52,
	73, 134, 6, 100, 100, 201, 60, 6, 118, 101,
	133, 60, 217, 218, 6, 4, 234, 113, 114, 115,
	116, 117, 269, 267, 120, 121, 122, 123, 214, 124,
	125, 126, 127, 194, 129, 59, 137, 130, 131, 132,
	86, 6, 6, 264, 243, 119, 202, 138, 204, 203,
	140, 81, 82, 262, 209, 173, 76, 77, 78, 147,
	148, 235, 227, 225, 84, 79, 83, 80, 85, 165,
	164, 155, 90, 156, 258, 100, 170, 153, 163, 167,
	169, 210, 211, 111, 112, 74, 75, 96, 222, 256,
	222, 223, 159, 221, 185, 135, 186, 187, 8, 108,
	190, 98, 188, 108, 107, 97, 213, 191, 136, 88,
	91, 176, 189, 175, 197, 1, 198, 199, 200, 59,
	166, 180, 205, 161, 31, 30, 29, 28, 27, 208,
	26, 212, 216, 81, 82, 228, 25, 24, 76, 77,
	78, 23, 229, 22, 230, 141, 84, 79, 83, 80,
	85, 21, 92, 224, 226, 20, 19, 102, 103, 104,
	105, 236, 238, 110, 18, 239, 34, 241, 168, 14,
	38, 13, 172, 12, 39, 15, 248, 10, 35, 11,
	240, 251, 242, 253, 254, 40, 32, 17, 33, 16,
	3, 0, 0, 0, 0, 0, 0, 0, 260, 0,
	0, 0, 263, 0, 0, 266, 142, 143, 144, 145,
	146, 0, 0, 149, 150, 151, 152, 206, 154, 0,
	0, 0, 0, 0, 0, 0, 0, 219, 0, 220,
	81, 82, 0, 177, 0, 76, 77, 78, 0, 231,
	183, 184, 0, 84, 79, 83, 80, 85, 233, 9,
	36, 37, 0, 0, 0, 0, 41, 179, 53, 182,
	181, 42, 43, 44, 45, 0, 0, 246, 46, 47,
	48, 49, 0, 158, 157, 0, 0, 0, 207, 0,
	255, 0, 0, 50, 0, 51, 0, 55, 54, 52,
	0, 0, 9, 36, 37, 0, 0, 0, 232, 41,
	247, 53, 0, 0, 42, 43, 44, 45, 0, 0,
	0, 46, 47, 48, 49, 0, 158, 157, 0, 0,
	0, 0, 244, 245, 0, 0, 50, 0, 51, 0,
	55, 54, 52, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 257, 0, 0, 259, 0, 0,
	0, 0, 0, 9, 36, 37, 265, 0, 0, 268,
	41, 237, 53, 0, 0, 42, 43, 44, 45, 0,
	0, 0, 46, 47, 48, 49, 0, 158, 157, 0,
	0, 0, 0, 0, 0, 0, 0, 50, 0, 51,
	0, 55, 54, 52, 9, 36, 37, 0, 0, 0,
	0, 41, 196, 53, 0, 0, 42, 43, 44, 45,
	0, 0, 0, 46, 47, 48, 49, 0, 158, 157,
	0, 0, 0, 0, 0, 0, 0, 0, 50, 0,
	51, 0, 55, 54, 52, 9, 36, 37, 0, 0,
	0, 0, 41, 193, 53, 0, 0, 42, 43, 44,
	45, 0, 0, 0, 46, 47, 48, 49, 0, 158,
	157, 0, 0, 0, 0, 0, 0, 0, 0, 50,
	0, 51, 0, 55, 54, 52, 9, 36, 37, 0,
	0, 0, 0, 41, 160, 53, 0, 0, 42, 43,
	44, 45, 0, 0, 0, 46, 47, 48, 49, 0,
	158, 157, 0, 0, 0, 0, 0, 0, 0, 0,
	50, 0, 51, 0, 55, 54, 52, 95, 94, 37,
	90, 0, 0, 96, 41, 0, 53, 0, 0, 42,
	43, 44, 45, 0, 0, 89, 46, 47, 48, 49,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 50, 0, 51, 0, 55, 54, 52, 9, 36,
	37, 0, 0, 0, 0, 41, 0, 53, 0, 0,
	42, 43, 44, 45, 0, 0, 0, 46, 47, 48,
	49, 0, 158, 157, 0, 0, 0, 0, 0, 0,
	0, 0, 50, 0, 51, 0, 55, 54, 52, 9,
	36, 37, 0, 0, 0, 0, 41, 0, 53, 0,
	0, 42, 43, 44, 45, 0, 0, 0, 46, 47,
	48, 49, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 50, 0, 51, 0, 55, 54, 52,
	9, 178, 37, 0, 0, 0, 0, 41, 0, 53,
	0, 0, 42, 43, 44, 45, 0, 0, 0, 46,
	47, 48, 49, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 50, 0, 51, 0, 55, 54,
	52,
}
var RubyPact = []int{

	-35, 46, -1000, -37, -1000, 655, 18, -1000, 18, -1000,
	18, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 39, -1000, -1000, -1000,
	-1000, 18, 18, 18, -1000, -1000, 18, 18, 18, 18,
	18, 18, -1000, 18, 17, 140, -1000, 270, 78, 573,
	160, 156, 18, 18, 655, 655, 655, 655, 159, -1000,
	655, -1000, -1000, 138, -1000, -1000, 18, 18, 18, 18,
	18, 48, 84, 18, 18, 18, 18, -1000, 18, 18,
	18, 18, 18, -1000, 39, -1000, 18, 18, 18, 41,
	149, 67, 18, 18, 18, 18, 18, -1000, -1000, 1,
	18, -1000, -1000, 655, 655, 655, 655, 655, 18, 18,
	655, 655, 655, 655, 128, 655, 159, 124, 270, 532,
	-16, 125, 125, -1000, 18, 18, -4, -1000, 6, -1000,
	696, 305, 18, 18, 18, 18, 18, 655, 655, 18,
	18, 18, 18, 18, 18, 18, 18, -1000, -1000, 18,
	-1000, 18, 155, -1000, -1000, -1000, -1000, -1000, 491, 64,
	-1000, -7, 450, 18, -1000, 18, 18, 18, 34, -1000,
	94, 18, -1000, 18, 18, 15, 106, 137, 59, -17,
	-1000, 18, 7, -1000, -1000, 147, -1000, 146, 114, 113,
	173, 18, -1000, 18, -1000, 655, 614, 18, -1000, -1000,
	-1000, -1000, -1000, 47, -1000, -1000, 112, -1000, -1000, 614,
	409, 18, -1000, -1000, 1, -1000, 1, -1000, 83, 655,
	655, 614, 18, 348, -1000, 18, 44, -1000, -1000, -28,
	1, -29, 1, 18, 18, 18, 614, -1000, 144, -1000,
	-1000, 655, -1000, 129, 655, 614, -1000, 18, -13, 18,
	91, 18, 82, 655, 18, 74, 655, -1000, 73, -1000,
}
var RubyPgo = []int{

	0, 251, 152, 250, 249, 248, 3, 247, 246, 245,
	239, 238, 237, 235, 234, 233, 231, 230, 229, 5,
	226, 224, 216, 215, 211, 203, 201, 197, 196, 190,
	188, 187, 186, 185, 184, 94, 6, 183, 181, 180,
	175, 173, 4, 172, 171, 170, 169, 168, 0, 14,
	1, 166,
}
var RubyR1 = []int{

	0, 40, 40, 40, 40, 3, 3, 35, 35, 35,
	35, 48, 48, 49, 49, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 12, 12, 12, 12, 12, 10, 10, 10,
	36, 36, 46, 46, 46, 46, 45, 45, 45, 45,
	45, 42, 42, 42, 42, 42, 50, 50, 50, 15,
	39, 39, 16, 16, 18, 19, 19, 47, 47, 13,
	13, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 4, 7, 8, 5, 5,
	41, 41, 41, 41, 44, 44, 44, 1, 1, 9,
	9, 17, 17, 14, 14, 20, 6, 6, 37, 43,
	43, 43, 51, 51, 11, 11, 38, 38, 38, 38,
	38,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 3, 0, 2, 2,
	2, 0, 2, 0, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 5, 3,
	5, 1, 1, 1, 5, 5, 1, 1, 5, 5,
	5, 0, 1, 1, 5, 5, 0, 2, 2, 9,
	0, 1, 6, 8, 6, 3, 6, 1, 4, 5,
	5, 3, 3, 3, 3, 5, 5, 5, 5, 5,
	6, 6, 5, 5, 5, 1, 1, 5, 9, 9,
	0, 6, 11, 12, 4, 9, 10, 0, 1, 2,
	2, 2, 2, 3, 3, 1, 3, 7, 3, 0,
	1, 5, 1, 2, 5, 6, 0, 5, 3, 4,
	2,
}
var RubyChk = []int{

	-1000, -40, 45, -3, 29, -48, 28, 45, -2, 4,
	-12, -10, -15, -16, -18, -13, -4, -7, -21, -22,
	-23, -24, -25, -26, -27, -28, -29, -30, -31, -32,
	-33, -34, -8, -5, -20, -11, 5, 6, -17, -14,
	-9, 11, 16, 17, 18, 19, 23, 24, 25, 26,
	38, 40, 44, 13, 43, 42, -48, -48, -48, -48,
	32, -48, -48, -48, -48, -48, -48, -48, -48, -48,
	-48, 5, 6, 43, 5, 6, 25, 26, 27, 34,
	36, 20, 21, 35, 33, 37, 22, -36, -46, 22,
	7, -45, -2, -6, 5, 4, 10, 5, 5, -19,
	-48, -19, -2, -2, -2, -2, -42, 5, 4, -49,
	-2, 5, 6, -48, -48, -48, -48, -48, 20, 21,
	-48, -48, -48, -48, -48, -48, -48, -48, -48, -35,
	-48, -48, -48, 29, 20, 6, -47, 29, -48, 29,
	-48, -35, -2, -2, -2, -2, -2, -48, -48, -2,
	-2, -2, -2, 9, -2, -42, 9, 29, 28, -2,
	12, -37, 33, -36, 5, 4, -39, -36, -35, -19,
	-48, 31, -35, 9, 39, -41, -44, -2, 5, 12,
	-38, 15, 14, -2, -2, -48, -48, -48, -48, -43,
	5, -42, -50, 12, 29, 31, 12, -48, -48, -48,
	-48, 31, 12, 15, 14, -48, -35, -2, -6, 8,
	4, 5, -6, -51, 29, 33, -48, 28, 29, -35,
	-35, 6, 4, 5, -49, 9, -49, 9, 22, -48,
	-48, -35, -2, -35, 29, 9, -50, 12, -48, -48,
	-49, -48, -49, 21, -2, -2, -35, 12, -48, 12,
	41, -48, 41, -48, -48, -35, 5, -2, 5, -2,
	-48, 31, 22, -48, 21, -2, -48, 9, -2, 9,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 0, 11, 4, 11, 15,
	-2, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, -2, 43, 44, 45,
	46, 11, 11, 11, 95, 96, 11, 11, 11, 11,
	11, 11, 115, 11, 0, 0, 12, 6, 0, 0,
	0, 0, 11, 11, 0, 0, 0, 0, 61, 13,
	0, 111, 112, 0, 109, 110, 11, 11, 11, 11,
	11, 0, 0, 11, 11, 11, 11, 47, -2, 11,
	11, -2, -2, 53, -2, 15, -2, 11, 11, 0,
	0, 0, -2, -2, -2, -2, 11, 62, 63, 11,
	-2, 113, 114, 0, 0, 0, 0, 0, 11, 11,
	0, 0, 0, 0, 0, 0, 61, 0, 0, 0,
	0, 0, 70, 7, 11, -2, 0, 7, 0, 14,
	100, 0, -2, -2, -2, -2, -2, 0, 0, -2,
	-2, -2, -2, 11, -2, 11, 11, 8, 9, -2,
	116, 11, -2, 48, 56, 57, 66, 71, 0, 0,
	75, 0, 0, 11, 97, 11, 11, 11, -2, 124,
	0, 11, 7, -2, -2, 0, 0, 0, 0, 0,
	-2, 11, 7, 72, 7, 0, 74, 0, 13, 13,
	0, 11, 125, 11, 7, 0, 130, -2, 55, 50,
	58, 59, 60, 7, 122, 118, 0, 67, 68, 66,
	0, -2, 64, 65, 11, 13, 11, 13, 0, 0,
	0, 128, -2, 0, 123, 11, 0, 73, 76, 0,
	11, 0, 11, 11, -2, -2, 129, 117, 0, 69,
	98, 0, 99, 0, 0, 127, -2, 11, 0, -2,
	0, 11, 0, 0, 11, -2, 0, 106, -2, 103,
}
var RubyTok1 = []int{

	1,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45,
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
		//line parser.y:140
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:142
		{
			Statements = []ast.Node{}
		}
	case 3:
		//line parser.y:144
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 4:
		//line parser.y:150
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 5:
		//line parser.y:156
		{
			RubyVAL.genericValue = nil
		}
	case 6:
		//line parser.y:157
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 7:
		//line parser.y:160
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 8:
		//line parser.y:162
		{ /* do nothing */
		}
	case 9:
		//line parser.y:164
		{ /* do nothing */
		}
	case 10:
		//line parser.y:166
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 44:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 45:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 46:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 47:
		//line parser.y:180
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 48:
		//line parser.y:187
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 49:
		//line parser.y:195
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 50:
		//line parser.y:203
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 51:
		//line parser.y:205
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 52:
		//line parser.y:208
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 53:
		//line parser.y:210
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 54:
		//line parser.y:212
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 55:
		//line parser.y:214
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 56:
		//line parser.y:217
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 57:
		//line parser.y:219
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 58:
		//line parser.y:221
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 59:
		//line parser.y:223
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 60:
		//line parser.y:225
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 61:
		//line parser.y:227
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 62:
		//line parser.y:229
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 63:
		//line parser.y:231
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 64:
		//line parser.y:233
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 65:
		//line parser.y:235
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 69:
		//line parser.y:245
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 70:
		//line parser.y:253
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 71:
		//line parser.y:254
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 72:
		//line parser.y:257
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 73:
		//line parser.y:264
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 74:
		//line parser.y:274
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 75:
		//line parser.y:283
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 76:
		//line parser.y:289
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 77:
		//line parser.y:297
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 78:
		//line parser.y:301
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 79:
		//line parser.y:306
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 80:
		//line parser.y:313
		{ // FIXME: this is a hack. If this rule is not present, "foo = 5" will not parse
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 81:
		//line parser.y:320
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 82:
		//line parser.y:321
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 83:
		//line parser.y:322
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 84:
		//line parser.y:323
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 85:
		//line parser.y:326
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 86:
		//line parser.y:335
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 87:
		//line parser.y:344
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 88:
		//line parser.y:353
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 89:
		//line parser.y:362
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "%"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 90:
		//line parser.y:371
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-5].genericValue,
				Func:   ast.BareReference{Name: "<<"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 91:
		//line parser.y:380
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-5].genericValue,
				Func:   ast.BareReference{Name: ">>"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 92:
		//line parser.y:389
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 93:
		//line parser.y:398
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 94:
		//line parser.y:407
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "^"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 95:
		//line parser.y:415
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 96:
		//line parser.y:416
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 97:
		//line parser.y:418
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 98:
		//line parser.y:422
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 99:
		//line parser.y:430
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 100:
		//line parser.y:438
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 101:
		//line parser.y:440
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 102:
		//line parser.y:442
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 103:
		//line parser.y:444
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 104:
		//line parser.y:447
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 105:
		//line parser.y:454
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 106:
		//line parser.y:461
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 107:
		//line parser.y:469
		{
			RubyVAL.genericValue = nil
		}
	case 108:
		//line parser.y:470
		{
			RubyVAL.genericValue = nil
		}
	case 109:
		//line parser.y:473
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 110:
		//line parser.y:475
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 111:
		//line parser.y:478
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 112:
		//line parser.y:480
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 113:
		//line parser.y:483
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 114:
		//line parser.y:485
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 115:
		//line parser.y:488
		{
			RubyVAL.genericValue = ast.FileNameConstReference{}
		}
	case 116:
		//line parser.y:491
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 117:
		//line parser.y:493
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 118:
		//line parser.y:501
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 119:
		//line parser.y:503
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 120:
		//line parser.y:505
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 121:
		//line parser.y:507
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 124:
		//line parser.y:512
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 125:
		//line parser.y:519
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 126:
		//line parser.y:527
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 127:
		//line parser.y:529
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 128:
		//line parser.y:536
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 129:
		//line parser.y:543
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 130:
		//line parser.y:550
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	}
	goto Rubystack /* stack new state and value */
}
