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

//line parser.y:900

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 26,
	47, 102,
	-2, 20,
	-1, 84,
	10, 75,
	11, 75,
	-2, 163,
	-1, 94,
	47, 102,
	-2, 20,
	-1, 100,
	13, 13,
	15, 13,
	18, 13,
	19, 13,
	20, 13,
	22, 13,
	24, 13,
	31, 13,
	32, 13,
	38, 13,
	39, 13,
	40, 13,
	41, 13,
	45, 13,
	-2, 11,
	-1, 115,
	47, 102,
	-2, 100,
	-1, 213,
	47, 103,
	-2, 101,
	-1, 220,
	10, 75,
	11, 75,
	-2, 163,
}

const RubyNprod = 189
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1602

var RubyAct = []int{

	9, 183, 246, 34, 181, 20, 92, 180, 86, 85,
	136, 206, 165, 24, 93, 65, 94, 84, 166, 89,
	99, 2, 3, 208, 70, 294, 91, 267, 111, 130,
	135, 131, 140, 201, 328, 208, 295, 189, 4, 107,
	108, 60, 61, 114, 116, 88, 317, 139, 264, 238,
	105, 83, 106, 68, 67, 207, 127, 127, 208, 70,
	159, 184, 87, 137, 100, 266, 59, 58, 138, 104,
	69, 293, 74, 145, 146, 147, 148, 205, 150, 151,
	152, 153, 132, 155, 208, 158, 126, 128, 160, 161,
	138, 74, 258, 138, 91, 157, 184, 185, 335, 182,
	103, 101, 158, 72, 73, 101, 186, 138, 175, 176,
	118, 333, 169, 167, 221, 70, 71, 70, 320, 74,
	82, 187, 72, 73, 300, 206, 198, 24, 93, 65,
	94, 84, 185, 102, 99, 71, 117, 236, 251, 82,
	211, 186, 158, 70, 319, 91, 202, 138, 208, 252,
	72, 73, 214, 218, 219, 60, 61, 201, 314, 312,
	223, 315, 313, 71, 226, 198, 224, 82, 260, 198,
	227, 163, 194, 195, 208, 327, 87, 284, 100, 141,
	59, 58, 309, 115, 277, 276, 70, 237, 259, 198,
	213, 198, 322, 243, 198, 24, 212, 65, 94, 66,
	302, 275, 253, 277, 276, 271, 236, 255, 261, 206,
	234, 206, 240, 91, 222, 206, 27, 239, 231, 72,
	73, 158, 263, 60, 61, 203, 138, 204, 235, 192,
	198, 262, 71, 198, 208, 241, 82, 133, 270, 134,
	179, 113, 95, 112, 62, 174, 63, 162, 59, 58,
	198, 198, 281, 144, 64, 288, 311, 256, 257, 90,
	297, 299, 210, 230, 200, 109, 254, 209, 110, 1,
	198, 142, 95, 95, 198, 198, 54, 53, 52, 95,
	198, 51, 50, 49, 198, 107, 108, 17, 16, 95,
	95, 95, 95, 15, 95, 95, 95, 95, 106, 95,
	14, 95, 41, 74, 95, 95, 296, 12, 198, 198,
	95, 198, 36, 11, 22, 21, 10, 303, 95, 290,
	19, 304, 198, 23, 95, 95, 13, 18, 35, 331,
	198, 37, 32, 31, 72, 73, 33, 30, 98, 75,
	76, 77, 0, 0, 0, 0, 318, 71, 80, 78,
	79, 82, 0, 0, 225, 0, 95, 0, 95, 0,
	0, 95, 0, 0, 0, 0, 0, 0, 98, 98,
	0, 0, 0, 0, 0, 98, 95, 0, 332, 334,
	95, 336, 0, 337, 0, 98, 98, 98, 98, 0,
	98, 98, 98, 98, 0, 98, 0, 98, 0, 0,
	98, 98, 0, 0, 0, 0, 98, 28, 0, 0,
	0, 0, 0, 0, 98, 0, 0, 0, 95, 0,
	98, 98, 0, 95, 0, 0, 0, 0, 0, 95,
	0, 0, 0, 96, 0, 0, 0, 95, 95, 0,
	24, 93, 65, 94, 66, 0, 0, 99, 0, 0,
	0, 0, 98, 0, 98, 0, 0, 98, 0, 0,
	0, 0, 0, 96, 96, 0, 0, 0, 60, 61,
	96, 0, 98, 0, 0, 0, 98, 95, 0, 0,
	96, 96, 96, 96, 0, 96, 96, 96, 96, 62,
	96, 100, 96, 59, 58, 96, 96, 0, 0, 0,
	0, 96, 29, 0, 0, 0, 125, 0, 0, 96,
	0, 0, 0, 0, 98, 96, 96, 0, 0, 98,
	0, 0, 0, 0, 0, 98, 0, 0, 97, 0,
	0, 0, 0, 98, 98, 24, 93, 65, 94, 220,
	0, 0, 99, 0, 0, 95, 0, 96, 0, 96,
	0, 0, 96, 0, 0, 0, 0, 0, 97, 97,
	0, 0, 0, 60, 61, 97, 0, 96, 0, 0,
	0, 96, 0, 98, 0, 97, 97, 97, 97, 0,
	97, 97, 97, 97, 62, 97, 100, 97, 59, 58,
	97, 97, 0, 0, 0, 0, 97, 24, 93, 65,
	94, 66, 0, 0, 97, 0, 164, 168, 0, 96,
	97, 97, 0, 0, 96, 0, 0, 0, 0, 0,
	96, 188, 0, 190, 0, 60, 61, 0, 96, 96,
	193, 0, 0, 0, 199, 0, 5, 0, 0, 0,
	0, 98, 97, 0, 97, 0, 62, 97, 63, 0,
	59, 58, 74, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 97, 0, 0, 0, 97, 0, 96, 0,
	0, 0, 229, 0, 232, 0, 0, 119, 120, 121,
	122, 123, 124, 72, 73, 0, 0, 0, 75, 76,
	77, 0, 129, 0, 0, 0, 71, 80, 78, 79,
	82, 249, 250, 0, 97, 143, 0, 0, 0, 97,
	0, 149, 0, 0, 0, 97, 154, 0, 156, 0,
	0, 0, 0, 97, 97, 24, 93, 65, 94, 84,
	0, 0, 99, 0, 0, 0, 96, 170, 171, 172,
	173, 0, 269, 0, 177, 178, 0, 0, 273, 0,
	274, 74, 191, 60, 61, 279, 0, 0, 0, 283,
	0, 0, 0, 97, 0, 81, 0, 0, 0, 0,
	0, 0, 0, 0, 87, 0, 100, 215, 59, 58,
	0, 0, 72, 73, 307, 308, 0, 75, 76, 77,
	0, 310, 0, 0, 0, 71, 80, 78, 79, 82,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 321, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 329,
	0, 97, 24, 25, 65, 26, 66, 0, 0, 0,
	38, 285, 46, 0, 0, 47, 39, 40, 0, 57,
	0, 48, 0, 0, 291, 292, 0, 0, 56, 55,
	60, 61, 0, 0, 0, 42, 43, 44, 45, 0,
	0, 286, 287, 272, 0, 0, 0, 0, 0, 0,
	0, 62, 278, 63, 0, 59, 58, 0, 0, 289,
	0, 0, 0, 0, 298, 0, 0, 0, 0, 301,
	0, 0, 0, 0, 24, 25, 65, 26, 66, 0,
	0, 306, 38, 280, 46, 248, 247, 47, 39, 40,
	0, 57, 0, 48, 0, 0, 0, 0, 0, 316,
	56, 55, 60, 61, 0, 0, 0, 42, 43, 44,
	45, 0, 0, 196, 197, 0, 0, 323, 324, 325,
	326, 0, 0, 62, 0, 63, 0, 59, 58, 0,
	0, 0, 330, 24, 25, 65, 26, 66, 0, 0,
	0, 38, 245, 46, 248, 247, 47, 39, 40, 0,
	57, 0, 48, 0, 0, 0, 0, 0, 0, 56,
	55, 60, 61, 0, 0, 0, 42, 43, 44, 45,
	0, 0, 196, 197, 0, 24, 25, 65, 26, 66,
	0, 0, 62, 38, 63, 46, 59, 58, 47, 39,
	40, 0, 57, 0, 48, 0, 0, 0, 0, 0,
	0, 56, 55, 60, 61, 0, 0, 0, 42, 43,
	44, 45, 0, 0, 6, 7, 0, 0, 0, 0,
	0, 0, 0, 0, 62, 0, 63, 0, 59, 58,
	0, 8, 24, 25, 65, 26, 66, 0, 0, 0,
	38, 305, 46, 0, 0, 47, 39, 40, 0, 57,
	0, 48, 0, 0, 0, 0, 0, 0, 56, 55,
	60, 61, 0, 0, 0, 42, 43, 44, 45, 0,
	0, 196, 197, 0, 0, 0, 0, 0, 0, 0,
	0, 62, 0, 63, 0, 59, 58, 24, 25, 65,
	26, 66, 0, 0, 0, 38, 282, 46, 0, 0,
	47, 39, 40, 0, 57, 0, 48, 0, 0, 0,
	0, 0, 0, 56, 55, 60, 61, 0, 0, 0,
	42, 43, 44, 45, 0, 0, 196, 197, 0, 24,
	25, 65, 26, 66, 0, 0, 62, 38, 63, 46,
	59, 58, 47, 39, 40, 0, 57, 0, 48, 0,
	0, 0, 0, 0, 0, 56, 55, 60, 61, 0,
	0, 0, 42, 43, 44, 45, 0, 0, 196, 197,
	0, 0, 0, 0, 0, 0, 0, 0, 62, 0,
	63, 268, 59, 58, 24, 25, 65, 26, 66, 0,
	0, 0, 38, 265, 46, 0, 0, 47, 39, 40,
	0, 57, 0, 48, 0, 0, 0, 0, 0, 0,
	56, 55, 60, 61, 0, 0, 0, 42, 43, 44,
	45, 0, 0, 196, 197, 0, 0, 0, 0, 0,
	0, 0, 0, 62, 0, 63, 0, 59, 58, 24,
	25, 65, 26, 66, 0, 0, 0, 38, 244, 46,
	0, 0, 47, 39, 40, 0, 57, 0, 48, 0,
	0, 0, 0, 0, 0, 56, 55, 60, 61, 0,
	0, 0, 42, 43, 44, 45, 0, 0, 196, 197,
	0, 0, 0, 0, 0, 0, 0, 0, 62, 0,
	63, 0, 59, 58, 24, 25, 65, 26, 66, 0,
	0, 0, 38, 242, 46, 0, 0, 47, 39, 40,
	0, 57, 0, 48, 0, 0, 0, 0, 0, 0,
	56, 55, 60, 61, 0, 0, 0, 42, 43, 44,
	45, 0, 0, 196, 197, 0, 24, 25, 65, 26,
	66, 0, 0, 62, 38, 63, 46, 59, 58, 47,
	39, 40, 0, 57, 0, 48, 0, 0, 0, 0,
	0, 0, 56, 55, 60, 61, 0, 0, 0, 42,
	43, 44, 45, 0, 0, 196, 197, 0, 0, 0,
	0, 0, 0, 0, 0, 62, 0, 63, 233, 59,
	58, 24, 25, 65, 26, 66, 0, 0, 0, 38,
	228, 46, 0, 0, 47, 39, 40, 0, 57, 0,
	48, 0, 0, 0, 0, 0, 0, 56, 55, 60,
	61, 0, 0, 0, 42, 43, 44, 45, 0, 0,
	196, 197, 0, 24, 25, 65, 26, 66, 0, 0,
	62, 38, 63, 46, 59, 58, 47, 39, 40, 0,
	57, 0, 48, 0, 0, 0, 0, 0, 0, 56,
	55, 60, 61, 0, 0, 0, 42, 43, 44, 45,
	0, 0, 196, 197, 0, 24, 25, 65, 26, 66,
	217, 0, 62, 38, 63, 46, 59, 58, 47, 39,
	40, 0, 57, 0, 48, 0, 0, 0, 0, 0,
	0, 56, 55, 60, 61, 0, 0, 0, 42, 43,
	44, 45, 0, 0, 0, 216, 0, 24, 25, 65,
	26, 66, 0, 0, 62, 38, 63, 46, 59, 58,
	47, 39, 40, 0, 57, 0, 48, 0, 0, 0,
	0, 0, 0, 56, 55, 60, 61, 0, 0, 0,
	42, 43, 44, 45, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 62, 0, 63, 0,
	59, 58,
}
var RubyPact = []int{

	-23, 1000, -1000, -1000, -1000, 9, -1000, -1000, -1000, 747,
	-1000, -1000, -1000, 33, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 8, 96, 63, 32, 13,
	-1000, -1000, -1000, -1000, -1000, -1000, 250, -20, 237, 175,
	175, 99, 1542, 1542, 1542, 1542, 1542, 1542, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 592, 592, 1542, 23, 231,
	-1000, -1000, 592, -1000, -15, 170, -1000, -1000, -1000, -1000,
	1542, 247, 592, 592, 592, 592, 1542, 592, 592, 592,
	592, 1542, 592, 1542, 592, -1000, 49, 592, 592, 241,
	160, 115, -1000, 720, 92, -1000, -1000, -1000, 4, -31,
	-31, 592, 1542, 1542, 1542, 1542, 239, 592, 592, 1542,
	1542, 234, 90, 90, 2, -1000, -1000, 1542, 223, 44,
	44, 44, 44, 44, 128, 1458, 146, 115, 146, 102,
	-1000, -1000, 219, -1000, -1000, 22, 0, 648, -1000, 190,
	182, 592, 1500, 44, 530, 115, 115, 115, 115, 44,
	115, 115, 115, 115, 44, 68, 44, 204, 648, 435,
	299, 115, -1000, 435, 1416, -1000, 212, -1000, 1361, 200,
	44, 44, 44, 44, -1000, 115, 115, 44, 44, -1000,
	-1000, 126, 55, -1000, 12, 211, 206, -1000, 1319, 175,
	1264, 44, -1000, 958, -1000, -1000, -1000, -1000, 747, 44,
	124, 592, -1000, -1000, -1000, -1000, 592, -1000, -1000, -1000,
	81, 184, 122, -1000, 198, 44, -1000, -1000, 49, -1000,
	592, 592, -1000, 115, -1000, 11, 115, -1000, -1000, 1209,
	16, -1000, 1154, -1000, -1000, -9, 55, 195, 1542, -1000,
	-1000, -9, -1000, -1000, -1000, -1000, 187, 1542, -1000, 899,
	1112, -1000, 169, 115, 827, 115, 14, -21, -1000, 1542,
	592, -1000, 114, 115, 1542, -1000, -1000, 194, -1000, 1458,
	-1000, -1000, 44, 1458, 1057, -1000, 1542, -1000, 44, 1458,
	-1000, 168, -1000, 1458, 252, -1000, -1000, -1000, 747, 44,
	-1000, 144, 143, -1000, 1542, -1000, 40, 747, 44, 115,
	-1000, 44, -1000, 130, 104, -1000, 44, 1458, 1458, -1000,
	1458, 186, 1542, 1542, 1542, 1542, 171, -12, -9, -1000,
	-1000, 1458, -1000, 44, 44, 44, 44, 1542, 592, 1458,
	100, 87, -9, -1000, -9, -1000, -9, -9,
}
var RubyPgo = []int{

	0, 634, 337, 336, 6, 333, 332, 331, 502, 328,
	327, 326, 323, 320, 0, 407, 319, 312, 316, 315,
	314, 5, 313, 1, 216, 307, 3, 302, 300, 293,
	288, 287, 283, 282, 281, 278, 277, 276, 506, 271,
	9, 12, 2, 269, 7, 267, 266, 264, 10, 263,
	30, 262, 259, 8, 4, 254, 47,
}
var RubyR1 = []int{

	0, 43, 43, 43, 43, 43, 43, 43, 43, 43,
	43, 56, 56, 38, 38, 38, 38, 38, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 21, 21, 21, 21, 21, 21, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 40, 40, 50, 50, 48, 48, 48, 48, 53,
	53, 53, 53, 52, 52, 52, 18, 18, 44, 44,
	23, 23, 23, 23, 54, 54, 54, 22, 22, 25,
	26, 26, 55, 55, 11, 11, 11, 11, 11, 11,
	8, 8, 24, 24, 15, 15, 27, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 2, 5,
	6, 6, 3, 3, 45, 45, 45, 45, 51, 51,
	51, 4, 4, 4, 4, 41, 49, 49, 49, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 42,
	42, 42, 42, 39, 39, 39, 7, 13, 47, 47,
	47, 47, 19, 20, 9, 12, 46, 46, 46, 46,
	46, 46, 16, 16, 16, 16, 16, 16, 16,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 4, 1, 4, 4, 2, 3, 3,
	4, 4, 3, 2, 3, 3, 3, 3, 3, 4,
	6, 3, 1, 1, 3, 0, 1, 1, 3, 1,
	1, 3, 3, 1, 3, 3, 7, 7, 1, 3,
	1, 2, 3, 2, 0, 1, 3, 4, 6, 4,
	1, 3, 1, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 3, 3, 3, 3, 2, 2,
	2, 2, 3, 3, 3, 3, 3, 3, 1, 1,
	3, 3, 5, 5, 0, 4, 7, 8, 3, 7,
	8, 3, 4, 4, 3, 3, 0, 1, 3, 4,
	5, 3, 3, 3, 3, 3, 5, 6, 5, 4,
	3, 3, 2, 0, 2, 2, 3, 4, 0, 3,
	4, 6, 2, 2, 5, 5, 0, 2, 2, 2,
	2, 2, 0, 1, 3, 3, 1, 3, 3,
}
var RubyChk = []int{

	-1000, -43, 44, 45, 61, -1, 44, 45, 61, -14,
	-18, -22, -25, -11, -28, -29, -30, -31, -10, -13,
	-21, -19, -20, -12, 5, 6, 8, -24, -15, -8,
	-2, -5, -6, -3, -26, -9, -17, -7, 13, 19,
	20, -27, 38, 39, 40, 41, 15, 18, 24, -32,
	-33, -34, -35, -36, -37, 32, 31, 22, 59, 58,
	33, 34, 54, 56, -55, 7, 9, 45, 44, 61,
	15, 48, 35, 36, 4, 40, 41, 42, 50, 51,
	49, 18, 52, 18, 9, -40, -53, 54, 37, 11,
	-52, -14, -4, 6, 8, -24, -15, -8, -17, 12,
	56, 9, 37, 37, 37, 37, 48, 35, 36, 15,
	18, 48, 6, 4, -26, 8, -26, 37, 11, -1,
	-1, -1, -1, -1, -1, -38, -50, -14, -50, -1,
	6, 8, 59, 6, 8, -50, -48, -14, -21, -56,
	47, 9, -39, -1, 6, -14, -14, -14, -14, -1,
	-14, -14, -14, -14, -1, -14, -1, -48, -14, 11,
	-14, -14, 6, 11, -38, -41, 49, -41, -38, -48,
	-1, -1, -1, -1, 6, -14, -14, -1, -1, 6,
	-44, -54, 9, -23, 6, 42, 51, -44, -38, 35,
	-38, -1, 6, -38, 44, 45, 44, 45, -14, -1,
	-47, 11, 44, 6, 8, 55, 11, 55, 44, -45,
	-51, -14, 6, 8, -48, -1, 45, 10, -53, -40,
	9, 46, 10, -14, -4, 55, -14, -4, 14, -38,
	-49, 6, -38, 57, 10, -56, 11, -54, 37, 6,
	6, -56, 14, -26, 14, 14, -42, 17, 16, -38,
	-38, 14, 25, -14, -46, -14, -56, -56, 11, 4,
	46, 10, -48, -14, 37, 14, 49, 11, 57, -38,
	-23, 10, -1, -38, -38, 14, 17, 16, -1, -38,
	14, -42, 14, -38, 8, 14, 44, 45, -14, -1,
	-16, 27, 28, 57, 11, 57, -56, -14, -1, -14,
	10, -1, 6, -56, -56, 14, -1, -38, -38, 14,
	-38, 4, 15, 18, 15, 18, -1, 6, -56, 14,
	14, -38, 6, -1, -1, -1, -1, 4, 46, -38,
	-1, -14, -56, 11, -56, 11, -56, -56,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 38,
	39, 40, 41, 42, 43, 44, 45, 46, 47, 48,
	49, 50, 51, 52, 18, 19, -2, 21, 22, 23,
	24, 25, 26, 27, 28, 29, 30, 31, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 13, 32,
	33, 34, 35, 36, 37, 0, 0, 0, 0, 0,
	128, 129, 75, 11, 0, 54, 163, 5, 6, 7,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, -2, 57, 63, 75, 0, 0,
	72, 79, 80, 19, -2, 21, 22, 23, 30, 13,
	-2, 75, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 94, 94, 13, -2, 13, 0, 0, 118,
	119, 120, 121, 13, 0, 168, 172, 73, 173, 0,
	112, 113, 0, 110, 111, 0, 0, 73, 77, 134,
	0, 75, 0, 151, 58, 64, 66, 68, 122, 123,
	124, 125, 126, 127, 153, 0, 155, 0, 76, 0,
	73, 104, 116, 0, 0, 13, 146, 13, 0, 0,
	105, 106, 107, 108, 59, 65, 67, 152, 154, 62,
	11, 88, 94, 95, 90, 0, 0, 11, 0, 0,
	0, 109, 117, 0, 13, 13, 14, 15, 16, 17,
	0, 0, 176, 114, 115, 130, 0, 131, 12, 11,
	11, 0, 19, -2, 0, 164, 165, 166, 60, 61,
	-2, 0, 53, 81, 82, 69, 84, 85, 141, 0,
	0, 147, 0, 144, 56, 13, 0, 0, 0, 91,
	93, 13, 97, 13, 99, 149, 0, 0, 13, 0,
	0, 167, 13, 74, 0, 78, 0, 0, 11, 0,
	0, 55, 0, 174, 0, 142, 145, 0, 143, 11,
	96, 89, 92, 11, 0, 150, 0, 13, 13, 162,
	156, 0, 158, 169, 13, 175, 177, 178, 38, 180,
	181, 183, 186, 132, 0, 133, 0, 38, 11, 138,
	71, 70, 148, 0, 0, 98, 13, 160, 161, 157,
	170, 0, 0, 0, 0, 0, 0, 0, 135, 86,
	87, 159, 13, 184, 185, 187, 188, 0, 0, 171,
	11, 11, 136, 11, 139, 11, 137, 140,
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
		//line parser.y:169
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:171
		{
		}
	case 3:
		//line parser.y:173
		{
		}
	case 4:
		//line parser.y:175
		{
		}
	case 5:
		//line parser.y:177
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:179
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:181
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:187
		{
		}
	case 11:
		//line parser.y:189
		{
		}
	case 12:
		//line parser.y:190
		{
		}
	case 13:
		//line parser.y:193
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 14:
		//line parser.y:195
		{
		}
	case 15:
		//line parser.y:197
		{
		}
	case 16:
		//line parser.y:199
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 17:
		//line parser.y:201
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 53:
		//line parser.y:211
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 54:
		//line parser.y:218
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 55:
		//line parser.y:222
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 56:
		//line parser.y:229
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 57:
		//line parser.y:236
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 58:
		//line parser.y:243
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 59:
		//line parser.y:250
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 60:
		//line parser.y:257
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 61:
		//line parser.y:265
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 62:
		//line parser.y:273
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 63:
		//line parser.y:283
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 64:
		//line parser.y:290
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 65:
		//line parser.y:298
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 66:
		//line parser.y:306
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 67:
		//line parser.y:314
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 68:
		//line parser.y:322
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 69:
		//line parser.y:332
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 70:
		//line parser.y:342
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 71:
		//line parser.y:351
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 72:
		//line parser.y:353
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 73:
		//line parser.y:356
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:358
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 75:
		//line parser.y:360
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 76:
		//line parser.y:362
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:364
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 78:
		//line parser.y:366
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
		//line parser.y:379
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 84:
		//line parser.y:381
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:383
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 86:
		//line parser.y:389
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 87:
		//line parser.y:397
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 88:
		//line parser.y:406
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 89:
		//line parser.y:408
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 90:
		//line parser.y:411
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 91:
		//line parser.y:413
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 92:
		//line parser.y:415
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 93:
		//line parser.y:417
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 94:
		//line parser.y:419
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 95:
		//line parser.y:421
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 96:
		//line parser.y:425
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 97:
		//line parser.y:430
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 98:
		//line parser.y:437
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 99:
		//line parser.y:447
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 100:
		//line parser.y:456
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 101:
		//line parser.y:462
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 102:
		//line parser.y:470
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 103:
		//line parser.y:474
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 104:
		//line parser.y:479
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 105:
		//line parser.y:486
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 106:
		//line parser.y:493
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 107:
		//line parser.y:500
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 108:
		//line parser.y:507
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 109:
		//line parser.y:514
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 110:
		//line parser.y:522
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 111:
		//line parser.y:524
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 112:
		//line parser.y:527
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 113:
		//line parser.y:529
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 114:
		//line parser.y:532
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 115:
		//line parser.y:534
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 116:
		//line parser.y:537
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 117:
		//line parser.y:539
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 118:
		//line parser.y:541
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 119:
		//line parser.y:542
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 120:
		//line parser.y:543
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 121:
		//line parser.y:544
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 122:
		//line parser.y:547
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 123:
		//line parser.y:556
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 124:
		//line parser.y:565
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 125:
		//line parser.y:574
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 126:
		//line parser.y:583
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 127:
		//line parser.y:592
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 128:
		//line parser.y:600
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 129:
		//line parser.y:601
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 130:
		//line parser.y:603
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 131:
		//line parser.y:605
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 132:
		//line parser.y:608
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 133:
		//line parser.y:616
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 134:
		//line parser.y:624
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 135:
		//line parser.y:626
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 136:
		//line parser.y:633
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 137:
		//line parser.y:640
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 138:
		//line parser.y:648
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 139:
		//line parser.y:655
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 140:
		//line parser.y:662
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 141:
		//line parser.y:670
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 142:
		//line parser.y:672
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 143:
		//line parser.y:679
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 144:
		//line parser.y:686
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 145:
		//line parser.y:689
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 146:
		//line parser.y:691
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 147:
		//line parser.y:693
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 148:
		//line parser.y:695
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 149:
		//line parser.y:698
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 150:
		//line parser.y:705
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 151:
		//line parser.y:713
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 152:
		//line parser.y:720
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 153:
		//line parser.y:727
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 154:
		//line parser.y:734
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 155:
		//line parser.y:741
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 156:
		//line parser.y:748
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 157:
		//line parser.y:755
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 158:
		//line parser.y:763
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 159:
		//line parser.y:772
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 160:
		//line parser.y:779
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 161:
		//line parser.y:786
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 162:
		//line parser.y:793
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 163:
		//line parser.y:800
		{
		}
	case 164:
		//line parser.y:801
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 165:
		//line parser.y:802
		{
		}
	case 166:
		//line parser.y:805
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 167:
		//line parser.y:808
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 168:
		//line parser.y:816
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 169:
		//line parser.y:818
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 170:
		//line parser.y:820
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 171:
		//line parser.y:829
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
	case 172:
		//line parser.y:844
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 173:
		//line parser.y:852
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 174:
		//line parser.y:861
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 175:
		//line parser.y:870
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 176:
		//line parser.y:873
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 177:
		//line parser.y:875
		{
		}
	case 178:
		//line parser.y:877
		{
		}
	case 179:
		//line parser.y:879
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 180:
		//line parser.y:881
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 181:
		//line parser.y:883
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 182:
		//line parser.y:885
		{
		}
	case 183:
		//line parser.y:887
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 184:
		//line parser.y:889
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 185:
		//line parser.y:891
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 186:
		//line parser.y:893
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 187:
		//line parser.y:895
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 188:
		//line parser.y:897
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	}
	goto Rubystack /* stack new state and value */
}
