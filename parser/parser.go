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

//line parser.y:621

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
	-1, 60,
	10, 12,
	-2, 53,
	-1, 85,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 139,
	-1, 97,
	5, 12,
	6, 12,
	7, 12,
	8, 12,
	-2, 46,
	-1, 98,
	10, 12,
	-2, 50,
	-1, 100,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 54,
	-1, 102,
	45, 12,
	-2, 8,
	-1, 107,
	5, 12,
	6, 12,
	7, 12,
	8, 12,
	-2, 47,
	-1, 117,
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
	-1, 118,
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
	-1, 119,
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
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 8,
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
	-2, 142,
	-1, 155,
	43, 83,
	-2, 12,
	-1, 162,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 51,
	-1, 163,
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
	-2, 95,
	-1, 165,
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
	-1, 166,
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
	-1, 167,
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
	-1, 168,
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
	-1, 169,
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
	-2, 132,
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
	-1, 176,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 11,
	-1, 179,
	45, 124,
	-2, 66,
	-1, 182,
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
	-1, 184,
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
	-1, 185,
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
	-1, 186,
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
	-1, 199,
	4, 17,
	17, 17,
	39, 17,
	44, 17,
	45, 17,
	46, 17,
	47, 17,
	-2, 12,
	-1, 207,
	45, 125,
	-2, 67,
	-1, 230,
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
	-1, 240,
	43, 84,
	-2, 12,
	-1, 252,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 8,
	-1, 264,
	10, 109,
	41, 109,
	53, 109,
	-2, 12,
	-1, 265,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 8,
	-1, 276,
	45, 126,
	-2, 70,
	-1, 279,
	10, 106,
	41, 106,
	53, 106,
	-2, 12,
	-1, 285,
	41, 110,
	53, 110,
	-2, 12,
	-1, 288,
	41, 107,
	53, 107,
	-2, 12,
}

const RubyNprod = 146
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 934

