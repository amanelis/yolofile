package helpers

import (
	"testing"
)

func TestYolofileExists(t *testing.T) {
	if YolofileExists() {
		t.Fail()
	}
}

func TestFileExists(t *testing.T) {
	if !FileExists("file.go") {
		t.Fail()
	}
}

func TestFileExistsNot(t *testing.T) {
	if FileExists("main.go") {
		t.Fail()
	}
}
