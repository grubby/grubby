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

//line parser.y:1085

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
	-2, 185,
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
	-1, 120,
	52, 123,
	-2, 121,
	-1, 224,
	52, 124,
	-2, 122,
	-1, 232,
	10, 82,
	11, 82,
	-2, 185,
	-1, 384,
	27, 201,
	28, 201,
	-2, 13,
	-1, 385,
	27, 201,
	28, 201,
	-2, 13,
}

const RubyNprod = 231
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 2406

var RubyAct = []int{

	9, 294, 349, 140, 260, 408, 36, 141, 189, 193,
	191, 92, 91, 177, 57, 105, 74, 2, 3, 329,
	269, 176, 219, 135, 219, 136, 20, 217, 97, 292,
	116, 289, 145, 394, 4, 330, 406, 78, 219, 288,
	252, 271, 112, 113, 74, 98, 199, 211, 119, 121,
	72, 71, 194, 219, 111, 192, 194, 219, 89, 132,
	194, 134, 131, 192, 149, 110, 142, 73, 109, 380,
	328, 219, 291, 76, 77, 78, 218, 152, 153, 154,
	155, 137, 157, 158, 159, 160, 75, 162, 163, 164,
	86, 167, 143, 195, 169, 170, 216, 195, 166, 190,
	97, 195, 196, 370, 219, 211, 196, 74, 167, 171,
	196, 76, 77, 184, 185, 179, 148, 143, 373, 175,
	143, 74, 78, 74, 75, 114, 74, 197, 115, 417,
	123, 208, 107, 168, 143, 275, 26, 99, 68, 100,
	69, 384, 385, 105, 415, 222, 219, 167, 74, 78,
	112, 113, 97, 219, 225, 204, 205, 383, 76, 77,
	212, 122, 250, 229, 230, 108, 70, 374, 173, 63,
	64, 75, 144, 143, 239, 208, 395, 265, 267, 208,
	219, 209, 322, 5, 323, 76, 77, 287, 267, 214,
	65, 215, 106, 107, 62, 61, 233, 376, 75, 208,
	377, 208, 276, 251, 208, 324, 257, 423, 146, 420,
	419, 393, 268, 266, 418, 382, 420, 419, 364, 240,
	362, 365, 74, 363, 97, 310, 124, 125, 126, 127,
	128, 129, 120, 167, 282, 335, 235, 285, 76, 77,
	281, 133, 357, 208, 303, 302, 208, 301, 224, 303,
	302, 75, 297, 250, 278, 217, 150, 299, 295, 143,
	296, 138, 156, 139, 208, 208, 372, 161, 307, 340,
	314, 165, 132, 326, 339, 325, 279, 332, 334, 118,
	283, 117, 286, 247, 217, 326, 234, 235, 337, 254,
	180, 181, 182, 183, 253, 344, 186, 187, 249, 327,
	344, 208, 244, 202, 201, 188, 208, 213, 172, 151,
	208, 327, 359, 67, 96, 221, 243, 210, 220, 1,
	347, 147, 56, 55, 54, 132, 53, 52, 367, 226,
	26, 99, 68, 100, 69, 51, 17, 16, 15, 14,
	43, 236, 316, 317, 22, 23, 238, 24, 25, 270,
	375, 132, 346, 379, 378, 12, 208, 208, 11, 208,
	70, 348, 248, 63, 64, 21, 10, 208, 19, 208,
	255, 13, 18, 39, 219, 38, 34, 33, 35, 32,
	344, 336, 208, 0, 65, 0, 66, 0, 62, 61,
	272, 0, 208, 273, 274, 405, 0, 0, 0, 0,
	208, 314, 314, 314, 400, 401, 402, 412, 284, 0,
	0, 0, 0, 280, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 314, 0, 0, 0, 0, 314, 314,
	314, 422, 0, 0, 298, 0, 0, 78, 0, 0,
	0, 427, 428, 304, 0, 0, 0, 429, 331, 0,
	0, 315, 0, 0, 0, 0, 0, 0, 333, 0,
	0, 0, 0, 0, 0, 0, 0, 341, 351, 0,
	338, 0, 352, 76, 77, 0, 345, 0, 79, 80,
	81, 345, 0, 0, 354, 0, 75, 84, 82, 83,
	86, 0, 0, 237, 0, 26, 27, 68, 28, 69,
	360, 361, 0, 40, 413, 318, 371, 0, 319, 41,
	42, 369, 59, 0, 50, 0, 0, 320, 321, 0,
	0, 0, 58, 0, 0, 70, 60, 0, 63, 64,
	0, 0, 0, 44, 45, 46, 47, 29, 0, 312,
	313, 392, 0, 0, 386, 387, 388, 389, 0, 65,
	0, 66, 398, 62, 61, 0, 0, 0, 396, 397,
	0, 345, 0, 0, 403, 101, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 404, 0, 414, 416, 0,
	37, 0, 315, 315, 315, 0, 0, 0, 424, 0,
	425, 421, 0, 0, 0, 0, 101, 0, 101, 0,
	0, 426, 0, 101, 315, 0, 0, 0, 104, 315,
	315, 315, 0, 0, 101, 101, 101, 101, 0, 101,
	101, 101, 101, 0, 101, 101, 101, 0, 101, 0,
	0, 101, 101, 0, 0, 0, 0, 101, 0, 104,
	0, 104, 0, 0, 0, 101, 104, 0, 0, 78,
	101, 101, 0, 0, 0, 0, 0, 104, 104, 104,
	104, 0, 104, 104, 104, 104, 0, 104, 104, 104,
	0, 104, 0, 0, 104, 104, 0, 0, 0, 0,
	104, 0, 101, 0, 101, 76, 77, 0, 104, 101,
	79, 80, 81, 104, 104, 0, 0, 0, 75, 84,
	82, 83, 86, 0, 26, 27, 68, 28, 69, 0,
	0, 101, 40, 0, 48, 0, 0, 49, 41, 42,
	0, 59, 0, 50, 0, 104, 0, 104, 0, 0,
	0, 58, 104, 0, 70, 60, 0, 63, 64, 0,
	0, 0, 44, 45, 46, 47, 0, 0, 0, 101,
	0, 0, 0, 0, 104, 0, 0, 0, 65, 0,
	66, 101, 62, 61, 0, 26, 99, 68, 100, 90,
	101, 101, 105, 0, 101, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	30, 0, 104, 0, 0, 70, 0, 0, 63, 64,
	0, 0, 0, 0, 104, 0, 0, 0, 0, 101,
	101, 277, 0, 104, 104, 101, 0, 104, 102, 93,
	0, 106, 101, 62, 61, 0, 0, 0, 0, 0,
	0, 0, 0, 31, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 102,
	0, 102, 104, 104, 0, 0, 102, 0, 104, 0,
	0, 103, 101, 0, 0, 104, 0, 102, 102, 102,
	102, 0, 102, 102, 102, 102, 0, 102, 102, 102,
	0, 102, 0, 0, 102, 102, 0, 0, 101, 0,
	102, 0, 103, 178, 103, 0, 0, 0, 102, 103,
	0, 0, 0, 102, 102, 104, 0, 0, 0, 0,
	103, 103, 103, 103, 0, 103, 103, 103, 103, 0,
	103, 103, 103, 0, 103, 0, 0, 103, 103, 0,
	0, 104, 101, 103, 0, 102, 0, 102, 0, 0,
	78, 103, 102, 0, 130, 0, 103, 103, 0, 0,
	0, 0, 0, 0, 85, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 102, 0, 0, 0, 0, 87,
	88, 0, 0, 0, 0, 104, 76, 77, 103, 0,
	103, 79, 80, 81, 0, 103, 0, 0, 0, 75,
	84, 82, 83, 86, 0, 0, 0, 0, 0, 174,
	0, 0, 102, 0, 0, 0, 0, 103, 0, 0,
	0, 0, 0, 198, 102, 200, 0, 0, 0, 0,
	0, 0, 203, 102, 102, 0, 0, 102, 0, 0,
	0, 26, 99, 68, 100, 69, 0, 0, 105, 0,
	0, 0, 0, 0, 0, 103, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 103, 0, 0,
	0, 70, 102, 102, 63, 64, 103, 103, 102, 242,
	103, 245, 0, 0, 0, 102, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 106, 0, 62,
	61, 0, 0, 0, 0, 0, 0, 0, 263, 264,
	0, 0, 0, 0, 0, 103, 103, 0, 0, 0,
	0, 103, 0, 0, 0, 102, 0, 0, 103, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 26, 27, 68, 28, 69,
	0, 102, 0, 40, 306, 48, 262, 261, 49, 41,
	42, 300, 59, 0, 50, 0, 305, 0, 103, 0,
	0, 309, 58, 0, 0, 70, 60, 0, 63, 64,
	0, 0, 0, 44, 45, 46, 47, 0, 0, 206,
	207, 0, 0, 0, 103, 102, 0, 0, 0, 65,
	0, 66, 0, 62, 61, 0, 0, 355, 356, 26,
	27, 68, 28, 69, 358, 0, 0, 40, 411, 318,
	410, 409, 319, 41, 42, 0, 59, 366, 50, 368,
	0, 320, 321, 0, 0, 0, 58, 0, 103, 70,
	60, 0, 63, 64, 0, 0, 0, 44, 45, 46,
	47, 0, 0, 312, 313, 0, 0, 0, 381, 0,
	0, 0, 0, 65, 203, 66, 0, 62, 61, 0,
	0, 391, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 399, 0, 263, 264,
	26, 27, 68, 28, 69, 0, 0, 0, 40, 407,
	318, 410, 409, 319, 41, 42, 0, 59, 0, 50,
	0, 0, 320, 321, 0, 0, 0, 58, 0, 0,
	70, 60, 0, 63, 64, 0, 0, 0, 44, 45,
	46, 47, 0, 0, 312, 313, 0, 0, 0, 0,
	0, 0, 0, 0, 65, 0, 66, 0, 62, 61,
	26, 27, 68, 28, 69, 0, 0, 0, 40, 311,
	318, 0, 0, 319, 41, 42, 0, 59, 0, 50,
	0, 0, 320, 321, 0, 0, 0, 58, 0, 0,
	70, 60, 0, 63, 64, 0, 0, 0, 44, 45,
	46, 47, 0, 0, 312, 313, 0, 0, 0, 0,
	0, 0, 0, 0, 65, 0, 66, 0, 62, 61,
	26, 27, 68, 28, 69, 0, 0, 0, 40, 259,
	48, 262, 261, 49, 41, 42, 0, 59, 0, 50,
	0, 0, 0, 0, 0, 0, 0, 58, 0, 0,
	70, 60, 0, 63, 64, 0, 0, 0, 44, 45,
	46, 47, 0, 0, 206, 207, 0, 26, 27, 68,
	28, 69, 0, 0, 65, 40, 66, 48, 62, 61,
	49, 41, 42, 0, 59, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 58, 0, 0, 70, 60, 0,
	63, 64, 0, 0, 0, 44, 45, 46, 47, 0,
	0, 6, 7, 0, 0, 0, 0, 0, 0, 0,
	0, 65, 0, 66, 0, 62, 61, 0, 8, 26,
	27, 68, 28, 69, 0, 0, 0, 40, 0, 318,
	0, 0, 319, 41, 42, 0, 59, 0, 50, 0,
	0, 320, 321, 0, 0, 0, 58, 0, 0, 70,
	60, 0, 63, 64, 0, 0, 0, 44, 45, 46,
	47, 0, 0, 312, 313, 0, 26, 27, 68, 28,
	69, 0, 0, 65, 40, 66, 48, 62, 61, 49,
	41, 42, 0, 59, 0, 50, 267, 0, 0, 0,
	0, 0, 350, 58, 0, 0, 70, 60, 0, 63,
	64, 0, 0, 0, 44, 45, 46, 47, 0, 0,
	342, 343, 0, 0, 0, 0, 0, 0, 0, 0,
	65, 0, 66, 0, 62, 61, 26, 27, 68, 28,
	69, 0, 0, 0, 40, 390, 48, 0, 0, 49,
	41, 42, 0, 59, 0, 50, 0, 0, 0, 0,
	0, 0, 0, 58, 0, 0, 70, 60, 0, 63,
	64, 0, 0, 0, 44, 45, 46, 47, 0, 0,
	206, 207, 0, 0, 0, 0, 0, 0, 0, 0,
	65, 0, 66, 0, 62, 61, 26, 27, 68, 28,
	69, 0, 0, 0, 40, 353, 48, 0, 0, 49,
	41, 42, 0, 59, 0, 50, 0, 0, 0, 0,
	0, 0, 0, 58, 0, 0, 70, 60, 0, 63,
	64, 0, 0, 0, 44, 45, 46, 47, 0, 0,
	206, 207, 0, 0, 0, 0, 0, 0, 0, 0,
	65, 0, 66, 0, 62, 61, 26, 27, 68, 28,
	69, 0, 0, 0, 40, 308, 48, 0, 0, 49,
	41, 42, 0, 59, 0, 50, 0, 0, 0, 0,
	0, 0, 0, 58, 0, 0, 70, 60, 0, 63,
	64, 0, 0, 0, 44, 45, 46, 47, 0, 0,
	206, 207, 0, 26, 27, 68, 28, 69, 0, 0,
	65, 40, 66, 48, 62, 61, 49, 41, 42, 0,
	59, 0, 50, 0, 0, 0, 0, 0, 0, 0,
	58, 0, 0, 70, 60, 0, 63, 64, 0, 0,
	0, 44, 45, 46, 47, 0, 0, 206, 207, 0,
	0, 0, 0, 0, 0, 0, 0, 65, 0, 66,
	293, 62, 61, 26, 27, 68, 28, 69, 0, 0,
	0, 40, 290, 48, 0, 0, 49, 41, 42, 0,
	59, 0, 50, 0, 0, 0, 0, 0, 0, 0,
	58, 0, 0, 70, 60, 0, 63, 64, 0, 0,
	0, 44, 45, 46, 47, 0, 0, 206, 207, 0,
	0, 0, 0, 0, 0, 0, 0, 65, 0, 66,
	0, 62, 61, 26, 27, 68, 28, 69, 0, 0,
	0, 40, 258, 48, 0, 0, 49, 41, 42, 0,
	59, 0, 50, 0, 0, 0, 0, 0, 0, 0,
	58, 0, 0, 70, 60, 0, 63, 64, 0, 0,
	0, 44, 45, 46, 47, 0, 0, 206, 207, 0,
	0, 0, 0, 0, 0, 0, 0, 65, 0, 66,
	0, 62, 61, 26, 27, 68, 28, 69, 0, 0,
	0, 40, 256, 48, 0, 0, 49, 41, 42, 0,
	59, 0, 50, 0, 0, 0, 0, 0, 0, 0,
	58, 0, 0, 70, 60, 0, 63, 64, 0, 0,
	0, 44, 45, 46, 47, 0, 0, 206, 207, 0,
	26, 27, 68, 28, 69, 0, 0, 65, 40, 66,
	48, 62, 61, 49, 41, 42, 0, 59, 0, 50,
	0, 0, 0, 0, 0, 0, 0, 58, 0, 0,
	70, 60, 0, 63, 64, 0, 0, 0, 44, 45,
	46, 47, 0, 0, 206, 207, 0, 0, 0, 0,
	0, 0, 0, 0, 65, 0, 66, 246, 62, 61,
	26, 27, 68, 28, 69, 0, 0, 0, 40, 241,
	48, 0, 0, 49, 41, 42, 0, 59, 0, 50,
	0, 0, 0, 0, 0, 0, 0, 58, 0, 0,
	70, 60, 0, 63, 64, 0, 0, 0, 44, 45,
	46, 47, 0, 0, 206, 207, 0, 26, 27, 68,
	28, 69, 0, 0, 65, 40, 66, 48, 62, 61,
	49, 41, 42, 0, 59, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 58, 0, 0, 70, 60, 0,
	63, 64, 0, 0, 0, 44, 45, 46, 47, 0,
	0, 206, 207, 0, 26, 27, 68, 28, 69, 228,
	0, 65, 40, 66, 48, 62, 61, 49, 41, 42,
	0, 59, 0, 50, 0, 0, 0, 0, 0, 0,
	0, 58, 0, 0, 70, 60, 0, 63, 64, 0,
	0, 0, 44, 45, 46, 47, 0, 0, 0, 227,
	26, 99, 68, 100, 90, 0, 95, 105, 65, 0,
	66, 0, 62, 61, 26, 99, 68, 100, 232, 0,
	0, 105, 0, 0, 0, 0, 0, 0, 0, 0,
	70, 0, 0, 63, 64, 0, 0, 94, 0, 0,
	0, 0, 0, 0, 70, 0, 0, 63, 64, 0,
	0, 231, 0, 0, 93, 0, 106, 0, 62, 61,
	26, 99, 68, 100, 69, 0, 0, 0, 65, 0,
	106, 0, 62, 61, 0, 0, 26, 223, 68, 100,
	69, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	70, 0, 0, 63, 64, 26, 99, 68, 100, 90,
	0, 0, 105, 0, 219, 0, 70, 0, 0, 63,
	64, 0, 0, 0, 65, 0, 66, 0, 62, 61,
	219, 0, 0, 0, 0, 70, 0, 0, 63, 64,
	65, 0, 66, 0, 62, 61, 26, 99, 68, 100,
	69, 0, 0, 0, 0, 0, 0, 0, 0, 93,
	0, 106, 0, 62, 61, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 70, 0, 0, 63,
	64, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	65, 0, 66, 0, 62, 61,
}
var RubyPact = []int{

	-32, 1442, -1000, -1000, -1000, 1, -1000, -1000, -1000, 936,
	-1000, -1000, -1000, 40, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 2205, 123, 26,
	23, 12, -1000, -1000, -1000, -1000, -1000, 110, -23, -1000,
	275, 224, 224, 119, 699, 699, 699, 699, 699, 699,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 2341, 699,
	2341, 17, 255, -1000, -1000, 2341, -1000, -20, 199, -1000,
	3, -1000, -1000, -1000, 699, 303, 2341, 2341, 2341, 2341,
	699, 2341, 2341, 2341, 2341, 699, 2341, 2341, 2341, 699,
	2341, -1000, 122, 2341, 2341, 302, 157, 71, -1000, 2300,
	184, -1000, -1000, -1000, 2, -33, -33, 2341, 699, 699,
	699, 699, 2341, 2341, 699, 699, 299, 46, 54, 6,
	-1000, -1000, 699, 297, 29, 29, 29, 29, 29, 106,
	2112, 94, 71, 111, 71, -1000, -1000, 183, -1000, -1000,
	36, 16, 645, -1000, 2281, 240, 2341, 2159, -1000, -33,
	29, 2219, 71, 71, 71, 71, 29, 71, 71, 71,
	71, 29, 145, 71, 71, 29, 276, 645, -1000, 433,
	33, -1000, -1000, 1026, 2065, -1000, 296, -1000, 2005, 273,
	29, 29, 29, 29, 71, 71, 29, 29, -1000, -1000,
	292, 151, 50, -1000, -2, 288, 283, -1000, 1958, 224,
	1898, 29, -1000, 1395, -1000, -1000, -1000, -1000, 936, 29,
	163, 2341, -1000, 4, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 124, 198, 760, -1000, 244, 29, -1000, -1000, 122,
	3, 699, 2341, 2341, 3, -1000, 131, -3, -25, 71,
	-1000, -1000, 1838, 18, -1000, 1778, -1000, -1000, -11, 54,
	50, 242, 699, -1000, -1000, -11, -1000, -1000, -1000, -1000,
	233, 699, -1000, 1130, 1731, -1000, -1000, 217, 71, 1335,
	168, 2341, 2265, 8, -27, -1000, 699, 2341, -1000, -1000,
	29, 225, 71, -1000, 325, 71, -1000, 282, 699, 268,
	-1000, -1000, 263, -1000, 1551, -1000, -1000, -1000, 29, 1551,
	1671, -1000, 699, -1000, 29, 2112, -1000, 228, -1000, 2112,
	308, -1000, -1000, -1000, 936, 29, -1000, -1000, 699, 699,
	205, 203, -1000, -1000, 2341, 94, 645, -1000, -1000, 699,
	-1000, 97, 936, 29, 71, -1000, 260, -1000, 29, -1000,
	-1000, 104, -1000, -1000, 936, 29, -1000, 153, 182, -1000,
	2341, -11, 55, -1000, 29, 2112, 2112, -1000, 2112, 209,
	108, 92, 699, 699, 699, 699, 1611, 94, 2112, 207,
	-18, -11, 166, -1000, -1000, -1000, 699, 699, 94, 1551,
	-1000, 2112, -1000, -1000, -1000, -1000, 29, 29, 29, 29,
	-1000, 2112, -11, 699, 2341, -1000, 29, 29, 22, 2112,
	1275, 1194, 490, -11, 133, 118, -1000, -1000, 200, 699,
	-1000, -1000, 193, -1000, -11, -1000, -11, -1000, -1000, 699,
	-1000, 29, 1504, -1000, -11, -11, 29, 1504, 1504, 1504,
}
var RubyPgo = []int{

	0, 181, 379, 378, 45, 377, 376, 375, 833, 373,
	2, 14, 372, 371, 368, 0, 790, 580, 366, 365,
	361, 26, 358, 9, 537, 355, 6, 352, 349, 348,
	347, 345, 344, 343, 342, 5, 340, 339, 338, 337,
	336, 335, 327, 326, 324, 323, 322, 893, 321, 320,
	12, 13, 4, 319, 8, 318, 20, 317, 7, 1,
	316, 3, 315, 314, 11, 10, 313, 172,
}
var RubyR1 = []int{

	0, 53, 53, 53, 53, 53, 53, 53, 53, 53,
	53, 67, 67, 47, 47, 47, 47, 47, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 21, 21, 21, 21, 21, 21, 21, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 50, 50, 50, 50,
	61, 61, 58, 58, 58, 58, 58, 64, 64, 64,
	64, 64, 63, 63, 63, 18, 18, 18, 27, 27,
	27, 27, 59, 59, 59, 59, 59, 59, 59, 54,
	54, 23, 23, 23, 23, 65, 65, 65, 22, 22,
	25, 26, 26, 66, 66, 13, 13, 13, 13, 13,
	13, 13, 8, 8, 24, 24, 16, 16, 36, 36,
	37, 38, 39, 40, 41, 42, 43, 44, 45, 46,
	2, 5, 6, 6, 3, 3, 55, 55, 55, 55,
	62, 62, 62, 4, 4, 4, 4, 51, 60, 60,
	60, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 52, 52, 52, 52, 48, 48, 48, 7, 14,
	10, 10, 10, 57, 57, 49, 49, 19, 20, 11,
	32, 56, 56, 56, 56, 56, 56, 56, 33, 33,
	33, 33, 33, 33, 33, 34, 34, 34, 34, 34,
	35, 35, 35, 35, 31, 30, 9, 29, 29, 28,
	28,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 4, 5, 1, 4,
	4, 2, 3, 4, 4, 5, 3, 5, 2, 3,
	3, 3, 3, 3, 4, 6, 3, 7, 1, 5,
	1, 3, 0, 1, 1, 4, 4, 1, 1, 4,
	4, 5, 1, 3, 3, 7, 9, 7, 0, 1,
	3, 3, 0, 2, 2, 2, 2, 2, 3, 1,
	3, 1, 2, 3, 2, 0, 1, 3, 4, 6,
	4, 1, 3, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 2, 2, 2, 3, 3, 3, 3,
	2, 2, 2, 2, 3, 3, 3, 3, 3, 3,
	1, 1, 3, 3, 5, 5, 0, 4, 7, 8,
	3, 7, 8, 3, 4, 4, 3, 3, 0, 1,
	3, 4, 5, 3, 3, 3, 3, 3, 5, 6,
	5, 4, 3, 3, 2, 0, 2, 2, 3, 4,
	2, 3, 5, 0, 2, 1, 2, 2, 2, 5,
	5, 0, 2, 2, 2, 2, 2, 2, 0, 1,
	3, 3, 1, 3, 3, 5, 6, 5, 6, 5,
	4, 3, 3, 2, 3, 3, 2, 5, 7, 4,
	5,
}
var RubyChk = []int{

	-1000, -53, 49, 50, 66, -1, 49, 50, 66, -15,
	-18, -22, -25, -13, -37, -38, -39, -40, -12, -14,
	-21, -19, -32, -31, -30, -29, 5, 6, 8, -24,
	-16, -8, -2, -5, -6, -3, -26, -17, -7, -9,
	13, 19, 20, -36, 43, 44, 45, 46, 15, 18,
	24, -41, -42, -43, -44, -45, -46, -11, 32, 22,
	36, 64, 63, 38, 39, 59, 61, -66, 7, 9,
	35, 50, 49, 66, 15, 53, 40, 41, 4, 45,
	46, 47, 55, 56, 54, 18, 57, 33, 34, 18,
	9, -50, -64, 59, 42, 11, -63, -15, -4, 6,
	8, -24, -16, -8, -17, 12, 61, 9, 42, 42,
	42, 42, 40, 41, 15, 18, 53, 6, 4, -26,
	8, -26, 42, 11, -1, -1, -1, -1, -1, -1,
	-47, -61, -15, -1, -15, 6, 8, 64, 6, 8,
	-61, -58, -15, -21, -67, 52, 9, -48, -4, 61,
	-1, 6, -15, -15, -15, -15, -1, -15, -15, -15,
	-15, -1, -15, -15, -15, -1, -58, -15, 11, -15,
	-15, -11, 6, 11, -47, -51, 54, -51, -47, -58,
	-1, -1, -1, -1, -15, -15, -1, -1, 6, -54,
	53, -65, 9, -23, 6, 47, 56, -54, -47, 40,
	-47, -1, 6, -47, 49, 50, 49, 50, -15, -1,
	-57, 11, 49, -67, 6, 8, 60, 11, 60, 49,
	-55, -62, -15, 6, 8, -58, -1, 50, 10, -64,
	-50, 42, 9, 51, 10, 11, -67, 60, -67, -15,
	-4, 14, -47, -60, 6, -47, 62, 10, -67, 6,
	11, -65, 42, 6, 6, -67, 14, -26, 14, 14,
	-52, 17, 16, -47, -47, 14, -10, 25, -15, -56,
	-28, 37, -67, -67, -67, 11, 4, 51, 10, -4,
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
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	13, 32, 33, 34, 35, 36, 37, 38, 0, 0,
	0, 0, 0, 150, 151, 82, 11, 0, 58, 185,
	0, 5, 6, 7, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	-2, 61, 68, 82, 0, 0, 78, 87, 88, 19,
	-2, 21, 22, 23, 29, 13, -2, 82, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 115, 115, 13,
	-2, 13, 0, 0, 140, 141, 142, 143, 13, 0,
	193, 197, 80, 0, 11, 134, 135, 0, 132, 133,
	0, 0, 80, 84, 156, 0, 82, 0, 226, 13,
	173, 62, 69, 71, 73, 144, 145, 146, 147, 148,
	149, 175, 0, 224, 225, 177, 0, 83, 11, 80,
	125, 126, 138, 11, 0, 13, 168, 13, 0, 0,
	127, 128, 129, 130, 70, 72, 174, 176, 66, 11,
	0, 109, 115, 116, 111, 0, 0, 11, 0, 0,
	0, 131, 139, 0, 13, 13, 14, 15, 16, 17,
	0, 0, 201, 0, 136, 137, 152, 11, 153, 12,
	11, 11, 0, 19, -2, 0, 186, 187, 188, 63,
	64, 0, -2, 0, 56, 11, 0, 74, 0, 93,
	94, 163, 0, 0, 169, 0, 166, 60, 102, 115,
	0, 0, 0, 112, 114, 102, 118, 13, 120, 171,
	0, 0, 13, 0, 0, 189, 194, 13, 81, 0,
	0, 0, 0, 0, 0, 11, 0, 0, 59, 65,
	67, 0, 199, 57, 0, 89, 90, 0, 0, 0,
	164, 167, 0, 165, 11, 11, 117, 110, 113, 11,
	0, 172, 0, 13, 13, 184, 178, 0, 180, 190,
	13, 200, 202, 203, 39, 205, 206, 207, 0, 0,
	209, 212, 227, 13, 0, 13, 85, 86, 154, 0,
	155, 0, 39, 11, 160, 76, 0, 91, 75, 79,
	170, 0, 103, 104, 39, 106, 107, 0, 99, 195,
	0, 102, 0, 119, 13, 182, 183, 179, 191, 0,
	13, 0, 0, 0, 0, 0, 0, 13, 11, 0,
	0, 157, 0, 95, 108, 196, 0, 0, 198, 11,
	97, 181, 13, 201, -2, -2, 210, 211, 213, 214,
	228, 11, 229, 0, 0, 77, 100, 101, 0, 192,
	0, 0, 0, 230, 11, 11, 96, 215, 0, 0,
	201, 217, 0, 219, 158, 11, 161, 11, 216, 0,
	201, 201, 208, 218, 159, 162, 201, 208, 208, 208,
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
		//line parser.y:193
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:195
		{
		}
	case 3:
		//line parser.y:197
		{
		}
	case 4:
		//line parser.y:199
		{
		}
	case 5:
		//line parser.y:201
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:203
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:205
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:211
		{
		}
	case 11:
		//line parser.y:213
		{
		}
	case 12:
		//line parser.y:214
		{
		}
	case 13:
		//line parser.y:217
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 14:
		//line parser.y:219
		{
		}
	case 15:
		//line parser.y:221
		{
		}
	case 16:
		//line parser.y:223
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 17:
		//line parser.y:225
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
		//line parser.y:235
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 57:
		//line parser.y:242
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-2].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 58:
		//line parser.y:249
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 59:
		//line parser.y:253
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 60:
		//line parser.y:260
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 61:
		//line parser.y:267
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 62:
		//line parser.y:274
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 63:
		//line parser.y:281
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 64:
		//line parser.y:289
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 65:
		//line parser.y:297
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 66:
		//line parser.y:305
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 67:
		//line parser.y:313
		{
			methodName := RubyS[Rubypt-2].genericValue.(ast.BareReference).Name + "="
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: methodName},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 68:
		//line parser.y:324
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 69:
		//line parser.y:331
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 70:
		//line parser.y:339
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 71:
		//line parser.y:347
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 72:
		//line parser.y:355
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 73:
		//line parser.y:363
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 74:
		//line parser.y:373
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 75:
		//line parser.y:383
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 76:
		//line parser.y:392
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 77:
		//line parser.y:394
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 78:
		//line parser.y:396
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 79:
		//line parser.y:398
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 80:
		//line parser.y:401
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 81:
		//line parser.y:403
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 82:
		//line parser.y:405
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 83:
		//line parser.y:407
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 84:
		//line parser.y:409
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:411
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 86:
		//line parser.y:413
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 87:
		//line parser.y:417
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 88:
		//line parser.y:419
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 89:
		//line parser.y:421
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 90:
		//line parser.y:423
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-1].genericValue)
		}
	case 91:
		//line parser.y:425
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 92:
		//line parser.y:428
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 93:
		//line parser.y:430
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 94:
		//line parser.y:432
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 95:
		//line parser.y:438
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 96:
		//line parser.y:446
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-7].genericValue,
				Name:   RubyS[Rubypt-5].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-4].genericSlice,
				Body:   RubyS[Rubypt-2].genericSlice,
			}
		}
	case 97:
		//line parser.y:455
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-5].operator},
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 98:
		//line parser.y:463
		{
		}
	case 99:
		//line parser.y:465
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 100:
		//line parser.y:467
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{RubyS[Rubypt-2].genericValue}}
		}
	case 101:
		//line parser.y:469
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{ast.Break{}},
			}
		}
	case 102:
		//line parser.y:477
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 103:
		//line parser.y:479
		{
		}
	case 104:
		//line parser.y:481
		{
		}
	case 105:
		//line parser.y:483
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 106:
		//line parser.y:485
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 107:
		//line parser.y:487
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 108:
		//line parser.y:489
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-1].genericSlice...)
		}
	case 109:
		//line parser.y:492
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 110:
		//line parser.y:494
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 111:
		//line parser.y:497
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 112:
		//line parser.y:499
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 113:
		//line parser.y:501
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 114:
		//line parser.y:503
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 115:
		//line parser.y:505
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 116:
		//line parser.y:507
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 117:
		//line parser.y:511
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 118:
		//line parser.y:516
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 119:
		//line parser.y:523
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 120:
		//line parser.y:533
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 121:
		//line parser.y:542
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 122:
		//line parser.y:548
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 123:
		//line parser.y:556
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 124:
		//line parser.y:560
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 125:
		//line parser.y:565
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 126:
		//line parser.y:572
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 127:
		//line parser.y:576
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 128:
		//line parser.y:583
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 129:
		//line parser.y:590
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 130:
		//line parser.y:597
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 131:
		//line parser.y:604
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 132:
		//line parser.y:612
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 133:
		//line parser.y:614
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 134:
		//line parser.y:617
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 135:
		//line parser.y:619
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 136:
		//line parser.y:622
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 137:
		//line parser.y:624
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 138:
		//line parser.y:627
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 139:
		//line parser.y:629
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 140:
		//line parser.y:631
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 141:
		//line parser.y:632
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 142:
		//line parser.y:633
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 143:
		//line parser.y:634
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 144:
		//line parser.y:637
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 145:
		//line parser.y:646
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 146:
		//line parser.y:655
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 147:
		//line parser.y:664
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 148:
		//line parser.y:673
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 149:
		//line parser.y:682
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 150:
		//line parser.y:690
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 151:
		//line parser.y:691
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 152:
		//line parser.y:693
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 153:
		//line parser.y:695
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
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
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 156:
		//line parser.y:714
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 157:
		//line parser.y:716
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 158:
		//line parser.y:723
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 159:
		//line parser.y:730
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 160:
		//line parser.y:738
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 161:
		//line parser.y:745
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 162:
		//line parser.y:752
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 163:
		//line parser.y:760
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 164:
		//line parser.y:762
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 165:
		//line parser.y:769
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 166:
		//line parser.y:776
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 167:
		//line parser.y:779
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 168:
		//line parser.y:781
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 169:
		//line parser.y:783
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 170:
		//line parser.y:785
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 171:
		//line parser.y:788
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 172:
		//line parser.y:795
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 173:
		//line parser.y:803
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 174:
		//line parser.y:810
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 175:
		//line parser.y:817
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 176:
		//line parser.y:824
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 177:
		//line parser.y:831
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 178:
		//line parser.y:838
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 179:
		//line parser.y:845
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 180:
		//line parser.y:853
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 181:
		//line parser.y:862
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 182:
		//line parser.y:869
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 183:
		//line parser.y:876
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 184:
		//line parser.y:883
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 185:
		//line parser.y:890
		{
		}
	case 186:
		//line parser.y:891
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 187:
		//line parser.y:892
		{
		}
	case 188:
		//line parser.y:895
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 189:
		//line parser.y:898
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 190:
		//line parser.y:906
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 191:
		//line parser.y:908
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 192:
		//line parser.y:917
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
	case 193:
		//line parser.y:932
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 194:
		//line parser.y:934
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 195:
		//line parser.y:937
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 196:
		//line parser.y:939
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 197:
		//line parser.y:942
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 198:
		//line parser.y:951
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 199:
		//line parser.y:960
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 200:
		//line parser.y:969
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 201:
		//line parser.y:972
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 202:
		//line parser.y:974
		{
		}
	case 203:
		//line parser.y:976
		{
		}
	case 204:
		//line parser.y:978
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 205:
		//line parser.y:980
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 206:
		//line parser.y:982
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 207:
		//line parser.y:984
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 208:
		//line parser.y:986
		{
		}
	case 209:
		//line parser.y:988
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 210:
		//line parser.y:990
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 211:
		//line parser.y:992
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 212:
		//line parser.y:994
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 213:
		//line parser.y:996
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 214:
		//line parser.y:998
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 215:
		//line parser.y:1001
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 216:
		//line parser.y:1008
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 217:
		//line parser.y:1016
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 218:
		//line parser.y:1023
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 219:
		//line parser.y:1031
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 220:
		//line parser.y:1039
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 221:
		//line parser.y:1046
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 222:
		//line parser.y:1053
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 223:
		//line parser.y:1060
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 224:
		//line parser.y:1068
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 225:
		//line parser.y:1071
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 226:
		//line parser.y:1073
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 227:
		//line parser.y:1076
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 228:
		//line parser.y:1078
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 229:
		//line parser.y:1081
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 230:
		//line parser.y:1083
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	}
	goto Rubystack /* stack new state and value */
}
