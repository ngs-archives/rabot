package app

import (
	"testing"
)

func TestFindPrefectureByID(t *testing.T) {
	res := FindPrefecture("JP26")
	if res.Name != "京都" {
		t.Errorf("Expected 京都 but got %v", res.Name)
	}
}

func TestFindPrefectureByName(t *testing.T) {
	res := FindPrefecture("京都")
	if res.ID != "JP26" {
		t.Errorf("Expected JP26 but got %v", res.ID)
	}
}
