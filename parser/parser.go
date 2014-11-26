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

//line parser.y:1252

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 44,
	54, 136,
	-2, 21,
	-1, 114,
	54, 136,
	-2, 134,
	-1, 116,
	10, 101,
	11, 101,
	-2, 209,
	-1, 128,
	13, 15,
	15, 15,
	18, 15,
	19, 15,
	20, 15,
	22, 15,
	24, 15,
	27, 15,
	28, 15,
	31, 15,
	32, 15,
	36, 15,
	52, 15,
	-2, 13,
	-1, 130,
	54, 136,
	-2, 21,
	-1, 278,
	54, 137,
	-2, 135,
	-1, 288,
	10, 101,
	11, 101,
	-2, 209,
	-1, 346,
	51, 13,
	-2, 228,
	-1, 451,
	51, 13,
	-2, 228,
}

const RubyNprod = 256
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 3242

var RubyAct = []int{

	259, 237, 5, 160, 465, 355, 206, 329, 125, 210,
	356, 25, 123, 29, 49, 208, 272, 161, 118, 236,
	117, 87, 2, 3, 88, 272, 312, 59, 138, 405,
	24, 176, 235, 177, 108, 139, 93, 109, 403, 4,
	221, 127, 127, 442, 294, 327, 294, 137, 269, 93,
	106, 148, 149, 261, 150, 113, 115, 85, 84, 106,
	326, 137, 153, 373, 151, 102, 103, 155, 272, 110,
	211, 107, 91, 92, 86, 370, 171, 172, 102, 103,
	107, 338, 166, 169, 452, 91, 92, 90, 181, 182,
	272, 178, 170, 170, 463, 313, 189, 293, 17, 270,
	90, 194, 45, 368, 268, 142, 199, 424, 372, 203,
	204, 212, 143, 376, 340, 215, 140, 217, 353, 214,
	201, 211, 213, 141, 209, 301, 220, 222, 272, 233,
	216, 224, 225, 369, 218, 211, 240, 241, 209, 243,
	244, 245, 246, 247, 248, 249, 131, 234, 228, 230,
	254, 227, 239, 87, 105, 261, 88, 264, 265, 266,
	267, 251, 212, 131, 368, 131, 131, 147, 163, 280,
	207, 134, 131, 213, 167, 87, 212, 453, 88, 369,
	131, 131, 131, 409, 309, 279, 438, 213, 336, 446,
	447, 93, 131, 87, 131, 131, 88, 131, 146, 131,
	131, 131, 131, 285, 131, 286, 345, 131, 297, 131,
	131, 255, 256, 343, 296, 163, 304, 145, 163, 131,
	102, 103, 131, 131, 131, 300, 93, 91, 92, 445,
	306, 272, 131, 163, 414, 147, 318, 131, 321, 134,
	131, 316, 90, 299, 87, 336, 93, 88, 144, 423,
	224, 225, 232, 411, 334, 102, 103, 332, 333, 359,
	299, 388, 91, 92, 336, 336, 163, 131, 131, 114,
	131, 335, 135, 278, 472, 102, 103, 90, 87, 136,
	262, 88, 91, 92, 325, 444, 87, 131, 350, 88,
	131, 93, 421, 158, 222, 349, 159, 90, 156, 131,
	131, 157, 420, 316, 295, 357, 351, 419, 363, 358,
	362, 396, 89, 397, 348, 269, 283, 364, 284, 93,
	102, 103, 480, 417, 477, 476, 320, 91, 92, 324,
	269, 382, 303, 383, 398, 131, 302, 163, 387, 392,
	385, 131, 90, 179, 399, 180, 101, 407, 102, 103,
	475, 298, 477, 476, 410, 91, 92, 430, 428, 413,
	381, 380, 308, 309, 131, 379, 412, 381, 380, 418,
	90, 253, 252, 412, 101, 250, 112, 422, 111, 231,
	131, 425, 205, 426, 427, 184, 451, 163, 471, 71,
	429, 131, 274, 131, 319, 431, 432, 131, 260, 433,
	273, 435, 434, 1, 131, 93, 168, 58, 57, 56,
	55, 54, 474, 131, 53, 102, 103, 36, 30, 440,
	35, 34, 91, 92, 33, 48, 393, 443, 19, 131,
	131, 38, 39, 254, 102, 103, 449, 90, 20, 339,
	401, 91, 92, 131, 131, 14, 456, 12, 332, 333,
	131, 412, 461, 11, 23, 22, 90, 21, 392, 392,
	392, 18, 126, 469, 401, 10, 26, 478, 131, 16,
	13, 37, 15, 32, 31, 27, 68, 482, 28, 126,
	392, 126, 126, 67, 392, 392, 392, 0, 126, 0,
	0, 0, 0, 0, 0, 0, 126, 126, 126, 0,
	0, 131, 46, 0, 0, 131, 0, 131, 126, 0,
	126, 126, 0, 126, 0, 126, 126, 126, 126, 131,
	126, 0, 0, 126, 0, 126, 126, 457, 458, 459,
	0, 0, 0, 0, 0, 126, 0, 0, 126, 126,
	126, 0, 0, 0, 131, 131, 132, 0, 126, 479,
	0, 0, 0, 126, 0, 131, 126, 0, 0, 483,
	484, 0, 0, 132, 485, 132, 132, 0, 0, 0,
	0, 0, 132, 0, 0, 0, 0, 0, 0, 0,
	132, 132, 132, 126, 126, 0, 126, 66, 129, 72,
	130, 73, 132, 0, 132, 132, 0, 132, 0, 132,
	132, 132, 132, 126, 132, 93, 126, 132, 0, 132,
	132, 0, 0, 0, 0, 126, 126, 74, 0, 132,
	82, 83, 132, 132, 132, 75, 76, 77, 78, 79,
	0, 0, 132, 272, 102, 103, 0, 132, 0, 0,
	132, 91, 92, 69, 0, 70, 93, 81, 80, 0,
	0, 126, 0, 0, 290, 346, 90, 126, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 132, 132, 0,
	132, 0, 0, 0, 0, 102, 103, 0, 0, 0,
	126, 0, 91, 92, 102, 103, 0, 132, 0, 0,
	132, 91, 92, 0, 0, 0, 126, 90, 0, 132,
	132, 0, 0, 0, 0, 0, 90, 126, 0, 126,
	0, 0, 0, 126, 66, 129, 72, 130, 116, 454,
	126, 127, 0, 0, 0, 0, 0, 0, 0, 126,
	0, 0, 0, 0, 0, 132, 0, 0, 0, 0,
	0, 132, 0, 0, 74, 126, 126, 82, 83, 0,
	0, 0, 75, 76, 77, 78, 79, 0, 0, 126,
	126, 0, 0, 0, 132, 0, 126, 0, 0, 0,
	238, 0, 128, 93, 81, 80, 0, 0, 0, 0,
	132, 0, 0, 0, 126, 0, 0, 0, 0, 0,
	0, 132, 0, 132, 0, 9, 0, 132, 0, 0,
	0, 0, 102, 103, 132, 0, 0, 0, 0, 91,
	92, 0, 0, 132, 94, 95, 96, 126, 0, 0,
	0, 126, 0, 126, 90, 99, 97, 98, 101, 132,
	132, 0, 0, 0, 0, 126, 0, 0, 0, 124,
	0, 0, 0, 132, 132, 0, 0, 0, 0, 0,
	132, 0, 0, 0, 0, 0, 152, 0, 154, 152,
	126, 126, 0, 0, 0, 162, 0, 0, 132, 0,
	0, 126, 0, 173, 174, 175, 0, 0, 0, 47,
	0, 0, 0, 0, 0, 183, 0, 185, 186, 0,
	188, 0, 190, 191, 192, 193, 0, 195, 0, 0,
	198, 132, 200, 202, 0, 132, 0, 132, 0, 0,
	0, 0, 219, 0, 0, 223, 226, 229, 0, 132,
	0, 0, 0, 133, 0, 124, 0, 0, 0, 0,
	219, 0, 0, 242, 0, 0, 0, 0, 0, 0,
	133, 0, 133, 133, 132, 132, 0, 0, 0, 133,
	0, 0, 0, 0, 0, 132, 0, 133, 133, 133,
	271, 276, 0, 219, 0, 0, 187, 0, 0, 133,
	0, 133, 133, 0, 133, 0, 133, 133, 133, 133,
	124, 133, 0, 289, 133, 0, 133, 133, 0, 0,
	0, 0, 291, 292, 0, 0, 133, 0, 0, 133,
	133, 133, 0, 0, 0, 0, 0, 0, 0, 133,
	0, 0, 0, 0, 133, 0, 0, 133, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 315, 0,
	0, 0, 0, 0, 323, 0, 0, 165, 0, 0,
	0, 0, 0, 0, 133, 133, 0, 133, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 337, 0, 0,
	0, 0, 0, 0, 133, 0, 0, 133, 0, 196,
	197, 0, 0, 124, 0, 0, 133, 133, 0, 0,
	0, 0, 0, 0, 219, 0, 352, 0, 0, 0,
	315, 0, 0, 0, 0, 165, 0, 360, 0, 0,
	0, 0, 0, 0, 0, 0, 366, 0, 0, 0,
	0, 0, 133, 0, 0, 0, 0, 0, 133, 0,
	0, 263, 377, 378, 0, 0, 0, 0, 0, 0,
	0, 0, 275, 0, 0, 0, 152, 400, 0, 0,
	0, 133, 0, 408, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 133, 0, 0,
	0, 400, 0, 0, 0, 0, 0, 0, 133, 0,
	133, 0, 0, 0, 133, 0, 0, 0, 0, 0,
	0, 133, 0, 0, 0, 0, 0, 0, 310, 0,
	133, 0, 0, 0, 152, 0, 0, 0, 436, 314,
	437, 0, 0, 0, 0, 0, 133, 133, 0, 0,
	0, 0, 436, 0, 0, 0, 0, 0, 0, 0,
	133, 133, 0, 0, 0, 0, 0, 133, 0, 0,
	0, 0, 0, 0, 0, 0, 341, 124, 455, 0,
	342, 344, 0, 0, 0, 133, 0, 0, 462, 0,
	0, 0, 0, 66, 129, 72, 130, 116, 0, 122,
	127, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 365, 0, 133, 0,
	0, 0, 133, 74, 133, 0, 82, 83, 0, 0,
	120, 75, 76, 77, 78, 79, 133, 121, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 119,
	404, 128, 406, 81, 80, 0, 66, 43, 72, 44,
	73, 133, 133, 0, 40, 468, 394, 467, 466, 395,
	41, 42, 133, 61, 0, 52, 0, 0, 64, 65,
	0, 0, 63, 60, 0, 0, 74, 62, 0, 82,
	83, 0, 0, 0, 75, 76, 77, 78, 79, 0,
	0, 0, 390, 391, 0, 0, 0, 0, 0, 0,
	0, 0, 69, 0, 70, 0, 81, 80, 0, 0,
	0, 0, 0, 0, 66, 43, 72, 44, 73, 0,
	0, 0, 40, 464, 394, 467, 466, 395, 41, 42,
	0, 61, 450, 52, 0, 0, 64, 65, 0, 0,
	63, 60, 0, 0, 74, 62, 460, 82, 83, 0,
	0, 0, 75, 76, 77, 78, 79, 0, 0, 473,
	390, 391, 0, 0, 0, 0, 0, 0, 0, 0,
	69, 481, 70, 0, 81, 80, 66, 43, 72, 44,
	73, 0, 0, 0, 40, 384, 50, 331, 330, 51,
	41, 42, 0, 61, 0, 52, 0, 0, 64, 65,
	0, 0, 63, 60, 0, 0, 74, 62, 0, 82,
	83, 0, 0, 0, 75, 76, 77, 78, 79, 0,
	0, 0, 257, 258, 0, 0, 0, 0, 0, 0,
	0, 0, 69, 0, 70, 0, 81, 80, 66, 43,
	72, 44, 73, 0, 0, 0, 40, 328, 50, 331,
	330, 51, 41, 42, 0, 61, 0, 52, 0, 0,
	64, 65, 0, 0, 63, 60, 0, 0, 74, 62,
	0, 82, 83, 0, 0, 0, 75, 76, 77, 78,
	79, 0, 0, 0, 257, 258, 0, 66, 43, 72,
	44, 73, 0, 0, 69, 40, 70, 50, 81, 80,
	51, 41, 42, 0, 61, 0, 52, 0, 0, 64,
	65, 0, 0, 63, 60, 0, 0, 74, 62, 0,
	82, 83, 0, 0, 0, 75, 76, 77, 78, 79,
	0, 0, 0, 6, 7, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 0, 70, 0, 81, 80, 0,
	8, 66, 43, 72, 44, 73, 0, 0, 0, 40,
	439, 50, 0, 0, 51, 41, 42, 0, 61, 0,
	52, 336, 0, 64, 65, 0, 0, 63, 60, 0,
	0, 74, 62, 0, 82, 83, 0, 0, 0, 75,
	76, 77, 78, 79, 0, 0, 0, 257, 258, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 0, 70,
	0, 81, 80, 66, 43, 72, 44, 73, 0, 0,
	0, 40, 361, 50, 0, 0, 51, 41, 42, 0,
	61, 0, 52, 336, 0, 64, 65, 0, 0, 63,
	60, 0, 0, 74, 62, 0, 82, 83, 0, 0,
	0, 75, 76, 77, 78, 79, 0, 0, 0, 257,
	258, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	0, 70, 0, 81, 80, 66, 43, 72, 44, 73,
	0, 0, 0, 40, 354, 50, 0, 0, 51, 41,
	42, 0, 61, 0, 52, 336, 0, 64, 65, 0,
	0, 63, 60, 0, 0, 74, 62, 0, 82, 83,
	0, 0, 0, 75, 76, 77, 78, 79, 0, 0,
	0, 257, 258, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 0, 70, 0, 81, 80, 66, 43, 72,
	44, 73, 0, 0, 0, 40, 470, 394, 0, 0,
	395, 41, 42, 0, 61, 0, 52, 0, 0, 64,
	65, 0, 0, 63, 60, 0, 0, 74, 62, 0,
	82, 83, 0, 0, 0, 75, 76, 77, 78, 79,
	0, 0, 0, 390, 391, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 0, 70, 0, 81, 80, 66,
	43, 72, 44, 73, 0, 0, 0, 40, 448, 50,
	0, 0, 51, 41, 42, 0, 61, 0, 52, 0,
	0, 64, 65, 0, 0, 63, 60, 0, 0, 74,
	62, 0, 82, 83, 0, 0, 0, 75, 76, 77,
	78, 79, 0, 0, 0, 257, 258, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 0, 70, 0, 81,
	80, 66, 43, 72, 44, 73, 0, 0, 0, 40,
	415, 50, 0, 0, 51, 41, 42, 0, 61, 0,
	52, 0, 0, 64, 65, 0, 0, 63, 60, 0,
	0, 74, 62, 0, 82, 83, 0, 0, 0, 75,
	76, 77, 78, 79, 0, 0, 0, 257, 258, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 0, 70,
	0, 81, 80, 66, 43, 72, 44, 73, 0, 0,
	0, 40, 389, 394, 0, 0, 395, 41, 42, 0,
	61, 0, 52, 0, 0, 64, 65, 0, 0, 63,
	60, 0, 0, 74, 62, 0, 82, 83, 0, 0,
	0, 75, 76, 77, 78, 79, 0, 0, 0, 390,
	391, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	0, 70, 0, 81, 80, 66, 43, 72, 44, 73,
	0, 0, 0, 40, 386, 50, 0, 0, 51, 41,
	42, 0, 61, 0, 52, 0, 0, 64, 65, 0,
	0, 63, 60, 0, 0, 74, 62, 0, 82, 83,
	0, 0, 0, 75, 76, 77, 78, 79, 0, 0,
	0, 257, 258, 0, 66, 43, 72, 44, 73, 0,
	0, 69, 40, 70, 50, 81, 80, 51, 41, 42,
	0, 61, 0, 52, 0, 0, 64, 65, 0, 0,
	63, 60, 0, 0, 74, 62, 0, 82, 83, 0,
	0, 0, 75, 76, 77, 78, 79, 0, 0, 0,
	257, 258, 0, 0, 0, 0, 0, 0, 0, 0,
	69, 0, 70, 374, 81, 80, 66, 43, 72, 44,
	73, 0, 0, 0, 40, 371, 50, 0, 0, 51,
	41, 42, 0, 61, 0, 52, 0, 0, 64, 65,
	0, 0, 63, 60, 0, 0, 74, 62, 0, 82,
	83, 0, 0, 0, 75, 76, 77, 78, 79, 0,
	0, 0, 257, 258, 0, 66, 43, 72, 44, 73,
	0, 0, 69, 40, 70, 50, 81, 80, 51, 41,
	42, 0, 61, 0, 52, 0, 0, 64, 65, 0,
	0, 63, 60, 0, 0, 74, 62, 0, 82, 83,
	0, 0, 0, 75, 76, 77, 78, 79, 0, 0,
	0, 257, 258, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 0, 70, 322, 81, 80, 66, 43, 72,
	44, 73, 0, 0, 0, 40, 317, 50, 0, 0,
	51, 41, 42, 0, 61, 0, 52, 0, 0, 64,
	65, 0, 0, 63, 60, 0, 0, 74, 62, 0,
	82, 83, 0, 0, 0, 75, 76, 77, 78, 79,
	0, 0, 0, 257, 258, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 0, 70, 0, 81, 80, 66,
	43, 72, 44, 73, 0, 0, 0, 40, 307, 50,
	0, 0, 51, 41, 42, 0, 61, 0, 52, 0,
	0, 64, 65, 0, 0, 63, 60, 0, 0, 74,
	62, 0, 82, 83, 0, 0, 0, 75, 76, 77,
	78, 79, 0, 0, 0, 257, 258, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 0, 70, 0, 81,
	80, 66, 43, 72, 44, 73, 0, 0, 0, 40,
	305, 50, 0, 0, 51, 41, 42, 0, 61, 0,
	52, 0, 0, 64, 65, 0, 0, 63, 60, 0,
	0, 74, 62, 0, 82, 83, 0, 0, 0, 75,
	76, 77, 78, 79, 0, 0, 0, 257, 258, 0,
	66, 43, 72, 44, 73, 0, 0, 69, 40, 70,
	394, 81, 80, 395, 41, 42, 0, 61, 0, 52,
	0, 0, 64, 65, 0, 0, 63, 60, 0, 0,
	74, 62, 0, 82, 83, 0, 0, 0, 75, 76,
	77, 78, 79, 0, 0, 0, 390, 391, 0, 66,
	43, 72, 44, 73, 0, 0, 69, 40, 70, 50,
	81, 80, 51, 41, 42, 0, 61, 0, 52, 0,
	0, 64, 65, 0, 0, 63, 60, 0, 0, 74,
	62, 0, 82, 83, 0, 0, 0, 75, 76, 77,
	78, 79, 0, 0, 0, 257, 258, 0, 66, 43,
	72, 44, 73, 282, 0, 69, 40, 70, 50, 81,
	80, 51, 41, 42, 0, 61, 0, 52, 0, 0,
	64, 65, 0, 0, 63, 60, 0, 0, 74, 62,
	0, 82, 83, 0, 0, 0, 75, 76, 77, 78,
	79, 0, 0, 0, 0, 281, 0, 66, 43, 72,
	44, 73, 0, 0, 69, 40, 70, 50, 81, 80,
	51, 41, 42, 0, 61, 0, 52, 0, 0, 64,
	65, 0, 0, 63, 60, 0, 0, 74, 62, 0,
	82, 83, 0, 0, 0, 75, 76, 77, 78, 79,
	66, 129, 72, 130, 73, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 0, 70, 0, 81, 80, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	74, 0, 0, 82, 83, 0, 0, 0, 75, 76,
	77, 78, 79, 0, 0, 0, 272, 66, 129, 72,
	130, 73, 0, 416, 0, 0, 69, 0, 70, 0,
	81, 80, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 74, 0, 0,
	82, 83, 0, 0, 0, 75, 76, 77, 78, 79,
	0, 0, 0, 272, 66, 129, 72, 130, 73, 0,
	402, 0, 0, 69, 0, 70, 0, 81, 80, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 74, 0, 0, 82, 83, 0,
	0, 0, 75, 76, 77, 78, 79, 0, 0, 0,
	272, 66, 129, 72, 130, 116, 0, 367, 127, 0,
	69, 0, 70, 0, 81, 80, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 74, 0, 0, 82, 83, 0, 0, 0, 75,
	76, 77, 78, 79, 66, 129, 72, 130, 288, 347,
	0, 127, 0, 0, 0, 0, 0, 238, 0, 128,
	0, 81, 80, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 74, 0, 0, 82, 83, 0,
	0, 287, 75, 76, 77, 78, 79, 66, 129, 72,
	130, 73, 0, 0, 127, 0, 0, 0, 0, 0,
	69, 0, 128, 0, 81, 80, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 74, 0, 0,
	82, 83, 0, 0, 0, 75, 76, 77, 78, 79,
	66, 129, 72, 130, 73, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 0, 128, 0, 81, 80, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	74, 0, 0, 82, 83, 0, 0, 0, 75, 76,
	77, 78, 79, 66, 277, 72, 130, 73, 0, 0,
	0, 0, 0, 164, 0, 0, 69, 0, 70, 0,
	81, 80, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 74, 0, 0, 82, 83, 0, 0,
	0, 75, 76, 77, 78, 79, 0, 0, 0, 272,
	66, 129, 72, 130, 116, 0, 0, 127, 0, 69,
	0, 70, 0, 81, 80, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	74, 0, 0, 82, 83, 0, 0, 0, 75, 76,
	77, 78, 79, 66, 129, 72, 130, 73, 0, 0,
	0, 0, 0, 0, 0, 0, 238, 0, 128, 0,
	81, 80, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 74, 0, 0, 82, 83, 0, 0,
	0, 75, 76, 77, 78, 79, 66, 441, 72, 130,
	73, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	0, 70, 0, 81, 80, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 74, 0, 0, 82,
	83, 93, 0, 0, 75, 76, 77, 78, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 93, 69, 0, 70, 0, 81, 80, 0, 0,
	102, 103, 0, 0, 0, 0, 0, 91, 92, 0,
	0, 0, 94, 95, 96, 104, 0, 0, 0, 0,
	102, 103, 90, 99, 97, 98, 101, 91, 92, 375,
	0, 0, 94, 95, 96, 104, 93, 0, 0, 0,
	0, 0, 90, 99, 97, 98, 101, 0, 0, 311,
	100, 0, 0, 0, 0, 0, 0, 89, 0, 0,
	0, 0, 0, 0, 0, 102, 103, 0, 0, 0,
	0, 0, 91, 92, 0, 0, 0, 94, 95, 96,
	104, 0, 0, 0, 0, 0, 0, 90, 99, 97,
	98, 101,
}
var RubyPact = []int{

	-29, 1552, -1000, -1000, -1000, 6, -1000, -1000, -1000, 3182,
	-1000, -1000, -1000, 136, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	19, 14, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	372, 261, 261, 1248, 230, -14, 74, 63, 206, 156,
	2592, 2592, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	3038, 2592, 3038, 3038, 283, 278, -1000, -1000, -1000, 2905,
	-1000, 28, 165, -1000, 30, 2592, 2592, 3038, 3038, 3038,
	25, 337, -1000, -1000, -1000, -1000, -1000, 2592, 2592, 3038,
	379, 3038, 3038, -1000, 3038, 2592, 3038, 3038, 3038, 3038,
	2592, 3038, -1000, -1000, 3038, 2592, 3038, 3038, 2592, 2592,
	376, 115, 129, 90, -1000, -1000, 2905, 30, 29, 2905,
	3038, 3038, 373, 241, 642, -1000, 10, -24, -24, 2995,
	162, 0, -1000, -1000, 2905, 2592, 2592, 3038, 2592, 2592,
	2592, 2592, 2592, 2592, 2592, 369, 366, 365, 271, 160,
	2494, 144, 642, 229, 642, 144, 2592, 2592, 2592, 2592,
	42, 37, 769, -1000, 3038, 2948, 265, 2905, 2543, -1000,
	-24, 271, 271, 642, 642, 642, -1000, -1000, 310, -1000,
	-1000, 271, 271, 642, 2819, 642, 642, 582, 642, 271,
	642, 642, 642, 642, 271, 601, 582, 582, 642, 271,
	642, 35, 242, 271, 271, 30, -1000, 345, 232, 64,
	-1000, 83, 330, 326, -1000, 2396, 261, 2334, 352, 769,
	-1000, -1000, -1000, 3137, -36, 33, 287, -1000, -1000, 315,
	-1000, -1000, 2862, 2272, -1000, 320, -1000, 2210, 2905, 319,
	271, 271, 222, 271, 271, 271, 271, 271, 271, 271,
	-1000, 224, -1, -16, 1503, -1000, -1000, -1000, -1000, 271,
	240, 3038, -1000, 77, 271, 271, 271, 271, -1000, -1000,
	-1000, 642, -1000, 202, 195, 39, 651, 2776, -1000, 304,
	271, -1000, -1000, -1000, -1000, 29, 30, 2592, 2905, 642,
	3038, 642, 642, -1000, 2862, 76, -1000, 1740, 129, 64,
	249, 3038, -1000, -1000, 1678, -1000, -1000, -1000, 30, -1000,
	2729, 122, -1000, -1000, 17, 642, -1000, -1000, 2161, 52,
	-1000, 2099, -1000, 3117, -1000, 71, 3038, 3038, -1000, 351,
	2592, -1000, 1441, 2050, -1000, -1000, 253, 642, 1988, 297,
	3038, 2682, -26, -1000, -35, -1000, 2592, 3038, -1000, -1000,
	271, 173, 642, 2592, -1000, 239, -1000, -1000, -1000, -1000,
	642, -1000, 220, 1926, -1000, 2635, 642, 317, 2592, 301,
	296, -1000, -1000, 286, -1000, 61, 2592, 187, 45, -1000,
	2592, -1000, 271, 2494, -1000, 344, -1000, 2494, 353, -1000,
	-1000, -1000, 271, -1000, 2592, 2592, -1000, -1000, 3038, 144,
	769, -1000, 3038, -1000, 582, -1000, 180, 271, 642, -1000,
	271, -1000, -1000, 1616, -1000, -1000, 3081, -1000, 271, -18,
	-1000, -1000, 271, 168, -1000, 271, 2494, 2494, -1000, 2494,
	279, 178, 138, 1864, 144, 2494, 642, 382, 31, -1000,
	163, 709, 3038, 2494, -1000, -1000, -1000, -1000, -1000, 2494,
	39, 2592, 3038, -1000, -1000, 32, 2494, 1379, 1311, 1802,
	39, 263, 401, -1000, -1000, 336, 2592, -1000, -1000, 308,
	-1000, -1000, -1000, 39, -1000, -1000, 2592, -1000, 271, 2445,
	-1000, 39, 271, 2445, 2445, 2445,
}
var RubyPgo = []int{

	0, 0, 483, 478, 11, 8, 476, 475, 474, 879,
	473, 10, 27, 472, 471, 470, 469, 795, 466, 502,
	418, 465, 461, 457, 455, 454, 98, 453, 9, 102,
	447, 445, 13, 439, 438, 432, 431, 30, 428, 426,
	4, 425, 424, 421, 420, 417, 414, 411, 410, 409,
	408, 407, 1, 406, 5, 20, 19, 7, 403, 6,
	400, 81, 398, 17, 394, 3, 392, 12, 14, 18,
	15, 389, 388, 966,
}
var RubyR1 = []int{

	0, 58, 58, 58, 58, 58, 58, 58, 58, 58,
	58, 72, 72, 73, 73, 52, 52, 52, 52, 18,
	18, 18, 18, 18, 18, 18, 18, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 26, 26, 26, 26, 26, 26, 26, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 37, 14, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 55, 55, 55, 55, 65,
	65, 63, 63, 63, 63, 63, 63, 63, 69, 69,
	69, 69, 69, 67, 67, 67, 21, 21, 21, 21,
	21, 21, 59, 59, 70, 70, 70, 28, 28, 28,
	28, 27, 27, 30, 32, 32, 71, 71, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 68, 68, 31,
	31, 31, 31, 31, 31, 9, 9, 29, 29, 19,
	19, 41, 41, 42, 43, 44, 45, 46, 47, 48,
	49, 50, 51, 2, 6, 7, 7, 3, 3, 60,
	60, 60, 60, 66, 66, 66, 5, 5, 5, 5,
	56, 64, 64, 64, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 57, 57, 57, 57, 53,
	53, 53, 8, 16, 11, 11, 11, 62, 62, 54,
	54, 22, 22, 23, 23, 25, 25, 25, 24, 24,
	24, 24, 12, 38, 61, 61, 61, 61, 61, 39,
	39, 39, 39, 39, 40, 40, 40, 40, 36, 35,
	10, 34, 34, 33, 33, 4,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 1, 0, 2, 0, 2, 2, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 3, 2, 4, 5, 1,
	4, 4, 2, 3, 3, 4, 4, 5, 3, 4,
	5, 2, 3, 3, 3, 4, 4, 4, 4, 4,
	4, 4, 6, 6, 6, 3, 7, 1, 5, 1,
	3, 0, 1, 1, 2, 4, 4, 5, 1, 1,
	4, 2, 5, 1, 3, 3, 5, 6, 7, 8,
	5, 6, 1, 3, 0, 1, 3, 1, 2, 3,
	2, 4, 6, 4, 1, 3, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 9, 6, 3,
	3, 3, 3, 3, 3, 2, 2, 2, 2, 3,
	3, 3, 3, 2, 2, 2, 2, 3, 3, 3,
	3, 3, 3, 1, 1, 3, 3, 5, 5, 0,
	1, 3, 7, 3, 7, 8, 3, 4, 4, 3,
	3, 0, 1, 3, 4, 5, 3, 3, 3, 3,
	3, 5, 6, 5, 3, 4, 3, 3, 2, 0,
	2, 2, 3, 4, 2, 3, 5, 0, 2, 1,
	2, 2, 1, 2, 1, 1, 3, 3, 0, 1,
	3, 3, 5, 5, 0, 2, 2, 2, 2, 5,
	6, 5, 6, 5, 4, 3, 3, 2, 4, 4,
	2, 5, 7, 4, 5, 3,
}
var RubyChk = []int{

	-1000, -58, 51, 52, 68, -1, 51, 52, 68, -17,
	-21, -27, -30, -15, -31, -13, -16, -26, -22, -38,
	-34, -23, -24, -25, -37, -4, -18, -7, -3, -32,
	-20, -8, -10, -42, -43, -44, -45, -14, -36, -35,
	13, 19, 20, 6, 8, -29, -19, -9, -41, -68,
	15, 18, 24, -46, -47, -48, -49, -50, -51, -12,
	32, 22, 36, 31, 27, 28, 5, -2, -6, 61,
	63, -71, 7, 9, 35, 43, 44, 45, 46, 47,
	66, 65, 38, 39, 52, 51, 68, 15, 18, 25,
	55, 40, 41, 4, 45, 46, 47, 57, 58, 56,
	18, 59, 33, 34, 48, 18, 40, 61, 15, 18,
	55, 6, 4, -32, 8, -32, 9, -55, -69, 61,
	42, 49, 11, -67, -17, -5, -20, 12, 63, 6,
	8, -29, -19, -9, 9, 42, 49, 61, 42, 49,
	42, 49, 42, 49, 42, 11, 42, 11, -1, -1,
	-52, -65, -17, -1, -17, -65, 15, 18, 15, 18,
	-65, -63, -17, -26, 58, -73, 54, 9, -53, -5,
	63, -1, -1, -17, -17, -17, 6, 8, 66, 6,
	8, -1, -1, -17, 6, -17, -17, -73, -17, -1,
	-17, -17, -17, -17, -1, -17, -73, -73, -17, -1,
	-17, -67, -17, -1, -1, 6, -59, 55, -70, 9,
	-28, 6, 47, 58, -59, -52, 40, -52, -63, -17,
	-5, 11, -5, -17, -4, -67, -17, -37, -12, -17,
	-12, 6, 11, -52, -56, 56, -56, -52, 61, -63,
	-1, -1, -17, -1, -1, -1, -1, -1, -1, -1,
	6, -68, 6, 6, -52, 51, 52, 51, 52, -1,
	-62, 11, 51, -73, -1, -1, -1, -1, 62, 11,
	62, -17, 51, -60, -66, -73, -17, 6, 8, -63,
	-1, 52, 10, 6, 8, -69, -55, 42, 9, -17,
	53, -17, -17, 62, 11, 62, -5, -52, 6, 11,
	-70, 42, 6, 6, -52, 14, -32, 14, 10, 11,
	-73, 62, 62, 62, -73, -17, -5, 14, -52, -64,
	6, -52, 64, -17, 10, 62, 61, 61, 14, -57,
	17, 16, -52, -52, 14, -11, 25, -17, -61, -33,
	37, -73, -73, 11, -73, 11, 4, 53, 10, -5,
	-1, -63, -17, 42, 14, -54, -11, -59, -28, 10,
	-17, 14, -54, -52, -5, -73, -17, 58, 42, 11,
	58, 14, 56, 11, 64, 62, 42, -17, -17, 14,
	17, 16, -1, -52, 14, -57, 14, -52, 8, 14,
	51, 52, -1, -39, 15, 18, 14, 16, 37, -65,
	-17, -26, 58, 64, -73, 64, -73, -1, -17, 10,
	-1, 14, -11, -52, 14, 14, 58, 6, -1, 6,
	6, 6, -1, 62, 62, -1, -52, -52, 14, -52,
	4, -1, -1, -52, -65, -52, -17, -17, 6, 14,
	-54, 6, 61, -52, 6, 51, 51, 52, 14, -52,
	-73, 4, 53, 14, 10, -17, -52, -61, -61, -61,
	-73, -1, -17, 62, 14, -40, 17, 16, 14, -40,
	14, -72, 11, -73, 11, 14, 17, 16, -1, -61,
	14, -73, -1, -61, -61, -61,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 48,
	49, 50, 51, 52, 53, 54, 55, 56, 57, 58,
	59, 60, 61, 62, 63, 64, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	0, 0, 0, 20, -2, 22, 23, 24, 0, 0,
	228, 228, 15, 41, 42, 43, 44, 45, 46, 47,
	222, 228, 0, 224, 229, 225, 19, 25, 26, 101,
	13, 0, 69, 209, 0, 228, 228, 0, 0, 0,
	0, 0, 173, 174, 5, 6, 7, 228, 228, 0,
	0, 0, 0, 13, 0, 228, 0, 0, 0, 0,
	228, 0, 13, 13, 0, 228, 0, 0, 228, 228,
	0, 124, 124, 15, -2, 15, -2, 72, 81, 101,
	0, 0, 0, 97, 108, 109, 31, 15, -2, 20,
	-2, 22, 23, 24, 101, 228, 228, 0, 228, 228,
	228, 228, 228, 228, 228, 0, 0, 0, 15, 0,
	217, 221, 99, 0, 13, 223, 228, 228, 228, 228,
	0, 0, 99, 103, 0, 13, 0, 101, 0, 250,
	15, 163, 164, 165, 166, 66, 157, 158, 0, 155,
	156, 196, 204, 65, 74, 82, 84, 0, 167, 168,
	169, 170, 171, 172, 198, 0, 0, 0, 255, 200,
	83, 0, 113, 197, 199, 78, 15, 0, 122, 124,
	125, 127, 0, 0, 15, 0, 0, 0, 0, 102,
	73, 13, 111, 99, 0, 0, 138, 139, 140, 149,
	150, 161, 13, 0, 15, 191, 15, 0, 101, 0,
	141, 151, 0, 142, 152, 143, 153, 144, 154, 145,
	162, 146, 0, 0, 0, 15, 15, 16, 17, 18,
	0, 0, 234, 0, 230, 231, 226, 227, 175, 13,
	176, 104, 14, 13, 13, 180, 0, 20, -2, 0,
	210, 211, 212, 159, 160, 75, 76, 228, -2, 85,
	0, 248, 249, 90, 0, 91, 79, 0, 124, 0,
	0, 0, 128, 130, 0, 131, 15, 133, 67, 13,
	0, 86, 88, 89, 0, 114, 115, 186, 0, 0,
	192, 0, 189, 99, 71, 87, 0, 0, 194, 0,
	228, 15, 0, 0, 213, 218, 15, 100, 0, 0,
	0, 0, 0, 13, 0, 13, -2, 0, 70, 77,
	80, 0, 232, 228, 116, 0, 219, 15, 126, 123,
	129, 120, 0, 0, 68, 0, 110, 0, 228, 0,
	0, 187, 190, 0, 188, 86, 228, 0, 0, 195,
	228, 15, 15, 208, 201, 0, 203, 214, 15, 233,
	235, 236, 237, 238, 228, 228, 251, 15, 0, 15,
	105, 106, 0, 177, 0, 178, 0, 181, 183, 95,
	94, 117, 220, 0, 121, 132, 0, 112, 92, 0,
	98, 193, 93, 0, 148, 15, 206, 207, 202, 215,
	0, 15, 0, 0, 15, 13, 107, 0, 0, 118,
	0, 20, 0, 205, 15, 234, 15, 15, 252, 13,
	253, -2, 0, 119, 96, 0, 216, 0, 0, 0,
	254, 11, 13, 147, 239, 0, 228, 234, 241, 0,
	243, 182, 12, 184, 13, 240, 228, 234, 234, 228,
	242, 185, 234, 228, 228, 228,
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
		//line parser.y:203
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:205
		{
		}
	case 3:
		//line parser.y:207
		{
		}
	case 4:
		//line parser.y:209
		{
		}
	case 5:
		//line parser.y:211
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:213
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:215
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:221
		{
		}
	case 11:
		//line parser.y:223
		{
		}
	case 12:
		//line parser.y:224
		{
		}
	case 13:
		//line parser.y:226
		{
		}
	case 14:
		//line parser.y:227
		{
		}
	case 15:
		//line parser.y:230
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 16:
		//line parser.y:232
		{
		}
	case 17:
		//line parser.y:234
		{
		}
	case 18:
		//line parser.y:236
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 64:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 65:
		//line parser.y:248
		{
			RubyVAL.genericValue = ast.RescueModifier{Statement: RubyS[Rubypt-2].genericValue, Rescue: RubyS[Rubypt-0].genericValue}
		}
	case 66:
		//line parser.y:251
		{
			RubyVAL.genericValue = ast.StarSplat{Value: RubyS[Rubypt-0].genericValue}
		}
	case 67:
		//line parser.y:254
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 68:
		//line parser.y:261
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-2].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 69:
		//line parser.y:268
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 70:
		//line parser.y:272
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 71:
		//line parser.y:279
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 72:
		//line parser.y:286
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 73:
		//line parser.y:293
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 74:
		//line parser.y:300
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 75:
		//line parser.y:307
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 76:
		//line parser.y:315
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 77:
		//line parser.y:323
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 78:
		//line parser.y:331
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 79:
		//line parser.y:339
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 80:
		//line parser.y:347
		{
			methodName := RubyS[Rubypt-2].genericValue.(ast.BareReference).Name + "="
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: methodName},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 81:
		//line parser.y:358
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 82:
		//line parser.y:365
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 83:
		//line parser.y:373
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 84:
		//line parser.y:381
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 85:
		//line parser.y:389
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 86:
		//line parser.y:399
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 87:
		//line parser.y:407
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 88:
		//line parser.y:415
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 89:
		//line parser.y:423
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 90:
		//line parser.y:431
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 91:
		//line parser.y:439
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 92:
		//line parser.y:449
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 93:
		//line parser.y:457
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 94:
		//line parser.y:465
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 95:
		//line parser.y:475
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 96:
		//line parser.y:477
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 97:
		//line parser.y:479
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 98:
		//line parser.y:481
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 99:
		//line parser.y:484
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 100:
		//line parser.y:486
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 101:
		//line parser.y:488
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 102:
		//line parser.y:490
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 103:
		//line parser.y:492
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 104:
		//line parser.y:494
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{Func: ast.BareReference{Name: "to_proc"}, Target: RubyS[Rubypt-0].genericValue})
		}
	case 105:
		//line parser.y:496
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 106:
		//line parser.y:498
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 107:
		//line parser.y:500
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{
				Func:   ast.BareReference{Name: "to_proc"},
				Target: RubyS[Rubypt-0].genericValue,
			})
		}
	case 108:
		//line parser.y:509
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 109:
		//line parser.y:511
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 110:
		//line parser.y:513
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 111:
		//line parser.y:515
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 112:
		//line parser.y:517
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 113:
		//line parser.y:520
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 114:
		//line parser.y:522
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 115:
		//line parser.y:524
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 116:
		//line parser.y:528
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 117:
		//line parser.y:536
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 118:
		//line parser.y:545
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 119:
		//line parser.y:554
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 120:
		//line parser.y:564
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 121:
		//line parser.y:572
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 122:
		//line parser.y:583
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 123:
		//line parser.y:585
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 124:
		//line parser.y:587
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 125:
		//line parser.y:589
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 126:
		//line parser.y:591
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 127:
		//line parser.y:594
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 128:
		//line parser.y:596
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 129:
		//line parser.y:598
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 130:
		//line parser.y:600
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 131:
		//line parser.y:604
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 132:
		//line parser.y:611
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 133:
		//line parser.y:621
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 134:
		//line parser.y:630
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 135:
		//line parser.y:636
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 136:
		//line parser.y:644
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 137:
		//line parser.y:648
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 138:
		//line parser.y:653
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 139:
		//line parser.y:660
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 140:
		//line parser.y:664
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 141:
		//line parser.y:668
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 142:
		//line parser.y:675
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 143:
		//line parser.y:682
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 144:
		//line parser.y:689
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 145:
		//line parser.y:696
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 146:
		//line parser.y:703
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: ast.Array{Nodes: RubyS[Rubypt-2].genericSlice},
				RHS: ast.Array{Nodes: RubyS[Rubypt-0].genericSlice},
			}
		}
	case 147:
		//line parser.y:711
		{
			RubyVAL.genericSlice = []ast.Node{
				ast.CallExpression{
					Target: RubyS[Rubypt-8].genericValue,
					Func:   ast.BareReference{Name: "[]="},
					Args:   []ast.Node{RubyS[Rubypt-6].genericValue},
				},
				ast.CallExpression{
					Target: RubyS[Rubypt-3].genericValue,
					Func:   ast.BareReference{Name: "[]="},
					Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
				},
			}
		}
	case 148:
		//line parser.y:726
		{
			tail := ast.CallExpression{Target: RubyS[Rubypt-3].genericValue, Func: ast.BareReference{Name: "[]="}, Args: []ast.Node{RubyS[Rubypt-1].genericValue}}
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, tail)
		}
	case 149:
		//line parser.y:732
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 150:
		//line parser.y:739
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 151:
		//line parser.y:743
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 152:
		//line parser.y:750
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 153:
		//line parser.y:757
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 154:
		//line parser.y:764
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 155:
		//line parser.y:772
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 156:
		//line parser.y:774
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 157:
		//line parser.y:777
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 158:
		//line parser.y:779
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 159:
		//line parser.y:782
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 160:
		//line parser.y:784
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 161:
		//line parser.y:787
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 162:
		//line parser.y:789
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 163:
		//line parser.y:791
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 164:
		//line parser.y:792
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 165:
		//line parser.y:793
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 166:
		//line parser.y:794
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 167:
		//line parser.y:797
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 168:
		//line parser.y:806
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 169:
		//line parser.y:815
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 170:
		//line parser.y:824
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 171:
		//line parser.y:833
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 172:
		//line parser.y:842
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 173:
		//line parser.y:850
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 174:
		//line parser.y:851
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 175:
		//line parser.y:853
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 176:
		//line parser.y:855
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 177:
		//line parser.y:858
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 178:
		//line parser.y:866
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 179:
		//line parser.y:874
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 180:
		//line parser.y:876
		{
		}
	case 181:
		//line parser.y:878
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 182:
		//line parser.y:885
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 183:
		//line parser.y:893
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 184:
		//line parser.y:900
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 185:
		//line parser.y:907
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 186:
		//line parser.y:915
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 187:
		//line parser.y:917
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 188:
		//line parser.y:924
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 189:
		//line parser.y:931
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 190:
		//line parser.y:934
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 191:
		//line parser.y:936
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 192:
		//line parser.y:938
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 193:
		//line parser.y:940
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 194:
		//line parser.y:943
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 195:
		//line parser.y:950
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 196:
		//line parser.y:958
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 197:
		//line parser.y:965
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 198:
		//line parser.y:972
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 199:
		//line parser.y:979
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 200:
		//line parser.y:986
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 201:
		//line parser.y:993
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 202:
		//line parser.y:1000
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 203:
		//line parser.y:1008
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 204:
		//line parser.y:1015
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 205:
		//line parser.y:1024
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 206:
		//line parser.y:1031
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 207:
		//line parser.y:1038
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 208:
		//line parser.y:1045
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 209:
		//line parser.y:1052
		{
		}
	case 210:
		//line parser.y:1053
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 211:
		//line parser.y:1054
		{
		}
	case 212:
		//line parser.y:1057
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 213:
		//line parser.y:1060
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 214:
		//line parser.y:1068
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 215:
		//line parser.y:1070
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 216:
		//line parser.y:1079
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
	case 217:
		//line parser.y:1094
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 218:
		//line parser.y:1096
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 219:
		//line parser.y:1099
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 220:
		//line parser.y:1101
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 221:
		//line parser.y:1104
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 222:
		//line parser.y:1111
		{
			RubyVAL.genericValue = ast.Yield{}
		}
	case 223:
		//line parser.y:1114
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 224:
		//line parser.y:1122
		{
			RubyVAL.genericValue = ast.Return{}
		}
	case 225:
		//line parser.y:1126
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 226:
		//line parser.y:1128
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 227:
		//line parser.y:1130
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 228:
		//line parser.y:1133
		{
		}
	case 229:
		//line parser.y:1135
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 230:
		//line parser.y:1137
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 231:
		//line parser.y:1139
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 232:
		//line parser.y:1143
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 233:
		//line parser.y:1152
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 234:
		//line parser.y:1155
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 235:
		//line parser.y:1157
		{
		}
	case 236:
		//line parser.y:1159
		{
		}
	case 237:
		//line parser.y:1161
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 238:
		//line parser.y:1163
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 239:
		//line parser.y:1166
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 240:
		//line parser.y:1173
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 241:
		//line parser.y:1181
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 242:
		//line parser.y:1188
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 243:
		//line parser.y:1196
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 244:
		//line parser.y:1204
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 245:
		//line parser.y:1211
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 246:
		//line parser.y:1218
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 247:
		//line parser.y:1225
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 248:
		//line parser.y:1233
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 249:
		//line parser.y:1236
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 250:
		//line parser.y:1238
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 251:
		//line parser.y:1241
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 252:
		//line parser.y:1243
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 253:
		//line parser.y:1246
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 254:
		//line parser.y:1248
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 255:
		//line parser.y:1250
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
