package webhook

import (
	"net/http"

	"github.com/blushft/strana"
	"github.com/blushft/strana/event"
	"github.com/blushft/strana/event/contexts"
	"github.com/blushft/strana/event/events"
	"github.com/blushft/strana/platform"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/logger"
	"github.com/blushft/strana/platform/store"
	"github.com/blushft/strana/processor"
	"github.com/gofiber/fiber"
	"github.com/mitchellh/mapstructure"
)

type pingdomHook struct {
	conf  config.Module
	opts  pingdomOptions
	log   *logger.Logger
	sub   strana.Subscriber
	pub   strana.Publisher
	procs []processor.EventProcessor
}

type pingdomOptions struct {
	Path       string `json:"path" structs:"path" mapstructure:"path"`
	TrackingID string `json:"trackingID" structs:"trackingID" mapstructure:"trackingID"`
}

func defaultPingdomOptions() *pingdomOptions {
	return &pingdomOptions{
		Path: "/webhooks/pingdom",
	}
}

func newPingdomHook(conf config.Module, opts HookOptions, procs ...config.Processor) (strana.Module, error) {
	var modprocs []processor.EventProcessor
	var err error

	if len(procs) > 0 {
		modprocs, err = platform.NewEventProcessorSet(procs)
		if err != nil {
			return nil, err
		}
	}

	modopts := defaultPingdomOptions()
	if err := mapstructure.Decode(opts, &modopts); err != nil {
		return nil, err
	}

	return &pingdomHook{
		conf:  conf,
		opts:  *modopts,
		procs: modprocs,
	}, nil
}

func (mod *pingdomHook) Routes(rtr fiber.Router) error {
	rtr.Post(mod.opts.Path, mod.collect)

	return nil
}

func (mod *pingdomHook) Events(eh strana.EventHandler) error {
	mod.pub = eh.Publisher()
	mod.sub = eh.Subscriber()

	return nil
}

func (mod *pingdomHook) Services(s *store.SQLStore) error {
	return nil
}

func (mod *pingdomHook) Logger(l *logger.Logger) {
	mod.log = l.WithFields(logger.Fields{
		"module": "pingdom_webhook",
	})
}

func (mod *pingdomHook) Subscribe(fn strana.SubscriptionHandlerFunc) error {
	return mod.sub.Subscribe(mod.conf.Sink, fn)
}

func (mod *pingdomHook) collect(ctx *fiber.Ctx) {
	var pl *pingdomPayload
	if err := ctx.BodyParser(&pl); err != nil {
		mod.log.WithError(err).Error("extract pingdom payload")
		ctx.SendStatus(http.StatusBadRequest)

		return
	}

	go mod.extract(pl)

	ctx.SendStatus(http.StatusOK)
}

func (mod *pingdomHook) extract(pl *pingdomPayload) {
	evtctx := []event.Context{
		&contexts.Action{
			Category: pl.CheckType,
			Action:   "pingdom_check",
			Label:    pl.CheckName,
			Property: "pingdom_state",
			Value:    pl.CurrentState,
		},
		contexts.NewNetwork(pl.FirstProbe.IP),
		&contexts.Page{
			Hostname: pl.CheckParams.Hostname,
			Path:     pl.CheckParams.URL,
		},
	}

	evt := event.New(
		events.EventTypeAction,
		event.TrackingID(mod.opts.TrackingID),
		event.Channel("webhook"),
		event.Platform("pingdom"),
		event.WithContexts(evtctx...),
		event.NonInteractive(),
	)

	mod.publish(evt)
}

func (mod *pingdomHook) publish(evt *event.Event) {
	evts, err := processor.Execute(mod.procs, evt)
	if err != nil {
		mod.log.WithError(err).Error("process event")
	}

	for _, re := range evts {
		m := message.NewMessage(re)
		if err := mod.pub.Publish(mod.conf.Sink, m); err != nil {
			mod.log.WithError(err).Error("publish event")
		}
	}
}

type pingdomPayload struct {
	CheckID               int64              `json:"check_id"`
	CheckName             string             `json:"check_name"`
	CheckType             string             `json:"check_type"`
	CheckParams           pingdomCheckParams `json:"check_params"`
	Tags                  []string           `json:"tags"`
	PreviousState         string             `json:"previous_state"`
	CurrentState          string             `json:"current_state"`
	ImportanceLevel       string             `json:"importance_level"`
	StateChangedTimestamp int64              `json:"state_changed_timestamp"`
	StateChangedUTCTime   string             `json:"state_changed_utc_time"`
	LongDescription       string             `json:"long_description"`
	Description           string             `json:"description"`
	FirstProbe            pingdomProbe       `json:"first_probe"`
	SecondProbe           pingdomProbe       `json:"second_probe"`
}

type pingdomCheckParams struct {
	FullURL    string `json:"full_url"`
	Header     string `json:"header"`
	Hostname   string `json:"hostname"`
	URL        string `json:"url"`
	Port       int64  `json:"port"`
	BasicAuth  bool   `json:"basic_auth"`
	Encryption bool   `json:"encryption"`
	Ipv6       bool   `json:"ipv6"`
}

type pingdomProbe struct {
	IP       string `json:"ip"`
	Ipv6     string `json:"ipv6"`
	Location string `json:"location"`
	Version  *int64 `json:"version,omitempty"`
}
