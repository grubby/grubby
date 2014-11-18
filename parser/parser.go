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
const RANGE = 57390
const OR_EQUALS = 57391
const WHITESPACE = 57392
const NEWLINE = 57393
const SEMICOLON = 57394
const COLON = 57395
const DOUBLECOLON = 57396
const DOT = 57397
const PIPE = 57398
const SLASH = 57399
const AMPERSAND = 57400
const QUESTIONMARK = 57401
const CARET = 57402
const LBRACKET = 57403
const RBRACKET = 57404
const LBRACE = 57405
const RBRACE = 57406
const DOLLARSIGN = 57407
const ATSIGN = 57408
const FILE_CONST_REF = 57409
const EOF = 57410

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
	"RANGE",
	"OR_EQUALS",
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

//line parser.y:1180

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 42,
	54, 129,
	-2, 21,
	-1, 108,
	54, 129,
	-2, 127,
	-1, 110,
	10, 94,
	11, 94,
	-2, 199,
	-1, 122,
	13, 15,
	15, 15,
	18, 15,
	19, 15,
	20, 15,
	22, 15,
	24, 15,
	31, 15,
	32, 15,
	36, 15,
	52, 15,
	-2, 13,
	-1, 124,
	54, 129,
	-2, 21,
	-1, 255,
	54, 130,
	-2, 128,
	-1, 265,
	10, 94,
	11, 94,
	-2, 199,
	-1, 412,
	27, 216,
	28, 216,
	-2, 15,
	-1, 413,
	27, 216,
	28, 216,
	-2, 15,
}

const RubyNprod = 246
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 3092

