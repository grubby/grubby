package parser

const whitespace = " \t"
const newline = "\n"

func lexNewlines(l *StatefulRubyLexer) stateFn {
	for l.accept(newline) {
		l.emit(tokenTypeNewline)
	}
	return lexAnything
}
