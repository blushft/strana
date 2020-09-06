package event

import "encoding/json"

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
