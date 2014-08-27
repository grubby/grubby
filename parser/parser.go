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

//line parser.y:544

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 12,
	-1, 11,
	23, 12,
	-2, 17,
	-1, 33,
	1, 39,
	4, 39,
	10, 39,
	13, 39,
	15, 39,
	16, 39,
	28, 39,
	30, 39,
	31, 39,
	34, 39,
	35, 39,
	36, 39,
	42, 39,
	46, 39,
	-2, 12,
	-1, 81,
	5, 12,
	6, 12,
	8, 12,
	-2, 46,
	-1, 84,
	10, 12,
	-2, 48,
	-1, 87,
	10, 12,
	-2, 51,
	-1, 88,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 52,
	-1, 90,
	1, 39,
	4, 39,
	10, 39,
	13, 39,
	15, 39,
	16, 39,
	28, 39,
	30, 39,
	31, 39,
	34, 39,
	35, 39,
	36, 39,
	42, 39,
	46, 39,
	-2, 12,
	-1, 92,
	34, 12,
	-2, 8,
	-1, 98,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 81,
	-1, 99,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 82,
	-1, 100,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 83,
	-1, 101,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 84,
	-1, 106,
	4, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 8,
	-1, 129,
	32, 77,
	-2, 12,
	-1, 136,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 49,
	-1, 137,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 85,
	-1, 138,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 86,
	-1, 139,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 87,
	-1, 140,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 88,
	-1, 141,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 89,
	-1, 142,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 90,
	-1, 146,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 80,
	-1, 148,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 79,
	-1, 153,
	4, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 11,
	-1, 156,
	34, 115,
	-2, 61,
	-1, 170,
	4, 39,
	28, 39,
	34, 39,
	35, 39,
	36, 39,
	-2, 12,
	-1, 180,
	34, 116,
	-2, 62,
	-1, 197,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 54,
	-1, 211,
	32, 78,
	-2, 12,
	-1, 222,
	4, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 8,
	-1, 234,
	10, 100,
	30, 100,
	42, 100,
	-2, 12,
	-1, 235,
	4, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 8,
	-1, 246,
	34, 117,
	-2, 65,
	-1, 249,
	10, 97,
	30, 97,
	42, 97,
	-2, 12,
	-1, 255,
	30, 101,
	42, 101,
	-2, 12,
	-1, 258,
	30, 98,
	42, 98,
	-2, 12,
}

const RubyNprod = 127
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 709

