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

//line parser.y:562

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 12,
	-1, 11,
	1, 17,
	4, 17,
	10, 17,
	13, 17,
	15, 17,
	16, 17,
	28, 17,
	30, 17,
	31, 17,
	34, 17,
	35, 17,
	36, 17,
	42, 17,
	46, 17,
	-2, 12,
	-1, 12,
	23, 12,
	-2, 18,
	-1, 13,
	23, 12,
	-2, 19,
	-1, 14,
	23, 12,
	-2, 20,
	-1, 15,
	23, 12,
	-2, 21,
	-1, 83,
	5, 12,
	6, 12,
	8, 12,
	-2, 45,
	-1, 85,
	10, 12,
	-2, 47,
	-1, 88,
	10, 12,
	-2, 50,
	-1, 89,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 51,
	-1, 91,
	1, 17,
	4, 17,
	10, 17,
	13, 17,
	15, 17,
	16, 17,
	28, 17,
	30, 17,
	31, 17,
	34, 17,
	35, 17,
	36, 17,
	42, 17,
	46, 17,
	-2, 12,
	-1, 93,
	34, 12,
	-2, 8,
	-1, 105,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 83,
	-1, 106,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 84,
	-1, 107,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 85,
	-1, 108,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 86,
	-1, 113,
	4, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 8,
	-1, 137,
	32, 76,
	-2, 12,
	-1, 144,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 48,
	-1, 145,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 87,
	-1, 146,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 88,
	-1, 147,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 89,
	-1, 148,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 90,
	-1, 149,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 91,
	-1, 150,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 92,
	-1, 155,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 78,
	-1, 160,
	4, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 11,
	-1, 163,
	34, 117,
	-2, 60,
	-1, 165,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 79,
	-1, 166,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 80,
	-1, 167,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 81,
	-1, 168,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 82,
	-1, 181,
	4, 17,
	28, 17,
	34, 17,
	35, 17,
	36, 17,
	-2, 12,
	-1, 191,
	34, 118,
	-2, 61,
	-1, 208,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 53,
	-1, 222,
	32, 77,
	-2, 12,
	-1, 233,
	4, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 8,
	-1, 245,
	10, 102,
	30, 102,
	42, 102,
	-2, 12,
	-1, 246,
	4, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 8,
	-1, 257,
	34, 119,
	-2, 64,
	-1, 260,
	10, 99,
	30, 99,
	42, 99,
	-2, 12,
	-1, 266,
	30, 103,
	42, 103,
	-2, 12,
	-1, 269,
	30, 100,
	42, 100,
	-2, 12,
}

const RubyNprod = 129
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 690

