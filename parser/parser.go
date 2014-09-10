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

//line parser.y:666

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
	-2, 17,
	-1, 12,
	5, 12,
	6, 12,
	7, 12,
	8, 12,
	34, 12,
	43, 82,
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
	-1, 62,
	10, 12,
	-2, 54,
	-1, 87,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 137,
	-1, 101,
	5, 12,
	6, 12,
	7, 12,
	8, 12,
	-2, 47,
	-1, 102,
	10, 12,
	-2, 51,
	-1, 104,
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
	-1, 106,
	45, 12,
	-2, 8,
	-1, 111,
	5, 12,
	6, 12,
	7, 12,
	8, 12,
	-2, 48,
	-1, 119,
	43, 82,
	-2, 80,
	-1, 121,
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
	-1, 122,
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
	-2, 91,
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
	-2, 92,
	-1, 129,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 8,
	-1, 130,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 140,
	-1, 167,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 11,
	-1, 169,
	43, 83,
	-2, 81,
	-1, 170,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 52,
	-1, 171,
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
	-2, 94,
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
	-2, 95,
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
	-2, 96,
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
	-2, 97,
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
	-2, 98,
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
	-2, 129,
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
	-2, 130,
	-1, 181,
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
	-1, 184,
	45, 122,
	-2, 66,
	-1, 187,
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
	-2, 86,
	-1, 190,
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
	-2, 88,
	-1, 202,
	4, 17,
	17, 17,
	39, 17,
	44, 17,
	45, 17,
	46, 17,
	47, 17,
	-2, 12,
	-1, 212,
	45, 123,
	-2, 67,
	-1, 234,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 57,
	-1, 256,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 8,
	-1, 269,
	10, 108,
	41, 108,
	53, 108,
	-2, 12,
	-1, 270,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 8,
	-1, 281,
	10, 105,
	41, 105,
	53, 105,
	-2, 12,
	-1, 284,
	45, 124,
	-2, 70,
	-1, 298,
	41, 109,
	53, 109,
	-2, 12,
	-1, 300,
	41, 106,
	53, 106,
	-2, 12,
}

const RubyNprod = 149
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 943

