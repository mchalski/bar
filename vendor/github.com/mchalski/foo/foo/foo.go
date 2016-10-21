package foo

//comment
//comment
//comment

type Foo struct {
}

func NewFoo() *Foo {
	return &Foo{}
}

func (f *Foo) FooMethod() error {
	return nil
}
