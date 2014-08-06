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
	yys          int
	genericValue ast.Node
	genericSlice ast.Nodes
	stringSlice  []string
}

const NODE = 57346
const REF = 57347
const CAPITAL_REF = 57348
const LPAREN = 57349
const RPAREN = 57350
const COMMA = 57351
const DEF = 57352
const END = 57353
const CLASS = 57354
const MODULE = 57355
const TRUE = 57356
const FALSE = 57357
const LESSTHAN = 57358
const GREATERTHAN = 57359
const EQUALTO = 57360
const BANG = 57361
const COMPLEMENT = 57362
const POSITIVE = 57363
const NEGATIVE = 57364
const STAR = 57365
const HASHROCKET = 57366
const COLON = 57367
const DOT = 57368
const LBRACKET = 57369
const RBRACKET = 57370
const LBRACE = 57371
const RBRACE = 57372

var RubyToknames = []string{
	"NODE",
	"REF",
	"CAPITAL_REF",
	"LPAREN",
	"RPAREN",
	"COMMA",
	"DEF",
	"END",
	"CLASS",
	"MODULE",
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
	"HASHROCKET",
	"COLON",
	"DOT",
	"LBRACKET",
	"RBRACKET",
	"LBRACE",
	"RBRACE",
}
var RubyStatenames = []string{}

const RubyEofCode = 1
const RubyErrCode = 2
const RubyMaxDepth = 200

//line parser.y:303

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 7,
	7, 56,
	9, 56,
	18, 56,
	-2, 12,
	-1, 42,
	9, 35,
	-2, 30,
	-1, 60,
	9, 37,
	-2, 31,
	-1, 61,
	9, 36,
	-2, 32,
	-1, 63,
	9, 56,
	-2, 34,
	-1, 69,
	18, 56,
	-2, 12,
	-1, 84,
	9, 35,
	-2, 30,
	-1, 87,
	25, 46,
	-2, 56,
	-1, 126,
	25, 47,
	-2, 56,
	-1, 139,
	21, 56,
	22, 56,
	23, 56,
	-2, 67,
	-1, 143,
	21, 56,
	22, 56,
	23, 56,
	-2, 68,
}

const RubyNprod = 69
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 349

