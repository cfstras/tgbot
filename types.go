package tgbot

import (
	"encoding/json"
)

type TGResponse struct {
	Ok          bool
	Description string `json:",omitempty"`
	Result      json.RawMessage
}

type TGUser struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:",omitempty"`
}
