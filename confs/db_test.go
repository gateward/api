package confs

import (
	"testing"
)

func TestConnectDB(t *testing.T) {
	_, err := ConnectDB("127.0.0.1", "mysql", "mysql", "gateward", "3306")

	if err != nil {
		t.Log("failed to connect database")
		t.Fail()
	}

	_, err = ConnectDB("127.0.0.1", "wrong", "wrong", "wrong", "3306")

	if err == nil {
		t.Log("wrong connection is ok")
		t.Fail()
	}
}
