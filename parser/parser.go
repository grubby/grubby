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

//line parser.y:807

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 25,
	43, 93,
	-2, 20,
	-1, 77,
	9, 65,
	10, 65,
	-2, 157,
	-1, 88,
	43, 93,
	-2, 20,
	-1, 93,
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
	-1, 103,
	43, 93,
	-2, 91,
	-1, 114,
	43, 93,
	-2, 20,
	-1, 199,
	9, 65,
	10, 65,
	-2, 157,
	-1, 236,
	43, 94,
	-2, 92,
}

const RubyNprod = 166
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1263

var RubyAct = []int{

	9, 165, 85, 167, 164, 79, 33, 78, 62, 257,
	155, 123, 23, 87, 88, 77, 156, 82, 92, 191,
	10, 72, 22, 189, 187, 84, 76, 117, 118, 281,
	258, 2, 3, 64, 235, 60, 59, 54, 55, 196,
	191, 81, 102, 104, 112, 86, 4, 73, 241, 129,
	61, 256, 277, 191, 72, 171, 80, 125, 93, 276,
	53, 52, 274, 190, 188, 132, 133, 106, 136, 137,
	138, 139, 288, 63, 142, 119, 238, 127, 125, 126,
	191, 149, 150, 240, 217, 95, 134, 191, 84, 147,
	74, 105, 99, 75, 143, 98, 125, 191, 127, 97,
	126, 127, 151, 126, 157, 286, 169, 159, 86, 62,
	62, 96, 148, 73, 227, 233, 127, 182, 126, 215,
	72, 23, 87, 88, 44, 228, 95, 92, 249, 194,
	251, 250, 84, 280, 262, 189, 26, 197, 153, 198,
	245, 215, 279, 62, 213, 189, 54, 55, 168, 201,
	166, 202, 86, 255, 204, 182, 206, 200, 189, 182,
	236, 89, 185, 186, 103, 56, 264, 93, 216, 53,
	52, 182, 168, 182, 205, 210, 182, 182, 220, 174,
	89, 120, 121, 101, 152, 100, 146, 141, 131, 229,
	230, 272, 234, 89, 58, 83, 193, 122, 209, 184,
	125, 89, 89, 192, 89, 89, 89, 89, 1, 182,
	89, 237, 182, 223, 89, 115, 51, 89, 89, 244,
	127, 50, 126, 49, 89, 48, 47, 46, 18, 17,
	16, 27, 89, 15, 37, 13, 261, 12, 11, 21,
	14, 19, 128, 31, 182, 30, 32, 29, 182, 182,
	0, 0, 0, 0, 182, 182, 90, 0, 0, 0,
	0, 0, 0, 0, 0, 89, 0, 0, 89, 0,
	182, 182, 182, 28, 0, 90, 0, 0, 0, 182,
	0, 116, 284, 182, 0, 89, 0, 0, 90, 0,
	89, 0, 0, 0, 0, 0, 90, 90, 91, 90,
	90, 90, 90, 0, 0, 90, 0, 0, 0, 90,
	0, 0, 90, 90, 0, 0, 0, 91, 0, 90,
	0, 0, 0, 0, 0, 0, 89, 90, 0, 0,
	91, 0, 0, 23, 113, 114, 89, 0, 91, 91,
	0, 91, 91, 91, 91, 0, 0, 91, 183, 0,
	5, 91, 0, 0, 91, 91, 0, 0, 54, 55,
	90, 91, 0, 90, 0, 0, 0, 0, 0, 91,
	0, 0, 89, 0, 154, 158, 0, 56, 0, 57,
	90, 53, 52, 0, 170, 90, 172, 107, 108, 109,
	110, 111, 0, 175, 176, 0, 0, 0, 0, 0,
	0, 0, 91, 20, 0, 91, 0, 214, 0, 0,
	0, 130, 218, 0, 0, 135, 0, 0, 89, 0,
	140, 90, 91, 144, 145, 0, 0, 91, 94, 0,
	0, 90, 0, 0, 0, 231, 232, 208, 0, 211,
	0, 0, 0, 0, 0, 160, 161, 162, 163, 23,
	87, 88, 199, 0, 173, 92, 0, 0, 0, 0,
	94, 0, 0, 91, 177, 0, 0, 90, 0, 94,
	0, 0, 0, 91, 54, 55, 259, 94, 0, 0,
	0, 94, 0, 0, 94, 94, 265, 0, 0, 0,
	266, 94, 0, 56, 0, 93, 243, 53, 52, 94,
	247, 0, 248, 275, 0, 0, 0, 253, 0, 91,
	254, 0, 0, 90, 23, 87, 88, 77, 0, 0,
	92, 23, 87, 88, 44, 0, 285, 287, 0, 289,
	0, 290, 0, 269, 270, 94, 0, 271, 0, 54,
	55, 0, 0, 0, 0, 0, 54, 55, 0, 0,
	278, 0, 23, 195, 114, 91, 0, 94, 80, 0,
	93, 282, 53, 52, 0, 56, 246, 57, 0, 53,
	52, 0, 0, 252, 0, 0, 0, 54, 55, 0,
	0, 0, 0, 260, 0, 0, 0, 263, 191, 0,
	0, 0, 0, 0, 0, 0, 56, 0, 57, 268,
	53, 52, 0, 94, 0, 0, 273, 23, 24, 25,
	44, 0, 0, 0, 34, 222, 42, 225, 224, 43,
	35, 36, 0, 0, 0, 45, 0, 0, 0, 283,
	0, 0, 54, 55, 0, 0, 0, 38, 39, 40,
	41, 0, 0, 180, 181, 23, 24, 25, 44, 0,
	0, 56, 34, 57, 42, 53, 52, 43, 35, 36,
	0, 0, 0, 45, 0, 0, 0, 0, 0, 0,
	54, 55, 0, 0, 0, 38, 39, 40, 41, 0,
	0, 6, 7, 0, 0, 0, 0, 0, 0, 56,
	0, 57, 0, 53, 52, 0, 8, 23, 24, 25,
	44, 0, 0, 0, 34, 267, 42, 0, 0, 43,
	35, 36, 0, 0, 0, 45, 0, 0, 0, 0,
	0, 0, 54, 55, 0, 0, 0, 38, 39, 40,
	41, 0, 0, 180, 181, 23, 24, 25, 44, 0,
	0, 56, 34, 57, 42, 53, 52, 43, 35, 36,
	0, 0, 0, 45, 0, 0, 0, 0, 0, 0,
	54, 55, 0, 0, 0, 38, 39, 40, 41, 0,
	0, 180, 181, 0, 0, 0, 0, 0, 0, 56,
	0, 57, 242, 53, 52, 23, 24, 25, 44, 0,
	0, 0, 34, 239, 42, 0, 0, 43, 35, 36,
	0, 0, 0, 45, 0, 0, 0, 0, 0, 0,
	54, 55, 0, 0, 0, 38, 39, 40, 41, 0,
	0, 180, 181, 0, 0, 0, 0, 0, 0, 56,
	0, 57, 0, 53, 52, 23, 24, 25, 44, 0,
	0, 0, 34, 226, 42, 0, 0, 43, 35, 36,
	0, 0, 0, 45, 0, 0, 0, 0, 0, 0,
	54, 55, 0, 0, 0, 38, 39, 40, 41, 0,
	0, 180, 181, 0, 0, 0, 0, 0, 0, 56,
	0, 57, 0, 53, 52, 23, 24, 25, 44, 0,
	0, 0, 34, 221, 42, 0, 0, 43, 35, 36,
	0, 0, 0, 45, 0, 0, 0, 0, 0, 0,
	54, 55, 0, 0, 0, 38, 39, 40, 41, 0,
	0, 180, 181, 0, 0, 0, 0, 0, 0, 56,
	0, 57, 0, 53, 52, 23, 24, 25, 44, 0,
	0, 0, 34, 219, 42, 0, 0, 43, 35, 36,
	0, 0, 0, 45, 0, 0, 0, 0, 0, 0,
	54, 55, 0, 0, 0, 38, 39, 40, 41, 0,
	0, 180, 181, 23, 24, 25, 44, 0, 0, 56,
	34, 57, 42, 53, 52, 43, 35, 36, 0, 0,
	0, 45, 0, 0, 0, 0, 0, 0, 54, 55,
	0, 0, 0, 38, 39, 40, 41, 0, 0, 180,
	181, 0, 0, 0, 0, 0, 0, 56, 0, 57,
	212, 53, 52, 23, 24, 25, 44, 0, 0, 0,
	34, 207, 42, 0, 0, 43, 35, 36, 0, 0,
	0, 45, 0, 0, 0, 0, 0, 0, 54, 55,
	0, 0, 0, 38, 39, 40, 41, 0, 0, 180,
	181, 23, 24, 25, 44, 0, 0, 56, 34, 57,
	42, 53, 52, 43, 35, 36, 0, 0, 0, 45,
	0, 0, 0, 0, 0, 0, 54, 55, 0, 0,
	0, 38, 39, 40, 41, 0, 0, 180, 181, 23,
	24, 25, 44, 179, 0, 56, 34, 57, 42, 53,
	52, 43, 35, 36, 0, 0, 0, 45, 0, 0,
	0, 0, 0, 0, 54, 55, 0, 0, 0, 38,
	39, 40, 41, 0, 0, 0, 178, 23, 24, 25,
	44, 0, 0, 56, 34, 57, 42, 53, 52, 43,
	35, 36, 23, 113, 114, 45, 0, 0, 92, 0,
	0, 0, 54, 55, 0, 0, 0, 38, 39, 40,
	41, 0, 0, 0, 64, 0, 0, 54, 55, 0,
	0, 56, 0, 57, 0, 53, 52, 124, 87, 88,
	44, 0, 0, 0, 0, 0, 56, 0, 93, 0,
	53, 52, 0, 0, 0, 64, 0, 65, 66, 67,
	0, 0, 54, 55, 63, 70, 68, 69, 71, 64,
	203, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 56, 0, 57, 0, 53, 52, 0, 65, 66,
	67, 0, 0, 0, 0, 63, 70, 68, 69, 0,
	0, 0, 65, 66, 67, 0, 0, 0, 0, 63,
	70, 68, 69,
}
var RubyPact = []int{

	-10, 640, -1000, -1000, -1000, -6, -1000, -1000, -1000, 1201,
	76, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-18, -1000, -1000, -1000, 7, 77, 65, 61, 58, -1000,
	-1000, -1000, -1000, -1000, 179, 157, 157, 57, 1132, 1132,
	1132, 1132, 1132, 328, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 21, 175, -1000, -1000, 1182, -1000, 6, -1000,
	-1000, -1000, 1132, 182, 328, 516, 1132, 328, 328, 328,
	328, 1132, 181, 516, 1132, 1132, 180, 516, -1000, 102,
	1182, 516, 178, 128, 29, -1000, -23, 509, 118, -1000,
	-1000, -1000, -29, -29, -18, 516, 1132, 1132, 1132, 1132,
	142, 142, 23, -1000, -1000, 1132, 173, 96, 96, 96,
	96, 96, -1000, -1000, -1000, 1094, 1056, -1000, -1000, 156,
	-1000, -1000, 14, 13, -1000, 1215, -1000, 10, 547, -4,
	96, 444, -1000, 29, -23, 96, -1000, -1000, -1000, -1000,
	96, -1000, 29, -23, 96, 96, -1000, 148, 1147, 1170,
	29, -23, -1000, 116, 1018, -1000, 169, -1000, 968, 135,
	96, 96, 96, 96, -1000, 109, 166, -1000, 50, -1000,
	930, 157, 880, 96, -1000, 602, 830, 96, -1000, -1000,
	-1000, -1000, 1201, 96, 101, -1000, -1000, 184, -1000, 328,
	-1000, -1000, -1000, 105, 188, -9, 153, 102, -1000, 516,
	-1000, -1000, -1000, 42, 29, -23, -1000, -1000, 780, 38,
	-1000, 730, -1000, -1000, 12, 166, 131, 1132, 12, -1000,
	-1000, -1000, -1000, 115, 1132, -1000, -1000, -1000, 146, -1000,
	-1000, -1, -22, -1000, 1132, 328, -1000, 125, 1132, -1000,
	-1000, 160, -1000, 1056, -1000, -1000, 96, 1056, 692, -1000,
	1132, -1000, 96, 1056, 1056, 187, -1000, 1132, -1000, 56,
	96, -1000, -1000, 96, -1000, 46, 39, -1000, 96, 1056,
	1056, 1056, 136, 129, -14, 12, -1000, -1000, 1056, -1000,
	1132, 328, 1056, 95, 62, 12, -1000, 12, -1000, 12,
	12,
}
var RubyPgo = []int{

	0, 348, 247, 246, 2, 245, 243, 403, 273, 241,
	240, 239, 0, 231, 20, 238, 22, 237, 136, 235,
	6, 3, 234, 233, 230, 229, 228, 227, 226, 225,
	223, 221, 216, 281, 215, 7, 10, 213, 208, 4,
	203, 199, 11, 198, 197, 196, 195, 5, 1, 194,
	242,
}
var RubyR1 = []int{

	0, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 50, 50, 33, 33, 33, 33, 33, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 16,
	16, 16, 16, 16, 16, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 35, 35, 44, 44, 42, 42, 42, 42, 42,
	47, 47, 47, 47, 46, 46, 46, 46, 46, 15,
	15, 39, 39, 21, 21, 48, 48, 48, 17, 17,
	19, 20, 20, 49, 49, 10, 10, 10, 10, 10,
	10, 10, 8, 8, 18, 18, 13, 13, 22, 22,
	23, 24, 25, 26, 27, 27, 27, 27, 28, 29,
	30, 31, 32, 2, 5, 6, 6, 3, 3, 40,
	40, 40, 40, 45, 45, 45, 45, 4, 4, 4,
	4, 36, 43, 43, 43, 9, 9, 9, 9, 9,
	9, 9, 37, 37, 37, 37, 37, 34, 34, 34,
	7, 11, 41, 41, 41, 41,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 4,
	4, 2, 3, 3, 4, 4, 3, 2, 3, 4,
	6, 3, 1, 1, 3, 0, 1, 1, 1, 3,
	1, 1, 3, 3, 1, 1, 3, 3, 3, 7,
	7, 1, 3, 1, 3, 0, 1, 3, 4, 6,
	4, 1, 4, 1, 4, 3, 3, 3, 3, 3,
	3, 3, 2, 2, 2, 2, 3, 3, 3, 3,
	2, 2, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 1, 1, 3, 3, 5, 5, 0,
	4, 7, 8, 3, 3, 7, 8, 3, 4, 4,
	3, 3, 0, 1, 3, 4, 5, 3, 3, 3,
	3, 4, 0, 4, 3, 3, 2, 0, 2, 2,
	3, 4, 0, 3, 4, 6,
}
var RubyChk = []int{

	-1000, -38, 41, 42, 56, -1, 41, 42, 56, -12,
	-14, -15, -17, -19, -10, -23, -24, -25, -26, -9,
	-7, -11, -16, 5, 6, 7, -18, -13, -8, -2,
	-5, -6, -3, -20, 12, 18, 19, -22, 35, 36,
	37, 38, 14, 17, 8, 23, -27, -28, -29, -30,
	-31, -32, 54, 53, 30, 31, 49, 51, -49, 42,
	41, 56, 14, 44, 4, 37, 38, 39, 46, 47,
	45, 17, 44, 37, 14, 17, 44, 8, -35, -47,
	49, 34, 10, -46, -12, -4, -14, 6, 7, -18,
	-13, -8, 11, 51, -7, 8, 34, 34, 34, 34,
	6, 4, -20, 7, -20, 34, 10, -1, -1, -1,
	-1, -1, -12, 6, 7, -34, -33, 6, 7, 54,
	6, 7, -44, -42, 5, -12, -16, -14, -50, 43,
	-1, 6, -12, -12, -14, -1, -12, -12, -12, -12,
	-1, 6, -12, -14, -1, -1, 6, -42, 10, -12,
	-12, -14, 6, 10, -33, -36, 45, -36, -33, -42,
	-1, -1, -1, -1, -39, -48, 8, -21, 6, -39,
	-33, 32, -33, -1, 6, -33, -33, -1, 42, 9,
	41, 42, -12, -1, -41, 6, 7, 10, 50, 10,
	50, 41, -40, -45, -12, 6, 43, -47, -35, 8,
	9, -12, -4, 50, -12, -14, -4, 13, -33, -43,
	6, -33, 52, 9, -50, 10, -48, 34, -50, 13,
	-20, 13, 13, -37, 16, 15, 13, 13, 24, 5,
	-12, -50, -50, 10, 4, 43, 7, -42, 34, 13,
	45, 10, 52, -33, -21, 9, -1, -33, -33, 13,
	16, 15, -1, -33, -33, 7, 52, 10, 52, -50,
	-1, -12, 9, -1, 6, -50, -50, 13, -1, -33,
	-33, -33, 4, -1, 6, -50, 13, 13, -33, 6,
	4, 43, -33, -1, -12, -50, 10, -50, 10, -50,
	-50,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
	46, 47, 48, 18, 19, -2, 21, 22, 23, 24,
	25, 26, 27, 28, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 157, 13, 29, 30, 31, 32,
	33, 34, 0, 0, 123, 124, 65, 11, 0, 5,
	6, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, -2, 51, 57,
	65, 0, 0, 62, 70, 71, 75, 19, -2, 21,
	22, 23, 13, -2, 0, 65, 0, 0, 0, 0,
	85, 85, 13, -2, 13, 0, 0, 110, 111, 112,
	113, 13, 13, 19, -2, 0, 162, 104, 105, 0,
	102, 103, 0, 0, 18, 66, 67, 68, 129, 0,
	147, 52, 58, 114, 116, 118, 119, 120, 121, 122,
	149, 53, 115, 117, 148, 150, 56, 0, 0, 66,
	95, 96, 108, 0, 0, 13, 142, 13, 0, 0,
	97, 98, 99, 100, 11, 81, 85, 86, 83, 11,
	0, 0, 0, 101, 109, 0, 0, 158, 159, 160,
	14, 15, 16, 17, 0, 106, 107, 0, 125, 0,
	126, 12, 11, 11, 0, 19, 0, 54, 55, -2,
	49, 72, 73, 59, 76, 77, 78, 137, 0, 0,
	143, 0, 140, 50, 13, 0, 0, 0, 13, 88,
	13, 90, 145, 0, 0, 13, 151, 161, 13, 64,
	69, 0, 0, 11, 0, 0, -2, 0, 0, 138,
	141, 0, 139, 11, 87, 82, 84, 11, 0, 146,
	0, 13, 13, 156, 163, 13, 127, 0, 128, 0,
	11, 133, 61, 60, 144, 0, 0, 89, 13, 154,
	155, 164, 0, 0, 0, 130, 79, 80, 153, 13,
	0, 0, 165, 11, 11, 131, 11, 135, 11, 132,
	136,
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
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 51:
		//line parser.y:214
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 52:
		//line parser.y:221
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 53:
		//line parser.y:228
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 54:
		//line parser.y:235
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 55:
		//line parser.y:243
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 56:
		//line parser.y:251
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 57:
		//line parser.y:261
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 58:
		//line parser.y:268
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 59:
		//line parser.y:278
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 60:
		//line parser.y:288
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 61:
		//line parser.y:297
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 62:
		//line parser.y:299
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 63:
		//line parser.y:302
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 64:
		//line parser.y:304
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 65:
		//line parser.y:306
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 66:
		//line parser.y:308
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:310
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:312
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 69:
		//line parser.y:314
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 70:
		//line parser.y:318
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 71:
		//line parser.y:320
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 72:
		//line parser.y:322
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:324
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:327
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 75:
		//line parser.y:329
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:331
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:333
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 78:
		//line parser.y:335
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 79:
		//line parser.y:341
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 80:
		//line parser.y:349
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 81:
		//line parser.y:358
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 82:
		//line parser.y:360
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 83:
		//line parser.y:363
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 84:
		//line parser.y:365
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 85:
		//line parser.y:367
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 86:
		//line parser.y:369
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 87:
		//line parser.y:373
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 88:
		//line parser.y:378
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 89:
		//line parser.y:385
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 90:
		//line parser.y:395
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 91:
		//line parser.y:404
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 92:
		//line parser.y:410
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 93:
		//line parser.y:418
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 94:
		//line parser.y:422
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
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
		//line parser.y:462
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 101:
		//line parser.y:469
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 102:
		//line parser.y:477
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 103:
		//line parser.y:479
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 104:
		//line parser.y:482
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 105:
		//line parser.y:484
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 106:
		//line parser.y:487
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 107:
		//line parser.y:489
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 108:
		//line parser.y:492
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 109:
		//line parser.y:494
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 110:
		//line parser.y:496
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 111:
		//line parser.y:497
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 112:
		//line parser.y:498
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 113:
		//line parser.y:499
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 114:
		//line parser.y:502
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 115:
		//line parser.y:510
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 116:
		//line parser.y:518
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 117:
		//line parser.y:526
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 118:
		//line parser.y:535
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 119:
		//line parser.y:544
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 120:
		//line parser.y:553
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 121:
		//line parser.y:562
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 122:
		//line parser.y:571
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 123:
		//line parser.y:579
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 124:
		//line parser.y:580
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 125:
		//line parser.y:582
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 126:
		//line parser.y:584
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 127:
		//line parser.y:587
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 128:
		//line parser.y:595
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 129:
		//line parser.y:603
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 130:
		//line parser.y:605
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 131:
		//line parser.y:612
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 132:
		//line parser.y:619
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 133:
		//line parser.y:627
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 134:
		//line parser.y:634
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 135:
		//line parser.y:641
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 136:
		//line parser.y:648
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 137:
		//line parser.y:656
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 138:
		//line parser.y:658
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 139:
		//line parser.y:665
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 140:
		//line parser.y:672
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 141:
		//line parser.y:675
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 142:
		//line parser.y:677
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 143:
		//line parser.y:679
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 144:
		//line parser.y:681
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 145:
		//line parser.y:684
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 146:
		//line parser.y:691
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 147:
		//line parser.y:699
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 148:
		//line parser.y:706
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 149:
		//line parser.y:713
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 150:
		//line parser.y:720
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 151:
		//line parser.y:727
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-2].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 152:
		//line parser.y:734
		{
			RubyVAL.genericSlice = []ast.Node{}
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
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 156:
		//line parser.y:757
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 157:
		//line parser.y:764
		{
		}
	case 158:
		//line parser.y:765
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 159:
		//line parser.y:766
		{
		}
	case 160:
		//line parser.y:769
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 161:
		//line parser.y:772
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 162:
		//line parser.y:780
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 163:
		//line parser.y:782
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 164:
		//line parser.y:784
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 165:
		//line parser.y:793
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
