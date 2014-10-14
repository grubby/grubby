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
const NEXT = 57370
const REDO = 57371
const RETRY = 57372
const RETURN = 57373
const YIELD = 57374
const TRUE = 57375
const FALSE = 57376
const LESSTHAN = 57377
const GREATERTHAN = 57378
const EQUALTO = 57379
const BANG = 57380
const COMPLEMENT = 57381
const POSITIVE = 57382
const NEGATIVE = 57383
const STAR = 57384
const WHITESPACE = 57385
const NEWLINE = 57386
const SEMICOLON = 57387
const COLON = 57388
const DOUBLECOLON = 57389
const DOT = 57390
const PIPE = 57391
const SLASH = 57392
const AMPERSAND = 57393
const QUESTIONMARK = 57394
const CARET = 57395
const LBRACKET = 57396
const RBRACKET = 57397
const LBRACE = 57398
const RBRACE = 57399
const DOLLARSIGN = 57400
const ATSIGN = 57401
const FILE_CONST_REF = 57402
const EOF = 57403

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
	"NEXT",
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

//line parser.y:876

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 27,
	47, 105,
	-2, 20,
	-1, 87,
	10, 78,
	11, 78,
	-2, 166,
	-1, 97,
	47, 105,
	-2, 20,
	-1, 103,
	5, 11,
	6, 11,
	7, 11,
	8, 11,
	9, 11,
	11, 11,
	33, 11,
	34, 11,
	44, 11,
	54, 11,
	56, 11,
	57, 11,
	58, 11,
	59, 11,
	-2, 13,
	-1, 118,
	47, 105,
	-2, 103,
	-1, 216,
	47, 106,
	-2, 104,
	-1, 223,
	10, 78,
	11, 78,
	-2, 166,
}

const RubyNprod = 179
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 1604

