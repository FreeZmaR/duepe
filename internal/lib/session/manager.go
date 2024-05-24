package session

import (
	"context"
	"duepe/internal/domain/models"
	"errors"
	"sync"
)

var (
	ErrNotFound = errors.New("not found session")
)

var (
	managerInstance *Manager
	initFN          sync.Once
)

type Manager struct {
	container  *sync.Map
	controller *managerController
}

func GetManager() *Manager {
	initFN.Do(func() {
		managerInstance = newManager()
	})

	return managerInstance
}

func newManager() *Manager {
	return &Manager{
		container:  new(sync.Map),
		controller: newManagerController(),
	}
}

func (s *Manager) GetSession(_ context.Context, token string) (*Session, error) {
	id, err := parseToken(token)
	if err != nil {
		return nil, err
	}

	rawSession, ok := s.container.Load(id)
	if !ok {
		return nil, ErrNotFound
	}

	session, ok := rawSession.(*Session)
	if !ok {
		s.deleteSession(id)

		return nil, ErrNotFound
	}

	return session, nil
}

func (s *Manager) NewSession(_ context.Context, user *models.User) (*Session, error) {
	id := s.controller.getID()

	token, err := makeToken(id)
	if err != nil {
		return nil, err
	}

	newUser := *user

	session := newSession(id, &newUser, token, s.deleteSession)
	s.container.Store(id, session)

	return session, nil
}

func (s *Manager) SaveSessionToContext(ctx context.Context, session *Session) context.Context {
	return putSessionToContext(ctx, session)
}

func (s *Manager) deleteSession(id int64) {
	s.container.Delete(id)
	s.controller.putID(id)
}
