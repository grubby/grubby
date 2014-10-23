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
const WHITESPACE = 57390
const NEWLINE = 57391
const SEMICOLON = 57392
const COLON = 57393
const DOUBLECOLON = 57394
const DOT = 57395
const PIPE = 57396
const SLASH = 57397
const AMPERSAND = 57398
const QUESTIONMARK = 57399
const CARET = 57400
const LBRACKET = 57401
const RBRACKET = 57402
const LBRACE = 57403
const RBRACE = 57404
const DOLLARSIGN = 57405
const ATSIGN = 57406
const FILE_CONST_REF = 57407
const EOF = 57408

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

//line parser.y:1077

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 28,
	52, 123,
	-2, 20,
	-1, 90,
	10, 82,
	11, 82,
	-2, 184,
	-1, 100,
	52, 123,
	-2, 20,
	-1, 106,
	13, 13,
	15, 13,
	18, 13,
	19, 13,
	20, 13,
	22, 13,
	24, 13,
	32, 13,
	36, 13,
	43, 13,
	44, 13,
	45, 13,
	46, 13,
	50, 13,
	-2, 11,
	-1, 121,
	52, 123,
	-2, 121,
	-1, 225,
	52, 124,
	-2, 122,
	-1, 233,
	10, 82,
	11, 82,
	-2, 184,
	-1, 384,
	27, 200,
	28, 200,
	-2, 13,
	-1, 385,
	27, 200,
	28, 200,
	-2, 13,
}

const RubyNprod = 230
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 2413

