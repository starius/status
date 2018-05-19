package status

import (
	"fmt"
	"testing"
)

func TestPropagation(t *testing.T) {
	err := Format("file not found")
	err = AttachCode(403, err)
	err = Format("can't open database: %v", err)
	if err.Error() != "can't open database: file not found" {
		t.Errorf("Bad error message: %q", err)
	}
	if Code(err) != 403 {
		t.Errorf("Bad error code: %d (want %d)", Code(err), 403)
	}
}

func TestWithCode(t *testing.T) {
	err := WithCode(403, "file not found: %s", "1.txt")
	err = Format("can't open database: %v", err)
	if Code(err) != 403 {
		t.Errorf("Bad error code: %d (want %d)", Code(err), 403)
	}
}

func TestBreak(t *testing.T) {
	err := WithCode(403, "file not found: %s", "1.txt")
	err = fmt.Errorf("can't open database: %v", err)
	if Code(err) != 0 {
		t.Errorf("Bad error code: %d (want %d)", Code(err), 0)
	}
}
