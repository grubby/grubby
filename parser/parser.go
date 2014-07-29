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
const COLON = 57365
const DOT = 57366

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
	"COLON",
	"DOT",
}
var RubyStatenames = []string{}

const RubyEofCode = 1
const RubyErrCode = 2
const RubyMaxDepth = 200

//line parser.y:252

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 7,
	18, 52,
	-2, 13,
	-1, 53,
	18, 52,
	-2, 13,
	-1, 75,
	23, 42,
	-2, 52,
	-1, 97,
	23, 43,
	-2, 52,
}

const RubyNprod = 57
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 239

var RubyAct = []int{

	36, 91, 6, 81, 59, 35, 37, 80, 38, 4,
	22, 23, 74, 39, 40, 77, 41, 22, 23, 54,
	92, 73, 84, 42, 43, 44, 45, 97, 47, 56,
	75, 48, 49, 50, 51, 22, 23, 39, 40, 58,
	41, 46, 57, 34, 52, 55, 60, 76, 60, 86,
	87, 56, 61, 69, 66, 67, 1, 68, 62, 63,
	64, 65, 21, 39, 40, 72, 41, 98, 99, 20,
	79, 71, 70, 19, 18, 60, 83, 17, 78, 82,
	13, 85, 6, 89, 88, 2, 6, 11, 94, 6,
	10, 14, 3, 12, 16, 96, 15, 6, 101, 8,
	7, 9, 0, 0, 0, 24, 100, 25, 27, 28,
	29, 0, 0, 0, 30, 31, 32, 33, 26, 0,
	5, 22, 23, 8, 7, 9, 0, 0, 0, 24,
	95, 25, 27, 28, 29, 0, 0, 0, 30, 31,
	32, 33, 26, 0, 5, 22, 23, 8, 7, 9,
	0, 0, 0, 24, 93, 25, 27, 28, 29, 0,
	0, 0, 30, 31, 32, 33, 26, 0, 5, 22,
	23, 8, 7, 9, 0, 0, 0, 24, 90, 25,
	27, 28, 29, 0, 0, 0, 30, 31, 32, 33,
	26, 0, 5, 22, 23, 8, 7, 9, 0, 0,
	0, 24, 0, 25, 27, 28, 29, 0, 0, 0,
	30, 31, 32, 33, 26, 0, 5, 22, 23, 8,
	53, 9, 0, 0, 0, 24, 0, 25, 27, 28,
	29, 0, 0, 0, 30, 31, 32, 33, 26,
}
var RubyPact = []int{

	-1000, 191, -1000, -1000, -16, -1000, -1000, 9, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -9, -9, -9, -9, 36, -9, -1000, -1000,
	-9, -9, -9, -9, 215, -1000, -2, -1000, 33, -1000,
	-1000, -9, -1000, -1000, 34, -9, -1000, -9, 215, 215,
	215, 215, -16, -9, -9, -1000, -9, 67, 59, -4,
	24, -10, -16, -16, -16, -16, 11, -1000, 215, -9,
	-1000, -1000, -18, -1000, -9, -9, -1, -1000, -16, 41,
	-1000, 167, -24, -1000, -3, 143, -1000, -9, 119, -1000,
	-1000, -1000, 21, -1000, 63, -1000, 95, -9, -1000, -1000,
	-1000, -1000,
}
var RubyPgo = []int{

	0, 0, 9, 96, 94, 93, 92, 83, 91, 90,
	87, 80, 4, 77, 74, 73, 69, 62, 3, 6,
	56, 53, 47, 43,
}
var RubyR1 = []int{

	0, 20, 20, 18, 18, 5, 7, 7, 7, 7,
	6, 6, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 19, 19, 19,
	19, 21, 21, 21, 21, 21, 9, 10, 10, 11,
	12, 12, 22, 22, 8, 13, 14, 15, 16, 17,
	23, 23, 1, 1, 1, 3, 4,
}
var RubyR2 = []int{

	0, 0, 2, 0, 2, 2, 1, 1, 1, 1,
	2, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 0, 1, 1,
	5, 0, 1, 1, 5, 5, 7, 6, 8, 6,
	3, 6, 1, 4, 5, 3, 3, 3, 3, 3,
	1, 3, 0, 2, 2, 1, 1,
}
var RubyChk = []int{

	-1000, -20, -7, -6, -2, 25, -1, 5, 4, 6,
	-9, -10, -5, -11, -8, -3, -4, -13, -14, -15,
	-16, -17, 26, 27, 10, 12, 23, 13, 14, 15,
	19, 20, 21, 22, -23, 21, -1, -19, -1, 4,
	5, 7, -1, -1, -1, -1, 5, -1, -1, -1,
	-1, -1, -2, 5, 21, -19, 18, -1, 5, -12,
	-1, -12, -2, -2, -2, -2, -1, -1, -1, -21,
	5, 4, -19, 25, 16, 6, -22, 25, -2, -1,
	25, -18, -12, -1, 23, -18, 8, 9, -18, -7,
	11, 25, 23, 11, -1, 11, -18, 6, 4, 5,
	11, -1,
}
var RubyDef = []int{

	1, -2, 2, 6, 7, 8, 9, -2, 12, 14,
	15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
	25, 26, 52, 52, 52, 52, 0, 52, 55, 56,
	52, 52, 52, 52, 0, 50, 0, 10, 27, 28,
	29, 52, 53, 54, 0, 52, 5, 52, 0, 0,
	0, 0, 49, -2, 52, 11, 52, 31, 27, 0,
	0, 0, 45, 46, 47, 48, 0, 51, 0, 52,
	32, 33, 0, 3, 52, -2, 0, 3, 44, 0,
	3, 0, 0, 40, 0, 0, 30, 52, 0, 4,
	37, 3, 0, 39, 0, 36, 0, -2, 34, 35,
	38, 41,
}
var RubyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 27,
	25, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 26,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24,
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
		//line parser.y:94
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:96
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 3:
		//line parser.y:103
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 4:
		//line parser.y:105
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
			}
		}
	case 5:
		//line parser.y:112
		{
			RubyVAL.genericValue = ast.Symbol{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 6:
		//line parser.y:117
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 7:
		//line parser.y:119
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 8:
		//line parser.y:121
		{
			RubyVAL.genericValue = nil /* ignores new lines around statements */
		}
	case 9:
		//line parser.y:123
		{
			RubyVAL.genericValue = nil /* ignores whitespace around statements */
		}
	case 10:
		//line parser.y:126
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue,
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 11:
		//line parser.y:133
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue,
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
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
		//line parser.y:142
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 28:
		//line parser.y:144
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 29:
		//line parser.y:146
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 30:
		//line parser.y:148
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 31:
		//line parser.y:152
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 32:
		//line parser.y:154
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 33:
		//line parser.y:156
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 34:
		//line parser.y:158
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 35:
		//line parser.y:160
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 36:
		//line parser.y:163
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-3].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 37:
		//line parser.y:172
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 38:
		//line parser.y:179
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 39:
		//line parser.y:189
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 40:
		//line parser.y:198
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 41:
		//line parser.y:204
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 42:
		//line parser.y:212
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 43:
		//line parser.y:216
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 44:
		//line parser.y:221
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 45:
		//line parser.y:228
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 46:
		//line parser.y:229
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 47:
		//line parser.y:230
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 48:
		//line parser.y:231
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 49:
		//line parser.y:234
		{
			RubyVAL.genericValue = ast.Addition{
				LHS: RubyS[Rubypt-2].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 52:
		//line parser.y:243
		{
			RubyVAL.genericValue = nil
		}
	case 53:
		//line parser.y:245
		{
			RubyVAL.genericValue = nil
		}
	case 54:
		//line parser.y:247
		{
			RubyVAL.genericValue = nil
		}
	case 55:
		//line parser.y:249
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 56:
		//line parser.y:250
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	}
	goto Rubystack /* stack new state and value */
}
