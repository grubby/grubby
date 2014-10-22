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

//line parser.y:1068

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 28,
	52, 122,
	-2, 20,
	-1, 90,
	10, 82,
	11, 82,
	-2, 183,
	-1, 100,
	52, 122,
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
	52, 122,
	-2, 120,
	-1, 224,
	52, 123,
	-2, 121,
	-1, 232,
	10, 82,
	11, 82,
	-2, 183,
	-1, 379,
	27, 199,
	28, 199,
	-2, 13,
	-1, 380,
	27, 199,
	28, 199,
	-2, 13,
}

const RubyNprod = 229
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 2332

var RubyAct = []int{

	9, 193, 346, 259, 292, 36, 401, 191, 92, 141,
	91, 142, 177, 190, 326, 98, 74, 2, 3, 194,
	145, 105, 219, 136, 219, 137, 176, 217, 97, 268,
	117, 287, 20, 146, 4, 327, 389, 219, 286, 251,
	111, 211, 113, 114, 270, 89, 110, 109, 120, 122,
	72, 71, 219, 290, 199, 112, 219, 211, 74, 133,
	195, 135, 74, 74, 375, 325, 143, 73, 132, 196,
	150, 26, 99, 68, 100, 69, 218, 153, 154, 155,
	156, 138, 158, 159, 160, 161, 149, 163, 164, 165,
	216, 168, 379, 380, 170, 171, 289, 378, 144, 219,
	97, 70, 167, 74, 63, 64, 169, 107, 168, 366,
	369, 370, 124, 78, 185, 186, 274, 74, 175, 179,
	410, 107, 266, 144, 178, 65, 144, 66, 372, 62,
	61, 373, 208, 197, 78, 390, 264, 204, 205, 115,
	144, 194, 116, 123, 192, 219, 222, 266, 168, 76,
	77, 212, 219, 97, 108, 416, 213, 413, 412, 225,
	249, 229, 75, 230, 113, 114, 86, 332, 235, 408,
	76, 77, 78, 74, 239, 208, 131, 112, 275, 208,
	144, 233, 195, 75, 319, 147, 320, 86, 173, 240,
	236, 196, 307, 411, 238, 413, 412, 294, 249, 208,
	250, 208, 121, 360, 208, 256, 361, 321, 76, 77,
	224, 248, 267, 265, 76, 77, 358, 377, 254, 359,
	353, 75, 300, 299, 97, 86, 298, 75, 300, 299,
	174, 86, 368, 168, 280, 277, 217, 283, 271, 247,
	217, 272, 273, 208, 279, 198, 208, 200, 337, 388,
	281, 293, 284, 119, 203, 118, 282, 234, 235, 296,
	74, 336, 334, 208, 208, 144, 304, 253, 214, 311,
	215, 133, 323, 139, 252, 140, 329, 331, 244, 202,
	322, 189, 184, 323, 172, 152, 355, 67, 96, 209,
	221, 5, 243, 341, 210, 328, 220, 341, 208, 1,
	242, 344, 245, 208, 324, 148, 57, 208, 56, 55,
	54, 53, 52, 338, 17, 324, 16, 348, 15, 14,
	44, 313, 133, 314, 22, 23, 24, 25, 269, 262,
	263, 363, 343, 12, 11, 125, 126, 127, 128, 129,
	130, 345, 21, 10, 19, 13, 18, 371, 133, 134,
	37, 367, 208, 208, 40, 208, 39, 374, 34, 33,
	35, 32, 0, 208, 151, 208, 0, 0, 0, 0,
	157, 0, 0, 0, 0, 162, 0, 208, 0, 166,
	0, 297, 0, 0, 0, 387, 302, 208, 0, 0,
	399, 306, 0, 0, 208, 311, 311, 311, 180, 181,
	182, 183, 405, 0, 0, 187, 188, 397, 394, 395,
	396, 0, 0, 201, 0, 0, 311, 0, 0, 407,
	409, 311, 311, 311, 0, 351, 352, 0, 0, 417,
	0, 418, 354, 415, 0, 0, 0, 0, 226, 0,
	0, 0, 0, 420, 421, 362, 0, 364, 0, 422,
	0, 0, 26, 27, 68, 28, 69, 0, 0, 0,
	41, 303, 49, 261, 260, 50, 42, 43, 0, 59,
	0, 51, 0, 0, 0, 376, 0, 0, 0, 58,
	0, 203, 70, 60, 0, 63, 64, 0, 386, 0,
	45, 46, 47, 48, 0, 0, 206, 207, 0, 0,
	0, 0, 393, 0, 262, 263, 65, 0, 66, 0,
	62, 61, 0, 0, 0, 26, 27, 68, 28, 69,
	0, 278, 0, 41, 404, 315, 403, 402, 316, 42,
	43, 0, 59, 0, 51, 0, 0, 317, 318, 0,
	0, 295, 58, 0, 0, 70, 60, 0, 63, 64,
	301, 0, 0, 45, 46, 47, 48, 0, 312, 309,
	310, 0, 0, 0, 0, 330, 0, 0, 0, 65,
	0, 66, 0, 62, 61, 0, 335, 0, 0, 0,
	0, 0, 342, 0, 0, 0, 342, 0, 0, 350,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	26, 27, 68, 28, 69, 356, 357, 0, 41, 400,
	315, 403, 402, 316, 42, 43, 365, 59, 0, 51,
	0, 0, 317, 318, 0, 0, 0, 58, 0, 0,
	70, 60, 0, 63, 64, 0, 0, 29, 45, 46,
	47, 48, 0, 0, 309, 310, 0, 0, 381, 382,
	383, 384, 0, 0, 65, 0, 66, 0, 62, 61,
	0, 0, 391, 392, 0, 101, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 398, 0,
	0, 0, 0, 0, 312, 312, 312, 0, 0, 0,
	0, 0, 414, 0, 0, 0, 101, 0, 101, 0,
	0, 0, 419, 101, 0, 312, 0, 0, 0, 0,
	312, 312, 312, 78, 101, 101, 101, 101, 0, 101,
	101, 101, 101, 0, 101, 101, 101, 0, 101, 0,
	38, 101, 101, 0, 0, 0, 0, 101, 0, 26,
	99, 68, 100, 232, 0, 101, 105, 0, 0, 76,
	77, 101, 101, 0, 79, 80, 81, 0, 104, 0,
	0, 0, 75, 84, 82, 83, 86, 0, 0, 70,
	0, 0, 63, 64, 0, 0, 231, 0, 0, 0,
	0, 0, 0, 101, 0, 101, 0, 0, 0, 104,
	101, 104, 0, 65, 0, 106, 104, 62, 61, 26,
	99, 68, 100, 69, 0, 0, 105, 104, 104, 104,
	104, 101, 104, 104, 104, 104, 0, 104, 104, 104,
	0, 104, 0, 0, 104, 104, 0, 0, 0, 70,
	104, 0, 63, 64, 0, 0, 0, 0, 104, 0,
	0, 0, 0, 219, 104, 104, 0, 0, 0, 101,
	285, 0, 0, 65, 0, 106, 0, 62, 61, 0,
	0, 101, 0, 0, 0, 0, 0, 0, 0, 0,
	101, 101, 0, 0, 101, 0, 104, 0, 104, 0,
	0, 0, 0, 104, 0, 26, 99, 68, 100, 90,
	0, 95, 105, 0, 26, 99, 68, 100, 69, 0,
	0, 0, 0, 0, 104, 0, 0, 0, 101, 101,
	0, 0, 0, 0, 101, 70, 0, 0, 63, 64,
	101, 0, 94, 0, 70, 0, 0, 63, 64, 0,
	0, 0, 26, 99, 68, 100, 69, 0, 219, 93,
	0, 106, 104, 62, 61, 333, 0, 0, 65, 0,
	66, 30, 62, 61, 104, 0, 0, 0, 0, 101,
	0, 0, 70, 104, 104, 63, 64, 104, 0, 0,
	0, 0, 0, 0, 0, 0, 219, 0, 0, 102,
	0, 0, 0, 0, 0, 101, 65, 0, 66, 0,
	62, 61, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 104, 104, 0, 0, 0, 0, 104, 0, 0,
	102, 0, 102, 104, 0, 0, 0, 102, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 101, 102, 102,
	102, 102, 78, 102, 102, 102, 102, 0, 102, 102,
	102, 0, 102, 0, 31, 102, 102, 0, 0, 0,
	0, 102, 104, 0, 0, 0, 0, 0, 0, 102,
	26, 99, 68, 100, 90, 102, 102, 105, 76, 77,
	0, 0, 103, 79, 80, 81, 0, 0, 104, 0,
	0, 75, 84, 82, 83, 86, 0, 0, 237, 0,
	70, 0, 0, 63, 64, 0, 0, 102, 0, 102,
	0, 0, 0, 103, 102, 103, 276, 0, 0, 0,
	103, 0, 0, 0, 93, 0, 106, 0, 62, 61,
	104, 103, 103, 103, 103, 102, 103, 103, 103, 103,
	0, 103, 103, 103, 0, 103, 0, 0, 103, 103,
	0, 0, 0, 0, 103, 0, 26, 99, 68, 100,
	69, 0, 103, 105, 0, 0, 0, 0, 103, 103,
	0, 0, 0, 102, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 102, 70, 0, 0, 63,
	64, 0, 0, 0, 102, 102, 0, 0, 102, 0,
	103, 0, 103, 0, 0, 0, 0, 103, 0, 0,
	65, 0, 106, 0, 62, 61, 26, 223, 68, 100,
	69, 0, 0, 26, 99, 68, 100, 90, 103, 0,
	105, 0, 102, 102, 0, 0, 0, 0, 102, 0,
	0, 0, 0, 0, 102, 0, 70, 0, 0, 63,
	64, 0, 0, 70, 0, 0, 63, 64, 0, 0,
	219, 0, 0, 0, 0, 0, 103, 0, 0, 0,
	65, 0, 66, 0, 62, 61, 0, 93, 103, 106,
	0, 62, 61, 102, 0, 0, 0, 103, 103, 0,
	0, 103, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 102,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 26,
	27, 68, 28, 69, 0, 103, 103, 41, 406, 315,
	0, 103, 316, 42, 43, 0, 59, 103, 51, 0,
	0, 317, 318, 0, 0, 0, 58, 0, 0, 70,
	60, 102, 63, 64, 0, 0, 0, 45, 46, 47,
	48, 0, 0, 309, 310, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 0, 66, 103, 62, 61, 0,
	0, 26, 27, 68, 28, 69, 0, 0, 0, 41,
	308, 315, 0, 0, 316, 42, 43, 0, 59, 0,
	51, 0, 103, 317, 318, 0, 0, 0, 58, 0,
	0, 70, 60, 0, 63, 64, 0, 0, 0, 45,
	46, 47, 48, 0, 0, 309, 310, 0, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 66, 0, 62,
	61, 0, 0, 0, 103, 26, 27, 68, 28, 69,
	0, 0, 0, 41, 258, 49, 261, 260, 50, 42,
	43, 0, 59, 0, 51, 0, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 70, 60, 0, 63, 64,
	0, 0, 0, 45, 46, 47, 48, 0, 0, 206,
	207, 0, 26, 27, 68, 28, 69, 0, 0, 65,
	41, 66, 49, 62, 61, 50, 42, 43, 0, 59,
	0, 51, 0, 0, 0, 0, 0, 0, 0, 58,
	0, 0, 70, 60, 0, 63, 64, 0, 0, 0,
	45, 46, 47, 48, 0, 0, 6, 7, 0, 0,
	0, 0, 0, 0, 0, 0, 65, 0, 66, 0,
	62, 61, 0, 8, 26, 27, 68, 28, 69, 0,
	0, 0, 41, 0, 315, 0, 0, 316, 42, 43,
	0, 59, 0, 51, 0, 0, 317, 318, 0, 0,
	0, 58, 0, 0, 70, 60, 0, 63, 64, 0,
	0, 0, 45, 46, 47, 48, 0, 0, 309, 310,
	0, 26, 27, 68, 28, 69, 0, 0, 65, 41,
	66, 49, 62, 61, 50, 42, 43, 0, 59, 0,
	51, 266, 0, 0, 0, 0, 0, 347, 58, 0,
	0, 70, 60, 0, 63, 64, 0, 0, 0, 45,
	46, 47, 48, 0, 0, 339, 340, 0, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 66, 0, 62,
	61, 26, 27, 68, 28, 69, 0, 0, 0, 41,
	385, 49, 0, 0, 50, 42, 43, 0, 59, 0,
	51, 0, 0, 0, 0, 0, 0, 0, 58, 0,
	0, 70, 60, 0, 63, 64, 0, 0, 0, 45,
	46, 47, 48, 0, 0, 206, 207, 0, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 66, 0, 62,
	61, 26, 27, 68, 28, 69, 0, 0, 0, 41,
	349, 49, 0, 0, 50, 42, 43, 0, 59, 0,
	51, 0, 0, 0, 0, 0, 0, 0, 58, 0,
	0, 70, 60, 0, 63, 64, 0, 0, 0, 45,
	46, 47, 48, 0, 0, 206, 207, 0, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 66, 0, 62,
	61, 26, 27, 68, 28, 69, 0, 0, 0, 41,
	305, 49, 0, 0, 50, 42, 43, 0, 59, 0,
	51, 0, 0, 0, 0, 0, 0, 0, 58, 0,
	0, 70, 60, 0, 63, 64, 0, 0, 0, 45,
	46, 47, 48, 0, 0, 206, 207, 0, 26, 27,
	68, 28, 69, 0, 0, 65, 41, 66, 49, 62,
	61, 50, 42, 43, 0, 59, 0, 51, 0, 0,
	0, 0, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 206, 207, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 66, 291, 62, 61, 26, 27,
	68, 28, 69, 0, 0, 0, 41, 288, 49, 0,
	0, 50, 42, 43, 0, 59, 0, 51, 0, 0,
	0, 0, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 206, 207, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 66, 0, 62, 61, 26, 27,
	68, 28, 69, 0, 0, 0, 41, 257, 49, 0,
	0, 50, 42, 43, 0, 59, 0, 51, 0, 0,
	0, 0, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 206, 207, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 66, 0, 62, 61, 26, 27,
	68, 28, 69, 0, 0, 0, 41, 255, 49, 0,
	0, 50, 42, 43, 0, 59, 0, 51, 0, 0,
	0, 0, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 206, 207, 0, 26, 27, 68, 28, 69,
	0, 0, 65, 41, 66, 49, 62, 61, 50, 42,
	43, 0, 59, 0, 51, 0, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 70, 60, 0, 63, 64,
	0, 0, 0, 45, 46, 47, 48, 0, 0, 206,
	207, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	0, 66, 246, 62, 61, 26, 27, 68, 28, 69,
	0, 0, 0, 41, 241, 49, 0, 0, 50, 42,
	43, 0, 59, 0, 51, 0, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 70, 60, 0, 63, 64,
	0, 0, 0, 45, 46, 47, 48, 0, 0, 206,
	207, 0, 26, 27, 68, 28, 69, 0, 0, 65,
	41, 66, 49, 62, 61, 50, 42, 43, 0, 59,
	0, 51, 0, 0, 0, 0, 0, 0, 0, 58,
	0, 0, 70, 60, 0, 63, 64, 0, 0, 0,
	45, 46, 47, 48, 0, 0, 206, 207, 0, 26,
	27, 68, 28, 69, 228, 0, 65, 41, 66, 49,
	62, 61, 50, 42, 43, 0, 59, 0, 51, 0,
	0, 0, 0, 0, 0, 0, 58, 0, 0, 70,
	60, 0, 63, 64, 0, 0, 0, 45, 46, 47,
	48, 0, 0, 0, 227, 0, 26, 27, 68, 28,
	69, 0, 0, 65, 41, 66, 49, 62, 61, 50,
	42, 43, 0, 59, 0, 51, 0, 0, 0, 0,
	0, 0, 0, 58, 0, 0, 70, 60, 78, 63,
	64, 0, 0, 0, 45, 46, 47, 48, 0, 0,
	0, 0, 85, 0, 0, 0, 0, 0, 0, 0,
	65, 0, 66, 0, 62, 61, 0, 87, 88, 0,
	0, 0, 0, 0, 76, 77, 0, 0, 0, 79,
	80, 81, 0, 0, 0, 0, 0, 75, 84, 82,
	83, 86,
}
var RubyPact = []int{

	-32, 1477, -1000, -1000, -1000, 1, -1000, -1000, -1000, 2274,
	-1000, -1000, -1000, 27, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 880, 112, 5,
	4, -2, -1000, -1000, -1000, -1000, -1000, -1000, 124, -23,
	-1000, 249, 194, 194, 101, 2241, 2241, 2241, 2241, 2241,
	2241, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 66, 2241,
	66, 17, 267, -1000, -1000, 66, -1000, -19, 176, -1000,
	9, -1000, -1000, -1000, 2241, 279, 66, 66, 66, 66,
	2241, 66, 66, 66, 66, 2241, 66, 66, 66, 2241,
	66, -1000, 95, 66, 66, 278, 177, 168, -1000, 1208,
	98, -1000, -1000, -1000, 2, -28, -28, 66, 2241, 2241,
	2241, 2241, 276, 66, 66, 2241, 2241, 275, 135, 135,
	14, -1000, -1000, 2241, 273, 47, 47, 47, 47, 47,
	88, 2147, 46, 168, 102, 168, -1000, -1000, 262, -1000,
	-1000, 30, 16, 709, -1000, 1201, 202, 66, 2194, -1000,
	-28, 47, 734, 168, 168, 168, 168, 47, 168, 168,
	168, 168, 47, 130, 168, 168, 47, 247, 709, -1000,
	1028, 168, -1000, 1141, 2100, -1000, 272, -1000, 2040, 229,
	47, 47, 47, 47, -1000, 168, 168, 47, 47, -1000,
	-1000, 149, 13, -1000, -3, 268, 261, -1000, 1993, 194,
	1933, 47, -1000, 1430, -1000, -1000, -1000, -1000, 2274, 47,
	122, 66, -1000, 7, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 105, 174, 1055, -1000, 225, 47, -1000, -1000, 95,
	-1000, 2241, 66, 66, 9, -1000, 794, -4, -25, 168,
	-1000, -1000, 1873, 42, -1000, 1813, -1000, -1000, -12, 13,
	187, 2241, -1000, -1000, -12, -1000, -1000, -1000, -1000, 212,
	2241, -1000, 447, 1766, -1000, -1000, 184, 168, 1366, 170,
	66, 927, 3, -27, -1000, 2241, 66, -1000, 47, 157,
	168, -1000, 889, 168, -1000, 256, 2241, 255, -1000, -1000,
	242, -1000, 1586, -1000, -1000, 47, 1586, 1706, -1000, 2241,
	-1000, 47, 2147, -1000, 206, -1000, 2147, 282, -1000, -1000,
	-1000, 2274, 47, -1000, -1000, 2241, 2241, 201, 188, -1000,
	-1000, 66, 46, 709, -1000, -1000, 2241, -1000, 103, 2274,
	47, 168, -1000, 226, -1000, 47, -1000, -1000, 96, -1000,
	-1000, 2274, 47, -1000, 97, 113, -1000, 66, 50, -1000,
	47, 2147, 2147, -1000, 2147, 211, 48, 43, 2241, 2241,
	2241, 2241, 1646, 46, 2147, 245, -15, -12, 125, -1000,
	-1000, -1000, 2241, 2241, 46, -1000, 2147, -1000, -1000, -1000,
	-1000, 47, 47, 47, 47, -1000, 2147, -12, 2241, 66,
	-1000, 47, 47, 2147, 595, 510, 1304, -12, 158, 109,
	-1000, 179, 2241, -1000, -1000, 141, -1000, -12, -1000, -12,
	-1000, -1000, 2241, -1000, 47, 1539, -1000, -12, -12, 47,
	1539, 1539, 1539,
}
var RubyPgo = []int{

	0, 289, 361, 360, 15, 359, 358, 356, 1044, 354,
	2, 350, 346, 345, 344, 0, 951, 730, 343, 342,
	341, 32, 334, 1, 637, 333, 5, 332, 328, 327,
	326, 325, 324, 323, 321, 6, 320, 319, 318, 316,
	314, 312, 311, 310, 309, 308, 306, 124, 305, 301,
	10, 12, 3, 299, 13, 296, 29, 294, 11, 4,
	292, 9, 290, 288, 8, 7, 287, 20,
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
	64, 64, 63, 63, 63, 18, 18, 27, 27, 27,
	27, 59, 59, 59, 59, 59, 59, 59, 54, 54,
	23, 23, 23, 23, 65, 65, 65, 22, 22, 25,
	26, 26, 66, 66, 13, 13, 13, 13, 13, 13,
	8, 8, 24, 24, 16, 16, 36, 36, 37, 38,
	39, 40, 41, 42, 43, 44, 45, 46, 2, 5,
	6, 6, 3, 3, 55, 55, 55, 55, 62, 62,
	62, 4, 4, 4, 4, 51, 60, 60, 60, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 52,
	52, 52, 52, 48, 48, 48, 7, 14, 10, 10,
	10, 57, 57, 49, 49, 19, 20, 11, 32, 56,
	56, 56, 56, 56, 56, 56, 33, 33, 33, 33,
	33, 33, 33, 34, 34, 34, 34, 34, 35, 35,
	35, 35, 31, 30, 9, 29, 29, 28, 28,
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
	4, 5, 1, 3, 3, 7, 7, 0, 1, 3,
	3, 0, 2, 2, 2, 2, 2, 3, 1, 3,
	1, 2, 3, 2, 0, 1, 3, 4, 6, 4,
	1, 3, 1, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 3, 3, 3, 3, 2, 2,
	2, 2, 3, 3, 3, 3, 3, 3, 1, 1,
	3, 3, 5, 5, 0, 4, 7, 8, 3, 7,
	8, 3, 4, 4, 3, 3, 0, 1, 3, 4,
	5, 3, 3, 3, 3, 3, 5, 6, 5, 4,
	3, 3, 2, 0, 2, 2, 3, 4, 2, 3,
	5, 0, 2, 1, 2, 2, 2, 5, 5, 0,
	2, 2, 2, 2, 2, 2, 0, 1, 3, 3,
	1, 3, 3, 5, 6, 5, 6, 5, 4, 3,
	3, 2, 3, 3, 2, 5, 7, 4, 5,
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
	-54, -65, 9, -23, 6, 47, 56, -54, -47, 40,
	-47, -1, 6, -47, 49, 50, 49, 50, -15, -1,
	-57, 11, 49, -67, 6, 8, 60, 11, 60, 49,
	-55, -62, -15, 6, 8, -58, -1, 50, 10, -64,
	-50, 42, 9, 51, 10, 11, -67, 60, -67, -15,
	-4, 14, -47, -60, 6, -47, 62, 10, -67, 11,
	-65, 42, 6, 6, -67, 14, -26, 14, 14, -52,
	17, 16, -47, -47, 14, -10, 25, -15, -56, -28,
	37, -67, -67, -67, 11, 4, 51, 10, -1, -58,
	-15, -4, -67, -15, -4, 56, 42, 56, 14, 54,
	11, 62, -59, -23, 10, -1, -59, -47, 14, 17,
	16, -1, -47, 14, -52, 14, -47, 8, 14, 49,
	50, -15, -1, -34, -33, 15, 18, 27, 28, 14,
	16, 37, -61, -15, -21, 62, 11, 62, -67, -15,
	-1, -15, 10, 56, 6, -1, 6, 6, -67, 49,
	50, -15, -1, -27, -49, -20, -10, 31, -67, 14,
	-1, -47, -47, 14, -47, 4, -1, -1, 15, 18,
	15, 18, -47, -61, -47, -1, 6, -67, 6, 14,
	14, -10, 15, 18, -61, 14, -47, 6, 49, 49,
	50, -1, -1, -1, -1, 14, -47, -67, 4, 51,
	10, -1, -1, -47, -56, -56, -56, -67, -1, -15,
	14, -35, 17, 16, 14, -35, 14, -67, 11, -67,
	11, 14, 17, 16, -1, -56, 14, -67, -67, -1,
	-56, -56, -56,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 39,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 18, 19, -2, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 13, 33, 34, 35, 36, 37, 38, 0, 0,
	0, 0, 0, 148, 149, 82, 11, 0, 58, 183,
	0, 5, 6, 7, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	-2, 61, 68, 82, 0, 0, 78, 87, 88, 19,
	-2, 21, 22, 23, 30, 13, -2, 82, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 114, 114,
	13, -2, 13, 0, 0, 138, 139, 140, 141, 13,
	0, 191, 195, 80, 0, 11, 132, 133, 0, 130,
	131, 0, 0, 80, 84, 154, 0, 82, 0, 224,
	13, 171, 62, 69, 71, 73, 142, 143, 144, 145,
	146, 147, 173, 0, 222, 223, 175, 0, 83, 11,
	80, 124, 136, 11, 0, 13, 166, 13, 0, 0,
	125, 126, 127, 128, 63, 70, 72, 172, 174, 66,
	11, 108, 114, 115, 110, 0, 0, 11, 0, 0,
	0, 129, 137, 0, 13, 13, 14, 15, 16, 17,
	0, 0, 199, 0, 134, 135, 150, 11, 151, 12,
	11, 11, 0, 19, -2, 0, 184, 185, 186, 64,
	65, 0, -2, 0, 56, 11, 0, 74, 0, 93,
	94, 161, 0, 0, 167, 0, 164, 60, 101, 0,
	0, 0, 111, 113, 101, 117, 13, 119, 169, 0,
	0, 13, 0, 0, 187, 192, 13, 81, 0, 0,
	0, 0, 0, 0, 11, 0, 0, 59, 67, 0,
	197, 57, 0, 89, 90, 0, 0, 0, 162, 165,
	0, 163, 11, 116, 109, 112, 11, 0, 170, 0,
	13, 13, 182, 176, 0, 178, 188, 13, 198, 200,
	201, 39, 203, 204, 205, 0, 0, 207, 210, 225,
	13, 0, 13, 85, 86, 152, 0, 153, 0, 39,
	11, 158, 76, 0, 91, 75, 79, 168, 0, 102,
	103, 39, 105, 106, 0, 98, 193, 0, 0, 118,
	13, 180, 181, 177, 189, 0, 13, 0, 0, 0,
	0, 0, 0, 13, 11, 0, 0, 155, 0, 95,
	107, 194, 0, 0, 196, 96, 179, 13, 199, -2,
	-2, 208, 209, 211, 212, 226, 11, 227, 0, 0,
	77, 99, 100, 190, 0, 0, 0, 228, 11, 11,
	213, 0, 0, 199, 215, 0, 217, 156, 11, 159,
	11, 214, 0, 199, 199, 206, 216, 157, 160, 199,
	206, 206, 206,
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
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 97:
		//line parser.y:450
		{
		}
	case 98:
		//line parser.y:452
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 99:
		//line parser.y:454
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{RubyS[Rubypt-2].genericValue}}
		}
	case 100:
		//line parser.y:456
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{ast.Break{}},
			}
		}
	case 101:
		//line parser.y:464
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 102:
		//line parser.y:466
		{
		}
	case 103:
		//line parser.y:468
		{
		}
	case 104:
		//line parser.y:470
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 105:
		//line parser.y:472
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 106:
		//line parser.y:474
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 107:
		//line parser.y:476
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-1].genericSlice...)
		}
	case 108:
		//line parser.y:479
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 109:
		//line parser.y:481
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 110:
		//line parser.y:484
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 111:
		//line parser.y:486
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 112:
		//line parser.y:488
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 113:
		//line parser.y:490
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 114:
		//line parser.y:492
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 115:
		//line parser.y:494
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 116:
		//line parser.y:498
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 117:
		//line parser.y:503
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 118:
		//line parser.y:510
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 119:
		//line parser.y:520
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 120:
		//line parser.y:529
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 121:
		//line parser.y:535
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 122:
		//line parser.y:543
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 123:
		//line parser.y:547
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 124:
		//line parser.y:552
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 125:
		//line parser.y:559
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 126:
		//line parser.y:566
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 127:
		//line parser.y:573
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 128:
		//line parser.y:580
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 129:
		//line parser.y:587
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 130:
		//line parser.y:595
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 131:
		//line parser.y:597
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 132:
		//line parser.y:600
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 133:
		//line parser.y:602
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 134:
		//line parser.y:605
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 135:
		//line parser.y:607
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 136:
		//line parser.y:610
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 137:
		//line parser.y:612
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 138:
		//line parser.y:614
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 139:
		//line parser.y:615
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 140:
		//line parser.y:616
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 141:
		//line parser.y:617
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 142:
		//line parser.y:620
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 143:
		//line parser.y:629
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 144:
		//line parser.y:638
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 145:
		//line parser.y:647
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 146:
		//line parser.y:656
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 147:
		//line parser.y:665
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 148:
		//line parser.y:673
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 149:
		//line parser.y:674
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 150:
		//line parser.y:676
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 151:
		//line parser.y:678
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 152:
		//line parser.y:681
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 153:
		//line parser.y:689
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 154:
		//line parser.y:697
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 155:
		//line parser.y:699
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 156:
		//line parser.y:706
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 157:
		//line parser.y:713
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 158:
		//line parser.y:721
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 159:
		//line parser.y:728
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 160:
		//line parser.y:735
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 161:
		//line parser.y:743
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 162:
		//line parser.y:745
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 163:
		//line parser.y:752
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 164:
		//line parser.y:759
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 165:
		//line parser.y:762
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 166:
		//line parser.y:764
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 167:
		//line parser.y:766
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 168:
		//line parser.y:768
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 169:
		//line parser.y:771
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 170:
		//line parser.y:778
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 171:
		//line parser.y:786
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 172:
		//line parser.y:793
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 173:
		//line parser.y:800
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 174:
		//line parser.y:807
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 175:
		//line parser.y:814
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 176:
		//line parser.y:821
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 177:
		//line parser.y:828
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 178:
		//line parser.y:836
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 179:
		//line parser.y:845
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 180:
		//line parser.y:852
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 181:
		//line parser.y:859
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 182:
		//line parser.y:866
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 183:
		//line parser.y:873
		{
		}
	case 184:
		//line parser.y:874
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 185:
		//line parser.y:875
		{
		}
	case 186:
		//line parser.y:878
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 187:
		//line parser.y:881
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 188:
		//line parser.y:889
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 189:
		//line parser.y:891
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 190:
		//line parser.y:900
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
	case 191:
		//line parser.y:915
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 192:
		//line parser.y:917
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 193:
		//line parser.y:920
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 194:
		//line parser.y:922
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 195:
		//line parser.y:925
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 196:
		//line parser.y:934
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 197:
		//line parser.y:943
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 198:
		//line parser.y:952
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 199:
		//line parser.y:955
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 200:
		//line parser.y:957
		{
		}
	case 201:
		//line parser.y:959
		{
		}
	case 202:
		//line parser.y:961
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 203:
		//line parser.y:963
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 204:
		//line parser.y:965
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 205:
		//line parser.y:967
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 206:
		//line parser.y:969
		{
		}
	case 207:
		//line parser.y:971
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 208:
		//line parser.y:973
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 209:
		//line parser.y:975
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 210:
		//line parser.y:977
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 211:
		//line parser.y:979
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 212:
		//line parser.y:981
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 213:
		//line parser.y:984
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 214:
		//line parser.y:991
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 215:
		//line parser.y:999
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 216:
		//line parser.y:1006
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 217:
		//line parser.y:1014
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 218:
		//line parser.y:1022
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 219:
		//line parser.y:1029
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 220:
		//line parser.y:1036
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 221:
		//line parser.y:1043
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 222:
		//line parser.y:1051
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 223:
		//line parser.y:1054
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 224:
		//line parser.y:1056
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 225:
		//line parser.y:1059
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 226:
		//line parser.y:1061
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 227:
		//line parser.y:1064
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 228:
		//line parser.y:1066
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	}
	goto Rubystack /* stack new state and value */
}
