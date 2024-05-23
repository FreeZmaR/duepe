package session

import "context"

type (
	sessionCTXKey struct{}
)

func putSessionToContext(ctx context.Context, session *Session) context.Context {
	return context.WithValue(ctx, sessionCTXKey{}, session)
}

func GetSessionFromContext(ctx context.Context) (*Session, bool) {
	session, ok := ctx.Value(sessionCTXKey{}).(*Session)

	return session, ok
}
