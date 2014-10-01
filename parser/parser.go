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

//line parser.y:799

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 25,
	43, 92,
	-2, 20,
	-1, 76,
	10, 65,
	-2, 156,
	-1, 87,
	43, 92,
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
	-1, 100,
	43, 92,
	-2, 90,
	-1, 194,
	9, 65,
	10, 65,
	-2, 156,
	-1, 230,
	43, 93,
	-2, 91,
}

const RubyNprod = 165
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1221

var RubyAct = []int{

	9, 163, 22, 33, 161, 152, 119, 84, 62, 78,
	77, 186, 64, 2, 3, 75, 184, 272, 153, 113,
	114, 178, 251, 5, 229, 83, 191, 125, 4, 124,
	186, 232, 72, 212, 97, 60, 59, 96, 182, 99,
	101, 235, 166, 95, 109, 65, 66, 67, 94, 10,
	61, 64, 63, 70, 68, 69, 185, 121, 199, 122,
	104, 105, 106, 107, 108, 128, 129, 115, 132, 133,
	134, 135, 250, 137, 85, 266, 234, 143, 183, 122,
	146, 147, 122, 144, 126, 62, 73, 83, 131, 74,
	221, 63, 268, 136, 279, 139, 140, 103, 154, 255,
	184, 222, 242, 186, 244, 243, 123, 20, 277, 72,
	186, 145, 62, 177, 249, 130, 156, 157, 158, 159,
	186, 102, 138, 227, 168, 189, 123, 210, 83, 123,
	148, 271, 93, 172, 239, 210, 85, 192, 193, 196,
	184, 62, 150, 164, 270, 162, 197, 180, 181, 116,
	117, 200, 177, 198, 248, 230, 177, 100, 202, 257,
	164, 206, 169, 149, 93, 141, 177, 211, 177, 127,
	214, 177, 177, 93, 98, 223, 264, 85, 228, 58,
	93, 82, 188, 118, 93, 224, 64, 93, 93, 205,
	209, 179, 187, 160, 93, 121, 1, 122, 217, 71,
	201, 231, 111, 51, 50, 177, 49, 48, 177, 47,
	46, 26, 238, 18, 17, 16, 15, 225, 226, 65,
	66, 67, 37, 13, 12, 11, 63, 70, 68, 69,
	254, 21, 14, 19, 240, 93, 88, 31, 177, 30,
	245, 32, 177, 29, 123, 0, 0, 177, 177, 0,
	253, 0, 0, 27, 256, 88, 112, 252, 93, 0,
	0, 0, 177, 177, 177, 260, 0, 258, 88, 0,
	177, 0, 265, 275, 177, 0, 88, 88, 89, 88,
	88, 88, 88, 267, 88, 0, 0, 0, 88, 0,
	0, 88, 88, 274, 0, 0, 0, 89, 88, 0,
	0, 0, 93, 28, 276, 278, 0, 280, 0, 281,
	89, 64, 0, 0, 0, 0, 195, 0, 89, 89,
	0, 89, 89, 89, 89, 0, 89, 0, 90, 0,
	89, 0, 0, 89, 89, 0, 88, 0, 0, 88,
	89, 0, 0, 0, 65, 66, 67, 90, 151, 155,
	0, 63, 70, 68, 69, 0, 165, 88, 167, 0,
	90, 0, 88, 0, 0, 170, 171, 0, 90, 90,
	0, 90, 90, 90, 90, 0, 90, 0, 89, 64,
	90, 89, 0, 90, 90, 0, 0, 0, 0, 0,
	90, 23, 86, 87, 44, 0, 88, 91, 0, 89,
	0, 0, 0, 0, 89, 0, 88, 0, 0, 204,
	0, 207, 65, 66, 67, 0, 54, 55, 0, 63,
	70, 68, 69, 0, 0, 0, 0, 0, 90, 0,
	0, 90, 0, 0, 0, 56, 0, 92, 89, 53,
	52, 88, 0, 0, 0, 0, 0, 0, 89, 90,
	0, 0, 0, 0, 90, 0, 0, 0, 0, 0,
	23, 24, 25, 44, 0, 0, 237, 34, 259, 42,
	0, 241, 43, 35, 36, 0, 246, 0, 45, 247,
	0, 0, 0, 89, 88, 54, 55, 0, 90, 0,
	38, 39, 40, 41, 0, 0, 175, 176, 90, 0,
	0, 261, 262, 0, 56, 263, 57, 0, 53, 52,
	0, 23, 24, 25, 44, 0, 0, 269, 34, 216,
	42, 219, 218, 43, 35, 36, 89, 273, 0, 45,
	0, 0, 0, 90, 0, 0, 54, 55, 0, 0,
	0, 38, 39, 40, 41, 0, 0, 175, 176, 23,
	24, 25, 44, 0, 0, 56, 34, 57, 42, 53,
	52, 43, 35, 36, 0, 0, 0, 45, 0, 0,
	0, 0, 0, 0, 54, 55, 90, 0, 0, 38,
	39, 40, 41, 0, 0, 6, 7, 0, 0, 0,
	0, 0, 0, 56, 0, 57, 0, 53, 52, 0,
	8, 23, 24, 25, 44, 0, 0, 0, 34, 0,
	42, 0, 0, 43, 35, 36, 0, 0, 0, 45,
	0, 0, 0, 0, 0, 0, 54, 55, 0, 0,
	0, 38, 39, 40, 41, 0, 0, 175, 176, 0,
	0, 0, 0, 0, 0, 56, 0, 57, 236, 53,
	52, 23, 24, 25, 44, 0, 0, 0, 34, 233,
	42, 0, 0, 43, 35, 36, 0, 0, 0, 45,
	0, 0, 0, 0, 0, 0, 54, 55, 0, 0,
	0, 38, 39, 40, 41, 0, 0, 175, 176, 0,
	0, 0, 0, 0, 0, 56, 0, 57, 0, 53,
	52, 23, 24, 25, 44, 0, 0, 0, 34, 220,
	42, 0, 0, 43, 35, 36, 0, 0, 0, 45,
	0, 0, 0, 0, 0, 0, 54, 55, 0, 0,
	0, 38, 39, 40, 41, 0, 0, 175, 176, 0,
	0, 0, 0, 0, 0, 56, 0, 57, 0, 53,
	52, 23, 24, 25, 44, 0, 0, 0, 34, 215,
	42, 0, 0, 43, 35, 36, 0, 0, 0, 45,
	0, 0, 0, 0, 0, 0, 54, 55, 0, 0,
	0, 38, 39, 40, 41, 0, 0, 175, 176, 0,
	0, 0, 0, 0, 0, 56, 0, 57, 0, 53,
	52, 23, 24, 25, 44, 0, 0, 0, 34, 213,
	42, 0, 0, 43, 35, 36, 0, 0, 0, 45,
	0, 0, 0, 0, 0, 0, 54, 55, 0, 0,
	0, 38, 39, 40, 41, 0, 0, 175, 176, 23,
	24, 25, 44, 0, 0, 56, 34, 57, 42, 53,
	52, 43, 35, 36, 0, 0, 0, 45, 0, 0,
	0, 0, 0, 0, 54, 55, 0, 0, 0, 38,
	39, 40, 41, 0, 0, 175, 176, 0, 0, 0,
	0, 0, 0, 56, 0, 57, 208, 53, 52, 23,
	24, 25, 44, 0, 0, 0, 34, 203, 42, 0,
	0, 43, 35, 36, 0, 0, 0, 45, 0, 0,
	0, 0, 0, 0, 54, 55, 0, 0, 0, 38,
	39, 40, 41, 0, 0, 175, 176, 23, 24, 25,
	44, 0, 0, 56, 34, 57, 42, 53, 52, 43,
	35, 36, 0, 0, 0, 45, 0, 0, 0, 0,
	0, 0, 54, 55, 0, 0, 0, 38, 39, 40,
	41, 0, 0, 175, 176, 23, 24, 25, 44, 174,
	0, 56, 34, 57, 42, 53, 52, 43, 35, 36,
	0, 0, 0, 45, 0, 0, 0, 0, 0, 0,
	54, 55, 0, 0, 0, 38, 39, 40, 41, 0,
	0, 0, 173, 23, 24, 25, 44, 0, 0, 56,
	34, 57, 42, 53, 52, 43, 35, 36, 0, 0,
	0, 45, 0, 0, 0, 0, 0, 0, 54, 55,
	0, 0, 0, 38, 39, 40, 41, 23, 86, 87,
	76, 0, 81, 91, 0, 0, 0, 56, 0, 57,
	0, 53, 52, 23, 86, 87, 194, 0, 0, 91,
	0, 0, 54, 55, 0, 0, 80, 23, 86, 87,
	76, 0, 0, 91, 0, 0, 0, 0, 54, 55,
	0, 79, 0, 92, 0, 53, 52, 0, 0, 0,
	0, 0, 54, 55, 0, 0, 0, 56, 0, 92,
	0, 53, 52, 23, 86, 87, 44, 142, 0, 0,
	0, 79, 0, 92, 0, 53, 52, 23, 86, 87,
	44, 0, 0, 0, 0, 0, 0, 0, 54, 55,
	23, 110, 87, 0, 0, 0, 91, 0, 0, 0,
	0, 0, 54, 55, 23, 190, 87, 56, 0, 57,
	0, 53, 52, 0, 0, 54, 55, 120, 86, 87,
	44, 56, 0, 57, 0, 53, 52, 0, 0, 54,
	55, 23, 110, 87, 56, 0, 92, 0, 53, 52,
	186, 0, 54, 55, 0, 0, 0, 0, 56, 0,
	57, 0, 53, 52, 0, 0, 54, 55, 0, 0,
	0, 56, 0, 57, 0, 53, 52, 0, 0, 0,
	0, 0, 0, 0, 0, 56, 0, 57, 0, 53,
	52,
}
var RubyPact = []int{

	-28, 544, -1000, -1000, -1000, -6, -1000, -1000, -1000, 182,
	72, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-29, -1000, -1000, -1000, 1032, 14, 9, 3, 0, -1000,
	-1000, -1000, -1000, -1000, 168, 150, 150, 87, 998, 998,
	998, 998, 998, 1166, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 13, 143, -1000, -1000, 1152, -1000, -16, -1000,
	-1000, -1000, 998, 163, 1166, 1112, 998, 1166, 1166, 1166,
	1166, 998, 1112, 998, 998, 159, 1098, -1000, 101, 1152,
	1112, 157, 132, 47, -1000, -1000, 1062, -1000, -1000, -1000,
	-1000, -27, -27, -29, 998, 998, 998, 998, 137, 10,
	-1000, -1000, 998, 156, 71, 71, 71, 71, 71, -1000,
	-1000, 960, 922, -1000, -1000, 141, -1000, -1000, 28, 6,
	-1000, 375, -1000, -5, 1139, -17, 71, 1048, -1000, 47,
	-1000, 71, -1000, -1000, -1000, -1000, 71, 47, -1000, 71,
	71, -1000, -1000, 307, 130, 1125, 8, 47, -1000, -1000,
	386, 884, -1000, 155, -1000, 834, 71, 71, 71, 71,
	-1000, 117, 154, -1000, -1, 796, 150, 746, 71, -1000,
	506, 696, 71, -1000, -1000, -1000, -1000, 182, 71, 77,
	-1000, -1000, 170, -1000, 1166, -1000, -1000, -1000, 113, 174,
	-19, 148, 101, -1000, 1112, -1000, -1000, -1000, -1000, -3,
	47, -1000, -1000, -1000, 646, 31, -1000, 596, -1000, -11,
	154, 125, 998, -1000, -1000, -1000, -1000, 89, 998, -1000,
	-1000, -1000, 147, -1000, -1000, 62, -30, -1000, 998, 1166,
	-1000, 90, 998, -1000, -1000, 153, -1000, 922, -1000, -1000,
	71, 455, -1000, 998, -1000, 71, 922, 922, 172, -1000,
	998, -1000, 69, 71, -1000, -1000, 71, -1000, 79, -1000,
	71, 922, 922, 922, 138, 127, -26, -11, -1000, 922,
	-1000, 998, 1166, 922, 98, 84, -11, -1000, -11, -1000,
	-11, -11,
}
var RubyPgo = []int{

	0, 21, 243, 241, 7, 239, 237, 107, 303, 233,
	232, 231, 0, 253, 49, 225, 2, 224, 211, 223,
	3, 1, 222, 216, 215, 214, 213, 210, 209, 207,
	206, 204, 203, 256, 202, 10, 5, 198, 196, 193,
	192, 191, 6, 189, 183, 182, 181, 9, 4, 179,
	29,
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
	39, 39, 21, 21, 48, 48, 48, 17, 17, 19,
	20, 20, 49, 49, 10, 10, 10, 10, 10, 10,
	10, 8, 8, 18, 18, 13, 13, 22, 22, 23,
	24, 25, 26, 27, 27, 27, 27, 28, 29, 30,
	31, 32, 2, 5, 6, 6, 3, 3, 40, 40,
	40, 40, 45, 45, 45, 45, 4, 4, 4, 4,
	36, 43, 43, 43, 9, 9, 9, 9, 9, 9,
	9, 37, 37, 37, 37, 37, 34, 34, 34, 7,
	11, 41, 41, 41, 41,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
	4, 4, 2, 3, 4, 4, 3, 2, 3, 4,
	6, 3, 1, 1, 3, 0, 1, 1, 1, 3,
	1, 1, 3, 3, 1, 1, 3, 3, 3, 7,
	1, 3, 1, 3, 0, 1, 3, 4, 6, 4,
	1, 4, 1, 4, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 2, 2, 3, 3, 3, 3, 2,
	2, 2, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 1, 1, 3, 3, 5, 5, 0, 4,
	7, 8, 3, 3, 7, 8, 3, 4, 4, 3,
	3, 0, 1, 3, 4, 5, 3, 3, 3, 3,
	4, 0, 4, 3, 3, 2, 0, 2, 2, 3,
	4, 0, 3, 4, 6,
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
	-8, 11, 51, -7, 34, 34, 34, 34, 6, -20,
	7, -20, 34, 10, -1, -1, -1, -1, -1, -12,
	6, -34, -33, 6, 7, 54, 6, 7, -44, -42,
	5, -12, -16, -14, -50, 43, -1, 6, -12, -12,
	-14, -1, -12, -12, -12, -12, -1, -12, -14, -1,
	-1, 6, 9, -12, -42, 10, -12, -12, -14, 6,
	10, -33, -36, 45, -36, -33, -1, -1, -1, -1,
	-39, -48, 8, -21, 6, -33, 32, -33, -1, 6,
	-33, -33, -1, 42, 9, 41, 42, -12, -1, -41,
	6, 7, 10, 50, 10, 50, 41, -40, -45, -12,
	6, 43, -47, -35, 8, 9, 9, -12, -4, 50,
	-12, -14, -4, 13, -33, -43, 6, -33, 52, -50,
	10, -48, 34, 13, -20, 13, 13, -37, 16, 15,
	13, 13, 24, 5, -12, -50, -50, 10, 4, 43,
	7, -42, 34, 13, 45, 10, 52, -33, -21, 9,
	-1, -33, 13, 16, 15, -1, -33, -33, 7, 52,
	10, 52, -50, -1, -12, 9, -1, 6, -50, 13,
	-1, -33, -33, -33, 4, -1, 6, -50, 13, -33,
	6, 4, 43, -33, -1, -12, -50, 10, -50, 10,
	-50, -50,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
	46, 47, 48, 18, 19, -2, 21, 22, 23, 24,
	25, 26, 27, 28, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 156, 13, 29, 30, 31, 32,
	33, 34, 0, 0, 122, 123, 65, 11, 0, 5,
	6, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, -2, 52, 57, 65,
	0, 0, 62, 70, 71, 75, 19, -2, 21, 22,
	23, 13, -2, 0, 0, 0, 0, 0, 84, 13,
	-2, 13, 0, 0, 109, 110, 111, 112, 13, 13,
	19, 0, 161, 103, 104, 0, 101, 102, 0, 0,
	18, 66, 67, 68, 128, 0, 146, 53, 58, 113,
	115, 117, 118, 119, 120, 121, 148, 114, 116, 147,
	149, 56, 49, 66, 0, 0, 66, 94, 95, 107,
	0, 0, 13, 141, 13, 0, 96, 97, 98, 99,
	11, 80, 84, 85, 82, 0, 0, 0, 100, 108,
	0, 0, 157, 158, 159, 14, 15, 16, 17, 0,
	105, 106, 0, 124, 0, 125, 12, 11, 11, 0,
	19, 0, 54, 55, -2, 50, 51, 72, 73, 59,
	76, 77, 78, 136, 0, 0, 142, 0, 139, 13,
	0, 0, 0, 87, 13, 89, 144, 0, 0, 13,
	150, 160, 13, 64, 69, 0, 0, 11, 0, 0,
	-2, 0, 0, 137, 140, 0, 138, 11, 86, 81,
	83, 0, 145, 0, 13, 13, 155, 162, 13, 126,
	0, 127, 0, 11, 132, 61, 60, 143, 0, 88,
	13, 153, 154, 163, 0, 0, 0, 129, 79, 152,
	13, 0, 0, 164, 11, 11, 130, 11, 134, 11,
	131, 135,
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
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: []ast.Node{},
			}
		}
	case 50:
		//line parser.y:207
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 51:
		//line parser.y:214
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 52:
		//line parser.y:221
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
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
		//line parser.y:350
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 81:
		//line parser.y:352
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 82:
		//line parser.y:355
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 83:
		//line parser.y:357
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 84:
		//line parser.y:359
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 85:
		//line parser.y:361
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 86:
		//line parser.y:365
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 87:
		//line parser.y:370
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 88:
		//line parser.y:377
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 89:
		//line parser.y:387
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 90:
		//line parser.y:396
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 91:
		//line parser.y:402
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 92:
		//line parser.y:410
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 93:
		//line parser.y:414
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 94:
		//line parser.y:419
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 95:
		//line parser.y:426
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 96:
		//line parser.y:433
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 97:
		//line parser.y:440
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 98:
		//line parser.y:447
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 99:
		//line parser.y:454
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 100:
		//line parser.y:461
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 101:
		//line parser.y:469
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 102:
		//line parser.y:471
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 103:
		//line parser.y:474
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 104:
		//line parser.y:476
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 105:
		//line parser.y:479
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 106:
		//line parser.y:481
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 107:
		//line parser.y:484
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 108:
		//line parser.y:486
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 109:
		//line parser.y:488
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 110:
		//line parser.y:489
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 111:
		//line parser.y:490
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 112:
		//line parser.y:491
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 113:
		//line parser.y:494
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
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
		//line parser.y:527
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 118:
		//line parser.y:536
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 119:
		//line parser.y:545
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 120:
		//line parser.y:554
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 121:
		//line parser.y:563
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 122:
		//line parser.y:571
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 123:
		//line parser.y:572
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 124:
		//line parser.y:574
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 125:
		//line parser.y:576
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 126:
		//line parser.y:579
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
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
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 129:
		//line parser.y:597
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 130:
		//line parser.y:604
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 131:
		//line parser.y:611
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 132:
		//line parser.y:619
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 133:
		//line parser.y:626
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 134:
		//line parser.y:633
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 135:
		//line parser.y:640
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 136:
		//line parser.y:648
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 137:
		//line parser.y:650
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 138:
		//line parser.y:657
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 139:
		//line parser.y:664
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 140:
		//line parser.y:667
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 141:
		//line parser.y:669
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 142:
		//line parser.y:671
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 143:
		//line parser.y:673
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 144:
		//line parser.y:676
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 145:
		//line parser.y:683
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 146:
		//line parser.y:691
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 147:
		//line parser.y:698
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 148:
		//line parser.y:705
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 149:
		//line parser.y:712
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 150:
		//line parser.y:719
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-2].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 151:
		//line parser.y:726
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 152:
		//line parser.y:728
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 153:
		//line parser.y:735
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 154:
		//line parser.y:742
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 155:
		//line parser.y:749
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 156:
		//line parser.y:756
		{
		}
	case 157:
		//line parser.y:757
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 158:
		//line parser.y:758
		{
		}
	case 159:
		//line parser.y:761
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 160:
		//line parser.y:764
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 161:
		//line parser.y:772
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 162:
		//line parser.y:774
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 163:
		//line parser.y:776
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 164:
		//line parser.y:785
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
