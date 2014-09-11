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
const TRUE = 57372
const FALSE = 57373
const LESSTHAN = 57374
const GREATERTHAN = 57375
const EQUALTO = 57376
const BANG = 57377
const COMPLEMENT = 57378
const POSITIVE = 57379
const NEGATIVE = 57380
const STAR = 57381
const WHITESPACE = 57382
const NEWLINE = 57383
const SEMICOLON = 57384
const COLON = 57385
const DOT = 57386
const PIPE = 57387
const SLASH = 57388
const AMPERSAND = 57389
const MODULO = 57390
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
	"MODULO",
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

//line parser.y:667

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 13,
	-1, 12,
	1, 18,
	4, 18,
	9, 18,
	10, 18,
	13, 18,
	15, 18,
	16, 18,
	17, 18,
	24, 18,
	39, 18,
	41, 18,
	42, 18,
	44, 18,
	45, 18,
	46, 18,
	47, 18,
	53, 18,
	57, 18,
	-2, 13,
	-1, 13,
	34, 13,
	43, 80,
	-2, 19,
	-1, 14,
	34, 13,
	-2, 20,
	-1, 15,
	34, 13,
	-2, 21,
	-1, 16,
	34, 13,
	-2, 22,
	-1, 22,
	14, 13,
	-2, 29,
	-1, 88,
	1, 28,
	9, 28,
	10, 28,
	12, 28,
	13, 28,
	15, 28,
	16, 28,
	18, 28,
	19, 28,
	23, 28,
	24, 28,
	35, 28,
	36, 28,
	41, 28,
	42, 28,
	53, 28,
	57, 28,
	-2, 13,
	-1, 101,
	10, 13,
	-2, 56,
	-1, 102,
	4, 13,
	5, 13,
	6, 13,
	7, 13,
	8, 13,
	30, 13,
	31, 13,
	50, 13,
	52, 13,
	54, 13,
	55, 13,
	-2, 64,
	-1, 104,
	1, 18,
	4, 18,
	9, 18,
	10, 18,
	13, 18,
	15, 18,
	16, 18,
	24, 18,
	41, 18,
	42, 18,
	44, 18,
	53, 18,
	57, 18,
	-2, 13,
	-1, 105,
	43, 80,
	-2, 19,
	-1, 109,
	5, 13,
	6, 13,
	7, 13,
	8, 13,
	30, 13,
	31, 13,
	50, 13,
	52, 13,
	54, 13,
	55, 13,
	-2, 51,
	-1, 113,
	10, 13,
	-2, 53,
	-1, 117,
	45, 13,
	-2, 9,
	-1, 131,
	43, 80,
	-2, 78,
	-1, 173,
	43, 81,
	-2, 79,
	-1, 199,
	45, 118,
	-2, 57,
	-1, 228,
	45, 119,
	-2, 18,
	-1, 246,
	4, 13,
	5, 13,
	6, 13,
	7, 13,
	8, 13,
	30, 13,
	31, 13,
	50, 13,
	52, 13,
	54, 13,
	55, 13,
	-2, 66,
	-1, 298,
	45, 120,
	-2, 18,
}

const RubyNprod = 146
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1211

