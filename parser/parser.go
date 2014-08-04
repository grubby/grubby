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
const COLON = 57366
const DOT = 57367
const LBRACKET = 57368
const RBRACKET = 57369
const LBRACE = 57370
const RBRACE = 57371

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

//line parser.y:278

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 7,
	7, 55,
	9, 55,
	18, 55,
	-2, 12,
	-1, 40,
	9, 34,
	-2, 29,
	-1, 57,
	9, 36,
	-2, 30,
	-1, 58,
	9, 35,
	-2, 31,
	-1, 60,
	9, 55,
	-2, 33,
	-1, 66,
	18, 55,
	-2, 12,
	-1, 79,
	9, 34,
	-2, 29,
	-1, 82,
	24, 45,
	-2, 55,
	-1, 114,
	24, 46,
	-2, 55,
}

const RubyNprod = 61
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 258

var RubyAct = []int{

	85, 4, 6, 81, 38, 39, 95, 55, 40, 25,
	26, 112, 60, 106, 84, 93, 62, 80, 52, 53,
	54, 56, 107, 98, 109, 93, 41, 42, 43, 44,
	93, 46, 114, 100, 47, 48, 49, 50, 51, 57,
	58, 61, 59, 110, 111, 63, 82, 63, 45, 65,
	67, 68, 69, 73, 74, 75, 83, 76, 72, 71,
	77, 78, 79, 64, 70, 57, 58, 86, 59, 1,
	104, 87, 2, 23, 22, 88, 89, 90, 91, 56,
	21, 20, 63, 97, 19, 18, 17, 94, 13, 11,
	92, 99, 10, 101, 102, 103, 6, 14, 96, 3,
	6, 12, 24, 16, 15, 0, 0, 0, 0, 0,
	0, 0, 0, 113, 6, 117, 6, 0, 0, 115,
	8, 7, 9, 0, 0, 0, 27, 118, 28, 30,
	31, 32, 0, 0, 0, 33, 34, 35, 36, 0,
	29, 0, 37, 8, 7, 9, 5, 25, 26, 27,
	116, 28, 30, 31, 32, 0, 0, 0, 33, 34,
	35, 36, 0, 29, 0, 37, 8, 7, 9, 5,
	25, 26, 27, 108, 28, 30, 31, 32, 0, 0,
	0, 33, 34, 35, 36, 0, 29, 0, 37, 8,
	7, 9, 5, 25, 26, 27, 105, 28, 30, 31,
	32, 0, 0, 0, 33, 34, 35, 36, 0, 29,
	0, 37, 8, 7, 9, 5, 25, 26, 27, 0,
	28, 30, 31, 32, 0, 0, 0, 33, 34, 35,
	36, 0, 29, 0, 37, 8, 66, 9, 5, 25,
	26, 27, 0, 28, 30, 31, 32, 0, 0, 0,
	33, 34, 35, 36, 0, 29, 0, 37,
}
var RubyPact = []int{

	-1000, 208, -1000, -22, -22, -1000, -1000, -22, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -22, -22, -22, -22, 43,
	-22, -1000, -1000, -22, -22, -22, -22, -22, -1000, -3,
	61, -1000, -1000, 36, -22, -1000, -22, 231, 231, 231,
	231, 54, -22, -22, -22, -1000, -22, -1000, -1000, -22,
	-22, -22, -13, 40, -16, -22, -22, -22, -22, -22,
	-22, -1000, -1000, 231, 231, 231, 231, 54, 21, 35,
	-1000, -22, -22, -1, -1000, -3, 3, 6, -22, -22,
	-22, -22, -22, -22, -22, 185, -17, -1000, -2, 162,
	-1000, 16, 39, -19, -1000, -1000, -1000, 26, -1000, -1000,
	-1000, -1000, -1000, 139, -22, 116, -1000, -1000, -1000,
}
var RubyPgo = []int{

	0, 0, 1, 104, 103, 102, 101, 99, 70, 97,
	92, 89, 88, 16, 86, 85, 84, 81, 80, 74,
	73, 6, 7, 69, 12, 56,
}
var RubyR1 = []int{

	0, 23, 23, 21, 21, 6, 8, 8, 8, 8,
	7, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 22,
	22, 22, 22, 22, 24, 24, 24, 24, 24, 10,
	11, 11, 12, 13, 13, 25, 25, 9, 14, 15,
	16, 17, 18, 19, 20, 1, 1, 1, 3, 4,
	5,
}
var RubyR2 = []int{

	0, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
	1, 1, 5, 1, 0, 1, 1, 5, 5, 9,
	6, 8, 6, 3, 6, 1, 4, 5, 3, 3,
	3, 3, 5, 5, 5, 0, 2, 2, 1, 1,
	5,
}
var RubyChk = []int{

	-1000, -23, -8, -7, -2, 30, -1, 5, 4, 6,
	-10, -11, -6, -12, -9, -3, -4, -14, -15, -16,
	-17, -18, -19, -20, -5, 31, 32, 10, 12, 24,
	13, 14, 15, 19, 20, 21, 22, 26, -1, -1,
	-1, -1, -1, -1, -1, 5, -1, -1, -1, -1,
	-1, -1, 21, 22, 23, -22, 18, 4, 5, 7,
	-24, 5, -13, -1, -13, -2, 5, -2, -2, -2,
	-24, 5, 4, -1, -1, -1, -1, -1, -1, -1,
	30, 16, 6, -25, 30, -1, -1, -1, -2, -2,
	-2, -2, -24, 9, -22, -21, -13, -1, 24, -21,
	27, -1, -1, -1, -8, 11, 30, 24, 11, 8,
	4, 5, 30, -21, 6, -21, 11, -1, 11,
}
var RubyDef = []int{

	1, -2, 2, 55, 55, 8, 9, -2, 11, 13,
	14, 15, 16, 17, 18, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 28, 55, 55, 55, 55, 0,
	55, 58, 59, 55, 55, 55, 55, 55, 6, 7,
	-2, 56, 57, 0, 55, 5, 55, 0, 0, 0,
	0, 34, 55, 55, 55, 10, 55, -2, -2, 55,
	-2, 55, 0, 0, 0, 48, -2, 49, 50, 51,
	55, 35, 36, 0, 0, 0, 0, 34, 0, -2,
	3, 55, -2, 0, 3, 0, 0, 0, 52, 53,
	54, 47, 55, 55, 55, 0, 0, 43, 0, 0,
	60, 0, 0, 0, 4, 40, 3, 0, 42, 32,
	37, 38, 3, 0, -2, 0, 41, 44, 39,
}
var RubyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 32,
	30, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 31,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29,
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
		//line parser.y:102
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:104
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 3:
		//line parser.y:111
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 4:
		//line parser.y:113
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
			}
		}
	case 5:
		//line parser.y:120
		{
			RubyVAL.genericValue = ast.Symbol{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 6:
		//line parser.y:125
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 7:
		//line parser.y:127
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 8:
		//line parser.y:129
		{
			RubyVAL.genericValue = nil /* ignores new lines around statements */
		}
	case 9:
		//line parser.y:131
		{
			RubyVAL.genericValue = nil /* ignores whitespace around statements */
		}
	case 10:
		//line parser.y:134
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
		//line parser.y:143
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 30:
		//line parser.y:145
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 31:
		//line parser.y:147
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 32:
		//line parser.y:149
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 33:
		//line parser.y:153
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 34:
		//line parser.y:157
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 35:
		//line parser.y:159
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 36:
		//line parser.y:161
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 37:
		//line parser.y:163
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 38:
		//line parser.y:165
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 39:
		//line parser.y:170
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 40:
		//line parser.y:179
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 41:
		//line parser.y:186
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 42:
		//line parser.y:196
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 43:
		//line parser.y:205
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 44:
		//line parser.y:211
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 45:
		//line parser.y:219
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 46:
		//line parser.y:223
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 47:
		//line parser.y:228
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 48:
		//line parser.y:235
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 49:
		//line parser.y:236
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 50:
		//line parser.y:237
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 51:
		//line parser.y:238
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 52:
		//line parser.y:241
		{
			RubyVAL.genericValue = ast.Addition{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 53:
		//line parser.y:249
		{
			RubyVAL.genericValue = ast.Subtraction{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 54:
		//line parser.y:257
		{
			RubyVAL.genericValue = ast.Multiplication{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 55:
		//line parser.y:265
		{
			RubyVAL.genericValue = nil
		}
	case 56:
		//line parser.y:267
		{
			RubyVAL.genericValue = nil
		}
	case 57:
		//line parser.y:269
		{
			RubyVAL.genericValue = nil
		}
	case 58:
		//line parser.y:271
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 59:
		//line parser.y:272
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 60:
		//line parser.y:274
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	}
	goto Rubystack /* stack new state and value */
}
