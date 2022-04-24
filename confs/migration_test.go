package confs

import (
	"testing"
)

func TestMigrateUsersOrgs(t *testing.T) {
	db, _ := ConnectDB("127.0.0.1", "mysql", "mysql", "gateward", "3306")
	MigrateUsersOrgs(db)

	toTest := []string{"users", "organizations", "org_invits"}

	for _, test := range toTest {
		rows, table_check := db.Raw("select * from " + test + ";").Rows()
		defer rows.Close()
		if table_check != nil {
			t.Log("migration is not ok")
			t.Fail()
		}
	}

	_, table_check := db.Raw("select * from unexisting_table;").Rows()
	if table_check == nil {
		t.Log("not migrated table is working")
		t.Fail()
	}
}
