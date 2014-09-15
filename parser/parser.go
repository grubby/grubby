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

//line parser.y:733

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
	43, 90,
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
	43, 90,
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
	43, 90,
	-2, 88,
	-1, 187,
	43, 91,
	-2, 89,
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
	-2, 101,
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
	-2, 102,
	-1, 233,
	10, 17,
	-2, 83,
	-1, 271,
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
	-1, 277,
	34, 17,
	-2, 59,
}

const RubyNprod = 158
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1404

var RubyAct = []int{

	162, 21, 6, 258, 134, 120, 233, 91, 58, 8,
	59, 61, 2, 63, 64, 65, 66, 67, 69, 70,
	321, 319, 217, 68, 131, 211, 92, 11, 105, 106,
	101, 7, 338, 7, 225, 257, 280, 60, 167, 135,
	74, 75, 76, 77, 78, 79, 80, 81, 82, 83,
	84, 88, 37, 38, 7, 4, 5, 324, 34, 7,
	185, 150, 269, 111, 327, 7, 227, 71, 302, 7,
	288, 289, 39, 60, 40, 226, 36, 35, 279, 185,
	137, 139, 237, 238, 288, 289, 147, 113, 90, 169,
	149, 236, 127, 151, 152, 153, 154, 155, 156, 157,
	158, 159, 160, 161, 126, 317, 163, 62, 125, 124,
	123, 164, 349, 168, 170, 126, 171, 172, 173, 174,
	7, 114, 176, 177, 178, 179, 180, 181, 182, 122,
	243, 183, 184, 114, 168, 186, 166, 188, 189, 7,
	190, 264, 348, 266, 265, 326, 291, 270, 211, 247,
	211, 291, 285, 92, 11, 105, 106, 101, 283, 343,
	212, 211, 208, 99, 128, 129, 100, 210, 316, 235,
	114, 234, 209, 72, 73, 138, 187, 332, 325, 37,
	38, 235, 136, 175, 110, 207, 93, 94, 95, 281,
	194, 341, 335, 286, 98, 96, 97, 10, 278, 39,
	244, 40, 41, 36, 35, 116, 191, 102, 245, 246,
	229, 130, 248, 249, 250, 195, 228, 251, 232, 114,
	1, 240, 216, 85, 252, 57, 253, 56, 55, 254,
	255, 256, 114, 54, 259, 260, 53, 52, 262, 30,
	29, 86, 267, 28, 27, 25, 24, 23, 33, 26,
	31, 14, 273, 259, 119, 276, 32, 19, 287, 282,
	284, 18, 20, 17, 294, 3, 296, 292, 0, 0,
	300, 140, 141, 142, 143, 144, 145, 0, 301, 0,
	0, 0, 0, 304, 0, 306, 0, 308, 305, 0,
	307, 0, 311, 312, 310, 313, 0, 0, 261, 22,
	263, 0, 0, 0, 268, 0, 320, 0, 322, 0,
	0, 107, 0, 0, 107, 0, 0, 329, 330, 0,
	290, 0, 0, 0, 0, 0, 107, 89, 297, 0,
	0, 299, 337, 0, 9, 0, 340, 0, 0, 342,
	303, 0, 344, 0, 200, 0, 0, 11, 105, 106,
	206, 0, 0, 121, 119, 0, 0, 0, 0, 104,
	315, 0, 115, 107, 0, 214, 0, 0, 218, 219,
	220, 221, 37, 38, 115, 0, 328, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 103, 0, 333,
	112, 334, 39, 0, 40, 0, 36, 35, 11, 197,
	106, 107, 133, 107, 107, 346, 107, 107, 107, 107,
	107, 115, 107, 0, 107, 0, 107, 0, 0, 0,
	107, 0, 0, 37, 38, 107, 0, 0, 0, 0,
	0, 15, 0, 107, 107, 0, 0, 0, 107, 133,
	275, 0, 0, 39, 0, 40, 0, 36, 35, 104,
	0, 0, 199, 0, 0, 0, 0, 0, 298, 0,
	115, 0, 104, 0, 104, 0, 0, 16, 104, 0,
	0, 0, 0, 115, 0, 0, 0, 103, 309, 196,
	198, 223, 201, 202, 203, 204, 205, 314, 133, 0,
	103, 108, 103, 0, 108, 0, 103, 0, 107, 323,
	107, 213, 107, 0, 0, 0, 108, 0, 0, 222,
	224, 331, 0, 0, 230, 0, 0, 0, 0, 0,
	0, 336, 92, 11, 105, 106, 101, 109, 0, 0,
	109, 0, 0, 345, 0, 347, 0, 0, 0, 0,
	0, 0, 109, 108, 0, 0, 272, 0, 37, 38,
	104, 0, 0, 0, 0, 93, 94, 95, 0, 0,
	0, 0, 0, 98, 96, 97, 0, 0, 39, 277,
	40, 0, 36, 35, 271, 0, 274, 0, 103, 109,
	0, 108, 0, 108, 108, 0, 108, 108, 108, 108,
	108, 0, 108, 0, 108, 0, 108, 0, 0, 0,
	108, 0, 0, 0, 0, 108, 0, 0, 0, 0,
	0, 0, 0, 108, 108, 0, 0, 109, 108, 109,
	109, 0, 109, 109, 109, 109, 109, 0, 109, 0,
	109, 0, 109, 0, 0, 0, 109, 0, 0, 0,
	0, 109, 0, 0, 0, 0, 0, 0, 0, 109,
	109, 0, 0, 0, 109, 0, 0, 0, 11, 12,
	13, 50, 0, 0, 0, 42, 239, 49, 242, 241,
	0, 43, 44, 0, 0, 0, 51, 0, 108, 0,
	108, 0, 108, 37, 38, 0, 0, 0, 45, 46,
	47, 48, 0, 193, 192, 0, 0, 0, 0, 0,
	0, 0, 0, 39, 0, 40, 0, 36, 35, 0,
	0, 0, 0, 0, 109, 0, 109, 0, 109, 11,
	12, 13, 50, 0, 0, 0, 42, 339, 49, 0,
	0, 0, 43, 44, 0, 0, 0, 51, 0, 0,
	0, 0, 0, 0, 37, 38, 0, 0, 0, 45,
	46, 47, 48, 0, 193, 192, 0, 0, 0, 0,
	0, 0, 0, 0, 39, 0, 40, 0, 36, 35,
	11, 12, 13, 50, 0, 0, 0, 42, 318, 49,
	0, 0, 0, 43, 44, 0, 0, 0, 51, 0,
	0, 0, 0, 0, 0, 37, 38, 0, 0, 0,
	45, 46, 47, 48, 0, 193, 192, 0, 0, 0,
	0, 0, 0, 0, 0, 39, 0, 40, 0, 36,
	35, 11, 12, 13, 50, 0, 0, 0, 42, 295,
	49, 0, 0, 0, 43, 44, 0, 0, 0, 51,
	0, 0, 0, 0, 0, 0, 37, 38, 0, 0,
	0, 45, 46, 47, 48, 0, 193, 192, 0, 0,
	0, 0, 0, 0, 0, 0, 39, 0, 40, 0,
	36, 35, 11, 12, 13, 50, 0, 0, 0, 42,
	293, 49, 0, 0, 0, 43, 44, 0, 0, 0,
	51, 0, 0, 0, 0, 0, 0, 37, 38, 0,
	0, 0, 45, 46, 47, 48, 0, 193, 192, 0,
	0, 0, 0, 0, 0, 0, 0, 39, 0, 40,
	0, 36, 35, 11, 12, 13, 50, 0, 0, 0,
	42, 215, 49, 0, 0, 0, 43, 44, 0, 0,
	0, 51, 0, 0, 0, 0, 0, 0, 37, 38,
	0, 0, 0, 45, 46, 47, 48, 0, 193, 192,
	0, 0, 11, 12, 13, 50, 148, 0, 39, 42,
	40, 49, 36, 35, 0, 43, 44, 0, 0, 0,
	51, 0, 0, 0, 0, 0, 0, 37, 38, 0,
	0, 0, 45, 46, 47, 48, 0, 7, 0, 146,
	0, 11, 12, 13, 50, 0, 0, 39, 42, 40,
	49, 36, 35, 0, 43, 44, 0, 0, 0, 51,
	0, 0, 0, 0, 0, 0, 37, 38, 0, 0,
	0, 45, 46, 47, 48, 0, 193, 192, 0, 0,
	0, 0, 0, 0, 0, 0, 39, 0, 40, 0,
	36, 35, 11, 12, 13, 50, 0, 0, 121, 42,
	0, 49, 0, 0, 0, 43, 44, 0, 0, 0,
	51, 0, 0, 0, 0, 0, 0, 37, 38, 0,
	0, 118, 45, 46, 47, 48, 0, 0, 0, 0,
	0, 11, 12, 13, 50, 0, 0, 117, 42, 40,
	49, 36, 35, 0, 43, 44, 0, 0, 0, 51,
	0, 0, 0, 0, 0, 0, 37, 38, 0, 0,
	0, 45, 46, 47, 48, 0, 7, 0, 87, 0,
	0, 0, 11, 105, 106, 101, 39, 0, 40, 0,
	36, 35, 11, 12, 13, 50, 0, 0, 121, 42,
	0, 49, 0, 0, 0, 43, 44, 37, 38, 0,
	51, 0, 0, 0, 0, 0, 0, 37, 38, 0,
	0, 0, 45, 46, 47, 48, 0, 39, 0, 40,
	0, 36, 35, 0, 0, 0, 0, 39, 0, 40,
	0, 36, 35, 11, 12, 13, 50, 0, 0, 121,
	42, 0, 49, 0, 0, 0, 43, 44, 0, 0,
	0, 51, 0, 0, 0, 0, 0, 0, 37, 38,
	0, 0, 0, 45, 46, 47, 48, 0, 0, 0,
	0, 0, 11, 12, 13, 50, 0, 0, 117, 42,
	40, 49, 36, 35, 0, 43, 44, 0, 0, 0,
	51, 0, 0, 0, 0, 0, 0, 37, 38, 0,
	0, 0, 45, 46, 47, 48, 92, 11, 105, 106,
	101, 0, 0, 0, 0, 0, 0, 39, 0, 40,
	0, 36, 35, 0, 0, 11, 105, 106, 0, 165,
	0, 0, 37, 38, 11, 105, 106, 0, 0, 93,
	94, 95, 11, 231, 106, 0, 0, 98, 96, 97,
	37, 38, 39, 0, 40, 0, 36, 35, 0, 37,
	38, 0, 0, 0, 0, 0, 0, 37, 38, 7,
	39, 0, 40, 0, 36, 35, 11, 105, 106, 39,
	0, 40, 0, 36, 35, 0, 0, 39, 0, 40,
	0, 36, 35, 132, 105, 106, 0, 0, 0, 0,
	0, 37, 38, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 37, 38,
	0, 39, 0, 40, 0, 36, 35, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 39, 0,
	40, 0, 36, 35,
}
var RubyPact = []int{

	-45, 14, -1000, -48, -1000, -1000, 1227, 25, -1000, -7,
	25, -1000, 99, 25, 25, 25, 25, -1000, -1000, -1000,
	-1000, -1000, 25, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 12, 167, -1000, -1000, 25,
	25, -1, 25, 25, 25, 25, 25, 25, 25, 25,
	1086, 47, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 149,
	178, -1000, 1289, 1047, 95, 76, 75, 74, 78, -1000,
	-1000, 158, -1000, -1000, 1348, -1000, -4, 176, 168, 168,
	1227, 1227, 1227, 1227, 1227, 957, -1000, -1000, -1000, -7,
	-1000, -1000, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, -7, -1000, 99, -1000, -1000, -1000, -1000,
	25, 1280, 29, 80, -1000, 25, 25, 25, 25, -1000,
	-1000, 25, 25, 25, 25, 25, 25, 25, -1000, -1000,
	25, 25, -1000, -7, 19, 169, 25, 25, -1000, 25,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 149,
	996, 393, 1331, 1227, 393, 393, 393, 393, 393, 1227,
	1331, 152, 22, 1188, 1127, -1000, 25, -1000, 1262, -1000,
	151, 67, 150, 1348, 1227, 918, -23, 1227, 1227, 1227,
	1227, 1331, 393, 24, 15, -1000, 1297, -1000, 163, 50,
	42, 653, -1000, -1000, -1000, 117, -1000, -1000, -7, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 25, 25, -1000,
	140, 25, 25, -7, -1000, -1000, 25, 175, -1000, -1000,
	-1000, -1000, -7, -1000, -1000, 25, -1000, -1000, 25, 25,
	25, -8, -1000, 25, 25, -1000, -1000, 25, -1000, -1000,
	128, 25, -1000, -1000, 38, 138, 342, -1000, 393, 1137,
	518, 37, -9, 184, 148, 142, 189, 25, 30, 141,
	175, 867, 168, 816, -1000, 25, -1000, 1227, 996, 25,
	-1000, -7, -1000, -1000, -1000, -1000, -1000, 25, 27, -1000,
	-1000, -1000, 19, -1000, 19, -1000, 25, 1227, -1000, -1000,
	996, 25, 25, -1000, 25, -1000, 1227, 996, -1000, 996,
	161, 71, -1000, 765, -32, 19, -33, 19, 1227, -1000,
	44, 172, 136, 23, -1000, 996, 25, 25, -1000, -1000,
	1227, -1000, 171, -1000, -1000, -1000, -1000, -1000, 996, 188,
	1227, 25, -11, 714, 996, 25, -1000, 187, 25, -1000,
	153, 25, 1227, -1000, 1227, 132, 996, 102, -1000, -1000,
}
var RubyPgo = []int{

	0, 190, 265, 263, 262, 5, 261, 257, 256, 467,
	250, 249, 248, 327, 431, 299, 247, 58, 246, 251,
	245, 1, 244, 243, 240, 239, 237, 236, 233, 228,
	227, 225, 61, 223, 7, 222, 221, 220, 218, 216,
	215, 24, 6, 211, 210, 207, 205, 202, 0, 4,
	198, 3,
}
var RubyR1 = []int{

	0, 37, 37, 37, 37, 2, 2, 2, 2, 32,
	32, 32, 32, 49, 49, 50, 50, 48, 48, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	17, 17, 17, 17, 17, 17, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 34, 34, 43, 43, 41, 41, 41, 41, 41,
	46, 46, 46, 46, 45, 45, 45, 45, 45, 51,
	51, 51, 16, 38, 38, 18, 18, 20, 21, 21,
	47, 47, 11, 11, 11, 11, 11, 22, 23, 24,
	25, 26, 26, 26, 26, 27, 28, 29, 30, 31,
	3, 6, 7, 7, 4, 4, 39, 39, 39, 39,
	44, 44, 44, 9, 9, 19, 19, 14, 14, 5,
	5, 35, 42, 42, 42, 10, 10, 10, 10, 10,
	36, 36, 36, 36, 36, 33, 33, 33, 33, 33,
	33, 33, 8, 12, 40, 40, 40, 40,
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
	2, 2, 9, 1, 5, 7, 11, 7, 1, 4,
	1, 4, 5, 5, 5, 5, 5, 3, 3, 3,
	3, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	1, 1, 5, 5, 9, 9, 0, 5, 10, 11,
	4, 9, 10, 2, 2, 2, 2, 3, 3, 3,
	7, 3, 0, 1, 5, 5, 6, 5, 5, 5,
	0, 5, 3, 4, 2, 0, 1, 1, 1, 2,
	2, 2, 3, 5, 0, 4, 7, 10,
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
	6, -48, -13, -41, -17, -15, -46, 50, 34, -1,
	-5, 11, 34, 34, 34, 34, 37, 14, 6, 7,
	-43, -41, 5, -13, -49, 43, 6, -21, 7, -21,
	-1, -1, -1, -1, -1, -1, 42, -48, 9, -48,
	-32, -48, -48, -48, -48, -48, -48, -48, -48, -48,
	-48, -48, -48, -48, -48, 9, -41, 9, -48, 9,
	-48, -48, -48, -48, -48, -32, -48, -48, -48, -48,
	-48, -48, -48, -48, -48, 41, -48, 7, -48, -48,
	-48, -32, 41, 40, -1, -40, -13, 6, -13, -15,
	-1, -13, -13, -13, -13, -13, -1, -41, 10, -34,
	-48, 10, 10, -13, -1, 13, -35, 45, -1, -1,
	-1, -1, -13, -15, -13, 10, 51, 51, -39, -44,
	-13, 6, -38, -42, 8, 6, 41, 32, 41, 13,
	-36, 16, 15, 13, -49, -48, -48, 9, -48, -48,
	-48, -48, -42, -48, -48, -48, -48, 43, -51, -48,
	-48, -32, -48, -32, 13, 16, 15, -48, -32, 24,
	9, -13, -15, -5, -13, -1, -5, 51, -50, 41,
	45, 5, -49, 10, -49, 10, 4, -48, 40, 41,
	-32, 10, -42, 13, -21, 13, -48, -32, -1, -32,
	-48, -48, 41, -32, -48, -49, -48, -49, -48, -1,
	-51, -48, -48, -48, -1, -32, 7, 34, 13, 53,
	-48, 53, -48, -1, 13, 6, 9, 41, -32, -48,
	-48, -1, 6, -32, -32, 4, -1, -48, 43, 13,
	-48, 4, -48, 6, -48, -1, -32, -1, 10, 10,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 6, 0, 17, 4, 17,
	17, 19, -2, -2, -2, -2, -2, 25, 26, 27,
	28, 29, -2, 38, 39, 40, 41, 42, 43, 44,
	45, 46, 47, 48, 49, 0, 0, 110, 111, 17,
	17, 0, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 0, 30, 31, 32, 33, 34, 35, 18, 7,
	0, 8, 17, 0, 0, 0, 0, 0, 0, 125,
	126, 0, 123, 124, 65, 13, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 146, 147, 148, -2,
	9, 54, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, -2, -2, 75, -2, -2, 22, 23, 24,
	-2, 65, -2, 17, 67, -2, -2, 17, 17, 70,
	71, -2, 17, 17, 17, 17, 17, 17, 127, 128,
	17, 17, 19, -2, 17, 0, 17, 17, -2, 17,
	97, 98, 99, 100, 9, 149, 150, 151, 152, 0,
	154, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	65, 0, 0, 0, 0, 50, 17, 51, 0, 52,
	0, 0, 0, 65, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 14, 116, -2, 132, 0,
	0, 0, 10, 11, 12, 13, 58, 20, -2, 103,
	105, 106, 107, 108, 109, 137, 139, 17, 17, 56,
	0, 17, 17, -2, 92, 129, 17, 132, 93, 94,
	95, 96, -2, 104, 138, 17, 112, 113, 17, 17,
	17, 20, 79, -2, 17, 133, 9, 17, 9, 135,
	0, 17, 9, 153, 0, 0, 0, 53, 0, 0,
	0, 0, 17, 0, 13, 13, 0, 17, 9, 0,
	132, 0, 0, 0, 136, 17, 9, 0, 144, 9,
	61, -2, 77, 78, 69, 72, 73, -2, 9, 15,
	131, 64, 17, 13, 17, 13, 17, 0, 80, 81,
	79, 17, 17, 85, 17, 87, 0, 142, 9, 155,
	0, 0, 16, 0, 0, 17, 0, 17, 0, 120,
	0, 0, 0, 0, 9, 143, 17, 17, 130, 114,
	0, 115, 0, 117, 82, 134, 84, 9, 141, 9,
	0, 17, 0, 0, 156, 17, 60, 0, 17, 86,
	0, 17, 0, 9, 0, 121, 157, 118, 122, 119,
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
		//line parser.y:268
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-6].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 60:
		//line parser.y:278
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-10].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-6].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 61:
		//line parser.y:287
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 62:
		//line parser.y:289
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 63:
		//line parser.y:292
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 64:
		//line parser.y:294
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 65:
		//line parser.y:296
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 66:
		//line parser.y:298
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:300
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 68:
		//line parser.y:302
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 69:
		//line parser.y:304
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 70:
		//line parser.y:308
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 71:
		//line parser.y:310
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 72:
		//line parser.y:312
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 73:
		//line parser.y:314
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:317
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 75:
		//line parser.y:319
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:321
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:323
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 78:
		//line parser.y:325
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 82:
		//line parser.y:335
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 83:
		//line parser.y:344
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 84:
		//line parser.y:346
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 85:
		//line parser.y:349
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 86:
		//line parser.y:356
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-8].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-4].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-8].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 87:
		//line parser.y:366
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 88:
		//line parser.y:375
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 89:
		//line parser.y:381
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 90:
		//line parser.y:389
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 91:
		//line parser.y:393
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 92:
		//line parser.y:398
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 93:
		//line parser.y:405
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 94:
		//line parser.y:412
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 95:
		//line parser.y:419
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 96:
		//line parser.y:426
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 97:
		//line parser.y:433
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 98:
		//line parser.y:434
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 99:
		//line parser.y:435
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 100:
		//line parser.y:436
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 101:
		//line parser.y:439
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 102:
		//line parser.y:447
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 103:
		//line parser.y:455
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 104:
		//line parser.y:463
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 105:
		//line parser.y:472
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 106:
		//line parser.y:481
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 107:
		//line parser.y:490
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 108:
		//line parser.y:499
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 109:
		//line parser.y:508
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 110:
		//line parser.y:516
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 111:
		//line parser.y:517
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 112:
		//line parser.y:519
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 113:
		//line parser.y:521
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 114:
		//line parser.y:524
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 115:
		//line parser.y:532
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 116:
		//line parser.y:540
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 117:
		//line parser.y:542
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 118:
		//line parser.y:549
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 119:
		//line parser.y:556
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 120:
		//line parser.y:564
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 121:
		//line parser.y:571
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 122:
		//line parser.y:578
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 123:
		//line parser.y:586
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 124:
		//line parser.y:588
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 125:
		//line parser.y:591
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 126:
		//line parser.y:593
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 127:
		//line parser.y:596
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 128:
		//line parser.y:598
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 129:
		//line parser.y:601
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 130:
		//line parser.y:603
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 131:
		//line parser.y:611
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 132:
		//line parser.y:613
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 133:
		//line parser.y:615
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 134:
		//line parser.y:617
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 135:
		//line parser.y:620
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 136:
		//line parser.y:627
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 137:
		//line parser.y:635
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 138:
		//line parser.y:642
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 139:
		//line parser.y:649
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-4].genericValue},
			}
		}
	case 140:
		//line parser.y:656
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 141:
		//line parser.y:658
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 142:
		//line parser.y:665
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 143:
		//line parser.y:672
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 144:
		//line parser.y:679
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 145:
		//line parser.y:686
		{
		}
	case 146:
		//line parser.y:687
		{
			RubyVAL.genericSlice = []ast.Node{RubyS[Rubypt-0].genericValue}
		}
	case 147:
		//line parser.y:688
		{
		}
	case 148:
		//line parser.y:689
		{
		}
	case 149:
		//line parser.y:690
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 150:
		//line parser.y:691
		{
		}
	case 151:
		//line parser.y:692
		{
		}
	case 152:
		//line parser.y:695
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 153:
		//line parser.y:698
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 154:
		//line parser.y:706
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 155:
		//line parser.y:708
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 156:
		//line parser.y:710
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				},
			})
		}
	case 157:
		//line parser.y:719
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
