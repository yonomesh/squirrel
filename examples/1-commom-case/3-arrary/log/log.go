package log

import (
	"encoding/json"
)

type ArrayMsg []any

func (m ArrayMsg) MsgToString() (string, error) {
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
