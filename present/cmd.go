package present

import "github.com/jpillora/opts"

type Command struct {
	Foo int
}

func New() opts.Opts {
	return opts.New(&Command{
		Foo: 42,
	})
}

func (c *Command) Run() error {
	return nil
}
