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

//line parser.y:1143

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 42,
	54, 125,
	-2, 22,
	-1, 107,
	54, 125,
	-2, 123,
	-1, 109,
	10, 92,
	11, 92,
	-2, 194,
	-1, 121,
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
	-1, 123,
	54, 125,
	-2, 22,
	-1, 251,
	54, 126,
	-2, 124,
	-1, 261,
	10, 92,
	11, 92,
	-2, 194,
	-1, 400,
	27, 210,
	28, 210,
	-2, 15,
	-1, 401,
	27, 210,
	28, 210,
	-2, 15,
}

const RubyNprod = 240
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 2689

var RubyAct = []int{

	237, 314, 5, 421, 315, 146, 192, 118, 289, 190,
	194, 298, 28, 56, 111, 110, 163, 218, 164, 217,
	103, 151, 2, 3, 82, 25, 195, 24, 245, 245,
	120, 61, 122, 67, 123, 109, 147, 115, 120, 4,
	245, 361, 359, 331, 410, 245, 327, 328, 138, 139,
	243, 100, 106, 108, 205, 120, 150, 82, 239, 143,
	399, 69, 393, 141, 77, 78, 145, 196, 113, 70,
	71, 72, 73, 74, 17, 114, 165, 154, 197, 267,
	82, 155, 29, 168, 82, 200, 82, 112, 330, 121,
	175, 76, 75, 80, 79, 180, 156, 159, 160, 161,
	185, 244, 187, 188, 99, 134, 155, 245, 87, 242,
	81, 300, 135, 82, 239, 198, 400, 401, 204, 206,
	232, 233, 240, 387, 119, 245, 388, 210, 212, 221,
	222, 223, 224, 225, 226, 227, 228, 229, 216, 149,
	119, 209, 119, 119, 85, 86, 202, 119, 132, 88,
	89, 90, 98, 412, 253, 133, 411, 119, 127, 84,
	93, 91, 92, 95, 220, 130, 119, 296, 119, 119,
	119, 119, 131, 119, 119, 119, 119, 305, 119, 119,
	119, 119, 303, 119, 149, 258, 259, 149, 409, 252,
	428, 128, 119, 87, 82, 119, 119, 119, 129, 265,
	266, 241, 149, 137, 370, 119, 101, 248, 195, 102,
	119, 193, 367, 272, 83, 296, 214, 61, 122, 67,
	123, 68, 280, 296, 85, 86, 256, 149, 257, 85,
	86, 100, 98, 119, 136, 119, 342, 98, 127, 84,
	166, 294, 167, 295, 84, 366, 275, 69, 95, 196,
	77, 78, 296, 119, 152, 70, 71, 72, 73, 74,
	197, 310, 276, 245, 318, 265, 206, 309, 319, 107,
	372, 278, 321, 64, 316, 65, 317, 76, 75, 251,
	398, 9, 323, 436, 382, 433, 432, 61, 157, 396,
	158, 336, 353, 431, 354, 433, 432, 119, 311, 346,
	301, 339, 385, 302, 304, 386, 356, 364, 380, 376,
	335, 334, 308, 243, 195, 355, 105, 193, 104, 368,
	77, 78, 119, 117, 287, 243, 368, 333, 374, 335,
	334, 375, 324, 119, 373, 377, 149, 274, 275, 142,
	284, 144, 142, 269, 119, 119, 148, 76, 75, 268,
	383, 384, 264, 230, 213, 196, 162, 189, 170, 119,
	360, 390, 362, 191, 427, 169, 197, 171, 172, 173,
	174, 395, 176, 177, 178, 179, 358, 181, 182, 183,
	184, 66, 186, 119, 119, 116, 402, 403, 404, 405,
	119, 203, 247, 87, 207, 208, 211, 283, 238, 358,
	368, 246, 1, 153, 117, 55, 54, 119, 53, 203,
	418, 414, 415, 416, 52, 346, 346, 346, 51, 425,
	50, 35, 236, 434, 87, 34, 33, 32, 46, 85,
	86, 347, 249, 438, 203, 435, 346, 98, 119, 348,
	346, 346, 346, 119, 84, 439, 440, 19, 408, 20,
	441, 21, 117, 22, 306, 299, 14, 12, 11, 23,
	85, 86, 18, 10, 417, 16, 13, 37, 98, 61,
	122, 67, 123, 68, 15, 84, 429, 31, 30, 95,
	26, 236, 87, 236, 63, 36, 27, 437, 62, 430,
	85, 86, 0, 119, 0, 0, 279, 236, 98, 69,
	0, 236, 77, 78, 0, 84, 0, 70, 71, 72,
	73, 74, 87, 236, 0, 245, 0, 0, 85, 86,
	0, 297, 326, 0, 0, 64, 98, 65, 0, 76,
	75, 0, 117, 84, 0, 0, 0, 0, 61, 122,
	67, 123, 68, 203, 312, 236, 0, 0, 85, 86,
	0, 0, 236, 0, 0, 0, 98, 0, 325, 0,
	0, 262, 0, 84, 236, 0, 0, 236, 69, 0,
	0, 77, 78, 0, 236, 236, 70, 71, 72, 73,
	74, 0, 142, 357, 245, 0, 0, 0, 363, 365,
	0, 0, 0, 0, 64, 0, 65, 0, 76, 75,
	0, 0, 0, 0, 236, 0, 357, 0, 0, 0,
	0, 0, 0, 0, 61, 41, 67, 42, 68, 236,
	0, 0, 38, 236, 47, 43, 0, 48, 39, 40,
	0, 58, 0, 49, 0, 0, 0, 142, 0, 0,
	60, 57, 392, 0, 69, 59, 0, 77, 78, 0,
	0, 236, 70, 71, 72, 73, 74, 0, 0, 0,
	236, 236, 0, 236, 0, 0, 0, 124, 0, 0,
	64, 236, 65, 236, 76, 75, 0, 0, 0, 236,
	0, 0, 0, 124, 0, 124, 124, 0, 0, 236,
	124, 363, 419, 0, 0, 236, 124, 124, 124, 124,
	124, 0, 0, 61, 122, 67, 123, 109, 0, 124,
	120, 124, 124, 124, 124, 0, 124, 124, 124, 124,
	0, 124, 124, 124, 124, 0, 124, 0, 0, 44,
	0, 0, 0, 69, 0, 124, 77, 78, 124, 124,
	124, 70, 71, 72, 73, 74, 0, 0, 124, 0,
	0, 307, 0, 124, 0, 0, 0, 0, 0, 112,
	0, 121, 0, 76, 75, 0, 0, 0, 0, 0,
	0, 125, 0, 0, 0, 0, 124, 0, 124, 0,
	0, 0, 0, 0, 0, 0, 219, 125, 0, 125,
	125, 0, 0, 0, 125, 0, 124, 0, 0, 0,
	125, 125, 125, 125, 125, 0, 0, 0, 0, 0,
	0, 0, 0, 125, 0, 125, 125, 125, 125, 0,
	125, 125, 125, 125, 0, 125, 125, 125, 125, 0,
	125, 0, 0, 0, 0, 0, 140, 0, 0, 125,
	124, 0, 125, 125, 125, 0, 0, 0, 0, 0,
	0, 0, 125, 0, 0, 0, 0, 125, 0, 0,
	0, 0, 0, 0, 0, 124, 0, 0, 0, 87,
	0, 0, 0, 0, 45, 0, 124, 0, 0, 0,
	125, 0, 125, 94, 0, 0, 0, 124, 124, 0,
	83, 0, 0, 199, 0, 201, 0, 0, 96, 97,
	125, 0, 124, 0, 0, 85, 86, 215, 0, 0,
	88, 89, 90, 98, 0, 0, 126, 0, 0, 0,
	84, 93, 91, 92, 95, 231, 124, 124, 0, 0,
	0, 0, 126, 124, 126, 126, 0, 0, 0, 126,
	0, 0, 0, 0, 125, 126, 126, 126, 126, 126,
	124, 0, 0, 0, 0, 0, 0, 0, 126, 0,
	126, 126, 126, 126, 0, 126, 126, 126, 126, 125,
	126, 126, 126, 126, 0, 126, 0, 263, 0, 0,
	125, 124, 0, 0, 126, 270, 124, 126, 126, 126,
	0, 125, 125, 0, 0, 0, 0, 126, 0, 0,
	0, 0, 126, 282, 0, 285, 125, 0, 0, 61,
	122, 67, 123, 261, 0, 0, 120, 0, 0, 292,
	293, 0, 0, 0, 0, 126, 0, 126, 0, 0,
	125, 125, 0, 0, 0, 0, 124, 125, 0, 69,
	0, 0, 77, 78, 0, 126, 260, 70, 71, 72,
	73, 74, 0, 0, 125, 87, 0, 0, 0, 322,
	0, 0, 0, 0, 0, 64, 0, 121, 0, 76,
	75, 0, 0, 0, 0, 0, 0, 0, 337, 0,
	0, 0, 0, 341, 0, 125, 0, 0, 0, 126,
	125, 85, 86, 0, 0, 0, 88, 89, 90, 98,
	0, 0, 0, 369, 0, 0, 84, 93, 91, 92,
	95, 0, 0, 277, 126, 0, 0, 0, 0, 0,
	0, 0, 378, 379, 0, 126, 0, 0, 0, 381,
	0, 0, 0, 0, 0, 0, 126, 126, 0, 0,
	125, 389, 0, 391, 0, 0, 0, 0, 0, 0,
	0, 126, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 397, 0, 0, 0, 0, 0,
	231, 0, 0, 0, 0, 126, 126, 407, 0, 0,
	0, 0, 126, 0, 0, 413, 0, 292, 293, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 126,
	0, 0, 0, 0, 0, 0, 61, 41, 67, 42,
	68, 0, 0, 0, 38, 424, 349, 423, 422, 350,
	39, 40, 0, 58, 0, 49, 0, 0, 351, 352,
	126, 0, 60, 57, 0, 126, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 344, 345, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 0, 76, 75, 0, 61,
	41, 67, 42, 68, 0, 0, 0, 38, 420, 349,
	423, 422, 350, 39, 40, 126, 58, 0, 49, 0,
	0, 351, 352, 0, 0, 60, 57, 0, 0, 69,
	59, 0, 77, 78, 0, 0, 0, 70, 71, 72,
	73, 74, 0, 0, 0, 344, 345, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 0, 65, 0, 76,
	75, 61, 41, 67, 42, 68, 0, 0, 0, 38,
	426, 349, 0, 0, 350, 39, 40, 0, 58, 0,
	49, 0, 0, 351, 352, 0, 0, 60, 57, 0,
	0, 69, 59, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 0, 0, 0, 344, 345, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 0, 65,
	0, 76, 75, 61, 41, 67, 42, 68, 0, 0,
	0, 38, 343, 349, 0, 0, 350, 39, 40, 0,
	58, 0, 49, 0, 0, 351, 352, 0, 0, 60,
	57, 0, 0, 69, 59, 0, 77, 78, 0, 0,
	0, 70, 71, 72, 73, 74, 0, 0, 0, 344,
	345, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	0, 65, 0, 76, 75, 61, 41, 67, 42, 68,
	0, 0, 0, 38, 338, 47, 291, 290, 48, 39,
	40, 0, 58, 0, 49, 0, 0, 0, 0, 0,
	0, 60, 57, 0, 0, 69, 59, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 0, 0,
	0, 234, 235, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 0, 65, 0, 76, 75, 61, 41, 67,
	42, 68, 0, 0, 0, 38, 288, 47, 291, 290,
	48, 39, 40, 0, 58, 0, 49, 0, 0, 0,
	0, 0, 0, 60, 57, 0, 0, 69, 59, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 234, 235, 0, 61, 41, 67, 42,
	68, 0, 0, 64, 38, 65, 47, 76, 75, 48,
	39, 40, 0, 58, 0, 49, 0, 0, 0, 0,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 6, 7, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 0, 76, 75, 0, 8,
	61, 41, 67, 42, 68, 0, 0, 0, 38, 0,
	349, 0, 0, 350, 39, 40, 0, 58, 0, 49,
	0, 0, 351, 352, 0, 0, 60, 57, 0, 0,
	69, 59, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 0, 0, 0, 344, 345, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 65, 0,
	76, 75, 61, 41, 67, 42, 68, 0, 0, 0,
	38, 394, 47, 0, 0, 48, 39, 40, 0, 58,
	0, 49, 296, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 234, 235,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 0,
	65, 0, 76, 75, 61, 41, 67, 42, 68, 0,
	0, 0, 38, 320, 47, 0, 0, 48, 39, 40,
	0, 58, 0, 49, 296, 0, 0, 0, 0, 0,
	60, 57, 0, 0, 69, 59, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 0, 0, 0,
	234, 235, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 0, 65, 0, 76, 75, 61, 41, 67, 42,
	68, 0, 0, 0, 38, 313, 47, 0, 0, 48,
	39, 40, 0, 58, 0, 49, 296, 0, 0, 0,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 234, 235, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 0, 76, 75, 61, 41,
	67, 42, 68, 0, 0, 0, 38, 406, 47, 0,
	0, 48, 39, 40, 0, 58, 0, 49, 0, 0,
	0, 0, 0, 0, 60, 57, 0, 0, 69, 59,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 234, 235, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 0, 65, 0, 76, 75,
	61, 41, 67, 42, 68, 0, 0, 0, 38, 371,
	47, 0, 0, 48, 39, 40, 0, 58, 0, 49,
	0, 0, 0, 0, 0, 0, 60, 57, 0, 0,
	69, 59, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 0, 0, 0, 234, 235, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 65, 0,
	76, 75, 61, 41, 67, 42, 68, 0, 0, 0,
	38, 340, 47, 0, 0, 48, 39, 40, 0, 58,
	0, 49, 0, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 234, 235,
	0, 61, 41, 67, 42, 68, 0, 0, 64, 38,
	65, 47, 76, 75, 48, 39, 40, 0, 58, 0,
	49, 0, 0, 0, 0, 0, 0, 60, 57, 0,
	0, 69, 59, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 0, 0, 0, 234, 235, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 0, 65,
	332, 76, 75, 61, 41, 67, 42, 68, 0, 0,
	0, 38, 329, 47, 0, 0, 48, 39, 40, 0,
	58, 0, 49, 0, 0, 0, 0, 0, 0, 60,
	57, 0, 0, 69, 59, 0, 77, 78, 0, 0,
	0, 70, 71, 72, 73, 74, 0, 0, 0, 234,
	235, 0, 61, 41, 67, 42, 68, 0, 0, 64,
	38, 65, 47, 76, 75, 48, 39, 40, 0, 58,
	0, 49, 0, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 234, 235,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 0,
	65, 286, 76, 75, 61, 41, 67, 42, 68, 0,
	0, 0, 38, 281, 47, 0, 0, 48, 39, 40,
	0, 58, 0, 49, 0, 0, 0, 0, 0, 0,
	60, 57, 0, 0, 69, 59, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 0, 0, 0,
	234, 235, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 0, 65, 0, 76, 75, 61, 41, 67, 42,
	68, 0, 0, 0, 38, 273, 47, 0, 0, 48,
	39, 40, 0, 58, 0, 49, 0, 0, 0, 0,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 234, 235, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 0, 76, 75, 61, 41,
	67, 42, 68, 0, 0, 0, 38, 271, 47, 0,
	0, 48, 39, 40, 0, 58, 0, 49, 0, 0,
	0, 0, 0, 0, 60, 57, 0, 0, 69, 59,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 234, 235, 0, 61, 41, 67,
	42, 68, 0, 0, 64, 38, 65, 47, 76, 75,
	48, 39, 40, 0, 58, 0, 49, 0, 0, 0,
	0, 0, 0, 60, 57, 0, 0, 69, 59, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 234, 235, 0, 61, 41, 67, 42,
	68, 255, 0, 64, 38, 65, 47, 76, 75, 48,
	39, 40, 0, 58, 0, 49, 0, 0, 0, 0,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 0, 254, 61, 122, 67, 123, 68, 0,
	0, 120, 64, 0, 65, 0, 76, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 0, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 61, 250, 67,
	123, 68, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 0, 121, 0, 76, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 0, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 245, 61, 122, 67, 123, 109, 0,
	0, 120, 0, 64, 0, 65, 0, 76, 75, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 0, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 61, 122, 67,
	123, 68, 0, 0, 0, 0, 0, 0, 0, 0,
	112, 0, 121, 0, 76, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 0, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 0, 65, 0, 76, 75,
}
var RubyPact = []int{

	-29, 1561, -1000, -1000, -1000, 42, -1000, -1000, -1000, 865,
	-1000, -1000, -1000, 86, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 191,
	-35, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 312, 261,
	261, 26, 149, 123, 106, 63, 192, 609, 609, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 2622, 609, 2622,
	2622, -1000, -1000, -1000, 2622, -1000, -33, 245, -1000, 18,
	282, 282, 282, 282, 2622, 10, 234, -1000, -1000, -1000,
	-1000, -1000, 609, 2622, 352, 2622, 2622, 2622, 2622, 609,
	2622, 2622, 2622, 2622, 609, 2622, 2622, 2622, 2622, 609,
	2622, 609, 609, 351, 308, 202, 45, -1000, -1000, 2622,
	18, 43, 2622, 2622, 2622, 348, 205, 389, -1000, 11,
	-37, -37, 2579, 229, -1000, -1000, -1000, 2622, 609, 609,
	609, 609, 609, 609, 609, 609, 609, 347, 98, 69,
	2392, 103, 389, 71, 389, 103, 47, 39, 104, -1000,
	2532, 271, 2622, 2441, -1000, -37, -1000, -1000, -1000, -1000,
	-1000, -1000, 389, -1000, -1000, 220, -1000, -1000, 98, 389,
	1004, 389, 389, 389, 389, 98, 389, 389, 389, 389,
	98, 508, 389, 389, 389, 98, 389, 98, 98, -1000,
	-1000, 346, 188, 20, -1000, 37, 343, 337, -1000, 2343,
	261, 2281, 327, 104, -1000, -1000, -1000, 1051, 189, -1000,
	-1000, 420, -1000, -1000, 2489, 2219, -1000, 334, -1000, 2157,
	314, 98, 98, 98, 98, 98, 98, 98, 98, 98,
	-1000, 1512, -1000, -1000, -1000, -1000, 865, 98, 227, 2622,
	-1000, 74, -1000, -1000, -1000, -1000, 171, 166, -6, 450,
	698, -1000, 302, 98, -1000, -1000, -1000, -1000, 43, 18,
	609, 2622, 2622, 1811, 202, 20, 254, 609, -1000, -1000,
	1749, -1000, -1000, -1000, 18, -1000, 464, 4, -11, 389,
	-1000, -1000, 2108, 32, -1000, 2046, -1000, -1000, -1000, 313,
	609, -1000, 1450, 1997, -1000, -1000, 228, 389, 1388, 278,
	2622, 533, -22, -1000, -23, -1000, 609, 2622, -1000, -1000,
	98, 235, 389, -1000, 198, -1000, -1000, -1000, -1000, 98,
	-1000, 190, 1935, -1000, 212, 389, 328, 609, 325, -1000,
	-1000, 303, -1000, -1000, 609, -1000, 98, 2392, -1000, 294,
	-1000, 2392, 280, -1000, -1000, -1000, 98, -1000, -1000, 609,
	609, 287, 108, -1000, -1000, 2622, 103, 104, -1000, -1000,
	533, -1000, 56, 865, 98, 389, -1000, -1000, -1000, 1687,
	-1000, -1000, 283, -1000, 98, -1000, -1000, 98, 2392, 2392,
	-1000, 2392, 274, 9, 65, 609, 609, 609, 609, 1873,
	103, 2392, 184, -9, -1000, 142, 143, 2392, -1000, -1000,
	-1000, -1000, 98, 98, 98, 98, -1000, 2392, -6, 609,
	2622, -1000, -1000, 2392, 1264, 1201, 1326, -6, 179, 478,
	-1000, 279, 609, -1000, -1000, 269, -1000, -1000, -1000, -6,
	-1000, -1000, 609, -1000, 98, 1625, -1000, -6, 98, 1625,
	1625, 1625,
}
var RubyPgo = []int{

	0, 0, 488, 486, 485, 7, 484, 480, 478, 874,
	477, 4, 13, 474, 467, 466, 465, 281, 25, 729,
	82, 463, 462, 459, 74, 458, 10, 625, 457, 456,
	12, 455, 453, 451, 449, 27, 447, 439, 431, 3,
	428, 427, 426, 425, 421, 420, 418, 414, 408, 406,
	405, 786, 403, 1, 15, 17, 8, 402, 9, 401,
	11, 398, 36, 397, 5, 392, 385, 14, 6, 381,
	364, 56,
}
var RubyR1 = []int{

	0, 57, 57, 57, 57, 57, 57, 57, 57, 57,
	57, 70, 70, 71, 71, 51, 51, 51, 51, 51,
	18, 18, 18, 18, 18, 18, 18, 18, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 24, 24, 24, 24, 24, 24, 24, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 35, 14, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 54, 54, 54, 54,
	64, 64, 62, 62, 62, 62, 62, 67, 67, 67,
	67, 67, 66, 66, 66, 21, 21, 21, 21, 21,
	21, 58, 58, 26, 26, 26, 26, 68, 68, 68,
	25, 25, 28, 30, 30, 69, 69, 15, 15, 15,
	15, 15, 15, 15, 15, 29, 29, 29, 29, 29,
	29, 9, 9, 27, 27, 19, 19, 40, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 2,
	6, 7, 7, 3, 3, 59, 59, 59, 59, 65,
	65, 65, 5, 5, 5, 5, 55, 63, 63, 63,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	56, 56, 56, 56, 52, 52, 52, 8, 16, 11,
	11, 11, 61, 61, 53, 53, 22, 23, 12, 36,
	60, 60, 60, 60, 60, 60, 37, 37, 37, 37,
	37, 37, 37, 38, 38, 38, 38, 38, 39, 39,
	39, 39, 34, 33, 10, 32, 32, 31, 31, 4,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 1, 0, 2, 0, 2, 2, 2, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 3, 2, 4, 5, 1, 4,
	4, 2, 3, 3, 4, 4, 5, 3, 5, 2,
	3, 3, 3, 3, 4, 6, 3, 7, 1, 5,
	1, 3, 0, 1, 1, 4, 4, 1, 1, 4,
	2, 5, 1, 3, 3, 5, 6, 7, 8, 5,
	6, 1, 3, 1, 2, 3, 2, 0, 1, 3,
	4, 6, 4, 1, 3, 1, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 2, 2, 3, 3, 3, 3, 2,
	2, 2, 2, 3, 3, 3, 3, 3, 3, 1,
	1, 3, 3, 5, 5, 0, 1, 3, 7, 3,
	7, 8, 3, 4, 4, 3, 3, 0, 1, 3,
	4, 5, 3, 3, 3, 3, 3, 5, 6, 5,
	4, 3, 3, 2, 0, 2, 2, 3, 4, 2,
	3, 5, 0, 2, 1, 2, 2, 2, 5, 5,
	0, 2, 2, 2, 2, 2, 0, 1, 3, 3,
	1, 3, 3, 5, 6, 5, 6, 5, 4, 3,
	3, 2, 3, 3, 2, 5, 7, 4, 5, 3,
}
var RubyChk = []int{

	-1000, -57, 51, 52, 68, -1, 51, 52, 68, -17,
	-21, -25, -28, -15, -29, -13, -16, -24, -22, -36,
	-34, -33, -32, -23, -35, -18, -7, -3, -30, -20,
	-8, -10, -41, -42, -43, -44, -4, -14, 13, 19,
	20, 6, 8, -27, -19, -9, -40, 15, 18, 24,
	-45, -46, -47, -48, -49, -50, -12, 32, 22, 36,
	31, 5, -2, -6, 61, 63, -69, 7, 9, 35,
	43, 44, 45, 46, 47, 66, 65, 38, 39, 52,
	51, 68, 15, 25, 55, 40, 41, 4, 45, 46,
	47, 57, 58, 56, 18, 59, 33, 34, 48, 18,
	40, 15, 18, 55, 6, 4, -30, 8, -30, 9,
	-54, -67, 61, 42, 49, 11, -66, -17, -5, -20,
	12, 63, 6, 8, -27, -19, -9, 9, 42, 49,
	42, 49, 42, 49, 42, 49, 42, 11, -1, -1,
	-51, -64, -17, -1, -17, -64, -64, -62, -17, -24,
	-71, 54, 9, -52, -5, 63, -18, 6, 8, -18,
	-18, -18, -17, 6, 8, 66, 6, 8, -1, -17,
	6, -17, -17, -17, -17, -1, -17, -17, -17, -17,
	-1, -17, -17, -17, -17, -1, -17, -1, -1, 6,
	-58, 55, -68, 9, -26, 6, 47, 58, -58, -51,
	40, -51, -62, -17, -5, 11, -5, -17, -17, -35,
	-12, -17, -12, 6, 11, -51, -55, 56, -55, -51,
	-62, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	6, -51, 51, 52, 51, 52, -17, -1, -61, 11,
	51, -71, 62, 11, 62, 51, -59, -65, -71, -17,
	6, 8, -62, -1, 52, 10, 6, 8, -67, -54,
	42, 9, 53, -51, 6, 11, -68, 42, 6, 6,
	-51, 14, -30, 14, 10, 11, -71, 62, -71, -17,
	-5, 14, -51, -63, 6, -51, 64, 10, 14, -56,
	17, 16, -51, -51, 14, -11, 25, -17, -60, -31,
	37, -71, -71, 11, -71, 11, 4, 53, 10, -5,
	-1, -62, -17, 14, -53, -11, -58, -26, 10, -1,
	14, -53, -51, -5, -71, -17, 58, 42, 58, 14,
	56, 11, 64, 14, 17, 16, -1, -51, 14, -56,
	14, -51, 8, 14, 51, 52, -1, -38, -37, 15,
	18, 27, 28, 14, 16, 37, -64, -17, -24, 64,
	-71, 64, -71, -17, -1, -17, 10, 14, -11, -51,
	14, 14, 58, 6, -1, 6, 6, -1, -51, -51,
	14, -51, 4, -1, -1, 15, 18, 15, 18, -51,
	-64, -51, -17, 6, 14, -53, 6, -51, 6, 51,
	51, 52, -1, -1, -1, -1, 14, -51, -71, 4,
	53, 14, 10, -51, -60, -60, -60, -71, -1, -17,
	14, -39, 17, 16, 14, -39, 14, -70, 11, -71,
	11, 14, 17, 16, -1, -60, 14, -71, -1, -60,
	-60, -60,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 48,
	49, 50, 51, 52, 53, 54, 55, 56, 57, 58,
	59, 60, 61, 62, 63, 28, 29, 30, 31, 32,
	33, 34, 35, 36, 37, 38, 39, 40, 0, 0,
	0, 21, -2, 23, 24, 25, 0, 0, 0, 15,
	41, 42, 43, 44, 45, 46, 47, 0, 0, 0,
	0, 20, 26, 27, 92, 13, 0, 68, 194, 0,
	0, 0, 0, 0, 0, 0, 0, 159, 160, 5,
	6, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 117, 117, 15, -2, 15, -2,
	71, 79, 92, 0, 0, 0, 88, 97, 98, 32,
	15, -2, 21, -2, 23, 24, 25, 92, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 15, 0,
	202, 206, 90, 0, 13, 207, 0, 0, 90, 94,
	13, 0, 92, 0, 234, 15, 149, 21, 22, 150,
	151, 152, 65, 143, 144, 0, 141, 142, 182, 64,
	73, 80, 82, 83, 153, 154, 155, 156, 157, 158,
	184, 0, 232, 233, 239, 186, 81, 183, 185, 77,
	15, 0, 111, 117, 118, 113, 0, 0, 15, 0,
	0, 0, 0, 93, 72, 13, 100, 90, 127, 128,
	129, 135, 136, 147, 13, 0, 15, 177, 15, 0,
	0, 130, 137, 131, 138, 132, 139, 133, 140, 134,
	148, 0, 15, 15, 16, 17, 18, 19, 0, 0,
	210, 0, 161, 13, 162, 14, 13, 13, 166, 0,
	21, -2, 0, 195, 196, 197, 145, 146, 74, 75,
	0, -2, 0, 0, 117, 0, 0, 0, 114, 116,
	0, 120, 15, 122, 66, 13, 0, 84, 0, 103,
	104, 172, 0, 0, 178, 0, 175, 70, 180, 0,
	0, 15, 0, 0, 198, 203, 15, 91, 0, 0,
	0, 0, 0, 13, 0, 13, 0, 0, 69, 76,
	78, 0, 208, 105, 0, 204, 15, 119, 112, 115,
	109, 0, 0, 67, 0, 99, 0, 0, 0, 173,
	176, 0, 174, 181, 0, 15, 15, 193, 187, 0,
	189, 199, 15, 209, 211, 212, 213, 214, 215, 0,
	0, 217, 220, 235, 15, 0, 15, 95, 96, 163,
	0, 164, 0, 48, 167, 169, 86, 106, 205, 0,
	110, 121, 0, 101, 85, 89, 179, 15, 191, 192,
	188, 200, 0, 15, 0, 0, 0, 0, 0, 0,
	15, 13, 0, 0, 107, 0, 0, 190, 15, 210,
	-2, -2, 218, 219, 221, 222, 236, 13, 237, 0,
	0, 108, 87, 201, 0, 0, 0, 238, 11, 13,
	223, 0, 0, 210, 225, 0, 227, 168, 12, 170,
	13, 224, 0, 210, 210, 216, 226, 171, 210, 216,
	216, 216,
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
		//line parser.y:236
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 61:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 62:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 63:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 64:
		//line parser.y:248
		{
			RubyVAL.genericValue = ast.RescueModifier{Statement: RubyS[Rubypt-2].genericValue, Rescue: RubyS[Rubypt-0].genericValue}
		}
	case 65:
		//line parser.y:251
		{
			RubyVAL.genericValue = ast.StarSplat{Value: RubyS[Rubypt-0].genericValue}
		}
	case 66:
		//line parser.y:254
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 67:
		//line parser.y:261
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-2].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 68:
		//line parser.y:268
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 69:
		//line parser.y:272
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 70:
		//line parser.y:279
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 71:
		//line parser.y:286
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 72:
		//line parser.y:293
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 73:
		//line parser.y:300
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 74:
		//line parser.y:307
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 75:
		//line parser.y:315
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 76:
		//line parser.y:323
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 77:
		//line parser.y:331
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 78:
		//line parser.y:339
		{
			methodName := RubyS[Rubypt-2].genericValue.(ast.BareReference).Name + "="
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: methodName},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 79:
		//line parser.y:350
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 80:
		//line parser.y:357
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 81:
		//line parser.y:365
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 82:
		//line parser.y:373
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 83:
		//line parser.y:381
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 84:
		//line parser.y:391
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 85:
		//line parser.y:401
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 86:
		//line parser.y:410
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 87:
		//line parser.y:412
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 88:
		//line parser.y:414
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 89:
		//line parser.y:416
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 90:
		//line parser.y:419
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 91:
		//line parser.y:421
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 92:
		//line parser.y:423
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 93:
		//line parser.y:425
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 94:
		//line parser.y:427
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 95:
		//line parser.y:429
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 96:
		//line parser.y:431
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 97:
		//line parser.y:435
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 98:
		//line parser.y:437
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 99:
		//line parser.y:439
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 100:
		//line parser.y:441
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 101:
		//line parser.y:443
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 102:
		//line parser.y:446
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 103:
		//line parser.y:448
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 104:
		//line parser.y:450
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 105:
		//line parser.y:456
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 106:
		//line parser.y:464
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 107:
		//line parser.y:473
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 108:
		//line parser.y:482
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 109:
		//line parser.y:492
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 110:
		//line parser.y:500
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 111:
		//line parser.y:511
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 112:
		//line parser.y:513
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 113:
		//line parser.y:516
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 114:
		//line parser.y:518
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 115:
		//line parser.y:520
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 116:
		//line parser.y:522
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 117:
		//line parser.y:524
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 118:
		//line parser.y:526
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 119:
		//line parser.y:530
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 120:
		//line parser.y:535
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 121:
		//line parser.y:542
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 122:
		//line parser.y:552
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 123:
		//line parser.y:561
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 124:
		//line parser.y:567
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 125:
		//line parser.y:575
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 126:
		//line parser.y:579
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 127:
		//line parser.y:584
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 128:
		//line parser.y:591
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 129:
		//line parser.y:595
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 130:
		//line parser.y:599
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 131:
		//line parser.y:606
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 132:
		//line parser.y:613
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 133:
		//line parser.y:620
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
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
		//line parser.y:635
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 136:
		//line parser.y:642
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 137:
		//line parser.y:646
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 138:
		//line parser.y:653
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 139:
		//line parser.y:660
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 140:
		//line parser.y:667
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 141:
		//line parser.y:675
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 142:
		//line parser.y:677
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 143:
		//line parser.y:680
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 144:
		//line parser.y:682
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 145:
		//line parser.y:685
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 146:
		//line parser.y:687
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 147:
		//line parser.y:690
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 148:
		//line parser.y:692
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 149:
		//line parser.y:694
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 150:
		//line parser.y:695
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 151:
		//line parser.y:696
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 152:
		//line parser.y:697
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 153:
		//line parser.y:700
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 154:
		//line parser.y:709
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 155:
		//line parser.y:718
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 156:
		//line parser.y:727
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 157:
		//line parser.y:736
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 158:
		//line parser.y:745
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 159:
		//line parser.y:753
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 160:
		//line parser.y:754
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 161:
		//line parser.y:756
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 162:
		//line parser.y:758
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 163:
		//line parser.y:761
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 164:
		//line parser.y:769
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 165:
		//line parser.y:777
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 166:
		//line parser.y:779
		{
		}
	case 167:
		//line parser.y:781
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 168:
		//line parser.y:788
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 169:
		//line parser.y:796
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 170:
		//line parser.y:803
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 171:
		//line parser.y:810
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 172:
		//line parser.y:818
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 173:
		//line parser.y:820
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 174:
		//line parser.y:827
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 175:
		//line parser.y:834
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 176:
		//line parser.y:837
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 177:
		//line parser.y:839
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 178:
		//line parser.y:841
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 179:
		//line parser.y:843
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 180:
		//line parser.y:846
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 181:
		//line parser.y:853
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 182:
		//line parser.y:861
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 183:
		//line parser.y:868
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 184:
		//line parser.y:875
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 185:
		//line parser.y:882
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 186:
		//line parser.y:889
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 187:
		//line parser.y:896
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 188:
		//line parser.y:903
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 189:
		//line parser.y:911
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 190:
		//line parser.y:920
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 191:
		//line parser.y:927
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 192:
		//line parser.y:934
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 193:
		//line parser.y:941
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 194:
		//line parser.y:948
		{
		}
	case 195:
		//line parser.y:949
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 196:
		//line parser.y:950
		{
		}
	case 197:
		//line parser.y:953
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 198:
		//line parser.y:956
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 199:
		//line parser.y:964
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 200:
		//line parser.y:966
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 201:
		//line parser.y:975
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
	case 202:
		//line parser.y:990
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 203:
		//line parser.y:992
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 204:
		//line parser.y:995
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 205:
		//line parser.y:997
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 206:
		//line parser.y:1000
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 207:
		//line parser.y:1009
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 208:
		//line parser.y:1018
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 209:
		//line parser.y:1027
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 210:
		//line parser.y:1030
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 211:
		//line parser.y:1032
		{
		}
	case 212:
		//line parser.y:1034
		{
		}
	case 213:
		//line parser.y:1036
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 214:
		//line parser.y:1038
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 215:
		//line parser.y:1040
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 216:
		//line parser.y:1042
		{
		}
	case 217:
		//line parser.y:1044
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 218:
		//line parser.y:1046
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 219:
		//line parser.y:1048
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 220:
		//line parser.y:1050
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 221:
		//line parser.y:1052
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 222:
		//line parser.y:1054
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 223:
		//line parser.y:1057
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 224:
		//line parser.y:1064
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 225:
		//line parser.y:1072
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 226:
		//line parser.y:1079
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 227:
		//line parser.y:1087
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 228:
		//line parser.y:1095
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 229:
		//line parser.y:1102
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 230:
		//line parser.y:1109
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 231:
		//line parser.y:1116
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 232:
		//line parser.y:1124
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 233:
		//line parser.y:1127
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 234:
		//line parser.y:1129
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 235:
		//line parser.y:1132
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 236:
		//line parser.y:1134
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 237:
		//line parser.y:1137
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 238:
		//line parser.y:1139
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 239:
		//line parser.y:1141
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
