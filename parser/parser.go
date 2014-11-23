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

//line parser.y:1220

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 42,
	54, 134,
	-2, 21,
	-1, 109,
	54, 134,
	-2, 132,
	-1, 111,
	10, 99,
	11, 99,
	-2, 204,
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
	54, 134,
	-2, 21,
	-1, 259,
	54, 135,
	-2, 133,
	-1, 269,
	10, 99,
	11, 99,
	-2, 204,
	-1, 383,
	61, 83,
	-2, 48,
	-1, 423,
	27, 222,
	28, 222,
	-2, 15,
	-1, 424,
	27, 222,
	28, 222,
	-2, 15,
}

const RubyNprod = 252
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 3162

var RubyAct = []int{

	244, 333, 5, 444, 315, 332, 199, 306, 195, 149,
	120, 27, 197, 17, 113, 112, 2, 3, 225, 292,
	103, 150, 253, 104, 88, 253, 82, 118, 56, 83,
	22, 132, 23, 4, 165, 381, 166, 122, 379, 133,
	253, 224, 349, 210, 122, 101, 134, 346, 141, 142,
	105, 108, 110, 97, 98, 101, 88, 155, 132, 146,
	86, 87, 80, 79, 274, 433, 102, 144, 274, 250,
	148, 160, 161, 253, 205, 85, 102, 200, 152, 81,
	158, 140, 304, 170, 171, 97, 98, 348, 159, 351,
	345, 178, 86, 87, 167, 159, 183, 89, 90, 91,
	99, 188, 246, 317, 192, 193, 9, 85, 94, 92,
	93, 96, 139, 100, 291, 293, 203, 253, 201, 273,
	251, 330, 281, 209, 211, 152, 88, 154, 152, 202,
	190, 228, 229, 207, 231, 232, 233, 234, 235, 236,
	237, 223, 214, 152, 217, 219, 216, 213, 119, 82,
	129, 227, 83, 249, 82, 97, 98, 83, 261, 386,
	289, 137, 86, 87, 145, 135, 147, 145, 138, 451,
	152, 151, 136, 82, 416, 246, 83, 85, 260, 162,
	163, 164, 322, 200, 275, 422, 198, 129, 266, 267,
	247, 172, 320, 174, 175, 176, 177, 279, 179, 180,
	181, 182, 221, 184, 432, 276, 187, 82, 189, 191,
	83, 280, 61, 124, 67, 125, 68, 286, 208, 253,
	130, 212, 215, 218, 201, 185, 186, 131, 336, 279,
	156, 119, 296, 97, 98, 202, 208, 434, 391, 230,
	86, 87, 69, 423, 424, 77, 78, 312, 313, 313,
	70, 71, 72, 73, 74, 85, 388, 361, 253, 88,
	252, 257, 264, 208, 265, 393, 109, 313, 64, 327,
	65, 259, 76, 75, 372, 248, 373, 211, 326, 421,
	119, 200, 256, 152, 198, 296, 335, 334, 97, 98,
	339, 328, 271, 272, 311, 86, 87, 374, 397, 341,
	89, 90, 91, 396, 82, 313, 88, 83, 355, 394,
	85, 94, 92, 93, 96, 300, 365, 358, 459, 283,
	456, 455, 201, 82, 384, 282, 83, 375, 295, 404,
	196, 387, 377, 202, 389, 97, 98, 454, 290, 456,
	455, 389, 86, 87, 409, 407, 395, 410, 408, 294,
	88, 278, 398, 314, 399, 238, 377, 85, 450, 240,
	241, 96, 325, 250, 95, 119, 303, 250, 220, 405,
	406, 84, 402, 194, 354, 353, 208, 329, 318, 97,
	98, 295, 319, 321, 412, 66, 86, 87, 337, 288,
	289, 89, 90, 91, 99, 173, 418, 343, 168, 255,
	169, 85, 94, 92, 93, 96, 299, 88, 425, 426,
	427, 428, 245, 352, 453, 354, 353, 342, 43, 107,
	389, 106, 254, 1, 145, 376, 157, 437, 438, 439,
	383, 385, 55, 441, 54, 53, 97, 98, 365, 365,
	365, 52, 448, 86, 87, 51, 457, 50, 380, 376,
	382, 458, 34, 33, 32, 31, 461, 46, 85, 365,
	126, 462, 463, 365, 365, 365, 464, 366, 367, 19,
	36, 37, 20, 316, 14, 12, 126, 11, 126, 126,
	21, 145, 18, 126, 10, 414, 24, 415, 16, 13,
	35, 126, 126, 126, 15, 88, 30, 29, 25, 63,
	414, 26, 62, 126, 0, 126, 126, 126, 126, 0,
	126, 126, 126, 126, 0, 126, 84, 0, 126, 0,
	126, 126, 0, 88, 97, 98, 119, 0, 0, 0,
	126, 86, 87, 126, 126, 126, 0, 0, 0, 383,
	442, 431, 0, 126, 0, 0, 85, 0, 126, 0,
	96, 126, 97, 98, 0, 0, 0, 0, 440, 86,
	87, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	452, 0, 126, 126, 85, 126, 61, 41, 67, 42,
	68, 460, 0, 0, 38, 447, 368, 446, 445, 369,
	39, 40, 126, 58, 0, 49, 0, 0, 370, 371,
	0, 0, 60, 57, 126, 126, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 363, 364, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 323, 65, 0, 76, 75, 0, 0,
	126, 61, 41, 67, 42, 68, 0, 0, 0, 38,
	0, 47, 0, 0, 48, 39, 40, 0, 58, 0,
	49, 0, 97, 98, 0, 126, 0, 60, 57, 86,
	87, 69, 59, 0, 77, 78, 88, 126, 0, 70,
	71, 72, 73, 74, 85, 0, 0, 0, 126, 126,
	0, 0, 28, 126, 0, 0, 0, 64, 0, 65,
	126, 76, 75, 0, 0, 97, 98, 0, 0, 126,
	0, 0, 86, 87, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 270, 0, 85, 0, 0,
	0, 0, 0, 0, 121, 0, 126, 126, 0, 0,
	0, 0, 0, 126, 0, 0, 0, 0, 0, 0,
	121, 0, 121, 121, 0, 0, 0, 121, 0, 0,
	0, 126, 0, 0, 0, 121, 121, 121, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 121, 0, 121,
	121, 121, 121, 0, 121, 121, 121, 121, 0, 121,
	0, 0, 121, 126, 121, 121, 0, 126, 0, 126,
	0, 0, 0, 0, 121, 0, 0, 121, 121, 121,
	0, 0, 126, 0, 0, 0, 0, 121, 0, 0,
	0, 0, 121, 0, 0, 121, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 126, 0,
	0, 0, 0, 0, 0, 0, 121, 121, 0, 121,
	0, 0, 126, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 121, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 121, 121,
	0, 0, 61, 41, 67, 42, 68, 0, 0, 0,
	38, 443, 368, 446, 445, 369, 39, 40, 0, 58,
	0, 49, 0, 0, 370, 371, 0, 0, 60, 57,
	0, 0, 69, 59, 121, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 363, 364,
	0, 61, 124, 67, 125, 68, 0, 0, 64, 121,
	65, 0, 76, 75, 0, 0, 0, 0, 0, 0,
	0, 121, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 121, 121, 77, 78, 44, 121, 0, 70,
	71, 72, 73, 74, 121, 0, 0, 253, 0, 0,
	0, 0, 0, 121, 378, 0, 0, 64, 0, 65,
	0, 76, 75, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 127, 0,
	121, 121, 0, 0, 0, 0, 0, 121, 0, 0,
	0, 0, 0, 0, 127, 0, 127, 127, 0, 0,
	0, 127, 0, 0, 0, 121, 0, 0, 0, 127,
	127, 127, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 127, 0, 127, 127, 127, 127, 0, 127, 127,
	127, 127, 0, 127, 0, 0, 127, 121, 127, 127,
	0, 121, 0, 121, 0, 0, 0, 0, 127, 0,
	0, 127, 127, 127, 0, 0, 121, 0, 0, 0,
	0, 127, 0, 0, 0, 0, 127, 0, 0, 127,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 121, 0, 0, 0, 0, 0, 0, 0,
	127, 127, 0, 127, 0, 0, 121, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	127, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	226, 0, 127, 127, 0, 0, 61, 41, 67, 42,
	68, 0, 0, 0, 38, 449, 368, 0, 0, 369,
	39, 40, 0, 58, 0, 49, 0, 0, 370, 371,
	0, 0, 60, 57, 0, 0, 69, 59, 127, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	143, 0, 363, 364, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 127, 65, 0, 76, 75, 0, 0,
	0, 0, 0, 0, 0, 127, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 127, 127, 0, 0,
	45, 127, 0, 0, 0, 0, 0, 0, 127, 0,
	0, 0, 0, 0, 0, 0, 0, 127, 0, 204,
	0, 206, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 222, 0, 0, 0, 0, 0, 0,
	0, 0, 128, 0, 127, 127, 0, 0, 0, 0,
	0, 127, 239, 0, 0, 0, 0, 0, 128, 0,
	128, 128, 0, 0, 0, 128, 0, 0, 0, 127,
	0, 0, 0, 128, 128, 128, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 128, 0, 128, 128, 128,
	128, 0, 128, 128, 128, 128, 0, 128, 0, 0,
	128, 127, 128, 128, 0, 127, 277, 127, 0, 0,
	0, 0, 128, 0, 284, 128, 128, 128, 0, 0,
	127, 0, 0, 0, 0, 128, 0, 0, 0, 0,
	128, 0, 0, 128, 298, 0, 301, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 127, 0, 0, 0,
	0, 309, 310, 0, 128, 128, 0, 128, 0, 0,
	127, 0, 0, 0, 0, 61, 41, 67, 42, 68,
	0, 0, 0, 38, 128, 47, 0, 0, 48, 39,
	40, 0, 58, 0, 49, 0, 128, 128, 0, 0,
	0, 60, 57, 0, 0, 69, 59, 340, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 0, 0,
	0, 6, 7, 0, 0, 0, 0, 0, 0, 356,
	0, 64, 128, 65, 360, 76, 75, 0, 8, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 390, 0, 128, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 128,
	0, 0, 0, 0, 0, 400, 401, 0, 0, 0,
	128, 128, 403, 0, 0, 128, 0, 0, 0, 0,
	0, 0, 128, 0, 411, 0, 413, 0, 0, 0,
	0, 128, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	420, 61, 124, 67, 125, 68, 239, 0, 128, 128,
	0, 0, 0, 430, 0, 128, 0, 0, 0, 0,
	0, 0, 436, 0, 309, 310, 0, 0, 0, 0,
	0, 69, 0, 128, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 0, 0, 0, 253, 0, 0,
	0, 0, 0, 0, 344, 0, 0, 64, 0, 65,
	0, 76, 75, 0, 0, 128, 0, 0, 0, 128,
	0, 128, 0, 61, 41, 67, 42, 68, 0, 0,
	0, 38, 362, 368, 128, 0, 369, 39, 40, 0,
	58, 0, 49, 0, 0, 370, 371, 0, 0, 60,
	57, 0, 0, 69, 59, 0, 77, 78, 0, 0,
	128, 70, 71, 72, 73, 74, 0, 0, 0, 363,
	364, 0, 0, 0, 128, 0, 0, 0, 0, 64,
	0, 65, 0, 76, 75, 61, 41, 67, 42, 68,
	0, 0, 0, 38, 357, 47, 308, 307, 48, 39,
	40, 0, 58, 0, 49, 0, 0, 0, 0, 0,
	0, 60, 57, 0, 0, 69, 59, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 0, 0,
	0, 242, 243, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 0, 65, 0, 76, 75, 61, 41, 67,
	42, 68, 0, 0, 0, 38, 305, 47, 308, 307,
	48, 39, 40, 0, 58, 0, 49, 0, 0, 0,
	0, 0, 0, 60, 57, 0, 0, 69, 59, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 242, 243, 0, 61, 41, 67, 42,
	68, 0, 0, 64, 38, 65, 368, 76, 75, 369,
	39, 40, 0, 58, 0, 49, 0, 0, 370, 371,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 363, 364, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 0, 76, 75, 61, 41,
	67, 42, 68, 0, 0, 0, 38, 417, 47, 0,
	0, 48, 39, 40, 0, 58, 0, 49, 313, 0,
	0, 0, 0, 0, 60, 57, 0, 0, 69, 59,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 242, 243, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 0, 65, 0, 76, 75,
	61, 41, 67, 42, 68, 0, 0, 0, 38, 338,
	47, 0, 0, 48, 39, 40, 0, 58, 0, 49,
	313, 0, 0, 0, 0, 0, 60, 57, 0, 0,
	69, 59, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 0, 0, 0, 242, 243, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 65, 0,
	76, 75, 61, 41, 67, 42, 68, 0, 0, 0,
	38, 331, 47, 0, 0, 48, 39, 40, 0, 58,
	0, 49, 313, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 242, 243,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 0,
	65, 0, 76, 75, 61, 41, 67, 42, 68, 0,
	0, 0, 38, 429, 47, 0, 0, 48, 39, 40,
	0, 58, 0, 49, 0, 0, 0, 0, 0, 0,
	60, 57, 0, 0, 69, 59, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 0, 0, 0,
	242, 243, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 0, 65, 0, 76, 75, 61, 41, 67, 42,
	68, 0, 0, 0, 38, 392, 47, 0, 0, 48,
	39, 40, 0, 58, 0, 49, 0, 0, 0, 0,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 242, 243, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 0, 76, 75, 61, 41,
	67, 42, 68, 0, 0, 0, 38, 359, 47, 0,
	0, 48, 39, 40, 0, 58, 0, 49, 0, 0,
	0, 0, 0, 0, 60, 57, 0, 0, 69, 59,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 242, 243, 0, 61, 41, 67,
	42, 68, 0, 0, 64, 38, 65, 47, 76, 75,
	48, 39, 40, 0, 58, 0, 49, 0, 0, 0,
	0, 0, 0, 60, 57, 0, 0, 69, 59, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 242, 243, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 0, 65, 350, 76, 75, 61,
	41, 67, 42, 68, 0, 0, 0, 38, 347, 47,
	0, 0, 48, 39, 40, 0, 58, 0, 49, 0,
	0, 0, 0, 0, 0, 60, 57, 0, 0, 69,
	59, 0, 77, 78, 0, 0, 0, 70, 71, 72,
	73, 74, 0, 0, 0, 242, 243, 0, 61, 41,
	67, 42, 68, 0, 0, 64, 38, 65, 47, 76,
	75, 48, 39, 40, 0, 58, 0, 49, 0, 0,
	0, 0, 0, 0, 60, 57, 0, 0, 69, 59,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 242, 243, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 0, 65, 302, 76, 75,
	61, 41, 67, 42, 68, 0, 0, 0, 38, 297,
	47, 0, 0, 48, 39, 40, 0, 58, 0, 49,
	0, 0, 0, 0, 0, 0, 60, 57, 0, 0,
	69, 59, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 0, 0, 0, 242, 243, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 65, 0,
	76, 75, 61, 41, 67, 42, 68, 0, 0, 0,
	38, 287, 47, 0, 0, 48, 39, 40, 0, 58,
	0, 49, 0, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 242, 243,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 0,
	65, 0, 76, 75, 61, 41, 67, 42, 68, 0,
	0, 0, 38, 285, 47, 0, 0, 48, 39, 40,
	0, 58, 0, 49, 0, 0, 0, 0, 0, 0,
	60, 57, 0, 0, 69, 59, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 0, 0, 0,
	242, 243, 0, 61, 41, 67, 42, 68, 0, 0,
	64, 38, 65, 47, 76, 75, 48, 39, 40, 0,
	58, 0, 49, 0, 0, 0, 0, 0, 0, 60,
	57, 0, 0, 69, 59, 0, 77, 78, 0, 0,
	0, 70, 71, 72, 73, 74, 0, 0, 0, 242,
	243, 0, 61, 41, 67, 42, 68, 263, 0, 64,
	38, 65, 47, 76, 75, 48, 39, 40, 0, 58,
	0, 49, 0, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 0, 262,
	61, 124, 67, 125, 111, 0, 117, 122, 64, 0,
	65, 0, 76, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	69, 0, 0, 77, 78, 0, 0, 115, 70, 71,
	72, 73, 74, 0, 116, 61, 124, 67, 125, 111,
	435, 0, 122, 0, 0, 0, 114, 0, 123, 0,
	76, 75, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 0, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 61, 124,
	67, 125, 111, 0, 0, 122, 0, 0, 0, 0,
	0, 114, 0, 123, 0, 76, 75, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 0,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 61, 124, 67, 125, 269, 324, 0, 122, 0,
	0, 0, 0, 0, 114, 0, 123, 0, 76, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 0, 0, 77, 78, 0, 0, 268, 70,
	71, 72, 73, 74, 61, 124, 67, 125, 68, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 0, 123,
	0, 76, 75, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 0, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 0, 0, 0,
	253, 61, 124, 67, 125, 68, 0, 0, 122, 0,
	64, 0, 65, 0, 76, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 0, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 61, 124, 67, 125, 68, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 0, 123,
	0, 76, 75, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 0, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 61, 258, 67,
	125, 68, 0, 0, 0, 0, 0, 153, 0, 0,
	64, 0, 65, 0, 76, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 0, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 253, 61, 124, 67, 125, 111, 0,
	0, 122, 0, 64, 0, 65, 0, 76, 75, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 0, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 61, 124, 67,
	125, 68, 0, 0, 0, 0, 0, 0, 0, 0,
	114, 0, 123, 0, 76, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 0, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	61, 419, 67, 125, 68, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 0, 65, 0, 76, 75, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	69, 0, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 65, 0,
	76, 75,
}
var RubyPact = []int{

	-35, 1400, -1000, -1000, -1000, 11, -1000, -1000, -1000, 346,
	-1000, -1000, -1000, 95, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 5, -5,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 415, 258,
	258, 2655, 178, -3, 123, 119, 70, 636, 636, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 3052, 636, 3052,
	3052, -1000, -1000, -1000, 2919, -1000, 3, 221, -1000, 25,
	636, 636, 3052, 3052, 3052, 28, 392, -1000, -1000, -1000,
	-1000, -1000, 636, 636, 3052, 389, 3052, 3052, 3052, 3052,
	636, 3052, 3052, 3052, 3052, 636, 3052, -1000, -1000, 3052,
	636, 3052, 3052, 636, 636, 367, 275, 177, 34, -1000,
	-1000, 2919, 25, 32, 2919, 3052, 3052, 362, 191, 519,
	-1000, 15, -15, -15, 3009, 141, -30, -1000, -1000, 2919,
	636, 636, 3052, 636, 636, 636, 636, 636, 636, 636,
	349, 289, 308, 2558, 164, 519, 139, 519, 164, 91,
	58, 255, -1000, 3052, 2962, 263, 2919, 2607, -1000, -15,
	289, 289, 519, 519, 519, -1000, -1000, 256, -1000, -1000,
	289, 289, 519, 2786, 519, 519, 519, 519, 289, 519,
	519, 519, 519, 289, 672, 2829, 2829, 519, 289, 519,
	57, 122, 289, 289, 25, -1000, 345, 186, 71, -1000,
	80, 319, 313, -1000, 2509, 258, 2447, 379, 255, -1000,
	-1000, -1000, 52, -43, 53, 491, -1000, -1000, 302, -1000,
	-1000, 2876, 2385, -1000, 309, -1000, 2323, 356, 289, 289,
	20, 289, 289, 289, 289, 289, 289, 289, -1000, 1742,
	-1000, -1000, -1000, -1000, 289, 280, 3052, -1000, 66, -1000,
	-1000, -1000, 519, -1000, 181, 171, 22, 629, 2743, -1000,
	352, 289, -1000, -1000, -1000, -1000, 32, 25, 636, 2919,
	3052, 519, 519, -1000, 2876, 79, -1000, 1977, 177, 71,
	218, 3052, -1000, -1000, 1915, -1000, -1000, -1000, 25, -1000,
	1546, 48, -1000, -1000, -11, 519, -1000, -1000, 2274, 31,
	-1000, 2212, -1000, -1000, 47, -1000, 399, 636, -1000, 1680,
	2163, -1000, -1000, 249, 519, 1618, 260, 3052, 926, -26,
	-1000, -29, -1000, 636, 3052, -1000, -1000, 289, 149, 519,
	636, -1000, 242, -1000, -1000, -1000, -1000, 519, -1000, 224,
	2101, -1000, 207, 519, 303, 636, 297, -1000, -1000, 292,
	-1000, 636, -1000, 636, -1000, 289, 2558, -1000, 358, -1000,
	2558, 325, -1000, -1000, -1000, 289, -1000, -1000, 636, 636,
	330, 329, -1000, -1000, 3052, 164, 255, -1000, 3052, -1000,
	2829, -1000, 168, 346, 289, 519, -1000, 289, -1000, -1000,
	1853, -1000, -1000, 3095, -1000, 289, -1000, -1000, 289, 289,
	2558, 2558, -1000, 2558, 273, 134, 192, 636, 636, 636,
	636, 2039, 164, 2558, 519, 200, 12, -1000, 223, 2700,
	2558, -1000, -1000, -1000, -1000, 289, 289, 289, 289, -1000,
	2558, 22, 636, 3052, -1000, -1000, 2558, 877, 571, 1151,
	22, 158, 403, -1000, 323, 636, -1000, -1000, 304, -1000,
	-1000, -1000, 22, -1000, -1000, 636, -1000, 289, 1791, -1000,
	22, 289, 1791, 1791, 1791,
}
var RubyPgo = []int{

	0, 0, 502, 501, 32, 10, 499, 498, 497, 1240,
	496, 1, 28, 494, 490, 489, 488, 106, 486, 966,
	692, 484, 482, 480, 13, 477, 6, 418, 475, 474,
	11, 473, 472, 471, 470, 30, 469, 468, 467, 3,
	457, 455, 454, 453, 452, 447, 445, 441, 435, 434,
	432, 1150, 426, 5, 15, 18, 7, 423, 8, 422,
	4, 412, 21, 406, 9, 399, 27, 14, 12, 385,
	358, 127,
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
	20, 20, 20, 54, 54, 54, 54, 64, 64, 62,
	62, 62, 62, 62, 62, 62, 67, 67, 67, 67,
	67, 66, 66, 66, 21, 21, 21, 21, 21, 21,
	58, 58, 68, 68, 68, 26, 26, 26, 26, 25,
	25, 28, 30, 30, 69, 69, 15, 15, 15, 15,
	15, 15, 15, 15, 29, 29, 29, 29, 29, 29,
	9, 9, 27, 27, 19, 19, 40, 40, 41, 42,
	43, 44, 45, 46, 47, 48, 49, 50, 2, 6,
	7, 7, 3, 3, 59, 59, 59, 59, 65, 65,
	65, 5, 5, 5, 5, 55, 63, 63, 63, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	56, 56, 56, 56, 52, 52, 52, 8, 16, 11,
	11, 11, 61, 61, 53, 53, 22, 22, 23, 23,
	12, 36, 60, 60, 60, 60, 60, 60, 37, 37,
	37, 37, 37, 37, 37, 38, 38, 38, 38, 38,
	39, 39, 39, 39, 34, 33, 10, 32, 32, 31,
	31, 4,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 1, 0, 2, 0, 2, 2, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 2, 4, 5, 1, 4, 4,
	2, 3, 3, 4, 4, 5, 3, 4, 5, 2,
	3, 3, 3, 3, 4, 4, 4, 4, 4, 4,
	6, 6, 6, 3, 7, 1, 5, 1, 3, 0,
	1, 1, 2, 4, 4, 5, 1, 1, 4, 2,
	5, 1, 3, 3, 5, 6, 7, 8, 5, 6,
	1, 3, 0, 1, 3, 1, 2, 3, 2, 4,
	6, 4, 1, 3, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 3, 3, 3, 3, 2, 2,
	2, 2, 3, 3, 3, 3, 3, 3, 1, 1,
	3, 3, 5, 5, 0, 1, 3, 7, 3, 7,
	8, 3, 4, 4, 3, 3, 0, 1, 3, 4,
	5, 3, 3, 3, 3, 3, 5, 6, 5, 3,
	4, 3, 3, 2, 0, 2, 2, 3, 4, 2,
	3, 5, 0, 2, 1, 2, 2, 1, 2, 1,
	5, 5, 0, 2, 2, 2, 2, 2, 0, 1,
	3, 3, 1, 3, 3, 5, 6, 5, 6, 5,
	4, 3, 3, 2, 4, 4, 2, 5, 7, 4,
	5, 3,
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
	53, -17, -17, 62, 11, 62, -5, -51, 6, 11,
	-68, 42, 6, 6, -51, 14, -30, 14, 10, 11,
	-71, 62, 62, 62, -71, -17, -5, 14, -51, -63,
	6, -51, 64, 10, 62, 14, -56, 17, 16, -51,
	-51, 14, -11, 25, -17, -60, -31, 37, -71, -71,
	11, -71, 11, 4, 53, 10, -5, -1, -62, -17,
	42, 14, -53, -11, -58, -26, 10, -17, 14, -53,
	-51, -5, -71, -17, 58, 42, 58, 14, 56, 11,
	64, 42, 14, 17, 16, -1, -51, 14, -56, 14,
	-51, 8, 14, 51, 52, -1, -38, -37, 15, 18,
	27, 28, 14, 16, 37, -64, -17, -24, 58, 64,
	-71, 64, -71, -17, -1, -17, 10, -1, 14, -11,
	-51, 14, 14, 58, 6, -1, 6, 6, -1, -1,
	-51, -51, 14, -51, 4, -1, -1, 15, 18, 15,
	18, -51, -64, -51, -17, -17, 6, 14, -53, 6,
	-51, 6, 51, 51, 52, -1, -1, -1, -1, 14,
	-51, -71, 4, 53, 14, 10, -51, -60, -60, -60,
	-71, -1, -17, 14, -39, 17, 16, 14, -39, 14,
	-70, 11, -71, 11, 14, 17, 16, -1, -60, 14,
	-71, -1, -60, -60, -60,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 48,
	49, 50, 51, 52, 53, 54, 55, 56, 57, 58,
	59, 60, 61, 62, 27, 28, 29, 30, 31, 32,
	33, 34, 35, 36, 37, 38, 39, 40, 0, 0,
	0, 20, -2, 22, 23, 24, 0, 0, 0, 15,
	41, 42, 43, 44, 45, 46, 47, 217, 0, 0,
	219, 19, 25, 26, 99, 13, 0, 67, 204, 0,
	0, 0, 0, 0, 0, 0, 0, 168, 169, 5,
	6, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 13, 13, 0,
	0, 0, 0, 0, 0, 0, 122, 122, 15, -2,
	15, -2, 70, 79, 99, 0, 0, 0, 95, 106,
	107, 31, 15, -2, 20, -2, 22, 23, 24, 99,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 15, 0, 212, 216, 97, 0, 13, 218, 0,
	0, 97, 101, 0, 13, 0, 99, 0, 246, 15,
	158, 159, 160, 161, 64, 152, 153, 0, 150, 151,
	191, 199, 63, 72, 80, 82, 83, 162, 163, 164,
	165, 166, 167, 193, 0, 0, 0, 251, 195, 81,
	0, 111, 192, 194, 76, 15, 0, 120, 122, 123,
	125, 0, 0, 15, 0, 0, 0, 0, 100, 71,
	13, 109, 97, 0, 0, 136, 137, 138, 144, 145,
	156, 13, 0, 15, 186, 15, 0, 0, 139, 146,
	0, 140, 147, 141, 148, 142, 149, 143, 157, 0,
	15, 15, 16, 17, 18, 0, 0, 222, 0, 170,
	13, 171, 102, 14, 13, 13, 175, 0, 20, -2,
	0, 205, 206, 207, 154, 155, 73, 74, 0, -2,
	0, 244, 245, 88, 0, 89, 77, 0, 122, 0,
	0, 0, 126, 128, 0, 129, 15, 131, 65, 13,
	0, 84, 86, 87, 0, 112, 113, 181, 0, 0,
	187, 0, 184, 69, 85, 189, 0, 0, 15, 0,
	0, 208, 213, 15, 98, 0, 0, 0, 0, 0,
	13, 0, 13, 0, 0, 68, 75, 78, 0, 220,
	0, 114, 0, 214, 15, 124, 121, 127, 118, 0,
	0, 66, 0, 108, 0, 0, 0, 182, 185, 0,
	183, 0, 190, 0, 15, 15, 203, 196, 0, 198,
	209, 15, 221, 223, 224, 225, 226, 227, 0, 0,
	229, 232, 247, 15, 0, 15, 103, 104, 0, 172,
	0, 173, 0, -2, 176, 178, 93, 92, 115, 215,
	0, 119, 130, 0, 110, 90, 96, 188, 91, 15,
	201, 202, 197, 210, 0, 15, 0, 0, 0, 0,
	0, 0, 15, 13, 105, 0, 0, 116, 0, 20,
	200, 15, 222, -2, -2, 230, 231, 233, 234, 248,
	13, 249, 0, 0, 117, 94, 211, 0, 0, 0,
	250, 11, 13, 235, 0, 0, 222, 237, 0, 239,
	177, 12, 179, 13, 236, 0, 222, 222, 228, 238,
	180, 222, 228, 228, 228,
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
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 78:
		//line parser.y:345
		{
			methodName := RubyS[Rubypt-2].genericValue.(ast.BareReference).Name + "="
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: methodName},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 79:
		//line parser.y:356
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
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
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 82:
		//line parser.y:379
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 83:
		//line parser.y:387
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
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
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
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
		//line parser.y:429
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 89:
		//line parser.y:437
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 90:
		//line parser.y:447
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 91:
		//line parser.y:455
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 92:
		//line parser.y:463
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 93:
		//line parser.y:473
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 94:
		//line parser.y:475
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 95:
		//line parser.y:477
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 96:
		//line parser.y:479
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 97:
		//line parser.y:482
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 98:
		//line parser.y:484
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 99:
		//line parser.y:486
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 100:
		//line parser.y:488
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 101:
		//line parser.y:490
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 102:
		//line parser.y:492
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{Func: ast.BareReference{Name: "to_proc"}, Target: RubyS[Rubypt-0].genericValue})
		}
	case 103:
		//line parser.y:494
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 104:
		//line parser.y:496
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 105:
		//line parser.y:498
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{
				Func:   ast.BareReference{Name: "to_proc"},
				Target: RubyS[Rubypt-0].genericValue,
			})
		}
	case 106:
		//line parser.y:507
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 107:
		//line parser.y:509
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 108:
		//line parser.y:511
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 109:
		//line parser.y:513
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 110:
		//line parser.y:515
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 111:
		//line parser.y:518
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 112:
		//line parser.y:520
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 113:
		//line parser.y:522
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 114:
		//line parser.y:526
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 115:
		//line parser.y:534
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 116:
		//line parser.y:543
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 117:
		//line parser.y:552
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 118:
		//line parser.y:562
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 119:
		//line parser.y:570
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 120:
		//line parser.y:581
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 121:
		//line parser.y:583
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 122:
		//line parser.y:585
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 123:
		//line parser.y:587
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 124:
		//line parser.y:589
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 125:
		//line parser.y:592
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 126:
		//line parser.y:594
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 127:
		//line parser.y:596
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 128:
		//line parser.y:598
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 129:
		//line parser.y:602
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 130:
		//line parser.y:609
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 131:
		//line parser.y:619
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 132:
		//line parser.y:628
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 133:
		//line parser.y:634
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 134:
		//line parser.y:642
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 135:
		//line parser.y:646
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 136:
		//line parser.y:651
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 137:
		//line parser.y:658
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 138:
		//line parser.y:662
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 139:
		//line parser.y:666
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 140:
		//line parser.y:673
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 141:
		//line parser.y:680
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 142:
		//line parser.y:687
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 143:
		//line parser.y:694
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 144:
		//line parser.y:702
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 145:
		//line parser.y:709
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 146:
		//line parser.y:713
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 147:
		//line parser.y:720
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 148:
		//line parser.y:727
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 149:
		//line parser.y:734
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 150:
		//line parser.y:742
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 151:
		//line parser.y:744
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 152:
		//line parser.y:747
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 153:
		//line parser.y:749
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 154:
		//line parser.y:752
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 155:
		//line parser.y:754
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 156:
		//line parser.y:757
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 157:
		//line parser.y:759
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 158:
		//line parser.y:761
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 159:
		//line parser.y:762
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 160:
		//line parser.y:763
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 161:
		//line parser.y:764
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 162:
		//line parser.y:767
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 163:
		//line parser.y:776
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 164:
		//line parser.y:785
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 165:
		//line parser.y:794
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 166:
		//line parser.y:803
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 167:
		//line parser.y:812
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 168:
		//line parser.y:820
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 169:
		//line parser.y:821
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 170:
		//line parser.y:823
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 171:
		//line parser.y:825
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 172:
		//line parser.y:828
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 173:
		//line parser.y:836
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 174:
		//line parser.y:844
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 175:
		//line parser.y:846
		{
		}
	case 176:
		//line parser.y:848
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 177:
		//line parser.y:855
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 178:
		//line parser.y:863
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 179:
		//line parser.y:870
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 180:
		//line parser.y:877
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 181:
		//line parser.y:885
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 182:
		//line parser.y:887
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 183:
		//line parser.y:894
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 184:
		//line parser.y:901
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 185:
		//line parser.y:904
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 186:
		//line parser.y:906
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 187:
		//line parser.y:908
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 188:
		//line parser.y:910
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 189:
		//line parser.y:913
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 190:
		//line parser.y:920
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 191:
		//line parser.y:928
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 192:
		//line parser.y:935
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 193:
		//line parser.y:942
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 194:
		//line parser.y:949
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 195:
		//line parser.y:956
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
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
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 198:
		//line parser.y:978
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 199:
		//line parser.y:985
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 200:
		//line parser.y:994
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 201:
		//line parser.y:1001
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 202:
		//line parser.y:1008
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 203:
		//line parser.y:1015
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 204:
		//line parser.y:1022
		{
		}
	case 205:
		//line parser.y:1023
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 206:
		//line parser.y:1024
		{
		}
	case 207:
		//line parser.y:1027
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 208:
		//line parser.y:1030
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 209:
		//line parser.y:1038
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 210:
		//line parser.y:1040
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 211:
		//line parser.y:1049
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
	case 212:
		//line parser.y:1064
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 213:
		//line parser.y:1066
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 214:
		//line parser.y:1069
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 215:
		//line parser.y:1071
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 216:
		//line parser.y:1074
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 217:
		//line parser.y:1081
		{
			RubyVAL.genericValue = ast.Yield{}
		}
	case 218:
		//line parser.y:1084
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 219:
		//line parser.y:1092
		{
			RubyVAL.genericValue = ast.Return{}
		}
	case 220:
		//line parser.y:1095
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 221:
		//line parser.y:1104
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 222:
		//line parser.y:1107
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 223:
		//line parser.y:1109
		{
		}
	case 224:
		//line parser.y:1111
		{
		}
	case 225:
		//line parser.y:1113
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 226:
		//line parser.y:1115
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 227:
		//line parser.y:1117
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 228:
		//line parser.y:1119
		{
		}
	case 229:
		//line parser.y:1121
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 230:
		//line parser.y:1123
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 231:
		//line parser.y:1125
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 232:
		//line parser.y:1127
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 233:
		//line parser.y:1129
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 234:
		//line parser.y:1131
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 235:
		//line parser.y:1134
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 236:
		//line parser.y:1141
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 237:
		//line parser.y:1149
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 238:
		//line parser.y:1156
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 239:
		//line parser.y:1164
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 240:
		//line parser.y:1172
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 241:
		//line parser.y:1179
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 242:
		//line parser.y:1186
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 243:
		//line parser.y:1193
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 244:
		//line parser.y:1201
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 245:
		//line parser.y:1204
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 246:
		//line parser.y:1206
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 247:
		//line parser.y:1209
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 248:
		//line parser.y:1211
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 249:
		//line parser.y:1214
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 250:
		//line parser.y:1216
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 251:
		//line parser.y:1218
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
