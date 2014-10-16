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

//line parser.y:1001

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
	-1, 336,
	27, 185,
	28, 185,
	-2, 13,
	-1, 337,
	27, 185,
	28, 185,
	-2, 13,
}

const RubyNprod = 208
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1940

var RubyAct = []int{

	9, 132, 353, 243, 251, 266, 180, 33, 20, 84,
	133, 90, 178, 83, 177, 162, 2, 3, 294, 190,
	68, 205, 127, 181, 128, 89, 203, 105, 106, 163,
	109, 136, 137, 4, 295, 343, 198, 205, 261, 235,
	104, 103, 81, 102, 264, 101, 112, 114, 186, 66,
	65, 205, 68, 72, 181, 125, 124, 179, 332, 182,
	362, 134, 99, 116, 293, 328, 67, 123, 183, 135,
	204, 142, 143, 144, 145, 129, 147, 148, 149, 150,
	202, 152, 263, 155, 70, 71, 157, 158, 205, 115,
	182, 135, 89, 154, 135, 205, 72, 69, 248, 183,
	155, 80, 68, 326, 300, 203, 172, 173, 135, 249,
	166, 198, 268, 233, 164, 99, 68, 161, 165, 68,
	329, 72, 323, 330, 195, 324, 184, 70, 71, 68,
	156, 335, 185, 196, 187, 5, 255, 208, 218, 155,
	69, 205, 89, 100, 80, 336, 337, 135, 199, 211,
	360, 215, 70, 71, 68, 216, 342, 220, 191, 192,
	321, 223, 195, 322, 233, 69, 195, 68, 221, 80,
	160, 368, 224, 365, 364, 117, 118, 119, 120, 121,
	122, 200, 226, 201, 229, 138, 195, 334, 195, 126,
	281, 195, 234, 363, 240, 365, 364, 258, 203, 250,
	107, 113, 140, 108, 252, 231, 203, 210, 146, 232,
	89, 246, 247, 151, 302, 153, 238, 237, 155, 260,
	105, 106, 256, 219, 203, 236, 135, 195, 259, 316,
	195, 274, 273, 104, 167, 168, 169, 170, 253, 254,
	267, 174, 175, 272, 270, 274, 273, 195, 195, 188,
	278, 228, 285, 70, 71, 189, 176, 297, 299, 130,
	271, 131, 111, 171, 110, 276, 69, 306, 159, 280,
	80, 306, 195, 212, 141, 318, 62, 195, 88, 207,
	227, 195, 197, 206, 1, 139, 53, 296, 52, 51,
	50, 49, 48, 17, 314, 315, 16, 15, 303, 14,
	40, 317, 311, 287, 288, 22, 308, 12, 11, 309,
	21, 125, 331, 10, 19, 195, 195, 13, 195, 18,
	34, 36, 31, 30, 32, 29, 0, 0, 0, 0,
	327, 0, 0, 333, 195, 23, 91, 63, 92, 64,
	347, 348, 349, 0, 351, 0, 0, 195, 285, 285,
	285, 357, 0, 0, 346, 0, 246, 247, 0, 0,
	367, 0, 0, 58, 59, 0, 0, 0, 285, 269,
	372, 373, 0, 285, 285, 285, 374, 0, 275, 0,
	0, 0, 359, 361, 60, 286, 61, 0, 57, 56,
	298, 0, 369, 0, 370, 301, 0, 0, 0, 0,
	307, 0, 0, 0, 307, 72, 0, 313, 0, 0,
	0, 0, 0, 26, 0, 0, 0, 0, 23, 91,
	63, 92, 82, 319, 320, 97, 0, 0, 325, 0,
	0, 0, 0, 0, 0, 0, 70, 71, 93, 0,
	0, 73, 74, 75, 0, 0, 58, 59, 0, 69,
	78, 76, 77, 80, 35, 338, 339, 340, 341, 257,
	0, 0, 0, 344, 345, 0, 0, 85, 93, 98,
	0, 57, 56, 0, 93, 0, 350, 0, 0, 96,
	0, 286, 286, 286, 93, 93, 93, 93, 366, 93,
	93, 93, 93, 0, 93, 0, 93, 0, 371, 93,
	93, 286, 0, 0, 0, 93, 286, 286, 286, 96,
	0, 0, 0, 93, 0, 96, 0, 0, 0, 93,
	93, 0, 0, 0, 0, 96, 96, 96, 96, 0,
	96, 96, 96, 96, 0, 96, 0, 96, 0, 0,
	96, 96, 0, 0, 0, 0, 96, 0, 0, 0,
	93, 0, 93, 0, 96, 93, 0, 0, 0, 0,
	96, 96, 0, 0, 0, 0, 0, 0, 0, 27,
	93, 0, 0, 0, 93, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 72,
	0, 96, 0, 96, 94, 0, 96, 0, 0, 0,
	0, 0, 0, 79, 0, 0, 0, 0, 0, 0,
	0, 96, 93, 28, 0, 96, 0, 93, 0, 0,
	70, 71, 0, 93, 94, 73, 74, 75, 0, 0,
	94, 93, 93, 69, 78, 76, 77, 80, 95, 0,
	94, 94, 94, 94, 0, 94, 94, 94, 94, 0,
	94, 0, 94, 96, 0, 94, 94, 0, 96, 0,
	0, 94, 0, 0, 96, 0, 0, 0, 95, 94,
	0, 93, 96, 96, 95, 94, 94, 0, 0, 0,
	0, 0, 0, 0, 95, 95, 95, 95, 0, 95,
	95, 95, 95, 0, 95, 0, 95, 0, 0, 95,
	95, 0, 0, 0, 0, 95, 94, 0, 94, 0,
	0, 94, 96, 95, 0, 0, 0, 0, 0, 95,
	95, 0, 0, 0, 93, 0, 94, 0, 0, 0,
	94, 23, 91, 63, 92, 82, 0, 87, 97, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	95, 0, 95, 0, 0, 95, 0, 93, 0, 58,
	59, 0, 0, 86, 0, 96, 0, 0, 94, 0,
	95, 0, 0, 94, 95, 0, 0, 0, 0, 94,
	85, 0, 98, 0, 57, 56, 0, 94, 94, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 96, 0,
	23, 91, 63, 92, 64, 0, 0, 97, 0, 0,
	0, 0, 95, 0, 0, 0, 0, 95, 0, 0,
	0, 0, 0, 95, 0, 0, 0, 94, 58, 59,
	0, 95, 95, 23, 24, 63, 25, 64, 0, 0,
	0, 37, 356, 289, 355, 354, 290, 38, 39, 60,
	55, 98, 47, 57, 56, 291, 292, 0, 0, 0,
	54, 58, 59, 0, 0, 0, 41, 42, 43, 44,
	0, 95, 283, 284, 0, 0, 0, 0, 0, 0,
	94, 0, 60, 0, 61, 0, 57, 56, 0, 0,
	0, 0, 23, 24, 63, 25, 64, 0, 0, 0,
	37, 352, 289, 355, 354, 290, 38, 39, 0, 55,
	0, 47, 0, 94, 291, 292, 0, 0, 0, 54,
	58, 59, 0, 0, 95, 41, 42, 43, 44, 0,
	0, 283, 284, 0, 0, 0, 0, 0, 0, 0,
	0, 60, 0, 61, 0, 57, 56, 0, 0, 0,
	23, 24, 63, 25, 64, 0, 0, 95, 37, 358,
	289, 0, 0, 290, 38, 39, 0, 55, 0, 47,
	0, 0, 291, 292, 0, 0, 0, 54, 58, 59,
	0, 0, 0, 41, 42, 43, 44, 0, 0, 283,
	284, 0, 0, 0, 0, 0, 0, 0, 0, 60,
	0, 61, 0, 57, 56, 23, 24, 63, 25, 64,
	0, 0, 0, 37, 282, 289, 0, 0, 290, 38,
	39, 0, 55, 0, 47, 0, 0, 291, 292, 0,
	0, 0, 54, 58, 59, 0, 0, 0, 41, 42,
	43, 44, 0, 0, 283, 284, 0, 0, 0, 0,
	0, 0, 0, 0, 60, 0, 61, 0, 57, 56,
	23, 24, 63, 25, 64, 0, 0, 0, 37, 277,
	45, 245, 244, 46, 38, 39, 0, 55, 0, 47,
	0, 0, 0, 0, 0, 0, 0, 54, 58, 59,
	0, 0, 0, 41, 42, 43, 44, 0, 0, 193,
	194, 0, 0, 0, 0, 0, 0, 0, 0, 60,
	0, 61, 0, 57, 56, 23, 24, 63, 25, 64,
	0, 0, 0, 37, 242, 45, 245, 244, 46, 38,
	39, 0, 55, 0, 47, 0, 0, 0, 0, 0,
	0, 0, 54, 58, 59, 0, 0, 0, 41, 42,
	43, 44, 0, 0, 193, 194, 0, 23, 24, 63,
	25, 64, 0, 0, 60, 37, 61, 45, 57, 56,
	46, 38, 39, 0, 55, 0, 47, 0, 0, 0,
	0, 0, 0, 0, 54, 58, 59, 0, 0, 0,
	41, 42, 43, 44, 0, 0, 6, 7, 0, 0,
	0, 0, 0, 0, 0, 0, 60, 0, 61, 0,
	57, 56, 0, 8, 23, 24, 63, 25, 64, 0,
	0, 0, 37, 0, 289, 0, 0, 290, 38, 39,
	0, 55, 0, 47, 0, 0, 291, 292, 0, 0,
	0, 54, 58, 59, 0, 0, 0, 41, 42, 43,
	44, 0, 0, 283, 284, 0, 0, 0, 0, 0,
	0, 0, 0, 60, 0, 61, 0, 57, 56, 23,
	24, 63, 25, 64, 0, 0, 0, 37, 312, 45,
	0, 0, 46, 38, 39, 0, 55, 0, 47, 0,
	0, 0, 0, 0, 0, 0, 54, 58, 59, 0,
	0, 0, 41, 42, 43, 44, 0, 0, 193, 194,
	0, 23, 24, 63, 25, 64, 0, 0, 60, 37,
	61, 45, 57, 56, 46, 38, 39, 0, 55, 0,
	47, 0, 0, 0, 0, 0, 0, 310, 54, 58,
	59, 0, 0, 0, 41, 42, 43, 44, 0, 0,
	304, 305, 0, 0, 0, 0, 0, 0, 0, 0,
	60, 0, 61, 0, 57, 56, 23, 24, 63, 25,
	64, 0, 0, 0, 37, 279, 45, 0, 0, 46,
	38, 39, 0, 55, 0, 47, 0, 0, 0, 0,
	0, 0, 0, 54, 58, 59, 0, 0, 0, 41,
	42, 43, 44, 0, 0, 193, 194, 0, 23, 24,
	63, 25, 64, 0, 0, 60, 37, 61, 45, 57,
	56, 46, 38, 39, 0, 55, 0, 47, 0, 0,
	0, 0, 0, 0, 0, 54, 58, 59, 0, 0,
	0, 41, 42, 43, 44, 0, 0, 193, 194, 0,
	0, 0, 0, 0, 0, 0, 0, 60, 0, 61,
	265, 57, 56, 23, 24, 63, 25, 64, 0, 0,
	0, 37, 262, 45, 0, 0, 46, 38, 39, 0,
	55, 0, 47, 0, 0, 0, 0, 0, 0, 0,
	54, 58, 59, 0, 0, 0, 41, 42, 43, 44,
	0, 0, 193, 194, 0, 0, 0, 0, 0, 0,
	0, 0, 60, 0, 61, 0, 57, 56, 23, 24,
	63, 25, 64, 0, 0, 0, 37, 241, 45, 0,
	0, 46, 38, 39, 0, 55, 0, 47, 0, 0,
	0, 0, 0, 0, 0, 54, 58, 59, 0, 0,
	0, 41, 42, 43, 44, 0, 0, 193, 194, 0,
	0, 0, 0, 0, 0, 0, 0, 60, 0, 61,
	0, 57, 56, 23, 24, 63, 25, 64, 0, 0,
	0, 37, 239, 45, 0, 0, 46, 38, 39, 0,
	55, 0, 47, 0, 0, 0, 0, 0, 0, 0,
	54, 58, 59, 0, 0, 0, 41, 42, 43, 44,
	0, 0, 193, 194, 0, 23, 24, 63, 25, 64,
	0, 0, 60, 37, 61, 45, 57, 56, 46, 38,
	39, 0, 55, 0, 47, 0, 0, 0, 0, 0,
	0, 0, 54, 58, 59, 0, 0, 0, 41, 42,
	43, 44, 0, 0, 193, 194, 0, 0, 0, 0,
	0, 0, 0, 0, 60, 0, 61, 230, 57, 56,
	23, 24, 63, 25, 64, 0, 0, 0, 37, 225,
	45, 0, 0, 46, 38, 39, 0, 55, 0, 47,
	0, 0, 0, 0, 0, 0, 0, 54, 58, 59,
	0, 0, 0, 41, 42, 43, 44, 0, 0, 193,
	194, 0, 23, 24, 63, 25, 64, 0, 0, 60,
	37, 61, 45, 57, 56, 46, 38, 39, 0, 55,
	0, 47, 0, 0, 0, 0, 0, 0, 0, 54,
	58, 59, 0, 0, 0, 41, 42, 43, 44, 0,
	0, 193, 194, 0, 23, 24, 63, 25, 64, 214,
	0, 60, 37, 61, 45, 57, 56, 46, 38, 39,
	0, 55, 0, 47, 0, 0, 0, 0, 0, 0,
	0, 54, 58, 59, 0, 0, 0, 41, 42, 43,
	44, 0, 0, 0, 213, 0, 23, 24, 63, 25,
	64, 0, 0, 60, 37, 61, 45, 57, 56, 46,
	38, 39, 0, 55, 0, 47, 0, 0, 0, 0,
	0, 0, 0, 54, 58, 59, 0, 0, 0, 41,
	42, 43, 44, 23, 91, 63, 92, 217, 0, 0,
	97, 0, 0, 0, 0, 60, 0, 61, 0, 57,
	56, 23, 209, 63, 92, 64, 0, 0, 0, 0,
	0, 58, 59, 23, 91, 63, 92, 82, 0, 0,
	97, 0, 0, 0, 0, 0, 0, 0, 0, 58,
	59, 0, 60, 0, 98, 0, 57, 56, 72, 0,
	205, 58, 59, 0, 0, 0, 0, 0, 0, 0,
	60, 0, 61, 0, 57, 56, 0, 0, 0, 0,
	0, 0, 85, 0, 98, 0, 57, 56, 0, 70,
	71, 0, 0, 0, 73, 74, 75, 0, 0, 0,
	0, 0, 69, 78, 76, 77, 80, 0, 0, 222,
}
var RubyPact = []int{

	-28, 1152, -1000, -1000, -1000, 5, -1000, -1000, -1000, 585,
	-1000, -1000, -1000, 24, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 726, 106, 8, 6, 4, -1000,
	-1000, -1000, -1000, -1000, -1000, 185, -18, 258, 193, 193,
	52, 1791, 1791, 1791, 1791, 1791, 1791, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 330, 1791, 16, 253, -1000, -1000,
	330, -1000, -15, 176, -1000, -1000, -1000, -1000, 1791, 268,
	330, 330, 330, 330, 1791, 330, 330, 330, 330, 1791,
	330, 1791, 330, -1000, 119, 330, 330, 262, 159, 117,
	-1000, 1858, 53, -1000, -1000, -1000, -8, -20, -20, 330,
	1791, 1791, 1791, 1791, 257, 330, 330, 1791, 1791, 250,
	48, 48, 13, -1000, -1000, 1791, 249, 37, 37, 37,
	37, 37, 114, 1707, 100, 117, 104, -1000, -1000, 175,
	-1000, -1000, 25, 15, 401, -1000, 1846, 199, 330, 1749,
	37, 1828, 117, 117, 117, 117, 37, 117, 117, 117,
	117, 37, 92, 37, 213, 401, 795, 1884, 117, -1000,
	795, 1665, -1000, 245, -1000, 1610, 195, 37, 37, 37,
	37, -1000, 117, 117, 37, 37, -1000, -1000, 153, 17,
	-1000, 2, 219, 211, -1000, 1568, 193, 1513, 37, -1000,
	1110, -1000, -1000, -1000, -1000, 585, 37, 84, 330, -1000,
	-1000, -1000, -1000, 330, -1000, -1000, -1000, 125, 218, 413,
	-1000, 187, 37, -1000, -1000, 119, -1000, 330, 330, -1000,
	117, -1000, 1, 117, -1000, -1000, 1458, 33, -1000, 1403,
	-1000, -1000, -7, 17, 102, 1791, -1000, -1000, -7, -1000,
	-1000, -1000, -1000, 229, 1791, -1000, 1055, 1361, -1000, 182,
	117, 1000, 117, 7, -23, -1000, 1791, 330, -1000, 94,
	117, 1791, -1000, -1000, 208, -1000, 1306, -1000, -1000, 37,
	1306, 1264, -1000, 1791, -1000, 37, 1707, -1000, 215, -1000,
	1707, 271, -1000, -1000, -1000, 585, 37, -1000, -1000, 1791,
	1791, 145, 107, -1000, 1791, -1000, 97, 585, 37, 117,
	-1000, 37, -1000, 51, -1000, -1000, 585, 37, -1000, 105,
	330, 44, -1000, 37, 1707, 1707, -1000, 1707, 181, 87,
	101, 1791, 1791, 1791, 1791, 152, -11, -7, -1000, 1791,
	1791, 100, -1000, 1707, -1000, -1000, -1000, -1000, 37, 37,
	37, 37, 1791, 330, 37, 37, 1707, 887, 828, 945,
	139, 49, -1000, 179, 1791, -1000, -1000, 157, -1000, -7,
	-1000, -7, -1000, -1000, 1791, -1000, 37, 1209, -1000, -7,
	-7, 37, 1209, 1209, 1209,
}
var RubyPgo = []int{

	0, 133, 325, 324, 11, 323, 322, 321, 613, 320,
	319, 317, 314, 0, 569, 454, 313, 310, 309, 8,
	308, 6, 413, 307, 7, 306, 305, 304, 303, 2,
	300, 299, 297, 296, 293, 292, 291, 290, 289, 288,
	286, 19, 285, 13, 15, 3, 284, 14, 283, 4,
	282, 10, 5, 280, 1, 279, 278, 9, 12, 276,
	31,
}
var RubyR1 = []int{

	0, 46, 46, 46, 46, 46, 46, 46, 46, 46,
	46, 60, 60, 41, 41, 41, 41, 41, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 19, 19, 19, 19, 19, 19, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	43, 43, 54, 54, 51, 51, 51, 51, 57, 57,
	57, 57, 56, 56, 56, 16, 16, 25, 25, 25,
	25, 52, 52, 52, 52, 52, 52, 47, 47, 21,
	21, 21, 21, 58, 58, 58, 20, 20, 23, 24,
	24, 59, 59, 11, 11, 11, 11, 11, 11, 8,
	8, 22, 22, 14, 14, 30, 30, 31, 32, 33,
	34, 35, 36, 37, 38, 39, 40, 2, 5, 6,
	6, 3, 3, 48, 48, 48, 48, 55, 55, 55,
	4, 4, 4, 4, 44, 53, 53, 53, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 45, 45,
	45, 45, 42, 42, 42, 7, 12, 50, 50, 50,
	50, 17, 18, 9, 26, 49, 49, 49, 49, 49,
	49, 49, 27, 27, 27, 27, 27, 27, 27, 28,
	28, 28, 28, 28, 29, 29, 29, 29,
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
	2, 2, 0, 1, 3, 3, 1, 3, 3, 5,
	6, 5, 6, 5, 4, 3, 3, 2,
}
var RubyChk = []int{

	-1000, -46, 44, 45, 61, -1, 44, 45, 61, -13,
	-16, -20, -23, -11, -31, -32, -33, -34, -10, -12,
	-19, -17, -26, 5, 6, 8, -22, -14, -8, -2,
	-5, -6, -3, -24, -9, -15, -7, 13, 19, 20,
	-30, 38, 39, 40, 41, 15, 18, 24, -35, -36,
	-37, -38, -39, -40, 32, 22, 59, 58, 33, 34,
	54, 56, -59, 7, 9, 45, 44, 61, 15, 48,
	35, 36, 4, 40, 41, 42, 50, 51, 49, 18,
	52, 18, 9, -43, -57, 54, 37, 11, -56, -13,
	-4, 6, 8, -22, -14, -8, -15, 12, 56, 9,
	37, 37, 37, 37, 48, 35, 36, 15, 18, 48,
	6, 4, -24, 8, -24, 37, 11, -1, -1, -1,
	-1, -1, -1, -41, -54, -13, -1, 6, 8, 59,
	6, 8, -54, -51, -13, -19, -60, 47, 9, -42,
	-1, 6, -13, -13, -13, -13, -1, -13, -13, -13,
	-13, -1, -13, -1, -51, -13, 11, -13, -13, 6,
	11, -41, -44, 49, -44, -41, -51, -1, -1, -1,
	-1, 6, -13, -13, -1, -1, 6, -47, -58, 9,
	-21, 6, 42, 51, -47, -41, 35, -41, -1, 6,
	-41, 44, 45, 44, 45, -13, -1, -50, 11, 44,
	6, 8, 55, 11, 55, 44, -48, -55, -13, 6,
	8, -51, -1, 45, 10, -57, -43, 9, 46, 10,
	-13, -4, 55, -13, -4, 14, -41, -53, 6, -41,
	57, 10, -60, 11, -58, 37, 6, 6, -60, 14,
	-24, 14, 14, -45, 17, 16, -41, -41, 14, 25,
	-13, -49, -13, -60, -60, 11, 4, 46, 10, -51,
	-13, 37, 14, 49, 11, 57, -52, -21, 10, -1,
	-52, -41, 14, 17, 16, -1, -41, 14, -45, 14,
	-41, 8, 14, 44, 45, -13, -1, -28, -27, 15,
	18, 27, 28, 57, 11, 57, -60, -13, -1, -13,
	10, -1, 6, -60, 44, 45, -13, -1, -25, -18,
	31, -60, 14, -1, -41, -41, 14, -41, 4, -1,
	-1, 15, 18, 15, 18, -1, 6, -60, 14, 15,
	18, -54, 14, -41, 6, 44, 44, 45, -1, -1,
	-1, -1, 4, 46, -1, -1, -41, -49, -49, -49,
	-1, -13, 14, -29, 17, 16, 14, -29, 14, -60,
	11, -60, 11, 14, 17, 16, -1, -49, 14, -60,
	-60, -1, -49, -49, -49,
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
	178, 13, 184, 186, 187, 38, 189, 190, 191, 0,
	0, 193, 196, 141, 0, 142, 0, 38, 11, 147,
	70, 69, 157, 0, 92, 93, 38, 95, 96, 88,
	0, 0, 107, 13, 169, 170, 166, 179, 0, 13,
	0, 0, 0, 0, 0, 0, 0, 144, 85, 0,
	0, 182, 86, 168, 13, 185, -2, -2, 194, 195,
	197, 198, 0, 0, 89, 90, 180, 0, 0, 0,
	11, 11, 199, 0, 0, 185, 201, 0, 203, 145,
	11, 148, 11, 200, 0, 185, 185, 192, 202, 146,
	149, 185, 192, 192, 192,
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
		//line parser.y:175
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:177
		{
		}
	case 3:
		//line parser.y:179
		{
		}
	case 4:
		//line parser.y:181
		{
		}
	case 5:
		//line parser.y:183
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:185
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:187
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:193
		{
		}
	case 11:
		//line parser.y:195
		{
		}
	case 12:
		//line parser.y:196
		{
		}
	case 13:
		//line parser.y:199
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 14:
		//line parser.y:201
		{
		}
	case 15:
		//line parser.y:203
		{
		}
	case 16:
		//line parser.y:205
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 17:
		//line parser.y:207
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
		//line parser.y:217
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 53:
		//line parser.y:224
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 54:
		//line parser.y:228
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 55:
		//line parser.y:235
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 56:
		//line parser.y:242
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 57:
		//line parser.y:249
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 58:
		//line parser.y:256
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 59:
		//line parser.y:263
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 60:
		//line parser.y:271
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 61:
		//line parser.y:279
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 62:
		//line parser.y:289
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 63:
		//line parser.y:296
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 64:
		//line parser.y:304
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 65:
		//line parser.y:312
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 66:
		//line parser.y:320
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 67:
		//line parser.y:328
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 68:
		//line parser.y:338
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 69:
		//line parser.y:348
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 70:
		//line parser.y:357
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 71:
		//line parser.y:359
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 72:
		//line parser.y:362
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:364
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:366
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 75:
		//line parser.y:368
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:370
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:372
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 78:
		//line parser.y:376
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 79:
		//line parser.y:378
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 80:
		//line parser.y:380
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 81:
		//line parser.y:382
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 82:
		//line parser.y:385
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 83:
		//line parser.y:387
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 84:
		//line parser.y:389
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:395
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 86:
		//line parser.y:403
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 87:
		//line parser.y:411
		{
		}
	case 88:
		//line parser.y:413
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 89:
		//line parser.y:415
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{RubyS[Rubypt-2].genericValue}}
		}
	case 90:
		//line parser.y:417
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{ast.Break{}},
			}
		}
	case 91:
		//line parser.y:425
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 92:
		//line parser.y:427
		{
		}
	case 93:
		//line parser.y:429
		{
		}
	case 94:
		//line parser.y:431
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 95:
		//line parser.y:433
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 96:
		//line parser.y:435
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 97:
		//line parser.y:438
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 98:
		//line parser.y:440
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 99:
		//line parser.y:443
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 100:
		//line parser.y:445
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 101:
		//line parser.y:447
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 102:
		//line parser.y:449
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 103:
		//line parser.y:451
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 104:
		//line parser.y:453
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 105:
		//line parser.y:457
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 106:
		//line parser.y:462
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 107:
		//line parser.y:469
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 108:
		//line parser.y:479
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 109:
		//line parser.y:488
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 110:
		//line parser.y:494
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 111:
		//line parser.y:502
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 112:
		//line parser.y:506
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 113:
		//line parser.y:511
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 114:
		//line parser.y:518
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 115:
		//line parser.y:525
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 116:
		//line parser.y:532
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 117:
		//line parser.y:539
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 118:
		//line parser.y:546
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 119:
		//line parser.y:554
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 120:
		//line parser.y:556
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 121:
		//line parser.y:559
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 122:
		//line parser.y:561
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 123:
		//line parser.y:564
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 124:
		//line parser.y:566
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 125:
		//line parser.y:569
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 126:
		//line parser.y:571
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 127:
		//line parser.y:573
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 128:
		//line parser.y:574
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 129:
		//line parser.y:575
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 130:
		//line parser.y:576
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 131:
		//line parser.y:579
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 132:
		//line parser.y:588
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 133:
		//line parser.y:597
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 134:
		//line parser.y:606
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 135:
		//line parser.y:615
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 136:
		//line parser.y:624
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 137:
		//line parser.y:632
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 138:
		//line parser.y:633
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 139:
		//line parser.y:635
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 140:
		//line parser.y:637
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 141:
		//line parser.y:640
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 142:
		//line parser.y:648
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 143:
		//line parser.y:656
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 144:
		//line parser.y:658
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 145:
		//line parser.y:665
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 146:
		//line parser.y:672
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 147:
		//line parser.y:680
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 148:
		//line parser.y:687
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 149:
		//line parser.y:694
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 150:
		//line parser.y:702
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 151:
		//line parser.y:704
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 152:
		//line parser.y:711
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 153:
		//line parser.y:718
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 154:
		//line parser.y:721
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 155:
		//line parser.y:723
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 156:
		//line parser.y:725
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 157:
		//line parser.y:727
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 158:
		//line parser.y:730
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 159:
		//line parser.y:737
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 160:
		//line parser.y:745
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 161:
		//line parser.y:752
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 162:
		//line parser.y:759
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 163:
		//line parser.y:766
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 164:
		//line parser.y:773
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 165:
		//line parser.y:780
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 166:
		//line parser.y:787
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 167:
		//line parser.y:795
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 168:
		//line parser.y:804
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 169:
		//line parser.y:811
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 170:
		//line parser.y:818
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 171:
		//line parser.y:825
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 172:
		//line parser.y:832
		{
		}
	case 173:
		//line parser.y:833
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 174:
		//line parser.y:834
		{
		}
	case 175:
		//line parser.y:837
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 176:
		//line parser.y:840
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 177:
		//line parser.y:848
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 178:
		//line parser.y:850
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 179:
		//line parser.y:852
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 180:
		//line parser.y:861
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
		//line parser.y:876
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 182:
		//line parser.y:885
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 183:
		//line parser.y:894
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 184:
		//line parser.y:903
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 185:
		//line parser.y:906
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 186:
		//line parser.y:908
		{
		}
	case 187:
		//line parser.y:910
		{
		}
	case 188:
		//line parser.y:912
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 189:
		//line parser.y:914
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 190:
		//line parser.y:916
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 191:
		//line parser.y:918
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 192:
		//line parser.y:920
		{
		}
	case 193:
		//line parser.y:922
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 194:
		//line parser.y:924
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 195:
		//line parser.y:926
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 196:
		//line parser.y:928
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 197:
		//line parser.y:930
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 198:
		//line parser.y:932
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 199:
		//line parser.y:935
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 200:
		//line parser.y:942
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 201:
		//line parser.y:950
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 202:
		//line parser.y:957
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 203:
		//line parser.y:965
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 204:
		//line parser.y:973
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 205:
		//line parser.y:980
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 206:
		//line parser.y:987
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 207:
		//line parser.y:994
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	}
	goto Rubystack /* stack new state and value */
}
