package session

import (
	"encoding/base64"
	"encoding/binary"
)

func parseToken(token string) (int64, error) {
	b, err := base64.RawStdEncoding.DecodeString(token)
	if err != nil {
		return 0, err
	}

	num, _ := binary.Varint(b)

	return num, nil
}

func makeToken(sessionID int64) (string, error) {
	buf := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(buf, sessionID)

	return base64.RawStdEncoding.EncodeToString(buf), nil
}
