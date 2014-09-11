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

//line parser.y:674

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
	43, 81,
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
	-2, 57,
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
	-2, 65,
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
	43, 81,
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
	-2, 52,
	-1, 114,
	10, 13,
	-2, 54,
	-1, 118,
	45, 13,
	-2, 9,
	-1, 132,
	43, 81,
	-2, 79,
	-1, 176,
	43, 82,
	-2, 80,
	-1, 203,
	45, 119,
	-2, 58,
	-1, 232,
	45, 120,
	-2, 18,
	-1, 249,
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
	-2, 67,
	-1, 301,
	45, 121,
	-2, 18,
}

const RubyNprod = 147
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1202

var RubyAct = []int{

	156, 14, 6, 238, 128, 117, 8, 2, 57, 21,
	58, 60, 90, 62, 63, 64, 65, 66, 68, 69,
	294, 292, 7, 67, 257, 203, 59, 198, 308, 112,
	237, 91, 11, 104, 105, 100, 7, 4, 5, 298,
	129, 98, 75, 297, 99, 73, 74, 22, 76, 77,
	78, 79, 80, 81, 82, 83, 87, 42, 43, 247,
	106, 277, 110, 106, 92, 93, 94, 70, 209, 144,
	265, 266, 97, 95, 96, 106, 174, 44, 256, 45,
	218, 41, 40, 183, 89, 141, 7, 131, 133, 143,
	10, 122, 145, 146, 147, 148, 149, 150, 151, 152,
	153, 154, 155, 126, 61, 157, 103, 7, 174, 121,
	158, 162, 106, 163, 120, 164, 165, 265, 266, 167,
	168, 169, 170, 171, 172, 119, 217, 173, 123, 175,
	223, 177, 178, 319, 179, 216, 7, 318, 161, 85,
	160, 242, 7, 244, 243, 106, 116, 106, 106, 278,
	106, 106, 106, 106, 106, 262, 106, 260, 106, 199,
	106, 197, 198, 134, 135, 136, 137, 138, 139, 248,
	198, 196, 227, 198, 106, 195, 289, 106, 132, 106,
	124, 125, 71, 72, 194, 176, 313, 303, 166, 224,
	215, 103, 130, 109, 311, 225, 226, 306, 263, 228,
	229, 255, 46, 230, 103, 106, 103, 114, 180, 101,
	211, 234, 235, 236, 231, 184, 210, 1, 240, 214,
	220, 202, 245, 84, 36, 103, 35, 34, 106, 33,
	106, 187, 251, 233, 258, 254, 32, 193, 264, 259,
	261, 116, 31, 30, 271, 29, 28, 27, 275, 200,
	269, 25, 204, 205, 206, 207, 24, 23, 39, 26,
	279, 88, 281, 37, 283, 280, 38, 282, 9, 19,
	286, 285, 18, 20, 250, 17, 3, 0, 0, 291,
	0, 293, 0, 295, 0, 0, 239, 0, 241, 0,
	300, 0, 246, 106, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 307, 15, 0, 0, 310, 267, 312,
	0, 0, 314, 253, 272, 0, 0, 274, 0, 0,
	102, 0, 0, 111, 0, 276, 11, 104, 105, 273,
	0, 0, 118, 0, 0, 127, 0, 0, 91, 11,
	104, 105, 100, 288, 0, 0, 0, 16, 284, 0,
	0, 42, 43, 0, 0, 287, 0, 299, 0, 0,
	0, 0, 0, 107, 42, 43, 107, 296, 304, 0,
	305, 44, 127, 45, 0, 41, 40, 302, 107, 0,
	0, 0, 0, 316, 44, 0, 45, 0, 41, 40,
	0, 0, 0, 0, 0, 0, 315, 0, 317, 0,
	0, 0, 0, 0, 0, 102, 108, 185, 186, 108,
	188, 189, 190, 191, 192, 107, 127, 0, 102, 0,
	102, 108, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 208, 0, 0, 212, 0, 102,
	0, 0, 0, 0, 0, 0, 0, 0, 107, 0,
	107, 107, 0, 107, 107, 107, 107, 107, 108, 107,
	0, 107, 0, 107, 0, 127, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 107, 0, 0,
	107, 0, 107, 0, 0, 0, 0, 0, 249, 0,
	252, 108, 0, 108, 108, 0, 108, 108, 108, 108,
	108, 0, 108, 0, 108, 0, 108, 0, 107, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	108, 0, 0, 108, 0, 108, 0, 0, 0, 0,
	0, 107, 0, 107, 11, 12, 13, 55, 0, 0,
	0, 47, 219, 54, 222, 221, 0, 48, 49, 0,
	0, 108, 56, 252, 0, 0, 0, 0, 0, 42,
	43, 0, 0, 0, 50, 51, 52, 53, 0, 182,
	181, 0, 0, 0, 108, 0, 108, 0, 0, 44,
	0, 45, 0, 41, 40, 11, 12, 13, 55, 0,
	0, 0, 47, 309, 54, 0, 107, 0, 48, 49,
	0, 0, 0, 56, 0, 0, 0, 0, 0, 0,
	42, 43, 11, 301, 105, 50, 51, 52, 53, 0,
	182, 181, 0, 0, 0, 0, 0, 0, 0, 0,
	44, 0, 45, 0, 41, 40, 0, 42, 43, 108,
	11, 12, 13, 55, 0, 0, 0, 47, 290, 54,
	0, 0, 0, 48, 49, 0, 0, 44, 56, 45,
	0, 41, 40, 0, 0, 42, 43, 0, 0, 0,
	50, 51, 52, 53, 0, 182, 181, 0, 0, 0,
	0, 0, 0, 0, 0, 44, 0, 45, 0, 41,
	40, 11, 12, 13, 55, 0, 0, 0, 47, 270,
	54, 0, 0, 0, 48, 49, 0, 0, 0, 56,
	0, 0, 0, 0, 0, 0, 42, 43, 0, 0,
	0, 50, 51, 52, 53, 0, 182, 181, 0, 0,
	0, 0, 0, 0, 0, 0, 44, 0, 45, 0,
	41, 40, 11, 12, 13, 55, 0, 0, 0, 47,
	268, 54, 0, 0, 0, 48, 49, 0, 0, 0,
	56, 0, 0, 0, 0, 0, 0, 42, 43, 0,
	0, 0, 50, 51, 52, 53, 0, 182, 181, 0,
	0, 0, 0, 0, 0, 0, 0, 44, 0, 45,
	0, 41, 40, 11, 12, 13, 55, 0, 0, 0,
	47, 201, 54, 0, 0, 0, 48, 49, 0, 0,
	0, 56, 0, 0, 0, 0, 0, 0, 42, 43,
	0, 0, 0, 50, 51, 52, 53, 0, 182, 181,
	0, 0, 11, 12, 13, 55, 142, 0, 44, 47,
	45, 54, 41, 40, 0, 48, 49, 0, 0, 0,
	56, 0, 0, 0, 0, 0, 0, 42, 43, 0,
	0, 0, 50, 51, 52, 53, 0, 7, 0, 140,
	0, 11, 12, 13, 55, 0, 0, 44, 47, 45,
	54, 41, 40, 0, 48, 49, 0, 0, 0, 56,
	0, 0, 0, 0, 0, 0, 42, 43, 0, 0,
	0, 50, 51, 52, 53, 0, 182, 181, 0, 0,
	0, 0, 0, 0, 0, 0, 44, 0, 45, 0,
	41, 40, 11, 12, 13, 55, 0, 0, 118, 47,
	0, 54, 0, 0, 0, 48, 49, 0, 0, 0,
	56, 0, 0, 0, 0, 0, 0, 42, 43, 0,
	0, 115, 50, 51, 52, 53, 0, 0, 0, 0,
	0, 11, 12, 13, 55, 0, 0, 44, 47, 45,
	54, 41, 40, 0, 48, 49, 0, 0, 0, 56,
	0, 0, 0, 0, 0, 0, 42, 43, 0, 0,
	0, 50, 51, 52, 53, 0, 7, 0, 86, 0,
	0, 0, 0, 0, 0, 0, 44, 0, 45, 0,
	41, 40, 11, 12, 13, 55, 0, 0, 118, 47,
	0, 54, 0, 0, 0, 48, 49, 0, 0, 0,
	56, 0, 0, 0, 0, 0, 0, 42, 43, 0,
	0, 0, 50, 51, 52, 53, 0, 0, 0, 0,
	0, 11, 12, 13, 55, 0, 0, 44, 47, 45,
	54, 41, 40, 0, 48, 49, 0, 0, 0, 56,
	0, 0, 0, 0, 0, 0, 42, 43, 11, 113,
	105, 50, 51, 52, 53, 11, 104, 105, 100, 0,
	0, 11, 113, 105, 0, 159, 44, 0, 45, 0,
	41, 40, 0, 42, 43, 0, 0, 0, 0, 0,
	42, 43, 0, 7, 0, 0, 42, 43, 11, 113,
	105, 0, 0, 44, 0, 45, 0, 41, 40, 0,
	44, 0, 45, 0, 41, 40, 44, 0, 45, 0,
	41, 40, 0, 42, 43, 11, 232, 105, 0, 0,
	0, 11, 213, 105, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 44, 0, 45, 0, 41, 40, 0,
	42, 43, 0, 0, 0, 0, 42, 43, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	44, 0, 45, 0, 41, 40, 44, 0, 45, 0,
	41, 40,
}
var RubyPact = []int{

	-50, -4, -1000, -51, -1000, -1000, 1046, 46, -1000, -18,
	46, -1000, 96, 46, 46, 46, 46, -1000, -1000, -1000,
	-1000, -1000, 46, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	12, 176, -1000, -1000, 46, 46, -1, 46, 46, 46,
	46, 46, 46, 46, 46, 956, 43, -1000, 27, 187,
	-1000, 1073, 917, 91, 80, 75, 57, 114, -1000, -1000,
	174, -1000, -1000, 1113, -1000, -3, 186, 171, 171, 1046,
	1046, 1046, 1046, 1046, 827, -1000, -1000, -1000, -18, -1000,
	-1000, 46, 46, 46, 46, 46, 46, 46, 46, 46,
	46, 46, -18, -1000, 96, -1000, -1000, -1000, -1000, 46,
	1086, 129, 102, -1000, 46, 46, -1000, -1000, 46, 46,
	46, 46, 46, 46, -1000, -1000, 46, -1000, 67, 178,
	46, 46, -1000, 46, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 27, 866, 1113, 1113, 1046, 1113, 1113,
	1113, 1113, 1113, 1046, 1113, 165, 334, 1007, 1080, -1000,
	46, -1000, -1000, 152, 149, 1046, 788, -20, 1046, 1046,
	1046, 1046, 1113, 17, -1000, 1146, -1000, 1080, 94, 39,
	529, -1000, -1000, -1000, 117, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 46, 46, -1000, 163, 46, 46,
	-1000, -1000, 46, 1140, -1000, -1000, -1000, -1000, -1000, -1000,
	46, 46, 46, -13, -1000, -1000, -1000, 46, -1000, -1000,
	128, 46, -1000, -1000, 35, 160, 321, -1000, 1113, 1007,
	37, -21, -1000, 46, 147, 145, 194, 46, 77, 737,
	171, 686, -1000, 46, -1000, 1046, 866, 46, -1000, -18,
	-1000, -1000, -1000, -1000, -1000, 20, -1000, -1000, 139, 67,
	-1000, 67, -1000, 46, 1046, -1000, -1000, 866, -1000, 46,
	-1000, 1046, 866, -1000, 866, 169, 635, -1000, 46, -32,
	67, -33, 67, 1046, -1000, 30, -2, -1000, 866, 46,
	-1000, 607, -1000, 1046, -1000, 181, -1000, -1000, -1000, 866,
	193, -1000, 46, -15, 580, 866, 46, 190, 46, -1000,
	180, 46, 1046, -1000, 1046, 127, 866, 123, -1000, -1000,
}
var RubyPgo = []int{

	0, 83, 276, 275, 273, 5, 272, 269, 266, 347,
	47, 263, 259, 258, 261, 304, 257, 256, 1, 251,
	9, 247, 246, 245, 243, 242, 236, 229, 227, 226,
	224, 69, 223, 12, 221, 220, 219, 217, 216, 215,
	29, 214, 210, 209, 207, 202, 0, 4, 3, 201,
}
var RubyR1 = []int{

	0, 37, 37, 37, 37, 2, 2, 2, 2, 31,
	31, 31, 31, 46, 46, 47, 47, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 33, 33, 40, 40,
	40, 44, 44, 44, 44, 43, 43, 43, 43, 43,
	48, 48, 48, 16, 36, 36, 17, 17, 19, 20,
	20, 45, 45, 12, 12, 12, 12, 12, 21, 22,
	23, 24, 25, 26, 27, 28, 29, 30, 3, 6,
	7, 4, 4, 38, 38, 38, 38, 42, 42, 42,
	9, 9, 18, 18, 15, 15, 5, 5, 34, 41,
	41, 41, 49, 49, 11, 11, 11, 11, 11, 35,
	35, 35, 35, 35, 32, 32, 32, 32, 32, 32,
	32, 8, 13, 39, 39, 39, 39,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 1, 3, 3, 0,
	2, 2, 2, 0, 2, 0, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 4, 4,
	6, 3, 3, 5, 3, 5, 5, 1, 0, 1,
	5, 1, 1, 5, 5, 1, 1, 5, 5, 5,
	0, 2, 2, 9, 0, 1, 7, 11, 7, 1,
	4, 1, 4, 5, 5, 5, 5, 5, 3, 3,
	3, 3, 5, 5, 5, 5, 5, 5, 1, 1,
	5, 9, 9, 0, 5, 10, 11, 4, 9, 10,
	2, 2, 2, 2, 3, 3, 3, 7, 3, 0,
	1, 5, 1, 2, 5, 6, 5, 5, 5, 0,
	5, 3, 4, 2, 0, 1, 1, 1, 2, 2,
	2, 3, 5, 0, 4, 7, 10,
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
	-46, -14, -40, 6, -44, 34, -1, -5, 11, 34,
	34, 34, 34, 14, 6, 7, -40, -14, -47, 43,
	6, -20, 7, -20, -1, -1, -1, -1, -1, -1,
	42, -46, 9, -46, -31, -46, -46, -46, -46, -46,
	-46, -46, -46, -46, -46, -46, -46, -46, -46, 9,
	-40, 9, 9, -46, -46, -46, -31, -46, -46, -46,
	-46, -46, -46, -46, 41, -46, 7, -46, -46, -46,
	-31, 41, 40, -1, -39, -14, -14, -1, -14, -14,
	-14, -14, -14, -1, -40, 10, -33, -46, 10, 10,
	-1, 13, -34, 45, -1, -1, -1, -1, -14, 51,
	-38, -42, -14, 6, -36, -33, 41, 32, 41, 13,
	-35, 16, 15, 13, -47, -46, -46, 9, -46, -46,
	-46, -41, 6, -40, -46, -46, -46, 43, -48, -31,
	-46, -31, 13, 16, 15, -46, -31, 24, 9, -14,
	-10, -5, -14, -1, -5, -49, 41, 45, -46, -47,
	10, -47, 10, 4, -46, 40, 41, -31, 13, -20,
	13, -46, -31, -1, -31, -46, -31, 41, 10, -46,
	-47, -46, -47, -46, -1, -48, -46, -1, -31, 7,
	13, -46, 53, -46, 53, -46, -1, 13, 41, -31,
	-46, 6, -1, 6, -31, -31, 4, -46, 43, 13,
	-46, 4, -46, 6, -46, -1, -31, -1, 10, 10,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 6, 0, 13, 4, 13,
	13, 17, -2, -2, -2, -2, -2, 23, 24, 25,
	26, 27, -2, 30, 31, 32, 33, 34, 35, 36,
	37, 38, 39, 40, 41, 42, 43, 44, 45, 46,
	0, 0, 98, 99, 13, 13, 0, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 0, 14, 7, 0,
	8, 13, 0, 0, 0, 0, 0, 0, 112, 113,
	0, 110, 111, 58, 15, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 135, 136, 137, -2, 9,
	51, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, -2, -2, 66, -2, -2, 20, 21, 22, -2,
	58, 59, 13, 18, -2, 13, 61, 62, -2, 13,
	13, 13, 13, 13, 114, 115, 13, 59, 13, 0,
	13, 13, -2, 13, 88, 89, 90, 91, 9, 138,
	139, 140, 141, 0, 143, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 58, 0, 0, 0, 0, 47,
	13, 48, 49, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 16, 103, -2, 74, 0, 0,
	0, 10, 11, 12, 15, 55, 92, 93, 94, 95,
	96, 97, 126, 128, 13, 13, 53, 0, 13, 13,
	83, 116, 13, -2, 84, 85, 86, 87, 127, 100,
	13, 13, 13, 18, 70, 75, 9, 13, 9, 124,
	0, 13, 9, 142, 0, 0, 0, 50, 0, 0,
	0, 0, -2, 13, 15, 15, 0, 13, 9, 0,
	0, 0, 125, 13, 9, 0, 133, 9, 56, -2,
	68, 69, 60, 63, 64, 9, 122, 118, 0, 13,
	15, 13, 15, 13, 0, 71, 72, 70, 76, 13,
	78, 0, 131, 9, 144, 0, 0, 123, 13, 0,
	13, 0, 13, 0, 107, 0, 0, 9, 132, 13,
	117, 0, 101, 0, 102, 0, 104, 73, 9, 130,
	9, -2, 13, 0, 0, 145, 13, 0, 13, 77,
	0, 13, 0, 9, 0, 108, 146, 105, 109, 106,
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
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 50:
		//line parser.y:215
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 51:
		//line parser.y:222
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 52:
		//line parser.y:229
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 53:
		//line parser.y:236
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 54:
		//line parser.y:246
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 55:
		//line parser.y:253
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 56:
		//line parser.y:262
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 57:
		//line parser.y:264
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 58:
		//line parser.y:266
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 59:
		//line parser.y:268
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 60:
		//line parser.y:270
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 61:
		//line parser.y:274
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 62:
		//line parser.y:276
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 63:
		//line parser.y:278
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 64:
		//line parser.y:280
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 65:
		//line parser.y:283
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 66:
		//line parser.y:285
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:287
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:289
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 69:
		//line parser.y:291
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:301
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 74:
		//line parser.y:309
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 75:
		//line parser.y:310
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 76:
		//line parser.y:313
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 77:
		//line parser.y:320
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-8].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-4].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-8].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 78:
		//line parser.y:330
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 79:
		//line parser.y:339
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 80:
		//line parser.y:345
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 81:
		//line parser.y:353
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 82:
		//line parser.y:357
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
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
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 88:
		//line parser.y:397
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 89:
		//line parser.y:398
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 90:
		//line parser.y:399
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 91:
		//line parser.y:400
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 92:
		//line parser.y:403
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 93:
		//line parser.y:412
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 94:
		//line parser.y:421
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 95:
		//line parser.y:430
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 96:
		//line parser.y:439
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 97:
		//line parser.y:448
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 98:
		//line parser.y:456
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 99:
		//line parser.y:457
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 100:
		//line parser.y:459
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 101:
		//line parser.y:463
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 102:
		//line parser.y:471
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 103:
		//line parser.y:479
		{
			RubyVAL.genericSlice = ast.Nodes{}
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
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 106:
		//line parser.y:495
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-1].genericValue})
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
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 109:
		//line parser.y:517
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 110:
		//line parser.y:525
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 111:
		//line parser.y:527
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 112:
		//line parser.y:530
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 113:
		//line parser.y:532
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 114:
		//line parser.y:535
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 115:
		//line parser.y:537
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 116:
		//line parser.y:540
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 117:
		//line parser.y:542
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 118:
		//line parser.y:550
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 119:
		//line parser.y:552
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 120:
		//line parser.y:554
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 121:
		//line parser.y:556
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 124:
		//line parser.y:561
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 125:
		//line parser.y:568
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
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
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 128:
		//line parser.y:590
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 129:
		//line parser.y:597
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 130:
		//line parser.y:599
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 131:
		//line parser.y:606
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 132:
		//line parser.y:613
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 133:
		//line parser.y:620
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 134:
		//line parser.y:627
		{
		}
	case 135:
		//line parser.y:628
		{
			RubyVAL.genericSlice = []ast.Node{RubyS[Rubypt-0].genericValue}
		}
	case 136:
		//line parser.y:629
		{
		}
	case 137:
		//line parser.y:630
		{
		}
	case 138:
		//line parser.y:631
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 139:
		//line parser.y:632
		{
		}
	case 140:
		//line parser.y:633
		{
		}
	case 141:
		//line parser.y:636
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 142:
		//line parser.y:639
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 143:
		//line parser.y:647
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 144:
		//line parser.y:649
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 145:
		//line parser.y:651
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				},
			})
		}
	case 146:
		//line parser.y:660
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
