package command

import (
	"strconv"
	"time"

	"github.com/blushft/strana/domain/entity"
	"github.com/blushft/strana/platform/store/ent"
	"github.com/google/uuid"
)

type tracker struct {
	apps     entity.AppReporter
	sessions entity.SessionManager
	users    entity.UserManager
}

func newTracker(ar entity.AppReporter,
	ses entity.SessionManager,
	usr entity.UserManager,
) *tracker {
	return &tracker{
		apps:     ar,
		sessions: ses,
		users:    usr,
	}
}

func (t *tracker) getSessionOrNew(rm *entity.RawMessage, app *entity.App, usr *entity.User) (*entity.Session, error) {
	session, err := rm.GetSession(t.sessions)
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
	if err := t.sessions.Create(session); err != nil {
		return nil, err
	}

	return session, nil
}

func (t *tracker) getUserOrNew(rm *entity.RawMessage) (*entity.User, error) {
	usr, err := rm.GetUser(t.users)
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

	if err := t.users.Create(usr); err != nil {
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
