package command

import (
	"errors"

	"github.com/blushft/strana/domain/entity"
	"github.com/google/uuid"
)

type TrackPageviewCommand interface {
	Track(*entity.RawMessage) error
}

type trackPageviewCmd struct {
	tracker  *tracker
	pagviews entity.PageviewManager
}

func NewTrackPageviewCommand(
	ar entity.AppReporter,
	pv entity.PageviewManager,
	ses entity.SessionManager,
	usr entity.UserManager,
) TrackPageviewCommand {
	return &trackPageviewCmd{
		tracker:  newTracker(ar, ses, usr),
		pagviews: pv,
	}
}

func (cmd *trackPageviewCmd) Track(msg *entity.RawMessage) error {
	if msg.AppID == "" && msg.TrackingID == "" {
		return errors.New("cannot track event: no app or tracking id")
	}

	app, err := msg.GetApp(cmd.tracker.apps)
	if err != nil {
		return err
	}

	usr, err := cmd.tracker.getUserOrNew(msg)
	if err != nil {
		return err
	}

	session, err := cmd.tracker.getSessionOrNew(msg, app, usr)
	if err != nil {
		return err
	}

	pvid, err := uuid.Parse(msg.EventID)
	if err != nil {
		return err
	}

	pv := &entity.Pageview{
		ID:        pvid,
		AppID:     app.ID,
		SessionID: session.ID,
		Hostname:  msg.URL,
		Pathname:  msg.Path,
		IsEntry:   parseBool(msg.NewSession),
		Referrer:  msg.Referrer,
		Timestamp: parseTimestamp(msg.Timestamp),
		UserAgent: msg.UserAgent,
		IPAddress: msg.IPAddress,
	}

	if err := cmd.pagviews.Create(pv); err != nil {
		return err
	}

	return nil
}
