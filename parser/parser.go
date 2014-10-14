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
const SPECIAL_CHAR_REF = 57349
const CAPITAL_REF = 57350
const LPAREN = 57351
const RPAREN = 57352
const COMMA = 57353
const DO = 57354
const DEF = 57355
const END = 57356
const IF = 57357
const ELSE = 57358
const ELSIF = 57359
const UNLESS = 57360
const CLASS = 57361
const MODULE = 57362
const FOR = 57363
const WHILE = 57364
const UNTIL = 57365
const BEGIN = 57366
const RESCUE = 57367
const ENSURE = 57368
const BREAK = 57369
const REDO = 57370
const RETRY = 57371
const RETURN = 57372
const YIELD = 57373
const TRUE = 57374
const FALSE = 57375
const LESSTHAN = 57376
const GREATERTHAN = 57377
const EQUALTO = 57378
const BANG = 57379
const COMPLEMENT = 57380
const POSITIVE = 57381
const NEGATIVE = 57382
const STAR = 57383
const WHITESPACE = 57384
const NEWLINE = 57385
const SEMICOLON = 57386
const COLON = 57387
const DOUBLECOLON = 57388
const DOT = 57389
const PIPE = 57390
const SLASH = 57391
const AMPERSAND = 57392
const QUESTIONMARK = 57393
const CARET = 57394
const LBRACKET = 57395
const RBRACKET = 57396
const LBRACE = 57397
const RBRACE = 57398
const DOLLARSIGN = 57399
const ATSIGN = 57400
const FILE_CONST_REF = 57401
const EOF = 57402

var RubyToknames = []string{
	"OPERATOR",
	"NODE",
	"REF",
	"SPECIAL_CHAR_REF",
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
	"DOUBLECOLON",
	"DOT",
	"PIPE",
	"SLASH",
	"AMPERSAND",
	"QUESTIONMARK",
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

//line parser.y:869

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 26,
	46, 102,
	-2, 20,
	-1, 84,
	10, 75,
	11, 75,
	-2, 163,
	-1, 94,
	46, 102,
	-2, 20,
	-1, 100,
	13, 13,
	15, 13,
	18, 13,
	19, 13,
	20, 13,
	22, 13,
	24, 13,
	30, 13,
	31, 13,
	37, 13,
	38, 13,
	39, 13,
	40, 13,
	44, 13,
	-2, 11,
	-1, 115,
	46, 102,
	-2, 100,
	-1, 213,
	46, 103,
	-2, 101,
	-1, 220,
	10, 75,
	11, 75,
	-2, 163,
}

const RubyNprod = 176
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1581

