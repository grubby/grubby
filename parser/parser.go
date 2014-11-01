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

//line parser.y:1161

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 39,
	54, 130,
	-2, 22,
	-1, 102,
	54, 130,
	-2, 128,
	-1, 104,
	10, 87,
	11, 87,
	-2, 198,
	-1, 116,
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
	-1, 118,
	54, 130,
	-2, 22,
	-1, 241,
	54, 131,
	-2, 129,
	-1, 251,
	10, 87,
	11, 87,
	-2, 198,
	-1, 404,
	27, 214,
	28, 214,
	-2, 15,
	-1, 405,
	27, 214,
	28, 214,
	-2, 15,
}

const RubyNprod = 245
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 2826

var RubyAct = []int{

	9, 304, 425, 279, 288, 253, 26, 17, 141, 187,
	140, 310, 185, 183, 106, 105, 208, 53, 57, 151,
	235, 152, 77, 2, 3, 235, 115, 235, 98, 156,
	77, 157, 23, 360, 326, 207, 145, 414, 358, 112,
	4, 235, 325, 101, 103, 257, 290, 129, 77, 329,
	113, 72, 73, 233, 130, 137, 229, 139, 75, 74,
	235, 142, 127, 415, 125, 136, 404, 405, 143, 128,
	188, 126, 193, 186, 286, 76, 188, 149, 71, 70,
	163, 164, 165, 166, 403, 168, 169, 170, 171, 158,
	173, 174, 175, 176, 328, 178, 179, 144, 81, 150,
	153, 154, 155, 395, 234, 196, 93, 232, 198, 199,
	201, 189, 143, 195, 191, 143, 148, 189, 112, 184,
	77, 81, 190, 196, 77, 432, 200, 202, 190, 77,
	143, 210, 206, 77, 79, 80, 226, 81, 229, 82,
	83, 84, 92, 197, 434, 239, 122, 196, 235, 78,
	87, 85, 86, 89, 143, 242, 267, 79, 80, 96,
	230, 188, 97, 112, 186, 92, 416, 81, 389, 222,
	223, 390, 78, 79, 80, 295, 89, 248, 249, 123,
	227, 92, 5, 413, 94, 95, 124, 81, 78, 94,
	95, 293, 352, 226, 353, 226, 372, 260, 132, 256,
	262, 255, 189, 79, 80, 269, 226, 286, 365, 265,
	226, 92, 296, 190, 122, 354, 252, 366, 78, 79,
	80, 284, 226, 79, 80, 133, 134, 92, 286, 131,
	287, 92, 286, 204, 78, 146, 138, 231, 78, 340,
	285, 112, 238, 440, 102, 437, 436, 387, 79, 80,
	388, 241, 196, 302, 307, 270, 92, 402, 161, 143,
	301, 307, 318, 78, 167, 314, 384, 322, 313, 172,
	315, 255, 368, 226, 177, 369, 226, 180, 181, 435,
	400, 437, 436, 226, 226, 382, 337, 333, 332, 344,
	378, 137, 356, 298, 233, 266, 377, 362, 364, 357,
	299, 355, 268, 375, 211, 212, 213, 214, 215, 216,
	217, 218, 219, 137, 274, 320, 367, 323, 259, 371,
	226, 258, 356, 370, 331, 81, 333, 332, 243, 357,
	367, 291, 277, 233, 292, 294, 226, 264, 265, 246,
	226, 247, 159, 100, 160, 99, 254, 220, 203, 182,
	162, 431, 62, 111, 237, 137, 273, 228, 236, 1,
	394, 79, 80, 321, 147, 392, 82, 83, 84, 92,
	52, 51, 307, 399, 50, 49, 78, 87, 85, 86,
	89, 226, 226, 48, 226, 47, 33, 32, 31, 30,
	43, 359, 226, 361, 226, 346, 347, 19, 20, 21,
	22, 289, 226, 309, 14, 12, 11, 311, 418, 419,
	420, 367, 226, 18, 362, 423, 10, 16, 226, 344,
	344, 344, 429, 13, 15, 29, 28, 24, 59, 34,
	25, 300, 439, 58, 308, 0, 0, 0, 316, 0,
	344, 308, 443, 444, 344, 344, 344, 445, 0, 0,
	0, 0, 0, 0, 0, 0, 57, 117, 63, 118,
	104, 334, 110, 115, 0, 0, 0, 0, 0, 345,
	0, 0, 0, 0, 0, 0, 0, 363, 0, 0,
	0, 0, 0, 0, 0, 0, 65, 0, 0, 72,
	73, 412, 0, 108, 66, 67, 68, 69, 0, 0,
	109, 0, 0, 0, 81, 0, 376, 0, 0, 421,
	0, 0, 107, 379, 116, 0, 71, 70, 88, 0,
	0, 433, 0, 40, 0, 0, 0, 0, 0, 385,
	386, 0, 441, 90, 91, 0, 0, 0, 0, 0,
	79, 80, 0, 0, 0, 82, 83, 84, 92, 396,
	397, 0, 308, 0, 0, 78, 87, 85, 86, 89,
	0, 0, 119, 0, 0, 0, 0, 0, 406, 407,
	408, 409, 0, 0, 0, 0, 0, 0, 119, 0,
	119, 0, 0, 0, 119, 0, 0, 0, 0, 0,
	119, 119, 119, 119, 422, 0, 0, 41, 0, 345,
	345, 345, 0, 119, 119, 119, 119, 438, 119, 119,
	119, 119, 0, 119, 119, 119, 119, 442, 119, 119,
	345, 0, 0, 0, 345, 345, 345, 0, 119, 0,
	0, 119, 119, 119, 0, 0, 120, 0, 0, 0,
	0, 119, 0, 0, 0, 0, 119, 0, 0, 0,
	0, 0, 120, 0, 120, 0, 0, 0, 120, 0,
	0, 0, 0, 0, 120, 120, 120, 120, 119, 0,
	119, 0, 0, 0, 0, 0, 0, 120, 120, 120,
	120, 0, 120, 120, 120, 120, 119, 120, 120, 120,
	120, 0, 120, 120, 0, 0, 57, 117, 63, 118,
	251, 0, 120, 115, 0, 120, 120, 120, 0, 0,
	0, 0, 0, 0, 0, 120, 0, 0, 0, 0,
	120, 0, 0, 0, 0, 0, 65, 0, 119, 72,
	73, 0, 0, 250, 66, 67, 68, 69, 0, 0,
	0, 0, 120, 0, 120, 0, 0, 0, 0, 0,
	0, 0, 60, 119, 116, 0, 71, 70, 0, 0,
	120, 0, 0, 0, 119, 0, 0, 0, 0, 0,
	57, 38, 63, 39, 64, 119, 119, 0, 35, 428,
	348, 427, 426, 349, 36, 37, 0, 55, 0, 46,
	119, 0, 350, 351, 0, 0, 0, 54, 0, 0,
	65, 56, 120, 72, 73, 0, 0, 0, 66, 67,
	68, 69, 0, 0, 119, 119, 342, 343, 0, 0,
	0, 119, 0, 0, 0, 0, 60, 120, 61, 0,
	71, 70, 0, 0, 0, 42, 119, 0, 120, 0,
	0, 0, 0, 0, 0, 119, 0, 0, 0, 120,
	120, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 120, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 121, 0, 0, 0, 119, 0,
	0, 0, 0, 119, 0, 0, 0, 0, 120, 120,
	121, 0, 121, 0, 0, 120, 121, 0, 0, 0,
	0, 0, 121, 121, 121, 121, 0, 0, 0, 27,
	120, 0, 0, 0, 0, 121, 121, 121, 121, 120,
	121, 121, 121, 121, 0, 121, 121, 121, 121, 0,
	121, 121, 0, 0, 0, 0, 0, 0, 119, 0,
	121, 0, 0, 121, 121, 121, 209, 0, 114, 0,
	0, 0, 120, 121, 0, 0, 0, 120, 121, 0,
	0, 0, 0, 0, 114, 0, 114, 0, 0, 0,
	114, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	121, 0, 121, 0, 0, 0, 0, 0, 0, 114,
	114, 114, 114, 135, 114, 114, 114, 114, 121, 114,
	114, 114, 114, 0, 114, 114, 0, 0, 0, 0,
	0, 0, 120, 0, 114, 0, 0, 114, 114, 114,
	0, 0, 0, 0, 0, 0, 0, 114, 0, 0,
	0, 0, 114, 0, 0, 0, 0, 0, 0, 0,
	121, 0, 0, 0, 0, 0, 0, 0, 192, 0,
	194, 0, 0, 0, 114, 0, 114, 0, 0, 0,
	0, 0, 205, 0, 0, 121, 0, 0, 0, 0,
	0, 0, 114, 0, 0, 0, 121, 0, 0, 0,
	221, 0, 57, 38, 63, 39, 64, 121, 121, 0,
	35, 424, 348, 427, 426, 349, 36, 37, 0, 55,
	0, 46, 121, 0, 350, 351, 0, 0, 0, 54,
	0, 0, 65, 56, 114, 72, 73, 0, 0, 0,
	66, 67, 68, 69, 0, 0, 121, 121, 342, 343,
	0, 0, 0, 121, 0, 0, 0, 0, 60, 114,
	61, 0, 71, 70, 0, 0, 0, 0, 121, 0,
	114, 0, 0, 272, 0, 275, 0, 121, 0, 0,
	0, 114, 114, 0, 0, 0, 0, 0, 0, 282,
	283, 0, 0, 0, 0, 0, 114, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	121, 0, 0, 0, 0, 121, 0, 0, 0, 0,
	114, 114, 0, 0, 0, 0, 0, 114, 0, 319,
	0, 0, 0, 0, 57, 117, 63, 118, 64, 0,
	0, 115, 114, 0, 0, 0, 0, 0, 335, 0,
	0, 114, 0, 339, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 65, 0, 0, 72, 73, 0,
	121, 0, 66, 67, 68, 69, 0, 0, 0, 0,
	235, 0, 0, 0, 114, 0, 0, 324, 0, 114,
	60, 0, 116, 0, 71, 70, 0, 0, 0, 0,
	380, 381, 57, 38, 63, 39, 64, 383, 0, 0,
	35, 430, 348, 0, 0, 349, 36, 37, 0, 55,
	391, 46, 393, 0, 350, 351, 0, 0, 0, 54,
	0, 0, 65, 56, 0, 72, 73, 0, 0, 0,
	66, 67, 68, 69, 114, 0, 401, 0, 342, 343,
	0, 0, 221, 0, 0, 0, 0, 0, 60, 411,
	61, 0, 71, 70, 0, 0, 0, 0, 0, 417,
	0, 282, 283, 57, 38, 63, 39, 64, 0, 0,
	0, 35, 398, 44, 0, 0, 45, 36, 37, 0,
	55, 0, 46, 286, 0, 0, 0, 0, 0, 312,
	54, 0, 0, 65, 56, 0, 72, 73, 0, 0,
	0, 66, 67, 68, 69, 0, 0, 0, 0, 305,
	306, 0, 0, 0, 0, 0, 0, 0, 0, 60,
	0, 61, 0, 71, 70, 57, 38, 63, 39, 64,
	0, 0, 0, 35, 341, 348, 0, 0, 349, 36,
	37, 0, 55, 0, 46, 0, 0, 350, 351, 0,
	0, 0, 54, 0, 0, 65, 56, 0, 72, 73,
	0, 0, 0, 66, 67, 68, 69, 0, 0, 0,
	0, 342, 343, 0, 0, 0, 0, 0, 0, 0,
	0, 60, 0, 61, 0, 71, 70, 57, 38, 63,
	39, 64, 0, 0, 0, 35, 336, 44, 281, 280,
	45, 36, 37, 0, 55, 0, 46, 0, 0, 0,
	0, 0, 0, 0, 54, 0, 0, 65, 56, 0,
	72, 73, 0, 0, 0, 66, 67, 68, 69, 0,
	0, 0, 0, 224, 225, 0, 0, 0, 0, 0,
	0, 0, 0, 60, 0, 61, 0, 71, 70, 57,
	38, 63, 39, 64, 0, 0, 0, 35, 317, 44,
	0, 0, 45, 36, 37, 0, 55, 0, 46, 286,
	0, 0, 0, 0, 0, 312, 54, 0, 0, 65,
	56, 0, 72, 73, 0, 0, 0, 66, 67, 68,
	69, 0, 0, 0, 0, 305, 306, 0, 0, 0,
	0, 0, 0, 0, 0, 60, 0, 61, 0, 71,
	70, 57, 38, 63, 39, 64, 0, 0, 0, 35,
	303, 44, 0, 0, 45, 36, 37, 0, 55, 0,
	46, 286, 0, 0, 0, 0, 0, 312, 54, 0,
	0, 65, 56, 0, 72, 73, 0, 0, 0, 66,
	67, 68, 69, 0, 0, 0, 0, 305, 306, 0,
	0, 0, 0, 0, 0, 0, 0, 60, 0, 61,
	0, 71, 70, 57, 38, 63, 39, 64, 0, 0,
	0, 35, 278, 44, 281, 280, 45, 36, 37, 0,
	55, 0, 46, 0, 0, 0, 0, 0, 0, 0,
	54, 0, 0, 65, 56, 0, 72, 73, 0, 0,
	0, 66, 67, 68, 69, 0, 0, 0, 0, 224,
	225, 0, 57, 38, 63, 39, 64, 0, 0, 60,
	35, 61, 44, 71, 70, 45, 36, 37, 0, 55,
	0, 46, 0, 0, 0, 0, 0, 0, 0, 54,
	0, 0, 65, 56, 0, 72, 73, 0, 0, 0,
	66, 67, 68, 69, 0, 0, 0, 0, 6, 7,
	0, 0, 0, 0, 0, 0, 0, 0, 60, 0,
	61, 0, 71, 70, 0, 8, 57, 38, 63, 39,
	64, 0, 0, 0, 35, 0, 348, 0, 0, 349,
	36, 37, 0, 55, 0, 46, 0, 0, 350, 351,
	0, 0, 0, 54, 0, 0, 65, 56, 0, 72,
	73, 0, 0, 0, 66, 67, 68, 69, 0, 0,
	0, 0, 342, 343, 0, 0, 0, 0, 0, 0,
	0, 0, 60, 0, 61, 0, 71, 70, 57, 38,
	63, 39, 64, 0, 0, 0, 35, 410, 44, 0,
	0, 45, 36, 37, 0, 55, 0, 46, 0, 0,
	0, 0, 0, 0, 0, 54, 0, 0, 65, 56,
	0, 72, 73, 0, 0, 0, 66, 67, 68, 69,
	0, 0, 0, 0, 224, 225, 0, 0, 0, 0,
	0, 0, 0, 0, 60, 0, 61, 0, 71, 70,
	57, 38, 63, 39, 64, 0, 0, 0, 35, 373,
	44, 0, 0, 45, 36, 37, 0, 55, 0, 46,
	0, 0, 0, 0, 0, 0, 0, 54, 0, 0,
	65, 56, 0, 72, 73, 0, 0, 0, 66, 67,
	68, 69, 0, 0, 0, 0, 224, 225, 0, 0,
	0, 0, 0, 0, 0, 0, 60, 0, 61, 0,
	71, 70, 57, 38, 63, 39, 64, 0, 0, 0,
	35, 338, 44, 0, 0, 45, 36, 37, 0, 55,
	0, 46, 0, 0, 0, 0, 0, 0, 0, 54,
	0, 0, 65, 56, 0, 72, 73, 0, 0, 0,
	66, 67, 68, 69, 0, 0, 0, 0, 224, 225,
	0, 57, 38, 63, 39, 64, 0, 0, 60, 35,
	61, 44, 71, 70, 45, 36, 37, 0, 55, 0,
	46, 0, 0, 0, 0, 0, 0, 0, 54, 0,
	0, 65, 56, 0, 72, 73, 0, 0, 0, 66,
	67, 68, 69, 0, 0, 0, 0, 224, 225, 0,
	0, 0, 0, 0, 0, 0, 0, 60, 0, 61,
	330, 71, 70, 57, 38, 63, 39, 64, 0, 0,
	0, 35, 327, 44, 0, 0, 45, 36, 37, 0,
	55, 0, 46, 0, 0, 0, 0, 0, 0, 0,
	54, 0, 0, 65, 56, 0, 72, 73, 0, 0,
	0, 66, 67, 68, 69, 0, 0, 0, 0, 224,
	225, 0, 57, 38, 63, 39, 64, 0, 0, 60,
	35, 61, 44, 71, 70, 45, 36, 37, 0, 55,
	0, 46, 0, 0, 0, 0, 0, 0, 0, 54,
	0, 0, 65, 56, 0, 72, 73, 0, 0, 0,
	66, 67, 68, 69, 0, 0, 0, 0, 224, 225,
	0, 0, 0, 0, 0, 0, 0, 0, 60, 0,
	61, 276, 71, 70, 57, 38, 63, 39, 64, 0,
	0, 0, 35, 271, 44, 0, 0, 45, 36, 37,
	0, 55, 0, 46, 0, 0, 0, 0, 0, 0,
	0, 54, 0, 0, 65, 56, 0, 72, 73, 0,
	0, 0, 66, 67, 68, 69, 0, 0, 0, 0,
	224, 225, 0, 0, 0, 0, 0, 0, 0, 0,
	60, 0, 61, 0, 71, 70, 57, 38, 63, 39,
	64, 0, 0, 0, 35, 263, 44, 0, 0, 45,
	36, 37, 0, 55, 0, 46, 0, 0, 0, 0,
	0, 0, 0, 54, 0, 0, 65, 56, 0, 72,
	73, 0, 0, 0, 66, 67, 68, 69, 0, 0,
	0, 0, 224, 225, 0, 0, 0, 0, 0, 0,
	0, 0, 60, 0, 61, 0, 71, 70, 57, 38,
	63, 39, 64, 0, 0, 0, 35, 261, 44, 0,
	0, 45, 36, 37, 0, 55, 0, 46, 0, 0,
	0, 0, 0, 0, 0, 54, 0, 0, 65, 56,
	0, 72, 73, 0, 0, 0, 66, 67, 68, 69,
	0, 0, 0, 0, 224, 225, 0, 57, 38, 63,
	39, 64, 0, 0, 60, 35, 61, 44, 71, 70,
	45, 36, 37, 0, 55, 0, 46, 0, 0, 0,
	0, 0, 0, 0, 54, 0, 0, 65, 56, 0,
	72, 73, 0, 0, 0, 66, 67, 68, 69, 0,
	0, 0, 0, 224, 225, 0, 57, 38, 63, 39,
	64, 245, 0, 60, 35, 61, 44, 71, 70, 45,
	36, 37, 0, 55, 0, 46, 0, 0, 0, 0,
	0, 0, 0, 54, 0, 0, 65, 56, 0, 72,
	73, 0, 0, 0, 66, 67, 68, 69, 0, 0,
	0, 0, 0, 244, 0, 57, 38, 63, 39, 64,
	0, 0, 60, 35, 61, 44, 71, 70, 45, 36,
	37, 0, 55, 0, 46, 0, 0, 0, 0, 0,
	0, 0, 54, 0, 0, 65, 56, 0, 72, 73,
	0, 0, 0, 66, 67, 68, 69, 57, 117, 63,
	118, 64, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 60, 0, 61, 0, 71, 70, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 65, 0, 0,
	72, 73, 0, 0, 0, 66, 67, 68, 69, 0,
	0, 0, 0, 235, 57, 117, 63, 118, 104, 0,
	374, 115, 0, 60, 0, 61, 0, 71, 70, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 65, 0, 0, 72, 73, 0,
	0, 0, 66, 67, 68, 69, 57, 117, 63, 118,
	64, 0, 297, 0, 0, 0, 0, 0, 0, 0,
	107, 0, 116, 0, 71, 70, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 65, 0, 0, 72,
	73, 0, 0, 0, 66, 67, 68, 69, 0, 0,
	0, 0, 235, 57, 117, 63, 118, 64, 0, 0,
	115, 0, 60, 0, 61, 0, 71, 70, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 0, 0, 72, 73, 0, 0,
	0, 66, 67, 68, 69, 57, 240, 63, 118, 64,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 60,
	0, 116, 0, 71, 70, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 0, 72, 73,
	0, 0, 0, 66, 67, 68, 69, 0, 0, 0,
	0, 235, 57, 117, 63, 118, 104, 0, 0, 115,
	0, 60, 0, 61, 0, 71, 70, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 0, 72, 73, 0, 0, 0,
	66, 67, 68, 69, 57, 117, 63, 118, 64, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 107, 0,
	116, 0, 71, 70, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 65, 0, 0, 72, 73, 0,
	0, 0, 66, 67, 68, 69, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	60, 0, 61, 0, 71, 70,
}
var RubyPact = []int{

	-28, 1707, -1000, -1000, -1000, 7, -1000, -1000, -1000, 500,
	-1000, -1000, -1000, 88, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 144, -27, -1000,
	-1000, -1000, -1000, -1000, -1000, 339, 236, 236, 451, 137,
	22, 20, 5, 187, 2450, 2450, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 2759, 2450, 2759, -1000, -1000, -1000,
	2759, -1000, -18, 226, -1000, 14, 13, 13, 13, 13,
	23, 336, -1000, -1000, -1000, -1000, -1000, 2450, 344, 2759,
	2759, 2759, 2759, 2450, 2759, 2759, 2759, 2759, 2450, 2759,
	2759, 2759, 2759, 2450, 2759, 2759, 2450, 2450, 343, 64,
	155, 32, -1000, -1000, 2759, -1000, 132, 2759, 2759, 2759,
	342, 222, 183, -1000, 149, -21, -21, 2717, 205, -1000,
	-1000, -1000, 2759, 2450, 2450, 2450, 2450, 2450, 2450, 2450,
	2450, 2450, 341, 105, 118, 2352, 127, 183, 109, 183,
	45, 42, 321, -1000, 2670, 243, 2759, 2401, -1000, -21,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 333, -1000,
	-1000, 105, 691, 183, 183, 183, 183, 105, 183, 183,
	183, 183, 105, 163, 183, 183, 183, 105, 183, 183,
	105, 105, -1000, -1000, 340, 190, 70, -1000, 3, 315,
	312, -1000, 2303, 236, 2241, 327, 321, -1000, 94, 117,
	-1000, 117, -1000, -1000, 2628, 2179, -1000, 308, -1000, 2117,
	322, 105, 105, 105, 105, 105, 105, 105, 105, 105,
	-1000, 1658, -1000, -1000, -1000, -1000, 500, 105, 207, 2759,
	-1000, 9, -1000, -1000, -1000, -1000, 180, 164, -10, 208,
	2539, -1000, 283, 105, -1000, -1000, -1000, -1000, 132, 14,
	2450, 2759, 2759, 1596, 155, 70, 260, 2450, -1000, -1000,
	1534, -1000, -1000, -1000, 14, -1000, 1209, 0, -24, 183,
	-1000, -1000, 2068, 38, -1000, 2006, -1000, -1000, -1000, 310,
	2450, -1000, 1472, 1957, -1000, -1000, 231, 183, 1410, 178,
	2759, 2581, -26, -1000, -31, -1000, 2450, 2759, -1000, -1000,
	105, 198, 183, -1000, 203, -1000, -1000, 500, 105, -1000,
	-1000, 257, 2759, -1000, -1000, -1000, 105, -1000, 182, 1895,
	-1000, 2492, 183, -1000, 297, 2450, 290, -1000, -1000, 284,
	-1000, -1000, 2450, -1000, 105, 2352, -1000, 271, -1000, 2352,
	262, -1000, -1000, -1000, 500, 105, -1000, -1000, 2450, 2450,
	232, 153, -1000, -1000, 2759, 127, 321, -1000, -1000, 2581,
	-1000, 97, 500, 105, 183, -1000, -1000, -1000, 2450, 2450,
	127, 1348, -1000, -1000, 274, -1000, 105, -1000, -1000, 105,
	2352, 2352, -1000, 2352, 251, 33, 15, 2450, 2450, 2450,
	2450, 1833, 127, 2352, 179, -16, 105, 105, -1000, 49,
	156, 2352, -1000, -1000, -1000, -1000, 105, 105, 105, 105,
	-1000, 2352, -10, 2450, 2759, -1000, -1000, 2352, 1077, 765,
	1277, -10, 114, 133, -1000, 265, 2450, -1000, -1000, 229,
	-1000, -1000, -1000, -10, -1000, -1000, 2450, -1000, 105, 1771,
	-1000, -10, 105, 1771, 1771, 1771,
}
var RubyPgo = []int{

	0, 180, 433, 430, 429, 50, 428, 427, 426, 835,
	425, 11, 17, 424, 423, 417, 0, 32, 597, 909,
	416, 413, 407, 7, 406, 9, 523, 405, 404, 6,
	403, 401, 400, 399, 398, 397, 396, 395, 2, 390,
	389, 388, 387, 386, 385, 383, 375, 374, 371, 370,
	946, 364, 1, 15, 16, 3, 359, 13, 358, 4,
	357, 8, 5, 356, 10, 354, 353, 14, 12, 352,
	351, 97,
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
	19, 53, 53, 53, 53, 64, 64, 61, 61, 61,
	61, 61, 67, 67, 67, 67, 67, 66, 66, 66,
	20, 20, 20, 20, 20, 20, 30, 30, 30, 30,
	62, 62, 62, 62, 62, 62, 57, 57, 25, 25,
	25, 25, 68, 68, 68, 24, 24, 27, 29, 29,
	69, 69, 14, 14, 14, 14, 14, 14, 14, 28,
	28, 28, 28, 28, 28, 9, 9, 26, 26, 18,
	18, 39, 39, 40, 41, 42, 43, 44, 45, 46,
	47, 48, 49, 2, 6, 7, 7, 3, 3, 58,
	58, 58, 58, 65, 65, 65, 5, 5, 5, 5,
	54, 63, 63, 63, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 55, 55, 55, 55, 51, 51,
	51, 8, 15, 11, 11, 11, 60, 60, 52, 52,
	21, 22, 12, 35, 59, 59, 59, 59, 59, 59,
	59, 36, 36, 36, 36, 36, 36, 36, 37, 37,
	37, 37, 37, 38, 38, 38, 38, 34, 33, 10,
	32, 32, 31, 31, 4,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 1, 0, 2, 0, 2, 2, 2, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 4, 5, 1, 4, 4, 2, 3, 4, 4,
	5, 3, 5, 2, 3, 3, 3, 3, 3, 4,
	6, 3, 7, 1, 5, 1, 3, 0, 1, 1,
	4, 4, 1, 1, 4, 4, 5, 1, 3, 3,
	5, 6, 7, 8, 5, 6, 0, 1, 3, 3,
	0, 2, 2, 2, 2, 2, 1, 3, 1, 2,
	3, 2, 0, 1, 3, 4, 6, 4, 1, 3,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 2, 2, 2, 3,
	3, 3, 3, 2, 2, 2, 2, 3, 3, 3,
	3, 3, 3, 1, 1, 3, 3, 5, 5, 0,
	1, 3, 7, 3, 7, 8, 3, 4, 4, 3,
	3, 0, 1, 3, 4, 5, 3, 3, 3, 3,
	3, 5, 6, 5, 4, 3, 3, 2, 0, 2,
	2, 3, 4, 2, 3, 5, 0, 2, 1, 2,
	2, 2, 5, 5, 0, 2, 2, 2, 2, 2,
	2, 0, 1, 3, 3, 1, 3, 3, 5, 6,
	5, 6, 5, 4, 3, 3, 2, 3, 3, 2,
	5, 7, 4, 5, 3,
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
	33, 34, 48, 18, 40, 41, 15, 18, 55, 6,
	4, -29, 8, -29, 9, -53, -67, 61, 42, 49,
	11, -66, -16, -5, -19, 12, 63, 6, 8, -26,
	-18, -9, 9, 42, 49, 42, 49, 42, 49, 42,
	49, 42, 11, -1, -1, -50, -64, -16, -1, -16,
	-64, -61, -16, -23, -71, 54, 9, -51, -5, 63,
	-17, 6, 8, -17, -17, -17, 6, 8, 66, 6,
	8, -1, 6, -16, -16, -16, -16, -1, -16, -16,
	-16, -16, -1, -16, -16, -16, -16, -1, -16, -16,
	-1, -1, 6, -57, 55, -68, 9, -25, 6, 47,
	58, -57, -50, 40, -50, -61, -16, 11, -16, -16,
	-12, -16, -12, 6, 11, -50, -54, 56, -54, -50,
	-61, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	6, -50, 51, 52, 51, 52, -16, -1, -60, 11,
	51, -71, 62, 11, 62, 51, -58, -65, -71, -16,
	6, 8, -61, -1, 52, 10, 6, 8, -67, -53,
	42, 9, 53, -62, 6, 11, -68, 42, 6, 6,
	-62, 14, -29, 14, 10, 11, -71, 62, -71, -16,
	-5, 14, -50, -63, 6, -50, 64, 10, 14, -55,
	17, 16, -50, -50, 14, -11, 25, -16, -59, -31,
	37, -71, -71, 11, -71, 11, 4, 53, 10, -5,
	-1, -61, -16, 14, -52, 51, 52, -16, -1, -30,
	-11, -22, 31, -57, -25, 10, -1, 14, -52, -50,
	-5, -71, -16, -5, 58, 42, 58, 14, 56, 11,
	64, 14, 17, 16, -1, -50, 14, -55, 14, -50,
	8, 14, 51, 52, -16, -1, -37, -36, 15, 18,
	27, 28, 14, 16, 37, -64, -16, -23, 64, -71,
	64, -71, -16, -1, -16, 10, 14, -11, 15, 18,
	-64, -62, 14, 14, 58, 6, -1, 6, 6, -1,
	-50, -50, 14, -50, 4, -1, -1, 15, 18, 15,
	18, -50, -64, -50, -16, 6, -1, -1, 14, -52,
	6, -50, 6, 51, 51, 52, -1, -1, -1, -1,
	14, -50, -71, 4, 53, 14, 10, -50, -59, -59,
	-59, -71, -1, -16, 14, -38, 17, 16, 14, -38,
	14, -70, 11, -71, 11, 14, 17, 16, -1, -59,
	14, -71, -1, -59, -59, -59,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 47,
	48, 49, 50, 51, 52, 53, 54, 55, 56, 57,
	58, 59, 60, 28, 29, 30, 31, 32, 33, 34,
	35, 36, 37, 38, 39, 0, 0, 0, 21, -2,
	23, 24, 25, 0, 0, 0, 15, 40, 41, 42,
	43, 44, 45, 46, 0, 0, 0, 20, 26, 27,
	87, 13, 0, 63, 198, 0, 0, 0, 0, 0,
	0, 0, 163, 164, 5, 6, 7, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 122,
	122, 15, -2, 15, -2, 66, 73, 87, 0, 0,
	0, 83, 92, 93, 32, 15, -2, 21, -2, 23,
	24, 25, 87, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 15, 0, 206, 210, 85, 0, 13,
	0, 0, 85, 89, 13, 0, 87, 0, 239, 15,
	153, 21, 22, 154, 155, 156, 147, 148, 0, 145,
	146, 186, 67, 74, 76, 78, 157, 158, 159, 160,
	161, 162, 188, 0, 237, 238, 244, 190, 75, 77,
	187, 189, 71, 110, 0, 116, 122, 123, 118, 0,
	0, 110, 0, 0, 0, 0, 88, 13, 85, 132,
	133, 139, 140, 151, 13, 0, 15, 181, 15, 0,
	0, 134, 141, 135, 142, 136, 143, 137, 144, 138,
	152, 0, 15, 15, 16, 17, 18, 19, 0, 0,
	214, 0, 165, 13, 166, 14, 13, 13, 170, 0,
	21, -2, 0, 199, 200, 201, 149, 150, 68, 69,
	0, -2, 0, 0, 122, 0, 0, 0, 119, 121,
	0, 125, 15, 127, 61, 13, 0, 79, 0, 98,
	99, 176, 0, 0, 182, 0, 179, 65, 184, 0,
	0, 15, 0, 0, 202, 207, 15, 86, 0, 0,
	0, 0, 0, 13, 0, 13, 0, 0, 64, 70,
	72, 0, 212, 100, 0, 111, 112, 47, 114, 115,
	208, 107, 0, 110, 124, 117, 120, 104, 0, 0,
	62, 0, 94, 95, 0, 0, 0, 177, 180, 0,
	178, 185, 0, 15, 15, 197, 191, 0, 193, 203,
	15, 213, 215, 216, 47, 218, 219, 220, 0, 0,
	222, 225, 240, 15, 0, 15, 90, 91, 167, 0,
	168, 0, 47, 171, 173, 81, 101, 209, 0, 0,
	211, 0, 105, 126, 0, 96, 80, 84, 183, 15,
	195, 196, 192, 204, 0, 15, 0, 0, 0, 0,
	0, 0, 15, 13, 0, 0, 108, 109, 102, 0,
	0, 194, 15, 214, -2, -2, 223, 224, 226, 227,
	241, 13, 242, 0, 0, 103, 82, 205, 0, 0,
	0, 243, 11, 13, 228, 0, 0, 214, 230, 0,
	232, 172, 12, 174, 13, 229, 0, 214, 214, 221,
	231, 175, 214, 221, 221, 221,
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
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 78:
		//line parser.y:375
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 79:
		//line parser.y:385
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 80:
		//line parser.y:395
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 81:
		//line parser.y:404
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 82:
		//line parser.y:406
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 83:
		//line parser.y:408
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 84:
		//line parser.y:410
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 85:
		//line parser.y:413
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 86:
		//line parser.y:415
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 87:
		//line parser.y:417
		{
			RubyVAL.genericSlice = ast.Nodes{}
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
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 91:
		//line parser.y:425
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 92:
		//line parser.y:429
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
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
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-1].genericValue)
		}
	case 96:
		//line parser.y:437
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 97:
		//line parser.y:440
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 98:
		//line parser.y:442
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 99:
		//line parser.y:444
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 100:
		//line parser.y:450
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 101:
		//line parser.y:458
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 102:
		//line parser.y:467
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 103:
		//line parser.y:476
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 104:
		//line parser.y:486
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 105:
		//line parser.y:494
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 106:
		//line parser.y:504
		{
		}
	case 107:
		//line parser.y:506
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 108:
		//line parser.y:508
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{RubyS[Rubypt-2].genericValue}}
		}
	case 109:
		//line parser.y:510
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{ast.Break{}},
			}
		}
	case 110:
		//line parser.y:518
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 111:
		//line parser.y:520
		{
		}
	case 112:
		//line parser.y:522
		{
		}
	case 113:
		//line parser.y:524
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 114:
		//line parser.y:526
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 115:
		//line parser.y:528
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 116:
		//line parser.y:531
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 117:
		//line parser.y:533
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
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
		//line parser.y:544
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 123:
		//line parser.y:546
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 124:
		//line parser.y:550
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 125:
		//line parser.y:555
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 126:
		//line parser.y:562
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 127:
		//line parser.y:572
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 128:
		//line parser.y:581
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 129:
		//line parser.y:587
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 130:
		//line parser.y:595
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 131:
		//line parser.y:599
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 132:
		//line parser.y:604
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 133:
		//line parser.y:611
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 134:
		//line parser.y:615
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 135:
		//line parser.y:622
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 136:
		//line parser.y:629
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 137:
		//line parser.y:636
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 138:
		//line parser.y:643
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 139:
		//line parser.y:651
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 140:
		//line parser.y:658
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 141:
		//line parser.y:662
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 142:
		//line parser.y:669
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 143:
		//line parser.y:676
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 144:
		//line parser.y:683
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 145:
		//line parser.y:691
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 146:
		//line parser.y:693
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 147:
		//line parser.y:696
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 148:
		//line parser.y:698
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 149:
		//line parser.y:701
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 150:
		//line parser.y:703
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 151:
		//line parser.y:706
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 152:
		//line parser.y:708
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 153:
		//line parser.y:710
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 154:
		//line parser.y:711
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 155:
		//line parser.y:712
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 156:
		//line parser.y:713
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 157:
		//line parser.y:716
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 158:
		//line parser.y:725
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 159:
		//line parser.y:734
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 160:
		//line parser.y:743
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 161:
		//line parser.y:752
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 162:
		//line parser.y:761
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 163:
		//line parser.y:769
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 164:
		//line parser.y:770
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 165:
		//line parser.y:772
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 166:
		//line parser.y:774
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
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
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 169:
		//line parser.y:793
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 170:
		//line parser.y:795
		{
		}
	case 171:
		//line parser.y:797
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 172:
		//line parser.y:804
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 173:
		//line parser.y:812
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 174:
		//line parser.y:819
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 175:
		//line parser.y:826
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 176:
		//line parser.y:834
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 177:
		//line parser.y:836
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 178:
		//line parser.y:843
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 179:
		//line parser.y:850
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 180:
		//line parser.y:853
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 181:
		//line parser.y:855
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 182:
		//line parser.y:857
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 183:
		//line parser.y:859
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 184:
		//line parser.y:862
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 185:
		//line parser.y:869
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 186:
		//line parser.y:877
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 187:
		//line parser.y:884
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 188:
		//line parser.y:891
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 189:
		//line parser.y:898
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 190:
		//line parser.y:905
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 191:
		//line parser.y:912
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 192:
		//line parser.y:919
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 193:
		//line parser.y:927
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 194:
		//line parser.y:936
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 195:
		//line parser.y:943
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 196:
		//line parser.y:950
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 197:
		//line parser.y:957
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 198:
		//line parser.y:964
		{
		}
	case 199:
		//line parser.y:965
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 200:
		//line parser.y:966
		{
		}
	case 201:
		//line parser.y:969
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 202:
		//line parser.y:972
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 203:
		//line parser.y:980
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 204:
		//line parser.y:982
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 205:
		//line parser.y:991
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
	case 206:
		//line parser.y:1006
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 207:
		//line parser.y:1008
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 208:
		//line parser.y:1011
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 209:
		//line parser.y:1013
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 210:
		//line parser.y:1016
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 211:
		//line parser.y:1025
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 212:
		//line parser.y:1034
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 213:
		//line parser.y:1043
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 214:
		//line parser.y:1046
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 215:
		//line parser.y:1048
		{
		}
	case 216:
		//line parser.y:1050
		{
		}
	case 217:
		//line parser.y:1052
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 218:
		//line parser.y:1054
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 219:
		//line parser.y:1056
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 220:
		//line parser.y:1058
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 221:
		//line parser.y:1060
		{
		}
	case 222:
		//line parser.y:1062
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 223:
		//line parser.y:1064
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 224:
		//line parser.y:1066
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 225:
		//line parser.y:1068
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 226:
		//line parser.y:1070
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 227:
		//line parser.y:1072
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 228:
		//line parser.y:1075
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 229:
		//line parser.y:1082
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 230:
		//line parser.y:1090
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 231:
		//line parser.y:1097
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 232:
		//line parser.y:1105
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 233:
		//line parser.y:1113
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 234:
		//line parser.y:1120
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 235:
		//line parser.y:1127
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 236:
		//line parser.y:1134
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 237:
		//line parser.y:1142
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 238:
		//line parser.y:1145
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 239:
		//line parser.y:1147
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 240:
		//line parser.y:1150
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 241:
		//line parser.y:1152
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 242:
		//line parser.y:1155
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 243:
		//line parser.y:1157
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 244:
		//line parser.y:1159
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
