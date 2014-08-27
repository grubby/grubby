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

//line parser.y:583

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
	-1, 88,
	5, 11,
	6, 11,
	8, 11,
	-2, 49,
	-1, 91,
	10, 11,
	-2, 51,
	-1, 94,
	10, 11,
	-2, 54,
	-1, 95,
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
	-1, 97,
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
	-1, 99,
	34, 11,
	-2, 7,
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
	-1, 107,
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
	-1, 108,
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
	-1, 113,
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
	-1, 140,
	32, 80,
	-2, 11,
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
	-2, 52,
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
	-2, 88,
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
	-2, 89,
	-1, 150,
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
	-1, 151,
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
	-2, 92,
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
	-2, 95,
	-1, 156,
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
	-2, 96,
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
	-2, 97,
	-1, 161,
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
	-1, 163,
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
	-1, 168,
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
	-1, 171,
	34, 122,
	-2, 64,
	-1, 185,
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
	-1, 190,
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
	-1, 191,
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
	-1, 197,
	34, 123,
	-2, 65,
	-1, 214,
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
	-2, 57,
	-1, 228,
	32, 81,
	-2, 11,
	-1, 239,
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
	-1, 251,
	10, 107,
	30, 107,
	42, 107,
	-2, 11,
	-1, 252,
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
	-1, 263,
	34, 124,
	-2, 68,
	-1, 266,
	10, 104,
	30, 104,
	42, 104,
	-2, 11,
	-1, 272,
	30, 108,
	42, 108,
	-2, 11,
	-1, 275,
	30, 105,
	42, 105,
	-2, 11,
}

const RubyNprod = 134
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 790

