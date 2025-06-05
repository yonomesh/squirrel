package log

import (
	"encoding/json"
)

type MapMsg map[string]any

func (m MapMsg) MsgToString() (string, error) {
	result, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

type StringExtra string

func (e StringExtra) ExtraToString() (string, error) {
	return string(e), nil
}
