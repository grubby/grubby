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
const CARET = 57390
const LBRACKET = 57391
const RBRACKET = 57392
const LBRACE = 57393
const RBRACE = 57394
const DOLLARSIGN = 57395
const ATSIGN = 57396
const FILE_CONST_REF = 57397
const EOF = 57398

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

//line parser.y:793

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 25,
	43, 91,
	-2, 20,
	-1, 76,
	9, 63,
	10, 63,
	-2, 155,
	-1, 87,
	43, 91,
	-2, 20,
	-1, 92,
	8, 13,
	12, 13,
	14, 13,
	17, 13,
	18, 13,
	19, 13,
	23, 13,
	35, 13,
	36, 13,
	37, 13,
	38, 13,
	42, 13,
	-2, 11,
	-1, 101,
	43, 91,
	-2, 89,
	-1, 194,
	9, 63,
	10, 63,
	-2, 155,
	-1, 230,
	43, 92,
	-2, 90,
}

const RubyNprod = 164
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1206

var RubyAct = []int{

	9, 120, 162, 160, 78, 22, 159, 84, 151, 33,
	77, 186, 62, 2, 3, 114, 115, 125, 251, 64,
	152, 184, 252, 235, 182, 83, 75, 275, 4, 229,
	191, 126, 268, 186, 72, 232, 211, 97, 96, 60,
	59, 95, 166, 280, 110, 100, 102, 62, 62, 186,
	23, 86, 87, 76, 61, 81, 91, 122, 234, 63,
	250, 185, 123, 116, 183, 129, 130, 186, 133, 134,
	135, 136, 271, 138, 270, 54, 55, 122, 143, 80,
	145, 146, 123, 94, 104, 123, 73, 83, 243, 74,
	245, 244, 221, 282, 79, 178, 92, 5, 53, 52,
	186, 153, 186, 222, 144, 274, 164, 10, 103, 72,
	227, 256, 184, 209, 177, 62, 239, 209, 195, 184,
	149, 163, 273, 161, 180, 181, 189, 117, 118, 83,
	249, 230, 85, 192, 105, 106, 107, 108, 109, 193,
	101, 99, 258, 98, 163, 196, 205, 169, 148, 142,
	199, 177, 197, 128, 26, 177, 223, 201, 127, 266,
	228, 58, 132, 82, 124, 210, 177, 137, 177, 140,
	141, 177, 177, 131, 188, 119, 214, 208, 204, 88,
	139, 179, 212, 187, 124, 224, 1, 124, 147, 217,
	155, 156, 157, 158, 85, 122, 231, 112, 88, 168,
	123, 51, 50, 49, 177, 225, 226, 177, 172, 48,
	47, 88, 238, 46, 18, 17, 16, 15, 37, 88,
	88, 13, 88, 88, 88, 88, 12, 88, 11, 21,
	255, 88, 14, 19, 88, 88, 85, 31, 177, 30,
	32, 88, 177, 177, 29, 253, 0, 0, 177, 177,
	0, 0, 0, 0, 0, 259, 0, 200, 113, 260,
	0, 27, 0, 0, 177, 177, 177, 23, 86, 87,
	44, 0, 269, 177, 0, 0, 278, 177, 0, 0,
	88, 0, 0, 88, 0, 0, 89, 0, 0, 0,
	0, 0, 54, 55, 0, 279, 281, 0, 283, 88,
	284, 0, 124, 0, 88, 89, 0, 240, 28, 0,
	0, 56, 0, 57, 246, 53, 52, 0, 89, 0,
	0, 0, 0, 0, 254, 0, 89, 89, 257, 89,
	89, 89, 89, 90, 89, 0, 0, 0, 89, 88,
	262, 89, 89, 0, 0, 0, 0, 267, 89, 88,
	150, 154, 90, 0, 0, 0, 0, 0, 0, 165,
	0, 167, 0, 0, 0, 90, 0, 20, 170, 171,
	277, 0, 0, 90, 90, 0, 90, 90, 90, 90,
	0, 90, 0, 0, 88, 90, 0, 89, 90, 90,
	89, 0, 93, 0, 0, 90, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 89, 0, 0, 0,
	203, 89, 206, 23, 86, 87, 44, 0, 0, 91,
	0, 0, 0, 0, 93, 0, 23, 86, 87, 194,
	88, 0, 91, 93, 90, 0, 0, 90, 54, 55,
	93, 0, 0, 0, 93, 0, 89, 93, 93, 0,
	0, 54, 55, 90, 93, 0, 89, 56, 90, 92,
	0, 53, 52, 23, 86, 87, 76, 237, 0, 91,
	56, 241, 92, 242, 53, 52, 0, 0, 247, 0,
	0, 248, 0, 0, 0, 0, 0, 0, 54, 55,
	0, 89, 0, 90, 0, 0, 93, 0, 0, 0,
	0, 0, 0, 90, 263, 264, 0, 79, 265, 92,
	0, 53, 52, 0, 0, 0, 0, 93, 0, 0,
	0, 272, 0, 0, 0, 0, 0, 121, 86, 87,
	44, 0, 276, 23, 24, 25, 44, 89, 90, 0,
	34, 216, 42, 219, 218, 43, 35, 36, 0, 0,
	0, 45, 54, 55, 0, 0, 0, 0, 54, 55,
	0, 0, 93, 38, 39, 40, 41, 0, 0, 175,
	176, 56, 0, 57, 0, 53, 52, 56, 0, 57,
	0, 53, 52, 0, 90, 23, 24, 25, 44, 0,
	0, 0, 34, 0, 42, 0, 0, 43, 35, 36,
	0, 0, 0, 45, 0, 0, 0, 0, 0, 0,
	54, 55, 0, 0, 0, 38, 39, 40, 41, 0,
	0, 6, 7, 0, 0, 0, 0, 0, 0, 56,
	0, 57, 0, 53, 52, 0, 8, 23, 24, 25,
	44, 0, 0, 0, 34, 261, 42, 0, 0, 43,
	35, 36, 0, 0, 0, 45, 0, 0, 0, 0,
	0, 0, 54, 55, 0, 0, 0, 38, 39, 40,
	41, 0, 0, 175, 176, 23, 24, 25, 44, 0,
	0, 56, 34, 57, 42, 53, 52, 43, 35, 36,
	0, 0, 0, 45, 0, 0, 0, 0, 0, 0,
	54, 55, 0, 0, 0, 38, 39, 40, 41, 0,
	0, 175, 176, 0, 0, 0, 0, 0, 0, 56,
	0, 57, 236, 53, 52, 23, 24, 25, 44, 0,
	0, 0, 34, 233, 42, 0, 0, 43, 35, 36,
	0, 0, 0, 45, 0, 0, 0, 0, 0, 0,
	54, 55, 0, 0, 0, 38, 39, 40, 41, 0,
	0, 175, 176, 0, 0, 0, 0, 0, 0, 56,
	0, 57, 0, 53, 52, 23, 24, 25, 44, 0,
	0, 0, 34, 220, 42, 0, 0, 43, 35, 36,
	0, 0, 0, 45, 0, 0, 0, 0, 0, 0,
	54, 55, 0, 0, 0, 38, 39, 40, 41, 0,
	0, 175, 176, 0, 0, 0, 0, 0, 0, 56,
	0, 57, 0, 53, 52, 23, 24, 25, 44, 0,
	0, 0, 34, 215, 42, 0, 0, 43, 35, 36,
	0, 0, 0, 45, 0, 0, 0, 0, 0, 0,
	54, 55, 0, 0, 0, 38, 39, 40, 41, 0,
	0, 175, 176, 0, 0, 0, 0, 0, 0, 56,
	0, 57, 0, 53, 52, 23, 24, 25, 44, 0,
	0, 0, 34, 213, 42, 0, 0, 43, 35, 36,
	0, 0, 0, 45, 0, 0, 0, 0, 0, 0,
	54, 55, 0, 0, 0, 38, 39, 40, 41, 0,
	0, 175, 176, 23, 24, 25, 44, 0, 0, 56,
	34, 57, 42, 53, 52, 43, 35, 36, 0, 0,
	0, 45, 0, 0, 0, 0, 0, 0, 54, 55,
	0, 0, 0, 38, 39, 40, 41, 0, 0, 175,
	176, 0, 0, 0, 0, 0, 0, 56, 0, 57,
	207, 53, 52, 23, 24, 25, 44, 0, 0, 0,
	34, 202, 42, 0, 0, 43, 35, 36, 0, 0,
	0, 45, 0, 0, 0, 0, 0, 0, 54, 55,
	0, 0, 0, 38, 39, 40, 41, 0, 0, 175,
	176, 23, 24, 25, 44, 0, 0, 56, 34, 57,
	42, 53, 52, 43, 35, 36, 0, 0, 0, 45,
	0, 0, 0, 0, 0, 0, 54, 55, 0, 0,
	0, 38, 39, 40, 41, 0, 0, 175, 176, 23,
	24, 25, 44, 174, 0, 56, 34, 57, 42, 53,
	52, 43, 35, 36, 0, 0, 0, 45, 0, 0,
	0, 0, 0, 0, 54, 55, 0, 0, 0, 38,
	39, 40, 41, 0, 0, 0, 173, 23, 24, 25,
	44, 0, 0, 56, 34, 57, 42, 53, 52, 43,
	35, 36, 23, 111, 87, 45, 0, 0, 91, 0,
	0, 0, 54, 55, 0, 0, 0, 38, 39, 40,
	41, 0, 0, 0, 64, 0, 0, 54, 55, 0,
	0, 56, 0, 57, 0, 53, 52, 23, 190, 87,
	0, 0, 0, 23, 111, 87, 56, 0, 92, 0,
	53, 52, 0, 0, 0, 0, 0, 65, 66, 67,
	0, 64, 54, 55, 63, 70, 68, 69, 54, 55,
	198, 0, 64, 186, 71, 0, 0, 0, 0, 0,
	0, 56, 0, 57, 0, 53, 52, 56, 0, 57,
	0, 53, 52, 0, 65, 66, 67, 0, 0, 0,
	0, 63, 70, 68, 69, 65, 66, 67, 0, 0,
	0, 0, 63, 70, 68, 69,
}
var RubyPact = []int{

	-28, 580, -1000, -1000, -1000, -2, -1000, -1000, -1000, 1147,
	72, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-18, -1000, -1000, -1000, 45, 49, 7, 4, 3, -1000,
	-1000, -1000, -1000, -1000, 137, 133, 133, 74, 1072, 1072,
	1072, 1072, 1072, 1128, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 9, 121, -1000, -1000, 522, -1000, -12, -1000,
	-1000, -1000, 1072, 147, 1128, 262, 1072, 1128, 1128, 1128,
	1128, 1072, 262, 1072, 1072, 143, 262, -1000, 94, 522,
	262, 142, 110, 15, -1000, -1000, 458, -1000, -1000, -1000,
	-1000, -25, -25, -18, 1072, 1072, 1072, 1072, 115, 115,
	10, -1000, -1000, 1072, 141, 34, 34, 34, 34, 34,
	-1000, -1000, 1034, 996, -1000, -1000, 118, -1000, -1000, 14,
	11, -1000, 1158, -1000, -3, 1122, -13, 34, 421, -1000,
	15, -1000, 34, -1000, -1000, -1000, -1000, 34, 15, -1000,
	34, 34, -1000, 109, 1087, 1110, 15, -1000, -1000, 408,
	958, -1000, 140, -1000, 908, 34, 34, 34, 34, -1000,
	103, 138, -1000, 2, -1000, 870, 133, 820, 34, -1000,
	528, 770, 34, -1000, -1000, -1000, -1000, 1147, 34, 79,
	-1000, -1000, 151, -1000, 1128, -1000, -1000, -1000, 100, 156,
	-14, 124, 94, -1000, 262, -1000, -1000, -1000, 1, 15,
	-1000, -1000, -1000, 720, 13, -1000, 670, -1000, -8, 138,
	107, 1072, -8, -1000, -1000, -1000, -1000, 75, 1072, -1000,
	-1000, -1000, 123, -1000, -1000, 8, -30, -1000, 1072, 1128,
	-1000, 102, 1072, -1000, -1000, 136, -1000, 996, -1000, -1000,
	34, 996, 632, -1000, 1072, -1000, 34, 996, 996, 155,
	-1000, 1072, -1000, 26, 34, -1000, -1000, 34, -1000, 61,
	59, -1000, 34, 996, 996, 996, 116, 101, -16, -8,
	-1000, -1000, 996, -1000, 1072, 1128, 996, 33, 83, -8,
	-1000, -8, -1000, -8, -8,
}
var RubyPgo = []int{

	0, 95, 244, 240, 7, 239, 237, 367, 308, 233,
	232, 229, 0, 261, 107, 228, 5, 226, 154, 221,
	9, 2, 218, 217, 216, 215, 214, 213, 210, 209,
	203, 202, 201, 258, 197, 10, 8, 189, 186, 6,
	183, 181, 1, 178, 175, 174, 163, 4, 3, 161,
	17,
}
var RubyR1 = []int{

	0, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 50, 50, 33, 33, 33, 33, 33, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 16,
	16, 16, 16, 16, 16, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 35,
	35, 44, 44, 42, 42, 42, 42, 42, 47, 47,
	47, 47, 46, 46, 46, 46, 46, 15, 15, 39,
	39, 21, 21, 48, 48, 48, 17, 17, 19, 20,
	20, 49, 49, 10, 10, 10, 10, 10, 10, 10,
	8, 8, 18, 18, 13, 13, 22, 22, 23, 24,
	25, 26, 27, 27, 27, 27, 28, 29, 30, 31,
	32, 2, 5, 6, 6, 3, 3, 40, 40, 40,
	40, 45, 45, 45, 45, 4, 4, 4, 4, 36,
	43, 43, 43, 9, 9, 9, 9, 9, 9, 9,
	37, 37, 37, 37, 37, 34, 34, 34, 7, 11,
	41, 41, 41, 41,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 4,
	2, 3, 4, 4, 3, 2, 3, 4, 6, 3,
	1, 1, 3, 0, 1, 1, 1, 3, 1, 1,
	3, 3, 1, 1, 3, 3, 3, 7, 7, 1,
	3, 1, 3, 0, 1, 3, 4, 6, 4, 1,
	4, 1, 4, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 3, 3, 3, 3, 2, 2,
	2, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 1, 1, 3, 3, 5, 5, 0, 4, 7,
	8, 3, 3, 7, 8, 3, 4, 4, 3, 3,
	0, 1, 3, 4, 5, 3, 3, 3, 3, 4,
	0, 4, 3, 3, 2, 0, 2, 2, 3, 4,
	0, 3, 4, 6,
}
var RubyChk = []int{

	-1000, -38, 41, 42, 56, -1, 41, 42, 56, -12,
	-14, -15, -17, -19, -10, -23, -24, -25, -26, -9,
	-7, -11, -16, 5, 6, 7, -18, -13, -8, -2,
	-5, -6, -3, -20, 12, 18, 19, -22, 35, 36,
	37, 38, 14, 17, 8, 23, -27, -28, -29, -30,
	-31, -32, 54, 53, 30, 31, 49, 51, -49, 42,
	41, 56, 14, 44, 4, 37, 38, 39, 46, 47,
	45, 17, 37, 14, 17, 44, 8, -35, -47, 49,
	34, 10, -46, -12, -4, -14, 6, 7, -18, -13,
	-8, 11, 51, -7, 34, 34, 34, 34, 6, 4,
	-20, 7, -20, 34, 10, -1, -1, -1, -1, -1,
	-12, 6, -34, -33, 6, 7, 54, 6, 7, -44,
	-42, 5, -12, -16, -14, -50, 43, -1, 6, -12,
	-12, -14, -1, -12, -12, -12, -12, -1, -12, -14,
	-1, -1, 6, -42, 10, -12, -12, -14, 6, 10,
	-33, -36, 45, -36, -33, -1, -1, -1, -1, -39,
	-48, 8, -21, 6, -39, -33, 32, -33, -1, 6,
	-33, -33, -1, 42, 9, 41, 42, -12, -1, -41,
	6, 7, 10, 50, 10, 50, 41, -40, -45, -12,
	6, 43, -47, -35, 8, 9, -12, -4, 50, -12,
	-14, -4, 13, -33, -43, 6, -33, 52, -50, 10,
	-48, 34, -50, 13, -20, 13, 13, -37, 16, 15,
	13, 13, 24, 5, -12, -50, -50, 10, 4, 43,
	7, -42, 34, 13, 45, 10, 52, -33, -21, 9,
	-1, -33, -33, 13, 16, 15, -1, -33, -33, 7,
	52, 10, 52, -50, -1, -12, 9, -1, 6, -50,
	-50, 13, -1, -33, -33, -33, 4, -1, 6, -50,
	13, 13, -33, 6, 4, 43, -33, -1, -12, -50,
	10, -50, 10, -50, -50,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
	46, 47, 48, 18, 19, -2, 21, 22, 23, 24,
	25, 26, 27, 28, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 155, 13, 29, 30, 31, 32,
	33, 34, 0, 0, 121, 122, 63, 11, 0, 5,
	6, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, -2, 50, 55, 63,
	0, 0, 60, 68, 69, 73, 19, -2, 21, 22,
	23, 13, -2, 0, 0, 0, 0, 0, 83, 83,
	13, -2, 13, 0, 0, 108, 109, 110, 111, 13,
	13, 19, 0, 160, 102, 103, 0, 100, 101, 0,
	0, 18, 64, 65, 66, 127, 0, 145, 51, 56,
	112, 114, 116, 117, 118, 119, 120, 147, 113, 115,
	146, 148, 54, 0, 0, 64, 93, 94, 106, 0,
	0, 13, 140, 13, 0, 95, 96, 97, 98, 11,
	79, 83, 84, 81, 11, 0, 0, 0, 99, 107,
	0, 0, 156, 157, 158, 14, 15, 16, 17, 0,
	104, 105, 0, 123, 0, 124, 12, 11, 11, 0,
	19, 0, 52, 53, -2, 49, 70, 71, 57, 74,
	75, 76, 135, 0, 0, 141, 0, 138, 13, 0,
	0, 0, 13, 86, 13, 88, 143, 0, 0, 13,
	149, 159, 13, 62, 67, 0, 0, 11, 0, 0,
	-2, 0, 0, 136, 139, 0, 137, 11, 85, 80,
	82, 11, 0, 144, 0, 13, 13, 154, 161, 13,
	125, 0, 126, 0, 11, 131, 59, 58, 142, 0,
	0, 87, 13, 152, 153, 162, 0, 0, 0, 128,
	77, 78, 151, 13, 0, 0, 163, 11, 11, 129,
	11, 133, 11, 130, 134,
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
	52, 53, 54, 55, 56,
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
		//line parser.y:158
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:160
		{
		}
	case 3:
		//line parser.y:162
		{
		}
	case 4:
		//line parser.y:164
		{
		}
	case 5:
		//line parser.y:166
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:168
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:170
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:176
		{
		}
	case 11:
		//line parser.y:178
		{
		}
	case 12:
		//line parser.y:179
		{
		}
	case 13:
		//line parser.y:182
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 14:
		//line parser.y:184
		{
		}
	case 15:
		//line parser.y:186
		{
		}
	case 16:
		//line parser.y:188
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 17:
		//line parser.y:190
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
		//line parser.y:200
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 50:
		//line parser.y:207
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 51:
		//line parser.y:214
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 52:
		//line parser.y:221
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 53:
		//line parser.y:229
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 54:
		//line parser.y:237
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 55:
		//line parser.y:247
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 56:
		//line parser.y:254
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 57:
		//line parser.y:264
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 58:
		//line parser.y:274
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 59:
		//line parser.y:283
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 60:
		//line parser.y:285
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 61:
		//line parser.y:288
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 62:
		//line parser.y:290
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 63:
		//line parser.y:292
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 64:
		//line parser.y:294
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 65:
		//line parser.y:296
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 66:
		//line parser.y:298
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:300
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:304
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 69:
		//line parser.y:306
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 70:
		//line parser.y:308
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 71:
		//line parser.y:310
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 72:
		//line parser.y:313
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:315
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:317
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 75:
		//line parser.y:319
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:321
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:327
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 78:
		//line parser.y:335
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 79:
		//line parser.y:344
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 80:
		//line parser.y:346
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 81:
		//line parser.y:349
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 82:
		//line parser.y:351
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 83:
		//line parser.y:353
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 84:
		//line parser.y:355
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:359
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 86:
		//line parser.y:364
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 87:
		//line parser.y:371
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 88:
		//line parser.y:381
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 89:
		//line parser.y:390
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 90:
		//line parser.y:396
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 91:
		//line parser.y:404
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 92:
		//line parser.y:408
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 93:
		//line parser.y:413
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 94:
		//line parser.y:420
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 95:
		//line parser.y:427
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 96:
		//line parser.y:434
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 97:
		//line parser.y:441
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 98:
		//line parser.y:448
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 99:
		//line parser.y:455
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 100:
		//line parser.y:463
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 101:
		//line parser.y:465
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 102:
		//line parser.y:468
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 103:
		//line parser.y:470
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 104:
		//line parser.y:473
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 105:
		//line parser.y:475
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 106:
		//line parser.y:478
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 107:
		//line parser.y:480
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 108:
		//line parser.y:482
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 109:
		//line parser.y:483
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 110:
		//line parser.y:484
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 111:
		//line parser.y:485
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 112:
		//line parser.y:488
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 113:
		//line parser.y:496
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 114:
		//line parser.y:504
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 115:
		//line parser.y:512
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 116:
		//line parser.y:521
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 117:
		//line parser.y:530
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 118:
		//line parser.y:539
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 119:
		//line parser.y:548
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 120:
		//line parser.y:557
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 121:
		//line parser.y:565
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 122:
		//line parser.y:566
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 123:
		//line parser.y:568
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 124:
		//line parser.y:570
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 125:
		//line parser.y:573
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 126:
		//line parser.y:581
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 127:
		//line parser.y:589
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 128:
		//line parser.y:591
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 129:
		//line parser.y:598
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 130:
		//line parser.y:605
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 131:
		//line parser.y:613
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 132:
		//line parser.y:620
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 133:
		//line parser.y:627
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 134:
		//line parser.y:634
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 135:
		//line parser.y:642
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 136:
		//line parser.y:644
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 137:
		//line parser.y:651
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 138:
		//line parser.y:658
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 139:
		//line parser.y:661
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 140:
		//line parser.y:663
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 141:
		//line parser.y:665
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 142:
		//line parser.y:667
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 143:
		//line parser.y:670
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 144:
		//line parser.y:677
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 145:
		//line parser.y:685
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 146:
		//line parser.y:692
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 147:
		//line parser.y:699
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 148:
		//line parser.y:706
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 149:
		//line parser.y:713
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-2].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 150:
		//line parser.y:720
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 151:
		//line parser.y:722
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 152:
		//line parser.y:729
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 153:
		//line parser.y:736
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 154:
		//line parser.y:743
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 155:
		//line parser.y:750
		{
		}
	case 156:
		//line parser.y:751
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 157:
		//line parser.y:752
		{
		}
	case 158:
		//line parser.y:755
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 159:
		//line parser.y:758
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 160:
		//line parser.y:766
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 161:
		//line parser.y:768
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 162:
		//line parser.y:770
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 163:
		//line parser.y:779
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
	}
	goto Rubystack /* stack new state and value */
}
