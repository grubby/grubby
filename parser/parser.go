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
const SPECIAL_CHAR_REF = 57349
const CAPITAL_REF = 57350
const LPAREN = 57351
const RPAREN = 57352
const COMMA = 57353
const DO = 57354
const DEF = 57355
const END = 57356
const IF = 57357
const ELSE = 57358
const ELSIF = 57359
const UNLESS = 57360
const CLASS = 57361
const MODULE = 57362
const FOR = 57363
const WHILE = 57364
const UNTIL = 57365
const BEGIN = 57366
const RESCUE = 57367
const ENSURE = 57368
const BREAK = 57369
const REDO = 57370
const RETRY = 57371
const RETURN = 57372
const YIELD = 57373
const TRUE = 57374
const FALSE = 57375
const LESSTHAN = 57376
const GREATERTHAN = 57377
const EQUALTO = 57378
const BANG = 57379
const COMPLEMENT = 57380
const POSITIVE = 57381
const NEGATIVE = 57382
const STAR = 57383
const WHITESPACE = 57384
const NEWLINE = 57385
const SEMICOLON = 57386
const COLON = 57387
const DOUBLECOLON = 57388
const DOT = 57389
const PIPE = 57390
const SLASH = 57391
const AMPERSAND = 57392
const QUESTIONMARK = 57393
const CARET = 57394
const LBRACKET = 57395
const RBRACKET = 57396
const LBRACE = 57397
const RBRACE = 57398
const DOLLARSIGN = 57399
const ATSIGN = 57400
const FILE_CONST_REF = 57401
const EOF = 57402

var RubyToknames = []string{
	"OPERATOR",
	"NODE",
	"REF",
	"SPECIAL_CHAR_REF",
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
	"DOUBLECOLON",
	"DOT",
	"PIPE",
	"SLASH",
	"AMPERSAND",
	"QUESTIONMARK",
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

//line parser.y:826

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 24,
	46, 99,
	-2, 20,
	-1, 80,
	10, 72,
	11, 72,
	-2, 158,
	-1, 90,
	46, 99,
	-2, 20,
	-1, 96,
	13, 13,
	15, 13,
	18, 13,
	19, 13,
	20, 13,
	24, 13,
	31, 13,
	37, 13,
	38, 13,
	39, 13,
	40, 13,
	44, 13,
	-2, 11,
	-1, 111,
	46, 99,
	-2, 97,
	-1, 202,
	46, 100,
	-2, 98,
	-1, 208,
	10, 72,
	11, 72,
	-2, 158,
}

const RubyNprod = 169
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1406

