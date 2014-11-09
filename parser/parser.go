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

//line parser.y:1141

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 42,
	54, 124,
	-2, 21,
	-1, 107,
	54, 124,
	-2, 122,
	-1, 109,
	10, 91,
	11, 91,
	-2, 193,
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
	54, 124,
	-2, 21,
	-1, 250,
	54, 125,
	-2, 123,
	-1, 260,
	10, 91,
	11, 91,
	-2, 193,
	-1, 399,
	27, 209,
	28, 209,
	-2, 15,
	-1, 400,
	27, 209,
	28, 209,
	-2, 15,
}

const RubyNprod = 239
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 2788

var RubyAct = []int{

	236, 314, 5, 420, 288, 313, 190, 146, 194, 192,
	56, 111, 28, 17, 24, 25, 110, 218, 120, 87,
	2, 3, 244, 244, 61, 157, 217, 158, 205, 120,
	103, 163, 151, 164, 118, 360, 358, 4, 409, 147,
	244, 244, 242, 238, 134, 82, 326, 327, 138, 139,
	266, 135, 106, 108, 100, 85, 86, 77, 78, 143,
	88, 89, 90, 98, 195, 141, 330, 193, 145, 155,
	84, 93, 91, 92, 95, 200, 392, 276, 149, 82,
	155, 80, 79, 168, 76, 75, 156, 159, 160, 161,
	175, 165, 99, 243, 241, 180, 297, 43, 81, 195,
	185, 82, 187, 188, 154, 196, 82, 195, 299, 238,
	193, 329, 198, 191, 82, 398, 197, 365, 274, 304,
	150, 244, 244, 149, 210, 212, 149, 82, 209, 221,
	222, 223, 224, 225, 226, 227, 228, 229, 216, 124,
	196, 149, 399, 400, 302, 204, 206, 137, 196, 202,
	239, 197, 427, 132, 252, 124, 82, 124, 124, 197,
	133, 264, 124, 232, 233, 410, 149, 220, 124, 124,
	124, 124, 124, 214, 101, 127, 295, 102, 136, 411,
	87, 124, 257, 124, 124, 124, 124, 258, 124, 124,
	124, 124, 251, 124, 124, 124, 124, 130, 124, 100,
	87, 83, 386, 265, 131, 387, 384, 124, 128, 385,
	124, 124, 124, 271, 94, 129, 85, 86, 397, 127,
	124, 83, 369, 366, 98, 124, 255, 293, 256, 96,
	97, 84, 152, 295, 295, 95, 85, 86, 295, 294,
	87, 88, 89, 90, 98, 317, 264, 341, 124, 279,
	124, 84, 93, 91, 92, 95, 352, 107, 353, 435,
	309, 432, 431, 307, 242, 240, 250, 318, 124, 395,
	315, 247, 381, 316, 149, 320, 85, 86, 375, 354,
	430, 87, 432, 431, 98, 286, 242, 87, 429, 426,
	335, 84, 206, 308, 374, 95, 338, 372, 345, 379,
	310, 334, 333, 273, 274, 283, 363, 355, 322, 408,
	268, 105, 124, 104, 357, 367, 267, 85, 86, 166,
	87, 167, 367, 85, 86, 98, 275, 373, 88, 89,
	90, 98, 84, 263, 376, 277, 124, 357, 84, 93,
	91, 92, 95, 230, 213, 85, 86, 124, 189, 382,
	383, 170, 66, 98, 116, 87, 85, 86, 124, 124,
	84, 246, 389, 300, 98, 282, 301, 303, 237, 261,
	245, 84, 1, 124, 394, 305, 332, 153, 334, 333,
	55, 54, 53, 52, 51, 401, 402, 403, 404, 44,
	50, 85, 86, 35, 34, 323, 367, 124, 124, 98,
	33, 32, 46, 346, 124, 347, 84, 19, 20, 417,
	21, 85, 86, 22, 345, 345, 345, 298, 424, 98,
	14, 124, 433, 359, 12, 361, 84, 11, 23, 18,
	10, 125, 437, 16, 13, 345, 37, 15, 31, 345,
	345, 345, 30, 26, 63, 36, 27, 125, 62, 125,
	125, 0, 124, 0, 125, 0, 0, 124, 0, 0,
	125, 125, 125, 125, 125, 0, 0, 0, 0, 0,
	0, 0, 0, 125, 0, 125, 125, 125, 125, 0,
	125, 125, 125, 125, 0, 125, 125, 125, 125, 0,
	125, 0, 0, 45, 0, 413, 414, 415, 0, 125,
	0, 0, 125, 125, 125, 0, 0, 124, 0, 0,
	0, 407, 125, 0, 0, 0, 0, 125, 0, 434,
	0, 0, 0, 0, 0, 0, 0, 416, 0, 438,
	439, 0, 0, 0, 440, 126, 0, 0, 0, 428,
	125, 0, 125, 0, 0, 0, 0, 0, 0, 0,
	436, 126, 0, 126, 126, 0, 0, 0, 126, 0,
	125, 0, 0, 0, 126, 126, 126, 126, 126, 0,
	0, 0, 0, 0, 0, 0, 0, 126, 0, 126,
	126, 126, 126, 0, 126, 126, 126, 126, 0, 126,
	126, 126, 126, 0, 126, 0, 61, 122, 67, 123,
	68, 0, 0, 126, 125, 0, 126, 126, 126, 0,
	0, 0, 0, 9, 0, 0, 126, 0, 0, 0,
	0, 126, 0, 0, 0, 0, 69, 0, 125, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 125,
	0, 0, 244, 0, 126, 0, 126, 0, 0, 371,
	125, 125, 64, 0, 65, 117, 76, 75, 0, 0,
	0, 0, 0, 0, 126, 125, 0, 0, 0, 0,
	0, 142, 0, 144, 142, 0, 0, 0, 148, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 162, 125,
	125, 0, 0, 0, 0, 0, 125, 169, 0, 171,
	172, 173, 174, 0, 176, 177, 178, 179, 126, 181,
	182, 183, 184, 125, 186, 0, 0, 29, 0, 0,
	0, 0, 0, 203, 0, 0, 207, 208, 211, 0,
	0, 0, 126, 0, 0, 0, 117, 0, 0, 0,
	0, 203, 0, 126, 125, 0, 0, 0, 0, 125,
	0, 0, 0, 0, 126, 126, 0, 0, 0, 119,
	0, 0, 0, 0, 248, 0, 203, 0, 0, 126,
	219, 0, 0, 0, 0, 119, 0, 119, 119, 0,
	0, 0, 119, 0, 117, 0, 0, 0, 0, 0,
	0, 0, 119, 126, 126, 0, 0, 0, 0, 125,
	126, 119, 0, 119, 119, 119, 119, 0, 119, 119,
	119, 119, 0, 119, 119, 119, 119, 126, 119, 0,
	140, 0, 0, 0, 0, 0, 0, 119, 278, 0,
	119, 119, 119, 0, 0, 0, 0, 0, 0, 0,
	119, 0, 0, 0, 0, 119, 0, 0, 126, 0,
	0, 0, 296, 126, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 117, 0, 0, 0, 0, 119, 0,
	119, 0, 0, 0, 203, 311, 0, 199, 0, 201,
	0, 0, 0, 0, 0, 0, 0, 0, 119, 324,
	0, 215, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 126, 0, 0, 0, 0, 0, 231,
	0, 0, 0, 142, 356, 0, 0, 0, 0, 362,
	364, 0, 0, 0, 61, 122, 67, 123, 68, 0,
	0, 0, 119, 0, 0, 0, 0, 356, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 0, 119, 77, 78, 0,
	0, 262, 70, 71, 72, 73, 74, 119, 142, 269,
	244, 0, 0, 391, 0, 0, 0, 325, 119, 119,
	64, 0, 65, 0, 76, 75, 0, 281, 0, 284,
	0, 0, 0, 119, 0, 61, 122, 67, 123, 109,
	0, 115, 120, 291, 292, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 119, 119, 0,
	0, 0, 362, 418, 119, 69, 0, 0, 77, 78,
	0, 0, 113, 70, 71, 72, 73, 74, 0, 114,
	0, 119, 321, 0, 0, 0, 0, 0, 0, 0,
	0, 112, 0, 121, 0, 76, 75, 0, 0, 0,
	0, 336, 0, 0, 0, 0, 340, 0, 0, 0,
	0, 0, 119, 0, 0, 0, 0, 119, 0, 0,
	61, 122, 67, 123, 109, 0, 368, 120, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 377, 378, 0, 0, 0,
	69, 0, 380, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 0, 388, 0, 390, 119, 306, 0,
	0, 0, 0, 0, 0, 0, 112, 0, 121, 0,
	76, 75, 0, 0, 0, 0, 0, 396, 0, 0,
	0, 0, 0, 231, 0, 0, 0, 0, 0, 0,
	406, 0, 0, 0, 0, 0, 0, 0, 412, 0,
	291, 292, 61, 41, 67, 42, 68, 0, 0, 0,
	38, 423, 348, 422, 421, 349, 39, 40, 0, 58,
	0, 49, 0, 0, 350, 351, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 343, 344,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 0,
	65, 0, 76, 75, 61, 41, 67, 42, 68, 0,
	0, 0, 38, 419, 348, 422, 421, 349, 39, 40,
	0, 58, 0, 49, 0, 0, 350, 351, 0, 0,
	60, 57, 0, 0, 69, 59, 0, 77, 78, 0,
	0, 0, 70, 71, 72, 73, 74, 0, 0, 0,
	343, 344, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 0, 65, 0, 76, 75, 61, 41, 67, 42,
	68, 0, 0, 0, 38, 425, 348, 0, 0, 349,
	39, 40, 0, 58, 0, 49, 0, 0, 350, 351,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 343, 344, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 0, 76, 75, 61, 41,
	67, 42, 68, 0, 0, 0, 38, 342, 348, 0,
	0, 349, 39, 40, 0, 58, 0, 49, 0, 0,
	350, 351, 0, 0, 60, 57, 0, 0, 69, 59,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 343, 344, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 0, 65, 0, 76, 75,
	61, 41, 67, 42, 68, 0, 0, 0, 38, 337,
	47, 290, 289, 48, 39, 40, 0, 58, 0, 49,
	0, 0, 0, 0, 0, 0, 60, 57, 0, 0,
	69, 59, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 0, 0, 0, 234, 235, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 0, 65, 0,
	76, 75, 61, 41, 67, 42, 68, 0, 0, 0,
	38, 287, 47, 290, 289, 48, 39, 40, 0, 58,
	0, 49, 0, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 234, 235,
	0, 61, 41, 67, 42, 68, 0, 0, 64, 38,
	65, 47, 76, 75, 48, 39, 40, 0, 58, 0,
	49, 0, 0, 0, 0, 0, 0, 60, 57, 0,
	0, 69, 59, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 0, 0, 0, 6, 7, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 0, 65,
	0, 76, 75, 0, 8, 61, 41, 67, 42, 68,
	0, 0, 0, 38, 0, 348, 0, 0, 349, 39,
	40, 0, 58, 0, 49, 0, 0, 350, 351, 0,
	0, 60, 57, 0, 0, 69, 59, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 0, 0,
	0, 343, 344, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 0, 65, 0, 76, 75, 61, 41, 67,
	42, 68, 0, 0, 0, 38, 393, 47, 0, 0,
	48, 39, 40, 0, 58, 0, 49, 295, 0, 0,
	0, 0, 0, 60, 57, 0, 0, 69, 59, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 234, 235, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 0, 65, 0, 76, 75, 61,
	41, 67, 42, 68, 0, 0, 0, 38, 319, 47,
	0, 0, 48, 39, 40, 0, 58, 0, 49, 295,
	0, 0, 0, 0, 0, 60, 57, 0, 0, 69,
	59, 0, 77, 78, 0, 0, 0, 70, 71, 72,
	73, 74, 0, 0, 0, 234, 235, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 0, 65, 0, 76,
	75, 61, 41, 67, 42, 68, 0, 0, 0, 38,
	312, 47, 0, 0, 48, 39, 40, 0, 58, 0,
	49, 295, 0, 0, 0, 0, 0, 60, 57, 0,
	0, 69, 59, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 0, 0, 0, 234, 235, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 0, 65,
	0, 76, 75, 61, 41, 67, 42, 68, 0, 0,
	0, 38, 405, 47, 0, 0, 48, 39, 40, 0,
	58, 0, 49, 0, 0, 0, 0, 0, 0, 60,
	57, 0, 0, 69, 59, 0, 77, 78, 0, 0,
	0, 70, 71, 72, 73, 74, 0, 0, 0, 234,
	235, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	0, 65, 0, 76, 75, 61, 41, 67, 42, 68,
	0, 0, 0, 38, 370, 47, 0, 0, 48, 39,
	40, 0, 58, 0, 49, 0, 0, 0, 0, 0,
	0, 60, 57, 0, 0, 69, 59, 0, 77, 78,
	0, 0, 0, 70, 71, 72, 73, 74, 0, 0,
	0, 234, 235, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 0, 65, 0, 76, 75, 61, 41, 67,
	42, 68, 0, 0, 0, 38, 339, 47, 0, 0,
	48, 39, 40, 0, 58, 0, 49, 0, 0, 0,
	0, 0, 0, 60, 57, 0, 0, 69, 59, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 234, 235, 0, 61, 41, 67, 42,
	68, 0, 0, 64, 38, 65, 47, 76, 75, 48,
	39, 40, 0, 58, 0, 49, 0, 0, 0, 0,
	0, 0, 60, 57, 0, 0, 69, 59, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 234, 235, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 331, 76, 75, 61, 41,
	67, 42, 68, 0, 0, 0, 38, 328, 47, 0,
	0, 48, 39, 40, 0, 58, 0, 49, 0, 0,
	0, 0, 0, 0, 60, 57, 0, 0, 69, 59,
	0, 77, 78, 0, 0, 0, 70, 71, 72, 73,
	74, 0, 0, 0, 234, 235, 0, 61, 41, 67,
	42, 68, 0, 0, 64, 38, 65, 47, 76, 75,
	48, 39, 40, 0, 58, 0, 49, 0, 0, 0,
	0, 0, 0, 60, 57, 0, 0, 69, 59, 0,
	77, 78, 0, 0, 0, 70, 71, 72, 73, 74,
	0, 0, 0, 234, 235, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 0, 65, 285, 76, 75, 61,
	41, 67, 42, 68, 0, 0, 0, 38, 280, 47,
	0, 0, 48, 39, 40, 0, 58, 0, 49, 0,
	0, 0, 0, 0, 0, 60, 57, 0, 0, 69,
	59, 0, 77, 78, 0, 0, 0, 70, 71, 72,
	73, 74, 0, 0, 0, 234, 235, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 0, 65, 0, 76,
	75, 61, 41, 67, 42, 68, 0, 0, 0, 38,
	272, 47, 0, 0, 48, 39, 40, 0, 58, 0,
	49, 0, 0, 0, 0, 0, 0, 60, 57, 0,
	0, 69, 59, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 0, 0, 0, 234, 235, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 0, 65,
	0, 76, 75, 61, 41, 67, 42, 68, 0, 0,
	0, 38, 270, 47, 0, 0, 48, 39, 40, 0,
	58, 0, 49, 0, 0, 0, 0, 0, 0, 60,
	57, 0, 0, 69, 59, 0, 77, 78, 0, 0,
	0, 70, 71, 72, 73, 74, 0, 0, 0, 234,
	235, 0, 61, 41, 67, 42, 68, 0, 0, 64,
	38, 65, 47, 76, 75, 48, 39, 40, 0, 58,
	0, 49, 0, 0, 0, 0, 0, 0, 60, 57,
	0, 0, 69, 59, 0, 77, 78, 0, 0, 0,
	70, 71, 72, 73, 74, 0, 0, 0, 234, 235,
	0, 61, 41, 67, 42, 68, 254, 0, 64, 38,
	65, 47, 76, 75, 48, 39, 40, 0, 58, 0,
	49, 0, 0, 0, 0, 0, 0, 60, 57, 0,
	0, 69, 59, 0, 77, 78, 0, 0, 0, 70,
	71, 72, 73, 74, 0, 0, 0, 0, 253, 0,
	61, 41, 67, 42, 68, 0, 0, 64, 38, 65,
	47, 76, 75, 48, 39, 40, 0, 58, 0, 49,
	0, 0, 0, 0, 0, 0, 60, 57, 0, 0,
	69, 59, 0, 77, 78, 0, 0, 0, 70, 71,
	72, 73, 74, 61, 122, 67, 123, 260, 0, 0,
	120, 0, 0, 0, 0, 0, 64, 0, 65, 0,
	76, 75, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 0, 0, 77, 78, 0, 0,
	259, 70, 71, 72, 73, 74, 61, 122, 67, 123,
	68, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	0, 121, 0, 76, 75, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 69, 0, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 244, 61, 122, 67, 123, 68, 0, 0,
	120, 0, 64, 0, 65, 0, 76, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 0, 0, 77, 78, 0, 0,
	0, 70, 71, 72, 73, 74, 61, 249, 67, 123,
	68, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	0, 121, 0, 76, 75, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 69, 0, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 244, 61, 122, 67, 123, 109, 0, 0,
	120, 0, 64, 0, 65, 0, 76, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 0, 0, 77, 78, 0, 0,
	0, 70, 71, 72, 73, 74, 61, 122, 67, 123,
	68, 0, 0, 0, 0, 0, 0, 0, 0, 112,
	0, 121, 0, 76, 75, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 69, 0, 0, 77,
	78, 0, 0, 0, 70, 71, 72, 73, 74, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 0, 76, 75,
}
var RubyPact = []int{

	-31, 1526, -1000, -1000, -1000, 30, -1000, -1000, -1000, 196,
	-1000, -1000, -1000, 74, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 159,
	-25, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 307, 249,
	249, 990, 166, 155, 111, 2, 136, 2455, 2455, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 2721, 2455, 2721,
	2721, -1000, -1000, -1000, 2721, -1000, -22, 223, -1000, 6,
	19, 19, 19, 19, 2721, 25, 313, -1000, -1000, -1000,
	-1000, -1000, 2455, 2721, 345, 2721, 2721, 2721, 2721, 2455,
	2721, 2721, 2721, 2721, 2455, 2721, 2721, 2721, 2721, 2455,
	2721, 2455, 2455, 342, 58, 101, 35, -1000, -1000, 2721,
	6, 17, 2721, 2721, 2721, 338, 162, 351, -1000, 14,
	-30, -30, 2678, 210, -1000, -1000, -1000, 2721, 2455, 2455,
	2455, 2455, 2455, 2455, 2455, 2455, 2455, 337, 86, 112,
	2357, 98, 351, 99, 351, 98, 32, 31, 283, -1000,
	2631, 258, 2721, 2406, -1000, -30, -1000, -1000, -1000, -1000,
	-1000, -1000, 351, -1000, -1000, 220, -1000, -1000, 86, 351,
	2498, 351, 351, 351, 351, 86, 351, 351, 351, 351,
	86, 316, 351, 351, 351, 86, 351, 86, 86, -1000,
	-1000, 327, 150, 93, -1000, 8, 310, 304, -1000, 2308,
	249, 2246, 293, 283, -1000, -1000, -1000, 15, 176, -1000,
	-1000, 236, -1000, -1000, 2588, 2184, -1000, 299, -1000, 2122,
	275, 86, 86, 86, 86, 86, 86, 86, 86, 86,
	-1000, 1477, -1000, -1000, -1000, -1000, 86, 213, 2721, -1000,
	71, -1000, -1000, -1000, -1000, 133, 108, -10, 371, 1075,
	-1000, 253, 86, -1000, -1000, -1000, -1000, 17, 6, 2455,
	2721, 2721, 1776, 101, 93, 235, 2455, -1000, -1000, 1714,
	-1000, -1000, -1000, 6, -1000, 919, 4, -11, 351, -1000,
	-1000, 2073, 55, -1000, 2011, -1000, -1000, -1000, 362, 2455,
	-1000, 1415, 1962, -1000, -1000, 239, 351, 1353, 242, 2721,
	2541, -28, -1000, -29, -1000, 2455, 2721, -1000, -1000, 86,
	107, 351, -1000, 209, -1000, -1000, -1000, -1000, 86, -1000,
	208, 1900, -1000, 591, 351, 291, 2455, 288, -1000, -1000,
	272, -1000, -1000, 2455, -1000, 86, 2357, -1000, 285, -1000,
	2357, 268, -1000, -1000, -1000, 86, -1000, -1000, 2455, 2455,
	191, 187, -1000, -1000, 2721, 98, 283, -1000, -1000, 2541,
	-1000, 70, 196, 86, 351, -1000, -1000, -1000, 1652, -1000,
	-1000, 263, -1000, 86, -1000, -1000, 86, 2357, 2357, -1000,
	2357, 212, 64, 91, 2455, 2455, 2455, 2455, 1838, 98,
	2357, 305, -15, -1000, 151, 169, 2357, -1000, -1000, -1000,
	-1000, 86, 86, 86, 86, -1000, 2357, -10, 2455, 2721,
	-1000, -1000, 2357, 1229, 1167, 1291, -10, 141, 277, -1000,
	266, 2455, -1000, -1000, 245, -1000, -1000, -1000, -10, -1000,
	-1000, 2455, -1000, 86, 1590, -1000, -10, 86, 1590, 1590,
	1590,
}
var RubyPgo = []int{

	0, 0, 448, 446, 445, 34, 444, 443, 442, 493,
	438, 1, 10, 437, 436, 434, 433, 613, 15, 389,
	717, 430, 429, 428, 13, 427, 8, 97, 424, 420,
	12, 417, 413, 410, 408, 14, 407, 405, 403, 3,
	402, 401, 400, 394, 393, 390, 384, 383, 382, 381,
	380, 770, 377, 5, 16, 17, 4, 372, 6, 370,
	96, 368, 39, 365, 7, 361, 354, 11, 9, 352,
	289, 120,
}
var RubyR1 = []int{

	0, 57, 57, 57, 57, 57, 57, 57, 57, 57,
	57, 70, 70, 71, 71, 51, 51, 51, 51, 18,
	18, 18, 18, 18, 18, 18, 18, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	24, 24, 24, 24, 24, 24, 24, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 35, 14, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 54, 54, 54, 54, 64,
	64, 62, 62, 62, 62, 62, 67, 67, 67, 67,
	67, 66, 66, 66, 21, 21, 21, 21, 21, 21,
	58, 58, 26, 26, 26, 26, 68, 68, 68, 25,
	25, 28, 30, 30, 69, 69, 15, 15, 15, 15,
	15, 15, 15, 15, 29, 29, 29, 29, 29, 29,
	9, 9, 27, 27, 19, 19, 40, 40, 41, 42,
	43, 44, 45, 46, 47, 48, 49, 50, 2, 6,
	7, 7, 3, 3, 59, 59, 59, 59, 65, 65,
	65, 5, 5, 5, 5, 55, 63, 63, 63, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 56,
	56, 56, 56, 52, 52, 52, 8, 16, 11, 11,
	11, 61, 61, 53, 53, 22, 23, 12, 36, 60,
	60, 60, 60, 60, 60, 37, 37, 37, 37, 37,
	37, 37, 38, 38, 38, 38, 38, 39, 39, 39,
	39, 34, 33, 10, 32, 32, 31, 31, 4,
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
	3, 3, 3, 4, 6, 3, 7, 1, 5, 1,
	3, 0, 1, 1, 4, 4, 1, 1, 4, 2,
	5, 1, 3, 3, 5, 6, 7, 8, 5, 6,
	1, 3, 1, 2, 3, 2, 0, 1, 3, 4,
	6, 4, 1, 3, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 3, 3, 3, 3, 2, 2,
	2, 2, 3, 3, 3, 3, 3, 3, 1, 1,
	3, 3, 5, 5, 0, 1, 3, 7, 3, 7,
	8, 3, 4, 4, 3, 3, 0, 1, 3, 4,
	5, 3, 3, 3, 3, 3, 5, 6, 5, 4,
	3, 3, 2, 0, 2, 2, 3, 4, 2, 3,
	5, 0, 2, 1, 2, 2, 2, 5, 5, 0,
	2, 2, 2, 2, 2, 0, 1, 3, 3, 1,
	3, 3, 5, 6, 5, 6, 5, 4, 3, 3,
	2, 3, 3, 2, 5, 7, 4, 5, 3,
}
var RubyChk = []int{

	-1000, -57, 51, 52, 68, -1, 51, 52, 68, -17,
	-21, -25, -28, -15, -29, -13, -16, -24, -22, -36,
	-34, -33, -32, -23, -35, -18, -7, -3, -30, -20,
	-8, -10, -41, -42, -43, -44, -4, -14, 13, 19,
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
	-71, 54, 9, -52, -5, 63, -18, 6, 8, -18,
	-18, -18, -17, 6, 8, 66, 6, 8, -1, -17,
	6, -17, -17, -17, -17, -1, -17, -17, -17, -17,
	-1, -17, -17, -17, -17, -1, -17, -1, -1, 6,
	-58, 55, -68, 9, -26, 6, 47, 58, -58, -51,
	40, -51, -62, -17, -5, 11, -5, -17, -17, -35,
	-12, -17, -12, 6, 11, -51, -55, 56, -55, -51,
	-62, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	6, -51, 51, 52, 51, 52, -1, -61, 11, 51,
	-71, 62, 11, 62, 51, -59, -65, -71, -17, 6,
	8, -62, -1, 52, 10, 6, 8, -67, -54, 42,
	9, 53, -51, 6, 11, -68, 42, 6, 6, -51,
	14, -30, 14, 10, 11, -71, 62, -71, -17, -5,
	14, -51, -63, 6, -51, 64, 10, 14, -56, 17,
	16, -51, -51, 14, -11, 25, -17, -60, -31, 37,
	-71, -71, 11, -71, 11, 4, 53, 10, -5, -1,
	-62, -17, 14, -53, -11, -58, -26, 10, -1, 14,
	-53, -51, -5, -71, -17, 58, 42, 58, 14, 56,
	11, 64, 14, 17, 16, -1, -51, 14, -56, 14,
	-51, 8, 14, 51, 52, -1, -38, -37, 15, 18,
	27, 28, 14, 16, 37, -64, -17, -24, 64, -71,
	64, -71, -17, -1, -17, 10, 14, -11, -51, 14,
	14, 58, 6, -1, 6, 6, -1, -51, -51, 14,
	-51, 4, -1, -1, 15, 18, 15, 18, -51, -64,
	-51, -17, 6, 14, -53, 6, -51, 6, 51, 51,
	52, -1, -1, -1, -1, 14, -51, -71, 4, 53,
	14, 10, -51, -60, -60, -60, -71, -1, -17, 14,
	-39, 17, 16, 14, -39, 14, -70, 11, -71, 11,
	14, 17, 16, -1, -60, 14, -71, -1, -60, -60,
	-60,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 47,
	48, 49, 50, 51, 52, 53, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 0, 0,
	0, 20, -2, 22, 23, 24, 0, 0, 0, 15,
	40, 41, 42, 43, 44, 45, 46, 0, 0, 0,
	0, 19, 25, 26, 91, 13, 0, 67, 193, 0,
	0, 0, 0, 0, 0, 0, 0, 158, 159, 5,
	6, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 116, 116, 15, -2, 15, -2,
	70, 78, 91, 0, 0, 0, 87, 96, 97, 31,
	15, -2, 20, -2, 22, 23, 24, 91, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 15, 0,
	201, 205, 89, 0, 13, 206, 0, 0, 89, 93,
	13, 0, 91, 0, 233, 15, 148, 20, 21, 149,
	150, 151, 64, 142, 143, 0, 140, 141, 181, 63,
	72, 79, 81, 82, 152, 153, 154, 155, 156, 157,
	183, 0, 231, 232, 238, 185, 80, 182, 184, 76,
	15, 0, 110, 116, 117, 112, 0, 0, 15, 0,
	0, 0, 0, 92, 71, 13, 99, 89, 126, 127,
	128, 134, 135, 146, 13, 0, 15, 176, 15, 0,
	0, 129, 136, 130, 137, 131, 138, 132, 139, 133,
	147, 0, 15, 15, 16, 17, 18, 0, 0, 209,
	0, 160, 13, 161, 14, 13, 13, 165, 0, 20,
	-2, 0, 194, 195, 196, 144, 145, 73, 74, 0,
	-2, 0, 0, 116, 0, 0, 0, 113, 115, 0,
	119, 15, 121, 65, 13, 0, 83, 0, 102, 103,
	171, 0, 0, 177, 0, 174, 69, 179, 0, 0,
	15, 0, 0, 197, 202, 15, 90, 0, 0, 0,
	0, 0, 13, 0, 13, 0, 0, 68, 75, 77,
	0, 207, 104, 0, 203, 15, 118, 111, 114, 108,
	0, 0, 66, 0, 98, 0, 0, 0, 172, 175,
	0, 173, 180, 0, 15, 15, 192, 186, 0, 188,
	198, 15, 208, 210, 211, 212, 213, 214, 0, 0,
	216, 219, 234, 15, 0, 15, 94, 95, 162, 0,
	163, 0, 47, 166, 168, 85, 105, 204, 0, 109,
	120, 0, 100, 84, 88, 178, 15, 190, 191, 187,
	199, 0, 15, 0, 0, 0, 0, 0, 0, 15,
	13, 0, 0, 106, 0, 0, 189, 15, 209, -2,
	-2, 217, 218, 220, 221, 235, 13, 236, 0, 0,
	107, 86, 200, 0, 0, 0, 237, 11, 13, 222,
	0, 0, 209, 224, 0, 226, 167, 12, 169, 13,
	223, 0, 209, 209, 215, 225, 170, 209, 215, 215,
	215,
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
		//line parser.y:399
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 85:
		//line parser.y:408
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 86:
		//line parser.y:410
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 87:
		//line parser.y:412
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 88:
		//line parser.y:414
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 89:
		//line parser.y:417
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 90:
		//line parser.y:419
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 91:
		//line parser.y:421
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 92:
		//line parser.y:423
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 93:
		//line parser.y:425
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 94:
		//line parser.y:427
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 95:
		//line parser.y:429
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 96:
		//line parser.y:433
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 97:
		//line parser.y:435
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 98:
		//line parser.y:437
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 99:
		//line parser.y:439
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 100:
		//line parser.y:441
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 101:
		//line parser.y:444
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 102:
		//line parser.y:446
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 103:
		//line parser.y:448
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 104:
		//line parser.y:454
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 105:
		//line parser.y:462
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 106:
		//line parser.y:471
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 107:
		//line parser.y:480
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 108:
		//line parser.y:490
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 109:
		//line parser.y:498
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 110:
		//line parser.y:509
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 111:
		//line parser.y:511
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 112:
		//line parser.y:514
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 113:
		//line parser.y:516
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 114:
		//line parser.y:518
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 115:
		//line parser.y:520
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 116:
		//line parser.y:522
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 117:
		//line parser.y:524
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 118:
		//line parser.y:528
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 119:
		//line parser.y:533
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 120:
		//line parser.y:540
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 121:
		//line parser.y:550
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 122:
		//line parser.y:559
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 123:
		//line parser.y:565
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 124:
		//line parser.y:573
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 125:
		//line parser.y:577
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 126:
		//line parser.y:582
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 127:
		//line parser.y:589
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 128:
		//line parser.y:593
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 129:
		//line parser.y:597
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 130:
		//line parser.y:604
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 131:
		//line parser.y:611
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 132:
		//line parser.y:618
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 133:
		//line parser.y:625
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 134:
		//line parser.y:633
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 135:
		//line parser.y:640
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 136:
		//line parser.y:644
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 137:
		//line parser.y:651
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 138:
		//line parser.y:658
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 139:
		//line parser.y:665
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 140:
		//line parser.y:673
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 141:
		//line parser.y:675
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 142:
		//line parser.y:678
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 143:
		//line parser.y:680
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 144:
		//line parser.y:683
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 145:
		//line parser.y:685
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 146:
		//line parser.y:688
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 147:
		//line parser.y:690
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 148:
		//line parser.y:692
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 149:
		//line parser.y:693
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 150:
		//line parser.y:694
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 151:
		//line parser.y:695
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 152:
		//line parser.y:698
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 153:
		//line parser.y:707
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 154:
		//line parser.y:716
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 155:
		//line parser.y:725
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 156:
		//line parser.y:734
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 157:
		//line parser.y:743
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 158:
		//line parser.y:751
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 159:
		//line parser.y:752
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 160:
		//line parser.y:754
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 161:
		//line parser.y:756
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 162:
		//line parser.y:759
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 163:
		//line parser.y:767
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 164:
		//line parser.y:775
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 165:
		//line parser.y:777
		{
		}
	case 166:
		//line parser.y:779
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 167:
		//line parser.y:786
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 168:
		//line parser.y:794
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 169:
		//line parser.y:801
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 170:
		//line parser.y:808
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 171:
		//line parser.y:816
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 172:
		//line parser.y:818
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 173:
		//line parser.y:825
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 174:
		//line parser.y:832
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 175:
		//line parser.y:835
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 176:
		//line parser.y:837
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 177:
		//line parser.y:839
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 178:
		//line parser.y:841
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 179:
		//line parser.y:844
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 180:
		//line parser.y:851
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 181:
		//line parser.y:859
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 182:
		//line parser.y:866
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 183:
		//line parser.y:873
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 184:
		//line parser.y:880
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 185:
		//line parser.y:887
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 186:
		//line parser.y:894
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 187:
		//line parser.y:901
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 188:
		//line parser.y:909
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 189:
		//line parser.y:918
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 190:
		//line parser.y:925
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 191:
		//line parser.y:932
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 192:
		//line parser.y:939
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 193:
		//line parser.y:946
		{
		}
	case 194:
		//line parser.y:947
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 195:
		//line parser.y:948
		{
		}
	case 196:
		//line parser.y:951
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 197:
		//line parser.y:954
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 198:
		//line parser.y:962
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 199:
		//line parser.y:964
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 200:
		//line parser.y:973
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
	case 201:
		//line parser.y:988
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 202:
		//line parser.y:990
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 203:
		//line parser.y:993
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 204:
		//line parser.y:995
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 205:
		//line parser.y:998
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 206:
		//line parser.y:1007
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 207:
		//line parser.y:1016
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 208:
		//line parser.y:1025
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 209:
		//line parser.y:1028
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 210:
		//line parser.y:1030
		{
		}
	case 211:
		//line parser.y:1032
		{
		}
	case 212:
		//line parser.y:1034
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 213:
		//line parser.y:1036
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 214:
		//line parser.y:1038
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 215:
		//line parser.y:1040
		{
		}
	case 216:
		//line parser.y:1042
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 217:
		//line parser.y:1044
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 218:
		//line parser.y:1046
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 219:
		//line parser.y:1048
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 220:
		//line parser.y:1050
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 221:
		//line parser.y:1052
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 222:
		//line parser.y:1055
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 223:
		//line parser.y:1062
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 224:
		//line parser.y:1070
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 225:
		//line parser.y:1077
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 226:
		//line parser.y:1085
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 227:
		//line parser.y:1093
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 228:
		//line parser.y:1100
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 229:
		//line parser.y:1107
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 230:
		//line parser.y:1114
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 231:
		//line parser.y:1122
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 232:
		//line parser.y:1125
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 233:
		//line parser.y:1127
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 234:
		//line parser.y:1130
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 235:
		//line parser.y:1132
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 236:
		//line parser.y:1135
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 237:
		//line parser.y:1137
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 238:
		//line parser.y:1139
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
