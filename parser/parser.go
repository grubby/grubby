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

//line parser.y:1127

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 41,
	54, 122,
	-2, 22,
	-1, 105,
	54, 122,
	-2, 120,
	-1, 107,
	10, 89,
	11, 89,
	-2, 190,
	-1, 119,
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
	-1, 121,
	54, 122,
	-2, 22,
	-1, 245,
	54, 123,
	-2, 121,
	-1, 255,
	10, 89,
	11, 89,
	-2, 190,
	-1, 395,
	27, 206,
	28, 206,
	-2, 15,
	-1, 396,
	27, 206,
	28, 206,
	-2, 15,
}

const RubyNprod = 236
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 2853

var RubyAct = []int{

	231, 308, 5, 283, 309, 187, 416, 144, 191, 116,
	27, 292, 148, 17, 145, 189, 109, 108, 212, 24,
	60, 155, 239, 156, 2, 3, 239, 239, 118, 55,
	211, 149, 85, 323, 101, 356, 405, 161, 294, 162,
	354, 4, 239, 237, 322, 261, 92, 136, 137, 104,
	106, 233, 239, 76, 77, 97, 81, 98, 141, 81,
	197, 94, 95, 81, 139, 192, 326, 143, 83, 84,
	388, 233, 135, 86, 87, 88, 96, 147, 152, 153,
	75, 74, 166, 82, 91, 89, 90, 93, 172, 154,
	157, 158, 159, 177, 238, 79, 78, 163, 182, 394,
	184, 185, 236, 134, 192, 192, 193, 190, 190, 195,
	125, 325, 80, 132, 406, 239, 201, 194, 81, 407,
	133, 147, 199, 81, 147, 290, 85, 215, 216, 217,
	218, 219, 220, 221, 222, 223, 81, 210, 365, 147,
	214, 204, 206, 126, 130, 193, 193, 361, 269, 290,
	127, 131, 247, 188, 234, 235, 194, 194, 299, 395,
	396, 242, 83, 84, 147, 246, 297, 86, 87, 88,
	96, 382, 226, 227, 383, 128, 259, 82, 91, 89,
	90, 93, 129, 208, 252, 253, 60, 40, 66, 41,
	67, 362, 312, 259, 37, 333, 46, 285, 284, 47,
	38, 39, 290, 57, 99, 48, 260, 100, 266, 302,
	237, 288, 59, 56, 270, 125, 68, 58, 274, 76,
	77, 272, 290, 85, 69, 70, 71, 72, 73, 98,
	425, 150, 228, 229, 348, 404, 349, 289, 380, 423,
	85, 381, 63, 81, 64, 337, 75, 74, 281, 237,
	295, 105, 85, 296, 298, 304, 250, 350, 251, 83,
	84, 245, 313, 303, 310, 9, 315, 96, 311, 147,
	305, 83, 84, 164, 82, 165, 83, 84, 317, 96,
	320, 103, 318, 102, 96, 331, 82, 393, 83, 84,
	334, 82, 391, 341, 377, 93, 96, 300, 268, 269,
	371, 359, 351, 82, 370, 431, 115, 428, 427, 353,
	355, 368, 357, 363, 426, 278, 428, 427, 263, 262,
	363, 258, 140, 369, 142, 140, 224, 422, 207, 146,
	372, 186, 353, 83, 84, 375, 167, 330, 329, 160,
	65, 96, 114, 85, 241, 378, 379, 277, 82, 168,
	169, 170, 171, 232, 173, 174, 175, 176, 385, 178,
	179, 180, 181, 240, 183, 328, 390, 330, 329, 1,
	151, 54, 53, 200, 52, 51, 202, 203, 205, 83,
	84, 397, 398, 399, 400, 50, 115, 96, 49, 34,
	33, 200, 256, 32, 82, 363, 31, 45, 342, 403,
	343, 19, 20, 21, 230, 413, 409, 410, 411, 22,
	341, 341, 341, 293, 243, 412, 200, 420, 429, 14,
	12, 11, 23, 18, 10, 16, 13, 424, 433, 36,
	430, 341, 15, 115, 30, 341, 341, 341, 432, 29,
	434, 435, 25, 62, 35, 436, 26, 61, 60, 120,
	66, 121, 107, 0, 113, 118, 0, 0, 0, 0,
	0, 0, 230, 0, 230, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 273, 230, 0, 0, 68, 230,
	0, 76, 77, 0, 0, 111, 69, 70, 71, 72,
	73, 230, 112, 0, 0, 0, 85, 0, 0, 291,
	0, 0, 0, 0, 110, 0, 119, 0, 75, 74,
	115, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 200, 306, 230, 0, 0, 0, 0, 0, 0,
	230, 0, 83, 84, 0, 0, 319, 86, 87, 88,
	96, 0, 230, 0, 0, 230, 0, 82, 91, 89,
	90, 93, 230, 230, 271, 0, 0, 0, 0, 0,
	140, 352, 0, 0, 0, 0, 358, 360, 0, 0,
	0, 0, 0, 60, 40, 66, 41, 67, 0, 0,
	0, 37, 230, 46, 352, 0, 47, 38, 39, 0,
	57, 0, 48, 0, 0, 0, 0, 0, 230, 59,
	56, 0, 230, 68, 58, 42, 76, 77, 0, 0,
	0, 69, 70, 71, 72, 73, 140, 0, 0, 6,
	7, 387, 0, 0, 0, 0, 0, 0, 0, 63,
	230, 64, 0, 75, 74, 0, 8, 0, 0, 230,
	230, 0, 230, 0, 0, 0, 122, 0, 0, 0,
	230, 0, 230, 0, 0, 0, 0, 0, 230, 0,
	0, 0, 122, 0, 122, 122, 0, 0, 230, 122,
	358, 414, 0, 0, 230, 122, 122, 122, 122, 122,
	0, 60, 120, 66, 121, 67, 0, 0, 118, 122,
	122, 122, 122, 0, 122, 122, 122, 122, 0, 122,
	122, 122, 122, 0, 122, 0, 0, 43, 0, 0,
	0, 68, 0, 122, 76, 77, 122, 122, 122, 69,
	70, 71, 72, 73, 0, 0, 122, 239, 0, 0,
	0, 122, 0, 0, 321, 0, 0, 63, 0, 119,
	0, 75, 74, 0, 0, 0, 0, 0, 123, 0,
	0, 0, 0, 0, 122, 0, 122, 0, 0, 0,
	0, 0, 0, 0, 123, 0, 123, 123, 0, 0,
	0, 123, 0, 122, 0, 0, 0, 123, 123, 123,
	123, 123, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 123, 123, 123, 123, 0, 123, 123, 123, 123,
	0, 123, 123, 123, 123, 0, 123, 60, 120, 66,
	121, 67, 0, 0, 122, 123, 0, 0, 123, 123,
	123, 0, 0, 0, 0, 0, 0, 0, 123, 0,
	0, 0, 0, 123, 0, 0, 0, 68, 0, 122,
	76, 77, 0, 0, 0, 69, 70, 71, 72, 73,
	122, 0, 0, 239, 0, 0, 123, 0, 123, 0,
	367, 122, 122, 63, 0, 64, 0, 75, 74, 0,
	0, 0, 0, 0, 0, 123, 122, 0, 0, 60,
	40, 66, 41, 67, 0, 0, 0, 37, 282, 46,
	285, 284, 47, 38, 39, 0, 57, 0, 48, 0,
	122, 122, 0, 0, 0, 59, 56, 122, 0, 68,
	58, 0, 76, 77, 0, 0, 123, 69, 70, 71,
	72, 73, 0, 0, 122, 228, 229, 0, 0, 0,
	0, 0, 0, 0, 0, 63, 0, 64, 0, 75,
	74, 123, 0, 0, 0, 0, 0, 0, 44, 0,
	0, 0, 123, 0, 0, 0, 122, 0, 0, 0,
	0, 122, 0, 123, 123, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 123, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 124,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 123, 123, 0, 124, 0, 124, 124, 123,
	0, 122, 124, 0, 0, 0, 0, 0, 124, 124,
	124, 124, 124, 0, 0, 0, 123, 0, 0, 0,
	0, 0, 124, 124, 124, 124, 0, 124, 124, 124,
	124, 0, 124, 124, 124, 124, 0, 124, 0, 0,
	28, 0, 0, 0, 0, 0, 124, 0, 123, 124,
	124, 124, 0, 123, 0, 0, 0, 0, 0, 124,
	0, 0, 0, 0, 124, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 213,
	0, 117, 0, 0, 0, 0, 0, 124, 0, 124,
	0, 0, 0, 0, 0, 0, 0, 117, 0, 117,
	117, 0, 0, 123, 117, 0, 124, 0, 0, 0,
	0, 0, 0, 0, 117, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 117, 117, 117, 117, 138, 117,
	117, 117, 117, 0, 117, 117, 117, 117, 0, 117,
	0, 0, 0, 0, 0, 0, 0, 124, 117, 0,
	0, 117, 117, 117, 0, 0, 0, 0, 0, 0,
	0, 117, 0, 0, 0, 0, 117, 0, 0, 0,
	0, 0, 124, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 124, 196, 0, 198, 0, 0, 117,
	0, 117, 0, 0, 124, 124, 0, 0, 209, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 117, 124,
	0, 0, 0, 0, 0, 0, 225, 0, 0, 0,
	60, 120, 66, 121, 107, 0, 0, 118, 0, 0,
	0, 0, 0, 124, 124, 0, 0, 0, 0, 0,
	124, 0, 0, 0, 0, 0, 0, 0, 0, 117,
	68, 0, 0, 76, 77, 0, 0, 124, 69, 70,
	71, 72, 73, 0, 0, 0, 0, 257, 301, 0,
	0, 0, 0, 0, 117, 264, 110, 0, 119, 0,
	75, 74, 0, 0, 0, 117, 0, 0, 0, 124,
	276, 0, 279, 0, 124, 0, 117, 117, 0, 0,
	0, 0, 0, 0, 0, 0, 286, 287, 0, 0,
	0, 117, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 117, 117, 0, 0, 0,
	0, 0, 117, 0, 124, 0, 316, 60, 40, 66,
	41, 67, 0, 0, 0, 37, 389, 46, 0, 117,
	47, 38, 39, 0, 57, 332, 48, 290, 0, 0,
	336, 0, 0, 59, 56, 0, 0, 68, 58, 0,
	76, 77, 0, 0, 0, 69, 70, 71, 72, 73,
	364, 117, 0, 228, 229, 0, 117, 0, 0, 0,
	0, 0, 0, 63, 0, 64, 0, 75, 74, 0,
	373, 374, 0, 0, 0, 0, 0, 376, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 384,
	0, 386, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 117, 0, 0, 0,
	0, 0, 392, 0, 0, 0, 0, 0, 225, 0,
	0, 0, 0, 0, 0, 402, 0, 0, 0, 0,
	0, 0, 0, 408, 0, 286, 287, 60, 40, 66,
	41, 67, 0, 0, 0, 37, 419, 344, 418, 417,
	345, 38, 39, 0, 57, 0, 48, 0, 0, 346,
	347, 0, 0, 59, 56, 0, 0, 68, 58, 0,
	76, 77, 0, 0, 0, 69, 70, 71, 72, 73,
	0, 0, 0, 339, 340, 0, 0, 0, 0, 0,
	0, 0, 0, 63, 0, 64, 0, 75, 74, 60,
	40, 66, 41, 67, 0, 0, 0, 37, 415, 344,
	418, 417, 345, 38, 39, 0, 57, 0, 48, 0,
	0, 346, 347, 0, 0, 59, 56, 0, 0, 68,
	58, 0, 76, 77, 0, 0, 0, 69, 70, 71,
	72, 73, 0, 0, 0, 339, 340, 0, 0, 0,
	0, 0, 0, 0, 0, 63, 0, 64, 0, 75,
	74, 60, 40, 66, 41, 67, 0, 0, 0, 37,
	421, 344, 0, 0, 345, 38, 39, 0, 57, 0,
	48, 0, 0, 346, 347, 0, 0, 59, 56, 0,
	0, 68, 58, 0, 76, 77, 0, 0, 0, 69,
	70, 71, 72, 73, 0, 0, 0, 339, 340, 0,
	0, 0, 0, 0, 0, 0, 0, 63, 0, 64,
	0, 75, 74, 60, 40, 66, 41, 67, 0, 0,
	0, 37, 338, 344, 0, 0, 345, 38, 39, 0,
	57, 0, 48, 0, 0, 346, 347, 0, 0, 59,
	56, 0, 0, 68, 58, 0, 76, 77, 0, 0,
	0, 69, 70, 71, 72, 73, 0, 0, 0, 339,
	340, 0, 60, 40, 66, 41, 67, 0, 0, 63,
	37, 64, 344, 75, 74, 345, 38, 39, 0, 57,
	0, 48, 0, 0, 346, 347, 0, 0, 59, 56,
	0, 0, 68, 58, 0, 76, 77, 0, 0, 0,
	69, 70, 71, 72, 73, 0, 0, 0, 339, 340,
	0, 0, 0, 0, 0, 0, 0, 0, 63, 0,
	64, 0, 75, 74, 60, 40, 66, 41, 67, 0,
	0, 0, 37, 314, 46, 0, 0, 47, 38, 39,
	0, 57, 0, 48, 290, 0, 0, 0, 0, 0,
	59, 56, 0, 0, 68, 58, 0, 76, 77, 0,
	0, 0, 69, 70, 71, 72, 73, 0, 0, 0,
	228, 229, 0, 0, 0, 0, 0, 0, 0, 0,
	63, 0, 64, 0, 75, 74, 60, 40, 66, 41,
	67, 0, 0, 0, 37, 307, 46, 0, 0, 47,
	38, 39, 0, 57, 0, 48, 290, 0, 0, 0,
	0, 0, 59, 56, 0, 0, 68, 58, 0, 76,
	77, 0, 0, 0, 69, 70, 71, 72, 73, 0,
	0, 0, 228, 229, 0, 0, 0, 0, 0, 0,
	0, 0, 63, 0, 64, 0, 75, 74, 60, 40,
	66, 41, 67, 0, 0, 0, 37, 401, 46, 0,
	0, 47, 38, 39, 0, 57, 0, 48, 0, 0,
	0, 0, 0, 0, 59, 56, 0, 0, 68, 58,
	0, 76, 77, 0, 0, 0, 69, 70, 71, 72,
	73, 0, 0, 0, 228, 229, 0, 0, 0, 0,
	0, 0, 0, 0, 63, 0, 64, 0, 75, 74,
	60, 40, 66, 41, 67, 0, 0, 0, 37, 366,
	46, 0, 0, 47, 38, 39, 0, 57, 0, 48,
	0, 0, 0, 0, 0, 0, 59, 56, 0, 0,
	68, 58, 0, 76, 77, 0, 0, 0, 69, 70,
	71, 72, 73, 0, 0, 0, 228, 229, 0, 0,
	0, 0, 0, 0, 0, 0, 63, 0, 64, 0,
	75, 74, 60, 40, 66, 41, 67, 0, 0, 0,
	37, 335, 46, 0, 0, 47, 38, 39, 0, 57,
	0, 48, 0, 0, 0, 0, 0, 0, 59, 56,
	0, 0, 68, 58, 0, 76, 77, 0, 0, 0,
	69, 70, 71, 72, 73, 0, 0, 0, 228, 229,
	0, 60, 40, 66, 41, 67, 0, 0, 63, 37,
	64, 46, 75, 74, 47, 38, 39, 0, 57, 0,
	48, 0, 0, 0, 0, 0, 0, 59, 56, 0,
	0, 68, 58, 0, 76, 77, 0, 0, 0, 69,
	70, 71, 72, 73, 0, 0, 0, 228, 229, 0,
	0, 0, 0, 0, 0, 0, 0, 63, 0, 64,
	327, 75, 74, 60, 40, 66, 41, 67, 0, 0,
	0, 37, 324, 46, 0, 0, 47, 38, 39, 0,
	57, 0, 48, 0, 0, 0, 0, 0, 0, 59,
	56, 0, 0, 68, 58, 0, 76, 77, 0, 0,
	0, 69, 70, 71, 72, 73, 0, 0, 0, 228,
	229, 0, 60, 40, 66, 41, 67, 0, 0, 63,
	37, 64, 46, 75, 74, 47, 38, 39, 0, 57,
	0, 48, 0, 0, 0, 0, 0, 0, 59, 56,
	0, 0, 68, 58, 0, 76, 77, 0, 0, 0,
	69, 70, 71, 72, 73, 0, 0, 0, 228, 229,
	0, 0, 0, 0, 0, 0, 0, 0, 63, 0,
	64, 280, 75, 74, 60, 40, 66, 41, 67, 0,
	0, 0, 37, 275, 46, 0, 0, 47, 38, 39,
	0, 57, 0, 48, 0, 0, 0, 0, 0, 0,
	59, 56, 0, 0, 68, 58, 0, 76, 77, 0,
	0, 0, 69, 70, 71, 72, 73, 0, 0, 0,
	228, 229, 0, 0, 0, 0, 0, 0, 0, 0,
	63, 0, 64, 0, 75, 74, 60, 40, 66, 41,
	67, 0, 0, 0, 37, 267, 46, 0, 0, 47,
	38, 39, 0, 57, 0, 48, 0, 0, 0, 0,
	0, 0, 59, 56, 0, 0, 68, 58, 0, 76,
	77, 0, 0, 0, 69, 70, 71, 72, 73, 0,
	0, 0, 228, 229, 0, 0, 0, 0, 0, 0,
	0, 0, 63, 0, 64, 0, 75, 74, 60, 40,
	66, 41, 67, 0, 0, 0, 37, 265, 46, 0,
	0, 47, 38, 39, 0, 57, 0, 48, 0, 0,
	0, 0, 0, 0, 59, 56, 0, 0, 68, 58,
	0, 76, 77, 0, 0, 0, 69, 70, 71, 72,
	73, 0, 0, 0, 228, 229, 0, 60, 40, 66,
	41, 67, 0, 0, 63, 37, 64, 46, 75, 74,
	47, 38, 39, 0, 57, 0, 48, 0, 0, 0,
	0, 0, 0, 59, 56, 0, 0, 68, 58, 0,
	76, 77, 0, 0, 0, 69, 70, 71, 72, 73,
	0, 0, 0, 228, 229, 0, 60, 40, 66, 41,
	67, 249, 0, 63, 37, 64, 46, 75, 74, 47,
	38, 39, 0, 57, 0, 48, 0, 0, 0, 0,
	0, 0, 59, 56, 0, 0, 68, 58, 0, 76,
	77, 0, 0, 0, 69, 70, 71, 72, 73, 0,
	0, 0, 0, 248, 0, 60, 40, 66, 41, 67,
	0, 0, 63, 37, 64, 46, 75, 74, 47, 38,
	39, 0, 57, 0, 48, 0, 0, 0, 0, 0,
	0, 59, 56, 0, 0, 68, 58, 0, 76, 77,
	0, 0, 0, 69, 70, 71, 72, 73, 60, 120,
	66, 121, 255, 0, 0, 118, 0, 0, 0, 0,
	0, 63, 0, 64, 0, 75, 74, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 68, 0,
	0, 76, 77, 0, 0, 254, 69, 70, 71, 72,
	73, 60, 120, 66, 121, 67, 0, 0, 0, 0,
	0, 0, 0, 0, 63, 0, 119, 0, 75, 74,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 0, 0, 76, 77, 0, 0, 0, 69,
	70, 71, 72, 73, 0, 0, 0, 239, 60, 120,
	66, 121, 67, 0, 0, 118, 0, 63, 0, 64,
	0, 75, 74, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 68, 0,
	0, 76, 77, 0, 0, 0, 69, 70, 71, 72,
	73, 60, 244, 66, 121, 67, 0, 0, 0, 0,
	0, 0, 0, 0, 63, 0, 119, 0, 75, 74,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 0, 0, 76, 77, 0, 0, 0, 69,
	70, 71, 72, 73, 0, 0, 0, 239, 60, 120,
	66, 121, 107, 0, 0, 118, 0, 63, 0, 64,
	0, 75, 74, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 68, 0,
	0, 76, 77, 0, 0, 0, 69, 70, 71, 72,
	73, 60, 120, 66, 121, 67, 0, 0, 0, 0,
	0, 0, 0, 0, 110, 0, 119, 0, 75, 74,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 0, 0, 76, 77, 0, 0, 0, 69,
	70, 71, 72, 73, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 63, 0, 64,
	0, 75, 74,
}
var RubyPact = []int{

	-27, 568, -1000, -1000, -1000, 44, -1000, -1000, -1000, 28,
	-1000, -1000, -1000, 37, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 189, -21,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 277, 243, 243,
	443, 101, 133, 102, 71, 61, 2520, 2520, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 2786, 2520, 2786, 2786,
	-1000, -1000, -1000, 2786, -1000, -23, 222, -1000, 16, 15,
	15, 15, 15, 2786, 31, 267, -1000, -1000, -1000, -1000,
	-1000, 2520, 330, 2786, 2786, 2786, 2786, 2520, 2786, 2786,
	2786, 2786, 2520, 2786, 2786, 2786, 2786, 2520, 2786, 2520,
	2520, 325, 98, 99, 20, -1000, -1000, 2786, -1000, 105,
	2786, 2786, 2786, 322, 172, 248, -1000, 17, -26, -26,
	2743, 206, -1000, -1000, -1000, 2786, 2520, 2520, 2520, 2520,
	2520, 2520, 2520, 2520, 2520, 320, 41, 121, 2422, 60,
	248, 103, 248, 60, 40, 32, 122, -1000, 2696, 253,
	2786, 2471, -1000, -26, -1000, -1000, -1000, -1000, -1000, -1000,
	248, -1000, -1000, 250, -1000, -1000, 41, 2563, 248, 248,
	248, 248, 41, 248, 248, 248, 248, 41, 339, 248,
	248, 248, 41, 248, 41, 41, -1000, -1000, 315, 165,
	59, -1000, 3, 313, 312, -1000, 2373, 243, 2311, 288,
	122, -1000, 492, 236, -1000, 236, -1000, -1000, 2653, 2249,
	-1000, 309, -1000, 2187, 238, 41, 41, 41, 41, 41,
	41, 41, 41, 41, -1000, 874, -1000, -1000, -1000, -1000,
	28, 41, 197, 2786, -1000, 1, -1000, -1000, -1000, -1000,
	155, 147, -9, 293, 1225, -1000, 199, 41, -1000, -1000,
	-1000, -1000, 105, 16, 2520, 2786, 2786, 1841, 99, 59,
	182, 2520, -1000, -1000, 1779, -1000, -1000, -1000, 16, -1000,
	676, 2, -25, 248, -1000, -1000, 2138, 55, -1000, 2076,
	-1000, -1000, -1000, 351, 2520, -1000, 181, 2027, -1000, -1000,
	237, 248, 1668, 220, 2786, 2606, -24, -1000, -29, -1000,
	2520, 2786, -1000, -1000, 41, 137, 248, -1000, 177, -1000,
	-1000, -1000, -1000, 41, -1000, 124, 1965, -1000, 802, 248,
	-1000, 305, 2520, 298, -1000, -1000, 294, -1000, -1000, 2520,
	-1000, 41, 2422, -1000, 321, -1000, 2422, 290, -1000, -1000,
	-1000, 41, -1000, -1000, 2520, 2520, 223, 156, -1000, -1000,
	2786, 60, 122, -1000, -1000, 2606, -1000, 64, 28, 41,
	248, -1000, -1000, -1000, 1352, -1000, -1000, 286, -1000, 41,
	-1000, -1000, 41, 2422, 2422, -1000, 2422, 281, 48, 108,
	2520, 2520, 2520, 2520, 1903, 60, 2422, 231, -17, -1000,
	100, 109, 2422, -1000, -1000, -1000, -1000, 41, 41, 41,
	41, -1000, 2422, -9, 2520, 2786, -1000, -1000, 2422, 1544,
	1482, 1606, -9, 228, 219, -1000, 300, 2520, -1000, -1000,
	291, -1000, -1000, -1000, -9, -1000, -1000, 2520, -1000, 41,
	1717, -1000, -9, 41, 1717, 1717, 1717,
}
var RubyPgo = []int{

	0, 0, 447, 446, 444, 9, 443, 442, 439, 948,
	434, 4, 29, 432, 429, 426, 425, 265, 19, 707,
	1050, 424, 423, 422, 13, 421, 8, 605, 420, 419,
	10, 413, 409, 403, 402, 401, 400, 398, 6, 397,
	396, 393, 390, 389, 388, 385, 375, 374, 372, 371,
	1089, 370, 1, 17, 18, 3, 369, 5, 363, 11,
	353, 14, 347, 7, 344, 342, 16, 15, 340, 327,
	12,
}
var RubyR1 = []int{

	0, 56, 56, 56, 56, 56, 56, 56, 56, 56,
	56, 69, 69, 70, 70, 50, 50, 50, 50, 50,
	18, 18, 18, 18, 18, 18, 18, 18, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 24, 24, 24, 24, 24, 24, 24, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 14, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 53, 53, 53, 53, 63, 63, 61,
	61, 61, 61, 61, 66, 66, 66, 66, 66, 65,
	65, 65, 21, 21, 21, 21, 21, 21, 57, 57,
	26, 26, 26, 26, 67, 67, 67, 25, 25, 28,
	30, 30, 68, 68, 15, 15, 15, 15, 15, 15,
	15, 29, 29, 29, 29, 29, 29, 9, 9, 27,
	27, 19, 19, 39, 39, 40, 41, 42, 43, 44,
	45, 46, 47, 48, 49, 2, 6, 7, 7, 3,
	3, 58, 58, 58, 58, 64, 64, 64, 5, 5,
	5, 5, 54, 62, 62, 62, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 55, 55, 55, 55,
	51, 51, 51, 8, 16, 11, 11, 11, 60, 60,
	52, 52, 22, 23, 12, 35, 59, 59, 59, 59,
	59, 59, 36, 36, 36, 36, 36, 36, 36, 37,
	37, 37, 37, 37, 38, 38, 38, 38, 34, 33,
	10, 32, 32, 31, 31, 4,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 1, 0, 2, 0, 2, 2, 2, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 4, 5, 1, 4, 4, 2,
	3, 4, 4, 5, 3, 5, 2, 3, 3, 3,
	3, 4, 6, 3, 7, 1, 5, 1, 3, 0,
	1, 1, 4, 4, 1, 1, 4, 4, 5, 1,
	3, 3, 5, 6, 7, 8, 5, 6, 1, 3,
	1, 2, 3, 2, 0, 1, 3, 4, 6, 4,
	1, 3, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 2,
	2, 3, 3, 3, 3, 2, 2, 2, 2, 3,
	3, 3, 3, 3, 3, 1, 1, 3, 3, 5,
	5, 0, 1, 3, 7, 3, 7, 8, 3, 4,
	4, 3, 3, 0, 1, 3, 4, 5, 3, 3,
	3, 3, 3, 5, 6, 5, 4, 3, 3, 2,
	0, 2, 2, 3, 4, 2, 3, 5, 0, 2,
	1, 2, 2, 2, 5, 5, 0, 2, 2, 2,
	2, 2, 0, 1, 3, 3, 1, 3, 3, 5,
	6, 5, 6, 5, 4, 3, 3, 2, 3, 3,
	2, 5, 7, 4, 5, 3,
}
var RubyChk = []int{

	-1000, -56, 51, 52, 68, -1, 51, 52, 68, -17,
	-21, -25, -28, -15, -29, -13, -16, -24, -22, -35,
	-34, -33, -32, -23, -18, -7, -3, -30, -20, -8,
	-10, -40, -41, -42, -43, -4, -14, 13, 19, 20,
	6, 8, -27, -19, -9, -39, 15, 18, 24, -44,
	-45, -46, -47, -48, -49, -12, 32, 22, 36, 31,
	5, -2, -6, 61, 63, -68, 7, 9, 35, 43,
	44, 45, 46, 47, 66, 65, 38, 39, 52, 51,
	68, 15, 55, 40, 41, 4, 45, 46, 47, 57,
	58, 56, 18, 59, 33, 34, 48, 18, 40, 15,
	18, 55, 6, 4, -30, 8, -30, 9, -53, -66,
	61, 42, 49, 11, -65, -17, -5, -20, 12, 63,
	6, 8, -27, -19, -9, 9, 42, 49, 42, 49,
	42, 49, 42, 49, 42, 11, -1, -1, -50, -63,
	-17, -1, -17, -63, -63, -61, -17, -24, -70, 54,
	9, -51, -5, 63, -18, 6, 8, -18, -18, -18,
	-17, 6, 8, 66, 6, 8, -1, 6, -17, -17,
	-17, -17, -1, -17, -17, -17, -17, -1, -17, -17,
	-17, -17, -1, -17, -1, -1, 6, -57, 55, -67,
	9, -26, 6, 47, 58, -57, -50, 40, -50, -61,
	-17, 11, -17, -17, -12, -17, -12, 6, 11, -50,
	-54, 56, -54, -50, -61, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, 6, -50, 51, 52, 51, 52,
	-17, -1, -60, 11, 51, -70, 62, 11, 62, 51,
	-58, -64, -70, -17, 6, 8, -61, -1, 52, 10,
	6, 8, -66, -53, 42, 9, 53, -50, 6, 11,
	-67, 42, 6, 6, -50, 14, -30, 14, 10, 11,
	-70, 62, -70, -17, -5, 14, -50, -62, 6, -50,
	64, 10, 14, -55, 17, 16, -50, -50, 14, -11,
	25, -17, -59, -31, 37, -70, -70, 11, -70, 11,
	4, 53, 10, -5, -1, -61, -17, 14, -52, -11,
	-57, -26, 10, -1, 14, -52, -50, -5, -70, -17,
	-5, 58, 42, 58, 14, 56, 11, 64, 14, 17,
	16, -1, -50, 14, -55, 14, -50, 8, 14, 51,
	52, -1, -37, -36, 15, 18, 27, 28, 14, 16,
	37, -63, -17, -24, 64, -70, 64, -70, -17, -1,
	-17, 10, 14, -11, -50, 14, 14, 58, 6, -1,
	6, 6, -1, -50, -50, 14, -50, 4, -1, -1,
	15, 18, 15, 18, -50, -63, -50, -17, 6, 14,
	-52, 6, -50, 6, 51, 51, 52, -1, -1, -1,
	-1, 14, -50, -70, 4, 53, 14, 10, -50, -59,
	-59, -59, -70, -1, -17, 14, -38, 17, 16, 14,
	-38, 14, -69, 11, -70, 11, 14, 17, 16, -1,
	-59, 14, -70, -1, -59, -59, -59,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 48,
	49, 50, 51, 52, 53, 54, 55, 56, 57, 58,
	59, 60, 61, 62, 28, 29, 30, 31, 32, 33,
	34, 35, 36, 37, 38, 39, 40, 0, 0, 0,
	21, -2, 23, 24, 25, 0, 0, 0, 15, 41,
	42, 43, 44, 45, 46, 47, 0, 0, 0, 0,
	20, 26, 27, 89, 13, 0, 66, 190, 0, 0,
	0, 0, 0, 0, 0, 0, 155, 156, 5, 6,
	7, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 114, 114, 15, -2, 15, -2, 69, 76,
	89, 0, 0, 0, 85, 94, 95, 32, 15, -2,
	21, -2, 23, 24, 25, 89, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 15, 0, 198, 202,
	87, 0, 13, 203, 0, 0, 87, 91, 13, 0,
	89, 0, 230, 15, 145, 21, 22, 146, 147, 148,
	63, 139, 140, 0, 137, 138, 178, 70, 77, 79,
	80, 149, 150, 151, 152, 153, 154, 180, 0, 228,
	229, 235, 182, 78, 179, 181, 74, 15, 0, 108,
	114, 115, 110, 0, 0, 15, 0, 0, 0, 0,
	90, 13, 87, 124, 125, 131, 132, 143, 13, 0,
	15, 173, 15, 0, 0, 126, 133, 127, 134, 128,
	135, 129, 136, 130, 144, 0, 15, 15, 16, 17,
	18, 19, 0, 0, 206, 0, 157, 13, 158, 14,
	13, 13, 162, 0, 21, -2, 0, 191, 192, 193,
	141, 142, 71, 72, 0, -2, 0, 0, 114, 0,
	0, 0, 111, 113, 0, 117, 15, 119, 64, 13,
	0, 81, 0, 100, 101, 168, 0, 0, 174, 0,
	171, 68, 176, 0, 0, 15, 0, 0, 194, 199,
	15, 88, 0, 0, 0, 0, 0, 13, 0, 13,
	0, 0, 67, 73, 75, 0, 204, 102, 0, 200,
	15, 116, 109, 112, 106, 0, 0, 65, 0, 96,
	97, 0, 0, 0, 169, 172, 0, 170, 177, 0,
	15, 15, 189, 183, 0, 185, 195, 15, 205, 207,
	208, 209, 210, 211, 0, 0, 213, 216, 231, 15,
	0, 15, 92, 93, 159, 0, 160, 0, 48, 163,
	165, 83, 103, 201, 0, 107, 118, 0, 98, 82,
	86, 175, 15, 187, 188, 184, 196, 0, 15, 0,
	0, 0, 0, 0, 0, 15, 13, 0, 0, 104,
	0, 0, 186, 15, 206, -2, -2, 214, 215, 217,
	218, 232, 13, 233, 0, 0, 105, 84, 197, 0,
	0, 0, 234, 11, 13, 219, 0, 0, 206, 221,
	0, 223, 164, 12, 166, 13, 220, 0, 206, 206,
	212, 222, 167, 206, 212, 212, 212,
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
		//line parser.y:199
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:201
		{
		}
	case 3:
		//line parser.y:203
		{
		}
	case 4:
		//line parser.y:205
		{
		}
	case 5:
		//line parser.y:207
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:209
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:211
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:217
		{
		}
	case 11:
		//line parser.y:219
		{
		}
	case 12:
		//line parser.y:220
		{
		}
	case 13:
		//line parser.y:222
		{
		}
	case 14:
		//line parser.y:223
		{
		}
	case 15:
		//line parser.y:226
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 16:
		//line parser.y:228
		{
		}
	case 17:
		//line parser.y:230
		{
		}
	case 18:
		//line parser.y:232
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 19:
		//line parser.y:234
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
		//line parser.y:246
		{
			RubyVAL.genericValue = ast.StarSplat{Value: RubyS[Rubypt-0].genericValue}
		}
	case 64:
		//line parser.y:249
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 65:
		//line parser.y:256
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-2].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 66:
		//line parser.y:263
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 67:
		//line parser.y:267
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 68:
		//line parser.y:274
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 69:
		//line parser.y:281
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 70:
		//line parser.y:288
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 71:
		//line parser.y:295
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 72:
		//line parser.y:303
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 73:
		//line parser.y:311
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 74:
		//line parser.y:319
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 75:
		//line parser.y:327
		{
			methodName := RubyS[Rubypt-2].genericValue.(ast.BareReference).Name + "="
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: methodName},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 76:
		//line parser.y:338
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 77:
		//line parser.y:345
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 78:
		//line parser.y:353
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 79:
		//line parser.y:361
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 80:
		//line parser.y:369
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 81:
		//line parser.y:379
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 82:
		//line parser.y:389
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 83:
		//line parser.y:398
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 84:
		//line parser.y:400
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 85:
		//line parser.y:402
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 86:
		//line parser.y:404
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 87:
		//line parser.y:407
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 88:
		//line parser.y:409
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 89:
		//line parser.y:411
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 90:
		//line parser.y:413
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 91:
		//line parser.y:415
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 92:
		//line parser.y:417
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 93:
		//line parser.y:419
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 94:
		//line parser.y:423
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 95:
		//line parser.y:425
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 96:
		//line parser.y:427
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 97:
		//line parser.y:429
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-1].genericValue)
		}
	case 98:
		//line parser.y:431
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 99:
		//line parser.y:434
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 100:
		//line parser.y:436
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 101:
		//line parser.y:438
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 102:
		//line parser.y:444
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 103:
		//line parser.y:452
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 104:
		//line parser.y:461
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 105:
		//line parser.y:470
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 106:
		//line parser.y:480
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 107:
		//line parser.y:488
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 108:
		//line parser.y:499
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 109:
		//line parser.y:501
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 110:
		//line parser.y:504
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 111:
		//line parser.y:506
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 112:
		//line parser.y:508
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 113:
		//line parser.y:510
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 114:
		//line parser.y:512
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 115:
		//line parser.y:514
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 116:
		//line parser.y:518
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 117:
		//line parser.y:523
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 118:
		//line parser.y:530
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 119:
		//line parser.y:540
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 120:
		//line parser.y:549
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 121:
		//line parser.y:555
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 122:
		//line parser.y:563
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 123:
		//line parser.y:567
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 124:
		//line parser.y:572
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 125:
		//line parser.y:579
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 126:
		//line parser.y:583
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 127:
		//line parser.y:590
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 128:
		//line parser.y:597
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 129:
		//line parser.y:604
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 130:
		//line parser.y:611
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 131:
		//line parser.y:619
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 132:
		//line parser.y:626
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 133:
		//line parser.y:630
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 134:
		//line parser.y:637
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 135:
		//line parser.y:644
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 136:
		//line parser.y:651
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 137:
		//line parser.y:659
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 138:
		//line parser.y:661
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 139:
		//line parser.y:664
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 140:
		//line parser.y:666
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 141:
		//line parser.y:669
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 142:
		//line parser.y:671
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 143:
		//line parser.y:674
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 144:
		//line parser.y:676
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 145:
		//line parser.y:678
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 146:
		//line parser.y:679
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 147:
		//line parser.y:680
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 148:
		//line parser.y:681
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 149:
		//line parser.y:684
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 150:
		//line parser.y:693
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 151:
		//line parser.y:702
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 152:
		//line parser.y:711
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 153:
		//line parser.y:720
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 154:
		//line parser.y:729
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 155:
		//line parser.y:737
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 156:
		//line parser.y:738
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 157:
		//line parser.y:740
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 158:
		//line parser.y:742
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 159:
		//line parser.y:745
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 160:
		//line parser.y:753
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 161:
		//line parser.y:761
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 162:
		//line parser.y:763
		{
		}
	case 163:
		//line parser.y:765
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 164:
		//line parser.y:772
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 165:
		//line parser.y:780
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 166:
		//line parser.y:787
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 167:
		//line parser.y:794
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 168:
		//line parser.y:802
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 169:
		//line parser.y:804
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 170:
		//line parser.y:811
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 171:
		//line parser.y:818
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 172:
		//line parser.y:821
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 173:
		//line parser.y:823
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 174:
		//line parser.y:825
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 175:
		//line parser.y:827
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 176:
		//line parser.y:830
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 177:
		//line parser.y:837
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 178:
		//line parser.y:845
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 179:
		//line parser.y:852
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 180:
		//line parser.y:859
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 181:
		//line parser.y:866
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 182:
		//line parser.y:873
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 183:
		//line parser.y:880
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 184:
		//line parser.y:887
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 185:
		//line parser.y:895
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 186:
		//line parser.y:904
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 187:
		//line parser.y:911
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 188:
		//line parser.y:918
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 189:
		//line parser.y:925
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 190:
		//line parser.y:932
		{
		}
	case 191:
		//line parser.y:933
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 192:
		//line parser.y:934
		{
		}
	case 193:
		//line parser.y:937
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 194:
		//line parser.y:940
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 195:
		//line parser.y:948
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 196:
		//line parser.y:950
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 197:
		//line parser.y:959
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
	case 198:
		//line parser.y:974
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 199:
		//line parser.y:976
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 200:
		//line parser.y:979
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 201:
		//line parser.y:981
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 202:
		//line parser.y:984
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 203:
		//line parser.y:993
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 204:
		//line parser.y:1002
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 205:
		//line parser.y:1011
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 206:
		//line parser.y:1014
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 207:
		//line parser.y:1016
		{
		}
	case 208:
		//line parser.y:1018
		{
		}
	case 209:
		//line parser.y:1020
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 210:
		//line parser.y:1022
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 211:
		//line parser.y:1024
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 212:
		//line parser.y:1026
		{
		}
	case 213:
		//line parser.y:1028
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 214:
		//line parser.y:1030
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 215:
		//line parser.y:1032
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 216:
		//line parser.y:1034
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 217:
		//line parser.y:1036
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 218:
		//line parser.y:1038
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 219:
		//line parser.y:1041
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 220:
		//line parser.y:1048
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 221:
		//line parser.y:1056
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 222:
		//line parser.y:1063
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 223:
		//line parser.y:1071
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 224:
		//line parser.y:1079
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 225:
		//line parser.y:1086
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 226:
		//line parser.y:1093
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 227:
		//line parser.y:1100
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 228:
		//line parser.y:1108
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 229:
		//line parser.y:1111
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 230:
		//line parser.y:1113
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 231:
		//line parser.y:1116
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 232:
		//line parser.y:1118
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 233:
		//line parser.y:1121
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 234:
		//line parser.y:1123
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 235:
		//line parser.y:1125
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
