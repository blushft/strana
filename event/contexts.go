package event

import (
	"encoding/json"
	"sort"
)

type Contexts map[string]Context

func (c Contexts) Get(ct ContextType) (Context, bool) {
	v, ok := c[string(ct)]
	return v, ok
}

func (c Contexts) Bind(v interface{}) {}

func (c Contexts) MarshalJSON() ([]byte, error) {
	cm := make(map[string]interface{}, len(c))
	for t, c := range c {
		cm[string(t)] = c.Values()
	}

	return json.Marshal(cm)
}

func (c *Contexts) UnmarshalJSON(b []byte) error {
	tc := make(Contexts)
	cm := make(map[string]json.RawMessage)

	if err := json.Unmarshal(b, &cm); err != nil {
		return err
	}

	for t, ce := range cm {
		nc, err := decodeContext(t, ce)
		if err != nil {
			return err
		}

		tc[t] = nc
	}

	*c = tc

	return nil
}

func (c Contexts) Map() map[string]interface{} {
	m := make(map[string]interface{}, len(c))
	for k, v := range c {
		m[k] = v
	}

	return m
}

func (c Contexts) List() []string {
	l := make([]string, 0, len(c))
	for k := range c {
		l = append(l, k)
	}

	return sort.StringSlice(l)
}

func (c Contexts) Visit(fn func(ctx Context)) {
	for _, ctx := range c {
		fn(ctx)
	}
}

type ContextIterator interface {
	First() Context
	Next() Context
}

type ctxIter struct {
	v    Context
	next *ctxIter
}

type ctxList struct {
	head    *ctxIter
	current *ctxIter
}

func (l *ctxList) add(ctx Context) {
	i := &ctxIter{
		v: ctx,
	}

	if l.head == nil {
		l.head = i
	} else {
		cur := l.head
		for cur.next != nil {
			cur = cur.next
		}

		cur.next = i
	}
}

func (l *ctxList) First() Context {
	l.current = l.head

	if l.head == nil {
		return nil
	}

	return l.head.v
}

func (l *ctxList) Next() Context {
	if l.current.next == nil {
		return nil
	}

	l.current = l.current.next
	return l.current.v
}

func (c Contexts) Iter() ContextIterator {
	ll := &ctxList{}

	cl := c.List()

	for _, cn := range cl {
		ll.add(c[cn])
	}

	return ll
}
