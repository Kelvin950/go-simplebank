package api

import (
	"testing"
	"time"

	"github.com/kelvin950/simplebank/db"
	"github.com/kelvin950/simplebank/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T ,store db.Store)*Server{

 
	config := util.Config{
		TokenSymmetricKey: util.RandomString(32) ,
		AcessTokenDuration: time.Minute,
	}


	server , err:=  NewServer(config , store) 

	require.NoError(t  , err) 

	
	return server
}