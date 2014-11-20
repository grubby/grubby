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

//line parser.y:1204

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 42,
	54, 132,
	-2, 21,
	-1, 109,
	54, 132,
	-2, 130,
	-1, 111,
	10, 97,
	11, 97,
	-2, 202,
	-1, 123,
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
	-1, 125,
	54, 132,
	-2, 21,
	-1, 259,
	54, 133,
	-2, 131,
	-1, 269,
	10, 97,
	11, 97,
	-2, 202,
	-1, 382,
	61, 82,
	-2, 48,
	-1, 422,
	27, 219,
	28, 219,
	-2, 15,
	-1, 423,
	27, 219,
	28, 219,
	-2, 15,
}

const RubyNprod = 249
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 3100

var RubyAct = []int{

	244, 331, 5, 443, 332, 149, 120, 305, 195, 197,
	27, 199, 150, 17, 113, 118, 291, 22, 225, 112,
	253, 82, 2, 3, 83, 253, 122, 133, 165, 56,
	166, 88, 23, 380, 134, 210, 122, 314, 378, 4,
	132, 155, 224, 274, 103, 105, 132, 104, 141, 142,
	108, 110, 274, 432, 253, 250, 88, 80, 79, 146,
	97, 98, 246, 144, 140, 316, 148, 86, 87, 101,
	205, 160, 161, 137, 81, 350, 158, 159, 152, 253,
	138, 348, 85, 170, 171, 97, 98, 159, 167, 303,
	102, 178, 86, 87, 292, 139, 183, 89, 90, 91,
	99, 188, 9, 273, 192, 193, 251, 85, 94, 92,
	93, 96, 200, 249, 290, 198, 203, 129, 190, 209,
	211, 415, 253, 344, 207, 152, 347, 101, 152, 345,
	214, 228, 229, 216, 231, 232, 233, 234, 235, 236,
	237, 223, 227, 152, 119, 217, 219, 213, 102, 88,
	130, 329, 135, 201, 280, 100, 129, 131, 261, 136,
	145, 196, 147, 145, 202, 246, 253, 151, 200, 260,
	152, 198, 433, 390, 200, 162, 163, 164, 97, 98,
	371, 321, 372, 312, 312, 86, 87, 172, 266, 174,
	175, 176, 177, 267, 179, 180, 181, 182, 82, 184,
	85, 83, 187, 373, 189, 191, 319, 275, 279, 201,
	88, 387, 310, 278, 208, 201, 285, 212, 215, 218,
	202, 221, 312, 312, 95, 156, 202, 119, 295, 450,
	360, 84, 208, 82, 421, 230, 83, 385, 288, 97,
	98, 82, 88, 109, 83, 82, 86, 87, 83, 259,
	311, 89, 90, 91, 99, 82, 252, 257, 83, 208,
	420, 85, 94, 92, 93, 96, 88, 408, 403, 326,
	409, 97, 98, 211, 325, 396, 119, 247, 86, 87,
	395, 295, 327, 152, 393, 338, 333, 84, 271, 272,
	334, 422, 423, 85, 340, 97, 98, 96, 335, 278,
	299, 82, 86, 87, 83, 88, 458, 354, 455, 454,
	453, 282, 455, 454, 449, 364, 357, 85, 324, 250,
	281, 96, 374, 383, 294, 406, 302, 250, 407, 277,
	386, 376, 287, 288, 97, 98, 388, 240, 241, 238,
	220, 86, 87, 388, 194, 394, 89, 90, 91, 313,
	264, 397, 265, 398, 173, 376, 85, 94, 92, 93,
	96, 119, 66, 401, 88, 353, 352, 88, 404, 405,
	255, 452, 208, 328, 168, 298, 169, 294, 351, 411,
	353, 352, 107, 336, 106, 245, 254, 1, 157, 55,
	54, 417, 342, 97, 98, 53, 97, 98, 431, 52,
	86, 87, 51, 86, 87, 50, 34, 424, 425, 426,
	427, 33, 43, 32, 31, 85, 270, 46, 85, 145,
	375, 365, 388, 366, 19, 382, 384, 97, 98, 36,
	37, 20, 440, 315, 86, 87, 14, 364, 364, 364,
	12, 447, 11, 21, 375, 456, 18, 10, 24, 85,
	16, 13, 35, 15, 126, 460, 30, 29, 364, 436,
	437, 438, 364, 364, 364, 25, 63, 26, 62, 0,
	126, 0, 126, 126, 0, 0, 145, 126, 0, 0,
	413, 0, 414, 457, 0, 126, 126, 126, 0, 88,
	28, 0, 0, 461, 462, 413, 0, 126, 463, 126,
	126, 126, 126, 0, 126, 126, 126, 126, 0, 126,
	0, 0, 126, 0, 126, 126, 0, 322, 97, 98,
	0, 119, 0, 0, 126, 86, 87, 126, 126, 126,
	0, 0, 121, 0, 382, 441, 0, 126, 0, 0,
	85, 0, 126, 0, 0, 126, 97, 98, 121, 0,
	121, 121, 0, 86, 87, 121, 0, 0, 0, 0,
	0, 0, 0, 121, 121, 121, 126, 126, 85, 126,
	61, 124, 67, 125, 68, 121, 0, 121, 121, 121,
	121, 0, 121, 121, 121, 121, 126, 121, 0, 0,
	121, 0, 121, 121, 0, 0, 0, 0, 126, 126,
	69, 0, 121, 77, 78, 121, 121, 121, 70, 71,
	72, 73, 74, 0, 0, 121, 253, 0, 0, 0,
	121, 0, 0, 121, 0, 0, 64, 0, 65, 0,
	76, 75, 0, 0, 126, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 121, 121, 0, 121, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 126,
	0, 0, 0, 0, 121, 0, 61, 124, 67, 125,
	111, 126, 117, 122, 0, 0, 121, 121, 0, 0,
	0, 0, 126, 126, 0, 0, 0, 126, 0, 0,
	0, 0, 0, 126, 0, 0, 69, 0, 0, 77,
	78, 0, 126, 115, 70, 71, 72, 73, 74, 0,
	116, 0, 121, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 114, 0, 123, 0, 76, 75, 0, 126,
	126, 0, 0, 0, 0, 0, 126, 121, 0, 0,
	0, 0, 0, 61, 124, 67, 125, 111, 434, 121,
	122, 0, 0, 0, 126, 0, 0, 0, 0, 0,
	121, 121, 0, 0, 0, 121, 0, 0, 0, 0,
	0, 121, 0, 69, 0, 0, 77, 78, 0, 0,
	121, 70, 71, 72, 73, 74, 126, 0, 0, 0,
	126, 0, 126, 0, 0, 0, 0, 0, 0, 114,
	44, 123, 0, 76, 75, 126, 0, 121, 121, 0,
	0, 0, 0, 0, 121, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 126, 121, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 127, 0, 0, 126, 0, 0, 0, 0,
	154, 0, 0, 0, 0, 0, 0, 0, 127, 0,
	127, 127, 0, 0, 121, 127, 0, 0, 121, 0,
	121, 0, 0, 127, 127, 127, 0, 0, 45, 0,
	0, 0, 0, 121, 0, 127, 0, 127, 127, 127,
	127, 0, 127, 127, 127, 127, 0, 127, 0, 0,
	127, 0, 127, 127, 0, 0, 0, 0, 0, 121,
	0, 0, 127, 0, 0, 127, 127, 127, 0, 0,
	128, 0, 0, 121, 0, 127, 0, 0, 0, 0,
	127, 0, 0, 127, 0, 0, 128, 0, 128, 128,
	0, 0, 0, 128, 0, 0, 0, 0, 185, 186,
	0, 128, 128, 128, 127, 127, 0, 127, 0, 0,
	0, 0, 0, 128, 0, 128, 128, 128, 128, 0,
	128, 128, 128, 128, 127, 128, 0, 0, 128, 0,
	128, 128, 0, 0, 0, 0, 127, 127, 226, 0,
	128, 0, 0, 128, 128, 128, 0, 0, 248, 0,
	0, 0, 0, 128, 0, 256, 0, 0, 128, 0,
	0, 128, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 127, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 128, 128, 0, 128, 0, 0, 143, 0,
	0, 0, 0, 0, 0, 0, 0, 127, 0, 0,
	0, 0, 128, 0, 61, 124, 67, 125, 68, 127,
	0, 289, 0, 0, 128, 128, 0, 0, 0, 0,
	127, 127, 293, 0, 0, 127, 0, 0, 0, 0,
	0, 127, 0, 0, 69, 0, 0, 77, 78, 0,
	127, 0, 70, 71, 72, 73, 74, 204, 0, 206,
	128, 317, 0, 0, 0, 318, 320, 153, 0, 0,
	64, 222, 65, 0, 76, 75, 0, 127, 127, 0,
	0, 0, 0, 0, 127, 128, 0, 0, 0, 0,
	239, 0, 0, 0, 0, 0, 0, 128, 0, 341,
	0, 0, 127, 0, 0, 0, 0, 0, 128, 128,
	0, 0, 0, 128, 0, 0, 0, 0, 0, 128,
	0, 0, 0, 0, 0, 0, 0, 0, 128, 0,
	379, 0, 381, 0, 127, 0, 0, 0, 127, 0,
	127, 0, 0, 0, 276, 0, 0, 0, 0, 0,
	0, 0, 283, 127, 0, 128, 128, 0, 0, 0,
	0, 0, 128, 0, 61, 124, 67, 125, 68, 0,
	0, 0, 297, 0, 300, 0, 0, 0, 0, 127,
	128, 0, 0, 0, 0, 0, 0, 0, 0, 308,
	309, 0, 0, 127, 69, 0, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 0, 0, 0,
	253, 0, 128, 0, 0, 0, 128, 392, 128, 0,
	64, 0, 65, 430, 76, 75, 0, 0, 0, 0,
	0, 128, 0, 0, 339, 0, 0, 0, 0, 0,
	439, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 451, 0, 0, 0, 355, 128, 0, 0,
	0, 359, 0, 459, 0, 0, 61, 41, 67, 42,
	68, 128, 0, 0, 38, 446, 367, 445, 444, 368,
	39, 40, 389, 58, 0, 49, 0, 0, 369, 370,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 399, 400, 70, 71, 72, 73, 74, 402,
	0, 0, 362, 363, 0, 0, 0, 0, 0, 0,
	0, 410, 64, 412, 65, 0, 76, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 419, 0, 0,
	0, 0, 0, 239, 0, 0, 0, 0, 0, 0,
	429, 0, 0, 0, 0, 0, 0, 0, 0, 435,
	0, 308, 309, 61, 41, 67, 42, 68, 0, 0,
	0, 38, 442, 367, 445, 444, 368, 39, 40, 0,
	58, 0, 49, 0, 0, 369, 370, 0, 0, 60,
	57, 0, 0, 69, 59, 0, 77, 78, 0, 0,
	0, 70, 71, 72, 73, 74, 0, 0, 0, 362,
	363, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	0, 65, 0, 76, 75, 61, 41, 67, 42, 68,
	0, 0, 0, 38, 448, 367, 0, 0, 368, 39,
	40, 0, 58, 0, 49, 0, 0, 369, 370, 0,
	0, 60, 57, 0, 0, 69, 59, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 0, 0,
	0, 362, 363, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 0, 65, 0, 76, 75, 61, 41, 67,
	42, 68, 0, 0, 0, 38, 361, 367, 0, 0,
	368, 39, 40, 0, 58, 0, 49, 0, 0, 369,
	370, 0, 0, 60, 57, 0, 0, 69, 59, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 362, 363, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 0, 65, 0, 76, 75, 61,
	41, 67, 42, 68, 0, 0, 0, 38, 356, 47,
	307, 306, 48, 39, 40, 0, 58, 0, 49, 0,
	0, 0, 0, 0, 0, 60, 57, 0, 0, 69,
	59, 0, 77, 78, 0, 0, 0, 70, 71, 72,
	73, 74, 0, 0, 0, 242, 243, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 0, 65, 0, 76,
	75, 61, 41, 67, 42, 68, 0, 0, 0, 38,
	304, 47, 307, 306, 48, 39, 40, 0, 58, 0,
	49, 0, 0, 0, 0, 0, 0, 60, 57, 0,
	0, 69, 59, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 0, 0, 0, 242, 243, 0,
	61, 41, 67, 42, 68, 0, 0, 64, 38, 65,
	47, 76, 75, 48, 39, 40, 0, 58, 0, 49,
	0, 0, 0, 0, 0, 0, 60, 57, 0, 0,
	69, 59, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 0, 0, 0, 6, 7, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 65, 0,
	76, 75, 0, 8, 61, 41, 67, 42, 68, 0,
	0, 0, 38, 0, 367, 0, 0, 368, 39, 40,
	0, 58, 0, 49, 0, 0, 369, 370, 0, 0,
	60, 57, 0, 0, 69, 59, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 0, 0, 0,
	362, 363, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 0, 65, 0, 76, 75, 61, 41, 67, 42,
	68, 0, 0, 0, 38, 416, 47, 0, 0, 48,
	39, 40, 0, 58, 0, 49, 312, 0, 0, 0,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 242, 243, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 0, 76, 75, 61, 41,
	67, 42, 68, 0, 0, 0, 38, 337, 47, 0,
	0, 48, 39, 40, 0, 58, 0, 49, 312, 0,
	0, 0, 0, 0, 60, 57, 0, 0, 69, 59,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 242, 243, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 0, 65, 0, 76, 75,
	61, 41, 67, 42, 68, 0, 0, 0, 38, 330,
	47, 0, 0, 48, 39, 40, 0, 58, 0, 49,
	312, 0, 0, 0, 0, 0, 60, 57, 0, 0,
	69, 59, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 0, 0, 0, 242, 243, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 65, 0,
	76, 75, 61, 41, 67, 42, 68, 0, 0, 0,
	38, 428, 47, 0, 0, 48, 39, 40, 0, 58,
	0, 49, 0, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 242, 243,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 0,
	65, 0, 76, 75, 61, 41, 67, 42, 68, 0,
	0, 0, 38, 391, 47, 0, 0, 48, 39, 40,
	0, 58, 0, 49, 0, 0, 0, 0, 0, 0,
	60, 57, 0, 0, 69, 59, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 0, 0, 0,
	242, 243, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 0, 65, 0, 76, 75, 61, 41, 67, 42,
	68, 0, 0, 0, 38, 358, 47, 0, 0, 48,
	39, 40, 0, 58, 0, 49, 0, 0, 0, 0,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 242, 243, 0, 61, 41, 67, 42, 68,
	0, 0, 64, 38, 65, 47, 76, 75, 48, 39,
	40, 0, 58, 0, 49, 0, 0, 0, 0, 0,
	0, 60, 57, 0, 0, 69, 59, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 0, 0,
	0, 242, 243, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 0, 65, 349, 76, 75, 61, 41, 67,
	42, 68, 0, 0, 0, 38, 346, 47, 0, 0,
	48, 39, 40, 0, 58, 0, 49, 0, 0, 0,
	0, 0, 0, 60, 57, 0, 0, 69, 59, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 242, 243, 0, 61, 41, 67, 42,
	68, 0, 0, 64, 38, 65, 47, 76, 75, 48,
	39, 40, 0, 58, 0, 49, 0, 0, 0, 0,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 242, 243, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 301, 76, 75, 61, 41,
	67, 42, 68, 0, 0, 0, 38, 296, 47, 0,
	0, 48, 39, 40, 0, 58, 0, 49, 0, 0,
	0, 0, 0, 0, 60, 57, 0, 0, 69, 59,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 242, 243, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 0, 65, 0, 76, 75,
	61, 41, 67, 42, 68, 0, 0, 0, 38, 286,
	47, 0, 0, 48, 39, 40, 0, 58, 0, 49,
	0, 0, 0, 0, 0, 0, 60, 57, 0, 0,
	69, 59, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 0, 0, 0, 242, 243, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 65, 0,
	76, 75, 61, 41, 67, 42, 68, 0, 0, 0,
	38, 284, 47, 0, 0, 48, 39, 40, 0, 58,
	0, 49, 0, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 242, 243,
	0, 61, 41, 67, 42, 68, 0, 0, 64, 38,
	65, 47, 76, 75, 48, 39, 40, 0, 58, 0,
	49, 0, 0, 0, 0, 0, 0, 60, 57, 0,
	0, 69, 59, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 0, 0, 0, 242, 243, 0,
	61, 41, 67, 42, 68, 263, 0, 64, 38, 65,
	47, 76, 75, 48, 39, 40, 0, 58, 0, 49,
	0, 0, 0, 0, 0, 0, 60, 57, 0, 0,
	69, 59, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 0, 0, 0, 0, 262, 0, 61,
	41, 67, 42, 68, 0, 0, 64, 38, 65, 47,
	76, 75, 48, 39, 40, 0, 58, 0, 49, 0,
	0, 0, 0, 0, 0, 60, 57, 0, 0, 69,
	59, 0, 77, 78, 0, 0, 0, 70, 71, 72,
	73, 74, 61, 124, 67, 125, 68, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 0, 65, 0, 76,
	75, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 69, 0, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 253, 61,
	124, 67, 125, 68, 0, 377, 0, 0, 64, 0,
	65, 0, 76, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	0, 0, 77, 78, 0, 0, 0, 70, 71, 72,
	73, 74, 0, 0, 0, 253, 61, 124, 67, 125,
	111, 0, 343, 122, 0, 64, 0, 65, 0, 76,
	75, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 69, 0, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 61,
	124, 67, 125, 269, 323, 0, 122, 0, 0, 0,
	0, 0, 114, 0, 123, 0, 76, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	0, 0, 77, 78, 0, 0, 268, 70, 71, 72,
	73, 74, 61, 124, 67, 125, 68, 0, 0, 122,
	0, 0, 0, 0, 0, 64, 0, 123, 0, 76,
	75, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 69, 0, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 61, 258, 67, 125, 68,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 0,
	123, 0, 76, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 0, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 0, 0,
	0, 253, 61, 124, 67, 125, 111, 0, 0, 122,
	0, 64, 0, 65, 0, 76, 75, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 69, 0, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 61, 124, 67, 125, 68,
	0, 0, 0, 0, 0, 0, 0, 0, 114, 0,
	123, 0, 76, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 0, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 61, 418,
	67, 125, 68, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 0, 65, 0, 76, 75, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 0,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 0, 65, 0, 76, 75,
}
var RubyPact = []int{

	-29, 1705, -1000, -1000, -1000, 6, -1000, -1000, -1000, 206,
	-1000, -1000, -1000, 137, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 29, -10,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 378, 235,
	235, 661, 108, -15, 110, 31, 53, 2634, 2634, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 2990, 2634, 2990,
	2990, -1000, -1000, -1000, 1049, -1000, -13, 216, -1000, 14,
	2634, 2634, 2990, 2990, 2990, 22, 368, -1000, -1000, -1000,
	-1000, -1000, 2634, 2634, 2990, 348, 2990, 2990, 2990, 2990,
	2634, 2990, 2990, 2990, 2990, 2634, 2990, -1000, -1000, 2990,
	2634, 2990, 2990, 2634, 2634, 338, 106, 162, 30, -1000,
	-1000, 1049, 14, 24, 1049, 2990, 2990, 334, 210, 485,
	-1000, 87, -14, -14, 2947, 147, -21, -1000, -1000, 1049,
	2634, 2634, 2990, 2634, 2634, 2634, 2634, 2634, 2634, 2634,
	333, 230, 286, 2536, 154, 485, 226, 485, 154, 51,
	44, 301, -1000, 2990, 2900, 241, 1049, 2585, -1000, -14,
	230, 230, 485, 485, 485, -1000, -1000, 344, -1000, -1000,
	230, 230, 485, 2814, 485, 485, 485, 485, 230, 485,
	485, 485, 485, 230, 363, 565, 565, 485, 230, 485,
	41, 145, 230, 230, -1000, -1000, 323, 202, 168, -1000,
	112, 314, 305, -1000, 2487, 235, 2425, 322, 301, -1000,
	-1000, -1000, 52, -46, 32, 262, -1000, -1000, 238, -1000,
	-1000, 2857, 2363, -1000, 294, -1000, 2301, 316, 230, 230,
	27, 230, 230, 230, 230, 230, 230, 230, -1000, 1656,
	-1000, -1000, -1000, -1000, 230, 198, 2990, -1000, 28, -1000,
	-1000, -1000, 485, -1000, 195, 170, 3, 513, 2771, -1000,
	308, 230, -1000, -1000, -1000, -1000, 24, 14, 2634, 1049,
	2990, 485, 485, -1000, 2857, 109, 1955, 162, 168, 288,
	2990, -1000, -1000, 1893, -1000, -1000, -1000, 14, -1000, 2724,
	81, -1000, -1000, 71, 485, -1000, -1000, 2252, 70, -1000,
	2190, -1000, -1000, 33, -1000, 364, 2634, -1000, 1594, 2141,
	-1000, -1000, 222, 485, 1532, 166, 2990, 2677, -26, -1000,
	-31, -1000, 2634, 2990, -1000, -1000, 230, 227, 485, 2634,
	-1000, 197, -1000, -1000, -1000, -1000, 485, -1000, 159, 2079,
	-1000, 1199, 485, 278, 2634, 274, -1000, -1000, 269, -1000,
	2634, -1000, 2634, -1000, 230, 2536, -1000, 349, -1000, 2536,
	264, -1000, -1000, -1000, 230, -1000, -1000, 2634, 2634, 310,
	252, -1000, -1000, 2990, 154, 301, -1000, 2990, -1000, 565,
	-1000, 115, 206, 230, 485, -1000, 230, -1000, -1000, 1831,
	-1000, -1000, 3033, -1000, 230, -1000, -1000, 230, 230, 2536,
	2536, -1000, 2536, 254, 183, 240, 2634, 2634, 2634, 2634,
	2017, 154, 2536, 485, 394, 0, -1000, 158, 738, 2536,
	-1000, -1000, -1000, -1000, 230, 230, 230, 230, -1000, 2536,
	3, 2634, 2990, -1000, -1000, 2536, 1408, 1301, 1470, 3,
	218, 360, -1000, 296, 2634, -1000, -1000, 292, -1000, -1000,
	-1000, 3, -1000, -1000, 2634, -1000, 230, 1769, -1000, 3,
	230, 1769, 1769, 1769,
}
var RubyPgo = []int{

	0, 0, 468, 467, 32, 6, 466, 465, 457, 878,
	456, 4, 29, 453, 452, 451, 450, 102, 448, 800,
	490, 447, 446, 443, 13, 442, 11, 412, 440, 436,
	10, 433, 431, 430, 429, 17, 424, 423, 421, 3,
	417, 414, 413, 411, 406, 405, 402, 399, 395, 390,
	389, 988, 388, 1, 19, 18, 7, 387, 8, 386,
	37, 385, 12, 375, 5, 370, 15, 14, 9, 362,
	314, 850,
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
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 54, 54, 54, 54, 64, 64, 62, 62, 62,
	62, 62, 62, 62, 67, 67, 67, 67, 67, 66,
	66, 66, 21, 21, 21, 21, 21, 21, 58, 58,
	68, 68, 68, 26, 26, 26, 26, 25, 25, 28,
	30, 30, 69, 69, 15, 15, 15, 15, 15, 15,
	15, 15, 29, 29, 29, 29, 29, 29, 9, 9,
	27, 27, 19, 19, 40, 40, 41, 42, 43, 44,
	45, 46, 47, 48, 49, 50, 2, 6, 7, 7,
	3, 3, 59, 59, 59, 59, 65, 65, 65, 5,
	5, 5, 5, 55, 63, 63, 63, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 56, 56,
	56, 56, 52, 52, 52, 8, 16, 11, 11, 11,
	61, 61, 53, 53, 22, 23, 23, 12, 36, 60,
	60, 60, 60, 60, 60, 37, 37, 37, 37, 37,
	37, 37, 38, 38, 38, 38, 38, 39, 39, 39,
	39, 34, 33, 10, 32, 32, 31, 31, 4,
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
	3, 3, 3, 4, 4, 4, 4, 4, 6, 6,
	6, 3, 7, 1, 5, 1, 3, 0, 1, 1,
	2, 4, 4, 5, 1, 1, 4, 2, 5, 1,
	3, 3, 5, 6, 7, 8, 5, 6, 1, 3,
	0, 1, 3, 1, 2, 3, 2, 4, 6, 4,
	1, 3, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	2, 2, 3, 3, 3, 3, 2, 2, 2, 2,
	3, 3, 3, 3, 3, 3, 1, 1, 3, 3,
	5, 5, 0, 1, 3, 7, 3, 7, 8, 3,
	4, 4, 3, 3, 0, 1, 3, 4, 5, 3,
	3, 3, 3, 3, 5, 6, 5, 3, 4, 3,
	3, 2, 0, 2, 2, 3, 4, 2, 3, 5,
	0, 2, 1, 2, 2, 2, 1, 5, 5, 0,
	2, 2, 2, 2, 2, 0, 1, 3, 3, 1,
	3, 3, 5, 6, 5, 6, 5, 4, 3, 3,
	2, 4, 4, 2, 5, 7, 4, 5, 3,
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
	18, 40, 61, 15, 18, 55, 6, 4, -30, 8,
	-30, 9, -54, -67, 61, 42, 49, 11, -66, -17,
	-5, -20, 12, 63, 6, 8, -27, -19, -9, 9,
	42, 49, 61, 42, 49, 42, 49, 42, 49, 42,
	11, -1, -1, -51, -64, -17, -1, -17, -64, -64,
	-62, -17, -24, 58, -71, 54, 9, -52, -5, 63,
	-1, -1, -17, -17, -17, 6, 8, 66, 6, 8,
	-1, -1, -17, 6, -17, -17, -17, -17, -1, -17,
	-17, -17, -17, -1, -17, -71, -71, -17, -1, -17,
	-66, -17, -1, -1, 6, -58, 55, -68, 9, -26,
	6, 47, 58, -58, -51, 40, -51, -62, -17, -5,
	11, -5, -17, -4, -66, -17, -35, -12, -17, -12,
	6, 11, -51, -55, 56, -55, -51, -62, -1, -1,
	-17, -1, -1, -1, -1, -1, -1, -1, 6, -51,
	51, 52, 51, 52, -1, -61, 11, 51, -71, 62,
	11, 62, -17, 51, -59, -65, -71, -17, 6, 8,
	-62, -1, 52, 10, 6, 8, -67, -54, 42, 9,
	53, -17, -17, 62, 11, 62, -51, 6, 11, -68,
	42, 6, 6, -51, 14, -30, 14, 10, 11, -71,
	62, 62, 62, -71, -17, -5, 14, -51, -63, 6,
	-51, 64, 10, 62, 14, -56, 17, 16, -51, -51,
	14, -11, 25, -17, -60, -31, 37, -71, -71, 11,
	-71, 11, 4, 53, 10, -5, -1, -62, -17, 42,
	14, -53, -11, -58, -26, 10, -17, 14, -53, -51,
	-5, -71, -17, 58, 42, 58, 14, 56, 11, 64,
	42, 14, 17, 16, -1, -51, 14, -56, 14, -51,
	8, 14, 51, 52, -1, -38, -37, 15, 18, 27,
	28, 14, 16, 37, -64, -17, -24, 58, 64, -71,
	64, -71, -17, -1, -17, 10, -1, 14, -11, -51,
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
	216, 19, 25, 26, 97, 13, 0, 67, 202, 0,
	0, 0, 0, 0, 0, 0, 0, 166, 167, 5,
	6, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 13, 13, 0,
	0, 0, 0, 0, 0, 0, 120, 120, 15, -2,
	15, -2, 70, 78, 97, 0, 0, 0, 93, 104,
	105, 31, 15, -2, 20, -2, 22, 23, 24, 97,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 15, 0, 210, 214, 95, 0, 13, 215, 0,
	0, 95, 99, 0, 13, 0, 97, 0, 243, 15,
	156, 157, 158, 159, 64, 150, 151, 0, 148, 149,
	189, 197, 63, 72, 79, 81, 82, 160, 161, 162,
	163, 164, 165, 191, 0, 0, 0, 248, 193, 80,
	0, 109, 190, 192, 76, 15, 0, 118, 120, 121,
	123, 0, 0, 15, 0, 0, 0, 0, 98, 71,
	13, 107, 95, 0, 0, 134, 135, 136, 142, 143,
	154, 13, 0, 15, 184, 15, 0, 0, 137, 144,
	0, 138, 145, 139, 146, 140, 147, 141, 155, 0,
	15, 15, 16, 17, 18, 0, 0, 219, 0, 168,
	13, 169, 100, 14, 13, 13, 173, 0, 20, -2,
	0, 203, 204, 205, 152, 153, 73, 74, 0, -2,
	0, 241, 242, 87, 0, 0, 0, 120, 0, 0,
	0, 124, 126, 0, 127, 15, 129, 65, 13, 0,
	83, 85, 86, 0, 110, 111, 179, 0, 0, 185,
	0, 182, 69, 84, 187, 0, 0, 15, 0, 0,
	206, 211, 15, 96, 0, 0, 0, 0, 0, 13,
	0, 13, 0, 0, 68, 75, 77, 0, 217, 0,
	112, 0, 212, 15, 122, 119, 125, 116, 0, 0,
	66, 0, 106, 0, 0, 0, 180, 183, 0, 181,
	0, 188, 0, 15, 15, 201, 194, 0, 196, 207,
	15, 218, 220, 221, 222, 223, 224, 0, 0, 226,
	229, 244, 15, 0, 15, 101, 102, 0, 170, 0,
	171, 0, -2, 174, 176, 91, 90, 113, 213, 0,
	117, 128, 0, 108, 88, 94, 186, 89, 15, 199,
	200, 195, 208, 0, 15, 0, 0, 0, 0, 0,
	0, 15, 13, 103, 0, 0, 114, 0, 20, 198,
	15, 219, -2, -2, 227, 228, 230, 231, 245, 13,
	246, 0, 0, 115, 92, 209, 0, 0, 0, 247,
	11, 13, 232, 0, 0, 219, 234, 0, 236, 175,
	12, 177, 13, 233, 0, 219, 219, 225, 235, 178,
	219, 225, 225, 225,
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
		//line parser.y:413
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 87:
		//line parser.y:421
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 88:
		//line parser.y:432
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 89:
		//line parser.y:440
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 90:
		//line parser.y:448
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 91:
		//line parser.y:458
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 92:
		//line parser.y:460
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 93:
		//line parser.y:462
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 94:
		//line parser.y:464
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 95:
		//line parser.y:467
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 96:
		//line parser.y:469
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 97:
		//line parser.y:471
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 98:
		//line parser.y:473
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 99:
		//line parser.y:475
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 100:
		//line parser.y:477
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{Func: ast.BareReference{Name: "to_proc"}, Target: RubyS[Rubypt-0].genericValue})
		}
	case 101:
		//line parser.y:479
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 102:
		//line parser.y:481
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 103:
		//line parser.y:483
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{
				Func:   ast.BareReference{Name: "to_proc"},
				Target: RubyS[Rubypt-0].genericValue,
			})
		}
	case 104:
		//line parser.y:492
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 105:
		//line parser.y:494
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 106:
		//line parser.y:496
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 107:
		//line parser.y:498
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 108:
		//line parser.y:500
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 109:
		//line parser.y:503
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 110:
		//line parser.y:505
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 111:
		//line parser.y:507
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 112:
		//line parser.y:511
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 113:
		//line parser.y:519
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 114:
		//line parser.y:528
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 115:
		//line parser.y:537
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 116:
		//line parser.y:547
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 117:
		//line parser.y:555
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 118:
		//line parser.y:566
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 119:
		//line parser.y:568
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 120:
		//line parser.y:570
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 121:
		//line parser.y:572
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 122:
		//line parser.y:574
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 123:
		//line parser.y:577
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 124:
		//line parser.y:579
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 125:
		//line parser.y:581
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 126:
		//line parser.y:583
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 127:
		//line parser.y:587
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 128:
		//line parser.y:594
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 129:
		//line parser.y:604
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 130:
		//line parser.y:613
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 131:
		//line parser.y:619
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 132:
		//line parser.y:627
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 133:
		//line parser.y:631
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 134:
		//line parser.y:636
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 135:
		//line parser.y:643
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 136:
		//line parser.y:647
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 137:
		//line parser.y:651
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 138:
		//line parser.y:658
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 139:
		//line parser.y:665
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 140:
		//line parser.y:672
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 141:
		//line parser.y:679
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 142:
		//line parser.y:687
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 143:
		//line parser.y:694
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 144:
		//line parser.y:698
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 145:
		//line parser.y:705
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 146:
		//line parser.y:712
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 147:
		//line parser.y:719
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 148:
		//line parser.y:727
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 149:
		//line parser.y:729
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 150:
		//line parser.y:732
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 151:
		//line parser.y:734
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 152:
		//line parser.y:737
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 153:
		//line parser.y:739
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 154:
		//line parser.y:742
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 155:
		//line parser.y:744
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 156:
		//line parser.y:746
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 157:
		//line parser.y:747
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 158:
		//line parser.y:748
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 159:
		//line parser.y:749
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 160:
		//line parser.y:752
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 161:
		//line parser.y:761
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 162:
		//line parser.y:770
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 163:
		//line parser.y:779
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 164:
		//line parser.y:788
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 165:
		//line parser.y:797
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 166:
		//line parser.y:805
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 167:
		//line parser.y:806
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 168:
		//line parser.y:808
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 169:
		//line parser.y:810
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 170:
		//line parser.y:813
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 171:
		//line parser.y:821
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 172:
		//line parser.y:829
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 173:
		//line parser.y:831
		{
		}
	case 174:
		//line parser.y:833
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 175:
		//line parser.y:840
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 176:
		//line parser.y:848
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 177:
		//line parser.y:855
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 178:
		//line parser.y:862
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 179:
		//line parser.y:870
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 180:
		//line parser.y:872
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 181:
		//line parser.y:879
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 182:
		//line parser.y:886
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 183:
		//line parser.y:889
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 184:
		//line parser.y:891
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 185:
		//line parser.y:893
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 186:
		//line parser.y:895
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 187:
		//line parser.y:898
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 188:
		//line parser.y:905
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 189:
		//line parser.y:913
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 190:
		//line parser.y:920
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 191:
		//line parser.y:927
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 192:
		//line parser.y:934
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 193:
		//line parser.y:941
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 194:
		//line parser.y:948
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 195:
		//line parser.y:955
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 196:
		//line parser.y:963
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 197:
		//line parser.y:970
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 198:
		//line parser.y:979
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 199:
		//line parser.y:986
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 200:
		//line parser.y:993
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 201:
		//line parser.y:1000
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 202:
		//line parser.y:1007
		{
		}
	case 203:
		//line parser.y:1008
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 204:
		//line parser.y:1009
		{
		}
	case 205:
		//line parser.y:1012
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 206:
		//line parser.y:1015
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 207:
		//line parser.y:1023
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 208:
		//line parser.y:1025
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 209:
		//line parser.y:1034
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
	case 210:
		//line parser.y:1049
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 211:
		//line parser.y:1051
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 212:
		//line parser.y:1054
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 213:
		//line parser.y:1056
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 214:
		//line parser.y:1059
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 215:
		//line parser.y:1068
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 216:
		//line parser.y:1076
		{
			RubyVAL.genericValue = ast.Return{}
		}
	case 217:
		//line parser.y:1079
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 218:
		//line parser.y:1088
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 219:
		//line parser.y:1091
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 220:
		//line parser.y:1093
		{
		}
	case 221:
		//line parser.y:1095
		{
		}
	case 222:
		//line parser.y:1097
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 223:
		//line parser.y:1099
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 224:
		//line parser.y:1101
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 225:
		//line parser.y:1103
		{
		}
	case 226:
		//line parser.y:1105
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 227:
		//line parser.y:1107
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 228:
		//line parser.y:1109
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 229:
		//line parser.y:1111
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 230:
		//line parser.y:1113
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 231:
		//line parser.y:1115
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 232:
		//line parser.y:1118
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 233:
		//line parser.y:1125
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 234:
		//line parser.y:1133
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 235:
		//line parser.y:1140
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 236:
		//line parser.y:1148
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 237:
		//line parser.y:1156
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 238:
		//line parser.y:1163
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 239:
		//line parser.y:1170
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 240:
		//line parser.y:1177
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 241:
		//line parser.y:1185
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 242:
		//line parser.y:1188
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 243:
		//line parser.y:1190
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 244:
		//line parser.y:1193
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 245:
		//line parser.y:1195
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 246:
		//line parser.y:1198
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 247:
		//line parser.y:1200
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 248:
		//line parser.y:1202
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
