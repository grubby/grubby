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

//line parser.y:803

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 13,
	43, 92,
	-2, 20,
	-1, 81,
	43, 92,
	-2, 20,
	-1, 86,
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
	-2, 9,
	-1, 109,
	43, 92,
	-2, 90,
	-1, 214,
	43, 93,
	-2, 91,
}

const RubyNprod = 167
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1251

var RubyAct = []int{

	114, 173, 9, 72, 171, 34, 71, 78, 188, 100,
	10, 2, 3, 77, 164, 233, 146, 147, 267, 213,
	169, 94, 95, 106, 162, 237, 4, 164, 91, 112,
	160, 230, 218, 90, 89, 88, 87, 227, 176, 69,
	102, 277, 236, 61, 139, 103, 216, 22, 228, 21,
	232, 257, 119, 111, 268, 113, 115, 116, 117, 118,
	79, 123, 126, 127, 163, 130, 131, 132, 133, 96,
	161, 137, 129, 164, 140, 141, 103, 134, 135, 103,
	138, 77, 164, 60, 235, 92, 164, 104, 93, 211,
	266, 14, 154, 108, 110, 144, 150, 151, 152, 153,
	69, 156, 157, 148, 82, 174, 167, 172, 91, 246,
	128, 248, 247, 275, 253, 162, 252, 69, 104, 109,
	178, 104, 142, 243, 216, 187, 77, 105, 79, 190,
	182, 82, 191, 194, 162, 158, 159, 97, 98, 155,
	195, 214, 270, 82, 255, 198, 187, 196, 174, 204,
	187, 179, 200, 82, 82, 143, 82, 82, 82, 82,
	125, 107, 82, 208, 207, 82, 82, 265, 212, 41,
	76, 166, 82, 79, 99, 203, 187, 217, 187, 189,
	165, 187, 187, 82, 170, 1, 11, 168, 81, 223,
	121, 59, 199, 102, 58, 57, 56, 82, 103, 55,
	54, 30, 229, 187, 29, 28, 187, 27, 45, 25,
	24, 37, 38, 23, 240, 15, 124, 82, 242, 33,
	26, 239, 164, 31, 32, 19, 220, 244, 83, 18,
	39, 82, 40, 249, 36, 35, 82, 20, 17, 254,
	104, 5, 187, 0, 0, 256, 187, 0, 16, 0,
	0, 187, 187, 0, 82, 83, 261, 0, 0, 0,
	0, 84, 0, 187, 187, 187, 0, 83, 272, 0,
	187, 0, 0, 0, 187, 271, 0, 83, 83, 0,
	83, 83, 83, 83, 82, 0, 83, 0, 84, 83,
	83, 0, 0, 209, 210, 0, 83, 0, 215, 0,
	84, 0, 145, 149, 0, 82, 0, 83, 0, 0,
	84, 84, 0, 84, 84, 84, 84, 0, 0, 84,
	0, 83, 84, 84, 0, 175, 0, 177, 0, 84,
	0, 0, 0, 0, 0, 180, 181, 0, 0, 238,
	84, 83, 0, 0, 0, 0, 11, 80, 81, 70,
	0, 75, 85, 0, 84, 83, 0, 0, 0, 82,
	83, 0, 0, 202, 0, 205, 0, 258, 0, 259,
	0, 37, 38, 0, 84, 74, 0, 0, 83, 61,
	0, 0, 0, 0, 193, 0, 0, 0, 84, 0,
	73, 0, 86, 84, 36, 35, 0, 0, 0, 274,
	276, 0, 0, 278, 0, 279, 0, 0, 83, 0,
	0, 84, 62, 63, 64, 0, 0, 0, 0, 60,
	67, 65, 66, 0, 11, 80, 81, 192, 0, 83,
	85, 0, 241, 11, 80, 81, 70, 245, 0, 85,
	0, 84, 250, 0, 0, 251, 0, 0, 0, 37,
	38, 0, 0, 0, 0, 0, 0, 0, 37, 38,
	0, 0, 84, 0, 0, 262, 263, 0, 39, 264,
	86, 0, 36, 35, 0, 0, 0, 73, 269, 86,
	0, 36, 35, 83, 0, 0, 0, 273, 11, 12,
	13, 52, 0, 0, 0, 42, 222, 50, 225, 224,
	51, 43, 44, 0, 0, 0, 53, 0, 0, 0,
	0, 0, 0, 37, 38, 0, 84, 0, 46, 47,
	48, 49, 0, 0, 185, 186, 11, 12, 13, 52,
	0, 0, 39, 42, 40, 50, 36, 35, 51, 43,
	44, 0, 0, 0, 53, 0, 0, 0, 0, 0,
	0, 37, 38, 0, 0, 0, 46, 47, 48, 49,
	0, 0, 6, 7, 0, 0, 0, 0, 0, 0,
	39, 0, 40, 0, 36, 35, 0, 8, 11, 12,
	13, 52, 0, 0, 0, 42, 260, 50, 0, 0,
	51, 43, 44, 0, 0, 0, 53, 0, 0, 0,
	0, 0, 0, 37, 38, 0, 0, 0, 46, 47,
	48, 49, 0, 0, 185, 186, 11, 12, 13, 52,
	0, 0, 39, 42, 40, 50, 36, 35, 51, 43,
	44, 0, 0, 0, 53, 0, 0, 0, 0, 0,
	0, 37, 38, 0, 0, 0, 46, 47, 48, 49,
	0, 0, 185, 186, 0, 0, 0, 0, 0, 0,
	39, 0, 40, 234, 36, 35, 11, 12, 13, 52,
	0, 0, 0, 42, 231, 50, 0, 0, 51, 43,
	44, 0, 0, 0, 53, 0, 0, 0, 0, 0,
	0, 37, 38, 0, 0, 0, 46, 47, 48, 49,
	0, 0, 185, 186, 0, 0, 0, 0, 0, 0,
	39, 0, 40, 0, 36, 35, 11, 12, 13, 52,
	0, 0, 0, 42, 226, 50, 0, 0, 51, 43,
	44, 0, 0, 0, 53, 0, 0, 0, 0, 0,
	0, 37, 38, 0, 0, 0, 46, 47, 48, 49,
	0, 0, 185, 186, 0, 0, 0, 0, 0, 0,
	39, 0, 40, 0, 36, 35, 11, 12, 13, 52,
	0, 0, 0, 42, 221, 50, 0, 0, 51, 43,
	44, 0, 0, 0, 53, 0, 0, 0, 0, 0,
	0, 37, 38, 0, 0, 0, 46, 47, 48, 49,
	0, 0, 185, 186, 0, 0, 0, 0, 0, 0,
	39, 0, 40, 0, 36, 35, 11, 12, 13, 52,
	0, 0, 0, 42, 219, 50, 0, 0, 51, 43,
	44, 0, 0, 0, 53, 0, 0, 0, 0, 0,
	0, 37, 38, 0, 0, 0, 46, 47, 48, 49,
	0, 0, 185, 186, 11, 12, 13, 52, 0, 0,
	39, 42, 40, 50, 36, 35, 51, 43, 44, 0,
	0, 0, 53, 0, 0, 0, 0, 0, 0, 37,
	38, 0, 0, 0, 46, 47, 48, 49, 0, 0,
	185, 186, 0, 0, 0, 0, 0, 0, 39, 0,
	40, 206, 36, 35, 11, 12, 13, 52, 0, 0,
	0, 42, 201, 50, 0, 0, 51, 43, 44, 0,
	0, 0, 53, 0, 0, 0, 0, 0, 0, 37,
	38, 0, 0, 0, 46, 47, 48, 49, 0, 0,
	185, 186, 11, 12, 13, 52, 0, 0, 39, 42,
	40, 50, 36, 35, 51, 43, 44, 0, 0, 0,
	53, 0, 0, 0, 0, 0, 0, 37, 38, 0,
	0, 0, 46, 47, 48, 49, 0, 0, 185, 186,
	11, 12, 13, 52, 184, 0, 39, 42, 40, 50,
	36, 35, 51, 43, 44, 0, 0, 0, 53, 0,
	0, 0, 0, 0, 0, 37, 38, 0, 0, 0,
	46, 47, 48, 49, 0, 0, 0, 183, 11, 12,
	13, 52, 0, 0, 39, 42, 40, 50, 36, 35,
	51, 43, 44, 0, 0, 0, 53, 0, 0, 0,
	0, 0, 0, 37, 38, 0, 0, 0, 46, 47,
	48, 49, 0, 0, 0, 122, 11, 12, 13, 52,
	0, 0, 39, 42, 40, 50, 36, 35, 51, 43,
	44, 11, 80, 81, 53, 0, 0, 85, 0, 0,
	0, 37, 38, 0, 0, 0, 46, 47, 48, 49,
	0, 0, 0, 0, 0, 0, 37, 38, 0, 0,
	39, 0, 40, 0, 36, 35, 11, 120, 81, 0,
	0, 0, 85, 0, 0, 39, 0, 86, 0, 36,
	35, 11, 80, 81, 0, 136, 0, 0, 0, 0,
	0, 37, 38, 11, 120, 81, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 37, 38, 0, 0,
	39, 0, 86, 0, 36, 35, 0, 0, 37, 38,
	11, 80, 81, 0, 0, 39, 0, 40, 0, 36,
	35, 101, 80, 81, 0, 0, 0, 39, 0, 40,
	0, 36, 35, 0, 0, 37, 38, 0, 61, 0,
	0, 0, 0, 0, 0, 0, 37, 38, 0, 0,
	0, 0, 0, 61, 39, 0, 40, 61, 36, 35,
	0, 0, 0, 0, 0, 39, 68, 40, 0, 36,
	35, 62, 63, 64, 0, 0, 0, 0, 60, 67,
	65, 66, 0, 0, 197, 0, 62, 63, 64, 0,
	62, 63, 64, 60, 67, 65, 66, 60, 67, 65,
	66,
}
var RubyPact = []int{

	-30, 521, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 1199,
	25, -1000, 341, 2, 1, 0, -1, -1000, -1000, -1000,
	-1000, -1000, 71, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 15, 131, -1000, -1000, 1166,
	-1000, -20, 155, 112, 112, 19, 1051, 1051, 1051, 1051,
	1051, 1128, 1013, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	154, 1128, 1155, 1051, 1128, 1128, 1128, 1128, 1051, 1051,
	1116, -1000, 34, 1166, 1155, 149, 85, 39, -1000, -1000,
	428, -1000, -1000, -1000, -1000, -28, -28, 1051, 1051, 1051,
	1051, 1155, 1051, 1051, -1000, -1000, 129, -1000, -1000, 20,
	14, -1000, 1203, -1000, -9, 181, -23, 99, 6, -1000,
	-1000, 1051, 145, 25, 1199, 25, 25, 25, 25, -1000,
	-1000, 975, -1000, 25, 937, 419, -1000, 39, -1000, 25,
	-1000, -1000, -1000, -1000, 25, 25, -1000, 375, 124, 1101,
	1184, 39, -1000, -1000, 1066, 899, -1000, 143, -1000, 849,
	25, 25, 25, 25, 39, -1000, 25, 25, -1000, -1000,
	159, -1000, 1128, -1000, -1000, -1000, 79, 164, -24, 134,
	-1000, 36, 142, -1000, -2, 811, 112, 761, 25, -1000,
	483, 711, 25, -1000, -1000, -1000, -1000, 1199, 25, 24,
	34, -1000, 1155, -1000, -1000, -1000, -1000, -3, 39, -1000,
	-1000, -1000, 661, 5, -1000, 611, -1000, -1000, -1000, 32,
	-27, -1000, 1051, 1128, -1000, -14, 142, 114, 1051, -1000,
	-1000, -1000, -1000, 96, 1051, -1000, -1000, -1000, 109, 105,
	1051, -1000, -1000, 138, -1000, -1000, 1051, -1000, 45, 25,
	-1000, 937, -1000, -1000, 25, 573, -1000, 1051, -1000, 25,
	937, 937, 163, -1000, 25, -1000, 86, -25, -14, 41,
	-1000, 25, 937, 937, 937, 136, 1051, 1128, -1000, 937,
	-1000, 103, 31, 937, -14, -1000, -14, -1000, -14, -14,
}
var RubyPgo = []int{

	0, 8, 241, 238, 237, 7, 229, 225, 224, 248,
	223, 220, 219, 0, 215, 47, 213, 5, 210, 91,
	209, 49, 1, 208, 207, 205, 204, 201, 200, 199,
	196, 195, 194, 191, 216, 190, 6, 16, 189, 185,
	184, 180, 179, 9, 175, 174, 171, 170, 3, 4,
	169, 127,
}
var RubyR1 = []int{

	0, 39, 39, 39, 39, 39, 39, 39, 39, 51,
	51, 2, 2, 34, 34, 34, 34, 34, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 17,
	17, 17, 17, 17, 17, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 36, 36, 45, 45, 43, 43, 43, 43, 43,
	48, 48, 48, 48, 47, 47, 47, 47, 47, 16,
	40, 40, 22, 22, 49, 49, 49, 18, 18, 20,
	21, 21, 50, 50, 11, 11, 11, 11, 11, 11,
	11, 9, 9, 19, 19, 14, 14, 23, 23, 24,
	25, 26, 27, 28, 28, 28, 28, 29, 30, 31,
	32, 33, 3, 6, 7, 7, 4, 4, 41, 41,
	41, 41, 46, 46, 46, 46, 5, 5, 5, 5,
	37, 44, 44, 44, 10, 10, 10, 10, 10, 10,
	10, 38, 38, 38, 38, 38, 35, 35, 35, 35,
	35, 8, 12, 42, 42, 42, 42,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 2, 2, 2, 2, 0,
	2, 1, 1, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
	4, 4, 4, 2, 3, 4, 4, 2, 3, 4,
	6, 3, 1, 1, 3, 0, 1, 1, 1, 3,
	1, 1, 3, 3, 1, 1, 3, 3, 3, 7,
	1, 3, 1, 3, 0, 1, 3, 4, 6, 4,
	1, 4, 1, 4, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 2, 2, 3, 3, 3, 3, 2,
	2, 2, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 1, 1, 3, 3, 5, 5, 0, 4,
	7, 8, 3, 3, 7, 8, 3, 4, 4, 3,
	3, 0, 1, 3, 4, 5, 3, 3, 3, 3,
	4, 0, 4, 3, 3, 2, 0, 1, 1, 2,
	2, 3, 4, 0, 3, 4, 6,
}
var RubyChk = []int{

	-1000, -39, 41, 42, 56, -2, 41, 42, 56, -13,
	-1, 5, 6, 7, -19, -14, -9, -3, -6, -7,
	-4, -21, -15, -16, -18, -20, -11, -24, -25, -26,
	-27, -10, -8, -12, -17, 54, 53, 30, 31, 49,
	51, -50, 12, 18, 19, -23, 35, 36, 37, 38,
	14, 17, 8, 23, -28, -29, -30, -31, -32, -33,
	44, 4, 37, 38, 39, 46, 47, 45, 17, 14,
	8, -36, -48, 49, 34, 10, -47, -13, -5, -15,
	6, 7, -19, -14, -9, 11, 51, 34, 34, 34,
	34, 37, 14, 17, 6, 7, 54, 6, 7, -45,
	-43, 5, -13, -17, -15, -51, 43, 6, -21, 7,
	-21, 34, 10, -1, -13, -1, -1, -1, -1, -13,
	6, -35, 42, -1, -34, 6, -13, -13, -15, -1,
	-13, -13, -13, -13, -1, -1, 9, -13, -43, 10,
	-13, -13, -15, 6, 10, -34, -37, 45, -37, -34,
	-1, -1, -1, -1, -13, -15, -1, -1, 6, 7,
	10, 50, 10, 50, 41, -41, -46, -13, 6, 43,
	-40, -49, 8, -22, 6, -34, 32, -34, -1, 6,
	-34, -34, -1, 42, 9, 41, 42, -13, -1, -42,
	-48, -36, 8, 9, 9, -13, -5, 50, -13, -15,
	-5, 13, -34, -44, 6, -34, 52, 5, -13, -51,
	-51, 10, 4, 43, 7, -51, 10, -49, 34, 13,
	-21, 13, 13, -38, 16, 15, 13, 13, 24, -43,
	34, 13, 45, 10, 52, 52, 10, 52, -51, -1,
	-13, -34, -22, 9, -1, -34, 13, 16, 15, -1,
	-34, -34, 7, 9, -1, 6, -1, 6, -51, -51,
	13, -1, -34, -34, -34, 4, 4, 43, 13, -34,
	6, -1, -13, -34, -51, 10, -51, 10, -51, -51,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 5, 6, 7, 8, 11,
	12, 18, 19, -2, 21, 22, 23, 24, 25, 26,
	27, 28, 36, 37, 38, 39, 40, 41, 42, 43,
	44, 45, 46, 47, 48, 0, 0, 122, 123, 65,
	9, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 156, 13, 29, 30, 31, 32, 33, 34,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	65, 53, 57, 65, 0, 0, 62, 70, 71, 75,
	19, -2, 21, 22, 23, 13, -2, 0, 0, 0,
	0, 0, 0, 0, 103, 104, 0, 101, 102, 0,
	0, 18, 66, 67, 68, 128, 0, 84, 13, -2,
	13, 0, 0, 109, 35, 110, 111, 112, 13, 13,
	19, 0, 157, 158, 163, 54, 58, 113, 115, 117,
	118, 119, 120, 121, 148, 146, 49, 66, 0, 0,
	66, 94, 95, 107, 0, 0, 13, 141, 13, 0,
	96, 97, 98, 99, 114, 116, 147, 149, 105, 106,
	0, 124, 0, 125, 10, 9, 9, 0, 19, 0,
	9, 80, 84, 85, 82, 0, 0, 0, 100, 108,
	0, 0, 159, 160, 161, 14, 15, 16, 17, 0,
	55, 56, 65, 50, 51, 72, 73, 59, 76, 77,
	78, 136, 0, 0, 142, 0, 139, 64, 69, 0,
	0, 9, 0, 0, -2, 13, 0, 0, 0, 87,
	13, 89, 144, 0, 0, 13, 150, 162, 13, 0,
	0, 137, 140, 0, 138, 126, 0, 127, 0, 9,
	132, 9, 86, 81, 83, 0, 145, 0, 13, 13,
	155, 164, 13, 61, 60, 143, 0, 0, 129, 0,
	88, 13, 153, 154, 165, 0, 0, 0, 79, 152,
	13, 9, 9, 166, 130, 9, 134, 9, 131, 135,
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
		//line parser.y:159
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:161
		{
		}
	case 3:
		//line parser.y:163
		{
		}
	case 4:
		//line parser.y:165
		{
		}
	case 5:
		//line parser.y:167
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 6:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 7:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 8:
		//line parser.y:175
		{
		}
	case 9:
		//line parser.y:177
		{
		}
	case 10:
		//line parser.y:178
		{
		}
	case 11:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 12:
		//line parser.y:182
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
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
		//line parser.y:203
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: []ast.Node{},
			}
		}
	case 50:
		//line parser.y:210
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 51:
		//line parser.y:217
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 52:
		//line parser.y:224
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 53:
		//line parser.y:231
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 54:
		//line parser.y:238
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 55:
		//line parser.y:245
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 56:
		//line parser.y:253
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 57:
		//line parser.y:263
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 58:
		//line parser.y:270
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 59:
		//line parser.y:280
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 60:
		//line parser.y:290
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 61:
		//line parser.y:299
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 62:
		//line parser.y:301
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 63:
		//line parser.y:304
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 64:
		//line parser.y:306
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 65:
		//line parser.y:308
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 66:
		//line parser.y:310
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:312
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:314
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 69:
		//line parser.y:316
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 70:
		//line parser.y:320
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 71:
		//line parser.y:322
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 72:
		//line parser.y:324
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:326
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:329
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 75:
		//line parser.y:331
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:333
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:335
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 78:
		//line parser.y:337
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 79:
		//line parser.y:343
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 80:
		//line parser.y:352
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 81:
		//line parser.y:354
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 82:
		//line parser.y:357
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 83:
		//line parser.y:359
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 84:
		//line parser.y:361
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 85:
		//line parser.y:363
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 86:
		//line parser.y:367
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 87:
		//line parser.y:372
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 88:
		//line parser.y:379
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 89:
		//line parser.y:389
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 90:
		//line parser.y:398
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 91:
		//line parser.y:404
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 92:
		//line parser.y:412
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 93:
		//line parser.y:416
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 94:
		//line parser.y:421
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 95:
		//line parser.y:428
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 96:
		//line parser.y:435
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 97:
		//line parser.y:442
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 98:
		//line parser.y:449
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 99:
		//line parser.y:456
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 100:
		//line parser.y:463
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 101:
		//line parser.y:471
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 102:
		//line parser.y:473
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 103:
		//line parser.y:476
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 104:
		//line parser.y:478
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 105:
		//line parser.y:481
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 106:
		//line parser.y:483
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 107:
		//line parser.y:486
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 108:
		//line parser.y:488
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 109:
		//line parser.y:490
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 110:
		//line parser.y:491
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 111:
		//line parser.y:492
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 112:
		//line parser.y:493
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
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
		//line parser.y:520
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 117:
		//line parser.y:529
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 118:
		//line parser.y:538
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 119:
		//line parser.y:547
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 120:
		//line parser.y:556
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 121:
		//line parser.y:565
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 122:
		//line parser.y:573
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 123:
		//line parser.y:574
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 124:
		//line parser.y:576
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 125:
		//line parser.y:578
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
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
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 128:
		//line parser.y:597
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 129:
		//line parser.y:599
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 130:
		//line parser.y:606
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 131:
		//line parser.y:613
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 132:
		//line parser.y:621
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 133:
		//line parser.y:628
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 134:
		//line parser.y:635
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 135:
		//line parser.y:642
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 136:
		//line parser.y:650
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 137:
		//line parser.y:652
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 138:
		//line parser.y:659
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 139:
		//line parser.y:666
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 140:
		//line parser.y:669
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 141:
		//line parser.y:671
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 142:
		//line parser.y:673
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 143:
		//line parser.y:675
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 144:
		//line parser.y:678
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 145:
		//line parser.y:685
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 146:
		//line parser.y:693
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 147:
		//line parser.y:700
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 148:
		//line parser.y:707
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 149:
		//line parser.y:714
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 150:
		//line parser.y:721
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-2].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 151:
		//line parser.y:728
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 152:
		//line parser.y:730
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 153:
		//line parser.y:737
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 154:
		//line parser.y:744
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 155:
		//line parser.y:751
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 156:
		//line parser.y:758
		{
		}
	case 157:
		//line parser.y:759
		{
		}
	case 158:
		//line parser.y:760
		{
			RubyVAL.genericSlice = []ast.Node{RubyS[Rubypt-0].genericValue}
		}
	case 159:
		//line parser.y:761
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 160:
		//line parser.y:762
		{
		}
	case 161:
		//line parser.y:765
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 162:
		//line parser.y:768
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 163:
		//line parser.y:776
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 164:
		//line parser.y:778
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 165:
		//line parser.y:780
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 166:
		//line parser.y:789
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
