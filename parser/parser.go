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

//line parser.y:1255

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 45,
	54, 137,
	-2, 21,
	-1, 116,
	54, 137,
	-2, 135,
	-1, 118,
	10, 102,
	11, 102,
	-2, 211,
	-1, 132,
	54, 137,
	-2, 21,
	-1, 278,
	51, 13,
	64, 13,
	-2, 31,
	-1, 281,
	54, 138,
	-2, 136,
	-1, 292,
	10, 102,
	11, 102,
	-2, 211,
	-1, 351,
	51, 13,
	-2, 230,
	-1, 460,
	51, 13,
	-2, 230,
}

const RubyNprod = 258
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 3174

var RubyAct = []int{

	261, 361, 5, 474, 360, 333, 162, 127, 209, 125,
	213, 25, 29, 129, 120, 316, 211, 163, 325, 119,
	274, 50, 88, 2, 3, 89, 178, 60, 179, 274,
	24, 274, 140, 430, 274, 451, 94, 298, 331, 141,
	4, 214, 413, 330, 411, 17, 107, 409, 224, 129,
	298, 139, 150, 151, 271, 115, 117, 263, 86, 85,
	139, 378, 94, 155, 172, 103, 104, 108, 153, 274,
	238, 157, 92, 93, 274, 87, 375, 173, 174, 238,
	112, 168, 215, 171, 461, 447, 180, 91, 317, 183,
	184, 103, 104, 216, 472, 274, 373, 191, 92, 93,
	172, 297, 196, 344, 46, 272, 377, 201, 270, 382,
	205, 206, 207, 91, 358, 305, 165, 274, 203, 88,
	433, 88, 89, 217, 89, 219, 106, 223, 225, 88,
	274, 228, 89, 227, 416, 313, 221, 88, 242, 243,
	89, 245, 246, 247, 248, 249, 250, 251, 237, 133,
	231, 233, 189, 230, 241, 455, 456, 257, 258, 266,
	267, 268, 269, 374, 165, 454, 133, 165, 133, 133,
	253, 283, 144, 264, 110, 133, 263, 111, 342, 145,
	149, 142, 165, 133, 133, 133, 147, 282, 143, 88,
	364, 303, 89, 214, 373, 133, 212, 133, 133, 107,
	133, 289, 133, 133, 133, 133, 290, 133, 109, 374,
	133, 148, 133, 133, 349, 165, 300, 146, 214, 462,
	108, 212, 421, 133, 167, 94, 133, 133, 133, 304,
	340, 347, 310, 340, 215, 287, 133, 288, 418, 101,
	338, 133, 210, 320, 133, 216, 90, 149, 136, 340,
	228, 340, 227, 303, 103, 104, 198, 199, 235, 215,
	160, 92, 93, 161, 339, 136, 95, 96, 97, 105,
	216, 133, 133, 402, 133, 403, 91, 100, 98, 99,
	102, 137, 489, 239, 486, 485, 165, 484, 138, 486,
	485, 133, 355, 394, 133, 481, 404, 225, 354, 88,
	353, 271, 89, 133, 133, 437, 320, 387, 386, 265,
	356, 362, 169, 367, 363, 385, 116, 387, 386, 94,
	369, 328, 271, 453, 94, 286, 483, 158, 94, 428,
	159, 312, 313, 281, 181, 388, 182, 114, 165, 113,
	133, 427, 391, 398, 133, 133, 426, 424, 103, 104,
	324, 405, 414, 103, 104, 92, 93, 103, 104, 417,
	92, 93, 419, 307, 92, 93, 306, 302, 133, 419,
	91, 255, 254, 252, 425, 91, 234, 314, 208, 91,
	186, 439, 432, 431, 480, 133, 329, 434, 318, 72,
	277, 407, 323, 262, 276, 1, 170, 133, 59, 133,
	58, 440, 441, 133, 67, 131, 73, 132, 74, 57,
	133, 443, 56, 55, 54, 37, 407, 36, 35, 133,
	34, 49, 399, 19, 345, 449, 39, 40, 20, 346,
	348, 350, 343, 14, 75, 133, 133, 83, 84, 30,
	12, 11, 76, 77, 78, 79, 80, 23, 22, 133,
	133, 419, 21, 18, 10, 31, 26, 133, 16, 13,
	70, 470, 71, 38, 82, 81, 370, 398, 398, 398,
	15, 478, 33, 32, 27, 133, 487, 69, 28, 380,
	68, 0, 0, 0, 128, 0, 491, 0, 0, 398,
	0, 0, 0, 398, 398, 398, 0, 0, 0, 0,
	410, 128, 412, 128, 128, 0, 0, 0, 0, 133,
	128, 0, 0, 133, 94, 133, 0, 0, 128, 128,
	128, 0, 47, 0, 0, 0, 0, 0, 133, 0,
	128, 0, 128, 128, 0, 128, 0, 128, 128, 128,
	128, 0, 128, 103, 104, 128, 0, 128, 128, 0,
	92, 93, 0, 94, 0, 133, 133, 0, 128, 0,
	0, 128, 128, 128, 0, 91, 133, 134, 0, 0,
	0, 128, 299, 0, 90, 0, 128, 0, 0, 128,
	0, 0, 103, 104, 134, 0, 134, 134, 0, 92,
	93, 0, 0, 134, 0, 0, 0, 459, 0, 0,
	0, 134, 134, 134, 91, 0, 128, 278, 102, 128,
	0, 469, 0, 134, 0, 134, 134, 0, 134, 0,
	134, 134, 134, 134, 482, 134, 128, 0, 134, 128,
	134, 134, 0, 466, 467, 468, 490, 0, 128, 128,
	0, 134, 0, 0, 134, 134, 134, 94, 0, 0,
	0, 0, 0, 0, 134, 488, 0, 0, 0, 134,
	0, 0, 134, 0, 94, 492, 493, 0, 0, 0,
	494, 0, 0, 0, 0, 128, 103, 104, 0, 278,
	128, 0, 0, 92, 93, 0, 0, 0, 0, 134,
	134, 0, 134, 103, 104, 0, 0, 0, 91, 94,
	92, 93, 102, 128, 0, 95, 96, 97, 0, 134,
	0, 0, 134, 0, 0, 91, 100, 98, 99, 102,
	128, 134, 134, 460, 0, 0, 0, 0, 103, 104,
	0, 0, 128, 0, 128, 92, 93, 0, 128, 67,
	131, 73, 132, 118, 463, 128, 129, 0, 294, 0,
	91, 0, 103, 104, 128, 0, 0, 0, 134, 92,
	93, 0, 134, 134, 0, 0, 0, 0, 0, 75,
	128, 128, 83, 84, 91, 94, 0, 76, 77, 78,
	79, 80, 0, 0, 128, 128, 134, 0, 0, 0,
	0, 0, 128, 0, 0, 240, 0, 130, 0, 82,
	81, 0, 0, 134, 103, 104, 0, 0, 0, 0,
	128, 92, 93, 0, 0, 134, 0, 134, 0, 0,
	0, 134, 0, 0, 0, 0, 91, 0, 134, 0,
	0, 0, 0, 0, 0, 0, 0, 134, 0, 67,
	131, 73, 132, 118, 128, 124, 129, 0, 128, 0,
	128, 0, 0, 134, 134, 0, 0, 9, 0, 0,
	0, 0, 0, 128, 0, 0, 0, 134, 134, 75,
	0, 0, 83, 84, 351, 134, 122, 76, 77, 78,
	79, 80, 0, 123, 0, 0, 0, 0, 0, 0,
	128, 128, 0, 134, 0, 121, 0, 130, 0, 82,
	81, 128, 126, 103, 104, 0, 0, 0, 0, 0,
	92, 93, 0, 0, 0, 0, 0, 0, 0, 154,
	0, 156, 154, 0, 0, 91, 0, 134, 164, 0,
	0, 134, 0, 134, 0, 0, 175, 176, 177, 0,
	48, 0, 0, 0, 0, 0, 134, 0, 185, 0,
	187, 188, 0, 190, 0, 192, 193, 194, 195, 0,
	197, 0, 0, 200, 0, 202, 204, 0, 0, 0,
	0, 0, 0, 134, 134, 0, 222, 0, 0, 226,
	229, 232, 0, 0, 134, 135, 0, 0, 0, 126,
	0, 0, 0, 0, 222, 0, 0, 244, 0, 0,
	0, 0, 135, 0, 135, 135, 0, 0, 0, 0,
	0, 135, 0, 0, 0, 0, 256, 0, 0, 135,
	135, 135, 0, 0, 273, 279, 0, 222, 0, 0,
	0, 135, 0, 135, 135, 0, 135, 0, 135, 135,
	135, 135, 0, 135, 126, 0, 135, 293, 135, 135,
	0, 0, 0, 0, 0, 0, 295, 296, 0, 135,
	0, 0, 135, 135, 135, 0, 0, 0, 0, 0,
	152, 0, 135, 0, 0, 0, 0, 135, 0, 0,
	135, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 319, 0, 0, 0, 279, 327, 0,
	0, 0, 0, 0, 0, 0, 0, 135, 135, 0,
	135, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 341, 0, 0, 0, 0, 0, 135, 0, 0,
	135, 0, 218, 0, 220, 0, 0, 0, 126, 135,
	135, 0, 0, 0, 0, 0, 236, 0, 0, 0,
	222, 0, 357, 0, 0, 0, 319, 67, 131, 73,
	132, 118, 0, 365, 129, 0, 0, 0, 0, 0,
	0, 0, 371, 0, 0, 0, 135, 0, 0, 0,
	135, 135, 0, 0, 0, 0, 0, 75, 383, 384,
	83, 84, 0, 0, 0, 76, 77, 78, 79, 80,
	0, 0, 154, 406, 135, 352, 0, 0, 0, 0,
	415, 0, 0, 240, 0, 130, 0, 82, 81, 0,
	0, 135, 0, 0, 0, 0, 301, 0, 406, 0,
	0, 0, 0, 135, 308, 135, 0, 0, 0, 135,
	0, 0, 0, 0, 0, 0, 135, 0, 0, 0,
	0, 0, 0, 0, 322, 135, 326, 67, 131, 73,
	132, 292, 154, 0, 129, 0, 445, 0, 446, 0,
	0, 135, 135, 0, 336, 337, 0, 0, 0, 0,
	0, 445, 0, 0, 0, 135, 135, 75, 0, 0,
	83, 84, 0, 135, 291, 76, 77, 78, 79, 80,
	0, 0, 0, 326, 0, 0, 0, 0, 126, 464,
	0, 135, 0, 70, 0, 130, 0, 82, 81, 471,
	0, 0, 0, 0, 0, 0, 0, 368, 0, 67,
	280, 73, 132, 74, 0, 0, 0, 0, 0, 0,
	0, 0, 379, 0, 0, 135, 0, 0, 0, 135,
	0, 135, 389, 0, 0, 0, 0, 393, 0, 75,
	0, 0, 83, 84, 135, 0, 0, 76, 77, 78,
	79, 80, 0, 0, 0, 274, 0, 0, 0, 420,
	238, 0, 0, 0, 0, 70, 0, 71, 275, 82,
	81, 135, 135, 0, 67, 131, 73, 132, 74, 0,
	0, 0, 135, 0, 435, 436, 0, 0, 0, 0,
	0, 438, 0, 0, 0, 0, 0, 0, 0, 0,
	442, 0, 444, 0, 75, 0, 0, 83, 84, 0,
	0, 0, 76, 77, 78, 79, 80, 0, 0, 0,
	274, 0, 0, 0, 0, 0, 0, 423, 0, 0,
	70, 452, 71, 0, 82, 81, 0, 0, 0, 0,
	458, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	465, 0, 336, 337, 67, 44, 73, 45, 74, 0,
	0, 0, 41, 477, 400, 476, 475, 401, 42, 43,
	0, 62, 0, 53, 0, 0, 65, 66, 0, 0,
	64, 61, 0, 0, 75, 63, 0, 83, 84, 0,
	0, 0, 76, 77, 78, 79, 80, 0, 0, 0,
	396, 397, 0, 0, 0, 0, 0, 0, 0, 0,
	70, 0, 71, 0, 82, 81, 67, 44, 73, 45,
	74, 0, 0, 0, 41, 473, 400, 476, 475, 401,
	42, 43, 0, 62, 0, 53, 0, 0, 65, 66,
	0, 0, 64, 61, 0, 0, 75, 63, 0, 83,
	84, 0, 0, 0, 76, 77, 78, 79, 80, 0,
	0, 0, 396, 397, 0, 0, 0, 0, 0, 0,
	0, 0, 70, 0, 71, 0, 82, 81, 67, 44,
	73, 45, 74, 0, 0, 0, 41, 390, 51, 335,
	334, 52, 42, 43, 0, 62, 0, 53, 0, 0,
	65, 66, 0, 0, 64, 61, 0, 0, 75, 63,
	0, 83, 84, 0, 0, 0, 76, 77, 78, 79,
	80, 0, 0, 0, 259, 260, 0, 0, 0, 0,
	0, 0, 0, 0, 70, 0, 71, 0, 82, 81,
	67, 44, 73, 45, 74, 0, 0, 0, 41, 332,
	51, 335, 334, 52, 42, 43, 0, 62, 0, 53,
	0, 0, 65, 66, 0, 0, 64, 61, 0, 0,
	75, 63, 0, 83, 84, 0, 0, 0, 76, 77,
	78, 79, 80, 0, 0, 0, 259, 260, 0, 67,
	44, 73, 45, 74, 0, 0, 70, 41, 71, 51,
	82, 81, 52, 42, 43, 0, 62, 0, 53, 0,
	0, 65, 66, 0, 0, 64, 61, 0, 0, 75,
	63, 0, 83, 84, 0, 0, 0, 76, 77, 78,
	79, 80, 0, 0, 0, 6, 7, 0, 0, 0,
	0, 0, 0, 0, 0, 70, 0, 71, 0, 82,
	81, 0, 8, 67, 44, 73, 45, 74, 0, 0,
	0, 41, 448, 51, 0, 0, 52, 42, 43, 0,
	62, 0, 53, 340, 0, 65, 66, 0, 0, 64,
	61, 0, 0, 75, 63, 0, 83, 84, 0, 0,
	0, 76, 77, 78, 79, 80, 0, 0, 0, 259,
	260, 0, 0, 0, 0, 0, 0, 0, 0, 70,
	0, 71, 0, 82, 81, 67, 44, 73, 45, 74,
	0, 0, 0, 41, 366, 51, 0, 0, 52, 42,
	43, 0, 62, 0, 53, 340, 0, 65, 66, 0,
	0, 64, 61, 0, 0, 75, 63, 0, 83, 84,
	0, 0, 0, 76, 77, 78, 79, 80, 0, 0,
	0, 259, 260, 0, 0, 0, 0, 0, 0, 0,
	0, 70, 0, 71, 0, 82, 81, 67, 44, 73,
	45, 74, 0, 0, 0, 41, 359, 51, 0, 0,
	52, 42, 43, 0, 62, 0, 53, 340, 0, 65,
	66, 0, 0, 64, 61, 0, 0, 75, 63, 0,
	83, 84, 0, 0, 0, 76, 77, 78, 79, 80,
	0, 0, 0, 259, 260, 0, 0, 0, 0, 0,
	0, 0, 0, 70, 0, 71, 0, 82, 81, 67,
	44, 73, 45, 74, 0, 0, 0, 41, 479, 400,
	0, 0, 401, 42, 43, 0, 62, 0, 53, 0,
	0, 65, 66, 0, 0, 64, 61, 0, 0, 75,
	63, 0, 83, 84, 0, 0, 0, 76, 77, 78,
	79, 80, 0, 0, 0, 396, 397, 0, 0, 0,
	0, 0, 0, 0, 0, 70, 0, 71, 0, 82,
	81, 67, 44, 73, 45, 74, 0, 0, 0, 41,
	457, 51, 0, 0, 52, 42, 43, 0, 62, 0,
	53, 0, 0, 65, 66, 0, 0, 64, 61, 0,
	0, 75, 63, 0, 83, 84, 0, 0, 0, 76,
	77, 78, 79, 80, 0, 0, 0, 259, 260, 0,
	67, 44, 73, 45, 74, 0, 0, 70, 41, 71,
	51, 82, 81, 52, 42, 43, 0, 62, 0, 53,
	0, 0, 65, 66, 0, 0, 64, 61, 0, 0,
	75, 63, 0, 83, 84, 0, 0, 0, 76, 77,
	78, 79, 80, 0, 0, 0, 259, 260, 0, 0,
	0, 0, 0, 0, 0, 0, 70, 0, 71, 429,
	82, 81, 67, 44, 73, 45, 74, 0, 0, 0,
	41, 422, 51, 0, 0, 52, 42, 43, 0, 62,
	0, 53, 0, 0, 65, 66, 0, 0, 64, 61,
	0, 0, 75, 63, 0, 83, 84, 0, 0, 0,
	76, 77, 78, 79, 80, 0, 0, 0, 259, 260,
	0, 0, 0, 0, 0, 0, 0, 0, 70, 0,
	71, 0, 82, 81, 67, 44, 73, 45, 74, 0,
	0, 0, 41, 395, 400, 0, 0, 401, 42, 43,
	0, 62, 0, 53, 0, 0, 65, 66, 0, 0,
	64, 61, 0, 0, 75, 63, 0, 83, 84, 0,
	0, 0, 76, 77, 78, 79, 80, 0, 0, 0,
	396, 397, 0, 0, 0, 0, 0, 0, 0, 0,
	70, 0, 71, 0, 82, 81, 67, 44, 73, 45,
	74, 0, 0, 0, 41, 392, 51, 0, 0, 52,
	42, 43, 0, 62, 0, 53, 0, 0, 65, 66,
	0, 0, 64, 61, 0, 0, 75, 63, 0, 83,
	84, 0, 0, 0, 76, 77, 78, 79, 80, 0,
	0, 0, 259, 260, 0, 0, 0, 0, 0, 0,
	0, 0, 70, 0, 71, 0, 82, 81, 67, 44,
	73, 45, 74, 0, 0, 0, 41, 376, 51, 0,
	0, 52, 42, 43, 0, 62, 0, 53, 0, 0,
	65, 66, 0, 0, 64, 61, 0, 0, 75, 63,
	0, 83, 84, 0, 0, 0, 76, 77, 78, 79,
	80, 0, 0, 0, 259, 260, 0, 0, 0, 0,
	0, 0, 0, 0, 70, 0, 71, 0, 82, 81,
	67, 44, 73, 45, 74, 0, 0, 0, 41, 321,
	51, 0, 0, 52, 42, 43, 0, 62, 0, 53,
	0, 0, 65, 66, 0, 0, 64, 61, 0, 0,
	75, 63, 0, 83, 84, 0, 0, 0, 76, 77,
	78, 79, 80, 0, 0, 0, 259, 260, 0, 0,
	0, 0, 0, 0, 0, 0, 70, 0, 71, 0,
	82, 81, 67, 44, 73, 45, 74, 0, 0, 0,
	41, 311, 51, 0, 0, 52, 42, 43, 0, 62,
	0, 53, 0, 0, 65, 66, 0, 0, 64, 61,
	0, 0, 75, 63, 0, 83, 84, 0, 0, 0,
	76, 77, 78, 79, 80, 0, 0, 0, 259, 260,
	0, 0, 0, 0, 0, 0, 0, 0, 70, 0,
	71, 0, 82, 81, 67, 44, 73, 45, 74, 0,
	0, 0, 41, 309, 51, 0, 0, 52, 42, 43,
	0, 62, 0, 53, 0, 0, 65, 66, 0, 0,
	64, 61, 0, 0, 75, 63, 0, 83, 84, 0,
	0, 0, 76, 77, 78, 79, 80, 0, 0, 0,
	259, 260, 0, 67, 44, 73, 45, 74, 0, 0,
	70, 41, 71, 400, 82, 81, 401, 42, 43, 0,
	62, 0, 53, 0, 0, 65, 66, 0, 0, 64,
	61, 0, 0, 75, 63, 0, 83, 84, 0, 0,
	0, 76, 77, 78, 79, 80, 0, 0, 0, 396,
	397, 0, 67, 44, 73, 45, 74, 0, 0, 70,
	41, 71, 51, 82, 81, 52, 42, 43, 0, 62,
	0, 53, 0, 0, 65, 66, 0, 0, 64, 61,
	0, 0, 75, 63, 0, 83, 84, 0, 0, 0,
	76, 77, 78, 79, 80, 0, 0, 0, 259, 260,
	0, 67, 44, 73, 45, 74, 285, 0, 70, 41,
	71, 51, 82, 81, 52, 42, 43, 0, 62, 0,
	53, 0, 0, 65, 66, 0, 0, 64, 61, 0,
	0, 75, 63, 0, 83, 84, 0, 0, 0, 76,
	77, 78, 79, 80, 0, 0, 0, 0, 284, 0,
	67, 44, 73, 45, 74, 0, 0, 70, 41, 71,
	51, 82, 81, 52, 42, 43, 0, 62, 0, 53,
	0, 0, 65, 66, 0, 0, 64, 61, 0, 0,
	75, 63, 0, 83, 84, 0, 0, 0, 76, 77,
	78, 79, 80, 67, 131, 73, 132, 74, 0, 0,
	0, 0, 0, 0, 0, 0, 70, 0, 71, 0,
	82, 81, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 75, 0, 0, 83, 84, 0, 0,
	0, 76, 77, 78, 79, 80, 0, 0, 0, 274,
	67, 131, 73, 132, 74, 0, 408, 0, 0, 70,
	0, 71, 0, 82, 81, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	75, 0, 0, 83, 84, 0, 0, 0, 76, 77,
	78, 79, 80, 0, 0, 0, 274, 67, 280, 73,
	132, 74, 0, 372, 0, 0, 70, 0, 71, 0,
	82, 81, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 75, 0, 0,
	83, 84, 0, 0, 0, 76, 77, 78, 79, 80,
	0, 0, 0, 274, 67, 131, 73, 132, 74, 0,
	0, 0, 0, 70, 0, 71, 275, 82, 81, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 75, 0, 0, 83, 84, 0,
	0, 0, 76, 77, 78, 79, 80, 0, 0, 0,
	274, 67, 131, 73, 132, 74, 0, 0, 129, 0,
	70, 0, 71, 0, 82, 81, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 75, 0, 0, 83, 84, 0, 0, 0, 76,
	77, 78, 79, 80, 67, 131, 73, 132, 74, 0,
	0, 0, 0, 0, 0, 0, 0, 70, 0, 130,
	0, 82, 81, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 75, 0, 0, 83, 84, 0,
	0, 0, 76, 77, 78, 79, 80, 67, 131, 73,
	132, 118, 0, 0, 129, 0, 0, 166, 0, 0,
	70, 0, 71, 0, 82, 81, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 75, 0, 0,
	83, 84, 0, 0, 0, 76, 77, 78, 79, 80,
	67, 450, 73, 132, 74, 0, 0, 0, 0, 0,
	0, 0, 0, 240, 0, 130, 0, 82, 81, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	75, 0, 0, 83, 84, 94, 0, 0, 76, 77,
	78, 79, 80, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 94, 70, 0, 71, 0,
	82, 81, 0, 0, 103, 104, 0, 0, 0, 0,
	0, 92, 93, 0, 0, 0, 95, 96, 97, 105,
	0, 0, 0, 0, 103, 104, 91, 100, 98, 99,
	102, 92, 93, 381, 0, 0, 95, 96, 97, 105,
	0, 0, 0, 0, 0, 0, 91, 100, 98, 99,
	102, 0, 0, 315,
}
var RubyPact = []int{

	-28, 1704, -1000, -1000, -1000, 7, -1000, -1000, -1000, 221,
	-1000, -1000, -1000, 108, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	159, -1000, 25, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 333, 308, 308, 834, 239, -10, 139, 130, 175,
	169, 2695, 2695, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 399, 2695, 399, 399, 312, 245, -1000, -1000, -1000,
	2969, -1000, 27, 303, -1000, 1, 2695, 2695, 399, 399,
	399, 20, 328, -1000, -1000, -1000, -1000, -1000, 2695, 2695,
	399, 374, 399, 399, -1000, 399, 2695, 399, 399, 399,
	399, 2695, 399, -1000, -1000, 399, 2695, 399, 399, 2695,
	2695, 2695, 372, 187, 212, 85, -1000, -1000, 2969, 1,
	37, 2969, 399, 399, 370, 247, 771, -1000, 6, 14,
	-1000, 3012, 256, -1, -1000, -1000, 2969, 2695, 2695, 399,
	2695, 2695, 2695, 2695, 2695, 2695, 2695, 367, 366, 365,
	174, 106, 2597, 165, 771, 122, 771, 165, 2695, 2695,
	2695, 2695, 46, 43, 660, -1000, 399, 2832, 325, 2969,
	2646, -1000, -1000, 174, 174, 771, 771, 771, -1000, -1000,
	229, -1000, -1000, 174, 174, 771, 1252, 771, 771, 2879,
	771, 174, 771, 771, 771, 771, 174, 695, 2879, 2879,
	771, 174, 771, 39, 510, 174, 174, 174, 1, -1000,
	361, 242, 35, -1000, 73, 360, 357, -1000, 2499, 308,
	2437, 321, 660, -1000, -1000, -1000, 3111, -47, 26, 549,
	-1000, -1000, 643, -1000, -1000, 2926, 2375, -1000, 344, 1324,
	2969, 311, 174, 174, 324, 174, 174, 174, 174, 174,
	174, 174, -1000, 236, -18, -23, 1655, -1000, -1000, -1000,
	-1000, 174, 226, 399, -1000, 66, 174, 174, 174, 174,
	-1000, -1000, -1000, 771, -1000, -1000, 220, 203, 6, 870,
	1152, -1000, 290, 174, -1000, -1000, 23, -1000, -1000, 37,
	1, 2695, 2969, 771, 399, 771, 771, -1000, 2926, 72,
	-1000, 1892, 212, 35, 180, 399, -1000, -1000, 1830, -1000,
	-1000, -1000, 1, -1000, 2785, 152, -1000, -1000, 18, 771,
	-1000, -1000, 2313, 50, -1000, -1000, 2597, 3091, -1000, 67,
	399, 399, -1000, 301, 2695, -1000, 1593, 2251, -1000, -1000,
	285, 771, 2189, 259, 399, 2738, -17, -1000, -20, -1000,
	-22, 2695, 399, -1000, -1000, 174, 124, 771, 2695, -1000,
	224, -1000, -1000, -1000, -1000, 771, -1000, 208, 2127, -1000,
	1389, 771, 341, 2695, 340, 335, -1000, -1000, 323, 2065,
	-31, 54, 2695, 320, 58, -1000, 2695, -1000, 174, 2597,
	-1000, 291, -1000, 2597, 377, -1000, -1000, -1000, 174, -1000,
	2695, 2695, -1000, -1000, 399, 165, 660, -1000, 399, -1000,
	2879, -1000, 79, -1000, 174, 771, -1000, 174, -1000, -1000,
	1768, -1000, -1000, 3055, -1000, 174, -26, -1000, -1000, -1000,
	-1000, 174, 198, -1000, 174, 2597, 2597, -1000, 2597, 317,
	114, 104, 2016, 165, 2597, 771, 719, 31, -1000, 205,
	734, 399, 2597, -1000, -1000, -1000, -1000, -1000, 2597, 44,
	2695, 399, -1000, -1000, 32, 2597, 1531, 1469, 1954, 44,
	284, 315, -1000, -1000, 273, 2695, -1000, -1000, 268, -1000,
	-1000, -1000, 44, -1000, -1000, 2695, -1000, 174, 2548, -1000,
	44, 174, 2548, 2548, 2548,
}
var RubyPgo = []int{

	0, 0, 480, 478, 11, 7, 477, 474, 473, 940,
	472, 1, 27, 470, 463, 459, 458, 857, 456, 522,
	439, 455, 454, 453, 452, 448, 447, 45, 441, 10,
	104, 440, 433, 12, 432, 428, 427, 426, 30, 423,
	422, 3, 421, 420, 418, 417, 415, 414, 413, 412,
	409, 400, 398, 1016, 396, 4, 19, 18, 5, 395,
	8, 394, 178, 393, 17, 392, 6, 390, 9, 21,
	14, 16, 389, 384, 152,
}
var RubyR1 = []int{

	0, 59, 59, 59, 59, 59, 59, 59, 59, 59,
	59, 73, 73, 74, 74, 53, 53, 53, 53, 18,
	18, 18, 18, 18, 18, 18, 18, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 27, 27, 27, 27, 27, 27, 27, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 38, 14, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 21, 56, 56, 56, 56,
	66, 66, 64, 64, 64, 64, 64, 64, 64, 70,
	70, 70, 70, 70, 68, 68, 68, 22, 22, 22,
	22, 22, 22, 60, 60, 71, 71, 71, 29, 29,
	29, 29, 28, 28, 31, 33, 33, 72, 72, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 69, 69,
	32, 32, 32, 32, 32, 32, 32, 9, 9, 30,
	30, 19, 19, 42, 42, 43, 44, 45, 46, 47,
	48, 49, 50, 51, 52, 2, 6, 7, 7, 3,
	3, 3, 3, 61, 61, 67, 67, 67, 5, 5,
	5, 5, 57, 65, 65, 65, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 58, 58, 58,
	58, 54, 54, 54, 8, 16, 11, 11, 11, 63,
	63, 55, 55, 23, 23, 24, 24, 26, 26, 26,
	25, 25, 25, 25, 12, 39, 62, 62, 62, 62,
	62, 40, 40, 40, 40, 40, 41, 41, 41, 41,
	37, 36, 10, 35, 35, 34, 34, 4,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 1, 0, 2, 0, 2, 2, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 2, 4, 5,
	1, 4, 4, 2, 3, 3, 4, 4, 5, 3,
	4, 5, 2, 3, 3, 3, 4, 4, 4, 4,
	4, 4, 6, 6, 6, 4, 3, 7, 1, 5,
	1, 3, 0, 1, 1, 2, 4, 4, 5, 1,
	1, 4, 2, 5, 1, 3, 3, 5, 6, 7,
	8, 5, 6, 1, 3, 0, 1, 3, 1, 2,
	3, 2, 4, 6, 4, 1, 3, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 9, 6,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 2,
	2, 3, 3, 3, 3, 2, 2, 2, 2, 3,
	3, 3, 3, 3, 3, 1, 1, 3, 3, 3,
	5, 5, 5, 3, 7, 3, 7, 8, 3, 4,
	5, 5, 3, 0, 1, 3, 4, 5, 3, 3,
	3, 3, 3, 5, 6, 5, 3, 4, 3, 3,
	2, 0, 2, 2, 3, 4, 2, 3, 5, 0,
	2, 1, 2, 2, 1, 2, 1, 1, 3, 3,
	0, 1, 3, 3, 5, 5, 0, 2, 2, 2,
	2, 5, 6, 5, 6, 5, 4, 3, 3, 2,
	4, 4, 2, 5, 7, 4, 5, 3,
}
var RubyChk = []int{

	-1000, -59, 51, 52, 68, -1, 51, 52, 68, -17,
	-22, -28, -31, -15, -32, -13, -16, -27, -23, -39,
	-35, -24, -25, -26, -38, -4, -18, -7, -3, -33,
	-20, -21, -8, -10, -43, -44, -45, -46, -14, -37,
	-36, 13, 19, 20, 6, 8, -30, -19, -9, -42,
	-69, 15, 18, 24, -47, -48, -49, -50, -51, -52,
	-12, 32, 22, 36, 31, 27, 28, 5, -2, -6,
	61, 63, -72, 7, 9, 35, 43, 44, 45, 46,
	47, 66, 65, 38, 39, 52, 51, 68, 15, 18,
	25, 55, 40, 41, 4, 45, 46, 47, 57, 58,
	56, 18, 59, 33, 34, 48, 18, 40, 61, 49,
	15, 18, 55, 6, 4, -33, 8, -33, 9, -56,
	-70, 61, 42, 49, 11, -68, -17, -5, -20, 12,
	63, 6, 8, -30, -19, -9, 9, 42, 49, 61,
	42, 49, 42, 49, 42, 49, 42, 11, 42, 11,
	-1, -1, -53, -66, -17, -1, -17, -66, 15, 18,
	15, 18, -66, -64, -17, -27, 58, -74, 54, 9,
	-54, -5, 63, -1, -1, -17, -17, -17, 6, 8,
	66, 6, 8, -1, -1, -17, 6, -17, -17, -74,
	-17, -1, -17, -17, -17, -17, -1, -17, -74, -74,
	-17, -1, -17, -68, -17, -1, -1, -1, 6, -60,
	55, -71, 9, -29, 6, 47, 58, -60, -53, 40,
	-53, -64, -17, -5, 11, -5, -17, -4, -68, -17,
	-38, -12, -17, -12, 6, 11, -53, -57, 56, -74,
	61, -64, -1, -1, -17, -1, -1, -1, -1, -1,
	-1, -1, 6, -69, 6, 6, -53, 51, 52, 51,
	52, -1, -63, 11, 51, -74, -1, -1, -1, -1,
	62, 11, 62, -17, 51, 64, -61, -67, -20, -17,
	6, 8, -64, -1, 52, 10, -74, 6, 8, -70,
	-56, 42, 9, -17, 53, -17, -17, 62, 11, 62,
	-5, -53, 6, 11, -71, 42, 6, 6, -53, 14,
	-33, 14, 10, 11, -74, 62, 62, 62, -74, -17,
	-5, 14, -53, -65, 6, -57, -53, -17, 10, 62,
	61, 61, 14, -58, 17, 16, -53, -53, 14, -11,
	25, -17, -62, -34, 37, -74, -74, 11, -74, 11,
	-74, 4, 53, 10, -5, -1, -64, -17, 42, 14,
	-55, -11, -60, -29, 10, -17, 14, -55, -53, -5,
	-74, -17, 58, 42, 11, 58, 14, 56, 11, -53,
	-74, 62, 42, -17, -17, 14, 17, 16, -1, -53,
	14, -58, 14, -53, 8, 14, 51, 52, -1, -40,
	15, 18, 14, 16, 37, -66, -17, -27, 58, 64,
	-74, 64, -74, 64, -1, -17, 10, -1, 14, -11,
	-53, 14, 14, 58, 6, -1, 6, 6, 6, 64,
	64, -1, 62, 62, -1, -53, -53, 14, -53, 4,
	-1, -1, -53, -66, -53, -17, -17, 6, 14, -55,
	6, 61, -53, 6, 51, 51, 52, 14, -53, -74,
	4, 53, 14, 10, -17, -53, -62, -62, -62, -74,
	-1, -17, 62, 14, -41, 17, 16, 14, -41, 14,
	-73, 11, -74, 11, 14, 17, 16, -1, -62, 14,
	-74, -1, -62, -62, -62,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 49,
	50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
	60, 61, 62, 63, 64, 65, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 0, 0, 0, 20, -2, 22, 23, 24, 0,
	0, 230, 230, 15, 42, 43, 44, 45, 46, 47,
	48, 224, 230, 0, 226, 231, 227, 19, 25, 26,
	102, 13, 0, 70, 211, 0, 230, 230, 0, 0,
	0, 0, 0, 175, 176, 5, 6, 7, 230, 230,
	0, 0, 0, 0, 13, 0, 230, 0, 0, 0,
	0, 230, 0, 13, 13, 0, 230, 0, 0, 230,
	230, 230, 0, 125, 125, 15, -2, 15, -2, 73,
	82, 102, 0, 0, 0, 98, 109, 110, 31, 15,
	13, 20, -2, 22, 23, 24, 102, 230, 230, 0,
	230, 230, 230, 230, 230, 230, 230, 0, 0, 0,
	15, 0, 219, 223, 100, 0, 13, 225, 230, 230,
	230, 230, 0, 0, 100, 104, 0, 0, 0, 102,
	0, 252, 13, 165, 166, 167, 168, 67, 159, 160,
	0, 157, 158, 198, 206, 66, 75, 83, 85, 0,
	169, 170, 171, 172, 173, 174, 200, 0, 0, 0,
	257, 202, 84, 0, 114, 156, 199, 201, 79, 15,
	0, 123, 125, 126, 128, 0, 0, 15, 0, 0,
	0, 0, 103, 74, 13, 112, 100, 0, 0, 139,
	140, 141, 150, 151, 163, 13, 0, 15, 193, 15,
	102, 0, 142, 152, 0, 143, 153, 144, 154, 145,
	155, 146, 164, 147, 0, 0, 0, 15, 15, 16,
	17, 18, 0, 0, 236, 0, 232, 233, 228, 229,
	177, 13, 178, 105, 14, 179, 13, 13, -2, 0,
	20, -2, 0, 212, 213, 214, 15, 161, 162, 76,
	77, 230, -2, 95, 0, 250, 251, 90, 0, 91,
	80, 0, 125, 0, 0, 0, 129, 131, 0, 132,
	15, 134, 68, 13, 0, 86, 88, 89, 0, 115,
	116, 188, 0, 0, 194, 15, 13, 100, 72, 87,
	0, 0, 196, 0, 230, 15, 0, 0, 215, 220,
	15, 101, 0, 0, 0, 0, 0, 13, 0, 13,
	0, -2, 0, 71, 78, 81, 0, 234, 230, 117,
	0, 221, 15, 127, 124, 130, 121, 0, 0, 69,
	0, 111, 0, 230, 0, 0, 189, 192, 0, 0,
	0, 86, 230, 0, 0, 197, 230, 15, 15, 210,
	203, 0, 205, 216, 15, 235, 237, 238, 239, 240,
	230, 230, 253, 15, 0, 15, 106, 107, 0, 180,
	0, 181, 0, 182, 183, 185, 96, 94, 118, 222,
	0, 122, 133, 0, 113, 92, 0, 99, 195, 190,
	191, 93, 0, 149, 15, 208, 209, 204, 217, 0,
	15, 0, 0, 15, 13, 108, 0, 0, 119, 0,
	20, 0, 207, 15, 236, 15, 15, 254, 13, 255,
	-2, 0, 120, 97, 0, 218, 0, 0, 0, 256,
	11, 13, 148, 241, 0, 230, 236, 243, 0, 245,
	184, 12, 186, 13, 242, 230, 236, 236, 230, 244,
	187, 236, 230, 230, 230,
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
		//line parser.y:204
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:206
		{
		}
	case 3:
		//line parser.y:208
		{
		}
	case 4:
		//line parser.y:210
		{
		}
	case 5:
		//line parser.y:212
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:214
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:216
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:222
		{
		}
	case 11:
		//line parser.y:224
		{
		}
	case 12:
		//line parser.y:225
		{
		}
	case 13:
		//line parser.y:227
		{
		}
	case 14:
		//line parser.y:228
		{
		}
	case 15:
		//line parser.y:231
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 16:
		//line parser.y:233
		{
		}
	case 17:
		//line parser.y:235
		{
		}
	case 18:
		//line parser.y:237
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
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 64:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 65:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 66:
		//line parser.y:249
		{
			RubyVAL.genericValue = ast.RescueModifier{Statement: RubyS[Rubypt-2].genericValue, Rescue: RubyS[Rubypt-0].genericValue}
		}
	case 67:
		//line parser.y:252
		{
			RubyVAL.genericValue = ast.StarSplat{Value: RubyS[Rubypt-0].genericValue}
		}
	case 68:
		//line parser.y:255
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 69:
		//line parser.y:262
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-2].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 70:
		//line parser.y:269
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 71:
		//line parser.y:273
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 72:
		//line parser.y:280
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 73:
		//line parser.y:287
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 74:
		//line parser.y:294
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 75:
		//line parser.y:301
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 76:
		//line parser.y:308
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 77:
		//line parser.y:316
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 78:
		//line parser.y:324
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 79:
		//line parser.y:332
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 80:
		//line parser.y:340
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 81:
		//line parser.y:348
		{
			methodName := RubyS[Rubypt-2].genericValue.(ast.BareReference).Name + "="
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: methodName},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 82:
		//line parser.y:359
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 83:
		//line parser.y:366
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 84:
		//line parser.y:374
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 85:
		//line parser.y:382
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 86:
		//line parser.y:392
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 87:
		//line parser.y:400
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 88:
		//line parser.y:408
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 89:
		//line parser.y:416
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 90:
		//line parser.y:424
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 91:
		//line parser.y:432
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 92:
		//line parser.y:442
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 93:
		//line parser.y:450
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 94:
		//line parser.y:458
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 95:
		//line parser.y:468
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 96:
		//line parser.y:478
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 97:
		//line parser.y:480
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 98:
		//line parser.y:482
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 99:
		//line parser.y:484
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 100:
		//line parser.y:487
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 101:
		//line parser.y:489
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 102:
		//line parser.y:491
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 103:
		//line parser.y:493
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 104:
		//line parser.y:495
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 105:
		//line parser.y:497
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{Func: ast.BareReference{Name: "to_proc"}, Target: RubyS[Rubypt-0].genericValue})
		}
	case 106:
		//line parser.y:499
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 107:
		//line parser.y:501
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 108:
		//line parser.y:503
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.CallExpression{
				Func:   ast.BareReference{Name: "to_proc"},
				Target: RubyS[Rubypt-0].genericValue,
			})
		}
	case 109:
		//line parser.y:512
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 110:
		//line parser.y:514
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 111:
		//line parser.y:516
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 112:
		//line parser.y:518
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 113:
		//line parser.y:520
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 114:
		//line parser.y:523
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 115:
		//line parser.y:525
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 116:
		//line parser.y:527
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 117:
		//line parser.y:531
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 118:
		//line parser.y:539
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 119:
		//line parser.y:548
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 120:
		//line parser.y:557
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 121:
		//line parser.y:567
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 122:
		//line parser.y:575
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 123:
		//line parser.y:586
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 124:
		//line parser.y:588
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 125:
		//line parser.y:590
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 126:
		//line parser.y:592
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 127:
		//line parser.y:594
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 128:
		//line parser.y:597
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 129:
		//line parser.y:599
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 130:
		//line parser.y:601
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 131:
		//line parser.y:603
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 132:
		//line parser.y:607
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 133:
		//line parser.y:614
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 134:
		//line parser.y:624
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 135:
		//line parser.y:633
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 136:
		//line parser.y:639
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 137:
		//line parser.y:647
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 138:
		//line parser.y:651
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 139:
		//line parser.y:656
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 140:
		//line parser.y:663
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 141:
		//line parser.y:667
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 142:
		//line parser.y:671
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 143:
		//line parser.y:678
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 144:
		//line parser.y:685
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 145:
		//line parser.y:692
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 146:
		//line parser.y:699
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 147:
		//line parser.y:706
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: ast.Array{Nodes: RubyS[Rubypt-2].genericSlice},
				RHS: ast.Array{Nodes: RubyS[Rubypt-0].genericSlice},
			}
		}
	case 148:
		//line parser.y:714
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
	case 149:
		//line parser.y:729
		{
			tail := ast.CallExpression{Target: RubyS[Rubypt-3].genericValue, Func: ast.BareReference{Name: "[]="}, Args: []ast.Node{RubyS[Rubypt-1].genericValue}}
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, tail)
		}
	case 150:
		//line parser.y:735
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 151:
		//line parser.y:742
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 152:
		//line parser.y:746
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 153:
		//line parser.y:753
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 154:
		//line parser.y:760
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 155:
		//line parser.y:767
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 156:
		//line parser.y:774
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 157:
		//line parser.y:777
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 158:
		//line parser.y:779
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 159:
		//line parser.y:782
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 160:
		//line parser.y:784
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 161:
		//line parser.y:787
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 162:
		//line parser.y:789
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 163:
		//line parser.y:792
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 164:
		//line parser.y:794
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 165:
		//line parser.y:796
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 166:
		//line parser.y:797
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 167:
		//line parser.y:798
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 168:
		//line parser.y:799
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 169:
		//line parser.y:802
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 170:
		//line parser.y:811
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 171:
		//line parser.y:820
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 172:
		//line parser.y:829
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 173:
		//line parser.y:838
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 174:
		//line parser.y:847
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 175:
		//line parser.y:855
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 176:
		//line parser.y:856
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 177:
		//line parser.y:858
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 178:
		//line parser.y:860
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 179:
		//line parser.y:863
		{
			RubyVAL.genericValue = ast.Hash{}
		}
	case 180:
		//line parser.y:865
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 181:
		//line parser.y:873
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 182:
		//line parser.y:881
		{
			RubyVAL.genericValue = ast.Block{Body: ast.Nodes{RubyS[Rubypt-2].genericValue}}
		}
	case 183:
		//line parser.y:884
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 184:
		//line parser.y:891
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 185:
		//line parser.y:899
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 186:
		//line parser.y:906
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 187:
		//line parser.y:913
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 188:
		//line parser.y:921
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 189:
		//line parser.y:923
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 190:
		//line parser.y:930
		{
			RubyVAL.genericValue = ast.Block{Args: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 191:
		//line parser.y:934
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-2].genericSlice}
		}
	case 192:
		//line parser.y:937
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 193:
		//line parser.y:939
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 194:
		//line parser.y:941
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 195:
		//line parser.y:943
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 196:
		//line parser.y:946
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 197:
		//line parser.y:953
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 198:
		//line parser.y:961
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 199:
		//line parser.y:968
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 200:
		//line parser.y:975
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 201:
		//line parser.y:982
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 202:
		//line parser.y:989
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 203:
		//line parser.y:996
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 204:
		//line parser.y:1003
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 205:
		//line parser.y:1011
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 206:
		//line parser.y:1018
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 207:
		//line parser.y:1027
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 208:
		//line parser.y:1034
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 209:
		//line parser.y:1041
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 210:
		//line parser.y:1048
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 211:
		//line parser.y:1055
		{
		}
	case 212:
		//line parser.y:1056
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 213:
		//line parser.y:1057
		{
		}
	case 214:
		//line parser.y:1060
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 215:
		//line parser.y:1063
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 216:
		//line parser.y:1071
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 217:
		//line parser.y:1073
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 218:
		//line parser.y:1082
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
	case 219:
		//line parser.y:1097
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 220:
		//line parser.y:1099
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 221:
		//line parser.y:1102
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 222:
		//line parser.y:1104
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 223:
		//line parser.y:1107
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 224:
		//line parser.y:1114
		{
			RubyVAL.genericValue = ast.Yield{}
		}
	case 225:
		//line parser.y:1117
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 226:
		//line parser.y:1125
		{
			RubyVAL.genericValue = ast.Return{}
		}
	case 227:
		//line parser.y:1129
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 228:
		//line parser.y:1131
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 229:
		//line parser.y:1133
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 230:
		//line parser.y:1136
		{
		}
	case 231:
		//line parser.y:1138
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 232:
		//line parser.y:1140
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 233:
		//line parser.y:1142
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 234:
		//line parser.y:1146
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 235:
		//line parser.y:1155
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 236:
		//line parser.y:1158
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 237:
		//line parser.y:1160
		{
		}
	case 238:
		//line parser.y:1162
		{
		}
	case 239:
		//line parser.y:1164
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 240:
		//line parser.y:1166
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 241:
		//line parser.y:1169
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 242:
		//line parser.y:1176
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 243:
		//line parser.y:1184
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 244:
		//line parser.y:1191
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 245:
		//line parser.y:1199
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 246:
		//line parser.y:1207
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 247:
		//line parser.y:1214
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 248:
		//line parser.y:1221
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 249:
		//line parser.y:1228
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 250:
		//line parser.y:1236
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 251:
		//line parser.y:1239
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-3].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 252:
		//line parser.y:1241
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 253:
		//line parser.y:1244
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 254:
		//line parser.y:1246
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 255:
		//line parser.y:1249
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 256:
		//line parser.y:1251
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 257:
		//line parser.y:1253
		{
			RubyVAL.genericValue = ast.Range{Start: RubyS[Rubypt-2].genericValue, End: RubyS[Rubypt-0].genericValue}
		}
	}
	goto Rubystack /* stack new state and value */
}
