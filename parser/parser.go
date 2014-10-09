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

//line parser.y:810

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 26,
	44, 95,
	-2, 20,
	-1, 81,
	9, 67,
	10, 67,
	-2, 157,
	-1, 92,
	44, 95,
	-2, 20,
	-1, 97,
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
	-1, 108,
	44, 95,
	-2, 93,
	-1, 119,
	44, 95,
	-2, 20,
	-1, 206,
	9, 67,
	10, 67,
	-2, 157,
	-1, 242,
	44, 96,
	-2, 94,
}

const RubyNprod = 167
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1192

var RubyAct = []int{

	9, 172, 174, 34, 90, 83, 22, 82, 64, 171,
	2, 3, 198, 66, 162, 163, 247, 196, 128, 80,
	89, 194, 287, 264, 241, 4, 88, 203, 66, 122,
	123, 263, 133, 198, 75, 244, 62, 61, 27, 223,
	107, 109, 79, 104, 280, 117, 178, 67, 68, 69,
	283, 63, 246, 103, 65, 72, 70, 71, 197, 129,
	210, 102, 195, 198, 93, 130, 282, 136, 137, 65,
	140, 141, 142, 143, 262, 74, 146, 147, 124, 198,
	198, 64, 154, 93, 75, 156, 158, 28, 130, 292,
	157, 130, 88, 64, 100, 198, 132, 93, 268, 196,
	153, 154, 24, 118, 119, 93, 93, 130, 93, 93,
	93, 93, 164, 94, 93, 93, 176, 111, 76, 166,
	93, 101, 189, 93, 93, 74, 233, 294, 56, 57,
	93, 155, 94, 201, 286, 239, 88, 234, 221, 93,
	100, 204, 110, 205, 64, 160, 94, 58, 175, 59,
	173, 55, 54, 261, 94, 94, 208, 94, 94, 94,
	94, 158, 189, 94, 94, 211, 189, 251, 221, 94,
	242, 93, 94, 94, 93, 222, 209, 108, 189, 94,
	189, 212, 226, 189, 189, 219, 196, 255, 94, 257,
	256, 207, 196, 285, 93, 235, 66, 236, 60, 93,
	192, 193, 24, 91, 92, 45, 270, 154, 96, 73,
	125, 126, 106, 130, 105, 189, 175, 121, 189, 216,
	94, 181, 66, 94, 250, 243, 159, 152, 56, 57,
	67, 68, 69, 93, 145, 93, 29, 65, 72, 70,
	71, 135, 267, 94, 278, 93, 240, 58, 94, 97,
	189, 55, 54, 87, 189, 189, 67, 68, 69, 75,
	189, 189, 95, 65, 72, 70, 71, 200, 220, 77,
	127, 215, 78, 224, 191, 199, 189, 189, 189, 1,
	93, 95, 94, 229, 94, 189, 120, 52, 290, 189,
	51, 50, 49, 76, 94, 95, 237, 238, 48, 190,
	74, 5, 47, 95, 95, 18, 95, 95, 95, 95,
	17, 16, 95, 95, 161, 165, 15, 38, 95, 13,
	12, 95, 95, 23, 11, 177, 93, 179, 95, 94,
	21, 14, 19, 32, 182, 183, 265, 95, 31, 112,
	113, 114, 115, 116, 33, 30, 271, 0, 0, 0,
	272, 0, 0, 0, 0, 10, 0, 0, 0, 0,
	0, 0, 0, 281, 134, 0, 0, 0, 139, 95,
	0, 0, 95, 144, 0, 94, 0, 149, 150, 151,
	214, 98, 217, 0, 0, 0, 291, 293, 0, 295,
	0, 296, 95, 0, 0, 0, 0, 95, 0, 0,
	0, 167, 168, 169, 170, 0, 0, 0, 0, 0,
	180, 0, 0, 20, 131, 0, 0, 0, 0, 0,
	184, 0, 0, 138, 24, 91, 92, 81, 0, 86,
	96, 95, 148, 95, 0, 0, 0, 131, 249, 99,
	131, 98, 253, 95, 254, 0, 0, 98, 0, 259,
	56, 57, 260, 0, 85, 0, 131, 24, 91, 92,
	206, 0, 0, 96, 0, 0, 0, 0, 0, 84,
	0, 97, 99, 55, 54, 275, 276, 0, 95, 277,
	0, 99, 0, 56, 57, 0, 24, 91, 92, 45,
	99, 98, 284, 0, 0, 99, 0, 0, 99, 99,
	0, 0, 58, 288, 97, 99, 55, 54, 0, 0,
	0, 0, 56, 57, 99, 0, 98, 0, 0, 0,
	0, 0, 0, 252, 95, 0, 0, 0, 0, 0,
	258, 58, 0, 59, 0, 55, 54, 0, 0, 0,
	266, 0, 0, 0, 269, 0, 0, 0, 0, 99,
	0, 0, 0, 0, 0, 0, 274, 24, 25, 26,
	45, 0, 131, 279, 35, 228, 43, 231, 230, 44,
	36, 37, 0, 0, 99, 46, 0, 0, 0, 0,
	0, 0, 53, 56, 57, 0, 289, 0, 39, 40,
	41, 42, 0, 0, 187, 188, 24, 25, 26, 45,
	0, 0, 58, 35, 59, 43, 55, 54, 44, 36,
	37, 0, 0, 0, 46, 0, 0, 0, 0, 0,
	99, 53, 56, 57, 0, 0, 0, 39, 40, 41,
	42, 0, 0, 6, 7, 0, 0, 0, 0, 0,
	0, 58, 0, 59, 0, 55, 54, 0, 8, 24,
	25, 26, 45, 0, 0, 0, 35, 273, 43, 0,
	0, 44, 36, 37, 0, 0, 0, 46, 0, 0,
	0, 0, 0, 0, 53, 56, 57, 0, 0, 0,
	39, 40, 41, 42, 0, 0, 187, 188, 24, 25,
	26, 45, 0, 0, 58, 35, 59, 43, 55, 54,
	44, 36, 37, 0, 0, 0, 46, 0, 0, 0,
	0, 0, 0, 53, 56, 57, 0, 0, 0, 39,
	40, 41, 42, 0, 0, 187, 188, 0, 0, 0,
	0, 0, 0, 58, 0, 59, 248, 55, 54, 24,
	25, 26, 45, 0, 0, 0, 35, 245, 43, 0,
	0, 44, 36, 37, 0, 0, 0, 46, 0, 0,
	0, 0, 0, 0, 53, 56, 57, 0, 0, 0,
	39, 40, 41, 42, 0, 0, 187, 188, 0, 0,
	0, 0, 0, 0, 58, 0, 59, 0, 55, 54,
	24, 25, 26, 45, 0, 0, 0, 35, 232, 43,
	0, 0, 44, 36, 37, 0, 0, 0, 46, 0,
	0, 0, 0, 0, 0, 53, 56, 57, 0, 0,
	0, 39, 40, 41, 42, 0, 0, 187, 188, 0,
	0, 0, 0, 0, 0, 58, 0, 59, 0, 55,
	54, 24, 25, 26, 45, 0, 0, 0, 35, 227,
	43, 0, 0, 44, 36, 37, 0, 0, 0, 46,
	0, 0, 0, 0, 0, 0, 53, 56, 57, 0,
	0, 0, 39, 40, 41, 42, 0, 0, 187, 188,
	0, 0, 0, 0, 0, 0, 58, 0, 59, 0,
	55, 54, 24, 25, 26, 45, 0, 0, 0, 35,
	225, 43, 0, 0, 44, 36, 37, 0, 0, 0,
	46, 0, 0, 0, 0, 0, 0, 53, 56, 57,
	0, 0, 0, 39, 40, 41, 42, 0, 0, 187,
	188, 24, 25, 26, 45, 0, 0, 58, 35, 59,
	43, 55, 54, 44, 36, 37, 0, 0, 0, 46,
	0, 0, 0, 0, 0, 0, 53, 56, 57, 0,
	0, 0, 39, 40, 41, 42, 0, 0, 187, 188,
	0, 0, 0, 0, 0, 0, 58, 0, 59, 218,
	55, 54, 24, 25, 26, 45, 0, 0, 0, 35,
	213, 43, 0, 0, 44, 36, 37, 0, 0, 0,
	46, 0, 0, 0, 0, 0, 0, 53, 56, 57,
	0, 0, 0, 39, 40, 41, 42, 0, 0, 187,
	188, 24, 25, 26, 45, 0, 0, 58, 35, 59,
	43, 55, 54, 44, 36, 37, 0, 0, 0, 46,
	0, 0, 0, 0, 0, 0, 53, 56, 57, 0,
	0, 0, 39, 40, 41, 42, 0, 0, 187, 188,
	24, 25, 26, 45, 186, 0, 58, 35, 59, 43,
	55, 54, 44, 36, 37, 0, 0, 0, 46, 0,
	0, 0, 0, 0, 0, 53, 56, 57, 0, 0,
	0, 39, 40, 41, 42, 0, 0, 0, 185, 24,
	25, 26, 45, 0, 0, 58, 35, 59, 43, 55,
	54, 44, 36, 37, 0, 0, 0, 46, 0, 0,
	0, 0, 0, 0, 53, 56, 57, 24, 202, 119,
	39, 40, 41, 42, 24, 91, 92, 81, 0, 0,
	96, 24, 118, 119, 58, 0, 59, 96, 55, 54,
	0, 0, 0, 56, 57, 0, 0, 0, 0, 0,
	56, 57, 0, 0, 198, 0, 0, 56, 57, 0,
	0, 0, 58, 0, 59, 0, 55, 54, 0, 84,
	0, 97, 0, 55, 54, 0, 58, 0, 97, 0,
	55, 54,
}
var RubyPact = []int{

	-32, 591, -1000, -1000, -1000, -6, -1000, -1000, -1000, 192,
	255, -1000, -1000, -1000, 25, -1000, -1000, -1000, -1000, -1000,
	-26, -1000, -1000, -1000, -1000, 419, 86, 26, 18, 8,
	-1000, -1000, -1000, -1000, -1000, 208, 170, 170, 107, 1094,
	1094, 1094, 1094, 1094, 97, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 23, 204, -1000, -1000, 481, -1000,
	-12, -1000, -1000, -1000, 1094, 235, 97, 481, 1094, 97,
	97, 97, 97, 1094, 228, 97, 481, 1094, 1094, 1094,
	221, 481, -1000, 121, 481, 481, 220, 135, 24, -1000,
	-1000, 1129, 132, -1000, -1000, -1000, -31, -31, 30, -26,
	481, 1094, 1094, 1094, 1094, 142, 142, 13, -1000, -1000,
	1094, 215, 67, 67, 67, 67, 67, -1000, -1000, -1000,
	1055, 1016, -1000, -1000, 194, -1000, -1000, 11, 7, 218,
	-1000, 80, 1122, -17, 67, 452, -1000, 24, 30, 67,
	-1000, -1000, -1000, -1000, 67, -1000, -1000, 24, 30, 67,
	67, 67, -1000, 182, 218, 1136, 9, -1000, 24, -1000,
	197, 977, -1000, 213, -1000, 926, 176, 67, 67, 67,
	67, -1000, 128, 210, -1000, 4, -1000, 887, 170, 836,
	67, -1000, 552, 785, 67, -1000, -1000, -1000, -1000, 192,
	67, 113, -1000, -1000, 97, -1000, 97, -1000, -1000, -1000,
	125, 242, -20, 163, 121, -1000, 481, -1000, -1000, -1000,
	0, -1000, -1000, -1000, 734, 6, -1000, 683, -1000, -1000,
	-9, 210, 158, 1094, -9, -1000, -1000, -1000, -1000, 174,
	1094, -1000, -1000, -1000, 146, -1000, -1000, 21, -30, -1000,
	1094, 97, -1000, 89, 1094, -1000, -1000, 200, -1000, 1016,
	-1000, -1000, 67, 1016, 644, -1000, 1094, -1000, 67, 1016,
	1016, 240, -1000, 1094, -1000, 38, 67, -1000, -1000, 67,
	-1000, 53, 37, -1000, 67, 1016, 1016, 1016, 187, 130,
	-22, -9, -1000, -1000, 1016, -1000, 1094, 97, 1016, 79,
	117, -9, -1000, -9, -1000, -9, -9,
}
var RubyPgo = []int{

	0, 299, 345, 344, 20, 338, 333, 413, 236, 332,
	331, 330, 0, 87, 355, 324, 323, 6, 320, 2,
	38, 319, 3, 4, 317, 316, 311, 310, 305, 302,
	298, 292, 291, 290, 287, 217, 286, 7, 14, 283,
	279, 9, 275, 274, 18, 271, 270, 267, 253, 5,
	1, 198, 96,
}
var RubyR1 = []int{

	0, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 52, 52, 35, 35, 35, 35, 35, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 17,
	17, 17, 17, 17, 17, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 37, 37, 46, 46, 44, 44, 44,
	44, 44, 49, 49, 49, 49, 23, 23, 48, 48,
	48, 15, 15, 41, 41, 19, 19, 50, 50, 50,
	18, 18, 21, 22, 22, 51, 51, 10, 10, 10,
	10, 10, 10, 8, 8, 20, 20, 13, 13, 24,
	24, 25, 26, 27, 28, 29, 29, 29, 29, 30,
	31, 32, 33, 34, 2, 5, 6, 6, 3, 3,
	42, 42, 42, 42, 47, 47, 47, 4, 4, 4,
	4, 38, 45, 45, 45, 9, 9, 9, 9, 9,
	9, 9, 9, 39, 39, 39, 39, 36, 36, 36,
	7, 11, 43, 43, 43, 43, 16,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	4, 4, 2, 3, 3, 4, 4, 3, 2, 3,
	3, 4, 6, 3, 1, 1, 3, 0, 1, 1,
	1, 3, 1, 1, 3, 3, 1, 1, 1, 3,
	3, 7, 7, 1, 3, 1, 3, 0, 1, 3,
	4, 6, 4, 1, 4, 1, 4, 3, 3, 3,
	3, 3, 3, 2, 2, 2, 2, 3, 3, 3,
	3, 2, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 1, 1, 3, 3, 5, 5,
	0, 4, 7, 8, 3, 7, 8, 3, 4, 4,
	3, 3, 0, 1, 3, 4, 5, 3, 3, 3,
	3, 3, 4, 4, 3, 3, 2, 0, 2, 2,
	3, 4, 0, 3, 4, 6, 1,
}
var RubyChk = []int{

	-1000, -40, 42, 43, 57, -1, 42, 43, 57, -12,
	-14, -15, -18, -21, -10, -25, -26, -27, -28, -9,
	-7, -11, -17, -16, 5, 6, 7, -20, -13, -8,
	-2, -5, -6, -3, -22, 12, 18, 19, -24, 36,
	37, 38, 39, 14, 17, 8, 23, -29, -30, -31,
	-32, -33, -34, 30, 55, 54, 31, 32, 50, 52,
	-51, 43, 42, 57, 14, 45, 4, 38, 39, 40,
	47, 48, 46, 17, 45, 4, 38, 14, 17, 17,
	45, 8, -37, -49, 50, 35, 10, -48, -12, -4,
	-23, 6, 7, -20, -13, -8, 11, 52, -14, -7,
	8, 35, 35, 35, 35, 6, 4, -22, 7, -22,
	35, 10, -1, -1, -1, -1, -1, -12, 6, 7,
	-36, -35, 6, 7, 55, 6, 7, -46, -44, -12,
	-17, -14, -52, 44, -1, 6, -12, -12, -14, -1,
	-12, -12, -12, -12, -1, 6, -12, -12, -14, -1,
	-1, -1, 6, -44, -12, 10, -12, -23, -12, 6,
	10, -35, -38, 46, -38, -35, -44, -1, -1, -1,
	-1, -41, -50, 8, -19, 6, -41, -35, 33, -35,
	-1, 6, -35, -35, -1, 43, 9, 42, 43, -12,
	-1, -43, 6, 7, 10, 51, 10, 51, 42, -42,
	-47, -12, 6, 44, -49, -37, 8, 9, -12, -4,
	51, -23, -4, 13, -35, -45, 6, -35, 53, 9,
	-52, 10, -50, 35, -52, 13, -22, 13, 13, -39,
	16, 15, 13, 13, 24, -12, -12, -52, -52, 10,
	4, 44, 7, -44, 35, 13, 46, 10, 53, -35,
	-19, 9, -1, -35, -35, 13, 16, 15, -1, -35,
	-35, 7, 53, 10, 53, -52, -1, -12, 9, -1,
	6, -52, -52, 13, -1, -35, -35, -35, 4, -1,
	6, -52, 13, 13, -35, 6, 4, 44, -35, -1,
	-12, -52, 10, -52, 10, -52, -52,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
	46, 47, 48, 49, 18, 19, -2, 21, 22, 23,
	24, 25, 26, 27, 28, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 157, 13, 29, 30, 31,
	32, 33, 34, 166, 0, 0, 124, 125, 67, 11,
	0, 5, 6, 7, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, -2, 52, 58, 67, 0, 0, 64, 72, 73,
	78, 19, -2, 21, 22, 23, 13, -2, 77, 0,
	67, 0, 0, 0, 0, 87, 87, 13, -2, 13,
	0, 0, 111, 112, 113, 114, 13, 13, 19, -2,
	0, 162, 105, 106, 0, 103, 104, 0, 0, 65,
	69, 70, 130, 0, 147, 53, 59, 115, 117, 119,
	120, 121, 122, 123, 149, 54, 60, 116, 118, 148,
	150, 151, 57, 0, 68, 0, 65, 97, 76, 109,
	0, 0, 13, 142, 13, 0, 0, 98, 99, 100,
	101, 11, 83, 87, 88, 85, 11, 0, 0, 0,
	102, 110, 0, 0, 158, 159, 160, 14, 15, 16,
	17, 0, 107, 108, 0, 126, 0, 127, 12, 11,
	11, 0, 19, 0, 55, 56, -2, 50, 74, 75,
	61, 79, 80, 137, 0, 0, 143, 0, 140, 51,
	13, 0, 0, 0, 13, 90, 13, 92, 145, 0,
	0, 13, 152, 161, 13, 66, 71, 0, 0, 11,
	0, 0, -2, 0, 0, 138, 141, 0, 139, 11,
	89, 84, 86, 11, 0, 146, 0, 13, 13, 156,
	163, 13, 128, 0, 129, 0, 11, 134, 63, 62,
	144, 0, 0, 91, 13, 154, 155, 164, 0, 0,
	0, 131, 81, 82, 153, 13, 0, 0, 165, 11,
	11, 132, 11, 135, 11, 133, 136,
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
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 60:
		//line parser.y:279
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 61:
		//line parser.y:289
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 62:
		//line parser.y:299
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 63:
		//line parser.y:308
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 64:
		//line parser.y:310
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 65:
		//line parser.y:313
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 66:
		//line parser.y:315
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:317
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 68:
		//line parser.y:319
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 69:
		//line parser.y:321
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 70:
		//line parser.y:323
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 71:
		//line parser.y:325
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 77:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 78:
		//line parser.y:340
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 79:
		//line parser.y:342
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 80:
		//line parser.y:344
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 81:
		//line parser.y:350
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 82:
		//line parser.y:358
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 83:
		//line parser.y:367
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 84:
		//line parser.y:369
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 85:
		//line parser.y:372
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 86:
		//line parser.y:374
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 87:
		//line parser.y:376
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 88:
		//line parser.y:378
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 89:
		//line parser.y:382
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 90:
		//line parser.y:387
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 91:
		//line parser.y:394
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 92:
		//line parser.y:404
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 93:
		//line parser.y:413
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 94:
		//line parser.y:419
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 95:
		//line parser.y:427
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 96:
		//line parser.y:431
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 97:
		//line parser.y:436
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 98:
		//line parser.y:443
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 99:
		//line parser.y:450
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 100:
		//line parser.y:457
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 101:
		//line parser.y:464
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 102:
		//line parser.y:471
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 103:
		//line parser.y:479
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 104:
		//line parser.y:481
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 105:
		//line parser.y:484
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 106:
		//line parser.y:486
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 107:
		//line parser.y:489
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 108:
		//line parser.y:491
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 109:
		//line parser.y:494
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 110:
		//line parser.y:496
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 111:
		//line parser.y:498
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 112:
		//line parser.y:499
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 113:
		//line parser.y:500
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 114:
		//line parser.y:501
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 115:
		//line parser.y:504
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 116:
		//line parser.y:512
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 117:
		//line parser.y:520
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 118:
		//line parser.y:528
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 119:
		//line parser.y:537
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 120:
		//line parser.y:546
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 121:
		//line parser.y:555
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 122:
		//line parser.y:564
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 123:
		//line parser.y:573
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 124:
		//line parser.y:581
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 125:
		//line parser.y:582
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 126:
		//line parser.y:584
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 127:
		//line parser.y:586
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 128:
		//line parser.y:589
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 129:
		//line parser.y:597
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 130:
		//line parser.y:605
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 131:
		//line parser.y:607
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 132:
		//line parser.y:614
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 133:
		//line parser.y:621
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 134:
		//line parser.y:629
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 135:
		//line parser.y:636
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 136:
		//line parser.y:643
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 137:
		//line parser.y:651
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 138:
		//line parser.y:653
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 139:
		//line parser.y:660
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 140:
		//line parser.y:667
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 141:
		//line parser.y:670
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 142:
		//line parser.y:672
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 143:
		//line parser.y:674
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 144:
		//line parser.y:676
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 145:
		//line parser.y:679
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 146:
		//line parser.y:686
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 147:
		//line parser.y:694
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 148:
		//line parser.y:701
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 149:
		//line parser.y:708
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 150:
		//line parser.y:715
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 151:
		//line parser.y:722
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 152:
		//line parser.y:729
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-2].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 153:
		//line parser.y:737
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 154:
		//line parser.y:744
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
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
		}
	case 158:
		//line parser.y:766
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 159:
		//line parser.y:767
		{
		}
	case 160:
		//line parser.y:770
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 161:
		//line parser.y:773
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 162:
		//line parser.y:781
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 163:
		//line parser.y:783
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 164:
		//line parser.y:785
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 165:
		//line parser.y:794
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
	case 166:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	}
	goto Rubystack /* stack new state and value */
}
