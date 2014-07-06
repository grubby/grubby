//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:3
var Statements []interface{}

//line parser.y:11
type RubySymType struct {
	yys      int
	intval   int
	floatval float64
}

const DIGIT = 57346
const UMINUS = 57347

var RubyToknames = []string{
	"DIGIT",
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

//line parser.y:66

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const RubyNprod = 15
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 52

var RubyAct = []int{

	3, 11, 12, 13, 8, 16, 17, 2, 1, 6,
	18, 19, 20, 21, 22, 23, 24, 15, 14, 9,
	10, 11, 12, 13, 0, 0, 0, 25, 15, 14,
	9, 10, 11, 12, 13, 14, 9, 10, 11, 12,
	13, 7, 0, 0, 0, 5, 9, 10, 11, 12,
	13, 4,
}
var RubyPact = []int{

	-1000, 37, -9, 23, 37, 37, -1000, -1000, -1000, 37,
	37, 37, 37, 37, 37, 37, 12, -1000, -8, -8,
	-1000, -1000, -1000, 39, 29, -1000,
}
var RubyPgo = []int{

	0, 0, 9, 8, 7,
}
var RubyR1 = []int{

	0, 3, 3, 4, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2,
}
var RubyR2 = []int{

	0, 0, 3, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 1, 1,
}
var RubyChk = []int{

	-1000, -3, -4, -1, 14, 8, -2, 4, 13, 7,
	8, 9, 10, 11, 6, 5, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, 15,
}
var RubyDef = []int{

	1, -2, 0, 3, 0, 0, 13, 14, 2, 0,
	0, 0, 0, 0, 0, 0, 0, 12, 5, 6,
	7, 8, 9, 10, 11, 4,
}
var RubyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	13, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 11, 6, 3,
	14, 15, 9, 7, 3, 8, 3, 10, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 5,
}
var RubyTok2 = []int{

	2, 3, 4, 12,
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
		//line parser.y:37
		{
			Statements = append(Statements, RubyS[Rubypt-0].intval)
		}
	case 4:
		//line parser.y:42
		{
			RubyVAL.intval = RubyS[Rubypt-1].intval
		}
	case 5:
		//line parser.y:44
		{
			RubyVAL.intval = RubyS[Rubypt-2].intval + RubyS[Rubypt-0].intval
		}
	case 6:
		//line parser.y:46
		{
			RubyVAL.intval = RubyS[Rubypt-2].intval - RubyS[Rubypt-0].intval
		}
	case 7:
		//line parser.y:48
		{
			RubyVAL.intval = RubyS[Rubypt-2].intval * RubyS[Rubypt-0].intval
		}
	case 8:
		//line parser.y:50
		{
			RubyVAL.intval = RubyS[Rubypt-2].intval / RubyS[Rubypt-0].intval
		}
	case 9:
		//line parser.y:52
		{
			RubyVAL.intval = RubyS[Rubypt-2].intval % RubyS[Rubypt-0].intval
		}
	case 10:
		//line parser.y:54
		{
			RubyVAL.intval = RubyS[Rubypt-2].intval & RubyS[Rubypt-0].intval
		}
	case 11:
		//line parser.y:56
		{
			RubyVAL.intval = RubyS[Rubypt-2].intval | RubyS[Rubypt-0].intval
		}
	case 12:
		//line parser.y:58
		{
			RubyVAL.intval = -RubyS[Rubypt-0].intval
		}
	case 13:
		RubyVAL.intval = RubyS[Rubypt-0].intval
	case 14:
		//line parser.y:63
		{
			RubyVAL.intval = RubyS[Rubypt-0].intval
		}
	}
	goto Rubystack /* stack new state and value */
}
