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

//line parser.y:1145

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 42,
	54, 125,
	-2, 21,
	-1, 107,
	54, 125,
	-2, 123,
	-1, 109,
	10, 92,
	11, 92,
	-2, 194,
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
	54, 125,
	-2, 21,
	-1, 249,
	54, 126,
	-2, 124,
	-1, 259,
	10, 92,
	11, 92,
	-2, 194,
	-1, 401,
	27, 210,
	28, 210,
	-2, 15,
	-1, 402,
	27, 210,
	28, 210,
	-2, 15,
}

const RubyNprod = 240
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 2815

var RubyAct = []int{

	235, 315, 5, 422, 316, 146, 188, 290, 192, 190,
	147, 111, 23, 17, 27, 82, 2, 3, 110, 217,
	161, 243, 162, 243, 203, 120, 241, 216, 278, 332,
	56, 118, 22, 4, 362, 120, 360, 237, 243, 103,
	151, 411, 243, 301, 193, 329, 193, 191, 138, 139,
	328, 80, 79, 193, 106, 108, 191, 243, 134, 143,
	100, 267, 198, 141, 101, 135, 145, 102, 81, 394,
	82, 156, 157, 82, 331, 82, 155, 242, 149, 150,
	163, 82, 127, 166, 132, 194, 155, 194, 240, 100,
	173, 133, 299, 189, 194, 178, 195, 9, 195, 99,
	183, 154, 185, 186, 82, 195, 401, 402, 412, 231,
	232, 400, 196, 371, 243, 128, 137, 238, 237, 297,
	200, 130, 129, 149, 297, 206, 149, 306, 131, 220,
	221, 222, 223, 224, 225, 226, 227, 228, 219, 117,
	215, 149, 202, 204, 209, 211, 208, 136, 368, 295,
	354, 304, 355, 388, 251, 142, 389, 144, 142, 297,
	297, 386, 148, 250, 387, 437, 149, 434, 433, 265,
	158, 159, 160, 356, 429, 213, 180, 181, 82, 413,
	256, 167, 127, 169, 170, 171, 172, 257, 174, 175,
	176, 177, 87, 179, 367, 275, 182, 432, 184, 434,
	433, 266, 87, 381, 152, 336, 335, 201, 399, 343,
	205, 207, 210, 272, 334, 107, 336, 335, 319, 265,
	117, 96, 97, 83, 239, 201, 87, 249, 85, 86,
	246, 96, 97, 431, 309, 241, 288, 241, 85, 86,
	254, 296, 255, 84, 164, 281, 165, 95, 247, 397,
	201, 274, 275, 84, 377, 96, 97, 95, 105, 311,
	104, 87, 85, 86, 376, 374, 117, 285, 269, 268,
	312, 317, 322, 149, 318, 264, 229, 84, 261, 262,
	212, 187, 168, 276, 383, 428, 66, 116, 204, 310,
	96, 97, 337, 279, 245, 284, 236, 85, 86, 244,
	347, 340, 1, 153, 55, 54, 324, 357, 365, 53,
	260, 280, 84, 410, 52, 51, 359, 50, 34, 33,
	369, 302, 32, 31, 303, 305, 46, 369, 348, 375,
	87, 349, 19, 36, 37, 298, 378, 20, 300, 359,
	14, 307, 96, 97, 12, 11, 117, 21, 18, 85,
	86, 384, 385, 10, 24, 325, 16, 201, 313, 96,
	97, 13, 391, 35, 84, 320, 85, 86, 15, 30,
	96, 97, 396, 29, 326, 25, 63, 85, 86, 26,
	62, 84, 0, 0, 361, 0, 363, 403, 404, 405,
	406, 43, 84, 0, 0, 0, 0, 0, 0, 142,
	358, 369, 0, 0, 0, 364, 366, 0, 0, 0,
	0, 419, 0, 0, 0, 0, 347, 347, 347, 0,
	426, 0, 0, 358, 435, 0, 0, 0, 0, 0,
	0, 0, 0, 124, 439, 0, 0, 347, 0, 0,
	0, 347, 347, 347, 0, 0, 0, 0, 0, 124,
	0, 124, 124, 0, 142, 0, 124, 0, 0, 393,
	0, 0, 0, 0, 124, 124, 124, 0, 0, 0,
	0, 0, 409, 0, 0, 124, 0, 124, 124, 124,
	124, 0, 124, 124, 124, 124, 0, 124, 418, 0,
	124, 0, 124, 415, 416, 417, 28, 0, 0, 0,
	430, 124, 0, 0, 124, 124, 124, 0, 364, 420,
	0, 438, 0, 0, 124, 0, 0, 436, 0, 124,
	0, 0, 0, 0, 0, 0, 0, 440, 441, 0,
	0, 0, 442, 0, 0, 0, 0, 0, 119, 0,
	0, 0, 124, 0, 124, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 119, 0, 119, 119, 0, 0,
	124, 119, 0, 0, 0, 0, 0, 0, 0, 119,
	119, 119, 124, 124, 61, 122, 67, 123, 68, 0,
	119, 0, 119, 119, 119, 119, 0, 119, 119, 119,
	119, 0, 119, 0, 0, 119, 0, 119, 0, 0,
	0, 0, 0, 0, 69, 124, 119, 77, 78, 119,
	119, 119, 70, 71, 72, 73, 74, 0, 0, 119,
	243, 0, 0, 0, 119, 0, 0, 373, 0, 124,
	64, 0, 65, 0, 76, 75, 0, 0, 0, 0,
	124, 0, 0, 0, 0, 0, 0, 119, 0, 119,
	0, 124, 124, 0, 0, 0, 0, 0, 0, 124,
	0, 0, 0, 0, 0, 119, 0, 0, 124, 0,
	0, 0, 0, 0, 0, 0, 0, 119, 119, 0,
	0, 61, 122, 67, 123, 109, 0, 115, 120, 0,
	0, 0, 0, 124, 124, 0, 0, 0, 0, 0,
	124, 44, 0, 0, 0, 0, 0, 0, 0, 0,
	119, 69, 0, 0, 77, 78, 0, 124, 113, 70,
	71, 72, 73, 74, 0, 114, 0, 0, 0, 0,
	0, 0, 0, 0, 119, 0, 0, 112, 0, 121,
	0, 76, 75, 125, 0, 119, 0, 0, 124, 0,
	0, 0, 0, 124, 0, 0, 119, 119, 0, 125,
	0, 125, 125, 0, 119, 0, 125, 218, 0, 0,
	0, 0, 0, 119, 125, 125, 125, 0, 0, 0,
	0, 0, 0, 0, 0, 125, 0, 125, 125, 125,
	125, 87, 125, 125, 125, 125, 0, 125, 119, 119,
	125, 0, 125, 124, 0, 119, 45, 0, 0, 0,
	0, 125, 0, 0, 125, 125, 125, 140, 0, 0,
	96, 97, 119, 0, 125, 0, 0, 85, 86, 125,
	0, 0, 88, 89, 90, 98, 0, 0, 0, 0,
	0, 0, 84, 93, 91, 92, 95, 0, 126, 277,
	0, 0, 125, 119, 125, 0, 0, 0, 119, 0,
	0, 0, 0, 0, 126, 0, 126, 126, 0, 0,
	125, 126, 0, 0, 197, 0, 199, 0, 0, 126,
	126, 126, 125, 125, 0, 0, 0, 0, 214, 0,
	126, 0, 126, 126, 126, 126, 0, 126, 126, 126,
	126, 0, 126, 0, 0, 126, 230, 126, 119, 0,
	0, 0, 0, 0, 0, 125, 126, 0, 0, 126,
	126, 126, 0, 0, 0, 0, 0, 0, 0, 126,
	0, 0, 0, 0, 126, 0, 0, 0, 0, 125,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	125, 0, 0, 0, 0, 0, 263, 126, 0, 126,
	0, 125, 125, 0, 270, 0, 0, 0, 0, 125,
	0, 0, 0, 0, 0, 126, 0, 0, 125, 0,
	0, 0, 0, 283, 0, 286, 0, 126, 126, 0,
	0, 0, 61, 122, 67, 123, 68, 0, 0, 293,
	294, 0, 0, 125, 125, 0, 0, 0, 0, 0,
	125, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	126, 0, 69, 0, 0, 77, 78, 125, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 243, 87,
	323, 0, 0, 0, 126, 327, 0, 0, 64, 0,
	65, 0, 76, 75, 0, 126, 0, 0, 125, 0,
	338, 0, 0, 125, 0, 342, 126, 126, 96, 97,
	0, 0, 0, 0, 126, 85, 86, 0, 0, 0,
	88, 89, 90, 126, 0, 370, 0, 0, 0, 0,
	84, 93, 91, 92, 95, 0, 0, 61, 122, 67,
	123, 68, 0, 0, 379, 380, 0, 0, 126, 126,
	0, 382, 0, 125, 0, 126, 0, 0, 0, 0,
	0, 0, 0, 390, 0, 392, 0, 69, 0, 0,
	77, 78, 126, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 243, 0, 0, 398, 0, 0, 0,
	0, 0, 230, 64, 0, 65, 0, 76, 75, 408,
	0, 0, 0, 126, 0, 0, 0, 414, 126, 293,
	294, 0, 0, 61, 41, 67, 42, 68, 0, 0,
	0, 38, 425, 350, 424, 423, 351, 39, 40, 0,
	58, 0, 49, 0, 0, 352, 353, 0, 0, 60,
	57, 0, 0, 69, 59, 0, 77, 78, 0, 0,
	0, 70, 71, 72, 73, 74, 0, 0, 126, 345,
	346, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	0, 65, 0, 76, 75, 61, 41, 67, 42, 68,
	0, 0, 0, 38, 421, 350, 424, 423, 351, 39,
	40, 0, 58, 0, 49, 0, 0, 352, 353, 0,
	0, 60, 57, 0, 0, 69, 59, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 0, 0,
	0, 345, 346, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 0, 65, 0, 76, 75, 61, 41, 67,
	42, 68, 0, 0, 0, 38, 427, 350, 0, 0,
	351, 39, 40, 0, 58, 0, 49, 0, 0, 352,
	353, 0, 0, 60, 57, 0, 0, 69, 59, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 345, 346, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 0, 65, 0, 76, 75, 61,
	41, 67, 42, 68, 0, 0, 0, 38, 344, 350,
	0, 0, 351, 39, 40, 0, 58, 0, 49, 0,
	0, 352, 353, 0, 0, 60, 57, 0, 0, 69,
	59, 0, 77, 78, 0, 0, 0, 70, 71, 72,
	73, 74, 0, 0, 0, 345, 346, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 0, 65, 0, 76,
	75, 61, 41, 67, 42, 68, 0, 0, 0, 38,
	339, 47, 292, 291, 48, 39, 40, 0, 58, 0,
	49, 0, 0, 0, 0, 0, 0, 60, 57, 0,
	0, 69, 59, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 0, 0, 0, 233, 234, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 0, 65,
	0, 76, 75, 61, 41, 67, 42, 68, 0, 0,
	0, 38, 289, 47, 292, 291, 48, 39, 40, 0,
	58, 0, 49, 0, 0, 0, 0, 0, 0, 60,
	57, 0, 0, 69, 59, 0, 77, 78, 0, 0,
	0, 70, 71, 72, 73, 74, 0, 0, 0, 233,
	234, 0, 61, 41, 67, 42, 68, 0, 0, 64,
	38, 65, 47, 76, 75, 48, 39, 40, 0, 58,
	0, 49, 0, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 6, 7,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 0,
	65, 0, 76, 75, 0, 8, 61, 41, 67, 42,
	68, 0, 0, 0, 38, 0, 350, 0, 0, 351,
	39, 40, 0, 58, 0, 49, 0, 0, 352, 353,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 345, 346, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 0, 76, 75, 61, 41,
	67, 42, 68, 0, 0, 0, 38, 395, 47, 0,
	0, 48, 39, 40, 0, 58, 0, 49, 297, 0,
	0, 0, 0, 0, 60, 57, 0, 0, 69, 59,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 233, 234, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 0, 65, 0, 76, 75,
	61, 41, 67, 42, 68, 0, 0, 0, 38, 321,
	47, 0, 0, 48, 39, 40, 0, 58, 0, 49,
	297, 0, 0, 0, 0, 0, 60, 57, 0, 0,
	69, 59, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 0, 0, 0, 233, 234, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 65, 0,
	76, 75, 61, 41, 67, 42, 68, 0, 0, 0,
	38, 314, 47, 0, 0, 48, 39, 40, 0, 58,
	0, 49, 297, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 233, 234,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 0,
	65, 0, 76, 75, 61, 41, 67, 42, 68, 0,
	0, 0, 38, 407, 47, 0, 0, 48, 39, 40,
	0, 58, 0, 49, 0, 0, 0, 0, 0, 0,
	60, 57, 0, 0, 69, 59, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 0, 0, 0,
	233, 234, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 0, 65, 0, 76, 75, 61, 41, 67, 42,
	68, 0, 0, 0, 38, 372, 47, 0, 0, 48,
	39, 40, 0, 58, 0, 49, 0, 0, 0, 0,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 233, 234, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 0, 76, 75, 61, 41,
	67, 42, 68, 0, 0, 0, 38, 341, 47, 0,
	0, 48, 39, 40, 0, 58, 0, 49, 0, 0,
	0, 0, 0, 0, 60, 57, 0, 0, 69, 59,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 233, 234, 0, 61, 41, 67,
	42, 68, 0, 0, 64, 38, 65, 47, 76, 75,
	48, 39, 40, 0, 58, 0, 49, 0, 0, 0,
	0, 0, 0, 60, 57, 0, 0, 69, 59, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 233, 234, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 0, 65, 333, 76, 75, 61,
	41, 67, 42, 68, 0, 0, 0, 38, 330, 47,
	0, 0, 48, 39, 40, 0, 58, 0, 49, 0,
	0, 0, 0, 0, 0, 60, 57, 0, 0, 69,
	59, 0, 77, 78, 0, 0, 0, 70, 71, 72,
	73, 74, 0, 0, 0, 233, 234, 0, 61, 41,
	67, 42, 68, 0, 0, 64, 38, 65, 47, 76,
	75, 48, 39, 40, 0, 58, 0, 49, 0, 0,
	0, 0, 0, 0, 60, 57, 0, 0, 69, 59,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 233, 234, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 0, 65, 287, 76, 75,
	61, 41, 67, 42, 68, 0, 0, 0, 38, 282,
	47, 0, 0, 48, 39, 40, 0, 58, 0, 49,
	0, 0, 0, 0, 0, 0, 60, 57, 0, 0,
	69, 59, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 0, 0, 0, 233, 234, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 65, 0,
	76, 75, 61, 41, 67, 42, 68, 0, 0, 0,
	38, 273, 47, 0, 0, 48, 39, 40, 0, 58,
	0, 49, 0, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 233, 234,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 0,
	65, 0, 76, 75, 61, 41, 67, 42, 68, 0,
	0, 0, 38, 271, 47, 0, 0, 48, 39, 40,
	0, 58, 0, 49, 0, 0, 0, 0, 0, 0,
	60, 57, 0, 0, 69, 59, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 0, 0, 0,
	233, 234, 0, 61, 41, 67, 42, 68, 0, 0,
	64, 38, 65, 47, 76, 75, 48, 39, 40, 0,
	58, 0, 49, 0, 0, 0, 0, 0, 0, 60,
	57, 0, 0, 69, 59, 0, 77, 78, 0, 0,
	0, 70, 71, 72, 73, 74, 0, 0, 0, 233,
	234, 0, 61, 41, 67, 42, 68, 253, 0, 64,
	38, 65, 47, 76, 75, 48, 39, 40, 0, 58,
	0, 49, 0, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 0, 252,
	0, 61, 41, 67, 42, 68, 0, 0, 64, 38,
	65, 47, 76, 75, 48, 39, 40, 0, 58, 0,
	49, 0, 0, 0, 0, 0, 0, 60, 57, 0,
	0, 69, 59, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 61, 122, 67, 123, 109, 0,
	0, 120, 0, 0, 0, 0, 0, 64, 0, 65,
	0, 76, 75, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 0, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 61, 122, 67,
	123, 259, 308, 0, 120, 0, 0, 0, 0, 0,
	112, 0, 121, 0, 76, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 0, 0,
	77, 78, 0, 0, 258, 70, 71, 72, 73, 74,
	61, 122, 67, 123, 68, 0, 0, 120, 0, 0,
	0, 0, 0, 64, 0, 121, 0, 76, 75, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	69, 0, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 61, 248, 67, 123, 68, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 121, 0,
	76, 75, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 0, 0, 77, 78, 0, 0,
	0, 70, 71, 72, 73, 74, 0, 0, 0, 243,
	61, 122, 67, 123, 109, 0, 0, 120, 0, 64,
	0, 65, 0, 76, 75, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	69, 0, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 61, 122, 67, 123, 68, 0, 0,
	0, 0, 0, 0, 0, 0, 112, 0, 121, 0,
	76, 75, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 0, 0, 77, 78, 0, 87,
	0, 70, 71, 72, 73, 74, 0, 0, 0, 0,
	0, 0, 0, 94, 0, 0, 0, 0, 0, 64,
	83, 65, 0, 76, 75, 0, 0, 0, 96, 97,
	0, 0, 0, 0, 0, 85, 86, 0, 0, 0,
	88, 89, 90, 98, 0, 0, 0, 0, 0, 0,
	84, 93, 91, 92, 95,
}
var RubyPact = []int{

	-35, 1527, -1000, -1000, -1000, 0, -1000, -1000, -1000, 2755,
	-1000, -1000, -1000, 81, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 49, -16,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 254, 207,
	207, 676, 73, 79, 42, 16, 105, 2456, 2456, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 2718, 2456, 2718,
	2718, -1000, -1000, -1000, 2718, -1000, -14, 195, -1000, 23,
	2456, 2456, 2718, 2718, 2718, 14, 238, -1000, -1000, -1000,
	-1000, -1000, 2456, 2718, 276, 2718, 2718, 2718, 2718, 2456,
	2718, 2718, 2718, 2718, 2456, 2718, -1000, -1000, 2718, 2456,
	2718, 2456, 2456, 275, 38, 47, 22, -1000, -1000, 2718,
	23, 13, 2718, 2718, 2718, 274, 164, 326, -1000, 20,
	-29, -29, 2675, 173, -1000, -1000, -1000, 2718, 2456, 2456,
	2456, 2456, 2456, 2456, 2456, 2456, 2456, 270, 89, 58,
	2358, 107, 326, 66, 326, 107, 26, 15, 1035, -1000,
	2628, 219, 2718, 2407, -1000, -29, 89, 89, 326, 326,
	326, -1000, -1000, 234, -1000, -1000, 89, 326, 2542, 326,
	326, 326, 326, 89, 326, 326, 326, 326, 89, 257,
	1092, 1092, 326, 89, 326, 89, 89, -1000, -1000, 269,
	158, 40, -1000, 19, 263, 262, -1000, 2309, 207, 2247,
	241, 1035, -1000, -1000, -1000, 787, -34, 198, -1000, -1000,
	188, -1000, -1000, 2585, 2185, -1000, 261, -1000, 2123, 226,
	89, 89, 89, 89, 89, 89, 89, 89, 89, -1000,
	1478, -1000, -1000, -1000, -1000, 89, 135, 2718, -1000, 6,
	-1000, -1000, -1000, -1000, 140, 116, -9, 337, 2499, -1000,
	224, 89, -1000, -1000, -1000, -1000, 13, 23, 2456, 2718,
	2718, 326, 326, 1777, 47, 40, 208, 2718, -1000, -1000,
	1715, -1000, -1000, -1000, 23, -1000, 987, 8, -1000, -13,
	326, -1000, -1000, 2074, 18, -1000, 2012, -1000, -1000, -1000,
	200, 2456, -1000, 1416, 1963, -1000, -1000, 201, 326, 1354,
	136, 2718, 1092, -28, -1000, -30, -1000, 2456, 2718, -1000,
	-1000, 89, 184, 326, -1000, 134, -1000, -1000, -1000, -1000,
	326, -1000, 99, 1901, -1000, 569, 326, 259, 2456, 258,
	-1000, -1000, 248, -1000, -1000, 2456, -1000, 89, 2358, -1000,
	189, -1000, 2358, 280, -1000, -1000, -1000, 89, -1000, -1000,
	2456, 2456, 146, 138, -1000, -1000, 2718, 107, 1035, -1000,
	-1000, 1092, -1000, 63, 2755, 89, 326, -1000, -1000, -1000,
	1653, -1000, -1000, 243, -1000, 89, -1000, -1000, 89, 2358,
	2358, -1000, 2358, 202, 60, 55, 2456, 2456, 2456, 2456,
	1839, 107, 2358, 309, -12, -1000, 94, 169, 2358, -1000,
	-1000, -1000, -1000, 89, 89, 89, 89, -1000, 2358, -9,
	2456, 2718, -1000, -1000, 2358, 1230, 1168, 1292, -9, 163,
	222, -1000, 183, 2456, -1000, -1000, 151, -1000, -1000, -1000,
	-9, -1000, -1000, 2456, -1000, 89, 1591, -1000, -9, 89,
	1591, 1591, 1591,
}
var RubyPgo = []int{

	0, 0, 380, 379, 12, 31, 376, 375, 373, 806,
	369, 4, 30, 368, 363, 361, 356, 97, 354, 701,
	496, 353, 348, 347, 13, 345, 8, 391, 344, 340,
	14, 338, 337, 334, 333, 32, 332, 331, 328, 3,
	326, 323, 322, 319, 318, 317, 315, 314, 309, 305,
	304, 767, 303, 1, 18, 19, 7, 302, 6, 299,
	92, 296, 10, 295, 5, 294, 287, 11, 9, 286,
	285, 79,
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
	64, 64, 62, 62, 62, 62, 62, 67, 67, 67,
	67, 67, 66, 66, 66, 21, 21, 21, 21, 21,
	21, 58, 58, 68, 68, 68, 26, 26, 26, 26,
	25, 25, 28, 30, 30, 69, 69, 15, 15, 15,
	15, 15, 15, 15, 15, 29, 29, 29, 29, 29,
	29, 9, 9, 27, 27, 19, 19, 40, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 2,
	6, 7, 7, 3, 3, 59, 59, 59, 59, 65,
	65, 65, 5, 5, 5, 5, 55, 63, 63, 63,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	56, 56, 56, 56, 52, 52, 52, 8, 16, 11,
	11, 11, 61, 61, 53, 53, 22, 23, 12, 36,
	60, 60, 60, 60, 60, 60, 37, 37, 37, 37,
	37, 37, 37, 38, 38, 38, 38, 38, 39, 39,
	39, 39, 34, 33, 10, 32, 32, 31, 31, 4,
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
	1, 3, 0, 1, 1, 4, 4, 1, 1, 4,
	2, 5, 1, 3, 3, 5, 6, 7, 8, 5,
	6, 1, 3, 0, 1, 3, 1, 2, 3, 2,
	4, 6, 4, 1, 3, 1, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 2, 2, 3, 3, 3, 3, 2,
	2, 2, 2, 3, 3, 3, 3, 3, 3, 1,
	1, 3, 3, 5, 5, 0, 1, 3, 7, 3,
	7, 8, 3, 4, 4, 3, 3, 0, 1, 3,
	4, 5, 3, 3, 3, 3, 3, 5, 6, 5,
	4, 3, 3, 2, 0, 2, 2, 3, 4, 2,
	3, 5, 0, 2, 1, 2, 2, 2, 5, 5,
	0, 2, 2, 2, 2, 2, 0, 1, 3, 3,
	1, 3, 3, 5, 6, 5, 6, 5, 4, 3,
	3, 2, 4, 4, 2, 5, 7, 4, 5, 3,
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
	-71, 54, 9, -52, -5, 63, -1, -1, -17, -17,
	-17, 6, 8, 66, 6, 8, -1, -17, 6, -17,
	-17, -17, -17, -1, -17, -17, -17, -17, -1, -17,
	-71, -71, -17, -1, -17, -1, -1, 6, -58, 55,
	-68, 9, -26, 6, 47, 58, -58, -51, 40, -51,
	-62, -17, -5, 11, -5, -17, -4, -17, -35, -12,
	-17, -12, 6, 11, -51, -55, 56, -55, -51, -62,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, 6,
	-51, 51, 52, 51, 52, -1, -61, 11, 51, -71,
	62, 11, 62, 51, -59, -65, -71, -17, 6, 8,
	-62, -1, 52, 10, 6, 8, -67, -54, 42, 9,
	53, -17, -17, -51, 6, 11, -68, 42, 6, 6,
	-51, 14, -30, 14, 10, 11, -71, 62, 62, -71,
	-17, -5, 14, -51, -63, 6, -51, 64, 10, 14,
	-56, 17, 16, -51, -51, 14, -11, 25, -17, -60,
	-31, 37, -71, -71, 11, -71, 11, 4, 53, 10,
	-5, -1, -62, -17, 14, -53, -11, -58, -26, 10,
	-17, 14, -53, -51, -5, -71, -17, 58, 42, 58,
	14, 56, 11, 64, 14, 17, 16, -1, -51, 14,
	-56, 14, -51, 8, 14, 51, 52, -1, -38, -37,
	15, 18, 27, 28, 14, 16, 37, -64, -17, -24,
	64, -71, 64, -71, -17, -1, -17, 10, 14, -11,
	-51, 14, 14, 58, 6, -1, 6, 6, -1, -51,
	-51, 14, -51, 4, -1, -1, 15, 18, 15, 18,
	-51, -64, -51, -17, 6, 14, -53, 6, -51, 6,
	51, 51, 52, -1, -1, -1, -1, 14, -51, -71,
	4, 53, 14, 10, -51, -60, -60, -60, -71, -1,
	-17, 14, -39, 17, 16, 14, -39, 14, -70, 11,
	-71, 11, 14, 17, 16, -1, -60, 14, -71, -1,
	-60, -60, -60,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 48,
	49, 50, 51, 52, 53, 54, 55, 56, 57, 58,
	59, 60, 61, 62, 27, 28, 29, 30, 31, 32,
	33, 34, 35, 36, 37, 38, 39, 40, 0, 0,
	0, 20, -2, 22, 23, 24, 0, 0, 0, 15,
	41, 42, 43, 44, 45, 46, 47, 0, 0, 0,
	0, 19, 25, 26, 92, 13, 0, 67, 194, 0,
	0, 0, 0, 0, 0, 0, 0, 159, 160, 5,
	6, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 13, 13, 0, 0,
	0, 0, 0, 0, 113, 113, 15, -2, 15, -2,
	70, 78, 92, 0, 0, 0, 88, 97, 98, 31,
	15, -2, 20, -2, 22, 23, 24, 92, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 15, 0,
	202, 206, 90, 0, 13, 207, 0, 0, 90, 94,
	13, 0, 92, 0, 234, 15, 149, 150, 151, 152,
	64, 143, 144, 0, 141, 142, 182, 63, 72, 79,
	81, 82, 153, 154, 155, 156, 157, 158, 184, 0,
	0, 0, 239, 186, 80, 183, 185, 76, 15, 0,
	111, 113, 114, 116, 0, 0, 15, 0, 0, 0,
	0, 93, 71, 13, 100, 90, 0, 127, 128, 129,
	135, 136, 147, 13, 0, 15, 177, 15, 0, 0,
	130, 137, 131, 138, 132, 139, 133, 140, 134, 148,
	0, 15, 15, 16, 17, 18, 0, 0, 210, 0,
	161, 13, 162, 14, 13, 13, 166, 0, 20, -2,
	0, 195, 196, 197, 145, 146, 73, 74, 0, -2,
	0, 232, 233, 0, 113, 0, 0, 0, 117, 119,
	0, 120, 15, 122, 65, 13, 0, 83, 84, 0,
	103, 104, 172, 0, 0, 178, 0, 175, 69, 180,
	0, 0, 15, 0, 0, 198, 203, 15, 91, 0,
	0, 0, 0, 0, 13, 0, 13, 0, 0, 68,
	75, 77, 0, 208, 105, 0, 204, 15, 115, 112,
	118, 109, 0, 0, 66, 0, 99, 0, 0, 0,
	173, 176, 0, 174, 181, 0, 15, 15, 193, 187,
	0, 189, 199, 15, 209, 211, 212, 213, 214, 215,
	0, 0, 217, 220, 235, 15, 0, 15, 95, 96,
	163, 0, 164, 0, 48, 167, 169, 86, 106, 205,
	0, 110, 121, 0, 101, 85, 89, 179, 15, 191,
	192, 188, 200, 0, 15, 0, 0, 0, 0, 0,
	0, 15, 13, 0, 0, 107, 0, 0, 190, 15,
	210, -2, -2, 218, 219, 221, 222, 236, 13, 237,
	0, 0, 108, 87, 201, 0, 0, 0, 238, 11,
	13, 223, 0, 0, 210, 225, 0, 227, 168, 12,
	170, 13, 224, 0, 210, 210, 216, 226, 171, 210,
	216, 216, 216,
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
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 96:
		//line parser.y:438
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 97:
		//line parser.y:442
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 98:
		//line parser.y:444
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 99:
		//line parser.y:446
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 100:
		//line parser.y:448
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 101:
		//line parser.y:450
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 102:
		//line parser.y:453
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 103:
		//line parser.y:455
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 104:
		//line parser.y:457
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 105:
		//line parser.y:461
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 106:
		//line parser.y:469
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 107:
		//line parser.y:478
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 108:
		//line parser.y:487
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 109:
		//line parser.y:497
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 110:
		//line parser.y:505
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 111:
		//line parser.y:516
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 112:
		//line parser.y:518
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 113:
		//line parser.y:520
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 114:
		//line parser.y:522
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 115:
		//line parser.y:524
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 116:
		//line parser.y:527
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 117:
		//line parser.y:529
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 118:
		//line parser.y:531
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 119:
		//line parser.y:533
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 120:
		//line parser.y:537
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 121:
		//line parser.y:544
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 122:
		//line parser.y:554
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 123:
		//line parser.y:563
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 124:
		//line parser.y:569
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 125:
		//line parser.y:577
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 126:
		//line parser.y:581
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 127:
		//line parser.y:586
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 128:
		//line parser.y:593
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 129:
		//line parser.y:597
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 130:
		//line parser.y:601
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 131:
		//line parser.y:608
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 132:
		//line parser.y:615
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 133:
		//line parser.y:622
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 134:
		//line parser.y:629
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 135:
		//line parser.y:637
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 136:
		//line parser.y:644
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 137:
		//line parser.y:648
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 138:
		//line parser.y:655
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 139:
		//line parser.y:662
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 140:
		//line parser.y:669
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 141:
		//line parser.y:677
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 142:
		//line parser.y:679
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 143:
		//line parser.y:682
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 144:
		//line parser.y:684
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 145:
		//line parser.y:687
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 146:
		//line parser.y:689
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 147:
		//line parser.y:692
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 148:
		//line parser.y:694
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 149:
		//line parser.y:696
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 150:
		//line parser.y:697
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 151:
		//line parser.y:698
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 152:
		//line parser.y:699
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 153:
		//line parser.y:702
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 154:
		//line parser.y:711
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 155:
		//line parser.y:720
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 156:
		//line parser.y:729
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 157:
		//line parser.y:738
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 158:
		//line parser.y:747
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 159:
		//line parser.y:755
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 160:
		//line parser.y:756
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 161:
		//line parser.y:758
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 162:
		//line parser.y:760
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 163:
		//line parser.y:763
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 164:
		//line parser.y:771
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 165:
		//line parser.y:779
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 166:
		//line parser.y:781
		{
		}
	case 167:
		//line parser.y:783
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 168:
		//line parser.y:790
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 169:
		//line parser.y:798
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 170:
		//line parser.y:805
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 171:
		//line parser.y:812
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 172:
		//line parser.y:820
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 173:
		//line parser.y:822
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 174:
		//line parser.y:829
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 175:
		//line parser.y:836
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 176:
		//line parser.y:839
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 177:
		//line parser.y:841
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 178:
		//line parser.y:843
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 179:
		//line parser.y:845
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 180:
		//line parser.y:848
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 181:
		//line parser.y:855
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 182:
		//line parser.y:863
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 183:
		//line parser.y:870
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 184:
		//line parser.y:877
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 185:
		//line parser.y:884
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 186:
		//line parser.y:891
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 187:
		//line parser.y:898
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 188:
		//line parser.y:905
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 189:
		//line parser.y:913
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 190:
		//line parser.y:922
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 191:
		//line parser.y:929
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 192:
		//line parser.y:936
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 193:
		//line parser.y:943
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 194:
		//line parser.y:950
		{
		}
	case 195:
		//line parser.y:951
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 196:
		//line parser.y:952
		{
		}
	case 197:
		//line parser.y:955
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 198:
		//line parser.y:958
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 199:
		//line parser.y:966
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 200:
		//line parser.y:968
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 201:
		//line parser.y:977
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
		//line parser.y:992
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 203:
		//line parser.y:994
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 204:
		//line parser.y:997
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 205:
		//line parser.y:999
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 206:
		//line parser.y:1002
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 207:
		//line parser.y:1011
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 208:
		//line parser.y:1020
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 209:
		//line parser.y:1029
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 210:
		//line parser.y:1032
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 211:
		//line parser.y:1034
		{
		}
	case 212:
		//line parser.y:1036
		{
		}
	case 213:
		//line parser.y:1038
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 214:
		//line parser.y:1040
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 215:
		//line parser.y:1042
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 216:
		//line parser.y:1044
		{
		}
	case 217:
		//line parser.y:1046
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 218:
		//line parser.y:1048
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 219:
		//line parser.y:1050
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 220:
		//line parser.y:1052
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 221:
		//line parser.y:1054
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 222:
		//line parser.y:1056
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 223:
		//line parser.y:1059
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 224:
		//line parser.y:1066
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 225:
		//line parser.y:1074
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 226:
		//line parser.y:1081
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 227:
		//line parser.y:1089
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 228:
		//line parser.y:1097
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 229:
		//line parser.y:1104
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 230:
		//line parser.y:1111
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 231:
		//line parser.y:1118
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 232:
		//line parser.y:1126
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 233:
		//line parser.y:1129
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 234:
		//line parser.y:1131
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 235:
		//line parser.y:1134
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 236:
		//line parser.y:1136
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 237:
		//line parser.y:1139
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 238:
		//line parser.y:1141
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 239:
		//line parser.y:1143
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
