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

//line parser.y:713

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 17,
	-1, 12,
	1, 20,
	4, 20,
	9, 20,
	10, 20,
	13, 20,
	15, 20,
	16, 20,
	17, 20,
	24, 20,
	39, 20,
	41, 20,
	42, 20,
	44, 20,
	45, 20,
	46, 20,
	47, 20,
	51, 20,
	53, 20,
	57, 20,
	-2, 17,
	-1, 13,
	34, 17,
	43, 88,
	-2, 21,
	-1, 14,
	34, 17,
	-2, 22,
	-1, 15,
	34, 17,
	-2, 23,
	-1, 16,
	34, 17,
	-2, 24,
	-1, 22,
	14, 17,
	37, 17,
	-2, 37,
	-1, 89,
	1, 36,
	9, 36,
	10, 36,
	12, 36,
	13, 36,
	15, 36,
	16, 36,
	18, 36,
	19, 36,
	23, 36,
	24, 36,
	35, 36,
	36, 36,
	41, 36,
	42, 36,
	51, 36,
	53, 36,
	57, 36,
	-2, 17,
	-1, 102,
	10, 17,
	-2, 60,
	-1, 103,
	4, 17,
	5, 17,
	6, 17,
	7, 17,
	8, 17,
	30, 17,
	31, 17,
	50, 17,
	52, 17,
	54, 17,
	55, 17,
	-2, 72,
	-1, 105,
	1, 20,
	4, 20,
	9, 20,
	10, 20,
	13, 20,
	15, 20,
	16, 20,
	24, 20,
	39, 20,
	41, 20,
	42, 20,
	44, 20,
	45, 20,
	46, 20,
	47, 20,
	51, 20,
	53, 20,
	57, 20,
	-2, 17,
	-1, 106,
	43, 88,
	-2, 21,
	-1, 110,
	5, 17,
	6, 17,
	7, 17,
	8, 17,
	30, 17,
	31, 17,
	50, 17,
	52, 17,
	54, 17,
	55, 17,
	-2, 55,
	-1, 112,
	10, 64,
	-2, 17,
	-1, 115,
	37, 17,
	-2, 66,
	-1, 116,
	10, 17,
	-2, 57,
	-1, 120,
	45, 17,
	-2, 9,
	-1, 132,
	9, 64,
	10, 64,
	51, 64,
	-2, 17,
	-1, 137,
	43, 88,
	-2, 86,
	-1, 185,
	43, 89,
	-2, 87,
	-1, 196,
	4, 17,
	5, 17,
	6, 17,
	7, 17,
	8, 17,
	30, 17,
	31, 17,
	50, 17,
	52, 17,
	54, 17,
	55, 17,
	-2, 99,
	-1, 219,
	4, 17,
	5, 17,
	6, 17,
	7, 17,
	8, 17,
	30, 17,
	31, 17,
	50, 17,
	52, 17,
	54, 17,
	55, 17,
	-2, 100,
	-1, 230,
	10, 17,
	-2, 81,
	-1, 267,
	4, 17,
	5, 17,
	6, 17,
	7, 17,
	8, 17,
	30, 17,
	31, 17,
	50, 17,
	52, 17,
	54, 17,
	55, 17,
	-2, 74,
}

const RubyNprod = 156
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1330

