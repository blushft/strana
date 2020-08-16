package command

import (
	"strconv"
	"time"

	"github.com/blushft/strana/domain/entity"
	"github.com/blushft/strana/pkg/event"
	"github.com/blushft/strana/platform/store/ent"
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

func (t *tracker) getUserOrNew(evt *event.Event) (*entity.User, error) {
	usr, err := t.users.Get(evt.UserID)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if usr != nil {
		return usr, nil
	}

	usr = &entity.User{
		ID:        evt.UserID,
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
