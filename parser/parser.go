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

//line parser.y:1325

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 46,
	54, 144,
	-2, 21,
	-1, 118,
	54, 144,
	-2, 142,
	-1, 120,
	10, 107,
	11, 107,
	-2, 228,
	-1, 134,
	54, 144,
	-2, 21,
	-1, 175,
	54, 144,
	-2, 21,
	-1, 304,
	51, 13,
	64, 13,
	-2, 31,
	-1, 307,
	54, 145,
	-2, 143,
	-1, 318,
	10, 107,
	11, 107,
	-2, 228,
}

const RubyNprod = 275
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 3401

var RubyAct = []int{

	286, 391, 5, 390, 509, 362, 168, 221, 351, 225,
	122, 30, 223, 61, 356, 51, 201, 121, 25, 2,
	3, 90, 127, 342, 91, 300, 190, 275, 191, 109,
	300, 96, 169, 18, 269, 13, 4, 26, 463, 300,
	264, 246, 371, 446, 236, 131, 300, 96, 485, 252,
	110, 138, 444, 155, 156, 117, 119, 88, 87, 442,
	105, 106, 360, 359, 160, 161, 142, 94, 95, 158,
	131, 114, 180, 163, 89, 107, 105, 106, 324, 185,
	186, 96, 93, 94, 95, 129, 192, 83, 300, 355,
	179, 195, 196, 496, 83, 405, 184, 143, 93, 203,
	83, 83, 92, 139, 208, 507, 171, 324, 172, 213,
	105, 106, 217, 218, 219, 408, 142, 94, 95, 324,
	300, 184, 210, 211, 229, 252, 300, 297, 481, 357,
	47, 288, 93, 215, 226, 403, 104, 224, 243, 245,
	251, 259, 260, 242, 262, 263, 240, 267, 268, 253,
	272, 273, 274, 233, 171, 96, 172, 171, 343, 172,
	407, 239, 258, 183, 404, 292, 293, 294, 295, 278,
	323, 255, 171, 300, 172, 227, 135, 257, 298, 291,
	374, 149, 296, 309, 105, 106, 228, 138, 150, 413,
	412, 94, 95, 135, 300, 403, 135, 135, 388, 331,
	149, 312, 112, 176, 146, 113, 93, 235, 237, 315,
	104, 135, 135, 135, 308, 171, 316, 172, 154, 152,
	140, 145, 231, 135, 108, 135, 135, 109, 135, 148,
	135, 135, 135, 135, 288, 135, 111, 330, 135, 139,
	135, 135, 404, 336, 379, 96, 377, 138, 110, 153,
	151, 176, 143, 340, 176, 135, 135, 247, 181, 144,
	146, 434, 154, 435, 135, 96, 344, 147, 329, 176,
	135, 142, 226, 135, 105, 106, 265, 240, 249, 270,
	140, 94, 95, 276, 436, 425, 497, 141, 171, 368,
	172, 313, 239, 314, 105, 106, 93, 369, 454, 139,
	118, 94, 95, 467, 135, 135, 326, 307, 226, 369,
	135, 224, 176, 227, 375, 451, 93, 367, 385, 376,
	378, 380, 90, 466, 228, 91, 369, 90, 369, 135,
	91, 90, 135, 372, 91, 346, 392, 488, 397, 393,
	90, 135, 135, 91, 193, 90, 194, 90, 91, 227,
	91, 386, 171, 116, 172, 115, 400, 222, 490, 491,
	228, 96, 166, 489, 419, 167, 473, 290, 518, 410,
	461, 422, 429, 429, 460, 164, 282, 283, 165, 459,
	135, 437, 447, 289, 135, 176, 449, 339, 457, 450,
	105, 106, 452, 350, 443, 515, 445, 94, 95, 452,
	333, 237, 384, 524, 458, 521, 520, 394, 329, 440,
	346, 439, 93, 516, 465, 383, 297, 90, 468, 135,
	91, 332, 328, 96, 399, 519, 74, 521, 520, 464,
	354, 297, 474, 475, 440, 280, 439, 135, 471, 279,
	418, 417, 416, 477, 418, 417, 96, 338, 339, 176,
	220, 135, 105, 106, 198, 135, 303, 483, 349, 94,
	95, 287, 135, 302, 1, 486, 182, 60, 59, 58,
	57, 135, 56, 55, 93, 105, 106, 38, 37, 36,
	48, 358, 94, 95, 35, 452, 50, 97, 98, 99,
	135, 135, 430, 20, 40, 494, 505, 93, 102, 100,
	101, 104, 429, 429, 429, 135, 176, 513, 96, 41,
	504, 522, 21, 135, 373, 15, 12, 11, 24, 96,
	23, 526, 22, 517, 429, 19, 136, 10, 429, 429,
	429, 176, 501, 502, 503, 525, 32, 105, 106, 27,
	17, 14, 39, 136, 94, 95, 136, 136, 105, 106,
	16, 34, 33, 177, 523, 94, 95, 28, 71, 93,
	29, 136, 136, 136, 527, 528, 325, 135, 320, 529,
	93, 70, 135, 136, 135, 136, 136, 0, 136, 0,
	136, 136, 136, 136, 0, 136, 0, 135, 136, 0,
	136, 136, 0, 0, 0, 495, 0, 0, 0, 0,
	0, 177, 0, 0, 177, 136, 136, 248, 96, 0,
	0, 0, 381, 0, 136, 135, 135, 0, 0, 177,
	136, 0, 0, 136, 105, 106, 266, 135, 0, 271,
	0, 94, 95, 277, 0, 0, 0, 105, 106, 0,
	0, 105, 106, 0, 94, 95, 93, 0, 94, 95,
	31, 0, 0, 0, 136, 136, 0, 0, 0, 93,
	136, 0, 177, 93, 69, 133, 75, 134, 120, 0,
	126, 131, 0, 0, 0, 0, 0, 0, 0, 136,
	0, 0, 136, 0, 0, 0, 0, 0, 0, 0,
	0, 136, 136, 0, 77, 0, 130, 85, 86, 0,
	0, 124, 78, 79, 80, 81, 82, 0, 125, 0,
	0, 0, 0, 130, 0, 0, 130, 130, 0, 0,
	123, 0, 132, 130, 84, 83, 0, 0, 0, 0,
	136, 130, 130, 130, 136, 177, 0, 0, 0, 0,
	0, 0, 0, 130, 0, 130, 130, 0, 130, 0,
	130, 130, 130, 130, 0, 130, 0, 0, 130, 0,
	130, 130, 0, 0, 0, 0, 0, 0, 0, 136,
	0, 130, 0, 0, 130, 130, 130, 69, 306, 75,
	134, 76, 0, 0, 130, 0, 0, 136, 0, 130,
	130, 0, 0, 130, 0, 0, 0, 0, 0, 177,
	0, 136, 0, 0, 0, 136, 0, 77, 0, 0,
	85, 86, 136, 0, 0, 78, 79, 80, 81, 82,
	0, 136, 0, 300, 130, 130, 0, 0, 252, 0,
	304, 0, 130, 72, 0, 73, 301, 84, 83, 0,
	136, 136, 0, 0, 0, 0, 0, 0, 0, 130,
	0, 0, 130, 0, 0, 136, 177, 0, 0, 0,
	0, 130, 130, 136, 0, 0, 0, 0, 0, 69,
	133, 75, 134, 120, 498, 0, 131, 0, 0, 0,
	0, 177, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 77,
	130, 0, 85, 86, 304, 130, 0, 78, 79, 80,
	81, 82, 0, 0, 0, 0, 0, 136, 0, 0,
	0, 0, 136, 0, 136, 254, 0, 132, 0, 84,
	83, 0, 0, 0, 0, 0, 0, 136, 0, 130,
	0, 0, 0, 0, 0, 0, 69, 174, 75, 175,
	76, 0, 0, 0, 0, 0, 0, 130, 0, 0,
	0, 0, 0, 0, 0, 136, 136, 0, 0, 130,
	0, 130, 0, 0, 9, 130, 77, 136, 0, 85,
	86, 0, 130, 0, 78, 79, 80, 81, 82, 0,
	0, 130, 300, 0, 0, 0, 0, 0, 0, 456,
	0, 0, 72, 0, 73, 0, 84, 83, 281, 0,
	130, 130, 0, 0, 0, 0, 0, 0, 0, 0,
	128, 0, 0, 0, 0, 130, 130, 0, 0, 0,
	0, 0, 0, 130, 0, 0, 0, 159, 0, 0,
	162, 159, 0, 0, 0, 0, 0, 170, 0, 0,
	0, 130, 0, 0, 0, 187, 188, 189, 0, 0,
	0, 0, 0, 157, 0, 0, 0, 197, 0, 199,
	200, 0, 202, 0, 204, 205, 206, 207, 0, 209,
	0, 0, 212, 0, 214, 216, 0, 130, 0, 0,
	0, 0, 130, 0, 130, 234, 0, 0, 238, 241,
	244, 0, 0, 0, 0, 0, 0, 130, 128, 0,
	0, 0, 0, 234, 256, 0, 96, 261, 0, 0,
	0, 0, 0, 0, 0, 0, 230, 0, 232, 0,
	0, 0, 0, 0, 0, 130, 130, 0, 0, 0,
	250, 0, 0, 0, 49, 105, 106, 130, 299, 128,
	0, 0, 94, 95, 305, 0, 234, 97, 98, 99,
	107, 0, 69, 174, 75, 175, 76, 93, 102, 100,
	101, 104, 0, 128, 411, 0, 319, 0, 0, 0,
	0, 0, 0, 0, 0, 321, 322, 0, 0, 0,
	137, 0, 77, 0, 0, 85, 86, 0, 0, 0,
	78, 79, 80, 81, 82, 0, 0, 137, 300, 0,
	137, 137, 0, 0, 0, 441, 0, 178, 72, 0,
	73, 0, 84, 83, 345, 137, 137, 137, 305, 353,
	327, 0, 0, 0, 0, 0, 0, 137, 334, 137,
	137, 0, 137, 0, 137, 137, 137, 137, 0, 137,
	0, 0, 137, 0, 137, 137, 0, 0, 0, 0,
	348, 0, 352, 370, 0, 178, 0, 0, 178, 137,
	137, 0, 0, 0, 0, 0, 0, 0, 137, 0,
	0, 128, 0, 178, 137, 0, 0, 137, 0, 0,
	0, 365, 366, 234, 0, 387, 0, 0, 0, 345,
	0, 0, 0, 0, 0, 0, 395, 0, 0, 0,
	0, 0, 0, 0, 0, 401, 0, 0, 137, 137,
	0, 352, 0, 0, 137, 0, 178, 0, 69, 133,
	75, 134, 76, 0, 414, 415, 0, 0, 0, 0,
	0, 0, 0, 137, 0, 398, 137, 0, 0, 159,
	438, 0, 0, 0, 0, 137, 137, 448, 77, 0,
	409, 85, 86, 0, 0, 0, 78, 79, 80, 81,
	82, 0, 0, 420, 300, 438, 0, 0, 424, 0,
	0, 402, 0, 0, 72, 0, 73, 0, 84, 83,
	0, 0, 0, 0, 137, 0, 0, 0, 137, 178,
	0, 453, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 159, 0, 0, 0, 0, 479, 0, 480, 0,
	0, 0, 0, 0, 0, 0, 0, 469, 470, 0,
	0, 479, 0, 137, 472, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 476, 0, 478, 0, 0, 0,
	0, 137, 0, 0, 69, 133, 75, 134, 120, 128,
	499, 131, 0, 178, 0, 137, 0, 0, 0, 137,
	0, 506, 0, 0, 0, 0, 137, 487, 0, 0,
	0, 0, 0, 0, 77, 137, 493, 85, 86, 0,
	0, 0, 78, 79, 80, 81, 82, 500, 0, 365,
	366, 0, 382, 0, 137, 137, 0, 0, 0, 0,
	254, 0, 132, 0, 84, 83, 0, 0, 0, 137,
	178, 0, 0, 0, 0, 0, 0, 137, 0, 69,
	45, 75, 46, 76, 0, 0, 0, 42, 512, 431,
	511, 510, 432, 43, 44, 178, 63, 64, 54, 0,
	0, 67, 68, 0, 0, 66, 62, 0, 0, 77,
	65, 0, 85, 86, 0, 0, 0, 78, 79, 80,
	81, 82, 0, 0, 0, 427, 428, 0, 0, 0,
	0, 137, 0, 0, 0, 72, 137, 73, 137, 84,
	83, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 137, 0, 0, 0, 69, 45, 75, 46, 76,
	0, 0, 0, 42, 508, 431, 511, 510, 432, 43,
	44, 0, 63, 64, 54, 0, 0, 67, 68, 137,
	137, 66, 62, 0, 0, 77, 65, 0, 85, 86,
	0, 137, 0, 78, 79, 80, 81, 82, 0, 0,
	0, 427, 428, 0, 0, 0, 0, 0, 0, 0,
	0, 72, 0, 73, 0, 84, 83, 69, 45, 75,
	46, 76, 0, 0, 0, 42, 421, 52, 364, 363,
	53, 43, 44, 0, 63, 64, 54, 0, 0, 67,
	68, 0, 0, 66, 62, 0, 0, 77, 65, 0,
	85, 86, 0, 0, 0, 78, 79, 80, 81, 82,
	0, 0, 0, 284, 285, 0, 0, 0, 0, 0,
	0, 0, 0, 72, 0, 73, 0, 84, 83, 69,
	45, 75, 46, 76, 0, 0, 0, 42, 361, 52,
	364, 363, 53, 43, 44, 0, 63, 64, 54, 0,
	0, 67, 68, 0, 0, 66, 62, 0, 0, 77,
	65, 0, 85, 86, 0, 0, 0, 78, 79, 80,
	81, 82, 0, 0, 0, 284, 285, 0, 69, 45,
	75, 46, 76, 0, 0, 72, 42, 73, 52, 84,
	83, 53, 43, 44, 0, 63, 64, 54, 0, 0,
	67, 68, 0, 0, 66, 62, 0, 0, 77, 65,
	0, 85, 86, 0, 0, 0, 78, 79, 80, 81,
	82, 0, 0, 0, 6, 7, 0, 0, 0, 0,
	0, 0, 0, 0, 72, 0, 73, 0, 84, 83,
	0, 8, 69, 45, 75, 46, 76, 0, 0, 0,
	42, 482, 52, 0, 0, 53, 43, 44, 0, 63,
	64, 54, 369, 0, 67, 68, 0, 0, 66, 62,
	0, 0, 77, 65, 0, 85, 86, 0, 0, 0,
	78, 79, 80, 81, 82, 0, 0, 0, 284, 285,
	0, 0, 0, 0, 0, 0, 0, 0, 72, 0,
	73, 0, 84, 83, 69, 45, 75, 46, 76, 0,
	0, 0, 42, 396, 52, 0, 0, 53, 43, 44,
	0, 63, 64, 54, 369, 0, 67, 68, 0, 0,
	66, 62, 0, 0, 77, 65, 0, 85, 86, 0,
	0, 0, 78, 79, 80, 81, 82, 0, 0, 0,
	284, 285, 0, 0, 0, 0, 0, 0, 0, 0,
	72, 0, 73, 0, 84, 83, 69, 45, 75, 46,
	76, 0, 0, 0, 42, 389, 52, 0, 0, 53,
	43, 44, 0, 63, 64, 54, 369, 0, 67, 68,
	0, 0, 66, 62, 0, 0, 77, 65, 0, 85,
	86, 0, 0, 0, 78, 79, 80, 81, 82, 0,
	0, 0, 284, 285, 0, 0, 0, 0, 0, 0,
	0, 0, 72, 0, 73, 0, 84, 83, 69, 45,
	75, 46, 76, 0, 0, 0, 42, 514, 431, 0,
	0, 432, 43, 44, 0, 63, 64, 54, 0, 0,
	67, 68, 0, 0, 66, 62, 0, 0, 77, 65,
	0, 85, 86, 0, 0, 0, 78, 79, 80, 81,
	82, 0, 0, 0, 427, 428, 0, 0, 0, 0,
	0, 0, 0, 0, 72, 0, 73, 0, 84, 83,
	69, 45, 75, 46, 76, 0, 0, 0, 42, 492,
	52, 0, 0, 53, 43, 44, 0, 63, 64, 54,
	0, 0, 67, 68, 0, 0, 66, 62, 0, 0,
	77, 65, 0, 85, 86, 0, 0, 0, 78, 79,
	80, 81, 82, 0, 0, 0, 284, 285, 0, 69,
	45, 75, 46, 76, 0, 0, 72, 42, 73, 52,
	84, 83, 53, 43, 44, 0, 63, 64, 54, 0,
	0, 67, 68, 0, 0, 66, 62, 0, 0, 77,
	65, 0, 85, 86, 0, 0, 0, 78, 79, 80,
	81, 82, 0, 0, 0, 284, 285, 0, 0, 0,
	0, 0, 0, 0, 0, 72, 0, 73, 462, 84,
	83, 69, 45, 75, 46, 76, 0, 0, 0, 42,
	455, 52, 0, 0, 53, 43, 44, 0, 63, 64,
	54, 0, 0, 67, 68, 0, 0, 66, 62, 0,
	0, 77, 65, 0, 85, 86, 0, 0, 0, 78,
	79, 80, 81, 82, 0, 0, 0, 284, 285, 0,
	0, 0, 0, 0, 0, 0, 0, 72, 0, 73,
	0, 84, 83, 69, 45, 75, 46, 76, 0, 0,
	0, 42, 433, 431, 0, 0, 432, 43, 44, 0,
	63, 64, 54, 0, 0, 67, 68, 0, 0, 66,
	62, 0, 0, 77, 65, 0, 85, 86, 0, 0,
	0, 78, 79, 80, 81, 82, 0, 0, 0, 427,
	428, 0, 0, 0, 0, 0, 0, 0, 0, 72,
	0, 73, 0, 84, 83, 69, 45, 75, 46, 76,
	0, 0, 0, 42, 426, 431, 0, 0, 432, 43,
	44, 0, 63, 64, 54, 0, 0, 67, 68, 0,
	0, 66, 62, 0, 0, 77, 65, 0, 85, 86,
	0, 0, 0, 78, 79, 80, 81, 82, 0, 0,
	0, 427, 428, 0, 0, 0, 0, 0, 0, 0,
	0, 72, 0, 73, 0, 84, 83, 69, 45, 75,
	46, 76, 0, 0, 0, 42, 423, 52, 0, 0,
	53, 43, 44, 0, 63, 64, 54, 0, 0, 67,
	68, 0, 0, 66, 62, 0, 0, 77, 65, 0,
	85, 86, 0, 0, 0, 78, 79, 80, 81, 82,
	0, 0, 0, 284, 285, 0, 0, 0, 0, 0,
	0, 0, 0, 72, 0, 73, 0, 84, 83, 69,
	45, 75, 46, 76, 0, 0, 0, 42, 406, 52,
	0, 0, 53, 43, 44, 0, 63, 64, 54, 0,
	0, 67, 68, 0, 0, 66, 62, 0, 0, 77,
	65, 0, 85, 86, 0, 0, 0, 78, 79, 80,
	81, 82, 0, 0, 0, 284, 285, 0, 0, 0,
	0, 0, 0, 0, 0, 72, 0, 73, 0, 84,
	83, 69, 45, 75, 46, 76, 0, 0, 0, 42,
	347, 52, 0, 0, 53, 43, 44, 0, 63, 64,
	54, 0, 0, 67, 68, 0, 0, 66, 62, 0,
	0, 77, 65, 0, 85, 86, 0, 0, 0, 78,
	79, 80, 81, 82, 0, 0, 0, 284, 285, 0,
	0, 0, 0, 0, 0, 0, 0, 72, 0, 73,
	0, 84, 83, 69, 45, 75, 46, 76, 0, 0,
	0, 42, 337, 52, 0, 0, 53, 43, 44, 0,
	63, 64, 54, 0, 0, 67, 68, 0, 0, 66,
	62, 0, 0, 77, 65, 0, 85, 86, 0, 0,
	0, 78, 79, 80, 81, 82, 0, 0, 0, 284,
	285, 0, 0, 0, 0, 0, 0, 0, 0, 72,
	0, 73, 0, 84, 83, 69, 45, 75, 46, 76,
	0, 0, 0, 42, 335, 52, 0, 0, 53, 43,
	44, 0, 63, 64, 54, 0, 0, 67, 68, 0,
	0, 66, 62, 0, 0, 77, 65, 0, 85, 86,
	0, 0, 0, 78, 79, 80, 81, 82, 0, 0,
	0, 284, 285, 0, 69, 45, 75, 46, 76, 0,
	0, 72, 42, 73, 431, 84, 83, 432, 43, 44,
	0, 63, 64, 54, 0, 0, 67, 68, 0, 0,
	66, 62, 0, 0, 77, 65, 0, 85, 86, 0,
	0, 0, 78, 79, 80, 81, 82, 0, 0, 0,
	427, 428, 0, 69, 45, 75, 46, 76, 0, 0,
	72, 42, 73, 52, 84, 83, 53, 43, 44, 0,
	63, 64, 54, 0, 0, 67, 68, 0, 0, 66,
	62, 0, 0, 77, 65, 0, 85, 86, 0, 0,
	0, 78, 79, 80, 81, 82, 0, 0, 0, 284,
	285, 0, 69, 45, 75, 46, 76, 311, 0, 72,
	42, 73, 52, 84, 83, 53, 43, 44, 0, 63,
	64, 54, 0, 0, 67, 68, 0, 0, 66, 62,
	0, 0, 77, 65, 0, 85, 86, 0, 0, 0,
	78, 79, 80, 81, 82, 0, 0, 0, 0, 310,
	0, 69, 45, 75, 46, 76, 0, 0, 72, 42,
	73, 52, 84, 83, 53, 43, 44, 0, 63, 64,
	54, 0, 0, 67, 68, 0, 0, 66, 62, 0,
	0, 77, 65, 0, 85, 86, 0, 0, 0, 78,
	79, 80, 81, 82, 0, 0, 0, 300, 0, 0,
	69, 45, 75, 46, 76, 0, 0, 72, 42, 73,
	52, 84, 83, 53, 43, 44, 0, 63, 64, 54,
	0, 0, 67, 68, 0, 0, 66, 62, 0, 0,
	77, 65, 0, 85, 86, 0, 0, 0, 78, 79,
	80, 81, 82, 69, 133, 75, 134, 318, 0, 0,
	131, 0, 0, 0, 0, 0, 72, 0, 73, 0,
	84, 83, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 77, 0, 0, 85, 86, 0, 0,
	317, 78, 79, 80, 81, 82, 69, 306, 75, 134,
	76, 0, 0, 0, 0, 0, 0, 0, 0, 72,
	0, 132, 0, 84, 83, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 77, 0, 0, 85,
	86, 0, 0, 0, 78, 79, 80, 81, 82, 0,
	0, 0, 300, 69, 133, 75, 134, 120, 0, 0,
	131, 0, 72, 0, 73, 301, 84, 83, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 77, 0, 0, 85, 86, 0, 0,
	124, 78, 79, 80, 81, 82, 69, 133, 75, 134,
	76, 0, 0, 0, 0, 0, 0, 0, 0, 254,
	0, 132, 0, 84, 83, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 77, 0, 0, 85,
	86, 0, 0, 0, 78, 79, 80, 81, 82, 0,
	0, 0, 300, 69, 133, 75, 134, 76, 0, 0,
	131, 0, 72, 0, 73, 0, 84, 83, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 77, 0, 0, 85, 86, 0, 0,
	0, 78, 79, 80, 81, 82, 69, 174, 75, 175,
	76, 0, 0, 0, 0, 0, 0, 0, 0, 72,
	0, 132, 0, 84, 83, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 77, 0, 0, 85,
	86, 0, 0, 0, 78, 79, 80, 81, 82, 69,
	133, 75, 134, 120, 0, 0, 131, 0, 0, 173,
	0, 0, 72, 0, 73, 0, 84, 83, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 77,
	0, 0, 85, 86, 0, 0, 0, 78, 79, 80,
	81, 82, 69, 133, 75, 134, 76, 0, 0, 0,
	0, 0, 0, 0, 0, 254, 0, 132, 0, 84,
	83, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 77, 0, 0, 85, 86, 0, 0, 0,
	78, 79, 80, 81, 82, 69, 484, 75, 134, 76,
	0, 0, 0, 0, 0, 0, 0, 0, 72, 0,
	73, 0, 84, 83, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 77, 0, 0, 85, 86,
	96, 0, 0, 78, 79, 80, 81, 82, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 72, 0, 73, 0, 84, 83, 0, 0, 105,
	106, 0, 0, 0, 0, 96, 94, 95, 0, 0,
	0, 97, 98, 99, 107, 0, 0, 0, 0, 103,
	0, 93, 102, 100, 101, 104, 92, 0, 341, 0,
	0, 0, 0, 0, 105, 106, 0, 0, 0, 0,
	0, 94, 95, 0, 0, 0, 97, 98, 99, 107,
	0, 0, 0, 0, 0, 0, 93, 102, 100, 101,
	104,
}
var RubyPact = []int{

	-32, 1773, -1000, -1000, -1000, 6, -1000, -1000, -1000, 3341,
	-1000, -1000, -1000, 206, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 187, -1000, 16, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 349, 292, 292, 659, 238, 210, 218, 139,
	208, 207, 2875, 2875, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 3227, 2875, 2875, 3227, 3227, 360, 347, -1000,
	-1000, -1000, 3141, -1000, 18, 249, -1000, 58, 2875, 2875,
	3227, 3227, 3227, 20, 338, -1000, -1000, -1000, -1000, -1000,
	2875, 2875, 3227, 448, 3227, 3227, -1000, 3227, 2875, 3227,
	3227, 3227, 3227, 2875, 3227, -1000, -1000, 3227, 2875, 3227,
	3227, 2875, 2875, 2875, 444, 302, 128, 182, -1000, -1000,
	3141, 58, 33, 3141, 3227, 3227, 35, 267, 604, -1000,
	-11, -7, -1000, 3184, 42, 5, -1000, -1000, 3141, 3227,
	2875, 2875, 3227, 2875, 2875, 34, 2875, 2875, 28, 2875,
	2875, 2875, 21, 433, 429, 330, 325, 2728, 223, 604,
	332, 316, 604, 223, 2875, 2875, 2875, 2875, 120, 116,
	442, -1000, -1000, 3227, 3008, 178, 55, 162, 158, 2961,
	299, 3141, 2777, -1000, -1000, 330, 330, 604, 604, 604,
	-1000, -1000, 285, -1000, -1000, 330, 330, 604, 2918, 604,
	604, 3051, 604, 330, 604, 604, 604, 604, 330, 515,
	3051, 3051, 604, 330, 604, 108, 504, 330, 330, 330,
	58, -1000, 416, 257, 266, -1000, 157, 415, 394, -1000,
	2630, 292, 2568, 437, 442, -1000, -1000, -1000, 3306, -39,
	96, 77, -1000, -1000, 151, -1000, -1000, -1000, -1000, 3098,
	2506, -1000, 387, 772, 3141, 420, 27, -48, 67, 330,
	330, 419, 330, 330, -1000, -1000, -1000, 330, 330, -1000,
	-1000, -1000, 330, 330, 330, -1000, -1000, -1000, 251, 2,
	1, 1724, -1000, -1000, -1000, -1000, 330, 303, 3227, -1000,
	-1000, 143, 330, 330, 330, 330, -1000, -1000, -1000, 604,
	-1000, -1000, 235, 233, -11, 608, 1449, -1000, 405, 330,
	-1000, -1000, 69, -1000, -1000, 33, 58, 2875, 3141, 604,
	3227, 604, 604, -1000, 3098, 156, -1000, 1961, 128, 266,
	397, 3227, -1000, -1000, 1899, -1000, -1000, -1000, 58, -1000,
	1323, 153, -1000, -1000, 37, 604, -1000, -1000, 2444, 104,
	-1000, -1000, 2728, 1112, -1000, 148, -1000, -1000, 147, 3227,
	3227, -1000, 428, 2875, -1000, 1662, 2382, -1000, -1000, 277,
	604, 2320, 2258, 247, 3227, 1157, -5, -1000, -12, -1000,
	-21, 2875, 3227, -1000, -1000, 330, 376, 604, 2875, -1000,
	301, -1000, -1000, -1000, -1000, 604, -1000, 284, 2196, -1000,
	941, 604, 382, 2875, 373, 368, -1000, -1000, 364, 2134,
	-26, 93, -1000, 2875, 261, 241, -1000, 2875, -1000, 330,
	2728, -1000, 424, -1000, 2728, 362, -1000, -1000, -1000, 330,
	-1000, 2875, 2875, -1000, -1000, -1000, 3227, 223, 442, -1000,
	-1000, 3227, -1000, 3051, -1000, 122, -1000, 330, 604, -1000,
	330, -1000, -1000, 1837, -1000, -1000, 3270, -1000, 330, -13,
	-1000, -1000, -1000, -1000, 2826, 330, 231, -1000, 330, 2728,
	2728, -1000, 2728, 331, 312, 307, 2085, 223, 2728, 604,
	591, 40, -1000, 272, 864, 3227, 330, 2728, -1000, -1000,
	-1000, -1000, -1000, 2728, 75, 2875, 3227, -1000, -1000, 43,
	2728, 1600, 1524, 2023, 75, 402, 357, -1000, -1000, 411,
	2875, -1000, -1000, 389, -1000, -1000, -1000, 75, -1000, -1000,
	2875, -1000, 330, 2679, -1000, 75, 330, 2679, 2679, 2679,
}
var RubyPgo = []int{

	0, 0, 571, 560, 37, 85, 558, 557, 552, 1144,
	551, 1, 13, 550, 542, 35, 541, 540, 974, 539,
	480, 650, 536, 527, 525, 522, 520, 518, 33, 517,
	9, 130, 516, 515, 11, 514, 512, 509, 494, 18,
	493, 492, 4, 486, 484, 479, 478, 477, 473, 472,
	470, 469, 468, 467, 1008, 466, 3, 17, 8, 5,
	464, 7, 463, 42, 461, 32, 458, 6, 456, 22,
	15, 10, 12, 426, 395, 16,
}
var RubyR1 = []int{

	0, 60, 60, 60, 60, 60, 60, 60, 60, 60,
	60, 74, 74, 75, 75, 54, 54, 54, 54, 19,
	19, 19, 19, 19, 19, 19, 19, 18, 18, 18,
	18, 18, 18, 18, 18, 18, 18, 18, 18, 18,
	18, 18, 28, 28, 28, 28, 28, 28, 28, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 39, 14, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	22, 57, 57, 57, 57, 67, 67, 65, 65, 65,
	65, 65, 65, 65, 65, 65, 71, 71, 71, 71,
	71, 69, 69, 69, 23, 23, 23, 23, 23, 23,
	61, 61, 72, 72, 72, 30, 30, 30, 30, 29,
	29, 32, 34, 34, 73, 73, 15, 15, 15, 15,
	15, 15, 15, 16, 16, 70, 70, 33, 33, 33,
	33, 33, 33, 33, 9, 9, 31, 31, 20, 20,
	43, 43, 43, 43, 43, 43, 43, 43, 43, 43,
	43, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 2, 6, 7, 7, 3, 3, 3, 3,
	62, 62, 68, 68, 68, 5, 5, 5, 5, 58,
	66, 66, 66, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 59, 59, 59, 59, 55, 55,
	55, 8, 17, 11, 11, 11, 64, 64, 56, 56,
	24, 24, 25, 25, 27, 27, 27, 26, 26, 26,
	12, 40, 40, 63, 63, 63, 63, 63, 41, 41,
	41, 41, 41, 42, 42, 42, 42, 38, 37, 10,
	36, 36, 35, 35, 4,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 1, 0, 2, 0, 2, 2, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 2, 4,
	5, 1, 4, 4, 2, 3, 3, 4, 4, 5,
	3, 4, 5, 2, 3, 3, 3, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 6, 7, 6, 6,
	4, 3, 7, 1, 5, 1, 3, 0, 1, 1,
	1, 2, 4, 4, 4, 5, 1, 1, 4, 2,
	5, 1, 3, 3, 5, 6, 7, 8, 5, 6,
	1, 3, 0, 1, 3, 1, 2, 3, 2, 4,
	6, 4, 1, 3, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 9, 6, 3, 3, 3,
	3, 3, 3, 3, 2, 2, 2, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 1, 1, 3, 3, 3, 5, 5, 5,
	3, 7, 3, 7, 8, 3, 4, 5, 5, 3,
	0, 1, 3, 4, 5, 3, 3, 3, 3, 3,
	5, 6, 5, 3, 4, 3, 3, 2, 0, 2,
	2, 3, 4, 2, 3, 5, 0, 2, 1, 2,
	2, 1, 2, 1, 1, 3, 3, 1, 3, 3,
	5, 5, 5, 0, 2, 2, 2, 2, 5, 6,
	5, 6, 5, 4, 3, 3, 2, 4, 4, 2,
	5, 7, 4, 5, 3,
}
var RubyChk = []int{

	-1000, -60, 51, 52, 68, -1, 51, 52, 68, -18,
	-23, -29, -32, -15, -16, -33, -13, -17, -28, -24,
	-40, -36, -25, -26, -27, -39, -4, -19, -7, -3,
	-34, -21, -22, -8, -10, -44, -45, -46, -47, -14,
	-38, -37, 13, 19, 20, 6, 8, -31, -20, -9,
	-43, -70, 15, 18, 24, -48, -49, -50, -51, -52,
	-53, -12, 32, 22, 23, 36, 31, 27, 28, 5,
	-2, -6, 61, 63, -73, 7, 9, 35, 43, 44,
	45, 46, 47, 66, 65, 38, 39, 52, 51, 68,
	15, 18, 25, 55, 40, 41, 4, 45, 46, 47,
	57, 58, 56, 18, 59, 33, 34, 48, 18, 40,
	61, 49, 15, 18, 55, 6, 4, -34, 8, -34,
	9, -57, -71, 61, 42, 49, 11, -69, -18, -5,
	-21, 12, 63, 6, 8, -31, -20, -9, 9, 61,
	42, 49, 61, 42, 49, 11, 42, 49, 11, 42,
	49, 42, 11, 42, 11, -1, -1, -54, -67, -18,
	-1, -1, -18, -67, 15, 18, 15, 18, -67, -65,
	-18, -28, -15, 58, 6, 8, -31, -20, -9, -75,
	54, 9, -55, -5, 63, -1, -1, -18, -18, -18,
	6, 8, 66, 6, 8, -1, -1, -18, 6, -18,
	-18, -75, -18, -1, -18, -18, -18, -18, -1, -18,
	-75, -75, -18, -1, -18, -69, -18, -1, -1, -1,
	6, -61, 55, -72, 9, -30, 6, 47, 58, -61,
	-54, 40, -54, -65, -18, -5, 11, -5, -18, -4,
	-69, -18, -39, -12, -18, -12, 6, -31, -20, 11,
	-54, -58, 56, -75, 61, -65, -18, -4, -69, -1,
	-1, -18, -1, -1, 6, -31, -20, -1, -1, 6,
	-31, -20, -1, -1, -1, 6, -31, -20, -70, 6,
	6, -54, 51, 52, 51, 52, -1, -64, 11, 51,
	51, -75, -1, -1, -1, -1, 62, 11, 62, -18,
	51, 64, -62, -68, -21, -18, 6, 8, -65, -1,
	52, 10, -75, 6, 8, -71, -57, 42, 9, -18,
	53, -18, -18, 62, 11, 62, -5, -54, 6, 11,
	-72, 42, 6, 6, -54, 14, -34, 14, 10, 11,
	-75, 62, 62, 62, -75, -18, -5, 14, -54, -66,
	6, -58, -54, -18, 10, 62, 62, 62, 62, 61,
	61, 14, -59, 17, 16, -54, -54, 14, -11, 25,
	-18, -63, -63, -35, 37, -75, -75, 11, -75, 11,
	-75, 4, 53, 10, -5, -1, -65, -18, 42, 14,
	-56, -11, -61, -30, 10, -18, 14, -56, -54, -5,
	-75, -18, 58, 42, 11, 58, 14, 56, 11, -54,
	-75, 62, 42, 42, -18, -18, 14, 17, 16, -1,
	-54, 14, -59, 14, -54, 8, 14, 51, 52, -1,
	-41, 15, 18, 14, 14, 16, 37, -67, -18, -15,
	-28, 58, 64, -75, 64, -75, 64, -1, -18, 10,
	-1, 14, -11, -54, 14, 14, 58, 6, -1, 6,
	6, 6, 64, 64, -75, -1, 62, 62, -1, -54,
	-54, 14, -54, 4, -1, -1, -54, -67, -54, -18,
	-18, 6, 14, -56, 6, 61, -1, -54, 6, 51,
	51, 52, 14, -54, -75, 4, 53, 14, 10, -18,
	-54, -63, -63, -63, -75, -1, -18, 62, 14, -42,
	17, 16, 14, -42, 14, -74, 11, -75, 11, 14,
	17, 16, -1, -63, 14, -75, -1, -63, -63, -63,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 49,
	50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
	60, 61, 62, 63, 64, 65, 66, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 0, 0, 0, 20, -2, 22, 23, 24,
	0, 0, 0, 0, 15, 42, 43, 44, 45, 46,
	47, 48, 241, 0, 0, 0, 243, 247, 244, 19,
	25, 26, 107, 13, 0, 71, 228, 0, 0, 0,
	0, 0, 0, 0, 0, 192, 193, 5, 6, 7,
	0, 0, 0, 0, 0, 0, 13, 0, 0, 0,
	0, 0, 0, 0, 0, 13, 13, 0, 0, 0,
	0, 0, 0, 0, 0, 132, 132, 15, -2, 15,
	-2, 74, 83, 107, 0, 0, 0, 103, 116, 117,
	31, 15, 13, 20, -2, 22, 23, 24, 107, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 15, 0, 236, 240, 105,
	0, 0, 13, 242, 0, 0, 0, 0, 0, 0,
	105, 109, 110, 0, 20, -2, 22, 23, 24, 0,
	0, 107, 0, 269, 13, 182, 183, 184, 185, 68,
	166, 167, 0, 164, 165, 215, 223, 67, 76, 84,
	86, 0, 186, 187, 188, 189, 190, 191, 217, 0,
	0, 0, 274, 219, 85, 0, 121, 163, 216, 218,
	80, 15, 0, 130, 132, 133, 135, 0, 0, 15,
	0, 0, 0, 0, 108, 75, 13, 119, 105, 0,
	0, 146, 147, 148, 157, 158, 170, 171, 172, 13,
	0, 15, 210, 15, 107, 0, 121, 0, 0, 149,
	159, 0, 150, 160, 173, 174, 175, 151, 161, 176,
	177, 178, 152, 162, 153, 179, 180, 181, 154, 0,
	0, 0, 15, 15, 16, 17, 18, 0, 0, 253,
	253, 0, 248, 249, 245, 246, 194, 13, 195, 111,
	14, 196, 13, 13, -2, 0, 20, -2, 0, 229,
	230, 231, 15, 168, 169, 77, 78, 0, -2, 100,
	0, 267, 268, 94, 0, 95, 81, 0, 132, 0,
	0, 0, 136, 138, 0, 139, 15, 141, 69, 13,
	0, 87, 90, 92, 0, 122, 123, 205, 0, 0,
	211, 15, 13, 105, 73, 88, 91, 93, 89, 0,
	0, 213, 0, 0, 15, 0, 0, 232, 237, 15,
	106, 0, 0, 0, 0, 0, 0, 13, 0, 13,
	0, 13, 0, 72, 79, 82, 0, 250, 0, 124,
	0, 238, 15, 134, 131, 137, 128, 0, 0, 70,
	0, 118, 0, 0, 0, 0, 206, 209, 0, 0,
	0, 87, 13, 0, 0, 0, 214, 0, 15, 15,
	227, 220, 0, 222, 233, 15, 251, 254, 255, 256,
	257, 0, 0, 252, 270, 15, 0, 15, 112, 113,
	114, 0, 197, 0, 198, 0, 199, 200, 202, 101,
	99, 125, 239, 0, 129, 140, 0, 120, 96, 0,
	104, 212, 207, 208, 0, 98, 0, 156, 15, 225,
	226, 221, 234, 0, 15, 0, 0, 15, 13, 115,
	0, 0, 126, 0, 20, 0, 97, 224, 15, 253,
	15, 15, 271, 13, 272, 13, 0, 127, 102, 0,
	235, 0, 0, 0, 273, 11, 13, 155, 258, 0,
	0, 253, 260, 0, 262, 201, 12, 203, 13, 259,
	0, 253, 253, 266, 261, 204, 253, 264, 265, 263,
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
		//line parser.y:205
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:207
		{
		}
	case 3:
		//line parser.y:209
		{
		}
	case 4:
		//line parser.y:211
		{
		}
	case 5:
		//line parser.y:213
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:215
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:217
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:223
		{
		}
	case 11:
		//line parser.y:225
		{
		}
	case 12:
		//line parser.y:226
		{
		}
	case 13:
		//line parser.y:228
		{
		}
	case 14:
		//line parser.y:229
		{
		}
	case 15:
		//line parser.y:232
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 16:
		//line parser.y:234
		{
		}
	case 17:
		//line parser.y:236
		{
		}
	case 18:
		//line parser.y:238
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 67:
		//line parser.y:250
		{
			RubyVAL.genericValue = ast.RescueModifier{Statement: RubyS[Rubypt-2].genericValue, Rescue: RubyS[Rubypt-0].genericValue}
		}
	case 68:
		//line parser.y:253
		{
			RubyVAL.genericValue = ast.StarSplat{Value: RubyS[Rubypt-0].genericValue}
		}
	case 69:
		//line parser.y:256
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 70:
		//line parser.y:263
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-2].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 71:
		//line parser.y:270
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 72:
		//line parser.y:274
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 73:
		//line parser.y:281
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 74:
		//line parser.y:288
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 75:
		//line parser.y:295
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 76:
		//line parser.y:302
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 77:
		//line parser.y:309
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 78:
		//line parser.y:317
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 79:
		//line parser.y:325
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 80:
		//line parser.y:333
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 81:
		//line parser.y:341
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 82:
		//line parser.y:349
		{
			methodName := RubyS[Rubypt-2].genericValue.(ast.BareReference).Name + "="
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: methodName},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 83:
		//line parser.y:360
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 84:
		//line parser.y:367
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 85:
		//line parser.y:375
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 86:
		//line parser.y:383
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 87:
		//line parser.y:393
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 88:
		//line parser.y:401
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 89:
		//line parser.y:409
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 90:
		//line parser.y:417
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 91:
		//line parser.y:425
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 92:
		//line parser.y:433
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 93:
		//line parser.y:441
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 94:
		//line parser.y:449
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 95:
		//line parser.y:457
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 96:
		//line parser.y:467
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 97:
		//line parser.y:475
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
	case 98:
		//line parser.y:486
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 99:
		//line parser.y:494
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 100:
		//line parser.y:504
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 101:
		//line parser.y:514
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 102:
		//line parser.y:516
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 103:
		//line parser.y:518
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 104:
		//line parser.y:520
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 105:
		//line parser.y:523
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 106:
		//line parser.y:525
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 107:
		//line parser.y:527
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 108:
		//line parser.y:529
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 109:
		//line parser.y:531
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 110:
		//line parser.y:533
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 111:
		//line parser.y:535
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{
				Func:   ast.BareReference{Name: "to_proc"},
				Target: RubyS[Rubypt-0].genericValue,
			})
		}
	case 112:
		//line parser.y:542
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 113:
		//line parser.y:544
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 114:
		//line parser.y:546
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 115:
		//line parser.y:548
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{
				Func:   ast.BareReference{Name: "to_proc"},
				Target: RubyS[Rubypt-0].genericValue,
			})
		}
	case 116:
		//line parser.y:557
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 117:
		//line parser.y:559
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 118:
		//line parser.y:561
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 119:
		//line parser.y:563
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 120:
		//line parser.y:565
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 121:
		//line parser.y:568
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 122:
		//line parser.y:570
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 123:
		//line parser.y:572
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 124:
		//line parser.y:576
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 125:
		//line parser.y:584
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 126:
		//line parser.y:593
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 127:
		//line parser.y:602
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 128:
		//line parser.y:612
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 129:
		//line parser.y:620
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 130:
		//line parser.y:631
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 131:
		//line parser.y:633
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 132:
		//line parser.y:635
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 133:
		//line parser.y:637
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 134:
		//line parser.y:639
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 135:
		//line parser.y:642
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 136:
		//line parser.y:644
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 137:
		//line parser.y:646
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 138:
		//line parser.y:648
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 139:
		//line parser.y:652
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 140:
		//line parser.y:659
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 141:
		//line parser.y:669
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 142:
		//line parser.y:678
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 143:
		//line parser.y:684
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 144:
		//line parser.y:692
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 145:
		//line parser.y:696
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 146:
		//line parser.y:701
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 147:
		//line parser.y:708
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 148:
		//line parser.y:712
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 149:
		//line parser.y:716
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 150:
		//line parser.y:723
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 151:
		//line parser.y:730
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 152:
		//line parser.y:737
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 153:
		//line parser.y:745
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 154:
		//line parser.y:752
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: ast.Array{Nodes: RubyS[Rubypt-2].genericSlice},
				RHS: ast.Array{Nodes: RubyS[Rubypt-0].genericSlice},
			}
		}
	case 155:
		//line parser.y:760
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
	case 156:
		//line parser.y:775
		{
			tail := ast.CallExpression{Target: RubyS[Rubypt-3].genericValue, Func: ast.BareReference{Name: "[]="}, Args: []ast.Node{RubyS[Rubypt-1].genericValue}}
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, tail)
		}
	case 157:
		//line parser.y:781
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 158:
		//line parser.y:788
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 159:
		//line parser.y:792
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 160:
		//line parser.y:799
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 161:
		//line parser.y:806
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 162:
		//line parser.y:813
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 163:
		//line parser.y:820
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 164:
		//line parser.y:823
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 165:
		//line parser.y:825
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 166:
		//line parser.y:828
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 167:
		//line parser.y:830
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 168:
		//line parser.y:833
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 169:
		//line parser.y:835
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 170:
		//line parser.y:838
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 171:
		//line parser.y:840
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 172:
		//line parser.y:842
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 173:
		//line parser.y:845
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 174:
		//line parser.y:847
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 175:
		//line parser.y:849
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 176:
		//line parser.y:852
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 177:
		//line parser.y:854
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 178:
		//line parser.y:856
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 179:
		//line parser.y:859
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 180:
		//line parser.y:861
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 181:
		//line parser.y:863
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 182:
		//line parser.y:865
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 183:
		//line parser.y:866
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 184:
		//line parser.y:867
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 185:
		//line parser.y:868
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 186:
		//line parser.y:871
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 187:
		//line parser.y:880
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 188:
		//line parser.y:889
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 189:
		//line parser.y:898
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 190:
		//line parser.y:907
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 191:
		//line parser.y:916
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 192:
		//line parser.y:924
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 193:
		//line parser.y:925
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 194:
		//line parser.y:927
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 195:
		//line parser.y:929
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 196:
		//line parser.y:932
		{
			RubyVAL.genericValue = ast.Hash{}
		}
	case 197:
		//line parser.y:934
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 198:
		//line parser.y:942
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 199:
		//line parser.y:950
		{
			RubyVAL.genericValue = ast.Block{Body: ast.Nodes{RubyS[Rubypt-2].genericValue}}
		}
	case 200:
		//line parser.y:953
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 201:
		//line parser.y:960
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 202:
		//line parser.y:968
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 203:
		//line parser.y:975
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 204:
		//line parser.y:982
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 205:
		//line parser.y:990
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 206:
		//line parser.y:992
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 207:
		//line parser.y:999
		{
			RubyVAL.genericValue = ast.Block{Args: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 208:
		//line parser.y:1003
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-2].genericSlice}
		}
	case 209:
		//line parser.y:1006
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 210:
		//line parser.y:1008
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 211:
		//line parser.y:1010
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 212:
		//line parser.y:1012
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 213:
		//line parser.y:1015
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 214:
		//line parser.y:1022
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 215:
		//line parser.y:1030
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 216:
		//line parser.y:1037
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 217:
		//line parser.y:1044
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 218:
		//line parser.y:1051
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 219:
		//line parser.y:1058
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 220:
		//line parser.y:1065
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 221:
		//line parser.y:1072
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 222:
		//line parser.y:1080
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 223:
		//line parser.y:1087
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 224:
		//line parser.y:1096
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 225:
		//line parser.y:1103
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 226:
		//line parser.y:1110
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 227:
		//line parser.y:1117
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 228:
		//line parser.y:1124
		{
		}
	case 229:
		//line parser.y:1125
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 230:
		//line parser.y:1126
		{
		}
	case 231:
		//line parser.y:1129
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 232:
		//line parser.y:1132
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 233:
		//line parser.y:1140
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 234:
		//line parser.y:1142
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 235:
		//line parser.y:1151
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
	case 236:
		//line parser.y:1166
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 237:
		//line parser.y:1168
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 238:
		//line parser.y:1171
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 239:
		//line parser.y:1173
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 240:
		//line parser.y:1176
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 241:
		//line parser.y:1183
		{
			RubyVAL.genericValue = ast.Yield{}
		}
	case 242:
		//line parser.y:1186
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 243:
		//line parser.y:1194
		{
			RubyVAL.genericValue = ast.Return{}
		}
	case 244:
		//line parser.y:1198
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 245:
		//line parser.y:1200
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 246:
		//line parser.y:1202
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 247:
		//line parser.y:1206
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 248:
		//line parser.y:1208
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 249:
		//line parser.y:1210
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 250:
		//line parser.y:1214
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 251:
		//line parser.y:1223
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 252:
		//line parser.y:1225
		{
			RubyVAL.genericValue = ast.Loop{Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue}, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 253:
		//line parser.y:1228
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 254:
		//line parser.y:1230
		{
		}
	case 255:
		//line parser.y:1232
		{
		}
	case 256:
		//line parser.y:1234
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 257:
		//line parser.y:1236
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 258:
		//line parser.y:1239
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 259:
		//line parser.y:1246
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 260:
		//line parser.y:1254
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 261:
		//line parser.y:1261
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 262:
		//line parser.y:1269
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 263:
		//line parser.y:1277
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 264:
		//line parser.y:1284
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 265:
		//line parser.y:1291
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 266:
		//line parser.y:1298
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 267:
		//line parser.y:1306
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 268:
		//line parser.y:1309
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 269:
		//line parser.y:1311
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 270:
		//line parser.y:1314
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 271:
		//line parser.y:1316
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 272:
		//line parser.y:1319
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 273:
		//line parser.y:1321
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 274:
		//line parser.y:1323
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
