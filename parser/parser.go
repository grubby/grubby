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

//line parser.y:803

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 10,
	43, 90,
	-2, 18,
	-1, 79,
	43, 90,
	-2, 18,
	-1, 84,
	8, 11,
	12, 11,
	14, 11,
	17, 11,
	18, 11,
	19, 11,
	23, 11,
	35, 11,
	36, 11,
	37, 11,
	38, 11,
	42, 11,
	-2, 5,
	-1, 107,
	43, 90,
	-2, 88,
	-1, 212,
	43, 91,
	-2, 89,
}

const RubyNprod = 165
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1282

var RubyAct = []int{

	112, 171, 6, 18, 169, 70, 144, 69, 76, 57,
	75, 2, 98, 162, 31, 234, 160, 231, 158, 92,
	93, 145, 59, 265, 211, 235, 167, 104, 266, 162,
	186, 90, 7, 89, 91, 255, 228, 100, 216, 67,
	251, 160, 103, 88, 106, 108, 162, 110, 87, 117,
	86, 101, 230, 85, 89, 174, 162, 161, 233, 159,
	124, 125, 58, 128, 129, 130, 131, 59, 94, 135,
	162, 109, 138, 139, 111, 113, 114, 115, 116, 75,
	121, 136, 275, 101, 225, 137, 101, 8, 118, 79,
	152, 146, 127, 11, 214, 226, 264, 132, 133, 209,
	60, 61, 62, 80, 165, 142, 67, 58, 65, 63,
	64, 273, 34, 35, 195, 67, 148, 149, 150, 151,
	250, 154, 155, 185, 75, 244, 107, 246, 245, 188,
	80, 189, 36, 172, 37, 170, 33, 32, 193, 212,
	176, 268, 80, 196, 185, 253, 194, 172, 185, 202,
	180, 198, 177, 80, 80, 141, 80, 80, 80, 80,
	205, 206, 80, 241, 214, 80, 80, 192, 160, 156,
	157, 123, 80, 105, 185, 215, 185, 263, 218, 185,
	185, 95, 96, 80, 210, 59, 38, 74, 164, 97,
	201, 100, 187, 163, 168, 1, 221, 80, 66, 119,
	56, 185, 55, 227, 185, 101, 207, 208, 54, 53,
	52, 213, 238, 51, 27, 26, 240, 80, 60, 61,
	62, 25, 24, 42, 22, 58, 65, 63, 64, 21,
	20, 80, 30, 23, 28, 29, 80, 16, 12, 15,
	185, 237, 17, 14, 185, 3, 0, 242, 81, 185,
	185, 0, 236, 247, 80, 0, 0, 0, 0, 252,
	13, 185, 185, 185, 0, 254, 270, 0, 185, 0,
	82, 122, 185, 0, 0, 81, 259, 0, 0, 0,
	256, 0, 257, 0, 80, 0, 0, 81, 0, 0,
	0, 0, 0, 0, 0, 269, 0, 82, 81, 81,
	0, 81, 81, 81, 81, 80, 0, 81, 0, 82,
	81, 81, 272, 274, 0, 0, 276, 81, 277, 0,
	82, 82, 0, 82, 82, 82, 82, 0, 81, 82,
	0, 0, 82, 82, 0, 0, 19, 0, 0, 82,
	0, 0, 81, 0, 0, 0, 77, 0, 0, 0,
	82, 0, 59, 0, 0, 143, 147, 191, 0, 80,
	0, 0, 81, 0, 82, 0, 0, 0, 0, 0,
	0, 0, 0, 102, 0, 0, 81, 0, 173, 0,
	175, 81, 0, 0, 82, 60, 61, 62, 178, 179,
	0, 0, 58, 65, 63, 64, 0, 126, 82, 81,
	0, 0, 0, 82, 0, 102, 0, 0, 102, 140,
	0, 0, 0, 0, 0, 77, 200, 0, 203, 0,
	0, 82, 0, 0, 0, 0, 153, 0, 0, 81,
	0, 8, 9, 10, 49, 0, 0, 0, 39, 220,
	47, 223, 222, 48, 40, 41, 0, 0, 0, 50,
	81, 82, 0, 0, 0, 0, 34, 35, 0, 59,
	77, 43, 44, 45, 46, 0, 0, 183, 184, 0,
	0, 0, 82, 0, 0, 0, 36, 0, 37, 197,
	33, 32, 0, 0, 0, 239, 0, 0, 0, 0,
	243, 0, 60, 61, 62, 248, 0, 0, 249, 58,
	65, 63, 64, 0, 81, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 260, 261,
	0, 0, 262, 0, 0, 0, 82, 102, 0, 0,
	0, 267, 0, 0, 8, 9, 10, 49, 0, 0,
	271, 39, 258, 47, 0, 0, 48, 40, 41, 0,
	0, 0, 50, 0, 0, 0, 0, 0, 0, 34,
	35, 0, 0, 0, 43, 44, 45, 46, 0, 0,
	183, 184, 0, 8, 9, 10, 49, 0, 0, 36,
	39, 37, 47, 33, 32, 48, 40, 41, 0, 0,
	0, 50, 0, 0, 0, 0, 0, 0, 34, 35,
	0, 0, 0, 43, 44, 45, 46, 0, 0, 183,
	184, 0, 0, 0, 0, 0, 0, 0, 36, 0,
	37, 232, 33, 32, 8, 9, 10, 49, 0, 0,
	0, 39, 229, 47, 0, 0, 48, 40, 41, 0,
	0, 0, 50, 0, 0, 0, 0, 0, 0, 34,
	35, 0, 0, 0, 43, 44, 45, 46, 0, 0,
	183, 184, 0, 0, 0, 0, 0, 0, 0, 36,
	0, 37, 0, 33, 32, 8, 9, 10, 49, 0,
	0, 0, 39, 224, 47, 0, 0, 48, 40, 41,
	0, 0, 0, 50, 0, 0, 0, 0, 0, 0,
	34, 35, 0, 0, 0, 43, 44, 45, 46, 0,
	0, 183, 184, 0, 0, 0, 0, 0, 0, 0,
	36, 0, 37, 0, 33, 32, 8, 9, 10, 49,
	0, 0, 0, 39, 219, 47, 0, 0, 48, 40,
	41, 0, 0, 0, 50, 0, 0, 0, 0, 0,
	0, 34, 35, 0, 0, 0, 43, 44, 45, 46,
	0, 0, 183, 184, 0, 0, 0, 0, 0, 0,
	0, 36, 0, 37, 0, 33, 32, 8, 9, 10,
	49, 0, 0, 0, 39, 217, 47, 0, 0, 48,
	40, 41, 0, 0, 0, 50, 0, 0, 0, 0,
	0, 0, 34, 35, 0, 0, 0, 43, 44, 45,
	46, 0, 0, 183, 184, 0, 8, 9, 10, 49,
	0, 0, 36, 39, 37, 47, 33, 32, 48, 40,
	41, 0, 0, 0, 50, 0, 0, 0, 0, 0,
	0, 34, 35, 0, 0, 0, 43, 44, 45, 46,
	0, 0, 183, 184, 0, 0, 0, 0, 0, 0,
	0, 36, 0, 37, 204, 33, 32, 8, 9, 10,
	49, 0, 0, 0, 39, 199, 47, 0, 0, 48,
	40, 41, 0, 0, 0, 50, 0, 0, 0, 0,
	0, 0, 34, 35, 0, 0, 0, 43, 44, 45,
	46, 0, 0, 183, 184, 0, 8, 9, 10, 49,
	0, 0, 36, 39, 37, 47, 33, 32, 48, 40,
	41, 0, 0, 0, 50, 0, 0, 0, 0, 0,
	0, 34, 35, 0, 0, 0, 43, 44, 45, 46,
	0, 0, 183, 184, 0, 8, 9, 10, 49, 182,
	0, 36, 39, 37, 47, 33, 32, 48, 40, 41,
	0, 0, 0, 50, 0, 0, 0, 0, 0, 0,
	34, 35, 0, 0, 0, 43, 44, 45, 46, 0,
	0, 0, 181, 0, 8, 9, 10, 49, 0, 0,
	36, 39, 37, 47, 33, 32, 48, 40, 41, 0,
	0, 0, 50, 0, 0, 0, 0, 0, 0, 34,
	35, 0, 0, 0, 43, 44, 45, 46, 0, 0,
	4, 5, 0, 8, 9, 10, 49, 0, 0, 36,
	39, 37, 47, 33, 32, 48, 40, 41, 0, 0,
	0, 50, 0, 0, 0, 0, 0, 0, 34, 35,
	0, 0, 0, 43, 44, 45, 46, 0, 0, 0,
	120, 0, 8, 9, 10, 49, 0, 0, 36, 39,
	37, 47, 33, 32, 48, 40, 41, 0, 0, 0,
	50, 0, 0, 0, 0, 0, 0, 34, 35, 0,
	0, 0, 43, 44, 45, 46, 8, 78, 79, 68,
	0, 73, 83, 0, 0, 0, 0, 36, 0, 37,
	0, 33, 32, 8, 78, 79, 190, 0, 0, 83,
	0, 34, 35, 0, 0, 72, 8, 78, 79, 68,
	0, 0, 83, 0, 0, 0, 0, 0, 34, 35,
	0, 71, 0, 84, 0, 33, 32, 0, 0, 0,
	0, 34, 35, 8, 78, 79, 0, 0, 36, 83,
	84, 0, 33, 32, 8, 118, 79, 0, 0, 0,
	83, 71, 0, 84, 0, 33, 32, 0, 34, 35,
	8, 166, 79, 0, 0, 0, 0, 0, 0, 34,
	35, 0, 8, 78, 79, 0, 134, 0, 36, 0,
	84, 0, 33, 32, 0, 34, 35, 0, 0, 36,
	0, 84, 0, 33, 32, 0, 162, 34, 35, 8,
	78, 79, 0, 0, 0, 36, 0, 37, 0, 33,
	32, 99, 78, 79, 0, 0, 0, 36, 0, 37,
	0, 33, 32, 0, 34, 35, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 34, 35, 0, 0,
	0, 0, 0, 0, 36, 0, 37, 0, 33, 32,
	0, 0, 0, 0, 0, 0, 36, 0, 37, 0,
	33, 32,
}
var RubyPact = []int{

	-46, 979, -1000, -48, -1000, -1000, 181, 25, -1000, 1091,
	19, 16, 14, 9, -1000, -1000, -1000, -1000, -1000, 17,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 13, 175, -1000, -1000, 1226, -1000, -16, 167,
	119, 119, 37, 1057, 1057, 1057, 1057, 1057, 82, 1018,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 165, 82,
	1214, 1057, 82, 82, 82, 82, 1057, 1057, 1187, -1000,
	75, 1226, 1214, 149, 95, 18, -1000, -1000, 1121, -1000,
	-1000, -1000, -1000, -24, -24, 1057, 1057, 1057, 1057, 1214,
	1057, 1057, -1000, -1000, 163, -1000, -1000, 8, 6, -1000,
	455, -1000, -4, 1175, -17, 127, 23, -1000, -1000, 1057,
	146, 25, 181, 25, 25, 25, 25, -1000, -1000, 940,
	-1000, 25, 901, 1108, -1000, 18, -1000, 25, -1000, -1000,
	-1000, -1000, 25, 25, -1000, 348, 158, 1159, 63, 18,
	-1000, -1000, 1148, 862, -1000, 143, -1000, 811, 25, 25,
	25, 25, 18, -1000, 25, 25, -1000, -1000, 155, -1000,
	82, -1000, -1000, -1000, 89, 180, -19, 132, -1000, 84,
	141, -1000, 4, 772, 119, 721, 25, -1000, 426, 670,
	25, -1000, -1000, -1000, -1000, 181, 25, 71, 75, -1000,
	1214, -1000, -1000, -1000, -1000, 2, 18, -1000, -1000, -1000,
	619, 7, -1000, 568, -1000, -1000, -1000, 5, -28, -1000,
	1057, 82, -1000, -12, 141, 154, 1057, -1000, -1000, -1000,
	-1000, 112, 1057, -1000, -1000, -1000, 113, 31, 1057, -1000,
	-1000, 139, -1000, -1000, 1057, -1000, 29, 25, -1000, 901,
	-1000, -1000, 25, 529, -1000, 1057, -1000, 25, 901, 901,
	173, -1000, 25, -1000, 92, -20, -12, 15, -1000, 25,
	901, 901, 901, 135, 1057, 82, -1000, 901, -1000, 101,
	72, 901, -12, -1000, -12, -1000, -12, -12,
}
var RubyPgo = []int{

	0, 30, 245, 243, 242, 8, 239, 237, 235, 260,
	234, 233, 232, 0, 238, 336, 230, 14, 229, 93,
	224, 3, 1, 223, 222, 221, 215, 214, 213, 210,
	209, 208, 202, 200, 271, 199, 7, 6, 196, 195,
	194, 193, 192, 12, 190, 189, 188, 187, 5, 4,
	186, 42,
}
var RubyR1 = []int{

	0, 39, 39, 39, 39, 51, 51, 2, 2, 2,
	2, 34, 34, 34, 34, 34, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 17, 17, 17,
	17, 17, 17, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 36,
	36, 45, 45, 43, 43, 43, 43, 43, 48, 48,
	48, 48, 47, 47, 47, 47, 47, 16, 40, 40,
	22, 22, 49, 49, 49, 18, 18, 20, 21, 21,
	50, 50, 11, 11, 11, 11, 11, 11, 11, 9,
	9, 19, 19, 14, 14, 23, 23, 24, 25, 26,
	27, 28, 28, 28, 28, 29, 30, 31, 32, 33,
	3, 6, 7, 7, 4, 4, 41, 41, 41, 41,
	46, 46, 46, 46, 5, 5, 5, 5, 37, 44,
	44, 44, 10, 10, 10, 10, 10, 10, 10, 38,
	38, 38, 38, 38, 35, 35, 35, 35, 35, 8,
	12, 42, 42, 42, 42,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 0, 2, 1, 1, 1,
	1, 0, 2, 2, 2, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 4, 4,
	4, 2, 3, 4, 4, 2, 3, 4, 6, 3,
	1, 1, 3, 0, 1, 1, 1, 3, 1, 1,
	3, 3, 1, 1, 3, 3, 3, 7, 1, 3,
	1, 3, 0, 1, 3, 4, 6, 4, 1, 4,
	1, 4, 3, 3, 3, 3, 3, 3, 3, 2,
	2, 2, 2, 3, 3, 3, 3, 2, 2, 2,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	1, 1, 3, 3, 5, 5, 0, 4, 7, 8,
	3, 3, 7, 8, 3, 4, 4, 3, 3, 0,
	1, 3, 4, 5, 3, 3, 3, 3, 4, 0,
	4, 3, 3, 2, 0, 1, 1, 2, 2, 3,
	4, 0, 3, 4, 6,
}
var RubyChk = []int{

	-1000, -39, 57, -2, 41, 42, -13, -1, 5, 6,
	7, -19, -14, -9, -3, -6, -7, -4, -21, -15,
	-16, -18, -20, -11, -24, -25, -26, -27, -10, -8,
	-12, -17, 55, 54, 30, 31, 50, 52, -50, 12,
	18, 19, -23, 35, 36, 37, 38, 14, 17, 8,
	23, -28, -29, -30, -31, -32, -33, 57, 44, 4,
	37, 38, 39, 46, 47, 45, 17, 14, 8, -36,
	-48, 50, 34, 10, -47, -13, -5, -15, 6, 7,
	-19, -14, -9, 11, 52, 34, 34, 34, 34, 37,
	14, 17, 6, 7, 55, 6, 7, -45, -43, 5,
	-13, -17, -15, -51, 43, 6, -21, 7, -21, 34,
	10, -1, -13, -1, -1, -1, -1, -13, 6, -35,
	42, -1, -34, 6, -13, -13, -15, -1, -13, -13,
	-13, -13, -1, -1, 9, -13, -43, 10, -13, -13,
	-15, 6, 10, -34, -37, 45, -37, -34, -1, -1,
	-1, -1, -13, -15, -1, -1, 6, 7, 10, 51,
	10, 51, 41, -41, -46, -13, 6, 43, -40, -49,
	8, -22, 6, -34, 32, -34, -1, 6, -34, -34,
	-1, 42, 9, 41, 42, -13, -1, -42, -48, -36,
	8, 9, 9, -13, -5, 51, -13, -15, -5, 13,
	-34, -44, 6, -34, 53, 5, -13, -51, -51, 10,
	4, 43, 7, -51, 10, -49, 34, 13, -21, 13,
	13, -38, 16, 15, 13, 13, 24, -43, 34, 13,
	45, 10, 53, 53, 10, 53, -51, -1, -13, -34,
	-22, 9, -1, -34, 13, 16, 15, -1, -34, -34,
	7, 9, -1, 6, -1, 6, -51, -51, 13, -1,
	-34, -34, -34, 4, 4, 43, 13, -34, 6, -1,
	-13, -34, -51, 10, -51, 10, -51, -51,
}
var RubyDef = []int{

	1, -2, 2, 3, 7, 8, 9, 10, 16, 17,
	-2, 19, 20, 21, 22, 23, 24, 25, 26, 34,
	35, 36, 37, 38, 39, 40, 41, 42, 43, 44,
	45, 46, 0, 0, 120, 121, 63, 5, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 154,
	11, 27, 28, 29, 30, 31, 32, 4, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 63, 51,
	55, 63, 0, 0, 60, 68, 69, 73, 17, -2,
	19, 20, 21, 11, -2, 0, 0, 0, 0, 0,
	0, 0, 101, 102, 0, 99, 100, 0, 0, 16,
	64, 65, 66, 126, 0, 82, 11, -2, 11, 0,
	0, 107, 33, 108, 109, 110, 11, 11, 17, 0,
	155, 156, 161, 52, 56, 111, 113, 115, 116, 117,
	118, 119, 146, 144, 47, 64, 0, 0, 64, 92,
	93, 105, 0, 0, 11, 139, 11, 0, 94, 95,
	96, 97, 112, 114, 145, 147, 103, 104, 0, 122,
	0, 123, 6, 5, 5, 0, 17, 0, 5, 78,
	82, 83, 80, 0, 0, 0, 98, 106, 0, 0,
	157, 158, 159, 12, 13, 14, 15, 0, 53, 54,
	63, 48, 49, 70, 71, 57, 74, 75, 76, 134,
	0, 0, 140, 0, 137, 62, 67, 0, 0, 5,
	0, 0, -2, 11, 0, 0, 0, 85, 11, 87,
	142, 0, 0, 11, 148, 160, 11, 0, 0, 135,
	138, 0, 136, 124, 0, 125, 0, 5, 130, 5,
	84, 79, 81, 0, 143, 0, 11, 11, 153, 162,
	11, 59, 58, 141, 0, 0, 127, 0, 86, 11,
	151, 152, 163, 0, 0, 0, 77, 150, 11, 5,
	5, 164, 128, 5, 132, 5, 129, 133,
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
		//line parser.y:160
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:162
		{
			Statements = []ast.Node{}
		}
	case 3:
		//line parser.y:164
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 4:
		//line parser.y:170
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 5:
		//line parser.y:176
		{
		}
	case 6:
		//line parser.y:177
		{
		}
	case 7:
		//line parser.y:179
		{
			RubyVAL.genericValue = nil
		}
	case 8:
		//line parser.y:180
		{
			RubyVAL.genericValue = nil
		}
	case 9:
		//line parser.y:181
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 10:
		//line parser.y:182
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 11:
		//line parser.y:185
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 12:
		//line parser.y:187
		{
		}
	case 13:
		//line parser.y:189
		{
		}
	case 14:
		//line parser.y:191
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 15:
		//line parser.y:193
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 16:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
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
		//line parser.y:203
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: []ast.Node{},
			}
		}
	case 48:
		//line parser.y:210
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 49:
		//line parser.y:217
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 50:
		//line parser.y:224
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 51:
		//line parser.y:231
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 52:
		//line parser.y:238
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 53:
		//line parser.y:245
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 54:
		//line parser.y:253
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 55:
		//line parser.y:263
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 56:
		//line parser.y:270
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 57:
		//line parser.y:280
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 58:
		//line parser.y:290
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 59:
		//line parser.y:299
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 60:
		//line parser.y:301
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 61:
		//line parser.y:304
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 62:
		//line parser.y:306
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 63:
		//line parser.y:308
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 64:
		//line parser.y:310
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 65:
		//line parser.y:312
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 66:
		//line parser.y:314
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:316
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:320
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 69:
		//line parser.y:322
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 70:
		//line parser.y:324
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 71:
		//line parser.y:326
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 72:
		//line parser.y:329
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:331
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:333
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 75:
		//line parser.y:335
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:337
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:343
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 78:
		//line parser.y:352
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 79:
		//line parser.y:354
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 80:
		//line parser.y:357
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 81:
		//line parser.y:359
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 82:
		//line parser.y:361
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 83:
		//line parser.y:363
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 84:
		//line parser.y:367
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:372
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 86:
		//line parser.y:379
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 87:
		//line parser.y:389
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 88:
		//line parser.y:398
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 89:
		//line parser.y:404
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 90:
		//line parser.y:412
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 91:
		//line parser.y:416
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 92:
		//line parser.y:421
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 93:
		//line parser.y:428
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 94:
		//line parser.y:435
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 95:
		//line parser.y:442
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 96:
		//line parser.y:449
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 97:
		//line parser.y:456
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 98:
		//line parser.y:463
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 99:
		//line parser.y:471
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 100:
		//line parser.y:473
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 101:
		//line parser.y:476
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 102:
		//line parser.y:478
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 103:
		//line parser.y:481
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 104:
		//line parser.y:483
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 105:
		//line parser.y:486
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 106:
		//line parser.y:488
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 107:
		//line parser.y:490
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 108:
		//line parser.y:491
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 109:
		//line parser.y:492
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 110:
		//line parser.y:493
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 111:
		//line parser.y:496
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 112:
		//line parser.y:504
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 113:
		//line parser.y:512
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 114:
		//line parser.y:520
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 115:
		//line parser.y:529
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 116:
		//line parser.y:538
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 117:
		//line parser.y:547
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 118:
		//line parser.y:556
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 119:
		//line parser.y:565
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 120:
		//line parser.y:573
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 121:
		//line parser.y:574
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 122:
		//line parser.y:576
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 123:
		//line parser.y:578
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 124:
		//line parser.y:581
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 125:
		//line parser.y:589
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 126:
		//line parser.y:597
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 127:
		//line parser.y:599
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 128:
		//line parser.y:606
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 129:
		//line parser.y:613
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 130:
		//line parser.y:621
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 131:
		//line parser.y:628
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 132:
		//line parser.y:635
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 133:
		//line parser.y:642
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 134:
		//line parser.y:650
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 135:
		//line parser.y:652
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 136:
		//line parser.y:659
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 137:
		//line parser.y:666
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 138:
		//line parser.y:669
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 139:
		//line parser.y:671
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 140:
		//line parser.y:673
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 141:
		//line parser.y:675
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 142:
		//line parser.y:678
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 143:
		//line parser.y:685
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 144:
		//line parser.y:693
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 145:
		//line parser.y:700
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 146:
		//line parser.y:707
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 147:
		//line parser.y:714
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 148:
		//line parser.y:721
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-2].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 149:
		//line parser.y:728
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 150:
		//line parser.y:730
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 151:
		//line parser.y:737
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 152:
		//line parser.y:744
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 153:
		//line parser.y:751
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 154:
		//line parser.y:758
		{
		}
	case 155:
		//line parser.y:759
		{
		}
	case 156:
		//line parser.y:760
		{
			RubyVAL.genericSlice = []ast.Node{RubyS[Rubypt-0].genericValue}
		}
	case 157:
		//line parser.y:761
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 158:
		//line parser.y:762
		{
		}
	case 159:
		//line parser.y:765
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 160:
		//line parser.y:768
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 161:
		//line parser.y:776
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 162:
		//line parser.y:778
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 163:
		//line parser.y:780
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 164:
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
