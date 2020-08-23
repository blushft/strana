package main

import (
	"log"
	"time"

	"github.com/blushft/strana/event"
	"github.com/blushft/strana/platform/logger"
	"github.com/blushft/strana/tracker"
	"github.com/google/uuid"
)

func main() {
	t, err := tracker.New(
		tracker.TrackingID("1234"),
		tracker.SetAppInfo(
			&event.App{
				Name:    "Basic Example",
				Version: "v0.0.1",
			},
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	user := &event.User{
		UserID: uuid.New().String(),
		Name:   "fred",
		Traits: event.Traits{
			Email: "fred@fakemail.int",
		},
	}

	logger.Log().Info("identifying user")
	_ = t.Identify(user)

	logger.Log().Info("submitting action")
	_ = t.Action(&event.Action{
		Category: "documents",
		Action:   "open_document",
		Label:    "doc",
		Property: "file_name",
		Value:    "test_doc.txt",
	})

	logger.Log().Info("start timing tracking")
	timer := t.TimingStart("downloads", "product_install", "download")

	time.Sleep(time.Millisecond * 5022)

	timing := timer.End()

	logger.Log().Info("submit timing")
	_ = t.Timing(timing, event.Channel("my_product"))

	time.Sleep(time.Second)

	logger.Log().Info("all done")
}
