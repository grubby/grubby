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

//line parser.y:800

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 25,
	43, 92,
	-2, 20,
	-1, 76,
	9, 64,
	10, 64,
	-2, 156,
	-1, 87,
	43, 92,
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
	-1, 102,
	43, 92,
	-2, 90,
	-1, 113,
	43, 92,
	-2, 20,
	-1, 197,
	9, 64,
	10, 64,
	-2, 156,
	-1, 234,
	43, 93,
	-2, 91,
}

const RubyNprod = 165
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1178

var RubyAct = []int{

	9, 78, 33, 165, 84, 163, 77, 62, 162, 2,
	3, 189, 153, 239, 187, 122, 154, 75, 64, 116,
	117, 279, 256, 233, 4, 83, 22, 194, 128, 189,
	169, 26, 255, 72, 60, 59, 185, 272, 101, 103,
	236, 215, 64, 98, 111, 97, 96, 62, 238, 61,
	275, 65, 66, 67, 188, 105, 88, 124, 63, 70,
	68, 69, 274, 189, 201, 131, 132, 118, 135, 136,
	137, 138, 189, 140, 254, 88, 186, 124, 189, 104,
	147, 148, 63, 125, 260, 187, 286, 83, 88, 94,
	189, 73, 145, 146, 74, 124, 88, 88, 231, 88,
	88, 88, 88, 125, 88, 155, 125, 213, 88, 167,
	157, 88, 88, 225, 72, 95, 180, 247, 88, 249,
	248, 125, 151, 284, 226, 94, 88, 62, 192, 243,
	213, 83, 195, 211, 187, 27, 278, 196, 23, 86,
	87, 76, 253, 81, 91, 234, 62, 199, 198, 187,
	102, 200, 202, 180, 183, 184, 204, 180, 277, 88,
	89, 262, 88, 54, 55, 166, 127, 80, 208, 180,
	214, 180, 218, 172, 180, 180, 119, 120, 88, 89,
	150, 144, 79, 88, 92, 130, 53, 52, 228, 166,
	227, 164, 89, 270, 58, 23, 193, 113, 124, 232,
	89, 89, 82, 89, 89, 89, 89, 180, 89, 100,
	180, 99, 89, 235, 115, 89, 89, 242, 191, 88,
	54, 55, 89, 121, 125, 207, 182, 190, 1, 88,
	89, 189, 28, 221, 259, 114, 51, 50, 49, 56,
	48, 57, 180, 53, 52, 47, 180, 180, 46, 18,
	17, 64, 180, 180, 16, 15, 37, 90, 10, 13,
	12, 11, 21, 89, 71, 88, 89, 14, 180, 180,
	180, 19, 31, 30, 32, 29, 90, 180, 0, 0,
	282, 180, 89, 85, 65, 66, 67, 89, 0, 90,
	0, 63, 70, 68, 69, 0, 0, 90, 90, 0,
	90, 90, 90, 90, 0, 90, 152, 156, 181, 90,
	5, 88, 90, 90, 0, 126, 168, 0, 170, 90,
	0, 0, 0, 89, 133, 173, 174, 90, 0, 212,
	0, 141, 0, 89, 216, 126, 0, 0, 126, 149,
	0, 0, 0, 0, 0, 85, 0, 106, 107, 108,
	109, 110, 0, 126, 0, 0, 0, 229, 230, 0,
	90, 0, 0, 90, 0, 0, 0, 0, 206, 89,
	209, 129, 0, 0, 0, 134, 0, 0, 0, 90,
	139, 0, 142, 143, 90, 0, 0, 0, 0, 85,
	0, 23, 86, 87, 44, 0, 0, 91, 257, 0,
	0, 0, 0, 20, 158, 159, 160, 161, 263, 0,
	203, 0, 264, 171, 0, 89, 54, 55, 0, 0,
	90, 0, 0, 175, 0, 273, 0, 241, 93, 0,
	90, 245, 0, 246, 0, 56, 0, 92, 251, 53,
	52, 252, 0, 0, 0, 0, 0, 0, 283, 285,
	0, 287, 0, 288, 0, 0, 126, 0, 0, 0,
	93, 23, 112, 113, 267, 268, 90, 91, 269, 93,
	0, 23, 86, 87, 197, 0, 93, 91, 0, 0,
	93, 276, 0, 93, 93, 0, 54, 55, 0, 0,
	93, 0, 280, 0, 0, 64, 54, 55, 93, 0,
	23, 86, 87, 76, 0, 56, 91, 92, 0, 53,
	52, 0, 90, 0, 0, 56, 0, 92, 0, 53,
	52, 0, 0, 0, 244, 54, 55, 0, 65, 66,
	67, 250, 0, 0, 93, 63, 70, 68, 69, 0,
	0, 258, 0, 0, 79, 261, 92, 0, 53, 52,
	0, 0, 0, 0, 0, 93, 0, 266, 23, 24,
	25, 44, 0, 0, 271, 34, 220, 42, 223, 222,
	43, 35, 36, 0, 0, 0, 45, 0, 0, 0,
	0, 0, 0, 54, 55, 0, 0, 281, 38, 39,
	40, 41, 0, 0, 178, 179, 23, 24, 25, 44,
	0, 93, 56, 34, 57, 42, 53, 52, 43, 35,
	36, 0, 0, 0, 45, 0, 0, 0, 0, 0,
	0, 54, 55, 0, 0, 0, 38, 39, 40, 41,
	0, 0, 6, 7, 0, 0, 0, 0, 0, 0,
	56, 0, 57, 0, 53, 52, 0, 8, 23, 24,
	25, 44, 0, 0, 0, 34, 265, 42, 0, 0,
	43, 35, 36, 0, 0, 0, 45, 0, 0, 0,
	0, 0, 0, 54, 55, 0, 0, 0, 38, 39,
	40, 41, 0, 0, 178, 179, 23, 24, 25, 44,
	0, 0, 56, 34, 57, 42, 53, 52, 43, 35,
	36, 0, 0, 0, 45, 0, 0, 0, 0, 0,
	0, 54, 55, 0, 0, 0, 38, 39, 40, 41,
	0, 0, 178, 179, 0, 0, 0, 0, 0, 0,
	56, 0, 57, 240, 53, 52, 23, 24, 25, 44,
	0, 0, 0, 34, 237, 42, 0, 0, 43, 35,
	36, 0, 0, 0, 45, 0, 0, 0, 0, 0,
	0, 54, 55, 0, 0, 0, 38, 39, 40, 41,
	0, 0, 178, 179, 0, 0, 0, 0, 0, 0,
	56, 0, 57, 0, 53, 52, 23, 24, 25, 44,
	0, 0, 0, 34, 224, 42, 0, 0, 43, 35,
	36, 0, 0, 0, 45, 0, 0, 0, 0, 0,
	0, 54, 55, 0, 0, 0, 38, 39, 40, 41,
	0, 0, 178, 179, 0, 0, 0, 0, 0, 0,
	56, 0, 57, 0, 53, 52, 23, 24, 25, 44,
	0, 0, 0, 34, 219, 42, 0, 0, 43, 35,
	36, 0, 0, 0, 45, 0, 0, 0, 0, 0,
	0, 54, 55, 0, 0, 0, 38, 39, 40, 41,
	0, 0, 178, 179, 0, 0, 0, 0, 0, 0,
	56, 0, 57, 0, 53, 52, 23, 24, 25, 44,
	0, 0, 0, 34, 217, 42, 0, 0, 43, 35,
	36, 0, 0, 0, 45, 0, 0, 0, 0, 0,
	0, 54, 55, 0, 0, 0, 38, 39, 40, 41,
	0, 0, 178, 179, 23, 24, 25, 44, 0, 0,
	56, 34, 57, 42, 53, 52, 43, 35, 36, 0,
	0, 0, 45, 0, 0, 0, 0, 0, 0, 54,
	55, 0, 0, 0, 38, 39, 40, 41, 0, 0,
	178, 179, 0, 0, 0, 0, 0, 0, 56, 0,
	57, 210, 53, 52, 23, 24, 25, 44, 0, 0,
	0, 34, 205, 42, 0, 0, 43, 35, 36, 0,
	0, 0, 45, 0, 0, 0, 0, 0, 0, 54,
	55, 0, 0, 0, 38, 39, 40, 41, 0, 0,
	178, 179, 23, 24, 25, 44, 0, 0, 56, 34,
	57, 42, 53, 52, 43, 35, 36, 0, 0, 0,
	45, 0, 0, 0, 0, 0, 0, 54, 55, 0,
	0, 0, 38, 39, 40, 41, 0, 0, 178, 179,
	23, 24, 25, 44, 177, 0, 56, 34, 57, 42,
	53, 52, 43, 35, 36, 0, 0, 0, 45, 0,
	0, 0, 0, 0, 0, 54, 55, 0, 0, 0,
	38, 39, 40, 41, 0, 0, 0, 176, 23, 24,
	25, 44, 0, 0, 56, 34, 57, 42, 53, 52,
	43, 35, 36, 0, 0, 0, 45, 0, 0, 0,
	0, 0, 0, 54, 55, 23, 112, 113, 38, 39,
	40, 41, 23, 86, 87, 44, 0, 0, 123, 86,
	87, 44, 56, 0, 57, 0, 53, 52, 0, 0,
	54, 55, 0, 0, 0, 0, 0, 54, 55, 0,
	0, 0, 0, 54, 55, 0, 0, 0, 0, 56,
	0, 57, 0, 53, 52, 0, 56, 0, 57, 0,
	53, 52, 56, 0, 57, 0, 53, 52,
}
var RubyPact = []int{

	-32, 591, -1000, -1000, -1000, -7, -1000, -1000, -1000, 247,
	77, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-27, -1000, -1000, -1000, 133, 81, 12, 11, 9, -1000,
	-1000, -1000, -1000, -1000, 205, 143, 143, 45, 1083, 1083,
	1083, 1083, 1083, 1110, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 13, 170, -1000, -1000, 1123, -1000, -15, -1000,
	-1000, -1000, 1083, 179, 1110, 1117, 1083, 1110, 1110, 1110,
	1110, 1083, 1117, 1083, 1083, 175, 1117, -1000, 83, 1123,
	1117, 174, 112, 38, -1000, -1000, 495, 117, -1000, -1000,
	-1000, -29, -29, -27, 1117, 1083, 1083, 1083, 1083, 183,
	183, -2, -1000, -1000, 1083, 167, 33, 33, 33, 33,
	33, -1000, -1000, -1000, 1045, 1007, -1000, -1000, 148, -1000,
	-1000, 26, 4, -1000, 491, -1000, -4, 190, -16, 33,
	466, -1000, 38, -1000, 33, -1000, -1000, -1000, -1000, 33,
	38, -1000, 33, 33, -1000, 139, 456, 14, 38, -1000,
	-1000, 386, 969, -1000, 162, -1000, 919, 124, 33, 33,
	33, 33, -1000, 97, 159, -1000, 7, -1000, 881, 143,
	831, 33, -1000, 553, 781, 33, -1000, -1000, -1000, -1000,
	247, 33, 100, -1000, -1000, 185, -1000, 1110, -1000, -1000,
	-1000, 88, 195, -20, 138, 83, -1000, 1117, -1000, -1000,
	-1000, 6, 38, -1000, -1000, -1000, 731, 3, -1000, 681,
	-1000, -1000, -12, 159, 120, 1083, -12, -1000, -1000, -1000,
	-1000, 104, 1083, -1000, -1000, -1000, 135, -1000, -1000, 22,
	-30, -1000, 1083, 1110, -1000, 75, 1083, -1000, -1000, 155,
	-1000, 1007, -1000, -1000, 33, 1007, 643, -1000, 1083, -1000,
	33, 1007, 1007, 189, -1000, 1083, -1000, 31, 33, -1000,
	-1000, 33, -1000, 49, 37, -1000, 33, 1007, 1007, 1007,
	152, 132, -22, -12, -1000, -1000, 1007, -1000, 1083, 1110,
	1007, 113, 76, -12, -1000, -12, -1000, -12, -12,
}
var RubyPgo = []int{

	0, 308, 275, 274, 4, 273, 272, 403, 232, 271,
	267, 262, 0, 135, 258, 261, 26, 260, 31, 259,
	2, 3, 256, 255, 254, 250, 249, 248, 245, 240,
	238, 237, 236, 214, 235, 6, 12, 233, 228, 8,
	227, 226, 15, 225, 223, 218, 202, 1, 5, 194,
	166,
}
var RubyR1 = []int{

	0, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 50, 50, 33, 33, 33, 33, 33, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 16,
	16, 16, 16, 16, 16, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	35, 35, 44, 44, 42, 42, 42, 42, 42, 47,
	47, 47, 47, 46, 46, 46, 46, 46, 15, 15,
	39, 39, 21, 21, 48, 48, 48, 17, 17, 19,
	20, 20, 49, 49, 10, 10, 10, 10, 10, 10,
	10, 8, 8, 18, 18, 13, 13, 22, 22, 23,
	24, 25, 26, 27, 27, 27, 27, 28, 29, 30,
	31, 32, 2, 5, 6, 6, 3, 3, 40, 40,
	40, 40, 45, 45, 45, 45, 4, 4, 4, 4,
	36, 43, 43, 43, 9, 9, 9, 9, 9, 9,
	9, 37, 37, 37, 37, 37, 34, 34, 34, 7,
	11, 41, 41, 41, 41,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 4,
	4, 2, 3, 4, 4, 3, 2, 3, 4, 6,
	3, 1, 1, 3, 0, 1, 1, 1, 3, 1,
	1, 3, 3, 1, 1, 3, 3, 3, 7, 7,
	1, 3, 1, 3, 0, 1, 3, 4, 6, 4,
	1, 4, 1, 4, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 2, 2, 3, 3, 3, 3, 2,
	2, 2, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 1, 1, 3, 3, 5, 5, 0, 4,
	7, 8, 3, 3, 7, 8, 3, 4, 4, 3,
	3, 0, 1, 3, 4, 5, 3, 3, 3, 3,
	4, 0, 4, 3, 3, 2, 0, 2, 2, 3,
	4, 0, 3, 4, 6,
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
	-8, 11, 51, -7, 8, 34, 34, 34, 34, 6,
	4, -20, 7, -20, 34, 10, -1, -1, -1, -1,
	-1, -12, 6, 7, -34, -33, 6, 7, 54, 6,
	7, -44, -42, 5, -12, -16, -14, -50, 43, -1,
	6, -12, -12, -14, -1, -12, -12, -12, -12, -1,
	-12, -14, -1, -1, 6, -42, 10, -12, -12, -14,
	6, 10, -33, -36, 45, -36, -33, -42, -1, -1,
	-1, -1, -39, -48, 8, -21, 6, -39, -33, 32,
	-33, -1, 6, -33, -33, -1, 42, 9, 41, 42,
	-12, -1, -41, 6, 7, 10, 50, 10, 50, 41,
	-40, -45, -12, 6, 43, -47, -35, 8, 9, -12,
	-4, 50, -12, -14, -4, 13, -33, -43, 6, -33,
	52, 9, -50, 10, -48, 34, -50, 13, -20, 13,
	13, -37, 16, 15, 13, 13, 24, 5, -12, -50,
	-50, 10, 4, 43, 7, -42, 34, 13, 45, 10,
	52, -33, -21, 9, -1, -33, -33, 13, 16, 15,
	-1, -33, -33, 7, 52, 10, 52, -50, -1, -12,
	9, -1, 6, -50, -50, 13, -1, -33, -33, -33,
	4, -1, 6, -50, 13, 13, -33, 6, 4, 43,
	-33, -1, -12, -50, 10, -50, 10, -50, -50,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
	46, 47, 48, 18, 19, -2, 21, 22, 23, 24,
	25, 26, 27, 28, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 156, 13, 29, 30, 31, 32,
	33, 34, 0, 0, 122, 123, 64, 11, 0, 5,
	6, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, -2, 51, 56, 64,
	0, 0, 61, 69, 70, 74, 19, -2, 21, 22,
	23, 13, -2, 0, 64, 0, 0, 0, 0, 84,
	84, 13, -2, 13, 0, 0, 109, 110, 111, 112,
	13, 13, 19, -2, 0, 161, 103, 104, 0, 101,
	102, 0, 0, 18, 65, 66, 67, 128, 0, 146,
	52, 57, 113, 115, 117, 118, 119, 120, 121, 148,
	114, 116, 147, 149, 55, 0, 0, 65, 94, 95,
	107, 0, 0, 13, 141, 13, 0, 0, 96, 97,
	98, 99, 11, 80, 84, 85, 82, 11, 0, 0,
	0, 100, 108, 0, 0, 157, 158, 159, 14, 15,
	16, 17, 0, 105, 106, 0, 124, 0, 125, 12,
	11, 11, 0, 19, 0, 53, 54, -2, 49, 71,
	72, 58, 75, 76, 77, 136, 0, 0, 142, 0,
	139, 50, 13, 0, 0, 0, 13, 87, 13, 89,
	144, 0, 0, 13, 150, 160, 13, 63, 68, 0,
	0, 11, 0, 0, -2, 0, 0, 137, 140, 0,
	138, 11, 86, 81, 83, 11, 0, 145, 0, 13,
	13, 155, 162, 13, 126, 0, 127, 0, 11, 132,
	60, 59, 143, 0, 0, 88, 13, 153, 154, 163,
	0, 0, 0, 129, 78, 79, 152, 13, 0, 0,
	164, 11, 11, 130, 11, 134, 11, 131, 135,
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
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 50:
		//line parser.y:207
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 51:
		//line parser.y:214
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 52:
		//line parser.y:221
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 53:
		//line parser.y:228
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 54:
		//line parser.y:236
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 55:
		//line parser.y:244
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 56:
		//line parser.y:254
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 57:
		//line parser.y:261
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 58:
		//line parser.y:271
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 59:
		//line parser.y:281
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 60:
		//line parser.y:290
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 61:
		//line parser.y:292
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 62:
		//line parser.y:295
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 63:
		//line parser.y:297
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 64:
		//line parser.y:299
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 65:
		//line parser.y:301
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 66:
		//line parser.y:303
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:305
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:307
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 69:
		//line parser.y:311
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 70:
		//line parser.y:313
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 71:
		//line parser.y:315
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 72:
		//line parser.y:317
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:320
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:322
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 75:
		//line parser.y:324
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:326
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:328
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 78:
		//line parser.y:334
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 79:
		//line parser.y:342
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 80:
		//line parser.y:351
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 81:
		//line parser.y:353
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 82:
		//line parser.y:356
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 83:
		//line parser.y:358
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 84:
		//line parser.y:360
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 85:
		//line parser.y:362
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 86:
		//line parser.y:366
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 87:
		//line parser.y:371
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 88:
		//line parser.y:378
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 89:
		//line parser.y:388
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 90:
		//line parser.y:397
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 91:
		//line parser.y:403
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 92:
		//line parser.y:411
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 93:
		//line parser.y:415
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 94:
		//line parser.y:420
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
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
		//line parser.y:470
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 102:
		//line parser.y:472
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 103:
		//line parser.y:475
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 104:
		//line parser.y:477
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 105:
		//line parser.y:480
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 106:
		//line parser.y:482
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 107:
		//line parser.y:485
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 108:
		//line parser.y:487
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 109:
		//line parser.y:489
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 110:
		//line parser.y:490
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 111:
		//line parser.y:491
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 112:
		//line parser.y:492
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 113:
		//line parser.y:495
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 114:
		//line parser.y:503
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 115:
		//line parser.y:511
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 116:
		//line parser.y:519
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 117:
		//line parser.y:528
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 118:
		//line parser.y:537
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 119:
		//line parser.y:546
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 120:
		//line parser.y:555
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 121:
		//line parser.y:564
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 122:
		//line parser.y:572
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 123:
		//line parser.y:573
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 124:
		//line parser.y:575
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 125:
		//line parser.y:577
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 126:
		//line parser.y:580
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 127:
		//line parser.y:588
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 128:
		//line parser.y:596
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 129:
		//line parser.y:598
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
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
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 132:
		//line parser.y:620
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
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
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 135:
		//line parser.y:641
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 136:
		//line parser.y:649
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 137:
		//line parser.y:651
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
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
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 140:
		//line parser.y:668
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 141:
		//line parser.y:670
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 142:
		//line parser.y:672
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 143:
		//line parser.y:674
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 144:
		//line parser.y:677
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 145:
		//line parser.y:684
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 146:
		//line parser.y:692
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
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
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 149:
		//line parser.y:713
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 150:
		//line parser.y:720
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-2].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 151:
		//line parser.y:727
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 152:
		//line parser.y:729
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 153:
		//line parser.y:736
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 154:
		//line parser.y:743
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 155:
		//line parser.y:750
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 156:
		//line parser.y:757
		{
		}
	case 157:
		//line parser.y:758
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 158:
		//line parser.y:759
		{
		}
	case 159:
		//line parser.y:762
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 160:
		//line parser.y:765
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 161:
		//line parser.y:773
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 162:
		//line parser.y:775
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 163:
		//line parser.y:777
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 164:
		//line parser.y:786
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
