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
const YIELD = 57372
const TRUE = 57373
const FALSE = 57374
const LESSTHAN = 57375
const GREATERTHAN = 57376
const EQUALTO = 57377
const BANG = 57378
const COMPLEMENT = 57379
const POSITIVE = 57380
const NEGATIVE = 57381
const STAR = 57382
const WHITESPACE = 57383
const NEWLINE = 57384
const SEMICOLON = 57385
const COLON = 57386
const DOT = 57387
const PIPE = 57388
const SLASH = 57389
const AMPERSAND = 57390
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
	"YIELD",
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

//line parser.y:814

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 26,
	44, 97,
	-2, 20,
	-1, 81,
	9, 67,
	10, 67,
	-2, 159,
	-1, 92,
	44, 97,
	-2, 20,
	-1, 97,
	5, 11,
	6, 11,
	7, 11,
	10, 11,
	31, 11,
	32, 11,
	42, 11,
	50, 11,
	52, 11,
	53, 11,
	54, 11,
	55, 11,
	-2, 13,
	-1, 108,
	44, 97,
	-2, 95,
	-1, 119,
	44, 97,
	-2, 20,
	-1, 208,
	9, 67,
	10, 67,
	-2, 159,
	-1, 246,
	44, 98,
	-2, 96,
}

const RubyNprod = 169
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1264

