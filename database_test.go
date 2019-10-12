package workday

import (
	"reflect"
	"testing"
)

func TestConnectToDBShouldReturnDBStruct(t *testing.T) {
	cfg := DBConfig{}
	db, err := ConnectDB(cfg)
	if err != nil {
		t.Errorf("Should get a DB struct but got an error: %v", err.Error())
	}
	if reflect.TypeOf(db).String() != "*sql.DB" {
		t.Errorf("Should get a DB struct but got: %v", reflect.TypeOf(db).String())
	}
}
