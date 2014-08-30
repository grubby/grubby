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

//line parser.y:609

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 12,
	-1, 11,
	1, 17,
	4, 17,
	10, 17,
	13, 17,
	15, 17,
	16, 17,
	17, 17,
	39, 17,
	41, 17,
	42, 17,
	44, 17,
	45, 17,
	46, 17,
	47, 17,
	53, 17,
	57, 17,
	-2, 12,
	-1, 12,
	5, 12,
	6, 12,
	7, 12,
	8, 12,
	34, 12,
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
	-1, 85,
	5, 12,
	6, 12,
	7, 12,
	8, 12,
	-2, 45,
	-1, 87,
	10, 12,
	-2, 49,
	-1, 90,
	10, 12,
	-2, 52,
	-1, 91,
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
	-1, 93,
	1, 17,
	4, 17,
	10, 17,
	13, 17,
	15, 17,
	16, 17,
	17, 17,
	39, 17,
	41, 17,
	42, 17,
	44, 17,
	45, 17,
	46, 17,
	47, 17,
	53, 17,
	57, 17,
	-2, 12,
	-1, 94,
	5, 12,
	6, 12,
	7, 12,
	8, 12,
	34, 12,
	-2, 18,
	-1, 96,
	45, 12,
	-2, 8,
	-1, 102,
	5, 12,
	6, 12,
	7, 12,
	8, 12,
	-2, 46,
	-1, 112,
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
	-1, 113,
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
	-1, 114,
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
	-1, 115,
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
	-1, 120,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 8,
	-1, 146,
	43, 82,
	-2, 12,
	-1, 153,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 50,
	-1, 154,
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
	-1, 155,
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
	-1, 156,
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
	-1, 157,
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
	-1, 158,
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
	-1, 159,
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
	-1, 160,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 130,
	-1, 161,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 131,
	-1, 164,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 84,
	-1, 169,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 11,
	-1, 172,
	45, 123,
	-2, 65,
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
	-2, 85,
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
	-2, 86,
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
	-2, 87,
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
	-2, 88,
	-1, 190,
	4, 17,
	17, 17,
	39, 17,
	44, 17,
	45, 17,
	46, 17,
	47, 17,
	-2, 12,
	-1, 200,
	45, 124,
	-2, 66,
	-1, 217,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 55,
	-1, 233,
	43, 83,
	-2, 12,
	-1, 245,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 8,
	-1, 257,
	10, 108,
	41, 108,
	53, 108,
	-2, 12,
	-1, 258,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 8,
	-1, 269,
	45, 125,
	-2, 69,
	-1, 272,
	10, 105,
	41, 105,
	53, 105,
	-2, 12,
	-1, 278,
	41, 109,
	53, 109,
	-2, 12,
	-1, 281,
	41, 106,
	53, 106,
	-2, 12,
}

const RubyNprod = 137
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 840