var RubyAct = []int{

	161, 21, 6, 254, 133, 91, 230, 34, 58, 8,
	59, 61, 2, 63, 64, 65, 66, 67, 69, 70,
	314, 312, 7, 68, 209, 222, 214, 275, 7, 119,
	113, 317, 60, 166, 7, 4, 5, 7, 329, 253,
	74, 75, 134, 77, 78, 79, 80, 81, 82, 83,
	84, 88, 76, 7, 183, 265, 168, 149, 283, 284,
	283, 284, 234, 111, 7, 224, 223, 71, 60, 320,
	114, 233, 183, 296, 274, 235, 62, 90, 125, 124,
	136, 138, 114, 123, 126, 122, 146, 7, 121, 260,
	148, 262, 261, 150, 151, 152, 153, 154, 155, 156,
	157, 158, 159, 160, 240, 130, 162, 125, 7, 319,
	286, 163, 340, 167, 169, 339, 170, 171, 172, 114,
	286, 174, 175, 176, 177, 178, 179, 180, 266, 209,
	181, 182, 280, 167, 184, 278, 186, 187, 210, 188,
	244, 209, 165, 209, 206, 232, 334, 231, 127, 128,
	11, 12, 13, 50, 147, 72, 73, 42, 310, 49,
	137, 185, 324, 43, 44, 318, 208, 114, 51, 207,
	232, 135, 110, 276, 332, 37, 38, 327, 173, 281,
	45, 46, 47, 48, 273, 7, 192, 145, 41, 116,
	205, 102, 226, 10, 129, 39, 193, 40, 241, 36,
	35, 189, 225, 229, 1, 237, 242, 243, 213, 85,
	245, 246, 57, 56, 247, 55, 54, 53, 52, 30,
	29, 248, 28, 249, 27, 25, 250, 251, 252, 24,
	23, 255, 256, 33, 26, 258, 31, 86, 32, 263,
	19, 18, 20, 17, 89, 3, 0, 0, 0, 255,
	118, 9, 0, 0, 282, 277, 279, 0, 0, 0,
	289, 0, 291, 287, 0, 0, 295, 139, 140, 141,
	142, 143, 144, 269, 0, 0, 272, 0, 298, 0,
	300, 0, 302, 299, 0, 301, 0, 305, 306, 304,
	307, 257, 0, 259, 0, 0, 0, 264, 0, 0,
	313, 0, 315, 0, 103, 0, 0, 112, 0, 0,
	0, 322, 285, 0, 0, 0, 0, 0, 0, 132,
	292, 0, 0, 294, 328, 0, 0, 0, 331, 0,
	333, 297, 0, 335, 0, 0, 0, 0, 0, 198,
	0, 0, 0, 0, 0, 204, 0, 14, 0, 118,
	0, 309, 0, 0, 0, 0, 132, 0, 0, 211,
	0, 0, 215, 216, 217, 218, 321, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 325, 0,
	326, 0, 15, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 337, 103, 0, 194, 196, 0, 199, 200,
	201, 202, 203, 0, 132, 0, 103, 107, 103, 0,
	107, 0, 103, 0, 0, 0, 0, 16, 11, 105,
	106, 101, 107, 0, 219, 221, 11, 105, 106, 227,
	0, 0, 120, 271, 0, 92, 11, 105, 106, 101,
	0, 0, 108, 37, 38, 108, 0, 0, 0, 0,
	293, 37, 38, 0, 0, 0, 0, 108, 0, 107,
	0, 37, 38, 39, 0, 40, 0, 36, 35, 303,
	0, 39, 0, 40, 0, 36, 35, 109, 308, 0,
	109, 39, 0, 40, 0, 36, 35, 0, 267, 316,
	270, 0, 109, 0, 108, 0, 107, 0, 107, 107,
	323, 107, 107, 107, 107, 107, 0, 107, 0, 107,
	0, 107, 0, 0, 0, 107, 0, 0, 0, 0,
	336, 0, 338, 0, 22, 0, 0, 107, 107, 109,
	0, 108, 107, 108, 108, 0, 108, 108, 108, 108,
	108, 0, 108, 0, 108, 0, 108, 0, 0, 0,
	108, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 108, 108, 0, 0, 109, 108, 109, 109,
	0, 109, 109, 109, 109, 109, 0, 109, 0, 109,
	0, 109, 0, 0, 104, 109, 0, 115, 0, 0,
	0, 107, 0, 107, 0, 0, 0, 109, 109, 115,
	0, 0, 109, 0, 0, 0, 0, 11, 12, 13,
	50, 0, 0, 0, 42, 236, 49, 239, 238, 0,
	43, 44, 0, 0, 0, 51, 108, 0, 108, 0,
	0, 0, 37, 38, 0, 0, 115, 45, 46, 47,
	48, 0, 191, 190, 0, 0, 0, 0, 0, 0,
	0, 0, 39, 0, 40, 0, 36, 35, 0, 0,
	0, 109, 0, 109, 0, 11, 105, 106, 0, 164,
	0, 0, 0, 104, 0, 0, 197, 0, 0, 0,
	0, 0, 0, 0, 115, 0, 104, 0, 104, 0,
	37, 38, 104, 11, 12, 13, 50, 0, 0, 0,
	42, 330, 49, 0, 220, 0, 43, 44, 0, 0,
	39, 51, 40, 0, 36, 35, 0, 0, 37, 38,
	0, 0, 0, 45, 46, 47, 48, 0, 191, 190,
	0, 0, 0, 0, 0, 0, 0, 0, 39, 0,
	40, 0, 36, 35, 11, 12, 13, 50, 0, 0,
	0, 42, 311, 49, 0, 0, 0, 43, 44, 0,
	0, 0, 51, 0, 0, 0, 0, 0, 268, 37,
	38, 0, 0, 0, 45, 46, 47, 48, 0, 191,
	190, 0, 0, 0, 0, 0, 0, 0, 0, 39,
	0, 40, 0, 36, 35, 11, 12, 13, 50, 0,
	0, 0, 42, 290, 49, 0, 0, 0, 43, 44,
	0, 0, 0, 51, 0, 0, 0, 0, 0, 0,
	37, 38, 0, 0, 0, 45, 46, 47, 48, 0,
	191, 190, 0, 0, 0, 0, 0, 0, 0, 0,
	39, 0, 40, 0, 36, 35, 11, 12, 13, 50,
	0, 0, 0, 42, 288, 49, 0, 0, 0, 43,
	44, 0, 0, 0, 51, 0, 0, 0, 0, 0,
	0, 37, 38, 0, 0, 0, 45, 46, 47, 48,
	0, 191, 190, 0, 0, 0, 0, 0, 0, 0,
	0, 39, 0, 40, 0, 36, 35, 11, 12, 13,
	50, 0, 0, 0, 42, 212, 49, 0, 0, 0,
	43, 44, 0, 0, 0, 51, 0, 0, 0, 0,
	0, 0, 37, 38, 0, 0, 0, 45, 46, 47,
	48, 0, 191, 190, 0, 0, 11, 12, 13, 50,
	0, 0, 39, 42, 40, 49, 36, 35, 0, 43,
	44, 0, 0, 0, 51, 0, 0, 0, 0, 0,
	0, 37, 38, 0, 0, 0, 45, 46, 47, 48,
	0, 191, 190, 0, 0, 0, 0, 0, 0, 0,
	0, 39, 0, 40, 0, 36, 35, 11, 12, 13,
	50, 0, 0, 120, 42, 0, 49, 0, 0, 0,
	43, 44, 0, 0, 0, 51, 0, 0, 0, 0,
	0, 0, 37, 38, 0, 0, 117, 45, 46, 47,
	48, 0, 0, 0, 0, 0, 11, 12, 13, 50,
	0, 0, 39, 42, 40, 49, 36, 35, 0, 43,
	44, 0, 0, 0, 51, 0, 0, 0, 0, 0,
	0, 37, 38, 0, 0, 0, 45, 46, 47, 48,
	0, 7, 0, 87, 0, 0, 0, 0, 0, 0,
	0, 39, 0, 40, 0, 36, 35, 11, 12, 13,
	50, 0, 0, 120, 42, 0, 49, 0, 0, 0,
	43, 44, 0, 0, 0, 51, 0, 0, 0, 0,
	0, 0, 37, 38, 0, 0, 0, 45, 46, 47,
	48, 0, 0, 0, 0, 0, 11, 12, 13, 50,
	0, 0, 39, 42, 40, 49, 36, 35, 0, 43,
	44, 0, 0, 0, 51, 0, 0, 0, 0, 0,
	0, 37, 38, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 0, 0, 92, 11, 105, 106, 101, 0,
	0, 39, 0, 40, 99, 36, 35, 100, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	37, 38, 92, 11, 105, 106, 101, 93, 94, 95,
	0, 0, 0, 0, 0, 98, 96, 97, 0, 0,
	39, 0, 40, 0, 36, 35, 0, 0, 37, 38,
	11, 105, 106, 0, 0, 93, 94, 95, 11, 195,
	106, 0, 0, 98, 96, 97, 0, 0, 39, 0,
	40, 0, 36, 35, 0, 37, 38, 0, 0, 0,
	0, 0, 0, 37, 38, 7, 11, 228, 106, 0,
	0, 0, 11, 105, 106, 39, 0, 40, 0, 36,
	35, 0, 0, 39, 0, 40, 0, 36, 35, 0,
	0, 37, 38, 0, 0, 0, 0, 37, 38, 131,
	105, 106, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 39, 0, 40, 0, 36, 35, 39, 0, 40,
	0, 36, 35, 0, 37, 38, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 39, 0, 40, 0, 36, 35,
}
var RubyPact = []int{

	-45, -6, -1000, -48, -1000, -1000, 1111, -3, -1000, -12,
	-3, -1000, 68, -3, -3, -3, -3, -1000, -1000, -1000,
	-1000, -1000, -3, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 12, 149, -1000, -1000, -3,
	-3, 9, -3, -3, -3, -3, -3, -3, -3, -3,
	1021, 36, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 1150,
	166, -1000, 1205, 982, 54, 51, 49, 45, 70, -1000,
	-1000, 142, -1000, -1000, 1274, -1000, -1, 165, 153, 153,
	1111, 1111, 1111, 1111, 1111, 145, -1000, -1000, -1000, -12,
	-1000, -1000, -3, -3, -3, -3, -3, -3, -3, -3,
	-3, -3, -3, -12, -1000, 68, -1000, -1000, -1000, -1000,
	-3, 660, 24, 47, -1000, -3, -3, -3, -1000, -1000,
	-3, -3, -3, -3, -3, -3, -3, -1000, -1000, -3,
	-3, -1000, -12, 13, 154, -3, -3, -1000, -3, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 1150, 931,
	1213, 1247, 1111, 1213, 1213, 1213, 1213, 1213, 1111, 1247,
	134, 431, 1072, 413, -1000, -3, -1000, 1178, -1000, 133,
	41, 128, 1111, 892, -19, 1111, 1111, 1111, 1111, 1247,
	1213, 15, 14, -1000, 1241, -1000, 139, 30, 34, 602,
	-1000, -1000, -1000, 91, -1000, -1000, -12, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -3, -3, -1000, 131, -3,
	-3, -1000, -1000, -3, 164, -1000, -1000, -1000, -1000, -12,
	-1000, -1000, -3, -1000, -1000, -3, -3, -3, -4, -1000,
	-3, -3, -1000, -1000, -3, -1000, -1000, 76, -3, -1000,
	-1000, 31, 119, 421, -1000, 1213, 1072, 33, -18, 168,
	125, 122, 175, -3, 20, 110, 164, 841, 153, 790,
	-1000, -3, -1000, 1111, 931, -3, -1000, -12, -1000, -1000,
	-1000, -1000, -1000, 32, -1000, -1000, -1000, 13, -1000, 13,
	-1000, -3, 1111, -1000, -1000, 931, -3, -3, -1000, -3,
	-1000, 1111, 931, -1000, 931, 151, -1000, 739, -32, 13,
	-33, 13, 1111, -1000, 18, 159, 100, 28, -1000, 931,
	-3, -1000, -1000, 1111, -1000, 156, -1000, -1000, -1000, -1000,
	-1000, 931, 173, -3, -5, 688, 931, -3, 170, -3,
	-1000, 140, -3, 1111, -1000, 1111, 105, 931, 102, -1000,
	-1000,
}
var RubyPgo = []int{

	0, 186, 245, 243, 242, 29, 241, 240, 238, 417,
	236, 234, 233, 244, 382, 524, 230, 7, 229, 347,
	225, 1, 224, 222, 220, 219, 218, 217, 216, 215,
	213, 212, 57, 209, 5, 208, 205, 204, 203, 202,
	196, 30, 6, 194, 192, 191, 189, 188, 0, 4,
	184, 3,
}
var RubyR1 = []int{

	0, 37, 37, 37, 37, 2, 2, 2, 2, 32,
	32, 32, 32, 49, 49, 50, 50, 48, 48, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	17, 17, 17, 17, 17, 17, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 34,
	34, 43, 43, 41, 41, 41, 41, 41, 46, 46,
	46, 46, 45, 45, 45, 45, 45, 51, 51, 51,
	16, 38, 38, 18, 18, 20, 21, 21, 47, 47,
	11, 11, 11, 11, 11, 22, 23, 24, 25, 26,
	26, 26, 26, 27, 28, 29, 30, 31, 3, 6,
	7, 7, 4, 4, 39, 39, 39, 39, 44, 44,
	44, 9, 9, 19, 19, 14, 14, 5, 5, 35,
	42, 42, 42, 10, 10, 10, 10, 10, 36, 36,
	36, 36, 36, 33, 33, 33, 33, 33, 33, 33,
	8, 12, 40, 40, 40, 40,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 1, 3, 3, 0,
	2, 2, 2, 0, 2, 1, 2, 0, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	4, 4, 4, 6, 3, 3, 5, 3, 5, 5,
	1, 1, 5, 0, 1, 1, 1, 5, 1, 1,
	5, 5, 1, 1, 5, 5, 5, 0, 2, 2,
	9, 1, 5, 7, 11, 7, 1, 4, 1, 4,
	5, 5, 5, 5, 5, 3, 3, 3, 3, 5,
	5, 5, 5, 5, 5, 5, 5, 5, 1, 1,
	5, 5, 9, 9, 0, 5, 10, 11, 4, 9,
	10, 2, 2, 2, 2, 3, 3, 3, 7, 3,
	0, 1, 5, 5, 6, 5, 5, 5, 0, 5,
	3, 4, 2, 0, 1, 1, 1, 2, 2, 2,
	3, 5, 0, 4, 7, 10,
}
var RubyChk = []int{

	-1000, -37, 57, -2, 41, 42, -48, 40, 57, -13,
	-1, 5, 6, 7, -19, -14, -9, -3, -6, -7,
	-4, -21, -15, -16, -18, -20, -11, -22, -23, -24,
	-25, -10, -8, -12, -17, 55, 54, 30, 31, 50,
	52, -47, 12, 18, 19, 35, 36, 37, 38, 14,
	8, 23, -26, -27, -28, -29, -30, -31, -48, -48,
	44, -48, 8, -48, -48, -48, -48, -48, -48, 6,
	7, 55, 6, 7, -48, -48, 43, -48, -48, -48,
	-48, -48, -48, -48, -48, -33, -1, 42, -48, -13,
	41, -34, 4, 37, 38, 39, 46, 47, 45, 14,
	17, 8, -45, -13, -15, 6, 7, -19, -14, -9,
	6, -48, -13, -41, -17, -15, -46, 34, -1, -5,
	11, 34, 34, 34, 34, 37, 14, 6, 7, -43,
	-41, 5, -13, -49, 43, 6, -21, 7, -21, -1,
	-1, -1, -1, -1, -1, 42, -48, 9, -48, -32,
	-48, -48, -48, -48, -48, -48, -48, -48, -48, -48,
	-48, -48, -48, -48, 9, -41, 9, -48, 9, -48,
	-48, -48, -48, -32, -48, -48, -48, -48, -48, -48,
	-48, -48, -48, 41, -48, 7, -48, -48, -48, -32,
	41, 40, -1, -40, -13, 6, -13, -15, -1, -13,
	-13, -13, -13, -13, -1, -41, 10, -34, -48, 10,
	10, -1, 13, -35, 45, -1, -1, -1, -1, -13,
	-15, -13, 10, 51, 51, -39, -44, -13, 6, -38,
	-42, 8, 6, 41, 32, 41, 13, -36, 16, 15,
	13, -49, -48, -48, 9, -48, -48, -48, -42, -48,
	-48, -48, -48, 43, -51, -48, -48, -32, -48, -32,
	13, 16, 15, -48, -32, 24, 9, -13, -15, -5,
	-13, -1, -5, -50, 41, 45, 5, -49, 10, -49,
	10, 4, -48, 40, 41, -32, 10, -42, 13, -21,
	13, -48, -32, -1, -32, -48, 41, -32, -48, -49,
	-48, -49, -48, -1, -51, -48, -48, -48, -1, -32,
	7, 13, 53, -48, 53, -48, -1, 13, 6, 9,
	41, -32, -48, -1, 6, -32, -32, 4, -48, 43,
	13, -48, 4, -48, 6, -48, -1, -32, -1, 10,
	10,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 6, 0, 17, 4, 17,
	17, 19, -2, -2, -2, -2, -2, 25, 26, 27,
	28, 29, -2, 38, 39, 40, 41, 42, 43, 44,
	45, 46, 47, 48, 49, 0, 0, 108, 109, 17,
	17, 0, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 0, 30, 31, 32, 33, 34, 35, 18, 7,
	0, 8, 17, 0, 0, 0, 0, 0, 0, 123,
	124, 0, 121, 122, 63, 13, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 144, 145, 146, -2,
	9, 54, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, -2, -2, 73, -2, -2, 22, 23, 24,
	-2, 63, -2, 17, 65, -2, -2, 17, 68, 69,
	-2, 17, 17, 17, 17, 17, 17, 125, 126, 17,
	17, 19, -2, 17, 0, 17, 17, -2, 17, 95,
	96, 97, 98, 9, 147, 148, 149, 150, 0, 152,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 63,
	0, 0, 0, 0, 50, 17, 51, 0, 52, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 14, 114, -2, 130, 0, 0, 0,
	10, 11, 12, 13, 58, 20, -2, 101, 103, 104,
	105, 106, 107, 135, 137, 17, 17, 56, 0, 17,
	17, 90, 127, 17, 130, 91, 92, 93, 94, -2,
	102, 136, 17, 110, 111, 17, 17, 17, 20, 77,
	-2, 17, 131, 9, 17, 9, 133, 0, 17, 9,
	151, 0, 0, 0, 53, 0, 0, 0, 17, 0,
	13, 13, 0, 17, 9, 0, 130, 0, 0, 0,
	134, 17, 9, 0, 142, 9, 59, -2, 75, 76,
	67, 70, 71, 9, 15, 129, 62, 17, 13, 17,
	13, 17, 0, 78, 79, 77, 17, 17, 83, 17,
	85, 0, 140, 9, 153, 0, 16, 0, 0, 17,
	0, 17, 0, 118, 0, 0, 0, 0, 9, 141,
	17, 128, 112, 0, 113, 0, 115, 80, 132, 82,
	9, 139, 9, 17, 0, 0, 154, 17, 0, 17,
	84, 0, 17, 0, 9, 0, 119, 155, 116, 120,
	117,
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
		//line parser.y:170
		{
			RubyVAL.genericValue = nil
		}
	case 7:
		//line parser.y:171
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 8:
		//line parser.y:172
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 9:
		//line parser.y:175
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 10:
		//line parser.y:177
		{
		}
	case 11:
		//line parser.y:179
		{
		}
	case 12:
		//line parser.y:181
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
			}
		}
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
		//line parser.y:199
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: []ast.Node{},
			}
		}
	case 51:
		//line parser.y:206
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 52:
		//line parser.y:213
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 53:
		//line parser.y:220
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 54:
		//line parser.y:227
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 55:
		//line parser.y:234
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 56:
		//line parser.y:241
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 57:
		//line parser.y:251
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 58:
		//line parser.y:258
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 59:
		//line parser.y:267
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 60:
		//line parser.y:269
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 61:
		//line parser.y:272
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 62:
		//line parser.y:274
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 63:
		//line parser.y:276
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 64:
		//line parser.y:278
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 65:
		//line parser.y:280
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 66:
		//line parser.y:282
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:284
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:288
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 69:
		//line parser.y:290
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 70:
		//line parser.y:292
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 71:
		//line parser.y:294
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 72:
		//line parser.y:297
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:299
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:301
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 75:
		//line parser.y:303
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:305
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 80:
		//line parser.y:315
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 81:
		//line parser.y:324
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 82:
		//line parser.y:326
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 83:
		//line parser.y:329
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 84:
		//line parser.y:336
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-8].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-4].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-8].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 85:
		//line parser.y:346
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 86:
		//line parser.y:355
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 87:
		//line parser.y:361
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 88:
		//line parser.y:369
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 89:
		//line parser.y:373
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 90:
		//line parser.y:378
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 91:
		//line parser.y:385
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 92:
		//line parser.y:392
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 93:
		//line parser.y:399
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 94:
		//line parser.y:406
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 95:
		//line parser.y:413
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 96:
		//line parser.y:414
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 97:
		//line parser.y:415
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 98:
		//line parser.y:416
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 99:
		//line parser.y:419
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 100:
		//line parser.y:427
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 101:
		//line parser.y:435
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 102:
		//line parser.y:443
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 103:
		//line parser.y:452
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 104:
		//line parser.y:461
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 105:
		//line parser.y:470
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 106:
		//line parser.y:479
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 107:
		//line parser.y:488
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 108:
		//line parser.y:496
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 109:
		//line parser.y:497
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 110:
		//line parser.y:499
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 111:
		//line parser.y:501
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 112:
		//line parser.y:504
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 113:
		//line parser.y:512
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 114:
		//line parser.y:520
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 115:
		//line parser.y:522
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 116:
		//line parser.y:529
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 117:
		//line parser.y:536
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 118:
		//line parser.y:544
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 119:
		//line parser.y:551
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 120:
		//line parser.y:558
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 121:
		//line parser.y:566
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 122:
		//line parser.y:568
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 123:
		//line parser.y:571
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 124:
		//line parser.y:573
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 125:
		//line parser.y:576
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 126:
		//line parser.y:578
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 127:
		//line parser.y:581
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 128:
		//line parser.y:583
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 129:
		//line parser.y:591
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 130:
		//line parser.y:593
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 131:
		//line parser.y:595
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 132:
		//line parser.y:597
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 133:
		//line parser.y:600
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 134:
		//line parser.y:607
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 135:
		//line parser.y:615
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 136:
		//line parser.y:622
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 137:
		//line parser.y:629
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 138:
		//line parser.y:636
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 139:
		//line parser.y:638
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 140:
		//line parser.y:645
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 141:
		//line parser.y:652
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 142:
		//line parser.y:659
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 143:
		//line parser.y:666
		{
		}
	case 144:
		//line parser.y:667
		{
			RubyVAL.genericSlice = []ast.Node{RubyS[Rubypt-0].genericValue}
		}
	case 145:
		//line parser.y:668
		{
		}
	case 146:
		//line parser.y:669
		{
		}
	case 147:
		//line parser.y:670
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 148:
		//line parser.y:671
		{
		}
	case 149:
		//line parser.y:672
		{
		}
	case 150:
		//line parser.y:675
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 151:
		//line parser.y:678
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 152:
		//line parser.y:686
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 153:
		//line parser.y:688
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 154:
		//line parser.y:690
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				},
			})
		}
	case 155:
		//line parser.y:699
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
