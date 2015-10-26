package scream

import (
	"encoding/json"
)

//Config is struct to store configuration parameters
type Config struct {
	Address string
	Key     string
}

//Cfg is runtime configuration instance
var Cfg Config

type notification struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Key     string `json:"key"`
}

func (n *notification) ToJSON() (string, error) {
	b, err := json.Marshal(n)
	return string(b), err
}
