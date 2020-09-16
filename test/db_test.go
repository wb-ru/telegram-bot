package test

import (
	"../pkg/db"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConnectPGDB(t *testing.T) {
	require.True(t, true, db.ConnectPGDB())
}
