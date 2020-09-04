package bus

type Options struct {
	Embedded bool
}

type Option func(*Options)

func NewOptions(opts ...Option) Options {
	options := Options{
		Embedded: true,
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

func Embedded(b bool) Option {
	return func(o *Options) {
		o.Embedded = b
	}
}
