package session

import "testing"

func TestManagerController(t *testing.T) {
	type testCase struct {
		name         string
		getCount     int
		putIDS       []int
		assertLastID int
	}

	testCases := []testCase{
		{
			name:         "Get 1 ID",
			getCount:     1,
			assertLastID: 2,
		},
		{
			name:         "Get 1 ID and put it back",
			getCount:     1,
			putIDS:       []int{1},
			assertLastID: 1,
		},
		{
			name:         "Get 2 IDs and put them back",
			getCount:     2,
			putIDS:       []int{1, 2},
			assertLastID: 1,
		},
		{
			name:         "Get 2 IDs and put 1 back",
			getCount:     2,
			putIDS:       []int{1},
			assertLastID: 1,
		},
		{
			name:         "Get 2 IDs and put last back",
			getCount:     2,
			putIDS:       []int{2},
			assertLastID: 2,
		},
		{
			name:         "Get 2 IDS",
			getCount:     2,
			assertLastID: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			controller := newManagerController()

			for i := 0; i < tc.getCount; i++ {
				controller.getID()
			}

			for _, id := range tc.putIDS {
				controller.putID(int64(id))
			}

			if id := controller.getID(); id != int64(tc.assertLastID) {
				t.Errorf("Expected last ID to be %d, but got %d", tc.assertLastID, id)
			}
		})
	}
}
