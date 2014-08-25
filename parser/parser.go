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

//line parser.y:568

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 11,
	-1, 10,
	23, 11,
	-2, 16,
	-1, 36,
	1, 42,
	4, 42,
	10, 42,
	13, 42,
	15, 42,
	16, 42,
	21, 42,
	22, 42,
	28, 42,
	30, 42,
	34, 42,
	35, 42,
	36, 42,
	37, 42,
	38, 42,
	42, 42,
	46, 42,
	-2, 11,
	-1, 89,
	10, 11,
	-2, 49,
	-1, 92,
	10, 11,
	-2, 52,
	-1, 93,
	4, 11,
	21, 11,
	22, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 53,
	-1, 95,
	1, 42,
	4, 42,
	10, 42,
	13, 42,
	15, 42,
	16, 42,
	21, 42,
	22, 42,
	28, 42,
	30, 42,
	34, 42,
	35, 42,
	36, 42,
	37, 42,
	38, 42,
	42, 42,
	46, 42,
	-2, 11,
	-1, 97,
	34, 11,
	-2, 7,
	-1, 103,
	4, 11,
	21, 11,
	22, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 82,
	-1, 104,
	4, 11,
	21, 11,
	22, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 83,
	-1, 105,
	4, 11,
	21, 11,
	22, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 84,
	-1, 106,
	4, 11,
	21, 11,
	22, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 85,
	-1, 111,
	4, 11,
	21, 11,
	22, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 7,
	-1, 137,
	32, 78,
	-2, 11,
	-1, 144,
	4, 11,
	21, 11,
	22, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 50,
	-1, 145,
	4, 11,
	21, 11,
	22, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 86,
	-1, 146,
	4, 11,
	21, 11,
	22, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 87,
	-1, 147,
	4, 11,
	21, 11,
	22, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 88,
	-1, 148,
	4, 11,
	21, 11,
	22, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 89,
	-1, 149,
	4, 11,
	21, 11,
	22, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 90,
	-1, 152,
	4, 11,
	21, 11,
	22, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 93,
	-1, 153,
	4, 11,
	21, 11,
	22, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 94,
	-1, 154,
	4, 11,
	21, 11,
	22, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 95,
	-1, 155,
	4, 11,
	21, 11,
	22, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 81,
	-1, 157,
	4, 11,
	21, 11,
	22, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 80,
	-1, 162,
	4, 11,
	21, 11,
	22, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 10,
	-1, 165,
	34, 120,
	-2, 62,
	-1, 181,
	4, 42,
	21, 42,
	22, 42,
	28, 42,
	34, 42,
	35, 42,
	36, 42,
	37, 42,
	38, 42,
	-2, 11,
	-1, 186,
	4, 11,
	21, 11,
	22, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 91,
	-1, 187,
	4, 11,
	21, 11,
	22, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 92,
	-1, 193,
	34, 121,
	-2, 63,
	-1, 210,
	4, 11,
	21, 11,
	22, 11,
	26, 11,
	27, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 55,
	-1, 224,
	32, 79,
	-2, 11,
	-1, 235,
	4, 11,
	21, 11,
	22, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 7,
	-1, 247,
	10, 105,
	30, 105,
	42, 105,
	-2, 11,
	-1, 248,
	4, 11,
	21, 11,
	22, 11,
	28, 11,
	34, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	-2, 7,
	-1, 259,
	34, 122,
	-2, 66,
	-1, 262,
	10, 102,
	30, 102,
	42, 102,
	-2, 11,
	-1, 268,
	30, 106,
	42, 106,
	-2, 11,
	-1, 271,
	30, 103,
	42, 103,
	-2, 11,
}

const RubyNprod = 132
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 788

