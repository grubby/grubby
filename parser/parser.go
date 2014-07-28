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
const COLON = 57361
const DOT = 57362

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
	"COLON",
	"DOT",
}
var RubyStatenames = []string{}

const RubyEofCode = 1
const RubyErrCode = 2
const RubyMaxDepth = 200

//line parser.y:216

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 39,
	19, 37,
	-2, 35,
	-1, 65,
	19, 38,
	-2, 36,
}

const RubyNprod = 42
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 140

var RubyAct = []int{

	54, 38, 4, 23, 8, 7, 9, 46, 47, 64,
	17, 71, 18, 20, 21, 22, 58, 52, 42, 19,
	31, 5, 6, 29, 25, 26, 28, 27, 32, 8,
	7, 9, 69, 41, 53, 17, 67, 18, 20, 21,
	22, 45, 24, 49, 19, 50, 5, 6, 25, 26,
	57, 27, 56, 48, 61, 33, 55, 43, 44, 39,
	25, 26, 33, 27, 65, 37, 68, 8, 7, 9,
	70, 59, 60, 17, 66, 18, 20, 21, 22, 36,
	35, 30, 19, 40, 5, 6, 8, 7, 9, 62,
	34, 2, 17, 63, 18, 20, 21, 22, 1, 13,
	11, 19, 10, 5, 6, 8, 7, 9, 14, 3,
	12, 17, 16, 18, 20, 21, 22, 15, 0, 0,
	19, 0, 5, 6, 8, 51, 9, 0, 0, 0,
	17, 0, 18, 20, 21, 22, 0, 0, 0, 19,
}
var RubyPact = []int{

	-1000, 101, -1000, -1000, -1000, -1000, -1000, 20, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 4, 1, 76,
	-2, -1000, -1000, -1000, 44, -1000, -1000, 75, 60, 53,
	-1000, 53, -1000, -4, 49, -1000, -1000, 56, -14, -1000,
	34, 22, 120, -1000, -5, 13, -1000, 40, 33, -1000,
	-1000, -6, 67, -1000, 82, -13, 58, 63, 37, -1000,
	-1000, 25, -1000, -1000, 53, -1000, -1000, -1000, 11, -1000,
	0, -1000,
}
var RubyPgo = []int{

	0, 2, 117, 112, 110, 109, 89, 108, 102, 100,
	99, 1, 0, 3, 98, 90, 83,
}
var RubyR1 = []int{

	0, 14, 14, 12, 12, 4, 6, 6, 6, 6,
	5, 5, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 13, 13, 13, 13, 15, 15, 15, 15,
	15, 8, 9, 9, 10, 11, 11, 16, 16, 7,
	2, 3,
}
var RubyR2 = []int{

	0, 0, 2, 0, 2, 2, 1, 1, 1, 1,
	2, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 0, 1, 1, 3, 0, 1, 1, 4,
	4, 7, 6, 10, 6, 1, 4, 1, 4, 5,
	1, 1,
}
var RubyChk = []int{

	-1000, -14, -6, -5, -1, 21, 22, 5, 4, 6,
	-8, -9, -4, -10, -7, -2, -3, 10, 12, 19,
	13, 14, 15, -13, 22, 4, 5, 7, 22, 22,
	5, 22, -13, 18, -15, 5, 4, 5, -11, 6,
	-16, -11, 22, 8, 9, -13, 21, 22, 19, 21,
	-1, 5, 22, 21, -12, 16, 19, -12, 22, 4,
	5, -12, -6, 11, 22, 6, 11, 11, -11, 21,
	-12, 11,
}
var RubyDef = []int{

	1, -2, 2, 6, 7, 8, 9, 13, 12, 14,
	15, 16, 17, 18, 19, 20, 21, 0, 0, 0,
	0, 40, 41, 10, 22, 23, 24, 26, 0, 0,
	5, 0, 11, 0, 0, 27, 28, 22, 0, -2,
	0, 0, 0, 25, 0, 0, 3, 0, 0, 3,
	39, 13, 0, 3, 0, 0, 0, 0, 0, 29,
	30, 0, 4, 32, 0, -2, 34, 31, 0, 3,
	0, 33,
}
var RubyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	21, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 22,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20,
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
		//line parser.y:79
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:81
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 3:
		//line parser.y:88
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 4:
		//line parser.y:90
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
			}
		}
	case 5:
		//line parser.y:97
		{
			RubyVAL.genericValue = ast.Symbol{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 6:
		//line parser.y:102
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 7:
		//line parser.y:104
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 8:
		//line parser.y:106
		{
			RubyVAL.genericValue = nil /* ignores new lines around statements */
		}
	case 9:
		//line parser.y:108
		{
			RubyVAL.genericValue = nil /* ignores whitespace around statements */
		}
	case 10:
		//line parser.y:111
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue,
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 11:
		//line parser.y:118
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
		//line parser.y:127
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 23:
		//line parser.y:129
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 24:
		//line parser.y:131
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 25:
		//line parser.y:133
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 26:
		//line parser.y:137
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 27:
		//line parser.y:139
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 28:
		//line parser.y:141
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 29:
		//line parser.y:143
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 30:
		//line parser.y:145
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 31:
		//line parser.y:148
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-3].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 32:
		//line parser.y:157
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 33:
		//line parser.y:164
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-7].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-7].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 34:
		//line parser.y:174
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 35:
		//line parser.y:183
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
			}
		}
	case 36:
		//line parser.y:189
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-0].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-3].stringSlice, "::"),
			}
		}
	case 37:
		//line parser.y:197
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 38:
		//line parser.y:201
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 39:
		//line parser.y:206
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 40:
		//line parser.y:213
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 41:
		//line parser.y:214
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	}
	goto Rubystack /* stack new state and value */
}
