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

//line parser.y:858

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 27,
	45, 102,
	-2, 20,
	-1, 87,
	9, 72,
	10, 72,
	-2, 164,
	-1, 98,
	45, 102,
	-2, 20,
	-1, 103,
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
	-1, 114,
	45, 102,
	-2, 100,
	-1, 125,
	45, 102,
	-2, 20,
	-1, 216,
	45, 103,
	-2, 101,
	-1, 219,
	9, 72,
	10, 72,
	-2, 164,
}

const RubyNprod = 175
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1336

var RubyAct = []int{

	9, 185, 22, 183, 173, 182, 35, 134, 96, 65,
	211, 95, 138, 69, 89, 2, 3, 174, 280, 88,
	86, 128, 129, 281, 209, 139, 76, 94, 264, 304,
	257, 220, 4, 211, 10, 300, 207, 63, 62, 261,
	237, 117, 67, 68, 113, 115, 123, 70, 71, 72,
	211, 110, 109, 297, 64, 66, 75, 73, 74, 77,
	135, 104, 136, 279, 211, 263, 116, 210, 142, 143,
	144, 145, 130, 148, 149, 150, 151, 28, 153, 208,
	155, 156, 157, 158, 25, 97, 98, 46, 165, 211,
	136, 167, 169, 136, 137, 164, 108, 106, 94, 191,
	168, 299, 85, 303, 99, 146, 186, 165, 175, 136,
	57, 58, 65, 65, 177, 309, 249, 159, 189, 65,
	285, 209, 137, 99, 107, 137, 104, 250, 202, 311,
	211, 59, 104, 60, 166, 56, 55, 99, 255, 214,
	187, 137, 94, 268, 235, 99, 99, 99, 99, 188,
	99, 99, 99, 99, 235, 99, 217, 99, 99, 99,
	99, 218, 233, 209, 171, 99, 106, 222, 99, 99,
	81, 278, 169, 202, 114, 99, 104, 202, 223, 302,
	225, 221, 209, 226, 99, 205, 206, 216, 236, 287,
	186, 202, 184, 202, 239, 234, 202, 202, 242, 79,
	80, 69, 240, 295, 82, 81, 104, 112, 251, 111,
	252, 238, 78, 25, 215, 125, 99, 131, 132, 99,
	165, 260, 136, 230, 187, 253, 254, 258, 194, 202,
	67, 68, 202, 188, 79, 80, 170, 267, 163, 57,
	58, 154, 141, 66, 99, 256, 81, 78, 61, 99,
	211, 93, 213, 133, 137, 259, 83, 229, 284, 84,
	59, 204, 60, 29, 56, 55, 212, 202, 282, 1,
	245, 202, 202, 126, 53, 79, 80, 202, 202, 288,
	82, 52, 51, 289, 50, 99, 49, 99, 78, 272,
	100, 274, 273, 202, 202, 202, 298, 99, 99, 48,
	127, 18, 202, 30, 17, 307, 202, 16, 15, 100,
	39, 13, 12, 23, 11, 21, 14, 19, 24, 308,
	310, 33, 312, 100, 313, 32, 34, 31, 0, 0,
	101, 100, 100, 100, 100, 99, 100, 100, 100, 100,
	0, 100, 0, 100, 100, 100, 100, 0, 0, 101,
	0, 100, 0, 0, 100, 100, 0, 0, 0, 0,
	20, 100, 0, 101, 0, 0, 25, 124, 125, 0,
	100, 101, 101, 101, 101, 0, 101, 101, 101, 101,
	0, 101, 99, 101, 101, 101, 101, 105, 0, 0,
	0, 101, 57, 58, 101, 101, 0, 0, 0, 0,
	0, 101, 100, 172, 176, 100, 0, 0, 0, 0,
	101, 0, 0, 59, 190, 60, 192, 56, 55, 203,
	105, 5, 0, 195, 196, 0, 0, 0, 0, 0,
	100, 105, 0, 0, 0, 100, 0, 0, 0, 0,
	0, 0, 101, 105, 0, 101, 0, 0, 105, 0,
	0, 105, 105, 0, 0, 0, 0, 0, 105, 0,
	118, 119, 120, 121, 122, 0, 0, 105, 0, 0,
	101, 100, 0, 100, 228, 101, 231, 0, 0, 0,
	0, 0, 0, 100, 100, 140, 0, 0, 0, 0,
	0, 147, 0, 25, 124, 125, 152, 0, 0, 102,
	0, 0, 105, 160, 161, 162, 0, 0, 0, 0,
	0, 101, 0, 101, 0, 0, 0, 0, 0, 57,
	58, 100, 0, 101, 101, 0, 0, 178, 179, 180,
	181, 0, 105, 0, 0, 266, 193, 0, 0, 0,
	59, 270, 103, 271, 56, 55, 197, 0, 276, 0,
	0, 277, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 101, 0, 0, 0, 0, 0, 0, 100, 0,
	0, 0, 0, 0, 0, 292, 293, 0, 0, 294,
	105, 105, 0, 0, 0, 0, 25, 26, 27, 46,
	0, 0, 301, 36, 244, 44, 247, 246, 45, 37,
	38, 0, 0, 305, 47, 0, 0, 0, 101, 0,
	69, 54, 57, 58, 0, 0, 0, 40, 41, 42,
	43, 0, 0, 200, 201, 0, 0, 0, 0, 0,
	0, 0, 0, 59, 0, 60, 0, 56, 55, 67,
	68, 0, 0, 0, 70, 71, 72, 0, 0, 0,
	0, 0, 66, 75, 73, 74, 0, 269, 0, 224,
	0, 0, 0, 0, 0, 0, 275, 0, 0, 0,
	25, 26, 27, 46, 0, 0, 283, 36, 0, 44,
	0, 286, 45, 37, 38, 0, 0, 0, 47, 0,
	0, 0, 0, 291, 0, 54, 57, 58, 0, 0,
	296, 40, 41, 42, 43, 0, 0, 6, 7, 0,
	0, 0, 0, 0, 0, 0, 0, 59, 0, 60,
	0, 56, 55, 306, 8, 25, 26, 27, 46, 0,
	0, 0, 36, 290, 44, 0, 0, 45, 37, 38,
	0, 0, 0, 47, 0, 0, 0, 0, 0, 0,
	54, 57, 58, 0, 0, 0, 40, 41, 42, 43,
	0, 0, 200, 201, 0, 0, 25, 26, 27, 46,
	0, 0, 59, 36, 60, 44, 56, 55, 45, 37,
	38, 0, 0, 0, 47, 0, 0, 0, 0, 0,
	0, 54, 57, 58, 0, 0, 0, 40, 41, 42,
	43, 0, 0, 200, 201, 0, 0, 0, 0, 0,
	0, 0, 0, 59, 0, 60, 265, 56, 55, 25,
	26, 27, 46, 0, 0, 0, 36, 262, 44, 0,
	0, 45, 37, 38, 0, 0, 0, 47, 0, 0,
	0, 0, 0, 0, 54, 57, 58, 0, 0, 0,
	40, 41, 42, 43, 0, 0, 200, 201, 0, 0,
	0, 0, 0, 0, 0, 0, 59, 0, 60, 0,
	56, 55, 25, 26, 27, 46, 0, 0, 0, 36,
	248, 44, 0, 0, 45, 37, 38, 0, 0, 0,
	47, 0, 0, 0, 0, 0, 0, 54, 57, 58,
	0, 0, 0, 40, 41, 42, 43, 0, 0, 200,
	201, 0, 0, 0, 0, 0, 0, 0, 0, 59,
	0, 60, 0, 56, 55, 25, 26, 27, 46, 0,
	0, 0, 36, 243, 44, 0, 0, 45, 37, 38,
	0, 0, 0, 47, 0, 0, 0, 0, 0, 0,
	54, 57, 58, 0, 0, 0, 40, 41, 42, 43,
	0, 0, 200, 201, 0, 0, 0, 0, 0, 0,
	0, 0, 59, 0, 60, 0, 56, 55, 25, 26,
	27, 46, 0, 0, 0, 36, 241, 44, 0, 0,
	45, 37, 38, 0, 0, 0, 47, 0, 0, 0,
	0, 0, 0, 54, 57, 58, 0, 0, 0, 40,
	41, 42, 43, 0, 0, 200, 201, 0, 0, 25,
	26, 27, 46, 0, 0, 59, 36, 60, 44, 56,
	55, 45, 37, 38, 0, 0, 0, 47, 0, 0,
	0, 0, 0, 0, 54, 57, 58, 0, 0, 0,
	40, 41, 42, 43, 0, 0, 200, 201, 0, 0,
	0, 0, 0, 0, 0, 0, 59, 0, 60, 232,
	56, 55, 25, 26, 27, 46, 0, 0, 0, 36,
	227, 44, 0, 0, 45, 37, 38, 0, 0, 0,
	47, 0, 0, 0, 0, 0, 0, 54, 57, 58,
	0, 0, 0, 40, 41, 42, 43, 0, 0, 200,
	201, 0, 0, 25, 26, 27, 46, 0, 0, 59,
	36, 60, 44, 56, 55, 45, 37, 38, 0, 0,
	0, 47, 0, 0, 0, 0, 0, 0, 54, 57,
	58, 0, 0, 0, 40, 41, 42, 43, 0, 0,
	200, 201, 0, 0, 25, 26, 27, 46, 199, 0,
	59, 36, 60, 44, 56, 55, 45, 37, 38, 0,
	0, 0, 47, 0, 0, 0, 0, 0, 0, 54,
	57, 58, 0, 0, 0, 40, 41, 42, 43, 0,
	0, 0, 198, 0, 0, 25, 26, 27, 46, 0,
	0, 59, 36, 60, 44, 56, 55, 45, 37, 38,
	0, 0, 0, 47, 0, 0, 0, 0, 0, 0,
	54, 57, 58, 0, 0, 0, 40, 41, 42, 43,
	25, 97, 98, 87, 0, 92, 102, 25, 97, 98,
	46, 0, 59, 102, 60, 0, 56, 55, 25, 97,
	98, 219, 0, 0, 102, 0, 57, 58, 0, 0,
	91, 0, 0, 57, 58, 25, 97, 98, 87, 0,
	0, 102, 0, 0, 57, 58, 0, 90, 0, 103,
	0, 56, 55, 0, 59, 0, 103, 0, 56, 55,
	69, 57, 58, 0, 0, 59, 0, 103, 0, 56,
	55, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 90, 0, 103, 0, 56, 55, 0, 67,
	68, 0, 0, 0, 70, 71, 72, 0, 0, 0,
	0, 0, 66, 75, 73, 74,
}
var RubyPact = []int{

	-27, 665, -1000, -1000, -1000, -5, -1000, -1000, -1000, 9,
	242, -1000, -1000, -1000, 85, -1000, -1000, -1000, -1000, -1000,
	-26, -1000, -1000, -1000, -1000, -1000, 1225, 89, 61, 17,
	16, -1000, -1000, -1000, -1000, -1000, 203, 167, 167, 31,
	1190, 1190, 1190, 1190, 1190, 361, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 15, 211, -1000, -1000, 79,
	-1000, -20, -1000, -1000, -1000, 1190, 236, 361, 361, 361,
	79, 1190, 361, 361, 361, 361, 1190, 361, 235, 361,
	361, 361, 79, 1190, 1190, 1190, 232, 79, -1000, 124,
	79, 79, 230, 154, 197, -1000, -1000, 1260, 158, -1000,
	-1000, -1000, -30, -30, 201, -26, 79, 1190, 1190, 1190,
	1190, 184, 184, 66, -1000, -1000, 1190, 222, 98, 98,
	98, 98, 98, -1000, -1000, -1000, 1149, 1108, -1000, -1000,
	179, -1000, -1000, 26, 14, 1286, -1000, 166, 208, 180,
	98, 1243, -1000, -1000, -1000, 197, 201, 98, -1000, -1000,
	-1000, -1000, 98, -13, -1000, -1000, -1000, -1000, 197, 201,
	98, 98, 98, -1000, 172, 1286, 488, 606, -1000, 197,
	-1000, 1232, 1067, -1000, 217, -1000, 1014, 153, 98, 98,
	98, 98, -1000, 144, 100, -1000, 5, 205, 188, -1000,
	973, 167, 920, 98, -1000, 581, 867, 98, -1000, -1000,
	-1000, -1000, 9, 98, 103, -1000, -1000, 361, -1000, 361,
	-1000, -1000, -1000, 128, 241, -14, -1000, 124, -1000, 79,
	79, -1000, -1000, -1000, 4, -1000, -1000, -1000, 814, 18,
	-1000, 761, -1000, -1000, -9, 100, 134, 1190, -1000, -1000,
	-9, -1000, -1000, -1000, -1000, 276, 1190, -1000, -1000, -1000,
	164, -1000, -1000, 8, -32, -1000, 1190, 361, 111, 201,
	197, 1190, -1000, -1000, 183, -1000, 1108, -1000, -1000, 98,
	1108, 720, -1000, 1190, -1000, 98, 1108, 1108, 199, -1000,
	1190, -1000, 47, 98, -1000, -1000, 98, -1000, 88, 22,
	-1000, 98, 1108, 1108, 1108, 173, 99, -15, -9, -1000,
	-1000, 1108, -1000, 1190, 361, 1108, 105, 119, -9, -1000,
	-9, -1000, -9, -9,
}
var RubyPgo = []int{

	0, 419, 327, 326, 11, 325, 321, 360, 303, 318,
	317, 316, 315, 0, 263, 34, 314, 313, 2, 312,
	1, 77, 311, 6, 8, 310, 308, 307, 304, 301,
	299, 286, 284, 282, 281, 274, 300, 273, 19, 4,
	270, 269, 5, 266, 261, 7, 257, 253, 252, 251,
	14, 3, 248, 12,
}
var RubyR1 = []int{

	0, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 53, 53, 36, 36, 36, 36, 36, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 18,
	18, 18, 18, 18, 18, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 38, 38,
	47, 47, 45, 45, 45, 45, 45, 50, 50, 50,
	50, 24, 24, 49, 49, 49, 16, 16, 42, 42,
	20, 20, 20, 20, 51, 51, 51, 19, 19, 22,
	23, 23, 52, 52, 11, 11, 11, 11, 11, 11,
	8, 8, 21, 21, 14, 14, 25, 25, 26, 27,
	28, 29, 30, 30, 30, 30, 31, 32, 33, 34,
	35, 2, 5, 6, 6, 3, 3, 43, 43, 43,
	43, 48, 48, 48, 4, 4, 4, 4, 39, 46,
	46, 46, 10, 10, 10, 10, 10, 10, 10, 10,
	40, 40, 40, 40, 37, 37, 37, 7, 12, 44,
	44, 44, 44, 17, 9,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 4, 4, 2, 3, 3, 4, 4, 3, 2,
	3, 3, 3, 3, 3, 3, 4, 6, 3, 1,
	1, 3, 0, 1, 1, 1, 3, 1, 1, 3,
	3, 1, 1, 1, 3, 3, 7, 7, 1, 3,
	1, 2, 3, 2, 0, 1, 3, 4, 6, 4,
	1, 3, 1, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 3, 3, 3, 3, 2, 2,
	2, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 1, 1, 3, 3, 5, 5, 0, 4, 7,
	8, 3, 7, 8, 3, 4, 4, 3, 3, 0,
	1, 3, 4, 5, 3, 3, 3, 3, 3, 4,
	4, 3, 3, 2, 0, 2, 2, 3, 4, 0,
	3, 4, 6, 1, 5,
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
	34, 4, 38, 14, 17, 17, 46, 8, -38, -50,
	52, 35, 10, -49, -13, -4, -24, 6, 7, -21,
	-14, -8, 11, 54, -15, -7, 8, 35, 35, 35,
	35, 6, 4, -23, 7, -23, 35, 10, -1, -1,
	-1, -1, -1, -13, 6, 7, -37, -36, 6, 7,
	57, 6, 7, -47, -45, -13, -18, -15, -53, 45,
	-1, 6, -13, -13, -13, -13, -15, -1, -13, -13,
	-13, -13, -1, -13, 6, -13, -13, -13, -13, -15,
	-1, -1, -1, 6, -45, -13, 10, -13, -24, -13,
	6, 10, -36, -39, 47, -39, -36, -45, -1, -1,
	-1, -1, -42, -51, 8, -20, 6, 40, 49, -42,
	-36, 33, -36, -1, 6, -36, -36, -1, 43, 9,
	42, 43, -13, -1, -44, 6, 7, 10, 53, 10,
	53, 42, -43, -48, -13, 6, 7, -50, -38, 8,
	44, 9, -13, -4, 53, -24, -4, 13, -36, -46,
	6, -36, 55, 9, -53, 10, -51, 35, 6, 6,
	-53, 13, -23, 13, 13, -40, 16, 15, 13, 13,
	24, -13, -13, -53, -53, 10, 4, 44, -45, -15,
	-13, 35, 13, 47, 10, 55, -36, -20, 9, -1,
	-36, -36, 13, 16, 15, -1, -36, -36, 7, 55,
	10, 55, -53, -1, -13, 9, -1, 6, -53, -53,
	13, -1, -36, -36, -36, 4, -1, 6, -53, 13,
	13, -36, 6, 4, 44, -36, -1, -13, -53, 10,
	-53, 10, -53, -53,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
	46, 47, 48, 49, 50, 18, 19, -2, 21, 22,
	23, 24, 25, 26, 27, 28, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 164, 13, 29, 30,
	31, 32, 33, 34, 173, 0, 0, 131, 132, 72,
	11, 0, 5, 6, 7, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, -2, 53, 59,
	72, 0, 0, 69, 77, 78, 83, 19, -2, 21,
	22, 23, 13, -2, 82, 0, 72, 0, 0, 0,
	0, 94, 94, 13, -2, 13, 0, 0, 118, 119,
	120, 121, 13, 13, 19, -2, 0, 169, 112, 113,
	0, 110, 111, 0, 0, 70, 74, 75, 137, 0,
	154, 54, 60, 62, 64, 122, 124, 126, 127, 128,
	129, 130, 156, 0, 55, 61, 63, 65, 123, 125,
	155, 157, 158, 58, 0, 73, 0, 70, 104, 81,
	116, 0, 0, 13, 149, 13, 0, 0, 105, 106,
	107, 108, 11, 88, 94, 95, 90, 0, 0, 11,
	0, 0, 0, 109, 117, 0, 0, 165, 166, 167,
	14, 15, 16, 17, 0, 114, 115, 0, 133, 0,
	134, 12, 11, 11, 0, 19, -2, 56, 57, -2,
	0, 51, 79, 80, 66, 84, 85, 144, 0, 0,
	150, 0, 147, 52, 13, 0, 0, 0, 91, 93,
	13, 97, 13, 99, 152, 0, 0, 13, 159, 168,
	13, 71, 76, 0, 0, 11, 0, 0, 0, 174,
	0, 0, 145, 148, 0, 146, 11, 96, 89, 92,
	11, 0, 153, 0, 13, 13, 163, 170, 13, 135,
	0, 136, 0, 11, 141, 68, 67, 151, 0, 0,
	98, 13, 161, 162, 171, 0, 0, 0, 138, 86,
	87, 160, 13, 0, 0, 172, 11, 11, 139, 11,
	142, 11, 140, 143,
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
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 53:
		//line parser.y:220
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 54:
		//line parser.y:227
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 55:
		//line parser.y:234
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 56:
		//line parser.y:241
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 57:
		//line parser.y:249
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 58:
		//line parser.y:257
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 59:
		//line parser.y:267
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 60:
		//line parser.y:274
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 61:
		//line parser.y:282
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 62:
		//line parser.y:290
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 63:
		//line parser.y:298
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 64:
		//line parser.y:306
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 65:
		//line parser.y:314
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 66:
		//line parser.y:324
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 67:
		//line parser.y:334
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 68:
		//line parser.y:343
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 69:
		//line parser.y:345
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 70:
		//line parser.y:348
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 71:
		//line parser.y:350
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 72:
		//line parser.y:352
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 73:
		//line parser.y:354
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:356
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 82:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 83:
		//line parser.y:375
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 84:
		//line parser.y:377
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:379
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 86:
		//line parser.y:385
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 87:
		//line parser.y:393
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 88:
		//line parser.y:402
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 89:
		//line parser.y:404
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 90:
		//line parser.y:407
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 91:
		//line parser.y:409
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 92:
		//line parser.y:411
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 93:
		//line parser.y:413
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 94:
		//line parser.y:415
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 95:
		//line parser.y:417
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 96:
		//line parser.y:421
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 97:
		//line parser.y:426
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 98:
		//line parser.y:433
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 99:
		//line parser.y:443
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 100:
		//line parser.y:452
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 101:
		//line parser.y:458
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 102:
		//line parser.y:466
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 103:
		//line parser.y:470
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 104:
		//line parser.y:475
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 105:
		//line parser.y:482
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 106:
		//line parser.y:489
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 107:
		//line parser.y:496
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 108:
		//line parser.y:503
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 109:
		//line parser.y:510
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 110:
		//line parser.y:518
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 111:
		//line parser.y:520
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 112:
		//line parser.y:523
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 113:
		//line parser.y:525
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 114:
		//line parser.y:528
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 115:
		//line parser.y:530
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 116:
		//line parser.y:533
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 117:
		//line parser.y:535
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 118:
		//line parser.y:537
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 119:
		//line parser.y:538
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 120:
		//line parser.y:539
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 121:
		//line parser.y:540
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 122:
		//line parser.y:543
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 123:
		//line parser.y:551
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 124:
		//line parser.y:559
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 125:
		//line parser.y:567
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 126:
		//line parser.y:576
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 127:
		//line parser.y:585
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 128:
		//line parser.y:594
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 129:
		//line parser.y:603
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 130:
		//line parser.y:612
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 131:
		//line parser.y:620
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 132:
		//line parser.y:621
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 133:
		//line parser.y:623
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 134:
		//line parser.y:625
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 135:
		//line parser.y:628
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 136:
		//line parser.y:636
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 137:
		//line parser.y:644
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 138:
		//line parser.y:646
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 139:
		//line parser.y:653
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 140:
		//line parser.y:660
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 141:
		//line parser.y:668
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 142:
		//line parser.y:675
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 143:
		//line parser.y:682
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 144:
		//line parser.y:690
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 145:
		//line parser.y:692
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 146:
		//line parser.y:699
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 147:
		//line parser.y:706
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 148:
		//line parser.y:709
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 149:
		//line parser.y:711
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 150:
		//line parser.y:713
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 151:
		//line parser.y:715
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 152:
		//line parser.y:718
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 153:
		//line parser.y:725
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 154:
		//line parser.y:733
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 155:
		//line parser.y:740
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 156:
		//line parser.y:747
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 157:
		//line parser.y:754
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 158:
		//line parser.y:761
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 159:
		//line parser.y:768
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-2].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 160:
		//line parser.y:776
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 161:
		//line parser.y:783
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 162:
		//line parser.y:790
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 163:
		//line parser.y:797
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 164:
		//line parser.y:804
		{
		}
	case 165:
		//line parser.y:805
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 166:
		//line parser.y:806
		{
		}
	case 167:
		//line parser.y:809
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 168:
		//line parser.y:812
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 169:
		//line parser.y:820
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 170:
		//line parser.y:822
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 171:
		//line parser.y:824
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 172:
		//line parser.y:833
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
	case 173:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 174:
		//line parser.y:850
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
