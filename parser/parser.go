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
	yys             int
	operator        string
	genericValue    ast.Node
	genericSlice    ast.Nodes
	stringSlice     []string
	switchCaseSlice []ast.SwitchCase
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
const CASE = 57378
const WHEN = 57379
const TRUE = 57380
const FALSE = 57381
const LESSTHAN = 57382
const GREATERTHAN = 57383
const EQUALTO = 57384
const BANG = 57385
const COMPLEMENT = 57386
const POSITIVE = 57387
const NEGATIVE = 57388
const STAR = 57389
const WHITESPACE = 57390
const NEWLINE = 57391
const SEMICOLON = 57392
const COLON = 57393
const DOUBLECOLON = 57394
const DOT = 57395
const PIPE = 57396
const SLASH = 57397
const AMPERSAND = 57398
const QUESTIONMARK = 57399
const CARET = 57400
const LBRACKET = 57401
const RBRACKET = 57402
const LBRACE = 57403
const RBRACE = 57404
const DOLLARSIGN = 57405
const ATSIGN = 57406
const FILE_CONST_REF = 57407
const EOF = 57408

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
	"CASE",
	"WHEN",
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

//line parser.y:1052

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 28,
	52, 120,
	-2, 20,
	-1, 90,
	10, 80,
	11, 80,
	-2, 181,
	-1, 100,
	52, 120,
	-2, 20,
	-1, 106,
	13, 13,
	15, 13,
	18, 13,
	19, 13,
	20, 13,
	22, 13,
	24, 13,
	32, 13,
	36, 13,
	43, 13,
	44, 13,
	45, 13,
	46, 13,
	50, 13,
	-2, 11,
	-1, 121,
	52, 120,
	-2, 118,
	-1, 224,
	52, 121,
	-2, 119,
	-1, 231,
	10, 80,
	11, 80,
	-2, 181,
	-1, 376,
	27, 197,
	28, 197,
	-2, 13,
	-1, 377,
	27, 197,
	28, 197,
	-2, 13,
}

const RubyNprod = 227
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 2412

