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
const FOR = 57361
const WHILE = 57362
const UNTIL = 57363
const BEGIN = 57364
const RESCUE = 57365
const ENSURE = 57366
const BREAK = 57367
const REDO = 57368
const RETRY = 57369
const RETURN = 57370
const TRUE = 57371
const FALSE = 57372
const LESSTHAN = 57373
const GREATERTHAN = 57374
const EQUALTO = 57375
const BANG = 57376
const COMPLEMENT = 57377
const POSITIVE = 57378
const NEGATIVE = 57379
const STAR = 57380
const WHITESPACE = 57381
const NEWLINE = 57382
const SEMICOLON = 57383
const COLON = 57384
const DOT = 57385
const PIPE = 57386
const SLASH = 57387
const AMPERSAND = 57388
const MODULO = 57389
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

//line parser.y:564

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
	38, 17,
	40, 17,
	41, 17,
	43, 17,
	44, 17,
	45, 17,
	46, 17,
	52, 17,
	56, 17,
	-2, 12,
	-1, 12,
	33, 12,
	-2, 18,
	-1, 13,
	33, 12,
	-2, 19,
	-1, 14,
	33, 12,
	-2, 20,
	-1, 15,
	33, 12,
	-2, 21,
	-1, 82,
	5, 12,
	6, 12,
	8, 12,
	-2, 44,
	-1, 84,
	10, 12,
	-2, 46,
	-1, 87,
	10, 12,
	-2, 49,
	-1, 88,
	4, 12,
	36, 12,
	37, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 50,
	-1, 90,
	1, 17,
	4, 17,
	10, 17,
	13, 17,
	15, 17,
	16, 17,
	38, 17,
	40, 17,
	41, 17,
	43, 17,
	44, 17,
	45, 17,
	46, 17,
	52, 17,
	56, 17,
	-2, 12,
	-1, 92,
	44, 12,
	-2, 8,
	-1, 103,
	4, 12,
	36, 12,
	37, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 82,
	-1, 104,
	4, 12,
	36, 12,
	37, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 83,
	-1, 105,
	4, 12,
	36, 12,
	37, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 84,
	-1, 106,
	4, 12,
	36, 12,
	37, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 85,
	-1, 111,
	4, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 8,
	-1, 134,
	42, 75,
	-2, 12,
	-1, 141,
	4, 12,
	36, 12,
	37, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 47,
	-1, 142,
	4, 12,
	36, 12,
	37, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 86,
	-1, 143,
	4, 12,
	36, 12,
	37, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 87,
	-1, 144,
	4, 12,
	36, 12,
	37, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 88,
	-1, 145,
	4, 12,
	36, 12,
	37, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 89,
	-1, 146,
	4, 12,
	36, 12,
	37, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 90,
	-1, 147,
	4, 12,
	36, 12,
	37, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 91,
	-1, 152,
	4, 12,
	36, 12,
	37, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 77,
	-1, 157,
	4, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 11,
	-1, 160,
	44, 116,
	-2, 59,
	-1, 161,
	4, 12,
	36, 12,
	37, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 78,
	-1, 162,
	4, 12,
	36, 12,
	37, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 79,
	-1, 163,
	4, 12,
	36, 12,
	37, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 80,
	-1, 164,
	4, 12,
	36, 12,
	37, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 81,
	-1, 177,
	4, 17,
	38, 17,
	43, 17,
	44, 17,
	45, 17,
	46, 17,
	-2, 12,
	-1, 187,
	44, 117,
	-2, 60,
	-1, 204,
	4, 12,
	36, 12,
	37, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 52,
	-1, 218,
	42, 76,
	-2, 12,
	-1, 229,
	4, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 8,
	-1, 241,
	10, 101,
	40, 101,
	52, 101,
	-2, 12,
	-1, 242,
	4, 12,
	38, 12,
	44, 12,
	45, 12,
	46, 12,
	-2, 8,
	-1, 253,
	44, 118,
	-2, 63,
	-1, 256,
	10, 98,
	40, 98,
	52, 98,
	-2, 12,
	-1, 262,
	40, 102,
	52, 102,
	-2, 12,
	-1, 265,
	40, 99,
	52, 99,
	-2, 12,
}

const RubyNprod = 128
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 853

