package models

type Device struct {
	Id      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Brand   string `json:"brand,omitempty"`
	Created string `json:"created,omitempty"`
}