var RubyAct = []int{

	130, 215, 6, 101, 57, 210, 211, 212, 54, 121,
	55, 102, 58, 64, 66, 67, 68, 8, 143, 241,
	276, 124, 241, 242, 2, 114, 69, 70, 272, 270,
	234, 194, 289, 63, 61, 62, 59, 179, 7, 4,
	5, 74, 75, 76, 287, 281, 77, 78, 79, 80,
	81, 82, 269, 83, 87, 7, 213, 254, 7, 56,
	103, 104, 7, 7, 7, 159, 56, 65, 7, 105,
	243, 224, 195, 243, 7, 71, 115, 115, 56, 236,
	237, 176, 110, 218, 192, 128, 236, 237, 9, 131,
	132, 133, 134, 135, 136, 137, 138, 139, 140, 141,
	142, 233, 116, 144, 217, 157, 109, 147, 148, 149,
	150, 151, 154, 145, 152, 63, 61, 62, 59, 108,
	284, 153, 158, 263, 225, 160, 227, 226, 255, 63,
	61, 62, 59, 247, 232, 85, 209, 194, 245, 172,
	100, 146, 111, 112, 161, 171, 180, 181, 72, 73,
	7, 240, 155, 183, 278, 115, 191, 188, 106, 117,
	118, 119, 120, 123, 207, 125, 126, 63, 61, 62,
	59, 113, 189, 204, 123, 122, 193, 107, 97, 205,
	190, 156, 98, 60, 197, 214, 206, 196, 1, 208,
	187, 201, 178, 84, 32, 220, 31, 221, 222, 223,
	58, 30, 29, 228, 28, 27, 26, 25, 231, 235,
	24, 23, 35, 162, 163, 164, 165, 166, 167, 168,
	169, 170, 229, 19, 173, 249, 13, 250, 18, 182,
	17, 184, 185, 186, 238, 14, 239, 20, 36, 16,
	256, 258, 198, 244, 246, 259, 251, 261, 15, 37,
	33, 253, 22, 34, 21, 3, 268, 0, 0, 0,
	0, 271, 0, 273, 274, 0, 0, 260, 0, 262,
	0, 266, 0, 0, 0, 0, 0, 0, 280, 0,
	0, 0, 283, 0, 275, 286, 230, 0, 10, 11,
	12, 53, 0, 0, 0, 40, 200, 52, 203, 202,
	0, 41, 42, 0, 0, 0, 0, 0, 0, 0,
	252, 0, 0, 43, 44, 0, 0, 0, 45, 46,
	47, 48, 0, 175, 174, 0, 0, 0, 0, 0,
	0, 264, 265, 49, 0, 50, 0, 39, 38, 51,
	0, 10, 11, 12, 53, 0, 0, 0, 40, 267,
	52, 0, 0, 277, 41, 42, 279, 0, 0, 0,
	0, 0, 0, 0, 0, 285, 43, 44, 288, 0,
	88, 45, 46, 47, 48, 0, 175, 174, 0, 0,
	95, 0, 0, 96, 0, 0, 49, 0, 50, 0,
	39, 38, 51, 10, 11, 12, 53, 0, 0, 0,
	40, 257, 52, 89, 90, 91, 41, 42, 0, 0,
	0, 94, 92, 93, 0, 0, 0, 0, 43, 44,
	0, 0, 0, 45, 46, 47, 48, 0, 175, 174,
	0, 0, 0, 0, 0, 0, 0, 0, 49, 0,
	50, 0, 39, 38, 51, 10, 11, 12, 53, 0,
	0, 0, 40, 219, 52, 0, 0, 0, 41, 42,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	43, 44, 0, 0, 0, 45, 46, 47, 48, 0,
	175, 174, 0, 0, 0, 0, 0, 0, 0, 0,
	49, 0, 50, 0, 39, 38, 51, 10, 11, 12,
	53, 0, 0, 0, 40, 216, 52, 0, 0, 0,
	41, 42, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 43, 44, 0, 0, 0, 45, 46, 47,
	48, 0, 175, 174, 0, 0, 0, 0, 0, 0,
	0, 0, 49, 0, 50, 0, 39, 38, 51, 10,
	11, 12, 53, 0, 0, 0, 40, 177, 52, 0,
	0, 0, 41, 42, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 43, 44, 0, 0, 0, 45,
	46, 47, 48, 0, 175, 174, 0, 0, 10, 11,
	12, 53, 129, 0, 49, 40, 50, 52, 39, 38,
	51, 41, 42, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 43, 44, 0, 0, 0, 45, 46,
	47, 48, 0, 7, 0, 127, 0, 10, 11, 12,
	53, 0, 0, 49, 40, 50, 52, 39, 38, 51,
	41, 42, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 43, 44, 0, 0, 0, 45, 46, 47,
	48, 0, 175, 174, 0, 0, 0, 0, 0, 0,
	0, 0, 49, 0, 50, 0, 39, 38, 51, 10,
	11, 12, 53, 0, 0, 102, 40, 0, 52, 0,
	0, 0, 41, 42, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 43, 44, 0, 0, 99, 45,
	46, 47, 48, 0, 0, 0, 0, 0, 10, 11,
	12, 53, 0, 0, 49, 40, 50, 52, 39, 38,
	51, 41, 42, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 43, 44, 0, 0, 0, 45, 46,
	47, 48, 0, 7, 0, 86, 0, 0, 0, 0,
	0, 0, 0, 49, 0, 50, 0, 39, 38, 51,
	10, 11, 12, 53, 0, 0, 102, 40, 0, 52,
	0, 0, 0, 41, 42, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 43, 44, 0, 0, 0,
	45, 46, 47, 48, 0, 0, 0, 0, 0, 10,
	11, 12, 53, 0, 0, 49, 40, 50, 52, 39,
	38, 51, 41, 42, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 43, 44, 0, 0, 0, 45,
	46, 47, 48, 0, 0, 0, 0, 0, 10, 199,
	12, 53, 0, 0, 49, 40, 50, 52, 39, 38,
	51, 41, 42, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 88, 43, 44, 0, 0, 0, 45, 46,
	47, 48, 95, 0, 0, 96, 0, 0, 0, 0,
	88, 0, 0, 49, 0, 50, 0, 39, 38, 51,
	95, 0, 282, 96, 0, 89, 90, 91, 0, 0,
	0, 0, 0, 94, 92, 93, 0, 0, 0, 0,
	248, 0, 0, 89, 90, 91, 0, 0, 0, 0,
	0, 94, 92, 93,
}
var RubyPact = []int{

	-33, -2, -1000, -40, -1000, -1000, 804, 18, -1000, 15,
	-1000, 110, 23, 18, 18, 18, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 20, 142,
	18, 18, 18, -1000, -1000, 18, 18, 18, 18, 18,
	18, -1000, 18, 713, -1000, 366, 172, -1000, 674, 18,
	18, -1000, -1000, -1000, 124, 171, 85, 72, 48, -1000,
	-1000, 136, -1000, -1000, 165, 18, 18, 804, 804, 804,
	804, 169, -1000, 804, 583, 15, -1000, -1000, 18, 18,
	18, 18, 18, 18, 18, 18, 18, 18, 18, 18,
	15, -1000, 18, 169, 131, -1000, 18, 18, 18, 18,
	18, -1000, -1000, 18, 80, 145, 64, 15, 15, 15,
	15, 18, -1000, -1000, 24, 15, 15, -1000, -1000, -1000,
	366, 804, 804, 804, 804, 804, 804, 804, 804, 804,
	162, 129, 804, 544, -8, 18, 18, 804, 162, 804,
	804, 804, 162, -1000, 18, 18, 41, -1000, 21, -1000,
	843, 283, 15, 15, 15, 15, 15, 15, 15, 15,
	15, -1000, 18, 15, -1000, -1000, 15, -1000, 18, 158,
	127, 0, 15, -1000, 15, 15, 15, -1000, -1000, 492,
	63, -1000, 40, 440, 18, -1000, 18, 18, 15, 28,
	-1000, 111, 18, -1000, 765, 60, -15, -1000, 18, -1000,
	-1000, -1000, -1000, -1000, -1000, 46, -1000, -1000, 144, -1000,
	17, 128, 123, 886, 18, -1000, 18, -1000, 804, 622,
	15, -1000, 16, -1000, -1000, 118, -1000, -1000, 622, 388,
	18, -1000, -1000, -1000, 24, -1000, 24, -1000, 90, 804,
	804, 622, 15, 336, -1000, 18, 39, -1000, -1000, -24,
	24, -25, 24, 18, 15, 15, 622, -1000, 14, -1000,
	-1000, 804, -1000, 148, 804, 622, -1000, 15, 2, 15,
	868, 18, 87, 804, 18, 34, 804, -1000, 22, -1000,
}
var RubyPgo = []int{

	0, 257, 81, 255, 254, 253, 3, 252, 250, 249,
	248, 239, 238, 237, 235, 230, 228, 226, 223, 25,
	212, 211, 210, 207, 206, 205, 204, 202, 201, 196,
	194, 18, 193, 4, 192, 191, 190, 188, 187, 9,
	186, 184, 183, 182, 181, 0, 21, 1, 134,
}
var RubyR1 = []int{

	0, 37, 37, 37, 37, 3, 3, 3, 31, 31,
	31, 31, 45, 45, 46, 46, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 11, 11, 11, 11, 11, 11,
	11, 11, 33, 33, 43, 43, 43, 43, 42, 42,
	42, 42, 42, 42, 42, 42, 39, 39, 39, 39,
	39, 39, 47, 47, 47, 15, 36, 36, 16, 16,
	18, 19, 19, 44, 44, 13, 13, 13, 13, 13,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	4, 7, 8, 5, 5, 38, 38, 38, 38, 41,
	41, 41, 1, 1, 10, 10, 17, 17, 14, 14,
	20, 6, 6, 34, 40, 40, 40, 48, 48, 12,
	12, 12, 12, 35, 35, 35, 35, 35, 32, 32,
	32, 32, 32, 32, 32, 9,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 1, 3, 0, 2,
	2, 2, 0, 2, 0, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 3, 3, 3, 5, 5,
	3, 5, 5, 1, 1, 1, 5, 5, 1, 1,
	1, 5, 5, 5, 5, 5, 0, 1, 1, 5,
	5, 5, 0, 2, 2, 9, 0, 1, 6, 8,
	6, 3, 6, 1, 4, 5, 5, 5, 5, 5,
	3, 3, 3, 3, 5, 5, 5, 5, 5, 5,
	1, 1, 5, 9, 9, 0, 6, 11, 12, 4,
	9, 10, 0, 1, 2, 2, 2, 2, 3, 3,
	1, 3, 7, 3, 0, 1, 5, 1, 2, 5,
	6, 5, 5, 0, 5, 3, 4, 2, 0, 1,
	1, 1, 2, 2, 2, 3,
}
var RubyChk = []int{

	-1000, -37, 57, -3, 41, 42, -45, 40, 57, -2,
	5, 6, 7, -17, -14, -10, -11, -15, -16, -18,
	-13, -4, -7, -21, -22, -23, -24, -25, -26, -27,
	-28, -29, -30, -8, -5, -20, -12, -9, 55, 54,
	12, 18, 19, 30, 31, 35, 36, 37, 38, 50,
	52, 56, 14, 8, -45, -45, 44, -33, -45, 8,
	-42, 6, 7, 5, -45, 44, -45, -45, -45, 6,
	7, 55, 6, 7, -45, -45, -45, -45, -45, -45,
	-45, -45, -45, -45, -32, -2, 42, -45, 4, 37,
	38, 39, 46, 47, 45, 14, 17, 6, -43, 34,
	-2, -6, 11, -45, -45, -33, 34, 6, 34, 34,
	34, 6, 7, 6, -19, -45, -19, -2, -2, -2,
	-2, -39, 6, 5, -46, -2, -2, 42, -45, 9,
	-45, -45, -45, -45, -45, -45, -45, -45, -45, -45,
	-45, -45, -45, -31, -45, -39, 10, -45, -45, -45,
	-45, -45, -45, 41, 32, 7, -44, 41, -45, 41,
	-45, -31, -2, -2, -2, -2, -2, -2, -2, -2,
	-2, -33, 10, -2, 41, 40, -2, 13, -34, 45,
	-45, -45, -2, -33, -2, -2, -2, -36, -33, -31,
	-19, -45, 43, -31, 10, 51, -38, -41, -2, 6,
	13, -35, 16, 15, -45, -45, -40, 6, -39, 9,
	5, 6, 7, 56, -6, -47, 13, 41, 43, 13,
	-45, -45, -45, -45, 43, 13, 16, 15, -45, -31,
	-2, -6, -48, 41, 45, -45, 40, 41, -31, -31,
	7, 5, 6, 56, -46, 10, -46, 10, 34, -45,
	-45, -31, -2, -31, 41, 10, -47, 13, -45, -45,
	-46, -45, -46, 33, -2, -2, -31, 13, -45, 13,
	53, -45, 53, -45, -45, -31, 6, -2, 6, -2,
	-45, 43, 34, -45, 33, -2, -45, 10, -2, 10,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 6, 0, 12, 4, 12,
	16, -2, -2, -2, -2, -2, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 0, 0,
	12, 12, 12, 100, 101, 12, 12, 12, 12, 12,
	12, 120, 12, 12, 13, 7, 0, 44, 0, 12,
	-2, 58, 59, 60, 0, 0, 0, 0, 0, 116,
	117, 0, 114, 115, 0, 12, 12, 0, 0, 0,
	0, 66, 14, 0, 0, -2, 140, 141, 12, 12,
	12, 12, 12, 12, 12, 12, 12, -2, -2, 12,
	-2, 55, -2, 66, 0, 45, 12, -2, 12, 12,
	12, 118, 119, 12, 0, 0, 0, -2, -2, -2,
	-2, 12, 67, 68, 12, -2, -2, 143, 144, 145,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 12, 12, 0, 0, 0,
	0, 0, 76, 8, 12, -2, 0, 8, 0, 15,
	105, 0, -2, -2, -2, -2, -2, -2, -2, -2,
	-2, 48, 12, -2, 9, 10, -2, 121, 12, -2,
	0, 0, -2, 49, -2, -2, -2, 72, 77, 0,
	0, 81, 0, 0, 12, 102, 12, 12, 12, -2,
	129, 0, 12, 8, 0, 0, 0, -2, 12, 52,
	61, 62, 63, 64, 65, 8, 78, 8, 0, 80,
	0, 14, 14, 0, 12, 130, 12, 8, 0, 137,
	-2, 57, 8, 127, 123, 0, 73, 74, 72, 0,
	-2, 69, 70, 71, 12, 14, 12, 14, 0, 0,
	0, 135, -2, 0, 128, 12, 0, 79, 82, 0,
	12, 0, 12, 12, -2, -2, 136, 122, 0, 75,
	103, 0, 104, 0, 0, 134, -2, 12, 0, -2,
	0, 12, 0, 0, 12, -2, 0, 111, -2, 108,
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
		//line parser.y:167
		{
			RubyVAL.genericValue = nil
		}
	case 7:
		//line parser.y:168
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 8:
		//line parser.y:171
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 9:
		//line parser.y:173
		{
		}
	case 10:
		//line parser.y:175
		{
		}
	case 11:
		//line parser.y:177
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
		//line parser.y:189
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 45:
		//line parser.y:196
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 46:
		//line parser.y:203
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 47:
		//line parser.y:210
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 48:
		//line parser.y:217
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 49:
		//line parser.y:225
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 50:
		//line parser.y:233
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 51:
		//line parser.y:240
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 52:
		//line parser.y:249
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 53:
		//line parser.y:251
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 54:
		//line parser.y:254
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 55:
		//line parser.y:256
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 56:
		//line parser.y:258
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 57:
		//line parser.y:260
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 58:
		//line parser.y:263
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 59:
		//line parser.y:265
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 60:
		//line parser.y:267
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 61:
		//line parser.y:269
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 62:
		//line parser.y:271
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 63:
		//line parser.y:273
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 64:
		//line parser.y:275
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 65:
		//line parser.y:277
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 66:
		//line parser.y:279
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 67:
		//line parser.y:281
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:283
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 69:
		//line parser.y:285
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 70:
		//line parser.y:287
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 71:
		//line parser.y:289
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 75:
		//line parser.y:299
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 76:
		//line parser.y:307
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 77:
		//line parser.y:308
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 78:
		//line parser.y:311
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 79:
		//line parser.y:318
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 80:
		//line parser.y:328
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 81:
		//line parser.y:337
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 82:
		//line parser.y:343
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 83:
		//line parser.y:351
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 84:
		//line parser.y:355
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 85:
		//line parser.y:360
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 86:
		//line parser.y:367
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 87:
		//line parser.y:374
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 88:
		//line parser.y:381
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 89:
		//line parser.y:388
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 90:
		//line parser.y:395
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 91:
		//line parser.y:396
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 92:
		//line parser.y:397
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 93:
		//line parser.y:398
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 94:
		//line parser.y:401
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 95:
		//line parser.y:410
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 96:
		//line parser.y:419
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 97:
		//line parser.y:428
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 98:
		//line parser.y:437
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 99:
		//line parser.y:446
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 100:
		//line parser.y:454
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 101:
		//line parser.y:455
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 102:
		//line parser.y:457
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 103:
		//line parser.y:461
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 104:
		//line parser.y:469
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 105:
		//line parser.y:477
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 106:
		//line parser.y:479
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 107:
		//line parser.y:481
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 108:
		//line parser.y:483
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 109:
		//line parser.y:486
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 110:
		//line parser.y:493
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 111:
		//line parser.y:500
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 112:
		//line parser.y:508
		{
			RubyVAL.genericValue = nil
		}
	case 113:
		//line parser.y:509
		{
			RubyVAL.genericValue = nil
		}
	case 114:
		//line parser.y:512
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 115:
		//line parser.y:514
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 116:
		//line parser.y:517
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 117:
		//line parser.y:519
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 118:
		//line parser.y:522
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 119:
		//line parser.y:524
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 120:
		//line parser.y:527
		{
			RubyVAL.genericValue = ast.FileNameConstReference{}
		}
	case 121:
		//line parser.y:530
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 122:
		//line parser.y:532
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 123:
		//line parser.y:540
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 124:
		//line parser.y:542
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 125:
		//line parser.y:544
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 126:
		//line parser.y:546
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 129:
		//line parser.y:551
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 130:
		//line parser.y:558
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 131:
		//line parser.y:566
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 132:
		//line parser.y:573
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 133:
		//line parser.y:580
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 134:
		//line parser.y:582
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 135:
		//line parser.y:589
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 136:
		//line parser.y:596
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 137:
		//line parser.y:603
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 138:
		//line parser.y:610
		{
		}
	case 139:
		//line parser.y:611
		{
			RubyVAL.genericSlice = []ast.Node{RubyS[Rubypt-0].genericValue}
		}
	case 140:
		//line parser.y:612
		{
		}
	case 141:
		//line parser.y:613
		{
		}
	case 142:
		//line parser.y:614
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 143:
		//line parser.y:615
		{
		}
	case 144:
		//line parser.y:616
		{
		}
	case 145:
		//line parser.y:619
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	}
	goto Rubystack /* stack new state and value */
}
