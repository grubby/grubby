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
}
var RubyStatenames = []string{}

const RubyEofCode = 1
const RubyErrCode = 2
const RubyMaxDepth = 200

//line parser.y:269

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 7,
	7, 54,
	9, 54,
	18, 54,
	-2, 12,
	-1, 38,
	9, 33,
	-2, 28,
	-1, 54,
	9, 35,
	-2, 29,
	-1, 55,
	9, 34,
	-2, 30,
	-1, 57,
	9, 54,
	-2, 32,
	-1, 63,
	18, 54,
	-2, 12,
	-1, 73,
	9, 33,
	-2, 28,
	-1, 76,
	24, 44,
	-2, 54,
	-1, 108,
	24, 45,
	-2, 54,
}

const RubyNprod = 59
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 260

var RubyAct = []int{

	79, 52, 6, 106, 36, 37, 90, 75, 38, 100,
	57, 59, 24, 25, 78, 88, 101, 74, 49, 50,
	51, 53, 93, 103, 88, 39, 40, 41, 42, 108,
	44, 76, 4, 45, 46, 47, 48, 104, 105, 54,
	55, 58, 56, 60, 43, 60, 54, 55, 77, 56,
	67, 68, 69, 53, 70, 1, 61, 71, 72, 73,
	87, 86, 23, 98, 80, 2, 22, 21, 20, 19,
	18, 17, 13, 11, 10, 89, 60, 92, 62, 64,
	65, 66, 85, 14, 3, 94, 95, 91, 12, 96,
	97, 6, 16, 15, 0, 6, 0, 0, 0, 0,
	81, 82, 83, 84, 0, 0, 0, 107, 6, 111,
	6, 0, 0, 109, 8, 7, 9, 0, 0, 0,
	26, 112, 27, 29, 30, 31, 0, 0, 0, 32,
	33, 34, 35, 0, 28, 0, 5, 24, 25, 8,
	7, 9, 0, 0, 0, 26, 110, 27, 29, 30,
	31, 0, 0, 0, 32, 33, 34, 35, 0, 28,
	0, 5, 24, 25, 8, 7, 9, 0, 0, 0,
	26, 102, 27, 29, 30, 31, 0, 0, 0, 32,
	33, 34, 35, 0, 28, 0, 5, 24, 25, 8,
	7, 9, 0, 0, 0, 26, 99, 27, 29, 30,
	31, 0, 0, 0, 32, 33, 34, 35, 0, 28,
	0, 5, 24, 25, 8, 7, 9, 0, 0, 0,
	26, 0, 27, 29, 30, 31, 0, 0, 0, 32,
	33, 34, 35, 0, 28, 0, 5, 24, 25, 8,
	63, 9, 0, 0, 0, 26, 0, 27, 29, 30,
	31, 0, 0, 0, 32, 33, 34, 35, 0, 28,
}
var RubyPact = []int{

	-1000, 210, -1000, -15, -15, -1000, -1000, -15, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -15, -15, -15, -15, 39, -15,
	-1000, -1000, -15, -15, -15, -15, -1000, -3, 35, -1000,
	-1000, 36, -15, -1000, -15, 235, 235, 235, 235, -15,
	-15, -15, -1000, -15, -1000, -1000, -15, -15, -15, -9,
	25, -12, -15, -15, -15, -15, -15, 235, 235, 235,
	235, 56, 6, 42, -1000, -15, -15, -2, -1000, -3,
	3, -15, -15, -15, -15, -15, -1000, -1000, -15, -15,
	185, -17, -1000, -8, 160, 15, 33, -23, -1000, -1000,
	-1000, 23, -1000, -1000, -1000, -1000, -1000, 135, -15, 110,
	-1000, -1000, -1000,
}
var RubyPgo = []int{

	0, 0, 32, 93, 92, 88, 84, 63, 83, 74,
	73, 72, 11, 71, 70, 69, 68, 67, 66, 62,
	6, 1, 55, 10, 48,
}
var RubyR1 = []int{

	0, 22, 22, 20, 20, 5, 7, 7, 7, 7,
	6, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 21, 21,
	21, 21, 21, 23, 23, 23, 23, 23, 9, 10,
	10, 11, 12, 12, 24, 24, 8, 13, 14, 15,
	16, 17, 18, 19, 1, 1, 1, 3, 4,
}
var RubyR2 = []int{

	0, 0, 2, 0, 2, 2, 2, 2, 1, 1,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 0, 1,
	1, 5, 1, 0, 1, 1, 5, 5, 9, 6,
	8, 6, 3, 6, 1, 4, 5, 3, 3, 3,
	3, 5, 5, 5, 0, 2, 2, 1, 1,
}
var RubyChk = []int{

	-1000, -22, -7, -6, -2, 26, -1, 5, 4, 6,
	-9, -10, -5, -11, -8, -3, -4, -13, -14, -15,
	-16, -17, -18, -19, 27, 28, 10, 12, 24, 13,
	14, 15, 19, 20, 21, 22, -1, -1, -1, -1,
	-1, -1, -1, 5, -1, -1, -1, -1, -1, 21,
	22, 23, -21, 18, 4, 5, 7, -23, 5, -12,
	-1, -12, -2, 5, -2, -2, -2, -1, -1, -1,
	-1, -1, -1, -1, 26, 16, 6, -24, 26, -1,
	-1, -2, -2, -2, -2, -23, 5, 4, 9, -21,
	-20, -12, -1, 24, -20, -1, -1, -1, -7, 11,
	26, 24, 11, 8, 4, 5, 26, -20, 6, -20,
	11, -1, 11,
}
var RubyDef = []int{

	1, -2, 2, 54, 54, 8, 9, -2, 11, 13,
	14, 15, 16, 17, 18, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 54, 54, 54, 54, 0, 54,
	57, 58, 54, 54, 54, 54, 6, 7, -2, 55,
	56, 0, 54, 5, 54, 0, 0, 0, 0, 54,
	54, 54, 10, 54, -2, -2, 54, -2, 54, 0,
	0, 0, 47, -2, 48, 49, 50, 0, 0, 0,
	0, 33, 0, -2, 3, 54, -2, 0, 3, 0,
	0, 51, 52, 53, 46, 54, 34, 35, 54, 54,
	0, 0, 42, 0, 0, 0, 0, 0, 4, 39,
	3, 0, 41, 31, 36, 37, 3, 0, -2, 0,
	40, 43, 38,
}
var RubyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 28,
	26, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 27,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25,
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
		//line parser.y:97
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:99
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 3:
		//line parser.y:106
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 4:
		//line parser.y:108
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
			}
		}
	case 5:
		//line parser.y:115
		{
			RubyVAL.genericValue = ast.Symbol{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 6:
		//line parser.y:120
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 7:
		//line parser.y:122
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 8:
		//line parser.y:124
		{
			RubyVAL.genericValue = nil /* ignores new lines around statements */
		}
	case 9:
		//line parser.y:126
		{
			RubyVAL.genericValue = nil /* ignores whitespace around statements */
		}
	case 10:
		//line parser.y:129
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
		//line parser.y:138
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 29:
		//line parser.y:140
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 30:
		//line parser.y:142
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 31:
		//line parser.y:144
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 32:
		//line parser.y:148
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 33:
		//line parser.y:152
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 34:
		//line parser.y:154
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 35:
		//line parser.y:156
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 36:
		//line parser.y:158
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 37:
		//line parser.y:160
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 38:
		//line parser.y:165
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 39:
		//line parser.y:174
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 40:
		//line parser.y:181
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 41:
		//line parser.y:191
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 42:
		//line parser.y:200
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 43:
		//line parser.y:206
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 44:
		//line parser.y:214
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 45:
		//line parser.y:218
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 46:
		//line parser.y:223
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 47:
		//line parser.y:230
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 48:
		//line parser.y:231
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 49:
		//line parser.y:232
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 50:
		//line parser.y:233
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 51:
		//line parser.y:236
		{
			RubyVAL.genericValue = ast.Addition{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 52:
		//line parser.y:244
		{
			RubyVAL.genericValue = ast.Subtraction{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 53:
		//line parser.y:252
		{
			RubyVAL.genericValue = ast.Multiplication{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 54:
		//line parser.y:260
		{
			RubyVAL.genericValue = nil
		}
	case 55:
		//line parser.y:262
		{
			RubyVAL.genericValue = nil
		}
	case 56:
		//line parser.y:264
		{
			RubyVAL.genericValue = nil
		}
	case 57:
		//line parser.y:266
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 58:
		//line parser.y:267
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	}
	goto Rubystack /* stack new state and value */
}
