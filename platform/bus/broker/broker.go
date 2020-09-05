package broker

import "github.com/blushft/strana"

type Broker interface {
	Publisher() strana.Publisher
	Subscriber() strana.Subscriber
	Connect() error
	Disconnect() error
}
