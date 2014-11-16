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

//line parser.y:1154

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 42,
	54, 127,
	-2, 21,
	-1, 107,
	54, 127,
	-2, 125,
	-1, 109,
	10, 92,
	11, 92,
	-2, 196,
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
	54, 127,
	-2, 21,
	-1, 251,
	54, 128,
	-2, 126,
	-1, 261,
	10, 92,
	11, 92,
	-2, 196,
	-1, 405,
	27, 212,
	28, 212,
	-2, 15,
	-1, 406,
	27, 212,
	28, 212,
	-2, 15,
}

const RubyNprod = 242
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 3069

var RubyAct = []int{

	236, 318, 5, 317, 301, 426, 189, 292, 146, 193,
	191, 111, 27, 56, 82, 162, 110, 163, 218, 2,
	3, 245, 245, 120, 280, 87, 242, 151, 22, 118,
	23, 238, 194, 334, 365, 363, 4, 147, 245, 204,
	120, 217, 415, 103, 152, 331, 83, 245, 138, 139,
	80, 79, 106, 108, 96, 97, 87, 137, 82, 143,
	330, 85, 86, 435, 100, 269, 141, 81, 87, 145,
	82, 157, 158, 195, 156, 164, 84, 243, 333, 199,
	95, 82, 241, 167, 196, 96, 97, 99, 136, 82,
	174, 156, 85, 86, 404, 179, 9, 96, 97, 155,
	184, 398, 186, 187, 85, 86, 239, 84, 17, 88,
	89, 90, 197, 82, 238, 414, 308, 405, 406, 84,
	93, 91, 92, 95, 181, 182, 306, 210, 212, 221,
	222, 223, 224, 225, 226, 227, 228, 229, 117, 216,
	203, 205, 209, 207, 96, 97, 245, 201, 416, 232,
	233, 85, 86, 267, 142, 253, 144, 142, 433, 299,
	303, 148, 82, 214, 194, 220, 84, 192, 134, 159,
	160, 161, 240, 149, 245, 135, 370, 277, 87, 248,
	168, 258, 170, 171, 172, 173, 259, 175, 176, 177,
	178, 252, 180, 194, 127, 183, 192, 185, 345, 321,
	267, 87, 356, 268, 357, 195, 202, 96, 97, 206,
	208, 211, 274, 190, 85, 86, 196, 107, 149, 117,
	127, 149, 153, 391, 202, 358, 392, 251, 132, 84,
	96, 97, 278, 95, 195, 133, 149, 85, 86, 298,
	101, 403, 281, 102, 283, 196, 374, 244, 249, 130,
	202, 389, 84, 128, 390, 380, 131, 299, 371, 386,
	129, 313, 149, 311, 242, 100, 117, 379, 377, 299,
	304, 297, 287, 319, 305, 307, 324, 320, 263, 264,
	290, 242, 299, 61, 122, 67, 123, 261, 205, 312,
	120, 276, 277, 441, 339, 438, 437, 271, 436, 314,
	438, 437, 349, 342, 270, 327, 326, 256, 266, 257,
	368, 282, 359, 69, 230, 213, 77, 78, 188, 372,
	260, 70, 71, 72, 73, 74, 372, 87, 165, 432,
	166, 378, 309, 169, 364, 300, 366, 66, 381, 64,
	105, 121, 104, 76, 75, 116, 384, 117, 338, 337,
	247, 286, 237, 387, 388, 246, 96, 97, 202, 315,
	1, 96, 97, 85, 86, 154, 322, 394, 85, 86,
	149, 336, 55, 338, 337, 328, 262, 400, 84, 54,
	53, 52, 51, 84, 50, 34, 33, 32, 31, 46,
	407, 408, 409, 410, 350, 351, 19, 36, 37, 20,
	142, 360, 372, 302, 14, 12, 367, 369, 11, 419,
	420, 421, 21, 361, 18, 423, 10, 24, 16, 13,
	349, 349, 349, 413, 360, 35, 430, 15, 439, 43,
	30, 29, 25, 440, 63, 26, 361, 62, 443, 0,
	422, 349, 0, 444, 445, 349, 349, 349, 446, 0,
	0, 0, 434, 0, 0, 142, 0, 0, 0, 396,
	0, 397, 0, 442, 0, 0, 0, 0, 0, 0,
	0, 124, 0, 396, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 124, 0, 124,
	124, 0, 0, 0, 124, 0, 0, 0, 117, 0,
	0, 0, 124, 124, 124, 0, 0, 0, 0, 0,
	0, 367, 424, 124, 0, 124, 124, 124, 124, 87,
	124, 124, 124, 124, 0, 124, 0, 0, 124, 0,
	124, 0, 0, 28, 0, 0, 0, 0, 0, 124,
	0, 0, 124, 124, 124, 0, 0, 0, 96, 97,
	0, 0, 124, 0, 0, 85, 86, 124, 0, 0,
	88, 89, 90, 98, 0, 0, 0, 0, 0, 0,
	84, 93, 91, 92, 95, 119, 0, 279, 0, 0,
	124, 124, 0, 124, 0, 0, 0, 0, 0, 0,
	0, 119, 0, 119, 119, 0, 0, 0, 119, 124,
	0, 0, 0, 0, 0, 0, 119, 119, 119, 0,
	0, 124, 124, 0, 0, 0, 0, 119, 0, 119,
	119, 119, 119, 0, 119, 119, 119, 119, 0, 119,
	0, 0, 119, 0, 119, 0, 0, 61, 122, 67,
	123, 68, 0, 119, 124, 0, 119, 119, 119, 0,
	0, 0, 0, 0, 0, 0, 119, 0, 0, 0,
	0, 119, 0, 0, 0, 0, 0, 69, 124, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	124, 0, 0, 245, 119, 119, 0, 119, 0, 0,
	376, 124, 124, 64, 0, 65, 0, 76, 75, 124,
	0, 0, 0, 119, 0, 0, 0, 0, 124, 0,
	0, 0, 0, 0, 0, 119, 119, 0, 0, 0,
	0, 61, 122, 67, 123, 109, 0, 115, 120, 0,
	0, 0, 0, 124, 124, 0, 0, 0, 0, 0,
	124, 0, 0, 0, 0, 0, 0, 0, 119, 0,
	0, 69, 0, 0, 77, 78, 0, 124, 113, 70,
	71, 72, 73, 74, 0, 114, 0, 0, 0, 0,
	0, 0, 119, 0, 0, 0, 0, 112, 0, 121,
	0, 76, 75, 0, 119, 0, 0, 0, 124, 0,
	44, 0, 124, 0, 124, 119, 119, 0, 0, 0,
	0, 0, 0, 119, 0, 0, 124, 0, 0, 0,
	0, 0, 119, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 124, 125, 0, 0, 0, 0, 119, 119, 0,
	0, 0, 0, 0, 119, 124, 0, 0, 125, 0,
	125, 125, 0, 0, 0, 125, 0, 0, 0, 0,
	0, 119, 0, 125, 125, 125, 0, 0, 0, 0,
	0, 0, 0, 0, 125, 0, 125, 125, 125, 125,
	0, 125, 125, 125, 125, 0, 125, 0, 0, 125,
	0, 125, 119, 0, 45, 0, 119, 0, 119, 0,
	125, 0, 0, 125, 125, 125, 0, 0, 0, 0,
	119, 0, 0, 125, 0, 0, 0, 0, 125, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 219,
	0, 0, 0, 0, 0, 119, 126, 0, 0, 0,
	0, 125, 125, 0, 125, 0, 0, 0, 0, 119,
	0, 0, 126, 0, 126, 126, 0, 0, 0, 126,
	125, 0, 0, 0, 0, 0, 0, 126, 126, 126,
	0, 0, 125, 125, 0, 0, 0, 0, 126, 140,
	126, 126, 126, 126, 0, 126, 126, 126, 126, 0,
	126, 0, 0, 126, 0, 126, 0, 0, 0, 0,
	0, 0, 0, 0, 126, 125, 0, 126, 126, 126,
	0, 0, 0, 0, 0, 0, 0, 126, 0, 0,
	0, 0, 126, 0, 0, 0, 0, 0, 0, 125,
	0, 0, 0, 0, 0, 0, 198, 0, 200, 0,
	0, 125, 0, 0, 0, 126, 126, 0, 126, 0,
	215, 0, 125, 125, 0, 0, 0, 0, 0, 0,
	125, 0, 0, 0, 126, 0, 0, 0, 231, 125,
	0, 0, 0, 0, 0, 0, 126, 126, 0, 0,
	0, 0, 61, 122, 67, 123, 109, 417, 0, 120,
	0, 0, 0, 0, 125, 125, 0, 0, 0, 0,
	0, 125, 0, 0, 0, 0, 0, 0, 0, 126,
	0, 0, 69, 0, 0, 77, 78, 0, 125, 265,
	70, 71, 72, 73, 74, 0, 0, 272, 0, 0,
	0, 0, 0, 126, 0, 0, 0, 0, 112, 0,
	121, 0, 76, 75, 0, 126, 285, 0, 288, 125,
	0, 0, 0, 125, 0, 125, 126, 126, 0, 0,
	0, 0, 295, 296, 126, 0, 0, 125, 0, 0,
	0, 0, 0, 126, 0, 0, 0, 61, 122, 67,
	123, 68, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 125, 0, 0, 0, 0, 0, 126, 126,
	0, 0, 0, 0, 325, 126, 125, 69, 0, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 126, 245, 340, 0, 0, 0, 0, 344,
	362, 0, 0, 64, 0, 65, 0, 76, 75, 0,
	0, 0, 0, 0, 61, 122, 67, 123, 68, 373,
	0, 0, 0, 126, 0, 0, 0, 126, 0, 126,
	0, 0, 0, 0, 0, 0, 0, 0, 382, 383,
	0, 126, 0, 0, 69, 385, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 393, 0, 395,
	245, 0, 0, 0, 0, 0, 126, 329, 0, 0,
	64, 0, 65, 0, 76, 75, 0, 0, 0, 0,
	126, 402, 0, 0, 0, 0, 0, 231, 0, 0,
	0, 0, 0, 0, 412, 0, 0, 0, 0, 0,
	0, 0, 0, 418, 0, 295, 296, 61, 41, 67,
	42, 68, 0, 0, 0, 38, 429, 352, 428, 427,
	353, 39, 40, 0, 58, 0, 49, 0, 0, 354,
	355, 0, 0, 60, 57, 0, 0, 69, 59, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 347, 348, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 0, 65, 0, 76, 75, 61,
	41, 67, 42, 68, 0, 0, 0, 38, 425, 352,
	428, 427, 353, 39, 40, 0, 58, 0, 49, 0,
	0, 354, 355, 0, 0, 60, 57, 0, 0, 69,
	59, 0, 77, 78, 0, 0, 0, 70, 71, 72,
	73, 74, 0, 0, 0, 347, 348, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 0, 65, 0, 76,
	75, 61, 41, 67, 42, 68, 0, 0, 0, 38,
	431, 352, 0, 0, 353, 39, 40, 0, 58, 0,
	49, 0, 0, 354, 355, 0, 0, 60, 57, 0,
	0, 69, 59, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 0, 0, 0, 347, 348, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 0, 65,
	0, 76, 75, 61, 41, 67, 42, 68, 0, 0,
	0, 38, 346, 352, 0, 0, 353, 39, 40, 0,
	58, 0, 49, 0, 0, 354, 355, 0, 0, 60,
	57, 0, 0, 69, 59, 0, 77, 78, 0, 0,
	0, 70, 71, 72, 73, 74, 0, 0, 0, 347,
	348, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	0, 65, 0, 76, 75, 61, 41, 67, 42, 68,
	0, 0, 0, 38, 341, 47, 294, 293, 48, 39,
	40, 0, 58, 0, 49, 0, 0, 0, 0, 0,
	0, 60, 57, 0, 0, 69, 59, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 0, 0,
	0, 234, 235, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 0, 65, 0, 76, 75, 61, 41, 67,
	42, 68, 0, 0, 0, 38, 291, 47, 294, 293,
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
	352, 0, 0, 353, 39, 40, 0, 58, 0, 49,
	0, 0, 354, 355, 0, 0, 60, 57, 0, 0,
	69, 59, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 0, 0, 0, 347, 348, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 65, 0,
	76, 75, 61, 41, 67, 42, 68, 0, 0, 0,
	38, 399, 47, 0, 0, 48, 39, 40, 0, 58,
	0, 49, 299, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 234, 235,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 0,
	65, 0, 76, 75, 61, 41, 67, 42, 68, 0,
	0, 0, 38, 323, 47, 0, 0, 48, 39, 40,
	0, 58, 0, 49, 299, 0, 0, 0, 0, 0,
	60, 57, 0, 0, 69, 59, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 0, 0, 0,
	234, 235, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 0, 65, 0, 76, 75, 61, 41, 67, 42,
	68, 0, 0, 0, 38, 316, 47, 0, 0, 48,
	39, 40, 0, 58, 0, 49, 299, 0, 0, 0,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 234, 235, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 0, 76, 75, 61, 41,
	67, 42, 68, 0, 0, 0, 38, 411, 47, 0,
	0, 48, 39, 40, 0, 58, 0, 49, 0, 0,
	0, 0, 0, 0, 60, 57, 0, 0, 69, 59,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 234, 235, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 0, 65, 0, 76, 75,
	61, 41, 67, 42, 68, 0, 0, 0, 38, 375,
	47, 0, 0, 48, 39, 40, 0, 58, 0, 49,
	0, 0, 0, 0, 0, 0, 60, 57, 0, 0,
	69, 59, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 0, 0, 0, 234, 235, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 65, 0,
	76, 75, 61, 41, 67, 42, 68, 0, 0, 0,
	38, 343, 47, 0, 0, 48, 39, 40, 0, 58,
	0, 49, 0, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 234, 235,
	0, 61, 41, 67, 42, 68, 0, 0, 64, 38,
	65, 47, 76, 75, 48, 39, 40, 0, 58, 0,
	49, 0, 0, 0, 0, 0, 0, 60, 57, 0,
	0, 69, 59, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 0, 0, 0, 234, 235, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 0, 65,
	335, 76, 75, 61, 41, 67, 42, 68, 0, 0,
	0, 38, 332, 47, 0, 0, 48, 39, 40, 0,
	58, 0, 49, 0, 0, 0, 0, 0, 0, 60,
	57, 0, 0, 69, 59, 0, 77, 78, 0, 0,
	0, 70, 71, 72, 73, 74, 0, 0, 0, 234,
	235, 0, 61, 41, 67, 42, 68, 0, 0, 64,
	38, 65, 47, 76, 75, 48, 39, 40, 0, 58,
	0, 49, 0, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 234, 235,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 0,
	65, 289, 76, 75, 61, 41, 67, 42, 68, 0,
	0, 0, 38, 284, 47, 0, 0, 48, 39, 40,
	0, 58, 0, 49, 0, 0, 0, 0, 0, 0,
	60, 57, 0, 0, 69, 59, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 0, 0, 0,
	234, 235, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 0, 65, 0, 76, 75, 61, 41, 67, 42,
	68, 0, 0, 0, 38, 275, 47, 0, 0, 48,
	39, 40, 0, 58, 0, 49, 0, 0, 0, 0,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 234, 235, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 0, 76, 75, 61, 41,
	67, 42, 68, 0, 0, 0, 38, 273, 47, 0,
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
	0, 0, 0, 254, 0, 61, 41, 67, 42, 68,
	0, 0, 64, 38, 65, 47, 76, 75, 48, 39,
	40, 0, 58, 0, 49, 0, 0, 0, 0, 0,
	0, 60, 57, 0, 0, 69, 59, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 61, 122,
	67, 123, 109, 0, 0, 120, 0, 0, 0, 0,
	0, 64, 0, 65, 0, 76, 75, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 0,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 61, 122, 67, 123, 68, 310, 0, 0, 0,
	0, 0, 0, 0, 112, 0, 121, 0, 76, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 0, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 0, 0, 0, 245, 61, 122,
	67, 123, 68, 0, 0, 0, 0, 64, 0, 65,
	0, 76, 75, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 0,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 61, 122, 67, 123, 68, 0, 0, 120, 0,
	0, 150, 0, 0, 64, 0, 65, 0, 76, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 0, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 61, 250, 67, 123, 68, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 0, 121,
	0, 76, 75, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 0, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 0, 0, 0,
	245, 61, 122, 67, 123, 109, 0, 0, 120, 0,
	64, 0, 65, 0, 76, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 0, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 61, 122, 67, 123, 68, 0,
	0, 0, 0, 0, 0, 0, 0, 112, 0, 121,
	0, 76, 75, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 0, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 61, 401, 67,
	123, 68, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 0, 65, 0, 76, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 0, 0,
	77, 78, 0, 87, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 0, 0, 0, 0, 94, 0, 0,
	0, 0, 0, 64, 83, 65, 0, 76, 75, 0,
	0, 0, 96, 97, 0, 0, 0, 0, 0, 85,
	86, 0, 0, 0, 88, 89, 90, 98, 0, 0,
	0, 0, 0, 0, 84, 93, 91, 92, 95,
}
var RubyPact = []int{

	-32, 1691, -1000, -1000, -1000, -1, -1000, -1000, -1000, 3009,
	-1000, -1000, -1000, 69, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 225, -12,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 336, 209,
	209, 716, 211, 207, 186, 126, 46, 2620, 2620, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 2929, 2620, 2929,
	2929, -1000, -1000, -1000, 2753, -1000, -10, 213, -1000, 11,
	2620, 2620, 2929, 2929, 2929, 9, 322, -1000, -1000, -1000,
	-1000, -1000, 2620, 2929, 327, 2929, 2929, 2929, 2929, 2620,
	2929, 2929, 2929, 2929, 2620, 2929, -1000, -1000, 2929, 2620,
	2929, 2620, 2620, 312, 158, 187, 39, -1000, -1000, 2753,
	11, 28, 2753, 2929, 2929, 309, 152, 197, -1000, 24,
	-15, -15, 2886, 185, -1000, -1000, -1000, 2753, 2620, 2620,
	2620, 2620, 2620, 2620, 2620, 2620, 2620, 308, 74, 98,
	2522, 103, 197, 55, 197, 103, 20, 15, 64, -1000,
	2929, 2839, 219, 2753, 2571, -1000, -15, 74, 74, 197,
	197, 197, -1000, -1000, 301, -1000, -1000, 74, 197, 278,
	197, 197, 197, 197, 74, 197, 197, 197, 197, 74,
	323, 2706, 2706, 197, 74, 197, 74, 74, -1000, -1000,
	302, 142, 26, -1000, 23, 298, 291, -1000, 2473, 209,
	2411, 281, 64, -1000, -1000, -1000, 515, -38, 21, -1000,
	-1000, 174, -1000, -1000, 2796, 2349, -1000, 266, -1000, 2287,
	270, 74, 74, 74, 74, 74, 74, 74, 74, 74,
	-1000, 1642, -1000, -1000, -1000, -1000, 74, 257, 2929, -1000,
	123, -1000, -1000, -1000, 197, -1000, 115, 105, -4, 328,
	2663, -1000, 253, 74, -1000, -1000, -1000, -1000, 28, 11,
	2620, 2753, 2929, 197, 197, 1941, 187, 26, 189, 2929,
	-1000, -1000, 1879, -1000, -1000, -1000, 11, -1000, 1239, 18,
	-1000, -13, 197, -1000, -1000, 2238, 22, -1000, 2176, -1000,
	-1000, -1000, 357, 2620, -1000, 1580, 2127, -1000, -1000, 190,
	197, 1518, 188, 2929, 1172, -29, -1000, -30, -1000, 2620,
	2929, -1000, -1000, 74, 166, 197, -1000, 244, -1000, -1000,
	-1000, -1000, 197, -1000, 232, 2065, -1000, 632, 197, 262,
	2620, 261, -1000, -1000, 249, -1000, -1000, 2620, -1000, 74,
	2522, -1000, 332, -1000, 2522, 255, -1000, -1000, -1000, 74,
	-1000, -1000, 2620, 2620, 236, 208, -1000, -1000, 2929, 103,
	64, -1000, 2929, -1000, 2706, -1000, 95, 3009, 74, 197,
	-1000, -1000, -1000, 1817, -1000, -1000, 2972, -1000, 74, -1000,
	-1000, 74, 2522, 2522, -1000, 2522, 235, 43, 66, 2620,
	2620, 2620, 2620, 2003, 103, 2522, 197, 111, -11, -1000,
	134, 1077, 2522, -1000, -1000, -1000, -1000, 74, 74, 74,
	74, -1000, 2522, -4, 2620, 2929, -1000, -1000, 2522, 1394,
	1332, 1456, -4, 147, 52, -1000, 284, 2620, -1000, -1000,
	279, -1000, -1000, -1000, -4, -1000, -1000, 2620, -1000, 74,
	1755, -1000, -4, 74, 1755, 1755, 1755,
}
var RubyPgo = []int{

	0, 0, 437, 435, 30, 29, 434, 432, 431, 894,
	430, 1, 13, 427, 425, 419, 418, 96, 417, 790,
	533, 416, 414, 412, 108, 408, 9, 429, 405, 404,
	12, 403, 399, 398, 397, 28, 396, 395, 394, 5,
	389, 388, 387, 386, 385, 384, 382, 381, 380, 379,
	372, 929, 365, 3, 16, 18, 7, 360, 6, 355,
	4, 352, 37, 351, 8, 350, 345, 11, 10, 337,
	329, 27,
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
	20, 20, 20, 20, 20, 20, 54, 54, 54, 54,
	64, 64, 62, 62, 62, 62, 62, 62, 62, 67,
	67, 67, 67, 67, 66, 66, 66, 21, 21, 21,
	21, 21, 21, 58, 58, 68, 68, 68, 26, 26,
	26, 26, 25, 25, 28, 30, 30, 69, 69, 15,
	15, 15, 15, 15, 15, 15, 15, 29, 29, 29,
	29, 29, 29, 9, 9, 27, 27, 19, 19, 40,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
	50, 2, 6, 7, 7, 3, 3, 59, 59, 59,
	59, 65, 65, 65, 5, 5, 5, 5, 55, 63,
	63, 63, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 56, 56, 56, 56, 52, 52, 52, 8,
	16, 11, 11, 11, 61, 61, 53, 53, 22, 23,
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
	2, 3, 3, 4, 4, 5, 3, 5, 2, 3,
	3, 3, 3, 4, 4, 6, 3, 7, 1, 5,
	1, 3, 0, 1, 1, 2, 4, 4, 5, 1,
	1, 4, 2, 5, 1, 3, 3, 5, 6, 7,
	8, 5, 6, 1, 3, 0, 1, 3, 1, 2,
	3, 2, 4, 6, 4, 1, 3, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 2, 2, 2, 3, 3, 3,
	3, 2, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 1, 1, 3, 3, 5, 5, 0, 1, 3,
	7, 3, 7, 8, 3, 4, 4, 3, 3, 0,
	1, 3, 4, 5, 3, 3, 3, 3, 3, 5,
	6, 5, 4, 3, 3, 2, 0, 2, 2, 3,
	4, 2, 3, 5, 0, 2, 1, 2, 2, 2,
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
	51, 68, 15, 25, 55, 40, 41, 4, 45, 46,
	47, 57, 58, 56, 18, 59, 33, 34, 48, 18,
	40, 15, 18, 55, 6, 4, -30, 8, -30, 9,
	-54, -67, 61, 42, 49, 11, -66, -17, -5, -20,
	12, 63, 6, 8, -27, -19, -9, 9, 42, 49,
	42, 49, 42, 49, 42, 49, 42, 11, -1, -1,
	-51, -64, -17, -1, -17, -64, -64, -62, -17, -24,
	58, -71, 54, 9, -52, -5, 63, -1, -1, -17,
	-17, -17, 6, 8, 66, 6, 8, -1, -17, 6,
	-17, -17, -17, -17, -1, -17, -17, -17, -17, -1,
	-17, -71, -71, -17, -1, -17, -1, -1, 6, -58,
	55, -68, 9, -26, 6, 47, 58, -58, -51, 40,
	-51, -62, -17, -5, 11, -5, -17, -4, -17, -35,
	-12, -17, -12, 6, 11, -51, -55, 56, -55, -51,
	-62, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	6, -51, 51, 52, 51, 52, -1, -61, 11, 51,
	-71, 62, 11, 62, -17, 51, -59, -65, -71, -17,
	6, 8, -62, -1, 52, 10, 6, 8, -67, -54,
	42, 9, 53, -17, -17, -51, 6, 11, -68, 42,
	6, 6, -51, 14, -30, 14, 10, 11, -71, 62,
	62, -71, -17, -5, 14, -51, -63, 6, -51, 64,
	10, 14, -56, 17, 16, -51, -51, 14, -11, 25,
	-17, -60, -31, 37, -71, -71, 11, -71, 11, 4,
	53, 10, -5, -1, -62, -17, 14, -53, -11, -58,
	-26, 10, -17, 14, -53, -51, -5, -71, -17, 58,
	42, 58, 14, 56, 11, 64, 14, 17, 16, -1,
	-51, 14, -56, 14, -51, 8, 14, 51, 52, -1,
	-38, -37, 15, 18, 27, 28, 14, 16, 37, -64,
	-17, -24, 58, 64, -71, 64, -71, -17, -1, -17,
	10, 14, -11, -51, 14, 14, 58, 6, -1, 6,
	6, -1, -51, -51, 14, -51, 4, -1, -1, 15,
	18, 15, 18, -51, -64, -51, -17, -17, 6, 14,
	-53, 6, -51, 6, 51, 51, 52, -1, -1, -1,
	-1, 14, -51, -71, 4, 53, 14, 10, -51, -60,
	-60, -60, -71, -1, -17, 14, -39, 17, 16, 14,
	-39, 14, -70, 11, -71, 11, 14, 17, 16, -1,
	-60, 14, -71, -1, -60, -60, -60,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 48,
	49, 50, 51, 52, 53, 54, 55, 56, 57, 58,
	59, 60, 61, 62, 27, 28, 29, 30, 31, 32,
	33, 34, 35, 36, 37, 38, 39, 40, 0, 0,
	0, 20, -2, 22, 23, 24, 0, 0, 0, 15,
	41, 42, 43, 44, 45, 46, 47, 0, 0, 0,
	0, 19, 25, 26, 92, 13, 0, 67, 196, 0,
	0, 0, 0, 0, 0, 0, 0, 161, 162, 5,
	6, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 13, 13, 0, 0,
	0, 0, 0, 0, 115, 115, 15, -2, 15, -2,
	70, 78, 92, 0, 0, 0, 88, 99, 100, 31,
	15, -2, 20, -2, 22, 23, 24, 92, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 15, 0,
	204, 208, 90, 0, 13, 209, 0, 0, 90, 94,
	0, 13, 0, 92, 0, 236, 15, 151, 152, 153,
	154, 64, 145, 146, 0, 143, 144, 184, 63, 72,
	79, 81, 82, 155, 156, 157, 158, 159, 160, 186,
	0, 0, 0, 241, 188, 80, 185, 187, 76, 15,
	0, 113, 115, 116, 118, 0, 0, 15, 0, 0,
	0, 0, 93, 71, 13, 102, 90, 0, 129, 130,
	131, 137, 138, 149, 13, 0, 15, 179, 15, 0,
	0, 132, 139, 133, 140, 134, 141, 135, 142, 136,
	150, 0, 15, 15, 16, 17, 18, 0, 0, 212,
	0, 163, 13, 164, 95, 14, 13, 13, 168, 0,
	20, -2, 0, 197, 198, 199, 147, 148, 73, 74,
	0, -2, 0, 234, 235, 0, 115, 0, 0, 0,
	119, 121, 0, 122, 15, 124, 65, 13, 0, 83,
	84, 0, 105, 106, 174, 0, 0, 180, 0, 177,
	69, 182, 0, 0, 15, 0, 0, 200, 205, 15,
	91, 0, 0, 0, 0, 0, 13, 0, 13, 0,
	0, 68, 75, 77, 0, 210, 107, 0, 206, 15,
	117, 114, 120, 111, 0, 0, 66, 0, 101, 0,
	0, 0, 175, 178, 0, 176, 183, 0, 15, 15,
	195, 189, 0, 191, 201, 15, 211, 213, 214, 215,
	216, 217, 0, 0, 219, 222, 237, 15, 0, 15,
	96, 97, 0, 165, 0, 166, 0, 48, 169, 171,
	86, 108, 207, 0, 112, 123, 0, 103, 85, 89,
	181, 15, 193, 194, 190, 202, 0, 15, 0, 0,
	0, 0, 0, 0, 15, 13, 98, 0, 0, 109,
	0, 20, 192, 15, 212, -2, -2, 220, 221, 223,
	224, 238, 13, 239, 0, 0, 110, 87, 203, 0,
	0, 0, 240, 11, 13, 225, 0, 0, 212, 227,
	0, 229, 170, 12, 172, 13, 226, 0, 212, 212,
	218, 228, 173, 212, 218, 218, 218,
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
		//line parser.y:408
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 86:
		//line parser.y:417
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 87:
		//line parser.y:419
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 88:
		//line parser.y:421
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 89:
		//line parser.y:423
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 90:
		//line parser.y:426
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 91:
		//line parser.y:428
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 92:
		//line parser.y:430
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 93:
		//line parser.y:432
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 94:
		//line parser.y:434
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 95:
		//line parser.y:436
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{Func: ast.BareReference{Name: "to_proc"}, Target: RubyS[Rubypt-0].genericValue})
		}
	case 96:
		//line parser.y:438
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 97:
		//line parser.y:440
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 98:
		//line parser.y:442
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{
				Func:   ast.BareReference{Name: "to_proc"},
				Target: RubyS[Rubypt-0].genericValue,
			})
		}
	case 99:
		//line parser.y:451
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 100:
		//line parser.y:453
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 101:
		//line parser.y:455
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 102:
		//line parser.y:457
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 103:
		//line parser.y:459
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 104:
		//line parser.y:462
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 105:
		//line parser.y:464
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 106:
		//line parser.y:466
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 107:
		//line parser.y:470
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 108:
		//line parser.y:478
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 109:
		//line parser.y:487
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 110:
		//line parser.y:496
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 111:
		//line parser.y:506
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 112:
		//line parser.y:514
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 113:
		//line parser.y:525
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 114:
		//line parser.y:527
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 115:
		//line parser.y:529
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 116:
		//line parser.y:531
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 117:
		//line parser.y:533
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 118:
		//line parser.y:536
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 119:
		//line parser.y:538
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 120:
		//line parser.y:540
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 121:
		//line parser.y:542
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 122:
		//line parser.y:546
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 123:
		//line parser.y:553
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 124:
		//line parser.y:563
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 125:
		//line parser.y:572
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 126:
		//line parser.y:578
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 127:
		//line parser.y:586
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 128:
		//line parser.y:590
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 129:
		//line parser.y:595
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 130:
		//line parser.y:602
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 131:
		//line parser.y:606
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 132:
		//line parser.y:610
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 133:
		//line parser.y:617
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 134:
		//line parser.y:624
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 135:
		//line parser.y:631
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 136:
		//line parser.y:638
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
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
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 139:
		//line parser.y:657
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 140:
		//line parser.y:664
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 141:
		//line parser.y:671
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 142:
		//line parser.y:678
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 143:
		//line parser.y:686
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 144:
		//line parser.y:688
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 145:
		//line parser.y:691
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 146:
		//line parser.y:693
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 147:
		//line parser.y:696
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 148:
		//line parser.y:698
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 149:
		//line parser.y:701
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 150:
		//line parser.y:703
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 151:
		//line parser.y:705
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 152:
		//line parser.y:706
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 153:
		//line parser.y:707
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 154:
		//line parser.y:708
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 155:
		//line parser.y:711
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 156:
		//line parser.y:720
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 157:
		//line parser.y:729
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 158:
		//line parser.y:738
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 159:
		//line parser.y:747
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 160:
		//line parser.y:756
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 161:
		//line parser.y:764
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 162:
		//line parser.y:765
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 163:
		//line parser.y:767
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 164:
		//line parser.y:769
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 165:
		//line parser.y:772
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 166:
		//line parser.y:780
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 167:
		//line parser.y:788
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 168:
		//line parser.y:790
		{
		}
	case 169:
		//line parser.y:792
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 170:
		//line parser.y:799
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 171:
		//line parser.y:807
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 172:
		//line parser.y:814
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 173:
		//line parser.y:821
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 174:
		//line parser.y:829
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 175:
		//line parser.y:831
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 176:
		//line parser.y:838
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 177:
		//line parser.y:845
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 178:
		//line parser.y:848
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 179:
		//line parser.y:850
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 180:
		//line parser.y:852
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 181:
		//line parser.y:854
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 182:
		//line parser.y:857
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 183:
		//line parser.y:864
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 184:
		//line parser.y:872
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 185:
		//line parser.y:879
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 186:
		//line parser.y:886
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 187:
		//line parser.y:893
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 188:
		//line parser.y:900
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 189:
		//line parser.y:907
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 190:
		//line parser.y:914
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 191:
		//line parser.y:922
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 192:
		//line parser.y:931
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 193:
		//line parser.y:938
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 194:
		//line parser.y:945
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 195:
		//line parser.y:952
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 196:
		//line parser.y:959
		{
		}
	case 197:
		//line parser.y:960
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 198:
		//line parser.y:961
		{
		}
	case 199:
		//line parser.y:964
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 200:
		//line parser.y:967
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 201:
		//line parser.y:975
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 202:
		//line parser.y:977
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 203:
		//line parser.y:986
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
	case 204:
		//line parser.y:1001
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 205:
		//line parser.y:1003
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 206:
		//line parser.y:1006
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 207:
		//line parser.y:1008
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 208:
		//line parser.y:1011
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 209:
		//line parser.y:1020
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 210:
		//line parser.y:1029
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 211:
		//line parser.y:1038
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 212:
		//line parser.y:1041
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 213:
		//line parser.y:1043
		{
		}
	case 214:
		//line parser.y:1045
		{
		}
	case 215:
		//line parser.y:1047
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 216:
		//line parser.y:1049
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 217:
		//line parser.y:1051
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 218:
		//line parser.y:1053
		{
		}
	case 219:
		//line parser.y:1055
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 220:
		//line parser.y:1057
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 221:
		//line parser.y:1059
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 222:
		//line parser.y:1061
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 223:
		//line parser.y:1063
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 224:
		//line parser.y:1065
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 225:
		//line parser.y:1068
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 226:
		//line parser.y:1075
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 227:
		//line parser.y:1083
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 228:
		//line parser.y:1090
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 229:
		//line parser.y:1098
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 230:
		//line parser.y:1106
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 231:
		//line parser.y:1113
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 232:
		//line parser.y:1120
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 233:
		//line parser.y:1127
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 234:
		//line parser.y:1135
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 235:
		//line parser.y:1138
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 236:
		//line parser.y:1140
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 237:
		//line parser.y:1143
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 238:
		//line parser.y:1145
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 239:
		//line parser.y:1148
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 240:
		//line parser.y:1150
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 241:
		//line parser.y:1152
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
