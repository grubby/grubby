package parser

func lexReference(l *BetterRubyLexer) stateFn {
	l.acceptRun(alphaNumericUnderscore)
	l.emit(tokenTypeReference)
	return lexAnything
}
