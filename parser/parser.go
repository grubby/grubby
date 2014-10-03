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

//line parser.y:822

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 25,
	43, 94,
	-2, 20,
	-1, 79,
	9, 66,
	10, 66,
	-2, 159,
	-1, 90,
	43, 94,
	-2, 20,
	-1, 95,
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
	-1, 105,
	43, 94,
	-2, 92,
	-1, 116,
	43, 94,
	-2, 20,
	-1, 203,
	9, 66,
	10, 66,
	-2, 159,
	-1, 240,
	43, 95,
	-2, 93,
}

const RubyNprod = 168
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1277

var RubyAct = []int{

	9, 171, 22, 33, 81, 169, 125, 80, 87, 168,
	195, 2, 3, 159, 62, 119, 120, 160, 193, 78,
	191, 262, 285, 239, 200, 86, 4, 131, 195, 245,
	108, 242, 221, 101, 231, 100, 99, 26, 73, 104,
	106, 60, 59, 175, 114, 232, 261, 23, 89, 90,
	79, 281, 84, 94, 107, 62, 61, 127, 194, 128,
	192, 64, 91, 121, 244, 134, 135, 10, 138, 139,
	140, 141, 54, 55, 144, 145, 83, 195, 72, 195,
	127, 91, 128, 153, 154, 128, 151, 77, 260, 73,
	86, 82, 88, 95, 91, 53, 52, 278, 127, 292,
	128, 63, 91, 91, 163, 91, 91, 91, 91, 161,
	97, 91, 91, 173, 152, 73, 130, 91, 280, 186,
	91, 91, 74, 237, 129, 75, 290, 91, 76, 72,
	62, 198, 195, 136, 86, 91, 98, 253, 201, 255,
	254, 202, 146, 219, 284, 157, 195, 129, 74, 97,
	129, 155, 259, 205, 62, 72, 240, 88, 208, 186,
	172, 206, 170, 186, 105, 129, 210, 283, 91, 266,
	193, 91, 249, 219, 268, 186, 220, 186, 172, 224,
	186, 186, 217, 193, 204, 193, 189, 190, 122, 123,
	91, 103, 214, 102, 234, 91, 178, 156, 27, 150,
	143, 88, 133, 64, 127, 233, 128, 276, 238, 58,
	241, 85, 197, 186, 124, 213, 186, 188, 196, 1,
	227, 248, 117, 92, 51, 209, 50, 49, 48, 47,
	46, 91, 18, 17, 16, 28, 65, 66, 67, 15,
	265, 91, 92, 63, 70, 68, 69, 37, 186, 207,
	13, 12, 186, 186, 11, 92, 21, 14, 186, 186,
	93, 19, 31, 92, 92, 30, 92, 92, 92, 92,
	32, 129, 92, 92, 186, 186, 186, 91, 92, 93,
	29, 92, 92, 186, 0, 218, 288, 186, 92, 0,
	222, 0, 93, 0, 0, 187, 92, 5, 0, 0,
	93, 93, 0, 93, 93, 93, 93, 64, 0, 93,
	93, 0, 0, 235, 236, 93, 118, 0, 93, 93,
	71, 0, 0, 91, 0, 93, 0, 0, 0, 92,
	0, 0, 92, 93, 109, 110, 111, 112, 113, 0,
	65, 66, 67, 0, 0, 0, 0, 63, 70, 68,
	69, 92, 0, 0, 263, 0, 92, 0, 132, 0,
	20, 0, 137, 0, 269, 0, 93, 142, 270, 93,
	0, 147, 148, 149, 0, 0, 0, 0, 0, 0,
	0, 279, 0, 0, 0, 96, 0, 0, 93, 0,
	0, 0, 92, 93, 164, 165, 166, 167, 0, 0,
	0, 0, 92, 177, 289, 291, 0, 293, 0, 294,
	0, 158, 162, 181, 64, 0, 0, 96, 0, 0,
	0, 174, 0, 176, 0, 0, 96, 0, 0, 93,
	179, 180, 0, 0, 0, 96, 0, 0, 92, 93,
	96, 0, 0, 96, 96, 0, 0, 65, 66, 67,
	96, 0, 0, 0, 63, 70, 68, 69, 96, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 93, 212, 0, 215, 23,
	24, 25, 44, 0, 92, 0, 34, 226, 42, 229,
	228, 43, 35, 36, 96, 0, 0, 45, 0, 0,
	0, 0, 0, 0, 54, 55, 0, 0, 0, 38,
	39, 40, 41, 0, 0, 184, 185, 250, 96, 0,
	0, 93, 0, 56, 256, 57, 0, 53, 52, 0,
	0, 0, 0, 0, 264, 247, 0, 0, 267, 251,
	0, 252, 0, 0, 0, 0, 257, 0, 0, 258,
	272, 0, 0, 0, 0, 0, 0, 277, 0, 0,
	0, 0, 0, 0, 96, 0, 0, 0, 0, 0,
	0, 0, 273, 274, 0, 0, 275, 0, 0, 0,
	287, 23, 24, 25, 44, 0, 0, 0, 34, 282,
	42, 0, 0, 43, 35, 36, 0, 0, 0, 45,
	286, 0, 0, 0, 0, 0, 54, 55, 0, 0,
	0, 38, 39, 40, 41, 0, 0, 6, 7, 0,
	0, 0, 0, 0, 0, 56, 0, 57, 0, 53,
	52, 0, 8, 23, 24, 25, 44, 0, 0, 0,
	34, 271, 42, 0, 0, 43, 35, 36, 0, 0,
	0, 45, 0, 0, 0, 0, 0, 0, 54, 55,
	0, 0, 0, 38, 39, 40, 41, 0, 0, 184,
	185, 23, 24, 25, 44, 0, 0, 56, 34, 57,
	42, 53, 52, 43, 35, 36, 0, 0, 0, 45,
	0, 0, 0, 0, 0, 0, 54, 55, 0, 0,
	0, 38, 39, 40, 41, 0, 0, 184, 185, 0,
	0, 0, 0, 0, 0, 56, 0, 57, 246, 53,
	52, 23, 24, 25, 44, 0, 0, 0, 34, 243,
	42, 0, 0, 43, 35, 36, 0, 0, 0, 45,
	0, 0, 0, 0, 0, 0, 54, 55, 0, 0,
	0, 38, 39, 40, 41, 0, 0, 184, 185, 0,
	0, 0, 0, 0, 0, 56, 0, 57, 0, 53,
	52, 23, 24, 25, 44, 0, 0, 0, 34, 230,
	42, 0, 0, 43, 35, 36, 0, 0, 0, 45,
	0, 0, 0, 0, 0, 0, 54, 55, 0, 0,
	0, 38, 39, 40, 41, 0, 0, 184, 185, 0,
	0, 0, 0, 0, 0, 56, 0, 57, 0, 53,
	52, 23, 24, 25, 44, 0, 0, 0, 34, 225,
	42, 0, 0, 43, 35, 36, 0, 0, 0, 45,
	0, 0, 0, 0, 0, 0, 54, 55, 0, 0,
	0, 38, 39, 40, 41, 0, 0, 184, 185, 0,
	0, 0, 0, 0, 0, 56, 0, 57, 0, 53,
	52, 23, 24, 25, 44, 0, 0, 0, 34, 223,
	42, 0, 0, 43, 35, 36, 0, 0, 0, 45,
	0, 0, 0, 0, 0, 0, 54, 55, 0, 0,
	0, 38, 39, 40, 41, 0, 0, 184, 185, 23,
	24, 25, 44, 0, 0, 56, 34, 57, 42, 53,
	52, 43, 35, 36, 0, 0, 0, 45, 0, 0,
	0, 0, 0, 0, 54, 55, 0, 0, 0, 38,
	39, 40, 41, 0, 0, 184, 185, 0, 0, 0,
	0, 0, 0, 56, 0, 57, 216, 53, 52, 23,
	24, 25, 44, 0, 0, 0, 34, 211, 42, 0,
	0, 43, 35, 36, 0, 0, 0, 45, 0, 0,
	0, 0, 0, 0, 54, 55, 0, 0, 0, 38,
	39, 40, 41, 0, 0, 184, 185, 23, 24, 25,
	44, 0, 0, 56, 34, 57, 42, 53, 52, 43,
	35, 36, 0, 0, 0, 45, 0, 0, 0, 0,
	0, 0, 54, 55, 0, 0, 0, 38, 39, 40,
	41, 0, 0, 184, 185, 23, 24, 25, 44, 183,
	0, 56, 34, 57, 42, 53, 52, 43, 35, 36,
	0, 0, 0, 45, 0, 0, 0, 0, 0, 0,
	54, 55, 0, 0, 0, 38, 39, 40, 41, 0,
	0, 0, 182, 23, 24, 25, 44, 0, 0, 56,
	34, 57, 42, 53, 52, 43, 35, 36, 0, 0,
	0, 45, 0, 0, 0, 0, 0, 0, 54, 55,
	0, 0, 0, 38, 39, 40, 41, 23, 89, 90,
	44, 0, 0, 94, 0, 0, 0, 56, 0, 57,
	0, 53, 52, 23, 89, 90, 203, 0, 0, 94,
	0, 0, 54, 55, 23, 89, 90, 79, 0, 0,
	94, 0, 0, 0, 0, 0, 0, 0, 54, 55,
	0, 56, 0, 95, 0, 53, 52, 0, 0, 54,
	55, 23, 89, 90, 44, 0, 0, 56, 0, 95,
	0, 53, 52, 23, 115, 116, 0, 0, 82, 94,
	95, 0, 53, 52, 0, 0, 54, 55, 23, 199,
	116, 0, 0, 0, 0, 0, 0, 0, 54, 55,
	126, 89, 90, 44, 0, 56, 0, 57, 0, 53,
	52, 0, 0, 54, 55, 0, 0, 56, 0, 95,
	0, 53, 52, 0, 195, 54, 55, 23, 115, 116,
	0, 0, 56, 0, 57, 0, 53, 52, 0, 0,
	0, 0, 0, 0, 56, 0, 57, 0, 53, 52,
	0, 0, 54, 55, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 56, 0, 57, 0, 53, 52,
}
var RubyPact = []int{

	-30, 576, -1000, -1000, -1000, 0, -1000, -1000, -1000, 303,
	111, -1000, -1000, -1000, 70, -1000, -1000, -1000, -1000, -1000,
	-25, -1000, -1000, -1000, 42, 102, 2, 1, -1, -1000,
	-1000, -1000, -1000, -1000, 187, 157, 157, 20, 1068, 1068,
	1068, 1068, 1068, 1222, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 9, 182, -1000, -1000, 1195, -1000, -16, -1000,
	-1000, -1000, 1068, 196, 1222, 1156, 1068, 1222, 1222, 1222,
	1222, 1068, 194, 1222, 1156, 1068, 1068, 1068, 193, 1156,
	-1000, 104, 1195, 1156, 191, 135, 57, -1000, 34, 1129,
	141, -1000, -1000, -1000, -28, -28, -25, 1156, 1068, 1068,
	1068, 1068, 154, 154, 11, -1000, -1000, 1068, 190, 41,
	41, 41, 41, 41, -1000, -1000, -1000, 1030, 992, -1000,
	-1000, 180, -1000, -1000, 10, 8, -1000, 410, -1000, 85,
	1183, -19, 41, 1118, -1000, 57, 34, 41, -1000, -1000,
	-1000, -1000, 41, -1000, -1000, 57, 34, 41, 41, 41,
	-1000, 175, 1168, 199, 57, 34, -1000, 1102, 954, -1000,
	186, -1000, 904, 173, 41, 41, 41, 41, -1000, 133,
	172, -1000, -2, -1000, 866, 157, 816, 41, -1000, 474,
	766, 41, -1000, -1000, -1000, -1000, 303, 41, 21, -1000,
	-1000, 200, -1000, 1222, -1000, -1000, -1000, 113, 204, -20,
	149, 104, -1000, 1156, -1000, -1000, -1000, -3, 57, 34,
	-1000, -1000, 716, 19, -1000, 666, -1000, -1000, -13, 172,
	163, 1068, -13, -1000, -1000, -1000, -1000, 124, 1068, -1000,
	-1000, -1000, 145, -1000, -1000, 36, -31, -1000, 1068, 1222,
	-1000, 160, 1068, -1000, -1000, 168, -1000, 992, -1000, -1000,
	41, 992, 628, -1000, 1068, -1000, 41, 992, 992, 203,
	-1000, 1068, -1000, 91, 41, -1000, -1000, 41, -1000, 105,
	38, -1000, 41, 992, 992, 992, 161, 140, -21, -13,
	-1000, -1000, 992, -1000, 1068, 1222, 992, 116, 89, -13,
	-1000, -13, -1000, -13, -13,
}
var RubyPgo = []int{

	0, 295, 280, 270, 8, 265, 262, 360, 235, 261,
	257, 256, 0, 198, 67, 254, 2, 251, 37, 250,
	3, 1, 247, 239, 234, 233, 232, 230, 229, 228,
	227, 226, 224, 316, 222, 7, 13, 220, 219, 9,
	218, 217, 6, 215, 214, 212, 211, 4, 5, 209,
	116,
}
var RubyR1 = []int{

	0, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 50, 50, 33, 33, 33, 33, 33, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 16,
	16, 16, 16, 16, 16, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 35, 35, 44, 44, 42, 42, 42, 42,
	42, 47, 47, 47, 47, 46, 46, 46, 46, 46,
	15, 15, 39, 39, 21, 21, 48, 48, 48, 17,
	17, 19, 20, 20, 49, 49, 10, 10, 10, 10,
	10, 10, 10, 8, 8, 18, 18, 13, 13, 22,
	22, 23, 24, 25, 26, 27, 27, 27, 27, 28,
	29, 30, 31, 32, 2, 5, 6, 6, 3, 3,
	40, 40, 40, 40, 45, 45, 45, 45, 4, 4,
	4, 4, 36, 43, 43, 43, 9, 9, 9, 9,
	9, 9, 9, 9, 37, 37, 37, 37, 37, 34,
	34, 34, 7, 11, 41, 41, 41, 41,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 4,
	4, 2, 3, 3, 4, 4, 3, 2, 3, 3,
	4, 6, 3, 1, 1, 3, 0, 1, 1, 1,
	3, 1, 1, 3, 3, 1, 1, 3, 3, 3,
	7, 7, 1, 3, 1, 3, 0, 1, 3, 4,
	6, 4, 1, 4, 1, 4, 3, 3, 3, 3,
	3, 3, 3, 2, 2, 2, 2, 3, 3, 3,
	3, 2, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 1, 1, 3, 3, 5, 5,
	0, 4, 7, 8, 3, 3, 7, 8, 3, 4,
	4, 3, 3, 0, 1, 3, 4, 5, 3, 3,
	3, 3, 3, 4, 0, 4, 3, 3, 2, 0,
	2, 2, 3, 4, 0, 3, 4, 6,
}
var RubyChk = []int{

	-1000, -38, 41, 42, 56, -1, 41, 42, 56, -12,
	-14, -15, -17, -19, -10, -23, -24, -25, -26, -9,
	-7, -11, -16, 5, 6, 7, -18, -13, -8, -2,
	-5, -6, -3, -20, 12, 18, 19, -22, 35, 36,
	37, 38, 14, 17, 8, 23, -27, -28, -29, -30,
	-31, -32, 54, 53, 30, 31, 49, 51, -49, 42,
	41, 56, 14, 44, 4, 37, 38, 39, 46, 47,
	45, 17, 44, 4, 37, 14, 17, 17, 44, 8,
	-35, -47, 49, 34, 10, -46, -12, -4, -14, 6,
	7, -18, -13, -8, 11, 51, -7, 8, 34, 34,
	34, 34, 6, 4, -20, 7, -20, 34, 10, -1,
	-1, -1, -1, -1, -12, 6, 7, -34, -33, 6,
	7, 54, 6, 7, -44, -42, 5, -12, -16, -14,
	-50, 43, -1, 6, -12, -12, -14, -1, -12, -12,
	-12, -12, -1, 6, -12, -12, -14, -1, -1, -1,
	6, -42, 10, -12, -12, -14, 6, 10, -33, -36,
	45, -36, -33, -42, -1, -1, -1, -1, -39, -48,
	8, -21, 6, -39, -33, 32, -33, -1, 6, -33,
	-33, -1, 42, 9, 41, 42, -12, -1, -41, 6,
	7, 10, 50, 10, 50, 41, -40, -45, -12, 6,
	43, -47, -35, 8, 9, -12, -4, 50, -12, -14,
	-4, 13, -33, -43, 6, -33, 52, 9, -50, 10,
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
	0, 0, 0, 0, 159, 13, 29, 30, 31, 32,
	33, 34, 0, 0, 124, 125, 66, 11, 0, 5,
	6, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, -2,
	51, 57, 66, 0, 0, 63, 71, 72, 76, 19,
	-2, 21, 22, 23, 13, -2, 0, 66, 0, 0,
	0, 0, 86, 86, 13, -2, 13, 0, 0, 111,
	112, 113, 114, 13, 13, 19, -2, 0, 164, 105,
	106, 0, 103, 104, 0, 0, 18, 67, 68, 69,
	130, 0, 148, 52, 58, 115, 117, 119, 120, 121,
	122, 123, 150, 53, 59, 116, 118, 149, 151, 152,
	56, 0, 0, 67, 96, 97, 109, 0, 0, 13,
	143, 13, 0, 0, 98, 99, 100, 101, 11, 82,
	86, 87, 84, 11, 0, 0, 0, 102, 110, 0,
	0, 160, 161, 162, 14, 15, 16, 17, 0, 107,
	108, 0, 126, 0, 127, 12, 11, 11, 0, 19,
	0, 54, 55, -2, 49, 73, 74, 60, 77, 78,
	79, 138, 0, 0, 144, 0, 141, 50, 13, 0,
	0, 0, 13, 89, 13, 91, 146, 0, 0, 13,
	153, 163, 13, 65, 70, 0, 0, 11, 0, 0,
	-2, 0, 0, 139, 142, 0, 140, 11, 88, 83,
	85, 11, 0, 147, 0, 13, 13, 158, 165, 13,
	128, 0, 129, 0, 11, 134, 62, 61, 145, 0,
	0, 90, 13, 156, 157, 166, 0, 0, 0, 131,
	80, 81, 155, 13, 0, 0, 167, 11, 11, 132,
	11, 136, 11, 133, 137,
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
		//line parser.y:276
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 60:
		//line parser.y:286
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 61:
		//line parser.y:296
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 62:
		//line parser.y:305
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 63:
		//line parser.y:307
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
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
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 67:
		//line parser.y:316
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:318
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 69:
		//line parser.y:320
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 70:
		//line parser.y:322
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 71:
		//line parser.y:326
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 72:
		//line parser.y:328
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:330
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:332
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
		//line parser.y:339
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 78:
		//line parser.y:341
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 79:
		//line parser.y:343
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 80:
		//line parser.y:349
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 81:
		//line parser.y:357
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 82:
		//line parser.y:366
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 83:
		//line parser.y:368
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 84:
		//line parser.y:371
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 85:
		//line parser.y:373
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 86:
		//line parser.y:375
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 87:
		//line parser.y:377
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 88:
		//line parser.y:381
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 89:
		//line parser.y:386
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 90:
		//line parser.y:393
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 91:
		//line parser.y:403
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 92:
		//line parser.y:412
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 93:
		//line parser.y:418
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 94:
		//line parser.y:426
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 95:
		//line parser.y:430
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
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
		//line parser.y:470
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 102:
		//line parser.y:477
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 103:
		//line parser.y:485
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 104:
		//line parser.y:487
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 105:
		//line parser.y:490
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 106:
		//line parser.y:492
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 107:
		//line parser.y:495
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 108:
		//line parser.y:497
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 109:
		//line parser.y:500
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 110:
		//line parser.y:502
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 111:
		//line parser.y:504
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 112:
		//line parser.y:505
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 113:
		//line parser.y:506
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 114:
		//line parser.y:507
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
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
		//line parser.y:534
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 119:
		//line parser.y:543
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 120:
		//line parser.y:552
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 121:
		//line parser.y:561
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 122:
		//line parser.y:570
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 123:
		//line parser.y:579
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 124:
		//line parser.y:587
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 125:
		//line parser.y:588
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 126:
		//line parser.y:590
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 127:
		//line parser.y:592
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
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
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 130:
		//line parser.y:611
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 131:
		//line parser.y:613
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 132:
		//line parser.y:620
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 133:
		//line parser.y:627
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 134:
		//line parser.y:635
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 135:
		//line parser.y:642
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 136:
		//line parser.y:649
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 137:
		//line parser.y:656
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 138:
		//line parser.y:664
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 139:
		//line parser.y:666
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 140:
		//line parser.y:673
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 141:
		//line parser.y:680
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 142:
		//line parser.y:683
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 143:
		//line parser.y:685
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 144:
		//line parser.y:687
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 145:
		//line parser.y:689
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 146:
		//line parser.y:692
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 147:
		//line parser.y:699
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 148:
		//line parser.y:707
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 149:
		//line parser.y:714
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 150:
		//line parser.y:721
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 151:
		//line parser.y:728
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 152:
		//line parser.y:735
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 153:
		//line parser.y:742
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-2].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 154:
		//line parser.y:749
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 155:
		//line parser.y:751
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 156:
		//line parser.y:758
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 157:
		//line parser.y:765
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 158:
		//line parser.y:772
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 159:
		//line parser.y:779
		{
		}
	case 160:
		//line parser.y:780
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 161:
		//line parser.y:781
		{
		}
	case 162:
		//line parser.y:784
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 163:
		//line parser.y:787
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 164:
		//line parser.y:795
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 165:
		//line parser.y:797
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 166:
		//line parser.y:799
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 167:
		//line parser.y:808
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
