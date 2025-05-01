package hasher

import (
	"reflect"
	"testing"
)

func TestHashValue_String(t *testing.T) {
	h := NewHasher(MD5, SHA256, SHA512, SHA3)
	value := "test"
	results, err := h.HashValue(value)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 4 {
		t.Errorf("expected 4 results, got %d", len(results))
	}
	for algorithm, hash := range results {
		if hash == "" {
			t.Errorf("expected non-empty hash for %s", algorithm)
		}
	}
}

func TestHashValue_Int(t *testing.T) {
	h := NewHasher(SHA256)
	value := 42
	results, err := h.HashValue(value)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Errorf("expected 1 result, got %d", len(results))
	}
}

func TestHashValue_Float64(t *testing.T) {
	h := NewHasher(SHA512)
	value := 3.14
	results, err := h.HashValue(value)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Errorf("expected 1 result, got %d", len(results))
	}
}

func TestHashValue_UnsupportedType(t *testing.T) {
	h := NewHasher(SHA256)
	_, err := h.HashValue([]byte("test"))
	if err == nil {
		t.Error("expected error for unsupported type, got nil")
	}
}

func TestHashValues(t *testing.T) {
	h := NewHasher(SHA256)
	values := []interface{}{1, "foo", 2.5}
	results, err := h.HashValues(values)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != len(values) {
		t.Errorf("expected %d results, got %d", len(values), len(results))
	}
	for i, r := range results {
		if len(r.Hashes) != 1 {
			t.Errorf("expected 1 hash result for value %d, got %d", i, len(r.Hashes))
		}
		if !reflect.DeepEqual(r.Original, values[i]) {
			t.Errorf("expected original value to be %v, got %v", values[i], r.Original)
		}
	}
}
