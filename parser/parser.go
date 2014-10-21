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

//line parser.y:1059

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 28,
	52, 121,
	-2, 20,
	-1, 90,
	10, 81,
	11, 81,
	-2, 182,
	-1, 100,
	52, 121,
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
	52, 121,
	-2, 119,
	-1, 224,
	52, 122,
	-2, 120,
	-1, 231,
	10, 81,
	11, 81,
	-2, 182,
	-1, 377,
	27, 198,
	28, 198,
	-2, 13,
	-1, 378,
	27, 198,
	28, 198,
	-2, 13,
}

const RubyNprod = 228
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 2408

var RubyAct = []int{

	9, 290, 399, 141, 267, 344, 36, 258, 92, 193,
	98, 142, 191, 91, 190, 177, 324, 74, 2, 3,
	78, 105, 217, 219, 219, 211, 176, 408, 97, 288,
	117, 285, 136, 20, 137, 4, 325, 146, 115, 387,
	219, 116, 284, 269, 74, 250, 111, 199, 89, 120,
	122, 72, 71, 211, 219, 219, 76, 77, 169, 133,
	78, 135, 132, 113, 114, 273, 143, 323, 73, 75,
	150, 218, 287, 86, 216, 194, 112, 153, 154, 155,
	156, 149, 158, 159, 160, 161, 78, 163, 164, 165,
	138, 168, 113, 114, 170, 171, 76, 77, 373, 144,
	97, 274, 167, 107, 110, 112, 364, 232, 168, 75,
	74, 388, 367, 86, 185, 186, 195, 194, 74, 179,
	192, 175, 76, 77, 144, 196, 107, 144, 74, 124,
	74, 109, 208, 219, 197, 75, 108, 76, 77, 86,
	209, 144, 5, 248, 377, 378, 222, 219, 168, 219,
	75, 173, 376, 97, 86, 317, 147, 318, 195, 225,
	123, 229, 204, 205, 212, 368, 230, 196, 414, 263,
	411, 410, 386, 406, 238, 208, 265, 74, 319, 208,
	265, 144, 370, 74, 239, 371, 125, 126, 127, 128,
	129, 130, 409, 305, 411, 410, 330, 234, 121, 208,
	134, 208, 224, 375, 208, 249, 255, 358, 356, 353,
	359, 357, 266, 366, 335, 151, 264, 351, 67, 298,
	297, 157, 292, 248, 97, 296, 162, 298, 297, 214,
	166, 215, 168, 278, 334, 145, 281, 276, 217, 246,
	217, 332, 208, 277, 279, 208, 282, 233, 234, 180,
	181, 182, 183, 252, 251, 294, 187, 188, 291, 139,
	243, 140, 208, 208, 201, 144, 202, 189, 309, 302,
	133, 321, 119, 320, 118, 327, 329, 184, 172, 152,
	96, 321, 221, 26, 99, 68, 100, 69, 242, 226,
	105, 339, 210, 220, 1, 339, 208, 342, 148, 57,
	56, 208, 55, 54, 322, 208, 53, 52, 17, 16,
	15, 14, 44, 70, 322, 311, 63, 64, 312, 22,
	133, 23, 24, 361, 25, 268, 341, 219, 12, 11,
	343, 21, 10, 19, 283, 13, 18, 65, 37, 106,
	40, 62, 61, 39, 34, 33, 133, 35, 369, 372,
	208, 208, 32, 208, 0, 0, 0, 0, 0, 0,
	0, 208, 0, 208, 0, 0, 0, 0, 0, 0,
	0, 213, 0, 0, 0, 208, 0, 0, 0, 0,
	0, 392, 393, 394, 0, 208, 0, 0, 397, 0,
	0, 293, 208, 309, 309, 309, 403, 0, 0, 0,
	299, 0, 0, 0, 0, 235, 413, 0, 310, 237,
	0, 0, 0, 0, 309, 328, 418, 419, 0, 309,
	309, 309, 420, 0, 0, 333, 247, 0, 0, 0,
	0, 340, 0, 253, 0, 340, 0, 0, 348, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 26, 99,
	68, 100, 69, 270, 354, 355, 271, 272, 0, 0,
	0, 0, 0, 0, 0, 363, 0, 0, 0, 0,
	280, 0, 0, 0, 0, 0, 0, 0, 70, 0,
	0, 63, 64, 0, 26, 99, 68, 100, 90, 0,
	95, 105, 219, 0, 0, 0, 0, 379, 380, 381,
	382, 0, 65, 0, 66, 0, 62, 61, 29, 326,
	0, 389, 390, 0, 70, 0, 0, 63, 64, 0,
	0, 94, 0, 0, 0, 0, 336, 396, 0, 0,
	346, 0, 0, 310, 310, 310, 101, 0, 93, 0,
	106, 412, 62, 61, 0, 0, 0, 0, 0, 0,
	0, 417, 0, 0, 310, 0, 0, 0, 0, 310,
	310, 310, 0, 0, 365, 0, 0, 101, 0, 101,
	0, 0, 0, 0, 101, 0, 0, 0, 0, 0,
	26, 99, 68, 100, 69, 101, 101, 101, 101, 38,
	101, 101, 101, 101, 0, 101, 101, 101, 385, 101,
	0, 0, 101, 101, 0, 0, 0, 0, 101, 0,
	70, 0, 0, 63, 64, 0, 101, 104, 0, 0,
	395, 0, 101, 101, 219, 0, 0, 0, 0, 0,
	0, 331, 405, 407, 65, 0, 66, 0, 62, 61,
	0, 0, 415, 0, 416, 0, 0, 0, 104, 0,
	104, 0, 0, 0, 101, 104, 101, 0, 0, 0,
	0, 101, 0, 0, 0, 0, 104, 104, 104, 104,
	0, 104, 104, 104, 104, 0, 104, 104, 104, 0,
	104, 0, 101, 104, 104, 0, 0, 0, 0, 104,
	0, 0, 26, 99, 68, 100, 90, 104, 0, 105,
	0, 0, 0, 104, 104, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 30, 78, 0, 0, 0,
	101, 0, 70, 0, 0, 63, 64, 0, 0, 0,
	0, 0, 101, 0, 0, 104, 0, 104, 275, 0,
	101, 101, 104, 102, 101, 0, 93, 0, 106, 0,
	62, 61, 76, 77, 0, 0, 0, 79, 80, 81,
	0, 0, 0, 104, 0, 75, 84, 82, 83, 86,
	0, 0, 236, 0, 102, 0, 102, 0, 101, 101,
	0, 102, 0, 0, 101, 0, 0, 0, 0, 101,
	0, 0, 102, 102, 102, 102, 0, 102, 102, 102,
	102, 104, 102, 102, 102, 0, 102, 0, 0, 102,
	102, 0, 0, 104, 0, 102, 0, 0, 0, 0,
	0, 104, 104, 102, 0, 104, 0, 0, 101, 102,
	102, 0, 0, 0, 0, 0, 0, 0, 0, 26,
	99, 68, 100, 69, 0, 0, 105, 0, 0, 0,
	0, 0, 0, 0, 101, 0, 0, 0, 0, 104,
	104, 102, 0, 102, 0, 104, 0, 0, 102, 70,
	104, 0, 63, 64, 0, 31, 0, 0, 178, 0,
	0, 0, 0, 0, 26, 99, 68, 100, 231, 102,
	0, 105, 0, 65, 0, 106, 101, 62, 61, 0,
	0, 0, 0, 103, 0, 0, 0, 0, 0, 104,
	0, 0, 0, 0, 70, 0, 0, 63, 64, 0,
	0, 0, 0, 0, 0, 0, 0, 102, 0, 0,
	131, 0, 0, 0, 103, 104, 103, 0, 65, 102,
	106, 103, 62, 61, 0, 0, 0, 102, 102, 0,
	0, 102, 103, 103, 103, 103, 0, 103, 103, 103,
	103, 0, 103, 103, 103, 0, 103, 0, 0, 103,
	103, 0, 0, 0, 0, 103, 0, 104, 0, 0,
	0, 0, 0, 103, 174, 102, 102, 0, 0, 103,
	103, 102, 0, 0, 0, 0, 102, 0, 0, 198,
	0, 200, 0, 0, 0, 0, 0, 0, 203, 0,
	0, 0, 26, 27, 68, 28, 69, 228, 0, 0,
	41, 103, 49, 103, 0, 50, 42, 43, 103, 59,
	0, 51, 0, 0, 0, 102, 0, 0, 0, 58,
	0, 0, 70, 60, 0, 63, 64, 0, 0, 103,
	45, 46, 47, 48, 241, 0, 244, 227, 0, 0,
	0, 102, 0, 0, 0, 0, 65, 0, 66, 0,
	62, 61, 26, 223, 68, 100, 69, 0, 0, 0,
	0, 0, 0, 261, 262, 0, 0, 103, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 103,
	0, 0, 70, 102, 0, 63, 64, 103, 103, 0,
	0, 103, 0, 0, 0, 0, 219, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 65, 0, 66, 0,
	62, 61, 0, 0, 295, 0, 0, 0, 0, 300,
	0, 0, 0, 0, 304, 103, 103, 0, 0, 0,
	0, 103, 0, 0, 0, 0, 103, 26, 99, 68,
	100, 90, 0, 0, 105, 0, 0, 0, 0, 0,
	26, 99, 68, 100, 69, 0, 0, 349, 350, 0,
	0, 0, 0, 0, 352, 0, 0, 70, 0, 0,
	63, 64, 0, 0, 78, 103, 0, 360, 0, 362,
	70, 0, 0, 63, 64, 0, 0, 0, 0, 0,
	0, 93, 0, 106, 0, 62, 61, 0, 0, 0,
	0, 103, 0, 0, 65, 0, 66, 374, 62, 61,
	76, 77, 0, 203, 0, 79, 80, 81, 0, 0,
	384, 0, 0, 75, 84, 82, 83, 86, 0, 0,
	0, 0, 0, 0, 391, 0, 261, 262, 26, 27,
	68, 28, 69, 103, 0, 0, 41, 402, 313, 401,
	400, 314, 42, 43, 0, 59, 0, 51, 0, 0,
	315, 316, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 307, 308, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 66, 0, 62, 61, 26, 27,
	68, 28, 69, 0, 0, 0, 41, 398, 313, 401,
	400, 314, 42, 43, 0, 59, 0, 51, 0, 0,
	315, 316, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 307, 308, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 66, 0, 62, 61, 26, 27,
	68, 28, 69, 0, 0, 0, 41, 404, 313, 0,
	0, 314, 42, 43, 0, 59, 0, 51, 0, 0,
	315, 316, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 307, 308, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 66, 0, 62, 61, 26, 27,
	68, 28, 69, 0, 0, 0, 41, 306, 313, 0,
	0, 314, 42, 43, 0, 59, 0, 51, 0, 0,
	315, 316, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 307, 308, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 66, 0, 62, 61, 26, 27,
	68, 28, 69, 0, 0, 0, 41, 301, 49, 260,
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
	28, 69, 0, 0, 0, 41, 0, 313, 0, 0,
	314, 42, 43, 0, 59, 0, 51, 0, 0, 315,
	316, 0, 0, 0, 58, 0, 0, 70, 60, 0,
	63, 64, 0, 0, 0, 45, 46, 47, 48, 0,
	0, 307, 308, 0, 26, 27, 68, 28, 69, 0,
	0, 65, 41, 66, 49, 62, 61, 50, 42, 43,
	0, 59, 0, 51, 265, 0, 0, 0, 0, 0,
	345, 58, 0, 0, 70, 60, 0, 63, 64, 0,
	0, 0, 45, 46, 47, 48, 0, 0, 337, 338,
	0, 0, 0, 0, 0, 0, 0, 0, 65, 0,
	66, 0, 62, 61, 26, 27, 68, 28, 69, 0,
	0, 0, 41, 383, 49, 0, 0, 50, 42, 43,
	0, 59, 0, 51, 0, 0, 0, 0, 0, 0,
	0, 58, 0, 0, 70, 60, 0, 63, 64, 0,
	0, 0, 45, 46, 47, 48, 0, 0, 206, 207,
	0, 0, 0, 0, 0, 0, 0, 0, 65, 0,
	66, 0, 62, 61, 26, 27, 68, 28, 69, 0,
	0, 0, 41, 347, 49, 0, 0, 50, 42, 43,
	0, 59, 0, 51, 0, 0, 0, 0, 0, 0,
	0, 58, 0, 0, 70, 60, 0, 63, 64, 0,
	0, 0, 45, 46, 47, 48, 0, 0, 206, 207,
	0, 0, 0, 0, 0, 0, 0, 0, 65, 0,
	66, 0, 62, 61, 26, 27, 68, 28, 69, 0,
	0, 0, 41, 303, 49, 0, 0, 50, 42, 43,
	0, 59, 0, 51, 0, 0, 0, 0, 0, 0,
	0, 58, 0, 0, 70, 60, 0, 63, 64, 0,
	0, 0, 45, 46, 47, 48, 0, 0, 206, 207,
	0, 26, 27, 68, 28, 69, 0, 0, 65, 41,
	66, 49, 62, 61, 50, 42, 43, 0, 59, 0,
	51, 0, 0, 0, 0, 0, 0, 0, 58, 0,
	0, 70, 60, 0, 63, 64, 0, 0, 0, 45,
	46, 47, 48, 0, 0, 206, 207, 0, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 66, 289, 62,
	61, 26, 27, 68, 28, 69, 0, 0, 0, 41,
	286, 49, 0, 0, 50, 42, 43, 0, 59, 0,
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
	0, 0, 70, 60, 78, 63, 64, 0, 0, 0,
	45, 46, 47, 48, 0, 0, 0, 0, 85, 0,
	0, 0, 0, 0, 0, 0, 65, 0, 66, 0,
	62, 61, 0, 87, 88, 0, 0, 0, 0, 0,
	76, 77, 0, 0, 0, 79, 80, 81, 0, 0,
	0, 0, 0, 75, 84, 82, 83, 86,
}
var RubyPact = []int{

	-31, 1600, -1000, -1000, -1000, 2, -1000, -1000, -1000, 2350,
	-1000, -1000, -1000, 30, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 479, 94, 89,
	62, 4, -1000, -1000, -1000, -1000, -1000, -1000, 23, -23,
	-1000, 268, 190, 190, 118, 2317, 2317, 2317, 2317, 2317,
	2317, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 1165, 2317,
	1165, 26, 253, -1000, -1000, 1165, -1000, -15, 147, -1000,
	9, -1000, -1000, -1000, 2317, 273, 1165, 1165, 1165, 1165,
	2317, 1165, 1165, 1165, 1165, 2317, 1165, 1165, 1165, 2317,
	1165, -1000, 47, 1165, 1165, 272, 140, 82, -1000, 1152,
	117, -1000, -1000, -1000, 52, -28, -28, 1165, 2317, 2317,
	2317, 2317, 271, 1165, 1165, 2317, 2317, 261, 111, 111,
	7, -1000, -1000, 2317, 260, 29, 29, 29, 29, 29,
	113, 2270, 42, 82, 115, 82, -1000, -1000, 223, -1000,
	-1000, 14, 11, 1190, -1000, 1067, 194, 1165, 1007, -1000,
	-28, 29, 879, 82, 82, 82, 82, 29, 82, 82,
	82, 82, 29, 56, 82, 82, 29, 237, 1190, -1000,
	712, 82, -1000, 834, 2223, -1000, 254, -1000, 2163, 229,
	29, 29, 29, 29, -1000, 82, 82, 29, 29, -1000,
	-1000, 132, 69, -1000, 3, 248, 247, -1000, 2116, 190,
	2056, 29, -1000, 1553, -1000, -1000, -1000, -1000, 2350, 29,
	155, 1165, -1000, 6, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 54, 97, 687, -1000, 227, 29, -1000, -1000, 47,
	-1000, 1165, 1165, 9, -1000, 278, 0, -25, 82, -1000,
	-1000, 1996, 18, -1000, 1936, -1000, -1000, -9, 69, 212,
	2317, -1000, -1000, -9, -1000, -1000, -1000, -1000, 211, 2317,
	-1000, 1493, 1889, -1000, -1000, 185, 82, 1433, 141, 1165,
	443, 5, -26, -1000, 2317, 1165, -1000, 186, 82, -1000,
	575, 82, -1000, 235, 2317, 228, -1000, -1000, 208, -1000,
	1709, -1000, -1000, 29, 1709, 1829, -1000, 2317, -1000, 29,
	2270, -1000, 203, -1000, 2270, 205, -1000, -1000, -1000, 2350,
	29, -1000, -1000, 2317, 2317, 193, 192, -1000, -1000, 1165,
	42, 1190, -1000, -1000, 2317, -1000, 100, 2350, 29, 82,
	-1000, 207, -1000, 29, -1000, -1000, 98, -1000, -1000, 2350,
	29, -1000, 151, 167, -1000, 1165, 84, -1000, 29, 2270,
	2270, -1000, 2270, 197, 103, 95, 2317, 2317, 2317, 2317,
	1769, 42, 2270, 168, -12, -9, 101, -1000, -1000, -1000,
	2317, 2317, 42, -1000, 2270, -1000, -1000, -1000, -1000, 29,
	29, 29, 29, -1000, 2270, -9, 2317, 1165, -1000, 29,
	29, 2270, 1313, 1253, 1373, -9, 162, 16, -1000, 178,
	2317, -1000, -1000, 154, -1000, -9, -1000, -9, -1000, -1000,
	2317, -1000, 29, 1662, -1000, -9, -9, 29, 1662, 1662,
	1662,
}
var RubyPgo = []int{

	0, 140, 352, 347, 10, 345, 344, 343, 875, 340,
	5, 338, 336, 335, 333, 0, 715, 589, 332, 331,
	330, 33, 329, 9, 508, 328, 6, 326, 325, 324,
	322, 321, 319, 318, 315, 2, 312, 311, 310, 309,
	308, 307, 306, 303, 302, 300, 299, 878, 298, 297,
	13, 15, 7, 294, 14, 293, 4, 292, 11, 1,
	288, 3, 282, 280, 8, 12, 218, 235,
}
var RubyR1 = []int{

	0, 53, 53, 53, 53, 53, 53, 53, 53, 53,
	53, 67, 67, 47, 47, 47, 47, 47, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 21, 21, 21, 21, 21, 21, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 50, 50, 50, 50, 61,
	61, 58, 58, 58, 58, 58, 64, 64, 64, 64,
	64, 63, 63, 63, 18, 18, 27, 27, 27, 27,
	59, 59, 59, 59, 59, 59, 59, 54, 54, 23,
	23, 23, 23, 65, 65, 65, 22, 22, 25, 26,
	26, 66, 66, 13, 13, 13, 13, 13, 13, 8,
	8, 24, 24, 16, 16, 36, 36, 37, 38, 39,
	40, 41, 42, 43, 44, 45, 46, 2, 5, 6,
	6, 3, 3, 55, 55, 55, 55, 62, 62, 62,
	4, 4, 4, 4, 51, 60, 60, 60, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 52, 52,
	52, 52, 48, 48, 48, 7, 14, 10, 10, 10,
	57, 57, 49, 49, 19, 20, 11, 32, 56, 56,
	56, 56, 56, 56, 56, 33, 33, 33, 33, 33,
	33, 33, 34, 34, 34, 34, 34, 35, 35, 35,
	35, 31, 30, 9, 29, 29, 28, 28,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 4, 5, 1, 4,
	4, 2, 3, 3, 4, 4, 3, 2, 3, 3,
	3, 3, 3, 4, 6, 3, 7, 1, 5, 1,
	3, 0, 1, 1, 4, 4, 1, 1, 4, 4,
	5, 1, 3, 3, 7, 7, 0, 1, 3, 3,
	0, 2, 2, 2, 2, 2, 3, 1, 3, 1,
	2, 3, 2, 0, 1, 3, 4, 6, 4, 1,
	3, 1, 3, 3, 3, 3, 3, 3, 3, 2,
	2, 2, 2, 3, 3, 3, 3, 2, 2, 2,
	2, 3, 3, 3, 3, 3, 3, 1, 1, 3,
	3, 5, 5, 0, 4, 7, 8, 3, 7, 8,
	3, 4, 4, 3, 3, 0, 1, 3, 4, 5,
	3, 3, 3, 3, 3, 5, 6, 5, 4, 3,
	3, 2, 0, 2, 2, 3, 4, 2, 3, 5,
	0, 2, 1, 2, 2, 2, 5, 5, 0, 2,
	2, 2, 2, 2, 2, 0, 1, 3, 3, 1,
	3, 3, 5, 6, 5, 6, 5, 4, 3, 3,
	2, 3, 3, 2, 5, 7, 4, 5,
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
	-67, -67, -67, 11, 4, 51, 10, -58, -15, -4,
	-67, -15, -4, 56, 42, 56, 14, 54, 11, 62,
	-59, -23, 10, -1, -59, -47, 14, 17, 16, -1,
	-47, 14, -52, 14, -47, 8, 14, 49, 50, -15,
	-1, -34, -33, 15, 18, 27, 28, 14, 16, 37,
	-61, -15, -21, 62, 11, 62, -67, -15, -1, -15,
	10, 56, 6, -1, 6, 6, -67, 49, 50, -15,
	-1, -27, -49, -20, -10, 31, -67, 14, -1, -47,
	-47, 14, -47, 4, -1, -1, 15, 18, 15, 18,
	-47, -61, -47, -1, 6, -67, 6, 14, 14, -10,
	15, 18, -61, 14, -47, 6, 49, 49, 50, -1,
	-1, -1, -1, 14, -47, -67, 4, 51, 10, -1,
	-1, -47, -56, -56, -56, -67, -1, -15, 14, -35,
	17, 16, 14, -35, 14, -67, 11, -67, 11, 14,
	17, 16, -1, -56, 14, -67, -67, -1, -56, -56,
	-56,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 39,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 18, 19, -2, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 13, 33, 34, 35, 36, 37, 38, 0, 0,
	0, 0, 0, 147, 148, 81, 11, 0, 58, 182,
	0, 5, 6, 7, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	-2, 61, 67, 81, 0, 0, 77, 86, 87, 19,
	-2, 21, 22, 23, 30, 13, -2, 81, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 113, 113,
	13, -2, 13, 0, 0, 137, 138, 139, 140, 13,
	0, 190, 194, 79, 0, 11, 131, 132, 0, 129,
	130, 0, 0, 79, 83, 153, 0, 81, 0, 223,
	13, 170, 62, 68, 70, 72, 141, 142, 143, 144,
	145, 146, 172, 0, 221, 222, 174, 0, 82, 11,
	79, 123, 135, 11, 0, 13, 165, 13, 0, 0,
	124, 125, 126, 127, 63, 69, 71, 171, 173, 66,
	11, 107, 113, 114, 109, 0, 0, 11, 0, 0,
	0, 128, 136, 0, 13, 13, 14, 15, 16, 17,
	0, 0, 198, 0, 133, 134, 149, 11, 150, 12,
	11, 11, 0, 19, -2, 0, 183, 184, 185, 64,
	65, -2, 0, 56, 11, 0, 73, 0, 92, 93,
	160, 0, 0, 166, 0, 163, 60, 100, 0, 0,
	0, 110, 112, 100, 116, 13, 118, 168, 0, 0,
	13, 0, 0, 186, 191, 13, 80, 0, 0, 0,
	0, 0, 0, 11, 0, 0, 59, 0, 196, 57,
	0, 88, 89, 0, 0, 0, 161, 164, 0, 162,
	11, 115, 108, 111, 11, 0, 169, 0, 13, 13,
	181, 175, 0, 177, 187, 13, 197, 199, 200, 39,
	202, 203, 204, 0, 0, 206, 209, 224, 13, 0,
	13, 84, 85, 151, 0, 152, 0, 39, 11, 157,
	75, 0, 90, 74, 78, 167, 0, 101, 102, 39,
	104, 105, 0, 97, 192, 0, 0, 117, 13, 179,
	180, 176, 188, 0, 13, 0, 0, 0, 0, 0,
	0, 13, 11, 0, 0, 154, 0, 94, 106, 193,
	0, 0, 195, 95, 178, 13, 198, -2, -2, 207,
	208, 210, 211, 225, 11, 226, 0, 0, 76, 98,
	99, 189, 0, 0, 0, 227, 11, 11, 212, 0,
	0, 198, 214, 0, 216, 155, 11, 158, 11, 213,
	0, 198, 198, 205, 215, 156, 159, 198, 205, 205,
	205,
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
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-2].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 58:
		//line parser.y:246
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
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
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 61:
		//line parser.y:264
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
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
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 64:
		//line parser.y:285
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 65:
		//line parser.y:293
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 66:
		//line parser.y:301
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 67:
		//line parser.y:311
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 68:
		//line parser.y:318
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 69:
		//line parser.y:326
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 70:
		//line parser.y:334
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 71:
		//line parser.y:342
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 72:
		//line parser.y:350
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 73:
		//line parser.y:360
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 74:
		//line parser.y:370
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 75:
		//line parser.y:379
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 76:
		//line parser.y:381
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 77:
		//line parser.y:383
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 78:
		//line parser.y:385
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 79:
		//line parser.y:388
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 80:
		//line parser.y:390
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 81:
		//line parser.y:392
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 82:
		//line parser.y:394
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 83:
		//line parser.y:396
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 84:
		//line parser.y:398
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:400
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 86:
		//line parser.y:404
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 87:
		//line parser.y:406
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 88:
		//line parser.y:408
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 89:
		//line parser.y:410
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-1].genericValue)
		}
	case 90:
		//line parser.y:412
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 91:
		//line parser.y:415
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 92:
		//line parser.y:417
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 93:
		//line parser.y:419
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 94:
		//line parser.y:425
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 95:
		//line parser.y:433
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 96:
		//line parser.y:441
		{
		}
	case 97:
		//line parser.y:443
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 98:
		//line parser.y:445
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{RubyS[Rubypt-2].genericValue}}
		}
	case 99:
		//line parser.y:447
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{ast.Break{}},
			}
		}
	case 100:
		//line parser.y:455
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 101:
		//line parser.y:457
		{
		}
	case 102:
		//line parser.y:459
		{
		}
	case 103:
		//line parser.y:461
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 104:
		//line parser.y:463
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 105:
		//line parser.y:465
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 106:
		//line parser.y:467
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-1].genericSlice...)
		}
	case 107:
		//line parser.y:470
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 108:
		//line parser.y:472
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 109:
		//line parser.y:475
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 110:
		//line parser.y:477
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 111:
		//line parser.y:479
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 112:
		//line parser.y:481
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 113:
		//line parser.y:483
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 114:
		//line parser.y:485
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 115:
		//line parser.y:489
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 116:
		//line parser.y:494
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 117:
		//line parser.y:501
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 118:
		//line parser.y:511
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 119:
		//line parser.y:520
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 120:
		//line parser.y:526
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 121:
		//line parser.y:534
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 122:
		//line parser.y:538
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
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
		//line parser.y:578
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 129:
		//line parser.y:586
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 130:
		//line parser.y:588
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 131:
		//line parser.y:591
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 132:
		//line parser.y:593
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 133:
		//line parser.y:596
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 134:
		//line parser.y:598
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 135:
		//line parser.y:601
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 136:
		//line parser.y:603
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 137:
		//line parser.y:605
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 138:
		//line parser.y:606
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 139:
		//line parser.y:607
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 140:
		//line parser.y:608
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 141:
		//line parser.y:611
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 142:
		//line parser.y:620
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 143:
		//line parser.y:629
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 144:
		//line parser.y:638
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 145:
		//line parser.y:647
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 146:
		//line parser.y:656
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 147:
		//line parser.y:664
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 148:
		//line parser.y:665
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 149:
		//line parser.y:667
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 150:
		//line parser.y:669
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 151:
		//line parser.y:672
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 152:
		//line parser.y:680
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 153:
		//line parser.y:688
		{
			RubyVAL.genericSlice = ast.Nodes{}
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
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 156:
		//line parser.y:704
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 157:
		//line parser.y:712
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 158:
		//line parser.y:719
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 159:
		//line parser.y:726
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 160:
		//line parser.y:734
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
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
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 163:
		//line parser.y:750
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 164:
		//line parser.y:753
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 165:
		//line parser.y:755
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 166:
		//line parser.y:757
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 167:
		//line parser.y:759
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 168:
		//line parser.y:762
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 169:
		//line parser.y:769
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
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
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 172:
		//line parser.y:791
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
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
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 175:
		//line parser.y:812
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 176:
		//line parser.y:819
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 177:
		//line parser.y:827
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 178:
		//line parser.y:836
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 179:
		//line parser.y:843
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 180:
		//line parser.y:850
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 181:
		//line parser.y:857
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 182:
		//line parser.y:864
		{
		}
	case 183:
		//line parser.y:865
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 184:
		//line parser.y:866
		{
		}
	case 185:
		//line parser.y:869
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 186:
		//line parser.y:872
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 187:
		//line parser.y:880
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 188:
		//line parser.y:882
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 189:
		//line parser.y:891
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
	case 190:
		//line parser.y:906
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 191:
		//line parser.y:908
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 192:
		//line parser.y:911
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 193:
		//line parser.y:913
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 194:
		//line parser.y:916
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 195:
		//line parser.y:925
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 196:
		//line parser.y:934
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 197:
		//line parser.y:943
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 198:
		//line parser.y:946
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 199:
		//line parser.y:948
		{
		}
	case 200:
		//line parser.y:950
		{
		}
	case 201:
		//line parser.y:952
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 202:
		//line parser.y:954
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 203:
		//line parser.y:956
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 204:
		//line parser.y:958
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 205:
		//line parser.y:960
		{
		}
	case 206:
		//line parser.y:962
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 207:
		//line parser.y:964
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 208:
		//line parser.y:966
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 209:
		//line parser.y:968
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 210:
		//line parser.y:970
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 211:
		//line parser.y:972
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 212:
		//line parser.y:975
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 213:
		//line parser.y:982
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 214:
		//line parser.y:990
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 215:
		//line parser.y:997
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 216:
		//line parser.y:1005
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 217:
		//line parser.y:1013
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 218:
		//line parser.y:1020
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 219:
		//line parser.y:1027
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 220:
		//line parser.y:1034
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 221:
		//line parser.y:1042
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 222:
		//line parser.y:1045
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 223:
		//line parser.y:1047
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 224:
		//line parser.y:1050
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 225:
		//line parser.y:1052
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 226:
		//line parser.y:1055
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 227:
		//line parser.y:1057
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	}
	goto Rubystack /* stack new state and value */
}
