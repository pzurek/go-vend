package vend

import (
	"net/http"
)

const (
	baseUrlString = "vendhq.com/api/"
)

// Vend client which talks to the Vend API
type Client struct {

	// HTTP client
	client *http.Client

	Config  *Config
	Users   *[]User
	Outlets *[]Outlet
}
