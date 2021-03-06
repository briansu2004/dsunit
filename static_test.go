package dsunit_test

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/viant/dsunit"
	"time"
)

func TestTableMapping(t *testing.T) {
	dsunit.InitFromURL(t, "test/static/init.json")
	dsunit.PrepareFor(t, "static", "test/static/data", "mapping")
	//business test logic comes here
	dsunit.ExpectFor(t, "static", dsunit.SnapshotDatasetCheckPolicy, "test/static/data", "mapping")

}

func TestTableMappingRemote(t *testing.T) {

	go func() {
		dsunit.StartServer("8877")
	}()
	time.Sleep(time.Second)
	tester := dsunit.NewRemoveTester("http://127.0.0.1:8877")
	tester.InitFromURL(t, "test/static/init.json")
	tester.PrepareFor(t, "static", "test/static/data", "mapping")
	//business test logic comes here
	tester.ExpectFor(t, "static", dsunit.SnapshotDatasetCheckPolicy, "test/static/data", "mapping")

}
