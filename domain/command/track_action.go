package command

import (
	"github.com/pkg/errors"

	"github.com/blushft/strana/domain/entity"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
)

type TrackActionCommand interface {
	Track(*entity.RawMessage) error
}

type trackActionCmd struct {
	tracker *tracker
}

func NewTrackActionCommand(
	ar entity.AppReporter,
	ses entity.SessionManager,
	usr entity.UserManager,
) TrackActionCommand {
	return &trackActionCmd{
		tracker: newTracker(ar, ses, usr),
	}
}

func (cmd *trackActionCmd) Track(msg *entity.RawMessage) error {
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

	aid, err := uuid.Parse(msg.EventID)
	if err != nil {
		return errors.Wrap(err, "extract event uuid")
	}

	act := &entity.Action{
		ID:             aid,
		AppID:          app.ID,
		SessionID:      session.ID,
		UserID:         usr.ID,
		NonInteractive: parseBool(msg.NonInteractive),
		Category:       msg.Category,
		Action:         msg.Action,
		Label:          msg.Label,
		Property:       msg.Property,
		Value:          msg.Value,
	}

	spew.Dump(act)

	return nil
}