var RubyAct = []int{

	134, 38, 6, 219, 59, 8, 128, 105, 56, 2,
	57, 279, 60, 66, 68, 69, 70, 71, 72, 125,
	277, 7, 238, 197, 184, 58, 294, 7, 7, 4,
	5, 67, 292, 136, 91, 276, 99, 7, 162, 100,
	302, 233, 76, 77, 78, 275, 195, 79, 80, 81,
	82, 83, 84, 85, 89, 194, 301, 135, 162, 93,
	94, 95, 107, 108, 198, 260, 73, 98, 96, 97,
	7, 109, 240, 241, 58, 65, 63, 64, 61, 118,
	120, 237, 65, 63, 64, 61, 7, 132, 240, 241,
	58, 196, 90, 137, 138, 139, 140, 141, 142, 143,
	144, 145, 146, 147, 148, 7, 114, 150, 113, 112,
	7, 153, 154, 155, 156, 157, 207, 7, 158, 159,
	227, 160, 228, 261, 230, 229, 161, 151, 251, 163,
	249, 65, 63, 64, 61, 180, 252, 152, 215, 216,
	217, 115, 116, 272, 106, 296, 99, 214, 197, 100,
	119, 179, 185, 186, 65, 63, 64, 61, 169, 188,
	110, 167, 287, 193, 149, 74, 75, 117, 9, 93,
	94, 95, 246, 284, 111, 208, 101, 98, 96, 97,
	289, 209, 246, 247, 210, 127, 212, 164, 127, 126,
	236, 55, 102, 62, 218, 200, 221, 211, 223, 168,
	224, 225, 226, 60, 213, 199, 231, 1, 192, 204,
	183, 86, 32, 31, 239, 87, 30, 235, 29, 28,
	27, 26, 104, 244, 25, 24, 23, 19, 253, 13,
	254, 248, 250, 18, 258, 17, 14, 37, 20, 35,
	16, 121, 122, 123, 124, 263, 262, 129, 130, 264,
	15, 266, 220, 268, 222, 36, 265, 33, 267, 22,
	34, 21, 274, 3, 232, 0, 278, 0, 280, 0,
	0, 0, 0, 283, 0, 0, 0, 242, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 291, 255, 0,
	293, 257, 0, 295, 259, 297, 0, 0, 0, 170,
	171, 172, 173, 174, 175, 176, 177, 178, 0, 0,
	181, 0, 0, 0, 271, 187, 0, 189, 190, 191,
	0, 0, 0, 92, 0, 201, 0, 0, 282, 0,
	0, 0, 0, 99, 285, 0, 100, 10, 11, 12,
	53, 288, 0, 0, 41, 203, 52, 206, 205, 0,
	42, 43, 0, 0, 299, 54, 93, 94, 95, 0,
	0, 0, 44, 45, 98, 96, 97, 46, 47, 48,
	49, 234, 166, 165, 0, 0, 0, 0, 0, 0,
	0, 0, 50, 0, 51, 0, 40, 39, 10, 11,
	12, 53, 0, 256, 0, 41, 290, 52, 0, 0,
	0, 42, 43, 0, 0, 0, 54, 0, 0, 0,
	0, 0, 0, 44, 45, 269, 270, 0, 46, 47,
	48, 49, 0, 166, 165, 0, 0, 0, 0, 0,
	281, 0, 0, 50, 0, 51, 0, 40, 39, 0,
	286, 0, 10, 11, 12, 53, 0, 0, 0, 41,
	273, 52, 0, 0, 0, 42, 43, 298, 0, 300,
	54, 0, 0, 0, 0, 0, 0, 44, 45, 0,
	0, 0, 46, 47, 48, 49, 0, 166, 165, 0,
	0, 0, 0, 0, 0, 0, 0, 50, 0, 51,
	0, 40, 39, 10, 11, 12, 53, 0, 0, 0,
	41, 245, 52, 0, 0, 0, 42, 43, 0, 0,
	0, 54, 0, 0, 0, 0, 0, 0, 44, 45,
	0, 0, 0, 46, 47, 48, 49, 0, 166, 165,
	0, 0, 0, 0, 0, 0, 0, 0, 50, 0,
	51, 0, 40, 39, 10, 11, 12, 53, 0, 0,
	0, 41, 243, 52, 0, 0, 0, 42, 43, 0,
	0, 0, 54, 0, 0, 0, 0, 0, 0, 44,
	45, 0, 0, 0, 46, 47, 48, 49, 0, 166,
	165, 0, 0, 0, 0, 0, 0, 0, 0, 50,
	0, 51, 0, 40, 39, 10, 11, 12, 53, 0,
	0, 0, 41, 182, 52, 0, 0, 0, 42, 43,
	0, 0, 0, 54, 0, 0, 0, 0, 0, 0,
	44, 45, 0, 0, 0, 46, 47, 48, 49, 0,
	166, 165, 0, 0, 10, 11, 12, 53, 133, 0,
	50, 41, 51, 52, 40, 39, 0, 42, 43, 0,
	0, 0, 54, 0, 0, 0, 0, 0, 0, 44,
	45, 0, 0, 0, 46, 47, 48, 49, 0, 7,
	0, 131, 0, 10, 11, 12, 53, 0, 0, 50,
	41, 51, 52, 40, 39, 0, 42, 43, 0, 0,
	0, 54, 0, 0, 0, 0, 0, 0, 44, 45,
	0, 0, 0, 46, 47, 48, 49, 0, 166, 165,
	0, 0, 0, 0, 0, 0, 0, 0, 50, 0,
	51, 0, 40, 39, 10, 11, 12, 53, 0, 0,
	106, 41, 0, 52, 0, 0, 0, 42, 43, 0,
	0, 0, 54, 0, 0, 0, 0, 0, 0, 44,
	45, 0, 0, 103, 46, 47, 48, 49, 0, 0,
	0, 0, 0, 10, 11, 12, 53, 0, 0, 50,
	41, 51, 52, 40, 39, 0, 42, 43, 0, 0,
	0, 54, 0, 0, 0, 0, 0, 0, 44, 45,
	0, 0, 0, 46, 47, 48, 49, 0, 7, 0,
	88, 0, 0, 0, 0, 0, 0, 0, 50, 0,
	51, 0, 40, 39, 10, 11, 12, 53, 0, 0,
	106, 41, 0, 52, 0, 0, 0, 42, 43, 0,
	0, 0, 54, 0, 0, 0, 0, 0, 0, 44,
	45, 0, 0, 0, 46, 47, 48, 49, 0, 0,
	0, 0, 0, 10, 11, 12, 53, 0, 0, 50,
	41, 51, 52, 40, 39, 0, 42, 43, 0, 0,
	0, 54, 0, 0, 0, 0, 0, 0, 44, 45,
	0, 0, 0, 46, 47, 48, 49, 0, 0, 0,
	0, 0, 10, 202, 12, 53, 0, 0, 50, 41,
	51, 52, 40, 39, 0, 42, 43, 0, 0, 0,
	54, 0, 0, 0, 0, 0, 0, 44, 45, 0,
	0, 0, 46, 47, 48, 49, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 50, 0, 51,
	0, 40, 39,
}
var RubyPact = []int{

	-48, -12, -1000, -52, -1000, -1000, 848, 65, -1000, -19,
	-1000, 70, -13, 65, 65, 65, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 11,
	159, 65, 65, 65, -1000, -1000, 65, 65, 65, 65,
	65, 65, 65, 758, 51, -9, -1000, 319, 170, -1000,
	719, 65, 65, -1000, -1000, -1000, 126, 168, 75, 74,
	72, -1000, -1000, 135, -1000, -1000, 161, 143, 143, 848,
	848, 848, 848, 183, -1000, 848, 629, -19, -1000, -1000,
	-1000, -10, 65, 65, 65, 65, 65, 65, 65, 65,
	65, 65, 65, 65, -19, -1000, 65, 183, 127, -1000,
	65, 65, 65, 65, 65, -1000, -1000, 65, 65, -1000,
	65, -19, -19, -19, -19, 65, -1000, -1000, -3, -19,
	-19, -1000, -1000, -1000, 319, 668, 151, 848, 848, 848,
	848, 848, 848, 848, 848, 848, 149, 125, 848, 590,
	-21, 65, 65, 848, 149, 848, 848, 848, 149, 14,
	50, 13, -1000, 887, 332, -1000, -1000, -19, 103, -1000,
	-19, -19, -19, -19, -19, -19, -19, -19, -19, -1000,
	65, -19, -1000, 65, 180, 138, 133, -19, -1000, -19,
	-19, -19, -1000, -1000, -1000, 65, -1000, 65, -1000, 65,
	65, -19, 77, -1000, 109, 65, -1000, -1000, 17, 809,
	40, -23, -1000, 65, -1000, -1000, -1000, -1000, -1000, 48,
	539, 143, 488, 177, 120, 118, 132, 65, -1000, 65,
	-1000, 848, 668, 65, -19, -1000, 24, -1000, -1000, 113,
	-1000, -1000, 668, -1000, 65, -1000, -1000, -1000, -3, -1000,
	-3, -1000, 65, 848, 848, 668, -19, 668, 136, 437,
	-1000, 65, 32, -6, -33, -3, -42, -3, 848, -19,
	-19, 668, 65, -1000, 167, -1000, -1000, -1000, 848, -1000,
	156, -19, 668, 176, -1000, 383, -19, -11, 668, 65,
	-1000, 22, 65, 139, 65, 848, -1000, 848, 46, 668,
	30, -1000, -1000,
}
var RubyPgo = []int{

	0, 265, 161, 263, 261, 260, 7, 259, 257, 255,
	250, 240, 239, 238, 237, 236, 235, 233, 229, 227,
	1, 226, 225, 224, 221, 220, 219, 218, 216, 213,
	212, 57, 211, 4, 210, 209, 208, 207, 205, 199,
	19, 197, 195, 193, 192, 191, 0, 6, 3, 190,
}
var RubyR1 = []int{

	0, 37, 37, 37, 37, 3, 3, 3, 31, 31,
	31, 31, 46, 46, 47, 47, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 11, 11, 11, 11, 11,
	11, 11, 11, 33, 33, 44, 44, 44, 44, 43,
	43, 43, 43, 43, 43, 43, 40, 40, 40, 40,
	40, 48, 48, 48, 16, 36, 36, 17, 17, 19,
	20, 20, 45, 45, 13, 13, 13, 13, 13, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 4,
	7, 8, 5, 5, 38, 38, 38, 38, 42, 42,
	42, 1, 1, 10, 10, 18, 18, 15, 15, 6,
	6, 34, 41, 41, 41, 49, 49, 12, 12, 12,
	12, 35, 35, 35, 35, 35, 32, 32, 32, 32,
	32, 32, 32, 9, 14, 39, 39, 39, 39,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 1, 3, 0, 2,
	2, 2, 0, 2, 0, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 3, 3, 3, 5,
	5, 3, 5, 5, 1, 1, 1, 5, 5, 1,
	1, 1, 5, 5, 5, 5, 0, 1, 1, 5,
	5, 0, 2, 2, 9, 0, 1, 7, 11, 7,
	1, 4, 1, 4, 5, 5, 5, 5, 5, 3,
	3, 3, 3, 5, 5, 5, 5, 5, 5, 1,
	1, 5, 9, 9, 0, 5, 10, 11, 4, 9,
	10, 0, 1, 2, 2, 2, 2, 3, 3, 3,
	7, 3, 0, 1, 5, 1, 2, 5, 6, 5,
	5, 0, 5, 3, 4, 2, 0, 1, 1, 1,
	2, 2, 2, 3, 5, 0, 4, 7, 10,
}
var RubyChk = []int{

	-1000, -37, 57, -3, 41, 42, -46, 40, 57, -2,
	5, 6, 7, -18, -15, -10, -11, -16, -17, -19,
	-13, -4, -7, -21, -22, -23, -24, -25, -26, -27,
	-28, -29, -30, -8, -5, -12, -9, -14, -20, 55,
	54, 12, 18, 19, 30, 31, 35, 36, 37, 38,
	50, 52, 14, 8, 23, -45, -46, -46, 44, -33,
	-46, 8, -43, 6, 7, 5, -46, 44, -46, -46,
	-46, 6, 7, 55, 6, 7, -46, -46, -46, -46,
	-46, -46, -46, -46, -46, -46, -32, -2, 42, -46,
	41, 43, 4, 37, 38, 39, 46, 47, 45, 14,
	17, 6, -44, 34, -2, -6, 11, -46, -46, -33,
	34, 6, 34, 34, 34, 6, 7, 6, -20, 7,
	-20, -2, -2, -2, -2, -40, 6, 5, -47, -2,
	-2, 42, -46, 9, -46, -31, 43, -46, -46, -46,
	-46, -46, -46, -46, -46, -46, -46, -46, -46, -31,
	-46, -40, 10, -46, -46, -46, -46, -46, -46, -46,
	-46, -46, 41, -46, -31, 41, 40, -2, -39, 7,
	-2, -2, -2, -2, -2, -2, -2, -2, -2, -33,
	10, -2, 13, -34, 45, -46, -46, -2, -33, -2,
	-2, -2, -36, -33, 41, 32, 41, 10, 51, -38,
	-42, -2, 6, 13, -35, 16, 15, 13, -47, -46,
	-46, -41, 6, -40, 9, 5, 6, 7, -6, -48,
	-31, -46, -31, -46, -46, -46, -46, 43, 13, 16,
	15, -46, -31, 24, -2, -6, -49, 41, 45, -46,
	40, 41, -31, 13, -20, 13, 5, 6, -47, 10,
	-47, 10, 4, -46, -46, -31, -2, -31, -46, -31,
	41, 10, -48, -46, -46, -47, -46, -47, -46, -2,
	-2, -31, 7, 13, -46, 13, 41, 53, -46, 53,
	-46, -2, -31, -46, 6, -31, -2, 6, -31, 4,
	13, -46, 43, -46, 4, -46, 6, -46, -2, -31,
	-2, 10, 10,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 6, 0, 12, 4, 12,
	16, -2, -2, -2, -2, -2, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 0,
	0, 12, 12, 12, 99, 100, 12, 12, 12, 12,
	12, 12, 12, 12, 0, 0, 13, 7, 0, 45,
	0, 12, -2, 59, 60, 61, 0, 0, 0, 0,
	0, 115, 116, 0, 113, 114, 0, 0, 0, 0,
	0, 0, 0, 66, 14, 0, 0, -2, 138, 139,
	8, 0, 12, 12, 12, 12, 12, 12, 12, 12,
	12, -2, -2, 12, -2, 56, -2, 66, 0, 46,
	12, -2, 12, 12, 12, 117, 118, 12, 12, -2,
	12, -2, -2, -2, -2, 12, 67, 68, 12, -2,
	-2, 141, 142, 143, 0, 145, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 12, 12, 0, 0, 0, 0, 0, 75, 0,
	0, 0, 15, 104, 0, 9, 10, -2, 14, -2,
	-2, -2, -2, -2, -2, -2, -2, -2, -2, 49,
	12, -2, 119, 12, -2, 0, 0, -2, 50, -2,
	-2, -2, 71, 76, 8, 12, 8, 12, 101, 12,
	12, 12, -2, 127, 0, 12, 8, 144, 0, 0,
	0, 0, -2, 12, 53, 62, 63, 64, 65, 8,
	0, 0, 0, 0, 14, 14, 0, 12, 128, 12,
	8, 0, 135, 8, -2, 58, 8, 125, 121, 0,
	72, 73, 71, 77, 12, 79, 69, 70, 12, 14,
	12, 14, 12, 0, 0, 133, -2, 146, 0, 0,
	126, 12, 0, 0, 0, 12, 0, 12, 0, -2,
	-2, 134, 12, 120, 0, 74, 8, 102, 0, 103,
	0, -2, 132, 8, -2, 0, 12, 0, 147, 12,
	78, 0, 12, 0, 12, 0, 8, 0, -2, 148,
	-2, 110, 107,
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
		//line parser.y:152
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:154
		{
			Statements = []ast.Node{}
		}
	case 3:
		//line parser.y:156
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 4:
		//line parser.y:162
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 5:
		//line parser.y:168
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
		//line parser.y:172
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 9:
		//line parser.y:174
		{
		}
	case 10:
		//line parser.y:176
		{
		}
	case 11:
		//line parser.y:178
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
		//line parser.y:190
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 46:
		//line parser.y:197
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 47:
		//line parser.y:204
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 48:
		//line parser.y:211
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 49:
		//line parser.y:218
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 50:
		//line parser.y:226
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 51:
		//line parser.y:234
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 52:
		//line parser.y:241
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 53:
		//line parser.y:250
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 54:
		//line parser.y:252
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 55:
		//line parser.y:255
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 56:
		//line parser.y:257
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 57:
		//line parser.y:259
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 58:
		//line parser.y:261
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 59:
		//line parser.y:264
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 60:
		//line parser.y:266
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 61:
		//line parser.y:268
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 62:
		//line parser.y:270
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 63:
		//line parser.y:272
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 64:
		//line parser.y:274
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 65:
		//line parser.y:276
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 66:
		//line parser.y:278
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 67:
		//line parser.y:280
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:282
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 69:
		//line parser.y:284
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 70:
		//line parser.y:286
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:296
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 75:
		//line parser.y:304
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 76:
		//line parser.y:305
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 77:
		//line parser.y:308
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 78:
		//line parser.y:315
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-8].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-4].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-8].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 79:
		//line parser.y:325
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 80:
		//line parser.y:334
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 81:
		//line parser.y:340
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 82:
		//line parser.y:348
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 83:
		//line parser.y:352
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 84:
		//line parser.y:357
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 85:
		//line parser.y:364
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 86:
		//line parser.y:371
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 87:
		//line parser.y:378
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 88:
		//line parser.y:385
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 89:
		//line parser.y:392
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 90:
		//line parser.y:393
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 91:
		//line parser.y:394
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 92:
		//line parser.y:395
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 93:
		//line parser.y:398
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 94:
		//line parser.y:407
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 95:
		//line parser.y:416
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 96:
		//line parser.y:425
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 97:
		//line parser.y:434
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 98:
		//line parser.y:443
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 99:
		//line parser.y:451
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 100:
		//line parser.y:452
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 101:
		//line parser.y:454
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 102:
		//line parser.y:458
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 103:
		//line parser.y:466
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 104:
		//line parser.y:474
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 105:
		//line parser.y:476
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 106:
		//line parser.y:483
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 107:
		//line parser.y:490
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 108:
		//line parser.y:498
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 109:
		//line parser.y:505
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 110:
		//line parser.y:512
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 111:
		//line parser.y:520
		{
			RubyVAL.genericValue = nil
		}
	case 112:
		//line parser.y:521
		{
			RubyVAL.genericValue = nil
		}
	case 113:
		//line parser.y:524
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 114:
		//line parser.y:526
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 115:
		//line parser.y:529
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 116:
		//line parser.y:531
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 117:
		//line parser.y:534
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 118:
		//line parser.y:536
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 119:
		//line parser.y:539
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 120:
		//line parser.y:541
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 121:
		//line parser.y:549
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 122:
		//line parser.y:551
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 123:
		//line parser.y:553
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 124:
		//line parser.y:555
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 127:
		//line parser.y:560
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 128:
		//line parser.y:567
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 129:
		//line parser.y:575
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 130:
		//line parser.y:582
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 131:
		//line parser.y:589
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 132:
		//line parser.y:591
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 133:
		//line parser.y:598
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 134:
		//line parser.y:605
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 135:
		//line parser.y:612
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 136:
		//line parser.y:619
		{
		}
	case 137:
		//line parser.y:620
		{
			RubyVAL.genericSlice = []ast.Node{RubyS[Rubypt-0].genericValue}
		}
	case 138:
		//line parser.y:621
		{
		}
	case 139:
		//line parser.y:622
		{
		}
	case 140:
		//line parser.y:623
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 141:
		//line parser.y:624
		{
		}
	case 142:
		//line parser.y:625
		{
		}
	case 143:
		//line parser.y:628
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 144:
		//line parser.y:631
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 145:
		//line parser.y:639
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 146:
		//line parser.y:641
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 147:
		//line parser.y:643
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				},
			})
		}
	case 148:
		//line parser.y:652
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
