package elements

import (
	"testing"
)

func TestWebMSaver(t *testing.T) {
	saver := NewWebmSaver(WebmSaverConfig{
		ID:   "id",
		Path: "./",
	})
	saver.Close()
}
