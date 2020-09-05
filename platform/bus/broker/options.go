package broker

type Options struct {
	Embedded bool
	Web      bool
}

type Option func(*Options)

func NewOptions(opts ...Option) Options {
	options := Options{
		Embedded: true,
		Web:      true,
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}
