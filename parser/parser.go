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

//line parser.y:1151

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 39,
	54, 129,
	-2, 22,
	-1, 101,
	54, 129,
	-2, 127,
	-1, 103,
	10, 86,
	11, 86,
	-2, 197,
	-1, 115,
	13, 15,
	15, 15,
	18, 15,
	19, 15,
	20, 15,
	22, 15,
	24, 15,
	32, 15,
	36, 15,
	52, 15,
	-2, 13,
	-1, 117,
	54, 129,
	-2, 22,
	-1, 239,
	54, 130,
	-2, 128,
	-1, 249,
	10, 86,
	11, 86,
	-2, 197,
	-1, 401,
	27, 213,
	28, 213,
	-2, 15,
	-1, 402,
	27, 213,
	28, 213,
	-2, 15,
}

const RubyNprod = 243
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 2770

var RubyAct = []int{

	225, 302, 5, 422, 139, 308, 251, 277, 181, 183,
	185, 26, 140, 17, 77, 105, 233, 327, 206, 112,
	104, 57, 150, 233, 151, 2, 3, 286, 53, 357,
	57, 116, 63, 117, 64, 23, 355, 114, 155, 205,
	156, 114, 4, 97, 231, 132, 133, 144, 100, 102,
	75, 74, 288, 411, 72, 73, 137, 94, 227, 135,
	65, 233, 326, 72, 73, 233, 233, 76, 66, 67,
	68, 69, 324, 128, 142, 77, 233, 77, 160, 131,
	129, 71, 70, 322, 166, 147, 60, 323, 115, 171,
	71, 70, 148, 40, 176, 232, 178, 179, 157, 186,
	392, 255, 149, 152, 153, 154, 77, 77, 189, 230,
	130, 401, 402, 220, 221, 186, 193, 142, 184, 95,
	142, 191, 96, 209, 210, 211, 212, 213, 214, 215,
	216, 217, 118, 204, 208, 142, 198, 200, 93, 77,
	187, 412, 400, 228, 94, 233, 186, 241, 118, 184,
	118, 188, 284, 429, 118, 227, 187, 77, 240, 142,
	118, 118, 118, 118, 182, 195, 349, 188, 350, 81,
	293, 121, 81, 118, 118, 118, 118, 246, 118, 118,
	118, 118, 247, 118, 118, 118, 118, 187, 118, 351,
	386, 291, 81, 387, 254, 253, 258, 118, 188, 121,
	118, 118, 118, 260, 122, 79, 80, 202, 79, 80,
	118, 123, 413, 92, 384, 118, 92, 385, 126, 9,
	78, 250, 268, 78, 89, 127, 145, 365, 79, 80,
	366, 124, 283, 82, 83, 84, 92, 118, 125, 118,
	369, 81, 363, 78, 87, 85, 86, 89, 431, 298,
	265, 284, 306, 284, 282, 118, 314, 338, 111, 306,
	316, 311, 299, 142, 312, 284, 437, 297, 434, 433,
	432, 101, 434, 433, 136, 239, 138, 79, 80, 332,
	141, 244, 318, 245, 321, 92, 399, 342, 335, 362,
	263, 397, 78, 352, 375, 360, 118, 374, 410, 162,
	163, 164, 165, 354, 167, 168, 169, 170, 364, 172,
	173, 174, 175, 372, 177, 367, 272, 379, 368, 331,
	330, 118, 364, 194, 373, 381, 196, 197, 199, 313,
	253, 376, 118, 354, 79, 80, 111, 296, 231, 257,
	81, 194, 92, 118, 118, 256, 382, 383, 329, 78,
	331, 330, 275, 231, 224, 428, 389, 294, 118, 262,
	263, 252, 158, 237, 159, 194, 393, 394, 218, 306,
	396, 99, 201, 98, 180, 161, 79, 80, 62, 110,
	235, 111, 118, 118, 92, 403, 404, 405, 406, 118,
	271, 78, 226, 79, 80, 234, 1, 146, 52, 51,
	50, 92, 364, 49, 118, 48, 47, 33, 78, 32,
	224, 419, 224, 118, 31, 30, 342, 342, 342, 43,
	426, 343, 267, 224, 435, 344, 19, 224, 415, 416,
	417, 20, 21, 22, 439, 287, 307, 342, 14, 224,
	12, 342, 342, 342, 81, 118, 11, 285, 309, 18,
	118, 10, 436, 16, 13, 15, 29, 28, 111, 24,
	59, 34, 440, 441, 25, 58, 0, 442, 0, 194,
	300, 305, 0, 0, 0, 0, 0, 0, 305, 0,
	79, 80, 0, 0, 320, 82, 83, 84, 92, 0,
	224, 0, 0, 224, 0, 78, 87, 85, 86, 89,
	224, 224, 0, 0, 0, 118, 0, 0, 136, 353,
	0, 0, 0, 0, 359, 361, 0, 0, 0, 0,
	0, 57, 116, 63, 117, 103, 0, 109, 114, 0,
	136, 0, 0, 0, 41, 0, 0, 224, 0, 353,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 65, 0, 224, 72, 73, 0, 224, 107, 66,
	67, 68, 69, 0, 0, 108, 0, 0, 0, 0,
	0, 136, 0, 119, 0, 0, 391, 106, 0, 115,
	0, 71, 70, 143, 0, 0, 0, 0, 305, 119,
	0, 119, 0, 0, 0, 119, 0, 224, 224, 0,
	224, 119, 119, 119, 119, 0, 0, 42, 224, 0,
	224, 0, 0, 0, 119, 119, 119, 119, 224, 119,
	119, 119, 119, 0, 119, 119, 119, 119, 224, 119,
	359, 420, 0, 0, 224, 0, 0, 0, 119, 0,
	0, 119, 119, 119, 0, 0, 120, 0, 0, 0,
	0, 119, 0, 0, 0, 0, 119, 0, 0, 0,
	0, 0, 120, 0, 120, 0, 0, 0, 120, 0,
	0, 0, 0, 0, 120, 120, 120, 120, 119, 0,
	119, 0, 0, 0, 0, 0, 0, 120, 120, 120,
	120, 0, 120, 120, 120, 120, 119, 120, 120, 120,
	120, 0, 120, 0, 0, 0, 0, 0, 0, 0,
	0, 120, 0, 0, 120, 120, 120, 0, 0, 0,
	0, 0, 229, 0, 120, 0, 0, 236, 0, 120,
	0, 0, 0, 0, 0, 0, 0, 119, 0, 0,
	0, 0, 57, 116, 63, 117, 103, 0, 207, 114,
	0, 120, 0, 120, 0, 0, 0, 0, 0, 0,
	0, 0, 119, 0, 0, 0, 0, 0, 0, 120,
	0, 0, 65, 119, 0, 72, 73, 0, 0, 264,
	66, 67, 68, 69, 119, 119, 266, 0, 0, 0,
	295, 0, 0, 0, 0, 134, 0, 0, 106, 119,
	115, 0, 71, 70, 0, 0, 0, 0, 0, 0,
	120, 0, 0, 0, 0, 289, 0, 0, 290, 292,
	0, 0, 0, 119, 119, 0, 0, 0, 0, 0,
	119, 0, 0, 0, 0, 120, 0, 0, 0, 0,
	0, 0, 0, 27, 0, 119, 120, 319, 0, 190,
	0, 192, 0, 0, 119, 0, 0, 120, 120, 0,
	0, 0, 0, 203, 0, 0, 0, 0, 0, 0,
	0, 0, 120, 0, 0, 356, 0, 358, 0, 0,
	0, 219, 113, 0, 0, 0, 119, 0, 0, 0,
	0, 119, 0, 0, 0, 0, 120, 120, 113, 0,
	113, 0, 0, 120, 113, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 120, 0,
	0, 0, 0, 113, 113, 113, 113, 120, 113, 113,
	113, 113, 0, 113, 113, 113, 113, 0, 113, 0,
	0, 0, 0, 0, 0, 0, 119, 113, 0, 0,
	113, 113, 113, 270, 0, 273, 0, 0, 0, 120,
	113, 0, 0, 0, 120, 113, 0, 0, 0, 280,
	281, 0, 0, 0, 409, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 113, 0, 113,
	0, 0, 418, 0, 0, 0, 57, 116, 63, 117,
	64, 0, 0, 0, 430, 113, 0, 0, 0, 317,
	0, 0, 0, 0, 0, 438, 0, 0, 0, 120,
	0, 0, 0, 0, 0, 0, 65, 0, 333, 72,
	73, 0, 0, 337, 66, 67, 68, 69, 0, 0,
	0, 0, 233, 0, 0, 0, 113, 0, 0, 371,
	0, 0, 60, 0, 61, 0, 71, 70, 0, 0,
	0, 81, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 113, 0, 0, 0, 88, 0, 0, 0, 0,
	377, 378, 113, 0, 0, 0, 0, 380, 0, 0,
	90, 91, 0, 113, 113, 0, 0, 79, 80, 388,
	0, 390, 82, 83, 84, 92, 0, 0, 113, 0,
	0, 0, 78, 87, 85, 86, 89, 0, 0, 0,
	0, 0, 0, 0, 0, 398, 0, 0, 0, 0,
	0, 219, 113, 113, 0, 0, 0, 0, 408, 113,
	0, 0, 0, 0, 0, 0, 0, 0, 414, 0,
	280, 281, 0, 0, 113, 0, 0, 0, 57, 38,
	63, 39, 64, 113, 0, 0, 35, 425, 345, 424,
	423, 346, 36, 37, 0, 55, 0, 46, 0, 0,
	347, 348, 0, 0, 0, 54, 0, 0, 65, 56,
	0, 72, 73, 0, 0, 113, 66, 67, 68, 69,
	113, 0, 0, 0, 340, 341, 0, 0, 0, 0,
	0, 0, 0, 0, 60, 0, 61, 0, 71, 70,
	57, 38, 63, 39, 64, 0, 0, 0, 35, 421,
	345, 424, 423, 346, 36, 37, 0, 55, 0, 46,
	0, 0, 347, 348, 0, 0, 0, 54, 0, 0,
	65, 56, 0, 72, 73, 113, 0, 0, 66, 67,
	68, 69, 0, 0, 0, 0, 340, 341, 0, 0,
	0, 0, 0, 0, 0, 0, 60, 0, 61, 0,
	71, 70, 57, 38, 63, 39, 64, 0, 0, 0,
	35, 427, 345, 0, 0, 346, 36, 37, 0, 55,
	0, 46, 0, 0, 347, 348, 0, 0, 0, 54,
	0, 0, 65, 56, 0, 72, 73, 0, 0, 0,
	66, 67, 68, 69, 0, 0, 0, 0, 340, 341,
	0, 0, 0, 0, 0, 0, 0, 0, 60, 0,
	61, 0, 71, 70, 57, 38, 63, 39, 64, 0,
	0, 0, 35, 395, 44, 0, 0, 45, 36, 37,
	0, 55, 0, 46, 284, 0, 0, 0, 0, 0,
	310, 54, 0, 0, 65, 56, 0, 72, 73, 0,
	0, 0, 66, 67, 68, 69, 0, 0, 0, 0,
	303, 304, 0, 0, 0, 0, 0, 0, 0, 0,
	60, 0, 61, 0, 71, 70, 57, 38, 63, 39,
	64, 0, 0, 0, 35, 339, 345, 0, 0, 346,
	36, 37, 0, 55, 0, 46, 0, 0, 347, 348,
	0, 0, 0, 54, 0, 0, 65, 56, 0, 72,
	73, 0, 0, 0, 66, 67, 68, 69, 0, 0,
	0, 0, 340, 341, 0, 0, 0, 0, 0, 0,
	0, 0, 60, 0, 61, 0, 71, 70, 57, 38,
	63, 39, 64, 0, 0, 0, 35, 334, 44, 279,
	278, 45, 36, 37, 0, 55, 0, 46, 0, 0,
	0, 0, 0, 0, 0, 54, 0, 0, 65, 56,
	0, 72, 73, 0, 0, 0, 66, 67, 68, 69,
	0, 0, 0, 0, 222, 223, 0, 0, 0, 0,
	0, 0, 0, 0, 60, 0, 61, 0, 71, 70,
	57, 38, 63, 39, 64, 0, 0, 0, 35, 315,
	44, 0, 0, 45, 36, 37, 0, 55, 0, 46,
	284, 0, 0, 0, 0, 0, 310, 54, 0, 0,
	65, 56, 0, 72, 73, 0, 0, 0, 66, 67,
	68, 69, 0, 0, 0, 0, 303, 304, 0, 0,
	0, 0, 0, 0, 0, 0, 60, 0, 61, 0,
	71, 70, 57, 38, 63, 39, 64, 0, 0, 0,
	35, 301, 44, 0, 0, 45, 36, 37, 0, 55,
	0, 46, 284, 0, 0, 0, 0, 0, 310, 54,
	0, 0, 65, 56, 0, 72, 73, 0, 0, 0,
	66, 67, 68, 69, 0, 0, 0, 0, 303, 304,
	0, 0, 0, 0, 0, 0, 0, 0, 60, 0,
	61, 0, 71, 70, 57, 38, 63, 39, 64, 0,
	0, 0, 35, 276, 44, 279, 278, 45, 36, 37,
	0, 55, 0, 46, 0, 0, 0, 0, 0, 0,
	0, 54, 0, 0, 65, 56, 0, 72, 73, 0,
	0, 0, 66, 67, 68, 69, 0, 0, 0, 0,
	222, 223, 0, 57, 38, 63, 39, 64, 0, 0,
	60, 35, 61, 44, 71, 70, 45, 36, 37, 0,
	55, 0, 46, 0, 0, 0, 0, 0, 0, 0,
	54, 0, 0, 65, 56, 0, 72, 73, 0, 0,
	0, 66, 67, 68, 69, 0, 0, 0, 0, 6,
	7, 0, 0, 0, 0, 0, 0, 0, 0, 60,
	0, 61, 0, 71, 70, 0, 8, 57, 38, 63,
	39, 64, 0, 0, 0, 35, 0, 345, 0, 0,
	346, 36, 37, 0, 55, 0, 46, 0, 0, 347,
	348, 0, 0, 0, 54, 0, 0, 65, 56, 0,
	72, 73, 0, 0, 0, 66, 67, 68, 69, 0,
	0, 0, 0, 340, 341, 0, 0, 0, 0, 0,
	0, 0, 0, 60, 0, 61, 0, 71, 70, 57,
	38, 63, 39, 64, 0, 0, 0, 35, 407, 44,
	0, 0, 45, 36, 37, 0, 55, 0, 46, 0,
	0, 0, 0, 0, 0, 0, 54, 0, 0, 65,
	56, 0, 72, 73, 0, 0, 0, 66, 67, 68,
	69, 0, 0, 0, 0, 222, 223, 0, 0, 0,
	0, 0, 0, 0, 0, 60, 0, 61, 0, 71,
	70, 57, 38, 63, 39, 64, 0, 0, 0, 35,
	370, 44, 0, 0, 45, 36, 37, 0, 55, 0,
	46, 0, 0, 0, 0, 0, 0, 0, 54, 0,
	0, 65, 56, 0, 72, 73, 0, 0, 0, 66,
	67, 68, 69, 0, 0, 0, 0, 222, 223, 0,
	0, 0, 0, 0, 0, 0, 0, 60, 0, 61,
	0, 71, 70, 57, 38, 63, 39, 64, 0, 0,
	0, 35, 336, 44, 0, 0, 45, 36, 37, 0,
	55, 0, 46, 0, 0, 0, 0, 0, 0, 0,
	54, 0, 0, 65, 56, 0, 72, 73, 0, 0,
	0, 66, 67, 68, 69, 0, 0, 0, 0, 222,
	223, 0, 57, 38, 63, 39, 64, 0, 0, 60,
	35, 61, 44, 71, 70, 45, 36, 37, 0, 55,
	0, 46, 0, 0, 0, 0, 0, 0, 0, 54,
	0, 0, 65, 56, 0, 72, 73, 0, 0, 0,
	66, 67, 68, 69, 0, 0, 0, 0, 222, 223,
	0, 0, 0, 0, 0, 0, 0, 0, 60, 0,
	61, 328, 71, 70, 57, 38, 63, 39, 64, 0,
	0, 0, 35, 325, 44, 0, 0, 45, 36, 37,
	0, 55, 0, 46, 0, 0, 0, 0, 0, 0,
	0, 54, 0, 0, 65, 56, 0, 72, 73, 0,
	0, 0, 66, 67, 68, 69, 0, 0, 0, 0,
	222, 223, 0, 57, 38, 63, 39, 64, 0, 0,
	60, 35, 61, 44, 71, 70, 45, 36, 37, 0,
	55, 0, 46, 0, 0, 0, 0, 0, 0, 0,
	54, 0, 0, 65, 56, 0, 72, 73, 0, 0,
	0, 66, 67, 68, 69, 0, 0, 0, 0, 222,
	223, 0, 0, 0, 0, 0, 0, 0, 0, 60,
	0, 61, 274, 71, 70, 57, 38, 63, 39, 64,
	0, 0, 0, 35, 269, 44, 0, 0, 45, 36,
	37, 0, 55, 0, 46, 0, 0, 0, 0, 0,
	0, 0, 54, 0, 0, 65, 56, 0, 72, 73,
	0, 0, 0, 66, 67, 68, 69, 0, 0, 0,
	0, 222, 223, 0, 0, 0, 0, 0, 0, 0,
	0, 60, 0, 61, 0, 71, 70, 57, 38, 63,
	39, 64, 0, 0, 0, 35, 261, 44, 0, 0,
	45, 36, 37, 0, 55, 0, 46, 0, 0, 0,
	0, 0, 0, 0, 54, 0, 0, 65, 56, 0,
	72, 73, 0, 0, 0, 66, 67, 68, 69, 0,
	0, 0, 0, 222, 223, 0, 0, 0, 0, 0,
	0, 0, 0, 60, 0, 61, 0, 71, 70, 57,
	38, 63, 39, 64, 0, 0, 0, 35, 259, 44,
	0, 0, 45, 36, 37, 0, 55, 0, 46, 0,
	0, 0, 0, 0, 0, 0, 54, 0, 0, 65,
	56, 0, 72, 73, 0, 0, 0, 66, 67, 68,
	69, 0, 0, 0, 0, 222, 223, 0, 57, 38,
	63, 39, 64, 0, 0, 60, 35, 61, 44, 71,
	70, 45, 36, 37, 0, 55, 0, 46, 0, 0,
	0, 0, 0, 0, 0, 54, 0, 0, 65, 56,
	0, 72, 73, 0, 0, 0, 66, 67, 68, 69,
	0, 0, 0, 0, 222, 223, 0, 57, 38, 63,
	39, 64, 243, 0, 60, 35, 61, 44, 71, 70,
	45, 36, 37, 0, 55, 0, 46, 0, 0, 0,
	0, 0, 0, 0, 54, 0, 0, 65, 56, 0,
	72, 73, 0, 0, 0, 66, 67, 68, 69, 0,
	0, 0, 0, 0, 242, 0, 57, 38, 63, 39,
	64, 0, 0, 60, 35, 61, 44, 71, 70, 45,
	36, 37, 0, 55, 0, 46, 0, 0, 0, 0,
	0, 0, 0, 54, 0, 0, 65, 56, 0, 72,
	73, 0, 0, 0, 66, 67, 68, 69, 57, 116,
	63, 117, 249, 0, 0, 114, 0, 0, 0, 0,
	0, 0, 60, 0, 61, 0, 71, 70, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 65, 0,
	0, 72, 73, 0, 0, 248, 66, 67, 68, 69,
	57, 116, 63, 117, 64, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 60, 0, 115, 0, 71, 70,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	65, 0, 0, 72, 73, 0, 0, 0, 66, 67,
	68, 69, 0, 0, 0, 0, 233, 57, 116, 63,
	117, 64, 0, 0, 114, 0, 60, 0, 61, 0,
	71, 70, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 65, 0, 0,
	72, 73, 0, 0, 0, 66, 67, 68, 69, 57,
	238, 63, 117, 64, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 60, 0, 115, 0, 71, 70, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	0, 0, 72, 73, 0, 0, 0, 66, 67, 68,
	69, 0, 0, 0, 0, 233, 57, 116, 63, 117,
	103, 0, 0, 114, 0, 60, 0, 61, 0, 71,
	70, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 65, 0, 0, 72,
	73, 0, 0, 0, 66, 67, 68, 69, 57, 116,
	63, 117, 64, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 106, 0, 115, 0, 71, 70, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 65, 0,
	0, 72, 73, 0, 0, 0, 66, 67, 68, 69,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 60, 0, 61, 0, 71, 70,
}
var RubyPact = []int{

	-26, 1698, -1000, -1000, -1000, -1, -1000, -1000, -1000, 1057,
	-1000, -1000, -1000, 120, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 104, -12, -1000,
	-1000, -1000, -1000, -1000, -1000, 367, 263, 263, 516, 162,
	189, 176, 31, 68, 2441, 2441, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 2703, 2441, 2703, -1000, -1000, -1000,
	2703, -1000, -7, 217, -1000, 29, 16, 16, 16, 16,
	32, 356, -1000, -1000, -1000, -1000, -1000, 2441, 369, 2703,
	2703, 2703, 2703, 2441, 2703, 2703, 2703, 2703, 2441, 2703,
	2703, 2703, 2703, 2441, 2703, 2441, 2441, 368, 109, 140,
	81, -1000, -1000, 2703, -1000, 154, 2703, 2703, 2703, 366,
	196, 336, -1000, 17, -17, -17, 2661, 190, -1000, -1000,
	-1000, 2703, 2441, 2441, 2441, 2441, 2441, 2441, 2441, 2441,
	2441, 362, 124, 62, 2343, 144, 336, 92, 336, 47,
	33, 440, -1000, 2614, 267, 2703, 2392, -1000, -17, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 275, -1000, -1000,
	124, 2483, 336, 336, 336, 336, 124, 336, 336, 336,
	336, 124, 168, 336, 336, 336, 124, 336, 124, 124,
	-1000, -1000, 355, 184, 93, -1000, 59, 339, 333, -1000,
	2294, 263, 2232, 349, 440, -1000, 188, 165, -1000, 165,
	-1000, -1000, 2572, 2170, -1000, 310, -1000, 2108, 342, 124,
	124, 124, 124, 124, 124, 124, 124, 124, -1000, 1649,
	-1000, -1000, -1000, -1000, 1057, 124, 240, 2703, -1000, 15,
	-1000, -1000, -1000, -1000, 180, 159, 10, 353, 737, -1000,
	327, 124, -1000, -1000, -1000, -1000, 154, 29, 2441, 2703,
	2703, 1587, 140, 93, 319, 2441, -1000, -1000, 1525, -1000,
	-1000, -1000, 29, -1000, 25, 45, 14, 336, -1000, -1000,
	2059, 6, -1000, 1997, -1000, -1000, -1000, 334, 2441, -1000,
	1463, 1948, -1000, -1000, 249, 336, 1401, 152, 2703, 2525,
	-28, -1000, -35, -1000, 2441, 2703, -1000, -1000, 124, 279,
	336, -1000, 228, -1000, -1000, 1057, 124, -1000, -1000, 212,
	2703, -1000, -1000, -1000, 124, -1000, 226, 1886, -1000, 991,
	336, -1000, 307, 2441, 291, -1000, -1000, 288, -1000, -1000,
	2441, -1000, 124, 2343, -1000, 303, -1000, 2343, 321, -1000,
	-1000, -1000, 124, -1000, -1000, 2441, 2441, 199, 175, -1000,
	-1000, 2703, 144, 440, -1000, -1000, 2525, -1000, 94, 1057,
	124, 336, -1000, -1000, -1000, 2441, 2441, 144, 1339, -1000,
	-1000, 285, -1000, 124, -1000, -1000, 124, 2343, 2343, -1000,
	2343, 280, 91, 60, 2441, 2441, 2441, 2441, 1824, 144,
	2343, 294, 0, 124, 124, -1000, 127, 202, 2343, -1000,
	-1000, -1000, -1000, 124, 124, 124, 124, -1000, 2343, 10,
	2441, 2703, -1000, -1000, 2343, 1215, 1153, 1277, 10, 142,
	237, -1000, 256, 2441, -1000, -1000, 252, -1000, -1000, -1000,
	10, -1000, -1000, 2441, -1000, 124, 1762, -1000, 10, 124,
	1762, 1762, 1762,
}
var RubyPgo = []int{

	0, 0, 465, 464, 461, 19, 460, 459, 457, 607,
	456, 5, 28, 455, 454, 453, 219, 35, 534, 843,
	451, 449, 448, 13, 446, 10, 93, 440, 438, 11,
	436, 435, 433, 432, 431, 426, 425, 421, 3, 419,
	415, 414, 409, 407, 406, 405, 403, 400, 399, 398,
	748, 397, 1, 20, 18, 7, 396, 8, 395, 27,
	392, 12, 6, 390, 4, 380, 379, 15, 9, 378,
	355, 583,
}
var RubyR1 = []int{

	0, 56, 56, 56, 56, 56, 56, 56, 56, 56,
	56, 70, 70, 71, 71, 50, 50, 50, 50, 50,
	17, 17, 17, 17, 17, 17, 17, 17, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	23, 23, 23, 23, 23, 23, 23, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 19, 19, 19, 19, 19, 19, 19, 19, 19,
	19, 19, 19, 19, 19, 19, 19, 19, 19, 19,
	53, 53, 53, 53, 64, 64, 61, 61, 61, 61,
	61, 67, 67, 67, 67, 67, 66, 66, 66, 20,
	20, 20, 20, 20, 20, 30, 30, 30, 30, 62,
	62, 62, 62, 62, 62, 57, 57, 25, 25, 25,
	25, 68, 68, 68, 24, 24, 27, 29, 29, 69,
	69, 14, 14, 14, 14, 14, 14, 14, 28, 28,
	28, 28, 28, 28, 9, 9, 26, 26, 18, 18,
	39, 39, 40, 41, 42, 43, 44, 45, 46, 47,
	48, 49, 2, 6, 7, 7, 3, 3, 58, 58,
	58, 58, 65, 65, 65, 5, 5, 5, 5, 54,
	63, 63, 63, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 55, 55, 55, 55, 51, 51, 51,
	8, 15, 11, 11, 11, 60, 60, 52, 52, 21,
	22, 12, 35, 59, 59, 59, 59, 59, 59, 36,
	36, 36, 36, 36, 36, 36, 37, 37, 37, 37,
	37, 38, 38, 38, 38, 34, 33, 10, 32, 32,
	31, 31, 4,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 1, 0, 2, 0, 2, 2, 2, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 4, 5, 1, 4, 4, 2, 3, 4, 4,
	5, 3, 5, 2, 3, 3, 3, 3, 4, 6,
	3, 7, 1, 5, 1, 3, 0, 1, 1, 4,
	4, 1, 1, 4, 4, 5, 1, 3, 3, 5,
	6, 7, 8, 5, 6, 0, 1, 3, 3, 0,
	2, 2, 2, 2, 2, 1, 3, 1, 2, 3,
	2, 0, 1, 3, 4, 6, 4, 1, 3, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 2, 2, 2, 3, 3,
	3, 3, 2, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 1, 1, 3, 3, 5, 5, 0, 1,
	3, 7, 3, 7, 8, 3, 4, 4, 3, 3,
	0, 1, 3, 4, 5, 3, 3, 3, 3, 3,
	5, 6, 5, 4, 3, 3, 2, 0, 2, 2,
	3, 4, 2, 3, 5, 0, 2, 1, 2, 2,
	2, 5, 5, 0, 2, 2, 2, 2, 2, 0,
	1, 3, 3, 1, 3, 3, 5, 6, 5, 6,
	5, 4, 3, 3, 2, 3, 3, 2, 5, 7,
	4, 5, 3,
}
var RubyChk = []int{

	-1000, -56, 51, 52, 68, -1, 51, 52, 68, -16,
	-20, -24, -27, -14, -28, -13, -15, -23, -21, -35,
	-34, -33, -32, -17, -7, -3, -29, -19, -8, -10,
	-40, -41, -42, -43, -4, 13, 19, 20, 6, 8,
	-26, -18, -9, -39, 15, 18, 24, -44, -45, -46,
	-47, -48, -49, -12, 32, 22, 36, 5, -2, -6,
	61, 63, -69, 7, 9, 35, 43, 44, 45, 46,
	66, 65, 38, 39, 52, 51, 68, 15, 55, 40,
	41, 4, 45, 46, 47, 57, 58, 56, 18, 59,
	33, 34, 48, 18, 40, 15, 18, 55, 6, 4,
	-29, 8, -29, 9, -53, -67, 61, 42, 49, 11,
	-66, -16, -5, -19, 12, 63, 6, 8, -26, -18,
	-9, 9, 42, 49, 42, 49, 42, 49, 42, 49,
	42, 11, -1, -1, -50, -64, -16, -1, -16, -64,
	-61, -16, -23, -71, 54, 9, -51, -5, 63, -17,
	6, 8, -17, -17, -17, 6, 8, 66, 6, 8,
	-1, 6, -16, -16, -16, -16, -1, -16, -16, -16,
	-16, -1, -16, -16, -16, -16, -1, -16, -1, -1,
	6, -57, 55, -68, 9, -25, 6, 47, 58, -57,
	-50, 40, -50, -61, -16, 11, -16, -16, -12, -16,
	-12, 6, 11, -50, -54, 56, -54, -50, -61, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, 6, -50,
	51, 52, 51, 52, -16, -1, -60, 11, 51, -71,
	62, 11, 62, 51, -58, -65, -71, -16, 6, 8,
	-61, -1, 52, 10, 6, 8, -67, -53, 42, 9,
	53, -62, 6, 11, -68, 42, 6, 6, -62, 14,
	-29, 14, 10, 11, -71, 62, -71, -16, -5, 14,
	-50, -63, 6, -50, 64, 10, 14, -55, 17, 16,
	-50, -50, 14, -11, 25, -16, -59, -31, 37, -71,
	-71, 11, -71, 11, 4, 53, 10, -5, -1, -61,
	-16, 14, -52, 51, 52, -16, -1, -30, -11, -22,
	31, -57, -25, 10, -1, 14, -52, -50, -5, -71,
	-16, -5, 58, 42, 58, 14, 56, 11, 64, 14,
	17, 16, -1, -50, 14, -55, 14, -50, 8, 14,
	51, 52, -1, -37, -36, 15, 18, 27, 28, 14,
	16, 37, -64, -16, -23, 64, -71, 64, -71, -16,
	-1, -16, 10, 14, -11, 15, 18, -64, -62, 14,
	14, 58, 6, -1, 6, 6, -1, -50, -50, 14,
	-50, 4, -1, -1, 15, 18, 15, 18, -50, -64,
	-50, -16, 6, -1, -1, 14, -52, 6, -50, 6,
	51, 51, 52, -1, -1, -1, -1, 14, -50, -71,
	4, 53, 14, 10, -50, -59, -59, -59, -71, -1,
	-16, 14, -38, 17, 16, 14, -38, 14, -70, 11,
	-71, 11, 14, 17, 16, -1, -59, 14, -71, -1,
	-59, -59, -59,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 47,
	48, 49, 50, 51, 52, 53, 54, 55, 56, 57,
	58, 59, 60, 28, 29, 30, 31, 32, 33, 34,
	35, 36, 37, 38, 39, 0, 0, 0, 21, -2,
	23, 24, 25, 0, 0, 0, 15, 40, 41, 42,
	43, 44, 45, 46, 0, 0, 0, 20, 26, 27,
	86, 13, 0, 63, 197, 0, 0, 0, 0, 0,
	0, 0, 162, 163, 5, 6, 7, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 121, 121,
	15, -2, 15, -2, 66, 73, 86, 0, 0, 0,
	82, 91, 92, 32, 15, -2, 21, -2, 23, 24,
	25, 86, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 15, 0, 205, 209, 84, 0, 13, 0,
	0, 84, 88, 13, 0, 86, 0, 237, 15, 152,
	21, 22, 153, 154, 155, 146, 147, 0, 144, 145,
	185, 67, 74, 76, 77, 156, 157, 158, 159, 160,
	161, 187, 0, 235, 236, 242, 189, 75, 186, 188,
	71, 109, 0, 115, 121, 122, 117, 0, 0, 109,
	0, 0, 0, 0, 87, 13, 84, 131, 132, 138,
	139, 150, 13, 0, 15, 180, 15, 0, 0, 133,
	140, 134, 141, 135, 142, 136, 143, 137, 151, 0,
	15, 15, 16, 17, 18, 19, 0, 0, 213, 0,
	164, 13, 165, 14, 13, 13, 169, 0, 21, -2,
	0, 198, 199, 200, 148, 149, 68, 69, 0, -2,
	0, 0, 121, 0, 0, 0, 118, 120, 0, 124,
	15, 126, 61, 13, 0, 78, 0, 97, 98, 175,
	0, 0, 181, 0, 178, 65, 183, 0, 0, 15,
	0, 0, 201, 206, 15, 85, 0, 0, 0, 0,
	0, 13, 0, 13, 0, 0, 64, 70, 72, 0,
	211, 99, 0, 110, 111, 47, 113, 114, 207, 106,
	0, 109, 123, 116, 119, 103, 0, 0, 62, 0,
	93, 94, 0, 0, 0, 176, 179, 0, 177, 184,
	0, 15, 15, 196, 190, 0, 192, 202, 15, 212,
	214, 215, 216, 217, 218, 0, 0, 220, 223, 238,
	15, 0, 15, 89, 90, 166, 0, 167, 0, 47,
	170, 172, 80, 100, 208, 0, 0, 210, 0, 104,
	125, 0, 95, 79, 83, 182, 15, 194, 195, 191,
	203, 0, 15, 0, 0, 0, 0, 0, 0, 15,
	13, 0, 0, 107, 108, 101, 0, 0, 193, 15,
	213, -2, -2, 221, 222, 224, 225, 239, 13, 240,
	0, 0, 102, 81, 204, 0, 0, 0, 241, 11,
	13, 226, 0, 0, 213, 228, 0, 230, 171, 12,
	173, 13, 227, 0, 213, 213, 219, 229, 174, 213,
	219, 219, 219,
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
		//line parser.y:200
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:202
		{
		}
	case 3:
		//line parser.y:204
		{
		}
	case 4:
		//line parser.y:206
		{
		}
	case 5:
		//line parser.y:208
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:210
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:212
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:218
		{
		}
	case 11:
		//line parser.y:220
		{
		}
	case 12:
		//line parser.y:221
		{
		}
	case 13:
		//line parser.y:223
		{
		}
	case 14:
		//line parser.y:224
		{
		}
	case 15:
		//line parser.y:227
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 16:
		//line parser.y:229
		{
		}
	case 17:
		//line parser.y:231
		{
		}
	case 18:
		//line parser.y:233
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 19:
		//line parser.y:235
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
		//line parser.y:247
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 62:
		//line parser.y:254
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-2].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 63:
		//line parser.y:261
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 64:
		//line parser.y:265
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 65:
		//line parser.y:272
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 66:
		//line parser.y:279
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 67:
		//line parser.y:286
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 68:
		//line parser.y:293
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 69:
		//line parser.y:301
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 70:
		//line parser.y:309
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 71:
		//line parser.y:317
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 72:
		//line parser.y:325
		{
			methodName := RubyS[Rubypt-2].genericValue.(ast.BareReference).Name + "="
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: methodName},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 73:
		//line parser.y:336
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 74:
		//line parser.y:343
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 75:
		//line parser.y:351
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 76:
		//line parser.y:359
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 77:
		//line parser.y:367
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 78:
		//line parser.y:377
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 79:
		//line parser.y:387
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 80:
		//line parser.y:396
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 81:
		//line parser.y:398
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 82:
		//line parser.y:400
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 83:
		//line parser.y:402
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
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
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 87:
		//line parser.y:411
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 88:
		//line parser.y:413
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 89:
		//line parser.y:415
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 90:
		//line parser.y:417
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 91:
		//line parser.y:421
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
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
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-1].genericValue)
		}
	case 95:
		//line parser.y:429
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 96:
		//line parser.y:432
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 97:
		//line parser.y:434
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 98:
		//line parser.y:436
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 99:
		//line parser.y:442
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 100:
		//line parser.y:450
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 101:
		//line parser.y:459
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 102:
		//line parser.y:468
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 103:
		//line parser.y:478
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 104:
		//line parser.y:486
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 105:
		//line parser.y:496
		{
		}
	case 106:
		//line parser.y:498
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 107:
		//line parser.y:500
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{RubyS[Rubypt-2].genericValue}}
		}
	case 108:
		//line parser.y:502
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{ast.Break{}},
			}
		}
	case 109:
		//line parser.y:510
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 110:
		//line parser.y:512
		{
		}
	case 111:
		//line parser.y:514
		{
		}
	case 112:
		//line parser.y:516
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 113:
		//line parser.y:518
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 114:
		//line parser.y:520
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 115:
		//line parser.y:523
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 116:
		//line parser.y:525
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 117:
		//line parser.y:528
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 118:
		//line parser.y:530
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 119:
		//line parser.y:532
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 120:
		//line parser.y:534
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 121:
		//line parser.y:536
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 122:
		//line parser.y:538
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 123:
		//line parser.y:542
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 124:
		//line parser.y:547
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 125:
		//line parser.y:554
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 126:
		//line parser.y:564
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 127:
		//line parser.y:573
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 128:
		//line parser.y:579
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 129:
		//line parser.y:587
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 130:
		//line parser.y:591
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 131:
		//line parser.y:596
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 132:
		//line parser.y:603
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 133:
		//line parser.y:607
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 134:
		//line parser.y:614
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 135:
		//line parser.y:621
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 136:
		//line parser.y:628
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 137:
		//line parser.y:635
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 138:
		//line parser.y:643
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 139:
		//line parser.y:650
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 140:
		//line parser.y:654
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 141:
		//line parser.y:661
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 142:
		//line parser.y:668
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 143:
		//line parser.y:675
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 144:
		//line parser.y:683
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 145:
		//line parser.y:685
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 146:
		//line parser.y:688
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 147:
		//line parser.y:690
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 148:
		//line parser.y:693
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 149:
		//line parser.y:695
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 150:
		//line parser.y:698
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 151:
		//line parser.y:700
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 152:
		//line parser.y:702
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 153:
		//line parser.y:703
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 154:
		//line parser.y:704
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 155:
		//line parser.y:705
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 156:
		//line parser.y:708
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 157:
		//line parser.y:717
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 158:
		//line parser.y:726
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 159:
		//line parser.y:735
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 160:
		//line parser.y:744
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 161:
		//line parser.y:753
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 162:
		//line parser.y:761
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 163:
		//line parser.y:762
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 164:
		//line parser.y:764
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 165:
		//line parser.y:766
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 166:
		//line parser.y:769
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 167:
		//line parser.y:777
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 168:
		//line parser.y:785
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 169:
		//line parser.y:787
		{
		}
	case 170:
		//line parser.y:789
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 171:
		//line parser.y:796
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 172:
		//line parser.y:804
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 173:
		//line parser.y:811
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 174:
		//line parser.y:818
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 175:
		//line parser.y:826
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 176:
		//line parser.y:828
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 177:
		//line parser.y:835
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 178:
		//line parser.y:842
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 179:
		//line parser.y:845
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 180:
		//line parser.y:847
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 181:
		//line parser.y:849
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 182:
		//line parser.y:851
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 183:
		//line parser.y:854
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 184:
		//line parser.y:861
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 185:
		//line parser.y:869
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 186:
		//line parser.y:876
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 187:
		//line parser.y:883
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 188:
		//line parser.y:890
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 189:
		//line parser.y:897
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 190:
		//line parser.y:904
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 191:
		//line parser.y:911
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 192:
		//line parser.y:919
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 193:
		//line parser.y:928
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 194:
		//line parser.y:935
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 195:
		//line parser.y:942
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 196:
		//line parser.y:949
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 197:
		//line parser.y:956
		{
		}
	case 198:
		//line parser.y:957
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 199:
		//line parser.y:958
		{
		}
	case 200:
		//line parser.y:961
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 201:
		//line parser.y:964
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 202:
		//line parser.y:972
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 203:
		//line parser.y:974
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 204:
		//line parser.y:983
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
	case 205:
		//line parser.y:998
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 206:
		//line parser.y:1000
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 207:
		//line parser.y:1003
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 208:
		//line parser.y:1005
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 209:
		//line parser.y:1008
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 210:
		//line parser.y:1017
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 211:
		//line parser.y:1026
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 212:
		//line parser.y:1035
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 213:
		//line parser.y:1038
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 214:
		//line parser.y:1040
		{
		}
	case 215:
		//line parser.y:1042
		{
		}
	case 216:
		//line parser.y:1044
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 217:
		//line parser.y:1046
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 218:
		//line parser.y:1048
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 219:
		//line parser.y:1050
		{
		}
	case 220:
		//line parser.y:1052
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 221:
		//line parser.y:1054
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 222:
		//line parser.y:1056
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 223:
		//line parser.y:1058
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 224:
		//line parser.y:1060
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 225:
		//line parser.y:1062
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 226:
		//line parser.y:1065
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 227:
		//line parser.y:1072
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 228:
		//line parser.y:1080
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 229:
		//line parser.y:1087
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 230:
		//line parser.y:1095
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 231:
		//line parser.y:1103
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 232:
		//line parser.y:1110
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 233:
		//line parser.y:1117
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 234:
		//line parser.y:1124
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 235:
		//line parser.y:1132
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 236:
		//line parser.y:1135
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 237:
		//line parser.y:1137
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 238:
		//line parser.y:1140
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 239:
		//line parser.y:1142
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 240:
		//line parser.y:1145
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 241:
		//line parser.y:1147
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 242:
		//line parser.y:1149
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
