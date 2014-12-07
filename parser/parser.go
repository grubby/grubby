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

//line parser.y:1312

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 45,
	54, 141,
	-2, 21,
	-1, 116,
	54, 141,
	-2, 139,
	-1, 118,
	10, 106,
	11, 106,
	-2, 225,
	-1, 132,
	54, 141,
	-2, 21,
	-1, 294,
	51, 13,
	64, 13,
	-2, 31,
	-1, 297,
	54, 142,
	-2, 140,
	-1, 308,
	10, 106,
	11, 106,
	-2, 225,
}

const RubyNprod = 271
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 3294

var RubyAct = []int{

	277, 380, 5, 379, 216, 496, 165, 352, 212, 214,
	17, 29, 120, 125, 119, 361, 127, 50, 24, 2,
	3, 88, 25, 341, 89, 181, 266, 182, 260, 290,
	136, 60, 290, 166, 290, 255, 4, 290, 107, 346,
	237, 332, 450, 94, 112, 433, 472, 431, 350, 349,
	429, 136, 153, 154, 115, 117, 140, 86, 85, 108,
	243, 227, 129, 158, 129, 314, 314, 314, 156, 94,
	287, 160, 103, 104, 87, 279, 290, 176, 177, 92,
	93, 168, 137, 394, 138, 183, 81, 105, 81, 186,
	187, 139, 174, 397, 91, 81, 483, 194, 103, 104,
	81, 345, 199, 137, 290, 92, 93, 204, 290, 243,
	208, 209, 210, 175, 46, 175, 347, 333, 313, 392,
	91, 288, 206, 220, 171, 393, 286, 494, 468, 168,
	143, 402, 168, 401, 217, 231, 226, 228, 396, 250,
	251, 233, 253, 254, 230, 258, 259, 168, 263, 264,
	265, 249, 224, 242, 234, 236, 392, 363, 152, 133,
	248, 141, 282, 283, 284, 285, 150, 377, 142, 269,
	246, 290, 321, 290, 299, 218, 133, 222, 133, 133,
	140, 147, 106, 168, 110, 133, 219, 111, 148, 151,
	484, 279, 441, 133, 133, 133, 88, 149, 393, 89,
	368, 359, 305, 359, 306, 133, 298, 133, 133, 107,
	133, 146, 133, 133, 133, 133, 438, 133, 109, 366,
	133, 217, 133, 133, 215, 320, 217, 359, 316, 215,
	108, 172, 476, 133, 326, 152, 133, 133, 133, 238,
	94, 422, 144, 423, 88, 163, 133, 89, 164, 145,
	88, 133, 133, 89, 319, 133, 168, 336, 256, 231,
	94, 261, 218, 357, 424, 267, 240, 218, 230, 103,
	104, 436, 329, 219, 359, 213, 92, 93, 219, 414,
	358, 95, 96, 97, 133, 133, 280, 133, 116, 103,
	104, 91, 100, 98, 99, 102, 92, 93, 88, 503,
	94, 89, 161, 88, 133, 162, 89, 133, 374, 297,
	94, 91, 88, 383, 319, 89, 133, 133, 454, 168,
	372, 287, 228, 373, 382, 344, 287, 381, 386, 103,
	104, 336, 328, 329, 477, 478, 92, 93, 475, 103,
	104, 511, 375, 508, 507, 388, 92, 93, 273, 274,
	94, 91, 460, 448, 408, 133, 502, 447, 453, 133,
	133, 91, 418, 411, 506, 446, 508, 507, 348, 94,
	425, 434, 444, 340, 458, 427, 407, 406, 437, 103,
	104, 439, 405, 323, 407, 406, 92, 93, 439, 303,
	90, 304, 322, 445, 133, 184, 318, 185, 103, 104,
	427, 91, 271, 452, 94, 92, 93, 455, 315, 270,
	114, 133, 113, 211, 189, 72, 293, 94, 339, 278,
	91, 461, 462, 133, 102, 133, 292, 1, 173, 133,
	59, 464, 58, 103, 104, 57, 133, 56, 55, 54,
	92, 93, 37, 36, 470, 133, 103, 104, 35, 34,
	49, 94, 473, 92, 93, 91, 47, 419, 505, 102,
	19, 39, 40, 20, 133, 133, 310, 362, 91, 14,
	12, 11, 439, 23, 22, 21, 18, 10, 133, 133,
	103, 104, 31, 492, 26, 16, 133, 92, 93, 418,
	418, 418, 488, 489, 490, 500, 13, 38, 509, 15,
	33, 134, 91, 32, 133, 27, 69, 28, 513, 68,
	0, 418, 0, 0, 510, 418, 418, 418, 134, 0,
	134, 134, 0, 0, 514, 515, 0, 134, 0, 516,
	0, 482, 0, 0, 0, 134, 134, 134, 0, 133,
	94, 0, 0, 133, 30, 133, 0, 134, 0, 134,
	134, 0, 134, 0, 134, 134, 134, 134, 133, 134,
	103, 104, 134, 0, 134, 134, 0, 92, 93, 103,
	104, 0, 0, 0, 0, 134, 92, 93, 134, 134,
	134, 239, 91, 370, 0, 0, 133, 133, 134, 128,
	0, 91, 0, 134, 134, 0, 0, 134, 133, 0,
	257, 0, 0, 262, 0, 0, 128, 268, 128, 128,
	0, 0, 103, 104, 0, 128, 0, 0, 0, 92,
	93, 0, 0, 128, 128, 128, 134, 134, 0, 134,
	0, 0, 0, 0, 91, 128, 0, 128, 128, 0,
	128, 0, 128, 128, 128, 128, 134, 128, 0, 134,
	128, 0, 128, 128, 0, 0, 0, 0, 134, 134,
	0, 0, 0, 128, 0, 0, 128, 128, 128, 0,
	67, 296, 73, 132, 74, 0, 128, 0, 0, 0,
	0, 128, 128, 0, 0, 128, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 134, 0, 0,
	75, 134, 134, 83, 84, 0, 0, 0, 76, 77,
	78, 79, 80, 0, 128, 294, 290, 128, 0, 0,
	0, 243, 0, 0, 0, 0, 70, 0, 71, 291,
	82, 81, 0, 0, 128, 0, 134, 128, 0, 0,
	94, 0, 0, 0, 0, 0, 128, 128, 0, 0,
	0, 0, 0, 134, 101, 0, 0, 0, 0, 0,
	0, 90, 0, 0, 0, 134, 0, 134, 0, 103,
	104, 134, 0, 0, 0, 0, 92, 93, 134, 0,
	0, 95, 96, 97, 105, 128, 0, 134, 0, 294,
	128, 91, 100, 98, 99, 102, 67, 131, 73, 132,
	118, 0, 124, 129, 0, 0, 134, 134, 0, 67,
	131, 73, 132, 118, 485, 0, 129, 0, 0, 0,
	134, 134, 0, 0, 128, 0, 75, 0, 134, 83,
	84, 0, 0, 122, 76, 77, 78, 79, 80, 75,
	123, 128, 83, 84, 0, 0, 134, 76, 77, 78,
	79, 80, 121, 128, 130, 128, 82, 81, 0, 128,
	94, 0, 9, 0, 0, 245, 128, 130, 0, 82,
	81, 0, 0, 0, 0, 128, 0, 0, 0, 0,
	0, 134, 0, 0, 0, 134, 0, 134, 0, 103,
	104, 0, 0, 0, 128, 128, 92, 93, 0, 0,
	134, 95, 96, 97, 105, 0, 0, 126, 128, 128,
	0, 91, 100, 98, 99, 102, 128, 0, 400, 0,
	0, 0, 0, 0, 157, 0, 159, 157, 134, 134,
	0, 0, 0, 167, 128, 0, 0, 0, 0, 0,
	134, 178, 179, 180, 0, 0, 0, 0, 0, 0,
	48, 0, 0, 188, 0, 190, 191, 0, 193, 0,
	195, 196, 197, 198, 0, 200, 0, 0, 203, 128,
	205, 207, 0, 128, 0, 128, 0, 0, 0, 0,
	0, 225, 0, 0, 229, 232, 235, 0, 128, 0,
	0, 0, 0, 0, 126, 135, 0, 0, 0, 225,
	247, 0, 0, 252, 0, 0, 0, 0, 0, 0,
	0, 0, 135, 0, 135, 135, 128, 128, 0, 0,
	0, 135, 0, 192, 0, 0, 0, 0, 128, 135,
	135, 135, 289, 295, 0, 225, 0, 0, 0, 0,
	0, 135, 0, 135, 135, 272, 135, 0, 135, 135,
	135, 135, 126, 135, 0, 309, 135, 0, 135, 135,
	0, 0, 0, 0, 311, 312, 0, 0, 0, 135,
	0, 0, 135, 135, 135, 0, 67, 131, 73, 132,
	74, 0, 135, 0, 0, 0, 0, 135, 135, 0,
	0, 135, 0, 0, 0, 170, 0, 0, 0, 155,
	0, 0, 0, 335, 0, 0, 75, 295, 343, 83,
	84, 0, 0, 0, 76, 77, 78, 79, 80, 0,
	135, 135, 290, 135, 0, 0, 0, 201, 202, 443,
	0, 0, 70, 0, 71, 0, 82, 81, 0, 0,
	135, 0, 360, 135, 0, 0, 0, 0, 0, 0,
	0, 0, 135, 135, 244, 0, 0, 0, 0, 126,
	0, 221, 0, 223, 67, 131, 73, 132, 74, 0,
	0, 225, 0, 376, 0, 241, 0, 335, 0, 0,
	0, 0, 0, 281, 384, 0, 0, 0, 0, 0,
	0, 135, 0, 390, 75, 135, 135, 83, 84, 302,
	0, 0, 76, 77, 78, 79, 80, 0, 0, 0,
	290, 0, 403, 404, 0, 0, 0, 428, 0, 0,
	70, 0, 71, 0, 82, 81, 157, 426, 0, 0,
	135, 0, 0, 0, 435, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 135, 0, 0,
	0, 330, 426, 0, 0, 0, 0, 0, 317, 135,
	0, 135, 0, 0, 334, 135, 324, 0, 0, 0,
	0, 0, 135, 0, 0, 0, 0, 0, 0, 0,
	0, 135, 0, 0, 0, 0, 0, 157, 338, 0,
	342, 466, 0, 467, 0, 0, 0, 0, 0, 0,
	135, 135, 0, 0, 0, 0, 466, 0, 0, 0,
	0, 364, 0, 0, 135, 135, 365, 367, 369, 355,
	356, 0, 135, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 126, 486, 0, 0, 0, 0,
	135, 0, 0, 0, 0, 0, 493, 0, 342, 0,
	0, 0, 0, 389, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 399, 0, 0, 0,
	0, 0, 387, 0, 0, 135, 0, 0, 0, 135,
	0, 135, 0, 0, 0, 0, 0, 398, 0, 0,
	430, 0, 432, 0, 135, 0, 0, 0, 0, 0,
	409, 0, 0, 0, 0, 413, 67, 131, 73, 132,
	74, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 135, 135, 0, 451, 0, 440, 0, 0,
	0, 0, 0, 0, 135, 0, 75, 0, 0, 83,
	84, 0, 0, 0, 76, 77, 78, 79, 80, 0,
	0, 0, 290, 456, 457, 0, 0, 0, 0, 391,
	459, 0, 70, 0, 71, 0, 82, 81, 0, 463,
	0, 465, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 481,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 474, 0, 0, 491, 0, 0, 0, 0, 0,
	480, 0, 0, 0, 0, 0, 0, 504, 0, 0,
	0, 487, 0, 355, 356, 0, 0, 0, 0, 512,
	67, 44, 73, 45, 74, 0, 0, 0, 41, 499,
	420, 498, 497, 421, 42, 43, 0, 62, 0, 53,
	0, 0, 65, 66, 0, 0, 64, 61, 0, 0,
	75, 63, 0, 83, 84, 0, 0, 0, 76, 77,
	78, 79, 80, 0, 0, 0, 416, 417, 0, 0,
	0, 0, 0, 0, 0, 0, 70, 0, 71, 0,
	82, 81, 67, 44, 73, 45, 74, 0, 0, 0,
	41, 495, 420, 498, 497, 421, 42, 43, 0, 62,
	0, 53, 0, 0, 65, 66, 0, 0, 64, 61,
	0, 0, 75, 63, 0, 83, 84, 0, 0, 0,
	76, 77, 78, 79, 80, 0, 0, 0, 416, 417,
	0, 0, 0, 0, 0, 0, 0, 0, 70, 0,
	71, 0, 82, 81, 67, 44, 73, 45, 74, 0,
	0, 0, 41, 410, 51, 354, 353, 52, 42, 43,
	0, 62, 0, 53, 0, 0, 65, 66, 0, 0,
	64, 61, 0, 0, 75, 63, 0, 83, 84, 0,
	0, 0, 76, 77, 78, 79, 80, 0, 0, 0,
	275, 276, 0, 0, 0, 0, 0, 0, 0, 0,
	70, 0, 71, 0, 82, 81, 67, 44, 73, 45,
	74, 0, 0, 0, 41, 351, 51, 354, 353, 52,
	42, 43, 0, 62, 0, 53, 0, 0, 65, 66,
	0, 0, 64, 61, 0, 0, 75, 63, 0, 83,
	84, 0, 0, 0, 76, 77, 78, 79, 80, 0,
	0, 0, 275, 276, 0, 67, 44, 73, 45, 74,
	0, 0, 70, 41, 71, 51, 82, 81, 52, 42,
	43, 0, 62, 0, 53, 0, 0, 65, 66, 0,
	0, 64, 61, 0, 0, 75, 63, 0, 83, 84,
	0, 0, 0, 76, 77, 78, 79, 80, 0, 0,
	0, 6, 7, 0, 0, 0, 0, 0, 0, 0,
	0, 70, 0, 71, 0, 82, 81, 0, 8, 67,
	44, 73, 45, 74, 0, 0, 0, 41, 469, 51,
	0, 0, 52, 42, 43, 0, 62, 0, 53, 359,
	0, 65, 66, 0, 0, 64, 61, 0, 0, 75,
	63, 0, 83, 84, 0, 0, 0, 76, 77, 78,
	79, 80, 0, 0, 0, 275, 276, 0, 0, 0,
	0, 0, 0, 0, 0, 70, 0, 71, 0, 82,
	81, 67, 44, 73, 45, 74, 0, 0, 0, 41,
	385, 51, 0, 0, 52, 42, 43, 0, 62, 0,
	53, 359, 0, 65, 66, 0, 0, 64, 61, 0,
	0, 75, 63, 0, 83, 84, 0, 0, 0, 76,
	77, 78, 79, 80, 0, 0, 0, 275, 276, 0,
	0, 0, 0, 0, 0, 0, 0, 70, 0, 71,
	0, 82, 81, 67, 44, 73, 45, 74, 0, 0,
	0, 41, 378, 51, 0, 0, 52, 42, 43, 0,
	62, 0, 53, 359, 0, 65, 66, 0, 0, 64,
	61, 0, 0, 75, 63, 0, 83, 84, 0, 0,
	0, 76, 77, 78, 79, 80, 0, 0, 0, 275,
	276, 0, 0, 0, 0, 0, 0, 0, 0, 70,
	0, 71, 0, 82, 81, 67, 44, 73, 45, 74,
	0, 0, 0, 41, 501, 420, 0, 0, 421, 42,
	43, 0, 62, 0, 53, 0, 0, 65, 66, 0,
	0, 64, 61, 0, 0, 75, 63, 0, 83, 84,
	0, 0, 0, 76, 77, 78, 79, 80, 0, 0,
	0, 416, 417, 0, 0, 0, 0, 0, 0, 0,
	0, 70, 0, 71, 0, 82, 81, 67, 44, 73,
	45, 74, 0, 0, 0, 41, 479, 51, 0, 0,
	52, 42, 43, 0, 62, 0, 53, 0, 0, 65,
	66, 0, 0, 64, 61, 0, 0, 75, 63, 0,
	83, 84, 0, 0, 0, 76, 77, 78, 79, 80,
	0, 0, 0, 275, 276, 0, 67, 44, 73, 45,
	74, 0, 0, 70, 41, 71, 51, 82, 81, 52,
	42, 43, 0, 62, 0, 53, 0, 0, 65, 66,
	0, 0, 64, 61, 0, 0, 75, 63, 0, 83,
	84, 0, 0, 0, 76, 77, 78, 79, 80, 0,
	0, 0, 275, 276, 0, 0, 0, 0, 0, 0,
	0, 0, 70, 0, 71, 449, 82, 81, 67, 44,
	73, 45, 74, 0, 0, 0, 41, 442, 51, 0,
	0, 52, 42, 43, 0, 62, 0, 53, 0, 0,
	65, 66, 0, 0, 64, 61, 0, 0, 75, 63,
	0, 83, 84, 0, 0, 0, 76, 77, 78, 79,
	80, 0, 0, 0, 275, 276, 0, 0, 0, 0,
	0, 0, 0, 0, 70, 0, 71, 0, 82, 81,
	67, 44, 73, 45, 74, 0, 0, 0, 41, 415,
	420, 0, 0, 421, 42, 43, 0, 62, 0, 53,
	0, 0, 65, 66, 0, 0, 64, 61, 0, 0,
	75, 63, 0, 83, 84, 0, 0, 0, 76, 77,
	78, 79, 80, 0, 0, 0, 416, 417, 0, 0,
	0, 0, 0, 0, 0, 0, 70, 0, 71, 0,
	82, 81, 67, 44, 73, 45, 74, 0, 0, 0,
	41, 412, 51, 0, 0, 52, 42, 43, 0, 62,
	0, 53, 0, 0, 65, 66, 0, 0, 64, 61,
	0, 0, 75, 63, 0, 83, 84, 0, 0, 0,
	76, 77, 78, 79, 80, 0, 0, 0, 275, 276,
	0, 0, 0, 0, 0, 0, 0, 0, 70, 0,
	71, 0, 82, 81, 67, 44, 73, 45, 74, 0,
	0, 0, 41, 395, 51, 0, 0, 52, 42, 43,
	0, 62, 0, 53, 0, 0, 65, 66, 0, 0,
	64, 61, 0, 0, 75, 63, 0, 83, 84, 0,
	0, 0, 76, 77, 78, 79, 80, 0, 0, 0,
	275, 276, 0, 0, 0, 0, 0, 0, 0, 0,
	70, 0, 71, 0, 82, 81, 67, 44, 73, 45,
	74, 0, 0, 0, 41, 337, 51, 0, 0, 52,
	42, 43, 0, 62, 0, 53, 0, 0, 65, 66,
	0, 0, 64, 61, 0, 0, 75, 63, 0, 83,
	84, 0, 0, 0, 76, 77, 78, 79, 80, 0,
	0, 0, 275, 276, 0, 0, 0, 0, 0, 0,
	0, 0, 70, 0, 71, 0, 82, 81, 67, 44,
	73, 45, 74, 0, 0, 0, 41, 327, 51, 0,
	0, 52, 42, 43, 0, 62, 0, 53, 0, 0,
	65, 66, 0, 0, 64, 61, 0, 0, 75, 63,
	0, 83, 84, 0, 0, 0, 76, 77, 78, 79,
	80, 0, 0, 0, 275, 276, 0, 0, 0, 0,
	0, 0, 0, 0, 70, 0, 71, 0, 82, 81,
	67, 44, 73, 45, 74, 0, 0, 0, 41, 325,
	51, 0, 0, 52, 42, 43, 0, 62, 0, 53,
	0, 0, 65, 66, 0, 0, 64, 61, 0, 0,
	75, 63, 0, 83, 84, 0, 0, 0, 76, 77,
	78, 79, 80, 0, 0, 0, 275, 276, 0, 67,
	44, 73, 45, 74, 0, 0, 70, 41, 71, 420,
	82, 81, 421, 42, 43, 0, 62, 0, 53, 0,
	0, 65, 66, 0, 0, 64, 61, 0, 0, 75,
	63, 0, 83, 84, 0, 0, 0, 76, 77, 78,
	79, 80, 0, 0, 0, 416, 417, 0, 67, 44,
	73, 45, 74, 0, 0, 70, 41, 71, 51, 82,
	81, 52, 42, 43, 0, 62, 0, 53, 0, 0,
	65, 66, 0, 0, 64, 61, 0, 0, 75, 63,
	0, 83, 84, 0, 0, 0, 76, 77, 78, 79,
	80, 0, 0, 0, 275, 276, 0, 67, 44, 73,
	45, 74, 301, 0, 70, 41, 71, 51, 82, 81,
	52, 42, 43, 0, 62, 0, 53, 0, 0, 65,
	66, 0, 0, 64, 61, 0, 0, 75, 63, 0,
	83, 84, 0, 0, 0, 76, 77, 78, 79, 80,
	0, 0, 0, 0, 300, 0, 67, 44, 73, 45,
	74, 0, 0, 70, 41, 71, 51, 82, 81, 52,
	42, 43, 0, 62, 0, 53, 0, 0, 65, 66,
	0, 0, 64, 61, 0, 0, 75, 63, 0, 83,
	84, 0, 0, 0, 76, 77, 78, 79, 80, 0,
	0, 0, 290, 0, 0, 67, 44, 73, 45, 74,
	0, 0, 70, 41, 71, 51, 82, 81, 52, 42,
	43, 0, 62, 0, 53, 0, 0, 65, 66, 0,
	0, 64, 61, 0, 0, 75, 63, 0, 83, 84,
	0, 0, 0, 76, 77, 78, 79, 80, 67, 131,
	73, 132, 118, 0, 0, 129, 0, 0, 0, 0,
	0, 70, 0, 71, 0, 82, 81, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 75, 0,
	0, 83, 84, 0, 0, 0, 76, 77, 78, 79,
	80, 67, 131, 73, 132, 308, 371, 0, 129, 0,
	0, 0, 0, 0, 245, 0, 130, 0, 82, 81,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 75, 0, 0, 83, 84, 0, 0, 307, 76,
	77, 78, 79, 80, 67, 296, 73, 132, 74, 0,
	0, 0, 0, 0, 0, 0, 0, 70, 0, 130,
	0, 82, 81, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 75, 0, 0, 83, 84, 0,
	0, 0, 76, 77, 78, 79, 80, 0, 0, 0,
	290, 67, 131, 73, 132, 74, 0, 0, 0, 0,
	70, 0, 71, 291, 82, 81, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 75, 0, 0, 83, 84, 0, 0, 0, 76,
	77, 78, 79, 80, 0, 0, 0, 290, 67, 131,
	73, 132, 74, 0, 0, 129, 0, 70, 0, 71,
	0, 82, 81, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 75, 0,
	0, 83, 84, 0, 0, 0, 76, 77, 78, 79,
	80, 67, 131, 73, 132, 74, 0, 0, 0, 0,
	0, 0, 0, 0, 70, 0, 130, 0, 82, 81,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 75, 0, 0, 83, 84, 0, 0, 0, 76,
	77, 78, 79, 80, 67, 131, 73, 132, 118, 0,
	0, 129, 0, 0, 169, 0, 0, 70, 0, 71,
	0, 82, 81, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 75, 0, 0, 83, 84, 0,
	0, 0, 76, 77, 78, 79, 80, 67, 131, 73,
	132, 74, 0, 0, 0, 0, 0, 0, 0, 0,
	245, 0, 130, 0, 82, 81, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 75, 0, 0,
	83, 84, 0, 0, 0, 76, 77, 78, 79, 80,
	67, 471, 73, 132, 74, 0, 0, 0, 0, 0,
	0, 0, 0, 70, 0, 71, 0, 82, 81, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	75, 0, 0, 83, 84, 94, 0, 0, 76, 77,
	78, 79, 80, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 70, 0, 71, 0,
	82, 81, 0, 0, 103, 104, 0, 0, 0, 0,
	0, 92, 93, 0, 0, 0, 95, 96, 97, 105,
	0, 0, 0, 0, 0, 0, 91, 100, 98, 99,
	102, 0, 0, 331,
}
var RubyPact = []int{

	-32, 1760, -1000, -1000, -1000, 6, -1000, -1000, -1000, 736,
	-1000, -1000, -1000, 164, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	169, -1000, -11, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 406, 280, 280, 791, 42, 119, 200, 139, 155,
	147, 2800, 2800, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 3152, 2800, 3152, 3152, 287, 230, -1000, -1000, -1000,
	3066, -1000, 70, 222, -1000, 52, 2800, 2800, 3152, 3152,
	3152, 19, 389, -1000, -1000, -1000, -1000, -1000, 2800, 2800,
	3152, 408, 3152, 3152, -1000, 3152, 2800, 3152, 3152, 3152,
	3152, 2800, 3152, -1000, -1000, 3152, 2800, 3152, 3152, 2800,
	2800, 2800, 407, 220, 215, 137, -1000, -1000, 3066, 52,
	50, 3066, 3152, 3152, 34, 255, 536, -1000, -2, 4,
	-1000, 3109, 21, -5, -1000, -1000, 3066, 3152, 2800, 2800,
	3152, 2800, 2800, 29, 2800, 2800, 22, 2800, 2800, 2800,
	20, 403, 396, 229, 297, 2653, 180, 536, 235, 536,
	180, 2800, 2800, 2800, 2800, 64, 59, 236, -1000, 3152,
	2929, 301, 3066, 2702, -1000, -1000, 229, 229, 536, 536,
	536, -1000, -1000, 383, -1000, -1000, 229, 229, 536, 2886,
	536, 536, 2976, 536, 229, 536, 536, 536, 536, 229,
	413, 2976, 2976, 536, 229, 536, 56, 346, 229, 229,
	229, 52, -1000, 390, 243, 128, -1000, 130, 386, 377,
	-1000, 2555, 280, 2493, 322, 236, -1000, -1000, -1000, 3231,
	-21, 55, 365, -1000, -1000, 400, -1000, -1000, -1000, -1000,
	3023, 2431, -1000, 367, 665, 3066, 315, 39, -23, 54,
	229, 229, 306, 229, 229, -1000, -1000, -1000, 229, 229,
	-1000, -1000, -1000, 229, 229, 229, -1000, -1000, -1000, 224,
	-12, -13, 1711, -1000, -1000, -1000, -1000, 229, 249, 3152,
	-1000, 120, 229, 229, 229, 229, -1000, -1000, -1000, 536,
	-1000, -1000, 208, 189, -2, 579, 2843, -1000, 310, 229,
	-1000, -1000, 53, -1000, -1000, 50, 52, 2800, 3066, 536,
	3152, 536, 536, -1000, 3023, 125, -1000, 1948, 215, 128,
	303, 3152, -1000, -1000, 1886, -1000, -1000, -1000, 52, -1000,
	1401, 114, -1000, -1000, 25, 536, -1000, -1000, 2369, 82,
	-1000, -1000, 2653, 856, -1000, 91, -1000, -1000, 89, 3152,
	3152, -1000, 368, 2800, -1000, 1649, 2307, -1000, -1000, 271,
	536, 2245, 227, 3152, 1159, -14, -1000, -17, -1000, -19,
	2800, 3152, -1000, -1000, 229, 261, 536, 2800, -1000, 202,
	-1000, -1000, -1000, -1000, 536, -1000, 178, 2183, -1000, 1071,
	536, 366, 2800, 359, 351, -1000, -1000, 347, 2121, -22,
	77, -1000, 2800, 296, 256, -1000, 2800, -1000, 229, 2653,
	-1000, 360, -1000, 2653, 348, -1000, -1000, -1000, 229, -1000,
	2800, 2800, -1000, -1000, 3152, 180, 236, -1000, 3152, -1000,
	2976, -1000, 122, -1000, 229, 536, -1000, 229, -1000, -1000,
	1824, -1000, -1000, 3195, -1000, 229, -15, -1000, -1000, -1000,
	-1000, 2751, 229, 187, -1000, 229, 2653, 2653, -1000, 2653,
	332, 181, 283, 2072, 180, 2653, 536, 527, 43, -1000,
	176, 804, 3152, 229, 2653, -1000, -1000, -1000, -1000, -1000,
	2653, 57, 2800, 3152, -1000, -1000, 65, 2653, 1587, 1525,
	2010, 57, 288, 447, -1000, -1000, 350, 2800, -1000, -1000,
	327, -1000, -1000, -1000, 57, -1000, -1000, 2800, -1000, 229,
	2604, -1000, 57, 229, 2604, 2604, 2604,
}
var RubyPgo = []int{

	0, 0, 509, 507, 22, 16, 506, 505, 503, 950,
	500, 1, 31, 499, 497, 496, 485, 862, 484, 456,
	544, 482, 477, 476, 475, 474, 473, 10, 471, 4,
	114, 470, 469, 11, 467, 463, 462, 461, 18, 460,
	457, 5, 450, 449, 448, 443, 442, 439, 438, 437,
	435, 432, 430, 1045, 428, 3, 14, 23, 7, 427,
	8, 426, 15, 419, 33, 418, 6, 416, 13, 17,
	12, 9, 415, 356, 1023,
}
var RubyR1 = []int{

	0, 59, 59, 59, 59, 59, 59, 59, 59, 59,
	59, 73, 73, 74, 74, 53, 53, 53, 53, 18,
	18, 18, 18, 18, 18, 18, 18, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 27, 27, 27, 27, 27, 27, 27, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 38, 14, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 21,
	56, 56, 56, 56, 66, 66, 64, 64, 64, 64,
	64, 64, 64, 70, 70, 70, 70, 70, 68, 68,
	68, 22, 22, 22, 22, 22, 22, 60, 60, 71,
	71, 71, 29, 29, 29, 29, 28, 28, 31, 33,
	33, 72, 72, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 69, 69, 32, 32, 32, 32, 32, 32,
	32, 9, 9, 30, 30, 19, 19, 42, 42, 42,
	42, 42, 42, 42, 42, 42, 42, 42, 42, 43,
	44, 45, 46, 47, 48, 49, 50, 51, 52, 2,
	6, 7, 7, 3, 3, 3, 3, 61, 61, 67,
	67, 67, 5, 5, 5, 5, 57, 65, 65, 65,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 58, 58, 58, 58, 54, 54, 54, 8, 16,
	11, 11, 11, 63, 63, 55, 55, 23, 23, 24,
	24, 26, 26, 26, 25, 25, 25, 12, 39, 62,
	62, 62, 62, 62, 40, 40, 40, 40, 40, 41,
	41, 41, 41, 37, 36, 10, 35, 35, 34, 34,
	4,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 1, 0, 2, 0, 2, 2, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 2, 4, 5,
	1, 4, 4, 2, 3, 3, 4, 4, 5, 3,
	4, 5, 2, 3, 3, 3, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 6, 7, 6, 6, 4,
	3, 7, 1, 5, 1, 3, 0, 1, 1, 2,
	4, 4, 5, 1, 1, 4, 2, 5, 1, 3,
	3, 5, 6, 7, 8, 5, 6, 1, 3, 0,
	1, 3, 1, 2, 3, 2, 4, 6, 4, 1,
	3, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 9, 6, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	2, 2, 2, 3, 3, 3, 3, 3, 3, 1,
	1, 3, 3, 3, 5, 5, 5, 3, 7, 3,
	7, 8, 3, 4, 5, 5, 3, 0, 1, 3,
	4, 5, 3, 3, 3, 3, 3, 5, 6, 5,
	3, 4, 3, 3, 2, 0, 2, 2, 3, 4,
	2, 3, 5, 0, 2, 1, 2, 2, 1, 2,
	1, 1, 3, 3, 1, 3, 3, 5, 5, 0,
	2, 2, 2, 2, 5, 6, 5, 6, 5, 4,
	3, 3, 2, 4, 4, 2, 5, 7, 4, 5,
	3,
}
var RubyChk = []int{

	-1000, -59, 51, 52, 68, -1, 51, 52, 68, -17,
	-22, -28, -31, -15, -32, -13, -16, -27, -23, -39,
	-35, -24, -25, -26, -38, -4, -18, -7, -3, -33,
	-20, -21, -8, -10, -43, -44, -45, -46, -14, -37,
	-36, 13, 19, 20, 6, 8, -30, -19, -9, -42,
	-69, 15, 18, 24, -47, -48, -49, -50, -51, -52,
	-12, 32, 22, 36, 31, 27, 28, 5, -2, -6,
	61, 63, -72, 7, 9, 35, 43, 44, 45, 46,
	47, 66, 65, 38, 39, 52, 51, 68, 15, 18,
	25, 55, 40, 41, 4, 45, 46, 47, 57, 58,
	56, 18, 59, 33, 34, 48, 18, 40, 61, 49,
	15, 18, 55, 6, 4, -33, 8, -33, 9, -56,
	-70, 61, 42, 49, 11, -68, -17, -5, -20, 12,
	63, 6, 8, -30, -19, -9, 9, 61, 42, 49,
	61, 42, 49, 11, 42, 49, 11, 42, 49, 42,
	11, 42, 11, -1, -1, -53, -66, -17, -1, -17,
	-66, 15, 18, 15, 18, -66, -64, -17, -27, 58,
	-74, 54, 9, -54, -5, 63, -1, -1, -17, -17,
	-17, 6, 8, 66, 6, 8, -1, -1, -17, 6,
	-17, -17, -74, -17, -1, -17, -17, -17, -17, -1,
	-17, -74, -74, -17, -1, -17, -68, -17, -1, -1,
	-1, 6, -60, 55, -71, 9, -29, 6, 47, 58,
	-60, -53, 40, -53, -64, -17, -5, 11, -5, -17,
	-4, -68, -17, -38, -12, -17, -12, 6, -30, -19,
	11, -53, -57, 56, -74, 61, -64, -17, -4, -68,
	-1, -1, -17, -1, -1, 6, -30, -19, -1, -1,
	6, -30, -19, -1, -1, -1, 6, -30, -19, -69,
	6, 6, -53, 51, 52, 51, 52, -1, -63, 11,
	51, -74, -1, -1, -1, -1, 62, 11, 62, -17,
	51, 64, -61, -67, -20, -17, 6, 8, -64, -1,
	52, 10, -74, 6, 8, -70, -56, 42, 9, -17,
	53, -17, -17, 62, 11, 62, -5, -53, 6, 11,
	-71, 42, 6, 6, -53, 14, -33, 14, 10, 11,
	-74, 62, 62, 62, -74, -17, -5, 14, -53, -65,
	6, -57, -53, -17, 10, 62, 62, 62, 62, 61,
	61, 14, -58, 17, 16, -53, -53, 14, -11, 25,
	-17, -62, -34, 37, -74, -74, 11, -74, 11, -74,
	4, 53, 10, -5, -1, -64, -17, 42, 14, -55,
	-11, -60, -29, 10, -17, 14, -55, -53, -5, -74,
	-17, 58, 42, 11, 58, 14, 56, 11, -53, -74,
	62, 42, 42, -17, -17, 14, 17, 16, -1, -53,
	14, -58, 14, -53, 8, 14, 51, 52, -1, -40,
	15, 18, 14, 16, 37, -66, -17, -27, 58, 64,
	-74, 64, -74, 64, -1, -17, 10, -1, 14, -11,
	-53, 14, 14, 58, 6, -1, 6, 6, 6, 64,
	64, -74, -1, 62, 62, -1, -53, -53, 14, -53,
	4, -1, -1, -53, -66, -53, -17, -17, 6, 14,
	-55, 6, 61, -1, -53, 6, 51, 51, 52, 14,
	-53, -74, 4, 53, 14, 10, -17, -53, -62, -62,
	-62, -74, -1, -17, 62, 14, -41, 17, 16, 14,
	-41, 14, -73, 11, -74, 11, 14, 17, 16, -1,
	-62, 14, -74, -1, -62, -62, -62,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 49,
	50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
	60, 61, 62, 63, 64, 65, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 0, 0, 0, 20, -2, 22, 23, 24, 0,
	0, 0, 0, 15, 42, 43, 44, 45, 46, 47,
	48, 238, 0, 0, 240, 244, 241, 19, 25, 26,
	106, 13, 0, 70, 225, 0, 0, 0, 0, 0,
	0, 0, 0, 189, 190, 5, 6, 7, 0, 0,
	0, 0, 0, 0, 13, 0, 0, 0, 0, 0,
	0, 0, 0, 13, 13, 0, 0, 0, 0, 0,
	0, 0, 0, 129, 129, 15, -2, 15, -2, 73,
	82, 106, 0, 0, 0, 102, 113, 114, 31, 15,
	13, 20, -2, 22, 23, 24, 106, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 15, 0, 233, 237, 104, 0, 13,
	239, 0, 0, 0, 0, 0, 0, 104, 108, 0,
	0, 0, 106, 0, 265, 13, 179, 180, 181, 182,
	67, 163, 164, 0, 161, 162, 212, 220, 66, 75,
	83, 85, 0, 183, 184, 185, 186, 187, 188, 214,
	0, 0, 0, 270, 216, 84, 0, 118, 160, 213,
	215, 79, 15, 0, 127, 129, 130, 132, 0, 0,
	15, 0, 0, 0, 0, 107, 74, 13, 116, 104,
	0, 0, 143, 144, 145, 154, 155, 167, 168, 169,
	13, 0, 15, 207, 15, 106, 0, 118, 0, 0,
	146, 156, 0, 147, 157, 170, 171, 172, 148, 158,
	173, 174, 175, 149, 159, 150, 176, 177, 178, 151,
	0, 0, 0, 15, 15, 16, 17, 18, 0, 0,
	249, 0, 245, 246, 242, 243, 191, 13, 192, 109,
	14, 193, 13, 13, -2, 0, 20, -2, 0, 226,
	227, 228, 15, 165, 166, 76, 77, 0, -2, 99,
	0, 263, 264, 93, 0, 94, 80, 0, 129, 0,
	0, 0, 133, 135, 0, 136, 15, 138, 68, 13,
	0, 86, 89, 91, 0, 119, 120, 202, 0, 0,
	208, 15, 13, 104, 72, 87, 90, 92, 88, 0,
	0, 210, 0, 0, 15, 0, 0, 229, 234, 15,
	105, 0, 0, 0, 0, 0, 13, 0, 13, 0,
	13, 0, 71, 78, 81, 0, 247, 0, 121, 0,
	235, 15, 131, 128, 134, 125, 0, 0, 69, 0,
	115, 0, 0, 0, 0, 203, 206, 0, 0, 0,
	86, 13, 0, 0, 0, 211, 0, 15, 15, 224,
	217, 0, 219, 230, 15, 248, 250, 251, 252, 253,
	0, 0, 266, 15, 0, 15, 110, 111, 0, 194,
	0, 195, 0, 196, 197, 199, 100, 98, 122, 236,
	0, 126, 137, 0, 117, 95, 0, 103, 209, 204,
	205, 0, 97, 0, 153, 15, 222, 223, 218, 231,
	0, 15, 0, 0, 15, 13, 112, 0, 0, 123,
	0, 20, 0, 96, 221, 15, 249, 15, 15, 267,
	13, 268, 13, 0, 124, 101, 0, 232, 0, 0,
	0, 269, 11, 13, 152, 254, 0, 0, 249, 256,
	0, 258, 198, 12, 200, 13, 255, 0, 249, 249,
	262, 257, 201, 249, 260, 261, 259,
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
		//line parser.y:204
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:206
		{
		}
	case 3:
		//line parser.y:208
		{
		}
	case 4:
		//line parser.y:210
		{
		}
	case 5:
		//line parser.y:212
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:214
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:216
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:222
		{
		}
	case 11:
		//line parser.y:224
		{
		}
	case 12:
		//line parser.y:225
		{
		}
	case 13:
		//line parser.y:227
		{
		}
	case 14:
		//line parser.y:228
		{
		}
	case 15:
		//line parser.y:231
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 16:
		//line parser.y:233
		{
		}
	case 17:
		//line parser.y:235
		{
		}
	case 18:
		//line parser.y:237
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 66:
		//line parser.y:249
		{
			RubyVAL.genericValue = ast.RescueModifier{Statement: RubyS[Rubypt-2].genericValue, Rescue: RubyS[Rubypt-0].genericValue}
		}
	case 67:
		//line parser.y:252
		{
			RubyVAL.genericValue = ast.StarSplat{Value: RubyS[Rubypt-0].genericValue}
		}
	case 68:
		//line parser.y:255
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 69:
		//line parser.y:262
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-2].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 70:
		//line parser.y:269
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 71:
		//line parser.y:273
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 72:
		//line parser.y:280
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 73:
		//line parser.y:287
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 74:
		//line parser.y:294
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 75:
		//line parser.y:301
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 76:
		//line parser.y:308
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 77:
		//line parser.y:316
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 78:
		//line parser.y:324
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 79:
		//line parser.y:332
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 80:
		//line parser.y:340
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 81:
		//line parser.y:348
		{
			methodName := RubyS[Rubypt-2].genericValue.(ast.BareReference).Name + "="
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: methodName},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 82:
		//line parser.y:359
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 83:
		//line parser.y:366
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 84:
		//line parser.y:374
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 85:
		//line parser.y:382
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 86:
		//line parser.y:392
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 87:
		//line parser.y:400
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 88:
		//line parser.y:408
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 89:
		//line parser.y:416
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 90:
		//line parser.y:424
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 91:
		//line parser.y:432
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 92:
		//line parser.y:440
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 93:
		//line parser.y:448
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 94:
		//line parser.y:456
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 95:
		//line parser.y:466
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 96:
		//line parser.y:474
		{
			if RubyS[Rubypt-0].genericValue == nil {
				panic("WHAT THE EVER COMPILING FUCK")
			}
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-6].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-4].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 97:
		//line parser.y:485
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 98:
		//line parser.y:493
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 99:
		//line parser.y:503
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 100:
		//line parser.y:513
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 101:
		//line parser.y:515
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 102:
		//line parser.y:517
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 103:
		//line parser.y:519
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 104:
		//line parser.y:522
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 105:
		//line parser.y:524
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 106:
		//line parser.y:526
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 107:
		//line parser.y:528
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 108:
		//line parser.y:530
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 109:
		//line parser.y:532
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{Func: ast.BareReference{Name: "to_proc"}, Target: RubyS[Rubypt-0].genericValue})
		}
	case 110:
		//line parser.y:534
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 111:
		//line parser.y:536
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 112:
		//line parser.y:538
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{
				Func:   ast.BareReference{Name: "to_proc"},
				Target: RubyS[Rubypt-0].genericValue,
			})
		}
	case 113:
		//line parser.y:547
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 114:
		//line parser.y:549
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 115:
		//line parser.y:551
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 116:
		//line parser.y:553
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 117:
		//line parser.y:555
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 118:
		//line parser.y:558
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 119:
		//line parser.y:560
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 120:
		//line parser.y:562
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 121:
		//line parser.y:566
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 122:
		//line parser.y:574
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 123:
		//line parser.y:583
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 124:
		//line parser.y:592
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 125:
		//line parser.y:602
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 126:
		//line parser.y:610
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 127:
		//line parser.y:621
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 128:
		//line parser.y:623
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 129:
		//line parser.y:625
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 130:
		//line parser.y:627
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 131:
		//line parser.y:629
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 132:
		//line parser.y:632
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 133:
		//line parser.y:634
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 134:
		//line parser.y:636
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 135:
		//line parser.y:638
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 136:
		//line parser.y:642
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 137:
		//line parser.y:649
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 138:
		//line parser.y:659
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 139:
		//line parser.y:668
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 140:
		//line parser.y:674
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 141:
		//line parser.y:682
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 142:
		//line parser.y:686
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 143:
		//line parser.y:691
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 144:
		//line parser.y:698
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 145:
		//line parser.y:702
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 146:
		//line parser.y:706
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 147:
		//line parser.y:713
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 148:
		//line parser.y:720
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 149:
		//line parser.y:727
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 150:
		//line parser.y:734
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 151:
		//line parser.y:741
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: ast.Array{Nodes: RubyS[Rubypt-2].genericSlice},
				RHS: ast.Array{Nodes: RubyS[Rubypt-0].genericSlice},
			}
		}
	case 152:
		//line parser.y:749
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
	case 153:
		//line parser.y:764
		{
			tail := ast.CallExpression{Target: RubyS[Rubypt-3].genericValue, Func: ast.BareReference{Name: "[]="}, Args: []ast.Node{RubyS[Rubypt-1].genericValue}}
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, tail)
		}
	case 154:
		//line parser.y:770
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 155:
		//line parser.y:777
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 156:
		//line parser.y:781
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 157:
		//line parser.y:788
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 158:
		//line parser.y:795
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 159:
		//line parser.y:802
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 160:
		//line parser.y:809
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 161:
		//line parser.y:812
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 162:
		//line parser.y:814
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 163:
		//line parser.y:817
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 164:
		//line parser.y:819
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 165:
		//line parser.y:822
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 166:
		//line parser.y:824
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 167:
		//line parser.y:827
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 168:
		//line parser.y:829
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 169:
		//line parser.y:831
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 170:
		//line parser.y:834
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 171:
		//line parser.y:836
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 172:
		//line parser.y:838
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 173:
		//line parser.y:841
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 174:
		//line parser.y:843
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 175:
		//line parser.y:845
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 176:
		//line parser.y:848
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 177:
		//line parser.y:850
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 178:
		//line parser.y:852
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 179:
		//line parser.y:854
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 180:
		//line parser.y:855
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 181:
		//line parser.y:856
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 182:
		//line parser.y:857
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 183:
		//line parser.y:860
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 184:
		//line parser.y:869
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 185:
		//line parser.y:878
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 186:
		//line parser.y:887
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 187:
		//line parser.y:896
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 188:
		//line parser.y:905
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 189:
		//line parser.y:913
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 190:
		//line parser.y:914
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 191:
		//line parser.y:916
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 192:
		//line parser.y:918
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 193:
		//line parser.y:921
		{
			RubyVAL.genericValue = ast.Hash{}
		}
	case 194:
		//line parser.y:923
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 195:
		//line parser.y:931
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 196:
		//line parser.y:939
		{
			RubyVAL.genericValue = ast.Block{Body: ast.Nodes{RubyS[Rubypt-2].genericValue}}
		}
	case 197:
		//line parser.y:942
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 198:
		//line parser.y:949
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 199:
		//line parser.y:957
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 200:
		//line parser.y:964
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 201:
		//line parser.y:971
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 202:
		//line parser.y:979
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 203:
		//line parser.y:981
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 204:
		//line parser.y:988
		{
			RubyVAL.genericValue = ast.Block{Args: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 205:
		//line parser.y:992
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-2].genericSlice}
		}
	case 206:
		//line parser.y:995
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 207:
		//line parser.y:997
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 208:
		//line parser.y:999
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 209:
		//line parser.y:1001
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 210:
		//line parser.y:1004
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 211:
		//line parser.y:1011
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 212:
		//line parser.y:1019
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 213:
		//line parser.y:1026
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 214:
		//line parser.y:1033
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 215:
		//line parser.y:1040
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 216:
		//line parser.y:1047
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 217:
		//line parser.y:1054
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 218:
		//line parser.y:1061
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 219:
		//line parser.y:1069
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 220:
		//line parser.y:1076
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 221:
		//line parser.y:1085
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 222:
		//line parser.y:1092
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 223:
		//line parser.y:1099
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 224:
		//line parser.y:1106
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 225:
		//line parser.y:1113
		{
		}
	case 226:
		//line parser.y:1114
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 227:
		//line parser.y:1115
		{
		}
	case 228:
		//line parser.y:1118
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 229:
		//line parser.y:1121
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 230:
		//line parser.y:1129
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 231:
		//line parser.y:1131
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 232:
		//line parser.y:1140
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
	case 233:
		//line parser.y:1155
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 234:
		//line parser.y:1157
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 235:
		//line parser.y:1160
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 236:
		//line parser.y:1162
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 237:
		//line parser.y:1165
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 238:
		//line parser.y:1172
		{
			RubyVAL.genericValue = ast.Yield{}
		}
	case 239:
		//line parser.y:1175
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 240:
		//line parser.y:1183
		{
			RubyVAL.genericValue = ast.Return{}
		}
	case 241:
		//line parser.y:1187
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 242:
		//line parser.y:1189
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 243:
		//line parser.y:1191
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 244:
		//line parser.y:1195
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 245:
		//line parser.y:1197
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 246:
		//line parser.y:1199
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 247:
		//line parser.y:1203
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 248:
		//line parser.y:1212
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 249:
		//line parser.y:1215
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 250:
		//line parser.y:1217
		{
		}
	case 251:
		//line parser.y:1219
		{
		}
	case 252:
		//line parser.y:1221
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 253:
		//line parser.y:1223
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 254:
		//line parser.y:1226
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 255:
		//line parser.y:1233
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 256:
		//line parser.y:1241
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 257:
		//line parser.y:1248
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 258:
		//line parser.y:1256
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 259:
		//line parser.y:1264
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 260:
		//line parser.y:1271
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 261:
		//line parser.y:1278
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 262:
		//line parser.y:1285
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 263:
		//line parser.y:1293
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 264:
		//line parser.y:1296
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 265:
		//line parser.y:1298
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 266:
		//line parser.y:1301
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 267:
		//line parser.y:1303
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 268:
		//line parser.y:1306
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 269:
		//line parser.y:1308
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 270:
		//line parser.y:1310
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