var RubyAct = []int{

	133, 109, 5, 199, 102, 96, 134, 56, 7, 57,
	90, 59, 2, 259, 112, 257, 180, 222, 171, 9,
	36, 37, 72, 73, 58, 99, 41, 268, 53, 202,
	178, 42, 43, 44, 45, 241, 6, 60, 46, 47,
	48, 49, 62, 63, 64, 256, 181, 65, 66, 67,
	68, 69, 70, 50, 71, 51, 221, 55, 54, 52,
	74, 224, 225, 6, 103, 103, 208, 61, 6, 104,
	6, 144, 61, 224, 225, 139, 6, 4, 116, 117,
	118, 119, 120, 121, 138, 276, 124, 125, 126, 127,
	128, 201, 129, 130, 131, 132, 274, 142, 60, 89,
	135, 136, 137, 271, 6, 250, 123, 242, 122, 209,
	143, 211, 210, 145, 234, 6, 217, 218, 216, 180,
	146, 228, 99, 153, 154, 232, 160, 159, 165, 93,
	162, 114, 115, 164, 75, 76, 229, 263, 158, 140,
	103, 177, 229, 230, 176, 175, 265, 172, 174, 179,
	111, 197, 111, 110, 101, 100, 168, 88, 220, 141,
	91, 94, 8, 192, 183, 193, 194, 196, 182, 1,
	173, 195, 187, 198, 170, 31, 30, 29, 28, 27,
	26, 204, 25, 205, 206, 207, 60, 24, 23, 212,
	22, 21, 20, 19, 18, 34, 213, 14, 215, 223,
	219, 38, 13, 12, 39, 15, 226, 10, 227, 236,
	35, 237, 11, 40, 32, 17, 33, 95, 238, 16,
	231, 233, 105, 106, 107, 108, 3, 240, 113, 245,
	243, 0, 246, 0, 248, 0, 0, 0, 0, 0,
	0, 0, 0, 255, 0, 0, 253, 247, 258, 249,
	260, 261, 0, 0, 0, 0, 0, 0, 0, 262,
	0, 0, 0, 0, 0, 267, 0, 0, 0, 270,
	0, 0, 273, 147, 148, 149, 150, 151, 152, 0,
	0, 155, 156, 157, 0, 161, 0, 163, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 184, 0, 0, 9, 36, 37, 0, 0,
	190, 191, 41, 186, 53, 189, 188, 42, 43, 44,
	45, 0, 0, 0, 46, 47, 48, 49, 0, 167,
	166, 0, 0, 0, 0, 0, 0, 0, 0, 50,
	0, 51, 0, 55, 54, 52, 9, 36, 37, 214,
	0, 0, 0, 41, 254, 53, 0, 0, 42, 43,
	44, 45, 0, 0, 0, 46, 47, 48, 49, 239,
	167, 166, 0, 77, 0, 0, 0, 0, 0, 0,
	50, 0, 51, 0, 55, 54, 52, 0, 0, 0,
	83, 84, 269, 251, 252, 78, 79, 80, 0, 0,
	0, 0, 0, 86, 81, 85, 82, 87, 0, 0,
	0, 0, 0, 0, 0, 264, 0, 0, 266, 0,
	0, 0, 0, 0, 9, 36, 37, 272, 0, 0,
	275, 41, 244, 53, 0, 0, 42, 43, 44, 45,
	0, 0, 0, 46, 47, 48, 49, 0, 167, 166,
	0, 0, 0, 0, 0, 0, 0, 0, 50, 0,
	51, 0, 55, 54, 52, 9, 36, 37, 0, 0,
	0, 0, 41, 203, 53, 0, 0, 42, 43, 44,
	45, 0, 0, 0, 46, 47, 48, 49, 0, 167,
	166, 0, 0, 0, 0, 0, 0, 0, 0, 50,
	0, 51, 0, 55, 54, 52, 9, 36, 37, 0,
	0, 0, 0, 41, 200, 53, 0, 0, 42, 43,
	44, 45, 0, 0, 0, 46, 47, 48, 49, 0,
	167, 166, 0, 0, 0, 0, 0, 0, 0, 0,
	50, 0, 51, 0, 55, 54, 52, 9, 36, 37,
	0, 0, 0, 0, 41, 169, 53, 0, 0, 42,
	43, 44, 45, 0, 0, 0, 46, 47, 48, 49,
	0, 167, 166, 0, 0, 0, 0, 0, 0, 0,
	0, 50, 0, 51, 0, 55, 54, 52, 98, 97,
	37, 93, 0, 0, 99, 41, 0, 53, 0, 0,
	42, 43, 44, 45, 0, 0, 92, 46, 47, 48,
	49, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 50, 0, 51, 0, 55, 54, 52, 9,
	36, 37, 0, 0, 0, 0, 41, 0, 53, 0,
	0, 42, 43, 44, 45, 0, 0, 0, 46, 47,
	48, 49, 0, 167, 166, 0, 0, 0, 0, 0,
	0, 0, 0, 50, 0, 51, 0, 55, 54, 52,
	9, 36, 37, 0, 0, 0, 0, 41, 0, 53,
	0, 0, 42, 43, 44, 45, 0, 0, 0, 46,
	47, 48, 49, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 50, 0, 51, 0, 55, 54,
	52, 9, 185, 37, 0, 0, 0, 0, 41, 0,
	53, 0, 0, 42, 43, 44, 45, 0, 0, 0,
	46, 47, 48, 49, 0, 77, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 50, 0, 51, 0, 55,
	54, 52, 83, 84, 235, 77, 0, 78, 79, 80,
	0, 0, 0, 0, 0, 86, 81, 85, 82, 87,
	0, 0, 83, 84, 0, 0, 0, 78, 79, 80,
	0, 0, 0, 0, 0, 86, 81, 85, 82, 87,
}
var RubyPact = []int{

	-34, 47, -1000, -38, -1000, 665, 7, -1000, 7, -9,
	7, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 39, -1000, -1000, -1000,
	-1000, 7, 7, 7, -1000, -1000, 7, 7, 7, 7,
	7, 7, -1000, 7, 16, 128, -1000, 751, 151, 76,
	583, 149, 148, 7, 7, 665, 665, 665, 665, 147,
	-1000, 665, -1000, -1000, 125, -1000, -1000, 7, 7, 7,
	7, 7, 7, 87, 84, 7, 7, 7, 7, 7,
	-1000, 7, 7, 7, 7, 7, -1000, 39, -9, 7,
	7, 7, 54, 132, 67, 7, 7, 7, 7, 7,
	-1000, -1000, 41, 7, -1000, -1000, 665, 665, 665, 665,
	665, 665, 7, 7, 665, 665, 665, 121, 665, 120,
	665, 147, 118, 751, 542, -16, 121, 121, -1000, 7,
	7, -2, -1000, 6, -1000, 706, 300, 7, 7, 7,
	7, 7, 7, 665, 665, 7, 7, 7, -1000, -1000,
	-1000, 7, 7, 7, 7, 7, -1000, -1000, 7, -1000,
	7, 145, -1000, -1000, -1000, 501, 61, -1000, -3, 460,
	7, -1000, 7, 7, 7, 34, -1000, 96, 7, -1000,
	7, 7, 14, 109, 111, 26, -17, -1000, 7, 44,
	-1000, -1000, 114, -1000, 137, 115, 104, 731, 7, -1000,
	7, -1000, 665, 624, 7, -1000, -1000, -1000, -1000, -1000,
	5, -1000, -1000, 97, -1000, -1000, 624, 419, 7, -1000,
	-1000, 41, -1000, 41, -1000, 83, 665, 665, 624, 7,
	341, -1000, 7, 32, -1000, -1000, -27, 41, -29, 41,
	7, 7, 7, 624, -1000, 131, -1000, -1000, 665, -1000,
	140, 665, 624, -1000, 7, -5, 7, 369, 7, 81,
	665, 7, 86, 665, -1000, 75, -1000,
}
var RubyPgo = []int{

	0, 231, 156, 226, 219, 216, 5, 215, 214, 213,
	212, 210, 207, 205, 204, 203, 202, 201, 197, 4,
	195, 194, 193, 192, 191, 190, 188, 187, 182, 180,
	179, 178, 177, 176, 175, 6, 10, 174, 172, 170,
	169, 168, 1, 167, 164, 161, 160, 159, 0, 14,
	3, 158,
}
var RubyR1 = []int{

	0, 40, 40, 40, 40, 3, 3, 35, 35, 35,
	35, 48, 48, 49, 49, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 12, 12, 12, 12, 12, 10, 10, 10,
	10, 10, 10, 36, 36, 46, 46, 46, 46, 45,
	45, 45, 45, 45, 42, 42, 42, 42, 42, 50,
	50, 50, 15, 39, 39, 16, 16, 18, 19, 19,
	47, 47, 13, 13, 21, 22, 23, 24, 25, 26,
	27, 28, 29, 30, 31, 32, 33, 34, 4, 7,
	8, 5, 5, 41, 41, 41, 41, 44, 44, 44,
	1, 1, 9, 9, 17, 17, 14, 14, 20, 6,
	6, 37, 43, 43, 43, 51, 51, 11, 11, 38,
	38, 38, 38, 38,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 3, 0, 2, 2,
	2, 0, 2, 0, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 5, 3,
	5, 3, 5, 5, 1, 1, 1, 5, 5, 1,
	1, 5, 5, 5, 0, 1, 1, 5, 5, 0,
	2, 2, 9, 0, 1, 6, 8, 6, 3, 6,
	1, 4, 5, 5, 3, 3, 3, 3, 5, 5,
	5, 5, 5, 6, 6, 5, 5, 5, 1, 1,
	5, 9, 9, 0, 6, 11, 12, 4, 9, 10,
	0, 1, 2, 2, 2, 2, 3, 3, 1, 3,
	7, 3, 0, 1, 5, 1, 2, 5, 6, 0,
	5, 3, 4, 2,
}
var RubyChk = []int{

	-1000, -40, 46, -3, 30, -48, 29, 46, -2, 5,
	-12, -10, -15, -16, -18, -13, -4, -7, -21, -22,
	-23, -24, -25, -26, -27, -28, -29, -30, -31, -32,
	-33, -34, -8, -5, -20, -11, 6, 7, -17, -14,
	-9, 12, 17, 18, 19, 20, 24, 25, 26, 27,
	39, 41, 45, 14, 44, 43, -48, -48, 33, -48,
	-48, 33, -48, -48, -48, -48, -48, -48, -48, -48,
	-48, -48, 6, 7, 44, 6, 7, 4, 26, 27,
	28, 35, 37, 21, 22, 36, 34, 38, 6, 23,
	-36, -46, 23, 8, -45, -2, -6, 6, 5, 11,
	6, 6, -19, -48, -19, -2, -2, -2, -2, -42,
	6, 5, -49, -2, 6, 7, -48, -48, -48, -48,
	-48, -48, 21, 22, -48, -48, -48, -48, -48, -48,
	-48, -48, -48, -48, -35, -48, -48, -48, 30, 21,
	7, -47, 30, -48, 30, -48, -35, -2, -2, -2,
	-2, -2, -2, -48, -48, -2, -2, -2, -36, 6,
	5, -2, 10, -2, -42, 10, 30, 29, -2, 13,
	-37, 34, -36, -39, -36, -35, -19, -48, 32, -35,
	10, 40, -41, -44, -2, 6, 13, -38, 16, 15,
	-2, -2, -48, -48, -48, -48, -43, 6, -42, -50,
	13, 30, 32, 13, -48, -48, -48, -48, 32, 13,
	16, 15, -48, -35, -2, -6, 9, 5, 6, -6,
	-51, 30, 34, -48, 29, 30, -35, -35, 7, 5,
	6, -49, 10, -49, 10, 23, -48, -48, -35, -2,
	-35, 30, 10, -50, 13, -48, -48, -49, -48, -49,
	22, -2, -2, -35, 13, -48, 13, 42, -48, 42,
	-48, -48, -35, 6, -2, 6, -2, -48, 32, 23,
	-48, 22, -2, -48, 10, -2, 10,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 0, 11, 4, 11, 15,
	-2, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, -2, 43, 44, 45,
	46, 11, 11, 11, 98, 99, 11, 11, 11, 11,
	11, 11, 118, 11, 0, 0, 12, 6, 0, 0,
	0, 0, 0, 11, 11, 0, 0, 0, 0, 64,
	13, 0, 114, 115, 0, 112, 113, 11, 11, 11,
	11, 11, 11, 0, 0, 11, 11, 11, -2, 11,
	47, -2, 11, 11, -2, -2, 56, -2, 15, -2,
	11, 11, 0, 0, 0, -2, -2, -2, -2, 11,
	65, 66, 11, -2, 116, 117, 0, 0, 0, 0,
	0, 0, 11, 11, 0, 0, 0, 0, 0, 0,
	0, 64, 0, 0, 0, 0, 0, 73, 7, 11,
	-2, 0, 7, 0, 14, 103, 0, -2, -2, -2,
	-2, -2, -2, 0, 0, -2, -2, -2, 48, 59,
	60, -2, 11, -2, 11, 11, 8, 9, -2, 119,
	11, -2, 50, 69, 74, 0, 0, 78, 0, 0,
	11, 100, 11, 11, 11, -2, 127, 0, 11, 7,
	-2, -2, 0, 0, 0, 0, 0, -2, 11, 7,
	75, 7, 0, 77, 0, 13, 13, 0, 11, 128,
	11, 7, 0, 133, -2, 58, 53, 61, 62, 63,
	7, 125, 121, 0, 70, 71, 69, 0, -2, 67,
	68, 11, 13, 11, 13, 0, 0, 0, 131, -2,
	0, 126, 11, 0, 76, 79, 0, 11, 0, 11,
	11, -2, -2, 132, 120, 0, 72, 101, 0, 102,
	0, 0, 130, -2, 11, 0, -2, 0, 11, 0,
	0, 11, -2, 0, 109, -2, 106,
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
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 49:
		//line parser.y:198
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 50:
		//line parser.y:205
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 51:
		//line parser.y:213
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 52:
		//line parser.y:220
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 53:
		//line parser.y:229
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 54:
		//line parser.y:231
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 55:
		//line parser.y:234
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 56:
		//line parser.y:236
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 57:
		//line parser.y:238
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 58:
		//line parser.y:240
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 59:
		//line parser.y:243
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 60:
		//line parser.y:245
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 61:
		//line parser.y:247
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 62:
		//line parser.y:249
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 63:
		//line parser.y:251
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 64:
		//line parser.y:253
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 65:
		//line parser.y:255
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 66:
		//line parser.y:257
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:259
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:261
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 72:
		//line parser.y:271
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 73:
		//line parser.y:279
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 74:
		//line parser.y:280
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 75:
		//line parser.y:283
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 76:
		//line parser.y:290
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 77:
		//line parser.y:300
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 78:
		//line parser.y:309
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 79:
		//line parser.y:315
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 80:
		//line parser.y:323
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 81:
		//line parser.y:327
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 82:
		//line parser.y:332
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 83:
		//line parser.y:339
		{ // FIXME: this is a hack. If this rule is not present, "foo = 5" will not parse
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 84:
		//line parser.y:346
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 85:
		//line parser.y:347
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 86:
		//line parser.y:348
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 87:
		//line parser.y:349
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 88:
		//line parser.y:352
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 89:
		//line parser.y:361
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 90:
		//line parser.y:370
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 91:
		//line parser.y:379
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 92:
		//line parser.y:388
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "%"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 93:
		//line parser.y:397
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-5].genericValue,
				Func:   ast.BareReference{Name: "<<"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 94:
		//line parser.y:406
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-5].genericValue,
				Func:   ast.BareReference{Name: ">>"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 95:
		//line parser.y:415
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 96:
		//line parser.y:424
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 97:
		//line parser.y:433
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "^"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 98:
		//line parser.y:441
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 99:
		//line parser.y:442
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 100:
		//line parser.y:444
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 101:
		//line parser.y:448
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 102:
		//line parser.y:456
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 103:
		//line parser.y:464
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 104:
		//line parser.y:466
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 105:
		//line parser.y:468
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 106:
		//line parser.y:470
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 107:
		//line parser.y:473
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 108:
		//line parser.y:480
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 109:
		//line parser.y:487
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 110:
		//line parser.y:495
		{
			RubyVAL.genericValue = nil
		}
	case 111:
		//line parser.y:496
		{
			RubyVAL.genericValue = nil
		}
	case 112:
		//line parser.y:499
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 113:
		//line parser.y:501
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 114:
		//line parser.y:504
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 115:
		//line parser.y:506
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 116:
		//line parser.y:509
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 117:
		//line parser.y:511
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 118:
		//line parser.y:514
		{
			RubyVAL.genericValue = ast.FileNameConstReference{}
		}
	case 119:
		//line parser.y:517
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 120:
		//line parser.y:519
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 121:
		//line parser.y:527
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 122:
		//line parser.y:529
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 123:
		//line parser.y:531
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 124:
		//line parser.y:533
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 127:
		//line parser.y:538
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 128:
		//line parser.y:545
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 129:
		//line parser.y:553
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 130:
		//line parser.y:555
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 131:
		//line parser.y:562
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 132:
		//line parser.y:569
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 133:
		//line parser.y:576
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	}
	goto Rubystack /* stack new state and value */
}