var RubyAct = []int{

	9, 174, 128, 32, 172, 82, 88, 81, 20, 66,
	171, 156, 2, 3, 122, 268, 123, 197, 70, 157,
	195, 252, 107, 132, 87, 300, 103, 104, 293, 4,
	269, 286, 197, 289, 249, 193, 226, 64, 63, 102,
	25, 110, 112, 101, 288, 180, 120, 197, 68, 69,
	79, 100, 175, 175, 65, 173, 99, 66, 251, 129,
	267, 67, 197, 196, 91, 78, 124, 130, 197, 136,
	137, 138, 139, 197, 141, 142, 143, 144, 194, 146,
	150, 149, 114, 148, 151, 152, 91, 176, 176, 130,
	87, 97, 130, 274, 195, 244, 177, 177, 149, 91,
	160, 224, 238, 70, 166, 167, 130, 113, 158, 91,
	91, 91, 91, 239, 91, 91, 91, 91, 98, 91,
	178, 91, 188, 70, 91, 91, 154, 298, 245, 97,
	91, 66, 200, 68, 69, 266, 87, 260, 91, 262,
	261, 206, 291, 207, 91, 91, 67, 256, 224, 105,
	78, 211, 106, 68, 69, 214, 188, 212, 68, 69,
	188, 215, 292, 191, 209, 192, 67, 111, 103, 104,
	78, 67, 91, 66, 202, 78, 91, 125, 225, 126,
	188, 102, 188, 60, 231, 188, 188, 222, 195, 210,
	195, 91, 276, 228, 240, 91, 241, 109, 227, 108,
	219, 183, 87, 170, 22, 89, 61, 90, 62, 149,
	248, 247, 165, 153, 135, 284, 86, 130, 188, 199,
	127, 188, 218, 190, 198, 1, 255, 34, 234, 133,
	52, 56, 57, 51, 91, 131, 91, 50, 49, 48,
	47, 17, 91, 16, 15, 14, 271, 273, 39, 91,
	91, 94, 58, 12, 59, 188, 55, 54, 11, 188,
	188, 21, 10, 19, 13, 188, 188, 18, 33, 35,
	30, 29, 31, 94, 28, 0, 26, 0, 0, 0,
	0, 0, 188, 188, 188, 0, 94, 91, 0, 0,
	0, 188, 0, 0, 296, 188, 94, 94, 94, 94,
	92, 94, 94, 94, 94, 0, 94, 0, 94, 0,
	0, 94, 94, 0, 0, 0, 0, 94, 0, 0,
	0, 0, 92, 0, 0, 94, 0, 0, 0, 0,
	0, 94, 94, 0, 91, 92, 0, 22, 89, 61,
	90, 80, 0, 0, 95, 92, 92, 92, 92, 0,
	92, 92, 92, 92, 0, 92, 0, 92, 0, 94,
	92, 92, 0, 94, 56, 57, 92, 0, 121, 0,
	0, 27, 0, 0, 92, 0, 0, 246, 94, 0,
	92, 92, 94, 0, 0, 83, 0, 96, 0, 55,
	54, 0, 0, 0, 0, 93, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 223, 92, 0,
	0, 0, 92, 0, 229, 0, 0, 93, 0, 0,
	0, 94, 0, 94, 0, 0, 0, 92, 0, 94,
	93, 92, 0, 0, 242, 243, 94, 94, 0, 0,
	93, 93, 93, 93, 0, 93, 93, 93, 93, 0,
	93, 0, 93, 0, 0, 93, 93, 189, 0, 5,
	0, 93, 0, 0, 155, 159, 0, 0, 0, 93,
	92, 0, 92, 0, 94, 93, 93, 0, 92, 179,
	270, 181, 0, 0, 0, 92, 92, 0, 184, 185,
	277, 0, 0, 0, 278, 0, 0, 0, 115, 116,
	117, 118, 119, 93, 0, 0, 0, 93, 287, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 94, 93, 92, 134, 217, 93, 220, 0, 0,
	140, 297, 299, 0, 301, 145, 302, 147, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 161, 162, 163, 164,
	0, 0, 0, 168, 169, 93, 0, 93, 0, 0,
	92, 182, 0, 93, 0, 22, 89, 61, 90, 62,
	93, 93, 95, 0, 0, 0, 22, 89, 61, 90,
	208, 203, 254, 95, 0, 0, 0, 0, 258, 0,
	259, 0, 56, 57, 0, 264, 0, 0, 265, 0,
	0, 0, 0, 56, 57, 0, 0, 0, 93, 0,
	0, 0, 0, 58, 0, 96, 0, 55, 54, 0,
	0, 281, 282, 0, 58, 283, 96, 0, 55, 54,
	0, 0, 22, 23, 61, 24, 62, 0, 0, 290,
	36, 233, 44, 236, 235, 45, 37, 38, 0, 0,
	294, 46, 0, 0, 0, 93, 0, 0, 53, 56,
	57, 0, 0, 0, 40, 41, 42, 43, 0, 0,
	186, 187, 0, 0, 257, 0, 0, 0, 0, 0,
	58, 0, 59, 263, 55, 54, 22, 23, 61, 24,
	62, 0, 0, 272, 36, 0, 44, 275, 0, 45,
	37, 38, 0, 0, 0, 46, 0, 0, 0, 280,
	0, 0, 53, 56, 57, 0, 285, 0, 40, 41,
	42, 43, 0, 0, 6, 7, 0, 0, 0, 0,
	0, 0, 0, 0, 58, 0, 59, 0, 55, 54,
	295, 8, 22, 23, 61, 24, 62, 0, 0, 0,
	36, 279, 44, 0, 0, 45, 37, 38, 0, 0,
	0, 46, 0, 0, 0, 0, 0, 0, 53, 56,
	57, 0, 0, 0, 40, 41, 42, 43, 0, 0,
	186, 187, 0, 22, 23, 61, 24, 62, 0, 0,
	58, 36, 59, 44, 55, 54, 45, 37, 38, 0,
	0, 0, 46, 0, 0, 0, 0, 0, 0, 53,
	56, 57, 0, 0, 0, 40, 41, 42, 43, 0,
	0, 186, 187, 0, 0, 0, 0, 0, 0, 0,
	0, 58, 0, 59, 253, 55, 54, 22, 23, 61,
	24, 62, 0, 0, 0, 36, 250, 44, 0, 0,
	45, 37, 38, 0, 0, 0, 46, 0, 0, 0,
	0, 0, 0, 53, 56, 57, 0, 0, 0, 40,
	41, 42, 43, 0, 0, 186, 187, 0, 0, 0,
	0, 0, 0, 0, 0, 58, 0, 59, 0, 55,
	54, 22, 23, 61, 24, 62, 0, 0, 0, 36,
	237, 44, 0, 0, 45, 37, 38, 0, 0, 0,
	46, 0, 0, 0, 0, 0, 0, 53, 56, 57,
	0, 0, 0, 40, 41, 42, 43, 0, 0, 186,
	187, 0, 0, 0, 0, 0, 0, 0, 0, 58,
	0, 59, 0, 55, 54, 22, 23, 61, 24, 62,
	0, 0, 0, 36, 232, 44, 0, 0, 45, 37,
	38, 0, 0, 0, 46, 0, 0, 0, 0, 0,
	0, 53, 56, 57, 0, 0, 0, 40, 41, 42,
	43, 0, 0, 186, 187, 0, 0, 0, 0, 0,
	0, 0, 0, 58, 0, 59, 0, 55, 54, 22,
	23, 61, 24, 62, 0, 0, 0, 36, 230, 44,
	0, 0, 45, 37, 38, 0, 0, 0, 46, 0,
	0, 0, 0, 0, 0, 53, 56, 57, 0, 0,
	0, 40, 41, 42, 43, 0, 0, 186, 187, 0,
	22, 23, 61, 24, 62, 0, 0, 58, 36, 59,
	44, 55, 54, 45, 37, 38, 0, 0, 0, 46,
	0, 0, 0, 0, 0, 0, 53, 56, 57, 0,
	0, 0, 40, 41, 42, 43, 0, 0, 186, 187,
	0, 0, 0, 0, 0, 0, 0, 0, 58, 0,
	59, 221, 55, 54, 22, 23, 61, 24, 62, 0,
	0, 0, 36, 216, 44, 0, 0, 45, 37, 38,
	0, 0, 0, 46, 0, 0, 0, 0, 0, 0,
	53, 56, 57, 0, 0, 0, 40, 41, 42, 43,
	0, 0, 186, 187, 0, 22, 23, 61, 24, 62,
	0, 0, 58, 36, 59, 44, 55, 54, 45, 37,
	38, 0, 0, 0, 46, 0, 0, 0, 0, 0,
	0, 53, 56, 57, 0, 0, 0, 40, 41, 42,
	43, 0, 0, 186, 187, 0, 22, 23, 61, 24,
	62, 205, 0, 58, 36, 59, 44, 55, 54, 45,
	37, 38, 0, 0, 0, 46, 0, 0, 0, 0,
	0, 0, 53, 56, 57, 0, 0, 0, 40, 41,
	42, 43, 0, 0, 0, 204, 0, 22, 23, 61,
	24, 62, 0, 0, 58, 36, 59, 44, 55, 54,
	45, 37, 38, 0, 0, 0, 46, 0, 0, 0,
	0, 0, 0, 53, 56, 57, 0, 0, 0, 40,
	41, 42, 43, 22, 89, 61, 90, 80, 0, 85,
	95, 0, 0, 0, 0, 58, 0, 59, 0, 55,
	54, 22, 201, 61, 90, 62, 0, 0, 0, 0,
	56, 57, 0, 0, 84, 22, 89, 61, 90, 80,
	0, 0, 95, 0, 0, 0, 0, 0, 56, 57,
	0, 83, 0, 96, 0, 55, 54, 0, 0, 197,
	70, 0, 56, 57, 0, 0, 0, 0, 0, 58,
	0, 59, 0, 55, 54, 0, 0, 0, 70, 0,
	0, 0, 0, 83, 0, 96, 0, 55, 54, 0,
	68, 69, 77, 0, 0, 71, 72, 73, 70, 0,
	0, 0, 0, 67, 76, 74, 75, 78, 68, 69,
	213, 0, 0, 71, 72, 73, 0, 0, 0, 0,
	0, 67, 76, 74, 75, 78, 0, 0, 68, 69,
	0, 0, 0, 71, 72, 73, 0, 0, 0, 0,
	0, 67, 76, 74, 75, 78,
}
var RubyPact = []int{

	-31, 691, -1000, -1000, -1000, -6, -1000, -1000, -1000, 1334,
	-1000, -1000, -1000, 32, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 1258, 82, 20, 15, 7, -1000, -1000,
	-1000, -1000, -1000, -1000, 134, -25, 193, 159, 159, 71,
	1222, 1222, 1222, 1222, 1222, 199, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 8, 171, -1000, -1000, 199, -1000,
	-23, -1000, -1000, -1000, -1000, -1000, 1222, 208, 199, 199,
	199, 199, 1222, 199, 199, 199, 199, 1222, 199, 1222,
	199, -1000, 69, 199, 199, 207, 115, 99, -1000, 1290,
	120, -1000, -1000, -1000, -8, -29, -29, 199, 1222, 1222,
	1222, 1222, 206, 199, 199, 1222, 1222, 197, 46, 46,
	11, -1000, -1000, 1222, 195, 42, 42, 42, 42, 42,
	99, 1140, -1000, -1000, 157, -1000, -1000, 24, 9, 1354,
	-1000, 1276, 166, 1181, 42, 581, 99, 99, 99, 99,
	42, 99, 99, 99, 99, 42, 119, 42, 179, 1354,
	570, 1316, 99, -1000, 570, 1099, -1000, 194, -1000, 1045,
	177, 42, 42, 42, 42, -1000, 99, 99, 42, 42,
	-1000, -1000, 90, 47, -1000, 0, 192, 187, -1000, 1004,
	159, 950, 42, -1000, 637, 896, -1000, -1000, 1334, 42,
	88, -1000, -1000, 199, -1000, 199, -1000, -1000, -1000, 84,
	124, 332, -1000, 42, -1000, -1000, 69, -1000, 199, 199,
	-1000, 99, -1000, -2, 99, -1000, -1000, 842, 10, -1000,
	788, -1000, -1000, -11, 47, 137, 1222, -1000, -1000, -11,
	-1000, -1000, -1000, -1000, 123, 1222, -1000, -1000, -1000, 127,
	99, 99, 4, -26, -1000, 1222, 199, 83, 99, 1222,
	-1000, -1000, 186, -1000, 1140, -1000, -1000, 42, 1140, 747,
	-1000, 1222, -1000, 42, 1140, 1140, 211, -1000, 1222, -1000,
	25, 1334, 42, 99, -1000, 42, -1000, 30, 19, -1000,
	42, 1140, 1140, 1140, 136, 158, -17, -11, -1000, -1000,
	1140, -1000, 1222, 199, 1140, 116, 14, -11, -1000, -11,
	-1000, -11, -11,
}
var RubyPgo = []int{

	0, 457, 274, 272, 6, 271, 270, 269, 371, 268,
	267, 264, 263, 0, 276, 227, 262, 261, 8, 258,
	1, 40, 253, 3, 248, 245, 244, 243, 241, 240,
	239, 238, 237, 233, 230, 368, 229, 7, 11, 228,
	225, 10, 224, 223, 2, 222, 220, 219, 216, 5,
	4, 183, 235,
}
var RubyR1 = []int{

	0, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 52, 52, 35, 35, 35, 35, 35, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 18, 18, 18, 18, 18, 18, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 37, 37,
	46, 46, 44, 44, 44, 44, 49, 49, 49, 49,
	48, 48, 48, 16, 16, 41, 41, 20, 20, 20,
	20, 50, 50, 50, 19, 19, 22, 23, 23, 51,
	51, 11, 11, 11, 11, 11, 11, 8, 8, 21,
	21, 14, 14, 24, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 2, 5, 6, 6, 3,
	3, 42, 42, 42, 42, 47, 47, 47, 4, 4,
	4, 4, 38, 45, 45, 45, 10, 10, 10, 10,
	10, 10, 10, 10, 39, 39, 39, 39, 36, 36,
	36, 7, 12, 43, 43, 43, 43, 17, 9,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 4, 1, 4, 2, 3, 3, 4, 4, 3,
	2, 3, 3, 3, 3, 3, 4, 6, 3, 1,
	1, 3, 0, 1, 1, 3, 1, 1, 3, 3,
	1, 3, 3, 7, 7, 1, 3, 1, 2, 3,
	2, 0, 1, 3, 4, 6, 4, 1, 3, 1,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 2,
	2, 3, 3, 3, 3, 2, 2, 2, 2, 3,
	3, 3, 3, 3, 3, 1, 1, 3, 3, 5,
	5, 0, 4, 7, 8, 3, 7, 8, 3, 4,
	4, 3, 3, 0, 1, 3, 4, 5, 3, 3,
	3, 3, 3, 4, 4, 3, 3, 2, 0, 2,
	2, 3, 4, 0, 3, 4, 6, 1, 5,
}
var RubyChk = []int{

	-1000, -40, 43, 44, 60, -1, 43, 44, 60, -13,
	-16, -19, -22, -11, -25, -26, -27, -28, -10, -12,
	-18, -17, 5, 6, 8, -21, -14, -8, -2, -5,
	-6, -3, -23, -9, -15, -7, 13, 19, 20, -24,
	37, 38, 39, 40, 15, 18, 24, -29, -30, -31,
	-32, -33, -34, 31, 58, 57, 32, 33, 53, 55,
	-51, 7, 9, 44, 43, 60, 15, 47, 34, 35,
	4, 39, 40, 41, 49, 50, 48, 18, 51, 18,
	9, -37, -49, 53, 36, 11, -48, -13, -4, 6,
	8, -21, -14, -8, -15, 12, 55, 9, 36, 36,
	36, 36, 47, 34, 35, 15, 18, 47, 6, 4,
	-23, 8, -23, 36, 11, -1, -1, -1, -1, -1,
	-13, -35, 6, 8, 58, 6, 8, -46, -44, -13,
	-18, -52, 46, -36, -1, 6, -13, -13, -13, -13,
	-1, -13, -13, -13, -13, -1, -13, -1, -44, -13,
	11, -13, -13, 6, 11, -35, -38, 48, -38, -35,
	-44, -1, -1, -1, -1, 6, -13, -13, -1, -1,
	6, -41, -50, 9, -20, 6, 41, 50, -41, -35,
	34, -35, -1, 6, -35, -35, 43, 44, -13, -1,
	-43, 6, 8, 11, 54, 11, 54, 43, -42, -47,
	-13, 6, 8, -1, 44, 10, -49, -37, 9, 45,
	10, -13, -4, 54, -13, -4, 14, -35, -45, 6,
	-35, 56, 10, -52, 11, -50, 36, 6, 6, -52,
	14, -23, 14, 14, -39, 17, 16, 14, 14, 25,
	-13, -13, -52, -52, 11, 4, 45, -44, -13, 36,
	14, 48, 11, 56, -35, -20, 10, -1, -35, -35,
	14, 17, 16, -1, -35, -35, 8, 56, 11, 56,
	-52, -13, -1, -13, 10, -1, 6, -52, -52, 14,
	-1, -35, -35, -35, 4, -1, 6, -52, 14, 14,
	-35, 6, 4, 45, -35, -1, -13, -52, 11, -52,
	11, -52, -52,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 38,
	39, 40, 41, 42, 43, 44, 45, 46, 47, 48,
	49, 50, 18, 19, -2, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 13, 32, 33, 34,
	35, 36, 37, 167, 0, 0, 125, 126, 72, 11,
	0, 52, 158, 5, 6, 7, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	-2, 54, 60, 72, 0, 0, 69, 76, 77, 19,
	-2, 21, 22, 23, 30, 13, -2, 72, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 91, 91,
	13, -2, 13, 0, 0, 115, 116, 117, 118, 13,
	13, 163, 109, 110, 0, 107, 108, 0, 0, 70,
	74, 131, 0, 0, 148, 55, 61, 63, 65, 119,
	120, 121, 122, 123, 124, 150, 0, 152, 0, 73,
	0, 70, 101, 113, 0, 0, 13, 143, 13, 0,
	0, 102, 103, 104, 105, 56, 62, 64, 149, 151,
	59, 11, 85, 91, 92, 87, 0, 0, 11, 0,
	0, 0, 106, 114, 0, 0, 14, 15, 16, 17,
	0, 111, 112, 0, 127, 0, 128, 12, 11, 11,
	0, 19, -2, 159, 160, 161, 57, 58, -2, 0,
	51, 78, 79, 66, 81, 82, 138, 0, 0, 144,
	0, 141, 53, 13, 0, 0, 0, 88, 90, 13,
	94, 13, 96, 146, 0, 0, 13, 153, 162, 13,
	71, 75, 0, 0, 11, 0, 0, 0, 168, 0,
	139, 142, 0, 140, 11, 93, 86, 89, 11, 0,
	147, 0, 13, 13, 157, 164, 13, 129, 0, 130,
	0, 38, 11, 135, 68, 67, 145, 0, 0, 95,
	13, 155, 156, 165, 0, 0, 0, 132, 83, 84,
	154, 13, 0, 0, 166, 11, 11, 133, 11, 136,
	11, 134, 137,
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
	52, 53, 54, 55, 56, 57, 58, 59, 60,
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
		//line parser.y:164
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:166
		{
		}
	case 3:
		//line parser.y:168
		{
		}
	case 4:
		//line parser.y:170
		{
		}
	case 5:
		//line parser.y:172
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:174
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:176
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:182
		{
		}
	case 11:
		//line parser.y:184
		{
		}
	case 12:
		//line parser.y:185
		{
		}
	case 13:
		//line parser.y:188
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 14:
		//line parser.y:190
		{
		}
	case 15:
		//line parser.y:192
		{
		}
	case 16:
		//line parser.y:194
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 17:
		//line parser.y:196
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 51:
		//line parser.y:206
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 52:
		//line parser.y:213
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 53:
		//line parser.y:217
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 54:
		//line parser.y:224
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 55:
		//line parser.y:231
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 56:
		//line parser.y:238
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 57:
		//line parser.y:245
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 58:
		//line parser.y:253
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 59:
		//line parser.y:261
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 60:
		//line parser.y:271
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 61:
		//line parser.y:278
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 62:
		//line parser.y:286
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 63:
		//line parser.y:294
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 64:
		//line parser.y:302
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 65:
		//line parser.y:310
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 66:
		//line parser.y:320
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 67:
		//line parser.y:330
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 68:
		//line parser.y:339
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 69:
		//line parser.y:341
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 70:
		//line parser.y:344
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 71:
		//line parser.y:346
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 72:
		//line parser.y:348
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 73:
		//line parser.y:350
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:352
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 75:
		//line parser.y:354
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:358
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:360
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 78:
		//line parser.y:362
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 79:
		//line parser.y:364
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 80:
		//line parser.y:367
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 81:
		//line parser.y:369
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 82:
		//line parser.y:371
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 83:
		//line parser.y:377
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 84:
		//line parser.y:385
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 85:
		//line parser.y:394
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 86:
		//line parser.y:396
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 87:
		//line parser.y:399
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 88:
		//line parser.y:401
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 89:
		//line parser.y:403
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 90:
		//line parser.y:405
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 91:
		//line parser.y:407
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 92:
		//line parser.y:409
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 93:
		//line parser.y:413
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 94:
		//line parser.y:418
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 95:
		//line parser.y:425
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 96:
		//line parser.y:435
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 97:
		//line parser.y:444
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 98:
		//line parser.y:450
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 99:
		//line parser.y:458
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 100:
		//line parser.y:462
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 101:
		//line parser.y:467
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 102:
		//line parser.y:474
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 103:
		//line parser.y:481
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 104:
		//line parser.y:488
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 105:
		//line parser.y:495
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 106:
		//line parser.y:502
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 107:
		//line parser.y:510
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 108:
		//line parser.y:512
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 109:
		//line parser.y:515
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 110:
		//line parser.y:517
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 111:
		//line parser.y:520
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 112:
		//line parser.y:522
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 113:
		//line parser.y:525
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 114:
		//line parser.y:527
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 115:
		//line parser.y:529
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 116:
		//line parser.y:530
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 117:
		//line parser.y:531
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 118:
		//line parser.y:532
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 119:
		//line parser.y:535
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 120:
		//line parser.y:544
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 121:
		//line parser.y:553
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 122:
		//line parser.y:562
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 123:
		//line parser.y:571
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 124:
		//line parser.y:580
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 125:
		//line parser.y:588
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 126:
		//line parser.y:589
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 127:
		//line parser.y:591
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 128:
		//line parser.y:593
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 129:
		//line parser.y:596
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 130:
		//line parser.y:604
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 131:
		//line parser.y:612
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 132:
		//line parser.y:614
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 133:
		//line parser.y:621
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 134:
		//line parser.y:628
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 135:
		//line parser.y:636
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 136:
		//line parser.y:643
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 137:
		//line parser.y:650
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 138:
		//line parser.y:658
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 139:
		//line parser.y:660
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 140:
		//line parser.y:667
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 141:
		//line parser.y:674
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 142:
		//line parser.y:677
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 143:
		//line parser.y:679
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 144:
		//line parser.y:681
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 145:
		//line parser.y:683
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 146:
		//line parser.y:686
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 147:
		//line parser.y:693
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 148:
		//line parser.y:701
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 149:
		//line parser.y:708
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 150:
		//line parser.y:715
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 151:
		//line parser.y:722
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 152:
		//line parser.y:729
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 153:
		//line parser.y:736
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-2].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 154:
		//line parser.y:744
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 155:
		//line parser.y:751
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 156:
		//line parser.y:758
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 157:
		//line parser.y:765
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 158:
		//line parser.y:772
		{
		}
	case 159:
		//line parser.y:773
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 160:
		//line parser.y:774
		{
		}
	case 161:
		//line parser.y:777
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 162:
		//line parser.y:780
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 163:
		//line parser.y:788
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 164:
		//line parser.y:790
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 165:
		//line parser.y:792
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 166:
		//line parser.y:801
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
	case 167:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 168:
		//line parser.y:818
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	}
	goto Rubystack /* stack new state and value */
}
