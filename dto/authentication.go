package dto

import (
)

//Authentication represents struct that will contain data important for oauth2 and request sources
type Authentication struct {
	
	ClientID                 string   `json:"clientId"`
	ClientSecret             string   `json:"clientSecret"`
	TokenURL                 string   `json:"tokenUrl"`
	ApiHostName              string   `json:"apiHostName"`
	PoweredBy                string   `json:"poweredBy"`
}

// SetRequester sets requester to the command
func (auth *Authentication) SetClientID(id string) {
	auth.ClientID = id
}