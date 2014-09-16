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

//line parser.y:750

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
	43, 95,
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
	-2, 62,
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
	-2, 74,
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
	43, 95,
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
	10, 66,
	-2, 17,
	-1, 115,
	37, 17,
	-2, 68,
	-1, 116,
	10, 17,
	-2, 57,
	-1, 121,
	45, 17,
	-2, 9,
	-1, 133,
	9, 66,
	10, 66,
	51, 66,
	-2, 17,
	-1, 138,
	43, 95,
	-2, 93,
	-1, 187,
	43, 96,
	-2, 94,
	-1, 198,
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
	-2, 106,
	-1, 213,
	10, 66,
	-2, 17,
	-1, 222,
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
	-2, 107,
	-1, 233,
	10, 17,
	-2, 83,
	-1, 274,
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
	-2, 76,
	-1, 280,
	34, 17,
	-2, 59,
}

const RubyNprod = 163
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1405

var RubyAct = []int{

	162, 235, 6, 34, 21, 233, 91, 260, 58, 131,
	59, 61, 8, 63, 64, 65, 66, 67, 2, 134,
	69, 70, 329, 68, 120, 327, 211, 225, 217, 92,
	11, 105, 106, 101, 7, 7, 4, 5, 7, 283,
	74, 75, 60, 77, 78, 79, 80, 81, 82, 83,
	84, 88, 348, 259, 167, 37, 38, 272, 135, 336,
	332, 150, 76, 111, 7, 185, 114, 227, 226, 71,
	292, 293, 113, 238, 185, 39, 307, 40, 114, 36,
	35, 282, 237, 137, 139, 7, 147, 292, 293, 60,
	149, 239, 90, 151, 152, 153, 154, 155, 156, 157,
	158, 159, 160, 161, 7, 169, 163, 62, 126, 324,
	297, 164, 125, 168, 170, 114, 171, 172, 173, 174,
	127, 166, 176, 177, 178, 179, 180, 181, 182, 124,
	244, 183, 184, 359, 168, 186, 7, 188, 189, 7,
	190, 123, 122, 126, 267, 358, 269, 268, 334, 295,
	273, 211, 92, 11, 105, 106, 101, 309, 248, 211,
	295, 289, 99, 287, 114, 100, 212, 210, 211, 208,
	207, 209, 236, 353, 234, 128, 129, 114, 37, 38,
	323, 194, 138, 175, 187, 93, 94, 95, 10, 72,
	73, 342, 340, 98, 96, 97, 236, 254, 39, 136,
	40, 110, 36, 35, 285, 351, 191, 345, 246, 247,
	290, 281, 249, 250, 251, 245, 41, 252, 116, 102,
	229, 130, 253, 195, 228, 232, 255, 1, 241, 256,
	257, 258, 86, 216, 261, 262, 85, 263, 57, 265,
	56, 55, 54, 270, 53, 119, 52, 30, 29, 28,
	27, 25, 24, 23, 284, 33, 26, 31, 32, 19,
	291, 18, 140, 141, 142, 143, 144, 145, 296, 301,
	299, 20, 276, 305, 17, 279, 286, 288, 3, 0,
	0, 306, 0, 0, 0, 0, 0, 310, 0, 312,
	0, 314, 0, 0, 0, 0, 317, 318, 319, 264,
	320, 266, 316, 0, 0, 271, 89, 311, 0, 313,
	326, 0, 328, 9, 330, 0, 0, 0, 0, 333,
	0, 0, 294, 0, 338, 339, 0, 0, 0, 0,
	0, 302, 0, 0, 304, 200, 0, 0, 0, 0,
	0, 206, 347, 308, 0, 119, 350, 0, 0, 352,
	0, 0, 354, 0, 0, 0, 214, 0, 0, 218,
	219, 220, 221, 0, 0, 322, 103, 0, 0, 112,
	14, 0, 0, 0, 0, 0, 0, 11, 12, 13,
	50, 133, 0, 337, 42, 240, 49, 243, 242, 0,
	43, 44, 0, 0, 0, 51, 0, 0, 343, 0,
	344, 0, 37, 38, 0, 0, 0, 45, 46, 47,
	48, 0, 193, 192, 0, 356, 0, 0, 133, 0,
	0, 0, 39, 0, 40, 0, 36, 35, 0, 0,
	107, 0, 278, 107, 15, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 107, 0, 0, 0, 0,
	0, 0, 303, 11, 105, 106, 103, 165, 196, 198,
	0, 201, 202, 203, 204, 205, 0, 133, 0, 103,
	0, 103, 0, 315, 0, 103, 0, 0, 37, 38,
	213, 0, 107, 321, 0, 0, 0, 0, 222, 224,
	0, 0, 0, 230, 108, 0, 331, 108, 39, 0,
	40, 335, 36, 35, 0, 0, 0, 16, 0, 108,
	341, 0, 0, 0, 0, 11, 105, 106, 101, 0,
	107, 346, 107, 107, 0, 107, 107, 107, 107, 107,
	0, 107, 0, 107, 355, 107, 357, 0, 22, 107,
	37, 38, 0, 0, 107, 0, 108, 0, 0, 0,
	0, 0, 107, 107, 274, 0, 277, 107, 103, 0,
	39, 0, 40, 0, 36, 35, 0, 109, 0, 0,
	109, 0, 0, 0, 0, 0, 0, 0, 0, 11,
	105, 106, 109, 0, 108, 121, 108, 108, 0, 108,
	108, 108, 108, 108, 0, 108, 0, 108, 104, 108,
	0, 115, 0, 108, 37, 38, 0, 0, 108, 0,
	0, 0, 0, 115, 0, 0, 108, 108, 107, 109,
	107, 108, 107, 0, 39, 0, 40, 0, 36, 35,
	11, 197, 106, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	115, 0, 0, 0, 0, 37, 38, 109, 0, 109,
	109, 0, 109, 109, 109, 109, 109, 0, 109, 0,
	109, 0, 109, 0, 0, 39, 109, 40, 0, 36,
	35, 109, 108, 0, 108, 0, 108, 0, 104, 109,
	109, 199, 0, 0, 109, 0, 0, 0, 0, 115,
	0, 104, 0, 104, 0, 0, 0, 104, 0, 0,
	0, 0, 115, 0, 0, 11, 12, 13, 50, 0,
	223, 0, 42, 349, 49, 0, 0, 0, 43, 44,
	0, 0, 0, 51, 0, 0, 0, 0, 0, 0,
	37, 38, 0, 0, 0, 45, 46, 47, 48, 0,
	193, 192, 0, 0, 0, 109, 0, 109, 0, 109,
	39, 0, 40, 0, 36, 35, 11, 12, 13, 50,
	0, 0, 0, 42, 325, 49, 0, 0, 0, 43,
	44, 0, 0, 0, 51, 0, 275, 0, 0, 0,
	104, 37, 38, 0, 0, 0, 45, 46, 47, 48,
	0, 193, 192, 0, 0, 0, 0, 0, 0, 0,
	0, 39, 0, 40, 0, 36, 35, 11, 12, 13,
	50, 0, 0, 0, 42, 300, 49, 0, 0, 0,
	43, 44, 0, 0, 0, 51, 0, 0, 0, 0,
	0, 0, 37, 38, 0, 0, 0, 45, 46, 47,
	48, 0, 193, 192, 0, 0, 0, 0, 0, 0,
	0, 0, 39, 0, 40, 0, 36, 35, 11, 12,
	13, 50, 0, 0, 0, 42, 298, 49, 0, 0,
	0, 43, 44, 0, 0, 0, 51, 0, 0, 0,
	0, 0, 0, 37, 38, 0, 0, 0, 45, 46,
	47, 48, 0, 193, 192, 0, 0, 0, 0, 0,
	0, 0, 0, 39, 0, 40, 0, 36, 35, 11,
	12, 13, 50, 0, 0, 0, 42, 215, 49, 0,
	0, 0, 43, 44, 0, 0, 0, 51, 0, 0,
	0, 0, 0, 0, 37, 38, 0, 0, 0, 45,
	46, 47, 48, 0, 193, 192, 0, 0, 11, 12,
	13, 50, 148, 0, 39, 42, 40, 49, 36, 35,
	0, 43, 44, 0, 0, 0, 51, 0, 0, 0,
	0, 0, 0, 37, 38, 0, 0, 0, 45, 46,
	47, 48, 0, 7, 0, 146, 0, 11, 12, 13,
	50, 0, 0, 39, 42, 40, 49, 36, 35, 0,
	43, 44, 0, 0, 0, 51, 0, 0, 0, 0,
	0, 0, 37, 38, 0, 0, 0, 45, 46, 47,
	48, 0, 193, 192, 0, 0, 0, 0, 0, 0,
	0, 0, 39, 0, 40, 0, 36, 35, 11, 12,
	13, 50, 0, 0, 121, 42, 0, 49, 0, 0,
	0, 43, 44, 0, 0, 0, 51, 0, 0, 0,
	0, 0, 0, 37, 38, 0, 0, 118, 45, 46,
	47, 48, 0, 0, 0, 0, 0, 11, 12, 13,
	50, 0, 0, 117, 42, 40, 49, 36, 35, 0,
	43, 44, 0, 0, 0, 51, 0, 0, 0, 0,
	0, 0, 37, 38, 0, 0, 0, 45, 46, 47,
	48, 0, 7, 0, 87, 0, 0, 0, 11, 231,
	106, 0, 39, 0, 40, 0, 36, 35, 11, 12,
	13, 50, 0, 0, 121, 42, 0, 49, 0, 0,
	0, 43, 44, 37, 38, 0, 51, 0, 0, 0,
	0, 0, 0, 37, 38, 0, 0, 0, 45, 46,
	47, 48, 0, 39, 0, 40, 0, 36, 35, 0,
	0, 0, 0, 39, 0, 40, 0, 36, 35, 11,
	12, 13, 50, 0, 0, 121, 42, 0, 49, 0,
	0, 0, 43, 44, 0, 0, 0, 51, 0, 0,
	0, 0, 0, 0, 37, 38, 0, 0, 0, 45,
	46, 47, 48, 0, 0, 0, 0, 0, 11, 12,
	13, 50, 0, 0, 117, 42, 40, 49, 36, 35,
	0, 43, 44, 0, 0, 0, 51, 0, 0, 0,
	0, 0, 0, 37, 38, 0, 0, 0, 45, 46,
	47, 48, 92, 11, 105, 106, 101, 0, 0, 0,
	0, 0, 0, 39, 0, 40, 0, 36, 35, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 37, 38,
	92, 11, 105, 106, 101, 93, 94, 95, 0, 0,
	0, 0, 0, 98, 96, 97, 0, 0, 39, 280,
	40, 0, 36, 35, 0, 0, 37, 38, 11, 105,
	106, 0, 0, 93, 94, 95, 11, 105, 106, 0,
	0, 98, 96, 97, 0, 0, 39, 0, 40, 0,
	36, 35, 0, 37, 38, 0, 0, 0, 0, 0,
	0, 37, 38, 7, 132, 105, 106, 0, 0, 0,
	0, 0, 0, 39, 0, 40, 0, 36, 35, 0,
	0, 39, 0, 40, 0, 36, 35, 0, 0, 37,
	38, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 39,
	0, 40, 0, 36, 35,
}
var RubyPact = []int{

	-39, -5, -1000, -45, -1000, -1000, 1223, 64, -1000, -2,
	64, -1000, 99, 64, 64, 64, 64, -1000, -1000, -1000,
	-1000, -1000, 64, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 14, 183, -1000, -1000, 64,
	64, 19, 64, 64, 64, 64, 64, 64, 64, 64,
	1082, 51, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 148,
	195, -1000, 1313, 1043, 108, 107, 95, 78, 106, -1000,
	-1000, 169, -1000, -1000, 1349, -1000, 15, 193, 175, 175,
	1223, 1223, 1223, 1223, 1223, 953, -1000, -1000, -1000, -2,
	-1000, -1000, 64, 64, 64, 64, 64, 64, 64, 64,
	64, 64, 64, -2, -1000, 99, -1000, -1000, -1000, -1000,
	64, 448, 45, 96, -1000, 64, 64, 64, 64, -1000,
	-1000, 64, 64, 64, 64, 64, 64, 64, -1000, -1000,
	64, 64, -1000, -2, 24, 177, 64, 64, -1000, 64,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 148,
	992, 625, 1321, 1223, 625, 625, 625, 625, 625, 1223,
	1321, 159, 25, 1184, 510, -1000, 64, -1000, 1286, -1000,
	158, 71, 156, 1349, 1223, 914, -17, 1223, 1223, 1223,
	1223, 1321, 625, 17, 16, -1000, 1123, -1000, 166, 41,
	50, 372, -1000, -1000, -1000, 117, -1000, -1000, -2, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 64, 64, -1000,
	149, 64, 64, -2, -1000, -1000, 64, 191, -1000, -1000,
	-1000, -1000, -2, -1000, -1000, 64, -1000, -1000, 64, 64,
	64, 10, -1000, 64, 64, -1000, 64, -1000, 64, -1000,
	-1000, 131, 64, -1000, -1000, 33, 141, 574, -1000, 625,
	1133, 1258, 40, -6, -1000, 199, 153, 151, 206, 64,
	30, 150, 190, 76, 863, 175, 812, -1000, 64, -1000,
	1223, 992, 64, -1000, -2, -1000, -1000, -1000, -1000, -1000,
	64, 35, -1000, -1000, 147, -1000, 24, -1000, 24, -1000,
	64, 1223, -1000, -1000, 992, 64, 64, 64, -1000, 64,
	-1000, 1223, 992, -1000, 992, 173, 75, -1000, 761, 64,
	-28, 24, -31, 24, 1223, -1000, 47, 190, 139, 1223,
	18, -1000, 992, 64, 64, -1000, 186, -1000, 1223, -1000,
	185, -1000, -1000, -1000, -1000, -1000, -1000, 992, 203, 1223,
	-1000, 64, 9, 710, 992, 64, -1000, 201, 64, -1000,
	167, 64, 1223, -1000, 1223, 135, 992, 123, -1000, -1000,
}
var RubyPgo = []int{

	0, 181, 278, 274, 271, 24, 261, 259, 258, 507,
	257, 256, 255, 306, 434, 538, 253, 3, 252, 370,
	251, 4, 1, 250, 249, 248, 247, 246, 244, 242,
	241, 240, 238, 61, 236, 6, 233, 228, 227, 225,
	224, 223, 9, 222, 221, 220, 219, 218, 5, 216,
	0, 19, 211, 7,
}
var RubyR1 = []int{

	0, 38, 38, 38, 38, 2, 2, 2, 2, 33,
	33, 33, 33, 51, 51, 52, 52, 50, 50, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	17, 17, 17, 17, 17, 17, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 35, 35, 44, 44, 42, 42, 42, 42, 42,
	47, 47, 47, 47, 46, 46, 46, 46, 46, 53,
	53, 53, 16, 39, 39, 22, 22, 48, 48, 48,
	18, 18, 20, 21, 21, 49, 49, 11, 11, 11,
	11, 11, 23, 24, 25, 26, 27, 27, 27, 27,
	28, 29, 30, 31, 32, 3, 6, 7, 7, 4,
	4, 40, 40, 40, 40, 45, 45, 45, 9, 9,
	19, 19, 14, 14, 5, 5, 36, 43, 43, 43,
	10, 10, 10, 10, 10, 37, 37, 37, 37, 37,
	34, 34, 34, 34, 34, 34, 34, 8, 12, 41,
	41, 41, 41,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 1, 3, 3, 0,
	2, 2, 2, 0, 2, 1, 2, 0, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	4, 4, 4, 6, 3, 3, 5, 3, 5, 7,
	11, 5, 1, 1, 5, 0, 1, 1, 1, 5,
	1, 1, 5, 5, 1, 1, 5, 5, 5, 0,
	2, 2, 9, 1, 5, 2, 5, 0, 1, 5,
	7, 11, 7, 1, 4, 1, 4, 5, 5, 5,
	5, 5, 3, 3, 3, 3, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 1, 1, 5, 5, 9,
	9, 0, 5, 10, 11, 4, 9, 10, 2, 2,
	2, 2, 3, 3, 3, 7, 3, 0, 1, 5,
	5, 6, 5, 5, 5, 0, 5, 3, 4, 2,
	0, 1, 1, 1, 2, 2, 2, 3, 5, 0,
	4, 7, 10,
}
var RubyChk = []int{

	-1000, -38, 57, -2, 41, 42, -50, 40, 57, -13,
	-1, 5, 6, 7, -19, -14, -9, -3, -6, -7,
	-4, -21, -15, -16, -18, -20, -11, -23, -24, -25,
	-26, -10, -8, -12, -17, 55, 54, 30, 31, 50,
	52, -49, 12, 18, 19, 35, 36, 37, 38, 14,
	8, 23, -27, -28, -29, -30, -31, -32, -50, -50,
	44, -50, 8, -50, -50, -50, -50, -50, -50, 6,
	7, 55, 6, 7, -50, -50, 43, -50, -50, -50,
	-50, -50, -50, -50, -50, -34, -1, 42, -50, -13,
	41, -35, 4, 37, 38, 39, 46, 47, 45, 14,
	17, 8, -46, -13, -15, 6, 7, -19, -14, -9,
	6, -50, -13, -42, -17, -15, -47, 50, 34, -1,
	-5, 11, 34, 34, 34, 34, 37, 14, 6, 7,
	-44, -42, 5, -13, -51, 43, 6, -21, 7, -21,
	-1, -1, -1, -1, -1, -1, 42, -50, 9, -50,
	-33, -50, -50, -50, -50, -50, -50, -50, -50, -50,
	-50, -50, -50, -50, -50, 9, -42, 9, -50, 9,
	-50, -50, -50, -50, -50, -33, -50, -50, -50, -50,
	-50, -50, -50, -50, -50, 41, -50, 7, -50, -50,
	-50, -33, 41, 40, -1, -41, -13, 6, -13, -15,
	-1, -13, -13, -13, -13, -13, -1, -42, 10, -35,
	-50, 10, 10, -13, -1, 13, -36, 45, -1, -1,
	-1, -1, -13, -15, -13, 10, 51, 51, -40, -45,
	-13, 6, -39, -48, 8, -22, 6, 41, 32, 41,
	13, -37, 16, 15, 13, -51, -50, -50, 9, -50,
	-50, -50, -50, -43, 6, -50, -50, -50, -50, 43,
	-53, -50, -50, -50, -33, -50, -33, 13, 16, 15,
	-50, -33, 24, 9, -13, -15, -5, -13, -1, -5,
	51, -52, 41, 45, -50, 5, -51, 10, -51, 10,
	4, -50, 40, 41, -33, 10, -48, 34, 13, -21,
	13, -50, -33, -1, -33, -50, -50, 41, -33, 10,
	-50, -51, -50, -51, -50, -1, -53, -50, -50, -50,
	-50, -1, -33, 7, 34, 13, -50, 53, -50, 53,
	-50, -1, 13, -22, 9, -1, 41, -33, -50, -50,
	6, -1, 6, -33, -33, 4, -1, -50, 43, 13,
	-50, 4, -50, 6, -50, -1, -33, -1, 10, 10,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 6, 0, 17, 4, 17,
	17, 19, -2, -2, -2, -2, -2, 25, 26, 27,
	28, 29, -2, 38, 39, 40, 41, 42, 43, 44,
	45, 46, 47, 48, 49, 0, 0, 115, 116, 17,
	17, 0, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 0, 30, 31, 32, 33, 34, 35, 18, 7,
	0, 8, 17, 0, 0, 0, 0, 0, 0, 130,
	131, 0, 128, 129, 65, 13, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 151, 152, 153, -2,
	9, 54, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, -2, -2, 75, -2, -2, 22, 23, 24,
	-2, 65, -2, 17, 67, -2, -2, 17, 17, 70,
	71, -2, 17, 17, 17, 17, 17, 17, 132, 133,
	17, 17, 19, -2, 17, 0, 17, 17, -2, 17,
	102, 103, 104, 105, 9, 154, 155, 156, 157, 0,
	159, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	65, 0, 0, 0, 0, 50, 17, 51, 0, 52,
	0, 0, 0, 65, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 14, 121, -2, 87, 0,
	0, 0, 10, 11, 12, 13, 58, 20, -2, 108,
	110, 111, 112, 113, 114, 142, 144, 17, 17, 56,
	0, 17, 17, -2, 97, 134, 17, 137, 98, 99,
	100, 101, -2, 109, 143, 17, 117, 118, 17, 17,
	17, 20, 79, -2, 17, 88, 17, 9, 17, 9,
	140, 0, 17, 9, 158, 0, 0, 0, 53, 0,
	0, 0, 0, 17, 138, 0, 13, 13, 0, 17,
	9, 0, 87, 85, 0, 0, 0, 141, 17, 9,
	0, 149, 9, 61, -2, 77, 78, 69, 72, 73,
	-2, 9, 15, 136, 0, 64, 17, 13, 17, 13,
	17, 0, 80, 81, 79, 17, 17, 17, 90, 17,
	92, 0, 147, 9, 160, 0, 0, 16, 0, 17,
	0, 17, 0, 17, 0, 125, 0, 0, 0, 0,
	0, 9, 148, 17, 17, 135, 0, 119, 0, 120,
	0, 122, 82, 89, 84, 86, 9, 146, 9, 0,
	139, 17, 0, 0, 161, 17, 60, 0, 17, 91,
	0, 17, 0, 9, 0, 126, 162, 123, 127, 124,
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
		//line parser.y:155
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:157
		{
			Statements = []ast.Node{}
		}
	case 3:
		//line parser.y:159
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 4:
		//line parser.y:165
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 5:
		//line parser.y:171
		{
			RubyVAL.genericValue = nil
		}
	case 6:
		//line parser.y:172
		{
			RubyVAL.genericValue = nil
		}
	case 7:
		//line parser.y:173
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 8:
		//line parser.y:174
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 9:
		//line parser.y:177
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 10:
		//line parser.y:179
		{
		}
	case 11:
		//line parser.y:181
		{
		}
	case 12:
		//line parser.y:183
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
		//line parser.y:201
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: []ast.Node{},
			}
		}
	case 51:
		//line parser.y:208
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 52:
		//line parser.y:215
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 53:
		//line parser.y:222
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 54:
		//line parser.y:229
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 55:
		//line parser.y:236
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 56:
		//line parser.y:243
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 57:
		//line parser.y:253
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 58:
		//line parser.y:260
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 59:
		//line parser.y:270
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-6].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 60:
		//line parser.y:280
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-10].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-6].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 61:
		//line parser.y:289
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 62:
		//line parser.y:291
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 63:
		//line parser.y:294
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 64:
		//line parser.y:296
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 65:
		//line parser.y:298
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 66:
		//line parser.y:300
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:302
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:304
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 69:
		//line parser.y:306
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 70:
		//line parser.y:310
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 71:
		//line parser.y:312
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 72:
		//line parser.y:314
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:316
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:319
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 75:
		//line parser.y:321
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:323
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:325
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 78:
		//line parser.y:327
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 82:
		//line parser.y:337
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 83:
		//line parser.y:346
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 84:
		//line parser.y:348
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 85:
		//line parser.y:351
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-1].genericValue.(ast.BareReference)}
		}
	case 86:
		//line parser.y:353
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 87:
		//line parser.y:355
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 88:
		//line parser.y:357
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 89:
		//line parser.y:361
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 90:
		//line parser.y:366
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 91:
		//line parser.y:373
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-8].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-4].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-8].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 92:
		//line parser.y:383
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 93:
		//line parser.y:392
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 94:
		//line parser.y:398
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 95:
		//line parser.y:406
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 96:
		//line parser.y:410
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 97:
		//line parser.y:415
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 98:
		//line parser.y:422
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 99:
		//line parser.y:429
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 100:
		//line parser.y:436
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 101:
		//line parser.y:443
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 102:
		//line parser.y:450
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 103:
		//line parser.y:451
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 104:
		//line parser.y:452
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 105:
		//line parser.y:453
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 106:
		//line parser.y:456
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 107:
		//line parser.y:464
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 108:
		//line parser.y:472
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 109:
		//line parser.y:480
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 110:
		//line parser.y:489
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 111:
		//line parser.y:498
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 112:
		//line parser.y:507
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 113:
		//line parser.y:516
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 114:
		//line parser.y:525
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 115:
		//line parser.y:533
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 116:
		//line parser.y:534
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 117:
		//line parser.y:536
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 118:
		//line parser.y:538
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 119:
		//line parser.y:541
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 120:
		//line parser.y:549
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 121:
		//line parser.y:557
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 122:
		//line parser.y:559
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 123:
		//line parser.y:566
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 124:
		//line parser.y:573
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 125:
		//line parser.y:581
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 126:
		//line parser.y:588
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 127:
		//line parser.y:595
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 128:
		//line parser.y:603
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 129:
		//line parser.y:605
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 130:
		//line parser.y:608
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 131:
		//line parser.y:610
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 132:
		//line parser.y:613
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 133:
		//line parser.y:615
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 134:
		//line parser.y:618
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 135:
		//line parser.y:620
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 136:
		//line parser.y:628
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 137:
		//line parser.y:630
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 138:
		//line parser.y:632
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 139:
		//line parser.y:634
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 140:
		//line parser.y:637
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 141:
		//line parser.y:644
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 142:
		//line parser.y:652
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 143:
		//line parser.y:659
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 144:
		//line parser.y:666
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 145:
		//line parser.y:673
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 146:
		//line parser.y:675
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 147:
		//line parser.y:682
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 148:
		//line parser.y:689
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 149:
		//line parser.y:696
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 150:
		//line parser.y:703
		{
		}
	case 151:
		//line parser.y:704
		{
			RubyVAL.genericSlice = []ast.Node{RubyS[Rubypt-0].genericValue}
		}
	case 152:
		//line parser.y:705
		{
		}
	case 153:
		//line parser.y:706
		{
		}
	case 154:
		//line parser.y:707
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 155:
		//line parser.y:708
		{
		}
	case 156:
		//line parser.y:709
		{
		}
	case 157:
		//line parser.y:712
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 158:
		//line parser.y:715
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 159:
		//line parser.y:723
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 160:
		//line parser.y:725
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 161:
		//line parser.y:727
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				},
			})
		}
	case 162:
		//line parser.y:736
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
