package app

import (
	"testing"
)

func TestFindPrefectureByID(t *testing.T) {
	res, err := FindPrefecture("JP26")
	if res.Name != "京都" {
		t.Errorf("Expected 京都 but got %v. %v", res.Name, err)
	}
}

func TestFindPrefectureByName(t *testing.T) {
	res, err := FindPrefecture("京都")
	if res.ID != "JP26" {
		t.Errorf("Expected JP26 but got %v. %v", res.ID, err)
	}
}

func TestFindPrefectureNotFound(t *testing.T) {
	_, err := FindPrefecture("SFO")
	if err.Error() != "Could not find a prefecture with id or name SFO" {
		t.Errorf("Expected error but got `%v`", err)
	}
}
