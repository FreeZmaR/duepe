package session

import (
	"context"
	"duepe/internal/domain/models"
	"errors"
	"reflect"
	"sync"
	"testing"
	"unsafe"
)

func TestManager_GetManager(t *testing.T) {
	type testCase struct {
		name      string
		callCount int
	}

	tests := []testCase{
		{
			name:      "Test 1",
			callCount: 1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg := sync.WaitGroup{}
			wg.Add(tt.callCount)

			mInstance := GetManager()
			ptr := unsafe.Pointer(mInstance)

			for i := range tt.callCount {
				go func(i int) {
					defer wg.Done()

					m := GetManager()

					if unsafe.Pointer(m) != ptr {
						t.Error("Manager instance is not the same: ", i)
					}
				}(i)
			}

			wg.Wait()
		})
	}
}

func TestManager_NewSessionAndDeleteSession(t *testing.T) {
	type testCase struct {
		name      string
		user      []*models.User
		assertIDS []int
	}

	tests := []testCase{
		{
			name: "One user",
			user: []*models.User{
				{Name: "User 1"},
			},
			assertIDS: []int{1},
		},
		{
			name: "Many users",
			user: []*models.User{
				{Name: "User 1"},
				{Name: "User 2"},
				{Name: "User 3"},
			},
			assertIDS: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mInstance := GetManager()

			defer func() {
				for _, id := range tt.assertIDS {
					mInstance.deleteSession(int64(id))
				}
			}()

			for i, user := range tt.user {
				session, err := mInstance.NewSession(context.Background(), user)
				if err != nil {
					t.Error(err)
				}

				if session.id != int64(tt.assertIDS[i]) {
					t.Errorf("Expected ID: %d, got: %d", tt.assertIDS[i], session.id)
				}
			}
		})
	}
}

func TestManager_GetSession(t *testing.T) {
	type testCase struct {
		name       string
		token      string
		users      []*models.User
		assertUser *models.User
		assertErr  error
	}

	tests := []testCase{
		{
			name:  "One User in session",
			token: mockToken(1),
			users: []*models.User{
				{Name: "User 1"},
			},
			assertUser: &models.User{Name: "User 1"},
		},
		{
			name:  "Many Users in session",
			token: mockToken(2),
			users: []*models.User{
				{Name: "User 1"},
				{Name: "User 2"},
				{Name: "User 3"},
			},
			assertUser: &models.User{Name: "User 2"},
		},
		{
			name:  "Invalid token",
			token: "test",
			users: []*models.User{
				{Name: "User 1"},
			},
			assertErr: ErrorInvalidToken,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mInstance := GetManager()
			sessions := make([]*Session, 0, len(tt.users))

			defer func() {
				for _, session := range sessions {
					mInstance.deleteSession(session.id)
				}
			}()

			for _, user := range tt.users {
				_, err := mInstance.NewSession(context.Background(), user)
				if err != nil && nil == tt.assertErr {
					t.Error("Error creating session: ", err)

					return
				}

				if tt.assertErr != nil && err != nil && !errors.Is(err, tt.assertErr) {
					t.Errorf(
						"Unexpected  NewSession error. Got: %s Expect: %s",
						err.Error(),
						tt.assertErr.Error(),
					)

					return
				}
			}

			session, err := mInstance.GetSession(context.Background(), tt.token)
			if err != nil && nil == tt.assertErr {
				t.Error("Unexpected GetSession error: ", err)

				return
			}

			sessions = append(sessions, session)

			if tt.assertErr != nil && err != nil {
				if !errors.Is(err, tt.assertErr) {
					t.Errorf(
						"Unexpected GetSession error. Got: %s Expect: %s",
						err.Error(),
						tt.assertErr.Error(),
					)

					return
				}

				return
			}

			if tt.assertErr != nil && nil == err {
				t.Error("Expected error, but got nil")

				return
			}

			if !reflect.DeepEqual(*session.user, *tt.assertUser) {
				t.Errorf(
					"Unexpected user. Got: %v Expect: %v",
					*session.user,
					*tt.assertUser,
				)
			}
		})
	}
}

func mockToken(id int64) string {
	token, _ := makeToken(id)

	return token
}