var RubyAct = []int{

	240, 297, 5, 322, 306, 433, 192, 194, 148, 196,
	119, 112, 27, 17, 22, 323, 111, 249, 23, 221,
	82, 249, 284, 83, 2, 3, 121, 164, 56, 165,
	371, 88, 131, 249, 369, 246, 207, 121, 149, 242,
	336, 4, 220, 132, 104, 154, 422, 249, 140, 141,
	133, 136, 107, 109, 197, 341, 80, 79, 137, 145,
	97, 98, 131, 88, 339, 335, 143, 86, 87, 147,
	101, 159, 160, 81, 405, 273, 139, 158, 151, 153,
	157, 202, 85, 169, 170, 100, 247, 166, 158, 295,
	245, 177, 97, 98, 308, 198, 182, 242, 9, 86,
	87, 187, 128, 189, 190, 423, 199, 138, 249, 338,
	376, 281, 197, 200, 85, 195, 304, 82, 96, 249,
	83, 134, 206, 208, 151, 88, 313, 151, 135, 212,
	224, 225, 210, 227, 228, 229, 230, 231, 232, 233,
	118, 219, 151, 213, 215, 88, 84, 311, 448, 204,
	445, 444, 442, 198, 97, 98, 144, 257, 146, 144,
	271, 86, 87, 150, 199, 82, 128, 223, 83, 151,
	217, 161, 162, 163, 97, 98, 85, 184, 185, 155,
	96, 86, 87, 171, 262, 173, 174, 175, 176, 263,
	178, 179, 180, 181, 256, 183, 85, 380, 186, 129,
	188, 411, 421, 272, 393, 88, 130, 88, 304, 205,
	377, 439, 209, 211, 214, 278, 326, 271, 302, 351,
	102, 304, 118, 103, 316, 246, 244, 205, 287, 304,
	226, 97, 98, 252, 97, 98, 97, 98, 86, 87,
	108, 86, 87, 86, 87, 101, 294, 246, 89, 90,
	91, 248, 253, 85, 205, 255, 85, 303, 85, 94,
	92, 93, 96, 410, 82, 318, 362, 83, 363, 197,
	386, 118, 195, 208, 317, 82, 88, 324, 83, 151,
	329, 325, 398, 267, 268, 399, 440, 282, 385, 364,
	82, 331, 82, 83, 260, 83, 261, 285, 383, 345,
	412, 413, 348, 291, 319, 97, 98, 355, 280, 281,
	198, 243, 86, 87, 275, 374, 286, 365, 193, 396,
	314, 199, 397, 367, 66, 266, 309, 85, 236, 237,
	310, 312, 443, 88, 445, 444, 384, 391, 378, 344,
	343, 305, 387, 274, 388, 378, 367, 95, 270, 97,
	98, 234, 216, 118, 84, 191, 86, 87, 172, 394,
	395, 332, 97, 98, 205, 320, 167, 117, 168, 86,
	87, 85, 327, 401, 89, 90, 91, 99, 106, 251,
	105, 333, 290, 407, 85, 94, 92, 93, 96, 241,
	250, 370, 342, 372, 344, 343, 1, 414, 415, 416,
	417, 43, 156, 55, 54, 53, 52, 144, 366, 51,
	50, 34, 33, 373, 375, 32, 426, 427, 428, 31,
	46, 356, 430, 378, 357, 19, 36, 355, 355, 355,
	37, 366, 20, 437, 307, 446, 14, 12, 11, 21,
	447, 18, 10, 125, 24, 450, 16, 13, 355, 35,
	451, 452, 355, 355, 355, 453, 15, 30, 29, 125,
	25, 125, 125, 144, 63, 26, 125, 403, 62, 404,
	0, 0, 0, 0, 125, 125, 125, 0, 0, 0,
	0, 403, 420, 0, 0, 0, 125, 0, 125, 125,
	125, 125, 0, 125, 125, 125, 125, 0, 125, 429,
	0, 125, 0, 125, 0, 0, 0, 118, 88, 0,
	0, 441, 125, 0, 0, 125, 125, 125, 0, 0,
	373, 431, 449, 0, 0, 125, 0, 0, 0, 0,
	125, 0, 0, 125, 0, 0, 0, 97, 98, 0,
	0, 0, 0, 0, 86, 87, 0, 0, 0, 89,
	90, 91, 99, 0, 125, 125, 0, 125, 0, 85,
	94, 92, 93, 96, 0, 0, 283, 0, 0, 0,
	0, 0, 0, 0, 125, 61, 41, 67, 42, 68,
	0, 0, 0, 38, 438, 358, 125, 125, 359, 39,
	40, 0, 58, 0, 49, 0, 0, 360, 361, 0,
	0, 60, 57, 0, 28, 69, 59, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 0, 125,
	0, 353, 354, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 0, 65, 0, 76, 75, 0, 0, 0,
	0, 0, 0, 0, 125, 0, 120, 0, 0, 0,
	0, 0, 0, 222, 0, 0, 125, 0, 0, 0,
	0, 0, 120, 0, 120, 120, 0, 125, 125, 120,
	0, 0, 0, 0, 0, 125, 0, 120, 120, 120,
	0, 0, 0, 0, 125, 0, 0, 0, 0, 120,
	0, 120, 120, 120, 120, 0, 120, 120, 120, 120,
	0, 120, 0, 142, 120, 0, 120, 0, 0, 0,
	125, 125, 0, 0, 0, 120, 0, 125, 120, 120,
	120, 0, 0, 0, 0, 0, 0, 0, 120, 0,
	0, 0, 0, 120, 125, 0, 120, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 120, 120, 0,
	120, 201, 0, 203, 0, 0, 125, 0, 0, 0,
	125, 0, 125, 0, 0, 218, 0, 120, 61, 41,
	67, 42, 68, 259, 125, 0, 38, 0, 47, 120,
	120, 48, 39, 40, 235, 58, 0, 49, 0, 0,
	0, 0, 0, 0, 60, 57, 0, 44, 69, 59,
	125, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 0, 120, 0, 125, 258, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 0, 65, 0, 76, 75,
	0, 0, 0, 0, 0, 0, 269, 120, 0, 126,
	0, 0, 0, 0, 276, 0, 0, 0, 0, 120,
	0, 0, 0, 0, 0, 126, 0, 126, 126, 0,
	120, 120, 126, 289, 0, 292, 0, 0, 120, 0,
	126, 126, 126, 0, 0, 0, 0, 120, 0, 0,
	300, 301, 126, 0, 126, 126, 126, 126, 0, 126,
	126, 126, 126, 0, 126, 0, 0, 126, 0, 126,
	0, 0, 0, 120, 120, 0, 0, 0, 126, 0,
	120, 126, 126, 126, 0, 0, 0, 0, 0, 0,
	0, 126, 330, 0, 0, 0, 126, 120, 0, 126,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 346, 0, 0, 0, 0, 350, 0,
	126, 126, 0, 126, 0, 0, 0, 0, 0, 120,
	0, 0, 0, 120, 0, 120, 0, 0, 379, 0,
	126, 0, 0, 0, 0, 0, 0, 120, 0, 0,
	0, 0, 126, 126, 0, 0, 0, 0, 389, 390,
	0, 0, 0, 0, 0, 392, 0, 0, 0, 0,
	45, 0, 0, 120, 0, 0, 0, 400, 0, 402,
	0, 0, 0, 0, 0, 126, 0, 120, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 409, 0, 0, 0, 0, 0, 235, 0,
	126, 0, 127, 0, 0, 419, 0, 0, 0, 0,
	0, 0, 126, 0, 425, 0, 300, 301, 127, 0,
	127, 127, 0, 126, 126, 127, 0, 0, 0, 0,
	0, 126, 0, 127, 127, 127, 0, 0, 0, 0,
	126, 0, 0, 0, 0, 127, 0, 127, 127, 127,
	127, 0, 127, 127, 127, 127, 0, 127, 0, 0,
	127, 0, 127, 0, 0, 0, 126, 126, 0, 0,
	0, 127, 0, 126, 127, 127, 127, 0, 0, 0,
	0, 0, 0, 0, 127, 0, 0, 0, 0, 127,
	126, 0, 127, 0, 0, 0, 0, 61, 123, 67,
	124, 110, 0, 116, 121, 0, 0, 0, 0, 0,
	0, 0, 0, 127, 127, 0, 127, 0, 0, 0,
	0, 0, 126, 0, 0, 0, 126, 69, 126, 0,
	77, 78, 0, 127, 114, 70, 71, 72, 73, 74,
	126, 115, 0, 0, 0, 127, 127, 0, 0, 0,
	0, 0, 0, 113, 0, 122, 0, 76, 75, 0,
	61, 41, 67, 42, 68, 0, 126, 0, 38, 436,
	358, 435, 434, 359, 39, 40, 0, 58, 127, 49,
	126, 0, 360, 361, 0, 0, 60, 57, 0, 0,
	69, 59, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 127, 0, 0, 353, 354, 0, 0,
	0, 0, 0, 0, 0, 127, 64, 0, 65, 0,
	76, 75, 0, 0, 0, 0, 127, 127, 0, 0,
	0, 0, 0, 0, 127, 0, 0, 0, 0, 0,
	0, 0, 0, 127, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 127,
	127, 0, 0, 0, 0, 0, 127, 61, 41, 67,
	42, 68, 0, 0, 0, 38, 432, 358, 435, 434,
	359, 39, 40, 127, 58, 0, 49, 0, 0, 360,
	361, 0, 0, 60, 57, 0, 0, 69, 59, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 353, 354, 127, 0, 0, 0, 127,
	0, 127, 0, 64, 0, 65, 0, 76, 75, 0,
	0, 0, 0, 127, 0, 0, 0, 0, 61, 41,
	67, 42, 68, 0, 0, 0, 38, 352, 358, 0,
	0, 359, 39, 40, 0, 58, 0, 49, 0, 127,
	360, 361, 0, 0, 60, 57, 0, 0, 69, 59,
	0, 77, 78, 127, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 353, 354, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 0, 65, 0, 76, 75,
	61, 41, 67, 42, 68, 0, 0, 0, 38, 347,
	47, 299, 298, 48, 39, 40, 0, 58, 0, 49,
	0, 0, 0, 0, 0, 0, 60, 57, 0, 0,
	69, 59, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 0, 0, 0, 238, 239, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 65, 0,
	76, 75, 61, 41, 67, 42, 68, 0, 0, 0,
	38, 296, 47, 299, 298, 48, 39, 40, 0, 58,
	0, 49, 0, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 238, 239,
	0, 61, 41, 67, 42, 68, 0, 0, 64, 38,
	65, 47, 76, 75, 48, 39, 40, 0, 58, 0,
	49, 0, 0, 0, 0, 0, 0, 60, 57, 0,
	0, 69, 59, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 0, 0, 0, 6, 7, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 0, 65,
	0, 76, 75, 0, 8, 61, 41, 67, 42, 68,
	0, 0, 0, 38, 0, 358, 0, 0, 359, 39,
	40, 0, 58, 0, 49, 0, 0, 360, 361, 0,
	0, 60, 57, 0, 0, 69, 59, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 0, 0,
	0, 353, 354, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 0, 65, 0, 76, 75, 61, 41, 67,
	42, 68, 0, 0, 0, 38, 406, 47, 0, 0,
	48, 39, 40, 0, 58, 0, 49, 304, 0, 0,
	0, 0, 0, 60, 57, 0, 0, 69, 59, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 238, 239, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 0, 65, 0, 76, 75, 61,
	41, 67, 42, 68, 0, 0, 0, 38, 328, 47,
	0, 0, 48, 39, 40, 0, 58, 0, 49, 304,
	0, 0, 0, 0, 0, 60, 57, 0, 0, 69,
	59, 0, 77, 78, 0, 0, 0, 70, 71, 72,
	73, 74, 0, 0, 0, 238, 239, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 0, 65, 0, 76,
	75, 61, 41, 67, 42, 68, 0, 0, 0, 38,
	321, 47, 0, 0, 48, 39, 40, 0, 58, 0,
	49, 304, 0, 0, 0, 0, 0, 60, 57, 0,
	0, 69, 59, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 0, 0, 0, 238, 239, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 0, 65,
	0, 76, 75, 61, 41, 67, 42, 68, 0, 0,
	0, 38, 418, 47, 0, 0, 48, 39, 40, 0,
	58, 0, 49, 0, 0, 0, 0, 0, 0, 60,
	57, 0, 0, 69, 59, 0, 77, 78, 0, 0,
	0, 70, 71, 72, 73, 74, 0, 0, 0, 238,
	239, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	0, 65, 0, 76, 75, 61, 41, 67, 42, 68,
	0, 0, 0, 38, 381, 47, 0, 0, 48, 39,
	40, 0, 58, 0, 49, 0, 0, 0, 0, 0,
	0, 60, 57, 0, 0, 69, 59, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 0, 0,
	0, 238, 239, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 0, 65, 0, 76, 75, 61, 41, 67,
	42, 68, 0, 0, 0, 38, 349, 47, 0, 0,
	48, 39, 40, 0, 58, 0, 49, 0, 0, 0,
	0, 0, 0, 60, 57, 0, 0, 69, 59, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 238, 239, 0, 61, 41, 67, 42,
	68, 0, 0, 64, 38, 65, 47, 76, 75, 48,
	39, 40, 0, 58, 0, 49, 0, 0, 0, 0,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 238, 239, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 340, 76, 75, 61, 41,
	67, 42, 68, 0, 0, 0, 38, 337, 47, 0,
	0, 48, 39, 40, 0, 58, 0, 49, 0, 0,
	0, 0, 0, 0, 60, 57, 0, 0, 69, 59,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 238, 239, 0, 61, 41, 67,
	42, 68, 0, 0, 64, 38, 65, 47, 76, 75,
	48, 39, 40, 0, 58, 0, 49, 0, 0, 0,
	0, 0, 0, 60, 57, 0, 0, 69, 59, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 238, 239, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 0, 65, 293, 76, 75, 61,
	41, 67, 42, 68, 0, 0, 0, 38, 288, 47,
	0, 0, 48, 39, 40, 0, 58, 0, 49, 0,
	0, 0, 0, 0, 0, 60, 57, 0, 0, 69,
	59, 0, 77, 78, 0, 0, 0, 70, 71, 72,
	73, 74, 0, 0, 0, 238, 239, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 0, 65, 0, 76,
	75, 61, 41, 67, 42, 68, 0, 0, 0, 38,
	279, 47, 0, 0, 48, 39, 40, 0, 58, 0,
	49, 0, 0, 0, 0, 0, 0, 60, 57, 0,
	0, 69, 59, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 0, 0, 0, 238, 239, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 0, 65,
	0, 76, 75, 61, 41, 67, 42, 68, 0, 0,
	0, 38, 277, 47, 0, 0, 48, 39, 40, 0,
	58, 0, 49, 0, 0, 0, 0, 0, 0, 60,
	57, 0, 0, 69, 59, 0, 77, 78, 0, 0,
	0, 70, 71, 72, 73, 74, 0, 0, 0, 238,
	239, 0, 61, 41, 67, 42, 68, 0, 0, 64,
	38, 65, 47, 76, 75, 48, 39, 40, 0, 58,
	0, 49, 0, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 238, 239,
	0, 61, 41, 67, 42, 68, 0, 0, 64, 38,
	65, 47, 76, 75, 48, 39, 40, 0, 58, 0,
	49, 0, 0, 0, 0, 0, 0, 60, 57, 0,
	0, 69, 59, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 61, 123, 67, 124, 110, 424,
	0, 121, 0, 0, 0, 0, 0, 64, 0, 65,
	0, 76, 75, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 0, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 61, 123, 67,
	124, 68, 0, 0, 0, 0, 0, 0, 0, 0,
	113, 0, 122, 0, 76, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 0, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 249, 61, 123, 67, 124, 68, 0,
	382, 0, 0, 64, 0, 65, 0, 76, 75, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 0, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 0, 0, 0,
	249, 61, 123, 67, 124, 68, 0, 368, 0, 0,
	64, 0, 65, 0, 76, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 0, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 0, 0, 0, 249, 61, 123,
	67, 124, 110, 0, 334, 121, 0, 64, 0, 65,
	0, 76, 75, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 0,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 61, 123, 67, 124, 265, 315, 0, 121, 0,
	0, 0, 0, 0, 113, 0, 122, 0, 76, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 0, 0, 77, 78, 0, 0, 264, 70,
	71, 72, 73, 74, 61, 123, 67, 124, 68, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 0, 122,
	0, 76, 75, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 0, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 0, 0, 0,
	249, 61, 123, 67, 124, 68, 0, 0, 0, 0,
	64, 0, 65, 0, 76, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 0, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 61, 123, 67, 124, 68, 0,
	0, 121, 0, 0, 152, 0, 0, 64, 0, 65,
	0, 76, 75, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 0, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 61, 254, 67,
	124, 68, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 0, 122, 0, 76, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 0, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 249, 61, 123, 67, 124, 110, 0,
	0, 121, 0, 64, 0, 65, 0, 76, 75, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 0, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 61, 123, 67,
	124, 68, 0, 0, 0, 0, 0, 0, 0, 0,
	113, 0, 122, 0, 76, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 0, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	61, 408, 67, 124, 68, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 0, 65, 0, 76, 75, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	69, 0, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 65, 0,
	76, 75,
}
var RubyPact = []int{

	-27, 1566, -1000, -1000, -1000, 5, -1000, -1000, -1000, 329,
	-1000, -1000, -1000, 67, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 205, -11,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 374, 232,
	232, 1142, 157, 1, 79, 9, 65, 2446, 2446, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 2982, 2446, 2982,
	2982, -1000, -1000, -1000, 2806, -1000, -9, 170, -1000, 14,
	2446, 2446, 2982, 2982, 2982, 21, 360, -1000, -1000, -1000,
	-1000, -1000, 2446, 2446, 2982, 352, 2982, 2982, 2982, 2982,
	2446, 2982, 2982, 2982, 2982, 2446, 2982, -1000, -1000, 2982,
	2446, 2982, 2446, 2446, 349, 263, 106, 41, -1000, -1000,
	2806, 14, 25, 2806, 2982, 2982, 346, 159, 201, -1000,
	30, -14, -14, 2939, 93, -29, -1000, -1000, 2806, 2446,
	2446, 2982, 2446, 2446, 2446, 2446, 2446, 2446, 2446, 345,
	102, 277, 2397, 86, 201, 260, 201, 86, 28, 24,
	203, -1000, 2982, 2892, 247, 2806, 773, -1000, -14, 102,
	102, 201, 201, 201, -1000, -1000, 288, -1000, -1000, 102,
	102, 201, 2716, 201, 201, 201, 201, 102, 201, 201,
	201, 201, 102, 272, 2759, 2759, 201, 102, 201, 102,
	102, -1000, -1000, 342, 149, 48, -1000, 33, 337, 308,
	-1000, 2348, 232, 2286, 298, 203, -1000, -1000, -1000, 504,
	-40, 121, -1000, -1000, 59, -1000, -1000, 2849, 2224, -1000,
	297, -1000, 2162, 236, 102, 102, 27, 102, 102, 102,
	102, 102, 102, 102, -1000, 1517, -1000, -1000, -1000, -1000,
	102, 204, 2982, -1000, 57, -1000, -1000, -1000, 201, -1000,
	136, 115, -4, 316, 2673, -1000, 214, 102, -1000, -1000,
	-1000, -1000, 25, 14, 2446, 2806, 2982, 201, 201, 1816,
	106, 48, 206, 2982, -1000, -1000, 1754, -1000, -1000, -1000,
	14, -1000, 2626, 23, -1000, -18, 201, -1000, -1000, 2113,
	53, -1000, 2051, -1000, -1000, 13, -1000, 378, 2446, -1000,
	1455, 2002, -1000, -1000, 211, 201, 1393, 252, 2982, 2579,
	-30, -1000, -34, -1000, 2446, 2982, -1000, -1000, 102, 100,
	201, -1000, 196, -1000, -1000, -1000, -1000, 201, -1000, 183,
	1940, -1000, 2532, 201, 292, 2446, 282, -1000, -1000, 264,
	-1000, 2446, -1000, 2446, -1000, 102, 2397, -1000, 323, -1000,
	2397, 200, -1000, -1000, -1000, 102, -1000, -1000, 2446, 2446,
	304, 267, -1000, -1000, 2982, 86, 203, -1000, 2982, -1000,
	2759, -1000, 68, 329, 102, 201, -1000, -1000, -1000, 1692,
	-1000, -1000, 3025, -1000, 102, -1000, -1000, 102, 102, 2397,
	2397, -1000, 2397, 257, 150, 249, 2446, 2446, 2446, 2446,
	1878, 86, 2397, 201, 198, -7, -1000, 91, 2489, 2397,
	-1000, -1000, -1000, -1000, 102, 102, 102, 102, -1000, 2397,
	-4, 2446, 2982, -1000, -1000, 2397, 1322, 1205, 570, -4,
	275, 141, -1000, 318, 2446, -1000, -1000, 134, -1000, -1000,
	-1000, -4, -1000, -1000, 2446, -1000, 102, 1630, -1000, -4,
	102, 1630, 1630, 1630,
}
var RubyPgo = []int{

	0, 0, 468, 465, 18, 10, 464, 460, 458, 1010,
	457, 15, 28, 456, 449, 447, 446, 98, 444, 807,
	604, 442, 441, 439, 13, 438, 9, 401, 437, 436,
	12, 434, 432, 430, 426, 14, 425, 424, 421, 5,
	420, 419, 415, 412, 411, 410, 409, 406, 405, 404,
	403, 653, 402, 3, 16, 19, 1, 396, 6, 390,
	4, 389, 38, 382, 8, 379, 367, 11, 7, 324,
	211, 79,
}
var RubyR1 = []int{

	0, 57, 57, 57, 57, 57, 57, 57, 57, 57,
	57, 70, 70, 71, 71, 51, 51, 51, 51, 18,
	18, 18, 18, 18, 18, 18, 18, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 24, 24, 24, 24, 24, 24, 24, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 35, 14, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 54, 54,
	54, 54, 64, 64, 62, 62, 62, 62, 62, 62,
	62, 67, 67, 67, 67, 67, 66, 66, 66, 21,
	21, 21, 21, 21, 21, 58, 58, 68, 68, 68,
	26, 26, 26, 26, 25, 25, 28, 30, 30, 69,
	69, 15, 15, 15, 15, 15, 15, 15, 15, 29,
	29, 29, 29, 29, 29, 9, 9, 27, 27, 19,
	19, 40, 40, 41, 42, 43, 44, 45, 46, 47,
	48, 49, 50, 2, 6, 7, 7, 3, 3, 59,
	59, 59, 59, 65, 65, 65, 5, 5, 5, 5,
	55, 63, 63, 63, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 56, 56, 56, 56, 52,
	52, 52, 8, 16, 11, 11, 11, 61, 61, 53,
	53, 22, 23, 23, 12, 36, 60, 60, 60, 60,
	60, 60, 37, 37, 37, 37, 37, 37, 37, 38,
	38, 38, 38, 38, 39, 39, 39, 39, 34, 33,
	10, 32, 32, 31, 31, 4,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 1, 0, 2, 0, 2, 2, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 2, 4, 5, 1, 4, 4,
	2, 3, 3, 4, 4, 5, 3, 5, 2, 3,
	3, 3, 3, 4, 4, 4, 6, 6, 3, 7,
	1, 5, 1, 3, 0, 1, 1, 2, 4, 4,
	5, 1, 1, 4, 2, 5, 1, 3, 3, 5,
	6, 7, 8, 5, 6, 1, 3, 0, 1, 3,
	1, 2, 3, 2, 4, 6, 4, 1, 3, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 2, 2, 2, 3,
	3, 3, 3, 2, 2, 2, 2, 3, 3, 3,
	3, 3, 3, 1, 1, 3, 3, 5, 5, 0,
	1, 3, 7, 3, 7, 8, 3, 4, 4, 3,
	3, 0, 1, 3, 4, 5, 3, 3, 3, 3,
	3, 5, 6, 5, 3, 4, 3, 3, 2, 0,
	2, 2, 3, 4, 2, 3, 5, 0, 2, 1,
	2, 2, 2, 1, 5, 5, 0, 2, 2, 2,
	2, 2, 0, 1, 3, 3, 1, 3, 3, 5,
	6, 5, 6, 5, 4, 3, 3, 2, 4, 4,
	2, 5, 7, 4, 5, 3,
}
var RubyChk = []int{

	-1000, -57, 51, 52, 68, -1, 51, 52, 68, -17,
	-21, -25, -28, -15, -29, -13, -16, -24, -22, -36,
	-32, -23, -35, -4, -18, -7, -3, -30, -20, -8,
	-10, -41, -42, -43, -44, -14, -34, -33, 13, 19,
	20, 6, 8, -27, -19, -9, -40, 15, 18, 24,
	-45, -46, -47, -48, -49, -50, -12, 32, 22, 36,
	31, 5, -2, -6, 61, 63, -69, 7, 9, 35,
	43, 44, 45, 46, 47, 66, 65, 38, 39, 52,
	51, 68, 15, 18, 25, 55, 40, 41, 4, 45,
	46, 47, 57, 58, 56, 18, 59, 33, 34, 48,
	18, 40, 15, 18, 55, 6, 4, -30, 8, -30,
	9, -54, -67, 61, 42, 49, 11, -66, -17, -5,
	-20, 12, 63, 6, 8, -27, -19, -9, 9, 42,
	49, 61, 42, 49, 42, 49, 42, 49, 42, 11,
	-1, -1, -51, -64, -17, -1, -17, -64, -64, -62,
	-17, -24, 58, -71, 54, 9, -52, -5, 63, -1,
	-1, -17, -17, -17, 6, 8, 66, 6, 8, -1,
	-1, -17, 6, -17, -17, -17, -17, -1, -17, -17,
	-17, -17, -1, -17, -71, -71, -17, -1, -17, -1,
	-1, 6, -58, 55, -68, 9, -26, 6, 47, 58,
	-58, -51, 40, -51, -62, -17, -5, 11, -5, -17,
	-4, -17, -35, -12, -17, -12, 6, 11, -51, -55,
	56, -55, -51, -62, -1, -1, -17, -1, -1, -1,
	-1, -1, -1, -1, 6, -51, 51, 52, 51, 52,
	-1, -61, 11, 51, -71, 62, 11, 62, -17, 51,
	-59, -65, -71, -17, 6, 8, -62, -1, 52, 10,
	6, 8, -67, -54, 42, 9, 53, -17, -17, -51,
	6, 11, -68, 42, 6, 6, -51, 14, -30, 14,
	10, 11, -71, 62, 62, -71, -17, -5, 14, -51,
	-63, 6, -51, 64, 10, 62, 14, -56, 17, 16,
	-51, -51, 14, -11, 25, -17, -60, -31, 37, -71,
	-71, 11, -71, 11, 4, 53, 10, -5, -1, -62,
	-17, 14, -53, -11, -58, -26, 10, -17, 14, -53,
	-51, -5, -71, -17, 58, 42, 58, 14, 56, 11,
	64, 42, 14, 17, 16, -1, -51, 14, -56, 14,
	-51, 8, 14, 51, 52, -1, -38, -37, 15, 18,
	27, 28, 14, 16, 37, -64, -17, -24, 58, 64,
	-71, 64, -71, -17, -1, -17, 10, 14, -11, -51,
	14, 14, 58, 6, -1, 6, 6, -1, -1, -51,
	-51, 14, -51, 4, -1, -1, 15, 18, 15, 18,
	-51, -64, -51, -17, -17, 6, 14, -53, 6, -51,
	6, 51, 51, 52, -1, -1, -1, -1, 14, -51,
	-71, 4, 53, 14, 10, -51, -60, -60, -60, -71,
	-1, -17, 14, -39, 17, 16, 14, -39, 14, -70,
	11, -71, 11, 14, 17, 16, -1, -60, 14, -71,
	-1, -60, -60, -60,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 48,
	49, 50, 51, 52, 53, 54, 55, 56, 57, 58,
	59, 60, 61, 62, 27, 28, 29, 30, 31, 32,
	33, 34, 35, 36, 37, 38, 39, 40, 0, 0,
	0, 20, -2, 22, 23, 24, 0, 0, 0, 15,
	41, 42, 43, 44, 45, 46, 47, 0, 0, 0,
	213, 19, 25, 26, 94, 13, 0, 67, 199, 0,
	0, 0, 0, 0, 0, 0, 0, 163, 164, 5,
	6, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 13, 13, 0,
	0, 0, 0, 0, 0, 117, 117, 15, -2, 15,
	-2, 70, 78, 94, 0, 0, 0, 90, 101, 102,
	31, 15, -2, 20, -2, 22, 23, 24, 94, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	15, 0, 207, 211, 92, 0, 13, 212, 0, 0,
	92, 96, 0, 13, 0, 94, 0, 240, 15, 153,
	154, 155, 156, 64, 147, 148, 0, 145, 146, 186,
	194, 63, 72, 79, 81, 82, 157, 158, 159, 160,
	161, 162, 188, 0, 0, 0, 245, 190, 80, 187,
	189, 76, 15, 0, 115, 117, 118, 120, 0, 0,
	15, 0, 0, 0, 0, 95, 71, 13, 104, 92,
	0, 131, 132, 133, 139, 140, 151, 13, 0, 15,
	181, 15, 0, 0, 134, 141, 0, 135, 142, 136,
	143, 137, 144, 138, 152, 0, 15, 15, 16, 17,
	18, 0, 0, 216, 0, 165, 13, 166, 97, 14,
	13, 13, 170, 0, 20, -2, 0, 200, 201, 202,
	149, 150, 73, 74, 0, -2, 0, 238, 239, 0,
	117, 0, 0, 0, 121, 123, 0, 124, 15, 126,
	65, 13, 0, 83, 85, 0, 107, 108, 176, 0,
	0, 182, 0, 179, 69, 84, 184, 0, 0, 15,
	0, 0, 203, 208, 15, 93, 0, 0, 0, 0,
	0, 13, 0, 13, 0, 0, 68, 75, 77, 0,
	214, 109, 0, 209, 15, 119, 116, 122, 113, 0,
	0, 66, 0, 103, 0, 0, 0, 177, 180, 0,
	178, 0, 185, 0, 15, 15, 198, 191, 0, 193,
	204, 15, 215, 217, 218, 219, 220, 221, 0, 0,
	223, 226, 241, 15, 0, 15, 98, 99, 0, 167,
	0, 168, 0, 48, 171, 173, 88, 110, 210, 0,
	114, 125, 0, 105, 86, 91, 183, 87, 15, 196,
	197, 192, 205, 0, 15, 0, 0, 0, 0, 0,
	0, 15, 13, 100, 0, 0, 111, 0, 20, 195,
	15, 216, -2, -2, 224, 225, 227, 228, 242, 13,
	243, 0, 0, 112, 89, 206, 0, 0, 0, 244,
	11, 13, 229, 0, 0, 216, 231, 0, 233, 172,
	12, 174, 13, 230, 0, 216, 216, 222, 232, 175,
	216, 222, 222, 222,
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
	62, 63, 64, 65, 66, 67, 68,
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
		//line parser.y:201
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:203
		{
		}
	case 3:
		//line parser.y:205
		{
		}
	case 4:
		//line parser.y:207
		{
		}
	case 5:
		//line parser.y:209
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:211
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:213
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:219
		{
		}
	case 11:
		//line parser.y:221
		{
		}
	case 12:
		//line parser.y:222
		{
		}
	case 13:
		//line parser.y:224
		{
		}
	case 14:
		//line parser.y:225
		{
		}
	case 15:
		//line parser.y:228
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 16:
		//line parser.y:230
		{
		}
	case 17:
		//line parser.y:232
		{
		}
	case 18:
		//line parser.y:234
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 57:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 58:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 59:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 60:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 61:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 62:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 63:
		//line parser.y:246
		{
			RubyVAL.genericValue = ast.RescueModifier{Statement: RubyS[Rubypt-2].genericValue, Rescue: RubyS[Rubypt-0].genericValue}
		}
	case 64:
		//line parser.y:249
		{
			RubyVAL.genericValue = ast.StarSplat{Value: RubyS[Rubypt-0].genericValue}
		}
	case 65:
		//line parser.y:252
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 66:
		//line parser.y:259
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-2].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 67:
		//line parser.y:266
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 68:
		//line parser.y:270
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 69:
		//line parser.y:277
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 70:
		//line parser.y:284
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 71:
		//line parser.y:291
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 72:
		//line parser.y:298
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 73:
		//line parser.y:305
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 74:
		//line parser.y:313
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 75:
		//line parser.y:321
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 76:
		//line parser.y:329
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 77:
		//line parser.y:337
		{
			methodName := RubyS[Rubypt-2].genericValue.(ast.BareReference).Name + "="
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: methodName},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 78:
		//line parser.y:348
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 79:
		//line parser.y:355
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 80:
		//line parser.y:363
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 81:
		//line parser.y:371
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 82:
		//line parser.y:379
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 83:
		//line parser.y:389
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 84:
		//line parser.y:397
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 85:
		//line parser.y:405
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 86:
		//line parser.y:416
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 87:
		//line parser.y:424
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 88:
		//line parser.y:434
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 89:
		//line parser.y:436
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 90:
		//line parser.y:438
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 91:
		//line parser.y:440
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 92:
		//line parser.y:443
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 93:
		//line parser.y:445
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 94:
		//line parser.y:447
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 95:
		//line parser.y:449
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 96:
		//line parser.y:451
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 97:
		//line parser.y:453
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{Func: ast.BareReference{Name: "to_proc"}, Target: RubyS[Rubypt-0].genericValue})
		}
	case 98:
		//line parser.y:455
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 99:
		//line parser.y:457
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 100:
		//line parser.y:459
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{
				Func:   ast.BareReference{Name: "to_proc"},
				Target: RubyS[Rubypt-0].genericValue,
			})
		}
	case 101:
		//line parser.y:468
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 102:
		//line parser.y:470
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 103:
		//line parser.y:472
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 104:
		//line parser.y:474
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 105:
		//line parser.y:476
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 106:
		//line parser.y:479
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 107:
		//line parser.y:481
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 108:
		//line parser.y:483
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 109:
		//line parser.y:487
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 110:
		//line parser.y:495
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 111:
		//line parser.y:504
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 112:
		//line parser.y:513
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 113:
		//line parser.y:523
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 114:
		//line parser.y:531
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 115:
		//line parser.y:542
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 116:
		//line parser.y:544
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 117:
		//line parser.y:546
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 118:
		//line parser.y:548
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 119:
		//line parser.y:550
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 120:
		//line parser.y:553
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 121:
		//line parser.y:555
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 122:
		//line parser.y:557
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 123:
		//line parser.y:559
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 124:
		//line parser.y:563
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 125:
		//line parser.y:570
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 126:
		//line parser.y:580
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 127:
		//line parser.y:589
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 128:
		//line parser.y:595
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 129:
		//line parser.y:603
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 130:
		//line parser.y:607
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 131:
		//line parser.y:612
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 132:
		//line parser.y:619
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 133:
		//line parser.y:623
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 134:
		//line parser.y:627
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 135:
		//line parser.y:634
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 136:
		//line parser.y:641
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 137:
		//line parser.y:648
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 138:
		//line parser.y:655
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 139:
		//line parser.y:663
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 140:
		//line parser.y:670
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 141:
		//line parser.y:674
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 142:
		//line parser.y:681
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 143:
		//line parser.y:688
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 144:
		//line parser.y:695
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 145:
		//line parser.y:703
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 146:
		//line parser.y:705
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 147:
		//line parser.y:708
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 148:
		//line parser.y:710
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 149:
		//line parser.y:713
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 150:
		//line parser.y:715
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 151:
		//line parser.y:718
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 152:
		//line parser.y:720
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 153:
		//line parser.y:722
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 154:
		//line parser.y:723
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 155:
		//line parser.y:724
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 156:
		//line parser.y:725
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 157:
		//line parser.y:728
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 158:
		//line parser.y:737
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 159:
		//line parser.y:746
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 160:
		//line parser.y:755
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 161:
		//line parser.y:764
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 162:
		//line parser.y:773
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 163:
		//line parser.y:781
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 164:
		//line parser.y:782
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 165:
		//line parser.y:784
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 166:
		//line parser.y:786
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 167:
		//line parser.y:789
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 168:
		//line parser.y:797
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 169:
		//line parser.y:805
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 170:
		//line parser.y:807
		{
		}
	case 171:
		//line parser.y:809
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 172:
		//line parser.y:816
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 173:
		//line parser.y:824
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 174:
		//line parser.y:831
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 175:
		//line parser.y:838
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 176:
		//line parser.y:846
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 177:
		//line parser.y:848
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 178:
		//line parser.y:855
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 179:
		//line parser.y:862
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 180:
		//line parser.y:865
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 181:
		//line parser.y:867
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 182:
		//line parser.y:869
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 183:
		//line parser.y:871
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 184:
		//line parser.y:874
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 185:
		//line parser.y:881
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 186:
		//line parser.y:889
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 187:
		//line parser.y:896
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 188:
		//line parser.y:903
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 189:
		//line parser.y:910
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 190:
		//line parser.y:917
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 191:
		//line parser.y:924
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 192:
		//line parser.y:931
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 193:
		//line parser.y:939
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 194:
		//line parser.y:946
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 195:
		//line parser.y:955
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 196:
		//line parser.y:962
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 197:
		//line parser.y:969
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 198:
		//line parser.y:976
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 199:
		//line parser.y:983
		{
		}
	case 200:
		//line parser.y:984
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 201:
		//line parser.y:985
		{
		}
	case 202:
		//line parser.y:988
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 203:
		//line parser.y:991
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 204:
		//line parser.y:999
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 205:
		//line parser.y:1001
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 206:
		//line parser.y:1010
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
	case 207:
		//line parser.y:1025
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 208:
		//line parser.y:1027
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 209:
		//line parser.y:1030
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 210:
		//line parser.y:1032
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 211:
		//line parser.y:1035
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 212:
		//line parser.y:1044
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 213:
		//line parser.y:1052
		{
			RubyVAL.genericValue = ast.Return{}
		}
	case 214:
		//line parser.y:1055
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 215:
		//line parser.y:1064
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 216:
		//line parser.y:1067
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 217:
		//line parser.y:1069
		{
		}
	case 218:
		//line parser.y:1071
		{
		}
	case 219:
		//line parser.y:1073
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 220:
		//line parser.y:1075
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 221:
		//line parser.y:1077
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 222:
		//line parser.y:1079
		{
		}
	case 223:
		//line parser.y:1081
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 224:
		//line parser.y:1083
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 225:
		//line parser.y:1085
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 226:
		//line parser.y:1087
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 227:
		//line parser.y:1089
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 228:
		//line parser.y:1091
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 229:
		//line parser.y:1094
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 230:
		//line parser.y:1101
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 231:
		//line parser.y:1109
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 232:
		//line parser.y:1116
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 233:
		//line parser.y:1124
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 234:
		//line parser.y:1132
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 235:
		//line parser.y:1139
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 236:
		//line parser.y:1146
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 237:
		//line parser.y:1153
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 238:
		//line parser.y:1161
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 239:
		//line parser.y:1164
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 240:
		//line parser.y:1166
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 241:
		//line parser.y:1169
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 242:
		//line parser.y:1171
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 243:
		//line parser.y:1174
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 244:
		//line parser.y:1176
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 245:
		//line parser.y:1178
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
