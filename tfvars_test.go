package tfvars

import (
	"path/filepath"
	"testing"
)

func TestRead(t *testing.T) {
	matches, err := filepath.Glob(filepath.Join("test-data", "*.tfvars"))
	if err != nil {
		t.Fatalf("failed to locate any tfvars files: %v", err)
	}

	for _, path := range matches {
		t.Run(path, func(t *testing.T) {
			tfvars, err := New(path)
			if err != nil {
				t.Fatal(err)
			}
			for _, key := range tfvars.Keys() {
				t.Logf("%s: %s", key, tfvars.Get(key))
			}
		})
	}
}
