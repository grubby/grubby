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
const UNLESS = 57359
const CLASS = 57360
const MODULE = 57361
const FOR = 57362
const WHILE = 57363
const UNTIL = 57364
const BEGIN = 57365
const RESCUE = 57366
const ENSURE = 57367
const BREAK = 57368
const REDO = 57369
const RETRY = 57370
const RETURN = 57371
const TRUE = 57372
const FALSE = 57373
const LESSTHAN = 57374
const GREATERTHAN = 57375
const EQUALTO = 57376
const BANG = 57377
const COMPLEMENT = 57378
const POSITIVE = 57379
const NEGATIVE = 57380
const STAR = 57381
const WHITESPACE = 57382
const NEWLINE = 57383
const SEMICOLON = 57384
const COLON = 57385
const DOT = 57386
const PIPE = 57387
const SLASH = 57388
const AMPERSAND = 57389
const MODULO = 57390
const CARET = 57391
const LBRACKET = 57392
const RBRACKET = 57393
const LBRACE = 57394
const RBRACE = 57395
const DOLLARSIGN = 57396
const ATSIGN = 57397
const FILE_CONST_REF = 57398
const EOF = 57399

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
	"UNLESS",
	"CLASS",
	"MODULE",
	"FOR",
	"WHILE",
	"UNTIL",
	"BEGIN",
	"RESCUE",
	"ENSURE",
	"BREAK",
	"REDO",
	"RETRY",
	"RETURN",
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

//line parser.y:782

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 10,
	43, 90,
	-2, 18,
	-1, 77,
	43, 90,
	-2, 18,
	-1, 82,
	8, 11,
	12, 11,
	14, 11,
	17, 11,
	18, 11,
	19, 11,
	23, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	42, 11,
	-2, 5,
	-1, 105,
	43, 90,
	-2, 88,
	-1, 205,
	43, 91,
	-2, 89,
}

const RubyNprod = 161
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1240

