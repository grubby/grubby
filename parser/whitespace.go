package parser

const whitespace = " \t"
const newline = "\n"

func lexNewlines(l StatefulRubyLexer) stateFn {
	for l.accept(newline) {
		l.parsedNewLine()
		l.emit(tokenTypeNewline)
	}
	return lexSomething
}

func lexWhiteSpaceIncludingNewlineAndComments(l StatefulRubyLexer) stateFn {
	// accept any whitespace we might currently have at the end of the line
	l.acceptRun(whitespace)

	// accept a newline and following whitespace, if present
	if l.accept(newline) {
		l.parsedNewLine()
	}
	l.acceptRun(whitespace)

	// ignore a comment, if this line had one
	// as of MRI ruby 2.1.1 only a single newline + comment is allowed
	if l.accept("#") {
		for r := l.next(); r != '\n' && r != eof; r = l.next() {
		}

		l.backup()
		l.ignore()

		if l.accept(newline) {
			l.parsedNewLine()
			l.acceptRun(whitespace)
		}
	}

	l.ignore()
	return lexSomething
}
