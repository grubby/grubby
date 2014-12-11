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

//line parser.y:1323

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
	-2, 227,
	-1, 134,
	54, 144,
	-2, 21,
	-1, 174,
	54, 144,
	-2, 21,
	-1, 301,
	51, 13,
	64, 13,
	-2, 31,
	-1, 304,
	54, 145,
	-2, 143,
	-1, 315,
	10, 107,
	11, 107,
	-2, 227,
}

const RubyNprod = 274
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 3388

var RubyAct = []int{

	284, 158, 5, 387, 506, 388, 220, 224, 359, 122,
	168, 30, 25, 222, 18, 51, 121, 2, 3, 26,
	109, 353, 127, 339, 348, 90, 273, 297, 91, 189,
	96, 190, 129, 267, 4, 61, 262, 244, 297, 297,
	460, 110, 482, 321, 297, 357, 234, 131, 250, 114,
	131, 443, 441, 155, 156, 117, 119, 439, 356, 105,
	106, 88, 87, 142, 160, 161, 94, 95, 163, 321,
	321, 97, 98, 99, 107, 294, 143, 145, 89, 184,
	185, 93, 102, 100, 101, 104, 83, 170, 408, 191,
	179, 194, 195, 83, 354, 142, 83, 83, 183, 202,
	13, 183, 297, 200, 207, 138, 138, 250, 143, 212,
	182, 297, 216, 217, 218, 144, 493, 368, 402, 297,
	340, 320, 371, 228, 400, 149, 295, 142, 410, 405,
	47, 232, 150, 214, 96, 170, 297, 240, 170, 140,
	409, 257, 258, 237, 260, 261, 238, 265, 266, 253,
	270, 271, 272, 170, 233, 235, 249, 139, 139, 255,
	241, 243, 256, 105, 106, 290, 291, 292, 293, 276,
	94, 95, 90, 171, 404, 91, 135, 178, 107, 385,
	90, 112, 306, 91, 113, 93, 328, 401, 149, 154,
	146, 305, 352, 135, 230, 170, 135, 135, 90, 108,
	286, 91, 478, 175, 446, 336, 109, 312, 486, 209,
	210, 135, 135, 135, 313, 111, 487, 488, 400, 225,
	153, 171, 223, 135, 171, 135, 135, 110, 135, 401,
	135, 135, 135, 135, 288, 135, 251, 327, 135, 171,
	135, 135, 333, 376, 138, 225, 90, 297, 223, 91,
	494, 175, 323, 374, 175, 135, 135, 245, 152, 90,
	226, 366, 91, 96, 135, 96, 289, 170, 221, 175,
	135, 227, 237, 135, 148, 238, 263, 140, 154, 268,
	343, 171, 287, 274, 141, 166, 226, 309, 167, 151,
	326, 365, 105, 106, 105, 106, 139, 227, 247, 94,
	95, 94, 95, 135, 135, 146, 97, 98, 99, 135,
	180, 175, 147, 451, 93, 382, 93, 102, 100, 101,
	104, 504, 225, 422, 366, 431, 383, 432, 135, 96,
	170, 135, 389, 448, 390, 394, 364, 118, 337, 90,
	135, 135, 91, 304, 366, 235, 381, 366, 433, 164,
	92, 341, 165, 171, 343, 391, 326, 485, 105, 106,
	513, 416, 458, 226, 90, 94, 95, 91, 396, 426,
	426, 419, 457, 434, 227, 280, 281, 456, 135, 444,
	93, 470, 135, 175, 104, 454, 447, 437, 347, 521,
	96, 518, 517, 449, 516, 330, 518, 517, 372, 329,
	449, 455, 325, 373, 375, 377, 369, 96, 380, 294,
	278, 462, 437, 512, 277, 465, 171, 135, 219, 105,
	106, 468, 197, 415, 414, 74, 94, 95, 300, 471,
	472, 351, 294, 346, 135, 474, 105, 106, 335, 336,
	397, 93, 285, 94, 95, 104, 175, 299, 135, 48,
	1, 181, 135, 407, 480, 492, 60, 59, 93, 135,
	96, 413, 483, 415, 414, 464, 58, 515, 135, 310,
	192, 311, 193, 436, 116, 57, 115, 56, 440, 55,
	442, 38, 37, 36, 105, 106, 449, 135, 135, 105,
	106, 94, 95, 502, 35, 136, 94, 95, 436, 426,
	426, 426, 135, 175, 510, 50, 93, 427, 519, 20,
	135, 93, 136, 461, 40, 136, 136, 41, 523, 21,
	370, 426, 176, 15, 12, 426, 426, 426, 175, 11,
	136, 136, 136, 24, 23, 22, 19, 10, 32, 27,
	17, 14, 136, 39, 136, 136, 16, 136, 34, 136,
	136, 136, 136, 33, 136, 28, 71, 136, 29, 136,
	136, 70, 0, 0, 135, 0, 0, 0, 0, 135,
	176, 135, 0, 176, 136, 136, 246, 0, 0, 491,
	0, 0, 0, 136, 135, 0, 0, 0, 176, 136,
	0, 0, 136, 0, 501, 264, 0, 0, 269, 0,
	0, 0, 275, 0, 498, 499, 500, 514, 0, 0,
	0, 0, 135, 135, 0, 0, 96, 31, 0, 522,
	0, 0, 136, 136, 135, 0, 520, 0, 136, 0,
	176, 69, 133, 75, 134, 120, 524, 525, 131, 0,
	0, 526, 0, 0, 0, 105, 106, 136, 0, 0,
	136, 0, 94, 95, 0, 0, 0, 0, 0, 136,
	136, 77, 0, 130, 85, 86, 0, 93, 0, 78,
	79, 80, 81, 82, 463, 0, 0, 0, 0, 379,
	130, 0, 0, 130, 130, 0, 0, 252, 0, 132,
	130, 84, 83, 0, 96, 0, 0, 136, 130, 130,
	130, 136, 176, 96, 0, 0, 0, 0, 0, 0,
	130, 0, 130, 130, 0, 130, 0, 130, 130, 130,
	130, 0, 130, 105, 106, 130, 0, 130, 130, 0,
	94, 95, 105, 106, 0, 0, 136, 0, 130, 94,
	95, 130, 130, 130, 96, 93, 96, 0, 0, 0,
	0, 130, 355, 136, 93, 0, 130, 130, 0, 0,
	130, 322, 0, 0, 0, 176, 0, 136, 0, 0,
	0, 136, 378, 105, 106, 105, 106, 0, 136, 0,
	94, 95, 94, 95, 0, 0, 0, 136, 0, 0,
	130, 130, 0, 317, 0, 93, 301, 93, 130, 0,
	0, 105, 106, 0, 0, 0, 136, 136, 94, 95,
	0, 0, 0, 0, 0, 130, 0, 0, 130, 0,
	0, 136, 176, 93, 0, 0, 0, 130, 130, 136,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 133,
	75, 134, 120, 0, 126, 131, 0, 176, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 130, 0, 0, 77, 301,
	130, 85, 86, 0, 0, 124, 78, 79, 80, 81,
	82, 0, 125, 136, 0, 0, 0, 0, 136, 0,
	136, 0, 0, 0, 123, 0, 132, 0, 84, 83,
	0, 0, 0, 136, 130, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 130, 0, 69, 133, 75, 134, 120, 495, 0,
	131, 136, 136, 130, 0, 130, 9, 0, 0, 130,
	0, 0, 0, 136, 0, 0, 130, 0, 0, 0,
	0, 0, 0, 77, 0, 130, 85, 86, 0, 0,
	0, 78, 79, 80, 81, 82, 0, 0, 0, 279,
	0, 0, 0, 0, 130, 130, 0, 0, 0, 252,
	0, 132, 128, 84, 83, 0, 0, 0, 0, 130,
	130, 0, 0, 0, 0, 0, 0, 130, 0, 159,
	0, 0, 162, 159, 0, 0, 0, 0, 0, 169,
	0, 0, 0, 0, 0, 130, 0, 186, 187, 188,
	0, 0, 0, 0, 157, 0, 0, 0, 0, 196,
	0, 198, 199, 0, 201, 0, 203, 204, 205, 206,
	0, 208, 0, 0, 211, 0, 213, 215, 0, 0,
	0, 130, 0, 0, 0, 0, 130, 169, 130, 0,
	236, 239, 242, 0, 0, 0, 0, 0, 0, 0,
	128, 130, 0, 0, 0, 169, 254, 96, 0, 259,
	0, 0, 0, 0, 0, 0, 0, 229, 0, 231,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 130,
	130, 248, 0, 0, 49, 0, 105, 106, 0, 296,
	128, 130, 0, 94, 95, 302, 0, 169, 97, 98,
	99, 107, 69, 303, 75, 134, 76, 0, 93, 102,
	100, 101, 104, 0, 128, 338, 0, 316, 0, 0,
	0, 0, 0, 0, 0, 0, 318, 319, 0, 0,
	137, 0, 77, 0, 0, 85, 86, 0, 0, 0,
	78, 79, 80, 81, 82, 0, 0, 137, 297, 0,
	137, 137, 0, 250, 0, 0, 0, 177, 72, 0,
	73, 298, 84, 83, 342, 137, 137, 137, 302, 350,
	324, 0, 0, 0, 0, 0, 0, 137, 331, 137,
	137, 0, 137, 0, 137, 137, 137, 137, 0, 137,
	0, 0, 137, 0, 137, 137, 0, 0, 0, 345,
	0, 349, 0, 367, 0, 177, 0, 0, 177, 137,
	137, 0, 0, 0, 0, 0, 0, 0, 137, 0,
	128, 0, 0, 177, 137, 0, 0, 137, 0, 0,
	362, 363, 169, 0, 384, 0, 0, 0, 342, 0,
	0, 0, 0, 0, 0, 392, 0, 0, 0, 0,
	0, 0, 0, 0, 398, 0, 0, 137, 137, 349,
	0, 0, 0, 137, 0, 177, 0, 69, 173, 75,
	174, 76, 0, 411, 412, 0, 0, 0, 0, 0,
	0, 0, 137, 395, 0, 137, 0, 0, 159, 435,
	0, 0, 0, 0, 137, 137, 445, 77, 406, 0,
	85, 86, 0, 0, 0, 78, 79, 80, 81, 82,
	0, 417, 0, 297, 435, 0, 421, 0, 0, 0,
	453, 0, 0, 72, 0, 73, 0, 84, 83, 0,
	0, 0, 137, 0, 0, 0, 137, 177, 0, 450,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	159, 0, 0, 0, 0, 476, 0, 477, 0, 0,
	0, 0, 0, 0, 0, 466, 467, 0, 0, 0,
	476, 137, 469, 0, 0, 96, 0, 0, 0, 0,
	0, 0, 473, 0, 475, 0, 0, 0, 137, 103,
	0, 0, 0, 0, 0, 0, 92, 0, 128, 496,
	177, 0, 137, 0, 105, 106, 137, 0, 0, 0,
	503, 94, 95, 137, 0, 484, 97, 98, 99, 107,
	0, 0, 137, 0, 490, 0, 93, 102, 100, 101,
	104, 0, 0, 0, 0, 497, 0, 362, 363, 0,
	0, 137, 137, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 137, 177, 0, 0,
	0, 0, 0, 0, 137, 0, 69, 45, 75, 46,
	76, 0, 0, 0, 42, 509, 428, 508, 507, 429,
	43, 44, 177, 63, 64, 54, 0, 0, 67, 68,
	0, 0, 66, 62, 0, 0, 77, 65, 0, 85,
	86, 0, 0, 0, 78, 79, 80, 81, 82, 0,
	0, 0, 424, 425, 0, 0, 0, 0, 137, 0,
	0, 0, 72, 137, 73, 137, 84, 83, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 137, 0,
	0, 0, 69, 45, 75, 46, 76, 0, 0, 0,
	42, 505, 428, 508, 507, 429, 43, 44, 0, 63,
	64, 54, 0, 0, 67, 68, 137, 137, 66, 62,
	0, 0, 77, 65, 0, 85, 86, 0, 137, 0,
	78, 79, 80, 81, 82, 0, 0, 0, 424, 425,
	0, 0, 0, 0, 0, 0, 0, 0, 72, 0,
	73, 0, 84, 83, 69, 45, 75, 46, 76, 0,
	0, 0, 42, 418, 52, 361, 360, 53, 43, 44,
	0, 63, 64, 54, 0, 0, 67, 68, 0, 0,
	66, 62, 0, 0, 77, 65, 0, 85, 86, 0,
	0, 0, 78, 79, 80, 81, 82, 0, 0, 0,
	282, 283, 0, 0, 0, 0, 0, 0, 0, 0,
	72, 0, 73, 0, 84, 83, 69, 45, 75, 46,
	76, 0, 0, 0, 42, 358, 52, 361, 360, 53,
	43, 44, 0, 63, 64, 54, 0, 0, 67, 68,
	0, 0, 66, 62, 0, 0, 77, 65, 0, 85,
	86, 0, 0, 0, 78, 79, 80, 81, 82, 0,
	0, 0, 282, 283, 0, 69, 45, 75, 46, 76,
	0, 0, 72, 42, 73, 52, 84, 83, 53, 43,
	44, 0, 63, 64, 54, 0, 0, 67, 68, 0,
	0, 66, 62, 0, 0, 77, 65, 0, 85, 86,
	0, 0, 0, 78, 79, 80, 81, 82, 0, 0,
	0, 6, 7, 0, 0, 0, 0, 0, 0, 0,
	0, 72, 0, 73, 0, 84, 83, 0, 8, 69,
	45, 75, 46, 76, 0, 0, 0, 42, 479, 52,
	0, 0, 53, 43, 44, 0, 63, 64, 54, 366,
	0, 67, 68, 0, 0, 66, 62, 0, 0, 77,
	65, 0, 85, 86, 0, 0, 0, 78, 79, 80,
	81, 82, 0, 0, 0, 282, 283, 0, 0, 0,
	0, 0, 0, 0, 0, 72, 0, 73, 0, 84,
	83, 69, 45, 75, 46, 76, 0, 0, 0, 42,
	393, 52, 0, 0, 53, 43, 44, 0, 63, 64,
	54, 366, 0, 67, 68, 0, 0, 66, 62, 0,
	0, 77, 65, 0, 85, 86, 0, 0, 0, 78,
	79, 80, 81, 82, 0, 0, 0, 282, 283, 0,
	0, 0, 0, 0, 0, 0, 0, 72, 0, 73,
	0, 84, 83, 69, 45, 75, 46, 76, 0, 0,
	0, 42, 386, 52, 0, 0, 53, 43, 44, 0,
	63, 64, 54, 366, 0, 67, 68, 0, 0, 66,
	62, 0, 0, 77, 65, 0, 85, 86, 0, 0,
	0, 78, 79, 80, 81, 82, 0, 0, 0, 282,
	283, 0, 0, 0, 0, 0, 0, 0, 0, 72,
	0, 73, 0, 84, 83, 69, 45, 75, 46, 76,
	0, 0, 0, 42, 511, 428, 0, 0, 429, 43,
	44, 0, 63, 64, 54, 0, 0, 67, 68, 0,
	0, 66, 62, 0, 0, 77, 65, 0, 85, 86,
	0, 0, 0, 78, 79, 80, 81, 82, 0, 0,
	0, 424, 425, 0, 0, 0, 0, 0, 0, 0,
	0, 72, 0, 73, 0, 84, 83, 69, 45, 75,
	46, 76, 0, 0, 0, 42, 489, 52, 0, 0,
	53, 43, 44, 0, 63, 64, 54, 0, 0, 67,
	68, 0, 0, 66, 62, 0, 0, 77, 65, 0,
	85, 86, 0, 0, 0, 78, 79, 80, 81, 82,
	0, 0, 0, 282, 283, 0, 69, 45, 75, 46,
	76, 0, 0, 72, 42, 73, 52, 84, 83, 53,
	43, 44, 0, 63, 64, 54, 0, 0, 67, 68,
	0, 0, 66, 62, 0, 0, 77, 65, 0, 85,
	86, 0, 0, 0, 78, 79, 80, 81, 82, 0,
	0, 0, 282, 283, 0, 0, 0, 0, 0, 0,
	0, 0, 72, 0, 73, 459, 84, 83, 69, 45,
	75, 46, 76, 0, 0, 0, 42, 452, 52, 0,
	0, 53, 43, 44, 0, 63, 64, 54, 0, 0,
	67, 68, 0, 0, 66, 62, 0, 0, 77, 65,
	0, 85, 86, 0, 0, 0, 78, 79, 80, 81,
	82, 0, 0, 0, 282, 283, 0, 0, 0, 0,
	0, 0, 0, 0, 72, 0, 73, 0, 84, 83,
	69, 45, 75, 46, 76, 0, 0, 0, 42, 430,
	428, 0, 0, 429, 43, 44, 0, 63, 64, 54,
	0, 0, 67, 68, 0, 0, 66, 62, 0, 0,
	77, 65, 0, 85, 86, 0, 0, 0, 78, 79,
	80, 81, 82, 0, 0, 0, 424, 425, 0, 0,
	0, 0, 0, 0, 0, 0, 72, 0, 73, 0,
	84, 83, 69, 45, 75, 46, 76, 0, 0, 0,
	42, 423, 428, 0, 0, 429, 43, 44, 0, 63,
	64, 54, 0, 0, 67, 68, 0, 0, 66, 62,
	0, 0, 77, 65, 0, 85, 86, 0, 0, 0,
	78, 79, 80, 81, 82, 0, 0, 0, 424, 425,
	0, 0, 0, 0, 0, 0, 0, 0, 72, 0,
	73, 0, 84, 83, 69, 45, 75, 46, 76, 0,
	0, 0, 42, 420, 52, 0, 0, 53, 43, 44,
	0, 63, 64, 54, 0, 0, 67, 68, 0, 0,
	66, 62, 0, 0, 77, 65, 0, 85, 86, 0,
	0, 0, 78, 79, 80, 81, 82, 0, 0, 0,
	282, 283, 0, 0, 0, 0, 0, 0, 0, 0,
	72, 0, 73, 0, 84, 83, 69, 45, 75, 46,
	76, 0, 0, 0, 42, 403, 52, 0, 0, 53,
	43, 44, 0, 63, 64, 54, 0, 0, 67, 68,
	0, 0, 66, 62, 0, 0, 77, 65, 0, 85,
	86, 0, 0, 0, 78, 79, 80, 81, 82, 0,
	0, 0, 282, 283, 0, 0, 0, 0, 0, 0,
	0, 0, 72, 0, 73, 0, 84, 83, 69, 45,
	75, 46, 76, 0, 0, 0, 42, 344, 52, 0,
	0, 53, 43, 44, 0, 63, 64, 54, 0, 0,
	67, 68, 0, 0, 66, 62, 0, 0, 77, 65,
	0, 85, 86, 0, 0, 0, 78, 79, 80, 81,
	82, 0, 0, 0, 282, 283, 0, 0, 0, 0,
	0, 0, 0, 0, 72, 0, 73, 0, 84, 83,
	69, 45, 75, 46, 76, 0, 0, 0, 42, 334,
	52, 0, 0, 53, 43, 44, 0, 63, 64, 54,
	0, 0, 67, 68, 0, 0, 66, 62, 0, 0,
	77, 65, 0, 85, 86, 0, 0, 0, 78, 79,
	80, 81, 82, 0, 0, 0, 282, 283, 0, 0,
	0, 0, 0, 0, 0, 0, 72, 0, 73, 0,
	84, 83, 69, 45, 75, 46, 76, 0, 0, 0,
	42, 332, 52, 0, 0, 53, 43, 44, 0, 63,
	64, 54, 0, 0, 67, 68, 0, 0, 66, 62,
	0, 0, 77, 65, 0, 85, 86, 0, 0, 0,
	78, 79, 80, 81, 82, 0, 0, 0, 282, 283,
	0, 69, 45, 75, 46, 76, 0, 0, 72, 42,
	73, 428, 84, 83, 429, 43, 44, 0, 63, 64,
	54, 0, 0, 67, 68, 0, 0, 66, 62, 0,
	0, 77, 65, 0, 85, 86, 0, 0, 0, 78,
	79, 80, 81, 82, 0, 0, 0, 424, 425, 0,
	69, 45, 75, 46, 76, 0, 0, 72, 42, 73,
	52, 84, 83, 53, 43, 44, 0, 63, 64, 54,
	0, 0, 67, 68, 0, 0, 66, 62, 0, 0,
	77, 65, 0, 85, 86, 0, 0, 0, 78, 79,
	80, 81, 82, 0, 0, 0, 282, 283, 0, 69,
	45, 75, 46, 76, 308, 0, 72, 42, 73, 52,
	84, 83, 53, 43, 44, 0, 63, 64, 54, 0,
	0, 67, 68, 0, 0, 66, 62, 0, 0, 77,
	65, 0, 85, 86, 0, 0, 0, 78, 79, 80,
	81, 82, 0, 0, 0, 0, 307, 0, 69, 45,
	75, 46, 76, 0, 0, 72, 42, 73, 52, 84,
	83, 53, 43, 44, 0, 63, 64, 54, 0, 0,
	67, 68, 0, 0, 66, 62, 0, 0, 77, 65,
	0, 85, 86, 0, 0, 0, 78, 79, 80, 81,
	82, 0, 0, 0, 297, 0, 0, 69, 45, 75,
	46, 76, 0, 0, 72, 42, 73, 52, 84, 83,
	53, 43, 44, 0, 63, 64, 54, 0, 0, 67,
	68, 0, 0, 66, 62, 0, 0, 77, 65, 0,
	85, 86, 0, 0, 0, 78, 79, 80, 81, 82,
	69, 173, 75, 174, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 72, 0, 73, 0, 84, 83, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	77, 0, 0, 85, 86, 0, 0, 0, 78, 79,
	80, 81, 82, 0, 0, 0, 297, 69, 133, 75,
	134, 76, 0, 438, 0, 0, 72, 0, 73, 0,
	84, 83, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 77, 0, 0,
	85, 86, 0, 0, 0, 78, 79, 80, 81, 82,
	0, 0, 0, 297, 69, 133, 75, 134, 315, 0,
	399, 131, 0, 72, 0, 73, 0, 84, 83, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 77, 0, 0, 85, 86, 0,
	0, 314, 78, 79, 80, 81, 82, 69, 303, 75,
	134, 76, 0, 0, 0, 0, 0, 0, 0, 0,
	72, 0, 132, 0, 84, 83, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 77, 0, 0,
	85, 86, 0, 0, 0, 78, 79, 80, 81, 82,
	0, 0, 0, 297, 69, 133, 75, 134, 120, 0,
	0, 131, 0, 72, 0, 73, 298, 84, 83, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 77, 0, 0, 85, 86, 0,
	0, 124, 78, 79, 80, 81, 82, 69, 133, 75,
	134, 76, 0, 0, 0, 0, 0, 0, 0, 0,
	252, 0, 132, 0, 84, 83, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 77, 0, 0,
	85, 86, 0, 0, 0, 78, 79, 80, 81, 82,
	0, 0, 0, 297, 69, 133, 75, 134, 76, 0,
	0, 131, 0, 72, 0, 73, 0, 84, 83, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 77, 0, 0, 85, 86, 0,
	0, 0, 78, 79, 80, 81, 82, 69, 173, 75,
	174, 76, 0, 0, 0, 0, 0, 0, 0, 0,
	72, 0, 132, 0, 84, 83, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 77, 0, 0,
	85, 86, 0, 0, 0, 78, 79, 80, 81, 82,
	69, 133, 75, 134, 120, 0, 0, 131, 0, 0,
	172, 0, 0, 72, 0, 73, 0, 84, 83, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	77, 0, 0, 85, 86, 0, 0, 0, 78, 79,
	80, 81, 82, 69, 133, 75, 134, 76, 0, 0,
	0, 0, 0, 0, 0, 0, 252, 0, 132, 0,
	84, 83, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 77, 0, 0, 85, 86, 0, 0,
	0, 78, 79, 80, 81, 82, 69, 481, 75, 134,
	76, 0, 0, 0, 0, 0, 0, 0, 0, 72,
	0, 73, 0, 84, 83, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 77, 0, 0, 85,
	86, 0, 0, 0, 78, 79, 80, 81, 82, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 72, 0, 73, 0, 84, 83,
}
var RubyPact = []int{

	-34, 1730, -1000, -1000, -1000, 10, -1000, -1000, -1000, 1391,
	-1000, -1000, -1000, 181, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 166, -1000, -6, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 470, 329, 329, 833, 235, 66, 263, 83,
	247, 178, 2832, 2832, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 3278, 2832, 2832, 3278, 3278, 334, 270, -1000,
	-1000, -1000, 3192, -1000, 36, 301, -1000, 38, 2832, 2832,
	3278, 3278, 3278, 23, 464, -1000, -1000, -1000, -1000, -1000,
	2832, 2832, 3278, 416, 3278, 3278, -1000, 3278, 2832, 3278,
	3278, 3278, 3278, 2832, 3278, -1000, -1000, 3278, 2832, 3278,
	3278, 2832, 2832, 2832, 412, 213, 239, 154, -1000, -1000,
	3192, 38, 35, 3192, 3278, 3278, 31, 287, 742, -1000,
	-20, -8, -1000, 3235, 96, 2, -1000, -1000, 3192, 3278,
	2832, 2832, 3278, 2832, 2832, 30, 2832, 2832, 27, 2832,
	2832, 2832, 20, 408, 404, 244, 324, 2685, 189, 742,
	231, 183, 742, 189, 2832, 2832, 2832, 2832, 64, 261,
	-1000, -1000, 3278, 3059, 97, 34, 148, 146, 3012, 335,
	3192, 2734, -1000, -1000, 244, 244, 742, 742, 742, -1000,
	-1000, 463, -1000, -1000, 244, 244, 742, 2969, 742, 742,
	3102, 742, 244, 742, 742, 742, 742, 244, 740, 3102,
	3102, 742, 244, 742, 59, 699, 244, 244, 244, 38,
	-1000, 396, 279, 316, -1000, 144, 393, 389, -1000, 2587,
	329, 2525, 428, -1000, -1000, -1000, 1073, -39, 58, 325,
	-1000, -1000, 386, -1000, -1000, -1000, -1000, 3149, 2463, -1000,
	382, 1117, 3192, 421, 130, -41, 32, 244, 244, 690,
	244, 244, -1000, -1000, -1000, 244, 244, -1000, -1000, -1000,
	244, 244, 244, -1000, -1000, -1000, 267, -3, -16, 1681,
	-1000, -1000, -1000, -1000, 244, 322, 3278, -1000, -1000, 85,
	244, 244, 244, 244, -1000, -1000, 742, -1000, -1000, 242,
	232, -20, 768, 626, -1000, 398, 244, -1000, -1000, 51,
	-1000, -1000, 35, 38, 2832, 3192, 742, 3278, 742, 742,
	-1000, 3149, 137, -1000, 1918, 239, 316, 345, 3278, -1000,
	-1000, 1856, -1000, -1000, -1000, 38, -1000, 2922, 176, -1000,
	-1000, 60, 742, -1000, -1000, 2401, 118, -1000, -1000, 2685,
	26, -1000, 98, -1000, -1000, 86, 3278, 3278, -1000, 447,
	2832, -1000, 1619, 2339, -1000, -1000, 315, 742, 2277, 2215,
	311, 3278, 2875, -7, -1000, -12, -1000, -13, 2832, 3278,
	-1000, -1000, 244, 194, 742, 2832, -1000, 319, -1000, -1000,
	-1000, -1000, 742, -1000, 299, 2153, -1000, 1282, 742, 379,
	2832, 371, 366, -1000, -1000, 356, 2091, -24, 82, -1000,
	2832, 612, 403, -1000, 2832, -1000, 244, 2685, -1000, 407,
	-1000, 2685, 377, -1000, -1000, -1000, 244, -1000, 2832, 2832,
	-1000, -1000, -1000, 3278, 189, 261, -1000, -1000, 3278, -1000,
	3102, -1000, 196, -1000, 244, 742, -1000, 244, -1000, -1000,
	1794, -1000, -1000, 3321, -1000, 244, -19, -1000, -1000, -1000,
	-1000, 2783, 244, 218, -1000, 244, 2685, 2685, -1000, 2685,
	351, 157, 165, 2042, 189, 2685, 742, 451, 63, -1000,
	236, 918, 3278, 244, 2685, -1000, -1000, -1000, -1000, -1000,
	2685, 68, 2832, 3278, -1000, -1000, 259, 2685, 1557, 1481,
	1980, 68, 349, 456, -1000, -1000, 380, 2832, -1000, -1000,
	375, -1000, -1000, -1000, 68, -1000, -1000, 2832, -1000, 244,
	2636, -1000, 68, 244, 2636, 2636, 2636,
}
var RubyPgo = []int{

	0, 0, 561, 558, 19, 32, 556, 555, 553, 1104,
	548, 5, 35, 546, 543, 100, 541, 540, 936, 539,
	449, 617, 538, 537, 536, 535, 534, 533, 14, 529,
	7, 130, 524, 523, 11, 520, 519, 517, 514, 12,
	509, 507, 4, 505, 494, 483, 482, 481, 479, 477,
	475, 466, 457, 456, 969, 451, 3, 16, 24, 8,
	450, 6, 447, 117, 442, 10, 433, 1, 428, 22,
	15, 9, 13, 425, 413, 103,
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
	52, 53, 2, 6, 7, 3, 3, 3, 3, 62,
	62, 68, 68, 68, 5, 5, 5, 5, 58, 66,
	66, 66, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 59, 59, 59, 59, 55, 55, 55,
	8, 17, 11, 11, 11, 64, 64, 56, 56, 24,
	24, 25, 25, 27, 27, 27, 26, 26, 26, 12,
	40, 40, 63, 63, 63, 63, 63, 41, 41, 41,
	41, 41, 42, 42, 42, 42, 38, 37, 10, 36,
	36, 35, 35, 4,
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
	3, 3, 1, 1, 3, 3, 5, 5, 5, 3,
	7, 3, 7, 8, 3, 4, 5, 5, 3, 0,
	1, 3, 4, 5, 3, 3, 3, 3, 3, 5,
	6, 5, 3, 4, 3, 3, 2, 0, 2, 2,
	3, 4, 2, 3, 5, 0, 2, 1, 2, 2,
	1, 2, 1, 1, 3, 3, 1, 3, 3, 5,
	5, 5, 0, 2, 2, 2, 2, 5, 6, 5,
	6, 5, 4, 3, 3, 2, 4, 4, 2, 5,
	7, 4, 5, 3,
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
	-1, -1, -18, -67, 15, 18, 15, 18, -65, -18,
	-28, -15, 58, 6, 8, -31, -20, -9, -75, 54,
	9, -55, -5, 63, -1, -1, -18, -18, -18, 6,
	8, 66, 6, 8, -1, -1, -18, 6, -18, -18,
	-75, -18, -1, -18, -18, -18, -18, -1, -18, -75,
	-75, -18, -1, -18, -69, -18, -1, -1, -1, 6,
	-61, 55, -72, 9, -30, 6, 47, 58, -61, -54,
	40, -54, -65, -5, 11, -5, -18, -4, -69, -18,
	-39, -12, -18, -12, 6, -31, -20, 11, -54, -58,
	56, -75, 61, -65, -18, -4, -69, -1, -1, -18,
	-1, -1, 6, -31, -20, -1, -1, 6, -31, -20,
	-1, -1, -1, 6, -31, -20, -70, 6, 6, -54,
	51, 52, 51, 52, -1, -64, 11, 51, 51, -75,
	-1, -1, -1, -1, 11, 62, -18, 51, 64, -62,
	-68, -21, -18, 6, 8, -65, -1, 52, 10, -75,
	6, 8, -71, -57, 42, 9, -18, 53, -18, -18,
	62, 11, 62, -5, -54, 6, 11, -72, 42, 6,
	6, -54, 14, -34, 14, 10, 11, -75, 62, 62,
	62, -75, -18, -5, 14, -54, -66, 6, -58, -54,
	-18, 10, 62, 62, 62, 62, 61, 61, 14, -59,
	17, 16, -54, -54, 14, -11, 25, -18, -63, -63,
	-35, 37, -75, -75, 11, -75, 11, -75, 4, 53,
	10, -5, -1, -65, -18, 42, 14, -56, -11, -61,
	-30, 10, -18, 14, -56, -54, -5, -75, -18, 58,
	42, 11, 58, 14, 56, 11, -54, -75, 62, 42,
	42, -18, -18, 14, 17, 16, -1, -54, 14, -59,
	14, -54, 8, 14, 51, 52, -1, -41, 15, 18,
	14, 14, 16, 37, -67, -18, -15, -28, 58, 64,
	-75, 64, -75, 64, -1, -18, 10, -1, 14, -11,
	-54, 14, 14, 58, 6, -1, 6, 6, 6, 64,
	64, -75, -1, 62, 62, -1, -54, -54, 14, -54,
	4, -1, -1, -54, -67, -54, -18, -18, 6, 14,
	-56, 6, 61, -1, -54, 6, 51, 51, 52, 14,
	-54, -75, 4, 53, 14, 10, -18, -54, -63, -63,
	-63, -75, -1, -18, 62, 14, -42, 17, 16, 14,
	-42, 14, -74, 11, -75, 11, 14, 17, 16, -1,
	-63, 14, -75, -1, -63, -63, -63,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 49,
	50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
	60, 61, 62, 63, 64, 65, 66, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 0, 0, 0, 20, -2, 22, 23, 24,
	0, 0, 0, 0, 15, 42, 43, 44, 45, 46,
	47, 48, 240, 0, 0, 0, 242, 246, 243, 19,
	25, 26, 107, 13, 0, 71, 227, 0, 0, 0,
	0, 0, 0, 0, 0, 192, 193, 5, 6, 7,
	0, 0, 0, 0, 0, 0, 13, 0, 0, 0,
	0, 0, 0, 0, 0, 13, 13, 0, 0, 0,
	0, 0, 0, 0, 0, 132, 132, 15, -2, 15,
	-2, 74, 83, 107, 0, 0, 0, 103, 116, 117,
	31, 15, 13, 20, -2, 22, 23, 24, 107, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 15, 0, 235, 239, 105,
	0, 0, 13, 241, 0, 0, 0, 0, 0, 108,
	109, 110, 0, 20, -2, 22, 23, 24, 0, 0,
	107, 0, 268, 13, 182, 183, 184, 185, 68, 166,
	167, 0, 164, 165, 214, 222, 67, 76, 84, 86,
	0, 186, 187, 188, 189, 190, 191, 216, 0, 0,
	0, 273, 218, 85, 0, 121, 163, 215, 217, 80,
	15, 0, 130, 132, 133, 135, 0, 0, 15, 0,
	0, 0, 0, 75, 13, 119, 108, 0, 0, 146,
	147, 148, 157, 158, 170, 171, 172, 13, 0, 15,
	209, 15, 107, 0, 121, 0, 0, 149, 159, 0,
	150, 160, 173, 174, 175, 151, 161, 176, 177, 178,
	152, 162, 153, 179, 180, 181, 154, 0, 0, 0,
	15, 15, 16, 17, 18, 0, 0, 252, 252, 0,
	247, 248, 244, 245, 13, 194, 111, 14, 195, 13,
	13, -2, 0, 20, -2, 0, 228, 229, 230, 15,
	168, 169, 77, 78, 0, -2, 100, 0, 266, 267,
	94, 0, 95, 81, 0, 132, 0, 0, 0, 136,
	138, 0, 139, 15, 141, 69, 13, 0, 87, 90,
	92, 0, 122, 123, 204, 0, 0, 210, 15, 13,
	108, 73, 88, 91, 93, 89, 0, 0, 212, 0,
	0, 15, 0, 0, 231, 236, 15, 106, 0, 0,
	0, 0, 0, 0, 13, 0, 13, 0, 13, 0,
	72, 79, 82, 0, 249, 0, 124, 0, 237, 15,
	134, 131, 137, 128, 0, 0, 70, 0, 118, 0,
	0, 0, 0, 205, 208, 0, 0, 0, 87, 13,
	0, 0, 0, 213, 0, 15, 15, 226, 219, 0,
	221, 232, 15, 250, 253, 254, 255, 256, 0, 0,
	251, 269, 15, 0, 15, 112, 113, 114, 0, 196,
	0, 197, 0, 198, 199, 201, 101, 99, 125, 238,
	0, 129, 140, 0, 120, 96, 0, 104, 211, 206,
	207, 0, 98, 0, 156, 15, 224, 225, 220, 233,
	0, 15, 0, 0, 15, 13, 115, 0, 0, 126,
	0, 20, 0, 97, 223, 15, 252, 15, 15, 270,
	13, 271, 13, 0, 127, 102, 0, 234, 0, 0,
	0, 272, 11, 13, 155, 257, 0, 0, 252, 259,
	0, 261, 200, 12, 202, 13, 258, 0, 252, 252,
	265, 260, 203, 252, 263, 264, 262,
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
		//line parser.y:930
		{
			RubyVAL.genericValue = ast.Hash{}
		}
	case 196:
		//line parser.y:932
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 197:
		//line parser.y:940
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 198:
		//line parser.y:948
		{
			RubyVAL.genericValue = ast.Block{Body: ast.Nodes{RubyS[Rubypt-2].genericValue}}
		}
	case 199:
		//line parser.y:951
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 200:
		//line parser.y:958
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 201:
		//line parser.y:966
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 202:
		//line parser.y:973
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 203:
		//line parser.y:980
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 204:
		//line parser.y:988
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 205:
		//line parser.y:990
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 206:
		//line parser.y:997
		{
			RubyVAL.genericValue = ast.Block{Args: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 207:
		//line parser.y:1001
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-2].genericSlice}
		}
	case 208:
		//line parser.y:1004
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 209:
		//line parser.y:1006
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 210:
		//line parser.y:1008
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 211:
		//line parser.y:1010
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 212:
		//line parser.y:1013
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 213:
		//line parser.y:1020
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 214:
		//line parser.y:1028
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 215:
		//line parser.y:1035
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 216:
		//line parser.y:1042
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 217:
		//line parser.y:1049
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 218:
		//line parser.y:1056
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 219:
		//line parser.y:1063
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 220:
		//line parser.y:1070
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 221:
		//line parser.y:1078
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 222:
		//line parser.y:1085
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 223:
		//line parser.y:1094
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 224:
		//line parser.y:1101
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 225:
		//line parser.y:1108
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 226:
		//line parser.y:1115
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 227:
		//line parser.y:1122
		{
		}
	case 228:
		//line parser.y:1123
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 229:
		//line parser.y:1124
		{
		}
	case 230:
		//line parser.y:1127
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 231:
		//line parser.y:1130
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 232:
		//line parser.y:1138
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 233:
		//line parser.y:1140
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 234:
		//line parser.y:1149
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
	case 235:
		//line parser.y:1164
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 236:
		//line parser.y:1166
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 237:
		//line parser.y:1169
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 238:
		//line parser.y:1171
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 239:
		//line parser.y:1174
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 240:
		//line parser.y:1181
		{
			RubyVAL.genericValue = ast.Yield{}
		}
	case 241:
		//line parser.y:1184
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 242:
		//line parser.y:1192
		{
			RubyVAL.genericValue = ast.Return{}
		}
	case 243:
		//line parser.y:1196
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 244:
		//line parser.y:1198
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 245:
		//line parser.y:1200
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 246:
		//line parser.y:1204
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 247:
		//line parser.y:1206
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 248:
		//line parser.y:1208
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 249:
		//line parser.y:1212
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 250:
		//line parser.y:1221
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 251:
		//line parser.y:1223
		{
			RubyVAL.genericValue = ast.Loop{Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue}, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 252:
		//line parser.y:1226
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 253:
		//line parser.y:1228
		{
		}
	case 254:
		//line parser.y:1230
		{
		}
	case 255:
		//line parser.y:1232
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 256:
		//line parser.y:1234
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 257:
		//line parser.y:1237
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 258:
		//line parser.y:1244
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 259:
		//line parser.y:1252
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 260:
		//line parser.y:1259
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 261:
		//line parser.y:1267
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 262:
		//line parser.y:1275
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 263:
		//line parser.y:1282
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 264:
		//line parser.y:1289
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 265:
		//line parser.y:1296
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 266:
		//line parser.y:1304
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 267:
		//line parser.y:1307
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 268:
		//line parser.y:1309
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 269:
		//line parser.y:1312
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 270:
		//line parser.y:1314
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 271:
		//line parser.y:1317
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 272:
		//line parser.y:1319
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 273:
		//line parser.y:1321
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
