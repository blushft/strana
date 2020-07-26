package config

type Enhancer struct {
	Subscriber Sink
	Publisher  Source
}

func DefaultEnhancerConfig() Enhancer {
	return Enhancer{
		Subscriber: Sink{
			Source: "collector",
		},
		Publisher: Source{
			Module: "enhancer",
			Topic:  "enhanced_raw_message",
			Broker: "in_process",
		},
	}
}