var RubyAct = []int{

	124, 189, 6, 107, 125, 89, 100, 110, 52, 8,
	53, 83, 55, 56, 57, 58, 59, 60, 61, 2,
	249, 247, 75, 7, 266, 212, 172, 54, 7, 160,
	258, 198, 7, 4, 5, 7, 192, 170, 7, 138,
	65, 66, 67, 231, 246, 68, 69, 70, 71, 72,
	73, 259, 74, 7, 76, 77, 78, 54, 214, 215,
	211, 191, 81, 79, 80, 62, 173, 101, 101, 264,
	214, 215, 75, 136, 102, 133, 112, 113, 114, 115,
	116, 117, 118, 119, 132, 120, 121, 122, 123, 96,
	95, 55, 261, 126, 127, 128, 129, 130, 7, 94,
	131, 225, 54, 93, 76, 77, 78, 240, 137, 207,
	208, 139, 81, 79, 80, 92, 140, 75, 199, 232,
	201, 200, 224, 206, 172, 222, 153, 150, 149, 154,
	86, 148, 218, 151, 101, 169, 134, 167, 97, 98,
	168, 171, 255, 166, 63, 64, 219, 253, 99, 76,
	77, 78, 182, 82, 183, 184, 157, 81, 79, 80,
	185, 219, 220, 9, 188, 109, 187, 109, 108, 210,
	135, 84, 87, 194, 175, 195, 196, 197, 55, 186,
	174, 202, 1, 165, 179, 159, 203, 32, 205, 213,
	209, 31, 30, 29, 216, 28, 217, 27, 26, 226,
	25, 227, 24, 221, 223, 23, 228, 35, 19, 13,
	18, 17, 88, 14, 20, 230, 36, 16, 233, 235,
	15, 33, 236, 22, 238, 103, 104, 105, 106, 34,
	237, 111, 239, 245, 243, 21, 3, 0, 248, 0,
	250, 251, 0, 0, 0, 0, 0, 252, 0, 0,
	0, 0, 0, 0, 0, 257, 0, 0, 0, 260,
	0, 0, 263, 0, 0, 0, 0, 0, 0, 141,
	142, 143, 144, 145, 146, 147, 0, 0, 152, 0,
	0, 10, 11, 12, 161, 162, 163, 164, 39, 178,
	51, 181, 180, 40, 41, 0, 176, 0, 0, 0,
	0, 0, 0, 0, 0, 42, 43, 0, 0, 0,
	44, 45, 46, 47, 0, 156, 155, 0, 0, 0,
	0, 0, 0, 0, 0, 48, 0, 49, 0, 38,
	37, 50, 0, 0, 0, 0, 0, 0, 0, 204,
	0, 10, 11, 12, 0, 0, 0, 0, 39, 244,
	51, 0, 0, 40, 41, 0, 0, 0, 0, 229,
	0, 0, 0, 0, 0, 42, 43, 0, 0, 0,
	44, 45, 46, 47, 0, 156, 155, 0, 0, 0,
	0, 0, 0, 241, 242, 48, 0, 49, 0, 38,
	37, 50, 0, 0, 10, 11, 12, 0, 0, 0,
	0, 39, 234, 51, 0, 254, 40, 41, 256, 0,
	0, 0, 0, 0, 0, 0, 0, 262, 42, 43,
	265, 0, 0, 44, 45, 46, 47, 0, 156, 155,
	0, 0, 0, 0, 0, 0, 0, 0, 48, 0,
	49, 0, 38, 37, 50, 10, 11, 12, 0, 0,
	0, 0, 39, 193, 51, 0, 0, 40, 41, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 42,
	43, 0, 0, 0, 44, 45, 46, 47, 0, 156,
	155, 0, 0, 0, 0, 0, 0, 0, 0, 48,
	0, 49, 0, 38, 37, 50, 10, 11, 12, 0,
	0, 0, 0, 39, 190, 51, 0, 0, 40, 41,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	42, 43, 0, 0, 0, 44, 45, 46, 47, 0,
	156, 155, 0, 0, 0, 0, 0, 0, 0, 0,
	48, 0, 49, 0, 38, 37, 50, 10, 11, 12,
	0, 0, 0, 0, 39, 158, 51, 0, 0, 40,
	41, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 42, 43, 0, 0, 0, 44, 45, 46, 47,
	0, 156, 155, 0, 0, 0, 0, 0, 0, 0,
	0, 48, 0, 49, 0, 38, 37, 50, 91, 90,
	12, 86, 0, 0, 92, 39, 0, 51, 0, 0,
	40, 41, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 42, 43, 0, 0, 85, 44, 45, 46,
	47, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 48, 0, 49, 0, 38, 37, 50, 10,
	11, 12, 0, 0, 0, 0, 39, 0, 51, 0,
	0, 40, 41, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 42, 43, 0, 0, 0, 44, 45,
	46, 47, 0, 156, 155, 0, 0, 0, 0, 0,
	0, 0, 0, 48, 0, 49, 0, 38, 37, 50,
	10, 11, 12, 0, 0, 0, 92, 39, 0, 51,
	0, 0, 40, 41, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 42, 43, 0, 0, 0, 44,
	45, 46, 47, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 48, 0, 49, 0, 38, 37,
	50, 10, 11, 12, 0, 0, 0, 0, 39, 0,
	51, 0, 0, 40, 41, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 42, 43, 0, 0, 0,
	44, 45, 46, 47, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 48, 0, 49, 0, 38,
	37, 50, 10, 177, 12, 0, 0, 0, 0, 39,
	0, 51, 0, 0, 40, 41, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 42, 43, 0, 0,
	0, 44, 45, 46, 47, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 48, 0, 49, 0,
	38, 37, 50,
}
var RubyPact = []int{

	-37, -7, -1000, -47, -1000, -1000, 746, -4, -1000, -16,
	-1000, -4, -4, -4, -4, -4, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 11, 138, -4,
	-4, -4, -1000, -1000, -4, -4, -4, -4, -4, -4,
	-1000, -4, -1000, 113, 147, 593, 70, 66, 57, 56,
	-1000, -1000, 132, -1000, -1000, 142, -4, -4, 746, 746,
	746, 746, 162, -1000, 746, -4, -4, -4, -4, -4,
	-4, -4, -4, -1000, -4, -4, -4, -4, -16, -1000,
	-4, -1000, -4, -4, -4, -4, -4, -1000, -1000, -4,
	44, 129, 33, -16, -16, -16, -16, -4, -1000, -1000,
	-1, -16, 746, 746, 746, 746, 746, 746, 746, 122,
	123, 746, 162, 119, 113, 542, -15, 746, 746, 746,
	746, 122, -1000, -4, -4, -5, -1000, 16, -1000, 797,
	276, -16, -16, -16, -16, -16, -16, -16, -1000, -1000,
	-1000, -4, -16, -4, -4, -1000, -1000, -16, -1000, -4,
	160, -16, -16, -16, -16, -1000, -1000, 491, 21, -1000,
	-6, 440, -4, -1000, -4, -4, -16, -11, -1000, 105,
	-4, -1000, 695, 114, 104, 20, -19, -1000, -4, 19,
	-1000, -1000, 125, -1000, 156, 115, 112, 68, -4, -1000,
	-4, -1000, 746, 644, -16, -1000, -1000, -1000, -1000, -1000,
	3, -1000, -1000, 109, -1000, -1000, 644, 389, -4, -1000,
	-1000, -1, -1000, -1, -1000, 75, 746, 746, 644, -16,
	336, -1000, -4, 31, -1000, -1000, -31, -1, -32, -1,
	-4, -16, -16, 644, -1000, 141, -1000, -1000, 746, -1000,
	136, 746, 644, -1000, -16, -12, -16, 18, -4, 60,
	746, -4, 59, 746, -1000, 14, -1000,
}
var RubyPgo = []int{

	0, 237, 156, 236, 235, 229, 5, 223, 221, 220,
	217, 216, 214, 213, 211, 210, 209, 208, 6, 207,
	205, 202, 200, 198, 197, 195, 193, 192, 191, 187,
	4, 11, 185, 184, 183, 182, 180, 3, 179, 174,
	172, 171, 170, 0, 7, 1, 169,
}
var RubyR1 = []int{

	0, 35, 35, 35, 35, 3, 3, 3, 30, 30,
	30, 30, 43, 43, 44, 44, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 10, 10, 10, 10, 10, 31, 31,
	41, 41, 41, 41, 40, 40, 40, 40, 40, 37,
	37, 37, 37, 37, 45, 45, 45, 14, 34, 34,
	15, 15, 17, 18, 18, 42, 42, 12, 12, 12,
	12, 12, 20, 21, 22, 23, 24, 25, 26, 27,
	28, 29, 4, 7, 8, 5, 5, 36, 36, 36,
	36, 39, 39, 39, 1, 1, 9, 9, 16, 16,
	13, 13, 19, 6, 6, 32, 38, 38, 38, 46,
	46, 11, 11, 33, 33, 33, 33, 33,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 1, 3, 0, 2,
	2, 2, 0, 2, 0, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 3, 5, 3, 5, 5, 1,
	1, 1, 5, 5, 1, 1, 5, 5, 5, 0,
	1, 1, 5, 5, 0, 2, 2, 9, 0, 1,
	6, 8, 6, 3, 6, 1, 4, 5, 5, 5,
	5, 5, 3, 3, 3, 3, 5, 5, 5, 5,
	5, 5, 1, 1, 5, 9, 9, 0, 6, 11,
	12, 4, 9, 10, 0, 1, 2, 2, 2, 2,
	3, 3, 1, 3, 7, 3, 0, 1, 5, 1,
	2, 5, 6, 0, 5, 3, 4, 2,
}
var RubyChk = []int{

	-1000, -35, 56, -3, 40, 41, -43, 39, 56, -2,
	5, 6, 7, -16, -13, -9, -10, -14, -15, -17,
	-12, -4, -7, -20, -21, -22, -23, -24, -25, -26,
	-27, -28, -29, -8, -5, -19, -11, 54, 53, 12,
	17, 18, 29, 30, 34, 35, 36, 37, 49, 51,
	55, 14, -43, -43, 43, -43, -43, -43, -43, -43,
	6, 7, 54, 6, 7, -43, -43, -43, -43, -43,
	-43, -43, -43, -43, -43, 4, 36, 37, 38, 45,
	46, 44, 6, -31, -41, 33, 8, -40, -2, -6,
	6, 5, 11, 33, 33, 33, 33, 6, 7, 6,
	-18, -43, -18, -2, -2, -2, -2, -37, 6, 5,
	-44, -2, -43, -43, -43, -43, -43, -43, -43, -43,
	-43, -43, -43, -43, -43, -30, -43, -43, -43, -43,
	-43, -43, 40, 31, 7, -42, 40, -43, 40, -43,
	-30, -2, -2, -2, -2, -2, -2, -2, -31, 6,
	5, 10, -2, -37, 10, 40, 39, -2, 13, -32,
	44, -2, -2, -2, -2, -34, -31, -30, -18, -43,
	42, -30, 10, 50, -36, -39, -2, 6, 13, -33,
	16, 15, -43, -43, -43, -43, -38, 6, -37, -45,
	13, 40, 42, 13, -43, -43, -43, -43, 42, 13,
	16, 15, -43, -30, -2, -6, 9, 5, 6, -6,
	-46, 40, 44, -43, 39, 40, -30, -30, 7, 5,
	6, -44, 10, -44, 10, 33, -43, -43, -30, -2,
	-30, 40, 10, -45, 13, -43, -43, -44, -43, -44,
	32, -2, -2, -30, 13, -43, 13, 52, -43, 52,
	-43, -43, -30, 6, -2, 6, -2, -43, 42, 33,
	-43, 32, -2, -43, 10, -2, 10,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 6, 0, 12, 4, 12,
	16, -2, -2, -2, -2, -2, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, 42, 0, 0, 12,
	12, 12, 92, 93, 12, 12, 12, 12, 12, 12,
	112, 12, 13, 7, 0, 0, 0, 0, 0, 0,
	108, 109, 0, 106, 107, 0, 12, 12, 0, 0,
	0, 0, 59, 14, 0, 12, 12, 12, 12, 12,
	12, 12, -2, 43, -2, 12, 12, -2, -2, 51,
	-2, 16, -2, 12, 12, 12, 12, 110, 111, 12,
	0, 0, 0, -2, -2, -2, -2, 12, 60, 61,
	12, -2, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 59, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 8, 12, -2, 0, 8, 0, 15, 97,
	0, -2, -2, -2, -2, -2, -2, -2, 45, 54,
	55, 12, -2, 12, 12, 9, 10, -2, 113, 12,
	-2, -2, -2, -2, -2, 64, 69, 0, 0, 73,
	0, 0, 12, 94, 12, 12, 12, -2, 121, 0,
	12, 8, 0, 0, 0, 0, 0, -2, 12, 8,
	70, 8, 0, 72, 0, 14, 14, 0, 12, 122,
	12, 8, 0, 127, -2, 53, 48, 56, 57, 58,
	8, 119, 115, 0, 65, 66, 64, 0, -2, 62,
	63, 12, 14, 12, 14, 0, 0, 0, 125, -2,
	0, 120, 12, 0, 71, 74, 0, 12, 0, 12,
	12, -2, -2, 126, 114, 0, 67, 95, 0, 96,
	0, 0, 124, -2, 12, 0, -2, 0, 12, 0,
	0, 12, -2, 0, 103, -2, 100,
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
		//line parser.y:148
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:150
		{
			Statements = []ast.Node{}
		}
	case 3:
		//line parser.y:152
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 4:
		//line parser.y:158
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 5:
		//line parser.y:164
		{
			RubyVAL.genericValue = nil
		}
	case 6:
		//line parser.y:165
		{
			RubyVAL.genericValue = nil
		}
	case 7:
		//line parser.y:166
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 8:
		//line parser.y:169
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 9:
		//line parser.y:171
		{ /* do nothing */
		}
	case 10:
		//line parser.y:173
		{ /* do nothing */
		}
	case 11:
		//line parser.y:175
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
		//line parser.y:187
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 44:
		//line parser.y:194
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 45:
		//line parser.y:201
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 46:
		//line parser.y:209
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 47:
		//line parser.y:216
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 48:
		//line parser.y:225
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 49:
		//line parser.y:227
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 50:
		//line parser.y:230
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 51:
		//line parser.y:232
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 52:
		//line parser.y:234
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 53:
		//line parser.y:236
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 54:
		//line parser.y:239
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 55:
		//line parser.y:241
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 56:
		//line parser.y:243
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 57:
		//line parser.y:245
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 58:
		//line parser.y:247
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 59:
		//line parser.y:249
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 60:
		//line parser.y:251
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 61:
		//line parser.y:253
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 62:
		//line parser.y:255
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 63:
		//line parser.y:257
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:267
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 68:
		//line parser.y:275
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 69:
		//line parser.y:276
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 70:
		//line parser.y:279
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 71:
		//line parser.y:286
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 72:
		//line parser.y:296
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 73:
		//line parser.y:305
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 74:
		//line parser.y:311
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 75:
		//line parser.y:319
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 76:
		//line parser.y:323
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 77:
		//line parser.y:328
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 78:
		//line parser.y:335
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 79:
		//line parser.y:342
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 80:
		//line parser.y:349
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 81:
		//line parser.y:356
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 82:
		//line parser.y:363
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 83:
		//line parser.y:364
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 84:
		//line parser.y:365
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 85:
		//line parser.y:366
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 86:
		//line parser.y:369
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 87:
		//line parser.y:378
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 88:
		//line parser.y:387
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 89:
		//line parser.y:396
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 90:
		//line parser.y:405
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 91:
		//line parser.y:414
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 92:
		//line parser.y:422
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 93:
		//line parser.y:423
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 94:
		//line parser.y:425
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 95:
		//line parser.y:429
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 96:
		//line parser.y:437
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 97:
		//line parser.y:445
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 98:
		//line parser.y:447
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 99:
		//line parser.y:449
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 100:
		//line parser.y:451
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 101:
		//line parser.y:454
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 102:
		//line parser.y:461
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 103:
		//line parser.y:468
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 104:
		//line parser.y:476
		{
			RubyVAL.genericValue = nil
		}
	case 105:
		//line parser.y:477
		{
			RubyVAL.genericValue = nil
		}
	case 106:
		//line parser.y:480
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 107:
		//line parser.y:482
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 108:
		//line parser.y:485
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 109:
		//line parser.y:487
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 110:
		//line parser.y:490
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 111:
		//line parser.y:492
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 112:
		//line parser.y:495
		{
			RubyVAL.genericValue = ast.FileNameConstReference{}
		}
	case 113:
		//line parser.y:498
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 114:
		//line parser.y:500
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 115:
		//line parser.y:508
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 116:
		//line parser.y:510
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 117:
		//line parser.y:512
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 118:
		//line parser.y:514
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 121:
		//line parser.y:519
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 122:
		//line parser.y:526
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 123:
		//line parser.y:534
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 124:
		//line parser.y:536
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 125:
		//line parser.y:543
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 126:
		//line parser.y:550
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 127:
		//line parser.y:557
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	}
	goto Rubystack /* stack new state and value */
}
