package session

import (
	"testing"
)

type tokenSet struct {
	name  string
	token string
	id    int64
}

func TestParseToken(t *testing.T) {
	type testCase struct {
		name   string
		token  string
		assert int64
	}

	set := getTokenSet()

	tt := make([]testCase, 0, len(set))
	for _, s := range set {
		tt = append(tt, testCase{
			name:   "Token set:" + s.name,
			token:  s.token,
			assert: s.id,
		})
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseToken(tc.token)
			if err != nil {
				t.Errorf("unexpected error: %v", err)

				return
			}

			if got != tc.assert {
				t.Errorf("expected %d, got %d", tc.assert, got)
			}
		})
	}
}

func TestMakeToken(t *testing.T) {
	type testCase struct {
		name   string
		id     int64
		assert string
	}

	set := getTokenSet()

	tt := make([]testCase, 0, len(set))
	for _, s := range set {
		tt = append(tt, testCase{
			name:   "Token set:" + s.name,
			id:     s.id,
			assert: s.token,
		})

	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := makeToken(tc.id)
			if err != nil {
				t.Errorf("unexpected error: %v", err)

				return
			}

			if got != tc.assert {
				t.Errorf("expected %s, got %s", tc.assert, got)
			}
		})
	}
}

func getTokenSet() []tokenSet {
	return []tokenSet{
		{
			name:  "zero value",
			token: "AAAAAAAAAAAAAA",
			id:    0,
		},
		{
			name:  "minimal value",
			token: "AgAAAAAAAAAAAA",
			id:    1,
		},
		{
			name:  "middle value",
			token: "mrYQAAAAAAAAAA",
			id:    134541,
		},
		{
			name:  "big value",
			token: "mqu2tOoHAAAAAA",
			id:    134541134541,
		},
	}
}
