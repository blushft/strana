package enhancer

import (
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/blushft/strana/domain/entity"
	"github.com/blushft/strana/platform"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/store"
	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber"
)

type Enhancer interface {
	platform.Module
	platform.Consumer
	platform.Producer
}

type enhancer struct {
	conf config.Enhancer
	pub  message.Publisher
	sub  message.Subscriber
}

func NewEnhancer(conf config.Enhancer) Enhancer {
	return &enhancer{
		conf: conf,
	}
}

func (e *enhancer) Routes(fiber.Router) {}

func (e *enhancer) Services(*store.Store) {}

func (e *enhancer) Events(eh platform.EventHandler) error {
	topic, sub, err := eh.Source(e.conf.Subscriber.Source)
	if err != nil {
		return err
	}

	pb, err := eh.Broker(e.conf.Publisher.Broker)
	if err != nil {
		return err
	}

	e.sub = sub.Subscriber()
	e.pub = pb.Publisher()

	eh.Register(e.conf.Publisher, e)

	eh.Router().AddHandler(
		"enhancer",
		topic,
		e.sub,
		e.conf.Publisher.Topic,
		e.pub,
		e.handle,
	)

	return nil
}

func (e *enhancer) Publisher() message.Publisher {
	return e.pub
}

func (e *enhancer) Subscriber() message.Subscriber {
	return e.sub
}

func (e *enhancer) handle(msg *message.Message) ([]*message.Message, error) {

	if msg.Payload == nil {
		log.Println("message is nil bro")
		return nil, nil
	}

	var rm *entity.RawMessage
	if err := json.Unmarshal(msg.Payload, &rm); err != nil {
		panic(err)
		//return nil, err
	}

	log.Println("enhancer got an event")
	spew.Dump(rm)

	return []*message.Message{msg}, nil
}
