package session

import (
	"context"
	"duepe/internal/domain/models"
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