var RubyAct = []int{

	9, 174, 172, 34, 162, 83, 82, 128, 90, 64,
	171, 200, 89, 66, 2, 3, 163, 122, 123, 80,
	198, 291, 268, 196, 245, 205, 88, 133, 200, 4,
	100, 267, 248, 251, 66, 225, 104, 62, 61, 27,
	107, 109, 75, 103, 79, 117, 175, 67, 68, 69,
	287, 102, 63, 180, 65, 72, 70, 71, 286, 129,
	212, 199, 75, 200, 197, 93, 124, 136, 137, 250,
	140, 141, 142, 143, 266, 65, 146, 147, 28, 200,
	176, 64, 154, 74, 93, 156, 158, 200, 177, 153,
	284, 296, 88, 111, 157, 64, 76, 298, 93, 155,
	100, 154, 164, 74, 94, 243, 93, 93, 166, 93,
	93, 93, 93, 272, 198, 93, 93, 178, 110, 255,
	223, 93, 191, 94, 93, 93, 200, 101, 237, 223,
	22, 93, 160, 203, 290, 132, 88, 94, 265, 238,
	93, 206, 207, 246, 64, 94, 94, 108, 94, 94,
	94, 94, 221, 198, 94, 94, 210, 209, 198, 289,
	94, 158, 191, 94, 94, 10, 191, 274, 211, 213,
	94, 227, 93, 214, 226, 93, 224, 194, 195, 94,
	191, 106, 191, 105, 230, 191, 191, 218, 175, 130,
	173, 98, 125, 126, 60, 93, 183, 239, 159, 240,
	93, 259, 152, 261, 260, 145, 135, 282, 244, 154,
	87, 94, 130, 202, 94, 130, 247, 191, 127, 217,
	191, 193, 176, 201, 131, 254, 66, 29, 1, 233,
	177, 130, 120, 138, 94, 121, 93, 52, 93, 94,
	51, 50, 148, 49, 48, 47, 271, 131, 93, 18,
	131, 98, 17, 95, 191, 16, 15, 98, 191, 191,
	67, 68, 69, 38, 191, 191, 131, 65, 72, 70,
	71, 13, 95, 12, 23, 94, 11, 94, 21, 14,
	191, 191, 191, 19, 32, 93, 95, 94, 31, 191,
	33, 30, 294, 191, 95, 95, 0, 95, 95, 95,
	95, 98, 0, 95, 95, 0, 192, 222, 5, 95,
	0, 0, 95, 95, 228, 0, 0, 0, 0, 95,
	0, 0, 0, 66, 94, 0, 98, 0, 95, 0,
	0, 93, 161, 165, 0, 0, 73, 241, 242, 130,
	0, 0, 0, 179, 0, 181, 112, 113, 114, 115,
	116, 0, 184, 185, 0, 0, 0, 67, 68, 69,
	95, 0, 0, 95, 65, 72, 70, 71, 0, 20,
	94, 134, 0, 0, 131, 139, 0, 0, 0, 269,
	144, 75, 0, 95, 149, 150, 151, 0, 95, 275,
	0, 77, 0, 276, 78, 99, 0, 0, 216, 0,
	219, 0, 0, 0, 0, 0, 285, 0, 167, 168,
	169, 170, 0, 0, 0, 76, 0, 182, 24, 91,
	92, 45, 74, 0, 95, 0, 95, 186, 99, 295,
	297, 0, 299, 0, 300, 0, 95, 99, 0, 0,
	0, 0, 0, 0, 56, 57, 99, 0, 0, 0,
	0, 99, 0, 0, 99, 99, 0, 0, 253, 0,
	0, 99, 0, 58, 257, 59, 258, 55, 54, 0,
	99, 263, 0, 95, 264, 0, 0, 24, 25, 26,
	45, 188, 0, 0, 35, 0, 43, 0, 0, 44,
	36, 37, 0, 0, 0, 46, 0, 279, 280, 0,
	0, 281, 53, 56, 57, 99, 0, 0, 39, 40,
	41, 42, 0, 0, 288, 187, 24, 204, 119, 95,
	0, 0, 58, 0, 59, 292, 55, 54, 0, 0,
	99, 0, 256, 0, 0, 0, 0, 0, 0, 0,
	0, 262, 56, 57, 0, 0, 0, 0, 0, 0,
	0, 270, 0, 200, 0, 273, 0, 0, 0, 0,
	0, 58, 0, 59, 0, 55, 54, 278, 0, 0,
	0, 0, 0, 0, 283, 0, 0, 0, 99, 0,
	24, 25, 26, 45, 0, 0, 0, 35, 232, 43,
	235, 234, 44, 36, 37, 0, 0, 293, 46, 0,
	0, 0, 0, 0, 0, 53, 56, 57, 0, 0,
	0, 39, 40, 41, 42, 0, 0, 189, 190, 24,
	25, 26, 45, 0, 0, 58, 35, 59, 43, 55,
	54, 44, 36, 37, 0, 0, 0, 46, 0, 0,
	0, 0, 0, 0, 53, 56, 57, 0, 0, 0,
	39, 40, 41, 42, 0, 0, 6, 7, 0, 0,
	0, 0, 0, 0, 58, 0, 59, 0, 55, 54,
	0, 8, 24, 25, 26, 45, 0, 0, 0, 35,
	277, 43, 0, 0, 44, 36, 37, 0, 0, 0,
	46, 0, 0, 0, 0, 0, 0, 53, 56, 57,
	0, 0, 0, 39, 40, 41, 42, 0, 0, 189,
	190, 24, 25, 26, 45, 0, 0, 58, 35, 59,
	43, 55, 54, 44, 36, 37, 0, 0, 0, 46,
	0, 0, 0, 0, 0, 0, 53, 56, 57, 0,
	0, 0, 39, 40, 41, 42, 0, 0, 189, 190,
	0, 0, 0, 0, 0, 0, 58, 0, 59, 252,
	55, 54, 24, 25, 26, 45, 0, 0, 0, 35,
	249, 43, 0, 0, 44, 36, 37, 0, 0, 0,
	46, 0, 0, 0, 0, 0, 0, 53, 56, 57,
	0, 0, 0, 39, 40, 41, 42, 0, 0, 189,
	190, 0, 0, 0, 0, 0, 0, 58, 0, 59,
	0, 55, 54, 24, 25, 26, 45, 0, 0, 0,
	35, 236, 43, 0, 0, 44, 36, 37, 0, 0,
	0, 46, 0, 0, 0, 0, 0, 0, 53, 56,
	57, 0, 0, 0, 39, 40, 41, 42, 0, 0,
	189, 190, 0, 0, 0, 0, 0, 0, 58, 0,
	59, 0, 55, 54, 24, 25, 26, 45, 0, 0,
	0, 35, 231, 43, 0, 0, 44, 36, 37, 0,
	0, 0, 46, 0, 0, 0, 0, 0, 0, 53,
	56, 57, 0, 0, 0, 39, 40, 41, 42, 0,
	0, 189, 190, 0, 0, 0, 0, 0, 0, 58,
	0, 59, 0, 55, 54, 24, 25, 26, 45, 0,
	0, 0, 35, 229, 43, 0, 0, 44, 36, 37,
	0, 0, 0, 46, 0, 0, 0, 0, 0, 0,
	53, 56, 57, 0, 0, 0, 39, 40, 41, 42,
	0, 0, 189, 190, 24, 25, 26, 45, 0, 0,
	58, 35, 59, 43, 55, 54, 44, 36, 37, 0,
	0, 0, 46, 0, 0, 0, 0, 0, 0, 53,
	56, 57, 0, 0, 0, 39, 40, 41, 42, 0,
	0, 189, 190, 0, 0, 0, 0, 0, 0, 58,
	0, 59, 220, 55, 54, 24, 25, 26, 45, 0,
	0, 0, 35, 215, 43, 0, 0, 44, 36, 37,
	0, 0, 0, 46, 0, 0, 0, 0, 0, 0,
	53, 56, 57, 0, 0, 0, 39, 40, 41, 42,
	0, 0, 189, 190, 24, 25, 26, 45, 0, 0,
	58, 35, 59, 43, 55, 54, 44, 36, 37, 0,
	0, 0, 46, 0, 0, 0, 0, 0, 0, 53,
	56, 57, 0, 0, 0, 39, 40, 41, 42, 0,
	0, 189, 190, 24, 25, 26, 45, 0, 0, 58,
	35, 59, 43, 55, 54, 44, 36, 37, 0, 0,
	0, 46, 0, 0, 0, 0, 0, 0, 53, 56,
	57, 0, 0, 0, 39, 40, 41, 42, 24, 91,
	92, 81, 0, 86, 96, 0, 0, 0, 58, 0,
	59, 0, 55, 54, 24, 91, 92, 45, 0, 0,
	96, 0, 0, 0, 56, 57, 0, 0, 85, 24,
	91, 92, 208, 0, 0, 96, 0, 0, 0, 0,
	56, 57, 0, 84, 0, 97, 0, 55, 54, 0,
	0, 0, 0, 0, 0, 56, 57, 0, 0, 58,
	0, 97, 0, 55, 54, 24, 91, 92, 81, 0,
	0, 96, 0, 0, 58, 0, 97, 0, 55, 54,
	24, 118, 119, 0, 0, 0, 96, 0, 0, 0,
	0, 56, 57, 24, 118, 119, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 56, 57, 0, 0,
	84, 0, 97, 0, 55, 54, 0, 0, 0, 56,
	57, 0, 0, 0, 0, 58, 0, 97, 0, 55,
	54, 0, 0, 0, 0, 0, 0, 0, 58, 0,
	59, 0, 55, 54,
}
var RubyPact = []int{

	-28, 614, -1000, -1000, -1000, -5, -1000, -1000, -1000, 319,
	377, -1000, -1000, -1000, 27, -1000, -1000, -1000, -1000, -1000,
	-26, -1000, -1000, -1000, -1000, 1113, 92, 16, 8, 1,
	-1000, -1000, -1000, -1000, -1000, 177, 140, 140, 83, 1078,
	1078, 1078, 1078, 1078, 1208, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 11, 186, -1000, -1000, 413, -1000,
	-17, -1000, -1000, -1000, 1078, 200, 1208, 413, 1078, 1208,
	1208, 1208, 1208, 1078, 199, 1208, 413, 1078, 1078, 1078,
	196, 413, -1000, 89, 413, 413, 192, 122, 30, -1000,
	-1000, 1180, 22, -1000, -1000, -1000, -30, -30, 38, -26,
	413, 1078, 1078, 1078, 1078, 182, 182, 20, -1000, -1000,
	1078, 190, 67, 67, 67, 67, 67, -1000, -1000, -1000,
	472, 1039, -1000, -1000, 171, -1000, -1000, 13, 10, 222,
	-1000, 58, 511, -19, 67, 1144, -1000, 30, 38, 67,
	-1000, -1000, -1000, -1000, 67, -1000, -1000, 30, 38, 67,
	67, 67, -1000, 148, 222, 1195, 9, -1000, 30, -1000,
	1129, 1000, -1000, 181, -1000, 949, 143, 67, 67, 67,
	67, -1000, 119, 40, -1000, 0, 168, 165, -1000, 910,
	140, 859, 67, -1000, 575, 808, 67, -1000, -1000, -1000,
	-1000, 319, 67, 115, -1000, -1000, 1208, -1000, 1208, -1000,
	-1000, -1000, 95, 204, -20, 136, 89, -1000, 413, -1000,
	-1000, -1000, -3, -1000, -1000, -1000, 757, 23, -1000, 706,
	-1000, -1000, -14, 40, 110, 1078, -1000, -1000, -14, -1000,
	-1000, -1000, -1000, 188, 1078, -1000, -1000, -1000, 131, -1000,
	-1000, 21, -31, -1000, 1078, 1208, -1000, 104, 1078, -1000,
	-1000, 161, -1000, 1039, -1000, -1000, 67, 1039, 667, -1000,
	1078, -1000, 67, 1039, 1039, 203, -1000, 1078, -1000, 84,
	67, -1000, -1000, 67, -1000, 45, 37, -1000, 67, 1039,
	1039, 1039, 153, 130, -23, -14, -1000, -1000, 1039, -1000,
	1078, 1208, 1039, 81, 87, -14, -1000, -14, -1000, -14,
	-14,
}
var RubyPgo = []int{

	0, 306, 291, 290, 12, 288, 284, 369, 227, 283,
	279, 278, 0, 78, 165, 276, 274, 130, 273, 1,
	39, 271, 3, 8, 263, 256, 255, 252, 249, 245,
	244, 243, 241, 240, 237, 235, 232, 6, 4, 229,
	228, 10, 223, 221, 7, 219, 218, 213, 210, 5,
	2, 194, 135,
}
var RubyR1 = []int{

	0, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 52, 52, 35, 35, 35, 35, 35, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 17,
	17, 17, 17, 17, 17, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 37, 37, 46, 46, 44, 44, 44,
	44, 44, 49, 49, 49, 49, 23, 23, 48, 48,
	48, 15, 15, 41, 41, 19, 19, 19, 19, 50,
	50, 50, 18, 18, 21, 22, 22, 51, 51, 10,
	10, 10, 10, 10, 10, 8, 8, 20, 20, 13,
	13, 24, 24, 25, 26, 27, 28, 29, 29, 29,
	29, 30, 31, 32, 33, 34, 2, 5, 6, 6,
	3, 3, 42, 42, 42, 42, 47, 47, 47, 4,
	4, 4, 4, 38, 45, 45, 45, 9, 9, 9,
	9, 9, 9, 9, 9, 39, 39, 39, 39, 36,
	36, 36, 7, 11, 43, 43, 43, 43, 16,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	4, 4, 2, 3, 3, 4, 4, 3, 2, 3,
	3, 4, 6, 3, 1, 1, 3, 0, 1, 1,
	1, 3, 1, 1, 3, 3, 1, 1, 1, 3,
	3, 7, 7, 1, 3, 1, 2, 3, 2, 0,
	1, 3, 4, 6, 4, 1, 4, 1, 4, 3,
	3, 3, 3, 3, 3, 2, 2, 2, 2, 3,
	3, 3, 3, 2, 2, 2, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 1, 1, 3, 3,
	5, 5, 0, 4, 7, 8, 3, 7, 8, 3,
	4, 4, 3, 3, 0, 1, 3, 4, 5, 3,
	3, 3, 3, 3, 4, 4, 3, 3, 2, 0,
	2, 2, 3, 4, 0, 3, 4, 6, 1,
}
var RubyChk = []int{

	-1000, -40, 42, 43, 57, -1, 42, 43, 57, -12,
	-14, -15, -18, -21, -10, -25, -26, -27, -28, -9,
	-7, -11, -17, -16, 5, 6, 7, -20, -13, -8,
	-2, -5, -6, -3, -22, 12, 18, 19, -24, 36,
	37, 38, 39, 14, 17, 8, 23, -29, -30, -31,
	-32, -33, -34, 30, 55, 54, 31, 32, 50, 52,
	-51, 43, 42, 57, 14, 45, 4, 38, 39, 40,
	47, 48, 46, 17, 45, 4, 38, 14, 17, 17,
	45, 8, -37, -49, 50, 35, 10, -48, -12, -4,
	-23, 6, 7, -20, -13, -8, 11, 52, -14, -7,
	8, 35, 35, 35, 35, 6, 4, -22, 7, -22,
	35, 10, -1, -1, -1, -1, -1, -12, 6, 7,
	-36, -35, 6, 7, 55, 6, 7, -46, -44, -12,
	-17, -14, -52, 44, -1, 6, -12, -12, -14, -1,
	-12, -12, -12, -12, -1, 6, -12, -12, -14, -1,
	-1, -1, 6, -44, -12, 10, -12, -23, -12, 6,
	10, -35, -38, 46, -38, -35, -44, -1, -1, -1,
	-1, -41, -50, 8, -19, 6, 40, 48, -41, -35,
	33, -35, -1, 6, -35, -35, -1, 43, 9, 42,
	43, -12, -1, -43, 6, 7, 10, 51, 10, 51,
	42, -42, -47, -12, 6, 44, -49, -37, 8, 9,
	-12, -4, 51, -23, -4, 13, -35, -45, 6, -35,
	53, 9, -52, 10, -50, 35, 6, 6, -52, 13,
	-22, 13, 13, -39, 16, 15, 13, 13, 24, -12,
	-12, -52, -52, 10, 4, 44, 7, -44, 35, 13,
	46, 10, 53, -35, -19, 9, -1, -35, -35, 13,
	16, 15, -1, -35, -35, 7, 53, 10, 53, -52,
	-1, -12, 9, -1, 6, -52, -52, 13, -1, -35,
	-35, -35, 4, -1, 6, -52, 13, 13, -35, 6,
	4, 44, -35, -1, -12, -52, 10, -52, 10, -52,
	-52,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
	46, 47, 48, 49, 18, 19, -2, 21, 22, 23,
	24, 25, 26, 27, 28, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 159, 13, 29, 30, 31,
	32, 33, 34, 168, 0, 0, 126, 127, 67, 11,
	0, 5, 6, 7, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, -2, 52, 58, 67, 0, 0, 64, 72, 73,
	78, 19, -2, 21, 22, 23, 13, -2, 77, 0,
	67, 0, 0, 0, 0, 89, 89, 13, -2, 13,
	0, 0, 113, 114, 115, 116, 13, 13, 19, -2,
	0, 164, 107, 108, 0, 105, 106, 0, 0, 65,
	69, 70, 132, 0, 149, 53, 59, 117, 119, 121,
	122, 123, 124, 125, 151, 54, 60, 118, 120, 150,
	152, 153, 57, 0, 68, 0, 65, 99, 76, 111,
	0, 0, 13, 144, 13, 0, 0, 100, 101, 102,
	103, 11, 83, 89, 90, 85, 0, 0, 11, 0,
	0, 0, 104, 112, 0, 0, 160, 161, 162, 14,
	15, 16, 17, 0, 109, 110, 0, 128, 0, 129,
	12, 11, 11, 0, 19, 0, 55, 56, -2, 50,
	74, 75, 61, 79, 80, 139, 0, 0, 145, 0,
	142, 51, 13, 0, 0, 0, 86, 88, 13, 92,
	13, 94, 147, 0, 0, 13, 154, 163, 13, 66,
	71, 0, 0, 11, 0, 0, -2, 0, 0, 140,
	143, 0, 141, 11, 91, 84, 87, 11, 0, 148,
	0, 13, 13, 158, 165, 13, 130, 0, 131, 0,
	11, 136, 63, 62, 146, 0, 0, 93, 13, 156,
	157, 166, 0, 0, 0, 133, 81, 82, 155, 13,
	0, 0, 167, 11, 11, 134, 11, 137, 11, 135,
	138,
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
		//line parser.y:161
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:163
		{
		}
	case 3:
		//line parser.y:165
		{
		}
	case 4:
		//line parser.y:167
		{
		}
	case 5:
		//line parser.y:169
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:171
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:173
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:179
		{
		}
	case 11:
		//line parser.y:181
		{
		}
	case 12:
		//line parser.y:182
		{
		}
	case 13:
		//line parser.y:185
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 14:
		//line parser.y:187
		{
		}
	case 15:
		//line parser.y:189
		{
		}
	case 16:
		//line parser.y:191
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 17:
		//line parser.y:193
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 50:
		//line parser.y:203
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 51:
		//line parser.y:210
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 52:
		//line parser.y:217
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 53:
		//line parser.y:224
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 54:
		//line parser.y:231
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 55:
		//line parser.y:238
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 56:
		//line parser.y:246
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 57:
		//line parser.y:254
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 58:
		//line parser.y:264
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 59:
		//line parser.y:271
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 60:
		//line parser.y:279
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 61:
		//line parser.y:289
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 62:
		//line parser.y:299
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 63:
		//line parser.y:308
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 64:
		//line parser.y:310
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 65:
		//line parser.y:313
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 66:
		//line parser.y:315
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:317
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 68:
		//line parser.y:319
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 69:
		//line parser.y:321
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 70:
		//line parser.y:323
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 71:
		//line parser.y:325
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 72:
		//line parser.y:329
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:331
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:333
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 75:
		//line parser.y:335
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 77:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 78:
		//line parser.y:340
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 79:
		//line parser.y:342
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 80:
		//line parser.y:344
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 81:
		//line parser.y:350
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 82:
		//line parser.y:358
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 83:
		//line parser.y:367
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 84:
		//line parser.y:369
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 85:
		//line parser.y:372
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 86:
		//line parser.y:374
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 87:
		//line parser.y:376
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 88:
		//line parser.y:378
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 89:
		//line parser.y:380
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 90:
		//line parser.y:382
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 91:
		//line parser.y:386
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 92:
		//line parser.y:391
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 93:
		//line parser.y:398
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 94:
		//line parser.y:408
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 95:
		//line parser.y:417
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 96:
		//line parser.y:423
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 97:
		//line parser.y:431
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 98:
		//line parser.y:435
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 99:
		//line parser.y:440
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 100:
		//line parser.y:447
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 101:
		//line parser.y:454
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 102:
		//line parser.y:461
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 103:
		//line parser.y:468
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 104:
		//line parser.y:475
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 105:
		//line parser.y:483
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 106:
		//line parser.y:485
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 107:
		//line parser.y:488
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 108:
		//line parser.y:490
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 109:
		//line parser.y:493
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 110:
		//line parser.y:495
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 111:
		//line parser.y:498
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 112:
		//line parser.y:500
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 113:
		//line parser.y:502
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 114:
		//line parser.y:503
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 115:
		//line parser.y:504
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 116:
		//line parser.y:505
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 117:
		//line parser.y:508
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 118:
		//line parser.y:516
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 119:
		//line parser.y:524
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 120:
		//line parser.y:532
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 121:
		//line parser.y:541
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 122:
		//line parser.y:550
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 123:
		//line parser.y:559
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 124:
		//line parser.y:568
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 125:
		//line parser.y:577
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 126:
		//line parser.y:585
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 127:
		//line parser.y:586
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 128:
		//line parser.y:588
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 129:
		//line parser.y:590
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 130:
		//line parser.y:593
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 131:
		//line parser.y:601
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 132:
		//line parser.y:609
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 133:
		//line parser.y:611
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 134:
		//line parser.y:618
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 135:
		//line parser.y:625
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 136:
		//line parser.y:633
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 137:
		//line parser.y:640
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 138:
		//line parser.y:647
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 139:
		//line parser.y:655
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 140:
		//line parser.y:657
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 141:
		//line parser.y:664
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 142:
		//line parser.y:671
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 143:
		//line parser.y:674
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 144:
		//line parser.y:676
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 145:
		//line parser.y:678
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 146:
		//line parser.y:680
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 147:
		//line parser.y:683
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 148:
		//line parser.y:690
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 149:
		//line parser.y:698
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 150:
		//line parser.y:705
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 151:
		//line parser.y:712
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 152:
		//line parser.y:719
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 153:
		//line parser.y:726
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 154:
		//line parser.y:733
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-2].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 155:
		//line parser.y:741
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 156:
		//line parser.y:748
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 157:
		//line parser.y:755
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 158:
		//line parser.y:762
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 159:
		//line parser.y:769
		{
		}
	case 160:
		//line parser.y:770
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 161:
		//line parser.y:771
		{
		}
	case 162:
		//line parser.y:774
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 163:
		//line parser.y:777
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 164:
		//line parser.y:785
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 165:
		//line parser.y:787
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 166:
		//line parser.y:789
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 167:
		//line parser.y:798
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
	case 168:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	}
	goto Rubystack /* stack new state and value */
}
