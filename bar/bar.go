package bar

import (
	"github.com/mchalski/foo/foo"
)

type Bar struct {
}

func NewBar() *Bar {
	_ = foo.NewFoo()
	return &Bar{}
}

func (b *Bar) UseFoo(f *foo.Foo) error {
	return nil
}
