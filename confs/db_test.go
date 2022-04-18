package confs

import (
	"testing"
)

func TestConnectDB(t *testing.T) {
	_, err := ConnectDB("127.0.0.1", "postgres", "postgres", "postgres", "5432")

	if err != nil {
		t.Log("failed to connect database")
		t.Fail()
	}
}
