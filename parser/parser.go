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

//line parser.y:579

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
	17, 17,
	39, 17,
	41, 17,
	42, 17,
	44, 17,
	45, 17,
	46, 17,
	47, 17,
	53, 17,
	57, 17,
	-2, 12,
	-1, 12,
	34, 12,
	-2, 18,
	-1, 13,
	34, 12,
	-2, 19,
	-1, 14,
	34, 12,
	-2, 20,
	-1, 15,
	34, 12,
	-2, 21,
	-1, 84,
	5, 12,
	6, 12,
	8, 12,
	-2, 44,
	-1, 86,
	10, 12,
	-2, 46,
	-1, 89,
	10, 12,
	-2, 49,
	-1, 90,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 50,
	-1, 92,
	1, 17,
	4, 17,
	10, 17,
	13, 17,
	15, 17,
	16, 17,
	17, 17,
	39, 17,
	41, 17,
	42, 17,
	44, 17,
	45, 17,
	46, 17,
	47, 17,
	53, 17,
	57, 17,
	-2, 12,
	-1, 94,
	45, 12,
	-2, 8,
	-1, 105,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 82,
	-1, 106,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 83,
	-1, 107,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 84,
	-1, 108,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 85,
	-1, 113,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 8,
	-1, 138,
	43, 75,
	-2, 12,
	-1, 145,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 47,
	-1, 146,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 86,
	-1, 147,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 87,
	-1, 148,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 88,
	-1, 149,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 89,
	-1, 150,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 90,
	-1, 151,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 91,
	-1, 152,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 123,
	-1, 153,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 124,
	-1, 158,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 77,
	-1, 163,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 11,
	-1, 166,
	45, 116,
	-2, 59,
	-1, 167,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 78,
	-1, 168,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 79,
	-1, 169,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 80,
	-1, 170,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 81,
	-1, 183,
	4, 17,
	17, 17,
	39, 17,
	44, 17,
	45, 17,
	46, 17,
	47, 17,
	-2, 12,
	-1, 193,
	45, 117,
	-2, 60,
	-1, 210,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 52,
	-1, 224,
	43, 76,
	-2, 12,
	-1, 235,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 8,
	-1, 247,
	10, 101,
	41, 101,
	53, 101,
	-2, 12,
	-1, 248,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 8,
	-1, 259,
	45, 118,
	-2, 63,
	-1, 262,
	10, 98,
	41, 98,
	53, 98,
	-2, 12,
	-1, 268,
	41, 102,
	53, 102,
	-2, 12,
	-1, 271,
	41, 99,
	53, 99,
	-2, 12,
}

const RubyNprod = 130
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 809