var RubyAct = []int{

	9, 193, 343, 398, 267, 258, 289, 36, 142, 141,
	98, 92, 191, 91, 190, 177, 2, 3, 74, 176,
	146, 78, 219, 219, 117, 219, 105, 217, 97, 386,
	284, 269, 20, 4, 136, 324, 137, 211, 283, 250,
	111, 113, 114, 219, 110, 199, 109, 89, 74, 211,
	120, 122, 72, 71, 112, 329, 234, 76, 77, 133,
	78, 135, 79, 80, 81, 169, 143, 145, 132, 73,
	75, 84, 82, 83, 86, 150, 218, 153, 154, 155,
	156, 149, 158, 159, 160, 161, 216, 163, 164, 165,
	323, 168, 138, 74, 170, 171, 76, 77, 144, 167,
	97, 194, 273, 26, 99, 68, 100, 69, 168, 75,
	105, 107, 287, 86, 185, 186, 179, 194, 124, 369,
	192, 175, 370, 144, 248, 363, 144, 375, 219, 74,
	372, 387, 208, 70, 197, 74, 63, 64, 173, 405,
	144, 322, 195, 74, 108, 78, 222, 219, 168, 123,
	366, 196, 407, 97, 282, 286, 225, 65, 195, 106,
	367, 62, 61, 212, 229, 219, 230, 196, 219, 376,
	377, 265, 107, 274, 238, 208, 263, 374, 78, 208,
	144, 76, 77, 74, 239, 219, 413, 265, 410, 409,
	214, 385, 215, 316, 75, 317, 304, 357, 86, 208,
	358, 208, 74, 213, 208, 249, 209, 255, 5, 76,
	77, 147, 266, 264, 76, 77, 318, 204, 205, 79,
	80, 81, 75, 139, 97, 140, 86, 75, 84, 82,
	83, 86, 168, 278, 236, 355, 280, 235, 356, 121,
	277, 237, 208, 224, 365, 208, 281, 119, 115, 118,
	290, 116, 125, 126, 127, 128, 129, 130, 247, 334,
	293, 333, 208, 208, 144, 253, 134, 301, 308, 331,
	133, 320, 252, 113, 114, 326, 328, 291, 248, 319,
	320, 151, 276, 217, 78, 270, 112, 157, 271, 272,
	338, 251, 162, 243, 338, 208, 166, 67, 85, 202,
	208, 189, 279, 321, 208, 408, 184, 410, 409, 246,
	217, 172, 321, 87, 88, 180, 181, 182, 183, 133,
	76, 77, 187, 188, 152, 79, 80, 81, 360, 350,
	201, 297, 296, 75, 84, 82, 83, 86, 233, 234,
	295, 325, 297, 296, 368, 133, 352, 96, 221, 208,
	208, 242, 208, 210, 371, 226, 220, 335, 1, 341,
	208, 345, 208, 148, 57, 56, 55, 54, 53, 52,
	17, 16, 15, 14, 208, 78, 44, 310, 311, 22,
	391, 392, 393, 23, 208, 24, 25, 396, 268, 340,
	12, 208, 308, 308, 308, 364, 402, 11, 342, 21,
	26, 99, 68, 100, 90, 412, 95, 105, 10, 19,
	13, 76, 77, 308, 18, 417, 418, 37, 308, 308,
	308, 419, 232, 40, 75, 39, 34, 29, 86, 384,
	70, 33, 35, 63, 64, 32, 0, 94, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 394, 0, 0, 93, 101, 106, 292, 62, 61,
	0, 0, 0, 404, 406, 0, 298, 0, 0, 0,
	0, 0, 0, 414, 309, 415, 0, 0, 0, 0,
	0, 327, 0, 0, 0, 0, 101, 0, 101, 0,
	332, 0, 0, 101, 0, 0, 339, 0, 0, 0,
	339, 0, 0, 347, 101, 101, 101, 101, 0, 101,
	101, 101, 101, 0, 101, 101, 101, 0, 101, 353,
	354, 101, 101, 0, 0, 0, 0, 101, 0, 0,
	362, 0, 0, 0, 0, 101, 0, 0, 0, 0,
	0, 101, 101, 0, 0, 0, 0, 0, 0, 0,
	0, 26, 99, 68, 100, 69, 0, 0, 0, 0,
	0, 0, 378, 379, 380, 381, 0, 0, 0, 0,
	0, 0, 0, 101, 0, 101, 388, 389, 0, 0,
	101, 70, 0, 0, 63, 64, 0, 38, 0, 0,
	0, 0, 395, 0, 0, 219, 0, 0, 309, 309,
	309, 101, 330, 0, 0, 65, 411, 66, 0, 62,
	61, 0, 0, 0, 0, 104, 416, 0, 0, 309,
	0, 0, 0, 0, 309, 309, 309, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 101,
	0, 0, 0, 0, 0, 0, 104, 0, 104, 0,
	0, 101, 0, 104, 0, 0, 0, 0, 0, 101,
	101, 0, 0, 101, 104, 104, 104, 104, 0, 104,
	104, 104, 104, 0, 104, 104, 104, 0, 104, 0,
	0, 104, 104, 0, 0, 0, 30, 104, 0, 0,
	0, 0, 0, 0, 0, 104, 0, 101, 101, 0,
	0, 104, 104, 101, 0, 0, 0, 101, 0, 0,
	0, 0, 0, 0, 102, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 104, 0, 104, 0, 0, 0, 0,
	104, 0, 0, 0, 0, 102, 101, 102, 0, 0,
	0, 0, 102, 0, 26, 99, 68, 100, 69, 0,
	0, 104, 0, 102, 102, 102, 102, 0, 102, 102,
	102, 102, 101, 102, 102, 102, 0, 102, 0, 0,
	102, 102, 0, 0, 70, 0, 102, 63, 64, 0,
	0, 0, 0, 0, 102, 0, 0, 0, 219, 104,
	102, 102, 0, 0, 0, 0, 0, 0, 65, 0,
	66, 104, 62, 61, 101, 0, 0, 0, 0, 104,
	104, 0, 0, 104, 0, 0, 0, 0, 0, 0,
	0, 0, 102, 0, 102, 0, 0, 0, 0, 102,
	0, 0, 0, 0, 0, 0, 31, 0, 0, 178,
	0, 0, 0, 0, 0, 0, 0, 104, 104, 0,
	102, 0, 0, 104, 0, 0, 0, 104, 0, 0,
	0, 0, 0, 0, 103, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 102, 0,
	0, 131, 0, 0, 0, 103, 104, 103, 0, 0,
	102, 0, 103, 0, 0, 0, 0, 0, 102, 102,
	0, 0, 102, 103, 103, 103, 103, 0, 103, 103,
	103, 103, 104, 103, 103, 103, 0, 103, 0, 0,
	103, 103, 0, 0, 0, 0, 103, 0, 0, 0,
	0, 0, 0, 0, 103, 174, 102, 102, 0, 0,
	103, 103, 102, 0, 0, 0, 102, 0, 0, 0,
	198, 0, 200, 0, 104, 0, 0, 0, 0, 203,
	0, 0, 0, 26, 27, 68, 28, 69, 228, 0,
	0, 41, 103, 49, 103, 0, 50, 42, 43, 103,
	59, 0, 51, 0, 0, 102, 0, 0, 0, 0,
	58, 0, 0, 70, 60, 0, 63, 64, 0, 0,
	103, 45, 46, 47, 48, 241, 0, 244, 227, 0,
	0, 102, 0, 0, 0, 0, 0, 65, 0, 66,
	0, 62, 61, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 261, 262, 0, 0, 103, 0,
	0, 0, 0, 0, 0, 26, 99, 68, 100, 90,
	103, 0, 105, 102, 0, 0, 0, 0, 103, 103,
	0, 0, 103, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 70, 0, 0, 63, 64,
	0, 0, 0, 0, 0, 294, 0, 0, 0, 0,
	299, 275, 0, 0, 0, 303, 103, 103, 0, 93,
	0, 106, 103, 62, 61, 0, 103, 26, 99, 68,
	100, 69, 0, 0, 105, 0, 0, 0, 0, 0,
	0, 0, 26, 99, 68, 100, 231, 348, 349, 105,
	0, 0, 0, 0, 351, 0, 0, 70, 0, 0,
	63, 64, 0, 0, 0, 103, 0, 359, 0, 361,
	0, 0, 70, 0, 0, 63, 64, 0, 0, 0,
	0, 65, 0, 106, 0, 62, 61, 0, 0, 0,
	0, 103, 0, 0, 0, 0, 65, 373, 106, 0,
	62, 61, 0, 203, 0, 0, 0, 0, 0, 0,
	383, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 390, 0, 261, 262, 26, 27,
	68, 28, 69, 103, 0, 0, 41, 401, 312, 400,
	399, 313, 42, 43, 0, 59, 0, 51, 0, 0,
	314, 315, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 306, 307, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 66, 0, 62, 61, 26, 27,
	68, 28, 69, 0, 0, 0, 41, 397, 312, 400,
	399, 313, 42, 43, 0, 59, 0, 51, 0, 0,
	314, 315, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 306, 307, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 66, 0, 62, 61, 26, 27,
	68, 28, 69, 0, 0, 0, 41, 403, 312, 0,
	0, 313, 42, 43, 0, 59, 0, 51, 0, 0,
	314, 315, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 306, 307, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 66, 0, 62, 61, 26, 27,
	68, 28, 69, 0, 0, 0, 41, 305, 312, 0,
	0, 313, 42, 43, 0, 59, 0, 51, 0, 0,
	314, 315, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 306, 307, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 66, 0, 62, 61, 26, 27,
	68, 28, 69, 0, 0, 0, 41, 300, 49, 260,
	259, 50, 42, 43, 0, 59, 0, 51, 0, 0,
	0, 0, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 206, 207, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 66, 0, 62, 61, 26, 27,
	68, 28, 69, 0, 0, 0, 41, 257, 49, 260,
	259, 50, 42, 43, 0, 59, 0, 51, 0, 0,
	0, 0, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 206, 207, 0, 26, 27, 68, 28, 69,
	0, 0, 65, 41, 66, 49, 62, 61, 50, 42,
	43, 0, 59, 0, 51, 0, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 70, 60, 0, 63, 64,
	0, 0, 0, 45, 46, 47, 48, 0, 0, 6,
	7, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	0, 66, 0, 62, 61, 0, 8, 26, 27, 68,
	28, 69, 0, 0, 0, 41, 0, 312, 0, 0,
	313, 42, 43, 0, 59, 0, 51, 0, 0, 314,
	315, 0, 0, 0, 58, 0, 0, 70, 60, 0,
	63, 64, 0, 0, 0, 45, 46, 47, 48, 0,
	0, 306, 307, 0, 26, 27, 68, 28, 69, 0,
	0, 65, 41, 66, 49, 62, 61, 50, 42, 43,
	0, 59, 0, 51, 265, 0, 0, 0, 0, 0,
	344, 58, 0, 0, 70, 60, 0, 63, 64, 0,
	0, 0, 45, 46, 47, 48, 0, 0, 336, 337,
	0, 0, 0, 0, 0, 0, 0, 0, 65, 0,
	66, 0, 62, 61, 26, 27, 68, 28, 69, 0,
	0, 0, 41, 382, 49, 0, 0, 50, 42, 43,
	0, 59, 0, 51, 0, 0, 0, 0, 0, 0,
	0, 58, 0, 0, 70, 60, 0, 63, 64, 0,
	0, 0, 45, 46, 47, 48, 0, 0, 206, 207,
	0, 0, 0, 0, 0, 0, 0, 0, 65, 0,
	66, 0, 62, 61, 26, 27, 68, 28, 69, 0,
	0, 0, 41, 346, 49, 0, 0, 50, 42, 43,
	0, 59, 0, 51, 0, 0, 0, 0, 0, 0,
	0, 58, 0, 0, 70, 60, 0, 63, 64, 0,
	0, 0, 45, 46, 47, 48, 0, 0, 206, 207,
	0, 0, 0, 0, 0, 0, 0, 0, 65, 0,
	66, 0, 62, 61, 26, 27, 68, 28, 69, 0,
	0, 0, 41, 302, 49, 0, 0, 50, 42, 43,
	0, 59, 0, 51, 0, 0, 0, 0, 0, 0,
	0, 58, 0, 0, 70, 60, 0, 63, 64, 0,
	0, 0, 45, 46, 47, 48, 0, 0, 206, 207,
	0, 26, 27, 68, 28, 69, 0, 0, 65, 41,
	66, 49, 62, 61, 50, 42, 43, 0, 59, 0,
	51, 0, 0, 0, 0, 0, 0, 0, 58, 0,
	0, 70, 60, 0, 63, 64, 0, 0, 0, 45,
	46, 47, 48, 0, 0, 206, 207, 0, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 66, 288, 62,
	61, 26, 27, 68, 28, 69, 0, 0, 0, 41,
	285, 49, 0, 0, 50, 42, 43, 0, 59, 0,
	51, 0, 0, 0, 0, 0, 0, 0, 58, 0,
	0, 70, 60, 0, 63, 64, 0, 0, 0, 45,
	46, 47, 48, 0, 0, 206, 207, 0, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 66, 0, 62,
	61, 26, 27, 68, 28, 69, 0, 0, 0, 41,
	256, 49, 0, 0, 50, 42, 43, 0, 59, 0,
	51, 0, 0, 0, 0, 0, 0, 0, 58, 0,
	0, 70, 60, 0, 63, 64, 0, 0, 0, 45,
	46, 47, 48, 0, 0, 206, 207, 0, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 66, 0, 62,
	61, 26, 27, 68, 28, 69, 0, 0, 0, 41,
	254, 49, 0, 0, 50, 42, 43, 0, 59, 0,
	51, 0, 0, 0, 0, 0, 0, 0, 58, 0,
	0, 70, 60, 0, 63, 64, 0, 0, 0, 45,
	46, 47, 48, 0, 0, 206, 207, 0, 26, 27,
	68, 28, 69, 0, 0, 65, 41, 66, 49, 62,
	61, 50, 42, 43, 0, 59, 0, 51, 0, 0,
	0, 0, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 206, 207, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 66, 245, 62, 61, 26, 27,
	68, 28, 69, 0, 0, 0, 41, 240, 49, 0,
	0, 50, 42, 43, 0, 59, 0, 51, 0, 0,
	0, 0, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 206, 207, 0, 26, 27, 68, 28, 69,
	0, 0, 65, 41, 66, 49, 62, 61, 50, 42,
	43, 0, 59, 0, 51, 0, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 70, 60, 0, 63, 64,
	0, 0, 0, 45, 46, 47, 48, 0, 0, 206,
	207, 0, 26, 27, 68, 28, 69, 0, 0, 65,
	41, 66, 49, 62, 61, 50, 42, 43, 0, 59,
	0, 51, 0, 0, 0, 0, 0, 0, 0, 58,
	0, 0, 70, 60, 0, 63, 64, 0, 0, 0,
	45, 46, 47, 48, 26, 223, 68, 100, 69, 0,
	26, 99, 68, 100, 90, 0, 65, 105, 66, 0,
	62, 61, 26, 99, 68, 100, 69, 0, 0, 0,
	0, 0, 0, 0, 70, 0, 0, 63, 64, 0,
	70, 0, 0, 63, 64, 0, 0, 0, 219, 0,
	0, 0, 70, 0, 0, 63, 64, 0, 65, 0,
	66, 0, 62, 61, 93, 0, 106, 0, 62, 61,
	0, 0, 0, 0, 0, 0, 65, 0, 66, 0,
	62, 61,
}
var RubyPact = []int{

	-33, 1570, -1000, -1000, -1000, 3, -1000, -1000, -1000, 280,
	-1000, -1000, -1000, 29, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 395, 102, 4,
	2, -2, -1000, -1000, -1000, -1000, -1000, -1000, 233, -29,
	-1000, 243, 231, 231, 107, 2287, 2287, 2287, 2287, 2287,
	2287, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 2347, 2287,
	2347, 28, 217, -1000, -1000, 2347, -1000, -32, 202, -1000,
	14, -1000, -1000, -1000, 2287, 318, 2347, 2347, 2347, 2347,
	2287, 2347, 2347, 2347, 2347, 2287, 2347, 2347, 2347, 2287,
	2347, -1000, 54, 2347, 2347, 305, 127, 56, -1000, 2335,
	163, -1000, -1000, -1000, 1, -35, -35, 2347, 2287, 2287,
	2287, 2287, 300, 2347, 2347, 2287, 2287, 295, 111, 111,
	5, -1000, -1000, 2287, 293, 33, 33, 33, 33, 33,
	168, 2240, 38, 56, 114, 56, -1000, -1000, 184, -1000,
	-1000, 26, 16, 17, -1000, 2329, 235, 2347, 978, -1000,
	-35, 33, 1137, 56, 56, 56, 56, 33, 56, 56,
	56, 56, 33, 371, 56, 56, 33, 328, 17, -1000,
	174, 56, -1000, 1122, 2193, -1000, 287, -1000, 2133, 299,
	33, 33, 33, 33, -1000, 56, 56, 33, 33, -1000,
	-1000, 113, 95, -1000, -3, 285, 266, -1000, 2086, 231,
	2026, 33, -1000, 1523, -1000, -1000, -1000, -1000, 280, 33,
	162, 2347, -1000, -6, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 91, 169, 1060, -1000, 272, 33, -1000, -1000, 54,
	-1000, 2347, 2347, -1000, -1000, 98, -4, -26, 56, -1000,
	-1000, 1966, 101, -1000, 1906, -1000, -1000, -24, 95, 267,
	2287, -1000, -1000, -24, -1000, -1000, -1000, -1000, 326, 2287,
	-1000, 1463, 1859, -1000, -1000, 188, 56, 1403, 179, 2347,
	749, 79, -27, -1000, 2287, 2347, -1000, 45, 56, 546,
	56, -1000, 263, 2287, 255, -1000, -1000, 253, -1000, 1679,
	-1000, -1000, 33, 1679, 1799, -1000, 2287, -1000, 33, 2240,
	-1000, 315, -1000, 2240, 342, -1000, -1000, -1000, 280, 33,
	-1000, -1000, 2287, 2287, 220, 182, -1000, -1000, 2347, 38,
	17, -1000, -1000, 2287, -1000, 119, 280, 33, 56, -1000,
	238, -1000, 33, -1000, -1000, 136, -1000, -1000, 280, 33,
	-1000, 146, 104, -1000, 2347, 116, -1000, 33, 2240, 2240,
	-1000, 2240, 171, 78, 120, 2287, 2287, 2287, 2287, 1739,
	38, 2240, 187, -22, -24, 121, -1000, -1000, -1000, 2287,
	2287, 38, -1000, 2240, -1000, -1000, -1000, -1000, 33, 33,
	33, 33, -1000, 2240, -24, 2287, 2347, -1000, 33, 33,
	2240, 1283, 1223, 1343, -24, 128, 141, -1000, 291, 2287,
	-1000, -1000, 172, -1000, -24, -1000, -24, -1000, -1000, 2287,
	-1000, 33, 1632, -1000, -24, -24, 33, 1632, 1632, 1632,
}
var RubyPgo = []int{

	0, 206, 435, 432, 10, 431, 426, 425, 846, 423,
	2, 417, 414, 410, 409, 0, 686, 587, 408, 399,
	398, 32, 397, 1, 427, 390, 7, 389, 388, 386,
	385, 383, 379, 378, 377, 3, 376, 373, 372, 371,
	370, 369, 368, 367, 366, 365, 364, 849, 363, 359,
	13, 15, 5, 358, 14, 356, 4, 353, 8, 6,
	351, 9, 348, 347, 11, 12, 297, 67,
}
var RubyR1 = []int{

	0, 53, 53, 53, 53, 53, 53, 53, 53, 53,
	53, 67, 67, 47, 47, 47, 47, 47, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 21, 21, 21, 21, 21, 21, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 50, 50, 50, 50, 61, 61,
	58, 58, 58, 58, 58, 64, 64, 64, 64, 64,
	63, 63, 63, 18, 18, 27, 27, 27, 27, 59,
	59, 59, 59, 59, 59, 59, 54, 54, 23, 23,
	23, 23, 65, 65, 65, 22, 22, 25, 26, 26,
	66, 66, 13, 13, 13, 13, 13, 13, 8, 8,
	24, 24, 16, 16, 36, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 2, 5, 6, 6,
	3, 3, 55, 55, 55, 55, 62, 62, 62, 4,
	4, 4, 4, 51, 60, 60, 60, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 52, 52, 52,
	52, 48, 48, 48, 7, 14, 10, 10, 10, 57,
	57, 49, 49, 19, 20, 11, 32, 56, 56, 56,
	56, 56, 56, 56, 33, 33, 33, 33, 33, 33,
	33, 34, 34, 34, 34, 34, 35, 35, 35, 35,
	31, 30, 9, 29, 29, 28, 28,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 4, 1, 4, 4,
	2, 3, 3, 4, 4, 3, 2, 3, 3, 3,
	3, 3, 4, 6, 3, 7, 1, 5, 1, 3,
	0, 1, 1, 4, 4, 1, 1, 4, 4, 5,
	1, 3, 3, 7, 7, 0, 1, 3, 3, 0,
	2, 2, 2, 2, 2, 3, 1, 3, 1, 2,
	3, 2, 0, 1, 3, 4, 6, 4, 1, 3,
	1, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	2, 2, 3, 3, 3, 3, 2, 2, 2, 2,
	3, 3, 3, 3, 3, 3, 1, 1, 3, 3,
	5, 5, 0, 4, 7, 8, 3, 7, 8, 3,
	4, 4, 3, 3, 0, 1, 3, 4, 5, 3,
	3, 3, 3, 3, 5, 6, 5, 4, 3, 3,
	2, 0, 2, 2, 3, 4, 2, 3, 5, 0,
	2, 1, 2, 2, 2, 5, 5, 0, 2, 2,
	2, 2, 2, 2, 0, 1, 3, 3, 1, 3,
	3, 5, 6, 5, 6, 5, 4, 3, 3, 2,
	3, 3, 2, 5, 7, 4, 5,
}
var RubyChk = []int{

	-1000, -53, 49, 50, 66, -1, 49, 50, 66, -15,
	-18, -22, -25, -13, -37, -38, -39, -40, -12, -14,
	-21, -19, -32, -31, -30, -29, 5, 6, 8, -24,
	-16, -8, -2, -5, -6, -3, -26, -11, -17, -7,
	-9, 13, 19, 20, -36, 43, 44, 45, 46, 15,
	18, 24, -41, -42, -43, -44, -45, -46, 32, 22,
	36, 64, 63, 38, 39, 59, 61, -66, 7, 9,
	35, 50, 49, 66, 15, 53, 40, 41, 4, 45,
	46, 47, 55, 56, 54, 18, 57, 33, 34, 18,
	9, -50, -64, 59, 42, 11, -63, -15, -4, 6,
	8, -24, -16, -8, -17, 12, 61, 9, 42, 42,
	42, 42, 53, 40, 41, 15, 18, 53, 6, 4,
	-26, 8, -26, 42, 11, -1, -1, -1, -1, -1,
	-1, -47, -61, -15, -1, -15, 6, 8, 64, 6,
	8, -61, -58, -15, -21, -67, 52, 9, -48, -4,
	61, -1, 6, -15, -15, -15, -15, -1, -15, -15,
	-15, -15, -1, -15, -15, -15, -1, -58, -15, 11,
	-15, -15, 6, 11, -47, -51, 54, -51, -47, -58,
	-1, -1, -1, -1, 6, -15, -15, -1, -1, 6,
	-54, -65, 9, -23, 6, 47, 56, -54, -47, 40,
	-47, -1, 6, -47, 49, 50, 49, 50, -15, -1,
	-57, 11, 49, -67, 6, 8, 60, 11, 60, 49,
	-55, -62, -15, 6, 8, -58, -1, 50, 10, -64,
	-50, 9, 51, 10, 11, -67, 60, -67, -15, -4,
	14, -47, -60, 6, -47, 62, 10, -67, 11, -65,
	42, 6, 6, -67, 14, -26, 14, 14, -52, 17,
	16, -47, -47, 14, -10, 25, -15, -56, -28, 37,
	-67, -67, -67, 11, 4, 51, 10, -58, -15, -67,
	-15, -4, 56, 42, 56, 14, 54, 11, 62, -59,
	-23, 10, -1, -59, -47, 14, 17, 16, -1, -47,
	14, -52, 14, -47, 8, 14, 49, 50, -15, -1,
	-34, -33, 15, 18, 27, 28, 14, 16, 37, -61,
	-15, -21, 62, 11, 62, -67, -15, -1, -15, 10,
	56, 6, -1, 6, 6, -67, 49, 50, -15, -1,
	-27, -49, -20, -10, 31, -67, 14, -1, -47, -47,
	14, -47, 4, -1, -1, 15, 18, 15, 18, -47,
	-61, -47, -1, 6, -67, 6, 14, 14, -10, 15,
	18, -61, 14, -47, 6, 49, 49, 50, -1, -1,
	-1, -1, 14, -47, -67, 4, 51, 10, -1, -1,
	-47, -56, -56, -56, -67, -1, -15, 14, -35, 17,
	16, 14, -35, 14, -67, 11, -67, 11, 14, 17,
	16, -1, -56, 14, -67, -67, -1, -56, -56, -56,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 39,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 18, 19, -2, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 13, 33, 34, 35, 36, 37, 38, 0, 0,
	0, 0, 0, 146, 147, 80, 11, 0, 57, 181,
	0, 5, 6, 7, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	-2, 60, 66, 80, 0, 0, 76, 85, 86, 19,
	-2, 21, 22, 23, 30, 13, -2, 80, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 112, 112,
	13, -2, 13, 0, 0, 136, 137, 138, 139, 13,
	0, 189, 193, 78, 0, 11, 130, 131, 0, 128,
	129, 0, 0, 78, 82, 152, 0, 80, 0, 222,
	13, 169, 61, 67, 69, 71, 140, 141, 142, 143,
	144, 145, 171, 0, 220, 221, 173, 0, 81, 11,
	78, 122, 134, 11, 0, 13, 164, 13, 0, 0,
	123, 124, 125, 126, 62, 68, 70, 170, 172, 65,
	11, 106, 112, 113, 108, 0, 0, 11, 0, 0,
	0, 127, 135, 0, 13, 13, 14, 15, 16, 17,
	0, 0, 197, 0, 132, 133, 148, 11, 149, 12,
	11, 11, 0, 19, -2, 0, 182, 183, 184, 63,
	64, -2, 0, 56, 11, 0, 72, 0, 91, 92,
	159, 0, 0, 165, 0, 162, 59, 99, 0, 0,
	0, 109, 111, 99, 115, 13, 117, 167, 0, 0,
	13, 0, 0, 185, 190, 13, 79, 0, 0, 0,
	0, 0, 0, 11, 0, 0, 58, 0, 195, 0,
	87, 88, 0, 0, 0, 160, 163, 0, 161, 11,
	114, 107, 110, 11, 0, 168, 0, 13, 13, 180,
	174, 0, 176, 186, 13, 196, 198, 199, 39, 201,
	202, 203, 0, 0, 205, 208, 223, 13, 0, 13,
	83, 84, 150, 0, 151, 0, 39, 11, 156, 74,
	0, 89, 73, 77, 166, 0, 100, 101, 39, 103,
	104, 0, 96, 191, 0, 0, 116, 13, 178, 179,
	175, 187, 0, 13, 0, 0, 0, 0, 0, 0,
	13, 11, 0, 0, 153, 0, 93, 105, 192, 0,
	0, 194, 94, 177, 13, 197, -2, -2, 206, 207,
	209, 210, 224, 11, 225, 0, 0, 75, 97, 98,
	188, 0, 0, 0, 226, 11, 11, 211, 0, 0,
	197, 213, 0, 215, 154, 11, 157, 11, 212, 0,
	197, 197, 204, 214, 155, 158, 197, 204, 204, 204,
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
	62, 63, 64, 65, 66,
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
		//line parser.y:190
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:192
		{
		}
	case 3:
		//line parser.y:194
		{
		}
	case 4:
		//line parser.y:196
		{
		}
	case 5:
		//line parser.y:198
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:200
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:202
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:208
		{
		}
	case 11:
		//line parser.y:210
		{
		}
	case 12:
		//line parser.y:211
		{
		}
	case 13:
		//line parser.y:214
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 14:
		//line parser.y:216
		{
		}
	case 15:
		//line parser.y:218
		{
		}
	case 16:
		//line parser.y:220
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 17:
		//line parser.y:222
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 56:
		//line parser.y:232
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 57:
		//line parser.y:239
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 58:
		//line parser.y:243
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 59:
		//line parser.y:250
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 60:
		//line parser.y:257
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 61:
		//line parser.y:264
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 62:
		//line parser.y:271
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 63:
		//line parser.y:278
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 64:
		//line parser.y:286
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 65:
		//line parser.y:294
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 66:
		//line parser.y:304
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 67:
		//line parser.y:311
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 68:
		//line parser.y:319
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 69:
		//line parser.y:327
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 70:
		//line parser.y:335
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 71:
		//line parser.y:343
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 72:
		//line parser.y:353
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 73:
		//line parser.y:363
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 74:
		//line parser.y:372
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 75:
		//line parser.y:374
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 76:
		//line parser.y:376
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 77:
		//line parser.y:378
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 78:
		//line parser.y:381
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 79:
		//line parser.y:383
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 80:
		//line parser.y:385
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 81:
		//line parser.y:387
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 82:
		//line parser.y:389
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 83:
		//line parser.y:391
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 84:
		//line parser.y:393
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:397
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 86:
		//line parser.y:399
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 87:
		//line parser.y:401
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 88:
		//line parser.y:403
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-1].genericValue)
		}
	case 89:
		//line parser.y:405
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 90:
		//line parser.y:408
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 91:
		//line parser.y:410
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 92:
		//line parser.y:412
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 93:
		//line parser.y:418
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 94:
		//line parser.y:426
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 95:
		//line parser.y:434
		{
		}
	case 96:
		//line parser.y:436
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 97:
		//line parser.y:438
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{RubyS[Rubypt-2].genericValue}}
		}
	case 98:
		//line parser.y:440
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{ast.Break{}},
			}
		}
	case 99:
		//line parser.y:448
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 100:
		//line parser.y:450
		{
		}
	case 101:
		//line parser.y:452
		{
		}
	case 102:
		//line parser.y:454
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 103:
		//line parser.y:456
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 104:
		//line parser.y:458
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 105:
		//line parser.y:460
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-1].genericSlice...)
		}
	case 106:
		//line parser.y:463
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 107:
		//line parser.y:465
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 108:
		//line parser.y:468
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 109:
		//line parser.y:470
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 110:
		//line parser.y:472
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 111:
		//line parser.y:474
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 112:
		//line parser.y:476
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 113:
		//line parser.y:478
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 114:
		//line parser.y:482
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 115:
		//line parser.y:487
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 116:
		//line parser.y:494
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 117:
		//line parser.y:504
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 118:
		//line parser.y:513
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 119:
		//line parser.y:519
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 120:
		//line parser.y:527
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 121:
		//line parser.y:531
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 122:
		//line parser.y:536
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 123:
		//line parser.y:543
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 124:
		//line parser.y:550
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 125:
		//line parser.y:557
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 126:
		//line parser.y:564
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 127:
		//line parser.y:571
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 128:
		//line parser.y:579
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 129:
		//line parser.y:581
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 130:
		//line parser.y:584
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 131:
		//line parser.y:586
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 132:
		//line parser.y:589
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 133:
		//line parser.y:591
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 134:
		//line parser.y:594
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 135:
		//line parser.y:596
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 136:
		//line parser.y:598
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 137:
		//line parser.y:599
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 138:
		//line parser.y:600
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 139:
		//line parser.y:601
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 140:
		//line parser.y:604
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 141:
		//line parser.y:613
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 142:
		//line parser.y:622
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 143:
		//line parser.y:631
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 144:
		//line parser.y:640
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 145:
		//line parser.y:649
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 146:
		//line parser.y:657
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 147:
		//line parser.y:658
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 148:
		//line parser.y:660
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 149:
		//line parser.y:662
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 150:
		//line parser.y:665
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 151:
		//line parser.y:673
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 152:
		//line parser.y:681
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 153:
		//line parser.y:683
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 154:
		//line parser.y:690
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 155:
		//line parser.y:697
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 156:
		//line parser.y:705
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 157:
		//line parser.y:712
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 158:
		//line parser.y:719
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 159:
		//line parser.y:727
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 160:
		//line parser.y:729
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 161:
		//line parser.y:736
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 162:
		//line parser.y:743
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 163:
		//line parser.y:746
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 164:
		//line parser.y:748
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 165:
		//line parser.y:750
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 166:
		//line parser.y:752
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 167:
		//line parser.y:755
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 168:
		//line parser.y:762
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 169:
		//line parser.y:770
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 170:
		//line parser.y:777
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 171:
		//line parser.y:784
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 172:
		//line parser.y:791
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 173:
		//line parser.y:798
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 174:
		//line parser.y:805
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 175:
		//line parser.y:812
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 176:
		//line parser.y:820
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 177:
		//line parser.y:829
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 178:
		//line parser.y:836
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 179:
		//line parser.y:843
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 180:
		//line parser.y:850
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 181:
		//line parser.y:857
		{
		}
	case 182:
		//line parser.y:858
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 183:
		//line parser.y:859
		{
		}
	case 184:
		//line parser.y:862
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 185:
		//line parser.y:865
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 186:
		//line parser.y:873
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 187:
		//line parser.y:875
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 188:
		//line parser.y:884
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
	case 189:
		//line parser.y:899
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 190:
		//line parser.y:901
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 191:
		//line parser.y:904
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 192:
		//line parser.y:906
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 193:
		//line parser.y:909
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 194:
		//line parser.y:918
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 195:
		//line parser.y:927
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 196:
		//line parser.y:936
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 197:
		//line parser.y:939
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 198:
		//line parser.y:941
		{
		}
	case 199:
		//line parser.y:943
		{
		}
	case 200:
		//line parser.y:945
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 201:
		//line parser.y:947
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 202:
		//line parser.y:949
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 203:
		//line parser.y:951
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 204:
		//line parser.y:953
		{
		}
	case 205:
		//line parser.y:955
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 206:
		//line parser.y:957
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 207:
		//line parser.y:959
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 208:
		//line parser.y:961
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 209:
		//line parser.y:963
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 210:
		//line parser.y:965
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 211:
		//line parser.y:968
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 212:
		//line parser.y:975
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 213:
		//line parser.y:983
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 214:
		//line parser.y:990
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 215:
		//line parser.y:998
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 216:
		//line parser.y:1006
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 217:
		//line parser.y:1013
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 218:
		//line parser.y:1020
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 219:
		//line parser.y:1027
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 220:
		//line parser.y:1035
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 221:
		//line parser.y:1038
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 222:
		//line parser.y:1040
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 223:
		//line parser.y:1043
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 224:
		//line parser.y:1045
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 225:
		//line parser.y:1048
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 226:
		//line parser.y:1050
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	}
	goto Rubystack /* stack new state and value */
}