var RubyAct = []int{

	90, 4, 6, 76, 40, 41, 65, 101, 42, 58,
	26, 27, 128, 63, 137, 108, 86, 124, 115, 89,
	77, 116, 99, 55, 56, 57, 141, 43, 44, 45,
	46, 85, 48, 104, 77, 49, 50, 51, 52, 53,
	54, 106, 55, 56, 57, 129, 59, 66, 99, 66,
	126, 68, 70, 71, 72, 67, 78, 79, 80, 64,
	81, 121, 99, 82, 83, 84, 87, 73, 122, 123,
	91, 55, 56, 57, 92, 60, 61, 93, 62, 47,
	94, 95, 96, 97, 75, 74, 88, 66, 103, 59,
	113, 107, 2, 102, 100, 109, 98, 105, 1, 110,
	111, 112, 6, 23, 60, 61, 6, 62, 118, 119,
	120, 22, 21, 20, 19, 18, 17, 13, 11, 10,
	14, 3, 127, 125, 12, 24, 6, 132, 133, 134,
	135, 6, 130, 16, 25, 15, 138, 139, 0, 140,
	0, 0, 142, 0, 143, 8, 7, 9, 0, 0,
	0, 28, 136, 29, 31, 32, 33, 0, 0, 0,
	34, 35, 36, 37, 0, 0, 30, 0, 38, 0,
	39, 0, 5, 26, 27, 8, 7, 9, 0, 0,
	0, 28, 131, 29, 31, 32, 33, 0, 0, 0,
	34, 35, 36, 37, 0, 0, 30, 0, 38, 0,
	39, 0, 5, 26, 27, 8, 7, 9, 0, 0,
	0, 28, 117, 29, 31, 32, 33, 0, 0, 0,
	34, 35, 36, 37, 0, 0, 30, 0, 38, 0,
	39, 0, 5, 26, 27, 8, 7, 9, 0, 0,
	0, 28, 114, 29, 31, 32, 33, 0, 0, 0,
	34, 35, 36, 37, 0, 0, 30, 0, 38, 0,
	39, 0, 5, 26, 27, 8, 7, 9, 0, 0,
	0, 28, 0, 29, 31, 32, 33, 0, 0, 0,
	34, 35, 36, 37, 0, 0, 30, 0, 38, 0,
	39, 0, 5, 26, 27, 8, 69, 9, 0, 0,
	0, 28, 0, 29, 31, 32, 33, 0, 0, 0,
	34, 35, 36, 37, 0, 0, 30, 0, 38, 0,
	39, 0, 108, 8, 69, 9, 0, 0, 0, 28,
	0, 29, 31, 32, 33, 0, 0, 0, 34, 35,
	36, 37, 0, 0, 30, 0, 38, 0, 39,
}
var RubyPact = []int{

	-1000, 261, -1000, -22, -22, -1000, -1000, -22, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -22, -22, -22, -22,
	74, -22, -1000, -1000, -22, -22, -22, -22, -22, -22,
	-1000, 50, 71, -1000, -1000, 54, -22, -1000, -22, 319,
	319, 319, 319, 80, -11, -22, -22, -22, -1000, -22,
	-1000, -1000, -22, -22, -22, 0, 60, -12, -22, -22,
	-22, -22, -22, -22, -1000, -1000, -22, -1000, 319, 319,
	319, 319, 80, 39, 100, -1000, -22, -22, 8, -1000,
	50, 28, 13, 291, -22, -22, -22, -22, -22, -22,
	-22, 231, -13, -1000, -4, 201, -1000, -22, -22, -22,
	53, 64, -14, -1000, -1000, -1000, 44, -1000, 3, -1000,
	21, -1000, -1000, -1000, -1000, 171, -22, -22, -22, -22,
	141, -1000, -1000, -16, 319, 319, -1000, -1000, -22, -22,
	2, -22, 319, -22,
}
var RubyPgo = []int{

	0, 0, 3, 1, 135, 134, 133, 125, 124, 121,
	90, 120, 119, 118, 117, 6, 116, 115, 114, 113,
	112, 111, 103, 7, 9, 98, 91, 13, 86,
}
var RubyR1 = []int{

	0, 25, 25, 23, 23, 8, 10, 10, 10, 10,
	9, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	24, 24, 24, 24, 24, 27, 27, 27, 27, 27,
	12, 13, 13, 14, 15, 15, 28, 28, 11, 16,
	17, 18, 19, 20, 21, 22, 1, 1, 1, 4,
	6, 7, 5, 2, 2, 2, 26, 26, 26,
}
var RubyR2 = []int{

	0, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	0, 1, 1, 5, 1, 0, 1, 1, 5, 5,
	9, 6, 8, 6, 3, 6, 1, 4, 5, 3,
	3, 3, 3, 5, 5, 5, 0, 2, 2, 1,
	1, 5, 9, 0, 1, 4, 0, 5, 9,
}
var RubyChk = []int{

	-1000, -25, -10, -9, -3, 31, -1, 5, 4, 6,
	-12, -13, -8, -14, -11, -4, -6, -16, -17, -18,
	-19, -20, -21, -22, -7, -5, 32, 33, 10, 12,
	25, 13, 14, 15, 19, 20, 21, 22, 27, 29,
	-1, -1, -1, -1, -1, -1, -1, 5, -1, -1,
	-1, -1, -1, -1, -1, 21, 22, 23, -24, 18,
	4, 5, 7, -27, 5, -15, -1, -15, -3, 5,
	-3, -3, -3, -27, 5, 4, -2, 31, -1, -1,
	-1, -1, -1, -1, -1, 31, 16, 6, -28, 31,
	-1, -1, -1, -1, -3, -3, -3, -3, -27, 9,
	-24, -23, -15, -1, 25, -23, 28, -26, 31, -3,
	-1, -1, -1, -10, 11, 31, 25, 11, -1, -1,
	-1, 8, 4, 5, 31, -23, 6, -2, 9, 24,
	-23, 11, -1, -1, -1, -1, 11, 30, -3, -3,
	-1, 24, -1, -3,
}
var RubyDef = []int{

	1, -2, 2, 56, 56, 8, 9, -2, 11, 13,
	14, 15, 16, 17, 18, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 28, 29, 56, 56, 56, 56,
	0, 56, 59, 60, 56, 56, 56, 56, 56, 56,
	6, 7, -2, 57, 58, 0, 56, 5, 56, 0,
	0, 0, 0, 35, 63, 56, 56, 56, 10, 56,
	-2, -2, 56, -2, 56, 0, 0, 0, 49, -2,
	50, 51, 52, 56, 36, 37, 56, 64, 0, 0,
	0, 0, 35, 0, -2, 3, 56, -2, 0, 3,
	0, 0, 0, 66, 53, 54, 55, 48, 56, 56,
	56, 0, 0, 44, 0, 0, 61, 56, 56, 56,
	0, 0, 0, 4, 41, 3, 0, 43, 63, 65,
	0, 33, 38, 39, 3, 0, -2, 56, 56, 56,
	0, 42, 45, 0, 0, 0, 40, 62, 56, -2,
	0, 56, 0, -2,
}
var RubyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 33,
	31, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 32,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30,
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
		//line parser.y:106
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:108
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 3:
		//line parser.y:115
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 4:
		//line parser.y:117
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
			}
		}
	case 5:
		//line parser.y:124
		{
			RubyVAL.genericValue = ast.Symbol{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 6:
		//line parser.y:129
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 7:
		//line parser.y:131
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 8:
		//line parser.y:133
		{
			RubyVAL.genericValue = nil /* ignores new lines around statements */
		}
	case 9:
		//line parser.y:135
		{
			RubyVAL.genericValue = nil /* ignores whitespace around statements */
		}
	case 10:
		//line parser.y:138
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue,
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 11:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 12:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 13:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 14:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 15:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 16:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 17:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
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
		//line parser.y:147
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 31:
		//line parser.y:149
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 32:
		//line parser.y:151
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 33:
		//line parser.y:153
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 34:
		//line parser.y:157
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 35:
		//line parser.y:161
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 36:
		//line parser.y:163
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 37:
		//line parser.y:165
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 38:
		//line parser.y:167
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 39:
		//line parser.y:169
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 40:
		//line parser.y:174
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 41:
		//line parser.y:183
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 42:
		//line parser.y:190
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 43:
		//line parser.y:200
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 44:
		//line parser.y:209
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 45:
		//line parser.y:215
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 46:
		//line parser.y:223
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 47:
		//line parser.y:227
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 48:
		//line parser.y:232
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 49:
		//line parser.y:239
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 50:
		//line parser.y:240
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 51:
		//line parser.y:241
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 52:
		//line parser.y:242
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 53:
		//line parser.y:245
		{
			RubyVAL.genericValue = ast.Addition{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 54:
		//line parser.y:253
		{
			RubyVAL.genericValue = ast.Subtraction{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 55:
		//line parser.y:261
		{
			RubyVAL.genericValue = ast.Multiplication{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 56:
		//line parser.y:269
		{
			RubyVAL.genericValue = nil
		}
	case 57:
		//line parser.y:271
		{
			RubyVAL.genericValue = nil
		}
	case 58:
		//line parser.y:273
		{
			RubyVAL.genericValue = nil
		}
	case 59:
		//line parser.y:275
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 60:
		//line parser.y:276
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 61:
		//line parser.y:278
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 62:
		//line parser.y:282
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 63:
		//line parser.y:291
		{
			RubyVAL.genericValue = nil
		}
	case 64:
		//line parser.y:293
		{
			RubyVAL.genericValue = nil
		}
	case 65:
		//line parser.y:295
		{
			RubyVAL.genericValue = nil
		}
	case 66:
		//line parser.y:297
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 67:
		//line parser.y:299
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 68:
		//line parser.y:301
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-4].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	}
	goto Rubystack /* stack new state and value */
}
