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

//line parser.y:1249

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 42,
	54, 134,
	-2, 21,
	-1, 110,
	54, 134,
	-2, 132,
	-1, 112,
	10, 99,
	11, 99,
	-2, 207,
	-1, 124,
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
	-1, 126,
	54, 134,
	-2, 21,
	-1, 266,
	54, 135,
	-2, 133,
	-1, 276,
	10, 99,
	11, 99,
	-2, 207,
	-1, 441,
	27, 225,
	28, 225,
	-2, 15,
	-1, 442,
	27, 225,
	28, 225,
	-2, 15,
}

const RubyNprod = 255
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 3317

var RubyAct = []int{

	251, 317, 5, 464, 326, 344, 343, 202, 152, 198,
	23, 119, 27, 57, 200, 47, 114, 22, 113, 2,
	3, 83, 228, 300, 84, 89, 168, 260, 169, 260,
	123, 153, 17, 437, 315, 314, 4, 121, 213, 123,
	396, 104, 394, 179, 105, 282, 133, 102, 282, 144,
	145, 227, 109, 111, 98, 99, 257, 81, 80, 203,
	149, 87, 88, 106, 260, 253, 102, 147, 103, 158,
	151, 358, 163, 164, 82, 134, 86, 451, 361, 328,
	260, 162, 135, 462, 173, 174, 170, 103, 83, 356,
	162, 84, 181, 260, 133, 208, 301, 186, 155, 281,
	204, 364, 191, 341, 433, 195, 196, 258, 161, 101,
	157, 205, 289, 130, 43, 193, 256, 83, 206, 203,
	84, 253, 201, 360, 440, 203, 216, 217, 201, 357,
	220, 222, 232, 233, 219, 235, 236, 237, 238, 239,
	240, 241, 188, 189, 210, 155, 226, 83, 155, 260,
	84, 212, 214, 441, 442, 357, 127, 333, 243, 452,
	204, 268, 231, 155, 130, 138, 204, 143, 157, 331,
	324, 205, 139, 127, 199, 127, 127, 205, 136, 143,
	127, 287, 141, 247, 248, 137, 356, 224, 127, 127,
	127, 267, 155, 273, 255, 274, 83, 131, 142, 84,
	127, 263, 127, 127, 132, 127, 159, 127, 127, 127,
	127, 376, 127, 140, 89, 127, 288, 127, 127, 83,
	421, 294, 84, 387, 89, 388, 405, 127, 402, 110,
	127, 127, 127, 400, 297, 284, 322, 324, 266, 324,
	127, 216, 217, 98, 99, 127, 389, 324, 127, 89,
	87, 88, 471, 98, 99, 254, 83, 298, 323, 84,
	87, 88, 304, 155, 479, 86, 476, 475, 302, 97,
	426, 127, 127, 427, 127, 86, 338, 424, 98, 99,
	425, 474, 415, 476, 475, 87, 88, 419, 470, 369,
	368, 127, 347, 287, 127, 346, 345, 336, 257, 350,
	86, 329, 439, 127, 127, 330, 332, 414, 339, 155,
	412, 214, 337, 367, 411, 369, 368, 312, 257, 370,
	304, 410, 373, 89, 296, 297, 271, 380, 272, 171,
	108, 172, 107, 408, 352, 398, 308, 390, 291, 127,
	290, 353, 401, 286, 245, 127, 244, 242, 223, 403,
	197, 176, 98, 99, 67, 262, 403, 409, 307, 87,
	88, 252, 392, 261, 1, 413, 89, 160, 127, 416,
	56, 55, 54, 53, 86, 395, 52, 397, 51, 34,
	127, 313, 33, 32, 422, 423, 392, 31, 46, 381,
	382, 127, 19, 127, 36, 98, 99, 127, 429, 37,
	20, 327, 87, 88, 127, 14, 12, 90, 91, 92,
	100, 435, 11, 127, 21, 18, 10, 86, 95, 93,
	94, 97, 24, 16, 363, 443, 444, 445, 446, 127,
	127, 13, 35, 15, 30, 29, 25, 64, 26, 63,
	0, 403, 0, 127, 127, 456, 457, 458, 0, 0,
	127, 460, 28, 0, 0, 0, 0, 380, 380, 380,
	0, 468, 0, 0, 0, 0, 477, 0, 127, 0,
	89, 478, 0, 0, 449, 0, 481, 0, 0, 380,
	0, 482, 483, 380, 380, 380, 484, 0, 0, 0,
	0, 85, 459, 0, 122, 0, 0, 0, 0, 98,
	99, 0, 0, 0, 127, 472, 87, 88, 127, 0,
	127, 122, 0, 122, 122, 0, 0, 480, 122, 0,
	0, 86, 127, 0, 0, 97, 122, 122, 122, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 122, 0,
	122, 122, 0, 122, 450, 122, 122, 122, 122, 0,
	122, 127, 127, 122, 0, 122, 122, 0, 0, 0,
	62, 125, 68, 126, 69, 122, 127, 0, 122, 122,
	122, 0, 0, 98, 99, 0, 0, 0, 122, 0,
	87, 88, 0, 122, 0, 0, 122, 0, 0, 0,
	70, 0, 0, 78, 79, 86, 44, 89, 71, 72,
	73, 74, 75, 0, 473, 0, 260, 0, 0, 122,
	122, 0, 122, 407, 0, 89, 65, 89, 66, 0,
	77, 76, 0, 0, 0, 0, 98, 99, 0, 122,
	0, 0, 122, 87, 88, 0, 0, 0, 128, 0,
	0, 122, 122, 0, 98, 99, 98, 99, 86, 0,
	0, 87, 88, 87, 88, 128, 0, 128, 128, 0,
	0, 0, 128, 0, 0, 0, 86, 0, 86, 0,
	128, 128, 128, 283, 89, 0, 0, 122, 0, 0,
	0, 0, 128, 122, 128, 128, 0, 128, 0, 128,
	128, 128, 128, 0, 128, 0, 0, 128, 0, 128,
	128, 0, 0, 98, 99, 0, 122, 0, 0, 128,
	87, 88, 128, 128, 128, 0, 0, 0, 122, 0,
	0, 0, 128, 278, 0, 86, 0, 128, 0, 122,
	128, 122, 0, 0, 0, 122, 0, 0, 0, 0,
	0, 334, 122, 0, 0, 0, 0, 0, 0, 0,
	0, 122, 0, 128, 128, 0, 128, 62, 125, 68,
	126, 112, 0, 118, 123, 0, 0, 122, 122, 0,
	98, 99, 0, 128, 0, 0, 128, 87, 88, 0,
	0, 122, 122, 0, 0, 128, 128, 70, 122, 0,
	78, 79, 86, 0, 116, 71, 72, 73, 74, 75,
	0, 117, 0, 0, 0, 0, 122, 0, 0, 0,
	0, 0, 0, 115, 0, 124, 0, 77, 76, 0,
	0, 128, 0, 0, 0, 0, 0, 128, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 122, 0, 0, 0, 122, 0, 122, 0,
	128, 0, 0, 0, 0, 89, 0, 0, 0, 0,
	122, 0, 128, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 128, 0, 128, 0, 0, 0, 128,
	9, 0, 0, 0, 98, 99, 128, 0, 0, 122,
	122, 87, 88, 0, 0, 128, 90, 91, 92, 100,
	0, 0, 0, 0, 122, 0, 86, 95, 93, 94,
	97, 128, 128, 299, 0, 0, 0, 0, 0, 0,
	0, 0, 120, 0, 0, 128, 128, 0, 0, 0,
	0, 0, 128, 0, 0, 0, 0, 0, 0, 148,
	0, 150, 148, 0, 0, 0, 154, 0, 0, 0,
	128, 0, 0, 0, 165, 166, 167, 0, 0, 0,
	0, 62, 125, 68, 126, 69, 175, 0, 177, 178,
	0, 180, 0, 182, 183, 184, 185, 0, 187, 0,
	0, 190, 0, 192, 194, 0, 128, 0, 0, 0,
	128, 70, 128, 211, 78, 79, 215, 218, 221, 71,
	72, 73, 74, 75, 128, 0, 120, 260, 0, 0,
	0, 211, 0, 0, 234, 0, 0, 65, 0, 66,
	0, 77, 76, 0, 45, 0, 0, 0, 0, 0,
	0, 0, 0, 128, 128, 0, 0, 259, 264, 0,
	211, 0, 0, 0, 0, 0, 0, 0, 128, 0,
	0, 0, 0, 0, 0, 0, 0, 120, 0, 0,
	277, 0, 0, 0, 0, 0, 129, 0, 0, 279,
	280, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 229, 129, 0, 129, 129, 0, 0, 0,
	129, 0, 0, 0, 0, 0, 0, 0, 129, 129,
	129, 0, 0, 0, 0, 303, 0, 0, 0, 0,
	129, 311, 129, 129, 0, 129, 0, 129, 129, 129,
	129, 0, 129, 0, 0, 129, 0, 129, 129, 0,
	0, 0, 0, 146, 325, 89, 0, 129, 0, 0,
	129, 129, 129, 0, 0, 0, 120, 0, 0, 0,
	129, 0, 0, 0, 0, 129, 0, 211, 129, 340,
	0, 0, 0, 303, 98, 99, 0, 0, 0, 0,
	348, 87, 88, 0, 0, 0, 90, 91, 92, 354,
	0, 129, 129, 0, 129, 0, 86, 95, 93, 94,
	97, 0, 207, 0, 209, 365, 366, 0, 0, 0,
	0, 129, 0, 0, 129, 0, 225, 0, 0, 148,
	391, 0, 0, 129, 129, 0, 399, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 246, 0, 0,
	0, 0, 0, 0, 391, 0, 0, 0, 0, 0,
	0, 0, 0, 89, 0, 0, 0, 0, 0, 129,
	0, 0, 0, 0, 0, 129, 0, 96, 0, 0,
	0, 0, 0, 0, 85, 0, 0, 0, 0, 0,
	148, 0, 98, 99, 431, 0, 432, 0, 129, 87,
	88, 285, 0, 0, 90, 91, 92, 100, 431, 292,
	129, 0, 0, 0, 86, 95, 93, 94, 97, 0,
	0, 129, 0, 129, 0, 0, 0, 129, 0, 306,
	0, 309, 0, 0, 129, 0, 0, 120, 454, 0,
	0, 0, 0, 129, 0, 0, 0, 0, 0, 0,
	320, 321, 461, 0, 0, 0, 0, 0, 0, 129,
	129, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 129, 129, 62, 41, 68, 42, 69,
	129, 0, 0, 38, 0, 48, 0, 0, 49, 39,
	40, 0, 59, 0, 50, 0, 0, 351, 129, 0,
	0, 61, 58, 0, 0, 70, 60, 0, 78, 79,
	0, 0, 0, 71, 72, 73, 74, 75, 0, 0,
	0, 0, 371, 0, 0, 0, 0, 375, 0, 0,
	0, 65, 0, 66, 129, 77, 76, 0, 129, 0,
	129, 0, 0, 0, 0, 0, 0, 0, 404, 0,
	0, 0, 129, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 417, 418, 0, 0, 0, 0, 0, 420,
	0, 129, 129, 0, 0, 0, 0, 0, 0, 0,
	0, 428, 0, 430, 0, 0, 129, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 438,
	0, 0, 0, 0, 0, 246, 0, 0, 0, 0,
	0, 0, 448, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 455, 0, 320, 321, 62, 41, 68, 42,
	69, 0, 0, 0, 38, 467, 383, 466, 465, 384,
	39, 40, 0, 59, 0, 50, 0, 0, 385, 386,
	0, 0, 61, 58, 0, 0, 70, 60, 0, 78,
	79, 0, 0, 0, 71, 72, 73, 74, 75, 0,
	0, 0, 378, 379, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 66, 0, 77, 76, 62, 41,
	68, 42, 69, 0, 0, 0, 38, 463, 383, 466,
	465, 384, 39, 40, 0, 59, 0, 50, 0, 0,
	385, 386, 0, 0, 61, 58, 0, 0, 70, 60,
	0, 78, 79, 0, 0, 0, 71, 72, 73, 74,
	75, 0, 0, 0, 378, 379, 0, 0, 0, 0,
	0, 0, 0, 0, 65, 0, 66, 0, 77, 76,
	62, 41, 68, 42, 69, 0, 0, 0, 38, 469,
	383, 0, 0, 384, 39, 40, 0, 59, 0, 50,
	0, 0, 385, 386, 0, 0, 61, 58, 0, 0,
	70, 60, 0, 78, 79, 0, 0, 0, 71, 72,
	73, 74, 75, 0, 0, 0, 378, 379, 0, 0,
	0, 0, 0, 0, 0, 0, 65, 0, 66, 0,
	77, 76, 62, 41, 68, 42, 69, 0, 0, 0,
	38, 377, 383, 0, 0, 384, 39, 40, 0, 59,
	0, 50, 0, 0, 385, 386, 0, 0, 61, 58,
	0, 0, 70, 60, 0, 78, 79, 0, 0, 0,
	71, 72, 73, 74, 75, 0, 0, 0, 378, 379,
	0, 0, 0, 0, 0, 0, 0, 0, 65, 0,
	66, 0, 77, 76, 62, 41, 68, 42, 69, 0,
	0, 0, 38, 372, 48, 319, 318, 49, 39, 40,
	0, 59, 0, 50, 0, 0, 0, 0, 0, 0,
	61, 58, 0, 0, 70, 60, 0, 78, 79, 0,
	0, 0, 71, 72, 73, 74, 75, 0, 0, 0,
	249, 250, 0, 0, 0, 0, 0, 0, 0, 0,
	65, 0, 66, 0, 77, 76, 62, 41, 68, 42,
	69, 0, 0, 0, 38, 316, 48, 319, 318, 49,
	39, 40, 0, 59, 0, 50, 0, 0, 0, 0,
	0, 0, 61, 58, 0, 0, 70, 60, 0, 78,
	79, 0, 0, 0, 71, 72, 73, 74, 75, 0,
	0, 0, 249, 250, 0, 62, 41, 68, 42, 69,
	0, 0, 65, 38, 66, 48, 77, 76, 49, 39,
	40, 0, 59, 0, 50, 0, 0, 0, 0, 0,
	0, 61, 58, 0, 0, 70, 60, 0, 78, 79,
	0, 0, 0, 71, 72, 73, 74, 75, 0, 0,
	0, 6, 7, 0, 0, 0, 0, 0, 0, 0,
	0, 65, 0, 66, 0, 77, 76, 0, 8, 62,
	41, 68, 42, 69, 0, 0, 0, 38, 0, 383,
	0, 0, 384, 39, 40, 0, 59, 0, 50, 0,
	0, 385, 386, 0, 0, 61, 58, 0, 0, 70,
	60, 0, 78, 79, 0, 0, 0, 71, 72, 73,
	74, 75, 0, 0, 0, 378, 379, 0, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 66, 0, 77,
	76, 62, 41, 68, 42, 69, 0, 0, 0, 38,
	434, 48, 0, 0, 49, 39, 40, 0, 59, 0,
	50, 324, 0, 0, 0, 0, 0, 61, 58, 0,
	0, 70, 60, 0, 78, 79, 0, 0, 0, 71,
	72, 73, 74, 75, 0, 0, 0, 249, 250, 0,
	0, 0, 0, 0, 0, 0, 0, 65, 0, 66,
	0, 77, 76, 62, 41, 68, 42, 69, 0, 0,
	0, 38, 349, 48, 0, 0, 49, 39, 40, 0,
	59, 0, 50, 324, 0, 0, 0, 0, 0, 61,
	58, 0, 0, 70, 60, 0, 78, 79, 0, 0,
	0, 71, 72, 73, 74, 75, 0, 0, 0, 249,
	250, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	0, 66, 0, 77, 76, 62, 41, 68, 42, 69,
	0, 0, 0, 38, 342, 48, 0, 0, 49, 39,
	40, 0, 59, 0, 50, 324, 0, 0, 0, 0,
	0, 61, 58, 0, 0, 70, 60, 0, 78, 79,
	0, 0, 0, 71, 72, 73, 74, 75, 0, 0,
	0, 249, 250, 0, 0, 0, 0, 0, 0, 0,
	0, 65, 0, 66, 0, 77, 76, 62, 41, 68,
	42, 69, 0, 0, 0, 38, 447, 48, 0, 0,
	49, 39, 40, 0, 59, 0, 50, 0, 0, 0,
	0, 0, 0, 61, 58, 0, 0, 70, 60, 0,
	78, 79, 0, 0, 0, 71, 72, 73, 74, 75,
	0, 0, 0, 249, 250, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 0, 66, 0, 77, 76, 62,
	41, 68, 42, 69, 0, 0, 0, 38, 406, 48,
	0, 0, 49, 39, 40, 0, 59, 0, 50, 0,
	0, 0, 0, 0, 0, 61, 58, 0, 0, 70,
	60, 0, 78, 79, 0, 0, 0, 71, 72, 73,
	74, 75, 0, 0, 0, 249, 250, 0, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 66, 0, 77,
	76, 62, 41, 68, 42, 69, 0, 0, 0, 38,
	374, 48, 0, 0, 49, 39, 40, 0, 59, 0,
	50, 0, 0, 0, 0, 0, 0, 61, 58, 0,
	0, 70, 60, 0, 78, 79, 0, 0, 0, 71,
	72, 73, 74, 75, 0, 0, 0, 249, 250, 0,
	62, 41, 68, 42, 69, 0, 0, 65, 38, 66,
	48, 77, 76, 49, 39, 40, 0, 59, 0, 50,
	0, 0, 0, 0, 0, 0, 61, 58, 0, 0,
	70, 60, 0, 78, 79, 0, 0, 0, 71, 72,
	73, 74, 75, 0, 0, 0, 249, 250, 0, 0,
	0, 0, 0, 0, 0, 0, 65, 0, 66, 362,
	77, 76, 62, 41, 68, 42, 69, 0, 0, 0,
	38, 359, 48, 0, 0, 49, 39, 40, 0, 59,
	0, 50, 0, 0, 0, 0, 0, 0, 61, 58,
	0, 0, 70, 60, 0, 78, 79, 0, 0, 0,
	71, 72, 73, 74, 75, 0, 0, 0, 249, 250,
	0, 62, 41, 68, 42, 69, 0, 0, 65, 38,
	66, 48, 77, 76, 49, 39, 40, 0, 59, 0,
	50, 0, 0, 0, 0, 0, 0, 61, 58, 0,
	0, 70, 60, 0, 78, 79, 0, 0, 0, 71,
	72, 73, 74, 75, 0, 0, 0, 249, 250, 0,
	0, 0, 0, 0, 0, 0, 0, 65, 0, 66,
	310, 77, 76, 62, 41, 68, 42, 69, 0, 0,
	0, 38, 305, 48, 0, 0, 49, 39, 40, 0,
	59, 0, 50, 0, 0, 0, 0, 0, 0, 61,
	58, 0, 0, 70, 60, 0, 78, 79, 0, 0,
	0, 71, 72, 73, 74, 75, 0, 0, 0, 249,
	250, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	0, 66, 0, 77, 76, 62, 41, 68, 42, 69,
	0, 0, 0, 38, 295, 48, 0, 0, 49, 39,
	40, 0, 59, 0, 50, 0, 0, 0, 0, 0,
	0, 61, 58, 0, 0, 70, 60, 0, 78, 79,
	0, 0, 0, 71, 72, 73, 74, 75, 0, 0,
	0, 249, 250, 0, 0, 0, 0, 0, 0, 0,
	0, 65, 0, 66, 0, 77, 76, 62, 41, 68,
	42, 69, 0, 0, 0, 38, 293, 48, 0, 0,
	49, 39, 40, 0, 59, 0, 50, 0, 0, 0,
	0, 0, 0, 61, 58, 0, 0, 70, 60, 0,
	78, 79, 0, 0, 0, 71, 72, 73, 74, 75,
	0, 0, 0, 249, 250, 0, 62, 41, 68, 42,
	69, 0, 0, 65, 38, 66, 48, 77, 76, 49,
	39, 40, 0, 59, 0, 50, 0, 0, 0, 0,
	0, 0, 61, 58, 0, 0, 70, 60, 0, 78,
	79, 0, 0, 0, 71, 72, 73, 74, 75, 0,
	0, 0, 249, 250, 0, 62, 41, 68, 42, 69,
	270, 0, 65, 38, 66, 48, 77, 76, 49, 39,
	40, 0, 59, 0, 50, 0, 0, 0, 0, 0,
	0, 61, 58, 0, 0, 70, 60, 0, 78, 79,
	0, 0, 0, 71, 72, 73, 74, 75, 0, 0,
	0, 0, 269, 62, 125, 68, 126, 112, 453, 0,
	123, 65, 0, 66, 0, 77, 76, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 70, 0, 0, 78, 79, 0, 0,
	0, 71, 72, 73, 74, 75, 62, 125, 68, 126,
	69, 0, 0, 0, 0, 0, 0, 0, 0, 230,
	0, 124, 0, 77, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 70, 0, 0, 78,
	79, 0, 0, 0, 71, 72, 73, 74, 75, 0,
	0, 0, 260, 62, 125, 68, 126, 69, 0, 393,
	0, 0, 65, 0, 66, 0, 77, 76, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 70, 0, 0, 78, 79, 0, 0,
	0, 71, 72, 73, 74, 75, 0, 0, 0, 260,
	62, 125, 68, 126, 112, 0, 355, 123, 0, 65,
	0, 66, 0, 77, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	70, 0, 0, 78, 79, 0, 0, 0, 71, 72,
	73, 74, 75, 62, 125, 68, 126, 276, 335, 0,
	123, 0, 0, 0, 0, 0, 230, 0, 124, 0,
	77, 76, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 70, 0, 0, 78, 79, 0, 0,
	275, 71, 72, 73, 74, 75, 62, 125, 68, 126,
	69, 0, 0, 123, 0, 0, 0, 0, 0, 65,
	0, 124, 0, 77, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 70, 0, 0, 78,
	79, 0, 0, 0, 71, 72, 73, 74, 75, 62,
	125, 68, 126, 69, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 124, 0, 77, 76, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 70,
	0, 0, 78, 79, 0, 0, 0, 71, 72, 73,
	74, 75, 62, 265, 68, 126, 69, 0, 0, 0,
	0, 0, 156, 0, 0, 65, 0, 66, 0, 77,
	76, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 70, 0, 0, 78, 79, 0, 0, 0,
	71, 72, 73, 74, 75, 0, 0, 0, 260, 62,
	125, 68, 126, 112, 0, 0, 123, 0, 65, 0,
	66, 0, 77, 76, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 70,
	0, 0, 78, 79, 0, 0, 0, 71, 72, 73,
	74, 75, 62, 125, 68, 126, 69, 0, 0, 0,
	0, 0, 0, 0, 0, 230, 0, 124, 0, 77,
	76, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 70, 0, 0, 78, 79, 0, 0, 0,
	71, 72, 73, 74, 75, 62, 436, 68, 126, 69,
	0, 0, 0, 0, 0, 0, 0, 0, 65, 0,
	66, 0, 77, 76, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 70, 0, 0, 78, 79,
	0, 0, 0, 71, 72, 73, 74, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 65, 0, 66, 0, 77, 76,
}
var RubyPact = []int{

	-32, 1880, -1000, -1000, -1000, 6, -1000, -1000, -1000, 1239,
	-1000, -1000, -1000, 91, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 26, 8,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 326, 221,
	221, 752, 155, 33, 136, 123, 171, 156, 1350, 1350,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 3207, 1350,
	3207, 3207, -1000, -1000, -1000, 3074, -1000, 15, 197, -1000,
	18, 1350, 1350, 3207, 3207, 3207, 20, 323, -1000, -1000,
	-1000, -1000, -1000, 1350, 1350, 3207, 345, 3207, 3207, -1000,
	3207, 1350, 3207, 3207, 3207, 3207, 1350, 3207, -1000, -1000,
	3207, 1350, 3207, 3207, 1350, 1350, 344, 119, 113, 55,
	-1000, -1000, 3074, 18, 27, 3074, 3207, 3207, 342, 176,
	613, -1000, 7, -5, -5, 3164, 104, -15, -1000, -1000,
	3074, 1350, 1350, 3207, 1350, 1350, 1350, 1350, 1350, 1350,
	1350, 341, 340, 338, 181, 132, 2711, 110, 613, 204,
	613, 110, 54, 45, 1131, -1000, 3207, 3117, 230, 3074,
	2760, -1000, -5, 181, 181, 613, 613, 613, -1000, -1000,
	320, -1000, -1000, 181, 181, 613, 2988, 613, 613, 956,
	613, 181, 613, 613, 613, 613, 181, 670, 956, 956,
	613, 181, 613, 37, 611, 181, 181, 18, -1000, 337,
	170, 53, -1000, 70, 334, 332, -1000, 2662, 221, 2600,
	314, 1131, -1000, -1000, -1000, 851, -39, 34, 466, -1000,
	-1000, 210, -1000, -1000, 3031, 2538, -1000, 330, -1000, 2476,
	3074, 307, 181, 181, 319, 181, 181, 181, 181, 181,
	181, 181, -1000, 168, -26, -27, 1831, -1000, -1000, -1000,
	-1000, 181, 222, 3207, -1000, 42, -1000, -1000, -1000, 613,
	-1000, 158, 146, 29, 737, 2945, -1000, 287, 181, -1000,
	-1000, -1000, -1000, 27, 18, 1350, 3074, 613, 3207, 613,
	613, -1000, 3031, 61, -1000, 2130, 113, 53, 282, 3207,
	-1000, -1000, 2068, -1000, -1000, -1000, 18, -1000, 2898, 144,
	-1000, -1000, 13, 613, -1000, -1000, 2427, 67, -1000, 2365,
	-1000, 362, -1000, 59, 3207, 3207, -1000, 299, 1350, -1000,
	1769, 2316, -1000, -1000, 203, 613, 1707, 209, 3207, 2851,
	-22, -1000, -24, -1000, 1350, 3207, -1000, -1000, 181, 223,
	613, 1350, -1000, 214, -1000, -1000, -1000, -1000, 613, -1000,
	212, 2254, -1000, 555, 613, 327, 1350, 315, 308, -1000,
	-1000, 304, -1000, 47, 1350, 245, 220, -1000, 1350, -1000,
	181, 2711, -1000, 273, -1000, 2711, 216, -1000, -1000, -1000,
	181, -1000, -1000, 1350, 1350, 262, 255, -1000, -1000, 3207,
	110, 1131, -1000, 3207, -1000, 956, -1000, 98, 181, 613,
	-1000, 181, -1000, -1000, 2006, -1000, -1000, 3250, -1000, 181,
	-28, -1000, -1000, 181, 118, -1000, 181, 2711, 2711, -1000,
	2711, 296, 73, 102, 1350, 1350, 1350, 1350, 2192, 110,
	2711, 613, 540, 24, -1000, 145, 2808, 3207, 2711, -1000,
	-1000, -1000, -1000, 181, 181, 181, 181, -1000, 2711, 29,
	1350, 3207, -1000, -1000, 21, 2711, 1583, 1521, 1645, 29,
	241, 593, -1000, -1000, 267, 1350, -1000, -1000, 250, -1000,
	-1000, -1000, 29, -1000, -1000, 1350, -1000, 181, 1944, -1000,
	29, 181, 1944, 1944, 1944,
}
var RubyPgo = []int{

	0, 0, 439, 438, 10, 37, 437, 436, 435, 1024,
	434, 5, 13, 433, 432, 431, 423, 880, 422, 596,
	452, 416, 415, 414, 32, 412, 7, 114, 406, 405,
	12, 401, 400, 399, 394, 17, 392, 390, 389, 3,
	388, 387, 383, 382, 379, 378, 376, 373, 372, 371,
	370, 1082, 367, 6, 18, 22, 1, 364, 9, 363,
	4, 361, 31, 358, 8, 355, 11, 15, 16, 14,
	354, 288, 43,
}
var RubyR1 = []int{

	0, 57, 57, 57, 57, 57, 57, 57, 57, 57,
	57, 71, 71, 72, 72, 51, 51, 51, 51, 18,
	18, 18, 18, 18, 18, 18, 18, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 24, 24, 24, 24, 24, 24, 24, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 35, 14, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 54, 54, 54, 54, 64, 64, 62,
	62, 62, 62, 62, 62, 62, 68, 68, 68, 68,
	68, 66, 66, 66, 21, 21, 21, 21, 21, 21,
	58, 58, 69, 69, 69, 26, 26, 26, 26, 25,
	25, 28, 30, 30, 70, 70, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 67, 67, 29, 29, 29,
	29, 29, 29, 9, 9, 27, 27, 19, 19, 40,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
	50, 2, 6, 7, 7, 3, 3, 59, 59, 59,
	59, 65, 65, 65, 5, 5, 5, 5, 55, 63,
	63, 63, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 56, 56, 56, 56, 52, 52, 52,
	8, 16, 11, 11, 11, 61, 61, 53, 53, 22,
	22, 23, 23, 12, 36, 60, 60, 60, 60, 60,
	60, 37, 37, 37, 37, 37, 37, 37, 38, 38,
	38, 38, 38, 39, 39, 39, 39, 34, 33, 10,
	32, 32, 31, 31, 4,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 1, 0, 2, 0, 2, 2, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 2, 4, 5, 1, 4, 4,
	2, 3, 3, 4, 4, 5, 3, 4, 5, 2,
	3, 3, 3, 4, 4, 4, 4, 4, 4, 4,
	6, 6, 6, 3, 7, 1, 5, 1, 3, 0,
	1, 1, 2, 4, 4, 5, 1, 1, 4, 2,
	5, 1, 3, 3, 5, 6, 7, 8, 5, 6,
	1, 3, 0, 1, 3, 1, 2, 3, 2, 4,
	6, 4, 1, 3, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 9, 6, 3, 3, 3,
	3, 3, 3, 2, 2, 2, 2, 3, 3, 3,
	3, 2, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 1, 1, 3, 3, 5, 5, 0, 1, 3,
	7, 3, 7, 8, 3, 4, 4, 3, 3, 0,
	1, 3, 4, 5, 3, 3, 3, 3, 3, 5,
	6, 5, 3, 4, 3, 3, 2, 0, 2, 2,
	3, 4, 2, 3, 5, 0, 2, 1, 2, 2,
	1, 2, 1, 5, 5, 0, 2, 2, 2, 2,
	2, 0, 1, 3, 3, 1, 3, 3, 5, 6,
	5, 6, 5, 4, 3, 3, 2, 4, 4, 2,
	5, 7, 4, 5, 3,
}
var RubyChk = []int{

	-1000, -57, 51, 52, 68, -1, 51, 52, 68, -17,
	-21, -25, -28, -15, -29, -13, -16, -24, -22, -36,
	-32, -23, -35, -4, -18, -7, -3, -30, -20, -8,
	-10, -41, -42, -43, -44, -14, -34, -33, 13, 19,
	20, 6, 8, -27, -19, -9, -40, -67, 15, 18,
	24, -45, -46, -47, -48, -49, -50, -12, 32, 22,
	36, 31, 5, -2, -6, 61, 63, -70, 7, 9,
	35, 43, 44, 45, 46, 47, 66, 65, 38, 39,
	52, 51, 68, 15, 18, 25, 55, 40, 41, 4,
	45, 46, 47, 57, 58, 56, 18, 59, 33, 34,
	48, 18, 40, 61, 15, 18, 55, 6, 4, -30,
	8, -30, 9, -54, -68, 61, 42, 49, 11, -66,
	-17, -5, -20, 12, 63, 6, 8, -27, -19, -9,
	9, 42, 49, 61, 42, 49, 42, 49, 42, 49,
	42, 11, 42, 11, -1, -1, -51, -64, -17, -1,
	-17, -64, -64, -62, -17, -24, 58, -72, 54, 9,
	-52, -5, 63, -1, -1, -17, -17, -17, 6, 8,
	66, 6, 8, -1, -1, -17, 6, -17, -17, -72,
	-17, -1, -17, -17, -17, -17, -1, -17, -72, -72,
	-17, -1, -17, -66, -17, -1, -1, 6, -58, 55,
	-69, 9, -26, 6, 47, 58, -58, -51, 40, -51,
	-62, -17, -5, 11, -5, -17, -4, -66, -17, -35,
	-12, -17, -12, 6, 11, -51, -55, 56, -55, -51,
	61, -62, -1, -1, -17, -1, -1, -1, -1, -1,
	-1, -1, 6, -67, 6, 6, -51, 51, 52, 51,
	52, -1, -61, 11, 51, -72, 62, 11, 62, -17,
	51, -59, -65, -72, -17, 6, 8, -62, -1, 52,
	10, 6, 8, -68, -54, 42, 9, -17, 53, -17,
	-17, 62, 11, 62, -5, -51, 6, 11, -69, 42,
	6, 6, -51, 14, -30, 14, 10, 11, -72, 62,
	62, 62, -72, -17, -5, 14, -51, -63, 6, -51,
	64, -17, 10, 62, 61, 61, 14, -56, 17, 16,
	-51, -51, 14, -11, 25, -17, -60, -31, 37, -72,
	-72, 11, -72, 11, 4, 53, 10, -5, -1, -62,
	-17, 42, 14, -53, -11, -58, -26, 10, -17, 14,
	-53, -51, -5, -72, -17, 58, 42, 11, 58, 14,
	56, 11, 64, 62, 42, -17, -17, 14, 17, 16,
	-1, -51, 14, -56, 14, -51, 8, 14, 51, 52,
	-1, -38, -37, 15, 18, 27, 28, 14, 16, 37,
	-64, -17, -24, 58, 64, -72, 64, -72, -1, -17,
	10, -1, 14, -11, -51, 14, 14, 58, 6, -1,
	6, 6, 6, -1, 62, 62, -1, -51, -51, 14,
	-51, 4, -1, -1, 15, 18, 15, 18, -51, -64,
	-51, -17, -17, 6, 14, -53, 6, 61, -51, 6,
	51, 51, 52, -1, -1, -1, -1, 14, -51, -72,
	4, 53, 14, 10, -17, -51, -60, -60, -60, -72,
	-1, -17, 62, 14, -39, 17, 16, 14, -39, 14,
	-71, 11, -72, 11, 14, 17, 16, -1, -60, 14,
	-72, -1, -60, -60, -60,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 48,
	49, 50, 51, 52, 53, 54, 55, 56, 57, 58,
	59, 60, 61, 62, 27, 28, 29, 30, 31, 32,
	33, 34, 35, 36, 37, 38, 39, 40, 0, 0,
	0, 20, -2, 22, 23, 24, 0, 0, 0, 0,
	15, 41, 42, 43, 44, 45, 46, 47, 220, 0,
	0, 222, 19, 25, 26, 99, 13, 0, 67, 207,
	0, 0, 0, 0, 0, 0, 0, 0, 171, 172,
	5, 6, 7, 0, 0, 0, 0, 0, 0, 13,
	0, 0, 0, 0, 0, 0, 0, 0, 13, 13,
	0, 0, 0, 0, 0, 0, 0, 122, 122, 15,
	-2, 15, -2, 70, 79, 99, 0, 0, 0, 95,
	106, 107, 31, 15, -2, 20, -2, 22, 23, 24,
	99, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 15, 0, 215, 219, 97, 0,
	13, 221, 0, 0, 97, 101, 0, 13, 0, 99,
	0, 249, 15, 161, 162, 163, 164, 64, 155, 156,
	0, 153, 154, 194, 202, 63, 72, 80, 82, 0,
	165, 166, 167, 168, 169, 170, 196, 0, 0, 0,
	254, 198, 81, 0, 111, 195, 197, 76, 15, 0,
	120, 122, 123, 125, 0, 0, 15, 0, 0, 0,
	0, 100, 71, 13, 109, 97, 0, 0, 136, 137,
	138, 147, 148, 159, 13, 0, 15, 189, 15, 0,
	99, 0, 139, 149, 0, 140, 150, 141, 151, 142,
	152, 143, 160, 144, 0, 0, 0, 15, 15, 16,
	17, 18, 0, 0, 225, 0, 173, 13, 174, 102,
	14, 13, 13, 178, 0, 20, -2, 0, 208, 209,
	210, 157, 158, 73, 74, 0, -2, 83, 0, 247,
	248, 88, 0, 89, 77, 0, 122, 0, 0, 0,
	126, 128, 0, 129, 15, 131, 65, 13, 0, 84,
	86, 87, 0, 112, 113, 184, 0, 0, 190, 0,
	187, 97, 69, 85, 0, 0, 192, 0, 0, 15,
	0, 0, 211, 216, 15, 98, 0, 0, 0, 0,
	0, 13, 0, 13, 13, 0, 68, 75, 78, 0,
	223, 0, 114, 0, 217, 15, 124, 121, 127, 118,
	0, 0, 66, 0, 108, 0, 0, 0, 0, 185,
	188, 0, 186, 84, 0, 0, 0, 193, 0, 15,
	15, 206, 199, 0, 201, 212, 15, 224, 226, 227,
	228, 229, 230, 0, 0, 232, 235, 250, 15, 0,
	15, 103, 104, 0, 175, 0, 176, 0, 179, 181,
	93, 92, 115, 218, 0, 119, 130, 0, 110, 90,
	0, 96, 191, 91, 0, 146, 15, 204, 205, 200,
	213, 0, 15, 0, 0, 0, 0, 0, 0, 15,
	13, 105, 0, 0, 116, 0, 20, 0, 203, 15,
	225, -2, -2, 233, 234, 236, 237, 251, 13, 252,
	13, 0, 117, 94, 0, 214, 0, 0, 0, 253,
	11, 13, 145, 238, 0, 0, 225, 240, 0, 242,
	180, 12, 182, 13, 239, 0, 225, 225, 231, 241,
	183, 225, 231, 231, 231,
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
		//line parser.y:202
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:204
		{
		}
	case 3:
		//line parser.y:206
		{
		}
	case 4:
		//line parser.y:208
		{
		}
	case 5:
		//line parser.y:210
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:212
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:214
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:220
		{
		}
	case 11:
		//line parser.y:222
		{
		}
	case 12:
		//line parser.y:223
		{
		}
	case 13:
		//line parser.y:225
		{
		}
	case 14:
		//line parser.y:226
		{
		}
	case 15:
		//line parser.y:229
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 16:
		//line parser.y:231
		{
		}
	case 17:
		//line parser.y:233
		{
		}
	case 18:
		//line parser.y:235
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
		//line parser.y:247
		{
			RubyVAL.genericValue = ast.RescueModifier{Statement: RubyS[Rubypt-2].genericValue, Rescue: RubyS[Rubypt-0].genericValue}
		}
	case 64:
		//line parser.y:250
		{
			RubyVAL.genericValue = ast.StarSplat{Value: RubyS[Rubypt-0].genericValue}
		}
	case 65:
		//line parser.y:253
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 66:
		//line parser.y:260
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-2].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 67:
		//line parser.y:267
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 68:
		//line parser.y:271
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 69:
		//line parser.y:278
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 70:
		//line parser.y:285
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 71:
		//line parser.y:292
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 72:
		//line parser.y:299
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 73:
		//line parser.y:306
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 74:
		//line parser.y:314
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 75:
		//line parser.y:322
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 76:
		//line parser.y:330
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 77:
		//line parser.y:338
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 78:
		//line parser.y:346
		{
			methodName := RubyS[Rubypt-2].genericValue.(ast.BareReference).Name + "="
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: methodName},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 79:
		//line parser.y:357
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 80:
		//line parser.y:364
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 81:
		//line parser.y:372
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 82:
		//line parser.y:380
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 83:
		//line parser.y:388
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 84:
		//line parser.y:398
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 85:
		//line parser.y:406
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 86:
		//line parser.y:414
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 87:
		//line parser.y:422
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 88:
		//line parser.y:430
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 89:
		//line parser.y:438
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 90:
		//line parser.y:448
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 91:
		//line parser.y:456
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 92:
		//line parser.y:464
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 93:
		//line parser.y:474
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 94:
		//line parser.y:476
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 95:
		//line parser.y:478
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 96:
		//line parser.y:480
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 97:
		//line parser.y:483
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 98:
		//line parser.y:485
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 99:
		//line parser.y:487
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 100:
		//line parser.y:489
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 101:
		//line parser.y:491
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 102:
		//line parser.y:493
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{Func: ast.BareReference{Name: "to_proc"}, Target: RubyS[Rubypt-0].genericValue})
		}
	case 103:
		//line parser.y:495
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 104:
		//line parser.y:497
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 105:
		//line parser.y:499
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{
				Func:   ast.BareReference{Name: "to_proc"},
				Target: RubyS[Rubypt-0].genericValue,
			})
		}
	case 106:
		//line parser.y:508
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 107:
		//line parser.y:510
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 108:
		//line parser.y:512
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 109:
		//line parser.y:514
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 110:
		//line parser.y:516
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 111:
		//line parser.y:519
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 112:
		//line parser.y:521
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 113:
		//line parser.y:523
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 114:
		//line parser.y:527
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 115:
		//line parser.y:535
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 116:
		//line parser.y:544
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 117:
		//line parser.y:553
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 118:
		//line parser.y:563
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 119:
		//line parser.y:571
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 120:
		//line parser.y:582
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 121:
		//line parser.y:584
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 122:
		//line parser.y:586
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 123:
		//line parser.y:588
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 124:
		//line parser.y:590
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 125:
		//line parser.y:593
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 126:
		//line parser.y:595
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 127:
		//line parser.y:597
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 128:
		//line parser.y:599
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 129:
		//line parser.y:603
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 130:
		//line parser.y:610
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 131:
		//line parser.y:620
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 132:
		//line parser.y:629
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 133:
		//line parser.y:635
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 134:
		//line parser.y:643
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 135:
		//line parser.y:647
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 136:
		//line parser.y:652
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 137:
		//line parser.y:659
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 138:
		//line parser.y:663
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 139:
		//line parser.y:667
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 140:
		//line parser.y:674
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 141:
		//line parser.y:681
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 142:
		//line parser.y:688
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 143:
		//line parser.y:695
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 144:
		//line parser.y:702
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: ast.Array{Nodes: RubyS[Rubypt-2].genericSlice},
				RHS: ast.Array{Nodes: RubyS[Rubypt-0].genericSlice},
			}
		}
	case 145:
		//line parser.y:710
		{
			RubyVAL.genericSlice = []ast.Node{
				ast.CallExpression{
					Target: RubyS[Rubypt-8].genericValue,
					Func:   ast.BareReference{Name: "[]="},
					Args:   []ast.Node{RubyS[Rubypt-6].genericValue},
				},
				ast.CallExpression{
					Target: RubyS[Rubypt-3].genericValue,
					Func:   ast.BareReference{Name: "[]="},
					Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
				},
			}
		}
	case 146:
		//line parser.y:725
		{
			tail := ast.CallExpression{Target: RubyS[Rubypt-3].genericValue, Func: ast.BareReference{Name: "[]="}, Args: []ast.Node{RubyS[Rubypt-1].genericValue}}
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, tail)
		}
	case 147:
		//line parser.y:731
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 148:
		//line parser.y:738
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 149:
		//line parser.y:742
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 150:
		//line parser.y:749
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 151:
		//line parser.y:756
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 152:
		//line parser.y:763
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 153:
		//line parser.y:771
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 154:
		//line parser.y:773
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 155:
		//line parser.y:776
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 156:
		//line parser.y:778
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 157:
		//line parser.y:781
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 158:
		//line parser.y:783
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 159:
		//line parser.y:786
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 160:
		//line parser.y:788
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 161:
		//line parser.y:790
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 162:
		//line parser.y:791
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 163:
		//line parser.y:792
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 164:
		//line parser.y:793
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 165:
		//line parser.y:796
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 166:
		//line parser.y:805
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 167:
		//line parser.y:814
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 168:
		//line parser.y:823
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 169:
		//line parser.y:832
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 170:
		//line parser.y:841
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 171:
		//line parser.y:849
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 172:
		//line parser.y:850
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 173:
		//line parser.y:852
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 174:
		//line parser.y:854
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 175:
		//line parser.y:857
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 176:
		//line parser.y:865
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 177:
		//line parser.y:873
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 178:
		//line parser.y:875
		{
		}
	case 179:
		//line parser.y:877
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 180:
		//line parser.y:884
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 181:
		//line parser.y:892
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 182:
		//line parser.y:899
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 183:
		//line parser.y:906
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 184:
		//line parser.y:914
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 185:
		//line parser.y:916
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 186:
		//line parser.y:923
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 187:
		//line parser.y:930
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 188:
		//line parser.y:933
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 189:
		//line parser.y:935
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 190:
		//line parser.y:937
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 191:
		//line parser.y:939
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 192:
		//line parser.y:942
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 193:
		//line parser.y:949
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 194:
		//line parser.y:957
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 195:
		//line parser.y:964
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 196:
		//line parser.y:971
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 197:
		//line parser.y:978
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 198:
		//line parser.y:985
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 199:
		//line parser.y:992
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 200:
		//line parser.y:999
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 201:
		//line parser.y:1007
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 202:
		//line parser.y:1014
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 203:
		//line parser.y:1023
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 204:
		//line parser.y:1030
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 205:
		//line parser.y:1037
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 206:
		//line parser.y:1044
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 207:
		//line parser.y:1051
		{
		}
	case 208:
		//line parser.y:1052
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 209:
		//line parser.y:1053
		{
		}
	case 210:
		//line parser.y:1056
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 211:
		//line parser.y:1059
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 212:
		//line parser.y:1067
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 213:
		//line parser.y:1069
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 214:
		//line parser.y:1078
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
	case 215:
		//line parser.y:1093
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 216:
		//line parser.y:1095
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 217:
		//line parser.y:1098
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 218:
		//line parser.y:1100
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 219:
		//line parser.y:1103
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 220:
		//line parser.y:1110
		{
			RubyVAL.genericValue = ast.Yield{}
		}
	case 221:
		//line parser.y:1113
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 222:
		//line parser.y:1121
		{
			RubyVAL.genericValue = ast.Return{}
		}
	case 223:
		//line parser.y:1124
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 224:
		//line parser.y:1133
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 225:
		//line parser.y:1136
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 226:
		//line parser.y:1138
		{
		}
	case 227:
		//line parser.y:1140
		{
		}
	case 228:
		//line parser.y:1142
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 229:
		//line parser.y:1144
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 230:
		//line parser.y:1146
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 231:
		//line parser.y:1148
		{
		}
	case 232:
		//line parser.y:1150
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 233:
		//line parser.y:1152
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 234:
		//line parser.y:1154
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 235:
		//line parser.y:1156
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 236:
		//line parser.y:1158
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 237:
		//line parser.y:1160
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 238:
		//line parser.y:1163
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 239:
		//line parser.y:1170
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 240:
		//line parser.y:1178
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 241:
		//line parser.y:1185
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 242:
		//line parser.y:1193
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 243:
		//line parser.y:1201
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 244:
		//line parser.y:1208
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 245:
		//line parser.y:1215
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 246:
		//line parser.y:1222
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 247:
		//line parser.y:1230
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 248:
		//line parser.y:1233
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 249:
		//line parser.y:1235
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 250:
		//line parser.y:1238
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 251:
		//line parser.y:1240
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 252:
		//line parser.y:1243
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 253:
		//line parser.y:1245
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 254:
		//line parser.y:1247
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
