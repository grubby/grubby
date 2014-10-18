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
const AND = 57375
const OR = 57376
const LAMBDA = 57377
const TRUE = 57378
const FALSE = 57379
const LESSTHAN = 57380
const GREATERTHAN = 57381
const EQUALTO = 57382
const BANG = 57383
const COMPLEMENT = 57384
const POSITIVE = 57385
const NEGATIVE = 57386
const STAR = 57387
const WHITESPACE = 57388
const NEWLINE = 57389
const SEMICOLON = 57390
const COLON = 57391
const DOUBLECOLON = 57392
const DOT = 57393
const PIPE = 57394
const SLASH = 57395
const AMPERSAND = 57396
const QUESTIONMARK = 57397
const CARET = 57398
const LBRACKET = 57399
const RBRACKET = 57400
const LBRACE = 57401
const RBRACE = 57402
const DOLLARSIGN = 57403
const ATSIGN = 57404
const FILE_CONST_REF = 57405
const EOF = 57406

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
	"AND",
	"OR",
	"LAMBDA",
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

//line parser.y:1028

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 28,
	50, 115,
	-2, 20,
	-1, 88,
	10, 77,
	11, 77,
	-2, 176,
	-1, 98,
	50, 115,
	-2, 20,
	-1, 104,
	13, 13,
	15, 13,
	18, 13,
	19, 13,
	20, 13,
	22, 13,
	24, 13,
	32, 13,
	35, 13,
	41, 13,
	42, 13,
	43, 13,
	44, 13,
	48, 13,
	-2, 11,
	-1, 119,
	50, 115,
	-2, 113,
	-1, 220,
	50, 116,
	-2, 114,
	-1, 227,
	10, 77,
	11, 77,
	-2, 176,
	-1, 351,
	27, 192,
	28, 192,
	-2, 13,
	-1, 352,
	27, 192,
	28, 192,
	-2, 13,
}

const RubyNprod = 218
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 2136

