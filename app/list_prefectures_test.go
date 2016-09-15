package app

import (
	"strings"
	"testing"
)

func TestListPrefectures(t *testing.T) {
	app := &App{}
	table := app.ListPrefectures()
	t.Log(table)
	if c := strings.Count(table, "\n"); c != 52 {
		t.Errorf("Expected \b%v\b to have 52 line breaks but got %d", table, c)
	}

}
