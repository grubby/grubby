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

//line parser.y:834

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 25,
	46, 100,
	-2, 20,
	-1, 81,
	10, 73,
	11, 73,
	-2, 159,
	-1, 91,
	46, 100,
	-2, 20,
	-1, 97,
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
	-1, 113,
	46, 100,
	-2, 98,
	-1, 208,
	46, 101,
	-2, 99,
	-1, 211,
	10, 73,
	11, 73,
	-2, 159,
}

const RubyNprod = 170
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1561

var RubyAct = []int{

	9, 177, 21, 33, 175, 83, 89, 82, 174, 66,
	131, 159, 2, 3, 203, 201, 271, 203, 160, 80,
	70, 125, 255, 126, 135, 88, 296, 272, 252, 4,
	79, 229, 103, 102, 77, 199, 101, 64, 63, 301,
	183, 112, 114, 66, 66, 178, 122, 153, 203, 99,
	68, 69, 292, 289, 65, 71, 72, 73, 202, 254,
	132, 270, 133, 67, 76, 74, 75, 78, 70, 138,
	139, 140, 141, 127, 143, 144, 145, 146, 200, 148,
	179, 203, 152, 247, 133, 154, 155, 133, 291, 180,
	203, 88, 151, 70, 116, 99, 227, 294, 68, 69,
	152, 157, 133, 71, 72, 73, 169, 170, 171, 161,
	163, 67, 76, 74, 75, 78, 241, 203, 216, 115,
	181, 295, 100, 68, 69, 194, 263, 242, 265, 264,
	277, 201, 66, 269, 212, 206, 67, 70, 88, 70,
	78, 259, 227, 209, 303, 210, 23, 90, 62, 91,
	81, 113, 86, 96, 214, 225, 201, 208, 217, 194,
	215, 287, 248, 194, 218, 213, 201, 68, 69, 68,
	69, 279, 178, 57, 58, 176, 197, 85, 198, 231,
	67, 228, 67, 194, 78, 194, 78, 234, 194, 194,
	230, 222, 68, 69, 84, 107, 97, 70, 56, 55,
	243, 128, 244, 129, 111, 67, 110, 179, 88, 78,
	186, 168, 152, 251, 133, 156, 180, 150, 137, 61,
	87, 194, 250, 205, 194, 105, 106, 68, 69, 258,
	26, 130, 71, 72, 73, 221, 196, 204, 104, 1,
	67, 76, 74, 75, 78, 237, 123, 53, 52, 274,
	276, 51, 50, 49, 48, 92, 17, 16, 194, 15,
	14, 39, 194, 194, 12, 11, 22, 10, 194, 194,
	20, 35, 13, 18, 34, 31, 92, 30, 32, 29,
	0, 0, 0, 0, 0, 194, 194, 194, 0, 0,
	92, 0, 0, 0, 194, 0, 95, 299, 194, 92,
	92, 92, 92, 0, 92, 92, 92, 92, 0, 92,
	0, 0, 92, 0, 0, 92, 92, 95, 0, 0,
	0, 92, 0, 0, 0, 0, 0, 0, 0, 0,
	92, 95, 0, 0, 0, 0, 92, 92, 92, 0,
	95, 95, 95, 95, 107, 95, 95, 95, 95, 0,
	95, 0, 0, 95, 0, 108, 95, 95, 109, 0,
	0, 0, 95, 0, 0, 92, 0, 0, 92, 0,
	0, 95, 0, 0, 105, 106, 0, 95, 95, 95,
	0, 0, 0, 0, 92, 0, 0, 104, 92, 0,
	27, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 134, 0, 0, 0, 0, 95, 0, 0, 95,
	0, 0, 0, 0, 0, 93, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 95, 0, 0, 0, 95,
	92, 28, 92, 0, 0, 0, 93, 0, 92, 0,
	0, 0, 92, 92, 0, 0, 0, 0, 0, 124,
	93, 0, 0, 0, 0, 0, 94, 0, 0, 93,
	93, 93, 93, 0, 93, 93, 93, 93, 0, 93,
	0, 95, 93, 95, 0, 93, 93, 94, 0, 95,
	92, 93, 0, 95, 95, 0, 0, 0, 0, 0,
	93, 94, 0, 0, 0, 0, 93, 93, 93, 0,
	94, 94, 94, 94, 0, 94, 94, 94, 94, 0,
	94, 0, 0, 94, 0, 0, 94, 94, 0, 0,
	0, 95, 94, 0, 0, 93, 0, 92, 93, 0,
	0, 94, 0, 0, 0, 0, 0, 94, 94, 94,
	0, 0, 0, 0, 93, 0, 158, 162, 93, 0,
	19, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 182, 0, 184, 0, 94, 0, 95, 94,
	0, 187, 188, 0, 0, 98, 226, 0, 0, 0,
	0, 0, 0, 232, 0, 94, 0, 0, 0, 94,
	93, 0, 93, 0, 0, 0, 98, 0, 93, 0,
	0, 0, 93, 93, 0, 0, 245, 246, 0, 220,
	98, 223, 0, 0, 0, 0, 0, 0, 0, 98,
	98, 98, 98, 0, 98, 98, 98, 98, 0, 98,
	0, 94, 98, 94, 0, 98, 98, 0, 0, 94,
	93, 98, 0, 94, 94, 0, 0, 0, 0, 273,
	98, 0, 0, 0, 0, 0, 98, 98, 98, 280,
	0, 0, 0, 281, 0, 0, 0, 0, 195, 0,
	5, 0, 0, 0, 0, 0, 257, 290, 0, 0,
	0, 94, 261, 0, 262, 98, 0, 93, 98, 267,
	0, 0, 268, 0, 0, 0, 0, 0, 0, 0,
	300, 302, 0, 304, 98, 305, 0, 0, 98, 117,
	118, 119, 120, 121, 0, 284, 285, 0, 0, 286,
	0, 0, 0, 0, 0, 0, 0, 0, 94, 0,
	0, 0, 0, 293, 0, 136, 0, 0, 0, 0,
	0, 142, 0, 0, 297, 0, 147, 0, 149, 0,
	98, 0, 98, 0, 0, 0, 0, 0, 98, 0,
	0, 0, 98, 98, 0, 0, 0, 0, 0, 164,
	165, 166, 167, 0, 0, 0, 0, 172, 173, 0,
	0, 0, 0, 0, 185, 0, 23, 207, 62, 91,
	46, 0, 189, 23, 24, 62, 25, 46, 0, 0,
	98, 36, 236, 44, 239, 238, 45, 37, 38, 0,
	0, 0, 47, 57, 58, 0, 0, 0, 0, 54,
	57, 58, 0, 0, 203, 40, 41, 42, 43, 0,
	0, 192, 193, 0, 59, 0, 60, 0, 56, 55,
	0, 59, 0, 60, 0, 56, 55, 98, 0, 23,
	24, 62, 25, 46, 0, 0, 0, 36, 0, 44,
	0, 0, 45, 37, 38, 0, 0, 0, 47, 0,
	0, 0, 0, 0, 0, 54, 57, 58, 0, 0,
	0, 40, 41, 42, 43, 0, 0, 6, 7, 0,
	0, 0, 0, 0, 0, 0, 0, 59, 260, 60,
	0, 56, 55, 0, 8, 0, 0, 266, 0, 0,
	23, 24, 62, 25, 46, 0, 0, 275, 36, 282,
	44, 278, 0, 45, 37, 38, 0, 0, 0, 47,
	0, 0, 0, 283, 0, 0, 54, 57, 58, 0,
	288, 0, 40, 41, 42, 43, 0, 0, 192, 193,
	0, 0, 0, 0, 0, 0, 0, 0, 59, 0,
	60, 0, 56, 55, 298, 23, 24, 62, 25, 46,
	0, 0, 0, 36, 0, 44, 0, 0, 45, 37,
	38, 0, 0, 0, 47, 0, 0, 0, 0, 0,
	0, 54, 57, 58, 0, 0, 0, 40, 41, 42,
	43, 0, 0, 192, 193, 0, 0, 0, 0, 0,
	0, 0, 0, 59, 0, 60, 256, 56, 55, 23,
	24, 62, 25, 46, 0, 0, 0, 36, 253, 44,
	0, 0, 45, 37, 38, 0, 0, 0, 47, 0,
	0, 0, 0, 0, 0, 54, 57, 58, 0, 0,
	0, 40, 41, 42, 43, 0, 0, 192, 193, 0,
	0, 0, 0, 0, 0, 0, 0, 59, 0, 60,
	0, 56, 55, 23, 24, 62, 25, 46, 0, 0,
	0, 36, 240, 44, 0, 0, 45, 37, 38, 0,
	0, 0, 47, 0, 0, 0, 0, 0, 0, 54,
	57, 58, 0, 0, 0, 40, 41, 42, 43, 0,
	0, 192, 193, 0, 0, 0, 0, 0, 0, 0,
	0, 59, 0, 60, 0, 56, 55, 23, 24, 62,
	25, 46, 0, 0, 0, 36, 235, 44, 0, 0,
	45, 37, 38, 0, 0, 0, 47, 0, 0, 0,
	0, 0, 0, 54, 57, 58, 0, 0, 0, 40,
	41, 42, 43, 0, 0, 192, 193, 0, 0, 0,
	0, 0, 0, 0, 0, 59, 0, 60, 0, 56,
	55, 23, 24, 62, 25, 46, 0, 0, 0, 36,
	233, 44, 0, 0, 45, 37, 38, 0, 0, 0,
	47, 0, 0, 0, 0, 0, 0, 54, 57, 58,
	0, 0, 0, 40, 41, 42, 43, 0, 0, 192,
	193, 0, 23, 24, 62, 25, 46, 0, 0, 59,
	36, 60, 44, 56, 55, 45, 37, 38, 0, 0,
	0, 47, 0, 0, 0, 0, 0, 0, 54, 57,
	58, 0, 0, 0, 40, 41, 42, 43, 0, 0,
	192, 193, 0, 0, 0, 0, 0, 0, 0, 0,
	59, 0, 60, 224, 56, 55, 23, 24, 62, 25,
	46, 0, 0, 0, 36, 219, 44, 0, 0, 45,
	37, 38, 0, 0, 0, 47, 0, 0, 0, 0,
	0, 0, 54, 57, 58, 0, 0, 0, 40, 41,
	42, 43, 0, 0, 192, 193, 0, 23, 24, 62,
	25, 46, 0, 0, 59, 36, 60, 44, 56, 55,
	45, 37, 38, 0, 0, 0, 47, 0, 0, 0,
	0, 0, 0, 54, 57, 58, 0, 0, 0, 40,
	41, 42, 43, 0, 0, 192, 193, 0, 23, 24,
	62, 25, 46, 191, 0, 59, 36, 60, 44, 56,
	55, 45, 37, 38, 0, 0, 0, 47, 0, 0,
	0, 0, 0, 0, 54, 57, 58, 0, 0, 0,
	40, 41, 42, 43, 0, 0, 0, 190, 0, 23,
	24, 62, 25, 46, 0, 0, 59, 36, 60, 44,
	56, 55, 45, 37, 38, 0, 0, 0, 47, 0,
	0, 0, 0, 0, 0, 54, 57, 58, 0, 0,
	0, 40, 41, 42, 43, 23, 90, 62, 91, 81,
	0, 0, 96, 0, 0, 0, 0, 59, 0, 60,
	0, 56, 55, 23, 90, 62, 91, 46, 0, 0,
	96, 0, 57, 58, 23, 90, 62, 91, 211, 0,
	0, 96, 0, 0, 0, 249, 0, 0, 0, 0,
	57, 58, 0, 84, 0, 97, 0, 56, 55, 0,
	0, 57, 58, 23, 90, 62, 91, 81, 0, 0,
	96, 59, 0, 97, 0, 56, 55, 23, 90, 62,
	91, 46, 59, 0, 97, 0, 56, 55, 0, 0,
	57, 58, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 57, 58, 0, 0, 0, 0,
	0, 84, 0, 97, 0, 56, 55, 0, 0, 0,
	0, 0, 0, 0, 0, 59, 0, 60, 0, 56,
	55,
}
var RubyPact = []int{

	-31, 844, -1000, -1000, -1000, -6, -1000, -1000, -1000, 16,
	-1000, -1000, -1000, 12, -1000, -1000, -1000, -1000, -1000, -28,
	-1000, -1000, -1000, -1000, 141, 86, 0, -3, -4, -1000,
	-1000, -1000, -1000, -1000, -1000, 340, 200, 143, 143, 83,
	1394, 1394, 1394, 1394, 1394, 1502, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 15, 195, -1000, -1000, 1502,
	-1000, -22, -1000, -1000, -1000, -1000, 1394, 212, 1502, 1502,
	1502, 1502, 1394, 1502, 1502, 1502, 1502, 1394, 1502, 1394,
	211, 1502, -1000, 36, 1502, 1502, 209, 90, 135, -1000,
	1488, 40, -1000, -1000, -1000, 191, -30, -30, -28, 1502,
	1394, 1394, 1394, 1394, 205, 1502, 1502, 1502, 1394, 1394,
	166, 166, 6, -1000, -1000, 1394, 204, 29, 29, 29,
	29, 29, 135, 1353, 1312, -1000, -1000, 170, -1000, -1000,
	24, 4, 193, -1000, 781, 149, 29, 1459, 135, 135,
	135, 135, 29, 135, 135, 135, 135, 29, 89, 29,
	-1000, 155, 193, 1448, 64, 135, -1000, 1448, 1271, -1000,
	185, -1000, 1217, 145, 29, 29, 29, 29, -1000, 135,
	135, 135, 29, 29, -1000, 85, 39, -1000, -5, 184,
	173, -1000, 1176, 143, 1122, 29, -1000, 788, 1068, 29,
	-1000, -1000, -1000, -1000, 16, 29, 102, -1000, -1000, 1502,
	-1000, 1502, -1000, -1000, -1000, 72, 158, 1430, -1000, 36,
	-1000, 1502, 1502, -1000, 135, -1000, -8, 135, -1000, -1000,
	1014, 11, -1000, 960, -1000, -1000, -26, 39, 131, 1394,
	-1000, -1000, -26, -1000, -1000, -1000, -1000, 112, 1394, -1000,
	-1000, -1000, 125, 135, 135, 5, -29, -1000, 1394, 1502,
	120, 135, 1394, -1000, -1000, 165, -1000, 1312, -1000, -1000,
	29, 1312, 905, -1000, 1394, -1000, 29, 1312, 1312, 157,
	-1000, 1394, -1000, 47, 16, 29, 135, -1000, 29, -1000,
	74, 38, -1000, 29, 1312, 1312, 1312, 91, 117, -19,
	-26, -1000, -1000, 1312, -1000, 1394, 1502, 1312, 28, 133,
	-26, -1000, -26, -1000, -26, -26,
}
var RubyPgo = []int{

	0, 668, 279, 278, 6, 277, 275, 550, 431, 274,
	273, 272, 270, 0, 390, 271, 267, 266, 2, 265,
	1, 230, 264, 3, 261, 260, 259, 257, 256, 254,
	253, 252, 251, 248, 247, 449, 246, 7, 11, 245,
	239, 8, 237, 236, 10, 235, 231, 223, 220, 5,
	4, 219, 401,
}
var RubyR1 = []int{

	0, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 52, 52, 35, 35, 35, 35, 35, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 18, 18, 18, 18, 18, 18, 1, 1, 1,
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
	10, 10, 10, 10, 10, 39, 39, 39, 39, 36,
	36, 36, 7, 12, 43, 43, 43, 43, 17, 9,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 4, 1, 4, 2, 3, 3, 4, 4, 3,
	2, 3, 3, 3, 3, 3, 3, 4, 6, 3,
	1, 1, 3, 0, 1, 1, 3, 1, 1, 3,
	3, 1, 3, 3, 7, 7, 1, 3, 1, 2,
	3, 2, 0, 1, 3, 4, 6, 4, 1, 3,
	1, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	2, 2, 3, 3, 3, 3, 2, 2, 2, 2,
	3, 3, 3, 3, 3, 3, 1, 1, 3, 3,
	5, 5, 0, 4, 7, 8, 3, 7, 8, 3,
	4, 4, 3, 3, 0, 1, 3, 4, 5, 3,
	3, 3, 3, 3, 4, 4, 3, 3, 2, 0,
	2, 2, 3, 4, 0, 3, 4, 6, 1, 5,
}
var RubyChk = []int{

	-1000, -40, 43, 44, 60, -1, 43, 44, 60, -13,
	-16, -19, -22, -11, -25, -26, -27, -28, -10, -7,
	-12, -18, -17, 5, 6, 8, -21, -14, -8, -2,
	-5, -6, -3, -23, -9, -15, 13, 19, 20, -24,
	37, 38, 39, 40, 15, 18, 9, 24, -29, -30,
	-31, -32, -33, -34, 31, 58, 57, 32, 33, 53,
	55, -51, 7, 44, 43, 60, 15, 47, 34, 35,
	4, 39, 40, 41, 49, 50, 48, 18, 51, 18,
	47, 9, -37, -49, 53, 36, 11, -48, -13, -4,
	6, 8, -21, -14, -8, -15, 12, 55, -7, 9,
	36, 36, 36, 36, 47, 34, 35, 4, 15, 18,
	6, 4, -23, 8, -23, 36, 11, -1, -1, -1,
	-1, -1, -13, -36, -35, 6, 8, 58, 6, 8,
	-46, -44, -13, -18, -52, 46, -1, 6, -13, -13,
	-13, -13, -1, -13, -13, -13, -13, -1, -13, -1,
	6, -44, -13, 11, -13, -13, 6, 11, -35, -38,
	48, -38, -35, -44, -1, -1, -1, -1, 6, -13,
	-13, -13, -1, -1, -41, -50, 9, -20, 6, 41,
	50, -41, -35, 34, -35, -1, 6, -35, -35, -1,
	44, 10, 43, 44, -13, -1, -43, 6, 8, 11,
	54, 11, 54, 43, -42, -47, -13, 6, 8, -49,
	-37, 9, 45, 10, -13, -4, 54, -13, -4, 14,
	-35, -45, 6, -35, 56, 10, -52, 11, -50, 36,
	6, 6, -52, 14, -23, 14, 14, -39, 17, 16,
	14, 14, 25, -13, -13, -52, -52, 11, 4, 45,
	-44, -13, 36, 14, 48, 11, 56, -35, -20, 10,
	-1, -35, -35, 14, 17, 16, -1, -35, -35, 8,
	56, 11, 56, -52, -13, -1, -13, 10, -1, 6,
	-52, -52, 14, -1, -35, -35, -35, 4, -1, 6,
	-52, 14, 14, -35, 6, 4, 45, -35, -1, -13,
	-52, 11, -52, 11, -52, -52,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 37,
	38, 39, 40, 41, 42, 43, 44, 45, 46, 47,
	48, 49, 50, 18, 19, -2, 21, 22, 23, 24,
	25, 26, 27, 28, 29, 30, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 159, 13, 31, 32,
	33, 34, 35, 36, 168, 0, 0, 126, 127, 73,
	11, 0, 52, 5, 6, 7, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, -2, 54, 60, 73, 0, 0, 70, 77, 78,
	19, -2, 21, 22, 23, 30, 13, -2, 0, 73,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	92, 92, 13, -2, 13, 0, 0, 116, 117, 118,
	119, 13, 13, 0, 164, 110, 111, 0, 108, 109,
	0, 0, 71, 75, 132, 0, 149, 55, 61, 63,
	65, 120, 121, 122, 123, 124, 125, 151, 0, 153,
	59, 0, 74, 0, 71, 102, 114, 0, 0, 13,
	144, 13, 0, 0, 103, 104, 105, 106, 56, 62,
	64, 66, 150, 152, 11, 86, 92, 93, 88, 0,
	0, 11, 0, 0, 0, 107, 115, 0, 0, 160,
	161, 162, 14, 15, 16, 17, 0, 112, 113, 0,
	128, 0, 129, 12, 11, 11, 0, 19, -2, 57,
	58, -2, 0, 51, 79, 80, 67, 82, 83, 139,
	0, 0, 145, 0, 142, 53, 13, 0, 0, 0,
	89, 91, 13, 95, 13, 97, 147, 0, 0, 13,
	154, 163, 13, 72, 76, 0, 0, 11, 0, 0,
	0, 169, 0, 140, 143, 0, 141, 11, 94, 87,
	90, 11, 0, 148, 0, 13, 13, 158, 165, 13,
	130, 0, 131, 0, 37, 11, 136, 69, 68, 146,
	0, 0, 96, 13, 156, 157, 166, 0, 0, 0,
	133, 84, 85, 155, 13, 0, 0, 167, 11, 11,
	134, 11, 137, 11, 135, 138,
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
		//line parser.y:318
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 67:
		//line parser.y:328
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 68:
		//line parser.y:338
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 69:
		//line parser.y:347
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 70:
		//line parser.y:349
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 71:
		//line parser.y:352
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 72:
		//line parser.y:354
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:356
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 74:
		//line parser.y:358
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 75:
		//line parser.y:360
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:362
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:366
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 78:
		//line parser.y:368
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 79:
		//line parser.y:370
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 80:
		//line parser.y:372
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 81:
		//line parser.y:375
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 82:
		//line parser.y:377
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 83:
		//line parser.y:379
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 84:
		//line parser.y:385
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 85:
		//line parser.y:393
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 86:
		//line parser.y:402
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 87:
		//line parser.y:404
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 88:
		//line parser.y:407
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 89:
		//line parser.y:409
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 90:
		//line parser.y:411
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 91:
		//line parser.y:413
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 92:
		//line parser.y:415
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 93:
		//line parser.y:417
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 94:
		//line parser.y:421
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 95:
		//line parser.y:426
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 96:
		//line parser.y:433
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 97:
		//line parser.y:443
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 98:
		//line parser.y:452
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 99:
		//line parser.y:458
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 100:
		//line parser.y:466
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 101:
		//line parser.y:470
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 102:
		//line parser.y:475
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 103:
		//line parser.y:482
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 104:
		//line parser.y:489
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 105:
		//line parser.y:496
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 106:
		//line parser.y:503
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 107:
		//line parser.y:510
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 108:
		//line parser.y:518
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 109:
		//line parser.y:520
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 110:
		//line parser.y:523
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 111:
		//line parser.y:525
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 112:
		//line parser.y:528
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 113:
		//line parser.y:530
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 114:
		//line parser.y:533
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 115:
		//line parser.y:535
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 116:
		//line parser.y:537
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 117:
		//line parser.y:538
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 118:
		//line parser.y:539
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 119:
		//line parser.y:540
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 120:
		//line parser.y:543
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 121:
		//line parser.y:552
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 122:
		//line parser.y:561
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 123:
		//line parser.y:570
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 124:
		//line parser.y:579
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 125:
		//line parser.y:588
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 126:
		//line parser.y:596
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 127:
		//line parser.y:597
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 128:
		//line parser.y:599
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 129:
		//line parser.y:601
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
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
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 132:
		//line parser.y:620
		{
			RubyVAL.genericSlice = ast.Nodes{}
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
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 135:
		//line parser.y:636
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 136:
		//line parser.y:644
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 137:
		//line parser.y:651
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 138:
		//line parser.y:658
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 139:
		//line parser.y:666
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
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
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 142:
		//line parser.y:682
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 143:
		//line parser.y:685
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 144:
		//line parser.y:687
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 145:
		//line parser.y:689
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 146:
		//line parser.y:691
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 147:
		//line parser.y:694
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 148:
		//line parser.y:701
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
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
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 151:
		//line parser.y:723
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
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
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 154:
		//line parser.y:744
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-2].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 155:
		//line parser.y:752
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 156:
		//line parser.y:759
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 157:
		//line parser.y:766
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 158:
		//line parser.y:773
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 159:
		//line parser.y:780
		{
		}
	case 160:
		//line parser.y:781
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 161:
		//line parser.y:782
		{
		}
	case 162:
		//line parser.y:785
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 163:
		//line parser.y:788
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 164:
		//line parser.y:796
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 165:
		//line parser.y:798
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 166:
		//line parser.y:800
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 167:
		//line parser.y:809
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
	case 169:
		//line parser.y:826
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
