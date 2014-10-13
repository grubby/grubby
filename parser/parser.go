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

//line parser.y:865

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 25,
	46, 101,
	-2, 20,
	-1, 82,
	10, 74,
	11, 74,
	-2, 162,
	-1, 92,
	46, 101,
	-2, 20,
	-1, 98,
	13, 13,
	15, 13,
	18, 13,
	19, 13,
	20, 13,
	24, 13,
	30, 13,
	31, 13,
	37, 13,
	38, 13,
	39, 13,
	40, 13,
	44, 13,
	-2, 11,
	-1, 113,
	46, 101,
	-2, 99,
	-1, 209,
	46, 102,
	-2, 100,
	-1, 216,
	10, 74,
	11, 74,
	-2, 162,
}

const RubyNprod = 174
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1489

var RubyAct = []int{

	9, 180, 242, 33, 178, 20, 90, 177, 84, 83,
	133, 202, 162, 2, 3, 204, 163, 23, 91, 63,
	92, 82, 109, 87, 97, 89, 262, 137, 282, 68,
	4, 307, 23, 91, 63, 92, 64, 204, 127, 97,
	128, 259, 112, 114, 58, 59, 234, 103, 86, 102,
	68, 198, 101, 132, 203, 125, 125, 191, 192, 58,
	59, 134, 186, 261, 300, 85, 135, 98, 81, 57,
	56, 142, 143, 144, 145, 68, 147, 148, 149, 150,
	60, 152, 98, 155, 57, 56, 157, 158, 135, 281,
	129, 135, 89, 154, 201, 23, 91, 63, 92, 82,
	155, 204, 97, 66, 65, 135, 172, 173, 124, 126,
	166, 164, 303, 296, 99, 272, 271, 302, 156, 184,
	67, 204, 58, 59, 195, 23, 91, 63, 92, 216,
	181, 99, 97, 179, 280, 255, 138, 207, 72, 155,
	72, 204, 89, 85, 135, 98, 204, 57, 56, 210,
	214, 215, 58, 59, 181, 116, 253, 219, 100, 105,
	106, 222, 195, 220, 232, 182, 195, 223, 70, 71,
	70, 71, 104, 60, 183, 98, 198, 57, 56, 217,
	115, 69, 160, 69, 233, 80, 195, 80, 195, 182,
	239, 195, 23, 208, 63, 92, 64, 247, 183, 249,
	199, 312, 200, 250, 107, 68, 279, 108, 248, 89,
	23, 91, 63, 92, 64, 113, 209, 155, 258, 58,
	59, 306, 135, 105, 106, 305, 195, 257, 289, 195,
	204, 130, 68, 131, 265, 236, 104, 58, 59, 111,
	60, 110, 61, 235, 57, 56, 195, 195, 276, 227,
	72, 26, 287, 202, 189, 284, 286, 314, 60, 176,
	61, 171, 57, 56, 270, 195, 272, 271, 298, 195,
	195, 254, 266, 232, 159, 195, 93, 141, 62, 195,
	70, 71, 88, 23, 91, 63, 92, 82, 206, 226,
	97, 35, 197, 69, 205, 195, 195, 80, 195, 256,
	202, 70, 71, 230, 202, 195, 93, 93, 310, 195,
	58, 59, 93, 1, 69, 139, 96, 53, 80, 218,
	202, 52, 93, 93, 93, 93, 51, 93, 93, 93,
	93, 85, 93, 98, 93, 57, 56, 93, 93, 50,
	49, 48, 17, 93, 16, 15, 96, 96, 14, 40,
	12, 93, 96, 11, 22, 21, 10, 93, 93, 19,
	13, 18, 96, 96, 96, 96, 34, 96, 96, 96,
	96, 36, 96, 31, 96, 30, 32, 96, 96, 29,
	0, 0, 0, 96, 0, 0, 0, 0, 93, 0,
	93, 96, 0, 93, 0, 0, 0, 96, 96, 0,
	0, 0, 0, 0, 0, 0, 27, 0, 93, 0,
	0, 0, 93, 0, 0, 136, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 96, 0,
	96, 94, 0, 96, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 96, 28,
	93, 0, 96, 0, 93, 0, 0, 0, 0, 0,
	93, 94, 94, 0, 0, 0, 0, 94, 93, 93,
	0, 0, 0, 0, 95, 0, 0, 94, 94, 94,
	94, 0, 94, 94, 94, 94, 0, 94, 0, 94,
	96, 0, 94, 94, 96, 0, 0, 0, 94, 0,
	96, 0, 0, 0, 95, 95, 94, 93, 96, 96,
	95, 0, 94, 94, 0, 0, 0, 0, 0, 0,
	95, 95, 95, 95, 0, 95, 95, 95, 95, 0,
	95, 0, 95, 0, 0, 95, 95, 0, 196, 0,
	5, 95, 0, 94, 0, 94, 0, 96, 94, 95,
	0, 0, 0, 0, 0, 95, 95, 123, 0, 93,
	0, 0, 0, 94, 0, 0, 0, 94, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	117, 118, 119, 120, 121, 122, 95, 0, 95, 0,
	0, 95, 0, 231, 0, 0, 0, 0, 0, 96,
	237, 0, 0, 0, 0, 94, 95, 140, 0, 94,
	95, 0, 0, 146, 0, 94, 0, 0, 151, 0,
	153, 251, 252, 94, 94, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 167,
	168, 169, 170, 0, 0, 0, 174, 175, 95, 0,
	0, 0, 95, 0, 188, 161, 165, 0, 95, 0,
	0, 0, 94, 72, 0, 0, 95, 95, 0, 283,
	185, 0, 187, 0, 0, 0, 0, 0, 211, 190,
	290, 0, 0, 0, 291, 0, 0, 0, 0, 0,
	0, 0, 0, 70, 71, 0, 0, 0, 73, 74,
	75, 301, 0, 0, 0, 95, 69, 78, 76, 77,
	80, 0, 0, 221, 94, 0, 0, 0, 0, 0,
	225, 0, 228, 0, 0, 311, 313, 0, 315, 0,
	316, 0, 0, 23, 24, 63, 25, 64, 0, 0,
	0, 37, 275, 45, 244, 243, 46, 38, 39, 245,
	246, 0, 47, 0, 0, 0, 0, 95, 55, 54,
	58, 59, 0, 0, 0, 41, 42, 43, 44, 0,
	0, 193, 194, 267, 0, 0, 0, 0, 0, 0,
	0, 60, 273, 61, 0, 57, 56, 0, 0, 264,
	0, 0, 0, 285, 0, 268, 0, 269, 288, 0,
	0, 0, 274, 0, 0, 0, 278, 0, 0, 0,
	293, 0, 0, 0, 23, 24, 63, 25, 64, 0,
	299, 0, 37, 241, 45, 244, 243, 46, 38, 39,
	294, 295, 0, 47, 0, 0, 0, 297, 0, 55,
	54, 58, 59, 0, 0, 309, 41, 42, 43, 44,
	0, 304, 193, 194, 0, 0, 0, 0, 0, 0,
	0, 0, 60, 308, 61, 0, 57, 56, 23, 24,
	63, 25, 64, 0, 0, 0, 37, 0, 45, 0,
	0, 46, 38, 39, 0, 0, 0, 47, 0, 0,
	0, 0, 0, 55, 54, 58, 59, 0, 0, 0,
	41, 42, 43, 44, 0, 0, 6, 7, 0, 0,
	0, 0, 0, 0, 0, 0, 60, 0, 61, 0,
	57, 56, 0, 8, 23, 24, 63, 25, 64, 0,
	0, 0, 37, 292, 45, 0, 0, 46, 38, 39,
	0, 0, 0, 47, 0, 0, 0, 0, 0, 55,
	54, 58, 59, 0, 0, 0, 41, 42, 43, 44,
	0, 0, 193, 194, 0, 0, 0, 0, 0, 0,
	0, 0, 60, 0, 61, 0, 57, 56, 23, 24,
	63, 25, 64, 0, 0, 0, 37, 277, 45, 0,
	0, 46, 38, 39, 0, 0, 0, 47, 0, 0,
	0, 0, 0, 55, 54, 58, 59, 0, 0, 0,
	41, 42, 43, 44, 0, 0, 193, 194, 0, 23,
	24, 63, 25, 64, 0, 0, 60, 37, 61, 45,
	57, 56, 46, 38, 39, 0, 0, 0, 47, 0,
	0, 0, 0, 0, 55, 54, 58, 59, 0, 0,
	0, 41, 42, 43, 44, 0, 0, 193, 194, 0,
	0, 0, 0, 0, 0, 0, 0, 60, 0, 61,
	263, 57, 56, 23, 24, 63, 25, 64, 0, 0,
	0, 37, 260, 45, 0, 0, 46, 38, 39, 0,
	0, 0, 47, 0, 0, 0, 0, 0, 55, 54,
	58, 59, 0, 0, 0, 41, 42, 43, 44, 0,
	0, 193, 194, 0, 0, 0, 0, 0, 0, 0,
	0, 60, 0, 61, 0, 57, 56, 23, 24, 63,
	25, 64, 0, 0, 0, 37, 240, 45, 0, 0,
	46, 38, 39, 0, 0, 0, 47, 0, 0, 0,
	0, 0, 55, 54, 58, 59, 0, 0, 0, 41,
	42, 43, 44, 0, 0, 193, 194, 0, 0, 0,
	0, 0, 0, 0, 0, 60, 0, 61, 0, 57,
	56, 23, 24, 63, 25, 64, 0, 0, 0, 37,
	238, 45, 0, 0, 46, 38, 39, 0, 0, 0,
	47, 0, 0, 0, 0, 0, 55, 54, 58, 59,
	0, 0, 0, 41, 42, 43, 44, 0, 0, 193,
	194, 0, 23, 24, 63, 25, 64, 0, 0, 60,
	37, 61, 45, 57, 56, 46, 38, 39, 0, 0,
	0, 47, 0, 0, 0, 0, 0, 55, 54, 58,
	59, 0, 0, 0, 41, 42, 43, 44, 0, 0,
	193, 194, 0, 0, 0, 0, 0, 0, 0, 0,
	60, 0, 61, 229, 57, 56, 23, 24, 63, 25,
	64, 0, 0, 0, 37, 224, 45, 0, 0, 46,
	38, 39, 0, 0, 0, 47, 0, 0, 0, 0,
	0, 55, 54, 58, 59, 0, 0, 0, 41, 42,
	43, 44, 0, 0, 193, 194, 0, 23, 24, 63,
	25, 64, 0, 0, 60, 37, 61, 45, 57, 56,
	46, 38, 39, 0, 0, 0, 47, 0, 0, 0,
	0, 0, 55, 54, 58, 59, 0, 0, 0, 41,
	42, 43, 44, 0, 0, 193, 194, 0, 23, 24,
	63, 25, 64, 213, 0, 60, 37, 61, 45, 57,
	56, 46, 38, 39, 0, 0, 0, 47, 0, 0,
	0, 0, 0, 55, 54, 58, 59, 0, 0, 0,
	41, 42, 43, 44, 0, 0, 0, 212, 0, 23,
	24, 63, 25, 64, 0, 0, 60, 37, 61, 45,
	57, 56, 46, 38, 39, 0, 0, 0, 47, 0,
	0, 0, 0, 72, 55, 54, 58, 59, 0, 0,
	0, 41, 42, 43, 44, 0, 0, 79, 0, 0,
	0, 72, 0, 0, 0, 0, 0, 60, 0, 61,
	0, 57, 56, 70, 71, 0, 0, 0, 73, 74,
	75, 0, 0, 0, 0, 0, 69, 78, 76, 77,
	80, 70, 71, 0, 0, 0, 73, 74, 75, 0,
	0, 0, 0, 0, 69, 78, 76, 77, 80,
}
var RubyPact = []int{

	-30, 863, -1000, -1000, -1000, 60, -1000, -1000, -1000, 1419,
	-1000, -1000, -1000, 50, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 12, 122, 16, 13, 11, -1000,
	-1000, -1000, -1000, -1000, -1000, 189, -25, 235, 207, 207,
	144, 1394, 1394, 1394, 1394, 1394, 1394, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 205, 205, 32, 225, -1000, -1000,
	205, -1000, -19, 127, -1000, -1000, -1000, -1000, 1394, 271,
	205, 205, 205, 205, 1394, 205, 205, 205, 205, 1394,
	205, 1394, 205, -1000, 107, 205, 205, 268, 171, 136,
	-1000, 278, 105, -1000, -1000, -1000, 125, -32, -32, 205,
	1394, 1394, 1394, 1394, 255, 205, 205, 1394, 1394, 253,
	124, 124, 28, -1000, -1000, 1394, 248, 35, 35, 35,
	35, 35, 14, 1312, 165, 136, 165, -1000, -1000, 194,
	-1000, -1000, 40, 0, 1437, -1000, 187, 208, 205, 1353,
	35, 120, 136, 136, 136, 136, 35, 136, 136, 136,
	136, 35, 134, 35, 309, 1437, 27, 659, 136, -1000,
	27, 1271, -1000, 243, -1000, 1217, 293, 35, 35, 35,
	35, -1000, 136, 136, 35, 35, -1000, -1000, 153, 148,
	-1000, 10, 237, 229, -1000, 1176, 207, 1122, 35, -1000,
	809, -1000, -1000, -1000, -1000, 1419, 35, 183, 205, -1000,
	-1000, -1000, 205, -1000, -1000, -1000, 145, 267, 90, -1000,
	289, 35, -1000, -1000, 107, -1000, 205, 205, -1000, 136,
	-1000, 5, 136, -1000, -1000, 1068, 15, -1000, 1014, -1000,
	-1000, -6, 148, 262, 1394, -1000, -1000, -6, -1000, -1000,
	-1000, -1000, 250, 1394, -1000, 728, 973, -1000, 198, 136,
	136, 78, -28, -1000, 1394, 205, -1000, 242, 136, 1394,
	-1000, -1000, 222, -1000, 1312, -1000, -1000, 35, 1312, 919,
	-1000, 1394, -1000, 35, 1312, -1000, 99, -1000, 1312, 264,
	-1000, 1394, -1000, 58, 1419, 35, 136, -1000, 35, -1000,
	103, 98, -1000, 35, 1312, 1312, -1000, 1312, 219, 217,
	-14, -6, -1000, -1000, 1312, -1000, 1394, 205, 1312, 190,
	246, -6, -1000, -6, -1000, -6, -6,
}
var RubyPgo = []int{

	0, 538, 379, 376, 6, 375, 373, 371, 449, 366,
	361, 360, 359, 0, 406, 291, 356, 355, 354, 5,
	353, 1, 251, 350, 3, 349, 348, 345, 344, 342,
	341, 340, 339, 326, 321, 317, 557, 315, 9, 12,
	2, 313, 7, 294, 292, 10, 289, 53, 288, 282,
	8, 4, 278, 415,
}
var RubyR1 = []int{

	0, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 53, 53, 36, 36, 36, 36, 36, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 19, 19, 19, 19, 19, 19, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	38, 38, 47, 47, 45, 45, 45, 45, 50, 50,
	50, 50, 49, 49, 49, 16, 16, 42, 42, 21,
	21, 21, 21, 51, 51, 51, 20, 20, 23, 24,
	24, 52, 52, 11, 11, 11, 11, 11, 11, 8,
	8, 22, 22, 14, 14, 25, 25, 26, 27, 28,
	29, 30, 31, 32, 33, 34, 35, 2, 5, 6,
	6, 3, 3, 43, 43, 43, 43, 48, 48, 48,
	4, 4, 4, 4, 39, 46, 46, 46, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 40, 40,
	40, 40, 37, 37, 37, 7, 12, 44, 44, 44,
	44, 17, 18, 9,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 4, 1, 4, 4, 2, 3, 3, 4,
	4, 3, 2, 3, 3, 3, 3, 3, 4, 6,
	3, 1, 1, 3, 0, 1, 1, 3, 1, 1,
	3, 3, 1, 3, 3, 7, 7, 1, 3, 1,
	2, 3, 2, 0, 1, 3, 4, 6, 4, 1,
	3, 1, 3, 3, 3, 3, 3, 3, 3, 2,
	2, 2, 2, 3, 3, 3, 3, 2, 2, 2,
	2, 3, 3, 3, 3, 3, 3, 1, 1, 3,
	3, 5, 5, 0, 4, 7, 8, 3, 7, 8,
	3, 4, 4, 3, 3, 0, 1, 3, 4, 5,
	3, 3, 3, 3, 3, 5, 6, 5, 4, 3,
	3, 2, 0, 2, 2, 3, 4, 0, 3, 4,
	6, 2, 2, 5,
}
var RubyChk = []int{

	-1000, -41, 43, 44, 60, -1, 43, 44, 60, -13,
	-16, -20, -23, -11, -26, -27, -28, -29, -10, -12,
	-19, -17, -18, 5, 6, 8, -22, -14, -8, -2,
	-5, -6, -3, -24, -9, -15, -7, 13, 19, 20,
	-25, 37, 38, 39, 40, 15, 18, 24, -30, -31,
	-32, -33, -34, -35, 31, 30, 58, 57, 32, 33,
	53, 55, -52, 7, 9, 44, 43, 60, 15, 47,
	34, 35, 4, 39, 40, 41, 49, 50, 48, 18,
	51, 18, 9, -38, -50, 53, 36, 11, -49, -13,
	-4, 6, 8, -22, -14, -8, -15, 12, 55, 9,
	36, 36, 36, 36, 47, 34, 35, 15, 18, 47,
	6, 4, -24, 8, -24, 36, 11, -1, -1, -1,
	-1, -1, -1, -36, -47, -13, -47, 6, 8, 58,
	6, 8, -47, -45, -13, -19, -53, 46, 9, -37,
	-1, 6, -13, -13, -13, -13, -1, -13, -13, -13,
	-13, -1, -13, -1, -45, -13, 11, -13, -13, 6,
	11, -36, -39, 48, -39, -36, -45, -1, -1, -1,
	-1, 6, -13, -13, -1, -1, 6, -42, -51, 9,
	-21, 6, 41, 50, -42, -36, 34, -36, -1, 6,
	-36, 43, 44, 43, 44, -13, -1, -44, 11, 6,
	8, 54, 11, 54, 43, -43, -48, -13, 6, 8,
	-45, -1, 44, 10, -50, -38, 9, 45, 10, -13,
	-4, 54, -13, -4, 14, -36, -46, 6, -36, 56,
	10, -53, 11, -51, 36, 6, 6, -53, 14, -24,
	14, 14, -40, 17, 16, -36, -36, 14, 25, -13,
	-13, -53, -53, 11, 4, 45, 10, -45, -13, 36,
	14, 48, 11, 56, -36, -21, 10, -1, -36, -36,
	14, 17, 16, -1, -36, 14, -40, 14, -36, 8,
	56, 11, 56, -53, -13, -1, -13, 10, -1, 6,
	-53, -53, 14, -1, -36, -36, 14, -36, 4, -1,
	6, -53, 14, 14, -36, 6, 4, 45, -36, -1,
	-13, -53, 11, -53, 11, -53, -53,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 38,
	39, 40, 41, 42, 43, 44, 45, 46, 47, 48,
	49, 50, 51, 18, 19, -2, 21, 22, 23, 24,
	25, 26, 27, 28, 29, 30, 31, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 13, 32, 33,
	34, 35, 36, 37, 0, 0, 0, 0, 127, 128,
	74, 11, 0, 53, 162, 5, 6, 7, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, -2, 56, 62, 74, 0, 0, 71, 78,
	79, 19, -2, 21, 22, 23, 30, 13, -2, 74,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	93, 93, 13, -2, 13, 0, 0, 117, 118, 119,
	120, 13, 0, 167, 171, 72, 172, 111, 112, 0,
	109, 110, 0, 0, 72, 76, 133, 0, 74, 0,
	150, 57, 63, 65, 67, 121, 122, 123, 124, 125,
	126, 152, 0, 154, 0, 75, 0, 72, 103, 115,
	0, 0, 13, 145, 13, 0, 0, 104, 105, 106,
	107, 58, 64, 66, 151, 153, 61, 11, 87, 93,
	94, 89, 0, 0, 11, 0, 0, 0, 108, 116,
	0, 13, 13, 14, 15, 16, 17, 0, 0, 113,
	114, 129, 0, 130, 12, 11, 11, 0, 19, -2,
	0, 163, 164, 165, 59, 60, -2, 0, 52, 80,
	81, 68, 83, 84, 140, 0, 0, 146, 0, 143,
	55, 13, 0, 0, 0, 90, 92, 13, 96, 13,
	98, 148, 0, 0, 13, 0, 0, 166, 13, 73,
	77, 0, 0, 11, 0, 0, 54, 0, 173, 0,
	141, 144, 0, 142, 11, 95, 88, 91, 11, 0,
	149, 0, 13, 13, 161, 155, 0, 157, 168, 13,
	131, 0, 132, 0, 38, 11, 137, 70, 69, 147,
	0, 0, 97, 13, 159, 160, 156, 169, 0, 0,
	0, 134, 85, 86, 158, 13, 0, 0, 170, 11,
	11, 135, 11, 138, 11, 136, 139,
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
		//line parser.y:165
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:167
		{
		}
	case 3:
		//line parser.y:169
		{
		}
	case 4:
		//line parser.y:171
		{
		}
	case 5:
		//line parser.y:173
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:175
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:177
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:183
		{
		}
	case 11:
		//line parser.y:185
		{
		}
	case 12:
		//line parser.y:186
		{
		}
	case 13:
		//line parser.y:189
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 14:
		//line parser.y:191
		{
		}
	case 15:
		//line parser.y:193
		{
		}
	case 16:
		//line parser.y:195
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 17:
		//line parser.y:197
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 52:
		//line parser.y:207
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 53:
		//line parser.y:214
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 54:
		//line parser.y:218
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 55:
		//line parser.y:225
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 56:
		//line parser.y:232
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 57:
		//line parser.y:239
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 58:
		//line parser.y:246
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 59:
		//line parser.y:253
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 60:
		//line parser.y:261
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 61:
		//line parser.y:269
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 62:
		//line parser.y:279
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 63:
		//line parser.y:286
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 64:
		//line parser.y:294
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 65:
		//line parser.y:302
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 66:
		//line parser.y:310
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 67:
		//line parser.y:318
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 68:
		//line parser.y:328
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 69:
		//line parser.y:338
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 70:
		//line parser.y:347
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 71:
		//line parser.y:349
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 72:
		//line parser.y:352
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:354
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:356
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 75:
		//line parser.y:358
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:360
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:362
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 78:
		//line parser.y:366
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 79:
		//line parser.y:368
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 80:
		//line parser.y:370
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 81:
		//line parser.y:372
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 82:
		//line parser.y:375
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 83:
		//line parser.y:377
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 84:
		//line parser.y:379
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:385
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 86:
		//line parser.y:393
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 87:
		//line parser.y:402
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 88:
		//line parser.y:404
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 89:
		//line parser.y:407
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 90:
		//line parser.y:409
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 91:
		//line parser.y:411
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 92:
		//line parser.y:413
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 93:
		//line parser.y:415
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 94:
		//line parser.y:417
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 95:
		//line parser.y:421
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 96:
		//line parser.y:426
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 97:
		//line parser.y:433
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 98:
		//line parser.y:443
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 99:
		//line parser.y:452
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 100:
		//line parser.y:458
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 101:
		//line parser.y:466
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 102:
		//line parser.y:470
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 103:
		//line parser.y:475
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 104:
		//line parser.y:482
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 105:
		//line parser.y:489
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 106:
		//line parser.y:496
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 107:
		//line parser.y:503
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 108:
		//line parser.y:510
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 109:
		//line parser.y:518
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 110:
		//line parser.y:520
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 111:
		//line parser.y:523
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 112:
		//line parser.y:525
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 113:
		//line parser.y:528
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 114:
		//line parser.y:530
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 115:
		//line parser.y:533
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 116:
		//line parser.y:535
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 117:
		//line parser.y:537
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 118:
		//line parser.y:538
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 119:
		//line parser.y:539
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 120:
		//line parser.y:540
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 121:
		//line parser.y:543
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 122:
		//line parser.y:552
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 123:
		//line parser.y:561
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 124:
		//line parser.y:570
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 125:
		//line parser.y:579
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 126:
		//line parser.y:588
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 127:
		//line parser.y:596
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 128:
		//line parser.y:597
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 129:
		//line parser.y:599
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 130:
		//line parser.y:601
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 131:
		//line parser.y:604
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 132:
		//line parser.y:612
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 133:
		//line parser.y:620
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 134:
		//line parser.y:622
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 135:
		//line parser.y:629
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 136:
		//line parser.y:636
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 137:
		//line parser.y:644
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 138:
		//line parser.y:651
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 139:
		//line parser.y:658
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 140:
		//line parser.y:666
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 141:
		//line parser.y:668
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 142:
		//line parser.y:675
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 143:
		//line parser.y:682
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 144:
		//line parser.y:685
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 145:
		//line parser.y:687
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 146:
		//line parser.y:689
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 147:
		//line parser.y:691
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 148:
		//line parser.y:694
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 149:
		//line parser.y:701
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 150:
		//line parser.y:709
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 151:
		//line parser.y:716
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 152:
		//line parser.y:723
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 153:
		//line parser.y:730
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 154:
		//line parser.y:737
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 155:
		//line parser.y:744
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 156:
		//line parser.y:751
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 157:
		//line parser.y:759
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 158:
		//line parser.y:768
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 159:
		//line parser.y:775
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 160:
		//line parser.y:782
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 161:
		//line parser.y:789
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 162:
		//line parser.y:796
		{
		}
	case 163:
		//line parser.y:797
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 164:
		//line parser.y:798
		{
		}
	case 165:
		//line parser.y:801
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 166:
		//line parser.y:804
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 167:
		//line parser.y:812
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 168:
		//line parser.y:814
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 169:
		//line parser.y:816
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 170:
		//line parser.y:825
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
	case 171:
		//line parser.y:840
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 172:
		//line parser.y:848
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 173:
		//line parser.y:857
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