var RubyAct = []int{

	126, 193, 6, 109, 102, 84, 61, 62, 52, 112,
	53, 90, 55, 57, 58, 59, 60, 92, 91, 12,
	87, 8, 2, 93, 39, 253, 51, 251, 176, 40,
	41, 42, 43, 216, 54, 86, 44, 45, 46, 47,
	66, 67, 68, 163, 63, 69, 70, 71, 72, 73,
	74, 48, 75, 49, 127, 38, 37, 50, 177, 7,
	7, 262, 202, 56, 56, 7, 4, 5, 103, 103,
	196, 174, 235, 104, 250, 7, 141, 114, 115, 116,
	117, 118, 119, 120, 121, 215, 122, 123, 124, 125,
	218, 219, 55, 76, 128, 129, 130, 131, 132, 133,
	160, 195, 134, 218, 219, 136, 139, 9, 7, 98,
	140, 270, 263, 142, 135, 77, 78, 79, 97, 96,
	76, 95, 265, 82, 80, 81, 244, 151, 156, 268,
	7, 203, 236, 205, 204, 164, 228, 103, 173, 229,
	170, 172, 77, 78, 79, 210, 176, 226, 7, 157,
	82, 80, 81, 211, 212, 186, 89, 187, 188, 93,
	154, 153, 152, 189, 87, 76, 222, 192, 143, 137,
	105, 106, 107, 108, 99, 100, 113, 198, 259, 199,
	200, 201, 55, 64, 65, 206, 101, 77, 78, 79,
	171, 223, 257, 217, 175, 82, 80, 81, 209, 94,
	213, 223, 224, 230, 83, 231, 111, 191, 214, 225,
	227, 111, 110, 138, 85, 144, 145, 146, 147, 148,
	149, 150, 237, 239, 155, 88, 240, 179, 242, 190,
	178, 165, 166, 167, 168, 1, 241, 249, 243, 169,
	207, 183, 252, 180, 254, 255, 162, 32, 220, 31,
	221, 30, 29, 28, 27, 26, 25, 24, 23, 261,
	232, 35, 19, 264, 13, 18, 267, 17, 14, 234,
	10, 11, 12, 20, 36, 16, 15, 39, 182, 51,
	185, 184, 40, 41, 42, 43, 33, 208, 247, 44,
	45, 46, 47, 22, 159, 158, 34, 21, 3, 0,
	0, 256, 0, 0, 48, 0, 49, 233, 38, 37,
	50, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 10, 11, 12, 0, 0, 0, 0, 39, 248,
	51, 245, 246, 40, 41, 42, 43, 0, 0, 0,
	44, 45, 46, 47, 0, 159, 158, 0, 0, 0,
	0, 0, 0, 258, 0, 48, 260, 49, 0, 38,
	37, 50, 10, 11, 12, 266, 0, 0, 269, 39,
	238, 51, 0, 0, 40, 41, 42, 43, 0, 0,
	0, 44, 45, 46, 47, 0, 159, 158, 0, 0,
	0, 0, 0, 0, 0, 0, 48, 0, 49, 0,
	38, 37, 50, 10, 11, 12, 0, 0, 0, 0,
	39, 197, 51, 0, 0, 40, 41, 42, 43, 0,
	0, 0, 44, 45, 46, 47, 0, 159, 158, 0,
	0, 0, 0, 0, 0, 0, 0, 48, 0, 49,
	0, 38, 37, 50, 10, 11, 12, 0, 0, 0,
	0, 39, 194, 51, 0, 0, 40, 41, 42, 43,
	0, 0, 0, 44, 45, 46, 47, 0, 159, 158,
	0, 0, 0, 0, 0, 0, 0, 0, 48, 0,
	49, 0, 38, 37, 50, 10, 11, 12, 0, 0,
	0, 0, 39, 161, 51, 0, 0, 40, 41, 42,
	43, 0, 0, 0, 44, 45, 46, 47, 0, 159,
	158, 0, 0, 0, 0, 0, 0, 0, 0, 48,
	0, 49, 0, 38, 37, 50, 10, 11, 12, 0,
	0, 0, 0, 39, 0, 51, 0, 0, 40, 41,
	42, 43, 0, 0, 0, 44, 45, 46, 47, 0,
	159, 158, 0, 0, 0, 0, 0, 0, 0, 0,
	48, 0, 49, 0, 38, 37, 50, 10, 11, 12,
	0, 0, 0, 93, 39, 0, 51, 0, 0, 40,
	41, 42, 43, 0, 0, 0, 44, 45, 46, 47,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 48, 0, 49, 0, 38, 37, 50, 10, 11,
	12, 0, 0, 0, 0, 39, 0, 51, 0, 0,
	40, 41, 42, 43, 0, 0, 0, 44, 45, 46,
	47, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 48, 0, 49, 0, 38, 37, 50, 10,
	181, 12, 0, 0, 0, 0, 39, 0, 51, 0,
	0, 40, 41, 42, 43, 0, 0, 0, 44, 45,
	46, 47, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 48, 0, 49, 0, 38, 37, 50,
}
var RubyPact = []int{

	-24, 36, -1000, -25, -1000, -1000, 603, 79, -1000, 79,
	1, 31, 79, 79, 79, 79, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 0, 177, 79,
	79, 79, -1000, -1000, 79, 79, 79, 79, 79, 79,
	-1000, 79, -1000, 161, 198, 12, 193, 98, 96, 95,
	86, -1000, -1000, 168, -1000, -1000, 180, 79, 79, 603,
	603, 603, 603, 206, -1000, 603, 79, 79, 79, 79,
	79, 79, 79, 79, -1000, 79, 79, 79, 79, 79,
	-1000, 31, 1, 79, 79, 79, 79, 79, 79, -1000,
	-1000, 79, 84, 162, 76, 79, 79, 79, 79, 79,
	-1000, -1000, 46, 79, 603, 603, 603, 603, 603, 603,
	603, 156, 150, 603, 206, 139, 161, 480, 9, 156,
	603, 603, 603, 603, 156, -1000, 79, 79, 39, -1000,
	18, -1000, 644, 265, 79, 79, 79, 79, 79, 79,
	79, -1000, -1000, -1000, 79, 79, 79, 79, -1000, -1000,
	79, -1000, 79, 201, -1000, 79, 79, 79, 79, -1000,
	-1000, 439, 71, -1000, 38, 398, 79, -1000, 79, 79,
	79, 30, -1000, 118, 79, -1000, 562, 136, 148, 55,
	-1, -1000, 79, 74, -1000, -1000, 159, -1000, 196, 137,
	126, 116, 79, -1000, 79, -1000, 603, 521, 79, -1000,
	-1000, -1000, -1000, -1000, 42, -1000, -1000, 122, -1000, -1000,
	521, 357, 79, -1000, -1000, 46, -1000, 46, -1000, 104,
	603, 603, 521, 79, 316, -1000, 79, 61, -1000, -1000,
	-15, 46, -17, 46, 79, 79, 79, 521, -1000, 186,
	-1000, -1000, 603, -1000, 172, 603, 521, -1000, 79, 29,
	79, 89, 79, 100, 603, 79, 119, 603, -1000, 101,
	-1000,
}
var RubyPgo = []int{

	0, 299, 100, 298, 297, 296, 11, 293, 286, 276,
	275, 274, 273, 268, 267, 265, 264, 262, 4, 261,
	258, 257, 256, 255, 254, 253, 252, 251, 249, 247,
	54, 5, 246, 241, 239, 235, 230, 3, 229, 227,
	225, 214, 213, 0, 9, 1, 208,
}
var RubyR1 = []int{

	0, 35, 35, 35, 35, 3, 3, 3, 30, 30,
	30, 30, 43, 43, 44, 44, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 10, 10, 10, 10, 10, 10, 31,
	31, 41, 41, 41, 41, 40, 40, 40, 40, 40,
	37, 37, 37, 37, 37, 45, 45, 45, 14, 34,
	34, 15, 15, 17, 18, 18, 42, 42, 12, 12,
	12, 12, 12, 20, 21, 22, 23, 24, 25, 26,
	27, 28, 29, 4, 7, 8, 5, 5, 36, 36,
	36, 36, 39, 39, 39, 1, 1, 9, 9, 16,
	16, 13, 13, 19, 6, 6, 32, 38, 38, 38,
	46, 46, 11, 11, 33, 33, 33, 33, 33,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 1, 3, 0, 2,
	2, 2, 0, 2, 0, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 5, 3, 5, 3, 5, 5,
	1, 1, 1, 5, 5, 1, 1, 5, 5, 5,
	0, 1, 1, 5, 5, 0, 2, 2, 9, 0,
	1, 6, 8, 6, 3, 6, 1, 4, 5, 5,
	5, 5, 5, 3, 3, 3, 3, 5, 5, 5,
	5, 5, 5, 1, 1, 5, 9, 9, 0, 6,
	11, 12, 4, 9, 10, 0, 1, 2, 2, 2,
	2, 3, 3, 1, 3, 7, 3, 0, 1, 5,
	1, 2, 5, 6, 0, 5, 3, 4, 2,
}
var RubyChk = []int{

	-1000, -35, 46, -3, 30, 31, -43, 29, 46, -2,
	5, 6, 7, -16, -13, -9, -10, -14, -15, -17,
	-12, -4, -7, -20, -21, -22, -23, -24, -25, -26,
	-27, -28, -29, -8, -5, -19, -11, 44, 43, 12,
	17, 18, 19, 20, 24, 25, 26, 27, 39, 41,
	45, 14, -43, -43, 33, -43, 33, -43, -43, -43,
	-43, 6, 7, 44, 6, 7, -43, -43, -43, -43,
	-43, -43, -43, -43, -43, -43, 4, 26, 27, 28,
	35, 36, 34, 6, -31, -41, 23, 8, -40, -2,
	-6, 6, 5, 11, 6, 23, 23, 23, 23, 6,
	7, 6, -18, -43, -18, -2, -2, -2, -2, -37,
	6, 5, -44, -2, -43, -43, -43, -43, -43, -43,
	-43, -43, -43, -43, -43, -43, -43, -30, -43, -43,
	-43, -43, -43, -43, -43, 30, 21, 7, -42, 30,
	-43, 30, -43, -30, -2, -2, -2, -2, -2, -2,
	-2, -31, 6, 5, 10, -2, -37, 10, 30, 29,
	-2, 13, -32, 34, -31, -2, -2, -2, -2, -34,
	-31, -30, -18, -43, 32, -30, 10, 40, -36, -39,
	-2, 6, 13, -33, 16, 15, -43, -43, -43, -43,
	-38, 6, -37, -45, 13, 30, 32, 13, -43, -43,
	-43, -43, 32, 13, 16, 15, -43, -30, -2, -6,
	9, 5, 6, -6, -46, 30, 34, -43, 29, 30,
	-30, -30, 7, 5, 6, -44, 10, -44, 10, 23,
	-43, -43, -30, -2, -30, 30, 10, -45, 13, -43,
	-43, -44, -43, -44, 22, -2, -2, -30, 13, -43,
	13, 42, -43, 42, -43, -43, -30, 6, -2, 6,
	-2, -43, 32, 23, -43, 22, -2, -43, 10, -2,
	10,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 6, 0, 12, 4, 12,
	16, -2, -2, -2, -2, -2, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, 42, 0, 0, 12,
	12, 12, 93, 94, 12, 12, 12, 12, 12, 12,
	113, 12, 13, 7, 0, 0, 0, 0, 0, 0,
	0, 109, 110, 0, 107, 108, 0, 12, 12, 0,
	0, 0, 0, 60, 14, 0, 12, 12, 12, 12,
	12, 12, 12, -2, 43, -2, 12, 12, -2, -2,
	52, -2, 16, -2, 12, 12, 12, 12, 12, 111,
	112, 12, 0, 0, 0, -2, -2, -2, -2, 12,
	61, 62, 12, -2, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 60, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 8, 12, -2, 0, 8,
	0, 15, 98, 0, -2, -2, -2, -2, -2, -2,
	-2, 44, 55, 56, 12, -2, 12, 12, 9, 10,
	-2, 114, 12, -2, 46, -2, -2, -2, -2, 65,
	70, 0, 0, 74, 0, 0, 12, 95, 12, 12,
	12, -2, 122, 0, 12, 8, 0, 0, 0, 0,
	0, -2, 12, 8, 71, 8, 0, 73, 0, 14,
	14, 0, 12, 123, 12, 8, 0, 128, -2, 54,
	49, 57, 58, 59, 8, 120, 116, 0, 66, 67,
	65, 0, -2, 63, 64, 12, 14, 12, 14, 0,
	0, 0, 126, -2, 0, 121, 12, 0, 72, 75,
	0, 12, 0, 12, 12, -2, -2, 127, 115, 0,
	68, 96, 0, 97, 0, 0, 125, -2, 12, 0,
	-2, 0, 12, 0, 0, 12, -2, 0, 104, -2,
	101,
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
		//line parser.y:138
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:140
		{
			Statements = []ast.Node{}
		}
	case 3:
		//line parser.y:142
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 4:
		//line parser.y:148
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 5:
		//line parser.y:154
		{
			RubyVAL.genericValue = nil
		}
	case 6:
		//line parser.y:155
		{
			RubyVAL.genericValue = nil
		}
	case 7:
		//line parser.y:156
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 8:
		//line parser.y:159
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 9:
		//line parser.y:161
		{ /* do nothing */
		}
	case 10:
		//line parser.y:163
		{ /* do nothing */
		}
	case 11:
		//line parser.y:165
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
		//line parser.y:177
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 44:
		//line parser.y:184
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 45:
		//line parser.y:192
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 46:
		//line parser.y:199
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 47:
		//line parser.y:207
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 48:
		//line parser.y:214
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 49:
		//line parser.y:223
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 50:
		//line parser.y:225
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 51:
		//line parser.y:228
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 52:
		//line parser.y:230
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 53:
		//line parser.y:232
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 54:
		//line parser.y:234
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 55:
		//line parser.y:237
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 56:
		//line parser.y:239
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 57:
		//line parser.y:241
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 58:
		//line parser.y:243
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 59:
		//line parser.y:245
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 60:
		//line parser.y:247
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 61:
		//line parser.y:249
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 62:
		//line parser.y:251
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 63:
		//line parser.y:253
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 64:
		//line parser.y:255
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:265
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 69:
		//line parser.y:273
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 70:
		//line parser.y:274
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 71:
		//line parser.y:277
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 72:
		//line parser.y:284
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 73:
		//line parser.y:294
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 74:
		//line parser.y:303
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 75:
		//line parser.y:309
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 76:
		//line parser.y:317
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 77:
		//line parser.y:321
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 78:
		//line parser.y:326
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 79:
		//line parser.y:333
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 80:
		//line parser.y:340
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 81:
		//line parser.y:347
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 82:
		//line parser.y:354
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 83:
		//line parser.y:361
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 84:
		//line parser.y:362
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 85:
		//line parser.y:363
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 86:
		//line parser.y:364
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 87:
		//line parser.y:367
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 88:
		//line parser.y:376
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 89:
		//line parser.y:385
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 90:
		//line parser.y:394
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 91:
		//line parser.y:403
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 92:
		//line parser.y:412
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 93:
		//line parser.y:420
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 94:
		//line parser.y:421
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 95:
		//line parser.y:423
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 96:
		//line parser.y:427
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 97:
		//line parser.y:435
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 98:
		//line parser.y:443
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 99:
		//line parser.y:445
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 100:
		//line parser.y:447
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 101:
		//line parser.y:449
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 102:
		//line parser.y:452
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 103:
		//line parser.y:459
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 104:
		//line parser.y:466
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 105:
		//line parser.y:474
		{
			RubyVAL.genericValue = nil
		}
	case 106:
		//line parser.y:475
		{
			RubyVAL.genericValue = nil
		}
	case 107:
		//line parser.y:478
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 108:
		//line parser.y:480
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 109:
		//line parser.y:483
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 110:
		//line parser.y:485
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 111:
		//line parser.y:488
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 112:
		//line parser.y:490
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 113:
		//line parser.y:493
		{
			RubyVAL.genericValue = ast.FileNameConstReference{}
		}
	case 114:
		//line parser.y:496
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 115:
		//line parser.y:498
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 116:
		//line parser.y:506
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 117:
		//line parser.y:508
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 118:
		//line parser.y:510
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 119:
		//line parser.y:512
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 122:
		//line parser.y:517
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 123:
		//line parser.y:524
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 124:
		//line parser.y:532
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 125:
		//line parser.y:534
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 126:
		//line parser.y:541
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 127:
		//line parser.y:548
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 128:
		//line parser.y:555
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	}
	goto Rubystack /* stack new state and value */
}
