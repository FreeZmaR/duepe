package session

import (
	"duepe/internal/domain/models"
	"errors"
	"sync/atomic"
)

var (
	ErrSessionIsPurged = errors.New("session is purged")
)

type Session struct {
	id              int64
	user            *models.User
	token           string
	deleteSessionFN deleteSessionFunc
	isPurged        atomic.Bool
}

type deleteSessionFunc func(id int64)

func newSession(id int64, user *models.User, token string, deleteSessionFN deleteSessionFunc) *Session {
	return &Session{
		id:              id,
		user:            user,
		token:           token,
		deleteSessionFN: deleteSessionFN,
	}
}

func (s *Session) User() *models.User {
	return s.user
}

func (s *Session) Token() string {
	return s.token
}

func (s *Session) Purge() error {
	if s.isPurged.Load() == true {
		return ErrSessionIsPurged
	}

	s.deleteSessionFN(s.id)
	s.isPurged.Store(true)

	return nil
}
