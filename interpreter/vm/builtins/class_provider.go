package builtins

type ClassProvider interface {
	ClassWithName(string) Class
}
