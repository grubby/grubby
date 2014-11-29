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

//line parser.y:1278

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 45,
	54, 137,
	-2, 21,
	-1, 116,
	54, 137,
	-2, 135,
	-1, 118,
	10, 102,
	11, 102,
	-2, 221,
	-1, 132,
	54, 137,
	-2, 21,
	-1, 290,
	51, 13,
	64, 13,
	-2, 31,
	-1, 293,
	54, 138,
	-2, 136,
	-1, 304,
	10, 102,
	11, 102,
	-2, 221,
	-1, 363,
	51, 13,
	-2, 240,
	-1, 472,
	51, 13,
	-2, 240,
}

const RubyNprod = 268
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 3290

var RubyAct = []int{

	273, 373, 5, 372, 164, 211, 486, 345, 213, 215,
	337, 29, 125, 120, 50, 354, 165, 2, 3, 119,
	262, 25, 88, 24, 107, 89, 286, 60, 191, 286,
	286, 94, 256, 17, 4, 180, 251, 181, 236, 442,
	129, 463, 425, 423, 286, 108, 343, 226, 129, 328,
	310, 342, 152, 153, 115, 117, 142, 421, 86, 85,
	103, 104, 139, 157, 242, 112, 155, 92, 93, 159,
	170, 310, 95, 96, 97, 87, 283, 175, 176, 94,
	81, 390, 91, 100, 98, 99, 102, 140, 286, 185,
	186, 174, 81, 242, 141, 182, 81, 193, 81, 174,
	169, 329, 198, 275, 167, 473, 139, 203, 103, 104,
	207, 208, 209, 46, 286, 92, 93, 286, 146, 216,
	219, 205, 309, 127, 387, 147, 389, 284, 216, 459,
	91, 214, 200, 201, 230, 223, 385, 484, 246, 247,
	241, 249, 250, 229, 254, 255, 232, 259, 260, 261,
	233, 235, 167, 245, 282, 167, 394, 136, 133, 243,
	217, 278, 279, 280, 281, 265, 370, 356, 317, 217,
	167, 218, 221, 295, 286, 133, 216, 133, 133, 214,
	218, 286, 88, 162, 133, 89, 163, 277, 294, 474,
	137, 106, 133, 133, 133, 110, 145, 138, 111, 173,
	352, 275, 301, 298, 133, 167, 133, 133, 302, 133,
	94, 133, 133, 133, 133, 88, 133, 217, 89, 133,
	107, 133, 133, 316, 386, 212, 361, 143, 218, 109,
	359, 90, 133, 322, 144, 133, 133, 133, 237, 103,
	104, 108, 386, 225, 227, 133, 92, 93, 151, 149,
	133, 467, 468, 133, 151, 326, 252, 230, 94, 257,
	94, 91, 472, 263, 433, 102, 229, 495, 330, 136,
	406, 94, 315, 385, 88, 352, 351, 89, 167, 150,
	148, 160, 133, 133, 161, 133, 239, 103, 104, 103,
	104, 103, 104, 171, 92, 93, 92, 93, 92, 93,
	103, 104, 133, 88, 367, 133, 89, 92, 93, 91,
	466, 91, 357, 91, 133, 133, 445, 358, 360, 362,
	374, 368, 91, 88, 379, 375, 89, 116, 414, 444,
	415, 430, 350, 299, 312, 300, 94, 493, 167, 269,
	270, 88, 352, 352, 89, 293, 94, 400, 428, 325,
	183, 416, 184, 133, 382, 410, 403, 133, 133, 276,
	114, 417, 113, 332, 426, 103, 104, 392, 376, 315,
	465, 429, 92, 93, 431, 103, 104, 501, 440, 498,
	497, 431, 92, 93, 365, 283, 437, 91, 422, 133,
	424, 419, 439, 94, 341, 443, 438, 91, 496, 446,
	498, 497, 340, 283, 311, 449, 133, 399, 398, 94,
	324, 325, 436, 452, 453, 336, 419, 319, 133, 318,
	133, 455, 103, 104, 133, 227, 366, 314, 267, 92,
	93, 133, 266, 210, 332, 188, 461, 492, 103, 104,
	133, 451, 306, 72, 91, 92, 93, 397, 381, 399,
	398, 289, 335, 274, 288, 1, 133, 133, 172, 59,
	91, 58, 57, 431, 102, 56, 55, 54, 37, 36,
	133, 133, 35, 482, 34, 49, 411, 19, 133, 410,
	410, 410, 478, 479, 480, 471, 490, 39, 499, 40,
	47, 94, 20, 355, 14, 12, 133, 11, 503, 481,
	23, 410, 363, 22, 500, 410, 410, 410, 21, 18,
	10, 31, 494, 26, 504, 505, 16, 13, 38, 506,
	103, 104, 15, 33, 502, 32, 27, 92, 93, 69,
	133, 103, 104, 28, 133, 134, 133, 68, 92, 93,
	0, 0, 91, 0, 0, 0, 0, 0, 0, 133,
	0, 0, 134, 91, 134, 134, 0, 0, 0, 0,
	0, 134, 0, 0, 0, 0, 0, 0, 0, 134,
	134, 134, 0, 0, 0, 0, 133, 133, 0, 0,
	0, 134, 0, 134, 134, 0, 134, 133, 134, 134,
	134, 134, 0, 134, 0, 0, 134, 0, 134, 134,
	0, 0, 94, 0, 0, 0, 0, 0, 0, 134,
	0, 0, 134, 134, 134, 238, 101, 0, 0, 0,
	0, 0, 134, 90, 0, 0, 0, 134, 0, 0,
	134, 103, 104, 253, 0, 0, 258, 0, 92, 93,
	264, 0, 0, 95, 96, 97, 105, 0, 0, 0,
	0, 0, 30, 91, 100, 98, 99, 102, 0, 134,
	134, 0, 134, 0, 0, 0, 67, 131, 73, 132,
	118, 0, 124, 129, 0, 0, 0, 0, 0, 134,
	0, 0, 134, 0, 0, 0, 0, 0, 0, 0,
	0, 134, 134, 0, 0, 0, 75, 128, 0, 83,
	84, 0, 0, 122, 76, 77, 78, 79, 80, 0,
	123, 0, 0, 0, 128, 0, 128, 128, 0, 0,
	0, 0, 121, 128, 130, 0, 82, 81, 0, 0,
	134, 128, 128, 128, 134, 134, 0, 0, 0, 0,
	0, 0, 0, 128, 0, 128, 128, 0, 128, 0,
	128, 128, 128, 128, 0, 128, 0, 0, 128, 0,
	128, 128, 0, 0, 0, 0, 134, 0, 0, 0,
	0, 128, 0, 0, 128, 128, 128, 0, 0, 0,
	0, 0, 0, 134, 128, 0, 0, 0, 0, 128,
	0, 0, 128, 0, 0, 134, 0, 134, 0, 0,
	0, 134, 0, 0, 0, 0, 0, 0, 134, 0,
	0, 0, 0, 0, 0, 0, 0, 134, 0, 0,
	0, 128, 290, 0, 128, 0, 67, 292, 73, 132,
	74, 0, 0, 134, 134, 0, 0, 0, 0, 0,
	0, 128, 0, 0, 128, 0, 0, 134, 134, 0,
	0, 0, 0, 128, 128, 134, 75, 0, 0, 83,
	84, 0, 0, 0, 76, 77, 78, 79, 80, 0,
	0, 0, 286, 134, 0, 0, 0, 242, 0, 0,
	0, 0, 70, 0, 71, 287, 82, 81, 0, 0,
	0, 0, 128, 0, 0, 0, 290, 128, 0, 0,
	67, 131, 73, 132, 74, 0, 0, 134, 0, 0,
	0, 134, 0, 134, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 0, 0, 134, 0, 128, 0,
	75, 0, 0, 83, 84, 0, 0, 0, 76, 77,
	78, 79, 80, 0, 0, 128, 286, 0, 0, 0,
	0, 0, 0, 134, 134, 0, 70, 128, 71, 128,
	82, 81, 0, 128, 134, 126, 0, 0, 0, 0,
	128, 0, 0, 0, 0, 0, 0, 0, 0, 128,
	0, 0, 156, 0, 158, 156, 0, 0, 0, 0,
	0, 166, 0, 0, 0, 128, 128, 0, 0, 177,
	178, 179, 0, 0, 0, 0, 0, 0, 0, 128,
	128, 187, 0, 189, 190, 0, 192, 128, 194, 195,
	196, 197, 0, 199, 0, 0, 202, 0, 204, 206,
	0, 0, 0, 0, 0, 128, 0, 0, 0, 224,
	0, 0, 228, 231, 234, 0, 0, 0, 0, 0,
	0, 0, 126, 0, 0, 0, 0, 224, 0, 0,
	248, 0, 0, 0, 0, 0, 0, 0, 0, 128,
	0, 0, 0, 128, 0, 128, 0, 0, 0, 0,
	0, 0, 48, 0, 0, 0, 0, 0, 128, 285,
	291, 0, 224, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 126,
	0, 0, 305, 0, 0, 128, 128, 0, 0, 0,
	0, 307, 308, 0, 0, 0, 128, 135, 0, 0,
	0, 0, 268, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 135, 0, 135, 135, 0, 0,
	0, 0, 0, 135, 0, 0, 0, 0, 0, 0,
	331, 135, 135, 135, 291, 339, 0, 0, 0, 0,
	0, 0, 0, 135, 0, 135, 135, 0, 135, 0,
	135, 135, 135, 135, 0, 135, 154, 0, 135, 0,
	135, 135, 0, 0, 0, 0, 353, 0, 0, 0,
	0, 135, 0, 0, 135, 135, 135, 0, 0, 0,
	0, 0, 0, 126, 135, 0, 0, 0, 0, 135,
	0, 0, 135, 0, 0, 224, 0, 369, 0, 0,
	0, 331, 0, 0, 0, 0, 0, 0, 377, 0,
	0, 0, 0, 0, 0, 0, 0, 383, 220, 0,
	222, 135, 135, 0, 135, 0, 0, 0, 0, 0,
	0, 0, 240, 395, 396, 0, 0, 0, 0, 0,
	0, 135, 0, 0, 135, 0, 0, 156, 418, 0,
	0, 0, 0, 135, 135, 427, 0, 0, 0, 0,
	67, 131, 73, 132, 118, 475, 0, 129, 0, 0,
	0, 0, 0, 418, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	75, 0, 135, 83, 84, 0, 135, 135, 76, 77,
	78, 79, 80, 0, 0, 0, 0, 156, 0, 0,
	0, 457, 0, 458, 313, 0, 244, 0, 130, 0,
	82, 81, 320, 0, 0, 0, 457, 0, 135, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 334, 135, 338, 67, 131, 73,
	132, 74, 0, 126, 476, 0, 0, 135, 0, 135,
	0, 0, 0, 135, 483, 0, 0, 0, 0, 0,
	135, 0, 348, 349, 0, 0, 0, 75, 0, 135,
	83, 84, 0, 0, 0, 76, 77, 78, 79, 80,
	0, 0, 0, 286, 0, 135, 135, 0, 0, 0,
	435, 338, 0, 70, 0, 71, 0, 82, 81, 135,
	135, 0, 0, 0, 0, 0, 0, 135, 67, 131,
	73, 132, 74, 0, 0, 380, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 135, 0, 0, 0, 0,
	391, 0, 0, 0, 0, 0, 0, 0, 75, 0,
	401, 83, 84, 0, 0, 405, 76, 77, 78, 79,
	80, 0, 0, 0, 286, 0, 0, 0, 0, 135,
	0, 420, 0, 135, 70, 135, 71, 432, 82, 81,
	0, 0, 0, 0, 0, 0, 0, 0, 135, 0,
	0, 0, 0, 0, 67, 131, 73, 132, 74, 0,
	0, 0, 447, 448, 0, 0, 0, 0, 0, 450,
	0, 0, 0, 0, 0, 135, 135, 0, 454, 0,
	456, 0, 0, 0, 75, 0, 135, 83, 84, 0,
	0, 0, 76, 77, 78, 79, 80, 0, 0, 0,
	286, 0, 0, 0, 0, 0, 0, 384, 0, 464,
	70, 0, 71, 0, 82, 81, 0, 0, 470, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 477, 0,
	348, 349, 67, 44, 73, 45, 74, 0, 0, 0,
	41, 489, 412, 488, 487, 413, 42, 43, 0, 62,
	0, 53, 0, 0, 65, 66, 0, 0, 64, 61,
	0, 0, 75, 63, 0, 83, 84, 0, 0, 0,
	76, 77, 78, 79, 80, 0, 0, 0, 408, 409,
	0, 0, 0, 0, 0, 0, 0, 0, 70, 0,
	71, 0, 82, 81, 67, 44, 73, 45, 74, 0,
	0, 0, 41, 485, 412, 488, 487, 413, 42, 43,
	0, 62, 0, 53, 0, 0, 65, 66, 0, 0,
	64, 61, 0, 0, 75, 63, 0, 83, 84, 0,
	0, 0, 76, 77, 78, 79, 80, 0, 0, 0,
	408, 409, 0, 0, 0, 0, 0, 0, 0, 0,
	70, 0, 71, 0, 82, 81, 67, 44, 73, 45,
	74, 0, 0, 0, 41, 402, 51, 347, 346, 52,
	42, 43, 0, 62, 0, 53, 0, 0, 65, 66,
	0, 0, 64, 61, 0, 0, 75, 63, 0, 83,
	84, 0, 0, 0, 76, 77, 78, 79, 80, 0,
	0, 0, 271, 272, 0, 0, 0, 0, 0, 0,
	0, 0, 70, 0, 71, 0, 82, 81, 67, 44,
	73, 45, 74, 0, 0, 0, 41, 344, 51, 347,
	346, 52, 42, 43, 0, 62, 0, 53, 0, 0,
	65, 66, 0, 0, 64, 61, 0, 0, 75, 63,
	0, 83, 84, 0, 0, 0, 76, 77, 78, 79,
	80, 0, 0, 0, 271, 272, 0, 67, 44, 73,
	45, 74, 0, 0, 70, 41, 71, 51, 82, 81,
	52, 42, 43, 0, 62, 0, 53, 0, 0, 65,
	66, 0, 0, 64, 61, 0, 0, 75, 63, 0,
	83, 84, 0, 0, 0, 76, 77, 78, 79, 80,
	0, 0, 0, 6, 7, 0, 0, 0, 0, 0,
	0, 0, 0, 70, 0, 71, 0, 82, 81, 0,
	8, 67, 44, 73, 45, 74, 0, 0, 0, 41,
	460, 51, 0, 0, 52, 42, 43, 0, 62, 0,
	53, 352, 0, 65, 66, 0, 0, 64, 61, 0,
	0, 75, 63, 0, 83, 84, 0, 0, 0, 76,
	77, 78, 79, 80, 0, 0, 0, 271, 272, 0,
	0, 0, 0, 0, 0, 0, 0, 70, 0, 71,
	0, 82, 81, 67, 44, 73, 45, 74, 0, 0,
	0, 41, 378, 51, 0, 0, 52, 42, 43, 0,
	62, 0, 53, 352, 0, 65, 66, 0, 0, 64,
	61, 0, 0, 75, 63, 0, 83, 84, 0, 0,
	0, 76, 77, 78, 79, 80, 0, 0, 0, 271,
	272, 0, 0, 0, 0, 0, 0, 0, 0, 70,
	0, 71, 0, 82, 81, 67, 44, 73, 45, 74,
	0, 0, 0, 41, 371, 51, 0, 0, 52, 42,
	43, 0, 62, 0, 53, 352, 0, 65, 66, 0,
	0, 64, 61, 0, 0, 75, 63, 0, 83, 84,
	0, 0, 0, 76, 77, 78, 79, 80, 0, 0,
	0, 271, 272, 0, 0, 0, 0, 0, 0, 0,
	0, 70, 0, 71, 0, 82, 81, 67, 44, 73,
	45, 74, 0, 0, 0, 41, 491, 412, 0, 0,
	413, 42, 43, 0, 62, 0, 53, 0, 0, 65,
	66, 0, 0, 64, 61, 0, 0, 75, 63, 0,
	83, 84, 0, 0, 0, 76, 77, 78, 79, 80,
	0, 0, 0, 408, 409, 0, 0, 0, 0, 0,
	0, 0, 0, 70, 0, 71, 0, 82, 81, 67,
	44, 73, 45, 74, 0, 0, 0, 41, 469, 51,
	0, 0, 52, 42, 43, 0, 62, 0, 53, 0,
	0, 65, 66, 0, 0, 64, 61, 0, 0, 75,
	63, 0, 83, 84, 0, 0, 0, 76, 77, 78,
	79, 80, 0, 0, 0, 271, 272, 0, 67, 44,
	73, 45, 74, 0, 0, 70, 41, 71, 51, 82,
	81, 52, 42, 43, 0, 62, 0, 53, 0, 0,
	65, 66, 0, 0, 64, 61, 0, 0, 75, 63,
	0, 83, 84, 0, 0, 0, 76, 77, 78, 79,
	80, 0, 0, 0, 271, 272, 0, 0, 0, 0,
	0, 0, 0, 0, 70, 0, 71, 441, 82, 81,
	67, 44, 73, 45, 74, 0, 0, 0, 41, 434,
	51, 0, 0, 52, 42, 43, 0, 62, 0, 53,
	0, 0, 65, 66, 0, 0, 64, 61, 0, 0,
	75, 63, 0, 83, 84, 0, 0, 0, 76, 77,
	78, 79, 80, 0, 0, 0, 271, 272, 0, 0,
	0, 0, 0, 0, 0, 0, 70, 0, 71, 0,
	82, 81, 67, 44, 73, 45, 74, 0, 0, 0,
	41, 407, 412, 0, 0, 413, 42, 43, 0, 62,
	0, 53, 0, 0, 65, 66, 0, 0, 64, 61,
	0, 0, 75, 63, 0, 83, 84, 0, 0, 0,
	76, 77, 78, 79, 80, 0, 0, 0, 408, 409,
	0, 0, 0, 0, 0, 0, 0, 0, 70, 0,
	71, 0, 82, 81, 67, 44, 73, 45, 74, 0,
	0, 0, 41, 404, 51, 0, 0, 52, 42, 43,
	0, 62, 0, 53, 0, 0, 65, 66, 0, 0,
	64, 61, 0, 0, 75, 63, 0, 83, 84, 0,
	0, 0, 76, 77, 78, 79, 80, 0, 0, 0,
	271, 272, 0, 0, 0, 0, 0, 0, 0, 0,
	70, 0, 71, 0, 82, 81, 67, 44, 73, 45,
	74, 0, 0, 0, 41, 388, 51, 0, 0, 52,
	42, 43, 0, 62, 0, 53, 0, 0, 65, 66,
	0, 0, 64, 61, 0, 0, 75, 63, 0, 83,
	84, 0, 0, 0, 76, 77, 78, 79, 80, 0,
	0, 0, 271, 272, 0, 0, 0, 0, 0, 0,
	0, 0, 70, 0, 71, 0, 82, 81, 67, 44,
	73, 45, 74, 0, 0, 0, 41, 333, 51, 0,
	0, 52, 42, 43, 0, 62, 0, 53, 0, 0,
	65, 66, 0, 0, 64, 61, 0, 0, 75, 63,
	0, 83, 84, 0, 0, 0, 76, 77, 78, 79,
	80, 0, 0, 0, 271, 272, 0, 0, 0, 0,
	0, 0, 0, 0, 70, 0, 71, 0, 82, 81,
	67, 44, 73, 45, 74, 0, 0, 0, 41, 323,
	51, 0, 0, 52, 42, 43, 0, 62, 0, 53,
	0, 0, 65, 66, 0, 0, 64, 61, 0, 0,
	75, 63, 0, 83, 84, 0, 0, 0, 76, 77,
	78, 79, 80, 0, 0, 0, 271, 272, 0, 0,
	0, 0, 0, 0, 0, 0, 70, 0, 71, 0,
	82, 81, 67, 44, 73, 45, 74, 0, 0, 0,
	41, 321, 51, 0, 0, 52, 42, 43, 0, 62,
	0, 53, 0, 0, 65, 66, 0, 0, 64, 61,
	0, 0, 75, 63, 0, 83, 84, 0, 0, 0,
	76, 77, 78, 79, 80, 0, 0, 0, 271, 272,
	0, 67, 44, 73, 45, 74, 0, 0, 70, 41,
	71, 412, 82, 81, 413, 42, 43, 0, 62, 0,
	53, 0, 0, 65, 66, 0, 0, 64, 61, 0,
	0, 75, 63, 0, 83, 84, 0, 0, 0, 76,
	77, 78, 79, 80, 0, 0, 0, 408, 409, 0,
	67, 44, 73, 45, 74, 0, 0, 70, 41, 71,
	51, 82, 81, 52, 42, 43, 0, 62, 0, 53,
	0, 0, 65, 66, 0, 0, 64, 61, 0, 0,
	75, 63, 0, 83, 84, 0, 0, 0, 76, 77,
	78, 79, 80, 0, 0, 0, 271, 272, 0, 67,
	44, 73, 45, 74, 297, 0, 70, 41, 71, 51,
	82, 81, 52, 42, 43, 0, 62, 0, 53, 0,
	0, 65, 66, 0, 0, 64, 61, 0, 0, 75,
	63, 0, 83, 84, 0, 0, 0, 76, 77, 78,
	79, 80, 0, 0, 0, 0, 296, 0, 67, 44,
	73, 45, 74, 0, 0, 70, 41, 71, 51, 82,
	81, 52, 42, 43, 0, 62, 0, 53, 0, 0,
	65, 66, 0, 0, 64, 61, 0, 0, 75, 63,
	0, 83, 84, 0, 0, 0, 76, 77, 78, 79,
	80, 67, 131, 73, 132, 118, 0, 0, 129, 0,
	0, 0, 0, 0, 70, 0, 71, 0, 82, 81,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 75, 0, 0, 83, 84, 0, 0, 0, 76,
	77, 78, 79, 80, 67, 131, 73, 132, 304, 364,
	0, 129, 0, 0, 0, 0, 0, 244, 0, 130,
	0, 82, 81, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 75, 0, 0, 83, 84, 0,
	0, 303, 76, 77, 78, 79, 80, 67, 292, 73,
	132, 74, 0, 0, 0, 0, 0, 0, 0, 0,
	70, 0, 130, 0, 82, 81, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 75, 0, 0,
	83, 84, 0, 0, 0, 76, 77, 78, 79, 80,
	0, 0, 0, 286, 67, 131, 73, 132, 74, 0,
	0, 129, 0, 70, 0, 71, 287, 82, 81, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 75, 0, 0, 83, 84, 0,
	0, 0, 76, 77, 78, 79, 80, 67, 131, 73,
	132, 74, 0, 0, 0, 0, 0, 0, 0, 0,
	70, 0, 130, 0, 82, 81, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 75, 0, 0,
	83, 84, 0, 0, 0, 76, 77, 78, 79, 80,
	67, 131, 73, 132, 118, 0, 0, 129, 0, 0,
	168, 0, 0, 70, 0, 71, 0, 82, 81, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	75, 0, 0, 83, 84, 0, 0, 0, 76, 77,
	78, 79, 80, 67, 131, 73, 132, 74, 0, 0,
	0, 0, 0, 0, 0, 0, 244, 0, 130, 0,
	82, 81, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 75, 0, 0, 83, 84, 0, 0,
	0, 76, 77, 78, 79, 80, 67, 462, 73, 132,
	74, 0, 0, 0, 0, 0, 0, 0, 0, 70,
	0, 71, 0, 82, 81, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 75, 0, 0, 83,
	84, 94, 0, 0, 76, 77, 78, 79, 80, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 94, 70, 0, 71, 0, 82, 81, 0, 0,
	103, 104, 0, 0, 0, 0, 0, 92, 93, 0,
	0, 0, 95, 96, 97, 105, 0, 0, 0, 0,
	103, 104, 91, 100, 98, 99, 102, 92, 93, 393,
	0, 0, 95, 96, 97, 105, 0, 0, 0, 0,
	0, 0, 91, 100, 98, 99, 102, 0, 0, 327,
}
var RubyPact = []int{

	-34, 1832, -1000, -1000, -1000, 7, -1000, -1000, -1000, 598,
	-1000, -1000, -1000, 173, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	180, -1000, 10, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 356, 319, 319, 661, 148, 45, 185, 76, 238,
	237, 2823, 2823, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 3128, 2823, 3128, 3128, 266, 168, -1000, -1000, -1000,
	3042, -1000, 16, 284, -1000, 28, 2823, 2823, 3128, 3128,
	3128, 29, 344, -1000, -1000, -1000, -1000, -1000, 2823, 2823,
	3128, 429, 3128, 3128, -1000, 3128, 2823, 3128, 3128, 3128,
	3128, 2823, 3128, -1000, -1000, 3128, 2823, 3128, 3128, 2823,
	2823, 2823, 427, 170, 122, 132, -1000, -1000, 3042, 28,
	36, 3042, 3128, 3128, 32, 275, 487, -1000, -16, 8,
	-1000, 3085, 260, 1, -1000, -1000, 3042, 2823, 2823, 3128,
	2823, 2823, 30, 2823, 2823, 26, 2823, 2823, 2823, 14,
	426, 422, 167, 288, 2725, 190, 487, 308, 487, 190,
	2823, 2823, 2823, 2823, 92, 65, 27, -1000, 3128, 2952,
	337, 3042, 2774, -1000, -1000, 167, 167, 487, 487, 487,
	-1000, -1000, 327, -1000, -1000, 167, 167, 487, 2909, 487,
	487, 895, 487, 167, 487, 487, 487, 487, 167, 389,
	895, 895, 487, 167, 487, 60, 342, 167, 167, 167,
	28, -1000, 421, 261, 113, -1000, 126, 413, 411, -1000,
	2627, 319, 2565, 400, 27, -1000, -1000, -1000, 3227, -13,
	39, 206, -1000, -1000, 405, -1000, -1000, -1000, -1000, 2999,
	2503, -1000, 409, 821, 3042, 392, 167, 167, 332, 167,
	167, -1000, -1000, -1000, 167, 167, -1000, -1000, -1000, 167,
	167, 167, -1000, -1000, -1000, 243, -10, -15, 1783, -1000,
	-1000, -1000, -1000, 167, 318, 3128, -1000, 130, 167, 167,
	167, 167, -1000, -1000, -1000, 487, -1000, -1000, 219, 215,
	-16, 498, 2866, -1000, 374, 167, -1000, -1000, 37, -1000,
	-1000, 36, 28, 2823, 3042, 487, 3128, 487, 487, -1000,
	2999, 124, -1000, 2020, 122, 113, 358, 3128, -1000, -1000,
	1958, -1000, -1000, -1000, 28, -1000, 1519, 231, -1000, -1000,
	66, 487, -1000, -1000, 2441, 70, -1000, -1000, 2725, 3207,
	-1000, 114, 3128, 3128, -1000, 433, 2823, -1000, 1721, 2379,
	-1000, -1000, 262, 487, 2317, 314, 3128, 1443, -7, -1000,
	-21, -1000, -22, 2823, 3128, -1000, -1000, 167, 338, 487,
	2823, -1000, 317, -1000, -1000, -1000, -1000, 487, -1000, 250,
	2255, -1000, 1372, 487, 406, 2823, 390, 386, -1000, -1000,
	372, 2193, -25, 94, 2823, 267, 254, -1000, 2823, -1000,
	167, 2725, -1000, 391, -1000, 2725, 437, -1000, -1000, -1000,
	167, -1000, 2823, 2823, -1000, -1000, 3128, 190, 27, -1000,
	3128, -1000, 895, -1000, 123, -1000, 167, 487, -1000, 167,
	-1000, -1000, 1896, -1000, -1000, 3171, -1000, 167, -20, -1000,
	-1000, -1000, -1000, 167, 213, -1000, 167, 2725, 2725, -1000,
	2725, 364, 259, 200, 2144, 190, 2725, 487, 258, 52,
	-1000, 175, 1285, 3128, 2725, -1000, -1000, -1000, -1000, -1000,
	2725, 63, 2823, 3128, -1000, -1000, 75, 2725, 1659, 1597,
	2082, 63, 326, 256, -1000, -1000, 384, 2823, -1000, -1000,
	363, -1000, -1000, -1000, 63, -1000, -1000, 2823, -1000, 167,
	2676, -1000, 63, 167, 2676, 2676, 2676,
}
var RubyPgo = []int{

	0, 0, 537, 533, 21, 123, 529, 526, 525, 1082,
	523, 1, 27, 522, 518, 517, 516, 920, 513, 490,
	652, 511, 510, 509, 508, 503, 500, 33, 497, 9,
	113, 495, 494, 11, 493, 492, 489, 487, 23, 477,
	476, 6, 475, 474, 472, 469, 468, 467, 466, 465,
	462, 461, 459, 1132, 458, 3, 19, 10, 7, 455,
	5, 454, 15, 453, 16, 452, 4, 451, 12, 14,
	13, 8, 443, 437, 28,
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
	20, 20, 20, 20, 20, 21, 56, 56, 56, 56,
	66, 66, 64, 64, 64, 64, 64, 64, 64, 70,
	70, 70, 70, 70, 68, 68, 68, 22, 22, 22,
	22, 22, 22, 60, 60, 71, 71, 71, 29, 29,
	29, 29, 28, 28, 31, 33, 33, 72, 72, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 69, 69,
	32, 32, 32, 32, 32, 32, 32, 9, 9, 30,
	30, 19, 19, 42, 42, 42, 42, 42, 42, 42,
	42, 42, 42, 42, 42, 43, 44, 45, 46, 47,
	48, 49, 50, 51, 52, 2, 6, 7, 7, 3,
	3, 3, 3, 61, 61, 67, 67, 67, 5, 5,
	5, 5, 57, 65, 65, 65, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 58, 58, 58,
	58, 54, 54, 54, 8, 16, 11, 11, 11, 63,
	63, 55, 55, 23, 23, 24, 24, 26, 26, 26,
	25, 25, 25, 25, 12, 39, 62, 62, 62, 62,
	62, 40, 40, 40, 40, 40, 41, 41, 41, 41,
	37, 36, 10, 35, 35, 34, 34, 4,
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
	4, 4, 6, 6, 6, 4, 3, 7, 1, 5,
	1, 3, 0, 1, 1, 2, 4, 4, 5, 1,
	1, 4, 2, 5, 1, 3, 3, 5, 6, 7,
	8, 5, 6, 1, 3, 0, 1, 3, 1, 2,
	3, 2, 4, 6, 4, 1, 3, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 9, 6,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 2,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 2, 2, 2, 3,
	3, 3, 3, 3, 3, 1, 1, 3, 3, 3,
	5, 5, 5, 3, 7, 3, 7, 8, 3, 4,
	5, 5, 3, 0, 1, 3, 4, 5, 3, 3,
	3, 3, 3, 5, 6, 5, 3, 4, 3, 3,
	2, 0, 2, 2, 3, 4, 2, 3, 5, 0,
	2, 1, 2, 2, 1, 2, 1, 1, 3, 3,
	0, 1, 3, 3, 5, 5, 0, 2, 2, 2,
	2, 5, 6, 5, 6, 5, 4, 3, 3, 2,
	4, 4, 2, 5, 7, 4, 5, 3,
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
	63, 6, 8, -30, -19, -9, 9, 42, 49, 61,
	42, 49, 11, 42, 49, 11, 42, 49, 42, 11,
	42, 11, -1, -1, -53, -66, -17, -1, -17, -66,
	15, 18, 15, 18, -66, -64, -17, -27, 58, -74,
	54, 9, -54, -5, 63, -1, -1, -17, -17, -17,
	6, 8, 66, 6, 8, -1, -1, -17, 6, -17,
	-17, -74, -17, -1, -17, -17, -17, -17, -1, -17,
	-74, -74, -17, -1, -17, -68, -17, -1, -1, -1,
	6, -60, 55, -71, 9, -29, 6, 47, 58, -60,
	-53, 40, -53, -64, -17, -5, 11, -5, -17, -4,
	-68, -17, -38, -12, -17, -12, 6, -30, -19, 11,
	-53, -57, 56, -74, 61, -64, -1, -1, -17, -1,
	-1, 6, -30, -19, -1, -1, 6, -30, -19, -1,
	-1, -1, 6, -30, -19, -69, 6, 6, -53, 51,
	52, 51, 52, -1, -63, 11, 51, -74, -1, -1,
	-1, -1, 62, 11, 62, -17, 51, 64, -61, -67,
	-20, -17, 6, 8, -64, -1, 52, 10, -74, 6,
	8, -70, -56, 42, 9, -17, 53, -17, -17, 62,
	11, 62, -5, -53, 6, 11, -71, 42, 6, 6,
	-53, 14, -33, 14, 10, 11, -74, 62, 62, 62,
	-74, -17, -5, 14, -53, -65, 6, -57, -53, -17,
	10, 62, 61, 61, 14, -58, 17, 16, -53, -53,
	14, -11, 25, -17, -62, -34, 37, -74, -74, 11,
	-74, 11, -74, 4, 53, 10, -5, -1, -64, -17,
	42, 14, -55, -11, -60, -29, 10, -17, 14, -55,
	-53, -5, -74, -17, 58, 42, 11, 58, 14, 56,
	11, -53, -74, 62, 42, -17, -17, 14, 17, 16,
	-1, -53, 14, -58, 14, -53, 8, 14, 51, 52,
	-1, -40, 15, 18, 14, 16, 37, -66, -17, -27,
	58, 64, -74, 64, -74, 64, -1, -17, 10, -1,
	14, -11, -53, 14, 14, 58, 6, -1, 6, 6,
	6, 64, 64, -1, 62, 62, -1, -53, -53, 14,
	-53, 4, -1, -1, -53, -66, -53, -17, -17, 6,
	14, -55, 6, 61, -53, 6, 51, 51, 52, 14,
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
	0, 240, 240, 15, 42, 43, 44, 45, 46, 47,
	48, 234, 240, 0, 236, 241, 237, 19, 25, 26,
	102, 13, 0, 70, 221, 0, 240, 240, 0, 0,
	0, 0, 0, 185, 186, 5, 6, 7, 240, 240,
	0, 0, 0, 0, 13, 0, 240, 0, 0, 0,
	0, 240, 0, 13, 13, 0, 240, 0, 0, 240,
	240, 240, 0, 125, 125, 15, -2, 15, -2, 73,
	82, 102, 0, 0, 0, 98, 109, 110, 31, 15,
	13, 20, -2, 22, 23, 24, 102, 240, 240, 0,
	240, 240, 0, 240, 240, 0, 240, 240, 240, 0,
	0, 0, 15, 0, 229, 233, 100, 0, 13, 235,
	240, 240, 240, 240, 0, 0, 100, 104, 0, 0,
	0, 102, 0, 262, 13, 175, 176, 177, 178, 67,
	159, 160, 0, 157, 158, 208, 216, 66, 75, 83,
	85, 0, 179, 180, 181, 182, 183, 184, 210, 0,
	0, 0, 267, 212, 84, 0, 114, 156, 209, 211,
	79, 15, 0, 123, 125, 126, 128, 0, 0, 15,
	0, 0, 0, 0, 103, 74, 13, 112, 100, 0,
	0, 139, 140, 141, 150, 151, 163, 164, 165, 13,
	0, 15, 203, 15, 102, 0, 142, 152, 0, 143,
	153, 166, 167, 168, 144, 154, 169, 170, 171, 145,
	155, 146, 172, 173, 174, 147, 0, 0, 0, 15,
	15, 16, 17, 18, 0, 0, 246, 0, 242, 243,
	238, 239, 187, 13, 188, 105, 14, 189, 13, 13,
	-2, 0, 20, -2, 0, 222, 223, 224, 15, 161,
	162, 76, 77, 240, -2, 95, 0, 260, 261, 90,
	0, 91, 80, 0, 125, 0, 0, 0, 129, 131,
	0, 132, 15, 134, 68, 13, 0, 86, 88, 89,
	0, 115, 116, 198, 0, 0, 204, 15, 13, 100,
	72, 87, 0, 0, 206, 0, 240, 15, 0, 0,
	225, 230, 15, 101, 0, 0, 0, 0, 0, 13,
	0, 13, 0, -2, 0, 71, 78, 81, 0, 244,
	240, 117, 0, 231, 15, 127, 124, 130, 121, 0,
	0, 69, 0, 111, 0, 240, 0, 0, 199, 202,
	0, 0, 0, 86, 240, 0, 0, 207, 240, 15,
	15, 220, 213, 0, 215, 226, 15, 245, 247, 248,
	249, 250, 240, 240, 263, 15, 0, 15, 106, 107,
	0, 190, 0, 191, 0, 192, 193, 195, 96, 94,
	118, 232, 0, 122, 133, 0, 113, 92, 0, 99,
	205, 200, 201, 93, 0, 149, 15, 218, 219, 214,
	227, 0, 15, 0, 0, 15, 13, 108, 0, 0,
	119, 0, 20, 0, 217, 15, 246, 15, 15, 264,
	13, 265, -2, 0, 120, 97, 0, 228, 0, 0,
	0, 266, 11, 13, 148, 251, 0, 240, 246, 253,
	0, 255, 194, 12, 196, 13, 252, 240, 246, 246,
	240, 254, 197, 246, 240, 240, 240,
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
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 90:
		//line parser.y:424
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 91:
		//line parser.y:432
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 92:
		//line parser.y:442
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 93:
		//line parser.y:450
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 94:
		//line parser.y:458
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 95:
		//line parser.y:468
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 96:
		//line parser.y:478
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 97:
		//line parser.y:480
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 98:
		//line parser.y:482
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 99:
		//line parser.y:484
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 100:
		//line parser.y:487
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 101:
		//line parser.y:489
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 102:
		//line parser.y:491
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 103:
		//line parser.y:493
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 104:
		//line parser.y:495
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 105:
		//line parser.y:497
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{Func: ast.BareReference{Name: "to_proc"}, Target: RubyS[Rubypt-0].genericValue})
		}
	case 106:
		//line parser.y:499
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 107:
		//line parser.y:501
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 108:
		//line parser.y:503
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{
				Func:   ast.BareReference{Name: "to_proc"},
				Target: RubyS[Rubypt-0].genericValue,
			})
		}
	case 109:
		//line parser.y:512
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 110:
		//line parser.y:514
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 111:
		//line parser.y:516
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 112:
		//line parser.y:518
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 113:
		//line parser.y:520
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 114:
		//line parser.y:523
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 115:
		//line parser.y:525
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 116:
		//line parser.y:527
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 117:
		//line parser.y:531
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 118:
		//line parser.y:539
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 119:
		//line parser.y:548
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 120:
		//line parser.y:557
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 121:
		//line parser.y:567
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 122:
		//line parser.y:575
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 123:
		//line parser.y:586
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 124:
		//line parser.y:588
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 125:
		//line parser.y:590
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 126:
		//line parser.y:592
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 127:
		//line parser.y:594
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 128:
		//line parser.y:597
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 129:
		//line parser.y:599
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 130:
		//line parser.y:601
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 131:
		//line parser.y:603
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 132:
		//line parser.y:607
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 133:
		//line parser.y:614
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 134:
		//line parser.y:624
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 135:
		//line parser.y:633
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 136:
		//line parser.y:639
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 137:
		//line parser.y:647
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 138:
		//line parser.y:651
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 139:
		//line parser.y:656
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 140:
		//line parser.y:663
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 141:
		//line parser.y:667
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 142:
		//line parser.y:671
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 143:
		//line parser.y:678
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 144:
		//line parser.y:685
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 145:
		//line parser.y:692
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 146:
		//line parser.y:699
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 147:
		//line parser.y:706
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: ast.Array{Nodes: RubyS[Rubypt-2].genericSlice},
				RHS: ast.Array{Nodes: RubyS[Rubypt-0].genericSlice},
			}
		}
	case 148:
		//line parser.y:714
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
	case 149:
		//line parser.y:729
		{
			tail := ast.CallExpression{Target: RubyS[Rubypt-3].genericValue, Func: ast.BareReference{Name: "[]="}, Args: []ast.Node{RubyS[Rubypt-1].genericValue}}
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, tail)
		}
	case 150:
		//line parser.y:735
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 151:
		//line parser.y:742
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 152:
		//line parser.y:746
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 153:
		//line parser.y:753
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 154:
		//line parser.y:760
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 155:
		//line parser.y:767
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 156:
		//line parser.y:774
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 157:
		//line parser.y:777
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 158:
		//line parser.y:779
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 159:
		//line parser.y:782
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 160:
		//line parser.y:784
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 161:
		//line parser.y:787
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 162:
		//line parser.y:789
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 163:
		//line parser.y:792
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 164:
		//line parser.y:794
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 165:
		//line parser.y:796
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 166:
		//line parser.y:799
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 167:
		//line parser.y:801
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 168:
		//line parser.y:803
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 169:
		//line parser.y:806
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 170:
		//line parser.y:808
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 171:
		//line parser.y:810
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 172:
		//line parser.y:813
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 173:
		//line parser.y:815
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 174:
		//line parser.y:817
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 175:
		//line parser.y:819
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 176:
		//line parser.y:820
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 177:
		//line parser.y:821
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 178:
		//line parser.y:822
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 179:
		//line parser.y:825
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 180:
		//line parser.y:834
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 181:
		//line parser.y:843
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 182:
		//line parser.y:852
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 183:
		//line parser.y:861
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 184:
		//line parser.y:870
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 185:
		//line parser.y:878
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 186:
		//line parser.y:879
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 187:
		//line parser.y:881
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 188:
		//line parser.y:883
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 189:
		//line parser.y:886
		{
			RubyVAL.genericValue = ast.Hash{}
		}
	case 190:
		//line parser.y:888
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 191:
		//line parser.y:896
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 192:
		//line parser.y:904
		{
			RubyVAL.genericValue = ast.Block{Body: ast.Nodes{RubyS[Rubypt-2].genericValue}}
		}
	case 193:
		//line parser.y:907
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 194:
		//line parser.y:914
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 195:
		//line parser.y:922
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 196:
		//line parser.y:929
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 197:
		//line parser.y:936
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 198:
		//line parser.y:944
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 199:
		//line parser.y:946
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 200:
		//line parser.y:953
		{
			RubyVAL.genericValue = ast.Block{Args: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 201:
		//line parser.y:957
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-2].genericSlice}
		}
	case 202:
		//line parser.y:960
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 203:
		//line parser.y:962
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 204:
		//line parser.y:964
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 205:
		//line parser.y:966
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 206:
		//line parser.y:969
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 207:
		//line parser.y:976
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 208:
		//line parser.y:984
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 209:
		//line parser.y:991
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 210:
		//line parser.y:998
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 211:
		//line parser.y:1005
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 212:
		//line parser.y:1012
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 213:
		//line parser.y:1019
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 214:
		//line parser.y:1026
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 215:
		//line parser.y:1034
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 216:
		//line parser.y:1041
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 217:
		//line parser.y:1050
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 218:
		//line parser.y:1057
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 219:
		//line parser.y:1064
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 220:
		//line parser.y:1071
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 221:
		//line parser.y:1078
		{
		}
	case 222:
		//line parser.y:1079
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 223:
		//line parser.y:1080
		{
		}
	case 224:
		//line parser.y:1083
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 225:
		//line parser.y:1086
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 226:
		//line parser.y:1094
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 227:
		//line parser.y:1096
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 228:
		//line parser.y:1105
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
	case 229:
		//line parser.y:1120
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 230:
		//line parser.y:1122
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 231:
		//line parser.y:1125
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 232:
		//line parser.y:1127
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 233:
		//line parser.y:1130
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 234:
		//line parser.y:1137
		{
			RubyVAL.genericValue = ast.Yield{}
		}
	case 235:
		//line parser.y:1140
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 236:
		//line parser.y:1148
		{
			RubyVAL.genericValue = ast.Return{}
		}
	case 237:
		//line parser.y:1152
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 238:
		//line parser.y:1154
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 239:
		//line parser.y:1156
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 240:
		//line parser.y:1159
		{
		}
	case 241:
		//line parser.y:1161
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 242:
		//line parser.y:1163
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 243:
		//line parser.y:1165
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 244:
		//line parser.y:1169
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 245:
		//line parser.y:1178
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 246:
		//line parser.y:1181
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 247:
		//line parser.y:1183
		{
		}
	case 248:
		//line parser.y:1185
		{
		}
	case 249:
		//line parser.y:1187
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 250:
		//line parser.y:1189
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 251:
		//line parser.y:1192
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 252:
		//line parser.y:1199
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 253:
		//line parser.y:1207
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 254:
		//line parser.y:1214
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 255:
		//line parser.y:1222
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 256:
		//line parser.y:1230
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 257:
		//line parser.y:1237
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 258:
		//line parser.y:1244
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 259:
		//line parser.y:1251
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 260:
		//line parser.y:1259
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 261:
		//line parser.y:1262
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 262:
		//line parser.y:1264
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 263:
		//line parser.y:1267
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 264:
		//line parser.y:1269
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 265:
		//line parser.y:1272
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 266:
		//line parser.y:1274
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 267:
		//line parser.y:1276
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
