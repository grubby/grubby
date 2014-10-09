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

//line parser.y:818

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 26,
	44, 95,
	-2, 20,
	-1, 81,
	9, 67,
	10, 67,
	-2, 158,
	-1, 92,
	44, 95,
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
	-1, 107,
	44, 95,
	-2, 93,
	-1, 118,
	44, 95,
	-2, 20,
	-1, 205,
	9, 67,
	10, 67,
	-2, 158,
	-1, 242,
	44, 96,
	-2, 94,
}

const RubyNprod = 168
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1281

var RubyAct = []int{

	9, 83, 173, 171, 89, 161, 34, 82, 24, 117,
	118, 170, 197, 127, 96, 195, 22, 263, 64, 162,
	80, 2, 3, 264, 193, 75, 88, 247, 75, 121,
	122, 287, 241, 66, 56, 57, 4, 202, 27, 133,
	197, 244, 177, 106, 108, 116, 62, 61, 280, 197,
	283, 282, 79, 58, 223, 97, 196, 55, 54, 129,
	262, 63, 76, 246, 93, 194, 74, 136, 137, 74,
	140, 141, 142, 143, 65, 130, 146, 147, 123, 197,
	197, 10, 129, 93, 197, 155, 156, 103, 102, 101,
	292, 233, 88, 110, 64, 153, 64, 93, 130, 99,
	129, 130, 234, 163, 294, 93, 93, 90, 93, 93,
	93, 93, 75, 165, 93, 93, 130, 175, 109, 154,
	93, 188, 77, 93, 93, 78, 100, 268, 195, 255,
	93, 257, 256, 200, 286, 239, 88, 203, 93, 221,
	131, 132, 159, 204, 64, 99, 76, 251, 221, 138,
	219, 195, 174, 74, 172, 207, 206, 195, 148, 208,
	210, 188, 261, 131, 212, 188, 131, 157, 191, 192,
	242, 93, 107, 90, 93, 285, 222, 188, 105, 188,
	104, 131, 188, 188, 226, 270, 24, 91, 92, 81,
	174, 86, 96, 93, 124, 125, 236, 216, 93, 180,
	158, 152, 145, 135, 66, 278, 129, 235, 240, 60,
	87, 199, 56, 57, 126, 188, 85, 90, 188, 243,
	215, 190, 130, 28, 250, 24, 91, 92, 45, 198,
	1, 84, 229, 97, 93, 55, 54, 119, 67, 68,
	69, 211, 267, 52, 93, 65, 72, 70, 71, 94,
	188, 56, 57, 51, 188, 188, 50, 29, 49, 48,
	188, 188, 47, 18, 17, 16, 15, 38, 94, 13,
	58, 12, 59, 23, 55, 54, 188, 188, 188, 11,
	93, 21, 94, 95, 14, 188, 19, 131, 290, 188,
	94, 94, 32, 94, 94, 94, 94, 31, 33, 94,
	94, 30, 95, 0, 0, 94, 0, 0, 94, 94,
	0, 66, 220, 0, 0, 94, 95, 224, 0, 0,
	189, 0, 5, 94, 95, 95, 93, 95, 95, 95,
	95, 66, 0, 95, 95, 0, 0, 0, 0, 95,
	237, 238, 95, 95, 73, 67, 68, 69, 0, 95,
	0, 120, 65, 72, 70, 71, 94, 95, 209, 94,
	111, 112, 113, 114, 115, 67, 68, 69, 0, 0,
	0, 0, 65, 72, 70, 71, 0, 0, 94, 0,
	0, 265, 0, 94, 0, 134, 0, 20, 0, 139,
	95, 271, 0, 95, 144, 272, 0, 0, 149, 150,
	151, 0, 0, 0, 0, 0, 0, 0, 281, 0,
	0, 0, 95, 98, 0, 0, 0, 95, 0, 94,
	0, 166, 167, 168, 169, 0, 0, 0, 0, 94,
	179, 291, 293, 0, 295, 0, 296, 0, 0, 0,
	183, 0, 0, 0, 0, 0, 98, 0, 160, 164,
	0, 0, 0, 95, 0, 98, 0, 0, 176, 0,
	178, 0, 0, 95, 98, 94, 0, 181, 182, 98,
	0, 0, 98, 98, 0, 0, 0, 0, 0, 98,
	0, 0, 0, 0, 0, 0, 0, 98, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 95,
	0, 0, 0, 0, 0, 0, 0, 24, 25, 26,
	45, 94, 0, 214, 35, 217, 43, 0, 0, 44,
	36, 37, 0, 98, 0, 46, 0, 0, 0, 0,
	0, 0, 53, 56, 57, 0, 0, 0, 39, 40,
	41, 42, 0, 0, 252, 95, 0, 98, 0, 0,
	0, 258, 58, 0, 59, 0, 55, 54, 0, 0,
	0, 266, 0, 0, 0, 269, 0, 0, 0, 24,
	201, 118, 249, 0, 0, 0, 253, 274, 254, 0,
	0, 0, 0, 259, 279, 0, 260, 0, 0, 0,
	0, 0, 0, 98, 0, 56, 57, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 197, 289, 0, 275,
	276, 0, 0, 277, 58, 0, 59, 0, 55, 54,
	24, 25, 26, 45, 0, 0, 284, 35, 228, 43,
	231, 230, 44, 36, 37, 0, 0, 288, 46, 0,
	0, 0, 0, 0, 0, 53, 56, 57, 0, 0,
	0, 39, 40, 41, 42, 0, 0, 186, 187, 24,
	25, 26, 45, 0, 0, 58, 35, 59, 43, 55,
	54, 44, 36, 37, 0, 0, 0, 46, 0, 0,
	0, 0, 0, 0, 53, 56, 57, 0, 0, 0,
	39, 40, 41, 42, 0, 0, 6, 7, 0, 0,
	0, 0, 0, 0, 58, 0, 59, 0, 55, 54,
	0, 8, 24, 25, 26, 45, 0, 0, 0, 35,
	273, 43, 0, 0, 44, 36, 37, 0, 0, 0,
	46, 0, 0, 0, 0, 0, 0, 53, 56, 57,
	0, 0, 0, 39, 40, 41, 42, 0, 0, 186,
	187, 24, 25, 26, 45, 0, 0, 58, 35, 59,
	43, 55, 54, 44, 36, 37, 0, 0, 0, 46,
	0, 0, 0, 0, 0, 0, 53, 56, 57, 0,
	0, 0, 39, 40, 41, 42, 0, 0, 186, 187,
	0, 0, 0, 0, 0, 0, 58, 0, 59, 248,
	55, 54, 24, 25, 26, 45, 0, 0, 0, 35,
	245, 43, 0, 0, 44, 36, 37, 0, 0, 0,
	46, 0, 0, 0, 0, 0, 0, 53, 56, 57,
	0, 0, 0, 39, 40, 41, 42, 0, 0, 186,
	187, 0, 0, 0, 0, 0, 0, 58, 0, 59,
	0, 55, 54, 24, 25, 26, 45, 0, 0, 0,
	35, 232, 43, 0, 0, 44, 36, 37, 0, 0,
	0, 46, 0, 0, 0, 0, 0, 0, 53, 56,
	57, 0, 0, 0, 39, 40, 41, 42, 0, 0,
	186, 187, 0, 0, 0, 0, 0, 0, 58, 0,
	59, 0, 55, 54, 24, 25, 26, 45, 0, 0,
	0, 35, 227, 43, 0, 0, 44, 36, 37, 0,
	0, 0, 46, 0, 0, 0, 0, 0, 0, 53,
	56, 57, 0, 0, 0, 39, 40, 41, 42, 0,
	0, 186, 187, 0, 0, 0, 0, 0, 0, 58,
	0, 59, 0, 55, 54, 24, 25, 26, 45, 0,
	0, 0, 35, 225, 43, 0, 0, 44, 36, 37,
	0, 0, 0, 46, 0, 0, 0, 0, 0, 0,
	53, 56, 57, 0, 0, 0, 39, 40, 41, 42,
	0, 0, 186, 187, 24, 25, 26, 45, 0, 0,
	58, 35, 59, 43, 55, 54, 44, 36, 37, 0,
	0, 0, 46, 0, 0, 0, 0, 0, 0, 53,
	56, 57, 0, 0, 0, 39, 40, 41, 42, 0,
	0, 186, 187, 0, 0, 0, 0, 0, 0, 58,
	0, 59, 218, 55, 54, 24, 25, 26, 45, 0,
	0, 0, 35, 213, 43, 0, 0, 44, 36, 37,
	0, 0, 0, 46, 0, 0, 0, 0, 0, 0,
	53, 56, 57, 0, 0, 0, 39, 40, 41, 42,
	0, 0, 186, 187, 24, 25, 26, 45, 0, 0,
	58, 35, 59, 43, 55, 54, 44, 36, 37, 0,
	0, 0, 46, 0, 0, 0, 0, 0, 0, 53,
	56, 57, 0, 0, 0, 39, 40, 41, 42, 0,
	0, 186, 187, 24, 25, 26, 45, 185, 0, 58,
	35, 59, 43, 55, 54, 44, 36, 37, 0, 0,
	0, 46, 0, 0, 0, 0, 0, 0, 53, 56,
	57, 0, 0, 0, 39, 40, 41, 42, 0, 0,
	0, 184, 0, 24, 91, 92, 45, 0, 58, 96,
	59, 0, 55, 54, 24, 91, 92, 205, 0, 0,
	96, 0, 0, 0, 0, 0, 0, 0, 0, 56,
	57, 24, 91, 92, 81, 0, 0, 96, 0, 0,
	56, 57, 128, 91, 92, 45, 0, 0, 58, 0,
	97, 0, 55, 54, 0, 0, 0, 56, 57, 58,
	0, 97, 0, 55, 54, 0, 0, 0, 56, 57,
	24, 117, 118, 0, 0, 0, 84, 0, 97, 0,
	55, 54, 0, 0, 0, 0, 0, 58, 0, 59,
	0, 55, 54, 0, 0, 0, 56, 57, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 58, 0, 59, 0, 55,
	54,
}
var RubyPact = []int{

	-21, 654, -1000, -1000, -1000, 4, -1000, -1000, -1000, 327,
	108, -1000, -1000, -1000, 35, -1000, -1000, -1000, -1000, -1000,
	-25, -1000, -1000, -1000, -1000, 181, 91, 54, 53, 52,
	-1000, -1000, -1000, -1000, -1000, 174, 165, 165, 83, 502,
	502, 502, 502, 502, 1225, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 23, 188, -1000, -1000, 1197, -1000,
	-5, -1000, -1000, -1000, 502, 197, 1225, 220, 502, 1225,
	1225, 1225, 1225, 502, 196, 1225, 220, 502, 502, 502,
	195, 220, -1000, 109, 1197, 220, 194, 132, 29, -1000,
	21, 1186, 137, -1000, -1000, -1000, -27, -27, -25, 220,
	502, 502, 502, 502, 146, 146, 9, -1000, -1000, 502,
	193, 82, 82, 82, 82, 82, -1000, -1000, -1000, 1118,
	1079, -1000, -1000, 162, -1000, -1000, 14, 5, -1000, 200,
	-1000, 24, 564, -7, 82, 1169, -1000, 29, 21, 82,
	-1000, -1000, -1000, -1000, 82, -1000, -1000, 29, 21, 82,
	82, 82, -1000, 147, 3, 307, 29, 21, -1000, 1158,
	1040, -1000, 191, -1000, 989, 141, 82, 82, 82, 82,
	-1000, 129, 184, -1000, 19, -1000, 950, 165, 899, 82,
	-1000, 615, 848, 82, -1000, -1000, -1000, -1000, 327, 82,
	78, -1000, -1000, 202, -1000, 1225, -1000, -1000, -1000, 125,
	204, -12, 163, 109, -1000, 220, -1000, -1000, -1000, 6,
	29, 21, -1000, -1000, 797, 17, -1000, 746, -1000, -1000,
	-2, 184, 138, 502, -2, -1000, -1000, -1000, -1000, 116,
	502, -1000, -1000, -1000, 155, -1000, -1000, 7, -30, -1000,
	502, 1225, -1000, 118, 502, -1000, -1000, 179, -1000, 1079,
	-1000, -1000, 82, 1079, 707, -1000, 502, -1000, 82, 1079,
	1079, 201, -1000, 502, -1000, 42, 82, -1000, -1000, 82,
	-1000, 38, 37, -1000, 82, 1079, 1079, 1079, 169, 130,
	-13, -2, -1000, -1000, 1079, -1000, 502, 1225, 1079, 80,
	94, -2, -1000, -2, -1000, -2, -2,
}
var RubyPgo = []int{

	0, 320, 301, 298, 4, 297, 292, 387, 257, 286,
	284, 281, 0, 223, 81, 279, 273, 16, 271, 38,
	269, 6, 2, 267, 266, 265, 264, 263, 262, 259,
	258, 256, 253, 243, 351, 237, 7, 5, 232, 230,
	11, 229, 221, 13, 220, 214, 211, 210, 1, 3,
	209, 141,
}
var RubyR1 = []int{

	0, 39, 39, 39, 39, 39, 39, 39, 39, 39,
	39, 51, 51, 34, 34, 34, 34, 34, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 17,
	17, 17, 17, 17, 17, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 36, 36, 45, 45, 43, 43, 43,
	43, 43, 48, 48, 48, 48, 47, 47, 47, 47,
	47, 15, 15, 40, 40, 22, 22, 49, 49, 49,
	18, 18, 20, 21, 21, 50, 50, 10, 10, 10,
	10, 10, 10, 10, 8, 8, 19, 19, 13, 13,
	23, 23, 24, 25, 26, 27, 28, 28, 28, 28,
	29, 30, 31, 32, 33, 2, 5, 6, 6, 3,
	3, 41, 41, 41, 41, 46, 46, 46, 4, 4,
	4, 4, 37, 44, 44, 44, 9, 9, 9, 9,
	9, 9, 9, 9, 38, 38, 38, 38, 35, 35,
	35, 7, 11, 42, 42, 42, 42, 16,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	4, 4, 2, 3, 3, 4, 4, 3, 2, 3,
	3, 4, 6, 3, 1, 1, 3, 0, 1, 1,
	1, 3, 1, 1, 3, 3, 1, 1, 3, 3,
	3, 7, 7, 1, 3, 1, 3, 0, 1, 3,
	4, 6, 4, 1, 4, 1, 4, 3, 3, 3,
	3, 3, 3, 3, 2, 2, 2, 2, 3, 3,
	3, 3, 2, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 1, 1, 3, 3, 5,
	5, 0, 4, 7, 8, 3, 7, 8, 3, 4,
	4, 3, 3, 0, 1, 3, 4, 5, 3, 3,
	3, 3, 3, 4, 4, 3, 3, 2, 0, 2,
	2, 3, 4, 0, 3, 4, 6, 1,
}
var RubyChk = []int{

	-1000, -39, 42, 43, 57, -1, 42, 43, 57, -12,
	-14, -15, -18, -20, -10, -24, -25, -26, -27, -9,
	-7, -11, -17, -16, 5, 6, 7, -19, -13, -8,
	-2, -5, -6, -3, -21, 12, 18, 19, -23, 36,
	37, 38, 39, 14, 17, 8, 23, -28, -29, -30,
	-31, -32, -33, 30, 55, 54, 31, 32, 50, 52,
	-50, 43, 42, 57, 14, 45, 4, 38, 39, 40,
	47, 48, 46, 17, 45, 4, 38, 14, 17, 17,
	45, 8, -36, -48, 50, 35, 10, -47, -12, -4,
	-14, 6, 7, -19, -13, -8, 11, 52, -7, 8,
	35, 35, 35, 35, 6, 4, -21, 7, -21, 35,
	10, -1, -1, -1, -1, -1, -12, 6, 7, -35,
	-34, 6, 7, 55, 6, 7, -45, -43, 5, -12,
	-17, -14, -51, 44, -1, 6, -12, -12, -14, -1,
	-12, -12, -12, -12, -1, 6, -12, -12, -14, -1,
	-1, -1, 6, -43, 10, -12, -12, -14, 6, 10,
	-34, -37, 46, -37, -34, -43, -1, -1, -1, -1,
	-40, -49, 8, -22, 6, -40, -34, 33, -34, -1,
	6, -34, -34, -1, 43, 9, 42, 43, -12, -1,
	-42, 6, 7, 10, 51, 10, 51, 42, -41, -46,
	-12, 6, 44, -48, -36, 8, 9, -12, -4, 51,
	-12, -14, -4, 13, -34, -44, 6, -34, 53, 9,
	-51, 10, -49, 35, -51, 13, -21, 13, 13, -38,
	16, 15, 13, 13, 24, 5, -12, -51, -51, 10,
	4, 44, 7, -43, 35, 13, 46, 10, 53, -34,
	-22, 9, -1, -34, -34, 13, 16, 15, -1, -34,
	-34, 7, 53, 10, 53, -51, -1, -12, 9, -1,
	6, -51, -51, 13, -1, -34, -34, -34, 4, -1,
	6, -51, 13, 13, -34, 6, 4, 44, -34, -1,
	-12, -51, 10, -51, 10, -51, -51,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
	46, 47, 48, 49, 18, 19, -2, 21, 22, 23,
	24, 25, 26, 27, 28, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 158, 13, 29, 30, 31,
	32, 33, 34, 167, 0, 0, 125, 126, 67, 11,
	0, 5, 6, 7, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, -2, 52, 58, 67, 0, 0, 64, 72, 73,
	77, 19, -2, 21, 22, 23, 13, -2, 0, 67,
	0, 0, 0, 0, 87, 87, 13, -2, 13, 0,
	0, 112, 113, 114, 115, 13, 13, 19, -2, 0,
	163, 106, 107, 0, 104, 105, 0, 0, 18, 68,
	69, 70, 131, 0, 148, 53, 59, 116, 118, 120,
	121, 122, 123, 124, 150, 54, 60, 117, 119, 149,
	151, 152, 57, 0, 0, 68, 97, 98, 110, 0,
	0, 13, 143, 13, 0, 0, 99, 100, 101, 102,
	11, 83, 87, 88, 85, 11, 0, 0, 0, 103,
	111, 0, 0, 159, 160, 161, 14, 15, 16, 17,
	0, 108, 109, 0, 127, 0, 128, 12, 11, 11,
	0, 19, 0, 55, 56, -2, 50, 74, 75, 61,
	78, 79, 80, 138, 0, 0, 144, 0, 141, 51,
	13, 0, 0, 0, 13, 90, 13, 92, 146, 0,
	0, 13, 153, 162, 13, 66, 71, 0, 0, 11,
	0, 0, -2, 0, 0, 139, 142, 0, 140, 11,
	89, 84, 86, 11, 0, 147, 0, 13, 13, 157,
	164, 13, 129, 0, 130, 0, 11, 135, 63, 62,
	145, 0, 0, 91, 13, 155, 156, 165, 0, 0,
	0, 132, 81, 82, 154, 13, 0, 0, 166, 11,
	11, 133, 11, 136, 11, 134, 137,
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
		//line parser.y:160
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:162
		{
		}
	case 3:
		//line parser.y:164
		{
		}
	case 4:
		//line parser.y:166
		{
		}
	case 5:
		//line parser.y:168
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:170
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:172
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:178
		{
		}
	case 11:
		//line parser.y:180
		{
		}
	case 12:
		//line parser.y:181
		{
		}
	case 13:
		//line parser.y:184
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 14:
		//line parser.y:186
		{
		}
	case 15:
		//line parser.y:188
		{
		}
	case 16:
		//line parser.y:190
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 17:
		//line parser.y:192
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
		//line parser.y:202
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 51:
		//line parser.y:209
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 52:
		//line parser.y:216
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 53:
		//line parser.y:223
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 54:
		//line parser.y:230
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 55:
		//line parser.y:237
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 56:
		//line parser.y:245
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 57:
		//line parser.y:253
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 58:
		//line parser.y:263
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 59:
		//line parser.y:270
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 60:
		//line parser.y:278
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 61:
		//line parser.y:288
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 62:
		//line parser.y:298
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 63:
		//line parser.y:307
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 64:
		//line parser.y:309
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 65:
		//line parser.y:312
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 66:
		//line parser.y:314
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:316
		{
			RubyVAL.genericSlice = ast.Nodes{}
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
		//line parser.y:328
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:330
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:332
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 75:
		//line parser.y:334
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:337
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:339
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 78:
		//line parser.y:341
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 79:
		//line parser.y:343
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 80:
		//line parser.y:345
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 81:
		//line parser.y:351
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 82:
		//line parser.y:359
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 83:
		//line parser.y:368
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 84:
		//line parser.y:370
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 85:
		//line parser.y:373
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 86:
		//line parser.y:375
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 87:
		//line parser.y:377
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 88:
		//line parser.y:379
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 89:
		//line parser.y:383
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 90:
		//line parser.y:388
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 91:
		//line parser.y:395
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 92:
		//line parser.y:405
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 93:
		//line parser.y:414
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 94:
		//line parser.y:420
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 95:
		//line parser.y:428
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 96:
		//line parser.y:432
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 97:
		//line parser.y:437
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 98:
		//line parser.y:444
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 99:
		//line parser.y:451
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 100:
		//line parser.y:458
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 101:
		//line parser.y:465
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 102:
		//line parser.y:472
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 103:
		//line parser.y:479
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 104:
		//line parser.y:487
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 105:
		//line parser.y:489
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 106:
		//line parser.y:492
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 107:
		//line parser.y:494
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 108:
		//line parser.y:497
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 109:
		//line parser.y:499
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 110:
		//line parser.y:502
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 111:
		//line parser.y:504
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 112:
		//line parser.y:506
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 113:
		//line parser.y:507
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 114:
		//line parser.y:508
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 115:
		//line parser.y:509
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 116:
		//line parser.y:512
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 117:
		//line parser.y:520
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 118:
		//line parser.y:528
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 119:
		//line parser.y:536
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 120:
		//line parser.y:545
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 121:
		//line parser.y:554
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 122:
		//line parser.y:563
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 123:
		//line parser.y:572
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 124:
		//line parser.y:581
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 125:
		//line parser.y:589
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 126:
		//line parser.y:590
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 127:
		//line parser.y:592
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 128:
		//line parser.y:594
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 129:
		//line parser.y:597
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 130:
		//line parser.y:605
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 131:
		//line parser.y:613
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 132:
		//line parser.y:615
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 133:
		//line parser.y:622
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 134:
		//line parser.y:629
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 135:
		//line parser.y:637
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 136:
		//line parser.y:644
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 137:
		//line parser.y:651
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 138:
		//line parser.y:659
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 139:
		//line parser.y:661
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 140:
		//line parser.y:668
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 141:
		//line parser.y:675
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 142:
		//line parser.y:678
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 143:
		//line parser.y:680
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 144:
		//line parser.y:682
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 145:
		//line parser.y:684
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 146:
		//line parser.y:687
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 147:
		//line parser.y:694
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 148:
		//line parser.y:702
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 149:
		//line parser.y:709
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 150:
		//line parser.y:716
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 151:
		//line parser.y:723
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 152:
		//line parser.y:730
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 153:
		//line parser.y:737
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-2].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 154:
		//line parser.y:745
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 155:
		//line parser.y:752
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 156:
		//line parser.y:759
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 157:
		//line parser.y:766
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 158:
		//line parser.y:773
		{
		}
	case 159:
		//line parser.y:774
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 160:
		//line parser.y:775
		{
		}
	case 161:
		//line parser.y:778
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 162:
		//line parser.y:781
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 163:
		//line parser.y:789
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 164:
		//line parser.y:791
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 165:
		//line parser.y:793
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 166:
		//line parser.y:802
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
	}
	goto Rubystack /* stack new state and value */
}