var RubyAct = []int{

	135, 202, 6, 92, 109, 86, 234, 269, 52, 8,
	53, 2, 55, 56, 58, 59, 60, 61, 62, 119,
	136, 220, 221, 222, 265, 116, 263, 96, 7, 185,
	282, 227, 54, 280, 172, 7, 234, 235, 247, 57,
	66, 67, 68, 274, 226, 69, 70, 71, 72, 73,
	74, 205, 75, 7, 169, 183, 211, 236, 7, 150,
	7, 9, 97, 7, 54, 262, 63, 54, 110, 110,
	186, 204, 223, 111, 7, 4, 5, 121, 122, 123,
	124, 125, 126, 127, 128, 129, 130, 236, 131, 132,
	133, 134, 229, 230, 55, 56, 148, 137, 145, 138,
	229, 230, 7, 139, 140, 141, 142, 144, 105, 143,
	91, 101, 99, 100, 89, 104, 103, 149, 277, 76,
	151, 256, 219, 185, 112, 113, 114, 115, 233, 83,
	120, 212, 84, 214, 213, 248, 162, 240, 238, 166,
	98, 152, 163, 106, 107, 174, 110, 182, 146, 179,
	181, 271, 77, 78, 79, 101, 99, 100, 89, 165,
	82, 80, 81, 108, 195, 180, 196, 197, 102, 184,
	64, 65, 198, 118, 200, 85, 153, 154, 155, 156,
	157, 158, 159, 160, 161, 225, 207, 164, 208, 209,
	210, 55, 147, 173, 215, 175, 176, 177, 201, 218,
	87, 224, 228, 118, 117, 90, 189, 188, 199, 187,
	1, 178, 242, 192, 243, 216, 171, 32, 31, 30,
	29, 28, 27, 231, 26, 232, 25, 24, 237, 239,
	23, 35, 19, 249, 251, 244, 13, 18, 252, 17,
	254, 14, 20, 36, 16, 15, 246, 33, 22, 261,
	217, 76, 34, 21, 264, 3, 266, 267, 253, 0,
	255, 83, 0, 0, 84, 0, 259, 0, 0, 0,
	245, 273, 0, 0, 0, 276, 0, 0, 279, 268,
	0, 275, 0, 0, 77, 78, 79, 0, 0, 0,
	0, 0, 82, 80, 81, 0, 0, 257, 258, 0,
	0, 0, 0, 0, 0, 0, 0, 10, 11, 12,
	0, 0, 0, 0, 39, 191, 51, 194, 193, 270,
	40, 41, 272, 0, 0, 0, 0, 0, 0, 0,
	0, 278, 42, 43, 281, 0, 0, 44, 45, 46,
	47, 0, 168, 167, 0, 0, 0, 0, 0, 0,
	0, 0, 48, 0, 49, 0, 38, 37, 50, 10,
	11, 12, 0, 0, 0, 0, 39, 260, 51, 0,
	0, 0, 40, 41, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 42, 43, 0, 0, 0, 44,
	45, 46, 47, 0, 168, 167, 0, 0, 0, 0,
	0, 0, 0, 0, 48, 0, 49, 0, 38, 37,
	50, 10, 11, 12, 0, 0, 0, 0, 39, 250,
	51, 0, 0, 0, 40, 41, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 42, 43, 0, 0,
	0, 44, 45, 46, 47, 0, 168, 167, 0, 0,
	0, 0, 0, 0, 0, 0, 48, 0, 49, 0,
	38, 37, 50, 10, 11, 12, 0, 0, 0, 0,
	39, 206, 51, 0, 0, 0, 40, 41, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 42, 43,
	0, 0, 0, 44, 45, 46, 47, 0, 168, 167,
	0, 0, 0, 0, 0, 0, 0, 0, 48, 0,
	49, 0, 38, 37, 50, 10, 11, 12, 0, 0,
	0, 0, 39, 203, 51, 0, 0, 0, 40, 41,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	42, 43, 0, 0, 0, 44, 45, 46, 47, 0,
	168, 167, 0, 0, 0, 0, 0, 0, 0, 0,
	48, 0, 49, 0, 38, 37, 50, 10, 11, 12,
	0, 0, 0, 0, 39, 170, 51, 0, 0, 0,
	40, 41, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 42, 43, 0, 0, 0, 44, 45, 46,
	47, 0, 168, 167, 0, 0, 0, 0, 0, 0,
	0, 0, 48, 0, 49, 0, 38, 37, 50, 95,
	93, 94, 89, 0, 0, 96, 39, 0, 51, 0,
	0, 0, 40, 41, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 42, 43, 0, 0, 88, 44,
	45, 46, 47, 0, 0, 0, 0, 0, 10, 11,
	12, 0, 0, 0, 48, 39, 49, 51, 38, 37,
	50, 40, 41, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 42, 43, 0, 0, 0, 44, 45,
	46, 47, 0, 168, 167, 0, 0, 0, 0, 0,
	0, 0, 0, 48, 0, 49, 0, 38, 37, 50,
	10, 11, 12, 0, 0, 0, 96, 39, 0, 51,
	10, 11, 12, 40, 41, 0, 0, 39, 0, 51,
	0, 0, 0, 40, 41, 42, 43, 0, 0, 0,
	44, 45, 46, 47, 0, 42, 43, 0, 0, 0,
	44, 45, 46, 47, 0, 48, 0, 49, 0, 38,
	37, 50, 0, 0, 0, 48, 0, 49, 0, 38,
	37, 50, 10, 190, 12, 0, 0, 0, 0, 39,
	0, 51, 0, 0, 0, 40, 41, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 76, 42, 43, 0,
	0, 0, 44, 45, 46, 47, 83, 0, 0, 84,
	0, 0, 0, 0, 0, 0, 0, 48, 0, 49,
	0, 38, 37, 50, 0, 0, 241, 0, 0, 77,
	78, 79, 0, 0, 0, 0, 0, 82, 80, 81,
}
var RubyPact = []int{

	-46, 34, -1000, -48, -1000, -1000, 715, 62, -1000, -12,
	-1000, 62, -5, 62, 62, 62, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 11, 164, 62,
	62, 62, -1000, -1000, 62, 62, 62, 62, 62, 62,
	-1000, 62, -1000, 115, 169, 614, 106, 162, 82, 81,
	74, -1000, -1000, 137, -1000, -1000, 157, 62, 62, 715,
	715, 715, 715, 198, -1000, 715, 62, 62, 62, 62,
	62, 62, 62, 62, 62, 62, -1000, 62, 62, 62,
	62, -12, -1000, 62, -5, -1000, 62, -1000, 62, -1000,
	-1000, -1000, 62, 62, 62, 62, -1000, -1000, 62, 66,
	141, 55, -12, -12, -12, -12, 62, -1000, -1000, 18,
	-12, 715, 715, 715, 715, 715, 715, 715, 715, 715,
	150, 132, 715, 198, 129, 115, 562, -11, 715, 150,
	715, 715, 715, 150, -1000, 62, 62, 12, -1000, 19,
	-1000, 767, 302, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -1000, 62, -12, 62, 62, -1000, -1000, -12,
	-1000, 62, 168, -12, -1000, -12, -12, -12, -1000, -1000,
	510, 30, -1000, 8, 458, 62, -1000, 62, 62, -12,
	13, -1000, 118, 62, -1000, 705, 113, 16, 3, -14,
	-1000, 62, 60, -1000, -1000, 121, -1000, 31, 128, 127,
	792, 62, -1000, 62, -1000, 715, 653, -12, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -3, -1000, -1000, 125, -1000,
	-1000, 653, 406, 62, -1000, -1000, -1000, 18, -1000, 18,
	-1000, 88, 715, 715, 653, -12, 354, -1000, 62, 52,
	-1000, -1000, -27, 18, -29, 18, 62, -12, -12, 653,
	-1000, 1, -1000, -1000, 715, -1000, 145, 715, 653, -1000,
	-12, 0, -12, 247, 62, 85, 715, 62, 23, 715,
	-1000, 20, -1000,
}
var RubyPgo = []int{

	0, 259, 54, 255, 253, 252, 3, 248, 247, 245,
	244, 243, 242, 241, 239, 237, 236, 232, 4, 231,
	230, 227, 226, 224, 222, 221, 220, 219, 218, 217,
	20, 5, 216, 213, 211, 210, 209, 25, 208, 207,
	205, 200, 192, 0, 19, 1, 185,
}
var RubyR1 = []int{

	0, 35, 35, 35, 35, 3, 3, 3, 30, 30,
	30, 30, 43, 43, 44, 44, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 10, 10, 10, 10, 10, 10, 10,
	10, 31, 31, 41, 41, 41, 41, 40, 40, 40,
	40, 40, 40, 40, 40, 37, 37, 37, 37, 37,
	37, 45, 45, 45, 14, 34, 34, 15, 15, 17,
	18, 18, 42, 42, 12, 12, 12, 12, 12, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 4,
	7, 8, 5, 5, 36, 36, 36, 36, 39, 39,
	39, 1, 1, 9, 9, 16, 16, 13, 13, 19,
	6, 6, 32, 38, 38, 38, 46, 46, 11, 11,
	11, 11, 33, 33, 33, 33, 33,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 1, 3, 0, 2,
	2, 2, 0, 2, 0, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 3, 3, 3, 5, 5, 3,
	5, 5, 1, 1, 1, 5, 5, 1, 1, 1,
	5, 5, 5, 5, 5, 0, 1, 1, 5, 5,
	5, 0, 2, 2, 9, 0, 1, 6, 8, 6,
	3, 6, 1, 4, 5, 5, 5, 5, 5, 3,
	3, 3, 3, 5, 5, 5, 5, 5, 5, 1,
	1, 5, 9, 9, 0, 6, 11, 12, 4, 9,
	10, 0, 1, 2, 2, 2, 2, 3, 3, 1,
	3, 7, 3, 0, 1, 5, 1, 2, 5, 6,
	5, 5, 0, 5, 3, 4, 2,
}
var RubyChk = []int{

	-1000, -35, 57, -3, 41, 42, -43, 40, 57, -2,
	5, 6, 7, -16, -13, -9, -10, -14, -15, -17,
	-12, -4, -7, -20, -21, -22, -23, -24, -25, -26,
	-27, -28, -29, -8, -5, -19, -11, 55, 54, 12,
	18, 19, 30, 31, 35, 36, 37, 38, 50, 52,
	56, 14, -43, -43, 44, -43, -43, 44, -43, -43,
	-43, 6, 7, 55, 6, 7, -43, -43, -43, -43,
	-43, -43, -43, -43, -43, -43, 4, 37, 38, 39,
	46, 47, 45, 14, 17, 6, -31, -41, 34, 8,
	-40, -2, -6, 6, 7, 5, 11, -31, 34, 6,
	7, 5, 6, 34, 34, 34, 6, 7, 6, -18,
	-43, -18, -2, -2, -2, -2, -37, 6, 5, -44,
	-2, -43, -43, -43, -43, -43, -43, -43, -43, -43,
	-43, -43, -43, -43, -43, -43, -30, -43, -43, -43,
	-43, -43, -43, -43, 41, 32, 7, -42, 41, -43,
	41, -43, -30, -2, -2, -2, -2, -2, -2, -2,
	-2, -2, -31, 10, -2, -37, 10, 41, 40, -2,
	13, -32, 45, -2, -31, -2, -2, -2, -34, -31,
	-30, -18, -43, 43, -30, 10, 51, -36, -39, -2,
	6, 13, -33, 16, 15, -43, -43, -43, -43, -38,
	6, -37, -45, 13, 41, 43, 13, -43, -43, -43,
	-43, 43, 13, 16, 15, -43, -30, -2, -6, 9,
	5, 6, 7, 56, -6, -46, 41, 45, -43, 40,
	41, -30, -30, 7, 5, 6, 56, -44, 10, -44,
	10, 34, -43, -43, -30, -2, -30, 41, 10, -45,
	13, -43, -43, -44, -43, -44, 33, -2, -2, -30,
	13, -43, 13, 53, -43, 53, -43, -43, -30, 6,
	-2, 6, -2, -43, 43, 34, -43, 33, -2, -43,
	10, -2, 10,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 6, 0, 12, 4, 12,
	16, -2, -2, -2, -2, -2, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, 42, 0, 0, 12,
	12, 12, 99, 100, 12, 12, 12, 12, 12, 12,
	119, 12, 13, 7, 0, 0, 0, 0, 0, 0,
	0, 115, 116, 0, 113, 114, 0, 12, 12, 0,
	0, 0, 0, 65, 14, 0, 12, 12, 12, 12,
	12, 12, 12, 12, 12, -2, 43, -2, 12, 12,
	-2, -2, 54, -2, -2, 16, -2, 44, 12, 57,
	58, 59, -2, 12, 12, 12, 117, 118, 12, 0,
	0, 0, -2, -2, -2, -2, 12, 66, 67, 12,
	-2, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 75, 8, 12, -2, 0, 8, 0,
	15, 104, 0, -2, -2, -2, -2, -2, -2, -2,
	-2, -2, 47, 12, -2, 12, 12, 9, 10, -2,
	120, 12, -2, -2, 48, -2, -2, -2, 71, 76,
	0, 0, 80, 0, 0, 12, 101, 12, 12, 12,
	-2, 128, 0, 12, 8, 0, 0, 0, 0, 0,
	-2, 12, 8, 77, 8, 0, 79, 0, 14, 14,
	0, 12, 129, 12, 8, 0, 136, -2, 56, 51,
	60, 61, 62, 63, 64, 8, 126, 122, 0, 72,
	73, 71, 0, -2, 68, 69, 70, 12, 14, 12,
	14, 0, 0, 0, 134, -2, 0, 127, 12, 0,
	78, 81, 0, 12, 0, 12, 12, -2, -2, 135,
	121, 0, 74, 102, 0, 103, 0, 0, 133, -2,
	12, 0, -2, 0, 12, 0, 0, 12, -2, 0,
	110, -2, 107,
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
		//line parser.y:149
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:151
		{
			Statements = []ast.Node{}
		}
	case 3:
		//line parser.y:153
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 4:
		//line parser.y:159
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 5:
		//line parser.y:165
		{
			RubyVAL.genericValue = nil
		}
	case 6:
		//line parser.y:166
		{
			RubyVAL.genericValue = nil
		}
	case 7:
		//line parser.y:167
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 8:
		//line parser.y:170
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 9:
		//line parser.y:172
		{ /* do nothing */
		}
	case 10:
		//line parser.y:174
		{ /* do nothing */
		}
	case 11:
		//line parser.y:176
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
		//line parser.y:188
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 44:
		//line parser.y:195
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 45:
		//line parser.y:202
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 46:
		//line parser.y:209
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 47:
		//line parser.y:216
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 48:
		//line parser.y:224
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 49:
		//line parser.y:232
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 50:
		//line parser.y:239
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 51:
		//line parser.y:248
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 52:
		//line parser.y:250
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 53:
		//line parser.y:253
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 54:
		//line parser.y:255
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 55:
		//line parser.y:257
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 56:
		//line parser.y:259
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 57:
		//line parser.y:262
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 58:
		//line parser.y:264
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 59:
		//line parser.y:266
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 60:
		//line parser.y:268
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 61:
		//line parser.y:270
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 62:
		//line parser.y:272
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 63:
		//line parser.y:274
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 64:
		//line parser.y:276
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 65:
		//line parser.y:278
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 66:
		//line parser.y:280
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:282
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:284
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 69:
		//line parser.y:286
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 70:
		//line parser.y:288
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:298
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 75:
		//line parser.y:306
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 76:
		//line parser.y:307
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 77:
		//line parser.y:310
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 78:
		//line parser.y:317
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 79:
		//line parser.y:327
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 80:
		//line parser.y:336
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 81:
		//line parser.y:342
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 82:
		//line parser.y:350
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 83:
		//line parser.y:354
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 84:
		//line parser.y:359
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 85:
		//line parser.y:366
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 86:
		//line parser.y:373
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 87:
		//line parser.y:380
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 88:
		//line parser.y:387
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 89:
		//line parser.y:394
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 90:
		//line parser.y:395
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 91:
		//line parser.y:396
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 92:
		//line parser.y:397
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 93:
		//line parser.y:400
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 94:
		//line parser.y:409
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 95:
		//line parser.y:418
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 96:
		//line parser.y:427
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 97:
		//line parser.y:436
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 98:
		//line parser.y:445
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 99:
		//line parser.y:453
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 100:
		//line parser.y:454
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 101:
		//line parser.y:456
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 102:
		//line parser.y:460
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 103:
		//line parser.y:468
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 104:
		//line parser.y:476
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 105:
		//line parser.y:478
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 106:
		//line parser.y:480
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 107:
		//line parser.y:482
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 108:
		//line parser.y:485
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 109:
		//line parser.y:492
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 110:
		//line parser.y:499
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 111:
		//line parser.y:507
		{
			RubyVAL.genericValue = nil
		}
	case 112:
		//line parser.y:508
		{
			RubyVAL.genericValue = nil
		}
	case 113:
		//line parser.y:511
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 114:
		//line parser.y:513
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 115:
		//line parser.y:516
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 116:
		//line parser.y:518
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 117:
		//line parser.y:521
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 118:
		//line parser.y:523
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 119:
		//line parser.y:526
		{
			RubyVAL.genericValue = ast.FileNameConstReference{}
		}
	case 120:
		//line parser.y:529
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 121:
		//line parser.y:531
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 122:
		//line parser.y:539
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 123:
		//line parser.y:541
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 124:
		//line parser.y:543
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 125:
		//line parser.y:545
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 128:
		//line parser.y:550
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 129:
		//line parser.y:557
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 130:
		//line parser.y:565
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 131:
		//line parser.y:572
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 132:
		//line parser.y:579
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 133:
		//line parser.y:581
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 134:
		//line parser.y:588
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 135:
		//line parser.y:595
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 136:
		//line parser.y:602
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	}
	goto Rubystack /* stack new state and value */
}