var RubyAct = []int{

	128, 109, 6, 195, 102, 85, 8, 2, 52, 112,
	53, 91, 55, 56, 57, 58, 59, 60, 61, 255,
	253, 7, 178, 272, 270, 54, 7, 218, 166, 204,
	7, 4, 5, 252, 264, 198, 176, 7, 142, 237,
	65, 66, 67, 220, 221, 68, 69, 70, 71, 72,
	73, 217, 74, 7, 7, 129, 137, 54, 54, 197,
	220, 221, 140, 179, 98, 136, 62, 103, 103, 7,
	97, 96, 104, 95, 267, 246, 114, 115, 116, 117,
	118, 119, 120, 121, 122, 123, 238, 124, 125, 126,
	127, 212, 178, 55, 230, 130, 131, 132, 133, 134,
	224, 163, 135, 205, 228, 207, 206, 160, 9, 157,
	141, 213, 214, 143, 138, 156, 155, 94, 88, 99,
	100, 75, 63, 64, 225, 259, 225, 226, 159, 154,
	261, 82, 111, 193, 83, 111, 110, 101, 103, 175,
	84, 172, 174, 216, 139, 86, 89, 181, 192, 180,
	1, 265, 171, 185, 76, 77, 78, 90, 188, 165,
	189, 190, 81, 79, 80, 32, 191, 31, 194, 144,
	105, 106, 107, 108, 30, 29, 113, 28, 27, 200,
	26, 201, 202, 203, 55, 25, 24, 208, 23, 35,
	19, 13, 173, 18, 17, 219, 177, 14, 20, 36,
	211, 16, 215, 15, 33, 232, 22, 233, 34, 21,
	3, 227, 229, 0, 0, 0, 145, 146, 147, 148,
	149, 150, 151, 152, 153, 241, 239, 158, 242, 0,
	244, 0, 0, 167, 168, 169, 170, 0, 243, 251,
	245, 0, 0, 209, 254, 182, 256, 257, 0, 0,
	0, 222, 0, 223, 75, 0, 0, 0, 0, 0,
	0, 263, 0, 234, 82, 266, 0, 83, 269, 0,
	0, 0, 236, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 231, 0, 0, 76, 77, 78,
	210, 249, 10, 11, 12, 81, 79, 80, 0, 39,
	184, 51, 187, 186, 258, 40, 41, 0, 0, 0,
	235, 0, 0, 0, 0, 0, 0, 42, 43, 0,
	0, 0, 44, 45, 46, 47, 0, 162, 161, 0,
	0, 0, 0, 0, 247, 248, 0, 48, 0, 49,
	0, 38, 37, 50, 10, 11, 12, 0, 0, 0,
	0, 39, 250, 51, 0, 0, 260, 40, 41, 262,
	0, 0, 0, 0, 0, 0, 0, 0, 268, 42,
	43, 271, 0, 75, 44, 45, 46, 47, 0, 162,
	161, 0, 0, 82, 0, 0, 83, 0, 0, 48,
	0, 49, 0, 38, 37, 50, 10, 11, 12, 0,
	0, 0, 0, 39, 240, 51, 76, 77, 78, 40,
	41, 0, 0, 0, 81, 79, 80, 0, 0, 0,
	0, 42, 43, 0, 0, 0, 44, 45, 46, 47,
	0, 162, 161, 0, 0, 0, 0, 0, 0, 0,
	0, 48, 0, 49, 0, 38, 37, 50, 10, 11,
	12, 0, 0, 0, 0, 39, 199, 51, 0, 0,
	0, 40, 41, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 42, 43, 0, 0, 0, 44, 45,
	46, 47, 0, 162, 161, 0, 0, 0, 0, 0,
	0, 0, 0, 48, 0, 49, 0, 38, 37, 50,
	10, 11, 12, 0, 0, 0, 0, 39, 196, 51,
	0, 0, 0, 40, 41, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 42, 43, 0, 0, 0,
	44, 45, 46, 47, 0, 162, 161, 0, 0, 0,
	0, 0, 0, 0, 0, 48, 0, 49, 0, 38,
	37, 50, 10, 11, 12, 0, 0, 0, 0, 39,
	164, 51, 0, 0, 0, 40, 41, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 42, 43, 0,
	0, 0, 44, 45, 46, 47, 0, 162, 161, 0,
	0, 0, 0, 0, 0, 0, 0, 48, 0, 49,
	0, 38, 37, 50, 93, 92, 12, 88, 0, 0,
	94, 39, 0, 51, 0, 0, 0, 40, 41, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 42,
	43, 0, 0, 87, 44, 45, 46, 47, 0, 0,
	0, 0, 0, 10, 11, 12, 0, 0, 0, 48,
	39, 49, 51, 38, 37, 50, 40, 41, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 42, 43,
	0, 0, 0, 44, 45, 46, 47, 0, 162, 161,
	0, 0, 0, 0, 0, 0, 0, 0, 48, 0,
	49, 0, 38, 37, 50, 10, 11, 12, 0, 0,
	0, 94, 39, 0, 51, 10, 11, 12, 40, 41,
	0, 0, 39, 0, 51, 0, 0, 0, 40, 41,
	42, 43, 0, 0, 0, 44, 45, 46, 47, 0,
	42, 43, 0, 0, 0, 44, 45, 46, 47, 0,
	48, 0, 49, 0, 38, 37, 50, 0, 0, 0,
	48, 0, 49, 0, 38, 37, 50, 10, 183, 12,
	0, 0, 0, 0, 39, 0, 51, 0, 0, 0,
	40, 41, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 42, 43, 0, 0, 0, 44, 45, 46,
	47, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 48, 0, 49, 0, 38, 37, 50,
}
var RubyPact = []int{

	-50, -10, -1000, -51, -1000, -1000, 700, 29, -1000, -19,
	-1000, 29, 29, 29, 29, 29, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 11, 116, 29,
	29, 29, -1000, -1000, 29, 29, 29, 29, 29, 29,
	-1000, 29, -1000, 369, 134, 599, 39, 37, 36, 30,
	-1000, -1000, 113, -1000, -1000, 131, 29, 29, 700, 700,
	700, 700, 130, -1000, 700, 29, 29, 29, 29, 29,
	29, 29, 29, 29, 29, -1000, 29, 29, 29, 29,
	-19, -1000, 29, -1000, 29, 29, 29, 29, 29, -1000,
	-1000, 29, 24, 107, 21, -19, -19, -19, -19, 29,
	-1000, -1000, -3, -19, 700, 700, 700, 700, 700, 700,
	700, 700, 700, 110, 99, 700, 130, 97, 369, 547,
	-17, 700, 700, 700, 700, 110, -1000, 29, 29, -7,
	-1000, 12, -1000, 752, 287, -19, -19, -19, -19, -19,
	-19, -19, -19, -19, -1000, -1000, -1000, 29, -19, 29,
	29, -1000, -1000, -19, -1000, 29, 127, -19, -19, -19,
	-19, -1000, -1000, 495, 18, -1000, -8, 443, 29, -1000,
	29, 29, -19, -14, -1000, 90, 29, -1000, 690, 82,
	106, 10, -18, -1000, 29, 3, -1000, -1000, 93, -1000,
	121, 94, 84, 250, 29, -1000, 29, -1000, 700, 638,
	-19, -1000, -1000, -1000, -1000, -1000, -2, -1000, -1000, 76,
	-1000, -1000, 638, 391, 29, -1000, -1000, -3, -1000, -3,
	-1000, 42, 700, 700, 638, -19, 339, -1000, 29, 20,
	-1000, -1000, -33, -3, -34, -3, 29, -19, -19, 638,
	-1000, 119, -1000, -1000, 700, -1000, 124, 700, 638, -1000,
	-19, -9, -19, 117, 29, 41, 700, 29, 14, 700,
	-1000, 13, -1000,
}
var RubyPgo = []int{

	0, 213, 101, 210, 209, 208, 11, 206, 204, 203,
	201, 199, 198, 197, 194, 193, 191, 190, 4, 189,
	188, 186, 185, 180, 178, 177, 175, 174, 167, 165,
	55, 5, 159, 153, 152, 150, 149, 1, 148, 147,
	146, 145, 144, 0, 9, 3, 143,
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
	46, 11, 11, 11, 11, 33, 33, 33, 33, 33,
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
	2, 5, 6, 5, 5, 0, 5, 3, 4, 2,
}
var RubyChk = []int{

	-1000, -35, 57, -3, 41, 42, -43, 40, 57, -2,
	5, 6, 7, -16, -13, -9, -10, -14, -15, -17,
	-12, -4, -7, -20, -21, -22, -23, -24, -25, -26,
	-27, -28, -29, -8, -5, -19, -11, 55, 54, 12,
	18, 19, 30, 31, 35, 36, 37, 38, 50, 52,
	56, 14, -43, -43, 44, -43, -43, -43, -43, -43,
	6, 7, 55, 6, 7, -43, -43, -43, -43, -43,
	-43, -43, -43, -43, -43, 4, 37, 38, 39, 46,
	47, 45, 14, 17, 6, -31, -41, 34, 8, -40,
	-2, -6, 6, 5, 11, 34, 34, 34, 34, 6,
	7, 6, -18, -43, -18, -2, -2, -2, -2, -37,
	6, 5, -44, -2, -43, -43, -43, -43, -43, -43,
	-43, -43, -43, -43, -43, -43, -43, -43, -43, -30,
	-43, -43, -43, -43, -43, -43, 41, 32, 7, -42,
	41, -43, 41, -43, -30, -2, -2, -2, -2, -2,
	-2, -2, -2, -2, -31, 6, 5, 10, -2, -37,
	10, 41, 40, -2, 13, -32, 45, -2, -2, -2,
	-2, -34, -31, -30, -18, -43, 43, -30, 10, 51,
	-36, -39, -2, 6, 13, -33, 16, 15, -43, -43,
	-43, -43, -38, 6, -37, -45, 13, 41, 43, 13,
	-43, -43, -43, -43, 43, 13, 16, 15, -43, -30,
	-2, -6, 9, 5, 6, -6, -46, 41, 45, -43,
	40, 41, -30, -30, 7, 5, 6, -44, 10, -44,
	10, 34, -43, -43, -30, -2, -30, 41, 10, -45,
	13, -43, -43, -44, -43, -44, 33, -2, -2, -30,
	13, -43, 13, 53, -43, 53, -43, -43, -30, 6,
	-2, 6, -2, -43, 43, 34, -43, 33, -2, -43,
	10, -2, 10,
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
	12, 12, 12, 12, -2, 43, -2, 12, 12, -2,
	-2, 51, -2, 16, -2, 12, 12, 12, 12, 110,
	111, 12, 0, 0, 0, -2, -2, -2, -2, 12,
	60, 61, 12, -2, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 59, 0, 0, 0,
	0, 0, 0, 0, 0, 68, 8, 12, -2, 0,
	8, 0, 15, 97, 0, -2, -2, -2, -2, -2,
	-2, -2, -2, -2, 45, 54, 55, 12, -2, 12,
	12, 9, 10, -2, 113, 12, -2, -2, -2, -2,
	-2, 64, 69, 0, 0, 73, 0, 0, 12, 94,
	12, 12, 12, -2, 121, 0, 12, 8, 0, 0,
	0, 0, 0, -2, 12, 8, 70, 8, 0, 72,
	0, 14, 14, 0, 12, 122, 12, 8, 0, 129,
	-2, 53, 48, 56, 57, 58, 8, 119, 115, 0,
	65, 66, 64, 0, -2, 62, 63, 12, 14, 12,
	14, 0, 0, 0, 127, -2, 0, 120, 12, 0,
	71, 74, 0, 12, 0, 12, 12, -2, -2, 128,
	114, 0, 67, 95, 0, 96, 0, 0, 126, -2,
	12, 0, -2, 0, 12, 0, 0, 12, -2, 0,
	103, -2, 100,
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
		//line parser.y:149
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:151
		{
			Statements = []ast.Node{}
		}
	case 3:
		//line parser.y:153
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 4:
		//line parser.y:159
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 5:
		//line parser.y:165
		{
			RubyVAL.genericValue = nil
		}
	case 6:
		//line parser.y:166
		{
			RubyVAL.genericValue = nil
		}
	case 7:
		//line parser.y:167
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 8:
		//line parser.y:170
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 9:
		//line parser.y:172
		{ /* do nothing */
		}
	case 10:
		//line parser.y:174
		{ /* do nothing */
		}
	case 11:
		//line parser.y:176
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
		//line parser.y:188
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 44:
		//line parser.y:195
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 45:
		//line parser.y:202
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 46:
		//line parser.y:210
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 47:
		//line parser.y:217
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 48:
		//line parser.y:226
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 49:
		//line parser.y:228
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 50:
		//line parser.y:231
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 51:
		//line parser.y:233
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 52:
		//line parser.y:235
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 53:
		//line parser.y:237
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 54:
		//line parser.y:240
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 55:
		//line parser.y:242
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 56:
		//line parser.y:244
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 57:
		//line parser.y:246
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 58:
		//line parser.y:248
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 59:
		//line parser.y:250
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 60:
		//line parser.y:252
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 61:
		//line parser.y:254
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 62:
		//line parser.y:256
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 63:
		//line parser.y:258
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:268
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 68:
		//line parser.y:276
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 69:
		//line parser.y:277
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 70:
		//line parser.y:280
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 71:
		//line parser.y:287
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 72:
		//line parser.y:297
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 73:
		//line parser.y:306
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 74:
		//line parser.y:312
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 75:
		//line parser.y:320
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 76:
		//line parser.y:324
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 77:
		//line parser.y:329
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 78:
		//line parser.y:336
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 79:
		//line parser.y:343
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 80:
		//line parser.y:350
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 81:
		//line parser.y:357
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 82:
		//line parser.y:364
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 83:
		//line parser.y:365
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 84:
		//line parser.y:366
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 85:
		//line parser.y:367
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 86:
		//line parser.y:370
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 87:
		//line parser.y:379
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 88:
		//line parser.y:388
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 89:
		//line parser.y:397
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 90:
		//line parser.y:406
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 91:
		//line parser.y:415
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 92:
		//line parser.y:423
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 93:
		//line parser.y:424
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 94:
		//line parser.y:426
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 95:
		//line parser.y:430
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 96:
		//line parser.y:438
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 97:
		//line parser.y:446
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 98:
		//line parser.y:448
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 99:
		//line parser.y:450
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 100:
		//line parser.y:452
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 101:
		//line parser.y:455
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 102:
		//line parser.y:462
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 103:
		//line parser.y:469
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 104:
		//line parser.y:477
		{
			RubyVAL.genericValue = nil
		}
	case 105:
		//line parser.y:478
		{
			RubyVAL.genericValue = nil
		}
	case 106:
		//line parser.y:481
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 107:
		//line parser.y:483
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 108:
		//line parser.y:486
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 109:
		//line parser.y:488
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 110:
		//line parser.y:491
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 111:
		//line parser.y:493
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 112:
		//line parser.y:496
		{
			RubyVAL.genericValue = ast.FileNameConstReference{}
		}
	case 113:
		//line parser.y:499
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 114:
		//line parser.y:501
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 115:
		//line parser.y:509
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 116:
		//line parser.y:511
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 117:
		//line parser.y:513
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 118:
		//line parser.y:515
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 121:
		//line parser.y:520
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 122:
		//line parser.y:527
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 123:
		//line parser.y:535
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 124:
		//line parser.y:542
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 125:
		//line parser.y:549
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 126:
		//line parser.y:551
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 127:
		//line parser.y:558
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 128:
		//line parser.y:565
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 129:
		//line parser.y:572
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	}
	goto Rubystack /* stack new state and value */
}
