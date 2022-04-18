package confs

import (
	"testing"
)

func TestMigrateUsersOrgs(t *testing.T) {
	db, _ := ConnectDB("127.0.0.1", "postgres", "postgres", "postgres", "5432")
	MigrateUsersOrgs(db)

	toTest := []string{"users", "organizations", "orgs_invits", "ahah"}

	for _, test := range toTest {
		table_check := db.Raw("select * from " + test + ";")
		if table_check != nil {
			t.Log(table_check)
			t.Fail()
		}
	}
}