var RubyAct = []int{

	9, 184, 249, 186, 168, 35, 89, 95, 183, 2,
	3, 88, 211, 169, 138, 73, 139, 133, 209, 134,
	114, 290, 110, 111, 143, 291, 4, 94, 20, 204,
	112, 270, 316, 113, 211, 109, 25, 96, 68, 97,
	87, 267, 92, 102, 71, 70, 117, 119, 192, 121,
	110, 111, 86, 241, 211, 73, 108, 130, 130, 107,
	309, 72, 210, 109, 63, 64, 140, 289, 91, 269,
	135, 129, 131, 208, 28, 120, 148, 149, 150, 151,
	73, 153, 154, 155, 156, 90, 158, 103, 161, 62,
	61, 163, 164, 312, 141, 73, 311, 94, 211, 187,
	104, 98, 185, 106, 160, 161, 254, 162, 170, 197,
	198, 178, 179, 261, 321, 239, 141, 255, 73, 141,
	204, 172, 166, 211, 205, 190, 211, 315, 105, 201,
	104, 98, 98, 141, 305, 188, 280, 279, 73, 187,
	98, 296, 209, 214, 189, 161, 274, 239, 94, 144,
	98, 98, 98, 98, 221, 98, 98, 98, 98, 222,
	98, 217, 98, 226, 77, 98, 98, 229, 201, 287,
	227, 98, 201, 141, 230, 188, 264, 209, 278, 98,
	280, 279, 77, 118, 189, 98, 98, 240, 314, 323,
	237, 209, 201, 216, 201, 75, 76, 201, 246, 225,
	209, 298, 206, 77, 207, 256, 224, 136, 74, 137,
	258, 243, 85, 75, 76, 242, 94, 98, 262, 98,
	307, 116, 98, 115, 161, 266, 74, 234, 195, 182,
	85, 177, 165, 201, 75, 76, 201, 98, 147, 67,
	265, 98, 93, 273, 37, 213, 233, 74, 203, 75,
	76, 85, 141, 201, 201, 284, 212, 1, 201, 145,
	55, 54, 74, 293, 295, 53, 85, 52, 51, 50,
	17, 101, 16, 201, 15, 142, 14, 201, 201, 98,
	42, 12, 11, 201, 98, 29, 22, 201, 21, 10,
	98, 24, 19, 23, 13, 18, 36, 38, 98, 98,
	33, 101, 101, 32, 201, 201, 34, 201, 31, 0,
	101, 0, 99, 0, 201, 0, 0, 319, 201, 0,
	101, 101, 101, 101, 0, 101, 101, 101, 101, 0,
	101, 0, 101, 0, 0, 101, 101, 0, 98, 0,
	0, 101, 99, 99, 0, 0, 0, 0, 0, 101,
	0, 99, 0, 0, 0, 101, 101, 0, 0, 30,
	0, 99, 99, 99, 99, 0, 99, 99, 99, 99,
	0, 99, 0, 99, 0, 0, 99, 99, 0, 0,
	0, 0, 99, 0, 0, 0, 100, 101, 0, 101,
	99, 98, 101, 0, 0, 0, 99, 99, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 101, 0, 0,
	0, 101, 0, 0, 0, 0, 100, 100, 0, 0,
	0, 0, 0, 0, 0, 100, 0, 0, 99, 0,
	99, 0, 202, 99, 5, 100, 100, 100, 100, 0,
	100, 100, 100, 100, 0, 100, 0, 100, 99, 101,
	100, 100, 99, 0, 101, 0, 100, 0, 0, 238,
	101, 0, 128, 0, 100, 0, 244, 0, 101, 101,
	100, 100, 0, 0, 0, 0, 122, 123, 124, 125,
	126, 127, 0, 0, 0, 0, 0, 0, 259, 260,
	99, 132, 0, 0, 0, 99, 0, 0, 0, 0,
	0, 99, 100, 0, 100, 0, 146, 100, 101, 99,
	99, 0, 152, 0, 0, 0, 0, 157, 0, 159,
	0, 0, 100, 0, 0, 0, 100, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 292, 173, 174,
	175, 176, 0, 0, 0, 180, 181, 0, 299, 99,
	0, 0, 300, 194, 0, 0, 0, 0, 0, 0,
	0, 101, 0, 0, 100, 167, 171, 0, 0, 100,
	310, 0, 0, 77, 0, 100, 0, 0, 218, 0,
	191, 0, 193, 100, 100, 0, 0, 0, 0, 196,
	0, 0, 0, 0, 320, 322, 0, 324, 0, 325,
	0, 0, 99, 0, 75, 76, 0, 0, 0, 78,
	79, 80, 0, 0, 0, 0, 0, 74, 83, 81,
	82, 85, 0, 100, 228, 0, 0, 0, 0, 0,
	0, 232, 0, 235, 0, 0, 25, 26, 68, 27,
	69, 0, 0, 0, 39, 0, 47, 0, 0, 48,
	40, 41, 0, 58, 0, 49, 0, 0, 59, 60,
	252, 253, 57, 56, 63, 64, 0, 0, 257, 43,
	44, 45, 46, 0, 275, 0, 100, 25, 96, 68,
	97, 69, 0, 281, 102, 65, 0, 66, 0, 62,
	61, 0, 0, 0, 0, 294, 0, 0, 0, 0,
	297, 272, 0, 0, 0, 63, 64, 276, 0, 277,
	0, 0, 302, 0, 282, 0, 0, 0, 286, 0,
	0, 0, 0, 308, 0, 0, 65, 0, 103, 0,
	62, 61, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 303, 304, 0, 0, 0, 318, 0,
	306, 25, 96, 68, 97, 223, 0, 0, 102, 25,
	26, 68, 27, 69, 0, 313, 0, 39, 283, 47,
	251, 250, 48, 40, 41, 0, 58, 317, 49, 63,
	64, 59, 60, 0, 0, 57, 56, 63, 64, 0,
	0, 0, 43, 44, 45, 46, 0, 0, 199, 200,
	65, 0, 103, 0, 62, 61, 0, 0, 65, 0,
	66, 0, 62, 61, 25, 26, 68, 27, 69, 0,
	0, 0, 39, 248, 47, 251, 250, 48, 40, 41,
	0, 58, 0, 49, 0, 0, 59, 60, 0, 0,
	57, 56, 63, 64, 0, 0, 0, 43, 44, 45,
	46, 0, 0, 199, 200, 0, 25, 26, 68, 27,
	69, 0, 0, 65, 39, 66, 47, 62, 61, 48,
	40, 41, 0, 58, 0, 49, 0, 0, 59, 60,
	0, 0, 57, 56, 63, 64, 0, 0, 0, 43,
	44, 45, 46, 0, 0, 6, 7, 0, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 66, 0, 62,
	61, 0, 8, 25, 26, 68, 27, 69, 0, 0,
	0, 39, 301, 47, 0, 0, 48, 40, 41, 0,
	58, 0, 49, 0, 0, 59, 60, 0, 0, 57,
	56, 63, 64, 0, 0, 0, 43, 44, 45, 46,
	0, 0, 199, 200, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 66, 0, 62, 61, 25, 26,
	68, 27, 69, 0, 0, 0, 39, 288, 47, 0,
	0, 48, 40, 41, 0, 58, 0, 49, 0, 0,
	59, 60, 0, 0, 57, 56, 63, 64, 0, 0,
	0, 43, 44, 45, 46, 0, 0, 199, 200, 0,
	0, 0, 0, 0, 0, 0, 0, 65, 0, 66,
	0, 62, 61, 25, 26, 68, 27, 69, 0, 0,
	0, 39, 285, 47, 0, 0, 48, 40, 41, 0,
	58, 0, 49, 0, 0, 59, 60, 0, 0, 57,
	56, 63, 64, 0, 0, 0, 43, 44, 45, 46,
	0, 0, 199, 200, 0, 25, 26, 68, 27, 69,
	0, 0, 65, 39, 66, 47, 62, 61, 48, 40,
	41, 0, 58, 0, 49, 0, 0, 59, 60, 0,
	0, 57, 56, 63, 64, 0, 0, 0, 43, 44,
	45, 46, 0, 0, 199, 200, 0, 0, 0, 0,
	0, 0, 0, 0, 65, 0, 66, 271, 62, 61,
	25, 26, 68, 27, 69, 0, 0, 0, 39, 268,
	47, 0, 0, 48, 40, 41, 0, 58, 0, 49,
	0, 0, 59, 60, 0, 0, 57, 56, 63, 64,
	0, 0, 0, 43, 44, 45, 46, 0, 0, 199,
	200, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	0, 66, 0, 62, 61, 25, 26, 68, 27, 69,
	0, 0, 0, 39, 247, 47, 0, 0, 48, 40,
	41, 0, 58, 0, 49, 0, 0, 59, 60, 0,
	0, 57, 56, 63, 64, 0, 0, 0, 43, 44,
	45, 46, 0, 0, 199, 200, 0, 0, 0, 0,
	0, 0, 0, 0, 65, 0, 66, 0, 62, 61,
	25, 26, 68, 27, 69, 0, 0, 0, 39, 245,
	47, 0, 0, 48, 40, 41, 0, 58, 0, 49,
	0, 0, 59, 60, 0, 0, 57, 56, 63, 64,
	0, 0, 0, 43, 44, 45, 46, 0, 0, 199,
	200, 0, 25, 26, 68, 27, 69, 0, 0, 65,
	39, 66, 47, 62, 61, 48, 40, 41, 0, 58,
	0, 49, 0, 0, 59, 60, 0, 0, 57, 56,
	63, 64, 0, 0, 0, 43, 44, 45, 46, 0,
	0, 199, 200, 0, 0, 0, 0, 0, 0, 0,
	0, 65, 0, 66, 236, 62, 61, 25, 26, 68,
	27, 69, 0, 0, 0, 39, 231, 47, 0, 0,
	48, 40, 41, 0, 58, 0, 49, 0, 0, 59,
	60, 0, 0, 57, 56, 63, 64, 0, 0, 0,
	43, 44, 45, 46, 0, 0, 199, 200, 0, 25,
	26, 68, 27, 69, 0, 0, 65, 39, 66, 47,
	62, 61, 48, 40, 41, 0, 58, 0, 49, 0,
	0, 59, 60, 0, 0, 57, 56, 63, 64, 0,
	0, 0, 43, 44, 45, 46, 0, 0, 199, 200,
	0, 25, 26, 68, 27, 69, 220, 0, 65, 39,
	66, 47, 62, 61, 48, 40, 41, 0, 58, 0,
	49, 0, 0, 59, 60, 0, 0, 57, 56, 63,
	64, 0, 0, 0, 43, 44, 45, 46, 0, 0,
	0, 219, 25, 96, 68, 97, 87, 0, 0, 102,
	65, 0, 66, 0, 62, 61, 0, 0, 25, 215,
	68, 97, 69, 0, 0, 0, 0, 0, 0, 0,
	63, 64, 25, 96, 68, 97, 87, 0, 0, 102,
	0, 0, 0, 263, 0, 0, 63, 64, 0, 0,
	0, 90, 0, 103, 0, 62, 61, 211, 0, 0,
	63, 64, 25, 96, 68, 97, 69, 65, 0, 66,
	0, 62, 61, 0, 0, 0, 0, 0, 0, 0,
	0, 90, 0, 103, 0, 62, 61, 77, 0, 0,
	63, 64, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 84, 0, 0, 0, 77, 0, 0, 0, 0,
	0, 65, 0, 66, 0, 62, 61, 0, 75, 76,
	0, 0, 0, 78, 79, 80, 0, 0, 0, 0,
	0, 74, 83, 81, 82, 85, 75, 76, 0, 0,
	0, 78, 79, 80, 0, 0, 0, 0, 0, 74,
	83, 81, 82, 85,
}
var RubyPact = []int{

	-35, 851, -1000, -1000, -1000, 0, -1000, -1000, -1000, 1533,
	-1000, -1000, -1000, 34, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 31, 91, 66, 22,
	19, -1000, -1000, -1000, -1000, -1000, -1000, 15, -28, 217,
	175, 175, 38, 631, 631, 631, 631, 631, 631, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 1507, 1507, 631, -1000,
	-1000, 11, 201, -1000, -1000, 1507, -1000, -23, 140, -1000,
	-1000, -1000, -1000, 631, 232, 1507, 1507, 1507, 1507, 631,
	1507, 1507, 1507, 1507, 631, 1507, 631, 1507, -1000, 96,
	1507, 1507, 226, 111, 199, -1000, 1477, 121, -1000, -1000,
	-1000, -13, -36, -36, 1507, 631, 631, 631, 631, 225,
	1507, 1507, 631, 631, 223, 93, 93, 13, -1000, -1000,
	631, 222, 40, 40, 40, 40, 40, 65, 1364, 109,
	199, 109, 80, -1000, -1000, 196, -1000, -1000, 18, 7,
	1551, -1000, 1463, 185, 1507, 1406, 40, 746, 199, 199,
	199, 199, 40, 199, 199, 199, 199, 40, 160, 40,
	189, 1551, 672, 569, 199, -1000, 672, 1322, -1000, 221,
	-1000, 1267, 180, 40, 40, 40, 40, -1000, 199, 199,
	40, 40, -1000, -1000, 104, 133, -1000, 16, 209, 205,
	-1000, 1225, 175, 1170, 40, -1000, 809, -1000, -1000, -1000,
	-1000, 1533, 40, 92, 1507, -1000, -1000, -1000, -1000, 1507,
	-1000, -1000, -1000, 102, 214, 1447, -1000, 166, 40, -1000,
	-1000, 96, -1000, 1507, 1507, -1000, 199, -1000, 4, 199,
	-1000, -1000, 1115, 20, -1000, 1060, -1000, -1000, -10, 133,
	136, 631, -1000, -1000, -10, -1000, -1000, -1000, -1000, 164,
	631, -1000, 754, 1018, -1000, 161, 199, 963, 199, 10,
	-32, -1000, 631, 1507, -1000, 131, 199, 631, -1000, -1000,
	195, -1000, 1364, -1000, -1000, 40, 1364, 908, -1000, 631,
	-1000, 40, 1364, -1000, 120, -1000, 1364, 216, -1000, -1000,
	631, -1000, 54, 1533, 40, 199, -1000, 40, -1000, 82,
	79, -1000, 40, 1364, 1364, -1000, 1364, 182, 123, -14,
	-10, -1000, -1000, 1364, -1000, 631, 1507, 1364, 103, 178,
	-10, -1000, -10, -1000, -10, -10,
}
var RubyPgo = []int{

	0, 432, 308, 306, 7, 303, 300, 297, 359, 296,
	295, 294, 293, 292, 0, 285, 291, 244, 289, 288,
	286, 28, 282, 3, 74, 281, 5, 280, 276, 274,
	272, 270, 269, 268, 267, 265, 261, 260, 462, 259,
	11, 4, 2, 257, 8, 256, 248, 16, 246, 14,
	245, 242, 6, 1, 239, 275,
}
var RubyR1 = []int{

	0, 43, 43, 43, 43, 43, 43, 43, 43, 43,
	43, 55, 55, 38, 38, 38, 38, 38, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 21, 21, 21, 21, 21, 21, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 16, 16, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 40, 40, 49, 49, 47, 47,
	47, 47, 52, 52, 52, 52, 51, 51, 51, 18,
	18, 44, 44, 23, 23, 23, 23, 53, 53, 53,
	22, 22, 25, 26, 26, 54, 54, 11, 11, 11,
	11, 11, 11, 8, 8, 24, 24, 15, 15, 27,
	27, 28, 29, 30, 31, 32, 33, 34, 35, 36,
	37, 2, 5, 6, 6, 3, 3, 45, 45, 45,
	45, 50, 50, 50, 4, 4, 4, 4, 41, 48,
	48, 48, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 42, 42, 42, 42, 39, 39, 39, 7,
	13, 46, 46, 46, 46, 19, 20, 9, 12,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 4, 1, 4, 4,
	2, 3, 3, 4, 4, 3, 2, 3, 3, 3,
	3, 3, 4, 6, 3, 1, 1, 3, 0, 1,
	1, 3, 1, 1, 3, 3, 1, 3, 3, 7,
	7, 1, 3, 1, 2, 3, 2, 0, 1, 3,
	4, 6, 4, 1, 3, 1, 3, 3, 3, 3,
	3, 3, 3, 2, 2, 2, 2, 3, 3, 3,
	3, 2, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 1, 1, 3, 3, 5, 5, 0, 4, 7,
	8, 3, 7, 8, 3, 4, 4, 3, 3, 0,
	1, 3, 4, 5, 3, 3, 3, 3, 3, 5,
	6, 5, 4, 3, 3, 2, 0, 2, 2, 3,
	4, 0, 3, 4, 6, 2, 2, 5, 5,
}
var RubyChk = []int{

	-1000, -43, 44, 45, 61, -1, 44, 45, 61, -14,
	-18, -22, -25, -11, -28, -29, -30, -31, -10, -13,
	-21, -19, -20, -12, -16, 5, 6, 8, -24, -15,
	-8, -2, -5, -6, -3, -26, -9, -17, -7, 13,
	19, 20, -27, 38, 39, 40, 41, 15, 18, 24,
	-32, -33, -34, -35, -36, -37, 32, 31, 22, 27,
	28, 59, 58, 33, 34, 54, 56, -54, 7, 9,
	45, 44, 61, 15, 48, 35, 36, 4, 40, 41,
	42, 50, 51, 49, 18, 52, 18, 9, -40, -52,
	54, 37, 11, -51, -14, -4, 6, 8, -24, -15,
	-8, -17, 12, 56, 9, 37, 37, 37, 37, 48,
	35, 36, 15, 18, 48, 6, 4, -26, 8, -26,
	37, 11, -1, -1, -1, -1, -1, -1, -38, -49,
	-14, -49, -1, 6, 8, 59, 6, 8, -49, -47,
	-14, -21, -55, 47, 9, -39, -1, 6, -14, -14,
	-14, -14, -1, -14, -14, -14, -14, -1, -14, -1,
	-47, -14, 11, -14, -14, 6, 11, -38, -41, 49,
	-41, -38, -47, -1, -1, -1, -1, 6, -14, -14,
	-1, -1, 6, -44, -53, 9, -23, 6, 42, 51,
	-44, -38, 35, -38, -1, 6, -38, 44, 45, 44,
	45, -14, -1, -46, 11, 44, 6, 8, 55, 11,
	55, 44, -45, -50, -14, 6, 8, -47, -1, 45,
	10, -52, -40, 9, 46, 10, -14, -4, 55, -14,
	-4, 14, -38, -48, 6, -38, 57, 10, -55, 11,
	-53, 37, 6, 6, -55, 14, -26, 14, 14, -42,
	17, 16, -38, -38, 14, 25, -14, -38, -14, -55,
	-55, 11, 4, 46, 10, -47, -14, 37, 14, 49,
	11, 57, -38, -23, 10, -1, -38, -38, 14, 17,
	16, -1, -38, 14, -42, 14, -38, 8, 14, 57,
	11, 57, -55, -14, -1, -14, 10, -1, 6, -55,
	-55, 14, -1, -38, -38, 14, -38, 4, -1, 6,
	-55, 14, 14, -38, 6, 4, 46, -38, -1, -14,
	-55, 11, -55, 11, -55, -55,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 38,
	39, 40, 41, 42, 43, 44, 45, 46, 47, 48,
	49, 50, 51, 52, 53, 18, 19, -2, 21, 22,
	23, 24, 25, 26, 27, 28, 29, 30, 31, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 13,
	32, 33, 34, 35, 36, 37, 0, 0, 0, 54,
	55, 0, 0, 131, 132, 78, 11, 0, 57, 166,
	5, 6, 7, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, -2, 60, 66,
	78, 0, 0, 75, 82, 83, 19, -2, 21, 22,
	23, 30, 13, -2, 78, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 97, 97, 13, -2, 13,
	0, 0, 121, 122, 123, 124, 13, 0, 171, 175,
	76, 176, 0, 115, 116, 0, 113, 114, 0, 0,
	76, 80, 137, 0, 78, 0, 154, 61, 67, 69,
	71, 125, 126, 127, 128, 129, 130, 156, 0, 158,
	0, 79, 0, 76, 107, 119, 0, 0, 13, 149,
	13, 0, 0, 108, 109, 110, 111, 62, 68, 70,
	155, 157, 65, 11, 91, 97, 98, 93, 0, 0,
	11, 0, 0, 0, 112, 120, 0, 13, 13, 14,
	15, 16, 17, 0, 0, 13, 117, 118, 133, 0,
	134, 12, 11, 11, 0, 19, -2, 0, 167, 168,
	169, 63, 64, -2, 0, 56, 84, 85, 72, 87,
	88, 144, 0, 0, 150, 0, 147, 59, 13, 0,
	0, 0, 94, 96, 13, 100, 13, 102, 152, 0,
	0, 13, 0, 0, 170, 13, 77, 0, 81, 0,
	0, 11, 0, 0, 58, 0, 177, 0, 145, 148,
	0, 146, 11, 99, 92, 95, 11, 0, 153, 0,
	13, 13, 165, 159, 0, 161, 172, 13, 178, 135,
	0, 136, 0, 38, 11, 141, 74, 73, 151, 0,
	0, 101, 13, 163, 164, 160, 173, 0, 0, 0,
	138, 89, 90, 162, 13, 0, 0, 174, 11, 11,
	139, 11, 142, 11, 140, 143,
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
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
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
		//line parser.y:168
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:170
		{
		}
	case 3:
		//line parser.y:172
		{
		}
	case 4:
		//line parser.y:174
		{
		}
	case 5:
		//line parser.y:176
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:178
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:180
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:186
		{
		}
	case 11:
		//line parser.y:188
		{
		}
	case 12:
		//line parser.y:189
		{
		}
	case 13:
		//line parser.y:192
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 14:
		//line parser.y:194
		{
		}
	case 15:
		//line parser.y:196
		{
		}
	case 16:
		//line parser.y:198
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 17:
		//line parser.y:200
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 54:
		//line parser.y:210
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 55:
		//line parser.y:212
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 56:
		//line parser.y:215
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 57:
		//line parser.y:222
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 58:
		//line parser.y:226
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 59:
		//line parser.y:233
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 60:
		//line parser.y:240
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 61:
		//line parser.y:247
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 62:
		//line parser.y:254
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 63:
		//line parser.y:261
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 64:
		//line parser.y:269
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 65:
		//line parser.y:277
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 66:
		//line parser.y:287
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 67:
		//line parser.y:294
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 68:
		//line parser.y:302
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 69:
		//line parser.y:310
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 70:
		//line parser.y:318
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 71:
		//line parser.y:326
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 72:
		//line parser.y:336
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 73:
		//line parser.y:346
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 74:
		//line parser.y:355
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 75:
		//line parser.y:357
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 76:
		//line parser.y:360
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:362
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 78:
		//line parser.y:364
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 79:
		//line parser.y:366
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 80:
		//line parser.y:368
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 81:
		//line parser.y:370
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 82:
		//line parser.y:374
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
		//line parser.y:383
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 87:
		//line parser.y:385
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 88:
		//line parser.y:387
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 89:
		//line parser.y:393
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 90:
		//line parser.y:401
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 91:
		//line parser.y:410
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 92:
		//line parser.y:412
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 93:
		//line parser.y:415
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 94:
		//line parser.y:417
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 95:
		//line parser.y:419
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 96:
		//line parser.y:421
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 97:
		//line parser.y:423
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 98:
		//line parser.y:425
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 99:
		//line parser.y:429
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 100:
		//line parser.y:434
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 101:
		//line parser.y:441
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 102:
		//line parser.y:451
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 103:
		//line parser.y:460
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 104:
		//line parser.y:466
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 105:
		//line parser.y:474
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 106:
		//line parser.y:478
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 107:
		//line parser.y:483
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 108:
		//line parser.y:490
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 109:
		//line parser.y:497
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 110:
		//line parser.y:504
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 111:
		//line parser.y:511
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 112:
		//line parser.y:518
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 113:
		//line parser.y:526
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 114:
		//line parser.y:528
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 115:
		//line parser.y:531
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 116:
		//line parser.y:533
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 117:
		//line parser.y:536
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 118:
		//line parser.y:538
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 119:
		//line parser.y:541
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 120:
		//line parser.y:543
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 121:
		//line parser.y:545
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 122:
		//line parser.y:546
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 123:
		//line parser.y:547
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 124:
		//line parser.y:548
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 125:
		//line parser.y:551
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 126:
		//line parser.y:560
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 127:
		//line parser.y:569
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 128:
		//line parser.y:578
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 129:
		//line parser.y:587
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 130:
		//line parser.y:596
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 131:
		//line parser.y:604
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 132:
		//line parser.y:605
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 133:
		//line parser.y:607
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 134:
		//line parser.y:609
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 135:
		//line parser.y:612
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 136:
		//line parser.y:620
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 137:
		//line parser.y:628
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 138:
		//line parser.y:630
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 139:
		//line parser.y:637
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 140:
		//line parser.y:644
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 141:
		//line parser.y:652
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 142:
		//line parser.y:659
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 143:
		//line parser.y:666
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 144:
		//line parser.y:674
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 145:
		//line parser.y:676
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 146:
		//line parser.y:683
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 147:
		//line parser.y:690
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 148:
		//line parser.y:693
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 149:
		//line parser.y:695
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 150:
		//line parser.y:697
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 151:
		//line parser.y:699
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 152:
		//line parser.y:702
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 153:
		//line parser.y:709
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 154:
		//line parser.y:717
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 155:
		//line parser.y:724
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 156:
		//line parser.y:731
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 157:
		//line parser.y:738
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 158:
		//line parser.y:745
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 159:
		//line parser.y:752
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 160:
		//line parser.y:759
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 161:
		//line parser.y:767
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 162:
		//line parser.y:776
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 163:
		//line parser.y:783
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 164:
		//line parser.y:790
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 165:
		//line parser.y:797
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 166:
		//line parser.y:804
		{
		}
	case 167:
		//line parser.y:805
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 168:
		//line parser.y:806
		{
		}
	case 169:
		//line parser.y:809
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 170:
		//line parser.y:812
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 171:
		//line parser.y:820
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 172:
		//line parser.y:822
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{Body: RubyS[Rubypt-0].genericSlice})
		}
	case 173:
		//line parser.y:824
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			})
		}
	case 174:
		//line parser.y:833
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
	case 175:
		//line parser.y:848
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 176:
		//line parser.y:856
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 177:
		//line parser.y:865
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 178:
		//line parser.y:874
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	}
	goto Rubystack /* stack new state and value */
}
