//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:3
import (
	"github.com/grubby/grubby/ast"
)

var Statements []ast.Node

//line parser.y:15
type RubySymType struct {
	yys          int
	genericValue ast.Node
	genericSlice ast.Nodes
}

const NODE = 57346
const REF = 57347
const LPAREN = 57348
const RPAREN = 57349
const COMMA = 57350
const DEF = 57351
const END = 57352
const CLASS = 57353
const MODULE = 57354
const LESSTHAN = 57355
const GREATERTHAN = 57356
const EQUALTO = 57357

var RubyToknames = []string{
	"NODE",
	"REF",
	"LPAREN",
	"RPAREN",
	"COMMA",
	"DEF",
	"END",
	"CLASS",
	"MODULE",
	"LESSTHAN",
	"GREATERTHAN",
	"EQUALTO",
}
var RubyStatenames = []string{}

const RubyEofCode = 1
const RubyErrCode = 2
const RubyMaxDepth = 200

//line parser.y:149

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const RubyNprod = 27
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 66

var RubyAct = []int{

	33, 13, 8, 7, 40, 8, 7, 11, 45, 12,
	11, 41, 12, 31, 5, 6, 20, 5, 6, 15,
	16, 17, 8, 7, 29, 30, 28, 11, 39, 12,
	19, 18, 14, 37, 5, 6, 8, 7, 43, 32,
	34, 11, 42, 12, 44, 26, 27, 25, 5, 6,
	15, 16, 17, 35, 36, 23, 22, 24, 38, 21,
	2, 1, 10, 9, 3, 4,
}
var RubyPact = []int{

	-1000, 32, -1000, -1000, -1000, -1000, -1000, 15, -1000, -1000,
	-1000, 14, 13, -1000, 46, -1000, -1000, 51, 52, 42,
	-1000, 38, -1000, -1000, 46, 8, -1000, -4, 23, -1000,
	27, 49, -1000, 18, -13, -1000, -1000, 1, -1000, -1000,
	37, -1000, 22, -1000, -2, -1000,
}
var RubyPgo = []int{

	0, 65, 64, 58, 63, 62, 0, 1, 61, 59,
}
var RubyR1 = []int{

	0, 8, 8, 6, 6, 3, 3, 3, 3, 2,
	2, 1, 1, 1, 1, 7, 7, 7, 7, 9,
	9, 9, 9, 9, 4, 5, 5,
}
var RubyR2 = []int{

	0, 0, 2, 0, 2, 1, 1, 1, 1, 2,
	3, 1, 1, 1, 1, 0, 1, 1, 3, 0,
	1, 1, 4, 4, 7, 6, 10,
}
var RubyChk = []int{

	-1000, -8, -3, -2, -1, 16, 17, 5, 4, -4,
	-5, 9, 11, -7, 17, 4, 5, 6, 17, 17,
	-7, -9, 5, 4, 5, 5, 7, 8, -7, 16,
	17, 17, 16, -6, 13, 4, 5, -6, -3, 10,
	17, 10, 5, 16, -6, 10,
}
var RubyDef = []int{

	1, -2, 2, 5, 6, 7, 8, 12, 11, 13,
	14, 0, 0, 9, 15, 16, 17, 19, 0, 0,
	10, 0, 20, 21, 15, 0, 18, 0, 0, 3,
	0, 0, 3, 0, 0, 22, 23, 0, 4, 25,
	0, 24, 0, 3, 0, 26,
}
var RubyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	16, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 17,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15,
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
		//line parser.y:61
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:63
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 3:
		//line parser.y:70
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 4:
		//line parser.y:72
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
			}
		}
	case 5:
		//line parser.y:79
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 6:
		//line parser.y:81
		{
			RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
		}
	case 7:
		//line parser.y:83
		{
			RubyVAL.genericValue = nil /* ignores new lines around statements */
		}
	case 8:
		//line parser.y:85
		{
			RubyVAL.genericValue = nil /* ignores whitespace around statements */
		}
	case 9:
		//line parser.y:88
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-1].genericValue,
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 10:
		//line parser.y:95
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
		//line parser.y:104
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 16:
		//line parser.y:106
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 17:
		//line parser.y:108
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 18:
		//line parser.y:110
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 19:
		//line parser.y:114
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 20:
		//line parser.y:116
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 21:
		//line parser.y:118
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 22:
		//line parser.y:120
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 23:
		//line parser.y:122
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 24:
		//line parser.y:125
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-4].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-3].genericSlice,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 25:
		//line parser.y:134
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 26:
		//line parser.y:141
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-7].genericValue.(ast.BareReference),
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.BareReference),
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	}
	goto Rubystack /* stack new state and value */
}
