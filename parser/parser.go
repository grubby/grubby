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
const OR_EQUALS = 57390
const WHITESPACE = 57391
const NEWLINE = 57392
const SEMICOLON = 57393
const COLON = 57394
const DOUBLECOLON = 57395
const DOT = 57396
const PIPE = 57397
const SLASH = 57398
const AMPERSAND = 57399
const QUESTIONMARK = 57400
const CARET = 57401
const LBRACKET = 57402
const RBRACKET = 57403
const LBRACE = 57404
const RBRACE = 57405
const DOLLARSIGN = 57406
const ATSIGN = 57407
const FILE_CONST_REF = 57408
const EOF = 57409

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

//line parser.y:1157

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 29,
	53, 126,
	-2, 20,
	-1, 91,
	10, 83,
	11, 83,
	-2, 194,
	-1, 102,
	53, 126,
	-2, 20,
	-1, 108,
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
	51, 13,
	-2, 11,
	-1, 126,
	53, 126,
	-2, 124,
	-1, 236,
	53, 127,
	-2, 125,
	-1, 244,
	10, 83,
	11, 83,
	-2, 194,
	-1, 397,
	27, 210,
	28, 210,
	-2, 13,
	-1, 398,
	27, 210,
	28, 210,
	-2, 13,
}

const RubyNprod = 240
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 2527

