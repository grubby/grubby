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
	-2, 140,
	-1, 100,
	5, 12,
	6, 12,
	7, 12,
	8, 12,
	-2, 47,
	-1, 101,
	10, 12,
	-2, 51,
	-1, 103,
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
	-1, 105,
	45, 12,
	-2, 8,
	-1, 110,
	5, 12,
	6, 12,
	7, 12,
	8, 12,
	-2, 48,
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
	-2, 91,
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
	-2, 92,
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
	-2, 93,
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
	-2, 94,
	-1, 128,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 8,
	-1, 129,
	4, 12,
	14, 12,
	17, 12,
	37, 12,
	38, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 143,
	-1, 159,
	43, 84,
	-2, 12,
	-1, 168,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 11,
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
	-2, 95,
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
	-2, 96,
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
	-2, 97,
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
	-2, 98,
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
	-2, 99,
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
	-2, 100,
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
	-2, 132,
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
	-2, 133,
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
	-2, 86,
	-1, 184,
	45, 125,
	-2, 67,
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
	-2, 87,
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
	-2, 89,
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
	-2, 90,
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
	45, 126,
	-2, 68,
	-1, 238,
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
	-1, 248,
	43, 85,
	-2, 12,
	-1, 260,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 8,
	-1, 274,
	10, 110,
	41, 110,
	53, 110,
	-2, 12,
	-1, 275,
	4, 12,
	17, 12,
	39, 12,
	45, 12,
	46, 12,
	47, 12,
	-2, 8,
	-1, 285,
	10, 107,
	41, 107,
	53, 107,
	-2, 12,
	-1, 288,
	45, 127,
	-2, 71,
	-1, 299,
	41, 111,
	53, 111,
	-2, 12,
	-1, 301,
	41, 108,
	53, 108,
	-2, 12,
}

const RubyNprod = 151
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 996

