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

//line parser.y:1367

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
	-2, 236,
	-1, 134,
	54, 144,
	-2, 21,
	-1, 229,
	54, 144,
	-2, 21,
	-1, 303,
	51, 13,
	64, 13,
	-2, 31,
	-1, 306,
	54, 145,
	-2, 143,
	-1, 317,
	10, 107,
	11, 107,
	-2, 236,
}

const RubyNprod = 285
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 3661

var RubyAct = []int{

	284, 30, 5, 523, 372, 393, 158, 18, 394, 363,
	211, 213, 215, 127, 357, 51, 129, 351, 122, 26,
	2, 3, 109, 121, 223, 90, 273, 180, 91, 181,
	61, 294, 25, 267, 131, 262, 294, 4, 96, 294,
	294, 244, 294, 110, 469, 117, 119, 96, 342, 449,
	250, 443, 447, 155, 156, 445, 323, 497, 361, 13,
	360, 88, 87, 138, 160, 161, 142, 105, 106, 234,
	131, 323, 323, 163, 94, 95, 105, 106, 89, 175,
	176, 96, 107, 94, 95, 174, 83, 182, 216, 93,
	411, 185, 186, 83, 173, 83, 414, 143, 93, 193,
	170, 83, 92, 114, 198, 521, 510, 358, 47, 203,
	105, 106, 207, 208, 209, 139, 142, 94, 95, 294,
	294, 174, 343, 322, 205, 138, 408, 219, 225, 217,
	406, 216, 93, 492, 214, 410, 104, 238, 233, 235,
	218, 257, 258, 237, 260, 261, 225, 265, 266, 249,
	270, 271, 272, 256, 135, 241, 243, 240, 140, 255,
	417, 90, 416, 253, 91, 290, 291, 292, 293, 276,
	391, 135, 217, 308, 135, 135, 297, 139, 294, 225,
	226, 294, 149, 218, 90, 96, 250, 91, 375, 135,
	135, 135, 149, 407, 146, 330, 307, 502, 226, 150,
	221, 135, 294, 135, 135, 108, 135, 314, 135, 135,
	135, 135, 315, 135, 105, 106, 135, 286, 135, 135,
	288, 94, 95, 335, 406, 148, 329, 325, 298, 230,
	511, 226, 135, 135, 135, 245, 93, 112, 96, 154,
	113, 370, 135, 474, 407, 480, 171, 230, 135, 457,
	152, 135, 481, 145, 263, 138, 146, 268, 452, 338,
	370, 274, 109, 147, 346, 382, 238, 105, 106, 96,
	153, 111, 237, 380, 94, 95, 90, 230, 135, 91,
	230, 151, 107, 110, 143, 397, 328, 90, 140, 93,
	91, 144, 96, 373, 369, 141, 356, 135, 105, 106,
	135, 386, 355, 142, 454, 94, 95, 139, 90, 135,
	135, 91, 503, 504, 377, 370, 90, 388, 154, 91,
	93, 105, 106, 287, 104, 225, 328, 168, 94, 95,
	247, 235, 387, 97, 98, 99, 135, 135, 395, 400,
	346, 396, 389, 93, 102, 100, 101, 104, 216, 354,
	355, 214, 280, 281, 402, 439, 135, 440, 118, 368,
	135, 135, 306, 166, 530, 423, 167, 500, 90, 96,
	370, 91, 430, 434, 434, 529, 426, 226, 441, 337,
	338, 538, 442, 535, 534, 450, 467, 466, 533, 217,
	535, 534, 453, 164, 465, 135, 165, 212, 105, 106,
	218, 169, 455, 463, 350, 94, 95, 464, 135, 455,
	312, 462, 313, 107, 135, 332, 331, 96, 472, 116,
	93, 115, 475, 462, 191, 327, 230, 341, 135, 183,
	278, 184, 135, 200, 201, 277, 210, 482, 483, 135,
	478, 188, 422, 421, 96, 74, 105, 106, 485, 135,
	295, 429, 489, 94, 95, 302, 349, 285, 103, 301,
	251, 1, 494, 461, 172, 92, 319, 60, 93, 135,
	135, 59, 498, 105, 106, 461, 420, 58, 422, 421,
	94, 95, 57, 501, 135, 97, 98, 99, 107, 56,
	289, 55, 38, 37, 135, 93, 102, 100, 101, 104,
	48, 36, 311, 455, 488, 35, 50, 515, 516, 517,
	519, 435, 230, 20, 40, 41, 434, 434, 434, 21,
	527, 374, 96, 15, 230, 536, 12, 11, 24, 23,
	537, 22, 19, 10, 32, 540, 27, 17, 434, 14,
	541, 542, 434, 434, 434, 543, 136, 39, 16, 34,
	135, 105, 106, 230, 33, 135, 28, 71, 94, 95,
	29, 70, 340, 136, 0, 0, 136, 136, 135, 0,
	0, 0, 0, 93, 0, 344, 0, 0, 0, 135,
	473, 136, 136, 136, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 136, 0, 136, 136, 0, 136, 135,
	136, 136, 136, 136, 135, 136, 135, 0, 136, 0,
	136, 136, 0, 0, 0, 0, 0, 0, 0, 135,
	0, 231, 0, 376, 136, 136, 136, 246, 0, 379,
	381, 383, 0, 0, 136, 0, 0, 0, 0, 231,
	136, 0, 0, 136, 0, 0, 264, 0, 0, 269,
	0, 0, 0, 275, 0, 0, 0, 69, 133, 75,
	134, 120, 512, 0, 131, 0, 403, 0, 0, 231,
	136, 0, 231, 0, 0, 0, 0, 0, 0, 0,
	413, 0, 0, 415, 0, 0, 0, 77, 0, 136,
	85, 86, 136, 0, 96, 78, 79, 80, 81, 82,
	0, 136, 136, 96, 0, 444, 0, 0, 446, 0,
	448, 0, 191, 252, 0, 132, 0, 84, 83, 0,
	0, 0, 0, 105, 106, 0, 0, 0, 136, 136,
	94, 95, 105, 106, 0, 0, 0, 279, 0, 94,
	95, 0, 0, 0, 471, 93, 0, 0, 136, 0,
	0, 0, 136, 136, 93, 69, 45, 75, 46, 76,
	0, 359, 0, 42, 526, 436, 525, 524, 437, 43,
	44, 0, 63, 64, 54, 0, 0, 67, 68, 0,
	0, 66, 62, 0, 0, 77, 65, 136, 85, 86,
	0, 0, 157, 78, 79, 80, 81, 82, 0, 509,
	136, 432, 433, 96, 0, 0, 136, 0, 0, 0,
	532, 72, 31, 73, 507, 84, 83, 0, 231, 0,
	136, 0, 0, 0, 136, 0, 0, 0, 105, 106,
	0, 136, 105, 106, 518, 94, 95, 191, 0, 94,
	95, 136, 0, 0, 0, 0, 0, 0, 531, 0,
	93, 96, 0, 0, 93, 220, 0, 222, 130, 0,
	539, 136, 136, 0, 0, 0, 0, 0, 0, 248,
	0, 0, 0, 0, 0, 130, 136, 0, 130, 130,
	105, 106, 0, 0, 0, 0, 136, 94, 95, 384,
	0, 0, 0, 130, 130, 130, 0, 0, 0, 0,
	0, 0, 93, 0, 231, 130, 0, 130, 130, 324,
	130, 0, 130, 130, 130, 130, 231, 130, 105, 106,
	130, 0, 130, 130, 0, 94, 95, 0, 0, 0,
	0, 0, 0, 130, 0, 0, 130, 130, 130, 0,
	93, 0, 136, 0, 0, 231, 130, 136, 0, 326,
	0, 130, 130, 0, 0, 130, 0, 333, 0, 0,
	136, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 136, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 130, 303, 0, 130, 0, 0, 348, 0, 352,
	0, 136, 0, 0, 0, 0, 136, 0, 136, 0,
	0, 130, 0, 0, 130, 0, 0, 0, 0, 0,
	0, 136, 0, 130, 130, 0, 0, 0, 366, 367,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	130, 130, 0, 0, 69, 45, 75, 46, 76, 352,
	0, 0, 42, 522, 436, 525, 524, 437, 43, 44,
	130, 63, 64, 54, 303, 130, 67, 68, 0, 0,
	66, 62, 0, 401, 77, 65, 0, 85, 86, 0,
	0, 0, 78, 79, 80, 81, 82, 0, 0, 412,
	432, 433, 0, 0, 0, 0, 0, 0, 0, 130,
	72, 0, 73, 424, 84, 83, 0, 0, 428, 0,
	0, 0, 130, 0, 0, 0, 0, 0, 130, 0,
	0, 0, 0, 0, 9, 0, 0, 0, 0, 0,
	130, 0, 130, 456, 0, 0, 130, 0, 0, 0,
	0, 0, 0, 130, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 130, 0, 0, 0, 0, 0, 0,
	476, 477, 0, 0, 0, 0, 0, 479, 0, 0,
	128, 0, 0, 130, 130, 0, 0, 0, 484, 0,
	486, 0, 0, 0, 0, 0, 0, 159, 130, 0,
	162, 159, 0, 0, 0, 0, 0, 0, 130, 0,
	0, 0, 0, 0, 0, 177, 178, 179, 0, 0,
	0, 0, 0, 499, 0, 0, 130, 187, 0, 189,
	190, 0, 192, 506, 194, 195, 196, 197, 130, 199,
	0, 0, 202, 0, 204, 206, 0, 0, 514, 0,
	0, 366, 367, 0, 0, 224, 0, 0, 236, 239,
	242, 0, 0, 0, 130, 0, 0, 130, 128, 130,
	0, 0, 0, 224, 254, 0, 0, 259, 0, 0,
	0, 0, 130, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 130, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 296, 304, 0, 224, 0, 0, 0,
	0, 0, 0, 130, 0, 0, 0, 0, 130, 0,
	130, 0, 0, 128, 0, 0, 318, 0, 0, 0,
	0, 0, 0, 130, 0, 320, 321, 0, 69, 45,
	75, 46, 76, 0, 0, 0, 42, 425, 52, 365,
	364, 53, 43, 44, 0, 63, 64, 54, 0, 0,
	67, 68, 339, 128, 66, 62, 0, 0, 77, 65,
	0, 85, 86, 0, 0, 0, 78, 79, 80, 81,
	82, 0, 345, 0, 282, 283, 304, 353, 0, 0,
	0, 0, 0, 0, 72, 0, 73, 0, 84, 83,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 371, 0, 0, 0, 0, 0, 0, 0, 69,
	133, 75, 134, 120, 378, 126, 131, 0, 0, 0,
	128, 0, 0, 0, 0, 0, 49, 0, 0, 0,
	0, 0, 224, 0, 390, 0, 0, 0, 345, 77,
	0, 0, 85, 86, 0, 398, 124, 78, 79, 80,
	81, 82, 0, 125, 0, 404, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 123, 0, 132, 0, 84,
	83, 0, 137, 0, 0, 418, 419, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 137,
	159, 0, 137, 137, 0, 0, 0, 0, 0, 0,
	451, 0, 0, 0, 0, 0, 0, 137, 137, 137,
	0, 0, 0, 0, 0, 0, 0, 0, 460, 137,
	0, 137, 137, 0, 137, 0, 137, 137, 137, 137,
	460, 137, 0, 0, 137, 0, 137, 137, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 232, 0, 0,
	137, 137, 137, 0, 0, 0, 159, 0, 0, 487,
	137, 491, 0, 0, 0, 232, 137, 0, 0, 137,
	0, 0, 0, 0, 496, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 496, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 232, 137, 0, 232, 0,
	0, 0, 0, 0, 0, 508, 0, 0, 0, 0,
	128, 0, 513, 0, 0, 137, 0, 0, 137, 0,
	0, 0, 0, 0, 0, 520, 0, 137, 137, 0,
	69, 45, 75, 46, 76, 0, 0, 0, 42, 362,
	52, 365, 364, 53, 43, 44, 0, 63, 64, 54,
	0, 0, 67, 68, 137, 137, 66, 62, 0, 0,
	77, 65, 0, 85, 86, 0, 0, 0, 78, 79,
	80, 81, 82, 0, 137, 0, 282, 283, 137, 137,
	0, 0, 0, 0, 0, 0, 72, 0, 73, 0,
	84, 83, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 137, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 137, 0, 0, 0,
	0, 0, 137, 0, 0, 69, 305, 75, 134, 76,
	0, 0, 0, 0, 232, 0, 137, 0, 0, 0,
	137, 0, 0, 0, 0, 0, 0, 137, 0, 0,
	0, 0, 0, 0, 0, 77, 0, 137, 85, 86,
	0, 0, 0, 78, 79, 80, 81, 82, 0, 0,
	0, 294, 0, 0, 0, 0, 250, 137, 137, 0,
	0, 72, 0, 73, 300, 84, 83, 0, 0, 0,
	0, 0, 137, 0, 0, 0, 69, 45, 75, 46,
	76, 0, 137, 0, 42, 0, 52, 0, 0, 53,
	43, 44, 0, 63, 64, 54, 0, 0, 67, 68,
	232, 0, 66, 62, 0, 0, 77, 65, 0, 85,
	86, 0, 232, 0, 78, 79, 80, 81, 82, 0,
	0, 0, 6, 7, 69, 228, 75, 229, 76, 0,
	0, 0, 72, 0, 73, 0, 84, 83, 137, 8,
	0, 232, 0, 137, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 77, 0, 137, 85, 86, 0,
	0, 0, 78, 79, 80, 81, 82, 137, 0, 0,
	294, 0, 0, 0, 0, 0, 0, 490, 0, 0,
	72, 0, 73, 0, 84, 83, 0, 137, 0, 0,
	0, 0, 137, 0, 137, 69, 45, 75, 46, 76,
	0, 0, 0, 42, 493, 52, 0, 137, 53, 43,
	44, 0, 63, 64, 54, 370, 0, 67, 68, 0,
	0, 66, 62, 0, 0, 77, 65, 0, 85, 86,
	0, 0, 0, 78, 79, 80, 81, 82, 0, 0,
	0, 282, 283, 0, 0, 0, 0, 0, 0, 0,
	0, 72, 0, 73, 0, 84, 83, 69, 45, 75,
	46, 76, 0, 0, 0, 42, 399, 52, 0, 0,
	53, 43, 44, 0, 63, 64, 54, 370, 0, 67,
	68, 0, 0, 66, 62, 0, 0, 77, 65, 0,
	85, 86, 0, 0, 0, 78, 79, 80, 81, 82,
	0, 0, 0, 282, 283, 0, 0, 0, 0, 0,
	0, 0, 0, 72, 0, 73, 0, 84, 83, 69,
	45, 75, 46, 76, 0, 0, 0, 42, 392, 52,
	0, 0, 53, 43, 44, 0, 63, 64, 54, 370,
	0, 67, 68, 0, 0, 66, 62, 0, 0, 77,
	65, 0, 85, 86, 0, 0, 0, 78, 79, 80,
	81, 82, 0, 0, 0, 282, 283, 0, 0, 0,
	0, 0, 0, 0, 0, 72, 0, 73, 0, 84,
	83, 69, 45, 75, 46, 76, 0, 0, 0, 42,
	528, 436, 0, 0, 437, 43, 44, 0, 63, 64,
	54, 0, 0, 67, 68, 0, 0, 66, 62, 0,
	0, 77, 65, 0, 85, 86, 0, 0, 0, 78,
	79, 80, 81, 82, 0, 0, 0, 432, 433, 0,
	0, 0, 0, 0, 0, 0, 0, 72, 0, 73,
	0, 84, 83, 69, 45, 75, 46, 76, 0, 0,
	0, 42, 505, 52, 0, 0, 53, 43, 44, 0,
	63, 64, 54, 0, 0, 67, 68, 0, 0, 66,
	62, 0, 0, 77, 65, 0, 85, 86, 0, 0,
	0, 78, 79, 80, 81, 82, 0, 0, 0, 282,
	283, 0, 69, 45, 75, 46, 76, 0, 0, 72,
	42, 73, 52, 84, 83, 53, 43, 44, 0, 63,
	64, 54, 0, 0, 67, 68, 0, 0, 66, 62,
	0, 0, 77, 65, 0, 85, 86, 0, 0, 0,
	78, 79, 80, 81, 82, 0, 0, 0, 282, 283,
	0, 0, 0, 0, 0, 0, 0, 0, 72, 0,
	73, 468, 84, 83, 69, 45, 75, 46, 76, 0,
	0, 0, 42, 458, 52, 0, 0, 53, 43, 44,
	0, 63, 64, 54, 0, 0, 67, 68, 0, 0,
	66, 62, 0, 0, 77, 65, 0, 85, 86, 0,
	0, 0, 78, 79, 80, 81, 82, 0, 0, 0,
	282, 283, 0, 0, 0, 0, 0, 0, 0, 0,
	72, 0, 73, 0, 84, 83, 69, 45, 75, 46,
	76, 0, 0, 0, 42, 438, 436, 0, 0, 437,
	43, 44, 0, 63, 64, 54, 0, 0, 67, 68,
	0, 0, 66, 62, 0, 0, 77, 65, 0, 85,
	86, 0, 0, 0, 78, 79, 80, 81, 82, 0,
	0, 0, 432, 433, 0, 0, 0, 0, 0, 0,
	0, 0, 72, 0, 73, 0, 84, 83, 69, 45,
	75, 46, 76, 0, 0, 0, 42, 431, 436, 0,
	0, 437, 43, 44, 0, 63, 64, 54, 0, 0,
	67, 68, 0, 0, 66, 62, 0, 0, 77, 65,
	0, 85, 86, 0, 0, 0, 78, 79, 80, 81,
	82, 0, 0, 0, 432, 433, 0, 0, 0, 0,
	0, 0, 0, 0, 72, 0, 73, 0, 84, 83,
	69, 45, 75, 46, 76, 0, 0, 0, 42, 427,
	52, 0, 0, 53, 43, 44, 0, 63, 64, 54,
	0, 0, 67, 68, 0, 0, 66, 62, 0, 0,
	77, 65, 0, 85, 86, 0, 0, 0, 78, 79,
	80, 81, 82, 0, 0, 0, 282, 283, 0, 0,
	0, 0, 0, 0, 0, 0, 72, 0, 73, 0,
	84, 83, 69, 45, 75, 46, 76, 0, 0, 0,
	42, 409, 52, 0, 0, 53, 43, 44, 0, 63,
	64, 54, 0, 0, 67, 68, 0, 0, 66, 62,
	0, 0, 77, 65, 0, 85, 86, 0, 0, 0,
	78, 79, 80, 81, 82, 0, 0, 0, 282, 283,
	0, 0, 0, 0, 0, 0, 0, 0, 72, 0,
	73, 0, 84, 83, 69, 45, 75, 46, 76, 0,
	0, 0, 42, 347, 52, 0, 0, 53, 43, 44,
	0, 63, 64, 54, 0, 0, 67, 68, 0, 0,
	66, 62, 0, 0, 77, 65, 0, 85, 86, 0,
	0, 0, 78, 79, 80, 81, 82, 0, 0, 0,
	282, 283, 0, 0, 0, 0, 0, 0, 0, 0,
	72, 0, 73, 0, 84, 83, 69, 45, 75, 46,
	76, 0, 0, 0, 42, 336, 52, 0, 0, 53,
	43, 44, 0, 63, 64, 54, 0, 0, 67, 68,
	0, 0, 66, 62, 0, 0, 77, 65, 0, 85,
	86, 0, 0, 0, 78, 79, 80, 81, 82, 0,
	0, 0, 282, 283, 0, 0, 0, 0, 0, 0,
	0, 0, 72, 0, 73, 0, 84, 83, 69, 45,
	75, 46, 76, 0, 0, 0, 42, 334, 52, 0,
	0, 53, 43, 44, 0, 63, 64, 54, 0, 0,
	67, 68, 0, 0, 66, 62, 0, 0, 77, 65,
	0, 85, 86, 0, 0, 0, 78, 79, 80, 81,
	82, 0, 0, 0, 282, 283, 0, 69, 45, 75,
	46, 76, 0, 0, 72, 42, 73, 436, 84, 83,
	437, 43, 44, 0, 63, 64, 54, 0, 0, 67,
	68, 0, 0, 66, 62, 0, 0, 77, 65, 0,
	85, 86, 0, 0, 0, 78, 79, 80, 81, 82,
	0, 0, 0, 432, 433, 0, 69, 45, 75, 46,
	76, 0, 0, 72, 42, 73, 52, 84, 83, 53,
	43, 44, 0, 63, 64, 54, 0, 0, 67, 68,
	0, 0, 66, 62, 0, 0, 77, 65, 0, 85,
	86, 0, 0, 0, 78, 79, 80, 81, 82, 0,
	0, 0, 282, 283, 0, 69, 45, 75, 46, 76,
	310, 0, 72, 42, 73, 52, 84, 83, 53, 43,
	44, 0, 63, 64, 54, 0, 0, 67, 68, 0,
	0, 66, 62, 0, 0, 77, 65, 0, 85, 86,
	0, 0, 0, 78, 79, 80, 81, 82, 0, 0,
	0, 0, 309, 0, 69, 45, 75, 46, 76, 0,
	0, 72, 42, 73, 52, 84, 83, 53, 43, 44,
	0, 63, 64, 54, 0, 0, 67, 68, 0, 0,
	66, 62, 0, 0, 77, 65, 0, 85, 86, 0,
	0, 0, 78, 79, 80, 81, 82, 0, 0, 0,
	294, 0, 0, 69, 45, 75, 46, 76, 0, 0,
	72, 42, 73, 52, 84, 83, 53, 43, 44, 0,
	63, 64, 54, 0, 0, 67, 68, 0, 0, 66,
	62, 0, 0, 77, 65, 0, 85, 86, 0, 0,
	0, 78, 79, 80, 81, 82, 69, 228, 75, 229,
	76, 0, 0, 0, 0, 0, 0, 0, 0, 72,
	0, 73, 0, 84, 83, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 77, 0, 0, 85,
	86, 0, 0, 0, 78, 79, 80, 81, 82, 0,
	0, 0, 294, 69, 228, 75, 229, 76, 0, 470,
	0, 0, 72, 0, 73, 0, 84, 83, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 77, 0, 0, 85, 86, 0, 0,
	0, 78, 79, 80, 81, 82, 0, 0, 0, 294,
	69, 133, 75, 134, 76, 0, 459, 0, 0, 72,
	0, 73, 0, 84, 83, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	77, 0, 0, 85, 86, 0, 0, 0, 78, 79,
	80, 81, 82, 0, 0, 0, 294, 69, 133, 75,
	134, 120, 0, 405, 131, 0, 72, 0, 73, 0,
	84, 83, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 77, 0, 0,
	85, 86, 0, 0, 0, 78, 79, 80, 81, 82,
	69, 133, 75, 134, 120, 385, 0, 131, 0, 0,
	0, 0, 0, 252, 0, 132, 0, 84, 83, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	77, 0, 0, 85, 86, 0, 0, 124, 78, 79,
	80, 81, 82, 69, 133, 75, 134, 317, 0, 0,
	131, 0, 0, 0, 0, 0, 252, 0, 132, 0,
	84, 83, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 77, 0, 0, 85, 86, 0, 0,
	316, 78, 79, 80, 81, 82, 69, 305, 75, 134,
	76, 0, 0, 0, 0, 0, 0, 0, 0, 72,
	0, 132, 0, 84, 83, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 77, 0, 0, 85,
	86, 0, 0, 0, 78, 79, 80, 81, 82, 0,
	0, 0, 294, 69, 228, 75, 229, 76, 0, 0,
	0, 0, 72, 0, 73, 300, 84, 83, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 77, 0, 0, 85, 86, 0, 0,
	0, 78, 79, 80, 81, 82, 0, 0, 0, 294,
	69, 133, 75, 134, 76, 0, 299, 0, 0, 72,
	0, 73, 0, 84, 83, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	77, 0, 0, 85, 86, 0, 0, 0, 78, 79,
	80, 81, 82, 0, 0, 0, 294, 69, 133, 75,
	134, 76, 0, 0, 131, 0, 72, 0, 73, 0,
	84, 83, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 77, 0, 0,
	85, 86, 0, 0, 0, 78, 79, 80, 81, 82,
	69, 228, 75, 229, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 72, 0, 132, 0, 84, 83, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	77, 0, 0, 85, 86, 0, 0, 0, 78, 79,
	80, 81, 82, 69, 133, 75, 134, 120, 0, 0,
	131, 0, 0, 227, 0, 0, 72, 0, 73, 0,
	84, 83, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 77, 0, 0, 85, 86, 0, 0,
	0, 78, 79, 80, 81, 82, 69, 133, 75, 134,
	76, 0, 0, 0, 0, 0, 0, 0, 0, 252,
	0, 132, 0, 84, 83, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 77, 0, 0, 85,
	86, 0, 0, 0, 78, 79, 80, 81, 82, 69,
	495, 75, 134, 76, 0, 0, 0, 0, 0, 0,
	0, 0, 72, 0, 73, 0, 84, 83, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 77,
	0, 0, 85, 86, 0, 0, 0, 78, 79, 80,
	81, 82, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 72, 0, 73, 0, 84,
	83,
}
var RubyPact = []int{

	-31, 1811, -1000, -1000, -1000, 10, -1000, -1000, -1000, 440,
	-1000, -1000, -1000, 187, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 222, -1000, 48, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 415, 350, 350, 1414, 246, 242, 214, 150,
	239, 228, 2968, 2968, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 3551, 2968, 2968, 3551, 3551, 378, 348, -1000,
	-1000, -1000, -1000, -1000, 46, 237, -1000, 22, 2968, 2968,
	3551, 3551, 3551, 21, 423, -1000, -1000, -1000, -1000, -1000,
	2968, 2968, 3551, 435, 3551, 3551, -1000, 3551, 2968, 3551,
	3551, 3551, 3551, 2968, 3551, -1000, -1000, 3551, 2968, 3551,
	3551, 2968, 2968, 2968, 430, 342, 125, 160, -1000, -1000,
	3465, 22, 58, 3551, 3551, 3551, 35, 319, 690, -1000,
	-18, -6, -1000, 3508, 54, 5, -1000, -1000, 3465, 3551,
	2968, 2968, 3551, 2968, 2968, 29, 2968, 2968, 27, 2968,
	2968, 2968, 20, 429, 424, 293, 301, 2821, 206, 690,
	272, 169, 690, 206, 2968, 2968, 2968, 2968, 3328, 3281,
	354, 3465, 2870, -1000, -1000, 293, 293, 690, 690, 690,
	-1000, -1000, 404, -1000, -1000, 293, 293, 690, 3238, 690,
	690, 3375, 690, 293, 690, 690, 690, 690, 293, 413,
	3375, 3375, 690, 293, 690, 61, 847, 293, 293, 293,
	22, -1000, 419, 315, 82, -1000, 153, 410, 409, -1000,
	2723, 350, 2661, 369, 288, -1000, -1000, 3551, 3195, 116,
	55, 152, 140, -1000, -1000, -1000, 365, -14, 60, 77,
	-1000, -1000, 265, -1000, -1000, -1000, -1000, 3422, 2599, -1000,
	398, 1740, 3551, 339, 234, -48, 45, 293, 293, 699,
	293, 293, -1000, -1000, -1000, 293, 293, -1000, -1000, -1000,
	293, 293, 293, -1000, -1000, -1000, 307, -1, -3, 1635,
	-1000, -1000, -1000, -1000, 293, 345, 3551, -1000, -1000, 151,
	293, 293, 293, 293, -1000, 303, 288, -1000, -1000, 3551,
	-1000, 262, 254, -18, 885, 3152, -1000, 291, 293, -1000,
	-1000, 130, -1000, -1000, 58, 22, 2968, 3465, 690, 3551,
	690, 690, -1000, 3422, 128, -1000, 2054, 125, 82, 275,
	3551, -1000, -1000, 1992, -1000, -1000, -1000, 22, -1000, 690,
	3105, 182, -1000, -1000, 68, 690, -1000, -1000, 2537, 79,
	-1000, -1000, 2821, 34, -1000, -1000, 120, -1000, -1000, 118,
	3551, 3551, -1000, 462, 2968, -1000, 1323, 2475, -1000, -1000,
	350, 690, 2413, 2351, 341, 3551, -11, -1000, 690, -9,
	-1000, -12, -1000, -15, 2968, 3551, -1000, -1000, 293, 248,
	690, 2968, -1000, 290, -1000, -1000, -1000, -1000, 690, -1000,
	235, 2289, -1000, 3058, 690, 397, 2968, 388, 381, -1000,
	-1000, 380, 2227, -20, 88, 3011, -1000, 2968, 518, 181,
	-1000, 2968, -1000, 293, 2821, -1000, 426, -1000, 2821, 241,
	-1000, -1000, -1000, -1000, 293, -1000, 2968, 2968, -1000, -1000,
	-1000, 3551, 206, -1000, 1859, -1000, 3375, -1000, 127, -1000,
	293, 690, -1000, 293, -1000, -1000, 1930, -1000, -1000, 3594,
	288, -1000, -1000, -1000, 293, -4, -1000, -1000, -1000, -1000,
	3551, 2919, 293, 233, -1000, 293, 2821, 2821, -1000, 2821,
	361, 350, 146, 261, 2178, 206, 2821, 288, -1000, -1000,
	3551, 795, 53, -1000, 216, 652, 690, 3551, 293, 2821,
	-1000, -1000, -1000, -1000, -1000, -1000, 2821, 69, 690, 2968,
	3551, -1000, -1000, 43, 2821, 1039, 750, 2116, 69, 353,
	799, -1000, -1000, 374, 2968, -1000, -1000, 367, -1000, -1000,
	-1000, 69, -1000, -1000, 2968, -1000, 293, 2772, -1000, 69,
	293, 2772, 2772, 2772,
}
var RubyPgo = []int{

	0, 0, 561, 560, 19, 16, 557, 556, 554, 1436,
	549, 8, 30, 548, 547, 59, 539, 537, 1124, 536,
	500, 812, 534, 533, 532, 531, 529, 528, 7, 527,
	12, 108, 526, 523, 1, 521, 519, 515, 514, 32,
	513, 511, 3, 506, 505, 501, 493, 492, 491, 489,
	482, 477, 471, 467, 737, 464, 5, 23, 17, 9,
	461, 10, 459, 4, 457, 24, 456, 6, 455, 13,
	15, 18, 11, 451, 450, 445, 375, 327,
}
var RubyR1 = []int{

	0, 60, 60, 60, 60, 60, 60, 60, 60, 60,
	60, 76, 76, 77, 77, 54, 54, 54, 54, 19,
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
	29, 32, 34, 34, 75, 75, 15, 15, 15, 15,
	15, 15, 15, 16, 16, 70, 70, 33, 33, 33,
	33, 33, 33, 33, 9, 9, 31, 31, 20, 20,
	43, 43, 43, 43, 43, 43, 43, 43, 43, 43,
	43, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 2, 6, 7, 74, 74, 74, 74, 74,
	74, 74, 74, 74, 3, 3, 3, 3, 62, 62,
	68, 68, 68, 5, 5, 5, 5, 58, 66, 66,
	66, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 59, 59, 59, 59, 55, 55, 55, 8,
	17, 11, 11, 11, 73, 73, 64, 64, 56, 56,
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
	3, 3, 1, 1, 5, 0, 1, 1, 1, 2,
	4, 4, 4, 5, 3, 5, 5, 5, 3, 7,
	3, 7, 8, 3, 4, 5, 5, 3, 0, 1,
	3, 4, 5, 3, 3, 3, 3, 3, 5, 6,
	5, 3, 4, 3, 3, 2, 0, 2, 2, 3,
	4, 2, 3, 5, 1, 3, 0, 2, 1, 2,
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
	-2, -6, 61, 63, -75, 7, 9, 35, 43, 44,
	45, 46, 47, 66, 65, 38, 39, 52, 51, 68,
	15, 18, 25, 55, 40, 41, 4, 45, 46, 47,
	57, 58, 56, 18, 59, 33, 34, 48, 18, 40,
	61, 49, 15, 18, 55, 6, 4, -34, 8, -34,
	9, -57, -71, 61, 42, 49, 11, -69, -18, -5,
	-21, 12, 63, 6, 8, -31, -20, -9, 9, 61,
	42, 49, 61, 42, 49, 11, 42, 49, 11, 42,
	49, 42, 11, 42, 11, -1, -1, -54, -67, -18,
	-1, -1, -18, -67, 15, 18, 15, 18, -77, -77,
	54, 9, -55, -5, 63, -1, -1, -18, -18, -18,
	6, 8, 66, 6, 8, -1, -1, -18, 6, -18,
	-18, -77, -18, -1, -18, -18, -18, -18, -1, -18,
	-77, -77, -18, -1, -18, -69, -18, -1, -1, -1,
	6, -61, 55, -72, 9, -30, 6, 47, 58, -61,
	-54, 40, -54, -65, -18, -28, -15, 58, 6, 8,
	-31, -20, -9, -5, 11, -5, -18, -4, -69, -18,
	-39, -12, -18, -12, 6, -31, -20, 11, -54, -58,
	56, -77, 61, -65, -18, -4, -69, -1, -1, -18,
	-1, -1, 6, -31, -20, -1, -1, 6, -31, -20,
	-1, -1, -1, 6, -31, -20, -70, 6, 6, -54,
	51, 52, 51, 52, -1, -64, 11, 51, 51, -77,
	-1, -1, -1, -1, 51, -74, -18, -28, -15, 58,
	64, -62, -68, -21, -18, 6, 8, -65, -1, 52,
	10, -77, 6, 8, -71, -57, 42, 9, -18, 53,
	-18, -18, 62, 11, 62, -5, -54, 6, 11, -72,
	42, 6, 6, -54, 14, -34, 14, 10, 11, -18,
	-77, 62, 62, 62, -77, -18, -5, 14, -54, -66,
	6, -58, -54, -18, 10, 11, 62, 62, 62, 62,
	61, 61, 14, -59, 17, 16, -54, -54, 14, -11,
	25, -18, -63, -63, -35, 37, -77, 11, -18, -77,
	11, -77, 11, -77, 4, 53, 10, -5, -1, -65,
	-18, 42, 14, -56, -11, -61, -30, 10, -18, 14,
	-56, -54, -5, -77, -18, 58, 42, 11, 58, 14,
	56, 11, -54, -77, 62, -77, 42, 42, -18, -18,
	14, 17, 16, -1, -54, 14, -59, 14, -54, -73,
	-34, 14, 51, 52, -1, -41, 15, 18, 14, 14,
	16, 37, -67, 62, -77, 64, -77, 64, -77, 64,
	-1, -18, 10, -1, 14, -11, -54, 14, 14, 58,
	-18, -15, -28, 6, -1, 6, 6, 6, 64, 64,
	58, -77, -1, 62, 62, -1, -54, -54, 14, -54,
	4, 11, -1, -1, -54, -67, -54, -18, -15, -28,
	58, -18, 6, 14, -56, 6, -18, 61, -1, -54,
	6, -34, 51, 51, 52, 14, -54, -77, -18, 4,
	53, 14, 10, -18, -54, -63, -63, -63, -77, -1,
	-18, 62, 14, -42, 17, 16, 14, -42, 14, -76,
	11, -77, 11, 14, 17, 16, -1, -63, 14, -77,
	-1, -63, -63, -63,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 49,
	50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
	60, 61, 62, 63, 64, 65, 66, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 0, 0, 0, 20, -2, 22, 23, 24,
	0, 0, 0, 0, 15, 42, 43, 44, 45, 46,
	47, 48, 251, 0, 0, 0, 253, 257, 254, 19,
	25, 26, 13, 13, 0, 71, 236, 0, 0, 0,
	0, 0, 0, 0, 0, 192, 193, 5, 6, 7,
	0, 0, 0, 0, 0, 0, 13, 0, 0, 0,
	0, 0, 0, 0, 0, 13, 13, 0, 0, 0,
	0, 0, 0, 0, 0, 132, 132, 15, -2, 15,
	-2, 74, 83, 13, 0, 0, 0, 103, 116, 117,
	31, 15, 13, 20, -2, 22, 23, 24, 107, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 15, 0, 246, 250, 105,
	0, 0, 13, 252, 0, 0, 0, 0, 195, 0,
	0, 107, 0, 279, 13, 182, 183, 184, 185, 68,
	166, 167, 0, 164, 165, 223, 231, 67, 76, 84,
	86, 0, 186, 187, 188, 189, 190, 191, 225, 0,
	0, 0, 284, 227, 85, 0, 121, 163, 224, 226,
	80, 15, 0, 130, 132, 133, 135, 0, 0, 15,
	0, 0, 0, 0, 108, 109, 110, 0, 20, -2,
	22, 23, 24, 75, 13, 119, 121, 0, 0, 146,
	147, 148, 157, 158, 170, 171, 172, 13, 0, 15,
	218, 15, 13, 0, 121, 0, 0, 149, 159, 0,
	150, 160, 173, 174, 175, 151, 161, 176, 177, 178,
	152, 162, 153, 179, 180, 181, 154, 0, 0, 0,
	15, 15, 16, 17, 18, 0, 0, 263, 263, 0,
	258, 259, 255, 256, 14, 13, 196, 197, 198, 0,
	204, 13, 13, -2, 0, 20, -2, 0, 237, 238,
	239, 15, 168, 169, 77, 78, 0, -2, 100, 0,
	277, 278, 94, 0, 95, 81, 0, 132, 0, 0,
	0, 136, 138, 0, 139, 15, 141, 69, 13, 111,
	0, 87, 90, 92, 0, 122, 123, 213, 0, 0,
	219, 15, 13, 121, 73, 13, 88, 91, 93, 89,
	0, 0, 221, 0, 0, 15, 0, 0, 240, 247,
	15, 106, 0, 0, 0, 0, 0, 13, 199, 0,
	13, 0, 13, 0, 13, 0, 72, 79, 82, 0,
	260, 0, 124, 0, 248, 15, 134, 131, 137, 128,
	0, 0, 70, 0, 118, 0, 0, 0, 0, 214,
	217, 0, 0, 0, 87, 0, 13, 0, 0, 0,
	222, 0, 15, 15, 235, 228, 0, 230, 241, 15,
	244, 261, 264, 265, 266, 267, 0, 0, 262, 280,
	15, 0, 15, 194, 0, 205, 0, 206, 0, 207,
	208, 210, 101, 99, 125, 249, 0, 129, 140, 0,
	112, 113, 114, 120, 96, 0, 104, 220, 215, 216,
	0, 0, 98, 0, 156, 15, 233, 234, 229, 242,
	0, 0, 15, 0, 0, 15, 13, 200, 201, 202,
	0, 0, 0, 126, 0, 20, 115, 0, 97, 232,
	15, 245, 263, 15, 15, 281, 13, 282, 203, 13,
	0, 127, 102, 0, 243, 0, 0, 0, 283, 11,
	13, 155, 268, 0, 0, 263, 270, 0, 272, 209,
	12, 211, 13, 269, 0, 263, 263, 276, 271, 212,
	263, 274, 275, 273,
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
		//line parser.y:207
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:209
		{
		}
	case 3:
		//line parser.y:211
		{
		}
	case 4:
		//line parser.y:213
		{
		}
	case 5:
		//line parser.y:215
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:217
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:219
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:225
		{
		}
	case 11:
		//line parser.y:227
		{
		}
	case 12:
		//line parser.y:228
		{
		}
	case 13:
		//line parser.y:230
		{
		}
	case 14:
		//line parser.y:231
		{
		}
	case 15:
		//line parser.y:234
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 16:
		//line parser.y:236
		{
		}
	case 17:
		//line parser.y:238
		{
		}
	case 18:
		//line parser.y:240
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
		//line parser.y:252
		{
			RubyVAL.genericValue = ast.RescueModifier{Statement: RubyS[Rubypt-2].genericValue, Rescue: RubyS[Rubypt-0].genericValue}
		}
	case 68:
		//line parser.y:255
		{
			RubyVAL.genericValue = ast.StarSplat{Value: RubyS[Rubypt-0].genericValue}
		}
	case 69:
		//line parser.y:258
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 70:
		//line parser.y:265
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-2].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 71:
		//line parser.y:272
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 72:
		//line parser.y:276
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 73:
		//line parser.y:283
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 74:
		//line parser.y:290
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 75:
		//line parser.y:297
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 76:
		//line parser.y:304
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 77:
		//line parser.y:311
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 78:
		//line parser.y:319
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 79:
		//line parser.y:327
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 80:
		//line parser.y:335
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 81:
		//line parser.y:343
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 82:
		//line parser.y:351
		{
			methodName := RubyS[Rubypt-2].genericValue.(ast.BareReference).Name + "="
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: methodName},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 83:
		//line parser.y:362
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 84:
		//line parser.y:369
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 85:
		//line parser.y:377
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 86:
		//line parser.y:385
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 87:
		//line parser.y:395
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 88:
		//line parser.y:403
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 89:
		//line parser.y:411
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 90:
		//line parser.y:419
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 91:
		//line parser.y:427
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 92:
		//line parser.y:435
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 93:
		//line parser.y:443
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 94:
		//line parser.y:451
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 95:
		//line parser.y:459
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 96:
		//line parser.y:469
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 97:
		//line parser.y:477
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
		//line parser.y:488
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 99:
		//line parser.y:496
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 100:
		//line parser.y:506
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 101:
		//line parser.y:516
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 102:
		//line parser.y:518
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 103:
		//line parser.y:520
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 104:
		//line parser.y:522
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 105:
		//line parser.y:525
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 106:
		//line parser.y:527
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 107:
		//line parser.y:529
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 108:
		//line parser.y:531
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 109:
		//line parser.y:533
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 110:
		//line parser.y:535
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 111:
		//line parser.y:537
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{
				Func:   ast.BareReference{Name: "to_proc"},
				Target: RubyS[Rubypt-0].genericValue,
			})
		}
	case 112:
		//line parser.y:544
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 113:
		//line parser.y:546
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 114:
		//line parser.y:548
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 115:
		//line parser.y:550
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{
				Func:   ast.BareReference{Name: "to_proc"},
				Target: RubyS[Rubypt-0].genericValue,
			})
		}
	case 116:
		//line parser.y:559
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 117:
		//line parser.y:561
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 118:
		//line parser.y:563
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 119:
		//line parser.y:565
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 120:
		//line parser.y:567
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 121:
		//line parser.y:570
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 122:
		//line parser.y:572
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 123:
		//line parser.y:574
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 124:
		//line parser.y:578
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 125:
		//line parser.y:586
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 126:
		//line parser.y:595
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 127:
		//line parser.y:604
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
		//line parser.y:614
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 129:
		//line parser.y:622
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 130:
		//line parser.y:633
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 131:
		//line parser.y:635
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 132:
		//line parser.y:637
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 133:
		//line parser.y:639
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 134:
		//line parser.y:641
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 135:
		//line parser.y:644
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 136:
		//line parser.y:646
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 137:
		//line parser.y:648
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 138:
		//line parser.y:650
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 139:
		//line parser.y:654
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 140:
		//line parser.y:661
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 141:
		//line parser.y:671
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 142:
		//line parser.y:680
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 143:
		//line parser.y:686
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 144:
		//line parser.y:694
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 145:
		//line parser.y:698
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 146:
		//line parser.y:703
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 147:
		//line parser.y:710
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 148:
		//line parser.y:714
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 149:
		//line parser.y:718
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 150:
		//line parser.y:725
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 151:
		//line parser.y:732
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 152:
		//line parser.y:739
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 153:
		//line parser.y:747
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 154:
		//line parser.y:754
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: ast.Array{Nodes: RubyS[Rubypt-2].genericSlice},
				RHS: ast.Array{Nodes: RubyS[Rubypt-0].genericSlice},
			}
		}
	case 155:
		//line parser.y:762
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
		//line parser.y:777
		{
			tail := ast.CallExpression{Target: RubyS[Rubypt-3].genericValue, Func: ast.BareReference{Name: "[]="}, Args: []ast.Node{RubyS[Rubypt-1].genericValue}}
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, tail)
		}
	case 157:
		//line parser.y:783
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 158:
		//line parser.y:790
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 159:
		//line parser.y:794
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 160:
		//line parser.y:801
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 161:
		//line parser.y:808
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 162:
		//line parser.y:815
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 163:
		//line parser.y:822
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 164:
		//line parser.y:825
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 165:
		//line parser.y:827
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 166:
		//line parser.y:830
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 167:
		//line parser.y:832
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 168:
		//line parser.y:835
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 169:
		//line parser.y:837
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 170:
		//line parser.y:840
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 171:
		//line parser.y:842
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 172:
		//line parser.y:844
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 173:
		//line parser.y:847
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 174:
		//line parser.y:849
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 175:
		//line parser.y:851
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 176:
		//line parser.y:854
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 177:
		//line parser.y:856
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 178:
		//line parser.y:858
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 179:
		//line parser.y:861
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 180:
		//line parser.y:863
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 181:
		//line parser.y:865
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 182:
		//line parser.y:867
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 183:
		//line parser.y:868
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 184:
		//line parser.y:869
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 185:
		//line parser.y:870
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 186:
		//line parser.y:873
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 187:
		//line parser.y:882
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 188:
		//line parser.y:891
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 189:
		//line parser.y:900
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 190:
		//line parser.y:909
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 191:
		//line parser.y:918
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 192:
		//line parser.y:926
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 193:
		//line parser.y:927
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 194:
		//line parser.y:929
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 195:
		//line parser.y:931
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 196:
		//line parser.y:933
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 197:
		//line parser.y:935
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 198:
		//line parser.y:937
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 199:
		//line parser.y:939
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{
				Func:   ast.BareReference{Name: "to_proc"},
				Target: RubyS[Rubypt-0].genericValue,
			})
		}
	case 200:
		//line parser.y:946
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 201:
		//line parser.y:948
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 202:
		//line parser.y:950
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 203:
		//line parser.y:952
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{
				Func:   ast.BareReference{Name: "to_proc"},
				Target: RubyS[Rubypt-0].genericValue,
			})
		}
	case 204:
		//line parser.y:960
		{
			RubyVAL.genericValue = ast.Hash{}
		}
	case 205:
		//line parser.y:962
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 206:
		//line parser.y:970
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 207:
		//line parser.y:978
		{
			RubyVAL.genericValue = ast.Block{Body: ast.Nodes{RubyS[Rubypt-2].genericValue}}
		}
	case 208:
		//line parser.y:981
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 209:
		//line parser.y:988
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 210:
		//line parser.y:996
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 211:
		//line parser.y:1003
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 212:
		//line parser.y:1010
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 213:
		//line parser.y:1018
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 214:
		//line parser.y:1020
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 215:
		//line parser.y:1027
		{
			RubyVAL.genericValue = ast.Block{Args: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 216:
		//line parser.y:1031
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-2].genericSlice}
		}
	case 217:
		//line parser.y:1034
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 218:
		//line parser.y:1036
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 219:
		//line parser.y:1038
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 220:
		//line parser.y:1040
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 221:
		//line parser.y:1043
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 222:
		//line parser.y:1050
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 223:
		//line parser.y:1058
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 224:
		//line parser.y:1065
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 225:
		//line parser.y:1072
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 226:
		//line parser.y:1079
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 227:
		//line parser.y:1086
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 228:
		//line parser.y:1093
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 229:
		//line parser.y:1100
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 230:
		//line parser.y:1108
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 231:
		//line parser.y:1115
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 232:
		//line parser.y:1124
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 233:
		//line parser.y:1131
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 234:
		//line parser.y:1138
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 235:
		//line parser.y:1145
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 236:
		//line parser.y:1152
		{
		}
	case 237:
		//line parser.y:1153
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 238:
		//line parser.y:1154
		{
		}
	case 239:
		//line parser.y:1157
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 240:
		//line parser.y:1160
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 241:
		//line parser.y:1168
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 242:
		//line parser.y:1170
		{
			classes := []ast.Class{}
			for _, class := range RubyS[Rubypt-1].genericSlice {
				classes = append(classes, class.(ast.Class))
			}
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Classes: classes,
				},
			}
		}
	case 243:
		//line parser.y:1183
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}

			classes := []ast.Class{}
			for _, class := range RubyS[Rubypt-3].genericSlice {
				classes = append(classes, class.(ast.Class))
			}

			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Var:     RubyS[Rubypt-1].genericValue.(ast.BareReference),
					Classes: classes,
				},
			}
		}
	case 244:
		//line parser.y:1203
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 245:
		//line parser.y:1205
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 246:
		//line parser.y:1208
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 247:
		//line parser.y:1210
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 248:
		//line parser.y:1213
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 249:
		//line parser.y:1215
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 250:
		//line parser.y:1218
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 251:
		//line parser.y:1225
		{
			RubyVAL.genericValue = ast.Yield{}
		}
	case 252:
		//line parser.y:1228
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 253:
		//line parser.y:1236
		{
			RubyVAL.genericValue = ast.Return{}
		}
	case 254:
		//line parser.y:1240
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 255:
		//line parser.y:1242
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 256:
		//line parser.y:1244
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 257:
		//line parser.y:1248
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 258:
		//line parser.y:1250
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 259:
		//line parser.y:1252
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 260:
		//line parser.y:1256
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 261:
		//line parser.y:1265
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 262:
		//line parser.y:1267
		{
			RubyVAL.genericValue = ast.Loop{Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue}, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 263:
		//line parser.y:1270
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 264:
		//line parser.y:1272
		{
		}
	case 265:
		//line parser.y:1274
		{
		}
	case 266:
		//line parser.y:1276
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 267:
		//line parser.y:1278
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 268:
		//line parser.y:1281
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 269:
		//line parser.y:1288
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 270:
		//line parser.y:1296
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 271:
		//line parser.y:1303
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 272:
		//line parser.y:1311
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 273:
		//line parser.y:1319
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 274:
		//line parser.y:1326
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 275:
		//line parser.y:1333
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 276:
		//line parser.y:1340
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 277:
		//line parser.y:1348
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 278:
		//line parser.y:1351
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 279:
		//line parser.y:1353
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 280:
		//line parser.y:1356
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 281:
		//line parser.y:1358
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 282:
		//line parser.y:1361
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 283:
		//line parser.y:1363
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 284:
		//line parser.y:1365
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
