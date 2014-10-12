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

//line parser.y:841

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 24,
	46, 100,
	-2, 20,
	-1, 80,
	10, 73,
	11, 73,
	-2, 160,
	-1, 90,
	46, 100,
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
	46, 100,
	-2, 98,
	-1, 204,
	46, 101,
	-2, 99,
	-1, 211,
	10, 73,
	11, 73,
	-2, 160,
}

const RubyNprod = 171
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1405

var RubyAct = []int{

	9, 175, 20, 32, 173, 172, 88, 82, 128, 81,
	157, 197, 22, 23, 61, 24, 62, 131, 2, 3,
	36, 199, 44, 275, 87, 45, 37, 38, 66, 158,
	122, 46, 123, 107, 276, 4, 132, 195, 53, 56,
	57, 110, 112, 176, 40, 41, 42, 43, 300, 199,
	6, 7, 254, 229, 198, 199, 64, 63, 101, 129,
	58, 130, 59, 296, 55, 54, 70, 8, 274, 137,
	138, 139, 140, 65, 142, 143, 144, 145, 177, 147,
	196, 150, 124, 130, 152, 153, 130, 178, 70, 149,
	87, 257, 199, 100, 99, 307, 68, 69, 150, 181,
	130, 71, 72, 73, 167, 168, 161, 159, 79, 67,
	76, 74, 75, 78, 114, 179, 216, 66, 68, 69,
	103, 104, 190, 22, 89, 61, 90, 80, 256, 85,
	95, 67, 202, 102, 150, 78, 130, 87, 66, 113,
	105, 295, 205, 106, 209, 176, 210, 293, 174, 305,
	56, 57, 214, 66, 84, 97, 217, 190, 215, 103,
	104, 190, 218, 242, 70, 151, 186, 187, 281, 197,
	199, 83, 102, 96, 243, 55, 54, 299, 77, 228,
	177, 190, 98, 190, 199, 234, 190, 248, 66, 178,
	226, 261, 227, 227, 68, 69, 244, 232, 245, 71,
	72, 73, 70, 265, 87, 267, 266, 67, 76, 74,
	75, 78, 150, 253, 130, 251, 197, 155, 246, 247,
	252, 190, 225, 197, 190, 70, 213, 197, 193, 260,
	194, 273, 68, 69, 25, 22, 89, 61, 90, 62,
	97, 190, 190, 212, 133, 67, 125, 111, 126, 78,
	278, 280, 204, 298, 283, 68, 69, 109, 91, 108,
	190, 231, 56, 57, 190, 190, 277, 230, 67, 249,
	190, 34, 78, 190, 222, 184, 171, 284, 166, 154,
	136, 285, 291, 58, 60, 59, 86, 55, 54, 190,
	190, 190, 201, 91, 127, 94, 221, 294, 190, 68,
	69, 303, 190, 91, 91, 91, 91, 192, 91, 91,
	91, 91, 67, 91, 200, 91, 78, 1, 91, 91,
	304, 306, 237, 308, 91, 309, 134, 52, 51, 50,
	94, 49, 91, 48, 47, 17, 16, 15, 91, 91,
	94, 94, 94, 94, 14, 94, 94, 94, 94, 39,
	94, 12, 94, 11, 21, 94, 94, 10, 19, 13,
	18, 94, 33, 35, 30, 29, 91, 31, 91, 94,
	28, 91, 0, 0, 0, 94, 94, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 91, 0, 0, 26,
	91, 22, 89, 61, 90, 62, 0, 0, 95, 0,
	0, 0, 0, 94, 0, 94, 0, 0, 94, 0,
	0, 0, 0, 92, 0, 0, 0, 0, 56, 57,
	0, 0, 0, 94, 0, 0, 0, 94, 0, 0,
	91, 121, 91, 0, 0, 0, 0, 0, 91, 58,
	0, 96, 0, 55, 54, 0, 91, 91, 92, 22,
	89, 61, 90, 80, 0, 0, 95, 0, 92, 92,
	92, 92, 0, 92, 92, 92, 92, 94, 92, 94,
	92, 0, 0, 92, 92, 94, 56, 57, 0, 92,
	0, 0, 0, 94, 94, 91, 0, 92, 0, 250,
	27, 0, 0, 92, 92, 0, 0, 83, 0, 96,
	0, 55, 54, 0, 0, 22, 23, 61, 24, 62,
	0, 0, 0, 36, 93, 44, 0, 0, 45, 37,
	38, 92, 94, 92, 46, 0, 92, 156, 160, 0,
	0, 53, 56, 57, 0, 91, 0, 40, 41, 42,
	43, 92, 180, 0, 182, 92, 0, 0, 0, 93,
	0, 185, 0, 58, 0, 59, 0, 55, 54, 93,
	93, 93, 93, 0, 93, 93, 93, 93, 70, 93,
	0, 93, 94, 0, 93, 93, 0, 0, 0, 0,
	93, 0, 0, 0, 0, 92, 0, 92, 93, 220,
	0, 223, 0, 92, 93, 93, 0, 0, 68, 69,
	0, 92, 92, 71, 72, 73, 0, 191, 0, 5,
	0, 67, 76, 74, 75, 78, 0, 0, 240, 241,
	0, 0, 93, 0, 93, 0, 0, 93, 0, 22,
	89, 61, 90, 211, 0, 0, 95, 0, 0, 0,
	92, 0, 93, 0, 0, 0, 93, 0, 115, 116,
	117, 118, 119, 120, 0, 0, 56, 57, 259, 0,
	0, 0, 0, 0, 263, 0, 264, 0, 0, 0,
	0, 269, 0, 0, 135, 272, 0, 58, 0, 96,
	141, 55, 54, 0, 0, 146, 93, 148, 93, 0,
	92, 0, 0, 0, 93, 0, 0, 0, 0, 288,
	289, 0, 93, 93, 0, 290, 162, 163, 164, 165,
	0, 0, 0, 169, 170, 0, 0, 0, 0, 297,
	0, 183, 0, 0, 22, 23, 61, 24, 62, 0,
	301, 0, 36, 236, 44, 239, 238, 45, 37, 38,
	0, 93, 206, 46, 0, 0, 0, 0, 0, 0,
	53, 56, 57, 0, 0, 0, 40, 41, 42, 43,
	0, 0, 188, 189, 0, 0, 0, 0, 0, 0,
	0, 0, 58, 0, 59, 0, 55, 54, 0, 22,
	23, 61, 24, 62, 0, 0, 0, 36, 286, 44,
	0, 93, 45, 37, 38, 0, 0, 0, 46, 0,
	0, 0, 0, 0, 0, 53, 56, 57, 0, 0,
	0, 40, 41, 42, 43, 0, 0, 188, 189, 0,
	0, 0, 0, 0, 0, 0, 0, 58, 0, 59,
	0, 55, 54, 0, 0, 0, 0, 262, 0, 0,
	0, 0, 0, 0, 0, 0, 268, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 279, 0, 0,
	0, 0, 282, 22, 23, 61, 24, 62, 0, 0,
	0, 36, 271, 44, 287, 0, 45, 37, 38, 0,
	0, 0, 46, 292, 0, 0, 0, 0, 0, 53,
	56, 57, 0, 0, 0, 40, 41, 42, 43, 0,
	0, 188, 189, 0, 0, 0, 0, 302, 0, 0,
	0, 58, 0, 59, 0, 55, 54, 22, 23, 61,
	24, 62, 0, 0, 0, 36, 270, 44, 0, 0,
	45, 37, 38, 0, 0, 0, 46, 0, 0, 0,
	0, 0, 0, 53, 56, 57, 0, 0, 0, 40,
	41, 42, 43, 0, 0, 188, 189, 0, 22, 23,
	61, 24, 62, 0, 0, 58, 36, 59, 44, 55,
	54, 45, 37, 38, 0, 0, 0, 46, 0, 0,
	0, 0, 0, 0, 53, 56, 57, 0, 0, 0,
	40, 41, 42, 43, 0, 0, 188, 189, 0, 0,
	0, 0, 0, 0, 0, 0, 58, 0, 59, 258,
	55, 54, 22, 23, 61, 24, 62, 0, 0, 0,
	36, 255, 44, 0, 0, 45, 37, 38, 0, 0,
	0, 46, 0, 0, 0, 0, 0, 0, 53, 56,
	57, 0, 0, 0, 40, 41, 42, 43, 0, 0,
	188, 189, 0, 0, 0, 0, 0, 0, 0, 0,
	58, 0, 59, 0, 55, 54, 22, 23, 61, 24,
	62, 0, 0, 0, 36, 235, 44, 0, 0, 45,
	37, 38, 0, 0, 0, 46, 0, 0, 0, 0,
	0, 0, 53, 56, 57, 0, 0, 0, 40, 41,
	42, 43, 0, 0, 188, 189, 0, 0, 0, 0,
	0, 0, 0, 0, 58, 0, 59, 0, 55, 54,
	22, 23, 61, 24, 62, 0, 0, 0, 36, 233,
	44, 0, 0, 45, 37, 38, 0, 0, 0, 46,
	0, 0, 0, 0, 0, 0, 53, 56, 57, 0,
	0, 0, 40, 41, 42, 43, 0, 0, 188, 189,
	0, 22, 23, 61, 24, 62, 0, 0, 58, 36,
	59, 44, 55, 54, 45, 37, 38, 0, 0, 0,
	46, 0, 0, 0, 0, 0, 0, 53, 56, 57,
	0, 0, 0, 40, 41, 42, 43, 0, 0, 188,
	189, 0, 0, 0, 0, 0, 0, 0, 0, 58,
	0, 59, 224, 55, 54, 22, 23, 61, 24, 62,
	0, 0, 0, 36, 219, 44, 0, 0, 45, 37,
	38, 0, 0, 0, 46, 0, 0, 0, 0, 0,
	0, 53, 56, 57, 0, 0, 0, 40, 41, 42,
	43, 0, 0, 188, 189, 0, 22, 23, 61, 24,
	62, 0, 0, 58, 36, 59, 44, 55, 54, 45,
	37, 38, 0, 0, 0, 46, 0, 0, 0, 0,
	0, 0, 53, 56, 57, 0, 0, 0, 40, 41,
	42, 43, 0, 0, 188, 189, 0, 22, 23, 61,
	24, 62, 208, 0, 58, 36, 59, 44, 55, 54,
	45, 37, 38, 0, 0, 0, 46, 0, 0, 0,
	0, 0, 0, 53, 56, 57, 0, 0, 0, 40,
	41, 42, 43, 0, 0, 0, 207, 22, 203, 61,
	90, 62, 0, 0, 0, 58, 0, 59, 0, 55,
	54, 22, 89, 61, 90, 80, 0, 0, 95, 0,
	0, 0, 0, 0, 56, 57, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 199, 0, 0, 56, 57,
	0, 0, 0, 0, 0, 58, 0, 59, 0, 55,
	54, 0, 0, 0, 0, 0, 0, 0, 0, 83,
	0, 96, 0, 55, 54,
}
var RubyPact = []int{

	-25, 7, -1000, -1000, -1000, 13, -1000, -1000, -1000, 160,
	-1000, -1000, -1000, 90, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 118, 146, 58, 57, 22, -1000, -1000,
	-1000, -1000, -1000, -1000, 125, -14, 253, 239, 239, 103,
	500, 500, 500, 500, 500, 500, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 24, 240, -1000, -1000, 230, -1000,
	-10, 235, -1000, -1000, -1000, -1000, 500, 274, 230, 230,
	230, 230, 500, 230, 230, 230, 230, 500, 230, 500,
	230, -1000, 154, 230, 230, 273, 206, 221, -1000, 1346,
	231, -1000, -1000, -1000, 86, -19, -19, 230, 500, 500,
	500, 500, 272, 230, 230, 500, 500, 270, 139, 139,
	65, -1000, -1000, 500, 269, 102, 102, 102, 102, 102,
	123, 1251, -1000, -1000, 222, -1000, -1000, 26, 0, 564,
	-1000, 1332, 244, 230, 1292, 102, 624, 221, 221, 221,
	221, 102, 221, 221, 221, 221, 102, 198, 102, 216,
	564, 386, 62, 221, -1000, 386, 1210, -1000, 268, -1000,
	1156, 212, 102, 102, 102, 102, -1000, 221, 221, 102,
	102, -1000, -1000, 182, 37, -1000, 17, 261, 255, -1000,
	1115, 239, 1061, 102, -1000, 719, -1000, -1000, -1000, -1000,
	160, 102, 149, -1000, -1000, 230, -1000, 230, -1000, -1000,
	-1000, 176, 265, 444, -1000, 205, 102, -1000, -1000, 154,
	-1000, 230, 230, -1000, 221, -1000, 16, 221, -1000, -1000,
	1007, 80, -1000, 953, -1000, -1000, 6, 37, 181, 500,
	-1000, -1000, 6, -1000, -1000, -1000, -1000, 189, 500, -1000,
	912, 858, -1000, 223, 221, 221, 12, -22, -1000, 500,
	230, -1000, 158, 221, 500, -1000, -1000, 248, -1000, 1251,
	-1000, -1000, 102, 1251, 774, -1000, 500, -1000, 102, 1251,
	-1000, -1000, 1251, 278, -1000, 500, -1000, 141, 160, 102,
	221, -1000, 102, -1000, 127, 49, -1000, 102, 1251, 1251,
	1251, 247, 173, 3, 6, -1000, -1000, 1251, -1000, 500,
	230, 1251, 138, 84, 6, -1000, 6, -1000, 6, 6,
}
var RubyPgo = []int{

	0, 607, 370, 367, 6, 365, 364, 363, 490, 362,
	360, 359, 358, 0, 389, 271, 357, 354, 2, 353,
	1, 234, 351, 3, 349, 344, 337, 336, 335, 334,
	333, 331, 329, 328, 327, 431, 326, 9, 10, 322,
	317, 5, 314, 307, 8, 296, 294, 292, 286, 7,
	4, 284, 17,
}
var RubyR1 = []int{

	0, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 52, 52, 35, 35, 35, 35, 35, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 18, 18, 18, 18, 18, 18, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 37,
	37, 46, 46, 44, 44, 44, 44, 49, 49, 49,
	49, 48, 48, 48, 16, 16, 41, 41, 20, 20,
	20, 20, 50, 50, 50, 19, 19, 22, 23, 23,
	51, 51, 11, 11, 11, 11, 11, 11, 8, 8,
	21, 21, 14, 14, 24, 24, 25, 26, 27, 28,
	29, 30, 31, 32, 33, 34, 2, 5, 6, 6,
	3, 3, 42, 42, 42, 42, 47, 47, 47, 4,
	4, 4, 4, 38, 45, 45, 45, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 39, 39, 39, 39,
	36, 36, 36, 7, 12, 43, 43, 43, 43, 17,
	9,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 4, 1, 4, 4, 2, 3, 3, 4, 4,
	3, 2, 3, 3, 3, 3, 3, 4, 6, 3,
	1, 1, 3, 0, 1, 1, 3, 1, 1, 3,
	3, 1, 3, 3, 7, 7, 1, 3, 1, 2,
	3, 2, 0, 1, 3, 4, 6, 4, 1, 3,
	1, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	2, 2, 3, 3, 3, 3, 2, 2, 2, 2,
	3, 3, 3, 3, 3, 3, 1, 1, 3, 3,
	5, 5, 0, 4, 7, 8, 3, 7, 8, 3,
	4, 4, 3, 3, 0, 1, 3, 4, 5, 3,
	3, 3, 3, 3, 5, 5, 4, 3, 3, 2,
	0, 2, 2, 3, 4, 0, 3, 4, 6, 1,
	5,
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
	-1, -35, 6, 8, 58, 6, 8, -46, -44, -13,
	-18, -52, 46, 9, -36, -1, 6, -13, -13, -13,
	-13, -1, -13, -13, -13, -13, -1, -13, -1, -44,
	-13, 11, -13, -13, 6, 11, -35, -38, 48, -38,
	-35, -44, -1, -1, -1, -1, 6, -13, -13, -1,
	-1, 6, -41, -50, 9, -20, 6, 41, 50, -41,
	-35, 34, -35, -1, 6, -35, 43, 44, 43, 44,
	-13, -1, -43, 6, 8, 11, 54, 11, 54, 43,
	-42, -47, -13, 6, 8, -44, -1, 44, 10, -49,
	-37, 9, 45, 10, -13, -4, 54, -13, -4, 14,
	-35, -45, 6, -35, 56, 10, -52, 11, -50, 36,
	6, 6, -52, 14, -23, 14, 14, -39, 17, 16,
	-35, -35, 14, 25, -13, -13, -52, -52, 11, 4,
	45, 10, -44, -13, 36, 14, 48, 11, 56, -35,
	-20, 10, -1, -35, -35, 14, 17, 16, -1, -35,
	14, 14, -35, 8, 56, 11, 56, -52, -13, -1,
	-13, 10, -1, 6, -52, -52, 14, -1, -35, -35,
	-35, 4, -1, 6, -52, 14, 14, -35, 6, 4,
	45, -35, -1, -13, -52, 11, -52, 11, -52, -52,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 38,
	39, 40, 41, 42, 43, 44, 45, 46, 47, 48,
	49, 50, 18, 19, -2, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 13, 32, 33, 34,
	35, 36, 37, 169, 0, 0, 126, 127, 73, 11,
	0, 52, 160, 5, 6, 7, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	-2, 55, 61, 73, 0, 0, 70, 77, 78, 19,
	-2, 21, 22, 23, 30, 13, -2, 73, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 92, 92,
	13, -2, 13, 0, 0, 116, 117, 118, 119, 13,
	0, 165, 110, 111, 0, 108, 109, 0, 0, 71,
	75, 132, 0, 73, 0, 149, 56, 62, 64, 66,
	120, 121, 122, 123, 124, 125, 151, 0, 153, 0,
	74, 0, 71, 102, 114, 0, 0, 13, 144, 13,
	0, 0, 103, 104, 105, 106, 57, 63, 65, 150,
	152, 60, 11, 86, 92, 93, 88, 0, 0, 11,
	0, 0, 0, 107, 115, 0, 13, 13, 14, 15,
	16, 17, 0, 112, 113, 0, 128, 0, 129, 12,
	11, 11, 0, 19, -2, 0, 161, 162, 163, 58,
	59, -2, 0, 51, 79, 80, 67, 82, 83, 139,
	0, 0, 145, 0, 142, 54, 13, 0, 0, 0,
	89, 91, 13, 95, 13, 97, 147, 0, 0, 13,
	0, 0, 164, 13, 72, 76, 0, 0, 11, 0,
	0, 53, 0, 170, 0, 140, 143, 0, 141, 11,
	94, 87, 90, 11, 0, 148, 0, 13, 13, 159,
	154, 155, 166, 13, 130, 0, 131, 0, 38, 11,
	136, 69, 68, 146, 0, 0, 96, 13, 157, 158,
	167, 0, 0, 0, 133, 84, 85, 156, 13, 0,
	0, 168, 11, 11, 134, 11, 137, 11, 135, 138,
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
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 55:
		//line parser.y:231
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
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
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 58:
		//line parser.y:252
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 59:
		//line parser.y:260
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 60:
		//line parser.y:268
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 61:
		//line parser.y:278
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 62:
		//line parser.y:285
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 63:
		//line parser.y:293
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 64:
		//line parser.y:301
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 65:
		//line parser.y:309
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 66:
		//line parser.y:317
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 67:
		//line parser.y:327
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 68:
		//line parser.y:337
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 69:
		//line parser.y:346
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 70:
		//line parser.y:348
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 71:
		//line parser.y:351
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 72:
		//line parser.y:353
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:355
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 74:
		//line parser.y:357
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 75:
		//line parser.y:359
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:361
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:365
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 78:
		//line parser.y:367
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 79:
		//line parser.y:369
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 80:
		//line parser.y:371
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 81:
		//line parser.y:374
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 82:
		//line parser.y:376
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 83:
		//line parser.y:378
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 84:
		//line parser.y:384
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 85:
		//line parser.y:392
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 86:
		//line parser.y:401
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 87:
		//line parser.y:403
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 88:
		//line parser.y:406
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 89:
		//line parser.y:408
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 90:
		//line parser.y:410
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 91:
		//line parser.y:412
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 92:
		//line parser.y:414
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 93:
		//line parser.y:416
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 94:
		//line parser.y:420
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 95:
		//line parser.y:425
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 96:
		//line parser.y:432
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 97:
		//line parser.y:442
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 98:
		//line parser.y:451
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 99:
		//line parser.y:457
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 100:
		//line parser.y:465
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 101:
		//line parser.y:469
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
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
		//line parser.y:509
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 108:
		//line parser.y:517
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 109:
		//line parser.y:519
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 110:
		//line parser.y:522
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 111:
		//line parser.y:524
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 112:
		//line parser.y:527
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 113:
		//line parser.y:529
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 114:
		//line parser.y:532
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 115:
		//line parser.y:534
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 116:
		//line parser.y:536
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 117:
		//line parser.y:537
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 118:
		//line parser.y:538
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 119:
		//line parser.y:539
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 120:
		//line parser.y:542
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 121:
		//line parser.y:551
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 122:
		//line parser.y:560
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 123:
		//line parser.y:569
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 124:
		//line parser.y:578
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 125:
		//line parser.y:587
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 126:
		//line parser.y:595
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 127:
		//line parser.y:596
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 128:
		//line parser.y:598
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 129:
		//line parser.y:600
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 130:
		//line parser.y:603
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 131:
		//line parser.y:611
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 132:
		//line parser.y:619
		{
			RubyVAL.genericSlice = ast.Nodes{}
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
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 135:
		//line parser.y:635
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 136:
		//line parser.y:643
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 137:
		//line parser.y:650
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 138:
		//line parser.y:657
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 139:
		//line parser.y:665
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
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
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 142:
		//line parser.y:681
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 143:
		//line parser.y:684
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 144:
		//line parser.y:686
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 145:
		//line parser.y:688
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 146:
		//line parser.y:690
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 147:
		//line parser.y:693
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 148:
		//line parser.y:700
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
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
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 151:
		//line parser.y:722
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
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
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 154:
		//line parser.y:743
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 155:
		//line parser.y:750
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
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
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 159:
		//line parser.y:780
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 160:
		//line parser.y:787
		{
		}
	case 161:
		//line parser.y:788
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 162:
		//line parser.y:789
		{
		}
	case 163:
		//line parser.y:792
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 164:
		//line parser.y:795
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 165:
		//line parser.y:803
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 166:
		//line parser.y:805
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 167:
		//line parser.y:807
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 168:
		//line parser.y:816
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
	case 169:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 170:
		//line parser.y:833
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
