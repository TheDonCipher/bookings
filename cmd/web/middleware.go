package main

import (
	"net/http"
)

//TODO: Add nosurf XSS protection and session management

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
