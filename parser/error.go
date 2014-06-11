package parser

type UnexpectedEndError struct{}

func NewUnexpectedEnd() error {
	return &UnexpectedEndError{}
}

func (err *UnexpectedEndError) Error() string {
	return "Read an 'end' where none was expected"
}