var RubyAct = []int{

	133, 104, 6, 8, 2, 222, 117, 283, 56, 124,
	57, 281, 60, 66, 68, 69, 70, 217, 218, 219,
	199, 127, 242, 105, 249, 288, 249, 250, 184, 7,
	303, 71, 72, 58, 65, 63, 64, 61, 293, 280,
	59, 7, 76, 77, 78, 67, 225, 79, 80, 81,
	82, 83, 84, 7, 85, 89, 65, 63, 64, 61,
	7, 200, 106, 107, 58, 197, 244, 245, 220, 7,
	7, 163, 231, 264, 302, 251, 241, 251, 118, 118,
	73, 168, 7, 4, 5, 119, 134, 131, 9, 244,
	245, 7, 135, 136, 137, 138, 139, 140, 141, 142,
	143, 144, 145, 146, 7, 295, 148, 108, 58, 237,
	151, 152, 153, 154, 155, 98, 149, 156, 99, 224,
	158, 161, 90, 113, 112, 162, 163, 111, 164, 157,
	232, 209, 234, 233, 216, 199, 87, 265, 92, 93,
	94, 255, 103, 253, 180, 150, 97, 95, 96, 277,
	185, 186, 248, 65, 63, 64, 61, 114, 115, 118,
	196, 120, 121, 122, 123, 195, 159, 128, 129, 65,
	63, 64, 61, 74, 75, 126, 214, 126, 125, 297,
	290, 211, 109, 116, 212, 179, 110, 100, 221, 291,
	240, 210, 147, 188, 215, 160, 101, 193, 62, 202,
	227, 213, 228, 229, 230, 60, 169, 201, 235, 1,
	192, 206, 183, 239, 86, 165, 243, 170, 171, 172,
	173, 174, 175, 176, 177, 178, 32, 31, 181, 30,
	29, 28, 257, 187, 258, 189, 190, 191, 262, 27,
	26, 25, 24, 23, 194, 35, 203, 19, 198, 268,
	252, 254, 266, 269, 13, 271, 18, 273, 17, 14,
	38, 20, 36, 16, 15, 37, 279, 33, 22, 34,
	21, 282, 256, 284, 3, 270, 0, 272, 287, 0,
	0, 0, 98, 0, 0, 99, 0, 0, 0, 0,
	292, 0, 294, 238, 296, 236, 298, 0, 0, 0,
	0, 0, 0, 0, 0, 92, 93, 94, 0, 246,
	91, 247, 0, 97, 95, 96, 0, 260, 0, 0,
	98, 259, 0, 99, 261, 0, 0, 263, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 274,
	275, 0, 0, 92, 93, 94, 0, 276, 0, 0,
	0, 97, 95, 96, 0, 285, 0, 0, 0, 0,
	0, 0, 286, 0, 289, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 299, 0,
	301, 0, 0, 0, 300, 10, 11, 12, 54, 0,
	0, 0, 41, 205, 53, 208, 207, 0, 42, 43,
	0, 0, 0, 55, 0, 0, 0, 0, 0, 0,
	44, 45, 0, 0, 0, 46, 47, 48, 49, 0,
	167, 166, 0, 0, 0, 0, 0, 0, 0, 0,
	50, 0, 51, 0, 40, 39, 52, 10, 11, 12,
	54, 0, 0, 0, 41, 278, 53, 0, 0, 0,
	42, 43, 0, 0, 0, 55, 0, 0, 0, 0,
	0, 0, 44, 45, 0, 0, 0, 46, 47, 48,
	49, 0, 167, 166, 0, 0, 0, 0, 0, 0,
	0, 0, 50, 0, 51, 0, 40, 39, 52, 10,
	11, 12, 54, 0, 0, 0, 41, 267, 53, 0,
	0, 0, 42, 43, 0, 0, 0, 55, 0, 0,
	0, 0, 0, 0, 44, 45, 0, 0, 0, 46,
	47, 48, 49, 0, 167, 166, 0, 0, 0, 0,
	0, 0, 0, 0, 50, 0, 51, 0, 40, 39,
	52, 10, 11, 12, 54, 0, 0, 0, 41, 226,
	53, 0, 0, 0, 42, 43, 0, 0, 0, 55,
	0, 0, 0, 0, 0, 0, 44, 45, 0, 0,
	0, 46, 47, 48, 49, 0, 167, 166, 0, 0,
	0, 0, 0, 0, 0, 0, 50, 0, 51, 0,
	40, 39, 52, 10, 11, 12, 54, 0, 0, 0,
	41, 223, 53, 0, 0, 0, 42, 43, 0, 0,
	0, 55, 0, 0, 0, 0, 0, 0, 44, 45,
	0, 0, 0, 46, 47, 48, 49, 0, 167, 166,
	0, 0, 0, 0, 0, 0, 0, 0, 50, 0,
	51, 0, 40, 39, 52, 10, 11, 12, 54, 0,
	0, 0, 41, 182, 53, 0, 0, 0, 42, 43,
	0, 0, 0, 55, 0, 0, 0, 0, 0, 0,
	44, 45, 0, 0, 0, 46, 47, 48, 49, 0,
	167, 166, 0, 0, 10, 11, 12, 54, 132, 0,
	50, 41, 51, 53, 40, 39, 52, 42, 43, 0,
	0, 0, 55, 0, 0, 0, 0, 0, 0, 44,
	45, 0, 0, 0, 46, 47, 48, 49, 0, 7,
	0, 130, 0, 10, 11, 12, 54, 0, 0, 50,
	41, 51, 53, 40, 39, 52, 42, 43, 0, 0,
	0, 55, 0, 0, 0, 0, 0, 0, 44, 45,
	0, 0, 0, 46, 47, 48, 49, 0, 167, 166,
	0, 0, 0, 0, 0, 0, 0, 0, 50, 0,
	51, 0, 40, 39, 52, 10, 11, 12, 54, 0,
	0, 105, 41, 0, 53, 0, 0, 0, 42, 43,
	0, 0, 0, 55, 0, 0, 0, 0, 0, 0,
	44, 45, 0, 0, 102, 46, 47, 48, 49, 0,
	0, 0, 0, 0, 10, 11, 12, 54, 0, 0,
	50, 41, 51, 53, 40, 39, 52, 42, 43, 0,
	0, 0, 55, 0, 0, 0, 0, 0, 0, 44,
	45, 0, 0, 0, 46, 47, 48, 49, 0, 7,
	0, 88, 0, 0, 0, 0, 0, 0, 0, 50,
	0, 51, 0, 40, 39, 52, 10, 11, 12, 54,
	0, 0, 105, 41, 0, 53, 0, 0, 0, 42,
	43, 0, 0, 0, 55, 0, 0, 0, 0, 0,
	0, 44, 45, 0, 0, 0, 46, 47, 48, 49,
	0, 0, 0, 0, 0, 10, 11, 12, 54, 0,
	0, 50, 41, 51, 53, 40, 39, 52, 42, 43,
	0, 0, 0, 55, 0, 0, 0, 0, 0, 0,
	44, 45, 0, 0, 0, 46, 47, 48, 49, 0,
	0, 0, 0, 0, 10, 204, 12, 54, 0, 0,
	50, 41, 51, 53, 40, 39, 52, 42, 43, 0,
	0, 0, 55, 0, 0, 0, 0, 0, 0, 44,
	45, 0, 0, 0, 46, 47, 48, 49, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 50,
	0, 51, 0, 40, 39, 52,
}
var RubyPact = []int{

	-53, 42, -1000, -54, -1000, -1000, 900, 13, -1000, -11,
	-1000, 51, 1, 13, 13, 13, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 25,
	167, 13, 13, 13, -1000, -1000, 13, 13, 13, 13,
	13, 13, -1000, 13, 809, 81, -1000, 306, 181, -1000,
	770, 13, 13, -1000, -1000, -1000, 148, 180, 93, 90,
	89, -1000, -1000, 151, -1000, -1000, 177, 13, 13, 900,
	900, 900, 900, 172, -1000, 900, 679, -11, -1000, -1000,
	-1000, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, -11, -1000, 13, 172, 135, -1000, 13,
	13, 13, 13, 13, -1000, -1000, 13, 88, 159, 80,
	-11, -11, -11, -11, 13, -1000, -1000, 30, -11, -11,
	-1000, -1000, -1000, 306, 718, 900, 900, 900, 900, 900,
	900, 900, 900, 900, 164, 134, 900, 640, -17, 13,
	13, 900, 164, 900, 900, 900, 164, -1000, 13, 13,
	22, -1000, 10, -1000, 939, 380, -1000, -1000, -11, 118,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -1000,
	13, -11, -1000, 13, 170, 125, 12, -11, -1000, -11,
	-11, -11, -1000, -1000, 588, 78, -1000, 3, 536, 13,
	-1000, 13, 13, -11, 29, -1000, 117, 13, -1000, -1000,
	85, 861, 35, -23, -1000, 13, -1000, -1000, -1000, -1000,
	-1000, -1000, 49, -1000, -1000, 145, -1000, 21, 133, 131,
	268, 13, -1000, 13, -1000, 900, 718, 13, -11, -1000,
	32, -1000, -1000, 127, -1000, -1000, 718, 484, 13, -1000,
	-1000, -1000, 30, -1000, 30, -1000, 13, 900, 900, 718,
	-11, 718, 142, 432, -1000, 13, 26, -1000, -1000, -42,
	30, -46, 30, 900, -11, -11, 718, 13, -1000, 19,
	-1000, -1000, 900, -1000, 174, -11, 718, 185, -1000, -11,
	-5, 13, 101, 13, 173, 13, 900, -1000, 900, 64,
	718, 20, -1000, -1000,
}
var RubyPgo = []int{

	0, 276, 81, 274, 270, 269, 1, 268, 267, 265,
	264, 263, 262, 261, 260, 259, 258, 256, 254, 247,
	6, 245, 243, 242, 241, 240, 239, 231, 230, 229,
	227, 226, 86, 214, 40, 212, 211, 210, 209, 207,
	206, 9, 201, 199, 198, 196, 195, 0, 21, 5,
	190,
}
var RubyR1 = []int{

	0, 38, 38, 38, 38, 3, 3, 3, 32, 32,
	32, 32, 47, 47, 48, 48, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 11, 11, 11, 11, 11,
	11, 11, 11, 34, 34, 45, 45, 45, 45, 44,
	44, 44, 44, 44, 44, 44, 44, 41, 41, 41,
	41, 41, 41, 49, 49, 49, 16, 37, 37, 17,
	17, 19, 20, 20, 46, 46, 13, 13, 13, 13,
	13, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	31, 4, 7, 8, 5, 5, 39, 39, 39, 39,
	43, 43, 43, 1, 1, 10, 10, 18, 18, 15,
	15, 21, 6, 6, 35, 42, 42, 42, 50, 50,
	12, 12, 12, 12, 36, 36, 36, 36, 36, 33,
	33, 33, 33, 33, 33, 33, 9, 14, 40, 40,
	40,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 1, 3, 0, 2,
	2, 2, 0, 2, 0, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 3, 3, 3, 5,
	5, 3, 5, 5, 1, 1, 1, 5, 5, 1,
	1, 1, 5, 5, 5, 5, 5, 0, 1, 1,
	5, 5, 5, 0, 2, 2, 9, 0, 1, 6,
	8, 6, 3, 6, 1, 4, 5, 5, 5, 5,
	5, 3, 3, 3, 3, 5, 5, 5, 5, 5,
	5, 1, 1, 5, 9, 9, 0, 5, 10, 11,
	4, 9, 10, 0, 1, 2, 2, 2, 2, 3,
	3, 1, 3, 7, 3, 0, 1, 5, 1, 2,
	5, 6, 5, 5, 0, 5, 3, 4, 2, 0,
	1, 1, 1, 2, 2, 2, 3, 5, 0, 4,
	10,
}
var RubyChk = []int{

	-1000, -38, 57, -3, 41, 42, -47, 40, 57, -2,
	5, 6, 7, -18, -15, -10, -11, -16, -17, -19,
	-13, -4, -7, -22, -23, -24, -25, -26, -27, -28,
	-29, -30, -31, -8, -5, -21, -12, -9, -14, 55,
	54, 12, 18, 19, 30, 31, 35, 36, 37, 38,
	50, 52, 56, 14, 8, 23, -47, -47, 44, -34,
	-47, 8, -44, 6, 7, 5, -47, 44, -47, -47,
	-47, 6, 7, 55, 6, 7, -47, -47, -47, -47,
	-47, -47, -47, -47, -47, -47, -33, -2, 42, -47,
	41, 4, 37, 38, 39, 46, 47, 45, 14, 17,
	6, -45, 34, -2, -6, 11, -47, -47, -34, 34,
	6, 34, 34, 34, 6, 7, 6, -20, -47, -20,
	-2, -2, -2, -2, -41, 6, 5, -48, -2, -2,
	42, -47, 9, -47, -32, -47, -47, -47, -47, -47,
	-47, -47, -47, -47, -47, -47, -47, -32, -47, -41,
	10, -47, -47, -47, -47, -47, -47, 41, 32, 7,
	-46, 41, -47, 41, -47, -32, 41, 40, -2, -40,
	-2, -2, -2, -2, -2, -2, -2, -2, -2, -34,
	10, -2, 13, -35, 45, -47, -47, -2, -34, -2,
	-2, -2, -37, -34, -32, -20, -47, 43, -32, 10,
	51, -39, -43, -2, 6, 13, -36, 16, 15, 13,
	-48, -47, -47, -42, 6, -41, 9, 5, 6, 7,
	56, -6, -49, 13, 41, 43, 13, -47, -47, -47,
	-47, 43, 13, 16, 15, -47, -32, 24, -2, -6,
	-50, 41, 45, -47, 40, 41, -32, -32, 7, 5,
	6, 56, -48, 10, -48, 10, 4, -47, -47, -32,
	-2, -32, -47, -32, 41, 10, -49, 13, -47, -47,
	-48, -47, -48, -47, -2, -2, -32, 7, 13, -47,
	13, 53, -47, 53, -47, -2, -32, -47, 6, -2,
	6, 4, -47, 43, -47, 4, -47, 6, -47, -2,
	-32, -2, 10, 10,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 6, 0, 12, 4, 12,
	16, -2, -2, -2, -2, -2, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 0,
	0, 12, 12, 12, 101, 102, 12, 12, 12, 12,
	12, 12, 121, 12, 12, 0, 13, 7, 0, 45,
	0, 12, -2, 59, 60, 61, 0, 0, 0, 0,
	0, 117, 118, 0, 115, 116, 0, 12, 12, 0,
	0, 0, 0, 67, 14, 0, 0, -2, 141, 142,
	8, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	-2, -2, 12, -2, 56, -2, 67, 0, 46, 12,
	-2, 12, 12, 12, 119, 120, 12, 0, 0, 0,
	-2, -2, -2, -2, 12, 68, 69, 12, -2, -2,
	144, 145, 146, 0, 148, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 12,
	12, 0, 0, 0, 0, 0, 77, 8, 12, -2,
	0, 8, 0, 15, 106, 0, 9, 10, -2, 14,
	-2, -2, -2, -2, -2, -2, -2, -2, -2, 49,
	12, -2, 122, 12, -2, 0, 0, -2, 50, -2,
	-2, -2, 73, 78, 0, 0, 82, 0, 0, 12,
	103, 12, 12, 12, -2, 130, 0, 12, 8, 147,
	0, 0, 0, 0, -2, 12, 53, 62, 63, 64,
	65, 66, 8, 79, 8, 0, 81, 0, 14, 14,
	0, 12, 131, 12, 8, 0, 138, 8, -2, 58,
	8, 128, 124, 0, 74, 75, 73, 0, -2, 70,
	71, 72, 12, 14, 12, 14, 12, 0, 0, 136,
	-2, 149, 0, 0, 129, 12, 0, 80, 83, 0,
	12, 0, 12, 0, -2, -2, 137, 12, 123, 0,
	76, 104, 0, 105, 0, -2, 135, 0, -2, 12,
	0, 12, 0, 12, 0, 12, 0, 8, 0, -2,
	150, -2, 112, 109,
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
		//line parser.y:191
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 46:
		//line parser.y:198
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 47:
		//line parser.y:205
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 48:
		//line parser.y:212
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 49:
		//line parser.y:219
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 50:
		//line parser.y:227
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 51:
		//line parser.y:235
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 52:
		//line parser.y:242
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 53:
		//line parser.y:251
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 54:
		//line parser.y:253
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
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
		//line parser.y:262
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
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:281
		{
			RubyVAL.genericSlice = ast.Nodes{}
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
	case 72:
		//line parser.y:291
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:301
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 77:
		//line parser.y:309
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 78:
		//line parser.y:310
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 79:
		//line parser.y:313
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 80:
		//line parser.y:320
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 81:
		//line parser.y:330
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 82:
		//line parser.y:339
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 83:
		//line parser.y:345
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 84:
		//line parser.y:353
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 85:
		//line parser.y:357
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 86:
		//line parser.y:362
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 87:
		//line parser.y:369
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 88:
		//line parser.y:376
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 89:
		//line parser.y:383
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 90:
		//line parser.y:390
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 91:
		//line parser.y:397
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 92:
		//line parser.y:398
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 93:
		//line parser.y:399
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 94:
		//line parser.y:400
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 95:
		//line parser.y:403
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 96:
		//line parser.y:412
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 97:
		//line parser.y:421
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 98:
		//line parser.y:430
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 99:
		//line parser.y:439
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 100:
		//line parser.y:448
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 101:
		//line parser.y:456
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 102:
		//line parser.y:457
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 103:
		//line parser.y:459
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 104:
		//line parser.y:463
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 105:
		//line parser.y:471
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 106:
		//line parser.y:479
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 107:
		//line parser.y:481
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 108:
		//line parser.y:488
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 109:
		//line parser.y:495
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 110:
		//line parser.y:503
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 111:
		//line parser.y:510
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 112:
		//line parser.y:517
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 113:
		//line parser.y:525
		{
			RubyVAL.genericValue = nil
		}
	case 114:
		//line parser.y:526
		{
			RubyVAL.genericValue = nil
		}
	case 115:
		//line parser.y:529
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 116:
		//line parser.y:531
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 117:
		//line parser.y:534
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 118:
		//line parser.y:536
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 119:
		//line parser.y:539
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 120:
		//line parser.y:541
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 121:
		//line parser.y:544
		{
			RubyVAL.genericValue = ast.FileNameConstReference{}
		}
	case 122:
		//line parser.y:547
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 123:
		//line parser.y:549
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 124:
		//line parser.y:557
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 125:
		//line parser.y:559
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 126:
		//line parser.y:561
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 127:
		//line parser.y:563
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 130:
		//line parser.y:568
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 131:
		//line parser.y:575
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 132:
		//line parser.y:583
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 133:
		//line parser.y:590
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 134:
		//line parser.y:597
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 135:
		//line parser.y:599
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 136:
		//line parser.y:606
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 137:
		//line parser.y:613
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 138:
		//line parser.y:620
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 139:
		//line parser.y:627
		{
		}
	case 140:
		//line parser.y:628
		{
			RubyVAL.genericSlice = []ast.Node{RubyS[Rubypt-0].genericValue}
		}
	case 141:
		//line parser.y:629
		{
		}
	case 142:
		//line parser.y:630
		{
		}
	case 143:
		//line parser.y:631
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 144:
		//line parser.y:632
		{
		}
	case 145:
		//line parser.y:633
		{
		}
	case 146:
		//line parser.y:636
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 147:
		//line parser.y:639
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 148:
		//line parser.y:647
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 149:
		//line parser.y:649
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 150:
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
