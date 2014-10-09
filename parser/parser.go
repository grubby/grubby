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

const RubyLast = 1286

var RubyAct = []int{

	9, 171, 173, 34, 89, 83, 82, 170, 64, 161,
	2, 3, 197, 66, 121, 122, 22, 195, 162, 80,
	193, 287, 127, 264, 241, 4, 88, 202, 75, 247,
	132, 197, 244, 75, 223, 103, 62, 61, 27, 66,
	106, 108, 79, 294, 177, 116, 263, 67, 68, 69,
	283, 63, 282, 102, 65, 72, 70, 71, 196, 128,
	209, 194, 76, 123, 93, 246, 280, 135, 136, 74,
	139, 140, 141, 142, 74, 129, 145, 146, 197, 197,
	65, 197, 153, 93, 101, 155, 156, 110, 233, 262,
	99, 255, 88, 257, 256, 64, 154, 93, 129, 234,
	153, 129, 197, 239, 152, 93, 93, 163, 93, 93,
	93, 93, 109, 175, 93, 93, 129, 100, 286, 221,
	93, 188, 165, 93, 93, 292, 159, 10, 64, 64,
	93, 131, 200, 268, 195, 88, 251, 221, 93, 99,
	203, 204, 24, 91, 92, 81, 261, 86, 96, 219,
	195, 206, 195, 90, 174, 207, 172, 191, 192, 208,
	210, 188, 124, 125, 212, 188, 242, 107, 56, 57,
	93, 285, 85, 93, 222, 270, 105, 188, 104, 188,
	174, 226, 188, 188, 216, 180, 130, 84, 158, 97,
	151, 55, 54, 93, 235, 137, 236, 75, 93, 60,
	144, 134, 278, 240, 147, 87, 153, 77, 199, 130,
	78, 126, 130, 157, 215, 188, 190, 198, 188, 90,
	1, 229, 129, 28, 250, 119, 52, 130, 243, 51,
	66, 76, 93, 50, 93, 49, 48, 47, 74, 18,
	17, 16, 267, 73, 93, 15, 38, 13, 12, 94,
	188, 23, 11, 21, 188, 188, 14, 19, 32, 31,
	188, 188, 90, 33, 67, 68, 69, 30, 94, 0,
	0, 65, 72, 70, 71, 29, 188, 188, 188, 0,
	93, 0, 94, 0, 0, 188, 0, 211, 290, 188,
	94, 94, 0, 94, 94, 94, 94, 0, 0, 94,
	94, 95, 220, 0, 0, 94, 0, 224, 94, 94,
	0, 0, 0, 0, 0, 94, 0, 0, 0, 0,
	95, 0, 0, 94, 0, 0, 93, 0, 0, 0,
	237, 238, 0, 130, 95, 0, 0, 0, 0, 24,
	201, 118, 95, 95, 0, 95, 95, 95, 95, 0,
	0, 95, 95, 0, 0, 94, 0, 95, 94, 66,
	95, 95, 0, 120, 0, 56, 57, 95, 0, 0,
	0, 265, 189, 0, 5, 95, 197, 0, 94, 0,
	0, 271, 0, 94, 58, 272, 59, 0, 55, 54,
	0, 0, 0, 67, 68, 69, 0, 0, 281, 0,
	65, 72, 70, 71, 0, 0, 0, 95, 0, 0,
	95, 0, 111, 112, 113, 114, 115, 94, 20, 94,
	0, 291, 293, 0, 295, 0, 296, 0, 0, 94,
	95, 0, 0, 0, 0, 95, 0, 133, 0, 0,
	0, 138, 0, 0, 98, 0, 143, 0, 0, 0,
	148, 149, 150, 0, 0, 0, 0, 0, 0, 0,
	160, 164, 0, 0, 0, 94, 0, 0, 0, 95,
	176, 95, 178, 166, 167, 168, 169, 98, 0, 181,
	182, 95, 179, 0, 0, 0, 98, 24, 91, 92,
	45, 0, 183, 96, 0, 98, 0, 0, 0, 0,
	98, 0, 0, 98, 98, 0, 0, 0, 0, 0,
	98, 94, 0, 56, 57, 0, 0, 95, 98, 0,
	0, 0, 0, 0, 0, 214, 0, 217, 0, 0,
	0, 0, 58, 0, 97, 0, 55, 54, 0, 0,
	0, 0, 24, 25, 26, 45, 0, 0, 0, 35,
	0, 43, 0, 98, 44, 36, 37, 0, 0, 0,
	46, 0, 0, 95, 0, 0, 0, 53, 56, 57,
	0, 0, 0, 39, 40, 41, 42, 0, 98, 6,
	7, 24, 117, 118, 249, 0, 0, 58, 253, 59,
	254, 55, 54, 0, 8, 259, 252, 0, 260, 0,
	0, 0, 0, 258, 0, 0, 0, 56, 57, 0,
	0, 0, 0, 266, 0, 0, 0, 269, 0, 0,
	0, 275, 276, 0, 98, 277, 58, 0, 59, 274,
	55, 54, 0, 0, 0, 0, 279, 0, 284, 0,
	0, 0, 0, 24, 25, 26, 45, 0, 0, 288,
	35, 228, 43, 231, 230, 44, 36, 37, 0, 289,
	0, 46, 0, 0, 0, 0, 0, 0, 53, 56,
	57, 0, 0, 0, 39, 40, 41, 42, 0, 0,
	186, 187, 0, 0, 0, 0, 0, 0, 58, 0,
	59, 0, 55, 54, 24, 25, 26, 45, 0, 0,
	0, 35, 273, 43, 0, 0, 44, 36, 37, 0,
	0, 0, 46, 0, 0, 0, 0, 0, 0, 53,
	56, 57, 0, 0, 0, 39, 40, 41, 42, 0,
	0, 186, 187, 24, 25, 26, 45, 0, 0, 58,
	35, 59, 43, 55, 54, 44, 36, 37, 0, 0,
	0, 46, 0, 0, 0, 0, 0, 0, 53, 56,
	57, 0, 0, 0, 39, 40, 41, 42, 0, 0,
	186, 187, 0, 0, 0, 0, 0, 0, 58, 0,
	59, 248, 55, 54, 24, 25, 26, 45, 0, 0,
	0, 35, 245, 43, 0, 0, 44, 36, 37, 0,
	0, 0, 46, 0, 0, 0, 0, 0, 0, 53,
	56, 57, 0, 0, 0, 39, 40, 41, 42, 0,
	0, 186, 187, 0, 0, 0, 0, 0, 0, 58,
	0, 59, 0, 55, 54, 24, 25, 26, 45, 0,
	0, 0, 35, 232, 43, 0, 0, 44, 36, 37,
	0, 0, 0, 46, 0, 0, 0, 0, 0, 0,
	53, 56, 57, 0, 0, 0, 39, 40, 41, 42,
	0, 0, 186, 187, 0, 0, 0, 0, 0, 0,
	58, 0, 59, 0, 55, 54, 24, 25, 26, 45,
	0, 0, 0, 35, 227, 43, 0, 0, 44, 36,
	37, 0, 0, 0, 46, 0, 0, 0, 0, 0,
	0, 53, 56, 57, 0, 0, 0, 39, 40, 41,
	42, 0, 0, 186, 187, 0, 0, 0, 0, 0,
	0, 58, 0, 59, 0, 55, 54, 24, 25, 26,
	45, 0, 0, 0, 35, 225, 43, 0, 0, 44,
	36, 37, 0, 0, 0, 46, 0, 0, 0, 0,
	0, 0, 53, 56, 57, 0, 0, 0, 39, 40,
	41, 42, 0, 0, 186, 187, 24, 25, 26, 45,
	0, 0, 58, 35, 59, 43, 55, 54, 44, 36,
	37, 0, 0, 0, 46, 0, 0, 0, 0, 0,
	0, 53, 56, 57, 0, 0, 0, 39, 40, 41,
	42, 0, 0, 186, 187, 0, 0, 0, 0, 0,
	0, 58, 0, 59, 218, 55, 54, 24, 25, 26,
	45, 0, 0, 0, 35, 213, 43, 0, 0, 44,
	36, 37, 0, 0, 0, 46, 0, 0, 0, 0,
	0, 0, 53, 56, 57, 0, 0, 0, 39, 40,
	41, 42, 0, 0, 186, 187, 24, 25, 26, 45,
	0, 0, 58, 35, 59, 43, 55, 54, 44, 36,
	37, 0, 0, 0, 46, 0, 0, 0, 0, 0,
	0, 53, 56, 57, 0, 0, 0, 39, 40, 41,
	42, 0, 0, 186, 187, 24, 25, 26, 45, 185,
	0, 58, 35, 59, 43, 55, 54, 44, 36, 37,
	0, 0, 0, 46, 0, 0, 0, 0, 0, 0,
	53, 56, 57, 0, 0, 0, 39, 40, 41, 42,
	0, 0, 0, 184, 24, 25, 26, 45, 0, 0,
	58, 35, 59, 43, 55, 54, 44, 36, 37, 0,
	0, 0, 46, 0, 0, 0, 0, 0, 0, 53,
	56, 57, 0, 0, 0, 39, 40, 41, 42, 24,
	91, 92, 205, 0, 0, 96, 0, 0, 0, 58,
	0, 59, 0, 55, 54, 24, 91, 92, 81, 0,
	0, 96, 0, 0, 0, 56, 57, 24, 91, 92,
	45, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 56, 57, 0, 58, 0, 97, 0, 55, 54,
	0, 0, 0, 56, 57, 24, 117, 118, 0, 0,
	84, 96, 97, 0, 55, 54, 0, 0, 0, 0,
	0, 0, 58, 0, 59, 0, 55, 54, 0, 0,
	0, 56, 57, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	58, 0, 97, 0, 55, 54,
}
var RubyPact = []int{

	-32, 537, -1000, -1000, -1000, -6, -1000, -1000, -1000, 226,
	193, -1000, -1000, -1000, 25, -1000, -1000, -1000, -1000, -1000,
	-26, -1000, -1000, -1000, -1000, 137, 82, 49, 18, 0,
	-1000, -1000, -1000, -1000, -1000, 172, 160, 160, 77, 1139,
	1139, 1139, 1139, 1139, 576, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 8, 156, -1000, -1000, 1202, -1000,
	-14, -1000, -1000, -1000, 1139, 195, 576, 1202, 1139, 576,
	576, 576, 576, 1139, 194, 576, 1202, 1139, 1139, 1139,
	184, 1202, -1000, 86, 1202, 1202, 182, 116, 35, -1000,
	29, 1190, 131, -1000, -1000, -1000, -28, -28, -26, 1202,
	1139, 1139, 1139, 1139, 148, 148, 11, -1000, -1000, 1139,
	179, 81, 81, 81, 81, 81, -1000, -1000, -1000, 1100,
	1061, -1000, -1000, 151, -1000, -1000, 10, 7, 355, -1000,
	24, 334, -17, 81, 1174, -1000, 35, 29, 81, -1000,
	-1000, -1000, -1000, 81, -1000, -1000, 35, 29, 81, 81,
	81, -1000, 142, 355, 1230, 9, 35, 29, -1000, 482,
	1022, -1000, 178, -1000, 971, 140, 81, 81, 81, 81,
	-1000, 109, 174, -1000, -1, -1000, 932, 160, 881, 81,
	-1000, 638, 830, 81, -1000, -1000, -1000, -1000, 226, 81,
	75, -1000, -1000, 576, -1000, 576, -1000, -1000, -1000, 93,
	199, -20, 159, 86, -1000, 1202, -1000, -1000, -1000, -3,
	35, 29, -1000, -1000, 779, 19, -1000, 728, -1000, -1000,
	-11, 174, 127, 1139, -11, -1000, -1000, -1000, -1000, 78,
	1139, -1000, -1000, -1000, 139, -1000, -1000, 36, -30, -1000,
	1139, 576, -1000, 124, 1139, -1000, -1000, 169, -1000, 1061,
	-1000, -1000, 81, 1061, 689, -1000, 1139, -1000, 81, 1061,
	1061, 198, -1000, 1139, -1000, 60, 81, -1000, -1000, 81,
	-1000, 39, 37, -1000, 81, 1061, 1061, 1061, 165, 114,
	-23, -11, -1000, -1000, 1061, -1000, 1139, 576, 1061, 115,
	33, -11, -1000, -11, -1000, -11, -11,
}
var RubyPgo = []int{

	0, 372, 267, 263, 4, 259, 258, 418, 275, 257,
	256, 253, 0, 223, 127, 252, 251, 16, 248, 38,
	247, 3, 2, 246, 245, 241, 240, 239, 237, 236,
	235, 233, 229, 226, 363, 225, 6, 9, 221, 220,
	7, 217, 216, 22, 214, 211, 208, 205, 5, 1,
	199, 131,
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
	-34, 6, 7, 55, 6, 7, -45, -43, -12, -17,
	-14, -51, 44, -1, 6, -12, -12, -14, -1, -12,
	-12, -12, -12, -1, 6, -12, -12, -14, -1, -1,
	-1, 6, -43, -12, 10, -12, -12, -14, 6, 10,
	-34, -37, 46, -37, -34, -43, -1, -1, -1, -1,
	-40, -49, 8, -22, 6, -40, -34, 33, -34, -1,
	6, -34, -34, -1, 43, 9, 42, 43, -12, -1,
	-42, 6, 7, 10, 51, 10, 51, 42, -41, -46,
	-12, 6, 44, -48, -36, 8, 9, -12, -4, 51,
	-12, -14, -4, 13, -34, -44, 6, -34, 53, 9,
	-51, 10, -49, 35, -51, 13, -21, 13, 13, -38,
	16, 15, 13, 13, 24, -12, -12, -51, -51, 10,
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
	163, 106, 107, 0, 104, 105, 0, 0, 65, 69,
	70, 131, 0, 148, 53, 59, 116, 118, 120, 121,
	122, 123, 124, 150, 54, 60, 117, 119, 149, 151,
	152, 57, 0, 68, 0, 65, 97, 98, 110, 0,
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
