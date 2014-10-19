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

//line parser.y:1036

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 27,
	50, 119,
	-2, 20,
	-1, 88,
	10, 79,
	11, 79,
	-2, 180,
	-1, 98,
	50, 119,
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
	41, 13,
	42, 13,
	43, 13,
	44, 13,
	48, 13,
	-2, 11,
	-1, 119,
	50, 119,
	-2, 117,
	-1, 220,
	50, 120,
	-2, 118,
	-1, 227,
	10, 79,
	11, 79,
	-2, 180,
	-1, 360,
	27, 196,
	28, 196,
	-2, 13,
	-1, 361,
	27, 196,
	28, 196,
	-2, 13,
}

const RubyNprod = 222
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 2138

var RubyAct = []int{

	9, 138, 330, 256, 265, 377, 284, 35, 20, 190,
	139, 90, 188, 89, 187, 96, 174, 2, 3, 72,
	312, 282, 215, 133, 103, 134, 213, 95, 173, 111,
	112, 115, 143, 72, 4, 313, 367, 142, 278, 215,
	248, 208, 110, 109, 87, 122, 108, 107, 72, 118,
	120, 70, 69, 384, 196, 208, 215, 72, 131, 130,
	76, 347, 281, 166, 140, 360, 361, 386, 71, 311,
	76, 147, 141, 214, 121, 150, 151, 152, 153, 135,
	155, 156, 157, 158, 146, 160, 161, 162, 212, 165,
	72, 356, 167, 168, 74, 75, 72, 141, 95, 164,
	141, 76, 215, 105, 74, 75, 165, 73, 191, 318,
	230, 84, 182, 183, 141, 228, 176, 73, 350, 270,
	172, 84, 201, 202, 215, 72, 191, 246, 359, 189,
	205, 349, 194, 351, 106, 74, 75, 170, 25, 97,
	66, 98, 67, 218, 263, 165, 271, 192, 73, 261,
	95, 215, 84, 141, 105, 221, 193, 209, 144, 353,
	263, 225, 354, 226, 210, 192, 211, 231, 68, 61,
	62, 236, 205, 175, 193, 392, 205, 389, 388, 299,
	74, 75, 232, 286, 246, 344, 237, 276, 345, 366,
	63, 119, 64, 73, 60, 59, 205, 84, 205, 220,
	72, 205, 247, 206, 253, 5, 273, 213, 342, 264,
	262, 343, 358, 387, 266, 389, 388, 244, 213, 136,
	95, 137, 267, 113, 129, 245, 114, 321, 165, 275,
	319, 266, 251, 229, 230, 117, 141, 116, 274, 267,
	205, 279, 337, 205, 292, 291, 111, 112, 123, 124,
	125, 126, 127, 128, 268, 269, 285, 339, 288, 110,
	205, 205, 132, 296, 277, 290, 303, 292, 291, 250,
	249, 241, 315, 317, 199, 186, 148, 171, 181, 169,
	149, 65, 154, 94, 217, 325, 240, 159, 207, 325,
	205, 163, 195, 216, 197, 205, 1, 328, 145, 205,
	56, 200, 55, 54, 53, 52, 51, 17, 314, 16,
	177, 178, 179, 180, 15, 14, 43, 184, 185, 305,
	306, 22, 322, 23, 24, 198, 332, 327, 12, 11,
	329, 352, 131, 355, 21, 10, 205, 205, 19, 205,
	13, 18, 36, 39, 38, 33, 239, 32, 242, 222,
	34, 31, 0, 0, 348, 0, 0, 0, 205, 0,
	0, 0, 0, 0, 371, 372, 373, 0, 375, 0,
	0, 205, 303, 303, 303, 259, 260, 0, 381, 25,
	97, 66, 98, 88, 391, 93, 103, 25, 219, 66,
	98, 67, 303, 0, 396, 397, 0, 303, 303, 303,
	398, 0, 0, 0, 0, 0, 0, 0, 0, 68,
	61, 62, 383, 385, 92, 0, 0, 68, 61, 62,
	0, 0, 393, 0, 394, 0, 0, 289, 0, 215,
	0, 91, 294, 104, 0, 60, 59, 298, 0, 63,
	0, 64, 0, 60, 59, 0, 25, 97, 66, 98,
	88, 0, 287, 103, 0, 0, 0, 0, 0, 0,
	0, 293, 0, 0, 0, 0, 335, 336, 0, 304,
	0, 0, 0, 338, 0, 316, 68, 61, 62, 0,
	0, 0, 320, 0, 0, 0, 0, 0, 326, 0,
	272, 0, 326, 0, 0, 334, 0, 0, 91, 0,
	104, 0, 60, 59, 76, 0, 0, 0, 357, 0,
	0, 340, 341, 0, 200, 0, 346, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 28,
	0, 0, 370, 0, 259, 260, 0, 0, 74, 75,
	0, 0, 0, 77, 78, 79, 362, 363, 364, 365,
	0, 73, 82, 80, 81, 84, 99, 368, 369, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	374, 0, 0, 0, 0, 304, 304, 304, 0, 0,
	0, 0, 390, 0, 0, 0, 0, 99, 0, 0,
	0, 0, 395, 99, 0, 304, 0, 0, 0, 0,
	304, 304, 304, 0, 99, 99, 99, 99, 0, 99,
	99, 99, 99, 0, 99, 99, 99, 0, 99, 76,
	0, 99, 99, 0, 0, 0, 0, 99, 37, 0,
	0, 0, 0, 0, 0, 99, 0, 0, 0, 0,
	0, 99, 99, 0, 0, 25, 97, 66, 98, 67,
	0, 0, 103, 74, 75, 102, 0, 0, 77, 78,
	79, 0, 0, 0, 0, 0, 73, 82, 80, 81,
	84, 0, 99, 234, 99, 68, 61, 62, 0, 99,
	0, 0, 0, 0, 0, 0, 102, 0, 0, 0,
	0, 0, 102, 0, 235, 0, 99, 63, 0, 104,
	99, 60, 59, 102, 102, 102, 102, 0, 102, 102,
	102, 102, 0, 102, 102, 102, 0, 102, 0, 0,
	102, 102, 0, 0, 0, 0, 102, 0, 29, 0,
	0, 0, 0, 0, 102, 0, 0, 0, 99, 0,
	102, 102, 0, 99, 25, 97, 66, 98, 67, 99,
	0, 103, 0, 0, 0, 100, 0, 99, 99, 0,
	99, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 102, 0, 102, 68, 61, 62, 0, 102, 0,
	0, 0, 0, 0, 0, 0, 100, 0, 0, 0,
	0, 0, 100, 233, 0, 102, 63, 0, 104, 102,
	60, 59, 99, 100, 100, 100, 100, 0, 100, 100,
	100, 100, 0, 100, 100, 100, 0, 100, 0, 0,
	100, 100, 0, 0, 0, 0, 100, 30, 0, 0,
	0, 0, 0, 0, 100, 0, 0, 102, 0, 0,
	100, 100, 102, 0, 0, 0, 0, 0, 102, 0,
	0, 0, 0, 0, 101, 0, 102, 102, 0, 102,
	0, 99, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 100, 0, 100, 0, 0, 0, 0, 100, 0,
	0, 0, 0, 0, 0, 101, 0, 0, 0, 0,
	0, 101, 0, 0, 0, 100, 0, 99, 0, 100,
	0, 102, 101, 101, 101, 101, 0, 101, 101, 101,
	101, 0, 101, 101, 101, 0, 101, 0, 0, 101,
	101, 0, 0, 0, 0, 101, 0, 0, 0, 0,
	0, 0, 0, 101, 0, 0, 0, 100, 0, 101,
	101, 0, 100, 25, 97, 66, 98, 227, 100, 0,
	103, 0, 0, 0, 0, 0, 100, 100, 0, 100,
	102, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	101, 0, 101, 68, 61, 62, 0, 101, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 101, 63, 102, 104, 101, 60,
	59, 100, 0, 0, 0, 0, 0, 25, 26, 66,
	27, 67, 0, 0, 0, 40, 380, 307, 379, 378,
	308, 41, 42, 0, 58, 0, 50, 0, 0, 309,
	310, 0, 0, 0, 57, 0, 101, 68, 61, 62,
	0, 101, 0, 44, 45, 46, 47, 101, 0, 301,
	302, 0, 0, 0, 0, 101, 101, 0, 101, 63,
	100, 64, 0, 60, 59, 0, 0, 25, 26, 66,
	27, 67, 0, 0, 0, 40, 376, 307, 379, 378,
	308, 41, 42, 0, 58, 0, 50, 0, 0, 309,
	310, 0, 0, 0, 57, 0, 100, 68, 61, 62,
	101, 0, 0, 44, 45, 46, 47, 0, 0, 301,
	302, 0, 0, 0, 0, 0, 0, 0, 0, 63,
	0, 64, 0, 60, 59, 25, 26, 66, 27, 67,
	0, 0, 0, 40, 382, 307, 0, 0, 308, 41,
	42, 0, 58, 0, 50, 0, 0, 309, 310, 0,
	0, 0, 57, 0, 0, 68, 61, 62, 0, 101,
	0, 44, 45, 46, 47, 0, 0, 301, 302, 0,
	0, 0, 0, 0, 0, 0, 0, 63, 0, 64,
	0, 60, 59, 25, 26, 66, 27, 67, 0, 0,
	0, 40, 300, 307, 0, 101, 308, 41, 42, 0,
	58, 0, 50, 0, 0, 309, 310, 0, 0, 0,
	57, 0, 0, 68, 61, 62, 0, 0, 0, 44,
	45, 46, 47, 0, 0, 301, 302, 0, 0, 0,
	0, 0, 0, 0, 0, 63, 0, 64, 0, 60,
	59, 25, 26, 66, 27, 67, 0, 0, 0, 40,
	295, 48, 258, 257, 49, 41, 42, 0, 58, 0,
	50, 0, 0, 0, 0, 0, 0, 0, 57, 0,
	0, 68, 61, 62, 0, 0, 0, 44, 45, 46,
	47, 0, 0, 203, 204, 0, 0, 0, 0, 0,
	0, 0, 0, 63, 0, 64, 0, 60, 59, 25,
	26, 66, 27, 67, 0, 0, 0, 40, 255, 48,
	258, 257, 49, 41, 42, 0, 58, 0, 50, 0,
	0, 0, 0, 0, 0, 0, 57, 0, 0, 68,
	61, 62, 0, 0, 0, 44, 45, 46, 47, 0,
	0, 203, 204, 0, 25, 26, 66, 27, 67, 0,
	0, 63, 40, 64, 48, 60, 59, 49, 41, 42,
	0, 58, 0, 50, 0, 0, 0, 0, 0, 0,
	0, 57, 0, 0, 68, 61, 62, 0, 0, 0,
	44, 45, 46, 47, 0, 0, 6, 7, 0, 0,
	0, 0, 0, 0, 0, 0, 63, 0, 64, 0,
	60, 59, 0, 8, 25, 26, 66, 27, 67, 0,
	0, 0, 40, 0, 307, 0, 0, 308, 41, 42,
	0, 58, 0, 50, 0, 0, 309, 310, 0, 0,
	0, 57, 0, 0, 68, 61, 62, 0, 0, 0,
	44, 45, 46, 47, 0, 0, 301, 302, 0, 25,
	26, 66, 27, 67, 0, 0, 63, 40, 64, 48,
	60, 59, 49, 41, 42, 0, 58, 0, 50, 263,
	0, 0, 0, 0, 0, 331, 57, 0, 0, 68,
	61, 62, 0, 0, 0, 44, 45, 46, 47, 0,
	0, 323, 324, 0, 0, 0, 0, 0, 0, 0,
	0, 63, 0, 64, 0, 60, 59, 25, 26, 66,
	27, 67, 0, 0, 0, 40, 333, 48, 0, 0,
	49, 41, 42, 0, 58, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 57, 0, 0, 68, 61, 62,
	0, 0, 0, 44, 45, 46, 47, 0, 0, 203,
	204, 0, 0, 0, 0, 0, 0, 0, 0, 63,
	0, 64, 0, 60, 59, 25, 26, 66, 27, 67,
	0, 0, 0, 40, 297, 48, 0, 0, 49, 41,
	42, 0, 58, 0, 50, 0, 0, 0, 0, 0,
	0, 0, 57, 0, 0, 68, 61, 62, 0, 0,
	0, 44, 45, 46, 47, 0, 0, 203, 204, 0,
	25, 26, 66, 27, 67, 0, 0, 63, 40, 64,
	48, 60, 59, 49, 41, 42, 0, 58, 0, 50,
	0, 0, 0, 0, 0, 0, 0, 57, 0, 0,
	68, 61, 62, 0, 0, 0, 44, 45, 46, 47,
	0, 0, 203, 204, 0, 0, 0, 0, 0, 0,
	0, 0, 63, 0, 64, 283, 60, 59, 25, 26,
	66, 27, 67, 0, 0, 0, 40, 280, 48, 0,
	0, 49, 41, 42, 0, 58, 0, 50, 0, 0,
	0, 0, 0, 0, 0, 57, 0, 0, 68, 61,
	62, 0, 0, 0, 44, 45, 46, 47, 0, 0,
	203, 204, 0, 0, 0, 0, 0, 0, 0, 0,
	63, 0, 64, 0, 60, 59, 25, 26, 66, 27,
	67, 0, 0, 0, 40, 254, 48, 0, 0, 49,
	41, 42, 0, 58, 0, 50, 0, 0, 0, 0,
	0, 0, 0, 57, 0, 0, 68, 61, 62, 0,
	0, 0, 44, 45, 46, 47, 0, 0, 203, 204,
	0, 0, 0, 0, 0, 0, 0, 0, 63, 0,
	64, 0, 60, 59, 25, 26, 66, 27, 67, 0,
	0, 0, 40, 252, 48, 0, 0, 49, 41, 42,
	0, 58, 0, 50, 0, 0, 0, 0, 0, 0,
	0, 57, 0, 0, 68, 61, 62, 0, 0, 0,
	44, 45, 46, 47, 0, 0, 203, 204, 0, 25,
	26, 66, 27, 67, 0, 0, 63, 40, 64, 48,
	60, 59, 49, 41, 42, 0, 58, 0, 50, 0,
	0, 0, 0, 0, 0, 0, 57, 0, 0, 68,
	61, 62, 0, 0, 0, 44, 45, 46, 47, 0,
	0, 203, 204, 0, 0, 0, 0, 0, 0, 0,
	0, 63, 0, 64, 243, 60, 59, 25, 26, 66,
	27, 67, 0, 0, 0, 40, 238, 48, 0, 0,
	49, 41, 42, 0, 58, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 57, 0, 0, 68, 61, 62,
	0, 0, 0, 44, 45, 46, 47, 0, 0, 203,
	204, 0, 25, 26, 66, 27, 67, 0, 0, 63,
	40, 64, 48, 60, 59, 49, 41, 42, 0, 58,
	0, 50, 0, 0, 0, 0, 0, 0, 0, 57,
	0, 0, 68, 61, 62, 0, 0, 0, 44, 45,
	46, 47, 0, 0, 203, 204, 0, 25, 26, 66,
	27, 67, 224, 0, 63, 40, 64, 48, 60, 59,
	49, 41, 42, 0, 58, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 57, 0, 0, 68, 61, 62,
	0, 0, 0, 44, 45, 46, 47, 0, 0, 0,
	223, 0, 25, 26, 66, 27, 67, 0, 0, 63,
	40, 64, 48, 60, 59, 49, 41, 42, 0, 58,
	0, 50, 0, 0, 0, 0, 0, 0, 0, 57,
	0, 0, 68, 61, 62, 0, 76, 0, 44, 45,
	46, 47, 25, 97, 66, 98, 88, 0, 0, 103,
	83, 0, 0, 0, 63, 0, 64, 0, 60, 59,
	25, 97, 66, 98, 67, 85, 86, 0, 0, 0,
	74, 75, 68, 61, 62, 77, 78, 79, 0, 0,
	0, 0, 0, 73, 82, 80, 81, 84, 0, 0,
	68, 61, 62, 0, 91, 0, 104, 0, 60, 59,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 63, 0, 64, 0, 60, 59,
}
var RubyPact = []int{

	-30, 1339, -1000, -1000, -1000, 4, -1000, -1000, -1000, 2052,
	-1000, -1000, -1000, 26, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 374, 94, 7, 6,
	3, -1000, -1000, -1000, -1000, -1000, -1000, 208, -20, -1000,
	231, 183, 183, 34, 2017, 2017, 2017, 2017, 2017, 2017,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 2075, 2017, 17,
	213, -1000, -1000, 2075, -1000, -18, 149, -1000, 12, -1000,
	-1000, -1000, 2017, 274, 2075, 2075, 2075, 2075, 2017, 2075,
	2075, 2075, 2075, 2017, 2075, 2075, 2075, 2017, 2075, -1000,
	52, 2075, 2075, 273, 126, 97, -1000, 2057, 145, -1000,
	-1000, -1000, -9, -24, -24, 2075, 2017, 2017, 2017, 2017,
	272, 2075, 2075, 2017, 2017, 269, 120, 120, 16, -1000,
	-1000, 2017, 268, 33, 33, 33, 33, 33, 75, 1927,
	44, 97, 110, -1000, -1000, 158, -1000, -1000, 30, 15,
	500, -1000, 382, 191, 2075, 1972, -1000, -24, 33, 938,
	97, 97, 97, 97, 33, 97, 97, 97, 97, 33,
	66, 97, 97, 33, 223, 500, 739, 615, 97, -1000,
	640, 1882, -1000, 265, -1000, 1824, 207, 33, 33, 33,
	33, -1000, 97, 97, 33, 33, -1000, -1000, 116, 102,
	-1000, 0, 264, 263, -1000, 1779, 183, 1721, 33, -1000,
	1294, -1000, -1000, -1000, -1000, 2052, 33, 135, 2075, -1000,
	-1000, -1000, -1000, 2075, -1000, -1000, -1000, 108, 142, 441,
	-1000, 196, 33, -1000, -1000, 52, -1000, 2075, 2075, -1000,
	133, 97, -1000, 258, -2, 235, 97, -1000, -1000, 1663,
	10, -1000, 1605, -1000, -1000, -8, 102, 173, 2017, -1000,
	-1000, -8, -1000, -1000, -1000, -1000, 251, 2017, -1000, 1236,
	1560, -1000, -1000, 171, 97, 1178, 500, -1000, 9, -25,
	-1000, 2017, 2075, -1000, 99, 97, 224, -1000, 2017, -1000,
	-1000, -1000, 221, -1000, 1444, -1000, -1000, 33, 1444, 1502,
	-1000, 2017, -1000, 33, 1927, -1000, 228, -1000, 1927, 253,
	-1000, -1000, -1000, 2052, 33, -1000, -1000, 2017, 2017, 193,
	170, -1000, 2017, -1000, 55, 2052, 33, 97, -1000, 121,
	33, -1000, 104, -1000, -1000, 2052, 33, -1000, 119, 144,
	-1000, 2075, 77, -1000, 33, 1927, 1927, -1000, 1927, 206,
	81, 18, 2017, 2017, 2017, 2017, 185, -13, -8, -1000,
	-1000, -1000, -1000, 2017, 2017, 44, -1000, 1927, -1000, -1000,
	-1000, -1000, 33, 33, 33, 33, 2017, 2075, 33, 33,
	1927, 1062, 1002, 1120, 42, 56, -1000, 199, 2017, -1000,
	-1000, 161, -1000, -8, -1000, -8, -1000, -1000, 2017, -1000,
	33, 1399, -1000, -8, -8, 33, 1399, 1399, 1399,
}
var RubyPgo = []int{

	0, 203, 351, 350, 15, 347, 345, 344, 827, 343,
	2, 342, 341, 340, 338, 0, 728, 628, 335, 334,
	330, 8, 329, 9, 529, 328, 7, 327, 324, 323,
	321, 320, 319, 5, 316, 315, 314, 309, 307, 306,
	305, 304, 303, 302, 300, 173, 298, 297, 13, 16,
	3, 296, 14, 293, 4, 288, 10, 6, 286, 1,
	284, 283, 11, 12, 281, 37,
}
var RubyR1 = []int{

	0, 51, 51, 51, 51, 51, 51, 51, 51, 51,
	51, 65, 65, 45, 45, 45, 45, 45, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 21, 21, 21, 21, 21, 21, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 48, 48, 48, 48, 59, 59, 56,
	56, 56, 56, 56, 62, 62, 62, 62, 62, 61,
	61, 61, 18, 18, 27, 27, 27, 27, 57, 57,
	57, 57, 57, 57, 57, 52, 52, 23, 23, 23,
	23, 63, 63, 63, 22, 22, 25, 26, 26, 64,
	64, 13, 13, 13, 13, 13, 13, 8, 8, 24,
	24, 16, 16, 34, 34, 35, 36, 37, 38, 39,
	40, 41, 42, 43, 44, 2, 5, 6, 6, 3,
	3, 53, 53, 53, 53, 60, 60, 60, 4, 4,
	4, 4, 49, 58, 58, 58, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 50, 50, 50, 50,
	46, 46, 46, 7, 14, 10, 10, 10, 55, 55,
	47, 47, 19, 20, 11, 30, 54, 54, 54, 54,
	54, 54, 54, 31, 31, 31, 31, 31, 31, 31,
	32, 32, 32, 32, 32, 33, 33, 33, 33, 29,
	28, 9,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 4, 1, 4, 4, 2,
	3, 3, 4, 4, 3, 2, 3, 3, 3, 3,
	3, 4, 6, 3, 6, 1, 4, 1, 3, 0,
	1, 1, 3, 3, 1, 1, 3, 3, 4, 1,
	3, 3, 7, 7, 0, 1, 3, 3, 0, 2,
	2, 2, 2, 2, 3, 1, 3, 1, 2, 3,
	2, 0, 1, 3, 4, 6, 4, 1, 3, 1,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 2,
	2, 3, 3, 3, 3, 2, 2, 2, 2, 3,
	3, 3, 3, 3, 3, 1, 1, 3, 3, 5,
	5, 0, 4, 7, 8, 3, 7, 8, 3, 4,
	4, 3, 3, 0, 1, 3, 4, 5, 3, 3,
	3, 3, 3, 5, 6, 5, 4, 3, 3, 2,
	0, 2, 2, 3, 4, 2, 3, 5, 0, 2,
	1, 2, 2, 2, 5, 5, 0, 2, 2, 2,
	2, 2, 2, 0, 1, 3, 3, 1, 3, 3,
	5, 6, 5, 6, 5, 4, 3, 3, 2, 3,
	3, 2,
}
var RubyChk = []int{

	-1000, -51, 47, 48, 64, -1, 47, 48, 64, -15,
	-18, -22, -25, -13, -35, -36, -37, -38, -12, -14,
	-21, -19, -30, -29, -28, 5, 6, 8, -24, -16,
	-8, -2, -5, -6, -3, -26, -11, -17, -7, -9,
	13, 19, 20, -34, 41, 42, 43, 44, 15, 18,
	24, -39, -40, -41, -42, -43, -44, 32, 22, 62,
	61, 36, 37, 57, 59, -64, 7, 9, 35, 48,
	47, 64, 15, 51, 38, 39, 4, 43, 44, 45,
	53, 54, 52, 18, 55, 33, 34, 18, 9, -48,
	-62, 57, 40, 11, -61, -15, -4, 6, 8, -24,
	-16, -8, -17, 12, 59, 9, 40, 40, 40, 40,
	51, 38, 39, 15, 18, 51, 6, 4, -26, 8,
	-26, 40, 11, -1, -1, -1, -1, -1, -1, -45,
	-59, -15, -1, 6, 8, 62, 6, 8, -59, -56,
	-15, -21, -65, 50, 9, -46, -4, 59, -1, 6,
	-15, -15, -15, -15, -1, -15, -15, -15, -15, -1,
	-15, -15, -15, -1, -56, -15, 11, -15, -15, 6,
	11, -45, -49, 52, -49, -45, -56, -1, -1, -1,
	-1, 6, -15, -15, -1, -1, 6, -52, -63, 9,
	-23, 6, 45, 54, -52, -45, 38, -45, -1, 6,
	-45, 47, 48, 47, 48, -15, -1, -55, 11, 47,
	6, 8, 58, 11, 58, 47, -53, -60, -15, 6,
	8, -56, -1, 48, 10, -62, -48, 9, 49, 10,
	11, -15, -4, 54, 58, 54, -15, -4, 14, -45,
	-58, 6, -45, 60, 10, -65, 11, -63, 40, 6,
	6, -65, 14, -26, 14, 14, -50, 17, 16, -45,
	-45, 14, -10, 25, -15, -54, -15, -21, -65, -65,
	11, 4, 49, 10, -56, -15, 54, 6, 40, 6,
	14, 52, 11, 60, -57, -23, 10, -1, -57, -45,
	14, 17, 16, -1, -45, 14, -50, 14, -45, 8,
	14, 47, 48, -15, -1, -32, -31, 15, 18, 27,
	28, 60, 11, 60, -65, -15, -1, -15, 10, 6,
	-1, 6, -65, 47, 48, -15, -1, -27, -47, -20,
	-10, 31, -65, 14, -1, -45, -45, 14, -45, 4,
	-1, -1, 15, 18, 15, 18, -1, 6, -65, 10,
	14, 14, -10, 15, 18, -59, 14, -45, 6, 47,
	47, 48, -1, -1, -1, -1, 4, 49, -1, -1,
	-45, -54, -54, -54, -1, -15, 14, -33, 17, 16,
	14, -33, 14, -65, 11, -65, 11, 14, 17, 16,
	-1, -54, 14, -65, -65, -1, -54, -54, -54,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 39,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 18, 19, -2, 21, 22,
	23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	13, 33, 34, 35, 36, 37, 38, 0, 0, 0,
	0, 145, 146, 79, 11, 0, 56, 180, 0, 5,
	6, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, -2, 59,
	65, 79, 0, 0, 75, 84, 85, 19, -2, 21,
	22, 23, 30, 13, -2, 79, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 111, 111, 13, -2,
	13, 0, 0, 135, 136, 137, 138, 13, 0, 188,
	192, 77, 0, 129, 130, 0, 127, 128, 0, 0,
	77, 81, 151, 0, 79, 0, 221, 13, 168, 60,
	66, 68, 70, 139, 140, 141, 142, 143, 144, 170,
	0, 219, 220, 172, 0, 80, 0, 77, 121, 133,
	0, 0, 13, 163, 13, 0, 0, 122, 123, 124,
	125, 61, 67, 69, 169, 171, 64, 11, 105, 111,
	112, 107, 0, 0, 11, 0, 0, 0, 126, 134,
	0, 13, 13, 14, 15, 16, 17, 0, 0, 196,
	131, 132, 147, 0, 148, 12, 11, 11, 0, 19,
	-2, 0, 181, 182, 183, 62, 63, -2, 0, 55,
	0, 86, 87, 0, 71, 0, 90, 91, 158, 0,
	0, 164, 0, 161, 58, 98, 0, 0, 0, 108,
	110, 98, 114, 13, 116, 166, 0, 0, 13, 0,
	0, 184, 189, 13, 78, 0, 82, 83, 0, 0,
	11, 0, 0, 57, 0, 194, 0, 88, 0, 76,
	159, 162, 0, 160, 11, 113, 106, 109, 11, 0,
	167, 0, 13, 13, 179, 173, 0, 175, 185, 13,
	195, 197, 198, 39, 200, 201, 202, 0, 0, 204,
	207, 149, 0, 150, 0, 39, 11, 155, 73, 0,
	72, 165, 0, 99, 100, 39, 102, 103, 0, 95,
	190, 0, 0, 115, 13, 177, 178, 174, 186, 0,
	13, 0, 0, 0, 0, 0, 0, 0, 152, 74,
	92, 104, 191, 0, 0, 193, 93, 176, 13, 196,
	-2, -2, 205, 206, 208, 209, 0, 0, 96, 97,
	187, 0, 0, 0, 11, 11, 210, 0, 0, 196,
	212, 0, 214, 153, 11, 156, 11, 211, 0, 196,
	196, 203, 213, 154, 157, 196, 203, 203, 203,
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
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 75:
		//line parser.y:370
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 76:
		//line parser.y:372
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-3].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 77:
		//line parser.y:375
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 78:
		//line parser.y:377
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 79:
		//line parser.y:379
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 80:
		//line parser.y:381
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 81:
		//line parser.y:383
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 82:
		//line parser.y:385
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 83:
		//line parser.y:387
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 84:
		//line parser.y:391
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:393
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 86:
		//line parser.y:395
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 87:
		//line parser.y:397
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 88:
		//line parser.y:399
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 89:
		//line parser.y:402
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 90:
		//line parser.y:404
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 91:
		//line parser.y:406
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 92:
		//line parser.y:412
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 93:
		//line parser.y:420
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 94:
		//line parser.y:428
		{
		}
	case 95:
		//line parser.y:430
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 96:
		//line parser.y:432
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{RubyS[Rubypt-2].genericValue}}
		}
	case 97:
		//line parser.y:434
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{ast.Break{}},
			}
		}
	case 98:
		//line parser.y:442
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 99:
		//line parser.y:444
		{
		}
	case 100:
		//line parser.y:446
		{
		}
	case 101:
		//line parser.y:448
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 102:
		//line parser.y:450
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 103:
		//line parser.y:452
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 104:
		//line parser.y:454
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-1].genericSlice...)
		}
	case 105:
		//line parser.y:457
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 106:
		//line parser.y:459
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 107:
		//line parser.y:462
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 108:
		//line parser.y:464
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 109:
		//line parser.y:466
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 110:
		//line parser.y:468
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 111:
		//line parser.y:470
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 112:
		//line parser.y:472
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 113:
		//line parser.y:476
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 114:
		//line parser.y:481
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 115:
		//line parser.y:488
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 116:
		//line parser.y:498
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 117:
		//line parser.y:507
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 118:
		//line parser.y:513
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 119:
		//line parser.y:521
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 120:
		//line parser.y:525
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 121:
		//line parser.y:530
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 122:
		//line parser.y:537
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 123:
		//line parser.y:544
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 124:
		//line parser.y:551
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 125:
		//line parser.y:558
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 126:
		//line parser.y:565
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 127:
		//line parser.y:573
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 128:
		//line parser.y:575
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 129:
		//line parser.y:578
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 130:
		//line parser.y:580
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 131:
		//line parser.y:583
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 132:
		//line parser.y:585
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 133:
		//line parser.y:588
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 134:
		//line parser.y:590
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 135:
		//line parser.y:592
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 136:
		//line parser.y:593
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 137:
		//line parser.y:594
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 138:
		//line parser.y:595
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 139:
		//line parser.y:598
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 140:
		//line parser.y:607
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 141:
		//line parser.y:616
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 142:
		//line parser.y:625
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 143:
		//line parser.y:634
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 144:
		//line parser.y:643
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 145:
		//line parser.y:651
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 146:
		//line parser.y:652
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 147:
		//line parser.y:654
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 148:
		//line parser.y:656
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 149:
		//line parser.y:659
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 150:
		//line parser.y:667
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 151:
		//line parser.y:675
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 152:
		//line parser.y:677
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 153:
		//line parser.y:684
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 154:
		//line parser.y:691
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 155:
		//line parser.y:699
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 156:
		//line parser.y:706
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 157:
		//line parser.y:713
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 158:
		//line parser.y:721
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 159:
		//line parser.y:723
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 160:
		//line parser.y:730
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 161:
		//line parser.y:737
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 162:
		//line parser.y:740
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 163:
		//line parser.y:742
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 164:
		//line parser.y:744
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 165:
		//line parser.y:746
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 166:
		//line parser.y:749
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 167:
		//line parser.y:756
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 168:
		//line parser.y:764
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 169:
		//line parser.y:771
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 170:
		//line parser.y:778
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 171:
		//line parser.y:785
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 172:
		//line parser.y:792
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 173:
		//line parser.y:799
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 174:
		//line parser.y:806
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 175:
		//line parser.y:814
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 176:
		//line parser.y:823
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 177:
		//line parser.y:830
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 178:
		//line parser.y:837
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 179:
		//line parser.y:844
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 180:
		//line parser.y:851
		{
		}
	case 181:
		//line parser.y:852
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 182:
		//line parser.y:853
		{
		}
	case 183:
		//line parser.y:856
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 184:
		//line parser.y:859
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 185:
		//line parser.y:867
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 186:
		//line parser.y:869
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 187:
		//line parser.y:878
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
	case 188:
		//line parser.y:893
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 189:
		//line parser.y:895
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 190:
		//line parser.y:898
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 191:
		//line parser.y:900
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 192:
		//line parser.y:903
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 193:
		//line parser.y:912
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 194:
		//line parser.y:921
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 195:
		//line parser.y:930
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 196:
		//line parser.y:933
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 197:
		//line parser.y:935
		{
		}
	case 198:
		//line parser.y:937
		{
		}
	case 199:
		//line parser.y:939
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 200:
		//line parser.y:941
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 201:
		//line parser.y:943
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 202:
		//line parser.y:945
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 203:
		//line parser.y:947
		{
		}
	case 204:
		//line parser.y:949
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 205:
		//line parser.y:951
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 206:
		//line parser.y:953
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 207:
		//line parser.y:955
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 208:
		//line parser.y:957
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 209:
		//line parser.y:959
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 210:
		//line parser.y:962
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 211:
		//line parser.y:969
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 212:
		//line parser.y:977
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 213:
		//line parser.y:984
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 214:
		//line parser.y:992
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 215:
		//line parser.y:1000
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 216:
		//line parser.y:1007
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 217:
		//line parser.y:1014
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 218:
		//line parser.y:1021
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 219:
		//line parser.y:1029
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 220:
		//line parser.y:1032
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 221:
		//line parser.y:1034
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	}
	goto Rubystack /* stack new state and value */
}
