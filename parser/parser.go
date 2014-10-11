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
const DOUBLECOLON = 57387
const DOT = 57388
const PIPE = 57389
const SLASH = 57390
const AMPERSAND = 57391
const QUESTIONMARK = 57392
const CARET = 57393
const LBRACKET = 57394
const RBRACKET = 57395
const LBRACE = 57396
const RBRACE = 57397
const DOLLARSIGN = 57398
const ATSIGN = 57399
const FILE_CONST_REF = 57400
const EOF = 57401

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

//line parser.y:868

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 27,
	45, 104,
	-2, 20,
	-1, 87,
	9, 74,
	10, 74,
	-2, 166,
	-1, 100,
	45, 104,
	-2, 20,
	-1, 105,
	5, 11,
	6, 11,
	7, 11,
	10, 11,
	31, 11,
	32, 11,
	42, 11,
	52, 11,
	54, 11,
	55, 11,
	56, 11,
	57, 11,
	-2, 13,
	-1, 116,
	45, 104,
	-2, 102,
	-1, 127,
	45, 104,
	-2, 20,
	-1, 218,
	45, 105,
	-2, 103,
	-1, 221,
	9, 74,
	10, 74,
	-2, 166,
}

const RubyNprod = 177
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1350

var RubyAct = []int{

	9, 187, 22, 97, 98, 185, 91, 90, 136, 35,
	65, 2, 3, 213, 184, 176, 175, 266, 282, 86,
	141, 306, 69, 259, 222, 213, 283, 96, 4, 263,
	302, 193, 301, 130, 131, 239, 211, 112, 63, 62,
	111, 110, 209, 188, 65, 28, 125, 115, 117, 313,
	213, 67, 68, 299, 265, 64, 70, 71, 72, 213,
	137, 213, 138, 281, 66, 75, 73, 74, 144, 145,
	146, 147, 101, 150, 151, 152, 153, 189, 155, 212,
	157, 158, 159, 160, 132, 210, 190, 119, 167, 213,
	138, 101, 251, 169, 171, 138, 166, 10, 170, 274,
	96, 276, 275, 252, 85, 101, 188, 168, 186, 167,
	108, 138, 118, 101, 101, 101, 101, 179, 101, 101,
	101, 101, 177, 101, 106, 101, 101, 101, 101, 191,
	204, 81, 311, 101, 69, 257, 65, 109, 101, 101,
	189, 216, 237, 305, 96, 101, 287, 211, 173, 190,
	219, 220, 81, 65, 101, 270, 237, 139, 235, 211,
	79, 80, 108, 67, 68, 82, 223, 211, 148, 224,
	207, 208, 225, 78, 171, 204, 66, 228, 227, 204,
	161, 79, 80, 133, 134, 139, 101, 280, 116, 101,
	139, 106, 238, 204, 78, 204, 218, 106, 204, 204,
	25, 126, 127, 244, 304, 114, 139, 113, 289, 241,
	253, 240, 254, 232, 101, 196, 172, 165, 156, 101,
	143, 140, 167, 262, 138, 297, 57, 58, 258, 61,
	260, 204, 95, 215, 204, 25, 99, 100, 46, 269,
	135, 106, 231, 206, 214, 1, 247, 59, 128, 60,
	53, 56, 55, 52, 29, 101, 51, 101, 50, 49,
	286, 57, 58, 48, 25, 126, 127, 101, 101, 204,
	104, 106, 18, 204, 204, 17, 16, 15, 39, 204,
	204, 102, 59, 13, 60, 12, 56, 55, 23, 11,
	57, 58, 21, 14, 19, 204, 204, 204, 24, 30,
	102, 33, 32, 34, 204, 101, 31, 309, 204, 0,
	0, 59, 0, 105, 102, 56, 55, 0, 0, 139,
	261, 0, 102, 102, 102, 102, 103, 102, 102, 102,
	102, 0, 102, 0, 102, 102, 102, 102, 0, 0,
	0, 0, 102, 0, 0, 103, 0, 102, 102, 129,
	0, 0, 101, 0, 102, 0, 0, 0, 0, 103,
	0, 0, 0, 102, 0, 0, 0, 103, 103, 103,
	103, 0, 103, 103, 103, 103, 205, 103, 5, 103,
	103, 103, 103, 25, 99, 100, 46, 103, 0, 104,
	0, 0, 103, 103, 0, 102, 0, 0, 102, 103,
	0, 0, 0, 0, 0, 0, 236, 0, 103, 57,
	58, 0, 0, 242, 0, 0, 81, 120, 121, 122,
	123, 124, 0, 102, 0, 0, 83, 0, 102, 84,
	59, 0, 105, 0, 56, 55, 255, 256, 0, 0,
	103, 0, 142, 103, 0, 79, 80, 0, 149, 0,
	82, 0, 0, 154, 174, 178, 0, 20, 78, 0,
	162, 163, 164, 0, 102, 192, 102, 194, 103, 0,
	0, 0, 0, 103, 197, 198, 102, 102, 0, 284,
	0, 0, 0, 0, 107, 0, 180, 181, 182, 183,
	290, 0, 0, 0, 291, 195, 0, 25, 99, 100,
	221, 0, 0, 104, 0, 199, 0, 300, 0, 103,
	0, 103, 0, 0, 102, 0, 0, 107, 0, 0,
	0, 103, 103, 57, 58, 230, 0, 233, 107, 0,
	310, 312, 0, 314, 0, 315, 25, 99, 100, 87,
	107, 0, 104, 0, 59, 107, 105, 0, 56, 55,
	107, 107, 0, 0, 0, 0, 0, 107, 0, 103,
	0, 102, 57, 58, 0, 0, 107, 89, 0, 0,
	25, 217, 127, 0, 0, 0, 0, 0, 0, 0,
	0, 88, 0, 92, 0, 105, 268, 56, 55, 0,
	0, 0, 272, 0, 273, 0, 57, 58, 0, 278,
	0, 107, 279, 0, 0, 0, 103, 213, 0, 0,
	0, 0, 0, 0, 0, 0, 271, 59, 0, 60,
	0, 56, 55, 0, 0, 277, 294, 295, 0, 0,
	296, 107, 0, 0, 0, 285, 0, 0, 0, 0,
	288, 0, 0, 303, 0, 0, 0, 0, 0, 0,
	0, 0, 293, 0, 307, 0, 0, 0, 0, 298,
	25, 26, 27, 46, 0, 0, 0, 36, 246, 44,
	249, 248, 45, 37, 38, 0, 0, 0, 47, 107,
	107, 0, 308, 0, 0, 54, 57, 58, 0, 0,
	0, 40, 41, 42, 43, 0, 0, 202, 203, 0,
	0, 25, 26, 27, 46, 0, 0, 59, 36, 60,
	44, 56, 55, 45, 37, 38, 0, 0, 0, 47,
	0, 0, 0, 0, 0, 0, 54, 57, 58, 0,
	0, 0, 40, 41, 42, 43, 0, 0, 6, 7,
	0, 0, 0, 0, 0, 0, 0, 0, 59, 0,
	60, 0, 56, 55, 0, 8, 25, 26, 27, 46,
	0, 0, 0, 36, 292, 44, 0, 0, 45, 37,
	38, 0, 0, 0, 47, 0, 0, 0, 0, 0,
	0, 54, 57, 58, 0, 0, 0, 40, 41, 42,
	43, 0, 0, 202, 203, 0, 0, 25, 26, 27,
	46, 0, 0, 59, 36, 60, 44, 56, 55, 45,
	37, 38, 0, 0, 0, 47, 0, 0, 0, 0,
	0, 0, 54, 57, 58, 0, 0, 0, 40, 41,
	42, 43, 0, 0, 202, 203, 0, 0, 0, 0,
	0, 0, 0, 0, 59, 0, 60, 267, 56, 55,
	25, 26, 27, 46, 0, 0, 0, 36, 264, 44,
	0, 0, 45, 37, 38, 0, 0, 0, 47, 0,
	0, 0, 0, 0, 0, 54, 57, 58, 0, 0,
	0, 40, 41, 42, 43, 0, 0, 202, 203, 0,
	0, 0, 0, 0, 0, 0, 0, 59, 0, 60,
	0, 56, 55, 25, 26, 27, 46, 0, 0, 0,
	36, 250, 44, 0, 0, 45, 37, 38, 0, 0,
	0, 47, 0, 0, 0, 0, 0, 0, 54, 57,
	58, 0, 0, 0, 40, 41, 42, 43, 0, 0,
	202, 203, 0, 0, 0, 0, 0, 0, 0, 0,
	59, 0, 60, 0, 56, 55, 25, 26, 27, 46,
	0, 0, 0, 36, 245, 44, 0, 0, 45, 37,
	38, 0, 0, 0, 47, 0, 0, 0, 0, 0,
	0, 54, 57, 58, 0, 0, 0, 40, 41, 42,
	43, 0, 0, 202, 203, 0, 0, 0, 0, 0,
	0, 0, 0, 59, 0, 60, 0, 56, 55, 25,
	26, 27, 46, 0, 0, 0, 36, 243, 44, 0,
	0, 45, 37, 38, 0, 0, 0, 47, 0, 0,
	0, 0, 0, 0, 54, 57, 58, 0, 0, 0,
	40, 41, 42, 43, 0, 0, 202, 203, 0, 0,
	25, 26, 27, 46, 0, 0, 59, 36, 60, 44,
	56, 55, 45, 37, 38, 0, 0, 0, 47, 0,
	0, 0, 0, 0, 0, 54, 57, 58, 0, 0,
	0, 40, 41, 42, 43, 0, 0, 202, 203, 0,
	0, 0, 0, 0, 0, 0, 0, 59, 0, 60,
	234, 56, 55, 25, 26, 27, 46, 0, 0, 0,
	36, 229, 44, 0, 0, 45, 37, 38, 0, 0,
	0, 47, 0, 0, 0, 0, 0, 0, 54, 57,
	58, 0, 0, 0, 40, 41, 42, 43, 0, 0,
	202, 203, 0, 0, 25, 26, 27, 46, 0, 0,
	59, 36, 60, 44, 56, 55, 45, 37, 38, 0,
	0, 0, 47, 0, 0, 0, 0, 0, 0, 54,
	57, 58, 0, 0, 0, 40, 41, 42, 43, 0,
	0, 202, 203, 0, 0, 25, 26, 27, 46, 201,
	0, 59, 36, 60, 44, 56, 55, 45, 37, 38,
	0, 0, 0, 47, 0, 0, 0, 0, 0, 0,
	54, 57, 58, 0, 0, 0, 40, 41, 42, 43,
	0, 0, 0, 200, 0, 0, 25, 26, 27, 46,
	0, 0, 59, 36, 60, 44, 56, 55, 45, 37,
	38, 0, 0, 0, 47, 0, 0, 0, 0, 0,
	0, 54, 57, 58, 0, 0, 0, 40, 41, 42,
	43, 25, 99, 100, 87, 0, 94, 104, 0, 0,
	0, 0, 0, 59, 0, 60, 0, 56, 55, 0,
	0, 0, 0, 0, 0, 69, 0, 57, 58, 0,
	0, 93, 89, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 0, 0, 88, 0, 92, 0,
	105, 0, 56, 55, 67, 68, 76, 0, 0, 70,
	71, 72, 0, 0, 0, 0, 0, 66, 75, 73,
	74, 0, 67, 68, 226, 0, 0, 70, 71, 72,
	0, 0, 0, 0, 0, 66, 75, 73, 74, 77,
}
var RubyPact = []int{

	-31, 696, -1000, -1000, -1000, -4, -1000, -1000, -1000, 1299,
	412, -1000, -1000, -1000, 87, -1000, -1000, -1000, -1000, -1000,
	-27, -1000, -1000, -1000, -1000, -1000, 1256, 102, 6, 5,
	2, -1000, -1000, -1000, -1000, -1000, 201, 181, 181, 77,
	1221, 1221, 1221, 1221, 1221, 195, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 27, 177, -1000, -1000, 230,
	-1000, -25, -1000, -1000, -1000, 1221, 214, 195, 195, 195,
	230, 1221, 195, 195, 195, 195, 1221, 195, 212, 195,
	195, 195, 230, 1221, 1221, 1221, 211, 230, -1000, -1000,
	-1000, 97, 230, 230, 210, 138, 130, -1000, -1000, 531,
	154, -1000, -1000, -1000, -32, -32, 148, -27, 230, 1221,
	1221, 1221, 1221, 100, 100, -2, -1000, -1000, 1221, 209,
	30, 30, 30, 30, 30, -1000, -1000, -1000, 1180, 1139,
	-1000, -1000, 164, -1000, -1000, 32, 26, 18, -1000, 127,
	565, 189, 30, 492, -1000, -1000, -1000, 130, 148, 30,
	-1000, -1000, -1000, -1000, 30, -20, -1000, -1000, -1000, -1000,
	130, 148, 30, 30, 30, -1000, 157, 18, 259, 1281,
	-1000, 130, -1000, 378, 1098, -1000, 207, -1000, 1045, 149,
	30, 30, 30, 30, -1000, 132, 37, -1000, 0, 205,
	203, -1000, 1004, 181, 951, 30, -1000, 655, 898, 30,
	-1000, -1000, -1000, -1000, 1299, 30, 79, -1000, -1000, 195,
	-1000, 195, -1000, -1000, -1000, 125, 224, -21, -1000, 97,
	-1000, 230, 230, -1000, -1000, -1000, -6, -1000, -1000, -1000,
	845, 7, -1000, 792, -1000, -1000, -17, 37, 146, 1221,
	-1000, -1000, -17, -1000, -1000, -1000, -1000, 86, 1221, -1000,
	-1000, -1000, 180, -1000, -1000, 8, -29, -1000, 1221, 195,
	137, 148, 130, 1221, -1000, -1000, 202, -1000, 1139, -1000,
	-1000, 30, 1139, 751, -1000, 1221, -1000, 30, 1139, 1139,
	221, -1000, 1221, -1000, 47, 30, -1000, -1000, 30, -1000,
	19, 17, -1000, 30, 1139, 1139, 1139, 198, 139, -23,
	-17, -1000, -1000, 1139, -1000, 1221, 195, 1139, 122, 39,
	-17, -1000, -17, -1000, -17, -17,
}
var RubyPgo = []int{

	0, 376, 306, 303, 3, 302, 301, 457, 299, 298,
	294, 293, 292, 0, 254, 97, 289, 288, 2, 285,
	1, 45, 283, 9, 4, 278, 277, 276, 275, 272,
	263, 259, 258, 256, 253, 250, 349, 248, 7, 16,
	246, 245, 14, 244, 243, 8, 242, 240, 233, 232,
	6, 5, 229, 221,
}
var RubyR1 = []int{

	0, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 53, 53, 36, 36, 36, 36, 36, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 18,
	18, 18, 18, 18, 18, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	38, 38, 47, 47, 45, 45, 45, 45, 45, 50,
	50, 50, 50, 24, 24, 49, 49, 49, 16, 16,
	42, 42, 20, 20, 20, 20, 51, 51, 51, 19,
	19, 22, 23, 23, 52, 52, 11, 11, 11, 11,
	11, 11, 8, 8, 21, 21, 14, 14, 25, 25,
	26, 27, 28, 29, 30, 30, 30, 30, 31, 32,
	33, 34, 35, 2, 5, 6, 6, 3, 3, 43,
	43, 43, 43, 48, 48, 48, 4, 4, 4, 4,
	39, 46, 46, 46, 10, 10, 10, 10, 10, 10,
	10, 10, 40, 40, 40, 40, 37, 37, 37, 7,
	12, 44, 44, 44, 44, 17, 9,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 4, 2, 2, 4, 2, 3, 3, 4, 4,
	3, 2, 3, 3, 3, 3, 3, 3, 4, 6,
	3, 1, 1, 3, 0, 1, 1, 1, 3, 1,
	1, 3, 3, 1, 1, 1, 3, 3, 7, 7,
	1, 3, 1, 2, 3, 2, 0, 1, 3, 4,
	6, 4, 1, 3, 1, 3, 3, 3, 3, 3,
	3, 3, 2, 2, 2, 2, 3, 3, 3, 3,
	2, 2, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 1, 1, 3, 3, 5, 5, 0,
	4, 7, 8, 3, 7, 8, 3, 4, 4, 3,
	3, 0, 1, 3, 4, 5, 3, 3, 3, 3,
	3, 4, 4, 3, 3, 2, 0, 2, 2, 3,
	4, 0, 3, 4, 6, 1, 5,
}
var RubyChk = []int{

	-1000, -41, 42, 43, 59, -1, 42, 43, 59, -13,
	-15, -16, -19, -22, -11, -26, -27, -28, -29, -10,
	-7, -12, -18, -17, -9, 5, 6, 7, -21, -14,
	-8, -2, -5, -6, -3, -23, 12, 18, 19, -25,
	36, 37, 38, 39, 14, 17, 8, 23, -30, -31,
	-32, -33, -34, -35, 30, 57, 56, 31, 32, 52,
	54, -52, 43, 42, 59, 14, 46, 33, 34, 4,
	38, 39, 40, 48, 49, 47, 17, 50, 46, 33,
	34, 4, 38, 14, 17, 17, 46, 8, 50, 36,
	-38, -50, 52, 35, 10, -49, -13, -4, -24, 6,
	7, -21, -14, -8, 11, 54, -15, -7, 8, 35,
	35, 35, 35, 6, 4, -23, 7, -23, 35, 10,
	-1, -1, -1, -1, -1, -13, 6, 7, -37, -36,
	6, 7, 57, 6, 7, -47, -45, -13, -18, -15,
	-53, 45, -1, 6, -13, -13, -13, -13, -15, -1,
	-13, -13, -13, -13, -1, -13, 6, -13, -13, -13,
	-13, -15, -1, -1, -1, 6, -45, -13, 10, -13,
	-24, -13, 6, 10, -36, -39, 47, -39, -36, -45,
	-1, -1, -1, -1, -42, -51, 8, -20, 6, 40,
	49, -42, -36, 33, -36, -1, 6, -36, -36, -1,
	43, 9, 42, 43, -13, -1, -44, 6, 7, 10,
	53, 10, 53, 42, -43, -48, -13, 6, 7, -50,
	-38, 8, 44, 9, -13, -4, 53, -24, -4, 13,
	-36, -46, 6, -36, 55, 9, -53, 10, -51, 35,
	6, 6, -53, 13, -23, 13, 13, -40, 16, 15,
	13, 13, 24, -13, -13, -53, -53, 10, 4, 44,
	-45, -15, -13, 35, 13, 47, 10, 55, -36, -20,
	9, -1, -36, -36, 13, 16, 15, -1, -36, -36,
	7, 55, 10, 55, -53, -1, -13, 9, -1, 6,
	-53, -53, 13, -1, -36, -36, -36, 4, -1, 6,
	-53, 13, 13, -36, 6, 4, 44, -36, -1, -13,
	-53, 10, -53, 10, -53, -53,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
	46, 47, 48, 49, 50, 18, 19, -2, 21, 22,
	23, 24, 25, 26, 27, 28, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 166, 13, 29, 30,
	31, 32, 33, 34, 175, 0, 0, 133, 134, 74,
	11, 0, 5, 6, 7, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, -2, 52, 53,
	55, 61, 74, 0, 0, 71, 79, 80, 85, 19,
	-2, 21, 22, 23, 13, -2, 84, 0, 74, 0,
	0, 0, 0, 96, 96, 13, -2, 13, 0, 0,
	120, 121, 122, 123, 13, 13, 19, -2, 0, 171,
	114, 115, 0, 112, 113, 0, 0, 72, 76, 77,
	139, 0, 156, 56, 62, 64, 66, 124, 126, 128,
	129, 130, 131, 132, 158, 0, 57, 63, 65, 67,
	125, 127, 157, 159, 160, 60, 0, 75, 0, 72,
	106, 83, 118, 0, 0, 13, 151, 13, 0, 0,
	107, 108, 109, 110, 11, 90, 96, 97, 92, 0,
	0, 11, 0, 0, 0, 111, 119, 0, 0, 167,
	168, 169, 14, 15, 16, 17, 0, 116, 117, 0,
	135, 0, 136, 12, 11, 11, 0, 19, -2, 58,
	59, -2, 0, 51, 81, 82, 68, 86, 87, 146,
	0, 0, 152, 0, 149, 54, 13, 0, 0, 0,
	93, 95, 13, 99, 13, 101, 154, 0, 0, 13,
	161, 170, 13, 73, 78, 0, 0, 11, 0, 0,
	0, 176, 0, 0, 147, 150, 0, 148, 11, 98,
	91, 94, 11, 0, 155, 0, 13, 13, 165, 172,
	13, 137, 0, 138, 0, 11, 143, 70, 69, 153,
	0, 0, 100, 13, 163, 164, 173, 0, 0, 0,
	140, 88, 89, 162, 13, 0, 0, 174, 11, 11,
	141, 11, 144, 11, 142, 145,
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
	52, 53, 54, 55, 56, 57, 58, 59,
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
			name := RubyS[Rubypt-1].genericValue.(ast.BareReference).Name
			RubyVAL.genericValue = ast.CallExpression{Func: ast.BareReference{Name: name + "?"}}
		}
	case 53:
		//line parser.y:218
		{
			name := RubyS[Rubypt-1].genericValue.(ast.BareReference).Name
			RubyVAL.genericValue = ast.CallExpression{Func: ast.BareReference{Name: name + "!"}}
		}
	case 54:
		//line parser.y:223
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 55:
		//line parser.y:230
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 56:
		//line parser.y:237
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 57:
		//line parser.y:244
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 58:
		//line parser.y:251
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 59:
		//line parser.y:259
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 60:
		//line parser.y:267
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 61:
		//line parser.y:277
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 62:
		//line parser.y:284
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 63:
		//line parser.y:292
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 64:
		//line parser.y:300
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 65:
		//line parser.y:308
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 66:
		//line parser.y:316
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 67:
		//line parser.y:324
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 68:
		//line parser.y:334
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 69:
		//line parser.y:344
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 70:
		//line parser.y:353
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 71:
		//line parser.y:355
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 72:
		//line parser.y:358
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:360
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:362
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 75:
		//line parser.y:364
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:366
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:368
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 78:
		//line parser.y:370
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 79:
		//line parser.y:374
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 80:
		//line parser.y:376
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 81:
		//line parser.y:378
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 82:
		//line parser.y:380
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 83:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 84:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 85:
		//line parser.y:385
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 86:
		//line parser.y:387
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 87:
		//line parser.y:389
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 88:
		//line parser.y:395
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 89:
		//line parser.y:403
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 90:
		//line parser.y:412
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 91:
		//line parser.y:414
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 92:
		//line parser.y:417
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 93:
		//line parser.y:419
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 94:
		//line parser.y:421
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 95:
		//line parser.y:423
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 96:
		//line parser.y:425
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 97:
		//line parser.y:427
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 98:
		//line parser.y:431
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 99:
		//line parser.y:436
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 100:
		//line parser.y:443
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 101:
		//line parser.y:453
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 102:
		//line parser.y:462
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 103:
		//line parser.y:468
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 104:
		//line parser.y:476
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 105:
		//line parser.y:480
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 106:
		//line parser.y:485
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 107:
		//line parser.y:492
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 108:
		//line parser.y:499
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 109:
		//line parser.y:506
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 110:
		//line parser.y:513
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 111:
		//line parser.y:520
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 112:
		//line parser.y:528
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 113:
		//line parser.y:530
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 114:
		//line parser.y:533
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 115:
		//line parser.y:535
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 116:
		//line parser.y:538
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 117:
		//line parser.y:540
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 118:
		//line parser.y:543
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 119:
		//line parser.y:545
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 120:
		//line parser.y:547
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 121:
		//line parser.y:548
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 122:
		//line parser.y:549
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 123:
		//line parser.y:550
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 124:
		//line parser.y:553
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 125:
		//line parser.y:561
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 126:
		//line parser.y:569
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 127:
		//line parser.y:577
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 128:
		//line parser.y:586
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 129:
		//line parser.y:595
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 130:
		//line parser.y:604
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 131:
		//line parser.y:613
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 132:
		//line parser.y:622
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 133:
		//line parser.y:630
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 134:
		//line parser.y:631
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 135:
		//line parser.y:633
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 136:
		//line parser.y:635
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 137:
		//line parser.y:638
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 138:
		//line parser.y:646
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 139:
		//line parser.y:654
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 140:
		//line parser.y:656
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 141:
		//line parser.y:663
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 142:
		//line parser.y:670
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 143:
		//line parser.y:678
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 144:
		//line parser.y:685
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 145:
		//line parser.y:692
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 146:
		//line parser.y:700
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 147:
		//line parser.y:702
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 148:
		//line parser.y:709
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 149:
		//line parser.y:716
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 150:
		//line parser.y:719
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 151:
		//line parser.y:721
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 152:
		//line parser.y:723
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 153:
		//line parser.y:725
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 154:
		//line parser.y:728
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 155:
		//line parser.y:735
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 156:
		//line parser.y:743
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 157:
		//line parser.y:750
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 158:
		//line parser.y:757
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 159:
		//line parser.y:764
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 160:
		//line parser.y:771
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 161:
		//line parser.y:778
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-2].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 162:
		//line parser.y:786
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 163:
		//line parser.y:793
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 164:
		//line parser.y:800
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 165:
		//line parser.y:807
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 166:
		//line parser.y:814
		{
		}
	case 167:
		//line parser.y:815
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 168:
		//line parser.y:816
		{
		}
	case 169:
		//line parser.y:819
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 170:
		//line parser.y:822
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 171:
		//line parser.y:830
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 172:
		//line parser.y:832
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 173:
		//line parser.y:834
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 174:
		//line parser.y:843
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
	case 175:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 176:
		//line parser.y:860
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
