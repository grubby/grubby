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
const NEXT = 57370
const REDO = 57371
const RETRY = 57372
const RETURN = 57373
const YIELD = 57374
const TRUE = 57375
const FALSE = 57376
const LESSTHAN = 57377
const GREATERTHAN = 57378
const EQUALTO = 57379
const BANG = 57380
const COMPLEMENT = 57381
const POSITIVE = 57382
const NEGATIVE = 57383
const STAR = 57384
const WHITESPACE = 57385
const NEWLINE = 57386
const SEMICOLON = 57387
const COLON = 57388
const DOUBLECOLON = 57389
const DOT = 57390
const PIPE = 57391
const SLASH = 57392
const AMPERSAND = 57393
const QUESTIONMARK = 57394
const CARET = 57395
const LBRACKET = 57396
const RBRACKET = 57397
const LBRACE = 57398
const RBRACE = 57399
const DOLLARSIGN = 57400
const ATSIGN = 57401
const FILE_CONST_REF = 57402
const EOF = 57403

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
	"NEXT",
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

//line parser.y:928

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 25,
	47, 111,
	-2, 20,
	-1, 82,
	10, 74,
	11, 74,
	-2, 172,
	-1, 92,
	47, 111,
	-2, 20,
	-1, 98,
	13, 13,
	15, 13,
	18, 13,
	19, 13,
	20, 13,
	22, 13,
	24, 13,
	32, 13,
	38, 13,
	39, 13,
	40, 13,
	41, 13,
	45, 13,
	-2, 11,
	-1, 113,
	47, 111,
	-2, 109,
	-1, 210,
	47, 112,
	-2, 110,
	-1, 217,
	10, 74,
	11, 74,
	-2, 172,
}

const RubyNprod = 198
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1650