var RubyAct = []int{

	9, 294, 408, 141, 349, 261, 190, 36, 20, 194,
	98, 177, 192, 92, 91, 105, 142, 2, 3, 74,
	329, 176, 195, 220, 218, 193, 212, 146, 97, 113,
	114, 270, 220, 292, 4, 136, 330, 137, 117, 289,
	220, 394, 112, 288, 272, 89, 376, 253, 111, 377,
	120, 122, 74, 72, 71, 406, 220, 200, 220, 133,
	110, 135, 132, 196, 150, 109, 143, 124, 380, 78,
	73, 328, 197, 219, 144, 217, 291, 153, 154, 155,
	156, 149, 158, 159, 160, 161, 383, 163, 164, 165,
	220, 168, 195, 138, 170, 171, 74, 370, 123, 144,
	97, 78, 144, 220, 373, 76, 77, 167, 168, 212,
	79, 80, 81, 145, 185, 186, 144, 175, 75, 84,
	82, 83, 86, 169, 179, 238, 198, 26, 99, 68,
	100, 69, 209, 196, 105, 107, 74, 76, 77, 220,
	220, 374, 197, 210, 266, 5, 223, 393, 168, 322,
	75, 323, 268, 97, 86, 268, 144, 70, 74, 415,
	63, 64, 78, 74, 226, 276, 230, 231, 108, 417,
	213, 220, 324, 78, 240, 209, 251, 364, 287, 209,
	365, 65, 173, 106, 241, 62, 61, 335, 236, 125,
	126, 127, 128, 129, 130, 395, 74, 107, 76, 77,
	209, 362, 209, 134, 363, 209, 252, 147, 258, 76,
	77, 75, 215, 269, 216, 86, 267, 139, 151, 140,
	234, 310, 75, 121, 157, 97, 86, 297, 251, 162,
	384, 385, 225, 166, 168, 282, 382, 423, 285, 420,
	419, 119, 144, 118, 209, 372, 283, 209, 286, 214,
	281, 340, 180, 181, 182, 183, 339, 295, 299, 187,
	188, 296, 277, 337, 359, 209, 209, 202, 74, 255,
	307, 314, 195, 133, 326, 193, 325, 254, 332, 334,
	279, 218, 327, 237, 250, 326, 418, 239, 420, 419,
	245, 115, 227, 327, 116, 344, 248, 218, 76, 77,
	344, 209, 205, 206, 249, 357, 209, 303, 302, 203,
	209, 75, 256, 196, 189, 86, 113, 114, 301, 191,
	303, 302, 197, 235, 236, 133, 184, 172, 367, 112,
	152, 67, 273, 96, 222, 274, 275, 244, 211, 221,
	26, 99, 68, 100, 90, 1, 95, 105, 347, 148,
	284, 133, 375, 379, 378, 57, 209, 209, 56, 209,
	55, 54, 53, 52, 17, 16, 15, 209, 14, 209,
	70, 44, 316, 63, 64, 317, 280, 94, 22, 23,
	344, 24, 209, 25, 271, 346, 12, 11, 348, 21,
	331, 10, 209, 19, 93, 405, 106, 298, 62, 61,
	209, 314, 314, 314, 412, 13, 304, 18, 341, 351,
	37, 40, 39, 352, 315, 400, 401, 402, 34, 33,
	35, 333, 32, 314, 0, 0, 0, 0, 314, 314,
	314, 0, 338, 0, 0, 78, 0, 0, 345, 0,
	0, 0, 422, 345, 0, 0, 354, 371, 0, 85,
	0, 0, 427, 428, 0, 0, 0, 0, 429, 0,
	0, 0, 360, 361, 87, 88, 0, 0, 0, 0,
	0, 76, 77, 369, 0, 0, 79, 80, 81, 0,
	0, 0, 392, 0, 75, 84, 82, 83, 86, 0,
	0, 0, 0, 398, 0, 0, 0, 0, 29, 0,
	0, 0, 0, 0, 0, 403, 386, 387, 388, 389,
	0, 0, 0, 0, 0, 0, 0, 0, 414, 416,
	396, 397, 0, 345, 0, 0, 101, 0, 0, 424,
	0, 425, 0, 0, 0, 0, 0, 404, 0, 0,
	0, 38, 0, 0, 315, 315, 315, 0, 0, 0,
	0, 0, 0, 421, 0, 0, 0, 101, 0, 101,
	0, 0, 0, 426, 101, 0, 315, 0, 0, 104,
	0, 315, 315, 315, 0, 101, 101, 101, 101, 0,
	101, 101, 101, 101, 0, 101, 101, 101, 0, 101,
	0, 0, 101, 101, 0, 0, 0, 0, 101, 0,
	104, 0, 104, 0, 0, 0, 101, 104, 0, 0,
	0, 0, 101, 101, 0, 0, 0, 0, 104, 104,
	104, 104, 0, 104, 104, 104, 104, 0, 104, 104,
	104, 0, 104, 0, 0, 104, 104, 0, 0, 0,
	0, 104, 0, 0, 101, 0, 101, 0, 0, 104,
	0, 101, 0, 0, 0, 104, 104, 0, 26, 99,
	68, 100, 69, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 101, 0, 0, 0, 178, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 104, 70, 104,
	0, 63, 64, 0, 104, 0, 0, 26, 99, 68,
	100, 90, 220, 0, 105, 0, 0, 0, 0, 336,
	0, 101, 65, 0, 66, 104, 62, 61, 0, 0,
	0, 0, 0, 101, 0, 0, 0, 70, 131, 0,
	63, 64, 101, 101, 0, 0, 101, 0, 0, 0,
	0, 0, 0, 278, 0, 0, 0, 0, 0, 0,
	30, 93, 78, 106, 104, 62, 61, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 104, 0, 0, 0,
	0, 101, 101, 0, 0, 104, 104, 101, 102, 104,
	0, 0, 174, 101, 0, 0, 0, 0, 76, 77,
	0, 0, 0, 79, 80, 81, 0, 199, 0, 201,
	0, 75, 84, 82, 83, 86, 204, 0, 0, 102,
	0, 102, 0, 0, 104, 104, 102, 0, 0, 0,
	104, 0, 0, 101, 0, 0, 104, 102, 102, 102,
	102, 0, 102, 102, 102, 102, 0, 102, 102, 102,
	0, 102, 0, 31, 102, 102, 0, 0, 0, 101,
	102, 0, 243, 0, 246, 0, 0, 0, 102, 0,
	0, 0, 0, 0, 102, 102, 104, 0, 0, 0,
	0, 103, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 264, 265, 0, 0, 0, 0, 0, 0,
	0, 0, 104, 101, 0, 0, 102, 0, 102, 0,
	0, 0, 103, 102, 103, 0, 0, 0, 0, 103,
	0, 0, 0, 0, 26, 99, 68, 100, 69, 0,
	103, 103, 103, 103, 102, 103, 103, 103, 103, 0,
	103, 103, 103, 0, 103, 300, 104, 103, 103, 0,
	305, 0, 0, 103, 70, 309, 0, 63, 64, 0,
	0, 103, 0, 0, 0, 0, 0, 103, 103, 0,
	0, 0, 0, 102, 0, 0, 0, 0, 65, 0,
	66, 0, 62, 61, 0, 102, 0, 0, 0, 0,
	355, 356, 0, 0, 102, 102, 0, 358, 102, 103,
	0, 103, 0, 0, 0, 0, 103, 0, 0, 0,
	366, 0, 368, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 103, 0, 0,
	0, 0, 0, 102, 102, 0, 0, 0, 0, 102,
	0, 381, 0, 0, 0, 102, 0, 204, 0, 0,
	0, 0, 0, 0, 391, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 103, 0, 0, 399,
	0, 264, 265, 0, 0, 0, 0, 0, 103, 0,
	0, 0, 0, 0, 0, 102, 0, 103, 103, 0,
	0, 103, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 26, 27, 68, 28,
	69, 102, 0, 0, 41, 411, 318, 410, 409, 319,
	42, 43, 0, 59, 0, 51, 103, 103, 320, 321,
	0, 0, 103, 58, 0, 0, 70, 60, 103, 63,
	64, 0, 0, 0, 45, 46, 47, 48, 0, 0,
	312, 313, 0, 0, 0, 102, 0, 0, 0, 0,
	65, 0, 66, 0, 62, 61, 0, 0, 0, 0,
	0, 0, 26, 27, 68, 28, 69, 0, 103, 0,
	41, 407, 318, 410, 409, 319, 42, 43, 0, 59,
	0, 51, 0, 0, 320, 321, 0, 0, 0, 58,
	0, 0, 70, 60, 103, 63, 64, 0, 0, 0,
	45, 46, 47, 48, 0, 0, 312, 313, 0, 0,
	0, 0, 0, 0, 0, 0, 65, 0, 66, 0,
	62, 61, 26, 27, 68, 28, 69, 0, 0, 0,
	41, 413, 318, 0, 0, 319, 42, 43, 103, 59,
	0, 51, 0, 0, 320, 321, 0, 0, 0, 58,
	0, 0, 70, 60, 0, 63, 64, 0, 0, 0,
	45, 46, 47, 48, 0, 0, 312, 313, 0, 0,
	0, 0, 0, 0, 0, 0, 65, 0, 66, 0,
	62, 61, 26, 27, 68, 28, 69, 0, 0, 0,
	41, 311, 318, 0, 0, 319, 42, 43, 0, 59,
	0, 51, 0, 0, 320, 321, 0, 0, 0, 58,
	0, 0, 70, 60, 0, 63, 64, 0, 0, 0,
	45, 46, 47, 48, 0, 0, 312, 313, 0, 0,
	0, 0, 0, 0, 0, 0, 65, 0, 66, 0,
	62, 61, 26, 27, 68, 28, 69, 0, 0, 0,
	41, 306, 49, 263, 262, 50, 42, 43, 0, 59,
	0, 51, 0, 0, 0, 0, 0, 0, 0, 58,
	0, 0, 70, 60, 0, 63, 64, 0, 0, 0,
	45, 46, 47, 48, 0, 0, 207, 208, 0, 0,
	0, 0, 0, 0, 0, 0, 65, 0, 66, 0,
	62, 61, 26, 27, 68, 28, 69, 0, 0, 0,
	41, 260, 49, 263, 262, 50, 42, 43, 0, 59,
	0, 51, 0, 0, 0, 0, 0, 0, 0, 58,
	0, 0, 70, 60, 0, 63, 64, 0, 0, 0,
	45, 46, 47, 48, 0, 0, 207, 208, 0, 26,
	27, 68, 28, 69, 0, 0, 65, 41, 66, 49,
	62, 61, 50, 42, 43, 0, 59, 0, 51, 0,
	0, 0, 0, 0, 0, 0, 58, 0, 0, 70,
	60, 0, 63, 64, 0, 0, 0, 45, 46, 47,
	48, 0, 0, 6, 7, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 0, 66, 0, 62, 61, 0,
	8, 26, 27, 68, 28, 69, 0, 0, 0, 41,
	0, 318, 0, 0, 319, 42, 43, 0, 59, 0,
	51, 0, 0, 320, 321, 0, 0, 0, 58, 0,
	0, 70, 60, 0, 63, 64, 0, 0, 0, 45,
	46, 47, 48, 0, 0, 312, 313, 0, 26, 27,
	68, 28, 69, 0, 0, 65, 41, 66, 49, 62,
	61, 50, 42, 43, 0, 59, 0, 51, 268, 0,
	0, 0, 0, 0, 350, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 342, 343, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 66, 0, 62, 61, 26, 27,
	68, 28, 69, 0, 0, 0, 41, 390, 49, 0,
	0, 50, 42, 43, 0, 59, 0, 51, 0, 0,
	0, 0, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 207, 208, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 66, 0, 62, 61, 26, 27,
	68, 28, 69, 0, 0, 0, 41, 353, 49, 0,
	0, 50, 42, 43, 0, 59, 0, 51, 0, 0,
	0, 0, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 207, 208, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 66, 0, 62, 61, 26, 27,
	68, 28, 69, 0, 0, 0, 41, 308, 49, 0,
	0, 50, 42, 43, 0, 59, 0, 51, 0, 0,
	0, 0, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 207, 208, 0, 26, 27, 68, 28, 69,
	0, 0, 65, 41, 66, 49, 62, 61, 50, 42,
	43, 0, 59, 0, 51, 0, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 70, 60, 0, 63, 64,
	0, 0, 0, 45, 46, 47, 48, 0, 0, 207,
	208, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	0, 66, 293, 62, 61, 26, 27, 68, 28, 69,
	0, 0, 0, 41, 290, 49, 0, 0, 50, 42,
	43, 0, 59, 0, 51, 0, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 70, 60, 0, 63, 64,
	0, 0, 0, 45, 46, 47, 48, 0, 0, 207,
	208, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	0, 66, 0, 62, 61, 26, 27, 68, 28, 69,
	0, 0, 0, 41, 259, 49, 0, 0, 50, 42,
	43, 0, 59, 0, 51, 0, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 70, 60, 0, 63, 64,
	0, 0, 0, 45, 46, 47, 48, 0, 0, 207,
	208, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	0, 66, 0, 62, 61, 26, 27, 68, 28, 69,
	0, 0, 0, 41, 257, 49, 0, 0, 50, 42,
	43, 0, 59, 0, 51, 0, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 70, 60, 0, 63, 64,
	0, 0, 0, 45, 46, 47, 48, 0, 0, 207,
	208, 0, 26, 27, 68, 28, 69, 0, 0, 65,
	41, 66, 49, 62, 61, 50, 42, 43, 0, 59,
	0, 51, 0, 0, 0, 0, 0, 0, 0, 58,
	0, 0, 70, 60, 0, 63, 64, 0, 0, 0,
	45, 46, 47, 48, 0, 0, 207, 208, 0, 0,
	0, 0, 0, 0, 0, 0, 65, 0, 66, 247,
	62, 61, 26, 27, 68, 28, 69, 0, 0, 0,
	41, 242, 49, 0, 0, 50, 42, 43, 0, 59,
	0, 51, 0, 0, 0, 0, 0, 0, 0, 58,
	0, 0, 70, 60, 0, 63, 64, 0, 0, 0,
	45, 46, 47, 48, 0, 0, 207, 208, 0, 26,
	27, 68, 28, 69, 0, 0, 65, 41, 66, 49,
	62, 61, 50, 42, 43, 0, 59, 0, 51, 0,
	0, 0, 0, 0, 0, 0, 58, 0, 0, 70,
	60, 0, 63, 64, 0, 0, 0, 45, 46, 47,
	48, 0, 0, 207, 208, 0, 26, 27, 68, 28,
	69, 229, 0, 65, 41, 66, 49, 62, 61, 50,
	42, 43, 0, 59, 0, 51, 0, 0, 0, 0,
	0, 0, 0, 58, 0, 0, 70, 60, 0, 63,
	64, 0, 0, 0, 45, 46, 47, 48, 0, 0,
	0, 228, 0, 26, 27, 68, 28, 69, 0, 0,
	65, 41, 66, 49, 62, 61, 50, 42, 43, 0,
	59, 0, 51, 0, 0, 0, 0, 0, 0, 0,
	58, 0, 0, 70, 60, 0, 63, 64, 0, 0,
	0, 45, 46, 47, 48, 26, 99, 68, 100, 233,
	0, 0, 105, 0, 0, 0, 0, 65, 0, 66,
	0, 62, 61, 26, 99, 68, 100, 69, 0, 0,
	0, 0, 0, 0, 0, 70, 0, 0, 63, 64,
	0, 0, 232, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 70, 0, 0, 63, 64, 0, 65,
	0, 106, 0, 62, 61, 0, 0, 220, 26, 99,
	68, 100, 69, 0, 0, 105, 0, 65, 0, 66,
	0, 62, 61, 0, 26, 224, 68, 100, 69, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 70, 0,
	0, 63, 64, 26, 99, 68, 100, 90, 0, 0,
	105, 0, 0, 0, 70, 0, 0, 63, 64, 0,
	0, 0, 65, 0, 106, 0, 62, 61, 220, 0,
	0, 0, 0, 70, 0, 0, 63, 64, 65, 0,
	66, 0, 62, 61, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 93, 0, 106,
	0, 62, 61,
}
var RubyPact = []int{

	-32, 1444, -1000, -1000, -1000, 4, -1000, -1000, -1000, 431,
	-1000, -1000, -1000, 27, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 335, 126, 23,
	18, 6, -1000, -1000, -1000, -1000, -1000, -1000, 276, -15,
	-1000, 237, 215, 215, 56, 2208, 2208, 2208, 2208, 2208,
	2208, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 909, 2208,
	909, 29, 211, -1000, -1000, 909, -1000, -25, 198, -1000,
	3, -1000, -1000, -1000, 2208, 324, 909, 909, 909, 909,
	2208, 909, 909, 909, 909, 2208, 909, 909, 909, 2208,
	909, -1000, 112, 909, 909, 321, 171, 97, -1000, 2348,
	188, -1000, -1000, -1000, -11, -33, -33, 909, 2208, 2208,
	2208, 2208, 320, 909, 909, 2208, 2208, 308, 266, 16,
	17, -1000, -1000, 2208, 303, 81, 81, 81, 81, 81,
	253, 2114, 98, 97, 121, 97, -1000, -1000, 206, -1000,
	-1000, 15, 13, 748, -1000, 2329, 224, 909, 2161, -1000,
	-33, 81, 2250, 97, 97, 97, 97, 81, 97, 97,
	97, 97, 81, 169, 97, 97, 81, 313, 748, -1000,
	65, 97, -1000, 2313, 2067, -1000, 284, -1000, 2007, 286,
	81, 81, 81, 81, -1000, 97, 97, 81, 81, -1000,
	-1000, 278, 165, 86, -1000, 5, 271, 263, -1000, 1960,
	215, 1900, 81, -1000, 1397, -1000, -1000, -1000, -1000, 431,
	81, 130, 909, -1000, 7, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 154, 258, 692, -1000, 270, 81, -1000, -1000,
	112, -1000, 2208, 909, 909, 3, -1000, 122, 1, -17,
	97, -1000, -1000, 1840, 22, -1000, 1780, -1000, -1000, -9,
	16, 86, 217, 2208, -1000, -1000, -9, -1000, -1000, -1000,
	-1000, 304, 2208, -1000, 1337, 1733, -1000, -1000, 213, 97,
	1277, 135, 909, 2268, 9, -26, -1000, 2208, 909, -1000,
	81, 177, 97, -1000, 653, 97, -1000, 257, 2208, 250,
	-1000, -1000, 245, -1000, 1553, -1000, -1000, -1000, 81, 1553,
	1673, -1000, 2208, -1000, 81, 2114, -1000, 291, -1000, 2114,
	260, -1000, -1000, -1000, 431, 81, -1000, -1000, 2208, 2208,
	186, 162, -1000, -1000, 909, 98, 748, -1000, -1000, 2208,
	-1000, 91, 431, 81, 97, -1000, 239, -1000, 81, -1000,
	-1000, 90, -1000, -1000, 431, 81, -1000, 127, 31, -1000,
	909, -9, 54, -1000, 81, 2114, 2114, -1000, 2114, 230,
	37, 181, 2208, 2208, 2208, 2208, 1613, 98, 2114, 143,
	-10, -9, 185, -1000, -1000, -1000, 2208, 2208, 98, 1553,
	-1000, 2114, -1000, -1000, -1000, -1000, 81, 81, 81, 81,
	-1000, 2114, -9, 2208, 909, -1000, 81, 81, 41, 2114,
	1157, 1091, 1217, -9, 148, 158, -1000, -1000, 272, 2208,
	-1000, -1000, 223, -1000, -9, -1000, -9, -1000, -1000, 2208,
	-1000, 81, 1506, -1000, -9, -9, 81, 1506, 1506, 1506,
}
var RubyPgo = []int{

	0, 143, 422, 420, 10, 419, 418, 412, 843, 411,
	4, 410, 407, 405, 393, 0, 750, 541, 391, 389,
	388, 8, 387, 9, 498, 386, 7, 385, 384, 383,
	381, 379, 378, 375, 372, 2, 371, 368, 366, 365,
	364, 363, 362, 361, 360, 358, 355, 676, 349, 348,
	14, 11, 5, 345, 6, 339, 31, 338, 16, 1,
	337, 3, 334, 333, 13, 12, 331, 113,
}
var RubyR1 = []int{

	0, 53, 53, 53, 53, 53, 53, 53, 53, 53,
	53, 67, 67, 47, 47, 47, 47, 47, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 21, 21, 21, 21, 21, 21, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 50, 50, 50, 50,
	61, 61, 58, 58, 58, 58, 58, 64, 64, 64,
	64, 64, 63, 63, 63, 18, 18, 18, 27, 27,
	27, 27, 59, 59, 59, 59, 59, 59, 59, 54,
	54, 23, 23, 23, 23, 65, 65, 65, 22, 22,
	25, 26, 26, 66, 66, 13, 13, 13, 13, 13,
	13, 8, 8, 24, 24, 16, 16, 36, 36, 37,
	38, 39, 40, 41, 42, 43, 44, 45, 46, 2,
	5, 6, 6, 3, 3, 55, 55, 55, 55, 62,
	62, 62, 4, 4, 4, 4, 51, 60, 60, 60,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	52, 52, 52, 52, 48, 48, 48, 7, 14, 10,
	10, 10, 57, 57, 49, 49, 19, 20, 11, 32,
	56, 56, 56, 56, 56, 56, 56, 33, 33, 33,
	33, 33, 33, 33, 34, 34, 34, 34, 34, 35,
	35, 35, 35, 31, 30, 9, 29, 29, 28, 28,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 4, 5, 1, 4,
	4, 2, 3, 3, 4, 4, 3, 5, 2, 3,
	3, 3, 3, 3, 4, 6, 3, 7, 1, 5,
	1, 3, 0, 1, 1, 4, 4, 1, 1, 4,
	4, 5, 1, 3, 3, 7, 9, 7, 0, 1,
	3, 3, 0, 2, 2, 2, 2, 2, 3, 1,
	3, 1, 2, 3, 2, 0, 1, 3, 4, 6,
	4, 1, 3, 1, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 2, 2, 3, 3, 3, 3, 2,
	2, 2, 2, 3, 3, 3, 3, 3, 3, 1,
	1, 3, 3, 5, 5, 0, 4, 7, 8, 3,
	7, 8, 3, 4, 4, 3, 3, 0, 1, 3,
	4, 5, 3, 3, 3, 3, 3, 5, 6, 5,
	4, 3, 3, 2, 0, 2, 2, 3, 4, 2,
	3, 5, 0, 2, 1, 2, 2, 2, 5, 5,
	0, 2, 2, 2, 2, 2, 2, 0, 1, 3,
	3, 1, 3, 3, 5, 6, 5, 6, 5, 4,
	3, 3, 2, 3, 3, 2, 5, 7, 4, 5,
}
var RubyChk = []int{

	-1000, -53, 49, 50, 66, -1, 49, 50, 66, -15,
	-18, -22, -25, -13, -37, -38, -39, -40, -12, -14,
	-21, -19, -32, -31, -30, -29, 5, 6, 8, -24,
	-16, -8, -2, -5, -6, -3, -26, -11, -17, -7,
	-9, 13, 19, 20, -36, 43, 44, 45, 46, 15,
	18, 24, -41, -42, -43, -44, -45, -46, 32, 22,
	36, 64, 63, 38, 39, 59, 61, -66, 7, 9,
	35, 50, 49, 66, 15, 53, 40, 41, 4, 45,
	46, 47, 55, 56, 54, 18, 57, 33, 34, 18,
	9, -50, -64, 59, 42, 11, -63, -15, -4, 6,
	8, -24, -16, -8, -17, 12, 61, 9, 42, 42,
	42, 42, 53, 40, 41, 15, 18, 53, 6, 4,
	-26, 8, -26, 42, 11, -1, -1, -1, -1, -1,
	-1, -47, -61, -15, -1, -15, 6, 8, 64, 6,
	8, -61, -58, -15, -21, -67, 52, 9, -48, -4,
	61, -1, 6, -15, -15, -15, -15, -1, -15, -15,
	-15, -15, -1, -15, -15, -15, -1, -58, -15, 11,
	-15, -15, 6, 11, -47, -51, 54, -51, -47, -58,
	-1, -1, -1, -1, 6, -15, -15, -1, -1, 6,
	-54, 53, -65, 9, -23, 6, 47, 56, -54, -47,
	40, -47, -1, 6, -47, 49, 50, 49, 50, -15,
	-1, -57, 11, 49, -67, 6, 8, 60, 11, 60,
	49, -55, -62, -15, 6, 8, -58, -1, 50, 10,
	-64, -50, 42, 9, 51, 10, 11, -67, 60, -67,
	-15, -4, 14, -47, -60, 6, -47, 62, 10, -67,
	6, 11, -65, 42, 6, 6, -67, 14, -26, 14,
	14, -52, 17, 16, -47, -47, 14, -10, 25, -15,
	-56, -28, 37, -67, -67, -67, 11, 4, 51, 10,
	-1, -58, -15, -4, -67, -15, -4, 56, 42, 56,
	14, 54, 11, 62, -59, -54, -23, 10, -1, -59,
	-47, 14, 17, 16, -1, -47, 14, -52, 14, -47,
	8, 14, 49, 50, -15, -1, -34, -33, 15, 18,
	27, 28, 14, 16, 37, -61, -15, -21, 62, 11,
	62, -67, -15, -1, -15, 10, 56, 6, -1, 6,
	6, -67, 49, 50, -15, -1, -27, -49, -20, -10,
	31, -67, -67, 14, -1, -47, -47, 14, -47, 4,
	-1, -1, 15, 18, 15, 18, -47, -61, -47, -1,
	6, -67, 6, 14, 14, -10, 15, 18, -61, -59,
	14, -47, 6, 49, 49, 50, -1, -1, -1, -1,
	14, -47, -67, 4, 51, 10, -1, -1, -67, -47,
	-56, -56, -56, -67, -1, -15, 14, 14, -35, 17,
	16, 14, -35, 14, -67, 11, -67, 11, 14, 17,
	16, -1, -56, 14, -67, -67, -1, -56, -56, -56,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 39,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 18, 19, -2, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 13, 33, 34, 35, 36, 37, 38, 0, 0,
	0, 0, 0, 149, 150, 82, 11, 0, 58, 184,
	0, 5, 6, 7, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	-2, 61, 68, 82, 0, 0, 78, 87, 88, 19,
	-2, 21, 22, 23, 30, 13, -2, 82, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 115, 115,
	13, -2, 13, 0, 0, 139, 140, 141, 142, 13,
	0, 192, 196, 80, 0, 11, 133, 134, 0, 131,
	132, 0, 0, 80, 84, 155, 0, 82, 0, 225,
	13, 172, 62, 69, 71, 73, 143, 144, 145, 146,
	147, 148, 174, 0, 223, 224, 176, 0, 83, 11,
	80, 125, 137, 11, 0, 13, 167, 13, 0, 0,
	126, 127, 128, 129, 63, 70, 72, 173, 175, 66,
	11, 0, 109, 115, 116, 111, 0, 0, 11, 0,
	0, 0, 130, 138, 0, 13, 13, 14, 15, 16,
	17, 0, 0, 200, 0, 135, 136, 151, 11, 152,
	12, 11, 11, 0, 19, -2, 0, 185, 186, 187,
	64, 65, 0, -2, 0, 56, 11, 0, 74, 0,
	93, 94, 162, 0, 0, 168, 0, 165, 60, 102,
	115, 0, 0, 0, 112, 114, 102, 118, 13, 120,
	170, 0, 0, 13, 0, 0, 188, 193, 13, 81,
	0, 0, 0, 0, 0, 0, 11, 0, 0, 59,
	67, 0, 198, 57, 0, 89, 90, 0, 0, 0,
	163, 166, 0, 164, 11, 11, 117, 110, 113, 11,
	0, 171, 0, 13, 13, 183, 177, 0, 179, 189,
	13, 199, 201, 202, 39, 204, 205, 206, 0, 0,
	208, 211, 226, 13, 0, 13, 85, 86, 153, 0,
	154, 0, 39, 11, 159, 76, 0, 91, 75, 79,
	169, 0, 103, 104, 39, 106, 107, 0, 99, 194,
	0, 102, 0, 119, 13, 181, 182, 178, 190, 0,
	13, 0, 0, 0, 0, 0, 0, 13, 11, 0,
	0, 156, 0, 95, 108, 195, 0, 0, 197, 11,
	97, 180, 13, 200, -2, -2, 209, 210, 212, 213,
	227, 11, 228, 0, 0, 77, 100, 101, 0, 191,
	0, 0, 0, 229, 11, 11, 96, 214, 0, 0,
	200, 216, 0, 218, 157, 11, 160, 11, 215, 0,
	200, 200, 207, 217, 158, 161, 200, 207, 207, 207,
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
	62, 63, 64, 65, 66,
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
		//line parser.y:190
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:192
		{
		}
	case 3:
		//line parser.y:194
		{
		}
	case 4:
		//line parser.y:196
		{
		}
	case 5:
		//line parser.y:198
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:200
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:202
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:208
		{
		}
	case 11:
		//line parser.y:210
		{
		}
	case 12:
		//line parser.y:211
		{
		}
	case 13:
		//line parser.y:214
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 14:
		//line parser.y:216
		{
		}
	case 15:
		//line parser.y:218
		{
		}
	case 16:
		//line parser.y:220
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 17:
		//line parser.y:222
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 18:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
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
		//line parser.y:232
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 57:
		//line parser.y:239
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-2].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 58:
		//line parser.y:246
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 59:
		//line parser.y:250
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 60:
		//line parser.y:257
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 61:
		//line parser.y:264
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 62:
		//line parser.y:271
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 63:
		//line parser.y:278
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 64:
		//line parser.y:285
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 65:
		//line parser.y:293
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 66:
		//line parser.y:301
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 67:
		//line parser.y:309
		{
			methodName := RubyS[Rubypt-2].genericValue.(ast.BareReference).Name + "="
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: methodName},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 68:
		//line parser.y:320
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 69:
		//line parser.y:327
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 70:
		//line parser.y:335
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 71:
		//line parser.y:343
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 72:
		//line parser.y:351
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 73:
		//line parser.y:359
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 74:
		//line parser.y:369
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 75:
		//line parser.y:379
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 76:
		//line parser.y:388
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 77:
		//line parser.y:390
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 78:
		//line parser.y:392
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 79:
		//line parser.y:394
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 80:
		//line parser.y:397
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 81:
		//line parser.y:399
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 82:
		//line parser.y:401
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 83:
		//line parser.y:403
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 84:
		//line parser.y:405
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:407
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 86:
		//line parser.y:409
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 87:
		//line parser.y:413
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 88:
		//line parser.y:415
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 89:
		//line parser.y:417
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 90:
		//line parser.y:419
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-1].genericValue)
		}
	case 91:
		//line parser.y:421
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 92:
		//line parser.y:424
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 93:
		//line parser.y:426
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 94:
		//line parser.y:428
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 95:
		//line parser.y:434
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 96:
		//line parser.y:442
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-7].genericValue,
				Name:   RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-4].genericSlice,
				Body:   RubyS[Rubypt-2].genericSlice,
			}
		}
	case 97:
		//line parser.y:451
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 98:
		//line parser.y:459
		{
		}
	case 99:
		//line parser.y:461
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 100:
		//line parser.y:463
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{RubyS[Rubypt-2].genericValue}}
		}
	case 101:
		//line parser.y:465
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{ast.Break{}},
			}
		}
	case 102:
		//line parser.y:473
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 103:
		//line parser.y:475
		{
		}
	case 104:
		//line parser.y:477
		{
		}
	case 105:
		//line parser.y:479
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 106:
		//line parser.y:481
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 107:
		//line parser.y:483
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 108:
		//line parser.y:485
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-1].genericSlice...)
		}
	case 109:
		//line parser.y:488
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 110:
		//line parser.y:490
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 111:
		//line parser.y:493
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 112:
		//line parser.y:495
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 113:
		//line parser.y:497
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 114:
		//line parser.y:499
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 115:
		//line parser.y:501
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 116:
		//line parser.y:503
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 117:
		//line parser.y:507
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 118:
		//line parser.y:512
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 119:
		//line parser.y:519
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 120:
		//line parser.y:529
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 121:
		//line parser.y:538
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 122:
		//line parser.y:544
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 123:
		//line parser.y:552
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 124:
		//line parser.y:556
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 125:
		//line parser.y:561
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 126:
		//line parser.y:568
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 127:
		//line parser.y:575
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 128:
		//line parser.y:582
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 129:
		//line parser.y:589
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 130:
		//line parser.y:596
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 131:
		//line parser.y:604
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 132:
		//line parser.y:606
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 133:
		//line parser.y:609
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 134:
		//line parser.y:611
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 135:
		//line parser.y:614
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 136:
		//line parser.y:616
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 137:
		//line parser.y:619
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 138:
		//line parser.y:621
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 139:
		//line parser.y:623
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 140:
		//line parser.y:624
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 141:
		//line parser.y:625
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 142:
		//line parser.y:626
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 143:
		//line parser.y:629
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 144:
		//line parser.y:638
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 145:
		//line parser.y:647
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 146:
		//line parser.y:656
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 147:
		//line parser.y:665
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 148:
		//line parser.y:674
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 149:
		//line parser.y:682
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 150:
		//line parser.y:683
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 151:
		//line parser.y:685
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 152:
		//line parser.y:687
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 153:
		//line parser.y:690
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 154:
		//line parser.y:698
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 155:
		//line parser.y:706
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 156:
		//line parser.y:708
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 157:
		//line parser.y:715
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 158:
		//line parser.y:722
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 159:
		//line parser.y:730
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 160:
		//line parser.y:737
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 161:
		//line parser.y:744
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 162:
		//line parser.y:752
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 163:
		//line parser.y:754
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 164:
		//line parser.y:761
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 165:
		//line parser.y:768
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 166:
		//line parser.y:771
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 167:
		//line parser.y:773
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 168:
		//line parser.y:775
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 169:
		//line parser.y:777
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 170:
		//line parser.y:780
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 171:
		//line parser.y:787
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 172:
		//line parser.y:795
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 173:
		//line parser.y:802
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 174:
		//line parser.y:809
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 175:
		//line parser.y:816
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 176:
		//line parser.y:823
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 177:
		//line parser.y:830
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 178:
		//line parser.y:837
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 179:
		//line parser.y:845
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 180:
		//line parser.y:854
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 181:
		//line parser.y:861
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 182:
		//line parser.y:868
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 183:
		//line parser.y:875
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 184:
		//line parser.y:882
		{
		}
	case 185:
		//line parser.y:883
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 186:
		//line parser.y:884
		{
		}
	case 187:
		//line parser.y:887
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 188:
		//line parser.y:890
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 189:
		//line parser.y:898
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 190:
		//line parser.y:900
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 191:
		//line parser.y:909
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
	case 192:
		//line parser.y:924
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 193:
		//line parser.y:926
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 194:
		//line parser.y:929
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 195:
		//line parser.y:931
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 196:
		//line parser.y:934
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 197:
		//line parser.y:943
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 198:
		//line parser.y:952
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 199:
		//line parser.y:961
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 200:
		//line parser.y:964
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 201:
		//line parser.y:966
		{
		}
	case 202:
		//line parser.y:968
		{
		}
	case 203:
		//line parser.y:970
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 204:
		//line parser.y:972
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 205:
		//line parser.y:974
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 206:
		//line parser.y:976
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 207:
		//line parser.y:978
		{
		}
	case 208:
		//line parser.y:980
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 209:
		//line parser.y:982
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 210:
		//line parser.y:984
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 211:
		//line parser.y:986
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 212:
		//line parser.y:988
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 213:
		//line parser.y:990
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 214:
		//line parser.y:993
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 215:
		//line parser.y:1000
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 216:
		//line parser.y:1008
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 217:
		//line parser.y:1015
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 218:
		//line parser.y:1023
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 219:
		//line parser.y:1031
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 220:
		//line parser.y:1038
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 221:
		//line parser.y:1045
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 222:
		//line parser.y:1052
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 223:
		//line parser.y:1060
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 224:
		//line parser.y:1063
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 225:
		//line parser.y:1065
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 226:
		//line parser.y:1068
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 227:
		//line parser.y:1070
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 228:
		//line parser.y:1073
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 229:
		//line parser.y:1075
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	}
	goto Rubystack /* stack new state and value */
}
