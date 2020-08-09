package command

import (
	"errors"
	"strconv"
	"time"

	"github.com/blushft/strana/domain/entity"
	"github.com/blushft/strana/platform/store/ent"
	"github.com/google/uuid"
)

type TrackPageviewCommand interface {
	Track(*entity.RawMessage) error
}

type trackPageviewCmd struct {
	apps     entity.AppReporter
	sessions entity.SessionManager
	pagviews entity.PageviewManager
	users    entity.UserManager
}

func NewTrackPageviewCommand(
	ar entity.AppReporter,
	pv entity.PageviewManager,
	ses entity.SessionManager,
	usr entity.UserManager,
) TrackPageviewCommand {
	return &trackPageviewCmd{
		apps:     ar,
		sessions: ses,
		pagviews: pv,
		users:    usr,
	}
}

func (cmd *trackPageviewCmd) Track(msg *entity.RawMessage) error {
	if msg.AppID == "" && msg.TrackingID == "" {
		return errors.New("cannot track event: no app or tracking id")
	}

	app, err := msg.GetApp(cmd.apps)
	if err != nil {
		return err
	}

	usr, err := cmd.getUserOrNew(msg)
	if err != nil {
		return err
	}

	session, err := cmd.getSessionOrNew(msg, app, usr)
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

func (cmd *trackPageviewCmd) getSessionOrNew(rm *entity.RawMessage, app *entity.App, usr *entity.User) (*entity.Session, error) {
	session, err := rm.GetSession(cmd.sessions)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if session != nil {
		return session, nil
	}

	sid, err := uuid.Parse(rm.SessionID)
	if err != nil {
		return nil, err
	}

	session = app.NewSession(sid)
	session.UserID = usr.ID
	if err := cmd.sessions.Create(session); err != nil {
		return nil, err
	}

	return session, nil
}

func (cmd *trackPageviewCmd) getUserOrNew(rm *entity.RawMessage) (*entity.User, error) {
	usr, err := rm.GetUser(cmd.users)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if usr != nil {
		return usr, nil
	}

	usr = &entity.User{
		ID:        rm.UserID,
		Anonymous: true,
	}

	if err := cmd.users.Create(usr); err != nil {
		return nil, err
	}

	return usr, nil
}

func parseBool(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}

func parseTimestamp(s string) time.Time {
	t, _ := time.Parse(time.RFC3339, s)
	return t
}
