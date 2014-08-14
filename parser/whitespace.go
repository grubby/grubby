package parser

const whitespace = " \t"

func lexWhitespace(l *StatefulRubyLexer) stateFn {
	l.acceptRun(whitespace)
	l.emit(tokenTypeWhitespace)
	return lexAnything
}

const newline = "\n"

func lexNewlines(l *StatefulRubyLexer) stateFn {
	for l.accept(newline) {
		l.emit(tokenTypeNewline)
	}
	return lexAnything
}
