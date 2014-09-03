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

//line parser.y:665

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 12,
	-1, 11,
	11, 12,
	12, 12,
	14, 12,
	18, 12,
	19, 12,
	23, 12,
	30, 12,
	31, 12,
	34, 12,
	35, 12,
	36, 12,
	37, 12,
	38, 12,
	50, 12,
	52, 12,
	54, 12,
	55, 12,
	56, 12,
	-2, 17,
	-1, 12,
	5, 12,
	6, 12,
	7, 12,
	8, 12,
	34, 12,
	43, 85,
	-2, 18,
	-1, 13,
	34, 12,
	-2, 19,
	-1, 14,
	34, 12,
	-2, 20,
	-1, 15,
	34, 12,
	-2, 21,
	-1, 64,
	10, 12,
	-2, 55,
	-1, 89,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 141,
	-1, 103,
	5, 12,
	6, 12,
	7, 12,
	8, 12,
	-2, 48,
	-1, 104,
	10, 12,
	-2, 52,
	-1, 106,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 56,
	-1, 108,
	45, 12,
	-2, 8,
	-1, 113,
	5, 12,
	6, 12,
	7, 12,
	8, 12,
	-2, 49,
	-1, 121,
	43, 85,
	-2, 83,
	-1, 123,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 92,
	-1, 124,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 93,
	-1, 125,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 94,
	-1, 126,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 95,
	-1, 131,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 8,
	-1, 132,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 144,
	-1, 169,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 11,
	-1, 171,
	43, 86,
	-2, 84,
	-1, 172,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 53,
	-1, 173,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 96,
	-1, 174,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 97,
	-1, 175,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 98,
	-1, 176,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 99,
	-1, 177,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 100,
	-1, 178,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 101,
	-1, 179,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 133,
	-1, 180,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 134,
	-1, 183,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 87,
	-1, 186,
	45, 126,
	-2, 68,
	-1, 189,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 88,
	-1, 191,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 89,
	-1, 192,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 90,
	-1, 193,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 91,
	-1, 204,
	4, 17,
	17, 17,
	39, 17,
	44, 17,
	45, 17,
	46, 17,
	47, 17,
	-2, 12,
	-1, 214,
	45, 127,
	-2, 69,
	-1, 237,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 58,
	-1, 260,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 8,
	-1, 273,
	10, 111,
	41, 111,
	53, 111,
	-2, 12,
	-1, 274,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 8,
	-1, 285,
	10, 108,
	41, 108,
	53, 108,
	-2, 12,
	-1, 288,
	45, 128,
	-2, 72,
	-1, 301,
	41, 112,
	53, 112,
	-2, 12,
	-1, 303,
	41, 109,
	53, 109,
	-2, 12,
}

const RubyNprod = 152
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 933