var RubyAct = []int{

	9, 132, 243, 266, 180, 33, 177, 84, 83, 136,
	178, 68, 2, 3, 133, 90, 162, 291, 205, 20,
	203, 72, 127, 198, 128, 89, 105, 106, 344, 4,
	163, 292, 264, 109, 137, 335, 205, 261, 235, 104,
	66, 65, 103, 186, 112, 114, 181, 102, 181, 179,
	205, 101, 70, 71, 81, 125, 124, 67, 342, 68,
	327, 134, 68, 290, 204, 69, 196, 202, 5, 80,
	263, 142, 143, 144, 145, 129, 147, 148, 149, 150,
	135, 152, 182, 155, 182, 72, 157, 158, 191, 192,
	205, 183, 89, 183, 321, 116, 323, 154, 324, 72,
	155, 325, 135, 198, 68, 135, 172, 173, 117, 118,
	119, 120, 121, 122, 166, 164, 70, 71, 184, 135,
	156, 115, 126, 99, 195, 68, 205, 218, 255, 69,
	70, 71, 205, 80, 248, 140, 233, 208, 160, 155,
	99, 146, 89, 69, 334, 249, 151, 80, 153, 215,
	216, 100, 256, 211, 199, 68, 318, 220, 135, 319,
	107, 223, 195, 108, 297, 203, 195, 167, 168, 169,
	170, 200, 221, 201, 174, 175, 224, 268, 233, 138,
	105, 106, 188, 70, 71, 281, 195, 232, 195, 113,
	234, 195, 240, 104, 238, 316, 69, 210, 317, 250,
	80, 315, 258, 203, 252, 313, 212, 274, 273, 272,
	89, 274, 273, 130, 26, 131, 253, 254, 155, 260,
	231, 203, 219, 203, 111, 329, 110, 195, 299, 237,
	195, 236, 259, 228, 189, 176, 171, 135, 267, 93,
	159, 141, 270, 62, 88, 207, 227, 195, 195, 278,
	197, 251, 285, 206, 1, 139, 53, 294, 296, 52,
	51, 50, 49, 48, 17, 293, 16, 303, 15, 93,
	14, 303, 195, 40, 305, 93, 300, 195, 12, 11,
	308, 195, 306, 21, 10, 93, 93, 93, 93, 287,
	93, 93, 93, 93, 19, 93, 22, 93, 13, 18,
	93, 93, 269, 34, 36, 322, 93, 31, 125, 326,
	30, 275, 195, 195, 93, 195, 32, 29, 286, 0,
	93, 93, 0, 295, 0, 0, 0, 0, 298, 195,
	0, 0, 0, 304, 0, 0, 340, 304, 0, 195,
	310, 0, 0, 0, 23, 91, 63, 92, 64, 341,
	343, 93, 345, 93, 346, 0, 93, 0, 320, 0,
	0, 0, 0, 0, 35, 0, 0, 0, 0, 0,
	0, 93, 58, 59, 0, 93, 0, 0, 0, 0,
	0, 0, 0, 330, 331, 332, 333, 0, 0, 96,
	0, 336, 337, 60, 0, 61, 0, 57, 56, 0,
	0, 339, 0, 0, 0, 27, 0, 0, 0, 0,
	0, 0, 0, 93, 0, 0, 0, 0, 93, 96,
	0, 0, 0, 0, 93, 96, 0, 0, 0, 0,
	94, 0, 93, 93, 0, 96, 96, 96, 96, 0,
	96, 96, 96, 96, 0, 96, 0, 96, 0, 0,
	96, 96, 0, 0, 0, 0, 96, 0, 0, 0,
	94, 0, 0, 0, 96, 0, 94, 0, 0, 0,
	96, 96, 93, 0, 0, 0, 94, 94, 94, 94,
	0, 94, 94, 94, 94, 0, 94, 0, 94, 0,
	0, 94, 94, 0, 0, 0, 0, 94, 0, 0,
	0, 96, 0, 96, 0, 94, 96, 0, 0, 0,
	0, 94, 94, 0, 0, 123, 0, 0, 0, 0,
	28, 96, 93, 0, 0, 96, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 23, 91, 63, 92,
	82, 0, 94, 97, 94, 95, 0, 94, 0, 0,
	93, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 94, 96, 58, 59, 94, 0, 96, 0,
	0, 0, 0, 0, 96, 95, 0, 257, 0, 0,
	0, 95, 96, 96, 0, 85, 0, 98, 0, 57,
	56, 95, 95, 95, 95, 0, 95, 95, 95, 95,
	0, 95, 72, 95, 94, 0, 95, 95, 0, 94,
	0, 0, 95, 161, 165, 94, 0, 0, 0, 0,
	95, 0, 96, 94, 94, 0, 95, 95, 185, 0,
	187, 0, 0, 70, 71, 0, 0, 190, 73, 74,
	75, 0, 0, 0, 0, 0, 69, 78, 76, 77,
	80, 0, 0, 222, 0, 0, 0, 95, 0, 95,
	0, 0, 95, 94, 0, 0, 0, 0, 0, 0,
	0, 0, 96, 0, 0, 0, 0, 95, 226, 0,
	229, 95, 23, 24, 63, 25, 64, 0, 0, 0,
	37, 282, 45, 0, 0, 46, 38, 39, 0, 55,
	96, 47, 0, 0, 288, 289, 0, 246, 247, 54,
	58, 59, 0, 94, 0, 41, 42, 43, 44, 95,
	0, 283, 284, 0, 95, 0, 0, 0, 0, 0,
	95, 60, 0, 61, 0, 57, 56, 0, 95, 95,
	0, 94, 0, 0, 0, 23, 91, 63, 92, 82,
	0, 87, 97, 0, 0, 0, 271, 0, 0, 0,
	0, 276, 0, 0, 0, 280, 0, 0, 0, 0,
	0, 0, 0, 58, 59, 0, 0, 86, 95, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	311, 312, 0, 0, 85, 0, 98, 314, 57, 56,
	0, 0, 0, 0, 23, 24, 63, 25, 64, 0,
	0, 0, 37, 277, 45, 245, 244, 46, 38, 39,
	0, 55, 0, 47, 0, 0, 328, 0, 95, 0,
	0, 54, 58, 59, 0, 0, 0, 41, 42, 43,
	44, 0, 0, 193, 194, 338, 0, 0, 0, 0,
	0, 0, 0, 60, 0, 61, 95, 57, 56, 23,
	24, 63, 25, 64, 0, 0, 0, 37, 242, 45,
	245, 244, 46, 38, 39, 0, 55, 0, 47, 0,
	0, 0, 0, 0, 0, 0, 54, 58, 59, 0,
	0, 0, 41, 42, 43, 44, 0, 0, 193, 194,
	0, 23, 24, 63, 25, 64, 0, 0, 60, 37,
	61, 45, 57, 56, 46, 38, 39, 0, 55, 0,
	47, 0, 0, 0, 0, 0, 0, 0, 54, 58,
	59, 0, 0, 0, 41, 42, 43, 44, 0, 0,
	6, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	60, 0, 61, 0, 57, 56, 0, 8, 23, 24,
	63, 25, 64, 0, 0, 0, 37, 309, 45, 0,
	0, 46, 38, 39, 0, 55, 0, 47, 0, 0,
	0, 0, 0, 0, 0, 54, 58, 59, 0, 0,
	0, 41, 42, 43, 44, 0, 0, 193, 194, 0,
	23, 24, 63, 25, 64, 0, 0, 60, 37, 61,
	45, 57, 56, 46, 38, 39, 0, 55, 0, 47,
	0, 0, 0, 0, 0, 0, 307, 54, 58, 59,
	0, 0, 0, 41, 42, 43, 44, 0, 0, 301,
	302, 0, 0, 0, 0, 0, 0, 0, 0, 60,
	0, 61, 0, 57, 56, 23, 24, 63, 25, 64,
	0, 0, 0, 37, 279, 45, 0, 0, 46, 38,
	39, 0, 55, 0, 47, 0, 0, 0, 0, 0,
	0, 0, 54, 58, 59, 0, 0, 0, 41, 42,
	43, 44, 0, 0, 193, 194, 0, 23, 24, 63,
	25, 64, 0, 0, 60, 37, 61, 45, 57, 56,
	46, 38, 39, 0, 55, 0, 47, 0, 0, 0,
	0, 0, 0, 0, 54, 58, 59, 0, 0, 0,
	41, 42, 43, 44, 0, 0, 193, 194, 0, 0,
	0, 0, 0, 0, 0, 0, 60, 0, 61, 265,
	57, 56, 23, 24, 63, 25, 64, 0, 0, 0,
	37, 262, 45, 0, 0, 46, 38, 39, 0, 55,
	0, 47, 0, 0, 0, 0, 0, 0, 0, 54,
	58, 59, 0, 0, 0, 41, 42, 43, 44, 0,
	0, 193, 194, 0, 0, 0, 0, 0, 0, 0,
	0, 60, 0, 61, 0, 57, 56, 23, 24, 63,
	25, 64, 0, 0, 0, 37, 241, 45, 0, 0,
	46, 38, 39, 0, 55, 0, 47, 0, 0, 0,
	0, 0, 0, 0, 54, 58, 59, 0, 0, 0,
	41, 42, 43, 44, 0, 0, 193, 194, 0, 0,
	0, 0, 0, 0, 0, 0, 60, 0, 61, 0,
	57, 56, 23, 24, 63, 25, 64, 0, 0, 0,
	37, 239, 45, 0, 0, 46, 38, 39, 0, 55,
	0, 47, 0, 0, 0, 0, 0, 0, 0, 54,
	58, 59, 0, 0, 0, 41, 42, 43, 44, 0,
	0, 193, 194, 0, 23, 24, 63, 25, 64, 0,
	0, 60, 37, 61, 45, 57, 56, 46, 38, 39,
	0, 55, 0, 47, 0, 0, 0, 0, 0, 0,
	0, 54, 58, 59, 0, 0, 0, 41, 42, 43,
	44, 0, 0, 193, 194, 0, 0, 0, 0, 0,
	0, 0, 0, 60, 0, 61, 230, 57, 56, 23,
	24, 63, 25, 64, 0, 0, 0, 37, 225, 45,
	0, 0, 46, 38, 39, 0, 55, 0, 47, 0,
	0, 0, 0, 0, 0, 0, 54, 58, 59, 0,
	0, 0, 41, 42, 43, 44, 0, 0, 193, 194,
	0, 23, 24, 63, 25, 64, 0, 0, 60, 37,
	61, 45, 57, 56, 46, 38, 39, 0, 55, 0,
	47, 0, 0, 0, 0, 0, 0, 0, 54, 58,
	59, 0, 0, 0, 41, 42, 43, 44, 0, 0,
	193, 194, 0, 23, 24, 63, 25, 64, 214, 0,
	60, 37, 61, 45, 57, 56, 46, 38, 39, 0,
	55, 0, 47, 0, 0, 0, 0, 0, 0, 0,
	54, 58, 59, 0, 0, 0, 41, 42, 43, 44,
	0, 0, 0, 213, 0, 23, 24, 63, 25, 64,
	0, 0, 60, 37, 61, 45, 57, 56, 46, 38,
	39, 0, 55, 0, 47, 0, 0, 0, 0, 0,
	0, 0, 54, 58, 59, 0, 0, 0, 41, 42,
	43, 44, 23, 91, 63, 92, 64, 0, 0, 97,
	0, 0, 0, 0, 60, 0, 61, 0, 57, 56,
	23, 91, 63, 92, 217, 0, 0, 97, 0, 0,
	58, 59, 0, 23, 209, 63, 92, 64, 0, 23,
	91, 63, 92, 82, 0, 0, 97, 0, 58, 59,
	0, 60, 0, 98, 0, 57, 56, 0, 0, 0,
	0, 58, 59, 72, 0, 0, 0, 58, 59, 60,
	0, 98, 205, 57, 56, 0, 0, 79, 0, 0,
	0, 72, 60, 0, 61, 0, 57, 56, 85, 0,
	98, 0, 57, 56, 70, 71, 0, 0, 0, 73,
	74, 75, 0, 0, 0, 0, 0, 69, 78, 76,
	77, 80, 70, 71, 0, 0, 0, 73, 74, 75,
	0, 0, 0, 0, 0, 69, 78, 76, 77, 80,
}
var RubyPact = []int{

	-32, 896, -1000, -1000, -1000, -4, -1000, -1000, -1000, 1579,
	-1000, -1000, -1000, 36, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 740, 114, 14, 10, 5, -1000,
	-1000, -1000, -1000, -1000, -1000, 145, -15, 220, 181, 181,
	84, 1480, 1480, 1480, 1480, 1480, 1480, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 339, 1480, 16, 207, -1000, -1000,
	339, -1000, -13, 170, -1000, -1000, -1000, -1000, 1480, 235,
	339, 339, 339, 339, 1480, 339, 339, 339, 339, 1480,
	339, 1480, 339, -1000, 109, 339, 339, 234, 127, 95,
	-1000, 1554, 131, -1000, -1000, -1000, -9, -19, -19, 339,
	1480, 1480, 1480, 1480, 230, 339, 339, 1480, 1480, 229,
	40, 40, 8, -1000, -1000, 1480, 228, 89, 89, 89,
	89, 89, 44, 1396, 92, 95, 110, -1000, -1000, 165,
	-1000, -1000, 12, 9, 1597, -1000, 1548, 189, 339, 1438,
	89, 1535, 95, 95, 95, 95, 89, 95, 95, 95,
	95, 89, 81, 89, 212, 1597, 1517, 598, 95, -1000,
	1517, 1354, -1000, 227, -1000, 1299, 210, 89, 89, 89,
	89, -1000, 95, 95, 89, 89, -1000, -1000, 125, 42,
	-1000, 1, 225, 223, -1000, 1257, 181, 1202, 89, -1000,
	854, -1000, -1000, -1000, -1000, 1579, 89, 120, 339, -1000,
	-1000, -1000, -1000, 339, -1000, -1000, -1000, 117, 148, 531,
	-1000, 192, 89, -1000, -1000, 109, -1000, 339, 339, -1000,
	95, -1000, 0, 95, -1000, -1000, 1147, 21, -1000, 1092,
	-1000, -1000, -8, 42, 167, 1480, -1000, -1000, -8, -1000,
	-1000, -1000, -1000, 195, 1480, -1000, 799, 1050, -1000, 177,
	95, 677, 95, 6, -26, -1000, 1480, 339, -1000, 154,
	95, 1480, -1000, -1000, 222, -1000, 995, -1000, -1000, 89,
	995, 953, -1000, 1480, -1000, 89, 1396, -1000, 191, -1000,
	1396, 197, -1000, -1000, -1000, 1579, 89, -1000, 180, 141,
	-1000, 1480, -1000, 88, 1579, 89, 95, -1000, 89, -1000,
	82, -1000, -1000, 1579, 89, -1000, 83, 339, 46, -1000,
	89, 1396, 1396, -1000, 1396, 219, 1480, 1480, 1480, 1480,
	140, -11, -8, -1000, 1480, 1480, 92, -1000, 1396, -1000,
	89, 89, 89, 89, 1480, 339, 89, 89, 1396, 47,
	17, -8, -1000, -8, -1000, -8, -8,
}
var RubyPgo = []int{

	0, 66, 317, 316, 15, 310, 307, 304, 520, 303,
	299, 298, 296, 294, 0, 405, 289, 364, 284, 283,
	282, 19, 279, 4, 214, 278, 5, 274, 273, 270,
	268, 266, 264, 263, 262, 261, 260, 259, 256, 515,
	255, 8, 16, 2, 254, 6, 253, 251, 250, 14,
	3, 246, 1, 245, 244, 7, 10, 243, 9,
}
var RubyR1 = []int{

	0, 44, 44, 44, 44, 44, 44, 44, 44, 44,
	44, 58, 58, 39, 39, 39, 39, 39, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 21, 21, 21, 21, 21, 21, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	41, 41, 52, 52, 49, 49, 49, 49, 55, 55,
	55, 55, 54, 54, 54, 18, 18, 27, 27, 27,
	27, 50, 50, 50, 50, 50, 50, 45, 45, 23,
	23, 23, 23, 56, 56, 56, 22, 22, 25, 26,
	26, 57, 57, 11, 11, 11, 11, 11, 11, 8,
	8, 24, 24, 15, 15, 28, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 2, 5, 6,
	6, 3, 3, 46, 46, 46, 46, 53, 53, 53,
	4, 4, 4, 4, 42, 51, 51, 51, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 43, 43,
	43, 43, 40, 40, 40, 7, 13, 48, 48, 48,
	48, 19, 20, 9, 12, 47, 47, 47, 47, 47,
	47, 16, 16, 16, 16, 16, 16, 16,
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
	3, 3, 1, 3, 3, 7, 7, 0, 1, 3,
	3, 0, 2, 2, 2, 2, 2, 1, 3, 1,
	2, 3, 2, 0, 1, 3, 4, 6, 4, 1,
	3, 1, 3, 3, 3, 3, 3, 3, 3, 2,
	2, 2, 2, 3, 3, 3, 3, 2, 2, 2,
	2, 3, 3, 3, 3, 3, 3, 1, 1, 3,
	3, 5, 5, 0, 4, 7, 8, 3, 7, 8,
	3, 4, 4, 3, 3, 0, 1, 3, 4, 5,
	3, 3, 3, 3, 3, 5, 6, 5, 4, 3,
	3, 2, 0, 2, 2, 3, 4, 0, 3, 4,
	6, 2, 2, 5, 5, 0, 2, 2, 2, 2,
	2, 0, 1, 3, 3, 1, 3, 3,
}
var RubyChk = []int{

	-1000, -44, 44, 45, 61, -1, 44, 45, 61, -14,
	-18, -22, -25, -11, -29, -30, -31, -32, -10, -13,
	-21, -19, -12, 5, 6, 8, -24, -15, -8, -2,
	-5, -6, -3, -26, -9, -17, -7, 13, 19, 20,
	-28, 38, 39, 40, 41, 15, 18, 24, -33, -34,
	-35, -36, -37, -38, 32, 22, 59, 58, 33, 34,
	54, 56, -57, 7, 9, 45, 44, 61, 15, 48,
	35, 36, 4, 40, 41, 42, 50, 51, 49, 18,
	52, 18, 9, -41, -55, 54, 37, 11, -54, -14,
	-4, 6, 8, -24, -15, -8, -17, 12, 56, 9,
	37, 37, 37, 37, 48, 35, 36, 15, 18, 48,
	6, 4, -26, 8, -26, 37, 11, -1, -1, -1,
	-1, -1, -1, -39, -52, -14, -1, 6, 8, 59,
	6, 8, -52, -49, -14, -21, -58, 47, 9, -40,
	-1, 6, -14, -14, -14, -14, -1, -14, -14, -14,
	-14, -1, -14, -1, -49, -14, 11, -14, -14, 6,
	11, -39, -42, 49, -42, -39, -49, -1, -1, -1,
	-1, 6, -14, -14, -1, -1, 6, -45, -56, 9,
	-23, 6, 42, 51, -45, -39, 35, -39, -1, 6,
	-39, 44, 45, 44, 45, -14, -1, -48, 11, 44,
	6, 8, 55, 11, 55, 44, -46, -53, -14, 6,
	8, -49, -1, 45, 10, -55, -41, 9, 46, 10,
	-14, -4, 55, -14, -4, 14, -39, -51, 6, -39,
	57, 10, -58, 11, -56, 37, 6, 6, -58, 14,
	-26, 14, 14, -43, 17, 16, -39, -39, 14, 25,
	-14, -47, -14, -58, -58, 11, 4, 46, 10, -49,
	-14, 37, 14, 49, 11, 57, -50, -23, 10, -1,
	-50, -39, 14, 17, 16, -1, -39, 14, -43, 14,
	-39, 8, 14, 44, 45, -14, -1, -16, 27, 28,
	57, 11, 57, -58, -14, -1, -14, 10, -1, 6,
	-58, 44, 45, -14, -1, -27, -20, 31, -58, 14,
	-1, -39, -39, 14, -39, 4, 15, 18, 15, 18,
	-1, 6, -58, 14, 15, 18, -52, 14, -39, 6,
	-1, -1, -1, -1, 4, 46, -1, -1, -39, -1,
	-14, -58, 11, -58, 11, -58, -58,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 38,
	39, 40, 41, 42, 43, 44, 45, 46, 47, 48,
	49, 50, 51, 18, 19, -2, 21, 22, 23, 24,
	25, 26, 27, 28, 29, 30, 31, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 13, 32, 33,
	34, 35, 36, 37, 0, 0, 0, 0, 137, 138,
	74, 11, 0, 53, 172, 5, 6, 7, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, -2, 56, 62, 74, 0, 0, 71, 78,
	79, 19, -2, 21, 22, 23, 30, 13, -2, 74,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	103, 103, 13, -2, 13, 0, 0, 127, 128, 129,
	130, 13, 0, 177, 181, 72, 0, 121, 122, 0,
	119, 120, 0, 0, 72, 76, 143, 0, 74, 0,
	160, 57, 63, 65, 67, 131, 132, 133, 134, 135,
	136, 162, 0, 164, 0, 75, 0, 72, 113, 125,
	0, 0, 13, 155, 13, 0, 0, 114, 115, 116,
	117, 58, 64, 66, 161, 163, 61, 11, 97, 103,
	104, 99, 0, 0, 11, 0, 0, 0, 118, 126,
	0, 13, 13, 14, 15, 16, 17, 0, 0, 185,
	123, 124, 139, 0, 140, 12, 11, 11, 0, 19,
	-2, 0, 173, 174, 175, 59, 60, -2, 0, 52,
	80, 81, 68, 83, 84, 150, 0, 0, 156, 0,
	153, 55, 91, 0, 0, 0, 100, 102, 91, 106,
	13, 108, 158, 0, 0, 13, 0, 0, 176, 13,
	73, 0, 77, 0, 0, 11, 0, 0, 54, 0,
	183, 0, 151, 154, 0, 152, 11, 105, 98, 101,
	11, 0, 159, 0, 13, 13, 171, 165, 0, 167,
	178, 13, 184, 186, 187, 38, 189, 190, 192, 195,
	141, 0, 142, 0, 38, 11, 147, 70, 69, 157,
	0, 92, 93, 38, 95, 96, 88, 0, 0, 107,
	13, 169, 170, 166, 179, 0, 0, 0, 0, 0,
	0, 0, 144, 85, 0, 0, 182, 86, 168, 13,
	193, 194, 196, 197, 0, 0, 89, 90, 180, 11,
	11, 145, 11, 148, 11, 146, 149,
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
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
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
		//line parser.y:171
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:173
		{
		}
	case 3:
		//line parser.y:175
		{
		}
	case 4:
		//line parser.y:177
		{
		}
	case 5:
		//line parser.y:179
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:181
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:183
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:189
		{
		}
	case 11:
		//line parser.y:191
		{
		}
	case 12:
		//line parser.y:192
		{
		}
	case 13:
		//line parser.y:195
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 14:
		//line parser.y:197
		{
		}
	case 15:
		//line parser.y:199
		{
		}
	case 16:
		//line parser.y:201
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 17:
		//line parser.y:203
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
		//line parser.y:213
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 53:
		//line parser.y:220
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
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
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 56:
		//line parser.y:238
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
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
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
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
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 61:
		//line parser.y:275
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 62:
		//line parser.y:285
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
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
				Func:   ast.BareReference{Name: "<"},
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
				Func:   ast.BareReference{Name: ">"},
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
		//line parser.y:372
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
		//line parser.y:381
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 83:
		//line parser.y:383
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 84:
		//line parser.y:385
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:391
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 86:
		//line parser.y:399
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 87:
		//line parser.y:407
		{
		}
	case 88:
		//line parser.y:409
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 89:
		//line parser.y:411
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{RubyS[Rubypt-2].genericValue}}
		}
	case 90:
		//line parser.y:413
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{ast.Break{}},
			}
		}
	case 91:
		//line parser.y:421
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 92:
		//line parser.y:423
		{
		}
	case 93:
		//line parser.y:425
		{
		}
	case 94:
		//line parser.y:427
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 95:
		//line parser.y:429
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 96:
		//line parser.y:431
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 97:
		//line parser.y:434
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 98:
		//line parser.y:436
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 99:
		//line parser.y:439
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 100:
		//line parser.y:441
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 101:
		//line parser.y:443
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 102:
		//line parser.y:445
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 103:
		//line parser.y:447
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 104:
		//line parser.y:449
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 105:
		//line parser.y:453
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 106:
		//line parser.y:458
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 107:
		//line parser.y:465
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 108:
		//line parser.y:475
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 109:
		//line parser.y:484
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 110:
		//line parser.y:490
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 111:
		//line parser.y:498
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 112:
		//line parser.y:502
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 113:
		//line parser.y:507
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 114:
		//line parser.y:514
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 115:
		//line parser.y:521
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 116:
		//line parser.y:528
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 117:
		//line parser.y:535
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 118:
		//line parser.y:542
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 119:
		//line parser.y:550
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 120:
		//line parser.y:552
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 121:
		//line parser.y:555
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 122:
		//line parser.y:557
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 123:
		//line parser.y:560
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 124:
		//line parser.y:562
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 125:
		//line parser.y:565
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 126:
		//line parser.y:567
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 127:
		//line parser.y:569
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 128:
		//line parser.y:570
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 129:
		//line parser.y:571
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 130:
		//line parser.y:572
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 131:
		//line parser.y:575
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 132:
		//line parser.y:584
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 133:
		//line parser.y:593
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 134:
		//line parser.y:602
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 135:
		//line parser.y:611
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 136:
		//line parser.y:620
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 137:
		//line parser.y:628
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 138:
		//line parser.y:629
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 139:
		//line parser.y:631
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 140:
		//line parser.y:633
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 141:
		//line parser.y:636
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 142:
		//line parser.y:644
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 143:
		//line parser.y:652
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 144:
		//line parser.y:654
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 145:
		//line parser.y:661
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 146:
		//line parser.y:668
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 147:
		//line parser.y:676
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 148:
		//line parser.y:683
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 149:
		//line parser.y:690
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 150:
		//line parser.y:698
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 151:
		//line parser.y:700
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 152:
		//line parser.y:707
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 153:
		//line parser.y:714
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 154:
		//line parser.y:717
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 155:
		//line parser.y:719
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 156:
		//line parser.y:721
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 157:
		//line parser.y:723
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 158:
		//line parser.y:726
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 159:
		//line parser.y:733
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 160:
		//line parser.y:741
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 161:
		//line parser.y:748
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 162:
		//line parser.y:755
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 163:
		//line parser.y:762
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 164:
		//line parser.y:769
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 165:
		//line parser.y:776
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 166:
		//line parser.y:783
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 167:
		//line parser.y:791
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 168:
		//line parser.y:800
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 169:
		//line parser.y:807
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 170:
		//line parser.y:814
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 171:
		//line parser.y:821
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 172:
		//line parser.y:828
		{
		}
	case 173:
		//line parser.y:829
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 174:
		//line parser.y:830
		{
		}
	case 175:
		//line parser.y:833
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 176:
		//line parser.y:836
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 177:
		//line parser.y:844
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 178:
		//line parser.y:846
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 179:
		//line parser.y:848
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 180:
		//line parser.y:857
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
	case 181:
		//line parser.y:872
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 182:
		//line parser.y:880
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 183:
		//line parser.y:889
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 184:
		//line parser.y:898
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 185:
		//line parser.y:901
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 186:
		//line parser.y:903
		{
		}
	case 187:
		//line parser.y:905
		{
		}
	case 188:
		//line parser.y:907
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 189:
		//line parser.y:909
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 190:
		//line parser.y:911
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 191:
		//line parser.y:913
		{
		}
	case 192:
		//line parser.y:915
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 193:
		//line parser.y:917
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 194:
		//line parser.y:919
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 195:
		//line parser.y:921
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 196:
		//line parser.y:923
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 197:
		//line parser.y:925
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	}
	goto Rubystack /* stack new state and value */
}
