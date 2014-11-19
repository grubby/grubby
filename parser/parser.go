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

//line parser.y:1188

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 42,
	54, 130,
	-2, 21,
	-1, 109,
	54, 130,
	-2, 128,
	-1, 111,
	10, 95,
	11, 95,
	-2, 200,
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
	54, 130,
	-2, 21,
	-1, 257,
	54, 131,
	-2, 129,
	-1, 267,
	10, 95,
	11, 95,
	-2, 200,
	-1, 377,
	61, 82,
	-2, 48,
	-1, 417,
	27, 217,
	28, 217,
	-2, 15,
	-1, 418,
	27, 217,
	28, 217,
	-2, 15,
}

const RubyNprod = 247
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 3059

var RubyAct = []int{

	242, 327, 5, 326, 309, 438, 194, 300, 149, 198,
	196, 113, 27, 17, 150, 23, 112, 22, 287, 223,
	82, 2, 3, 83, 251, 165, 251, 166, 122, 209,
	122, 56, 120, 101, 222, 132, 105, 375, 4, 373,
	103, 248, 244, 104, 133, 343, 155, 427, 141, 142,
	251, 134, 108, 110, 102, 251, 80, 79, 88, 146,
	311, 82, 340, 132, 83, 101, 144, 199, 410, 148,
	197, 160, 161, 81, 251, 140, 204, 199, 152, 159,
	154, 159, 129, 170, 171, 167, 102, 97, 98, 345,
	342, 178, 249, 247, 86, 87, 183, 417, 418, 9,
	137, 188, 158, 100, 191, 192, 139, 138, 200, 85,
	339, 324, 276, 251, 202, 130, 298, 82, 200, 201,
	83, 82, 131, 244, 83, 152, 206, 316, 152, 201,
	212, 226, 227, 214, 229, 230, 231, 232, 233, 234,
	235, 119, 221, 152, 225, 208, 210, 215, 217, 135,
	428, 385, 382, 238, 239, 305, 136, 145, 259, 147,
	145, 307, 307, 307, 151, 314, 307, 274, 88, 219,
	152, 258, 162, 163, 164, 380, 284, 82, 185, 186,
	83, 330, 274, 129, 172, 264, 174, 175, 176, 177,
	265, 179, 180, 181, 182, 156, 184, 97, 98, 187,
	355, 189, 190, 82, 86, 87, 83, 88, 275, 109,
	366, 207, 367, 416, 211, 213, 216, 281, 453, 85,
	450, 449, 199, 445, 119, 197, 271, 82, 246, 207,
	83, 415, 228, 368, 257, 254, 97, 98, 403, 245,
	88, 404, 398, 86, 87, 306, 391, 390, 89, 90,
	91, 99, 290, 250, 255, 262, 207, 263, 85, 94,
	92, 93, 96, 200, 168, 286, 169, 321, 388, 97,
	98, 195, 401, 119, 201, 402, 86, 87, 319, 248,
	328, 152, 322, 333, 329, 269, 270, 297, 248, 294,
	285, 85, 283, 284, 88, 96, 444, 210, 320, 278,
	288, 448, 349, 450, 449, 396, 277, 348, 347, 273,
	359, 352, 107, 236, 106, 84, 335, 218, 378, 289,
	369, 193, 88, 97, 98, 381, 371, 173, 383, 312,
	86, 87, 66, 313, 315, 383, 346, 118, 348, 347,
	389, 253, 293, 243, 308, 85, 392, 252, 393, 96,
	371, 97, 98, 1, 157, 55, 119, 426, 86, 87,
	54, 53, 88, 399, 400, 336, 52, 207, 323, 447,
	51, 268, 50, 85, 34, 33, 331, 406, 32, 31,
	46, 360, 361, 19, 36, 337, 97, 98, 412, 37,
	20, 97, 98, 86, 87, 374, 310, 376, 86, 87,
	14, 12, 419, 420, 421, 422, 43, 11, 85, 21,
	18, 145, 370, 85, 383, 10, 24, 377, 379, 16,
	13, 431, 432, 433, 35, 15, 30, 435, 29, 25,
	63, 26, 359, 359, 359, 62, 370, 0, 442, 0,
	451, 0, 224, 0, 0, 452, 88, 0, 126, 0,
	455, 0, 0, 359, 0, 456, 457, 359, 359, 359,
	458, 0, 0, 0, 126, 0, 126, 126, 145, 0,
	0, 126, 408, 0, 409, 97, 98, 0, 0, 126,
	126, 126, 86, 87, 0, 0, 0, 408, 425, 0,
	0, 126, 143, 126, 126, 126, 126, 85, 126, 126,
	126, 126, 0, 126, 317, 434, 126, 0, 126, 126,
	0, 0, 0, 119, 0, 0, 28, 446, 126, 0,
	0, 126, 126, 126, 0, 0, 377, 436, 454, 0,
	0, 126, 0, 97, 98, 0, 126, 0, 0, 126,
	86, 87, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 203, 0, 205, 0, 85, 0, 0, 121, 0,
	126, 126, 0, 126, 0, 220, 0, 0, 0, 0,
	0, 0, 0, 0, 121, 0, 121, 121, 0, 0,
	126, 121, 0, 0, 237, 0, 0, 0, 0, 121,
	121, 121, 126, 126, 0, 0, 0, 0, 0, 0,
	0, 121, 0, 121, 121, 121, 121, 0, 121, 121,
	121, 121, 0, 121, 0, 0, 121, 0, 121, 121,
	0, 0, 0, 0, 0, 0, 126, 0, 121, 0,
	0, 121, 121, 121, 0, 0, 0, 272, 0, 0,
	0, 121, 0, 0, 0, 279, 121, 0, 0, 121,
	0, 126, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 126, 292, 0, 295, 0, 0, 0,
	121, 121, 0, 121, 126, 126, 0, 0, 0, 0,
	0, 303, 304, 126, 0, 0, 0, 0, 0, 0,
	121, 0, 126, 61, 124, 67, 125, 111, 0, 117,
	122, 0, 121, 121, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 126, 126,
	0, 0, 0, 69, 334, 126, 77, 78, 0, 0,
	115, 70, 71, 72, 73, 74, 121, 116, 0, 0,
	0, 0, 0, 126, 0, 350, 0, 0, 0, 114,
	354, 123, 0, 76, 75, 0, 0, 0, 0, 0,
	0, 121, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 384, 0, 121, 0, 126, 0, 0, 0, 126,
	0, 126, 0, 0, 121, 121, 0, 0, 0, 0,
	0, 394, 395, 121, 126, 0, 0, 0, 397, 0,
	0, 0, 121, 0, 0, 0, 0, 0, 44, 0,
	405, 0, 407, 0, 0, 0, 0, 0, 0, 0,
	126, 0, 0, 0, 0, 0, 0, 0, 121, 121,
	0, 0, 0, 0, 126, 121, 414, 0, 0, 0,
	0, 0, 237, 0, 0, 0, 0, 0, 0, 424,
	127, 0, 0, 121, 0, 0, 0, 0, 430, 0,
	303, 304, 0, 0, 0, 0, 127, 0, 127, 127,
	0, 0, 0, 127, 0, 0, 0, 0, 0, 0,
	0, 127, 127, 127, 0, 121, 0, 0, 0, 121,
	0, 121, 0, 127, 0, 127, 127, 127, 127, 0,
	127, 127, 127, 127, 121, 127, 0, 0, 127, 0,
	127, 127, 61, 124, 67, 125, 111, 429, 45, 122,
	127, 0, 0, 127, 127, 127, 0, 0, 0, 0,
	121, 0, 0, 127, 0, 0, 0, 0, 127, 0,
	0, 127, 69, 0, 121, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 0, 0,
	128, 0, 127, 127, 0, 127, 0, 0, 114, 0,
	123, 0, 76, 75, 0, 0, 128, 0, 128, 128,
	0, 0, 127, 128, 0, 0, 0, 0, 0, 0,
	0, 128, 128, 128, 127, 127, 0, 0, 0, 0,
	0, 0, 0, 128, 0, 128, 128, 128, 128, 0,
	128, 128, 128, 128, 0, 128, 0, 0, 128, 0,
	128, 128, 61, 124, 67, 125, 111, 0, 127, 122,
	128, 0, 0, 128, 128, 128, 0, 0, 0, 0,
	0, 0, 0, 128, 0, 0, 0, 0, 128, 0,
	0, 128, 69, 127, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 127, 0, 0, 0, 0,
	318, 0, 128, 128, 0, 128, 127, 127, 114, 0,
	123, 0, 76, 75, 0, 127, 0, 0, 0, 0,
	0, 0, 128, 0, 127, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 128, 128, 0, 0, 0, 61,
	41, 67, 42, 68, 0, 0, 0, 38, 443, 362,
	127, 127, 363, 39, 40, 0, 58, 127, 49, 0,
	0, 364, 365, 0, 0, 60, 57, 0, 128, 69,
	59, 0, 77, 78, 0, 127, 0, 70, 71, 72,
	73, 74, 0, 0, 0, 357, 358, 0, 0, 0,
	0, 0, 0, 128, 88, 64, 0, 65, 0, 76,
	75, 0, 0, 0, 0, 128, 0, 127, 95, 0,
	0, 127, 0, 127, 0, 84, 128, 128, 0, 0,
	0, 0, 0, 97, 98, 128, 127, 0, 0, 0,
	86, 87, 0, 0, 128, 89, 90, 91, 99, 0,
	0, 0, 0, 0, 0, 85, 94, 92, 93, 96,
	0, 0, 127, 0, 0, 0, 0, 0, 0, 0,
	128, 128, 0, 0, 0, 0, 127, 128, 0, 61,
	41, 67, 42, 68, 0, 0, 0, 38, 441, 362,
	440, 439, 363, 39, 40, 128, 58, 0, 49, 0,
	0, 364, 365, 0, 0, 60, 57, 0, 0, 69,
	59, 0, 77, 78, 0, 0, 0, 70, 71, 72,
	73, 74, 0, 0, 0, 357, 358, 128, 0, 0,
	0, 128, 0, 128, 0, 64, 0, 65, 0, 76,
	75, 0, 0, 0, 0, 0, 128, 0, 0, 0,
	0, 61, 41, 67, 42, 68, 0, 0, 0, 38,
	437, 362, 440, 439, 363, 39, 40, 0, 58, 0,
	49, 0, 128, 364, 365, 0, 0, 60, 57, 0,
	0, 69, 59, 0, 77, 78, 128, 0, 0, 70,
	71, 72, 73, 74, 0, 0, 0, 357, 358, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 0, 65,
	0, 76, 75, 61, 41, 67, 42, 68, 0, 0,
	0, 38, 356, 362, 0, 0, 363, 39, 40, 0,
	58, 0, 49, 0, 0, 364, 365, 0, 0, 60,
	57, 0, 0, 69, 59, 0, 77, 78, 0, 0,
	0, 70, 71, 72, 73, 74, 0, 0, 0, 357,
	358, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	0, 65, 0, 76, 75, 61, 41, 67, 42, 68,
	0, 0, 0, 38, 351, 47, 302, 301, 48, 39,
	40, 0, 58, 0, 49, 0, 0, 0, 0, 0,
	0, 60, 57, 0, 0, 69, 59, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 0, 0,
	0, 240, 241, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 0, 65, 0, 76, 75, 61, 41, 67,
	42, 68, 0, 0, 0, 38, 299, 47, 302, 301,
	48, 39, 40, 0, 58, 0, 49, 0, 0, 0,
	0, 0, 0, 60, 57, 0, 0, 69, 59, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 240, 241, 0, 61, 41, 67, 42,
	68, 0, 0, 64, 38, 65, 47, 76, 75, 48,
	39, 40, 0, 58, 0, 49, 0, 0, 0, 0,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 6, 7, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 0, 76, 75, 0, 8,
	61, 41, 67, 42, 68, 0, 0, 0, 38, 0,
	362, 0, 0, 363, 39, 40, 0, 58, 0, 49,
	0, 0, 364, 365, 0, 0, 60, 57, 0, 0,
	69, 59, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 0, 0, 0, 357, 358, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 65, 0,
	76, 75, 61, 41, 67, 42, 68, 0, 0, 0,
	38, 411, 47, 0, 0, 48, 39, 40, 0, 58,
	0, 49, 307, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 240, 241,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 0,
	65, 0, 76, 75, 61, 41, 67, 42, 68, 0,
	0, 0, 38, 332, 47, 0, 0, 48, 39, 40,
	0, 58, 0, 49, 307, 0, 0, 0, 0, 0,
	60, 57, 0, 0, 69, 59, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 0, 0, 0,
	240, 241, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 0, 65, 0, 76, 75, 61, 41, 67, 42,
	68, 0, 0, 0, 38, 325, 47, 0, 0, 48,
	39, 40, 0, 58, 0, 49, 307, 0, 0, 0,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 240, 241, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 0, 76, 75, 61, 41,
	67, 42, 68, 0, 0, 0, 38, 423, 47, 0,
	0, 48, 39, 40, 0, 58, 0, 49, 0, 0,
	0, 0, 0, 0, 60, 57, 0, 0, 69, 59,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 240, 241, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 0, 65, 0, 76, 75,
	61, 41, 67, 42, 68, 0, 0, 0, 38, 386,
	47, 0, 0, 48, 39, 40, 0, 58, 0, 49,
	0, 0, 0, 0, 0, 0, 60, 57, 0, 0,
	69, 59, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 0, 0, 0, 240, 241, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 65, 0,
	76, 75, 61, 41, 67, 42, 68, 0, 0, 0,
	38, 353, 47, 0, 0, 48, 39, 40, 0, 58,
	0, 49, 0, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 240, 241,
	0, 61, 41, 67, 42, 68, 0, 0, 64, 38,
	65, 47, 76, 75, 48, 39, 40, 0, 58, 0,
	49, 0, 0, 0, 0, 0, 0, 60, 57, 0,
	0, 69, 59, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 0, 0, 0, 240, 241, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 0, 65,
	344, 76, 75, 61, 41, 67, 42, 68, 0, 0,
	0, 38, 341, 47, 0, 0, 48, 39, 40, 0,
	58, 0, 49, 0, 0, 0, 0, 0, 0, 60,
	57, 0, 0, 69, 59, 0, 77, 78, 0, 0,
	0, 70, 71, 72, 73, 74, 0, 0, 0, 240,
	241, 0, 61, 41, 67, 42, 68, 0, 0, 64,
	38, 65, 47, 76, 75, 48, 39, 40, 0, 58,
	0, 49, 0, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 240, 241,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 0,
	65, 296, 76, 75, 61, 41, 67, 42, 68, 0,
	0, 0, 38, 291, 47, 0, 0, 48, 39, 40,
	0, 58, 0, 49, 0, 0, 0, 0, 0, 0,
	60, 57, 0, 0, 69, 59, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 0, 0, 0,
	240, 241, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 0, 65, 0, 76, 75, 61, 41, 67, 42,
	68, 0, 0, 0, 38, 282, 47, 0, 0, 48,
	39, 40, 0, 58, 0, 49, 0, 0, 0, 0,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 240, 241, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 0, 76, 75, 61, 41,
	67, 42, 68, 0, 0, 0, 38, 280, 47, 0,
	0, 48, 39, 40, 0, 58, 0, 49, 0, 0,
	0, 0, 0, 0, 60, 57, 0, 0, 69, 59,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 240, 241, 0, 61, 41, 67,
	42, 68, 0, 0, 64, 38, 65, 47, 76, 75,
	48, 39, 40, 0, 58, 0, 49, 0, 0, 0,
	0, 0, 0, 60, 57, 0, 0, 69, 59, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 240, 241, 0, 61, 41, 67, 42,
	68, 261, 0, 64, 38, 65, 47, 76, 75, 48,
	39, 40, 0, 58, 0, 49, 0, 0, 0, 0,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 0, 260, 0, 61, 41, 67, 42, 68,
	0, 0, 64, 38, 65, 47, 76, 75, 48, 39,
	40, 0, 58, 0, 49, 0, 0, 0, 0, 0,
	0, 60, 57, 0, 0, 69, 59, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 61, 124,
	67, 125, 68, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 0, 65, 0, 76, 75, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 0,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 251, 61, 124, 67, 125, 68,
	0, 387, 0, 0, 64, 0, 65, 0, 76, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 0, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 0, 0,
	0, 251, 61, 124, 67, 125, 68, 0, 372, 0,
	0, 64, 0, 65, 0, 76, 75, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 69, 0, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 251, 61,
	124, 67, 125, 267, 0, 338, 122, 0, 64, 0,
	65, 0, 76, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	0, 0, 77, 78, 0, 0, 266, 70, 71, 72,
	73, 74, 61, 124, 67, 125, 68, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 0, 123, 0, 76,
	75, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 69, 0, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 251, 61,
	124, 67, 125, 68, 0, 0, 0, 0, 64, 0,
	65, 0, 76, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	0, 0, 77, 78, 0, 0, 0, 70, 71, 72,
	73, 74, 61, 124, 67, 125, 68, 0, 0, 122,
	0, 0, 153, 0, 0, 64, 0, 65, 0, 76,
	75, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 69, 0, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 61, 256, 67, 125, 68,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 0,
	123, 0, 76, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 0, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 0, 0,
	0, 251, 61, 124, 67, 125, 111, 0, 0, 122,
	0, 64, 0, 65, 0, 76, 75, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 69, 0, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 61, 124, 67, 125, 68,
	0, 0, 0, 0, 0, 0, 0, 0, 114, 0,
	123, 0, 76, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 0, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 61, 413,
	67, 125, 68, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 0, 65, 0, 76, 75, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 0,
	0, 77, 78, 88, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 0, 65, 0, 76, 75,
	0, 0, 97, 98, 0, 0, 0, 0, 0, 86,
	87, 0, 0, 0, 89, 90, 91, 0, 0, 0,
	0, 0, 0, 0, 85, 94, 92, 93, 96,
}
var RubyPact = []int{

	-30, 1541, -1000, -1000, -1000, 5, -1000, -1000, -1000, 1160,
	-1000, -1000, -1000, 85, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 25, -19,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 308, 201,
	201, 688, 73, 2, 107, 58, 64, 2470, 2470, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 2920, 2470, 2920,
	2920, -1000, -1000, -1000, 2744, -1000, -8, 186, -1000, 16,
	2470, 2470, 2920, 2920, 2920, 19, 258, -1000, -1000, -1000,
	-1000, -1000, 2470, 2470, 2920, 321, 2920, 2920, 2920, 2920,
	2470, 2920, 2920, 2920, 2920, 2470, 2920, -1000, -1000, 2920,
	2470, 2920, 2920, 2470, 2470, 315, 216, 61, 36, -1000,
	-1000, 2744, 16, 18, 2744, 2920, 2920, 311, 158, 442,
	-1000, -7, -22, -22, 2877, 174, -26, -1000, -1000, 2744,
	2470, 2470, 2920, 2470, 2470, 2470, 2470, 2470, 2470, 2470,
	307, 106, 102, 2372, 112, 442, 188, 442, 112, 31,
	30, 2999, -1000, 2920, 2830, 226, 2744, 2421, -1000, -22,
	106, 106, 442, 442, 442, -1000, -1000, 249, -1000, -1000,
	106, 106, 442, 2654, 442, 442, 442, 442, 106, 442,
	442, 442, 442, 106, 318, 2697, 2697, 442, 106, 442,
	164, 106, 106, -1000, -1000, 303, 156, 71, -1000, 70,
	300, 293, -1000, 2323, 201, 2261, 282, 2999, -1000, -1000,
	-1000, 203, -44, 290, -1000, -1000, 236, -1000, -1000, 2787,
	2199, -1000, 283, -1000, 2137, 277, 106, 106, 54, 106,
	106, 106, 106, 106, 106, 106, -1000, 1492, -1000, -1000,
	-1000, -1000, 106, 141, 2920, -1000, 23, -1000, -1000, -1000,
	442, -1000, 154, 116, -1, 500, 1017, -1000, 268, 106,
	-1000, -1000, -1000, -1000, 18, 16, 2470, 2744, 2920, 442,
	442, 69, 1791, 61, 71, 171, 2920, -1000, -1000, 1729,
	-1000, -1000, -1000, 16, -1000, 2607, 68, -1000, 4, 442,
	-1000, -1000, 2088, 34, -1000, 2026, -1000, -1000, 47, -1000,
	322, 2470, -1000, 1430, 1977, -1000, -1000, 192, 442, 1368,
	196, 2920, 2560, -25, -1000, -27, -1000, 2470, 2920, -1000,
	-1000, 106, 165, 442, 2470, -1000, 138, -1000, -1000, -1000,
	-1000, 442, -1000, 137, 1915, -1000, 2513, 442, 262, 2470,
	241, -1000, -1000, 240, -1000, 2470, -1000, 2470, -1000, 106,
	2372, -1000, 291, -1000, 2372, 238, -1000, -1000, -1000, 106,
	-1000, -1000, 2470, 2470, 257, 223, -1000, -1000, 2920, 112,
	2999, -1000, 2920, -1000, 2697, -1000, 62, 1160, 106, 442,
	-1000, 106, -1000, -1000, 1667, -1000, -1000, 2963, -1000, 106,
	-1000, -1000, 106, 106, 2372, 2372, -1000, 2372, 225, 162,
	46, 2470, 2470, 2470, 2470, 1853, 112, 2372, 442, 353,
	-6, -1000, 136, 907, 2372, -1000, -1000, -1000, -1000, 106,
	106, 106, 106, -1000, 2372, -1, 2470, 2920, -1000, -1000,
	2372, 1306, 1234, 1104, -1, 212, 358, -1000, 287, 2470,
	-1000, -1000, 204, -1000, -1000, -1000, -1, -1000, -1000, 2470,
	-1000, 106, 1605, -1000, -1, 106, 1605, 1605, 1605,
}
var RubyPgo = []int{

	0, 0, 435, 431, 15, 32, 430, 429, 428, 918,
	426, 1, 31, 425, 424, 420, 419, 99, 416, 808,
	516, 415, 410, 409, 13, 407, 9, 406, 401, 400,
	12, 396, 390, 389, 384, 17, 383, 382, 381, 5,
	380, 379, 378, 375, 374, 372, 370, 366, 361, 360,
	355, 442, 354, 3, 16, 19, 7, 353, 6, 347,
	4, 343, 14, 342, 8, 341, 337, 11, 10, 332,
	296, 80,
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
	20, 20, 20, 20, 20, 20, 20, 20, 20, 54,
	54, 54, 54, 64, 64, 62, 62, 62, 62, 62,
	62, 62, 67, 67, 67, 67, 67, 66, 66, 66,
	21, 21, 21, 21, 21, 21, 58, 58, 68, 68,
	68, 26, 26, 26, 26, 25, 25, 28, 30, 30,
	69, 69, 15, 15, 15, 15, 15, 15, 15, 15,
	29, 29, 29, 29, 29, 29, 9, 9, 27, 27,
	19, 19, 40, 40, 41, 42, 43, 44, 45, 46,
	47, 48, 49, 50, 2, 6, 7, 7, 3, 3,
	59, 59, 59, 59, 65, 65, 65, 5, 5, 5,
	5, 55, 63, 63, 63, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 56, 56, 56, 56,
	52, 52, 52, 8, 16, 11, 11, 11, 61, 61,
	53, 53, 22, 23, 23, 12, 36, 60, 60, 60,
	60, 60, 60, 37, 37, 37, 37, 37, 37, 37,
	38, 38, 38, 38, 38, 39, 39, 39, 39, 34,
	33, 10, 32, 32, 31, 31, 4,
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
	3, 3, 3, 4, 4, 4, 6, 6, 6, 3,
	7, 1, 5, 1, 3, 0, 1, 1, 2, 4,
	4, 5, 1, 1, 4, 2, 5, 1, 3, 3,
	5, 6, 7, 8, 5, 6, 1, 3, 0, 1,
	3, 1, 2, 3, 2, 4, 6, 4, 1, 3,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 2, 2, 2,
	3, 3, 3, 3, 2, 2, 2, 2, 3, 3,
	3, 3, 3, 3, 1, 1, 3, 3, 5, 5,
	0, 1, 3, 7, 3, 7, 8, 3, 4, 4,
	3, 3, 0, 1, 3, 4, 5, 3, 3, 3,
	3, 3, 5, 6, 5, 3, 4, 3, 3, 2,
	0, 2, 2, 3, 4, 2, 3, 5, 0, 2,
	1, 2, 2, 2, 1, 5, 5, 0, 2, 2,
	2, 2, 2, 0, 1, 3, 3, 1, 3, 3,
	5, 6, 5, 6, 5, 4, 3, 3, 2, 4,
	4, 2, 5, 7, 4, 5, 3,
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
	-17, -1, -1, 6, -58, 55, -68, 9, -26, 6,
	47, 58, -58, -51, 40, -51, -62, -17, -5, 11,
	-5, -17, -4, -17, -35, -12, -17, -12, 6, 11,
	-51, -55, 56, -55, -51, -62, -1, -1, -17, -1,
	-1, -1, -1, -1, -1, -1, 6, -51, 51, 52,
	51, 52, -1, -61, 11, 51, -71, 62, 11, 62,
	-17, 51, -59, -65, -71, -17, 6, 8, -62, -1,
	52, 10, 6, 8, -67, -54, 42, 9, 53, -17,
	-17, 62, -51, 6, 11, -68, 42, 6, 6, -51,
	14, -30, 14, 10, 11, -71, 62, 62, -71, -17,
	-5, 14, -51, -63, 6, -51, 64, 10, 62, 14,
	-56, 17, 16, -51, -51, 14, -11, 25, -17, -60,
	-31, 37, -71, -71, 11, -71, 11, 4, 53, 10,
	-5, -1, -62, -17, 42, 14, -53, -11, -58, -26,
	10, -17, 14, -53, -51, -5, -71, -17, 58, 42,
	58, 14, 56, 11, 64, 42, 14, 17, 16, -1,
	-51, 14, -56, 14, -51, 8, 14, 51, 52, -1,
	-38, -37, 15, 18, 27, 28, 14, 16, 37, -64,
	-17, -24, 58, 64, -71, 64, -71, -17, -1, -17,
	10, -1, 14, -11, -51, 14, 14, 58, 6, -1,
	6, 6, -1, -1, -51, -51, 14, -51, 4, -1,
	-1, 15, 18, 15, 18, -51, -64, -51, -17, -17,
	6, 14, -53, 6, -51, 6, 51, 51, 52, -1,
	-1, -1, -1, 14, -51, -71, 4, 53, 14, 10,
	-51, -60, -60, -60, -71, -1, -17, 14, -39, 17,
	16, 14, -39, 14, -70, 11, -71, 11, 14, 17,
	16, -1, -60, 14, -71, -1, -60, -60, -60,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 48,
	49, 50, 51, 52, 53, 54, 55, 56, 57, 58,
	59, 60, 61, 62, 27, 28, 29, 30, 31, 32,
	33, 34, 35, 36, 37, 38, 39, 40, 0, 0,
	0, 20, -2, 22, 23, 24, 0, 0, 0, 15,
	41, 42, 43, 44, 45, 46, 47, 0, 0, 0,
	214, 19, 25, 26, 95, 13, 0, 67, 200, 0,
	0, 0, 0, 0, 0, 0, 0, 164, 165, 5,
	6, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 13, 13, 0,
	0, 0, 0, 0, 0, 0, 118, 118, 15, -2,
	15, -2, 70, 78, 95, 0, 0, 0, 91, 102,
	103, 31, 15, -2, 20, -2, 22, 23, 24, 95,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 15, 0, 208, 212, 93, 0, 13, 213, 0,
	0, 93, 97, 0, 13, 0, 95, 0, 241, 15,
	154, 155, 156, 157, 64, 148, 149, 0, 146, 147,
	187, 195, 63, 72, 79, 81, 82, 158, 159, 160,
	161, 162, 163, 189, 0, 0, 0, 246, 191, 80,
	0, 188, 190, 76, 15, 0, 116, 118, 119, 121,
	0, 0, 15, 0, 0, 0, 0, 96, 71, 13,
	105, 93, 0, 132, 133, 134, 140, 141, 152, 13,
	0, 15, 182, 15, 0, 0, 135, 142, 0, 136,
	143, 137, 144, 138, 145, 139, 153, 0, 15, 15,
	16, 17, 18, 0, 0, 217, 0, 166, 13, 167,
	98, 14, 13, 13, 171, 0, 20, -2, 0, 201,
	202, 203, 150, 151, 73, 74, 0, -2, 0, 239,
	240, 0, 0, 118, 0, 0, 0, 122, 124, 0,
	125, 15, 127, 65, 13, 0, 83, 85, 0, 108,
	109, 177, 0, 0, 183, 0, 180, 69, 84, 185,
	0, 0, 15, 0, 0, 204, 209, 15, 94, 0,
	0, 0, 0, 0, 13, 0, 13, 0, 0, 68,
	75, 77, 0, 215, 0, 110, 0, 210, 15, 120,
	117, 123, 114, 0, 0, 66, 0, 104, 0, 0,
	0, 178, 181, 0, 179, 0, 186, 0, 15, 15,
	199, 192, 0, 194, 205, 15, 216, 218, 219, 220,
	221, 222, 0, 0, 224, 227, 242, 15, 0, 15,
	99, 100, 0, 168, 0, 169, 0, -2, 172, 174,
	89, 88, 111, 211, 0, 115, 126, 0, 106, 86,
	92, 184, 87, 15, 197, 198, 193, 206, 0, 15,
	0, 0, 0, 0, 0, 0, 15, 13, 101, 0,
	0, 112, 0, 20, 196, 15, 217, -2, -2, 225,
	226, 228, 229, 243, 13, 244, 0, 0, 113, 90,
	207, 0, 0, 0, 245, 11, 13, 230, 0, 0,
	217, 232, 0, 234, 173, 12, 175, 13, 231, 0,
	217, 217, 223, 233, 176, 217, 223, 223, 223,
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
		//line parser.y:432
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 89:
		//line parser.y:442
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 90:
		//line parser.y:444
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 91:
		//line parser.y:446
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 92:
		//line parser.y:448
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 93:
		//line parser.y:451
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 94:
		//line parser.y:453
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 95:
		//line parser.y:455
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 96:
		//line parser.y:457
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 97:
		//line parser.y:459
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 98:
		//line parser.y:461
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{Func: ast.BareReference{Name: "to_proc"}, Target: RubyS[Rubypt-0].genericValue})
		}
	case 99:
		//line parser.y:463
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 100:
		//line parser.y:465
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 101:
		//line parser.y:467
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{
				Func:   ast.BareReference{Name: "to_proc"},
				Target: RubyS[Rubypt-0].genericValue,
			})
		}
	case 102:
		//line parser.y:476
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 103:
		//line parser.y:478
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 104:
		//line parser.y:480
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 105:
		//line parser.y:482
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 106:
		//line parser.y:484
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 107:
		//line parser.y:487
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 108:
		//line parser.y:489
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 109:
		//line parser.y:491
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 110:
		//line parser.y:495
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 111:
		//line parser.y:503
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 112:
		//line parser.y:512
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 113:
		//line parser.y:521
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 114:
		//line parser.y:531
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 115:
		//line parser.y:539
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 116:
		//line parser.y:550
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 117:
		//line parser.y:552
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 118:
		//line parser.y:554
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 119:
		//line parser.y:556
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 120:
		//line parser.y:558
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 121:
		//line parser.y:561
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 122:
		//line parser.y:563
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 123:
		//line parser.y:565
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 124:
		//line parser.y:567
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 125:
		//line parser.y:571
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 126:
		//line parser.y:578
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 127:
		//line parser.y:588
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 128:
		//line parser.y:597
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 129:
		//line parser.y:603
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 130:
		//line parser.y:611
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 131:
		//line parser.y:615
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 132:
		//line parser.y:620
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 133:
		//line parser.y:627
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 134:
		//line parser.y:631
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 135:
		//line parser.y:635
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 136:
		//line parser.y:642
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 137:
		//line parser.y:649
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 138:
		//line parser.y:656
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 139:
		//line parser.y:663
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 140:
		//line parser.y:671
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 141:
		//line parser.y:678
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 142:
		//line parser.y:682
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 143:
		//line parser.y:689
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 144:
		//line parser.y:696
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 145:
		//line parser.y:703
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 146:
		//line parser.y:711
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 147:
		//line parser.y:713
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 148:
		//line parser.y:716
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 149:
		//line parser.y:718
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 150:
		//line parser.y:721
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 151:
		//line parser.y:723
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 152:
		//line parser.y:726
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 153:
		//line parser.y:728
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 154:
		//line parser.y:730
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 155:
		//line parser.y:731
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 156:
		//line parser.y:732
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 157:
		//line parser.y:733
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 158:
		//line parser.y:736
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 159:
		//line parser.y:745
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 160:
		//line parser.y:754
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 161:
		//line parser.y:763
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 162:
		//line parser.y:772
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 163:
		//line parser.y:781
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 164:
		//line parser.y:789
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 165:
		//line parser.y:790
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 166:
		//line parser.y:792
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 167:
		//line parser.y:794
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
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
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 170:
		//line parser.y:813
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 171:
		//line parser.y:815
		{
		}
	case 172:
		//line parser.y:817
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 173:
		//line parser.y:824
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 174:
		//line parser.y:832
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 175:
		//line parser.y:839
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 176:
		//line parser.y:846
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 177:
		//line parser.y:854
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 178:
		//line parser.y:856
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 179:
		//line parser.y:863
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 180:
		//line parser.y:870
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 181:
		//line parser.y:873
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 182:
		//line parser.y:875
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 183:
		//line parser.y:877
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 184:
		//line parser.y:879
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 185:
		//line parser.y:882
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 186:
		//line parser.y:889
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 187:
		//line parser.y:897
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 188:
		//line parser.y:904
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 189:
		//line parser.y:911
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 190:
		//line parser.y:918
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 191:
		//line parser.y:925
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 192:
		//line parser.y:932
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 193:
		//line parser.y:939
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 194:
		//line parser.y:947
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 195:
		//line parser.y:954
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 196:
		//line parser.y:963
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 197:
		//line parser.y:970
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 198:
		//line parser.y:977
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 199:
		//line parser.y:984
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 200:
		//line parser.y:991
		{
		}
	case 201:
		//line parser.y:992
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 202:
		//line parser.y:993
		{
		}
	case 203:
		//line parser.y:996
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 204:
		//line parser.y:999
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 205:
		//line parser.y:1007
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 206:
		//line parser.y:1009
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 207:
		//line parser.y:1018
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
	case 208:
		//line parser.y:1033
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 209:
		//line parser.y:1035
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 210:
		//line parser.y:1038
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 211:
		//line parser.y:1040
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 212:
		//line parser.y:1043
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 213:
		//line parser.y:1052
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 214:
		//line parser.y:1060
		{
			RubyVAL.genericValue = ast.Return{}
		}
	case 215:
		//line parser.y:1063
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 216:
		//line parser.y:1072
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 217:
		//line parser.y:1075
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 218:
		//line parser.y:1077
		{
		}
	case 219:
		//line parser.y:1079
		{
		}
	case 220:
		//line parser.y:1081
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 221:
		//line parser.y:1083
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 222:
		//line parser.y:1085
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 223:
		//line parser.y:1087
		{
		}
	case 224:
		//line parser.y:1089
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 225:
		//line parser.y:1091
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 226:
		//line parser.y:1093
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 227:
		//line parser.y:1095
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 228:
		//line parser.y:1097
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 229:
		//line parser.y:1099
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 230:
		//line parser.y:1102
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 231:
		//line parser.y:1109
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 232:
		//line parser.y:1117
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 233:
		//line parser.y:1124
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 234:
		//line parser.y:1132
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 235:
		//line parser.y:1140
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 236:
		//line parser.y:1147
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 237:
		//line parser.y:1154
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 238:
		//line parser.y:1161
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 239:
		//line parser.y:1169
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 240:
		//line parser.y:1172
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 241:
		//line parser.y:1174
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 242:
		//line parser.y:1177
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 243:
		//line parser.y:1179
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 244:
		//line parser.y:1182
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 245:
		//line parser.y:1184
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 246:
		//line parser.y:1186
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