var RubyAct = []int{

	130, 195, 5, 94, 107, 100, 110, 56, 7, 57,
	2, 58, 88, 255, 253, 71, 72, 6, 218, 165,
	204, 60, 6, 264, 252, 176, 60, 198, 174, 6,
	141, 220, 221, 131, 6, 4, 237, 59, 272, 270,
	220, 221, 61, 62, 63, 87, 217, 64, 65, 66,
	67, 68, 69, 73, 70, 177, 136, 6, 6, 267,
	197, 139, 6, 101, 101, 135, 246, 121, 120, 102,
	205, 238, 207, 206, 212, 176, 230, 114, 115, 116,
	117, 118, 119, 213, 214, 122, 123, 124, 125, 97,
	126, 127, 128, 129, 228, 76, 59, 159, 132, 133,
	134, 168, 167, 156, 91, 112, 113, 224, 140, 74,
	75, 142, 82, 83, 265, 225, 259, 77, 78, 79,
	137, 150, 151, 225, 226, 85, 80, 84, 81, 86,
	109, 193, 261, 158, 109, 108, 99, 101, 173, 98,
	216, 138, 172, 89, 92, 143, 166, 170, 179, 192,
	178, 1, 169, 183, 164, 162, 31, 188, 30, 189,
	190, 8, 29, 28, 27, 191, 76, 26, 25, 171,
	194, 24, 23, 175, 22, 21, 20, 200, 19, 201,
	202, 203, 59, 82, 83, 208, 18, 34, 77, 78,
	79, 14, 211, 38, 215, 219, 85, 80, 84, 81,
	86, 13, 12, 39, 15, 232, 10, 233, 227, 229,
	35, 11, 40, 32, 17, 93, 33, 16, 3, 209,
	103, 104, 105, 106, 239, 241, 111, 0, 242, 222,
	244, 223, 0, 0, 0, 243, 0, 245, 0, 251,
	0, 234, 0, 0, 254, 0, 256, 257, 0, 0,
	236, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 263, 0, 0, 0, 266, 0, 0, 269, 249,
	144, 145, 146, 147, 148, 149, 0, 0, 152, 153,
	154, 155, 258, 157, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 180, 0,
	0, 9, 36, 37, 0, 0, 186, 187, 41, 182,
	53, 185, 184, 42, 43, 44, 45, 0, 0, 0,
	46, 47, 48, 49, 0, 161, 160, 0, 0, 0,
	0, 0, 0, 0, 0, 50, 0, 51, 0, 55,
	54, 52, 0, 0, 210, 0, 9, 36, 37, 0,
	0, 0, 0, 41, 250, 53, 0, 0, 42, 43,
	44, 45, 0, 0, 235, 46, 47, 48, 49, 0,
	161, 160, 0, 76, 0, 0, 0, 0, 0, 0,
	50, 0, 51, 0, 55, 54, 52, 0, 247, 248,
	82, 83, 231, 0, 0, 77, 78, 79, 0, 0,
	0, 0, 0, 85, 80, 84, 81, 86, 0, 0,
	260, 0, 0, 262, 0, 0, 0, 0, 0, 9,
	36, 37, 268, 0, 0, 271, 41, 240, 53, 0,
	0, 42, 43, 44, 45, 0, 0, 0, 46, 47,
	48, 49, 0, 161, 160, 0, 0, 0, 0, 0,
	0, 0, 0, 50, 0, 51, 0, 55, 54, 52,
	9, 36, 37, 0, 0, 0, 0, 41, 199, 53,
	0, 0, 42, 43, 44, 45, 0, 0, 0, 46,
	47, 48, 49, 0, 161, 160, 0, 0, 0, 0,
	0, 0, 0, 0, 50, 0, 51, 0, 55, 54,
	52, 9, 36, 37, 0, 0, 0, 0, 41, 196,
	53, 0, 0, 42, 43, 44, 45, 0, 0, 0,
	46, 47, 48, 49, 0, 161, 160, 0, 0, 0,
	0, 0, 0, 0, 0, 50, 0, 51, 0, 55,
	54, 52, 9, 36, 37, 0, 0, 0, 0, 41,
	163, 53, 0, 0, 42, 43, 44, 45, 0, 0,
	0, 46, 47, 48, 49, 0, 161, 160, 0, 0,
	0, 0, 0, 0, 0, 0, 50, 0, 51, 0,
	55, 54, 52, 96, 95, 37, 91, 0, 0, 97,
	41, 0, 53, 0, 0, 42, 43, 44, 45, 0,
	0, 90, 46, 47, 48, 49, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 50, 0, 51,
	0, 55, 54, 52, 9, 36, 37, 0, 0, 0,
	0, 41, 0, 53, 0, 0, 42, 43, 44, 45,
	0, 0, 0, 46, 47, 48, 49, 0, 161, 160,
	0, 0, 0, 0, 0, 0, 0, 0, 50, 0,
	51, 0, 55, 54, 52, 9, 36, 37, 0, 0,
	0, 97, 41, 0, 53, 0, 0, 42, 43, 44,
	45, 0, 0, 0, 46, 47, 48, 49, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 50,
	0, 51, 0, 55, 54, 52, 9, 36, 37, 0,
	0, 0, 0, 41, 0, 53, 0, 0, 42, 43,
	44, 45, 0, 0, 0, 46, 47, 48, 49, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	50, 0, 51, 0, 55, 54, 52, 9, 181, 37,
	0, 0, 0, 0, 41, 0, 53, 0, 0, 42,
	43, 44, 45, 0, 0, 0, 46, 47, 48, 49,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 50, 0, 51, 0, 55, 54, 52,
}
var RubyPact = []int{

	-36, 5, -1000, -38, -1000, 701, 33, -1000, 33, -1000,
	33, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -7, -1000, -1000, -1000,
	-1000, 33, 33, 33, -1000, -1000, 33, 33, 33, 33,
	33, 33, -1000, 33, 9, 103, -1000, 162, 22, 578,
	133, 130, 33, 33, 701, 701, 701, 701, 129, -1000,
	701, -1000, -1000, 99, -1000, -1000, 33, 33, 33, 33,
	33, 33, 47, 45, 33, 33, 33, 33, -1000, 33,
	33, 33, 33, 33, -1000, -7, -1000, 33, 33, 33,
	35, 113, 31, 33, 33, 33, 33, 33, -1000, -1000,
	0, 33, -1000, -1000, 701, 701, 701, 701, 701, 701,
	33, 33, 701, 701, 701, 701, 93, 701, 129, 87,
	162, 537, -15, 96, 96, -1000, 33, 33, -4, -1000,
	15, -1000, 742, 296, 33, 33, 33, 33, 33, 33,
	701, 701, 33, 33, 33, 33, 33, 33, 33, 33,
	-1000, -1000, 33, -1000, 33, 125, -1000, -1000, -1000, -1000,
	-1000, 496, 30, -1000, -5, 455, 33, -1000, 33, 33,
	33, -12, -1000, 57, 33, -1000, 33, 33, 660, 65,
	78, 16, -16, -1000, 33, 2, -1000, -1000, 100, -1000,
	118, 84, 66, 369, 33, -1000, 33, -1000, 701, 619,
	33, -1000, -1000, -1000, -1000, -1000, 6, -1000, -1000, 61,
	-1000, -1000, 619, 414, 33, -1000, -1000, 0, -1000, 0,
	-1000, 44, 701, 701, 619, 33, 341, -1000, 33, 11,
	-1000, -1000, -28, 0, -29, 0, 33, 33, 33, 619,
	-1000, 110, -1000, -1000, 701, -1000, 126, 701, 619, -1000,
	33, -9, 33, 91, 33, 37, 701, 33, 29, 701,
	-1000, 28, -1000,
}
var RubyPgo = []int{

	0, 227, 155, 218, 217, 216, 3, 214, 213, 212,
	211, 210, 206, 204, 203, 202, 201, 193, 191, 5,
	187, 186, 178, 176, 175, 174, 172, 171, 168, 167,
	164, 163, 162, 158, 156, 33, 12, 154, 153, 152,
	151, 150, 4, 149, 148, 144, 143, 141, 0, 6,
	1, 140,
}
var RubyR1 = []int{

	0, 40, 40, 40, 40, 3, 3, 35, 35, 35,
	35, 48, 48, 49, 49, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 12, 12, 12, 12, 12, 10, 10, 10,
	10, 36, 36, 46, 46, 46, 46, 45, 45, 45,
	45, 45, 42, 42, 42, 42, 42, 50, 50, 50,
	15, 39, 39, 16, 16, 18, 19, 19, 47, 47,
	13, 13, 21, 22, 23, 24, 25, 26, 27, 28,
	29, 30, 31, 32, 33, 34, 4, 7, 8, 5,
	5, 41, 41, 41, 41, 44, 44, 44, 1, 1,
	9, 9, 17, 17, 14, 14, 20, 6, 6, 37,
	43, 43, 43, 51, 51, 11, 11, 38, 38, 38,
	38, 38,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 3, 0, 2, 2,
	2, 0, 2, 0, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 5, 3,
	5, 5, 1, 1, 1, 5, 5, 1, 1, 5,
	5, 5, 0, 1, 1, 5, 5, 0, 2, 2,
	9, 0, 1, 6, 8, 6, 3, 6, 1, 4,
	5, 5, 3, 3, 3, 3, 5, 5, 5, 5,
	5, 6, 6, 5, 5, 5, 1, 1, 5, 9,
	9, 0, 6, 11, 12, 4, 9, 10, 0, 1,
	2, 2, 2, 2, 3, 3, 1, 3, 7, 3,
	0, 1, 5, 1, 2, 5, 6, 0, 5, 3,
	4, 2,
}
var RubyChk = []int{

	-1000, -40, 46, -3, 30, -48, 29, 46, -2, 5,
	-12, -10, -15, -16, -18, -13, -4, -7, -21, -22,
	-23, -24, -25, -26, -27, -28, -29, -30, -31, -32,
	-33, -34, -8, -5, -20, -11, 6, 7, -17, -14,
	-9, 12, 17, 18, 19, 20, 24, 25, 26, 27,
	39, 41, 45, 14, 44, 43, -48, -48, -48, -48,
	33, -48, -48, -48, -48, -48, -48, -48, -48, -48,
	-48, 6, 7, 44, 6, 7, 4, 26, 27, 28,
	35, 37, 21, 22, 36, 34, 38, 23, -36, -46,
	23, 8, -45, -2, -6, 6, 5, 11, 6, 6,
	-19, -48, -19, -2, -2, -2, -2, -42, 6, 5,
	-49, -2, 6, 7, -48, -48, -48, -48, -48, -48,
	21, 22, -48, -48, -48, -48, -48, -48, -48, -48,
	-48, -35, -48, -48, -48, 30, 21, 7, -47, 30,
	-48, 30, -48, -35, -2, -2, -2, -2, -2, -2,
	-48, -48, -2, -2, -2, -2, 10, -2, -42, 10,
	30, 29, -2, 13, -37, 34, -36, 6, 5, -39,
	-36, -35, -19, -48, 32, -35, 10, 40, -41, -44,
	-2, 6, 13, -38, 16, 15, -2, -2, -48, -48,
	-48, -48, -43, 6, -42, -50, 13, 30, 32, 13,
	-48, -48, -48, -48, 32, 13, 16, 15, -48, -35,
	-2, -6, 9, 5, 6, -6, -51, 30, 34, -48,
	29, 30, -35, -35, 7, 5, 6, -49, 10, -49,
	10, 23, -48, -48, -35, -2, -35, 30, 10, -50,
	13, -48, -48, -49, -48, -49, 22, -2, -2, -35,
	13, -48, 13, 42, -48, 42, -48, -48, -35, 6,
	-2, 6, -2, -48, 32, 23, -48, 22, -2, -48,
	10, -2, 10,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 0, 11, 4, 11, 15,
	-2, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, -2, 43, 44, 45,
	46, 11, 11, 11, 96, 97, 11, 11, 11, 11,
	11, 11, 116, 11, 0, 0, 12, 6, 0, 0,
	0, 0, 11, 11, 0, 0, 0, 0, 62, 13,
	0, 112, 113, 0, 110, 111, 11, 11, 11, 11,
	11, 11, 0, 0, 11, 11, 11, 11, 47, -2,
	11, 11, -2, -2, 54, -2, 15, -2, 11, 11,
	0, 0, 0, -2, -2, -2, -2, 11, 63, 64,
	11, -2, 114, 115, 0, 0, 0, 0, 0, 0,
	11, 11, 0, 0, 0, 0, 0, 0, 62, 0,
	0, 0, 0, 0, 71, 7, 11, -2, 0, 7,
	0, 14, 101, 0, -2, -2, -2, -2, -2, -2,
	0, 0, -2, -2, -2, -2, 11, -2, 11, 11,
	8, 9, -2, 117, 11, -2, 48, 57, 58, 67,
	72, 0, 0, 76, 0, 0, 11, 98, 11, 11,
	11, -2, 125, 0, 11, 7, -2, -2, 0, 0,
	0, 0, 0, -2, 11, 7, 73, 7, 0, 75,
	0, 13, 13, 0, 11, 126, 11, 7, 0, 131,
	-2, 56, 51, 59, 60, 61, 7, 123, 119, 0,
	68, 69, 67, 0, -2, 65, 66, 11, 13, 11,
	13, 0, 0, 0, 129, -2, 0, 124, 11, 0,
	74, 77, 0, 11, 0, 11, 11, -2, -2, 130,
	118, 0, 70, 99, 0, 100, 0, 0, 128, -2,
	11, 0, -2, 0, 11, 0, 0, 11, -2, 0,
	107, -2, 104,
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
		//line parser.y:143
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:145
		{
			Statements = []ast.Node{}
		}
	case 3:
		//line parser.y:147
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 4:
		//line parser.y:153
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 5:
		//line parser.y:159
		{
			RubyVAL.genericValue = nil
		}
	case 6:
		//line parser.y:160
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 7:
		//line parser.y:163
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 8:
		//line parser.y:165
		{ /* do nothing */
		}
	case 9:
		//line parser.y:167
		{ /* do nothing */
		}
	case 10:
		//line parser.y:169
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
		//line parser.y:183
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 48:
		//line parser.y:190
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 49:
		//line parser.y:198
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 50:
		//line parser.y:205
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 51:
		//line parser.y:214
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 52:
		//line parser.y:216
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 53:
		//line parser.y:219
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 54:
		//line parser.y:221
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 55:
		//line parser.y:223
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 56:
		//line parser.y:225
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 57:
		//line parser.y:228
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 58:
		//line parser.y:230
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 59:
		//line parser.y:232
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 60:
		//line parser.y:234
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 61:
		//line parser.y:236
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 62:
		//line parser.y:238
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 63:
		//line parser.y:240
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 64:
		//line parser.y:242
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 65:
		//line parser.y:244
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 66:
		//line parser.y:246
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 70:
		//line parser.y:256
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 71:
		//line parser.y:264
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 72:
		//line parser.y:265
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 73:
		//line parser.y:268
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 74:
		//line parser.y:275
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 75:
		//line parser.y:285
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 76:
		//line parser.y:294
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 77:
		//line parser.y:300
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 78:
		//line parser.y:308
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 79:
		//line parser.y:312
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 80:
		//line parser.y:317
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 81:
		//line parser.y:324
		{ // FIXME: this is a hack. If this rule is not present, "foo = 5" will not parse
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 82:
		//line parser.y:331
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 83:
		//line parser.y:332
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 84:
		//line parser.y:333
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 85:
		//line parser.y:334
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 86:
		//line parser.y:337
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 87:
		//line parser.y:346
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 88:
		//line parser.y:355
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 89:
		//line parser.y:364
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 90:
		//line parser.y:373
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "%"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 91:
		//line parser.y:382
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-5].genericValue,
				Func:   ast.BareReference{Name: "<<"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 92:
		//line parser.y:391
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-5].genericValue,
				Func:   ast.BareReference{Name: ">>"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 93:
		//line parser.y:400
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 94:
		//line parser.y:409
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 95:
		//line parser.y:418
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "^"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 96:
		//line parser.y:426
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 97:
		//line parser.y:427
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 98:
		//line parser.y:429
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 99:
		//line parser.y:433
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 100:
		//line parser.y:441
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 101:
		//line parser.y:449
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 102:
		//line parser.y:451
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 103:
		//line parser.y:453
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 104:
		//line parser.y:455
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 105:
		//line parser.y:458
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 106:
		//line parser.y:465
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 107:
		//line parser.y:472
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 108:
		//line parser.y:480
		{
			RubyVAL.genericValue = nil
		}
	case 109:
		//line parser.y:481
		{
			RubyVAL.genericValue = nil
		}
	case 110:
		//line parser.y:484
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 111:
		//line parser.y:486
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 112:
		//line parser.y:489
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 113:
		//line parser.y:491
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 114:
		//line parser.y:494
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 115:
		//line parser.y:496
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 116:
		//line parser.y:499
		{
			RubyVAL.genericValue = ast.FileNameConstReference{}
		}
	case 117:
		//line parser.y:502
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 118:
		//line parser.y:504
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 119:
		//line parser.y:512
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 120:
		//line parser.y:514
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 121:
		//line parser.y:516
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 122:
		//line parser.y:518
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 125:
		//line parser.y:523
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 126:
		//line parser.y:530
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 127:
		//line parser.y:538
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 128:
		//line parser.y:540
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 129:
		//line parser.y:547
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 130:
		//line parser.y:554
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 131:
		//line parser.y:561
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	}
	goto Rubystack /* stack new state and value */
}