var RubyAct = []int{

	155, 22, 6, 235, 21, 116, 90, 8, 57, 2,
	58, 60, 291, 62, 63, 64, 65, 66, 289, 127,
	205, 125, 7, 67, 68, 69, 59, 295, 253, 199,
	143, 7, 4, 5, 274, 305, 294, 234, 128, 75,
	91, 11, 104, 105, 100, 73, 74, 252, 76, 77,
	78, 79, 80, 81, 82, 83, 87, 7, 171, 244,
	103, 206, 110, 262, 263, 214, 42, 43, 262, 263,
	215, 180, 61, 70, 213, 89, 171, 7, 10, 121,
	120, 119, 130, 132, 118, 140, 44, 122, 45, 142,
	41, 40, 144, 145, 146, 147, 148, 149, 150, 151,
	152, 153, 154, 220, 7, 156, 239, 316, 241, 240,
	157, 245, 205, 315, 161, 162, 224, 205, 164, 165,
	166, 167, 168, 169, 275, 259, 170, 85, 172, 257,
	174, 175, 159, 176, 115, 195, 91, 11, 104, 105,
	100, 192, 160, 286, 103, 131, 98, 173, 163, 99,
	310, 133, 134, 135, 136, 137, 138, 103, 300, 103,
	194, 129, 42, 43, 193, 123, 124, 109, 177, 92,
	93, 94, 71, 72, 308, 191, 103, 97, 95, 96,
	303, 212, 44, 260, 45, 251, 41, 40, 46, 113,
	101, 208, 222, 223, 227, 181, 225, 207, 1, 226,
	211, 221, 217, 198, 84, 36, 230, 35, 231, 232,
	233, 34, 33, 32, 31, 237, 30, 29, 184, 242,
	28, 229, 27, 25, 190, 247, 24, 23, 115, 248,
	254, 250, 39, 26, 196, 261, 37, 200, 201, 202,
	203, 268, 266, 38, 236, 272, 238, 19, 88, 18,
	243, 256, 258, 20, 17, 9, 3, 276, 0, 278,
	0, 280, 0, 0, 0, 0, 264, 283, 282, 0,
	0, 0, 269, 0, 0, 271, 288, 277, 290, 279,
	292, 0, 273, 0, 14, 0, 0, 297, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 249, 0, 0,
	304, 285, 0, 0, 307, 0, 309, 102, 0, 311,
	111, 0, 0, 0, 270, 296, 0, 0, 0, 11,
	104, 105, 126, 0, 0, 117, 301, 0, 302, 11,
	104, 105, 100, 281, 0, 0, 0, 0, 0, 0,
	284, 313, 0, 106, 42, 43, 106, 0, 0, 0,
	0, 15, 293, 0, 42, 43, 0, 0, 106, 126,
	0, 0, 299, 0, 44, 0, 45, 0, 41, 40,
	0, 0, 0, 0, 44, 0, 45, 0, 41, 40,
	0, 312, 0, 314, 16, 0, 0, 0, 0, 0,
	0, 102, 0, 182, 183, 106, 185, 186, 187, 188,
	189, 0, 126, 0, 102, 0, 102, 0, 0, 0,
	107, 0, 0, 107, 0, 0, 0, 0, 204, 0,
	0, 209, 0, 102, 0, 107, 0, 106, 0, 106,
	106, 0, 106, 106, 106, 106, 106, 0, 106, 0,
	106, 0, 106, 108, 0, 0, 108, 0, 126, 0,
	0, 0, 0, 0, 106, 0, 0, 106, 108, 106,
	0, 0, 107, 0, 0, 0, 0, 11, 112, 105,
	0, 158, 246, 0, 0, 0, 0, 0, 0, 255,
	0, 0, 0, 0, 106, 0, 0, 0, 0, 0,
	0, 0, 42, 43, 107, 108, 107, 107, 0, 107,
	107, 107, 107, 107, 0, 107, 0, 107, 106, 107,
	0, 0, 44, 0, 45, 106, 41, 40, 11, 112,
	105, 107, 0, 0, 107, 0, 107, 108, 0, 108,
	108, 0, 108, 108, 108, 108, 108, 255, 108, 0,
	108, 0, 108, 42, 43, 0, 0, 0, 0, 0,
	0, 107, 0, 7, 108, 0, 0, 108, 0, 108,
	0, 0, 0, 44, 0, 45, 0, 41, 40, 0,
	0, 0, 0, 106, 0, 107, 0, 0, 0, 0,
	0, 0, 107, 0, 108, 0, 11, 12, 13, 55,
	0, 0, 0, 47, 216, 54, 219, 218, 0, 48,
	49, 0, 0, 0, 56, 0, 0, 0, 108, 0,
	0, 42, 43, 0, 0, 108, 50, 51, 52, 53,
	0, 179, 178, 0, 0, 0, 0, 0, 0, 0,
	0, 44, 0, 45, 0, 41, 40, 0, 0, 0,
	107, 0, 0, 0, 11, 12, 13, 55, 0, 0,
	0, 47, 306, 54, 0, 0, 0, 48, 49, 0,
	0, 0, 56, 0, 0, 0, 0, 0, 0, 42,
	43, 0, 0, 108, 50, 51, 52, 53, 0, 179,
	178, 0, 0, 0, 0, 0, 0, 0, 0, 44,
	0, 45, 0, 41, 40, 11, 12, 13, 55, 0,
	0, 0, 47, 287, 54, 0, 0, 0, 48, 49,
	0, 0, 0, 56, 0, 0, 0, 0, 0, 0,
	42, 43, 0, 0, 0, 50, 51, 52, 53, 0,
	179, 178, 0, 0, 0, 0, 0, 0, 0, 0,
	44, 0, 45, 0, 41, 40, 11, 12, 13, 55,
	0, 0, 0, 47, 267, 54, 0, 0, 0, 48,
	49, 0, 0, 0, 56, 0, 0, 0, 0, 0,
	0, 42, 43, 0, 0, 0, 50, 51, 52, 53,
	0, 179, 178, 0, 0, 0, 0, 0, 0, 0,
	0, 44, 0, 45, 0, 41, 40, 11, 12, 13,
	55, 0, 0, 0, 47, 265, 54, 0, 0, 0,
	48, 49, 0, 0, 0, 56, 0, 0, 0, 0,
	0, 0, 42, 43, 0, 0, 0, 50, 51, 52,
	53, 0, 179, 178, 0, 0, 0, 0, 0, 0,
	0, 0, 44, 0, 45, 0, 41, 40, 11, 12,
	13, 55, 0, 0, 0, 47, 197, 54, 0, 0,
	0, 48, 49, 0, 0, 0, 56, 0, 0, 0,
	0, 0, 0, 42, 43, 0, 0, 0, 50, 51,
	52, 53, 0, 179, 178, 0, 0, 11, 12, 13,
	55, 141, 0, 44, 47, 45, 54, 41, 40, 0,
	48, 49, 0, 0, 0, 56, 0, 0, 0, 0,
	0, 0, 42, 43, 0, 0, 0, 50, 51, 52,
	53, 0, 7, 0, 139, 0, 11, 12, 13, 55,
	0, 0, 44, 47, 45, 54, 41, 40, 0, 48,
	49, 0, 0, 0, 56, 0, 0, 0, 0, 0,
	0, 42, 43, 0, 0, 0, 50, 51, 52, 53,
	0, 179, 178, 0, 0, 0, 0, 0, 0, 0,
	0, 44, 0, 45, 0, 41, 40, 11, 12, 13,
	55, 0, 0, 117, 47, 0, 54, 0, 0, 0,
	48, 49, 0, 0, 0, 56, 0, 0, 0, 0,
	0, 0, 42, 43, 0, 0, 114, 50, 51, 52,
	53, 0, 0, 0, 0, 0, 11, 12, 13, 55,
	0, 0, 44, 47, 45, 54, 41, 40, 0, 48,
	49, 0, 0, 0, 56, 0, 0, 0, 0, 0,
	0, 42, 43, 0, 0, 0, 50, 51, 52, 53,
	0, 7, 0, 86, 0, 0, 0, 0, 0, 0,
	0, 44, 0, 45, 0, 41, 40, 11, 12, 13,
	55, 0, 0, 117, 47, 0, 54, 0, 0, 0,
	48, 49, 0, 0, 0, 56, 0, 0, 0, 0,
	0, 0, 42, 43, 0, 0, 0, 50, 51, 52,
	53, 0, 0, 0, 0, 0, 11, 12, 13, 55,
	0, 0, 44, 47, 45, 54, 41, 40, 0, 48,
	49, 11, 298, 105, 56, 0, 0, 0, 0, 0,
	0, 42, 43, 11, 112, 105, 50, 51, 52, 53,
	0, 0, 0, 0, 0, 0, 42, 43, 11, 228,
	105, 44, 0, 45, 0, 41, 40, 0, 42, 43,
	11, 210, 105, 0, 0, 0, 44, 0, 45, 0,
	41, 40, 0, 42, 43, 0, 0, 0, 44, 0,
	45, 0, 41, 40, 0, 42, 43, 0, 0, 0,
	0, 0, 0, 44, 0, 45, 0, 41, 40, 0,
	0, 0, 0, 0, 0, 44, 0, 45, 0, 41,
	40,
}
var RubyPact = []int{

	-48, -9, -1000, -50, -1000, -1000, 1101, 37, -1000, -18,
	37, -1000, 64, 37, 37, 37, 37, -1000, -1000, -1000,
	-1000, -1000, 37, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	18, 166, -1000, -1000, 37, 37, -4, 37, 37, 37,
	37, 37, 37, 37, 37, 1011, 34, -1000, 132, 161,
	-1000, 513, 972, 50, 47, 46, 45, 73, -1000, -1000,
	159, -1000, -1000, 1128, -1000, -5, 155, 138, 138, 1101,
	1101, 1101, 1101, 1101, 882, -1000, -1000, -1000, -18, -1000,
	-1000, 37, 37, 37, 37, 37, 37, 37, 37, 37,
	37, 37, -18, -1000, 64, -1000, -1000, -1000, -1000, 37,
	462, 133, -1000, 37, 37, -1000, -1000, 37, 37, 37,
	37, 37, 37, -1000, -1000, 37, -1000, 17, 140, 37,
	37, -1000, 37, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 132, 921, 1128, 1128, 1101, 1128, 1128, 1128,
	1128, 1128, 1101, 1128, 131, 36, 1062, 324, -1000, 37,
	-1000, 125, 1101, 843, -16, 1101, 1101, 1101, 1101, 1128,
	10, -1000, 1155, -1000, 324, 33, 29, 581, -1000, -1000,
	-1000, 90, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 37, 37, -1000, 107, 37, -1000, -1000, 37, 1143,
	-1000, -1000, -1000, -1000, -1000, 37, -1000, 37, 37, 37,
	-6, -1000, -1000, -1000, 37, -1000, -1000, 93, 37, -1000,
	-1000, 35, 102, 314, -1000, 1062, 6, -17, -1000, 37,
	1128, 119, 115, 179, 37, 28, 792, 138, 741, -1000,
	37, -1000, 1101, 921, 37, -1000, -18, -1000, -1000, -1000,
	-1000, -7, -1000, -1000, 114, -1000, 17, -1000, 17, -1000,
	37, 1101, -1000, -1000, 921, -1000, 37, -1000, 1101, 921,
	-1000, 921, 136, 690, -1000, 37, -35, 17, -41, 17,
	1101, -1000, 23, -14, -1000, 921, 37, -1000, 1116, -1000,
	1101, -1000, 152, -1000, -1000, -1000, 921, 176, -1000, 37,
	-8, 639, 921, 37, 170, 37, -1000, 144, 37, 1101,
	-1000, 1101, 103, 921, 97, -1000, -1000,
}
var RubyPgo = []int{

	0, 71, 256, 254, 253, 5, 249, 247, 243, 384,
	1, 236, 233, 232, 248, 351, 227, 226, 284, 223,
	4, 222, 220, 217, 216, 214, 213, 212, 211, 207,
	205, 30, 204, 6, 203, 202, 200, 198, 197, 195,
	21, 194, 191, 190, 189, 188, 0, 19, 3, 185,
}
var RubyR1 = []int{

	0, 37, 37, 37, 37, 2, 2, 2, 2, 31,
	31, 31, 31, 46, 46, 47, 47, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 10, 10, 10,
	10, 10, 10, 10, 10, 33, 33, 40, 40, 40,
	44, 44, 44, 44, 43, 43, 43, 43, 43, 48,
	48, 48, 16, 36, 36, 17, 17, 19, 20, 20,
	45, 45, 12, 12, 12, 12, 12, 21, 22, 23,
	24, 25, 26, 27, 28, 29, 30, 3, 6, 7,
	4, 4, 38, 38, 38, 38, 42, 42, 42, 9,
	9, 18, 18, 15, 15, 5, 5, 34, 41, 41,
	41, 49, 49, 11, 11, 11, 11, 11, 35, 35,
	35, 35, 35, 32, 32, 32, 32, 32, 32, 32,
	8, 13, 39, 39, 39, 39,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 1, 3, 3, 0,
	2, 2, 2, 0, 2, 0, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 4, 6,
	3, 3, 5, 3, 5, 5, 1, 0, 1, 5,
	1, 1, 5, 5, 1, 1, 5, 5, 5, 0,
	2, 2, 9, 0, 1, 7, 11, 7, 1, 4,
	1, 4, 5, 5, 5, 5, 5, 3, 3, 3,
	3, 5, 5, 5, 5, 5, 5, 1, 1, 5,
	9, 9, 0, 5, 10, 11, 4, 9, 10, 2,
	2, 2, 2, 3, 3, 3, 7, 3, 0, 1,
	5, 1, 2, 5, 6, 5, 5, 5, 0, 5,
	3, 4, 2, 0, 1, 1, 1, 2, 2, 2,
	3, 5, 0, 4, 7, 10,
}
var RubyChk = []int{

	-1000, -37, 57, -2, 41, 42, -46, 40, 57, -14,
	-1, 5, 6, 7, -18, -15, -9, -3, -6, -7,
	-4, -20, -10, -16, -17, -19, -12, -21, -22, -23,
	-24, -25, -26, -27, -28, -29, -30, -11, -8, -13,
	55, 54, 30, 31, 50, 52, -45, 12, 18, 19,
	35, 36, 37, 38, 14, 8, 23, -46, -46, 44,
	-46, 8, -46, -46, -46, -46, -46, -46, 6, 7,
	55, 6, 7, -46, -46, 43, -46, -46, -46, -46,
	-46, -46, -46, -46, -32, -1, 42, -46, -14, 41,
	-33, 4, 37, 38, 39, 46, 47, 45, 14, 17,
	8, -43, -14, -10, 6, 7, -18, -15, -9, 6,
	-46, -14, 6, -44, 34, -1, -5, 11, 34, 34,
	34, 34, 14, 6, 7, -40, -14, -47, 43, 6,
	-20, 7, -20, -1, -1, -1, -1, -1, -1, 42,
	-46, 9, -46, -31, -46, -46, -46, -46, -46, -46,
	-46, -46, -46, -46, -46, -46, -46, -46, 9, -40,
	9, -46, -46, -31, -46, -46, -46, -46, -46, -46,
	-46, 41, -46, 7, -46, -46, -46, -31, 41, 40,
	-1, -39, -14, -14, -1, -14, -14, -14, -14, -14,
	-1, -40, 10, -33, -46, 10, -1, 13, -34, 45,
	-1, -1, -1, -1, -14, 10, 51, -38, -42, -14,
	6, -36, -33, 41, 32, 41, 13, -35, 16, 15,
	13, -47, -46, -46, 9, -46, -46, -41, 6, -40,
	-46, -46, -46, -46, 43, -48, -31, -46, -31, 13,
	16, 15, -46, -31, 24, 9, -14, -10, -5, -1,
	-5, -49, 41, 45, -46, -14, -47, 10, -47, 10,
	4, -46, 40, 41, -31, 13, -20, 13, -46, -31,
	-1, -31, -46, -31, 41, 10, -46, -47, -46, -47,
	-46, -1, -48, -46, -1, -31, 7, 13, -46, 53,
	-46, 53, -46, -1, 13, 41, -31, -46, 6, -1,
	6, -31, -31, 4, -46, 43, 13, -46, 4, -46,
	6, -46, -1, -31, -1, 10, 10,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 6, 0, 13, 4, 13,
	13, 17, -2, -2, -2, -2, -2, 23, 24, 25,
	26, 27, -2, 30, 31, 32, 33, 34, 35, 36,
	37, 38, 39, 40, 41, 42, 43, 44, 45, 46,
	0, 0, 97, 98, 13, 13, 0, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 0, 14, 7, 0,
	8, 13, 0, 0, 0, 0, 0, 0, 111, 112,
	0, 109, 110, 57, 15, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 134, 135, 136, -2, 9,
	50, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, -2, -2, 65, -2, -2, 20, 21, 22, -2,
	57, 0, 18, -2, 13, 60, 61, -2, 13, 13,
	13, 13, 13, 113, 114, 13, 58, 13, 0, 13,
	13, -2, 13, 87, 88, 89, 90, 9, 137, 138,
	139, 140, 0, 142, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 57, 0, 0, 0, 0, 47, 13,
	48, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 16, 102, -2, 73, 0, 0, 0, 10, 11,
	12, 15, 54, 91, 92, 93, 94, 95, 96, 125,
	127, 13, 13, 52, 0, 13, 82, 115, 13, -2,
	83, 84, 85, 86, 126, 13, 99, 13, 13, 13,
	18, 69, 74, 9, 13, 9, 123, 0, 13, 9,
	141, 0, 0, 0, 49, 0, 0, 0, -2, 13,
	0, 15, 15, 0, 13, 9, 0, 0, 0, 124,
	13, 9, 0, 132, 9, 55, -2, 67, 68, 62,
	63, 9, 121, 117, 0, 59, 13, 15, 13, 15,
	13, 0, 70, 71, 69, 75, 13, 77, 0, 130,
	9, 143, 0, 0, 122, 13, 0, 13, 0, 13,
	0, 106, 0, 0, 9, 131, 13, 116, 0, 100,
	0, 101, 0, 103, 72, 9, 129, 9, -2, 13,
	0, 0, 144, 13, 0, 13, 76, 0, 13, 0,
	9, 0, 107, 145, 104, 108, 105,
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
		//line parser.y:151
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:153
		{
			Statements = []ast.Node{}
		}
	case 3:
		//line parser.y:155
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 4:
		//line parser.y:161
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 5:
		//line parser.y:167
		{
			RubyVAL.genericValue = nil
		}
	case 6:
		//line parser.y:168
		{
			RubyVAL.genericValue = nil
		}
	case 7:
		//line parser.y:169
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 8:
		//line parser.y:170
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 9:
		//line parser.y:173
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 10:
		//line parser.y:175
		{
		}
	case 11:
		//line parser.y:177
		{
		}
	case 12:
		//line parser.y:179
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
			}
		}
	case 17:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
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
		//line parser.y:194
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: []ast.Node{},
			}
		}
	case 48:
		//line parser.y:201
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 49:
		//line parser.y:208
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 50:
		//line parser.y:215
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 51:
		//line parser.y:222
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 52:
		//line parser.y:229
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 53:
		//line parser.y:239
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 54:
		//line parser.y:246
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 55:
		//line parser.y:255
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 56:
		//line parser.y:257
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 57:
		//line parser.y:259
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 58:
		//line parser.y:261
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 59:
		//line parser.y:263
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 60:
		//line parser.y:267
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 61:
		//line parser.y:269
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 62:
		//line parser.y:271
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 63:
		//line parser.y:273
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 64:
		//line parser.y:276
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 65:
		//line parser.y:278
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 66:
		//line parser.y:280
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:282
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:284
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 72:
		//line parser.y:294
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 73:
		//line parser.y:302
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 74:
		//line parser.y:303
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 75:
		//line parser.y:306
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 76:
		//line parser.y:313
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-8].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-4].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-8].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 77:
		//line parser.y:323
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 78:
		//line parser.y:332
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 79:
		//line parser.y:338
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 80:
		//line parser.y:346
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 81:
		//line parser.y:350
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 82:
		//line parser.y:355
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 83:
		//line parser.y:362
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 84:
		//line parser.y:369
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 85:
		//line parser.y:376
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 86:
		//line parser.y:383
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 87:
		//line parser.y:390
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 88:
		//line parser.y:391
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 89:
		//line parser.y:392
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 90:
		//line parser.y:393
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 91:
		//line parser.y:396
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 92:
		//line parser.y:405
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 93:
		//line parser.y:414
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 94:
		//line parser.y:423
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 95:
		//line parser.y:432
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 96:
		//line parser.y:441
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 97:
		//line parser.y:449
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 98:
		//line parser.y:450
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 99:
		//line parser.y:452
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 100:
		//line parser.y:456
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 101:
		//line parser.y:464
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 102:
		//line parser.y:472
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 103:
		//line parser.y:474
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 104:
		//line parser.y:481
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 105:
		//line parser.y:488
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 106:
		//line parser.y:496
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 107:
		//line parser.y:503
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 108:
		//line parser.y:510
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
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
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 116:
		//line parser.y:535
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 117:
		//line parser.y:543
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 118:
		//line parser.y:545
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 119:
		//line parser.y:547
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 120:
		//line parser.y:549
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 123:
		//line parser.y:554
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 124:
		//line parser.y:561
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 125:
		//line parser.y:569
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 126:
		//line parser.y:576
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 127:
		//line parser.y:583
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 128:
		//line parser.y:590
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 129:
		//line parser.y:592
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 130:
		//line parser.y:599
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 131:
		//line parser.y:606
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 132:
		//line parser.y:613
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 133:
		//line parser.y:620
		{
		}
	case 134:
		//line parser.y:621
		{
			RubyVAL.genericSlice = []ast.Node{RubyS[Rubypt-0].genericValue}
		}
	case 135:
		//line parser.y:622
		{
		}
	case 136:
		//line parser.y:623
		{
		}
	case 137:
		//line parser.y:624
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 138:
		//line parser.y:625
		{
		}
	case 139:
		//line parser.y:626
		{
		}
	case 140:
		//line parser.y:629
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 141:
		//line parser.y:632
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 142:
		//line parser.y:640
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 143:
		//line parser.y:642
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 144:
		//line parser.y:644
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				},
			})
		}
	case 145:
		//line parser.y:653
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}

			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Var:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
					Class: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				},
			})
		}
	}
	goto Rubystack /* stack new state and value */
}
