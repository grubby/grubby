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

//line parser.y:846

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 26,
	44, 101,
	-2, 20,
	-1, 85,
	9, 71,
	10, 71,
	-2, 163,
	-1, 96,
	44, 101,
	-2, 20,
	-1, 101,
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
	-1, 112,
	44, 101,
	-2, 99,
	-1, 123,
	44, 101,
	-2, 20,
	-1, 216,
	9, 71,
	10, 71,
	-2, 163,
	-1, 254,
	44, 102,
	-2, 100,
}

const RubyNprod = 173
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1298

var RubyAct = []int{

	9, 182, 180, 34, 93, 22, 206, 64, 87, 94,
	86, 204, 179, 170, 2, 3, 208, 132, 84, 68,
	259, 171, 126, 127, 299, 253, 92, 276, 213, 4,
	137, 208, 256, 115, 233, 62, 61, 108, 107, 27,
	111, 113, 106, 183, 275, 121, 292, 207, 66, 67,
	63, 188, 205, 69, 70, 71, 258, 10, 114, 133,
	65, 74, 72, 73, 134, 97, 220, 140, 141, 142,
	143, 128, 146, 147, 148, 149, 208, 184, 152, 153,
	154, 155, 208, 102, 97, 185, 162, 274, 83, 164,
	166, 134, 295, 245, 134, 294, 92, 304, 97, 165,
	183, 64, 181, 161, 246, 162, 97, 97, 97, 97,
	134, 97, 97, 97, 97, 172, 135, 97, 97, 97,
	97, 208, 174, 186, 208, 97, 199, 144, 97, 97,
	64, 79, 104, 306, 184, 97, 163, 211, 156, 251,
	92, 231, 185, 135, 97, 298, 135, 102, 214, 267,
	215, 269, 268, 102, 168, 64, 104, 68, 273, 105,
	77, 78, 135, 254, 218, 80, 136, 112, 219, 166,
	199, 297, 76, 222, 199, 282, 97, 60, 221, 97,
	280, 206, 263, 231, 232, 235, 66, 67, 199, 79,
	199, 79, 238, 199, 199, 229, 206, 102, 65, 81,
	217, 206, 82, 97, 110, 247, 109, 248, 97, 234,
	24, 95, 96, 85, 226, 90, 100, 162, 77, 78,
	77, 78, 134, 80, 191, 199, 102, 167, 199, 160,
	76, 151, 76, 262, 255, 28, 56, 57, 202, 203,
	89, 68, 129, 130, 97, 139, 97, 290, 252, 91,
	210, 131, 225, 201, 279, 88, 97, 101, 209, 55,
	54, 98, 199, 1, 241, 124, 199, 199, 52, 51,
	66, 67, 199, 199, 135, 69, 70, 71, 50, 49,
	98, 48, 65, 74, 72, 73, 47, 18, 199, 199,
	199, 17, 16, 97, 98, 15, 38, 199, 13, 12,
	302, 199, 98, 98, 98, 98, 23, 98, 98, 98,
	98, 29, 11, 98, 98, 98, 98, 24, 95, 96,
	45, 98, 21, 100, 98, 98, 14, 19, 32, 31,
	33, 98, 30, 0, 0, 0, 0, 99, 0, 97,
	98, 0, 125, 56, 57, 0, 230, 0, 0, 0,
	0, 0, 0, 236, 0, 0, 99, 0, 0, 0,
	0, 0, 58, 0, 101, 0, 55, 54, 0, 0,
	99, 0, 98, 0, 0, 98, 249, 250, 99, 99,
	99, 99, 0, 99, 99, 99, 99, 0, 0, 99,
	99, 99, 99, 0, 200, 0, 5, 99, 0, 98,
	99, 99, 0, 0, 98, 0, 0, 99, 0, 0,
	0, 0, 0, 0, 0, 0, 99, 0, 277, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 283, 0,
	0, 0, 284, 0, 116, 117, 118, 119, 120, 0,
	98, 0, 98, 169, 173, 293, 0, 0, 99, 0,
	0, 99, 98, 0, 187, 20, 189, 0, 0, 138,
	0, 0, 0, 192, 193, 145, 0, 0, 303, 305,
	150, 307, 0, 308, 0, 99, 157, 158, 159, 0,
	99, 103, 0, 0, 24, 95, 96, 216, 0, 98,
	100, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	175, 176, 177, 178, 0, 0, 0, 0, 0, 190,
	56, 57, 0, 224, 103, 227, 99, 0, 99, 194,
	0, 24, 95, 96, 85, 103, 0, 100, 99, 58,
	0, 101, 0, 55, 54, 98, 103, 0, 0, 0,
	0, 103, 0, 0, 103, 103, 0, 56, 57, 0,
	0, 103, 0, 0, 0, 0, 24, 95, 96, 45,
	103, 0, 0, 0, 0, 99, 88, 0, 101, 0,
	55, 54, 0, 261, 0, 0, 0, 0, 0, 265,
	0, 266, 56, 57, 0, 0, 271, 0, 0, 272,
	0, 0, 0, 0, 0, 103, 24, 122, 123, 0,
	0, 58, 100, 59, 0, 55, 54, 24, 212, 123,
	0, 99, 287, 288, 0, 0, 289, 0, 0, 0,
	0, 0, 56, 57, 103, 0, 0, 0, 264, 296,
	0, 0, 0, 56, 57, 0, 0, 270, 0, 0,
	300, 58, 0, 101, 208, 55, 54, 278, 0, 0,
	0, 281, 58, 0, 59, 0, 55, 54, 0, 0,
	0, 0, 0, 286, 0, 0, 0, 0, 0, 0,
	291, 0, 103, 24, 25, 26, 45, 0, 0, 0,
	35, 240, 43, 243, 242, 44, 36, 37, 0, 0,
	0, 46, 0, 301, 0, 0, 0, 0, 53, 56,
	57, 0, 0, 0, 39, 40, 41, 42, 0, 0,
	197, 198, 24, 25, 26, 45, 0, 0, 58, 35,
	59, 43, 55, 54, 44, 36, 37, 0, 0, 0,
	46, 0, 0, 0, 0, 0, 0, 53, 56, 57,
	0, 0, 0, 39, 40, 41, 42, 0, 0, 6,
	7, 0, 0, 0, 0, 0, 0, 58, 0, 59,
	0, 55, 54, 0, 8, 24, 25, 26, 45, 0,
	0, 0, 35, 285, 43, 0, 0, 44, 36, 37,
	0, 0, 0, 46, 0, 0, 0, 0, 0, 0,
	53, 56, 57, 0, 0, 0, 39, 40, 41, 42,
	0, 0, 197, 198, 24, 25, 26, 45, 0, 0,
	58, 35, 59, 43, 55, 54, 44, 36, 37, 0,
	0, 0, 46, 0, 0, 0, 0, 0, 0, 53,
	56, 57, 0, 0, 0, 39, 40, 41, 42, 0,
	0, 197, 198, 0, 0, 0, 0, 0, 0, 58,
	0, 59, 260, 55, 54, 24, 25, 26, 45, 0,
	0, 0, 35, 257, 43, 0, 0, 44, 36, 37,
	0, 0, 0, 46, 0, 0, 0, 0, 0, 0,
	53, 56, 57, 0, 0, 0, 39, 40, 41, 42,
	0, 0, 197, 198, 0, 0, 0, 0, 0, 0,
	58, 0, 59, 0, 55, 54, 24, 25, 26, 45,
	0, 0, 0, 35, 244, 43, 0, 0, 44, 36,
	37, 0, 0, 0, 46, 0, 0, 0, 0, 0,
	0, 53, 56, 57, 0, 0, 0, 39, 40, 41,
	42, 0, 0, 197, 198, 0, 0, 0, 0, 0,
	0, 58, 0, 59, 0, 55, 54, 24, 25, 26,
	45, 0, 0, 0, 35, 239, 43, 0, 0, 44,
	36, 37, 0, 0, 0, 46, 0, 0, 0, 0,
	0, 0, 53, 56, 57, 0, 0, 0, 39, 40,
	41, 42, 0, 0, 197, 198, 0, 0, 0, 0,
	0, 0, 58, 0, 59, 0, 55, 54, 24, 25,
	26, 45, 0, 0, 0, 35, 237, 43, 0, 0,
	44, 36, 37, 0, 0, 0, 46, 0, 0, 0,
	0, 0, 0, 53, 56, 57, 0, 0, 0, 39,
	40, 41, 42, 0, 0, 197, 198, 24, 25, 26,
	45, 0, 0, 58, 35, 59, 43, 55, 54, 44,
	36, 37, 0, 0, 0, 46, 0, 0, 0, 0,
	0, 0, 53, 56, 57, 0, 0, 0, 39, 40,
	41, 42, 0, 0, 197, 198, 0, 0, 0, 0,
	0, 0, 58, 0, 59, 228, 55, 54, 24, 25,
	26, 45, 0, 0, 0, 35, 223, 43, 0, 0,
	44, 36, 37, 0, 0, 0, 46, 0, 0, 0,
	0, 0, 0, 53, 56, 57, 0, 0, 0, 39,
	40, 41, 42, 0, 0, 197, 198, 24, 25, 26,
	45, 0, 0, 58, 35, 59, 43, 55, 54, 44,
	36, 37, 0, 0, 0, 46, 0, 0, 0, 0,
	0, 0, 53, 56, 57, 0, 0, 0, 39, 40,
	41, 42, 0, 0, 197, 198, 24, 25, 26, 45,
	196, 0, 58, 35, 59, 43, 55, 54, 44, 36,
	37, 0, 0, 0, 46, 0, 0, 0, 0, 0,
	0, 53, 56, 57, 0, 0, 0, 39, 40, 41,
	42, 0, 0, 0, 195, 24, 25, 26, 45, 0,
	0, 58, 35, 59, 43, 55, 54, 44, 36, 37,
	24, 122, 123, 46, 0, 0, 0, 0, 0, 0,
	53, 56, 57, 0, 0, 0, 39, 40, 41, 42,
	0, 0, 0, 68, 0, 0, 56, 57, 0, 0,
	58, 0, 59, 0, 55, 54, 75, 0, 0, 0,
	0, 0, 0, 0, 0, 58, 0, 59, 0, 55,
	54, 0, 66, 67, 0, 0, 0, 69, 70, 71,
	0, 0, 0, 0, 65, 74, 72, 73,
}
var RubyPact = []int{

	-28, 707, -1000, -1000, -1000, -7, -1000, -1000, -1000, 1249,
	185, -1000, -1000, -1000, 71, -1000, -1000, -1000, -1000, -1000,
	-27, -1000, -1000, -1000, -1000, 205, 124, 7, 3, 2,
	-1000, -1000, -1000, -1000, -1000, 200, 160, 160, 23, 1210,
	1210, 1210, 1210, 1210, 1225, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 16, 236, -1000, -1000, 551, -1000,
	-14, -1000, -1000, -1000, 1210, 239, 1225, 1225, 1225, 551,
	1210, 1225, 1225, 1225, 1225, 1210, 225, 1225, 1225, 1225,
	551, 1210, 1210, 1210, 223, 551, -1000, 126, 551, 551,
	221, 144, 153, -1000, -1000, 516, 148, -1000, -1000, -1000,
	-25, -25, 187, -27, 551, 1210, 1210, 1210, 1210, 94,
	94, 18, -1000, -1000, 1210, 218, 116, 116, 116, 116,
	116, -1000, -1000, -1000, 1171, 1132, -1000, -1000, 232, -1000,
	-1000, 1, -4, 237, -1000, 127, 602, -16, 116, 479,
	-1000, -1000, -1000, 153, 187, 116, -1000, -1000, -1000, -1000,
	116, -1000, -1000, -1000, -1000, 153, 187, 116, 116, 116,
	-1000, 191, 237, 591, 15, -1000, 153, -1000, 312, 1093,
	-1000, 208, -1000, 1042, 186, 116, 116, 116, 116, -1000,
	131, 37, -1000, -1, 203, 179, -1000, 1003, 160, 952,
	116, -1000, 668, 901, 116, -1000, -1000, -1000, -1000, 1249,
	116, 80, -1000, -1000, 1225, -1000, 1225, -1000, -1000, -1000,
	129, 244, -19, 156, 126, -1000, 551, -1000, -1000, -1000,
	-3, -1000, -1000, -1000, 850, 10, -1000, 799, -1000, -1000,
	-11, 37, 173, 1210, -1000, -1000, -11, -1000, -1000, -1000,
	-1000, 136, 1210, -1000, -1000, -1000, 151, -1000, -1000, 34,
	-26, -1000, 1210, 1225, -1000, 171, 1210, -1000, -1000, 169,
	-1000, 1132, -1000, -1000, 116, 1132, 760, -1000, 1210, -1000,
	116, 1132, 1132, 243, -1000, 1210, -1000, 40, 116, -1000,
	-1000, 116, -1000, 82, 79, -1000, 116, 1132, 1132, 1132,
	165, 141, -20, -11, -1000, -1000, 1132, -1000, 1210, 1225,
	1132, 87, 123, -11, -1000, -11, -1000, -11, -11,
}
var RubyPgo = []int{

	0, 394, 332, 330, 4, 329, 328, 455, 311, 327,
	326, 322, 0, 235, 57, 312, 306, 5, 299, 1,
	39, 298, 3, 9, 296, 295, 292, 291, 287, 286,
	281, 279, 278, 269, 268, 342, 265, 10, 13, 264,
	263, 12, 258, 253, 17, 252, 251, 250, 249, 8,
	2, 177, 166,
}
var RubyR1 = []int{

	0, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 52, 52, 35, 35, 35, 35, 35, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 17,
	17, 17, 17, 17, 17, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 37, 37, 46,
	46, 44, 44, 44, 44, 44, 49, 49, 49, 49,
	23, 23, 48, 48, 48, 15, 15, 41, 41, 19,
	19, 19, 19, 50, 50, 50, 18, 18, 21, 22,
	22, 51, 51, 10, 10, 10, 10, 10, 10, 8,
	8, 20, 20, 13, 13, 24, 24, 25, 26, 27,
	28, 29, 29, 29, 29, 30, 31, 32, 33, 34,
	2, 5, 6, 6, 3, 3, 42, 42, 42, 42,
	47, 47, 47, 4, 4, 4, 4, 38, 45, 45,
	45, 9, 9, 9, 9, 9, 9, 9, 9, 39,
	39, 39, 39, 36, 36, 36, 7, 11, 43, 43,
	43, 43, 16,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	4, 4, 2, 3, 3, 4, 4, 3, 2, 3,
	3, 3, 3, 3, 3, 4, 6, 3, 1, 1,
	3, 0, 1, 1, 1, 3, 1, 1, 3, 3,
	1, 1, 1, 3, 3, 7, 7, 1, 3, 1,
	2, 3, 2, 0, 1, 3, 4, 6, 4, 1,
	4, 1, 4, 3, 3, 3, 3, 3, 3, 2,
	2, 2, 2, 3, 3, 3, 3, 2, 2, 2,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	1, 1, 3, 3, 5, 5, 0, 4, 7, 8,
	3, 7, 8, 3, 4, 4, 3, 3, 0, 1,
	3, 4, 5, 3, 3, 3, 3, 3, 4, 4,
	3, 3, 2, 0, 2, 2, 3, 4, 0, 3,
	4, 6, 1,
}
var RubyChk = []int{

	-1000, -40, 42, 43, 57, -1, 42, 43, 57, -12,
	-14, -15, -18, -21, -10, -25, -26, -27, -28, -9,
	-7, -11, -17, -16, 5, 6, 7, -20, -13, -8,
	-2, -5, -6, -3, -22, 12, 18, 19, -24, 36,
	37, 38, 39, 14, 17, 8, 23, -29, -30, -31,
	-32, -33, -34, 30, 55, 54, 31, 32, 50, 52,
	-51, 43, 42, 57, 14, 45, 33, 34, 4, 38,
	39, 40, 47, 48, 46, 17, 45, 33, 34, 4,
	38, 14, 17, 17, 45, 8, -37, -49, 50, 35,
	10, -48, -12, -4, -23, 6, 7, -20, -13, -8,
	11, 52, -14, -7, 8, 35, 35, 35, 35, 6,
	4, -22, 7, -22, 35, 10, -1, -1, -1, -1,
	-1, -12, 6, 7, -36, -35, 6, 7, 55, 6,
	7, -46, -44, -12, -17, -14, -52, 44, -1, 6,
	-12, -12, -12, -12, -14, -1, -12, -12, -12, -12,
	-1, 6, -12, -12, -12, -12, -14, -1, -1, -1,
	6, -44, -12, 10, -12, -23, -12, 6, 10, -35,
	-38, 46, -38, -35, -44, -1, -1, -1, -1, -41,
	-50, 8, -19, 6, 40, 48, -41, -35, 33, -35,
	-1, 6, -35, -35, -1, 43, 9, 42, 43, -12,
	-1, -43, 6, 7, 10, 51, 10, 51, 42, -42,
	-47, -12, 6, 44, -49, -37, 8, 9, -12, -4,
	51, -23, -4, 13, -35, -45, 6, -35, 53, 9,
	-52, 10, -50, 35, 6, 6, -52, 13, -22, 13,
	13, -39, 16, 15, 13, 13, 24, -12, -12, -52,
	-52, 10, 4, 44, 7, -44, 35, 13, 46, 10,
	53, -35, -19, 9, -1, -35, -35, 13, 16, 15,
	-1, -35, -35, 7, 53, 10, 53, -52, -1, -12,
	9, -1, 6, -52, -52, 13, -1, -35, -35, -35,
	4, -1, 6, -52, 13, 13, -35, 6, 4, 44,
	-35, -1, -12, -52, 10, -52, 10, -52, -52,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
	46, 47, 48, 49, 18, 19, -2, 21, 22, 23,
	24, 25, 26, 27, 28, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 163, 13, 29, 30, 31,
	32, 33, 34, 172, 0, 0, 130, 131, 71, 11,
	0, 5, 6, 7, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, -2, 52, 58, 71, 0,
	0, 68, 76, 77, 82, 19, -2, 21, 22, 23,
	13, -2, 81, 0, 71, 0, 0, 0, 0, 93,
	93, 13, -2, 13, 0, 0, 117, 118, 119, 120,
	13, 13, 19, -2, 0, 168, 111, 112, 0, 109,
	110, 0, 0, 69, 73, 74, 136, 0, 153, 53,
	59, 61, 63, 121, 123, 125, 126, 127, 128, 129,
	155, 54, 60, 62, 64, 122, 124, 154, 156, 157,
	57, 0, 72, 0, 69, 103, 80, 115, 0, 0,
	13, 148, 13, 0, 0, 104, 105, 106, 107, 11,
	87, 93, 94, 89, 0, 0, 11, 0, 0, 0,
	108, 116, 0, 0, 164, 165, 166, 14, 15, 16,
	17, 0, 113, 114, 0, 132, 0, 133, 12, 11,
	11, 0, 19, 0, 55, 56, -2, 50, 78, 79,
	65, 83, 84, 143, 0, 0, 149, 0, 146, 51,
	13, 0, 0, 0, 90, 92, 13, 96, 13, 98,
	151, 0, 0, 13, 158, 167, 13, 70, 75, 0,
	0, 11, 0, 0, -2, 0, 0, 144, 147, 0,
	145, 11, 95, 88, 91, 11, 0, 152, 0, 13,
	13, 162, 169, 13, 134, 0, 135, 0, 11, 140,
	67, 66, 150, 0, 0, 97, 13, 160, 161, 170,
	0, 0, 0, 137, 85, 86, 159, 13, 0, 0,
	171, 11, 11, 138, 11, 141, 11, 139, 142,
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
		//line parser.y:161
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:163
		{
		}
	case 3:
		//line parser.y:165
		{
		}
	case 4:
		//line parser.y:167
		{
		}
	case 5:
		//line parser.y:169
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:171
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:173
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:179
		{
		}
	case 11:
		//line parser.y:181
		{
		}
	case 12:
		//line parser.y:182
		{
		}
	case 13:
		//line parser.y:185
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 14:
		//line parser.y:187
		{
		}
	case 15:
		//line parser.y:189
		{
		}
	case 16:
		//line parser.y:191
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 17:
		//line parser.y:193
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
		//line parser.y:203
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 51:
		//line parser.y:210
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 52:
		//line parser.y:217
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 53:
		//line parser.y:224
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 54:
		//line parser.y:231
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 55:
		//line parser.y:238
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 56:
		//line parser.y:246
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 57:
		//line parser.y:254
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 58:
		//line parser.y:264
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 59:
		//line parser.y:271
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 60:
		//line parser.y:279
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 61:
		//line parser.y:287
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 62:
		//line parser.y:295
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 63:
		//line parser.y:303
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 64:
		//line parser.y:311
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 65:
		//line parser.y:321
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 66:
		//line parser.y:331
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 67:
		//line parser.y:340
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 68:
		//line parser.y:342
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 69:
		//line parser.y:345
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 70:
		//line parser.y:347
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 71:
		//line parser.y:349
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 72:
		//line parser.y:351
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:353
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:355
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 75:
		//line parser.y:357
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:361
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:363
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 78:
		//line parser.y:365
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 79:
		//line parser.y:367
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 80:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 81:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 82:
		//line parser.y:372
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 83:
		//line parser.y:374
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 84:
		//line parser.y:376
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:382
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 86:
		//line parser.y:390
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 87:
		//line parser.y:399
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 88:
		//line parser.y:401
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 89:
		//line parser.y:404
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 90:
		//line parser.y:406
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 91:
		//line parser.y:408
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 92:
		//line parser.y:410
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 93:
		//line parser.y:412
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 94:
		//line parser.y:414
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 95:
		//line parser.y:418
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 96:
		//line parser.y:423
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 97:
		//line parser.y:430
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 98:
		//line parser.y:440
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 99:
		//line parser.y:449
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 100:
		//line parser.y:455
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 101:
		//line parser.y:463
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 102:
		//line parser.y:467
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 103:
		//line parser.y:472
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
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
		//line parser.y:515
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 110:
		//line parser.y:517
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 111:
		//line parser.y:520
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 112:
		//line parser.y:522
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 113:
		//line parser.y:525
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 114:
		//line parser.y:527
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 115:
		//line parser.y:530
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 116:
		//line parser.y:532
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 117:
		//line parser.y:534
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 118:
		//line parser.y:535
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 119:
		//line parser.y:536
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 120:
		//line parser.y:537
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 121:
		//line parser.y:540
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 122:
		//line parser.y:548
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
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 124:
		//line parser.y:564
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 125:
		//line parser.y:573
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 126:
		//line parser.y:582
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 127:
		//line parser.y:591
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 128:
		//line parser.y:600
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 129:
		//line parser.y:609
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 130:
		//line parser.y:617
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 131:
		//line parser.y:618
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 132:
		//line parser.y:620
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 133:
		//line parser.y:622
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 134:
		//line parser.y:625
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 135:
		//line parser.y:633
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 136:
		//line parser.y:641
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 137:
		//line parser.y:643
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 138:
		//line parser.y:650
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 139:
		//line parser.y:657
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 140:
		//line parser.y:665
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 141:
		//line parser.y:672
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 142:
		//line parser.y:679
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 143:
		//line parser.y:687
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 144:
		//line parser.y:689
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 145:
		//line parser.y:696
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 146:
		//line parser.y:703
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 147:
		//line parser.y:706
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 148:
		//line parser.y:708
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 149:
		//line parser.y:710
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 150:
		//line parser.y:712
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 151:
		//line parser.y:715
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 152:
		//line parser.y:722
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 153:
		//line parser.y:730
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 154:
		//line parser.y:737
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 155:
		//line parser.y:744
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 156:
		//line parser.y:751
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 157:
		//line parser.y:758
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 158:
		//line parser.y:765
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-2].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 159:
		//line parser.y:773
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 160:
		//line parser.y:780
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 161:
		//line parser.y:787
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 162:
		//line parser.y:794
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 163:
		//line parser.y:801
		{
		}
	case 164:
		//line parser.y:802
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 165:
		//line parser.y:803
		{
		}
	case 166:
		//line parser.y:806
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 167:
		//line parser.y:809
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 168:
		//line parser.y:817
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 169:
		//line parser.y:819
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 170:
		//line parser.y:821
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 171:
		//line parser.y:830
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	}
	goto Rubystack /* stack new state and value */
}
