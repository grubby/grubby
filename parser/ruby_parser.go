//line parser.y:8
package parser

import __yyfmt__ "fmt"

//line parser.y:9
import (
	"fmt"
)

var regs = make([]int, 26)
var base int

//line parser.y:22
type RubySymType struct {
	yys int
	val int
}

const DIGIT = 57346
const LETTER = 57347
const UMINUS = 57348

var RubyToknames = []string{
	"DIGIT",
	"LETTER",
	" |",
	" &",
	" +",
	" -",
	" *",
	" /",
	" %",
	"UMINUS",
}
var RubyStatenames = []string{}

const RubyEofCode = 1
const RubyErrCode = 2
const RubyMaxDepth = 200

//line parser.y:91

/*  start  of  programs  */

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const RubyNprod = 18
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 54

var RubyAct = []int{

	3, 10, 11, 12, 13, 14, 18, 20, 21, 17,
	9, 22, 23, 24, 25, 26, 27, 28, 29, 16,
	15, 10, 11, 12, 13, 14, 8, 19, 8, 4,
	30, 6, 2, 6, 1, 12, 13, 14, 5, 7,
	5, 16, 15, 10, 11, 12, 13, 14, 15, 10,
	11, 12, 13, 14,
}
var RubyPact = []int{

	-1000, 24, -4, 35, -6, 22, 22, 4, -1000, -1000,
	22, 22, 22, 22, 22, 22, 22, 22, 13, -1000,
	-1000, -1000, 25, 25, -1000, -1000, -1000, -7, 41, 35,
	-1000,
}
var RubyPgo = []int{

	0, 0, 39, 34, 32,
}
var RubyR1 = []int{

	0, 3, 3, 4, 4, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 2,
}
var RubyR2 = []int{

	0, 0, 3, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 1, 1, 1, 2,
}
var RubyChk = []int{

	-1000, -3, -4, -1, 5, 16, 9, -2, 4, 14,
	8, 9, 10, 11, 12, 7, 6, 15, -1, 5,
	-1, 4, -1, -1, -1, -1, -1, -1, -1, -1,
	17,
}
var RubyDef = []int{

	1, -2, 0, 3, 14, 0, 0, 15, 16, 2,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 14,
	13, 17, 6, 7, 8, 9, 10, 11, 12, 4,
	5,
}
var RubyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	14, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 12, 7, 3,
	16, 17, 10, 8, 3, 9, 3, 11, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 15, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 6,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 13,
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

	case 3:
		//line parser.y:46
		{
			fmt.Printf("you typed '%d'\n", RubyS[Rubypt-0].val)
		}
	case 4:
		//line parser.y:50
		{
			regs[RubyS[Rubypt-2].val] = RubyS[Rubypt-0].val
		}
	case 5:
		//line parser.y:56
		{
			RubyVAL.val = RubyS[Rubypt-1].val
		}
	case 6:
		//line parser.y:58
		{
			RubyVAL.val = RubyS[Rubypt-2].val + RubyS[Rubypt-0].val
		}
	case 7:
		//line parser.y:60
		{
			RubyVAL.val = RubyS[Rubypt-2].val - RubyS[Rubypt-0].val
		}
	case 8:
		//line parser.y:62
		{
			RubyVAL.val = RubyS[Rubypt-2].val * RubyS[Rubypt-0].val
		}
	case 9:
		//line parser.y:64
		{
			RubyVAL.val = RubyS[Rubypt-2].val / RubyS[Rubypt-0].val
		}
	case 10:
		//line parser.y:66
		{
			RubyVAL.val = RubyS[Rubypt-2].val % RubyS[Rubypt-0].val
		}
	case 11:
		//line parser.y:68
		{
			RubyVAL.val = RubyS[Rubypt-2].val & RubyS[Rubypt-0].val
		}
	case 12:
		//line parser.y:70
		{
			RubyVAL.val = RubyS[Rubypt-2].val | RubyS[Rubypt-0].val
		}
	case 13:
		//line parser.y:72
		{
			RubyVAL.val = -RubyS[Rubypt-0].val
		}
	case 14:
		//line parser.y:74
		{
			RubyVAL.val = regs[RubyS[Rubypt-0].val]
		}
	case 15:
		RubyVAL.val = RubyS[Rubypt-0].val
	case 16:
		//line parser.y:79
		{
			RubyVAL.val = RubyS[Rubypt-0].val
			if RubyS[Rubypt-0].val == 0 {
				base = 8
			} else {
				base = 10
			}
		}
	case 17:
		//line parser.y:88
		{
			RubyVAL.val = base*RubyS[Rubypt-1].val + RubyS[Rubypt-0].val
		}
	}
	goto Rubystack /* stack new state and value */
}
