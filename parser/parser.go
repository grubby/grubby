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
const CARET = 57390
const LBRACKET = 57391
const RBRACKET = 57392
const LBRACE = 57393
const RBRACE = 57394
const DOLLARSIGN = 57395
const ATSIGN = 57396
const FILE_CONST_REF = 57397
const EOF = 57398

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

//line parser.y:807

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 25,
	43, 93,
	-2, 20,
	-1, 76,
	10, 65,
	-2, 157,
	-1, 87,
	43, 93,
	-2, 20,
	-1, 92,
	8, 13,
	12, 13,
	14, 13,
	17, 13,
	18, 13,
	19, 13,
	23, 13,
	35, 13,
	36, 13,
	37, 13,
	38, 13,
	42, 13,
	-2, 11,
	-1, 101,
	43, 93,
	-2, 91,
	-1, 196,
	9, 65,
	10, 65,
	-2, 157,
	-1, 233,
	43, 94,
	-2, 92,
}

const RubyNprod = 166
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1303

var RubyAct = []int{

	9, 120, 33, 162, 84, 153, 164, 161, 78, 77,
	188, 2, 3, 62, 114, 115, 186, 238, 64, 254,
	154, 255, 75, 184, 278, 83, 4, 232, 22, 193,
	126, 188, 271, 72, 235, 214, 274, 97, 100, 102,
	60, 59, 168, 224, 110, 96, 273, 10, 62, 95,
	188, 104, 237, 94, 225, 61, 187, 122, 63, 20,
	64, 253, 116, 185, 188, 129, 130, 188, 133, 134,
	135, 136, 85, 138, 188, 103, 73, 144, 145, 74,
	147, 148, 285, 252, 93, 123, 246, 83, 248, 247,
	259, 186, 283, 65, 66, 67, 62, 146, 155, 72,
	63, 70, 68, 69, 124, 123, 201, 166, 123, 242,
	212, 230, 277, 131, 179, 212, 93, 151, 233, 180,
	139, 5, 62, 101, 124, 93, 191, 124, 149, 83,
	198, 186, 93, 165, 85, 163, 93, 194, 195, 93,
	93, 182, 183, 117, 118, 276, 93, 199, 99, 261,
	98, 200, 202, 179, 165, 208, 204, 179, 105, 106,
	107, 108, 109, 171, 150, 142, 128, 213, 179, 226,
	179, 217, 269, 179, 179, 231, 85, 58, 23, 86,
	87, 44, 127, 82, 91, 190, 132, 227, 93, 119,
	207, 137, 181, 140, 141, 64, 189, 122, 234, 203,
	1, 220, 125, 54, 55, 112, 51, 179, 71, 50,
	179, 93, 49, 48, 157, 158, 159, 160, 26, 241,
	47, 46, 56, 170, 92, 123, 53, 52, 65, 66,
	67, 18, 174, 258, 17, 63, 70, 68, 69, 16,
	15, 179, 37, 88, 124, 179, 179, 13, 12, 27,
	11, 179, 179, 21, 14, 19, 93, 31, 30, 32,
	29, 0, 88, 0, 0, 0, 0, 179, 179, 179,
	0, 0, 0, 0, 89, 88, 179, 0, 0, 281,
	179, 0, 0, 88, 88, 0, 88, 88, 88, 88,
	0, 88, 0, 89, 0, 88, 0, 0, 88, 88,
	0, 0, 0, 0, 0, 88, 89, 0, 0, 0,
	0, 0, 0, 0, 89, 89, 0, 89, 89, 89,
	89, 0, 89, 0, 0, 0, 89, 0, 0, 89,
	89, 0, 0, 0, 243, 0, 89, 0, 0, 0,
	0, 249, 0, 28, 88, 0, 0, 88, 0, 0,
	0, 257, 113, 0, 0, 260, 0, 0, 23, 86,
	87, 44, 143, 0, 211, 88, 0, 265, 90, 215,
	88, 0, 0, 0, 270, 89, 0, 0, 89, 0,
	0, 0, 0, 54, 55, 0, 0, 90, 0, 0,
	0, 0, 228, 229, 0, 0, 89, 280, 0, 0,
	90, 89, 56, 0, 57, 88, 53, 52, 90, 90,
	0, 90, 90, 90, 90, 88, 90, 0, 0, 0,
	90, 0, 0, 90, 90, 0, 0, 0, 0, 0,
	90, 0, 0, 256, 0, 0, 89, 0, 0, 0,
	0, 0, 0, 262, 152, 156, 89, 263, 0, 0,
	0, 88, 0, 167, 0, 169, 0, 0, 0, 64,
	272, 0, 172, 173, 197, 0, 0, 0, 0, 90,
	0, 0, 90, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 89, 282, 284, 0, 286, 0, 287, 0,
	90, 0, 65, 66, 67, 90, 0, 88, 0, 63,
	70, 68, 69, 0, 0, 0, 206, 0, 209, 0,
	0, 23, 24, 25, 44, 0, 0, 0, 34, 219,
	42, 222, 221, 43, 35, 36, 0, 0, 89, 45,
	90, 0, 0, 0, 0, 0, 54, 55, 0, 64,
	90, 38, 39, 40, 41, 0, 0, 177, 178, 0,
	0, 0, 0, 0, 0, 56, 0, 57, 0, 53,
	52, 0, 0, 0, 240, 0, 0, 0, 244, 0,
	245, 0, 65, 66, 67, 250, 90, 0, 251, 63,
	70, 68, 69, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 266, 267, 0, 0, 268, 0, 23, 24, 25,
	44, 0, 0, 0, 34, 0, 42, 0, 275, 43,
	35, 36, 90, 0, 0, 45, 0, 0, 0, 279,
	0, 0, 54, 55, 0, 0, 0, 38, 39, 40,
	41, 0, 0, 6, 7, 0, 0, 0, 0, 0,
	0, 56, 0, 57, 0, 53, 52, 0, 8, 23,
	24, 25, 44, 0, 0, 0, 34, 264, 42, 0,
	0, 43, 35, 36, 0, 0, 0, 45, 0, 0,
	0, 0, 0, 0, 54, 55, 0, 0, 0, 38,
	39, 40, 41, 0, 0, 177, 178, 23, 24, 25,
	44, 0, 0, 56, 34, 57, 42, 53, 52, 43,
	35, 36, 0, 0, 0, 45, 0, 0, 0, 0,
	0, 0, 54, 55, 0, 0, 0, 38, 39, 40,
	41, 0, 0, 177, 178, 0, 0, 0, 0, 0,
	0, 56, 0, 57, 239, 53, 52, 23, 24, 25,
	44, 0, 0, 0, 34, 236, 42, 0, 0, 43,
	35, 36, 0, 0, 0, 45, 0, 0, 0, 0,
	0, 0, 54, 55, 0, 0, 0, 38, 39, 40,
	41, 0, 0, 177, 178, 0, 0, 0, 0, 0,
	0, 56, 0, 57, 0, 53, 52, 23, 24, 25,
	44, 0, 0, 0, 34, 223, 42, 0, 0, 43,
	35, 36, 0, 0, 0, 45, 0, 0, 0, 0,
	0, 0, 54, 55, 0, 0, 0, 38, 39, 40,
	41, 0, 0, 177, 178, 0, 0, 0, 0, 0,
	0, 56, 0, 57, 0, 53, 52, 23, 24, 25,
	44, 0, 0, 0, 34, 218, 42, 0, 0, 43,
	35, 36, 0, 0, 0, 45, 0, 0, 0, 0,
	0, 0, 54, 55, 0, 0, 0, 38, 39, 40,
	41, 0, 0, 177, 178, 0, 0, 0, 0, 0,
	0, 56, 0, 57, 0, 53, 52, 23, 24, 25,
	44, 0, 0, 0, 34, 216, 42, 0, 0, 43,
	35, 36, 0, 0, 0, 45, 0, 0, 0, 0,
	0, 0, 54, 55, 0, 0, 0, 38, 39, 40,
	41, 0, 0, 177, 178, 23, 24, 25, 44, 0,
	0, 56, 34, 57, 42, 53, 52, 43, 35, 36,
	0, 0, 0, 45, 0, 0, 0, 0, 0, 0,
	54, 55, 0, 0, 0, 38, 39, 40, 41, 0,
	0, 177, 178, 0, 0, 0, 0, 0, 0, 56,
	0, 57, 210, 53, 52, 23, 24, 25, 44, 0,
	0, 0, 34, 205, 42, 0, 0, 43, 35, 36,
	0, 0, 0, 45, 0, 0, 0, 0, 0, 0,
	54, 55, 0, 0, 0, 38, 39, 40, 41, 0,
	0, 177, 178, 23, 24, 25, 44, 0, 0, 56,
	34, 57, 42, 53, 52, 43, 35, 36, 0, 0,
	0, 45, 0, 0, 0, 0, 0, 0, 54, 55,
	0, 0, 0, 38, 39, 40, 41, 0, 0, 177,
	178, 23, 24, 25, 44, 176, 0, 56, 34, 57,
	42, 53, 52, 43, 35, 36, 0, 0, 0, 45,
	0, 0, 0, 0, 0, 0, 54, 55, 0, 0,
	0, 38, 39, 40, 41, 0, 0, 0, 175, 23,
	24, 25, 44, 0, 0, 56, 34, 57, 42, 53,
	52, 43, 35, 36, 0, 0, 0, 45, 0, 0,
	0, 0, 0, 0, 54, 55, 0, 0, 0, 38,
	39, 40, 41, 23, 86, 87, 76, 0, 81, 91,
	0, 0, 0, 56, 0, 57, 0, 53, 52, 23,
	86, 87, 196, 0, 0, 91, 0, 0, 54, 55,
	0, 0, 80, 23, 86, 87, 76, 0, 0, 91,
	0, 0, 0, 0, 54, 55, 0, 79, 0, 92,
	0, 53, 52, 0, 0, 0, 0, 0, 54, 55,
	23, 111, 87, 56, 0, 92, 91, 53, 52, 23,
	86, 87, 44, 0, 0, 0, 0, 79, 0, 92,
	0, 53, 52, 0, 0, 54, 55, 0, 0, 0,
	0, 0, 0, 0, 54, 55, 23, 192, 87, 0,
	0, 0, 0, 0, 56, 0, 92, 0, 53, 52,
	23, 111, 87, 56, 0, 57, 0, 53, 52, 0,
	0, 54, 55, 121, 86, 87, 44, 0, 0, 0,
	0, 0, 188, 0, 0, 54, 55, 0, 0, 0,
	56, 0, 57, 0, 53, 52, 0, 0, 54, 55,
	0, 0, 0, 0, 56, 0, 57, 0, 53, 52,
	0, 0, 0, 0, 0, 0, 0, 56, 0, 57,
	0, 53, 52,
}
var RubyPact = []int{

	-30, 602, -1000, -1000, -1000, -1, -1000, -1000, -1000, 191,
	62, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-22, -1000, -1000, -1000, 1128, 19, 15, 11, 3, -1000,
	-1000, -1000, -1000, -1000, 144, 116, 116, 41, 1094, 1094,
	1094, 1094, 1094, 1235, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 8, 137, -1000, -1000, 1248, -1000, -13, -1000,
	-1000, -1000, 1094, 160, 1235, 1194, 1094, 1235, 1235, 1235,
	1235, 1094, 1194, 1094, 1094, 159, 353, -1000, 87, 1248,
	1194, 158, 107, 14, -1000, -1000, 1158, -1000, -1000, -1000,
	-1000, -25, -25, -22, 1094, 1094, 1094, 1094, 127, 127,
	10, -1000, -1000, 1094, 157, 34, 34, 34, 34, 34,
	-1000, -1000, 1056, 1018, -1000, -1000, 135, -1000, -1000, 13,
	6, -1000, 535, -1000, -4, 1221, -14, 34, 1144, -1000,
	14, -1000, 34, -1000, -1000, -1000, -1000, 34, 14, -1000,
	34, 34, -1000, -1000, 455, 121, 1185, 56, 14, -1000,
	-1000, 173, 980, -1000, 149, -1000, 930, 34, 34, 34,
	34, -1000, 105, 148, -1000, 1, -1000, 892, 116, 842,
	34, -1000, 506, 792, 34, -1000, -1000, -1000, -1000, 191,
	34, 30, -1000, -1000, 164, -1000, 1235, -1000, -1000, -1000,
	101, 171, -16, 111, 87, -1000, 1194, -1000, -1000, -1000,
	-1000, 0, 14, -1000, -1000, -1000, 742, 7, -1000, 692,
	-1000, -10, 148, 100, 1094, -10, -1000, -1000, -1000, -1000,
	73, 1094, -1000, -1000, -1000, 76, -1000, -1000, 9, -31,
	-1000, 1094, 1235, -1000, 81, 1094, -1000, -1000, 143, -1000,
	1018, -1000, -1000, 34, 1018, 654, -1000, 1094, -1000, 34,
	1018, 1018, 168, -1000, 1094, -1000, 26, 34, -1000, -1000,
	34, -1000, 33, 23, -1000, 34, 1018, 1018, 1018, 139,
	108, -19, -10, -1000, -1000, 1018, -1000, 1094, 1235, 1018,
	82, 72, -10, -1000, -10, -1000, -10, -10,
}
var RubyPgo = []int{

	0, 119, 260, 259, 4, 258, 257, 59, 343, 255,
	254, 253, 0, 249, 47, 250, 28, 248, 218, 247,
	2, 6, 242, 240, 239, 234, 231, 221, 220, 213,
	212, 209, 206, 352, 205, 9, 5, 201, 200, 7,
	196, 192, 1, 190, 189, 185, 183, 8, 3, 177,
	202,
}
var RubyR1 = []int{

	0, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 50, 50, 33, 33, 33, 33, 33, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 16,
	16, 16, 16, 16, 16, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 35, 35, 44, 44, 42, 42, 42, 42, 42,
	47, 47, 47, 47, 46, 46, 46, 46, 46, 15,
	15, 39, 39, 21, 21, 48, 48, 48, 17, 17,
	19, 20, 20, 49, 49, 10, 10, 10, 10, 10,
	10, 10, 8, 8, 18, 18, 13, 13, 22, 22,
	23, 24, 25, 26, 27, 27, 27, 27, 28, 29,
	30, 31, 32, 2, 5, 6, 6, 3, 3, 40,
	40, 40, 40, 45, 45, 45, 45, 4, 4, 4,
	4, 36, 43, 43, 43, 9, 9, 9, 9, 9,
	9, 9, 37, 37, 37, 37, 37, 34, 34, 34,
	7, 11, 41, 41, 41, 41,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
	4, 4, 2, 3, 4, 4, 3, 2, 3, 4,
	6, 3, 1, 1, 3, 0, 1, 1, 1, 3,
	1, 1, 3, 3, 1, 1, 3, 3, 3, 7,
	7, 1, 3, 1, 3, 0, 1, 3, 4, 6,
	4, 1, 4, 1, 4, 3, 3, 3, 3, 3,
	3, 3, 2, 2, 2, 2, 3, 3, 3, 3,
	2, 2, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 1, 1, 3, 3, 5, 5, 0,
	4, 7, 8, 3, 3, 7, 8, 3, 4, 4,
	3, 3, 0, 1, 3, 4, 5, 3, 3, 3,
	3, 4, 0, 4, 3, 3, 2, 0, 2, 2,
	3, 4, 0, 3, 4, 6,
}
var RubyChk = []int{

	-1000, -38, 41, 42, 56, -1, 41, 42, 56, -12,
	-14, -15, -17, -19, -10, -23, -24, -25, -26, -9,
	-7, -11, -16, 5, 6, 7, -18, -13, -8, -2,
	-5, -6, -3, -20, 12, 18, 19, -22, 35, 36,
	37, 38, 14, 17, 8, 23, -27, -28, -29, -30,
	-31, -32, 54, 53, 30, 31, 49, 51, -49, 42,
	41, 56, 14, 44, 4, 37, 38, 39, 46, 47,
	45, 17, 37, 14, 17, 44, 8, -35, -47, 49,
	34, 10, -46, -12, -4, -14, 6, 7, -18, -13,
	-8, 11, 51, -7, 34, 34, 34, 34, 6, 4,
	-20, 7, -20, 34, 10, -1, -1, -1, -1, -1,
	-12, 6, -34, -33, 6, 7, 54, 6, 7, -44,
	-42, 5, -12, -16, -14, -50, 43, -1, 6, -12,
	-12, -14, -1, -12, -12, -12, -12, -1, -12, -14,
	-1, -1, 6, 9, -12, -42, 10, -12, -12, -14,
	6, 10, -33, -36, 45, -36, -33, -1, -1, -1,
	-1, -39, -48, 8, -21, 6, -39, -33, 32, -33,
	-1, 6, -33, -33, -1, 42, 9, 41, 42, -12,
	-1, -41, 6, 7, 10, 50, 10, 50, 41, -40,
	-45, -12, 6, 43, -47, -35, 8, 9, 9, -12,
	-4, 50, -12, -14, -4, 13, -33, -43, 6, -33,
	52, -50, 10, -48, 34, -50, 13, -20, 13, 13,
	-37, 16, 15, 13, 13, 24, 5, -12, -50, -50,
	10, 4, 43, 7, -42, 34, 13, 45, 10, 52,
	-33, -21, 9, -1, -33, -33, 13, 16, 15, -1,
	-33, -33, 7, 52, 10, 52, -50, -1, -12, 9,
	-1, 6, -50, -50, 13, -1, -33, -33, -33, 4,
	-1, 6, -50, 13, 13, -33, 6, 4, 43, -33,
	-1, -12, -50, 10, -50, 10, -50, -50,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
	46, 47, 48, 18, 19, -2, 21, 22, 23, 24,
	25, 26, 27, 28, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 157, 13, 29, 30, 31, 32,
	33, 34, 0, 0, 123, 124, 65, 11, 0, 5,
	6, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, -2, 52, 57, 65,
	0, 0, 62, 70, 71, 75, 19, -2, 21, 22,
	23, 13, -2, 0, 0, 0, 0, 0, 85, 85,
	13, -2, 13, 0, 0, 110, 111, 112, 113, 13,
	13, 19, 0, 162, 104, 105, 0, 102, 103, 0,
	0, 18, 66, 67, 68, 129, 0, 147, 53, 58,
	114, 116, 118, 119, 120, 121, 122, 149, 115, 117,
	148, 150, 56, 49, 66, 0, 0, 66, 95, 96,
	108, 0, 0, 13, 142, 13, 0, 97, 98, 99,
	100, 11, 81, 85, 86, 83, 11, 0, 0, 0,
	101, 109, 0, 0, 158, 159, 160, 14, 15, 16,
	17, 0, 106, 107, 0, 125, 0, 126, 12, 11,
	11, 0, 19, 0, 54, 55, -2, 50, 51, 72,
	73, 59, 76, 77, 78, 137, 0, 0, 143, 0,
	140, 13, 0, 0, 0, 13, 88, 13, 90, 145,
	0, 0, 13, 151, 161, 13, 64, 69, 0, 0,
	11, 0, 0, -2, 0, 0, 138, 141, 0, 139,
	11, 87, 82, 84, 11, 0, 146, 0, 13, 13,
	156, 163, 13, 127, 0, 128, 0, 11, 133, 61,
	60, 144, 0, 0, 89, 13, 154, 155, 164, 0,
	0, 0, 130, 79, 80, 153, 13, 0, 0, 165,
	11, 11, 131, 11, 135, 11, 132, 136,
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
	52, 53, 54, 55, 56,
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
		}
	case 3:
		//line parser.y:162
		{
		}
	case 4:
		//line parser.y:164
		{
		}
	case 5:
		//line parser.y:166
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:168
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:170
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:176
		{
		}
	case 11:
		//line parser.y:178
		{
		}
	case 12:
		//line parser.y:179
		{
		}
	case 13:
		//line parser.y:182
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 14:
		//line parser.y:184
		{
		}
	case 15:
		//line parser.y:186
		{
		}
	case 16:
		//line parser.y:188
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 17:
		//line parser.y:190
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 48:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 49:
		//line parser.y:200
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: []ast.Node{},
			}
		}
	case 50:
		//line parser.y:207
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 51:
		//line parser.y:214
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 52:
		//line parser.y:221
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 53:
		//line parser.y:228
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 54:
		//line parser.y:235
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 55:
		//line parser.y:243
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 56:
		//line parser.y:251
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 57:
		//line parser.y:261
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 58:
		//line parser.y:268
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 59:
		//line parser.y:278
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 60:
		//line parser.y:288
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 61:
		//line parser.y:297
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 62:
		//line parser.y:299
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 63:
		//line parser.y:302
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 64:
		//line parser.y:304
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 65:
		//line parser.y:306
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 66:
		//line parser.y:308
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:310
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:312
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 69:
		//line parser.y:314
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 70:
		//line parser.y:318
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 71:
		//line parser.y:320
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 72:
		//line parser.y:322
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:324
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:327
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 75:
		//line parser.y:329
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:331
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:333
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 78:
		//line parser.y:335
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 79:
		//line parser.y:341
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 80:
		//line parser.y:349
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 81:
		//line parser.y:358
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 82:
		//line parser.y:360
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 83:
		//line parser.y:363
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 84:
		//line parser.y:365
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 85:
		//line parser.y:367
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 86:
		//line parser.y:369
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 87:
		//line parser.y:373
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 88:
		//line parser.y:378
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 89:
		//line parser.y:385
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 90:
		//line parser.y:395
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 91:
		//line parser.y:404
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 92:
		//line parser.y:410
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 93:
		//line parser.y:418
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 94:
		//line parser.y:422
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 95:
		//line parser.y:427
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 96:
		//line parser.y:434
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 97:
		//line parser.y:441
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 98:
		//line parser.y:448
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 99:
		//line parser.y:455
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 100:
		//line parser.y:462
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 101:
		//line parser.y:469
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 102:
		//line parser.y:477
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 103:
		//line parser.y:479
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 104:
		//line parser.y:482
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 105:
		//line parser.y:484
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 106:
		//line parser.y:487
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 107:
		//line parser.y:489
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 108:
		//line parser.y:492
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 109:
		//line parser.y:494
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 110:
		//line parser.y:496
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 111:
		//line parser.y:497
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 112:
		//line parser.y:498
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 113:
		//line parser.y:499
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 114:
		//line parser.y:502
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 115:
		//line parser.y:510
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 116:
		//line parser.y:518
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 117:
		//line parser.y:526
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 118:
		//line parser.y:535
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 119:
		//line parser.y:544
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 120:
		//line parser.y:553
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 121:
		//line parser.y:562
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 122:
		//line parser.y:571
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 123:
		//line parser.y:579
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 124:
		//line parser.y:580
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 125:
		//line parser.y:582
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 126:
		//line parser.y:584
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 127:
		//line parser.y:587
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 128:
		//line parser.y:595
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 129:
		//line parser.y:603
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 130:
		//line parser.y:605
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 131:
		//line parser.y:612
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 132:
		//line parser.y:619
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 133:
		//line parser.y:627
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 134:
		//line parser.y:634
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 135:
		//line parser.y:641
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 136:
		//line parser.y:648
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 137:
		//line parser.y:656
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 138:
		//line parser.y:658
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 139:
		//line parser.y:665
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 140:
		//line parser.y:672
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 141:
		//line parser.y:675
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 142:
		//line parser.y:677
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 143:
		//line parser.y:679
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 144:
		//line parser.y:681
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 145:
		//line parser.y:684
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 146:
		//line parser.y:691
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 147:
		//line parser.y:699
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 148:
		//line parser.y:706
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 149:
		//line parser.y:713
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 150:
		//line parser.y:720
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 151:
		//line parser.y:727
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-2].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 152:
		//line parser.y:734
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 153:
		//line parser.y:736
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 154:
		//line parser.y:743
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 155:
		//line parser.y:750
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 156:
		//line parser.y:757
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 157:
		//line parser.y:764
		{
		}
	case 158:
		//line parser.y:765
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 159:
		//line parser.y:766
		{
		}
	case 160:
		//line parser.y:769
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 161:
		//line parser.y:772
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 162:
		//line parser.y:780
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 163:
		//line parser.y:782
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 164:
		//line parser.y:784
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 165:
		//line parser.y:793
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
