package http

import (
	"net/http"

	"howett.net/spectre"
)

// TODO(DH) This file should be split up and have its comments cleaned.

type UserService interface {
	GetUserForRequest(r *http.Request) spectre.User
}

type PermitterProvider interface {
	GetPermitterForRequest(r *http.Request) spectre.Permitter
}