var RubyAct = []int{

	136, 222, 6, 39, 8, 2, 130, 107, 58, 61,
	59, 283, 62, 68, 70, 71, 72, 249, 288, 127,
	217, 218, 219, 249, 250, 281, 108, 73, 74, 199,
	241, 67, 65, 66, 63, 7, 186, 7, 7, 60,
	116, 69, 295, 78, 79, 80, 138, 305, 81, 82,
	83, 84, 85, 86, 93, 87, 91, 304, 67, 65,
	66, 63, 7, 164, 109, 110, 7, 115, 251, 230,
	200, 220, 243, 244, 251, 236, 75, 7, 111, 280,
	279, 60, 264, 120, 122, 240, 198, 7, 137, 134,
	92, 60, 164, 7, 114, 139, 140, 141, 142, 143,
	144, 145, 146, 147, 148, 149, 150, 243, 244, 152,
	7, 4, 5, 155, 156, 157, 158, 159, 197, 209,
	160, 161, 231, 162, 233, 232, 265, 196, 163, 153,
	255, 165, 216, 199, 67, 65, 66, 63, 67, 65,
	66, 63, 253, 276, 182, 154, 117, 118, 76, 77,
	121, 129, 214, 299, 187, 188, 171, 291, 181, 129,
	128, 169, 10, 11, 12, 55, 190, 112, 9, 42,
	195, 54, 119, 113, 103, 43, 44, 210, 292, 239,
	56, 57, 104, 211, 64, 202, 212, 45, 46, 213,
	170, 201, 47, 48, 49, 50, 221, 151, 224, 1,
	226, 194, 227, 228, 229, 62, 215, 51, 234, 52,
	206, 41, 40, 53, 185, 88, 242, 89, 32, 238,
	166, 31, 30, 29, 106, 28, 27, 26, 247, 25,
	24, 257, 23, 258, 252, 254, 35, 262, 19, 13,
	18, 297, 17, 123, 124, 125, 126, 266, 267, 131,
	132, 101, 14, 268, 102, 270, 38, 272, 20, 36,
	269, 16, 271, 15, 37, 33, 278, 22, 34, 21,
	282, 3, 284, 0, 95, 96, 97, 287, 0, 0,
	0, 0, 100, 98, 99, 223, 0, 225, 0, 0,
	0, 294, 0, 296, 0, 0, 298, 235, 300, 0,
	0, 172, 173, 174, 175, 176, 177, 178, 179, 180,
	0, 245, 183, 0, 0, 0, 0, 189, 0, 191,
	192, 193, 259, 0, 0, 261, 0, 203, 263, 10,
	11, 12, 55, 0, 0, 0, 42, 205, 54, 208,
	207, 0, 43, 44, 0, 0, 0, 56, 0, 275,
	0, 0, 0, 0, 45, 46, 0, 0, 0, 47,
	48, 49, 50, 286, 168, 167, 0, 0, 0, 289,
	0, 0, 0, 237, 51, 0, 52, 0, 41, 40,
	53, 0, 0, 0, 0, 0, 0, 0, 302, 0,
	0, 10, 11, 12, 55, 0, 260, 0, 42, 293,
	54, 0, 0, 0, 43, 44, 0, 0, 0, 56,
	0, 0, 0, 0, 0, 0, 45, 46, 0, 273,
	274, 47, 48, 49, 50, 0, 168, 167, 0, 0,
	0, 0, 0, 0, 285, 0, 51, 0, 52, 0,
	41, 40, 53, 0, 290, 10, 11, 12, 55, 0,
	0, 0, 42, 277, 54, 0, 0, 0, 43, 44,
	301, 0, 303, 56, 0, 0, 0, 0, 0, 0,
	45, 46, 0, 0, 0, 47, 48, 49, 50, 0,
	168, 167, 0, 0, 0, 0, 0, 0, 0, 0,
	51, 0, 52, 0, 41, 40, 53, 10, 11, 12,
	55, 0, 0, 0, 42, 248, 54, 0, 0, 0,
	43, 44, 0, 0, 0, 56, 0, 0, 0, 0,
	0, 0, 45, 46, 0, 0, 0, 47, 48, 49,
	50, 0, 168, 167, 0, 0, 0, 0, 0, 0,
	0, 0, 51, 0, 52, 0, 41, 40, 53, 10,
	11, 12, 55, 0, 0, 0, 42, 246, 54, 0,
	0, 0, 43, 44, 0, 0, 0, 56, 0, 0,
	0, 0, 0, 0, 45, 46, 0, 0, 0, 47,
	48, 49, 50, 0, 168, 167, 0, 0, 0, 0,
	0, 0, 0, 0, 51, 0, 52, 0, 41, 40,
	53, 10, 11, 12, 55, 0, 0, 0, 42, 184,
	54, 0, 0, 0, 43, 44, 0, 0, 0, 56,
	0, 0, 0, 0, 0, 0, 45, 46, 0, 0,
	0, 47, 48, 49, 50, 0, 168, 167, 0, 0,
	10, 11, 12, 55, 135, 0, 51, 42, 52, 54,
	41, 40, 53, 43, 44, 0, 0, 0, 56, 0,
	0, 0, 0, 0, 0, 45, 46, 0, 0, 0,
	47, 48, 49, 50, 0, 7, 0, 133, 0, 10,
	11, 12, 55, 0, 0, 51, 42, 52, 54, 41,
	40, 53, 43, 44, 0, 0, 0, 56, 0, 0,
	0, 0, 0, 0, 45, 46, 0, 0, 0, 47,
	48, 49, 50, 0, 168, 167, 0, 0, 0, 0,
	0, 0, 0, 0, 51, 0, 52, 0, 41, 40,
	53, 10, 11, 12, 55, 0, 0, 108, 42, 0,
	54, 0, 0, 0, 43, 44, 0, 0, 0, 56,
	0, 0, 0, 0, 0, 0, 45, 46, 0, 0,
	105, 47, 48, 49, 50, 0, 0, 0, 0, 0,
	10, 11, 12, 55, 0, 0, 51, 42, 52, 54,
	41, 40, 53, 43, 44, 0, 0, 0, 56, 0,
	0, 0, 0, 0, 0, 45, 46, 0, 0, 0,
	47, 48, 49, 50, 0, 7, 0, 90, 0, 0,
	0, 0, 0, 0, 0, 51, 0, 52, 0, 41,
	40, 53, 10, 11, 12, 55, 0, 0, 108, 42,
	0, 54, 0, 0, 0, 43, 44, 0, 0, 0,
	56, 0, 0, 0, 0, 0, 0, 45, 46, 0,
	0, 0, 47, 48, 49, 50, 0, 0, 0, 0,
	0, 10, 204, 12, 55, 0, 0, 51, 42, 52,
	54, 41, 40, 53, 43, 44, 0, 0, 0, 56,
	0, 0, 0, 0, 0, 256, 45, 46, 0, 94,
	0, 47, 48, 49, 50, 101, 0, 0, 102, 101,
	0, 0, 102, 0, 0, 0, 51, 0, 52, 0,
	41, 40, 53, 0, 0, 0, 0, 0, 95, 96,
	97, 0, 95, 96, 97, 0, 100, 98, 99, 0,
	100, 98, 99,
}
var RubyPact = []int{

	-52, 70, -1000, -53, -1000, -1000, 157, -2, -1000, -5,
	-1000, 53, -3, -2, -2, -2, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	21, 142, -2, -2, -2, -1000, -1000, -2, -2, -2,
	-2, -2, -2, -1000, -2, 765, 49, 11, -1000, 885,
	168, -1000, 726, -2, -2, -1000, -1000, -1000, 133, 167,
	60, 33, 6, -1000, -1000, 140, -1000, -1000, 166, 143,
	143, 157, 157, 157, 157, 154, -1000, 157, 635, -5,
	-1000, -1000, -1000, 3, -2, -2, -2, -2, -2, -2,
	-2, -2, -2, -2, -2, -2, -5, -1000, -2, 154,
	135, -1000, -2, -2, -2, -2, -2, -1000, -1000, -2,
	-2, -1000, -2, -5, -5, -5, -5, -2, -1000, -1000,
	22, -5, -5, -1000, -1000, -1000, 885, 674, 149, 157,
	157, 157, 157, 157, 157, 157, 157, 157, 129, 134,
	157, 596, -9, -2, -2, 157, 129, 157, 157, 157,
	129, 86, 45, 19, -1000, 856, 324, -1000, -1000, -5,
	106, -1000, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -1000, -2, -5, -1000, -2, 146, 123, 15, -5,
	-1000, -5, -5, -5, -1000, -1000, -1000, -2, -1000, -2,
	-1000, -2, -2, -5, 26, -1000, 109, -2, -1000, -1000,
	51, 817, 44, -15, -1000, -2, -1000, -1000, -1000, -1000,
	-1000, -1000, 32, 544, 143, 492, 18, 132, 120, 881,
	-2, -1000, -2, -1000, 157, 674, -2, -5, -1000, 41,
	-1000, -1000, 116, -1000, -1000, 674, -1000, -2, -1000, -1000,
	-1000, -1000, 22, -1000, 22, -1000, -2, 157, 157, 674,
	-5, 674, 136, 440, -1000, -2, 67, 38, -28, 22,
	-42, 22, 157, -5, -5, 674, -2, -1000, 12, -1000,
	-1000, -1000, 157, -1000, 151, -5, 674, 174, -1000, 386,
	-5, -1, -2, -1000, 237, -2, 147, -2, 157, -1000,
	157, 47, 674, 37, -1000, -1000,
}
var RubyPgo = []int{

	0, 273, 161, 271, 269, 268, 7, 267, 265, 264,
	263, 261, 259, 258, 256, 252, 242, 240, 239, 238,
	3, 236, 232, 230, 229, 227, 226, 225, 223, 222,
	221, 218, 88, 215, 9, 214, 210, 201, 199, 191,
	190, 19, 189, 185, 184, 182, 181, 0, 6, 1,
	179,
}
var RubyR1 = []int{

	0, 38, 38, 38, 38, 3, 3, 3, 32, 32,
	32, 32, 47, 47, 48, 48, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 11, 11, 11, 11,
	11, 11, 11, 11, 34, 34, 45, 45, 45, 45,
	44, 44, 44, 44, 44, 44, 44, 44, 41, 41,
	41, 41, 41, 41, 49, 49, 49, 16, 37, 37,
	17, 17, 19, 20, 20, 46, 46, 13, 13, 13,
	13, 13, 22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 4, 7, 8, 5, 5, 39, 39, 39,
	39, 43, 43, 43, 1, 1, 10, 10, 18, 18,
	15, 15, 21, 6, 6, 35, 42, 42, 42, 50,
	50, 12, 12, 12, 12, 36, 36, 36, 36, 36,
	33, 33, 33, 33, 33, 33, 33, 9, 14, 40,
	40, 40,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 1, 3, 0, 2,
	2, 2, 0, 2, 0, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 3, 3, 3,
	5, 5, 3, 5, 5, 1, 1, 1, 5, 5,
	1, 1, 1, 5, 5, 5, 5, 5, 0, 1,
	1, 5, 5, 5, 0, 2, 2, 9, 0, 1,
	7, 11, 7, 1, 4, 1, 4, 5, 5, 5,
	5, 5, 3, 3, 3, 3, 5, 5, 5, 5,
	5, 5, 1, 1, 5, 9, 9, 0, 5, 10,
	11, 4, 9, 10, 0, 1, 2, 2, 2, 2,
	3, 3, 1, 3, 7, 3, 0, 1, 5, 1,
	2, 5, 6, 5, 5, 0, 5, 3, 4, 2,
	0, 1, 1, 1, 2, 2, 2, 3, 5, 0,
	4, 10,
}
var RubyChk = []int{

	-1000, -38, 57, -3, 41, 42, -47, 40, 57, -2,
	5, 6, 7, -18, -15, -10, -11, -16, -17, -19,
	-13, -4, -7, -22, -23, -24, -25, -26, -27, -28,
	-29, -30, -31, -8, -5, -21, -12, -9, -14, -20,
	55, 54, 12, 18, 19, 30, 31, 35, 36, 37,
	38, 50, 52, 56, 14, 8, 23, -46, -47, -47,
	44, -34, -47, 8, -44, 6, 7, 5, -47, 44,
	-47, -47, -47, 6, 7, 55, 6, 7, -47, -47,
	-47, -47, -47, -47, -47, -47, -47, -47, -33, -2,
	42, -47, 41, 43, 4, 37, 38, 39, 46, 47,
	45, 14, 17, 6, -45, 34, -2, -6, 11, -47,
	-47, -34, 34, 6, 34, 34, 34, 6, 7, 6,
	-20, 7, -20, -2, -2, -2, -2, -41, 6, 5,
	-48, -2, -2, 42, -47, 9, -47, -32, 43, -47,
	-47, -47, -47, -47, -47, -47, -47, -47, -47, -47,
	-47, -32, -47, -41, 10, -47, -47, -47, -47, -47,
	-47, -47, -47, -47, 41, -47, -32, 41, 40, -2,
	-40, 7, -2, -2, -2, -2, -2, -2, -2, -2,
	-2, -34, 10, -2, 13, -35, 45, -47, -47, -2,
	-34, -2, -2, -2, -37, -34, 41, 32, 41, 10,
	51, -39, -43, -2, 6, 13, -36, 16, 15, 13,
	-48, -47, -47, -42, 6, -41, 9, 5, 6, 7,
	56, -6, -49, -32, -47, -32, -47, -47, -47, -47,
	43, 13, 16, 15, -47, -32, 24, -2, -6, -50,
	41, 45, -47, 40, 41, -32, 13, -20, 13, 5,
	6, 56, -48, 10, -48, 10, 4, -47, -47, -32,
	-2, -32, -47, -32, 41, 10, -49, -47, -47, -48,
	-47, -48, -47, -2, -2, -32, 7, 13, -47, 13,
	41, 53, -47, 53, -47, -2, -32, -47, 6, -32,
	-2, 6, 4, 13, -47, 43, -47, 4, -47, 6,
	-47, -2, -32, -2, 10, 10,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 6, 0, 12, 4, 12,
	16, -2, -2, -2, -2, -2, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
	0, 0, 12, 12, 12, 102, 103, 12, 12, 12,
	12, 12, 12, 122, 12, 12, 0, 0, 13, 7,
	0, 46, 0, 12, -2, 60, 61, 62, 0, 0,
	0, 0, 0, 118, 119, 0, 116, 117, 0, 0,
	0, 0, 0, 0, 0, 68, 14, 0, 0, -2,
	142, 143, 8, 0, 12, 12, 12, 12, 12, 12,
	12, 12, 12, -2, -2, 12, -2, 57, -2, 68,
	0, 47, 12, -2, 12, 12, 12, 120, 121, 12,
	12, -2, 12, -2, -2, -2, -2, 12, 69, 70,
	12, -2, -2, 145, 146, 147, 0, 149, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 12, 12, 0, 0, 0, 0, 0,
	78, 0, 0, 0, 15, 107, 0, 9, 10, -2,
	14, -2, -2, -2, -2, -2, -2, -2, -2, -2,
	-2, 50, 12, -2, 123, 12, -2, 0, 0, -2,
	51, -2, -2, -2, 74, 79, 8, 12, 8, 12,
	104, 12, 12, 12, -2, 131, 0, 12, 8, 148,
	0, 0, 0, 0, -2, 12, 54, 63, 64, 65,
	66, 67, 8, 0, 0, 0, 0, 14, 14, 0,
	12, 132, 12, 8, 0, 139, 8, -2, 59, 8,
	129, 125, 0, 75, 76, 74, 80, 12, 82, 71,
	72, 73, 12, 14, 12, 14, 12, 0, 0, 137,
	-2, 150, 0, 0, 130, 12, 0, 0, 0, 12,
	0, 12, 0, -2, -2, 138, 12, 124, 0, 77,
	8, 105, 0, 106, 0, -2, 136, 0, -2, 0,
	12, 0, 12, 81, 0, 12, 0, 12, 0, 8,
	0, -2, 151, -2, 113, 110,
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
		//line parser.y:153
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:155
		{
			Statements = []ast.Node{}
		}
	case 3:
		//line parser.y:157
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 4:
		//line parser.y:163
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 5:
		//line parser.y:169
		{
			RubyVAL.genericValue = nil
		}
	case 6:
		//line parser.y:169
		{
			RubyVAL.genericValue = nil
		}
	case 7:
		//line parser.y:170
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 8:
		//line parser.y:173
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 9:
		//line parser.y:175
		{
		}
	case 10:
		//line parser.y:177
		{
		}
	case 11:
		//line parser.y:179
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
			}
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
		//line parser.y:191
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 47:
		//line parser.y:198
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 48:
		//line parser.y:205
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 49:
		//line parser.y:212
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 50:
		//line parser.y:219
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 51:
		//line parser.y:227
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 52:
		//line parser.y:235
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 53:
		//line parser.y:242
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 54:
		//line parser.y:251
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 55:
		//line parser.y:253
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 56:
		//line parser.y:256
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 57:
		//line parser.y:258
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 58:
		//line parser.y:260
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 59:
		//line parser.y:262
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 60:
		//line parser.y:265
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 61:
		//line parser.y:267
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 62:
		//line parser.y:269
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 63:
		//line parser.y:271
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 64:
		//line parser.y:273
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 65:
		//line parser.y:275
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 66:
		//line parser.y:277
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:279
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:281
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 69:
		//line parser.y:283
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 70:
		//line parser.y:285
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 71:
		//line parser.y:287
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 72:
		//line parser.y:289
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:291
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:301
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 78:
		//line parser.y:309
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 79:
		//line parser.y:310
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 80:
		//line parser.y:313
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 81:
		//line parser.y:320
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-8].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-4].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-8].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 82:
		//line parser.y:330
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 83:
		//line parser.y:339
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 84:
		//line parser.y:345
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 85:
		//line parser.y:353
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 86:
		//line parser.y:357
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 87:
		//line parser.y:362
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 88:
		//line parser.y:369
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 89:
		//line parser.y:376
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 90:
		//line parser.y:383
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 91:
		//line parser.y:390
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 92:
		//line parser.y:397
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 93:
		//line parser.y:398
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 94:
		//line parser.y:399
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 95:
		//line parser.y:400
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 96:
		//line parser.y:403
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 97:
		//line parser.y:412
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 98:
		//line parser.y:421
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 99:
		//line parser.y:430
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 100:
		//line parser.y:439
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 101:
		//line parser.y:448
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 102:
		//line parser.y:456
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 103:
		//line parser.y:457
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 104:
		//line parser.y:459
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 105:
		//line parser.y:463
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 106:
		//line parser.y:471
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 107:
		//line parser.y:479
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 108:
		//line parser.y:481
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 109:
		//line parser.y:488
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 110:
		//line parser.y:495
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 111:
		//line parser.y:503
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 112:
		//line parser.y:510
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 113:
		//line parser.y:517
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 114:
		//line parser.y:525
		{
			RubyVAL.genericValue = nil
		}
	case 115:
		//line parser.y:526
		{
			RubyVAL.genericValue = nil
		}
	case 116:
		//line parser.y:529
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 117:
		//line parser.y:531
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 118:
		//line parser.y:534
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 119:
		//line parser.y:536
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 120:
		//line parser.y:539
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 121:
		//line parser.y:541
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 122:
		//line parser.y:544
		{
			RubyVAL.genericValue = ast.FileNameConstReference{}
		}
	case 123:
		//line parser.y:547
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 124:
		//line parser.y:549
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 125:
		//line parser.y:557
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 126:
		//line parser.y:559
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 127:
		//line parser.y:561
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 128:
		//line parser.y:563
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 131:
		//line parser.y:568
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 132:
		//line parser.y:575
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 133:
		//line parser.y:583
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 134:
		//line parser.y:590
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 135:
		//line parser.y:597
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 136:
		//line parser.y:599
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 137:
		//line parser.y:606
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 138:
		//line parser.y:613
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 139:
		//line parser.y:620
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 140:
		//line parser.y:627
		{
		}
	case 141:
		//line parser.y:628
		{
			RubyVAL.genericSlice = []ast.Node{RubyS[Rubypt-0].genericValue}
		}
	case 142:
		//line parser.y:629
		{
		}
	case 143:
		//line parser.y:630
		{
		}
	case 144:
		//line parser.y:631
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 145:
		//line parser.y:632
		{
		}
	case 146:
		//line parser.y:633
		{
		}
	case 147:
		//line parser.y:636
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 148:
		//line parser.y:639
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 149:
		//line parser.y:647
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 150:
		//line parser.y:649
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 151:
		//line parser.y:651
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
