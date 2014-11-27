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

//line parser.y:1254

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 44,
	54, 136,
	-2, 21,
	-1, 115,
	54, 136,
	-2, 134,
	-1, 117,
	10, 101,
	11, 101,
	-2, 210,
	-1, 129,
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
	-1, 131,
	54, 136,
	-2, 21,
	-1, 280,
	54, 137,
	-2, 135,
	-1, 290,
	10, 101,
	11, 101,
	-2, 210,
	-1, 348,
	51, 13,
	-2, 229,
	-1, 453,
	51, 13,
	-2, 229,
}

const RubyNprod = 257
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 3271

var RubyAct = []int{

	261, 358, 5, 467, 161, 331, 357, 208, 212, 119,
	210, 29, 162, 17, 124, 25, 118, 49, 238, 87,
	2, 3, 88, 274, 177, 274, 178, 128, 24, 314,
	111, 59, 126, 237, 93, 274, 407, 4, 405, 139,
	444, 329, 328, 138, 93, 167, 140, 223, 128, 296,
	370, 149, 150, 114, 116, 85, 84, 93, 138, 375,
	454, 93, 154, 102, 103, 152, 296, 271, 156, 263,
	91, 92, 86, 102, 103, 213, 172, 173, 171, 440,
	91, 92, 106, 164, 179, 90, 102, 103, 182, 183,
	102, 103, 465, 91, 92, 90, 190, 91, 92, 171,
	315, 195, 426, 107, 374, 378, 200, 170, 90, 204,
	205, 206, 90, 45, 342, 425, 214, 295, 272, 327,
	270, 216, 202, 213, 274, 274, 211, 215, 274, 135,
	220, 164, 372, 355, 164, 227, 226, 242, 243, 303,
	245, 246, 247, 248, 249, 250, 251, 236, 241, 164,
	229, 222, 224, 230, 232, 340, 218, 132, 266, 267,
	268, 269, 136, 109, 214, 253, 110, 371, 87, 137,
	282, 88, 209, 143, 132, 215, 132, 132, 141, 455,
	144, 281, 164, 132, 398, 142, 399, 93, 106, 416,
	338, 132, 132, 132, 188, 287, 213, 108, 370, 211,
	338, 135, 288, 132, 447, 132, 132, 400, 132, 107,
	132, 132, 132, 132, 105, 132, 102, 103, 132, 148,
	132, 132, 302, 91, 92, 87, 87, 146, 88, 88,
	308, 132, 413, 93, 132, 132, 132, 214, 90, 390,
	298, 263, 336, 338, 132, 297, 411, 311, 215, 132,
	147, 371, 132, 338, 164, 227, 226, 87, 145, 347,
	88, 264, 102, 103, 337, 166, 345, 318, 159, 91,
	92, 160, 361, 301, 94, 95, 96, 104, 87, 132,
	132, 88, 132, 148, 90, 99, 97, 98, 101, 157,
	352, 377, 158, 448, 449, 350, 271, 197, 198, 132,
	326, 271, 132, 353, 164, 310, 311, 301, 359, 234,
	360, 132, 132, 364, 257, 258, 115, 285, 474, 286,
	224, 351, 87, 280, 166, 88, 168, 446, 482, 318,
	479, 478, 423, 384, 66, 130, 72, 131, 73, 422,
	387, 394, 477, 366, 479, 478, 421, 401, 132, 409,
	265, 180, 432, 181, 132, 419, 412, 403, 322, 414,
	430, 277, 383, 382, 74, 305, 414, 82, 83, 304,
	300, 420, 75, 76, 77, 78, 79, 132, 473, 424,
	274, 403, 255, 427, 93, 254, 381, 418, 383, 382,
	69, 252, 70, 132, 81, 80, 233, 433, 434, 113,
	207, 112, 185, 71, 132, 436, 132, 276, 321, 262,
	132, 275, 1, 102, 103, 169, 93, 132, 312, 58,
	91, 92, 442, 476, 57, 56, 132, 55, 54, 316,
	53, 30, 36, 35, 34, 90, 33, 48, 395, 101,
	19, 38, 132, 132, 414, 102, 103, 93, 39, 20,
	341, 14, 91, 92, 463, 12, 132, 132, 11, 23,
	394, 394, 394, 132, 471, 22, 343, 90, 89, 480,
	344, 346, 21, 18, 10, 127, 102, 103, 26, 484,
	16, 132, 394, 91, 92, 13, 394, 394, 394, 37,
	15, 32, 127, 31, 127, 127, 27, 68, 90, 28,
	67, 127, 101, 0, 0, 0, 367, 0, 0, 127,
	127, 127, 0, 0, 132, 0, 46, 0, 132, 0,
	132, 127, 0, 127, 127, 0, 127, 0, 127, 127,
	127, 127, 132, 127, 0, 0, 127, 0, 127, 127,
	406, 0, 408, 0, 0, 0, 0, 0, 0, 127,
	0, 0, 127, 127, 127, 0, 0, 132, 132, 0,
	133, 0, 127, 0, 0, 0, 0, 127, 132, 0,
	127, 0, 0, 0, 0, 0, 0, 133, 0, 133,
	133, 0, 0, 0, 0, 0, 133, 0, 0, 0,
	0, 0, 0, 0, 133, 133, 133, 127, 127, 0,
	127, 0, 0, 459, 460, 461, 133, 0, 133, 133,
	0, 133, 93, 133, 133, 133, 133, 127, 133, 0,
	127, 133, 0, 133, 133, 481, 0, 0, 453, 127,
	127, 0, 452, 0, 133, 485, 486, 133, 133, 133,
	487, 102, 103, 0, 0, 0, 462, 133, 91, 92,
	0, 0, 133, 0, 0, 133, 0, 102, 103, 475,
	0, 292, 93, 90, 91, 92, 127, 348, 0, 0,
	0, 483, 127, 0, 0, 0, 0, 0, 0, 90,
	0, 0, 133, 133, 0, 133, 0, 0, 0, 0,
	0, 102, 103, 0, 0, 127, 102, 103, 91, 92,
	0, 0, 133, 91, 92, 133, 0, 0, 0, 0,
	0, 127, 0, 90, 133, 133, 0, 0, 90, 0,
	0, 0, 127, 0, 127, 0, 0, 0, 127, 66,
	130, 72, 131, 117, 456, 127, 128, 0, 0, 0,
	0, 0, 0, 0, 127, 0, 0, 0, 0, 0,
	0, 133, 0, 0, 0, 0, 0, 133, 0, 74,
	127, 127, 82, 83, 0, 0, 0, 75, 76, 77,
	78, 79, 0, 0, 127, 127, 0, 0, 0, 0,
	133, 127, 0, 0, 0, 240, 0, 129, 0, 81,
	80, 0, 0, 0, 0, 0, 133, 0, 0, 127,
	0, 0, 0, 0, 0, 0, 0, 133, 0, 133,
	0, 0, 0, 133, 0, 0, 0, 0, 0, 0,
	133, 0, 0, 0, 0, 0, 0, 0, 93, 133,
	0, 0, 127, 0, 9, 0, 127, 0, 127, 0,
	0, 0, 100, 0, 0, 133, 133, 0, 0, 89,
	127, 0, 0, 0, 0, 0, 0, 102, 103, 133,
	133, 0, 0, 0, 91, 92, 133, 0, 0, 94,
	95, 96, 104, 0, 0, 127, 127, 0, 125, 90,
	99, 97, 98, 101, 133, 0, 127, 0, 0, 0,
	0, 0, 0, 0, 0, 153, 0, 155, 153, 0,
	0, 0, 0, 0, 163, 0, 0, 0, 0, 0,
	0, 0, 174, 175, 176, 0, 0, 133, 0, 47,
	0, 133, 0, 133, 184, 0, 186, 187, 0, 189,
	0, 191, 192, 193, 194, 133, 196, 0, 0, 199,
	0, 201, 203, 0, 0, 0, 0, 0, 0, 239,
	0, 0, 221, 0, 0, 225, 228, 231, 0, 0,
	133, 133, 0, 134, 0, 125, 0, 0, 0, 0,
	221, 133, 0, 244, 0, 0, 0, 0, 0, 0,
	134, 0, 134, 134, 0, 0, 0, 0, 0, 134,
	0, 0, 0, 0, 0, 0, 0, 134, 134, 134,
	273, 278, 151, 221, 0, 0, 0, 0, 0, 134,
	0, 134, 134, 0, 134, 0, 134, 134, 134, 134,
	125, 134, 0, 291, 134, 0, 134, 134, 0, 0,
	0, 0, 293, 294, 0, 0, 0, 134, 0, 0,
	134, 134, 134, 0, 0, 0, 0, 0, 0, 0,
	134, 0, 0, 0, 0, 134, 0, 0, 134, 0,
	0, 0, 0, 0, 217, 0, 219, 0, 0, 317,
	0, 0, 0, 0, 0, 325, 0, 0, 235, 0,
	0, 0, 0, 0, 0, 134, 134, 0, 134, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 339, 256,
	0, 0, 0, 0, 0, 134, 0, 0, 134, 0,
	0, 0, 0, 0, 125, 0, 0, 134, 134, 0,
	0, 0, 0, 0, 0, 221, 0, 354, 0, 0,
	0, 317, 0, 0, 0, 0, 0, 0, 362, 93,
	0, 0, 0, 0, 0, 0, 0, 368, 0, 0,
	0, 0, 0, 0, 134, 0, 0, 0, 299, 0,
	134, 0, 0, 379, 380, 0, 306, 0, 102, 103,
	0, 0, 0, 0, 0, 91, 92, 153, 402, 0,
	94, 95, 96, 134, 410, 0, 320, 0, 323, 0,
	90, 99, 97, 98, 101, 0, 0, 0, 0, 134,
	0, 0, 402, 0, 0, 0, 0, 334, 335, 0,
	134, 0, 134, 0, 0, 0, 134, 0, 0, 0,
	0, 0, 0, 134, 0, 0, 0, 0, 0, 0,
	0, 0, 134, 0, 0, 153, 0, 0, 0, 438,
	0, 439, 0, 0, 0, 0, 0, 0, 134, 134,
	0, 0, 0, 438, 0, 0, 0, 0, 365, 0,
	0, 93, 134, 134, 0, 0, 0, 0, 0, 134,
	0, 0, 0, 0, 0, 0, 0, 0, 125, 457,
	0, 0, 0, 385, 0, 0, 0, 134, 389, 464,
	102, 103, 0, 0, 0, 0, 0, 91, 92, 0,
	0, 0, 94, 95, 96, 104, 0, 0, 0, 415,
	0, 0, 90, 99, 97, 98, 101, 0, 0, 313,
	134, 0, 0, 0, 134, 0, 134, 0, 0, 0,
	0, 0, 0, 428, 429, 0, 0, 0, 134, 0,
	431, 0, 0, 0, 0, 0, 0, 0, 0, 435,
	0, 437, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 134, 134, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 134, 0, 0, 445, 0, 0,
	0, 0, 0, 256, 0, 0, 451, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 458, 0, 334, 335,
	66, 43, 72, 44, 73, 0, 0, 0, 40, 470,
	396, 469, 468, 397, 41, 42, 0, 61, 0, 52,
	0, 0, 64, 65, 0, 0, 63, 60, 0, 0,
	74, 62, 0, 82, 83, 0, 0, 0, 75, 76,
	77, 78, 79, 0, 0, 0, 392, 393, 0, 0,
	0, 0, 0, 0, 0, 0, 69, 0, 70, 0,
	81, 80, 66, 43, 72, 44, 73, 0, 0, 0,
	40, 466, 396, 469, 468, 397, 41, 42, 0, 61,
	0, 52, 0, 0, 64, 65, 0, 0, 63, 60,
	0, 0, 74, 62, 0, 82, 83, 0, 0, 0,
	75, 76, 77, 78, 79, 0, 0, 0, 392, 393,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 0,
	70, 0, 81, 80, 66, 43, 72, 44, 73, 0,
	0, 0, 40, 386, 50, 333, 332, 51, 41, 42,
	0, 61, 0, 52, 0, 0, 64, 65, 0, 0,
	63, 60, 0, 0, 74, 62, 0, 82, 83, 0,
	0, 0, 75, 76, 77, 78, 79, 0, 0, 0,
	259, 260, 0, 0, 0, 0, 0, 0, 0, 0,
	69, 0, 70, 0, 81, 80, 66, 43, 72, 44,
	73, 0, 0, 0, 40, 330, 50, 333, 332, 51,
	41, 42, 0, 61, 0, 52, 0, 0, 64, 65,
	0, 0, 63, 60, 0, 0, 74, 62, 0, 82,
	83, 0, 0, 0, 75, 76, 77, 78, 79, 0,
	0, 0, 259, 260, 0, 66, 43, 72, 44, 73,
	0, 0, 69, 40, 70, 50, 81, 80, 51, 41,
	42, 0, 61, 0, 52, 0, 0, 64, 65, 0,
	0, 63, 60, 0, 0, 74, 62, 0, 82, 83,
	0, 0, 0, 75, 76, 77, 78, 79, 0, 0,
	0, 6, 7, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 0, 70, 0, 81, 80, 0, 8, 66,
	43, 72, 44, 73, 0, 0, 0, 40, 441, 50,
	0, 0, 51, 41, 42, 0, 61, 0, 52, 338,
	0, 64, 65, 0, 0, 63, 60, 0, 0, 74,
	62, 0, 82, 83, 0, 0, 0, 75, 76, 77,
	78, 79, 0, 0, 0, 259, 260, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 0, 70, 0, 81,
	80, 66, 43, 72, 44, 73, 0, 0, 0, 40,
	363, 50, 0, 0, 51, 41, 42, 0, 61, 0,
	52, 338, 0, 64, 65, 0, 0, 63, 60, 0,
	0, 74, 62, 0, 82, 83, 0, 0, 0, 75,
	76, 77, 78, 79, 0, 0, 0, 259, 260, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 0, 70,
	0, 81, 80, 66, 43, 72, 44, 73, 0, 0,
	0, 40, 356, 50, 0, 0, 51, 41, 42, 0,
	61, 0, 52, 338, 0, 64, 65, 0, 0, 63,
	60, 0, 0, 74, 62, 0, 82, 83, 0, 0,
	0, 75, 76, 77, 78, 79, 0, 0, 0, 259,
	260, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	0, 70, 0, 81, 80, 66, 43, 72, 44, 73,
	0, 0, 0, 40, 472, 396, 0, 0, 397, 41,
	42, 0, 61, 0, 52, 0, 0, 64, 65, 0,
	0, 63, 60, 0, 0, 74, 62, 0, 82, 83,
	0, 0, 0, 75, 76, 77, 78, 79, 0, 0,
	0, 392, 393, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 0, 70, 0, 81, 80, 66, 43, 72,
	44, 73, 0, 0, 0, 40, 450, 50, 0, 0,
	51, 41, 42, 0, 61, 0, 52, 0, 0, 64,
	65, 0, 0, 63, 60, 0, 0, 74, 62, 0,
	82, 83, 0, 0, 0, 75, 76, 77, 78, 79,
	0, 0, 0, 259, 260, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 0, 70, 0, 81, 80, 66,
	43, 72, 44, 73, 0, 0, 0, 40, 417, 50,
	0, 0, 51, 41, 42, 0, 61, 0, 52, 0,
	0, 64, 65, 0, 0, 63, 60, 0, 0, 74,
	62, 0, 82, 83, 0, 0, 0, 75, 76, 77,
	78, 79, 0, 0, 0, 259, 260, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 0, 70, 0, 81,
	80, 66, 43, 72, 44, 73, 0, 0, 0, 40,
	391, 396, 0, 0, 397, 41, 42, 0, 61, 0,
	52, 0, 0, 64, 65, 0, 0, 63, 60, 0,
	0, 74, 62, 0, 82, 83, 0, 0, 0, 75,
	76, 77, 78, 79, 0, 0, 0, 392, 393, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 0, 70,
	0, 81, 80, 66, 43, 72, 44, 73, 0, 0,
	0, 40, 388, 50, 0, 0, 51, 41, 42, 0,
	61, 0, 52, 0, 0, 64, 65, 0, 0, 63,
	60, 0, 0, 74, 62, 0, 82, 83, 0, 0,
	0, 75, 76, 77, 78, 79, 0, 0, 0, 259,
	260, 0, 66, 43, 72, 44, 73, 0, 0, 69,
	40, 70, 50, 81, 80, 51, 41, 42, 0, 61,
	0, 52, 0, 0, 64, 65, 0, 0, 63, 60,
	0, 0, 74, 62, 0, 82, 83, 0, 0, 0,
	75, 76, 77, 78, 79, 0, 0, 0, 259, 260,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 0,
	70, 376, 81, 80, 66, 43, 72, 44, 73, 0,
	0, 0, 40, 373, 50, 0, 0, 51, 41, 42,
	0, 61, 0, 52, 0, 0, 64, 65, 0, 0,
	63, 60, 0, 0, 74, 62, 0, 82, 83, 0,
	0, 0, 75, 76, 77, 78, 79, 0, 0, 0,
	259, 260, 0, 66, 43, 72, 44, 73, 0, 0,
	69, 40, 70, 50, 81, 80, 51, 41, 42, 0,
	61, 0, 52, 0, 0, 64, 65, 0, 0, 63,
	60, 0, 0, 74, 62, 0, 82, 83, 0, 0,
	0, 75, 76, 77, 78, 79, 0, 0, 0, 259,
	260, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	0, 70, 324, 81, 80, 66, 43, 72, 44, 73,
	0, 0, 0, 40, 319, 50, 0, 0, 51, 41,
	42, 0, 61, 0, 52, 0, 0, 64, 65, 0,
	0, 63, 60, 0, 0, 74, 62, 0, 82, 83,
	0, 0, 0, 75, 76, 77, 78, 79, 0, 0,
	0, 259, 260, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 0, 70, 0, 81, 80, 66, 43, 72,
	44, 73, 0, 0, 0, 40, 309, 50, 0, 0,
	51, 41, 42, 0, 61, 0, 52, 0, 0, 64,
	65, 0, 0, 63, 60, 0, 0, 74, 62, 0,
	82, 83, 0, 0, 0, 75, 76, 77, 78, 79,
	0, 0, 0, 259, 260, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 0, 70, 0, 81, 80, 66,
	43, 72, 44, 73, 0, 0, 0, 40, 307, 50,
	0, 0, 51, 41, 42, 0, 61, 0, 52, 0,
	0, 64, 65, 0, 0, 63, 60, 0, 0, 74,
	62, 0, 82, 83, 0, 0, 0, 75, 76, 77,
	78, 79, 0, 0, 0, 259, 260, 0, 66, 43,
	72, 44, 73, 0, 0, 69, 40, 70, 396, 81,
	80, 397, 41, 42, 0, 61, 0, 52, 0, 0,
	64, 65, 0, 0, 63, 60, 0, 0, 74, 62,
	0, 82, 83, 0, 0, 0, 75, 76, 77, 78,
	79, 0, 0, 0, 392, 393, 0, 66, 43, 72,
	44, 73, 0, 0, 69, 40, 70, 50, 81, 80,
	51, 41, 42, 0, 61, 0, 52, 0, 0, 64,
	65, 0, 0, 63, 60, 0, 0, 74, 62, 0,
	82, 83, 0, 0, 0, 75, 76, 77, 78, 79,
	0, 0, 0, 259, 260, 0, 66, 43, 72, 44,
	73, 284, 0, 69, 40, 70, 50, 81, 80, 51,
	41, 42, 0, 61, 0, 52, 0, 0, 64, 65,
	0, 0, 63, 60, 0, 0, 74, 62, 0, 82,
	83, 0, 0, 0, 75, 76, 77, 78, 79, 0,
	0, 0, 0, 283, 0, 66, 43, 72, 44, 73,
	0, 0, 69, 40, 70, 50, 81, 80, 51, 41,
	42, 0, 61, 0, 52, 0, 0, 64, 65, 0,
	0, 63, 60, 0, 0, 74, 62, 0, 82, 83,
	0, 0, 0, 75, 76, 77, 78, 79, 66, 130,
	72, 131, 117, 0, 123, 128, 0, 0, 0, 0,
	0, 69, 0, 70, 0, 81, 80, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 74, 0,
	0, 82, 83, 0, 0, 121, 75, 76, 77, 78,
	79, 0, 122, 66, 130, 72, 131, 73, 0, 0,
	0, 0, 0, 0, 120, 0, 129, 0, 81, 80,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 74, 0, 0, 82, 83, 0, 0,
	0, 75, 76, 77, 78, 79, 0, 0, 0, 274,
	66, 130, 72, 131, 73, 0, 404, 0, 0, 69,
	0, 70, 0, 81, 80, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	74, 0, 0, 82, 83, 0, 0, 0, 75, 76,
	77, 78, 79, 0, 0, 0, 274, 66, 130, 72,
	131, 117, 0, 369, 128, 0, 69, 0, 70, 0,
	81, 80, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 74, 0, 0,
	82, 83, 0, 0, 0, 75, 76, 77, 78, 79,
	66, 130, 72, 131, 290, 349, 0, 128, 0, 0,
	0, 0, 0, 240, 0, 129, 0, 81, 80, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	74, 0, 0, 82, 83, 0, 0, 289, 75, 76,
	77, 78, 79, 66, 130, 72, 131, 73, 0, 0,
	0, 0, 0, 0, 0, 0, 69, 0, 129, 0,
	81, 80, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 74, 0, 0, 82, 83, 0, 0,
	0, 75, 76, 77, 78, 79, 0, 0, 0, 274,
	66, 130, 72, 131, 73, 0, 0, 128, 0, 69,
	0, 70, 0, 81, 80, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	74, 0, 0, 82, 83, 0, 0, 0, 75, 76,
	77, 78, 79, 66, 130, 72, 131, 73, 0, 0,
	0, 0, 0, 0, 0, 0, 69, 0, 129, 0,
	81, 80, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 74, 0, 0, 82, 83, 0, 0,
	0, 75, 76, 77, 78, 79, 66, 279, 72, 131,
	73, 0, 0, 0, 0, 0, 165, 0, 0, 69,
	0, 70, 0, 81, 80, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 74, 0, 0, 82,
	83, 0, 0, 0, 75, 76, 77, 78, 79, 0,
	0, 0, 274, 66, 130, 72, 131, 117, 0, 0,
	128, 0, 69, 0, 70, 0, 81, 80, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 74, 0, 0, 82, 83, 0, 0,
	0, 75, 76, 77, 78, 79, 66, 130, 72, 131,
	73, 0, 0, 0, 0, 0, 0, 0, 0, 240,
	0, 129, 0, 81, 80, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 74, 0, 0, 82,
	83, 0, 0, 0, 75, 76, 77, 78, 79, 66,
	443, 72, 131, 73, 0, 0, 0, 0, 0, 0,
	0, 0, 69, 0, 70, 0, 81, 80, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 74,
	0, 0, 82, 83, 0, 0, 0, 75, 76, 77,
	78, 79, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 0, 70, 0, 81,
	80,
}
var RubyPact = []int{

	-31, 1630, -1000, -1000, -1000, 4, -1000, -1000, -1000, 824,
	-1000, -1000, -1000, 196, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	148, -25, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	395, 308, 308, 2713, 120, -3, 136, 131, 216, 208,
	2670, 2670, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	3161, 2670, 3161, 3161, 274, 253, -1000, -1000, -1000, 3028,
	-1000, -9, 317, -1000, 15, 2670, 2670, 3161, 3161, 3161,
	18, 345, -1000, -1000, -1000, -1000, -1000, 2670, 2670, 3161,
	396, 3161, 3161, -1000, 3161, 2670, 3161, 3161, 3161, 3161,
	2670, 3161, -1000, -1000, 3161, 2670, 3161, 3161, 2670, 2670,
	2670, 394, 117, 190, 116, -1000, -1000, 3028, 15, 36,
	3028, 3161, 3161, 390, 298, 658, -1000, 42, -23, -23,
	3118, 192, -18, -1000, -1000, 3028, 2670, 2670, 3161, 2670,
	2670, 2670, 2670, 2670, 2670, 2670, 385, 379, 376, 211,
	263, 2572, 230, 658, 210, 658, 230, 2670, 2670, 2670,
	2670, 58, 56, 1135, -1000, 3161, 3071, 315, 3028, 2621,
	-1000, -23, 211, 211, 658, 658, 658, -1000, -1000, 311,
	-1000, -1000, 211, 211, 658, 2895, 658, 658, 2938, 658,
	211, 658, 658, 658, 658, 211, 608, 2938, 2938, 658,
	211, 658, 55, 183, 211, 211, 211, 15, -1000, 364,
	296, 69, -1000, 97, 363, 359, -1000, 2474, 308, 2412,
	295, 1135, -1000, -1000, -1000, 1257, -33, 38, 443, -1000,
	-1000, 380, -1000, -1000, 2985, 2350, -1000, 352, -1000, 2288,
	3028, 290, 211, 211, 57, 211, 211, 211, 211, 211,
	211, 211, -1000, 272, -19, -20, 1581, -1000, -1000, -1000,
	-1000, 211, 228, 3161, -1000, 77, 211, 211, 211, 211,
	-1000, -1000, -1000, 658, -1000, 255, 248, -16, 663, 2852,
	-1000, 285, 211, -1000, -1000, -1000, -1000, 36, 15, 2670,
	3028, 658, 3161, 658, 658, -1000, 2985, 91, -1000, 1818,
	190, 69, 262, 3161, -1000, -1000, 1756, -1000, -1000, -1000,
	15, -1000, 2805, 156, -1000, -1000, 74, 658, -1000, -1000,
	2239, 48, -1000, 2177, -1000, 229, -1000, 63, 3161, 3161,
	-1000, 372, 2670, -1000, 1519, 2128, -1000, -1000, 231, 658,
	2066, 170, 3161, 2758, -26, -1000, -28, -1000, 2670, 3161,
	-1000, -1000, 211, 236, 658, 2670, -1000, 218, -1000, -1000,
	-1000, -1000, 658, -1000, 175, 2004, -1000, 329, 658, 349,
	2670, 340, 333, -1000, -1000, 326, -1000, 8, 2670, 53,
	40, -1000, 2670, -1000, 211, 2572, -1000, 346, -1000, 2572,
	348, -1000, -1000, -1000, 211, -1000, 2670, 2670, -1000, -1000,
	3161, 230, 1135, -1000, 3161, -1000, 2938, -1000, 73, 211,
	658, -1000, 211, -1000, -1000, 1694, -1000, -1000, 3204, -1000,
	211, -21, -1000, -1000, 211, 240, -1000, 211, 2572, 2572,
	-1000, 2572, 321, 153, 242, 1942, 230, 2572, 658, 624,
	7, -1000, 165, 724, 3161, 2572, -1000, -1000, -1000, -1000,
	-1000, 2572, -16, 2670, 3161, -1000, -1000, 30, 2572, 1457,
	1395, 1880, -16, 307, 412, -1000, -1000, 328, 2670, -1000,
	-1000, 314, -1000, -1000, -1000, -16, -1000, -1000, 2670, -1000,
	211, 2523, -1000, -16, 211, 2523, 2523, 2523,
}
var RubyPgo = []int{

	0, 0, 500, 499, 15, 32, 497, 496, 493, 919,
	491, 1, 31, 490, 489, 485, 480, 834, 478, 516,
	431, 474, 473, 472, 465, 459, 13, 458, 8, 113,
	455, 451, 11, 450, 449, 448, 441, 28, 440, 438,
	3, 437, 436, 434, 433, 432, 430, 428, 427, 425,
	424, 419, 949, 415, 6, 16, 18, 5, 412, 7,
	411, 155, 409, 12, 408, 4, 407, 14, 17, 9,
	10, 403, 378, 194,
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
	31, 31, 31, 31, 31, 31, 9, 9, 29, 29,
	19, 19, 41, 41, 42, 43, 44, 45, 46, 47,
	48, 49, 50, 51, 2, 6, 7, 7, 3, 3,
	60, 60, 60, 60, 66, 66, 66, 5, 5, 5,
	5, 56, 64, 64, 64, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 57, 57, 57, 57,
	53, 53, 53, 8, 16, 11, 11, 11, 62, 62,
	54, 54, 22, 22, 23, 23, 25, 25, 25, 24,
	24, 24, 24, 12, 38, 61, 61, 61, 61, 61,
	39, 39, 39, 39, 39, 40, 40, 40, 40, 36,
	35, 10, 34, 34, 33, 33, 4,
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
	3, 3, 3, 3, 3, 3, 2, 2, 2, 2,
	3, 3, 3, 3, 2, 2, 2, 2, 3, 3,
	3, 3, 3, 3, 1, 1, 3, 3, 5, 5,
	0, 1, 3, 7, 3, 7, 8, 3, 4, 4,
	3, 3, 0, 1, 3, 4, 5, 3, 3, 3,
	3, 3, 5, 6, 5, 3, 4, 3, 3, 2,
	0, 2, 2, 3, 4, 2, 3, 5, 0, 2,
	1, 2, 2, 1, 2, 1, 1, 3, 3, 0,
	1, 3, 3, 5, 5, 0, 2, 2, 2, 2,
	5, 6, 5, 6, 5, 4, 3, 3, 2, 4,
	4, 2, 5, 7, 4, 5, 3,
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
	18, 59, 33, 34, 48, 18, 40, 61, 49, 15,
	18, 55, 6, 4, -32, 8, -32, 9, -55, -69,
	61, 42, 49, 11, -67, -17, -5, -20, 12, 63,
	6, 8, -29, -19, -9, 9, 42, 49, 61, 42,
	49, 42, 49, 42, 49, 42, 11, 42, 11, -1,
	-1, -52, -65, -17, -1, -17, -65, 15, 18, 15,
	18, -65, -63, -17, -26, 58, -73, 54, 9, -53,
	-5, 63, -1, -1, -17, -17, -17, 6, 8, 66,
	6, 8, -1, -1, -17, 6, -17, -17, -73, -17,
	-1, -17, -17, -17, -17, -1, -17, -73, -73, -17,
	-1, -17, -67, -17, -1, -1, -1, 6, -59, 55,
	-70, 9, -28, 6, 47, 58, -59, -52, 40, -52,
	-63, -17, -5, 11, -5, -17, -4, -67, -17, -37,
	-12, -17, -12, 6, 11, -52, -56, 56, -56, -52,
	61, -63, -1, -1, -17, -1, -1, -1, -1, -1,
	-1, -1, 6, -68, 6, 6, -52, 51, 52, 51,
	52, -1, -62, 11, 51, -73, -1, -1, -1, -1,
	62, 11, 62, -17, 51, -60, -66, -73, -17, 6,
	8, -63, -1, 52, 10, 6, 8, -69, -55, 42,
	9, -17, 53, -17, -17, 62, 11, 62, -5, -52,
	6, 11, -70, 42, 6, 6, -52, 14, -32, 14,
	10, 11, -73, 62, 62, 62, -73, -17, -5, 14,
	-52, -64, 6, -52, 64, -17, 10, 62, 61, 61,
	14, -57, 17, 16, -52, -52, 14, -11, 25, -17,
	-61, -33, 37, -73, -73, 11, -73, 11, 4, 53,
	10, -5, -1, -63, -17, 42, 14, -54, -11, -59,
	-28, 10, -17, 14, -54, -52, -5, -73, -17, 58,
	42, 11, 58, 14, 56, 11, 64, 62, 42, -17,
	-17, 14, 17, 16, -1, -52, 14, -57, 14, -52,
	8, 14, 51, 52, -1, -39, 15, 18, 14, 16,
	37, -65, -17, -26, 58, 64, -73, 64, -73, -1,
	-17, 10, -1, 14, -11, -52, 14, 14, 58, 6,
	-1, 6, 6, 6, -1, 62, 62, -1, -52, -52,
	14, -52, 4, -1, -1, -52, -65, -52, -17, -17,
	6, 14, -54, 6, 61, -52, 6, 51, 51, 52,
	14, -52, -73, 4, 53, 14, 10, -17, -52, -61,
	-61, -61, -73, -1, -17, 62, 14, -40, 17, 16,
	14, -40, 14, -72, 11, -73, 11, 14, 17, 16,
	-1, -61, 14, -73, -1, -61, -61, -61,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 48,
	49, 50, 51, 52, 53, 54, 55, 56, 57, 58,
	59, 60, 61, 62, 63, 64, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	0, 0, 0, 20, -2, 22, 23, 24, 0, 0,
	229, 229, 15, 41, 42, 43, 44, 45, 46, 47,
	223, 229, 0, 225, 230, 226, 19, 25, 26, 101,
	13, 0, 69, 210, 0, 229, 229, 0, 0, 0,
	0, 0, 174, 175, 5, 6, 7, 229, 229, 0,
	0, 0, 0, 13, 0, 229, 0, 0, 0, 0,
	229, 0, 13, 13, 0, 229, 0, 0, 229, 229,
	229, 0, 124, 124, 15, -2, 15, -2, 72, 81,
	101, 0, 0, 0, 97, 108, 109, 31, 15, -2,
	20, -2, 22, 23, 24, 101, 229, 229, 0, 229,
	229, 229, 229, 229, 229, 229, 0, 0, 0, 15,
	0, 218, 222, 99, 0, 13, 224, 229, 229, 229,
	229, 0, 0, 99, 103, 0, 13, 0, 101, 0,
	251, 15, 164, 165, 166, 167, 66, 158, 159, 0,
	156, 157, 197, 205, 65, 74, 82, 84, 0, 168,
	169, 170, 171, 172, 173, 199, 0, 0, 0, 256,
	201, 83, 0, 113, 155, 198, 200, 78, 15, 0,
	122, 124, 125, 127, 0, 0, 15, 0, 0, 0,
	0, 102, 73, 13, 111, 99, 0, 0, 138, 139,
	140, 149, 150, 162, 13, 0, 15, 192, 15, 0,
	101, 0, 141, 151, 0, 142, 152, 143, 153, 144,
	154, 145, 163, 146, 0, 0, 0, 15, 15, 16,
	17, 18, 0, 0, 235, 0, 231, 232, 227, 228,
	176, 13, 177, 104, 14, 13, 13, 181, 0, 20,
	-2, 0, 211, 212, 213, 160, 161, 75, 76, 229,
	-2, 85, 0, 249, 250, 90, 0, 91, 79, 0,
	124, 0, 0, 0, 128, 130, 0, 131, 15, 133,
	67, 13, 0, 86, 88, 89, 0, 114, 115, 187,
	0, 0, 193, 0, 190, 99, 71, 87, 0, 0,
	195, 0, 229, 15, 0, 0, 214, 219, 15, 100,
	0, 0, 0, 0, 0, 13, 0, 13, -2, 0,
	70, 77, 80, 0, 233, 229, 116, 0, 220, 15,
	126, 123, 129, 120, 0, 0, 68, 0, 110, 0,
	229, 0, 0, 188, 191, 0, 189, 86, 229, 0,
	0, 196, 229, 15, 15, 209, 202, 0, 204, 215,
	15, 234, 236, 237, 238, 239, 229, 229, 252, 15,
	0, 15, 105, 106, 0, 178, 0, 179, 0, 182,
	184, 95, 94, 117, 221, 0, 121, 132, 0, 112,
	92, 0, 98, 194, 93, 0, 148, 15, 207, 208,
	203, 216, 0, 15, 0, 0, 15, 13, 107, 0,
	0, 118, 0, 20, 0, 206, 15, 235, 15, 15,
	253, 13, 254, -2, 0, 119, 96, 0, 217, 0,
	0, 0, 255, 11, 13, 147, 240, 0, 229, 235,
	242, 0, 244, 183, 12, 185, 13, 241, 229, 235,
	235, 229, 243, 186, 235, 229, 229, 229,
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
		//line parser.y:771
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 156:
		//line parser.y:774
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 157:
		//line parser.y:776
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 158:
		//line parser.y:779
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 159:
		//line parser.y:781
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 160:
		//line parser.y:784
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 161:
		//line parser.y:786
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 162:
		//line parser.y:789
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 163:
		//line parser.y:791
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 164:
		//line parser.y:793
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 165:
		//line parser.y:794
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 166:
		//line parser.y:795
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 167:
		//line parser.y:796
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 168:
		//line parser.y:799
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 169:
		//line parser.y:808
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 170:
		//line parser.y:817
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 171:
		//line parser.y:826
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 172:
		//line parser.y:835
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 173:
		//line parser.y:844
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 174:
		//line parser.y:852
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 175:
		//line parser.y:853
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 176:
		//line parser.y:855
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 177:
		//line parser.y:857
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 178:
		//line parser.y:860
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 179:
		//line parser.y:868
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 180:
		//line parser.y:876
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 181:
		//line parser.y:878
		{
		}
	case 182:
		//line parser.y:880
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 183:
		//line parser.y:887
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 184:
		//line parser.y:895
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 185:
		//line parser.y:902
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 186:
		//line parser.y:909
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 187:
		//line parser.y:917
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 188:
		//line parser.y:919
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 189:
		//line parser.y:926
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 190:
		//line parser.y:933
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 191:
		//line parser.y:936
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 192:
		//line parser.y:938
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 193:
		//line parser.y:940
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 194:
		//line parser.y:942
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 195:
		//line parser.y:945
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 196:
		//line parser.y:952
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 197:
		//line parser.y:960
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 198:
		//line parser.y:967
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 199:
		//line parser.y:974
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 200:
		//line parser.y:981
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 201:
		//line parser.y:988
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 202:
		//line parser.y:995
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 203:
		//line parser.y:1002
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 204:
		//line parser.y:1010
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 205:
		//line parser.y:1017
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 206:
		//line parser.y:1026
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 207:
		//line parser.y:1033
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 208:
		//line parser.y:1040
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 209:
		//line parser.y:1047
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 210:
		//line parser.y:1054
		{
		}
	case 211:
		//line parser.y:1055
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 212:
		//line parser.y:1056
		{
		}
	case 213:
		//line parser.y:1059
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 214:
		//line parser.y:1062
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 215:
		//line parser.y:1070
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 216:
		//line parser.y:1072
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 217:
		//line parser.y:1081
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
	case 218:
		//line parser.y:1096
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 219:
		//line parser.y:1098
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 220:
		//line parser.y:1101
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 221:
		//line parser.y:1103
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 222:
		//line parser.y:1106
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 223:
		//line parser.y:1113
		{
			RubyVAL.genericValue = ast.Yield{}
		}
	case 224:
		//line parser.y:1116
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 225:
		//line parser.y:1124
		{
			RubyVAL.genericValue = ast.Return{}
		}
	case 226:
		//line parser.y:1128
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 227:
		//line parser.y:1130
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 228:
		//line parser.y:1132
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 229:
		//line parser.y:1135
		{
		}
	case 230:
		//line parser.y:1137
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 231:
		//line parser.y:1139
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 232:
		//line parser.y:1141
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 233:
		//line parser.y:1145
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 234:
		//line parser.y:1154
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 235:
		//line parser.y:1157
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 236:
		//line parser.y:1159
		{
		}
	case 237:
		//line parser.y:1161
		{
		}
	case 238:
		//line parser.y:1163
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 239:
		//line parser.y:1165
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 240:
		//line parser.y:1168
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 241:
		//line parser.y:1175
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 242:
		//line parser.y:1183
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 243:
		//line parser.y:1190
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 244:
		//line parser.y:1198
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 245:
		//line parser.y:1206
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 246:
		//line parser.y:1213
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 247:
		//line parser.y:1220
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 248:
		//line parser.y:1227
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 249:
		//line parser.y:1235
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 250:
		//line parser.y:1238
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 251:
		//line parser.y:1240
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 252:
		//line parser.y:1243
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 253:
		//line parser.y:1245
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 254:
		//line parser.y:1248
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 255:
		//line parser.y:1250
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 256:
		//line parser.y:1252
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