var RubyAct = []int{

	9, 277, 322, 140, 190, 368, 36, 253, 20, 188,
	90, 89, 215, 187, 72, 262, 141, 174, 305, 96,
	2, 3, 103, 275, 213, 306, 173, 208, 95, 111,
	112, 115, 135, 145, 136, 358, 72, 4, 272, 215,
	245, 109, 110, 108, 144, 107, 70, 69, 118, 120,
	344, 87, 196, 345, 215, 339, 76, 383, 131, 380,
	379, 130, 72, 71, 274, 142, 72, 304, 350, 134,
	206, 214, 5, 143, 212, 150, 151, 152, 153, 133,
	155, 156, 157, 158, 191, 160, 161, 162, 137, 165,
	74, 75, 167, 168, 351, 352, 215, 143, 95, 342,
	143, 76, 72, 73, 122, 164, 165, 84, 377, 347,
	260, 208, 182, 183, 143, 123, 124, 125, 126, 127,
	128, 172, 176, 192, 72, 191, 311, 213, 189, 132,
	205, 194, 193, 121, 209, 74, 75, 26, 97, 67,
	98, 88, 215, 148, 103, 218, 166, 165, 73, 154,
	95, 266, 84, 243, 159, 143, 201, 202, 163, 113,
	225, 226, 114, 221, 192, 170, 105, 230, 62, 63,
	76, 233, 205, 193, 341, 292, 205, 177, 178, 179,
	180, 268, 111, 112, 184, 185, 231, 279, 243, 91,
	234, 104, 198, 61, 60, 110, 205, 106, 205, 244,
	267, 205, 258, 250, 74, 75, 105, 215, 336, 261,
	259, 337, 146, 260, 263, 228, 334, 73, 222, 335,
	95, 84, 119, 26, 97, 67, 98, 68, 165, 271,
	103, 210, 242, 211, 74, 75, 143, 205, 375, 248,
	205, 357, 72, 378, 270, 380, 379, 73, 278, 220,
	281, 84, 72, 349, 62, 63, 313, 205, 205, 269,
	213, 264, 265, 296, 289, 241, 213, 247, 308, 310,
	329, 246, 285, 284, 138, 64, 139, 104, 317, 61,
	60, 238, 317, 205, 283, 331, 285, 284, 205, 229,
	213, 117, 205, 116, 199, 186, 181, 169, 149, 66,
	94, 217, 237, 207, 26, 219, 67, 98, 68, 216,
	1, 307, 320, 147, 56, 55, 280, 54, 53, 52,
	51, 17, 314, 343, 131, 286, 324, 346, 205, 205,
	16, 205, 15, 297, 14, 62, 63, 43, 309, 298,
	299, 22, 23, 312, 24, 319, 215, 12, 318, 205,
	11, 321, 318, 21, 340, 326, 64, 10, 65, 366,
	61, 60, 205, 296, 296, 296, 362, 363, 364, 372,
	19, 332, 333, 13, 18, 37, 338, 25, 39, 34,
	33, 35, 32, 296, 0, 0, 382, 0, 296, 296,
	296, 29, 0, 0, 0, 0, 387, 388, 0, 0,
	0, 0, 389, 0, 0, 353, 354, 355, 356, 0,
	374, 376, 0, 0, 0, 359, 360, 0, 0, 99,
	384, 0, 385, 0, 0, 0, 0, 0, 365, 0,
	0, 0, 0, 297, 297, 297, 38, 0, 0, 0,
	381, 0, 0, 0, 0, 0, 0, 0, 0, 99,
	386, 0, 0, 297, 0, 0, 99, 0, 297, 297,
	297, 0, 0, 0, 102, 0, 99, 99, 99, 99,
	0, 99, 99, 99, 99, 0, 99, 99, 99, 0,
	99, 0, 0, 99, 99, 0, 0, 0, 0, 99,
	0, 0, 0, 0, 102, 0, 0, 99, 0, 0,
	0, 102, 0, 99, 99, 76, 0, 0, 0, 0,
	0, 102, 102, 102, 102, 0, 102, 102, 102, 102,
	0, 102, 102, 102, 0, 102, 76, 0, 102, 102,
	0, 0, 0, 0, 102, 0, 99, 0, 99, 74,
	75, 99, 102, 0, 77, 78, 79, 0, 102, 102,
	0, 0, 73, 82, 80, 81, 84, 0, 99, 232,
	74, 75, 99, 0, 0, 77, 78, 79, 0, 0,
	0, 0, 0, 73, 82, 80, 81, 84, 0, 0,
	0, 102, 0, 102, 0, 0, 102, 0, 0, 0,
	0, 0, 0, 0, 30, 0, 0, 0, 0, 0,
	99, 0, 0, 102, 0, 99, 0, 102, 0, 0,
	0, 99, 0, 0, 0, 0, 0, 0, 0, 99,
	99, 0, 100, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 102, 0, 0, 0, 0,
	102, 0, 100, 0, 0, 0, 102, 0, 0, 100,
	99, 0, 0, 0, 102, 102, 0, 0, 0, 100,
	100, 100, 100, 175, 100, 100, 100, 100, 0, 100,
	100, 100, 0, 100, 0, 0, 100, 100, 0, 0,
	0, 0, 100, 31, 0, 0, 0, 0, 0, 0,
	100, 0, 0, 0, 0, 102, 100, 100, 0, 0,
	0, 0, 0, 0, 0, 99, 0, 0, 0, 0,
	0, 101, 0, 0, 129, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 100,
	0, 100, 0, 0, 100, 0, 0, 0, 0, 0,
	99, 101, 0, 0, 0, 0, 0, 0, 101, 0,
	102, 100, 0, 0, 0, 100, 0, 0, 101, 101,
	101, 101, 0, 101, 101, 101, 101, 171, 101, 101,
	101, 0, 101, 0, 0, 101, 101, 0, 0, 0,
	0, 101, 195, 0, 197, 102, 0, 0, 0, 101,
	0, 200, 0, 100, 0, 101, 101, 0, 100, 0,
	0, 0, 0, 0, 100, 0, 0, 0, 0, 0,
	0, 0, 100, 100, 0, 0, 26, 97, 67, 98,
	88, 0, 93, 103, 0, 0, 0, 0, 101, 0,
	101, 0, 0, 101, 0, 0, 236, 0, 239, 0,
	0, 0, 0, 0, 0, 0, 0, 62, 63, 0,
	101, 92, 0, 100, 101, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 256, 257, 0, 91, 0,
	104, 0, 61, 60, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 101, 0, 0, 0, 0, 101, 0, 0,
	0, 0, 0, 101, 0, 0, 0, 0, 100, 0,
	0, 101, 101, 0, 282, 0, 0, 0, 0, 287,
	0, 0, 0, 0, 291, 0, 0, 26, 27, 67,
	28, 68, 0, 0, 0, 40, 288, 48, 255, 254,
	49, 41, 42, 100, 58, 0, 50, 0, 0, 327,
	328, 0, 101, 0, 57, 0, 330, 59, 62, 63,
	0, 0, 0, 44, 45, 46, 47, 0, 0, 203,
	204, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	0, 65, 0, 61, 60, 0, 0, 0, 0, 0,
	348, 0, 0, 0, 0, 0, 200, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 101, 0, 0,
	0, 0, 0, 361, 0, 256, 257, 26, 27, 67,
	28, 68, 0, 0, 0, 40, 371, 300, 370, 369,
	301, 41, 42, 0, 58, 0, 50, 0, 0, 302,
	303, 0, 101, 0, 57, 0, 0, 59, 62, 63,
	0, 0, 0, 44, 45, 46, 47, 0, 0, 294,
	295, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	0, 65, 0, 61, 60, 26, 27, 67, 28, 68,
	0, 0, 0, 40, 367, 300, 370, 369, 301, 41,
	42, 0, 58, 0, 50, 0, 0, 302, 303, 0,
	0, 0, 57, 0, 0, 59, 62, 63, 0, 0,
	0, 44, 45, 46, 47, 0, 0, 294, 295, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 0, 65,
	0, 61, 60, 26, 27, 67, 28, 68, 0, 0,
	0, 40, 373, 300, 0, 0, 301, 41, 42, 0,
	58, 0, 50, 0, 0, 302, 303, 0, 0, 0,
	57, 0, 0, 59, 62, 63, 0, 0, 0, 44,
	45, 46, 47, 0, 0, 294, 295, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 0, 65, 0, 61,
	60, 26, 27, 67, 28, 68, 0, 0, 0, 40,
	293, 300, 0, 0, 301, 41, 42, 0, 58, 0,
	50, 0, 0, 302, 303, 0, 0, 0, 57, 0,
	0, 59, 62, 63, 0, 0, 0, 44, 45, 46,
	47, 0, 0, 294, 295, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 0, 65, 0, 61, 60, 26,
	27, 67, 28, 68, 0, 0, 0, 40, 252, 48,
	255, 254, 49, 41, 42, 0, 58, 0, 50, 0,
	0, 0, 0, 0, 0, 0, 57, 0, 0, 59,
	62, 63, 0, 0, 0, 44, 45, 46, 47, 0,
	0, 203, 204, 0, 26, 27, 67, 28, 68, 0,
	0, 64, 40, 65, 48, 61, 60, 49, 41, 42,
	0, 58, 0, 50, 0, 0, 0, 0, 0, 0,
	0, 57, 0, 0, 59, 62, 63, 0, 0, 0,
	44, 45, 46, 47, 0, 0, 6, 7, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 65, 0,
	61, 60, 0, 8, 26, 27, 67, 28, 68, 0,
	0, 0, 40, 0, 300, 0, 0, 301, 41, 42,
	0, 58, 0, 50, 0, 0, 302, 303, 0, 0,
	0, 57, 0, 0, 59, 62, 63, 0, 0, 0,
	44, 45, 46, 47, 0, 0, 294, 295, 0, 26,
	27, 67, 28, 68, 0, 0, 64, 40, 65, 48,
	61, 60, 49, 41, 42, 0, 58, 0, 50, 260,
	0, 0, 0, 0, 0, 323, 57, 0, 0, 59,
	62, 63, 0, 0, 0, 44, 45, 46, 47, 0,
	0, 315, 316, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 0, 65, 0, 61, 60, 26, 27, 67,
	28, 68, 0, 0, 0, 40, 325, 48, 0, 0,
	49, 41, 42, 0, 58, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 57, 0, 0, 59, 62, 63,
	0, 0, 0, 44, 45, 46, 47, 0, 0, 203,
	204, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	0, 65, 0, 61, 60, 26, 27, 67, 28, 68,
	0, 0, 0, 40, 290, 48, 0, 0, 49, 41,
	42, 0, 58, 0, 50, 0, 0, 0, 0, 0,
	0, 0, 57, 0, 0, 59, 62, 63, 0, 0,
	0, 44, 45, 46, 47, 0, 0, 203, 204, 0,
	26, 27, 67, 28, 68, 0, 0, 64, 40, 65,
	48, 61, 60, 49, 41, 42, 0, 58, 0, 50,
	0, 0, 0, 0, 0, 0, 0, 57, 0, 0,
	59, 62, 63, 0, 0, 0, 44, 45, 46, 47,
	0, 0, 203, 204, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 276, 61, 60, 26, 27,
	67, 28, 68, 0, 0, 0, 40, 273, 48, 0,
	0, 49, 41, 42, 0, 58, 0, 50, 0, 0,
	0, 0, 0, 0, 0, 57, 0, 0, 59, 62,
	63, 0, 0, 0, 44, 45, 46, 47, 0, 0,
	203, 204, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 0, 65, 0, 61, 60, 26, 27, 67, 28,
	68, 0, 0, 0, 40, 251, 48, 0, 0, 49,
	41, 42, 0, 58, 0, 50, 0, 0, 0, 0,
	0, 0, 0, 57, 0, 0, 59, 62, 63, 0,
	0, 0, 44, 45, 46, 47, 0, 0, 203, 204,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 0,
	65, 0, 61, 60, 26, 27, 67, 28, 68, 0,
	0, 0, 40, 249, 48, 0, 0, 49, 41, 42,
	0, 58, 0, 50, 0, 0, 0, 0, 0, 0,
	0, 57, 0, 0, 59, 62, 63, 0, 0, 0,
	44, 45, 46, 47, 0, 0, 203, 204, 0, 26,
	27, 67, 28, 68, 0, 0, 64, 40, 65, 48,
	61, 60, 49, 41, 42, 0, 58, 0, 50, 0,
	0, 0, 0, 0, 0, 0, 57, 0, 0, 59,
	62, 63, 0, 0, 0, 44, 45, 46, 47, 0,
	0, 203, 204, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 0, 65, 240, 61, 60, 26, 27, 67,
	28, 68, 0, 0, 0, 40, 235, 48, 0, 0,
	49, 41, 42, 0, 58, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 57, 0, 0, 59, 62, 63,
	0, 0, 0, 44, 45, 46, 47, 0, 0, 203,
	204, 0, 26, 27, 67, 28, 68, 0, 0, 64,
	40, 65, 48, 61, 60, 49, 41, 42, 0, 58,
	0, 50, 0, 0, 0, 0, 0, 0, 0, 57,
	0, 0, 59, 62, 63, 0, 0, 0, 44, 45,
	46, 47, 0, 0, 203, 204, 0, 26, 27, 67,
	28, 68, 224, 0, 64, 40, 65, 48, 61, 60,
	49, 41, 42, 0, 58, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 57, 0, 0, 59, 62, 63,
	0, 0, 0, 44, 45, 46, 47, 0, 0, 0,
	223, 0, 26, 27, 67, 28, 68, 0, 0, 64,
	40, 65, 48, 61, 60, 49, 41, 42, 0, 58,
	0, 50, 0, 0, 0, 0, 0, 0, 0, 57,
	0, 0, 59, 62, 63, 0, 0, 0, 44, 45,
	46, 47, 26, 97, 67, 98, 227, 0, 0, 103,
	0, 0, 0, 0, 64, 0, 65, 0, 61, 60,
	26, 97, 67, 98, 88, 0, 0, 103, 0, 0,
	0, 0, 0, 62, 63, 26, 97, 67, 98, 68,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 62, 63, 0, 64, 0, 104, 0, 61, 60,
	0, 0, 0, 0, 76, 0, 62, 63, 0, 0,
	0, 0, 91, 0, 104, 0, 61, 60, 83, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 0, 65,
	0, 61, 60, 85, 86, 0, 0, 0, 74, 75,
	0, 0, 0, 77, 78, 79, 0, 0, 0, 0,
	0, 73, 82, 80, 81, 84,
}
var RubyPact = []int{

	-27, 1299, -1000, -1000, -1000, -1, -1000, -1000, -1000, 2080,
	-1000, -1000, -1000, 33, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 821, 157, 5,
	3, 1, -1000, -1000, -1000, -1000, -1000, -1000, 144, -20,
	287, 214, 214, 93, 1977, 1977, 1977, 1977, 1977, 1977,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 2050, 1977, 10,
	26, 268, -1000, -1000, 2050, -1000, -17, 203, -1000, -1000,
	-1000, -1000, 1977, 292, 2050, 2050, 2050, 2050, 1977, 2050,
	2050, 2050, 2050, 1977, 2050, 2050, 2050, 1977, 2050, -1000,
	135, 2050, 2050, 291, 154, 52, -1000, 2035, 197, -1000,
	-1000, -1000, -9, -26, -26, 2050, 1977, 1977, 1977, 1977,
	290, 2050, 2050, 1977, 1977, 289, 119, 119, 14, -1000,
	-1000, 1977, 288, 51, 51, 51, 51, 51, 109, 1887,
	100, 52, 87, -1000, -26, -1000, -1000, 225, -1000, -1000,
	16, 13, 522, -1000, 299, 241, 2050, 1932, 51, 2017,
	52, 52, 52, 52, 51, 52, 52, 52, 52, 51,
	166, 52, 52, 51, 279, 522, 218, 501, 52, -1000,
	218, 1842, -1000, 275, -1000, 1784, 255, 51, 51, 51,
	51, -1000, 52, 52, 51, 51, -1000, -1000, 142, 78,
	-1000, 0, 265, 261, -1000, 1739, 214, 1681, 51, -1000,
	1254, -1000, -1000, -1000, -1000, 2080, 51, 188, 2050, -1000,
	-1000, -1000, -1000, 2050, -1000, -1000, -1000, 140, 196, 132,
	-1000, 249, 51, -1000, -1000, 135, -1000, 2050, 2050, -1000,
	52, -1000, -2, 52, -1000, -1000, 1623, 12, -1000, 1565,
	-1000, -1000, -8, 78, 177, 1977, -1000, -1000, -8, -1000,
	-1000, -1000, -1000, 270, 1977, -1000, 932, 1520, -1000, -1000,
	167, 52, 1196, 52, 7, -35, -1000, 1977, 2050, -1000,
	116, 52, 1977, -1000, -1000, 250, -1000, 1404, -1000, -1000,
	51, 1404, 1462, -1000, 1977, -1000, 51, 1887, -1000, 256,
	-1000, 1887, 281, -1000, -1000, -1000, 2080, 51, -1000, -1000,
	1977, 1977, 201, 193, -1000, 1977, -1000, 49, 2080, 51,
	52, -1000, 51, -1000, 160, -1000, -1000, 2080, 51, -1000,
	85, 35, -1000, 2050, 95, -1000, 51, 1887, 1887, -1000,
	1887, 247, 21, 47, 1977, 1977, 1977, 1977, 237, -14,
	-8, -1000, -1000, -1000, 1977, 1977, 100, -1000, 1887, -1000,
	-1000, -1000, -1000, 51, 51, 51, 51, 1977, 2050, 51,
	51, 1887, 1080, 1022, 1138, 227, 97, -1000, 229, 1977,
	-1000, -1000, 43, -1000, -8, -1000, -8, -1000, -1000, 1977,
	-1000, 51, 1359, -1000, -8, -8, 51, 1359, 1359, 1359,
}
var RubyPgo = []int{

	0, 70, 382, 381, 19, 380, 379, 378, 693, 377,
	2, 375, 374, 373, 370, 0, 594, 436, 357, 353,
	351, 8, 350, 4, 391, 347, 6, 345, 344, 342,
	341, 340, 339, 5, 337, 334, 332, 330, 321, 320,
	319, 318, 317, 315, 314, 673, 313, 312, 11, 17,
	7, 310, 13, 309, 15, 303, 16, 1, 302, 3,
	301, 300, 10, 9, 299, 44,
}
var RubyR1 = []int{

	0, 51, 51, 51, 51, 51, 51, 51, 51, 51,
	51, 65, 65, 45, 45, 45, 45, 45, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 21, 21, 21, 21, 21, 21, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 48, 48, 59, 59, 56, 56, 56,
	56, 62, 62, 62, 62, 61, 61, 61, 18, 18,
	27, 27, 27, 27, 57, 57, 57, 57, 57, 57,
	57, 52, 52, 23, 23, 23, 23, 63, 63, 63,
	22, 22, 25, 26, 26, 64, 64, 13, 13, 13,
	13, 13, 13, 8, 8, 24, 24, 16, 16, 34,
	34, 35, 36, 37, 38, 39, 40, 41, 42, 43,
	44, 2, 5, 6, 6, 3, 3, 53, 53, 53,
	53, 60, 60, 60, 4, 4, 4, 4, 49, 58,
	58, 58, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 50, 50, 50, 50, 46, 46, 46, 7,
	14, 10, 10, 10, 55, 55, 47, 47, 19, 20,
	11, 30, 54, 54, 54, 54, 54, 54, 54, 31,
	31, 31, 31, 31, 31, 31, 32, 32, 32, 32,
	32, 33, 33, 33, 33, 29, 28, 9,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 4, 1, 4, 4, 2,
	3, 3, 4, 4, 3, 2, 3, 3, 3, 3,
	3, 4, 6, 3, 1, 1, 3, 0, 1, 1,
	3, 1, 1, 3, 3, 1, 3, 3, 7, 7,
	0, 1, 3, 3, 0, 2, 2, 2, 2, 2,
	3, 1, 3, 1, 2, 3, 2, 0, 1, 3,
	4, 6, 4, 1, 3, 1, 3, 3, 3, 3,
	3, 3, 3, 2, 2, 2, 2, 3, 3, 3,
	3, 2, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 1, 1, 3, 3, 5, 5, 0, 4, 7,
	8, 3, 7, 8, 3, 4, 4, 3, 3, 0,
	1, 3, 4, 5, 3, 3, 3, 3, 3, 5,
	6, 5, 4, 3, 3, 2, 0, 2, 2, 3,
	4, 2, 3, 5, 0, 2, 1, 2, 2, 2,
	5, 5, 0, 2, 2, 2, 2, 2, 2, 0,
	1, 3, 3, 1, 3, 3, 5, 6, 5, 6,
	5, 4, 3, 3, 2, 3, 3, 2,
}
var RubyChk = []int{

	-1000, -51, 47, 48, 64, -1, 47, 48, 64, -15,
	-18, -22, -25, -13, -35, -36, -37, -38, -12, -14,
	-21, -19, -30, -29, -28, -9, 5, 6, 8, -24,
	-16, -8, -2, -5, -6, -3, -26, -11, -17, -7,
	13, 19, 20, -34, 41, 42, 43, 44, 15, 18,
	24, -39, -40, -41, -42, -43, -44, 32, 22, 35,
	62, 61, 36, 37, 57, 59, -64, 7, 9, 48,
	47, 64, 15, 51, 38, 39, 4, 43, 44, 45,
	53, 54, 52, 18, 55, 33, 34, 18, 9, -48,
	-62, 57, 40, 11, -61, -15, -4, 6, 8, -24,
	-16, -8, -17, 12, 59, 9, 40, 40, 40, 40,
	51, 38, 39, 15, 18, 51, 6, 4, -26, 8,
	-26, 40, 11, -1, -1, -1, -1, -1, -1, -45,
	-59, -15, -1, -4, 59, 6, 8, 62, 6, 8,
	-59, -56, -15, -21, -65, 50, 9, -46, -1, 6,
	-15, -15, -15, -15, -1, -15, -15, -15, -15, -1,
	-15, -15, -15, -1, -56, -15, 11, -15, -15, 6,
	11, -45, -49, 52, -49, -45, -56, -1, -1, -1,
	-1, 6, -15, -15, -1, -1, 6, -52, -63, 9,
	-23, 6, 45, 54, -52, -45, 38, -45, -1, 6,
	-45, 47, 48, 47, 48, -15, -1, -55, 11, 47,
	6, 8, 58, 11, 58, 47, -53, -60, -15, 6,
	8, -56, -1, 48, 10, -62, -48, 9, 49, 10,
	-15, -4, 58, -15, -4, 14, -45, -58, 6, -45,
	60, 10, -65, 11, -63, 40, 6, 6, -65, 14,
	-26, 14, 14, -50, 17, 16, -45, -45, 14, -10,
	25, -15, -54, -15, -65, -65, 11, 4, 49, 10,
	-56, -15, 40, 14, 52, 11, 60, -57, -23, 10,
	-1, -57, -45, 14, 17, 16, -1, -45, 14, -50,
	14, -45, 8, 14, 47, 48, -15, -1, -32, -31,
	15, 18, 27, 28, 60, 11, 60, -65, -15, -1,
	-15, 10, -1, 6, -65, 47, 48, -15, -1, -27,
	-47, -20, -10, 31, -65, 14, -1, -45, -45, 14,
	-45, 4, -1, -1, 15, 18, 15, 18, -1, 6,
	-65, 14, 14, -10, 15, 18, -59, 14, -45, 6,
	47, 47, 48, -1, -1, -1, -1, 4, 49, -1,
	-1, -45, -54, -54, -54, -1, -15, 14, -33, 17,
	16, 14, -33, 14, -65, 11, -65, 11, 14, 17,
	16, -1, -54, 14, -65, -65, -1, -54, -54, -54,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 38,
	39, 40, 41, 42, 43, 44, 45, 46, 47, 48,
	49, 50, 51, 52, 53, 54, 18, 19, -2, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	13, 32, 33, 34, 35, 36, 37, 0, 0, 0,
	0, 0, 141, 142, 77, 11, 0, 56, 176, 5,
	6, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, -2, 59,
	65, 77, 0, 0, 74, 81, 82, 19, -2, 21,
	22, 23, 30, 13, -2, 77, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 107, 107, 13, -2,
	13, 0, 0, 131, 132, 133, 134, 13, 0, 184,
	188, 75, 0, 217, 13, 125, 126, 0, 123, 124,
	0, 0, 75, 79, 147, 0, 77, 0, 164, 60,
	66, 68, 70, 135, 136, 137, 138, 139, 140, 166,
	0, 215, 216, 168, 0, 78, 0, 75, 117, 129,
	0, 0, 13, 159, 13, 0, 0, 118, 119, 120,
	121, 61, 67, 69, 165, 167, 64, 11, 101, 107,
	108, 103, 0, 0, 11, 0, 0, 0, 122, 130,
	0, 13, 13, 14, 15, 16, 17, 0, 0, 192,
	127, 128, 143, 0, 144, 12, 11, 11, 0, 19,
	-2, 0, 177, 178, 179, 62, 63, -2, 0, 55,
	83, 84, 71, 86, 87, 154, 0, 0, 160, 0,
	157, 58, 94, 0, 0, 0, 104, 106, 94, 110,
	13, 112, 162, 0, 0, 13, 0, 0, 180, 185,
	13, 76, 0, 80, 0, 0, 11, 0, 0, 57,
	0, 190, 0, 155, 158, 0, 156, 11, 109, 102,
	105, 11, 0, 163, 0, 13, 13, 175, 169, 0,
	171, 181, 13, 191, 193, 194, 38, 196, 197, 198,
	0, 0, 200, 203, 145, 0, 146, 0, 38, 11,
	151, 73, 72, 161, 0, 95, 96, 38, 98, 99,
	0, 91, 186, 0, 0, 111, 13, 173, 174, 170,
	182, 0, 13, 0, 0, 0, 0, 0, 0, 0,
	148, 88, 100, 187, 0, 0, 189, 89, 172, 13,
	192, -2, -2, 201, 202, 204, 205, 0, 0, 92,
	93, 183, 0, 0, 0, 11, 11, 206, 0, 0,
	192, 208, 0, 210, 149, 11, 152, 11, 207, 0,
	192, 192, 199, 209, 150, 153, 192, 199, 199, 199,
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
	62, 63, 64,
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
		//line parser.y:184
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:186
		{
		}
	case 3:
		//line parser.y:188
		{
		}
	case 4:
		//line parser.y:190
		{
		}
	case 5:
		//line parser.y:192
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:194
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:196
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:202
		{
		}
	case 11:
		//line parser.y:204
		{
		}
	case 12:
		//line parser.y:205
		{
		}
	case 13:
		//line parser.y:208
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 14:
		//line parser.y:210
		{
		}
	case 15:
		//line parser.y:212
		{
		}
	case 16:
		//line parser.y:214
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 17:
		//line parser.y:216
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 55:
		//line parser.y:226
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 56:
		//line parser.y:233
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 57:
		//line parser.y:237
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 58:
		//line parser.y:244
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 59:
		//line parser.y:251
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 60:
		//line parser.y:258
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 61:
		//line parser.y:265
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 62:
		//line parser.y:272
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 63:
		//line parser.y:280
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 64:
		//line parser.y:288
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 65:
		//line parser.y:298
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 66:
		//line parser.y:305
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 67:
		//line parser.y:313
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 68:
		//line parser.y:321
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 69:
		//line parser.y:329
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 70:
		//line parser.y:337
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 71:
		//line parser.y:347
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 72:
		//line parser.y:357
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 73:
		//line parser.y:366
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 74:
		//line parser.y:368
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 75:
		//line parser.y:371
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 76:
		//line parser.y:373
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 77:
		//line parser.y:375
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 78:
		//line parser.y:377
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 79:
		//line parser.y:379
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 80:
		//line parser.y:381
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 81:
		//line parser.y:385
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 82:
		//line parser.y:387
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 83:
		//line parser.y:389
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 84:
		//line parser.y:391
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:394
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 86:
		//line parser.y:396
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 87:
		//line parser.y:398
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 88:
		//line parser.y:404
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 89:
		//line parser.y:412
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 90:
		//line parser.y:420
		{
		}
	case 91:
		//line parser.y:422
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 92:
		//line parser.y:424
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{RubyS[Rubypt-2].genericValue}}
		}
	case 93:
		//line parser.y:426
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{ast.Break{}},
			}
		}
	case 94:
		//line parser.y:434
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 95:
		//line parser.y:436
		{
		}
	case 96:
		//line parser.y:438
		{
		}
	case 97:
		//line parser.y:440
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 98:
		//line parser.y:442
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 99:
		//line parser.y:444
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 100:
		//line parser.y:446
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-1].genericSlice...)
		}
	case 101:
		//line parser.y:449
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 102:
		//line parser.y:451
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 103:
		//line parser.y:454
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 104:
		//line parser.y:456
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 105:
		//line parser.y:458
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 106:
		//line parser.y:460
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 107:
		//line parser.y:462
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 108:
		//line parser.y:464
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 109:
		//line parser.y:468
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 110:
		//line parser.y:473
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 111:
		//line parser.y:480
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 112:
		//line parser.y:490
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 113:
		//line parser.y:499
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 114:
		//line parser.y:505
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 115:
		//line parser.y:513
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 116:
		//line parser.y:517
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 117:
		//line parser.y:522
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 118:
		//line parser.y:529
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 119:
		//line parser.y:536
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 120:
		//line parser.y:543
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 121:
		//line parser.y:550
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 122:
		//line parser.y:557
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 123:
		//line parser.y:565
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 124:
		//line parser.y:567
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 125:
		//line parser.y:570
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 126:
		//line parser.y:572
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 127:
		//line parser.y:575
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 128:
		//line parser.y:577
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 129:
		//line parser.y:580
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 130:
		//line parser.y:582
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 131:
		//line parser.y:584
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 132:
		//line parser.y:585
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 133:
		//line parser.y:586
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 134:
		//line parser.y:587
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 135:
		//line parser.y:590
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 136:
		//line parser.y:599
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 137:
		//line parser.y:608
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 138:
		//line parser.y:617
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 139:
		//line parser.y:626
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 140:
		//line parser.y:635
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 141:
		//line parser.y:643
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 142:
		//line parser.y:644
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 143:
		//line parser.y:646
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 144:
		//line parser.y:648
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 145:
		//line parser.y:651
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 146:
		//line parser.y:659
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 147:
		//line parser.y:667
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 148:
		//line parser.y:669
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 149:
		//line parser.y:676
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 150:
		//line parser.y:683
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 151:
		//line parser.y:691
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 152:
		//line parser.y:698
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 153:
		//line parser.y:705
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 154:
		//line parser.y:713
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 155:
		//line parser.y:715
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 156:
		//line parser.y:722
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 157:
		//line parser.y:729
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 158:
		//line parser.y:732
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 159:
		//line parser.y:734
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 160:
		//line parser.y:736
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 161:
		//line parser.y:738
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 162:
		//line parser.y:741
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 163:
		//line parser.y:748
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 164:
		//line parser.y:756
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 165:
		//line parser.y:763
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 166:
		//line parser.y:770
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 167:
		//line parser.y:777
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 168:
		//line parser.y:784
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 169:
		//line parser.y:791
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 170:
		//line parser.y:798
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 171:
		//line parser.y:806
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 172:
		//line parser.y:815
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 173:
		//line parser.y:822
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 174:
		//line parser.y:829
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 175:
		//line parser.y:836
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 176:
		//line parser.y:843
		{
		}
	case 177:
		//line parser.y:844
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 178:
		//line parser.y:845
		{
		}
	case 179:
		//line parser.y:848
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 180:
		//line parser.y:851
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 181:
		//line parser.y:859
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 182:
		//line parser.y:861
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 183:
		//line parser.y:870
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}

			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Var:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
					Class: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				},
			}
		}
	case 184:
		//line parser.y:885
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 185:
		//line parser.y:887
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 186:
		//line parser.y:890
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 187:
		//line parser.y:892
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 188:
		//line parser.y:895
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 189:
		//line parser.y:904
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 190:
		//line parser.y:913
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 191:
		//line parser.y:922
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 192:
		//line parser.y:925
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 193:
		//line parser.y:927
		{
		}
	case 194:
		//line parser.y:929
		{
		}
	case 195:
		//line parser.y:931
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 196:
		//line parser.y:933
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 197:
		//line parser.y:935
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 198:
		//line parser.y:937
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 199:
		//line parser.y:939
		{
		}
	case 200:
		//line parser.y:941
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 201:
		//line parser.y:943
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 202:
		//line parser.y:945
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 203:
		//line parser.y:947
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 204:
		//line parser.y:949
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 205:
		//line parser.y:951
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 206:
		//line parser.y:954
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 207:
		//line parser.y:961
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 208:
		//line parser.y:969
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 209:
		//line parser.y:976
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 210:
		//line parser.y:984
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 211:
		//line parser.y:992
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 212:
		//line parser.y:999
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 213:
		//line parser.y:1006
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 214:
		//line parser.y:1013
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 215:
		//line parser.y:1021
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 216:
		//line parser.y:1024
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 217:
		//line parser.y:1026
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	}
	goto Rubystack /* stack new state and value */
}