var RubyAct = []int{

	108, 166, 6, 74, 164, 31, 69, 68, 18, 118,
	73, 96, 139, 56, 2, 101, 157, 90, 91, 224,
	155, 140, 258, 204, 227, 162, 102, 157, 228, 153,
	259, 87, 221, 209, 58, 86, 85, 98, 84, 83,
	248, 218, 99, 268, 169, 66, 257, 66, 113, 104,
	106, 133, 219, 266, 223, 157, 66, 66, 157, 120,
	121, 156, 124, 125, 126, 127, 92, 226, 131, 207,
	154, 134, 135, 99, 57, 157, 99, 73, 237, 132,
	239, 238, 202, 8, 76, 77, 67, 88, 147, 81,
	89, 138, 142, 244, 155, 141, 234, 207, 185, 155,
	137, 179, 160, 7, 167, 261, 165, 243, 34, 35,
	87, 105, 71, 205, 168, 246, 170, 151, 152, 178,
	73, 167, 171, 172, 93, 94, 181, 182, 70, 195,
	82, 119, 33, 32, 186, 103, 198, 187, 189, 178,
	256, 191, 203, 178, 107, 109, 110, 111, 112, 193,
	117, 196, 38, 72, 159, 95, 199, 8, 76, 77,
	183, 194, 123, 81, 180, 158, 163, 128, 129, 178,
	208, 178, 178, 178, 200, 201, 1, 214, 211, 206,
	115, 55, 34, 35, 98, 143, 144, 145, 146, 99,
	149, 150, 54, 53, 178, 220, 52, 178, 51, 50,
	27, 26, 36, 25, 82, 24, 33, 32, 22, 233,
	21, 20, 30, 23, 28, 29, 232, 173, 229, 16,
	15, 236, 17, 14, 3, 0, 241, 0, 0, 242,
	0, 0, 0, 178, 58, 0, 0, 178, 0, 0,
	0, 0, 178, 178, 0, 0, 249, 0, 250, 253,
	254, 0, 0, 255, 178, 178, 178, 58, 0, 0,
	0, 178, 260, 0, 0, 178, 0, 59, 60, 61,
	65, 264, 0, 11, 57, 64, 62, 63, 265, 267,
	0, 188, 269, 78, 270, 0, 0, 0, 0, 0,
	59, 60, 61, 0, 0, 0, 0, 57, 64, 62,
	63, 0, 0, 0, 0, 230, 231, 0, 0, 0,
	78, 235, 19, 58, 12, 0, 0, 240, 184, 0,
	0, 78, 75, 245, 79, 0, 0, 0, 0, 247,
	0, 0, 78, 78, 0, 78, 78, 78, 78, 0,
	252, 78, 0, 0, 78, 78, 59, 60, 61, 100,
	78, 79, 0, 57, 64, 62, 63, 0, 0, 262,
	263, 78, 79, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 122, 79, 79, 78, 79, 79, 79, 79,
	100, 0, 79, 100, 136, 79, 79, 0, 0, 75,
	0, 79, 0, 78, 8, 161, 77, 0, 0, 0,
	148, 0, 79, 13, 0, 0, 0, 78, 0, 0,
	0, 78, 0, 80, 0, 0, 79, 0, 0, 34,
	35, 0, 0, 0, 0, 0, 0, 0, 0, 78,
	157, 0, 75, 0, 79, 0, 0, 0, 0, 36,
	80, 37, 0, 33, 32, 0, 0, 0, 79, 0,
	190, 80, 79, 0, 0, 0, 0, 78, 0, 0,
	0, 0, 80, 80, 0, 80, 80, 80, 80, 0,
	79, 80, 0, 0, 80, 80, 0, 0, 0, 0,
	80, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 80, 0, 8, 76, 77, 100, 130, 79, 8,
	9, 10, 48, 0, 0, 80, 39, 213, 46, 216,
	215, 47, 40, 41, 0, 0, 0, 49, 34, 35,
	0, 0, 0, 80, 34, 35, 0, 0, 0, 42,
	43, 44, 45, 0, 0, 176, 177, 80, 36, 0,
	37, 80, 33, 32, 36, 0, 37, 0, 33, 32,
	0, 0, 0, 8, 9, 10, 48, 0, 0, 80,
	39, 251, 46, 0, 0, 47, 40, 41, 0, 0,
	0, 49, 0, 0, 0, 0, 0, 0, 34, 35,
	0, 0, 0, 42, 43, 44, 45, 80, 0, 176,
	177, 0, 8, 9, 10, 48, 0, 0, 36, 39,
	37, 46, 33, 32, 47, 40, 41, 0, 0, 0,
	49, 0, 0, 0, 0, 0, 0, 34, 35, 0,
	0, 0, 42, 43, 44, 45, 0, 0, 176, 177,
	0, 0, 0, 0, 0, 0, 0, 36, 0, 37,
	225, 33, 32, 8, 9, 10, 48, 0, 0, 0,
	39, 222, 46, 0, 0, 47, 40, 41, 0, 0,
	0, 49, 0, 0, 0, 0, 0, 0, 34, 35,
	0, 0, 0, 42, 43, 44, 45, 0, 0, 176,
	177, 0, 0, 0, 0, 0, 0, 0, 36, 0,
	37, 0, 33, 32, 8, 9, 10, 48, 0, 0,
	0, 39, 217, 46, 0, 0, 47, 40, 41, 0,
	0, 0, 49, 0, 0, 0, 0, 0, 0, 34,
	35, 0, 0, 0, 42, 43, 44, 45, 0, 0,
	176, 177, 0, 0, 0, 0, 0, 0, 0, 36,
	0, 37, 0, 33, 32, 8, 9, 10, 48, 0,
	0, 0, 39, 212, 46, 0, 0, 47, 40, 41,
	0, 0, 0, 49, 0, 0, 0, 0, 0, 0,
	34, 35, 0, 0, 0, 42, 43, 44, 45, 0,
	0, 176, 177, 0, 0, 0, 0, 0, 0, 0,
	36, 0, 37, 0, 33, 32, 8, 9, 10, 48,
	0, 0, 0, 39, 210, 46, 0, 0, 47, 40,
	41, 0, 0, 0, 49, 0, 0, 0, 0, 0,
	0, 34, 35, 0, 0, 0, 42, 43, 44, 45,
	0, 0, 176, 177, 0, 8, 9, 10, 48, 0,
	0, 36, 39, 37, 46, 33, 32, 47, 40, 41,
	0, 0, 0, 49, 0, 0, 0, 0, 0, 0,
	34, 35, 0, 0, 0, 42, 43, 44, 45, 0,
	0, 176, 177, 0, 0, 0, 0, 0, 0, 0,
	36, 0, 37, 197, 33, 32, 8, 9, 10, 48,
	0, 0, 0, 39, 192, 46, 0, 0, 47, 40,
	41, 0, 0, 0, 49, 0, 0, 0, 0, 0,
	0, 34, 35, 0, 0, 0, 42, 43, 44, 45,
	0, 0, 176, 177, 0, 8, 9, 10, 48, 0,
	0, 36, 39, 37, 46, 33, 32, 47, 40, 41,
	0, 0, 0, 49, 0, 0, 0, 0, 0, 0,
	34, 35, 0, 0, 0, 42, 43, 44, 45, 0,
	0, 176, 177, 0, 8, 9, 10, 48, 175, 0,
	36, 39, 37, 46, 33, 32, 47, 40, 41, 0,
	0, 0, 49, 0, 0, 0, 0, 0, 0, 34,
	35, 0, 0, 0, 42, 43, 44, 45, 0, 0,
	0, 174, 0, 8, 9, 10, 48, 0, 0, 36,
	39, 37, 46, 33, 32, 47, 40, 41, 0, 0,
	0, 49, 0, 0, 0, 0, 0, 0, 34, 35,
	0, 0, 0, 42, 43, 44, 45, 0, 0, 4,
	5, 0, 8, 9, 10, 48, 0, 0, 36, 39,
	37, 46, 33, 32, 47, 40, 41, 0, 0, 0,
	49, 0, 0, 0, 0, 0, 0, 34, 35, 0,
	0, 0, 42, 43, 44, 45, 0, 0, 0, 116,
	0, 8, 9, 10, 48, 0, 0, 36, 39, 37,
	46, 33, 32, 47, 40, 41, 0, 0, 0, 49,
	0, 0, 58, 0, 0, 0, 34, 35, 0, 0,
	0, 42, 43, 44, 45, 8, 76, 77, 67, 0,
	0, 81, 0, 8, 76, 77, 36, 0, 37, 81,
	33, 32, 8, 114, 77, 59, 60, 61, 81, 0,
	34, 35, 57, 64, 62, 63, 0, 0, 34, 35,
	8, 76, 77, 0, 0, 0, 0, 34, 35, 0,
	70, 0, 82, 0, 33, 32, 0, 0, 36, 0,
	82, 0, 33, 32, 0, 34, 35, 36, 0, 82,
	0, 33, 32, 8, 114, 77, 0, 0, 0, 97,
	76, 77, 0, 0, 0, 36, 0, 37, 0, 33,
	32, 0, 0, 0, 0, 0, 0, 0, 34, 35,
	0, 0, 0, 0, 34, 35, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 36, 0,
	37, 0, 33, 32, 36, 0, 37, 0, 33, 32,
}
var RubyPact = []int{

	-43, 998, -1000, -44, -1000, -1000, 253, 31, -1000, 78,
	5, 4, 2, 1, -1000, -1000, -1000, -1000, -1000, 73,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 11, 118, -1000, -1000, 1184, -1000, -17, 129,
	104, 104, 1076, 1076, 1076, 1076, 1076, 1178, 1037, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 125, 1178, 1145,
	1076, 1178, 1178, 1178, 1178, 1076, 1076, 488, -1000, 41,
	1184, 1145, 90, 30, -1000, -1000, 1110, -1000, -1000, -1000,
	-1000, -24, -24, 1076, 1076, 1076, 1076, 1145, 1076, 1076,
	-1000, -1000, 111, -1000, -1000, 19, 10, -1000, 1098, -1000,
	-6, 389, -18, 98, 12, -1000, -1000, 31, 253, 31,
	31, 31, 31, -1000, -1000, 959, -1000, 31, 920, 152,
	-1000, 30, -1000, 31, -1000, -1000, -1000, -1000, 31, 31,
	-1000, 309, 89, 1127, 230, 30, -1000, 1118, 881, -1000,
	123, -1000, 830, 31, 31, 31, 31, 30, -1000, 31,
	31, -1000, -1000, 131, -1000, 1178, -1000, -1000, -1000, 72,
	138, -20, 106, -1000, 59, 115, -1000, -1, 791, 104,
	740, 494, 689, 31, -1000, -1000, -1000, -1000, 253, 31,
	28, 41, -1000, 1145, -1000, -1000, -1000, -1000, -2, 30,
	-1000, -1000, -1000, 638, 9, -1000, 587, -1000, -1000, -1000,
	14, -25, -1000, 1076, 1076, -1000, -14, 115, 87, 1076,
	-1000, -1000, -1000, -1000, 65, 1076, -1000, -1000, -1000, 100,
	84, 1076, -1000, -1000, 109, -1000, -1000, 1076, -1000, 34,
	31, 31, 920, -1000, -1000, 31, 548, -1000, 1076, -1000,
	31, 920, 920, 136, -1000, 31, -1000, 42, -21, -14,
	17, -1000, 31, 920, 920, 920, 99, 1076, 1076, -1000,
	920, -1000, 43, 33, 920, -14, -1000, -14, -1000, -14,
	-14,
}
var RubyPgo = []int{

	0, 101, 224, 223, 222, 3, 220, 219, 215, 403,
	214, 213, 212, 0, 314, 312, 211, 5, 210, 273,
	208, 8, 1, 205, 203, 201, 200, 199, 198, 196,
	193, 192, 181, 9, 180, 7, 12, 177, 176, 166,
	165, 164, 11, 161, 155, 154, 153, 6, 4, 152,
	15,
}
var RubyR1 = []int{

	0, 38, 38, 38, 38, 50, 50, 2, 2, 2,
	2, 33, 33, 33, 33, 33, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 17, 17, 17,
	17, 17, 17, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 35,
	35, 44, 44, 42, 42, 42, 42, 42, 47, 47,
	47, 47, 46, 46, 46, 46, 46, 16, 39, 39,
	22, 22, 48, 48, 48, 18, 18, 20, 21, 21,
	49, 49, 11, 11, 11, 11, 11, 11, 23, 24,
	25, 26, 27, 27, 27, 27, 28, 29, 30, 31,
	32, 3, 6, 7, 7, 4, 4, 40, 40, 40,
	40, 45, 45, 45, 9, 9, 19, 19, 14, 14,
	5, 5, 5, 5, 36, 43, 43, 43, 10, 10,
	10, 10, 10, 10, 10, 37, 37, 37, 37, 37,
	34, 34, 34, 34, 34, 8, 12, 41, 41, 41,
	41,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 0, 2, 1, 1, 1,
	1, 0, 2, 2, 2, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 4, 4,
	4, 2, 3, 4, 4, 2, 3, 4, 6, 3,
	1, 1, 3, 0, 1, 1, 1, 3, 1, 1,
	3, 3, 1, 1, 3, 3, 3, 7, 1, 3,
	1, 3, 0, 1, 3, 4, 6, 4, 1, 4,
	1, 4, 3, 3, 3, 3, 3, 3, 2, 2,
	2, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 1, 1, 3, 3, 5, 5, 0, 4, 7,
	8, 3, 7, 8, 2, 2, 2, 2, 3, 3,
	3, 4, 4, 3, 3, 0, 1, 3, 4, 5,
	3, 3, 3, 3, 4, 0, 4, 3, 3, 2,
	0, 1, 1, 2, 2, 3, 4, 0, 3, 4,
	6,
}
var RubyChk = []int{

	-1000, -38, 57, -2, 41, 42, -13, -1, 5, 6,
	7, -19, -14, -9, -3, -6, -7, -4, -21, -15,
	-16, -18, -20, -11, -23, -24, -25, -26, -10, -8,
	-12, -17, 55, 54, 30, 31, 50, 52, -49, 12,
	18, 19, 35, 36, 37, 38, 14, 17, 8, 23,
	-27, -28, -29, -30, -31, -32, 57, 44, 4, 37,
	38, 39, 46, 47, 45, 17, 14, 8, -35, -47,
	50, 34, -46, -13, -5, -15, 6, 7, -19, -14,
	-9, 11, 52, 34, 34, 34, 34, 37, 14, 17,
	6, 7, 55, 6, 7, -44, -42, 5, -13, -17,
	-15, -50, 43, 6, -21, 7, -21, -1, -13, -1,
	-1, -1, -1, -13, 6, -34, 42, -1, -33, 6,
	-13, -13, -15, -1, -13, -13, -13, -13, -1, -1,
	9, -13, -42, 10, -13, -13, -15, 10, -33, -36,
	45, -36, -33, -1, -1, -1, -1, -13, -15, -1,
	-1, 6, 7, 10, 51, 10, 51, 41, -40, -45,
	-13, 6, 43, -39, -48, 8, -22, 6, -33, 32,
	-33, -33, -33, -1, 42, 9, 41, 42, -13, -1,
	-41, -47, -35, 8, 9, 9, -13, -5, 51, -13,
	-15, -5, 13, -33, -43, 6, -33, 53, 5, -13,
	-50, -50, 10, 4, 43, 7, -50, 10, -48, 34,
	13, -21, 13, 13, -37, 16, 15, 13, 13, 24,
	-42, 34, 13, 45, 10, 53, 53, 10, 53, -50,
	-1, -1, -33, -22, 9, -1, -33, 13, 16, 15,
	-1, -33, -33, 7, 9, -1, 6, -1, 6, -50,
	-50, 13, -1, -33, -33, -33, 4, 4, 43, 13,
	-33, 6, -1, -1, -33, -50, 10, -50, 10, -50,
	-50,
}
var RubyDef = []int{

	1, -2, 2, 3, 7, 8, 9, 10, 16, 17,
	-2, 19, 20, 21, 22, 23, 24, 25, 26, 34,
	35, 36, 37, 38, 39, 40, 41, 42, 43, 44,
	45, 46, 0, 0, 111, 112, 63, 5, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 150, 11,
	27, 28, 29, 30, 31, 32, 4, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 63, 51, 55,
	63, 0, 60, 68, 69, 73, 17, -2, 19, 20,
	21, 11, -2, 0, 0, 0, 0, 0, 0, 0,
	126, 127, 0, 124, 125, 0, 0, 16, 64, 65,
	66, 117, 0, 82, 11, -2, 11, 98, 33, 99,
	100, 101, 11, 11, 17, 0, 151, 152, 157, 52,
	56, 102, 104, 106, 107, 108, 109, 110, 142, 140,
	47, 64, 0, 0, 64, 92, 93, 0, 0, 11,
	135, 11, 0, 94, 95, 96, 97, 103, 105, 141,
	143, 128, 129, 0, 113, 0, 114, 6, 5, 5,
	0, 17, 0, 5, 78, 82, 83, 80, 0, 0,
	0, 0, 0, 153, 154, 155, 12, 13, 14, 15,
	0, 53, 54, 63, 48, 49, 70, 71, 57, 74,
	75, 76, 130, 0, 0, 136, 0, 133, 62, 67,
	0, 0, 5, 0, 0, -2, 11, 0, 0, 0,
	85, 11, 87, 138, 0, 0, 11, 144, 156, 11,
	0, 0, 131, 134, 0, 132, 115, 0, 116, 0,
	5, 121, 5, 84, 79, 81, 0, 139, 0, 11,
	11, 149, 158, 11, 59, 58, 137, 0, 0, 118,
	0, 86, 11, 147, 148, 159, 0, 0, 0, 77,
	146, 11, 5, 5, 160, 119, 5, 122, 5, 120,
	123,
}
var RubyTok1 = []int{

	1,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57,
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
		//line parser.y:158
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:160
		{
			Statements = []ast.Node{}
		}
	case 3:
		//line parser.y:162
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 4:
		//line parser.y:168
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 5:
		//line parser.y:174
		{
		}
	case 6:
		//line parser.y:175
		{
		}
	case 7:
		//line parser.y:177
		{
			RubyVAL.genericValue = nil
		}
	case 8:
		//line parser.y:178
		{
			RubyVAL.genericValue = nil
		}
	case 9:
		//line parser.y:179
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 10:
		//line parser.y:180
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 11:
		//line parser.y:183
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 12:
		//line parser.y:185
		{
		}
	case 13:
		//line parser.y:187
		{
		}
	case 14:
		//line parser.y:189
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 15:
		//line parser.y:191
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 45:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 46:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 47:
		//line parser.y:201
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: []ast.Node{},
			}
		}
	case 48:
		//line parser.y:208
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 49:
		//line parser.y:215
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 50:
		//line parser.y:222
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 51:
		//line parser.y:229
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 52:
		//line parser.y:236
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 53:
		//line parser.y:243
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 54:
		//line parser.y:251
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 55:
		//line parser.y:261
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 56:
		//line parser.y:268
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 57:
		//line parser.y:278
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 58:
		//line parser.y:288
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 59:
		//line parser.y:297
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 60:
		//line parser.y:299
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 61:
		//line parser.y:302
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 62:
		//line parser.y:304
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 63:
		//line parser.y:306
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 64:
		//line parser.y:308
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 65:
		//line parser.y:310
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 66:
		//line parser.y:312
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:314
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:318
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 69:
		//line parser.y:320
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 70:
		//line parser.y:322
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 71:
		//line parser.y:324
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 72:
		//line parser.y:327
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:329
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:331
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 75:
		//line parser.y:333
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:335
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:341
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 78:
		//line parser.y:350
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 79:
		//line parser.y:352
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 80:
		//line parser.y:355
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 81:
		//line parser.y:357
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 82:
		//line parser.y:359
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 83:
		//line parser.y:361
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 84:
		//line parser.y:365
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:370
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 86:
		//line parser.y:377
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 87:
		//line parser.y:387
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 88:
		//line parser.y:396
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 89:
		//line parser.y:402
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 90:
		//line parser.y:410
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 91:
		//line parser.y:414
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 92:
		//line parser.y:419
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 93:
		//line parser.y:426
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 94:
		//line parser.y:433
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 95:
		//line parser.y:440
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 96:
		//line parser.y:447
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 97:
		//line parser.y:454
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 98:
		//line parser.y:461
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 99:
		//line parser.y:462
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 100:
		//line parser.y:463
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 101:
		//line parser.y:464
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 102:
		//line parser.y:467
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 103:
		//line parser.y:475
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 104:
		//line parser.y:483
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 105:
		//line parser.y:491
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 106:
		//line parser.y:500
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 107:
		//line parser.y:509
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 108:
		//line parser.y:518
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 109:
		//line parser.y:527
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 110:
		//line parser.y:536
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 111:
		//line parser.y:544
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 112:
		//line parser.y:545
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 113:
		//line parser.y:547
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 114:
		//line parser.y:549
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 115:
		//line parser.y:552
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 116:
		//line parser.y:560
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 117:
		//line parser.y:568
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 118:
		//line parser.y:570
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 119:
		//line parser.y:577
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 120:
		//line parser.y:584
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 121:
		//line parser.y:592
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 122:
		//line parser.y:599
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 123:
		//line parser.y:606
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 124:
		//line parser.y:614
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 125:
		//line parser.y:616
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 126:
		//line parser.y:619
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 127:
		//line parser.y:621
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 128:
		//line parser.y:624
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 129:
		//line parser.y:626
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 130:
		//line parser.y:629
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 131:
		//line parser.y:631
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 132:
		//line parser.y:638
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 133:
		//line parser.y:645
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 134:
		//line parser.y:648
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 135:
		//line parser.y:650
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 136:
		//line parser.y:652
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 137:
		//line parser.y:654
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 138:
		//line parser.y:657
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 139:
		//line parser.y:664
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 140:
		//line parser.y:672
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 141:
		//line parser.y:679
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 142:
		//line parser.y:686
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 143:
		//line parser.y:693
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 144:
		//line parser.y:700
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-2].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 145:
		//line parser.y:707
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 146:
		//line parser.y:709
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 147:
		//line parser.y:716
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 148:
		//line parser.y:723
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 149:
		//line parser.y:730
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 150:
		//line parser.y:737
		{
		}
	case 151:
		//line parser.y:738
		{
		}
	case 152:
		//line parser.y:739
		{
			RubyVAL.genericSlice = []ast.Node{RubyS[Rubypt-0].genericValue}
		}
	case 153:
		//line parser.y:740
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 154:
		//line parser.y:741
		{
		}
	case 155:
		//line parser.y:744
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 156:
		//line parser.y:747
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 157:
		//line parser.y:755
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 158:
		//line parser.y:757
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 159:
		//line parser.y:759
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 160:
		//line parser.y:768
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}

			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Var:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
					Class: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				},
			})
		}
	}
	goto Rubystack /* stack new state and value */
}