var RubyAct = []int{

	9, 307, 418, 146, 260, 313, 21, 272, 201, 205,
	37, 147, 203, 93, 92, 185, 75, 351, 58, 2,
	3, 231, 107, 141, 229, 142, 223, 304, 231, 99,
	184, 122, 151, 281, 352, 301, 4, 407, 231, 300,
	75, 116, 283, 387, 114, 100, 129, 117, 211, 264,
	115, 73, 72, 125, 127, 231, 231, 90, 206, 75,
	138, 204, 140, 137, 112, 75, 223, 148, 74, 350,
	113, 303, 155, 149, 230, 396, 228, 128, 158, 159,
	160, 161, 143, 163, 164, 165, 166, 231, 168, 169,
	170, 174, 173, 79, 224, 175, 176, 178, 149, 207,
	408, 149, 99, 172, 79, 287, 202, 118, 119, 208,
	173, 427, 79, 30, 177, 179, 149, 154, 425, 196,
	197, 187, 75, 183, 109, 75, 206, 409, 206, 77,
	78, 204, 262, 209, 80, 81, 82, 220, 279, 181,
	77, 78, 103, 76, 85, 83, 84, 87, 77, 78,
	249, 234, 381, 173, 76, 382, 109, 110, 99, 149,
	397, 398, 76, 111, 237, 369, 87, 207, 152, 207,
	79, 241, 242, 103, 363, 103, 279, 208, 379, 208,
	103, 380, 251, 220, 226, 279, 227, 220, 144, 79,
	145, 103, 103, 103, 103, 75, 103, 103, 103, 103,
	332, 103, 103, 103, 288, 103, 77, 78, 103, 103,
	103, 220, 126, 220, 267, 103, 220, 263, 245, 344,
	76, 345, 269, 103, 280, 77, 78, 252, 278, 236,
	216, 217, 103, 103, 120, 365, 99, 121, 366, 76,
	77, 78, 346, 357, 247, 173, 294, 318, 262, 297,
	395, 149, 277, 389, 76, 220, 293, 362, 220, 118,
	119, 310, 406, 279, 103, 124, 103, 123, 310, 321,
	316, 103, 317, 75, 361, 359, 220, 220, 433, 266,
	430, 429, 336, 329, 138, 348, 265, 347, 291, 354,
	356, 349, 295, 261, 298, 103, 428, 348, 430, 429,
	290, 229, 256, 349, 214, 221, 200, 5, 27, 101,
	69, 102, 244, 364, 376, 107, 138, 259, 229, 367,
	180, 368, 374, 220, 325, 324, 157, 364, 220, 246,
	247, 323, 220, 325, 324, 68, 98, 103, 71, 233,
	255, 64, 65, 222, 232, 243, 1, 138, 153, 103,
	384, 130, 131, 132, 133, 134, 135, 57, 103, 103,
	56, 55, 103, 66, 54, 108, 139, 63, 62, 310,
	393, 53, 52, 220, 220, 18, 220, 17, 16, 15,
	44, 156, 338, 339, 220, 23, 220, 162, 24, 25,
	26, 282, 167, 312, 14, 220, 171, 103, 103, 364,
	12, 11, 314, 103, 22, 220, 10, 38, 416, 20,
	103, 220, 336, 336, 336, 422, 188, 189, 190, 191,
	192, 193, 194, 195, 13, 19, 198, 199, 40, 103,
	411, 412, 413, 336, 213, 39, 106, 35, 336, 336,
	336, 34, 36, 33, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 432, 0, 0, 0, 0, 238,
	103, 0, 0, 0, 437, 438, 0, 106, 0, 106,
	439, 0, 0, 0, 106, 0, 0, 27, 101, 69,
	102, 91, 0, 97, 107, 106, 106, 106, 106, 0,
	106, 106, 106, 106, 0, 106, 106, 106, 0, 106,
	0, 0, 106, 106, 106, 0, 0, 71, 0, 106,
	64, 65, 0, 0, 95, 0, 0, 106, 150, 0,
	96, 103, 0, 0, 0, 0, 106, 106, 0, 0,
	0, 0, 94, 0, 108, 0, 63, 62, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 292,
	0, 0, 0, 0, 0, 0, 0, 0, 106, 0,
	106, 0, 0, 0, 0, 106, 311, 0, 0, 0,
	319, 0, 0, 311, 0, 0, 0, 0, 0, 326,
	0, 0, 27, 28, 69, 29, 70, 337, 0, 106,
	41, 423, 340, 0, 355, 341, 42, 43, 0, 60,
	0, 51, 0, 0, 342, 343, 360, 0, 0, 59,
	0, 0, 71, 61, 0, 64, 65, 0, 0, 0,
	45, 46, 47, 48, 0, 0, 0, 334, 335, 0,
	371, 106, 0, 0, 0, 0, 0, 66, 0, 67,
	0, 63, 62, 106, 0, 0, 377, 378, 0, 0,
	0, 0, 106, 106, 0, 0, 106, 386, 0, 225,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 390, 391, 0, 311, 0, 0, 0, 0, 31,
	0, 0, 0, 0, 0, 399, 400, 401, 402, 0,
	0, 106, 106, 248, 0, 0, 0, 106, 0, 0,
	250, 0, 0, 0, 106, 0, 0, 0, 104, 0,
	0, 0, 415, 0, 0, 0, 0, 337, 337, 337,
	0, 0, 0, 106, 0, 431, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 436, 0, 0, 337, 104,
	0, 104, 79, 337, 337, 337, 104, 0, 284, 0,
	0, 285, 286, 0, 106, 0, 0, 104, 104, 104,
	104, 32, 104, 104, 104, 104, 296, 104, 104, 104,
	0, 104, 0, 0, 104, 104, 104, 0, 77, 78,
	0, 104, 0, 80, 81, 82, 0, 0, 0, 104,
	105, 0, 76, 85, 83, 84, 87, 0, 104, 104,
	0, 0, 0, 186, 0, 0, 353, 0, 0, 0,
	0, 0, 0, 0, 0, 106, 0, 0, 0, 0,
	0, 105, 0, 105, 0, 0, 0, 0, 105, 0,
	104, 0, 104, 0, 0, 0, 0, 104, 0, 105,
	105, 105, 105, 0, 105, 105, 105, 105, 79, 105,
	105, 105, 0, 105, 0, 136, 105, 105, 105, 0,
	0, 104, 86, 105, 0, 0, 0, 0, 0, 0,
	0, 105, 0, 0, 388, 0, 0, 88, 89, 0,
	105, 105, 0, 0, 77, 78, 0, 0, 0, 80,
	81, 82, 0, 0, 0, 0, 0, 0, 76, 85,
	83, 84, 87, 104, 405, 0, 0, 0, 0, 0,
	0, 182, 105, 0, 105, 104, 0, 0, 0, 105,
	0, 0, 0, 414, 104, 104, 0, 0, 104, 210,
	0, 212, 0, 0, 424, 426, 0, 0, 215, 0,
	0, 0, 0, 105, 434, 0, 435, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 104, 104, 0, 0, 0, 0, 104,
	27, 101, 69, 102, 70, 0, 104, 107, 0, 27,
	101, 69, 102, 70, 0, 105, 0, 254, 0, 257,
	0, 0, 0, 0, 0, 104, 0, 105, 0, 0,
	71, 0, 0, 64, 65, 0, 105, 105, 0, 71,
	105, 0, 64, 65, 0, 231, 0, 0, 0, 0,
	275, 276, 299, 0, 231, 66, 104, 108, 0, 63,
	62, 358, 0, 0, 66, 0, 67, 0, 63, 62,
	0, 0, 0, 0, 0, 105, 105, 0, 0, 0,
	0, 105, 0, 0, 0, 0, 0, 0, 105, 0,
	0, 0, 27, 101, 69, 102, 70, 0, 0, 0,
	0, 0, 0, 322, 0, 0, 0, 105, 327, 0,
	0, 0, 0, 331, 0, 0, 0, 104, 0, 0,
	0, 0, 71, 0, 0, 64, 65, 0, 0, 0,
	0, 27, 28, 69, 29, 70, 0, 231, 105, 41,
	421, 340, 420, 419, 341, 42, 43, 66, 60, 67,
	51, 63, 62, 342, 343, 0, 0, 0, 59, 372,
	373, 71, 61, 0, 64, 65, 375, 0, 0, 45,
	46, 47, 48, 0, 0, 0, 334, 335, 0, 383,
	0, 385, 0, 0, 0, 0, 66, 0, 67, 0,
	63, 62, 0, 0, 27, 101, 69, 102, 70, 105,
	0, 107, 0, 0, 0, 394, 0, 0, 0, 0,
	0, 215, 0, 0, 0, 0, 0, 0, 404, 0,
	0, 0, 0, 0, 71, 0, 0, 64, 65, 410,
	0, 275, 276, 27, 28, 69, 29, 70, 0, 0,
	0, 41, 417, 340, 420, 419, 341, 42, 43, 66,
	60, 108, 51, 63, 62, 342, 343, 0, 0, 0,
	59, 0, 0, 71, 61, 0, 64, 65, 0, 0,
	0, 45, 46, 47, 48, 0, 0, 0, 334, 335,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 0,
	67, 0, 63, 62, 27, 28, 69, 29, 70, 0,
	0, 0, 41, 392, 49, 0, 0, 50, 42, 43,
	0, 60, 0, 51, 279, 0, 0, 0, 0, 0,
	315, 59, 0, 0, 71, 61, 0, 64, 65, 0,
	0, 0, 45, 46, 47, 48, 0, 0, 0, 308,
	309, 0, 0, 0, 0, 0, 0, 0, 0, 66,
	0, 67, 0, 63, 62, 27, 28, 69, 29, 70,
	0, 0, 0, 41, 333, 340, 0, 0, 341, 42,
	43, 0, 60, 0, 51, 0, 0, 342, 343, 0,
	0, 0, 59, 0, 0, 71, 61, 0, 64, 65,
	0, 0, 0, 45, 46, 47, 48, 0, 0, 0,
	334, 335, 0, 0, 0, 0, 0, 0, 0, 0,
	66, 0, 67, 0, 63, 62, 27, 28, 69, 29,
	70, 0, 0, 0, 41, 328, 49, 274, 273, 50,
	42, 43, 0, 60, 0, 51, 0, 0, 0, 0,
	0, 0, 0, 59, 0, 0, 71, 61, 0, 64,
	65, 0, 0, 0, 45, 46, 47, 48, 0, 0,
	0, 218, 219, 0, 0, 0, 0, 0, 0, 0,
	0, 66, 0, 67, 0, 63, 62, 27, 28, 69,
	29, 70, 0, 0, 0, 41, 320, 49, 0, 0,
	50, 42, 43, 0, 60, 0, 51, 279, 0, 0,
	0, 0, 0, 315, 59, 0, 0, 71, 61, 0,
	64, 65, 0, 0, 0, 45, 46, 47, 48, 0,
	0, 0, 308, 309, 0, 0, 0, 0, 0, 0,
	0, 0, 66, 0, 67, 0, 63, 62, 27, 28,
	69, 29, 70, 0, 0, 0, 41, 306, 49, 0,
	0, 50, 42, 43, 0, 60, 0, 51, 279, 0,
	0, 0, 0, 0, 315, 59, 0, 0, 71, 61,
	0, 64, 65, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 0, 308, 309, 0, 0, 0, 0, 0,
	0, 0, 0, 66, 0, 67, 0, 63, 62, 27,
	28, 69, 29, 70, 0, 0, 0, 41, 271, 49,
	274, 273, 50, 42, 43, 0, 60, 0, 51, 0,
	0, 0, 0, 0, 0, 0, 59, 0, 0, 71,
	61, 0, 64, 65, 0, 0, 0, 45, 46, 47,
	48, 0, 0, 0, 218, 219, 0, 27, 28, 69,
	29, 70, 0, 0, 66, 41, 67, 49, 63, 62,
	50, 42, 43, 0, 60, 0, 51, 0, 0, 0,
	0, 0, 0, 0, 59, 0, 0, 71, 61, 0,
	64, 65, 0, 0, 0, 45, 46, 47, 48, 0,
	0, 0, 6, 7, 0, 0, 0, 0, 0, 0,
	0, 0, 66, 0, 67, 0, 63, 62, 0, 8,
	27, 28, 69, 29, 70, 0, 0, 0, 41, 0,
	340, 0, 0, 341, 42, 43, 0, 60, 0, 51,
	0, 0, 342, 343, 0, 0, 0, 59, 0, 0,
	71, 61, 0, 64, 65, 0, 0, 0, 45, 46,
	47, 48, 0, 0, 0, 334, 335, 0, 0, 0,
	0, 0, 0, 0, 0, 66, 0, 67, 0, 63,
	62, 27, 28, 69, 29, 70, 0, 0, 0, 41,
	403, 49, 0, 0, 50, 42, 43, 0, 60, 0,
	51, 0, 0, 0, 0, 0, 0, 0, 59, 0,
	0, 71, 61, 0, 64, 65, 0, 0, 0, 45,
	46, 47, 48, 0, 0, 0, 218, 219, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 0, 67, 0,
	63, 62, 27, 28, 69, 29, 70, 0, 0, 0,
	41, 370, 49, 0, 0, 50, 42, 43, 0, 60,
	0, 51, 0, 0, 0, 0, 0, 0, 0, 59,
	0, 0, 71, 61, 0, 64, 65, 0, 0, 0,
	45, 46, 47, 48, 0, 0, 0, 218, 219, 0,
	0, 0, 0, 0, 0, 0, 0, 66, 0, 67,
	0, 63, 62, 27, 28, 69, 29, 70, 0, 0,
	0, 41, 330, 49, 0, 0, 50, 42, 43, 0,
	60, 0, 51, 0, 0, 0, 0, 0, 0, 0,
	59, 0, 0, 71, 61, 0, 64, 65, 0, 0,
	0, 45, 46, 47, 48, 0, 0, 0, 218, 219,
	0, 27, 28, 69, 29, 70, 0, 0, 66, 41,
	67, 49, 63, 62, 50, 42, 43, 0, 60, 0,
	51, 0, 0, 0, 0, 0, 0, 0, 59, 0,
	0, 71, 61, 0, 64, 65, 0, 0, 0, 45,
	46, 47, 48, 0, 0, 0, 218, 219, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 0, 67, 305,
	63, 62, 27, 28, 69, 29, 70, 0, 0, 0,
	41, 302, 49, 0, 0, 50, 42, 43, 0, 60,
	0, 51, 0, 0, 0, 0, 0, 0, 0, 59,
	0, 0, 71, 61, 0, 64, 65, 0, 0, 0,
	45, 46, 47, 48, 0, 0, 0, 218, 219, 0,
	0, 0, 0, 0, 0, 0, 0, 66, 0, 67,
	0, 63, 62, 27, 28, 69, 29, 70, 0, 0,
	0, 41, 270, 49, 0, 0, 50, 42, 43, 0,
	60, 0, 51, 0, 0, 0, 0, 0, 0, 0,
	59, 0, 0, 71, 61, 0, 64, 65, 0, 0,
	0, 45, 46, 47, 48, 0, 0, 0, 218, 219,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 0,
	67, 0, 63, 62, 27, 28, 69, 29, 70, 0,
	0, 0, 41, 268, 49, 0, 0, 50, 42, 43,
	0, 60, 0, 51, 0, 0, 0, 0, 0, 0,
	0, 59, 0, 0, 71, 61, 0, 64, 65, 0,
	0, 0, 45, 46, 47, 48, 0, 0, 0, 218,
	219, 0, 27, 28, 69, 29, 70, 0, 0, 66,
	41, 67, 49, 63, 62, 50, 42, 43, 0, 60,
	0, 51, 0, 0, 0, 0, 0, 0, 0, 59,
	0, 0, 71, 61, 0, 64, 65, 0, 0, 0,
	45, 46, 47, 48, 0, 0, 0, 218, 219, 0,
	0, 0, 0, 0, 0, 0, 0, 66, 0, 67,
	258, 63, 62, 27, 28, 69, 29, 70, 0, 0,
	0, 41, 253, 49, 0, 0, 50, 42, 43, 0,
	60, 0, 51, 0, 0, 0, 0, 0, 0, 0,
	59, 0, 0, 71, 61, 0, 64, 65, 0, 0,
	0, 45, 46, 47, 48, 0, 0, 0, 218, 219,
	0, 27, 28, 69, 29, 70, 0, 0, 66, 41,
	67, 49, 63, 62, 50, 42, 43, 0, 60, 0,
	51, 0, 0, 0, 0, 0, 0, 0, 59, 0,
	0, 71, 61, 0, 64, 65, 0, 0, 0, 45,
	46, 47, 48, 0, 0, 0, 218, 219, 0, 27,
	28, 69, 29, 70, 240, 0, 66, 41, 67, 49,
	63, 62, 50, 42, 43, 0, 60, 0, 51, 0,
	0, 0, 0, 0, 0, 0, 59, 0, 0, 71,
	61, 0, 64, 65, 0, 0, 0, 45, 46, 47,
	48, 0, 0, 0, 0, 239, 0, 27, 28, 69,
	29, 70, 0, 0, 66, 41, 67, 49, 63, 62,
	50, 42, 43, 0, 60, 0, 51, 0, 0, 0,
	0, 0, 0, 0, 59, 0, 0, 71, 61, 0,
	64, 65, 0, 0, 0, 45, 46, 47, 48, 27,
	101, 69, 102, 91, 0, 0, 107, 27, 235, 69,
	102, 70, 66, 0, 67, 0, 63, 62, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 71,
	0, 0, 64, 65, 0, 0, 0, 71, 0, 0,
	64, 65, 0, 0, 0, 0, 289, 0, 0, 0,
	0, 0, 231, 0, 94, 0, 108, 0, 63, 62,
	0, 0, 66, 0, 67, 0, 63, 62, 27, 101,
	69, 102, 91, 0, 0, 107, 27, 101, 69, 102,
	70, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 71, 0,
	0, 64, 65, 0, 0, 0, 71, 0, 0, 64,
	65, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 94, 0, 108, 0, 63, 62, 0,
	0, 66, 0, 67, 0, 63, 62,
}
var RubyPact = []int{

	-31, 1612, -1000, -1000, -1000, 1, -1000, -1000, -1000, 844,
	-1000, -1000, -1000, 39, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 472, 115,
	22, 2, -1, -1000, -1000, -1000, -1000, -1000, 219, -23,
	-1000, 261, 204, 204, 35, 2342, 2342, 2342, 2342, 2342,
	2342, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 2461,
	2342, 2461, 17, 182, -1000, -1000, 2461, -1000, -21, 159,
	-1000, 10, -1000, -1000, -1000, 2342, 320, 2461, 2461, 2461,
	2461, 2342, 2461, 2461, 2461, 2461, 2342, 2461, 2461, 2461,
	2342, 2461, -1000, 80, 2461, 2461, 2461, 314, 128, 185,
	-1000, 2453, 147, -1000, -1000, -1000, 67, -25, -25, 2461,
	2342, 2342, 2342, 2342, 2342, 2342, 2342, 2342, 2461, 2461,
	2342, 2342, 300, 52, 122, 8, -1000, -1000, 2342, 298,
	50, 50, 50, 50, 50, 180, 2246, 55, 185, 44,
	185, -1000, -1000, 178, -1000, -1000, 15, 13, 738, -1000,
	2392, 221, 2461, 2294, -1000, -25, 50, 303, 185, 185,
	185, 185, 50, 185, 185, 185, 185, 50, 166, 185,
	185, 50, 319, 738, -1000, 89, 108, -1000, 108, -1000,
	-1000, 1159, 2198, -1000, 296, -1000, 2137, 307, 50, 50,
	50, 50, 50, 50, 50, 50, 185, 185, 50, 50,
	-1000, -1000, 287, 121, 120, -1000, 7, 280, 273, -1000,
	2089, 204, 2028, 50, -1000, 1564, -1000, -1000, -1000, -1000,
	844, 50, 238, 2461, -1000, 5, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 94, 200, 2384, -1000, 290, 50, -1000,
	-1000, 80, 10, 2342, 2461, 2461, 10, -1000, 965, -3,
	-22, 185, -1000, -1000, 1967, 16, -1000, 1906, -1000, -1000,
	1503, 122, 120, 237, 2342, -1000, -1000, 1442, -1000, -1000,
	-1000, -1000, 317, 2342, -1000, 1381, 1858, -1000, -1000, 192,
	185, 1320, 205, 2461, 1057, 6, -29, -1000, 2342, 2461,
	-1000, -1000, 50, 233, 185, -1000, 974, 185, -1000, 269,
	2342, 268, -1000, -1000, 251, -1000, -1000, 160, -1000, -1000,
	844, 50, -1000, -1000, 220, 2461, -1000, -1000, -1000, 50,
	-1000, 151, 1797, -1000, 2342, -1000, 50, 2246, -1000, 308,
	-1000, 2246, 310, -1000, -1000, -1000, 844, 50, -1000, -1000,
	2342, 2342, 163, 137, -1000, -1000, 2461, 55, 738, -1000,
	-1000, 2342, -1000, 37, 844, 50, 185, -1000, 247, -1000,
	50, -1000, -1000, -1000, -1000, 2342, 2342, 55, 1259, -1000,
	-1000, 50, 2246, 2246, -1000, 2246, 244, 25, 110, 2342,
	2342, 2342, 2342, 1736, 55, 2246, 258, -15, -12, 90,
	50, 50, -1000, 113, 2246, -1000, -1000, -1000, -1000, 50,
	50, 50, 50, -1000, 2246, -12, 2342, 2461, -1000, -1000,
	2246, 1198, 1096, 577, -12, 107, 100, -1000, 282, 2342,
	-1000, -1000, 264, -1000, -12, -1000, -12, -1000, -1000, 2342,
	-1000, 50, 1675, -1000, -12, -12, 50, 1675, 1675, 1675,
}
var RubyPgo = []int{

	0, 305, 443, 442, 45, 441, 437, 435, 761, 428,
	5, 18, 425, 424, 409, 0, 679, 407, 406, 404,
	402, 6, 401, 9, 113, 400, 394, 10, 393, 391,
	390, 389, 388, 385, 383, 382, 2, 380, 379, 378,
	377, 375, 372, 371, 364, 361, 360, 357, 803, 348,
	1, 14, 15, 7, 346, 8, 344, 33, 343, 11,
	4, 340, 3, 339, 336, 13, 12, 335, 518,
}
var RubyR1 = []int{

	0, 54, 54, 54, 54, 54, 54, 54, 54, 54,
	54, 68, 68, 48, 48, 48, 48, 48, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 21, 21, 21, 21, 21, 21, 21, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 51, 51, 51,
	51, 62, 62, 59, 59, 59, 59, 59, 65, 65,
	65, 65, 65, 64, 64, 64, 18, 18, 18, 18,
	18, 18, 28, 28, 28, 28, 60, 60, 60, 60,
	60, 60, 55, 55, 23, 23, 23, 23, 66, 66,
	66, 22, 22, 25, 27, 27, 67, 67, 13, 13,
	13, 13, 13, 13, 13, 26, 26, 26, 26, 26,
	26, 8, 8, 24, 24, 16, 16, 37, 37, 38,
	39, 40, 41, 42, 43, 44, 45, 46, 47, 2,
	5, 6, 6, 3, 3, 56, 56, 56, 56, 63,
	63, 63, 4, 4, 4, 4, 52, 61, 61, 61,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	53, 53, 53, 53, 49, 49, 49, 7, 14, 10,
	10, 10, 58, 58, 50, 50, 19, 20, 11, 33,
	57, 57, 57, 57, 57, 57, 57, 34, 34, 34,
	34, 34, 34, 34, 35, 35, 35, 35, 35, 36,
	36, 36, 36, 32, 31, 9, 30, 30, 29, 29,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 5, 1,
	4, 4, 2, 3, 4, 4, 5, 3, 5, 2,
	3, 3, 3, 3, 3, 4, 6, 3, 7, 1,
	5, 1, 3, 0, 1, 1, 4, 4, 1, 1,
	4, 4, 5, 1, 3, 3, 5, 6, 7, 8,
	5, 6, 0, 1, 3, 3, 0, 2, 2, 2,
	2, 2, 1, 3, 1, 2, 3, 2, 0, 1,
	3, 4, 6, 4, 1, 3, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
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

	-1000, -54, 50, 51, 67, -1, 50, 51, 67, -15,
	-18, -22, -25, -13, -26, -38, -39, -40, -41, -12,
	-14, -21, -19, -33, -32, -31, -30, 5, 6, 8,
	-24, -16, -8, -2, -5, -6, -3, -27, -17, -7,
	-9, 13, 19, 20, -37, 43, 44, 45, 46, 15,
	18, 24, -42, -43, -44, -45, -46, -47, -11, 32,
	22, 36, 65, 64, 38, 39, 60, 62, -67, 7,
	9, 35, 51, 50, 67, 15, 54, 40, 41, 4,
	45, 46, 47, 56, 57, 55, 18, 58, 33, 34,
	18, 9, -51, -65, 60, 42, 48, 11, -64, -15,
	-4, 6, 8, -24, -16, -8, -17, 12, 62, 9,
	42, 48, 42, 48, 42, 48, 42, 48, 40, 41,
	15, 18, 54, 6, 4, -27, 8, -27, 42, 11,
	-1, -1, -1, -1, -1, -1, -48, -62, -15, -1,
	-15, 6, 8, 65, 6, 8, -62, -59, -15, -21,
	-68, 53, 9, -49, -4, 62, -1, 6, -15, -15,
	-15, -15, -1, -15, -15, -15, -15, -1, -15, -15,
	-15, -1, -59, -15, 11, -15, -15, -11, -15, -11,
	6, 11, -48, -52, 55, -52, -48, -59, -1, -1,
	-1, -1, -1, -1, -1, -1, -15, -15, -1, -1,
	6, -55, 54, -66, 9, -23, 6, 47, 57, -55,
	-48, 40, -48, -1, 6, -48, 50, 51, 50, 51,
	-15, -1, -58, 11, 50, -68, 6, 8, 61, 11,
	61, 50, -56, -63, -15, 6, 8, -59, -1, 51,
	10, -65, -51, 42, 9, 52, 10, 11, -68, 61,
	-68, -15, -4, 14, -48, -61, 6, -48, 63, 10,
	-60, 6, 11, -66, 42, 6, 6, -60, 14, -27,
	14, 14, -53, 17, 16, -48, -48, 14, -10, 25,
	-15, -57, -29, 37, -68, -68, -68, 11, 4, 52,
	10, -4, -1, -59, -15, -4, -68, -15, -4, 57,
	42, 57, 14, 55, 11, 63, 14, -50, 50, 51,
	-15, -1, -28, -10, -20, 31, -55, -23, 10, -1,
	14, -50, -48, 14, 17, 16, -1, -48, 14, -53,
	14, -48, 8, 14, 50, 51, -15, -1, -35, -34,
	15, 18, 27, 28, 14, 16, 37, -62, -15, -21,
	63, 11, 63, -68, -15, -1, -15, 10, 57, 6,
	-1, 6, 6, 14, -10, 15, 18, -62, -60, 14,
	14, -1, -48, -48, 14, -48, 4, -1, -1, 15,
	18, 15, 18, -48, -62, -48, -1, 6, -68, 6,
	-1, -1, 14, -50, -48, 6, 50, 50, 51, -1,
	-1, -1, -1, 14, -48, -68, 4, 52, 10, 14,
	-48, -57, -57, -57, -68, -1, -15, 14, -36, 17,
	16, 14, -36, 14, -68, 11, -68, 11, 14, 17,
	16, -1, -57, 14, -68, -68, -1, -57, -57, -57,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 39,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 56, 18, 19, -2,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	31, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 13, 32, 33, 34, 35, 36, 37, 38, 0,
	0, 0, 0, 0, 159, 160, 83, 11, 0, 59,
	194, 0, 5, 6, 7, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, -2, 62, 69, 83, 0, 0, 0, 79, 88,
	89, 19, -2, 21, 22, 23, 29, 13, -2, 83,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 118, 118, 13, -2, 13, 0, 0,
	149, 150, 151, 152, 13, 0, 202, 206, 81, 0,
	11, 143, 144, 0, 141, 142, 0, 0, 81, 85,
	165, 0, 83, 0, 235, 13, 182, 63, 70, 72,
	74, 153, 154, 155, 156, 157, 158, 184, 0, 233,
	234, 186, 0, 84, 11, 81, 128, 129, 135, 136,
	147, 11, 0, 13, 177, 13, 0, 0, 130, 137,
	131, 138, 132, 139, 133, 140, 71, 73, 183, 185,
	67, 106, 0, 112, 118, 119, 114, 0, 0, 106,
	0, 0, 0, 134, 148, 0, 13, 13, 14, 15,
	16, 17, 0, 0, 210, 0, 145, 146, 161, 11,
	162, 12, 11, 11, 0, 19, -2, 0, 195, 196,
	197, 64, 65, 0, -2, 0, 57, 11, 0, 75,
	0, 94, 95, 172, 0, 0, 178, 0, 175, 61,
	0, 118, 0, 0, 0, 115, 117, 0, 121, 13,
	123, 180, 0, 0, 13, 0, 0, 198, 203, 13,
	82, 0, 0, 0, 0, 0, 0, 11, 0, 0,
	60, 66, 68, 0, 208, 58, 0, 90, 91, 0,
	0, 0, 173, 176, 0, 174, 96, 0, 107, 108,
	39, 110, 111, 204, 103, 0, 106, 120, 113, 116,
	100, 0, 0, 181, 0, 13, 13, 193, 187, 0,
	189, 199, 13, 209, 211, 212, 39, 214, 215, 216,
	0, 0, 218, 221, 236, 13, 0, 13, 86, 87,
	163, 0, 164, 0, 39, 11, 169, 77, 0, 92,
	76, 80, 179, 97, 205, 0, 0, 207, 0, 101,
	122, 13, 191, 192, 188, 200, 0, 13, 0, 0,
	0, 0, 0, 0, 13, 11, 0, 0, 166, 0,
	104, 105, 98, 0, 190, 13, 210, -2, -2, 219,
	220, 222, 223, 237, 11, 238, 0, 0, 78, 99,
	201, 0, 0, 0, 239, 11, 11, 224, 0, 0,
	210, 226, 0, 228, 167, 11, 170, 11, 225, 0,
	210, 210, 217, 227, 168, 171, 210, 217, 217, 217,
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
	62, 63, 64, 65, 66, 67,
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
		//line parser.y:196
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:198
		{
		}
	case 3:
		//line parser.y:200
		{
		}
	case 4:
		//line parser.y:202
		{
		}
	case 5:
		//line parser.y:204
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:206
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:208
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:214
		{
		}
	case 11:
		//line parser.y:216
		{
		}
	case 12:
		//line parser.y:217
		{
		}
	case 13:
		//line parser.y:220
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 14:
		//line parser.y:222
		{
		}
	case 15:
		//line parser.y:224
		{
		}
	case 16:
		//line parser.y:226
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 17:
		//line parser.y:228
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 57:
		//line parser.y:238
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 58:
		//line parser.y:245
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-2].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 59:
		//line parser.y:252
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 60:
		//line parser.y:256
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 61:
		//line parser.y:263
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 62:
		//line parser.y:270
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 63:
		//line parser.y:277
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 64:
		//line parser.y:284
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 65:
		//line parser.y:292
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 66:
		//line parser.y:300
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 67:
		//line parser.y:308
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 68:
		//line parser.y:316
		{
			methodName := RubyS[Rubypt-2].genericValue.(ast.BareReference).Name + "="
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: methodName},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 69:
		//line parser.y:327
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 70:
		//line parser.y:334
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 71:
		//line parser.y:342
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 72:
		//line parser.y:350
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 73:
		//line parser.y:358
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 74:
		//line parser.y:366
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 75:
		//line parser.y:376
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 76:
		//line parser.y:386
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 77:
		//line parser.y:395
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 78:
		//line parser.y:397
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 79:
		//line parser.y:399
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 80:
		//line parser.y:401
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 81:
		//line parser.y:404
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 82:
		//line parser.y:406
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 83:
		//line parser.y:408
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 84:
		//line parser.y:410
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:412
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 86:
		//line parser.y:414
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 87:
		//line parser.y:416
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 88:
		//line parser.y:420
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 89:
		//line parser.y:422
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 90:
		//line parser.y:424
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 91:
		//line parser.y:426
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-1].genericValue)
		}
	case 92:
		//line parser.y:428
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 93:
		//line parser.y:431
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 94:
		//line parser.y:433
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 95:
		//line parser.y:435
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 96:
		//line parser.y:441
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 97:
		//line parser.y:449
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 98:
		//line parser.y:458
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 99:
		//line parser.y:467
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 100:
		//line parser.y:477
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 101:
		//line parser.y:485
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 102:
		//line parser.y:495
		{
		}
	case 103:
		//line parser.y:497
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 104:
		//line parser.y:499
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{RubyS[Rubypt-2].genericValue}}
		}
	case 105:
		//line parser.y:501
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{ast.Break{}},
			}
		}
	case 106:
		//line parser.y:509
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 107:
		//line parser.y:511
		{
		}
	case 108:
		//line parser.y:513
		{
		}
	case 109:
		//line parser.y:515
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 110:
		//line parser.y:517
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 111:
		//line parser.y:519
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 112:
		//line parser.y:524
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 113:
		//line parser.y:526
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 114:
		//line parser.y:529
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 115:
		//line parser.y:531
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 116:
		//line parser.y:533
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 117:
		//line parser.y:535
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 118:
		//line parser.y:537
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 119:
		//line parser.y:539
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 120:
		//line parser.y:543
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 121:
		//line parser.y:548
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 122:
		//line parser.y:555
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 123:
		//line parser.y:565
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 124:
		//line parser.y:574
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 125:
		//line parser.y:580
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 126:
		//line parser.y:588
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 127:
		//line parser.y:592
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
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
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 130:
		//line parser.y:608
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 131:
		//line parser.y:615
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 132:
		//line parser.y:622
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 133:
		//line parser.y:629
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 134:
		//line parser.y:636
		{
			RubyVAL.genericValue = ast.Assignment{
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
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 137:
		//line parser.y:655
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 138:
		//line parser.y:662
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 139:
		//line parser.y:669
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 140:
		//line parser.y:676
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 141:
		//line parser.y:684
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 142:
		//line parser.y:686
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 143:
		//line parser.y:689
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 144:
		//line parser.y:691
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 145:
		//line parser.y:694
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 146:
		//line parser.y:696
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 147:
		//line parser.y:699
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 148:
		//line parser.y:701
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 149:
		//line parser.y:703
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 150:
		//line parser.y:704
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 151:
		//line parser.y:705
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 152:
		//line parser.y:706
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 153:
		//line parser.y:709
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 154:
		//line parser.y:718
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 155:
		//line parser.y:727
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 156:
		//line parser.y:736
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 157:
		//line parser.y:745
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 158:
		//line parser.y:754
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 159:
		//line parser.y:762
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 160:
		//line parser.y:763
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 161:
		//line parser.y:765
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 162:
		//line parser.y:767
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 163:
		//line parser.y:770
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 164:
		//line parser.y:778
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 165:
		//line parser.y:786
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 166:
		//line parser.y:788
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 167:
		//line parser.y:795
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 168:
		//line parser.y:802
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 169:
		//line parser.y:810
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 170:
		//line parser.y:817
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 171:
		//line parser.y:824
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 172:
		//line parser.y:832
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 173:
		//line parser.y:834
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 174:
		//line parser.y:841
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 175:
		//line parser.y:848
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 176:
		//line parser.y:851
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 177:
		//line parser.y:853
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 178:
		//line parser.y:855
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 179:
		//line parser.y:857
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 180:
		//line parser.y:860
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 181:
		//line parser.y:867
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 182:
		//line parser.y:875
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 183:
		//line parser.y:882
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 184:
		//line parser.y:889
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 185:
		//line parser.y:896
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 186:
		//line parser.y:903
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 187:
		//line parser.y:910
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 188:
		//line parser.y:917
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 189:
		//line parser.y:925
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 190:
		//line parser.y:934
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 191:
		//line parser.y:941
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 192:
		//line parser.y:948
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 193:
		//line parser.y:955
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 194:
		//line parser.y:962
		{
		}
	case 195:
		//line parser.y:963
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 196:
		//line parser.y:964
		{
		}
	case 197:
		//line parser.y:967
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 198:
		//line parser.y:970
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 199:
		//line parser.y:978
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 200:
		//line parser.y:980
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 201:
		//line parser.y:989
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
	case 202:
		//line parser.y:1004
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 203:
		//line parser.y:1006
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 204:
		//line parser.y:1009
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 205:
		//line parser.y:1011
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 206:
		//line parser.y:1014
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 207:
		//line parser.y:1023
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 208:
		//line parser.y:1032
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 209:
		//line parser.y:1041
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 210:
		//line parser.y:1044
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 211:
		//line parser.y:1046
		{
		}
	case 212:
		//line parser.y:1048
		{
		}
	case 213:
		//line parser.y:1050
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 214:
		//line parser.y:1052
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 215:
		//line parser.y:1054
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 216:
		//line parser.y:1056
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 217:
		//line parser.y:1058
		{
		}
	case 218:
		//line parser.y:1060
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 219:
		//line parser.y:1062
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 220:
		//line parser.y:1064
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 221:
		//line parser.y:1066
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 222:
		//line parser.y:1068
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 223:
		//line parser.y:1070
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 224:
		//line parser.y:1073
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 225:
		//line parser.y:1080
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 226:
		//line parser.y:1088
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 227:
		//line parser.y:1095
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 228:
		//line parser.y:1103
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 229:
		//line parser.y:1111
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 230:
		//line parser.y:1118
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 231:
		//line parser.y:1125
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 232:
		//line parser.y:1132
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 233:
		//line parser.y:1140
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 234:
		//line parser.y:1143
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 235:
		//line parser.y:1145
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 236:
		//line parser.y:1148
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 237:
		//line parser.y:1150
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 238:
		//line parser.y:1153
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 239:
		//line parser.y:1155
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	}
	goto Rubystack /* stack new state and value */
}
