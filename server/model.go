package server

import "encoding/json"

type relayProperties struct {
	State int `json:"state"`
}

func ParseRequestBody(body []byte) (relayProperties, error) {
	var r relayProperties
	if err := json.Unmarshal(body, &r); err != nil {
		return relayProperties{}, err
	}
	return r, nil
}

func MarshalReponseBody(r relayProperties) ([]byte, error) {
	return json.Marshal(r)
}
