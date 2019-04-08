package session

import (
	"github.com/gorilla/sessions"
)

var (
	store *sessions.CookieStore
)

func NewCookieStore(key []byte) *sessions.CookieStore {
	store = sessions.NewCookieStore(key)
	return store
}
