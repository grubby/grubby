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

//line parser.y:1160

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 30,
	54, 129,
	-2, 22,
	-1, 93,
	10, 86,
	11, 86,
	-2, 197,
	-1, 104,
	54, 129,
	-2, 22,
	-1, 110,
	13, 15,
	15, 15,
	18, 15,
	19, 15,
	20, 15,
	22, 15,
	24, 15,
	32, 15,
	36, 15,
	43, 15,
	44, 15,
	45, 15,
	46, 15,
	52, 15,
	-2, 13,
	-1, 128,
	54, 129,
	-2, 127,
	-1, 223,
	48, 41,
	-2, 18,
	-1, 240,
	54, 130,
	-2, 128,
	-1, 248,
	10, 86,
	11, 86,
	-2, 197,
	-1, 401,
	27, 213,
	28, 213,
	-2, 15,
	-1, 402,
	27, 213,
	28, 213,
	-2, 15,
}

const RubyNprod = 244
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 2586

var RubyAct = []int{

	9, 204, 422, 312, 285, 318, 208, 276, 21, 38,
	148, 95, 264, 206, 149, 94, 59, 188, 234, 76,
	2, 3, 234, 102, 143, 209, 144, 109, 232, 234,
	101, 357, 187, 411, 124, 355, 306, 4, 81, 153,
	234, 118, 116, 287, 226, 305, 92, 268, 119, 117,
	135, 214, 77, 127, 129, 74, 73, 234, 309, 209,
	226, 140, 207, 142, 120, 121, 210, 386, 150, 412,
	387, 139, 75, 76, 79, 80, 151, 211, 157, 233,
	161, 162, 163, 164, 145, 166, 167, 168, 169, 78,
	171, 172, 173, 89, 176, 231, 156, 178, 179, 181,
	210, 81, 151, 308, 101, 151, 77, 209, 175, 400,
	207, 211, 176, 392, 180, 182, 114, 177, 337, 413,
	151, 199, 200, 115, 76, 292, 190, 186, 212, 81,
	283, 410, 437, 131, 434, 433, 431, 79, 80, 223,
	28, 103, 70, 104, 93, 76, 99, 109, 210, 290,
	249, 81, 78, 238, 266, 176, 205, 77, 234, 211,
	227, 101, 184, 151, 130, 79, 80, 79, 80, 241,
	72, 111, 245, 65, 66, 81, 246, 97, 77, 154,
	78, 293, 78, 128, 98, 255, 223, 79, 80, 111,
	223, 240, 82, 83, 84, 152, 96, 76, 110, 399,
	64, 63, 78, 87, 85, 86, 89, 393, 256, 253,
	76, 79, 80, 349, 223, 350, 223, 79, 80, 223,
	384, 267, 112, 385, 273, 271, 78, 284, 374, 113,
	77, 282, 78, 401, 402, 122, 351, 368, 123, 283,
	101, 429, 229, 77, 230, 76, 219, 220, 283, 176,
	299, 367, 224, 302, 5, 281, 146, 151, 147, 223,
	120, 121, 223, 298, 366, 315, 283, 321, 370, 364,
	296, 371, 315, 322, 300, 326, 303, 126, 77, 125,
	223, 223, 432, 270, 434, 433, 341, 334, 140, 353,
	379, 269, 330, 329, 359, 361, 265, 354, 352, 132,
	133, 134, 353, 136, 137, 328, 260, 330, 329, 217,
	354, 362, 251, 203, 141, 323, 266, 183, 369, 295,
	232, 140, 263, 232, 250, 251, 160, 381, 223, 158,
	159, 372, 369, 223, 373, 428, 165, 223, 228, 69,
	100, 170, 236, 259, 225, 174, 235, 1, 237, 155,
	58, 57, 140, 56, 55, 54, 53, 391, 18, 17,
	16, 15, 389, 45, 343, 191, 192, 193, 194, 195,
	196, 197, 198, 252, 315, 201, 202, 397, 223, 223,
	254, 223, 344, 216, 23, 24, 25, 26, 286, 223,
	317, 223, 28, 103, 70, 104, 71, 14, 12, 223,
	11, 319, 22, 369, 10, 415, 416, 417, 242, 223,
	20, 359, 420, 13, 19, 223, 341, 341, 341, 426,
	81, 41, 72, 40, 36, 65, 66, 35, 288, 436,
	27, 289, 291, 37, 34, 0, 0, 341, 234, 440,
	441, 341, 341, 341, 442, 363, 0, 301, 67, 0,
	68, 0, 64, 63, 0, 0, 79, 80, 0, 0,
	0, 82, 83, 84, 0, 0, 0, 0, 0, 0,
	0, 78, 87, 85, 86, 89, 0, 0, 0, 0,
	0, 31, 0, 0, 0, 0, 356, 0, 358, 0,
	0, 0, 0, 0, 0, 28, 29, 70, 30, 71,
	297, 0, 0, 42, 0, 50, 0, 0, 51, 43,
	44, 105, 61, 0, 52, 0, 0, 316, 0, 0,
	0, 324, 60, 0, 316, 72, 62, 0, 65, 66,
	331, 105, 0, 46, 47, 48, 49, 0, 342, 0,
	0, 0, 105, 0, 105, 0, 360, 0, 0, 105,
	0, 67, 0, 68, 0, 64, 63, 0, 365, 0,
	0, 105, 105, 105, 105, 0, 105, 105, 105, 105,
	0, 105, 105, 105, 0, 105, 0, 81, 105, 105,
	105, 0, 376, 0, 0, 105, 409, 0, 0, 0,
	0, 88, 0, 105, 0, 0, 0, 0, 382, 383,
	0, 0, 105, 105, 418, 0, 90, 91, 0, 0,
	0, 0, 0, 79, 80, 0, 430, 0, 82, 83,
	84, 0, 0, 394, 395, 0, 316, 438, 78, 87,
	85, 86, 89, 0, 105, 0, 105, 403, 404, 405,
	406, 0, 105, 0, 0, 28, 103, 70, 104, 71,
	0, 0, 109, 0, 28, 103, 70, 104, 93, 0,
	0, 109, 0, 419, 0, 0, 105, 0, 342, 342,
	342, 0, 0, 0, 39, 72, 435, 0, 65, 66,
	0, 0, 0, 0, 72, 0, 439, 65, 66, 342,
	0, 234, 0, 342, 342, 342, 0, 0, 304, 0,
	0, 67, 294, 110, 108, 64, 63, 0, 105, 0,
	96, 0, 110, 0, 64, 63, 0, 0, 0, 0,
	0, 105, 0, 0, 108, 0, 0, 0, 0, 0,
	105, 105, 0, 0, 105, 108, 0, 108, 0, 0,
	0, 0, 108, 0, 0, 0, 0, 28, 103, 70,
	104, 71, 0, 0, 108, 108, 108, 108, 0, 108,
	108, 108, 108, 0, 108, 108, 108, 0, 108, 105,
	105, 108, 108, 108, 0, 0, 105, 72, 108, 0,
	65, 66, 0, 105, 0, 0, 108, 32, 0, 0,
	0, 0, 0, 234, 0, 108, 108, 0, 0, 0,
	0, 0, 105, 67, 0, 68, 0, 64, 63, 0,
	0, 0, 0, 0, 0, 0, 0, 106, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 108, 0, 108,
	0, 0, 0, 105, 0, 108, 189, 106, 105, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 106, 0,
	106, 0, 0, 0, 0, 106, 0, 0, 0, 108,
	0, 0, 0, 0, 0, 0, 0, 106, 106, 106,
	106, 0, 106, 106, 106, 106, 0, 106, 106, 106,
	0, 106, 0, 0, 106, 106, 106, 0, 0, 138,
	0, 106, 0, 105, 0, 0, 0, 0, 0, 106,
	0, 108, 0, 0, 0, 0, 0, 0, 106, 106,
	0, 0, 0, 0, 108, 0, 0, 0, 0, 0,
	0, 0, 0, 108, 108, 0, 0, 108, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	106, 0, 106, 0, 0, 0, 185, 0, 106, 0,
	0, 28, 103, 70, 104, 248, 0, 0, 109, 0,
	0, 0, 108, 108, 213, 0, 215, 0, 0, 108,
	0, 0, 106, 218, 0, 0, 108, 0, 0, 0,
	33, 72, 0, 0, 65, 66, 0, 0, 247, 0,
	0, 0, 0, 0, 0, 108, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 67, 0, 110,
	107, 64, 63, 0, 106, 0, 0, 0, 0, 0,
	0, 0, 0, 258, 0, 261, 108, 106, 0, 0,
	107, 108, 0, 0, 0, 0, 106, 106, 0, 0,
	106, 107, 0, 107, 0, 0, 0, 0, 107, 0,
	0, 0, 0, 0, 0, 0, 279, 280, 0, 0,
	107, 107, 107, 107, 0, 107, 107, 107, 107, 0,
	107, 107, 107, 0, 107, 106, 106, 107, 107, 107,
	0, 0, 106, 0, 107, 0, 108, 0, 0, 106,
	0, 0, 107, 0, 0, 0, 0, 0, 0, 0,
	0, 107, 107, 0, 0, 0, 0, 0, 106, 0,
	327, 0, 0, 0, 0, 332, 0, 0, 0, 0,
	336, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 107, 0, 107, 0, 0, 0, 106,
	0, 107, 0, 0, 106, 28, 29, 70, 30, 71,
	0, 0, 0, 42, 333, 50, 278, 277, 51, 43,
	44, 0, 61, 0, 52, 107, 0, 377, 378, 0,
	0, 0, 60, 0, 380, 72, 62, 0, 65, 66,
	0, 0, 0, 46, 47, 48, 49, 388, 0, 390,
	0, 221, 222, 0, 28, 239, 70, 104, 71, 106,
	0, 67, 0, 68, 0, 64, 63, 107, 0, 0,
	0, 0, 0, 398, 0, 0, 0, 0, 0, 218,
	107, 0, 0, 0, 72, 0, 408, 65, 66, 107,
	107, 0, 0, 107, 0, 0, 414, 0, 279, 280,
	234, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	67, 0, 68, 0, 64, 63, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 107, 107,
	0, 0, 0, 0, 0, 107, 0, 0, 0, 0,
	0, 0, 107, 28, 29, 70, 30, 71, 0, 0,
	0, 42, 425, 345, 424, 423, 346, 43, 44, 0,
	61, 107, 52, 0, 0, 347, 348, 0, 0, 0,
	60, 0, 0, 72, 62, 0, 65, 66, 0, 0,
	0, 46, 47, 48, 49, 0, 0, 0, 0, 339,
	340, 0, 107, 0, 0, 0, 0, 107, 0, 67,
	0, 68, 0, 64, 63, 0, 0, 28, 29, 70,
	30, 71, 0, 0, 0, 42, 421, 345, 424, 423,
	346, 43, 44, 0, 61, 0, 52, 0, 0, 347,
	348, 0, 0, 0, 60, 0, 0, 72, 62, 0,
	65, 66, 0, 0, 0, 46, 47, 48, 49, 0,
	0, 0, 107, 339, 340, 0, 0, 0, 0, 0,
	0, 0, 0, 67, 0, 68, 0, 64, 63, 28,
	29, 70, 30, 71, 0, 0, 0, 42, 427, 345,
	0, 0, 346, 43, 44, 0, 61, 0, 52, 0,
	0, 347, 348, 0, 0, 0, 60, 0, 0, 72,
	62, 0, 65, 66, 0, 0, 0, 46, 47, 48,
	49, 0, 0, 0, 0, 339, 340, 0, 0, 0,
	0, 0, 0, 0, 0, 67, 0, 68, 0, 64,
	63, 28, 29, 70, 30, 71, 0, 0, 0, 42,
	396, 50, 0, 0, 51, 43, 44, 0, 61, 0,
	52, 283, 0, 0, 0, 0, 0, 320, 60, 0,
	0, 72, 62, 0, 65, 66, 0, 0, 0, 46,
	47, 48, 49, 0, 0, 0, 0, 313, 314, 0,
	0, 0, 0, 0, 0, 0, 0, 67, 0, 68,
	0, 64, 63, 28, 29, 70, 30, 71, 0, 0,
	0, 42, 338, 345, 0, 0, 346, 43, 44, 0,
	61, 0, 52, 0, 0, 347, 348, 0, 0, 0,
	60, 0, 0, 72, 62, 0, 65, 66, 0, 0,
	0, 46, 47, 48, 49, 0, 0, 0, 0, 339,
	340, 0, 0, 0, 0, 0, 0, 0, 0, 67,
	0, 68, 0, 64, 63, 28, 29, 70, 30, 71,
	0, 0, 0, 42, 325, 50, 0, 0, 51, 43,
	44, 0, 61, 0, 52, 283, 0, 0, 0, 0,
	0, 320, 60, 0, 0, 72, 62, 0, 65, 66,
	0, 0, 0, 46, 47, 48, 49, 0, 0, 0,
	0, 313, 314, 0, 0, 0, 0, 0, 0, 0,
	0, 67, 0, 68, 0, 64, 63, 28, 29, 70,
	30, 71, 0, 0, 0, 42, 311, 50, 0, 0,
	51, 43, 44, 0, 61, 0, 52, 283, 0, 0,
	0, 0, 0, 320, 60, 0, 0, 72, 62, 0,
	65, 66, 0, 0, 0, 46, 47, 48, 49, 0,
	0, 0, 0, 313, 314, 0, 0, 0, 0, 0,
	0, 0, 0, 67, 0, 68, 0, 64, 63, 28,
	29, 70, 30, 71, 0, 0, 0, 42, 275, 50,
	278, 277, 51, 43, 44, 0, 61, 0, 52, 0,
	0, 0, 0, 0, 0, 0, 60, 0, 0, 72,
	62, 0, 65, 66, 0, 0, 0, 46, 47, 48,
	49, 0, 0, 0, 0, 221, 222, 0, 28, 29,
	70, 30, 71, 0, 0, 67, 42, 68, 50, 64,
	63, 51, 43, 44, 0, 61, 0, 52, 0, 0,
	0, 0, 0, 0, 0, 60, 0, 0, 72, 62,
	0, 65, 66, 0, 0, 0, 46, 47, 48, 49,
	0, 0, 0, 0, 6, 7, 0, 0, 0, 0,
	0, 0, 0, 0, 67, 0, 68, 0, 64, 63,
	0, 8, 28, 29, 70, 30, 71, 0, 0, 0,
	42, 0, 345, 0, 0, 346, 43, 44, 0, 61,
	0, 52, 0, 0, 347, 348, 0, 0, 0, 60,
	0, 0, 72, 62, 0, 65, 66, 0, 0, 0,
	46, 47, 48, 49, 0, 0, 0, 0, 339, 340,
	0, 0, 0, 0, 0, 0, 0, 0, 67, 0,
	68, 0, 64, 63, 28, 29, 70, 30, 71, 0,
	0, 0, 42, 407, 50, 0, 0, 51, 43, 44,
	0, 61, 0, 52, 0, 0, 0, 0, 0, 0,
	0, 60, 0, 0, 72, 62, 0, 65, 66, 0,
	0, 0, 46, 47, 48, 49, 0, 0, 0, 0,
	221, 222, 0, 0, 0, 0, 0, 0, 0, 0,
	67, 0, 68, 0, 64, 63, 28, 29, 70, 30,
	71, 0, 0, 0, 42, 375, 50, 0, 0, 51,
	43, 44, 0, 61, 0, 52, 0, 0, 0, 0,
	0, 0, 0, 60, 0, 0, 72, 62, 0, 65,
	66, 0, 0, 0, 46, 47, 48, 49, 0, 0,
	0, 0, 221, 222, 0, 0, 0, 0, 0, 0,
	0, 0, 67, 0, 68, 0, 64, 63, 28, 29,
	70, 30, 71, 0, 0, 0, 42, 335, 50, 0,
	0, 51, 43, 44, 0, 61, 0, 52, 0, 0,
	0, 0, 0, 0, 0, 60, 0, 0, 72, 62,
	0, 65, 66, 0, 0, 0, 46, 47, 48, 49,
	0, 0, 0, 0, 221, 222, 0, 28, 29, 70,
	30, 71, 0, 0, 67, 42, 68, 50, 64, 63,
	51, 43, 44, 0, 61, 0, 52, 0, 0, 0,
	0, 0, 0, 0, 60, 0, 0, 72, 62, 0,
	65, 66, 0, 0, 0, 46, 47, 48, 49, 0,
	0, 0, 0, 221, 222, 0, 0, 0, 0, 0,
	0, 0, 0, 67, 0, 68, 310, 64, 63, 28,
	29, 70, 30, 71, 0, 0, 0, 42, 307, 50,
	0, 0, 51, 43, 44, 0, 61, 0, 52, 0,
	0, 0, 0, 0, 0, 0, 60, 0, 0, 72,
	62, 0, 65, 66, 0, 0, 0, 46, 47, 48,
	49, 0, 0, 0, 0, 221, 222, 0, 0, 0,
	0, 0, 0, 0, 0, 67, 0, 68, 0, 64,
	63, 28, 29, 70, 30, 71, 0, 0, 0, 42,
	274, 50, 0, 0, 51, 43, 44, 0, 61, 0,
	52, 0, 0, 0, 0, 0, 0, 0, 60, 0,
	0, 72, 62, 0, 65, 66, 0, 0, 0, 46,
	47, 48, 49, 0, 0, 0, 0, 221, 222, 0,
	0, 0, 0, 0, 0, 0, 0, 67, 0, 68,
	0, 64, 63, 28, 29, 70, 30, 71, 0, 0,
	0, 42, 272, 50, 0, 0, 51, 43, 44, 0,
	61, 0, 52, 0, 0, 0, 0, 0, 0, 0,
	60, 0, 0, 72, 62, 0, 65, 66, 0, 0,
	0, 46, 47, 48, 49, 0, 0, 0, 0, 221,
	222, 0, 28, 29, 70, 30, 71, 0, 0, 67,
	42, 68, 50, 64, 63, 51, 43, 44, 0, 61,
	0, 52, 0, 0, 0, 0, 0, 0, 0, 60,
	0, 0, 72, 62, 0, 65, 66, 0, 0, 0,
	46, 47, 48, 49, 0, 0, 0, 0, 221, 222,
	0, 0, 0, 0, 0, 0, 0, 0, 67, 0,
	68, 262, 64, 63, 28, 29, 70, 30, 71, 0,
	0, 0, 42, 257, 50, 0, 0, 51, 43, 44,
	0, 61, 0, 52, 0, 0, 0, 0, 0, 0,
	0, 60, 0, 0, 72, 62, 0, 65, 66, 0,
	0, 0, 46, 47, 48, 49, 0, 0, 0, 0,
	221, 222, 0, 28, 29, 70, 30, 71, 0, 0,
	67, 42, 68, 50, 64, 63, 51, 43, 44, 0,
	61, 0, 52, 0, 0, 0, 0, 0, 0, 0,
	60, 0, 0, 72, 62, 0, 65, 66, 0, 0,
	0, 46, 47, 48, 49, 0, 0, 0, 0, 221,
	222, 0, 28, 29, 70, 30, 71, 244, 0, 67,
	42, 68, 50, 64, 63, 51, 43, 44, 0, 61,
	0, 52, 0, 28, 103, 70, 104, 71, 0, 60,
	109, 0, 72, 62, 0, 65, 66, 0, 0, 0,
	46, 47, 48, 49, 28, 103, 70, 104, 93, 243,
	0, 109, 0, 72, 0, 0, 65, 66, 67, 0,
	68, 0, 64, 63, 28, 103, 70, 104, 71, 0,
	0, 0, 0, 0, 72, 0, 0, 65, 66, 67,
	0, 110, 0, 64, 63, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 72, 0, 0, 65, 66, 0,
	96, 0, 110, 0, 64, 63, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	67, 0, 68, 0, 64, 63,
}
var RubyPact = []int{

	-31, 1763, -1000, -1000, -1000, 4, -1000, -1000, -1000, 573,
	-1000, -1000, -1000, 28, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 135,
	180, 74, 0, -1, -1000, -1000, -1000, -1000, -1000, 220,
	-21, -1000, 273, 175, 175, 122, 490, 490, 490, 2519,
	490, 490, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	2519, 490, 2519, 18, 250, -1000, -1000, 2519, -1000, -15,
	170, -1000, 15, -1000, -1000, -1000, 490, 490, 320, 2519,
	2519, 2519, 2519, 490, 2519, 2519, 2519, 2519, 490, 2519,
	2519, 2519, 490, 2519, -1000, 106, 2519, 2519, 2519, 311,
	151, 171, -1000, 2499, 162, -1000, -1000, -1000, 24, -24,
	-24, 2519, 490, 490, 490, 490, 490, 490, 490, 490,
	2519, 2519, 490, 490, 307, 101, 53, 11, -1000, -1000,
	490, 303, 130, 130, 130, 171, 130, 195, 2408, 49,
	171, 109, 171, -1000, -1000, 236, -1000, -1000, 33, 17,
	416, -1000, 1189, 183, 2519, 2457, -1000, -24, 130, 130,
	946, 171, 171, 171, 171, 130, 171, 171, 171, 171,
	130, 97, 171, 171, 130, 314, 416, -1000, 147, 34,
	-1000, 34, -1000, -1000, 2478, 2359, -1000, 300, -1000, 2297,
	312, 130, 130, 130, 130, 130, 130, 130, 130, 171,
	171, 130, 130, -1000, -1000, 290, 143, 19, -1000, 5,
	285, 277, -1000, 2248, 175, 2186, 130, -1000, 1714, -1000,
	-1000, -1000, -1000, 573, 130, 241, 2519, -1000, 6, -1000,
	-1000, -1000, -1000, -1000, -1000, 138, 114, -11, 177, 649,
	-1000, 309, 130, -1000, -1000, 106, 15, 490, 2519, 2519,
	15, -1000, 640, 3, -22, 171, -1000, -1000, 2124, 47,
	-1000, 2062, -1000, -1000, 1652, 53, 19, 305, 490, -1000,
	-1000, 1590, -1000, -1000, -1000, -1000, 291, 490, -1000, 1140,
	2013, -1000, -1000, 110, 171, 1528, 199, 2519, 742, -29,
	-1000, -33, -1000, 490, 2519, -1000, -1000, 130, 301, 171,
	-1000, 387, 171, -1000, 263, 490, 258, -1000, -1000, 245,
	-1000, -1000, 223, -1000, -1000, 573, 130, -1000, -1000, 253,
	2519, -1000, -1000, -1000, 130, -1000, 214, 1951, -1000, 490,
	-1000, 130, 2408, -1000, 276, -1000, 2408, 323, -1000, -1000,
	-1000, 573, 130, -1000, -1000, 490, 490, 205, 52, -1000,
	-1000, 2519, 49, 416, -1000, -1000, 742, -1000, 107, 573,
	130, 171, -1000, 201, -1000, 130, -1000, -1000, -1000, -1000,
	490, 490, 49, 1466, -1000, -1000, 130, 2408, 2408, -1000,
	2408, 193, 58, 182, 490, 490, 490, 490, 1889, 49,
	2408, 127, -20, 59, 130, 130, -1000, 105, 2408, -1000,
	-1000, -1000, -1000, 130, 130, 130, 130, -1000, 2408, -11,
	490, 2519, -1000, -1000, 2408, 1342, 1278, 1404, -11, 230,
	125, -1000, 268, 490, -1000, -1000, 118, -1000, -1000, -1000,
	-11, -1000, -1000, 490, -1000, 130, 1827, -1000, -11, 130,
	1827, 1827, 1827,
}
var RubyPgo = []int{

	0, 252, 434, 433, 430, 23, 427, 424, 423, 980,
	421, 5, 16, 414, 413, 410, 0, 787, 674, 404,
	402, 401, 8, 400, 6, 481, 398, 397, 9, 390,
	388, 387, 386, 385, 384, 382, 364, 2, 363, 361,
	360, 359, 358, 356, 355, 354, 353, 351, 350, 836,
	349, 3, 15, 17, 7, 347, 1, 346, 4, 344,
	14, 12, 343, 10, 342, 340, 11, 13, 339, 335,
	195,
}
var RubyR1 = []int{

	0, 55, 55, 55, 55, 55, 55, 55, 55, 55,
	55, 69, 69, 70, 70, 49, 49, 49, 49, 49,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 22, 22, 22, 22, 22, 22,
	22, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	18, 18, 18, 18, 18, 18, 18, 18, 18, 18,
	18, 18, 18, 18, 18, 18, 18, 18, 18, 18,
	52, 52, 52, 52, 63, 63, 60, 60, 60, 60,
	60, 66, 66, 66, 66, 66, 65, 65, 65, 19,
	19, 19, 19, 19, 19, 29, 29, 29, 29, 61,
	61, 61, 61, 61, 61, 56, 56, 24, 24, 24,
	24, 67, 67, 67, 23, 23, 26, 28, 28, 68,
	68, 14, 14, 14, 14, 14, 14, 14, 27, 27,
	27, 27, 27, 27, 9, 9, 25, 25, 17, 17,
	38, 38, 39, 40, 41, 42, 43, 44, 45, 46,
	47, 48, 2, 6, 7, 7, 3, 3, 57, 57,
	57, 57, 64, 64, 64, 5, 5, 5, 5, 53,
	62, 62, 62, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 54, 54, 54, 54, 50, 50, 50,
	8, 15, 11, 11, 11, 59, 59, 51, 51, 20,
	21, 12, 34, 58, 58, 58, 58, 58, 58, 58,
	35, 35, 35, 35, 35, 35, 35, 36, 36, 36,
	36, 36, 37, 37, 37, 37, 33, 32, 10, 31,
	31, 30, 30, 4,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 1, 0, 2, 0, 2, 2, 2, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	4, 5, 1, 4, 4, 2, 3, 4, 4, 5,
	3, 5, 2, 3, 3, 3, 3, 3, 4, 6,
	3, 7, 1, 5, 1, 3, 0, 1, 1, 4,
	4, 1, 1, 4, 4, 5, 1, 3, 3, 5,
	6, 7, 8, 5, 6, 0, 1, 3, 3, 0,
	2, 2, 2, 2, 2, 1, 3, 1, 2, 3,
	2, 0, 1, 3, 4, 6, 4, 1, 3, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 2, 2, 2, 3, 3,
	3, 3, 2, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 1, 1, 3, 3, 5, 5, 0, 1,
	3, 7, 3, 7, 8, 3, 4, 4, 3, 3,
	0, 1, 3, 4, 5, 3, 3, 3, 3, 3,
	5, 6, 5, 4, 3, 3, 2, 0, 2, 2,
	3, 4, 2, 3, 5, 0, 2, 1, 2, 2,
	2, 5, 5, 0, 2, 2, 2, 2, 2, 2,
	0, 1, 3, 3, 1, 3, 3, 5, 6, 5,
	6, 5, 4, 3, 3, 2, 3, 3, 2, 5,
	7, 4, 5, 3,
}
var RubyChk = []int{

	-1000, -55, 51, 52, 68, -1, 51, 52, 68, -16,
	-19, -23, -26, -14, -27, -39, -40, -41, -42, -13,
	-15, -22, -20, -34, -33, -32, -31, -4, 5, 6,
	8, -25, -17, -9, -2, -6, -7, -3, -28, -18,
	-8, -10, 13, 19, 20, -38, 43, 44, 45, 46,
	15, 18, 24, -43, -44, -45, -46, -47, -48, -12,
	32, 22, 36, 66, 65, 38, 39, 61, 63, -68,
	7, 9, 35, 52, 51, 68, 15, 48, 55, 40,
	41, 4, 45, 46, 47, 57, 58, 56, 18, 59,
	33, 34, 18, 9, -52, -66, 61, 42, 49, 11,
	-65, -16, -5, 6, 8, -25, -17, -9, -18, 12,
	63, 9, 42, 49, 42, 49, 42, 49, 42, 49,
	40, 41, 15, 18, 55, 6, 4, -28, 8, -28,
	42, 11, -1, -1, -1, -16, -1, -1, -49, -63,
	-16, -1, -16, 6, 8, 66, 6, 8, -63, -60,
	-16, -22, -70, 54, 9, -50, -5, 63, -1, -1,
	6, -16, -16, -16, -16, -1, -16, -16, -16, -16,
	-1, -16, -16, -16, -1, -60, -16, 11, -16, -16,
	-12, -16, -12, 6, 11, -49, -53, 56, -53, -49,
	-60, -1, -1, -1, -1, -1, -1, -1, -1, -16,
	-16, -1, -1, 6, -56, 55, -67, 9, -24, 6,
	47, 58, -56, -49, 40, -49, -1, 6, -49, 51,
	52, 51, 52, -16, -1, -59, 11, 51, -70, 6,
	8, 62, 11, 62, 51, -57, -64, -70, -16, 6,
	8, -60, -1, 52, 10, -66, -52, 42, 9, 53,
	10, 11, -70, 62, -70, -16, -5, 14, -49, -62,
	6, -49, 64, 10, -61, 6, 11, -67, 42, 6,
	6, -61, 14, -28, 14, 14, -54, 17, 16, -49,
	-49, 14, -11, 25, -16, -58, -30, 37, -70, -70,
	11, -70, 11, 4, 53, 10, -5, -1, -60, -16,
	-5, -70, -16, -5, 58, 42, 58, 14, 56, 11,
	64, 14, -51, 51, 52, -16, -1, -29, -11, -21,
	31, -56, -24, 10, -1, 14, -51, -49, 14, 17,
	16, -1, -49, 14, -54, 14, -49, 8, 14, 51,
	52, -16, -1, -36, -35, 15, 18, 27, 28, 14,
	16, 37, -63, -16, -22, 64, -70, 64, -70, -16,
	-1, -16, 10, 58, 6, -1, 6, 6, 14, -11,
	15, 18, -63, -61, 14, 14, -1, -49, -49, 14,
	-49, 4, -1, -1, 15, 18, 15, 18, -49, -63,
	-49, -16, 6, 6, -1, -1, 14, -51, -49, 6,
	51, 51, 52, -1, -1, -1, -1, 14, -49, -70,
	4, 53, 10, 14, -49, -58, -58, -58, -70, -1,
	-16, 14, -37, 17, 16, 14, -37, 14, -69, 11,
	-70, 11, 14, 17, 16, -1, -58, 14, -70, -1,
	-58, -58, -58,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 20, 21,
	-2, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 15, 34, 35, 36, 37, 38, 39, 40,
	0, 0, 0, 0, 0, 162, 163, 86, 13, 0,
	62, 197, 0, 5, 6, 7, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, -2, 65, 72, 86, 0, 0, 0,
	82, 91, 92, 21, -2, 23, 24, 25, 31, 15,
	-2, 86, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 121, 121, 15, -2, 15,
	0, 0, 152, 153, 154, 155, 15, 0, 205, 209,
	84, 0, 13, 146, 147, 0, 144, 145, 0, 0,
	84, 88, 13, 0, 86, 0, 238, 15, 185, 243,
	66, 73, 75, 77, 156, 157, 158, 159, 160, 161,
	187, 0, 236, 237, 189, 0, 87, 13, 84, 131,
	132, 138, 139, 150, 13, 0, 15, 180, 15, 0,
	0, 133, 140, 134, 141, 135, 142, 136, 143, 74,
	76, 186, 188, 70, 109, 0, 115, 121, 122, 117,
	0, 0, 109, 0, 0, 0, 137, 151, 0, 15,
	15, 16, 17, -2, 19, 0, 0, 213, 0, 148,
	149, 164, 13, 165, 14, 13, 13, 169, 0, 21,
	-2, 0, 198, 199, 200, 67, 68, 0, -2, 0,
	60, 13, 0, 78, 0, 97, 98, 175, 0, 0,
	181, 0, 178, 64, 0, 121, 0, 0, 0, 118,
	120, 0, 124, 15, 126, 183, 0, 0, 15, 0,
	0, 201, 206, 15, 85, 0, 0, 0, 0, 0,
	13, 0, 13, 0, 0, 63, 69, 71, 0, 211,
	61, 0, 93, 94, 0, 0, 0, 176, 179, 0,
	177, 99, 0, 110, 111, 41, 113, 114, 207, 106,
	0, 109, 123, 116, 119, 103, 0, 0, 184, 0,
	15, 15, 196, 190, 0, 192, 202, 15, 212, 214,
	215, 41, 217, 218, 219, 0, 0, 221, 224, 239,
	15, 0, 15, 89, 90, 166, 0, 167, 0, 41,
	170, 172, 80, 0, 95, 79, 83, 182, 100, 208,
	0, 0, 210, 0, 104, 125, 15, 194, 195, 191,
	203, 0, 15, 0, 0, 0, 0, 0, 0, 15,
	13, 0, 0, 0, 107, 108, 101, 0, 193, 15,
	213, -2, -2, 222, 223, 225, 226, 240, 13, 241,
	0, 0, 81, 102, 204, 0, 0, 0, 242, 11,
	13, 227, 0, 0, 213, 229, 0, 231, 171, 12,
	173, 13, 228, 0, 213, 213, 220, 230, 174, 213,
	220, 220, 220,
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
		//line parser.y:199
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:201
		{
		}
	case 3:
		//line parser.y:203
		{
		}
	case 4:
		//line parser.y:205
		{
		}
	case 5:
		//line parser.y:207
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:209
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:211
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:217
		{
		}
	case 11:
		//line parser.y:219
		{
		}
	case 12:
		//line parser.y:220
		{
		}
	case 13:
		//line parser.y:222
		{
		}
	case 14:
		//line parser.y:223
		{
		}
	case 15:
		//line parser.y:226
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 16:
		//line parser.y:228
		{
		}
	case 17:
		//line parser.y:230
		{
		}
	case 18:
		//line parser.y:232
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 19:
		//line parser.y:234
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
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
		//line parser.y:244
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 61:
		//line parser.y:251
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-2].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 62:
		//line parser.y:258
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 63:
		//line parser.y:262
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 64:
		//line parser.y:269
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 65:
		//line parser.y:276
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 66:
		//line parser.y:283
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 67:
		//line parser.y:290
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 68:
		//line parser.y:298
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 69:
		//line parser.y:306
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 70:
		//line parser.y:314
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 71:
		//line parser.y:322
		{
			methodName := RubyS[Rubypt-2].genericValue.(ast.BareReference).Name + "="
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: methodName},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 72:
		//line parser.y:333
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 73:
		//line parser.y:340
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 74:
		//line parser.y:348
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 75:
		//line parser.y:356
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 76:
		//line parser.y:364
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 77:
		//line parser.y:372
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 78:
		//line parser.y:382
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 79:
		//line parser.y:392
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 80:
		//line parser.y:401
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 81:
		//line parser.y:403
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 82:
		//line parser.y:405
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 83:
		//line parser.y:407
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 84:
		//line parser.y:410
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:412
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 86:
		//line parser.y:414
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 87:
		//line parser.y:416
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 88:
		//line parser.y:418
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 89:
		//line parser.y:420
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 90:
		//line parser.y:422
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 91:
		//line parser.y:426
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 92:
		//line parser.y:428
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 93:
		//line parser.y:430
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 94:
		//line parser.y:432
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-1].genericValue)
		}
	case 95:
		//line parser.y:434
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 96:
		//line parser.y:437
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 97:
		//line parser.y:439
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 98:
		//line parser.y:441
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 99:
		//line parser.y:447
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 100:
		//line parser.y:455
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 101:
		//line parser.y:464
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 102:
		//line parser.y:473
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 103:
		//line parser.y:483
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 104:
		//line parser.y:491
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 105:
		//line parser.y:501
		{
		}
	case 106:
		//line parser.y:503
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 107:
		//line parser.y:505
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{RubyS[Rubypt-2].genericValue}}
		}
	case 108:
		//line parser.y:507
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{ast.Break{}},
			}
		}
	case 109:
		//line parser.y:515
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 110:
		//line parser.y:517
		{
		}
	case 111:
		//line parser.y:519
		{
		}
	case 112:
		//line parser.y:521
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 113:
		//line parser.y:523
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 114:
		//line parser.y:525
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 115:
		//line parser.y:530
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 116:
		//line parser.y:532
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 117:
		//line parser.y:535
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 118:
		//line parser.y:537
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 119:
		//line parser.y:539
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 120:
		//line parser.y:541
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 121:
		//line parser.y:543
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 122:
		//line parser.y:545
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 123:
		//line parser.y:549
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 124:
		//line parser.y:554
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 125:
		//line parser.y:561
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 126:
		//line parser.y:571
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 127:
		//line parser.y:580
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 128:
		//line parser.y:586
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 129:
		//line parser.y:594
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 130:
		//line parser.y:598
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 131:
		//line parser.y:603
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 132:
		//line parser.y:610
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 133:
		//line parser.y:614
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 134:
		//line parser.y:621
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 135:
		//line parser.y:628
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 136:
		//line parser.y:635
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 137:
		//line parser.y:642
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 138:
		//line parser.y:650
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 139:
		//line parser.y:657
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 140:
		//line parser.y:661
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 141:
		//line parser.y:668
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 142:
		//line parser.y:675
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 143:
		//line parser.y:682
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 144:
		//line parser.y:690
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 145:
		//line parser.y:692
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 146:
		//line parser.y:695
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 147:
		//line parser.y:697
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 148:
		//line parser.y:700
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 149:
		//line parser.y:702
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 150:
		//line parser.y:705
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 151:
		//line parser.y:707
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 152:
		//line parser.y:709
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 153:
		//line parser.y:710
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 154:
		//line parser.y:711
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 155:
		//line parser.y:712
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 156:
		//line parser.y:715
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 157:
		//line parser.y:724
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 158:
		//line parser.y:733
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 159:
		//line parser.y:742
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 160:
		//line parser.y:751
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 161:
		//line parser.y:760
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 162:
		//line parser.y:768
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 163:
		//line parser.y:769
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 164:
		//line parser.y:771
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 165:
		//line parser.y:773
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 166:
		//line parser.y:776
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 167:
		//line parser.y:784
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 168:
		//line parser.y:792
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 169:
		//line parser.y:794
		{
		}
	case 170:
		//line parser.y:796
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 171:
		//line parser.y:803
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 172:
		//line parser.y:811
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 173:
		//line parser.y:818
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 174:
		//line parser.y:825
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 175:
		//line parser.y:833
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 176:
		//line parser.y:835
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 177:
		//line parser.y:842
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 178:
		//line parser.y:849
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 179:
		//line parser.y:852
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 180:
		//line parser.y:854
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 181:
		//line parser.y:856
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 182:
		//line parser.y:858
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 183:
		//line parser.y:861
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 184:
		//line parser.y:868
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 185:
		//line parser.y:876
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 186:
		//line parser.y:883
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 187:
		//line parser.y:890
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 188:
		//line parser.y:897
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 189:
		//line parser.y:904
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 190:
		//line parser.y:911
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 191:
		//line parser.y:918
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 192:
		//line parser.y:926
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 193:
		//line parser.y:935
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 194:
		//line parser.y:942
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 195:
		//line parser.y:949
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 196:
		//line parser.y:956
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 197:
		//line parser.y:963
		{
		}
	case 198:
		//line parser.y:964
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 199:
		//line parser.y:965
		{
		}
	case 200:
		//line parser.y:968
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 201:
		//line parser.y:971
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 202:
		//line parser.y:979
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 203:
		//line parser.y:981
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 204:
		//line parser.y:990
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
	case 205:
		//line parser.y:1005
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 206:
		//line parser.y:1007
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 207:
		//line parser.y:1010
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 208:
		//line parser.y:1012
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 209:
		//line parser.y:1015
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 210:
		//line parser.y:1024
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 211:
		//line parser.y:1033
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 212:
		//line parser.y:1042
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 213:
		//line parser.y:1045
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 214:
		//line parser.y:1047
		{
		}
	case 215:
		//line parser.y:1049
		{
		}
	case 216:
		//line parser.y:1051
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 217:
		//line parser.y:1053
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 218:
		//line parser.y:1055
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 219:
		//line parser.y:1057
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 220:
		//line parser.y:1059
		{
		}
	case 221:
		//line parser.y:1061
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 222:
		//line parser.y:1063
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 223:
		//line parser.y:1065
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 224:
		//line parser.y:1067
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 225:
		//line parser.y:1069
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 226:
		//line parser.y:1071
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 227:
		//line parser.y:1074
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 228:
		//line parser.y:1081
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 229:
		//line parser.y:1089
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 230:
		//line parser.y:1096
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 231:
		//line parser.y:1104
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 232:
		//line parser.y:1112
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 233:
		//line parser.y:1119
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 234:
		//line parser.y:1126
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 235:
		//line parser.y:1133
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 236:
		//line parser.y:1141
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 237:
		//line parser.y:1144
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 238:
		//line parser.y:1146
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 239:
		//line parser.y:1149
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 240:
		//line parser.y:1151
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 241:
		//line parser.y:1154
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 242:
		//line parser.y:1156
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 243:
		//line parser.y:1158
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
