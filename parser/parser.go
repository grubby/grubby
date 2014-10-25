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

//line parser.y:1114

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 28,
	52, 125,
	-2, 20,
	-1, 90,
	10, 82,
	11, 82,
	-2, 187,
	-1, 100,
	52, 125,
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
	52, 125,
	-2, 123,
	-1, 224,
	52, 126,
	-2, 124,
	-1, 232,
	10, 82,
	11, 82,
	-2, 187,
	-1, 385,
	27, 203,
	28, 203,
	-2, 13,
	-1, 386,
	27, 203,
	28, 203,
	-2, 13,
}

const RubyNprod = 233
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 2592

var RubyAct = []int{

	9, 295, 260, 140, 269, 301, 406, 248, 20, 189,
	193, 36, 92, 91, 141, 191, 57, 177, 2, 3,
	339, 74, 219, 135, 219, 136, 105, 217, 97, 289,
	292, 176, 98, 116, 145, 4, 395, 340, 271, 219,
	288, 252, 111, 211, 112, 113, 89, 110, 199, 109,
	219, 74, 74, 119, 121, 72, 71, 194, 219, 132,
	192, 134, 131, 74, 369, 194, 142, 370, 192, 194,
	396, 338, 73, 291, 143, 149, 218, 152, 153, 154,
	155, 137, 157, 158, 159, 160, 384, 162, 163, 164,
	78, 167, 216, 78, 169, 170, 211, 212, 195, 143,
	97, 123, 143, 148, 190, 166, 195, 196, 167, 367,
	195, 171, 368, 184, 185, 196, 143, 107, 332, 196,
	333, 78, 179, 175, 168, 74, 76, 77, 197, 76,
	77, 208, 122, 26, 99, 68, 100, 69, 74, 75,
	105, 334, 75, 86, 78, 222, 276, 167, 275, 144,
	108, 415, 97, 375, 250, 143, 173, 76, 77, 385,
	386, 225, 107, 70, 229, 230, 63, 64, 233, 214,
	75, 215, 204, 205, 239, 208, 146, 219, 78, 208,
	76, 77, 76, 77, 287, 353, 413, 65, 354, 106,
	74, 62, 61, 75, 320, 75, 219, 114, 397, 208,
	115, 208, 357, 120, 208, 255, 240, 351, 251, 267,
	224, 257, 268, 267, 76, 77, 266, 364, 267, 79,
	80, 81, 112, 113, 97, 265, 383, 75, 84, 82,
	83, 86, 377, 167, 282, 67, 267, 285, 350, 345,
	235, 143, 421, 208, 418, 417, 208, 281, 416, 298,
	418, 417, 362, 349, 313, 312, 298, 309, 347, 304,
	394, 305, 96, 279, 208, 208, 317, 283, 254, 286,
	324, 74, 132, 336, 138, 335, 139, 342, 344, 253,
	311, 337, 313, 312, 213, 336, 306, 250, 278, 217,
	247, 217, 249, 337, 209, 244, 5, 234, 235, 202,
	118, 352, 117, 188, 132, 172, 151, 355, 221, 243,
	210, 208, 356, 220, 1, 352, 208, 147, 236, 56,
	208, 55, 54, 238, 53, 52, 51, 17, 16, 15,
	14, 43, 326, 327, 22, 132, 23, 24, 372, 124,
	125, 126, 127, 128, 129, 25, 270, 300, 12, 11,
	302, 21, 10, 19, 133, 13, 18, 298, 381, 39,
	38, 208, 208, 34, 208, 33, 35, 272, 32, 150,
	273, 274, 208, 0, 208, 156, 0, 0, 0, 0,
	161, 0, 0, 208, 165, 284, 0, 352, 29, 399,
	400, 401, 0, 208, 0, 0, 404, 0, 0, 208,
	324, 324, 324, 180, 181, 182, 183, 410, 0, 186,
	187, 0, 0, 420, 0, 0, 101, 201, 0, 0,
	0, 324, 0, 425, 426, 341, 324, 324, 324, 427,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 226, 0, 0, 0, 0, 101, 0, 101,
	0, 0, 0, 0, 101, 0, 0, 26, 99, 68,
	100, 90, 0, 95, 105, 101, 101, 101, 101, 0,
	101, 101, 101, 101, 0, 101, 101, 101, 0, 101,
	0, 0, 101, 101, 0, 0, 0, 70, 101, 0,
	63, 64, 0, 376, 94, 0, 101, 0, 0, 0,
	0, 101, 101, 0, 0, 0, 0, 0, 0, 0,
	0, 93, 0, 106, 0, 62, 61, 0, 0, 0,
	0, 0, 0, 393, 0, 0, 280, 0, 0, 0,
	0, 0, 0, 101, 0, 101, 0, 0, 0, 0,
	101, 0, 402, 299, 0, 0, 0, 307, 0, 0,
	299, 0, 0, 412, 414, 0, 314, 26, 99, 68,
	100, 69, 101, 422, 325, 423, 0, 0, 0, 0,
	0, 343, 26, 99, 68, 100, 90, 0, 0, 105,
	0, 0, 0, 348, 0, 0, 0, 70, 0, 0,
	63, 64, 0, 0, 0, 0, 0, 0, 0, 0,
	101, 219, 70, 78, 0, 63, 64, 359, 346, 0,
	0, 65, 101, 66, 0, 62, 61, 0, 277, 0,
	0, 101, 101, 365, 366, 101, 93, 0, 106, 0,
	62, 61, 0, 0, 374, 0, 0, 0, 0, 76,
	77, 0, 0, 0, 79, 80, 81, 37, 378, 379,
	0, 299, 75, 84, 82, 83, 86, 0, 0, 237,
	101, 101, 387, 388, 389, 390, 101, 0, 0, 0,
	0, 0, 0, 101, 0, 104, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 403,
	0, 0, 101, 0, 325, 325, 325, 0, 0, 0,
	0, 0, 419, 0, 0, 0, 104, 0, 104, 0,
	0, 0, 424, 104, 0, 325, 0, 0, 0, 0,
	325, 325, 325, 101, 104, 104, 104, 104, 0, 104,
	104, 104, 104, 0, 104, 104, 104, 0, 104, 0,
	0, 104, 104, 0, 0, 0, 0, 104, 0, 0,
	30, 0, 0, 0, 0, 104, 0, 0, 0, 0,
	104, 104, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 102, 0,
	0, 0, 0, 0, 101, 0, 0, 0, 0, 0,
	0, 0, 104, 0, 104, 0, 0, 0, 0, 104,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 102,
	0, 102, 0, 0, 0, 0, 102, 0, 0, 0,
	0, 104, 0, 0, 0, 0, 0, 102, 102, 102,
	102, 0, 102, 102, 102, 102, 0, 102, 102, 102,
	0, 102, 0, 0, 102, 102, 0, 0, 0, 0,
	102, 0, 0, 26, 99, 68, 100, 232, 102, 104,
	105, 0, 0, 102, 102, 0, 0, 0, 0, 0,
	0, 104, 0, 0, 0, 0, 0, 0, 0, 0,
	104, 104, 178, 70, 104, 78, 63, 64, 0, 0,
	231, 0, 0, 0, 0, 102, 0, 102, 0, 85,
	0, 0, 102, 0, 0, 0, 0, 65, 0, 106,
	0, 62, 61, 0, 87, 88, 0, 0, 0, 104,
	104, 76, 77, 0, 102, 104, 79, 80, 81, 0,
	0, 0, 104, 130, 75, 84, 82, 83, 86, 26,
	27, 68, 28, 69, 0, 0, 0, 40, 0, 48,
	0, 104, 49, 41, 42, 0, 59, 0, 50, 0,
	0, 0, 102, 0, 0, 0, 58, 0, 31, 70,
	60, 0, 63, 64, 102, 0, 0, 44, 45, 46,
	47, 0, 104, 102, 102, 0, 0, 102, 174, 0,
	0, 0, 0, 65, 0, 66, 103, 62, 61, 0,
	0, 0, 198, 0, 200, 0, 0, 0, 0, 0,
	0, 203, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 102, 102, 0, 0, 0, 103, 102, 103,
	0, 0, 0, 0, 103, 102, 26, 99, 68, 100,
	69, 0, 0, 104, 0, 103, 103, 103, 103, 0,
	103, 103, 103, 103, 102, 103, 103, 103, 242, 103,
	245, 0, 103, 103, 0, 0, 70, 0, 103, 63,
	64, 26, 99, 68, 100, 69, 103, 0, 105, 0,
	219, 103, 103, 0, 0, 102, 0, 263, 264, 0,
	65, 0, 66, 0, 62, 61, 0, 0, 0, 0,
	0, 70, 0, 0, 63, 64, 0, 0, 0, 0,
	0, 0, 0, 103, 0, 103, 0, 0, 0, 0,
	103, 0, 0, 0, 0, 65, 0, 106, 0, 62,
	61, 26, 223, 68, 100, 69, 0, 0, 0, 0,
	310, 0, 103, 0, 0, 315, 102, 0, 0, 0,
	319, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 70, 0, 0, 63, 64, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 219, 0, 0, 0, 0,
	103, 0, 0, 0, 0, 65, 0, 66, 0, 62,
	61, 0, 103, 0, 0, 0, 360, 361, 0, 0,
	0, 103, 103, 363, 0, 103, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 371, 0, 373, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	103, 103, 382, 0, 0, 0, 103, 0, 203, 0,
	0, 0, 0, 103, 0, 392, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 398, 0, 263, 264,
	0, 0, 103, 0, 26, 27, 68, 28, 69, 0,
	0, 0, 40, 409, 328, 408, 407, 329, 41, 42,
	0, 59, 0, 50, 0, 0, 330, 331, 0, 0,
	0, 58, 0, 103, 70, 60, 0, 63, 64, 0,
	0, 0, 44, 45, 46, 47, 0, 0, 322, 323,
	0, 0, 0, 0, 0, 0, 0, 0, 65, 0,
	66, 0, 62, 61, 0, 26, 27, 68, 28, 69,
	0, 0, 0, 40, 405, 328, 408, 407, 329, 41,
	42, 0, 59, 0, 50, 0, 0, 330, 331, 0,
	0, 0, 58, 0, 103, 70, 60, 0, 63, 64,
	0, 0, 0, 44, 45, 46, 47, 0, 0, 322,
	323, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	0, 66, 0, 62, 61, 26, 27, 68, 28, 69,
	0, 0, 0, 40, 411, 328, 0, 0, 329, 41,
	42, 0, 59, 0, 50, 0, 0, 330, 331, 0,
	0, 0, 58, 0, 0, 70, 60, 0, 63, 64,
	0, 0, 0, 44, 45, 46, 47, 0, 0, 322,
	323, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	0, 66, 0, 62, 61, 26, 27, 68, 28, 69,
	0, 0, 0, 40, 380, 48, 0, 0, 49, 41,
	42, 0, 59, 0, 50, 267, 0, 0, 0, 0,
	0, 303, 58, 0, 0, 70, 60, 0, 63, 64,
	0, 0, 0, 44, 45, 46, 47, 0, 0, 296,
	297, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	0, 66, 0, 62, 61, 26, 27, 68, 28, 69,
	0, 0, 0, 40, 321, 328, 0, 0, 329, 41,
	42, 0, 59, 0, 50, 0, 0, 330, 331, 0,
	0, 0, 58, 0, 0, 70, 60, 0, 63, 64,
	0, 0, 0, 44, 45, 46, 47, 0, 0, 322,
	323, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	0, 66, 0, 62, 61, 26, 27, 68, 28, 69,
	0, 0, 0, 40, 316, 48, 262, 261, 49, 41,
	42, 0, 59, 0, 50, 0, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 70, 60, 0, 63, 64,
	0, 0, 0, 44, 45, 46, 47, 0, 0, 206,
	207, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	0, 66, 0, 62, 61, 26, 27, 68, 28, 69,
	0, 0, 0, 40, 308, 48, 0, 0, 49, 41,
	42, 0, 59, 0, 50, 267, 0, 0, 0, 0,
	0, 303, 58, 0, 0, 70, 60, 0, 63, 64,
	0, 0, 0, 44, 45, 46, 47, 0, 0, 296,
	297, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	0, 66, 0, 62, 61, 26, 27, 68, 28, 69,
	0, 0, 0, 40, 294, 48, 0, 0, 49, 41,
	42, 0, 59, 0, 50, 267, 0, 0, 0, 0,
	0, 303, 58, 0, 0, 70, 60, 0, 63, 64,
	0, 0, 0, 44, 45, 46, 47, 0, 0, 296,
	297, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	0, 66, 0, 62, 61, 26, 27, 68, 28, 69,
	0, 0, 0, 40, 259, 48, 262, 261, 49, 41,
	42, 0, 59, 0, 50, 0, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 70, 60, 0, 63, 64,
	0, 0, 0, 44, 45, 46, 47, 0, 0, 206,
	207, 0, 26, 27, 68, 28, 69, 0, 0, 65,
	40, 66, 48, 62, 61, 49, 41, 42, 0, 59,
	0, 50, 0, 0, 0, 0, 0, 0, 0, 58,
	0, 0, 70, 60, 0, 63, 64, 0, 0, 0,
	44, 45, 46, 47, 0, 0, 6, 7, 0, 0,
	0, 0, 0, 0, 0, 0, 65, 0, 66, 0,
	62, 61, 0, 8, 26, 27, 68, 28, 69, 0,
	0, 0, 40, 0, 328, 0, 0, 329, 41, 42,
	0, 59, 0, 50, 0, 0, 330, 331, 0, 0,
	0, 58, 0, 0, 70, 60, 0, 63, 64, 0,
	0, 0, 44, 45, 46, 47, 0, 0, 322, 323,
	0, 0, 0, 0, 0, 0, 0, 0, 65, 0,
	66, 0, 62, 61, 26, 27, 68, 28, 69, 0,
	0, 0, 40, 391, 48, 0, 0, 49, 41, 42,
	0, 59, 0, 50, 0, 0, 0, 0, 0, 0,
	0, 58, 0, 0, 70, 60, 0, 63, 64, 0,
	0, 0, 44, 45, 46, 47, 0, 0, 206, 207,
	0, 0, 0, 0, 0, 0, 0, 0, 65, 0,
	66, 0, 62, 61, 26, 27, 68, 28, 69, 0,
	0, 0, 40, 358, 48, 0, 0, 49, 41, 42,
	0, 59, 0, 50, 0, 0, 0, 0, 0, 0,
	0, 58, 0, 0, 70, 60, 0, 63, 64, 0,
	0, 0, 44, 45, 46, 47, 0, 0, 206, 207,
	0, 0, 0, 0, 0, 0, 0, 0, 65, 0,
	66, 0, 62, 61, 26, 27, 68, 28, 69, 0,
	0, 0, 40, 318, 48, 0, 0, 49, 41, 42,
	0, 59, 0, 50, 0, 0, 0, 0, 0, 0,
	0, 58, 0, 0, 70, 60, 0, 63, 64, 0,
	0, 0, 44, 45, 46, 47, 0, 0, 206, 207,
	0, 26, 27, 68, 28, 69, 0, 0, 65, 40,
	66, 48, 62, 61, 49, 41, 42, 0, 59, 0,
	50, 0, 0, 0, 0, 0, 0, 0, 58, 0,
	0, 70, 60, 0, 63, 64, 0, 0, 0, 44,
	45, 46, 47, 0, 0, 206, 207, 0, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 66, 293, 62,
	61, 26, 27, 68, 28, 69, 0, 0, 0, 40,
	290, 48, 0, 0, 49, 41, 42, 0, 59, 0,
	50, 0, 0, 0, 0, 0, 0, 0, 58, 0,
	0, 70, 60, 0, 63, 64, 0, 0, 0, 44,
	45, 46, 47, 0, 0, 206, 207, 0, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 66, 0, 62,
	61, 26, 27, 68, 28, 69, 0, 0, 0, 40,
	258, 48, 0, 0, 49, 41, 42, 0, 59, 0,
	50, 0, 0, 0, 0, 0, 0, 0, 58, 0,
	0, 70, 60, 0, 63, 64, 0, 0, 0, 44,
	45, 46, 47, 0, 0, 206, 207, 0, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 66, 0, 62,
	61, 26, 27, 68, 28, 69, 0, 0, 0, 40,
	256, 48, 0, 0, 49, 41, 42, 0, 59, 0,
	50, 0, 0, 0, 0, 0, 0, 0, 58, 0,
	0, 70, 60, 0, 63, 64, 0, 0, 0, 44,
	45, 46, 47, 0, 0, 206, 207, 0, 26, 27,
	68, 28, 69, 0, 0, 65, 40, 66, 48, 62,
	61, 49, 41, 42, 0, 59, 0, 50, 0, 0,
	0, 0, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 44, 45, 46, 47,
	0, 0, 206, 207, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 0, 66, 246, 62, 61, 26, 27,
	68, 28, 69, 0, 0, 0, 40, 241, 48, 0,
	0, 49, 41, 42, 0, 59, 0, 50, 0, 0,
	0, 0, 0, 0, 0, 58, 0, 0, 70, 60,
	0, 63, 64, 0, 0, 0, 44, 45, 46, 47,
	0, 0, 206, 207, 0, 26, 27, 68, 28, 69,
	0, 0, 65, 40, 66, 48, 62, 61, 49, 41,
	42, 0, 59, 0, 50, 0, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 70, 60, 0, 63, 64,
	0, 0, 0, 44, 45, 46, 47, 0, 0, 206,
	207, 0, 26, 27, 68, 28, 69, 228, 0, 65,
	40, 66, 48, 62, 61, 49, 41, 42, 0, 59,
	0, 50, 0, 0, 0, 0, 0, 0, 0, 58,
	0, 0, 70, 60, 0, 63, 64, 0, 0, 0,
	44, 45, 46, 47, 0, 0, 0, 227, 26, 99,
	68, 100, 90, 0, 0, 105, 65, 0, 66, 0,
	62, 61, 26, 99, 68, 100, 69, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 70, 0,
	0, 63, 64, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 70, 0, 0, 63, 64, 0, 0, 0,
	0, 0, 93, 0, 106, 0, 62, 61, 0, 0,
	0, 0, 0, 0, 0, 0, 65, 0, 66, 0,
	62, 61,
}
var RubyPact = []int{

	-31, 1797, -1000, -1000, -1000, 6, -1000, -1000, -1000, 881,
	-1000, -1000, -1000, 28, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 452, 108, 7,
	5, 0, -1000, -1000, -1000, -1000, -1000, 182, -20, -1000,
	296, 195, 195, 90, 934, 934, 934, 934, 934, 934,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 2527, 934,
	2527, 17, 268, -1000, -1000, 2527, -1000, -18, 167, -1000,
	14, -1000, -1000, -1000, 934, 300, 2527, 2527, 2527, 2527,
	934, 2527, 2527, 2527, 2527, 934, 2527, 2527, 2527, 934,
	2527, -1000, 113, 2527, 2527, 299, 145, 89, -1000, 2513,
	153, -1000, -1000, -1000, 4, -23, -23, 2527, 934, 934,
	934, 934, 2527, 2527, 934, 934, 297, 51, 59, 8,
	-1000, -1000, 934, 293, 36, 36, 36, 36, 36, 123,
	2420, 85, 89, 48, 89, -1000, -1000, 163, -1000, -1000,
	32, 16, 174, -1000, 1126, 202, 2527, 2467, -1000, -23,
	36, 848, 89, 89, 89, 89, 36, 89, 89, 89,
	89, 36, 117, 89, 89, 36, 287, 174, -1000, 599,
	86, -1000, -1000, 1066, 2373, -1000, 289, -1000, 2313, 280,
	36, 36, 36, 36, 89, 89, 36, 36, -1000, -1000,
	286, 143, 63, -1000, -1, 273, 262, -1000, 2266, 195,
	2206, 36, -1000, 1750, -1000, -1000, -1000, -1000, 881, 36,
	211, 2527, -1000, 1, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 137, 142, 567, -1000, 278, 36, -1000, -1000, 113,
	14, 934, 2527, 2527, 14, -1000, 128, -2, -27, 89,
	-1000, -1000, 2146, 19, -1000, 2086, -1000, -1000, 1690, 59,
	63, 276, 934, -1000, -1000, 1630, -1000, -1000, -1000, -1000,
	266, 934, -1000, 1570, 2039, -1000, -1000, 186, 89, 1510,
	104, 2527, 1031, 9, -25, -1000, 934, 2527, -1000, -1000,
	36, 229, 89, -1000, 552, 89, -1000, 252, 934, 247,
	-1000, -1000, 232, -1000, -1000, 193, -1000, -1000, 881, 36,
	-1000, -1000, 170, 2527, -1000, -1000, -1000, 36, -1000, 188,
	1979, -1000, 934, -1000, 36, 2420, -1000, 238, -1000, 2420,
	213, -1000, -1000, -1000, 881, 36, -1000, -1000, 934, 934,
	94, 49, -1000, -1000, 2527, 85, 174, -1000, -1000, 934,
	-1000, 147, 881, 36, 89, -1000, 226, -1000, 36, -1000,
	-1000, -1000, -1000, 934, 934, 85, 1450, -1000, -1000, 36,
	2420, 2420, -1000, 2420, 220, 37, 110, 934, 934, 934,
	934, 1919, 85, 2420, 256, -15, -10, 60, 36, 36,
	-1000, 184, 2420, -1000, -1000, -1000, -1000, 36, 36, 36,
	36, -1000, 2420, -10, 934, 2527, -1000, -1000, 2420, 1330,
	1269, 1390, -10, 175, 140, -1000, 234, 934, -1000, -1000,
	228, -1000, -10, -1000, -10, -1000, -1000, 934, -1000, 36,
	1859, -1000, -10, -10, 36, 1859, 1859, 1859,
}
var RubyPgo = []int{

	0, 294, 368, 366, 32, 365, 363, 360, 968, 359,
	5, 16, 356, 355, 353, 0, 750, 647, 352, 351,
	350, 8, 349, 10, 388, 348, 11, 347, 346, 345,
	337, 336, 334, 333, 332, 6, 331, 330, 329, 328,
	327, 326, 325, 324, 322, 321, 319, 882, 317, 1,
	13, 17, 2, 314, 9, 313, 4, 310, 14, 7,
	309, 3, 308, 262, 12, 15, 235, 149,
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
	64, 64, 63, 63, 63, 18, 18, 18, 18, 18,
	18, 27, 27, 27, 27, 59, 59, 59, 59, 59,
	59, 54, 54, 23, 23, 23, 23, 65, 65, 65,
	22, 22, 25, 26, 26, 66, 66, 13, 13, 13,
	13, 13, 13, 13, 8, 8, 24, 24, 16, 16,
	36, 36, 37, 38, 39, 40, 41, 42, 43, 44,
	45, 46, 2, 5, 6, 6, 3, 3, 55, 55,
	55, 55, 62, 62, 62, 4, 4, 4, 4, 51,
	60, 60, 60, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 52, 52, 52, 52, 48, 48, 48,
	7, 14, 10, 10, 10, 57, 57, 49, 49, 19,
	20, 11, 32, 56, 56, 56, 56, 56, 56, 56,
	33, 33, 33, 33, 33, 33, 33, 34, 34, 34,
	34, 34, 35, 35, 35, 35, 31, 30, 9, 29,
	29, 28, 28,
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
	4, 5, 1, 3, 3, 5, 6, 7, 8, 5,
	6, 0, 1, 3, 3, 0, 2, 2, 2, 2,
	2, 1, 3, 1, 2, 3, 2, 0, 1, 3,
	4, 6, 4, 1, 3, 1, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 2, 2, 2, 3, 3,
	3, 3, 2, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 1, 1, 3, 3, 5, 5, 0, 4,
	7, 8, 3, 7, 8, 3, 4, 4, 3, 3,
	0, 1, 3, 4, 5, 3, 3, 3, 3, 3,
	5, 6, 5, 4, 3, 3, 2, 0, 2, 2,
	3, 4, 2, 3, 5, 0, 2, 1, 2, 2,
	2, 5, 5, 0, 2, 2, 2, 2, 2, 2,
	0, 1, 3, 3, 1, 3, 3, 5, 6, 5,
	6, 5, 4, 3, 3, 2, 3, 3, 2, 5,
	7, 4, 5,
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
	-4, 14, -47, -60, 6, -47, 62, 10, -59, 6,
	11, -65, 42, 6, 6, -59, 14, -26, 14, 14,
	-52, 17, 16, -47, -47, 14, -10, 25, -15, -56,
	-28, 37, -67, -67, -67, 11, 4, 51, 10, -4,
	-1, -58, -15, -4, -67, -15, -4, 56, 42, 56,
	14, 54, 11, 62, 14, -49, 49, 50, -15, -1,
	-27, -10, -20, 31, -54, -23, 10, -1, 14, -49,
	-47, 14, 17, 16, -1, -47, 14, -52, 14, -47,
	8, 14, 49, 50, -15, -1, -34, -33, 15, 18,
	27, 28, 14, 16, 37, -61, -15, -21, 62, 11,
	62, -67, -15, -1, -15, 10, 56, 6, -1, 6,
	6, 14, -10, 15, 18, -61, -59, 14, 14, -1,
	-47, -47, 14, -47, 4, -1, -1, 15, 18, 15,
	18, -47, -61, -47, -1, 6, -67, 6, -1, -1,
	14, -49, -47, 6, 49, 49, 50, -1, -1, -1,
	-1, 14, -47, -67, 4, 51, 10, 14, -47, -56,
	-56, -56, -67, -1, -15, 14, -35, 17, 16, 14,
	-35, 14, -67, 11, -67, 11, 14, 17, 16, -1,
	-56, 14, -67, -67, -1, -56, -56, -56,
}
var RubyDef = []int{

	1, -2, 2, 3, 4, 0, 8, 9, 10, 39,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 18, 19, -2, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	13, 32, 33, 34, 35, 36, 37, 38, 0, 0,
	0, 0, 0, 152, 153, 82, 11, 0, 58, 187,
	0, 5, 6, 7, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	-2, 61, 68, 82, 0, 0, 78, 87, 88, 19,
	-2, 21, 22, 23, 29, 13, -2, 82, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 117, 117, 13,
	-2, 13, 0, 0, 142, 143, 144, 145, 13, 0,
	195, 199, 80, 0, 11, 136, 137, 0, 134, 135,
	0, 0, 80, 84, 158, 0, 82, 0, 228, 13,
	175, 62, 69, 71, 73, 146, 147, 148, 149, 150,
	151, 177, 0, 226, 227, 179, 0, 83, 11, 80,
	127, 128, 140, 11, 0, 13, 170, 13, 0, 0,
	129, 130, 131, 132, 70, 72, 176, 178, 66, 105,
	0, 111, 117, 118, 113, 0, 0, 105, 0, 0,
	0, 133, 141, 0, 13, 13, 14, 15, 16, 17,
	0, 0, 203, 0, 138, 139, 154, 11, 155, 12,
	11, 11, 0, 19, -2, 0, 188, 189, 190, 63,
	64, 0, -2, 0, 56, 11, 0, 74, 0, 93,
	94, 165, 0, 0, 171, 0, 168, 60, 0, 117,
	0, 0, 0, 114, 116, 0, 120, 13, 122, 173,
	0, 0, 13, 0, 0, 191, 196, 13, 81, 0,
	0, 0, 0, 0, 0, 11, 0, 0, 59, 65,
	67, 0, 201, 57, 0, 89, 90, 0, 0, 0,
	166, 169, 0, 167, 95, 0, 106, 107, 39, 109,
	110, 197, 102, 0, 105, 119, 112, 115, 99, 0,
	0, 174, 0, 13, 13, 186, 180, 0, 182, 192,
	13, 202, 204, 205, 39, 207, 208, 209, 0, 0,
	211, 214, 229, 13, 0, 13, 85, 86, 156, 0,
	157, 0, 39, 11, 162, 76, 0, 91, 75, 79,
	172, 96, 198, 0, 0, 200, 0, 100, 121, 13,
	184, 185, 181, 193, 0, 13, 0, 0, 0, 0,
	0, 0, 13, 11, 0, 0, 159, 0, 103, 104,
	97, 0, 183, 13, 203, -2, -2, 212, 213, 215,
	216, 230, 11, 231, 0, 0, 77, 98, 194, 0,
	0, 0, 232, 11, 11, 217, 0, 0, 203, 219,
	0, 221, 160, 11, 163, 11, 218, 0, 203, 203,
	210, 220, 161, 164, 203, 210, 210, 210,
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
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 96:
		//line parser.y:446
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 97:
		//line parser.y:455
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target: RubyS[Rubypt-5].genericValue,
				Name:   RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-2].genericSlice,
				Body:   RubyS[Rubypt-1].genericSlice,
			}
		}
	case 98:
		//line parser.y:464
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Target:  RubyS[Rubypt-6].genericValue,
				Name:    RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 99:
		//line parser.y:474
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: ast.BareReference{Name: RubyS[Rubypt-3].operator},
				Args: RubyS[Rubypt-2].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 100:
		//line parser.y:482
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name:    ast.BareReference{Name: RubyS[Rubypt-4].operator},
				Args:    RubyS[Rubypt-3].genericSlice,
				Body:    RubyS[Rubypt-2].genericSlice,
				Rescues: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 101:
		//line parser.y:492
		{
		}
	case 102:
		//line parser.y:494
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 103:
		//line parser.y:496
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{RubyS[Rubypt-2].genericValue}}
		}
	case 104:
		//line parser.y:498
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{ast.Break{}},
			}
		}
	case 105:
		//line parser.y:506
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 106:
		//line parser.y:508
		{
		}
	case 107:
		//line parser.y:510
		{
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
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 111:
		//line parser.y:521
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 112:
		//line parser.y:523
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 113:
		//line parser.y:526
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference)}
		}
	case 114:
		//line parser.y:528
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsSplat: true}
		}
	case 115:
		//line parser.y:530
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference), DefaultValue: RubyS[Rubypt-0].genericValue}
		}
	case 116:
		//line parser.y:532
		{
			RubyVAL.genericValue = ast.MethodParam{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference), IsProc: true}
		}
	case 117:
		//line parser.y:534
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 118:
		//line parser.y:536
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 119:
		//line parser.y:540
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 120:
		//line parser.y:545
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 121:
		//line parser.y:552
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-4].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-2].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-4].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 122:
		//line parser.y:562
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-2].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-2].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 123:
		//line parser.y:571
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 124:
		//line parser.y:577
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-2].stringSlice, "::"),
			}
		}
	case 125:
		//line parser.y:585
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 126:
		//line parser.y:589
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 127:
		//line parser.y:594
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 128:
		//line parser.y:601
		{
			RubyVAL.genericValue = ast.Assignment{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 129:
		//line parser.y:605
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 130:
		//line parser.y:612
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 131:
		//line parser.y:619
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 132:
		//line parser.y:626
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 133:
		//line parser.y:633
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 134:
		//line parser.y:641
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 135:
		//line parser.y:643
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 136:
		//line parser.y:646
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 137:
		//line parser.y:648
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 138:
		//line parser.y:651
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 139:
		//line parser.y:653
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 140:
		//line parser.y:656
		{
			RubyVAL.genericValue = ast.Array{Nodes: []ast.Node{RubyS[Rubypt-2].genericValue, RubyS[Rubypt-0].genericValue}}
		}
	case 141:
		//line parser.y:658
		{
			RubyVAL.genericValue = ast.Array{Nodes: append(RubyVAL.genericValue.(ast.Array).Nodes, RubyS[Rubypt-0].genericValue)}
		}
	case 142:
		//line parser.y:660
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 143:
		//line parser.y:661
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 144:
		//line parser.y:662
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 145:
		//line parser.y:663
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 146:
		//line parser.y:666
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 147:
		//line parser.y:675
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 148:
		//line parser.y:684
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 149:
		//line parser.y:693
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 150:
		//line parser.y:702
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 151:
		//line parser.y:711
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 152:
		//line parser.y:719
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 153:
		//line parser.y:720
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 154:
		//line parser.y:722
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 155:
		//line parser.y:724
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-1].genericSlice}
		}
	case 156:
		//line parser.y:727
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 157:
		//line parser.y:735
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-2].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 158:
		//line parser.y:743
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 159:
		//line parser.y:745
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 160:
		//line parser.y:752
		{
			if RubyS[Rubypt-2].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-3].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 161:
		//line parser.y:759
		{
			if RubyS[Rubypt-3].operator != "=>" {
				panic("FREAKOUT")
			}
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-2].genericValue})
		}
	case 162:
		//line parser.y:767
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-2].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 163:
		//line parser.y:774
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 164:
		//line parser.y:781
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-2].genericValue,
			})
		}
	case 165:
		//line parser.y:789
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 166:
		//line parser.y:791
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 167:
		//line parser.y:798
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 168:
		//line parser.y:805
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 169:
		//line parser.y:808
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 170:
		//line parser.y:810
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 171:
		//line parser.y:812
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 172:
		//line parser.y:814
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 173:
		//line parser.y:817
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 174:
		//line parser.y:824
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 175:
		//line parser.y:832
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 176:
		//line parser.y:839
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-0].genericValue,
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 177:
		//line parser.y:846
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      []ast.Node{RubyS[Rubypt-2].genericValue},
			}
		}
	case 178:
		//line parser.y:853
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 179:
		//line parser.y:860
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue},
				Body:      ast.Nodes{RubyS[Rubypt-2].genericValue},
			}
		}
	case 180:
		//line parser.y:867
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 181:
		//line parser.y:874
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 182:
		//line parser.y:882
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 183:
		//line parser.y:891
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 184:
		//line parser.y:898
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 185:
		//line parser.y:905
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 186:
		//line parser.y:912
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 187:
		//line parser.y:919
		{
		}
	case 188:
		//line parser.y:920
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 189:
		//line parser.y:921
		{
		}
	case 190:
		//line parser.y:924
		{
			RubyVAL.genericValue = ast.Group{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 191:
		//line parser.y:927
		{
			RubyVAL.genericValue = ast.Begin{
				Body:   RubyS[Rubypt-2].genericSlice,
				Rescue: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 192:
		//line parser.y:935
		{
			RubyVAL.genericValue = ast.Rescue{Body: RubyS[Rubypt-0].genericSlice}
		}
	case 193:
		//line parser.y:937
		{
			RubyVAL.genericValue = ast.Rescue{
				Body: RubyS[Rubypt-0].genericSlice,
				Exception: ast.RescueException{
					Class: RubyS[Rubypt-1].genericValue.(ast.BareReference),
				},
			}
		}
	case 194:
		//line parser.y:946
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
	case 195:
		//line parser.y:961
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 196:
		//line parser.y:963
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 197:
		//line parser.y:966
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 198:
		//line parser.y:968
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 199:
		//line parser.y:971
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Yield{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 200:
		//line parser.y:980
		{
			if len(RubyS[Rubypt-0].genericSlice) == 1 {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice[0]}
			} else {
				RubyVAL.genericValue = ast.Return{Value: RubyS[Rubypt-0].genericSlice}
			}
		}
	case 201:
		//line parser.y:989
		{
			RubyVAL.genericValue = ast.Ternary{
				Condition: RubyS[Rubypt-4].genericValue,
				True:      RubyS[Rubypt-2].genericValue,
				False:     RubyS[Rubypt-0].genericValue,
			}
		}
	case 202:
		//line parser.y:998
		{
			RubyVAL.genericValue = ast.Loop{Condition: RubyS[Rubypt-3].genericValue, Body: RubyS[Rubypt-1].genericSlice}
		}
	case 203:
		//line parser.y:1001
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 204:
		//line parser.y:1003
		{
		}
	case 205:
		//line parser.y:1005
		{
		}
	case 206:
		//line parser.y:1007
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 207:
		//line parser.y:1009
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
		//line parser.y:1015
		{
		}
	case 211:
		//line parser.y:1017
		{
			RubyVAL.genericValue = ast.Break{}
		}
	case 212:
		//line parser.y:1019
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Break{}}}
		}
	case 213:
		//line parser.y:1021
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Break{}}}
		}
	case 214:
		//line parser.y:1023
		{
			RubyVAL.genericValue = ast.Next{}
		}
	case 215:
		//line parser.y:1025
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: RubyS[Rubypt-0].genericValue, Body: []ast.Node{ast.Next{}}}
		}
	case 216:
		//line parser.y:1027
		{
			RubyVAL.genericValue = ast.IfBlock{Condition: ast.Negation{Target: RubyS[Rubypt-0].genericValue}, Body: []ast.Node{ast.Next{}}}
		}
	case 217:
		//line parser.y:1030
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 218:
		//line parser.y:1037
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-4].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 219:
		//line parser.y:1045
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 220:
		//line parser.y:1052
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-4].genericValue},
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 221:
		//line parser.y:1060
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: ast.Negation{Target: RubyS[Rubypt-3].genericValue},
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 222:
		//line parser.y:1068
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 223:
		//line parser.y:1075
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 224:
		//line parser.y:1082
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 225:
		//line parser.y:1089
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 226:
		//line parser.y:1097
		{
			RubyVAL.genericValue = ast.WeakLogicalAnd{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 227:
		//line parser.y:1100
		{
			RubyVAL.genericValue = ast.WeakLogicalOr{LHS: RubyS[Rubypt-2].genericValue, RHS: RubyS[Rubypt-0].genericValue}
		}
	case 228:
		//line parser.y:1102
		{
			RubyVAL.genericValue = ast.Lambda{Body: RubyS[Rubypt-0].genericValue.(ast.Block)}
		}
	case 229:
		//line parser.y:1105
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-3].genericValue, Cases: RubyS[Rubypt-1].switchCaseSlice}
		}
	case 230:
		//line parser.y:1107
		{
			RubyVAL.genericValue = ast.SwitchStatement{Condition: RubyS[Rubypt-5].genericValue, Cases: RubyS[Rubypt-3].switchCaseSlice, Else: RubyS[Rubypt-1].genericSlice}
		}
	case 231:
		//line parser.y:1110
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	case 232:
		//line parser.y:1112
		{
			RubyVAL.switchCaseSlice = append(RubyVAL.switchCaseSlice, ast.SwitchCase{Conditions: RubyS[Rubypt-2].genericSlice, Body: RubyS[Rubypt-1].genericSlice})
		}
	}
	goto Rubystack /* stack new state and value */
}
