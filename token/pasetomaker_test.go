package token

import (
	"testing"
	"time"

	"github.com/kelvin950/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {

	maker, err := NewPasetoMaker(util.RandomString(32))

	if err != nil {

		t.Errorf("failed %s", err)
	}

	username := util.RandomString(23)
	duration := time.Minute

	token, err := maker.CreateToken(username, duration)
	expiredAt := time.Now().Add(duration)
	if err != nil {
		t.Errorf("failed %s", err)
	}

	if token == "" {
		t.Error("token failed to create")
	}

	Payload, err := maker.VerifyToken(token)

	if err != nil {
		t.Errorf("failed to verify %s", err)
	}

	if Payload == nil {
		t.Error("failed to verify")
	}

	require.NotEmpty(t, Payload)
	require.Equal(t, username, Payload.Username)
	require.WithinDuration(t, expiredAt, Payload.ExpiredAt, time.Second)

}