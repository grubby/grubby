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

//line parser.y:263

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 4,
	23, 53,
	-2, 7,
	-1, 7,
	7, 53,
	18, 53,
	-2, 12,
	-1, 61,
	18, 53,
	-2, 12,
	-1, 73,
	24, 43,
	-2, 53,
	-1, 101,
	24, 44,
	-2, 53,
}

const RubyNprod = 58
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 264

var RubyAct = []int{

	36, 51, 6, 85, 24, 25, 72, 94, 37, 84,
	57, 75, 48, 49, 50, 52, 71, 95, 88, 97,
	98, 53, 54, 56, 55, 38, 39, 40, 41, 101,
	43, 105, 106, 44, 45, 46, 47, 73, 53, 54,
	42, 55, 58, 74, 58, 83, 82, 81, 4, 65,
	66, 67, 52, 68, 59, 92, 69, 2, 70, 1,
	23, 22, 76, 21, 20, 19, 18, 17, 13, 11,
	10, 14, 3, 58, 87, 12, 16, 15, 0, 89,
	0, 0, 90, 86, 0, 0, 6, 0, 91, 0,
	6, 0, 6, 60, 62, 63, 64, 0, 100, 102,
	0, 6, 104, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 77, 78, 79, 80, 8, 7,
	9, 0, 0, 0, 26, 103, 27, 29, 30, 31,
	0, 0, 0, 32, 33, 34, 35, 0, 28, 0,
	5, 24, 25, 8, 7, 9, 0, 0, 0, 26,
	99, 27, 29, 30, 31, 0, 0, 0, 32, 33,
	34, 35, 0, 28, 0, 5, 24, 25, 8, 7,
	9, 0, 0, 0, 26, 96, 27, 29, 30, 31,
	0, 0, 0, 32, 33, 34, 35, 0, 28, 0,
	5, 24, 25, 8, 7, 9, 0, 0, 0, 26,
	93, 27, 29, 30, 31, 0, 0, 0, 32, 33,
	34, 35, 0, 28, 0, 5, 24, 25, 8, 7,
	9, 0, 0, 0, 26, 0, 27, 29, 30, 31,
	0, 0, 0, 32, 33, 34, 35, 0, 28, 0,
	5, 24, 25, 8, 61, 9, 0, 0, 0, 26,
	0, 27, 29, 30, 31, 0, 0, 0, 32, 33,
	34, 35, 0, 28,
}
var RubyPact = []int{

	-1000, 214, -1000, -1000, -23, -1000, -1000, -23, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -23, -23, -23, -23, 35, -23,
	-1000, -1000, -23, -23, -23, -23, -9, 34, -1000, -1000,
	18, -23, -1000, -23, 239, 239, 239, 239, -23, -23,
	-23, -1000, -23, -1000, -1000, -23, 17, -10, 31, -15,
	-23, -23, -23, -23, -23, 239, 239, 239, 239, 41,
	-17, -1000, -23, -23, -6, -1000, -3, -23, -23, -23,
	-23, -23, -1000, -1000, -1000, 189, -19, -1000, -7, 164,
	11, 139, -1000, -1000, -1000, 23, -1000, -1000, -23, -1000,
	114, -23, 27, -1000, -1000, -1000, -1000,
}
var RubyPgo = []int{

	0, 0, 48, 77, 76, 75, 72, 55, 71, 70,
	69, 68, 10, 67, 66, 65, 64, 63, 61, 60,
	3, 1, 59, 47, 43,
}
var RubyR1 = []int{

	0, 22, 22, 20, 20, 5, 7, 7, 7, 7,
	6, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 21, 21,
	21, 21, 23, 23, 23, 23, 23, 9, 10, 10,
	11, 12, 12, 24, 24, 8, 13, 14, 15, 16,
	17, 18, 19, 1, 1, 1, 3, 4,
}
var RubyR2 = []int{

	0, 0, 2, 0, 2, 2, 1, 1, 1, 1,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 0, 1,
	1, 5, 0, 1, 1, 5, 5, 7, 6, 8,
	6, 3, 6, 1, 4, 5, 3, 3, 3, 3,
	5, 5, 5, 0, 2, 2, 1, 1,
}
var RubyChk = []int{

	-1000, -22, -7, -6, -2, 26, -1, 5, 4, 6,
	-9, -10, -5, -11, -8, -3, -4, -13, -14, -15,
	-16, -17, -18, -19, 27, 28, 10, 12, 24, 13,
	14, 15, 19, 20, 21, 22, -1, -1, -1, -1,
	-1, -1, 5, -1, -1, -1, -1, -1, 21, 22,
	23, -21, 18, 4, 5, 7, 5, -12, -1, -12,
	-2, 5, -2, -2, -2, -1, -1, -1, -1, -1,
	-21, 26, 16, 6, -24, 26, -1, -2, -2, -2,
	-2, -23, 5, 4, 26, -20, -12, -1, 24, -20,
	-1, -20, -7, 11, 26, 24, 11, 8, 9, 11,
	-20, 6, -1, 11, -1, 4, 5,
}
var RubyDef = []int{

	1, -2, 2, 6, -2, 8, 9, -2, 11, 13,
	14, 15, 16, 17, 18, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 53, 53, 53, 53, 0, 53,
	56, 57, 53, 53, 53, 53, 0, 28, 54, 55,
	0, 53, 5, 53, 0, 0, 0, 0, 53, 53,
	53, 10, 53, 29, 30, 53, 28, 0, 0, 0,
	46, -2, 47, 48, 49, 0, 0, 0, 0, 32,
	0, 3, 53, -2, 0, 3, 0, 50, 51, 52,
	45, 53, 33, 34, 3, 0, 0, 41, 0, 0,
	0, 0, 4, 38, 3, 0, 40, 31, 53, 37,
	0, -2, 0, 39, 42, 35, 36,
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
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 7:
		//line parser.y:122
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
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
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 33:
		//line parser.y:150
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 34:
		//line parser.y:152
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 35:
		//line parser.y:154
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 36:
		//line parser.y:156
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 37:
		//line parser.y:159
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-3].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 38:
		//line parser.y:168
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 39:
		//line parser.y:175
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 40:
		//line parser.y:185
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 41:
		//line parser.y:194
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 42:
		//line parser.y:200
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 43:
		//line parser.y:208
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 44:
		//line parser.y:212
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 45:
		//line parser.y:217
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 46:
		//line parser.y:224
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 47:
		//line parser.y:225
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 48:
		//line parser.y:226
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 49:
		//line parser.y:227
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 50:
		//line parser.y:230
		{
			RubyVAL.genericValue = ast.Addition{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 51:
		//line parser.y:238
		{
			RubyVAL.genericValue = ast.Subtraction{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 52:
		//line parser.y:246
		{
			RubyVAL.genericValue = ast.Multiplication{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 53:
		//line parser.y:254
		{
			RubyVAL.genericValue = nil
		}
	case 54:
		//line parser.y:256
		{
			RubyVAL.genericValue = nil
		}
	case 55:
		//line parser.y:258
		{
			RubyVAL.genericValue = nil
		}
	case 56:
		//line parser.y:260
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 57:
		//line parser.y:261
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	}
	goto Rubystack /* stack new state and value */
}
