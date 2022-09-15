package configdir

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"
)

func TestIsDir(t *testing.T) {
	d, err := dir()
	if err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(d)
	if err != nil {
		t.Fatal(err)
	}
	if !info.IsDir() {
		t.Fatalf("config path %q is not a directory", d)
	}
}

func tempDir() (string, error) {
	b := make([]byte, 10)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return "configdir-test-" + hex.EncodeToString(b), nil
}

func TestCreateDir(t *testing.T) {
	// Get temp name
	name, err := tempDir()
	if err != nil {
		t.Fatal(err)
	}

	// Create a new config directory
	d, err := New(name)
	if err != nil {
		t.Fatal(err)
	}
	if base := filepath.Base(d); base != name {
		t.Fatalf("got %q; want %q", base, name)
	}
	defer os.Remove(d)
	info, err := os.Stat(d)
	if err != nil {
		t.Fatal(err)
	}
	if !info.IsDir() {
		t.Fatalf("created config path %q is not a directory", d)
	}
}
