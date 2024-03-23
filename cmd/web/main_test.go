package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	db, err := run()
	if err != nil {
		t.Error("run() failed")
	}

	defer db.SQL.Close()
}