var RubyAct = []int{

	122, 182, 6, 89, 95, 102, 8, 105, 53, 2,
	54, 83, 56, 242, 240, 205, 10, 33, 34, 69,
	70, 156, 92, 38, 55, 50, 251, 185, 39, 40,
	41, 42, 163, 123, 57, 43, 44, 45, 46, 59,
	60, 61, 165, 7, 62, 63, 64, 65, 66, 67,
	47, 68, 48, 224, 52, 51, 49, 71, 7, 133,
	7, 96, 96, 191, 58, 7, 97, 207, 208, 58,
	204, 184, 166, 259, 131, 109, 110, 111, 112, 113,
	114, 115, 116, 117, 239, 118, 119, 120, 121, 153,
	74, 57, 7, 124, 125, 126, 9, 7, 4, 5,
	207, 208, 128, 132, 257, 82, 134, 254, 225, 252,
	233, 127, 75, 76, 77, 192, 217, 194, 193, 215,
	80, 78, 79, 7, 199, 165, 149, 150, 143, 96,
	162, 145, 144, 161, 86, 200, 201, 157, 159, 147,
	135, 92, 107, 108, 72, 73, 211, 88, 175, 129,
	176, 177, 98, 99, 100, 101, 178, 248, 106, 212,
	246, 160, 181, 212, 213, 164, 187, 94, 188, 189,
	190, 57, 104, 180, 195, 104, 103, 93, 81, 198,
	203, 202, 206, 130, 84, 87, 168, 179, 167, 1,
	158, 172, 219, 155, 220, 28, 214, 216, 27, 136,
	137, 138, 139, 140, 141, 142, 26, 146, 196, 148,
	25, 226, 228, 24, 23, 229, 209, 231, 210, 74,
	22, 21, 20, 230, 169, 232, 238, 19, 221, 31,
	15, 241, 35, 243, 244, 14, 74, 223, 218, 13,
	36, 75, 76, 77, 16, 11, 32, 12, 250, 80,
	78, 79, 253, 37, 29, 256, 236, 18, 75, 76,
	77, 30, 17, 3, 0, 197, 80, 78, 79, 245,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 10,
	33, 34, 0, 0, 0, 222, 38, 171, 50, 174,
	173, 39, 40, 41, 42, 0, 0, 0, 43, 44,
	45, 46, 0, 152, 151, 0, 0, 0, 0, 234,
	235, 0, 0, 47, 0, 48, 0, 52, 51, 49,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 247, 0, 0, 249, 0, 0, 0, 0, 0,
	10, 33, 34, 255, 0, 0, 258, 38, 237, 50,
	0, 0, 39, 40, 41, 42, 0, 0, 0, 43,
	44, 45, 46, 0, 152, 151, 0, 0, 0, 0,
	0, 0, 0, 0, 47, 0, 48, 0, 52, 51,
	49, 10, 33, 34, 0, 0, 0, 0, 38, 227,
	50, 0, 0, 39, 40, 41, 42, 0, 0, 0,
	43, 44, 45, 46, 0, 152, 151, 0, 0, 0,
	0, 0, 0, 0, 0, 47, 0, 48, 0, 52,
	51, 49, 10, 33, 34, 0, 0, 0, 0, 38,
	186, 50, 0, 0, 39, 40, 41, 42, 0, 0,
	0, 43, 44, 45, 46, 0, 152, 151, 0, 0,
	0, 0, 0, 0, 0, 0, 47, 0, 48, 0,
	52, 51, 49, 10, 33, 34, 0, 0, 0, 0,
	38, 183, 50, 0, 0, 39, 40, 41, 42, 0,
	0, 0, 43, 44, 45, 46, 0, 152, 151, 0,
	0, 0, 0, 0, 0, 0, 0, 47, 0, 48,
	0, 52, 51, 49, 10, 33, 34, 0, 0, 0,
	0, 38, 154, 50, 0, 0, 39, 40, 41, 42,
	0, 0, 0, 43, 44, 45, 46, 0, 152, 151,
	0, 0, 0, 0, 0, 0, 0, 0, 47, 0,
	48, 0, 52, 51, 49, 91, 90, 34, 86, 0,
	0, 92, 38, 0, 50, 0, 0, 39, 40, 41,
	42, 0, 0, 85, 43, 44, 45, 46, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 47,
	0, 48, 0, 52, 51, 49, 10, 33, 34, 0,
	0, 0, 0, 38, 0, 50, 0, 0, 39, 40,
	41, 42, 0, 0, 0, 43, 44, 45, 46, 0,
	152, 151, 0, 0, 0, 0, 0, 0, 0, 0,
	47, 0, 48, 0, 52, 51, 49, 10, 33, 34,
	0, 0, 0, 0, 38, 0, 50, 0, 0, 39,
	40, 41, 42, 0, 0, 0, 43, 44, 45, 46,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 47, 0, 48, 0, 52, 51, 49, 10, 170,
	34, 0, 0, 0, 0, 38, 0, 50, 0, 0,
	39, 40, 41, 42, 0, 0, 0, 43, 44, 45,
	46, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 47, 0, 48, 0, 52, 51, 49,
}
var RubyPact = []int{

	-37, 68, -1000, -40, -1000, -1000, 622, 14, -1000, 14,
	-9, 14, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 36, -1000, -1000, -1000, -1000, 14, 14,
	14, -1000, -1000, 14, 14, 14, 14, 14, 14, -1000,
	14, 13, 138, -1000, 232, 172, 82, 540, 171, 161,
	14, 14, 622, 622, 622, 622, 170, -1000, 622, -1000,
	-1000, 136, -1000, -1000, 14, 14, 14, 14, 14, 14,
	14, 14, 14, -1000, 14, 14, 14, 14, 14, -1000,
	36, -9, 14, 14, 14, 81, 142, 44, 14, 14,
	14, 14, 14, -1000, -1000, 29, 14, -1000, -1000, 622,
	622, 622, 622, 622, 622, 622, 126, 622, 129, 622,
	170, 117, 232, 499, -13, 126, 126, -1000, 14, 14,
	0, -1000, 32, -1000, 663, 274, 14, 14, 14, 14,
	14, 14, 14, -1000, -1000, -1000, 14, 14, 14, 14,
	14, -1000, -1000, 14, -1000, 14, 167, -1000, -1000, -1000,
	458, 41, -1000, -5, 417, 14, -1000, 14, 14, 14,
	31, -1000, 102, 14, -1000, 11, 115, 130, 40, -19,
	-1000, 14, 38, -1000, -1000, 139, -1000, 158, 109, 106,
	215, 14, -1000, 14, -1000, 622, 581, 14, -1000, -1000,
	-1000, -1000, -1000, 23, -1000, -1000, 98, -1000, -1000, 581,
	376, 14, -1000, -1000, 29, -1000, 29, -1000, 88, 622,
	622, 581, 14, 335, -1000, 14, 71, -1000, -1000, -28,
	29, -29, 29, 14, 14, 14, 581, -1000, 154, -1000,
	-1000, 622, -1000, 151, 622, 581, -1000, 14, -6, 14,
	86, 14, 85, 622, 14, 94, 622, -1000, 63, -1000,
}
var RubyPgo = []int{

	0, 264, 89, 263, 262, 261, 3, 257, 254, 253,
	247, 246, 245, 244, 240, 239, 235, 232, 230, 4,
	229, 227, 222, 221, 220, 214, 213, 210, 206, 198,
	195, 33, 11, 193, 191, 190, 189, 188, 5, 187,
	186, 185, 184, 183, 0, 7, 1, 180,
}
var RubyR1 = []int{

	0, 36, 36, 36, 36, 3, 3, 3, 31, 31,
	31, 31, 44, 44, 45, 45, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 12,
	12, 12, 12, 12, 10, 10, 10, 10, 10, 10,
	32, 32, 42, 42, 42, 42, 41, 41, 41, 41,
	41, 38, 38, 38, 38, 38, 46, 46, 46, 15,
	35, 35, 16, 16, 18, 19, 19, 43, 43, 13,
	13, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	30, 4, 7, 8, 5, 5, 37, 37, 37, 37,
	40, 40, 40, 1, 1, 9, 9, 17, 17, 14,
	14, 20, 6, 6, 33, 39, 39, 39, 47, 47,
	11, 11, 34, 34, 34, 34, 34,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 1, 3, 0, 2,
	2, 2, 0, 2, 0, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 3, 5, 3, 5, 3, 5,
	5, 1, 1, 1, 5, 5, 1, 1, 5, 5,
	5, 0, 1, 1, 5, 5, 0, 2, 2, 9,
	0, 1, 6, 8, 6, 3, 6, 1, 4, 5,
	5, 3, 3, 3, 3, 5, 5, 5, 5, 5,
	5, 1, 1, 5, 9, 9, 0, 6, 11, 12,
	4, 9, 10, 0, 1, 2, 2, 2, 2, 3,
	3, 1, 3, 7, 3, 0, 1, 5, 1, 2,
	5, 6, 0, 5, 3, 4, 2,
}
var RubyChk = []int{

	-1000, -36, 46, -3, 30, 31, -44, 29, 46, -2,
	5, -12, -10, -15, -16, -18, -13, -4, -7, -21,
	-22, -23, -24, -25, -26, -27, -28, -29, -30, -8,
	-5, -20, -11, 6, 7, -17, -14, -9, 12, 17,
	18, 19, 20, 24, 25, 26, 27, 39, 41, 45,
	14, 44, 43, -44, -44, 33, -44, -44, 33, -44,
	-44, -44, -44, -44, -44, -44, -44, -44, -44, 6,
	7, 44, 6, 7, 4, 26, 27, 28, 35, 36,
	34, 6, 23, -32, -42, 23, 8, -41, -2, -6,
	6, 5, 11, 6, 6, -19, -44, -19, -2, -2,
	-2, -2, -38, 6, 5, -45, -2, 6, 7, -44,
	-44, -44, -44, -44, -44, -44, -44, -44, -44, -44,
	-44, -44, -44, -31, -44, -44, -44, 30, 21, 7,
	-43, 30, -44, 30, -44, -31, -2, -2, -2, -2,
	-2, -2, -2, -32, 6, 5, -2, 10, -2, -38,
	10, 30, 29, -2, 13, -33, 34, -32, -35, -32,
	-31, -19, -44, 32, -31, 10, 40, -37, -40, -2,
	6, 13, -34, 16, 15, -44, -44, -44, -44, -39,
	6, -38, -46, 13, 30, 32, 13, -44, -44, -44,
	-44, 32, 13, 16, 15, -44, -31, -2, -6, 9,
	5, 6, -6, -47, 30, 34, -44, 29, 30, -31,
	-31, 7, 5, 6, -45, 10, -45, 10, 23, -44,
	-44, -31, -2, -31, 30, 10, -46, 13, -44, -44,
	-45, -44, -45, 22, -2, -2, -31, 13, -44, 13,
	42, -44, 42, -44, -44, -31, 6, -2, 6, -2,
	-44, 32, 23, -44, 22, -2, -44, 10, -2, 10,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 6, 0, 12, 4, 12,
	16, -2, 18, 19, 20, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, -2, 40, 41, 42, 43, 12, 12,
	12, 91, 92, 12, 12, 12, 12, 12, 12, 111,
	12, 0, 0, 13, 7, 0, 0, 0, 0, 0,
	12, 12, 0, 0, 0, 0, 61, 14, 0, 107,
	108, 0, 105, 106, 12, 12, 12, 12, 12, 12,
	12, -2, 12, 44, -2, 12, 12, -2, -2, 53,
	-2, 16, -2, 12, 12, 0, 0, 0, -2, -2,
	-2, -2, 12, 62, 63, 12, -2, 109, 110, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	61, 0, 0, 0, 0, 0, 70, 8, 12, -2,
	0, 8, 0, 15, 96, 0, -2, -2, -2, -2,
	-2, -2, -2, 45, 56, 57, -2, 12, -2, 12,
	12, 9, 10, -2, 112, 12, -2, 47, 66, 71,
	0, 0, 75, 0, 0, 12, 93, 12, 12, 12,
	-2, 120, 0, 12, 8, 0, 0, 0, 0, 0,
	-2, 12, 8, 72, 8, 0, 74, 0, 14, 14,
	0, 12, 121, 12, 8, 0, 126, -2, 55, 50,
	58, 59, 60, 8, 118, 114, 0, 67, 68, 66,
	0, -2, 64, 65, 12, 14, 12, 14, 0, 0,
	0, 124, -2, 0, 119, 12, 0, 73, 76, 0,
	12, 0, 12, 12, -2, -2, 125, 113, 0, 69,
	94, 0, 95, 0, 0, 123, -2, 12, 0, -2,
	0, 12, 0, 0, 12, -2, 0, 102, -2, 99,
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
			RubyVAL.genericValue = nil
		}
	case 7:
		//line parser.y:157
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 8:
		//line parser.y:160
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 9:
		//line parser.y:162
		{ /* do nothing */
		}
	case 10:
		//line parser.y:164
		{ /* do nothing */
		}
	case 11:
		//line parser.y:166
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
			}
		}
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
		//line parser.y:180
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 45:
		//line parser.y:187
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 46:
		//line parser.y:195
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 47:
		//line parser.y:202
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 48:
		//line parser.y:210
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 49:
		//line parser.y:217
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 50:
		//line parser.y:226
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 51:
		//line parser.y:228
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 52:
		//line parser.y:231
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 53:
		//line parser.y:233
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 54:
		//line parser.y:235
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 55:
		//line parser.y:237
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 56:
		//line parser.y:240
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 57:
		//line parser.y:242
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 58:
		//line parser.y:244
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 59:
		//line parser.y:246
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 60:
		//line parser.y:248
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 61:
		//line parser.y:250
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 62:
		//line parser.y:252
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 63:
		//line parser.y:254
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 64:
		//line parser.y:256
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 65:
		//line parser.y:258
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 69:
		//line parser.y:268
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 70:
		//line parser.y:276
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 71:
		//line parser.y:277
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 72:
		//line parser.y:280
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 73:
		//line parser.y:287
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 74:
		//line parser.y:297
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 75:
		//line parser.y:306
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 76:
		//line parser.y:312
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 77:
		//line parser.y:320
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 78:
		//line parser.y:324
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 79:
		//line parser.y:329
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 80:
		//line parser.y:336
		{ // FIXME: this is a hack. If this rule is not present, "foo = 5" will not parse
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 81:
		//line parser.y:343
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 82:
		//line parser.y:344
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 83:
		//line parser.y:345
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 84:
		//line parser.y:346
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 85:
		//line parser.y:349
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 86:
		//line parser.y:358
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 87:
		//line parser.y:367
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 88:
		//line parser.y:376
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 89:
		//line parser.y:385
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 90:
		//line parser.y:394
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 91:
		//line parser.y:402
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 92:
		//line parser.y:403
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 93:
		//line parser.y:405
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 94:
		//line parser.y:409
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 95:
		//line parser.y:417
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 96:
		//line parser.y:425
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 97:
		//line parser.y:427
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 98:
		//line parser.y:429
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 99:
		//line parser.y:431
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 100:
		//line parser.y:434
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 101:
		//line parser.y:441
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 102:
		//line parser.y:448
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 103:
		//line parser.y:456
		{
			RubyVAL.genericValue = nil
		}
	case 104:
		//line parser.y:457
		{
			RubyVAL.genericValue = nil
		}
	case 105:
		//line parser.y:460
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 106:
		//line parser.y:462
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 107:
		//line parser.y:465
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 108:
		//line parser.y:467
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 109:
		//line parser.y:470
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 110:
		//line parser.y:472
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 111:
		//line parser.y:475
		{
			RubyVAL.genericValue = ast.FileNameConstReference{}
		}
	case 112:
		//line parser.y:478
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 113:
		//line parser.y:480
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 114:
		//line parser.y:488
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 115:
		//line parser.y:490
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 116:
		//line parser.y:492
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 117:
		//line parser.y:494
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 120:
		//line parser.y:499
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 121:
		//line parser.y:506
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 122:
		//line parser.y:514
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 123:
		//line parser.y:516
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 124:
		//line parser.y:523
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 125:
		//line parser.y:530
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 126:
		//line parser.y:537
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	}
	goto Rubystack /* stack new state and value */
}
