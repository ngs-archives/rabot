package app

import (
	"strings"
	"testing"
)

func TestListStations(t *testing.T) {
	_SetupMockHTTP()
	app := &App{}
	table := app.ListStations("東京")
	t.Log(table)
	if c := strings.Count(table, "\n"); c != 18 {
		t.Errorf("Expected \b%v\b to have 18 line breaks but got %d", table, c)
	}
}

func TestListStationsNotFound(t *testing.T) {
	_SetupMockHTTP()
	app := &App{}
	table := app.ListStations("SFO")
	if table != "Could not find a prefecture with id or name SFO" {
		t.Errorf("Expected error but got `%v`", table)
	}
}