var RubyAct = []int{

	9, 183, 246, 34, 181, 180, 92, 206, 20, 86,
	136, 85, 165, 287, 24, 25, 65, 26, 66, 2,
	3, 166, 38, 140, 46, 111, 91, 47, 39, 40,
	313, 57, 70, 48, 135, 130, 4, 131, 139, 56,
	55, 60, 61, 114, 116, 208, 42, 43, 44, 45,
	207, 208, 6, 7, 201, 208, 127, 127, 286, 264,
	68, 67, 62, 137, 63, 189, 59, 58, 288, 8,
	184, 138, 70, 145, 146, 147, 148, 69, 150, 151,
	152, 153, 309, 155, 267, 158, 308, 132, 160, 161,
	126, 128, 74, 138, 91, 157, 138, 205, 238, 320,
	194, 195, 158, 118, 83, 185, 105, 104, 175, 176,
	138, 208, 169, 167, 186, 208, 74, 103, 251, 187,
	70, 266, 72, 73, 101, 306, 198, 312, 117, 252,
	24, 93, 65, 94, 84, 71, 89, 99, 70, 82,
	211, 184, 158, 74, 182, 91, 72, 73, 202, 318,
	138, 102, 214, 70, 218, 70, 219, 60, 61, 71,
	223, 88, 208, 82, 226, 198, 224, 293, 206, 198,
	227, 159, 101, 72, 73, 258, 185, 236, 87, 201,
	100, 259, 59, 58, 221, 186, 71, 237, 141, 198,
	82, 198, 163, 243, 198, 107, 108, 302, 284, 277,
	276, 275, 253, 277, 276, 271, 236, 255, 106, 261,
	206, 72, 73, 91, 24, 93, 65, 94, 66, 235,
	115, 158, 263, 203, 71, 204, 241, 213, 82, 138,
	198, 262, 304, 198, 234, 206, 222, 206, 270, 311,
	295, 60, 61, 27, 64, 133, 74, 134, 256, 257,
	198, 198, 281, 109, 240, 198, 110, 113, 239, 112,
	290, 292, 62, 231, 63, 192, 59, 58, 179, 95,
	198, 174, 107, 108, 198, 198, 72, 73, 162, 144,
	198, 75, 76, 77, 198, 106, 90, 36, 210, 71,
	80, 78, 79, 82, 230, 200, 225, 289, 209, 95,
	95, 198, 198, 1, 198, 142, 95, 54, 296, 53,
	52, 198, 297, 98, 316, 198, 95, 95, 95, 95,
	51, 95, 95, 95, 95, 50, 95, 49, 95, 17,
	307, 95, 95, 16, 15, 14, 41, 95, 12, 11,
	22, 21, 10, 98, 98, 95, 19, 23, 13, 18,
	98, 95, 95, 35, 317, 319, 37, 321, 32, 322,
	98, 98, 98, 98, 31, 98, 98, 98, 98, 33,
	98, 30, 98, 0, 0, 98, 98, 0, 0, 0,
	0, 98, 0, 95, 0, 95, 0, 0, 95, 98,
	0, 0, 0, 0, 0, 98, 98, 0, 0, 0,
	0, 0, 0, 95, 0, 0, 0, 95, 0, 24,
	93, 65, 94, 84, 0, 0, 99, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 98, 0, 98,
	0, 0, 98, 0, 0, 0, 60, 61, 0, 0,
	0, 0, 28, 0, 0, 95, 0, 98, 0, 260,
	95, 98, 0, 0, 0, 0, 95, 87, 0, 100,
	0, 59, 58, 0, 95, 95, 0, 0, 96, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 29, 0, 0, 98,
	0, 0, 0, 0, 98, 0, 0, 0, 96, 96,
	98, 0, 0, 0, 95, 96, 0, 0, 98, 98,
	0, 0, 97, 0, 0, 96, 96, 96, 96, 0,
	96, 96, 96, 96, 0, 96, 0, 96, 0, 0,
	96, 96, 0, 0, 0, 0, 96, 0, 0, 0,
	0, 0, 97, 97, 96, 0, 0, 0, 98, 97,
	96, 96, 0, 0, 0, 0, 0, 95, 0, 97,
	97, 97, 97, 0, 97, 97, 97, 97, 0, 97,
	0, 97, 0, 0, 97, 97, 199, 0, 5, 0,
	97, 0, 96, 0, 96, 0, 0, 96, 97, 0,
	0, 0, 0, 0, 97, 97, 125, 0, 0, 0,
	0, 98, 96, 0, 0, 0, 96, 0, 0, 24,
	93, 65, 94, 66, 0, 0, 99, 0, 0, 119,
	120, 121, 122, 123, 124, 0, 97, 0, 97, 0,
	0, 97, 0, 0, 129, 0, 60, 61, 0, 0,
	0, 0, 0, 0, 96, 0, 97, 143, 0, 96,
	97, 0, 0, 149, 0, 96, 0, 62, 154, 100,
	156, 59, 58, 96, 96, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 24, 212, 65, 94, 66, 170,
	171, 172, 173, 0, 0, 0, 177, 178, 97, 0,
	0, 0, 0, 97, 191, 0, 164, 168, 0, 97,
	0, 60, 61, 96, 0, 0, 0, 97, 97, 0,
	0, 188, 208, 190, 74, 0, 0, 0, 0, 215,
	193, 0, 62, 0, 63, 0, 59, 58, 81, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 72, 73, 0, 97, 0, 75,
	76, 77, 0, 0, 0, 0, 96, 71, 80, 78,
	79, 82, 229, 0, 232, 0, 0, 0, 0, 24,
	25, 65, 26, 66, 0, 0, 0, 38, 280, 46,
	248, 247, 47, 39, 40, 0, 57, 0, 48, 0,
	0, 249, 250, 0, 56, 55, 60, 61, 0, 254,
	97, 42, 43, 44, 45, 0, 0, 196, 197, 0,
	0, 0, 0, 0, 0, 272, 0, 62, 0, 63,
	0, 59, 58, 0, 278, 74, 0, 0, 0, 0,
	0, 0, 269, 0, 0, 0, 291, 0, 273, 0,
	274, 294, 0, 0, 0, 279, 0, 0, 0, 283,
	0, 0, 0, 299, 0, 72, 73, 0, 0, 0,
	75, 76, 77, 0, 305, 0, 0, 0, 71, 80,
	78, 79, 82, 0, 300, 301, 0, 0, 0, 0,
	0, 303, 0, 0, 0, 0, 0, 0, 0, 315,
	24, 25, 65, 26, 66, 0, 310, 0, 38, 245,
	46, 248, 247, 47, 39, 40, 0, 57, 314, 48,
	0, 0, 0, 0, 0, 56, 55, 60, 61, 0,
	0, 0, 42, 43, 44, 45, 0, 0, 196, 197,
	0, 0, 0, 0, 0, 0, 0, 0, 62, 0,
	63, 0, 59, 58, 24, 25, 65, 26, 66, 0,
	0, 0, 38, 298, 46, 0, 0, 47, 39, 40,
	0, 57, 0, 48, 0, 0, 0, 0, 0, 56,
	55, 60, 61, 0, 0, 0, 42, 43, 44, 45,
	0, 0, 196, 197, 0, 0, 0, 0, 0, 0,
	0, 0, 62, 0, 63, 0, 59, 58, 24, 25,
	65, 26, 66, 0, 0, 0, 38, 285, 46, 0,
	0, 47, 39, 40, 0, 57, 0, 48, 0, 0,
	0, 0, 0, 56, 55, 60, 61, 0, 0, 0,
	42, 43, 44, 45, 0, 0, 196, 197, 0, 0,
	0, 0, 0, 0, 0, 0, 62, 0, 63, 0,
	59, 58, 24, 25, 65, 26, 66, 0, 0, 0,
	38, 282, 46, 0, 0, 47, 39, 40, 0, 57,
	0, 48, 0, 0, 0, 0, 0, 56, 55, 60,
	61, 0, 0, 0, 42, 43, 44, 45, 0, 0,
	196, 197, 0, 24, 25, 65, 26, 66, 0, 0,
	62, 38, 63, 46, 59, 58, 47, 39, 40, 0,
	57, 0, 48, 0, 0, 0, 0, 0, 56, 55,
	60, 61, 0, 0, 0, 42, 43, 44, 45, 0,
	0, 196, 197, 0, 0, 0, 0, 0, 0, 0,
	0, 62, 0, 63, 268, 59, 58, 24, 25, 65,
	26, 66, 0, 0, 0, 38, 265, 46, 0, 0,
	47, 39, 40, 0, 57, 0, 48, 0, 0, 0,
	0, 0, 56, 55, 60, 61, 0, 0, 0, 42,
	43, 44, 45, 0, 0, 196, 197, 0, 0, 0,
	0, 0, 0, 0, 0, 62, 0, 63, 0, 59,
	58, 24, 25, 65, 26, 66, 0, 0, 0, 38,
	244, 46, 0, 0, 47, 39, 40, 0, 57, 0,
	48, 0, 0, 0, 0, 0, 56, 55, 60, 61,
	0, 0, 0, 42, 43, 44, 45, 0, 0, 196,
	197, 0, 0, 0, 0, 0, 0, 0, 0, 62,
	0, 63, 0, 59, 58, 24, 25, 65, 26, 66,
	0, 0, 0, 38, 242, 46, 0, 0, 47, 39,
	40, 0, 57, 0, 48, 0, 0, 0, 0, 0,
	56, 55, 60, 61, 0, 0, 0, 42, 43, 44,
	45, 0, 0, 196, 197, 0, 24, 25, 65, 26,
	66, 0, 0, 62, 38, 63, 46, 59, 58, 47,
	39, 40, 0, 57, 0, 48, 0, 0, 0, 0,
	0, 56, 55, 60, 61, 0, 0, 0, 42, 43,
	44, 45, 0, 0, 196, 197, 0, 0, 0, 0,
	0, 0, 0, 0, 62, 0, 63, 233, 59, 58,
	24, 25, 65, 26, 66, 0, 0, 0, 38, 228,
	46, 0, 0, 47, 39, 40, 0, 57, 0, 48,
	0, 0, 0, 0, 0, 56, 55, 60, 61, 0,
	0, 0, 42, 43, 44, 45, 0, 0, 196, 197,
	0, 24, 25, 65, 26, 66, 0, 0, 62, 38,
	63, 46, 59, 58, 47, 39, 40, 0, 57, 0,
	48, 0, 0, 0, 0, 0, 56, 55, 60, 61,
	0, 0, 0, 42, 43, 44, 45, 0, 0, 196,
	197, 0, 24, 25, 65, 26, 66, 217, 0, 62,
	38, 63, 46, 59, 58, 47, 39, 40, 0, 57,
	0, 48, 0, 0, 0, 0, 0, 56, 55, 60,
	61, 0, 0, 0, 42, 43, 44, 45, 0, 0,
	0, 216, 0, 24, 25, 65, 26, 66, 0, 0,
	62, 38, 63, 46, 59, 58, 47, 39, 40, 0,
	57, 0, 48, 0, 0, 0, 0, 0, 56, 55,
	60, 61, 0, 0, 0, 42, 43, 44, 45, 24,
	93, 65, 94, 220, 0, 0, 99, 0, 0, 0,
	0, 62, 0, 63, 0, 59, 58, 24, 93, 65,
	94, 84, 0, 0, 99, 0, 60, 61, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 60, 61, 0, 62, 0, 100,
	0, 59, 58, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 87, 0, 100, 0, 59,
	58,
}
var RubyPact = []int{

	-24, 9, -1000, -1000, -1000, 17, -1000, -1000, -1000, 710,
	-1000, -1000, -1000, 86, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 125, 115, 81, 71, 70,
	-1000, -1000, -1000, -1000, -1000, -1000, 238, -22, 253, 212,
	212, 92, 1468, 1468, 1468, 1468, 1468, 1468, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 209, 209, 1468, 29, 239,
	-1000, -1000, 209, -1000, -23, 179, -1000, -1000, -1000, -1000,
	1468, 273, 209, 209, 209, 209, 1468, 209, 209, 209,
	209, 1468, 209, 1468, 209, -1000, 160, 209, 209, 272,
	181, 112, -1000, 1522, 163, -1000, -1000, -1000, 161, -27,
	-27, 209, 1468, 1468, 1468, 1468, 265, 209, 209, 1468,
	1468, 262, 135, 135, 31, -1000, -1000, 1468, 259, 140,
	140, 140, 140, 140, 57, 1386, 168, 112, 168, 105,
	-1000, -1000, 217, -1000, -1000, 43, -4, 821, -1000, 669,
	219, 209, 1427, 140, 1504, 112, 112, 112, 112, 140,
	112, 112, 112, 112, 140, 139, 140, 226, 821, 604,
	242, 112, -1000, 604, 1345, -1000, 257, -1000, 1291, 224,
	140, 140, 140, 140, -1000, 112, 112, 140, 140, -1000,
	-1000, 166, 64, -1000, 62, 252, 248, -1000, 1250, 212,
	1196, 140, -1000, 885, -1000, -1000, -1000, -1000, 710, 140,
	104, 209, -1000, -1000, -1000, -1000, 209, -1000, -1000, -1000,
	164, 177, 404, -1000, 199, 140, -1000, -1000, 160, -1000,
	209, 209, -1000, 112, -1000, 23, 112, -1000, -1000, 1142,
	73, -1000, 1088, -1000, -1000, 8, 64, 195, 1468, -1000,
	-1000, 8, -1000, -1000, -1000, -1000, 187, 1468, -1000, 764,
	1047, -1000, 190, 112, 993, 112, 2, 12, -1000, 1468,
	209, -1000, 157, 112, 1468, -1000, -1000, 234, -1000, 1386,
	-1000, -1000, 140, 1386, 939, -1000, 1468, -1000, 140, 1386,
	-1000, 183, -1000, 1386, 228, -1000, -1000, 1468, -1000, 119,
	710, 140, 112, -1000, 140, -1000, 72, 68, -1000, 140,
	1386, 1386, -1000, 1386, 233, 123, -15, 8, -1000, -1000,
	1386, -1000, 1468, 209, 1386, 138, 88, 8, -1000, 8,
	-1000, 8, 8,
}
var RubyPgo = []int{

	0, 576, 371, 369, 6, 364, 358, 356, 486, 353,
	349, 348, 347, 346, 0, 442, 287, 342, 341, 340,
	8, 339, 1, 243, 338, 3, 336, 335, 334, 333,
	329, 327, 325, 320, 310, 309, 307, 596, 305, 11,
	12, 2, 303, 5, 298, 295, 10, 294, 34, 288,
	286, 9, 4, 244, 38,
}
var RubyR1 = []int{

	0, 42, 42, 42, 42, 42, 42, 42, 42, 42,
	42, 54, 54, 37, 37, 37, 37, 37, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 20, 20, 20, 20, 20, 20, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 39, 39, 48, 48, 46, 46, 46, 46, 51,
	51, 51, 51, 50, 50, 50, 17, 17, 43, 43,
	22, 22, 22, 22, 52, 52, 52, 21, 21, 24,
	25, 25, 53, 53, 11, 11, 11, 11, 11, 11,
	8, 8, 23, 23, 15, 15, 26, 26, 27, 28,
	29, 30, 31, 32, 33, 34, 35, 36, 2, 5,
	6, 6, 3, 3, 44, 44, 44, 44, 49, 49,
	49, 4, 4, 4, 4, 40, 47, 47, 47, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 41,
	41, 41, 41, 38, 38, 38, 7, 13, 45, 45,
	45, 45, 18, 19, 9, 12,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 4, 1, 4, 4, 2, 3, 3,
	4, 4, 3, 2, 3, 3, 3, 3, 3, 4,
	6, 3, 1, 1, 3, 0, 1, 1, 3, 1,
	1, 3, 3, 1, 3, 3, 7, 7, 1, 3,
	1, 2, 3, 2, 0, 1, 3, 4, 6, 4,
	1, 3, 1, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 3, 3, 3, 3, 2, 2,
	2, 2, 3, 3, 3, 3, 3, 3, 1, 1,
	3, 3, 5, 5, 0, 4, 7, 8, 3, 7,
	8, 3, 4, 4, 3, 3, 0, 1, 3, 4,
	5, 3, 3, 3, 3, 3, 5, 6, 5, 4,
	3, 3, 2, 0, 2, 2, 3, 4, 0, 3,
	4, 6, 2, 2, 5, 5,
}
var RubyChk = []int{

	-1000, -42, 43, 44, 60, -1, 43, 44, 60, -14,
	-17, -21, -24, -11, -27, -28, -29, -30, -10, -13,
	-20, -18, -19, -12, 5, 6, 8, -23, -15, -8,
	-2, -5, -6, -3, -25, -9, -16, -7, 13, 19,
	20, -26, 37, 38, 39, 40, 15, 18, 24, -31,
	-32, -33, -34, -35, -36, 31, 30, 22, 58, 57,
	32, 33, 53, 55, -53, 7, 9, 44, 43, 60,
	15, 47, 34, 35, 4, 39, 40, 41, 49, 50,
	48, 18, 51, 18, 9, -39, -51, 53, 36, 11,
	-50, -14, -4, 6, 8, -23, -15, -8, -16, 12,
	55, 9, 36, 36, 36, 36, 47, 34, 35, 15,
	18, 47, 6, 4, -25, 8, -25, 36, 11, -1,
	-1, -1, -1, -1, -1, -37, -48, -14, -48, -1,
	6, 8, 58, 6, 8, -48, -46, -14, -20, -54,
	46, 9, -38, -1, 6, -14, -14, -14, -14, -1,
	-14, -14, -14, -14, -1, -14, -1, -46, -14, 11,
	-14, -14, 6, 11, -37, -40, 48, -40, -37, -46,
	-1, -1, -1, -1, 6, -14, -14, -1, -1, 6,
	-43, -52, 9, -22, 6, 41, 50, -43, -37, 34,
	-37, -1, 6, -37, 43, 44, 43, 44, -14, -1,
	-45, 11, 43, 6, 8, 54, 11, 54, 43, -44,
	-49, -14, 6, 8, -46, -1, 44, 10, -51, -39,
	9, 45, 10, -14, -4, 54, -14, -4, 14, -37,
	-47, 6, -37, 56, 10, -54, 11, -52, 36, 6,
	6, -54, 14, -25, 14, 14, -41, 17, 16, -37,
	-37, 14, 25, -14, -37, -14, -54, -54, 11, 4,
	45, 10, -46, -14, 36, 14, 48, 11, 56, -37,
	-22, 10, -1, -37, -37, 14, 17, 16, -1, -37,
	14, -41, 14, -37, 8, 14, 56, 11, 56, -54,
	-14, -1, -14, 10, -1, 6, -54, -54, 14, -1,
	-37, -37, 14, -37, 4, -1, 6, -54, 14, 14,
	-37, 6, 4, 45, -37, -1, -14, -54, 11, -54,
	11, -54, -54,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 38,
	39, 40, 41, 42, 43, 44, 45, 46, 47, 48,
	49, 50, 51, 52, 18, 19, -2, 21, 22, 23,
	24, 25, 26, 27, 28, 29, 30, 31, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 13, 32,
	33, 34, 35, 36, 37, 0, 0, 0, 0, 0,
	128, 129, 75, 11, 0, 54, 163, 5, 6, 7,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, -2, 57, 63, 75, 0, 0,
	72, 79, 80, 19, -2, 21, 22, 23, 30, 13,
	-2, 75, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 94, 94, 13, -2, 13, 0, 0, 118,
	119, 120, 121, 13, 0, 168, 172, 73, 173, 0,
	112, 113, 0, 110, 111, 0, 0, 73, 77, 134,
	0, 75, 0, 151, 58, 64, 66, 68, 122, 123,
	124, 125, 126, 127, 153, 0, 155, 0, 76, 0,
	73, 104, 116, 0, 0, 13, 146, 13, 0, 0,
	105, 106, 107, 108, 59, 65, 67, 152, 154, 62,
	11, 88, 94, 95, 90, 0, 0, 11, 0, 0,
	0, 109, 117, 0, 13, 13, 14, 15, 16, 17,
	0, 0, 13, 114, 115, 130, 0, 131, 12, 11,
	11, 0, 19, -2, 0, 164, 165, 166, 60, 61,
	-2, 0, 53, 81, 82, 69, 84, 85, 141, 0,
	0, 147, 0, 144, 56, 13, 0, 0, 0, 91,
	93, 13, 97, 13, 99, 149, 0, 0, 13, 0,
	0, 167, 13, 74, 0, 78, 0, 0, 11, 0,
	0, 55, 0, 174, 0, 142, 145, 0, 143, 11,
	96, 89, 92, 11, 0, 150, 0, 13, 13, 162,
	156, 0, 158, 169, 13, 175, 132, 0, 133, 0,
	38, 11, 138, 71, 70, 148, 0, 0, 98, 13,
	160, 161, 157, 170, 0, 0, 0, 135, 86, 87,
	159, 13, 0, 0, 171, 11, 11, 136, 11, 139,
	11, 137, 140,
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
	52, 53, 54, 55, 56, 57, 58, 59, 60,
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
		//line parser.y:166
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:168
		{
		}
	case 3:
		//line parser.y:170
		{
		}
	case 4:
		//line parser.y:172
		{
		}
	case 5:
		//line parser.y:174
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:176
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:178
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:184
		{
		}
	case 11:
		//line parser.y:186
		{
		}
	case 12:
		//line parser.y:187
		{
		}
	case 13:
		//line parser.y:190
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 14:
		//line parser.y:192
		{
		}
	case 15:
		//line parser.y:194
		{
		}
	case 16:
		//line parser.y:196
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 17:
		//line parser.y:198
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 51:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 52:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 53:
		//line parser.y:208
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 54:
		//line parser.y:215
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 55:
		//line parser.y:219
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 56:
		//line parser.y:226
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 57:
		//line parser.y:233
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 58:
		//line parser.y:240
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 59:
		//line parser.y:247
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 60:
		//line parser.y:254
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 61:
		//line parser.y:262
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 62:
		//line parser.y:270
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 63:
		//line parser.y:280
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 64:
		//line parser.y:287
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 65:
		//line parser.y:295
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 66:
		//line parser.y:303
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 67:
		//line parser.y:311
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 68:
		//line parser.y:319
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 69:
		//line parser.y:329
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 70:
		//line parser.y:339
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 71:
		//line parser.y:348
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 72:
		//line parser.y:350
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 73:
		//line parser.y:353
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 74:
		//line parser.y:355
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 75:
		//line parser.y:357
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 76:
		//line parser.y:359
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:361
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 78:
		//line parser.y:363
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 79:
		//line parser.y:367
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 80:
		//line parser.y:369
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 81:
		//line parser.y:371
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 82:
		//line parser.y:373
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 83:
		//line parser.y:376
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 84:
		//line parser.y:378
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:380
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 86:
		//line parser.y:386
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 87:
		//line parser.y:394
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 88:
		//line parser.y:403
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 89:
		//line parser.y:405
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 90:
		//line parser.y:408
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 91:
		//line parser.y:410
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 92:
		//line parser.y:412
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 93:
		//line parser.y:414
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 94:
		//line parser.y:416
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 95:
		//line parser.y:418
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 96:
		//line parser.y:422
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 97:
		//line parser.y:427
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 98:
		//line parser.y:434
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 99:
		//line parser.y:444
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 100:
		//line parser.y:453
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 101:
		//line parser.y:459
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 102:
		//line parser.y:467
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 103:
		//line parser.y:471
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 104:
		//line parser.y:476
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 105:
		//line parser.y:483
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 106:
		//line parser.y:490
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 107:
		//line parser.y:497
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 108:
		//line parser.y:504
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 109:
		//line parser.y:511
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 110:
		//line parser.y:519
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 111:
		//line parser.y:521
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 112:
		//line parser.y:524
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 113:
		//line parser.y:526
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 114:
		//line parser.y:529
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 115:
		//line parser.y:531
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 116:
		//line parser.y:534
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 117:
		//line parser.y:536
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 118:
		//line parser.y:538
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 119:
		//line parser.y:539
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 120:
		//line parser.y:540
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 121:
		//line parser.y:541
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 122:
		//line parser.y:544
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 123:
		//line parser.y:553
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 124:
		//line parser.y:562
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 125:
		//line parser.y:571
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 126:
		//line parser.y:580
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 127:
		//line parser.y:589
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 128:
		//line parser.y:597
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 129:
		//line parser.y:598
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 130:
		//line parser.y:600
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 131:
		//line parser.y:602
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 132:
		//line parser.y:605
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 133:
		//line parser.y:613
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 134:
		//line parser.y:621
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 135:
		//line parser.y:623
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 136:
		//line parser.y:630
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 137:
		//line parser.y:637
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 138:
		//line parser.y:645
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 139:
		//line parser.y:652
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 140:
		//line parser.y:659
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 141:
		//line parser.y:667
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 142:
		//line parser.y:669
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 143:
		//line parser.y:676
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 144:
		//line parser.y:683
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 145:
		//line parser.y:686
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 146:
		//line parser.y:688
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 147:
		//line parser.y:690
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 148:
		//line parser.y:692
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 149:
		//line parser.y:695
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 150:
		//line parser.y:702
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 151:
		//line parser.y:710
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 152:
		//line parser.y:717
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 153:
		//line parser.y:724
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 154:
		//line parser.y:731
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 155:
		//line parser.y:738
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 156:
		//line parser.y:745
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 157:
		//line parser.y:752
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 158:
		//line parser.y:760
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 159:
		//line parser.y:769
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 160:
		//line parser.y:776
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 161:
		//line parser.y:783
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 162:
		//line parser.y:790
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 163:
		//line parser.y:797
		{
		}
	case 164:
		//line parser.y:798
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 165:
		//line parser.y:799
		{
		}
	case 166:
		//line parser.y:802
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 167:
		//line parser.y:805
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 168:
		//line parser.y:813
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 169:
		//line parser.y:815
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 170:
		//line parser.y:817
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 171:
		//line parser.y:826
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
	case 172:
		//line parser.y:841
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 173:
		//line parser.y:849
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 174:
		//line parser.y:858
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 175:
		//line parser.y:867
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	}
	goto Rubystack /* stack new state and value */
}
