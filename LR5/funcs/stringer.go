package funcs

type Stringer interface {
	Create(name string, author string) Stringer
	GetInfo()
}
