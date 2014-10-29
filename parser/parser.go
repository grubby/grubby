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

//line parser.y:1156

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 29,
	53, 128,
	-2, 22,
	-1, 91,
	10, 85,
	11, 85,
	-2, 196,
	-1, 102,
	53, 128,
	-2, 22,
	-1, 108,
	13, 15,
	15, 15,
	18, 15,
	19, 15,
	20, 15,
	22, 15,
	24, 15,
	32, 15,
	36, 15,
	43, 15,
	44, 15,
	45, 15,
	46, 15,
	51, 15,
	-2, 13,
	-1, 126,
	53, 128,
	-2, 126,
	-1, 237,
	53, 129,
	-2, 127,
	-1, 245,
	10, 85,
	11, 85,
	-2, 196,
	-1, 398,
	27, 212,
	28, 212,
	-2, 15,
	-1, 399,
	27, 212,
	28, 212,
	-2, 15,
}

const RubyNprod = 242
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 2580

var RubyAct = []int{

	9, 201, 419, 273, 315, 261, 309, 205, 146, 37,
	100, 147, 203, 93, 92, 185, 141, 58, 142, 150,
	21, 75, 2, 3, 186, 231, 231, 231, 306, 99,
	75, 231, 122, 303, 184, 75, 151, 408, 354, 4,
	352, 107, 75, 229, 223, 116, 282, 302, 211, 114,
	284, 117, 125, 127, 389, 115, 73, 72, 206, 90,
	138, 204, 140, 231, 265, 398, 399, 148, 137, 75,
	216, 217, 305, 74, 129, 143, 136, 397, 158, 159,
	160, 161, 154, 163, 164, 165, 166, 149, 168, 169,
	170, 155, 173, 230, 228, 175, 176, 178, 231, 207,
	75, 223, 99, 172, 79, 128, 202, 79, 174, 208,
	173, 428, 149, 177, 179, 149, 112, 118, 119, 196,
	197, 187, 113, 183, 409, 206, 209, 410, 204, 289,
	149, 371, 182, 206, 287, 224, 383, 220, 280, 384,
	77, 78, 280, 77, 78, 27, 101, 69, 102, 70,
	210, 235, 212, 173, 76, 359, 248, 76, 99, 215,
	225, 87, 407, 263, 238, 346, 207, 347, 181, 365,
	234, 242, 243, 149, 207, 71, 208, 109, 64, 65,
	280, 109, 252, 220, 208, 278, 79, 220, 348, 79,
	231, 426, 253, 290, 249, 75, 280, 360, 77, 78,
	66, 251, 67, 152, 63, 62, 120, 334, 255, 121,
	258, 220, 76, 220, 110, 268, 220, 264, 320, 263,
	111, 270, 77, 78, 281, 77, 78, 279, 396, 77,
	78, 118, 119, 126, 246, 381, 76, 99, 382, 76,
	237, 276, 277, 76, 390, 367, 173, 296, 368, 285,
	299, 364, 286, 288, 293, 363, 220, 295, 297, 220,
	300, 226, 312, 227, 318, 144, 149, 145, 298, 312,
	434, 319, 431, 430, 124, 323, 123, 220, 220, 429,
	331, 431, 430, 338, 361, 138, 350, 376, 267, 327,
	326, 356, 358, 349, 325, 324, 327, 326, 266, 350,
	329, 292, 229, 260, 229, 333, 351, 353, 262, 355,
	247, 248, 257, 214, 366, 200, 180, 157, 138, 351,
	378, 425, 68, 98, 370, 220, 369, 233, 366, 256,
	220, 222, 232, 1, 220, 153, 57, 56, 55, 54,
	53, 52, 18, 17, 221, 16, 5, 15, 44, 138,
	340, 341, 374, 375, 388, 23, 24, 386, 25, 377,
	26, 283, 314, 14, 12, 11, 316, 22, 10, 20,
	13, 312, 385, 19, 387, 220, 220, 394, 220, 40,
	39, 35, 34, 36, 33, 0, 220, 0, 220, 0,
	130, 131, 132, 133, 134, 135, 220, 0, 395, 366,
	0, 0, 0, 0, 215, 139, 220, 406, 356, 417,
	0, 405, 220, 338, 338, 338, 423, 0, 0, 0,
	156, 411, 30, 276, 277, 415, 162, 0, 0, 0,
	0, 167, 0, 0, 338, 171, 0, 427, 338, 338,
	338, 0, 0, 0, 412, 413, 414, 0, 435, 0,
	0, 103, 0, 0, 0, 188, 189, 190, 191, 192,
	193, 194, 195, 0, 0, 198, 199, 0, 433, 0,
	0, 0, 0, 213, 0, 0, 0, 0, 437, 438,
	0, 0, 103, 439, 103, 0, 79, 0, 0, 103,
	0, 0, 0, 0, 0, 0, 0, 0, 239, 0,
	103, 103, 103, 103, 0, 103, 103, 103, 103, 0,
	103, 103, 103, 0, 103, 0, 0, 103, 103, 103,
	0, 0, 77, 78, 103, 0, 0, 80, 81, 82,
	0, 0, 103, 0, 0, 0, 76, 85, 83, 84,
	87, 103, 103, 250, 0, 0, 0, 0, 0, 0,
	0, 0, 27, 28, 69, 29, 70, 0, 0, 0,
	41, 393, 49, 0, 0, 50, 42, 43, 0, 60,
	0, 51, 280, 103, 0, 103, 0, 0, 317, 59,
	103, 0, 71, 61, 0, 64, 65, 0, 0, 294,
	45, 46, 47, 48, 0, 0, 0, 310, 311, 0,
	0, 0, 0, 0, 103, 0, 313, 66, 0, 67,
	321, 63, 62, 313, 0, 0, 0, 0, 0, 328,
	0, 27, 101, 69, 102, 70, 0, 339, 107, 0,
	0, 0, 0, 0, 0, 357, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 103, 362, 0, 0,
	0, 71, 0, 0, 64, 65, 0, 0, 0, 103,
	0, 0, 0, 0, 0, 0, 231, 0, 103, 103,
	0, 373, 103, 301, 0, 0, 66, 0, 108, 0,
	63, 62, 0, 0, 0, 0, 0, 379, 380, 27,
	101, 69, 102, 91, 0, 0, 107, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 103, 103, 0,
	38, 0, 391, 392, 103, 313, 0, 0, 0, 71,
	0, 103, 64, 65, 0, 0, 400, 401, 402, 403,
	0, 0, 0, 0, 0, 0, 291, 0, 0, 106,
	103, 0, 0, 0, 94, 0, 108, 0, 63, 62,
	0, 0, 416, 0, 0, 0, 0, 339, 339, 339,
	0, 0, 0, 0, 0, 432, 0, 0, 0, 0,
	106, 103, 106, 0, 0, 436, 103, 106, 339, 0,
	0, 0, 339, 339, 339, 79, 31, 0, 106, 106,
	106, 106, 0, 106, 106, 106, 106, 0, 106, 106,
	106, 0, 106, 0, 0, 106, 106, 106, 0, 0,
	0, 0, 106, 0, 0, 104, 0, 0, 0, 0,
	106, 77, 78, 0, 0, 0, 80, 81, 82, 106,
	106, 103, 0, 0, 0, 76, 85, 83, 84, 87,
	0, 0, 0, 0, 0, 0, 104, 0, 104, 0,
	0, 0, 0, 104, 0, 0, 0, 0, 0, 0,
	0, 106, 0, 106, 104, 104, 104, 104, 106, 104,
	104, 104, 104, 0, 104, 104, 104, 0, 104, 0,
	0, 104, 104, 104, 0, 0, 0, 0, 104, 0,
	0, 0, 106, 0, 0, 0, 104, 0, 0, 0,
	0, 0, 32, 0, 0, 104, 104, 0, 0, 0,
	0, 0, 0, 0, 0, 27, 101, 69, 102, 91,
	0, 97, 107, 0, 0, 0, 0, 0, 0, 0,
	0, 105, 0, 0, 106, 0, 0, 104, 0, 104,
	0, 0, 0, 0, 104, 71, 0, 106, 64, 65,
	0, 0, 95, 0, 0, 0, 106, 106, 96, 0,
	106, 0, 105, 0, 105, 0, 0, 0, 104, 105,
	94, 0, 108, 0, 63, 62, 0, 0, 0, 0,
	105, 105, 105, 105, 0, 105, 105, 105, 105, 0,
	105, 105, 105, 0, 105, 106, 106, 105, 105, 105,
	0, 0, 106, 0, 105, 0, 0, 0, 0, 106,
	104, 0, 105, 0, 0, 0, 0, 0, 0, 0,
	0, 105, 105, 104, 0, 0, 0, 0, 106, 0,
	0, 0, 104, 104, 0, 0, 104, 0, 0, 0,
	27, 101, 69, 102, 245, 0, 0, 107, 0, 0,
	0, 0, 0, 105, 0, 105, 0, 0, 0, 106,
	105, 0, 0, 0, 106, 0, 0, 0, 0, 0,
	71, 104, 104, 64, 65, 0, 0, 244, 104, 0,
	0, 0, 0, 0, 105, 104, 0, 27, 101, 69,
	102, 70, 0, 0, 0, 66, 0, 108, 0, 63,
	62, 0, 0, 0, 104, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 71, 0, 106,
	64, 65, 0, 0, 0, 0, 105, 0, 0, 0,
	0, 0, 231, 0, 0, 104, 0, 0, 0, 105,
	104, 0, 66, 0, 67, 0, 63, 62, 105, 105,
	0, 0, 105, 0, 0, 0, 0, 0, 0, 27,
	28, 69, 29, 70, 0, 0, 0, 41, 422, 342,
	421, 420, 343, 42, 43, 0, 60, 0, 51, 0,
	0, 344, 345, 0, 0, 0, 59, 105, 105, 71,
	61, 0, 64, 65, 105, 104, 0, 45, 46, 47,
	48, 105, 0, 0, 336, 337, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 0, 67, 0, 63, 62,
	105, 0, 0, 0, 0, 0, 0, 27, 28, 69,
	29, 70, 0, 0, 0, 41, 418, 342, 421, 420,
	343, 42, 43, 0, 60, 0, 51, 0, 0, 344,
	345, 105, 0, 0, 59, 0, 105, 71, 61, 0,
	64, 65, 0, 0, 0, 45, 46, 47, 48, 0,
	0, 0, 336, 337, 0, 0, 0, 0, 0, 0,
	0, 0, 66, 0, 67, 0, 63, 62, 0, 0,
	27, 28, 69, 29, 70, 0, 0, 0, 41, 424,
	342, 0, 0, 343, 42, 43, 0, 60, 0, 51,
	0, 105, 344, 345, 0, 0, 0, 59, 0, 0,
	71, 61, 0, 64, 65, 0, 0, 0, 45, 46,
	47, 48, 0, 0, 0, 336, 337, 0, 0, 0,
	0, 0, 0, 0, 0, 66, 0, 67, 0, 63,
	62, 27, 28, 69, 29, 70, 0, 0, 0, 41,
	335, 342, 0, 0, 343, 42, 43, 0, 60, 0,
	51, 0, 0, 344, 345, 0, 0, 0, 59, 0,
	0, 71, 61, 0, 64, 65, 0, 0, 0, 45,
	46, 47, 48, 0, 0, 0, 336, 337, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 0, 67, 0,
	63, 62, 27, 28, 69, 29, 70, 0, 0, 0,
	41, 330, 49, 275, 274, 50, 42, 43, 0, 60,
	0, 51, 0, 0, 0, 0, 0, 0, 0, 59,
	0, 0, 71, 61, 0, 64, 65, 0, 0, 0,
	45, 46, 47, 48, 0, 0, 0, 218, 219, 0,
	0, 0, 0, 0, 0, 0, 0, 66, 0, 67,
	0, 63, 62, 27, 28, 69, 29, 70, 0, 0,
	0, 41, 322, 49, 0, 0, 50, 42, 43, 0,
	60, 0, 51, 280, 0, 0, 0, 0, 0, 317,
	59, 0, 0, 71, 61, 0, 64, 65, 0, 0,
	0, 45, 46, 47, 48, 0, 0, 0, 310, 311,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 0,
	67, 0, 63, 62, 27, 28, 69, 29, 70, 0,
	0, 0, 41, 308, 49, 0, 0, 50, 42, 43,
	0, 60, 0, 51, 280, 0, 0, 0, 0, 0,
	317, 59, 0, 0, 71, 61, 0, 64, 65, 0,
	0, 0, 45, 46, 47, 48, 0, 0, 0, 310,
	311, 0, 0, 0, 0, 0, 0, 0, 0, 66,
	0, 67, 0, 63, 62, 27, 28, 69, 29, 70,
	0, 0, 0, 41, 272, 49, 275, 274, 50, 42,
	43, 0, 60, 0, 51, 0, 0, 0, 0, 0,
	0, 0, 59, 0, 0, 71, 61, 0, 64, 65,
	0, 0, 0, 45, 46, 47, 48, 0, 0, 0,
	218, 219, 0, 27, 28, 69, 29, 70, 0, 0,
	66, 41, 67, 49, 63, 62, 50, 42, 43, 0,
	60, 0, 51, 0, 0, 0, 0, 0, 0, 0,
	59, 0, 0, 71, 61, 0, 64, 65, 0, 0,
	0, 45, 46, 47, 48, 0, 0, 0, 6, 7,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 0,
	67, 0, 63, 62, 0, 8, 27, 28, 69, 29,
	70, 0, 0, 0, 41, 0, 342, 0, 0, 343,
	42, 43, 0, 60, 0, 51, 0, 0, 344, 345,
	0, 0, 0, 59, 0, 0, 71, 61, 0, 64,
	65, 0, 0, 0, 45, 46, 47, 48, 0, 0,
	0, 336, 337, 0, 0, 0, 0, 0, 0, 0,
	0, 66, 0, 67, 0, 63, 62, 27, 28, 69,
	29, 70, 0, 0, 0, 41, 404, 49, 0, 0,
	50, 42, 43, 0, 60, 0, 51, 0, 0, 0,
	0, 0, 0, 0, 59, 0, 0, 71, 61, 0,
	64, 65, 0, 0, 0, 45, 46, 47, 48, 0,
	0, 0, 218, 219, 0, 0, 0, 0, 0, 0,
	0, 0, 66, 0, 67, 0, 63, 62, 27, 28,
	69, 29, 70, 0, 0, 0, 41, 372, 49, 0,
	0, 50, 42, 43, 0, 60, 0, 51, 0, 0,
	0, 0, 0, 0, 0, 59, 0, 0, 71, 61,
	0, 64, 65, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 0, 218, 219, 0, 0, 0, 0, 0,
	0, 0, 0, 66, 0, 67, 0, 63, 62, 27,
	28, 69, 29, 70, 0, 0, 0, 41, 332, 49,
	0, 0, 50, 42, 43, 0, 60, 0, 51, 0,
	0, 0, 0, 0, 0, 0, 59, 0, 0, 71,
	61, 0, 64, 65, 0, 0, 0, 45, 46, 47,
	48, 0, 0, 0, 218, 219, 0, 27, 28, 69,
	29, 70, 0, 0, 66, 41, 67, 49, 63, 62,
	50, 42, 43, 0, 60, 0, 51, 0, 0, 0,
	0, 0, 0, 0, 59, 0, 0, 71, 61, 0,
	64, 65, 0, 0, 0, 45, 46, 47, 48, 0,
	0, 0, 218, 219, 0, 0, 0, 0, 0, 0,
	0, 0, 66, 0, 67, 307, 63, 62, 27, 28,
	69, 29, 70, 0, 0, 0, 41, 304, 49, 0,
	0, 50, 42, 43, 0, 60, 0, 51, 0, 0,
	0, 0, 0, 0, 0, 59, 0, 0, 71, 61,
	0, 64, 65, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 0, 218, 219, 0, 0, 0, 0, 0,
	0, 0, 0, 66, 0, 67, 0, 63, 62, 27,
	28, 69, 29, 70, 0, 0, 0, 41, 271, 49,
	0, 0, 50, 42, 43, 0, 60, 0, 51, 0,
	0, 0, 0, 0, 0, 0, 59, 0, 0, 71,
	61, 0, 64, 65, 0, 0, 0, 45, 46, 47,
	48, 0, 0, 0, 218, 219, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 0, 67, 0, 63, 62,
	27, 28, 69, 29, 70, 0, 0, 0, 41, 269,
	49, 0, 0, 50, 42, 43, 0, 60, 0, 51,
	0, 0, 0, 0, 0, 0, 0, 59, 0, 0,
	71, 61, 0, 64, 65, 0, 0, 0, 45, 46,
	47, 48, 0, 0, 0, 218, 219, 0, 27, 28,
	69, 29, 70, 0, 0, 66, 41, 67, 49, 63,
	62, 50, 42, 43, 0, 60, 0, 51, 0, 0,
	0, 0, 0, 0, 0, 59, 0, 0, 71, 61,
	0, 64, 65, 0, 0, 0, 45, 46, 47, 48,
	0, 0, 0, 218, 219, 0, 0, 0, 0, 0,
	0, 0, 0, 66, 0, 67, 259, 63, 62, 27,
	28, 69, 29, 70, 0, 0, 0, 41, 254, 49,
	0, 0, 50, 42, 43, 0, 60, 0, 51, 0,
	0, 0, 0, 0, 0, 0, 59, 0, 0, 71,
	61, 0, 64, 65, 0, 0, 0, 45, 46, 47,
	48, 0, 0, 0, 218, 219, 0, 27, 28, 69,
	29, 70, 0, 0, 66, 41, 67, 49, 63, 62,
	50, 42, 43, 0, 60, 0, 51, 0, 0, 0,
	0, 0, 0, 0, 59, 0, 0, 71, 61, 0,
	64, 65, 0, 0, 0, 45, 46, 47, 48, 0,
	0, 0, 218, 219, 0, 27, 28, 69, 29, 70,
	241, 0, 66, 41, 67, 49, 63, 62, 50, 42,
	43, 0, 60, 0, 51, 0, 0, 0, 0, 0,
	0, 0, 59, 0, 0, 71, 61, 0, 64, 65,
	0, 0, 0, 45, 46, 47, 48, 0, 0, 0,
	0, 240, 0, 27, 28, 69, 29, 70, 0, 0,
	66, 41, 67, 49, 63, 62, 50, 42, 43, 0,
	60, 0, 51, 0, 0, 0, 0, 0, 0, 0,
	59, 0, 0, 71, 61, 0, 64, 65, 0, 0,
	0, 45, 46, 47, 48, 27, 101, 69, 102, 70,
	0, 0, 107, 27, 236, 69, 102, 70, 66, 0,
	67, 0, 63, 62, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 71, 0, 0, 64, 65,
	0, 0, 0, 71, 0, 0, 64, 65, 27, 101,
	69, 102, 91, 0, 0, 107, 0, 0, 231, 0,
	66, 0, 108, 0, 63, 62, 0, 0, 66, 0,
	67, 0, 63, 62, 0, 0, 0, 0, 71, 0,
	0, 64, 65, 27, 101, 69, 102, 70, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 94, 0, 108, 0, 63, 62, 0,
	0, 0, 0, 71, 0, 79, 64, 65, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 86,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 0,
	67, 0, 63, 62, 88, 89, 0, 0, 0, 0,
	0, 77, 78, 0, 0, 0, 80, 81, 82, 0,
	0, 0, 0, 0, 0, 76, 85, 83, 84, 87,
}
var RubyPact = []int{

	-28, 1638, -1000, -1000, -1000, 6, -1000, -1000, -1000, 2521,
	-1000, -1000, -1000, 41, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 910, 172,
	74, 7, 3, -1000, -1000, -1000, -1000, -1000, 191, -22,
	-1000, 270, 225, 225, 63, 2368, 2368, 2368, 2368, 2368,
	2368, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 2488,
	2368, 2488, 10, 259, -1000, -1000, 2488, -1000, -17, 194,
	-1000, 29, -1000, -1000, -1000, 2368, 311, 2488, 2488, 2488,
	2488, 2368, 2488, 2488, 2488, 2488, 2368, 2488, 2488, 2488,
	2368, 2488, -1000, 97, 2488, 2488, 2488, 310, 157, 185,
	-1000, 2453, 168, -1000, -1000, -1000, 77, -21, -21, 2488,
	2368, 2368, 2368, 2368, 2368, 2368, 2368, 2368, 2488, 2488,
	2368, 2368, 309, 52, 119, 8, -1000, -1000, 2368, 307,
	54, 54, 54, 54, 54, 20, 2272, 90, 185, 85,
	185, -1000, -1000, 255, -1000, -1000, 33, 32, 781, -1000,
	2418, 232, 2488, 2320, -1000, -21, 54, 1035, 185, 185,
	185, 185, 54, 185, 185, 185, 185, 54, 182, 185,
	185, 54, 300, 781, -1000, 482, 103, -1000, 103, -1000,
	-1000, 2410, 2224, -1000, 306, -1000, 2163, 293, 54, 54,
	54, 54, 54, 54, 54, 54, 185, 185, 54, 54,
	-1000, -1000, 302, 152, 127, -1000, 22, 292, 282, -1000,
	2115, 225, 2054, 54, -1000, 1590, -1000, -1000, -1000, -1000,
	2521, 54, 171, 2488, -1000, 13, -1000, -1000, -1000, -1000,
	-1000, -1000, 123, 118, -19, 189, 684, -1000, 291, 54,
	-1000, -1000, 97, 29, 2368, 2488, 2488, 29, -1000, 616,
	5, -24, 185, -1000, -1000, 1993, 17, -1000, 1932, -1000,
	-1000, 1529, 119, 127, 208, 2368, -1000, -1000, 1468, -1000,
	-1000, -1000, -1000, 280, 2368, -1000, 1407, 1884, -1000, -1000,
	199, 185, 1346, 151, 2488, 1082, -23, -1000, -25, -1000,
	2368, 2488, -1000, -1000, 54, 145, 185, -1000, 140, 185,
	-1000, 278, 2368, 249, -1000, -1000, 245, -1000, -1000, 155,
	-1000, -1000, 2521, 54, -1000, -1000, 230, 2488, -1000, -1000,
	-1000, 54, -1000, 117, 1823, -1000, 2368, -1000, 54, 2272,
	-1000, 273, -1000, 2272, 316, -1000, -1000, -1000, 2521, 54,
	-1000, -1000, 2368, 2368, 220, 121, -1000, -1000, 2488, 90,
	781, -1000, -1000, 1082, -1000, 48, 2521, 54, 185, -1000,
	238, -1000, 54, -1000, -1000, -1000, -1000, 2368, 2368, 90,
	547, -1000, -1000, 54, 2272, 2272, -1000, 2272, 222, 27,
	15, 2368, 2368, 2368, 2368, 1762, 90, 2272, 158, -15,
	114, 54, 54, -1000, 113, 2272, -1000, -1000, -1000, -1000,
	54, 54, 54, 54, -1000, 2272, -19, 2368, 2488, -1000,
	-1000, 2272, 1222, 1154, 1285, -19, 180, 100, -1000, 265,
	2368, -1000, -1000, 256, -1000, -1000, -1000, -19, -1000, -1000,
	2368, -1000, 54, 1701, -1000, -19, 54, 1701, 1701, 1701,
}
var RubyPgo = []int{

	0, 344, 384, 383, 10, 382, 381, 380, 902, 379,
	4, 17, 373, 370, 369, 0, 786, 710, 368, 367,
	366, 20, 365, 7, 422, 364, 363, 9, 362, 361,
	360, 358, 356, 355, 351, 350, 2, 348, 347, 345,
	343, 342, 341, 340, 339, 338, 337, 336, 24, 335,
	6, 14, 15, 3, 333, 1, 332, 46, 331, 11,
	5, 329, 8, 327, 323, 13, 12, 322, 321, 19,
}
var RubyR1 = []int{

	0, 54, 54, 54, 54, 54, 54, 54, 54, 54,
	54, 68, 68, 69, 69, 48, 48, 48, 48, 48,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 21, 21, 21, 21, 21, 21,
	21, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 51,
	51, 51, 51, 62, 62, 59, 59, 59, 59, 59,
	65, 65, 65, 65, 65, 64, 64, 64, 18, 18,
	18, 18, 18, 18, 28, 28, 28, 28, 60, 60,
	60, 60, 60, 60, 55, 55, 23, 23, 23, 23,
	66, 66, 66, 22, 22, 25, 27, 27, 67, 67,
	13, 13, 13, 13, 13, 13, 13, 26, 26, 26,
	26, 26, 26, 8, 8, 24, 24, 16, 16, 37,
	37, 38, 39, 40, 41, 42, 43, 44, 45, 46,
	47, 2, 5, 6, 6, 3, 3, 56, 56, 56,
	56, 63, 63, 63, 4, 4, 4, 4, 52, 61,
	61, 61, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 53, 53, 53, 53, 49, 49, 49, 7,
	14, 10, 10, 10, 58, 58, 50, 50, 19, 20,
	11, 33, 57, 57, 57, 57, 57, 57, 57, 34,
	34, 34, 34, 34, 34, 34, 35, 35, 35, 35,
	35, 36, 36, 36, 36, 32, 31, 9, 30, 30,
	29, 29,
}
var RubyR2 = []int{

	0, 0, 1, 1, 1, 3, 3, 3, 2, 2,
	2, 0, 1, 0, 2, 0, 2, 2, 2, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 4,
	5, 1, 4, 4, 2, 3, 4, 4, 5, 3,
	5, 2, 3, 3, 3, 3, 3, 4, 6, 3,
	7, 1, 5, 1, 3, 0, 1, 1, 4, 4,
	1, 1, 4, 4, 5, 1, 3, 3, 5, 6,
	7, 8, 5, 6, 0, 1, 3, 3, 0, 2,
	2, 2, 2, 2, 1, 3, 1, 2, 3, 2,
	0, 1, 3, 4, 6, 4, 1, 3, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 2, 2, 2, 3, 3, 3,
	3, 2, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 1, 1, 3, 3, 5, 5, 0, 1, 3,
	7, 3, 7, 8, 3, 4, 4, 3, 3, 0,
	1, 3, 4, 5, 3, 3, 3, 3, 3, 5,
	6, 5, 4, 3, 3, 2, 0, 2, 2, 3,
	4, 2, 3, 5, 0, 2, 1, 2, 2, 2,
	5, 5, 0, 2, 2, 2, 2, 2, 2, 0,
	1, 3, 3, 1, 3, 3, 5, 6, 5, 6,
	5, 4, 3, 3, 2, 3, 3, 2, 5, 7,
	4, 5,
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
	-69, 53, 9, -49, -4, 62, -1, 6, -15, -15,
	-15, -15, -1, -15, -15, -15, -15, -1, -15, -15,
	-15, -1, -59, -15, 11, -15, -15, -11, -15, -11,
	6, 11, -48, -52, 55, -52, -48, -59, -1, -1,
	-1, -1, -1, -1, -1, -1, -15, -15, -1, -1,
	6, -55, 54, -66, 9, -23, 6, 47, 57, -55,
	-48, 40, -48, -1, 6, -48, 50, 51, 50, 51,
	-15, -1, -58, 11, 50, -69, 6, 8, 61, 11,
	61, 50, -56, -63, -69, -15, 6, 8, -59, -1,
	51, 10, -65, -51, 42, 9, 52, 10, 11, -69,
	61, -69, -15, -4, 14, -48, -61, 6, -48, 63,
	10, -60, 6, 11, -66, 42, 6, 6, -60, 14,
	-27, 14, 14, -53, 17, 16, -48, -48, 14, -10,
	25, -15, -57, -29, 37, -69, -69, 11, -69, 11,
	4, 52, 10, -4, -1, -59, -15, -4, -69, -15,
	-4, 57, 42, 57, 14, 55, 11, 63, 14, -50,
	50, 51, -15, -1, -28, -10, -20, 31, -55, -23,
	10, -1, 14, -50, -48, 14, 17, 16, -1, -48,
	14, -53, 14, -48, 8, 14, 50, 51, -15, -1,
	-35, -34, 15, 18, 27, 28, 14, 16, 37, -62,
	-15, -21, 63, -69, 63, -69, -15, -1, -15, 10,
	57, 6, -1, 6, 6, 14, -10, 15, 18, -62,
	-60, 14, 14, -1, -48, -48, 14, -48, 4, -1,
	-1, 15, 18, 15, 18, -48, -62, -48, -15, 6,
	6, -1, -1, 14, -50, -48, 6, 50, 50, 51,
	-1, -1, -1, -1, 14, -48, -69, 4, 52, 10,
	14, -48, -57, -57, -57, -69, -1, -15, 14, -36,
	17, 16, 14, -36, 14, -68, 11, -69, 11, 14,
	17, 16, -1, -57, 14, -69, -1, -57, -57, -57,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 20, 21, -2,
	23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
	33, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 15, 34, 35, 36, 37, 38, 39, 40, 0,
	0, 0, 0, 0, 161, 162, 85, 13, 0, 61,
	196, 0, 5, 6, 7, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, -2, 64, 71, 85, 0, 0, 0, 81, 90,
	91, 21, -2, 23, 24, 25, 31, 15, -2, 85,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 120, 120, 15, -2, 15, 0, 0,
	151, 152, 153, 154, 15, 0, 204, 208, 83, 0,
	13, 145, 146, 0, 143, 144, 0, 0, 83, 87,
	13, 0, 85, 0, 237, 15, 184, 65, 72, 74,
	76, 155, 156, 157, 158, 159, 160, 186, 0, 235,
	236, 188, 0, 86, 13, 83, 130, 131, 137, 138,
	149, 13, 0, 15, 179, 15, 0, 0, 132, 139,
	133, 140, 134, 141, 135, 142, 73, 75, 185, 187,
	69, 108, 0, 114, 120, 121, 116, 0, 0, 108,
	0, 0, 0, 136, 150, 0, 15, 15, 16, 17,
	18, 19, 0, 0, 212, 0, 147, 148, 163, 13,
	164, 14, 13, 13, 168, 0, 21, -2, 0, 197,
	198, 199, 66, 67, 0, -2, 0, 59, 13, 0,
	77, 0, 96, 97, 174, 0, 0, 180, 0, 177,
	63, 0, 120, 0, 0, 0, 117, 119, 0, 123,
	15, 125, 182, 0, 0, 15, 0, 0, 200, 205,
	15, 84, 0, 0, 0, 0, 0, 13, 0, 13,
	0, 0, 62, 68, 70, 0, 210, 60, 0, 92,
	93, 0, 0, 0, 175, 178, 0, 176, 98, 0,
	109, 110, 41, 112, 113, 206, 105, 0, 108, 122,
	115, 118, 102, 0, 0, 183, 0, 15, 15, 195,
	189, 0, 191, 201, 15, 211, 213, 214, 41, 216,
	217, 218, 0, 0, 220, 223, 238, 15, 0, 15,
	88, 89, 165, 0, 166, 0, 41, 169, 171, 79,
	0, 94, 78, 82, 181, 99, 207, 0, 0, 209,
	0, 103, 124, 15, 193, 194, 190, 202, 0, 15,
	0, 0, 0, 0, 0, 0, 15, 13, 0, 0,
	0, 106, 107, 100, 0, 192, 15, 212, -2, -2,
	221, 222, 224, 225, 239, 13, 240, 0, 0, 80,
	101, 203, 0, 0, 0, 241, 11, 13, 226, 0,
	0, 212, 228, 0, 230, 170, 12, 172, 13, 227,
	0, 212, 212, 219, 229, 173, 212, 219, 219, 219,
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
		//line parser.y:197
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:199
		{
		}
	case 3:
		//line parser.y:201
		{
		}
	case 4:
		//line parser.y:203
		{
		}
	case 5:
		//line parser.y:205
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 6:
		//line parser.y:207
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 7:
		//line parser.y:209
		{
			Statements = append(Statements, RubyS[Rubypt-1].genericValue)
		}
	case 8:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 9:
		RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
	case 10:
		//line parser.y:215
		{
		}
	case 11:
		//line parser.y:217
		{
		}
	case 12:
		//line parser.y:218
		{
		}
	case 13:
		//line parser.y:220
		{
		}
	case 14:
		//line parser.y:221
		{
		}
	case 15:
		//line parser.y:224
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 16:
		//line parser.y:226
		{
		}
	case 17:
		//line parser.y:228
		{
		}
	case 18:
		//line parser.y:230
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 19:
		//line parser.y:232
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
		//line parser.y:242
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 60:
		//line parser.y:249
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: append(RubyS[Rubypt-2].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 61:
		//line parser.y:256
		{
			RubyVAL.genericValue = ast.CallExpression{Func: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 62:
		//line parser.y:260
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 63:
		//line parser.y:267
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 64:
		//line parser.y:274
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 65:
		//line parser.y:281
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 66:
		//line parser.y:288
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 67:
		//line parser.y:296
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-3].genericValue,
				Func:   RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 68:
		//line parser.y:304
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   append(RubyS[Rubypt-1].genericSlice, RubyS[Rubypt-0].genericValue),
			}
		}
	case 69:
		//line parser.y:312
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
				Args:   []ast.Node{},
			}
		}
	case 70:
		//line parser.y:320
		{
			methodName := RubyS[Rubypt-2].genericValue.(ast.BareReference).Name + "="
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: methodName},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 71:
		//line parser.y:331
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 72:
		//line parser.y:338
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 73:
		//line parser.y:346
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "<"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 74:
		//line parser.y:354
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 75:
		//line parser.y:362
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: ">"},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 76:
		//line parser.y:370
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-1].operator},
				Target: RubyS[Rubypt-2].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 77:
		//line parser.y:380
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]"},
				Target: RubyS[Rubypt-3].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-1].genericValue},
			}
		}
	case 78:
		//line parser.y:390
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: "[]="},
				Target: RubyS[Rubypt-5].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-3].genericValue, RubyS[Rubypt-0].genericValue},
			}
		}
	case 79:
		//line parser.y:399
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 80:
		//line parser.y:401
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-5].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-1].genericValue})
		}
	case 81:
		//line parser.y:403
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 82:
		//line parser.y:405
		{
			RubyVAL.genericSlice = append(RubyS[Rubypt-4].genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 83:
		//line parser.y:408
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 84:
		//line parser.y:410
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 85:
		//line parser.y:412
		{
			RubyVAL.genericSlice = ast.Nodes{}
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
		//line parser.y:418
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 89:
		//line parser.y:420
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
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 92:
		//line parser.y:428
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 93:
		//line parser.y:430
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-1].genericValue)
		}
	case 94:
		//line parser.y:432
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.ProcArg{Value: RubyS[Rubypt-0].genericValue})
		}
	case 95:
		//line parser.y:435
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 96:
		//line parser.y:437
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 97:
		//line parser.y:439
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 98:
		//line parser.y:445
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 99:
		//line parser.y:453
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 100:
		//line parser.y:462
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 101:
		//line parser.y:471
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 102:
		//line parser.y:481
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 103:
		//line parser.y:489
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 104:
		//line parser.y:499
		{
		}
	case 105:
		//line parser.y:501
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 106:
		//line parser.y:503
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{RubyS[Rubypt-2].genericValue}}
		}
	case 107:
		//line parser.y:505
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{ast.Break{}},
			}
		}
	case 108:
		//line parser.y:513
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 109:
		//line parser.y:515
		{
		}
	case 110:
		//line parser.y:517
		{
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
		//line parser.y:528
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 115:
		//line parser.y:530
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 116:
		//line parser.y:533
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 117:
		//line parser.y:535
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 118:
		//line parser.y:537
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 119:
		//line parser.y:539
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 120:
		//line parser.y:541
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 121:
		//line parser.y:543
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 122:
		//line parser.y:547
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 123:
		//line parser.y:552
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 124:
		//line parser.y:559
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 125:
		//line parser.y:569
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 126:
		//line parser.y:578
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 127:
		//line parser.y:584
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 128:
		//line parser.y:592
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 129:
		//line parser.y:596
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
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
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 132:
		//line parser.y:612
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 133:
		//line parser.y:619
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 134:
		//line parser.y:626
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 135:
		//line parser.y:633
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 136:
		//line parser.y:640
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
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
			RubyVAL.genericValue = ast.ConditionalAssignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 139:
		//line parser.y:659
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 140:
		//line parser.y:666
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 141:
		//line parser.y:673
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 142:
		//line parser.y:680
		{
			RubyVAL.genericValue = ast.ConditionalAssignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 143:
		//line parser.y:688
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 144:
		//line parser.y:690
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 145:
		//line parser.y:693
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 146:
		//line parser.y:695
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 147:
		//line parser.y:698
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 148:
		//line parser.y:700
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 149:
		//line parser.y:703
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 150:
		//line parser.y:705
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 151:
		//line parser.y:707
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 152:
		//line parser.y:708
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 153:
		//line parser.y:709
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 154:
		//line parser.y:710
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 155:
		//line parser.y:713
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 156:
		//line parser.y:722
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 157:
		//line parser.y:731
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 158:
		//line parser.y:740
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 159:
		//line parser.y:749
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 160:
		//line parser.y:758
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 161:
		//line parser.y:766
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 162:
		//line parser.y:767
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 163:
		//line parser.y:769
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 164:
		//line parser.y:771
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 165:
		//line parser.y:774
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 166:
		//line parser.y:782
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 167:
		//line parser.y:790
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 168:
		//line parser.y:792
		{
		}
	case 169:
		//line parser.y:794
		{
			if RubyS[Rubypt-1].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-2].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 170:
		//line parser.y:801
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 171:
		//line parser.y:809
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 172:
		//line parser.y:816
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 173:
		//line parser.y:823
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 174:
		//line parser.y:831
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 175:
		//line parser.y:833
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 176:
		//line parser.y:840
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 177:
		//line parser.y:847
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 178:
		//line parser.y:850
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 179:
		//line parser.y:852
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 180:
		//line parser.y:854
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 181:
		//line parser.y:856
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 182:
		//line parser.y:859
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 183:
		//line parser.y:866
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 184:
		//line parser.y:874
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 185:
		//line parser.y:881
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 186:
		//line parser.y:888
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 187:
		//line parser.y:895
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 188:
		//line parser.y:902
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 189:
		//line parser.y:909
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 190:
		//line parser.y:916
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 191:
		//line parser.y:924
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 192:
		//line parser.y:933
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 193:
		//line parser.y:940
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 194:
		//line parser.y:947
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 195:
		//line parser.y:954
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 196:
		//line parser.y:961
		{
		}
	case 197:
		//line parser.y:962
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 198:
		//line parser.y:963
		{
		}
	case 199:
		//line parser.y:966
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 200:
		//line parser.y:969
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 201:
		//line parser.y:977
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 202:
		//line parser.y:979
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 203:
		//line parser.y:988
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
	case 204:
		//line parser.y:1003
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 205:
		//line parser.y:1005
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 206:
		//line parser.y:1008
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 207:
		//line parser.y:1010
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 208:
		//line parser.y:1013
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 209:
		//line parser.y:1022
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 210:
		//line parser.y:1031
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 211:
		//line parser.y:1040
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 212:
		//line parser.y:1043
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 213:
		//line parser.y:1045
		{
		}
	case 214:
		//line parser.y:1047
		{
		}
	case 215:
		//line parser.y:1049
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 216:
		//line parser.y:1051
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 217:
		//line parser.y:1053
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 218:
		//line parser.y:1055
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 219:
		//line parser.y:1057
		{
		}
	case 220:
		//line parser.y:1059
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 221:
		//line parser.y:1061
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 222:
		//line parser.y:1063
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 223:
		//line parser.y:1065
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 224:
		//line parser.y:1067
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 225:
		//line parser.y:1069
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 226:
		//line parser.y:1072
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 227:
		//line parser.y:1079
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 228:
		//line parser.y:1087
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 229:
		//line parser.y:1094
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 230:
		//line parser.y:1102
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 231:
		//line parser.y:1110
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 232:
		//line parser.y:1117
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 233:
		//line parser.y:1124
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 234:
		//line parser.y:1131
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 235:
		//line parser.y:1139
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 236:
		//line parser.y:1142
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 237:
		//line parser.y:1144
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 238:
		//line parser.y:1147
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 239:
		//line parser.y:1149
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 240:
		//line parser.y:1152
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 241:
		//line parser.y:1154
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	}
	goto Rubystack /* stack new state and value */
}
